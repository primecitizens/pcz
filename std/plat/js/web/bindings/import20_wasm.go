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

//go:wasmimport plat/js/web new_CSSColor_CSSColor
//go:noescape
func NewCSSColorByCSSColor(
	colorSpace js.Ref,
	channels js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSColor_CSSColor1
//go:noescape
func NewCSSColorByCSSColor1(
	colorSpace js.Ref,
	channels js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSColor_ColorSpace
//go:noescape
func GetCSSColorColorSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSColor_ColorSpace
//go:noescape
func SetCSSColorColorSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSColor_Channels
//go:noescape
func GetCSSColorChannels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSColor_Channels
//go:noescape
func SetCSSColorChannels(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSColor_Alpha
//go:noescape
func GetCSSColorAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSColor_Alpha
//go:noescape
func SetCSSColorAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSColorProfileRule_Name
//go:noescape
func GetCSSColorProfileRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSColorProfileRule_Src
//go:noescape
func GetCSSColorProfileRuleSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSColorProfileRule_RenderingIntent
//go:noescape
func GetCSSColorProfileRuleRenderingIntent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSColorProfileRule_Components
//go:noescape
func GetCSSColorProfileRuleComponents(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSColorValue_Parse
//go:noescape
func HasFuncCSSColorValueParse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSColorValue_Parse
//go:noescape
func FuncCSSColorValueParse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSColorValue_Parse
//go:noescape
func CallCSSColorValueParse(
	this js.Ref, retPtr unsafe.Pointer,
	cssText js.Ref)

//go:wasmimport plat/js/web try_CSSColorValue_Parse
//go:noescape
func TryCSSColorValueParse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cssText js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSConditionRule_ConditionText
//go:noescape
func GetCSSConditionRuleConditionText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSContainerRule_ContainerName
//go:noescape
func GetCSSContainerRuleContainerName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSContainerRule_ContainerQuery
//go:noescape
func GetCSSContainerRuleContainerQuery(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Name
//go:noescape
func GetCSSCounterStyleRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Name
//go:noescape
func SetCSSCounterStyleRuleName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_System
//go:noescape
func GetCSSCounterStyleRuleSystem(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_System
//go:noescape
func SetCSSCounterStyleRuleSystem(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Symbols
//go:noescape
func GetCSSCounterStyleRuleSymbols(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Symbols
//go:noescape
func SetCSSCounterStyleRuleSymbols(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_AdditiveSymbols
//go:noescape
func GetCSSCounterStyleRuleAdditiveSymbols(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_AdditiveSymbols
//go:noescape
func SetCSSCounterStyleRuleAdditiveSymbols(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Negative
//go:noescape
func GetCSSCounterStyleRuleNegative(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Negative
//go:noescape
func SetCSSCounterStyleRuleNegative(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Prefix
//go:noescape
func GetCSSCounterStyleRulePrefix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Prefix
//go:noescape
func SetCSSCounterStyleRulePrefix(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Suffix
//go:noescape
func GetCSSCounterStyleRuleSuffix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Suffix
//go:noescape
func SetCSSCounterStyleRuleSuffix(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Range
//go:noescape
func GetCSSCounterStyleRuleRange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Range
//go:noescape
func SetCSSCounterStyleRuleRange(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Pad
//go:noescape
func GetCSSCounterStyleRulePad(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Pad
//go:noescape
func SetCSSCounterStyleRulePad(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_SpeakAs
//go:noescape
func GetCSSCounterStyleRuleSpeakAs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_SpeakAs
//go:noescape
func SetCSSCounterStyleRuleSpeakAs(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSCounterStyleRule_Fallback
//go:noescape
func GetCSSCounterStyleRuleFallback(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSCounterStyleRule_Fallback
//go:noescape
func SetCSSCounterStyleRuleFallback(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSFontFaceRule_Style
//go:noescape
func GetCSSFontFaceRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSFontFeatureValuesMap_Set
//go:noescape
func HasFuncCSSFontFeatureValuesMapSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSFontFeatureValuesMap_Set
//go:noescape
func FuncCSSFontFeatureValuesMapSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSFontFeatureValuesMap_Set
//go:noescape
func CallCSSFontFeatureValuesMapSet(
	this js.Ref, retPtr unsafe.Pointer,
	featureValueName js.Ref,
	values js.Ref)

//go:wasmimport plat/js/web try_CSSFontFeatureValuesMap_Set
//go:noescape
func TryCSSFontFeatureValuesMapSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	featureValueName js.Ref,
	values js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_FontFamily
//go:noescape
func GetCSSFontFeatureValuesRuleFontFamily(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSFontFeatureValuesRule_FontFamily
//go:noescape
func SetCSSFontFeatureValuesRuleFontFamily(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_Annotation
//go:noescape
func GetCSSFontFeatureValuesRuleAnnotation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_Ornaments
//go:noescape
func GetCSSFontFeatureValuesRuleOrnaments(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_Stylistic
//go:noescape
func GetCSSFontFeatureValuesRuleStylistic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_Swash
//go:noescape
func GetCSSFontFeatureValuesRuleSwash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_CharacterVariant
//go:noescape
func GetCSSFontFeatureValuesRuleCharacterVariant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontFeatureValuesRule_Styleset
//go:noescape
func GetCSSFontFeatureValuesRuleStyleset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontPaletteValuesRule_Name
//go:noescape
func GetCSSFontPaletteValuesRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontPaletteValuesRule_FontFamily
//go:noescape
func GetCSSFontPaletteValuesRuleFontFamily(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontPaletteValuesRule_BasePalette
//go:noescape
func GetCSSFontPaletteValuesRuleBasePalette(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSFontPaletteValuesRule_OverrideColors
//go:noescape
func GetCSSFontPaletteValuesRuleOverrideColors(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSGroupingRule_CssRules
//go:noescape
func GetCSSGroupingRuleCssRules(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSGroupingRule_InsertRule
//go:noescape
func HasFuncCSSGroupingRuleInsertRule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSGroupingRule_InsertRule
//go:noescape
func FuncCSSGroupingRuleInsertRule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSGroupingRule_InsertRule
//go:noescape
func CallCSSGroupingRuleInsertRule(
	this js.Ref, retPtr unsafe.Pointer,
	rule js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_CSSGroupingRule_InsertRule
//go:noescape
func TryCSSGroupingRuleInsertRule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rule js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSGroupingRule_InsertRule1
//go:noescape
func HasFuncCSSGroupingRuleInsertRule1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSGroupingRule_InsertRule1
//go:noescape
func FuncCSSGroupingRuleInsertRule1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSGroupingRule_InsertRule1
//go:noescape
func CallCSSGroupingRuleInsertRule1(
	this js.Ref, retPtr unsafe.Pointer,
	rule js.Ref)

//go:wasmimport plat/js/web try_CSSGroupingRule_InsertRule1
//go:noescape
func TryCSSGroupingRuleInsertRule1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rule js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSGroupingRule_DeleteRule
//go:noescape
func HasFuncCSSGroupingRuleDeleteRule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSGroupingRule_DeleteRule
//go:noescape
func FuncCSSGroupingRuleDeleteRule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSGroupingRule_DeleteRule
//go:noescape
func CallCSSGroupingRuleDeleteRule(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSGroupingRule_DeleteRule
//go:noescape
func TryCSSGroupingRuleDeleteRule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSHSL_CSSHSL
//go:noescape
func NewCSSHSLByCSSHSL(
	h js.Ref,
	s js.Ref,
	l js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSHSL_CSSHSL1
//go:noescape
func NewCSSHSLByCSSHSL1(
	h js.Ref,
	s js.Ref,
	l js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSHSL_H
//go:noescape
func GetCSSHSLH(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHSL_H
//go:noescape
func SetCSSHSLH(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHSL_S
//go:noescape
func GetCSSHSLS(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHSL_S
//go:noescape
func SetCSSHSLS(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHSL_L
//go:noescape
func GetCSSHSLL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHSL_L
//go:noescape
func SetCSSHSLL(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHSL_Alpha
//go:noescape
func GetCSSHSLAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHSL_Alpha
//go:noescape
func SetCSSHSLAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSHWB_CSSHWB
//go:noescape
func NewCSSHWBByCSSHWB(
	h js.Ref,
	w js.Ref,
	b js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSHWB_CSSHWB1
//go:noescape
func NewCSSHWBByCSSHWB1(
	h js.Ref,
	w js.Ref,
	b js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSHWB_H
//go:noescape
func GetCSSHWBH(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHWB_H
//go:noescape
func SetCSSHWBH(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHWB_W
//go:noescape
func GetCSSHWBW(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHWB_W
//go:noescape
func SetCSSHWBW(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHWB_B
//go:noescape
func GetCSSHWBB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHWB_B
//go:noescape
func SetCSSHWBB(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSHWB_Alpha
//go:noescape
func GetCSSHWBAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSHWB_Alpha
//go:noescape
func SetCSSHWBAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSImportRule_Href
//go:noescape
func GetCSSImportRuleHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSImportRule_Media
//go:noescape
func GetCSSImportRuleMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSImportRule_StyleSheet
//go:noescape
func GetCSSImportRuleStyleSheet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSImportRule_LayerName
//go:noescape
func GetCSSImportRuleLayerName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSImportRule_SupportsText
//go:noescape
func GetCSSImportRuleSupportsText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSKeyframeRule_KeyText
//go:noescape
func GetCSSKeyframeRuleKeyText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSKeyframeRule_KeyText
//go:noescape
func SetCSSKeyframeRuleKeyText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSKeyframeRule_Style
//go:noescape
func GetCSSKeyframeRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSKeyframesRule_Name
//go:noescape
func GetCSSKeyframesRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSKeyframesRule_Name
//go:noescape
func SetCSSKeyframesRuleName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSKeyframesRule_CssRules
//go:noescape
func GetCSSKeyframesRuleCssRules(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSKeyframesRule_Length
//go:noescape
func GetCSSKeyframesRuleLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSKeyframesRule_Get
//go:noescape
func HasFuncCSSKeyframesRuleGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSKeyframesRule_Get
//go:noescape
func FuncCSSKeyframesRuleGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSKeyframesRule_Get
//go:noescape
func CallCSSKeyframesRuleGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSKeyframesRule_Get
//go:noescape
func TryCSSKeyframesRuleGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSKeyframesRule_AppendRule
//go:noescape
func HasFuncCSSKeyframesRuleAppendRule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSKeyframesRule_AppendRule
//go:noescape
func FuncCSSKeyframesRuleAppendRule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSKeyframesRule_AppendRule
//go:noescape
func CallCSSKeyframesRuleAppendRule(
	this js.Ref, retPtr unsafe.Pointer,
	rule js.Ref)

//go:wasmimport plat/js/web try_CSSKeyframesRule_AppendRule
//go:noescape
func TryCSSKeyframesRuleAppendRule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rule js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSKeyframesRule_DeleteRule
//go:noescape
func HasFuncCSSKeyframesRuleDeleteRule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSKeyframesRule_DeleteRule
//go:noescape
func FuncCSSKeyframesRuleDeleteRule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSKeyframesRule_DeleteRule
//go:noescape
func CallCSSKeyframesRuleDeleteRule(
	this js.Ref, retPtr unsafe.Pointer,
	sel js.Ref)

//go:wasmimport plat/js/web try_CSSKeyframesRule_DeleteRule
//go:noescape
func TryCSSKeyframesRuleDeleteRule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSKeyframesRule_FindRule
//go:noescape
func HasFuncCSSKeyframesRuleFindRule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSKeyframesRule_FindRule
//go:noescape
func FuncCSSKeyframesRuleFindRule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSKeyframesRule_FindRule
//go:noescape
func CallCSSKeyframesRuleFindRule(
	this js.Ref, retPtr unsafe.Pointer,
	sel js.Ref)

//go:wasmimport plat/js/web try_CSSKeyframesRule_FindRule
//go:noescape
func TryCSSKeyframesRuleFindRule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSLCH_CSSLCH
//go:noescape
func NewCSSLCHByCSSLCH(
	l js.Ref,
	c js.Ref,
	h js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSLCH_CSSLCH1
//go:noescape
func NewCSSLCHByCSSLCH1(
	l js.Ref,
	c js.Ref,
	h js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSLCH_L
//go:noescape
func GetCSSLCHL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLCH_L
//go:noescape
func SetCSSLCHL(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLCH_C
//go:noescape
func GetCSSLCHC(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLCH_C
//go:noescape
func SetCSSLCHC(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLCH_H
//go:noescape
func GetCSSLCHH(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLCH_H
//go:noescape
func SetCSSLCHH(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLCH_Alpha
//go:noescape
func GetCSSLCHAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLCH_Alpha
//go:noescape
func SetCSSLCHAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSLab_CSSLab
//go:noescape
func NewCSSLabByCSSLab(
	l js.Ref,
	a js.Ref,
	b js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSLab_CSSLab1
//go:noescape
func NewCSSLabByCSSLab1(
	l js.Ref,
	a js.Ref,
	b js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSLab_L
//go:noescape
func GetCSSLabL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLab_L
//go:noescape
func SetCSSLabL(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLab_A
//go:noescape
func GetCSSLabA(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLab_A
//go:noescape
func SetCSSLabA(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLab_B
//go:noescape
func GetCSSLabB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLab_B
//go:noescape
func SetCSSLabB(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLab_Alpha
//go:noescape
func GetCSSLabAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSLab_Alpha
//go:noescape
func SetCSSLabAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSLayerBlockRule_Name
//go:noescape
func GetCSSLayerBlockRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSLayerStatementRule_NameList
//go:noescape
func GetCSSLayerStatementRuleNameList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSMarginRule_Name
//go:noescape
func GetCSSMarginRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSMarginRule_Style
//go:noescape
func GetCSSMarginRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathClamp_CSSMathClamp
//go:noescape
func NewCSSMathClampByCSSMathClamp(
	lower js.Ref,
	value js.Ref,
	upper js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSMathClamp_Lower
//go:noescape
func GetCSSMathClampLower(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSMathClamp_Value
//go:noescape
func GetCSSMathClampValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSMathClamp_Upper
//go:noescape
func GetCSSMathClampUpper(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathInvert_CSSMathInvert
//go:noescape
func NewCSSMathInvertByCSSMathInvert(
	arg js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSMathInvert_Value
//go:noescape
func GetCSSMathInvertValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathMax_CSSMathMax
//go:noescape
func NewCSSMathMaxByCSSMathMax(
	argsPtr unsafe.Pointer,
	argsCount uint32) js.Ref

//go:wasmimport plat/js/web get_CSSMathMax_Values
//go:noescape
func GetCSSMathMaxValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathMin_CSSMathMin
//go:noescape
func NewCSSMathMinByCSSMathMin(
	argsPtr unsafe.Pointer,
	argsCount uint32) js.Ref

//go:wasmimport plat/js/web get_CSSMathMin_Values
//go:noescape
func GetCSSMathMinValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSMathNegate_CSSMathNegate
//go:noescape
func NewCSSMathNegateByCSSMathNegate(
	arg js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSMathNegate_Value
//go:noescape
func GetCSSMathNegateValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_CSSMathOperator
//go:noescape
func ConstOfCSSMathOperator(str js.Ref) uint32

//go:wasmimport plat/js/web new_CSSMathProduct_CSSMathProduct
//go:noescape
func NewCSSMathProductByCSSMathProduct(
	argsPtr unsafe.Pointer,
	argsCount uint32) js.Ref

//go:wasmimport plat/js/web get_CSSMathProduct_Values
//go:noescape
func GetCSSMathProductValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSMathValue_Operator
//go:noescape
func GetCSSMathValueOperator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_DOMMatrixReadOnly_DOMMatrixReadOnly
//go:noescape
func NewDOMMatrixReadOnlyByDOMMatrixReadOnly(
	init js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMMatrixReadOnly_DOMMatrixReadOnly1
//go:noescape
func NewDOMMatrixReadOnlyByDOMMatrixReadOnly1() js.Ref

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_A
//go:noescape
func GetDOMMatrixReadOnlyA(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_B
//go:noescape
func GetDOMMatrixReadOnlyB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_C
//go:noescape
func GetDOMMatrixReadOnlyC(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_D
//go:noescape
func GetDOMMatrixReadOnlyD(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_E
//go:noescape
func GetDOMMatrixReadOnlyE(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_F
//go:noescape
func GetDOMMatrixReadOnlyF(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M11
//go:noescape
func GetDOMMatrixReadOnlyM11(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M12
//go:noescape
func GetDOMMatrixReadOnlyM12(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M13
//go:noescape
func GetDOMMatrixReadOnlyM13(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M14
//go:noescape
func GetDOMMatrixReadOnlyM14(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M21
//go:noescape
func GetDOMMatrixReadOnlyM21(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M22
//go:noescape
func GetDOMMatrixReadOnlyM22(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M23
//go:noescape
func GetDOMMatrixReadOnlyM23(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M24
//go:noescape
func GetDOMMatrixReadOnlyM24(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M31
//go:noescape
func GetDOMMatrixReadOnlyM31(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M32
//go:noescape
func GetDOMMatrixReadOnlyM32(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M33
//go:noescape
func GetDOMMatrixReadOnlyM33(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M34
//go:noescape
func GetDOMMatrixReadOnlyM34(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M41
//go:noescape
func GetDOMMatrixReadOnlyM41(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M42
//go:noescape
func GetDOMMatrixReadOnlyM42(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M43
//go:noescape
func GetDOMMatrixReadOnlyM43(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_M44
//go:noescape
func GetDOMMatrixReadOnlyM44(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_Is2D
//go:noescape
func GetDOMMatrixReadOnlyIs2D(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMMatrixReadOnly_IsIdentity
//go:noescape
func GetDOMMatrixReadOnlyIsIdentity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FromMatrix
//go:noescape
func HasFuncDOMMatrixReadOnlyFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FromMatrix
//go:noescape
func FuncDOMMatrixReadOnlyFromMatrix(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FromMatrix
//go:noescape
func CallDOMMatrixReadOnlyFromMatrix(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FromMatrix
//go:noescape
func TryDOMMatrixReadOnlyFromMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FromMatrix1
//go:noescape
func HasFuncDOMMatrixReadOnlyFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FromMatrix1
//go:noescape
func FuncDOMMatrixReadOnlyFromMatrix1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FromMatrix1
//go:noescape
func CallDOMMatrixReadOnlyFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FromMatrix1
//go:noescape
func TryDOMMatrixReadOnlyFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FromFloat32Array
//go:noescape
func HasFuncDOMMatrixReadOnlyFromFloat32Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FromFloat32Array
//go:noescape
func FuncDOMMatrixReadOnlyFromFloat32Array(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FromFloat32Array
//go:noescape
func CallDOMMatrixReadOnlyFromFloat32Array(
	this js.Ref, retPtr unsafe.Pointer,
	array32 js.Ref)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FromFloat32Array
//go:noescape
func TryDOMMatrixReadOnlyFromFloat32Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array32 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FromFloat64Array
//go:noescape
func HasFuncDOMMatrixReadOnlyFromFloat64Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FromFloat64Array
//go:noescape
func FuncDOMMatrixReadOnlyFromFloat64Array(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FromFloat64Array
//go:noescape
func CallDOMMatrixReadOnlyFromFloat64Array(
	this js.Ref, retPtr unsafe.Pointer,
	array64 js.Ref)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FromFloat64Array
//go:noescape
func TryDOMMatrixReadOnlyFromFloat64Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array64 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Translate
//go:noescape
func HasFuncDOMMatrixReadOnlyTranslate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Translate
//go:noescape
func FuncDOMMatrixReadOnlyTranslate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Translate
//go:noescape
func CallDOMMatrixReadOnlyTranslate(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64,
	ty float64,
	tz float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Translate
//go:noescape
func TryDOMMatrixReadOnlyTranslate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64,
	ty float64,
	tz float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Translate1
//go:noescape
func HasFuncDOMMatrixReadOnlyTranslate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Translate1
//go:noescape
func FuncDOMMatrixReadOnlyTranslate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Translate1
//go:noescape
func CallDOMMatrixReadOnlyTranslate1(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64,
	ty float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Translate1
//go:noescape
func TryDOMMatrixReadOnlyTranslate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64,
	ty float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Translate2
//go:noescape
func HasFuncDOMMatrixReadOnlyTranslate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Translate2
//go:noescape
func FuncDOMMatrixReadOnlyTranslate2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Translate2
//go:noescape
func CallDOMMatrixReadOnlyTranslate2(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Translate2
//go:noescape
func TryDOMMatrixReadOnlyTranslate2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Translate3
//go:noescape
func HasFuncDOMMatrixReadOnlyTranslate3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Translate3
//go:noescape
func FuncDOMMatrixReadOnlyTranslate3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Translate3
//go:noescape
func CallDOMMatrixReadOnlyTranslate3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Translate3
//go:noescape
func TryDOMMatrixReadOnlyTranslate3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale
//go:noescape
func HasFuncDOMMatrixReadOnlyScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale
//go:noescape
func FuncDOMMatrixReadOnlyScale(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale
//go:noescape
func CallDOMMatrixReadOnlyScale(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
	originZ float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale
//go:noescape
func TryDOMMatrixReadOnlyScale(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
	originZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale1
//go:noescape
func HasFuncDOMMatrixReadOnlyScale1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale1
//go:noescape
func FuncDOMMatrixReadOnlyScale1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale1
//go:noescape
func CallDOMMatrixReadOnlyScale1(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale1
//go:noescape
func TryDOMMatrixReadOnlyScale1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale2
//go:noescape
func HasFuncDOMMatrixReadOnlyScale2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale2
//go:noescape
func FuncDOMMatrixReadOnlyScale2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale2
//go:noescape
func CallDOMMatrixReadOnlyScale2(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale2
//go:noescape
func TryDOMMatrixReadOnlyScale2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3
//go:noescape
func FuncDOMMatrixReadOnlyScale3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3
//go:noescape
func CallDOMMatrixReadOnlyScale3(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3
//go:noescape
func TryDOMMatrixReadOnlyScale3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale4
//go:noescape
func HasFuncDOMMatrixReadOnlyScale4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale4
//go:noescape
func FuncDOMMatrixReadOnlyScale4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale4
//go:noescape
func CallDOMMatrixReadOnlyScale4(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale4
//go:noescape
func TryDOMMatrixReadOnlyScale4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale5
//go:noescape
func HasFuncDOMMatrixReadOnlyScale5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale5
//go:noescape
func FuncDOMMatrixReadOnlyScale5(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale5
//go:noescape
func CallDOMMatrixReadOnlyScale5(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale5
//go:noescape
func TryDOMMatrixReadOnlyScale5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale6
//go:noescape
func HasFuncDOMMatrixReadOnlyScale6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale6
//go:noescape
func FuncDOMMatrixReadOnlyScale6(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale6
//go:noescape
func CallDOMMatrixReadOnlyScale6(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale6
//go:noescape
func TryDOMMatrixReadOnlyScale6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ScaleNonUniform
//go:noescape
func HasFuncDOMMatrixReadOnlyScaleNonUniform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ScaleNonUniform
//go:noescape
func FuncDOMMatrixReadOnlyScaleNonUniform(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ScaleNonUniform
//go:noescape
func CallDOMMatrixReadOnlyScaleNonUniform(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ScaleNonUniform
//go:noescape
func TryDOMMatrixReadOnlyScaleNonUniform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ScaleNonUniform1
//go:noescape
func HasFuncDOMMatrixReadOnlyScaleNonUniform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ScaleNonUniform1
//go:noescape
func FuncDOMMatrixReadOnlyScaleNonUniform1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ScaleNonUniform1
//go:noescape
func CallDOMMatrixReadOnlyScaleNonUniform1(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ScaleNonUniform1
//go:noescape
func TryDOMMatrixReadOnlyScaleNonUniform1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ScaleNonUniform2
//go:noescape
func HasFuncDOMMatrixReadOnlyScaleNonUniform2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ScaleNonUniform2
//go:noescape
func FuncDOMMatrixReadOnlyScaleNonUniform2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ScaleNonUniform2
//go:noescape
func CallDOMMatrixReadOnlyScaleNonUniform2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ScaleNonUniform2
//go:noescape
func TryDOMMatrixReadOnlyScaleNonUniform2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3d
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3d
//go:noescape
func FuncDOMMatrixReadOnlyScale3d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3d
//go:noescape
func CallDOMMatrixReadOnlyScale3d(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64,
	originZ float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3d
//go:noescape
func TryDOMMatrixReadOnlyScale3d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64,
	originZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3d1
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3d1
//go:noescape
func FuncDOMMatrixReadOnlyScale3d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3d1
//go:noescape
func CallDOMMatrixReadOnlyScale3d1(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3d1
//go:noescape
func TryDOMMatrixReadOnlyScale3d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3d2
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3d2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3d2
//go:noescape
func FuncDOMMatrixReadOnlyScale3d2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3d2
//go:noescape
func CallDOMMatrixReadOnlyScale3d2(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3d2
//go:noescape
func TryDOMMatrixReadOnlyScale3d2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3d3
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3d3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3d3
//go:noescape
func FuncDOMMatrixReadOnlyScale3d3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3d3
//go:noescape
func CallDOMMatrixReadOnlyScale3d3(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3d3
//go:noescape
func TryDOMMatrixReadOnlyScale3d3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Scale3d4
//go:noescape
func HasFuncDOMMatrixReadOnlyScale3d4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Scale3d4
//go:noescape
func FuncDOMMatrixReadOnlyScale3d4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Scale3d4
//go:noescape
func CallDOMMatrixReadOnlyScale3d4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Scale3d4
//go:noescape
func TryDOMMatrixReadOnlyScale3d4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Rotate
//go:noescape
func HasFuncDOMMatrixReadOnlyRotate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Rotate
//go:noescape
func FuncDOMMatrixReadOnlyRotate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Rotate
//go:noescape
func CallDOMMatrixReadOnlyRotate(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64,
	rotY float64,
	rotZ float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Rotate
//go:noescape
func TryDOMMatrixReadOnlyRotate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64,
	rotY float64,
	rotZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Rotate1
//go:noescape
func HasFuncDOMMatrixReadOnlyRotate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Rotate1
//go:noescape
func FuncDOMMatrixReadOnlyRotate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Rotate1
//go:noescape
func CallDOMMatrixReadOnlyRotate1(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64,
	rotY float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Rotate1
//go:noescape
func TryDOMMatrixReadOnlyRotate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64,
	rotY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Rotate2
//go:noescape
func HasFuncDOMMatrixReadOnlyRotate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Rotate2
//go:noescape
func FuncDOMMatrixReadOnlyRotate2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Rotate2
//go:noescape
func CallDOMMatrixReadOnlyRotate2(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Rotate2
//go:noescape
func TryDOMMatrixReadOnlyRotate2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Rotate3
//go:noescape
func HasFuncDOMMatrixReadOnlyRotate3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Rotate3
//go:noescape
func FuncDOMMatrixReadOnlyRotate3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Rotate3
//go:noescape
func CallDOMMatrixReadOnlyRotate3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Rotate3
//go:noescape
func TryDOMMatrixReadOnlyRotate3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateFromVector
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateFromVector(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateFromVector
//go:noescape
func FuncDOMMatrixReadOnlyRotateFromVector(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateFromVector
//go:noescape
func CallDOMMatrixReadOnlyRotateFromVector(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateFromVector
//go:noescape
func TryDOMMatrixReadOnlyRotateFromVector(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateFromVector1
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateFromVector1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateFromVector1
//go:noescape
func FuncDOMMatrixReadOnlyRotateFromVector1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateFromVector1
//go:noescape
func CallDOMMatrixReadOnlyRotateFromVector1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateFromVector1
//go:noescape
func TryDOMMatrixReadOnlyRotateFromVector1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateFromVector2
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateFromVector2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateFromVector2
//go:noescape
func FuncDOMMatrixReadOnlyRotateFromVector2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateFromVector2
//go:noescape
func CallDOMMatrixReadOnlyRotateFromVector2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateFromVector2
//go:noescape
func TryDOMMatrixReadOnlyRotateFromVector2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateAxisAngle
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateAxisAngle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateAxisAngle
//go:noescape
func FuncDOMMatrixReadOnlyRotateAxisAngle(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateAxisAngle
//go:noescape
func CallDOMMatrixReadOnlyRotateAxisAngle(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64,
	angle float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateAxisAngle
//go:noescape
func TryDOMMatrixReadOnlyRotateAxisAngle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64,
	angle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateAxisAngle1
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateAxisAngle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateAxisAngle1
//go:noescape
func FuncDOMMatrixReadOnlyRotateAxisAngle1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateAxisAngle1
//go:noescape
func CallDOMMatrixReadOnlyRotateAxisAngle1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateAxisAngle1
//go:noescape
func TryDOMMatrixReadOnlyRotateAxisAngle1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateAxisAngle2
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateAxisAngle2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateAxisAngle2
//go:noescape
func FuncDOMMatrixReadOnlyRotateAxisAngle2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateAxisAngle2
//go:noescape
func CallDOMMatrixReadOnlyRotateAxisAngle2(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateAxisAngle2
//go:noescape
func TryDOMMatrixReadOnlyRotateAxisAngle2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateAxisAngle3
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateAxisAngle3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateAxisAngle3
//go:noescape
func FuncDOMMatrixReadOnlyRotateAxisAngle3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateAxisAngle3
//go:noescape
func CallDOMMatrixReadOnlyRotateAxisAngle3(
	this js.Ref, retPtr unsafe.Pointer,
	x float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateAxisAngle3
//go:noescape
func TryDOMMatrixReadOnlyRotateAxisAngle3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_RotateAxisAngle4
//go:noescape
func HasFuncDOMMatrixReadOnlyRotateAxisAngle4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_RotateAxisAngle4
//go:noescape
func FuncDOMMatrixReadOnlyRotateAxisAngle4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_RotateAxisAngle4
//go:noescape
func CallDOMMatrixReadOnlyRotateAxisAngle4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_RotateAxisAngle4
//go:noescape
func TryDOMMatrixReadOnlyRotateAxisAngle4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_SkewX
//go:noescape
func HasFuncDOMMatrixReadOnlySkewX(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_SkewX
//go:noescape
func FuncDOMMatrixReadOnlySkewX(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_SkewX
//go:noescape
func CallDOMMatrixReadOnlySkewX(
	this js.Ref, retPtr unsafe.Pointer,
	sx float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_SkewX
//go:noescape
func TryDOMMatrixReadOnlySkewX(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_SkewX1
//go:noescape
func HasFuncDOMMatrixReadOnlySkewX1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_SkewX1
//go:noescape
func FuncDOMMatrixReadOnlySkewX1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_SkewX1
//go:noescape
func CallDOMMatrixReadOnlySkewX1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_SkewX1
//go:noescape
func TryDOMMatrixReadOnlySkewX1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_SkewY
//go:noescape
func HasFuncDOMMatrixReadOnlySkewY(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_SkewY
//go:noescape
func FuncDOMMatrixReadOnlySkewY(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_SkewY
//go:noescape
func CallDOMMatrixReadOnlySkewY(
	this js.Ref, retPtr unsafe.Pointer,
	sy float64)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_SkewY
//go:noescape
func TryDOMMatrixReadOnlySkewY(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sy float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_SkewY1
//go:noescape
func HasFuncDOMMatrixReadOnlySkewY1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_SkewY1
//go:noescape
func FuncDOMMatrixReadOnlySkewY1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_SkewY1
//go:noescape
func CallDOMMatrixReadOnlySkewY1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_SkewY1
//go:noescape
func TryDOMMatrixReadOnlySkewY1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Multiply
//go:noescape
func HasFuncDOMMatrixReadOnlyMultiply(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Multiply
//go:noescape
func FuncDOMMatrixReadOnlyMultiply(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Multiply
//go:noescape
func CallDOMMatrixReadOnlyMultiply(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Multiply
//go:noescape
func TryDOMMatrixReadOnlyMultiply(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Multiply1
//go:noescape
func HasFuncDOMMatrixReadOnlyMultiply1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Multiply1
//go:noescape
func FuncDOMMatrixReadOnlyMultiply1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Multiply1
//go:noescape
func CallDOMMatrixReadOnlyMultiply1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Multiply1
//go:noescape
func TryDOMMatrixReadOnlyMultiply1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FlipX
//go:noescape
func HasFuncDOMMatrixReadOnlyFlipX(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FlipX
//go:noescape
func FuncDOMMatrixReadOnlyFlipX(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FlipX
//go:noescape
func CallDOMMatrixReadOnlyFlipX(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FlipX
//go:noescape
func TryDOMMatrixReadOnlyFlipX(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_FlipY
//go:noescape
func HasFuncDOMMatrixReadOnlyFlipY(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_FlipY
//go:noescape
func FuncDOMMatrixReadOnlyFlipY(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_FlipY
//go:noescape
func CallDOMMatrixReadOnlyFlipY(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_FlipY
//go:noescape
func TryDOMMatrixReadOnlyFlipY(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_Inverse
//go:noescape
func HasFuncDOMMatrixReadOnlyInverse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_Inverse
//go:noescape
func FuncDOMMatrixReadOnlyInverse(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_Inverse
//go:noescape
func CallDOMMatrixReadOnlyInverse(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_Inverse
//go:noescape
func TryDOMMatrixReadOnlyInverse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_TransformPoint
//go:noescape
func HasFuncDOMMatrixReadOnlyTransformPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_TransformPoint
//go:noescape
func FuncDOMMatrixReadOnlyTransformPoint(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_TransformPoint
//go:noescape
func CallDOMMatrixReadOnlyTransformPoint(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_TransformPoint
//go:noescape
func TryDOMMatrixReadOnlyTransformPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_TransformPoint1
//go:noescape
func HasFuncDOMMatrixReadOnlyTransformPoint1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_TransformPoint1
//go:noescape
func FuncDOMMatrixReadOnlyTransformPoint1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_TransformPoint1
//go:noescape
func CallDOMMatrixReadOnlyTransformPoint1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_TransformPoint1
//go:noescape
func TryDOMMatrixReadOnlyTransformPoint1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ToFloat32Array
//go:noescape
func HasFuncDOMMatrixReadOnlyToFloat32Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ToFloat32Array
//go:noescape
func FuncDOMMatrixReadOnlyToFloat32Array(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ToFloat32Array
//go:noescape
func CallDOMMatrixReadOnlyToFloat32Array(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ToFloat32Array
//go:noescape
func TryDOMMatrixReadOnlyToFloat32Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ToFloat64Array
//go:noescape
func HasFuncDOMMatrixReadOnlyToFloat64Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ToFloat64Array
//go:noescape
func FuncDOMMatrixReadOnlyToFloat64Array(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ToFloat64Array
//go:noescape
func CallDOMMatrixReadOnlyToFloat64Array(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ToFloat64Array
//go:noescape
func TryDOMMatrixReadOnlyToFloat64Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ToString
//go:noescape
func HasFuncDOMMatrixReadOnlyToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ToString
//go:noescape
func FuncDOMMatrixReadOnlyToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ToString
//go:noescape
func CallDOMMatrixReadOnlyToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ToString
//go:noescape
func TryDOMMatrixReadOnlyToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrixReadOnly_ToJSON
//go:noescape
func HasFuncDOMMatrixReadOnlyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrixReadOnly_ToJSON
//go:noescape
func FuncDOMMatrixReadOnlyToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMMatrixReadOnly_ToJSON
//go:noescape
func CallDOMMatrixReadOnlyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrixReadOnly_ToJSON
//go:noescape
func TryDOMMatrixReadOnlyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CSSMatrixComponentOptions
//go:noescape
func CSSMatrixComponentOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CSSMatrixComponentOptions
//go:noescape
func CSSMatrixComponentOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSMatrixComponent_CSSMatrixComponent
//go:noescape
func NewCSSMatrixComponentByCSSMatrixComponent(
	matrix js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CSSMatrixComponent_CSSMatrixComponent1
//go:noescape
func NewCSSMatrixComponentByCSSMatrixComponent1(
	matrix js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSMatrixComponent_Matrix
//go:noescape
func GetCSSMatrixComponentMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSMatrixComponent_Matrix
//go:noescape
func SetCSSMatrixComponentMatrix(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSMediaRule_Media
//go:noescape
func GetCSSMediaRuleMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSNamespaceRule_NamespaceURI
//go:noescape
func GetCSSNamespaceRuleNamespaceURI(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSNamespaceRule_Prefix
//go:noescape
func GetCSSNamespaceRulePrefix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSOKLCH_CSSOKLCH
//go:noescape
func NewCSSOKLCHByCSSOKLCH(
	l js.Ref,
	c js.Ref,
	h js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSOKLCH_CSSOKLCH1
//go:noescape
func NewCSSOKLCHByCSSOKLCH1(
	l js.Ref,
	c js.Ref,
	h js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSOKLCH_L
//go:noescape
func GetCSSOKLCHL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLCH_L
//go:noescape
func SetCSSOKLCHL(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLCH_C
//go:noescape
func GetCSSOKLCHC(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLCH_C
//go:noescape
func SetCSSOKLCHC(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLCH_H
//go:noescape
func GetCSSOKLCHH(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLCH_H
//go:noescape
func SetCSSOKLCHH(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLCH_Alpha
//go:noescape
func GetCSSOKLCHAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLCH_Alpha
//go:noescape
func SetCSSOKLCHAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_CSSOKLab_CSSOKLab
//go:noescape
func NewCSSOKLabByCSSOKLab(
	l js.Ref,
	a js.Ref,
	b js.Ref,
	alpha js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSOKLab_CSSOKLab1
//go:noescape
func NewCSSOKLabByCSSOKLab1(
	l js.Ref,
	a js.Ref,
	b js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSOKLab_L
//go:noescape
func GetCSSOKLabL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLab_L
//go:noescape
func SetCSSOKLabL(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLab_A
//go:noescape
func GetCSSOKLabA(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLab_A
//go:noescape
func SetCSSOKLabA(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLab_B
//go:noescape
func GetCSSOKLabB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLab_B
//go:noescape
func SetCSSOKLabB(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSOKLab_Alpha
//go:noescape
func GetCSSOKLabAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSOKLab_Alpha
//go:noescape
func SetCSSOKLabAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSPageRule_SelectorText
//go:noescape
func GetCSSPageRuleSelectorText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSPageRule_SelectorText
//go:noescape
func SetCSSPageRuleSelectorText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSPageRule_Style
//go:noescape
func GetCSSPageRuleStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSParserAtRule_CSSParserAtRule
//go:noescape
func NewCSSParserAtRuleByCSSParserAtRule(
	name js.Ref,
	prelude js.Ref,
	body js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserAtRule_CSSParserAtRule1
//go:noescape
func NewCSSParserAtRuleByCSSParserAtRule1(
	name js.Ref,
	prelude js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserAtRule_Name
//go:noescape
func GetCSSParserAtRuleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserAtRule_Prelude
//go:noescape
func GetCSSParserAtRulePrelude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserAtRule_Body
//go:noescape
func GetCSSParserAtRuleBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSParserAtRule_ToString
//go:noescape
func HasFuncCSSParserAtRuleToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserAtRule_ToString
//go:noescape
func FuncCSSParserAtRuleToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSParserAtRule_ToString
//go:noescape
func CallCSSParserAtRuleToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSParserAtRule_ToString
//go:noescape
func TryCSSParserAtRuleToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSParserBlock_CSSParserBlock
//go:noescape
func NewCSSParserBlockByCSSParserBlock(
	name js.Ref,
	body js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserBlock_Name
//go:noescape
func GetCSSParserBlockName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserBlock_Body
//go:noescape
func GetCSSParserBlockBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSParserBlock_ToString
//go:noescape
func HasFuncCSSParserBlockToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserBlock_ToString
//go:noescape
func FuncCSSParserBlockToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSParserBlock_ToString
//go:noescape
func CallCSSParserBlockToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSParserBlock_ToString
//go:noescape
func TryCSSParserBlockToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSParserFunction_CSSParserFunction
//go:noescape
func NewCSSParserFunctionByCSSParserFunction(
	name js.Ref,
	args js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserFunction_Name
//go:noescape
func GetCSSParserFunctionName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserFunction_Args
//go:noescape
func GetCSSParserFunctionArgs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSParserFunction_ToString
//go:noescape
func HasFuncCSSParserFunctionToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserFunction_ToString
//go:noescape
func FuncCSSParserFunctionToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSParserFunction_ToString
//go:noescape
func CallCSSParserFunctionToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSParserFunction_ToString
//go:noescape
func TryCSSParserFunctionToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_CSSParserQualifiedRule_CSSParserQualifiedRule
//go:noescape
func NewCSSParserQualifiedRuleByCSSParserQualifiedRule(
	prelude js.Ref,
	body js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserQualifiedRule_CSSParserQualifiedRule1
//go:noescape
func NewCSSParserQualifiedRuleByCSSParserQualifiedRule1(
	prelude js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserQualifiedRule_Prelude
//go:noescape
func GetCSSParserQualifiedRulePrelude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserQualifiedRule_Body
//go:noescape
func GetCSSParserQualifiedRuleBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSParserQualifiedRule_ToString
//go:noescape
func HasFuncCSSParserQualifiedRuleToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserQualifiedRule_ToString
//go:noescape
func FuncCSSParserQualifiedRuleToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSParserQualifiedRule_ToString
//go:noescape
func CallCSSParserQualifiedRuleToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSParserQualifiedRule_ToString
//go:noescape
func TryCSSParserQualifiedRuleToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
