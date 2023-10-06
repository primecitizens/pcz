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
// It returns ok=false if there is no such property.
func (this FontFacePalette) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UsableWithLightBackground returns the value of property "FontFacePalette.usableWithLightBackground".
//
// It returns ok=false if there is no such property.
func (this FontFacePalette) UsableWithLightBackground() (ret bool, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteUsableWithLightBackground(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UsableWithDarkBackground returns the value of property "FontFacePalette.usableWithDarkBackground".
//
// It returns ok=false if there is no such property.
func (this FontFacePalette) UsableWithDarkBackground() (ret bool, ok bool) {
	ok = js.True == bindings.GetFontFacePaletteUsableWithDarkBackground(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "FontFacePalette." exists.
func (this FontFacePalette) HasGet() bool {
	return js.True == bindings.HasFontFacePaletteGet(
		this.Ref(),
	)
}

// GetFunc returns the method "FontFacePalette.".
func (this FontFacePalette) GetFunc() (fn js.Func[func(index uint32) js.String]) {
	return fn.FromRef(
		bindings.FontFacePaletteGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "FontFacePalette.".
func (this FontFacePalette) Get(index uint32) (ret js.String) {
	bindings.CallFontFacePaletteGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "FontFacePalette."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFacePalette) TryGet(index uint32) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFacePaletteGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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
// It returns ok=false if there is no such property.
func (this FontFacePalettes) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetFontFacePalettesLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "FontFacePalettes." exists.
func (this FontFacePalettes) HasGet() bool {
	return js.True == bindings.HasFontFacePalettesGet(
		this.Ref(),
	)
}

// GetFunc returns the method "FontFacePalettes.".
func (this FontFacePalettes) GetFunc() (fn js.Func[func(index uint32) FontFacePalette]) {
	return fn.FromRef(
		bindings.FontFacePalettesGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "FontFacePalettes.".
func (this FontFacePalettes) Get(index uint32) (ret FontFacePalette) {
	bindings.CallFontFacePalettesGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "FontFacePalettes."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFacePalettes) TryGet(index uint32) (ret FontFacePalette, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFacePalettesGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
// It returns ok=false if there is no such property.
func (this FontFace) Family() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceFamily(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFamily sets the value of property "FontFace.family" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Style() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStyle sets the value of property "FontFace.style" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Weight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceWeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWeight sets the value of property "FontFace.weight" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Stretch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceStretch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStretch sets the value of property "FontFace.stretch" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) UnicodeRange() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceUnicodeRange(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUnicodeRange sets the value of property "FontFace.unicodeRange" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Variant() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariant(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVariant sets the value of property "FontFace.variant" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) FeatureSettings() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceFeatureSettings(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFeatureSettings sets the value of property "FontFace.featureSettings" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) VariationSettings() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationSettings(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVariationSettings sets the value of property "FontFace.variationSettings" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Display() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceDisplay(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisplay sets the value of property "FontFace.display" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) AscentOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceAscentOverride(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAscentOverride sets the value of property "FontFace.ascentOverride" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) DescentOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceDescentOverride(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDescentOverride sets the value of property "FontFace.descentOverride" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) LineGapOverride() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceLineGapOverride(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineGapOverride sets the value of property "FontFace.lineGapOverride" to val.
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
// It returns ok=false if there is no such property.
func (this FontFace) Status() (ret FontFaceLoadStatus, ok bool) {
	ok = js.True == bindings.GetFontFaceStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Loaded returns the value of property "FontFace.loaded".
//
// It returns ok=false if there is no such property.
func (this FontFace) Loaded() (ret js.Promise[FontFace], ok bool) {
	ok = js.True == bindings.GetFontFaceLoaded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Features returns the value of property "FontFace.features".
//
// It returns ok=false if there is no such property.
func (this FontFace) Features() (ret FontFaceFeatures, ok bool) {
	ok = js.True == bindings.GetFontFaceFeatures(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Variations returns the value of property "FontFace.variations".
//
// It returns ok=false if there is no such property.
func (this FontFace) Variations() (ret FontFaceVariations, ok bool) {
	ok = js.True == bindings.GetFontFaceVariations(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Palettes returns the value of property "FontFace.palettes".
//
// It returns ok=false if there is no such property.
func (this FontFace) Palettes() (ret FontFacePalettes, ok bool) {
	ok = js.True == bindings.GetFontFacePalettes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasLoad returns true if the method "FontFace.load" exists.
func (this FontFace) HasLoad() bool {
	return js.True == bindings.HasFontFaceLoad(
		this.Ref(),
	)
}

// LoadFunc returns the method "FontFace.load".
func (this FontFace) LoadFunc() (fn js.Func[func() js.Promise[FontFace]]) {
	return fn.FromRef(
		bindings.FontFaceLoadFunc(
			this.Ref(),
		),
	)
}

// Load calls the method "FontFace.load".
func (this FontFace) Load() (ret js.Promise[FontFace]) {
	bindings.CallFontFaceLoad(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLoad calls the method "FontFace.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFace) TryLoad() (ret js.Promise[FontFace], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceLoad(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
// It returns ok=false if there is no such property.
func (this FontFaceSet) Ready() (ret js.Promise[FontFaceSet], ok bool) {
	ok = js.True == bindings.GetFontFaceSetReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Status returns the value of property "FontFaceSet.status".
//
// It returns ok=false if there is no such property.
func (this FontFaceSet) Status() (ret FontFaceSetLoadStatus, ok bool) {
	ok = js.True == bindings.GetFontFaceSetStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAdd returns true if the method "FontFaceSet.add" exists.
func (this FontFaceSet) HasAdd() bool {
	return js.True == bindings.HasFontFaceSetAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "FontFaceSet.add".
func (this FontFaceSet) AddFunc() (fn js.Func[func(font FontFace) FontFaceSet]) {
	return fn.FromRef(
		bindings.FontFaceSetAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "FontFaceSet.add".
func (this FontFaceSet) Add(font FontFace) (ret FontFaceSet) {
	bindings.CallFontFaceSetAdd(
		this.Ref(), js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryAdd calls the method "FontFaceSet.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryAdd(font FontFace) (ret FontFaceSet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasDelete returns true if the method "FontFaceSet.delete" exists.
func (this FontFaceSet) HasDelete() bool {
	return js.True == bindings.HasFontFaceSetDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "FontFaceSet.delete".
func (this FontFaceSet) DeleteFunc() (fn js.Func[func(font FontFace) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "FontFaceSet.delete".
func (this FontFaceSet) Delete(font FontFace) (ret bool) {
	bindings.CallFontFaceSetDelete(
		this.Ref(), js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryDelete calls the method "FontFaceSet.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryDelete(font FontFace) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasClear returns true if the method "FontFaceSet.clear" exists.
func (this FontFaceSet) HasClear() bool {
	return js.True == bindings.HasFontFaceSetClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "FontFaceSet.clear".
func (this FontFaceSet) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FontFaceSetClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "FontFaceSet.clear".
func (this FontFaceSet) Clear() (ret js.Void) {
	bindings.CallFontFaceSetClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "FontFaceSet.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasLoad returns true if the method "FontFaceSet.load" exists.
func (this FontFaceSet) HasLoad() bool {
	return js.True == bindings.HasFontFaceSetLoad(
		this.Ref(),
	)
}

// LoadFunc returns the method "FontFaceSet.load".
func (this FontFaceSet) LoadFunc() (fn js.Func[func(font js.String, text js.String) js.Promise[js.Array[FontFace]]]) {
	return fn.FromRef(
		bindings.FontFaceSetLoadFunc(
			this.Ref(),
		),
	)
}

// Load calls the method "FontFaceSet.load".
func (this FontFaceSet) Load(font js.String, text js.String) (ret js.Promise[js.Array[FontFace]]) {
	bindings.CallFontFaceSetLoad(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
		text.Ref(),
	)

	return
}

// HasLoad1 returns true if the method "FontFaceSet.load" exists.
func (this FontFaceSet) HasLoad1() bool {
	return js.True == bindings.HasFontFaceSetLoad1(
		this.Ref(),
	)
}

// Load1Func returns the method "FontFaceSet.load".
func (this FontFaceSet) Load1Func() (fn js.Func[func(font js.String) js.Promise[js.Array[FontFace]]]) {
	return fn.FromRef(
		bindings.FontFaceSetLoad1Func(
			this.Ref(),
		),
	)
}

// Load1 calls the method "FontFaceSet.load".
func (this FontFaceSet) Load1(font js.String) (ret js.Promise[js.Array[FontFace]]) {
	bindings.CallFontFaceSetLoad1(
		this.Ref(), js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryLoad1 calls the method "FontFaceSet.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryLoad1(font js.String) (ret js.Promise[js.Array[FontFace]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetLoad1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
}

// HasCheck returns true if the method "FontFaceSet.check" exists.
func (this FontFaceSet) HasCheck() bool {
	return js.True == bindings.HasFontFaceSetCheck(
		this.Ref(),
	)
}

// CheckFunc returns the method "FontFaceSet.check".
func (this FontFaceSet) CheckFunc() (fn js.Func[func(font js.String, text js.String) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetCheckFunc(
			this.Ref(),
		),
	)
}

// Check calls the method "FontFaceSet.check".
func (this FontFaceSet) Check(font js.String, text js.String) (ret bool) {
	bindings.CallFontFaceSetCheck(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
		text.Ref(),
	)

	return
}

// HasCheck1 returns true if the method "FontFaceSet.check" exists.
func (this FontFaceSet) HasCheck1() bool {
	return js.True == bindings.HasFontFaceSetCheck1(
		this.Ref(),
	)
}

// Check1Func returns the method "FontFaceSet.check".
func (this FontFaceSet) Check1Func() (fn js.Func[func(font js.String) bool]) {
	return fn.FromRef(
		bindings.FontFaceSetCheck1Func(
			this.Ref(),
		),
	)
}

// Check1 calls the method "FontFaceSet.check".
func (this FontFaceSet) Check1(font js.String) (ret bool) {
	bindings.CallFontFaceSetCheck1(
		this.Ref(), js.Pointer(&ret),
		font.Ref(),
	)

	return
}

// TryCheck1 calls the method "FontFaceSet.check"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontFaceSet) TryCheck1(font js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontFaceSetCheck1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		font.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Document) Implementation() (ret DOMImplementation, ok bool) {
	ok = js.True == bindings.GetDocumentImplementation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// URL returns the value of property "Document.URL".
//
// It returns ok=false if there is no such property.
func (this Document) URL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DocumentURI returns the value of property "Document.documentURI".
//
// It returns ok=false if there is no such property.
func (this Document) DocumentURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDocumentURI(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CompatMode returns the value of property "Document.compatMode".
//
// It returns ok=false if there is no such property.
func (this Document) CompatMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCompatMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CharacterSet returns the value of property "Document.characterSet".
//
// It returns ok=false if there is no such property.
func (this Document) CharacterSet() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCharacterSet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Charset returns the value of property "Document.charset".
//
// It returns ok=false if there is no such property.
func (this Document) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCharset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InputEncoding returns the value of property "Document.inputEncoding".
//
// It returns ok=false if there is no such property.
func (this Document) InputEncoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentInputEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentType returns the value of property "Document.contentType".
//
// It returns ok=false if there is no such property.
func (this Document) ContentType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentContentType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Doctype returns the value of property "Document.doctype".
//
// It returns ok=false if there is no such property.
func (this Document) Doctype() (ret DocumentType, ok bool) {
	ok = js.True == bindings.GetDocumentDoctype(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DocumentElement returns the value of property "Document.documentElement".
//
// It returns ok=false if there is no such property.
func (this Document) DocumentElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentDocumentElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FragmentDirective returns the value of property "Document.fragmentDirective".
//
// It returns ok=false if there is no such property.
func (this Document) FragmentDirective() (ret FragmentDirective, ok bool) {
	ok = js.True == bindings.GetDocumentFragmentDirective(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PermissionsPolicy returns the value of property "Document.permissionsPolicy".
//
// It returns ok=false if there is no such property.
func (this Document) PermissionsPolicy() (ret PermissionsPolicy, ok bool) {
	ok = js.True == bindings.GetDocumentPermissionsPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WasDiscarded returns the value of property "Document.wasDiscarded".
//
// It returns ok=false if there is no such property.
func (this Document) WasDiscarded() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentWasDiscarded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FullscreenEnabled returns the value of property "Document.fullscreenEnabled".
//
// It returns ok=false if there is no such property.
func (this Document) FullscreenEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreenEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fullscreen returns the value of property "Document.fullscreen".
//
// It returns ok=false if there is no such property.
func (this Document) Fullscreen() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timeline returns the value of property "Document.timeline".
//
// It returns ok=false if there is no such property.
func (this Document) Timeline() (ret DocumentTimeline, ok bool) {
	ok = js.True == bindings.GetDocumentTimeline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PictureInPictureEnabled returns the value of property "Document.pictureInPictureEnabled".
//
// It returns ok=false if there is no such property.
func (this Document) PictureInPictureEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NamedFlows returns the value of property "Document.namedFlows".
//
// It returns ok=false if there is no such property.
func (this Document) NamedFlows() (ret NamedFlowMap, ok bool) {
	ok = js.True == bindings.GetDocumentNamedFlows(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScrollingElement returns the value of property "Document.scrollingElement".
//
// It returns ok=false if there is no such property.
func (this Document) ScrollingElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentScrollingElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RootElement returns the value of property "Document.rootElement".
//
// It returns ok=false if there is no such property.
func (this Document) RootElement() (ret SVGSVGElement, ok bool) {
	ok = js.True == bindings.GetDocumentRootElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Prerendering returns the value of property "Document.prerendering".
//
// It returns ok=false if there is no such property.
func (this Document) Prerendering() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentPrerendering(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FgColor returns the value of property "Document.fgColor".
//
// It returns ok=false if there is no such property.
func (this Document) FgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentFgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFgColor sets the value of property "Document.fgColor" to val.
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
// It returns ok=false if there is no such property.
func (this Document) LinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentLinkColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLinkColor sets the value of property "Document.linkColor" to val.
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
// It returns ok=false if there is no such property.
func (this Document) VlinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentVlinkColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVlinkColor sets the value of property "Document.vlinkColor" to val.
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
// It returns ok=false if there is no such property.
func (this Document) AlinkColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentAlinkColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlinkColor sets the value of property "Document.alinkColor" to val.
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
// It returns ok=false if there is no such property.
func (this Document) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "Document.bgColor" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Anchors() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentAnchors(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Applets returns the value of property "Document.applets".
//
// It returns ok=false if there is no such property.
func (this Document) Applets() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentApplets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// All returns the value of property "Document.all".
//
// It returns ok=false if there is no such property.
func (this Document) All() (ret HTMLAllCollection, ok bool) {
	ok = js.True == bindings.GetDocumentAll(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Location returns the value of property "Document.location".
//
// It returns ok=false if there is no such property.
func (this Document) Location() (ret Location, ok bool) {
	ok = js.True == bindings.GetDocumentLocation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Domain returns the value of property "Document.domain".
//
// It returns ok=false if there is no such property.
func (this Document) Domain() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDomain(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDomain sets the value of property "Document.domain" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentReferrer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cookie returns the value of property "Document.cookie".
//
// It returns ok=false if there is no such property.
func (this Document) Cookie() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentCookie(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCookie sets the value of property "Document.cookie" to val.
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
// It returns ok=false if there is no such property.
func (this Document) LastModified() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentLastModified(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "Document.readyState".
//
// It returns ok=false if there is no such property.
func (this Document) ReadyState() (ret DocumentReadyState, ok bool) {
	ok = js.True == bindings.GetDocumentReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Title returns the value of property "Document.title".
//
// It returns ok=false if there is no such property.
func (this Document) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentTitle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTitle sets the value of property "Document.title" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Dir() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDir(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDir sets the value of property "Document.dir" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Body() (ret HTMLElement, ok bool) {
	ok = js.True == bindings.GetDocumentBody(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBody sets the value of property "Document.body" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Head() (ret HTMLHeadElement, ok bool) {
	ok = js.True == bindings.GetDocumentHead(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Images returns the value of property "Document.images".
//
// It returns ok=false if there is no such property.
func (this Document) Images() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentImages(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Embeds returns the value of property "Document.embeds".
//
// It returns ok=false if there is no such property.
func (this Document) Embeds() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentEmbeds(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Plugins returns the value of property "Document.plugins".
//
// It returns ok=false if there is no such property.
func (this Document) Plugins() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentPlugins(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Links returns the value of property "Document.links".
//
// It returns ok=false if there is no such property.
func (this Document) Links() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentLinks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Forms returns the value of property "Document.forms".
//
// It returns ok=false if there is no such property.
func (this Document) Forms() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentForms(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scripts returns the value of property "Document.scripts".
//
// It returns ok=false if there is no such property.
func (this Document) Scripts() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentScripts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentScript returns the value of property "Document.currentScript".
//
// It returns ok=false if there is no such property.
func (this Document) CurrentScript() (ret HTMLOrSVGScriptElement, ok bool) {
	ok = js.True == bindings.GetDocumentCurrentScript(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultView returns the value of property "Document.defaultView".
//
// It returns ok=false if there is no such property.
func (this Document) DefaultView() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetDocumentDefaultView(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DesignMode returns the value of property "Document.designMode".
//
// It returns ok=false if there is no such property.
func (this Document) DesignMode() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDocumentDesignMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDesignMode sets the value of property "Document.designMode" to val.
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
// It returns ok=false if there is no such property.
func (this Document) Hidden() (ret bool, ok bool) {
	ok = js.True == bindings.GetDocumentHidden(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VisibilityState returns the value of property "Document.visibilityState".
//
// It returns ok=false if there is no such property.
func (this Document) VisibilityState() (ret DocumentVisibilityState, ok bool) {
	ok = js.True == bindings.GetDocumentVisibilityState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FullscreenElement returns the value of property "Document.fullscreenElement".
//
// It returns ok=false if there is no such property.
func (this Document) FullscreenElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFullscreenElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StyleSheets returns the value of property "Document.styleSheets".
//
// It returns ok=false if there is no such property.
func (this Document) StyleSheets() (ret StyleSheetList, ok bool) {
	ok = js.True == bindings.GetDocumentStyleSheets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AdoptedStyleSheets returns the value of property "Document.adoptedStyleSheets".
//
// It returns ok=false if there is no such property.
func (this Document) AdoptedStyleSheets() (ret js.ObservableArray[CSSStyleSheet], ok bool) {
	ok = js.True == bindings.GetDocumentAdoptedStyleSheets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAdoptedStyleSheets sets the value of property "Document.adoptedStyleSheets" to val.
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
// It returns ok=false if there is no such property.
func (this Document) PictureInPictureElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActiveElement returns the value of property "Document.activeElement".
//
// It returns ok=false if there is no such property.
func (this Document) ActiveElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentActiveElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PointerLockElement returns the value of property "Document.pointerLockElement".
//
// It returns ok=false if there is no such property.
func (this Document) PointerLockElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentPointerLockElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Children returns the value of property "Document.children".
//
// It returns ok=false if there is no such property.
func (this Document) Children() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetDocumentChildren(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstElementChild returns the value of property "Document.firstElementChild".
//
// It returns ok=false if there is no such property.
func (this Document) FirstElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentFirstElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastElementChild returns the value of property "Document.lastElementChild".
//
// It returns ok=false if there is no such property.
func (this Document) LastElementChild() (ret Element, ok bool) {
	ok = js.True == bindings.GetDocumentLastElementChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChildElementCount returns the value of property "Document.childElementCount".
//
// It returns ok=false if there is no such property.
func (this Document) ChildElementCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetDocumentChildElementCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Fonts returns the value of property "Document.fonts".
//
// It returns ok=false if there is no such property.
func (this Document) Fonts() (ret FontFaceSet, ok bool) {
	ok = js.True == bindings.GetDocumentFonts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetElementsByTagName returns true if the method "Document.getElementsByTagName" exists.
func (this Document) HasGetElementsByTagName() bool {
	return js.True == bindings.HasDocumentGetElementsByTagName(
		this.Ref(),
	)
}

// GetElementsByTagNameFunc returns the method "Document.getElementsByTagName".
func (this Document) GetElementsByTagNameFunc() (fn js.Func[func(qualifiedName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByTagNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagName calls the method "Document.getElementsByTagName".
func (this Document) GetElementsByTagName(qualifiedName js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByTagName(
		this.Ref(), js.Pointer(&ret),
		qualifiedName.Ref(),
	)

	return
}

// TryGetElementsByTagName calls the method "Document.getElementsByTagName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByTagName(qualifiedName js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByTagName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		qualifiedName.Ref(),
	)

	return
}

// HasGetElementsByTagNameNS returns true if the method "Document.getElementsByTagNameNS" exists.
func (this Document) HasGetElementsByTagNameNS() bool {
	return js.True == bindings.HasDocumentGetElementsByTagNameNS(
		this.Ref(),
	)
}

// GetElementsByTagNameNSFunc returns the method "Document.getElementsByTagNameNS".
func (this Document) GetElementsByTagNameNSFunc() (fn js.Func[func(namespace js.String, localName js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByTagNameNSFunc(
			this.Ref(),
		),
	)
}

// GetElementsByTagNameNS calls the method "Document.getElementsByTagNameNS".
func (this Document) GetElementsByTagNameNS(namespace js.String, localName js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByTagNameNS(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		localName.Ref(),
	)

	return
}

// HasGetElementsByClassName returns true if the method "Document.getElementsByClassName" exists.
func (this Document) HasGetElementsByClassName() bool {
	return js.True == bindings.HasDocumentGetElementsByClassName(
		this.Ref(),
	)
}

// GetElementsByClassNameFunc returns the method "Document.getElementsByClassName".
func (this Document) GetElementsByClassNameFunc() (fn js.Func[func(classNames js.String) HTMLCollection]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByClassNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByClassName calls the method "Document.getElementsByClassName".
func (this Document) GetElementsByClassName(classNames js.String) (ret HTMLCollection) {
	bindings.CallDocumentGetElementsByClassName(
		this.Ref(), js.Pointer(&ret),
		classNames.Ref(),
	)

	return
}

// TryGetElementsByClassName calls the method "Document.getElementsByClassName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByClassName(classNames js.String) (ret HTMLCollection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByClassName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		classNames.Ref(),
	)

	return
}

// HasCreateElement returns true if the method "Document.createElement" exists.
func (this Document) HasCreateElement() bool {
	return js.True == bindings.HasDocumentCreateElement(
		this.Ref(),
	)
}

// CreateElementFunc returns the method "Document.createElement".
func (this Document) CreateElementFunc() (fn js.Func[func(localName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementFunc(
			this.Ref(),
		),
	)
}

// CreateElement calls the method "Document.createElement".
func (this Document) CreateElement(localName js.String, options OneOf_String_ElementCreationOptions) (ret Element) {
	bindings.CallDocumentCreateElement(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
		options.Ref(),
	)

	return
}

// HasCreateElement1 returns true if the method "Document.createElement" exists.
func (this Document) HasCreateElement1() bool {
	return js.True == bindings.HasDocumentCreateElement1(
		this.Ref(),
	)
}

// CreateElement1Func returns the method "Document.createElement".
func (this Document) CreateElement1Func() (fn js.Func[func(localName js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElement1Func(
			this.Ref(),
		),
	)
}

// CreateElement1 calls the method "Document.createElement".
func (this Document) CreateElement1(localName js.String) (ret Element) {
	bindings.CallDocumentCreateElement1(
		this.Ref(), js.Pointer(&ret),
		localName.Ref(),
	)

	return
}

// TryCreateElement1 calls the method "Document.createElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateElement1(localName js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateElement1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
	)

	return
}

// HasCreateElementNS returns true if the method "Document.createElementNS" exists.
func (this Document) HasCreateElementNS() bool {
	return js.True == bindings.HasDocumentCreateElementNS(
		this.Ref(),
	)
}

// CreateElementNSFunc returns the method "Document.createElementNS".
func (this Document) CreateElementNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementNSFunc(
			this.Ref(),
		),
	)
}

// CreateElementNS calls the method "Document.createElementNS".
func (this Document) CreateElementNS(namespace js.String, qualifiedName js.String, options OneOf_String_ElementCreationOptions) (ret Element) {
	bindings.CallDocumentCreateElementNS(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
		options.Ref(),
	)

	return
}

// HasCreateElementNS1 returns true if the method "Document.createElementNS" exists.
func (this Document) HasCreateElementNS1() bool {
	return js.True == bindings.HasDocumentCreateElementNS1(
		this.Ref(),
	)
}

// CreateElementNS1Func returns the method "Document.createElementNS".
func (this Document) CreateElementNS1Func() (fn js.Func[func(namespace js.String, qualifiedName js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentCreateElementNS1Func(
			this.Ref(),
		),
	)
}

// CreateElementNS1 calls the method "Document.createElementNS".
func (this Document) CreateElementNS1(namespace js.String, qualifiedName js.String) (ret Element) {
	bindings.CallDocumentCreateElementNS1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasCreateDocumentFragment returns true if the method "Document.createDocumentFragment" exists.
func (this Document) HasCreateDocumentFragment() bool {
	return js.True == bindings.HasDocumentCreateDocumentFragment(
		this.Ref(),
	)
}

// CreateDocumentFragmentFunc returns the method "Document.createDocumentFragment".
func (this Document) CreateDocumentFragmentFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.DocumentCreateDocumentFragmentFunc(
			this.Ref(),
		),
	)
}

// CreateDocumentFragment calls the method "Document.createDocumentFragment".
func (this Document) CreateDocumentFragment() (ret DocumentFragment) {
	bindings.CallDocumentCreateDocumentFragment(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateDocumentFragment calls the method "Document.createDocumentFragment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateDocumentFragment() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateDocumentFragment(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateTextNode returns true if the method "Document.createTextNode" exists.
func (this Document) HasCreateTextNode() bool {
	return js.True == bindings.HasDocumentCreateTextNode(
		this.Ref(),
	)
}

// CreateTextNodeFunc returns the method "Document.createTextNode".
func (this Document) CreateTextNodeFunc() (fn js.Func[func(data js.String) Text]) {
	return fn.FromRef(
		bindings.DocumentCreateTextNodeFunc(
			this.Ref(),
		),
	)
}

// CreateTextNode calls the method "Document.createTextNode".
func (this Document) CreateTextNode(data js.String) (ret Text) {
	bindings.CallDocumentCreateTextNode(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateTextNode calls the method "Document.createTextNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTextNode(data js.String) (ret Text, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTextNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasCreateCDATASection returns true if the method "Document.createCDATASection" exists.
func (this Document) HasCreateCDATASection() bool {
	return js.True == bindings.HasDocumentCreateCDATASection(
		this.Ref(),
	)
}

// CreateCDATASectionFunc returns the method "Document.createCDATASection".
func (this Document) CreateCDATASectionFunc() (fn js.Func[func(data js.String) CDATASection]) {
	return fn.FromRef(
		bindings.DocumentCreateCDATASectionFunc(
			this.Ref(),
		),
	)
}

// CreateCDATASection calls the method "Document.createCDATASection".
func (this Document) CreateCDATASection(data js.String) (ret CDATASection) {
	bindings.CallDocumentCreateCDATASection(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateCDATASection calls the method "Document.createCDATASection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateCDATASection(data js.String) (ret CDATASection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateCDATASection(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasCreateComment returns true if the method "Document.createComment" exists.
func (this Document) HasCreateComment() bool {
	return js.True == bindings.HasDocumentCreateComment(
		this.Ref(),
	)
}

// CreateCommentFunc returns the method "Document.createComment".
func (this Document) CreateCommentFunc() (fn js.Func[func(data js.String) Comment]) {
	return fn.FromRef(
		bindings.DocumentCreateCommentFunc(
			this.Ref(),
		),
	)
}

// CreateComment calls the method "Document.createComment".
func (this Document) CreateComment(data js.String) (ret Comment) {
	bindings.CallDocumentCreateComment(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryCreateComment calls the method "Document.createComment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateComment(data js.String) (ret Comment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateComment(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasCreateProcessingInstruction returns true if the method "Document.createProcessingInstruction" exists.
func (this Document) HasCreateProcessingInstruction() bool {
	return js.True == bindings.HasDocumentCreateProcessingInstruction(
		this.Ref(),
	)
}

// CreateProcessingInstructionFunc returns the method "Document.createProcessingInstruction".
func (this Document) CreateProcessingInstructionFunc() (fn js.Func[func(target js.String, data js.String) ProcessingInstruction]) {
	return fn.FromRef(
		bindings.DocumentCreateProcessingInstructionFunc(
			this.Ref(),
		),
	)
}

// CreateProcessingInstruction calls the method "Document.createProcessingInstruction".
func (this Document) CreateProcessingInstruction(target js.String, data js.String) (ret ProcessingInstruction) {
	bindings.CallDocumentCreateProcessingInstruction(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
		data.Ref(),
	)

	return
}

// HasImportNode returns true if the method "Document.importNode" exists.
func (this Document) HasImportNode() bool {
	return js.True == bindings.HasDocumentImportNode(
		this.Ref(),
	)
}

// ImportNodeFunc returns the method "Document.importNode".
func (this Document) ImportNodeFunc() (fn js.Func[func(node Node, deep bool) Node]) {
	return fn.FromRef(
		bindings.DocumentImportNodeFunc(
			this.Ref(),
		),
	)
}

// ImportNode calls the method "Document.importNode".
func (this Document) ImportNode(node Node, deep bool) (ret Node) {
	bindings.CallDocumentImportNode(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		js.Bool(bool(deep)),
	)

	return
}

// HasImportNode1 returns true if the method "Document.importNode" exists.
func (this Document) HasImportNode1() bool {
	return js.True == bindings.HasDocumentImportNode1(
		this.Ref(),
	)
}

// ImportNode1Func returns the method "Document.importNode".
func (this Document) ImportNode1Func() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.DocumentImportNode1Func(
			this.Ref(),
		),
	)
}

// ImportNode1 calls the method "Document.importNode".
func (this Document) ImportNode1(node Node) (ret Node) {
	bindings.CallDocumentImportNode1(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryImportNode1 calls the method "Document.importNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryImportNode1(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentImportNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasAdoptNode returns true if the method "Document.adoptNode" exists.
func (this Document) HasAdoptNode() bool {
	return js.True == bindings.HasDocumentAdoptNode(
		this.Ref(),
	)
}

// AdoptNodeFunc returns the method "Document.adoptNode".
func (this Document) AdoptNodeFunc() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.DocumentAdoptNodeFunc(
			this.Ref(),
		),
	)
}

// AdoptNode calls the method "Document.adoptNode".
func (this Document) AdoptNode(node Node) (ret Node) {
	bindings.CallDocumentAdoptNode(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryAdoptNode calls the method "Document.adoptNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryAdoptNode(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentAdoptNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasCreateAttribute returns true if the method "Document.createAttribute" exists.
func (this Document) HasCreateAttribute() bool {
	return js.True == bindings.HasDocumentCreateAttribute(
		this.Ref(),
	)
}

// CreateAttributeFunc returns the method "Document.createAttribute".
func (this Document) CreateAttributeFunc() (fn js.Func[func(localName js.String) Attr]) {
	return fn.FromRef(
		bindings.DocumentCreateAttributeFunc(
			this.Ref(),
		),
	)
}

// CreateAttribute calls the method "Document.createAttribute".
func (this Document) CreateAttribute(localName js.String) (ret Attr) {
	bindings.CallDocumentCreateAttribute(
		this.Ref(), js.Pointer(&ret),
		localName.Ref(),
	)

	return
}

// TryCreateAttribute calls the method "Document.createAttribute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateAttribute(localName js.String) (ret Attr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateAttribute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		localName.Ref(),
	)

	return
}

// HasCreateAttributeNS returns true if the method "Document.createAttributeNS" exists.
func (this Document) HasCreateAttributeNS() bool {
	return js.True == bindings.HasDocumentCreateAttributeNS(
		this.Ref(),
	)
}

// CreateAttributeNSFunc returns the method "Document.createAttributeNS".
func (this Document) CreateAttributeNSFunc() (fn js.Func[func(namespace js.String, qualifiedName js.String) Attr]) {
	return fn.FromRef(
		bindings.DocumentCreateAttributeNSFunc(
			this.Ref(),
		),
	)
}

// CreateAttributeNS calls the method "Document.createAttributeNS".
func (this Document) CreateAttributeNS(namespace js.String, qualifiedName js.String) (ret Attr) {
	bindings.CallDocumentCreateAttributeNS(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
		qualifiedName.Ref(),
	)

	return
}

// HasCreateEvent returns true if the method "Document.createEvent" exists.
func (this Document) HasCreateEvent() bool {
	return js.True == bindings.HasDocumentCreateEvent(
		this.Ref(),
	)
}

// CreateEventFunc returns the method "Document.createEvent".
func (this Document) CreateEventFunc() (fn js.Func[func(iface js.String) Event]) {
	return fn.FromRef(
		bindings.DocumentCreateEventFunc(
			this.Ref(),
		),
	)
}

// CreateEvent calls the method "Document.createEvent".
func (this Document) CreateEvent(iface js.String) (ret Event) {
	bindings.CallDocumentCreateEvent(
		this.Ref(), js.Pointer(&ret),
		iface.Ref(),
	)

	return
}

// TryCreateEvent calls the method "Document.createEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateEvent(iface js.String) (ret Event, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		iface.Ref(),
	)

	return
}

// HasCreateRange returns true if the method "Document.createRange" exists.
func (this Document) HasCreateRange() bool {
	return js.True == bindings.HasDocumentCreateRange(
		this.Ref(),
	)
}

// CreateRangeFunc returns the method "Document.createRange".
func (this Document) CreateRangeFunc() (fn js.Func[func() Range]) {
	return fn.FromRef(
		bindings.DocumentCreateRangeFunc(
			this.Ref(),
		),
	)
}

// CreateRange calls the method "Document.createRange".
func (this Document) CreateRange() (ret Range) {
	bindings.CallDocumentCreateRange(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateRange calls the method "Document.createRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateRange() (ret Range, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateNodeIterator returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasCreateNodeIterator() bool {
	return js.True == bindings.HasDocumentCreateNodeIterator(
		this.Ref(),
	)
}

// CreateNodeIteratorFunc returns the method "Document.createNodeIterator".
func (this Document) CreateNodeIteratorFunc() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIteratorFunc(
			this.Ref(),
		),
	)
}

// CreateNodeIterator calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// HasCreateNodeIterator1 returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasCreateNodeIterator1() bool {
	return js.True == bindings.HasDocumentCreateNodeIterator1(
		this.Ref(),
	)
}

// CreateNodeIterator1Func returns the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator1Func() (fn js.Func[func(root Node, whatToShow uint32) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIterator1Func(
			this.Ref(),
		),
	)
}

// CreateNodeIterator1 calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator1(root Node, whatToShow uint32) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// HasCreateNodeIterator2 returns true if the method "Document.createNodeIterator" exists.
func (this Document) HasCreateNodeIterator2() bool {
	return js.True == bindings.HasDocumentCreateNodeIterator2(
		this.Ref(),
	)
}

// CreateNodeIterator2Func returns the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator2Func() (fn js.Func[func(root Node) NodeIterator]) {
	return fn.FromRef(
		bindings.DocumentCreateNodeIterator2Func(
			this.Ref(),
		),
	)
}

// CreateNodeIterator2 calls the method "Document.createNodeIterator".
func (this Document) CreateNodeIterator2(root Node) (ret NodeIterator) {
	bindings.CallDocumentCreateNodeIterator2(
		this.Ref(), js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryCreateNodeIterator2 calls the method "Document.createNodeIterator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNodeIterator2(root Node) (ret NodeIterator, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNodeIterator2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

// HasCreateTreeWalker returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasCreateTreeWalker() bool {
	return js.True == bindings.HasDocumentCreateTreeWalker(
		this.Ref(),
	)
}

// CreateTreeWalkerFunc returns the method "Document.createTreeWalker".
func (this Document) CreateTreeWalkerFunc() (fn js.Func[func(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalkerFunc(
			this.Ref(),
		),
	)
}

// CreateTreeWalker calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker(root Node, whatToShow uint32, filter js.Func[func(node Node) uint16]) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
		filter.Ref(),
	)

	return
}

// HasCreateTreeWalker1 returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasCreateTreeWalker1() bool {
	return js.True == bindings.HasDocumentCreateTreeWalker1(
		this.Ref(),
	)
}

// CreateTreeWalker1Func returns the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker1Func() (fn js.Func[func(root Node, whatToShow uint32) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalker1Func(
			this.Ref(),
		),
	)
}

// CreateTreeWalker1 calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker1(root Node, whatToShow uint32) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
		uint32(whatToShow),
	)

	return
}

// HasCreateTreeWalker2 returns true if the method "Document.createTreeWalker" exists.
func (this Document) HasCreateTreeWalker2() bool {
	return js.True == bindings.HasDocumentCreateTreeWalker2(
		this.Ref(),
	)
}

// CreateTreeWalker2Func returns the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker2Func() (fn js.Func[func(root Node) TreeWalker]) {
	return fn.FromRef(
		bindings.DocumentCreateTreeWalker2Func(
			this.Ref(),
		),
	)
}

// CreateTreeWalker2 calls the method "Document.createTreeWalker".
func (this Document) CreateTreeWalker2(root Node) (ret TreeWalker) {
	bindings.CallDocumentCreateTreeWalker2(
		this.Ref(), js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryCreateTreeWalker2 calls the method "Document.createTreeWalker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateTreeWalker2(root Node) (ret TreeWalker, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateTreeWalker2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

// HasStartViewTransition returns true if the method "Document.startViewTransition" exists.
func (this Document) HasStartViewTransition() bool {
	return js.True == bindings.HasDocumentStartViewTransition(
		this.Ref(),
	)
}

// StartViewTransitionFunc returns the method "Document.startViewTransition".
func (this Document) StartViewTransitionFunc() (fn js.Func[func(updateCallback js.Func[func() js.Promise[js.Any]]) ViewTransition]) {
	return fn.FromRef(
		bindings.DocumentStartViewTransitionFunc(
			this.Ref(),
		),
	)
}

// StartViewTransition calls the method "Document.startViewTransition".
func (this Document) StartViewTransition(updateCallback js.Func[func() js.Promise[js.Any]]) (ret ViewTransition) {
	bindings.CallDocumentStartViewTransition(
		this.Ref(), js.Pointer(&ret),
		updateCallback.Ref(),
	)

	return
}

// TryStartViewTransition calls the method "Document.startViewTransition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryStartViewTransition(updateCallback js.Func[func() js.Promise[js.Any]]) (ret ViewTransition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentStartViewTransition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		updateCallback.Ref(),
	)

	return
}

// HasStartViewTransition1 returns true if the method "Document.startViewTransition" exists.
func (this Document) HasStartViewTransition1() bool {
	return js.True == bindings.HasDocumentStartViewTransition1(
		this.Ref(),
	)
}

// StartViewTransition1Func returns the method "Document.startViewTransition".
func (this Document) StartViewTransition1Func() (fn js.Func[func() ViewTransition]) {
	return fn.FromRef(
		bindings.DocumentStartViewTransition1Func(
			this.Ref(),
		),
	)
}

// StartViewTransition1 calls the method "Document.startViewTransition".
func (this Document) StartViewTransition1() (ret ViewTransition) {
	bindings.CallDocumentStartViewTransition1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStartViewTransition1 calls the method "Document.startViewTransition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryStartViewTransition1() (ret ViewTransition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentStartViewTransition1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMeasureElement returns true if the method "Document.measureElement" exists.
func (this Document) HasMeasureElement() bool {
	return js.True == bindings.HasDocumentMeasureElement(
		this.Ref(),
	)
}

// MeasureElementFunc returns the method "Document.measureElement".
func (this Document) MeasureElementFunc() (fn js.Func[func(element Element) FontMetrics]) {
	return fn.FromRef(
		bindings.DocumentMeasureElementFunc(
			this.Ref(),
		),
	)
}

// MeasureElement calls the method "Document.measureElement".
func (this Document) MeasureElement(element Element) (ret FontMetrics) {
	bindings.CallDocumentMeasureElement(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryMeasureElement calls the method "Document.measureElement"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryMeasureElement(element Element) (ret FontMetrics, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentMeasureElement(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasMeasureText returns true if the method "Document.measureText" exists.
func (this Document) HasMeasureText() bool {
	return js.True == bindings.HasDocumentMeasureText(
		this.Ref(),
	)
}

// MeasureTextFunc returns the method "Document.measureText".
func (this Document) MeasureTextFunc() (fn js.Func[func(text js.String, styleMap StylePropertyMapReadOnly) FontMetrics]) {
	return fn.FromRef(
		bindings.DocumentMeasureTextFunc(
			this.Ref(),
		),
	)
}

// MeasureText calls the method "Document.measureText".
func (this Document) MeasureText(text js.String, styleMap StylePropertyMapReadOnly) (ret FontMetrics) {
	bindings.CallDocumentMeasureText(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		styleMap.Ref(),
	)

	return
}

// HasGetSelection returns true if the method "Document.getSelection" exists.
func (this Document) HasGetSelection() bool {
	return js.True == bindings.HasDocumentGetSelection(
		this.Ref(),
	)
}

// GetSelectionFunc returns the method "Document.getSelection".
func (this Document) GetSelectionFunc() (fn js.Func[func() Selection]) {
	return fn.FromRef(
		bindings.DocumentGetSelectionFunc(
			this.Ref(),
		),
	)
}

// GetSelection calls the method "Document.getSelection".
func (this Document) GetSelection() (ret Selection) {
	bindings.CallDocumentGetSelection(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSelection calls the method "Document.getSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetSelection() (ret Selection, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetSelection(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExitPointerLock returns true if the method "Document.exitPointerLock" exists.
func (this Document) HasExitPointerLock() bool {
	return js.True == bindings.HasDocumentExitPointerLock(
		this.Ref(),
	)
}

// ExitPointerLockFunc returns the method "Document.exitPointerLock".
func (this Document) ExitPointerLockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentExitPointerLockFunc(
			this.Ref(),
		),
	)
}

// ExitPointerLock calls the method "Document.exitPointerLock".
func (this Document) ExitPointerLock() (ret js.Void) {
	bindings.CallDocumentExitPointerLock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryExitPointerLock calls the method "Document.exitPointerLock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitPointerLock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitPointerLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExitFullscreen returns true if the method "Document.exitFullscreen" exists.
func (this Document) HasExitFullscreen() bool {
	return js.True == bindings.HasDocumentExitFullscreen(
		this.Ref(),
	)
}

// ExitFullscreenFunc returns the method "Document.exitFullscreen".
func (this Document) ExitFullscreenFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentExitFullscreenFunc(
			this.Ref(),
		),
	)
}

// ExitFullscreen calls the method "Document.exitFullscreen".
func (this Document) ExitFullscreen() (ret js.Promise[js.Void]) {
	bindings.CallDocumentExitFullscreen(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryExitFullscreen calls the method "Document.exitFullscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitFullscreen() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitFullscreen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHasStorageAccess returns true if the method "Document.hasStorageAccess" exists.
func (this Document) HasHasStorageAccess() bool {
	return js.True == bindings.HasDocumentHasStorageAccess(
		this.Ref(),
	)
}

// HasStorageAccessFunc returns the method "Document.hasStorageAccess".
func (this Document) HasStorageAccessFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasStorageAccessFunc(
			this.Ref(),
		),
	)
}

// HasStorageAccess calls the method "Document.hasStorageAccess".
func (this Document) HasStorageAccess() (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasStorageAccess(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHasStorageAccess calls the method "Document.hasStorageAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasStorageAccess() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasStorageAccess(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestStorageAccess returns true if the method "Document.requestStorageAccess" exists.
func (this Document) HasRequestStorageAccess() bool {
	return js.True == bindings.HasDocumentRequestStorageAccess(
		this.Ref(),
	)
}

// RequestStorageAccessFunc returns the method "Document.requestStorageAccess".
func (this Document) RequestStorageAccessFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentRequestStorageAccessFunc(
			this.Ref(),
		),
	)
}

// RequestStorageAccess calls the method "Document.requestStorageAccess".
func (this Document) RequestStorageAccess() (ret js.Promise[js.Void]) {
	bindings.CallDocumentRequestStorageAccess(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestStorageAccess calls the method "Document.requestStorageAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryRequestStorageAccess() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentRequestStorageAccess(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExitPictureInPicture returns true if the method "Document.exitPictureInPicture" exists.
func (this Document) HasExitPictureInPicture() bool {
	return js.True == bindings.HasDocumentExitPictureInPicture(
		this.Ref(),
	)
}

// ExitPictureInPictureFunc returns the method "Document.exitPictureInPicture".
func (this Document) ExitPictureInPictureFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentExitPictureInPictureFunc(
			this.Ref(),
		),
	)
}

// ExitPictureInPicture calls the method "Document.exitPictureInPicture".
func (this Document) ExitPictureInPicture() (ret js.Promise[js.Void]) {
	bindings.CallDocumentExitPictureInPicture(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryExitPictureInPicture calls the method "Document.exitPictureInPicture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExitPictureInPicture() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExitPictureInPicture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHasPrivateTokens returns true if the method "Document.hasPrivateTokens" exists.
func (this Document) HasHasPrivateTokens() bool {
	return js.True == bindings.HasDocumentHasPrivateTokens(
		this.Ref(),
	)
}

// HasPrivateTokensFunc returns the method "Document.hasPrivateTokens".
func (this Document) HasPrivateTokensFunc() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasPrivateTokensFunc(
			this.Ref(),
		),
	)
}

// HasPrivateTokens calls the method "Document.hasPrivateTokens".
func (this Document) HasPrivateTokens(issuer js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasPrivateTokens(
		this.Ref(), js.Pointer(&ret),
		issuer.Ref(),
	)

	return
}

// TryHasPrivateTokens calls the method "Document.hasPrivateTokens"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasPrivateTokens(issuer js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasPrivateTokens(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		issuer.Ref(),
	)

	return
}

// HasHasRedemptionRecord returns true if the method "Document.hasRedemptionRecord" exists.
func (this Document) HasHasRedemptionRecord() bool {
	return js.True == bindings.HasDocumentHasRedemptionRecord(
		this.Ref(),
	)
}

// HasRedemptionRecordFunc returns the method "Document.hasRedemptionRecord".
func (this Document) HasRedemptionRecordFunc() (fn js.Func[func(issuer js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.DocumentHasRedemptionRecordFunc(
			this.Ref(),
		),
	)
}

// HasRedemptionRecord calls the method "Document.hasRedemptionRecord".
func (this Document) HasRedemptionRecord(issuer js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallDocumentHasRedemptionRecord(
		this.Ref(), js.Pointer(&ret),
		issuer.Ref(),
	)

	return
}

// TryHasRedemptionRecord calls the method "Document.hasRedemptionRecord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasRedemptionRecord(issuer js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasRedemptionRecord(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		issuer.Ref(),
	)

	return
}

// HasElementFromPoint returns true if the method "Document.elementFromPoint" exists.
func (this Document) HasElementFromPoint() bool {
	return js.True == bindings.HasDocumentElementFromPoint(
		this.Ref(),
	)
}

// ElementFromPointFunc returns the method "Document.elementFromPoint".
func (this Document) ElementFromPointFunc() (fn js.Func[func(x float64, y float64) Element]) {
	return fn.FromRef(
		bindings.DocumentElementFromPointFunc(
			this.Ref(),
		),
	)
}

// ElementFromPoint calls the method "Document.elementFromPoint".
func (this Document) ElementFromPoint(x float64, y float64) (ret Element) {
	bindings.CallDocumentElementFromPoint(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasElementsFromPoint returns true if the method "Document.elementsFromPoint" exists.
func (this Document) HasElementsFromPoint() bool {
	return js.True == bindings.HasDocumentElementsFromPoint(
		this.Ref(),
	)
}

// ElementsFromPointFunc returns the method "Document.elementsFromPoint".
func (this Document) ElementsFromPointFunc() (fn js.Func[func(x float64, y float64) js.Array[Element]]) {
	return fn.FromRef(
		bindings.DocumentElementsFromPointFunc(
			this.Ref(),
		),
	)
}

// ElementsFromPoint calls the method "Document.elementsFromPoint".
func (this Document) ElementsFromPoint(x float64, y float64) (ret js.Array[Element]) {
	bindings.CallDocumentElementsFromPoint(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasCaretPositionFromPoint returns true if the method "Document.caretPositionFromPoint" exists.
func (this Document) HasCaretPositionFromPoint() bool {
	return js.True == bindings.HasDocumentCaretPositionFromPoint(
		this.Ref(),
	)
}

// CaretPositionFromPointFunc returns the method "Document.caretPositionFromPoint".
func (this Document) CaretPositionFromPointFunc() (fn js.Func[func(x float64, y float64) CaretPosition]) {
	return fn.FromRef(
		bindings.DocumentCaretPositionFromPointFunc(
			this.Ref(),
		),
	)
}

// CaretPositionFromPoint calls the method "Document.caretPositionFromPoint".
func (this Document) CaretPositionFromPoint(x float64, y float64) (ret CaretPosition) {
	bindings.CallDocumentCaretPositionFromPoint(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasClear returns true if the method "Document.clear" exists.
func (this Document) HasClear() bool {
	return js.True == bindings.HasDocumentClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "Document.clear".
func (this Document) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "Document.clear".
func (this Document) Clear() (ret js.Void) {
	bindings.CallDocumentClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "Document.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCaptureEvents returns true if the method "Document.captureEvents" exists.
func (this Document) HasCaptureEvents() bool {
	return js.True == bindings.HasDocumentCaptureEvents(
		this.Ref(),
	)
}

// CaptureEventsFunc returns the method "Document.captureEvents".
func (this Document) CaptureEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentCaptureEventsFunc(
			this.Ref(),
		),
	)
}

// CaptureEvents calls the method "Document.captureEvents".
func (this Document) CaptureEvents() (ret js.Void) {
	bindings.CallDocumentCaptureEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCaptureEvents calls the method "Document.captureEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCaptureEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCaptureEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReleaseEvents returns true if the method "Document.releaseEvents" exists.
func (this Document) HasReleaseEvents() bool {
	return js.True == bindings.HasDocumentReleaseEvents(
		this.Ref(),
	)
}

// ReleaseEventsFunc returns the method "Document.releaseEvents".
func (this Document) ReleaseEventsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentReleaseEventsFunc(
			this.Ref(),
		),
	)
}

// ReleaseEvents calls the method "Document.releaseEvents".
func (this Document) ReleaseEvents() (ret js.Void) {
	bindings.CallDocumentReleaseEvents(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReleaseEvents calls the method "Document.releaseEvents"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryReleaseEvents() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentReleaseEvents(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestStorageAccessFor returns true if the method "Document.requestStorageAccessFor" exists.
func (this Document) HasRequestStorageAccessFor() bool {
	return js.True == bindings.HasDocumentRequestStorageAccessFor(
		this.Ref(),
	)
}

// RequestStorageAccessForFunc returns the method "Document.requestStorageAccessFor".
func (this Document) RequestStorageAccessForFunc() (fn js.Func[func(requestedOrigin js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DocumentRequestStorageAccessForFunc(
			this.Ref(),
		),
	)
}

// RequestStorageAccessFor calls the method "Document.requestStorageAccessFor".
func (this Document) RequestStorageAccessFor(requestedOrigin js.String) (ret js.Promise[js.Void]) {
	bindings.CallDocumentRequestStorageAccessFor(
		this.Ref(), js.Pointer(&ret),
		requestedOrigin.Ref(),
	)

	return
}

// TryRequestStorageAccessFor calls the method "Document.requestStorageAccessFor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryRequestStorageAccessFor(requestedOrigin js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentRequestStorageAccessFor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		requestedOrigin.Ref(),
	)

	return
}

// HasGet returns true if the method "Document." exists.
func (this Document) HasGet() bool {
	return js.True == bindings.HasDocumentGet(
		this.Ref(),
	)
}

// GetFunc returns the method "Document.".
func (this Document) GetFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.DocumentGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "Document.".
func (this Document) Get(name js.String) (ret js.Object) {
	bindings.CallDocumentGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "Document."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGet(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetElementsByName returns true if the method "Document.getElementsByName" exists.
func (this Document) HasGetElementsByName() bool {
	return js.True == bindings.HasDocumentGetElementsByName(
		this.Ref(),
	)
}

// GetElementsByNameFunc returns the method "Document.getElementsByName".
func (this Document) GetElementsByNameFunc() (fn js.Func[func(elementName js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentGetElementsByNameFunc(
			this.Ref(),
		),
	)
}

// GetElementsByName calls the method "Document.getElementsByName".
func (this Document) GetElementsByName(elementName js.String) (ret NodeList) {
	bindings.CallDocumentGetElementsByName(
		this.Ref(), js.Pointer(&ret),
		elementName.Ref(),
	)

	return
}

// TryGetElementsByName calls the method "Document.getElementsByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementsByName(elementName js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementsByName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elementName.Ref(),
	)

	return
}

// HasOpen returns true if the method "Document.open" exists.
func (this Document) HasOpen() bool {
	return js.True == bindings.HasDocumentOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "Document.open".
func (this Document) OpenFunc() (fn js.Func[func(unused1 js.String, unused2 js.String) Document]) {
	return fn.FromRef(
		bindings.DocumentOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "Document.open".
func (this Document) Open(unused1 js.String, unused2 js.String) (ret Document) {
	bindings.CallDocumentOpen(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		unused1.Ref(),
		unused2.Ref(),
	)

	return
}

// HasOpen1 returns true if the method "Document.open" exists.
func (this Document) HasOpen1() bool {
	return js.True == bindings.HasDocumentOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "Document.open".
func (this Document) Open1Func() (fn js.Func[func(unused1 js.String) Document]) {
	return fn.FromRef(
		bindings.DocumentOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "Document.open".
func (this Document) Open1(unused1 js.String) (ret Document) {
	bindings.CallDocumentOpen1(
		this.Ref(), js.Pointer(&ret),
		unused1.Ref(),
	)

	return
}

// TryOpen1 calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen1(unused1 js.String) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		unused1.Ref(),
	)

	return
}

// HasOpen2 returns true if the method "Document.open" exists.
func (this Document) HasOpen2() bool {
	return js.True == bindings.HasDocumentOpen2(
		this.Ref(),
	)
}

// Open2Func returns the method "Document.open".
func (this Document) Open2Func() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.DocumentOpen2Func(
			this.Ref(),
		),
	)
}

// Open2 calls the method "Document.open".
func (this Document) Open2() (ret Document) {
	bindings.CallDocumentOpen2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpen2 calls the method "Document.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryOpen2() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentOpen2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpen3 returns true if the method "Document.open" exists.
func (this Document) HasOpen3() bool {
	return js.True == bindings.HasDocumentOpen3(
		this.Ref(),
	)
}

// Open3Func returns the method "Document.open".
func (this Document) Open3Func() (fn js.Func[func(url js.String, name js.String, features js.String) js.Object]) {
	return fn.FromRef(
		bindings.DocumentOpen3Func(
			this.Ref(),
		),
	)
}

// Open3 calls the method "Document.open".
func (this Document) Open3(url js.String, name js.String, features js.String) (ret js.Object) {
	bindings.CallDocumentOpen3(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		name.Ref(),
		features.Ref(),
	)

	return
}

// HasClose returns true if the method "Document.close" exists.
func (this Document) HasClose() bool {
	return js.True == bindings.HasDocumentClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "Document.close".
func (this Document) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.DocumentCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "Document.close".
func (this Document) Close() (ret js.Void) {
	bindings.CallDocumentClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "Document.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWrite returns true if the method "Document.write" exists.
func (this Document) HasWrite() bool {
	return js.True == bindings.HasDocumentWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "Document.write".
func (this Document) WriteFunc() (fn js.Func[func(text ...js.String)]) {
	return fn.FromRef(
		bindings.DocumentWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "Document.write".
func (this Document) Write(text ...js.String) (ret js.Void) {
	bindings.CallDocumentWrite(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// HasWriteln returns true if the method "Document.writeln" exists.
func (this Document) HasWriteln() bool {
	return js.True == bindings.HasDocumentWriteln(
		this.Ref(),
	)
}

// WritelnFunc returns the method "Document.writeln".
func (this Document) WritelnFunc() (fn js.Func[func(text ...js.String)]) {
	return fn.FromRef(
		bindings.DocumentWritelnFunc(
			this.Ref(),
		),
	)
}

// Writeln calls the method "Document.writeln".
func (this Document) Writeln(text ...js.String) (ret js.Void) {
	bindings.CallDocumentWriteln(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(text),
		js.SizeU(len(text)),
	)

	return
}

// HasHasFocus returns true if the method "Document.hasFocus" exists.
func (this Document) HasHasFocus() bool {
	return js.True == bindings.HasDocumentHasFocus(
		this.Ref(),
	)
}

// HasFocusFunc returns the method "Document.hasFocus".
func (this Document) HasFocusFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.DocumentHasFocusFunc(
			this.Ref(),
		),
	)
}

// HasFocus calls the method "Document.hasFocus".
func (this Document) HasFocus() (ret bool) {
	bindings.CallDocumentHasFocus(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHasFocus calls the method "Document.hasFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryHasFocus() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentHasFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExecCommand returns true if the method "Document.execCommand" exists.
func (this Document) HasExecCommand() bool {
	return js.True == bindings.HasDocumentExecCommand(
		this.Ref(),
	)
}

// ExecCommandFunc returns the method "Document.execCommand".
func (this Document) ExecCommandFunc() (fn js.Func[func(commandId js.String, showUI bool, value js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommandFunc(
			this.Ref(),
		),
	)
}

// ExecCommand calls the method "Document.execCommand".
func (this Document) ExecCommand(commandId js.String, showUI bool, value js.String) (ret bool) {
	bindings.CallDocumentExecCommand(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
		js.Bool(bool(showUI)),
		value.Ref(),
	)

	return
}

// HasExecCommand1 returns true if the method "Document.execCommand" exists.
func (this Document) HasExecCommand1() bool {
	return js.True == bindings.HasDocumentExecCommand1(
		this.Ref(),
	)
}

// ExecCommand1Func returns the method "Document.execCommand".
func (this Document) ExecCommand1Func() (fn js.Func[func(commandId js.String, showUI bool) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommand1Func(
			this.Ref(),
		),
	)
}

// ExecCommand1 calls the method "Document.execCommand".
func (this Document) ExecCommand1(commandId js.String, showUI bool) (ret bool) {
	bindings.CallDocumentExecCommand1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
		js.Bool(bool(showUI)),
	)

	return
}

// HasExecCommand2 returns true if the method "Document.execCommand" exists.
func (this Document) HasExecCommand2() bool {
	return js.True == bindings.HasDocumentExecCommand2(
		this.Ref(),
	)
}

// ExecCommand2Func returns the method "Document.execCommand".
func (this Document) ExecCommand2Func() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentExecCommand2Func(
			this.Ref(),
		),
	)
}

// ExecCommand2 calls the method "Document.execCommand".
func (this Document) ExecCommand2(commandId js.String) (ret bool) {
	bindings.CallDocumentExecCommand2(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryExecCommand2 calls the method "Document.execCommand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryExecCommand2(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentExecCommand2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasQueryCommandEnabled returns true if the method "Document.queryCommandEnabled" exists.
func (this Document) HasQueryCommandEnabled() bool {
	return js.True == bindings.HasDocumentQueryCommandEnabled(
		this.Ref(),
	)
}

// QueryCommandEnabledFunc returns the method "Document.queryCommandEnabled".
func (this Document) QueryCommandEnabledFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandEnabledFunc(
			this.Ref(),
		),
	)
}

// QueryCommandEnabled calls the method "Document.queryCommandEnabled".
func (this Document) QueryCommandEnabled(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandEnabled(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandEnabled calls the method "Document.queryCommandEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandEnabled(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasQueryCommandIndeterm returns true if the method "Document.queryCommandIndeterm" exists.
func (this Document) HasQueryCommandIndeterm() bool {
	return js.True == bindings.HasDocumentQueryCommandIndeterm(
		this.Ref(),
	)
}

// QueryCommandIndetermFunc returns the method "Document.queryCommandIndeterm".
func (this Document) QueryCommandIndetermFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandIndetermFunc(
			this.Ref(),
		),
	)
}

// QueryCommandIndeterm calls the method "Document.queryCommandIndeterm".
func (this Document) QueryCommandIndeterm(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandIndeterm(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandIndeterm calls the method "Document.queryCommandIndeterm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandIndeterm(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandIndeterm(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasQueryCommandState returns true if the method "Document.queryCommandState" exists.
func (this Document) HasQueryCommandState() bool {
	return js.True == bindings.HasDocumentQueryCommandState(
		this.Ref(),
	)
}

// QueryCommandStateFunc returns the method "Document.queryCommandState".
func (this Document) QueryCommandStateFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandStateFunc(
			this.Ref(),
		),
	)
}

// QueryCommandState calls the method "Document.queryCommandState".
func (this Document) QueryCommandState(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandState(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandState calls the method "Document.queryCommandState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandState(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasQueryCommandSupported returns true if the method "Document.queryCommandSupported" exists.
func (this Document) HasQueryCommandSupported() bool {
	return js.True == bindings.HasDocumentQueryCommandSupported(
		this.Ref(),
	)
}

// QueryCommandSupportedFunc returns the method "Document.queryCommandSupported".
func (this Document) QueryCommandSupportedFunc() (fn js.Func[func(commandId js.String) bool]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandSupportedFunc(
			this.Ref(),
		),
	)
}

// QueryCommandSupported calls the method "Document.queryCommandSupported".
func (this Document) QueryCommandSupported(commandId js.String) (ret bool) {
	bindings.CallDocumentQueryCommandSupported(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandSupported calls the method "Document.queryCommandSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandSupported(commandId js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasQueryCommandValue returns true if the method "Document.queryCommandValue" exists.
func (this Document) HasQueryCommandValue() bool {
	return js.True == bindings.HasDocumentQueryCommandValue(
		this.Ref(),
	)
}

// QueryCommandValueFunc returns the method "Document.queryCommandValue".
func (this Document) QueryCommandValueFunc() (fn js.Func[func(commandId js.String) js.String]) {
	return fn.FromRef(
		bindings.DocumentQueryCommandValueFunc(
			this.Ref(),
		),
	)
}

// QueryCommandValue calls the method "Document.queryCommandValue".
func (this Document) QueryCommandValue(commandId js.String) (ret js.String) {
	bindings.CallDocumentQueryCommandValue(
		this.Ref(), js.Pointer(&ret),
		commandId.Ref(),
	)

	return
}

// TryQueryCommandValue calls the method "Document.queryCommandValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQueryCommandValue(commandId js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQueryCommandValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandId.Ref(),
	)

	return
}

// HasGetBoxQuads returns true if the method "Document.getBoxQuads" exists.
func (this Document) HasGetBoxQuads() bool {
	return js.True == bindings.HasDocumentGetBoxQuads(
		this.Ref(),
	)
}

// GetBoxQuadsFunc returns the method "Document.getBoxQuads".
func (this Document) GetBoxQuadsFunc() (fn js.Func[func(options BoxQuadOptions) js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.DocumentGetBoxQuadsFunc(
			this.Ref(),
		),
	)
}

// GetBoxQuads calls the method "Document.getBoxQuads".
func (this Document) GetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad]) {
	bindings.CallDocumentGetBoxQuads(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetBoxQuads calls the method "Document.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetBoxQuads(options BoxQuadOptions) (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetBoxQuads(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetBoxQuads1 returns true if the method "Document.getBoxQuads" exists.
func (this Document) HasGetBoxQuads1() bool {
	return js.True == bindings.HasDocumentGetBoxQuads1(
		this.Ref(),
	)
}

// GetBoxQuads1Func returns the method "Document.getBoxQuads".
func (this Document) GetBoxQuads1Func() (fn js.Func[func() js.Array[DOMQuad]]) {
	return fn.FromRef(
		bindings.DocumentGetBoxQuads1Func(
			this.Ref(),
		),
	)
}

// GetBoxQuads1 calls the method "Document.getBoxQuads".
func (this Document) GetBoxQuads1() (ret js.Array[DOMQuad]) {
	bindings.CallDocumentGetBoxQuads1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBoxQuads1 calls the method "Document.getBoxQuads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetBoxQuads1() (ret js.Array[DOMQuad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetBoxQuads1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvertQuadFromNode returns true if the method "Document.convertQuadFromNode" exists.
func (this Document) HasConvertQuadFromNode() bool {
	return js.True == bindings.HasDocumentConvertQuadFromNode(
		this.Ref(),
	)
}

// ConvertQuadFromNodeFunc returns the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNodeFunc() (fn js.Func[func(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertQuadFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode calls the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNode(quad DOMQuadInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallDocumentConvertQuadFromNode(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertQuadFromNode1 returns true if the method "Document.convertQuadFromNode" exists.
func (this Document) HasConvertQuadFromNode1() bool {
	return js.True == bindings.HasDocumentConvertQuadFromNode1(
		this.Ref(),
	)
}

// ConvertQuadFromNode1Func returns the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNode1Func() (fn js.Func[func(quad DOMQuadInit, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertQuadFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertQuadFromNode1 calls the method "Document.convertQuadFromNode".
func (this Document) ConvertQuadFromNode1(quad DOMQuadInit, from GeometryNode) (ret DOMQuad) {
	bindings.CallDocumentConvertQuadFromNode1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&quad),
		from.Ref(),
	)

	return
}

// HasConvertRectFromNode returns true if the method "Document.convertRectFromNode" exists.
func (this Document) HasConvertRectFromNode() bool {
	return js.True == bindings.HasDocumentConvertRectFromNode(
		this.Ref(),
	)
}

// ConvertRectFromNodeFunc returns the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNodeFunc() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertRectFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode calls the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNode(rect DOMRectReadOnly, from GeometryNode, options ConvertCoordinateOptions) (ret DOMQuad) {
	bindings.CallDocumentConvertRectFromNode(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertRectFromNode1 returns true if the method "Document.convertRectFromNode" exists.
func (this Document) HasConvertRectFromNode1() bool {
	return js.True == bindings.HasDocumentConvertRectFromNode1(
		this.Ref(),
	)
}

// ConvertRectFromNode1Func returns the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNode1Func() (fn js.Func[func(rect DOMRectReadOnly, from GeometryNode) DOMQuad]) {
	return fn.FromRef(
		bindings.DocumentConvertRectFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertRectFromNode1 calls the method "Document.convertRectFromNode".
func (this Document) ConvertRectFromNode1(rect DOMRectReadOnly, from GeometryNode) (ret DOMQuad) {
	bindings.CallDocumentConvertRectFromNode1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		rect.Ref(),
		from.Ref(),
	)

	return
}

// HasConvertPointFromNode returns true if the method "Document.convertPointFromNode" exists.
func (this Document) HasConvertPointFromNode() bool {
	return js.True == bindings.HasDocumentConvertPointFromNode(
		this.Ref(),
	)
}

// ConvertPointFromNodeFunc returns the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNodeFunc() (fn js.Func[func(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) DOMPoint]) {
	return fn.FromRef(
		bindings.DocumentConvertPointFromNodeFunc(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode calls the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNode(point DOMPointInit, from GeometryNode, options ConvertCoordinateOptions) (ret DOMPoint) {
	bindings.CallDocumentConvertPointFromNode(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvertPointFromNode1 returns true if the method "Document.convertPointFromNode" exists.
func (this Document) HasConvertPointFromNode1() bool {
	return js.True == bindings.HasDocumentConvertPointFromNode1(
		this.Ref(),
	)
}

// ConvertPointFromNode1Func returns the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNode1Func() (fn js.Func[func(point DOMPointInit, from GeometryNode) DOMPoint]) {
	return fn.FromRef(
		bindings.DocumentConvertPointFromNode1Func(
			this.Ref(),
		),
	)
}

// ConvertPointFromNode1 calls the method "Document.convertPointFromNode".
func (this Document) ConvertPointFromNode1(point DOMPointInit, from GeometryNode) (ret DOMPoint) {
	bindings.CallDocumentConvertPointFromNode1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&point),
		from.Ref(),
	)

	return
}

// HasGetElementById returns true if the method "Document.getElementById" exists.
func (this Document) HasGetElementById() bool {
	return js.True == bindings.HasDocumentGetElementById(
		this.Ref(),
	)
}

// GetElementByIdFunc returns the method "Document.getElementById".
func (this Document) GetElementByIdFunc() (fn js.Func[func(elementId js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentGetElementByIdFunc(
			this.Ref(),
		),
	)
}

// GetElementById calls the method "Document.getElementById".
func (this Document) GetElementById(elementId js.String) (ret Element) {
	bindings.CallDocumentGetElementById(
		this.Ref(), js.Pointer(&ret),
		elementId.Ref(),
	)

	return
}

// TryGetElementById calls the method "Document.getElementById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetElementById(elementId js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetElementById(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		elementId.Ref(),
	)

	return
}

// HasGetAnimations returns true if the method "Document.getAnimations" exists.
func (this Document) HasGetAnimations() bool {
	return js.True == bindings.HasDocumentGetAnimations(
		this.Ref(),
	)
}

// GetAnimationsFunc returns the method "Document.getAnimations".
func (this Document) GetAnimationsFunc() (fn js.Func[func() js.Array[Animation]]) {
	return fn.FromRef(
		bindings.DocumentGetAnimationsFunc(
			this.Ref(),
		),
	)
}

// GetAnimations calls the method "Document.getAnimations".
func (this Document) GetAnimations() (ret js.Array[Animation]) {
	bindings.CallDocumentGetAnimations(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAnimations calls the method "Document.getAnimations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryGetAnimations() (ret js.Array[Animation], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentGetAnimations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPrepend returns true if the method "Document.prepend" exists.
func (this Document) HasPrepend() bool {
	return js.True == bindings.HasDocumentPrepend(
		this.Ref(),
	)
}

// PrependFunc returns the method "Document.prepend".
func (this Document) PrependFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentPrependFunc(
			this.Ref(),
		),
	)
}

// Prepend calls the method "Document.prepend".
func (this Document) Prepend(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentPrepend(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasAppend returns true if the method "Document.append" exists.
func (this Document) HasAppend() bool {
	return js.True == bindings.HasDocumentAppend(
		this.Ref(),
	)
}

// AppendFunc returns the method "Document.append".
func (this Document) AppendFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentAppendFunc(
			this.Ref(),
		),
	)
}

// Append calls the method "Document.append".
func (this Document) Append(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentAppend(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasReplaceChildren returns true if the method "Document.replaceChildren" exists.
func (this Document) HasReplaceChildren() bool {
	return js.True == bindings.HasDocumentReplaceChildren(
		this.Ref(),
	)
}

// ReplaceChildrenFunc returns the method "Document.replaceChildren".
func (this Document) ReplaceChildrenFunc() (fn js.Func[func(nodes ...OneOf_Node_String)]) {
	return fn.FromRef(
		bindings.DocumentReplaceChildrenFunc(
			this.Ref(),
		),
	)
}

// ReplaceChildren calls the method "Document.replaceChildren".
func (this Document) ReplaceChildren(nodes ...OneOf_Node_String) (ret js.Void) {
	bindings.CallDocumentReplaceChildren(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(nodes),
		js.SizeU(len(nodes)),
	)

	return
}

// HasQuerySelector returns true if the method "Document.querySelector" exists.
func (this Document) HasQuerySelector() bool {
	return js.True == bindings.HasDocumentQuerySelector(
		this.Ref(),
	)
}

// QuerySelectorFunc returns the method "Document.querySelector".
func (this Document) QuerySelectorFunc() (fn js.Func[func(selectors js.String) Element]) {
	return fn.FromRef(
		bindings.DocumentQuerySelectorFunc(
			this.Ref(),
		),
	)
}

// QuerySelector calls the method "Document.querySelector".
func (this Document) QuerySelector(selectors js.String) (ret Element) {
	bindings.CallDocumentQuerySelector(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelector calls the method "Document.querySelector"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQuerySelector(selectors js.String) (ret Element, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQuerySelector(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasQuerySelectorAll returns true if the method "Document.querySelectorAll" exists.
func (this Document) HasQuerySelectorAll() bool {
	return js.True == bindings.HasDocumentQuerySelectorAll(
		this.Ref(),
	)
}

// QuerySelectorAllFunc returns the method "Document.querySelectorAll".
func (this Document) QuerySelectorAllFunc() (fn js.Func[func(selectors js.String) NodeList]) {
	return fn.FromRef(
		bindings.DocumentQuerySelectorAllFunc(
			this.Ref(),
		),
	)
}

// QuerySelectorAll calls the method "Document.querySelectorAll".
func (this Document) QuerySelectorAll(selectors js.String) (ret NodeList) {
	bindings.CallDocumentQuerySelectorAll(
		this.Ref(), js.Pointer(&ret),
		selectors.Ref(),
	)

	return
}

// TryQuerySelectorAll calls the method "Document.querySelectorAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryQuerySelectorAll(selectors js.String) (ret NodeList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentQuerySelectorAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		selectors.Ref(),
	)

	return
}

// HasCreateExpression returns true if the method "Document.createExpression" exists.
func (this Document) HasCreateExpression() bool {
	return js.True == bindings.HasDocumentCreateExpression(
		this.Ref(),
	)
}

// CreateExpressionFunc returns the method "Document.createExpression".
func (this Document) CreateExpressionFunc() (fn js.Func[func(expression js.String, resolver js.Func[func(prefix js.String) js.String]) XPathExpression]) {
	return fn.FromRef(
		bindings.DocumentCreateExpressionFunc(
			this.Ref(),
		),
	)
}

// CreateExpression calls the method "Document.createExpression".
func (this Document) CreateExpression(expression js.String, resolver js.Func[func(prefix js.String) js.String]) (ret XPathExpression) {
	bindings.CallDocumentCreateExpression(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		resolver.Ref(),
	)

	return
}

// HasCreateExpression1 returns true if the method "Document.createExpression" exists.
func (this Document) HasCreateExpression1() bool {
	return js.True == bindings.HasDocumentCreateExpression1(
		this.Ref(),
	)
}

// CreateExpression1Func returns the method "Document.createExpression".
func (this Document) CreateExpression1Func() (fn js.Func[func(expression js.String) XPathExpression]) {
	return fn.FromRef(
		bindings.DocumentCreateExpression1Func(
			this.Ref(),
		),
	)
}

// CreateExpression1 calls the method "Document.createExpression".
func (this Document) CreateExpression1(expression js.String) (ret XPathExpression) {
	bindings.CallDocumentCreateExpression1(
		this.Ref(), js.Pointer(&ret),
		expression.Ref(),
	)

	return
}

// TryCreateExpression1 calls the method "Document.createExpression"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateExpression1(expression js.String) (ret XPathExpression, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateExpression1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
	)

	return
}

// HasCreateNSResolver returns true if the method "Document.createNSResolver" exists.
func (this Document) HasCreateNSResolver() bool {
	return js.True == bindings.HasDocumentCreateNSResolver(
		this.Ref(),
	)
}

// CreateNSResolverFunc returns the method "Document.createNSResolver".
func (this Document) CreateNSResolverFunc() (fn js.Func[func(nodeResolver Node) Node]) {
	return fn.FromRef(
		bindings.DocumentCreateNSResolverFunc(
			this.Ref(),
		),
	)
}

// CreateNSResolver calls the method "Document.createNSResolver".
func (this Document) CreateNSResolver(nodeResolver Node) (ret Node) {
	bindings.CallDocumentCreateNSResolver(
		this.Ref(), js.Pointer(&ret),
		nodeResolver.Ref(),
	)

	return
}

// TryCreateNSResolver calls the method "Document.createNSResolver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Document) TryCreateNSResolver(nodeResolver Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDocumentCreateNSResolver(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		nodeResolver.Ref(),
	)

	return
}

// HasEvaluate returns true if the method "Document.evaluate" exists.
func (this Document) HasEvaluate() bool {
	return js.True == bindings.HasDocumentEvaluate(
		this.Ref(),
	)
}

// EvaluateFunc returns the method "Document.evaluate".
func (this Document) EvaluateFunc() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluateFunc(
			this.Ref(),
		),
	)
}

// Evaluate calls the method "Document.evaluate".
func (this Document) Evaluate(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16, result XPathResult) (ret XPathResult) {
	bindings.CallDocumentEvaluate(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
		result.Ref(),
	)

	return
}

// HasEvaluate1 returns true if the method "Document.evaluate" exists.
func (this Document) HasEvaluate1() bool {
	return js.True == bindings.HasDocumentEvaluate1(
		this.Ref(),
	)
}

// Evaluate1Func returns the method "Document.evaluate".
func (this Document) Evaluate1Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate1Func(
			this.Ref(),
		),
	)
}

// Evaluate1 calls the method "Document.evaluate".
func (this Document) Evaluate1(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String], typ uint16) (ret XPathResult) {
	bindings.CallDocumentEvaluate1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
		uint32(typ),
	)

	return
}

// HasEvaluate2 returns true if the method "Document.evaluate" exists.
func (this Document) HasEvaluate2() bool {
	return js.True == bindings.HasDocumentEvaluate2(
		this.Ref(),
	)
}

// Evaluate2Func returns the method "Document.evaluate".
func (this Document) Evaluate2Func() (fn js.Func[func(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate2Func(
			this.Ref(),
		),
	)
}

// Evaluate2 calls the method "Document.evaluate".
func (this Document) Evaluate2(expression js.String, contextNode Node, resolver js.Func[func(prefix js.String) js.String]) (ret XPathResult) {
	bindings.CallDocumentEvaluate2(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
		resolver.Ref(),
	)

	return
}

// HasEvaluate3 returns true if the method "Document.evaluate" exists.
func (this Document) HasEvaluate3() bool {
	return js.True == bindings.HasDocumentEvaluate3(
		this.Ref(),
	)
}

// Evaluate3Func returns the method "Document.evaluate".
func (this Document) Evaluate3Func() (fn js.Func[func(expression js.String, contextNode Node) XPathResult]) {
	return fn.FromRef(
		bindings.DocumentEvaluate3Func(
			this.Ref(),
		),
	)
}

// Evaluate3 calls the method "Document.evaluate".
func (this Document) Evaluate3(expression js.String, contextNode Node) (ret XPathResult) {
	bindings.CallDocumentEvaluate3(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		expression.Ref(),
		contextNode.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Node) NodeType() (ret uint16, ok bool) {
	ok = js.True == bindings.GetNodeNodeType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NodeName returns the value of property "Node.nodeName".
//
// It returns ok=false if there is no such property.
func (this Node) NodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeNodeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BaseURI returns the value of property "Node.baseURI".
//
// It returns ok=false if there is no such property.
func (this Node) BaseURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeBaseURI(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsConnected returns the value of property "Node.isConnected".
//
// It returns ok=false if there is no such property.
func (this Node) IsConnected() (ret bool, ok bool) {
	ok = js.True == bindings.GetNodeIsConnected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OwnerDocument returns the value of property "Node.ownerDocument".
//
// It returns ok=false if there is no such property.
func (this Node) OwnerDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetNodeOwnerDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ParentNode returns the value of property "Node.parentNode".
//
// It returns ok=false if there is no such property.
func (this Node) ParentNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeParentNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ParentElement returns the value of property "Node.parentElement".
//
// It returns ok=false if there is no such property.
func (this Node) ParentElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetNodeParentElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChildNodes returns the value of property "Node.childNodes".
//
// It returns ok=false if there is no such property.
func (this Node) ChildNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetNodeChildNodes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstChild returns the value of property "Node.firstChild".
//
// It returns ok=false if there is no such property.
func (this Node) FirstChild() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeFirstChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastChild returns the value of property "Node.lastChild".
//
// It returns ok=false if there is no such property.
func (this Node) LastChild() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeLastChild(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "Node.previousSibling".
//
// It returns ok=false if there is no such property.
func (this Node) PreviousSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodePreviousSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "Node.nextSibling".
//
// It returns ok=false if there is no such property.
func (this Node) NextSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetNodeNextSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NodeValue returns the value of property "Node.nodeValue".
//
// It returns ok=false if there is no such property.
func (this Node) NodeValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeNodeValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNodeValue sets the value of property "Node.nodeValue" to val.
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
// It returns ok=false if there is no such property.
func (this Node) TextContent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNodeTextContent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTextContent sets the value of property "Node.textContent" to val.
//
// It returns false if the property cannot be set.
func (this Node) SetTextContent(val js.String) bool {
	return js.True == bindings.SetNodeTextContent(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetRootNode returns true if the method "Node.getRootNode" exists.
func (this Node) HasGetRootNode() bool {
	return js.True == bindings.HasNodeGetRootNode(
		this.Ref(),
	)
}

// GetRootNodeFunc returns the method "Node.getRootNode".
func (this Node) GetRootNodeFunc() (fn js.Func[func(options GetRootNodeOptions) Node]) {
	return fn.FromRef(
		bindings.NodeGetRootNodeFunc(
			this.Ref(),
		),
	)
}

// GetRootNode calls the method "Node.getRootNode".
func (this Node) GetRootNode(options GetRootNodeOptions) (ret Node) {
	bindings.CallNodeGetRootNode(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetRootNode calls the method "Node.getRootNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryGetRootNode(options GetRootNodeOptions) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeGetRootNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetRootNode1 returns true if the method "Node.getRootNode" exists.
func (this Node) HasGetRootNode1() bool {
	return js.True == bindings.HasNodeGetRootNode1(
		this.Ref(),
	)
}

// GetRootNode1Func returns the method "Node.getRootNode".
func (this Node) GetRootNode1Func() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeGetRootNode1Func(
			this.Ref(),
		),
	)
}

// GetRootNode1 calls the method "Node.getRootNode".
func (this Node) GetRootNode1() (ret Node) {
	bindings.CallNodeGetRootNode1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetRootNode1 calls the method "Node.getRootNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryGetRootNode1() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeGetRootNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasHasChildNodes returns true if the method "Node.hasChildNodes" exists.
func (this Node) HasHasChildNodes() bool {
	return js.True == bindings.HasNodeHasChildNodes(
		this.Ref(),
	)
}

// HasChildNodesFunc returns the method "Node.hasChildNodes".
func (this Node) HasChildNodesFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NodeHasChildNodesFunc(
			this.Ref(),
		),
	)
}

// HasChildNodes calls the method "Node.hasChildNodes".
func (this Node) HasChildNodes() (ret bool) {
	bindings.CallNodeHasChildNodes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHasChildNodes calls the method "Node.hasChildNodes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryHasChildNodes() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeHasChildNodes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasNormalize returns true if the method "Node.normalize" exists.
func (this Node) HasNormalize() bool {
	return js.True == bindings.HasNodeNormalize(
		this.Ref(),
	)
}

// NormalizeFunc returns the method "Node.normalize".
func (this Node) NormalizeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NodeNormalizeFunc(
			this.Ref(),
		),
	)
}

// Normalize calls the method "Node.normalize".
func (this Node) Normalize() (ret js.Void) {
	bindings.CallNodeNormalize(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryNormalize calls the method "Node.normalize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryNormalize() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeNormalize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCloneNode returns true if the method "Node.cloneNode" exists.
func (this Node) HasCloneNode() bool {
	return js.True == bindings.HasNodeCloneNode(
		this.Ref(),
	)
}

// CloneNodeFunc returns the method "Node.cloneNode".
func (this Node) CloneNodeFunc() (fn js.Func[func(deep bool) Node]) {
	return fn.FromRef(
		bindings.NodeCloneNodeFunc(
			this.Ref(),
		),
	)
}

// CloneNode calls the method "Node.cloneNode".
func (this Node) CloneNode(deep bool) (ret Node) {
	bindings.CallNodeCloneNode(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(deep)),
	)

	return
}

// TryCloneNode calls the method "Node.cloneNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCloneNode(deep bool) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCloneNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(deep)),
	)

	return
}

// HasCloneNode1 returns true if the method "Node.cloneNode" exists.
func (this Node) HasCloneNode1() bool {
	return js.True == bindings.HasNodeCloneNode1(
		this.Ref(),
	)
}

// CloneNode1Func returns the method "Node.cloneNode".
func (this Node) CloneNode1Func() (fn js.Func[func() Node]) {
	return fn.FromRef(
		bindings.NodeCloneNode1Func(
			this.Ref(),
		),
	)
}

// CloneNode1 calls the method "Node.cloneNode".
func (this Node) CloneNode1() (ret Node) {
	bindings.CallNodeCloneNode1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCloneNode1 calls the method "Node.cloneNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCloneNode1() (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCloneNode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsEqualNode returns true if the method "Node.isEqualNode" exists.
func (this Node) HasIsEqualNode() bool {
	return js.True == bindings.HasNodeIsEqualNode(
		this.Ref(),
	)
}

// IsEqualNodeFunc returns the method "Node.isEqualNode".
func (this Node) IsEqualNodeFunc() (fn js.Func[func(otherNode Node) bool]) {
	return fn.FromRef(
		bindings.NodeIsEqualNodeFunc(
			this.Ref(),
		),
	)
}

// IsEqualNode calls the method "Node.isEqualNode".
func (this Node) IsEqualNode(otherNode Node) (ret bool) {
	bindings.CallNodeIsEqualNode(
		this.Ref(), js.Pointer(&ret),
		otherNode.Ref(),
	)

	return
}

// TryIsEqualNode calls the method "Node.isEqualNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsEqualNode(otherNode Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsEqualNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		otherNode.Ref(),
	)

	return
}

// HasIsSameNode returns true if the method "Node.isSameNode" exists.
func (this Node) HasIsSameNode() bool {
	return js.True == bindings.HasNodeIsSameNode(
		this.Ref(),
	)
}

// IsSameNodeFunc returns the method "Node.isSameNode".
func (this Node) IsSameNodeFunc() (fn js.Func[func(otherNode Node) bool]) {
	return fn.FromRef(
		bindings.NodeIsSameNodeFunc(
			this.Ref(),
		),
	)
}

// IsSameNode calls the method "Node.isSameNode".
func (this Node) IsSameNode(otherNode Node) (ret bool) {
	bindings.CallNodeIsSameNode(
		this.Ref(), js.Pointer(&ret),
		otherNode.Ref(),
	)

	return
}

// TryIsSameNode calls the method "Node.isSameNode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsSameNode(otherNode Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsSameNode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		otherNode.Ref(),
	)

	return
}

// HasCompareDocumentPosition returns true if the method "Node.compareDocumentPosition" exists.
func (this Node) HasCompareDocumentPosition() bool {
	return js.True == bindings.HasNodeCompareDocumentPosition(
		this.Ref(),
	)
}

// CompareDocumentPositionFunc returns the method "Node.compareDocumentPosition".
func (this Node) CompareDocumentPositionFunc() (fn js.Func[func(other Node) uint16]) {
	return fn.FromRef(
		bindings.NodeCompareDocumentPositionFunc(
			this.Ref(),
		),
	)
}

// CompareDocumentPosition calls the method "Node.compareDocumentPosition".
func (this Node) CompareDocumentPosition(other Node) (ret uint16) {
	bindings.CallNodeCompareDocumentPosition(
		this.Ref(), js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryCompareDocumentPosition calls the method "Node.compareDocumentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryCompareDocumentPosition(other Node) (ret uint16, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeCompareDocumentPosition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasContains returns true if the method "Node.contains" exists.
func (this Node) HasContains() bool {
	return js.True == bindings.HasNodeContains(
		this.Ref(),
	)
}

// ContainsFunc returns the method "Node.contains".
func (this Node) ContainsFunc() (fn js.Func[func(other Node) bool]) {
	return fn.FromRef(
		bindings.NodeContainsFunc(
			this.Ref(),
		),
	)
}

// Contains calls the method "Node.contains".
func (this Node) Contains(other Node) (ret bool) {
	bindings.CallNodeContains(
		this.Ref(), js.Pointer(&ret),
		other.Ref(),
	)

	return
}

// TryContains calls the method "Node.contains"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryContains(other Node) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeContains(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		other.Ref(),
	)

	return
}

// HasLookupPrefix returns true if the method "Node.lookupPrefix" exists.
func (this Node) HasLookupPrefix() bool {
	return js.True == bindings.HasNodeLookupPrefix(
		this.Ref(),
	)
}

// LookupPrefixFunc returns the method "Node.lookupPrefix".
func (this Node) LookupPrefixFunc() (fn js.Func[func(namespace js.String) js.String]) {
	return fn.FromRef(
		bindings.NodeLookupPrefixFunc(
			this.Ref(),
		),
	)
}

// LookupPrefix calls the method "Node.lookupPrefix".
func (this Node) LookupPrefix(namespace js.String) (ret js.String) {
	bindings.CallNodeLookupPrefix(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
	)

	return
}

// TryLookupPrefix calls the method "Node.lookupPrefix"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryLookupPrefix(namespace js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeLookupPrefix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
	)

	return
}

// HasLookupNamespaceURI returns true if the method "Node.lookupNamespaceURI" exists.
func (this Node) HasLookupNamespaceURI() bool {
	return js.True == bindings.HasNodeLookupNamespaceURI(
		this.Ref(),
	)
}

// LookupNamespaceURIFunc returns the method "Node.lookupNamespaceURI".
func (this Node) LookupNamespaceURIFunc() (fn js.Func[func(prefix js.String) js.String]) {
	return fn.FromRef(
		bindings.NodeLookupNamespaceURIFunc(
			this.Ref(),
		),
	)
}

// LookupNamespaceURI calls the method "Node.lookupNamespaceURI".
func (this Node) LookupNamespaceURI(prefix js.String) (ret js.String) {
	bindings.CallNodeLookupNamespaceURI(
		this.Ref(), js.Pointer(&ret),
		prefix.Ref(),
	)

	return
}

// TryLookupNamespaceURI calls the method "Node.lookupNamespaceURI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryLookupNamespaceURI(prefix js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeLookupNamespaceURI(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		prefix.Ref(),
	)

	return
}

// HasIsDefaultNamespace returns true if the method "Node.isDefaultNamespace" exists.
func (this Node) HasIsDefaultNamespace() bool {
	return js.True == bindings.HasNodeIsDefaultNamespace(
		this.Ref(),
	)
}

// IsDefaultNamespaceFunc returns the method "Node.isDefaultNamespace".
func (this Node) IsDefaultNamespaceFunc() (fn js.Func[func(namespace js.String) bool]) {
	return fn.FromRef(
		bindings.NodeIsDefaultNamespaceFunc(
			this.Ref(),
		),
	)
}

// IsDefaultNamespace calls the method "Node.isDefaultNamespace".
func (this Node) IsDefaultNamespace(namespace js.String) (ret bool) {
	bindings.CallNodeIsDefaultNamespace(
		this.Ref(), js.Pointer(&ret),
		namespace.Ref(),
	)

	return
}

// TryIsDefaultNamespace calls the method "Node.isDefaultNamespace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryIsDefaultNamespace(namespace js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeIsDefaultNamespace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		namespace.Ref(),
	)

	return
}

// HasInsertBefore returns true if the method "Node.insertBefore" exists.
func (this Node) HasInsertBefore() bool {
	return js.True == bindings.HasNodeInsertBefore(
		this.Ref(),
	)
}

// InsertBeforeFunc returns the method "Node.insertBefore".
func (this Node) InsertBeforeFunc() (fn js.Func[func(node Node, child Node) Node]) {
	return fn.FromRef(
		bindings.NodeInsertBeforeFunc(
			this.Ref(),
		),
	)
}

// InsertBefore calls the method "Node.insertBefore".
func (this Node) InsertBefore(node Node, child Node) (ret Node) {
	bindings.CallNodeInsertBefore(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		child.Ref(),
	)

	return
}

// HasAppendChild returns true if the method "Node.appendChild" exists.
func (this Node) HasAppendChild() bool {
	return js.True == bindings.HasNodeAppendChild(
		this.Ref(),
	)
}

// AppendChildFunc returns the method "Node.appendChild".
func (this Node) AppendChildFunc() (fn js.Func[func(node Node) Node]) {
	return fn.FromRef(
		bindings.NodeAppendChildFunc(
			this.Ref(),
		),
	)
}

// AppendChild calls the method "Node.appendChild".
func (this Node) AppendChild(node Node) (ret Node) {
	bindings.CallNodeAppendChild(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryAppendChild calls the method "Node.appendChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryAppendChild(node Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeAppendChild(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

// HasReplaceChild returns true if the method "Node.replaceChild" exists.
func (this Node) HasReplaceChild() bool {
	return js.True == bindings.HasNodeReplaceChild(
		this.Ref(),
	)
}

// ReplaceChildFunc returns the method "Node.replaceChild".
func (this Node) ReplaceChildFunc() (fn js.Func[func(node Node, child Node) Node]) {
	return fn.FromRef(
		bindings.NodeReplaceChildFunc(
			this.Ref(),
		),
	)
}

// ReplaceChild calls the method "Node.replaceChild".
func (this Node) ReplaceChild(node Node, child Node) (ret Node) {
	bindings.CallNodeReplaceChild(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
		child.Ref(),
	)

	return
}

// HasRemoveChild returns true if the method "Node.removeChild" exists.
func (this Node) HasRemoveChild() bool {
	return js.True == bindings.HasNodeRemoveChild(
		this.Ref(),
	)
}

// RemoveChildFunc returns the method "Node.removeChild".
func (this Node) RemoveChildFunc() (fn js.Func[func(child Node) Node]) {
	return fn.FromRef(
		bindings.NodeRemoveChildFunc(
			this.Ref(),
		),
	)
}

// RemoveChild calls the method "Node.removeChild".
func (this Node) RemoveChild(child Node) (ret Node) {
	bindings.CallNodeRemoveChild(
		this.Ref(), js.Pointer(&ret),
		child.Ref(),
	)

	return
}

// TryRemoveChild calls the method "Node.removeChild"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Node) TryRemoveChild(child Node) (ret Node, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNodeRemoveChild(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		child.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this AbstractRange) StartContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetAbstractRangeStartContainer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StartOffset returns the value of property "AbstractRange.startOffset".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) StartOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAbstractRangeStartOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EndContainer returns the value of property "AbstractRange.endContainer".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) EndContainer() (ret Node, ok bool) {
	ok = js.True == bindings.GetAbstractRangeEndContainer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EndOffset returns the value of property "AbstractRange.endOffset".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) EndOffset() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAbstractRangeEndOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Collapsed returns the value of property "AbstractRange.collapsed".
//
// It returns ok=false if there is no such property.
func (this AbstractRange) Collapsed() (ret bool, ok bool) {
	ok = js.True == bindings.GetAbstractRangeCollapsed(
		this.Ref(), js.Pointer(&ret),
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
// It returns ok=false if there is no such property.
func (this Accelerometer) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "Accelerometer.y".
//
// It returns ok=false if there is no such property.
func (this Accelerometer) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "Accelerometer.z".
//
// It returns ok=false if there is no such property.
func (this Accelerometer) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetAccelerometerZ(
		this.Ref(), js.Pointer(&ret),
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

// New creates a new AdRender in the application heap.
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

// New creates a new AddressErrors in the application heap.
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

// New creates a new AddressInit in the application heap.
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

// New creates a new AesCbcParams in the application heap.
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

// New creates a new AesCtrParams in the application heap.
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

// New creates a new AesDerivedKeyParams in the application heap.
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

// New creates a new AesKeyAlgorithm in the application heap.
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

// New creates a new AesKeyGenParams in the application heap.
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

// New creates a new Algorithm in the application heap.
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

// New creates a new AllowedBluetoothDevice in the application heap.
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

// New creates a new AllowedUSBDevice in the application heap.
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

// New creates a new AmbientLightReadingValues in the application heap.
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Illuminance returns the value of property "AmbientLightSensor.illuminance".
//
// It returns ok=false if there is no such property.
func (this AmbientLightSensor) Illuminance() (ret float64, ok bool) {
	ok = js.True == bindings.GetAmbientLightSensorIlluminance(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type BiquadFilterType uint32
