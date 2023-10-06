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
// It returns ok=false if there is no such property.
func (this SVGFETileElement) In1() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementIn1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFETileElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFETileElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFETileElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFETileElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFETileElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFETileElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFETileElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFETileElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFETileElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFETileElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFETileElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) BaseFrequencyX() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementBaseFrequencyX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BaseFrequencyY returns the value of property "SVGFETurbulenceElement.baseFrequencyY".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) BaseFrequencyY() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementBaseFrequencyY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumOctaves returns the value of property "SVGFETurbulenceElement.numOctaves".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) NumOctaves() (ret SVGAnimatedInteger, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementNumOctaves(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Seed returns the value of property "SVGFETurbulenceElement.seed".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Seed() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementSeed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StitchTiles returns the value of property "SVGFETurbulenceElement.stitchTiles".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) StitchTiles() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementStitchTiles(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "SVGFETurbulenceElement.type".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Type() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFETurbulenceElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFETurbulenceElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFETurbulenceElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFETurbulenceElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "SVGFETurbulenceElement.result".
//
// It returns ok=false if there is no such property.
func (this SVGFETurbulenceElement) Result() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFETurbulenceElementResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGFilterElement) FilterUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementFilterUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PrimitiveUnits returns the value of property "SVGFilterElement.primitiveUnits".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) PrimitiveUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementPrimitiveUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGFilterElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGFilterElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGFilterElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGFilterElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGFilterElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGFilterElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGFilterElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGForeignObjectElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGForeignObjectElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGForeignObjectElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGForeignObjectElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGForeignObjectElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGForeignObjectElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGForeignObjectElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGForeignObjectElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGForeignObjectElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGForeignObjectElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGForeignObjectElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGGeometryElement) PathLength() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGGeometryElementPathLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasIsPointInFill returns true if the method "SVGGeometryElement.isPointInFill" exists.
func (this SVGGeometryElement) HasIsPointInFill() bool {
	return js.True == bindings.HasSVGGeometryElementIsPointInFill(
		this.Ref(),
	)
}

// IsPointInFillFunc returns the method "SVGGeometryElement.isPointInFill".
func (this SVGGeometryElement) IsPointInFillFunc() (fn js.Func[func(point DOMPointInit) bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInFillFunc(
			this.Ref(),
		),
	)
}

// IsPointInFill calls the method "SVGGeometryElement.isPointInFill".
func (this SVGGeometryElement) IsPointInFill(point DOMPointInit) (ret bool) {
	bindings.CallSVGGeometryElementIsPointInFill(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryIsPointInFill calls the method "SVGGeometryElement.isPointInFill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryIsPointInFill(point DOMPointInit) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementIsPointInFill(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasIsPointInFill1 returns true if the method "SVGGeometryElement.isPointInFill" exists.
func (this SVGGeometryElement) HasIsPointInFill1() bool {
	return js.True == bindings.HasSVGGeometryElementIsPointInFill1(
		this.Ref(),
	)
}

// IsPointInFill1Func returns the method "SVGGeometryElement.isPointInFill".
func (this SVGGeometryElement) IsPointInFill1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInFill1Func(
			this.Ref(),
		),
	)
}

// IsPointInFill1 calls the method "SVGGeometryElement.isPointInFill".
func (this SVGGeometryElement) IsPointInFill1() (ret bool) {
	bindings.CallSVGGeometryElementIsPointInFill1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsPointInFill1 calls the method "SVGGeometryElement.isPointInFill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryIsPointInFill1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementIsPointInFill1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsPointInStroke returns true if the method "SVGGeometryElement.isPointInStroke" exists.
func (this SVGGeometryElement) HasIsPointInStroke() bool {
	return js.True == bindings.HasSVGGeometryElementIsPointInStroke(
		this.Ref(),
	)
}

// IsPointInStrokeFunc returns the method "SVGGeometryElement.isPointInStroke".
func (this SVGGeometryElement) IsPointInStrokeFunc() (fn js.Func[func(point DOMPointInit) bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "SVGGeometryElement.isPointInStroke".
func (this SVGGeometryElement) IsPointInStroke(point DOMPointInit) (ret bool) {
	bindings.CallSVGGeometryElementIsPointInStroke(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryIsPointInStroke calls the method "SVGGeometryElement.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryIsPointInStroke(point DOMPointInit) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementIsPointInStroke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasIsPointInStroke1 returns true if the method "SVGGeometryElement.isPointInStroke" exists.
func (this SVGGeometryElement) HasIsPointInStroke1() bool {
	return js.True == bindings.HasSVGGeometryElementIsPointInStroke1(
		this.Ref(),
	)
}

// IsPointInStroke1Func returns the method "SVGGeometryElement.isPointInStroke".
func (this SVGGeometryElement) IsPointInStroke1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.SVGGeometryElementIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "SVGGeometryElement.isPointInStroke".
func (this SVGGeometryElement) IsPointInStroke1() (ret bool) {
	bindings.CallSVGGeometryElementIsPointInStroke1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsPointInStroke1 calls the method "SVGGeometryElement.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryIsPointInStroke1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementIsPointInStroke1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetTotalLength returns true if the method "SVGGeometryElement.getTotalLength" exists.
func (this SVGGeometryElement) HasGetTotalLength() bool {
	return js.True == bindings.HasSVGGeometryElementGetTotalLength(
		this.Ref(),
	)
}

// GetTotalLengthFunc returns the method "SVGGeometryElement.getTotalLength".
func (this SVGGeometryElement) GetTotalLengthFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGGeometryElementGetTotalLengthFunc(
			this.Ref(),
		),
	)
}

// GetTotalLength calls the method "SVGGeometryElement.getTotalLength".
func (this SVGGeometryElement) GetTotalLength() (ret float32) {
	bindings.CallSVGGeometryElementGetTotalLength(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTotalLength calls the method "SVGGeometryElement.getTotalLength"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryGetTotalLength() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementGetTotalLength(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPointAtLength returns true if the method "SVGGeometryElement.getPointAtLength" exists.
func (this SVGGeometryElement) HasGetPointAtLength() bool {
	return js.True == bindings.HasSVGGeometryElementGetPointAtLength(
		this.Ref(),
	)
}

// GetPointAtLengthFunc returns the method "SVGGeometryElement.getPointAtLength".
func (this SVGGeometryElement) GetPointAtLengthFunc() (fn js.Func[func(distance float32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGGeometryElementGetPointAtLengthFunc(
			this.Ref(),
		),
	)
}

// GetPointAtLength calls the method "SVGGeometryElement.getPointAtLength".
func (this SVGGeometryElement) GetPointAtLength(distance float32) (ret DOMPoint) {
	bindings.CallSVGGeometryElementGetPointAtLength(
		this.Ref(), js.Pointer(&ret),
		float32(distance),
	)

	return
}

// TryGetPointAtLength calls the method "SVGGeometryElement.getPointAtLength"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGeometryElement) TryGetPointAtLength(distance float32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGeometryElementGetPointAtLength(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(distance),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGGradientElement) GradientUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGGradientElementGradientUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// GradientTransform returns the value of property "SVGGradientElement.gradientTransform".
//
// It returns ok=false if there is no such property.
func (this SVGGradientElement) GradientTransform() (ret SVGAnimatedTransformList, ok bool) {
	ok = js.True == bindings.GetSVGGradientElementGradientTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SpreadMethod returns the value of property "SVGGradientElement.spreadMethod".
//
// It returns ok=false if there is no such property.
func (this SVGGradientElement) SpreadMethod() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGGradientElementSpreadMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGGradientElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGGradientElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGGradientElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGGraphicsElement) Transform() (ret SVGAnimatedTransformList, ok bool) {
	ok = js.True == bindings.GetSVGGraphicsElementTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RequiredExtensions returns the value of property "SVGGraphicsElement.requiredExtensions".
//
// It returns ok=false if there is no such property.
func (this SVGGraphicsElement) RequiredExtensions() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGGraphicsElementRequiredExtensions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SystemLanguage returns the value of property "SVGGraphicsElement.systemLanguage".
//
// It returns ok=false if there is no such property.
func (this SVGGraphicsElement) SystemLanguage() (ret SVGStringList, ok bool) {
	ok = js.True == bindings.GetSVGGraphicsElementSystemLanguage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetBBox returns true if the method "SVGGraphicsElement.getBBox" exists.
func (this SVGGraphicsElement) HasGetBBox() bool {
	return js.True == bindings.HasSVGGraphicsElementGetBBox(
		this.Ref(),
	)
}

// GetBBoxFunc returns the method "SVGGraphicsElement.getBBox".
func (this SVGGraphicsElement) GetBBoxFunc() (fn js.Func[func(options SVGBoundingBoxOptions) DOMRect]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetBBoxFunc(
			this.Ref(),
		),
	)
}

// GetBBox calls the method "SVGGraphicsElement.getBBox".
func (this SVGGraphicsElement) GetBBox(options SVGBoundingBoxOptions) (ret DOMRect) {
	bindings.CallSVGGraphicsElementGetBBox(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBBox calls the method "SVGGraphicsElement.getBBox"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGraphicsElement) TryGetBBox(options SVGBoundingBoxOptions) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGraphicsElementGetBBox(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetBBox1 returns true if the method "SVGGraphicsElement.getBBox" exists.
func (this SVGGraphicsElement) HasGetBBox1() bool {
	return js.True == bindings.HasSVGGraphicsElementGetBBox1(
		this.Ref(),
	)
}

// GetBBox1Func returns the method "SVGGraphicsElement.getBBox".
func (this SVGGraphicsElement) GetBBox1Func() (fn js.Func[func() DOMRect]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetBBox1Func(
			this.Ref(),
		),
	)
}

// GetBBox1 calls the method "SVGGraphicsElement.getBBox".
func (this SVGGraphicsElement) GetBBox1() (ret DOMRect) {
	bindings.CallSVGGraphicsElementGetBBox1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBBox1 calls the method "SVGGraphicsElement.getBBox"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGraphicsElement) TryGetBBox1() (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGraphicsElementGetBBox1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetCTM returns true if the method "SVGGraphicsElement.getCTM" exists.
func (this SVGGraphicsElement) HasGetCTM() bool {
	return js.True == bindings.HasSVGGraphicsElementGetCTM(
		this.Ref(),
	)
}

// GetCTMFunc returns the method "SVGGraphicsElement.getCTM".
func (this SVGGraphicsElement) GetCTMFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetCTMFunc(
			this.Ref(),
		),
	)
}

// GetCTM calls the method "SVGGraphicsElement.getCTM".
func (this SVGGraphicsElement) GetCTM() (ret DOMMatrix) {
	bindings.CallSVGGraphicsElementGetCTM(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCTM calls the method "SVGGraphicsElement.getCTM"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGraphicsElement) TryGetCTM() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGraphicsElementGetCTM(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetScreenCTM returns true if the method "SVGGraphicsElement.getScreenCTM" exists.
func (this SVGGraphicsElement) HasGetScreenCTM() bool {
	return js.True == bindings.HasSVGGraphicsElementGetScreenCTM(
		this.Ref(),
	)
}

// GetScreenCTMFunc returns the method "SVGGraphicsElement.getScreenCTM".
func (this SVGGraphicsElement) GetScreenCTMFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.SVGGraphicsElementGetScreenCTMFunc(
			this.Ref(),
		),
	)
}

// GetScreenCTM calls the method "SVGGraphicsElement.getScreenCTM".
func (this SVGGraphicsElement) GetScreenCTM() (ret DOMMatrix) {
	bindings.CallSVGGraphicsElementGetScreenCTM(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetScreenCTM calls the method "SVGGraphicsElement.getScreenCTM"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGGraphicsElement) TryGetScreenCTM() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGGraphicsElementGetScreenCTM(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGLineElement) X1() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLineElementX1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y1 returns the value of property "SVGLineElement.y1".
//
// It returns ok=false if there is no such property.
func (this SVGLineElement) Y1() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLineElementY1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X2 returns the value of property "SVGLineElement.x2".
//
// It returns ok=false if there is no such property.
func (this SVGLineElement) X2() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLineElementX2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y2 returns the value of property "SVGLineElement.y2".
//
// It returns ok=false if there is no such property.
func (this SVGLineElement) Y2() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLineElementY2(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGLinearGradientElement) X1() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLinearGradientElementX1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y1 returns the value of property "SVGLinearGradientElement.y1".
//
// It returns ok=false if there is no such property.
func (this SVGLinearGradientElement) Y1() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLinearGradientElementY1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X2 returns the value of property "SVGLinearGradientElement.x2".
//
// It returns ok=false if there is no such property.
func (this SVGLinearGradientElement) X2() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLinearGradientElementX2(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y2 returns the value of property "SVGLinearGradientElement.y2".
//
// It returns ok=false if there is no such property.
func (this SVGLinearGradientElement) Y2() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGLinearGradientElementY2(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGMPathElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGMPathElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) RefX() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementRefX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RefY returns the value of property "SVGMarkerElement.refY".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) RefY() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementRefY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MarkerUnits returns the value of property "SVGMarkerElement.markerUnits".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) MarkerUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementMarkerUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MarkerWidth returns the value of property "SVGMarkerElement.markerWidth".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) MarkerWidth() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementMarkerWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MarkerHeight returns the value of property "SVGMarkerElement.markerHeight".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) MarkerHeight() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementMarkerHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OrientType returns the value of property "SVGMarkerElement.orientType".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) OrientType() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementOrientType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OrientAngle returns the value of property "SVGMarkerElement.orientAngle".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) OrientAngle() (ret SVGAnimatedAngle, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementOrientAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Orient returns the value of property "SVGMarkerElement.orient".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) Orient() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementOrient(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOrient sets the value of property "SVGMarkerElement.orient" to val.
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
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementViewBox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGMarkerElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGMarkerElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGMarkerElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetOrientToAuto returns true if the method "SVGMarkerElement.setOrientToAuto" exists.
func (this SVGMarkerElement) HasSetOrientToAuto() bool {
	return js.True == bindings.HasSVGMarkerElementSetOrientToAuto(
		this.Ref(),
	)
}

// SetOrientToAutoFunc returns the method "SVGMarkerElement.setOrientToAuto".
func (this SVGMarkerElement) SetOrientToAutoFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGMarkerElementSetOrientToAutoFunc(
			this.Ref(),
		),
	)
}

// SetOrientToAuto calls the method "SVGMarkerElement.setOrientToAuto".
func (this SVGMarkerElement) SetOrientToAuto() (ret js.Void) {
	bindings.CallSVGMarkerElementSetOrientToAuto(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetOrientToAuto calls the method "SVGMarkerElement.setOrientToAuto"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGMarkerElement) TrySetOrientToAuto() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGMarkerElementSetOrientToAuto(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetOrientToAngle returns true if the method "SVGMarkerElement.setOrientToAngle" exists.
func (this SVGMarkerElement) HasSetOrientToAngle() bool {
	return js.True == bindings.HasSVGMarkerElementSetOrientToAngle(
		this.Ref(),
	)
}

// SetOrientToAngleFunc returns the method "SVGMarkerElement.setOrientToAngle".
func (this SVGMarkerElement) SetOrientToAngleFunc() (fn js.Func[func(angle SVGAngle)]) {
	return fn.FromRef(
		bindings.SVGMarkerElementSetOrientToAngleFunc(
			this.Ref(),
		),
	)
}

// SetOrientToAngle calls the method "SVGMarkerElement.setOrientToAngle".
func (this SVGMarkerElement) SetOrientToAngle(angle SVGAngle) (ret js.Void) {
	bindings.CallSVGMarkerElementSetOrientToAngle(
		this.Ref(), js.Pointer(&ret),
		angle.Ref(),
	)

	return
}

// TrySetOrientToAngle calls the method "SVGMarkerElement.setOrientToAngle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGMarkerElement) TrySetOrientToAngle(angle SVGAngle) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGMarkerElementSetOrientToAngle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		angle.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGMaskElement) MaskUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementMaskUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaskContentUnits returns the value of property "SVGMaskElement.maskContentUnits".
//
// It returns ok=false if there is no such property.
func (this SVGMaskElement) MaskContentUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementMaskContentUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGMaskElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGMaskElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGMaskElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGMaskElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGMaskElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGMaskElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGMaskElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGMaskElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGMaskElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGPatternElement) PatternUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementPatternUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PatternContentUnits returns the value of property "SVGPatternElement.patternContentUnits".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) PatternContentUnits() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementPatternContentUnits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PatternTransform returns the value of property "SVGPatternElement.patternTransform".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) PatternTransform() (ret SVGAnimatedTransformList, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementPatternTransform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "SVGPatternElement.x".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGPatternElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGPatternElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGPatternElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ViewBox returns the value of property "SVGPatternElement.viewBox".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementViewBox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGPatternElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGPatternElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGPatternElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGPatternElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGPointList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGPointListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfItems returns the value of property "SVGPointList.numberOfItems".
//
// It returns ok=false if there is no such property.
func (this SVGPointList) NumberOfItems() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSVGPointListNumberOfItems(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClear returns true if the method "SVGPointList.clear" exists.
func (this SVGPointList) HasClear() bool {
	return js.True == bindings.HasSVGPointListClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "SVGPointList.clear".
func (this SVGPointList) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SVGPointListClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SVGPointList.clear".
func (this SVGPointList) Clear() (ret js.Void) {
	bindings.CallSVGPointListClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SVGPointList.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitialize returns true if the method "SVGPointList.initialize" exists.
func (this SVGPointList) HasInitialize() bool {
	return js.True == bindings.HasSVGPointListInitialize(
		this.Ref(),
	)
}

// InitializeFunc returns the method "SVGPointList.initialize".
func (this SVGPointList) InitializeFunc() (fn js.Func[func(newItem DOMPoint) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListInitializeFunc(
			this.Ref(),
		),
	)
}

// Initialize calls the method "SVGPointList.initialize".
func (this SVGPointList) Initialize(newItem DOMPoint) (ret DOMPoint) {
	bindings.CallSVGPointListInitialize(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryInitialize calls the method "SVGPointList.initialize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryInitialize(newItem DOMPoint) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListInitialize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasGetItem returns true if the method "SVGPointList.getItem" exists.
func (this SVGPointList) HasGetItem() bool {
	return js.True == bindings.HasSVGPointListGetItem(
		this.Ref(),
	)
}

// GetItemFunc returns the method "SVGPointList.getItem".
func (this SVGPointList) GetItemFunc() (fn js.Func[func(index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListGetItemFunc(
			this.Ref(),
		),
	)
}

// GetItem calls the method "SVGPointList.getItem".
func (this SVGPointList) GetItem(index uint32) (ret DOMPoint) {
	bindings.CallSVGPointListGetItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetItem calls the method "SVGPointList.getItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryGetItem(index uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListGetItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasInsertItemBefore returns true if the method "SVGPointList.insertItemBefore" exists.
func (this SVGPointList) HasInsertItemBefore() bool {
	return js.True == bindings.HasSVGPointListInsertItemBefore(
		this.Ref(),
	)
}

// InsertItemBeforeFunc returns the method "SVGPointList.insertItemBefore".
func (this SVGPointList) InsertItemBeforeFunc() (fn js.Func[func(newItem DOMPoint, index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListInsertItemBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertItemBefore calls the method "SVGPointList.insertItemBefore".
func (this SVGPointList) InsertItemBefore(newItem DOMPoint, index uint32) (ret DOMPoint) {
	bindings.CallSVGPointListInsertItemBefore(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryInsertItemBefore calls the method "SVGPointList.insertItemBefore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryInsertItemBefore(newItem DOMPoint, index uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListInsertItemBefore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasReplaceItem returns true if the method "SVGPointList.replaceItem" exists.
func (this SVGPointList) HasReplaceItem() bool {
	return js.True == bindings.HasSVGPointListReplaceItem(
		this.Ref(),
	)
}

// ReplaceItemFunc returns the method "SVGPointList.replaceItem".
func (this SVGPointList) ReplaceItemFunc() (fn js.Func[func(newItem DOMPoint, index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListReplaceItemFunc(
			this.Ref(),
		),
	)
}

// ReplaceItem calls the method "SVGPointList.replaceItem".
func (this SVGPointList) ReplaceItem(newItem DOMPoint, index uint32) (ret DOMPoint) {
	bindings.CallSVGPointListReplaceItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// TryReplaceItem calls the method "SVGPointList.replaceItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryReplaceItem(newItem DOMPoint, index uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListReplaceItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
		uint32(index),
	)

	return
}

// HasRemoveItem returns true if the method "SVGPointList.removeItem" exists.
func (this SVGPointList) HasRemoveItem() bool {
	return js.True == bindings.HasSVGPointListRemoveItem(
		this.Ref(),
	)
}

// RemoveItemFunc returns the method "SVGPointList.removeItem".
func (this SVGPointList) RemoveItemFunc() (fn js.Func[func(index uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListRemoveItemFunc(
			this.Ref(),
		),
	)
}

// RemoveItem calls the method "SVGPointList.removeItem".
func (this SVGPointList) RemoveItem(index uint32) (ret DOMPoint) {
	bindings.CallSVGPointListRemoveItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryRemoveItem calls the method "SVGPointList.removeItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryRemoveItem(index uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListRemoveItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasAppendItem returns true if the method "SVGPointList.appendItem" exists.
func (this SVGPointList) HasAppendItem() bool {
	return js.True == bindings.HasSVGPointListAppendItem(
		this.Ref(),
	)
}

// AppendItemFunc returns the method "SVGPointList.appendItem".
func (this SVGPointList) AppendItemFunc() (fn js.Func[func(newItem DOMPoint) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGPointListAppendItemFunc(
			this.Ref(),
		),
	)
}

// AppendItem calls the method "SVGPointList.appendItem".
func (this SVGPointList) AppendItem(newItem DOMPoint) (ret DOMPoint) {
	bindings.CallSVGPointListAppendItem(
		this.Ref(), js.Pointer(&ret),
		newItem.Ref(),
	)

	return
}

// TryAppendItem calls the method "SVGPointList.appendItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TryAppendItem(newItem DOMPoint) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListAppendItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		newItem.Ref(),
	)

	return
}

// HasSet returns true if the method "SVGPointList." exists.
func (this SVGPointList) HasSet() bool {
	return js.True == bindings.HasSVGPointListSet(
		this.Ref(),
	)
}

// SetFunc returns the method "SVGPointList.".
func (this SVGPointList) SetFunc() (fn js.Func[func(index uint32, newItem DOMPoint)]) {
	return fn.FromRef(
		bindings.SVGPointListSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "SVGPointList.".
func (this SVGPointList) Set(index uint32, newItem DOMPoint) (ret js.Void) {
	bindings.CallSVGPointListSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		newItem.Ref(),
	)

	return
}

// TrySet calls the method "SVGPointList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGPointList) TrySet(index uint32, newItem DOMPoint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGPointListSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		newItem.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGPolygonElement) Points() (ret SVGPointList, ok bool) {
	ok = js.True == bindings.GetSVGPolygonElementPoints(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimatedPoints returns the value of property "SVGPolygonElement.animatedPoints".
//
// It returns ok=false if there is no such property.
func (this SVGPolygonElement) AnimatedPoints() (ret SVGPointList, ok bool) {
	ok = js.True == bindings.GetSVGPolygonElementAnimatedPoints(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGPolylineElement) Points() (ret SVGPointList, ok bool) {
	ok = js.True == bindings.GetSVGPolylineElementPoints(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AnimatedPoints returns the value of property "SVGPolylineElement.animatedPoints".
//
// It returns ok=false if there is no such property.
func (this SVGPolylineElement) AnimatedPoints() (ret SVGPointList, ok bool) {
	ok = js.True == bindings.GetSVGPolylineElementAnimatedPoints(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) Cx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementCx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cy returns the value of property "SVGRadialGradientElement.cy".
//
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) Cy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementCy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// R returns the value of property "SVGRadialGradientElement.r".
//
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) R() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementR(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fx returns the value of property "SVGRadialGradientElement.fx".
//
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) Fx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementFx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fy returns the value of property "SVGRadialGradientElement.fy".
//
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) Fy() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementFy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fr returns the value of property "SVGRadialGradientElement.fr".
//
// It returns ok=false if there is no such property.
func (this SVGRadialGradientElement) Fr() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRadialGradientElementFr(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGRectElement) X() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGRectElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGRectElement) Y() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "SVGRectElement.width".
//
// It returns ok=false if there is no such property.
func (this SVGRectElement) Width() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "SVGRectElement.height".
//
// It returns ok=false if there is no such property.
func (this SVGRectElement) Height() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rx returns the value of property "SVGRectElement.rx".
//
// It returns ok=false if there is no such property.
func (this SVGRectElement) Rx() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementRx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ry returns the value of property "SVGRectElement.ry".
//
// It returns ok=false if there is no such property.
func (this SVGRectElement) Ry() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGRectElementRy(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGStopElement) Offset() (ret SVGAnimatedNumber, ok bool) {
	ok = js.True == bindings.GetSVGStopElementOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGStyleElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGStyleElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "SVGStyleElement.type" to val.
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
// It returns ok=false if there is no such property.
func (this SVGStyleElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGStyleElementMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "SVGStyleElement.media" to val.
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
// It returns ok=false if there is no such property.
func (this SVGStyleElement) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSVGStyleElementTitle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "SVGStyleElement.title" to val.
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
// It returns ok=false if there is no such property.
func (this SVGStyleElement) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetSVGStyleElementSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGSymbolElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGSymbolElementViewBox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGSymbolElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGSymbolElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGSymbolElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGTextContentElement) TextLength() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGTextContentElementTextLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LengthAdjust returns the value of property "SVGTextContentElement.lengthAdjust".
//
// It returns ok=false if there is no such property.
func (this SVGTextContentElement) LengthAdjust() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGTextContentElementLengthAdjust(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetNumberOfChars returns true if the method "SVGTextContentElement.getNumberOfChars" exists.
func (this SVGTextContentElement) HasGetNumberOfChars() bool {
	return js.True == bindings.HasSVGTextContentElementGetNumberOfChars(
		this.Ref(),
	)
}

// GetNumberOfCharsFunc returns the method "SVGTextContentElement.getNumberOfChars".
func (this SVGTextContentElement) GetNumberOfCharsFunc() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetNumberOfCharsFunc(
			this.Ref(),
		),
	)
}

// GetNumberOfChars calls the method "SVGTextContentElement.getNumberOfChars".
func (this SVGTextContentElement) GetNumberOfChars() (ret int32) {
	bindings.CallSVGTextContentElementGetNumberOfChars(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetNumberOfChars calls the method "SVGTextContentElement.getNumberOfChars"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetNumberOfChars() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetNumberOfChars(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetComputedTextLength returns true if the method "SVGTextContentElement.getComputedTextLength" exists.
func (this SVGTextContentElement) HasGetComputedTextLength() bool {
	return js.True == bindings.HasSVGTextContentElementGetComputedTextLength(
		this.Ref(),
	)
}

// GetComputedTextLengthFunc returns the method "SVGTextContentElement.getComputedTextLength".
func (this SVGTextContentElement) GetComputedTextLengthFunc() (fn js.Func[func() float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetComputedTextLengthFunc(
			this.Ref(),
		),
	)
}

// GetComputedTextLength calls the method "SVGTextContentElement.getComputedTextLength".
func (this SVGTextContentElement) GetComputedTextLength() (ret float32) {
	bindings.CallSVGTextContentElementGetComputedTextLength(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetComputedTextLength calls the method "SVGTextContentElement.getComputedTextLength"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetComputedTextLength() (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetComputedTextLength(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSubStringLength returns true if the method "SVGTextContentElement.getSubStringLength" exists.
func (this SVGTextContentElement) HasGetSubStringLength() bool {
	return js.True == bindings.HasSVGTextContentElementGetSubStringLength(
		this.Ref(),
	)
}

// GetSubStringLengthFunc returns the method "SVGTextContentElement.getSubStringLength".
func (this SVGTextContentElement) GetSubStringLengthFunc() (fn js.Func[func(charnum uint32, nchars uint32) float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetSubStringLengthFunc(
			this.Ref(),
		),
	)
}

// GetSubStringLength calls the method "SVGTextContentElement.getSubStringLength".
func (this SVGTextContentElement) GetSubStringLength(charnum uint32, nchars uint32) (ret float32) {
	bindings.CallSVGTextContentElementGetSubStringLength(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
		uint32(nchars),
	)

	return
}

// TryGetSubStringLength calls the method "SVGTextContentElement.getSubStringLength"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetSubStringLength(charnum uint32, nchars uint32) (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetSubStringLength(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
		uint32(nchars),
	)

	return
}

// HasGetStartPositionOfChar returns true if the method "SVGTextContentElement.getStartPositionOfChar" exists.
func (this SVGTextContentElement) HasGetStartPositionOfChar() bool {
	return js.True == bindings.HasSVGTextContentElementGetStartPositionOfChar(
		this.Ref(),
	)
}

// GetStartPositionOfCharFunc returns the method "SVGTextContentElement.getStartPositionOfChar".
func (this SVGTextContentElement) GetStartPositionOfCharFunc() (fn js.Func[func(charnum uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetStartPositionOfCharFunc(
			this.Ref(),
		),
	)
}

// GetStartPositionOfChar calls the method "SVGTextContentElement.getStartPositionOfChar".
func (this SVGTextContentElement) GetStartPositionOfChar(charnum uint32) (ret DOMPoint) {
	bindings.CallSVGTextContentElementGetStartPositionOfChar(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
	)

	return
}

// TryGetStartPositionOfChar calls the method "SVGTextContentElement.getStartPositionOfChar"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetStartPositionOfChar(charnum uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetStartPositionOfChar(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
	)

	return
}

// HasGetEndPositionOfChar returns true if the method "SVGTextContentElement.getEndPositionOfChar" exists.
func (this SVGTextContentElement) HasGetEndPositionOfChar() bool {
	return js.True == bindings.HasSVGTextContentElementGetEndPositionOfChar(
		this.Ref(),
	)
}

// GetEndPositionOfCharFunc returns the method "SVGTextContentElement.getEndPositionOfChar".
func (this SVGTextContentElement) GetEndPositionOfCharFunc() (fn js.Func[func(charnum uint32) DOMPoint]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetEndPositionOfCharFunc(
			this.Ref(),
		),
	)
}

// GetEndPositionOfChar calls the method "SVGTextContentElement.getEndPositionOfChar".
func (this SVGTextContentElement) GetEndPositionOfChar(charnum uint32) (ret DOMPoint) {
	bindings.CallSVGTextContentElementGetEndPositionOfChar(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
	)

	return
}

// TryGetEndPositionOfChar calls the method "SVGTextContentElement.getEndPositionOfChar"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetEndPositionOfChar(charnum uint32) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetEndPositionOfChar(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
	)

	return
}

// HasGetExtentOfChar returns true if the method "SVGTextContentElement.getExtentOfChar" exists.
func (this SVGTextContentElement) HasGetExtentOfChar() bool {
	return js.True == bindings.HasSVGTextContentElementGetExtentOfChar(
		this.Ref(),
	)
}

// GetExtentOfCharFunc returns the method "SVGTextContentElement.getExtentOfChar".
func (this SVGTextContentElement) GetExtentOfCharFunc() (fn js.Func[func(charnum uint32) DOMRect]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetExtentOfCharFunc(
			this.Ref(),
		),
	)
}

// GetExtentOfChar calls the method "SVGTextContentElement.getExtentOfChar".
func (this SVGTextContentElement) GetExtentOfChar(charnum uint32) (ret DOMRect) {
	bindings.CallSVGTextContentElementGetExtentOfChar(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
	)

	return
}

// TryGetExtentOfChar calls the method "SVGTextContentElement.getExtentOfChar"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetExtentOfChar(charnum uint32) (ret DOMRect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetExtentOfChar(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
	)

	return
}

// HasGetRotationOfChar returns true if the method "SVGTextContentElement.getRotationOfChar" exists.
func (this SVGTextContentElement) HasGetRotationOfChar() bool {
	return js.True == bindings.HasSVGTextContentElementGetRotationOfChar(
		this.Ref(),
	)
}

// GetRotationOfCharFunc returns the method "SVGTextContentElement.getRotationOfChar".
func (this SVGTextContentElement) GetRotationOfCharFunc() (fn js.Func[func(charnum uint32) float32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetRotationOfCharFunc(
			this.Ref(),
		),
	)
}

// GetRotationOfChar calls the method "SVGTextContentElement.getRotationOfChar".
func (this SVGTextContentElement) GetRotationOfChar(charnum uint32) (ret float32) {
	bindings.CallSVGTextContentElementGetRotationOfChar(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
	)

	return
}

// TryGetRotationOfChar calls the method "SVGTextContentElement.getRotationOfChar"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetRotationOfChar(charnum uint32) (ret float32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetRotationOfChar(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
	)

	return
}

// HasGetCharNumAtPosition returns true if the method "SVGTextContentElement.getCharNumAtPosition" exists.
func (this SVGTextContentElement) HasGetCharNumAtPosition() bool {
	return js.True == bindings.HasSVGTextContentElementGetCharNumAtPosition(
		this.Ref(),
	)
}

// GetCharNumAtPositionFunc returns the method "SVGTextContentElement.getCharNumAtPosition".
func (this SVGTextContentElement) GetCharNumAtPositionFunc() (fn js.Func[func(point DOMPointInit) int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetCharNumAtPositionFunc(
			this.Ref(),
		),
	)
}

// GetCharNumAtPosition calls the method "SVGTextContentElement.getCharNumAtPosition".
func (this SVGTextContentElement) GetCharNumAtPosition(point DOMPointInit) (ret int32) {
	bindings.CallSVGTextContentElementGetCharNumAtPosition(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&point),
	)

	return
}

// TryGetCharNumAtPosition calls the method "SVGTextContentElement.getCharNumAtPosition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetCharNumAtPosition(point DOMPointInit) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetCharNumAtPosition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
	)

	return
}

// HasGetCharNumAtPosition1 returns true if the method "SVGTextContentElement.getCharNumAtPosition" exists.
func (this SVGTextContentElement) HasGetCharNumAtPosition1() bool {
	return js.True == bindings.HasSVGTextContentElementGetCharNumAtPosition1(
		this.Ref(),
	)
}

// GetCharNumAtPosition1Func returns the method "SVGTextContentElement.getCharNumAtPosition".
func (this SVGTextContentElement) GetCharNumAtPosition1Func() (fn js.Func[func() int32]) {
	return fn.FromRef(
		bindings.SVGTextContentElementGetCharNumAtPosition1Func(
			this.Ref(),
		),
	)
}

// GetCharNumAtPosition1 calls the method "SVGTextContentElement.getCharNumAtPosition".
func (this SVGTextContentElement) GetCharNumAtPosition1() (ret int32) {
	bindings.CallSVGTextContentElementGetCharNumAtPosition1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCharNumAtPosition1 calls the method "SVGTextContentElement.getCharNumAtPosition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TryGetCharNumAtPosition1() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementGetCharNumAtPosition1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSelectSubString returns true if the method "SVGTextContentElement.selectSubString" exists.
func (this SVGTextContentElement) HasSelectSubString() bool {
	return js.True == bindings.HasSVGTextContentElementSelectSubString(
		this.Ref(),
	)
}

// SelectSubStringFunc returns the method "SVGTextContentElement.selectSubString".
func (this SVGTextContentElement) SelectSubStringFunc() (fn js.Func[func(charnum uint32, nchars uint32)]) {
	return fn.FromRef(
		bindings.SVGTextContentElementSelectSubStringFunc(
			this.Ref(),
		),
	)
}

// SelectSubString calls the method "SVGTextContentElement.selectSubString".
func (this SVGTextContentElement) SelectSubString(charnum uint32, nchars uint32) (ret js.Void) {
	bindings.CallSVGTextContentElementSelectSubString(
		this.Ref(), js.Pointer(&ret),
		uint32(charnum),
		uint32(nchars),
	)

	return
}

// TrySelectSubString calls the method "SVGTextContentElement.selectSubString"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SVGTextContentElement) TrySelectSubString(charnum uint32, nchars uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySVGTextContentElementSelectSubString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(charnum),
		uint32(nchars),
	)

	return
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
// It returns ok=false if there is no such property.
func (this SVGTextPathElement) StartOffset() (ret SVGAnimatedLength, ok bool) {
	ok = js.True == bindings.GetSVGTextPathElementStartOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Method returns the value of property "SVGTextPathElement.method".
//
// It returns ok=false if there is no such property.
func (this SVGTextPathElement) Method() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGTextPathElementMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Spacing returns the value of property "SVGTextPathElement.spacing".
//
// It returns ok=false if there is no such property.
func (this SVGTextPathElement) Spacing() (ret SVGAnimatedEnumeration, ok bool) {
	ok = js.True == bindings.GetSVGTextPathElementSpacing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "SVGTextPathElement.href".
//
// It returns ok=false if there is no such property.
func (this SVGTextPathElement) Href() (ret SVGAnimatedString, ok bool) {
	ok = js.True == bindings.GetSVGTextPathElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGTextPositioningElement) X() (ret SVGAnimatedLengthList, ok bool) {
	ok = js.True == bindings.GetSVGTextPositioningElementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "SVGTextPositioningElement.y".
//
// It returns ok=false if there is no such property.
func (this SVGTextPositioningElement) Y() (ret SVGAnimatedLengthList, ok bool) {
	ok = js.True == bindings.GetSVGTextPositioningElementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dx returns the value of property "SVGTextPositioningElement.dx".
//
// It returns ok=false if there is no such property.
func (this SVGTextPositioningElement) Dx() (ret SVGAnimatedLengthList, ok bool) {
	ok = js.True == bindings.GetSVGTextPositioningElementDx(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dy returns the value of property "SVGTextPositioningElement.dy".
//
// It returns ok=false if there is no such property.
func (this SVGTextPositioningElement) Dy() (ret SVGAnimatedLengthList, ok bool) {
	ok = js.True == bindings.GetSVGTextPositioningElementDy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rotate returns the value of property "SVGTextPositioningElement.rotate".
//
// It returns ok=false if there is no such property.
func (this SVGTextPositioningElement) Rotate() (ret SVGAnimatedNumberList, ok bool) {
	ok = js.True == bindings.GetSVGTextPositioningElementRotate(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SVGViewElement) ViewBox() (ret SVGAnimatedRect, ok bool) {
	ok = js.True == bindings.GetSVGViewElementViewBox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreserveAspectRatio returns the value of property "SVGViewElement.preserveAspectRatio".
//
// It returns ok=false if there is no such property.
func (this SVGViewElement) PreserveAspectRatio() (ret SVGAnimatedPreserveAspectRatio, ok bool) {
	ok = js.True == bindings.GetSVGViewElementPreserveAspectRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type ScoreAdOutput struct {
	// Desirability is "ScoreAdOutput.desirability"
	//
	// Required
	Desirability float64
	// Bid is "ScoreAdOutput.bid"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bid MUST be set to true to make this field effective.
	Bid float64
	// BidCurrency is "ScoreAdOutput.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// IncomingBidInSellerCurrency is "ScoreAdOutput.incomingBidInSellerCurrency"
	//
	// Optional
	//
	// NOTE: FFI_USE_IncomingBidInSellerCurrency MUST be set to true to make this field effective.
	IncomingBidInSellerCurrency float64
	// AllowComponentAuction is "ScoreAdOutput.allowComponentAuction"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AllowComponentAuction MUST be set to true to make this field effective.
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

// New creates a new ScoreAdOutput in the application heap.
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
	//
	// NOTE: FFI_USE_DataVersion MUST be set to true to make this field effective.
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

// New creates a new ScoringBrowserSignals in the application heap.
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
// It returns ok=false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetScriptingPolicyReportBodyViolationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ViolationURL returns the value of property "ScriptingPolicyReportBody.violationURL".
//
// It returns ok=false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetScriptingPolicyReportBodyViolationURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ViolationSample returns the value of property "ScriptingPolicyReportBody.violationSample".
//
// It returns ok=false if there is no such property.
func (this ScriptingPolicyReportBody) ViolationSample() (ret js.String, ok bool) {
	ok = js.True == bindings.GetScriptingPolicyReportBodyViolationSample(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Lineno returns the value of property "ScriptingPolicyReportBody.lineno".
//
// It returns ok=false if there is no such property.
func (this ScriptingPolicyReportBody) Lineno() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScriptingPolicyReportBodyLineno(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Colno returns the value of property "ScriptingPolicyReportBody.colno".
//
// It returns ok=false if there is no such property.
func (this ScriptingPolicyReportBody) Colno() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScriptingPolicyReportBodyColno(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "ScriptingPolicyReportBody.toJSON" exists.
func (this ScriptingPolicyReportBody) HasToJSON() bool {
	return js.True == bindings.HasScriptingPolicyReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "ScriptingPolicyReportBody.toJSON".
func (this ScriptingPolicyReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.ScriptingPolicyReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "ScriptingPolicyReportBody.toJSON".
func (this ScriptingPolicyReportBody) ToJSON() (ret js.Object) {
	bindings.CallScriptingPolicyReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "ScriptingPolicyReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ScriptingPolicyReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryScriptingPolicyReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
