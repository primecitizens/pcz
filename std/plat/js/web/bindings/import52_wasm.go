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

//go:wasmimport plat/js/web get_SVGFETileElement_In1
//go:noescape
func GetSVGFETileElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETileElement_X
//go:noescape
func GetSVGFETileElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETileElement_Y
//go:noescape
func GetSVGFETileElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETileElement_Width
//go:noescape
func GetSVGFETileElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETileElement_Height
//go:noescape
func GetSVGFETileElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETileElement_Result
//go:noescape
func GetSVGFETileElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_BaseFrequencyX
//go:noescape
func GetSVGFETurbulenceElementBaseFrequencyX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_BaseFrequencyY
//go:noescape
func GetSVGFETurbulenceElementBaseFrequencyY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_NumOctaves
//go:noescape
func GetSVGFETurbulenceElementNumOctaves(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Seed
//go:noescape
func GetSVGFETurbulenceElementSeed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_StitchTiles
//go:noescape
func GetSVGFETurbulenceElementStitchTiles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Type
//go:noescape
func GetSVGFETurbulenceElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_X
//go:noescape
func GetSVGFETurbulenceElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Y
//go:noescape
func GetSVGFETurbulenceElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Width
//go:noescape
func GetSVGFETurbulenceElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Height
//go:noescape
func GetSVGFETurbulenceElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFETurbulenceElement_Result
//go:noescape
func GetSVGFETurbulenceElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_FilterUnits
//go:noescape
func GetSVGFilterElementFilterUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_PrimitiveUnits
//go:noescape
func GetSVGFilterElementPrimitiveUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_X
//go:noescape
func GetSVGFilterElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_Y
//go:noescape
func GetSVGFilterElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_Width
//go:noescape
func GetSVGFilterElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_Height
//go:noescape
func GetSVGFilterElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFilterElement_Href
//go:noescape
func GetSVGFilterElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGForeignObjectElement_X
//go:noescape
func GetSVGForeignObjectElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGForeignObjectElement_Y
//go:noescape
func GetSVGForeignObjectElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGForeignObjectElement_Width
//go:noescape
func GetSVGForeignObjectElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGForeignObjectElement_Height
//go:noescape
func GetSVGForeignObjectElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGeometryElement_PathLength
//go:noescape
func GetSVGGeometryElementPathLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_IsPointInFill
//go:noescape
func HasFuncSVGGeometryElementIsPointInFill(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_IsPointInFill
//go:noescape
func FuncSVGGeometryElementIsPointInFill(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_IsPointInFill
//go:noescape
func CallSVGGeometryElementIsPointInFill(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGeometryElement_IsPointInFill
//go:noescape
func TrySVGGeometryElementIsPointInFill(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_IsPointInFill1
//go:noescape
func HasFuncSVGGeometryElementIsPointInFill1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_IsPointInFill1
//go:noescape
func FuncSVGGeometryElementIsPointInFill1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_IsPointInFill1
//go:noescape
func CallSVGGeometryElementIsPointInFill1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGeometryElement_IsPointInFill1
//go:noescape
func TrySVGGeometryElementIsPointInFill1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_IsPointInStroke
//go:noescape
func HasFuncSVGGeometryElementIsPointInStroke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_IsPointInStroke
//go:noescape
func FuncSVGGeometryElementIsPointInStroke(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_IsPointInStroke
//go:noescape
func CallSVGGeometryElementIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGeometryElement_IsPointInStroke
//go:noescape
func TrySVGGeometryElementIsPointInStroke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_IsPointInStroke1
//go:noescape
func HasFuncSVGGeometryElementIsPointInStroke1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_IsPointInStroke1
//go:noescape
func FuncSVGGeometryElementIsPointInStroke1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_IsPointInStroke1
//go:noescape
func CallSVGGeometryElementIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGeometryElement_IsPointInStroke1
//go:noescape
func TrySVGGeometryElementIsPointInStroke1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_GetTotalLength
//go:noescape
func HasFuncSVGGeometryElementGetTotalLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_GetTotalLength
//go:noescape
func FuncSVGGeometryElementGetTotalLength(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_GetTotalLength
//go:noescape
func CallSVGGeometryElementGetTotalLength(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGeometryElement_GetTotalLength
//go:noescape
func TrySVGGeometryElementGetTotalLength(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGeometryElement_GetPointAtLength
//go:noescape
func HasFuncSVGGeometryElementGetPointAtLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGeometryElement_GetPointAtLength
//go:noescape
func FuncSVGGeometryElementGetPointAtLength(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGeometryElement_GetPointAtLength
//go:noescape
func CallSVGGeometryElementGetPointAtLength(
	this js.Ref, retPtr unsafe.Pointer,
	distance float32)

//go:wasmimport plat/js/web try_SVGGeometryElement_GetPointAtLength
//go:noescape
func TrySVGGeometryElementGetPointAtLength(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	distance float32) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGradientElement_GradientUnits
//go:noescape
func GetSVGGradientElementGradientUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGradientElement_GradientTransform
//go:noescape
func GetSVGGradientElementGradientTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGradientElement_SpreadMethod
//go:noescape
func GetSVGGradientElementSpreadMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGradientElement_Href
//go:noescape
func GetSVGGradientElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGraphicsElement_Transform
//go:noescape
func GetSVGGraphicsElementTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGraphicsElement_RequiredExtensions
//go:noescape
func GetSVGGraphicsElementRequiredExtensions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGGraphicsElement_SystemLanguage
//go:noescape
func GetSVGGraphicsElementSystemLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGraphicsElement_GetBBox
//go:noescape
func HasFuncSVGGraphicsElementGetBBox(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGraphicsElement_GetBBox
//go:noescape
func FuncSVGGraphicsElementGetBBox(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGraphicsElement_GetBBox
//go:noescape
func CallSVGGraphicsElementGetBBox(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGraphicsElement_GetBBox
//go:noescape
func TrySVGGraphicsElementGetBBox(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGraphicsElement_GetBBox1
//go:noescape
func HasFuncSVGGraphicsElementGetBBox1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGraphicsElement_GetBBox1
//go:noescape
func FuncSVGGraphicsElementGetBBox1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGraphicsElement_GetBBox1
//go:noescape
func CallSVGGraphicsElementGetBBox1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGraphicsElement_GetBBox1
//go:noescape
func TrySVGGraphicsElementGetBBox1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGraphicsElement_GetCTM
//go:noescape
func HasFuncSVGGraphicsElementGetCTM(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGraphicsElement_GetCTM
//go:noescape
func FuncSVGGraphicsElementGetCTM(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGraphicsElement_GetCTM
//go:noescape
func CallSVGGraphicsElementGetCTM(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGraphicsElement_GetCTM
//go:noescape
func TrySVGGraphicsElementGetCTM(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGGraphicsElement_GetScreenCTM
//go:noescape
func HasFuncSVGGraphicsElementGetScreenCTM(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGGraphicsElement_GetScreenCTM
//go:noescape
func FuncSVGGraphicsElementGetScreenCTM(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGGraphicsElement_GetScreenCTM
//go:noescape
func CallSVGGraphicsElementGetScreenCTM(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGGraphicsElement_GetScreenCTM
//go:noescape
func TrySVGGraphicsElementGetScreenCTM(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLineElement_X1
//go:noescape
func GetSVGLineElementX1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLineElement_Y1
//go:noescape
func GetSVGLineElementY1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLineElement_X2
//go:noescape
func GetSVGLineElementX2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLineElement_Y2
//go:noescape
func GetSVGLineElementY2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLinearGradientElement_X1
//go:noescape
func GetSVGLinearGradientElementX1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLinearGradientElement_Y1
//go:noescape
func GetSVGLinearGradientElementY1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLinearGradientElement_X2
//go:noescape
func GetSVGLinearGradientElementX2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLinearGradientElement_Y2
//go:noescape
func GetSVGLinearGradientElementY2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMPathElement_Href
//go:noescape
func GetSVGMPathElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_RefX
//go:noescape
func GetSVGMarkerElementRefX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_RefY
//go:noescape
func GetSVGMarkerElementRefY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_MarkerUnits
//go:noescape
func GetSVGMarkerElementMarkerUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_MarkerWidth
//go:noescape
func GetSVGMarkerElementMarkerWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_MarkerHeight
//go:noescape
func GetSVGMarkerElementMarkerHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_OrientType
//go:noescape
func GetSVGMarkerElementOrientType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_OrientAngle
//go:noescape
func GetSVGMarkerElementOrientAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_Orient
//go:noescape
func GetSVGMarkerElementOrient(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGMarkerElement_Orient
//go:noescape
func SetSVGMarkerElementOrient(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGMarkerElement_ViewBox
//go:noescape
func GetSVGMarkerElementViewBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMarkerElement_PreserveAspectRatio
//go:noescape
func GetSVGMarkerElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGMarkerElement_SetOrientToAuto
//go:noescape
func HasFuncSVGMarkerElementSetOrientToAuto(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGMarkerElement_SetOrientToAuto
//go:noescape
func FuncSVGMarkerElementSetOrientToAuto(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGMarkerElement_SetOrientToAuto
//go:noescape
func CallSVGMarkerElementSetOrientToAuto(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGMarkerElement_SetOrientToAuto
//go:noescape
func TrySVGMarkerElementSetOrientToAuto(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGMarkerElement_SetOrientToAngle
//go:noescape
func HasFuncSVGMarkerElementSetOrientToAngle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGMarkerElement_SetOrientToAngle
//go:noescape
func FuncSVGMarkerElementSetOrientToAngle(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGMarkerElement_SetOrientToAngle
//go:noescape
func CallSVGMarkerElementSetOrientToAngle(
	this js.Ref, retPtr unsafe.Pointer,
	angle js.Ref)

//go:wasmimport plat/js/web try_SVGMarkerElement_SetOrientToAngle
//go:noescape
func TrySVGMarkerElementSetOrientToAngle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_MaskUnits
//go:noescape
func GetSVGMaskElementMaskUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_MaskContentUnits
//go:noescape
func GetSVGMaskElementMaskContentUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_X
//go:noescape
func GetSVGMaskElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_Y
//go:noescape
func GetSVGMaskElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_Width
//go:noescape
func GetSVGMaskElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGMaskElement_Height
//go:noescape
func GetSVGMaskElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_PatternUnits
//go:noescape
func GetSVGPatternElementPatternUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_PatternContentUnits
//go:noescape
func GetSVGPatternElementPatternContentUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_PatternTransform
//go:noescape
func GetSVGPatternElementPatternTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_X
//go:noescape
func GetSVGPatternElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_Y
//go:noescape
func GetSVGPatternElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_Width
//go:noescape
func GetSVGPatternElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_Height
//go:noescape
func GetSVGPatternElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_ViewBox
//go:noescape
func GetSVGPatternElementViewBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_PreserveAspectRatio
//go:noescape
func GetSVGPatternElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPatternElement_Href
//go:noescape
func GetSVGPatternElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPointList_Length
//go:noescape
func GetSVGPointListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPointList_NumberOfItems
//go:noescape
func GetSVGPointListNumberOfItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_Clear
//go:noescape
func HasFuncSVGPointListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_Clear
//go:noescape
func FuncSVGPointListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_Clear
//go:noescape
func CallSVGPointListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGPointList_Clear
//go:noescape
func TrySVGPointListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_Initialize
//go:noescape
func HasFuncSVGPointListInitialize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_Initialize
//go:noescape
func FuncSVGPointListInitialize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_Initialize
//go:noescape
func CallSVGPointListInitialize(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGPointList_Initialize
//go:noescape
func TrySVGPointListInitialize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_GetItem
//go:noescape
func HasFuncSVGPointListGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_GetItem
//go:noescape
func FuncSVGPointListGetItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_GetItem
//go:noescape
func CallSVGPointListGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGPointList_GetItem
//go:noescape
func TrySVGPointListGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_InsertItemBefore
//go:noescape
func HasFuncSVGPointListInsertItemBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_InsertItemBefore
//go:noescape
func FuncSVGPointListInsertItemBefore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_InsertItemBefore
//go:noescape
func CallSVGPointListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGPointList_InsertItemBefore
//go:noescape
func TrySVGPointListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_ReplaceItem
//go:noescape
func HasFuncSVGPointListReplaceItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_ReplaceItem
//go:noescape
func FuncSVGPointListReplaceItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_ReplaceItem
//go:noescape
func CallSVGPointListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGPointList_ReplaceItem
//go:noescape
func TrySVGPointListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_RemoveItem
//go:noescape
func HasFuncSVGPointListRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_RemoveItem
//go:noescape
func FuncSVGPointListRemoveItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_RemoveItem
//go:noescape
func CallSVGPointListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGPointList_RemoveItem
//go:noescape
func TrySVGPointListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_AppendItem
//go:noescape
func HasFuncSVGPointListAppendItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_AppendItem
//go:noescape
func FuncSVGPointListAppendItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_AppendItem
//go:noescape
func CallSVGPointListAppendItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGPointList_AppendItem
//go:noescape
func TrySVGPointListAppendItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGPointList_Set
//go:noescape
func HasFuncSVGPointListSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGPointList_Set
//go:noescape
func FuncSVGPointListSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGPointList_Set
//go:noescape
func CallSVGPointListSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGPointList_Set
//go:noescape
func TrySVGPointListSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPolygonElement_Points
//go:noescape
func GetSVGPolygonElementPoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPolygonElement_AnimatedPoints
//go:noescape
func GetSVGPolygonElementAnimatedPoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPolylineElement_Points
//go:noescape
func GetSVGPolylineElementPoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPolylineElement_AnimatedPoints
//go:noescape
func GetSVGPolylineElementAnimatedPoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_Cx
//go:noescape
func GetSVGRadialGradientElementCx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_Cy
//go:noescape
func GetSVGRadialGradientElementCy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_R
//go:noescape
func GetSVGRadialGradientElementR(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_Fx
//go:noescape
func GetSVGRadialGradientElementFx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_Fy
//go:noescape
func GetSVGRadialGradientElementFy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRadialGradientElement_Fr
//go:noescape
func GetSVGRadialGradientElementFr(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_X
//go:noescape
func GetSVGRectElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_Y
//go:noescape
func GetSVGRectElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_Width
//go:noescape
func GetSVGRectElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_Height
//go:noescape
func GetSVGRectElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_Rx
//go:noescape
func GetSVGRectElementRx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGRectElement_Ry
//go:noescape
func GetSVGRectElementRy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGStopElement_Offset
//go:noescape
func GetSVGStopElementOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGStyleElement_Type
//go:noescape
func GetSVGStyleElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGStyleElement_Type
//go:noescape
func SetSVGStyleElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGStyleElement_Media
//go:noescape
func GetSVGStyleElementMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGStyleElement_Media
//go:noescape
func SetSVGStyleElementMedia(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGStyleElement_Title
//go:noescape
func GetSVGStyleElementTitle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGStyleElement_Title
//go:noescape
func SetSVGStyleElementTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGStyleElement_Sheet
//go:noescape
func GetSVGStyleElementSheet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSymbolElement_ViewBox
//go:noescape
func GetSVGSymbolElementViewBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSymbolElement_PreserveAspectRatio
//go:noescape
func GetSVGSymbolElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextContentElement_TextLength
//go:noescape
func GetSVGTextContentElementTextLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextContentElement_LengthAdjust
//go:noescape
func GetSVGTextContentElementLengthAdjust(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetNumberOfChars
//go:noescape
func HasFuncSVGTextContentElementGetNumberOfChars(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetNumberOfChars
//go:noescape
func FuncSVGTextContentElementGetNumberOfChars(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetNumberOfChars
//go:noescape
func CallSVGTextContentElementGetNumberOfChars(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetNumberOfChars
//go:noescape
func TrySVGTextContentElementGetNumberOfChars(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetComputedTextLength
//go:noescape
func HasFuncSVGTextContentElementGetComputedTextLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetComputedTextLength
//go:noescape
func FuncSVGTextContentElementGetComputedTextLength(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetComputedTextLength
//go:noescape
func CallSVGTextContentElementGetComputedTextLength(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetComputedTextLength
//go:noescape
func TrySVGTextContentElementGetComputedTextLength(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetSubStringLength
//go:noescape
func HasFuncSVGTextContentElementGetSubStringLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetSubStringLength
//go:noescape
func FuncSVGTextContentElementGetSubStringLength(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetSubStringLength
//go:noescape
func CallSVGTextContentElementGetSubStringLength(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32,
	nchars uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetSubStringLength
//go:noescape
func TrySVGTextContentElementGetSubStringLength(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32,
	nchars uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetStartPositionOfChar
//go:noescape
func HasFuncSVGTextContentElementGetStartPositionOfChar(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetStartPositionOfChar
//go:noescape
func FuncSVGTextContentElementGetStartPositionOfChar(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetStartPositionOfChar
//go:noescape
func CallSVGTextContentElementGetStartPositionOfChar(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetStartPositionOfChar
//go:noescape
func TrySVGTextContentElementGetStartPositionOfChar(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetEndPositionOfChar
//go:noescape
func HasFuncSVGTextContentElementGetEndPositionOfChar(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetEndPositionOfChar
//go:noescape
func FuncSVGTextContentElementGetEndPositionOfChar(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetEndPositionOfChar
//go:noescape
func CallSVGTextContentElementGetEndPositionOfChar(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetEndPositionOfChar
//go:noescape
func TrySVGTextContentElementGetEndPositionOfChar(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetExtentOfChar
//go:noescape
func HasFuncSVGTextContentElementGetExtentOfChar(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetExtentOfChar
//go:noescape
func FuncSVGTextContentElementGetExtentOfChar(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetExtentOfChar
//go:noescape
func CallSVGTextContentElementGetExtentOfChar(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetExtentOfChar
//go:noescape
func TrySVGTextContentElementGetExtentOfChar(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetRotationOfChar
//go:noescape
func HasFuncSVGTextContentElementGetRotationOfChar(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetRotationOfChar
//go:noescape
func FuncSVGTextContentElementGetRotationOfChar(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetRotationOfChar
//go:noescape
func CallSVGTextContentElementGetRotationOfChar(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetRotationOfChar
//go:noescape
func TrySVGTextContentElementGetRotationOfChar(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetCharNumAtPosition
//go:noescape
func HasFuncSVGTextContentElementGetCharNumAtPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetCharNumAtPosition
//go:noescape
func FuncSVGTextContentElementGetCharNumAtPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetCharNumAtPosition
//go:noescape
func CallSVGTextContentElementGetCharNumAtPosition(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetCharNumAtPosition
//go:noescape
func TrySVGTextContentElementGetCharNumAtPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_GetCharNumAtPosition1
//go:noescape
func HasFuncSVGTextContentElementGetCharNumAtPosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_GetCharNumAtPosition1
//go:noescape
func FuncSVGTextContentElementGetCharNumAtPosition1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_GetCharNumAtPosition1
//go:noescape
func CallSVGTextContentElementGetCharNumAtPosition1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTextContentElement_GetCharNumAtPosition1
//go:noescape
func TrySVGTextContentElementGetCharNumAtPosition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTextContentElement_SelectSubString
//go:noescape
func HasFuncSVGTextContentElementSelectSubString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTextContentElement_SelectSubString
//go:noescape
func FuncSVGTextContentElementSelectSubString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTextContentElement_SelectSubString
//go:noescape
func CallSVGTextContentElementSelectSubString(
	this js.Ref, retPtr unsafe.Pointer,
	charnum uint32,
	nchars uint32)

//go:wasmimport plat/js/web try_SVGTextContentElement_SelectSubString
//go:noescape
func TrySVGTextContentElementSelectSubString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	charnum uint32,
	nchars uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPathElement_StartOffset
//go:noescape
func GetSVGTextPathElementStartOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPathElement_Method
//go:noescape
func GetSVGTextPathElementMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPathElement_Spacing
//go:noescape
func GetSVGTextPathElementSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPathElement_Href
//go:noescape
func GetSVGTextPathElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPositioningElement_X
//go:noescape
func GetSVGTextPositioningElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPositioningElement_Y
//go:noescape
func GetSVGTextPositioningElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPositioningElement_Dx
//go:noescape
func GetSVGTextPositioningElementDx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPositioningElement_Dy
//go:noescape
func GetSVGTextPositioningElementDy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTextPositioningElement_Rotate
//go:noescape
func GetSVGTextPositioningElementRotate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGViewElement_ViewBox
//go:noescape
func GetSVGViewElementViewBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGViewElement_PreserveAspectRatio
//go:noescape
func GetSVGViewElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ScoreAdOutput
//go:noescape
func ScoreAdOutputJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScoreAdOutput
//go:noescape
func ScoreAdOutputJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ScoringBrowserSignals
//go:noescape
func ScoringBrowserSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScoringBrowserSignals
//go:noescape
func ScoringBrowserSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ScriptingPolicyReportBody_ViolationType
//go:noescape
func GetScriptingPolicyReportBodyViolationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScriptingPolicyReportBody_ViolationURL
//go:noescape
func GetScriptingPolicyReportBodyViolationURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScriptingPolicyReportBody_ViolationSample
//go:noescape
func GetScriptingPolicyReportBodyViolationSample(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScriptingPolicyReportBody_Lineno
//go:noescape
func GetScriptingPolicyReportBodyLineno(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScriptingPolicyReportBody_Colno
//go:noescape
func GetScriptingPolicyReportBodyColno(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ScriptingPolicyReportBody_ToJSON
//go:noescape
func HasFuncScriptingPolicyReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ScriptingPolicyReportBody_ToJSON
//go:noescape
func FuncScriptingPolicyReportBodyToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ScriptingPolicyReportBody_ToJSON
//go:noescape
func CallScriptingPolicyReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ScriptingPolicyReportBody_ToJSON
//go:noescape
func TryScriptingPolicyReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ScriptingPolicyViolationType
//go:noescape
func ConstOfScriptingPolicyViolationType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ScrollAxis
//go:noescape
func ConstOfScrollAxis(str js.Ref) uint32
