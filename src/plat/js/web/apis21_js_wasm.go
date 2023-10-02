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

func NewCSSPerspective(length CSSPerspectiveValue) CSSPerspective {
	return CSSPerspective{}.FromRef(
		bindings.NewCSSPerspectiveByCSSPerspective(
			length.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSPerspective) Length() (CSSPerspectiveValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSPerspectiveLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSPerspectiveValue{}.FromRef(_ret), _ok
}

// Length sets the value of property "CSSPerspective.length" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSPositionFallbackRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPositionFallbackRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this CSSPropertyRule) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPropertyRuleName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Syntax returns the value of property "CSSPropertyRule.syntax".
//
// The returned bool will be false if there is no such property.
func (this CSSPropertyRule) Syntax() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPropertyRuleSyntax(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Inherits returns the value of property "CSSPropertyRule.inherits".
//
// The returned bool will be false if there is no such property.
func (this CSSPropertyRule) Inherits() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCSSPropertyRuleInherits(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// InitialValue returns the value of property "CSSPropertyRule.initialValue".
//
// The returned bool will be false if there is no such property.
func (this CSSPropertyRule) InitialValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSPropertyRuleInitialValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

func NewCSSRGB(r CSSColorRGBComp, g CSSColorRGBComp, b CSSColorRGBComp, alpha CSSColorPercent) CSSRGB {
	return CSSRGB{}.FromRef(
		bindings.NewCSSRGBByCSSRGB(
			r.Ref(),
			g.Ref(),
			b.Ref(),
			alpha.Ref()),
	)
}

func NewCSSRGBByCSSRGB1(r CSSColorRGBComp, g CSSColorRGBComp, b CSSColorRGBComp) CSSRGB {
	return CSSRGB{}.FromRef(
		bindings.NewCSSRGBByCSSRGB1(
			r.Ref(),
			g.Ref(),
			b.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSRGB) R() (CSSColorRGBComp, bool) {
	var _ok bool
	_ret := bindings.GetCSSRGBR(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorRGBComp{}.FromRef(_ret), _ok
}

// R sets the value of property "CSSRGB.r" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRGB) G() (CSSColorRGBComp, bool) {
	var _ok bool
	_ret := bindings.GetCSSRGBG(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorRGBComp{}.FromRef(_ret), _ok
}

// G sets the value of property "CSSRGB.g" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRGB) B() (CSSColorRGBComp, bool) {
	var _ok bool
	_ret := bindings.GetCSSRGBB(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorRGBComp{}.FromRef(_ret), _ok
}

// B sets the value of property "CSSRGB.b" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRGB) Alpha() (CSSColorPercent, bool) {
	var _ok bool
	_ret := bindings.GetCSSRGBAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSColorPercent{}.FromRef(_ret), _ok
}

// Alpha sets the value of property "CSSRGB.alpha" to val.
//
// It returns false if the property cannot be set.
func (this CSSRGB) SetAlpha(val CSSColorPercent) bool {
	return js.True == bindings.SetCSSRGBAlpha(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSRotate(angle CSSNumericValue) CSSRotate {
	return CSSRotate{}.FromRef(
		bindings.NewCSSRotateByCSSRotate(
			angle.Ref()),
	)
}

func NewCSSRotateByCSSRotate1(x CSSNumberish, y CSSNumberish, z CSSNumberish, angle CSSNumericValue) CSSRotate {
	return CSSRotate{}.FromRef(
		bindings.NewCSSRotateByCSSRotate1(
			x.Ref(),
			y.Ref(),
			z.Ref(),
			angle.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSRotate) X() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSRotateX(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// X sets the value of property "CSSRotate.x" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRotate) Y() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSRotateY(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Y sets the value of property "CSSRotate.y" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRotate) Z() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSRotateZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Z sets the value of property "CSSRotate.z" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSRotate) Angle() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSRotateAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Angle sets the value of property "CSSRotate.angle" to val.
//
// It returns false if the property cannot be set.
func (this CSSRotate) SetAngle(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSRotateAngle(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSScale(x CSSNumberish, y CSSNumberish, z CSSNumberish) CSSScale {
	return CSSScale{}.FromRef(
		bindings.NewCSSScaleByCSSScale(
			x.Ref(),
			y.Ref(),
			z.Ref()),
	)
}

func NewCSSScaleByCSSScale1(x CSSNumberish, y CSSNumberish) CSSScale {
	return CSSScale{}.FromRef(
		bindings.NewCSSScaleByCSSScale1(
			x.Ref(),
			y.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSScale) X() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSScaleX(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// X sets the value of property "CSSScale.x" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSScale) Y() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSScaleY(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Y sets the value of property "CSSScale.y" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSScale) Z() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetCSSScaleZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// Z sets the value of property "CSSScale.z" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSScopeRule) Start() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSScopeRuleStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// End returns the value of property "CSSScopeRule.end".
//
// The returned bool will be false if there is no such property.
func (this CSSScopeRule) End() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSScopeRuleEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

func NewCSSSkew(ax CSSNumericValue, ay CSSNumericValue) CSSSkew {
	return CSSSkew{}.FromRef(
		bindings.NewCSSSkewByCSSSkew(
			ax.Ref(),
			ay.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSSkew) Ax() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSSkewAx(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Ax sets the value of property "CSSSkew.ax" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSSkew) Ay() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSSkewAy(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Ay sets the value of property "CSSSkew.ay" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkew) SetAy(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewAy(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSSkewX(ax CSSNumericValue) CSSSkewX {
	return CSSSkewX{}.FromRef(
		bindings.NewCSSSkewXByCSSSkewX(
			ax.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSSkewX) Ax() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSSkewXAx(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Ax sets the value of property "CSSSkewX.ax" to val.
//
// It returns false if the property cannot be set.
func (this CSSSkewX) SetAx(val CSSNumericValue) bool {
	return js.True == bindings.SetCSSSkewXAx(
		this.Ref(),
		val.Ref(),
	)
}

func NewCSSSkewY(ay CSSNumericValue) CSSSkewY {
	return CSSSkewY{}.FromRef(
		bindings.NewCSSSkewYByCSSSkewY(
			ay.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSSkewY) Ay() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSSkewYAy(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Ay sets the value of property "CSSSkewY.ay" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSStyleRule) SelectorText() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleRuleSelectorText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectorText sets the value of property "CSSStyleRule.selectorText" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSStyleRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// StyleMap returns the value of property "CSSStyleRule.styleMap".
//
// The returned bool will be false if there is no such property.
func (this CSSStyleRule) StyleMap() (StylePropertyMap, bool) {
	var _ok bool
	_ret := bindings.GetCSSStyleRuleStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return StylePropertyMap{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this CSSTransformComponent) Is2D() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCSSTransformComponentIs2D(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Is2D sets the value of property "CSSTransformComponent.is2D" to val.
//
// It returns false if the property cannot be set.
func (this CSSTransformComponent) SetIs2D(val bool) bool {
	return js.True == bindings.SetCSSTransformComponentIs2D(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ToString calls the method "CSSTransformComponent.toString".
//
// The returned bool will be false if there is no such method.
func (this CSSTransformComponent) ToString() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCSSTransformComponentToString(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToStringFunc returns the method "CSSTransformComponent.toString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSTransformComponent) ToStringFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.CSSTransformComponentToStringFunc(
			this.Ref(),
		),
	)
}

// ToMatrix calls the method "CSSTransformComponent.toMatrix".
//
// The returned bool will be false if there is no such method.
func (this CSSTransformComponent) ToMatrix() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallCSSTransformComponentToMatrix(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ToMatrixFunc returns the method "CSSTransformComponent.toMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSTransformComponent) ToMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CSSTransformComponentToMatrixFunc(
			this.Ref(),
		),
	)
}

func NewCSSTransformValue(transforms js.Array[CSSTransformComponent]) CSSTransformValue {
	return CSSTransformValue{}.FromRef(
		bindings.NewCSSTransformValueByCSSTransformValue(
			transforms.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSTransformValue) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSTransformValueLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Is2D returns the value of property "CSSTransformValue.is2D".
//
// The returned bool will be false if there is no such property.
func (this CSSTransformValue) Is2D() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCSSTransformValueIs2D(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Get calls the method "CSSTransformValue.".
//
// The returned bool will be false if there is no such method.
func (this CSSTransformValue) Get(index uint32) (CSSTransformComponent, bool) {
	var _ok bool
	_ret := bindings.CallCSSTransformValueGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSTransformComponent{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CSSTransformValue.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSTransformValue) GetFunc() (fn js.Func[func(index uint32) CSSTransformComponent]) {
	return fn.FromRef(
		bindings.CSSTransformValueGetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CSSTransformValue.".
//
// The returned bool will be false if there is no such method.
func (this CSSTransformValue) Set(index uint32, val CSSTransformComponent) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSTransformValueSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		val.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "CSSTransformValue.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSTransformValue) SetFunc() (fn js.Func[func(index uint32, val CSSTransformComponent)]) {
	return fn.FromRef(
		bindings.CSSTransformValueSetFunc(
			this.Ref(),
		),
	)
}

// ToMatrix calls the method "CSSTransformValue.toMatrix".
//
// The returned bool will be false if there is no such method.
func (this CSSTransformValue) ToMatrix() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallCSSTransformValueToMatrix(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// ToMatrixFunc returns the method "CSSTransformValue.toMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSTransformValue) ToMatrixFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CSSTransformValueToMatrixFunc(
			this.Ref(),
		),
	)
}

func NewCSSTransition(effect AnimationEffect, timeline AnimationTimeline) CSSTransition {
	return CSSTransition{}.FromRef(
		bindings.NewCSSTransitionByCSSTransition(
			effect.Ref(),
			timeline.Ref()),
	)
}

func NewCSSTransitionByCSSTransition1(effect AnimationEffect) CSSTransition {
	return CSSTransition{}.FromRef(
		bindings.NewCSSTransitionByCSSTransition1(
			effect.Ref()),
	)
}

func NewCSSTransitionByCSSTransition2() CSSTransition {
	return CSSTransition{}.FromRef(
		bindings.NewCSSTransitionByCSSTransition2(),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSTransition) TransitionProperty() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSTransitionTransitionProperty(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

func NewCSSTranslate(x CSSNumericValue, y CSSNumericValue, z CSSNumericValue) CSSTranslate {
	return CSSTranslate{}.FromRef(
		bindings.NewCSSTranslateByCSSTranslate(
			x.Ref(),
			y.Ref(),
			z.Ref()),
	)
}

func NewCSSTranslateByCSSTranslate1(x CSSNumericValue, y CSSNumericValue) CSSTranslate {
	return CSSTranslate{}.FromRef(
		bindings.NewCSSTranslateByCSSTranslate1(
			x.Ref(),
			y.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSTranslate) X() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSTranslateX(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// X sets the value of property "CSSTranslate.x" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSTranslate) Y() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSTranslateY(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Y sets the value of property "CSSTranslate.y" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSTranslate) Z() (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSTranslateZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericValue{}.FromRef(_ret), _ok
}

// Z sets the value of property "CSSTranslate.z" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSTryRule) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetCSSTryRuleStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

func NewCSSUnparsedValue(members js.Array[CSSUnparsedSegment]) CSSUnparsedValue {
	return CSSUnparsedValue{}.FromRef(
		bindings.NewCSSUnparsedValueByCSSUnparsedValue(
			members.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSUnparsedValue) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSUnparsedValueLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "CSSUnparsedValue.".
//
// The returned bool will be false if there is no such method.
func (this CSSUnparsedValue) Get(index uint32) (CSSUnparsedSegment, bool) {
	var _ok bool
	_ret := bindings.CallCSSUnparsedValueGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSUnparsedSegment{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CSSUnparsedValue.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSUnparsedValue) GetFunc() (fn js.Func[func(index uint32) CSSUnparsedSegment]) {
	return fn.FromRef(
		bindings.CSSUnparsedValueGetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "CSSUnparsedValue.".
//
// The returned bool will be false if there is no such method.
func (this CSSUnparsedValue) Set(index uint32, val CSSUnparsedSegment) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCSSUnparsedValueSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		val.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "CSSUnparsedValue.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSUnparsedValue) SetFunc() (fn js.Func[func(index uint32, val CSSUnparsedSegment)]) {
	return fn.FromRef(
		bindings.CSSUnparsedValueSetFunc(
			this.Ref(),
		),
	)
}

func NewCSSVariableReferenceValue(variable js.String, fallback CSSUnparsedValue) CSSVariableReferenceValue {
	return CSSVariableReferenceValue{}.FromRef(
		bindings.NewCSSVariableReferenceValueByCSSVariableReferenceValue(
			variable.Ref(),
			fallback.Ref()),
	)
}

func NewCSSVariableReferenceValueByCSSVariableReferenceValue1(variable js.String) CSSVariableReferenceValue {
	return CSSVariableReferenceValue{}.FromRef(
		bindings.NewCSSVariableReferenceValueByCSSVariableReferenceValue1(
			variable.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CSSVariableReferenceValue) Variable() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSVariableReferenceValueVariable(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Variable sets the value of property "CSSVariableReferenceValue.variable" to val.
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
// The returned bool will be false if there is no such property.
func (this CSSVariableReferenceValue) Fallback() (CSSUnparsedValue, bool) {
	var _ok bool
	_ret := bindings.GetCSSVariableReferenceValueFallback(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSUnparsedValue{}.FromRef(_ret), _ok
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

// Match calls the method "Cache.match".
//
// The returned bool will be false if there is no such method.
func (this Cache) Match(request RequestInfo, options CacheQueryOptions) (js.Promise[OneOf_Response_undefined], bool) {
	var _ok bool
	_ret := bindings.CallCacheMatch(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[OneOf_Response_undefined]{}.FromRef(_ret), _ok
}

// MatchFunc returns the method "Cache.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) MatchFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheMatchFunc(
			this.Ref(),
		),
	)
}

// Match1 calls the method "Cache.match".
//
// The returned bool will be false if there is no such method.
func (this Cache) Match1(request RequestInfo) (js.Promise[OneOf_Response_undefined], bool) {
	var _ok bool
	_ret := bindings.CallCacheMatch1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[OneOf_Response_undefined]{}.FromRef(_ret), _ok
}

// Match1Func returns the method "Cache.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheMatch1Func(
			this.Ref(),
		),
	)
}

// MatchAll calls the method "Cache.matchAll".
//
// The returned bool will be false if there is no such method.
func (this Cache) MatchAll(request RequestInfo, options CacheQueryOptions) (js.Promise[js.FrozenArray[Response]], bool) {
	var _ok bool
	_ret := bindings.CallCacheMatchAll(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.FrozenArray[Response]]{}.FromRef(_ret), _ok
}

// MatchAllFunc returns the method "Cache.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) MatchAllFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAllFunc(
			this.Ref(),
		),
	)
}

// MatchAll1 calls the method "Cache.matchAll".
//
// The returned bool will be false if there is no such method.
func (this Cache) MatchAll1(request RequestInfo) (js.Promise[js.FrozenArray[Response]], bool) {
	var _ok bool
	_ret := bindings.CallCacheMatchAll1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[js.FrozenArray[Response]]{}.FromRef(_ret), _ok
}

// MatchAll1Func returns the method "Cache.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) MatchAll1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAll1Func(
			this.Ref(),
		),
	)
}

// MatchAll2 calls the method "Cache.matchAll".
//
// The returned bool will be false if there is no such method.
func (this Cache) MatchAll2() (js.Promise[js.FrozenArray[Response]], bool) {
	var _ok bool
	_ret := bindings.CallCacheMatchAll2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.FrozenArray[Response]]{}.FromRef(_ret), _ok
}

// MatchAll2Func returns the method "Cache.matchAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) MatchAll2Func() (fn js.Func[func() js.Promise[js.FrozenArray[Response]]]) {
	return fn.FromRef(
		bindings.CacheMatchAll2Func(
			this.Ref(),
		),
	)
}

// Add calls the method "Cache.add".
//
// The returned bool will be false if there is no such method.
func (this Cache) Add(request RequestInfo) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCacheAdd(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddFunc returns the method "Cache.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) AddFunc() (fn js.Func[func(request RequestInfo) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CacheAddFunc(
			this.Ref(),
		),
	)
}

// AddAll calls the method "Cache.addAll".
//
// The returned bool will be false if there is no such method.
func (this Cache) AddAll(requests js.Array[RequestInfo]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCacheAddAll(
		this.Ref(), js.Pointer(&_ok),
		requests.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddAllFunc returns the method "Cache.addAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) AddAllFunc() (fn js.Func[func(requests js.Array[RequestInfo]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CacheAddAllFunc(
			this.Ref(),
		),
	)
}

// Put calls the method "Cache.put".
//
// The returned bool will be false if there is no such method.
func (this Cache) Put(request RequestInfo, response Response) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCachePut(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		response.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PutFunc returns the method "Cache.put".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) PutFunc() (fn js.Func[func(request RequestInfo, response Response) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CachePutFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "Cache.delete".
//
// The returned bool will be false if there is no such method.
func (this Cache) Delete(request RequestInfo, options CacheQueryOptions) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallCacheDelete(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "Cache.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) DeleteFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete1 calls the method "Cache.delete".
//
// The returned bool will be false if there is no such method.
func (this Cache) Delete1(request RequestInfo) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallCacheDelete1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// Delete1Func returns the method "Cache.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) Delete1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheDelete1Func(
			this.Ref(),
		),
	)
}

// Keys calls the method "Cache.keys".
//
// The returned bool will be false if there is no such method.
func (this Cache) Keys(request RequestInfo, options CacheQueryOptions) (js.Promise[js.FrozenArray[Request]], bool) {
	var _ok bool
	_ret := bindings.CallCacheKeys(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.FrozenArray[Request]]{}.FromRef(_ret), _ok
}

// KeysFunc returns the method "Cache.keys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) KeysFunc() (fn js.Func[func(request RequestInfo, options CacheQueryOptions) js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeysFunc(
			this.Ref(),
		),
	)
}

// Keys1 calls the method "Cache.keys".
//
// The returned bool will be false if there is no such method.
func (this Cache) Keys1(request RequestInfo) (js.Promise[js.FrozenArray[Request]], bool) {
	var _ok bool
	_ret := bindings.CallCacheKeys1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[js.FrozenArray[Request]]{}.FromRef(_ret), _ok
}

// Keys1Func returns the method "Cache.keys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) Keys1Func() (fn js.Func[func(request RequestInfo) js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeys1Func(
			this.Ref(),
		),
	)
}

// Keys2 calls the method "Cache.keys".
//
// The returned bool will be false if there is no such method.
func (this Cache) Keys2() (js.Promise[js.FrozenArray[Request]], bool) {
	var _ok bool
	_ret := bindings.CallCacheKeys2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.FrozenArray[Request]]{}.FromRef(_ret), _ok
}

// Keys2Func returns the method "Cache.keys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Cache) Keys2Func() (fn js.Func[func() js.Promise[js.FrozenArray[Request]]]) {
	return fn.FromRef(
		bindings.CacheKeys2Func(
			this.Ref(),
		),
	)
}

type MultiCacheQueryOptions struct {
	// CacheName is "MultiCacheQueryOptions.cacheName"
	//
	// Optional
	CacheName js.String
	// IgnoreSearch is "MultiCacheQueryOptions.ignoreSearch"
	//
	// Optional, defaults to false.
	IgnoreSearch bool
	// IgnoreMethod is "MultiCacheQueryOptions.ignoreMethod"
	//
	// Optional, defaults to false.
	IgnoreMethod bool
	// IgnoreVary is "MultiCacheQueryOptions.ignoreVary"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 MultiCacheQueryOptions MultiCacheQueryOptions [// MultiCacheQueryOptions] [0x14000543ea0 0x14000543f40 0x140005be0a0 0x140005be1e0 0x140005be000 0x140005be140 0x140005be280] 0x14000574258 {0 0}} in the application heap.
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

// Match calls the method "CacheStorage.match".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Match(request RequestInfo, options MultiCacheQueryOptions) (js.Promise[OneOf_Response_undefined], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageMatch(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[OneOf_Response_undefined]{}.FromRef(_ret), _ok
}

// MatchFunc returns the method "CacheStorage.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) MatchFunc() (fn js.Func[func(request RequestInfo, options MultiCacheQueryOptions) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheStorageMatchFunc(
			this.Ref(),
		),
	)
}

// Match1 calls the method "CacheStorage.match".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Match1(request RequestInfo) (js.Promise[OneOf_Response_undefined], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageMatch1(
		this.Ref(), js.Pointer(&_ok),
		request.Ref(),
	)

	return js.Promise[OneOf_Response_undefined]{}.FromRef(_ret), _ok
}

// Match1Func returns the method "CacheStorage.match".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) Match1Func() (fn js.Func[func(request RequestInfo) js.Promise[OneOf_Response_undefined]]) {
	return fn.FromRef(
		bindings.CacheStorageMatch1Func(
			this.Ref(),
		),
	)
}

// Has calls the method "CacheStorage.has".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Has(cacheName js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageHas(
		this.Ref(), js.Pointer(&_ok),
		cacheName.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// HasFunc returns the method "CacheStorage.has".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) HasFunc() (fn js.Func[func(cacheName js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheStorageHasFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "CacheStorage.open".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Open(cacheName js.String) (js.Promise[Cache], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageOpen(
		this.Ref(), js.Pointer(&_ok),
		cacheName.Ref(),
	)

	return js.Promise[Cache]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "CacheStorage.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) OpenFunc() (fn js.Func[func(cacheName js.String) js.Promise[Cache]]) {
	return fn.FromRef(
		bindings.CacheStorageOpenFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "CacheStorage.delete".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Delete(cacheName js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageDelete(
		this.Ref(), js.Pointer(&_ok),
		cacheName.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "CacheStorage.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) DeleteFunc() (fn js.Func[func(cacheName js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.CacheStorageDeleteFunc(
			this.Ref(),
		),
	)
}

// Keys calls the method "CacheStorage.keys".
//
// The returned bool will be false if there is no such method.
func (this CacheStorage) Keys() (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallCacheStorageKeys(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// KeysFunc returns the method "CacheStorage.keys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CacheStorage) KeysFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.CacheStorageKeysFunc(
			this.Ref(),
		),
	)
}

type CameraDevicePermissionDescriptor struct {
	// PanTiltZoom is "CameraDevicePermissionDescriptor.panTiltZoom"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CameraDevicePermissionDescriptor CameraDevicePermissionDescriptor [// CameraDevicePermissionDescriptor] [0x140005be320 0x140005be460 0x140005be3c0] 0x14000574300 {0 0}} in the application heap.
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

func NewCanMakePaymentEvent(typ js.String) CanMakePaymentEvent {
	return CanMakePaymentEvent{}.FromRef(
		bindings.NewCanMakePaymentEventByCanMakePaymentEvent(
			typ.Ref()),
	)
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

// RespondWith calls the method "CanMakePaymentEvent.respondWith".
//
// The returned bool will be false if there is no such method.
func (this CanMakePaymentEvent) RespondWith(canMakePaymentResponse js.Promise[js.Boolean]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanMakePaymentEventRespondWith(
		this.Ref(), js.Pointer(&_ok),
		canMakePaymentResponse.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RespondWithFunc returns the method "CanMakePaymentEvent.respondWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanMakePaymentEvent) RespondWithFunc() (fn js.Func[func(canMakePaymentResponse js.Promise[js.Boolean])]) {
	return fn.FromRef(
		bindings.CanMakePaymentEventRespondWithFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this CanvasCaptureMediaStreamTrack) Canvas() (HTMLCanvasElement, bool) {
	var _ok bool
	_ret := bindings.GetCanvasCaptureMediaStreamTrackCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCanvasElement{}.FromRef(_ret), _ok
}

// RequestFrame calls the method "CanvasCaptureMediaStreamTrack.requestFrame".
//
// The returned bool will be false if there is no such method.
func (this CanvasCaptureMediaStreamTrack) RequestFrame() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasCaptureMediaStreamTrackRequestFrame(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestFrameFunc returns the method "CanvasCaptureMediaStreamTrack.requestFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasCaptureMediaStreamTrack) RequestFrameFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasCaptureMediaStreamTrackRequestFrameFunc(
			this.Ref(),
		),
	)
}

type CaptureActionEventInit struct {
	// Action is "CaptureActionEventInit.action"
	//
	// Optional
	Action js.String
	// Bubbles is "CaptureActionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CaptureActionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CaptureActionEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CaptureActionEventInit CaptureActionEventInit [// CaptureActionEventInit] [0x140005be500 0x140005be5a0 0x140005be6e0 0x140005be820 0x140005be640 0x140005be780 0x140005be8c0] 0x14000574360 {0 0}} in the application heap.
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

func NewCaptureActionEvent(init CaptureActionEventInit) CaptureActionEvent {
	return CaptureActionEvent{}.FromRef(
		bindings.NewCaptureActionEventByCaptureActionEvent(
			js.Pointer(&init)),
	)
}

func NewCaptureActionEventByCaptureActionEvent1() CaptureActionEvent {
	return CaptureActionEvent{}.FromRef(
		bindings.NewCaptureActionEventByCaptureActionEvent1(),
	)
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
// The returned bool will be false if there is no such property.
func (this CaptureActionEvent) Action() (CaptureAction, bool) {
	var _ok bool
	_ret := bindings.GetCaptureActionEventAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return CaptureAction(_ret), _ok
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

// SetFocusBehavior calls the method "CaptureController.setFocusBehavior".
//
// The returned bool will be false if there is no such method.
func (this CaptureController) SetFocusBehavior(focusBehavior CaptureStartFocusBehavior) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCaptureControllerSetFocusBehavior(
		this.Ref(), js.Pointer(&_ok),
		uint32(focusBehavior),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFocusBehaviorFunc returns the method "CaptureController.setFocusBehavior".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CaptureController) SetFocusBehaviorFunc() (fn js.Func[func(focusBehavior CaptureStartFocusBehavior)]) {
	return fn.FromRef(
		bindings.CaptureControllerSetFocusBehaviorFunc(
			this.Ref(),
		),
	)
}

type CaptureHandleConfig struct {
	// ExposeOrigin is "CaptureHandleConfig.exposeOrigin"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CaptureHandleConfig CaptureHandleConfig [// CaptureHandleConfig] [0x140005be960 0x140005beaa0 0x140005beb40 0x140005bea00] 0x14000574450 {0 0}} in the application heap.
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
	SurfaceX int32
	// SurfaceY is "CapturedMouseEventInit.surfaceY"
	//
	// Optional, defaults to -1.
	SurfaceY int32
	// Bubbles is "CapturedMouseEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CapturedMouseEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CapturedMouseEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CapturedMouseEventInit CapturedMouseEventInit [// CapturedMouseEventInit] [0x140005bebe0 0x140005bed20 0x140005bee60 0x140005befa0 0x140005bf0e0 0x140005bec80 0x140005bedc0 0x140005bef00 0x140005bf040 0x140005bf180] 0x14000574498 {0 0}} in the application heap.
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

func NewCapturedMouseEvent(typ js.String, eventInitDict CapturedMouseEventInit) CapturedMouseEvent {
	return CapturedMouseEvent{}.FromRef(
		bindings.NewCapturedMouseEventByCapturedMouseEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewCapturedMouseEventByCapturedMouseEvent1(typ js.String) CapturedMouseEvent {
	return CapturedMouseEvent{}.FromRef(
		bindings.NewCapturedMouseEventByCapturedMouseEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CapturedMouseEvent) SurfaceX() (int32, bool) {
	var _ok bool
	_ret := bindings.GetCapturedMouseEventSurfaceX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SurfaceY returns the value of property "CapturedMouseEvent.surfaceY".
//
// The returned bool will be false if there is no such property.
func (this CapturedMouseEvent) SurfaceY() (int32, bool) {
	var _ok bool
	_ret := bindings.GetCapturedMouseEventSurfaceY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

type CharacterBoundsUpdateEventInit struct {
	// RangeStart is "CharacterBoundsUpdateEventInit.rangeStart"
	//
	// Optional
	RangeStart uint32
	// RangeEnd is "CharacterBoundsUpdateEventInit.rangeEnd"
	//
	// Optional
	RangeEnd uint32
	// Bubbles is "CharacterBoundsUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "CharacterBoundsUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "CharacterBoundsUpdateEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 CharacterBoundsUpdateEventInit CharacterBoundsUpdateEventInit [// CharacterBoundsUpdateEventInit] [0x140005bf2c0 0x140005bf400 0x140005bf540 0x140005bf680 0x140005bf7c0 0x140005bf360 0x140005bf4a0 0x140005bf5e0 0x140005bf720 0x140005bf860] 0x140005744f8 {0 0}} in the application heap.
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

func NewCharacterBoundsUpdateEvent(typ js.String, options CharacterBoundsUpdateEventInit) CharacterBoundsUpdateEvent {
	return CharacterBoundsUpdateEvent{}.FromRef(
		bindings.NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent(
			typ.Ref(),
			js.Pointer(&options)),
	)
}

func NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent1(typ js.String) CharacterBoundsUpdateEvent {
	return CharacterBoundsUpdateEvent{}.FromRef(
		bindings.NewCharacterBoundsUpdateEventByCharacterBoundsUpdateEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this CharacterBoundsUpdateEvent) RangeStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCharacterBoundsUpdateEventRangeStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// RangeEnd returns the value of property "CharacterBoundsUpdateEvent.rangeEnd".
//
// The returned bool will be false if there is no such property.
func (this CharacterBoundsUpdateEvent) RangeEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCharacterBoundsUpdateEventRangeEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this CharacterData) Data() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCharacterDataData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Data sets the value of property "CharacterData.data" to val.
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
// The returned bool will be false if there is no such property.
func (this CharacterData) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCharacterDataLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// PreviousElementSibling returns the value of property "CharacterData.previousElementSibling".
//
// The returned bool will be false if there is no such property.
func (this CharacterData) PreviousElementSibling() (Element, bool) {
	var _ok bool
	_ret := bindings.GetCharacterDataPreviousElementSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// NextElementSibling returns the value of property "CharacterData.nextElementSibling".
//
// The returned bool will be false if there is no such property.
func (this CharacterData) NextElementSibling() (Element, bool) {
	var _ok bool
	_ret := bindings.GetCharacterDataNextElementSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// SubstringData calls the method "CharacterData.substringData".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) SubstringData(offset uint32, count uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataSubstringData(
		this.Ref(), js.Pointer(&_ok),
		uint32(offset),
		uint32(count),
	)

	return js.String{}.FromRef(_ret), _ok
}

// SubstringDataFunc returns the method "CharacterData.substringData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) SubstringDataFunc() (fn js.Func[func(offset uint32, count uint32) js.String]) {
	return fn.FromRef(
		bindings.CharacterDataSubstringDataFunc(
			this.Ref(),
		),
	)
}

// AppendData calls the method "CharacterData.appendData".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) AppendData(data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataAppendData(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendDataFunc returns the method "CharacterData.appendData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) AppendDataFunc() (fn js.Func[func(data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataAppendDataFunc(
			this.Ref(),
		),
	)
}

// InsertData calls the method "CharacterData.insertData".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) InsertData(offset uint32, data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataInsertData(
		this.Ref(), js.Pointer(&_ok),
		uint32(offset),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDataFunc returns the method "CharacterData.insertData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) InsertDataFunc() (fn js.Func[func(offset uint32, data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataInsertDataFunc(
			this.Ref(),
		),
	)
}

// DeleteData calls the method "CharacterData.deleteData".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) DeleteData(offset uint32, count uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataDeleteData(
		this.Ref(), js.Pointer(&_ok),
		uint32(offset),
		uint32(count),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteDataFunc returns the method "CharacterData.deleteData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) DeleteDataFunc() (fn js.Func[func(offset uint32, count uint32)]) {
	return fn.FromRef(
		bindings.CharacterDataDeleteDataFunc(
			this.Ref(),
		),
	)
}

// ReplaceData calls the method "CharacterData.replaceData".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) ReplaceData(offset uint32, count uint32, data js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataReplaceData(
		this.Ref(), js.Pointer(&_ok),
		uint32(offset),
		uint32(count),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceDataFunc returns the method "CharacterData.replaceData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) ReplaceDataFunc() (fn js.Func[func(offset uint32, count uint32, data js.String)]) {
	return fn.FromRef(
		bindings.CharacterDataReplaceDataFunc(
			this.Ref(),
		),
	)
}

// Before calls the method "CharacterData.before".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) Before(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataBefore(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeforeFunc returns the method "CharacterData.before".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) BeforeFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataBeforeFunc(
			this.Ref(),
		),
	)
}

// After calls the method "CharacterData.after".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) After(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataAfter(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AfterFunc returns the method "CharacterData.after".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) AfterFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataAfterFunc(
			this.Ref(),
		),
	)
}

// ReplaceWith calls the method "CharacterData.replaceWith".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) ReplaceWith(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataReplaceWith(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceWithFunc returns the method "CharacterData.replaceWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) ReplaceWithFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.CharacterDataReplaceWithFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "CharacterData.remove".
//
// The returned bool will be false if there is no such method.
func (this CharacterData) Remove() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCharacterDataRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "CharacterData.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CharacterData) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CharacterDataRemoveFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this Client) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetClientUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FrameType returns the value of property "Client.frameType".
//
// The returned bool will be false if there is no such property.
func (this Client) FrameType() (FrameType, bool) {
	var _ok bool
	_ret := bindings.GetClientFrameType(
		this.Ref(), js.Pointer(&_ok),
	)
	return FrameType(_ret), _ok
}

// Id returns the value of property "Client.id".
//
// The returned bool will be false if there is no such property.
func (this Client) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetClientId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type returns the value of property "Client.type".
//
// The returned bool will be false if there is no such property.
func (this Client) Type() (ClientType, bool) {
	var _ok bool
	_ret := bindings.GetClientType(
		this.Ref(), js.Pointer(&_ok),
	)
	return ClientType(_ret), _ok
}

// LifecycleState returns the value of property "Client.lifecycleState".
//
// The returned bool will be false if there is no such property.
func (this Client) LifecycleState() (ClientLifecycleState, bool) {
	var _ok bool
	_ret := bindings.GetClientLifecycleState(
		this.Ref(), js.Pointer(&_ok),
	)
	return ClientLifecycleState(_ret), _ok
}

// PostMessage calls the method "Client.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Client) PostMessage(message js.Any, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallClientPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "Client.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Client) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.ClientPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "Client.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Client) PostMessage1(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallClientPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "Client.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Client) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.ClientPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "Client.postMessage".
//
// The returned bool will be false if there is no such method.
func (this Client) PostMessage2(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallClientPostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "Client.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Client) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.ClientPostMessage2Func(
			this.Ref(),
		),
	)
}

type ClientQueryOptions struct {
	// IncludeUncontrolled is "ClientQueryOptions.includeUncontrolled"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 ClientQueryOptions ClientQueryOptions [// ClientQueryOptions] [0x140005bfae0 0x140005bfc20 0x140005bfb80] 0x14000574588 {0 0}} in the application heap.
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
