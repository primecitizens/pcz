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

//go:wasmimport plat/js/web constof_VideoTransferCharacteristics
//go:noescape

func ConstOfVideoTransferCharacteristics(str js.Ref) uint32

//go:wasmimport plat/js/web constof_VideoMatrixCoefficients
//go:noescape

func ConstOfVideoMatrixCoefficients(str js.Ref) uint32

//go:wasmimport plat/js/web store_VideoColorSpaceInit
//go:noescape

func VideoColorSpaceInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoColorSpaceInit
//go:noescape

func VideoColorSpaceInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoFrameBufferInit
//go:noescape

func VideoFrameBufferInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoFrameBufferInit
//go:noescape

func VideoFrameBufferInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoFrameCopyToOptions
//go:noescape

func VideoFrameCopyToOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoFrameCopyToOptions
//go:noescape

func VideoFrameCopyToOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_VideoColorSpace_VideoColorSpace
//go:noescape

func NewVideoColorSpaceByVideoColorSpace(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_VideoColorSpace_VideoColorSpace1
//go:noescape

func NewVideoColorSpaceByVideoColorSpace1() js.Ref

//go:wasmimport plat/js/web get_VideoColorSpace_Primaries
//go:noescape

func GetVideoColorSpacePrimaries(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoColorSpace_Transfer
//go:noescape

func GetVideoColorSpaceTransfer(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoColorSpace_Matrix
//go:noescape

func GetVideoColorSpaceMatrix(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoColorSpace_FullRange
//go:noescape

func GetVideoColorSpaceFullRange(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_VideoColorSpace_ToJSON
//go:noescape

func CallVideoColorSpaceToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoColorSpace_ToJSON
//go:noescape

func VideoColorSpaceToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_VideoFrame_VideoFrame
//go:noescape

func NewVideoFrameByVideoFrame(
	image js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_VideoFrame_VideoFrame1
//go:noescape

func NewVideoFrameByVideoFrame1(
	image js.Ref) js.Ref

//go:wasmimport plat/js/web new_VideoFrame_VideoFrame2
//go:noescape

func NewVideoFrameByVideoFrame2(
	data js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_VideoFrame_Format
//go:noescape

func GetVideoFrameFormat(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoFrame_CodedWidth
//go:noescape

func GetVideoFrameCodedWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoFrame_CodedHeight
//go:noescape

func GetVideoFrameCodedHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoFrame_CodedRect
//go:noescape

func GetVideoFrameCodedRect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoFrame_VisibleRect
//go:noescape

func GetVideoFrameVisibleRect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoFrame_DisplayWidth
//go:noescape

func GetVideoFrameDisplayWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoFrame_DisplayHeight
//go:noescape

func GetVideoFrameDisplayHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoFrame_Duration
//go:noescape

func GetVideoFrameDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_VideoFrame_Timestamp
//go:noescape

func GetVideoFrameTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web get_VideoFrame_ColorSpace
//go:noescape

func GetVideoFrameColorSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_Metadata
//go:noescape

func CallVideoFrameMetadata(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoFrame_Metadata
//go:noescape

func VideoFrameMetadataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_AllocationSize
//go:noescape

func CallVideoFrameAllocationSize(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web func_VideoFrame_AllocationSize
//go:noescape

func VideoFrameAllocationSizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_AllocationSize1
//go:noescape

func CallVideoFrameAllocationSize1(
	this js.Ref,
	ptr unsafe.Pointer,

) uint32

//go:wasmimport plat/js/web func_VideoFrame_AllocationSize1
//go:noescape

func VideoFrameAllocationSize1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_CopyTo
//go:noescape

func CallVideoFrameCopyTo(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_VideoFrame_CopyTo
//go:noescape

func VideoFrameCopyToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_CopyTo1
//go:noescape

func CallVideoFrameCopyTo1(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_VideoFrame_CopyTo1
//go:noescape

func VideoFrameCopyTo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_Clone
//go:noescape

func CallVideoFrameClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoFrame_Clone
//go:noescape

func VideoFrameCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoFrame_Close
//go:noescape

func CallVideoFrameClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_VideoFrame_Close
//go:noescape

func VideoFrameCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebGLRenderingContext_Canvas
//go:noescape

func GetWebGLRenderingContextCanvas(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebGLRenderingContext_DrawingBufferWidth
//go:noescape

func GetWebGLRenderingContextDrawingBufferWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_WebGLRenderingContext_DrawingBufferHeight
//go:noescape

func GetWebGLRenderingContextDrawingBufferHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_WebGLRenderingContext_DrawingBufferColorSpace
//go:noescape

func GetWebGLRenderingContextDrawingBufferColorSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_WebGLRenderingContext_DrawingBufferColorSpace
//go:noescape

func SetWebGLRenderingContextDrawingBufferColorSpace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_WebGLRenderingContext_UnpackColorSpace
//go:noescape

func GetWebGLRenderingContextUnpackColorSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_WebGLRenderingContext_UnpackColorSpace
//go:noescape

func SetWebGLRenderingContextUnpackColorSpace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetContextAttributes
//go:noescape

func CallWebGLRenderingContextGetContextAttributes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetContextAttributes
//go:noescape

func WebGLRenderingContextGetContextAttributesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsContextLost
//go:noescape

func CallWebGLRenderingContextIsContextLost(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsContextLost
//go:noescape

func WebGLRenderingContextIsContextLostFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetSupportedExtensions
//go:noescape

func CallWebGLRenderingContextGetSupportedExtensions(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetSupportedExtensions
//go:noescape

func WebGLRenderingContextGetSupportedExtensionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetExtension
//go:noescape

func CallWebGLRenderingContextGetExtension(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetExtension
//go:noescape

func WebGLRenderingContextGetExtensionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ActiveTexture
//go:noescape

func CallWebGLRenderingContextActiveTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ActiveTexture
//go:noescape

func WebGLRenderingContextActiveTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_AttachShader
//go:noescape

func CallWebGLRenderingContextAttachShader(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_AttachShader
//go:noescape

func WebGLRenderingContextAttachShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BindAttribLocation
//go:noescape

func CallWebGLRenderingContextBindAttribLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BindAttribLocation
//go:noescape

func WebGLRenderingContextBindAttribLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BindBuffer
//go:noescape

func CallWebGLRenderingContextBindBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BindBuffer
//go:noescape

func WebGLRenderingContextBindBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BindFramebuffer
//go:noescape

func CallWebGLRenderingContextBindFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BindFramebuffer
//go:noescape

func WebGLRenderingContextBindFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BindRenderbuffer
//go:noescape

func CallWebGLRenderingContextBindRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BindRenderbuffer
//go:noescape

func WebGLRenderingContextBindRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BindTexture
//go:noescape

func CallWebGLRenderingContextBindTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BindTexture
//go:noescape

func WebGLRenderingContextBindTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BlendColor
//go:noescape

func CallWebGLRenderingContextBlendColor(
	this js.Ref,
	ptr unsafe.Pointer,

	red float32,
	green float32,
	blue float32,
	alpha float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BlendColor
//go:noescape

func WebGLRenderingContextBlendColorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BlendEquation
//go:noescape

func CallWebGLRenderingContextBlendEquation(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BlendEquation
//go:noescape

func WebGLRenderingContextBlendEquationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BlendEquationSeparate
//go:noescape

func CallWebGLRenderingContextBlendEquationSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	modeRGB uint32,
	modeAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BlendEquationSeparate
//go:noescape

func WebGLRenderingContextBlendEquationSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BlendFunc
//go:noescape

func CallWebGLRenderingContextBlendFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	sfactor uint32,
	dfactor uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BlendFunc
//go:noescape

func WebGLRenderingContextBlendFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BlendFuncSeparate
//go:noescape

func CallWebGLRenderingContextBlendFuncSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	srcRGB uint32,
	dstRGB uint32,
	srcAlpha uint32,
	dstAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BlendFuncSeparate
//go:noescape

func WebGLRenderingContextBlendFuncSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CheckFramebufferStatus
//go:noescape

func CallWebGLRenderingContextCheckFramebufferStatus(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
) uint32

//go:wasmimport plat/js/web func_WebGLRenderingContext_CheckFramebufferStatus
//go:noescape

func WebGLRenderingContextCheckFramebufferStatusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Clear
//go:noescape

func CallWebGLRenderingContextClear(
	this js.Ref,
	ptr unsafe.Pointer,

	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Clear
//go:noescape

func WebGLRenderingContextClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ClearColor
//go:noescape

func CallWebGLRenderingContextClearColor(
	this js.Ref,
	ptr unsafe.Pointer,

	red float32,
	green float32,
	blue float32,
	alpha float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ClearColor
//go:noescape

func WebGLRenderingContextClearColorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ClearDepth
//go:noescape

func CallWebGLRenderingContextClearDepth(
	this js.Ref,
	ptr unsafe.Pointer,

	depth float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ClearDepth
//go:noescape

func WebGLRenderingContextClearDepthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ClearStencil
//go:noescape

func CallWebGLRenderingContextClearStencil(
	this js.Ref,
	ptr unsafe.Pointer,

	s int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ClearStencil
//go:noescape

func WebGLRenderingContextClearStencilFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ColorMask
//go:noescape

func CallWebGLRenderingContextColorMask(
	this js.Ref,
	ptr unsafe.Pointer,

	red js.Ref,
	green js.Ref,
	blue js.Ref,
	alpha js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ColorMask
//go:noescape

func WebGLRenderingContextColorMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CompileShader
//go:noescape

func CallWebGLRenderingContextCompileShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CompileShader
//go:noescape

func WebGLRenderingContextCompileShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CopyTexImage2D
//go:noescape

func CallWebGLRenderingContextCopyTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	x int32,
	y int32,
	width int32,
	height int32,
	border int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CopyTexImage2D
//go:noescape

func WebGLRenderingContextCopyTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CopyTexSubImage2D
//go:noescape

func CallWebGLRenderingContextCopyTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CopyTexSubImage2D
//go:noescape

func WebGLRenderingContextCopyTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateBuffer
//go:noescape

func CallWebGLRenderingContextCreateBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateBuffer
//go:noescape

func WebGLRenderingContextCreateBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateFramebuffer
//go:noescape

func CallWebGLRenderingContextCreateFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateFramebuffer
//go:noescape

func WebGLRenderingContextCreateFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateProgram
//go:noescape

func CallWebGLRenderingContextCreateProgram(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateProgram
//go:noescape

func WebGLRenderingContextCreateProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateRenderbuffer
//go:noescape

func CallWebGLRenderingContextCreateRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateRenderbuffer
//go:noescape

func WebGLRenderingContextCreateRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateShader
//go:noescape

func CallWebGLRenderingContextCreateShader(
	this js.Ref,
	ptr unsafe.Pointer,

	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateShader
//go:noescape

func WebGLRenderingContextCreateShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CreateTexture
//go:noescape

func CallWebGLRenderingContextCreateTexture(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CreateTexture
//go:noescape

func WebGLRenderingContextCreateTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CullFace
//go:noescape

func CallWebGLRenderingContextCullFace(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CullFace
//go:noescape

func WebGLRenderingContextCullFaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteBuffer
//go:noescape

func CallWebGLRenderingContextDeleteBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteBuffer
//go:noescape

func WebGLRenderingContextDeleteBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteFramebuffer
//go:noescape

func CallWebGLRenderingContextDeleteFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteFramebuffer
//go:noescape

func WebGLRenderingContextDeleteFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteProgram
//go:noescape

func CallWebGLRenderingContextDeleteProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteProgram
//go:noescape

func WebGLRenderingContextDeleteProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteRenderbuffer
//go:noescape

func CallWebGLRenderingContextDeleteRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteRenderbuffer
//go:noescape

func WebGLRenderingContextDeleteRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteShader
//go:noescape

func CallWebGLRenderingContextDeleteShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteShader
//go:noescape

func WebGLRenderingContextDeleteShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DeleteTexture
//go:noescape

func CallWebGLRenderingContextDeleteTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DeleteTexture
//go:noescape

func WebGLRenderingContextDeleteTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DepthFunc
//go:noescape

func CallWebGLRenderingContextDepthFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	fn uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DepthFunc
//go:noescape

func WebGLRenderingContextDepthFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DepthMask
//go:noescape

func CallWebGLRenderingContextDepthMask(
	this js.Ref,
	ptr unsafe.Pointer,

	flag js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DepthMask
//go:noescape

func WebGLRenderingContextDepthMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DepthRange
//go:noescape

func CallWebGLRenderingContextDepthRange(
	this js.Ref,
	ptr unsafe.Pointer,

	zNear float32,
	zFar float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DepthRange
//go:noescape

func WebGLRenderingContextDepthRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DetachShader
//go:noescape

func CallWebGLRenderingContextDetachShader(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DetachShader
//go:noescape

func WebGLRenderingContextDetachShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Disable
//go:noescape

func CallWebGLRenderingContextDisable(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Disable
//go:noescape

func WebGLRenderingContextDisableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DisableVertexAttribArray
//go:noescape

func CallWebGLRenderingContextDisableVertexAttribArray(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DisableVertexAttribArray
//go:noescape

func WebGLRenderingContextDisableVertexAttribArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DrawArrays
//go:noescape

func CallWebGLRenderingContextDrawArrays(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	first int32,
	count int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DrawArrays
//go:noescape

func WebGLRenderingContextDrawArraysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_DrawElements
//go:noescape

func CallWebGLRenderingContextDrawElements(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	count int32,
	typ uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_DrawElements
//go:noescape

func WebGLRenderingContextDrawElementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Enable
//go:noescape

func CallWebGLRenderingContextEnable(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Enable
//go:noescape

func WebGLRenderingContextEnableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_EnableVertexAttribArray
//go:noescape

func CallWebGLRenderingContextEnableVertexAttribArray(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_EnableVertexAttribArray
//go:noescape

func WebGLRenderingContextEnableVertexAttribArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Finish
//go:noescape

func CallWebGLRenderingContextFinish(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Finish
//go:noescape

func WebGLRenderingContextFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Flush
//go:noescape

func CallWebGLRenderingContextFlush(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Flush
//go:noescape

func WebGLRenderingContextFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_FramebufferRenderbuffer
//go:noescape

func CallWebGLRenderingContextFramebufferRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	renderbuffertarget uint32,
	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_FramebufferRenderbuffer
//go:noescape

func WebGLRenderingContextFramebufferRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_FramebufferTexture2D
//go:noescape

func CallWebGLRenderingContextFramebufferTexture2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	textarget uint32,
	texture js.Ref,
	level int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_FramebufferTexture2D
//go:noescape

func WebGLRenderingContextFramebufferTexture2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_FrontFace
//go:noescape

func CallWebGLRenderingContextFrontFace(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_FrontFace
//go:noescape

func WebGLRenderingContextFrontFaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GenerateMipmap
//go:noescape

func CallWebGLRenderingContextGenerateMipmap(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GenerateMipmap
//go:noescape

func WebGLRenderingContextGenerateMipmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetActiveAttrib
//go:noescape

func CallWebGLRenderingContextGetActiveAttrib(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetActiveAttrib
//go:noescape

func WebGLRenderingContextGetActiveAttribFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetActiveUniform
//go:noescape

func CallWebGLRenderingContextGetActiveUniform(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetActiveUniform
//go:noescape

func WebGLRenderingContextGetActiveUniformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetAttachedShaders
//go:noescape

func CallWebGLRenderingContextGetAttachedShaders(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetAttachedShaders
//go:noescape

func WebGLRenderingContextGetAttachedShadersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetAttribLocation
//go:noescape

func CallWebGLRenderingContextGetAttribLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	name js.Ref,
) int32

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetAttribLocation
//go:noescape

func WebGLRenderingContextGetAttribLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetBufferParameter
//go:noescape

func CallWebGLRenderingContextGetBufferParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetBufferParameter
//go:noescape

func WebGLRenderingContextGetBufferParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetParameter
//go:noescape

func CallWebGLRenderingContextGetParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetParameter
//go:noescape

func WebGLRenderingContextGetParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetError
//go:noescape

func CallWebGLRenderingContextGetError(
	this js.Ref,
	ptr unsafe.Pointer,

) uint32

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetError
//go:noescape

func WebGLRenderingContextGetErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetFramebufferAttachmentParameter
//go:noescape

func CallWebGLRenderingContextGetFramebufferAttachmentParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetFramebufferAttachmentParameter
//go:noescape

func WebGLRenderingContextGetFramebufferAttachmentParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetProgramParameter
//go:noescape

func CallWebGLRenderingContextGetProgramParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetProgramParameter
//go:noescape

func WebGLRenderingContextGetProgramParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetProgramInfoLog
//go:noescape

func CallWebGLRenderingContextGetProgramInfoLog(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetProgramInfoLog
//go:noescape

func WebGLRenderingContextGetProgramInfoLogFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetRenderbufferParameter
//go:noescape

func CallWebGLRenderingContextGetRenderbufferParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetRenderbufferParameter
//go:noescape

func WebGLRenderingContextGetRenderbufferParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetShaderParameter
//go:noescape

func CallWebGLRenderingContextGetShaderParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetShaderParameter
//go:noescape

func WebGLRenderingContextGetShaderParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetShaderPrecisionFormat
//go:noescape

func CallWebGLRenderingContextGetShaderPrecisionFormat(
	this js.Ref,
	ptr unsafe.Pointer,

	shadertype uint32,
	precisiontype uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetShaderPrecisionFormat
//go:noescape

func WebGLRenderingContextGetShaderPrecisionFormatFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetShaderInfoLog
//go:noescape

func CallWebGLRenderingContextGetShaderInfoLog(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetShaderInfoLog
//go:noescape

func WebGLRenderingContextGetShaderInfoLogFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetShaderSource
//go:noescape

func CallWebGLRenderingContextGetShaderSource(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetShaderSource
//go:noescape

func WebGLRenderingContextGetShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetTexParameter
//go:noescape

func CallWebGLRenderingContextGetTexParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetTexParameter
//go:noescape

func WebGLRenderingContextGetTexParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetUniform
//go:noescape

func CallWebGLRenderingContextGetUniform(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	location js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetUniform
//go:noescape

func WebGLRenderingContextGetUniformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetUniformLocation
//go:noescape

func CallWebGLRenderingContextGetUniformLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetUniformLocation
//go:noescape

func WebGLRenderingContextGetUniformLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetVertexAttrib
//go:noescape

func CallWebGLRenderingContextGetVertexAttrib(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetVertexAttrib
//go:noescape

func WebGLRenderingContextGetVertexAttribFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_GetVertexAttribOffset
//go:noescape

func CallWebGLRenderingContextGetVertexAttribOffset(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	pname uint32,
) int64

//go:wasmimport plat/js/web func_WebGLRenderingContext_GetVertexAttribOffset
//go:noescape

func WebGLRenderingContextGetVertexAttribOffsetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Hint
//go:noescape

func CallWebGLRenderingContextHint(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Hint
//go:noescape

func WebGLRenderingContextHintFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsBuffer
//go:noescape

func CallWebGLRenderingContextIsBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsBuffer
//go:noescape

func WebGLRenderingContextIsBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsEnabled
//go:noescape

func CallWebGLRenderingContextIsEnabled(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsEnabled
//go:noescape

func WebGLRenderingContextIsEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsFramebuffer
//go:noescape

func CallWebGLRenderingContextIsFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsFramebuffer
//go:noescape

func WebGLRenderingContextIsFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsProgram
//go:noescape

func CallWebGLRenderingContextIsProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsProgram
//go:noescape

func WebGLRenderingContextIsProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsRenderbuffer
//go:noescape

func CallWebGLRenderingContextIsRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsRenderbuffer
//go:noescape

func WebGLRenderingContextIsRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsShader
//go:noescape

func CallWebGLRenderingContextIsShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsShader
//go:noescape

func WebGLRenderingContextIsShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_IsTexture
//go:noescape

func CallWebGLRenderingContextIsTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_IsTexture
//go:noescape

func WebGLRenderingContextIsTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_LineWidth
//go:noescape

func CallWebGLRenderingContextLineWidth(
	this js.Ref,
	ptr unsafe.Pointer,

	width float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_LineWidth
//go:noescape

func WebGLRenderingContextLineWidthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_LinkProgram
//go:noescape

func CallWebGLRenderingContextLinkProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_LinkProgram
//go:noescape

func WebGLRenderingContextLinkProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_PixelStorei
//go:noescape

func CallWebGLRenderingContextPixelStorei(
	this js.Ref,
	ptr unsafe.Pointer,

	pname uint32,
	param int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_PixelStorei
//go:noescape

func WebGLRenderingContextPixelStoreiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_PolygonOffset
//go:noescape

func CallWebGLRenderingContextPolygonOffset(
	this js.Ref,
	ptr unsafe.Pointer,

	factor float32,
	units float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_PolygonOffset
//go:noescape

func WebGLRenderingContextPolygonOffsetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_RenderbufferStorage
//go:noescape

func CallWebGLRenderingContextRenderbufferStorage(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	internalformat uint32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_RenderbufferStorage
//go:noescape

func WebGLRenderingContextRenderbufferStorageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_SampleCoverage
//go:noescape

func CallWebGLRenderingContextSampleCoverage(
	this js.Ref,
	ptr unsafe.Pointer,

	value float32,
	invert js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_SampleCoverage
//go:noescape

func WebGLRenderingContextSampleCoverageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Scissor
//go:noescape

func CallWebGLRenderingContextScissor(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Scissor
//go:noescape

func WebGLRenderingContextScissorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ShaderSource
//go:noescape

func CallWebGLRenderingContextShaderSource(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ShaderSource
//go:noescape

func WebGLRenderingContextShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilFunc
//go:noescape

func CallWebGLRenderingContextStencilFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	fn uint32,
	ref int32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilFunc
//go:noescape

func WebGLRenderingContextStencilFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilFuncSeparate
//go:noescape

func CallWebGLRenderingContextStencilFuncSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	fn uint32,
	ref int32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilFuncSeparate
//go:noescape

func WebGLRenderingContextStencilFuncSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilMask
//go:noescape

func CallWebGLRenderingContextStencilMask(
	this js.Ref,
	ptr unsafe.Pointer,

	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilMask
//go:noescape

func WebGLRenderingContextStencilMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilMaskSeparate
//go:noescape

func CallWebGLRenderingContextStencilMaskSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilMaskSeparate
//go:noescape

func WebGLRenderingContextStencilMaskSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilOp
//go:noescape

func CallWebGLRenderingContextStencilOp(
	this js.Ref,
	ptr unsafe.Pointer,

	fail uint32,
	zfail uint32,
	zpass uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilOp
//go:noescape

func WebGLRenderingContextStencilOpFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_StencilOpSeparate
//go:noescape

func CallWebGLRenderingContextStencilOpSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	fail uint32,
	zfail uint32,
	zpass uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_StencilOpSeparate
//go:noescape

func WebGLRenderingContextStencilOpSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexParameterf
//go:noescape

func CallWebGLRenderingContextTexParameterf(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
	param float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexParameterf
//go:noescape

func WebGLRenderingContextTexParameterfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexParameteri
//go:noescape

func CallWebGLRenderingContextTexParameteri(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
	param int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexParameteri
//go:noescape

func WebGLRenderingContextTexParameteriFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform1f
//go:noescape

func CallWebGLRenderingContextUniform1f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform1f
//go:noescape

func WebGLRenderingContextUniform1fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform2f
//go:noescape

func CallWebGLRenderingContextUniform2f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform2f
//go:noescape

func WebGLRenderingContextUniform2fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform3f
//go:noescape

func CallWebGLRenderingContextUniform3f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform3f
//go:noescape

func WebGLRenderingContextUniform3fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform4f
//go:noescape

func CallWebGLRenderingContextUniform4f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
	z float32,
	w float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform4f
//go:noescape

func WebGLRenderingContextUniform4fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform1i
//go:noescape

func CallWebGLRenderingContextUniform1i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform1i
//go:noescape

func WebGLRenderingContextUniform1iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform2i
//go:noescape

func CallWebGLRenderingContextUniform2i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform2i
//go:noescape

func WebGLRenderingContextUniform2iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform3i
//go:noescape

func CallWebGLRenderingContextUniform3i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
	z int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform3i
//go:noescape

func WebGLRenderingContextUniform3iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform4i
//go:noescape

func CallWebGLRenderingContextUniform4i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
	z int32,
	w int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform4i
//go:noescape

func WebGLRenderingContextUniform4iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_UseProgram
//go:noescape

func CallWebGLRenderingContextUseProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_UseProgram
//go:noescape

func WebGLRenderingContextUseProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ValidateProgram
//go:noescape

func CallWebGLRenderingContextValidateProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ValidateProgram
//go:noescape

func WebGLRenderingContextValidateProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib1f
//go:noescape

func CallWebGLRenderingContextVertexAttrib1f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib1f
//go:noescape

func WebGLRenderingContextVertexAttrib1fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib2f
//go:noescape

func CallWebGLRenderingContextVertexAttrib2f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib2f
//go:noescape

func WebGLRenderingContextVertexAttrib2fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib3f
//go:noescape

func CallWebGLRenderingContextVertexAttrib3f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib3f
//go:noescape

func WebGLRenderingContextVertexAttrib3fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib4f
//go:noescape

func CallWebGLRenderingContextVertexAttrib4f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
	z float32,
	w float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib4f
//go:noescape

func WebGLRenderingContextVertexAttrib4fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib1fv
//go:noescape

func CallWebGLRenderingContextVertexAttrib1fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib1fv
//go:noescape

func WebGLRenderingContextVertexAttrib1fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib2fv
//go:noescape

func CallWebGLRenderingContextVertexAttrib2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib2fv
//go:noescape

func WebGLRenderingContextVertexAttrib2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib3fv
//go:noescape

func CallWebGLRenderingContextVertexAttrib3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib3fv
//go:noescape

func WebGLRenderingContextVertexAttrib3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttrib4fv
//go:noescape

func CallWebGLRenderingContextVertexAttrib4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttrib4fv
//go:noescape

func WebGLRenderingContextVertexAttrib4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_VertexAttribPointer
//go:noescape

func CallWebGLRenderingContextVertexAttribPointer(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	size int32,
	typ uint32,
	normalized js.Ref,
	stride int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_VertexAttribPointer
//go:noescape

func WebGLRenderingContextVertexAttribPointerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Viewport
//go:noescape

func CallWebGLRenderingContextViewport(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Viewport
//go:noescape

func WebGLRenderingContextViewportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_MakeXRCompatible
//go:noescape

func CallWebGLRenderingContextMakeXRCompatible(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_MakeXRCompatible
//go:noescape

func WebGLRenderingContextMakeXRCompatibleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BufferData
//go:noescape

func CallWebGLRenderingContextBufferData(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	size float64,
	usage uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BufferData
//go:noescape

func WebGLRenderingContextBufferDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BufferData1
//go:noescape

func CallWebGLRenderingContextBufferData1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	data js.Ref,
	usage uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BufferData1
//go:noescape

func WebGLRenderingContextBufferData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_BufferSubData
//go:noescape

func CallWebGLRenderingContextBufferSubData(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	offset float64,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_BufferSubData
//go:noescape

func WebGLRenderingContextBufferSubDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CompressedTexImage2D
//go:noescape

func CallWebGLRenderingContextCompressedTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	border int32,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CompressedTexImage2D
//go:noescape

func WebGLRenderingContextCompressedTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_CompressedTexSubImage2D
//go:noescape

func CallWebGLRenderingContextCompressedTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_CompressedTexSubImage2D
//go:noescape

func WebGLRenderingContextCompressedTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_ReadPixels
//go:noescape

func CallWebGLRenderingContextReadPixels(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	pixels js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_ReadPixels
//go:noescape

func WebGLRenderingContextReadPixelsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexImage2D
//go:noescape

func CallWebGLRenderingContextTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	border int32,
	format uint32,
	typ uint32,
	pixels js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexImage2D
//go:noescape

func WebGLRenderingContextTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexImage2D1
//go:noescape

func CallWebGLRenderingContextTexImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexImage2D1
//go:noescape

func WebGLRenderingContextTexImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexSubImage2D
//go:noescape

func CallWebGLRenderingContextTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	pixels js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexSubImage2D
//go:noescape

func WebGLRenderingContextTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_TexSubImage2D1
//go:noescape

func CallWebGLRenderingContextTexSubImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_TexSubImage2D1
//go:noescape

func WebGLRenderingContextTexSubImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform1fv
//go:noescape

func CallWebGLRenderingContextUniform1fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform1fv
//go:noescape

func WebGLRenderingContextUniform1fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform2fv
//go:noescape

func CallWebGLRenderingContextUniform2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform2fv
//go:noescape

func WebGLRenderingContextUniform2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform3fv
//go:noescape

func CallWebGLRenderingContextUniform3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform3fv
//go:noescape

func WebGLRenderingContextUniform3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform4fv
//go:noescape

func CallWebGLRenderingContextUniform4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform4fv
//go:noescape

func WebGLRenderingContextUniform4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform1iv
//go:noescape

func CallWebGLRenderingContextUniform1iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform1iv
//go:noescape

func WebGLRenderingContextUniform1ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform2iv
//go:noescape

func CallWebGLRenderingContextUniform2iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform2iv
//go:noescape

func WebGLRenderingContextUniform2ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform3iv
//go:noescape

func CallWebGLRenderingContextUniform3iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform3iv
//go:noescape

func WebGLRenderingContextUniform3ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_Uniform4iv
//go:noescape

func CallWebGLRenderingContextUniform4iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_Uniform4iv
//go:noescape

func WebGLRenderingContextUniform4ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_UniformMatrix2fv
//go:noescape

func CallWebGLRenderingContextUniformMatrix2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_UniformMatrix2fv
//go:noescape

func WebGLRenderingContextUniformMatrix2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_UniformMatrix3fv
//go:noescape

func CallWebGLRenderingContextUniformMatrix3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_UniformMatrix3fv
//go:noescape

func WebGLRenderingContextUniformMatrix3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGLRenderingContext_UniformMatrix4fv
//go:noescape

func CallWebGLRenderingContextUniformMatrix4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGLRenderingContext_UniformMatrix4fv
//go:noescape

func WebGLRenderingContextUniformMatrix4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebGL2RenderingContext_Canvas
//go:noescape

func GetWebGL2RenderingContextCanvas(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WebGL2RenderingContext_DrawingBufferWidth
//go:noescape

func GetWebGL2RenderingContextDrawingBufferWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_WebGL2RenderingContext_DrawingBufferHeight
//go:noescape

func GetWebGL2RenderingContextDrawingBufferHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_WebGL2RenderingContext_DrawingBufferColorSpace
//go:noescape

func GetWebGL2RenderingContextDrawingBufferColorSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_WebGL2RenderingContext_DrawingBufferColorSpace
//go:noescape

func SetWebGL2RenderingContextDrawingBufferColorSpace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_WebGL2RenderingContext_UnpackColorSpace
//go:noescape

func GetWebGL2RenderingContextUnpackColorSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_WebGL2RenderingContext_UnpackColorSpace
//go:noescape

func SetWebGL2RenderingContextUnpackColorSpace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferData
//go:noescape

func CallWebGL2RenderingContextBufferData(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	size float64,
	usage uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferData
//go:noescape

func WebGL2RenderingContextBufferDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferData1
//go:noescape

func CallWebGL2RenderingContextBufferData1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcData js.Ref,
	usage uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferData1
//go:noescape

func WebGL2RenderingContextBufferData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferSubData
//go:noescape

func CallWebGL2RenderingContextBufferSubData(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	dstByteOffset float64,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferSubData
//go:noescape

func WebGL2RenderingContextBufferSubDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferData2
//go:noescape

func CallWebGL2RenderingContextBufferData2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcData js.Ref,
	usage uint32,
	srcOffset uint32,
	length uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferData2
//go:noescape

func WebGL2RenderingContextBufferData2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferData3
//go:noescape

func CallWebGL2RenderingContextBufferData3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcData js.Ref,
	usage uint32,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferData3
//go:noescape

func WebGL2RenderingContextBufferData3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferSubData1
//go:noescape

func CallWebGL2RenderingContextBufferSubData1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	dstByteOffset float64,
	srcData js.Ref,
	srcOffset uint32,
	length uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferSubData1
//go:noescape

func WebGL2RenderingContextBufferSubData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BufferSubData2
//go:noescape

func CallWebGL2RenderingContextBufferSubData2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	dstByteOffset float64,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BufferSubData2
//go:noescape

func WebGL2RenderingContextBufferSubData2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage2D
//go:noescape

func CallWebGL2RenderingContextTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	border int32,
	format uint32,
	typ uint32,
	pixels js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage2D
//go:noescape

func WebGL2RenderingContextTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage2D1
//go:noescape

func CallWebGL2RenderingContextTexImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage2D1
//go:noescape

func WebGL2RenderingContextTexImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage2D
//go:noescape

func CallWebGL2RenderingContextTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	pixels js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage2D
//go:noescape

func WebGL2RenderingContextTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage2D1
//go:noescape

func CallWebGL2RenderingContextTexSubImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage2D1
//go:noescape

func WebGL2RenderingContextTexSubImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage2D2
//go:noescape

func CallWebGL2RenderingContextTexImage2D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	border int32,
	format uint32,
	typ uint32,
	pboOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage2D2
//go:noescape

func WebGL2RenderingContextTexImage2D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage2D3
//go:noescape

func CallWebGL2RenderingContextTexImage2D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	border int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage2D3
//go:noescape

func WebGL2RenderingContextTexImage2D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage2D4
//go:noescape

func CallWebGL2RenderingContextTexImage2D4(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	border int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage2D4
//go:noescape

func WebGL2RenderingContextTexImage2D4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage2D2
//go:noescape

func CallWebGL2RenderingContextTexSubImage2D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	pboOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage2D2
//go:noescape

func WebGL2RenderingContextTexSubImage2D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage2D3
//go:noescape

func CallWebGL2RenderingContextTexSubImage2D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage2D3
//go:noescape

func WebGL2RenderingContextTexSubImage2D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage2D4
//go:noescape

func CallWebGL2RenderingContextTexSubImage2D4(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage2D4
//go:noescape

func WebGL2RenderingContextTexSubImage2D4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage2D
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	border int32,
	imageSize int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage2D
//go:noescape

func WebGL2RenderingContextCompressedTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage2D1
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	border int32,
	srcData js.Ref,
	srcOffset uint32,
	srcLengthOverride uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage2D1
//go:noescape

func WebGL2RenderingContextCompressedTexImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage2D2
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage2D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	border int32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage2D2
//go:noescape

func WebGL2RenderingContextCompressedTexImage2D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage2D3
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage2D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	border int32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage2D3
//go:noescape

func WebGL2RenderingContextCompressedTexImage2D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage2D
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	imageSize int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage2D
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage2D1
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage2D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	srcData js.Ref,
	srcOffset uint32,
	srcLengthOverride uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage2D1
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage2D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage2D2
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage2D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage2D2
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage2D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage2D3
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage2D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	width int32,
	height int32,
	format uint32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage2D3
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage2D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1fv
//go:noescape

func CallWebGL2RenderingContextUniform1fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1fv
//go:noescape

func WebGL2RenderingContextUniform1fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1fv1
//go:noescape

func CallWebGL2RenderingContextUniform1fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1fv1
//go:noescape

func WebGL2RenderingContextUniform1fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1fv2
//go:noescape

func CallWebGL2RenderingContextUniform1fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1fv2
//go:noescape

func WebGL2RenderingContextUniform1fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2fv
//go:noescape

func CallWebGL2RenderingContextUniform2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2fv
//go:noescape

func WebGL2RenderingContextUniform2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2fv1
//go:noescape

func CallWebGL2RenderingContextUniform2fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2fv1
//go:noescape

func WebGL2RenderingContextUniform2fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2fv2
//go:noescape

func CallWebGL2RenderingContextUniform2fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2fv2
//go:noescape

func WebGL2RenderingContextUniform2fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3fv
//go:noescape

func CallWebGL2RenderingContextUniform3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3fv
//go:noescape

func WebGL2RenderingContextUniform3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3fv1
//go:noescape

func CallWebGL2RenderingContextUniform3fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3fv1
//go:noescape

func WebGL2RenderingContextUniform3fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3fv2
//go:noescape

func CallWebGL2RenderingContextUniform3fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3fv2
//go:noescape

func WebGL2RenderingContextUniform3fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4fv
//go:noescape

func CallWebGL2RenderingContextUniform4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4fv
//go:noescape

func WebGL2RenderingContextUniform4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4fv1
//go:noescape

func CallWebGL2RenderingContextUniform4fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4fv1
//go:noescape

func WebGL2RenderingContextUniform4fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4fv2
//go:noescape

func CallWebGL2RenderingContextUniform4fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4fv2
//go:noescape

func WebGL2RenderingContextUniform4fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1iv
//go:noescape

func CallWebGL2RenderingContextUniform1iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1iv
//go:noescape

func WebGL2RenderingContextUniform1ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1iv1
//go:noescape

func CallWebGL2RenderingContextUniform1iv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1iv1
//go:noescape

func WebGL2RenderingContextUniform1iv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1iv2
//go:noescape

func CallWebGL2RenderingContextUniform1iv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1iv2
//go:noescape

func WebGL2RenderingContextUniform1iv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2iv
//go:noescape

func CallWebGL2RenderingContextUniform2iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2iv
//go:noescape

func WebGL2RenderingContextUniform2ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2iv1
//go:noescape

func CallWebGL2RenderingContextUniform2iv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2iv1
//go:noescape

func WebGL2RenderingContextUniform2iv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2iv2
//go:noescape

func CallWebGL2RenderingContextUniform2iv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2iv2
//go:noescape

func WebGL2RenderingContextUniform2iv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3iv
//go:noescape

func CallWebGL2RenderingContextUniform3iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3iv
//go:noescape

func WebGL2RenderingContextUniform3ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3iv1
//go:noescape

func CallWebGL2RenderingContextUniform3iv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3iv1
//go:noescape

func WebGL2RenderingContextUniform3iv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3iv2
//go:noescape

func CallWebGL2RenderingContextUniform3iv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3iv2
//go:noescape

func WebGL2RenderingContextUniform3iv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4iv
//go:noescape

func CallWebGL2RenderingContextUniform4iv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4iv
//go:noescape

func WebGL2RenderingContextUniform4ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4iv1
//go:noescape

func CallWebGL2RenderingContextUniform4iv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4iv1
//go:noescape

func WebGL2RenderingContextUniform4iv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4iv2
//go:noescape

func CallWebGL2RenderingContextUniform4iv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4iv2
//go:noescape

func WebGL2RenderingContextUniform4iv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2fv
//go:noescape

func WebGL2RenderingContextUniformMatrix2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix2fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix2fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3fv
//go:noescape

func WebGL2RenderingContextUniformMatrix3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix3fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix3fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4fv
//go:noescape

func WebGL2RenderingContextUniformMatrix4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix4fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix4fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ReadPixels
//go:noescape

func CallWebGL2RenderingContextReadPixels(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	dstData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ReadPixels
//go:noescape

func WebGL2RenderingContextReadPixelsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ReadPixels1
//go:noescape

func CallWebGL2RenderingContextReadPixels1(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ReadPixels1
//go:noescape

func WebGL2RenderingContextReadPixels1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ReadPixels2
//go:noescape

func CallWebGL2RenderingContextReadPixels2(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
	format uint32,
	typ uint32,
	dstData js.Ref,
	dstOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ReadPixels2
//go:noescape

func WebGL2RenderingContextReadPixels2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CopyBufferSubData
//go:noescape

func CallWebGL2RenderingContextCopyBufferSubData(
	this js.Ref,
	ptr unsafe.Pointer,

	readTarget uint32,
	writeTarget uint32,
	readOffset float64,
	writeOffset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CopyBufferSubData
//go:noescape

func WebGL2RenderingContextCopyBufferSubDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetBufferSubData
//go:noescape

func CallWebGL2RenderingContextGetBufferSubData(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcByteOffset float64,
	dstBuffer js.Ref,
	dstOffset uint32,
	length uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetBufferSubData
//go:noescape

func WebGL2RenderingContextGetBufferSubDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetBufferSubData1
//go:noescape

func CallWebGL2RenderingContextGetBufferSubData1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcByteOffset float64,
	dstBuffer js.Ref,
	dstOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetBufferSubData1
//go:noescape

func WebGL2RenderingContextGetBufferSubData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetBufferSubData2
//go:noescape

func CallWebGL2RenderingContextGetBufferSubData2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	srcByteOffset float64,
	dstBuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetBufferSubData2
//go:noescape

func WebGL2RenderingContextGetBufferSubData2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlitFramebuffer
//go:noescape

func CallWebGL2RenderingContextBlitFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	srcX0 int32,
	srcY0 int32,
	srcX1 int32,
	srcY1 int32,
	dstX0 int32,
	dstY0 int32,
	dstX1 int32,
	dstY1 int32,
	mask uint32,
	filter uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlitFramebuffer
//go:noescape

func WebGL2RenderingContextBlitFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_FramebufferTextureLayer
//go:noescape

func CallWebGL2RenderingContextFramebufferTextureLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	texture js.Ref,
	level int32,
	layer int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_FramebufferTextureLayer
//go:noescape

func WebGL2RenderingContextFramebufferTextureLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_InvalidateFramebuffer
//go:noescape

func CallWebGL2RenderingContextInvalidateFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachments js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_InvalidateFramebuffer
//go:noescape

func WebGL2RenderingContextInvalidateFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_InvalidateSubFramebuffer
//go:noescape

func CallWebGL2RenderingContextInvalidateSubFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachments js.Ref,
	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_InvalidateSubFramebuffer
//go:noescape

func WebGL2RenderingContextInvalidateSubFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ReadBuffer
//go:noescape

func CallWebGL2RenderingContextReadBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	src uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ReadBuffer
//go:noescape

func WebGL2RenderingContextReadBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetInternalformatParameter
//go:noescape

func CallWebGL2RenderingContextGetInternalformatParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	internalformat uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetInternalformatParameter
//go:noescape

func WebGL2RenderingContextGetInternalformatParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_RenderbufferStorageMultisample
//go:noescape

func CallWebGL2RenderingContextRenderbufferStorageMultisample(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	samples int32,
	internalformat uint32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_RenderbufferStorageMultisample
//go:noescape

func WebGL2RenderingContextRenderbufferStorageMultisampleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexStorage2D
//go:noescape

func CallWebGL2RenderingContextTexStorage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	levels int32,
	internalformat uint32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexStorage2D
//go:noescape

func WebGL2RenderingContextTexStorage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexStorage3D
//go:noescape

func CallWebGL2RenderingContextTexStorage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	levels int32,
	internalformat uint32,
	width int32,
	height int32,
	depth int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexStorage3D
//go:noescape

func WebGL2RenderingContextTexStorage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage3D
//go:noescape

func CallWebGL2RenderingContextTexImage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	depth int32,
	border int32,
	format uint32,
	typ uint32,
	pboOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage3D
//go:noescape

func WebGL2RenderingContextTexImage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage3D1
//go:noescape

func CallWebGL2RenderingContextTexImage3D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	depth int32,
	border int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage3D1
//go:noescape

func WebGL2RenderingContextTexImage3D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage3D2
//go:noescape

func CallWebGL2RenderingContextTexImage3D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	depth int32,
	border int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage3D2
//go:noescape

func WebGL2RenderingContextTexImage3D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexImage3D3
//go:noescape

func CallWebGL2RenderingContextTexImage3D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat int32,
	width int32,
	height int32,
	depth int32,
	border int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexImage3D3
//go:noescape

func WebGL2RenderingContextTexImage3D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage3D
//go:noescape

func CallWebGL2RenderingContextTexSubImage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	typ uint32,
	pboOffset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage3D
//go:noescape

func WebGL2RenderingContextTexSubImage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage3D1
//go:noescape

func CallWebGL2RenderingContextTexSubImage3D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	typ uint32,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage3D1
//go:noescape

func WebGL2RenderingContextTexSubImage3D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage3D2
//go:noescape

func CallWebGL2RenderingContextTexSubImage3D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage3D2
//go:noescape

func WebGL2RenderingContextTexSubImage3D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexSubImage3D3
//go:noescape

func CallWebGL2RenderingContextTexSubImage3D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	typ uint32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexSubImage3D3
//go:noescape

func WebGL2RenderingContextTexSubImage3D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CopyTexSubImage3D
//go:noescape

func CallWebGL2RenderingContextCopyTexSubImage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CopyTexSubImage3D
//go:noescape

func WebGL2RenderingContextCopyTexSubImage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage3D
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	depth int32,
	border int32,
	imageSize int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage3D
//go:noescape

func WebGL2RenderingContextCompressedTexImage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage3D1
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage3D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	depth int32,
	border int32,
	srcData js.Ref,
	srcOffset uint32,
	srcLengthOverride uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage3D1
//go:noescape

func WebGL2RenderingContextCompressedTexImage3D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage3D2
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage3D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	depth int32,
	border int32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage3D2
//go:noescape

func WebGL2RenderingContextCompressedTexImage3D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexImage3D3
//go:noescape

func CallWebGL2RenderingContextCompressedTexImage3D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	width int32,
	height int32,
	depth int32,
	border int32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexImage3D3
//go:noescape

func WebGL2RenderingContextCompressedTexImage3D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage3D
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage3D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	imageSize int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage3D
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage3DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage3D1
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage3D1(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	srcData js.Ref,
	srcOffset uint32,
	srcLengthOverride uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage3D1
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage3D1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage3D2
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage3D2(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	srcData js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage3D2
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage3D2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompressedTexSubImage3D3
//go:noescape

func CallWebGL2RenderingContextCompressedTexSubImage3D3(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	zoffset int32,
	width int32,
	height int32,
	depth int32,
	format uint32,
	srcData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompressedTexSubImage3D3
//go:noescape

func WebGL2RenderingContextCompressedTexSubImage3D3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetFragDataLocation
//go:noescape

func CallWebGL2RenderingContextGetFragDataLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	name js.Ref,
) int32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetFragDataLocation
//go:noescape

func WebGL2RenderingContextGetFragDataLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1ui
//go:noescape

func CallWebGL2RenderingContextUniform1ui(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v0 uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1ui
//go:noescape

func WebGL2RenderingContextUniform1uiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2ui
//go:noescape

func CallWebGL2RenderingContextUniform2ui(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v0 uint32,
	v1 uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2ui
//go:noescape

func WebGL2RenderingContextUniform2uiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3ui
//go:noescape

func CallWebGL2RenderingContextUniform3ui(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v0 uint32,
	v1 uint32,
	v2 uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3ui
//go:noescape

func WebGL2RenderingContextUniform3uiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4ui
//go:noescape

func CallWebGL2RenderingContextUniform4ui(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	v0 uint32,
	v1 uint32,
	v2 uint32,
	v3 uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4ui
//go:noescape

func WebGL2RenderingContextUniform4uiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1uiv
//go:noescape

func CallWebGL2RenderingContextUniform1uiv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1uiv
//go:noescape

func WebGL2RenderingContextUniform1uivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1uiv1
//go:noescape

func CallWebGL2RenderingContextUniform1uiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1uiv1
//go:noescape

func WebGL2RenderingContextUniform1uiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1uiv2
//go:noescape

func CallWebGL2RenderingContextUniform1uiv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1uiv2
//go:noescape

func WebGL2RenderingContextUniform1uiv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2uiv
//go:noescape

func CallWebGL2RenderingContextUniform2uiv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2uiv
//go:noescape

func WebGL2RenderingContextUniform2uivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2uiv1
//go:noescape

func CallWebGL2RenderingContextUniform2uiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2uiv1
//go:noescape

func WebGL2RenderingContextUniform2uiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2uiv2
//go:noescape

func CallWebGL2RenderingContextUniform2uiv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2uiv2
//go:noescape

func WebGL2RenderingContextUniform2uiv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3uiv
//go:noescape

func CallWebGL2RenderingContextUniform3uiv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3uiv
//go:noescape

func WebGL2RenderingContextUniform3uivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3uiv1
//go:noescape

func CallWebGL2RenderingContextUniform3uiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3uiv1
//go:noescape

func WebGL2RenderingContextUniform3uiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3uiv2
//go:noescape

func CallWebGL2RenderingContextUniform3uiv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3uiv2
//go:noescape

func WebGL2RenderingContextUniform3uiv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4uiv
//go:noescape

func CallWebGL2RenderingContextUniform4uiv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4uiv
//go:noescape

func WebGL2RenderingContextUniform4uivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4uiv1
//go:noescape

func CallWebGL2RenderingContextUniform4uiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4uiv1
//go:noescape

func WebGL2RenderingContextUniform4uiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4uiv2
//go:noescape

func CallWebGL2RenderingContextUniform4uiv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4uiv2
//go:noescape

func WebGL2RenderingContextUniform4uiv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x2fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x2fv
//go:noescape

func WebGL2RenderingContextUniformMatrix3x2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x2fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x2fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x2fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix3x2fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x2fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x2fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x2fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix3x2fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x2fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x2fv
//go:noescape

func WebGL2RenderingContextUniformMatrix4x2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x2fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x2fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x2fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix4x2fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x2fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x2fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x2fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix4x2fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x3fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x3fv
//go:noescape

func WebGL2RenderingContextUniformMatrix2x3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x3fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x3fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x3fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix2x3fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x3fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x3fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x3fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix2x3fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x3fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x3fv
//go:noescape

func WebGL2RenderingContextUniformMatrix4x3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x3fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x3fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x3fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix4x3fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix4x3fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix4x3fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix4x3fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix4x3fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x4fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x4fv
//go:noescape

func WebGL2RenderingContextUniformMatrix2x4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x4fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x4fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x4fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix2x4fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix2x4fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix2x4fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix2x4fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix2x4fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x4fv
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
	srcLength uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x4fv
//go:noescape

func WebGL2RenderingContextUniformMatrix3x4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x4fv1
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x4fv1(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x4fv1
//go:noescape

func WebGL2RenderingContextUniformMatrix3x4fv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformMatrix3x4fv2
//go:noescape

func CallWebGL2RenderingContextUniformMatrix3x4fv2(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	transpose js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformMatrix3x4fv2
//go:noescape

func WebGL2RenderingContextUniformMatrix3x4fv2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribI4i
//go:noescape

func CallWebGL2RenderingContextVertexAttribI4i(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x int32,
	y int32,
	z int32,
	w int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribI4i
//go:noescape

func WebGL2RenderingContextVertexAttribI4iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribI4iv
//go:noescape

func CallWebGL2RenderingContextVertexAttribI4iv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribI4iv
//go:noescape

func WebGL2RenderingContextVertexAttribI4ivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribI4ui
//go:noescape

func CallWebGL2RenderingContextVertexAttribI4ui(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x uint32,
	y uint32,
	z uint32,
	w uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribI4ui
//go:noescape

func WebGL2RenderingContextVertexAttribI4uiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribI4uiv
//go:noescape

func CallWebGL2RenderingContextVertexAttribI4uiv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribI4uiv
//go:noescape

func WebGL2RenderingContextVertexAttribI4uivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribIPointer
//go:noescape

func CallWebGL2RenderingContextVertexAttribIPointer(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	size int32,
	typ uint32,
	stride int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribIPointer
//go:noescape

func WebGL2RenderingContextVertexAttribIPointerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribDivisor
//go:noescape

func CallWebGL2RenderingContextVertexAttribDivisor(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	divisor uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribDivisor
//go:noescape

func WebGL2RenderingContextVertexAttribDivisorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawArraysInstanced
//go:noescape

func CallWebGL2RenderingContextDrawArraysInstanced(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	first int32,
	count int32,
	instanceCount int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawArraysInstanced
//go:noescape

func WebGL2RenderingContextDrawArraysInstancedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawElementsInstanced
//go:noescape

func CallWebGL2RenderingContextDrawElementsInstanced(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	count int32,
	typ uint32,
	offset float64,
	instanceCount int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawElementsInstanced
//go:noescape

func WebGL2RenderingContextDrawElementsInstancedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawRangeElements
//go:noescape

func CallWebGL2RenderingContextDrawRangeElements(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	start uint32,
	end uint32,
	count int32,
	typ uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawRangeElements
//go:noescape

func WebGL2RenderingContextDrawRangeElementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawBuffers
//go:noescape

func CallWebGL2RenderingContextDrawBuffers(
	this js.Ref,
	ptr unsafe.Pointer,

	buffers js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawBuffers
//go:noescape

func WebGL2RenderingContextDrawBuffersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferfv
//go:noescape

func CallWebGL2RenderingContextClearBufferfv(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferfv
//go:noescape

func WebGL2RenderingContextClearBufferfvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferfv1
//go:noescape

func CallWebGL2RenderingContextClearBufferfv1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferfv1
//go:noescape

func WebGL2RenderingContextClearBufferfv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferiv
//go:noescape

func CallWebGL2RenderingContextClearBufferiv(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferiv
//go:noescape

func WebGL2RenderingContextClearBufferivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferiv1
//go:noescape

func CallWebGL2RenderingContextClearBufferiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferiv1
//go:noescape

func WebGL2RenderingContextClearBufferiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferuiv
//go:noescape

func CallWebGL2RenderingContextClearBufferuiv(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
	srcOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferuiv
//go:noescape

func WebGL2RenderingContextClearBufferuivFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferuiv1
//go:noescape

func CallWebGL2RenderingContextClearBufferuiv1(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferuiv1
//go:noescape

func WebGL2RenderingContextClearBufferuiv1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearBufferfi
//go:noescape

func CallWebGL2RenderingContextClearBufferfi(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer uint32,
	drawbuffer int32,
	depth float32,
	stencil int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearBufferfi
//go:noescape

func WebGL2RenderingContextClearBufferfiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateQuery
//go:noescape

func CallWebGL2RenderingContextCreateQuery(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateQuery
//go:noescape

func WebGL2RenderingContextCreateQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteQuery
//go:noescape

func CallWebGL2RenderingContextDeleteQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	query js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteQuery
//go:noescape

func WebGL2RenderingContextDeleteQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsQuery
//go:noescape

func CallWebGL2RenderingContextIsQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	query js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsQuery
//go:noescape

func WebGL2RenderingContextIsQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BeginQuery
//go:noescape

func CallWebGL2RenderingContextBeginQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	query js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BeginQuery
//go:noescape

func WebGL2RenderingContextBeginQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_EndQuery
//go:noescape

func CallWebGL2RenderingContextEndQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_EndQuery
//go:noescape

func WebGL2RenderingContextEndQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetQuery
//go:noescape

func CallWebGL2RenderingContextGetQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetQuery
//go:noescape

func WebGL2RenderingContextGetQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetQueryParameter
//go:noescape

func CallWebGL2RenderingContextGetQueryParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	query js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetQueryParameter
//go:noescape

func WebGL2RenderingContextGetQueryParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateSampler
//go:noescape

func CallWebGL2RenderingContextCreateSampler(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateSampler
//go:noescape

func WebGL2RenderingContextCreateSamplerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteSampler
//go:noescape

func CallWebGL2RenderingContextDeleteSampler(
	this js.Ref,
	ptr unsafe.Pointer,

	sampler js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteSampler
//go:noescape

func WebGL2RenderingContextDeleteSamplerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsSampler
//go:noescape

func CallWebGL2RenderingContextIsSampler(
	this js.Ref,
	ptr unsafe.Pointer,

	sampler js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsSampler
//go:noescape

func WebGL2RenderingContextIsSamplerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindSampler
//go:noescape

func CallWebGL2RenderingContextBindSampler(
	this js.Ref,
	ptr unsafe.Pointer,

	unit uint32,
	sampler js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindSampler
//go:noescape

func WebGL2RenderingContextBindSamplerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_SamplerParameteri
//go:noescape

func CallWebGL2RenderingContextSamplerParameteri(
	this js.Ref,
	ptr unsafe.Pointer,

	sampler js.Ref,
	pname uint32,
	param int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_SamplerParameteri
//go:noescape

func WebGL2RenderingContextSamplerParameteriFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_SamplerParameterf
//go:noescape

func CallWebGL2RenderingContextSamplerParameterf(
	this js.Ref,
	ptr unsafe.Pointer,

	sampler js.Ref,
	pname uint32,
	param float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_SamplerParameterf
//go:noescape

func WebGL2RenderingContextSamplerParameterfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetSamplerParameter
//go:noescape

func CallWebGL2RenderingContextGetSamplerParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	sampler js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetSamplerParameter
//go:noescape

func WebGL2RenderingContextGetSamplerParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_FenceSync
//go:noescape

func CallWebGL2RenderingContextFenceSync(
	this js.Ref,
	ptr unsafe.Pointer,

	condition uint32,
	flags uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_FenceSync
//go:noescape

func WebGL2RenderingContextFenceSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsSync
//go:noescape

func CallWebGL2RenderingContextIsSync(
	this js.Ref,
	ptr unsafe.Pointer,

	sync js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsSync
//go:noescape

func WebGL2RenderingContextIsSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteSync
//go:noescape

func CallWebGL2RenderingContextDeleteSync(
	this js.Ref,
	ptr unsafe.Pointer,

	sync js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteSync
//go:noescape

func WebGL2RenderingContextDeleteSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClientWaitSync
//go:noescape

func CallWebGL2RenderingContextClientWaitSync(
	this js.Ref,
	ptr unsafe.Pointer,

	sync js.Ref,
	flags uint32,
	timeout float64,
) uint32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClientWaitSync
//go:noescape

func WebGL2RenderingContextClientWaitSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_WaitSync
//go:noescape

func CallWebGL2RenderingContextWaitSync(
	this js.Ref,
	ptr unsafe.Pointer,

	sync js.Ref,
	flags uint32,
	timeout float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_WaitSync
//go:noescape

func WebGL2RenderingContextWaitSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetSyncParameter
//go:noescape

func CallWebGL2RenderingContextGetSyncParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	sync js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetSyncParameter
//go:noescape

func WebGL2RenderingContextGetSyncParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateTransformFeedback
//go:noescape

func CallWebGL2RenderingContextCreateTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateTransformFeedback
//go:noescape

func WebGL2RenderingContextCreateTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteTransformFeedback
//go:noescape

func CallWebGL2RenderingContextDeleteTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

	tf js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteTransformFeedback
//go:noescape

func WebGL2RenderingContextDeleteTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsTransformFeedback
//go:noescape

func CallWebGL2RenderingContextIsTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

	tf js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsTransformFeedback
//go:noescape

func WebGL2RenderingContextIsTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindTransformFeedback
//go:noescape

func CallWebGL2RenderingContextBindTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	tf js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindTransformFeedback
//go:noescape

func WebGL2RenderingContextBindTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BeginTransformFeedback
//go:noescape

func CallWebGL2RenderingContextBeginTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

	primitiveMode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BeginTransformFeedback
//go:noescape

func WebGL2RenderingContextBeginTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_EndTransformFeedback
//go:noescape

func CallWebGL2RenderingContextEndTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_EndTransformFeedback
//go:noescape

func WebGL2RenderingContextEndTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TransformFeedbackVaryings
//go:noescape

func CallWebGL2RenderingContextTransformFeedbackVaryings(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	varyings js.Ref,
	bufferMode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TransformFeedbackVaryings
//go:noescape

func WebGL2RenderingContextTransformFeedbackVaryingsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetTransformFeedbackVarying
//go:noescape

func CallWebGL2RenderingContextGetTransformFeedbackVarying(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetTransformFeedbackVarying
//go:noescape

func WebGL2RenderingContextGetTransformFeedbackVaryingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_PauseTransformFeedback
//go:noescape

func CallWebGL2RenderingContextPauseTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_PauseTransformFeedback
//go:noescape

func WebGL2RenderingContextPauseTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ResumeTransformFeedback
//go:noescape

func CallWebGL2RenderingContextResumeTransformFeedback(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ResumeTransformFeedback
//go:noescape

func WebGL2RenderingContextResumeTransformFeedbackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindBufferBase
//go:noescape

func CallWebGL2RenderingContextBindBufferBase(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	index uint32,
	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindBufferBase
//go:noescape

func WebGL2RenderingContextBindBufferBaseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindBufferRange
//go:noescape

func CallWebGL2RenderingContextBindBufferRange(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	index uint32,
	buffer js.Ref,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindBufferRange
//go:noescape

func WebGL2RenderingContextBindBufferRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetIndexedParameter
//go:noescape

func CallWebGL2RenderingContextGetIndexedParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetIndexedParameter
//go:noescape

func WebGL2RenderingContextGetIndexedParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetUniformIndices
//go:noescape

func CallWebGL2RenderingContextGetUniformIndices(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformNames js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetUniformIndices
//go:noescape

func WebGL2RenderingContextGetUniformIndicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetActiveUniforms
//go:noescape

func CallWebGL2RenderingContextGetActiveUniforms(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformIndices js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetActiveUniforms
//go:noescape

func WebGL2RenderingContextGetActiveUniformsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetUniformBlockIndex
//go:noescape

func CallWebGL2RenderingContextGetUniformBlockIndex(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformBlockName js.Ref,
) uint32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetUniformBlockIndex
//go:noescape

func WebGL2RenderingContextGetUniformBlockIndexFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetActiveUniformBlockParameter
//go:noescape

func CallWebGL2RenderingContextGetActiveUniformBlockParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformBlockIndex uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetActiveUniformBlockParameter
//go:noescape

func WebGL2RenderingContextGetActiveUniformBlockParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetActiveUniformBlockName
//go:noescape

func CallWebGL2RenderingContextGetActiveUniformBlockName(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformBlockIndex uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetActiveUniformBlockName
//go:noescape

func WebGL2RenderingContextGetActiveUniformBlockNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UniformBlockBinding
//go:noescape

func CallWebGL2RenderingContextUniformBlockBinding(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	uniformBlockIndex uint32,
	uniformBlockBinding uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UniformBlockBinding
//go:noescape

func WebGL2RenderingContextUniformBlockBindingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateVertexArray
//go:noescape

func CallWebGL2RenderingContextCreateVertexArray(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateVertexArray
//go:noescape

func WebGL2RenderingContextCreateVertexArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteVertexArray
//go:noescape

func CallWebGL2RenderingContextDeleteVertexArray(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexArray js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteVertexArray
//go:noescape

func WebGL2RenderingContextDeleteVertexArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsVertexArray
//go:noescape

func CallWebGL2RenderingContextIsVertexArray(
	this js.Ref,
	ptr unsafe.Pointer,

	vertexArray js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsVertexArray
//go:noescape

func WebGL2RenderingContextIsVertexArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindVertexArray
//go:noescape

func CallWebGL2RenderingContextBindVertexArray(
	this js.Ref,
	ptr unsafe.Pointer,

	array js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindVertexArray
//go:noescape

func WebGL2RenderingContextBindVertexArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetContextAttributes
//go:noescape

func CallWebGL2RenderingContextGetContextAttributes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetContextAttributes
//go:noescape

func WebGL2RenderingContextGetContextAttributesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsContextLost
//go:noescape

func CallWebGL2RenderingContextIsContextLost(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsContextLost
//go:noescape

func WebGL2RenderingContextIsContextLostFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetSupportedExtensions
//go:noescape

func CallWebGL2RenderingContextGetSupportedExtensions(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetSupportedExtensions
//go:noescape

func WebGL2RenderingContextGetSupportedExtensionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetExtension
//go:noescape

func CallWebGL2RenderingContextGetExtension(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetExtension
//go:noescape

func WebGL2RenderingContextGetExtensionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ActiveTexture
//go:noescape

func CallWebGL2RenderingContextActiveTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ActiveTexture
//go:noescape

func WebGL2RenderingContextActiveTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_AttachShader
//go:noescape

func CallWebGL2RenderingContextAttachShader(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_AttachShader
//go:noescape

func WebGL2RenderingContextAttachShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindAttribLocation
//go:noescape

func CallWebGL2RenderingContextBindAttribLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindAttribLocation
//go:noescape

func WebGL2RenderingContextBindAttribLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindBuffer
//go:noescape

func CallWebGL2RenderingContextBindBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindBuffer
//go:noescape

func WebGL2RenderingContextBindBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindFramebuffer
//go:noescape

func CallWebGL2RenderingContextBindFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindFramebuffer
//go:noescape

func WebGL2RenderingContextBindFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindRenderbuffer
//go:noescape

func CallWebGL2RenderingContextBindRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindRenderbuffer
//go:noescape

func WebGL2RenderingContextBindRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BindTexture
//go:noescape

func CallWebGL2RenderingContextBindTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BindTexture
//go:noescape

func WebGL2RenderingContextBindTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlendColor
//go:noescape

func CallWebGL2RenderingContextBlendColor(
	this js.Ref,
	ptr unsafe.Pointer,

	red float32,
	green float32,
	blue float32,
	alpha float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlendColor
//go:noescape

func WebGL2RenderingContextBlendColorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlendEquation
//go:noescape

func CallWebGL2RenderingContextBlendEquation(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlendEquation
//go:noescape

func WebGL2RenderingContextBlendEquationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlendEquationSeparate
//go:noescape

func CallWebGL2RenderingContextBlendEquationSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	modeRGB uint32,
	modeAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlendEquationSeparate
//go:noescape

func WebGL2RenderingContextBlendEquationSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlendFunc
//go:noescape

func CallWebGL2RenderingContextBlendFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	sfactor uint32,
	dfactor uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlendFunc
//go:noescape

func WebGL2RenderingContextBlendFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_BlendFuncSeparate
//go:noescape

func CallWebGL2RenderingContextBlendFuncSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	srcRGB uint32,
	dstRGB uint32,
	srcAlpha uint32,
	dstAlpha uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_BlendFuncSeparate
//go:noescape

func WebGL2RenderingContextBlendFuncSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CheckFramebufferStatus
//go:noescape

func CallWebGL2RenderingContextCheckFramebufferStatus(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
) uint32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CheckFramebufferStatus
//go:noescape

func WebGL2RenderingContextCheckFramebufferStatusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Clear
//go:noescape

func CallWebGL2RenderingContextClear(
	this js.Ref,
	ptr unsafe.Pointer,

	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Clear
//go:noescape

func WebGL2RenderingContextClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearColor
//go:noescape

func CallWebGL2RenderingContextClearColor(
	this js.Ref,
	ptr unsafe.Pointer,

	red float32,
	green float32,
	blue float32,
	alpha float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearColor
//go:noescape

func WebGL2RenderingContextClearColorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearDepth
//go:noescape

func CallWebGL2RenderingContextClearDepth(
	this js.Ref,
	ptr unsafe.Pointer,

	depth float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearDepth
//go:noescape

func WebGL2RenderingContextClearDepthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ClearStencil
//go:noescape

func CallWebGL2RenderingContextClearStencil(
	this js.Ref,
	ptr unsafe.Pointer,

	s int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ClearStencil
//go:noescape

func WebGL2RenderingContextClearStencilFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ColorMask
//go:noescape

func CallWebGL2RenderingContextColorMask(
	this js.Ref,
	ptr unsafe.Pointer,

	red js.Ref,
	green js.Ref,
	blue js.Ref,
	alpha js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ColorMask
//go:noescape

func WebGL2RenderingContextColorMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CompileShader
//go:noescape

func CallWebGL2RenderingContextCompileShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CompileShader
//go:noescape

func WebGL2RenderingContextCompileShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CopyTexImage2D
//go:noescape

func CallWebGL2RenderingContextCopyTexImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	internalformat uint32,
	x int32,
	y int32,
	width int32,
	height int32,
	border int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CopyTexImage2D
//go:noescape

func WebGL2RenderingContextCopyTexImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CopyTexSubImage2D
//go:noescape

func CallWebGL2RenderingContextCopyTexSubImage2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	level int32,
	xoffset int32,
	yoffset int32,
	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CopyTexSubImage2D
//go:noescape

func WebGL2RenderingContextCopyTexSubImage2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateBuffer
//go:noescape

func CallWebGL2RenderingContextCreateBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateBuffer
//go:noescape

func WebGL2RenderingContextCreateBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateFramebuffer
//go:noescape

func CallWebGL2RenderingContextCreateFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateFramebuffer
//go:noescape

func WebGL2RenderingContextCreateFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateProgram
//go:noescape

func CallWebGL2RenderingContextCreateProgram(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateProgram
//go:noescape

func WebGL2RenderingContextCreateProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateRenderbuffer
//go:noescape

func CallWebGL2RenderingContextCreateRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateRenderbuffer
//go:noescape

func WebGL2RenderingContextCreateRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateShader
//go:noescape

func CallWebGL2RenderingContextCreateShader(
	this js.Ref,
	ptr unsafe.Pointer,

	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateShader
//go:noescape

func WebGL2RenderingContextCreateShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CreateTexture
//go:noescape

func CallWebGL2RenderingContextCreateTexture(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CreateTexture
//go:noescape

func WebGL2RenderingContextCreateTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_CullFace
//go:noescape

func CallWebGL2RenderingContextCullFace(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_CullFace
//go:noescape

func WebGL2RenderingContextCullFaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteBuffer
//go:noescape

func CallWebGL2RenderingContextDeleteBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteBuffer
//go:noescape

func WebGL2RenderingContextDeleteBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteFramebuffer
//go:noescape

func CallWebGL2RenderingContextDeleteFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteFramebuffer
//go:noescape

func WebGL2RenderingContextDeleteFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteProgram
//go:noescape

func CallWebGL2RenderingContextDeleteProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteProgram
//go:noescape

func WebGL2RenderingContextDeleteProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteRenderbuffer
//go:noescape

func CallWebGL2RenderingContextDeleteRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteRenderbuffer
//go:noescape

func WebGL2RenderingContextDeleteRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteShader
//go:noescape

func CallWebGL2RenderingContextDeleteShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteShader
//go:noescape

func WebGL2RenderingContextDeleteShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DeleteTexture
//go:noescape

func CallWebGL2RenderingContextDeleteTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DeleteTexture
//go:noescape

func WebGL2RenderingContextDeleteTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DepthFunc
//go:noescape

func CallWebGL2RenderingContextDepthFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	fn uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DepthFunc
//go:noescape

func WebGL2RenderingContextDepthFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DepthMask
//go:noescape

func CallWebGL2RenderingContextDepthMask(
	this js.Ref,
	ptr unsafe.Pointer,

	flag js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DepthMask
//go:noescape

func WebGL2RenderingContextDepthMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DepthRange
//go:noescape

func CallWebGL2RenderingContextDepthRange(
	this js.Ref,
	ptr unsafe.Pointer,

	zNear float32,
	zFar float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DepthRange
//go:noescape

func WebGL2RenderingContextDepthRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DetachShader
//go:noescape

func CallWebGL2RenderingContextDetachShader(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DetachShader
//go:noescape

func WebGL2RenderingContextDetachShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Disable
//go:noescape

func CallWebGL2RenderingContextDisable(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Disable
//go:noescape

func WebGL2RenderingContextDisableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DisableVertexAttribArray
//go:noescape

func CallWebGL2RenderingContextDisableVertexAttribArray(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DisableVertexAttribArray
//go:noescape

func WebGL2RenderingContextDisableVertexAttribArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawArrays
//go:noescape

func CallWebGL2RenderingContextDrawArrays(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	first int32,
	count int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawArrays
//go:noescape

func WebGL2RenderingContextDrawArraysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_DrawElements
//go:noescape

func CallWebGL2RenderingContextDrawElements(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	count int32,
	typ uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_DrawElements
//go:noescape

func WebGL2RenderingContextDrawElementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Enable
//go:noescape

func CallWebGL2RenderingContextEnable(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Enable
//go:noescape

func WebGL2RenderingContextEnableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_EnableVertexAttribArray
//go:noescape

func CallWebGL2RenderingContextEnableVertexAttribArray(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_EnableVertexAttribArray
//go:noescape

func WebGL2RenderingContextEnableVertexAttribArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Finish
//go:noescape

func CallWebGL2RenderingContextFinish(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Finish
//go:noescape

func WebGL2RenderingContextFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Flush
//go:noescape

func CallWebGL2RenderingContextFlush(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Flush
//go:noescape

func WebGL2RenderingContextFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_FramebufferRenderbuffer
//go:noescape

func CallWebGL2RenderingContextFramebufferRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	renderbuffertarget uint32,
	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_FramebufferRenderbuffer
//go:noescape

func WebGL2RenderingContextFramebufferRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_FramebufferTexture2D
//go:noescape

func CallWebGL2RenderingContextFramebufferTexture2D(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	textarget uint32,
	texture js.Ref,
	level int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_FramebufferTexture2D
//go:noescape

func WebGL2RenderingContextFramebufferTexture2DFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_FrontFace
//go:noescape

func CallWebGL2RenderingContextFrontFace(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_FrontFace
//go:noescape

func WebGL2RenderingContextFrontFaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GenerateMipmap
//go:noescape

func CallWebGL2RenderingContextGenerateMipmap(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GenerateMipmap
//go:noescape

func WebGL2RenderingContextGenerateMipmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetActiveAttrib
//go:noescape

func CallWebGL2RenderingContextGetActiveAttrib(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetActiveAttrib
//go:noescape

func WebGL2RenderingContextGetActiveAttribFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetActiveUniform
//go:noescape

func CallWebGL2RenderingContextGetActiveUniform(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetActiveUniform
//go:noescape

func WebGL2RenderingContextGetActiveUniformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetAttachedShaders
//go:noescape

func CallWebGL2RenderingContextGetAttachedShaders(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetAttachedShaders
//go:noescape

func WebGL2RenderingContextGetAttachedShadersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetAttribLocation
//go:noescape

func CallWebGL2RenderingContextGetAttribLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	name js.Ref,
) int32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetAttribLocation
//go:noescape

func WebGL2RenderingContextGetAttribLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetBufferParameter
//go:noescape

func CallWebGL2RenderingContextGetBufferParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetBufferParameter
//go:noescape

func WebGL2RenderingContextGetBufferParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetParameter
//go:noescape

func CallWebGL2RenderingContextGetParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetParameter
//go:noescape

func WebGL2RenderingContextGetParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetError
//go:noescape

func CallWebGL2RenderingContextGetError(
	this js.Ref,
	ptr unsafe.Pointer,

) uint32

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetError
//go:noescape

func WebGL2RenderingContextGetErrorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetFramebufferAttachmentParameter
//go:noescape

func CallWebGL2RenderingContextGetFramebufferAttachmentParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	attachment uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetFramebufferAttachmentParameter
//go:noescape

func WebGL2RenderingContextGetFramebufferAttachmentParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetProgramParameter
//go:noescape

func CallWebGL2RenderingContextGetProgramParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetProgramParameter
//go:noescape

func WebGL2RenderingContextGetProgramParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetProgramInfoLog
//go:noescape

func CallWebGL2RenderingContextGetProgramInfoLog(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetProgramInfoLog
//go:noescape

func WebGL2RenderingContextGetProgramInfoLogFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetRenderbufferParameter
//go:noescape

func CallWebGL2RenderingContextGetRenderbufferParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetRenderbufferParameter
//go:noescape

func WebGL2RenderingContextGetRenderbufferParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetShaderParameter
//go:noescape

func CallWebGL2RenderingContextGetShaderParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetShaderParameter
//go:noescape

func WebGL2RenderingContextGetShaderParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetShaderPrecisionFormat
//go:noescape

func CallWebGL2RenderingContextGetShaderPrecisionFormat(
	this js.Ref,
	ptr unsafe.Pointer,

	shadertype uint32,
	precisiontype uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetShaderPrecisionFormat
//go:noescape

func WebGL2RenderingContextGetShaderPrecisionFormatFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetShaderInfoLog
//go:noescape

func CallWebGL2RenderingContextGetShaderInfoLog(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetShaderInfoLog
//go:noescape

func WebGL2RenderingContextGetShaderInfoLogFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetShaderSource
//go:noescape

func CallWebGL2RenderingContextGetShaderSource(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetShaderSource
//go:noescape

func WebGL2RenderingContextGetShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetTexParameter
//go:noescape

func CallWebGL2RenderingContextGetTexParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetTexParameter
//go:noescape

func WebGL2RenderingContextGetTexParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetUniform
//go:noescape

func CallWebGL2RenderingContextGetUniform(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	location js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetUniform
//go:noescape

func WebGL2RenderingContextGetUniformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetUniformLocation
//go:noescape

func CallWebGL2RenderingContextGetUniformLocation(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetUniformLocation
//go:noescape

func WebGL2RenderingContextGetUniformLocationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetVertexAttrib
//go:noescape

func CallWebGL2RenderingContextGetVertexAttrib(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	pname uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetVertexAttrib
//go:noescape

func WebGL2RenderingContextGetVertexAttribFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_GetVertexAttribOffset
//go:noescape

func CallWebGL2RenderingContextGetVertexAttribOffset(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	pname uint32,
) int64

//go:wasmimport plat/js/web func_WebGL2RenderingContext_GetVertexAttribOffset
//go:noescape

func WebGL2RenderingContextGetVertexAttribOffsetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Hint
//go:noescape

func CallWebGL2RenderingContextHint(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Hint
//go:noescape

func WebGL2RenderingContextHintFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsBuffer
//go:noescape

func CallWebGL2RenderingContextIsBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	buffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsBuffer
//go:noescape

func WebGL2RenderingContextIsBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsEnabled
//go:noescape

func CallWebGL2RenderingContextIsEnabled(
	this js.Ref,
	ptr unsafe.Pointer,

	cap uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsEnabled
//go:noescape

func WebGL2RenderingContextIsEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsFramebuffer
//go:noescape

func CallWebGL2RenderingContextIsFramebuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	framebuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsFramebuffer
//go:noescape

func WebGL2RenderingContextIsFramebufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsProgram
//go:noescape

func CallWebGL2RenderingContextIsProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsProgram
//go:noescape

func WebGL2RenderingContextIsProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsRenderbuffer
//go:noescape

func CallWebGL2RenderingContextIsRenderbuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	renderbuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsRenderbuffer
//go:noescape

func WebGL2RenderingContextIsRenderbufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsShader
//go:noescape

func CallWebGL2RenderingContextIsShader(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsShader
//go:noescape

func WebGL2RenderingContextIsShaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_IsTexture
//go:noescape

func CallWebGL2RenderingContextIsTexture(
	this js.Ref,
	ptr unsafe.Pointer,

	texture js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_IsTexture
//go:noescape

func WebGL2RenderingContextIsTextureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_LineWidth
//go:noescape

func CallWebGL2RenderingContextLineWidth(
	this js.Ref,
	ptr unsafe.Pointer,

	width float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_LineWidth
//go:noescape

func WebGL2RenderingContextLineWidthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_LinkProgram
//go:noescape

func CallWebGL2RenderingContextLinkProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_LinkProgram
//go:noescape

func WebGL2RenderingContextLinkProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_PixelStorei
//go:noescape

func CallWebGL2RenderingContextPixelStorei(
	this js.Ref,
	ptr unsafe.Pointer,

	pname uint32,
	param int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_PixelStorei
//go:noescape

func WebGL2RenderingContextPixelStoreiFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_PolygonOffset
//go:noescape

func CallWebGL2RenderingContextPolygonOffset(
	this js.Ref,
	ptr unsafe.Pointer,

	factor float32,
	units float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_PolygonOffset
//go:noescape

func WebGL2RenderingContextPolygonOffsetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_RenderbufferStorage
//go:noescape

func CallWebGL2RenderingContextRenderbufferStorage(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	internalformat uint32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_RenderbufferStorage
//go:noescape

func WebGL2RenderingContextRenderbufferStorageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_SampleCoverage
//go:noescape

func CallWebGL2RenderingContextSampleCoverage(
	this js.Ref,
	ptr unsafe.Pointer,

	value float32,
	invert js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_SampleCoverage
//go:noescape

func WebGL2RenderingContextSampleCoverageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Scissor
//go:noescape

func CallWebGL2RenderingContextScissor(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Scissor
//go:noescape

func WebGL2RenderingContextScissorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ShaderSource
//go:noescape

func CallWebGL2RenderingContextShaderSource(
	this js.Ref,
	ptr unsafe.Pointer,

	shader js.Ref,
	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ShaderSource
//go:noescape

func WebGL2RenderingContextShaderSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilFunc
//go:noescape

func CallWebGL2RenderingContextStencilFunc(
	this js.Ref,
	ptr unsafe.Pointer,

	fn uint32,
	ref int32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilFunc
//go:noescape

func WebGL2RenderingContextStencilFuncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilFuncSeparate
//go:noescape

func CallWebGL2RenderingContextStencilFuncSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	fn uint32,
	ref int32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilFuncSeparate
//go:noescape

func WebGL2RenderingContextStencilFuncSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilMask
//go:noescape

func CallWebGL2RenderingContextStencilMask(
	this js.Ref,
	ptr unsafe.Pointer,

	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilMask
//go:noescape

func WebGL2RenderingContextStencilMaskFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilMaskSeparate
//go:noescape

func CallWebGL2RenderingContextStencilMaskSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	mask uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilMaskSeparate
//go:noescape

func WebGL2RenderingContextStencilMaskSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilOp
//go:noescape

func CallWebGL2RenderingContextStencilOp(
	this js.Ref,
	ptr unsafe.Pointer,

	fail uint32,
	zfail uint32,
	zpass uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilOp
//go:noescape

func WebGL2RenderingContextStencilOpFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_StencilOpSeparate
//go:noescape

func CallWebGL2RenderingContextStencilOpSeparate(
	this js.Ref,
	ptr unsafe.Pointer,

	face uint32,
	fail uint32,
	zfail uint32,
	zpass uint32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_StencilOpSeparate
//go:noescape

func WebGL2RenderingContextStencilOpSeparateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexParameterf
//go:noescape

func CallWebGL2RenderingContextTexParameterf(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
	param float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexParameterf
//go:noescape

func WebGL2RenderingContextTexParameterfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_TexParameteri
//go:noescape

func CallWebGL2RenderingContextTexParameteri(
	this js.Ref,
	ptr unsafe.Pointer,

	target uint32,
	pname uint32,
	param int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_TexParameteri
//go:noescape

func WebGL2RenderingContextTexParameteriFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1f
//go:noescape

func CallWebGL2RenderingContextUniform1f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1f
//go:noescape

func WebGL2RenderingContextUniform1fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2f
//go:noescape

func CallWebGL2RenderingContextUniform2f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2f
//go:noescape

func WebGL2RenderingContextUniform2fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3f
//go:noescape

func CallWebGL2RenderingContextUniform3f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3f
//go:noescape

func WebGL2RenderingContextUniform3fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4f
//go:noescape

func CallWebGL2RenderingContextUniform4f(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x float32,
	y float32,
	z float32,
	w float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4f
//go:noescape

func WebGL2RenderingContextUniform4fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform1i
//go:noescape

func CallWebGL2RenderingContextUniform1i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform1i
//go:noescape

func WebGL2RenderingContextUniform1iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform2i
//go:noescape

func CallWebGL2RenderingContextUniform2i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform2i
//go:noescape

func WebGL2RenderingContextUniform2iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform3i
//go:noescape

func CallWebGL2RenderingContextUniform3i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
	z int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform3i
//go:noescape

func WebGL2RenderingContextUniform3iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Uniform4i
//go:noescape

func CallWebGL2RenderingContextUniform4i(
	this js.Ref,
	ptr unsafe.Pointer,

	location js.Ref,
	x int32,
	y int32,
	z int32,
	w int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Uniform4i
//go:noescape

func WebGL2RenderingContextUniform4iFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_UseProgram
//go:noescape

func CallWebGL2RenderingContextUseProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_UseProgram
//go:noescape

func WebGL2RenderingContextUseProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_ValidateProgram
//go:noescape

func CallWebGL2RenderingContextValidateProgram(
	this js.Ref,
	ptr unsafe.Pointer,

	program js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_ValidateProgram
//go:noescape

func WebGL2RenderingContextValidateProgramFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib1f
//go:noescape

func CallWebGL2RenderingContextVertexAttrib1f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib1f
//go:noescape

func WebGL2RenderingContextVertexAttrib1fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib2f
//go:noescape

func CallWebGL2RenderingContextVertexAttrib2f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib2f
//go:noescape

func WebGL2RenderingContextVertexAttrib2fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib3f
//go:noescape

func CallWebGL2RenderingContextVertexAttrib3f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
	z float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib3f
//go:noescape

func WebGL2RenderingContextVertexAttrib3fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib4f
//go:noescape

func CallWebGL2RenderingContextVertexAttrib4f(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	x float32,
	y float32,
	z float32,
	w float32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib4f
//go:noescape

func WebGL2RenderingContextVertexAttrib4fFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib1fv
//go:noescape

func CallWebGL2RenderingContextVertexAttrib1fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib1fv
//go:noescape

func WebGL2RenderingContextVertexAttrib1fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib2fv
//go:noescape

func CallWebGL2RenderingContextVertexAttrib2fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib2fv
//go:noescape

func WebGL2RenderingContextVertexAttrib2fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib3fv
//go:noescape

func CallWebGL2RenderingContextVertexAttrib3fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib3fv
//go:noescape

func WebGL2RenderingContextVertexAttrib3fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttrib4fv
//go:noescape

func CallWebGL2RenderingContextVertexAttrib4fv(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	values js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttrib4fv
//go:noescape

func WebGL2RenderingContextVertexAttrib4fvFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_VertexAttribPointer
//go:noescape

func CallWebGL2RenderingContextVertexAttribPointer(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	size int32,
	typ uint32,
	normalized js.Ref,
	stride int32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_VertexAttribPointer
//go:noescape

func WebGL2RenderingContextVertexAttribPointerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_Viewport
//go:noescape

func CallWebGL2RenderingContextViewport(
	this js.Ref,
	ptr unsafe.Pointer,

	x int32,
	y int32,
	width int32,
	height int32,
) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_Viewport
//go:noescape

func WebGL2RenderingContextViewportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WebGL2RenderingContext_MakeXRCompatible
//go:noescape

func CallWebGL2RenderingContextMakeXRCompatible(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WebGL2RenderingContext_MakeXRCompatible
//go:noescape

func WebGL2RenderingContextMakeXRCompatibleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUBufferMapState
//go:noescape

func ConstOfGPUBufferMapState(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUBuffer_Size
//go:noescape

func GetGPUBufferSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_GPUBuffer_Usage
//go:noescape

func GetGPUBufferUsage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUBuffer_MapState
//go:noescape

func GetGPUBufferMapState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUBuffer_Label
//go:noescape

func GetGPUBufferLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUBuffer_Label
//go:noescape

func SetGPUBufferLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_MapAsync
//go:noescape

func CallGPUBufferMapAsync(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_MapAsync
//go:noescape

func GPUBufferMapAsyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_MapAsync1
//go:noescape

func CallGPUBufferMapAsync1(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_MapAsync1
//go:noescape

func GPUBufferMapAsync1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_MapAsync2
//go:noescape

func CallGPUBufferMapAsync2(
	this js.Ref,
	ptr unsafe.Pointer,

	mode uint32,
) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_MapAsync2
//go:noescape

func GPUBufferMapAsync2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_GetMappedRange
//go:noescape

func CallGPUBufferGetMappedRange(
	this js.Ref,
	ptr unsafe.Pointer,

	offset float64,
	size float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_GetMappedRange
//go:noescape

func GPUBufferGetMappedRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_GetMappedRange1
//go:noescape

func CallGPUBufferGetMappedRange1(
	this js.Ref,
	ptr unsafe.Pointer,

	offset float64,
) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_GetMappedRange1
//go:noescape

func GPUBufferGetMappedRange1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_GetMappedRange2
//go:noescape

func CallGPUBufferGetMappedRange2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_GetMappedRange2
//go:noescape

func GPUBufferGetMappedRange2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_Unmap
//go:noescape

func CallGPUBufferUnmap(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_Unmap
//go:noescape

func GPUBufferUnmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUBuffer_Destroy
//go:noescape

func CallGPUBufferDestroy(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUBuffer_Destroy
//go:noescape

func GPUBufferDestroyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUBufferDescriptor
//go:noescape

func GPUBufferDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUBufferDescriptor
//go:noescape

func GPUBufferDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUTextureView_Label
//go:noescape

func GetGPUTextureViewLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUTextureView_Label
//go:noescape

func SetGPUTextureViewLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web constof_GPUTextureFormat
//go:noescape

func ConstOfGPUTextureFormat(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUTextureViewDimension
//go:noescape

func ConstOfGPUTextureViewDimension(str js.Ref) uint32

//go:wasmimport plat/js/web constof_GPUTextureAspect
//go:noescape

func ConstOfGPUTextureAspect(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUTextureViewDescriptor
//go:noescape

func GPUTextureViewDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUTextureViewDescriptor
//go:noescape

func GPUTextureViewDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_GPUTextureDimension
//go:noescape

func ConstOfGPUTextureDimension(str js.Ref) uint32

//go:wasmimport plat/js/web get_GPUTexture_Width
//go:noescape

func GetGPUTextureWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_Height
//go:noescape

func GetGPUTextureHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_DepthOrArrayLayers
//go:noescape

func GetGPUTextureDepthOrArrayLayers(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_MipLevelCount
//go:noescape

func GetGPUTextureMipLevelCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_SampleCount
//go:noescape

func GetGPUTextureSampleCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_Dimension
//go:noescape

func GetGPUTextureDimension(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_Format
//go:noescape

func GetGPUTextureFormat(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_Usage
//go:noescape

func GetGPUTextureUsage(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_GPUTexture_Label
//go:noescape

func GetGPUTextureLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_GPUTexture_Label
//go:noescape

func SetGPUTextureLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_GPUTexture_CreateView
//go:noescape

func CallGPUTextureCreateView(
	this js.Ref,
	ptr unsafe.Pointer,

	descriptor unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_GPUTexture_CreateView
//go:noescape

func GPUTextureCreateViewFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUTexture_CreateView1
//go:noescape

func CallGPUTextureCreateView1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUTexture_CreateView1
//go:noescape

func GPUTextureCreateView1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUTexture_Destroy
//go:noescape

func CallGPUTextureDestroy(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GPUTexture_Destroy
//go:noescape

func GPUTextureDestroyFunc(this js.Ref) js.Ref
