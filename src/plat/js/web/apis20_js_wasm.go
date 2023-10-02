// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

func NewCSSColor(colorSpace CSSKeywordish, channels js.Array[CSSColorPercent], alpha CSSNumberish) CSSColor {
	return CSSColor{}.FromRef(
		bindings.NewCSSColorByCSSColor(
			colorSpace.Ref(),
			channels.Ref(),
			alpha.Ref()),
	)
}

func NewCSSColorByCSSColor1(colorSpace CSSKeywordish, channels js.Array[CSSColorPercent]) CSSColor {
	return CSSColor{}.FromRef(
		bindings.NewCSSColorByCSSColor1(
			colorSpace.Ref(),
			channels.Ref()),
	)
}

type CSSColor struct {
	CSSColorValue
}

func (this CSSColor) Once() CSSColor {
	this.Ref().Once()
	return this
}

func (this CSSColor) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSColor) FromRef(ref js.Ref) CSSColor {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSColor) Free() {
	this.Ref().Free()
}

// ColorSpace returns the value of property "CSSColor.colorSpace".
//
// The returned bool will be false if there is no such property.
func (this CSSColor) ColorSpace() (CSSKeywordish, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSKeywordish{}.FromRef(_ret), _ok
}

// ColorSpace sets the value of property "CSSColor.colorSpace" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetColorSpace(val CSSKeywordish) bool {
	return js.True == bindings.SetCSSColorColorSpace(
		this.Ref(),
		val.Ref(),
	)
}

// Channels returns the value of property "CSSColor.channels".
//
// The returned bool will be false if there is no such property.
func (this CSSColor) Channels() (js.ObservableArray[CSSColorPercent], bool) {
	var _ok bool
	_ret := bindings.GetCSSColorChannels(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ObservableArray[CSSColorPercent]{}.FromRef(_ret), _ok
}

// Channels sets the value of property "CSSColor.channels" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetChannels(val js.ObservableArray[CSSColorPercent]) bool {
	return js.True == bindings.SetCSSColorChannels(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSColor.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSColor) Alpha() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSColor.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetAlpha(val CSSNumberish) bool {
	return js.True == bindings.SetCSSColorAlpha(
		this.Ref(),
		val.Ref(),
	)
}

type CSSColorAngle = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue

type CSSColorNumber = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue

type CSSColorProfileRule struct {
	CSSRule
}

func (this CSSColorProfileRule) Once() CSSColorProfileRule {
	this.Ref().Once()
	return this
}

func (this CSSColorProfileRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSColorProfileRule) FromRef(ref js.Ref) CSSColorProfileRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSColorProfileRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSColorProfileRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSColorProfileRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorProfileRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src returns the value of property "CSSColorProfileRule.src".
//
// The returned bool will be false if there is no such property.
func (this CSSColorProfileRule) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorProfileRuleSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// RenderingIntent returns the value of property "CSSColorProfileRule.renderingIntent".
//
// The returned bool will be false if there is no such property.
func (this CSSColorProfileRule) RenderingIntent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorProfileRuleRenderingIntent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Components returns the value of property "CSSColorProfileRule.components".
//
// The returned bool will be false if there is no such property.
func (this CSSColorProfileRule) Components() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSColorProfileRuleComponents(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSColorRGBComp = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue

type OneOf_CSSColorValue_CSSStyleValue struct {
	ref js.Ref
}

func (x OneOf_CSSColorValue_CSSStyleValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CSSColorValue_CSSStyleValue) Free() {
	x.ref.Free()
}

func (x OneOf_CSSColorValue_CSSStyleValue) FromRef(ref js.Ref) OneOf_CSSColorValue_CSSStyleValue {
	return OneOf_CSSColorValue_CSSStyleValue{
		ref: ref,
	}
}

func (x OneOf_CSSColorValue_CSSStyleValue) CSSColorValue() CSSColorValue {
	return CSSColorValue{}.FromRef(x.ref)
}

func (x OneOf_CSSColorValue_CSSStyleValue) CSSStyleValue() CSSStyleValue {
	return CSSStyleValue{}.FromRef(x.ref)
}

type CSSColorValue struct {
	CSSStyleValue
}

func (this CSSColorValue) Once() CSSColorValue {
	this.Ref().Once()
	return this
}

func (this CSSColorValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSColorValue) FromRef(ref js.Ref) CSSColorValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSColorValue) Free() {
	this.Ref().Free()
}

// Parse calls the staticmethod "CSSColorValue.parse".
//
// The returned bool will be false if there is no such method.
func (this CSSColorValue) Parse(cssText js.String) (OneOf_CSSColorValue_CSSStyleValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSColorValueParse(
		this.Ref(), js.Pointer(&_ok),
		cssText.Ref(),
	)

	return OneOf_CSSColorValue_CSSStyleValue{}.FromRef(_ret), _ok
}

// ParseFunc returns the staticmethod "CSSColorValue.parse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSColorValue) ParseFunc() (fn js.Func[func(cssText js.String) OneOf_CSSColorValue_CSSStyleValue]) {
	return fn.FromRef(
		bindings.CSSColorValueParseFunc(
			this.Ref(),
		),
	)
}

type CSSConditionRule struct {
	CSSGroupingRule
}

func (this CSSConditionRule) Once() CSSConditionRule {
	this.Ref().Once()
	return this
}

func (this CSSConditionRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSConditionRule) FromRef(ref js.Ref) CSSConditionRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSConditionRule) Free() {
	this.Ref().Free()
}

// ConditionText returns the value of property "CSSConditionRule.conditionText".
//
// The returned bool will be false if there is no such property.
func (this CSSConditionRule) ConditionText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSConditionRuleConditionText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSContainerRule struct {
	CSSConditionRule
}

func (this CSSContainerRule) Once() CSSContainerRule {
	this.Ref().Once()
	return this
}

func (this CSSContainerRule) Ref() js.Ref {
	return this.CSSConditionRule.Ref()
}

func (this CSSContainerRule) FromRef(ref js.Ref) CSSContainerRule {
	this.CSSConditionRule = this.CSSConditionRule.FromRef(ref)
	return this
}

func (this CSSContainerRule) Free() {
	this.Ref().Free()
}

// ContainerName returns the value of property "CSSContainerRule.containerName".
//
// The returned bool will be false if there is no such property.
func (this CSSContainerRule) ContainerName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSContainerRuleContainerName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContainerQuery returns the value of property "CSSContainerRule.containerQuery".
//
// The returned bool will be false if there is no such property.
func (this CSSContainerRule) ContainerQuery() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSContainerRuleContainerQuery(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSCounterStyleRule struct {
	CSSRule
}

func (this CSSCounterStyleRule) Once() CSSCounterStyleRule {
	this.Ref().Once()
	return this
}

func (this CSSCounterStyleRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSCounterStyleRule) FromRef(ref js.Ref) CSSCounterStyleRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSCounterStyleRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSCounterStyleRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "CSSCounterStyleRule.name" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetName(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleName(
		this.Ref(),
		val.Ref(),
	)
}

// System returns the value of property "CSSCounterStyleRule.system".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) System() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleSystem(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// System sets the value of property "CSSCounterStyleRule.system" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSystem(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSystem(
		this.Ref(),
		val.Ref(),
	)
}

// Symbols returns the value of property "CSSCounterStyleRule.symbols".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Symbols() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleSymbols(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Symbols sets the value of property "CSSCounterStyleRule.symbols" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSymbols(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSymbols(
		this.Ref(),
		val.Ref(),
	)
}

// AdditiveSymbols returns the value of property "CSSCounterStyleRule.additiveSymbols".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) AdditiveSymbols() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleAdditiveSymbols(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AdditiveSymbols sets the value of property "CSSCounterStyleRule.additiveSymbols" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetAdditiveSymbols(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleAdditiveSymbols(
		this.Ref(),
		val.Ref(),
	)
}

// Negative returns the value of property "CSSCounterStyleRule.negative".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Negative() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleNegative(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Negative sets the value of property "CSSCounterStyleRule.negative" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetNegative(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleNegative(
		this.Ref(),
		val.Ref(),
	)
}

// Prefix returns the value of property "CSSCounterStyleRule.prefix".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Prefix() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRulePrefix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Prefix sets the value of property "CSSCounterStyleRule.prefix" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetPrefix(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRulePrefix(
		this.Ref(),
		val.Ref(),
	)
}

// Suffix returns the value of property "CSSCounterStyleRule.suffix".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Suffix() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleSuffix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Suffix sets the value of property "CSSCounterStyleRule.suffix" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSuffix(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSuffix(
		this.Ref(),
		val.Ref(),
	)
}

// Range returns the value of property "CSSCounterStyleRule.range".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Range() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleRange(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Range sets the value of property "CSSCounterStyleRule.range" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetRange(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleRange(
		this.Ref(),
		val.Ref(),
	)
}

// Pad returns the value of property "CSSCounterStyleRule.pad".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Pad() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRulePad(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pad sets the value of property "CSSCounterStyleRule.pad" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetPad(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRulePad(
		this.Ref(),
		val.Ref(),
	)
}

// SpeakAs returns the value of property "CSSCounterStyleRule.speakAs".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) SpeakAs() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleSpeakAs(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SpeakAs sets the value of property "CSSCounterStyleRule.speakAs" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSpeakAs(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSpeakAs(
		this.Ref(),
		val.Ref(),
	)
}

// Fallback returns the value of property "CSSCounterStyleRule.fallback".
//
// The returned bool will be false if there is no such property.
func (this CSSCounterStyleRule) Fallback() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSCounterStyleRuleFallback(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Fallback sets the value of property "CSSCounterStyleRule.fallback" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetFallback(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleFallback(
		this.Ref(),
		val.Ref(),
	)
}

type CSSFontFaceRule struct {
	CSSRule
}

func (this CSSFontFaceRule) Once() CSSFontFaceRule {
	this.Ref().Once()
	return this
}

func (this CSSFontFaceRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSFontFaceRule) FromRef(ref js.Ref) CSSFontFaceRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSFontFaceRule) Free() {
	this.Ref().Free()
}

// Style returns the value of property "CSSFontFaceRule.style".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFaceRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFaceRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

type OneOf_Uint32_ArrayUint32 struct {
	ref js.Ref
}

func (x OneOf_Uint32_ArrayUint32) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Uint32_ArrayUint32) Free() {
	x.ref.Free()
}

func (x OneOf_Uint32_ArrayUint32) FromRef(ref js.Ref) OneOf_Uint32_ArrayUint32 {
	return OneOf_Uint32_ArrayUint32{
		ref: ref,
	}
}

func (x OneOf_Uint32_ArrayUint32) Uint32() uint32 {
	return js.Number[uint32]{}.FromRef(x.ref).Get()
}

func (x OneOf_Uint32_ArrayUint32) ArrayUint32() js.Array[uint32] {
	return js.Array[uint32]{}.FromRef(x.ref)
}

type CSSFontFeatureValuesMap struct {
	ref js.Ref
}

func (this CSSFontFeatureValuesMap) Once() CSSFontFeatureValuesMap {
	this.Ref().Once()
	return this
}

func (this CSSFontFeatureValuesMap) Ref() js.Ref {
	return this.ref
}

func (this CSSFontFeatureValuesMap) FromRef(ref js.Ref) CSSFontFeatureValuesMap {
	this.ref = ref
	return this
}

func (this CSSFontFeatureValuesMap) Free() {
	this.Ref().Free()
}

// Set calls the method "CSSFontFeatureValuesMap.set".
//
// The returned bool will be false if there is no such method.
func (this CSSFontFeatureValuesMap) Set(featureValueName js.String, values OneOf_Uint32_ArrayUint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSFontFeatureValuesMapSet(
		this.Ref(), js.Pointer(&_ok),
		featureValueName.Ref(),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "CSSFontFeatureValuesMap.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSFontFeatureValuesMap) SetFunc() (fn js.Func[func(featureValueName js.String, values OneOf_Uint32_ArrayUint32)]) {
	return fn.FromRef(
		bindings.CSSFontFeatureValuesMapSetFunc(
			this.Ref(),
		),
	)
}

type CSSFontFeatureValuesRule struct {
	CSSRule
}

func (this CSSFontFeatureValuesRule) Once() CSSFontFeatureValuesRule {
	this.Ref().Once()
	return this
}

func (this CSSFontFeatureValuesRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSFontFeatureValuesRule) FromRef(ref js.Ref) CSSFontFeatureValuesRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSFontFeatureValuesRule) Free() {
	this.Ref().Free()
}

// FontFamily returns the value of property "CSSFontFeatureValuesRule.fontFamily".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) FontFamily() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleFontFamily(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FontFamily sets the value of property "CSSFontFeatureValuesRule.fontFamily" to val.
//
// It returns false if the property cannot be set.
func (this CSSFontFeatureValuesRule) SetFontFamily(val js.String) bool {
	return js.True == bindings.SetCSSFontFeatureValuesRuleFontFamily(
		this.Ref(),
		val.Ref(),
	)
}

// Annotation returns the value of property "CSSFontFeatureValuesRule.annotation".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) Annotation() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleAnnotation(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

// Ornaments returns the value of property "CSSFontFeatureValuesRule.ornaments".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) Ornaments() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleOrnaments(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

// Stylistic returns the value of property "CSSFontFeatureValuesRule.stylistic".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) Stylistic() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleStylistic(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

// Swash returns the value of property "CSSFontFeatureValuesRule.swash".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) Swash() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleSwash(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

// CharacterVariant returns the value of property "CSSFontFeatureValuesRule.characterVariant".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) CharacterVariant() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleCharacterVariant(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

// Styleset returns the value of property "CSSFontFeatureValuesRule.styleset".
//
// The returned bool will be false if there is no such property.
func (this CSSFontFeatureValuesRule) Styleset() (CSSFontFeatureValuesMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontFeatureValuesRuleStyleset(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSFontFeatureValuesMap{}.FromRef(_ret), _ok
}

type CSSFontPaletteValuesRule struct {
	CSSRule
}

func (this CSSFontPaletteValuesRule) Once() CSSFontPaletteValuesRule {
	this.Ref().Once()
	return this
}

func (this CSSFontPaletteValuesRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSFontPaletteValuesRule) FromRef(ref js.Ref) CSSFontPaletteValuesRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSFontPaletteValuesRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSFontPaletteValuesRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSFontPaletteValuesRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontPaletteValuesRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FontFamily returns the value of property "CSSFontPaletteValuesRule.fontFamily".
//
// The returned bool will be false if there is no such property.
func (this CSSFontPaletteValuesRule) FontFamily() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontPaletteValuesRuleFontFamily(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BasePalette returns the value of property "CSSFontPaletteValuesRule.basePalette".
//
// The returned bool will be false if there is no such property.
func (this CSSFontPaletteValuesRule) BasePalette() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontPaletteValuesRuleBasePalette(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OverrideColors returns the value of property "CSSFontPaletteValuesRule.overrideColors".
//
// The returned bool will be false if there is no such property.
func (this CSSFontPaletteValuesRule) OverrideColors() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSFontPaletteValuesRuleOverrideColors(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSGroupingRule struct {
	CSSRule
}

func (this CSSGroupingRule) Once() CSSGroupingRule {
	this.Ref().Once()
	return this
}

func (this CSSGroupingRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSGroupingRule) FromRef(ref js.Ref) CSSGroupingRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSGroupingRule) Free() {
	this.Ref().Free()
}

// CssRules returns the value of property "CSSGroupingRule.cssRules".
//
// The returned bool will be false if there is no such property.
func (this CSSGroupingRule) CssRules() (CSSRuleList, bool) {
	var _ok bool
	_ret := bindings.GetCSSGroupingRuleCssRules(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRuleList{}.FromRef(_ret), _ok
}

// InsertRule calls the method "CSSGroupingRule.insertRule".
//
// The returned bool will be false if there is no such method.
func (this CSSGroupingRule) InsertRule(rule js.String, index uint32) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallCSSGroupingRuleInsertRule(
		this.Ref(), js.Pointer(&_ok),
		rule.Ref(),
		uint32(index),
	)

	return uint32(_ret), _ok
}

// InsertRuleFunc returns the method "CSSGroupingRule.insertRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSGroupingRule) InsertRuleFunc() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleInsertRuleFunc(
			this.Ref(),
		),
	)
}

// InsertRule1 calls the method "CSSGroupingRule.insertRule".
//
// The returned bool will be false if there is no such method.
func (this CSSGroupingRule) InsertRule1(rule js.String) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallCSSGroupingRuleInsertRule1(
		this.Ref(), js.Pointer(&_ok),
		rule.Ref(),
	)

	return uint32(_ret), _ok
}

// InsertRule1Func returns the method "CSSGroupingRule.insertRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSGroupingRule) InsertRule1Func() (fn js.Func[func(rule js.String) uint32]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleInsertRule1Func(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSGroupingRule.deleteRule".
//
// The returned bool will be false if there is no such method.
func (this CSSGroupingRule) DeleteRule(index uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSGroupingRuleDeleteRule(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRuleFunc returns the method "CSSGroupingRule.deleteRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSGroupingRule) DeleteRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleDeleteRuleFunc(
			this.Ref(),
		),
	)
}

func NewCSSHSL(h CSSColorAngle, s CSSColorPercent, l CSSColorPercent, alpha CSSColorPercent) CSSHSL {
	return CSSHSL{}.FromRef(
		bindings.NewCSSHSLByCSSHSL(
			h.Ref(),
			s.Ref(),
			l.Ref(),
			alpha.Ref()),
	)
}

func NewCSSHSLByCSSHSL1(h CSSColorAngle, s CSSColorPercent, l CSSColorPercent) CSSHSL {
	return CSSHSL{}.FromRef(
		bindings.NewCSSHSLByCSSHSL1(
			h.Ref(),
			s.Ref(),
			l.Ref()),
	)
}

type CSSHSL struct {
	CSSColorValue
}

func (this CSSHSL) Once() CSSHSL {
	this.Ref().Once()
	return this
}

func (this CSSHSL) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSHSL) FromRef(ref js.Ref) CSSHSL {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSHSL) Free() {
	this.Ref().Free()
}

// H returns the value of property "CSSHSL.h".
//
// The returned bool will be false if there is no such property.
func (this CSSHSL) H() (CSSColorAngle, bool) {
	var _ok bool
	_ret := bindings.GetCSSHSLH(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorAngle{}.FromRef(_ret), _ok
}

// H sets the value of property "CSSHSL.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSHSLH(
		this.Ref(),
		val.Ref(),
	)
}

// S returns the value of property "CSSHSL.s".
//
// The returned bool will be false if there is no such property.
func (this CSSHSL) S() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSHSLS(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// S sets the value of property "CSSHSL.s" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetS(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLS(
		this.Ref(),
		val.Ref(),
	)
}

// L returns the value of property "CSSHSL.l".
//
// The returned bool will be false if there is no such property.
func (this CSSHSL) L() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSHSLL(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// L sets the value of property "CSSHSL.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLL(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSHSL.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSHSL) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSHSLAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSHSL.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSHWB(h CSSNumericValue, w CSSNumberish, b CSSNumberish, alpha CSSNumberish) CSSHWB {
	return CSSHWB{}.FromRef(
		bindings.NewCSSHWBByCSSHWB(
			h.Ref(),
			w.Ref(),
			b.Ref(),
			alpha.Ref()),
	)
}

func NewCSSHWBByCSSHWB1(h CSSNumericValue, w CSSNumberish, b CSSNumberish) CSSHWB {
	return CSSHWB{}.FromRef(
		bindings.NewCSSHWBByCSSHWB1(
			h.Ref(),
			w.Ref(),
			b.Ref()),
	)
}

type CSSHWB struct {
	CSSColorValue
}

func (this CSSHWB) Once() CSSHWB {
	this.Ref().Once()
	return this
}

func (this CSSHWB) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSHWB) FromRef(ref js.Ref) CSSHWB {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSHWB) Free() {
	this.Ref().Free()
}

// H returns the value of property "CSSHWB.h".
//
// The returned bool will be false if there is no such property.
func (this CSSHWB) H() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSHWBH(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// H sets the value of property "CSSHWB.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetH(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSHWBH(
		this.Ref(),
		val.Ref(),
	)
}

// W returns the value of property "CSSHWB.w".
//
// The returned bool will be false if there is no such property.
func (this CSSHWB) W() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSHWBW(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// W sets the value of property "CSSHWB.w" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetW(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBW(
		this.Ref(),
		val.Ref(),
	)
}

// B returns the value of property "CSSHWB.b".
//
// The returned bool will be false if there is no such property.
func (this CSSHWB) B() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSHWBB(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// B sets the value of property "CSSHWB.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetB(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBB(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSHWB.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSHWB) Alpha() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSHWBAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSHWB.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetAlpha(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBAlpha(
		this.Ref(),
		val.Ref(),
	)
}

type CSSImageValue struct {
	CSSStyleValue
}

func (this CSSImageValue) Once() CSSImageValue {
	this.Ref().Once()
	return this
}

func (this CSSImageValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSImageValue) FromRef(ref js.Ref) CSSImageValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSImageValue) Free() {
	this.Ref().Free()
}

type CSSImportRule struct {
	CSSRule
}

func (this CSSImportRule) Once() CSSImportRule {
	this.Ref().Once()
	return this
}

func (this CSSImportRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSImportRule) FromRef(ref js.Ref) CSSImportRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSImportRule) Free() {
	this.Ref().Free()
}

// Href returns the value of property "CSSImportRule.href".
//
// The returned bool will be false if there is no such property.
func (this CSSImportRule) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSImportRuleHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media returns the value of property "CSSImportRule.media".
//
// The returned bool will be false if there is no such property.
func (this CSSImportRule) Media() (MediaList, bool) {
	var _ok bool
	_ret := bindings.GetCSSImportRuleMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaList{}.FromRef(_ret), _ok
}

// StyleSheet returns the value of property "CSSImportRule.styleSheet".
//
// The returned bool will be false if there is no such property.
func (this CSSImportRule) StyleSheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetCSSImportRuleStyleSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

// LayerName returns the value of property "CSSImportRule.layerName".
//
// The returned bool will be false if there is no such property.
func (this CSSImportRule) LayerName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSImportRuleLayerName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SupportsText returns the value of property "CSSImportRule.supportsText".
//
// The returned bool will be false if there is no such property.
func (this CSSImportRule) SupportsText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSImportRuleSupportsText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSKeyframeRule struct {
	CSSRule
}

func (this CSSKeyframeRule) Once() CSSKeyframeRule {
	this.Ref().Once()
	return this
}

func (this CSSKeyframeRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSKeyframeRule) FromRef(ref js.Ref) CSSKeyframeRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSKeyframeRule) Free() {
	this.Ref().Free()
}

// KeyText returns the value of property "CSSKeyframeRule.keyText".
//
// The returned bool will be false if there is no such property.
func (this CSSKeyframeRule) KeyText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeyframeRuleKeyText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// KeyText sets the value of property "CSSKeyframeRule.keyText" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeyframeRule) SetKeyText(val js.String) bool {
	return js.True == bindings.SetCSSKeyframeRuleKeyText(
		this.Ref(),
		val.Ref(),
	)
}

// Style returns the value of property "CSSKeyframeRule.style".
//
// The returned bool will be false if there is no such property.
func (this CSSKeyframeRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeyframeRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

type CSSKeyframesRule struct {
	CSSRule
}

func (this CSSKeyframesRule) Once() CSSKeyframesRule {
	this.Ref().Once()
	return this
}

func (this CSSKeyframesRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSKeyframesRule) FromRef(ref js.Ref) CSSKeyframesRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSKeyframesRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSKeyframesRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSKeyframesRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeyframesRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "CSSKeyframesRule.name" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeyframesRule) SetName(val js.String) bool {
	return js.True == bindings.SetCSSKeyframesRuleName(
		this.Ref(),
		val.Ref(),
	)
}

// CssRules returns the value of property "CSSKeyframesRule.cssRules".
//
// The returned bool will be false if there is no such property.
func (this CSSKeyframesRule) CssRules() (CSSRuleList, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeyframesRuleCssRules(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSRuleList{}.FromRef(_ret), _ok
}

// Length returns the value of property "CSSKeyframesRule.length".
//
// The returned bool will be false if there is no such property.
func (this CSSKeyframesRule) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSKeyframesRuleLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "CSSKeyframesRule.".
//
// The returned bool will be false if there is no such method.
func (this CSSKeyframesRule) Get(index uint32) (CSSKeyframeRule, bool) {
	var _ok bool
	_ret := bindings.CallCSSKeyframesRuleGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSKeyframeRule{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CSSKeyframesRule.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSKeyframesRule) GetFunc() (fn js.Func[func(index uint32) CSSKeyframeRule]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleGetFunc(
			this.Ref(),
		),
	)
}

// AppendRule calls the method "CSSKeyframesRule.appendRule".
//
// The returned bool will be false if there is no such method.
func (this CSSKeyframesRule) AppendRule(rule js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSKeyframesRuleAppendRule(
		this.Ref(), js.Pointer(&_ok),
		rule.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendRuleFunc returns the method "CSSKeyframesRule.appendRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSKeyframesRule) AppendRuleFunc() (fn js.Func[func(rule js.String)]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleAppendRuleFunc(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSKeyframesRule.deleteRule".
//
// The returned bool will be false if there is no such method.
func (this CSSKeyframesRule) DeleteRule(sel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSKeyframesRuleDeleteRule(
		this.Ref(), js.Pointer(&_ok),
		sel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRuleFunc returns the method "CSSKeyframesRule.deleteRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSKeyframesRule) DeleteRuleFunc() (fn js.Func[func(sel js.String)]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleDeleteRuleFunc(
			this.Ref(),
		),
	)
}

// FindRule calls the method "CSSKeyframesRule.findRule".
//
// The returned bool will be false if there is no such method.
func (this CSSKeyframesRule) FindRule(sel js.String) (CSSKeyframeRule, bool) {
	var _ok bool
	_ret := bindings.CallCSSKeyframesRuleFindRule(
		this.Ref(), js.Pointer(&_ok),
		sel.Ref(),
	)

	return CSSKeyframeRule{}.FromRef(_ret), _ok
}

// FindRuleFunc returns the method "CSSKeyframesRule.findRule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSKeyframesRule) FindRuleFunc() (fn js.Func[func(sel js.String) CSSKeyframeRule]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleFindRuleFunc(
			this.Ref(),
		),
	)
}

func NewCSSLCH(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle, alpha CSSColorPercent) CSSLCH {
	return CSSLCH{}.FromRef(
		bindings.NewCSSLCHByCSSLCH(
			l.Ref(),
			c.Ref(),
			h.Ref(),
			alpha.Ref()),
	)
}

func NewCSSLCHByCSSLCH1(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle) CSSLCH {
	return CSSLCH{}.FromRef(
		bindings.NewCSSLCHByCSSLCH1(
			l.Ref(),
			c.Ref(),
			h.Ref()),
	)
}

type CSSLCH struct {
	CSSColorValue
}

func (this CSSLCH) Once() CSSLCH {
	this.Ref().Once()
	return this
}

func (this CSSLCH) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSLCH) FromRef(ref js.Ref) CSSLCH {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSLCH) Free() {
	this.Ref().Free()
}

// L returns the value of property "CSSLCH.l".
//
// The returned bool will be false if there is no such property.
func (this CSSLCH) L() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSLCHL(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// L sets the value of property "CSSLCH.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHL(
		this.Ref(),
		val.Ref(),
	)
}

// C returns the value of property "CSSLCH.c".
//
// The returned bool will be false if there is no such property.
func (this CSSLCH) C() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSLCHC(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// C sets the value of property "CSSLCH.c" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetC(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHC(
		this.Ref(),
		val.Ref(),
	)
}

// H returns the value of property "CSSLCH.h".
//
// The returned bool will be false if there is no such property.
func (this CSSLCH) H() (CSSColorAngle, bool) {
	var _ok bool
	_ret := bindings.GetCSSLCHH(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorAngle{}.FromRef(_ret), _ok
}

// H sets the value of property "CSSLCH.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSLCHH(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSLCH.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSLCH) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSLCHAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSLab(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber, alpha CSSColorPercent) CSSLab {
	return CSSLab{}.FromRef(
		bindings.NewCSSLabByCSSLab(
			l.Ref(),
			a.Ref(),
			b.Ref(),
			alpha.Ref()),
	)
}

func NewCSSLabByCSSLab1(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber) CSSLab {
	return CSSLab{}.FromRef(
		bindings.NewCSSLabByCSSLab1(
			l.Ref(),
			a.Ref(),
			b.Ref()),
	)
}

type CSSLab struct {
	CSSColorValue
}

func (this CSSLab) Once() CSSLab {
	this.Ref().Once()
	return this
}

func (this CSSLab) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSLab) FromRef(ref js.Ref) CSSLab {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSLab) Free() {
	this.Ref().Free()
}

// L returns the value of property "CSSLab.l".
//
// The returned bool will be false if there is no such property.
func (this CSSLab) L() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSLabL(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// L sets the value of property "CSSLab.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLabL(
		this.Ref(),
		val.Ref(),
	)
}

// A returns the value of property "CSSLab.a".
//
// The returned bool will be false if there is no such property.
func (this CSSLab) A() (CSSColorNumber, bool) {
	var _ok bool
	_ret := bindings.GetCSSLabA(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorNumber{}.FromRef(_ret), _ok
}

// A sets the value of property "CSSLab.a" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetA(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSLabA(
		this.Ref(),
		val.Ref(),
	)
}

// B returns the value of property "CSSLab.b".
//
// The returned bool will be false if there is no such property.
func (this CSSLab) B() (CSSColorNumber, bool) {
	var _ok bool
	_ret := bindings.GetCSSLabB(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorNumber{}.FromRef(_ret), _ok
}

// B sets the value of property "CSSLab.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetB(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSLabB(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSLab.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSLab) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSLabAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSLab.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLabAlpha(
		this.Ref(),
		val.Ref(),
	)
}

type CSSLayerBlockRule struct {
	CSSGroupingRule
}

func (this CSSLayerBlockRule) Once() CSSLayerBlockRule {
	this.Ref().Once()
	return this
}

func (this CSSLayerBlockRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSLayerBlockRule) FromRef(ref js.Ref) CSSLayerBlockRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSLayerBlockRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSLayerBlockRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSLayerBlockRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSLayerBlockRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSLayerStatementRule struct {
	CSSRule
}

func (this CSSLayerStatementRule) Once() CSSLayerStatementRule {
	this.Ref().Once()
	return this
}

func (this CSSLayerStatementRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSLayerStatementRule) FromRef(ref js.Ref) CSSLayerStatementRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSLayerStatementRule) Free() {
	this.Ref().Free()
}

// NameList returns the value of property "CSSLayerStatementRule.nameList".
//
// The returned bool will be false if there is no such property.
func (this CSSLayerStatementRule) NameList() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetCSSLayerStatementRuleNameList(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

type CSSMarginRule struct {
	CSSRule
}

func (this CSSMarginRule) Once() CSSMarginRule {
	this.Ref().Once()
	return this
}

func (this CSSMarginRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSMarginRule) FromRef(ref js.Ref) CSSMarginRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSMarginRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSMarginRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSMarginRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSMarginRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Style returns the value of property "CSSMarginRule.style".
//
// The returned bool will be false if there is no such property.
func (this CSSMarginRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSMarginRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

func NewCSSMathClamp(lower CSSNumberish, value CSSNumberish, upper CSSNumberish) CSSMathClamp {
	return CSSMathClamp{}.FromRef(
		bindings.NewCSSMathClampByCSSMathClamp(
			lower.Ref(),
			value.Ref(),
			upper.Ref()),
	)
}

type CSSMathClamp struct {
	CSSMathValue
}

func (this CSSMathClamp) Once() CSSMathClamp {
	this.Ref().Once()
	return this
}

func (this CSSMathClamp) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathClamp) FromRef(ref js.Ref) CSSMathClamp {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathClamp) Free() {
	this.Ref().Free()
}

// Lower returns the value of property "CSSMathClamp.lower".
//
// The returned bool will be false if there is no such property.
func (this CSSMathClamp) Lower() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathClampLower(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Value returns the value of property "CSSMathClamp.value".
//
// The returned bool will be false if there is no such property.
func (this CSSMathClamp) Value() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathClampValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Upper returns the value of property "CSSMathClamp.upper".
//
// The returned bool will be false if there is no such property.
func (this CSSMathClamp) Upper() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathClampUpper(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

func NewCSSMathInvert(arg CSSNumberish) CSSMathInvert {
	return CSSMathInvert{}.FromRef(
		bindings.NewCSSMathInvertByCSSMathInvert(
			arg.Ref()),
	)
}

type CSSMathInvert struct {
	CSSMathValue
}

func (this CSSMathInvert) Once() CSSMathInvert {
	this.Ref().Once()
	return this
}

func (this CSSMathInvert) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathInvert) FromRef(ref js.Ref) CSSMathInvert {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathInvert) Free() {
	this.Ref().Free()
}

// Value returns the value of property "CSSMathInvert.value".
//
// The returned bool will be false if there is no such property.
func (this CSSMathInvert) Value() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathInvertValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

func NewCSSMathMax(args ...CSSNumberish) CSSMathMax {
	return CSSMathMax{}.FromRef(
		bindings.NewCSSMathMaxByCSSMathMax(
			js.SliceData(args),
			js.SizeU(len(args))),
	)
}

type CSSMathMax struct {
	CSSMathValue
}

func (this CSSMathMax) Once() CSSMathMax {
	this.Ref().Once()
	return this
}

func (this CSSMathMax) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathMax) FromRef(ref js.Ref) CSSMathMax {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathMax) Free() {
	this.Ref().Free()
}

// Values returns the value of property "CSSMathMax.values".
//
// The returned bool will be false if there is no such property.
func (this CSSMathMax) Values() (CSSNumericArray, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathMaxValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericArray{}.FromRef(_ret), _ok
}

func NewCSSMathMin(args ...CSSNumberish) CSSMathMin {
	return CSSMathMin{}.FromRef(
		bindings.NewCSSMathMinByCSSMathMin(
			js.SliceData(args),
			js.SizeU(len(args))),
	)
}

type CSSMathMin struct {
	CSSMathValue
}

func (this CSSMathMin) Once() CSSMathMin {
	this.Ref().Once()
	return this
}

func (this CSSMathMin) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathMin) FromRef(ref js.Ref) CSSMathMin {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathMin) Free() {
	this.Ref().Free()
}

// Values returns the value of property "CSSMathMin.values".
//
// The returned bool will be false if there is no such property.
func (this CSSMathMin) Values() (CSSNumericArray, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathMinValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericArray{}.FromRef(_ret), _ok
}

func NewCSSMathNegate(arg CSSNumberish) CSSMathNegate {
	return CSSMathNegate{}.FromRef(
		bindings.NewCSSMathNegateByCSSMathNegate(
			arg.Ref()),
	)
}

type CSSMathNegate struct {
	CSSMathValue
}

func (this CSSMathNegate) Once() CSSMathNegate {
	this.Ref().Once()
	return this
}

func (this CSSMathNegate) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathNegate) FromRef(ref js.Ref) CSSMathNegate {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathNegate) Free() {
	this.Ref().Free()
}

// Value returns the value of property "CSSMathNegate.value".
//
// The returned bool will be false if there is no such property.
func (this CSSMathNegate) Value() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathNegateValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

type CSSMathOperator uint32

const (
	_ CSSMathOperator = iota

	CSSMathOperator_SUM
	CSSMathOperator_PRODUCT
	CSSMathOperator_NEGATE
	CSSMathOperator_INVERT
	CSSMathOperator_MIN
	CSSMathOperator_MAX
	CSSMathOperator_CLAMP
)

func (CSSMathOperator) FromRef(str js.Ref) CSSMathOperator {
	return CSSMathOperator(bindings.ConstOfCSSMathOperator(str))
}

func (x CSSMathOperator) String() (string, bool) {
	switch x {
	case CSSMathOperator_SUM:
		return "sum", true
	case CSSMathOperator_PRODUCT:
		return "product", true
	case CSSMathOperator_NEGATE:
		return "negate", true
	case CSSMathOperator_INVERT:
		return "invert", true
	case CSSMathOperator_MIN:
		return "min", true
	case CSSMathOperator_MAX:
		return "max", true
	case CSSMathOperator_CLAMP:
		return "clamp", true
	default:
		return "", false
	}
}

func NewCSSMathProduct(args ...CSSNumberish) CSSMathProduct {
	return CSSMathProduct{}.FromRef(
		bindings.NewCSSMathProductByCSSMathProduct(
			js.SliceData(args),
			js.SizeU(len(args))),
	)
}

type CSSMathProduct struct {
	CSSMathValue
}

func (this CSSMathProduct) Once() CSSMathProduct {
	this.Ref().Once()
	return this
}

func (this CSSMathProduct) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathProduct) FromRef(ref js.Ref) CSSMathProduct {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathProduct) Free() {
	this.Ref().Free()
}

// Values returns the value of property "CSSMathProduct.values".
//
// The returned bool will be false if there is no such property.
func (this CSSMathProduct) Values() (CSSNumericArray, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathProductValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericArray{}.FromRef(_ret), _ok
}

type CSSMathValue struct {
	CSSNumericValue
}

func (this CSSMathValue) Once() CSSMathValue {
	this.Ref().Once()
	return this
}

func (this CSSMathValue) Ref() js.Ref {
	return this.CSSNumericValue.Ref()
}

func (this CSSMathValue) FromRef(ref js.Ref) CSSMathValue {
	this.CSSNumericValue = this.CSSNumericValue.FromRef(ref)
	return this
}

func (this CSSMathValue) Free() {
	this.Ref().Free()
}

// Operator returns the value of property "CSSMathValue.operator".
//
// The returned bool will be false if there is no such property.
func (this CSSMathValue) Operator() (CSSMathOperator, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathValueOperator(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSMathOperator(_ret), _ok
}

func NewDOMMatrixReadOnly(init OneOf_String_ArrayFloat64) DOMMatrixReadOnly {
	return DOMMatrixReadOnly{}.FromRef(
		bindings.NewDOMMatrixReadOnlyByDOMMatrixReadOnly(
			init.Ref()),
	)
}

func NewDOMMatrixReadOnlyByDOMMatrixReadOnly1() DOMMatrixReadOnly {
	return DOMMatrixReadOnly{}.FromRef(
		bindings.NewDOMMatrixReadOnlyByDOMMatrixReadOnly1(),
	)
}

type DOMMatrixReadOnly struct {
	ref js.Ref
}

func (this DOMMatrixReadOnly) Once() DOMMatrixReadOnly {
	this.Ref().Once()
	return this
}

func (this DOMMatrixReadOnly) Ref() js.Ref {
	return this.ref
}

func (this DOMMatrixReadOnly) FromRef(ref js.Ref) DOMMatrixReadOnly {
	this.ref = ref
	return this
}

func (this DOMMatrixReadOnly) Free() {
	this.Ref().Free()
}

// A returns the value of property "DOMMatrixReadOnly.a".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) A() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyA(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// B returns the value of property "DOMMatrixReadOnly.b".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) B() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyB(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// C returns the value of property "DOMMatrixReadOnly.c".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) C() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyC(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// D returns the value of property "DOMMatrixReadOnly.d".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) D() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyD(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// E returns the value of property "DOMMatrixReadOnly.e".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) E() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyE(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// F returns the value of property "DOMMatrixReadOnly.f".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) F() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyF(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M11 returns the value of property "DOMMatrixReadOnly.m11".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M11() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM11(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M12 returns the value of property "DOMMatrixReadOnly.m12".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M12() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM12(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M13 returns the value of property "DOMMatrixReadOnly.m13".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M13() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM13(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M14 returns the value of property "DOMMatrixReadOnly.m14".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M14() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM14(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M21 returns the value of property "DOMMatrixReadOnly.m21".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M21() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM21(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M22 returns the value of property "DOMMatrixReadOnly.m22".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M22() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM22(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M23 returns the value of property "DOMMatrixReadOnly.m23".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M23() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM23(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M24 returns the value of property "DOMMatrixReadOnly.m24".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M24() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM24(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M31 returns the value of property "DOMMatrixReadOnly.m31".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M31() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM31(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M32 returns the value of property "DOMMatrixReadOnly.m32".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M32() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM32(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M33 returns the value of property "DOMMatrixReadOnly.m33".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M33() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM33(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M34 returns the value of property "DOMMatrixReadOnly.m34".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M34() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM34(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M41 returns the value of property "DOMMatrixReadOnly.m41".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M41() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM41(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M42 returns the value of property "DOMMatrixReadOnly.m42".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M42() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM42(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M43 returns the value of property "DOMMatrixReadOnly.m43".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M43() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM43(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// M44 returns the value of property "DOMMatrixReadOnly.m44".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) M44() (float64, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyM44(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Is2D returns the value of property "DOMMatrixReadOnly.is2D".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) Is2D() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyIs2D(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsIdentity returns the value of property "DOMMatrixReadOnly.isIdentity".
//
// The returned bool will be false if there is no such property.
func (this DOMMatrixReadOnly) IsIdentity() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDOMMatrixReadOnlyIsIdentity(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FromMatrix calls the staticmethod "DOMMatrixReadOnly.fromMatrix".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FromMatrix(other DOMMatrixInit) (DOMMatrixReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFromMatrix(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMMatrixReadOnly{}.FromRef(_ret), _ok
}

// FromMatrixFunc returns the staticmethod "DOMMatrixReadOnly.fromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FromMatrixFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromMatrixFunc(
			this.Ref(),
		),
	)
}

// FromMatrix1 calls the staticmethod "DOMMatrixReadOnly.fromMatrix".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FromMatrix1() (DOMMatrixReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFromMatrix1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrixReadOnly{}.FromRef(_ret), _ok
}

// FromMatrix1Func returns the staticmethod "DOMMatrixReadOnly.fromMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FromMatrix1Func() (fn js.Func[func() DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromMatrix1Func(
			this.Ref(),
		),
	)
}

// FromFloat32Array calls the staticmethod "DOMMatrixReadOnly.fromFloat32Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FromFloat32Array(array32 js.TypedArray[float32]) (DOMMatrixReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFromFloat32Array(
		this.Ref(), js.Pointer(&_ok),
		array32.Ref(),
	)

	return DOMMatrixReadOnly{}.FromRef(_ret), _ok
}

// FromFloat32ArrayFunc returns the staticmethod "DOMMatrixReadOnly.fromFloat32Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FromFloat32ArrayFunc() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat64Array calls the staticmethod "DOMMatrixReadOnly.fromFloat64Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FromFloat64Array(array64 js.TypedArray[float64]) (DOMMatrixReadOnly, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFromFloat64Array(
		this.Ref(), js.Pointer(&_ok),
		array64.Ref(),
	)

	return DOMMatrixReadOnly{}.FromRef(_ret), _ok
}

// FromFloat64ArrayFunc returns the staticmethod "DOMMatrixReadOnly.fromFloat64Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FromFloat64ArrayFunc() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "DOMMatrixReadOnly.translate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Translate(tx float64, ty float64, tz float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTranslate(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// TranslateFunc returns the method "DOMMatrixReadOnly.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) TranslateFunc() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslateFunc(
			this.Ref(),
		),
	)
}

// Translate1 calls the method "DOMMatrixReadOnly.translate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Translate1(tx float64, ty float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTranslate1(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
		float64(ty),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Translate1Func returns the method "DOMMatrixReadOnly.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Translate1Func() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate1Func(
			this.Ref(),
		),
	)
}

// Translate2 calls the method "DOMMatrixReadOnly.translate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Translate2(tx float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTranslate2(
		this.Ref(), js.Pointer(&_ok),
		float64(tx),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Translate2Func returns the method "DOMMatrixReadOnly.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Translate2Func() (fn js.Func[func(tx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate2Func(
			this.Ref(),
		),
	)
}

// Translate3 calls the method "DOMMatrixReadOnly.translate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Translate3() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTranslate3(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Translate3Func returns the method "DOMMatrixReadOnly.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Translate3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate3Func(
			this.Ref(),
		),
	)
}

// Scale calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleFunc returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ScaleFunc() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleFunc(
			this.Ref(),
		),
	)
}

// Scale1 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale1(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale1Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale1Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale1Func(
			this.Ref(),
		),
	)
}

// Scale2 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale2(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale2Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale2Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale2Func(
			this.Ref(),
		),
	)
}

// Scale3 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3(scaleX float64, scaleY float64, scaleZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3Func(
			this.Ref(),
		),
	)
}

// Scale4 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale4(scaleX float64, scaleY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale4(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale4Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale4Func() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale4Func(
			this.Ref(),
		),
	)
}

// Scale5 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale5(scaleX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale5(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale5Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale5Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale5Func(
			this.Ref(),
		),
	)
}

// Scale6 calls the method "DOMMatrixReadOnly.scale".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale6() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale6(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale6Func returns the method "DOMMatrixReadOnly.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale6Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale6Func(
			this.Ref(),
		),
	)
}

// ScaleNonUniform calls the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniform(scaleX float64, scaleY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScaleNonUniform(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
		float64(scaleY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleNonUniformFunc returns the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniformFunc() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniformFunc(
			this.Ref(),
		),
	)
}

// ScaleNonUniform1 calls the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniform1(scaleX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScaleNonUniform1(
		this.Ref(), js.Pointer(&_ok),
		float64(scaleX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleNonUniform1Func returns the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniform1Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniform1Func(
			this.Ref(),
		),
	)
}

// ScaleNonUniform2 calls the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniform2() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScaleNonUniform2(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ScaleNonUniform2Func returns the method "DOMMatrixReadOnly.scaleNonUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ScaleNonUniform2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniform2Func(
			this.Ref(),
		),
	)
}

// Scale3d calls the method "DOMMatrixReadOnly.scale3d".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3d(scale float64, originX float64, originY float64, originZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3d(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3dFunc returns the method "DOMMatrixReadOnly.scale3d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3dFunc() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3dFunc(
			this.Ref(),
		),
	)
}

// Scale3d1 calls the method "DOMMatrixReadOnly.scale3d".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3d1(scale float64, originX float64, originY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3d1(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3d1Func returns the method "DOMMatrixReadOnly.scale3d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3d1Func() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d1Func(
			this.Ref(),
		),
	)
}

// Scale3d2 calls the method "DOMMatrixReadOnly.scale3d".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3d2(scale float64, originX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3d2(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
		float64(originX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3d2Func returns the method "DOMMatrixReadOnly.scale3d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3d2Func() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d2Func(
			this.Ref(),
		),
	)
}

// Scale3d3 calls the method "DOMMatrixReadOnly.scale3d".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3d3(scale float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3d3(
		this.Ref(), js.Pointer(&_ok),
		float64(scale),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3d3Func returns the method "DOMMatrixReadOnly.scale3d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3d3Func() (fn js.Func[func(scale float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d3Func(
			this.Ref(),
		),
	)
}

// Scale3d4 calls the method "DOMMatrixReadOnly.scale3d".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Scale3d4() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyScale3d4(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Scale3d4Func returns the method "DOMMatrixReadOnly.scale3d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Scale3d4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d4Func(
			this.Ref(),
		),
	)
}

// Rotate calls the method "DOMMatrixReadOnly.rotate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Rotate(rotX float64, rotY float64, rotZ float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotate(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFunc returns the method "DOMMatrixReadOnly.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateFunc() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFunc(
			this.Ref(),
		),
	)
}

// Rotate1 calls the method "DOMMatrixReadOnly.rotate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Rotate1(rotX float64, rotY float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotate1(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
		float64(rotY),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Rotate1Func returns the method "DOMMatrixReadOnly.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Rotate1Func() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate1Func(
			this.Ref(),
		),
	)
}

// Rotate2 calls the method "DOMMatrixReadOnly.rotate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Rotate2(rotX float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotate2(
		this.Ref(), js.Pointer(&_ok),
		float64(rotX),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Rotate2Func returns the method "DOMMatrixReadOnly.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Rotate2Func() (fn js.Func[func(rotX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate2Func(
			this.Ref(),
		),
	)
}

// Rotate3 calls the method "DOMMatrixReadOnly.rotate".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Rotate3() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotate3(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Rotate3Func returns the method "DOMMatrixReadOnly.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Rotate3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate3Func(
			this.Ref(),
		),
	)
}

// RotateFromVector calls the method "DOMMatrixReadOnly.rotateFromVector".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVector(x float64, y float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateFromVector(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVectorFunc returns the method "DOMMatrixReadOnly.rotateFromVector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVectorFunc() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVectorFunc(
			this.Ref(),
		),
	)
}

// RotateFromVector1 calls the method "DOMMatrixReadOnly.rotateFromVector".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVector1(x float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateFromVector1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVector1Func returns the method "DOMMatrixReadOnly.rotateFromVector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVector1Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVector1Func(
			this.Ref(),
		),
	)
}

// RotateFromVector2 calls the method "DOMMatrixReadOnly.rotateFromVector".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVector2() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateFromVector2(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateFromVector2Func returns the method "DOMMatrixReadOnly.rotateFromVector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateFromVector2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVector2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle calls the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle(x float64, y float64, z float64, angle float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateAxisAngle(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngleFunc returns the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngleFunc() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngleFunc(
			this.Ref(),
		),
	)
}

// RotateAxisAngle1 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle1(x float64, y float64, z float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateAxisAngle1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(z),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngle1Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle1Func() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle1Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle2 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle2(x float64, y float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateAxisAngle2(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngle2Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle2Func() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle3 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle3(x float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateAxisAngle3(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngle3Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle3Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle3Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle4 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle4() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyRotateAxisAngle4(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// RotateAxisAngle4Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) RotateAxisAngle4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle4Func(
			this.Ref(),
		),
	)
}

// SkewX calls the method "DOMMatrixReadOnly.skewX".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) SkewX(sx float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlySkewX(
		this.Ref(), js.Pointer(&_ok),
		float64(sx),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewXFunc returns the method "DOMMatrixReadOnly.skewX".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) SkewXFunc() (fn js.Func[func(sx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewXFunc(
			this.Ref(),
		),
	)
}

// SkewX1 calls the method "DOMMatrixReadOnly.skewX".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) SkewX1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlySkewX1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewX1Func returns the method "DOMMatrixReadOnly.skewX".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) SkewX1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewX1Func(
			this.Ref(),
		),
	)
}

// SkewY calls the method "DOMMatrixReadOnly.skewY".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) SkewY(sy float64) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlySkewY(
		this.Ref(), js.Pointer(&_ok),
		float64(sy),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewYFunc returns the method "DOMMatrixReadOnly.skewY".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) SkewYFunc() (fn js.Func[func(sy float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewYFunc(
			this.Ref(),
		),
	)
}

// SkewY1 calls the method "DOMMatrixReadOnly.skewY".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) SkewY1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlySkewY1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// SkewY1Func returns the method "DOMMatrixReadOnly.skewY".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) SkewY1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewY1Func(
			this.Ref(),
		),
	)
}

// Multiply calls the method "DOMMatrixReadOnly.multiply".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Multiply(other DOMMatrixInit) (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyMultiply(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&other),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// MultiplyFunc returns the method "DOMMatrixReadOnly.multiply".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) MultiplyFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyMultiplyFunc(
			this.Ref(),
		),
	)
}

// Multiply1 calls the method "DOMMatrixReadOnly.multiply".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Multiply1() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyMultiply1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// Multiply1Func returns the method "DOMMatrixReadOnly.multiply".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) Multiply1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyMultiply1Func(
			this.Ref(),
		),
	)
}

// FlipX calls the method "DOMMatrixReadOnly.flipX".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FlipX() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFlipX(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FlipXFunc returns the method "DOMMatrixReadOnly.flipX".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FlipXFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFlipXFunc(
			this.Ref(),
		),
	)
}

// FlipY calls the method "DOMMatrixReadOnly.flipY".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) FlipY() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyFlipY(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// FlipYFunc returns the method "DOMMatrixReadOnly.flipY".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) FlipYFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFlipYFunc(
			this.Ref(),
		),
	)
}

// Inverse calls the method "DOMMatrixReadOnly.inverse".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) Inverse() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyInverse(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// InverseFunc returns the method "DOMMatrixReadOnly.inverse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) InverseFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyInverseFunc(
			this.Ref(),
		),
	)
}

// TransformPoint calls the method "DOMMatrixReadOnly.transformPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) TransformPoint(point DOMPointInit) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTransformPoint(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// TransformPointFunc returns the method "DOMMatrixReadOnly.transformPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) TransformPointFunc() (fn js.Func[func(point DOMPointInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTransformPointFunc(
			this.Ref(),
		),
	)
}

// TransformPoint1 calls the method "DOMMatrixReadOnly.transformPoint".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) TransformPoint1() (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyTransformPoint1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// TransformPoint1Func returns the method "DOMMatrixReadOnly.transformPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) TransformPoint1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTransformPoint1Func(
			this.Ref(),
		),
	)
}

// ToFloat32Array calls the method "DOMMatrixReadOnly.toFloat32Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ToFloat32Array() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyToFloat32Array(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// ToFloat32ArrayFunc returns the method "DOMMatrixReadOnly.toFloat32Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ToFloat32ArrayFunc() (fn js.Func[func() js.TypedArray[float32]]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// ToFloat64Array calls the method "DOMMatrixReadOnly.toFloat64Array".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ToFloat64Array() (js.TypedArray[float64], bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyToFloat64Array(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.TypedArray[float64]{}.FromRef(_ret), _ok
}

// ToFloat64ArrayFunc returns the method "DOMMatrixReadOnly.toFloat64Array".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ToFloat64ArrayFunc() (fn js.Func[func() js.TypedArray[float64]]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "DOMMatrixReadOnly.toString".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "DOMMatrixReadOnly.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToStringFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMMatrixReadOnly.toJSON".
//
// The returned bool will be false if there is no such method.
func (this DOMMatrixReadOnly) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDOMMatrixReadOnlyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "DOMMatrixReadOnly.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DOMMatrixReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
}

type CSSMatrixComponentOptions struct {
	// Is2D is "CSSMatrixComponentOptions.is2D"
	//
	// Optional
	//
	// NOTE: FFI_USE_Is2D MUST be set to true to make this field effective.
	Is2D bool

	FFI_USE_Is2D bool // for Is2D.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CSSMatrixComponentOptions with all fields set.
func (p CSSMatrixComponentOptions) FromRef(ref js.Ref) CSSMatrixComponentOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CSSMatrixComponentOptions in the application heap.
func (p CSSMatrixComponentOptions) New() js.Ref {
	return bindings.CSSMatrixComponentOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CSSMatrixComponentOptions) UpdateFrom(ref js.Ref) {
	bindings.CSSMatrixComponentOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CSSMatrixComponentOptions) Update(ref js.Ref) {
	bindings.CSSMatrixComponentOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCSSMatrixComponent(matrix DOMMatrixReadOnly, options CSSMatrixComponentOptions) CSSMatrixComponent {
	return CSSMatrixComponent{}.FromRef(
		bindings.NewCSSMatrixComponentByCSSMatrixComponent(
			matrix.Ref(),
			js.Pointer(&options)),
	)
}

func NewCSSMatrixComponentByCSSMatrixComponent1(matrix DOMMatrixReadOnly) CSSMatrixComponent {
	return CSSMatrixComponent{}.FromRef(
		bindings.NewCSSMatrixComponentByCSSMatrixComponent1(
			matrix.Ref()),
	)
}

type CSSMatrixComponent struct {
	CSSTransformComponent
}

func (this CSSMatrixComponent) Once() CSSMatrixComponent {
	this.Ref().Once()
	return this
}

func (this CSSMatrixComponent) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSMatrixComponent) FromRef(ref js.Ref) CSSMatrixComponent {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSMatrixComponent) Free() {
	this.Ref().Free()
}

// Matrix returns the value of property "CSSMatrixComponent.matrix".
//
// The returned bool will be false if there is no such property.
func (this CSSMatrixComponent) Matrix() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.GetCSSMatrixComponentMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMMatrix{}.FromRef(_ret), _ok
}

// Matrix sets the value of property "CSSMatrixComponent.matrix" to val.
//
// It returns false if the property cannot be set.
func (this CSSMatrixComponent) SetMatrix(val DOMMatrix) bool {
	return js.True == bindings.SetCSSMatrixComponentMatrix(
		this.Ref(),
		val.Ref(),
	)
}

type CSSMediaRule struct {
	CSSConditionRule
}

func (this CSSMediaRule) Once() CSSMediaRule {
	this.Ref().Once()
	return this
}

func (this CSSMediaRule) Ref() js.Ref {
	return this.CSSConditionRule.Ref()
}

func (this CSSMediaRule) FromRef(ref js.Ref) CSSMediaRule {
	this.CSSConditionRule = this.CSSConditionRule.FromRef(ref)
	return this
}

func (this CSSMediaRule) Free() {
	this.Ref().Free()
}

// Media returns the value of property "CSSMediaRule.media".
//
// The returned bool will be false if there is no such property.
func (this CSSMediaRule) Media() (MediaList, bool) {
	var _ok bool
	_ret := bindings.GetCSSMediaRuleMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaList{}.FromRef(_ret), _ok
}

type CSSNamespaceRule struct {
	CSSRule
}

func (this CSSNamespaceRule) Once() CSSNamespaceRule {
	this.Ref().Once()
	return this
}

func (this CSSNamespaceRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSNamespaceRule) FromRef(ref js.Ref) CSSNamespaceRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSNamespaceRule) Free() {
	this.Ref().Free()
}

// NamespaceURI returns the value of property "CSSNamespaceRule.namespaceURI".
//
// The returned bool will be false if there is no such property.
func (this CSSNamespaceRule) NamespaceURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSNamespaceRuleNamespaceURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Prefix returns the value of property "CSSNamespaceRule.prefix".
//
// The returned bool will be false if there is no such property.
func (this CSSNamespaceRule) Prefix() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSNamespaceRulePrefix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

func NewCSSOKLCH(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle, alpha CSSColorPercent) CSSOKLCH {
	return CSSOKLCH{}.FromRef(
		bindings.NewCSSOKLCHByCSSOKLCH(
			l.Ref(),
			c.Ref(),
			h.Ref(),
			alpha.Ref()),
	)
}

func NewCSSOKLCHByCSSOKLCH1(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle) CSSOKLCH {
	return CSSOKLCH{}.FromRef(
		bindings.NewCSSOKLCHByCSSOKLCH1(
			l.Ref(),
			c.Ref(),
			h.Ref()),
	)
}

type CSSOKLCH struct {
	CSSColorValue
}

func (this CSSOKLCH) Once() CSSOKLCH {
	this.Ref().Once()
	return this
}

func (this CSSOKLCH) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSOKLCH) FromRef(ref js.Ref) CSSOKLCH {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSOKLCH) Free() {
	this.Ref().Free()
}

// L returns the value of property "CSSOKLCH.l".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLCH) L() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLCHL(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// L sets the value of property "CSSOKLCH.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHL(
		this.Ref(),
		val.Ref(),
	)
}

// C returns the value of property "CSSOKLCH.c".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLCH) C() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLCHC(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// C sets the value of property "CSSOKLCH.c" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetC(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHC(
		this.Ref(),
		val.Ref(),
	)
}

// H returns the value of property "CSSOKLCH.h".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLCH) H() (CSSColorAngle, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLCHH(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorAngle{}.FromRef(_ret), _ok
}

// H sets the value of property "CSSOKLCH.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSOKLCHH(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSOKLCH.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLCH) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLCHAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSOKLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSOKLab(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber, alpha CSSColorPercent) CSSOKLab {
	return CSSOKLab{}.FromRef(
		bindings.NewCSSOKLabByCSSOKLab(
			l.Ref(),
			a.Ref(),
			b.Ref(),
			alpha.Ref()),
	)
}

func NewCSSOKLabByCSSOKLab1(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber) CSSOKLab {
	return CSSOKLab{}.FromRef(
		bindings.NewCSSOKLabByCSSOKLab1(
			l.Ref(),
			a.Ref(),
			b.Ref()),
	)
}

type CSSOKLab struct {
	CSSColorValue
}

func (this CSSOKLab) Once() CSSOKLab {
	this.Ref().Once()
	return this
}

func (this CSSOKLab) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSOKLab) FromRef(ref js.Ref) CSSOKLab {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSOKLab) Free() {
	this.Ref().Free()
}

// L returns the value of property "CSSOKLab.l".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLab) L() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLabL(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// L sets the value of property "CSSOKLab.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLabL(
		this.Ref(),
		val.Ref(),
	)
}

// A returns the value of property "CSSOKLab.a".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLab) A() (CSSColorNumber, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLabA(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorNumber{}.FromRef(_ret), _ok
}

// A sets the value of property "CSSOKLab.a" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetA(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSOKLabA(
		this.Ref(),
		val.Ref(),
	)
}

// B returns the value of property "CSSOKLab.b".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLab) B() (CSSColorNumber, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLabB(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorNumber{}.FromRef(_ret), _ok
}

// B sets the value of property "CSSOKLab.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetB(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSOKLabB(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSOKLab.alpha".
//
// The returned bool will be false if there is no such property.
func (this CSSOKLab) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSOKLabAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSOKLab.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLabAlpha(
		this.Ref(),
		val.Ref(),
	)
}

type CSSPageRule struct {
	CSSGroupingRule
}

func (this CSSPageRule) Once() CSSPageRule {
	this.Ref().Once()
	return this
}

func (this CSSPageRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSPageRule) FromRef(ref js.Ref) CSSPageRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSPageRule) Free() {
	this.Ref().Free()
}

// SelectorText returns the value of property "CSSPageRule.selectorText".
//
// The returned bool will be false if there is no such property.
func (this CSSPageRule) SelectorText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPageRuleSelectorText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectorText sets the value of property "CSSPageRule.selectorText" to val.
//
// It returns false if the property cannot be set.
func (this CSSPageRule) SetSelectorText(val js.String) bool {
	return js.True == bindings.SetCSSPageRuleSelectorText(
		this.Ref(),
		val.Ref(),
	)
}

// Style returns the value of property "CSSPageRule.style".
//
// The returned bool will be false if there is no such property.
func (this CSSPageRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSPageRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

func NewCSSParserAtRule(name js.String, prelude js.Array[CSSToken], body js.Array[CSSParserRule]) CSSParserAtRule {
	return CSSParserAtRule{}.FromRef(
		bindings.NewCSSParserAtRuleByCSSParserAtRule(
			name.Ref(),
			prelude.Ref(),
			body.Ref()),
	)
}

func NewCSSParserAtRuleByCSSParserAtRule1(name js.String, prelude js.Array[CSSToken]) CSSParserAtRule {
	return CSSParserAtRule{}.FromRef(
		bindings.NewCSSParserAtRuleByCSSParserAtRule1(
			name.Ref(),
			prelude.Ref()),
	)
}

type CSSParserAtRule struct {
	CSSParserRule
}

func (this CSSParserAtRule) Once() CSSParserAtRule {
	this.Ref().Once()
	return this
}

func (this CSSParserAtRule) Ref() js.Ref {
	return this.CSSParserRule.Ref()
}

func (this CSSParserAtRule) FromRef(ref js.Ref) CSSParserAtRule {
	this.CSSParserRule = this.CSSParserRule.FromRef(ref)
	return this
}

func (this CSSParserAtRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSParserAtRule.name".
//
// The returned bool will be false if there is no such property.
func (this CSSParserAtRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSParserAtRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Prelude returns the value of property "CSSParserAtRule.prelude".
//
// The returned bool will be false if there is no such property.
func (this CSSParserAtRule) Prelude() (js.FrozenArray[CSSParserValue], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserAtRulePrelude(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserValue]{}.FromRef(_ret), _ok
}

// Body returns the value of property "CSSParserAtRule.body".
//
// The returned bool will be false if there is no such property.
func (this CSSParserAtRule) Body() (js.FrozenArray[CSSParserRule], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserAtRuleBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserRule]{}.FromRef(_ret), _ok
}

// ToString calls the method "CSSParserAtRule.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSParserAtRule) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSParserAtRuleToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSParserAtRule.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSParserAtRule) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserAtRuleToStringFunc(
			this.Ref(),
		),
	)
}

func NewCSSParserBlock(name js.String, body js.Array[CSSParserValue]) CSSParserBlock {
	return CSSParserBlock{}.FromRef(
		bindings.NewCSSParserBlockByCSSParserBlock(
			name.Ref(),
			body.Ref()),
	)
}

type CSSParserBlock struct {
	CSSParserValue
}

func (this CSSParserBlock) Once() CSSParserBlock {
	this.Ref().Once()
	return this
}

func (this CSSParserBlock) Ref() js.Ref {
	return this.CSSParserValue.Ref()
}

func (this CSSParserBlock) FromRef(ref js.Ref) CSSParserBlock {
	this.CSSParserValue = this.CSSParserValue.FromRef(ref)
	return this
}

func (this CSSParserBlock) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSParserBlock.name".
//
// The returned bool will be false if there is no such property.
func (this CSSParserBlock) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSParserBlockName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Body returns the value of property "CSSParserBlock.body".
//
// The returned bool will be false if there is no such property.
func (this CSSParserBlock) Body() (js.FrozenArray[CSSParserValue], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserBlockBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserValue]{}.FromRef(_ret), _ok
}

// ToString calls the method "CSSParserBlock.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSParserBlock) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSParserBlockToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSParserBlock.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSParserBlock) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserBlockToStringFunc(
			this.Ref(),
		),
	)
}

func NewCSSParserFunction(name js.String, args js.Array[js.Array[CSSParserValue]]) CSSParserFunction {
	return CSSParserFunction{}.FromRef(
		bindings.NewCSSParserFunctionByCSSParserFunction(
			name.Ref(),
			args.Ref()),
	)
}

type CSSParserFunction struct {
	CSSParserValue
}

func (this CSSParserFunction) Once() CSSParserFunction {
	this.Ref().Once()
	return this
}

func (this CSSParserFunction) Ref() js.Ref {
	return this.CSSParserValue.Ref()
}

func (this CSSParserFunction) FromRef(ref js.Ref) CSSParserFunction {
	this.CSSParserValue = this.CSSParserValue.FromRef(ref)
	return this
}

func (this CSSParserFunction) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSParserFunction.name".
//
// The returned bool will be false if there is no such property.
func (this CSSParserFunction) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSParserFunctionName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Args returns the value of property "CSSParserFunction.args".
//
// The returned bool will be false if there is no such property.
func (this CSSParserFunction) Args() (js.FrozenArray[js.FrozenArray[CSSParserValue]], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserFunctionArgs(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.FrozenArray[CSSParserValue]]{}.FromRef(_ret), _ok
}

// ToString calls the method "CSSParserFunction.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSParserFunction) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSParserFunctionToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSParserFunction.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSParserFunction) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserFunctionToStringFunc(
			this.Ref(),
		),
	)
}

func NewCSSParserQualifiedRule(prelude js.Array[CSSToken], body js.Array[CSSParserRule]) CSSParserQualifiedRule {
	return CSSParserQualifiedRule{}.FromRef(
		bindings.NewCSSParserQualifiedRuleByCSSParserQualifiedRule(
			prelude.Ref(),
			body.Ref()),
	)
}

func NewCSSParserQualifiedRuleByCSSParserQualifiedRule1(prelude js.Array[CSSToken]) CSSParserQualifiedRule {
	return CSSParserQualifiedRule{}.FromRef(
		bindings.NewCSSParserQualifiedRuleByCSSParserQualifiedRule1(
			prelude.Ref()),
	)
}

type CSSParserQualifiedRule struct {
	CSSParserRule
}

func (this CSSParserQualifiedRule) Once() CSSParserQualifiedRule {
	this.Ref().Once()
	return this
}

func (this CSSParserQualifiedRule) Ref() js.Ref {
	return this.CSSParserRule.Ref()
}

func (this CSSParserQualifiedRule) FromRef(ref js.Ref) CSSParserQualifiedRule {
	this.CSSParserRule = this.CSSParserRule.FromRef(ref)
	return this
}

func (this CSSParserQualifiedRule) Free() {
	this.Ref().Free()
}

// Prelude returns the value of property "CSSParserQualifiedRule.prelude".
//
// The returned bool will be false if there is no such property.
func (this CSSParserQualifiedRule) Prelude() (js.FrozenArray[CSSParserValue], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserQualifiedRulePrelude(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserValue]{}.FromRef(_ret), _ok
}

// Body returns the value of property "CSSParserQualifiedRule.body".
//
// The returned bool will be false if there is no such property.
func (this CSSParserQualifiedRule) Body() (js.FrozenArray[CSSParserRule], bool) {
	var _ok bool
	_ret := bindings.GetCSSParserQualifiedRuleBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CSSParserRule]{}.FromRef(_ret), _ok
}

// ToString calls the method "CSSParserQualifiedRule.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSParserQualifiedRule) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSParserQualifiedRuleToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSParserQualifiedRule.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSParserQualifiedRule) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserQualifiedRuleToStringFunc(
			this.Ref(),
		),
	)
}

type OneOf_CSSNumericValue_String_CSSKeywordValue struct {
	ref js.Ref
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) Free() {
	x.ref.Free()
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) FromRef(ref js.Ref) OneOf_CSSNumericValue_String_CSSKeywordValue {
	return OneOf_CSSNumericValue_String_CSSKeywordValue{
		ref: ref,
	}
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_CSSNumericValue_String_CSSKeywordValue) CSSKeywordValue() CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(x.ref)
}

type CSSPerspectiveValue = OneOf_CSSNumericValue_String_CSSKeywordValue
