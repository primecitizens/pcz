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

//go:wasmimport plat/js/web constof_ImageSmoothingQuality
//go:noescape
func ConstOfImageSmoothingQuality(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasLineCap
//go:noescape
func ConstOfCanvasLineCap(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasLineJoin
//go:noescape
func ConstOfCanvasLineJoin(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasTextAlign
//go:noescape
func ConstOfCanvasTextAlign(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasTextBaseline
//go:noescape
func ConstOfCanvasTextBaseline(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasDirection
//go:noescape
func ConstOfCanvasDirection(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasFontKerning
//go:noescape
func ConstOfCanvasFontKerning(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasFontStretch
//go:noescape
func ConstOfCanvasFontStretch(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasFontVariantCaps
//go:noescape
func ConstOfCanvasFontVariantCaps(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CanvasTextRendering
//go:noescape
func ConstOfCanvasTextRendering(str js.Ref) uint32

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_Canvas
//go:noescape
func GetOffscreenCanvasRenderingContext2DCanvas(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_GlobalAlpha
//go:noescape
func GetOffscreenCanvasRenderingContext2DGlobalAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_GlobalAlpha
//go:noescape
func SetOffscreenCanvasRenderingContext2DGlobalAlpha(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_GlobalCompositeOperation
//go:noescape
func GetOffscreenCanvasRenderingContext2DGlobalCompositeOperation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_GlobalCompositeOperation
//go:noescape
func SetOffscreenCanvasRenderingContext2DGlobalCompositeOperation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ImageSmoothingEnabled
//go:noescape
func GetOffscreenCanvasRenderingContext2DImageSmoothingEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ImageSmoothingEnabled
//go:noescape
func SetOffscreenCanvasRenderingContext2DImageSmoothingEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ImageSmoothingQuality
//go:noescape
func GetOffscreenCanvasRenderingContext2DImageSmoothingQuality(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ImageSmoothingQuality
//go:noescape
func SetOffscreenCanvasRenderingContext2DImageSmoothingQuality(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_StrokeStyle
//go:noescape
func GetOffscreenCanvasRenderingContext2DStrokeStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_StrokeStyle
//go:noescape
func SetOffscreenCanvasRenderingContext2DStrokeStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_FillStyle
//go:noescape
func GetOffscreenCanvasRenderingContext2DFillStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_FillStyle
//go:noescape
func SetOffscreenCanvasRenderingContext2DFillStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ShadowOffsetX
//go:noescape
func GetOffscreenCanvasRenderingContext2DShadowOffsetX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ShadowOffsetX
//go:noescape
func SetOffscreenCanvasRenderingContext2DShadowOffsetX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ShadowOffsetY
//go:noescape
func GetOffscreenCanvasRenderingContext2DShadowOffsetY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ShadowOffsetY
//go:noescape
func SetOffscreenCanvasRenderingContext2DShadowOffsetY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ShadowBlur
//go:noescape
func GetOffscreenCanvasRenderingContext2DShadowBlur(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ShadowBlur
//go:noescape
func SetOffscreenCanvasRenderingContext2DShadowBlur(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_ShadowColor
//go:noescape
func GetOffscreenCanvasRenderingContext2DShadowColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_ShadowColor
//go:noescape
func SetOffscreenCanvasRenderingContext2DShadowColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_Filter
//go:noescape
func GetOffscreenCanvasRenderingContext2DFilter(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_Filter
//go:noescape
func SetOffscreenCanvasRenderingContext2DFilter(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_LineWidth
//go:noescape
func GetOffscreenCanvasRenderingContext2DLineWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_LineWidth
//go:noescape
func SetOffscreenCanvasRenderingContext2DLineWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_LineCap
//go:noescape
func GetOffscreenCanvasRenderingContext2DLineCap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_LineCap
//go:noescape
func SetOffscreenCanvasRenderingContext2DLineCap(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_LineJoin
//go:noescape
func GetOffscreenCanvasRenderingContext2DLineJoin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_LineJoin
//go:noescape
func SetOffscreenCanvasRenderingContext2DLineJoin(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_MiterLimit
//go:noescape
func GetOffscreenCanvasRenderingContext2DMiterLimit(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_MiterLimit
//go:noescape
func SetOffscreenCanvasRenderingContext2DMiterLimit(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_LineDashOffset
//go:noescape
func GetOffscreenCanvasRenderingContext2DLineDashOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_LineDashOffset
//go:noescape
func SetOffscreenCanvasRenderingContext2DLineDashOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_Font
//go:noescape
func GetOffscreenCanvasRenderingContext2DFont(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_Font
//go:noescape
func SetOffscreenCanvasRenderingContext2DFont(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_TextAlign
//go:noescape
func GetOffscreenCanvasRenderingContext2DTextAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_TextAlign
//go:noescape
func SetOffscreenCanvasRenderingContext2DTextAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_TextBaseline
//go:noescape
func GetOffscreenCanvasRenderingContext2DTextBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_TextBaseline
//go:noescape
func SetOffscreenCanvasRenderingContext2DTextBaseline(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_Direction
//go:noescape
func GetOffscreenCanvasRenderingContext2DDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_Direction
//go:noescape
func SetOffscreenCanvasRenderingContext2DDirection(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_LetterSpacing
//go:noescape
func GetOffscreenCanvasRenderingContext2DLetterSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_LetterSpacing
//go:noescape
func SetOffscreenCanvasRenderingContext2DLetterSpacing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_FontKerning
//go:noescape
func GetOffscreenCanvasRenderingContext2DFontKerning(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_FontKerning
//go:noescape
func SetOffscreenCanvasRenderingContext2DFontKerning(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_FontStretch
//go:noescape
func GetOffscreenCanvasRenderingContext2DFontStretch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_FontStretch
//go:noescape
func SetOffscreenCanvasRenderingContext2DFontStretch(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_FontVariantCaps
//go:noescape
func GetOffscreenCanvasRenderingContext2DFontVariantCaps(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_FontVariantCaps
//go:noescape
func SetOffscreenCanvasRenderingContext2DFontVariantCaps(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_TextRendering
//go:noescape
func GetOffscreenCanvasRenderingContext2DTextRendering(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_TextRendering
//go:noescape
func SetOffscreenCanvasRenderingContext2DTextRendering(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_OffscreenCanvasRenderingContext2D_WordSpacing
//go:noescape
func GetOffscreenCanvasRenderingContext2DWordSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_OffscreenCanvasRenderingContext2D_WordSpacing
//go:noescape
func SetOffscreenCanvasRenderingContext2DWordSpacing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Commit
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCommit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Commit
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCommit(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Commit
//go:noescape
func CallOffscreenCanvasRenderingContext2DCommit(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Commit
//go:noescape
func TryOffscreenCanvasRenderingContext2DCommit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Scale
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Scale
//go:noescape
func FuncOffscreenCanvasRenderingContext2DScale(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Scale
//go:noescape
func CallOffscreenCanvasRenderingContext2DScale(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Scale
//go:noescape
func TryOffscreenCanvasRenderingContext2DScale(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Rotate
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DRotate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Rotate
//go:noescape
func FuncOffscreenCanvasRenderingContext2DRotate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Rotate
//go:noescape
func CallOffscreenCanvasRenderingContext2DRotate(
	this js.Ref, retPtr unsafe.Pointer,
	angle float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Rotate
//go:noescape
func TryOffscreenCanvasRenderingContext2DRotate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Translate
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DTranslate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Translate
//go:noescape
func FuncOffscreenCanvasRenderingContext2DTranslate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Translate
//go:noescape
func CallOffscreenCanvasRenderingContext2DTranslate(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Translate
//go:noescape
func TryOffscreenCanvasRenderingContext2DTranslate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Transform
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Transform
//go:noescape
func FuncOffscreenCanvasRenderingContext2DTransform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Transform
//go:noescape
func CallOffscreenCanvasRenderingContext2DTransform(
	this js.Ref, retPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Transform
//go:noescape
func TryOffscreenCanvasRenderingContext2DTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_GetTransform
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DGetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_GetTransform
//go:noescape
func FuncOffscreenCanvasRenderingContext2DGetTransform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_GetTransform
//go:noescape
func CallOffscreenCanvasRenderingContext2DGetTransform(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_GetTransform
//go:noescape
func TryOffscreenCanvasRenderingContext2DGetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_SetTransform
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DSetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_SetTransform
//go:noescape
func FuncOffscreenCanvasRenderingContext2DSetTransform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_SetTransform
//go:noescape
func CallOffscreenCanvasRenderingContext2DSetTransform(
	this js.Ref, retPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_SetTransform
//go:noescape
func TryOffscreenCanvasRenderingContext2DSetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a float64,
	b float64,
	c float64,
	d float64,
	e float64,
	f float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_SetTransform1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DSetTransform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_SetTransform1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DSetTransform1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_SetTransform1
//go:noescape
func CallOffscreenCanvasRenderingContext2DSetTransform1(
	this js.Ref, retPtr unsafe.Pointer,
	transform unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_SetTransform1
//go:noescape
func TryOffscreenCanvasRenderingContext2DSetTransform1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transform unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_SetTransform2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DSetTransform2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_SetTransform2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DSetTransform2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_SetTransform2
//go:noescape
func CallOffscreenCanvasRenderingContext2DSetTransform2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_SetTransform2
//go:noescape
func TryOffscreenCanvasRenderingContext2DSetTransform2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_ResetTransform
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DResetTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_ResetTransform
//go:noescape
func FuncOffscreenCanvasRenderingContext2DResetTransform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_ResetTransform
//go:noescape
func CallOffscreenCanvasRenderingContext2DResetTransform(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_ResetTransform
//go:noescape
func TryOffscreenCanvasRenderingContext2DResetTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateLinearGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateLinearGradient(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateLinearGradient(
	this js.Ref, retPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	x1 float64,
	y1 float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateLinearGradient
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateLinearGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	x1 float64,
	y1 float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateRadialGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateRadialGradient(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateRadialGradient(
	this js.Ref, retPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	r0 float64,
	x1 float64,
	y1 float64,
	r1 float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateRadialGradient
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateRadialGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x0 float64,
	y0 float64,
	r0 float64,
	x1 float64,
	y1 float64,
	r1 float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateConicGradient
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateConicGradient(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateConicGradient
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateConicGradient(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateConicGradient
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateConicGradient(
	this js.Ref, retPtr unsafe.Pointer,
	startAngle float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateConicGradient
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateConicGradient(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	startAngle float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreatePattern
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreatePattern(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreatePattern
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreatePattern(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreatePattern
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreatePattern(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	repetition js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreatePattern
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreatePattern(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	repetition js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_ClearRect
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClearRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_ClearRect
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClearRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_ClearRect
//go:noescape
func CallOffscreenCanvasRenderingContext2DClearRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_ClearRect
//go:noescape
func TryOffscreenCanvasRenderingContext2DClearRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_FillRect
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFillRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_FillRect
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFillRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_FillRect
//go:noescape
func CallOffscreenCanvasRenderingContext2DFillRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_FillRect
//go:noescape
func TryOffscreenCanvasRenderingContext2DFillRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_StrokeRect
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DStrokeRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_StrokeRect
//go:noescape
func FuncOffscreenCanvasRenderingContext2DStrokeRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_StrokeRect
//go:noescape
func CallOffscreenCanvasRenderingContext2DStrokeRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_StrokeRect
//go:noescape
func TryOffscreenCanvasRenderingContext2DStrokeRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_BeginPath
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DBeginPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_BeginPath
//go:noescape
func FuncOffscreenCanvasRenderingContext2DBeginPath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_BeginPath
//go:noescape
func CallOffscreenCanvasRenderingContext2DBeginPath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_BeginPath
//go:noescape
func TryOffscreenCanvasRenderingContext2DBeginPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Fill
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFill(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Fill
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFill(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Fill
//go:noescape
func CallOffscreenCanvasRenderingContext2DFill(
	this js.Ref, retPtr unsafe.Pointer,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Fill
//go:noescape
func TryOffscreenCanvasRenderingContext2DFill(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Fill1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFill1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Fill1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFill1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Fill1
//go:noescape
func CallOffscreenCanvasRenderingContext2DFill1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Fill1
//go:noescape
func TryOffscreenCanvasRenderingContext2DFill1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Fill2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFill2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Fill2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFill2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Fill2
//go:noescape
func CallOffscreenCanvasRenderingContext2DFill2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Fill2
//go:noescape
func TryOffscreenCanvasRenderingContext2DFill2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Fill3
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFill3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Fill3
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFill3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Fill3
//go:noescape
func CallOffscreenCanvasRenderingContext2DFill3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Fill3
//go:noescape
func TryOffscreenCanvasRenderingContext2DFill3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Stroke
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DStroke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Stroke
//go:noescape
func FuncOffscreenCanvasRenderingContext2DStroke(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Stroke
//go:noescape
func CallOffscreenCanvasRenderingContext2DStroke(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Stroke
//go:noescape
func TryOffscreenCanvasRenderingContext2DStroke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Stroke1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DStroke1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Stroke1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DStroke1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Stroke1
//go:noescape
func CallOffscreenCanvasRenderingContext2DStroke1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Stroke1
//go:noescape
func TryOffscreenCanvasRenderingContext2DStroke1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Clip
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClip(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Clip
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClip(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Clip
//go:noescape
func CallOffscreenCanvasRenderingContext2DClip(
	this js.Ref, retPtr unsafe.Pointer,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Clip
//go:noescape
func TryOffscreenCanvasRenderingContext2DClip(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Clip1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClip1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Clip1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClip1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Clip1
//go:noescape
func CallOffscreenCanvasRenderingContext2DClip1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Clip1
//go:noescape
func TryOffscreenCanvasRenderingContext2DClip1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Clip2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClip2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Clip2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClip2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Clip2
//go:noescape
func CallOffscreenCanvasRenderingContext2DClip2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Clip2
//go:noescape
func TryOffscreenCanvasRenderingContext2DClip2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Clip3
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClip3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Clip3
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClip3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Clip3
//go:noescape
func CallOffscreenCanvasRenderingContext2DClip3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Clip3
//go:noescape
func TryOffscreenCanvasRenderingContext2DClip3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInPath
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInPath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInPath
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInPath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInPath
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInPath(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInPath
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInPath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInPath1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInPath1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInPath1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInPath1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInPath1
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInPath1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInPath1
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInPath1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInPath2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInPath2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInPath2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInPath2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInPath2
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInPath2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64,
	fillRule uint32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInPath2
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInPath2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64,
	fillRule uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInPath3
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInPath3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInPath3
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInPath3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInPath3
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInPath3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInPath3
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInPath3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInStroke
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInStroke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInStroke
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInStroke(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInStroke
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInStroke
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsPointInStroke1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsPointInStroke1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsPointInStroke1
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_FillText
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFillText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_FillText
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFillText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_FillText
//go:noescape
func CallOffscreenCanvasRenderingContext2DFillText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_FillText
//go:noescape
func TryOffscreenCanvasRenderingContext2DFillText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_FillText1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DFillText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_FillText1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DFillText1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_FillText1
//go:noescape
func CallOffscreenCanvasRenderingContext2DFillText1(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_FillText1
//go:noescape
func TryOffscreenCanvasRenderingContext2DFillText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_StrokeText
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DStrokeText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_StrokeText
//go:noescape
func FuncOffscreenCanvasRenderingContext2DStrokeText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_StrokeText
//go:noescape
func CallOffscreenCanvasRenderingContext2DStrokeText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_StrokeText
//go:noescape
func TryOffscreenCanvasRenderingContext2DStrokeText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64,
	maxWidth float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_StrokeText1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DStrokeText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_StrokeText1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DStrokeText1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_StrokeText1
//go:noescape
func CallOffscreenCanvasRenderingContext2DStrokeText1(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_StrokeText1
//go:noescape
func TryOffscreenCanvasRenderingContext2DStrokeText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_MeasureText
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DMeasureText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_MeasureText
//go:noescape
func FuncOffscreenCanvasRenderingContext2DMeasureText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_MeasureText
//go:noescape
func CallOffscreenCanvasRenderingContext2DMeasureText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_MeasureText
//go:noescape
func TryOffscreenCanvasRenderingContext2DMeasureText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_DrawImage
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DDrawImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_DrawImage
//go:noescape
func FuncOffscreenCanvasRenderingContext2DDrawImage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_DrawImage
//go:noescape
func CallOffscreenCanvasRenderingContext2DDrawImage(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_DrawImage
//go:noescape
func TryOffscreenCanvasRenderingContext2DDrawImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_DrawImage1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DDrawImage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_DrawImage1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DDrawImage1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_DrawImage1
//go:noescape
func CallOffscreenCanvasRenderingContext2DDrawImage1(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64,
	dw float64,
	dh float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_DrawImage1
//go:noescape
func TryOffscreenCanvasRenderingContext2DDrawImage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	dx float64,
	dy float64,
	dw float64,
	dh float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_DrawImage2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DDrawImage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_DrawImage2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DDrawImage2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_DrawImage2
//go:noescape
func CallOffscreenCanvasRenderingContext2DDrawImage2(
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

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_DrawImage2
//go:noescape
func TryOffscreenCanvasRenderingContext2DDrawImage2(
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

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateImageData
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateImageData
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateImageData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateImageData
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateImageData(
	this js.Ref, retPtr unsafe.Pointer,
	sw int32,
	sh int32,
	settings unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateImageData
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sw int32,
	sh int32,
	settings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateImageData1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateImageData1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateImageData1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateImageData1
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateImageData1
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_CreateImageData2
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DCreateImageData2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_CreateImageData2
//go:noescape
func FuncOffscreenCanvasRenderingContext2DCreateImageData2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_CreateImageData2
//go:noescape
func CallOffscreenCanvasRenderingContext2DCreateImageData2(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_CreateImageData2
//go:noescape
func TryOffscreenCanvasRenderingContext2DCreateImageData2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_GetImageData
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DGetImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_GetImageData
//go:noescape
func FuncOffscreenCanvasRenderingContext2DGetImageData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_GetImageData
//go:noescape
func CallOffscreenCanvasRenderingContext2DGetImageData(
	this js.Ref, retPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	settings unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_GetImageData
//go:noescape
func TryOffscreenCanvasRenderingContext2DGetImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32,
	settings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_GetImageData1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DGetImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_GetImageData1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DGetImageData1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_GetImageData1
//go:noescape
func CallOffscreenCanvasRenderingContext2DGetImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_GetImageData1
//go:noescape
func TryOffscreenCanvasRenderingContext2DGetImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx int32,
	sy int32,
	sw int32,
	sh int32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_PutImageData
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DPutImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_PutImageData
//go:noescape
func FuncOffscreenCanvasRenderingContext2DPutImageData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_PutImageData
//go:noescape
func CallOffscreenCanvasRenderingContext2DPutImageData(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_PutImageData
//go:noescape
func TryOffscreenCanvasRenderingContext2DPutImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_PutImageData1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DPutImageData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_PutImageData1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DPutImageData1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_PutImageData1
//go:noescape
func CallOffscreenCanvasRenderingContext2DPutImageData1(
	this js.Ref, retPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32,
	dirtyX int32,
	dirtyY int32,
	dirtyWidth int32,
	dirtyHeight int32)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_PutImageData1
//go:noescape
func TryOffscreenCanvasRenderingContext2DPutImageData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imagedata js.Ref,
	dx int32,
	dy int32,
	dirtyX int32,
	dirtyY int32,
	dirtyWidth int32,
	dirtyHeight int32) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_SetLineDash
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DSetLineDash(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_SetLineDash
//go:noescape
func FuncOffscreenCanvasRenderingContext2DSetLineDash(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_SetLineDash
//go:noescape
func CallOffscreenCanvasRenderingContext2DSetLineDash(
	this js.Ref, retPtr unsafe.Pointer,
	segments js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_SetLineDash
//go:noescape
func TryOffscreenCanvasRenderingContext2DSetLineDash(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	segments js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_GetLineDash
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DGetLineDash(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_GetLineDash
//go:noescape
func FuncOffscreenCanvasRenderingContext2DGetLineDash(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_GetLineDash
//go:noescape
func CallOffscreenCanvasRenderingContext2DGetLineDash(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_GetLineDash
//go:noescape
func TryOffscreenCanvasRenderingContext2DGetLineDash(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_ClosePath
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DClosePath(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_ClosePath
//go:noescape
func FuncOffscreenCanvasRenderingContext2DClosePath(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_ClosePath
//go:noescape
func CallOffscreenCanvasRenderingContext2DClosePath(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_ClosePath
//go:noescape
func TryOffscreenCanvasRenderingContext2DClosePath(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_MoveTo
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DMoveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_MoveTo
//go:noescape
func FuncOffscreenCanvasRenderingContext2DMoveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_MoveTo
//go:noescape
func CallOffscreenCanvasRenderingContext2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_MoveTo
//go:noescape
func TryOffscreenCanvasRenderingContext2DMoveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_LineTo
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DLineTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_LineTo
//go:noescape
func FuncOffscreenCanvasRenderingContext2DLineTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_LineTo
//go:noescape
func CallOffscreenCanvasRenderingContext2DLineTo(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_LineTo
//go:noescape
func TryOffscreenCanvasRenderingContext2DLineTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DQuadraticCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func FuncOffscreenCanvasRenderingContext2DQuadraticCurveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func CallOffscreenCanvasRenderingContext2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_QuadraticCurveTo
//go:noescape
func TryOffscreenCanvasRenderingContext2DQuadraticCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cpx float64,
	cpy float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_BezierCurveTo
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DBezierCurveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_BezierCurveTo
//go:noescape
func FuncOffscreenCanvasRenderingContext2DBezierCurveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_BezierCurveTo
//go:noescape
func CallOffscreenCanvasRenderingContext2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_BezierCurveTo
//go:noescape
func TryOffscreenCanvasRenderingContext2DBezierCurveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cp1x float64,
	cp1y float64,
	cp2x float64,
	cp2y float64,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_ArcTo
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DArcTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_ArcTo
//go:noescape
func FuncOffscreenCanvasRenderingContext2DArcTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_ArcTo
//go:noescape
func CallOffscreenCanvasRenderingContext2DArcTo(
	this js.Ref, retPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_ArcTo
//go:noescape
func TryOffscreenCanvasRenderingContext2DArcTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	radius float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Rect
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Rect
//go:noescape
func FuncOffscreenCanvasRenderingContext2DRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Rect
//go:noescape
func CallOffscreenCanvasRenderingContext2DRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Rect
//go:noescape
func TryOffscreenCanvasRenderingContext2DRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_RoundRect
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DRoundRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_RoundRect
//go:noescape
func FuncOffscreenCanvasRenderingContext2DRoundRect(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_RoundRect
//go:noescape
func CallOffscreenCanvasRenderingContext2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_RoundRect
//go:noescape
func TryOffscreenCanvasRenderingContext2DRoundRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64,
	radii js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_RoundRect1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DRoundRect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_RoundRect1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DRoundRect1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_RoundRect1
//go:noescape
func CallOffscreenCanvasRenderingContext2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_RoundRect1
//go:noescape
func TryOffscreenCanvasRenderingContext2DRoundRect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	w float64,
	h float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Arc
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DArc(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Arc
//go:noescape
func FuncOffscreenCanvasRenderingContext2DArc(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Arc
//go:noescape
func CallOffscreenCanvasRenderingContext2DArc(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Arc
//go:noescape
func TryOffscreenCanvasRenderingContext2DArc(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Arc1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DArc1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Arc1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DArc1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Arc1
//go:noescape
func CallOffscreenCanvasRenderingContext2DArc1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Arc1
//go:noescape
func TryOffscreenCanvasRenderingContext2DArc1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radius float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Ellipse
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DEllipse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Ellipse
//go:noescape
func FuncOffscreenCanvasRenderingContext2DEllipse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Ellipse
//go:noescape
func CallOffscreenCanvasRenderingContext2DEllipse(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Ellipse
//go:noescape
func TryOffscreenCanvasRenderingContext2DEllipse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64,
	counterclockwise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Ellipse1
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DEllipse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Ellipse1
//go:noescape
func FuncOffscreenCanvasRenderingContext2DEllipse1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Ellipse1
//go:noescape
func CallOffscreenCanvasRenderingContext2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Ellipse1
//go:noescape
func TryOffscreenCanvasRenderingContext2DEllipse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	radiusX float64,
	radiusY float64,
	rotation float64,
	startAngle float64,
	endAngle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Save
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DSave(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Save
//go:noescape
func FuncOffscreenCanvasRenderingContext2DSave(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Save
//go:noescape
func CallOffscreenCanvasRenderingContext2DSave(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Save
//go:noescape
func TryOffscreenCanvasRenderingContext2DSave(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Restore
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DRestore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Restore
//go:noescape
func FuncOffscreenCanvasRenderingContext2DRestore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Restore
//go:noescape
func CallOffscreenCanvasRenderingContext2DRestore(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Restore
//go:noescape
func TryOffscreenCanvasRenderingContext2DRestore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_Reset
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_Reset
//go:noescape
func FuncOffscreenCanvasRenderingContext2DReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_Reset
//go:noescape
func CallOffscreenCanvasRenderingContext2DReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_Reset
//go:noescape
func TryOffscreenCanvasRenderingContext2DReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_OffscreenCanvasRenderingContext2D_IsContextLost
//go:noescape
func HasFuncOffscreenCanvasRenderingContext2DIsContextLost(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_OffscreenCanvasRenderingContext2D_IsContextLost
//go:noescape
func FuncOffscreenCanvasRenderingContext2DIsContextLost(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_OffscreenCanvasRenderingContext2D_IsContextLost
//go:noescape
func CallOffscreenCanvasRenderingContext2DIsContextLost(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_OffscreenCanvasRenderingContext2D_IsContextLost
//go:noescape
func TryOffscreenCanvasRenderingContext2DIsContextLost(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ImageBitmapRenderingContext_Canvas
//go:noescape
func GetImageBitmapRenderingContextCanvas(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ImageBitmapRenderingContext_TransferFromImageBitmap
//go:noescape
func HasFuncImageBitmapRenderingContextTransferFromImageBitmap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ImageBitmapRenderingContext_TransferFromImageBitmap
//go:noescape
func FuncImageBitmapRenderingContextTransferFromImageBitmap(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ImageBitmapRenderingContext_TransferFromImageBitmap
//go:noescape
func CallImageBitmapRenderingContextTransferFromImageBitmap(
	this js.Ref, retPtr unsafe.Pointer,
	bitmap js.Ref)

//go:wasmimport plat/js/web try_ImageBitmapRenderingContext_TransferFromImageBitmap
//go:noescape
func TryImageBitmapRenderingContextTransferFromImageBitmap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bitmap js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_WebGLPowerPreference
//go:noescape
func ConstOfWebGLPowerPreference(str js.Ref) uint32

//go:wasmimport plat/js/web store_WebGLContextAttributes
//go:noescape
func WebGLContextAttributesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WebGLContextAttributes
//go:noescape
func WebGLContextAttributesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_WebGLActiveInfo_Size
//go:noescape
func GetWebGLActiveInfoSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebGLActiveInfo_Type
//go:noescape
func GetWebGLActiveInfoType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebGLActiveInfo_Name
//go:noescape
func GetWebGLActiveInfoName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebGLShaderPrecisionFormat_RangeMin
//go:noescape
func GetWebGLShaderPrecisionFormatRangeMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebGLShaderPrecisionFormat_RangeMax
//go:noescape
func GetWebGLShaderPrecisionFormatRangeMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WebGLShaderPrecisionFormat_Precision
//go:noescape
func GetWebGLShaderPrecisionFormatPrecision(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_VideoFrameMetadata
//go:noescape
func VideoFrameMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoFrameMetadata
//go:noescape
func VideoFrameMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_VideoFrameInit
//go:noescape
func VideoFrameInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoFrameInit
//go:noescape
func VideoFrameInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_VideoPixelFormat
//go:noescape
func ConstOfVideoPixelFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_PlaneLayout
//go:noescape
func PlaneLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PlaneLayout
//go:noescape
func PlaneLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_VideoColorPrimaries
//go:noescape
func ConstOfVideoColorPrimaries(str js.Ref) uint32
