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
func HasFuncVideoEncoderConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Configure
//go:noescape
func FuncVideoEncoderConfigure(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderEncode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode
//go:noescape
func FuncVideoEncoderEncode(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderEncode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Encode1
//go:noescape
func FuncVideoEncoderEncode1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Flush
//go:noescape
func FuncVideoEncoderFlush(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Reset
//go:noescape
func FuncVideoEncoderReset(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_Close
//go:noescape
func FuncVideoEncoderClose(this js.Ref, fn unsafe.Pointer)

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
func HasFuncVideoEncoderIsConfigSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoEncoder_IsConfigSupported
//go:noescape
func FuncVideoEncoderIsConfigSupported(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_compressed_texture_astcGetSupportedProfiles(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_compressed_texture_astc_GetSupportedProfiles
//go:noescape
func FuncWEBGL_compressed_texture_astcGetSupportedProfiles(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_debug_shadersGetTranslatedShaderSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_debug_shaders_GetTranslatedShaderSource
//go:noescape
func FuncWEBGL_debug_shadersGetTranslatedShaderSource(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_draw_buffersDrawBuffersWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_buffers_DrawBuffersWEBGL
//go:noescape
func FuncWEBGL_draw_buffersDrawBuffersWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func FuncWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_draw_instanced_base_vertex_base_instance_DrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func FuncWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_lose_contextLoseContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_LoseContext
//go:noescape
func FuncWEBGL_lose_contextLoseContext(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_lose_contextRestoreContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_lose_context_RestoreContext
//go:noescape
func FuncWEBGL_lose_contextRestoreContext(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_drawMultiDrawArraysWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysWEBGL
//go:noescape
func FuncWEBGL_multi_drawMultiDrawArraysWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_drawMultiDrawElementsWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsWEBGL
//go:noescape
func FuncWEBGL_multi_drawMultiDrawElementsWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawArraysInstancedWEBGL
//go:noescape
func FuncWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_MultiDrawElementsInstancedWEBGL
//go:noescape
func FuncWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawArraysInstancedBaseInstanceWEBGL
//go:noescape
func FuncWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_multi_draw_instanced_base_vertex_base_instance_MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL
//go:noescape
func FuncWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWEBGL_provoking_vertexProvokingVertexWEBGL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WEBGL_provoking_vertex_ProvokingVertexWEBGL
//go:noescape
func FuncWEBGL_provoking_vertexProvokingVertexWEBGL(this js.Ref, fn unsafe.Pointer)

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
func HasFuncWebAssemblyValidate() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Validate
//go:noescape
func FuncWebAssemblyValidate(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_Validate
//go:noescape
func CallWebAssemblyValidate(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Validate
//go:noescape
func TryWebAssemblyValidate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Compile
//go:noescape
func HasFuncWebAssemblyCompile() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Compile
//go:noescape
func FuncWebAssemblyCompile(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_Compile
//go:noescape
func CallWebAssemblyCompile(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Compile
//go:noescape
func TryWebAssemblyCompile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Instantiate
//go:noescape
func HasFuncWebAssemblyInstantiate() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate
//go:noescape
func FuncWebAssemblyInstantiate(fn unsafe.Pointer)

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
	importObject js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Instantiate1
//go:noescape
func HasFuncWebAssemblyInstantiate1() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate1
//go:noescape
func FuncWebAssemblyInstantiate1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_Instantiate1
//go:noescape
func CallWebAssemblyInstantiate1(
	retPtr unsafe.Pointer,
	bytes js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate1
//go:noescape
func TryWebAssemblyInstantiate1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bytes js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Instantiate2
//go:noescape
func HasFuncWebAssemblyInstantiate2() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate2
//go:noescape
func FuncWebAssemblyInstantiate2(fn unsafe.Pointer)

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
	importObject js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_Instantiate3
//go:noescape
func HasFuncWebAssemblyInstantiate3() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_Instantiate3
//go:noescape
func FuncWebAssemblyInstantiate3(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_Instantiate3
//go:noescape
func CallWebAssemblyInstantiate3(
	retPtr unsafe.Pointer,
	moduleObject js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_Instantiate3
//go:noescape
func TryWebAssemblyInstantiate3(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleObject js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_CompileStreaming
//go:noescape
func HasFuncWebAssemblyCompileStreaming() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_CompileStreaming
//go:noescape
func FuncWebAssemblyCompileStreaming(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_CompileStreaming
//go:noescape
func CallWebAssemblyCompileStreaming(
	retPtr unsafe.Pointer,
	source js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_CompileStreaming
//go:noescape
func TryWebAssemblyCompileStreaming(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_InstantiateStreaming
//go:noescape
func HasFuncWebAssemblyInstantiateStreaming() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming
//go:noescape
func FuncWebAssemblyInstantiateStreaming(fn unsafe.Pointer)

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
	importObject js.Ref) js.Ref

//go:wasmimport plat/js/web has_WebAssembly_InstantiateStreaming1
//go:noescape
func HasFuncWebAssemblyInstantiateStreaming1() js.Ref

//go:wasmimport plat/js/web func_WebAssembly_InstantiateStreaming1
//go:noescape
func FuncWebAssemblyInstantiateStreaming1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WebAssembly_InstantiateStreaming1
//go:noescape
func CallWebAssemblyInstantiateStreaming1(
	retPtr unsafe.Pointer,
	source js.Ref)

//go:wasmimport plat/js/web try_WebAssembly_InstantiateStreaming1
//go:noescape
func TryWebAssemblyInstantiateStreaming1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref) js.Ref

//go:wasmimport plat/js/web store_WebGLContextEventInit
//go:noescape
func WebGLContextEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebGLContextEventInit
//go:noescape
func WebGLContextEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
