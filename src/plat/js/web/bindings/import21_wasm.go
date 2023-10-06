// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web new_CSSPerspective_CSSPerspective
//go:noescape
func NewCSSPerspectiveByCSSPerspective(
	length js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSPerspective_Length
//go:noescape
func GetCSSPerspectiveLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSPerspective_Length
//go:noescape
func SetCSSPerspectiveLength(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSPositionFallbackRule_Name
//go:noescape
func GetCSSPositionFallbackRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSPropertyRule_Name
//go:noescape
func GetCSSPropertyRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSPropertyRule_Syntax
//go:noescape
func GetCSSPropertyRuleSyntax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSPropertyRule_Inherits
//go:noescape
func GetCSSPropertyRuleInherits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSPropertyRule_InitialValue
//go:noescape
func GetCSSPropertyRuleInitialValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSRGB_CSSRGB
//go:noescape
func NewCSSRGBByCSSRGB(
	r js.Ref,
	g js.Ref,
	b js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSRGB_CSSRGB1
//go:noescape
func NewCSSRGBByCSSRGB1(
	r js.Ref,
	g js.Ref,
	b js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSRGB_R
//go:noescape
func GetCSSRGBR(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRGB_R
//go:noescape
func SetCSSRGBR(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRGB_G
//go:noescape
func GetCSSRGBG(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRGB_G
//go:noescape
func SetCSSRGBG(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRGB_B
//go:noescape
func GetCSSRGBB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRGB_B
//go:noescape
func SetCSSRGBB(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRGB_Alpha
//go:noescape
func GetCSSRGBAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRGB_Alpha
//go:noescape
func SetCSSRGBAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSRotate_CSSRotate
//go:noescape
func NewCSSRotateByCSSRotate(
	angle js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSRotate_CSSRotate1
//go:noescape
func NewCSSRotateByCSSRotate1(
	x js.Ref,
	y js.Ref,
	z js.Ref,
	angle js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSRotate_X
//go:noescape
func GetCSSRotateX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRotate_X
//go:noescape
func SetCSSRotateX(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRotate_Y
//go:noescape
func GetCSSRotateY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRotate_Y
//go:noescape
func SetCSSRotateY(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRotate_Z
//go:noescape
func GetCSSRotateZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRotate_Z
//go:noescape
func SetCSSRotateZ(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRotate_Angle
//go:noescape
func GetCSSRotateAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSRotate_Angle
//go:noescape
func SetCSSRotateAngle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSScale_CSSScale
//go:noescape
func NewCSSScaleByCSSScale(
	x js.Ref,
	y js.Ref,
	z js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSScale_CSSScale1
//go:noescape
func NewCSSScaleByCSSScale1(
	x js.Ref,
	y js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSScale_X
//go:noescape
func GetCSSScaleX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSScale_X
//go:noescape
func SetCSSScaleX(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSScale_Y
//go:noescape
func GetCSSScaleY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSScale_Y
//go:noescape
func SetCSSScaleY(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSScale_Z
//go:noescape
func GetCSSScaleZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSScale_Z
//go:noescape
func SetCSSScaleZ(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSScopeRule_Start
//go:noescape
func GetCSSScopeRuleStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSScopeRule_End
//go:noescape
func GetCSSScopeRuleEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSSkew_CSSSkew
//go:noescape
func NewCSSSkewByCSSSkew(
	ax js.Ref,
	ay js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSSkew_Ax
//go:noescape
func GetCSSSkewAx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSSkew_Ax
//go:noescape
func SetCSSSkewAx(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSSkew_Ay
//go:noescape
func GetCSSSkewAy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSSkew_Ay
//go:noescape
func SetCSSSkewAy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSSkewX_CSSSkewX
//go:noescape
func NewCSSSkewXByCSSSkewX(
	ax js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSSkewX_Ax
//go:noescape
func GetCSSSkewXAx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSSkewX_Ax
//go:noescape
func SetCSSSkewXAx(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSSkewY_CSSSkewY
//go:noescape
func NewCSSSkewYByCSSSkewY(
	ay js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSSkewY_Ay
//go:noescape
func GetCSSSkewYAy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSSkewY_Ay
//go:noescape
func SetCSSSkewYAy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleRule_SelectorText
//go:noescape
func GetCSSStyleRuleSelectorText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSStyleRule_SelectorText
//go:noescape
func SetCSSStyleRuleSelectorText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleRule_Style
//go:noescape
func GetCSSStyleRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSStyleRule_StyleMap
//go:noescape
func GetCSSStyleRuleStyleMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSTransformComponent_Is2D
//go:noescape
func GetCSSTransformComponentIs2D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSTransformComponent_Is2D
//go:noescape
func SetCSSTransformComponentIs2D(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_CSSTransformComponent_ToString
//go:noescape
func HasCSSTransformComponentToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSTransformComponent_ToString
//go:noescape
func CSSTransformComponentToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSTransformComponent_ToString
//go:noescape
func CallCSSTransformComponentToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSTransformComponent_ToString
//go:noescape
func TryCSSTransformComponentToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSTransformComponent_ToMatrix
//go:noescape
func HasCSSTransformComponentToMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSTransformComponent_ToMatrix
//go:noescape
func CSSTransformComponentToMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSTransformComponent_ToMatrix
//go:noescape
func CallCSSTransformComponentToMatrix(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSTransformComponent_ToMatrix
//go:noescape
func TryCSSTransformComponentToMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSTransformValue_CSSTransformValue
//go:noescape
func NewCSSTransformValueByCSSTransformValue(
	transforms js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSTransformValue_Length
//go:noescape
func GetCSSTransformValueLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSTransformValue_Is2D
//go:noescape
func GetCSSTransformValueIs2D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSTransformValue_Get
//go:noescape
func HasCSSTransformValueGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSTransformValue_Get
//go:noescape
func CSSTransformValueGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSTransformValue_Get
//go:noescape
func CallCSSTransformValueGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSTransformValue_Get
//go:noescape
func TryCSSTransformValueGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSTransformValue_Set
//go:noescape
func HasCSSTransformValueSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSTransformValue_Set
//go:noescape
func CSSTransformValueSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSTransformValue_Set
//go:noescape
func CallCSSTransformValueSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	val js.Ref)

//go:wasmimport plat/js/web try_CSSTransformValue_Set
//go:noescape
func TryCSSTransformValueSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	val js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSTransformValue_ToMatrix
//go:noescape
func HasCSSTransformValueToMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSTransformValue_ToMatrix
//go:noescape
func CSSTransformValueToMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSTransformValue_ToMatrix
//go:noescape
func CallCSSTransformValueToMatrix(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSTransformValue_ToMatrix
//go:noescape
func TryCSSTransformValueToMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSTransition_CSSTransition
//go:noescape
func NewCSSTransitionByCSSTransition(
	effect js.Ref,
	timeline js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSTransition_CSSTransition1
//go:noescape
func NewCSSTransitionByCSSTransition1(
	effect js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSTransition_CSSTransition2
//go:noescape
func NewCSSTransitionByCSSTransition2() js.Ref

//go:wasmimport plat/js/web get_CSSTransition_TransitionProperty
//go:noescape
func GetCSSTransitionTransitionProperty(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSTranslate_CSSTranslate
//go:noescape
func NewCSSTranslateByCSSTranslate(
	x js.Ref,
	y js.Ref,
	z js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSTranslate_CSSTranslate1
//go:noescape
func NewCSSTranslateByCSSTranslate1(
	x js.Ref,
	y js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSTranslate_X
//go:noescape
func GetCSSTranslateX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSTranslate_X
//go:noescape
func SetCSSTranslateX(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSTranslate_Y
//go:noescape
func GetCSSTranslateY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSTranslate_Y
//go:noescape
func SetCSSTranslateY(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSTranslate_Z
//go:noescape
func GetCSSTranslateZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSTranslate_Z
//go:noescape
func SetCSSTranslateZ(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSTryRule_Style
//go:noescape
func GetCSSTryRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSUnparsedValue_CSSUnparsedValue
//go:noescape
func NewCSSUnparsedValueByCSSUnparsedValue(
	members js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSUnparsedValue_Length
//go:noescape
func GetCSSUnparsedValueLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSUnparsedValue_Get
//go:noescape
func HasCSSUnparsedValueGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSUnparsedValue_Get
//go:noescape
func CSSUnparsedValueGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSUnparsedValue_Get
//go:noescape
func CallCSSUnparsedValueGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSUnparsedValue_Get
//go:noescape
func TryCSSUnparsedValueGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSUnparsedValue_Set
//go:noescape
func HasCSSUnparsedValueSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSUnparsedValue_Set
//go:noescape
func CSSUnparsedValueSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSUnparsedValue_Set
//go:noescape
func CallCSSUnparsedValueSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	val js.Ref)

//go:wasmimport plat/js/web try_CSSUnparsedValue_Set
//go:noescape
func TryCSSUnparsedValueSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	val js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSVariableReferenceValue_CSSVariableReferenceValue
//go:noescape
func NewCSSVariableReferenceValueByCSSVariableReferenceValue(
	variable js.Ref,
	fallback js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSVariableReferenceValue_CSSVariableReferenceValue1
//go:noescape
func NewCSSVariableReferenceValueByCSSVariableReferenceValue1(
	variable js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSVariableReferenceValue_Variable
//go:noescape
func GetCSSVariableReferenceValueVariable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSVariableReferenceValue_Variable
//go:noescape
func SetCSSVariableReferenceValueVariable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSVariableReferenceValue_Fallback
//go:noescape
func GetCSSVariableReferenceValueFallback(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Match
//go:noescape
func HasCacheMatch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Match
//go:noescape
func CacheMatchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Match
//go:noescape
func CallCacheMatch(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_Match
//go:noescape
func TryCacheMatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Match1
//go:noescape
func HasCacheMatch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Match1
//go:noescape
func CacheMatch1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Match1
//go:noescape
func CallCacheMatch1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_Cache_Match1
//go:noescape
func TryCacheMatch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_MatchAll
//go:noescape
func HasCacheMatchAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_MatchAll
//go:noescape
func CacheMatchAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_MatchAll
//go:noescape
func CallCacheMatchAll(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_MatchAll
//go:noescape
func TryCacheMatchAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_MatchAll1
//go:noescape
func HasCacheMatchAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_MatchAll1
//go:noescape
func CacheMatchAll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_MatchAll1
//go:noescape
func CallCacheMatchAll1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_Cache_MatchAll1
//go:noescape
func TryCacheMatchAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_MatchAll2
//go:noescape
func HasCacheMatchAll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_MatchAll2
//go:noescape
func CacheMatchAll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_MatchAll2
//go:noescape
func CallCacheMatchAll2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_MatchAll2
//go:noescape
func TryCacheMatchAll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Add
//go:noescape
func HasCacheAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Add
//go:noescape
func CacheAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Add
//go:noescape
func CallCacheAdd(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_Cache_Add
//go:noescape
func TryCacheAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_AddAll
//go:noescape
func HasCacheAddAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_AddAll
//go:noescape
func CacheAddAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_AddAll
//go:noescape
func CallCacheAddAll(
	this js.Ref, retPtr unsafe.Pointer,
	requests js.Ref)

//go:wasmimport plat/js/web try_Cache_AddAll
//go:noescape
func TryCacheAddAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requests js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Put
//go:noescape
func HasCachePut(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Put
//go:noescape
func CachePutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Put
//go:noescape
func CallCachePut(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	response js.Ref)

//go:wasmimport plat/js/web try_Cache_Put
//go:noescape
func TryCachePut(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	response js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Delete
//go:noescape
func HasCacheDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Delete
//go:noescape
func CacheDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Delete
//go:noescape
func CallCacheDelete(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_Delete
//go:noescape
func TryCacheDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Delete1
//go:noescape
func HasCacheDelete1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Delete1
//go:noescape
func CacheDelete1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Delete1
//go:noescape
func CallCacheDelete1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_Cache_Delete1
//go:noescape
func TryCacheDelete1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Keys
//go:noescape
func HasCacheKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Keys
//go:noescape
func CacheKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Keys
//go:noescape
func CallCacheKeys(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_Keys
//go:noescape
func TryCacheKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Keys1
//go:noescape
func HasCacheKeys1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Keys1
//go:noescape
func CacheKeys1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Keys1
//go:noescape
func CallCacheKeys1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_Cache_Keys1
//go:noescape
func TryCacheKeys1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Cache_Keys2
//go:noescape
func HasCacheKeys2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Cache_Keys2
//go:noescape
func CacheKeys2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Cache_Keys2
//go:noescape
func CallCacheKeys2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Cache_Keys2
//go:noescape
func TryCacheKeys2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MultiCacheQueryOptions
//go:noescape
func MultiCacheQueryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MultiCacheQueryOptions
//go:noescape
func MultiCacheQueryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CacheStorage_Match
//go:noescape
func HasCacheStorageMatch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Match
//go:noescape
func CacheStorageMatchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Match
//go:noescape
func CallCacheStorageMatch(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CacheStorage_Match
//go:noescape
func TryCacheStorageMatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CacheStorage_Match1
//go:noescape
func HasCacheStorageMatch1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Match1
//go:noescape
func CacheStorageMatch1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Match1
//go:noescape
func CallCacheStorageMatch1(
	this js.Ref, retPtr unsafe.Pointer,
	request js.Ref)

//go:wasmimport plat/js/web try_CacheStorage_Match1
//go:noescape
func TryCacheStorageMatch1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CacheStorage_Has
//go:noescape
func HasCacheStorageHas(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Has
//go:noescape
func CacheStorageHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Has
//go:noescape
func CallCacheStorageHas(
	this js.Ref, retPtr unsafe.Pointer,
	cacheName js.Ref)

//go:wasmimport plat/js/web try_CacheStorage_Has
//go:noescape
func TryCacheStorageHas(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cacheName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CacheStorage_Open
//go:noescape
func HasCacheStorageOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Open
//go:noescape
func CacheStorageOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Open
//go:noescape
func CallCacheStorageOpen(
	this js.Ref, retPtr unsafe.Pointer,
	cacheName js.Ref)

//go:wasmimport plat/js/web try_CacheStorage_Open
//go:noescape
func TryCacheStorageOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cacheName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CacheStorage_Delete
//go:noescape
func HasCacheStorageDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Delete
//go:noescape
func CacheStorageDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Delete
//go:noescape
func CallCacheStorageDelete(
	this js.Ref, retPtr unsafe.Pointer,
	cacheName js.Ref)

//go:wasmimport plat/js/web try_CacheStorage_Delete
//go:noescape
func TryCacheStorageDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cacheName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CacheStorage_Keys
//go:noescape
func HasCacheStorageKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CacheStorage_Keys
//go:noescape
func CacheStorageKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CacheStorage_Keys
//go:noescape
func CallCacheStorageKeys(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CacheStorage_Keys
//go:noescape
func TryCacheStorageKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CameraDevicePermissionDescriptor
//go:noescape
func CameraDevicePermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CameraDevicePermissionDescriptor
//go:noescape
func CameraDevicePermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CanMakePaymentEvent_CanMakePaymentEvent
//go:noescape
func NewCanMakePaymentEventByCanMakePaymentEvent(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web has_CanMakePaymentEvent_RespondWith
//go:noescape
func HasCanMakePaymentEventRespondWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanMakePaymentEvent_RespondWith
//go:noescape
func CanMakePaymentEventRespondWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanMakePaymentEvent_RespondWith
//go:noescape
func CallCanMakePaymentEventRespondWith(
	this js.Ref, retPtr unsafe.Pointer,
	canMakePaymentResponse js.Ref)

//go:wasmimport plat/js/web try_CanMakePaymentEvent_RespondWith
//go:noescape
func TryCanMakePaymentEventRespondWith(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	canMakePaymentResponse js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_CanvasCaptureMediaStreamTrack_Canvas
//go:noescape
func GetCanvasCaptureMediaStreamTrackCanvas(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CanvasCaptureMediaStreamTrack_RequestFrame
//go:noescape
func HasCanvasCaptureMediaStreamTrackRequestFrame(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CanvasCaptureMediaStreamTrack_RequestFrame
//go:noescape
func CanvasCaptureMediaStreamTrackRequestFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CanvasCaptureMediaStreamTrack_RequestFrame
//go:noescape
func CallCanvasCaptureMediaStreamTrackRequestFrame(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CanvasCaptureMediaStreamTrack_RequestFrame
//go:noescape
func TryCanvasCaptureMediaStreamTrackRequestFrame(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CaptureActionEventInit
//go:noescape
func CaptureActionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CaptureActionEventInit
//go:noescape
func CaptureActionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CaptureActionEvent_CaptureActionEvent
//go:noescape
func NewCaptureActionEventByCaptureActionEvent(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CaptureActionEvent_CaptureActionEvent1
//go:noescape
func NewCaptureActionEventByCaptureActionEvent1() js.Ref

//go:wasmimport plat/js/web get_CaptureActionEvent_Action
//go:noescape
func GetCaptureActionEventAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_CaptureStartFocusBehavior
//go:noescape
func ConstOfCaptureStartFocusBehavior(str js.Ref) uint32

//go:wasmimport plat/js/web has_CaptureController_SetFocusBehavior
//go:noescape
func HasCaptureControllerSetFocusBehavior(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CaptureController_SetFocusBehavior
//go:noescape
func CaptureControllerSetFocusBehaviorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CaptureController_SetFocusBehavior
//go:noescape
func CallCaptureControllerSetFocusBehavior(
	this js.Ref, retPtr unsafe.Pointer,
	focusBehavior uint32)

//go:wasmimport plat/js/web try_CaptureController_SetFocusBehavior
//go:noescape
func TryCaptureControllerSetFocusBehavior(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	focusBehavior uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_CaptureHandleConfig
//go:noescape
func CaptureHandleConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CaptureHandleConfig
//go:noescape
func CaptureHandleConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CapturedMouseEventInit
//go:noescape
func CapturedMouseEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CapturedMouseEventInit
//go:noescape
func CapturedMouseEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CapturedMouseEvent_CapturedMouseEvent
//go:noescape
func NewCapturedMouseEventByCapturedMouseEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CapturedMouseEvent_CapturedMouseEvent1
//go:noescape
func NewCapturedMouseEventByCapturedMouseEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CapturedMouseEvent_SurfaceX
//go:noescape
func GetCapturedMouseEventSurfaceX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CapturedMouseEvent_SurfaceY
//go:noescape
func GetCapturedMouseEventSurfaceY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CharacterBoundsUpdateEventInit
//go:noescape
func CharacterBoundsUpdateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CharacterBoundsUpdateEventInit
//go:noescape
func CharacterBoundsUpdateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CharacterBoundsUpdateEvent_CharacterBoundsUpdateEvent
//go:noescape
func NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent(
	typ js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CharacterBoundsUpdateEvent_CharacterBoundsUpdateEvent1
//go:noescape
func NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CharacterBoundsUpdateEvent_RangeStart
//go:noescape
func GetCharacterBoundsUpdateEventRangeStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CharacterBoundsUpdateEvent_RangeEnd
//go:noescape
func GetCharacterBoundsUpdateEventRangeEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CharacterData_Data
//go:noescape
func GetCharacterDataData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CharacterData_Data
//go:noescape
func SetCharacterDataData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CharacterData_Length
//go:noescape
func GetCharacterDataLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CharacterData_PreviousElementSibling
//go:noescape
func GetCharacterDataPreviousElementSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CharacterData_NextElementSibling
//go:noescape
func GetCharacterDataNextElementSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_SubstringData
//go:noescape
func HasCharacterDataSubstringData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_SubstringData
//go:noescape
func CharacterDataSubstringDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_SubstringData
//go:noescape
func CallCharacterDataSubstringData(
	this js.Ref, retPtr unsafe.Pointer,
	offset uint32,
	count uint32)

//go:wasmimport plat/js/web try_CharacterData_SubstringData
//go:noescape
func TryCharacterDataSubstringData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset uint32,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_AppendData
//go:noescape
func HasCharacterDataAppendData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_AppendData
//go:noescape
func CharacterDataAppendDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_AppendData
//go:noescape
func CallCharacterDataAppendData(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_CharacterData_AppendData
//go:noescape
func TryCharacterDataAppendData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_InsertData
//go:noescape
func HasCharacterDataInsertData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_InsertData
//go:noescape
func CharacterDataInsertDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_InsertData
//go:noescape
func CallCharacterDataInsertData(
	this js.Ref, retPtr unsafe.Pointer,
	offset uint32,
	data js.Ref)

//go:wasmimport plat/js/web try_CharacterData_InsertData
//go:noescape
func TryCharacterDataInsertData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset uint32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_DeleteData
//go:noescape
func HasCharacterDataDeleteData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_DeleteData
//go:noescape
func CharacterDataDeleteDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_DeleteData
//go:noescape
func CallCharacterDataDeleteData(
	this js.Ref, retPtr unsafe.Pointer,
	offset uint32,
	count uint32)

//go:wasmimport plat/js/web try_CharacterData_DeleteData
//go:noescape
func TryCharacterDataDeleteData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset uint32,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_ReplaceData
//go:noescape
func HasCharacterDataReplaceData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_ReplaceData
//go:noescape
func CharacterDataReplaceDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_ReplaceData
//go:noescape
func CallCharacterDataReplaceData(
	this js.Ref, retPtr unsafe.Pointer,
	offset uint32,
	count uint32,
	data js.Ref)

//go:wasmimport plat/js/web try_CharacterData_ReplaceData
//go:noescape
func TryCharacterDataReplaceData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset uint32,
	count uint32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_Before
//go:noescape
func HasCharacterDataBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_Before
//go:noescape
func CharacterDataBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_Before
//go:noescape
func CallCharacterDataBefore(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_CharacterData_Before
//go:noescape
func TryCharacterDataBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_After
//go:noescape
func HasCharacterDataAfter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_After
//go:noescape
func CharacterDataAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_After
//go:noescape
func CallCharacterDataAfter(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_CharacterData_After
//go:noescape
func TryCharacterDataAfter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_ReplaceWith
//go:noescape
func HasCharacterDataReplaceWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_ReplaceWith
//go:noescape
func CharacterDataReplaceWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_ReplaceWith
//go:noescape
func CallCharacterDataReplaceWith(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_CharacterData_ReplaceWith
//go:noescape
func TryCharacterDataReplaceWith(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CharacterData_Remove
//go:noescape
func HasCharacterDataRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CharacterData_Remove
//go:noescape
func CharacterDataRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CharacterData_Remove
//go:noescape
func CallCharacterDataRemove(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CharacterData_Remove
//go:noescape
func TryCharacterDataRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ChildDisplayType
//go:noescape
func ConstOfChildDisplayType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_FrameType
//go:noescape
func ConstOfFrameType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ClientType
//go:noescape
func ConstOfClientType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ClientLifecycleState
//go:noescape
func ConstOfClientLifecycleState(str js.Ref) uint32

//go:wasmimport plat/js/web get_Client_Url
//go:noescape
func GetClientUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Client_FrameType
//go:noescape
func GetClientFrameType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Client_Id
//go:noescape
func GetClientId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Client_Type
//go:noescape
func GetClientType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Client_LifecycleState
//go:noescape
func GetClientLifecycleState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Client_PostMessage
//go:noescape
func HasClientPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Client_PostMessage
//go:noescape
func ClientPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Client_PostMessage
//go:noescape
func CallClientPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref)

//go:wasmimport plat/js/web try_Client_PostMessage
//go:noescape
func TryClientPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	transfer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Client_PostMessage1
//go:noescape
func HasClientPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Client_PostMessage1
//go:noescape
func ClientPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Client_PostMessage1
//go:noescape
func CallClientPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Client_PostMessage1
//go:noescape
func TryClientPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Client_PostMessage2
//go:noescape
func HasClientPostMessage2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Client_PostMessage2
//go:noescape
func ClientPostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Client_PostMessage2
//go:noescape
func CallClientPostMessage2(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_Client_PostMessage2
//go:noescape
func TryClientPostMessage2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_ClientQueryOptions
//go:noescape
func ClientQueryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ClientQueryOptions
//go:noescape
func ClientQueryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
