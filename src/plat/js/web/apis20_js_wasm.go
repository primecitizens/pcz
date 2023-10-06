// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

func NewCSSColor(colorSpace CSSKeywordish, channels js.Array[CSSColorPercent], alpha CSSNumberish) (ret CSSColor) {
	ret.ref = bindings.NewCSSColorByCSSColor(
		colorSpace.Ref(),
		channels.Ref(),
		alpha.Ref())
	return
}

func NewCSSColorByCSSColor1(colorSpace CSSKeywordish, channels js.Array[CSSColorPercent]) (ret CSSColor) {
	ret.ref = bindings.NewCSSColorByCSSColor1(
		colorSpace.Ref(),
		channels.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSColor) ColorSpace() (ret CSSKeywordish, ok bool) {
	ok = js.True == bindings.GetCSSColorColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetColorSpace sets the value of property "CSSColor.colorSpace" to val.
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
// It returns ok=false if there is no such property.
func (this CSSColor) Channels() (ret js.ObservableArray[CSSColorPercent], ok bool) {
	ok = js.True == bindings.GetCSSColorChannels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChannels sets the value of property "CSSColor.channels" to val.
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
// It returns ok=false if there is no such property.
func (this CSSColor) Alpha() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSColorAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSColor.alpha" to val.
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
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Src returns the value of property "CSSColorProfileRule.src".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RenderingIntent returns the value of property "CSSColorProfileRule.renderingIntent".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) RenderingIntent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleRenderingIntent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Components returns the value of property "CSSColorProfileRule.components".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Components() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleComponents(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasParse returns true if the staticmethod "CSSColorValue.parse" exists.
func (this CSSColorValue) HasParse() bool {
	return js.True == bindings.HasCSSColorValueParse(
		this.Ref(),
	)
}

// ParseFunc returns the staticmethod "CSSColorValue.parse".
func (this CSSColorValue) ParseFunc() (fn js.Func[func(cssText js.String) OneOf_CSSColorValue_CSSStyleValue]) {
	return fn.FromRef(
		bindings.CSSColorValueParseFunc(
			this.Ref(),
		),
	)
}

// Parse calls the staticmethod "CSSColorValue.parse".
func (this CSSColorValue) Parse(cssText js.String) (ret OneOf_CSSColorValue_CSSStyleValue) {
	bindings.CallCSSColorValueParse(
		this.Ref(), js.Pointer(&ret),
		cssText.Ref(),
	)

	return
}

// TryParse calls the staticmethod "CSSColorValue.parse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSColorValue) TryParse(cssText js.String) (ret OneOf_CSSColorValue_CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSColorValueParse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		cssText.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this CSSConditionRule) ConditionText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSConditionRuleConditionText(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSContainerRule) ContainerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSContainerRuleContainerName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContainerQuery returns the value of property "CSSContainerRule.containerQuery".
//
// It returns ok=false if there is no such property.
func (this CSSContainerRule) ContainerQuery() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSContainerRuleContainerQuery(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "CSSCounterStyleRule.name" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) System() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSystem(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSystem sets the value of property "CSSCounterStyleRule.system" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Symbols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSymbols(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSymbols sets the value of property "CSSCounterStyleRule.symbols" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) AdditiveSymbols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleAdditiveSymbols(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAdditiveSymbols sets the value of property "CSSCounterStyleRule.additiveSymbols" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Negative() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleNegative(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNegative sets the value of property "CSSCounterStyleRule.negative" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRulePrefix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPrefix sets the value of property "CSSCounterStyleRule.prefix" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Suffix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSuffix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSuffix sets the value of property "CSSCounterStyleRule.suffix" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Range() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleRange(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRange sets the value of property "CSSCounterStyleRule.range" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Pad() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRulePad(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPad sets the value of property "CSSCounterStyleRule.pad" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) SpeakAs() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSpeakAs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSpeakAs sets the value of property "CSSCounterStyleRule.speakAs" to val.
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
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Fallback() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleFallback(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFallback sets the value of property "CSSCounterStyleRule.fallback" to val.
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
// It returns ok=false if there is no such property.
func (this CSSFontFaceRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSFontFaceRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasSet returns true if the method "CSSFontFeatureValuesMap.set" exists.
func (this CSSFontFeatureValuesMap) HasSet() bool {
	return js.True == bindings.HasCSSFontFeatureValuesMapSet(
		this.Ref(),
	)
}

// SetFunc returns the method "CSSFontFeatureValuesMap.set".
func (this CSSFontFeatureValuesMap) SetFunc() (fn js.Func[func(featureValueName js.String, values OneOf_Uint32_ArrayUint32)]) {
	return fn.FromRef(
		bindings.CSSFontFeatureValuesMapSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CSSFontFeatureValuesMap.set".
func (this CSSFontFeatureValuesMap) Set(featureValueName js.String, values OneOf_Uint32_ArrayUint32) (ret js.Void) {
	bindings.CallCSSFontFeatureValuesMapSet(
		this.Ref(), js.Pointer(&ret),
		featureValueName.Ref(),
		values.Ref(),
	)

	return
}

// TrySet calls the method "CSSFontFeatureValuesMap.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSFontFeatureValuesMap) TrySet(featureValueName js.String, values OneOf_Uint32_ArrayUint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSFontFeatureValuesMapSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		featureValueName.Ref(),
		values.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) FontFamily() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleFontFamily(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFontFamily sets the value of property "CSSFontFeatureValuesRule.fontFamily" to val.
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
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Annotation() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleAnnotation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ornaments returns the value of property "CSSFontFeatureValuesRule.ornaments".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Ornaments() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleOrnaments(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Stylistic returns the value of property "CSSFontFeatureValuesRule.stylistic".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Stylistic() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleStylistic(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Swash returns the value of property "CSSFontFeatureValuesRule.swash".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Swash() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleSwash(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CharacterVariant returns the value of property "CSSFontFeatureValuesRule.characterVariant".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) CharacterVariant() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleCharacterVariant(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Styleset returns the value of property "CSSFontFeatureValuesRule.styleset".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Styleset() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleStyleset(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FontFamily returns the value of property "CSSFontPaletteValuesRule.fontFamily".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) FontFamily() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleFontFamily(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BasePalette returns the value of property "CSSFontPaletteValuesRule.basePalette".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) BasePalette() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleBasePalette(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OverrideColors returns the value of property "CSSFontPaletteValuesRule.overrideColors".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) OverrideColors() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleOverrideColors(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSGroupingRule) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSGroupingRuleCssRules(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInsertRule returns true if the method "CSSGroupingRule.insertRule" exists.
func (this CSSGroupingRule) HasInsertRule() bool {
	return js.True == bindings.HasCSSGroupingRuleInsertRule(
		this.Ref(),
	)
}

// InsertRuleFunc returns the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRuleFunc() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleInsertRuleFunc(
			this.Ref(),
		),
	)
}

// InsertRule calls the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRule(rule js.String, index uint32) (ret uint32) {
	bindings.CallCSSGroupingRuleInsertRule(
		this.Ref(), js.Pointer(&ret),
		rule.Ref(),
		uint32(index),
	)

	return
}

// TryInsertRule calls the method "CSSGroupingRule.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSGroupingRule) TryInsertRule(rule js.String, index uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGroupingRuleInsertRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
		uint32(index),
	)

	return
}

// HasInsertRule1 returns true if the method "CSSGroupingRule.insertRule" exists.
func (this CSSGroupingRule) HasInsertRule1() bool {
	return js.True == bindings.HasCSSGroupingRuleInsertRule1(
		this.Ref(),
	)
}

// InsertRule1Func returns the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRule1Func() (fn js.Func[func(rule js.String) uint32]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleInsertRule1Func(
			this.Ref(),
		),
	)
}

// InsertRule1 calls the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRule1(rule js.String) (ret uint32) {
	bindings.CallCSSGroupingRuleInsertRule1(
		this.Ref(), js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryInsertRule1 calls the method "CSSGroupingRule.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSGroupingRule) TryInsertRule1(rule js.String) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGroupingRuleInsertRule1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasDeleteRule returns true if the method "CSSGroupingRule.deleteRule" exists.
func (this CSSGroupingRule) HasDeleteRule() bool {
	return js.True == bindings.HasCSSGroupingRuleDeleteRule(
		this.Ref(),
	)
}

// DeleteRuleFunc returns the method "CSSGroupingRule.deleteRule".
func (this CSSGroupingRule) DeleteRuleFunc() (fn js.Func[func(index uint32)]) {
	return fn.FromRef(
		bindings.CSSGroupingRuleDeleteRuleFunc(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSGroupingRule.deleteRule".
func (this CSSGroupingRule) DeleteRule(index uint32) (ret js.Void) {
	bindings.CallCSSGroupingRuleDeleteRule(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDeleteRule calls the method "CSSGroupingRule.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSGroupingRule) TryDeleteRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGroupingRuleDeleteRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewCSSHSL(h CSSColorAngle, s CSSColorPercent, l CSSColorPercent, alpha CSSColorPercent) (ret CSSHSL) {
	ret.ref = bindings.NewCSSHSLByCSSHSL(
		h.Ref(),
		s.Ref(),
		l.Ref(),
		alpha.Ref())
	return
}

func NewCSSHSLByCSSHSL1(h CSSColorAngle, s CSSColorPercent, l CSSColorPercent) (ret CSSHSL) {
	ret.ref = bindings.NewCSSHSLByCSSHSL1(
		h.Ref(),
		s.Ref(),
		l.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSHSL) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSHSLH(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSHSL.h" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHSL) S() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLS(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetS sets the value of property "CSSHSL.s" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHSL) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSHSL.l" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHSL) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSHSL.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSHWB(h CSSNumericValue, w CSSNumberish, b CSSNumberish, alpha CSSNumberish) (ret CSSHWB) {
	ret.ref = bindings.NewCSSHWBByCSSHWB(
		h.Ref(),
		w.Ref(),
		b.Ref(),
		alpha.Ref())
	return
}

func NewCSSHWBByCSSHWB1(h CSSNumericValue, w CSSNumberish, b CSSNumberish) (ret CSSHWB) {
	ret.ref = bindings.NewCSSHWBByCSSHWB1(
		h.Ref(),
		w.Ref(),
		b.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSHWB) H() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSHWBH(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSHWB.h" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHWB) W() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBW(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetW sets the value of property "CSSHWB.w" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHWB) B() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSHWB.b" to val.
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
// It returns ok=false if there is no such property.
func (this CSSHWB) Alpha() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSHWB.alpha" to val.
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
// It returns ok=false if there is no such property.
func (this CSSImportRule) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Media returns the value of property "CSSImportRule.media".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) Media() (ret MediaList, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StyleSheet returns the value of property "CSSImportRule.styleSheet".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) StyleSheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleStyleSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LayerName returns the value of property "CSSImportRule.layerName".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) LayerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleLayerName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SupportsText returns the value of property "CSSImportRule.supportsText".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) SupportsText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleSupportsText(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSKeyframeRule) KeyText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeyframeRuleKeyText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetKeyText sets the value of property "CSSKeyframeRule.keyText" to val.
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
// It returns ok=false if there is no such property.
func (this CSSKeyframeRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSKeyframeRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "CSSKeyframesRule.name" to val.
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
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleCssRules(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "CSSKeyframesRule.length".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "CSSKeyframesRule." exists.
func (this CSSKeyframesRule) HasGet() bool {
	return js.True == bindings.HasCSSKeyframesRuleGet(
		this.Ref(),
	)
}

// GetFunc returns the method "CSSKeyframesRule.".
func (this CSSKeyframesRule) GetFunc() (fn js.Func[func(index uint32) CSSKeyframeRule]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "CSSKeyframesRule.".
func (this CSSKeyframesRule) Get(index uint32) (ret CSSKeyframeRule) {
	bindings.CallCSSKeyframesRuleGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "CSSKeyframesRule."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryGet(index uint32) (ret CSSKeyframeRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendRule returns true if the method "CSSKeyframesRule.appendRule" exists.
func (this CSSKeyframesRule) HasAppendRule() bool {
	return js.True == bindings.HasCSSKeyframesRuleAppendRule(
		this.Ref(),
	)
}

// AppendRuleFunc returns the method "CSSKeyframesRule.appendRule".
func (this CSSKeyframesRule) AppendRuleFunc() (fn js.Func[func(rule js.String)]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleAppendRuleFunc(
			this.Ref(),
		),
	)
}

// AppendRule calls the method "CSSKeyframesRule.appendRule".
func (this CSSKeyframesRule) AppendRule(rule js.String) (ret js.Void) {
	bindings.CallCSSKeyframesRuleAppendRule(
		this.Ref(), js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryAppendRule calls the method "CSSKeyframesRule.appendRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryAppendRule(rule js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleAppendRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasDeleteRule returns true if the method "CSSKeyframesRule.deleteRule" exists.
func (this CSSKeyframesRule) HasDeleteRule() bool {
	return js.True == bindings.HasCSSKeyframesRuleDeleteRule(
		this.Ref(),
	)
}

// DeleteRuleFunc returns the method "CSSKeyframesRule.deleteRule".
func (this CSSKeyframesRule) DeleteRuleFunc() (fn js.Func[func(sel js.String)]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleDeleteRuleFunc(
			this.Ref(),
		),
	)
}

// DeleteRule calls the method "CSSKeyframesRule.deleteRule".
func (this CSSKeyframesRule) DeleteRule(sel js.String) (ret js.Void) {
	bindings.CallCSSKeyframesRuleDeleteRule(
		this.Ref(), js.Pointer(&ret),
		sel.Ref(),
	)

	return
}

// TryDeleteRule calls the method "CSSKeyframesRule.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryDeleteRule(sel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleDeleteRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sel.Ref(),
	)

	return
}

// HasFindRule returns true if the method "CSSKeyframesRule.findRule" exists.
func (this CSSKeyframesRule) HasFindRule() bool {
	return js.True == bindings.HasCSSKeyframesRuleFindRule(
		this.Ref(),
	)
}

// FindRuleFunc returns the method "CSSKeyframesRule.findRule".
func (this CSSKeyframesRule) FindRuleFunc() (fn js.Func[func(sel js.String) CSSKeyframeRule]) {
	return fn.FromRef(
		bindings.CSSKeyframesRuleFindRuleFunc(
			this.Ref(),
		),
	)
}

// FindRule calls the method "CSSKeyframesRule.findRule".
func (this CSSKeyframesRule) FindRule(sel js.String) (ret CSSKeyframeRule) {
	bindings.CallCSSKeyframesRuleFindRule(
		this.Ref(), js.Pointer(&ret),
		sel.Ref(),
	)

	return
}

// TryFindRule calls the method "CSSKeyframesRule.findRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryFindRule(sel js.String) (ret CSSKeyframeRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleFindRule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sel.Ref(),
	)

	return
}

func NewCSSLCH(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle, alpha CSSColorPercent) (ret CSSLCH) {
	ret.ref = bindings.NewCSSLCHByCSSLCH(
		l.Ref(),
		c.Ref(),
		h.Ref(),
		alpha.Ref())
	return
}

func NewCSSLCHByCSSLCH1(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle) (ret CSSLCH) {
	ret.ref = bindings.NewCSSLCHByCSSLCH1(
		l.Ref(),
		c.Ref(),
		h.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSLCH) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSLCH.l" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLCH) C() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHC(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "CSSLCH.c" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLCH) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSLCHH(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSLCH.h" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLCH) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSLab(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber, alpha CSSColorPercent) (ret CSSLab) {
	ret.ref = bindings.NewCSSLabByCSSLab(
		l.Ref(),
		a.Ref(),
		b.Ref(),
		alpha.Ref())
	return
}

func NewCSSLabByCSSLab1(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber) (ret CSSLab) {
	ret.ref = bindings.NewCSSLabByCSSLab1(
		l.Ref(),
		a.Ref(),
		b.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSLab) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLabL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSLab.l" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLab) A() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSLabA(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "CSSLab.a" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLab) B() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSLabB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSLab.b" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLab) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLabAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSLab.alpha" to val.
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
// It returns ok=false if there is no such property.
func (this CSSLayerBlockRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSLayerBlockRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSLayerStatementRule) NameList() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetCSSLayerStatementRuleNameList(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSMarginRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSMarginRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "CSSMarginRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSMarginRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSMarginRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSMathClamp(lower CSSNumberish, value CSSNumberish, upper CSSNumberish) (ret CSSMathClamp) {
	ret.ref = bindings.NewCSSMathClampByCSSMathClamp(
		lower.Ref(),
		value.Ref(),
		upper.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Lower() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampLower(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "CSSMathClamp.value".
//
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Upper returns the value of property "CSSMathClamp.upper".
//
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Upper() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampUpper(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSMathInvert(arg CSSNumberish) (ret CSSMathInvert) {
	ret.ref = bindings.NewCSSMathInvertByCSSMathInvert(
		arg.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathInvert) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathInvertValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSMathMax(args ...CSSNumberish) (ret CSSMathMax) {
	ret.ref = bindings.NewCSSMathMaxByCSSMathMax(
		js.SliceData(args),
		js.SizeU(len(args)))
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathMax) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathMaxValues(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSMathMin(args ...CSSNumberish) (ret CSSMathMin) {
	ret.ref = bindings.NewCSSMathMinByCSSMathMin(
		js.SliceData(args),
		js.SizeU(len(args)))
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathMin) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathMinValues(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSMathNegate(arg CSSNumberish) (ret CSSMathNegate) {
	ret.ref = bindings.NewCSSMathNegateByCSSMathNegate(
		arg.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathNegate) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathNegateValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewCSSMathProduct(args ...CSSNumberish) (ret CSSMathProduct) {
	ret.ref = bindings.NewCSSMathProductByCSSMathProduct(
		js.SliceData(args),
		js.SizeU(len(args)))
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathProduct) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathProductValues(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSMathValue) Operator() (ret CSSMathOperator, ok bool) {
	ok = js.True == bindings.GetCSSMathValueOperator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewDOMMatrixReadOnly(init OneOf_String_ArrayFloat64) (ret DOMMatrixReadOnly) {
	ret.ref = bindings.NewDOMMatrixReadOnlyByDOMMatrixReadOnly(
		init.Ref())
	return
}

func NewDOMMatrixReadOnlyByDOMMatrixReadOnly1() (ret DOMMatrixReadOnly) {
	ret.ref = bindings.NewDOMMatrixReadOnlyByDOMMatrixReadOnly1()
	return
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
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) A() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyA(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// B returns the value of property "DOMMatrixReadOnly.b".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) B() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// C returns the value of property "DOMMatrixReadOnly.c".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) C() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyC(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// D returns the value of property "DOMMatrixReadOnly.d".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) D() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyD(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// E returns the value of property "DOMMatrixReadOnly.e".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) E() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyE(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// F returns the value of property "DOMMatrixReadOnly.f".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) F() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyF(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M11 returns the value of property "DOMMatrixReadOnly.m11".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M11() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM11(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M12 returns the value of property "DOMMatrixReadOnly.m12".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M12() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM12(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M13 returns the value of property "DOMMatrixReadOnly.m13".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M13() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM13(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M14 returns the value of property "DOMMatrixReadOnly.m14".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M14() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM14(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M21 returns the value of property "DOMMatrixReadOnly.m21".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M21() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM21(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M22 returns the value of property "DOMMatrixReadOnly.m22".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M22() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM22(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M23 returns the value of property "DOMMatrixReadOnly.m23".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M23() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM23(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M24 returns the value of property "DOMMatrixReadOnly.m24".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M24() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM24(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M31 returns the value of property "DOMMatrixReadOnly.m31".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M31() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM31(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M32 returns the value of property "DOMMatrixReadOnly.m32".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M32() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM32(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M33 returns the value of property "DOMMatrixReadOnly.m33".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M33() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM33(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M34 returns the value of property "DOMMatrixReadOnly.m34".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M34() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM34(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M41 returns the value of property "DOMMatrixReadOnly.m41".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M41() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM41(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M42 returns the value of property "DOMMatrixReadOnly.m42".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M42() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM42(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M43 returns the value of property "DOMMatrixReadOnly.m43".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M43() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM43(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// M44 returns the value of property "DOMMatrixReadOnly.m44".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M44() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM44(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Is2D returns the value of property "DOMMatrixReadOnly.is2D".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) Is2D() (ret bool, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyIs2D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsIdentity returns the value of property "DOMMatrixReadOnly.isIdentity".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) IsIdentity() (ret bool, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyIsIdentity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasFromMatrix returns true if the staticmethod "DOMMatrixReadOnly.fromMatrix" exists.
func (this DOMMatrixReadOnly) HasFromMatrix() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFromMatrix(
		this.Ref(),
	)
}

// FromMatrixFunc returns the staticmethod "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrixFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromMatrixFunc(
			this.Ref(),
		),
	)
}

// FromMatrix calls the staticmethod "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrix(other DOMMatrixInit) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromMatrix(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromMatrix calls the staticmethod "DOMMatrixReadOnly.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromMatrix(other DOMMatrixInit) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFromMatrix1 returns true if the staticmethod "DOMMatrixReadOnly.fromMatrix" exists.
func (this DOMMatrixReadOnly) HasFromMatrix1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFromMatrix1(
		this.Ref(),
	)
}

// FromMatrix1Func returns the staticmethod "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrix1Func() (fn js.Func[func() DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromMatrix1Func(
			this.Ref(),
		),
	)
}

// FromMatrix1 calls the staticmethod "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrix1() (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromMatrix1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFromMatrix1 calls the staticmethod "DOMMatrixReadOnly.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromMatrix1() (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromMatrix1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFromFloat32Array returns true if the staticmethod "DOMMatrixReadOnly.fromFloat32Array" exists.
func (this DOMMatrixReadOnly) HasFromFloat32Array() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFromFloat32Array(
		this.Ref(),
	)
}

// FromFloat32ArrayFunc returns the staticmethod "DOMMatrixReadOnly.fromFloat32Array".
func (this DOMMatrixReadOnly) FromFloat32ArrayFunc() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat32Array calls the staticmethod "DOMMatrixReadOnly.fromFloat32Array".
func (this DOMMatrixReadOnly) FromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromFloat32Array(
		this.Ref(), js.Pointer(&ret),
		array32.Ref(),
	)

	return
}

// TryFromFloat32Array calls the staticmethod "DOMMatrixReadOnly.fromFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromFloat32Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array32.Ref(),
	)

	return
}

// HasFromFloat64Array returns true if the staticmethod "DOMMatrixReadOnly.fromFloat64Array" exists.
func (this DOMMatrixReadOnly) HasFromFloat64Array() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFromFloat64Array(
		this.Ref(),
	)
}

// FromFloat64ArrayFunc returns the staticmethod "DOMMatrixReadOnly.fromFloat64Array".
func (this DOMMatrixReadOnly) FromFloat64ArrayFunc() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrixReadOnly]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFromFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// FromFloat64Array calls the staticmethod "DOMMatrixReadOnly.fromFloat64Array".
func (this DOMMatrixReadOnly) FromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromFloat64Array(
		this.Ref(), js.Pointer(&ret),
		array64.Ref(),
	)

	return
}

// TryFromFloat64Array calls the staticmethod "DOMMatrixReadOnly.fromFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromFloat64Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array64.Ref(),
	)

	return
}

// HasTranslate returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasTranslate() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTranslate(
		this.Ref(),
	)
}

// TranslateFunc returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) TranslateFunc() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate(tx float64, ty float64, tz float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// TryTranslate calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate(tx float64, ty float64, tz float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// HasTranslate1 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasTranslate1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTranslate1(
		this.Ref(),
	)
}

// Translate1Func returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate1Func() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate1Func(
			this.Ref(),
		),
	)
}

// Translate1 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate1(tx float64, ty float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate1(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
		float64(ty),
	)

	return
}

// TryTranslate1 calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate1(tx float64, ty float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
	)

	return
}

// HasTranslate2 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasTranslate2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTranslate2(
		this.Ref(),
	)
}

// Translate2Func returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate2Func() (fn js.Func[func(tx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate2Func(
			this.Ref(),
		),
	)
}

// Translate2 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate2(tx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate2(
		this.Ref(), js.Pointer(&ret),
		float64(tx),
	)

	return
}

// TryTranslate2 calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate2(tx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
	)

	return
}

// HasTranslate3 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasTranslate3() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTranslate3(
		this.Ref(),
	)
}

// Translate3Func returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTranslate3Func(
			this.Ref(),
		),
	)
}

// Translate3 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate3() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTranslate3 calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScale returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale(
		this.Ref(),
	)
}

// ScaleFunc returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) ScaleFunc() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// TryScale calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasScale1 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale1(
		this.Ref(),
	)
}

// Scale1Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale1Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale1Func(
			this.Ref(),
		),
	)
}

// Scale1 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale1(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScale1 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// HasScale2 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale2(
		this.Ref(),
	)
}

// Scale2Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale2Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale2Func(
			this.Ref(),
		),
	)
}

// Scale2 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale2(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// TryScale2 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// HasScale3 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale3() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3(
		this.Ref(),
	)
}

// Scale3Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale3Func() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3Func(
			this.Ref(),
		),
	)
}

// Scale3 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// TryScale3 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// HasScale4 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale4() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale4(
		this.Ref(),
	)
}

// Scale4Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale4Func() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale4Func(
			this.Ref(),
		),
	)
}

// Scale4 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale4(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale4(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// TryScale4 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale4(scaleX float64, scaleY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasScale5 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale5() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale5(
		this.Ref(),
	)
}

// Scale5Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale5Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale5Func(
			this.Ref(),
		),
	)
}

// Scale5 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale5(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale5(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScale5 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale5(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasScale6 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasScale6() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale6(
		this.Ref(),
	)
}

// Scale6Func returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale6Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale6Func(
			this.Ref(),
		),
	)
}

// Scale6 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale6() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale6(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScale6 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale6() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScaleNonUniform returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasScaleNonUniform() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScaleNonUniform(
		this.Ref(),
	)
}

// ScaleNonUniformFunc returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniformFunc() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniformFunc(
			this.Ref(),
		),
	)
}

// ScaleNonUniform calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// TryScaleNonUniform calls the method "DOMMatrixReadOnly.scaleNonUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScaleNonUniform(scaleX float64, scaleY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScaleNonUniform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasScaleNonUniform1 returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasScaleNonUniform1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScaleNonUniform1(
		this.Ref(),
	)
}

// ScaleNonUniform1Func returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform1Func() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniform1Func(
			this.Ref(),
		),
	)
}

// ScaleNonUniform1 calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform1(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform1(
		this.Ref(), js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScaleNonUniform1 calls the method "DOMMatrixReadOnly.scaleNonUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScaleNonUniform1(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScaleNonUniform1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasScaleNonUniform2 returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasScaleNonUniform2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScaleNonUniform2(
		this.Ref(),
	)
}

// ScaleNonUniform2Func returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScaleNonUniform2Func(
			this.Ref(),
		),
	)
}

// ScaleNonUniform2 calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform2() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScaleNonUniform2 calls the method "DOMMatrixReadOnly.scaleNonUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScaleNonUniform2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScaleNonUniform2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScale3d returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasScale3d() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3d(
		this.Ref(),
	)
}

// Scale3dFunc returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3dFunc() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3dFunc(
			this.Ref(),
		),
	)
}

// Scale3d calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// TryScale3d calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasScale3d1 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasScale3d1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3d1(
		this.Ref(),
	)
}

// Scale3d1Func returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d1Func() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d1Func(
			this.Ref(),
		),
	)
}

// Scale3d1 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d1(scale float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d1(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// TryScale3d1 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d1(scale float64, originX float64, originY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// HasScale3d2 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasScale3d2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3d2(
		this.Ref(),
	)
}

// Scale3d2Func returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d2Func() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d2Func(
			this.Ref(),
		),
	)
}

// Scale3d2 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d2(scale float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d2(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
		float64(originX),
	)

	return
}

// TryScale3d2 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d2(scale float64, originX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
	)

	return
}

// HasScale3d3 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasScale3d3() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3d3(
		this.Ref(),
	)
}

// Scale3d3Func returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d3Func() (fn js.Func[func(scale float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d3Func(
			this.Ref(),
		),
	)
}

// Scale3d3 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d3(scale float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d3(
		this.Ref(), js.Pointer(&ret),
		float64(scale),
	)

	return
}

// TryScale3d3 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d3(scale float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
	)

	return
}

// HasScale3d4 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasScale3d4() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyScale3d4(
		this.Ref(),
	)
}

// Scale3d4Func returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyScale3d4Func(
			this.Ref(),
		),
	)
}

// Scale3d4 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d4() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScale3d4 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotate returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasRotate() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotate(
		this.Ref(),
	)
}

// RotateFunc returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) RotateFunc() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// TryRotate calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// HasRotate1 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasRotate1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotate1(
		this.Ref(),
	)
}

// Rotate1Func returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate1Func() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate1Func(
			this.Ref(),
		),
	)
}

// Rotate1 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate1(rotX float64, rotY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate1(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
		float64(rotY),
	)

	return
}

// TryRotate1 calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate1(rotX float64, rotY float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
	)

	return
}

// HasRotate2 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasRotate2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotate2(
		this.Ref(),
	)
}

// Rotate2Func returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate2Func() (fn js.Func[func(rotX float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate2Func(
			this.Ref(),
		),
	)
}

// Rotate2 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate2(rotX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate2(
		this.Ref(), js.Pointer(&ret),
		float64(rotX),
	)

	return
}

// TryRotate2 calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate2(rotX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
	)

	return
}

// HasRotate3 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasRotate3() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotate3(
		this.Ref(),
	)
}

// Rotate3Func returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate3Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotate3Func(
			this.Ref(),
		),
	)
}

// Rotate3 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate3() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotate3 calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotateFromVector returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasRotateFromVector() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateFromVector(
		this.Ref(),
	)
}

// RotateFromVectorFunc returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVectorFunc() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVectorFunc(
			this.Ref(),
		),
	)
}

// RotateFromVector calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateFromVector calls the method "DOMMatrixReadOnly.rotateFromVector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateFromVector(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateFromVector(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotateFromVector1 returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasRotateFromVector1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateFromVector1(
		this.Ref(),
	)
}

// RotateFromVector1Func returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector1Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVector1Func(
			this.Ref(),
		),
	)
}

// RotateFromVector1 calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector1(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateFromVector1 calls the method "DOMMatrixReadOnly.rotateFromVector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateFromVector1(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateFromVector1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasRotateFromVector2 returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasRotateFromVector2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateFromVector2(
		this.Ref(),
	)
}

// RotateFromVector2Func returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector2Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateFromVector2Func(
			this.Ref(),
		),
	)
}

// RotateFromVector2 calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector2() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotateFromVector2 calls the method "DOMMatrixReadOnly.rotateFromVector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateFromVector2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateFromVector2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRotateAxisAngle returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasRotateAxisAngle() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateAxisAngle(
		this.Ref(),
	)
}

// RotateAxisAngleFunc returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngleFunc() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngleFunc(
			this.Ref(),
		),
	)
}

// RotateAxisAngle calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle(x float64, y float64, z float64, angle float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// TryRotateAxisAngle calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle(x float64, y float64, z float64, angle float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// HasRotateAxisAngle1 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasRotateAxisAngle1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateAxisAngle1(
		this.Ref(),
	)
}

// RotateAxisAngle1Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle1Func() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle1Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle1 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle1(x float64, y float64, z float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// TryRotateAxisAngle1 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle1(x float64, y float64, z float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// HasRotateAxisAngle2 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasRotateAxisAngle2() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateAxisAngle2(
		this.Ref(),
	)
}

// RotateAxisAngle2Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle2Func() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle2Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle2 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle2(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle2(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryRotateAxisAngle2 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle2(x float64, y float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotateAxisAngle3 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasRotateAxisAngle3() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateAxisAngle3(
		this.Ref(),
	)
}

// RotateAxisAngle3Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle3Func() (fn js.Func[func(x float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle3Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle3 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle3(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle3(
		this.Ref(), js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateAxisAngle3 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle3(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasRotateAxisAngle4 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasRotateAxisAngle4() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyRotateAxisAngle4(
		this.Ref(),
	)
}

// RotateAxisAngle4Func returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle4Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyRotateAxisAngle4Func(
			this.Ref(),
		),
	)
}

// RotateAxisAngle4 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle4() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle4(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRotateAxisAngle4 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSkewX returns true if the method "DOMMatrixReadOnly.skewX" exists.
func (this DOMMatrixReadOnly) HasSkewX() bool {
	return js.True == bindings.HasDOMMatrixReadOnlySkewX(
		this.Ref(),
	)
}

// SkewXFunc returns the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewXFunc() (fn js.Func[func(sx float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewXFunc(
			this.Ref(),
		),
	)
}

// SkewX calls the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewX(sx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewX(
		this.Ref(), js.Pointer(&ret),
		float64(sx),
	)

	return
}

// TrySkewX calls the method "DOMMatrixReadOnly.skewX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewX(sx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewX(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(sx),
	)

	return
}

// HasSkewX1 returns true if the method "DOMMatrixReadOnly.skewX" exists.
func (this DOMMatrixReadOnly) HasSkewX1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlySkewX1(
		this.Ref(),
	)
}

// SkewX1Func returns the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewX1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewX1Func(
			this.Ref(),
		),
	)
}

// SkewX1 calls the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewX1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewX1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySkewX1 calls the method "DOMMatrixReadOnly.skewX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewX1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewX1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSkewY returns true if the method "DOMMatrixReadOnly.skewY" exists.
func (this DOMMatrixReadOnly) HasSkewY() bool {
	return js.True == bindings.HasDOMMatrixReadOnlySkewY(
		this.Ref(),
	)
}

// SkewYFunc returns the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewYFunc() (fn js.Func[func(sy float64) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewYFunc(
			this.Ref(),
		),
	)
}

// SkewY calls the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewY(sy float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewY(
		this.Ref(), js.Pointer(&ret),
		float64(sy),
	)

	return
}

// TrySkewY calls the method "DOMMatrixReadOnly.skewY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewY(sy float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewY(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(sy),
	)

	return
}

// HasSkewY1 returns true if the method "DOMMatrixReadOnly.skewY" exists.
func (this DOMMatrixReadOnly) HasSkewY1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlySkewY1(
		this.Ref(),
	)
}

// SkewY1Func returns the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewY1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlySkewY1Func(
			this.Ref(),
		),
	)
}

// SkewY1 calls the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewY1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewY1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySkewY1 calls the method "DOMMatrixReadOnly.skewY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewY1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewY1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMultiply returns true if the method "DOMMatrixReadOnly.multiply" exists.
func (this DOMMatrixReadOnly) HasMultiply() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyMultiply(
		this.Ref(),
	)
}

// MultiplyFunc returns the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) MultiplyFunc() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyMultiplyFunc(
			this.Ref(),
		),
	)
}

// Multiply calls the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) Multiply(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyMultiply(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryMultiply calls the method "DOMMatrixReadOnly.multiply"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryMultiply(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyMultiply(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasMultiply1 returns true if the method "DOMMatrixReadOnly.multiply" exists.
func (this DOMMatrixReadOnly) HasMultiply1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyMultiply1(
		this.Ref(),
	)
}

// Multiply1Func returns the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) Multiply1Func() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyMultiply1Func(
			this.Ref(),
		),
	)
}

// Multiply1 calls the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) Multiply1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyMultiply1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMultiply1 calls the method "DOMMatrixReadOnly.multiply"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryMultiply1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyMultiply1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFlipX returns true if the method "DOMMatrixReadOnly.flipX" exists.
func (this DOMMatrixReadOnly) HasFlipX() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFlipX(
		this.Ref(),
	)
}

// FlipXFunc returns the method "DOMMatrixReadOnly.flipX".
func (this DOMMatrixReadOnly) FlipXFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFlipXFunc(
			this.Ref(),
		),
	)
}

// FlipX calls the method "DOMMatrixReadOnly.flipX".
func (this DOMMatrixReadOnly) FlipX() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyFlipX(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlipX calls the method "DOMMatrixReadOnly.flipX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFlipX() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFlipX(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFlipY returns true if the method "DOMMatrixReadOnly.flipY" exists.
func (this DOMMatrixReadOnly) HasFlipY() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyFlipY(
		this.Ref(),
	)
}

// FlipYFunc returns the method "DOMMatrixReadOnly.flipY".
func (this DOMMatrixReadOnly) FlipYFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyFlipYFunc(
			this.Ref(),
		),
	)
}

// FlipY calls the method "DOMMatrixReadOnly.flipY".
func (this DOMMatrixReadOnly) FlipY() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyFlipY(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlipY calls the method "DOMMatrixReadOnly.flipY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFlipY() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFlipY(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInverse returns true if the method "DOMMatrixReadOnly.inverse" exists.
func (this DOMMatrixReadOnly) HasInverse() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyInverse(
		this.Ref(),
	)
}

// InverseFunc returns the method "DOMMatrixReadOnly.inverse".
func (this DOMMatrixReadOnly) InverseFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyInverseFunc(
			this.Ref(),
		),
	)
}

// Inverse calls the method "DOMMatrixReadOnly.inverse".
func (this DOMMatrixReadOnly) Inverse() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyInverse(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryInverse calls the method "DOMMatrixReadOnly.inverse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryInverse() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyInverse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTransformPoint returns true if the method "DOMMatrixReadOnly.transformPoint" exists.
func (this DOMMatrixReadOnly) HasTransformPoint() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTransformPoint(
		this.Ref(),
	)
}

// TransformPointFunc returns the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPointFunc() (fn js.Func[func(point DOMPointInit) DOMPoint]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTransformPointFunc(
			this.Ref(),
		),
	)
}

// TransformPoint calls the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPoint(point DOMPointInit) (ret DOMPoint) {
	bindings.CallDOMMatrixReadOnlyTransformPoint(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryTransformPoint calls the method "DOMMatrixReadOnly.transformPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTransformPoint(point DOMPointInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTransformPoint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasTransformPoint1 returns true if the method "DOMMatrixReadOnly.transformPoint" exists.
func (this DOMMatrixReadOnly) HasTransformPoint1() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyTransformPoint1(
		this.Ref(),
	)
}

// TransformPoint1Func returns the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPoint1Func() (fn js.Func[func() DOMPoint]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyTransformPoint1Func(
			this.Ref(),
		),
	)
}

// TransformPoint1 calls the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPoint1() (ret DOMPoint) {
	bindings.CallDOMMatrixReadOnlyTransformPoint1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTransformPoint1 calls the method "DOMMatrixReadOnly.transformPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTransformPoint1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTransformPoint1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToFloat32Array returns true if the method "DOMMatrixReadOnly.toFloat32Array" exists.
func (this DOMMatrixReadOnly) HasToFloat32Array() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyToFloat32Array(
		this.Ref(),
	)
}

// ToFloat32ArrayFunc returns the method "DOMMatrixReadOnly.toFloat32Array".
func (this DOMMatrixReadOnly) ToFloat32ArrayFunc() (fn js.Func[func() js.TypedArray[float32]]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToFloat32ArrayFunc(
			this.Ref(),
		),
	)
}

// ToFloat32Array calls the method "DOMMatrixReadOnly.toFloat32Array".
func (this DOMMatrixReadOnly) ToFloat32Array() (ret js.TypedArray[float32]) {
	bindings.CallDOMMatrixReadOnlyToFloat32Array(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToFloat32Array calls the method "DOMMatrixReadOnly.toFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToFloat32Array() (ret js.TypedArray[float32], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToFloat32Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToFloat64Array returns true if the method "DOMMatrixReadOnly.toFloat64Array" exists.
func (this DOMMatrixReadOnly) HasToFloat64Array() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyToFloat64Array(
		this.Ref(),
	)
}

// ToFloat64ArrayFunc returns the method "DOMMatrixReadOnly.toFloat64Array".
func (this DOMMatrixReadOnly) ToFloat64ArrayFunc() (fn js.Func[func() js.TypedArray[float64]]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToFloat64ArrayFunc(
			this.Ref(),
		),
	)
}

// ToFloat64Array calls the method "DOMMatrixReadOnly.toFloat64Array".
func (this DOMMatrixReadOnly) ToFloat64Array() (ret js.TypedArray[float64]) {
	bindings.CallDOMMatrixReadOnlyToFloat64Array(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToFloat64Array calls the method "DOMMatrixReadOnly.toFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToFloat64Array() (ret js.TypedArray[float64], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToFloat64Array(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToString returns true if the method "DOMMatrixReadOnly.toString" exists.
func (this DOMMatrixReadOnly) HasToString() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "DOMMatrixReadOnly.toString".
func (this DOMMatrixReadOnly) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "DOMMatrixReadOnly.toString".
func (this DOMMatrixReadOnly) ToString() (ret js.String) {
	bindings.CallDOMMatrixReadOnlyToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "DOMMatrixReadOnly.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "DOMMatrixReadOnly.toJSON" exists.
func (this DOMMatrixReadOnly) HasToJSON() bool {
	return js.True == bindings.HasDOMMatrixReadOnlyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "DOMMatrixReadOnly.toJSON".
func (this DOMMatrixReadOnly) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.DOMMatrixReadOnlyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "DOMMatrixReadOnly.toJSON".
func (this DOMMatrixReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMMatrixReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMMatrixReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewCSSMatrixComponent(matrix DOMMatrixReadOnly, options CSSMatrixComponentOptions) (ret CSSMatrixComponent) {
	ret.ref = bindings.NewCSSMatrixComponentByCSSMatrixComponent(
		matrix.Ref(),
		js.Pointer(&options))
	return
}

func NewCSSMatrixComponentByCSSMatrixComponent1(matrix DOMMatrixReadOnly) (ret CSSMatrixComponent) {
	ret.ref = bindings.NewCSSMatrixComponentByCSSMatrixComponent1(
		matrix.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSMatrixComponent) Matrix() (ret DOMMatrix, ok bool) {
	ok = js.True == bindings.GetCSSMatrixComponentMatrix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMatrix sets the value of property "CSSMatrixComponent.matrix" to val.
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
// It returns ok=false if there is no such property.
func (this CSSMediaRule) Media() (ret MediaList, ok bool) {
	ok = js.True == bindings.GetCSSMediaRuleMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this CSSNamespaceRule) NamespaceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSNamespaceRuleNamespaceURI(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Prefix returns the value of property "CSSNamespaceRule.prefix".
//
// It returns ok=false if there is no such property.
func (this CSSNamespaceRule) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSNamespaceRulePrefix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSOKLCH(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle, alpha CSSColorPercent) (ret CSSOKLCH) {
	ret.ref = bindings.NewCSSOKLCHByCSSOKLCH(
		l.Ref(),
		c.Ref(),
		h.Ref(),
		alpha.Ref())
	return
}

func NewCSSOKLCHByCSSOKLCH1(l CSSColorPercent, c CSSColorPercent, h CSSColorAngle) (ret CSSOKLCH) {
	ret.ref = bindings.NewCSSOKLCHByCSSOKLCH1(
		l.Ref(),
		c.Ref(),
		h.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSOKLCH) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSOKLCH.l" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLCH) C() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHC(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "CSSOKLCH.c" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLCH) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHH(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSOKLCH.h" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLCH) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSOKLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSOKLab(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber, alpha CSSColorPercent) (ret CSSOKLab) {
	ret.ref = bindings.NewCSSOKLabByCSSOKLab(
		l.Ref(),
		a.Ref(),
		b.Ref(),
		alpha.Ref())
	return
}

func NewCSSOKLabByCSSOKLab1(l CSSColorPercent, a CSSColorNumber, b CSSColorNumber) (ret CSSOKLab) {
	ret.ref = bindings.NewCSSOKLabByCSSOKLab1(
		l.Ref(),
		a.Ref(),
		b.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSOKLab) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLabL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSOKLab.l" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLab) A() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSOKLabA(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "CSSOKLab.a" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLab) B() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSOKLabB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSOKLab.b" to val.
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
// It returns ok=false if there is no such property.
func (this CSSOKLab) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLabAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSOKLab.alpha" to val.
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
// It returns ok=false if there is no such property.
func (this CSSPageRule) SelectorText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPageRuleSelectorText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectorText sets the value of property "CSSPageRule.selectorText" to val.
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
// It returns ok=false if there is no such property.
func (this CSSPageRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSPageRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSParserAtRule(name js.String, prelude js.Array[CSSToken], body js.Array[CSSParserRule]) (ret CSSParserAtRule) {
	ret.ref = bindings.NewCSSParserAtRuleByCSSParserAtRule(
		name.Ref(),
		prelude.Ref(),
		body.Ref())
	return
}

func NewCSSParserAtRuleByCSSParserAtRule1(name js.String, prelude js.Array[CSSToken]) (ret CSSParserAtRule) {
	ret.ref = bindings.NewCSSParserAtRuleByCSSParserAtRule1(
		name.Ref(),
		prelude.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserAtRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Prelude returns the value of property "CSSParserAtRule.prelude".
//
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Prelude() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserAtRulePrelude(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserAtRule.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Body() (ret js.FrozenArray[CSSParserRule], ok bool) {
	ok = js.True == bindings.GetCSSParserAtRuleBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToString returns true if the method "CSSParserAtRule.toString" exists.
func (this CSSParserAtRule) HasToString() bool {
	return js.True == bindings.HasCSSParserAtRuleToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSParserAtRule.toString".
func (this CSSParserAtRule) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserAtRuleToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSParserAtRule.toString".
func (this CSSParserAtRule) ToString() (ret js.String) {
	bindings.CallCSSParserAtRuleToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserAtRule.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserAtRule) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserAtRuleToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewCSSParserBlock(name js.String, body js.Array[CSSParserValue]) (ret CSSParserBlock) {
	ret.ref = bindings.NewCSSParserBlockByCSSParserBlock(
		name.Ref(),
		body.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSParserBlock) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserBlockName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserBlock.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserBlock) Body() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserBlockBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToString returns true if the method "CSSParserBlock.toString" exists.
func (this CSSParserBlock) HasToString() bool {
	return js.True == bindings.HasCSSParserBlockToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSParserBlock.toString".
func (this CSSParserBlock) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserBlockToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSParserBlock.toString".
func (this CSSParserBlock) ToString() (ret js.String) {
	bindings.CallCSSParserBlockToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserBlock.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserBlock) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserBlockToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewCSSParserFunction(name js.String, args js.Array[js.Array[CSSParserValue]]) (ret CSSParserFunction) {
	ret.ref = bindings.NewCSSParserFunctionByCSSParserFunction(
		name.Ref(),
		args.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSParserFunction) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserFunctionName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Args returns the value of property "CSSParserFunction.args".
//
// It returns ok=false if there is no such property.
func (this CSSParserFunction) Args() (ret js.FrozenArray[js.FrozenArray[CSSParserValue]], ok bool) {
	ok = js.True == bindings.GetCSSParserFunctionArgs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToString returns true if the method "CSSParserFunction.toString" exists.
func (this CSSParserFunction) HasToString() bool {
	return js.True == bindings.HasCSSParserFunctionToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSParserFunction.toString".
func (this CSSParserFunction) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserFunctionToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSParserFunction.toString".
func (this CSSParserFunction) ToString() (ret js.String) {
	bindings.CallCSSParserFunctionToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserFunction.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserFunction) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserFunctionToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewCSSParserQualifiedRule(prelude js.Array[CSSToken], body js.Array[CSSParserRule]) (ret CSSParserQualifiedRule) {
	ret.ref = bindings.NewCSSParserQualifiedRuleByCSSParserQualifiedRule(
		prelude.Ref(),
		body.Ref())
	return
}

func NewCSSParserQualifiedRuleByCSSParserQualifiedRule1(prelude js.Array[CSSToken]) (ret CSSParserQualifiedRule) {
	ret.ref = bindings.NewCSSParserQualifiedRuleByCSSParserQualifiedRule1(
		prelude.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this CSSParserQualifiedRule) Prelude() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserQualifiedRulePrelude(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserQualifiedRule.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserQualifiedRule) Body() (ret js.FrozenArray[CSSParserRule], ok bool) {
	ok = js.True == bindings.GetCSSParserQualifiedRuleBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToString returns true if the method "CSSParserQualifiedRule.toString" exists.
func (this CSSParserQualifiedRule) HasToString() bool {
	return js.True == bindings.HasCSSParserQualifiedRuleToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSParserQualifiedRule.toString".
func (this CSSParserQualifiedRule) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSParserQualifiedRuleToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSParserQualifiedRule.toString".
func (this CSSParserQualifiedRule) ToString() (ret js.String) {
	bindings.CallCSSParserQualifiedRuleToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserQualifiedRule.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserQualifiedRule) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserQualifiedRuleToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
