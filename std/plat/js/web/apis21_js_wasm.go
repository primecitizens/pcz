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

func NewCSSPerspective(length CSSPerspectiveValue) (ret CSSPerspective) {
	ret.ref = bindings.NewCSSPerspectiveByCSSPerspective(
		length.Ref())
	return
}

type CSSPerspective struct {
	CSSTransformComponent
}

func (this CSSPerspective) Once() CSSPerspective {
	this.Ref().Once()
	return this
}

func (this CSSPerspective) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSPerspective) FromRef(ref js.Ref) CSSPerspective {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSPerspective) Free() {
	this.Ref().Free()
}

// Length returns the value of property "CSSPerspective.length".
//
// It returns ok=false if there is no such property.
func (this CSSPerspective) Length() (ret CSSPerspectiveValue, ok bool) {
	ok = js.True == bindings.GetCSSPerspectiveLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLength sets the value of property "CSSPerspective.length" to val.
//
// It returns false if the property cannot be set.
func (this CSSPerspective) SetLength(val CSSPerspectiveValue) bool {
	return js.True == bindings.SetCSSPerspectiveLength(
		this.Ref(),
		val.Ref(),
	)
}

type CSSPositionFallbackRule struct {
	CSSGroupingRule
}

func (this CSSPositionFallbackRule) Once() CSSPositionFallbackRule {
	this.Ref().Once()
	return this
}

func (this CSSPositionFallbackRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSPositionFallbackRule) FromRef(ref js.Ref) CSSPositionFallbackRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSPositionFallbackRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSPositionFallbackRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSPositionFallbackRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPositionFallbackRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CSSPropertyRule struct {
	CSSRule
}

func (this CSSPropertyRule) Once() CSSPropertyRule {
	this.Ref().Once()
	return this
}

func (this CSSPropertyRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSPropertyRule) FromRef(ref js.Ref) CSSPropertyRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSPropertyRule) Free() {
	this.Ref().Free()
}

// Name returns the value of property "CSSPropertyRule.name".
//
// It returns ok=false if there is no such property.
func (this CSSPropertyRule) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPropertyRuleName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Syntax returns the value of property "CSSPropertyRule.syntax".
//
// It returns ok=false if there is no such property.
func (this CSSPropertyRule) Syntax() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPropertyRuleSyntax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Inherits returns the value of property "CSSPropertyRule.inherits".
//
// It returns ok=false if there is no such property.
func (this CSSPropertyRule) Inherits() (ret bool, ok bool) {
	ok = js.True == bindings.GetCSSPropertyRuleInherits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InitialValue returns the value of property "CSSPropertyRule.initialValue".
//
// It returns ok=false if there is no such property.
func (this CSSPropertyRule) InitialValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSPropertyRuleInitialValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSRGB(r CSSColorRGBComp, g CSSColorRGBComp, b CSSColorRGBComp, alpha CSSColorPercent) (ret CSSRGB) {
	ret.ref = bindings.NewCSSRGBByCSSRGB(
		r.Ref(),
		g.Ref(),
		b.Ref(),
		alpha.Ref())
	return
}

func NewCSSRGBByCSSRGB1(r CSSColorRGBComp, g CSSColorRGBComp, b CSSColorRGBComp) (ret CSSRGB) {
	ret.ref = bindings.NewCSSRGBByCSSRGB1(
		r.Ref(),
		g.Ref(),
		b.Ref())
	return
}

type CSSRGB struct {
	CSSColorValue
}

func (this CSSRGB) Once() CSSRGB {
	this.Ref().Once()
	return this
}

func (this CSSRGB) Ref() js.Ref {
	return this.CSSColorValue.Ref()
}

func (this CSSRGB) FromRef(ref js.Ref) CSSRGB {
	this.CSSColorValue = this.CSSColorValue.FromRef(ref)
	return this
}

func (this CSSRGB) Free() {
	this.Ref().Free()
}

// R returns the value of property "CSSRGB.r".
//
// It returns ok=false if there is no such property.
func (this CSSRGB) R() (ret CSSColorRGBComp, ok bool) {
	ok = js.True == bindings.GetCSSRGBR(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetR sets the value of property "CSSRGB.r" to val.
//
// It returns false if the property cannot be set.
func (this CSSRGB) SetR(val CSSColorRGBComp) bool {
	return js.True == bindings.SetCSSRGBR(
		this.Ref(),
		val.Ref(),
	)
}

// G returns the value of property "CSSRGB.g".
//
// It returns ok=false if there is no such property.
func (this CSSRGB) G() (ret CSSColorRGBComp, ok bool) {
	ok = js.True == bindings.GetCSSRGBG(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetG sets the value of property "CSSRGB.g" to val.
//
// It returns false if the property cannot be set.
func (this CSSRGB) SetG(val CSSColorRGBComp) bool {
	return js.True == bindings.SetCSSRGBG(
		this.Ref(),
		val.Ref(),
	)
}

// B returns the value of property "CSSRGB.b".
//
// It returns ok=false if there is no such property.
func (this CSSRGB) B() (ret CSSColorRGBComp, ok bool) {
	ok = js.True == bindings.GetCSSRGBB(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetB sets the value of property "CSSRGB.b" to val.
//
// It returns false if the property cannot be set.
func (this CSSRGB) SetB(val CSSColorRGBComp) bool {
	return js.True == bindings.SetCSSRGBB(
		this.Ref(),
		val.Ref(),
	)
}

// Alpha returns the value of property "CSSRGB.alpha".
//
// It returns ok=false if there is no such property.
func (this CSSRGB) Alpha() (ret CSSColorPercent, ok bool) {
	ok = js.True == bindings.GetCSSRGBAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlpha sets the value of property "CSSRGB.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSRGB) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSRGBAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSRotate(angle CSSNumericValue) (ret CSSRotate) {
	ret.ref = bindings.NewCSSRotateByCSSRotate(
		angle.Ref())
	return
}

func NewCSSRotateByCSSRotate1(x CSSNumberish, y CSSNumberish, z CSSNumberish, angle CSSNumericValue) (ret CSSRotate) {
	ret.ref = bindings.NewCSSRotateByCSSRotate1(
		x.Ref(),
		y.Ref(),
		z.Ref(),
		angle.Ref())
	return
}

type CSSRotate struct {
	CSSTransformComponent
}

func (this CSSRotate) Once() CSSRotate {
	this.Ref().Once()
	return this
}

func (this CSSRotate) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSRotate) FromRef(ref js.Ref) CSSRotate {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSRotate) Free() {
	this.Ref().Free()
}

// X returns the value of property "CSSRotate.x".
//
// It returns ok=false if there is no such property.
func (this CSSRotate) X() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSRotateX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "CSSRotate.x" to val.
//
// It returns false if the property cannot be set.
func (this CSSRotate) SetX(val CSSNumberish) bool {
	return js.True == bindings.SetCSSRotateX(
		this.Ref(),
		val.Ref(),
	)
}

// Y returns the value of property "CSSRotate.y".
//
// It returns ok=false if there is no such property.
func (this CSSRotate) Y() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSRotateY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "CSSRotate.y" to val.
//
// It returns false if the property cannot be set.
func (this CSSRotate) SetY(val CSSNumberish) bool {
	return js.True == bindings.SetCSSRotateY(
		this.Ref(),
		val.Ref(),
	)
}

// Z returns the value of property "CSSRotate.z".
//
// It returns ok=false if there is no such property.
func (this CSSRotate) Z() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSRotateZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetZ sets the value of property "CSSRotate.z" to val.
//
// It returns false if the property cannot be set.
func (this CSSRotate) SetZ(val CSSNumberish) bool {
	return js.True == bindings.SetCSSRotateZ(
		this.Ref(),
		val.Ref(),
	)
}

// Angle returns the value of property "CSSRotate.angle".
//
// It returns ok=false if there is no such property.
func (this CSSRotate) Angle() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSRotateAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAngle sets the value of property "CSSRotate.angle" to val.
//
// It returns false if the property cannot be set.
func (this CSSRotate) SetAngle(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSRotateAngle(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSScale(x CSSNumberish, y CSSNumberish, z CSSNumberish) (ret CSSScale) {
	ret.ref = bindings.NewCSSScaleByCSSScale(
		x.Ref(),
		y.Ref(),
		z.Ref())
	return
}

func NewCSSScaleByCSSScale1(x CSSNumberish, y CSSNumberish) (ret CSSScale) {
	ret.ref = bindings.NewCSSScaleByCSSScale1(
		x.Ref(),
		y.Ref())
	return
}

type CSSScale struct {
	CSSTransformComponent
}

func (this CSSScale) Once() CSSScale {
	this.Ref().Once()
	return this
}

func (this CSSScale) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSScale) FromRef(ref js.Ref) CSSScale {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSScale) Free() {
	this.Ref().Free()
}

// X returns the value of property "CSSScale.x".
//
// It returns ok=false if there is no such property.
func (this CSSScale) X() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSScaleX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "CSSScale.x" to val.
//
// It returns false if the property cannot be set.
func (this CSSScale) SetX(val CSSNumberish) bool {
	return js.True == bindings.SetCSSScaleX(
		this.Ref(),
		val.Ref(),
	)
}

// Y returns the value of property "CSSScale.y".
//
// It returns ok=false if there is no such property.
func (this CSSScale) Y() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSScaleY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "CSSScale.y" to val.
//
// It returns false if the property cannot be set.
func (this CSSScale) SetY(val CSSNumberish) bool {
	return js.True == bindings.SetCSSScaleY(
		this.Ref(),
		val.Ref(),
	)
}

// Z returns the value of property "CSSScale.z".
//
// It returns ok=false if there is no such property.
func (this CSSScale) Z() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetCSSScaleZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetZ sets the value of property "CSSScale.z" to val.
//
// It returns false if the property cannot be set.
func (this CSSScale) SetZ(val CSSNumberish) bool {
	return js.True == bindings.SetCSSScaleZ(
		this.Ref(),
		val.Ref(),
	)
}

type CSSScopeRule struct {
	CSSGroupingRule
}

func (this CSSScopeRule) Once() CSSScopeRule {
	this.Ref().Once()
	return this
}

func (this CSSScopeRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSScopeRule) FromRef(ref js.Ref) CSSScopeRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSScopeRule) Free() {
	this.Ref().Free()
}

// Start returns the value of property "CSSScopeRule.start".
//
// It returns ok=false if there is no such property.
func (this CSSScopeRule) Start() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSScopeRuleStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// End returns the value of property "CSSScopeRule.end".
//
// It returns ok=false if there is no such property.
func (this CSSScopeRule) End() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSScopeRuleEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSSkew(ax CSSNumericValue, ay CSSNumericValue) (ret CSSSkew) {
	ret.ref = bindings.NewCSSSkewByCSSSkew(
		ax.Ref(),
		ay.Ref())
	return
}

type CSSSkew struct {
	CSSTransformComponent
}

func (this CSSSkew) Once() CSSSkew {
	this.Ref().Once()
	return this
}

func (this CSSSkew) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSSkew) FromRef(ref js.Ref) CSSSkew {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSSkew) Free() {
	this.Ref().Free()
}

// Ax returns the value of property "CSSSkew.ax".
//
// It returns ok=false if there is no such property.
func (this CSSSkew) Ax() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSSkewAx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAx sets the value of property "CSSSkew.ax" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkew) SetAx(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewAx(
		this.Ref(),
		val.Ref(),
	)
}

// Ay returns the value of property "CSSSkew.ay".
//
// It returns ok=false if there is no such property.
func (this CSSSkew) Ay() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSSkewAy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAy sets the value of property "CSSSkew.ay" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkew) SetAy(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewAy(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSSkewX(ax CSSNumericValue) (ret CSSSkewX) {
	ret.ref = bindings.NewCSSSkewXByCSSSkewX(
		ax.Ref())
	return
}

type CSSSkewX struct {
	CSSTransformComponent
}

func (this CSSSkewX) Once() CSSSkewX {
	this.Ref().Once()
	return this
}

func (this CSSSkewX) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSSkewX) FromRef(ref js.Ref) CSSSkewX {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSSkewX) Free() {
	this.Ref().Free()
}

// Ax returns the value of property "CSSSkewX.ax".
//
// It returns ok=false if there is no such property.
func (this CSSSkewX) Ax() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSSkewXAx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAx sets the value of property "CSSSkewX.ax" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkewX) SetAx(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewXAx(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSSkewY(ay CSSNumericValue) (ret CSSSkewY) {
	ret.ref = bindings.NewCSSSkewYByCSSSkewY(
		ay.Ref())
	return
}

type CSSSkewY struct {
	CSSTransformComponent
}

func (this CSSSkewY) Once() CSSSkewY {
	this.Ref().Once()
	return this
}

func (this CSSSkewY) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSSkewY) FromRef(ref js.Ref) CSSSkewY {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSSkewY) Free() {
	this.Ref().Free()
}

// Ay returns the value of property "CSSSkewY.ay".
//
// It returns ok=false if there is no such property.
func (this CSSSkewY) Ay() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSSkewYAy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAy sets the value of property "CSSSkewY.ay" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkewY) SetAy(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewYAy(
		this.Ref(),
		val.Ref(),
	)
}

type CSSStartingStyleRule struct {
	CSSGroupingRule
}

func (this CSSStartingStyleRule) Once() CSSStartingStyleRule {
	this.Ref().Once()
	return this
}

func (this CSSStartingStyleRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSStartingStyleRule) FromRef(ref js.Ref) CSSStartingStyleRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSStartingStyleRule) Free() {
	this.Ref().Free()
}

type CSSStyleRule struct {
	CSSGroupingRule
}

func (this CSSStyleRule) Once() CSSStyleRule {
	this.Ref().Once()
	return this
}

func (this CSSStyleRule) Ref() js.Ref {
	return this.CSSGroupingRule.Ref()
}

func (this CSSStyleRule) FromRef(ref js.Ref) CSSStyleRule {
	this.CSSGroupingRule = this.CSSGroupingRule.FromRef(ref)
	return this
}

func (this CSSStyleRule) Free() {
	this.Ref().Free()
}

// SelectorText returns the value of property "CSSStyleRule.selectorText".
//
// It returns ok=false if there is no such property.
func (this CSSStyleRule) SelectorText() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSStyleRuleSelectorText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectorText sets the value of property "CSSStyleRule.selectorText" to val.
//
// It returns false if the property cannot be set.
func (this CSSStyleRule) SetSelectorText(val js.String) bool {
	return js.True == bindings.SetCSSStyleRuleSelectorText(
		this.Ref(),
		val.Ref(),
	)
}

// Style returns the value of property "CSSStyleRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSStyleRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSStyleRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StyleMap returns the value of property "CSSStyleRule.styleMap".
//
// It returns ok=false if there is no such property.
func (this CSSStyleRule) StyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetCSSStyleRuleStyleMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CSSSupportsRule struct {
	CSSConditionRule
}

func (this CSSSupportsRule) Once() CSSSupportsRule {
	this.Ref().Once()
	return this
}

func (this CSSSupportsRule) Ref() js.Ref {
	return this.CSSConditionRule.Ref()
}

func (this CSSSupportsRule) FromRef(ref js.Ref) CSSSupportsRule {
	this.CSSConditionRule = this.CSSConditionRule.FromRef(ref)
	return this
}

func (this CSSSupportsRule) Free() {
	this.Ref().Free()
}

type CSSTransformComponent struct {
	ref js.Ref
}

func (this CSSTransformComponent) Once() CSSTransformComponent {
	this.Ref().Once()
	return this
}

func (this CSSTransformComponent) Ref() js.Ref {
	return this.ref
}

func (this CSSTransformComponent) FromRef(ref js.Ref) CSSTransformComponent {
	this.ref = ref
	return this
}

func (this CSSTransformComponent) Free() {
	this.Ref().Free()
}

// Is2D returns the value of property "CSSTransformComponent.is2D".
//
// It returns ok=false if there is no such property.
func (this CSSTransformComponent) Is2D() (ret bool, ok bool) {
	ok = js.True == bindings.GetCSSTransformComponentIs2D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIs2D sets the value of property "CSSTransformComponent.is2D" to val.
//
// It returns false if the property cannot be set.
func (this CSSTransformComponent) SetIs2D(val bool) bool {
	return js.True == bindings.SetCSSTransformComponentIs2D(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// HasToString returns true if the method "CSSTransformComponent.toString" exists.
func (this CSSTransformComponent) HasToString() bool {
	return js.True == bindings.HasCSSTransformComponentToString(
		this.Ref(),
	)
}

// ToStringFunc returns the method "CSSTransformComponent.toString".
func (this CSSTransformComponent) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSTransformComponentToStringFunc(
			this.Ref(),
		),
	)
}

// ToString calls the method "CSSTransformComponent.toString".
func (this CSSTransformComponent) ToString() (ret js.String) {
	bindings.CallCSSTransformComponentToString(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToString calls the method "CSSTransformComponent.toString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSTransformComponent) TryToString() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTransformComponentToString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToMatrix returns true if the method "CSSTransformComponent.toMatrix" exists.
func (this CSSTransformComponent) HasToMatrix() bool {
	return js.True == bindings.HasCSSTransformComponentToMatrix(
		this.Ref(),
	)
}

// ToMatrixFunc returns the method "CSSTransformComponent.toMatrix".
func (this CSSTransformComponent) ToMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CSSTransformComponentToMatrixFunc(
			this.Ref(),
		),
	)
}

// ToMatrix calls the method "CSSTransformComponent.toMatrix".
func (this CSSTransformComponent) ToMatrix() (ret DOMMatrix) {
	bindings.CallCSSTransformComponentToMatrix(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToMatrix calls the method "CSSTransformComponent.toMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSTransformComponent) TryToMatrix() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTransformComponentToMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewCSSTransformValue(transforms js.Array[CSSTransformComponent]) (ret CSSTransformValue) {
	ret.ref = bindings.NewCSSTransformValueByCSSTransformValue(
		transforms.Ref())
	return
}

type CSSTransformValue struct {
	CSSStyleValue
}

func (this CSSTransformValue) Once() CSSTransformValue {
	this.Ref().Once()
	return this
}

func (this CSSTransformValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSTransformValue) FromRef(ref js.Ref) CSSTransformValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSTransformValue) Free() {
	this.Ref().Free()
}

// Length returns the value of property "CSSTransformValue.length".
//
// It returns ok=false if there is no such property.
func (this CSSTransformValue) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSTransformValueLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Is2D returns the value of property "CSSTransformValue.is2D".
//
// It returns ok=false if there is no such property.
func (this CSSTransformValue) Is2D() (ret bool, ok bool) {
	ok = js.True == bindings.GetCSSTransformValueIs2D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "CSSTransformValue." exists.
func (this CSSTransformValue) HasGet() bool {
	return js.True == bindings.HasCSSTransformValueGet(
		this.Ref(),
	)
}

// GetFunc returns the method "CSSTransformValue.".
func (this CSSTransformValue) GetFunc() (fn js.Func[func(index uint32) CSSTransformComponent]) {
	return fn.FromRef(
		bindings.CSSTransformValueGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "CSSTransformValue.".
func (this CSSTransformValue) Get(index uint32) (ret CSSTransformComponent) {
	bindings.CallCSSTransformValueGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "CSSTransformValue."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSTransformValue) TryGet(index uint32) (ret CSSTransformComponent, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTransformValueGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasSet returns true if the method "CSSTransformValue." exists.
func (this CSSTransformValue) HasSet() bool {
	return js.True == bindings.HasCSSTransformValueSet(
		this.Ref(),
	)
}

// SetFunc returns the method "CSSTransformValue.".
func (this CSSTransformValue) SetFunc() (fn js.Func[func(index uint32, val CSSTransformComponent)]) {
	return fn.FromRef(
		bindings.CSSTransformValueSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CSSTransformValue.".
func (this CSSTransformValue) Set(index uint32, val CSSTransformComponent) (ret js.Void) {
	bindings.CallCSSTransformValueSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		val.Ref(),
	)

	return
}

// TrySet calls the method "CSSTransformValue."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSTransformValue) TrySet(index uint32, val CSSTransformComponent) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTransformValueSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		val.Ref(),
	)

	return
}

// HasToMatrix returns true if the method "CSSTransformValue.toMatrix" exists.
func (this CSSTransformValue) HasToMatrix() bool {
	return js.True == bindings.HasCSSTransformValueToMatrix(
		this.Ref(),
	)
}

// ToMatrixFunc returns the method "CSSTransformValue.toMatrix".
func (this CSSTransformValue) ToMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CSSTransformValueToMatrixFunc(
			this.Ref(),
		),
	)
}

// ToMatrix calls the method "CSSTransformValue.toMatrix".
func (this CSSTransformValue) ToMatrix() (ret DOMMatrix) {
	bindings.CallCSSTransformValueToMatrix(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToMatrix calls the method "CSSTransformValue.toMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSTransformValue) TryToMatrix() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSTransformValueToMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewCSSTransition(effect AnimationEffect, timeline AnimationTimeline) (ret CSSTransition) {
	ret.ref = bindings.NewCSSTransitionByCSSTransition(
		effect.Ref(),
		timeline.Ref())
	return
}

func NewCSSTransitionByCSSTransition1(effect AnimationEffect) (ret CSSTransition) {
	ret.ref = bindings.NewCSSTransitionByCSSTransition1(
		effect.Ref())
	return
}

func NewCSSTransitionByCSSTransition2() (ret CSSTransition) {
	ret.ref = bindings.NewCSSTransitionByCSSTransition2()
	return
}

type CSSTransition struct {
	Animation
}

func (this CSSTransition) Once() CSSTransition {
	this.Ref().Once()
	return this
}

func (this CSSTransition) Ref() js.Ref {
	return this.Animation.Ref()
}

func (this CSSTransition) FromRef(ref js.Ref) CSSTransition {
	this.Animation = this.Animation.FromRef(ref)
	return this
}

func (this CSSTransition) Free() {
	this.Ref().Free()
}

// TransitionProperty returns the value of property "CSSTransition.transitionProperty".
//
// It returns ok=false if there is no such property.
func (this CSSTransition) TransitionProperty() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSTransitionTransitionProperty(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSTranslate(x CSSNumericValue, y CSSNumericValue, z CSSNumericValue) (ret CSSTranslate) {
	ret.ref = bindings.NewCSSTranslateByCSSTranslate(
		x.Ref(),
		y.Ref(),
		z.Ref())
	return
}

func NewCSSTranslateByCSSTranslate1(x CSSNumericValue, y CSSNumericValue) (ret CSSTranslate) {
	ret.ref = bindings.NewCSSTranslateByCSSTranslate1(
		x.Ref(),
		y.Ref())
	return
}

type CSSTranslate struct {
	CSSTransformComponent
}

func (this CSSTranslate) Once() CSSTranslate {
	this.Ref().Once()
	return this
}

func (this CSSTranslate) Ref() js.Ref {
	return this.CSSTransformComponent.Ref()
}

func (this CSSTranslate) FromRef(ref js.Ref) CSSTranslate {
	this.CSSTransformComponent = this.CSSTransformComponent.FromRef(ref)
	return this
}

func (this CSSTranslate) Free() {
	this.Ref().Free()
}

// X returns the value of property "CSSTranslate.x".
//
// It returns ok=false if there is no such property.
func (this CSSTranslate) X() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSTranslateX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetX sets the value of property "CSSTranslate.x" to val.
//
// It returns false if the property cannot be set.
func (this CSSTranslate) SetX(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSTranslateX(
		this.Ref(),
		val.Ref(),
	)
}

// Y returns the value of property "CSSTranslate.y".
//
// It returns ok=false if there is no such property.
func (this CSSTranslate) Y() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSTranslateY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetY sets the value of property "CSSTranslate.y" to val.
//
// It returns false if the property cannot be set.
func (this CSSTranslate) SetY(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSTranslateY(
		this.Ref(),
		val.Ref(),
	)
}

// Z returns the value of property "CSSTranslate.z".
//
// It returns ok=false if there is no such property.
func (this CSSTranslate) Z() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetCSSTranslateZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetZ sets the value of property "CSSTranslate.z" to val.
//
// It returns false if the property cannot be set.
func (this CSSTranslate) SetZ(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSTranslateZ(
		this.Ref(),
		val.Ref(),
	)
}

type CSSTryRule struct {
	CSSRule
}

func (this CSSTryRule) Once() CSSTryRule {
	this.Ref().Once()
	return this
}

func (this CSSTryRule) Ref() js.Ref {
	return this.CSSRule.Ref()
}

func (this CSSTryRule) FromRef(ref js.Ref) CSSTryRule {
	this.CSSRule = this.CSSRule.FromRef(ref)
	return this
}

func (this CSSTryRule) Free() {
	this.Ref().Free()
}

// Style returns the value of property "CSSTryRule.style".
//
// It returns ok=false if there is no such property.
func (this CSSTryRule) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetCSSTryRuleStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewCSSUnparsedValue(members js.Array[CSSUnparsedSegment]) (ret CSSUnparsedValue) {
	ret.ref = bindings.NewCSSUnparsedValueByCSSUnparsedValue(
		members.Ref())
	return
}

type CSSUnparsedValue struct {
	CSSStyleValue
}

func (this CSSUnparsedValue) Once() CSSUnparsedValue {
	this.Ref().Once()
	return this
}

func (this CSSUnparsedValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSUnparsedValue) FromRef(ref js.Ref) CSSUnparsedValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSUnparsedValue) Free() {
	this.Ref().Free()
}

// Length returns the value of property "CSSUnparsedValue.length".
//
// It returns ok=false if there is no such property.
func (this CSSUnparsedValue) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSUnparsedValueLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "CSSUnparsedValue." exists.
func (this CSSUnparsedValue) HasGet() bool {
	return js.True == bindings.HasCSSUnparsedValueGet(
		this.Ref(),
	)
}

// GetFunc returns the method "CSSUnparsedValue.".
func (this CSSUnparsedValue) GetFunc() (fn js.Func[func(index uint32) CSSUnparsedSegment]) {
	return fn.FromRef(
		bindings.CSSUnparsedValueGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "CSSUnparsedValue.".
func (this CSSUnparsedValue) Get(index uint32) (ret CSSUnparsedSegment) {
	bindings.CallCSSUnparsedValueGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "CSSUnparsedValue."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSUnparsedValue) TryGet(index uint32) (ret CSSUnparsedSegment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSUnparsedValueGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasSet returns true if the method "CSSUnparsedValue." exists.
func (this CSSUnparsedValue) HasSet() bool {
	return js.True == bindings.HasCSSUnparsedValueSet(
		this.Ref(),
	)
}

// SetFunc returns the method "CSSUnparsedValue.".
func (this CSSUnparsedValue) SetFunc() (fn js.Func[func(index uint32, val CSSUnparsedSegment)]) {
	return fn.FromRef(
		bindings.CSSUnparsedValueSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CSSUnparsedValue.".
func (this CSSUnparsedValue) Set(index uint32, val CSSUnparsedSegment) (ret js.Void) {
	bindings.CallCSSUnparsedValueSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		val.Ref(),
	)

	return
}

// TrySet calls the method "CSSUnparsedValue."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSUnparsedValue) TrySet(index uint32, val CSSUnparsedSegment) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSUnparsedValueSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		val.Ref(),
	)

	return
}

func NewCSSVariableReferenceValue(variable js.String, fallback CSSUnparsedValue) (ret CSSVariableReferenceValue) {
	ret.ref = bindings.NewCSSVariableReferenceValueByCSSVariableReferenceValue(
		variable.Ref(),
		fallback.Ref())
	return
}

func NewCSSVariableReferenceValueByCSSVariableReferenceValue1(variable js.String) (ret CSSVariableReferenceValue) {
	ret.ref = bindings.NewCSSVariableReferenceValueByCSSVariableReferenceValue1(
		variable.Ref())
	return
}

type CSSVariableReferenceValue struct {
	ref js.Ref
}

func (this CSSVariableReferenceValue) Once() CSSVariableReferenceValue {
	this.Ref().Once()
	return this
}

func (this CSSVariableReferenceValue) Ref() js.Ref {
	return this.ref
}

func (this CSSVariableReferenceValue) FromRef(ref js.Ref) CSSVariableReferenceValue {
	this.ref = ref
	return this
}

func (this CSSVariableReferenceValue) Free() {
	this.Ref().Free()
}

// Variable returns the value of property "CSSVariableReferenceValue.variable".
//
// It returns ok=false if there is no such property.
func (this CSSVariableReferenceValue) Variable() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSVariableReferenceValueVariable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVariable sets the value of property "CSSVariableReferenceValue.variable" to val.
//
// It returns false if the property cannot be set.
func (this CSSVariableReferenceValue) SetVariable(val js.String) bool {
	return js.True == bindings.SetCSSVariableReferenceValueVariable(
		this.Ref(),
		val.Ref(),
	)
}

// Fallback returns the value of property "CSSVariableReferenceValue.fallback".
//
// It returns ok=false if there is no such property.
func (this CSSVariableReferenceValue) Fallback() (ret CSSUnparsedValue, ok bool) {
	ok = js.True == bindings.GetCSSVariableReferenceValueFallback(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_String_CSSVariableReferenceValue struct {
	ref js.Ref
}

func (x OneOf_String_CSSVariableReferenceValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_CSSVariableReferenceValue) Free() {
	x.ref.Free()
}

func (x OneOf_String_CSSVariableReferenceValue) FromRef(ref js.Ref) OneOf_String_CSSVariableReferenceValue {
	return OneOf_String_CSSVariableReferenceValue{
		ref: ref,
	}
}

func (x OneOf_String_CSSVariableReferenceValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_CSSVariableReferenceValue) CSSVariableReferenceValue() CSSVariableReferenceValue {
	return CSSVariableReferenceValue{}.FromRef(x.ref)
}

type CSSUnparsedSegment = OneOf_String_CSSVariableReferenceValue

type OneOf_Response_undefined struct {
	ref js.Ref
}

func (x OneOf_Response_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Response_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_Response_undefined) FromRef(ref js.Ref) OneOf_Response_undefined {
	return OneOf_Response_undefined{
		ref: ref,
	}
}

func (x OneOf_Response_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_Response_undefined) Response() Response {
	return Response{}.FromRef(x.ref)
}

type Cache struct {
	ref js.Ref
}

func (this Cache) Once() Cache {
	this.Ref().Once()
	return this
}

func (this Cache) Ref() js.Ref {
	return this.ref
}

func (this Cache) FromRef(ref js.Ref) Cache {
	this.ref = ref
	return this
}

func (this Cache) Free() {
	this.Ref().Free()
}

// HasMatch returns true if the method "Cache.match" exists.
func (this Cache) HasMatch() bool {
	return js.True == bindings.HasCacheMatch(
		this.Ref(),
	)
}

// MatchFunc returns the method "Cache.match".
func (this Cache) MatchFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheMatchFunc(
			this.Ref(),
		),
	)
}

// Match calls the method "Cache.match".
func (this Cache) Match(request RequestInfo, options CacheQueryOptions) (ret js.Promise[OneOf_Response_undefined]) {
	bindings.CallCacheMatch(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatch calls the method "Cache.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryMatch(request RequestInfo, options CacheQueryOptions) (ret js.Promise[OneOf_Response_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheMatch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMatch1 returns true if the method "Cache.match" exists.
func (this Cache) HasMatch1() bool {
	return js.True == bindings.HasCacheMatch1(
		this.Ref(),
	)
}

// Match1Func returns the method "Cache.match".
func (this Cache) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheMatch1Func(
			this.Ref(),
		),
	)
}

// Match1 calls the method "Cache.match".
func (this Cache) Match1(request RequestInfo) (ret js.Promise[OneOf_Response_undefined]) {
	bindings.CallCacheMatch1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatch1 calls the method "Cache.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryMatch1(request RequestInfo) (ret js.Promise[OneOf_Response_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheMatch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasMatchAll returns true if the method "Cache.matchAll" exists.
func (this Cache) HasMatchAll() bool {
	return js.True == bindings.HasCacheMatchAll(
		this.Ref(),
	)
}

// MatchAllFunc returns the method "Cache.matchAll".
func (this Cache) MatchAllFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "Cache.matchAll".
func (this Cache) MatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.FrozenArray[Response]]) {
	bindings.CallCacheMatchAll(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatchAll calls the method "Cache.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryMatchAll(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.FrozenArray[Response]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheMatchAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMatchAll1 returns true if the method "Cache.matchAll" exists.
func (this Cache) HasMatchAll1() bool {
	return js.True == bindings.HasCacheMatchAll1(
		this.Ref(),
	)
}

// MatchAll1Func returns the method "Cache.matchAll".
func (this Cache) MatchAll1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAll1Func(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "Cache.matchAll".
func (this Cache) MatchAll1(request RequestInfo) (ret js.Promise[js.FrozenArray[Response]]) {
	bindings.CallCacheMatchAll1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatchAll1 calls the method "Cache.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryMatchAll1(request RequestInfo) (ret js.Promise[js.FrozenArray[Response]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheMatchAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasMatchAll2 returns true if the method "Cache.matchAll" exists.
func (this Cache) HasMatchAll2() bool {
	return js.True == bindings.HasCacheMatchAll2(
		this.Ref(),
	)
}

// MatchAll2Func returns the method "Cache.matchAll".
func (this Cache) MatchAll2Func() (fn js.Func[func() js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAll2Func(
			this.Ref(),
		),
	)
}

// MatchAll2 calls the method "Cache.matchAll".
func (this Cache) MatchAll2() (ret js.Promise[js.FrozenArray[Response]]) {
	bindings.CallCacheMatchAll2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMatchAll2 calls the method "Cache.matchAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryMatchAll2() (ret js.Promise[js.FrozenArray[Response]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheMatchAll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAdd returns true if the method "Cache.add" exists.
func (this Cache) HasAdd() bool {
	return js.True == bindings.HasCacheAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "Cache.add".
func (this Cache) AddFunc() (fn js.Func[func(request RequestInfo) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CacheAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "Cache.add".
func (this Cache) Add(request RequestInfo) (ret js.Promise[js.Void]) {
	bindings.CallCacheAdd(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryAdd calls the method "Cache.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryAdd(request RequestInfo) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasAddAll returns true if the method "Cache.addAll" exists.
func (this Cache) HasAddAll() bool {
	return js.True == bindings.HasCacheAddAll(
		this.Ref(),
	)
}

// AddAllFunc returns the method "Cache.addAll".
func (this Cache) AddAllFunc() (fn js.Func[func(requests js.Array[RequestInfo]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CacheAddAllFunc(
			this.Ref(),
		),
	)
}

// AddAll calls the method "Cache.addAll".
func (this Cache) AddAll(requests js.Array[RequestInfo]) (ret js.Promise[js.Void]) {
	bindings.CallCacheAddAll(
		this.Ref(), js.Pointer(&ret),
		requests.Ref(),
	)

	return
}

// TryAddAll calls the method "Cache.addAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryAddAll(requests js.Array[RequestInfo]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheAddAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		requests.Ref(),
	)

	return
}

// HasPut returns true if the method "Cache.put" exists.
func (this Cache) HasPut() bool {
	return js.True == bindings.HasCachePut(
		this.Ref(),
	)
}

// PutFunc returns the method "Cache.put".
func (this Cache) PutFunc() (fn js.Func[func(request RequestInfo, response Response) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CachePutFunc(
			this.Ref(),
		),
	)
}

// Put calls the method "Cache.put".
func (this Cache) Put(request RequestInfo, response Response) (ret js.Promise[js.Void]) {
	bindings.CallCachePut(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		response.Ref(),
	)

	return
}

// TryPut calls the method "Cache.put"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryPut(request RequestInfo, response Response) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCachePut(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		response.Ref(),
	)

	return
}

// HasDelete returns true if the method "Cache.delete" exists.
func (this Cache) HasDelete() bool {
	return js.True == bindings.HasCacheDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "Cache.delete".
func (this Cache) DeleteFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "Cache.delete".
func (this Cache) Delete(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Boolean]) {
	bindings.CallCacheDelete(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryDelete calls the method "Cache.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryDelete(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasDelete1 returns true if the method "Cache.delete" exists.
func (this Cache) HasDelete1() bool {
	return js.True == bindings.HasCacheDelete1(
		this.Ref(),
	)
}

// Delete1Func returns the method "Cache.delete".
func (this Cache) Delete1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheDelete1Func(
			this.Ref(),
		),
	)
}

// Delete1 calls the method "Cache.delete".
func (this Cache) Delete1(request RequestInfo) (ret js.Promise[js.Boolean]) {
	bindings.CallCacheDelete1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryDelete1 calls the method "Cache.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryDelete1(request RequestInfo) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheDelete1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasKeys returns true if the method "Cache.keys" exists.
func (this Cache) HasKeys() bool {
	return js.True == bindings.HasCacheKeys(
		this.Ref(),
	)
}

// KeysFunc returns the method "Cache.keys".
func (this Cache) KeysFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeysFunc(
			this.Ref(),
		),
	)
}

// Keys calls the method "Cache.keys".
func (this Cache) Keys(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.FrozenArray[Request]]) {
	bindings.CallCacheKeys(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryKeys calls the method "Cache.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryKeys(request RequestInfo, options CacheQueryOptions) (ret js.Promise[js.FrozenArray[Request]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasKeys1 returns true if the method "Cache.keys" exists.
func (this Cache) HasKeys1() bool {
	return js.True == bindings.HasCacheKeys1(
		this.Ref(),
	)
}

// Keys1Func returns the method "Cache.keys".
func (this Cache) Keys1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeys1Func(
			this.Ref(),
		),
	)
}

// Keys1 calls the method "Cache.keys".
func (this Cache) Keys1(request RequestInfo) (ret js.Promise[js.FrozenArray[Request]]) {
	bindings.CallCacheKeys1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryKeys1 calls the method "Cache.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryKeys1(request RequestInfo) (ret js.Promise[js.FrozenArray[Request]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheKeys1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasKeys2 returns true if the method "Cache.keys" exists.
func (this Cache) HasKeys2() bool {
	return js.True == bindings.HasCacheKeys2(
		this.Ref(),
	)
}

// Keys2Func returns the method "Cache.keys".
func (this Cache) Keys2Func() (fn js.Func[func() js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeys2Func(
			this.Ref(),
		),
	)
}

// Keys2 calls the method "Cache.keys".
func (this Cache) Keys2() (ret js.Promise[js.FrozenArray[Request]]) {
	bindings.CallCacheKeys2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryKeys2 calls the method "Cache.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Cache) TryKeys2() (ret js.Promise[js.FrozenArray[Request]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheKeys2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MultiCacheQueryOptions struct {
	// CacheName is "MultiCacheQueryOptions.cacheName"
	//
	// Optional
	CacheName js.String
	// IgnoreSearch is "MultiCacheQueryOptions.ignoreSearch"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreSearch MUST be set to true to make this field effective.
	IgnoreSearch bool
	// IgnoreMethod is "MultiCacheQueryOptions.ignoreMethod"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreMethod MUST be set to true to make this field effective.
	IgnoreMethod bool
	// IgnoreVary is "MultiCacheQueryOptions.ignoreVary"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreVary MUST be set to true to make this field effective.
	IgnoreVary bool

	FFI_USE_IgnoreSearch bool // for IgnoreSearch.
	FFI_USE_IgnoreMethod bool // for IgnoreMethod.
	FFI_USE_IgnoreVary   bool // for IgnoreVary.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MultiCacheQueryOptions with all fields set.
func (p MultiCacheQueryOptions) FromRef(ref js.Ref) MultiCacheQueryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MultiCacheQueryOptions in the application heap.
func (p MultiCacheQueryOptions) New() js.Ref {
	return bindings.MultiCacheQueryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MultiCacheQueryOptions) UpdateFrom(ref js.Ref) {
	bindings.MultiCacheQueryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MultiCacheQueryOptions) Update(ref js.Ref) {
	bindings.MultiCacheQueryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CacheStorage struct {
	ref js.Ref
}

func (this CacheStorage) Once() CacheStorage {
	this.Ref().Once()
	return this
}

func (this CacheStorage) Ref() js.Ref {
	return this.ref
}

func (this CacheStorage) FromRef(ref js.Ref) CacheStorage {
	this.ref = ref
	return this
}

func (this CacheStorage) Free() {
	this.Ref().Free()
}

// HasMatch returns true if the method "CacheStorage.match" exists.
func (this CacheStorage) HasMatch() bool {
	return js.True == bindings.HasCacheStorageMatch(
		this.Ref(),
	)
}

// MatchFunc returns the method "CacheStorage.match".
func (this CacheStorage) MatchFunc() (fn js.Func[func(request RequestInfo, options MultiCacheQueryOptions) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheStorageMatchFunc(
			this.Ref(),
		),
	)
}

// Match calls the method "CacheStorage.match".
func (this CacheStorage) Match(request RequestInfo, options MultiCacheQueryOptions) (ret js.Promise[OneOf_Response_undefined]) {
	bindings.CallCacheStorageMatch(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMatch calls the method "CacheStorage.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryMatch(request RequestInfo, options MultiCacheQueryOptions) (ret js.Promise[OneOf_Response_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageMatch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMatch1 returns true if the method "CacheStorage.match" exists.
func (this CacheStorage) HasMatch1() bool {
	return js.True == bindings.HasCacheStorageMatch1(
		this.Ref(),
	)
}

// Match1Func returns the method "CacheStorage.match".
func (this CacheStorage) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheStorageMatch1Func(
			this.Ref(),
		),
	)
}

// Match1 calls the method "CacheStorage.match".
func (this CacheStorage) Match1(request RequestInfo) (ret js.Promise[OneOf_Response_undefined]) {
	bindings.CallCacheStorageMatch1(
		this.Ref(), js.Pointer(&ret),
		request.Ref(),
	)

	return
}

// TryMatch1 calls the method "CacheStorage.match"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryMatch1(request RequestInfo) (ret js.Promise[OneOf_Response_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageMatch1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		request.Ref(),
	)

	return
}

// HasHas returns true if the method "CacheStorage.has" exists.
func (this CacheStorage) HasHas() bool {
	return js.True == bindings.HasCacheStorageHas(
		this.Ref(),
	)
}

// HasFunc returns the method "CacheStorage.has".
func (this CacheStorage) HasFunc() (fn js.Func[func(cacheName js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheStorageHasFunc(
			this.Ref(),
		),
	)
}

// Has calls the method "CacheStorage.has".
func (this CacheStorage) Has(cacheName js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallCacheStorageHas(
		this.Ref(), js.Pointer(&ret),
		cacheName.Ref(),
	)

	return
}

// TryHas calls the method "CacheStorage.has"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryHas(cacheName js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageHas(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		cacheName.Ref(),
	)

	return
}

// HasOpen returns true if the method "CacheStorage.open" exists.
func (this CacheStorage) HasOpen() bool {
	return js.True == bindings.HasCacheStorageOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "CacheStorage.open".
func (this CacheStorage) OpenFunc() (fn js.Func[func(cacheName js.String) js.Promise[Cache]]) {
	return fn.FromRef(
		bindings.CacheStorageOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "CacheStorage.open".
func (this CacheStorage) Open(cacheName js.String) (ret js.Promise[Cache]) {
	bindings.CallCacheStorageOpen(
		this.Ref(), js.Pointer(&ret),
		cacheName.Ref(),
	)

	return
}

// TryOpen calls the method "CacheStorage.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryOpen(cacheName js.String) (ret js.Promise[Cache], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		cacheName.Ref(),
	)

	return
}

// HasDelete returns true if the method "CacheStorage.delete" exists.
func (this CacheStorage) HasDelete() bool {
	return js.True == bindings.HasCacheStorageDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "CacheStorage.delete".
func (this CacheStorage) DeleteFunc() (fn js.Func[func(cacheName js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheStorageDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "CacheStorage.delete".
func (this CacheStorage) Delete(cacheName js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallCacheStorageDelete(
		this.Ref(), js.Pointer(&ret),
		cacheName.Ref(),
	)

	return
}

// TryDelete calls the method "CacheStorage.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryDelete(cacheName js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		cacheName.Ref(),
	)

	return
}

// HasKeys returns true if the method "CacheStorage.keys" exists.
func (this CacheStorage) HasKeys() bool {
	return js.True == bindings.HasCacheStorageKeys(
		this.Ref(),
	)
}

// KeysFunc returns the method "CacheStorage.keys".
func (this CacheStorage) KeysFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.CacheStorageKeysFunc(
			this.Ref(),
		),
	)
}

// Keys calls the method "CacheStorage.keys".
func (this CacheStorage) Keys() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallCacheStorageKeys(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryKeys calls the method "CacheStorage.keys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CacheStorage) TryKeys() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCacheStorageKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CameraDevicePermissionDescriptor struct {
	// PanTiltZoom is "CameraDevicePermissionDescriptor.panTiltZoom"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PanTiltZoom MUST be set to true to make this field effective.
	PanTiltZoom bool
	// Name is "CameraDevicePermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE_PanTiltZoom bool // for PanTiltZoom.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CameraDevicePermissionDescriptor with all fields set.
func (p CameraDevicePermissionDescriptor) FromRef(ref js.Ref) CameraDevicePermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CameraDevicePermissionDescriptor in the application heap.
func (p CameraDevicePermissionDescriptor) New() js.Ref {
	return bindings.CameraDevicePermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CameraDevicePermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.CameraDevicePermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CameraDevicePermissionDescriptor) Update(ref js.Ref) {
	bindings.CameraDevicePermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCanMakePaymentEvent(typ js.String) (ret CanMakePaymentEvent) {
	ret.ref = bindings.NewCanMakePaymentEventByCanMakePaymentEvent(
		typ.Ref())
	return
}

type CanMakePaymentEvent struct {
	ExtendableEvent
}

func (this CanMakePaymentEvent) Once() CanMakePaymentEvent {
	this.Ref().Once()
	return this
}

func (this CanMakePaymentEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this CanMakePaymentEvent) FromRef(ref js.Ref) CanMakePaymentEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this CanMakePaymentEvent) Free() {
	this.Ref().Free()
}

// HasRespondWith returns true if the method "CanMakePaymentEvent.respondWith" exists.
func (this CanMakePaymentEvent) HasRespondWith() bool {
	return js.True == bindings.HasCanMakePaymentEventRespondWith(
		this.Ref(),
	)
}

// RespondWithFunc returns the method "CanMakePaymentEvent.respondWith".
func (this CanMakePaymentEvent) RespondWithFunc() (fn js.Func[func(canMakePaymentResponse js.Promise[js.Boolean])]) {
	return fn.FromRef(
		bindings.CanMakePaymentEventRespondWithFunc(
			this.Ref(),
		),
	)
}

// RespondWith calls the method "CanMakePaymentEvent.respondWith".
func (this CanMakePaymentEvent) RespondWith(canMakePaymentResponse js.Promise[js.Boolean]) (ret js.Void) {
	bindings.CallCanMakePaymentEventRespondWith(
		this.Ref(), js.Pointer(&ret),
		canMakePaymentResponse.Ref(),
	)

	return
}

// TryRespondWith calls the method "CanMakePaymentEvent.respondWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanMakePaymentEvent) TryRespondWith(canMakePaymentResponse js.Promise[js.Boolean]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanMakePaymentEventRespondWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		canMakePaymentResponse.Ref(),
	)

	return
}

type CanvasCaptureMediaStreamTrack struct {
	MediaStreamTrack
}

func (this CanvasCaptureMediaStreamTrack) Once() CanvasCaptureMediaStreamTrack {
	this.Ref().Once()
	return this
}

func (this CanvasCaptureMediaStreamTrack) Ref() js.Ref {
	return this.MediaStreamTrack.Ref()
}

func (this CanvasCaptureMediaStreamTrack) FromRef(ref js.Ref) CanvasCaptureMediaStreamTrack {
	this.MediaStreamTrack = this.MediaStreamTrack.FromRef(ref)
	return this
}

func (this CanvasCaptureMediaStreamTrack) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "CanvasCaptureMediaStreamTrack.canvas".
//
// It returns ok=false if there is no such property.
func (this CanvasCaptureMediaStreamTrack) Canvas() (ret HTMLCanvasElement, ok bool) {
	ok = js.True == bindings.GetCanvasCaptureMediaStreamTrackCanvas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestFrame returns true if the method "CanvasCaptureMediaStreamTrack.requestFrame" exists.
func (this CanvasCaptureMediaStreamTrack) HasRequestFrame() bool {
	return js.True == bindings.HasCanvasCaptureMediaStreamTrackRequestFrame(
		this.Ref(),
	)
}

// RequestFrameFunc returns the method "CanvasCaptureMediaStreamTrack.requestFrame".
func (this CanvasCaptureMediaStreamTrack) RequestFrameFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasCaptureMediaStreamTrackRequestFrameFunc(
			this.Ref(),
		),
	)
}

// RequestFrame calls the method "CanvasCaptureMediaStreamTrack.requestFrame".
func (this CanvasCaptureMediaStreamTrack) RequestFrame() (ret js.Void) {
	bindings.CallCanvasCaptureMediaStreamTrackRequestFrame(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestFrame calls the method "CanvasCaptureMediaStreamTrack.requestFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasCaptureMediaStreamTrack) TryRequestFrame() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasCaptureMediaStreamTrackRequestFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CaptureActionEventInit struct {
	// Action is "CaptureActionEventInit.action"
	//
	// Optional
	Action js.String
	// Bubbles is "CaptureActionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CaptureActionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CaptureActionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CaptureActionEventInit with all fields set.
func (p CaptureActionEventInit) FromRef(ref js.Ref) CaptureActionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CaptureActionEventInit in the application heap.
func (p CaptureActionEventInit) New() js.Ref {
	return bindings.CaptureActionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CaptureActionEventInit) UpdateFrom(ref js.Ref) {
	bindings.CaptureActionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CaptureActionEventInit) Update(ref js.Ref) {
	bindings.CaptureActionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCaptureActionEvent(init CaptureActionEventInit) (ret CaptureActionEvent) {
	ret.ref = bindings.NewCaptureActionEventByCaptureActionEvent(
		js.Pointer(&init))
	return
}

func NewCaptureActionEventByCaptureActionEvent1() (ret CaptureActionEvent) {
	ret.ref = bindings.NewCaptureActionEventByCaptureActionEvent1()
	return
}

type CaptureActionEvent struct {
	Event
}

func (this CaptureActionEvent) Once() CaptureActionEvent {
	this.Ref().Once()
	return this
}

func (this CaptureActionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CaptureActionEvent) FromRef(ref js.Ref) CaptureActionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CaptureActionEvent) Free() {
	this.Ref().Free()
}

// Action returns the value of property "CaptureActionEvent.action".
//
// It returns ok=false if there is no such property.
func (this CaptureActionEvent) Action() (ret CaptureAction, ok bool) {
	ok = js.True == bindings.GetCaptureActionEventAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CaptureStartFocusBehavior uint32

const (
	_ CaptureStartFocusBehavior = iota

	CaptureStartFocusBehavior_FOCUS_CAPTURING_APPLICATION
	CaptureStartFocusBehavior_FOCUS_CAPTURED_SURFACE
	CaptureStartFocusBehavior_NO_FOCUS_CHANGE
)

func (CaptureStartFocusBehavior) FromRef(str js.Ref) CaptureStartFocusBehavior {
	return CaptureStartFocusBehavior(bindings.ConstOfCaptureStartFocusBehavior(str))
}

func (x CaptureStartFocusBehavior) String() (string, bool) {
	switch x {
	case CaptureStartFocusBehavior_FOCUS_CAPTURING_APPLICATION:
		return "focus-capturing-application", true
	case CaptureStartFocusBehavior_FOCUS_CAPTURED_SURFACE:
		return "focus-captured-surface", true
	case CaptureStartFocusBehavior_NO_FOCUS_CHANGE:
		return "no-focus-change", true
	default:
		return "", false
	}
}

type CaptureController struct {
	EventTarget
}

func (this CaptureController) Once() CaptureController {
	this.Ref().Once()
	return this
}

func (this CaptureController) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this CaptureController) FromRef(ref js.Ref) CaptureController {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this CaptureController) Free() {
	this.Ref().Free()
}

// HasSetFocusBehavior returns true if the method "CaptureController.setFocusBehavior" exists.
func (this CaptureController) HasSetFocusBehavior() bool {
	return js.True == bindings.HasCaptureControllerSetFocusBehavior(
		this.Ref(),
	)
}

// SetFocusBehaviorFunc returns the method "CaptureController.setFocusBehavior".
func (this CaptureController) SetFocusBehaviorFunc() (fn js.Func[func(focusBehavior CaptureStartFocusBehavior)]) {
	return fn.FromRef(
		bindings.CaptureControllerSetFocusBehaviorFunc(
			this.Ref(),
		),
	)
}

// SetFocusBehavior calls the method "CaptureController.setFocusBehavior".
func (this CaptureController) SetFocusBehavior(focusBehavior CaptureStartFocusBehavior) (ret js.Void) {
	bindings.CallCaptureControllerSetFocusBehavior(
		this.Ref(), js.Pointer(&ret),
		uint32(focusBehavior),
	)

	return
}

// TrySetFocusBehavior calls the method "CaptureController.setFocusBehavior"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CaptureController) TrySetFocusBehavior(focusBehavior CaptureStartFocusBehavior) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaptureControllerSetFocusBehavior(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(focusBehavior),
	)

	return
}

type CaptureHandleConfig struct {
	// ExposeOrigin is "CaptureHandleConfig.exposeOrigin"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExposeOrigin MUST be set to true to make this field effective.
	ExposeOrigin bool
	// Handle is "CaptureHandleConfig.handle"
	//
	// Optional, defaults to "".
	Handle js.String
	// PermittedOrigins is "CaptureHandleConfig.permittedOrigins"
	//
	// Optional, defaults to [].
	PermittedOrigins js.Array[js.String]

	FFI_USE_ExposeOrigin bool // for ExposeOrigin.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CaptureHandleConfig with all fields set.
func (p CaptureHandleConfig) FromRef(ref js.Ref) CaptureHandleConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CaptureHandleConfig in the application heap.
func (p CaptureHandleConfig) New() js.Ref {
	return bindings.CaptureHandleConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CaptureHandleConfig) UpdateFrom(ref js.Ref) {
	bindings.CaptureHandleConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CaptureHandleConfig) Update(ref js.Ref) {
	bindings.CaptureHandleConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CapturedMouseEventInit struct {
	// SurfaceX is "CapturedMouseEventInit.surfaceX"
	//
	// Optional, defaults to -1.
	//
	// NOTE: FFI_USE_SurfaceX MUST be set to true to make this field effective.
	SurfaceX int32
	// SurfaceY is "CapturedMouseEventInit.surfaceY"
	//
	// Optional, defaults to -1.
	//
	// NOTE: FFI_USE_SurfaceY MUST be set to true to make this field effective.
	SurfaceY int32
	// Bubbles is "CapturedMouseEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CapturedMouseEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CapturedMouseEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_SurfaceX   bool // for SurfaceX.
	FFI_USE_SurfaceY   bool // for SurfaceY.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CapturedMouseEventInit with all fields set.
func (p CapturedMouseEventInit) FromRef(ref js.Ref) CapturedMouseEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CapturedMouseEventInit in the application heap.
func (p CapturedMouseEventInit) New() js.Ref {
	return bindings.CapturedMouseEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CapturedMouseEventInit) UpdateFrom(ref js.Ref) {
	bindings.CapturedMouseEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CapturedMouseEventInit) Update(ref js.Ref) {
	bindings.CapturedMouseEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCapturedMouseEvent(typ js.String, eventInitDict CapturedMouseEventInit) (ret CapturedMouseEvent) {
	ret.ref = bindings.NewCapturedMouseEventByCapturedMouseEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewCapturedMouseEventByCapturedMouseEvent1(typ js.String) (ret CapturedMouseEvent) {
	ret.ref = bindings.NewCapturedMouseEventByCapturedMouseEvent1(
		typ.Ref())
	return
}

type CapturedMouseEvent struct {
	Event
}

func (this CapturedMouseEvent) Once() CapturedMouseEvent {
	this.Ref().Once()
	return this
}

func (this CapturedMouseEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CapturedMouseEvent) FromRef(ref js.Ref) CapturedMouseEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CapturedMouseEvent) Free() {
	this.Ref().Free()
}

// SurfaceX returns the value of property "CapturedMouseEvent.surfaceX".
//
// It returns ok=false if there is no such property.
func (this CapturedMouseEvent) SurfaceX() (ret int32, ok bool) {
	ok = js.True == bindings.GetCapturedMouseEventSurfaceX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SurfaceY returns the value of property "CapturedMouseEvent.surfaceY".
//
// It returns ok=false if there is no such property.
func (this CapturedMouseEvent) SurfaceY() (ret int32, ok bool) {
	ok = js.True == bindings.GetCapturedMouseEventSurfaceY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CharacterBoundsUpdateEventInit struct {
	// RangeStart is "CharacterBoundsUpdateEventInit.rangeStart"
	//
	// Optional
	//
	// NOTE: FFI_USE_RangeStart MUST be set to true to make this field effective.
	RangeStart uint32
	// RangeEnd is "CharacterBoundsUpdateEventInit.rangeEnd"
	//
	// Optional
	//
	// NOTE: FFI_USE_RangeEnd MUST be set to true to make this field effective.
	RangeEnd uint32
	// Bubbles is "CharacterBoundsUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "CharacterBoundsUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "CharacterBoundsUpdateEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_RangeStart bool // for RangeStart.
	FFI_USE_RangeEnd   bool // for RangeEnd.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CharacterBoundsUpdateEventInit with all fields set.
func (p CharacterBoundsUpdateEventInit) FromRef(ref js.Ref) CharacterBoundsUpdateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CharacterBoundsUpdateEventInit in the application heap.
func (p CharacterBoundsUpdateEventInit) New() js.Ref {
	return bindings.CharacterBoundsUpdateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CharacterBoundsUpdateEventInit) UpdateFrom(ref js.Ref) {
	bindings.CharacterBoundsUpdateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CharacterBoundsUpdateEventInit) Update(ref js.Ref) {
	bindings.CharacterBoundsUpdateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewCharacterBoundsUpdateEvent(typ js.String, options CharacterBoundsUpdateEventInit) (ret CharacterBoundsUpdateEvent) {
	ret.ref = bindings.NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent(
		typ.Ref(),
		js.Pointer(&options))
	return
}

func NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent1(typ js.String) (ret CharacterBoundsUpdateEvent) {
	ret.ref = bindings.NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent1(
		typ.Ref())
	return
}

type CharacterBoundsUpdateEvent struct {
	Event
}

func (this CharacterBoundsUpdateEvent) Once() CharacterBoundsUpdateEvent {
	this.Ref().Once()
	return this
}

func (this CharacterBoundsUpdateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this CharacterBoundsUpdateEvent) FromRef(ref js.Ref) CharacterBoundsUpdateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this CharacterBoundsUpdateEvent) Free() {
	this.Ref().Free()
}

// RangeStart returns the value of property "CharacterBoundsUpdateEvent.rangeStart".
//
// It returns ok=false if there is no such property.
func (this CharacterBoundsUpdateEvent) RangeStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCharacterBoundsUpdateEventRangeStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RangeEnd returns the value of property "CharacterBoundsUpdateEvent.rangeEnd".
//
// It returns ok=false if there is no such property.
func (this CharacterBoundsUpdateEvent) RangeEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCharacterBoundsUpdateEventRangeEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type CharacterData struct {
	Node
}

func (this CharacterData) Once() CharacterData {
	this.Ref().Once()
	return this
}

func (this CharacterData) Ref() js.Ref {
	return this.Node.Ref()
}

func (this CharacterData) FromRef(ref js.Ref) CharacterData {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this CharacterData) Free() {
	this.Ref().Free()
}

// Data returns the value of property "CharacterData.data".
//
// It returns ok=false if there is no such property.
func (this CharacterData) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCharacterDataData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "CharacterData.data" to val.
//
// It returns false if the property cannot be set.
func (this CharacterData) SetData(val js.String) bool {
	return js.True == bindings.SetCharacterDataData(
		this.Ref(),
		val.Ref(),
	)
}

// Length returns the value of property "CharacterData.length".
//
// It returns ok=false if there is no such property.
func (this CharacterData) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCharacterDataLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousElementSibling returns the value of property "CharacterData.previousElementSibling".
//
// It returns ok=false if there is no such property.
func (this CharacterData) PreviousElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetCharacterDataPreviousElementSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextElementSibling returns the value of property "CharacterData.nextElementSibling".
//
// It returns ok=false if there is no such property.
func (this CharacterData) NextElementSibling() (ret Element, ok bool) {
	ok = js.True == bindings.GetCharacterDataNextElementSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSubstringData returns true if the method "CharacterData.substringData" exists.
func (this CharacterData) HasSubstringData() bool {
	return js.True == bindings.HasCharacterDataSubstringData(
		this.Ref(),
	)
}

// SubstringDataFunc returns the method "CharacterData.substringData".
func (this CharacterData) SubstringDataFunc() (fn js.Func[func(offset uint32, count uint32) js.String]) {
	return fn.FromRef(
		bindings.CharacterDataSubstringDataFunc(
			this.Ref(),
		),
	)
}

// SubstringData calls the method "CharacterData.substringData".
func (this CharacterData) SubstringData(offset uint32, count uint32) (ret js.String) {
	bindings.CallCharacterDataSubstringData(
		this.Ref(), js.Pointer(&ret),
		uint32(offset),
		uint32(count),
	)

	return
}

// TrySubstringData calls the method "CharacterData.substringData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TrySubstringData(offset uint32, count uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataSubstringData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
		uint32(count),
	)

	return
}

// HasAppendData returns true if the method "CharacterData.appendData" exists.
func (this CharacterData) HasAppendData() bool {
	return js.True == bindings.HasCharacterDataAppendData(
		this.Ref(),
	)
}

// AppendDataFunc returns the method "CharacterData.appendData".
func (this CharacterData) AppendDataFunc() (fn js.Func[func(data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataAppendDataFunc(
			this.Ref(),
		),
	)
}

// AppendData calls the method "CharacterData.appendData".
func (this CharacterData) AppendData(data js.String) (ret js.Void) {
	bindings.CallCharacterDataAppendData(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAppendData calls the method "CharacterData.appendData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryAppendData(data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataAppendData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasInsertData returns true if the method "CharacterData.insertData" exists.
func (this CharacterData) HasInsertData() bool {
	return js.True == bindings.HasCharacterDataInsertData(
		this.Ref(),
	)
}

// InsertDataFunc returns the method "CharacterData.insertData".
func (this CharacterData) InsertDataFunc() (fn js.Func[func(offset uint32, data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataInsertDataFunc(
			this.Ref(),
		),
	)
}

// InsertData calls the method "CharacterData.insertData".
func (this CharacterData) InsertData(offset uint32, data js.String) (ret js.Void) {
	bindings.CallCharacterDataInsertData(
		this.Ref(), js.Pointer(&ret),
		uint32(offset),
		data.Ref(),
	)

	return
}

// TryInsertData calls the method "CharacterData.insertData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryInsertData(offset uint32, data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataInsertData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
		data.Ref(),
	)

	return
}

// HasDeleteData returns true if the method "CharacterData.deleteData" exists.
func (this CharacterData) HasDeleteData() bool {
	return js.True == bindings.HasCharacterDataDeleteData(
		this.Ref(),
	)
}

// DeleteDataFunc returns the method "CharacterData.deleteData".
func (this CharacterData) DeleteDataFunc() (fn js.Func[func(offset uint32, count uint32)]) {
	return fn.FromRef(
		bindings.CharacterDataDeleteDataFunc(
			this.Ref(),
		),
	)
}

// DeleteData calls the method "CharacterData.deleteData".
func (this CharacterData) DeleteData(offset uint32, count uint32) (ret js.Void) {
	bindings.CallCharacterDataDeleteData(
		this.Ref(), js.Pointer(&ret),
		uint32(offset),
		uint32(count),
	)

	return
}

// TryDeleteData calls the method "CharacterData.deleteData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryDeleteData(offset uint32, count uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataDeleteData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
		uint32(count),
	)

	return
}

// HasReplaceData returns true if the method "CharacterData.replaceData" exists.
func (this CharacterData) HasReplaceData() bool {
	return js.True == bindings.HasCharacterDataReplaceData(
		this.Ref(),
	)
}

// ReplaceDataFunc returns the method "CharacterData.replaceData".
func (this CharacterData) ReplaceDataFunc() (fn js.Func[func(offset uint32, count uint32, data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataReplaceDataFunc(
			this.Ref(),
		),
	)
}

// ReplaceData calls the method "CharacterData.replaceData".
func (this CharacterData) ReplaceData(offset uint32, count uint32, data js.String) (ret js.Void) {
	bindings.CallCharacterDataReplaceData(
		this.Ref(), js.Pointer(&ret),
		uint32(offset),
		uint32(count),
		data.Ref(),
	)

	return
}

// TryReplaceData calls the method "CharacterData.replaceData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryReplaceData(offset uint32, count uint32, data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataReplaceData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(offset),
		uint32(count),
		data.Ref(),
	)

	return
}

// HasBefore returns true if the method "CharacterData.before" exists.
func (this CharacterData) HasBefore() bool {
	return js.True == bindings.HasCharacterDataBefore(
		this.Ref(),
	)
}

// BeforeFunc returns the method "CharacterData.before".
func (this CharacterData) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataBeforeFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "CharacterData.before".
func (this CharacterData) Before(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallCharacterDataBefore(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryBefore calls the method "CharacterData.before"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryBefore(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAfter returns true if the method "CharacterData.after" exists.
func (this CharacterData) HasAfter() bool {
	return js.True == bindings.HasCharacterDataAfter(
		this.Ref(),
	)
}

// AfterFunc returns the method "CharacterData.after".
func (this CharacterData) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataAfterFunc(
			this.Ref(),
		),
	)
}

// After calls the method "CharacterData.after".
func (this CharacterData) After(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallCharacterDataAfter(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAfter calls the method "CharacterData.after"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryAfter(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataAfter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceWith returns true if the method "CharacterData.replaceWith" exists.
func (this CharacterData) HasReplaceWith() bool {
	return js.True == bindings.HasCharacterDataReplaceWith(
		this.Ref(),
	)
}

// ReplaceWithFunc returns the method "CharacterData.replaceWith".
func (this CharacterData) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataReplaceWithFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "CharacterData.replaceWith".
func (this CharacterData) ReplaceWith(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallCharacterDataReplaceWith(
		this.Ref(), js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceWith calls the method "CharacterData.replaceWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryReplaceWith(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataReplaceWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasRemove returns true if the method "CharacterData.remove" exists.
func (this CharacterData) HasRemove() bool {
	return js.True == bindings.HasCharacterDataRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "CharacterData.remove".
func (this CharacterData) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CharacterDataRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "CharacterData.remove".
func (this CharacterData) Remove() (ret js.Void) {
	bindings.CallCharacterDataRemove(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "CharacterData.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CharacterData) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCharacterDataRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ChildDisplayType uint32

const (
	_ ChildDisplayType = iota

	ChildDisplayType_BLOCK
	ChildDisplayType_NORMAL
)

func (ChildDisplayType) FromRef(str js.Ref) ChildDisplayType {
	return ChildDisplayType(bindings.ConstOfChildDisplayType(str))
}

func (x ChildDisplayType) String() (string, bool) {
	switch x {
	case ChildDisplayType_BLOCK:
		return "block", true
	case ChildDisplayType_NORMAL:
		return "normal", true
	default:
		return "", false
	}
}

type FrameType uint32

const (
	_ FrameType = iota

	FrameType_AUXILIARY
	FrameType_TOP_LEVEL
	FrameType_NESTED
	FrameType_NONE
)

func (FrameType) FromRef(str js.Ref) FrameType {
	return FrameType(bindings.ConstOfFrameType(str))
}

func (x FrameType) String() (string, bool) {
	switch x {
	case FrameType_AUXILIARY:
		return "auxiliary", true
	case FrameType_TOP_LEVEL:
		return "top-level", true
	case FrameType_NESTED:
		return "nested", true
	case FrameType_NONE:
		return "none", true
	default:
		return "", false
	}
}

type ClientType uint32

const (
	_ ClientType = iota

	ClientType_WINDOW
	ClientType_WORKER
	ClientType_SHAREDWORKER
	ClientType_ALL
)

func (ClientType) FromRef(str js.Ref) ClientType {
	return ClientType(bindings.ConstOfClientType(str))
}

func (x ClientType) String() (string, bool) {
	switch x {
	case ClientType_WINDOW:
		return "window", true
	case ClientType_WORKER:
		return "worker", true
	case ClientType_SHAREDWORKER:
		return "sharedworker", true
	case ClientType_ALL:
		return "all", true
	default:
		return "", false
	}
}

type ClientLifecycleState uint32

const (
	_ ClientLifecycleState = iota

	ClientLifecycleState_ACTIVE
	ClientLifecycleState_FROZEN
)

func (ClientLifecycleState) FromRef(str js.Ref) ClientLifecycleState {
	return ClientLifecycleState(bindings.ConstOfClientLifecycleState(str))
}

func (x ClientLifecycleState) String() (string, bool) {
	switch x {
	case ClientLifecycleState_ACTIVE:
		return "active", true
	case ClientLifecycleState_FROZEN:
		return "frozen", true
	default:
		return "", false
	}
}

type Client struct {
	ref js.Ref
}

func (this Client) Once() Client {
	this.Ref().Once()
	return this
}

func (this Client) Ref() js.Ref {
	return this.ref
}

func (this Client) FromRef(ref js.Ref) Client {
	this.ref = ref
	return this
}

func (this Client) Free() {
	this.Ref().Free()
}

// Url returns the value of property "Client.url".
//
// It returns ok=false if there is no such property.
func (this Client) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetClientUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FrameType returns the value of property "Client.frameType".
//
// It returns ok=false if there is no such property.
func (this Client) FrameType() (ret FrameType, ok bool) {
	ok = js.True == bindings.GetClientFrameType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "Client.id".
//
// It returns ok=false if there is no such property.
func (this Client) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetClientId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "Client.type".
//
// It returns ok=false if there is no such property.
func (this Client) Type() (ret ClientType, ok bool) {
	ok = js.True == bindings.GetClientType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LifecycleState returns the value of property "Client.lifecycleState".
//
// It returns ok=false if there is no such property.
func (this Client) LifecycleState() (ret ClientLifecycleState, ok bool) {
	ok = js.True == bindings.GetClientLifecycleState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPostMessage returns true if the method "Client.postMessage" exists.
func (this Client) HasPostMessage() bool {
	return js.True == bindings.HasClientPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "Client.postMessage".
func (this Client) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.ClientPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "Client.postMessage".
func (this Client) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallClientPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "Client.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Client) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "Client.postMessage" exists.
func (this Client) HasPostMessage1() bool {
	return js.True == bindings.HasClientPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "Client.postMessage".
func (this Client) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.ClientPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Client.postMessage".
func (this Client) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallClientPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "Client.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Client) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage2 returns true if the method "Client.postMessage" exists.
func (this Client) HasPostMessage2() bool {
	return js.True == bindings.HasClientPostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "Client.postMessage".
func (this Client) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.ClientPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Client.postMessage".
func (this Client) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallClientPostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "Client.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Client) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClientPostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

type ClientQueryOptions struct {
	// IncludeUncontrolled is "ClientQueryOptions.includeUncontrolled"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IncludeUncontrolled MUST be set to true to make this field effective.
	IncludeUncontrolled bool
	// Type is "ClientQueryOptions.type"
	//
	// Optional, defaults to "window".
	Type ClientType

	FFI_USE_IncludeUncontrolled bool // for IncludeUncontrolled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClientQueryOptions with all fields set.
func (p ClientQueryOptions) FromRef(ref js.Ref) ClientQueryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClientQueryOptions in the application heap.
func (p ClientQueryOptions) New() js.Ref {
	return bindings.ClientQueryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ClientQueryOptions) UpdateFrom(ref js.Ref) {
	bindings.ClientQueryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ClientQueryOptions) Update(ref js.Ref) {
	bindings.ClientQueryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
