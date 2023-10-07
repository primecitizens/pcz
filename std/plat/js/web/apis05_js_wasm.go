// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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

// New creates a new FontFaceDescriptors in the application heap.
func (p FontFaceDescriptors) New() js.Ref {
	return bindings.FontFaceDescriptorsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FontFaceDescriptors) UpdateFrom(ref js.Ref) {
	bindings.FontFaceDescriptorsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FontFaceDescriptors) Update(ref js.Ref) {
	bindings.FontFaceDescriptorsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FontFaceDescriptors) FreeMembers(recursive bool) {
	js.Free(
		p.Style.Ref(),
		p.Weight.Ref(),
		p.Stretch.Ref(),
		p.UnicodeRange.Ref(),
		p.Variant.Ref(),
		p.FeatureSettings.Ref(),
		p.VariationSettings.Ref(),
		p.Display.Ref(),
		p.AscentOverride.Ref(),
		p.DescentOverride.Ref(),
		p.LineGapOverride.Ref(),
	)
	p.Style = p.Style.FromRef(js.Undefined)
	p.Weight = p.Weight.FromRef(js.Undefined)
	p.Stretch = p.Stretch.FromRef(js.Undefined)
	p.UnicodeRange = p.UnicodeRange.FromRef(js.Undefined)
	p.Variant = p.Variant.FromRef(js.Undefined)
	p.FeatureSettings = p.FeatureSettings.FromRef(js.Undefined)
	p.VariationSettings = p.VariationSettings.FromRef(js.Undefined)
	p.Display = p.Display.FromRef(js.Undefined)
	p.AscentOverride = p.AscentOverride.FromRef(js.Undefined)
	p.DescentOverride = p.DescentOverride.FromRef(js.Undefined)
	p.LineGapOverride = p.LineGapOverride.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

type FontFaceVariations struct {
	ref js.Ref
}

func (this FontFaceVariations) Once() FontFaceVariations {
	this.ref.Once()
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
	this.ref.Free()
}

type FontFacePalette struct {
	ref js.Ref
}

func (this FontFacePalette) Once() FontFacePalette {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "FontFacePalette.length".
//
// It returns ok=false if there is no such property.
func (this FontFacePalette) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsableWithLightBackground returns the value of property "FontFacePalette.usableWithLightBackground".
//
// It returns ok=false if there is no such property.
func (this FontFacePalette) UsableWithLightBackground() (ret bool, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteUsableWithLightBackground(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsableWithDarkBackground returns the value of property "FontFacePalette.usableWithDarkBackground".
//
// It returns ok=false if there is no such property.
func (this FontFacePalette) UsableWithDarkBackground() (ret bool, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteUsableWithDarkBackground(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "FontFacePalette." exists.
func (this FontFacePalette) HasFuncGet() bool {
	return js.True == bindings.HasFuncFontFacePaletteGet(
		this.ref,
	)
}

// FuncGet returns the method "FontFacePalette.".
func (this FontFacePalette) FuncGet() (fn js.Func[func(index uint32) js.String]) {
	bindings.FuncFontFacePaletteGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "FontFacePalette.".
func (this FontFacePalette) Get(index uint32) (ret js.String) {
	bindings.CallFontFacePaletteGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "FontFacePalette."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFacePalette) TryGet(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFacePaletteGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type FontFacePalettes struct {
	ref js.Ref
}

func (this FontFacePalettes) Once() FontFacePalettes {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "FontFacePalettes.length".
//
// It returns ok=false if there is no such property.
func (this FontFacePalettes) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontFacePalettesLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "FontFacePalettes." exists.
func (this FontFacePalettes) HasFuncGet() bool {
	return js.True == bindings.HasFuncFontFacePalettesGet(
		this.ref,
	)
}

// FuncGet returns the method "FontFacePalettes.".
func (this FontFacePalettes) FuncGet() (fn js.Func[func(index uint32) FontFacePalette]) {
	bindings.FuncFontFacePalettesGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "FontFacePalettes.".
func (this FontFacePalettes) Get(index uint32) (ret FontFacePalette) {
	bindings.CallFontFacePalettesGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "FontFacePalettes."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFacePalettes) TryGet(index uint32) (ret FontFacePalette, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFacePalettesGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewFontFace(family js.String, source OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView, descriptors FontFaceDescriptors) (ret FontFace) {
	ret.ref = bindings.NewFontFaceByFontFace(
		family.Ref(),
		source.Ref(),
		js.Pointer(&descriptors))
	return
}

func NewFontFaceByFontFace1(family js.String, source OneOf_String_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) (ret FontFace) {
	ret.ref = bindings.NewFontFaceByFontFace1(
		family.Ref(),
		source.Ref())
	return
}

type FontFace struct {
	ref js.Ref
}

func (this FontFace) Once() FontFace {
	this.ref.Once()
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
	this.ref.Free()
}

// Family returns the value of property "FontFace.family".
//
// It returns ok=false if there is no such property.
func (this FontFace) Family() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceFamily(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFamily sets the value of property "FontFace.family" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetFamily(val js.String) bool {
	return js.True == bindings.SetFontFaceFamily(
		this.ref,
		val.Ref(),
	)
}

// Style returns the value of property "FontFace.style".
//
// It returns ok=false if there is no such property.
func (this FontFace) Style() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStyle sets the value of property "FontFace.style" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetStyle(val js.String) bool {
	return js.True == bindings.SetFontFaceStyle(
		this.ref,
		val.Ref(),
	)
}

// Weight returns the value of property "FontFace.weight".
//
// It returns ok=false if there is no such property.
func (this FontFace) Weight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceWeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWeight sets the value of property "FontFace.weight" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetWeight(val js.String) bool {
	return js.True == bindings.SetFontFaceWeight(
		this.ref,
		val.Ref(),
	)
}

// Stretch returns the value of property "FontFace.stretch".
//
// It returns ok=false if there is no such property.
func (this FontFace) Stretch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceStretch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStretch sets the value of property "FontFace.stretch" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetStretch(val js.String) bool {
	return js.True == bindings.SetFontFaceStretch(
		this.ref,
		val.Ref(),
	)
}

// UnicodeRange returns the value of property "FontFace.unicodeRange".
//
// It returns ok=false if there is no such property.
func (this FontFace) UnicodeRange() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceUnicodeRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUnicodeRange sets the value of property "FontFace.unicodeRange" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetUnicodeRange(val js.String) bool {
	return js.True == bindings.SetFontFaceUnicodeRange(
		this.ref,
		val.Ref(),
	)
}

// Variant returns the value of property "FontFace.variant".
//
// It returns ok=false if there is no such property.
func (this FontFace) Variant() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVariant sets the value of property "FontFace.variant" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetVariant(val js.String) bool {
	return js.True == bindings.SetFontFaceVariant(
		this.ref,
		val.Ref(),
	)
}

// FeatureSettings returns the value of property "FontFace.featureSettings".
//
// It returns ok=false if there is no such property.
func (this FontFace) FeatureSettings() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceFeatureSettings(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFeatureSettings sets the value of property "FontFace.featureSettings" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetFeatureSettings(val js.String) bool {
	return js.True == bindings.SetFontFaceFeatureSettings(
		this.ref,
		val.Ref(),
	)
}

// VariationSettings returns the value of property "FontFace.variationSettings".
//
// It returns ok=false if there is no such property.
func (this FontFace) VariationSettings() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationSettings(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVariationSettings sets the value of property "FontFace.variationSettings" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetVariationSettings(val js.String) bool {
	return js.True == bindings.SetFontFaceVariationSettings(
		this.ref,
		val.Ref(),
	)
}

// Display returns the value of property "FontFace.display".
//
// It returns ok=false if there is no such property.
func (this FontFace) Display() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceDisplay(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisplay sets the value of property "FontFace.display" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetDisplay(val js.String) bool {
	return js.True == bindings.SetFontFaceDisplay(
		this.ref,
		val.Ref(),
	)
}

// AscentOverride returns the value of property "FontFace.ascentOverride".
//
// It returns ok=false if there is no such property.
func (this FontFace) AscentOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceAscentOverride(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAscentOverride sets the value of property "FontFace.ascentOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetAscentOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceAscentOverride(
		this.ref,
		val.Ref(),
	)
}

// DescentOverride returns the value of property "FontFace.descentOverride".
//
// It returns ok=false if there is no such property.
func (this FontFace) DescentOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceDescentOverride(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDescentOverride sets the value of property "FontFace.descentOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetDescentOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceDescentOverride(
		this.ref,
		val.Ref(),
	)
}

// LineGapOverride returns the value of property "FontFace.lineGapOverride".
//
// It returns ok=false if there is no such property.
func (this FontFace) LineGapOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceLineGapOverride(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineGapOverride sets the value of property "FontFace.lineGapOverride" to val.
//
// It returns false if the property cannot be set.
func (this FontFace) SetLineGapOverride(val js.String) bool {
	return js.True == bindings.SetFontFaceLineGapOverride(
		this.ref,
		val.Ref(),
	)
}

// Status returns the value of property "FontFace.status".
//
// It returns ok=false if there is no such property.
func (this FontFace) Status() (ret FontFaceLoadStatus, ok bool) {
	ok = js.True == bindings.GetFontFaceStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Loaded returns the value of property "FontFace.loaded".
//
// It returns ok=false if there is no such property.
func (this FontFace) Loaded() (ret js.Promise[FontFace], ok bool) {
	ok = js.True == bindings.GetFontFaceLoaded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Features returns the value of property "FontFace.features".
//
// It returns ok=false if there is no such property.
func (this FontFace) Features() (ret FontFaceFeatures, ok bool) {
	ok = js.True == bindings.GetFontFaceFeatures(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Variations returns the value of property "FontFace.variations".
//
// It returns ok=false if there is no such property.
func (this FontFace) Variations() (ret FontFaceVariations, ok bool) {
	ok = js.True == bindings.GetFontFaceVariations(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Palettes returns the value of property "FontFace.palettes".
//
// It returns ok=false if there is no such property.
func (this FontFace) Palettes() (ret FontFacePalettes, ok bool) {
	ok = js.True == bindings.GetFontFacePalettes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncLoad returns true if the method "FontFace.load" exists.
func (this FontFace) HasFuncLoad() bool {
	return js.True == bindings.HasFuncFontFaceLoad(
		this.ref,
	)
}

// FuncLoad returns the method "FontFace.load".
func (this FontFace) FuncLoad() (fn js.Func[func() js.Promise[FontFace]]) {
	bindings.FuncFontFaceLoad(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Load calls the method "FontFace.load".
func (this FontFace) Load() (ret js.Promise[FontFace]) {
	bindings.CallFontFaceLoad(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLoad calls the method "FontFace.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFace) TryLoad() (ret js.Promise[FontFace], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceLoad(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewFontFaceSet(initialFaces js.Array[FontFace]) (ret FontFaceSet) {
	ret.ref = bindings.NewFontFaceSetByFontFaceSet(
		initialFaces.Ref())
	return
}

type FontFaceSet struct {
	EventTarget
}

func (this FontFaceSet) Once() FontFaceSet {
	this.ref.Once()
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
	this.ref.Free()
}

// Ready returns the value of property "FontFaceSet.ready".
//
// It returns ok=false if there is no such property.
func (this FontFaceSet) Ready() (ret js.Promise[FontFaceSet], ok bool) {
	ok = js.True == bindings.GetFontFaceSetReady(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "FontFaceSet.status".
//
// It returns ok=false if there is no such property.
func (this FontFaceSet) Status() (ret FontFaceSetLoadStatus, ok bool) {
	ok = js.True == bindings.GetFontFaceSetStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAdd returns true if the method "FontFaceSet.add" exists.
func (this FontFaceSet) HasFuncAdd() bool {
	return js.True == bindings.HasFuncFontFaceSetAdd(
		this.ref,
	)
}

// FuncAdd returns the method "FontFaceSet.add".
func (this FontFaceSet) FuncAdd() (fn js.Func[func(font FontFace) FontFaceSet]) {
	bindings.FuncFontFaceSetAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "FontFaceSet.add".
func (this FontFaceSet) Add(font FontFace) (ret FontFaceSet) {
	bindings.CallFontFaceSetAdd(
		this.ref, js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryAdd calls the method "FontFaceSet.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryAdd(font FontFace) (ret FontFaceSet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "FontFaceSet.delete" exists.
func (this FontFaceSet) HasFuncDelete() bool {
	return js.True == bindings.HasFuncFontFaceSetDelete(
		this.ref,
	)
}

// FuncDelete returns the method "FontFaceSet.delete".
func (this FontFaceSet) FuncDelete() (fn js.Func[func(font FontFace) bool]) {
	bindings.FuncFontFaceSetDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "FontFaceSet.delete".
func (this FontFaceSet) Delete(font FontFace) (ret bool) {
	bindings.CallFontFaceSetDelete(
		this.ref, js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryDelete calls the method "FontFaceSet.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryDelete(font FontFace) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "FontFaceSet.clear" exists.
func (this FontFaceSet) HasFuncClear() bool {
	return js.True == bindings.HasFuncFontFaceSetClear(
		this.ref,
	)
}

// FuncClear returns the method "FontFaceSet.clear".
func (this FontFaceSet) FuncClear() (fn js.Func[func()]) {
	bindings.FuncFontFaceSetClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "FontFaceSet.clear".
func (this FontFaceSet) Clear() (ret js.Void) {
	bindings.CallFontFaceSetClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "FontFaceSet.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLoad returns true if the method "FontFaceSet.load" exists.
func (this FontFaceSet) HasFuncLoad() bool {
	return js.True == bindings.HasFuncFontFaceSetLoad(
		this.ref,
	)
}

// FuncLoad returns the method "FontFaceSet.load".
func (this FontFaceSet) FuncLoad() (fn js.Func[func(font js.String, text js.String) js.Promise[js.Array[FontFace]]]) {
	bindings.FuncFontFaceSetLoad(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Load calls the method "FontFaceSet.load".
func (this FontFaceSet) Load(font js.String, text js.String) (ret js.Promise[js.Array[FontFace]]) {
	bindings.CallFontFaceSetLoad(
		this.ref, js.Pointer(&ret),
		font.Ref(),
		text.Ref(),
	)

	return
}

// TryLoad calls the method "FontFaceSet.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryLoad(font js.String, text js.String) (ret js.Promise[js.Array[FontFace]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetLoad(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
		text.Ref(),
	)

	return
}

// HasFuncLoad1 returns true if the method "FontFaceSet.load" exists.
func (this FontFaceSet) HasFuncLoad1() bool {
	return js.True == bindings.HasFuncFontFaceSetLoad1(
		this.ref,
	)
}

// FuncLoad1 returns the method "FontFaceSet.load".
func (this FontFaceSet) FuncLoad1() (fn js.Func[func(font js.String) js.Promise[js.Array[FontFace]]]) {
	bindings.FuncFontFaceSetLoad1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Load1 calls the method "FontFaceSet.load".
func (this FontFaceSet) Load1(font js.String) (ret js.Promise[js.Array[FontFace]]) {
	bindings.CallFontFaceSetLoad1(
		this.ref, js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryLoad1 calls the method "FontFaceSet.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryLoad1(font js.String) (ret js.Promise[js.Array[FontFace]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetLoad1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasFuncCheck returns true if the method "FontFaceSet.check" exists.
func (this FontFaceSet) HasFuncCheck() bool {
	return js.True == bindings.HasFuncFontFaceSetCheck(
		this.ref,
	)
}

// FuncCheck returns the method "FontFaceSet.check".
func (this FontFaceSet) FuncCheck() (fn js.Func[func(font js.String, text js.String) bool]) {
	bindings.FuncFontFaceSetCheck(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Check calls the method "FontFaceSet.check".
func (this FontFaceSet) Check(font js.String, text js.String) (ret bool) {
	bindings.CallFontFaceSetCheck(
		this.ref, js.Pointer(&ret),
		font.Ref(),
		text.Ref(),
	)

	return
}

// TryCheck calls the method "FontFaceSet.check"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryCheck(font js.String, text js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetCheck(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
		text.Ref(),
	)

	return
}

// HasFuncCheck1 returns true if the method "FontFaceSet.check" exists.
func (this FontFaceSet) HasFuncCheck1() bool {
	return js.True == bindings.HasFuncFontFaceSetCheck1(
		this.ref,
	)
}

// FuncCheck1 returns the method "FontFaceSet.check".
func (this FontFaceSet) FuncCheck1() (fn js.Func[func(font js.String) bool]) {
	bindings.FuncFontFaceSetCheck1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Check1 calls the method "FontFaceSet.check".
func (this FontFaceSet) Check1(font js.String) (ret bool) {
	bindings.CallFontFaceSetCheck1(
		this.ref, js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryCheck1 calls the method "FontFaceSet.check"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryCheck1(font js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetCheck1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

type Document struct {
	Node
}

func (this Document) Once() Document {
	this.ref.Once()
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
	this.ref.Free()
}

// Implementation returns the value of property "Document.implementation".
//
// It returns ok=false if there is no such property.
func (this Document) Implementation() (ret DOMImplementation, ok bool) {
	ok = js.True == bindings.GetDocumentImplementation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// URL returns the value of property "Document.URL".
//
// It returns ok=false if there is no such property.
func (this Document) URL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DocumentURI returns the value of property "Document.documentURI".
//
// It returns ok=false if there is no such property.
func (this Document) DocumentURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDocumentURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CompatMode returns the value of property "Document.compatMode".
//
// It returns ok=false if there is no such property.
func (this Document) CompatMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCompatMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CharacterSet returns the value of property "Document.characterSet".
//
// It returns ok=false if there is no such property.
func (this Document) CharacterSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCharacterSet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Charset returns the value of property "Document.charset".
//
// It returns ok=false if there is no such property.
func (this Document) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCharset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputEncoding returns the value of property "Document.inputEncoding".
//
// It returns ok=false if there is no such property.
func (this Document) InputEncoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentInputEncoding(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentType returns the value of property "Document.contentType".
//
// It returns ok=false if there is no such property.
func (this Document) ContentType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentContentType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Doctype returns the value of property "Document.doctype".
//
// It returns ok=false if there is no such property.
func (this Document) Doctype() (ret DocumentType, ok bool) {
	ok = js.True == bindings.GetDocumentDoctype(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DocumentElement returns the value of property "Document.documentElement".
//
// It returns ok=false if there is no such property.
func (this Document) DocumentElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentDocumentElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FragmentDirective returns the value of property "Document.fragmentDirective".
//
// It returns ok=false if there is no such property.
func (this Document) FragmentDirective() (ret FragmentDirective, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentDirective(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PermissionsPolicy returns the value of property "Document.permissionsPolicy".
//
// It returns ok=false if there is no such property.
func (this Document) PermissionsPolicy() (ret PermissionsPolicy, ok bool) {
	ok = js.True == bindings.GetDocumentPermissionsPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WasDiscarded returns the value of property "Document.wasDiscarded".
//
// It returns ok=false if there is no such property.
func (this Document) WasDiscarded() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentWasDiscarded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FullscreenEnabled returns the value of property "Document.fullscreenEnabled".
//
// It returns ok=false if there is no such property.
func (this Document) FullscreenEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreenEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Fullscreen returns the value of property "Document.fullscreen".
//
// It returns ok=false if there is no such property.
func (this Document) Fullscreen() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timeline returns the value of property "Document.timeline".
//
// It returns ok=false if there is no such property.
func (this Document) Timeline() (ret DocumentTimeline, ok bool) {
	ok = js.True == bindings.GetDocumentTimeline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PictureInPictureEnabled returns the value of property "Document.pictureInPictureEnabled".
//
// It returns ok=false if there is no such property.
func (this Document) PictureInPictureEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NamedFlows returns the value of property "Document.namedFlows".
//
// It returns ok=false if there is no such property.
func (this Document) NamedFlows() (ret NamedFlowMap, ok bool) {
	ok = js.True == bindings.GetDocumentNamedFlows(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScrollingElement returns the value of property "Document.scrollingElement".
//
// It returns ok=false if there is no such property.
func (this Document) ScrollingElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentScrollingElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RootElement returns the value of property "Document.rootElement".
//
// It returns ok=false if there is no such property.
func (this Document) RootElement() (ret SVGSVGElement, ok bool) {
	ok = js.True == bindings.GetDocumentRootElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Prerendering returns the value of property "Document.prerendering".
//
// It returns ok=false if there is no such property.
func (this Document) Prerendering() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentPrerendering(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FgColor returns the value of property "Document.fgColor".
//
// It returns ok=false if there is no such property.
func (this Document) FgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentFgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFgColor sets the value of property "Document.fgColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetFgColor(val js.String) bool {
	return js.True == bindings.SetDocumentFgColor(
		this.ref,
		val.Ref(),
	)
}

// LinkColor returns the value of property "Document.linkColor".
//
// It returns ok=false if there is no such property.
func (this Document) LinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentLinkColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLinkColor sets the value of property "Document.linkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetLinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentLinkColor(
		this.ref,
		val.Ref(),
	)
}

// VlinkColor returns the value of property "Document.vlinkColor".
//
// It returns ok=false if there is no such property.
func (this Document) VlinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentVlinkColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVlinkColor sets the value of property "Document.vlinkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetVlinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentVlinkColor(
		this.ref,
		val.Ref(),
	)
}

// AlinkColor returns the value of property "Document.alinkColor".
//
// It returns ok=false if there is no such property.
func (this Document) AlinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentAlinkColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlinkColor sets the value of property "Document.alinkColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetAlinkColor(val js.String) bool {
	return js.True == bindings.SetDocumentAlinkColor(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "Document.bgColor".
//
// It returns ok=false if there is no such property.
func (this Document) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "Document.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetBgColor(val js.String) bool {
	return js.True == bindings.SetDocumentBgColor(
		this.ref,
		val.Ref(),
	)
}

// Anchors returns the value of property "Document.anchors".
//
// It returns ok=false if there is no such property.
func (this Document) Anchors() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentAnchors(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Applets returns the value of property "Document.applets".
//
// It returns ok=false if there is no such property.
func (this Document) Applets() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentApplets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// All returns the value of property "Document.all".
//
// It returns ok=false if there is no such property.
func (this Document) All() (ret HTMLAllCollection, ok bool) {
	ok = js.True == bindings.GetDocumentAll(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Location returns the value of property "Document.location".
//
// It returns ok=false if there is no such property.
func (this Document) Location() (ret Location, ok bool) {
	ok = js.True == bindings.GetDocumentLocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Domain returns the value of property "Document.domain".
//
// It returns ok=false if there is no such property.
func (this Document) Domain() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDomain(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDomain sets the value of property "Document.domain" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDomain(val js.String) bool {
	return js.True == bindings.SetDocumentDomain(
		this.ref,
		val.Ref(),
	)
}

// Referrer returns the value of property "Document.referrer".
//
// It returns ok=false if there is no such property.
func (this Document) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentReferrer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cookie returns the value of property "Document.cookie".
//
// It returns ok=false if there is no such property.
func (this Document) Cookie() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCookie(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCookie sets the value of property "Document.cookie" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetCookie(val js.String) bool {
	return js.True == bindings.SetDocumentCookie(
		this.ref,
		val.Ref(),
	)
}

// LastModified returns the value of property "Document.lastModified".
//
// It returns ok=false if there is no such property.
func (this Document) LastModified() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentLastModified(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "Document.readyState".
//
// It returns ok=false if there is no such property.
func (this Document) ReadyState() (ret DocumentReadyState, ok bool) {
	ok = js.True == bindings.GetDocumentReadyState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Title returns the value of property "Document.title".
//
// It returns ok=false if there is no such property.
func (this Document) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTitle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "Document.title" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetTitle(val js.String) bool {
	return js.True == bindings.SetDocumentTitle(
		this.ref,
		val.Ref(),
	)
}

// Dir returns the value of property "Document.dir".
//
// It returns ok=false if there is no such property.
func (this Document) Dir() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDir(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDir sets the value of property "Document.dir" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDir(val js.String) bool {
	return js.True == bindings.SetDocumentDir(
		this.ref,
		val.Ref(),
	)
}

// Body returns the value of property "Document.body".
//
// It returns ok=false if there is no such property.
func (this Document) Body() (ret HTMLElement, ok bool) {
	ok = js.True == bindings.GetDocumentBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBody sets the value of property "Document.body" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetBody(val HTMLElement) bool {
	return js.True == bindings.SetDocumentBody(
		this.ref,
		val.Ref(),
	)
}

// Head returns the value of property "Document.head".
//
// It returns ok=false if there is no such property.
func (this Document) Head() (ret HTMLHeadElement, ok bool) {
	ok = js.True == bindings.GetDocumentHead(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Images returns the value of property "Document.images".
//
// It returns ok=false if there is no such property.
func (this Document) Images() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentImages(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Embeds returns the value of property "Document.embeds".
//
// It returns ok=false if there is no such property.
func (this Document) Embeds() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentEmbeds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Plugins returns the value of property "Document.plugins".
//
// It returns ok=false if there is no such property.
func (this Document) Plugins() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentPlugins(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Links returns the value of property "Document.links".
//
// It returns ok=false if there is no such property.
func (this Document) Links() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentLinks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Forms returns the value of property "Document.forms".
//
// It returns ok=false if there is no such property.
func (this Document) Forms() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentForms(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scripts returns the value of property "Document.scripts".
//
// It returns ok=false if there is no such property.
func (this Document) Scripts() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentScripts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentScript returns the value of property "Document.currentScript".
//
// It returns ok=false if there is no such property.
func (this Document) CurrentScript() (ret HTMLOrSVGScriptElement, ok bool) {
	ok = js.True == bindings.GetDocumentCurrentScript(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DefaultView returns the value of property "Document.defaultView".
//
// It returns ok=false if there is no such property.
func (this Document) DefaultView() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetDocumentDefaultView(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DesignMode returns the value of property "Document.designMode".
//
// It returns ok=false if there is no such property.
func (this Document) DesignMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDesignMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDesignMode sets the value of property "Document.designMode" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetDesignMode(val js.String) bool {
	return js.True == bindings.SetDocumentDesignMode(
		this.ref,
		val.Ref(),
	)
}

// Hidden returns the value of property "Document.hidden".
//
// It returns ok=false if there is no such property.
func (this Document) Hidden() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentHidden(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VisibilityState returns the value of property "Document.visibilityState".
//
// It returns ok=false if there is no such property.
func (this Document) VisibilityState() (ret DocumentVisibilityState, ok bool) {
	ok = js.True == bindings.GetDocumentVisibilityState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FullscreenElement returns the value of property "Document.fullscreenElement".
//
// It returns ok=false if there is no such property.
func (this Document) FullscreenElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreenElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StyleSheets returns the value of property "Document.styleSheets".
//
// It returns ok=false if there is no such property.
func (this Document) StyleSheets() (ret StyleSheetList, ok bool) {
	ok = js.True == bindings.GetDocumentStyleSheets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AdoptedStyleSheets returns the value of property "Document.adoptedStyleSheets".
//
// It returns ok=false if there is no such property.
func (this Document) AdoptedStyleSheets() (ret js.ObservableArray[CSSStyleSheet], ok bool) {
	ok = js.True == bindings.GetDocumentAdoptedStyleSheets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAdoptedStyleSheets sets the value of property "Document.adoptedStyleSheets" to val.
//
// It returns false if the property cannot be set.
func (this Document) SetAdoptedStyleSheets(val js.ObservableArray[CSSStyleSheet]) bool {
	return js.True == bindings.SetDocumentAdoptedStyleSheets(
		this.ref,
		val.Ref(),
	)
}

// PictureInPictureElement returns the value of property "Document.pictureInPictureElement".
//
// It returns ok=false if there is no such property.
func (this Document) PictureInPictureElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActiveElement returns the value of property "Document.activeElement".
//
// It returns ok=false if there is no such property.
func (this Document) ActiveElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentActiveElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PointerLockElement returns the value of property "Document.pointerLockElement".
//
// It returns ok=false if there is no such property.
func (this Document) PointerLockElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentPointerLockElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Children returns the value of property "Document.children".
//
// It returns ok=false if there is no such property.
func (this Document) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentChildren(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "Document.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this Document) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFirstElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "Document.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this Document) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentLastElementChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "Document.childElementCount".
//
// It returns ok=false if there is no such property.
func (this Document) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDocumentChildElementCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "Document.fonts".
//
// It returns ok=false if there is no such property.
func (this Document) Fonts() (ret FontFaceSet, ok bool) {
	ok = js.True == bindings.GetDocumentFonts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetElementsByTagName returns true if the method "Document.getElementsByTagName" exists.
func (this Document) HasFuncGetElementsByTagName() bool {
	return js.True == bindings.HasFuncDocumentGetElementsByTagName(
		this.ref,
	)
}

// FuncGetElementsByTagName returns the method "Document.getElementsByTagName".
func (this Document) FuncGetElementsByTagName() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	bindings.FuncDocumentGetElementsByTagName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByTagName calls the method "Document.getElementsByTagName".
func (this Document) GetElementsByTagName(qualifiedName js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByTagName(
		this.ref, js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetElementsByTagName calls the method "Document.getElementsByTagName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByTagName(qualifiedName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByTagName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncGetElementsByTagNameNS returns true if the method "Document.getElementsByTagNameNS" exists.
func (this Document) HasFuncGetElementsByTagNameNS() bool {
	return js.True == bindings.HasFuncDocumentGetElementsByTagNameNS(
		this.ref,
	)
}

// FuncGetElementsByTagNameNS returns the method "Document.getElementsByTagNameNS".
func (this Document) FuncGetElementsByTagNameNS() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	bindings.FuncDocumentGetElementsByTagNameNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByTagNameNS calls the method "Document.getElementsByTagNameNS".
func (this Document) GetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByTagNameNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetElementsByTagNameNS calls the method "Document.getElementsByTagNameNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByTagNameNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncGetElementsByClassName returns true if the method "Document.getElementsByClassName" exists.
func (this Document) HasFuncGetElementsByClassName() bool {
	return js.True == bindings.HasFuncDocumentGetElementsByClassName(
		this.ref,
	)
}

// FuncGetElementsByClassName returns the method "Document.getElementsByClassName".
func (this Document) FuncGetElementsByClassName() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	bindings.FuncDocumentGetElementsByClassName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByClassName calls the method "Document.getElementsByClassName".
func (this Document) GetElementsByClassName(classNames js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByClassName(
		this.ref, js.Pointer(&ret),
		classNames.Ref(),
	)

	return
}

// TryGetElementsByClassName calls the method "Document.getElementsByClassName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByClassName(classNames js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByClassName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		classNames.Ref(),
	)

	return
}

// HasFuncCreateElement returns true if the method "Document.createElement" exists.
func (this Document) HasFuncCreateElement() bool {
	return js.True == bindings.HasFuncDocumentCreateElement(
		this.ref,
	)
}

// FuncCreateElement returns the method "Document.createElement".
func (this Document) FuncCreateElement() (fn js.Func[func(localName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	bindings.FuncDocumentCreateElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateElement calls the method "Document.createElement".
func (this Document) CreateElement(localName js.String, options OneOf_String_ElementCreationOptions) (ret Element) {
	bindings.CallDocumentCreateElement(
		this.ref, js.Pointer(&ret),
		localName.Ref(),
		options.Ref(),
	)

	return
}

// TryCreateElement calls the method "Document.createElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateElement(localName js.String, options OneOf_String_ElementCreationOptions) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncCreateElement1 returns true if the method "Document.createElement" exists.
func (this Document) HasFuncCreateElement1() bool {
	return js.True == bindings.HasFuncDocumentCreateElement1(
		this.ref,
	)
}

// FuncCreateElement1 returns the method "Document.createElement".
func (this Document) FuncCreateElement1() (fn js.Func[func(localName js.String) Element]) {
	bindings.FuncDocumentCreateElement1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateElement1 calls the method "Document.createElement".
func (this Document) CreateElement1(localName js.String) (ret Element) {
	bindings.CallDocumentCreateElement1(
		this.ref, js.Pointer(&ret),
		localName.Ref(),
	)

	return
}

// TryCreateElement1 calls the method "Document.createElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateElement1(localName js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateElement1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
	)

	return
}

// HasFuncCreateElementNS returns true if the method "Document.createElementNS" exists.
func (this Document) HasFuncCreateElementNS() bool {
	return js.True == bindings.HasFuncDocumentCreateElementNS(
		this.ref,
	)
}

// FuncCreateElementNS returns the method "Document.createElementNS".
func (this Document) FuncCreateElementNS() (fn js.Func[func(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	bindings.FuncDocumentCreateElementNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateElementNS calls the method "Document.createElementNS".
func (this Document) CreateElementNS(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) (ret Element) {
	bindings.CallDocumentCreateElementNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
		options.Ref(),
	)

	return
}

// TryCreateElementNS calls the method "Document.createElementNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateElementNS(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateElementNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncCreateElementNS1 returns true if the method "Document.createElementNS" exists.
func (this Document) HasFuncCreateElementNS1() bool {
	return js.True == bindings.HasFuncDocumentCreateElementNS1(
		this.ref,
	)
}

// FuncCreateElementNS1 returns the method "Document.createElementNS".
func (this Document) FuncCreateElementNS1() (fn js.Func[func(namespace js.String, qualifiedName js.String) Element]) {
	bindings.FuncDocumentCreateElementNS1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateElementNS1 calls the method "Document.createElementNS".
func (this Document) CreateElementNS1(namespace js.String, qualifiedName js.String) (ret Element) {
	bindings.CallDocumentCreateElementNS1(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// TryCreateElementNS1 calls the method "Document.createElementNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateElementNS1(namespace js.String, qualifiedName js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateElementNS1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncCreateDocumentFragment returns true if the method "Document.createDocumentFragment" exists.
func (this Document) HasFuncCreateDocumentFragment() bool {
	return js.True == bindings.HasFuncDocumentCreateDocumentFragment(
		this.ref,
	)
}

// FuncCreateDocumentFragment returns the method "Document.createDocumentFragment".
func (this Document) FuncCreateDocumentFragment() (fn js.Func[func() DocumentFragment]) {
	bindings.FuncDocumentCreateDocumentFragment(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDocumentFragment calls the method "Document.createDocumentFragment".
func (this Document) CreateDocumentFragment() (ret DocumentFragment) {
	bindings.CallDocumentCreateDocumentFragment(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateDocumentFragment calls the method "Document.createDocumentFragment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateDocumentFragment() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateDocumentFragment(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateTextNode returns true if the method "Document.createTextNode" exists.
func (this Document) HasFuncCreateTextNode() bool {
	return js.True == bindings.HasFuncDocumentCreateTextNode(
		this.ref,
	)
}

// FuncCreateTextNode returns the method "Document.createTextNode".
func (this Document) FuncCreateTextNode() (fn js.Func[func(data js.String) Text]) {
	bindings.FuncDocumentCreateTextNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTextNode calls the method "Document.createTextNode".
func (this Document) CreateTextNode(data js.String) (ret Text) {
	bindings.CallDocumentCreateTextNode(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateTextNode calls the method "Document.createTextNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTextNode(data js.String) (ret Text, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTextNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncCreateCDATASection returns true if the method "Document.createCDATASection" exists.
func (this Document) HasFuncCreateCDATASection() bool {
	return js.True == bindings.HasFuncDocumentCreateCDATASection(
		this.ref,
	)
}

// FuncCreateCDATASection returns the method "Document.createCDATASection".
func (this Document) FuncCreateCDATASection() (fn js.Func[func(data js.String) CDATASection]) {
	bindings.FuncDocumentCreateCDATASection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCDATASection calls the method "Document.createCDATASection".
func (this Document) CreateCDATASection(data js.String) (ret CDATASection) {
	bindings.CallDocumentCreateCDATASection(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateCDATASection calls the method "Document.createCDATASection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateCDATASection(data js.String) (ret CDATASection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateCDATASection(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncCreateComment returns true if the method "Document.createComment" exists.
func (this Document) HasFuncCreateComment() bool {
	return js.True == bindings.HasFuncDocumentCreateComment(
		this.ref,
	)
}

// FuncCreateComment returns the method "Document.createComment".
func (this Document) FuncCreateComment() (fn js.Func[func(data js.String) Comment]) {
	bindings.FuncDocumentCreateComment(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateComment calls the method "Document.createComment".
func (this Document) CreateComment(data js.String) (ret Comment) {
	bindings.CallDocumentCreateComment(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateComment calls the method "Document.createComment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateComment(data js.String) (ret Comment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateComment(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncCreateProcessingInstruction returns true if the method "Document.createProcessingInstruction" exists.
func (this Document) HasFuncCreateProcessingInstruction() bool {
	return js.True == bindings.HasFuncDocumentCreateProcessingInstruction(
		this.ref,
	)
}

// FuncCreateProcessingInstruction returns the method "Document.createProcessingInstruction".
func (this Document) FuncCreateProcessingInstruction() (fn js.Func[func(target js.String, data js.String) ProcessingInstruction]) {
	bindings.FuncDocumentCreateProcessingInstruction(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateProcessingInstruction calls the method "Document.createProcessingInstruction".
func (this Document) CreateProcessingInstruction(target js.String, data js.String) (ret ProcessingInstruction) {
	bindings.CallDocumentCreateProcessingInstruction(
		this.ref, js.Pointer(&ret),
		target.Ref(),
		data.Ref(),
	)

	return
}

// TryCreateProcessingInstruction calls the method "Document.createProcessingInstruction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateProcessingInstruction(target js.String, data js.String) (ret ProcessingInstruction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateProcessingInstruction(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncImportNode returns true if the method "Document.importNode" exists.
func (this Document) HasFuncImportNode() bool {
	return js.True == bindings.HasFuncDocumentImportNode(
		this.ref,
	)
}

// FuncImportNode returns the method "Document.importNode".
func (this Document) FuncImportNode() (fn js.Func[func(node Node, deep bool) Node]) {
	bindings.FuncDocumentImportNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ImportNode calls the method "Document.importNode".
func (this Document) ImportNode(node Node, deep bool) (ret Node) {
	bindings.CallDocumentImportNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		js.Bool(bool(deep)),
	)

	return
}

// TryImportNode calls the method "Document.importNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryImportNode(node Node, deep bool) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentImportNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		js.Bool(bool(deep)),
	)

	return
}

// HasFuncImportNode1 returns true if the method "Document.importNode" exists.
func (this Document) HasFuncImportNode1() bool {
	return js.True == bindings.HasFuncDocumentImportNode1(
		this.ref,
	)
}

// FuncImportNode1 returns the method "Document.importNode".
func (this Document) FuncImportNode1() (fn js.Func[func(node Node) Node]) {
	bindings.FuncDocumentImportNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ImportNode1 calls the method "Document.importNode".
func (this Document) ImportNode1(node Node) (ret Node) {
	bindings.CallDocumentImportNode1(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryImportNode1 calls the method "Document.importNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryImportNode1(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentImportNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncAdoptNode returns true if the method "Document.adoptNode" exists.
func (this Document) HasFuncAdoptNode() bool {
	return js.True == bindings.HasFuncDocumentAdoptNode(
		this.ref,
	)
}

// FuncAdoptNode returns the method "Document.adoptNode".
func (this Document) FuncAdoptNode() (fn js.Func[func(node Node) Node]) {
	bindings.FuncDocumentAdoptNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AdoptNode calls the method "Document.adoptNode".
func (this Document) AdoptNode(node Node) (ret Node) {
	bindings.CallDocumentAdoptNode(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryAdoptNode calls the method "Document.adoptNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryAdoptNode(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentAdoptNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncCreateAttribute returns true if the method "Document.createAttribute" exists.
func (this Document) HasFuncCreateAttribute() bool {
	return js.True == bindings.HasFuncDocumentCreateAttribute(
		this.ref,
	)
}

// FuncCreateAttribute returns the method "Document.createAttribute".
func (this Document) FuncCreateAttribute() (fn js.Func[func(localName js.String) Attr]) {
	bindings.FuncDocumentCreateAttribute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateAttribute calls the method "Document.createAttribute".
func (this Document) CreateAttribute(localName js.String) (ret Attr) {
	bindings.CallDocumentCreateAttribute(
		this.ref, js.Pointer(&ret),
		localName.Ref(),
	)

	return
}

// TryCreateAttribute calls the method "Document.createAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateAttribute(localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateAttribute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
	)

	return
}

// HasFuncCreateAttributeNS returns true if the method "Document.createAttributeNS" exists.
func (this Document) HasFuncCreateAttributeNS() bool {
	return js.True == bindings.HasFuncDocumentCreateAttributeNS(
		this.ref,
	)
}

// FuncCreateAttributeNS returns the method "Document.createAttributeNS".
func (this Document) FuncCreateAttributeNS() (fn js.Func[func(namespace js.String, qualifiedName js.String) Attr]) {
	bindings.FuncDocumentCreateAttributeNS(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateAttributeNS calls the method "Document.createAttributeNS".
func (this Document) CreateAttributeNS(namespace js.String, qualifiedName js.String) (ret Attr) {
	bindings.CallDocumentCreateAttributeNS(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// TryCreateAttributeNS calls the method "Document.createAttributeNS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateAttributeNS(namespace js.String, qualifiedName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateAttributeNS(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasFuncCreateEvent returns true if the method "Document.createEvent" exists.
func (this Document) HasFuncCreateEvent() bool {
	return js.True == bindings.HasFuncDocumentCreateEvent(
		this.ref,
	)
}

// FuncCreateEvent returns the method "Document.createEvent".
func (this Document) FuncCreateEvent() (fn js.Func[func(iface js.String) Event]) {
	bindings.FuncDocumentCreateEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateEvent calls the method "Document.createEvent".
func (this Document) CreateEvent(iface js.String) (ret Event) {
	bindings.CallDocumentCreateEvent(
		this.ref, js.Pointer(&ret),
		iface.Ref(),
	)

	return
}

// TryCreateEvent calls the method "Document.createEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateEvent(iface js.String) (ret Event, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		iface.Ref(),
	)

	return
}

// HasFuncCreateRange returns true if the method "Document.createRange" exists.
func (this Document) HasFuncCreateRange() bool {
	return js.True == bindings.HasFuncDocumentCreateRange(
		this.ref,
	)
}

// FuncCreateRange returns the method "Document.createRange".
func (this Document) FuncCreateRange() (fn js.Func[func() Range]) {
	bindings.FuncDocumentCreateRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRange calls the method "Document.createRange".
func (this Document) CreateRange() (ret Range) {
	bindings.CallDocumentCreateRange(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateRange calls the method "Document.createRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateRange() (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateNodeIterator returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasFuncCreateNodeIterator() bool {
	return js.True == bindings.HasFuncDocumentCreateNodeIterator(
		this.ref,
	)
}

// FuncCreateNodeIterator returns the method "Document.createNodeIterator".
func (this Document) FuncCreateNodeIterator() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) NodeIterator]) {
	bindings.FuncDocumentCreateNodeIterator(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateNodeIterator calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator(
		this.ref, js.Pointer(&ret),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// TryCreateNodeIterator calls the method "Document.createNodeIterator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNodeIterator(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret NodeIterator, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNodeIterator(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// HasFuncCreateNodeIterator1 returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasFuncCreateNodeIterator1() bool {
	return js.True == bindings.HasFuncDocumentCreateNodeIterator1(
		this.ref,
	)
}

// FuncCreateNodeIterator1 returns the method "Document.createNodeIterator".
func (this Document) FuncCreateNodeIterator1() (fn js.Func[func(root Node, whatToShow uint32) NodeIterator]) {
	bindings.FuncDocumentCreateNodeIterator1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateNodeIterator1 calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator1(root Node, whatToShow uint32) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator1(
		this.ref, js.Pointer(&ret),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// TryCreateNodeIterator1 calls the method "Document.createNodeIterator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNodeIterator1(root Node, whatToShow uint32) (ret NodeIterator, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNodeIterator1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// HasFuncCreateNodeIterator2 returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasFuncCreateNodeIterator2() bool {
	return js.True == bindings.HasFuncDocumentCreateNodeIterator2(
		this.ref,
	)
}

// FuncCreateNodeIterator2 returns the method "Document.createNodeIterator".
func (this Document) FuncCreateNodeIterator2() (fn js.Func[func(root Node) NodeIterator]) {
	bindings.FuncDocumentCreateNodeIterator2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateNodeIterator2 calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator2(root Node) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator2(
		this.ref, js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryCreateNodeIterator2 calls the method "Document.createNodeIterator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNodeIterator2(root Node) (ret NodeIterator, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNodeIterator2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

// HasFuncCreateTreeWalker returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasFuncCreateTreeWalker() bool {
	return js.True == bindings.HasFuncDocumentCreateTreeWalker(
		this.ref,
	)
}

// FuncCreateTreeWalker returns the method "Document.createTreeWalker".
func (this Document) FuncCreateTreeWalker() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) TreeWalker]) {
	bindings.FuncDocumentCreateTreeWalker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTreeWalker calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker(
		this.ref, js.Pointer(&ret),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// TryCreateTreeWalker calls the method "Document.createTreeWalker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTreeWalker(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret TreeWalker, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTreeWalker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// HasFuncCreateTreeWalker1 returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasFuncCreateTreeWalker1() bool {
	return js.True == bindings.HasFuncDocumentCreateTreeWalker1(
		this.ref,
	)
}

// FuncCreateTreeWalker1 returns the method "Document.createTreeWalker".
func (this Document) FuncCreateTreeWalker1() (fn js.Func[func(root Node, whatToShow uint32) TreeWalker]) {
	bindings.FuncDocumentCreateTreeWalker1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTreeWalker1 calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker1(root Node, whatToShow uint32) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker1(
		this.ref, js.Pointer(&ret),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// TryCreateTreeWalker1 calls the method "Document.createTreeWalker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTreeWalker1(root Node, whatToShow uint32) (ret TreeWalker, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTreeWalker1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// HasFuncCreateTreeWalker2 returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasFuncCreateTreeWalker2() bool {
	return js.True == bindings.HasFuncDocumentCreateTreeWalker2(
		this.ref,
	)
}

// FuncCreateTreeWalker2 returns the method "Document.createTreeWalker".
func (this Document) FuncCreateTreeWalker2() (fn js.Func[func(root Node) TreeWalker]) {
	bindings.FuncDocumentCreateTreeWalker2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTreeWalker2 calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker2(root Node) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker2(
		this.ref, js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryCreateTreeWalker2 calls the method "Document.createTreeWalker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTreeWalker2(root Node) (ret TreeWalker, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTreeWalker2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

// HasFuncStartViewTransition returns true if the method "Document.startViewTransition" exists.
func (this Document) HasFuncStartViewTransition() bool {
	return js.True == bindings.HasFuncDocumentStartViewTransition(
		this.ref,
	)
}

// FuncStartViewTransition returns the method "Document.startViewTransition".
func (this Document) FuncStartViewTransition() (fn js.Func[func(updateCallback js.Func[func() js.Promise[js.Any]]) ViewTransition]) {
	bindings.FuncDocumentStartViewTransition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StartViewTransition calls the method "Document.startViewTransition".
func (this Document) StartViewTransition(updateCallback js.Func[func() js.Promise[js.Any]]) (ret ViewTransition) {
	bindings.CallDocumentStartViewTransition(
		this.ref, js.Pointer(&ret),
		updateCallback.Ref(),
	)

	return
}

// TryStartViewTransition calls the method "Document.startViewTransition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryStartViewTransition(updateCallback js.Func[func() js.Promise[js.Any]]) (ret ViewTransition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentStartViewTransition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		updateCallback.Ref(),
	)

	return
}

// HasFuncStartViewTransition1 returns true if the method "Document.startViewTransition" exists.
func (this Document) HasFuncStartViewTransition1() bool {
	return js.True == bindings.HasFuncDocumentStartViewTransition1(
		this.ref,
	)
}

// FuncStartViewTransition1 returns the method "Document.startViewTransition".
func (this Document) FuncStartViewTransition1() (fn js.Func[func() ViewTransition]) {
	bindings.FuncDocumentStartViewTransition1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StartViewTransition1 calls the method "Document.startViewTransition".
func (this Document) StartViewTransition1() (ret ViewTransition) {
	bindings.CallDocumentStartViewTransition1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStartViewTransition1 calls the method "Document.startViewTransition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryStartViewTransition1() (ret ViewTransition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentStartViewTransition1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMeasureElement returns true if the method "Document.measureElement" exists.
func (this Document) HasFuncMeasureElement() bool {
	return js.True == bindings.HasFuncDocumentMeasureElement(
		this.ref,
	)
}

// FuncMeasureElement returns the method "Document.measureElement".
func (this Document) FuncMeasureElement() (fn js.Func[func(element Element) FontMetrics]) {
	bindings.FuncDocumentMeasureElement(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MeasureElement calls the method "Document.measureElement".
func (this Document) MeasureElement(element Element) (ret FontMetrics) {
	bindings.CallDocumentMeasureElement(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryMeasureElement calls the method "Document.measureElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryMeasureElement(element Element) (ret FontMetrics, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentMeasureElement(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasFuncMeasureText returns true if the method "Document.measureText" exists.
func (this Document) HasFuncMeasureText() bool {
	return js.True == bindings.HasFuncDocumentMeasureText(
		this.ref,
	)
}

// FuncMeasureText returns the method "Document.measureText".
func (this Document) FuncMeasureText() (fn js.Func[func(text js.String, styleMap StylePropertyMapReadOnly) FontMetrics]) {
	bindings.FuncDocumentMeasureText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MeasureText calls the method "Document.measureText".
func (this Document) MeasureText(text js.String, styleMap StylePropertyMapReadOnly) (ret FontMetrics) {
	bindings.CallDocumentMeasureText(
		this.ref, js.Pointer(&ret),
		text.Ref(),
		styleMap.Ref(),
	)

	return
}

// TryMeasureText calls the method "Document.measureText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryMeasureText(text js.String, styleMap StylePropertyMapReadOnly) (ret FontMetrics, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentMeasureText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		styleMap.Ref(),
	)

	return
}

// HasFuncGetSelection returns true if the method "Document.getSelection" exists.
func (this Document) HasFuncGetSelection() bool {
	return js.True == bindings.HasFuncDocumentGetSelection(
		this.ref,
	)
}

// FuncGetSelection returns the method "Document.getSelection".
func (this Document) FuncGetSelection() (fn js.Func[func() Selection]) {
	bindings.FuncDocumentGetSelection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSelection calls the method "Document.getSelection".
func (this Document) GetSelection() (ret Selection) {
	bindings.CallDocumentGetSelection(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSelection calls the method "Document.getSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetSelection() (ret Selection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetSelection(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExitPointerLock returns true if the method "Document.exitPointerLock" exists.
func (this Document) HasFuncExitPointerLock() bool {
	return js.True == bindings.HasFuncDocumentExitPointerLock(
		this.ref,
	)
}

// FuncExitPointerLock returns the method "Document.exitPointerLock".
func (this Document) FuncExitPointerLock() (fn js.Func[func()]) {
	bindings.FuncDocumentExitPointerLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExitPointerLock calls the method "Document.exitPointerLock".
func (this Document) ExitPointerLock() (ret js.Void) {
	bindings.CallDocumentExitPointerLock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExitPointerLock calls the method "Document.exitPointerLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitPointerLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitPointerLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExitFullscreen returns true if the method "Document.exitFullscreen" exists.
func (this Document) HasFuncExitFullscreen() bool {
	return js.True == bindings.HasFuncDocumentExitFullscreen(
		this.ref,
	)
}

// FuncExitFullscreen returns the method "Document.exitFullscreen".
func (this Document) FuncExitFullscreen() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDocumentExitFullscreen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExitFullscreen calls the method "Document.exitFullscreen".
func (this Document) ExitFullscreen() (ret js.Promise[js.Void]) {
	bindings.CallDocumentExitFullscreen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExitFullscreen calls the method "Document.exitFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitFullscreen() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitFullscreen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHasStorageAccess returns true if the method "Document.hasStorageAccess" exists.
func (this Document) HasFuncHasStorageAccess() bool {
	return js.True == bindings.HasFuncDocumentHasStorageAccess(
		this.ref,
	)
}

// FuncHasStorageAccess returns the method "Document.hasStorageAccess".
func (this Document) FuncHasStorageAccess() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncDocumentHasStorageAccess(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasStorageAccess calls the method "Document.hasStorageAccess".
func (this Document) HasStorageAccess() (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasStorageAccess(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasStorageAccess calls the method "Document.hasStorageAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasStorageAccess() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasStorageAccess(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestStorageAccess returns true if the method "Document.requestStorageAccess" exists.
func (this Document) HasFuncRequestStorageAccess() bool {
	return js.True == bindings.HasFuncDocumentRequestStorageAccess(
		this.ref,
	)
}

// FuncRequestStorageAccess returns the method "Document.requestStorageAccess".
func (this Document) FuncRequestStorageAccess() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDocumentRequestStorageAccess(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestStorageAccess calls the method "Document.requestStorageAccess".
func (this Document) RequestStorageAccess() (ret js.Promise[js.Void]) {
	bindings.CallDocumentRequestStorageAccess(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestStorageAccess calls the method "Document.requestStorageAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryRequestStorageAccess() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentRequestStorageAccess(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExitPictureInPicture returns true if the method "Document.exitPictureInPicture" exists.
func (this Document) HasFuncExitPictureInPicture() bool {
	return js.True == bindings.HasFuncDocumentExitPictureInPicture(
		this.ref,
	)
}

// FuncExitPictureInPicture returns the method "Document.exitPictureInPicture".
func (this Document) FuncExitPictureInPicture() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDocumentExitPictureInPicture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExitPictureInPicture calls the method "Document.exitPictureInPicture".
func (this Document) ExitPictureInPicture() (ret js.Promise[js.Void]) {
	bindings.CallDocumentExitPictureInPicture(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExitPictureInPicture calls the method "Document.exitPictureInPicture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitPictureInPicture() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitPictureInPicture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHasPrivateTokens returns true if the method "Document.hasPrivateTokens" exists.
func (this Document) HasFuncHasPrivateTokens() bool {
	return js.True == bindings.HasFuncDocumentHasPrivateTokens(
		this.ref,
	)
}

// FuncHasPrivateTokens returns the method "Document.hasPrivateTokens".
func (this Document) FuncHasPrivateTokens() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	bindings.FuncDocumentHasPrivateTokens(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasPrivateTokens calls the method "Document.hasPrivateTokens".
func (this Document) HasPrivateTokens(issuer js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasPrivateTokens(
		this.ref, js.Pointer(&ret),
		issuer.Ref(),
	)

	return
}

// TryHasPrivateTokens calls the method "Document.hasPrivateTokens"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasPrivateTokens(issuer js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasPrivateTokens(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		issuer.Ref(),
	)

	return
}

// HasFuncHasRedemptionRecord returns true if the method "Document.hasRedemptionRecord" exists.
func (this Document) HasFuncHasRedemptionRecord() bool {
	return js.True == bindings.HasFuncDocumentHasRedemptionRecord(
		this.ref,
	)
}

// FuncHasRedemptionRecord returns the method "Document.hasRedemptionRecord".
func (this Document) FuncHasRedemptionRecord() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	bindings.FuncDocumentHasRedemptionRecord(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasRedemptionRecord calls the method "Document.hasRedemptionRecord".
func (this Document) HasRedemptionRecord(issuer js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasRedemptionRecord(
		this.ref, js.Pointer(&ret),
		issuer.Ref(),
	)

	return
}

// TryHasRedemptionRecord calls the method "Document.hasRedemptionRecord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasRedemptionRecord(issuer js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasRedemptionRecord(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		issuer.Ref(),
	)

	return
}

// HasFuncElementFromPoint returns true if the method "Document.elementFromPoint" exists.
func (this Document) HasFuncElementFromPoint() bool {
	return js.True == bindings.HasFuncDocumentElementFromPoint(
		this.ref,
	)
}

// FuncElementFromPoint returns the method "Document.elementFromPoint".
func (this Document) FuncElementFromPoint() (fn js.Func[func(x float64, y float64) Element]) {
	bindings.FuncDocumentElementFromPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ElementFromPoint calls the method "Document.elementFromPoint".
func (this Document) ElementFromPoint(x float64, y float64) (ret Element) {
	bindings.CallDocumentElementFromPoint(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryElementFromPoint calls the method "Document.elementFromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryElementFromPoint(x float64, y float64) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentElementFromPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncElementsFromPoint returns true if the method "Document.elementsFromPoint" exists.
func (this Document) HasFuncElementsFromPoint() bool {
	return js.True == bindings.HasFuncDocumentElementsFromPoint(
		this.ref,
	)
}

// FuncElementsFromPoint returns the method "Document.elementsFromPoint".
func (this Document) FuncElementsFromPoint() (fn js.Func[func(x float64, y float64) js.Array[Element]]) {
	bindings.FuncDocumentElementsFromPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ElementsFromPoint calls the method "Document.elementsFromPoint".
func (this Document) ElementsFromPoint(x float64, y float64) (ret js.Array[Element]) {
	bindings.CallDocumentElementsFromPoint(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryElementsFromPoint calls the method "Document.elementsFromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryElementsFromPoint(x float64, y float64) (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentElementsFromPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncCaretPositionFromPoint returns true if the method "Document.caretPositionFromPoint" exists.
func (this Document) HasFuncCaretPositionFromPoint() bool {
	return js.True == bindings.HasFuncDocumentCaretPositionFromPoint(
		this.ref,
	)
}

// FuncCaretPositionFromPoint returns the method "Document.caretPositionFromPoint".
func (this Document) FuncCaretPositionFromPoint() (fn js.Func[func(x float64, y float64) CaretPosition]) {
	bindings.FuncDocumentCaretPositionFromPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CaretPositionFromPoint calls the method "Document.caretPositionFromPoint".
func (this Document) CaretPositionFromPoint(x float64, y float64) (ret CaretPosition) {
	bindings.CallDocumentCaretPositionFromPoint(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryCaretPositionFromPoint calls the method "Document.caretPositionFromPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCaretPositionFromPoint(x float64, y float64) (ret CaretPosition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCaretPositionFromPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncClear returns true if the method "Document.clear" exists.
func (this Document) HasFuncClear() bool {
	return js.True == bindings.HasFuncDocumentClear(
		this.ref,
	)
}

// FuncClear returns the method "Document.clear".
func (this Document) FuncClear() (fn js.Func[func()]) {
	bindings.FuncDocumentClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "Document.clear".
func (this Document) Clear() (ret js.Void) {
	bindings.CallDocumentClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "Document.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCaptureEvents returns true if the method "Document.captureEvents" exists.
func (this Document) HasFuncCaptureEvents() bool {
	return js.True == bindings.HasFuncDocumentCaptureEvents(
		this.ref,
	)
}

// FuncCaptureEvents returns the method "Document.captureEvents".
func (this Document) FuncCaptureEvents() (fn js.Func[func()]) {
	bindings.FuncDocumentCaptureEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CaptureEvents calls the method "Document.captureEvents".
func (this Document) CaptureEvents() (ret js.Void) {
	bindings.CallDocumentCaptureEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCaptureEvents calls the method "Document.captureEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCaptureEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCaptureEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReleaseEvents returns true if the method "Document.releaseEvents" exists.
func (this Document) HasFuncReleaseEvents() bool {
	return js.True == bindings.HasFuncDocumentReleaseEvents(
		this.ref,
	)
}

// FuncReleaseEvents returns the method "Document.releaseEvents".
func (this Document) FuncReleaseEvents() (fn js.Func[func()]) {
	bindings.FuncDocumentReleaseEvents(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReleaseEvents calls the method "Document.releaseEvents".
func (this Document) ReleaseEvents() (ret js.Void) {
	bindings.CallDocumentReleaseEvents(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReleaseEvents calls the method "Document.releaseEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryReleaseEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentReleaseEvents(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestStorageAccessFor returns true if the method "Document.requestStorageAccessFor" exists.
func (this Document) HasFuncRequestStorageAccessFor() bool {
	return js.True == bindings.HasFuncDocumentRequestStorageAccessFor(
		this.ref,
	)
}

// FuncRequestStorageAccessFor returns the method "Document.requestStorageAccessFor".
func (this Document) FuncRequestStorageAccessFor() (fn js.Func[func(requestedOrigin js.String) js.Promise[js.Void]]) {
	bindings.FuncDocumentRequestStorageAccessFor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestStorageAccessFor calls the method "Document.requestStorageAccessFor".
func (this Document) RequestStorageAccessFor(requestedOrigin js.String) (ret js.Promise[js.Void]) {
	bindings.CallDocumentRequestStorageAccessFor(
		this.ref, js.Pointer(&ret),
		requestedOrigin.Ref(),
	)

	return
}

// TryRequestStorageAccessFor calls the method "Document.requestStorageAccessFor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryRequestStorageAccessFor(requestedOrigin js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentRequestStorageAccessFor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		requestedOrigin.Ref(),
	)

	return
}

// HasFuncGet returns true if the method "Document." exists.
func (this Document) HasFuncGet() bool {
	return js.True == bindings.HasFuncDocumentGet(
		this.ref,
	)
}

// FuncGet returns the method "Document.".
func (this Document) FuncGet() (fn js.Func[func(name js.String) js.Object]) {
	bindings.FuncDocumentGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "Document.".
func (this Document) Get(name js.String) (ret js.Object) {
	bindings.CallDocumentGet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "Document."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGet(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetElementsByName returns true if the method "Document.getElementsByName" exists.
func (this Document) HasFuncGetElementsByName() bool {
	return js.True == bindings.HasFuncDocumentGetElementsByName(
		this.ref,
	)
}

// FuncGetElementsByName returns the method "Document.getElementsByName".
func (this Document) FuncGetElementsByName() (fn js.Func[func(elementName js.String) NodeList]) {
	bindings.FuncDocumentGetElementsByName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementsByName calls the method "Document.getElementsByName".
func (this Document) GetElementsByName(elementName js.String) (ret NodeList) {
	bindings.CallDocumentGetElementsByName(
		this.ref, js.Pointer(&ret),
		elementName.Ref(),
	)

	return
}

// TryGetElementsByName calls the method "Document.getElementsByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByName(elementName js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elementName.Ref(),
	)

	return
}

// HasFuncOpen returns true if the method "Document.open" exists.
func (this Document) HasFuncOpen() bool {
	return js.True == bindings.HasFuncDocumentOpen(
		this.ref,
	)
}

// FuncOpen returns the method "Document.open".
func (this Document) FuncOpen() (fn js.Func[func(unused1 js.String, unused2 js.String) Document]) {
	bindings.FuncDocumentOpen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open calls the method "Document.open".
func (this Document) Open(unused1 js.String, unused2 js.String) (ret Document) {
	bindings.CallDocumentOpen(
		this.ref, js.Pointer(&ret),
		unused1.Ref(),
		unused2.Ref(),
	)

	return
}

// TryOpen calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen(unused1 js.String, unused2 js.String) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		unused1.Ref(),
		unused2.Ref(),
	)

	return
}

// HasFuncOpen1 returns true if the method "Document.open" exists.
func (this Document) HasFuncOpen1() bool {
	return js.True == bindings.HasFuncDocumentOpen1(
		this.ref,
	)
}

// FuncOpen1 returns the method "Document.open".
func (this Document) FuncOpen1() (fn js.Func[func(unused1 js.String) Document]) {
	bindings.FuncDocumentOpen1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open1 calls the method "Document.open".
func (this Document) Open1(unused1 js.String) (ret Document) {
	bindings.CallDocumentOpen1(
		this.ref, js.Pointer(&ret),
		unused1.Ref(),
	)

	return
}

// TryOpen1 calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen1(unused1 js.String) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		unused1.Ref(),
	)

	return
}

// HasFuncOpen2 returns true if the method "Document.open" exists.
func (this Document) HasFuncOpen2() bool {
	return js.True == bindings.HasFuncDocumentOpen2(
		this.ref,
	)
}

// FuncOpen2 returns the method "Document.open".
func (this Document) FuncOpen2() (fn js.Func[func() Document]) {
	bindings.FuncDocumentOpen2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open2 calls the method "Document.open".
func (this Document) Open2() (ret Document) {
	bindings.CallDocumentOpen2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpen2 calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen2() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpen3 returns true if the method "Document.open" exists.
func (this Document) HasFuncOpen3() bool {
	return js.True == bindings.HasFuncDocumentOpen3(
		this.ref,
	)
}

// FuncOpen3 returns the method "Document.open".
func (this Document) FuncOpen3() (fn js.Func[func(url js.String, name js.String, features js.String) js.Object]) {
	bindings.FuncDocumentOpen3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Open3 calls the method "Document.open".
func (this Document) Open3(url js.String, name js.String, features js.String) (ret js.Object) {
	bindings.CallDocumentOpen3(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		name.Ref(),
		features.Ref(),
	)

	return
}

// TryOpen3 calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen3(url js.String, name js.String, features js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		name.Ref(),
		features.Ref(),
	)

	return
}

// HasFuncClose returns true if the method "Document.close" exists.
func (this Document) HasFuncClose() bool {
	return js.True == bindings.HasFuncDocumentClose(
		this.ref,
	)
}

// FuncClose returns the method "Document.close".
func (this Document) FuncClose() (fn js.Func[func()]) {
	bindings.FuncDocumentClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "Document.close".
func (this Document) Close() (ret js.Void) {
	bindings.CallDocumentClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "Document.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWrite returns true if the method "Document.write" exists.
func (this Document) HasFuncWrite() bool {
	return js.True == bindings.HasFuncDocumentWrite(
		this.ref,
	)
}

// FuncWrite returns the method "Document.write".
func (this Document) FuncWrite() (fn js.Func[func(text ...js.String)]) {
	bindings.FuncDocumentWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "Document.write".
func (this Document) Write(text ...js.String) (ret js.Void) {
	bindings.CallDocumentWrite(
		this.ref, js.Pointer(&ret),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// TryWrite calls the method "Document.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryWrite(text ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// HasFuncWriteln returns true if the method "Document.writeln" exists.
func (this Document) HasFuncWriteln() bool {
	return js.True == bindings.HasFuncDocumentWriteln(
		this.ref,
	)
}

// FuncWriteln returns the method "Document.writeln".
func (this Document) FuncWriteln() (fn js.Func[func(text ...js.String)]) {
	bindings.FuncDocumentWriteln(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Writeln calls the method "Document.writeln".
func (this Document) Writeln(text ...js.String) (ret js.Void) {
	bindings.CallDocumentWriteln(
		this.ref, js.Pointer(&ret),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// TryWriteln calls the method "Document.writeln"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryWriteln(text ...js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentWriteln(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// HasFuncHasFocus returns true if the method "Document.hasFocus" exists.
func (this Document) HasFuncHasFocus() bool {
	return js.True == bindings.HasFuncDocumentHasFocus(
		this.ref,
	)
}

// FuncHasFocus returns the method "Document.hasFocus".
func (this Document) FuncHasFocus() (fn js.Func[func() bool]) {
	bindings.FuncDocumentHasFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasFocus calls the method "Document.hasFocus".
func (this Document) HasFocus() (ret bool) {
	bindings.CallDocumentHasFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasFocus calls the method "Document.hasFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasFocus() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExecCommand returns true if the method "Document.execCommand" exists.
func (this Document) HasFuncExecCommand() bool {
	return js.True == bindings.HasFuncDocumentExecCommand(
		this.ref,
	)
}

// FuncExecCommand returns the method "Document.execCommand".
func (this Document) FuncExecCommand() (fn js.Func[func(commandId js.String, showUI bool, value js.String) bool]) {
	bindings.FuncDocumentExecCommand(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExecCommand calls the method "Document.execCommand".
func (this Document) ExecCommand(commandId js.String, showUI bool, value js.String) (ret bool) {
	bindings.CallDocumentExecCommand(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
		js.Bool(bool(showUI)),
		value.Ref(),
	)

	return
}

// TryExecCommand calls the method "Document.execCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExecCommand(commandId js.String, showUI bool, value js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExecCommand(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
		js.Bool(bool(showUI)),
		value.Ref(),
	)

	return
}

// HasFuncExecCommand1 returns true if the method "Document.execCommand" exists.
func (this Document) HasFuncExecCommand1() bool {
	return js.True == bindings.HasFuncDocumentExecCommand1(
		this.ref,
	)
}

// FuncExecCommand1 returns the method "Document.execCommand".
func (this Document) FuncExecCommand1() (fn js.Func[func(commandId js.String, showUI bool) bool]) {
	bindings.FuncDocumentExecCommand1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExecCommand1 calls the method "Document.execCommand".
func (this Document) ExecCommand1(commandId js.String, showUI bool) (ret bool) {
	bindings.CallDocumentExecCommand1(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
		js.Bool(bool(showUI)),
	)

	return
}

// TryExecCommand1 calls the method "Document.execCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExecCommand1(commandId js.String, showUI bool) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExecCommand1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
		js.Bool(bool(showUI)),
	)

	return
}

// HasFuncExecCommand2 returns true if the method "Document.execCommand" exists.
func (this Document) HasFuncExecCommand2() bool {
	return js.True == bindings.HasFuncDocumentExecCommand2(
		this.ref,
	)
}

// FuncExecCommand2 returns the method "Document.execCommand".
func (this Document) FuncExecCommand2() (fn js.Func[func(commandId js.String) bool]) {
	bindings.FuncDocumentExecCommand2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExecCommand2 calls the method "Document.execCommand".
func (this Document) ExecCommand2(commandId js.String) (ret bool) {
	bindings.CallDocumentExecCommand2(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryExecCommand2 calls the method "Document.execCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExecCommand2(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExecCommand2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncQueryCommandEnabled returns true if the method "Document.queryCommandEnabled" exists.
func (this Document) HasFuncQueryCommandEnabled() bool {
	return js.True == bindings.HasFuncDocumentQueryCommandEnabled(
		this.ref,
	)
}

// FuncQueryCommandEnabled returns the method "Document.queryCommandEnabled".
func (this Document) FuncQueryCommandEnabled() (fn js.Func[func(commandId js.String) bool]) {
	bindings.FuncDocumentQueryCommandEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCommandEnabled calls the method "Document.queryCommandEnabled".
func (this Document) QueryCommandEnabled(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandEnabled(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandEnabled calls the method "Document.queryCommandEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandEnabled(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncQueryCommandIndeterm returns true if the method "Document.queryCommandIndeterm" exists.
func (this Document) HasFuncQueryCommandIndeterm() bool {
	return js.True == bindings.HasFuncDocumentQueryCommandIndeterm(
		this.ref,
	)
}

// FuncQueryCommandIndeterm returns the method "Document.queryCommandIndeterm".
func (this Document) FuncQueryCommandIndeterm() (fn js.Func[func(commandId js.String) bool]) {
	bindings.FuncDocumentQueryCommandIndeterm(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCommandIndeterm calls the method "Document.queryCommandIndeterm".
func (this Document) QueryCommandIndeterm(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandIndeterm(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandIndeterm calls the method "Document.queryCommandIndeterm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandIndeterm(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandIndeterm(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncQueryCommandState returns true if the method "Document.queryCommandState" exists.
func (this Document) HasFuncQueryCommandState() bool {
	return js.True == bindings.HasFuncDocumentQueryCommandState(
		this.ref,
	)
}

// FuncQueryCommandState returns the method "Document.queryCommandState".
func (this Document) FuncQueryCommandState() (fn js.Func[func(commandId js.String) bool]) {
	bindings.FuncDocumentQueryCommandState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCommandState calls the method "Document.queryCommandState".
func (this Document) QueryCommandState(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandState(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandState calls the method "Document.queryCommandState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandState(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncQueryCommandSupported returns true if the method "Document.queryCommandSupported" exists.
func (this Document) HasFuncQueryCommandSupported() bool {
	return js.True == bindings.HasFuncDocumentQueryCommandSupported(
		this.ref,
	)
}

// FuncQueryCommandSupported returns the method "Document.queryCommandSupported".
func (this Document) FuncQueryCommandSupported() (fn js.Func[func(commandId js.String) bool]) {
	bindings.FuncDocumentQueryCommandSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCommandSupported calls the method "Document.queryCommandSupported".
func (this Document) QueryCommandSupported(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandSupported(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandSupported calls the method "Document.queryCommandSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandSupported(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncQueryCommandValue returns true if the method "Document.queryCommandValue" exists.
func (this Document) HasFuncQueryCommandValue() bool {
	return js.True == bindings.HasFuncDocumentQueryCommandValue(
		this.ref,
	)
}

// FuncQueryCommandValue returns the method "Document.queryCommandValue".
func (this Document) FuncQueryCommandValue() (fn js.Func[func(commandId js.String) js.String]) {
	bindings.FuncDocumentQueryCommandValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCommandValue calls the method "Document.queryCommandValue".
func (this Document) QueryCommandValue(commandId js.String) (ret js.String) {
	bindings.CallDocumentQueryCommandValue(
		this.ref, js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandValue calls the method "Document.queryCommandValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandValue(commandId js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasFuncGetBoxQuads returns true if the method "Document.getBoxQuads" exists.
func (this Document) HasFuncGetBoxQuads() bool {
	return js.True == bindings.HasFuncDocumentGetBoxQuads(
		this.ref,
	)
}

// FuncGetBoxQuads returns the method "Document.getBoxQuads".
func (this Document) FuncGetBoxQuads() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	bindings.FuncDocumentGetBoxQuads(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads calls the method "Document.getBoxQuads".
func (this Document) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallDocumentGetBoxQuads(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Document.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetBoxQuads(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetBoxQuads1 returns true if the method "Document.getBoxQuads" exists.
func (this Document) HasFuncGetBoxQuads1() bool {
	return js.True == bindings.HasFuncDocumentGetBoxQuads1(
		this.ref,
	)
}

// FuncGetBoxQuads1 returns the method "Document.getBoxQuads".
func (this Document) FuncGetBoxQuads1() (fn js.Func[func() js.Array[DOMQuad]]) {
	bindings.FuncDocumentGetBoxQuads1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBoxQuads1 calls the method "Document.getBoxQuads".
func (this Document) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallDocumentGetBoxQuads1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Document.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetBoxQuads1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvertQuadFromNode returns true if the method "Document.convertQuadFromNode" exists.
func (this Document) HasFuncConvertQuadFromNode() bool {
	return js.True == bindings.HasFuncDocumentConvertQuadFromNode(
		this.ref,
	)
}

// FuncConvertQuadFromNode returns the method "Document.convertQuadFromNode".
func (this Document) FuncConvertQuadFromNode() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncDocumentConvertQuadFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode calls the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallDocumentConvertQuadFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertQuadFromNode calls the method "Document.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertQuadFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertQuadFromNode1 returns true if the method "Document.convertQuadFromNode" exists.
func (this Document) HasFuncConvertQuadFromNode1() bool {
	return js.True == bindings.HasFuncDocumentConvertQuadFromNode1(
		this.ref,
	)
}

// FuncConvertQuadFromNode1 returns the method "Document.convertQuadFromNode".
func (this Document) FuncConvertQuadFromNode1() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	bindings.FuncDocumentConvertQuadFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertQuadFromNode1 calls the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallDocumentConvertQuadFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// TryConvertQuadFromNode1 calls the method "Document.convertQuadFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertQuadFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasFuncConvertRectFromNode returns true if the method "Document.convertRectFromNode" exists.
func (this Document) HasFuncConvertRectFromNode() bool {
	return js.True == bindings.HasFuncDocumentConvertRectFromNode(
		this.ref,
	)
}

// FuncConvertRectFromNode returns the method "Document.convertRectFromNode".
func (this Document) FuncConvertRectFromNode() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	bindings.FuncDocumentConvertRectFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode calls the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallDocumentConvertRectFromNode(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertRectFromNode calls the method "Document.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertRectFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertRectFromNode1 returns true if the method "Document.convertRectFromNode" exists.
func (this Document) HasFuncConvertRectFromNode1() bool {
	return js.True == bindings.HasFuncDocumentConvertRectFromNode1(
		this.ref,
	)
}

// FuncConvertRectFromNode1 returns the method "Document.convertRectFromNode".
func (this Document) FuncConvertRectFromNode1() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	bindings.FuncDocumentConvertRectFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertRectFromNode1 calls the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallDocumentConvertRectFromNode1(
		this.ref, js.Pointer(&ret),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// TryConvertRectFromNode1 calls the method "Document.convertRectFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertRectFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasFuncConvertPointFromNode returns true if the method "Document.convertPointFromNode" exists.
func (this Document) HasFuncConvertPointFromNode() bool {
	return js.True == bindings.HasFuncDocumentConvertPointFromNode(
		this.ref,
	)
}

// FuncConvertPointFromNode returns the method "Document.convertPointFromNode".
func (this Document) FuncConvertPointFromNode() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	bindings.FuncDocumentConvertPointFromNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode calls the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallDocumentConvertPointFromNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvertPointFromNode calls the method "Document.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertPointFromNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertPointFromNode1 returns true if the method "Document.convertPointFromNode" exists.
func (this Document) HasFuncConvertPointFromNode1() bool {
	return js.True == bindings.HasFuncDocumentConvertPointFromNode1(
		this.ref,
	)
}

// FuncConvertPointFromNode1 returns the method "Document.convertPointFromNode".
func (this Document) FuncConvertPointFromNode1() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	bindings.FuncDocumentConvertPointFromNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertPointFromNode1 calls the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallDocumentConvertPointFromNode1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// TryConvertPointFromNode1 calls the method "Document.convertPointFromNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentConvertPointFromNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// HasFuncGetElementById returns true if the method "Document.getElementById" exists.
func (this Document) HasFuncGetElementById() bool {
	return js.True == bindings.HasFuncDocumentGetElementById(
		this.ref,
	)
}

// FuncGetElementById returns the method "Document.getElementById".
func (this Document) FuncGetElementById() (fn js.Func[func(elementId js.String) Element]) {
	bindings.FuncDocumentGetElementById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetElementById calls the method "Document.getElementById".
func (this Document) GetElementById(elementId js.String) (ret Element) {
	bindings.CallDocumentGetElementById(
		this.ref, js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "Document.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasFuncGetAnimations returns true if the method "Document.getAnimations" exists.
func (this Document) HasFuncGetAnimations() bool {
	return js.True == bindings.HasFuncDocumentGetAnimations(
		this.ref,
	)
}

// FuncGetAnimations returns the method "Document.getAnimations".
func (this Document) FuncGetAnimations() (fn js.Func[func() js.Array[Animation]]) {
	bindings.FuncDocumentGetAnimations(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAnimations calls the method "Document.getAnimations".
func (this Document) GetAnimations() (ret js.Array[Animation]) {
	bindings.CallDocumentGetAnimations(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAnimations calls the method "Document.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetAnimations() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetAnimations(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPrepend returns true if the method "Document.prepend" exists.
func (this Document) HasFuncPrepend() bool {
	return js.True == bindings.HasFuncDocumentPrepend(
		this.ref,
	)
}

// FuncPrepend returns the method "Document.prepend".
func (this Document) FuncPrepend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentPrepend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prepend calls the method "Document.prepend".
func (this Document) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentPrepend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryPrepend calls the method "Document.prepend"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryPrepend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentPrepend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncAppend returns true if the method "Document.append" exists.
func (this Document) HasFuncAppend() bool {
	return js.True == bindings.HasFuncDocumentAppend(
		this.ref,
	)
}

// FuncAppend returns the method "Document.append".
func (this Document) FuncAppend() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "Document.append".
func (this Document) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentAppend(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryAppend calls the method "Document.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryAppend(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncReplaceChildren returns true if the method "Document.replaceChildren" exists.
func (this Document) HasFuncReplaceChildren() bool {
	return js.True == bindings.HasFuncDocumentReplaceChildren(
		this.ref,
	)
}

// FuncReplaceChildren returns the method "Document.replaceChildren".
func (this Document) FuncReplaceChildren() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	bindings.FuncDocumentReplaceChildren(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceChildren calls the method "Document.replaceChildren".
func (this Document) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentReplaceChildren(
		this.ref, js.Pointer(&ret),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// TryReplaceChildren calls the method "Document.replaceChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentReplaceChildren(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasFuncQuerySelector returns true if the method "Document.querySelector" exists.
func (this Document) HasFuncQuerySelector() bool {
	return js.True == bindings.HasFuncDocumentQuerySelector(
		this.ref,
	)
}

// FuncQuerySelector returns the method "Document.querySelector".
func (this Document) FuncQuerySelector() (fn js.Func[func(selectors js.String) Element]) {
	bindings.FuncDocumentQuerySelector(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelector calls the method "Document.querySelector".
func (this Document) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallDocumentQuerySelector(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "Document.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQuerySelector(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncQuerySelectorAll returns true if the method "Document.querySelectorAll" exists.
func (this Document) HasFuncQuerySelectorAll() bool {
	return js.True == bindings.HasFuncDocumentQuerySelectorAll(
		this.ref,
	)
}

// FuncQuerySelectorAll returns the method "Document.querySelectorAll".
func (this Document) FuncQuerySelectorAll() (fn js.Func[func(selectors js.String) NodeList]) {
	bindings.FuncDocumentQuerySelectorAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuerySelectorAll calls the method "Document.querySelectorAll".
func (this Document) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallDocumentQuerySelectorAll(
		this.ref, js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "Document.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQuerySelectorAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasFuncCreateExpression returns true if the method "Document.createExpression" exists.
func (this Document) HasFuncCreateExpression() bool {
	return js.True == bindings.HasFuncDocumentCreateExpression(
		this.ref,
	)
}

// FuncCreateExpression returns the method "Document.createExpression".
func (this Document) FuncCreateExpression() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	bindings.FuncDocumentCreateExpression(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateExpression calls the method "Document.createExpression".
func (this Document) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression) {
	bindings.CallDocumentCreateExpression(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// TryCreateExpression calls the method "Document.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateExpression(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// HasFuncCreateExpression1 returns true if the method "Document.createExpression" exists.
func (this Document) HasFuncCreateExpression1() bool {
	return js.True == bindings.HasFuncDocumentCreateExpression1(
		this.ref,
	)
}

// FuncCreateExpression1 returns the method "Document.createExpression".
func (this Document) FuncCreateExpression1() (fn js.Func[func(expression js.String) XPathExpression]) {
	bindings.FuncDocumentCreateExpression1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateExpression1 calls the method "Document.createExpression".
func (this Document) CreateExpression1(expression js.String) (ret XPathExpression) {
	bindings.CallDocumentCreateExpression1(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
	)

	return
}

// TryCreateExpression1 calls the method "Document.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateExpression1(expression js.String) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateExpression1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
	)

	return
}

// HasFuncCreateNSResolver returns true if the method "Document.createNSResolver" exists.
func (this Document) HasFuncCreateNSResolver() bool {
	return js.True == bindings.HasFuncDocumentCreateNSResolver(
		this.ref,
	)
}

// FuncCreateNSResolver returns the method "Document.createNSResolver".
func (this Document) FuncCreateNSResolver() (fn js.Func[func(nodeResolver Node) Node]) {
	bindings.FuncDocumentCreateNSResolver(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateNSResolver calls the method "Document.createNSResolver".
func (this Document) CreateNSResolver(nodeResolver Node) (ret Node) {
	bindings.CallDocumentCreateNSResolver(
		this.ref, js.Pointer(&ret),
		nodeResolver.Ref(),
	)

	return
}

// TryCreateNSResolver calls the method "Document.createNSResolver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNSResolver(nodeResolver Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNSResolver(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		nodeResolver.Ref(),
	)

	return
}

// HasFuncEvaluate returns true if the method "Document.evaluate" exists.
func (this Document) HasFuncEvaluate() bool {
	return js.True == bindings.HasFuncDocumentEvaluate(
		this.ref,
	)
}

// FuncEvaluate returns the method "Document.evaluate".
func (this Document) FuncEvaluate() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	bindings.FuncDocumentEvaluate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate calls the method "Document.evaluate".
func (this Document) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallDocumentEvaluate(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// TryEvaluate calls the method "Document.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryEvaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentEvaluate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasFuncEvaluate1 returns true if the method "Document.evaluate" exists.
func (this Document) HasFuncEvaluate1() bool {
	return js.True == bindings.HasFuncDocumentEvaluate1(
		this.ref,
	)
}

// FuncEvaluate1 returns the method "Document.evaluate".
func (this Document) FuncEvaluate1() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	bindings.FuncDocumentEvaluate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate1 calls the method "Document.evaluate".
func (this Document) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult) {
	bindings.CallDocumentEvaluate1(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// TryEvaluate1 calls the method "Document.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryEvaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentEvaluate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// HasFuncEvaluate2 returns true if the method "Document.evaluate" exists.
func (this Document) HasFuncEvaluate2() bool {
	return js.True == bindings.HasFuncDocumentEvaluate2(
		this.ref,
	)
}

// FuncEvaluate2 returns the method "Document.evaluate".
func (this Document) FuncEvaluate2() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	bindings.FuncDocumentEvaluate2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate2 calls the method "Document.evaluate".
func (this Document) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult) {
	bindings.CallDocumentEvaluate2(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// TryEvaluate2 calls the method "Document.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryEvaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentEvaluate2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// HasFuncEvaluate3 returns true if the method "Document.evaluate" exists.
func (this Document) HasFuncEvaluate3() bool {
	return js.True == bindings.HasFuncDocumentEvaluate3(
		this.ref,
	)
}

// FuncEvaluate3 returns the method "Document.evaluate".
func (this Document) FuncEvaluate3() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	bindings.FuncDocumentEvaluate3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Evaluate3 calls the method "Document.evaluate".
func (this Document) Evaluate3(expression js.String, contextNode Node) (ret XPathResult) {
	bindings.CallDocumentEvaluate3(
		this.ref, js.Pointer(&ret),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
}

// TryEvaluate3 calls the method "Document.evaluate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryEvaluate3(expression js.String, contextNode Node) (ret XPathResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentEvaluate3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
}

type Node struct {
	EventTarget
}

func (this Node) Once() Node {
	this.ref.Once()
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
	this.ref.Free()
}

// NodeType returns the value of property "Node.nodeType".
//
// It returns ok=false if there is no such property.
func (this Node) NodeType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetNodeNodeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NodeName returns the value of property "Node.nodeName".
//
// It returns ok=false if there is no such property.
func (this Node) NodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeNodeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BaseURI returns the value of property "Node.baseURI".
//
// It returns ok=false if there is no such property.
func (this Node) BaseURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeBaseURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsConnected returns the value of property "Node.isConnected".
//
// It returns ok=false if there is no such property.
func (this Node) IsConnected() (ret bool, ok bool) {
	ok = js.True == bindings.GetNodeIsConnected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OwnerDocument returns the value of property "Node.ownerDocument".
//
// It returns ok=false if there is no such property.
func (this Node) OwnerDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetNodeOwnerDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ParentNode returns the value of property "Node.parentNode".
//
// It returns ok=false if there is no such property.
func (this Node) ParentNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeParentNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ParentElement returns the value of property "Node.parentElement".
//
// It returns ok=false if there is no such property.
func (this Node) ParentElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetNodeParentElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChildNodes returns the value of property "Node.childNodes".
//
// It returns ok=false if there is no such property.
func (this Node) ChildNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetNodeChildNodes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstChild returns the value of property "Node.firstChild".
//
// It returns ok=false if there is no such property.
func (this Node) FirstChild() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeFirstChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastChild returns the value of property "Node.lastChild".
//
// It returns ok=false if there is no such property.
func (this Node) LastChild() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeLastChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "Node.previousSibling".
//
// It returns ok=false if there is no such property.
func (this Node) PreviousSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodePreviousSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "Node.nextSibling".
//
// It returns ok=false if there is no such property.
func (this Node) NextSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeNextSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NodeValue returns the value of property "Node.nodeValue".
//
// It returns ok=false if there is no such property.
func (this Node) NodeValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeNodeValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNodeValue sets the value of property "Node.nodeValue" to val.
//
// It returns false if the property cannot be set.
func (this Node) SetNodeValue(val js.String) bool {
	return js.True == bindings.SetNodeNodeValue(
		this.ref,
		val.Ref(),
	)
}

// TextContent returns the value of property "Node.textContent".
//
// It returns ok=false if there is no such property.
func (this Node) TextContent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeTextContent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextContent sets the value of property "Node.textContent" to val.
//
// It returns false if the property cannot be set.
func (this Node) SetTextContent(val js.String) bool {
	return js.True == bindings.SetNodeTextContent(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetRootNode returns true if the method "Node.getRootNode" exists.
func (this Node) HasFuncGetRootNode() bool {
	return js.True == bindings.HasFuncNodeGetRootNode(
		this.ref,
	)
}

// FuncGetRootNode returns the method "Node.getRootNode".
func (this Node) FuncGetRootNode() (fn js.Func[func(options GetRootNodeOptions) Node]) {
	bindings.FuncNodeGetRootNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRootNode calls the method "Node.getRootNode".
func (this Node) GetRootNode(options GetRootNodeOptions) (ret Node) {
	bindings.CallNodeGetRootNode(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetRootNode calls the method "Node.getRootNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryGetRootNode(options GetRootNodeOptions) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeGetRootNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetRootNode1 returns true if the method "Node.getRootNode" exists.
func (this Node) HasFuncGetRootNode1() bool {
	return js.True == bindings.HasFuncNodeGetRootNode1(
		this.ref,
	)
}

// FuncGetRootNode1 returns the method "Node.getRootNode".
func (this Node) FuncGetRootNode1() (fn js.Func[func() Node]) {
	bindings.FuncNodeGetRootNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRootNode1 calls the method "Node.getRootNode".
func (this Node) GetRootNode1() (ret Node) {
	bindings.CallNodeGetRootNode1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRootNode1 calls the method "Node.getRootNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryGetRootNode1() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeGetRootNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncHasChildNodes returns true if the method "Node.hasChildNodes" exists.
func (this Node) HasFuncHasChildNodes() bool {
	return js.True == bindings.HasFuncNodeHasChildNodes(
		this.ref,
	)
}

// FuncHasChildNodes returns the method "Node.hasChildNodes".
func (this Node) FuncHasChildNodes() (fn js.Func[func() bool]) {
	bindings.FuncNodeHasChildNodes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasChildNodes calls the method "Node.hasChildNodes".
func (this Node) HasChildNodes() (ret bool) {
	bindings.CallNodeHasChildNodes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasChildNodes calls the method "Node.hasChildNodes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryHasChildNodes() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeHasChildNodes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNormalize returns true if the method "Node.normalize" exists.
func (this Node) HasFuncNormalize() bool {
	return js.True == bindings.HasFuncNodeNormalize(
		this.ref,
	)
}

// FuncNormalize returns the method "Node.normalize".
func (this Node) FuncNormalize() (fn js.Func[func()]) {
	bindings.FuncNodeNormalize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Normalize calls the method "Node.normalize".
func (this Node) Normalize() (ret js.Void) {
	bindings.CallNodeNormalize(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryNormalize calls the method "Node.normalize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryNormalize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeNormalize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCloneNode returns true if the method "Node.cloneNode" exists.
func (this Node) HasFuncCloneNode() bool {
	return js.True == bindings.HasFuncNodeCloneNode(
		this.ref,
	)
}

// FuncCloneNode returns the method "Node.cloneNode".
func (this Node) FuncCloneNode() (fn js.Func[func(deep bool) Node]) {
	bindings.FuncNodeCloneNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CloneNode calls the method "Node.cloneNode".
func (this Node) CloneNode(deep bool) (ret Node) {
	bindings.CallNodeCloneNode(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(deep)),
	)

	return
}

// TryCloneNode calls the method "Node.cloneNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCloneNode(deep bool) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCloneNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(deep)),
	)

	return
}

// HasFuncCloneNode1 returns true if the method "Node.cloneNode" exists.
func (this Node) HasFuncCloneNode1() bool {
	return js.True == bindings.HasFuncNodeCloneNode1(
		this.ref,
	)
}

// FuncCloneNode1 returns the method "Node.cloneNode".
func (this Node) FuncCloneNode1() (fn js.Func[func() Node]) {
	bindings.FuncNodeCloneNode1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CloneNode1 calls the method "Node.cloneNode".
func (this Node) CloneNode1() (ret Node) {
	bindings.CallNodeCloneNode1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCloneNode1 calls the method "Node.cloneNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCloneNode1() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCloneNode1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsEqualNode returns true if the method "Node.isEqualNode" exists.
func (this Node) HasFuncIsEqualNode() bool {
	return js.True == bindings.HasFuncNodeIsEqualNode(
		this.ref,
	)
}

// FuncIsEqualNode returns the method "Node.isEqualNode".
func (this Node) FuncIsEqualNode() (fn js.Func[func(otherNode Node) bool]) {
	bindings.FuncNodeIsEqualNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsEqualNode calls the method "Node.isEqualNode".
func (this Node) IsEqualNode(otherNode Node) (ret bool) {
	bindings.CallNodeIsEqualNode(
		this.ref, js.Pointer(&ret),
		otherNode.Ref(),
	)

	return
}

// TryIsEqualNode calls the method "Node.isEqualNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsEqualNode(otherNode Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsEqualNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		otherNode.Ref(),
	)

	return
}

// HasFuncIsSameNode returns true if the method "Node.isSameNode" exists.
func (this Node) HasFuncIsSameNode() bool {
	return js.True == bindings.HasFuncNodeIsSameNode(
		this.ref,
	)
}

// FuncIsSameNode returns the method "Node.isSameNode".
func (this Node) FuncIsSameNode() (fn js.Func[func(otherNode Node) bool]) {
	bindings.FuncNodeIsSameNode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsSameNode calls the method "Node.isSameNode".
func (this Node) IsSameNode(otherNode Node) (ret bool) {
	bindings.CallNodeIsSameNode(
		this.ref, js.Pointer(&ret),
		otherNode.Ref(),
	)

	return
}

// TryIsSameNode calls the method "Node.isSameNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsSameNode(otherNode Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsSameNode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		otherNode.Ref(),
	)

	return
}

// HasFuncCompareDocumentPosition returns true if the method "Node.compareDocumentPosition" exists.
func (this Node) HasFuncCompareDocumentPosition() bool {
	return js.True == bindings.HasFuncNodeCompareDocumentPosition(
		this.ref,
	)
}

// FuncCompareDocumentPosition returns the method "Node.compareDocumentPosition".
func (this Node) FuncCompareDocumentPosition() (fn js.Func[func(other Node) uint16]) {
	bindings.FuncNodeCompareDocumentPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompareDocumentPosition calls the method "Node.compareDocumentPosition".
func (this Node) CompareDocumentPosition(other Node) (ret uint16) {
	bindings.CallNodeCompareDocumentPosition(
		this.ref, js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryCompareDocumentPosition calls the method "Node.compareDocumentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCompareDocumentPosition(other Node) (ret uint16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCompareDocumentPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasFuncContains returns true if the method "Node.contains" exists.
func (this Node) HasFuncContains() bool {
	return js.True == bindings.HasFuncNodeContains(
		this.ref,
	)
}

// FuncContains returns the method "Node.contains".
func (this Node) FuncContains() (fn js.Func[func(other Node) bool]) {
	bindings.FuncNodeContains(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Contains calls the method "Node.contains".
func (this Node) Contains(other Node) (ret bool) {
	bindings.CallNodeContains(
		this.ref, js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryContains calls the method "Node.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryContains(other Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeContains(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasFuncLookupPrefix returns true if the method "Node.lookupPrefix" exists.
func (this Node) HasFuncLookupPrefix() bool {
	return js.True == bindings.HasFuncNodeLookupPrefix(
		this.ref,
	)
}

// FuncLookupPrefix returns the method "Node.lookupPrefix".
func (this Node) FuncLookupPrefix() (fn js.Func[func(namespace js.String) js.String]) {
	bindings.FuncNodeLookupPrefix(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LookupPrefix calls the method "Node.lookupPrefix".
func (this Node) LookupPrefix(namespace js.String) (ret js.String) {
	bindings.CallNodeLookupPrefix(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
	)

	return
}

// TryLookupPrefix calls the method "Node.lookupPrefix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryLookupPrefix(namespace js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeLookupPrefix(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
	)

	return
}

// HasFuncLookupNamespaceURI returns true if the method "Node.lookupNamespaceURI" exists.
func (this Node) HasFuncLookupNamespaceURI() bool {
	return js.True == bindings.HasFuncNodeLookupNamespaceURI(
		this.ref,
	)
}

// FuncLookupNamespaceURI returns the method "Node.lookupNamespaceURI".
func (this Node) FuncLookupNamespaceURI() (fn js.Func[func(prefix js.String) js.String]) {
	bindings.FuncNodeLookupNamespaceURI(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LookupNamespaceURI calls the method "Node.lookupNamespaceURI".
func (this Node) LookupNamespaceURI(prefix js.String) (ret js.String) {
	bindings.CallNodeLookupNamespaceURI(
		this.ref, js.Pointer(&ret),
		prefix.Ref(),
	)

	return
}

// TryLookupNamespaceURI calls the method "Node.lookupNamespaceURI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryLookupNamespaceURI(prefix js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeLookupNamespaceURI(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		prefix.Ref(),
	)

	return
}

// HasFuncIsDefaultNamespace returns true if the method "Node.isDefaultNamespace" exists.
func (this Node) HasFuncIsDefaultNamespace() bool {
	return js.True == bindings.HasFuncNodeIsDefaultNamespace(
		this.ref,
	)
}

// FuncIsDefaultNamespace returns the method "Node.isDefaultNamespace".
func (this Node) FuncIsDefaultNamespace() (fn js.Func[func(namespace js.String) bool]) {
	bindings.FuncNodeIsDefaultNamespace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsDefaultNamespace calls the method "Node.isDefaultNamespace".
func (this Node) IsDefaultNamespace(namespace js.String) (ret bool) {
	bindings.CallNodeIsDefaultNamespace(
		this.ref, js.Pointer(&ret),
		namespace.Ref(),
	)

	return
}

// TryIsDefaultNamespace calls the method "Node.isDefaultNamespace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsDefaultNamespace(namespace js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsDefaultNamespace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
	)

	return
}

// HasFuncInsertBefore returns true if the method "Node.insertBefore" exists.
func (this Node) HasFuncInsertBefore() bool {
	return js.True == bindings.HasFuncNodeInsertBefore(
		this.ref,
	)
}

// FuncInsertBefore returns the method "Node.insertBefore".
func (this Node) FuncInsertBefore() (fn js.Func[func(node Node, child Node) Node]) {
	bindings.FuncNodeInsertBefore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertBefore calls the method "Node.insertBefore".
func (this Node) InsertBefore(node Node, child Node) (ret Node) {
	bindings.CallNodeInsertBefore(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		child.Ref(),
	)

	return
}

// TryInsertBefore calls the method "Node.insertBefore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryInsertBefore(node Node, child Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeInsertBefore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		child.Ref(),
	)

	return
}

// HasFuncAppendChild returns true if the method "Node.appendChild" exists.
func (this Node) HasFuncAppendChild() bool {
	return js.True == bindings.HasFuncNodeAppendChild(
		this.ref,
	)
}

// FuncAppendChild returns the method "Node.appendChild".
func (this Node) FuncAppendChild() (fn js.Func[func(node Node) Node]) {
	bindings.FuncNodeAppendChild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendChild calls the method "Node.appendChild".
func (this Node) AppendChild(node Node) (ret Node) {
	bindings.CallNodeAppendChild(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryAppendChild calls the method "Node.appendChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryAppendChild(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeAppendChild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasFuncReplaceChild returns true if the method "Node.replaceChild" exists.
func (this Node) HasFuncReplaceChild() bool {
	return js.True == bindings.HasFuncNodeReplaceChild(
		this.ref,
	)
}

// FuncReplaceChild returns the method "Node.replaceChild".
func (this Node) FuncReplaceChild() (fn js.Func[func(node Node, child Node) Node]) {
	bindings.FuncNodeReplaceChild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceChild calls the method "Node.replaceChild".
func (this Node) ReplaceChild(node Node, child Node) (ret Node) {
	bindings.CallNodeReplaceChild(
		this.ref, js.Pointer(&ret),
		node.Ref(),
		child.Ref(),
	)

	return
}

// TryReplaceChild calls the method "Node.replaceChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryReplaceChild(node Node, child Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeReplaceChild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		child.Ref(),
	)

	return
}

// HasFuncRemoveChild returns true if the method "Node.removeChild" exists.
func (this Node) HasFuncRemoveChild() bool {
	return js.True == bindings.HasFuncNodeRemoveChild(
		this.ref,
	)
}

// FuncRemoveChild returns the method "Node.removeChild".
func (this Node) FuncRemoveChild() (fn js.Func[func(child Node) Node]) {
	bindings.FuncNodeRemoveChild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveChild calls the method "Node.removeChild".
func (this Node) RemoveChild(child Node) (ret Node) {
	bindings.CallNodeRemoveChild(
		this.ref, js.Pointer(&ret),
		child.Ref(),
	)

	return
}

// TryRemoveChild calls the method "Node.removeChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryRemoveChild(child Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeRemoveChild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		child.Ref(),
	)

	return
}

type AbstractRange struct {
	ref js.Ref
}

func (this AbstractRange) Once() AbstractRange {
	this.ref.Once()
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
	this.ref.Free()
}

// StartContainer returns the value of property "AbstractRange.startContainer".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) StartContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetAbstractRangeStartContainer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StartOffset returns the value of property "AbstractRange.startOffset".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) StartOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAbstractRangeStartOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EndContainer returns the value of property "AbstractRange.endContainer".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) EndContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetAbstractRangeEndContainer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EndOffset returns the value of property "AbstractRange.endOffset".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) EndOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAbstractRangeEndOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Collapsed returns the value of property "AbstractRange.collapsed".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) Collapsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetAbstractRangeCollapsed(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AccelerometerSensorOptions with all fields set.
func (p AccelerometerSensorOptions) FromRef(ref js.Ref) AccelerometerSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AccelerometerSensorOptions in the application heap.
func (p AccelerometerSensorOptions) New() js.Ref {
	return bindings.AccelerometerSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AccelerometerSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.AccelerometerSensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AccelerometerSensorOptions) Update(ref js.Ref) {
	bindings.AccelerometerSensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AccelerometerSensorOptions) FreeMembers(recursive bool) {
}

func NewAccelerometer(options AccelerometerSensorOptions) (ret Accelerometer) {
	ret.ref = bindings.NewAccelerometerByAccelerometer(
		js.Pointer(&options))
	return
}

func NewAccelerometerByAccelerometer1() (ret Accelerometer) {
	ret.ref = bindings.NewAccelerometerByAccelerometer1()
	return
}

type Accelerometer struct {
	Sensor
}

func (this Accelerometer) Once() Accelerometer {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "Accelerometer.x".
//
// It returns ok=false if there is no such property.
func (this Accelerometer) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "Accelerometer.y".
//
// It returns ok=false if there is no such property.
func (this Accelerometer) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "Accelerometer.z".
//
// It returns ok=false if there is no such property.
func (this Accelerometer) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerZ(
		this.ref, js.Pointer(&ret),
	)
	return
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

// New creates a new AccelerometerReadingValues in the application heap.
func (p AccelerometerReadingValues) New() js.Ref {
	return bindings.AccelerometerReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AccelerometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AccelerometerReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AccelerometerReadingValues) Update(ref js.Ref) {
	bindings.AccelerometerReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AccelerometerReadingValues) FreeMembers(recursive bool) {
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

// New creates a new AdRender in the application heap.
func (p AdRender) New() js.Ref {
	return bindings.AdRenderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AdRender) UpdateFrom(ref js.Ref) {
	bindings.AdRenderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AdRender) Update(ref js.Ref) {
	bindings.AdRenderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AdRender) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Width.Ref(),
		p.Height.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Width = p.Width.FromRef(js.Undefined)
	p.Height = p.Height.FromRef(js.Undefined)
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

// New creates a new AddressErrors in the application heap.
func (p AddressErrors) New() js.Ref {
	return bindings.AddressErrorsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressErrors) UpdateFrom(ref js.Ref) {
	bindings.AddressErrorsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressErrors) Update(ref js.Ref) {
	bindings.AddressErrorsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressErrors) FreeMembers(recursive bool) {
	js.Free(
		p.AddressLine.Ref(),
		p.City.Ref(),
		p.Country.Ref(),
		p.DependentLocality.Ref(),
		p.Organization.Ref(),
		p.Phone.Ref(),
		p.PostalCode.Ref(),
		p.Recipient.Ref(),
		p.Region.Ref(),
		p.SortingCode.Ref(),
	)
	p.AddressLine = p.AddressLine.FromRef(js.Undefined)
	p.City = p.City.FromRef(js.Undefined)
	p.Country = p.Country.FromRef(js.Undefined)
	p.DependentLocality = p.DependentLocality.FromRef(js.Undefined)
	p.Organization = p.Organization.FromRef(js.Undefined)
	p.Phone = p.Phone.FromRef(js.Undefined)
	p.PostalCode = p.PostalCode.FromRef(js.Undefined)
	p.Recipient = p.Recipient.FromRef(js.Undefined)
	p.Region = p.Region.FromRef(js.Undefined)
	p.SortingCode = p.SortingCode.FromRef(js.Undefined)
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

// New creates a new AddressInit in the application heap.
func (p AddressInit) New() js.Ref {
	return bindings.AddressInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddressInit) UpdateFrom(ref js.Ref) {
	bindings.AddressInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddressInit) Update(ref js.Ref) {
	bindings.AddressInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddressInit) FreeMembers(recursive bool) {
	js.Free(
		p.Country.Ref(),
		p.AddressLine.Ref(),
		p.Region.Ref(),
		p.City.Ref(),
		p.DependentLocality.Ref(),
		p.PostalCode.Ref(),
		p.SortingCode.Ref(),
		p.Organization.Ref(),
		p.Recipient.Ref(),
		p.Phone.Ref(),
	)
	p.Country = p.Country.FromRef(js.Undefined)
	p.AddressLine = p.AddressLine.FromRef(js.Undefined)
	p.Region = p.Region.FromRef(js.Undefined)
	p.City = p.City.FromRef(js.Undefined)
	p.DependentLocality = p.DependentLocality.FromRef(js.Undefined)
	p.PostalCode = p.PostalCode.FromRef(js.Undefined)
	p.SortingCode = p.SortingCode.FromRef(js.Undefined)
	p.Organization = p.Organization.FromRef(js.Undefined)
	p.Recipient = p.Recipient.FromRef(js.Undefined)
	p.Phone = p.Phone.FromRef(js.Undefined)
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

// New creates a new AesCbcParams in the application heap.
func (p AesCbcParams) New() js.Ref {
	return bindings.AesCbcParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesCbcParams) UpdateFrom(ref js.Ref) {
	bindings.AesCbcParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesCbcParams) Update(ref js.Ref) {
	bindings.AesCbcParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesCbcParams) FreeMembers(recursive bool) {
	js.Free(
		p.Iv.Ref(),
		p.Name.Ref(),
	)
	p.Iv = p.Iv.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new AesCtrParams in the application heap.
func (p AesCtrParams) New() js.Ref {
	return bindings.AesCtrParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesCtrParams) UpdateFrom(ref js.Ref) {
	bindings.AesCtrParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesCtrParams) Update(ref js.Ref) {
	bindings.AesCtrParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesCtrParams) FreeMembers(recursive bool) {
	js.Free(
		p.Counter.Ref(),
		p.Name.Ref(),
	)
	p.Counter = p.Counter.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new AesDerivedKeyParams in the application heap.
func (p AesDerivedKeyParams) New() js.Ref {
	return bindings.AesDerivedKeyParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesDerivedKeyParams) UpdateFrom(ref js.Ref) {
	bindings.AesDerivedKeyParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesDerivedKeyParams) Update(ref js.Ref) {
	bindings.AesDerivedKeyParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesDerivedKeyParams) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
	//
	// NOTE: FFI_USE_TagLength MUST be set to true to make this field effective.
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

// New creates a new AesGcmParams in the application heap.
func (p AesGcmParams) New() js.Ref {
	return bindings.AesGcmParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesGcmParams) UpdateFrom(ref js.Ref) {
	bindings.AesGcmParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesGcmParams) Update(ref js.Ref) {
	bindings.AesGcmParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesGcmParams) FreeMembers(recursive bool) {
	js.Free(
		p.Iv.Ref(),
		p.AdditionalData.Ref(),
		p.Name.Ref(),
	)
	p.Iv = p.Iv.FromRef(js.Undefined)
	p.AdditionalData = p.AdditionalData.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new AesKeyAlgorithm in the application heap.
func (p AesKeyAlgorithm) New() js.Ref {
	return bindings.AesKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.AesKeyAlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesKeyAlgorithm) Update(ref js.Ref) {
	bindings.AesKeyAlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesKeyAlgorithm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new AesKeyGenParams in the application heap.
func (p AesKeyGenParams) New() js.Ref {
	return bindings.AesKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AesKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.AesKeyGenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AesKeyGenParams) Update(ref js.Ref) {
	bindings.AesKeyGenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AesKeyGenParams) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new Algorithm in the application heap.
func (p Algorithm) New() js.Ref {
	return bindings.AlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Algorithm) UpdateFrom(ref js.Ref) {
	bindings.AlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Algorithm) Update(ref js.Ref) {
	bindings.AlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Algorithm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new AllowedBluetoothDevice in the application heap.
func (p AllowedBluetoothDevice) New() js.Ref {
	return bindings.AllowedBluetoothDeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AllowedBluetoothDevice) UpdateFrom(ref js.Ref) {
	bindings.AllowedBluetoothDeviceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AllowedBluetoothDevice) Update(ref js.Ref) {
	bindings.AllowedBluetoothDeviceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AllowedBluetoothDevice) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceId.Ref(),
		p.AllowedServices.Ref(),
		p.AllowedManufacturerData.Ref(),
	)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.AllowedServices = p.AllowedServices.FromRef(js.Undefined)
	p.AllowedManufacturerData = p.AllowedManufacturerData.FromRef(js.Undefined)
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

// New creates a new AllowedUSBDevice in the application heap.
func (p AllowedUSBDevice) New() js.Ref {
	return bindings.AllowedUSBDeviceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AllowedUSBDevice) UpdateFrom(ref js.Ref) {
	bindings.AllowedUSBDeviceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AllowedUSBDevice) Update(ref js.Ref) {
	bindings.AllowedUSBDeviceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AllowedUSBDevice) FreeMembers(recursive bool) {
	js.Free(
		p.SerialNumber.Ref(),
	)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
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

// New creates a new AmbientLightReadingValues in the application heap.
func (p AmbientLightReadingValues) New() js.Ref {
	return bindings.AmbientLightReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AmbientLightReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AmbientLightReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AmbientLightReadingValues) Update(ref js.Ref) {
	bindings.AmbientLightReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AmbientLightReadingValues) FreeMembers(recursive bool) {
}

type SensorOptions struct {
	// Frequency is "SensorOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SensorOptions with all fields set.
func (p SensorOptions) FromRef(ref js.Ref) SensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SensorOptions in the application heap.
func (p SensorOptions) New() js.Ref {
	return bindings.SensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SensorOptions) UpdateFrom(ref js.Ref) {
	bindings.SensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SensorOptions) Update(ref js.Ref) {
	bindings.SensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SensorOptions) FreeMembers(recursive bool) {
}

func NewAmbientLightSensor(sensorOptions SensorOptions) (ret AmbientLightSensor) {
	ret.ref = bindings.NewAmbientLightSensorByAmbientLightSensor(
		js.Pointer(&sensorOptions))
	return
}

func NewAmbientLightSensorByAmbientLightSensor1() (ret AmbientLightSensor) {
	ret.ref = bindings.NewAmbientLightSensorByAmbientLightSensor1()
	return
}

type AmbientLightSensor struct {
	Sensor
}

func (this AmbientLightSensor) Once() AmbientLightSensor {
	this.ref.Once()
	return this
}

func (this AmbientLightSensor) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this AmbientLightSensor) FromRef(ref js.Ref) AmbientLightSensor {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this AmbientLightSensor) Free() {
	this.ref.Free()
}

// Illuminance returns the value of property "AmbientLightSensor.illuminance".
//
// It returns ok=false if there is no such property.
func (this AmbientLightSensor) Illuminance() (ret float64, ok bool) {
	ok = js.True == bindings.GetAmbientLightSensorIlluminance(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BiquadFilterType uint32
