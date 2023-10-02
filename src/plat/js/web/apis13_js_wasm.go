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

const (
	_ ImageSmoothingQuality = iota

	ImageSmoothingQuality_LOW
	ImageSmoothingQuality_MEDIUM
	ImageSmoothingQuality_HIGH
)

func (ImageSmoothingQuality) FromRef(str js.Ref) ImageSmoothingQuality {
	return ImageSmoothingQuality(bindings.ConstOfImageSmoothingQuality(str))
}

func (x ImageSmoothingQuality) String() (string, bool) {
	switch x {
	case ImageSmoothingQuality_LOW:
		return "low", true
	case ImageSmoothingQuality_MEDIUM:
		return "medium", true
	case ImageSmoothingQuality_HIGH:
		return "high", true
	default:
		return "", false
	}
}

type OneOf_String_CanvasGradient_CanvasPattern struct {
	ref js.Ref
}

func (x OneOf_String_CanvasGradient_CanvasPattern) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_CanvasGradient_CanvasPattern) Free() {
	x.ref.Free()
}

func (x OneOf_String_CanvasGradient_CanvasPattern) FromRef(ref js.Ref) OneOf_String_CanvasGradient_CanvasPattern {
	return OneOf_String_CanvasGradient_CanvasPattern{
		ref: ref,
	}
}

func (x OneOf_String_CanvasGradient_CanvasPattern) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_CanvasGradient_CanvasPattern) CanvasGradient() CanvasGradient {
	return CanvasGradient{}.FromRef(x.ref)
}

func (x OneOf_String_CanvasGradient_CanvasPattern) CanvasPattern() CanvasPattern {
	return CanvasPattern{}.FromRef(x.ref)
}

type CanvasLineCap uint32

const (
	_ CanvasLineCap = iota

	CanvasLineCap_BUTT
	CanvasLineCap_ROUND
	CanvasLineCap_SQUARE
)

func (CanvasLineCap) FromRef(str js.Ref) CanvasLineCap {
	return CanvasLineCap(bindings.ConstOfCanvasLineCap(str))
}

func (x CanvasLineCap) String() (string, bool) {
	switch x {
	case CanvasLineCap_BUTT:
		return "butt", true
	case CanvasLineCap_ROUND:
		return "round", true
	case CanvasLineCap_SQUARE:
		return "square", true
	default:
		return "", false
	}
}

type CanvasLineJoin uint32

const (
	_ CanvasLineJoin = iota

	CanvasLineJoin_ROUND
	CanvasLineJoin_BEVEL
	CanvasLineJoin_MITER
)

func (CanvasLineJoin) FromRef(str js.Ref) CanvasLineJoin {
	return CanvasLineJoin(bindings.ConstOfCanvasLineJoin(str))
}

func (x CanvasLineJoin) String() (string, bool) {
	switch x {
	case CanvasLineJoin_ROUND:
		return "round", true
	case CanvasLineJoin_BEVEL:
		return "bevel", true
	case CanvasLineJoin_MITER:
		return "miter", true
	default:
		return "", false
	}
}

type CanvasTextAlign uint32

const (
	_ CanvasTextAlign = iota

	CanvasTextAlign_START
	CanvasTextAlign_END
	CanvasTextAlign_LEFT
	CanvasTextAlign_RIGHT
	CanvasTextAlign_CENTER
)

func (CanvasTextAlign) FromRef(str js.Ref) CanvasTextAlign {
	return CanvasTextAlign(bindings.ConstOfCanvasTextAlign(str))
}

func (x CanvasTextAlign) String() (string, bool) {
	switch x {
	case CanvasTextAlign_START:
		return "start", true
	case CanvasTextAlign_END:
		return "end", true
	case CanvasTextAlign_LEFT:
		return "left", true
	case CanvasTextAlign_RIGHT:
		return "right", true
	case CanvasTextAlign_CENTER:
		return "center", true
	default:
		return "", false
	}
}

type CanvasTextBaseline uint32

const (
	_ CanvasTextBaseline = iota

	CanvasTextBaseline_TOP
	CanvasTextBaseline_HANGING
	CanvasTextBaseline_MIDDLE
	CanvasTextBaseline_ALPHABETIC
	CanvasTextBaseline_IDEOGRAPHIC
	CanvasTextBaseline_BOTTOM
)

func (CanvasTextBaseline) FromRef(str js.Ref) CanvasTextBaseline {
	return CanvasTextBaseline(bindings.ConstOfCanvasTextBaseline(str))
}

func (x CanvasTextBaseline) String() (string, bool) {
	switch x {
	case CanvasTextBaseline_TOP:
		return "top", true
	case CanvasTextBaseline_HANGING:
		return "hanging", true
	case CanvasTextBaseline_MIDDLE:
		return "middle", true
	case CanvasTextBaseline_ALPHABETIC:
		return "alphabetic", true
	case CanvasTextBaseline_IDEOGRAPHIC:
		return "ideographic", true
	case CanvasTextBaseline_BOTTOM:
		return "bottom", true
	default:
		return "", false
	}
}

type CanvasDirection uint32

const (
	_ CanvasDirection = iota

	CanvasDirection_LTR
	CanvasDirection_RTL
	CanvasDirection_INHERIT
)

func (CanvasDirection) FromRef(str js.Ref) CanvasDirection {
	return CanvasDirection(bindings.ConstOfCanvasDirection(str))
}

func (x CanvasDirection) String() (string, bool) {
	switch x {
	case CanvasDirection_LTR:
		return "ltr", true
	case CanvasDirection_RTL:
		return "rtl", true
	case CanvasDirection_INHERIT:
		return "inherit", true
	default:
		return "", false
	}
}

type CanvasFontKerning uint32

const (
	_ CanvasFontKerning = iota

	CanvasFontKerning_AUTO
	CanvasFontKerning_NORMAL
	CanvasFontKerning_NONE
)

func (CanvasFontKerning) FromRef(str js.Ref) CanvasFontKerning {
	return CanvasFontKerning(bindings.ConstOfCanvasFontKerning(str))
}

func (x CanvasFontKerning) String() (string, bool) {
	switch x {
	case CanvasFontKerning_AUTO:
		return "auto", true
	case CanvasFontKerning_NORMAL:
		return "normal", true
	case CanvasFontKerning_NONE:
		return "none", true
	default:
		return "", false
	}
}

type CanvasFontStretch uint32

const (
	_ CanvasFontStretch = iota

	CanvasFontStretch_ULTRA_CONDENSED
	CanvasFontStretch_EXTRA_CONDENSED
	CanvasFontStretch_CONDENSED
	CanvasFontStretch_SEMI_CONDENSED
	CanvasFontStretch_NORMAL
	CanvasFontStretch_SEMI_EXPANDED
	CanvasFontStretch_EXPANDED
	CanvasFontStretch_EXTRA_EXPANDED
	CanvasFontStretch_ULTRA_EXPANDED
)

func (CanvasFontStretch) FromRef(str js.Ref) CanvasFontStretch {
	return CanvasFontStretch(bindings.ConstOfCanvasFontStretch(str))
}

func (x CanvasFontStretch) String() (string, bool) {
	switch x {
	case CanvasFontStretch_ULTRA_CONDENSED:
		return "ultra-condensed", true
	case CanvasFontStretch_EXTRA_CONDENSED:
		return "extra-condensed", true
	case CanvasFontStretch_CONDENSED:
		return "condensed", true
	case CanvasFontStretch_SEMI_CONDENSED:
		return "semi-condensed", true
	case CanvasFontStretch_NORMAL:
		return "normal", true
	case CanvasFontStretch_SEMI_EXPANDED:
		return "semi-expanded", true
	case CanvasFontStretch_EXPANDED:
		return "expanded", true
	case CanvasFontStretch_EXTRA_EXPANDED:
		return "extra-expanded", true
	case CanvasFontStretch_ULTRA_EXPANDED:
		return "ultra-expanded", true
	default:
		return "", false
	}
}

type CanvasFontVariantCaps uint32

const (
	_ CanvasFontVariantCaps = iota

	CanvasFontVariantCaps_NORMAL
	CanvasFontVariantCaps_SMALL_CAPS
	CanvasFontVariantCaps_ALL_SMALL_CAPS
	CanvasFontVariantCaps_PETITE_CAPS
	CanvasFontVariantCaps_ALL_PETITE_CAPS
	CanvasFontVariantCaps_UNICASE
	CanvasFontVariantCaps_TITLING_CAPS
)

func (CanvasFontVariantCaps) FromRef(str js.Ref) CanvasFontVariantCaps {
	return CanvasFontVariantCaps(bindings.ConstOfCanvasFontVariantCaps(str))
}

func (x CanvasFontVariantCaps) String() (string, bool) {
	switch x {
	case CanvasFontVariantCaps_NORMAL:
		return "normal", true
	case CanvasFontVariantCaps_SMALL_CAPS:
		return "small-caps", true
	case CanvasFontVariantCaps_ALL_SMALL_CAPS:
		return "all-small-caps", true
	case CanvasFontVariantCaps_PETITE_CAPS:
		return "petite-caps", true
	case CanvasFontVariantCaps_ALL_PETITE_CAPS:
		return "all-petite-caps", true
	case CanvasFontVariantCaps_UNICASE:
		return "unicase", true
	case CanvasFontVariantCaps_TITLING_CAPS:
		return "titling-caps", true
	default:
		return "", false
	}
}

type CanvasTextRendering uint32

const (
	_ CanvasTextRendering = iota

	CanvasTextRendering_AUTO
	CanvasTextRendering_OPTIMIZE_SPEED
	CanvasTextRendering_OPTIMIZE_LEGIBILITY
	CanvasTextRendering_GEOMETRIC_PRECISION
)

func (CanvasTextRendering) FromRef(str js.Ref) CanvasTextRendering {
	return CanvasTextRendering(bindings.ConstOfCanvasTextRendering(str))
}

func (x CanvasTextRendering) String() (string, bool) {
	switch x {
	case CanvasTextRendering_AUTO:
		return "auto", true
	case CanvasTextRendering_OPTIMIZE_SPEED:
		return "optimizeSpeed", true
	case CanvasTextRendering_OPTIMIZE_LEGIBILITY:
		return "optimizeLegibility", true
	case CanvasTextRendering_GEOMETRIC_PRECISION:
		return "geometricPrecision", true
	default:
		return "", false
	}
}

type OffscreenCanvasRenderingContext2D struct {
	ref js.Ref
}

func (this OffscreenCanvasRenderingContext2D) Once() OffscreenCanvasRenderingContext2D {
	this.Ref().Once()
	return this
}

func (this OffscreenCanvasRenderingContext2D) Ref() js.Ref {
	return this.ref
}

func (this OffscreenCanvasRenderingContext2D) FromRef(ref js.Ref) OffscreenCanvasRenderingContext2D {
	this.ref = ref
	return this
}

func (this OffscreenCanvasRenderingContext2D) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "OffscreenCanvasRenderingContext2D.canvas".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) Canvas() (OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return OffscreenCanvas{}.FromRef(_ret), _ok
}

// GlobalAlpha returns the value of property "OffscreenCanvasRenderingContext2D.globalAlpha".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) GlobalAlpha() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DGlobalAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// GlobalAlpha sets the value of property "OffscreenCanvasRenderingContext2D.globalAlpha" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetGlobalAlpha(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DGlobalAlpha(
		this.Ref(),
		float64(val),
	)
}

// GlobalCompositeOperation returns the value of property "OffscreenCanvasRenderingContext2D.globalCompositeOperation".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) GlobalCompositeOperation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DGlobalCompositeOperation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GlobalCompositeOperation sets the value of property "OffscreenCanvasRenderingContext2D.globalCompositeOperation" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetGlobalCompositeOperation(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DGlobalCompositeOperation(
		this.Ref(),
		val.Ref(),
	)
}

// ImageSmoothingEnabled returns the value of property "OffscreenCanvasRenderingContext2D.imageSmoothingEnabled".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ImageSmoothingEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DImageSmoothingEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ImageSmoothingEnabled sets the value of property "OffscreenCanvasRenderingContext2D.imageSmoothingEnabled" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetImageSmoothingEnabled(val bool) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DImageSmoothingEnabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ImageSmoothingQuality returns the value of property "OffscreenCanvasRenderingContext2D.imageSmoothingQuality".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ImageSmoothingQuality() (ImageSmoothingQuality, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DImageSmoothingQuality(
		this.Ref(), js.Pointer(&_ok),
	)
	return ImageSmoothingQuality(_ret), _ok
}

// ImageSmoothingQuality sets the value of property "OffscreenCanvasRenderingContext2D.imageSmoothingQuality" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DImageSmoothingQuality(
		this.Ref(),
		uint32(val),
	)
}

// StrokeStyle returns the value of property "OffscreenCanvasRenderingContext2D.strokeStyle".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) StrokeStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DStrokeStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// StrokeStyle sets the value of property "OffscreenCanvasRenderingContext2D.strokeStyle" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetStrokeStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DStrokeStyle(
		this.Ref(),
		val.Ref(),
	)
}

// FillStyle returns the value of property "OffscreenCanvasRenderingContext2D.fillStyle".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) FillStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFillStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// FillStyle sets the value of property "OffscreenCanvasRenderingContext2D.fillStyle" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFillStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFillStyle(
		this.Ref(),
		val.Ref(),
	)
}

// ShadowOffsetX returns the value of property "OffscreenCanvasRenderingContext2D.shadowOffsetX".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ShadowOffsetX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DShadowOffsetX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetX sets the value of property "OffscreenCanvasRenderingContext2D.shadowOffsetX" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetShadowOffsetX(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DShadowOffsetX(
		this.Ref(),
		float64(val),
	)
}

// ShadowOffsetY returns the value of property "OffscreenCanvasRenderingContext2D.shadowOffsetY".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ShadowOffsetY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DShadowOffsetY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetY sets the value of property "OffscreenCanvasRenderingContext2D.shadowOffsetY" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetShadowOffsetY(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DShadowOffsetY(
		this.Ref(),
		float64(val),
	)
}

// ShadowBlur returns the value of property "OffscreenCanvasRenderingContext2D.shadowBlur".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ShadowBlur() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DShadowBlur(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowBlur sets the value of property "OffscreenCanvasRenderingContext2D.shadowBlur" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetShadowBlur(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DShadowBlur(
		this.Ref(),
		float64(val),
	)
}

// ShadowColor returns the value of property "OffscreenCanvasRenderingContext2D.shadowColor".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) ShadowColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DShadowColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ShadowColor sets the value of property "OffscreenCanvasRenderingContext2D.shadowColor" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetShadowColor(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DShadowColor(
		this.Ref(),
		val.Ref(),
	)
}

// Filter returns the value of property "OffscreenCanvasRenderingContext2D.filter".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) Filter() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFilter(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Filter sets the value of property "OffscreenCanvasRenderingContext2D.filter" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFilter(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFilter(
		this.Ref(),
		val.Ref(),
	)
}

// LineWidth returns the value of property "OffscreenCanvasRenderingContext2D.lineWidth".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) LineWidth() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DLineWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineWidth sets the value of property "OffscreenCanvasRenderingContext2D.lineWidth" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetLineWidth(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DLineWidth(
		this.Ref(),
		float64(val),
	)
}

// LineCap returns the value of property "OffscreenCanvasRenderingContext2D.lineCap".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) LineCap() (CanvasLineCap, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DLineCap(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineCap(_ret), _ok
}

// LineCap sets the value of property "OffscreenCanvasRenderingContext2D.lineCap" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetLineCap(val CanvasLineCap) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DLineCap(
		this.Ref(),
		uint32(val),
	)
}

// LineJoin returns the value of property "OffscreenCanvasRenderingContext2D.lineJoin".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) LineJoin() (CanvasLineJoin, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DLineJoin(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineJoin(_ret), _ok
}

// LineJoin sets the value of property "OffscreenCanvasRenderingContext2D.lineJoin" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetLineJoin(val CanvasLineJoin) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DLineJoin(
		this.Ref(),
		uint32(val),
	)
}

// MiterLimit returns the value of property "OffscreenCanvasRenderingContext2D.miterLimit".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) MiterLimit() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DMiterLimit(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MiterLimit sets the value of property "OffscreenCanvasRenderingContext2D.miterLimit" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetMiterLimit(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DMiterLimit(
		this.Ref(),
		float64(val),
	)
}

// LineDashOffset returns the value of property "OffscreenCanvasRenderingContext2D.lineDashOffset".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) LineDashOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DLineDashOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineDashOffset sets the value of property "OffscreenCanvasRenderingContext2D.lineDashOffset" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetLineDashOffset(val float64) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DLineDashOffset(
		this.Ref(),
		float64(val),
	)
}

// Font returns the value of property "OffscreenCanvasRenderingContext2D.font".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) Font() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFont(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Font sets the value of property "OffscreenCanvasRenderingContext2D.font" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFont(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFont(
		this.Ref(),
		val.Ref(),
	)
}

// TextAlign returns the value of property "OffscreenCanvasRenderingContext2D.textAlign".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) TextAlign() (CanvasTextAlign, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DTextAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextAlign(_ret), _ok
}

// TextAlign sets the value of property "OffscreenCanvasRenderingContext2D.textAlign" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetTextAlign(val CanvasTextAlign) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DTextAlign(
		this.Ref(),
		uint32(val),
	)
}

// TextBaseline returns the value of property "OffscreenCanvasRenderingContext2D.textBaseline".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) TextBaseline() (CanvasTextBaseline, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DTextBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextBaseline(_ret), _ok
}

// TextBaseline sets the value of property "OffscreenCanvasRenderingContext2D.textBaseline" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetTextBaseline(val CanvasTextBaseline) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DTextBaseline(
		this.Ref(),
		uint32(val),
	)
}

// Direction returns the value of property "OffscreenCanvasRenderingContext2D.direction".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) Direction() (CanvasDirection, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasDirection(_ret), _ok
}

// Direction sets the value of property "OffscreenCanvasRenderingContext2D.direction" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetDirection(val CanvasDirection) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DDirection(
		this.Ref(),
		uint32(val),
	)
}

// LetterSpacing returns the value of property "OffscreenCanvasRenderingContext2D.letterSpacing".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) LetterSpacing() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DLetterSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LetterSpacing sets the value of property "OffscreenCanvasRenderingContext2D.letterSpacing" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetLetterSpacing(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DLetterSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// FontKerning returns the value of property "OffscreenCanvasRenderingContext2D.fontKerning".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) FontKerning() (CanvasFontKerning, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFontKerning(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontKerning(_ret), _ok
}

// FontKerning sets the value of property "OffscreenCanvasRenderingContext2D.fontKerning" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFontKerning(val CanvasFontKerning) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFontKerning(
		this.Ref(),
		uint32(val),
	)
}

// FontStretch returns the value of property "OffscreenCanvasRenderingContext2D.fontStretch".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) FontStretch() (CanvasFontStretch, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFontStretch(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontStretch(_ret), _ok
}

// FontStretch sets the value of property "OffscreenCanvasRenderingContext2D.fontStretch" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFontStretch(val CanvasFontStretch) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFontStretch(
		this.Ref(),
		uint32(val),
	)
}

// FontVariantCaps returns the value of property "OffscreenCanvasRenderingContext2D.fontVariantCaps".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) FontVariantCaps() (CanvasFontVariantCaps, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DFontVariantCaps(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontVariantCaps(_ret), _ok
}

// FontVariantCaps sets the value of property "OffscreenCanvasRenderingContext2D.fontVariantCaps" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetFontVariantCaps(val CanvasFontVariantCaps) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DFontVariantCaps(
		this.Ref(),
		uint32(val),
	)
}

// TextRendering returns the value of property "OffscreenCanvasRenderingContext2D.textRendering".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) TextRendering() (CanvasTextRendering, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DTextRendering(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextRendering(_ret), _ok
}

// TextRendering sets the value of property "OffscreenCanvasRenderingContext2D.textRendering" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetTextRendering(val CanvasTextRendering) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DTextRendering(
		this.Ref(),
		uint32(val),
	)
}

// WordSpacing returns the value of property "OffscreenCanvasRenderingContext2D.wordSpacing".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvasRenderingContext2D) WordSpacing() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasRenderingContext2DWordSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// WordSpacing sets the value of property "OffscreenCanvasRenderingContext2D.wordSpacing" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvasRenderingContext2D) SetWordSpacing(val js.String) bool {
	return js.True == bindings.SetOffscreenCanvasRenderingContext2DWordSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// Commit calls the method "OffscreenCanvasRenderingContext2D.commit".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Commit() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCommit(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CommitFunc returns the method "OffscreenCanvasRenderingContext2D.commit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CommitFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCommitFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "OffscreenCanvasRenderingContext2D.scale".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Scale(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DScale(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScaleFunc returns the method "OffscreenCanvasRenderingContext2D.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ScaleFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DScaleFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "OffscreenCanvasRenderingContext2D.rotate".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Rotate(angle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DRotate(
		this.Ref(), js.Pointer(&_ok),
		float64(angle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RotateFunc returns the method "OffscreenCanvasRenderingContext2D.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RotateFunc() (fn js.Func[func(angle float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DRotateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "OffscreenCanvasRenderingContext2D.translate".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Translate(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TranslateFunc returns the method "OffscreenCanvasRenderingContext2D.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) TranslateFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DTranslateFunc(
			this.Ref(),
		),
	)
}

// Transform calls the method "OffscreenCanvasRenderingContext2D.transform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DTransform(
		this.Ref(), js.Pointer(&_ok),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TransformFunc returns the method "OffscreenCanvasRenderingContext2D.transform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) TransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DTransformFunc(
			this.Ref(),
		),
	)
}

// GetTransform calls the method "OffscreenCanvasRenderingContext2D.getTransform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetTransform() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// GetTransformFunc returns the method "OffscreenCanvasRenderingContext2D.getTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetTransformFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DGetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&_ok),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransformFunc returns the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform1Func returns the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransform1Func() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform2 calls the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransform2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform2Func returns the method "OffscreenCanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetTransform2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DSetTransform2Func(
			this.Ref(),
		),
	)
}

// ResetTransform calls the method "OffscreenCanvasRenderingContext2D.resetTransform".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ResetTransform() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetTransformFunc returns the method "OffscreenCanvasRenderingContext2D.resetTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ResetTransformFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DResetTransformFunc(
			this.Ref(),
		),
	)
}

// CreateLinearGradient calls the method "OffscreenCanvasRenderingContext2D.createLinearGradient".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateLinearGradientFunc returns the method "OffscreenCanvasRenderingContext2D.createLinearGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateLinearGradientFunc() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateLinearGradientFunc(
			this.Ref(),
		),
	)
}

// CreateRadialGradient calls the method "OffscreenCanvasRenderingContext2D.createRadialGradient".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateRadialGradientFunc returns the method "OffscreenCanvasRenderingContext2D.createRadialGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateRadialGradientFunc() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateRadialGradientFunc(
			this.Ref(),
		),
	)
}

// CreateConicGradient calls the method "OffscreenCanvasRenderingContext2D.createConicGradient".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateConicGradientFunc returns the method "OffscreenCanvasRenderingContext2D.createConicGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateConicGradientFunc() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateConicGradientFunc(
			this.Ref(),
		),
	)
}

// CreatePattern calls the method "OffscreenCanvasRenderingContext2D.createPattern".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		repetition.Ref(),
	)

	return CanvasPattern{}.FromRef(_ret), _ok
}

// CreatePatternFunc returns the method "OffscreenCanvasRenderingContext2D.createPattern".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreatePatternFunc() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreatePatternFunc(
			this.Ref(),
		),
	)
}

// ClearRect calls the method "OffscreenCanvasRenderingContext2D.clearRect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearRectFunc returns the method "OffscreenCanvasRenderingContext2D.clearRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ClearRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClearRectFunc(
			this.Ref(),
		),
	)
}

// FillRect calls the method "OffscreenCanvasRenderingContext2D.fillRect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillRectFunc returns the method "OffscreenCanvasRenderingContext2D.fillRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFillRectFunc(
			this.Ref(),
		),
	)
}

// StrokeRect calls the method "OffscreenCanvasRenderingContext2D.strokeRect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeRectFunc returns the method "OffscreenCanvasRenderingContext2D.strokeRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DStrokeRectFunc(
			this.Ref(),
		),
	)
}

// BeginPath calls the method "OffscreenCanvasRenderingContext2D.beginPath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) BeginPath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginPathFunc returns the method "OffscreenCanvasRenderingContext2D.beginPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) BeginPathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DBeginPathFunc(
			this.Ref(),
		),
	)
}

// Fill calls the method "OffscreenCanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFill(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillFunc returns the method "OffscreenCanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFillFunc(
			this.Ref(),
		),
	)
}

// Fill1 calls the method "OffscreenCanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFill1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill1Func returns the method "OffscreenCanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFill1Func(
			this.Ref(),
		),
	)
}

// Fill2 calls the method "OffscreenCanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFill2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill2Func returns the method "OffscreenCanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFill2Func(
			this.Ref(),
		),
	)
}

// Fill3 calls the method "OffscreenCanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFill3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill3Func returns the method "OffscreenCanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Fill3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFill3Func(
			this.Ref(),
		),
	)
}

// Stroke calls the method "OffscreenCanvasRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Stroke() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DStroke(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeFunc returns the method "OffscreenCanvasRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DStrokeFunc(
			this.Ref(),
		),
	)
}

// Stroke1 calls the method "OffscreenCanvasRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Stroke1(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Stroke1Func returns the method "OffscreenCanvasRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Stroke1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DStroke1Func(
			this.Ref(),
		),
	)
}

// Clip calls the method "OffscreenCanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClip(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClipFunc returns the method "OffscreenCanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ClipFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClipFunc(
			this.Ref(),
		),
	)
}

// Clip1 calls the method "OffscreenCanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClip1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip1Func returns the method "OffscreenCanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClip1Func(
			this.Ref(),
		),
	)
}

// Clip2 calls the method "OffscreenCanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClip2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip2Func returns the method "OffscreenCanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClip2Func(
			this.Ref(),
		),
	)
}

// Clip3 calls the method "OffscreenCanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClip3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip3Func returns the method "OffscreenCanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Clip3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClip3Func(
			this.Ref(),
		),
	)
}

// IsPointInPath calls the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPathFunc returns the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPathFunc() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInPathFunc(
			this.Ref(),
		),
	)
}

// IsPointInPath1 calls the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath1(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath1Func returns the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath1Func() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInPath1Func(
			this.Ref(),
		),
	)
}

// IsPointInPath2 calls the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPath2Func returns the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath2Func() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInPath2Func(
			this.Ref(),
		),
	)
}

// IsPointInPath3 calls the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath3Func returns the method "OffscreenCanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInPath3Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInPath3Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "OffscreenCanvasRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInStroke(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStrokeFunc returns the method "OffscreenCanvasRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInStrokeFunc() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "OffscreenCanvasRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStroke1Func returns the method "OffscreenCanvasRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsPointInStroke1Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// FillText calls the method "OffscreenCanvasRenderingContext2D.fillText".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillText(text js.String, x float64, y float64, maxWidth float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFillText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillTextFunc returns the method "OffscreenCanvasRenderingContext2D.fillText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFillTextFunc(
			this.Ref(),
		),
	)
}

// FillText1 calls the method "OffscreenCanvasRenderingContext2D.fillText".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillText1(text js.String, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DFillText1(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillText1Func returns the method "OffscreenCanvasRenderingContext2D.fillText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) FillText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DFillText1Func(
			this.Ref(),
		),
	)
}

// StrokeText calls the method "OffscreenCanvasRenderingContext2D.strokeText".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeText(text js.String, x float64, y float64, maxWidth float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DStrokeText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeTextFunc returns the method "OffscreenCanvasRenderingContext2D.strokeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DStrokeTextFunc(
			this.Ref(),
		),
	)
}

// StrokeText1 calls the method "OffscreenCanvasRenderingContext2D.strokeText".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeText1(text js.String, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DStrokeText1(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeText1Func returns the method "OffscreenCanvasRenderingContext2D.strokeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) StrokeText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DStrokeText1Func(
			this.Ref(),
		),
	)
}

// MeasureText calls the method "OffscreenCanvasRenderingContext2D.measureText".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) MeasureText(text js.String) (TextMetrics, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DMeasureText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
	)

	return TextMetrics{}.FromRef(_ret), _ok
}

// MeasureTextFunc returns the method "OffscreenCanvasRenderingContext2D.measureText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) MeasureTextFunc() (fn js.Func[func(text js.String) TextMetrics]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DMeasureTextFunc(
			this.Ref(),
		),
	)
}

// DrawImage calls the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImageFunc returns the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImageFunc() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DDrawImageFunc(
			this.Ref(),
		),
	)
}

// DrawImage1 calls the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImage1Func returns the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImage1Func() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DDrawImage1Func(
			this.Ref(),
		),
	)
}

// DrawImage2 calls the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(sx),
		float64(sy),
		float64(sw),
		float64(sh),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImage2Func returns the method "OffscreenCanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) DrawImage2Func() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DDrawImage2Func(
			this.Ref(),
		),
	)
}

// CreateImageData calls the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageData(sw int32, sh int32, settings ImageDataSettings) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateImageData(
		this.Ref(), js.Pointer(&_ok),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageDataFunc returns the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageDataFunc() (fn js.Func[func(sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateImageDataFunc(
			this.Ref(),
		),
	)
}

// CreateImageData1 calls the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageData1(sw int32, sh int32) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateImageData1(
		this.Ref(), js.Pointer(&_ok),
		int32(sw),
		int32(sh),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageData1Func returns the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageData1Func() (fn js.Func[func(sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateImageData1Func(
			this.Ref(),
		),
	)
}

// CreateImageData2 calls the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageData2(imagedata ImageData) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DCreateImageData2(
		this.Ref(), js.Pointer(&_ok),
		imagedata.Ref(),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageData2Func returns the method "OffscreenCanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) CreateImageData2Func() (fn js.Func[func(imagedata ImageData) ImageData]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DCreateImageData2Func(
			this.Ref(),
		),
	)
}

// GetImageData calls the method "OffscreenCanvasRenderingContext2D.getImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DGetImageData(
		this.Ref(), js.Pointer(&_ok),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// GetImageDataFunc returns the method "OffscreenCanvasRenderingContext2D.getImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetImageDataFunc() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DGetImageDataFunc(
			this.Ref(),
		),
	)
}

// GetImageData1 calls the method "OffscreenCanvasRenderingContext2D.getImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetImageData1(sx int32, sy int32, sw int32, sh int32) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DGetImageData1(
		this.Ref(), js.Pointer(&_ok),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// GetImageData1Func returns the method "OffscreenCanvasRenderingContext2D.getImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetImageData1Func() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DGetImageData1Func(
			this.Ref(),
		),
	)
}

// PutImageData calls the method "OffscreenCanvasRenderingContext2D.putImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) PutImageData(imagedata ImageData, dx int32, dy int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DPutImageData(
		this.Ref(), js.Pointer(&_ok),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PutImageDataFunc returns the method "OffscreenCanvasRenderingContext2D.putImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) PutImageDataFunc() (fn js.Func[func(imagedata ImageData, dx int32, dy int32)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DPutImageDataFunc(
			this.Ref(),
		),
	)
}

// PutImageData1 calls the method "OffscreenCanvasRenderingContext2D.putImageData".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) PutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DPutImageData1(
		this.Ref(), js.Pointer(&_ok),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
		int32(dirtyX),
		int32(dirtyY),
		int32(dirtyWidth),
		int32(dirtyHeight),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PutImageData1Func returns the method "OffscreenCanvasRenderingContext2D.putImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) PutImageData1Func() (fn js.Func[func(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DPutImageData1Func(
			this.Ref(),
		),
	)
}

// SetLineDash calls the method "OffscreenCanvasRenderingContext2D.setLineDash".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetLineDash(segments js.Array[float64]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&_ok),
		segments.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetLineDashFunc returns the method "OffscreenCanvasRenderingContext2D.setLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SetLineDashFunc() (fn js.Func[func(segments js.Array[float64])]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DSetLineDashFunc(
			this.Ref(),
		),
	)
}

// GetLineDash calls the method "OffscreenCanvasRenderingContext2D.getLineDash".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetLineDash() (js.Array[float64], bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[float64]{}.FromRef(_ret), _ok
}

// GetLineDashFunc returns the method "OffscreenCanvasRenderingContext2D.getLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) GetLineDashFunc() (fn js.Func[func() js.Array[float64]]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DGetLineDashFunc(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "OffscreenCanvasRenderingContext2D.closePath".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ClosePath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClosePathFunc returns the method "OffscreenCanvasRenderingContext2D.closePath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DClosePathFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "OffscreenCanvasRenderingContext2D.moveTo".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) MoveTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveToFunc returns the method "OffscreenCanvasRenderingContext2D.moveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DMoveToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "OffscreenCanvasRenderingContext2D.lineTo".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) LineTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineToFunc returns the method "OffscreenCanvasRenderingContext2D.lineTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DLineToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "OffscreenCanvasRenderingContext2D.quadraticCurveTo".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QuadraticCurveToFunc returns the method "OffscreenCanvasRenderingContext2D.quadraticCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "OffscreenCanvasRenderingContext2D.bezierCurveTo".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BezierCurveToFunc returns the method "OffscreenCanvasRenderingContext2D.bezierCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "OffscreenCanvasRenderingContext2D.arcTo".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcToFunc returns the method "OffscreenCanvasRenderingContext2D.arcTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DArcToFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "OffscreenCanvasRenderingContext2D.rect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RectFunc returns the method "OffscreenCanvasRenderingContext2D.rect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "OffscreenCanvasRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRectFunc returns the method "OffscreenCanvasRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "OffscreenCanvasRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRect1Func returns the method "OffscreenCanvasRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// Arc calls the method "OffscreenCanvasRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DArc(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcFunc returns the method "OffscreenCanvasRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "OffscreenCanvasRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DArc1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Arc1Func returns the method "OffscreenCanvasRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DArc1Func(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "OffscreenCanvasRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EllipseFunc returns the method "OffscreenCanvasRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "OffscreenCanvasRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Ellipse1Func returns the method "OffscreenCanvasRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DEllipse1Func(
			this.Ref(),
		),
	)
}

// Save calls the method "OffscreenCanvasRenderingContext2D.save".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Save() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DSave(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SaveFunc returns the method "OffscreenCanvasRenderingContext2D.save".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) SaveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DSaveFunc(
			this.Ref(),
		),
	)
}

// Restore calls the method "OffscreenCanvasRenderingContext2D.restore".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Restore() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DRestore(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RestoreFunc returns the method "OffscreenCanvasRenderingContext2D.restore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) RestoreFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DRestoreFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "OffscreenCanvasRenderingContext2D.reset".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "OffscreenCanvasRenderingContext2D.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DResetFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "OffscreenCanvasRenderingContext2D.isContextLost".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsContextLost() (bool, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsContextLostFunc returns the method "OffscreenCanvasRenderingContext2D.isContextLost".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvasRenderingContext2D) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.OffscreenCanvasRenderingContext2DIsContextLostFunc(
			this.Ref(),
		),
	)
}

type OneOf_HTMLCanvasElement_OffscreenCanvas struct {
	ref js.Ref
}

func (x OneOf_HTMLCanvasElement_OffscreenCanvas) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLCanvasElement_OffscreenCanvas) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLCanvasElement_OffscreenCanvas) FromRef(ref js.Ref) OneOf_HTMLCanvasElement_OffscreenCanvas {
	return OneOf_HTMLCanvasElement_OffscreenCanvas{
		ref: ref,
	}
}

func (x OneOf_HTMLCanvasElement_OffscreenCanvas) HTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLCanvasElement_OffscreenCanvas) OffscreenCanvas() OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(x.ref)
}

type ImageBitmapRenderingContext struct {
	ref js.Ref
}

func (this ImageBitmapRenderingContext) Once() ImageBitmapRenderingContext {
	this.Ref().Once()
	return this
}

func (this ImageBitmapRenderingContext) Ref() js.Ref {
	return this.ref
}

func (this ImageBitmapRenderingContext) FromRef(ref js.Ref) ImageBitmapRenderingContext {
	this.ref = ref
	return this
}

func (this ImageBitmapRenderingContext) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "ImageBitmapRenderingContext.canvas".
//
// The returned bool will be false if there is no such property.
func (this ImageBitmapRenderingContext) Canvas() (OneOf_HTMLCanvasElement_OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.GetImageBitmapRenderingContextCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_HTMLCanvasElement_OffscreenCanvas{}.FromRef(_ret), _ok
}

// TransferFromImageBitmap calls the method "ImageBitmapRenderingContext.transferFromImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this ImageBitmapRenderingContext) TransferFromImageBitmap(bitmap ImageBitmap) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallImageBitmapRenderingContextTransferFromImageBitmap(
		this.Ref(), js.Pointer(&_ok),
		bitmap.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TransferFromImageBitmapFunc returns the method "ImageBitmapRenderingContext.transferFromImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ImageBitmapRenderingContext) TransferFromImageBitmapFunc() (fn js.Func[func(bitmap ImageBitmap)]) {
	return fn.FromRef(
		bindings.ImageBitmapRenderingContextTransferFromImageBitmapFunc(
			this.Ref(),
		),
	)
}

const (
	WebGLRenderingContext_DEPTH_BUFFER_BIT                             GLenum = 0x00000100
	WebGLRenderingContext_STENCIL_BUFFER_BIT                           GLenum = 0x00000400
	WebGLRenderingContext_COLOR_BUFFER_BIT                             GLenum = 0x00004000
	WebGLRenderingContext_POINTS                                       GLenum = 0x0000
	WebGLRenderingContext_LINES                                        GLenum = 0x0001
	WebGLRenderingContext_LINE_LOOP                                    GLenum = 0x0002
	WebGLRenderingContext_LINE_STRIP                                   GLenum = 0x0003
	WebGLRenderingContext_TRIANGLES                                    GLenum = 0x0004
	WebGLRenderingContext_TRIANGLE_STRIP                               GLenum = 0x0005
	WebGLRenderingContext_TRIANGLE_FAN                                 GLenum = 0x0006
	WebGLRenderingContext_ZERO                                         GLenum = 0
	WebGLRenderingContext_ONE                                          GLenum = 1
	WebGLRenderingContext_SRC_COLOR                                    GLenum = 0x0300
	WebGLRenderingContext_ONE_MINUS_SRC_COLOR                          GLenum = 0x0301
	WebGLRenderingContext_SRC_ALPHA                                    GLenum = 0x0302
	WebGLRenderingContext_ONE_MINUS_SRC_ALPHA                          GLenum = 0x0303
	WebGLRenderingContext_DST_ALPHA                                    GLenum = 0x0304
	WebGLRenderingContext_ONE_MINUS_DST_ALPHA                          GLenum = 0x0305
	WebGLRenderingContext_DST_COLOR                                    GLenum = 0x0306
	WebGLRenderingContext_ONE_MINUS_DST_COLOR                          GLenum = 0x0307
	WebGLRenderingContext_SRC_ALPHA_SATURATE                           GLenum = 0x0308
	WebGLRenderingContext_FUNC_ADD                                     GLenum = 0x8006
	WebGLRenderingContext_BLEND_EQUATION                               GLenum = 0x8009
	WebGLRenderingContext_BLEND_EQUATION_RGB                           GLenum = 0x8009
	WebGLRenderingContext_BLEND_EQUATION_ALPHA                         GLenum = 0x883D
	WebGLRenderingContext_FUNC_SUBTRACT                                GLenum = 0x800A
	WebGLRenderingContext_FUNC_REVERSE_SUBTRACT                        GLenum = 0x800B
	WebGLRenderingContext_BLEND_DST_RGB                                GLenum = 0x80C8
	WebGLRenderingContext_BLEND_SRC_RGB                                GLenum = 0x80C9
	WebGLRenderingContext_BLEND_DST_ALPHA                              GLenum = 0x80CA
	WebGLRenderingContext_BLEND_SRC_ALPHA                              GLenum = 0x80CB
	WebGLRenderingContext_CONSTANT_COLOR                               GLenum = 0x8001
	WebGLRenderingContext_ONE_MINUS_CONSTANT_COLOR                     GLenum = 0x8002
	WebGLRenderingContext_CONSTANT_ALPHA                               GLenum = 0x8003
	WebGLRenderingContext_ONE_MINUS_CONSTANT_ALPHA                     GLenum = 0x8004
	WebGLRenderingContext_BLEND_COLOR                                  GLenum = 0x8005
	WebGLRenderingContext_ARRAY_BUFFER                                 GLenum = 0x8892
	WebGLRenderingContext_ELEMENT_ARRAY_BUFFER                         GLenum = 0x8893
	WebGLRenderingContext_ARRAY_BUFFER_BINDING                         GLenum = 0x8894
	WebGLRenderingContext_ELEMENT_ARRAY_BUFFER_BINDING                 GLenum = 0x8895
	WebGLRenderingContext_STREAM_DRAW                                  GLenum = 0x88E0
	WebGLRenderingContext_STATIC_DRAW                                  GLenum = 0x88E4
	WebGLRenderingContext_DYNAMIC_DRAW                                 GLenum = 0x88E8
	WebGLRenderingContext_BUFFER_SIZE                                  GLenum = 0x8764
	WebGLRenderingContext_BUFFER_USAGE                                 GLenum = 0x8765
	WebGLRenderingContext_CURRENT_VERTEX_ATTRIB                        GLenum = 0x8626
	WebGLRenderingContext_FRONT                                        GLenum = 0x0404
	WebGLRenderingContext_BACK                                         GLenum = 0x0405
	WebGLRenderingContext_FRONT_AND_BACK                               GLenum = 0x0408
	WebGLRenderingContext_CULL_FACE                                    GLenum = 0x0B44
	WebGLRenderingContext_BLEND                                        GLenum = 0x0BE2
	WebGLRenderingContext_DITHER                                       GLenum = 0x0BD0
	WebGLRenderingContext_STENCIL_TEST                                 GLenum = 0x0B90
	WebGLRenderingContext_DEPTH_TEST                                   GLenum = 0x0B71
	WebGLRenderingContext_SCISSOR_TEST                                 GLenum = 0x0C11
	WebGLRenderingContext_POLYGON_OFFSET_FILL                          GLenum = 0x8037
	WebGLRenderingContext_SAMPLE_ALPHA_TO_COVERAGE                     GLenum = 0x809E
	WebGLRenderingContext_SAMPLE_COVERAGE                              GLenum = 0x80A0
	WebGLRenderingContext_NO_ERROR                                     GLenum = 0
	WebGLRenderingContext_INVALID_ENUM                                 GLenum = 0x0500
	WebGLRenderingContext_INVALID_VALUE                                GLenum = 0x0501
	WebGLRenderingContext_INVALID_OPERATION                            GLenum = 0x0502
	WebGLRenderingContext_OUT_OF_MEMORY                                GLenum = 0x0505
	WebGLRenderingContext_CW                                           GLenum = 0x0900
	WebGLRenderingContext_CCW                                          GLenum = 0x0901
	WebGLRenderingContext_LINE_WIDTH                                   GLenum = 0x0B21
	WebGLRenderingContext_ALIASED_POINT_SIZE_RANGE                     GLenum = 0x846D
	WebGLRenderingContext_ALIASED_LINE_WIDTH_RANGE                     GLenum = 0x846E
	WebGLRenderingContext_CULL_FACE_MODE                               GLenum = 0x0B45
	WebGLRenderingContext_FRONT_FACE                                   GLenum = 0x0B46
	WebGLRenderingContext_DEPTH_RANGE                                  GLenum = 0x0B70
	WebGLRenderingContext_DEPTH_WRITEMASK                              GLenum = 0x0B72
	WebGLRenderingContext_DEPTH_CLEAR_VALUE                            GLenum = 0x0B73
	WebGLRenderingContext_DEPTH_FUNC                                   GLenum = 0x0B74
	WebGLRenderingContext_STENCIL_CLEAR_VALUE                          GLenum = 0x0B91
	WebGLRenderingContext_STENCIL_FUNC                                 GLenum = 0x0B92
	WebGLRenderingContext_STENCIL_FAIL                                 GLenum = 0x0B94
	WebGLRenderingContext_STENCIL_PASS_DEPTH_FAIL                      GLenum = 0x0B95
	WebGLRenderingContext_STENCIL_PASS_DEPTH_PASS                      GLenum = 0x0B96
	WebGLRenderingContext_STENCIL_REF                                  GLenum = 0x0B97
	WebGLRenderingContext_STENCIL_VALUE_MASK                           GLenum = 0x0B93
	WebGLRenderingContext_STENCIL_WRITEMASK                            GLenum = 0x0B98
	WebGLRenderingContext_STENCIL_BACK_FUNC                            GLenum = 0x8800
	WebGLRenderingContext_STENCIL_BACK_FAIL                            GLenum = 0x8801
	WebGLRenderingContext_STENCIL_BACK_PASS_DEPTH_FAIL                 GLenum = 0x8802
	WebGLRenderingContext_STENCIL_BACK_PASS_DEPTH_PASS                 GLenum = 0x8803
	WebGLRenderingContext_STENCIL_BACK_REF                             GLenum = 0x8CA3
	WebGLRenderingContext_STENCIL_BACK_VALUE_MASK                      GLenum = 0x8CA4
	WebGLRenderingContext_STENCIL_BACK_WRITEMASK                       GLenum = 0x8CA5
	WebGLRenderingContext_VIEWPORT                                     GLenum = 0x0BA2
	WebGLRenderingContext_SCISSOR_BOX                                  GLenum = 0x0C10
	WebGLRenderingContext_COLOR_CLEAR_VALUE                            GLenum = 0x0C22
	WebGLRenderingContext_COLOR_WRITEMASK                              GLenum = 0x0C23
	WebGLRenderingContext_UNPACK_ALIGNMENT                             GLenum = 0x0CF5
	WebGLRenderingContext_PACK_ALIGNMENT                               GLenum = 0x0D05
	WebGLRenderingContext_MAX_TEXTURE_SIZE                             GLenum = 0x0D33
	WebGLRenderingContext_MAX_VIEWPORT_DIMS                            GLenum = 0x0D3A
	WebGLRenderingContext_SUBPIXEL_BITS                                GLenum = 0x0D50
	WebGLRenderingContext_RED_BITS                                     GLenum = 0x0D52
	WebGLRenderingContext_GREEN_BITS                                   GLenum = 0x0D53
	WebGLRenderingContext_BLUE_BITS                                    GLenum = 0x0D54
	WebGLRenderingContext_ALPHA_BITS                                   GLenum = 0x0D55
	WebGLRenderingContext_DEPTH_BITS                                   GLenum = 0x0D56
	WebGLRenderingContext_STENCIL_BITS                                 GLenum = 0x0D57
	WebGLRenderingContext_POLYGON_OFFSET_UNITS                         GLenum = 0x2A00
	WebGLRenderingContext_POLYGON_OFFSET_FACTOR                        GLenum = 0x8038
	WebGLRenderingContext_TEXTURE_BINDING_2D                           GLenum = 0x8069
	WebGLRenderingContext_SAMPLE_BUFFERS                               GLenum = 0x80A8
	WebGLRenderingContext_SAMPLES                                      GLenum = 0x80A9
	WebGLRenderingContext_SAMPLE_COVERAGE_VALUE                        GLenum = 0x80AA
	WebGLRenderingContext_SAMPLE_COVERAGE_INVERT                       GLenum = 0x80AB
	WebGLRenderingContext_COMPRESSED_TEXTURE_FORMATS                   GLenum = 0x86A3
	WebGLRenderingContext_DONT_CARE                                    GLenum = 0x1100
	WebGLRenderingContext_FASTEST                                      GLenum = 0x1101
	WebGLRenderingContext_NICEST                                       GLenum = 0x1102
	WebGLRenderingContext_GENERATE_MIPMAP_HINT                         GLenum = 0x8192
	WebGLRenderingContext_BYTE                                         GLenum = 0x1400
	WebGLRenderingContext_UNSIGNED_BYTE                                GLenum = 0x1401
	WebGLRenderingContext_SHORT                                        GLenum = 0x1402
	WebGLRenderingContext_UNSIGNED_SHORT                               GLenum = 0x1403
	WebGLRenderingContext_INT                                          GLenum = 0x1404
	WebGLRenderingContext_UNSIGNED_INT                                 GLenum = 0x1405
	WebGLRenderingContext_FLOAT                                        GLenum = 0x1406
	WebGLRenderingContext_DEPTH_COMPONENT                              GLenum = 0x1902
	WebGLRenderingContext_ALPHA                                        GLenum = 0x1906
	WebGLRenderingContext_RGB                                          GLenum = 0x1907
	WebGLRenderingContext_RGBA                                         GLenum = 0x1908
	WebGLRenderingContext_LUMINANCE                                    GLenum = 0x1909
	WebGLRenderingContext_LUMINANCE_ALPHA                              GLenum = 0x190A
	WebGLRenderingContext_UNSIGNED_SHORT_4_4_4_4                       GLenum = 0x8033
	WebGLRenderingContext_UNSIGNED_SHORT_5_5_5_1                       GLenum = 0x8034
	WebGLRenderingContext_UNSIGNED_SHORT_5_6_5                         GLenum = 0x8363
	WebGLRenderingContext_FRAGMENT_SHADER                              GLenum = 0x8B30
	WebGLRenderingContext_VERTEX_SHADER                                GLenum = 0x8B31
	WebGLRenderingContext_MAX_VERTEX_ATTRIBS                           GLenum = 0x8869
	WebGLRenderingContext_MAX_VERTEX_UNIFORM_VECTORS                   GLenum = 0x8DFB
	WebGLRenderingContext_MAX_VARYING_VECTORS                          GLenum = 0x8DFC
	WebGLRenderingContext_MAX_COMBINED_TEXTURE_IMAGE_UNITS             GLenum = 0x8B4D
	WebGLRenderingContext_MAX_VERTEX_TEXTURE_IMAGE_UNITS               GLenum = 0x8B4C
	WebGLRenderingContext_MAX_TEXTURE_IMAGE_UNITS                      GLenum = 0x8872
	WebGLRenderingContext_MAX_FRAGMENT_UNIFORM_VECTORS                 GLenum = 0x8DFD
	WebGLRenderingContext_SHADER_TYPE                                  GLenum = 0x8B4F
	WebGLRenderingContext_DELETE_STATUS                                GLenum = 0x8B80
	WebGLRenderingContext_LINK_STATUS                                  GLenum = 0x8B82
	WebGLRenderingContext_VALIDATE_STATUS                              GLenum = 0x8B83
	WebGLRenderingContext_ATTACHED_SHADERS                             GLenum = 0x8B85
	WebGLRenderingContext_ACTIVE_UNIFORMS                              GLenum = 0x8B86
	WebGLRenderingContext_ACTIVE_ATTRIBUTES                            GLenum = 0x8B89
	WebGLRenderingContext_SHADING_LANGUAGE_VERSION                     GLenum = 0x8B8C
	WebGLRenderingContext_CURRENT_PROGRAM                              GLenum = 0x8B8D
	WebGLRenderingContext_NEVER                                        GLenum = 0x0200
	WebGLRenderingContext_LESS                                         GLenum = 0x0201
	WebGLRenderingContext_EQUAL                                        GLenum = 0x0202
	WebGLRenderingContext_LEQUAL                                       GLenum = 0x0203
	WebGLRenderingContext_GREATER                                      GLenum = 0x0204
	WebGLRenderingContext_NOTEQUAL                                     GLenum = 0x0205
	WebGLRenderingContext_GEQUAL                                       GLenum = 0x0206
	WebGLRenderingContext_ALWAYS                                       GLenum = 0x0207
	WebGLRenderingContext_KEEP                                         GLenum = 0x1E00
	WebGLRenderingContext_REPLACE                                      GLenum = 0x1E01
	WebGLRenderingContext_INCR                                         GLenum = 0x1E02
	WebGLRenderingContext_DECR                                         GLenum = 0x1E03
	WebGLRenderingContext_INVERT                                       GLenum = 0x150A
	WebGLRenderingContext_INCR_WRAP                                    GLenum = 0x8507
	WebGLRenderingContext_DECR_WRAP                                    GLenum = 0x8508
	WebGLRenderingContext_VENDOR                                       GLenum = 0x1F00
	WebGLRenderingContext_RENDERER                                     GLenum = 0x1F01
	WebGLRenderingContext_VERSION                                      GLenum = 0x1F02
	WebGLRenderingContext_NEAREST                                      GLenum = 0x2600
	WebGLRenderingContext_LINEAR                                       GLenum = 0x2601
	WebGLRenderingContext_NEAREST_MIPMAP_NEAREST                       GLenum = 0x2700
	WebGLRenderingContext_LINEAR_MIPMAP_NEAREST                        GLenum = 0x2701
	WebGLRenderingContext_NEAREST_MIPMAP_LINEAR                        GLenum = 0x2702
	WebGLRenderingContext_LINEAR_MIPMAP_LINEAR                         GLenum = 0x2703
	WebGLRenderingContext_TEXTURE_MAG_FILTER                           GLenum = 0x2800
	WebGLRenderingContext_TEXTURE_MIN_FILTER                           GLenum = 0x2801
	WebGLRenderingContext_TEXTURE_WRAP_S                               GLenum = 0x2802
	WebGLRenderingContext_TEXTURE_WRAP_T                               GLenum = 0x2803
	WebGLRenderingContext_TEXTURE_2D                                   GLenum = 0x0DE1
	WebGLRenderingContext_TEXTURE                                      GLenum = 0x1702
	WebGLRenderingContext_TEXTURE_CUBE_MAP                             GLenum = 0x8513
	WebGLRenderingContext_TEXTURE_BINDING_CUBE_MAP                     GLenum = 0x8514
	WebGLRenderingContext_TEXTURE_CUBE_MAP_POSITIVE_X                  GLenum = 0x8515
	WebGLRenderingContext_TEXTURE_CUBE_MAP_NEGATIVE_X                  GLenum = 0x8516
	WebGLRenderingContext_TEXTURE_CUBE_MAP_POSITIVE_Y                  GLenum = 0x8517
	WebGLRenderingContext_TEXTURE_CUBE_MAP_NEGATIVE_Y                  GLenum = 0x8518
	WebGLRenderingContext_TEXTURE_CUBE_MAP_POSITIVE_Z                  GLenum = 0x8519
	WebGLRenderingContext_TEXTURE_CUBE_MAP_NEGATIVE_Z                  GLenum = 0x851A
	WebGLRenderingContext_MAX_CUBE_MAP_TEXTURE_SIZE                    GLenum = 0x851C
	WebGLRenderingContext_TEXTURE0                                     GLenum = 0x84C0
	WebGLRenderingContext_TEXTURE1                                     GLenum = 0x84C1
	WebGLRenderingContext_TEXTURE2                                     GLenum = 0x84C2
	WebGLRenderingContext_TEXTURE3                                     GLenum = 0x84C3
	WebGLRenderingContext_TEXTURE4                                     GLenum = 0x84C4
	WebGLRenderingContext_TEXTURE5                                     GLenum = 0x84C5
	WebGLRenderingContext_TEXTURE6                                     GLenum = 0x84C6
	WebGLRenderingContext_TEXTURE7                                     GLenum = 0x84C7
	WebGLRenderingContext_TEXTURE8                                     GLenum = 0x84C8
	WebGLRenderingContext_TEXTURE9                                     GLenum = 0x84C9
	WebGLRenderingContext_TEXTURE10                                    GLenum = 0x84CA
	WebGLRenderingContext_TEXTURE11                                    GLenum = 0x84CB
	WebGLRenderingContext_TEXTURE12                                    GLenum = 0x84CC
	WebGLRenderingContext_TEXTURE13                                    GLenum = 0x84CD
	WebGLRenderingContext_TEXTURE14                                    GLenum = 0x84CE
	WebGLRenderingContext_TEXTURE15                                    GLenum = 0x84CF
	WebGLRenderingContext_TEXTURE16                                    GLenum = 0x84D0
	WebGLRenderingContext_TEXTURE17                                    GLenum = 0x84D1
	WebGLRenderingContext_TEXTURE18                                    GLenum = 0x84D2
	WebGLRenderingContext_TEXTURE19                                    GLenum = 0x84D3
	WebGLRenderingContext_TEXTURE20                                    GLenum = 0x84D4
	WebGLRenderingContext_TEXTURE21                                    GLenum = 0x84D5
	WebGLRenderingContext_TEXTURE22                                    GLenum = 0x84D6
	WebGLRenderingContext_TEXTURE23                                    GLenum = 0x84D7
	WebGLRenderingContext_TEXTURE24                                    GLenum = 0x84D8
	WebGLRenderingContext_TEXTURE25                                    GLenum = 0x84D9
	WebGLRenderingContext_TEXTURE26                                    GLenum = 0x84DA
	WebGLRenderingContext_TEXTURE27                                    GLenum = 0x84DB
	WebGLRenderingContext_TEXTURE28                                    GLenum = 0x84DC
	WebGLRenderingContext_TEXTURE29                                    GLenum = 0x84DD
	WebGLRenderingContext_TEXTURE30                                    GLenum = 0x84DE
	WebGLRenderingContext_TEXTURE31                                    GLenum = 0x84DF
	WebGLRenderingContext_ACTIVE_TEXTURE                               GLenum = 0x84E0
	WebGLRenderingContext_REPEAT                                       GLenum = 0x2901
	WebGLRenderingContext_CLAMP_TO_EDGE                                GLenum = 0x812F
	WebGLRenderingContext_MIRRORED_REPEAT                              GLenum = 0x8370
	WebGLRenderingContext_FLOAT_VEC2                                   GLenum = 0x8B50
	WebGLRenderingContext_FLOAT_VEC3                                   GLenum = 0x8B51
	WebGLRenderingContext_FLOAT_VEC4                                   GLenum = 0x8B52
	WebGLRenderingContext_INT_VEC2                                     GLenum = 0x8B53
	WebGLRenderingContext_INT_VEC3                                     GLenum = 0x8B54
	WebGLRenderingContext_INT_VEC4                                     GLenum = 0x8B55
	WebGLRenderingContext_BOOL                                         GLenum = 0x8B56
	WebGLRenderingContext_BOOL_VEC2                                    GLenum = 0x8B57
	WebGLRenderingContext_BOOL_VEC3                                    GLenum = 0x8B58
	WebGLRenderingContext_BOOL_VEC4                                    GLenum = 0x8B59
	WebGLRenderingContext_FLOAT_MAT2                                   GLenum = 0x8B5A
	WebGLRenderingContext_FLOAT_MAT3                                   GLenum = 0x8B5B
	WebGLRenderingContext_FLOAT_MAT4                                   GLenum = 0x8B5C
	WebGLRenderingContext_SAMPLER_2D                                   GLenum = 0x8B5E
	WebGLRenderingContext_SAMPLER_CUBE                                 GLenum = 0x8B60
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_ENABLED                  GLenum = 0x8622
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_SIZE                     GLenum = 0x8623
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_STRIDE                   GLenum = 0x8624
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_TYPE                     GLenum = 0x8625
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_NORMALIZED               GLenum = 0x886A
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_POINTER                  GLenum = 0x8645
	WebGLRenderingContext_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           GLenum = 0x889F
	WebGLRenderingContext_IMPLEMENTATION_COLOR_READ_TYPE               GLenum = 0x8B9A
	WebGLRenderingContext_IMPLEMENTATION_COLOR_READ_FORMAT             GLenum = 0x8B9B
	WebGLRenderingContext_COMPILE_STATUS                               GLenum = 0x8B81
	WebGLRenderingContext_LOW_FLOAT                                    GLenum = 0x8DF0
	WebGLRenderingContext_MEDIUM_FLOAT                                 GLenum = 0x8DF1
	WebGLRenderingContext_HIGH_FLOAT                                   GLenum = 0x8DF2
	WebGLRenderingContext_LOW_INT                                      GLenum = 0x8DF3
	WebGLRenderingContext_MEDIUM_INT                                   GLenum = 0x8DF4
	WebGLRenderingContext_HIGH_INT                                     GLenum = 0x8DF5
	WebGLRenderingContext_FRAMEBUFFER                                  GLenum = 0x8D40
	WebGLRenderingContext_RENDERBUFFER                                 GLenum = 0x8D41
	WebGLRenderingContext_RGBA4                                        GLenum = 0x8056
	WebGLRenderingContext_RGB5_A1                                      GLenum = 0x8057
	WebGLRenderingContext_RGB565                                       GLenum = 0x8D62
	WebGLRenderingContext_DEPTH_COMPONENT16                            GLenum = 0x81A5
	WebGLRenderingContext_STENCIL_INDEX8                               GLenum = 0x8D48
	WebGLRenderingContext_DEPTH_STENCIL                                GLenum = 0x84F9
	WebGLRenderingContext_RENDERBUFFER_WIDTH                           GLenum = 0x8D42
	WebGLRenderingContext_RENDERBUFFER_HEIGHT                          GLenum = 0x8D43
	WebGLRenderingContext_RENDERBUFFER_INTERNAL_FORMAT                 GLenum = 0x8D44
	WebGLRenderingContext_RENDERBUFFER_RED_SIZE                        GLenum = 0x8D50
	WebGLRenderingContext_RENDERBUFFER_GREEN_SIZE                      GLenum = 0x8D51
	WebGLRenderingContext_RENDERBUFFER_BLUE_SIZE                       GLenum = 0x8D52
	WebGLRenderingContext_RENDERBUFFER_ALPHA_SIZE                      GLenum = 0x8D53
	WebGLRenderingContext_RENDERBUFFER_DEPTH_SIZE                      GLenum = 0x8D54
	WebGLRenderingContext_RENDERBUFFER_STENCIL_SIZE                    GLenum = 0x8D55
	WebGLRenderingContext_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           GLenum = 0x8CD0
	WebGLRenderingContext_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           GLenum = 0x8CD1
	WebGLRenderingContext_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         GLenum = 0x8CD2
	WebGLRenderingContext_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE GLenum = 0x8CD3
	WebGLRenderingContext_COLOR_ATTACHMENT0                            GLenum = 0x8CE0
	WebGLRenderingContext_DEPTH_ATTACHMENT                             GLenum = 0x8D00
	WebGLRenderingContext_STENCIL_ATTACHMENT                           GLenum = 0x8D20
	WebGLRenderingContext_DEPTH_STENCIL_ATTACHMENT                     GLenum = 0x821A
	WebGLRenderingContext_NONE                                         GLenum = 0
	WebGLRenderingContext_FRAMEBUFFER_COMPLETE                         GLenum = 0x8CD5
	WebGLRenderingContext_FRAMEBUFFER_INCOMPLETE_ATTACHMENT            GLenum = 0x8CD6
	WebGLRenderingContext_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    GLenum = 0x8CD7
	WebGLRenderingContext_FRAMEBUFFER_INCOMPLETE_DIMENSIONS            GLenum = 0x8CD9
	WebGLRenderingContext_FRAMEBUFFER_UNSUPPORTED                      GLenum = 0x8CDD
	WebGLRenderingContext_FRAMEBUFFER_BINDING                          GLenum = 0x8CA6
	WebGLRenderingContext_RENDERBUFFER_BINDING                         GLenum = 0x8CA7
	WebGLRenderingContext_MAX_RENDERBUFFER_SIZE                        GLenum = 0x84E8
	WebGLRenderingContext_INVALID_FRAMEBUFFER_OPERATION                GLenum = 0x0506
	WebGLRenderingContext_UNPACK_FLIP_Y_WEBGL                          GLenum = 0x9240
	WebGLRenderingContext_UNPACK_PREMULTIPLY_ALPHA_WEBGL               GLenum = 0x9241
	WebGLRenderingContext_CONTEXT_LOST_WEBGL                           GLenum = 0x9242
	WebGLRenderingContext_UNPACK_COLORSPACE_CONVERSION_WEBGL           GLenum = 0x9243
	WebGLRenderingContext_BROWSER_DEFAULT_WEBGL                        GLenum = 0x9244
)

type WebGLPowerPreference uint32

const (
	_ WebGLPowerPreference = iota

	WebGLPowerPreference_DEFAULT
	WebGLPowerPreference_LOW_POWER
	WebGLPowerPreference_HIGH_PERFORMANCE
)

func (WebGLPowerPreference) FromRef(str js.Ref) WebGLPowerPreference {
	return WebGLPowerPreference(bindings.ConstOfWebGLPowerPreference(str))
}

func (x WebGLPowerPreference) String() (string, bool) {
	switch x {
	case WebGLPowerPreference_DEFAULT:
		return "default", true
	case WebGLPowerPreference_LOW_POWER:
		return "low-power", true
	case WebGLPowerPreference_HIGH_PERFORMANCE:
		return "high-performance", true
	default:
		return "", false
	}
}

type WebGLContextAttributes struct {
	// Alpha is "WebGLContextAttributes.alpha"
	//
	// Optional, defaults to true.
	Alpha bool
	// Depth is "WebGLContextAttributes.depth"
	//
	// Optional, defaults to true.
	Depth bool
	// Stencil is "WebGLContextAttributes.stencil"
	//
	// Optional, defaults to false.
	Stencil bool
	// Antialias is "WebGLContextAttributes.antialias"
	//
	// Optional, defaults to true.
	Antialias bool
	// PremultipliedAlpha is "WebGLContextAttributes.premultipliedAlpha"
	//
	// Optional, defaults to true.
	PremultipliedAlpha bool
	// PreserveDrawingBuffer is "WebGLContextAttributes.preserveDrawingBuffer"
	//
	// Optional, defaults to false.
	PreserveDrawingBuffer bool
	// PowerPreference is "WebGLContextAttributes.powerPreference"
	//
	// Optional, defaults to "default".
	PowerPreference WebGLPowerPreference
	// FailIfMajorPerformanceCaveat is "WebGLContextAttributes.failIfMajorPerformanceCaveat"
	//
	// Optional, defaults to false.
	FailIfMajorPerformanceCaveat bool
	// Desynchronized is "WebGLContextAttributes.desynchronized"
	//
	// Optional, defaults to false.
	Desynchronized bool
	// XrCompatible is "WebGLContextAttributes.xrCompatible"
	//
	// Optional, defaults to false.
	XrCompatible bool

	FFI_USE_Alpha                        bool // for Alpha.
	FFI_USE_Depth                        bool // for Depth.
	FFI_USE_Stencil                      bool // for Stencil.
	FFI_USE_Antialias                    bool // for Antialias.
	FFI_USE_PremultipliedAlpha           bool // for PremultipliedAlpha.
	FFI_USE_PreserveDrawingBuffer        bool // for PreserveDrawingBuffer.
	FFI_USE_FailIfMajorPerformanceCaveat bool // for FailIfMajorPerformanceCaveat.
	FFI_USE_Desynchronized               bool // for Desynchronized.
	FFI_USE_XrCompatible                 bool // for XrCompatible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebGLContextAttributes with all fields set.
func (p WebGLContextAttributes) FromRef(ref js.Ref) WebGLContextAttributes {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 WebGLContextAttributes WebGLContextAttributes [// WebGLContextAttributes] [0x140002c2f00 0x140002c30e0 0x140002c3220 0x140002c3360 0x140002c34a0 0x140002c35e0 0x140002c3720 0x140002c37c0 0x140002c3900 0x140002c3a40 0x140002c2fa0 0x140002c3180 0x140002c32c0 0x140002c3400 0x140002c3540 0x140002c3680 0x140002c3860 0x140002c39a0 0x140002c3ae0] 0x1400107f938 {0 0}} in the application heap.
func (p WebGLContextAttributes) New() js.Ref {
	return bindings.WebGLContextAttributesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebGLContextAttributes) UpdateFrom(ref js.Ref) {
	bindings.WebGLContextAttributesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebGLContextAttributes) Update(ref js.Ref) {
	bindings.WebGLContextAttributesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebGLProgram struct {
	WebGLObject
}

func (this WebGLProgram) Once() WebGLProgram {
	this.Ref().Once()
	return this
}

func (this WebGLProgram) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLProgram) FromRef(ref js.Ref) WebGLProgram {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLProgram) Free() {
	this.Ref().Free()
}

type WebGLShader struct {
	WebGLObject
}

func (this WebGLShader) Once() WebGLShader {
	this.Ref().Once()
	return this
}

func (this WebGLShader) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLShader) FromRef(ref js.Ref) WebGLShader {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLShader) Free() {
	this.Ref().Free()
}

type WebGLBuffer struct {
	WebGLObject
}

func (this WebGLBuffer) Once() WebGLBuffer {
	this.Ref().Once()
	return this
}

func (this WebGLBuffer) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLBuffer) FromRef(ref js.Ref) WebGLBuffer {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLBuffer) Free() {
	this.Ref().Free()
}

type WebGLFramebuffer struct {
	WebGLObject
}

func (this WebGLFramebuffer) Once() WebGLFramebuffer {
	this.Ref().Once()
	return this
}

func (this WebGLFramebuffer) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLFramebuffer) FromRef(ref js.Ref) WebGLFramebuffer {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLFramebuffer) Free() {
	this.Ref().Free()
}

type WebGLRenderbuffer struct {
	WebGLObject
}

func (this WebGLRenderbuffer) Once() WebGLRenderbuffer {
	this.Ref().Once()
	return this
}

func (this WebGLRenderbuffer) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLRenderbuffer) FromRef(ref js.Ref) WebGLRenderbuffer {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLRenderbuffer) Free() {
	this.Ref().Free()
}

type WebGLTexture struct {
	WebGLObject
}

func (this WebGLTexture) Once() WebGLTexture {
	this.Ref().Once()
	return this
}

func (this WebGLTexture) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLTexture) FromRef(ref js.Ref) WebGLTexture {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLTexture) Free() {
	this.Ref().Free()
}

type GLclampf float32

type GLbitfield uint32

type GLboolean bool

type WebGLActiveInfo struct {
	ref js.Ref
}

func (this WebGLActiveInfo) Once() WebGLActiveInfo {
	this.Ref().Once()
	return this
}

func (this WebGLActiveInfo) Ref() js.Ref {
	return this.ref
}

func (this WebGLActiveInfo) FromRef(ref js.Ref) WebGLActiveInfo {
	this.ref = ref
	return this
}

func (this WebGLActiveInfo) Free() {
	this.Ref().Free()
}

// Size returns the value of property "WebGLActiveInfo.size".
//
// The returned bool will be false if there is no such property.
func (this WebGLActiveInfo) Size() (GLint, bool) {
	var _ok bool
	_ret := bindings.GetWebGLActiveInfoSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLint(_ret), _ok
}

// Type returns the value of property "WebGLActiveInfo.type".
//
// The returned bool will be false if there is no such property.
func (this WebGLActiveInfo) Type() (GLenum, bool) {
	var _ok bool
	_ret := bindings.GetWebGLActiveInfoType(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLenum(_ret), _ok
}

// Name returns the value of property "WebGLActiveInfo.name".
//
// The returned bool will be false if there is no such property.
func (this WebGLActiveInfo) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetWebGLActiveInfoName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type WebGLShaderPrecisionFormat struct {
	ref js.Ref
}

func (this WebGLShaderPrecisionFormat) Once() WebGLShaderPrecisionFormat {
	this.Ref().Once()
	return this
}

func (this WebGLShaderPrecisionFormat) Ref() js.Ref {
	return this.ref
}

func (this WebGLShaderPrecisionFormat) FromRef(ref js.Ref) WebGLShaderPrecisionFormat {
	this.ref = ref
	return this
}

func (this WebGLShaderPrecisionFormat) Free() {
	this.Ref().Free()
}

// RangeMin returns the value of property "WebGLShaderPrecisionFormat.rangeMin".
//
// The returned bool will be false if there is no such property.
func (this WebGLShaderPrecisionFormat) RangeMin() (GLint, bool) {
	var _ok bool
	_ret := bindings.GetWebGLShaderPrecisionFormatRangeMin(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLint(_ret), _ok
}

// RangeMax returns the value of property "WebGLShaderPrecisionFormat.rangeMax".
//
// The returned bool will be false if there is no such property.
func (this WebGLShaderPrecisionFormat) RangeMax() (GLint, bool) {
	var _ok bool
	_ret := bindings.GetWebGLShaderPrecisionFormatRangeMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLint(_ret), _ok
}

// Precision returns the value of property "WebGLShaderPrecisionFormat.precision".
//
// The returned bool will be false if there is no such property.
func (this WebGLShaderPrecisionFormat) Precision() (GLint, bool) {
	var _ok bool
	_ret := bindings.GetWebGLShaderPrecisionFormatPrecision(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLint(_ret), _ok
}

type WebGLUniformLocation struct {
	ref js.Ref
}

func (this WebGLUniformLocation) Once() WebGLUniformLocation {
	this.Ref().Once()
	return this
}

func (this WebGLUniformLocation) Ref() js.Ref {
	return this.ref
}

func (this WebGLUniformLocation) FromRef(ref js.Ref) WebGLUniformLocation {
	this.ref = ref
	return this
}

func (this WebGLUniformLocation) Free() {
	this.Ref().Free()
}

type GLfloat float32

type OneOf_TypedArrayFloat32_ArrayGLfloat struct {
	ref js.Ref
}

func (x OneOf_TypedArrayFloat32_ArrayGLfloat) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayFloat32_ArrayGLfloat) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayFloat32_ArrayGLfloat) FromRef(ref js.Ref) OneOf_TypedArrayFloat32_ArrayGLfloat {
	return OneOf_TypedArrayFloat32_ArrayGLfloat{
		ref: ref,
	}
}

func (x OneOf_TypedArrayFloat32_ArrayGLfloat) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayFloat32_ArrayGLfloat) ArrayGLfloat() js.Array[GLfloat] {
	return js.Array[GLfloat]{}.FromRef(x.ref)
}

type Float32List = OneOf_TypedArrayFloat32_ArrayGLfloat

type GLsizeiptr int64

type VideoFrameMetadata struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoFrameMetadata with all fields set.
func (p VideoFrameMetadata) FromRef(ref js.Ref) VideoFrameMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 VideoFrameMetadata VideoFrameMetadata [// VideoFrameMetadata] [] 0x140037e8450 {0 0}} in the application heap.
func (p VideoFrameMetadata) New() js.Ref {
	return bindings.VideoFrameMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoFrameMetadata) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoFrameMetadata) Update(ref js.Ref) {
	bindings.VideoFrameMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoFrameInit struct {
	// Duration is "VideoFrameInit.duration"
	//
	// Optional
	Duration uint64
	// Timestamp is "VideoFrameInit.timestamp"
	//
	// Optional
	Timestamp int64
	// Alpha is "VideoFrameInit.alpha"
	//
	// Optional, defaults to "keep".
	Alpha AlphaOption
	// VisibleRect is "VideoFrameInit.visibleRect"
	//
	// Optional
	VisibleRect DOMRectInit
	// DisplayWidth is "VideoFrameInit.displayWidth"
	//
	// Optional
	DisplayWidth uint32
	// DisplayHeight is "VideoFrameInit.displayHeight"
	//
	// Optional
	DisplayHeight uint32
	// Metadata is "VideoFrameInit.metadata"
	//
	// Optional
	Metadata VideoFrameMetadata

	FFI_USE_Duration      bool // for Duration.
	FFI_USE_Timestamp     bool // for Timestamp.
	FFI_USE_DisplayWidth  bool // for DisplayWidth.
	FFI_USE_DisplayHeight bool // for DisplayHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoFrameInit with all fields set.
func (p VideoFrameInit) FromRef(ref js.Ref) VideoFrameInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 VideoFrameInit VideoFrameInit [// VideoFrameInit] [0x140002c3cc0 0x140002c3e00 0x140002c3f40 0x14000312000 0x140003120a0 0x140003121e0 0x14000312320 0x140002c3d60 0x140002c3ea0 0x14000312140 0x14000312280] 0x1400121b290 {0 0}} in the application heap.
func (p VideoFrameInit) New() js.Ref {
	return bindings.VideoFrameInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoFrameInit) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoFrameInit) Update(ref js.Ref) {
	bindings.VideoFrameInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoPixelFormat uint32

const (
	_ VideoPixelFormat = iota

	VideoPixelFormat_I420
	VideoPixelFormat_I420_A
	VideoPixelFormat_I422
	VideoPixelFormat_I444
	VideoPixelFormat_NV12
	VideoPixelFormat_RGBA
	VideoPixelFormat_RGBX
	VideoPixelFormat_BGRA
	VideoPixelFormat_BGRX
)

func (VideoPixelFormat) FromRef(str js.Ref) VideoPixelFormat {
	return VideoPixelFormat(bindings.ConstOfVideoPixelFormat(str))
}

func (x VideoPixelFormat) String() (string, bool) {
	switch x {
	case VideoPixelFormat_I420:
		return "I420", true
	case VideoPixelFormat_I420_A:
		return "I420A", true
	case VideoPixelFormat_I422:
		return "I422", true
	case VideoPixelFormat_I444:
		return "I444", true
	case VideoPixelFormat_NV12:
		return "NV12", true
	case VideoPixelFormat_RGBA:
		return "RGBA", true
	case VideoPixelFormat_RGBX:
		return "RGBX", true
	case VideoPixelFormat_BGRA:
		return "BGRA", true
	case VideoPixelFormat_BGRX:
		return "BGRX", true
	default:
		return "", false
	}
}

type PlaneLayout struct {
	// Offset is "PlaneLayout.offset"
	//
	// Required
	Offset uint32
	// Stride is "PlaneLayout.stride"
	//
	// Required
	Stride uint32

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PlaneLayout with all fields set.
func (p PlaneLayout) FromRef(ref js.Ref) PlaneLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PlaneLayout PlaneLayout [// PlaneLayout] [0x14000312780 0x14000312820] 0x14000baa360 {0 0}} in the application heap.
func (p PlaneLayout) New() js.Ref {
	return bindings.PlaneLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PlaneLayout) UpdateFrom(ref js.Ref) {
	bindings.PlaneLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PlaneLayout) Update(ref js.Ref) {
	bindings.PlaneLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoColorPrimaries uint32

const (
	_ VideoColorPrimaries = iota

	VideoColorPrimaries_BT709
	VideoColorPrimaries_BT_470BG
	VideoColorPrimaries_SMPTE_170M
	VideoColorPrimaries_BT2020
	VideoColorPrimaries_SMPTE432
)

func (VideoColorPrimaries) FromRef(str js.Ref) VideoColorPrimaries {
	return VideoColorPrimaries(bindings.ConstOfVideoColorPrimaries(str))
}

func (x VideoColorPrimaries) String() (string, bool) {
	switch x {
	case VideoColorPrimaries_BT709:
		return "bt709", true
	case VideoColorPrimaries_BT_470BG:
		return "bt470bg", true
	case VideoColorPrimaries_SMPTE_170M:
		return "smpte170m", true
	case VideoColorPrimaries_BT2020:
		return "bt2020", true
	case VideoColorPrimaries_SMPTE432:
		return "smpte432", true
	default:
		return "", false
	}
}
