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

//go:wasmimport plat/js/web constof_OffscreenRenderingContextId
//go:noescape
func ConstOfOffscreenRenderingContextId(str js.Ref) uint32

//go:wasmimport plat/js/web store_ImageEncodeOptions
//go:noescape
func ImageEncodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageEncodeOptions
//go:noescape
func ImageEncodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_OffscreenCanvas_OffscreenCanvas
//go:noescape
func NewOffscreenCanvasByOffscreenCanvas(
	width float64,
	height float64) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvas_Width
//go:noescape
func GetOffscreenCanvasWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvas_Width
//go:noescape
func SetOffscreenCanvasWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvas_Height
//go:noescape
func GetOffscreenCanvasHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvas_Height
//go:noescape
func SetOffscreenCanvasHeight(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_OffscreenCanvas_GetContext
//go:noescape
func HasOffscreenCanvasGetContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvas_GetContext
//go:noescape
func OffscreenCanvasGetContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OffscreenCanvas_GetContext
//go:noescape
func CallOffscreenCanvasGetContext(
	this js.Ref, retPtr unsafe.Pointer,
	contextId uint32,
	options js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvas_GetContext
//go:noescape
func TryOffscreenCanvasGetContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextId uint32,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvas_GetContext1
//go:noescape
func HasOffscreenCanvasGetContext1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvas_GetContext1
//go:noescape
func OffscreenCanvasGetContext1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OffscreenCanvas_GetContext1
//go:noescape
func CallOffscreenCanvasGetContext1(
	this js.Ref, retPtr unsafe.Pointer,
	contextId uint32)

//go:wasmimport plat/js/web try_OffscreenCanvas_GetContext1
//go:noescape
func TryOffscreenCanvasGetContext1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextId uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvas_TransferToImageBitmap
//go:noescape
func HasOffscreenCanvasTransferToImageBitmap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvas_TransferToImageBitmap
//go:noescape
func OffscreenCanvasTransferToImageBitmapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OffscreenCanvas_TransferToImageBitmap
//go:noescape
func CallOffscreenCanvasTransferToImageBitmap(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvas_TransferToImageBitmap
//go:noescape
func TryOffscreenCanvasTransferToImageBitmap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvas_ConvertToBlob
//go:noescape
func HasOffscreenCanvasConvertToBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvas_ConvertToBlob
//go:noescape
func OffscreenCanvasConvertToBlobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OffscreenCanvas_ConvertToBlob
//go:noescape
func CallOffscreenCanvasConvertToBlob(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvas_ConvertToBlob
//go:noescape
func TryOffscreenCanvasConvertToBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvas_ConvertToBlob1
//go:noescape
func HasOffscreenCanvasConvertToBlob1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvas_ConvertToBlob1
//go:noescape
func OffscreenCanvasConvertToBlob1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_OffscreenCanvas_ConvertToBlob1
//go:noescape
func CallOffscreenCanvasConvertToBlob1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvas_ConvertToBlob1
//go:noescape
func TryOffscreenCanvasConvertToBlob1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_Canvas
//go:noescape
func GetCanvasRenderingContext2DCanvas(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_GlobalAlpha
//go:noescape
func GetCanvasRenderingContext2DGlobalAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_GlobalAlpha
//go:noescape
func SetCanvasRenderingContext2DGlobalAlpha(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_GlobalCompositeOperation
//go:noescape
func GetCanvasRenderingContext2DGlobalCompositeOperation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_GlobalCompositeOperation
//go:noescape
func SetCanvasRenderingContext2DGlobalCompositeOperation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ImageSmoothingEnabled
//go:noescape
func GetCanvasRenderingContext2DImageSmoothingEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ImageSmoothingEnabled
//go:noescape
func SetCanvasRenderingContext2DImageSmoothingEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ImageSmoothingQuality
//go:noescape
func GetCanvasRenderingContext2DImageSmoothingQuality(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ImageSmoothingQuality
//go:noescape
func SetCanvasRenderingContext2DImageSmoothingQuality(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_StrokeStyle
//go:noescape
func GetCanvasRenderingContext2DStrokeStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_StrokeStyle
//go:noescape
func SetCanvasRenderingContext2DStrokeStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_FillStyle
//go:noescape
func GetCanvasRenderingContext2DFillStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_FillStyle
//go:noescape
func SetCanvasRenderingContext2DFillStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ShadowOffsetX
//go:noescape
func GetCanvasRenderingContext2DShadowOffsetX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ShadowOffsetX
//go:noescape
func SetCanvasRenderingContext2DShadowOffsetX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ShadowOffsetY
//go:noescape
func GetCanvasRenderingContext2DShadowOffsetY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ShadowOffsetY
//go:noescape
func SetCanvasRenderingContext2DShadowOffsetY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ShadowBlur
//go:noescape
func GetCanvasRenderingContext2DShadowBlur(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ShadowBlur
//go:noescape
func SetCanvasRenderingContext2DShadowBlur(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_ShadowColor
//go:noescape
func GetCanvasRenderingContext2DShadowColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_ShadowColor
//go:noescape
func SetCanvasRenderingContext2DShadowColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_Filter
//go:noescape
func GetCanvasRenderingContext2DFilter(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_Filter
//go:noescape
func SetCanvasRenderingContext2DFilter(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_LineWidth
//go:noescape
func GetCanvasRenderingContext2DLineWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_LineWidth
//go:noescape
func SetCanvasRenderingContext2DLineWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_LineCap
//go:noescape
func GetCanvasRenderingContext2DLineCap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_LineCap
//go:noescape
func SetCanvasRenderingContext2DLineCap(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_LineJoin
//go:noescape
func GetCanvasRenderingContext2DLineJoin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_LineJoin
//go:noescape
func SetCanvasRenderingContext2DLineJoin(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_MiterLimit
//go:noescape
func GetCanvasRenderingContext2DMiterLimit(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_MiterLimit
//go:noescape
func SetCanvasRenderingContext2DMiterLimit(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_LineDashOffset
//go:noescape
func GetCanvasRenderingContext2DLineDashOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_LineDashOffset
//go:noescape
func SetCanvasRenderingContext2DLineDashOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_Font
//go:noescape
func GetCanvasRenderingContext2DFont(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_Font
//go:noescape
func SetCanvasRenderingContext2DFont(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_TextAlign
//go:noescape
func GetCanvasRenderingContext2DTextAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_TextAlign
//go:noescape
func SetCanvasRenderingContext2DTextAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_TextBaseline
//go:noescape
func GetCanvasRenderingContext2DTextBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_TextBaseline
//go:noescape
func SetCanvasRenderingContext2DTextBaseline(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_Direction
//go:noescape
func GetCanvasRenderingContext2DDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_Direction
//go:noescape
func SetCanvasRenderingContext2DDirection(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_LetterSpacing
//go:noescape
func GetCanvasRenderingContext2DLetterSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_LetterSpacing
//go:noescape
func SetCanvasRenderingContext2DLetterSpacing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_FontKerning
//go:noescape
func GetCanvasRenderingContext2DFontKerning(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_FontKerning
//go:noescape
func SetCanvasRenderingContext2DFontKerning(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_FontStretch
//go:noescape
func GetCanvasRenderingContext2DFontStretch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_FontStretch
//go:noescape
func SetCanvasRenderingContext2DFontStretch(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_FontVariantCaps
//go:noescape
func GetCanvasRenderingContext2DFontVariantCaps(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_FontVariantCaps
//go:noescape
func SetCanvasRenderingContext2DFontVariantCaps(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_TextRendering
//go:noescape
func GetCanvasRenderingContext2DTextRendering(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_TextRendering
//go:noescape
func SetCanvasRenderingContext2DTextRendering(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_CanvasRenderingContext2D_WordSpacing
//go:noescape
func GetCanvasRenderingContext2DWordSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CanvasRenderingContext2D_WordSpacing
//go:noescape
func SetCanvasRenderingContext2DWordSpacing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_GetContextAttributes
//go:noescape
func HasCanvasRenderingContext2DGetContextAttributes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_GetContextAttributes
//go:noescape
func CanvasRenderingContext2DGetContextAttributesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_GetContextAttributes
//go:noescape
func CallCanvasRenderingContext2DGetContextAttributes(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_GetContextAttributes
//go:noescape
func TryCanvasRenderingContext2DGetContextAttributes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Save
//go:noescape
func HasCanvasRenderingContext2DSave(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Save
//go:noescape
func CanvasRenderingContext2DSaveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Save
//go:noescape
func CallCanvasRenderingContext2DSave(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Save
//go:noescape
func TryCanvasRenderingContext2DSave(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Restore
//go:noescape
func HasCanvasRenderingContext2DRestore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Restore
//go:noescape
func CanvasRenderingContext2DRestoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Restore
//go:noescape
func CallCanvasRenderingContext2DRestore(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Restore
//go:noescape
func TryCanvasRenderingContext2DRestore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Reset
//go:noescape
func HasCanvasRenderingContext2DReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Reset
//go:noescape
func CanvasRenderingContext2DResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Reset
//go:noescape
func CallCanvasRenderingContext2DReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Reset
//go:noescape
func TryCanvasRenderingContext2DReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsContextLost
//go:noescape
func HasCanvasRenderingContext2DIsContextLost(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsContextLost
//go:noescape
func CanvasRenderingContext2DIsContextLostFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsContextLost
//go:noescape
func CallCanvasRenderingContext2DIsContextLost(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsContextLost
//go:noescape
func TryCanvasRenderingContext2DIsContextLost(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Scale
//go:noescape
func HasCanvasRenderingContext2DScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Scale
//go:noescape
func CanvasRenderingContext2DScaleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Scale
//go:noescape
func CallCanvasRenderingContext2DScale(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Scale
//go:noescape
func TryCanvasRenderingContext2DScale(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Rotate
//go:noescape
func HasCanvasRenderingContext2DRotate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Rotate
//go:noescape
func CanvasRenderingContext2DRotateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Rotate
//go:noescape
func CallCanvasRenderingContext2DRotate(
	this js.Ref, retPtr unsafe.Pointer,
	angle float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Rotate
//go:noescape
func TryCanvasRenderingContext2DRotate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Translate
//go:noescape
func HasCanvasRenderingContext2DTranslate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Translate
//go:noescape
func CanvasRenderingContext2DTranslateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Translate
//go:noescape
func CallCanvasRenderingContext2DTranslate(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Translate
//go:noescape
func TryCanvasRenderingContext2DTranslate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Transform
//go:noescape
func HasCanvasRenderingContext2DTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Transform
//go:noescape
func CanvasRenderingContext2DTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Transform
//go:noescape
func CallCanvasRenderingContext2DTransform(
	this js.Ref, retPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Transform
//go:noescape
func TryCanvasRenderingContext2DTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_GetTransform
//go:noescape
func HasCanvasRenderingContext2DGetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_GetTransform
//go:noescape
func CanvasRenderingContext2DGetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_GetTransform
//go:noescape
func CallCanvasRenderingContext2DGetTransform(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_GetTransform
//go:noescape
func TryCanvasRenderingContext2DGetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_SetTransform
//go:noescape
func HasCanvasRenderingContext2DSetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_SetTransform
//go:noescape
func CanvasRenderingContext2DSetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_SetTransform
//go:noescape
func CallCanvasRenderingContext2DSetTransform(
	this js.Ref, retPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_SetTransform
//go:noescape
func TryCanvasRenderingContext2DSetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_SetTransform1
//go:noescape
func HasCanvasRenderingContext2DSetTransform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_SetTransform1
//go:noescape
func CanvasRenderingContext2DSetTransform1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_SetTransform1
//go:noescape
func CallCanvasRenderingContext2DSetTransform1(
	this js.Ref, retPtr unsafe.Pointer,
	transform unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_SetTransform1
//go:noescape
func TryCanvasRenderingContext2DSetTransform1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transform unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_SetTransform2
//go:noescape
func HasCanvasRenderingContext2DSetTransform2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_SetTransform2
//go:noescape
func CanvasRenderingContext2DSetTransform2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_SetTransform2
//go:noescape
func CallCanvasRenderingContext2DSetTransform2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_SetTransform2
//go:noescape
func TryCanvasRenderingContext2DSetTransform2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ResetTransform
//go:noescape
func HasCanvasRenderingContext2DResetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ResetTransform
//go:noescape
func CanvasRenderingContext2DResetTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ResetTransform
//go:noescape
func CallCanvasRenderingContext2DResetTransform(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ResetTransform
//go:noescape
func TryCanvasRenderingContext2DResetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func HasCanvasRenderingContext2DCreateLinearGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func CanvasRenderingContext2DCreateLinearGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func CallCanvasRenderingContext2DCreateLinearGradient(
	this js.Ref, retPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	x1 float64,
	y1 float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func TryCanvasRenderingContext2DCreateLinearGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	x1 float64,
	y1 float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func HasCanvasRenderingContext2DCreateRadialGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func CanvasRenderingContext2DCreateRadialGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func CallCanvasRenderingContext2DCreateRadialGradient(
	this js.Ref, retPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	r0 float64,
	x1 float64,
	y1 float64,
	r1 float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func TryCanvasRenderingContext2DCreateRadialGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	r0 float64,
	x1 float64,
	y1 float64,
	r1 float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateConicGradient
//go:noescape
func HasCanvasRenderingContext2DCreateConicGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateConicGradient
//go:noescape
func CanvasRenderingContext2DCreateConicGradientFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateConicGradient
//go:noescape
func CallCanvasRenderingContext2DCreateConicGradient(
	this js.Ref, retPtr unsafe.Pointer,
	startAngle float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateConicGradient
//go:noescape
func TryCanvasRenderingContext2DCreateConicGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	startAngle float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreatePattern
//go:noescape
func HasCanvasRenderingContext2DCreatePattern(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreatePattern
//go:noescape
func CanvasRenderingContext2DCreatePatternFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreatePattern
//go:noescape
func CallCanvasRenderingContext2DCreatePattern(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	repetition js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreatePattern
//go:noescape
func TryCanvasRenderingContext2DCreatePattern(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	repetition js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ClearRect
//go:noescape
func HasCanvasRenderingContext2DClearRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ClearRect
//go:noescape
func CanvasRenderingContext2DClearRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ClearRect
//go:noescape
func CallCanvasRenderingContext2DClearRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ClearRect
//go:noescape
func TryCanvasRenderingContext2DClearRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_FillRect
//go:noescape
func HasCanvasRenderingContext2DFillRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_FillRect
//go:noescape
func CanvasRenderingContext2DFillRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_FillRect
//go:noescape
func CallCanvasRenderingContext2DFillRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_FillRect
//go:noescape
func TryCanvasRenderingContext2DFillRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_StrokeRect
//go:noescape
func HasCanvasRenderingContext2DStrokeRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_StrokeRect
//go:noescape
func CanvasRenderingContext2DStrokeRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_StrokeRect
//go:noescape
func CallCanvasRenderingContext2DStrokeRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_StrokeRect
//go:noescape
func TryCanvasRenderingContext2DStrokeRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_BeginPath
//go:noescape
func HasCanvasRenderingContext2DBeginPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_BeginPath
//go:noescape
func CanvasRenderingContext2DBeginPathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_BeginPath
//go:noescape
func CallCanvasRenderingContext2DBeginPath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_BeginPath
//go:noescape
func TryCanvasRenderingContext2DBeginPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Fill
//go:noescape
func HasCanvasRenderingContext2DFill(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Fill
//go:noescape
func CanvasRenderingContext2DFillFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Fill
//go:noescape
func CallCanvasRenderingContext2DFill(
	this js.Ref, retPtr unsafe.Pointer,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Fill
//go:noescape
func TryCanvasRenderingContext2DFill(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Fill1
//go:noescape
func HasCanvasRenderingContext2DFill1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Fill1
//go:noescape
func CanvasRenderingContext2DFill1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Fill1
//go:noescape
func CallCanvasRenderingContext2DFill1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Fill1
//go:noescape
func TryCanvasRenderingContext2DFill1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Fill2
//go:noescape
func HasCanvasRenderingContext2DFill2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Fill2
//go:noescape
func CanvasRenderingContext2DFill2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Fill2
//go:noescape
func CallCanvasRenderingContext2DFill2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Fill2
//go:noescape
func TryCanvasRenderingContext2DFill2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Fill3
//go:noescape
func HasCanvasRenderingContext2DFill3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Fill3
//go:noescape
func CanvasRenderingContext2DFill3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Fill3
//go:noescape
func CallCanvasRenderingContext2DFill3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Fill3
//go:noescape
func TryCanvasRenderingContext2DFill3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Stroke
//go:noescape
func HasCanvasRenderingContext2DStroke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Stroke
//go:noescape
func CanvasRenderingContext2DStrokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Stroke
//go:noescape
func CallCanvasRenderingContext2DStroke(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Stroke
//go:noescape
func TryCanvasRenderingContext2DStroke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Stroke1
//go:noescape
func HasCanvasRenderingContext2DStroke1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Stroke1
//go:noescape
func CanvasRenderingContext2DStroke1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Stroke1
//go:noescape
func CallCanvasRenderingContext2DStroke1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Stroke1
//go:noescape
func TryCanvasRenderingContext2DStroke1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Clip
//go:noescape
func HasCanvasRenderingContext2DClip(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Clip
//go:noescape
func CanvasRenderingContext2DClipFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Clip
//go:noescape
func CallCanvasRenderingContext2DClip(
	this js.Ref, retPtr unsafe.Pointer,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Clip
//go:noescape
func TryCanvasRenderingContext2DClip(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Clip1
//go:noescape
func HasCanvasRenderingContext2DClip1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Clip1
//go:noescape
func CanvasRenderingContext2DClip1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Clip1
//go:noescape
func CallCanvasRenderingContext2DClip1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Clip1
//go:noescape
func TryCanvasRenderingContext2DClip1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Clip2
//go:noescape
func HasCanvasRenderingContext2DClip2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Clip2
//go:noescape
func CanvasRenderingContext2DClip2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Clip2
//go:noescape
func CallCanvasRenderingContext2DClip2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Clip2
//go:noescape
func TryCanvasRenderingContext2DClip2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Clip3
//go:noescape
func HasCanvasRenderingContext2DClip3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Clip3
//go:noescape
func CanvasRenderingContext2DClip3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Clip3
//go:noescape
func CallCanvasRenderingContext2DClip3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Clip3
//go:noescape
func TryCanvasRenderingContext2DClip3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInPath
//go:noescape
func HasCanvasRenderingContext2DIsPointInPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInPath
//go:noescape
func CanvasRenderingContext2DIsPointInPathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInPath
//go:noescape
func CallCanvasRenderingContext2DIsPointInPath(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInPath
//go:noescape
func TryCanvasRenderingContext2DIsPointInPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInPath1
//go:noescape
func HasCanvasRenderingContext2DIsPointInPath1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInPath1
//go:noescape
func CanvasRenderingContext2DIsPointInPath1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInPath1
//go:noescape
func CallCanvasRenderingContext2DIsPointInPath1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInPath1
//go:noescape
func TryCanvasRenderingContext2DIsPointInPath1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInPath2
//go:noescape
func HasCanvasRenderingContext2DIsPointInPath2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInPath2
//go:noescape
func CanvasRenderingContext2DIsPointInPath2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInPath2
//go:noescape
func CallCanvasRenderingContext2DIsPointInPath2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64,
	fillRule uint32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInPath2
//go:noescape
func TryCanvasRenderingContext2DIsPointInPath2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInPath3
//go:noescape
func HasCanvasRenderingContext2DIsPointInPath3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInPath3
//go:noescape
func CanvasRenderingContext2DIsPointInPath3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInPath3
//go:noescape
func CallCanvasRenderingContext2DIsPointInPath3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInPath3
//go:noescape
func TryCanvasRenderingContext2DIsPointInPath3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInStroke
//go:noescape
func HasCanvasRenderingContext2DIsPointInStroke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInStroke
//go:noescape
func CanvasRenderingContext2DIsPointInStrokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInStroke
//go:noescape
func CallCanvasRenderingContext2DIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInStroke
//go:noescape
func TryCanvasRenderingContext2DIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func HasCanvasRenderingContext2DIsPointInStroke1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func CanvasRenderingContext2DIsPointInStroke1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func CallCanvasRenderingContext2DIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func TryCanvasRenderingContext2DIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_DrawFocusIfNeeded
//go:noescape
func HasCanvasRenderingContext2DDrawFocusIfNeeded(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_DrawFocusIfNeeded
//go:noescape
func CanvasRenderingContext2DDrawFocusIfNeededFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_DrawFocusIfNeeded
//go:noescape
func CallCanvasRenderingContext2DDrawFocusIfNeeded(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_DrawFocusIfNeeded
//go:noescape
func TryCanvasRenderingContext2DDrawFocusIfNeeded(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_DrawFocusIfNeeded1
//go:noescape
func HasCanvasRenderingContext2DDrawFocusIfNeeded1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_DrawFocusIfNeeded1
//go:noescape
func CanvasRenderingContext2DDrawFocusIfNeeded1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_DrawFocusIfNeeded1
//go:noescape
func CallCanvasRenderingContext2DDrawFocusIfNeeded1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	element js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_DrawFocusIfNeeded1
//go:noescape
func TryCanvasRenderingContext2DDrawFocusIfNeeded1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ScrollPathIntoView
//go:noescape
func HasCanvasRenderingContext2DScrollPathIntoView(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ScrollPathIntoView
//go:noescape
func CanvasRenderingContext2DScrollPathIntoViewFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ScrollPathIntoView
//go:noescape
func CallCanvasRenderingContext2DScrollPathIntoView(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ScrollPathIntoView
//go:noescape
func TryCanvasRenderingContext2DScrollPathIntoView(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ScrollPathIntoView1
//go:noescape
func HasCanvasRenderingContext2DScrollPathIntoView1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ScrollPathIntoView1
//go:noescape
func CanvasRenderingContext2DScrollPathIntoView1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ScrollPathIntoView1
//go:noescape
func CallCanvasRenderingContext2DScrollPathIntoView1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ScrollPathIntoView1
//go:noescape
func TryCanvasRenderingContext2DScrollPathIntoView1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_FillText
//go:noescape
func HasCanvasRenderingContext2DFillText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_FillText
//go:noescape
func CanvasRenderingContext2DFillTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_FillText
//go:noescape
func CallCanvasRenderingContext2DFillText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_FillText
//go:noescape
func TryCanvasRenderingContext2DFillText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_FillText1
//go:noescape
func HasCanvasRenderingContext2DFillText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_FillText1
//go:noescape
func CanvasRenderingContext2DFillText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_FillText1
//go:noescape
func CallCanvasRenderingContext2DFillText1(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_FillText1
//go:noescape
func TryCanvasRenderingContext2DFillText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_StrokeText
//go:noescape
func HasCanvasRenderingContext2DStrokeText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_StrokeText
//go:noescape
func CanvasRenderingContext2DStrokeTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_StrokeText
//go:noescape
func CallCanvasRenderingContext2DStrokeText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_StrokeText
//go:noescape
func TryCanvasRenderingContext2DStrokeText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_StrokeText1
//go:noescape
func HasCanvasRenderingContext2DStrokeText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_StrokeText1
//go:noescape
func CanvasRenderingContext2DStrokeText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_StrokeText1
//go:noescape
func CallCanvasRenderingContext2DStrokeText1(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_StrokeText1
//go:noescape
func TryCanvasRenderingContext2DStrokeText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_MeasureText
//go:noescape
func HasCanvasRenderingContext2DMeasureText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_MeasureText
//go:noescape
func CanvasRenderingContext2DMeasureTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_MeasureText
//go:noescape
func CallCanvasRenderingContext2DMeasureText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_MeasureText
//go:noescape
func TryCanvasRenderingContext2DMeasureText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_DrawImage
//go:noescape
func HasCanvasRenderingContext2DDrawImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_DrawImage
//go:noescape
func CanvasRenderingContext2DDrawImageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_DrawImage
//go:noescape
func CallCanvasRenderingContext2DDrawImage(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_DrawImage
//go:noescape
func TryCanvasRenderingContext2DDrawImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_DrawImage1
//go:noescape
func HasCanvasRenderingContext2DDrawImage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_DrawImage1
//go:noescape
func CanvasRenderingContext2DDrawImage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_DrawImage1
//go:noescape
func CallCanvasRenderingContext2DDrawImage1(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64,
	dw float64,
	dh float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_DrawImage1
//go:noescape
func TryCanvasRenderingContext2DDrawImage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64,
	dw float64,
	dh float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_DrawImage2
//go:noescape
func HasCanvasRenderingContext2DDrawImage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_DrawImage2
//go:noescape
func CanvasRenderingContext2DDrawImage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_DrawImage2
//go:noescape
func CallCanvasRenderingContext2DDrawImage2(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	sx float64,
	sy float64,
	sw float64,
	sh float64,
	dx float64,
	dy float64,
	dw float64,
	dh float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_DrawImage2
//go:noescape
func TryCanvasRenderingContext2DDrawImage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	sx float64,
	sy float64,
	sw float64,
	sh float64,
	dx float64,
	dy float64,
	dw float64,
	dh float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateImageData
//go:noescape
func HasCanvasRenderingContext2DCreateImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateImageData
//go:noescape
func CanvasRenderingContext2DCreateImageDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateImageData
//go:noescape
func CallCanvasRenderingContext2DCreateImageData(
	this js.Ref, retPtr unsafe.Pointer,
	sw int32,
	sh int32,
	settings unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateImageData
//go:noescape
func TryCanvasRenderingContext2DCreateImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sw int32,
	sh int32,
	settings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateImageData1
//go:noescape
func HasCanvasRenderingContext2DCreateImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateImageData1
//go:noescape
func CanvasRenderingContext2DCreateImageData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateImageData1
//go:noescape
func CallCanvasRenderingContext2DCreateImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateImageData1
//go:noescape
func TryCanvasRenderingContext2DCreateImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_CreateImageData2
//go:noescape
func HasCanvasRenderingContext2DCreateImageData2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_CreateImageData2
//go:noescape
func CanvasRenderingContext2DCreateImageData2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_CreateImageData2
//go:noescape
func CallCanvasRenderingContext2DCreateImageData2(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_CreateImageData2
//go:noescape
func TryCanvasRenderingContext2DCreateImageData2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_GetImageData
//go:noescape
func HasCanvasRenderingContext2DGetImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_GetImageData
//go:noescape
func CanvasRenderingContext2DGetImageDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_GetImageData
//go:noescape
func CallCanvasRenderingContext2DGetImageData(
	this js.Ref, retPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	settings unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_GetImageData
//go:noescape
func TryCanvasRenderingContext2DGetImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	settings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_GetImageData1
//go:noescape
func HasCanvasRenderingContext2DGetImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_GetImageData1
//go:noescape
func CanvasRenderingContext2DGetImageData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_GetImageData1
//go:noescape
func CallCanvasRenderingContext2DGetImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_GetImageData1
//go:noescape
func TryCanvasRenderingContext2DGetImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_PutImageData
//go:noescape
func HasCanvasRenderingContext2DPutImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_PutImageData
//go:noescape
func CanvasRenderingContext2DPutImageDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_PutImageData
//go:noescape
func CallCanvasRenderingContext2DPutImageData(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_PutImageData
//go:noescape
func TryCanvasRenderingContext2DPutImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_PutImageData1
//go:noescape
func HasCanvasRenderingContext2DPutImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_PutImageData1
//go:noescape
func CanvasRenderingContext2DPutImageData1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_PutImageData1
//go:noescape
func CallCanvasRenderingContext2DPutImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32,
	dirtyX int32,
	dirtyY int32,
	dirtyWidth int32,
	dirtyHeight int32)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_PutImageData1
//go:noescape
func TryCanvasRenderingContext2DPutImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32,
	dirtyX int32,
	dirtyY int32,
	dirtyWidth int32,
	dirtyHeight int32) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_SetLineDash
//go:noescape
func HasCanvasRenderingContext2DSetLineDash(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_SetLineDash
//go:noescape
func CanvasRenderingContext2DSetLineDashFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_SetLineDash
//go:noescape
func CallCanvasRenderingContext2DSetLineDash(
	this js.Ref, retPtr unsafe.Pointer,
	segments js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_SetLineDash
//go:noescape
func TryCanvasRenderingContext2DSetLineDash(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	segments js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_GetLineDash
//go:noescape
func HasCanvasRenderingContext2DGetLineDash(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_GetLineDash
//go:noescape
func CanvasRenderingContext2DGetLineDashFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_GetLineDash
//go:noescape
func CallCanvasRenderingContext2DGetLineDash(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_GetLineDash
//go:noescape
func TryCanvasRenderingContext2DGetLineDash(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ClosePath
//go:noescape
func HasCanvasRenderingContext2DClosePath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ClosePath
//go:noescape
func CanvasRenderingContext2DClosePathFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ClosePath
//go:noescape
func CallCanvasRenderingContext2DClosePath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ClosePath
//go:noescape
func TryCanvasRenderingContext2DClosePath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_MoveTo
//go:noescape
func HasCanvasRenderingContext2DMoveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_MoveTo
//go:noescape
func CanvasRenderingContext2DMoveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_MoveTo
//go:noescape
func CallCanvasRenderingContext2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_MoveTo
//go:noescape
func TryCanvasRenderingContext2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_LineTo
//go:noescape
func HasCanvasRenderingContext2DLineTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_LineTo
//go:noescape
func CanvasRenderingContext2DLineToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_LineTo
//go:noescape
func CallCanvasRenderingContext2DLineTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_LineTo
//go:noescape
func TryCanvasRenderingContext2DLineTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func HasCanvasRenderingContext2DQuadraticCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func CanvasRenderingContext2DQuadraticCurveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func CallCanvasRenderingContext2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func TryCanvasRenderingContext2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_BezierCurveTo
//go:noescape
func HasCanvasRenderingContext2DBezierCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_BezierCurveTo
//go:noescape
func CanvasRenderingContext2DBezierCurveToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_BezierCurveTo
//go:noescape
func CallCanvasRenderingContext2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_BezierCurveTo
//go:noescape
func TryCanvasRenderingContext2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_ArcTo
//go:noescape
func HasCanvasRenderingContext2DArcTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_ArcTo
//go:noescape
func CanvasRenderingContext2DArcToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_ArcTo
//go:noescape
func CallCanvasRenderingContext2DArcTo(
	this js.Ref, retPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_ArcTo
//go:noescape
func TryCanvasRenderingContext2DArcTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Rect
//go:noescape
func HasCanvasRenderingContext2DRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Rect
//go:noescape
func CanvasRenderingContext2DRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Rect
//go:noescape
func CallCanvasRenderingContext2DRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Rect
//go:noescape
func TryCanvasRenderingContext2DRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_RoundRect
//go:noescape
func HasCanvasRenderingContext2DRoundRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_RoundRect
//go:noescape
func CanvasRenderingContext2DRoundRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_RoundRect
//go:noescape
func CallCanvasRenderingContext2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_RoundRect
//go:noescape
func TryCanvasRenderingContext2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_RoundRect1
//go:noescape
func HasCanvasRenderingContext2DRoundRect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_RoundRect1
//go:noescape
func CanvasRenderingContext2DRoundRect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_RoundRect1
//go:noescape
func CallCanvasRenderingContext2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_RoundRect1
//go:noescape
func TryCanvasRenderingContext2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Arc
//go:noescape
func HasCanvasRenderingContext2DArc(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Arc
//go:noescape
func CanvasRenderingContext2DArcFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Arc
//go:noescape
func CallCanvasRenderingContext2DArc(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Arc
//go:noescape
func TryCanvasRenderingContext2DArc(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Arc1
//go:noescape
func HasCanvasRenderingContext2DArc1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Arc1
//go:noescape
func CanvasRenderingContext2DArc1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Arc1
//go:noescape
func CallCanvasRenderingContext2DArc1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Arc1
//go:noescape
func TryCanvasRenderingContext2DArc1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Ellipse
//go:noescape
func HasCanvasRenderingContext2DEllipse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Ellipse
//go:noescape
func CanvasRenderingContext2DEllipseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Ellipse
//go:noescape
func CallCanvasRenderingContext2DEllipse(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Ellipse
//go:noescape
func TryCanvasRenderingContext2DEllipse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasRenderingContext2D_Ellipse1
//go:noescape
func HasCanvasRenderingContext2DEllipse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasRenderingContext2D_Ellipse1
//go:noescape
func CanvasRenderingContext2DEllipse1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasRenderingContext2D_Ellipse1
//go:noescape
func CallCanvasRenderingContext2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_CanvasRenderingContext2D_Ellipse1
//go:noescape
func TryCanvasRenderingContext2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLCanvasElement_HTMLCanvasElement
//go:noescape
func NewHTMLCanvasElementByHTMLCanvasElement() js.Ref

//go:wasmimport plat/js/web get_HTMLCanvasElement_Width
//go:noescape
func GetHTMLCanvasElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLCanvasElement_Width
//go:noescape
func SetHTMLCanvasElementWidth(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLCanvasElement_Height
//go:noescape
func GetHTMLCanvasElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLCanvasElement_Height
//go:noescape
func SetHTMLCanvasElementHeight(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web has_HTMLCanvasElement_GetContext
//go:noescape
func HasHTMLCanvasElementGetContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_GetContext
//go:noescape
func HTMLCanvasElementGetContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_GetContext
//go:noescape
func CallHTMLCanvasElementGetContext(
	this js.Ref, retPtr unsafe.Pointer,
	contextId js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_GetContext
//go:noescape
func TryHTMLCanvasElementGetContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextId js.Ref,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_GetContext1
//go:noescape
func HasHTMLCanvasElementGetContext1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_GetContext1
//go:noescape
func HTMLCanvasElementGetContext1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_GetContext1
//go:noescape
func CallHTMLCanvasElementGetContext1(
	this js.Ref, retPtr unsafe.Pointer,
	contextId js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_GetContext1
//go:noescape
func TryHTMLCanvasElementGetContext1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToDataURL
//go:noescape
func HasHTMLCanvasElementToDataURL(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToDataURL
//go:noescape
func HTMLCanvasElementToDataURLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToDataURL
//go:noescape
func CallHTMLCanvasElementToDataURL(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref,
	quality js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToDataURL
//go:noescape
func TryHTMLCanvasElementToDataURL(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref,
	quality js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToDataURL1
//go:noescape
func HasHTMLCanvasElementToDataURL1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToDataURL1
//go:noescape
func HTMLCanvasElementToDataURL1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToDataURL1
//go:noescape
func CallHTMLCanvasElementToDataURL1(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToDataURL1
//go:noescape
func TryHTMLCanvasElementToDataURL1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToDataURL2
//go:noescape
func HasHTMLCanvasElementToDataURL2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToDataURL2
//go:noescape
func HTMLCanvasElementToDataURL2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToDataURL2
//go:noescape
func CallHTMLCanvasElementToDataURL2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToDataURL2
//go:noescape
func TryHTMLCanvasElementToDataURL2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToBlob
//go:noescape
func HasHTMLCanvasElementToBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToBlob
//go:noescape
func HTMLCanvasElementToBlobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToBlob
//go:noescape
func CallHTMLCanvasElementToBlob(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref,
	typ js.Ref,
	quality js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToBlob
//go:noescape
func TryHTMLCanvasElementToBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref,
	typ js.Ref,
	quality js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToBlob1
//go:noescape
func HasHTMLCanvasElementToBlob1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToBlob1
//go:noescape
func HTMLCanvasElementToBlob1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToBlob1
//go:noescape
func CallHTMLCanvasElementToBlob1(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref,
	typ js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToBlob1
//go:noescape
func TryHTMLCanvasElementToBlob1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_ToBlob2
//go:noescape
func HasHTMLCanvasElementToBlob2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_ToBlob2
//go:noescape
func HTMLCanvasElementToBlob2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_ToBlob2
//go:noescape
func CallHTMLCanvasElementToBlob2(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_HTMLCanvasElement_ToBlob2
//go:noescape
func TryHTMLCanvasElementToBlob2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_TransferControlToOffscreen
//go:noescape
func HasHTMLCanvasElementTransferControlToOffscreen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_TransferControlToOffscreen
//go:noescape
func HTMLCanvasElementTransferControlToOffscreenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_TransferControlToOffscreen
//go:noescape
func CallHTMLCanvasElementTransferControlToOffscreen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLCanvasElement_TransferControlToOffscreen
//go:noescape
func TryHTMLCanvasElementTransferControlToOffscreen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_CaptureStream
//go:noescape
func HasHTMLCanvasElementCaptureStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_CaptureStream
//go:noescape
func HTMLCanvasElementCaptureStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_CaptureStream
//go:noescape
func CallHTMLCanvasElementCaptureStream(
	this js.Ref, retPtr unsafe.Pointer,
	frameRequestRate float64)

//go:wasmimport plat/js/web try_HTMLCanvasElement_CaptureStream
//go:noescape
func TryHTMLCanvasElementCaptureStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	frameRequestRate float64) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLCanvasElement_CaptureStream1
//go:noescape
func HasHTMLCanvasElementCaptureStream1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLCanvasElement_CaptureStream1
//go:noescape
func HTMLCanvasElementCaptureStream1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCanvasElement_CaptureStream1
//go:noescape
func CallHTMLCanvasElementCaptureStream1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLCanvasElement_CaptureStream1
//go:noescape
func TryHTMLCanvasElementCaptureStream1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_BarcodeDetector_BarcodeDetector
//go:noescape
func NewBarcodeDetectorByBarcodeDetector(
	barcodeDetectorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_BarcodeDetector_BarcodeDetector1
//go:noescape
func NewBarcodeDetectorByBarcodeDetector1() js.Ref

//go:wasmimport plat/js/web has_BarcodeDetector_GetSupportedFormats
//go:noescape
func HasBarcodeDetectorGetSupportedFormats(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BarcodeDetector_GetSupportedFormats
//go:noescape
func BarcodeDetectorGetSupportedFormatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BarcodeDetector_GetSupportedFormats
//go:noescape
func CallBarcodeDetectorGetSupportedFormats(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BarcodeDetector_GetSupportedFormats
//go:noescape
func TryBarcodeDetectorGetSupportedFormats(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BarcodeDetector_Detect
//go:noescape
func HasBarcodeDetectorDetect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BarcodeDetector_Detect
//go:noescape
func BarcodeDetectorDetectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BarcodeDetector_Detect
//go:noescape
func CallBarcodeDetectorDetect(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref)

//go:wasmimport plat/js/web try_BarcodeDetector_Detect
//go:noescape
func TryBarcodeDetectorDetect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_CompositeOperationOrAuto
//go:noescape
func ConstOfCompositeOperationOrAuto(str js.Ref) uint32

//go:wasmimport plat/js/web store_BaseComputedKeyframe
//go:noescape
func BaseComputedKeyframeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BaseComputedKeyframe
//go:noescape
func BaseComputedKeyframeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BaseKeyframe
//go:noescape
func BaseKeyframeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BaseKeyframe
//go:noescape
func BaseKeyframeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BasePropertyIndexedKeyframe
//go:noescape
func BasePropertyIndexedKeyframeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BasePropertyIndexedKeyframe
//go:noescape
func BasePropertyIndexedKeyframeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_BatteryManager_Charging
//go:noescape
func GetBatteryManagerCharging(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BatteryManager_ChargingTime
//go:noescape
func GetBatteryManagerChargingTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BatteryManager_DischargingTime
//go:noescape
func GetBatteryManagerDischargingTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BatteryManager_Level
//go:noescape
func GetBatteryManagerLevel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PromptResponseObject
//go:noescape
func PromptResponseObjectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PromptResponseObject
//go:noescape
func PromptResponseObjectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BeforeInstallPromptEvent_BeforeInstallPromptEvent
//go:noescape
func NewBeforeInstallPromptEventByBeforeInstallPromptEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_BeforeInstallPromptEvent_BeforeInstallPromptEvent1
//go:noescape
func NewBeforeInstallPromptEventByBeforeInstallPromptEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web has_BeforeInstallPromptEvent_Prompt
//go:noescape
func HasBeforeInstallPromptEventPrompt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BeforeInstallPromptEvent_Prompt
//go:noescape
func BeforeInstallPromptEventPromptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BeforeInstallPromptEvent_Prompt
//go:noescape
func CallBeforeInstallPromptEventPrompt(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BeforeInstallPromptEvent_Prompt
//go:noescape
func TryBeforeInstallPromptEventPrompt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_BeforeUnloadEvent_BeforeUnloadEvent
//go:noescape
func NewBeforeUnloadEventByBeforeUnloadEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_BeforeUnloadEvent_BeforeUnloadEvent1
//go:noescape
func NewBeforeUnloadEventByBeforeUnloadEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_BeforeUnloadEvent_ReturnValue
//go:noescape
func GetBeforeUnloadEventReturnValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_BeforeUnloadEvent_ReturnValue
//go:noescape
func SetBeforeUnloadEventReturnValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_PreviousWin
//go:noescape
func PreviousWinJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PreviousWin
//go:noescape
func PreviousWinJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BiddingBrowserSignals
//go:noescape
func BiddingBrowserSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BiddingBrowserSignals
//go:noescape
func BiddingBrowserSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_BinaryType
//go:noescape
func ConstOfBinaryType(str js.Ref) uint32

//go:wasmimport plat/js/web store_BlobEventInit
//go:noescape
func BlobEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BlobEventInit
//go:noescape
func BlobEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BlobEvent_BlobEvent
//go:noescape
func NewBlobEventByBlobEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_BlobEvent_Data
//go:noescape
func GetBlobEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BlobEvent_Timecode
//go:noescape
func GetBlobEventTimecode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_BlockFragmentationType
//go:noescape
func ConstOfBlockFragmentationType(str js.Ref) uint32

//go:wasmimport plat/js/web store_WatchAdvertisementsOptions
//go:noescape
func WatchAdvertisementsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WatchAdvertisementsOptions
//go:noescape
func WatchAdvertisementsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_BluetoothRemoteGATTDescriptor_Characteristic
//go:noescape
func GetBluetoothRemoteGATTDescriptorCharacteristic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTDescriptor_Uuid
//go:noescape
func GetBluetoothRemoteGATTDescriptorUuid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTDescriptor_Value
//go:noescape
func GetBluetoothRemoteGATTDescriptorValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTDescriptor_ReadValue
//go:noescape
func HasBluetoothRemoteGATTDescriptorReadValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTDescriptor_ReadValue
//go:noescape
func BluetoothRemoteGATTDescriptorReadValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTDescriptor_ReadValue
//go:noescape
func CallBluetoothRemoteGATTDescriptorReadValue(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTDescriptor_ReadValue
//go:noescape
func TryBluetoothRemoteGATTDescriptorReadValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTDescriptor_WriteValue
//go:noescape
func HasBluetoothRemoteGATTDescriptorWriteValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTDescriptor_WriteValue
//go:noescape
func BluetoothRemoteGATTDescriptorWriteValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTDescriptor_WriteValue
//go:noescape
func CallBluetoothRemoteGATTDescriptorWriteValue(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTDescriptor_WriteValue
//go:noescape
func TryBluetoothRemoteGATTDescriptorWriteValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_Broadcast
//go:noescape
func GetBluetoothCharacteristicPropertiesBroadcast(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_Read
//go:noescape
func GetBluetoothCharacteristicPropertiesRead(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_WriteWithoutResponse
//go:noescape
func GetBluetoothCharacteristicPropertiesWriteWithoutResponse(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_Write
//go:noescape
func GetBluetoothCharacteristicPropertiesWrite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_Notify
//go:noescape
func GetBluetoothCharacteristicPropertiesNotify(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_Indicate
//go:noescape
func GetBluetoothCharacteristicPropertiesIndicate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_AuthenticatedSignedWrites
//go:noescape
func GetBluetoothCharacteristicPropertiesAuthenticatedSignedWrites(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_ReliableWrite
//go:noescape
func GetBluetoothCharacteristicPropertiesReliableWrite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothCharacteristicProperties_WritableAuxiliaries
//go:noescape
func GetBluetoothCharacteristicPropertiesWritableAuxiliaries(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTCharacteristic_Service
//go:noescape
func GetBluetoothRemoteGATTCharacteristicService(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTCharacteristic_Uuid
//go:noescape
func GetBluetoothRemoteGATTCharacteristicUuid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTCharacteristic_Properties
//go:noescape
func GetBluetoothRemoteGATTCharacteristicProperties(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTCharacteristic_Value
//go:noescape
func GetBluetoothRemoteGATTCharacteristicValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_GetDescriptor
//go:noescape
func HasBluetoothRemoteGATTCharacteristicGetDescriptor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_GetDescriptor
//go:noescape
func BluetoothRemoteGATTCharacteristicGetDescriptorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_GetDescriptor
//go:noescape
func CallBluetoothRemoteGATTCharacteristicGetDescriptor(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_GetDescriptor
//go:noescape
func TryBluetoothRemoteGATTCharacteristicGetDescriptor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_GetDescriptors
//go:noescape
func HasBluetoothRemoteGATTCharacteristicGetDescriptors(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_GetDescriptors
//go:noescape
func BluetoothRemoteGATTCharacteristicGetDescriptorsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_GetDescriptors
//go:noescape
func CallBluetoothRemoteGATTCharacteristicGetDescriptors(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_GetDescriptors
//go:noescape
func TryBluetoothRemoteGATTCharacteristicGetDescriptors(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_GetDescriptors1
//go:noescape
func HasBluetoothRemoteGATTCharacteristicGetDescriptors1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_GetDescriptors1
//go:noescape
func BluetoothRemoteGATTCharacteristicGetDescriptors1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_GetDescriptors1
//go:noescape
func CallBluetoothRemoteGATTCharacteristicGetDescriptors1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_GetDescriptors1
//go:noescape
func TryBluetoothRemoteGATTCharacteristicGetDescriptors1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_ReadValue
//go:noescape
func HasBluetoothRemoteGATTCharacteristicReadValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_ReadValue
//go:noescape
func BluetoothRemoteGATTCharacteristicReadValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_ReadValue
//go:noescape
func CallBluetoothRemoteGATTCharacteristicReadValue(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_ReadValue
//go:noescape
func TryBluetoothRemoteGATTCharacteristicReadValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_WriteValue
//go:noescape
func HasBluetoothRemoteGATTCharacteristicWriteValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_WriteValue
//go:noescape
func BluetoothRemoteGATTCharacteristicWriteValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_WriteValue
//go:noescape
func CallBluetoothRemoteGATTCharacteristicWriteValue(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_WriteValue
//go:noescape
func TryBluetoothRemoteGATTCharacteristicWriteValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_WriteValueWithResponse
//go:noescape
func HasBluetoothRemoteGATTCharacteristicWriteValueWithResponse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_WriteValueWithResponse
//go:noescape
func BluetoothRemoteGATTCharacteristicWriteValueWithResponseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_WriteValueWithResponse
//go:noescape
func CallBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_WriteValueWithResponse
//go:noescape
func TryBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_WriteValueWithoutResponse
//go:noescape
func HasBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_WriteValueWithoutResponse
//go:noescape
func BluetoothRemoteGATTCharacteristicWriteValueWithoutResponseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_WriteValueWithoutResponse
//go:noescape
func CallBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_WriteValueWithoutResponse
//go:noescape
func TryBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_StartNotifications
//go:noescape
func HasBluetoothRemoteGATTCharacteristicStartNotifications(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_StartNotifications
//go:noescape
func BluetoothRemoteGATTCharacteristicStartNotificationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_StartNotifications
//go:noescape
func CallBluetoothRemoteGATTCharacteristicStartNotifications(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_StartNotifications
//go:noescape
func TryBluetoothRemoteGATTCharacteristicStartNotifications(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTCharacteristic_StopNotifications
//go:noescape
func HasBluetoothRemoteGATTCharacteristicStopNotifications(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTCharacteristic_StopNotifications
//go:noescape
func BluetoothRemoteGATTCharacteristicStopNotificationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTCharacteristic_StopNotifications
//go:noescape
func CallBluetoothRemoteGATTCharacteristicStopNotifications(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTCharacteristic_StopNotifications
//go:noescape
func TryBluetoothRemoteGATTCharacteristicStopNotifications(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTService_Device
//go:noescape
func GetBluetoothRemoteGATTServiceDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTService_Uuid
//go:noescape
func GetBluetoothRemoteGATTServiceUuid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTService_IsPrimary
//go:noescape
func GetBluetoothRemoteGATTServiceIsPrimary(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetCharacteristic
//go:noescape
func HasBluetoothRemoteGATTServiceGetCharacteristic(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetCharacteristic
//go:noescape
func BluetoothRemoteGATTServiceGetCharacteristicFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetCharacteristic
//go:noescape
func CallBluetoothRemoteGATTServiceGetCharacteristic(
	this js.Ref, retPtr unsafe.Pointer,
	characteristic js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetCharacteristic
//go:noescape
func TryBluetoothRemoteGATTServiceGetCharacteristic(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristic js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetCharacteristics
//go:noescape
func HasBluetoothRemoteGATTServiceGetCharacteristics(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetCharacteristics
//go:noescape
func BluetoothRemoteGATTServiceGetCharacteristicsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetCharacteristics
//go:noescape
func CallBluetoothRemoteGATTServiceGetCharacteristics(
	this js.Ref, retPtr unsafe.Pointer,
	characteristic js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetCharacteristics
//go:noescape
func TryBluetoothRemoteGATTServiceGetCharacteristics(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	characteristic js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetCharacteristics1
//go:noescape
func HasBluetoothRemoteGATTServiceGetCharacteristics1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetCharacteristics1
//go:noescape
func BluetoothRemoteGATTServiceGetCharacteristics1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetCharacteristics1
//go:noescape
func CallBluetoothRemoteGATTServiceGetCharacteristics1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetCharacteristics1
//go:noescape
func TryBluetoothRemoteGATTServiceGetCharacteristics1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetIncludedService
//go:noescape
func HasBluetoothRemoteGATTServiceGetIncludedService(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetIncludedService
//go:noescape
func BluetoothRemoteGATTServiceGetIncludedServiceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetIncludedService
//go:noescape
func CallBluetoothRemoteGATTServiceGetIncludedService(
	this js.Ref, retPtr unsafe.Pointer,
	service js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetIncludedService
//go:noescape
func TryBluetoothRemoteGATTServiceGetIncludedService(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	service js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetIncludedServices
//go:noescape
func HasBluetoothRemoteGATTServiceGetIncludedServices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetIncludedServices
//go:noescape
func BluetoothRemoteGATTServiceGetIncludedServicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetIncludedServices
//go:noescape
func CallBluetoothRemoteGATTServiceGetIncludedServices(
	this js.Ref, retPtr unsafe.Pointer,
	service js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetIncludedServices
//go:noescape
func TryBluetoothRemoteGATTServiceGetIncludedServices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	service js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTService_GetIncludedServices1
//go:noescape
func HasBluetoothRemoteGATTServiceGetIncludedServices1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTService_GetIncludedServices1
//go:noescape
func BluetoothRemoteGATTServiceGetIncludedServices1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTService_GetIncludedServices1
//go:noescape
func CallBluetoothRemoteGATTServiceGetIncludedServices1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTService_GetIncludedServices1
//go:noescape
func TryBluetoothRemoteGATTServiceGetIncludedServices1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTServer_Device
//go:noescape
func GetBluetoothRemoteGATTServerDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothRemoteGATTServer_Connected
//go:noescape
func GetBluetoothRemoteGATTServerConnected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTServer_Connect
//go:noescape
func HasBluetoothRemoteGATTServerConnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTServer_Connect
//go:noescape
func BluetoothRemoteGATTServerConnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTServer_Connect
//go:noescape
func CallBluetoothRemoteGATTServerConnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTServer_Connect
//go:noescape
func TryBluetoothRemoteGATTServerConnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTServer_Disconnect
//go:noescape
func HasBluetoothRemoteGATTServerDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTServer_Disconnect
//go:noescape
func BluetoothRemoteGATTServerDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTServer_Disconnect
//go:noescape
func CallBluetoothRemoteGATTServerDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTServer_Disconnect
//go:noescape
func TryBluetoothRemoteGATTServerDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTServer_GetPrimaryService
//go:noescape
func HasBluetoothRemoteGATTServerGetPrimaryService(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTServer_GetPrimaryService
//go:noescape
func BluetoothRemoteGATTServerGetPrimaryServiceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTServer_GetPrimaryService
//go:noescape
func CallBluetoothRemoteGATTServerGetPrimaryService(
	this js.Ref, retPtr unsafe.Pointer,
	service js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTServer_GetPrimaryService
//go:noescape
func TryBluetoothRemoteGATTServerGetPrimaryService(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	service js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTServer_GetPrimaryServices
//go:noescape
func HasBluetoothRemoteGATTServerGetPrimaryServices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTServer_GetPrimaryServices
//go:noescape
func BluetoothRemoteGATTServerGetPrimaryServicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTServer_GetPrimaryServices
//go:noescape
func CallBluetoothRemoteGATTServerGetPrimaryServices(
	this js.Ref, retPtr unsafe.Pointer,
	service js.Ref)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTServer_GetPrimaryServices
//go:noescape
func TryBluetoothRemoteGATTServerGetPrimaryServices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	service js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothRemoteGATTServer_GetPrimaryServices1
//go:noescape
func HasBluetoothRemoteGATTServerGetPrimaryServices1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothRemoteGATTServer_GetPrimaryServices1
//go:noescape
func BluetoothRemoteGATTServerGetPrimaryServices1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothRemoteGATTServer_GetPrimaryServices1
//go:noescape
func CallBluetoothRemoteGATTServerGetPrimaryServices1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothRemoteGATTServer_GetPrimaryServices1
//go:noescape
func TryBluetoothRemoteGATTServerGetPrimaryServices1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothDevice_Id
//go:noescape
func GetBluetoothDeviceId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothDevice_Name
//go:noescape
func GetBluetoothDeviceName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothDevice_Gatt
//go:noescape
func GetBluetoothDeviceGatt(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothDevice_WatchingAdvertisements
//go:noescape
func GetBluetoothDeviceWatchingAdvertisements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothDevice_Forget
//go:noescape
func HasBluetoothDeviceForget(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothDevice_Forget
//go:noescape
func BluetoothDeviceForgetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothDevice_Forget
//go:noescape
func CallBluetoothDeviceForget(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothDevice_Forget
//go:noescape
func TryBluetoothDeviceForget(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothDevice_WatchAdvertisements
//go:noescape
func HasBluetoothDeviceWatchAdvertisements(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothDevice_WatchAdvertisements
//go:noescape
func BluetoothDeviceWatchAdvertisementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothDevice_WatchAdvertisements
//go:noescape
func CallBluetoothDeviceWatchAdvertisements(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothDevice_WatchAdvertisements
//go:noescape
func TryBluetoothDeviceWatchAdvertisements(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothDevice_WatchAdvertisements1
//go:noescape
func HasBluetoothDeviceWatchAdvertisements1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothDevice_WatchAdvertisements1
//go:noescape
func BluetoothDeviceWatchAdvertisements1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothDevice_WatchAdvertisements1
//go:noescape
func CallBluetoothDeviceWatchAdvertisements1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BluetoothDevice_WatchAdvertisements1
//go:noescape
func TryBluetoothDeviceWatchAdvertisements1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BluetoothManufacturerDataFilterInit
//go:noescape
func BluetoothManufacturerDataFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothManufacturerDataFilterInit
//go:noescape
func BluetoothManufacturerDataFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
