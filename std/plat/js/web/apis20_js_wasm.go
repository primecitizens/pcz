// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// ColorSpace returns the value of property "CSSColor.colorSpace".
//
// It returns ok=false if there is no such property.
func (this CSSColor) ColorSpace() (ret CSSKeywordish, ok bool) {
	ok = js.True == bindings.GetCSSColorColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColorSpace sets the value of property "CSSColor.colorSpace" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetColorSpace(val CSSKeywordish) bool {
	return js.True == bindings.SetCSSColorColorSpace(
		this.ref,
		val.Ref(),
	)
}

// Channels returns the value of property "CSSColor.channels".
//
// It returns ok=false if there is no such property.
func (this CSSColor) Channels() (ret js.ObservableArray[CSSColorPercent], ok bool) {
	ok = js.True == bindings.GetCSSColorChannels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChannels sets the value of property "CSSColor.channels" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetChannels(val js.ObservableArray[CSSColorPercent]) bool {
	return js.True == bindings.SetCSSColorChannels(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSColor.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSColor) Alpha() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSColorAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSColor.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSColor) SetAlpha(val CSSNumberish) bool {
	return js.True == bindings.SetCSSColorAlpha(
		this.ref,
		val.Ref(),
	)
}

type CSSColorAngle = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue

type CSSColorNumber = OneOf_Float64_CSSNumericValue_String_CSSKeywordValue

type CSSColorProfileRule struct {
	CSSRule
}

func (this CSSColorProfileRule) Once() CSSColorProfileRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSColorProfileRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Src returns the value of property "CSSColorProfileRule.src".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RenderingIntent returns the value of property "CSSColorProfileRule.renderingIntent".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) RenderingIntent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleRenderingIntent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Components returns the value of property "CSSColorProfileRule.components".
//
// It returns ok=false if there is no such property.
func (this CSSColorProfileRule) Components() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSColorProfileRuleComponents(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncParse returns true if the static method "CSSColorValue.parse" exists.
func (this CSSColorValue) HasFuncParse() bool {
	return js.True == bindings.HasFuncCSSColorValueParse(
		this.ref,
	)
}

// FuncParse returns the static method "CSSColorValue.parse".
func (this CSSColorValue) FuncParse() (fn js.Func[func(cssText js.String) OneOf_CSSColorValue_CSSStyleValue]) {
	bindings.FuncCSSColorValueParse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Parse calls the static method "CSSColorValue.parse".
func (this CSSColorValue) Parse(cssText js.String) (ret OneOf_CSSColorValue_CSSStyleValue) {
	bindings.CallCSSColorValueParse(
		this.ref, js.Pointer(&ret),
		cssText.Ref(),
	)

	return
}

// TryParse calls the static method "CSSColorValue.parse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSColorValue) TryParse(cssText js.String) (ret OneOf_CSSColorValue_CSSStyleValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSColorValueParse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		cssText.Ref(),
	)

	return
}

type CSSConditionRule struct {
	CSSGroupingRule
}

func (this CSSConditionRule) Once() CSSConditionRule {
	this.ref.Once()
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
	this.ref.Free()
}

// ConditionText returns the value of property "CSSConditionRule.conditionText".
//
// It returns ok=false if there is no such property.
func (this CSSConditionRule) ConditionText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSConditionRuleConditionText(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSContainerRule struct {
	CSSConditionRule
}

func (this CSSContainerRule) Once() CSSContainerRule {
	this.ref.Once()
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
	this.ref.Free()
}

// ContainerName returns the value of property "CSSContainerRule.containerName".
//
// It returns ok=false if there is no such property.
func (this CSSContainerRule) ContainerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSContainerRuleContainerName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContainerQuery returns the value of property "CSSContainerRule.containerQuery".
//
// It returns ok=false if there is no such property.
func (this CSSContainerRule) ContainerQuery() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSContainerRuleContainerQuery(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSCounterStyleRule struct {
	CSSRule
}

func (this CSSCounterStyleRule) Once() CSSCounterStyleRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSCounterStyleRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "CSSCounterStyleRule.name" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetName(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleName(
		this.ref,
		val.Ref(),
	)
}

// System returns the value of property "CSSCounterStyleRule.system".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) System() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSystem(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSystem sets the value of property "CSSCounterStyleRule.system" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSystem(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSystem(
		this.ref,
		val.Ref(),
	)
}

// Symbols returns the value of property "CSSCounterStyleRule.symbols".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Symbols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSymbols(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSymbols sets the value of property "CSSCounterStyleRule.symbols" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSymbols(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSymbols(
		this.ref,
		val.Ref(),
	)
}

// AdditiveSymbols returns the value of property "CSSCounterStyleRule.additiveSymbols".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) AdditiveSymbols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleAdditiveSymbols(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAdditiveSymbols sets the value of property "CSSCounterStyleRule.additiveSymbols" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetAdditiveSymbols(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleAdditiveSymbols(
		this.ref,
		val.Ref(),
	)
}

// Negative returns the value of property "CSSCounterStyleRule.negative".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Negative() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleNegative(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNegative sets the value of property "CSSCounterStyleRule.negative" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetNegative(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleNegative(
		this.ref,
		val.Ref(),
	)
}

// Prefix returns the value of property "CSSCounterStyleRule.prefix".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRulePrefix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPrefix sets the value of property "CSSCounterStyleRule.prefix" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetPrefix(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRulePrefix(
		this.ref,
		val.Ref(),
	)
}

// Suffix returns the value of property "CSSCounterStyleRule.suffix".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Suffix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSuffix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSuffix sets the value of property "CSSCounterStyleRule.suffix" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSuffix(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSuffix(
		this.ref,
		val.Ref(),
	)
}

// Range returns the value of property "CSSCounterStyleRule.range".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Range() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRange sets the value of property "CSSCounterStyleRule.range" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetRange(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleRange(
		this.ref,
		val.Ref(),
	)
}

// Pad returns the value of property "CSSCounterStyleRule.pad".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Pad() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRulePad(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPad sets the value of property "CSSCounterStyleRule.pad" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetPad(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRulePad(
		this.ref,
		val.Ref(),
	)
}

// SpeakAs returns the value of property "CSSCounterStyleRule.speakAs".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) SpeakAs() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleSpeakAs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpeakAs sets the value of property "CSSCounterStyleRule.speakAs" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetSpeakAs(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleSpeakAs(
		this.ref,
		val.Ref(),
	)
}

// Fallback returns the value of property "CSSCounterStyleRule.fallback".
//
// It returns ok=false if there is no such property.
func (this CSSCounterStyleRule) Fallback() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSCounterStyleRuleFallback(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFallback sets the value of property "CSSCounterStyleRule.fallback" to val.
//
// It returns false if the property cannot be set.
func (this CSSCounterStyleRule) SetFallback(val js.String) bool {
	return js.True == bindings.SetCSSCounterStyleRuleFallback(
		this.ref,
		val.Ref(),
	)
}

type CSSFontFaceRule struct {
	CSSRule
}

func (this CSSFontFaceRule) Once() CSSFontFaceRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Style returns the value of property "CSSFontFaceRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSFontFaceRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSFontFaceRuleStyle(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSet returns true if the method "CSSFontFeatureValuesMap.set" exists.
func (this CSSFontFeatureValuesMap) HasFuncSet() bool {
	return js.True == bindings.HasFuncCSSFontFeatureValuesMapSet(
		this.ref,
	)
}

// FuncSet returns the method "CSSFontFeatureValuesMap.set".
func (this CSSFontFeatureValuesMap) FuncSet() (fn js.Func[func(featureValueName js.String, values OneOf_Uint32_ArrayUint32)]) {
	bindings.FuncCSSFontFeatureValuesMapSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "CSSFontFeatureValuesMap.set".
func (this CSSFontFeatureValuesMap) Set(featureValueName js.String, values OneOf_Uint32_ArrayUint32) (ret js.Void) {
	bindings.CallCSSFontFeatureValuesMapSet(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		featureValueName.Ref(),
		values.Ref(),
	)

	return
}

type CSSFontFeatureValuesRule struct {
	CSSRule
}

func (this CSSFontFeatureValuesRule) Once() CSSFontFeatureValuesRule {
	this.ref.Once()
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
	this.ref.Free()
}

// FontFamily returns the value of property "CSSFontFeatureValuesRule.fontFamily".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) FontFamily() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleFontFamily(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontFamily sets the value of property "CSSFontFeatureValuesRule.fontFamily" to val.
//
// It returns false if the property cannot be set.
func (this CSSFontFeatureValuesRule) SetFontFamily(val js.String) bool {
	return js.True == bindings.SetCSSFontFeatureValuesRuleFontFamily(
		this.ref,
		val.Ref(),
	)
}

// Annotation returns the value of property "CSSFontFeatureValuesRule.annotation".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Annotation() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleAnnotation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ornaments returns the value of property "CSSFontFeatureValuesRule.ornaments".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Ornaments() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleOrnaments(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Stylistic returns the value of property "CSSFontFeatureValuesRule.stylistic".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Stylistic() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleStylistic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Swash returns the value of property "CSSFontFeatureValuesRule.swash".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Swash() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleSwash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CharacterVariant returns the value of property "CSSFontFeatureValuesRule.characterVariant".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) CharacterVariant() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleCharacterVariant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Styleset returns the value of property "CSSFontFeatureValuesRule.styleset".
//
// It returns ok=false if there is no such property.
func (this CSSFontFeatureValuesRule) Styleset() (ret CSSFontFeatureValuesMap, ok bool) {
	ok = js.True == bindings.GetCSSFontFeatureValuesRuleStyleset(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSFontPaletteValuesRule struct {
	CSSRule
}

func (this CSSFontPaletteValuesRule) Once() CSSFontPaletteValuesRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSFontPaletteValuesRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FontFamily returns the value of property "CSSFontPaletteValuesRule.fontFamily".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) FontFamily() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleFontFamily(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BasePalette returns the value of property "CSSFontPaletteValuesRule.basePalette".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) BasePalette() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleBasePalette(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OverrideColors returns the value of property "CSSFontPaletteValuesRule.overrideColors".
//
// It returns ok=false if there is no such property.
func (this CSSFontPaletteValuesRule) OverrideColors() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSFontPaletteValuesRuleOverrideColors(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSGroupingRule struct {
	CSSRule
}

func (this CSSGroupingRule) Once() CSSGroupingRule {
	this.ref.Once()
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
	this.ref.Free()
}

// CssRules returns the value of property "CSSGroupingRule.cssRules".
//
// It returns ok=false if there is no such property.
func (this CSSGroupingRule) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSGroupingRuleCssRules(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInsertRule returns true if the method "CSSGroupingRule.insertRule" exists.
func (this CSSGroupingRule) HasFuncInsertRule() bool {
	return js.True == bindings.HasFuncCSSGroupingRuleInsertRule(
		this.ref,
	)
}

// FuncInsertRule returns the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) FuncInsertRule() (fn js.Func[func(rule js.String, index uint32) uint32]) {
	bindings.FuncCSSGroupingRuleInsertRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRule calls the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRule(rule js.String, index uint32) (ret uint32) {
	bindings.CallCSSGroupingRuleInsertRule(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
		uint32(index),
	)

	return
}

// HasFuncInsertRule1 returns true if the method "CSSGroupingRule.insertRule" exists.
func (this CSSGroupingRule) HasFuncInsertRule1() bool {
	return js.True == bindings.HasFuncCSSGroupingRuleInsertRule1(
		this.ref,
	)
}

// FuncInsertRule1 returns the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) FuncInsertRule1() (fn js.Func[func(rule js.String) uint32]) {
	bindings.FuncCSSGroupingRuleInsertRule1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRule1 calls the method "CSSGroupingRule.insertRule".
func (this CSSGroupingRule) InsertRule1(rule js.String) (ret uint32) {
	bindings.CallCSSGroupingRuleInsertRule1(
		this.ref, js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryInsertRule1 calls the method "CSSGroupingRule.insertRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSGroupingRule) TryInsertRule1(rule js.String) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGroupingRuleInsertRule1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasFuncDeleteRule returns true if the method "CSSGroupingRule.deleteRule" exists.
func (this CSSGroupingRule) HasFuncDeleteRule() bool {
	return js.True == bindings.HasFuncCSSGroupingRuleDeleteRule(
		this.ref,
	)
}

// FuncDeleteRule returns the method "CSSGroupingRule.deleteRule".
func (this CSSGroupingRule) FuncDeleteRule() (fn js.Func[func(index uint32)]) {
	bindings.FuncCSSGroupingRuleDeleteRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRule calls the method "CSSGroupingRule.deleteRule".
func (this CSSGroupingRule) DeleteRule(index uint32) (ret js.Void) {
	bindings.CallCSSGroupingRuleDeleteRule(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDeleteRule calls the method "CSSGroupingRule.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSGroupingRule) TryDeleteRule(index uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSGroupingRuleDeleteRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// H returns the value of property "CSSHSL.h".
//
// It returns ok=false if there is no such property.
func (this CSSHSL) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSHSLH(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSHSL.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSHSLH(
		this.ref,
		val.Ref(),
	)
}

// S returns the value of property "CSSHSL.s".
//
// It returns ok=false if there is no such property.
func (this CSSHSL) S() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLS(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetS sets the value of property "CSSHSL.s" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetS(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLS(
		this.ref,
		val.Ref(),
	)
}

// L returns the value of property "CSSHSL.l".
//
// It returns ok=false if there is no such property.
func (this CSSHSL) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSHSL.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLL(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSHSL.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSHSL) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSHSLAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSHSL.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSHSL) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSHSLAlpha(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// H returns the value of property "CSSHWB.h".
//
// It returns ok=false if there is no such property.
func (this CSSHWB) H() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSHWBH(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSHWB.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetH(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSHWBH(
		this.ref,
		val.Ref(),
	)
}

// W returns the value of property "CSSHWB.w".
//
// It returns ok=false if there is no such property.
func (this CSSHWB) W() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBW(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetW sets the value of property "CSSHWB.w" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetW(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBW(
		this.ref,
		val.Ref(),
	)
}

// B returns the value of property "CSSHWB.b".
//
// It returns ok=false if there is no such property.
func (this CSSHWB) B() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSHWB.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetB(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBB(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSHWB.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSHWB) Alpha() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSHWBAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSHWB.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSHWB) SetAlpha(val CSSNumberish) bool {
	return js.True == bindings.SetCSSHWBAlpha(
		this.ref,
		val.Ref(),
	)
}

type CSSImageValue struct {
	CSSStyleValue
}

func (this CSSImageValue) Once() CSSImageValue {
	this.ref.Once()
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
	this.ref.Free()
}

type CSSImportRule struct {
	CSSRule
}

func (this CSSImportRule) Once() CSSImportRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "CSSImportRule.href".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Media returns the value of property "CSSImportRule.media".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) Media() (ret MediaList, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StyleSheet returns the value of property "CSSImportRule.styleSheet".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) StyleSheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleStyleSheet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LayerName returns the value of property "CSSImportRule.layerName".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) LayerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleLayerName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SupportsText returns the value of property "CSSImportRule.supportsText".
//
// It returns ok=false if there is no such property.
func (this CSSImportRule) SupportsText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSImportRuleSupportsText(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSKeyframeRule struct {
	CSSRule
}

func (this CSSKeyframeRule) Once() CSSKeyframeRule {
	this.ref.Once()
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
	this.ref.Free()
}

// KeyText returns the value of property "CSSKeyframeRule.keyText".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframeRule) KeyText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeyframeRuleKeyText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetKeyText sets the value of property "CSSKeyframeRule.keyText" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeyframeRule) SetKeyText(val js.String) bool {
	return js.True == bindings.SetCSSKeyframeRuleKeyText(
		this.ref,
		val.Ref(),
	)
}

// Style returns the value of property "CSSKeyframeRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframeRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSKeyframeRuleStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSKeyframesRule struct {
	CSSRule
}

func (this CSSKeyframesRule) Once() CSSKeyframesRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSKeyframesRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "CSSKeyframesRule.name" to val.
//
// It returns false if the property cannot be set.
func (this CSSKeyframesRule) SetName(val js.String) bool {
	return js.True == bindings.SetCSSKeyframesRuleName(
		this.ref,
		val.Ref(),
	)
}

// CssRules returns the value of property "CSSKeyframesRule.cssRules".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) CssRules() (ret CSSRuleList, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleCssRules(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "CSSKeyframesRule.length".
//
// It returns ok=false if there is no such property.
func (this CSSKeyframesRule) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSKeyframesRuleLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "CSSKeyframesRule." exists.
func (this CSSKeyframesRule) HasFuncGet() bool {
	return js.True == bindings.HasFuncCSSKeyframesRuleGet(
		this.ref,
	)
}

// FuncGet returns the method "CSSKeyframesRule.".
func (this CSSKeyframesRule) FuncGet() (fn js.Func[func(index uint32) CSSKeyframeRule]) {
	bindings.FuncCSSKeyframesRuleGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "CSSKeyframesRule.".
func (this CSSKeyframesRule) Get(index uint32) (ret CSSKeyframeRule) {
	bindings.CallCSSKeyframesRuleGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "CSSKeyframesRule."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryGet(index uint32) (ret CSSKeyframeRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAppendRule returns true if the method "CSSKeyframesRule.appendRule" exists.
func (this CSSKeyframesRule) HasFuncAppendRule() bool {
	return js.True == bindings.HasFuncCSSKeyframesRuleAppendRule(
		this.ref,
	)
}

// FuncAppendRule returns the method "CSSKeyframesRule.appendRule".
func (this CSSKeyframesRule) FuncAppendRule() (fn js.Func[func(rule js.String)]) {
	bindings.FuncCSSKeyframesRuleAppendRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendRule calls the method "CSSKeyframesRule.appendRule".
func (this CSSKeyframesRule) AppendRule(rule js.String) (ret js.Void) {
	bindings.CallCSSKeyframesRuleAppendRule(
		this.ref, js.Pointer(&ret),
		rule.Ref(),
	)

	return
}

// TryAppendRule calls the method "CSSKeyframesRule.appendRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryAppendRule(rule js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleAppendRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rule.Ref(),
	)

	return
}

// HasFuncDeleteRule returns true if the method "CSSKeyframesRule.deleteRule" exists.
func (this CSSKeyframesRule) HasFuncDeleteRule() bool {
	return js.True == bindings.HasFuncCSSKeyframesRuleDeleteRule(
		this.ref,
	)
}

// FuncDeleteRule returns the method "CSSKeyframesRule.deleteRule".
func (this CSSKeyframesRule) FuncDeleteRule() (fn js.Func[func(sel js.String)]) {
	bindings.FuncCSSKeyframesRuleDeleteRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRule calls the method "CSSKeyframesRule.deleteRule".
func (this CSSKeyframesRule) DeleteRule(sel js.String) (ret js.Void) {
	bindings.CallCSSKeyframesRuleDeleteRule(
		this.ref, js.Pointer(&ret),
		sel.Ref(),
	)

	return
}

// TryDeleteRule calls the method "CSSKeyframesRule.deleteRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryDeleteRule(sel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleDeleteRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sel.Ref(),
	)

	return
}

// HasFuncFindRule returns true if the method "CSSKeyframesRule.findRule" exists.
func (this CSSKeyframesRule) HasFuncFindRule() bool {
	return js.True == bindings.HasFuncCSSKeyframesRuleFindRule(
		this.ref,
	)
}

// FuncFindRule returns the method "CSSKeyframesRule.findRule".
func (this CSSKeyframesRule) FuncFindRule() (fn js.Func[func(sel js.String) CSSKeyframeRule]) {
	bindings.FuncCSSKeyframesRuleFindRule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FindRule calls the method "CSSKeyframesRule.findRule".
func (this CSSKeyframesRule) FindRule(sel js.String) (ret CSSKeyframeRule) {
	bindings.CallCSSKeyframesRuleFindRule(
		this.ref, js.Pointer(&ret),
		sel.Ref(),
	)

	return
}

// TryFindRule calls the method "CSSKeyframesRule.findRule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSKeyframesRule) TryFindRule(sel js.String) (ret CSSKeyframeRule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSKeyframesRuleFindRule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// L returns the value of property "CSSLCH.l".
//
// It returns ok=false if there is no such property.
func (this CSSLCH) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSLCH.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHL(
		this.ref,
		val.Ref(),
	)
}

// C returns the value of property "CSSLCH.c".
//
// It returns ok=false if there is no such property.
func (this CSSLCH) C() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHC(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "CSSLCH.c" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetC(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHC(
		this.ref,
		val.Ref(),
	)
}

// H returns the value of property "CSSLCH.h".
//
// It returns ok=false if there is no such property.
func (this CSSLCH) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSLCHH(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSLCH.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSLCHH(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSLCH.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSLCH) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLCHAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLCHAlpha(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// L returns the value of property "CSSLab.l".
//
// It returns ok=false if there is no such property.
func (this CSSLab) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLabL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSLab.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLabL(
		this.ref,
		val.Ref(),
	)
}

// A returns the value of property "CSSLab.a".
//
// It returns ok=false if there is no such property.
func (this CSSLab) A() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSLabA(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "CSSLab.a" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetA(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSLabA(
		this.ref,
		val.Ref(),
	)
}

// B returns the value of property "CSSLab.b".
//
// It returns ok=false if there is no such property.
func (this CSSLab) B() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSLabB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSLab.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetB(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSLabB(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSLab.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSLab) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSLabAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSLab.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSLab) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSLabAlpha(
		this.ref,
		val.Ref(),
	)
}

type CSSLayerBlockRule struct {
	CSSGroupingRule
}

func (this CSSLayerBlockRule) Once() CSSLayerBlockRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSLayerBlockRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSLayerBlockRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSLayerBlockRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSLayerStatementRule struct {
	CSSRule
}

func (this CSSLayerStatementRule) Once() CSSLayerStatementRule {
	this.ref.Once()
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
	this.ref.Free()
}

// NameList returns the value of property "CSSLayerStatementRule.nameList".
//
// It returns ok=false if there is no such property.
func (this CSSLayerStatementRule) NameList() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetCSSLayerStatementRuleNameList(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSMarginRule struct {
	CSSRule
}

func (this CSSMarginRule) Once() CSSMarginRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSMarginRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSMarginRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSMarginRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "CSSMarginRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSMarginRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSMarginRuleStyle(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Lower returns the value of property "CSSMathClamp.lower".
//
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Lower() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampLower(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "CSSMathClamp.value".
//
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Upper returns the value of property "CSSMathClamp.upper".
//
// It returns ok=false if there is no such property.
func (this CSSMathClamp) Upper() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathClampUpper(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "CSSMathInvert.value".
//
// It returns ok=false if there is no such property.
func (this CSSMathInvert) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathInvertValue(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Values returns the value of property "CSSMathMax.values".
//
// It returns ok=false if there is no such property.
func (this CSSMathMax) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathMaxValues(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Values returns the value of property "CSSMathMin.values".
//
// It returns ok=false if there is no such property.
func (this CSSMathMin) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathMinValues(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "CSSMathNegate.value".
//
// It returns ok=false if there is no such property.
func (this CSSMathNegate) Value() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSMathNegateValue(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Values returns the value of property "CSSMathProduct.values".
//
// It returns ok=false if there is no such property.
func (this CSSMathProduct) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathProductValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSMathValue struct {
	CSSNumericValue
}

func (this CSSMathValue) Once() CSSMathValue {
	this.ref.Once()
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
	this.ref.Free()
}

// Operator returns the value of property "CSSMathValue.operator".
//
// It returns ok=false if there is no such property.
func (this CSSMathValue) Operator() (ret CSSMathOperator, ok bool) {
	ok = js.True == bindings.GetCSSMathValueOperator(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// A returns the value of property "DOMMatrixReadOnly.a".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) A() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyA(
		this.ref, js.Pointer(&ret),
	)
	return
}

// B returns the value of property "DOMMatrixReadOnly.b".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) B() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// C returns the value of property "DOMMatrixReadOnly.c".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) C() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyC(
		this.ref, js.Pointer(&ret),
	)
	return
}

// D returns the value of property "DOMMatrixReadOnly.d".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) D() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyD(
		this.ref, js.Pointer(&ret),
	)
	return
}

// E returns the value of property "DOMMatrixReadOnly.e".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) E() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyE(
		this.ref, js.Pointer(&ret),
	)
	return
}

// F returns the value of property "DOMMatrixReadOnly.f".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) F() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyF(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M11 returns the value of property "DOMMatrixReadOnly.m11".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M11() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM11(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M12 returns the value of property "DOMMatrixReadOnly.m12".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M12() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM12(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M13 returns the value of property "DOMMatrixReadOnly.m13".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M13() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM13(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M14 returns the value of property "DOMMatrixReadOnly.m14".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M14() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM14(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M21 returns the value of property "DOMMatrixReadOnly.m21".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M21() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM21(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M22 returns the value of property "DOMMatrixReadOnly.m22".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M22() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM22(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M23 returns the value of property "DOMMatrixReadOnly.m23".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M23() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM23(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M24 returns the value of property "DOMMatrixReadOnly.m24".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M24() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM24(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M31 returns the value of property "DOMMatrixReadOnly.m31".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M31() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM31(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M32 returns the value of property "DOMMatrixReadOnly.m32".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M32() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM32(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M33 returns the value of property "DOMMatrixReadOnly.m33".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M33() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM33(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M34 returns the value of property "DOMMatrixReadOnly.m34".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M34() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM34(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M41 returns the value of property "DOMMatrixReadOnly.m41".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M41() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM41(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M42 returns the value of property "DOMMatrixReadOnly.m42".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M42() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM42(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M43 returns the value of property "DOMMatrixReadOnly.m43".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M43() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM43(
		this.ref, js.Pointer(&ret),
	)
	return
}

// M44 returns the value of property "DOMMatrixReadOnly.m44".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) M44() (ret float64, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyM44(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Is2D returns the value of property "DOMMatrixReadOnly.is2D".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) Is2D() (ret bool, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyIs2D(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsIdentity returns the value of property "DOMMatrixReadOnly.isIdentity".
//
// It returns ok=false if there is no such property.
func (this DOMMatrixReadOnly) IsIdentity() (ret bool, ok bool) {
	ok = js.True == bindings.GetDOMMatrixReadOnlyIsIdentity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncFromMatrix returns true if the static method "DOMMatrixReadOnly.fromMatrix" exists.
func (this DOMMatrixReadOnly) HasFuncFromMatrix() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFromMatrix(
		this.ref,
	)
}

// FuncFromMatrix returns the static method "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FuncFromMatrix() (fn js.Func[func(other DOMMatrixInit) DOMMatrixReadOnly]) {
	bindings.FuncDOMMatrixReadOnlyFromMatrix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromMatrix calls the static method "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrix(other DOMMatrixInit) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromMatrix(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryFromMatrix calls the static method "DOMMatrixReadOnly.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromMatrix(other DOMMatrixInit) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromMatrix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncFromMatrix1 returns true if the static method "DOMMatrixReadOnly.fromMatrix" exists.
func (this DOMMatrixReadOnly) HasFuncFromMatrix1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFromMatrix1(
		this.ref,
	)
}

// FuncFromMatrix1 returns the static method "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FuncFromMatrix1() (fn js.Func[func() DOMMatrixReadOnly]) {
	bindings.FuncDOMMatrixReadOnlyFromMatrix1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromMatrix1 calls the static method "DOMMatrixReadOnly.fromMatrix".
func (this DOMMatrixReadOnly) FromMatrix1() (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromMatrix1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFromMatrix1 calls the static method "DOMMatrixReadOnly.fromMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromMatrix1() (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromMatrix1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFromFloat32Array returns true if the static method "DOMMatrixReadOnly.fromFloat32Array" exists.
func (this DOMMatrixReadOnly) HasFuncFromFloat32Array() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFromFloat32Array(
		this.ref,
	)
}

// FuncFromFloat32Array returns the static method "DOMMatrixReadOnly.fromFloat32Array".
func (this DOMMatrixReadOnly) FuncFromFloat32Array() (fn js.Func[func(array32 js.TypedArray[float32]) DOMMatrixReadOnly]) {
	bindings.FuncDOMMatrixReadOnlyFromFloat32Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromFloat32Array calls the static method "DOMMatrixReadOnly.fromFloat32Array".
func (this DOMMatrixReadOnly) FromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromFloat32Array(
		this.ref, js.Pointer(&ret),
		array32.Ref(),
	)

	return
}

// TryFromFloat32Array calls the static method "DOMMatrixReadOnly.fromFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromFloat32Array(array32 js.TypedArray[float32]) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromFloat32Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array32.Ref(),
	)

	return
}

// HasFuncFromFloat64Array returns true if the static method "DOMMatrixReadOnly.fromFloat64Array" exists.
func (this DOMMatrixReadOnly) HasFuncFromFloat64Array() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFromFloat64Array(
		this.ref,
	)
}

// FuncFromFloat64Array returns the static method "DOMMatrixReadOnly.fromFloat64Array".
func (this DOMMatrixReadOnly) FuncFromFloat64Array() (fn js.Func[func(array64 js.TypedArray[float64]) DOMMatrixReadOnly]) {
	bindings.FuncDOMMatrixReadOnlyFromFloat64Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FromFloat64Array calls the static method "DOMMatrixReadOnly.fromFloat64Array".
func (this DOMMatrixReadOnly) FromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrixReadOnly) {
	bindings.CallDOMMatrixReadOnlyFromFloat64Array(
		this.ref, js.Pointer(&ret),
		array64.Ref(),
	)

	return
}

// TryFromFloat64Array calls the static method "DOMMatrixReadOnly.fromFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFromFloat64Array(array64 js.TypedArray[float64]) (ret DOMMatrixReadOnly, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFromFloat64Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array64.Ref(),
	)

	return
}

// HasFuncTranslate returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasFuncTranslate() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTranslate(
		this.ref,
	)
}

// FuncTranslate returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) FuncTranslate() (fn js.Func[func(tx float64, ty float64, tz float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyTranslate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Translate calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate(tx float64, ty float64, tz float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
		float64(tz),
	)

	return
}

// HasFuncTranslate1 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasFuncTranslate1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTranslate1(
		this.ref,
	)
}

// FuncTranslate1 returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) FuncTranslate1() (fn js.Func[func(tx float64, ty float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyTranslate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Translate1 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate1(tx float64, ty float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
		float64(ty),
	)

	return
}

// HasFuncTranslate2 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasFuncTranslate2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTranslate2(
		this.ref,
	)
}

// FuncTranslate2 returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) FuncTranslate2() (fn js.Func[func(tx float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyTranslate2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Translate2 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate2(tx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate2(
		this.ref, js.Pointer(&ret),
		float64(tx),
	)

	return
}

// TryTranslate2 calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate2(tx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(tx),
	)

	return
}

// HasFuncTranslate3 returns true if the method "DOMMatrixReadOnly.translate" exists.
func (this DOMMatrixReadOnly) HasFuncTranslate3() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTranslate3(
		this.ref,
	)
}

// FuncTranslate3 returns the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) FuncTranslate3() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyTranslate3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Translate3 calls the method "DOMMatrixReadOnly.translate".
func (this DOMMatrixReadOnly) Translate3() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyTranslate3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTranslate3 calls the method "DOMMatrixReadOnly.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTranslate3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTranslate3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScale returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale(
		this.ref,
	)
}

// FuncScale returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasFuncScale1 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale1(
		this.ref,
	)
}

// FuncScale1 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale1() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale1 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale1(scaleX float64, scaleY float64, scaleZ float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
		float64(originY),
	)

	return
}

// HasFuncScale2 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale2(
		this.ref,
	)
}

// FuncScale2 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale2() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64, originX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale2 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale2(scaleX float64, scaleY float64, scaleZ float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
		float64(originX),
	)

	return
}

// HasFuncScale3 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale3() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3(
		this.ref,
	)
}

// FuncScale3 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale3() (fn js.Func[func(scaleX float64, scaleY float64, scaleZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale3(scaleX float64, scaleY float64, scaleZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
		float64(scaleZ),
	)

	return
}

// HasFuncScale4 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale4() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale4(
		this.ref,
	)
}

// FuncScale4 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale4() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale4 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale4(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale4(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasFuncScale5 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale5() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale5(
		this.ref,
	)
}

// FuncScale5 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale5() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale5 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale5(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale5(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScale5 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale5(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasFuncScale6 returns true if the method "DOMMatrixReadOnly.scale" exists.
func (this DOMMatrixReadOnly) HasFuncScale6() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale6(
		this.ref,
	)
}

// FuncScale6 returns the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) FuncScale6() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale6 calls the method "DOMMatrixReadOnly.scale".
func (this DOMMatrixReadOnly) Scale6() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale6(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScale6 calls the method "DOMMatrixReadOnly.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale6() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScaleNonUniform returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasFuncScaleNonUniform() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScaleNonUniform(
		this.ref,
	)
}

// FuncScaleNonUniform returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) FuncScaleNonUniform() (fn js.Func[func(scaleX float64, scaleY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScaleNonUniform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleNonUniform calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform(scaleX float64, scaleY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
		float64(scaleY),
	)

	return
}

// HasFuncScaleNonUniform1 returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasFuncScaleNonUniform1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScaleNonUniform1(
		this.ref,
	)
}

// FuncScaleNonUniform1 returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) FuncScaleNonUniform1() (fn js.Func[func(scaleX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScaleNonUniform1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleNonUniform1 calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform1(scaleX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform1(
		this.ref, js.Pointer(&ret),
		float64(scaleX),
	)

	return
}

// TryScaleNonUniform1 calls the method "DOMMatrixReadOnly.scaleNonUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScaleNonUniform1(scaleX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScaleNonUniform1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scaleX),
	)

	return
}

// HasFuncScaleNonUniform2 returns true if the method "DOMMatrixReadOnly.scaleNonUniform" exists.
func (this DOMMatrixReadOnly) HasFuncScaleNonUniform2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScaleNonUniform2(
		this.ref,
	)
}

// FuncScaleNonUniform2 returns the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) FuncScaleNonUniform2() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScaleNonUniform2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScaleNonUniform2 calls the method "DOMMatrixReadOnly.scaleNonUniform".
func (this DOMMatrixReadOnly) ScaleNonUniform2() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScaleNonUniform2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScaleNonUniform2 calls the method "DOMMatrixReadOnly.scaleNonUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScaleNonUniform2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScaleNonUniform2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScale3d returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasFuncScale3d() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3d(
		this.ref,
	)
}

// FuncScale3d returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) FuncScale3d() (fn js.Func[func(scale float64, originX float64, originY float64, originZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3d calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d(scale float64, originX float64, originY float64, originZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
		float64(originZ),
	)

	return
}

// HasFuncScale3d1 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasFuncScale3d1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3d1(
		this.ref,
	)
}

// FuncScale3d1 returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) FuncScale3d1() (fn js.Func[func(scale float64, originX float64, originY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3d1 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d1(scale float64, originX float64, originY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
		float64(originY),
	)

	return
}

// HasFuncScale3d2 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasFuncScale3d2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3d2(
		this.ref,
	)
}

// FuncScale3d2 returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) FuncScale3d2() (fn js.Func[func(scale float64, originX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3d2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3d2 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d2(scale float64, originX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
		float64(originX),
	)

	return
}

// HasFuncScale3d3 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasFuncScale3d3() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3d3(
		this.ref,
	)
}

// FuncScale3d3 returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) FuncScale3d3() (fn js.Func[func(scale float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3d3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3d3 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d3(scale float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d3(
		this.ref, js.Pointer(&ret),
		float64(scale),
	)

	return
}

// TryScale3d3 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d3(scale float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(scale),
	)

	return
}

// HasFuncScale3d4 returns true if the method "DOMMatrixReadOnly.scale3d" exists.
func (this DOMMatrixReadOnly) HasFuncScale3d4() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyScale3d4(
		this.ref,
	)
}

// FuncScale3d4 returns the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) FuncScale3d4() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyScale3d4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale3d4 calls the method "DOMMatrixReadOnly.scale3d".
func (this DOMMatrixReadOnly) Scale3d4() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyScale3d4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScale3d4 calls the method "DOMMatrixReadOnly.scale3d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryScale3d4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyScale3d4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotate returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasFuncRotate() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotate(
		this.ref,
	)
}

// FuncRotate returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) FuncRotate() (fn js.Func[func(rotX float64, rotY float64, rotZ float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rotate calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate(rotX float64, rotY float64, rotZ float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
		float64(rotZ),
	)

	return
}

// HasFuncRotate1 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasFuncRotate1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotate1(
		this.ref,
	)
}

// FuncRotate1 returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) FuncRotate1() (fn js.Func[func(rotX float64, rotY float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rotate1 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate1(rotX float64, rotY float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
		float64(rotY),
	)

	return
}

// HasFuncRotate2 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasFuncRotate2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotate2(
		this.ref,
	)
}

// FuncRotate2 returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) FuncRotate2() (fn js.Func[func(rotX float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotate2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rotate2 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate2(rotX float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate2(
		this.ref, js.Pointer(&ret),
		float64(rotX),
	)

	return
}

// TryRotate2 calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate2(rotX float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(rotX),
	)

	return
}

// HasFuncRotate3 returns true if the method "DOMMatrixReadOnly.rotate" exists.
func (this DOMMatrixReadOnly) HasFuncRotate3() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotate3(
		this.ref,
	)
}

// FuncRotate3 returns the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) FuncRotate3() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotate3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rotate3 calls the method "DOMMatrixReadOnly.rotate".
func (this DOMMatrixReadOnly) Rotate3() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotate3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotate3 calls the method "DOMMatrixReadOnly.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotate3() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotate3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotateFromVector returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasFuncRotateFromVector() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateFromVector(
		this.ref,
	)
}

// FuncRotateFromVector returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) FuncRotateFromVector() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateFromVector(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVector calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRotateFromVector1 returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasFuncRotateFromVector1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateFromVector1(
		this.ref,
	)
}

// FuncRotateFromVector1 returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) FuncRotateFromVector1() (fn js.Func[func(x float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateFromVector1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVector1 calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector1(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector1(
		this.ref, js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateFromVector1 calls the method "DOMMatrixReadOnly.rotateFromVector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateFromVector1(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateFromVector1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasFuncRotateFromVector2 returns true if the method "DOMMatrixReadOnly.rotateFromVector" exists.
func (this DOMMatrixReadOnly) HasFuncRotateFromVector2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateFromVector2(
		this.ref,
	)
}

// FuncRotateFromVector2 returns the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) FuncRotateFromVector2() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateFromVector2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateFromVector2 calls the method "DOMMatrixReadOnly.rotateFromVector".
func (this DOMMatrixReadOnly) RotateFromVector2() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateFromVector2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotateFromVector2 calls the method "DOMMatrixReadOnly.rotateFromVector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateFromVector2() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateFromVector2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRotateAxisAngle returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasFuncRotateAxisAngle() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateAxisAngle(
		this.ref,
	)
}

// FuncRotateAxisAngle returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) FuncRotateAxisAngle() (fn js.Func[func(x float64, y float64, z float64, angle float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateAxisAngle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngle calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle(x float64, y float64, z float64, angle float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
		float64(angle),
	)

	return
}

// HasFuncRotateAxisAngle1 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasFuncRotateAxisAngle1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateAxisAngle1(
		this.ref,
	)
}

// FuncRotateAxisAngle1 returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) FuncRotateAxisAngle1() (fn js.Func[func(x float64, y float64, z float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateAxisAngle1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngle1 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle1(x float64, y float64, z float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(z),
	)

	return
}

// HasFuncRotateAxisAngle2 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasFuncRotateAxisAngle2() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateAxisAngle2(
		this.ref,
	)
}

// FuncRotateAxisAngle2 returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) FuncRotateAxisAngle2() (fn js.Func[func(x float64, y float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateAxisAngle2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngle2 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle2(x float64, y float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRotateAxisAngle3 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasFuncRotateAxisAngle3() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateAxisAngle3(
		this.ref,
	)
}

// FuncRotateAxisAngle3 returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) FuncRotateAxisAngle3() (fn js.Func[func(x float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateAxisAngle3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngle3 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle3(x float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle3(
		this.ref, js.Pointer(&ret),
		float64(x),
	)

	return
}

// TryRotateAxisAngle3 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle3(x float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
	)

	return
}

// HasFuncRotateAxisAngle4 returns true if the method "DOMMatrixReadOnly.rotateAxisAngle" exists.
func (this DOMMatrixReadOnly) HasFuncRotateAxisAngle4() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyRotateAxisAngle4(
		this.ref,
	)
}

// FuncRotateAxisAngle4 returns the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) FuncRotateAxisAngle4() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyRotateAxisAngle4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RotateAxisAngle4 calls the method "DOMMatrixReadOnly.rotateAxisAngle".
func (this DOMMatrixReadOnly) RotateAxisAngle4() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyRotateAxisAngle4(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRotateAxisAngle4 calls the method "DOMMatrixReadOnly.rotateAxisAngle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryRotateAxisAngle4() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyRotateAxisAngle4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSkewX returns true if the method "DOMMatrixReadOnly.skewX" exists.
func (this DOMMatrixReadOnly) HasFuncSkewX() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlySkewX(
		this.ref,
	)
}

// FuncSkewX returns the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) FuncSkewX() (fn js.Func[func(sx float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlySkewX(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewX calls the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewX(sx float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewX(
		this.ref, js.Pointer(&ret),
		float64(sx),
	)

	return
}

// TrySkewX calls the method "DOMMatrixReadOnly.skewX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewX(sx float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewX(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(sx),
	)

	return
}

// HasFuncSkewX1 returns true if the method "DOMMatrixReadOnly.skewX" exists.
func (this DOMMatrixReadOnly) HasFuncSkewX1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlySkewX1(
		this.ref,
	)
}

// FuncSkewX1 returns the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) FuncSkewX1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlySkewX1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewX1 calls the method "DOMMatrixReadOnly.skewX".
func (this DOMMatrixReadOnly) SkewX1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewX1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkewX1 calls the method "DOMMatrixReadOnly.skewX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewX1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewX1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSkewY returns true if the method "DOMMatrixReadOnly.skewY" exists.
func (this DOMMatrixReadOnly) HasFuncSkewY() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlySkewY(
		this.ref,
	)
}

// FuncSkewY returns the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) FuncSkewY() (fn js.Func[func(sy float64) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlySkewY(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewY calls the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewY(sy float64) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewY(
		this.ref, js.Pointer(&ret),
		float64(sy),
	)

	return
}

// TrySkewY calls the method "DOMMatrixReadOnly.skewY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewY(sy float64) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewY(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(sy),
	)

	return
}

// HasFuncSkewY1 returns true if the method "DOMMatrixReadOnly.skewY" exists.
func (this DOMMatrixReadOnly) HasFuncSkewY1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlySkewY1(
		this.ref,
	)
}

// FuncSkewY1 returns the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) FuncSkewY1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlySkewY1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkewY1 calls the method "DOMMatrixReadOnly.skewY".
func (this DOMMatrixReadOnly) SkewY1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlySkewY1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkewY1 calls the method "DOMMatrixReadOnly.skewY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TrySkewY1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlySkewY1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMultiply returns true if the method "DOMMatrixReadOnly.multiply" exists.
func (this DOMMatrixReadOnly) HasFuncMultiply() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyMultiply(
		this.ref,
	)
}

// FuncMultiply returns the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) FuncMultiply() (fn js.Func[func(other DOMMatrixInit) DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyMultiply(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Multiply calls the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) Multiply(other DOMMatrixInit) (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyMultiply(
		this.ref, js.Pointer(&ret),
		js.Pointer(&other),
	)

	return
}

// TryMultiply calls the method "DOMMatrixReadOnly.multiply"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryMultiply(other DOMMatrixInit) (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyMultiply(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&other),
	)

	return
}

// HasFuncMultiply1 returns true if the method "DOMMatrixReadOnly.multiply" exists.
func (this DOMMatrixReadOnly) HasFuncMultiply1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyMultiply1(
		this.ref,
	)
}

// FuncMultiply1 returns the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) FuncMultiply1() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyMultiply1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Multiply1 calls the method "DOMMatrixReadOnly.multiply".
func (this DOMMatrixReadOnly) Multiply1() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyMultiply1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMultiply1 calls the method "DOMMatrixReadOnly.multiply"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryMultiply1() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyMultiply1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFlipX returns true if the method "DOMMatrixReadOnly.flipX" exists.
func (this DOMMatrixReadOnly) HasFuncFlipX() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFlipX(
		this.ref,
	)
}

// FuncFlipX returns the method "DOMMatrixReadOnly.flipX".
func (this DOMMatrixReadOnly) FuncFlipX() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyFlipX(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FlipX calls the method "DOMMatrixReadOnly.flipX".
func (this DOMMatrixReadOnly) FlipX() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyFlipX(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlipX calls the method "DOMMatrixReadOnly.flipX"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFlipX() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFlipX(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFlipY returns true if the method "DOMMatrixReadOnly.flipY" exists.
func (this DOMMatrixReadOnly) HasFuncFlipY() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyFlipY(
		this.ref,
	)
}

// FuncFlipY returns the method "DOMMatrixReadOnly.flipY".
func (this DOMMatrixReadOnly) FuncFlipY() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyFlipY(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FlipY calls the method "DOMMatrixReadOnly.flipY".
func (this DOMMatrixReadOnly) FlipY() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyFlipY(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlipY calls the method "DOMMatrixReadOnly.flipY"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryFlipY() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyFlipY(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInverse returns true if the method "DOMMatrixReadOnly.inverse" exists.
func (this DOMMatrixReadOnly) HasFuncInverse() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyInverse(
		this.ref,
	)
}

// FuncInverse returns the method "DOMMatrixReadOnly.inverse".
func (this DOMMatrixReadOnly) FuncInverse() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncDOMMatrixReadOnlyInverse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Inverse calls the method "DOMMatrixReadOnly.inverse".
func (this DOMMatrixReadOnly) Inverse() (ret DOMMatrix) {
	bindings.CallDOMMatrixReadOnlyInverse(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryInverse calls the method "DOMMatrixReadOnly.inverse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryInverse() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyInverse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTransformPoint returns true if the method "DOMMatrixReadOnly.transformPoint" exists.
func (this DOMMatrixReadOnly) HasFuncTransformPoint() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTransformPoint(
		this.ref,
	)
}

// FuncTransformPoint returns the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) FuncTransformPoint() (fn js.Func[func(point DOMPointInit) DOMPoint]) {
	bindings.FuncDOMMatrixReadOnlyTransformPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransformPoint calls the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPoint(point DOMPointInit) (ret DOMPoint) {
	bindings.CallDOMMatrixReadOnlyTransformPoint(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryTransformPoint calls the method "DOMMatrixReadOnly.transformPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTransformPoint(point DOMPointInit) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTransformPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasFuncTransformPoint1 returns true if the method "DOMMatrixReadOnly.transformPoint" exists.
func (this DOMMatrixReadOnly) HasFuncTransformPoint1() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyTransformPoint1(
		this.ref,
	)
}

// FuncTransformPoint1 returns the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) FuncTransformPoint1() (fn js.Func[func() DOMPoint]) {
	bindings.FuncDOMMatrixReadOnlyTransformPoint1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransformPoint1 calls the method "DOMMatrixReadOnly.transformPoint".
func (this DOMMatrixReadOnly) TransformPoint1() (ret DOMPoint) {
	bindings.CallDOMMatrixReadOnlyTransformPoint1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTransformPoint1 calls the method "DOMMatrixReadOnly.transformPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryTransformPoint1() (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyTransformPoint1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToFloat32Array returns true if the method "DOMMatrixReadOnly.toFloat32Array" exists.
func (this DOMMatrixReadOnly) HasFuncToFloat32Array() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyToFloat32Array(
		this.ref,
	)
}

// FuncToFloat32Array returns the method "DOMMatrixReadOnly.toFloat32Array".
func (this DOMMatrixReadOnly) FuncToFloat32Array() (fn js.Func[func() js.TypedArray[float32]]) {
	bindings.FuncDOMMatrixReadOnlyToFloat32Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToFloat32Array calls the method "DOMMatrixReadOnly.toFloat32Array".
func (this DOMMatrixReadOnly) ToFloat32Array() (ret js.TypedArray[float32]) {
	bindings.CallDOMMatrixReadOnlyToFloat32Array(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToFloat32Array calls the method "DOMMatrixReadOnly.toFloat32Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToFloat32Array() (ret js.TypedArray[float32], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToFloat32Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToFloat64Array returns true if the method "DOMMatrixReadOnly.toFloat64Array" exists.
func (this DOMMatrixReadOnly) HasFuncToFloat64Array() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyToFloat64Array(
		this.ref,
	)
}

// FuncToFloat64Array returns the method "DOMMatrixReadOnly.toFloat64Array".
func (this DOMMatrixReadOnly) FuncToFloat64Array() (fn js.Func[func() js.TypedArray[float64]]) {
	bindings.FuncDOMMatrixReadOnlyToFloat64Array(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToFloat64Array calls the method "DOMMatrixReadOnly.toFloat64Array".
func (this DOMMatrixReadOnly) ToFloat64Array() (ret js.TypedArray[float64]) {
	bindings.CallDOMMatrixReadOnlyToFloat64Array(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToFloat64Array calls the method "DOMMatrixReadOnly.toFloat64Array"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToFloat64Array() (ret js.TypedArray[float64], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToFloat64Array(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToString returns true if the method "DOMMatrixReadOnly.toString" exists.
func (this DOMMatrixReadOnly) HasFuncToString() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyToString(
		this.ref,
	)
}

// FuncToString returns the method "DOMMatrixReadOnly.toString".
func (this DOMMatrixReadOnly) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncDOMMatrixReadOnlyToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "DOMMatrixReadOnly.toString".
func (this DOMMatrixReadOnly) ToString() (ret js.String) {
	bindings.CallDOMMatrixReadOnlyToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "DOMMatrixReadOnly.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "DOMMatrixReadOnly.toJSON" exists.
func (this DOMMatrixReadOnly) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncDOMMatrixReadOnlyToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "DOMMatrixReadOnly.toJSON".
func (this DOMMatrixReadOnly) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncDOMMatrixReadOnlyToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "DOMMatrixReadOnly.toJSON".
func (this DOMMatrixReadOnly) ToJSON() (ret js.Object) {
	bindings.CallDOMMatrixReadOnlyToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "DOMMatrixReadOnly.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DOMMatrixReadOnly) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDOMMatrixReadOnlyToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *CSSMatrixComponentOptions) UpdateFrom(ref js.Ref) {
	bindings.CSSMatrixComponentOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CSSMatrixComponentOptions) Update(ref js.Ref) {
	bindings.CSSMatrixComponentOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CSSMatrixComponentOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Matrix returns the value of property "CSSMatrixComponent.matrix".
//
// It returns ok=false if there is no such property.
func (this CSSMatrixComponent) Matrix() (ret DOMMatrix, ok bool) {
	ok = js.True == bindings.GetCSSMatrixComponentMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMatrix sets the value of property "CSSMatrixComponent.matrix" to val.
//
// It returns false if the property cannot be set.
func (this CSSMatrixComponent) SetMatrix(val DOMMatrix) bool {
	return js.True == bindings.SetCSSMatrixComponentMatrix(
		this.ref,
		val.Ref(),
	)
}

type CSSMediaRule struct {
	CSSConditionRule
}

func (this CSSMediaRule) Once() CSSMediaRule {
	this.ref.Once()
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
	this.ref.Free()
}

// Media returns the value of property "CSSMediaRule.media".
//
// It returns ok=false if there is no such property.
func (this CSSMediaRule) Media() (ret MediaList, ok bool) {
	ok = js.True == bindings.GetCSSMediaRuleMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSNamespaceRule struct {
	CSSRule
}

func (this CSSNamespaceRule) Once() CSSNamespaceRule {
	this.ref.Once()
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
	this.ref.Free()
}

// NamespaceURI returns the value of property "CSSNamespaceRule.namespaceURI".
//
// It returns ok=false if there is no such property.
func (this CSSNamespaceRule) NamespaceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSNamespaceRuleNamespaceURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Prefix returns the value of property "CSSNamespaceRule.prefix".
//
// It returns ok=false if there is no such property.
func (this CSSNamespaceRule) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSNamespaceRulePrefix(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// L returns the value of property "CSSOKLCH.l".
//
// It returns ok=false if there is no such property.
func (this CSSOKLCH) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSOKLCH.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHL(
		this.ref,
		val.Ref(),
	)
}

// C returns the value of property "CSSOKLCH.c".
//
// It returns ok=false if there is no such property.
func (this CSSOKLCH) C() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHC(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetC sets the value of property "CSSOKLCH.c" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetC(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHC(
		this.ref,
		val.Ref(),
	)
}

// H returns the value of property "CSSOKLCH.h".
//
// It returns ok=false if there is no such property.
func (this CSSOKLCH) H() (ret CSSColorAngle, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHH(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetH sets the value of property "CSSOKLCH.h" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetH(val CSSColorAngle) bool {
	return js.True == bindings.SetCSSOKLCHH(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSOKLCH.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSOKLCH) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLCHAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSOKLCH.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLCH) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLCHAlpha(
		this.ref,
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
	this.ref.Once()
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
	this.ref.Free()
}

// L returns the value of property "CSSOKLab.l".
//
// It returns ok=false if there is no such property.
func (this CSSOKLab) L() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLabL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetL sets the value of property "CSSOKLab.l" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetL(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLabL(
		this.ref,
		val.Ref(),
	)
}

// A returns the value of property "CSSOKLab.a".
//
// It returns ok=false if there is no such property.
func (this CSSOKLab) A() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSOKLabA(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetA sets the value of property "CSSOKLab.a" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetA(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSOKLabA(
		this.ref,
		val.Ref(),
	)
}

// B returns the value of property "CSSOKLab.b".
//
// It returns ok=false if there is no such property.
func (this CSSOKLab) B() (ret CSSColorNumber, ok bool) {
	ok = js.True == bindings.GetCSSOKLabB(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSOKLab.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetB(val CSSColorNumber) bool {
	return js.True == bindings.SetCSSOKLabB(
		this.ref,
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSOKLab.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSOKLab) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSOKLabAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSOKLab.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSOKLab) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSOKLabAlpha(
		this.ref,
		val.Ref(),
	)
}

type CSSPageRule struct {
	CSSGroupingRule
}

func (this CSSPageRule) Once() CSSPageRule {
	this.ref.Once()
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
	this.ref.Free()
}

// SelectorText returns the value of property "CSSPageRule.selectorText".
//
// It returns ok=false if there is no such property.
func (this CSSPageRule) SelectorText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPageRuleSelectorText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectorText sets the value of property "CSSPageRule.selectorText" to val.
//
// It returns false if the property cannot be set.
func (this CSSPageRule) SetSelectorText(val js.String) bool {
	return js.True == bindings.SetCSSPageRuleSelectorText(
		this.ref,
		val.Ref(),
	)
}

// Style returns the value of property "CSSPageRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSPageRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSPageRuleStyle(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSParserAtRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserAtRuleName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Prelude returns the value of property "CSSParserAtRule.prelude".
//
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Prelude() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserAtRulePrelude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserAtRule.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserAtRule) Body() (ret js.FrozenArray[CSSParserRule], ok bool) {
	ok = js.True == bindings.GetCSSParserAtRuleBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToString returns true if the method "CSSParserAtRule.toString" exists.
func (this CSSParserAtRule) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSParserAtRuleToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSParserAtRule.toString".
func (this CSSParserAtRule) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSParserAtRuleToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSParserAtRule.toString".
func (this CSSParserAtRule) ToString() (ret js.String) {
	bindings.CallCSSParserAtRuleToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserAtRule.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserAtRule) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserAtRuleToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSParserBlock.name".
//
// It returns ok=false if there is no such property.
func (this CSSParserBlock) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserBlockName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserBlock.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserBlock) Body() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserBlockBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToString returns true if the method "CSSParserBlock.toString" exists.
func (this CSSParserBlock) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSParserBlockToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSParserBlock.toString".
func (this CSSParserBlock) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSParserBlockToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSParserBlock.toString".
func (this CSSParserBlock) ToString() (ret js.String) {
	bindings.CallCSSParserBlockToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserBlock.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserBlock) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserBlockToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "CSSParserFunction.name".
//
// It returns ok=false if there is no such property.
func (this CSSParserFunction) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSParserFunctionName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Args returns the value of property "CSSParserFunction.args".
//
// It returns ok=false if there is no such property.
func (this CSSParserFunction) Args() (ret js.FrozenArray[js.FrozenArray[CSSParserValue]], ok bool) {
	ok = js.True == bindings.GetCSSParserFunctionArgs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToString returns true if the method "CSSParserFunction.toString" exists.
func (this CSSParserFunction) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSParserFunctionToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSParserFunction.toString".
func (this CSSParserFunction) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSParserFunctionToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSParserFunction.toString".
func (this CSSParserFunction) ToString() (ret js.String) {
	bindings.CallCSSParserFunctionToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserFunction.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserFunction) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserFunctionToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Prelude returns the value of property "CSSParserQualifiedRule.prelude".
//
// It returns ok=false if there is no such property.
func (this CSSParserQualifiedRule) Prelude() (ret js.FrozenArray[CSSParserValue], ok bool) {
	ok = js.True == bindings.GetCSSParserQualifiedRulePrelude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "CSSParserQualifiedRule.body".
//
// It returns ok=false if there is no such property.
func (this CSSParserQualifiedRule) Body() (ret js.FrozenArray[CSSParserRule], ok bool) {
	ok = js.True == bindings.GetCSSParserQualifiedRuleBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToString returns true if the method "CSSParserQualifiedRule.toString" exists.
func (this CSSParserQualifiedRule) HasFuncToString() bool {
	return js.True == bindings.HasFuncCSSParserQualifiedRuleToString(
		this.ref,
	)
}

// FuncToString returns the method "CSSParserQualifiedRule.toString".
func (this CSSParserQualifiedRule) FuncToString() (fn js.Func[func() js.String]) {
	bindings.FuncCSSParserQualifiedRuleToString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToString calls the method "CSSParserQualifiedRule.toString".
func (this CSSParserQualifiedRule) ToString() (ret js.String) {
	bindings.CallCSSParserQualifiedRuleToString(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSParserQualifiedRule.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSParserQualifiedRule) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSParserQualifiedRuleToString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
