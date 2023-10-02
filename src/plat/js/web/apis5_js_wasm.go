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

type OneOf_HTMLScriptElement_SVGScriptElement struct {
	ref js.Ref
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) FromRef(ref js.Ref) OneOf_HTMLScriptElement_SVGScriptElement {
	return OneOf_HTMLScriptElement_SVGScriptElement{
		ref: ref,
	}
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) HTMLScriptElement() HTMLScriptElement {
	return HTMLScriptElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLScriptElement_SVGScriptElement) SVGScriptElement() SVGScriptElement {
	return SVGScriptElement{}.FromRef(x.ref)
}

type HTMLOrSVGScriptElement = OneOf_HTMLScriptElement_SVGScriptElement

type DocumentVisibilityState uint32

const (
	_ DocumentVisibilityState = iota

	DocumentVisibilityState_VISIBLE
	DocumentVisibilityState_HIDDEN
)

func (DocumentVisibilityState) FromRef(str js.Ref) DocumentVisibilityState {
	return DocumentVisibilityState(bindings.ConstOfDocumentVisibilityState(str))
}

func (x DocumentVisibilityState) String() (string, bool) {
	switch x {
	case DocumentVisibilityState_VISIBLE:
		return "visible", true
	case DocumentVisibilityState_HIDDEN:
		return "hidden", true
	default:
		return "", false
	}
}

type OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView struct {
	ref js.Ref
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FromRef(ref js.Ref) OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView {
	return OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView{
		ref: ref,
	}
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

type FontFaceDescriptors struct {
	// Style is "FontFaceDescriptors.style"
	//
	// Optional, defaults to "normal".
	Style js.String
	// Weight is "FontFaceDescriptors.weight"
	//
	// Optional, defaults to "normal".
	Weight js.String
	// Stretch is "FontFaceDescriptors.stretch"
	//
	// Optional, defaults to "normal".
	Stretch js.String
	// UnicodeRange is "FontFaceDescriptors.unicodeRange"
	//
	// Optional, defaults to "U+0-10FFFF".
	UnicodeRange js.String
	// Variant is "FontFaceDescriptors.variant"
	//
	// Optional, defaults to "normal".
	Variant js.String
	// FeatureSettings is "FontFaceDescriptors.featureSettings"
	//
	// Optional, defaults to "normal".
	FeatureSettings js.String
	// VariationSettings is "FontFaceDescriptors.variationSettings"
	//
	// Optional, defaults to "normal".
	VariationSettings js.String
	// Display is "FontFaceDescriptors.display"
	//
	// Optional, defaults to "auto".
	Display js.String
	// AscentOverride is "FontFaceDescriptors.ascentOverride"
	//
	// Optional, defaults to "normal".
	AscentOverride js.String
	// DescentOverride is "FontFaceDescriptors.descentOverride"
	//
	// Optional, defaults to "normal".
	DescentOverride js.String
	// LineGapOverride is "FontFaceDescriptors.lineGapOverride"
	//
	// Optional, defaults to "normal".
	LineGapOverride js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FontFaceDescriptors with all fields set.
func (p FontFaceDescriptors) FromRef(ref js.Ref) FontFaceDescriptors {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FontFaceDescriptors FontFaceDescriptors [// FontFaceDescriptors] [0x1400033e960 0x1400033ea00 0x1400033eb40 0x1400033ebe0 0x1400033ec80 0x1400033ed20 0x1400033edc0 0x1400033ee60 0x1400033ef00 0x1400033efa0 0x1400033f040] 0x14001af2bd0 {0 0}} in the application heap.
func (p FontFaceDescriptors) New() js.Ref {
	return bindings.FontFaceDescriptorsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FontFaceDescriptors) UpdateFrom(ref js.Ref) {
	bindings.FontFaceDescriptorsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FontFaceDescriptors) Update(ref js.Ref) {
	bindings.FontFaceDescriptorsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FontFaceLoadStatus uint32

const (
	_ FontFaceLoadStatus = iota

	FontFaceLoadStatus_UNLOADED
	FontFaceLoadStatus_LOADING
	FontFaceLoadStatus_LOADED
	FontFaceLoadStatus_ERROR
)

func (FontFaceLoadStatus) FromRef(str js.Ref) FontFaceLoadStatus {
	return FontFaceLoadStatus(bindings.ConstOfFontFaceLoadStatus(str))
}

func (x FontFaceLoadStatus) String() (string, bool) {
	switch x {
	case FontFaceLoadStatus_UNLOADED:
		return "unloaded", true
	case FontFaceLoadStatus_LOADING:
		return "loading", true
	case FontFaceLoadStatus_LOADED:
		return "loaded", true
	case FontFaceLoadStatus_ERROR:
		return "error", true
	default:
		return "", false
	}
}

type FontFaceFeatures struct {
	ref js.Ref
}

func (this FontFaceFeatures) Once() FontFaceFeatures {
	this.Ref().Once()
	return this
}

func (this FontFaceFeatures) Ref() js.Ref {
	return this.ref
}

func (this FontFaceFeatures) FromRef(ref js.Ref) FontFaceFeatures {
	this.ref = ref
	return this
}

func (this FontFaceFeatures) Free() {
	this.Ref().Free()
}

type FontFaceVariations struct {
	ref js.Ref
}

func (this FontFaceVariations) Once() FontFaceVariations {
	this.Ref().Once()
	return this
}

func (this FontFaceVariations) Ref() js.Ref {
	return this.ref
}

func (this FontFaceVariations) FromRef(ref js.Ref) FontFaceVariations {
	this.ref = ref
	return this
}

func (this FontFaceVariations) Free() {
	this.Ref().Free()
}

type FontFacePalette struct {
	ref js.Ref
}

func (this FontFacePalette) Once() FontFacePalette {
	this.Ref().Once()
	return this
}

func (this FontFacePalette) Ref() js.Ref {
	return this.ref
}

func (this FontFacePalette) FromRef(ref js.Ref) FontFacePalette {
	this.ref = ref
	return this
}

func (this FontFacePalette) Free() {
	this.Ref().Free()
}

// Length returns the value of property "FontFacePalette.length".
//
// The returned bool will be false if there is no such property.
func (this FontFacePalette) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetFontFacePaletteLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// UsableWithLightBackground returns the value of property "FontFacePalette.usableWithLightBackground".
//
// The returned bool will be false if there is no such property.
func (this FontFacePalette) UsableWithLightBackground() (bool, bool) {
	var _ok bool
	_ret := bindings.GetFontFacePaletteUsableWithLightBackground(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// UsableWithDarkBackground returns the value of property "FontFacePalette.usableWithDarkBackground".
//
// The returned bool will be false if there is no such property.
func (this FontFacePalette) UsableWithDarkBackground() (bool, bool) {
	var _ok bool
	_ret := bindings.GetFontFacePaletteUsableWithDarkBackground(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Get calls the method "FontFacePalette.".
//
// The returned bool will be false if there is no such method.
func (this FontFacePalette) Get(index uint32) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallFontFacePaletteGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetFunc returns the method "FontFacePalette.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFacePalette) GetFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.FontFacePaletteGetFunc(
			this.Ref(),
		),
	)
}

type FontFacePalettes struct {
	ref js.Ref
}

func (this FontFacePalettes) Once() FontFacePalettes {
	this.Ref().Once()
	return this
}

func (this FontFacePalettes) Ref() js.Ref {
	return this.ref
}

func (this FontFacePalettes) FromRef(ref js.Ref) FontFacePalettes {
	this.ref = ref
	return this
}

func (this FontFacePalettes) Free() {
	this.Ref().Free()
}

// Length returns the value of property "FontFacePalettes.length".
//
// The returned bool will be false if there is no such property.
func (this FontFacePalettes) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetFontFacePalettesLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "FontFacePalettes.".
//
// The returned bool will be false if there is no such method.
func (this FontFacePalettes) Get(index uint32) (FontFacePalette, bool) {
	var _ok bool
	_ret := bindings.CallFontFacePalettesGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return FontFacePalette{}.FromRef(_ret), _ok
}

// GetFunc returns the method "FontFacePalettes.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFacePalettes) GetFunc() (fn js.Func[func(index uint32) FontFacePalette]) {
	return fn.FromRef(
		bindings.FontFacePalettesGetFunc(
			this.Ref(),
		),
	)
}

func NewFontFace(family js.String, source OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView, descriptors FontFaceDescriptors) FontFace {
	return FontFace{}.FromRef(
		bindings.NewFontFaceByFontFace(
			family.Ref(),
			source.Ref(),
			js.Pointer(&descriptors)),
	)
}

func NewFontFaceByFontFace1(family js.String, source OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FontFace {
	return FontFace{}.FromRef(
		bindings.NewFontFaceByFontFace1(
			family.Ref(),
			source.Ref()),
	)
}

type FontFace struct {
	ref js.Ref
}

func (this FontFace) Once() FontFace {
	this.Ref().Once()
	return this
}

func (this FontFace) Ref() js.Ref {
	return this.ref
}

func (this FontFace) FromRef(ref js.Ref) FontFace {
	this.ref = ref
	return this
}

func (this FontFace) Free() {
	this.Ref().Free()
}

// Family returns the value of property "FontFace.family".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Family() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceFamily(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Family sets the value of property "FontFace.family" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetFamily(val js.String) bool {
	return js.True == bindings.SetFontFaceFamily(
		this.Ref(),
		val.Ref(),
	)
}

// Style returns the value of property "FontFace.style".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Style() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Style sets the value of property "FontFace.style" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetStyle(val js.String) bool {
	return js.True == bindings.SetFontFaceStyle(
		this.Ref(),
		val.Ref(),
	)
}

// Weight returns the value of property "FontFace.weight".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Weight() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceWeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Weight sets the value of property "FontFace.weight" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetWeight(val js.String) bool {
	return js.True == bindings.SetFontFaceWeight(
		this.Ref(),
		val.Ref(),
	)
}

// Stretch returns the value of property "FontFace.stretch".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Stretch() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceStretch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Stretch sets the value of property "FontFace.stretch" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetStretch(val js.String) bool {
	return js.True == bindings.SetFontFaceStretch(
		this.Ref(),
		val.Ref(),
	)
}

// UnicodeRange returns the value of property "FontFace.unicodeRange".
//
// The returned bool will be false if there is no such property.
func (this FontFace) UnicodeRange() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceUnicodeRange(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UnicodeRange sets the value of property "FontFace.unicodeRange" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetUnicodeRange(val js.String) bool {
	return js.True == bindings.SetFontFaceUnicodeRange(
		this.Ref(),
		val.Ref(),
	)
}

// Variant returns the value of property "FontFace.variant".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Variant() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariant(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Variant sets the value of property "FontFace.variant" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetVariant(val js.String) bool {
	return js.True == bindings.SetFontFaceVariant(
		this.Ref(),
		val.Ref(),
	)
}

// FeatureSettings returns the value of property "FontFace.featureSettings".
//
// The returned bool will be false if there is no such property.
func (this FontFace) FeatureSettings() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceFeatureSettings(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FeatureSettings sets the value of property "FontFace.featureSettings" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetFeatureSettings(val js.String) bool {
	return js.True == bindings.SetFontFaceFeatureSettings(
		this.Ref(),
		val.Ref(),
	)
}

// VariationSettings returns the value of property "FontFace.variationSettings".
//
// The returned bool will be false if there is no such property.
func (this FontFace) VariationSettings() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationSettings(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VariationSettings sets the value of property "FontFace.variationSettings" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetVariationSettings(val js.String) bool {
	return js.True == bindings.SetFontFaceVariationSettings(
		this.Ref(),
		val.Ref(),
	)
}

// Display returns the value of property "FontFace.display".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Display() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceDisplay(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Display sets the value of property "FontFace.display" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetDisplay(val js.String) bool {
	return js.True == bindings.SetFontFaceDisplay(
		this.Ref(),
		val.Ref(),
	)
}

// AscentOverride returns the value of property "FontFace.ascentOverride".
//
// The returned bool will be false if there is no such property.
func (this FontFace) AscentOverride() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceAscentOverride(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AscentOverride sets the value of property "FontFace.ascentOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetAscentOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceAscentOverride(
		this.Ref(),
		val.Ref(),
	)
}

// DescentOverride returns the value of property "FontFace.descentOverride".
//
// The returned bool will be false if there is no such property.
func (this FontFace) DescentOverride() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceDescentOverride(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DescentOverride sets the value of property "FontFace.descentOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetDescentOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceDescentOverride(
		this.Ref(),
		val.Ref(),
	)
}

// LineGapOverride returns the value of property "FontFace.lineGapOverride".
//
// The returned bool will be false if there is no such property.
func (this FontFace) LineGapOverride() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceLineGapOverride(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LineGapOverride sets the value of property "FontFace.lineGapOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetLineGapOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceLineGapOverride(
		this.Ref(),
		val.Ref(),
	)
}

// Status returns the value of property "FontFace.status".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Status() (FontFaceLoadStatus, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceLoadStatus(_ret), _ok
}

// Loaded returns the value of property "FontFace.loaded".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Loaded() (js.Promise[FontFace], bool) {
	var _ok bool
	_ret := bindings.GetFontFaceLoaded(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[FontFace]{}.FromRef(_ret), _ok
}

// Features returns the value of property "FontFace.features".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Features() (FontFaceFeatures, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceFeatures(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceFeatures{}.FromRef(_ret), _ok
}

// Variations returns the value of property "FontFace.variations".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Variations() (FontFaceVariations, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariations(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceVariations{}.FromRef(_ret), _ok
}

// Palettes returns the value of property "FontFace.palettes".
//
// The returned bool will be false if there is no such property.
func (this FontFace) Palettes() (FontFacePalettes, bool) {
	var _ok bool
	_ret := bindings.GetFontFacePalettes(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFacePalettes{}.FromRef(_ret), _ok
}

// Load calls the method "FontFace.load".
//
// The returned bool will be false if there is no such method.
func (this FontFace) Load() (js.Promise[FontFace], bool) {
	var _ok bool
	_ret := bindings.CallFontFaceLoad(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FontFace]{}.FromRef(_ret), _ok
}

// LoadFunc returns the method "FontFace.load".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFace) LoadFunc() (fn js.Func[func() js.Promise[FontFace]]) {
	return fn.FromRef(
		bindings.FontFaceLoadFunc(
			this.Ref(),
		),
	)
}

type FontFaceSetLoadStatus uint32

const (
	_ FontFaceSetLoadStatus = iota

	FontFaceSetLoadStatus_LOADING
	FontFaceSetLoadStatus_LOADED
)

func (FontFaceSetLoadStatus) FromRef(str js.Ref) FontFaceSetLoadStatus {
	return FontFaceSetLoadStatus(bindings.ConstOfFontFaceSetLoadStatus(str))
}

func (x FontFaceSetLoadStatus) String() (string, bool) {
	switch x {
	case FontFaceSetLoadStatus_LOADING:
		return "loading", true
	case FontFaceSetLoadStatus_LOADED:
		return "loaded", true
	default:
		return "", false
	}
}

func NewFontFaceSet(initialFaces js.Array[FontFace]) FontFaceSet {
	return FontFaceSet{}.FromRef(
		bindings.NewFontFaceSetByFontFaceSet(
			initialFaces.Ref()),
	)
}

type FontFaceSet struct {
	EventTarget
}

func (this FontFaceSet) Once() FontFaceSet {
	this.Ref().Once()
	return this
}

func (this FontFaceSet) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this FontFaceSet) FromRef(ref js.Ref) FontFaceSet {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this FontFaceSet) Free() {
	this.Ref().Free()
}

// Ready returns the value of property "FontFaceSet.ready".
//
// The returned bool will be false if there is no such property.
func (this FontFaceSet) Ready() (js.Promise[FontFaceSet], bool) {
	var _ok bool
	_ret := bindings.GetFontFaceSetReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[FontFaceSet]{}.FromRef(_ret), _ok
}

// Status returns the value of property "FontFaceSet.status".
//
// The returned bool will be false if there is no such property.
func (this FontFaceSet) Status() (FontFaceSetLoadStatus, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceSetStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceSetLoadStatus(_ret), _ok
}

// Add calls the method "FontFaceSet.add".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Add(font FontFace) (FontFaceSet, bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetAdd(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
	)

	return FontFaceSet{}.FromRef(_ret), _ok
}

// AddFunc returns the method "FontFaceSet.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) AddFunc() (fn js.Func[func(font FontFace) FontFaceSet]) {
	return fn.FromRef(
		bindings.FontFaceSetAddFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "FontFaceSet.delete".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Delete(font FontFace) (bool, bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetDelete(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
	)

	return _ret == js.True, _ok
}

// DeleteFunc returns the method "FontFaceSet.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) DeleteFunc() (fn js.Func[func(font FontFace) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetDeleteFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "FontFaceSet.clear".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "FontFaceSet.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FontFaceSetClearFunc(
			this.Ref(),
		),
	)
}

// Load calls the method "FontFaceSet.load".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Load(font js.String, text js.String) (js.Promise[js.Array[FontFace]], bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetLoad(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
		text.Ref(),
	)

	return js.Promise[js.Array[FontFace]]{}.FromRef(_ret), _ok
}

// LoadFunc returns the method "FontFaceSet.load".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) LoadFunc() (fn js.Func[func(font js.String, text js.String) js.Promise[js.Array[FontFace]]]) {
	return fn.FromRef(
		bindings.FontFaceSetLoadFunc(
			this.Ref(),
		),
	)
}

// Load1 calls the method "FontFaceSet.load".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Load1(font js.String) (js.Promise[js.Array[FontFace]], bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetLoad1(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
	)

	return js.Promise[js.Array[FontFace]]{}.FromRef(_ret), _ok
}

// Load1Func returns the method "FontFaceSet.load".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) Load1Func() (fn js.Func[func(font js.String) js.Promise[js.Array[FontFace]]]) {
	return fn.FromRef(
		bindings.FontFaceSetLoad1Func(
			this.Ref(),
		),
	)
}

// Check calls the method "FontFaceSet.check".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Check(font js.String, text js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetCheck(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
		text.Ref(),
	)

	return _ret == js.True, _ok
}

// CheckFunc returns the method "FontFaceSet.check".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) CheckFunc() (fn js.Func[func(font js.String, text js.String) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetCheckFunc(
			this.Ref(),
		),
	)
}

// Check1 calls the method "FontFaceSet.check".
//
// The returned bool will be false if there is no such method.
func (this FontFaceSet) Check1(font js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallFontFaceSetCheck1(
		this.Ref(), js.Pointer(&_ok),
		font.Ref(),
	)

	return _ret == js.True, _ok
}

// Check1Func returns the method "FontFaceSet.check".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontFaceSet) Check1Func() (fn js.Func[func(font js.String) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetCheck1Func(
			this.Ref(),
		),
	)
}

type Document struct {
	Node
}

func (this Document) Once() Document {
	this.Ref().Once()
	return this
}

func (this Document) Ref() js.Ref {
	return this.Node.Ref()
}

func (this Document) FromRef(ref js.Ref) Document {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this Document) Free() {
	this.Ref().Free()
}

// Implementation returns the value of property "Document.implementation".
//
// The returned bool will be false if there is no such property.
func (this Document) Implementation() (DOMImplementation, bool) {
	var _ok bool
	_ret := bindings.GetDocumentImplementation(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMImplementation{}.FromRef(_ret), _ok
}

// URL returns the value of property "Document.URL".
//
// The returned bool will be false if there is no such property.
func (this Document) URL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DocumentURI returns the value of property "Document.documentURI".
//
// The returned bool will be false if there is no such property.
func (this Document) DocumentURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDocumentURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CompatMode returns the value of property "Document.compatMode".
//
// The returned bool will be false if there is no such property.
func (this Document) CompatMode() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentCompatMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CharacterSet returns the value of property "Document.characterSet".
//
// The returned bool will be false if there is no such property.
func (this Document) CharacterSet() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentCharacterSet(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Charset returns the value of property "Document.charset".
//
// The returned bool will be false if there is no such property.
func (this Document) Charset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentCharset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InputEncoding returns the value of property "Document.inputEncoding".
//
// The returned bool will be false if there is no such property.
func (this Document) InputEncoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentInputEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContentType returns the value of property "Document.contentType".
//
// The returned bool will be false if there is no such property.
func (this Document) ContentType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentContentType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Doctype returns the value of property "Document.doctype".
//
// The returned bool will be false if there is no such property.
func (this Document) Doctype() (DocumentType, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDoctype(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentType{}.FromRef(_ret), _ok
}

// DocumentElement returns the value of property "Document.documentElement".
//
// The returned bool will be false if there is no such property.
func (this Document) DocumentElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDocumentElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// FragmentDirective returns the value of property "Document.fragmentDirective".
//
// The returned bool will be false if there is no such property.
func (this Document) FragmentDirective() (FragmentDirective, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFragmentDirective(
		this.Ref(), js.Pointer(&_ok),
	)
	return FragmentDirective{}.FromRef(_ret), _ok
}

// PermissionsPolicy returns the value of property "Document.permissionsPolicy".
//
// The returned bool will be false if there is no such property.
func (this Document) PermissionsPolicy() (PermissionsPolicy, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPermissionsPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return PermissionsPolicy{}.FromRef(_ret), _ok
}

// WasDiscarded returns the value of property "Document.wasDiscarded".
//
// The returned bool will be false if there is no such property.
func (this Document) WasDiscarded() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentWasDiscarded(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FullscreenEnabled returns the value of property "Document.fullscreenEnabled".
//
// The returned bool will be false if there is no such property.
func (this Document) FullscreenEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFullscreenEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Fullscreen returns the value of property "Document.fullscreen".
//
// The returned bool will be false if there is no such property.
func (this Document) Fullscreen() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFullscreen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Timeline returns the value of property "Document.timeline".
//
// The returned bool will be false if there is no such property.
func (this Document) Timeline() (DocumentTimeline, bool) {
	var _ok bool
	_ret := bindings.GetDocumentTimeline(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentTimeline{}.FromRef(_ret), _ok
}

// PictureInPictureEnabled returns the value of property "Document.pictureInPictureEnabled".
//
// The returned bool will be false if there is no such property.
func (this Document) PictureInPictureEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPictureInPictureEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NamedFlows returns the value of property "Document.namedFlows".
//
// The returned bool will be false if there is no such property.
func (this Document) NamedFlows() (NamedFlowMap, bool) {
	var _ok bool
	_ret := bindings.GetDocumentNamedFlows(
		this.Ref(), js.Pointer(&_ok),
	)
	return NamedFlowMap{}.FromRef(_ret), _ok
}

// ScrollingElement returns the value of property "Document.scrollingElement".
//
// The returned bool will be false if there is no such property.
func (this Document) ScrollingElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentScrollingElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// RootElement returns the value of property "Document.rootElement".
//
// The returned bool will be false if there is no such property.
func (this Document) RootElement() (SVGSVGElement, bool) {
	var _ok bool
	_ret := bindings.GetDocumentRootElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return SVGSVGElement{}.FromRef(_ret), _ok
}

// Prerendering returns the value of property "Document.prerendering".
//
// The returned bool will be false if there is no such property.
func (this Document) Prerendering() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPrerendering(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FgColor returns the value of property "Document.fgColor".
//
// The returned bool will be false if there is no such property.
func (this Document) FgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FgColor sets the value of property "Document.fgColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetFgColor(val js.String) bool {
	return js.True == bindings.SetDocumentFgColor(
		this.Ref(),
		val.Ref(),
	)
}

// LinkColor returns the value of property "Document.linkColor".
//
// The returned bool will be false if there is no such property.
func (this Document) LinkColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentLinkColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LinkColor sets the value of property "Document.linkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetLinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentLinkColor(
		this.Ref(),
		val.Ref(),
	)
}

// VlinkColor returns the value of property "Document.vlinkColor".
//
// The returned bool will be false if there is no such property.
func (this Document) VlinkColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentVlinkColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VlinkColor sets the value of property "Document.vlinkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetVlinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentVlinkColor(
		this.Ref(),
		val.Ref(),
	)
}

// AlinkColor returns the value of property "Document.alinkColor".
//
// The returned bool will be false if there is no such property.
func (this Document) AlinkColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentAlinkColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AlinkColor sets the value of property "Document.alinkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetAlinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentAlinkColor(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "Document.bgColor".
//
// The returned bool will be false if there is no such property.
func (this Document) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "Document.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetBgColor(val js.String) bool {
	return js.True == bindings.SetDocumentBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// Anchors returns the value of property "Document.anchors".
//
// The returned bool will be false if there is no such property.
func (this Document) Anchors() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentAnchors(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Applets returns the value of property "Document.applets".
//
// The returned bool will be false if there is no such property.
func (this Document) Applets() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentApplets(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// All returns the value of property "Document.all".
//
// The returned bool will be false if there is no such property.
func (this Document) All() (HTMLAllCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentAll(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLAllCollection{}.FromRef(_ret), _ok
}

// Location returns the value of property "Document.location".
//
// The returned bool will be false if there is no such property.
func (this Document) Location() (Location, bool) {
	var _ok bool
	_ret := bindings.GetDocumentLocation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Location{}.FromRef(_ret), _ok
}

// Domain returns the value of property "Document.domain".
//
// The returned bool will be false if there is no such property.
func (this Document) Domain() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDomain(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Domain sets the value of property "Document.domain" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDomain(val js.String) bool {
	return js.True == bindings.SetDocumentDomain(
		this.Ref(),
		val.Ref(),
	)
}

// Referrer returns the value of property "Document.referrer".
//
// The returned bool will be false if there is no such property.
func (this Document) Referrer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentReferrer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Cookie returns the value of property "Document.cookie".
//
// The returned bool will be false if there is no such property.
func (this Document) Cookie() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentCookie(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Cookie sets the value of property "Document.cookie" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetCookie(val js.String) bool {
	return js.True == bindings.SetDocumentCookie(
		this.Ref(),
		val.Ref(),
	)
}

// LastModified returns the value of property "Document.lastModified".
//
// The returned bool will be false if there is no such property.
func (this Document) LastModified() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentLastModified(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReadyState returns the value of property "Document.readyState".
//
// The returned bool will be false if there is no such property.
func (this Document) ReadyState() (DocumentReadyState, bool) {
	var _ok bool
	_ret := bindings.GetDocumentReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentReadyState(_ret), _ok
}

// Title returns the value of property "Document.title".
//
// The returned bool will be false if there is no such property.
func (this Document) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Title sets the value of property "Document.title" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetTitle(val js.String) bool {
	return js.True == bindings.SetDocumentTitle(
		this.Ref(),
		val.Ref(),
	)
}

// Dir returns the value of property "Document.dir".
//
// The returned bool will be false if there is no such property.
func (this Document) Dir() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDir(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Dir sets the value of property "Document.dir" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDir(val js.String) bool {
	return js.True == bindings.SetDocumentDir(
		this.Ref(),
		val.Ref(),
	)
}

// Body returns the value of property "Document.body".
//
// The returned bool will be false if there is no such property.
func (this Document) Body() (HTMLElement, bool) {
	var _ok bool
	_ret := bindings.GetDocumentBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLElement{}.FromRef(_ret), _ok
}

// Body sets the value of property "Document.body" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetBody(val HTMLElement) bool {
	return js.True == bindings.SetDocumentBody(
		this.Ref(),
		val.Ref(),
	)
}

// Head returns the value of property "Document.head".
//
// The returned bool will be false if there is no such property.
func (this Document) Head() (HTMLHeadElement, bool) {
	var _ok bool
	_ret := bindings.GetDocumentHead(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLHeadElement{}.FromRef(_ret), _ok
}

// Images returns the value of property "Document.images".
//
// The returned bool will be false if there is no such property.
func (this Document) Images() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentImages(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Embeds returns the value of property "Document.embeds".
//
// The returned bool will be false if there is no such property.
func (this Document) Embeds() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentEmbeds(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Plugins returns the value of property "Document.plugins".
//
// The returned bool will be false if there is no such property.
func (this Document) Plugins() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPlugins(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Links returns the value of property "Document.links".
//
// The returned bool will be false if there is no such property.
func (this Document) Links() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentLinks(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Forms returns the value of property "Document.forms".
//
// The returned bool will be false if there is no such property.
func (this Document) Forms() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentForms(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Scripts returns the value of property "Document.scripts".
//
// The returned bool will be false if there is no such property.
func (this Document) Scripts() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentScripts(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// CurrentScript returns the value of property "Document.currentScript".
//
// The returned bool will be false if there is no such property.
func (this Document) CurrentScript() (HTMLOrSVGScriptElement, bool) {
	var _ok bool
	_ret := bindings.GetDocumentCurrentScript(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLOrSVGScriptElement{}.FromRef(_ret), _ok
}

// DefaultView returns the value of property "Document.defaultView".
//
// The returned bool will be false if there is no such property.
func (this Document) DefaultView() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDefaultView(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// DesignMode returns the value of property "Document.designMode".
//
// The returned bool will be false if there is no such property.
func (this Document) DesignMode() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDocumentDesignMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DesignMode sets the value of property "Document.designMode" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDesignMode(val js.String) bool {
	return js.True == bindings.SetDocumentDesignMode(
		this.Ref(),
		val.Ref(),
	)
}

// Hidden returns the value of property "Document.hidden".
//
// The returned bool will be false if there is no such property.
func (this Document) Hidden() (bool, bool) {
	var _ok bool
	_ret := bindings.GetDocumentHidden(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// VisibilityState returns the value of property "Document.visibilityState".
//
// The returned bool will be false if there is no such property.
func (this Document) VisibilityState() (DocumentVisibilityState, bool) {
	var _ok bool
	_ret := bindings.GetDocumentVisibilityState(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentVisibilityState(_ret), _ok
}

// FullscreenElement returns the value of property "Document.fullscreenElement".
//
// The returned bool will be false if there is no such property.
func (this Document) FullscreenElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFullscreenElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// StyleSheets returns the value of property "Document.styleSheets".
//
// The returned bool will be false if there is no such property.
func (this Document) StyleSheets() (StyleSheetList, bool) {
	var _ok bool
	_ret := bindings.GetDocumentStyleSheets(
		this.Ref(), js.Pointer(&_ok),
	)
	return StyleSheetList{}.FromRef(_ret), _ok
}

// AdoptedStyleSheets returns the value of property "Document.adoptedStyleSheets".
//
// The returned bool will be false if there is no such property.
func (this Document) AdoptedStyleSheets() (js.ObservableArray[CSSStyleSheet], bool) {
	var _ok bool
	_ret := bindings.GetDocumentAdoptedStyleSheets(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ObservableArray[CSSStyleSheet]{}.FromRef(_ret), _ok
}

// AdoptedStyleSheets sets the value of property "Document.adoptedStyleSheets" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetAdoptedStyleSheets(val js.ObservableArray[CSSStyleSheet]) bool {
	return js.True == bindings.SetDocumentAdoptedStyleSheets(
		this.Ref(),
		val.Ref(),
	)
}

// PictureInPictureElement returns the value of property "Document.pictureInPictureElement".
//
// The returned bool will be false if there is no such property.
func (this Document) PictureInPictureElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPictureInPictureElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ActiveElement returns the value of property "Document.activeElement".
//
// The returned bool will be false if there is no such property.
func (this Document) ActiveElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentActiveElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// PointerLockElement returns the value of property "Document.pointerLockElement".
//
// The returned bool will be false if there is no such property.
func (this Document) PointerLockElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentPointerLockElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Children returns the value of property "Document.children".
//
// The returned bool will be false if there is no such property.
func (this Document) Children() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetDocumentChildren(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// FirstElementChild returns the value of property "Document.firstElementChild".
//
// The returned bool will be false if there is no such property.
func (this Document) FirstElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFirstElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// LastElementChild returns the value of property "Document.lastElementChild".
//
// The returned bool will be false if there is no such property.
func (this Document) LastElementChild() (Element, bool) {
	var _ok bool
	_ret := bindings.GetDocumentLastElementChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ChildElementCount returns the value of property "Document.childElementCount".
//
// The returned bool will be false if there is no such property.
func (this Document) ChildElementCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetDocumentChildElementCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Fonts returns the value of property "Document.fonts".
//
// The returned bool will be false if there is no such property.
func (this Document) Fonts() (FontFaceSet, bool) {
	var _ok bool
	_ret := bindings.GetDocumentFonts(
		this.Ref(), js.Pointer(&_ok),
	)
	return FontFaceSet{}.FromRef(_ret), _ok
}

// GetElementsByTagName calls the method "Document.getElementsByTagName".
//
// The returned bool will be false if there is no such method.
func (this Document) GetElementsByTagName(qualifiedName js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetElementsByTagName(
		this.Ref(), js.Pointer(&_ok),
		qualifiedName.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByTagNameFunc returns the method "Document.getElementsByTagName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetElementsByTagNameFunc() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByTagNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagNameNS calls the method "Document.getElementsByTagNameNS".
//
// The returned bool will be false if there is no such method.
func (this Document) GetElementsByTagNameNS(namespace js.String, localName js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetElementsByTagNameNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		localName.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByTagNameNSFunc returns the method "Document.getElementsByTagNameNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetElementsByTagNameNSFunc() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByTagNameNSFunc(
			this.Ref(),
		),
	)
}

// GetElementsByClassName calls the method "Document.getElementsByClassName".
//
// The returned bool will be false if there is no such method.
func (this Document) GetElementsByClassName(classNames js.String) (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetElementsByClassName(
		this.Ref(), js.Pointer(&_ok),
		classNames.Ref(),
	)

	return HTMLCollection{}.FromRef(_ret), _ok
}

// GetElementsByClassNameFunc returns the method "Document.getElementsByClassName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetElementsByClassNameFunc() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByClassNameFunc(
			this.Ref(),
		),
	)
}

// CreateElement calls the method "Document.createElement".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateElement(localName js.String, options OneOf_String_ElementCreationOptions) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateElement(
		this.Ref(), js.Pointer(&_ok),
		localName.Ref(),
		options.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// CreateElementFunc returns the method "Document.createElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateElementFunc() (fn js.Func[func(localName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementFunc(
			this.Ref(),
		),
	)
}

// CreateElement1 calls the method "Document.createElement".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateElement1(localName js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateElement1(
		this.Ref(), js.Pointer(&_ok),
		localName.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// CreateElement1Func returns the method "Document.createElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateElement1Func() (fn js.Func[func(localName js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElement1Func(
			this.Ref(),
		),
	)
}

// CreateElementNS calls the method "Document.createElementNS".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateElementNS(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateElementNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
		options.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// CreateElementNSFunc returns the method "Document.createElementNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateElementNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementNSFunc(
			this.Ref(),
		),
	)
}

// CreateElementNS1 calls the method "Document.createElementNS".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateElementNS1(namespace js.String, qualifiedName js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateElementNS1(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// CreateElementNS1Func returns the method "Document.createElementNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateElementNS1Func() (fn js.Func[func(namespace js.String, qualifiedName js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementNS1Func(
			this.Ref(),
		),
	)
}

// CreateDocumentFragment calls the method "Document.createDocumentFragment".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateDocumentFragment() (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateDocumentFragment(
		this.Ref(), js.Pointer(&_ok),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// CreateDocumentFragmentFunc returns the method "Document.createDocumentFragment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateDocumentFragmentFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.DocumentCreateDocumentFragmentFunc(
			this.Ref(),
		),
	)
}

// CreateTextNode calls the method "Document.createTextNode".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateTextNode(data js.String) (Text, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateTextNode(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return Text{}.FromRef(_ret), _ok
}

// CreateTextNodeFunc returns the method "Document.createTextNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateTextNodeFunc() (fn js.Func[func(data js.String) Text]) {
	return fn.FromRef(
		bindings.DocumentCreateTextNodeFunc(
			this.Ref(),
		),
	)
}

// CreateCDATASection calls the method "Document.createCDATASection".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateCDATASection(data js.String) (CDATASection, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateCDATASection(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return CDATASection{}.FromRef(_ret), _ok
}

// CreateCDATASectionFunc returns the method "Document.createCDATASection".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateCDATASectionFunc() (fn js.Func[func(data js.String) CDATASection]) {
	return fn.FromRef(
		bindings.DocumentCreateCDATASectionFunc(
			this.Ref(),
		),
	)
}

// CreateComment calls the method "Document.createComment".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateComment(data js.String) (Comment, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateComment(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return Comment{}.FromRef(_ret), _ok
}

// CreateCommentFunc returns the method "Document.createComment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateCommentFunc() (fn js.Func[func(data js.String) Comment]) {
	return fn.FromRef(
		bindings.DocumentCreateCommentFunc(
			this.Ref(),
		),
	)
}

// CreateProcessingInstruction calls the method "Document.createProcessingInstruction".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateProcessingInstruction(target js.String, data js.String) (ProcessingInstruction, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateProcessingInstruction(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
		data.Ref(),
	)

	return ProcessingInstruction{}.FromRef(_ret), _ok
}

// CreateProcessingInstructionFunc returns the method "Document.createProcessingInstruction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateProcessingInstructionFunc() (fn js.Func[func(target js.String, data js.String) ProcessingInstruction]) {
	return fn.FromRef(
		bindings.DocumentCreateProcessingInstructionFunc(
			this.Ref(),
		),
	)
}

// ImportNode calls the method "Document.importNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ImportNode(node Node, deep bool) (Node, bool) {
	var _ok bool
	_ret := bindings.CallDocumentImportNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		js.Bool(bool(deep)),
	)

	return Node{}.FromRef(_ret), _ok
}

// ImportNodeFunc returns the method "Document.importNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ImportNodeFunc() (fn js.Func[func(node Node, deep bool) Node]) {
	return fn.FromRef(
		bindings.DocumentImportNodeFunc(
			this.Ref(),
		),
	)
}

// ImportNode1 calls the method "Document.importNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ImportNode1(node Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallDocumentImportNode1(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// ImportNode1Func returns the method "Document.importNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ImportNode1Func() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.DocumentImportNode1Func(
			this.Ref(),
		),
	)
}

// AdoptNode calls the method "Document.adoptNode".
//
// The returned bool will be false if there is no such method.
func (this Document) AdoptNode(node Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallDocumentAdoptNode(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// AdoptNodeFunc returns the method "Document.adoptNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) AdoptNodeFunc() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.DocumentAdoptNodeFunc(
			this.Ref(),
		),
	)
}

// CreateAttribute calls the method "Document.createAttribute".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateAttribute(localName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateAttribute(
		this.Ref(), js.Pointer(&_ok),
		localName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// CreateAttributeFunc returns the method "Document.createAttribute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateAttributeFunc() (fn js.Func[func(localName js.String) Attr]) {
	return fn.FromRef(
		bindings.DocumentCreateAttributeFunc(
			this.Ref(),
		),
	)
}

// CreateAttributeNS calls the method "Document.createAttributeNS".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateAttributeNS(namespace js.String, qualifiedName js.String) (Attr, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateAttributeNS(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return Attr{}.FromRef(_ret), _ok
}

// CreateAttributeNSFunc returns the method "Document.createAttributeNS".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateAttributeNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.DocumentCreateAttributeNSFunc(
			this.Ref(),
		),
	)
}

// CreateEvent calls the method "Document.createEvent".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateEvent(iface js.String) (Event, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateEvent(
		this.Ref(), js.Pointer(&_ok),
		iface.Ref(),
	)

	return Event{}.FromRef(_ret), _ok
}

// CreateEventFunc returns the method "Document.createEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateEventFunc() (fn js.Func[func(iface js.String) Event]) {
	return fn.FromRef(
		bindings.DocumentCreateEventFunc(
			this.Ref(),
		),
	)
}

// CreateRange calls the method "Document.createRange".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateRange() (Range, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateRange(
		this.Ref(), js.Pointer(&_ok),
	)

	return Range{}.FromRef(_ret), _ok
}

// CreateRangeFunc returns the method "Document.createRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateRangeFunc() (fn js.Func[func() Range]) {
	return fn.FromRef(
		bindings.DocumentCreateRangeFunc(
			this.Ref(),
		),
	)
}

// CreateNodeIterator calls the method "Document.createNodeIterator".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateNodeIterator(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (NodeIterator, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateNodeIterator(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return NodeIterator{}.FromRef(_ret), _ok
}

// CreateNodeIteratorFunc returns the method "Document.createNodeIterator".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateNodeIteratorFunc() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIteratorFunc(
			this.Ref(),
		),
	)
}

// CreateNodeIterator1 calls the method "Document.createNodeIterator".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateNodeIterator1(root Node, whatToShow uint32) (NodeIterator, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateNodeIterator1(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
		uint32(whatToShow),
	)

	return NodeIterator{}.FromRef(_ret), _ok
}

// CreateNodeIterator1Func returns the method "Document.createNodeIterator".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateNodeIterator1Func() (fn js.Func[func(root Node, whatToShow uint32) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIterator1Func(
			this.Ref(),
		),
	)
}

// CreateNodeIterator2 calls the method "Document.createNodeIterator".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateNodeIterator2(root Node) (NodeIterator, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateNodeIterator2(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
	)

	return NodeIterator{}.FromRef(_ret), _ok
}

// CreateNodeIterator2Func returns the method "Document.createNodeIterator".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateNodeIterator2Func() (fn js.Func[func(root Node) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIterator2Func(
			this.Ref(),
		),
	)
}

// CreateTreeWalker calls the method "Document.createTreeWalker".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateTreeWalker(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (TreeWalker, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateTreeWalker(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return TreeWalker{}.FromRef(_ret), _ok
}

// CreateTreeWalkerFunc returns the method "Document.createTreeWalker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateTreeWalkerFunc() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalkerFunc(
			this.Ref(),
		),
	)
}

// CreateTreeWalker1 calls the method "Document.createTreeWalker".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateTreeWalker1(root Node, whatToShow uint32) (TreeWalker, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateTreeWalker1(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
		uint32(whatToShow),
	)

	return TreeWalker{}.FromRef(_ret), _ok
}

// CreateTreeWalker1Func returns the method "Document.createTreeWalker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateTreeWalker1Func() (fn js.Func[func(root Node, whatToShow uint32) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalker1Func(
			this.Ref(),
		),
	)
}

// CreateTreeWalker2 calls the method "Document.createTreeWalker".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateTreeWalker2(root Node) (TreeWalker, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateTreeWalker2(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
	)

	return TreeWalker{}.FromRef(_ret), _ok
}

// CreateTreeWalker2Func returns the method "Document.createTreeWalker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateTreeWalker2Func() (fn js.Func[func(root Node) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalker2Func(
			this.Ref(),
		),
	)
}

// StartViewTransition calls the method "Document.startViewTransition".
//
// The returned bool will be false if there is no such method.
func (this Document) StartViewTransition(updateCallback js.Func[func() js.Promise[js.Any]]) (ViewTransition, bool) {
	var _ok bool
	_ret := bindings.CallDocumentStartViewTransition(
		this.Ref(), js.Pointer(&_ok),
		updateCallback.Ref(),
	)

	return ViewTransition{}.FromRef(_ret), _ok
}

// StartViewTransitionFunc returns the method "Document.startViewTransition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) StartViewTransitionFunc() (fn js.Func[func(updateCallback js.Func[func() js.Promise[js.Any]]) ViewTransition]) {
	return fn.FromRef(
		bindings.DocumentStartViewTransitionFunc(
			this.Ref(),
		),
	)
}

// StartViewTransition1 calls the method "Document.startViewTransition".
//
// The returned bool will be false if there is no such method.
func (this Document) StartViewTransition1() (ViewTransition, bool) {
	var _ok bool
	_ret := bindings.CallDocumentStartViewTransition1(
		this.Ref(), js.Pointer(&_ok),
	)

	return ViewTransition{}.FromRef(_ret), _ok
}

// StartViewTransition1Func returns the method "Document.startViewTransition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) StartViewTransition1Func() (fn js.Func[func() ViewTransition]) {
	return fn.FromRef(
		bindings.DocumentStartViewTransition1Func(
			this.Ref(),
		),
	)
}

// MeasureElement calls the method "Document.measureElement".
//
// The returned bool will be false if there is no such method.
func (this Document) MeasureElement(element Element) (FontMetrics, bool) {
	var _ok bool
	_ret := bindings.CallDocumentMeasureElement(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	return FontMetrics{}.FromRef(_ret), _ok
}

// MeasureElementFunc returns the method "Document.measureElement".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) MeasureElementFunc() (fn js.Func[func(element Element) FontMetrics]) {
	return fn.FromRef(
		bindings.DocumentMeasureElementFunc(
			this.Ref(),
		),
	)
}

// MeasureText calls the method "Document.measureText".
//
// The returned bool will be false if there is no such method.
func (this Document) MeasureText(text js.String, styleMap StylePropertyMapReadOnly) (FontMetrics, bool) {
	var _ok bool
	_ret := bindings.CallDocumentMeasureText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		styleMap.Ref(),
	)

	return FontMetrics{}.FromRef(_ret), _ok
}

// MeasureTextFunc returns the method "Document.measureText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) MeasureTextFunc() (fn js.Func[func(text js.String, styleMap StylePropertyMapReadOnly) FontMetrics]) {
	return fn.FromRef(
		bindings.DocumentMeasureTextFunc(
			this.Ref(),
		),
	)
}

// GetSelection calls the method "Document.getSelection".
//
// The returned bool will be false if there is no such method.
func (this Document) GetSelection() (Selection, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetSelection(
		this.Ref(), js.Pointer(&_ok),
	)

	return Selection{}.FromRef(_ret), _ok
}

// GetSelectionFunc returns the method "Document.getSelection".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetSelectionFunc() (fn js.Func[func() Selection]) {
	return fn.FromRef(
		bindings.DocumentGetSelectionFunc(
			this.Ref(),
		),
	)
}

// ExitPointerLock calls the method "Document.exitPointerLock".
//
// The returned bool will be false if there is no such method.
func (this Document) ExitPointerLock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentExitPointerLock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ExitPointerLockFunc returns the method "Document.exitPointerLock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExitPointerLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentExitPointerLockFunc(
			this.Ref(),
		),
	)
}

// ExitFullscreen calls the method "Document.exitFullscreen".
//
// The returned bool will be false if there is no such method.
func (this Document) ExitFullscreen() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallDocumentExitFullscreen(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ExitFullscreenFunc returns the method "Document.exitFullscreen".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExitFullscreenFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentExitFullscreenFunc(
			this.Ref(),
		),
	)
}

// HasStorageAccess calls the method "Document.hasStorageAccess".
//
// The returned bool will be false if there is no such method.
func (this Document) HasStorageAccess() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallDocumentHasStorageAccess(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// HasStorageAccessFunc returns the method "Document.hasStorageAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) HasStorageAccessFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasStorageAccessFunc(
			this.Ref(),
		),
	)
}

// RequestStorageAccess calls the method "Document.requestStorageAccess".
//
// The returned bool will be false if there is no such method.
func (this Document) RequestStorageAccess() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallDocumentRequestStorageAccess(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RequestStorageAccessFunc returns the method "Document.requestStorageAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) RequestStorageAccessFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentRequestStorageAccessFunc(
			this.Ref(),
		),
	)
}

// ExitPictureInPicture calls the method "Document.exitPictureInPicture".
//
// The returned bool will be false if there is no such method.
func (this Document) ExitPictureInPicture() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallDocumentExitPictureInPicture(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ExitPictureInPictureFunc returns the method "Document.exitPictureInPicture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExitPictureInPictureFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentExitPictureInPictureFunc(
			this.Ref(),
		),
	)
}

// HasPrivateTokens calls the method "Document.hasPrivateTokens".
//
// The returned bool will be false if there is no such method.
func (this Document) HasPrivateTokens(issuer js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallDocumentHasPrivateTokens(
		this.Ref(), js.Pointer(&_ok),
		issuer.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// HasPrivateTokensFunc returns the method "Document.hasPrivateTokens".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) HasPrivateTokensFunc() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasPrivateTokensFunc(
			this.Ref(),
		),
	)
}

// HasRedemptionRecord calls the method "Document.hasRedemptionRecord".
//
// The returned bool will be false if there is no such method.
func (this Document) HasRedemptionRecord(issuer js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallDocumentHasRedemptionRecord(
		this.Ref(), js.Pointer(&_ok),
		issuer.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// HasRedemptionRecordFunc returns the method "Document.hasRedemptionRecord".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) HasRedemptionRecordFunc() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasRedemptionRecordFunc(
			this.Ref(),
		),
	)
}

// ElementFromPoint calls the method "Document.elementFromPoint".
//
// The returned bool will be false if there is no such method.
func (this Document) ElementFromPoint(x float64, y float64) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentElementFromPoint(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return Element{}.FromRef(_ret), _ok
}

// ElementFromPointFunc returns the method "Document.elementFromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ElementFromPointFunc() (fn js.Func[func(x float64, y float64) Element]) {
	return fn.FromRef(
		bindings.DocumentElementFromPointFunc(
			this.Ref(),
		),
	)
}

// ElementsFromPoint calls the method "Document.elementsFromPoint".
//
// The returned bool will be false if there is no such method.
func (this Document) ElementsFromPoint(x float64, y float64) (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallDocumentElementsFromPoint(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// ElementsFromPointFunc returns the method "Document.elementsFromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ElementsFromPointFunc() (fn js.Func[func(x float64, y float64) js.Array[Element]]) {
	return fn.FromRef(
		bindings.DocumentElementsFromPointFunc(
			this.Ref(),
		),
	)
}

// CaretPositionFromPoint calls the method "Document.caretPositionFromPoint".
//
// The returned bool will be false if there is no such method.
func (this Document) CaretPositionFromPoint(x float64, y float64) (CaretPosition, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCaretPositionFromPoint(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return CaretPosition{}.FromRef(_ret), _ok
}

// CaretPositionFromPointFunc returns the method "Document.caretPositionFromPoint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CaretPositionFromPointFunc() (fn js.Func[func(x float64, y float64) CaretPosition]) {
	return fn.FromRef(
		bindings.DocumentCaretPositionFromPointFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "Document.clear".
//
// The returned bool will be false if there is no such method.
func (this Document) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentClear(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "Document.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentClearFunc(
			this.Ref(),
		),
	)
}

// CaptureEvents calls the method "Document.captureEvents".
//
// The returned bool will be false if there is no such method.
func (this Document) CaptureEvents() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCaptureEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CaptureEventsFunc returns the method "Document.captureEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CaptureEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentCaptureEventsFunc(
			this.Ref(),
		),
	)
}

// ReleaseEvents calls the method "Document.releaseEvents".
//
// The returned bool will be false if there is no such method.
func (this Document) ReleaseEvents() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentReleaseEvents(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReleaseEventsFunc returns the method "Document.releaseEvents".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ReleaseEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentReleaseEventsFunc(
			this.Ref(),
		),
	)
}

// RequestStorageAccessFor calls the method "Document.requestStorageAccessFor".
//
// The returned bool will be false if there is no such method.
func (this Document) RequestStorageAccessFor(requestedOrigin js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallDocumentRequestStorageAccessFor(
		this.Ref(), js.Pointer(&_ok),
		requestedOrigin.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RequestStorageAccessForFunc returns the method "Document.requestStorageAccessFor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) RequestStorageAccessForFunc() (fn js.Func[func(requestedOrigin js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentRequestStorageAccessForFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Document.".
//
// The returned bool will be false if there is no such method.
func (this Document) Get(name js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// GetFunc returns the method "Document.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.DocumentGetFunc(
			this.Ref(),
		),
	)
}

// GetElementsByName calls the method "Document.getElementsByName".
//
// The returned bool will be false if there is no such method.
func (this Document) GetElementsByName(elementName js.String) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetElementsByName(
		this.Ref(), js.Pointer(&_ok),
		elementName.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// GetElementsByNameFunc returns the method "Document.getElementsByName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetElementsByNameFunc() (fn js.Func[func(elementName js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByNameFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "Document.open".
//
// The returned bool will be false if there is no such method.
func (this Document) Open(unused1 js.String, unused2 js.String) (Document, bool) {
	var _ok bool
	_ret := bindings.CallDocumentOpen(
		this.Ref(), js.Pointer(&_ok),
		unused1.Ref(),
		unused2.Ref(),
	)

	return Document{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "Document.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) OpenFunc() (fn js.Func[func(unused1 js.String, unused2 js.String) Document]) {
	return fn.FromRef(
		bindings.DocumentOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "Document.open".
//
// The returned bool will be false if there is no such method.
func (this Document) Open1(unused1 js.String) (Document, bool) {
	var _ok bool
	_ret := bindings.CallDocumentOpen1(
		this.Ref(), js.Pointer(&_ok),
		unused1.Ref(),
	)

	return Document{}.FromRef(_ret), _ok
}

// Open1Func returns the method "Document.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Open1Func() (fn js.Func[func(unused1 js.String) Document]) {
	return fn.FromRef(
		bindings.DocumentOpen1Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "Document.open".
//
// The returned bool will be false if there is no such method.
func (this Document) Open2() (Document, bool) {
	var _ok bool
	_ret := bindings.CallDocumentOpen2(
		this.Ref(), js.Pointer(&_ok),
	)

	return Document{}.FromRef(_ret), _ok
}

// Open2Func returns the method "Document.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Open2Func() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.DocumentOpen2Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "Document.open".
//
// The returned bool will be false if there is no such method.
func (this Document) Open3(url js.String, name js.String, features js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallDocumentOpen3(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		name.Ref(),
		features.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// Open3Func returns the method "Document.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Open3Func() (fn js.Func[func(url js.String, name js.String, features js.String) js.Object]) {
	return fn.FromRef(
		bindings.DocumentOpen3Func(
			this.Ref(),
		),
	)
}

// Close calls the method "Document.close".
//
// The returned bool will be false if there is no such method.
func (this Document) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "Document.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentCloseFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "Document.write".
//
// The returned bool will be false if there is no such method.
func (this Document) Write(text ...js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentWrite(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteFunc returns the method "Document.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) WriteFunc() (fn js.Func[func(text ...js.String)]) {
	return fn.FromRef(
		bindings.DocumentWriteFunc(
			this.Ref(),
		),
	)
}

// Writeln calls the method "Document.writeln".
//
// The returned bool will be false if there is no such method.
func (this Document) Writeln(text ...js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentWriteln(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WritelnFunc returns the method "Document.writeln".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) WritelnFunc() (fn js.Func[func(text ...js.String)]) {
	return fn.FromRef(
		bindings.DocumentWritelnFunc(
			this.Ref(),
		),
	)
}

// HasFocus calls the method "Document.hasFocus".
//
// The returned bool will be false if there is no such method.
func (this Document) HasFocus() (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentHasFocus(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// HasFocusFunc returns the method "Document.hasFocus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) HasFocusFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.DocumentHasFocusFunc(
			this.Ref(),
		),
	)
}

// ExecCommand calls the method "Document.execCommand".
//
// The returned bool will be false if there is no such method.
func (this Document) ExecCommand(commandId js.String, showUI bool, value js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentExecCommand(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
		js.Bool(bool(showUI)),
		value.Ref(),
	)

	return _ret == js.True, _ok
}

// ExecCommandFunc returns the method "Document.execCommand".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExecCommandFunc() (fn js.Func[func(commandId js.String, showUI bool, value js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommandFunc(
			this.Ref(),
		),
	)
}

// ExecCommand1 calls the method "Document.execCommand".
//
// The returned bool will be false if there is no such method.
func (this Document) ExecCommand1(commandId js.String, showUI bool) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentExecCommand1(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
		js.Bool(bool(showUI)),
	)

	return _ret == js.True, _ok
}

// ExecCommand1Func returns the method "Document.execCommand".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExecCommand1Func() (fn js.Func[func(commandId js.String, showUI bool) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommand1Func(
			this.Ref(),
		),
	)
}

// ExecCommand2 calls the method "Document.execCommand".
//
// The returned bool will be false if there is no such method.
func (this Document) ExecCommand2(commandId js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentExecCommand2(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return _ret == js.True, _ok
}

// ExecCommand2Func returns the method "Document.execCommand".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ExecCommand2Func() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommand2Func(
			this.Ref(),
		),
	)
}

// QueryCommandEnabled calls the method "Document.queryCommandEnabled".
//
// The returned bool will be false if there is no such method.
func (this Document) QueryCommandEnabled(commandId js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQueryCommandEnabled(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return _ret == js.True, _ok
}

// QueryCommandEnabledFunc returns the method "Document.queryCommandEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QueryCommandEnabledFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandEnabledFunc(
			this.Ref(),
		),
	)
}

// QueryCommandIndeterm calls the method "Document.queryCommandIndeterm".
//
// The returned bool will be false if there is no such method.
func (this Document) QueryCommandIndeterm(commandId js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQueryCommandIndeterm(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return _ret == js.True, _ok
}

// QueryCommandIndetermFunc returns the method "Document.queryCommandIndeterm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QueryCommandIndetermFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandIndetermFunc(
			this.Ref(),
		),
	)
}

// QueryCommandState calls the method "Document.queryCommandState".
//
// The returned bool will be false if there is no such method.
func (this Document) QueryCommandState(commandId js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQueryCommandState(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return _ret == js.True, _ok
}

// QueryCommandStateFunc returns the method "Document.queryCommandState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QueryCommandStateFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandStateFunc(
			this.Ref(),
		),
	)
}

// QueryCommandSupported calls the method "Document.queryCommandSupported".
//
// The returned bool will be false if there is no such method.
func (this Document) QueryCommandSupported(commandId js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQueryCommandSupported(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return _ret == js.True, _ok
}

// QueryCommandSupportedFunc returns the method "Document.queryCommandSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QueryCommandSupportedFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandSupportedFunc(
			this.Ref(),
		),
	)
}

// QueryCommandValue calls the method "Document.queryCommandValue".
//
// The returned bool will be false if there is no such method.
func (this Document) QueryCommandValue(commandId js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQueryCommandValue(
		this.Ref(), js.Pointer(&_ok),
		commandId.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// QueryCommandValueFunc returns the method "Document.queryCommandValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QueryCommandValueFunc() (fn js.Func[func(commandId js.String) js.String]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandValueFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Document.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Document) GetBoxQuads(options BoxQuadOptions) (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetBoxQuads(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuadsFunc returns the method "Document.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.DocumentGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Document.getBoxQuads".
//
// The returned bool will be false if there is no such method.
func (this Document) GetBoxQuads1() (js.Array[DOMQuad], bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetBoxQuads1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[DOMQuad]{}.FromRef(_ret), _ok
}

// GetBoxQuads1Func returns the method "Document.getBoxQuads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.DocumentGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Document.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertQuadFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNodeFunc returns the method "Document.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Document.convertQuadFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertQuadFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&quad),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertQuadFromNode1Func returns the method "Document.convertQuadFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Document.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertRectFromNode(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNodeFunc returns the method "Document.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Document.convertRectFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (DOMQuad, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertRectFromNode1(
		this.Ref(), js.Pointer(&_ok),
		rect.Ref(),
		from.Ref(),
	)

	return DOMQuad{}.FromRef(_ret), _ok
}

// ConvertRectFromNode1Func returns the method "Document.convertRectFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Document.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertPointFromNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNodeFunc returns the method "Document.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.DocumentConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Document.convertPointFromNode".
//
// The returned bool will be false if there is no such method.
func (this Document) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (DOMPoint, bool) {
	var _ok bool
	_ret := bindings.CallDocumentConvertPointFromNode1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&point),
		from.Ref(),
	)

	return DOMPoint{}.FromRef(_ret), _ok
}

// ConvertPointFromNode1Func returns the method "Document.convertPointFromNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.DocumentConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// GetElementById calls the method "Document.getElementById".
//
// The returned bool will be false if there is no such method.
func (this Document) GetElementById(elementId js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetElementById(
		this.Ref(), js.Pointer(&_ok),
		elementId.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// GetElementByIdFunc returns the method "Document.getElementById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// GetAnimations calls the method "Document.getAnimations".
//
// The returned bool will be false if there is no such method.
func (this Document) GetAnimations() (js.Array[Animation], bool) {
	var _ok bool
	_ret := bindings.CallDocumentGetAnimations(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Animation]{}.FromRef(_ret), _ok
}

// GetAnimationsFunc returns the method "Document.getAnimations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) GetAnimationsFunc() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.DocumentGetAnimationsFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "Document.prepend".
//
// The returned bool will be false if there is no such method.
func (this Document) Prepend(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentPrepend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PrependFunc returns the method "Document.prepend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentPrependFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "Document.append".
//
// The returned bool will be false if there is no such method.
func (this Document) Append(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentAppend(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendFunc returns the method "Document.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentAppendFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "Document.replaceChildren".
//
// The returned bool will be false if there is no such method.
func (this Document) ReplaceChildren(nodes ...OneOf_Node_String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallDocumentReplaceChildren(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceChildrenFunc returns the method "Document.replaceChildren".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "Document.querySelector".
//
// The returned bool will be false if there is no such method.
func (this Document) QuerySelector(selectors js.String) (Element, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQuerySelector(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return Element{}.FromRef(_ret), _ok
}

// QuerySelectorFunc returns the method "Document.querySelector".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "Document.querySelectorAll".
//
// The returned bool will be false if there is no such method.
func (this Document) QuerySelectorAll(selectors js.String) (NodeList, bool) {
	var _ok bool
	_ret := bindings.CallDocumentQuerySelectorAll(
		this.Ref(), js.Pointer(&_ok),
		selectors.Ref(),
	)

	return NodeList{}.FromRef(_ret), _ok
}

// QuerySelectorAllFunc returns the method "Document.querySelectorAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

// CreateExpression calls the method "Document.createExpression".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (XPathExpression, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateExpression(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		resolver.Ref(),
	)

	return XPathExpression{}.FromRef(_ret), _ok
}

// CreateExpressionFunc returns the method "Document.createExpression".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateExpressionFunc() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	return fn.FromRef(
		bindings.DocumentCreateExpressionFunc(
			this.Ref(),
		),
	)
}

// CreateExpression1 calls the method "Document.createExpression".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateExpression1(expression js.String) (XPathExpression, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateExpression1(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
	)

	return XPathExpression{}.FromRef(_ret), _ok
}

// CreateExpression1Func returns the method "Document.createExpression".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateExpression1Func() (fn js.Func[func(expression js.String) XPathExpression]) {
	return fn.FromRef(
		bindings.DocumentCreateExpression1Func(
			this.Ref(),
		),
	)
}

// CreateNSResolver calls the method "Document.createNSResolver".
//
// The returned bool will be false if there is no such method.
func (this Document) CreateNSResolver(nodeResolver Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallDocumentCreateNSResolver(
		this.Ref(), js.Pointer(&_ok),
		nodeResolver.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// CreateNSResolverFunc returns the method "Document.createNSResolver".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) CreateNSResolverFunc() (fn js.Func[func(nodeResolver Node) Node]) {
	return fn.FromRef(
		bindings.DocumentCreateNSResolverFunc(
			this.Ref(),
		),
	)
}

// Evaluate calls the method "Document.evaluate".
//
// The returned bool will be false if there is no such method.
func (this Document) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallDocumentEvaluate(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// EvaluateFunc returns the method "Document.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) EvaluateFunc() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "Document.evaluate".
//
// The returned bool will be false if there is no such method.
func (this Document) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallDocumentEvaluate1(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate1Func returns the method "Document.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Evaluate1Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "Document.evaluate".
//
// The returned bool will be false if there is no such method.
func (this Document) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallDocumentEvaluate2(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate2Func returns the method "Document.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Evaluate2Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate2Func(
			this.Ref(),
		),
	)
}

// Evaluate3 calls the method "Document.evaluate".
//
// The returned bool will be false if there is no such method.
func (this Document) Evaluate3(expression js.String, contextNode Node) (XPathResult, bool) {
	var _ok bool
	_ret := bindings.CallDocumentEvaluate3(
		this.Ref(), js.Pointer(&_ok),
		expression.Ref(),
		contextNode.Ref(),
	)

	return XPathResult{}.FromRef(_ret), _ok
}

// Evaluate3Func returns the method "Document.evaluate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Document) Evaluate3Func() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate3Func(
			this.Ref(),
		),
	)
}

type Node struct {
	EventTarget
}

func (this Node) Once() Node {
	this.Ref().Once()
	return this
}

func (this Node) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Node) FromRef(ref js.Ref) Node {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Node) Free() {
	this.Ref().Free()
}

// NodeType returns the value of property "Node.nodeType".
//
// The returned bool will be false if there is no such property.
func (this Node) NodeType() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetNodeNodeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// NodeName returns the value of property "Node.nodeName".
//
// The returned bool will be false if there is no such property.
func (this Node) NodeName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNodeNodeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BaseURI returns the value of property "Node.baseURI".
//
// The returned bool will be false if there is no such property.
func (this Node) BaseURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNodeBaseURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IsConnected returns the value of property "Node.isConnected".
//
// The returned bool will be false if there is no such property.
func (this Node) IsConnected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNodeIsConnected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// OwnerDocument returns the value of property "Node.ownerDocument".
//
// The returned bool will be false if there is no such property.
func (this Node) OwnerDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.GetNodeOwnerDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// ParentNode returns the value of property "Node.parentNode".
//
// The returned bool will be false if there is no such property.
func (this Node) ParentNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeParentNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// ParentElement returns the value of property "Node.parentElement".
//
// The returned bool will be false if there is no such property.
func (this Node) ParentElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetNodeParentElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// ChildNodes returns the value of property "Node.childNodes".
//
// The returned bool will be false if there is no such property.
func (this Node) ChildNodes() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetNodeChildNodes(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// FirstChild returns the value of property "Node.firstChild".
//
// The returned bool will be false if there is no such property.
func (this Node) FirstChild() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeFirstChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// LastChild returns the value of property "Node.lastChild".
//
// The returned bool will be false if there is no such property.
func (this Node) LastChild() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeLastChild(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// PreviousSibling returns the value of property "Node.previousSibling".
//
// The returned bool will be false if there is no such property.
func (this Node) PreviousSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodePreviousSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// NextSibling returns the value of property "Node.nextSibling".
//
// The returned bool will be false if there is no such property.
func (this Node) NextSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.GetNodeNextSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// NodeValue returns the value of property "Node.nodeValue".
//
// The returned bool will be false if there is no such property.
func (this Node) NodeValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNodeNodeValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NodeValue sets the value of property "Node.nodeValue" to val.
//
// It returns false if the property cannot be set.
func (this Node) SetNodeValue(val js.String) bool {
	return js.True == bindings.SetNodeNodeValue(
		this.Ref(),
		val.Ref(),
	)
}

// TextContent returns the value of property "Node.textContent".
//
// The returned bool will be false if there is no such property.
func (this Node) TextContent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNodeTextContent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// TextContent sets the value of property "Node.textContent" to val.
//
// It returns false if the property cannot be set.
func (this Node) SetTextContent(val js.String) bool {
	return js.True == bindings.SetNodeTextContent(
		this.Ref(),
		val.Ref(),
	)
}

// GetRootNode calls the method "Node.getRootNode".
//
// The returned bool will be false if there is no such method.
func (this Node) GetRootNode(options GetRootNodeOptions) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeGetRootNode(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return Node{}.FromRef(_ret), _ok
}

// GetRootNodeFunc returns the method "Node.getRootNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) GetRootNodeFunc() (fn js.Func[func(options GetRootNodeOptions) Node]) {
	return fn.FromRef(
		bindings.NodeGetRootNodeFunc(
			this.Ref(),
		),
	)
}

// GetRootNode1 calls the method "Node.getRootNode".
//
// The returned bool will be false if there is no such method.
func (this Node) GetRootNode1() (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeGetRootNode1(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// GetRootNode1Func returns the method "Node.getRootNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) GetRootNode1Func() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeGetRootNode1Func(
			this.Ref(),
		),
	)
}

// HasChildNodes calls the method "Node.hasChildNodes".
//
// The returned bool will be false if there is no such method.
func (this Node) HasChildNodes() (bool, bool) {
	var _ok bool
	_ret := bindings.CallNodeHasChildNodes(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// HasChildNodesFunc returns the method "Node.hasChildNodes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) HasChildNodesFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NodeHasChildNodesFunc(
			this.Ref(),
		),
	)
}

// Normalize calls the method "Node.normalize".
//
// The returned bool will be false if there is no such method.
func (this Node) Normalize() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNodeNormalize(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// NormalizeFunc returns the method "Node.normalize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) NormalizeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NodeNormalizeFunc(
			this.Ref(),
		),
	)
}

// CloneNode calls the method "Node.cloneNode".
//
// The returned bool will be false if there is no such method.
func (this Node) CloneNode(deep bool) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeCloneNode(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(deep)),
	)

	return Node{}.FromRef(_ret), _ok
}

// CloneNodeFunc returns the method "Node.cloneNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) CloneNodeFunc() (fn js.Func[func(deep bool) Node]) {
	return fn.FromRef(
		bindings.NodeCloneNodeFunc(
			this.Ref(),
		),
	)
}

// CloneNode1 calls the method "Node.cloneNode".
//
// The returned bool will be false if there is no such method.
func (this Node) CloneNode1() (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeCloneNode1(
		this.Ref(), js.Pointer(&_ok),
	)

	return Node{}.FromRef(_ret), _ok
}

// CloneNode1Func returns the method "Node.cloneNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) CloneNode1Func() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeCloneNode1Func(
			this.Ref(),
		),
	)
}

// IsEqualNode calls the method "Node.isEqualNode".
//
// The returned bool will be false if there is no such method.
func (this Node) IsEqualNode(otherNode Node) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNodeIsEqualNode(
		this.Ref(), js.Pointer(&_ok),
		otherNode.Ref(),
	)

	return _ret == js.True, _ok
}

// IsEqualNodeFunc returns the method "Node.isEqualNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) IsEqualNodeFunc() (fn js.Func[func(otherNode Node) bool]) {
	return fn.FromRef(
		bindings.NodeIsEqualNodeFunc(
			this.Ref(),
		),
	)
}

// IsSameNode calls the method "Node.isSameNode".
//
// The returned bool will be false if there is no such method.
func (this Node) IsSameNode(otherNode Node) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNodeIsSameNode(
		this.Ref(), js.Pointer(&_ok),
		otherNode.Ref(),
	)

	return _ret == js.True, _ok
}

// IsSameNodeFunc returns the method "Node.isSameNode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) IsSameNodeFunc() (fn js.Func[func(otherNode Node) bool]) {
	return fn.FromRef(
		bindings.NodeIsSameNodeFunc(
			this.Ref(),
		),
	)
}

// CompareDocumentPosition calls the method "Node.compareDocumentPosition".
//
// The returned bool will be false if there is no such method.
func (this Node) CompareDocumentPosition(other Node) (uint16, bool) {
	var _ok bool
	_ret := bindings.CallNodeCompareDocumentPosition(
		this.Ref(), js.Pointer(&_ok),
		other.Ref(),
	)

	return uint16(_ret), _ok
}

// CompareDocumentPositionFunc returns the method "Node.compareDocumentPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) CompareDocumentPositionFunc() (fn js.Func[func(other Node) uint16]) {
	return fn.FromRef(
		bindings.NodeCompareDocumentPositionFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "Node.contains".
//
// The returned bool will be false if there is no such method.
func (this Node) Contains(other Node) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNodeContains(
		this.Ref(), js.Pointer(&_ok),
		other.Ref(),
	)

	return _ret == js.True, _ok
}

// ContainsFunc returns the method "Node.contains".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) ContainsFunc() (fn js.Func[func(other Node) bool]) {
	return fn.FromRef(
		bindings.NodeContainsFunc(
			this.Ref(),
		),
	)
}

// LookupPrefix calls the method "Node.lookupPrefix".
//
// The returned bool will be false if there is no such method.
func (this Node) LookupPrefix(namespace js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallNodeLookupPrefix(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// LookupPrefixFunc returns the method "Node.lookupPrefix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) LookupPrefixFunc() (fn js.Func[func(namespace js.String) js.String]) {
	return fn.FromRef(
		bindings.NodeLookupPrefixFunc(
			this.Ref(),
		),
	)
}

// LookupNamespaceURI calls the method "Node.lookupNamespaceURI".
//
// The returned bool will be false if there is no such method.
func (this Node) LookupNamespaceURI(prefix js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallNodeLookupNamespaceURI(
		this.Ref(), js.Pointer(&_ok),
		prefix.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// LookupNamespaceURIFunc returns the method "Node.lookupNamespaceURI".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) LookupNamespaceURIFunc() (fn js.Func[func(prefix js.String) js.String]) {
	return fn.FromRef(
		bindings.NodeLookupNamespaceURIFunc(
			this.Ref(),
		),
	)
}

// IsDefaultNamespace calls the method "Node.isDefaultNamespace".
//
// The returned bool will be false if there is no such method.
func (this Node) IsDefaultNamespace(namespace js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNodeIsDefaultNamespace(
		this.Ref(), js.Pointer(&_ok),
		namespace.Ref(),
	)

	return _ret == js.True, _ok
}

// IsDefaultNamespaceFunc returns the method "Node.isDefaultNamespace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) IsDefaultNamespaceFunc() (fn js.Func[func(namespace js.String) bool]) {
	return fn.FromRef(
		bindings.NodeIsDefaultNamespaceFunc(
			this.Ref(),
		),
	)
}

// InsertBefore calls the method "Node.insertBefore".
//
// The returned bool will be false if there is no such method.
func (this Node) InsertBefore(node Node, child Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeInsertBefore(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		child.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// InsertBeforeFunc returns the method "Node.insertBefore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) InsertBeforeFunc() (fn js.Func[func(node Node, child Node) Node]) {
	return fn.FromRef(
		bindings.NodeInsertBeforeFunc(
			this.Ref(),
		),
	)
}

// AppendChild calls the method "Node.appendChild".
//
// The returned bool will be false if there is no such method.
func (this Node) AppendChild(node Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeAppendChild(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// AppendChildFunc returns the method "Node.appendChild".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) AppendChildFunc() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.NodeAppendChildFunc(
			this.Ref(),
		),
	)
}

// ReplaceChild calls the method "Node.replaceChild".
//
// The returned bool will be false if there is no such method.
func (this Node) ReplaceChild(node Node, child Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeReplaceChild(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
		child.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// ReplaceChildFunc returns the method "Node.replaceChild".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) ReplaceChildFunc() (fn js.Func[func(node Node, child Node) Node]) {
	return fn.FromRef(
		bindings.NodeReplaceChildFunc(
			this.Ref(),
		),
	)
}

// RemoveChild calls the method "Node.removeChild".
//
// The returned bool will be false if there is no such method.
func (this Node) RemoveChild(child Node) (Node, bool) {
	var _ok bool
	_ret := bindings.CallNodeRemoveChild(
		this.Ref(), js.Pointer(&_ok),
		child.Ref(),
	)

	return Node{}.FromRef(_ret), _ok
}

// RemoveChildFunc returns the method "Node.removeChild".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Node) RemoveChildFunc() (fn js.Func[func(child Node) Node]) {
	return fn.FromRef(
		bindings.NodeRemoveChildFunc(
			this.Ref(),
		),
	)
}

type AbstractRange struct {
	ref js.Ref
}

func (this AbstractRange) Once() AbstractRange {
	this.Ref().Once()
	return this
}

func (this AbstractRange) Ref() js.Ref {
	return this.ref
}

func (this AbstractRange) FromRef(ref js.Ref) AbstractRange {
	this.ref = ref
	return this
}

func (this AbstractRange) Free() {
	this.Ref().Free()
}

// StartContainer returns the value of property "AbstractRange.startContainer".
//
// The returned bool will be false if there is no such property.
func (this AbstractRange) StartContainer() (Node, bool) {
	var _ok bool
	_ret := bindings.GetAbstractRangeStartContainer(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// StartOffset returns the value of property "AbstractRange.startOffset".
//
// The returned bool will be false if there is no such property.
func (this AbstractRange) StartOffset() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAbstractRangeStartOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// EndContainer returns the value of property "AbstractRange.endContainer".
//
// The returned bool will be false if there is no such property.
func (this AbstractRange) EndContainer() (Node, bool) {
	var _ok bool
	_ret := bindings.GetAbstractRangeEndContainer(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// EndOffset returns the value of property "AbstractRange.endOffset".
//
// The returned bool will be false if there is no such property.
func (this AbstractRange) EndOffset() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAbstractRangeEndOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Collapsed returns the value of property "AbstractRange.collapsed".
//
// The returned bool will be false if there is no such property.
func (this AbstractRange) Collapsed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAbstractRangeCollapsed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type AccelerometerLocalCoordinateSystem uint32

const (
	_ AccelerometerLocalCoordinateSystem = iota

	AccelerometerLocalCoordinateSystem_DEVICE
	AccelerometerLocalCoordinateSystem_SCREEN
)

func (AccelerometerLocalCoordinateSystem) FromRef(str js.Ref) AccelerometerLocalCoordinateSystem {
	return AccelerometerLocalCoordinateSystem(bindings.ConstOfAccelerometerLocalCoordinateSystem(str))
}

func (x AccelerometerLocalCoordinateSystem) String() (string, bool) {
	switch x {
	case AccelerometerLocalCoordinateSystem_DEVICE:
		return "device", true
	case AccelerometerLocalCoordinateSystem_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type AccelerometerSensorOptions struct {
	// ReferenceFrame is "AccelerometerSensorOptions.referenceFrame"
	//
	// Optional, defaults to "device".
	ReferenceFrame AccelerometerLocalCoordinateSystem
	// Frequency is "AccelerometerSensorOptions.frequency"
	//
	// Optional
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccelerometerSensorOptions with all fields set.
func (p AccelerometerSensorOptions) FromRef(ref js.Ref) AccelerometerSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AccelerometerSensorOptions AccelerometerSensorOptions [// AccelerometerSensorOptions] [0x1400033f360 0x1400033f400 0x1400033f4a0] 0x14001af2ca8 {0 0}} in the application heap.
func (p AccelerometerSensorOptions) New() js.Ref {
	return bindings.AccelerometerSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AccelerometerSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.AccelerometerSensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AccelerometerSensorOptions) Update(ref js.Ref) {
	bindings.AccelerometerSensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAccelerometer(options AccelerometerSensorOptions) Accelerometer {
	return Accelerometer{}.FromRef(
		bindings.NewAccelerometerByAccelerometer(
			js.Pointer(&options)),
	)
}

func NewAccelerometerByAccelerometer1() Accelerometer {
	return Accelerometer{}.FromRef(
		bindings.NewAccelerometerByAccelerometer1(),
	)
}

type Accelerometer struct {
	Sensor
}

func (this Accelerometer) Once() Accelerometer {
	this.Ref().Once()
	return this
}

func (this Accelerometer) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this Accelerometer) FromRef(ref js.Ref) Accelerometer {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this Accelerometer) Free() {
	this.Ref().Free()
}

// X returns the value of property "Accelerometer.x".
//
// The returned bool will be false if there is no such property.
func (this Accelerometer) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAccelerometerX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "Accelerometer.y".
//
// The returned bool will be false if there is no such property.
func (this Accelerometer) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAccelerometerY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "Accelerometer.z".
//
// The returned bool will be false if there is no such property.
func (this Accelerometer) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAccelerometerZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type AccelerometerReadingValues struct {
	// X is "AccelerometerReadingValues.x"
	//
	// Required
	X float64
	// Y is "AccelerometerReadingValues.y"
	//
	// Required
	Y float64
	// Z is "AccelerometerReadingValues.z"
	//
	// Required
	Z float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccelerometerReadingValues with all fields set.
func (p AccelerometerReadingValues) FromRef(ref js.Ref) AccelerometerReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AccelerometerReadingValues AccelerometerReadingValues [// AccelerometerReadingValues] [0x1400033f5e0 0x1400033f680 0x1400033f720] 0x14001af2d08 {0 0}} in the application heap.
func (p AccelerometerReadingValues) New() js.Ref {
	return bindings.AccelerometerReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AccelerometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AccelerometerReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AccelerometerReadingValues) Update(ref js.Ref) {
	bindings.AccelerometerReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AdRender struct {
	// Url is "AdRender.url"
	//
	// Required
	Url js.String
	// Width is "AdRender.width"
	//
	// Optional
	Width js.String
	// Height is "AdRender.height"
	//
	// Optional
	Height js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AdRender with all fields set.
func (p AdRender) FromRef(ref js.Ref) AdRender {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AdRender AdRender [// AdRender] [0x1400033f7c0 0x1400033f860 0x1400033f900] 0x14001af2d80 {0 0}} in the application heap.
func (p AdRender) New() js.Ref {
	return bindings.AdRenderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AdRender) UpdateFrom(ref js.Ref) {
	bindings.AdRenderJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AdRender) Update(ref js.Ref) {
	bindings.AdRenderJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AddressErrors struct {
	// AddressLine is "AddressErrors.addressLine"
	//
	// Optional
	AddressLine js.String
	// City is "AddressErrors.city"
	//
	// Optional
	City js.String
	// Country is "AddressErrors.country"
	//
	// Optional
	Country js.String
	// DependentLocality is "AddressErrors.dependentLocality"
	//
	// Optional
	DependentLocality js.String
	// Organization is "AddressErrors.organization"
	//
	// Optional
	Organization js.String
	// Phone is "AddressErrors.phone"
	//
	// Optional
	Phone js.String
	// PostalCode is "AddressErrors.postalCode"
	//
	// Optional
	PostalCode js.String
	// Recipient is "AddressErrors.recipient"
	//
	// Optional
	Recipient js.String
	// Region is "AddressErrors.region"
	//
	// Optional
	Region js.String
	// SortingCode is "AddressErrors.sortingCode"
	//
	// Optional
	SortingCode js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressErrors with all fields set.
func (p AddressErrors) FromRef(ref js.Ref) AddressErrors {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AddressErrors AddressErrors [// AddressErrors] [0x1400033f9a0 0x1400033fa40 0x1400033fae0 0x1400033fb80 0x1400033fc20 0x1400033fcc0 0x1400033fd60 0x1400033fe00 0x1400033fea0 0x1400033ff40] 0x14001af2db0 {0 0}} in the application heap.
func (p AddressErrors) New() js.Ref {
	return bindings.AddressErrorsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AddressErrors) UpdateFrom(ref js.Ref) {
	bindings.AddressErrorsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AddressErrors) Update(ref js.Ref) {
	bindings.AddressErrorsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AddressInit struct {
	// Country is "AddressInit.country"
	//
	// Optional, defaults to "".
	Country js.String
	// AddressLine is "AddressInit.addressLine"
	//
	// Optional, defaults to [].
	AddressLine js.Array[js.String]
	// Region is "AddressInit.region"
	//
	// Optional, defaults to "".
	Region js.String
	// City is "AddressInit.city"
	//
	// Optional, defaults to "".
	City js.String
	// DependentLocality is "AddressInit.dependentLocality"
	//
	// Optional, defaults to "".
	DependentLocality js.String
	// PostalCode is "AddressInit.postalCode"
	//
	// Optional, defaults to "".
	PostalCode js.String
	// SortingCode is "AddressInit.sortingCode"
	//
	// Optional, defaults to "".
	SortingCode js.String
	// Organization is "AddressInit.organization"
	//
	// Optional, defaults to "".
	Organization js.String
	// Recipient is "AddressInit.recipient"
	//
	// Optional, defaults to "".
	Recipient js.String
	// Phone is "AddressInit.phone"
	//
	// Optional, defaults to "".
	Phone js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddressInit with all fields set.
func (p AddressInit) FromRef(ref js.Ref) AddressInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AddressInit AddressInit [// AddressInit] [0x14000126140 0x14000126280 0x140001263c0 0x14000126500 0x14000126640 0x14000126780 0x140001268c0 0x14000126a00 0x14000126b40 0x14000126c80] 0x14001af2de0 {0 0}} in the application heap.
func (p AddressInit) New() js.Ref {
	return bindings.AddressInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AddressInit) UpdateFrom(ref js.Ref) {
	bindings.AddressInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AddressInit) Update(ref js.Ref) {
	bindings.AddressInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

type BufferSource = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer

type AesCbcParams struct {
	// Iv is "AesCbcParams.iv"
	//
	// Required
	Iv BufferSource
	// Name is "AesCbcParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesCbcParams with all fields set.
func (p AesCbcParams) FromRef(ref js.Ref) AesCbcParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesCbcParams AesCbcParams [// AesCbcParams] [0x14000126dc0 0x14000126f00] 0x14001af2e10 {0 0}} in the application heap.
func (p AesCbcParams) New() js.Ref {
	return bindings.AesCbcParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesCbcParams) UpdateFrom(ref js.Ref) {
	bindings.AesCbcParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesCbcParams) Update(ref js.Ref) {
	bindings.AesCbcParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AesCtrParams struct {
	// Counter is "AesCtrParams.counter"
	//
	// Required
	Counter BufferSource
	// Length is "AesCtrParams.length"
	//
	// Required
	Length uint8
	// Name is "AesCtrParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesCtrParams with all fields set.
func (p AesCtrParams) FromRef(ref js.Ref) AesCtrParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesCtrParams AesCtrParams [// AesCtrParams] [0x14000127040 0x14000127180 0x140001272c0] 0x14001af2ed0 {0 0}} in the application heap.
func (p AesCtrParams) New() js.Ref {
	return bindings.AesCtrParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesCtrParams) UpdateFrom(ref js.Ref) {
	bindings.AesCtrParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesCtrParams) Update(ref js.Ref) {
	bindings.AesCtrParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AesDerivedKeyParams struct {
	// Length is "AesDerivedKeyParams.length"
	//
	// Required
	Length uint16
	// Name is "AesDerivedKeyParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesDerivedKeyParams with all fields set.
func (p AesDerivedKeyParams) FromRef(ref js.Ref) AesDerivedKeyParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesDerivedKeyParams AesDerivedKeyParams [// AesDerivedKeyParams] [0x14000127400 0x14000127540] 0x14001af2f18 {0 0}} in the application heap.
func (p AesDerivedKeyParams) New() js.Ref {
	return bindings.AesDerivedKeyParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesDerivedKeyParams) UpdateFrom(ref js.Ref) {
	bindings.AesDerivedKeyParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesDerivedKeyParams) Update(ref js.Ref) {
	bindings.AesDerivedKeyParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AesGcmParams struct {
	// Iv is "AesGcmParams.iv"
	//
	// Required
	Iv BufferSource
	// AdditionalData is "AesGcmParams.additionalData"
	//
	// Optional
	AdditionalData BufferSource
	// TagLength is "AesGcmParams.tagLength"
	//
	// Optional
	TagLength uint8
	// Name is "AesGcmParams.name"
	//
	// Required
	Name js.String

	FFI_USE_TagLength bool // for TagLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesGcmParams with all fields set.
func (p AesGcmParams) FromRef(ref js.Ref) AesGcmParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesGcmParams AesGcmParams [// AesGcmParams] [0x14000127680 0x140001277c0 0x14000127900 0x14000127b80 0x14000127a40] 0x14001af2f48 {0 0}} in the application heap.
func (p AesGcmParams) New() js.Ref {
	return bindings.AesGcmParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesGcmParams) UpdateFrom(ref js.Ref) {
	bindings.AesGcmParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesGcmParams) Update(ref js.Ref) {
	bindings.AesGcmParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AesKeyAlgorithm struct {
	// Length is "AesKeyAlgorithm.length"
	//
	// Required
	Length uint16
	// Name is "AesKeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesKeyAlgorithm with all fields set.
func (p AesKeyAlgorithm) FromRef(ref js.Ref) AesKeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesKeyAlgorithm AesKeyAlgorithm [// AesKeyAlgorithm] [0x14000127cc0 0x14000a36000] 0x14001af2fc0 {0 0}} in the application heap.
func (p AesKeyAlgorithm) New() js.Ref {
	return bindings.AesKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.AesKeyAlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesKeyAlgorithm) Update(ref js.Ref) {
	bindings.AesKeyAlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AesKeyGenParams struct {
	// Length is "AesKeyGenParams.length"
	//
	// Required
	Length uint16
	// Name is "AesKeyGenParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AesKeyGenParams with all fields set.
func (p AesKeyGenParams) FromRef(ref js.Ref) AesKeyGenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AesKeyGenParams AesKeyGenParams [// AesKeyGenParams] [0x14000a360a0 0x14000a36140] 0x14001af3008 {0 0}} in the application heap.
func (p AesKeyGenParams) New() js.Ref {
	return bindings.AesKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AesKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.AesKeyGenParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AesKeyGenParams) Update(ref js.Ref) {
	bindings.AesKeyGenParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Algorithm struct {
	// Name is "Algorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Algorithm with all fields set.
func (p Algorithm) FromRef(ref js.Ref) Algorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 Algorithm Algorithm [// Algorithm] [0x14000a361e0] 0x14001af3038 {0 0}} in the application heap.
func (p Algorithm) New() js.Ref {
	return bindings.AlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p Algorithm) UpdateFrom(ref js.Ref) {
	bindings.AlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p Algorithm) Update(ref js.Ref) {
	bindings.AlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Object_String struct {
	ref js.Ref
}

func (x OneOf_Object_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Object_String) Free() {
	x.ref.Free()
}

func (x OneOf_Object_String) FromRef(ref js.Ref) OneOf_Object_String {
	return OneOf_Object_String{
		ref: ref,
	}
}

func (x OneOf_Object_String) Object() js.Object {
	return js.Object{}.FromRef(x.ref)
}

func (x OneOf_Object_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type AlgorithmIdentifier = OneOf_Object_String

type AlignSetting uint32

const (
	_ AlignSetting = iota

	AlignSetting_START
	AlignSetting_CENTER
	AlignSetting_END
	AlignSetting_LEFT
	AlignSetting_RIGHT
)

func (AlignSetting) FromRef(str js.Ref) AlignSetting {
	return AlignSetting(bindings.ConstOfAlignSetting(str))
}

func (x AlignSetting) String() (string, bool) {
	switch x {
	case AlignSetting_START:
		return "start", true
	case AlignSetting_CENTER:
		return "center", true
	case AlignSetting_END:
		return "end", true
	case AlignSetting_LEFT:
		return "left", true
	case AlignSetting_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView struct {
	ref js.Ref
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FromRef(ref js.Ref) OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView {
	return OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView{
		ref: ref,
	}
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) SharedArrayBuffer() js.SharedArrayBuffer {
	return js.SharedArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

type AllowSharedBufferSource = OneOf_ArrayBuffer_SharedArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView

type UUID = js.String

type OneOf_String_ArrayUUID struct {
	ref js.Ref
}

func (x OneOf_String_ArrayUUID) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayUUID) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayUUID) FromRef(ref js.Ref) OneOf_String_ArrayUUID {
	return OneOf_String_ArrayUUID{
		ref: ref,
	}
}

func (x OneOf_String_ArrayUUID) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayUUID) ArrayUUID() js.Array[UUID] {
	return js.Array[UUID]{}.FromRef(x.ref)
}

type AllowedBluetoothDevice struct {
	// DeviceId is "AllowedBluetoothDevice.deviceId"
	//
	// Required
	DeviceId js.String
	// MayUseGATT is "AllowedBluetoothDevice.mayUseGATT"
	//
	// Required
	MayUseGATT bool
	// AllowedServices is "AllowedBluetoothDevice.allowedServices"
	//
	// Required
	AllowedServices OneOf_String_ArrayUUID
	// AllowedManufacturerData is "AllowedBluetoothDevice.allowedManufacturerData"
	//
	// Required
	AllowedManufacturerData js.Array[uint16]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AllowedBluetoothDevice with all fields set.
func (p AllowedBluetoothDevice) FromRef(ref js.Ref) AllowedBluetoothDevice {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AllowedBluetoothDevice AllowedBluetoothDevice [// AllowedBluetoothDevice] [0x14000a36280 0x14000a36320 0x14000a363c0 0x14000a36460] 0x14001af30f8 {0 0}} in the application heap.
func (p AllowedBluetoothDevice) New() js.Ref {
	return bindings.AllowedBluetoothDeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AllowedBluetoothDevice) UpdateFrom(ref js.Ref) {
	bindings.AllowedBluetoothDeviceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AllowedBluetoothDevice) Update(ref js.Ref) {
	bindings.AllowedBluetoothDeviceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AllowedUSBDevice struct {
	// VendorId is "AllowedUSBDevice.vendorId"
	//
	// Required
	VendorId uint8
	// ProductId is "AllowedUSBDevice.productId"
	//
	// Required
	ProductId uint8
	// SerialNumber is "AllowedUSBDevice.serialNumber"
	//
	// Optional
	SerialNumber js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AllowedUSBDevice with all fields set.
func (p AllowedUSBDevice) FromRef(ref js.Ref) AllowedUSBDevice {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AllowedUSBDevice AllowedUSBDevice [// AllowedUSBDevice] [0x14000a36500 0x14000a365a0 0x14000a36640] 0x14001af3128 {0 0}} in the application heap.
func (p AllowedUSBDevice) New() js.Ref {
	return bindings.AllowedUSBDeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AllowedUSBDevice) UpdateFrom(ref js.Ref) {
	bindings.AllowedUSBDeviceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AllowedUSBDevice) Update(ref js.Ref) {
	bindings.AllowedUSBDeviceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AlphaOption uint32

const (
	_ AlphaOption = iota

	AlphaOption_KEEP
	AlphaOption_DISCARD
)

func (AlphaOption) FromRef(str js.Ref) AlphaOption {
	return AlphaOption(bindings.ConstOfAlphaOption(str))
}

func (x AlphaOption) String() (string, bool) {
	switch x {
	case AlphaOption_KEEP:
		return "keep", true
	case AlphaOption_DISCARD:
		return "discard", true
	default:
		return "", false
	}
}

type AmbientLightReadingValues struct {
	// Illuminance is "AmbientLightReadingValues.illuminance"
	//
	// Required
	Illuminance float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AmbientLightReadingValues with all fields set.
func (p AmbientLightReadingValues) FromRef(ref js.Ref) AmbientLightReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 AmbientLightReadingValues AmbientLightReadingValues [// AmbientLightReadingValues] [0x14000a366e0] 0x14001af3140 {0 0}} in the application heap.
func (p AmbientLightReadingValues) New() js.Ref {
	return bindings.AmbientLightReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AmbientLightReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AmbientLightReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AmbientLightReadingValues) Update(ref js.Ref) {
	bindings.AmbientLightReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SensorOptions struct {
	// Frequency is "SensorOptions.frequency"
	//
	// Optional
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SensorOptions with all fields set.
func (p SensorOptions) FromRef(ref js.Ref) SensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SensorOptions SensorOptions [// SensorOptions] [0x14000a36780 0x14000a36820] 0x14001af3170 {0 0}} in the application heap.
func (p SensorOptions) New() js.Ref {
	return bindings.SensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SensorOptions) UpdateFrom(ref js.Ref) {
	bindings.SensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SensorOptions) Update(ref js.Ref) {
	bindings.SensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
