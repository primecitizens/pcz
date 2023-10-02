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

type OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext struct {
	ref js.Ref
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) Ref() js.Ref {
	return x.ref
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) Free() {
	x.ref.Free()
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) FromRef(ref js.Ref) OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext {
	return OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext{
		ref: ref,
	}
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) OffscreenCanvasRenderingContext2D() OffscreenCanvasRenderingContext2D {
	return OffscreenCanvasRenderingContext2D{}.FromRef(x.ref)
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) ImageBitmapRenderingContext() ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{}.FromRef(x.ref)
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) WebGLRenderingContext() WebGLRenderingContext {
	return WebGLRenderingContext{}.FromRef(x.ref)
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) WebGL2RenderingContext() WebGL2RenderingContext {
	return WebGL2RenderingContext{}.FromRef(x.ref)
}

func (x OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) GPUCanvasContext() GPUCanvasContext {
	return GPUCanvasContext{}.FromRef(x.ref)
}

type OffscreenRenderingContext = OneOf_OffscreenCanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext

type OffscreenRenderingContextId uint32

const (
	_ OffscreenRenderingContextId = iota

	OffscreenRenderingContextId_2D
	OffscreenRenderingContextId_BITMAPRENDERER
	OffscreenRenderingContextId_WEBGL
	OffscreenRenderingContextId_WEBGL2
	OffscreenRenderingContextId_WEBGPU
)

func (OffscreenRenderingContextId) FromRef(str js.Ref) OffscreenRenderingContextId {
	return OffscreenRenderingContextId(bindings.ConstOfOffscreenRenderingContextId(str))
}

func (x OffscreenRenderingContextId) String() (string, bool) {
	switch x {
	case OffscreenRenderingContextId_2D:
		return "2d", true
	case OffscreenRenderingContextId_BITMAPRENDERER:
		return "bitmaprenderer", true
	case OffscreenRenderingContextId_WEBGL:
		return "webgl", true
	case OffscreenRenderingContextId_WEBGL2:
		return "webgl2", true
	case OffscreenRenderingContextId_WEBGPU:
		return "webgpu", true
	default:
		return "", false
	}
}

type ImageEncodeOptions struct {
	// Type is "ImageEncodeOptions.type"
	//
	// Optional, defaults to "image/png".
	Type js.String
	// Quality is "ImageEncodeOptions.quality"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quality MUST be set to true to make this field effective.
	Quality float64

	FFI_USE_Quality bool // for Quality.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageEncodeOptions with all fields set.
func (p ImageEncodeOptions) FromRef(ref js.Ref) ImageEncodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageEncodeOptions in the application heap.
func (p ImageEncodeOptions) New() js.Ref {
	return bindings.ImageEncodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ImageEncodeOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageEncodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageEncodeOptions) Update(ref js.Ref) {
	bindings.ImageEncodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewOffscreenCanvas(width uint64, height uint64) OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(
		bindings.NewOffscreenCanvasByOffscreenCanvas(
			float64(width),
			float64(height)),
	)
}

type OffscreenCanvas struct {
	EventTarget
}

func (this OffscreenCanvas) Once() OffscreenCanvas {
	this.Ref().Once()
	return this
}

func (this OffscreenCanvas) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this OffscreenCanvas) FromRef(ref js.Ref) OffscreenCanvas {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this OffscreenCanvas) Free() {
	this.Ref().Free()
}

// Width returns the value of property "OffscreenCanvas.width".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvas) Width() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Width sets the value of property "OffscreenCanvas.width" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvas) SetWidth(val uint64) bool {
	return js.True == bindings.SetOffscreenCanvasWidth(
		this.Ref(),
		float64(val),
	)
}

// Height returns the value of property "OffscreenCanvas.height".
//
// The returned bool will be false if there is no such property.
func (this OffscreenCanvas) Height() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetOffscreenCanvasHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Height sets the value of property "OffscreenCanvas.height" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvas) SetHeight(val uint64) bool {
	return js.True == bindings.SetOffscreenCanvasHeight(
		this.Ref(),
		float64(val),
	)
}

// GetContext calls the method "OffscreenCanvas.getContext".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvas) GetContext(contextId OffscreenRenderingContextId, options js.Any) (OffscreenRenderingContext, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasGetContext(
		this.Ref(), js.Pointer(&_ok),
		uint32(contextId),
		options.Ref(),
	)

	return OffscreenRenderingContext{}.FromRef(_ret), _ok
}

// GetContextFunc returns the method "OffscreenCanvas.getContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvas) GetContextFunc() (fn js.Func[func(contextId OffscreenRenderingContextId, options js.Any) OffscreenRenderingContext]) {
	return fn.FromRef(
		bindings.OffscreenCanvasGetContextFunc(
			this.Ref(),
		),
	)
}

// GetContext1 calls the method "OffscreenCanvas.getContext".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvas) GetContext1(contextId OffscreenRenderingContextId) (OffscreenRenderingContext, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasGetContext1(
		this.Ref(), js.Pointer(&_ok),
		uint32(contextId),
	)

	return OffscreenRenderingContext{}.FromRef(_ret), _ok
}

// GetContext1Func returns the method "OffscreenCanvas.getContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvas) GetContext1Func() (fn js.Func[func(contextId OffscreenRenderingContextId) OffscreenRenderingContext]) {
	return fn.FromRef(
		bindings.OffscreenCanvasGetContext1Func(
			this.Ref(),
		),
	)
}

// TransferToImageBitmap calls the method "OffscreenCanvas.transferToImageBitmap".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvas) TransferToImageBitmap() (ImageBitmap, bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasTransferToImageBitmap(
		this.Ref(), js.Pointer(&_ok),
	)

	return ImageBitmap{}.FromRef(_ret), _ok
}

// TransferToImageBitmapFunc returns the method "OffscreenCanvas.transferToImageBitmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvas) TransferToImageBitmapFunc() (fn js.Func[func() ImageBitmap]) {
	return fn.FromRef(
		bindings.OffscreenCanvasTransferToImageBitmapFunc(
			this.Ref(),
		),
	)
}

// ConvertToBlob calls the method "OffscreenCanvas.convertToBlob".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvas) ConvertToBlob(options ImageEncodeOptions) (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasConvertToBlob(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// ConvertToBlobFunc returns the method "OffscreenCanvas.convertToBlob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvas) ConvertToBlobFunc() (fn js.Func[func(options ImageEncodeOptions) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.OffscreenCanvasConvertToBlobFunc(
			this.Ref(),
		),
	)
}

// ConvertToBlob1 calls the method "OffscreenCanvas.convertToBlob".
//
// The returned bool will be false if there is no such method.
func (this OffscreenCanvas) ConvertToBlob1() (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallOffscreenCanvasConvertToBlob1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// ConvertToBlob1Func returns the method "OffscreenCanvas.convertToBlob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OffscreenCanvas) ConvertToBlob1Func() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.OffscreenCanvasConvertToBlob1Func(
			this.Ref(),
		),
	)
}

type OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame struct {
	ref js.Ref
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) FromRef(ref js.Ref) OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame {
	return OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame{
		ref: ref,
	}
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) HTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) SVGImageElement() SVGImageElement {
	return SVGImageElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) HTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) HTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) ImageBitmap() ImageBitmap {
	return ImageBitmap{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) OffscreenCanvas() OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame) VideoFrame() VideoFrame {
	return VideoFrame{}.FromRef(x.ref)
}

type CanvasImageSource = OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame

type CanvasRenderingContext2D struct {
	ref js.Ref
}

func (this CanvasRenderingContext2D) Once() CanvasRenderingContext2D {
	this.Ref().Once()
	return this
}

func (this CanvasRenderingContext2D) Ref() js.Ref {
	return this.ref
}

func (this CanvasRenderingContext2D) FromRef(ref js.Ref) CanvasRenderingContext2D {
	this.ref = ref
	return this
}

func (this CanvasRenderingContext2D) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "CanvasRenderingContext2D.canvas".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) Canvas() (HTMLCanvasElement, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCanvasElement{}.FromRef(_ret), _ok
}

// GlobalAlpha returns the value of property "CanvasRenderingContext2D.globalAlpha".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) GlobalAlpha() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DGlobalAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// GlobalAlpha sets the value of property "CanvasRenderingContext2D.globalAlpha" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetGlobalAlpha(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DGlobalAlpha(
		this.Ref(),
		float64(val),
	)
}

// GlobalCompositeOperation returns the value of property "CanvasRenderingContext2D.globalCompositeOperation".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) GlobalCompositeOperation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DGlobalCompositeOperation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GlobalCompositeOperation sets the value of property "CanvasRenderingContext2D.globalCompositeOperation" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetGlobalCompositeOperation(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DGlobalCompositeOperation(
		this.Ref(),
		val.Ref(),
	)
}

// ImageSmoothingEnabled returns the value of property "CanvasRenderingContext2D.imageSmoothingEnabled".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DImageSmoothingEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ImageSmoothingEnabled sets the value of property "CanvasRenderingContext2D.imageSmoothingEnabled" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetImageSmoothingEnabled(val bool) bool {
	return js.True == bindings.SetCanvasRenderingContext2DImageSmoothingEnabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ImageSmoothingQuality returns the value of property "CanvasRenderingContext2D.imageSmoothingQuality".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingQuality() (ImageSmoothingQuality, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DImageSmoothingQuality(
		this.Ref(), js.Pointer(&_ok),
	)
	return ImageSmoothingQuality(_ret), _ok
}

// ImageSmoothingQuality sets the value of property "CanvasRenderingContext2D.imageSmoothingQuality" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) bool {
	return js.True == bindings.SetCanvasRenderingContext2DImageSmoothingQuality(
		this.Ref(),
		uint32(val),
	)
}

// StrokeStyle returns the value of property "CanvasRenderingContext2D.strokeStyle".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) StrokeStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DStrokeStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// StrokeStyle sets the value of property "CanvasRenderingContext2D.strokeStyle" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetStrokeStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetCanvasRenderingContext2DStrokeStyle(
		this.Ref(),
		val.Ref(),
	)
}

// FillStyle returns the value of property "CanvasRenderingContext2D.fillStyle".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) FillStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFillStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// FillStyle sets the value of property "CanvasRenderingContext2D.fillStyle" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFillStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFillStyle(
		this.Ref(),
		val.Ref(),
	)
}

// ShadowOffsetX returns the value of property "CanvasRenderingContext2D.shadowOffsetX".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DShadowOffsetX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetX sets the value of property "CanvasRenderingContext2D.shadowOffsetX" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowOffsetX(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowOffsetX(
		this.Ref(),
		float64(val),
	)
}

// ShadowOffsetY returns the value of property "CanvasRenderingContext2D.shadowOffsetY".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DShadowOffsetY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetY sets the value of property "CanvasRenderingContext2D.shadowOffsetY" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowOffsetY(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowOffsetY(
		this.Ref(),
		float64(val),
	)
}

// ShadowBlur returns the value of property "CanvasRenderingContext2D.shadowBlur".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ShadowBlur() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DShadowBlur(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowBlur sets the value of property "CanvasRenderingContext2D.shadowBlur" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowBlur(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowBlur(
		this.Ref(),
		float64(val),
	)
}

// ShadowColor returns the value of property "CanvasRenderingContext2D.shadowColor".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) ShadowColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DShadowColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ShadowColor sets the value of property "CanvasRenderingContext2D.shadowColor" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowColor(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowColor(
		this.Ref(),
		val.Ref(),
	)
}

// Filter returns the value of property "CanvasRenderingContext2D.filter".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) Filter() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFilter(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Filter sets the value of property "CanvasRenderingContext2D.filter" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFilter(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFilter(
		this.Ref(),
		val.Ref(),
	)
}

// LineWidth returns the value of property "CanvasRenderingContext2D.lineWidth".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) LineWidth() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DLineWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineWidth sets the value of property "CanvasRenderingContext2D.lineWidth" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineWidth(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineWidth(
		this.Ref(),
		float64(val),
	)
}

// LineCap returns the value of property "CanvasRenderingContext2D.lineCap".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) LineCap() (CanvasLineCap, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DLineCap(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineCap(_ret), _ok
}

// LineCap sets the value of property "CanvasRenderingContext2D.lineCap" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineCap(val CanvasLineCap) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineCap(
		this.Ref(),
		uint32(val),
	)
}

// LineJoin returns the value of property "CanvasRenderingContext2D.lineJoin".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) LineJoin() (CanvasLineJoin, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DLineJoin(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineJoin(_ret), _ok
}

// LineJoin sets the value of property "CanvasRenderingContext2D.lineJoin" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineJoin(val CanvasLineJoin) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineJoin(
		this.Ref(),
		uint32(val),
	)
}

// MiterLimit returns the value of property "CanvasRenderingContext2D.miterLimit".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) MiterLimit() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DMiterLimit(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MiterLimit sets the value of property "CanvasRenderingContext2D.miterLimit" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetMiterLimit(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DMiterLimit(
		this.Ref(),
		float64(val),
	)
}

// LineDashOffset returns the value of property "CanvasRenderingContext2D.lineDashOffset".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) LineDashOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DLineDashOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineDashOffset sets the value of property "CanvasRenderingContext2D.lineDashOffset" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineDashOffset(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineDashOffset(
		this.Ref(),
		float64(val),
	)
}

// Font returns the value of property "CanvasRenderingContext2D.font".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) Font() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFont(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Font sets the value of property "CanvasRenderingContext2D.font" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFont(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFont(
		this.Ref(),
		val.Ref(),
	)
}

// TextAlign returns the value of property "CanvasRenderingContext2D.textAlign".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) TextAlign() (CanvasTextAlign, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DTextAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextAlign(_ret), _ok
}

// TextAlign sets the value of property "CanvasRenderingContext2D.textAlign" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextAlign(val CanvasTextAlign) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextAlign(
		this.Ref(),
		uint32(val),
	)
}

// TextBaseline returns the value of property "CanvasRenderingContext2D.textBaseline".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) TextBaseline() (CanvasTextBaseline, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DTextBaseline(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextBaseline(_ret), _ok
}

// TextBaseline sets the value of property "CanvasRenderingContext2D.textBaseline" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextBaseline(val CanvasTextBaseline) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextBaseline(
		this.Ref(),
		uint32(val),
	)
}

// Direction returns the value of property "CanvasRenderingContext2D.direction".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) Direction() (CanvasDirection, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasDirection(_ret), _ok
}

// Direction sets the value of property "CanvasRenderingContext2D.direction" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetDirection(val CanvasDirection) bool {
	return js.True == bindings.SetCanvasRenderingContext2DDirection(
		this.Ref(),
		uint32(val),
	)
}

// LetterSpacing returns the value of property "CanvasRenderingContext2D.letterSpacing".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) LetterSpacing() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DLetterSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LetterSpacing sets the value of property "CanvasRenderingContext2D.letterSpacing" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLetterSpacing(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLetterSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// FontKerning returns the value of property "CanvasRenderingContext2D.fontKerning".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) FontKerning() (CanvasFontKerning, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFontKerning(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontKerning(_ret), _ok
}

// FontKerning sets the value of property "CanvasRenderingContext2D.fontKerning" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontKerning(val CanvasFontKerning) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontKerning(
		this.Ref(),
		uint32(val),
	)
}

// FontStretch returns the value of property "CanvasRenderingContext2D.fontStretch".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) FontStretch() (CanvasFontStretch, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFontStretch(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontStretch(_ret), _ok
}

// FontStretch sets the value of property "CanvasRenderingContext2D.fontStretch" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontStretch(val CanvasFontStretch) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontStretch(
		this.Ref(),
		uint32(val),
	)
}

// FontVariantCaps returns the value of property "CanvasRenderingContext2D.fontVariantCaps".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) FontVariantCaps() (CanvasFontVariantCaps, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DFontVariantCaps(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasFontVariantCaps(_ret), _ok
}

// FontVariantCaps sets the value of property "CanvasRenderingContext2D.fontVariantCaps" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontVariantCaps(val CanvasFontVariantCaps) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontVariantCaps(
		this.Ref(),
		uint32(val),
	)
}

// TextRendering returns the value of property "CanvasRenderingContext2D.textRendering".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) TextRendering() (CanvasTextRendering, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DTextRendering(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasTextRendering(_ret), _ok
}

// TextRendering sets the value of property "CanvasRenderingContext2D.textRendering" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextRendering(val CanvasTextRendering) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextRendering(
		this.Ref(),
		uint32(val),
	)
}

// WordSpacing returns the value of property "CanvasRenderingContext2D.wordSpacing".
//
// The returned bool will be false if there is no such property.
func (this CanvasRenderingContext2D) WordSpacing() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCanvasRenderingContext2DWordSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// WordSpacing sets the value of property "CanvasRenderingContext2D.wordSpacing" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetWordSpacing(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DWordSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// GetContextAttributes calls the method "CanvasRenderingContext2D.getContextAttributes".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) GetContextAttributes() (CanvasRenderingContext2DSettings, bool) {
	var _ret CanvasRenderingContext2DSettings
	_ok := js.True == bindings.CallCanvasRenderingContext2DGetContextAttributes(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetContextAttributesFunc returns the method "CanvasRenderingContext2D.getContextAttributes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) GetContextAttributesFunc() (fn js.Func[func() CanvasRenderingContext2DSettings]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// Save calls the method "CanvasRenderingContext2D.save".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Save() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DSave(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SaveFunc returns the method "CanvasRenderingContext2D.save".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) SaveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSaveFunc(
			this.Ref(),
		),
	)
}

// Restore calls the method "CanvasRenderingContext2D.restore".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Restore() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DRestore(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RestoreFunc returns the method "CanvasRenderingContext2D.restore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) RestoreFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRestoreFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "CanvasRenderingContext2D.reset".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "CanvasRenderingContext2D.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DResetFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "CanvasRenderingContext2D.isContextLost".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsContextLost() (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsContextLostFunc returns the method "CanvasRenderingContext2D.isContextLost".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsContextLostFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "CanvasRenderingContext2D.scale".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Scale(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DScale(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScaleFunc returns the method "CanvasRenderingContext2D.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ScaleFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScaleFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "CanvasRenderingContext2D.rotate".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Rotate(angle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DRotate(
		this.Ref(), js.Pointer(&_ok),
		float64(angle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RotateFunc returns the method "CanvasRenderingContext2D.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) RotateFunc() (fn js.Func[func(angle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRotateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "CanvasRenderingContext2D.translate".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Translate(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TranslateFunc returns the method "CanvasRenderingContext2D.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) TranslateFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DTranslateFunc(
			this.Ref(),
		),
	)
}

// Transform calls the method "CanvasRenderingContext2D.transform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DTransform(
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

// TransformFunc returns the method "CanvasRenderingContext2D.transform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) TransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DTransformFunc(
			this.Ref(),
		),
	)
}

// GetTransform calls the method "CanvasRenderingContext2D.getTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) GetTransform() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// GetTransformFunc returns the method "CanvasRenderingContext2D.getTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) GetTransformFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "CanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DSetTransform(
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

// SetTransformFunc returns the method "CanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) SetTransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "CanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform1Func returns the method "CanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) SetTransform1Func() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform2 calls the method "CanvasRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) SetTransform2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform2Func returns the method "CanvasRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) SetTransform2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransform2Func(
			this.Ref(),
		),
	)
}

// ResetTransform calls the method "CanvasRenderingContext2D.resetTransform".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ResetTransform() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetTransformFunc returns the method "CanvasRenderingContext2D.resetTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ResetTransformFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DResetTransformFunc(
			this.Ref(),
		),
	)
}

// CreateLinearGradient calls the method "CanvasRenderingContext2D.createLinearGradient".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateLinearGradientFunc returns the method "CanvasRenderingContext2D.createLinearGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateLinearGradientFunc() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateLinearGradientFunc(
			this.Ref(),
		),
	)
}

// CreateRadialGradient calls the method "CanvasRenderingContext2D.createRadialGradient".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateRadialGradient(
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

// CreateRadialGradientFunc returns the method "CanvasRenderingContext2D.createRadialGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateRadialGradientFunc() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateRadialGradientFunc(
			this.Ref(),
		),
	)
}

// CreateConicGradient calls the method "CanvasRenderingContext2D.createConicGradient".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateConicGradientFunc returns the method "CanvasRenderingContext2D.createConicGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateConicGradientFunc() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateConicGradientFunc(
			this.Ref(),
		),
	)
}

// CreatePattern calls the method "CanvasRenderingContext2D.createPattern".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		repetition.Ref(),
	)

	return CanvasPattern{}.FromRef(_ret), _ok
}

// CreatePatternFunc returns the method "CanvasRenderingContext2D.createPattern".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreatePatternFunc() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreatePatternFunc(
			this.Ref(),
		),
	)
}

// ClearRect calls the method "CanvasRenderingContext2D.clearRect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearRectFunc returns the method "CanvasRenderingContext2D.clearRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ClearRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClearRectFunc(
			this.Ref(),
		),
	)
}

// FillRect calls the method "CanvasRenderingContext2D.fillRect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillRectFunc returns the method "CanvasRenderingContext2D.fillRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) FillRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillRectFunc(
			this.Ref(),
		),
	)
}

// StrokeRect calls the method "CanvasRenderingContext2D.strokeRect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeRectFunc returns the method "CanvasRenderingContext2D.strokeRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) StrokeRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeRectFunc(
			this.Ref(),
		),
	)
}

// BeginPath calls the method "CanvasRenderingContext2D.beginPath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) BeginPath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginPathFunc returns the method "CanvasRenderingContext2D.beginPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) BeginPathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DBeginPathFunc(
			this.Ref(),
		),
	)
}

// Fill calls the method "CanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Fill(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFill(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillFunc returns the method "CanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) FillFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillFunc(
			this.Ref(),
		),
	)
}

// Fill1 calls the method "CanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Fill1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFill1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill1Func returns the method "CanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Fill1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill1Func(
			this.Ref(),
		),
	)
}

// Fill2 calls the method "CanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFill2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill2Func returns the method "CanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Fill2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill2Func(
			this.Ref(),
		),
	)
}

// Fill3 calls the method "CanvasRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Fill3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFill3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill3Func returns the method "CanvasRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Fill3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill3Func(
			this.Ref(),
		),
	)
}

// Stroke calls the method "CanvasRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Stroke() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DStroke(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeFunc returns the method "CanvasRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) StrokeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeFunc(
			this.Ref(),
		),
	)
}

// Stroke1 calls the method "CanvasRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Stroke1(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Stroke1Func returns the method "CanvasRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Stroke1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStroke1Func(
			this.Ref(),
		),
	)
}

// Clip calls the method "CanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Clip(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClip(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClipFunc returns the method "CanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ClipFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClipFunc(
			this.Ref(),
		),
	)
}

// Clip1 calls the method "CanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Clip1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClip1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip1Func returns the method "CanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Clip1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip1Func(
			this.Ref(),
		),
	)
}

// Clip2 calls the method "CanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClip2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip2Func returns the method "CanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Clip2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip2Func(
			this.Ref(),
		),
	)
}

// Clip3 calls the method "CanvasRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Clip3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClip3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip3Func returns the method "CanvasRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Clip3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip3Func(
			this.Ref(),
		),
	)
}

// IsPointInPath calls the method "CanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPathFunc returns the method "CanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPathFunc() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPathFunc(
			this.Ref(),
		),
	)
}

// IsPointInPath1 calls the method "CanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath1(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath1Func returns the method "CanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath1Func() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath1Func(
			this.Ref(),
		),
	)
}

// IsPointInPath2 calls the method "CanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPath2Func returns the method "CanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath2Func() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath2Func(
			this.Ref(),
		),
	)
}

// IsPointInPath3 calls the method "CanvasRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath3Func returns the method "CanvasRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInPath3Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath3Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "CanvasRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInStroke(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStrokeFunc returns the method "CanvasRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInStrokeFunc() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "CanvasRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStroke1Func returns the method "CanvasRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) IsPointInStroke1Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// DrawFocusIfNeeded calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) DrawFocusIfNeeded(element Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawFocusIfNeededFunc returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) DrawFocusIfNeededFunc() (fn js.Func[func(element Element)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawFocusIfNeededFunc(
			this.Ref(),
		),
	)
}

// DrawFocusIfNeeded1 calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) DrawFocusIfNeeded1(path Path2D, element Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		element.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawFocusIfNeeded1Func returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) DrawFocusIfNeeded1Func() (fn js.Func[func(path Path2D, element Element)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawFocusIfNeeded1Func(
			this.Ref(),
		),
	)
}

// ScrollPathIntoView calls the method "CanvasRenderingContext2D.scrollPathIntoView".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ScrollPathIntoView() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DScrollPathIntoView(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollPathIntoViewFunc returns the method "CanvasRenderingContext2D.scrollPathIntoView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ScrollPathIntoViewFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScrollPathIntoViewFunc(
			this.Ref(),
		),
	)
}

// ScrollPathIntoView1 calls the method "CanvasRenderingContext2D.scrollPathIntoView".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ScrollPathIntoView1(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DScrollPathIntoView1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollPathIntoView1Func returns the method "CanvasRenderingContext2D.scrollPathIntoView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ScrollPathIntoView1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScrollPathIntoView1Func(
			this.Ref(),
		),
	)
}

// FillText calls the method "CanvasRenderingContext2D.fillText".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) FillText(text js.String, x float64, y float64, maxWidth float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFillText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillTextFunc returns the method "CanvasRenderingContext2D.fillText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) FillTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillTextFunc(
			this.Ref(),
		),
	)
}

// FillText1 calls the method "CanvasRenderingContext2D.fillText".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) FillText1(text js.String, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DFillText1(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillText1Func returns the method "CanvasRenderingContext2D.fillText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) FillText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillText1Func(
			this.Ref(),
		),
	)
}

// StrokeText calls the method "CanvasRenderingContext2D.strokeText".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) StrokeText(text js.String, x float64, y float64, maxWidth float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DStrokeText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeTextFunc returns the method "CanvasRenderingContext2D.strokeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) StrokeTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeTextFunc(
			this.Ref(),
		),
	)
}

// StrokeText1 calls the method "CanvasRenderingContext2D.strokeText".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) StrokeText1(text js.String, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DStrokeText1(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeText1Func returns the method "CanvasRenderingContext2D.strokeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) StrokeText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeText1Func(
			this.Ref(),
		),
	)
}

// MeasureText calls the method "CanvasRenderingContext2D.measureText".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) MeasureText(text js.String) (TextMetrics, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DMeasureText(
		this.Ref(), js.Pointer(&_ok),
		text.Ref(),
	)

	return TextMetrics{}.FromRef(_ret), _ok
}

// MeasureTextFunc returns the method "CanvasRenderingContext2D.measureText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) MeasureTextFunc() (fn js.Func[func(text js.String) TextMetrics]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DMeasureTextFunc(
			this.Ref(),
		),
	)
}

// DrawImage calls the method "CanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImageFunc returns the method "CanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) DrawImageFunc() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImageFunc(
			this.Ref(),
		),
	)
}

// DrawImage1 calls the method "CanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DDrawImage1(
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

// DrawImage1Func returns the method "CanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) DrawImage1Func() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImage1Func(
			this.Ref(),
		),
	)
}

// DrawImage2 calls the method "CanvasRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DDrawImage2(
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

// DrawImage2Func returns the method "CanvasRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) DrawImage2Func() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImage2Func(
			this.Ref(),
		),
	)
}

// CreateImageData calls the method "CanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateImageData(sw int32, sh int32, settings ImageDataSettings) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateImageData(
		this.Ref(), js.Pointer(&_ok),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageDataFunc returns the method "CanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateImageDataFunc() (fn js.Func[func(sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageDataFunc(
			this.Ref(),
		),
	)
}

// CreateImageData1 calls the method "CanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateImageData1(sw int32, sh int32) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateImageData1(
		this.Ref(), js.Pointer(&_ok),
		int32(sw),
		int32(sh),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageData1Func returns the method "CanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateImageData1Func() (fn js.Func[func(sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageData1Func(
			this.Ref(),
		),
	)
}

// CreateImageData2 calls the method "CanvasRenderingContext2D.createImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) CreateImageData2(imagedata ImageData) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DCreateImageData2(
		this.Ref(), js.Pointer(&_ok),
		imagedata.Ref(),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// CreateImageData2Func returns the method "CanvasRenderingContext2D.createImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) CreateImageData2Func() (fn js.Func[func(imagedata ImageData) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageData2Func(
			this.Ref(),
		),
	)
}

// GetImageData calls the method "CanvasRenderingContext2D.getImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) GetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DGetImageData(
		this.Ref(), js.Pointer(&_ok),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// GetImageDataFunc returns the method "CanvasRenderingContext2D.getImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) GetImageDataFunc() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetImageDataFunc(
			this.Ref(),
		),
	)
}

// GetImageData1 calls the method "CanvasRenderingContext2D.getImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) GetImageData1(sx int32, sy int32, sw int32, sh int32) (ImageData, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DGetImageData1(
		this.Ref(), js.Pointer(&_ok),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return ImageData{}.FromRef(_ret), _ok
}

// GetImageData1Func returns the method "CanvasRenderingContext2D.getImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) GetImageData1Func() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetImageData1Func(
			this.Ref(),
		),
	)
}

// PutImageData calls the method "CanvasRenderingContext2D.putImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) PutImageData(imagedata ImageData, dx int32, dy int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DPutImageData(
		this.Ref(), js.Pointer(&_ok),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PutImageDataFunc returns the method "CanvasRenderingContext2D.putImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) PutImageDataFunc() (fn js.Func[func(imagedata ImageData, dx int32, dy int32)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DPutImageDataFunc(
			this.Ref(),
		),
	)
}

// PutImageData1 calls the method "CanvasRenderingContext2D.putImageData".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) PutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DPutImageData1(
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

// PutImageData1Func returns the method "CanvasRenderingContext2D.putImageData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) PutImageData1Func() (fn js.Func[func(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DPutImageData1Func(
			this.Ref(),
		),
	)
}

// SetLineDash calls the method "CanvasRenderingContext2D.setLineDash".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) SetLineDash(segments js.Array[float64]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&_ok),
		segments.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetLineDashFunc returns the method "CanvasRenderingContext2D.setLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) SetLineDashFunc() (fn js.Func[func(segments js.Array[float64])]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetLineDashFunc(
			this.Ref(),
		),
	)
}

// GetLineDash calls the method "CanvasRenderingContext2D.getLineDash".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) GetLineDash() (js.Array[float64], bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[float64]{}.FromRef(_ret), _ok
}

// GetLineDashFunc returns the method "CanvasRenderingContext2D.getLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) GetLineDashFunc() (fn js.Func[func() js.Array[float64]]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetLineDashFunc(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "CanvasRenderingContext2D.closePath".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ClosePath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClosePathFunc returns the method "CanvasRenderingContext2D.closePath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClosePathFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "CanvasRenderingContext2D.moveTo".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) MoveTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveToFunc returns the method "CanvasRenderingContext2D.moveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DMoveToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "CanvasRenderingContext2D.lineTo".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) LineTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineToFunc returns the method "CanvasRenderingContext2D.lineTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DLineToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "CanvasRenderingContext2D.quadraticCurveTo".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QuadraticCurveToFunc returns the method "CanvasRenderingContext2D.quadraticCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "CanvasRenderingContext2D.bezierCurveTo".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DBezierCurveTo(
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

// BezierCurveToFunc returns the method "CanvasRenderingContext2D.bezierCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "CanvasRenderingContext2D.arcTo".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DArcTo(
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

// ArcToFunc returns the method "CanvasRenderingContext2D.arcTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArcToFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "CanvasRenderingContext2D.rect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RectFunc returns the method "CanvasRenderingContext2D.rect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "CanvasRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DRoundRect(
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

// RoundRectFunc returns the method "CanvasRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "CanvasRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRect1Func returns the method "CanvasRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// Arc calls the method "CanvasRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DArc(
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

// ArcFunc returns the method "CanvasRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "CanvasRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DArc1(
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

// Arc1Func returns the method "CanvasRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArc1Func(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "CanvasRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DEllipse(
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

// EllipseFunc returns the method "CanvasRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "CanvasRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this CanvasRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCanvasRenderingContext2DEllipse1(
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

// Ellipse1Func returns the method "CanvasRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CanvasRenderingContext2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DEllipse1Func(
			this.Ref(),
		),
	)
}

type OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext struct {
	ref js.Ref
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) Free() {
	x.ref.Free()
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) FromRef(ref js.Ref) OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext {
	return OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext{
		ref: ref,
	}
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) CanvasRenderingContext2D() CanvasRenderingContext2D {
	return CanvasRenderingContext2D{}.FromRef(x.ref)
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) ImageBitmapRenderingContext() ImageBitmapRenderingContext {
	return ImageBitmapRenderingContext{}.FromRef(x.ref)
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) WebGLRenderingContext() WebGLRenderingContext {
	return WebGLRenderingContext{}.FromRef(x.ref)
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) WebGL2RenderingContext() WebGL2RenderingContext {
	return WebGL2RenderingContext{}.FromRef(x.ref)
}

func (x OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext) GPUCanvasContext() GPUCanvasContext {
	return GPUCanvasContext{}.FromRef(x.ref)
}

type RenderingContext = OneOf_CanvasRenderingContext2D_ImageBitmapRenderingContext_WebGLRenderingContext_WebGL2RenderingContext_GPUCanvasContext

type BlobCallbackFunc func(this js.Ref, blob Blob) js.Ref

func (fn BlobCallbackFunc) Register() js.Func[func(blob Blob)] {
	return js.RegisterCallback[func(blob Blob)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BlobCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Blob{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BlobCallback[T any] struct {
	Fn  func(arg T, this js.Ref, blob Blob) js.Ref
	Arg T
}

func (cb *BlobCallback[T]) Register() js.Func[func(blob Blob)] {
	return js.RegisterCallback[func(blob Blob)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BlobCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		Blob{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

func NewHTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(
		bindings.NewHTMLCanvasElementByHTMLCanvasElement(),
	)
}

type HTMLCanvasElement struct {
	HTMLElement
}

func (this HTMLCanvasElement) Once() HTMLCanvasElement {
	this.Ref().Once()
	return this
}

func (this HTMLCanvasElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLCanvasElement) FromRef(ref js.Ref) HTMLCanvasElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLCanvasElement) Free() {
	this.Ref().Free()
}

// Width returns the value of property "HTMLCanvasElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLCanvasElement) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLCanvasElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Width sets the value of property "HTMLCanvasElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLCanvasElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLCanvasElementWidth(
		this.Ref(),
		uint32(val),
	)
}

// Height returns the value of property "HTMLCanvasElement.height".
//
// The returned bool will be false if there is no such property.
func (this HTMLCanvasElement) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLCanvasElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height sets the value of property "HTMLCanvasElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLCanvasElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLCanvasElementHeight(
		this.Ref(),
		uint32(val),
	)
}

// GetContext calls the method "HTMLCanvasElement.getContext".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) GetContext(contextId js.String, options js.Any) (RenderingContext, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementGetContext(
		this.Ref(), js.Pointer(&_ok),
		contextId.Ref(),
		options.Ref(),
	)

	return RenderingContext{}.FromRef(_ret), _ok
}

// GetContextFunc returns the method "HTMLCanvasElement.getContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) GetContextFunc() (fn js.Func[func(contextId js.String, options js.Any) RenderingContext]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementGetContextFunc(
			this.Ref(),
		),
	)
}

// GetContext1 calls the method "HTMLCanvasElement.getContext".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) GetContext1(contextId js.String) (RenderingContext, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementGetContext1(
		this.Ref(), js.Pointer(&_ok),
		contextId.Ref(),
	)

	return RenderingContext{}.FromRef(_ret), _ok
}

// GetContext1Func returns the method "HTMLCanvasElement.getContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) GetContext1Func() (fn js.Func[func(contextId js.String) RenderingContext]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementGetContext1Func(
			this.Ref(),
		),
	)
}

// ToDataURL calls the method "HTMLCanvasElement.toDataURL".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToDataURL(typ js.String, quality js.Any) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToDataURL(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		quality.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToDataURLFunc returns the method "HTMLCanvasElement.toDataURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToDataURLFunc() (fn js.Func[func(typ js.String, quality js.Any) js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURLFunc(
			this.Ref(),
		),
	)
}

// ToDataURL1 calls the method "HTMLCanvasElement.toDataURL".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToDataURL1(typ js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToDataURL1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToDataURL1Func returns the method "HTMLCanvasElement.toDataURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToDataURL1Func() (fn js.Func[func(typ js.String) js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURL1Func(
			this.Ref(),
		),
	)
}

// ToDataURL2 calls the method "HTMLCanvasElement.toDataURL".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToDataURL2() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToDataURL2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToDataURL2Func returns the method "HTMLCanvasElement.toDataURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToDataURL2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURL2Func(
			this.Ref(),
		),
	)
}

// ToBlob calls the method "HTMLCanvasElement.toBlob".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToBlob(callback js.Func[func(blob Blob)], typ js.String, quality js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToBlob(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
		typ.Ref(),
		quality.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ToBlobFunc returns the method "HTMLCanvasElement.toBlob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToBlobFunc() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String, quality js.Any)]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlobFunc(
			this.Ref(),
		),
	)
}

// ToBlob1 calls the method "HTMLCanvasElement.toBlob".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToBlob1(callback js.Func[func(blob Blob)], typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToBlob1(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ToBlob1Func returns the method "HTMLCanvasElement.toBlob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToBlob1Func() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String)]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlob1Func(
			this.Ref(),
		),
	)
}

// ToBlob2 calls the method "HTMLCanvasElement.toBlob".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) ToBlob2(callback js.Func[func(blob Blob)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementToBlob2(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ToBlob2Func returns the method "HTMLCanvasElement.toBlob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) ToBlob2Func() (fn js.Func[func(callback js.Func[func(blob Blob)])]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlob2Func(
			this.Ref(),
		),
	)
}

// TransferControlToOffscreen calls the method "HTMLCanvasElement.transferControlToOffscreen".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) TransferControlToOffscreen() (OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementTransferControlToOffscreen(
		this.Ref(), js.Pointer(&_ok),
	)

	return OffscreenCanvas{}.FromRef(_ret), _ok
}

// TransferControlToOffscreenFunc returns the method "HTMLCanvasElement.transferControlToOffscreen".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) TransferControlToOffscreenFunc() (fn js.Func[func() OffscreenCanvas]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementTransferControlToOffscreenFunc(
			this.Ref(),
		),
	)
}

// CaptureStream calls the method "HTMLCanvasElement.captureStream".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) CaptureStream(frameRequestRate float64) (MediaStream, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementCaptureStream(
		this.Ref(), js.Pointer(&_ok),
		float64(frameRequestRate),
	)

	return MediaStream{}.FromRef(_ret), _ok
}

// CaptureStreamFunc returns the method "HTMLCanvasElement.captureStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) CaptureStreamFunc() (fn js.Func[func(frameRequestRate float64) MediaStream]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementCaptureStreamFunc(
			this.Ref(),
		),
	)
}

// CaptureStream1 calls the method "HTMLCanvasElement.captureStream".
//
// The returned bool will be false if there is no such method.
func (this HTMLCanvasElement) CaptureStream1() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.CallHTMLCanvasElementCaptureStream1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaStream{}.FromRef(_ret), _ok
}

// CaptureStream1Func returns the method "HTMLCanvasElement.captureStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLCanvasElement) CaptureStream1Func() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementCaptureStream1Func(
			this.Ref(),
		),
	)
}

type OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData struct {
	ref js.Ref
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) FromRef(ref js.Ref) OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData {
	return OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData{
		ref: ref,
	}
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) HTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) SVGImageElement() SVGImageElement {
	return SVGImageElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) HTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) HTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) ImageBitmap() ImageBitmap {
	return ImageBitmap{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) OffscreenCanvas() OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) VideoFrame() VideoFrame {
	return VideoFrame{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData) ImageData() ImageData {
	return ImageData{}.FromRef(x.ref)
}

type ImageBitmapSource = OneOf_HTMLImageElement_SVGImageElement_HTMLVideoElement_HTMLCanvasElement_ImageBitmap_OffscreenCanvas_VideoFrame_Blob_ImageData

func NewBarcodeDetector(barcodeDetectorOptions BarcodeDetectorOptions) BarcodeDetector {
	return BarcodeDetector{}.FromRef(
		bindings.NewBarcodeDetectorByBarcodeDetector(
			js.Pointer(&barcodeDetectorOptions)),
	)
}

func NewBarcodeDetectorByBarcodeDetector1() BarcodeDetector {
	return BarcodeDetector{}.FromRef(
		bindings.NewBarcodeDetectorByBarcodeDetector1(),
	)
}

type BarcodeDetector struct {
	ref js.Ref
}

func (this BarcodeDetector) Once() BarcodeDetector {
	this.Ref().Once()
	return this
}

func (this BarcodeDetector) Ref() js.Ref {
	return this.ref
}

func (this BarcodeDetector) FromRef(ref js.Ref) BarcodeDetector {
	this.ref = ref
	return this
}

func (this BarcodeDetector) Free() {
	this.Ref().Free()
}

// GetSupportedFormats calls the staticmethod "BarcodeDetector.getSupportedFormats".
//
// The returned bool will be false if there is no such method.
func (this BarcodeDetector) GetSupportedFormats() (js.Promise[js.Array[BarcodeFormat]], bool) {
	var _ok bool
	_ret := bindings.CallBarcodeDetectorGetSupportedFormats(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BarcodeFormat]]{}.FromRef(_ret), _ok
}

// GetSupportedFormatsFunc returns the staticmethod "BarcodeDetector.getSupportedFormats".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BarcodeDetector) GetSupportedFormatsFunc() (fn js.Func[func() js.Promise[js.Array[BarcodeFormat]]]) {
	return fn.FromRef(
		bindings.BarcodeDetectorGetSupportedFormatsFunc(
			this.Ref(),
		),
	)
}

// Detect calls the method "BarcodeDetector.detect".
//
// The returned bool will be false if there is no such method.
func (this BarcodeDetector) Detect(image ImageBitmapSource) (js.Promise[js.Array[DetectedBarcode]], bool) {
	var _ok bool
	_ret := bindings.CallBarcodeDetectorDetect(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
	)

	return js.Promise[js.Array[DetectedBarcode]]{}.FromRef(_ret), _ok
}

// DetectFunc returns the method "BarcodeDetector.detect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BarcodeDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedBarcode]]]) {
	return fn.FromRef(
		bindings.BarcodeDetectorDetectFunc(
			this.Ref(),
		),
	)
}

type CompositeOperationOrAuto uint32

const (
	_ CompositeOperationOrAuto = iota

	CompositeOperationOrAuto_REPLACE
	CompositeOperationOrAuto_ADD
	CompositeOperationOrAuto_ACCUMULATE
	CompositeOperationOrAuto_AUTO
)

func (CompositeOperationOrAuto) FromRef(str js.Ref) CompositeOperationOrAuto {
	return CompositeOperationOrAuto(bindings.ConstOfCompositeOperationOrAuto(str))
}

func (x CompositeOperationOrAuto) String() (string, bool) {
	switch x {
	case CompositeOperationOrAuto_REPLACE:
		return "replace", true
	case CompositeOperationOrAuto_ADD:
		return "add", true
	case CompositeOperationOrAuto_ACCUMULATE:
		return "accumulate", true
	case CompositeOperationOrAuto_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type BaseComputedKeyframe struct {
	// Offset is "BaseComputedKeyframe.offset"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset float64
	// ComputedOffset is "BaseComputedKeyframe.computedOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_ComputedOffset MUST be set to true to make this field effective.
	ComputedOffset float64
	// Easing is "BaseComputedKeyframe.easing"
	//
	// Optional, defaults to "linear".
	Easing js.String
	// Composite is "BaseComputedKeyframe.composite"
	//
	// Optional, defaults to "auto".
	Composite CompositeOperationOrAuto

	FFI_USE_Offset         bool // for Offset.
	FFI_USE_ComputedOffset bool // for ComputedOffset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BaseComputedKeyframe with all fields set.
func (p BaseComputedKeyframe) FromRef(ref js.Ref) BaseComputedKeyframe {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BaseComputedKeyframe in the application heap.
func (p BaseComputedKeyframe) New() js.Ref {
	return bindings.BaseComputedKeyframeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BaseComputedKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BaseComputedKeyframeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BaseComputedKeyframe) Update(ref js.Ref) {
	bindings.BaseComputedKeyframeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BaseKeyframe struct {
	// Offset is "BaseKeyframe.offset"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset float64
	// Easing is "BaseKeyframe.easing"
	//
	// Optional, defaults to "linear".
	Easing js.String
	// Composite is "BaseKeyframe.composite"
	//
	// Optional, defaults to "auto".
	Composite CompositeOperationOrAuto

	FFI_USE_Offset bool // for Offset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BaseKeyframe with all fields set.
func (p BaseKeyframe) FromRef(ref js.Ref) BaseKeyframe {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BaseKeyframe in the application heap.
func (p BaseKeyframe) New() js.Ref {
	return bindings.BaseKeyframeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BaseKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BaseKeyframeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BaseKeyframe) Update(ref js.Ref) {
	bindings.BaseKeyframeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Float64_ArrayFloat64_Null struct {
	ref js.Ref
}

func (x OneOf_Float64_ArrayFloat64_Null) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_ArrayFloat64_Null) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_ArrayFloat64_Null) FromRef(ref js.Ref) OneOf_Float64_ArrayFloat64_Null {
	return OneOf_Float64_ArrayFloat64_Null{
		ref: ref,
	}
}

func (x OneOf_Float64_ArrayFloat64_Null) Null() bool {
	return x.ref == js.Null
}

func (x OneOf_Float64_ArrayFloat64_Null) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_ArrayFloat64_Null) ArrayFloat64() js.Array[float64] {
	return js.Array[float64]{}.FromRef(x.ref)
}

type OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto struct {
	ref js.Ref
}

func (x OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto) Free() {
	x.ref.Free()
}

func (x OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto) FromRef(ref js.Ref) OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto {
	return OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto{
		ref: ref,
	}
}

func (x OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto) CompositeOperationOrAuto() CompositeOperationOrAuto {
	return CompositeOperationOrAuto(0).FromRef(x.ref)
}

func (x OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto) ArrayCompositeOperationOrAuto() js.Array[CompositeOperationOrAuto] {
	return js.Array[CompositeOperationOrAuto]{}.FromRef(x.ref)
}

type BasePropertyIndexedKeyframe struct {
	// Offset is "BasePropertyIndexedKeyframe.offset"
	//
	// Optional, defaults to [].
	Offset OneOf_Float64_ArrayFloat64_Null
	// Easing is "BasePropertyIndexedKeyframe.easing"
	//
	// Optional, defaults to [].
	Easing OneOf_String_ArrayString
	// Composite is "BasePropertyIndexedKeyframe.composite"
	//
	// Optional, defaults to [].
	Composite OneOf_CompositeOperationOrAuto_ArrayCompositeOperationOrAuto

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BasePropertyIndexedKeyframe with all fields set.
func (p BasePropertyIndexedKeyframe) FromRef(ref js.Ref) BasePropertyIndexedKeyframe {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BasePropertyIndexedKeyframe in the application heap.
func (p BasePropertyIndexedKeyframe) New() js.Ref {
	return bindings.BasePropertyIndexedKeyframeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BasePropertyIndexedKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BasePropertyIndexedKeyframeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BasePropertyIndexedKeyframe) Update(ref js.Ref) {
	bindings.BasePropertyIndexedKeyframeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BatteryManager struct {
	EventTarget
}

func (this BatteryManager) Once() BatteryManager {
	this.Ref().Once()
	return this
}

func (this BatteryManager) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BatteryManager) FromRef(ref js.Ref) BatteryManager {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BatteryManager) Free() {
	this.Ref().Free()
}

// Charging returns the value of property "BatteryManager.charging".
//
// The returned bool will be false if there is no such property.
func (this BatteryManager) Charging() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBatteryManagerCharging(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ChargingTime returns the value of property "BatteryManager.chargingTime".
//
// The returned bool will be false if there is no such property.
func (this BatteryManager) ChargingTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetBatteryManagerChargingTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DischargingTime returns the value of property "BatteryManager.dischargingTime".
//
// The returned bool will be false if there is no such property.
func (this BatteryManager) DischargingTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetBatteryManagerDischargingTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Level returns the value of property "BatteryManager.level".
//
// The returned bool will be false if there is no such property.
func (this BatteryManager) Level() (float64, bool) {
	var _ok bool
	_ret := bindings.GetBatteryManagerLevel(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type PromptResponseObject struct {
	// UserChoice is "PromptResponseObject.userChoice"
	//
	// Optional
	UserChoice AppBannerPromptOutcome

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PromptResponseObject with all fields set.
func (p PromptResponseObject) FromRef(ref js.Ref) PromptResponseObject {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PromptResponseObject in the application heap.
func (p PromptResponseObject) New() js.Ref {
	return bindings.PromptResponseObjectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PromptResponseObject) UpdateFrom(ref js.Ref) {
	bindings.PromptResponseObjectJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PromptResponseObject) Update(ref js.Ref) {
	bindings.PromptResponseObjectJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewBeforeInstallPromptEvent(typ js.String, eventInitDict EventInit) BeforeInstallPromptEvent {
	return BeforeInstallPromptEvent{}.FromRef(
		bindings.NewBeforeInstallPromptEventByBeforeInstallPromptEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewBeforeInstallPromptEventByBeforeInstallPromptEvent1(typ js.String) BeforeInstallPromptEvent {
	return BeforeInstallPromptEvent{}.FromRef(
		bindings.NewBeforeInstallPromptEventByBeforeInstallPromptEvent1(
			typ.Ref()),
	)
}

type BeforeInstallPromptEvent struct {
	Event
}

func (this BeforeInstallPromptEvent) Once() BeforeInstallPromptEvent {
	this.Ref().Once()
	return this
}

func (this BeforeInstallPromptEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this BeforeInstallPromptEvent) FromRef(ref js.Ref) BeforeInstallPromptEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this BeforeInstallPromptEvent) Free() {
	this.Ref().Free()
}

// Prompt calls the method "BeforeInstallPromptEvent.prompt".
//
// The returned bool will be false if there is no such method.
func (this BeforeInstallPromptEvent) Prompt() (js.Promise[PromptResponseObject], bool) {
	var _ok bool
	_ret := bindings.CallBeforeInstallPromptEventPrompt(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PromptResponseObject]{}.FromRef(_ret), _ok
}

// PromptFunc returns the method "BeforeInstallPromptEvent.prompt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BeforeInstallPromptEvent) PromptFunc() (fn js.Func[func() js.Promise[PromptResponseObject]]) {
	return fn.FromRef(
		bindings.BeforeInstallPromptEventPromptFunc(
			this.Ref(),
		),
	)
}

func NewBeforeUnloadEvent(typ js.String, eventInitDict EventInit) BeforeUnloadEvent {
	return BeforeUnloadEvent{}.FromRef(
		bindings.NewBeforeUnloadEventByBeforeUnloadEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewBeforeUnloadEventByBeforeUnloadEvent1(typ js.String) BeforeUnloadEvent {
	return BeforeUnloadEvent{}.FromRef(
		bindings.NewBeforeUnloadEventByBeforeUnloadEvent1(
			typ.Ref()),
	)
}

type BeforeUnloadEvent struct {
	Event
}

func (this BeforeUnloadEvent) Once() BeforeUnloadEvent {
	this.Ref().Once()
	return this
}

func (this BeforeUnloadEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this BeforeUnloadEvent) FromRef(ref js.Ref) BeforeUnloadEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this BeforeUnloadEvent) Free() {
	this.Ref().Free()
}

// ReturnValue returns the value of property "BeforeUnloadEvent.returnValue".
//
// The returned bool will be false if there is no such property.
func (this BeforeUnloadEvent) ReturnValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBeforeUnloadEventReturnValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReturnValue sets the value of property "BeforeUnloadEvent.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this BeforeUnloadEvent) SetReturnValue(val js.String) bool {
	return js.True == bindings.SetBeforeUnloadEventReturnValue(
		this.Ref(),
		val.Ref(),
	)
}

type PreviousWin struct {
	// TimeDelta is "PreviousWin.timeDelta"
	//
	// Required
	TimeDelta int64
	// AdJSON is "PreviousWin.adJSON"
	//
	// Required
	AdJSON js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PreviousWin with all fields set.
func (p PreviousWin) FromRef(ref js.Ref) PreviousWin {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PreviousWin in the application heap.
func (p PreviousWin) New() js.Ref {
	return bindings.PreviousWinJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PreviousWin) UpdateFrom(ref js.Ref) {
	bindings.PreviousWinJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PreviousWin) Update(ref js.Ref) {
	bindings.PreviousWinJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BiddingBrowserSignals struct {
	// TopWindowHostname is "BiddingBrowserSignals.topWindowHostname"
	//
	// Required
	TopWindowHostname js.String
	// Seller is "BiddingBrowserSignals.seller"
	//
	// Required
	Seller js.String
	// JoinCount is "BiddingBrowserSignals.joinCount"
	//
	// Required
	JoinCount int32
	// BidCount is "BiddingBrowserSignals.bidCount"
	//
	// Required
	BidCount int32
	// Recency is "BiddingBrowserSignals.recency"
	//
	// Required
	Recency int32
	// TopLevelSeller is "BiddingBrowserSignals.topLevelSeller"
	//
	// Optional
	TopLevelSeller js.String
	// PrevWinsMs is "BiddingBrowserSignals.prevWinsMs"
	//
	// Optional
	PrevWinsMs js.Array[PreviousWin]
	// WasmHelper is "BiddingBrowserSignals.wasmHelper"
	//
	// Optional
	WasmHelper js.Object
	// DataVersion is "BiddingBrowserSignals.dataVersion"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataVersion MUST be set to true to make this field effective.
	DataVersion uint32

	FFI_USE_DataVersion bool // for DataVersion.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BiddingBrowserSignals with all fields set.
func (p BiddingBrowserSignals) FromRef(ref js.Ref) BiddingBrowserSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BiddingBrowserSignals in the application heap.
func (p BiddingBrowserSignals) New() js.Ref {
	return bindings.BiddingBrowserSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BiddingBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.BiddingBrowserSignalsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BiddingBrowserSignals) Update(ref js.Ref) {
	bindings.BiddingBrowserSignalsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BigInteger = js.TypedArray[uint8]

type OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView struct {
	ref js.Ref
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FromRef(ref js.Ref) OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView {
	return OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView{
		ref: ref,
	}
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

type BinaryData = OneOf_ArrayBuffer_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView

type BinaryType uint32

const (
	_ BinaryType = iota

	BinaryType_BLOB
	BinaryType_ARRAYBUFFER
)

func (BinaryType) FromRef(str js.Ref) BinaryType {
	return BinaryType(bindings.ConstOfBinaryType(str))
}

func (x BinaryType) String() (string, bool) {
	switch x {
	case BinaryType_BLOB:
		return "blob", true
	case BinaryType_ARRAYBUFFER:
		return "arraybuffer", true
	default:
		return "", false
	}
}

type BlobEventInit struct {
	// Data is "BlobEventInit.data"
	//
	// Required
	Data Blob
	// Timecode is "BlobEventInit.timecode"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timecode MUST be set to true to make this field effective.
	Timecode DOMHighResTimeStamp

	FFI_USE_Timecode bool // for Timecode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BlobEventInit with all fields set.
func (p BlobEventInit) FromRef(ref js.Ref) BlobEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BlobEventInit in the application heap.
func (p BlobEventInit) New() js.Ref {
	return bindings.BlobEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BlobEventInit) UpdateFrom(ref js.Ref) {
	bindings.BlobEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BlobEventInit) Update(ref js.Ref) {
	bindings.BlobEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewBlobEvent(typ js.String, eventInitDict BlobEventInit) BlobEvent {
	return BlobEvent{}.FromRef(
		bindings.NewBlobEventByBlobEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type BlobEvent struct {
	Event
}

func (this BlobEvent) Once() BlobEvent {
	this.Ref().Once()
	return this
}

func (this BlobEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this BlobEvent) FromRef(ref js.Ref) BlobEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this BlobEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "BlobEvent.data".
//
// The returned bool will be false if there is no such property.
func (this BlobEvent) Data() (Blob, bool) {
	var _ok bool
	_ret := bindings.GetBlobEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return Blob{}.FromRef(_ret), _ok
}

// Timecode returns the value of property "BlobEvent.timecode".
//
// The returned bool will be false if there is no such property.
func (this BlobEvent) Timecode() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetBlobEventTimecode(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

type BlockFragmentationType uint32

const (
	_ BlockFragmentationType = iota

	BlockFragmentationType_NONE
	BlockFragmentationType_PAGE
	BlockFragmentationType_COLUMN
	BlockFragmentationType_REGION
)

func (BlockFragmentationType) FromRef(str js.Ref) BlockFragmentationType {
	return BlockFragmentationType(bindings.ConstOfBlockFragmentationType(str))
}

func (x BlockFragmentationType) String() (string, bool) {
	switch x {
	case BlockFragmentationType_NONE:
		return "none", true
	case BlockFragmentationType_PAGE:
		return "page", true
	case BlockFragmentationType_COLUMN:
		return "column", true
	case BlockFragmentationType_REGION:
		return "region", true
	default:
		return "", false
	}
}

type WatchAdvertisementsOptions struct {
	// Signal is "WatchAdvertisementsOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WatchAdvertisementsOptions with all fields set.
func (p WatchAdvertisementsOptions) FromRef(ref js.Ref) WatchAdvertisementsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WatchAdvertisementsOptions in the application heap.
func (p WatchAdvertisementsOptions) New() js.Ref {
	return bindings.WatchAdvertisementsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WatchAdvertisementsOptions) UpdateFrom(ref js.Ref) {
	bindings.WatchAdvertisementsOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WatchAdvertisementsOptions) Update(ref js.Ref) {
	bindings.WatchAdvertisementsOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type BluetoothRemoteGATTDescriptor struct {
	ref js.Ref
}

func (this BluetoothRemoteGATTDescriptor) Once() BluetoothRemoteGATTDescriptor {
	this.Ref().Once()
	return this
}

func (this BluetoothRemoteGATTDescriptor) Ref() js.Ref {
	return this.ref
}

func (this BluetoothRemoteGATTDescriptor) FromRef(ref js.Ref) BluetoothRemoteGATTDescriptor {
	this.ref = ref
	return this
}

func (this BluetoothRemoteGATTDescriptor) Free() {
	this.Ref().Free()
}

// Characteristic returns the value of property "BluetoothRemoteGATTDescriptor.characteristic".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Characteristic() (BluetoothRemoteGATTCharacteristic, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTDescriptorCharacteristic(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothRemoteGATTCharacteristic{}.FromRef(_ret), _ok
}

// Uuid returns the value of property "BluetoothRemoteGATTDescriptor.uuid".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Uuid() (UUID, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTDescriptorUuid(
		this.Ref(), js.Pointer(&_ok),
	)
	return UUID{}.FromRef(_ret), _ok
}

// Value returns the value of property "BluetoothRemoteGATTDescriptor.value".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Value() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTDescriptorValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// ReadValue calls the method "BluetoothRemoteGATTDescriptor.readValue".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTDescriptor) ReadValue() (js.Promise[js.DataView], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTDescriptorReadValue(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.DataView]{}.FromRef(_ret), _ok
}

// ReadValueFunc returns the method "BluetoothRemoteGATTDescriptor.readValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTDescriptor) ReadValueFunc() (fn js.Func[func() js.Promise[js.DataView]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTDescriptorReadValueFunc(
			this.Ref(),
		),
	)
}

// WriteValue calls the method "BluetoothRemoteGATTDescriptor.writeValue".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTDescriptor) WriteValue(value BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTDescriptorWriteValue(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteValueFunc returns the method "BluetoothRemoteGATTDescriptor.writeValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTDescriptor) WriteValueFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTDescriptorWriteValueFunc(
			this.Ref(),
		),
	)
}

type OneOf_String_Uint32 struct {
	ref js.Ref
}

func (x OneOf_String_Uint32) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Uint32) Free() {
	x.ref.Free()
}

func (x OneOf_String_Uint32) FromRef(ref js.Ref) OneOf_String_Uint32 {
	return OneOf_String_Uint32{
		ref: ref,
	}
}

func (x OneOf_String_Uint32) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Uint32) Uint32() uint32 {
	return js.Number[uint32]{}.FromRef(x.ref).Get()
}

type BluetoothDescriptorUUID = OneOf_String_Uint32

type BluetoothCharacteristicProperties struct {
	ref js.Ref
}

func (this BluetoothCharacteristicProperties) Once() BluetoothCharacteristicProperties {
	this.Ref().Once()
	return this
}

func (this BluetoothCharacteristicProperties) Ref() js.Ref {
	return this.ref
}

func (this BluetoothCharacteristicProperties) FromRef(ref js.Ref) BluetoothCharacteristicProperties {
	this.ref = ref
	return this
}

func (this BluetoothCharacteristicProperties) Free() {
	this.Ref().Free()
}

// Broadcast returns the value of property "BluetoothCharacteristicProperties.broadcast".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) Broadcast() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesBroadcast(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Read returns the value of property "BluetoothCharacteristicProperties.read".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) Read() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesRead(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// WriteWithoutResponse returns the value of property "BluetoothCharacteristicProperties.writeWithoutResponse".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) WriteWithoutResponse() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesWriteWithoutResponse(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Write returns the value of property "BluetoothCharacteristicProperties.write".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) Write() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesWrite(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Notify returns the value of property "BluetoothCharacteristicProperties.notify".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) Notify() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesNotify(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Indicate returns the value of property "BluetoothCharacteristicProperties.indicate".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) Indicate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesIndicate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AuthenticatedSignedWrites returns the value of property "BluetoothCharacteristicProperties.authenticatedSignedWrites".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) AuthenticatedSignedWrites() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesAuthenticatedSignedWrites(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReliableWrite returns the value of property "BluetoothCharacteristicProperties.reliableWrite".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) ReliableWrite() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesReliableWrite(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// WritableAuxiliaries returns the value of property "BluetoothCharacteristicProperties.writableAuxiliaries".
//
// The returned bool will be false if there is no such property.
func (this BluetoothCharacteristicProperties) WritableAuxiliaries() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothCharacteristicPropertiesWritableAuxiliaries(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type BluetoothRemoteGATTCharacteristic struct {
	EventTarget
}

func (this BluetoothRemoteGATTCharacteristic) Once() BluetoothRemoteGATTCharacteristic {
	this.Ref().Once()
	return this
}

func (this BluetoothRemoteGATTCharacteristic) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BluetoothRemoteGATTCharacteristic) FromRef(ref js.Ref) BluetoothRemoteGATTCharacteristic {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BluetoothRemoteGATTCharacteristic) Free() {
	this.Ref().Free()
}

// Service returns the value of property "BluetoothRemoteGATTCharacteristic.service".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Service() (BluetoothRemoteGATTService, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTCharacteristicService(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothRemoteGATTService{}.FromRef(_ret), _ok
}

// Uuid returns the value of property "BluetoothRemoteGATTCharacteristic.uuid".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Uuid() (UUID, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTCharacteristicUuid(
		this.Ref(), js.Pointer(&_ok),
	)
	return UUID{}.FromRef(_ret), _ok
}

// Properties returns the value of property "BluetoothRemoteGATTCharacteristic.properties".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Properties() (BluetoothCharacteristicProperties, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTCharacteristicProperties(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothCharacteristicProperties{}.FromRef(_ret), _ok
}

// Value returns the value of property "BluetoothRemoteGATTCharacteristic.value".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Value() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTCharacteristicValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// GetDescriptor calls the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptor(descriptor BluetoothDescriptorUUID) (js.Promise[BluetoothRemoteGATTDescriptor], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.Ref(), js.Pointer(&_ok),
		descriptor.Ref(),
	)

	return js.Promise[BluetoothRemoteGATTDescriptor]{}.FromRef(_ret), _ok
}

// GetDescriptorFunc returns the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptorFunc() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[BluetoothRemoteGATTDescriptor]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptorFunc(
			this.Ref(),
		),
	)
}

// GetDescriptors calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors(descriptor BluetoothDescriptorUUID) (js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.Ref(), js.Pointer(&_ok),
		descriptor.Ref(),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]{}.FromRef(_ret), _ok
}

// GetDescriptorsFunc returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptorsFunc() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptorsFunc(
			this.Ref(),
		),
	)
}

// GetDescriptors1 calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors1() (js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]{}.FromRef(_ret), _ok
}

// GetDescriptors1Func returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptors1Func(
			this.Ref(),
		),
	)
}

// ReadValue calls the method "BluetoothRemoteGATTCharacteristic.readValue".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) ReadValue() (js.Promise[js.DataView], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicReadValue(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.DataView]{}.FromRef(_ret), _ok
}

// ReadValueFunc returns the method "BluetoothRemoteGATTCharacteristic.readValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) ReadValueFunc() (fn js.Func[func() js.Promise[js.DataView]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicReadValueFunc(
			this.Ref(),
		),
	)
}

// WriteValue calls the method "BluetoothRemoteGATTCharacteristic.writeValue".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValue(value BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicWriteValue(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteValueFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValueFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueFunc(
			this.Ref(),
		),
	)
}

// WriteValueWithResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithResponse(value BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteValueWithResponseFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithResponseFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueWithResponseFunc(
			this.Ref(),
		),
	)
}

// WriteValueWithoutResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithoutResponse(value BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteValueWithoutResponseFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithoutResponseFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueWithoutResponseFunc(
			this.Ref(),
		),
	)
}

// StartNotifications calls the method "BluetoothRemoteGATTCharacteristic.startNotifications".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) StartNotifications() (js.Promise[BluetoothRemoteGATTCharacteristic], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicStartNotifications(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[BluetoothRemoteGATTCharacteristic]{}.FromRef(_ret), _ok
}

// StartNotificationsFunc returns the method "BluetoothRemoteGATTCharacteristic.startNotifications".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) StartNotificationsFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicStartNotificationsFunc(
			this.Ref(),
		),
	)
}

// StopNotifications calls the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) StopNotifications() (js.Promise[BluetoothRemoteGATTCharacteristic], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTCharacteristicStopNotifications(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[BluetoothRemoteGATTCharacteristic]{}.FromRef(_ret), _ok
}

// StopNotificationsFunc returns the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTCharacteristic) StopNotificationsFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicStopNotificationsFunc(
			this.Ref(),
		),
	)
}

type BluetoothCharacteristicUUID = OneOf_String_Uint32

type BluetoothServiceUUID = OneOf_String_Uint32

type BluetoothRemoteGATTService struct {
	EventTarget
}

func (this BluetoothRemoteGATTService) Once() BluetoothRemoteGATTService {
	this.Ref().Once()
	return this
}

func (this BluetoothRemoteGATTService) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BluetoothRemoteGATTService) FromRef(ref js.Ref) BluetoothRemoteGATTService {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BluetoothRemoteGATTService) Free() {
	this.Ref().Free()
}

// Device returns the value of property "BluetoothRemoteGATTService.device".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTService) Device() (BluetoothDevice, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTServiceDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothDevice{}.FromRef(_ret), _ok
}

// Uuid returns the value of property "BluetoothRemoteGATTService.uuid".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTService) Uuid() (UUID, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTServiceUuid(
		this.Ref(), js.Pointer(&_ok),
	)
	return UUID{}.FromRef(_ret), _ok
}

// IsPrimary returns the value of property "BluetoothRemoteGATTService.isPrimary".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTService) IsPrimary() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTServiceIsPrimary(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetCharacteristic calls the method "BluetoothRemoteGATTService.getCharacteristic".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristic(characteristic BluetoothCharacteristicUUID) (js.Promise[BluetoothRemoteGATTCharacteristic], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetCharacteristic(
		this.Ref(), js.Pointer(&_ok),
		characteristic.Ref(),
	)

	return js.Promise[BluetoothRemoteGATTCharacteristic]{}.FromRef(_ret), _ok
}

// GetCharacteristicFunc returns the method "BluetoothRemoteGATTService.getCharacteristic".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristicFunc() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristicFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristics calls the method "BluetoothRemoteGATTService.getCharacteristics".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristics(characteristic BluetoothCharacteristicUUID) (js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetCharacteristics(
		this.Ref(), js.Pointer(&_ok),
		characteristic.Ref(),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]{}.FromRef(_ret), _ok
}

// GetCharacteristicsFunc returns the method "BluetoothRemoteGATTService.getCharacteristics".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristicsFunc() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristicsFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristics1 calls the method "BluetoothRemoteGATTService.getCharacteristics".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristics1() (js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetCharacteristics1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]{}.FromRef(_ret), _ok
}

// GetCharacteristics1Func returns the method "BluetoothRemoteGATTService.getCharacteristics".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetCharacteristics1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristics1Func(
			this.Ref(),
		),
	)
}

// GetIncludedService calls the method "BluetoothRemoteGATTService.getIncludedService".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedService(service BluetoothServiceUUID) (js.Promise[BluetoothRemoteGATTService], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetIncludedService(
		this.Ref(), js.Pointer(&_ok),
		service.Ref(),
	)

	return js.Promise[BluetoothRemoteGATTService]{}.FromRef(_ret), _ok
}

// GetIncludedServiceFunc returns the method "BluetoothRemoteGATTService.getIncludedService".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedServiceFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServiceFunc(
			this.Ref(),
		),
	)
}

// GetIncludedServices calls the method "BluetoothRemoteGATTService.getIncludedServices".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedServices(service BluetoothServiceUUID) (js.Promise[js.Array[BluetoothRemoteGATTService]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetIncludedServices(
		this.Ref(), js.Pointer(&_ok),
		service.Ref(),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTService]]{}.FromRef(_ret), _ok
}

// GetIncludedServicesFunc returns the method "BluetoothRemoteGATTService.getIncludedServices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedServicesFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServicesFunc(
			this.Ref(),
		),
	)
}

// GetIncludedServices1 calls the method "BluetoothRemoteGATTService.getIncludedServices".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedServices1() (js.Promise[js.Array[BluetoothRemoteGATTService]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServiceGetIncludedServices1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTService]]{}.FromRef(_ret), _ok
}

// GetIncludedServices1Func returns the method "BluetoothRemoteGATTService.getIncludedServices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTService) GetIncludedServices1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServices1Func(
			this.Ref(),
		),
	)
}

type BluetoothRemoteGATTServer struct {
	ref js.Ref
}

func (this BluetoothRemoteGATTServer) Once() BluetoothRemoteGATTServer {
	this.Ref().Once()
	return this
}

func (this BluetoothRemoteGATTServer) Ref() js.Ref {
	return this.ref
}

func (this BluetoothRemoteGATTServer) FromRef(ref js.Ref) BluetoothRemoteGATTServer {
	this.ref = ref
	return this
}

func (this BluetoothRemoteGATTServer) Free() {
	this.Ref().Free()
}

// Device returns the value of property "BluetoothRemoteGATTServer.device".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTServer) Device() (BluetoothDevice, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTServerDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothDevice{}.FromRef(_ret), _ok
}

// Connected returns the value of property "BluetoothRemoteGATTServer.connected".
//
// The returned bool will be false if there is no such property.
func (this BluetoothRemoteGATTServer) Connected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothRemoteGATTServerConnected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Connect calls the method "BluetoothRemoteGATTServer.connect".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTServer) Connect() (js.Promise[BluetoothRemoteGATTServer], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServerConnect(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[BluetoothRemoteGATTServer]{}.FromRef(_ret), _ok
}

// ConnectFunc returns the method "BluetoothRemoteGATTServer.connect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTServer) ConnectFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTServer]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerConnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "BluetoothRemoteGATTServer.disconnect".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTServer) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServerDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "BluetoothRemoteGATTServer.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTServer) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerDisconnectFunc(
			this.Ref(),
		),
	)
}

// GetPrimaryService calls the method "BluetoothRemoteGATTServer.getPrimaryService".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryService(service BluetoothServiceUUID) (js.Promise[BluetoothRemoteGATTService], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServerGetPrimaryService(
		this.Ref(), js.Pointer(&_ok),
		service.Ref(),
	)

	return js.Promise[BluetoothRemoteGATTService]{}.FromRef(_ret), _ok
}

// GetPrimaryServiceFunc returns the method "BluetoothRemoteGATTServer.getPrimaryService".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryServiceFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServiceFunc(
			this.Ref(),
		),
	)
}

// GetPrimaryServices calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryServices(service BluetoothServiceUUID) (js.Promise[js.Array[BluetoothRemoteGATTService]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServerGetPrimaryServices(
		this.Ref(), js.Pointer(&_ok),
		service.Ref(),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTService]]{}.FromRef(_ret), _ok
}

// GetPrimaryServicesFunc returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryServicesFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServicesFunc(
			this.Ref(),
		),
	)
}

// GetPrimaryServices1 calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
//
// The returned bool will be false if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryServices1() (js.Promise[js.Array[BluetoothRemoteGATTService]], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothRemoteGATTServerGetPrimaryServices1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[BluetoothRemoteGATTService]]{}.FromRef(_ret), _ok
}

// GetPrimaryServices1Func returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothRemoteGATTServer) GetPrimaryServices1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServices1Func(
			this.Ref(),
		),
	)
}

type BluetoothDevice struct {
	EventTarget
}

func (this BluetoothDevice) Once() BluetoothDevice {
	this.Ref().Once()
	return this
}

func (this BluetoothDevice) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BluetoothDevice) FromRef(ref js.Ref) BluetoothDevice {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BluetoothDevice) Free() {
	this.Ref().Free()
}

// Id returns the value of property "BluetoothDevice.id".
//
// The returned bool will be false if there is no such property.
func (this BluetoothDevice) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothDeviceId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "BluetoothDevice.name".
//
// The returned bool will be false if there is no such property.
func (this BluetoothDevice) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothDeviceName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Gatt returns the value of property "BluetoothDevice.gatt".
//
// The returned bool will be false if there is no such property.
func (this BluetoothDevice) Gatt() (BluetoothRemoteGATTServer, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothDeviceGatt(
		this.Ref(), js.Pointer(&_ok),
	)
	return BluetoothRemoteGATTServer{}.FromRef(_ret), _ok
}

// WatchingAdvertisements returns the value of property "BluetoothDevice.watchingAdvertisements".
//
// The returned bool will be false if there is no such property.
func (this BluetoothDevice) WatchingAdvertisements() (bool, bool) {
	var _ok bool
	_ret := bindings.GetBluetoothDeviceWatchingAdvertisements(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Forget calls the method "BluetoothDevice.forget".
//
// The returned bool will be false if there is no such method.
func (this BluetoothDevice) Forget() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothDeviceForget(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ForgetFunc returns the method "BluetoothDevice.forget".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothDevice) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceForgetFunc(
			this.Ref(),
		),
	)
}

// WatchAdvertisements calls the method "BluetoothDevice.watchAdvertisements".
//
// The returned bool will be false if there is no such method.
func (this BluetoothDevice) WatchAdvertisements(options WatchAdvertisementsOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothDeviceWatchAdvertisements(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WatchAdvertisementsFunc returns the method "BluetoothDevice.watchAdvertisements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothDevice) WatchAdvertisementsFunc() (fn js.Func[func(options WatchAdvertisementsOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceWatchAdvertisementsFunc(
			this.Ref(),
		),
	)
}

// WatchAdvertisements1 calls the method "BluetoothDevice.watchAdvertisements".
//
// The returned bool will be false if there is no such method.
func (this BluetoothDevice) WatchAdvertisements1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallBluetoothDeviceWatchAdvertisements1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WatchAdvertisements1Func returns the method "BluetoothDevice.watchAdvertisements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BluetoothDevice) WatchAdvertisements1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceWatchAdvertisements1Func(
			this.Ref(),
		),
	)
}

type BluetoothManufacturerDataFilterInit struct {
	// CompanyIdentifier is "BluetoothManufacturerDataFilterInit.companyIdentifier"
	//
	// Required
	CompanyIdentifier uint16
	// DataPrefix is "BluetoothManufacturerDataFilterInit.dataPrefix"
	//
	// Optional
	DataPrefix BufferSource
	// Mask is "BluetoothManufacturerDataFilterInit.mask"
	//
	// Optional
	Mask BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BluetoothManufacturerDataFilterInit with all fields set.
func (p BluetoothManufacturerDataFilterInit) FromRef(ref js.Ref) BluetoothManufacturerDataFilterInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BluetoothManufacturerDataFilterInit in the application heap.
func (p BluetoothManufacturerDataFilterInit) New() js.Ref {
	return bindings.BluetoothManufacturerDataFilterInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p BluetoothManufacturerDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothManufacturerDataFilterInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p BluetoothManufacturerDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothManufacturerDataFilterInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
