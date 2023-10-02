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

type SVGFETileElement struct {
	SVGElement
}

func (this SVGFETileElement) Once() SVGFETileElement {
	this.Ref().Once()
	return this
}

func (this SVGFETileElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFETileElement) FromRef(ref js.Ref) SVGFETileElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFETileElement) Free() {
	this.Ref().Free()
}

// In1 returns the value of property "SVGFETileElement.in1".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) In1() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementIn1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFETileElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFETileElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFETileElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFETileElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFETileElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFETileElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETileElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

const (
	SVGFETurbulenceElement_SVG_TURBULENCE_TYPE_UNKNOWN      uint16 = 0
	SVGFETurbulenceElement_SVG_TURBULENCE_TYPE_FRACTALNOISE uint16 = 1
	SVGFETurbulenceElement_SVG_TURBULENCE_TYPE_TURBULENCE   uint16 = 2
	SVGFETurbulenceElement_SVG_STITCHTYPE_UNKNOWN           uint16 = 0
	SVGFETurbulenceElement_SVG_STITCHTYPE_STITCH            uint16 = 1
	SVGFETurbulenceElement_SVG_STITCHTYPE_NOSTITCH          uint16 = 2
)

type SVGFETurbulenceElement struct {
	SVGElement
}

func (this SVGFETurbulenceElement) Once() SVGFETurbulenceElement {
	this.Ref().Once()
	return this
}

func (this SVGFETurbulenceElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFETurbulenceElement) FromRef(ref js.Ref) SVGFETurbulenceElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFETurbulenceElement) Free() {
	this.Ref().Free()
}

// BaseFrequencyX returns the value of property "SVGFETurbulenceElement.baseFrequencyX".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) BaseFrequencyX() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementBaseFrequencyX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// BaseFrequencyY returns the value of property "SVGFETurbulenceElement.baseFrequencyY".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) BaseFrequencyY() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementBaseFrequencyY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// NumOctaves returns the value of property "SVGFETurbulenceElement.numOctaves".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) NumOctaves() (SVGAnimatedInteger, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementNumOctaves(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedInteger{}.FromRef(_ret), _ok
}

// Seed returns the value of property "SVGFETurbulenceElement.seed".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Seed() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementSeed(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// StitchTiles returns the value of property "SVGFETurbulenceElement.stitchTiles".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) StitchTiles() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementStitchTiles(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Type returns the value of property "SVGFETurbulenceElement.type".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Type() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFETurbulenceElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFETurbulenceElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFETurbulenceElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFETurbulenceElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Result returns the value of property "SVGFETurbulenceElement.result".
//
// The returned bool will be false if there is no such property.
func (this SVGFETurbulenceElement) Result() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFETurbulenceElementResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type SVGFilterElement struct {
	SVGElement
}

func (this SVGFilterElement) Once() SVGFilterElement {
	this.Ref().Once()
	return this
}

func (this SVGFilterElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGFilterElement) FromRef(ref js.Ref) SVGFilterElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGFilterElement) Free() {
	this.Ref().Free()
}

// FilterUnits returns the value of property "SVGFilterElement.filterUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) FilterUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementFilterUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// PrimitiveUnits returns the value of property "SVGFilterElement.primitiveUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) PrimitiveUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementPrimitiveUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGFilterElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGFilterElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGFilterElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGFilterElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGFilterElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGFilterElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGFilterElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type SVGForeignObjectElement struct {
	SVGGraphicsElement
}

func (this SVGForeignObjectElement) Once() SVGForeignObjectElement {
	this.Ref().Once()
	return this
}

func (this SVGForeignObjectElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGForeignObjectElement) FromRef(ref js.Ref) SVGForeignObjectElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGForeignObjectElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGForeignObjectElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGForeignObjectElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGForeignObjectElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGForeignObjectElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGForeignObjectElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGForeignObjectElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGForeignObjectElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGForeignObjectElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGForeignObjectElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGForeignObjectElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGForeignObjectElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGForeignObjectElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGGElement struct {
	SVGGraphicsElement
}

func (this SVGGElement) Once() SVGGElement {
	this.Ref().Once()
	return this
}

func (this SVGGElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGGElement) FromRef(ref js.Ref) SVGGElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGGElement) Free() {
	this.Ref().Free()
}

type SVGGeometryElement struct {
	SVGGraphicsElement
}

func (this SVGGeometryElement) Once() SVGGeometryElement {
	this.Ref().Once()
	return this
}

func (this SVGGeometryElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGGeometryElement) FromRef(ref js.Ref) SVGGeometryElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGGeometryElement) Free() {
	this.Ref().Free()
}

// PathLength returns the value of property "SVGGeometryElement.pathLength".
//
// The returned bool will be false if there is no such property.
func (this SVGGeometryElement) PathLength() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGGeometryElementPathLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

// IsPointInFill calls the method "SVGGeometryElement.isPointInFill".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) IsPointInFill(point DOMPointInit) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementIsPointInFill(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
	)

	return _ret == js.True, _ok
}

// IsPointInFillFunc returns the method "SVGGeometryElement.isPointInFill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) IsPointInFillFunc() (fn js.Func[func(point DOMPointInit) bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInFillFunc(
			this.Ref(),
		),
	)
}

// IsPointInFill1 calls the method "SVGGeometryElement.isPointInFill".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) IsPointInFill1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementIsPointInFill1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsPointInFill1Func returns the method "SVGGeometryElement.isPointInFill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) IsPointInFill1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInFill1Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "SVGGeometryElement.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) IsPointInStroke(point DOMPointInit) (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementIsPointInStroke(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
	)

	return _ret == js.True, _ok
}

// IsPointInStrokeFunc returns the method "SVGGeometryElement.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) IsPointInStrokeFunc() (fn js.Func[func(point DOMPointInit) bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "SVGGeometryElement.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) IsPointInStroke1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementIsPointInStroke1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsPointInStroke1Func returns the method "SVGGeometryElement.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) IsPointInStroke1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// GetTotalLength calls the method "SVGGeometryElement.getTotalLength".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) GetTotalLength() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementGetTotalLength(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetTotalLengthFunc returns the method "SVGGeometryElement.getTotalLength".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) GetTotalLengthFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGGeometryElementGetTotalLengthFunc(
			this.Ref(),
		),
	)
}

// GetPointAtLength calls the method "SVGGeometryElement.getPointAtLength".
//
// The returned bool will be false if there is no such method.
func (this SVGGeometryElement) GetPointAtLength(distance float32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGGeometryElementGetPointAtLength(
		this.Ref(), js.Pointer(&_ok),
		float32(distance),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// GetPointAtLengthFunc returns the method "SVGGeometryElement.getPointAtLength".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGeometryElement) GetPointAtLengthFunc() (fn js.Func[func(distance float32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGGeometryElementGetPointAtLengthFunc(
			this.Ref(),
		),
	)
}

const (
	SVGGradientElement_SVG_SPREADMETHOD_UNKNOWN uint16 = 0
	SVGGradientElement_SVG_SPREADMETHOD_PAD     uint16 = 1
	SVGGradientElement_SVG_SPREADMETHOD_REFLECT uint16 = 2
	SVGGradientElement_SVG_SPREADMETHOD_REPEAT  uint16 = 3
)

type SVGGradientElement struct {
	SVGElement
}

func (this SVGGradientElement) Once() SVGGradientElement {
	this.Ref().Once()
	return this
}

func (this SVGGradientElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGGradientElement) FromRef(ref js.Ref) SVGGradientElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGGradientElement) Free() {
	this.Ref().Free()
}

// GradientUnits returns the value of property "SVGGradientElement.gradientUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGGradientElement) GradientUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGGradientElementGradientUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// GradientTransform returns the value of property "SVGGradientElement.gradientTransform".
//
// The returned bool will be false if there is no such property.
func (this SVGGradientElement) GradientTransform() (SVGAnimatedTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGGradientElementGradientTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedTransformList{}.FromRef(_ret), _ok
}

// SpreadMethod returns the value of property "SVGGradientElement.spreadMethod".
//
// The returned bool will be false if there is no such property.
func (this SVGGradientElement) SpreadMethod() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGGradientElementSpreadMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGGradientElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGGradientElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGGradientElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type SVGGraphicsElement struct {
	SVGElement
}

func (this SVGGraphicsElement) Once() SVGGraphicsElement {
	this.Ref().Once()
	return this
}

func (this SVGGraphicsElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGGraphicsElement) FromRef(ref js.Ref) SVGGraphicsElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGGraphicsElement) Free() {
	this.Ref().Free()
}

// Transform returns the value of property "SVGGraphicsElement.transform".
//
// The returned bool will be false if there is no such property.
func (this SVGGraphicsElement) Transform() (SVGAnimatedTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGGraphicsElementTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedTransformList{}.FromRef(_ret), _ok
}

// RequiredExtensions returns the value of property "SVGGraphicsElement.requiredExtensions".
//
// The returned bool will be false if there is no such property.
func (this SVGGraphicsElement) RequiredExtensions() (SVGStringList, bool) {
	var _ok bool
	_ret := bindings.GetSVGGraphicsElementRequiredExtensions(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGStringList{}.FromRef(_ret), _ok
}

// SystemLanguage returns the value of property "SVGGraphicsElement.systemLanguage".
//
// The returned bool will be false if there is no such property.
func (this SVGGraphicsElement) SystemLanguage() (SVGStringList, bool) {
	var _ok bool
	_ret := bindings.GetSVGGraphicsElementSystemLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGStringList{}.FromRef(_ret), _ok
}

// GetBBox calls the method "SVGGraphicsElement.getBBox".
//
// The returned bool will be false if there is no such method.
func (this SVGGraphicsElement) GetBBox(options SVGBoundingBoxOptions) (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallSVGGraphicsElementGetBBox(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetBBoxFunc returns the method "SVGGraphicsElement.getBBox".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGraphicsElement) GetBBoxFunc() (fn js.Func[func(options SVGBoundingBoxOptions) DOMRect]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetBBoxFunc(
			this.Ref(),
		),
	)
}

// GetBBox1 calls the method "SVGGraphicsElement.getBBox".
//
// The returned bool will be false if there is no such method.
func (this SVGGraphicsElement) GetBBox1() (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallSVGGraphicsElementGetBBox1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetBBox1Func returns the method "SVGGraphicsElement.getBBox".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGraphicsElement) GetBBox1Func() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetBBox1Func(
			this.Ref(),
		),
	)
}

// GetCTM calls the method "SVGGraphicsElement.getCTM".
//
// The returned bool will be false if there is no such method.
func (this SVGGraphicsElement) GetCTM() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallSVGGraphicsElementGetCTM(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// GetCTMFunc returns the method "SVGGraphicsElement.getCTM".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGraphicsElement) GetCTMFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetCTMFunc(
			this.Ref(),
		),
	)
}

// GetScreenCTM calls the method "SVGGraphicsElement.getScreenCTM".
//
// The returned bool will be false if there is no such method.
func (this SVGGraphicsElement) GetScreenCTM() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallSVGGraphicsElementGetScreenCTM(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// GetScreenCTMFunc returns the method "SVGGraphicsElement.getScreenCTM".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGGraphicsElement) GetScreenCTMFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetScreenCTMFunc(
			this.Ref(),
		),
	)
}

type SVGLineElement struct {
	SVGGeometryElement
}

func (this SVGLineElement) Once() SVGLineElement {
	this.Ref().Once()
	return this
}

func (this SVGLineElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGLineElement) FromRef(ref js.Ref) SVGLineElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGLineElement) Free() {
	this.Ref().Free()
}

// X1 returns the value of property "SVGLineElement.x1".
//
// The returned bool will be false if there is no such property.
func (this SVGLineElement) X1() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLineElementX1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y1 returns the value of property "SVGLineElement.y1".
//
// The returned bool will be false if there is no such property.
func (this SVGLineElement) Y1() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLineElementY1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// X2 returns the value of property "SVGLineElement.x2".
//
// The returned bool will be false if there is no such property.
func (this SVGLineElement) X2() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLineElementX2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y2 returns the value of property "SVGLineElement.y2".
//
// The returned bool will be false if there is no such property.
func (this SVGLineElement) Y2() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLineElementY2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGLinearGradientElement struct {
	SVGGradientElement
}

func (this SVGLinearGradientElement) Once() SVGLinearGradientElement {
	this.Ref().Once()
	return this
}

func (this SVGLinearGradientElement) Ref() js.Ref {
	return this.SVGGradientElement.Ref()
}

func (this SVGLinearGradientElement) FromRef(ref js.Ref) SVGLinearGradientElement {
	this.SVGGradientElement = this.SVGGradientElement.FromRef(ref)
	return this
}

func (this SVGLinearGradientElement) Free() {
	this.Ref().Free()
}

// X1 returns the value of property "SVGLinearGradientElement.x1".
//
// The returned bool will be false if there is no such property.
func (this SVGLinearGradientElement) X1() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLinearGradientElementX1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y1 returns the value of property "SVGLinearGradientElement.y1".
//
// The returned bool will be false if there is no such property.
func (this SVGLinearGradientElement) Y1() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLinearGradientElementY1(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// X2 returns the value of property "SVGLinearGradientElement.x2".
//
// The returned bool will be false if there is no such property.
func (this SVGLinearGradientElement) X2() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLinearGradientElementX2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y2 returns the value of property "SVGLinearGradientElement.y2".
//
// The returned bool will be false if there is no such property.
func (this SVGLinearGradientElement) Y2() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGLinearGradientElementY2(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGMPathElement struct {
	SVGElement
}

func (this SVGMPathElement) Once() SVGMPathElement {
	this.Ref().Once()
	return this
}

func (this SVGMPathElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGMPathElement) FromRef(ref js.Ref) SVGMPathElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGMPathElement) Free() {
	this.Ref().Free()
}

// Href returns the value of property "SVGMPathElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGMPathElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGMPathElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

const (
	SVGMarkerElement_SVG_MARKERUNITS_UNKNOWN        uint16 = 0
	SVGMarkerElement_SVG_MARKERUNITS_USERSPACEONUSE uint16 = 1
	SVGMarkerElement_SVG_MARKERUNITS_STROKEWIDTH    uint16 = 2
	SVGMarkerElement_SVG_MARKER_ORIENT_UNKNOWN      uint16 = 0
	SVGMarkerElement_SVG_MARKER_ORIENT_AUTO         uint16 = 1
	SVGMarkerElement_SVG_MARKER_ORIENT_ANGLE        uint16 = 2
)

type SVGMarkerElement struct {
	SVGElement
}

func (this SVGMarkerElement) Once() SVGMarkerElement {
	this.Ref().Once()
	return this
}

func (this SVGMarkerElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGMarkerElement) FromRef(ref js.Ref) SVGMarkerElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGMarkerElement) Free() {
	this.Ref().Free()
}

// RefX returns the value of property "SVGMarkerElement.refX".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) RefX() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementRefX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// RefY returns the value of property "SVGMarkerElement.refY".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) RefY() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementRefY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// MarkerUnits returns the value of property "SVGMarkerElement.markerUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) MarkerUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementMarkerUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// MarkerWidth returns the value of property "SVGMarkerElement.markerWidth".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) MarkerWidth() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementMarkerWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// MarkerHeight returns the value of property "SVGMarkerElement.markerHeight".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) MarkerHeight() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementMarkerHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// OrientType returns the value of property "SVGMarkerElement.orientType".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) OrientType() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementOrientType(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// OrientAngle returns the value of property "SVGMarkerElement.orientAngle".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) OrientAngle() (SVGAnimatedAngle, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementOrientAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedAngle{}.FromRef(_ret), _ok
}

// Orient returns the value of property "SVGMarkerElement.orient".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) Orient() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementOrient(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Orient sets the value of property "SVGMarkerElement.orient" to val.
//
// It returns false if the property cannot be set.
func (this SVGMarkerElement) SetOrient(val js.String) bool {
	return js.True == bindings.SetSVGMarkerElementOrient(
		this.Ref(),
		val.Ref(),
	)
}

// ViewBox returns the value of property "SVGMarkerElement.viewBox".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) ViewBox() (SVGAnimatedRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementViewBox(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedRect{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGMarkerElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGMarkerElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGMarkerElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

// SetOrientToAuto calls the method "SVGMarkerElement.setOrientToAuto".
//
// The returned bool will be false if there is no such method.
func (this SVGMarkerElement) SetOrientToAuto() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGMarkerElementSetOrientToAuto(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetOrientToAutoFunc returns the method "SVGMarkerElement.setOrientToAuto".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGMarkerElement) SetOrientToAutoFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGMarkerElementSetOrientToAutoFunc(
			this.Ref(),
		),
	)
}

// SetOrientToAngle calls the method "SVGMarkerElement.setOrientToAngle".
//
// The returned bool will be false if there is no such method.
func (this SVGMarkerElement) SetOrientToAngle(angle SVGAngle) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGMarkerElementSetOrientToAngle(
		this.Ref(), js.Pointer(&_ok),
		angle.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetOrientToAngleFunc returns the method "SVGMarkerElement.setOrientToAngle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGMarkerElement) SetOrientToAngleFunc() (fn js.Func[func(angle SVGAngle)]) {
	return fn.FromRef(
		bindings.SVGMarkerElementSetOrientToAngleFunc(
			this.Ref(),
		),
	)
}

type SVGMaskElement struct {
	SVGElement
}

func (this SVGMaskElement) Once() SVGMaskElement {
	this.Ref().Once()
	return this
}

func (this SVGMaskElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGMaskElement) FromRef(ref js.Ref) SVGMaskElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGMaskElement) Free() {
	this.Ref().Free()
}

// MaskUnits returns the value of property "SVGMaskElement.maskUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) MaskUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementMaskUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// MaskContentUnits returns the value of property "SVGMaskElement.maskContentUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) MaskContentUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementMaskContentUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGMaskElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGMaskElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGMaskElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGMaskElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGMaskElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGMaskElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGMetadataElement struct {
	SVGElement
}

func (this SVGMetadataElement) Once() SVGMetadataElement {
	this.Ref().Once()
	return this
}

func (this SVGMetadataElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGMetadataElement) FromRef(ref js.Ref) SVGMetadataElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGMetadataElement) Free() {
	this.Ref().Free()
}

type SVGPathElement struct {
	SVGGeometryElement
}

func (this SVGPathElement) Once() SVGPathElement {
	this.Ref().Once()
	return this
}

func (this SVGPathElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGPathElement) FromRef(ref js.Ref) SVGPathElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGPathElement) Free() {
	this.Ref().Free()
}

type SVGPatternElement struct {
	SVGElement
}

func (this SVGPatternElement) Once() SVGPatternElement {
	this.Ref().Once()
	return this
}

func (this SVGPatternElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGPatternElement) FromRef(ref js.Ref) SVGPatternElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGPatternElement) Free() {
	this.Ref().Free()
}

// PatternUnits returns the value of property "SVGPatternElement.patternUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) PatternUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementPatternUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// PatternContentUnits returns the value of property "SVGPatternElement.patternContentUnits".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) PatternContentUnits() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementPatternContentUnits(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// PatternTransform returns the value of property "SVGPatternElement.patternTransform".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) PatternTransform() (SVGAnimatedTransformList, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementPatternTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedTransformList{}.FromRef(_ret), _ok
}

// X returns the value of property "SVGPatternElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGPatternElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGPatternElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGPatternElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// ViewBox returns the value of property "SVGPatternElement.viewBox".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) ViewBox() (SVGAnimatedRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementViewBox(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedRect{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGPatternElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGPatternElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGPatternElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGPatternElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type SVGPointList struct {
	ref js.Ref
}

func (this SVGPointList) Once() SVGPointList {
	this.Ref().Once()
	return this
}

func (this SVGPointList) Ref() js.Ref {
	return this.ref
}

func (this SVGPointList) FromRef(ref js.Ref) SVGPointList {
	this.ref = ref
	return this
}

func (this SVGPointList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SVGPointList.length".
//
// The returned bool will be false if there is no such property.
func (this SVGPointList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGPointListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfItems returns the value of property "SVGPointList.numberOfItems".
//
// The returned bool will be false if there is no such property.
func (this SVGPointList) NumberOfItems() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSVGPointListNumberOfItems(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Clear calls the method "SVGPointList.clear".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "SVGPointList.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGPointListClearFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGPointList.initialize".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) Initialize(newItem DOMPoint) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListInitialize(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// InitializeFunc returns the method "SVGPointList.initialize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) InitializeFunc() (fn js.Func[func(newItem DOMPoint) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListInitializeFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGPointList.getItem".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) GetItem(index uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListGetItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// GetItemFunc returns the method "SVGPointList.getItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) GetItemFunc() (fn js.Func[func(index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListGetItemFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGPointList.insertItemBefore".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) InsertItemBefore(newItem DOMPoint, index uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListInsertItemBefore(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// InsertItemBeforeFunc returns the method "SVGPointList.insertItemBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) InsertItemBeforeFunc() (fn js.Func[func(newItem DOMPoint, index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGPointList.replaceItem".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) ReplaceItem(newItem DOMPoint, index uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListReplaceItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
		uint32(index),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ReplaceItemFunc returns the method "SVGPointList.replaceItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) ReplaceItemFunc() (fn js.Func[func(newItem DOMPoint, index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGPointList.removeItem".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) RemoveItem(index uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListRemoveItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// RemoveItemFunc returns the method "SVGPointList.removeItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) RemoveItemFunc() (fn js.Func[func(index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGPointList.appendItem".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) AppendItem(newItem DOMPoint) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListAppendItem(
		this.Ref(), js.Pointer(&_ok),
		newItem.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// AppendItemFunc returns the method "SVGPointList.appendItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) AppendItemFunc() (fn js.Func[func(newItem DOMPoint) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListAppendItemFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGPointList.".
//
// The returned bool will be false if there is no such method.
func (this SVGPointList) Set(index uint32, newItem DOMPoint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGPointListSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		newItem.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "SVGPointList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGPointList) SetFunc() (fn js.Func[func(index uint32, newItem DOMPoint)]) {
	return fn.FromRef(
		bindings.SVGPointListSetFunc(
			this.Ref(),
		),
	)
}

type SVGPolygonElement struct {
	SVGGeometryElement
}

func (this SVGPolygonElement) Once() SVGPolygonElement {
	this.Ref().Once()
	return this
}

func (this SVGPolygonElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGPolygonElement) FromRef(ref js.Ref) SVGPolygonElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGPolygonElement) Free() {
	this.Ref().Free()
}

// Points returns the value of property "SVGPolygonElement.points".
//
// The returned bool will be false if there is no such property.
func (this SVGPolygonElement) Points() (SVGPointList, bool) {
	var _ok bool
	_ret := bindings.GetSVGPolygonElementPoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPointList{}.FromRef(_ret), _ok
}

// AnimatedPoints returns the value of property "SVGPolygonElement.animatedPoints".
//
// The returned bool will be false if there is no such property.
func (this SVGPolygonElement) AnimatedPoints() (SVGPointList, bool) {
	var _ok bool
	_ret := bindings.GetSVGPolygonElementAnimatedPoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPointList{}.FromRef(_ret), _ok
}

type SVGPolylineElement struct {
	SVGGeometryElement
}

func (this SVGPolylineElement) Once() SVGPolylineElement {
	this.Ref().Once()
	return this
}

func (this SVGPolylineElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGPolylineElement) FromRef(ref js.Ref) SVGPolylineElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGPolylineElement) Free() {
	this.Ref().Free()
}

// Points returns the value of property "SVGPolylineElement.points".
//
// The returned bool will be false if there is no such property.
func (this SVGPolylineElement) Points() (SVGPointList, bool) {
	var _ok bool
	_ret := bindings.GetSVGPolylineElementPoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPointList{}.FromRef(_ret), _ok
}

// AnimatedPoints returns the value of property "SVGPolylineElement.animatedPoints".
//
// The returned bool will be false if there is no such property.
func (this SVGPolylineElement) AnimatedPoints() (SVGPointList, bool) {
	var _ok bool
	_ret := bindings.GetSVGPolylineElementAnimatedPoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGPointList{}.FromRef(_ret), _ok
}

type SVGRadialGradientElement struct {
	SVGGradientElement
}

func (this SVGRadialGradientElement) Once() SVGRadialGradientElement {
	this.Ref().Once()
	return this
}

func (this SVGRadialGradientElement) Ref() js.Ref {
	return this.SVGGradientElement.Ref()
}

func (this SVGRadialGradientElement) FromRef(ref js.Ref) SVGRadialGradientElement {
	this.SVGGradientElement = this.SVGGradientElement.FromRef(ref)
	return this
}

func (this SVGRadialGradientElement) Free() {
	this.Ref().Free()
}

// Cx returns the value of property "SVGRadialGradientElement.cx".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) Cx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementCx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Cy returns the value of property "SVGRadialGradientElement.cy".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) Cy() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementCy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// R returns the value of property "SVGRadialGradientElement.r".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) R() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementR(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Fx returns the value of property "SVGRadialGradientElement.fx".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) Fx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementFx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Fy returns the value of property "SVGRadialGradientElement.fy".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) Fy() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementFy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Fr returns the value of property "SVGRadialGradientElement.fr".
//
// The returned bool will be false if there is no such property.
func (this SVGRadialGradientElement) Fr() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRadialGradientElementFr(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGRectElement struct {
	SVGGeometryElement
}

func (this SVGRectElement) Once() SVGRectElement {
	this.Ref().Once()
	return this
}

func (this SVGRectElement) Ref() js.Ref {
	return this.SVGGeometryElement.Ref()
}

func (this SVGRectElement) FromRef(ref js.Ref) SVGRectElement {
	this.SVGGeometryElement = this.SVGGeometryElement.FromRef(ref)
	return this
}

func (this SVGRectElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGRectElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) X() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGRectElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) Y() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Width returns the value of property "SVGRectElement.width".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) Width() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Height returns the value of property "SVGRectElement.height".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) Height() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Rx returns the value of property "SVGRectElement.rx".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) Rx() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementRx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Ry returns the value of property "SVGRectElement.ry".
//
// The returned bool will be false if there is no such property.
func (this SVGRectElement) Ry() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGRectElementRy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

type SVGSetElement struct {
	SVGAnimationElement
}

func (this SVGSetElement) Once() SVGSetElement {
	this.Ref().Once()
	return this
}

func (this SVGSetElement) Ref() js.Ref {
	return this.SVGAnimationElement.Ref()
}

func (this SVGSetElement) FromRef(ref js.Ref) SVGSetElement {
	this.SVGAnimationElement = this.SVGAnimationElement.FromRef(ref)
	return this
}

func (this SVGSetElement) Free() {
	this.Ref().Free()
}

type SVGStopElement struct {
	SVGElement
}

func (this SVGStopElement) Once() SVGStopElement {
	this.Ref().Once()
	return this
}

func (this SVGStopElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGStopElement) FromRef(ref js.Ref) SVGStopElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGStopElement) Free() {
	this.Ref().Free()
}

// Offset returns the value of property "SVGStopElement.offset".
//
// The returned bool will be false if there is no such property.
func (this SVGStopElement) Offset() (SVGAnimatedNumber, bool) {
	var _ok bool
	_ret := bindings.GetSVGStopElementOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumber{}.FromRef(_ret), _ok
}

type SVGStyleElement struct {
	SVGElement
}

func (this SVGStyleElement) Once() SVGStyleElement {
	this.Ref().Once()
	return this
}

func (this SVGStyleElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGStyleElement) FromRef(ref js.Ref) SVGStyleElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGStyleElement) Free() {
	this.Ref().Free()
}

// Type returns the value of property "SVGStyleElement.type".
//
// The returned bool will be false if there is no such property.
func (this SVGStyleElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGStyleElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "SVGStyleElement.type" to val.
//
// It returns false if the property cannot be set.
func (this SVGStyleElement) SetType(val js.String) bool {
	return js.True == bindings.SetSVGStyleElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Media returns the value of property "SVGStyleElement.media".
//
// The returned bool will be false if there is no such property.
func (this SVGStyleElement) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGStyleElementMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media sets the value of property "SVGStyleElement.media" to val.
//
// It returns false if the property cannot be set.
func (this SVGStyleElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetSVGStyleElementMedia(
		this.Ref(),
		val.Ref(),
	)
}

// Title returns the value of property "SVGStyleElement.title".
//
// The returned bool will be false if there is no such property.
func (this SVGStyleElement) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSVGStyleElementTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Title sets the value of property "SVGStyleElement.title" to val.
//
// It returns false if the property cannot be set.
func (this SVGStyleElement) SetTitle(val js.String) bool {
	return js.True == bindings.SetSVGStyleElementTitle(
		this.Ref(),
		val.Ref(),
	)
}

// Sheet returns the value of property "SVGStyleElement.sheet".
//
// The returned bool will be false if there is no such property.
func (this SVGStyleElement) Sheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetSVGStyleElementSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

type SVGSwitchElement struct {
	SVGGraphicsElement
}

func (this SVGSwitchElement) Once() SVGSwitchElement {
	this.Ref().Once()
	return this
}

func (this SVGSwitchElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGSwitchElement) FromRef(ref js.Ref) SVGSwitchElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGSwitchElement) Free() {
	this.Ref().Free()
}

type SVGSymbolElement struct {
	SVGGraphicsElement
}

func (this SVGSymbolElement) Once() SVGSymbolElement {
	this.Ref().Once()
	return this
}

func (this SVGSymbolElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGSymbolElement) FromRef(ref js.Ref) SVGSymbolElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGSymbolElement) Free() {
	this.Ref().Free()
}

// ViewBox returns the value of property "SVGSymbolElement.viewBox".
//
// The returned bool will be false if there is no such property.
func (this SVGSymbolElement) ViewBox() (SVGAnimatedRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGSymbolElementViewBox(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedRect{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGSymbolElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGSymbolElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGSymbolElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

type SVGTSpanElement struct {
	SVGTextPositioningElement
}

func (this SVGTSpanElement) Once() SVGTSpanElement {
	this.Ref().Once()
	return this
}

func (this SVGTSpanElement) Ref() js.Ref {
	return this.SVGTextPositioningElement.Ref()
}

func (this SVGTSpanElement) FromRef(ref js.Ref) SVGTSpanElement {
	this.SVGTextPositioningElement = this.SVGTextPositioningElement.FromRef(ref)
	return this
}

func (this SVGTSpanElement) Free() {
	this.Ref().Free()
}

const (
	SVGTextContentElement_LENGTHADJUST_UNKNOWN          uint16 = 0
	SVGTextContentElement_LENGTHADJUST_SPACING          uint16 = 1
	SVGTextContentElement_LENGTHADJUST_SPACINGANDGLYPHS uint16 = 2
)

type SVGTextContentElement struct {
	SVGGraphicsElement
}

func (this SVGTextContentElement) Once() SVGTextContentElement {
	this.Ref().Once()
	return this
}

func (this SVGTextContentElement) Ref() js.Ref {
	return this.SVGGraphicsElement.Ref()
}

func (this SVGTextContentElement) FromRef(ref js.Ref) SVGTextContentElement {
	this.SVGGraphicsElement = this.SVGGraphicsElement.FromRef(ref)
	return this
}

func (this SVGTextContentElement) Free() {
	this.Ref().Free()
}

// TextLength returns the value of property "SVGTextContentElement.textLength".
//
// The returned bool will be false if there is no such property.
func (this SVGTextContentElement) TextLength() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextContentElementTextLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// LengthAdjust returns the value of property "SVGTextContentElement.lengthAdjust".
//
// The returned bool will be false if there is no such property.
func (this SVGTextContentElement) LengthAdjust() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextContentElementLengthAdjust(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// GetNumberOfChars calls the method "SVGTextContentElement.getNumberOfChars".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetNumberOfChars() (int32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetNumberOfChars(
		this.Ref(), js.Pointer(&_ok),
	)

	return int32(_ret), _ok
}

// GetNumberOfCharsFunc returns the method "SVGTextContentElement.getNumberOfChars".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetNumberOfCharsFunc() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetNumberOfCharsFunc(
			this.Ref(),
		),
	)
}

// GetComputedTextLength calls the method "SVGTextContentElement.getComputedTextLength".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetComputedTextLength() (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetComputedTextLength(
		this.Ref(), js.Pointer(&_ok),
	)

	return float32(_ret), _ok
}

// GetComputedTextLengthFunc returns the method "SVGTextContentElement.getComputedTextLength".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetComputedTextLengthFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetComputedTextLengthFunc(
			this.Ref(),
		),
	)
}

// GetSubStringLength calls the method "SVGTextContentElement.getSubStringLength".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetSubStringLength(charnum uint32, nchars uint32) (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetSubStringLength(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
		uint32(nchars),
	)

	return float32(_ret), _ok
}

// GetSubStringLengthFunc returns the method "SVGTextContentElement.getSubStringLength".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetSubStringLengthFunc() (fn js.Func[func(charnum uint32, nchars uint32) float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetSubStringLengthFunc(
			this.Ref(),
		),
	)
}

// GetStartPositionOfChar calls the method "SVGTextContentElement.getStartPositionOfChar".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetStartPositionOfChar(charnum uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetStartPositionOfChar(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// GetStartPositionOfCharFunc returns the method "SVGTextContentElement.getStartPositionOfChar".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetStartPositionOfCharFunc() (fn js.Func[func(charnum uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetStartPositionOfCharFunc(
			this.Ref(),
		),
	)
}

// GetEndPositionOfChar calls the method "SVGTextContentElement.getEndPositionOfChar".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetEndPositionOfChar(charnum uint32) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetEndPositionOfChar(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// GetEndPositionOfCharFunc returns the method "SVGTextContentElement.getEndPositionOfChar".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetEndPositionOfCharFunc() (fn js.Func[func(charnum uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetEndPositionOfCharFunc(
			this.Ref(),
		),
	)
}

// GetExtentOfChar calls the method "SVGTextContentElement.getExtentOfChar".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetExtentOfChar(charnum uint32) (DOMRect, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetExtentOfChar(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
	)

	return DOMRect{}.FromRef(_ret), _ok
}

// GetExtentOfCharFunc returns the method "SVGTextContentElement.getExtentOfChar".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetExtentOfCharFunc() (fn js.Func[func(charnum uint32) DOMRect]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetExtentOfCharFunc(
			this.Ref(),
		),
	)
}

// GetRotationOfChar calls the method "SVGTextContentElement.getRotationOfChar".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetRotationOfChar(charnum uint32) (float32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetRotationOfChar(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
	)

	return float32(_ret), _ok
}

// GetRotationOfCharFunc returns the method "SVGTextContentElement.getRotationOfChar".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetRotationOfCharFunc() (fn js.Func[func(charnum uint32) float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetRotationOfCharFunc(
			this.Ref(),
		),
	)
}

// GetCharNumAtPosition calls the method "SVGTextContentElement.getCharNumAtPosition".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetCharNumAtPosition(point DOMPointInit) (int32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetCharNumAtPosition(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
	)

	return int32(_ret), _ok
}

// GetCharNumAtPositionFunc returns the method "SVGTextContentElement.getCharNumAtPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetCharNumAtPositionFunc() (fn js.Func[func(point DOMPointInit) int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetCharNumAtPositionFunc(
			this.Ref(),
		),
	)
}

// GetCharNumAtPosition1 calls the method "SVGTextContentElement.getCharNumAtPosition".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) GetCharNumAtPosition1() (int32, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementGetCharNumAtPosition1(
		this.Ref(), js.Pointer(&_ok),
	)

	return int32(_ret), _ok
}

// GetCharNumAtPosition1Func returns the method "SVGTextContentElement.getCharNumAtPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) GetCharNumAtPosition1Func() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetCharNumAtPosition1Func(
			this.Ref(),
		),
	)
}

// SelectSubString calls the method "SVGTextContentElement.selectSubString".
//
// The returned bool will be false if there is no such method.
func (this SVGTextContentElement) SelectSubString(charnum uint32, nchars uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSVGTextContentElementSelectSubString(
		this.Ref(), js.Pointer(&_ok),
		uint32(charnum),
		uint32(nchars),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectSubStringFunc returns the method "SVGTextContentElement.selectSubString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SVGTextContentElement) SelectSubStringFunc() (fn js.Func[func(charnum uint32, nchars uint32)]) {
	return fn.FromRef(
		bindings.SVGTextContentElementSelectSubStringFunc(
			this.Ref(),
		),
	)
}

type SVGTextElement struct {
	SVGTextPositioningElement
}

func (this SVGTextElement) Once() SVGTextElement {
	this.Ref().Once()
	return this
}

func (this SVGTextElement) Ref() js.Ref {
	return this.SVGTextPositioningElement.Ref()
}

func (this SVGTextElement) FromRef(ref js.Ref) SVGTextElement {
	this.SVGTextPositioningElement = this.SVGTextPositioningElement.FromRef(ref)
	return this
}

func (this SVGTextElement) Free() {
	this.Ref().Free()
}

const (
	SVGTextPathElement_TEXTPATH_METHODTYPE_UNKNOWN  uint16 = 0
	SVGTextPathElement_TEXTPATH_METHODTYPE_ALIGN    uint16 = 1
	SVGTextPathElement_TEXTPATH_METHODTYPE_STRETCH  uint16 = 2
	SVGTextPathElement_TEXTPATH_SPACINGTYPE_UNKNOWN uint16 = 0
	SVGTextPathElement_TEXTPATH_SPACINGTYPE_AUTO    uint16 = 1
	SVGTextPathElement_TEXTPATH_SPACINGTYPE_EXACT   uint16 = 2
)

type SVGTextPathElement struct {
	SVGTextContentElement
}

func (this SVGTextPathElement) Once() SVGTextPathElement {
	this.Ref().Once()
	return this
}

func (this SVGTextPathElement) Ref() js.Ref {
	return this.SVGTextContentElement.Ref()
}

func (this SVGTextPathElement) FromRef(ref js.Ref) SVGTextPathElement {
	this.SVGTextContentElement = this.SVGTextContentElement.FromRef(ref)
	return this
}

func (this SVGTextPathElement) Free() {
	this.Ref().Free()
}

// StartOffset returns the value of property "SVGTextPathElement.startOffset".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPathElement) StartOffset() (SVGAnimatedLength, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPathElementStartOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLength{}.FromRef(_ret), _ok
}

// Method returns the value of property "SVGTextPathElement.method".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPathElement) Method() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPathElementMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Spacing returns the value of property "SVGTextPathElement.spacing".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPathElement) Spacing() (SVGAnimatedEnumeration, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPathElementSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedEnumeration{}.FromRef(_ret), _ok
}

// Href returns the value of property "SVGTextPathElement.href".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPathElement) Href() (SVGAnimatedString, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPathElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedString{}.FromRef(_ret), _ok
}

type SVGTextPositioningElement struct {
	SVGTextContentElement
}

func (this SVGTextPositioningElement) Once() SVGTextPositioningElement {
	this.Ref().Once()
	return this
}

func (this SVGTextPositioningElement) Ref() js.Ref {
	return this.SVGTextContentElement.Ref()
}

func (this SVGTextPositioningElement) FromRef(ref js.Ref) SVGTextPositioningElement {
	this.SVGTextContentElement = this.SVGTextContentElement.FromRef(ref)
	return this
}

func (this SVGTextPositioningElement) Free() {
	this.Ref().Free()
}

// X returns the value of property "SVGTextPositioningElement.x".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPositioningElement) X() (SVGAnimatedLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPositioningElementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLengthList{}.FromRef(_ret), _ok
}

// Y returns the value of property "SVGTextPositioningElement.y".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPositioningElement) Y() (SVGAnimatedLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPositioningElementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLengthList{}.FromRef(_ret), _ok
}

// Dx returns the value of property "SVGTextPositioningElement.dx".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPositioningElement) Dx() (SVGAnimatedLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPositioningElementDx(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLengthList{}.FromRef(_ret), _ok
}

// Dy returns the value of property "SVGTextPositioningElement.dy".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPositioningElement) Dy() (SVGAnimatedLengthList, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPositioningElementDy(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedLengthList{}.FromRef(_ret), _ok
}

// Rotate returns the value of property "SVGTextPositioningElement.rotate".
//
// The returned bool will be false if there is no such property.
func (this SVGTextPositioningElement) Rotate() (SVGAnimatedNumberList, bool) {
	var _ok bool
	_ret := bindings.GetSVGTextPositioningElementRotate(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedNumberList{}.FromRef(_ret), _ok
}

type SVGTitleElement struct {
	SVGElement
}

func (this SVGTitleElement) Once() SVGTitleElement {
	this.Ref().Once()
	return this
}

func (this SVGTitleElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGTitleElement) FromRef(ref js.Ref) SVGTitleElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGTitleElement) Free() {
	this.Ref().Free()
}

const (
	SVGUnitTypes_SVG_UNIT_TYPE_UNKNOWN           uint16 = 0
	SVGUnitTypes_SVG_UNIT_TYPE_USERSPACEONUSE    uint16 = 1
	SVGUnitTypes_SVG_UNIT_TYPE_OBJECTBOUNDINGBOX uint16 = 2
)

type SVGUnitTypes struct {
	ref js.Ref
}

func (this SVGUnitTypes) Once() SVGUnitTypes {
	this.Ref().Once()
	return this
}

func (this SVGUnitTypes) Ref() js.Ref {
	return this.ref
}

func (this SVGUnitTypes) FromRef(ref js.Ref) SVGUnitTypes {
	this.ref = ref
	return this
}

func (this SVGUnitTypes) Free() {
	this.Ref().Free()
}

type SVGUseElementShadowRoot struct {
	ShadowRoot
}

func (this SVGUseElementShadowRoot) Once() SVGUseElementShadowRoot {
	this.Ref().Once()
	return this
}

func (this SVGUseElementShadowRoot) Ref() js.Ref {
	return this.ShadowRoot.Ref()
}

func (this SVGUseElementShadowRoot) FromRef(ref js.Ref) SVGUseElementShadowRoot {
	this.ShadowRoot = this.ShadowRoot.FromRef(ref)
	return this
}

func (this SVGUseElementShadowRoot) Free() {
	this.Ref().Free()
}

type SVGViewElement struct {
	SVGElement
}

func (this SVGViewElement) Once() SVGViewElement {
	this.Ref().Once()
	return this
}

func (this SVGViewElement) Ref() js.Ref {
	return this.SVGElement.Ref()
}

func (this SVGViewElement) FromRef(ref js.Ref) SVGViewElement {
	this.SVGElement = this.SVGElement.FromRef(ref)
	return this
}

func (this SVGViewElement) Free() {
	this.Ref().Free()
}

// ViewBox returns the value of property "SVGViewElement.viewBox".
//
// The returned bool will be false if there is no such property.
func (this SVGViewElement) ViewBox() (SVGAnimatedRect, bool) {
	var _ok bool
	_ret := bindings.GetSVGViewElementViewBox(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedRect{}.FromRef(_ret), _ok
}

// PreserveAspectRatio returns the value of property "SVGViewElement.preserveAspectRatio".
//
// The returned bool will be false if there is no such property.
func (this SVGViewElement) PreserveAspectRatio() (SVGAnimatedPreserveAspectRatio, bool) {
	var _ok bool
	_ret := bindings.GetSVGViewElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGAnimatedPreserveAspectRatio{}.FromRef(_ret), _ok
}

type ScoreAdOutput struct {
	// Desirability is "ScoreAdOutput.desirability"
	//
	// Required
	Desirability float64
	// Bid is "ScoreAdOutput.bid"
	//
	// Optional
	Bid float64
	// BidCurrency is "ScoreAdOutput.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// IncomingBidInSellerCurrency is "ScoreAdOutput.incomingBidInSellerCurrency"
	//
	// Optional
	IncomingBidInSellerCurrency float64
	// AllowComponentAuction is "ScoreAdOutput.allowComponentAuction"
	//
	// Optional, defaults to false.
	AllowComponentAuction bool

	FFI_USE_Bid                         bool // for Bid.
	FFI_USE_IncomingBidInSellerCurrency bool // for IncomingBidInSellerCurrency.
	FFI_USE_AllowComponentAuction       bool // for AllowComponentAuction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScoreAdOutput with all fields set.
func (p ScoreAdOutput) FromRef(ref js.Ref) ScoreAdOutput {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ScoreAdOutput ScoreAdOutput [// ScoreAdOutput] [0x14000a430e0 0x14000a43180 0x14000a432c0 0x14000a43360 0x14000a434a0 0x14000a43220 0x14000a43400 0x14000a43540] 0x14000a02120 {0 0}} in the application heap.
func (p ScoreAdOutput) New() js.Ref {
	return bindings.ScoreAdOutputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScoreAdOutput) UpdateFrom(ref js.Ref) {
	bindings.ScoreAdOutputJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScoreAdOutput) Update(ref js.Ref) {
	bindings.ScoreAdOutputJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScoringBrowserSignals struct {
	// TopWindowHostname is "ScoringBrowserSignals.topWindowHostname"
	//
	// Required
	TopWindowHostname js.String
	// InterestGroupOwner is "ScoringBrowserSignals.interestGroupOwner"
	//
	// Required
	InterestGroupOwner js.String
	// RenderURL is "ScoringBrowserSignals.renderURL"
	//
	// Required
	RenderURL js.String
	// BiddingDurationMsec is "ScoringBrowserSignals.biddingDurationMsec"
	//
	// Required
	BiddingDurationMsec uint32
	// BidCurrency is "ScoringBrowserSignals.bidCurrency"
	//
	// Required
	BidCurrency js.String
	// DataVersion is "ScoringBrowserSignals.dataVersion"
	//
	// Optional
	DataVersion uint32
	// AdComponents is "ScoringBrowserSignals.adComponents"
	//
	// Optional
	AdComponents js.Array[js.String]

	FFI_USE_DataVersion bool // for DataVersion.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScoringBrowserSignals with all fields set.
func (p ScoringBrowserSignals) FromRef(ref js.Ref) ScoringBrowserSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ScoringBrowserSignals ScoringBrowserSignals [// ScoringBrowserSignals] [0x14000a435e0 0x14000a43680 0x14000a43720 0x14000a437c0 0x14000a43860 0x14000a43900 0x14000a43a40 0x14000a439a0] 0x14000a02168 {0 0}} in the application heap.
func (p ScoringBrowserSignals) New() js.Ref {
	return bindings.ScoringBrowserSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScoringBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.ScoringBrowserSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScoringBrowserSignals) Update(ref js.Ref) {
	bindings.ScoringBrowserSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScriptString = js.String

type ScriptURLString = js.String

type ScriptingPolicyReportBody struct {
	ReportBody
}

func (this ScriptingPolicyReportBody) Once() ScriptingPolicyReportBody {
	this.Ref().Once()
	return this
}

func (this ScriptingPolicyReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this ScriptingPolicyReportBody) FromRef(ref js.Ref) ScriptingPolicyReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this ScriptingPolicyReportBody) Free() {
	this.Ref().Free()
}

// ViolationType returns the value of property "ScriptingPolicyReportBody.violationType".
//
// The returned bool will be false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetScriptingPolicyReportBodyViolationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ViolationURL returns the value of property "ScriptingPolicyReportBody.violationURL".
//
// The returned bool will be false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetScriptingPolicyReportBodyViolationURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ViolationSample returns the value of property "ScriptingPolicyReportBody.violationSample".
//
// The returned bool will be false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationSample() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetScriptingPolicyReportBodyViolationSample(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lineno returns the value of property "ScriptingPolicyReportBody.lineno".
//
// The returned bool will be false if there is no such property.
func (this ScriptingPolicyReportBody) Lineno() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetScriptingPolicyReportBodyLineno(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Colno returns the value of property "ScriptingPolicyReportBody.colno".
//
// The returned bool will be false if there is no such property.
func (this ScriptingPolicyReportBody) Colno() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetScriptingPolicyReportBodyColno(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ToJSON calls the method "ScriptingPolicyReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this ScriptingPolicyReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallScriptingPolicyReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "ScriptingPolicyReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ScriptingPolicyReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ScriptingPolicyReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

type ScriptingPolicyViolationType uint32

const (
	_ ScriptingPolicyViolationType = iota

	ScriptingPolicyViolationType_EXTERNAL_SCRIPT
	ScriptingPolicyViolationType_INLINE_SCRIPT
	ScriptingPolicyViolationType_INLINE_EVENT_HANDLER
	ScriptingPolicyViolationType_EVAL
)

func (ScriptingPolicyViolationType) FromRef(str js.Ref) ScriptingPolicyViolationType {
	return ScriptingPolicyViolationType(bindings.ConstOfScriptingPolicyViolationType(str))
}

func (x ScriptingPolicyViolationType) String() (string, bool) {
	switch x {
	case ScriptingPolicyViolationType_EXTERNAL_SCRIPT:
		return "externalScript", true
	case ScriptingPolicyViolationType_INLINE_SCRIPT:
		return "inlineScript", true
	case ScriptingPolicyViolationType_INLINE_EVENT_HANDLER:
		return "inlineEventHandler", true
	case ScriptingPolicyViolationType_EVAL:
		return "eval", true
	default:
		return "", false
	}
}

type ScrollAxis uint32

const (
	_ ScrollAxis = iota

	ScrollAxis_BLOCK
	ScrollAxis_INLINE
	ScrollAxis_X
	ScrollAxis_Y
)

func (ScrollAxis) FromRef(str js.Ref) ScrollAxis {
	return ScrollAxis(bindings.ConstOfScrollAxis(str))
}

func (x ScrollAxis) String() (string, bool) {
	switch x {
	case ScrollAxis_BLOCK:
		return "block", true
	case ScrollAxis_INLINE:
		return "inline", true
	case ScrollAxis_X:
		return "x", true
	case ScrollAxis_Y:
		return "y", true
	default:
		return "", false
	}
}
