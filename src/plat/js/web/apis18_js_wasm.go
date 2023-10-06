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

func NewOffscreenCanvas(width uint64, height uint64) (ret OffscreenCanvas) {
	ret.ref = bindings.NewOffscreenCanvasByOffscreenCanvas(
		float64(width),
		float64(height))
	return
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
// It returns ok=false if there is no such property.
func (this OffscreenCanvas) Width() (ret uint64, ok bool) {
	ok = js.True == bindings.GetOffscreenCanvasWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "OffscreenCanvas.width" to val.
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
// It returns ok=false if there is no such property.
func (this OffscreenCanvas) Height() (ret uint64, ok bool) {
	ok = js.True == bindings.GetOffscreenCanvasHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "OffscreenCanvas.height" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvas) SetHeight(val uint64) bool {
	return js.True == bindings.SetOffscreenCanvasHeight(
		this.Ref(),
		float64(val),
	)
}

// HasGetContext returns true if the method "OffscreenCanvas.getContext" exists.
func (this OffscreenCanvas) HasGetContext() bool {
	return js.True == bindings.HasOffscreenCanvasGetContext(
		this.Ref(),
	)
}

// GetContextFunc returns the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContextFunc() (fn js.Func[func(contextId OffscreenRenderingContextId, options js.Any) OffscreenRenderingContext]) {
	return fn.FromRef(
		bindings.OffscreenCanvasGetContextFunc(
			this.Ref(),
		),
	)
}

// GetContext calls the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContext(contextId OffscreenRenderingContextId, options js.Any) (ret OffscreenRenderingContext) {
	bindings.CallOffscreenCanvasGetContext(
		this.Ref(), js.Pointer(&ret),
		uint32(contextId),
		options.Ref(),
	)

	return
}

// TryGetContext calls the method "OffscreenCanvas.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OffscreenCanvas) TryGetContext(contextId OffscreenRenderingContextId, options js.Any) (ret OffscreenRenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasGetContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(contextId),
		options.Ref(),
	)

	return
}

// HasGetContext1 returns true if the method "OffscreenCanvas.getContext" exists.
func (this OffscreenCanvas) HasGetContext1() bool {
	return js.True == bindings.HasOffscreenCanvasGetContext1(
		this.Ref(),
	)
}

// GetContext1Func returns the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContext1Func() (fn js.Func[func(contextId OffscreenRenderingContextId) OffscreenRenderingContext]) {
	return fn.FromRef(
		bindings.OffscreenCanvasGetContext1Func(
			this.Ref(),
		),
	)
}

// GetContext1 calls the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContext1(contextId OffscreenRenderingContextId) (ret OffscreenRenderingContext) {
	bindings.CallOffscreenCanvasGetContext1(
		this.Ref(), js.Pointer(&ret),
		uint32(contextId),
	)

	return
}

// TryGetContext1 calls the method "OffscreenCanvas.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OffscreenCanvas) TryGetContext1(contextId OffscreenRenderingContextId) (ret OffscreenRenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasGetContext1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(contextId),
	)

	return
}

// HasTransferToImageBitmap returns true if the method "OffscreenCanvas.transferToImageBitmap" exists.
func (this OffscreenCanvas) HasTransferToImageBitmap() bool {
	return js.True == bindings.HasOffscreenCanvasTransferToImageBitmap(
		this.Ref(),
	)
}

// TransferToImageBitmapFunc returns the method "OffscreenCanvas.transferToImageBitmap".
func (this OffscreenCanvas) TransferToImageBitmapFunc() (fn js.Func[func() ImageBitmap]) {
	return fn.FromRef(
		bindings.OffscreenCanvasTransferToImageBitmapFunc(
			this.Ref(),
		),
	)
}

// TransferToImageBitmap calls the method "OffscreenCanvas.transferToImageBitmap".
func (this OffscreenCanvas) TransferToImageBitmap() (ret ImageBitmap) {
	bindings.CallOffscreenCanvasTransferToImageBitmap(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTransferToImageBitmap calls the method "OffscreenCanvas.transferToImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OffscreenCanvas) TryTransferToImageBitmap() (ret ImageBitmap, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasTransferToImageBitmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvertToBlob returns true if the method "OffscreenCanvas.convertToBlob" exists.
func (this OffscreenCanvas) HasConvertToBlob() bool {
	return js.True == bindings.HasOffscreenCanvasConvertToBlob(
		this.Ref(),
	)
}

// ConvertToBlobFunc returns the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlobFunc() (fn js.Func[func(options ImageEncodeOptions) js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.OffscreenCanvasConvertToBlobFunc(
			this.Ref(),
		),
	)
}

// ConvertToBlob calls the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlob(options ImageEncodeOptions) (ret js.Promise[Blob]) {
	bindings.CallOffscreenCanvasConvertToBlob(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryConvertToBlob calls the method "OffscreenCanvas.convertToBlob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OffscreenCanvas) TryConvertToBlob(options ImageEncodeOptions) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasConvertToBlob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasConvertToBlob1 returns true if the method "OffscreenCanvas.convertToBlob" exists.
func (this OffscreenCanvas) HasConvertToBlob1() bool {
	return js.True == bindings.HasOffscreenCanvasConvertToBlob1(
		this.Ref(),
	)
}

// ConvertToBlob1Func returns the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlob1Func() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.OffscreenCanvasConvertToBlob1Func(
			this.Ref(),
		),
	)
}

// ConvertToBlob1 calls the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlob1() (ret js.Promise[Blob]) {
	bindings.CallOffscreenCanvasConvertToBlob1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryConvertToBlob1 calls the method "OffscreenCanvas.convertToBlob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OffscreenCanvas) TryConvertToBlob1() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasConvertToBlob1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Canvas() (ret HTMLCanvasElement, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DCanvas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// GlobalAlpha returns the value of property "CanvasRenderingContext2D.globalAlpha".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) GlobalAlpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DGlobalAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetGlobalAlpha sets the value of property "CanvasRenderingContext2D.globalAlpha" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) GlobalCompositeOperation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DGlobalCompositeOperation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetGlobalCompositeOperation sets the value of property "CanvasRenderingContext2D.globalCompositeOperation" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DImageSmoothingEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingEnabled sets the value of property "CanvasRenderingContext2D.imageSmoothingEnabled" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingQuality() (ret ImageSmoothingQuality, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DImageSmoothingQuality(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingQuality sets the value of property "CanvasRenderingContext2D.imageSmoothingQuality" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) StrokeStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DStrokeStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStrokeStyle sets the value of property "CanvasRenderingContext2D.strokeStyle" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FillStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFillStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFillStyle sets the value of property "CanvasRenderingContext2D.fillStyle" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetX() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowOffsetX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetX sets the value of property "CanvasRenderingContext2D.shadowOffsetX" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetY() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowOffsetY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetY sets the value of property "CanvasRenderingContext2D.shadowOffsetY" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowBlur() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowBlur(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowBlur sets the value of property "CanvasRenderingContext2D.shadowBlur" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowColor sets the value of property "CanvasRenderingContext2D.shadowColor" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Filter() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFilter(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFilter sets the value of property "CanvasRenderingContext2D.filter" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineWidth() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineWidth sets the value of property "CanvasRenderingContext2D.lineWidth" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineCap() (ret CanvasLineCap, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineCap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineCap sets the value of property "CanvasRenderingContext2D.lineCap" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineJoin() (ret CanvasLineJoin, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineJoin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineJoin sets the value of property "CanvasRenderingContext2D.lineJoin" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) MiterLimit() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DMiterLimit(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMiterLimit sets the value of property "CanvasRenderingContext2D.miterLimit" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineDashOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineDashOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineDashOffset sets the value of property "CanvasRenderingContext2D.lineDashOffset" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Font() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFont(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFont sets the value of property "CanvasRenderingContext2D.font" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextAlign() (ret CanvasTextAlign, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTextAlign sets the value of property "CanvasRenderingContext2D.textAlign" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextBaseline() (ret CanvasTextBaseline, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextBaseline(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTextBaseline sets the value of property "CanvasRenderingContext2D.textBaseline" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Direction() (ret CanvasDirection, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDirection sets the value of property "CanvasRenderingContext2D.direction" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LetterSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLetterSpacing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLetterSpacing sets the value of property "CanvasRenderingContext2D.letterSpacing" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontKerning() (ret CanvasFontKerning, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontKerning(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFontKerning sets the value of property "CanvasRenderingContext2D.fontKerning" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontStretch() (ret CanvasFontStretch, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontStretch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFontStretch sets the value of property "CanvasRenderingContext2D.fontStretch" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontVariantCaps() (ret CanvasFontVariantCaps, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontVariantCaps(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFontVariantCaps sets the value of property "CanvasRenderingContext2D.fontVariantCaps" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextRendering() (ret CanvasTextRendering, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextRendering(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTextRendering sets the value of property "CanvasRenderingContext2D.textRendering" to val.
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
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) WordSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DWordSpacing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWordSpacing sets the value of property "CanvasRenderingContext2D.wordSpacing" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetWordSpacing(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DWordSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetContextAttributes returns true if the method "CanvasRenderingContext2D.getContextAttributes" exists.
func (this CanvasRenderingContext2D) HasGetContextAttributes() bool {
	return js.True == bindings.HasCanvasRenderingContext2DGetContextAttributes(
		this.Ref(),
	)
}

// GetContextAttributesFunc returns the method "CanvasRenderingContext2D.getContextAttributes".
func (this CanvasRenderingContext2D) GetContextAttributesFunc() (fn js.Func[func() CanvasRenderingContext2DSettings]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// GetContextAttributes calls the method "CanvasRenderingContext2D.getContextAttributes".
func (this CanvasRenderingContext2D) GetContextAttributes() (ret CanvasRenderingContext2DSettings) {
	bindings.CallCanvasRenderingContext2DGetContextAttributes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "CanvasRenderingContext2D.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryGetContextAttributes() (ret CanvasRenderingContext2DSettings, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetContextAttributes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSave returns true if the method "CanvasRenderingContext2D.save" exists.
func (this CanvasRenderingContext2D) HasSave() bool {
	return js.True == bindings.HasCanvasRenderingContext2DSave(
		this.Ref(),
	)
}

// SaveFunc returns the method "CanvasRenderingContext2D.save".
func (this CanvasRenderingContext2D) SaveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSaveFunc(
			this.Ref(),
		),
	)
}

// Save calls the method "CanvasRenderingContext2D.save".
func (this CanvasRenderingContext2D) Save() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSave(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySave calls the method "CanvasRenderingContext2D.save"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TrySave() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSave(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRestore returns true if the method "CanvasRenderingContext2D.restore" exists.
func (this CanvasRenderingContext2D) HasRestore() bool {
	return js.True == bindings.HasCanvasRenderingContext2DRestore(
		this.Ref(),
	)
}

// RestoreFunc returns the method "CanvasRenderingContext2D.restore".
func (this CanvasRenderingContext2D) RestoreFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRestoreFunc(
			this.Ref(),
		),
	)
}

// Restore calls the method "CanvasRenderingContext2D.restore".
func (this CanvasRenderingContext2D) Restore() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRestore(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRestore calls the method "CanvasRenderingContext2D.restore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryRestore() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRestore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "CanvasRenderingContext2D.reset" exists.
func (this CanvasRenderingContext2D) HasReset() bool {
	return js.True == bindings.HasCanvasRenderingContext2DReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "CanvasRenderingContext2D.reset".
func (this CanvasRenderingContext2D) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "CanvasRenderingContext2D.reset".
func (this CanvasRenderingContext2D) Reset() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "CanvasRenderingContext2D.reset"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsContextLost returns true if the method "CanvasRenderingContext2D.isContextLost" exists.
func (this CanvasRenderingContext2D) HasIsContextLost() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsContextLost(
		this.Ref(),
	)
}

// IsContextLostFunc returns the method "CanvasRenderingContext2D.isContextLost".
func (this CanvasRenderingContext2D) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsContextLostFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "CanvasRenderingContext2D.isContextLost".
func (this CanvasRenderingContext2D) IsContextLost() (ret bool) {
	bindings.CallCanvasRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "CanvasRenderingContext2D.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScale returns true if the method "CanvasRenderingContext2D.scale" exists.
func (this CanvasRenderingContext2D) HasScale() bool {
	return js.True == bindings.HasCanvasRenderingContext2DScale(
		this.Ref(),
	)
}

// ScaleFunc returns the method "CanvasRenderingContext2D.scale".
func (this CanvasRenderingContext2D) ScaleFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScaleFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "CanvasRenderingContext2D.scale".
func (this CanvasRenderingContext2D) Scale(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScale(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScale calls the method "CanvasRenderingContext2D.scale"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryScale(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScale(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotate returns true if the method "CanvasRenderingContext2D.rotate" exists.
func (this CanvasRenderingContext2D) HasRotate() bool {
	return js.True == bindings.HasCanvasRenderingContext2DRotate(
		this.Ref(),
	)
}

// RotateFunc returns the method "CanvasRenderingContext2D.rotate".
func (this CanvasRenderingContext2D) RotateFunc() (fn js.Func[func(angle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRotateFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "CanvasRenderingContext2D.rotate".
func (this CanvasRenderingContext2D) Rotate(angle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRotate(
		this.Ref(), js.Pointer(&ret),
		float64(angle),
	)

	return
}

// TryRotate calls the method "CanvasRenderingContext2D.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryRotate(angle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRotate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(angle),
	)

	return
}

// HasTranslate returns true if the method "CanvasRenderingContext2D.translate" exists.
func (this CanvasRenderingContext2D) HasTranslate() bool {
	return js.True == bindings.HasCanvasRenderingContext2DTranslate(
		this.Ref(),
	)
}

// TranslateFunc returns the method "CanvasRenderingContext2D.translate".
func (this CanvasRenderingContext2D) TranslateFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DTranslateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "CanvasRenderingContext2D.translate".
func (this CanvasRenderingContext2D) Translate(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryTranslate calls the method "CanvasRenderingContext2D.translate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryTranslate(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasTransform returns true if the method "CanvasRenderingContext2D.transform" exists.
func (this CanvasRenderingContext2D) HasTransform() bool {
	return js.True == bindings.HasCanvasRenderingContext2DTransform(
		this.Ref(),
	)
}

// TransformFunc returns the method "CanvasRenderingContext2D.transform".
func (this CanvasRenderingContext2D) TransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DTransformFunc(
			this.Ref(),
		),
	)
}

// Transform calls the method "CanvasRenderingContext2D.transform".
func (this CanvasRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DTransform(
		this.Ref(), js.Pointer(&ret),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// TryTransform calls the method "CanvasRenderingContext2D.transform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasGetTransform returns true if the method "CanvasRenderingContext2D.getTransform" exists.
func (this CanvasRenderingContext2D) HasGetTransform() bool {
	return js.True == bindings.HasCanvasRenderingContext2DGetTransform(
		this.Ref(),
	)
}

// GetTransformFunc returns the method "CanvasRenderingContext2D.getTransform".
func (this CanvasRenderingContext2D) GetTransformFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetTransformFunc(
			this.Ref(),
		),
	)
}

// GetTransform calls the method "CanvasRenderingContext2D.getTransform".
func (this CanvasRenderingContext2D) GetTransform() (ret DOMMatrix) {
	bindings.CallCanvasRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTransform calls the method "CanvasRenderingContext2D.getTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryGetTransform() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetTransform returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasSetTransform() bool {
	return js.True == bindings.HasCanvasRenderingContext2DSetTransform(
		this.Ref(),
	)
}

// SetTransformFunc returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&ret),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// TrySetTransform calls the method "CanvasRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasSetTransform1 returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasSetTransform1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DSetTransform1(
		this.Ref(),
	)
}

// SetTransform1Func returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform1Func() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TrySetTransform1 calls the method "CanvasRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform1(transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasSetTransform2 returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasSetTransform2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DSetTransform2(
		this.Ref(),
	)
}

// SetTransform2Func returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetTransform2Func(
			this.Ref(),
		),
	)
}

// SetTransform2 calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform2() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetTransform2 calls the method "CanvasRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResetTransform returns true if the method "CanvasRenderingContext2D.resetTransform" exists.
func (this CanvasRenderingContext2D) HasResetTransform() bool {
	return js.True == bindings.HasCanvasRenderingContext2DResetTransform(
		this.Ref(),
	)
}

// ResetTransformFunc returns the method "CanvasRenderingContext2D.resetTransform".
func (this CanvasRenderingContext2D) ResetTransformFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DResetTransformFunc(
			this.Ref(),
		),
	)
}

// ResetTransform calls the method "CanvasRenderingContext2D.resetTransform".
func (this CanvasRenderingContext2D) ResetTransform() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResetTransform calls the method "CanvasRenderingContext2D.resetTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryResetTransform() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateLinearGradient returns true if the method "CanvasRenderingContext2D.createLinearGradient" exists.
func (this CanvasRenderingContext2D) HasCreateLinearGradient() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateLinearGradient(
		this.Ref(),
	)
}

// CreateLinearGradientFunc returns the method "CanvasRenderingContext2D.createLinearGradient".
func (this CanvasRenderingContext2D) CreateLinearGradientFunc() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateLinearGradientFunc(
			this.Ref(),
		),
	)
}

// CreateLinearGradient calls the method "CanvasRenderingContext2D.createLinearGradient".
func (this CanvasRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&ret),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// TryCreateLinearGradient calls the method "CanvasRenderingContext2D.createLinearGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// HasCreateRadialGradient returns true if the method "CanvasRenderingContext2D.createRadialGradient" exists.
func (this CanvasRenderingContext2D) HasCreateRadialGradient() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateRadialGradient(
		this.Ref(),
	)
}

// CreateRadialGradientFunc returns the method "CanvasRenderingContext2D.createRadialGradient".
func (this CanvasRenderingContext2D) CreateRadialGradientFunc() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateRadialGradientFunc(
			this.Ref(),
		),
	)
}

// CreateRadialGradient calls the method "CanvasRenderingContext2D.createRadialGradient".
func (this CanvasRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&ret),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return
}

// TryCreateRadialGradient calls the method "CanvasRenderingContext2D.createRadialGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return
}

// HasCreateConicGradient returns true if the method "CanvasRenderingContext2D.createConicGradient" exists.
func (this CanvasRenderingContext2D) HasCreateConicGradient() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateConicGradient(
		this.Ref(),
	)
}

// CreateConicGradientFunc returns the method "CanvasRenderingContext2D.createConicGradient".
func (this CanvasRenderingContext2D) CreateConicGradientFunc() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateConicGradientFunc(
			this.Ref(),
		),
	)
}

// CreateConicGradient calls the method "CanvasRenderingContext2D.createConicGradient".
func (this CanvasRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&ret),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// TryCreateConicGradient calls the method "CanvasRenderingContext2D.createConicGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// HasCreatePattern returns true if the method "CanvasRenderingContext2D.createPattern" exists.
func (this CanvasRenderingContext2D) HasCreatePattern() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreatePattern(
		this.Ref(),
	)
}

// CreatePatternFunc returns the method "CanvasRenderingContext2D.createPattern".
func (this CanvasRenderingContext2D) CreatePatternFunc() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreatePatternFunc(
			this.Ref(),
		),
	)
}

// CreatePattern calls the method "CanvasRenderingContext2D.createPattern".
func (this CanvasRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern) {
	bindings.CallCanvasRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// TryCreatePattern calls the method "CanvasRenderingContext2D.createPattern"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// HasClearRect returns true if the method "CanvasRenderingContext2D.clearRect" exists.
func (this CanvasRenderingContext2D) HasClearRect() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClearRect(
		this.Ref(),
	)
}

// ClearRectFunc returns the method "CanvasRenderingContext2D.clearRect".
func (this CanvasRenderingContext2D) ClearRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClearRectFunc(
			this.Ref(),
		),
	)
}

// ClearRect calls the method "CanvasRenderingContext2D.clearRect".
func (this CanvasRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryClearRect calls the method "CanvasRenderingContext2D.clearRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClearRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFillRect returns true if the method "CanvasRenderingContext2D.fillRect" exists.
func (this CanvasRenderingContext2D) HasFillRect() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFillRect(
		this.Ref(),
	)
}

// FillRectFunc returns the method "CanvasRenderingContext2D.fillRect".
func (this CanvasRenderingContext2D) FillRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillRectFunc(
			this.Ref(),
		),
	)
}

// FillRect calls the method "CanvasRenderingContext2D.fillRect".
func (this CanvasRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryFillRect calls the method "CanvasRenderingContext2D.fillRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFillRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasStrokeRect returns true if the method "CanvasRenderingContext2D.strokeRect" exists.
func (this CanvasRenderingContext2D) HasStrokeRect() bool {
	return js.True == bindings.HasCanvasRenderingContext2DStrokeRect(
		this.Ref(),
	)
}

// StrokeRectFunc returns the method "CanvasRenderingContext2D.strokeRect".
func (this CanvasRenderingContext2D) StrokeRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeRectFunc(
			this.Ref(),
		),
	)
}

// StrokeRect calls the method "CanvasRenderingContext2D.strokeRect".
func (this CanvasRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryStrokeRect calls the method "CanvasRenderingContext2D.strokeRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasBeginPath returns true if the method "CanvasRenderingContext2D.beginPath" exists.
func (this CanvasRenderingContext2D) HasBeginPath() bool {
	return js.True == bindings.HasCanvasRenderingContext2DBeginPath(
		this.Ref(),
	)
}

// BeginPathFunc returns the method "CanvasRenderingContext2D.beginPath".
func (this CanvasRenderingContext2D) BeginPathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DBeginPathFunc(
			this.Ref(),
		),
	)
}

// BeginPath calls the method "CanvasRenderingContext2D.beginPath".
func (this CanvasRenderingContext2D) BeginPath() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBeginPath calls the method "CanvasRenderingContext2D.beginPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryBeginPath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFill returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFill() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFill(
		this.Ref(),
	)
}

// FillFunc returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) FillFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillFunc(
			this.Ref(),
		),
	)
}

// Fill calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill(
		this.Ref(), js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryFill calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFill(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasFill1 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFill1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFill1(
		this.Ref(),
	)
}

// Fill1Func returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill1Func(
			this.Ref(),
		),
	)
}

// Fill1 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill1() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFill1 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFill1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFill2 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFill2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFill2(
		this.Ref(),
	)
}

// Fill2Func returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill2Func(
			this.Ref(),
		),
	)
}

// Fill2 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryFill2 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFill2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasFill3 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFill3() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFill3(
		this.Ref(),
	)
}

// Fill3Func returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFill3Func(
			this.Ref(),
		),
	)
}

// Fill3 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill3(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryFill3 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFill3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasStroke returns true if the method "CanvasRenderingContext2D.stroke" exists.
func (this CanvasRenderingContext2D) HasStroke() bool {
	return js.True == bindings.HasCanvasRenderingContext2DStroke(
		this.Ref(),
	)
}

// StrokeFunc returns the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) StrokeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeFunc(
			this.Ref(),
		),
	)
}

// Stroke calls the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) Stroke() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStroke(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStroke calls the method "CanvasRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryStroke() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStroke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStroke1 returns true if the method "CanvasRenderingContext2D.stroke" exists.
func (this CanvasRenderingContext2D) HasStroke1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DStroke1(
		this.Ref(),
	)
}

// Stroke1Func returns the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) Stroke1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStroke1Func(
			this.Ref(),
		),
	)
}

// Stroke1 calls the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) Stroke1(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryStroke1 calls the method "CanvasRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryStroke1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasClip returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasClip() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClip(
		this.Ref(),
	)
}

// ClipFunc returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) ClipFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClipFunc(
			this.Ref(),
		),
	)
}

// Clip calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip(
		this.Ref(), js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryClip calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClip(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasClip1 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasClip1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClip1(
		this.Ref(),
	)
}

// Clip1Func returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip1Func(
			this.Ref(),
		),
	)
}

// Clip1 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip1() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClip1 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClip1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClip2 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasClip2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClip2(
		this.Ref(),
	)
}

// Clip2Func returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip2Func(
			this.Ref(),
		),
	)
}

// Clip2 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryClip2 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClip2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasClip3 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasClip3() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClip3(
		this.Ref(),
	)
}

// Clip3Func returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClip3Func(
			this.Ref(),
		),
	)
}

// Clip3 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip3(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryClip3 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClip3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasIsPointInPath returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasIsPointInPath() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInPath(
		this.Ref(),
	)
}

// IsPointInPathFunc returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPathFunc() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPathFunc(
			this.Ref(),
		),
	)
}

// IsPointInPath calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasIsPointInPath1 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasIsPointInPath1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInPath1(
		this.Ref(),
	)
}

// IsPointInPath1Func returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath1Func() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath1Func(
			this.Ref(),
		),
	)
}

// IsPointInPath1 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath1(x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath1 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath1(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInPath2 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasIsPointInPath2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInPath2(
		this.Ref(),
	)
}

// IsPointInPath2Func returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath2Func() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath2Func(
			this.Ref(),
		),
	)
}

// IsPointInPath2 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath2 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasIsPointInPath3 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasIsPointInPath3() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInPath3(
		this.Ref(),
	)
}

// IsPointInPath3Func returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath3Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInPath3Func(
			this.Ref(),
		),
	)
}

// IsPointInPath3 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath3 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath3(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInStroke returns true if the method "CanvasRenderingContext2D.isPointInStroke" exists.
func (this CanvasRenderingContext2D) HasIsPointInStroke() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInStroke(
		this.Ref(),
	)
}

// IsPointInStrokeFunc returns the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStrokeFunc() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStroke(x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke calls the method "CanvasRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInStroke(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInStroke1 returns true if the method "CanvasRenderingContext2D.isPointInStroke" exists.
func (this CanvasRenderingContext2D) HasIsPointInStroke1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DIsPointInStroke1(
		this.Ref(),
	)
}

// IsPointInStroke1Func returns the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStroke1Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke1 calls the method "CanvasRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInStroke1(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasDrawFocusIfNeeded returns true if the method "CanvasRenderingContext2D.drawFocusIfNeeded" exists.
func (this CanvasRenderingContext2D) HasDrawFocusIfNeeded() bool {
	return js.True == bindings.HasCanvasRenderingContext2DDrawFocusIfNeeded(
		this.Ref(),
	)
}

// DrawFocusIfNeededFunc returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeededFunc() (fn js.Func[func(element Element)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawFocusIfNeededFunc(
			this.Ref(),
		),
	)
}

// DrawFocusIfNeeded calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeeded(element Element) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryDrawFocusIfNeeded calls the method "CanvasRenderingContext2D.drawFocusIfNeeded"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawFocusIfNeeded(element Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawFocusIfNeeded(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasDrawFocusIfNeeded1 returns true if the method "CanvasRenderingContext2D.drawFocusIfNeeded" exists.
func (this CanvasRenderingContext2D) HasDrawFocusIfNeeded1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.Ref(),
	)
}

// DrawFocusIfNeeded1Func returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeeded1Func() (fn js.Func[func(path Path2D, element Element)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawFocusIfNeeded1Func(
			this.Ref(),
		),
	)
}

// DrawFocusIfNeeded1 calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeeded1(path Path2D, element Element) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		element.Ref(),
	)

	return
}

// TryDrawFocusIfNeeded1 calls the method "CanvasRenderingContext2D.drawFocusIfNeeded"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawFocusIfNeeded1(path Path2D, element Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		element.Ref(),
	)

	return
}

// HasScrollPathIntoView returns true if the method "CanvasRenderingContext2D.scrollPathIntoView" exists.
func (this CanvasRenderingContext2D) HasScrollPathIntoView() bool {
	return js.True == bindings.HasCanvasRenderingContext2DScrollPathIntoView(
		this.Ref(),
	)
}

// ScrollPathIntoViewFunc returns the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoViewFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScrollPathIntoViewFunc(
			this.Ref(),
		),
	)
}

// ScrollPathIntoView calls the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoView() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScrollPathIntoView(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScrollPathIntoView calls the method "CanvasRenderingContext2D.scrollPathIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryScrollPathIntoView() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScrollPathIntoView(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScrollPathIntoView1 returns true if the method "CanvasRenderingContext2D.scrollPathIntoView" exists.
func (this CanvasRenderingContext2D) HasScrollPathIntoView1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DScrollPathIntoView1(
		this.Ref(),
	)
}

// ScrollPathIntoView1Func returns the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoView1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DScrollPathIntoView1Func(
			this.Ref(),
		),
	)
}

// ScrollPathIntoView1 calls the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoView1(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScrollPathIntoView1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryScrollPathIntoView1 calls the method "CanvasRenderingContext2D.scrollPathIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryScrollPathIntoView1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScrollPathIntoView1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFillText returns true if the method "CanvasRenderingContext2D.fillText" exists.
func (this CanvasRenderingContext2D) HasFillText() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFillText(
		this.Ref(),
	)
}

// FillTextFunc returns the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillTextFunc(
			this.Ref(),
		),
	)
}

// FillText calls the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillText(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// TryFillText calls the method "CanvasRenderingContext2D.fillText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFillText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// HasFillText1 returns true if the method "CanvasRenderingContext2D.fillText" exists.
func (this CanvasRenderingContext2D) HasFillText1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DFillText1(
		this.Ref(),
	)
}

// FillText1Func returns the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DFillText1Func(
			this.Ref(),
		),
	)
}

// FillText1 calls the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillText1(text js.String, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillText1(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryFillText1 calls the method "CanvasRenderingContext2D.fillText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryFillText1(text js.String, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasStrokeText returns true if the method "CanvasRenderingContext2D.strokeText" exists.
func (this CanvasRenderingContext2D) HasStrokeText() bool {
	return js.True == bindings.HasCanvasRenderingContext2DStrokeText(
		this.Ref(),
	)
}

// StrokeTextFunc returns the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeTextFunc() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeTextFunc(
			this.Ref(),
		),
	)
}

// StrokeText calls the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeText(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// TryStrokeText calls the method "CanvasRenderingContext2D.strokeText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// HasStrokeText1 returns true if the method "CanvasRenderingContext2D.strokeText" exists.
func (this CanvasRenderingContext2D) HasStrokeText1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DStrokeText1(
		this.Ref(),
	)
}

// StrokeText1Func returns the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeText1Func() (fn js.Func[func(text js.String, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DStrokeText1Func(
			this.Ref(),
		),
	)
}

// StrokeText1 calls the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeText1(text js.String, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeText1(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryStrokeText1 calls the method "CanvasRenderingContext2D.strokeText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeText1(text js.String, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasMeasureText returns true if the method "CanvasRenderingContext2D.measureText" exists.
func (this CanvasRenderingContext2D) HasMeasureText() bool {
	return js.True == bindings.HasCanvasRenderingContext2DMeasureText(
		this.Ref(),
	)
}

// MeasureTextFunc returns the method "CanvasRenderingContext2D.measureText".
func (this CanvasRenderingContext2D) MeasureTextFunc() (fn js.Func[func(text js.String) TextMetrics]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DMeasureTextFunc(
			this.Ref(),
		),
	)
}

// MeasureText calls the method "CanvasRenderingContext2D.measureText".
func (this CanvasRenderingContext2D) MeasureText(text js.String) (ret TextMetrics) {
	bindings.CallCanvasRenderingContext2DMeasureText(
		this.Ref(), js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryMeasureText calls the method "CanvasRenderingContext2D.measureText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryMeasureText(text js.String) (ret TextMetrics, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DMeasureText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasDrawImage returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasDrawImage() bool {
	return js.True == bindings.HasCanvasRenderingContext2DDrawImage(
		this.Ref(),
	)
}

// DrawImageFunc returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImageFunc() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImageFunc(
			this.Ref(),
		),
	)
}

// DrawImage calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// TryDrawImage calls the method "CanvasRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// HasDrawImage1 returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasDrawImage1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DDrawImage1(
		this.Ref(),
	)
}

// DrawImage1Func returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage1Func() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImage1Func(
			this.Ref(),
		),
	)
}

// DrawImage1 calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// TryDrawImage1 calls the method "CanvasRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// HasDrawImage2 returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasDrawImage2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DDrawImage2(
		this.Ref(),
	)
}

// DrawImage2Func returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage2Func() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DDrawImage2Func(
			this.Ref(),
		),
	)
}

// DrawImage2 calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryDrawImage2 calls the method "CanvasRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCreateImageData returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasCreateImageData() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateImageData(
		this.Ref(),
	)
}

// CreateImageDataFunc returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageDataFunc() (fn js.Func[func(sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageDataFunc(
			this.Ref(),
		),
	)
}

// CreateImageData calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData(sw int32, sh int32, settings ImageDataSettings) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData(
		this.Ref(), js.Pointer(&ret),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// TryCreateImageData calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData(sw int32, sh int32, settings ImageDataSettings) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// HasCreateImageData1 returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasCreateImageData1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateImageData1(
		this.Ref(),
	)
}

// CreateImageData1Func returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData1Func() (fn js.Func[func(sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageData1Func(
			this.Ref(),
		),
	)
}

// CreateImageData1 calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData1(sw int32, sh int32) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData1(
		this.Ref(), js.Pointer(&ret),
		int32(sw),
		int32(sh),
	)

	return
}

// TryCreateImageData1 calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData1(sw int32, sh int32) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(sw),
		int32(sh),
	)

	return
}

// HasCreateImageData2 returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasCreateImageData2() bool {
	return js.True == bindings.HasCanvasRenderingContext2DCreateImageData2(
		this.Ref(),
	)
}

// CreateImageData2Func returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData2Func() (fn js.Func[func(imagedata ImageData) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DCreateImageData2Func(
			this.Ref(),
		),
	)
}

// CreateImageData2 calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData2(imagedata ImageData) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData2(
		this.Ref(), js.Pointer(&ret),
		imagedata.Ref(),
	)

	return
}

// TryCreateImageData2 calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData2(imagedata ImageData) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		imagedata.Ref(),
	)

	return
}

// HasGetImageData returns true if the method "CanvasRenderingContext2D.getImageData" exists.
func (this CanvasRenderingContext2D) HasGetImageData() bool {
	return js.True == bindings.HasCanvasRenderingContext2DGetImageData(
		this.Ref(),
	)
}

// GetImageDataFunc returns the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageDataFunc() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetImageDataFunc(
			this.Ref(),
		),
	)
}

// GetImageData calls the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DGetImageData(
		this.Ref(), js.Pointer(&ret),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// TryGetImageData calls the method "CanvasRenderingContext2D.getImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryGetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetImageData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// HasGetImageData1 returns true if the method "CanvasRenderingContext2D.getImageData" exists.
func (this CanvasRenderingContext2D) HasGetImageData1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DGetImageData1(
		this.Ref(),
	)
}

// GetImageData1Func returns the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageData1Func() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32) ImageData]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetImageData1Func(
			this.Ref(),
		),
	)
}

// GetImageData1 calls the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageData1(sx int32, sy int32, sw int32, sh int32) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DGetImageData1(
		this.Ref(), js.Pointer(&ret),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// TryGetImageData1 calls the method "CanvasRenderingContext2D.getImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryGetImageData1(sx int32, sy int32, sw int32, sh int32) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetImageData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasPutImageData returns true if the method "CanvasRenderingContext2D.putImageData" exists.
func (this CanvasRenderingContext2D) HasPutImageData() bool {
	return js.True == bindings.HasCanvasRenderingContext2DPutImageData(
		this.Ref(),
	)
}

// PutImageDataFunc returns the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageDataFunc() (fn js.Func[func(imagedata ImageData, dx int32, dy int32)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DPutImageDataFunc(
			this.Ref(),
		),
	)
}

// PutImageData calls the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageData(imagedata ImageData, dx int32, dy int32) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DPutImageData(
		this.Ref(), js.Pointer(&ret),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	return
}

// TryPutImageData calls the method "CanvasRenderingContext2D.putImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryPutImageData(imagedata ImageData, dx int32, dy int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DPutImageData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	return
}

// HasPutImageData1 returns true if the method "CanvasRenderingContext2D.putImageData" exists.
func (this CanvasRenderingContext2D) HasPutImageData1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DPutImageData1(
		this.Ref(),
	)
}

// PutImageData1Func returns the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageData1Func() (fn js.Func[func(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DPutImageData1Func(
			this.Ref(),
		),
	)
}

// PutImageData1 calls the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DPutImageData1(
		this.Ref(), js.Pointer(&ret),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
		int32(dirtyX),
		int32(dirtyY),
		int32(dirtyWidth),
		int32(dirtyHeight),
	)

	return
}

// TryPutImageData1 calls the method "CanvasRenderingContext2D.putImageData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryPutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DPutImageData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
		int32(dirtyX),
		int32(dirtyY),
		int32(dirtyWidth),
		int32(dirtyHeight),
	)

	return
}

// HasSetLineDash returns true if the method "CanvasRenderingContext2D.setLineDash" exists.
func (this CanvasRenderingContext2D) HasSetLineDash() bool {
	return js.True == bindings.HasCanvasRenderingContext2DSetLineDash(
		this.Ref(),
	)
}

// SetLineDashFunc returns the method "CanvasRenderingContext2D.setLineDash".
func (this CanvasRenderingContext2D) SetLineDashFunc() (fn js.Func[func(segments js.Array[float64])]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DSetLineDashFunc(
			this.Ref(),
		),
	)
}

// SetLineDash calls the method "CanvasRenderingContext2D.setLineDash".
func (this CanvasRenderingContext2D) SetLineDash(segments js.Array[float64]) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&ret),
		segments.Ref(),
	)

	return
}

// TrySetLineDash calls the method "CanvasRenderingContext2D.setLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TrySetLineDash(segments js.Array[float64]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		segments.Ref(),
	)

	return
}

// HasGetLineDash returns true if the method "CanvasRenderingContext2D.getLineDash" exists.
func (this CanvasRenderingContext2D) HasGetLineDash() bool {
	return js.True == bindings.HasCanvasRenderingContext2DGetLineDash(
		this.Ref(),
	)
}

// GetLineDashFunc returns the method "CanvasRenderingContext2D.getLineDash".
func (this CanvasRenderingContext2D) GetLineDashFunc() (fn js.Func[func() js.Array[float64]]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DGetLineDashFunc(
			this.Ref(),
		),
	)
}

// GetLineDash calls the method "CanvasRenderingContext2D.getLineDash".
func (this CanvasRenderingContext2D) GetLineDash() (ret js.Array[float64]) {
	bindings.CallCanvasRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetLineDash calls the method "CanvasRenderingContext2D.getLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryGetLineDash() (ret js.Array[float64], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClosePath returns true if the method "CanvasRenderingContext2D.closePath" exists.
func (this CanvasRenderingContext2D) HasClosePath() bool {
	return js.True == bindings.HasCanvasRenderingContext2DClosePath(
		this.Ref(),
	)
}

// ClosePathFunc returns the method "CanvasRenderingContext2D.closePath".
func (this CanvasRenderingContext2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DClosePathFunc(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "CanvasRenderingContext2D.closePath".
func (this CanvasRenderingContext2D) ClosePath() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClosePath calls the method "CanvasRenderingContext2D.closePath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryClosePath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMoveTo returns true if the method "CanvasRenderingContext2D.moveTo" exists.
func (this CanvasRenderingContext2D) HasMoveTo() bool {
	return js.True == bindings.HasCanvasRenderingContext2DMoveTo(
		this.Ref(),
	)
}

// MoveToFunc returns the method "CanvasRenderingContext2D.moveTo".
func (this CanvasRenderingContext2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DMoveToFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "CanvasRenderingContext2D.moveTo".
func (this CanvasRenderingContext2D) MoveTo(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryMoveTo calls the method "CanvasRenderingContext2D.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryMoveTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasLineTo returns true if the method "CanvasRenderingContext2D.lineTo" exists.
func (this CanvasRenderingContext2D) HasLineTo() bool {
	return js.True == bindings.HasCanvasRenderingContext2DLineTo(
		this.Ref(),
	)
}

// LineToFunc returns the method "CanvasRenderingContext2D.lineTo".
func (this CanvasRenderingContext2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DLineToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "CanvasRenderingContext2D.lineTo".
func (this CanvasRenderingContext2D) LineTo(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryLineTo calls the method "CanvasRenderingContext2D.lineTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryLineTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasQuadraticCurveTo returns true if the method "CanvasRenderingContext2D.quadraticCurveTo" exists.
func (this CanvasRenderingContext2D) HasQuadraticCurveTo() bool {
	return js.True == bindings.HasCanvasRenderingContext2DQuadraticCurveTo(
		this.Ref(),
	)
}

// QuadraticCurveToFunc returns the method "CanvasRenderingContext2D.quadraticCurveTo".
func (this CanvasRenderingContext2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "CanvasRenderingContext2D.quadraticCurveTo".
func (this CanvasRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// TryQuadraticCurveTo calls the method "CanvasRenderingContext2D.quadraticCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryQuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// HasBezierCurveTo returns true if the method "CanvasRenderingContext2D.bezierCurveTo" exists.
func (this CanvasRenderingContext2D) HasBezierCurveTo() bool {
	return js.True == bindings.HasCanvasRenderingContext2DBezierCurveTo(
		this.Ref(),
	)
}

// BezierCurveToFunc returns the method "CanvasRenderingContext2D.bezierCurveTo".
func (this CanvasRenderingContext2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "CanvasRenderingContext2D.bezierCurveTo".
func (this CanvasRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// TryBezierCurveTo calls the method "CanvasRenderingContext2D.bezierCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryBezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// HasArcTo returns true if the method "CanvasRenderingContext2D.arcTo" exists.
func (this CanvasRenderingContext2D) HasArcTo() bool {
	return js.True == bindings.HasCanvasRenderingContext2DArcTo(
		this.Ref(),
	)
}

// ArcToFunc returns the method "CanvasRenderingContext2D.arcTo".
func (this CanvasRenderingContext2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArcToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "CanvasRenderingContext2D.arcTo".
func (this CanvasRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&ret),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// TryArcTo calls the method "CanvasRenderingContext2D.arcTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// HasRect returns true if the method "CanvasRenderingContext2D.rect" exists.
func (this CanvasRenderingContext2D) HasRect() bool {
	return js.True == bindings.HasCanvasRenderingContext2DRect(
		this.Ref(),
	)
}

// RectFunc returns the method "CanvasRenderingContext2D.rect".
func (this CanvasRenderingContext2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRectFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "CanvasRenderingContext2D.rect".
func (this CanvasRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRect calls the method "CanvasRenderingContext2D.rect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasRoundRect returns true if the method "CanvasRenderingContext2D.roundRect" exists.
func (this CanvasRenderingContext2D) HasRoundRect() bool {
	return js.True == bindings.HasCanvasRenderingContext2DRoundRect(
		this.Ref(),
	)
}

// RoundRectFunc returns the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// TryRoundRect calls the method "CanvasRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryRoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// HasRoundRect1 returns true if the method "CanvasRenderingContext2D.roundRect" exists.
func (this CanvasRenderingContext2D) HasRoundRect1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DRoundRect1(
		this.Ref(),
	)
}

// RoundRect1Func returns the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRoundRect1 calls the method "CanvasRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryRoundRect1(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasArc returns true if the method "CanvasRenderingContext2D.arc" exists.
func (this CanvasRenderingContext2D) HasArc() bool {
	return js.True == bindings.HasCanvasRenderingContext2DArc(
		this.Ref(),
	)
}

// ArcFunc returns the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc calls the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArc(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryArc calls the method "CanvasRenderingContext2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryArc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasArc1 returns true if the method "CanvasRenderingContext2D.arc" exists.
func (this CanvasRenderingContext2D) HasArc1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DArc1(
		this.Ref(),
	)
}

// Arc1Func returns the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DArc1Func(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArc1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryArc1 calls the method "CanvasRenderingContext2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryArc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArc1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasEllipse returns true if the method "CanvasRenderingContext2D.ellipse" exists.
func (this CanvasRenderingContext2D) HasEllipse() bool {
	return js.True == bindings.HasCanvasRenderingContext2DEllipse(
		this.Ref(),
	)
}

// EllipseFunc returns the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryEllipse calls the method "CanvasRenderingContext2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryEllipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasEllipse1 returns true if the method "CanvasRenderingContext2D.ellipse" exists.
func (this CanvasRenderingContext2D) HasEllipse1() bool {
	return js.True == bindings.HasCanvasRenderingContext2DEllipse1(
		this.Ref(),
	)
}

// Ellipse1Func returns the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.CanvasRenderingContext2DEllipse1Func(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryEllipse1 calls the method "CanvasRenderingContext2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this CanvasRenderingContext2D) TryEllipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
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

func NewHTMLCanvasElement() (ret HTMLCanvasElement) {
	ret.ref = bindings.NewHTMLCanvasElementByHTMLCanvasElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLCanvasElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCanvasElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLCanvasElement.width" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLCanvasElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCanvasElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLCanvasElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLCanvasElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLCanvasElementHeight(
		this.Ref(),
		uint32(val),
	)
}

// HasGetContext returns true if the method "HTMLCanvasElement.getContext" exists.
func (this HTMLCanvasElement) HasGetContext() bool {
	return js.True == bindings.HasHTMLCanvasElementGetContext(
		this.Ref(),
	)
}

// GetContextFunc returns the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContextFunc() (fn js.Func[func(contextId js.String, options js.Any) RenderingContext]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementGetContextFunc(
			this.Ref(),
		),
	)
}

// GetContext calls the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContext(contextId js.String, options js.Any) (ret RenderingContext) {
	bindings.CallHTMLCanvasElementGetContext(
		this.Ref(), js.Pointer(&ret),
		contextId.Ref(),
		options.Ref(),
	)

	return
}

// TryGetContext calls the method "HTMLCanvasElement.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryGetContext(contextId js.String, options js.Any) (ret RenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementGetContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextId.Ref(),
		options.Ref(),
	)

	return
}

// HasGetContext1 returns true if the method "HTMLCanvasElement.getContext" exists.
func (this HTMLCanvasElement) HasGetContext1() bool {
	return js.True == bindings.HasHTMLCanvasElementGetContext1(
		this.Ref(),
	)
}

// GetContext1Func returns the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContext1Func() (fn js.Func[func(contextId js.String) RenderingContext]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementGetContext1Func(
			this.Ref(),
		),
	)
}

// GetContext1 calls the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContext1(contextId js.String) (ret RenderingContext) {
	bindings.CallHTMLCanvasElementGetContext1(
		this.Ref(), js.Pointer(&ret),
		contextId.Ref(),
	)

	return
}

// TryGetContext1 calls the method "HTMLCanvasElement.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryGetContext1(contextId js.String) (ret RenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementGetContext1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextId.Ref(),
	)

	return
}

// HasToDataURL returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasToDataURL() bool {
	return js.True == bindings.HasHTMLCanvasElementToDataURL(
		this.Ref(),
	)
}

// ToDataURLFunc returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURLFunc() (fn js.Func[func(typ js.String, quality js.Any) js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURLFunc(
			this.Ref(),
		),
	)
}

// ToDataURL calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL(typ js.String, quality js.Any) (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// TryToDataURL calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL(typ js.String, quality js.Any) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// HasToDataURL1 returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasToDataURL1() bool {
	return js.True == bindings.HasHTMLCanvasElementToDataURL1(
		this.Ref(),
	)
}

// ToDataURL1Func returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL1Func() (fn js.Func[func(typ js.String) js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURL1Func(
			this.Ref(),
		),
	)
}

// ToDataURL1 calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL1(typ js.String) (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL1(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryToDataURL1 calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL1(typ js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasToDataURL2 returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasToDataURL2() bool {
	return js.True == bindings.HasHTMLCanvasElementToDataURL2(
		this.Ref(),
	)
}

// ToDataURL2Func returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL2Func() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToDataURL2Func(
			this.Ref(),
		),
	)
}

// ToDataURL2 calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL2() (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToDataURL2 calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL2() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToBlob returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasToBlob() bool {
	return js.True == bindings.HasHTMLCanvasElementToBlob(
		this.Ref(),
	)
}

// ToBlobFunc returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlobFunc() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String, quality js.Any)]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlobFunc(
			this.Ref(),
		),
	)
}

// ToBlob calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob(callback js.Func[func(blob Blob)], typ js.String, quality js.Any) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// TryToBlob calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToBlob(callback js.Func[func(blob Blob)], typ js.String, quality js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// HasToBlob1 returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasToBlob1() bool {
	return js.True == bindings.HasHTMLCanvasElementToBlob1(
		this.Ref(),
	)
}

// ToBlob1Func returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob1Func() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String)]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlob1Func(
			this.Ref(),
		),
	)
}

// ToBlob1 calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob1(callback js.Func[func(blob Blob)], typ js.String) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob1(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
		typ.Ref(),
	)

	return
}

// TryToBlob1 calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToBlob1(callback js.Func[func(blob Blob)], typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		typ.Ref(),
	)

	return
}

// HasToBlob2 returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasToBlob2() bool {
	return js.True == bindings.HasHTMLCanvasElementToBlob2(
		this.Ref(),
	)
}

// ToBlob2Func returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob2Func() (fn js.Func[func(callback js.Func[func(blob Blob)])]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementToBlob2Func(
			this.Ref(),
		),
	)
}

// ToBlob2 calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob2(callback js.Func[func(blob Blob)]) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob2(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryToBlob2 calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryToBlob2(callback js.Func[func(blob Blob)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasTransferControlToOffscreen returns true if the method "HTMLCanvasElement.transferControlToOffscreen" exists.
func (this HTMLCanvasElement) HasTransferControlToOffscreen() bool {
	return js.True == bindings.HasHTMLCanvasElementTransferControlToOffscreen(
		this.Ref(),
	)
}

// TransferControlToOffscreenFunc returns the method "HTMLCanvasElement.transferControlToOffscreen".
func (this HTMLCanvasElement) TransferControlToOffscreenFunc() (fn js.Func[func() OffscreenCanvas]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementTransferControlToOffscreenFunc(
			this.Ref(),
		),
	)
}

// TransferControlToOffscreen calls the method "HTMLCanvasElement.transferControlToOffscreen".
func (this HTMLCanvasElement) TransferControlToOffscreen() (ret OffscreenCanvas) {
	bindings.CallHTMLCanvasElementTransferControlToOffscreen(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTransferControlToOffscreen calls the method "HTMLCanvasElement.transferControlToOffscreen"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryTransferControlToOffscreen() (ret OffscreenCanvas, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementTransferControlToOffscreen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCaptureStream returns true if the method "HTMLCanvasElement.captureStream" exists.
func (this HTMLCanvasElement) HasCaptureStream() bool {
	return js.True == bindings.HasHTMLCanvasElementCaptureStream(
		this.Ref(),
	)
}

// CaptureStreamFunc returns the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStreamFunc() (fn js.Func[func(frameRequestRate float64) MediaStream]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementCaptureStreamFunc(
			this.Ref(),
		),
	)
}

// CaptureStream calls the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStream(frameRequestRate float64) (ret MediaStream) {
	bindings.CallHTMLCanvasElementCaptureStream(
		this.Ref(), js.Pointer(&ret),
		float64(frameRequestRate),
	)

	return
}

// TryCaptureStream calls the method "HTMLCanvasElement.captureStream"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryCaptureStream(frameRequestRate float64) (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementCaptureStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(frameRequestRate),
	)

	return
}

// HasCaptureStream1 returns true if the method "HTMLCanvasElement.captureStream" exists.
func (this HTMLCanvasElement) HasCaptureStream1() bool {
	return js.True == bindings.HasHTMLCanvasElementCaptureStream1(
		this.Ref(),
	)
}

// CaptureStream1Func returns the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStream1Func() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.HTMLCanvasElementCaptureStream1Func(
			this.Ref(),
		),
	)
}

// CaptureStream1 calls the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStream1() (ret MediaStream) {
	bindings.CallHTMLCanvasElementCaptureStream1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCaptureStream1 calls the method "HTMLCanvasElement.captureStream"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLCanvasElement) TryCaptureStream1() (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementCaptureStream1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewBarcodeDetector(barcodeDetectorOptions BarcodeDetectorOptions) (ret BarcodeDetector) {
	ret.ref = bindings.NewBarcodeDetectorByBarcodeDetector(
		js.Pointer(&barcodeDetectorOptions))
	return
}

func NewBarcodeDetectorByBarcodeDetector1() (ret BarcodeDetector) {
	ret.ref = bindings.NewBarcodeDetectorByBarcodeDetector1()
	return
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

// HasGetSupportedFormats returns true if the staticmethod "BarcodeDetector.getSupportedFormats" exists.
func (this BarcodeDetector) HasGetSupportedFormats() bool {
	return js.True == bindings.HasBarcodeDetectorGetSupportedFormats(
		this.Ref(),
	)
}

// GetSupportedFormatsFunc returns the staticmethod "BarcodeDetector.getSupportedFormats".
func (this BarcodeDetector) GetSupportedFormatsFunc() (fn js.Func[func() js.Promise[js.Array[BarcodeFormat]]]) {
	return fn.FromRef(
		bindings.BarcodeDetectorGetSupportedFormatsFunc(
			this.Ref(),
		),
	)
}

// GetSupportedFormats calls the staticmethod "BarcodeDetector.getSupportedFormats".
func (this BarcodeDetector) GetSupportedFormats() (ret js.Promise[js.Array[BarcodeFormat]]) {
	bindings.CallBarcodeDetectorGetSupportedFormats(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedFormats calls the staticmethod "BarcodeDetector.getSupportedFormats"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BarcodeDetector) TryGetSupportedFormats() (ret js.Promise[js.Array[BarcodeFormat]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBarcodeDetectorGetSupportedFormats(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDetect returns true if the method "BarcodeDetector.detect" exists.
func (this BarcodeDetector) HasDetect() bool {
	return js.True == bindings.HasBarcodeDetectorDetect(
		this.Ref(),
	)
}

// DetectFunc returns the method "BarcodeDetector.detect".
func (this BarcodeDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedBarcode]]]) {
	return fn.FromRef(
		bindings.BarcodeDetectorDetectFunc(
			this.Ref(),
		),
	)
}

// Detect calls the method "BarcodeDetector.detect".
func (this BarcodeDetector) Detect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedBarcode]]) {
	bindings.CallBarcodeDetectorDetect(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryDetect calls the method "BarcodeDetector.detect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BarcodeDetector) TryDetect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedBarcode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBarcodeDetectorDetect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this BatteryManager) Charging() (ret bool, ok bool) {
	ok = js.True == bindings.GetBatteryManagerCharging(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChargingTime returns the value of property "BatteryManager.chargingTime".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) ChargingTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerChargingTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DischargingTime returns the value of property "BatteryManager.dischargingTime".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) DischargingTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerDischargingTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Level returns the value of property "BatteryManager.level".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) Level() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerLevel(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewBeforeInstallPromptEvent(typ js.String, eventInitDict EventInit) (ret BeforeInstallPromptEvent) {
	ret.ref = bindings.NewBeforeInstallPromptEventByBeforeInstallPromptEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewBeforeInstallPromptEventByBeforeInstallPromptEvent1(typ js.String) (ret BeforeInstallPromptEvent) {
	ret.ref = bindings.NewBeforeInstallPromptEventByBeforeInstallPromptEvent1(
		typ.Ref())
	return
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

// HasPrompt returns true if the method "BeforeInstallPromptEvent.prompt" exists.
func (this BeforeInstallPromptEvent) HasPrompt() bool {
	return js.True == bindings.HasBeforeInstallPromptEventPrompt(
		this.Ref(),
	)
}

// PromptFunc returns the method "BeforeInstallPromptEvent.prompt".
func (this BeforeInstallPromptEvent) PromptFunc() (fn js.Func[func() js.Promise[PromptResponseObject]]) {
	return fn.FromRef(
		bindings.BeforeInstallPromptEventPromptFunc(
			this.Ref(),
		),
	)
}

// Prompt calls the method "BeforeInstallPromptEvent.prompt".
func (this BeforeInstallPromptEvent) Prompt() (ret js.Promise[PromptResponseObject]) {
	bindings.CallBeforeInstallPromptEventPrompt(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPrompt calls the method "BeforeInstallPromptEvent.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BeforeInstallPromptEvent) TryPrompt() (ret js.Promise[PromptResponseObject], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBeforeInstallPromptEventPrompt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewBeforeUnloadEvent(typ js.String, eventInitDict EventInit) (ret BeforeUnloadEvent) {
	ret.ref = bindings.NewBeforeUnloadEventByBeforeUnloadEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewBeforeUnloadEventByBeforeUnloadEvent1(typ js.String) (ret BeforeUnloadEvent) {
	ret.ref = bindings.NewBeforeUnloadEventByBeforeUnloadEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this BeforeUnloadEvent) ReturnValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBeforeUnloadEventReturnValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReturnValue sets the value of property "BeforeUnloadEvent.returnValue" to val.
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

func NewBlobEvent(typ js.String, eventInitDict BlobEventInit) (ret BlobEvent) {
	ret.ref = bindings.NewBlobEventByBlobEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this BlobEvent) Data() (ret Blob, ok bool) {
	ok = js.True == bindings.GetBlobEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timecode returns the value of property "BlobEvent.timecode".
//
// It returns ok=false if there is no such property.
func (this BlobEvent) Timecode() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetBlobEventTimecode(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Characteristic() (ret BluetoothRemoteGATTCharacteristic, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorCharacteristic(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTDescriptor.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorUuid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "BluetoothRemoteGATTDescriptor.value".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Value() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasReadValue returns true if the method "BluetoothRemoteGATTDescriptor.readValue" exists.
func (this BluetoothRemoteGATTDescriptor) HasReadValue() bool {
	return js.True == bindings.HasBluetoothRemoteGATTDescriptorReadValue(
		this.Ref(),
	)
}

// ReadValueFunc returns the method "BluetoothRemoteGATTDescriptor.readValue".
func (this BluetoothRemoteGATTDescriptor) ReadValueFunc() (fn js.Func[func() js.Promise[js.DataView]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTDescriptorReadValueFunc(
			this.Ref(),
		),
	)
}

// ReadValue calls the method "BluetoothRemoteGATTDescriptor.readValue".
func (this BluetoothRemoteGATTDescriptor) ReadValue() (ret js.Promise[js.DataView]) {
	bindings.CallBluetoothRemoteGATTDescriptorReadValue(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReadValue calls the method "BluetoothRemoteGATTDescriptor.readValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTDescriptor) TryReadValue() (ret js.Promise[js.DataView], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTDescriptorReadValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWriteValue returns true if the method "BluetoothRemoteGATTDescriptor.writeValue" exists.
func (this BluetoothRemoteGATTDescriptor) HasWriteValue() bool {
	return js.True == bindings.HasBluetoothRemoteGATTDescriptorWriteValue(
		this.Ref(),
	)
}

// WriteValueFunc returns the method "BluetoothRemoteGATTDescriptor.writeValue".
func (this BluetoothRemoteGATTDescriptor) WriteValueFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTDescriptorWriteValueFunc(
			this.Ref(),
		),
	)
}

// WriteValue calls the method "BluetoothRemoteGATTDescriptor.writeValue".
func (this BluetoothRemoteGATTDescriptor) WriteValue(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTDescriptorWriteValue(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValue calls the method "BluetoothRemoteGATTDescriptor.writeValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTDescriptor) TryWriteValue(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTDescriptorWriteValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Broadcast() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesBroadcast(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Read returns the value of property "BluetoothCharacteristicProperties.read".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Read() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesRead(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WriteWithoutResponse returns the value of property "BluetoothCharacteristicProperties.writeWithoutResponse".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) WriteWithoutResponse() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWriteWithoutResponse(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Write returns the value of property "BluetoothCharacteristicProperties.write".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Write() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWrite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Notify returns the value of property "BluetoothCharacteristicProperties.notify".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Notify() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesNotify(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Indicate returns the value of property "BluetoothCharacteristicProperties.indicate".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Indicate() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesIndicate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AuthenticatedSignedWrites returns the value of property "BluetoothCharacteristicProperties.authenticatedSignedWrites".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) AuthenticatedSignedWrites() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesAuthenticatedSignedWrites(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReliableWrite returns the value of property "BluetoothCharacteristicProperties.reliableWrite".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) ReliableWrite() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesReliableWrite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WritableAuxiliaries returns the value of property "BluetoothCharacteristicProperties.writableAuxiliaries".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) WritableAuxiliaries() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWritableAuxiliaries(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Service() (ret BluetoothRemoteGATTService, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicService(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTCharacteristic.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicUuid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Properties returns the value of property "BluetoothRemoteGATTCharacteristic.properties".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Properties() (ret BluetoothCharacteristicProperties, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicProperties(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "BluetoothRemoteGATTCharacteristic.value".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Value() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetDescriptor returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptor" exists.
func (this BluetoothRemoteGATTCharacteristic) HasGetDescriptor() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.Ref(),
	)
}

// GetDescriptorFunc returns the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptorFunc() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[BluetoothRemoteGATTDescriptor]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptorFunc(
			this.Ref(),
		),
	)
}

// GetDescriptor calls the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptor(descriptor BluetoothDescriptorUUID) (ret js.Promise[BluetoothRemoteGATTDescriptor]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.Ref(), js.Pointer(&ret),
		descriptor.Ref(),
	)

	return
}

// TryGetDescriptor calls the method "BluetoothRemoteGATTCharacteristic.getDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptor(descriptor BluetoothDescriptorUUID) (ret js.Promise[BluetoothRemoteGATTDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		descriptor.Ref(),
	)

	return
}

// HasGetDescriptors returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptors" exists.
func (this BluetoothRemoteGATTCharacteristic) HasGetDescriptors() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.Ref(),
	)
}

// GetDescriptorsFunc returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptorsFunc() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptorsFunc(
			this.Ref(),
		),
	)
}

// GetDescriptors calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors(descriptor BluetoothDescriptorUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.Ref(), js.Pointer(&ret),
		descriptor.Ref(),
	)

	return
}

// TryGetDescriptors calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptors(descriptor BluetoothDescriptorUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		descriptor.Ref(),
	)

	return
}

// HasGetDescriptors1 returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptors" exists.
func (this BluetoothRemoteGATTCharacteristic) HasGetDescriptors1() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.Ref(),
	)
}

// GetDescriptors1Func returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicGetDescriptors1Func(
			this.Ref(),
		),
	)
}

// GetDescriptors1 calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors1() (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDescriptors1 calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptors1() (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReadValue returns true if the method "BluetoothRemoteGATTCharacteristic.readValue" exists.
func (this BluetoothRemoteGATTCharacteristic) HasReadValue() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicReadValue(
		this.Ref(),
	)
}

// ReadValueFunc returns the method "BluetoothRemoteGATTCharacteristic.readValue".
func (this BluetoothRemoteGATTCharacteristic) ReadValueFunc() (fn js.Func[func() js.Promise[js.DataView]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicReadValueFunc(
			this.Ref(),
		),
	)
}

// ReadValue calls the method "BluetoothRemoteGATTCharacteristic.readValue".
func (this BluetoothRemoteGATTCharacteristic) ReadValue() (ret js.Promise[js.DataView]) {
	bindings.CallBluetoothRemoteGATTCharacteristicReadValue(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReadValue calls the method "BluetoothRemoteGATTCharacteristic.readValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryReadValue() (ret js.Promise[js.DataView], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicReadValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWriteValue returns true if the method "BluetoothRemoteGATTCharacteristic.writeValue" exists.
func (this BluetoothRemoteGATTCharacteristic) HasWriteValue() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicWriteValue(
		this.Ref(),
	)
}

// WriteValueFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValue".
func (this BluetoothRemoteGATTCharacteristic) WriteValueFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueFunc(
			this.Ref(),
		),
	)
}

// WriteValue calls the method "BluetoothRemoteGATTCharacteristic.writeValue".
func (this BluetoothRemoteGATTCharacteristic) WriteValue(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValue(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValue calls the method "BluetoothRemoteGATTCharacteristic.writeValue"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValue(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasWriteValueWithResponse returns true if the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse" exists.
func (this BluetoothRemoteGATTCharacteristic) HasWriteValueWithResponse() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.Ref(),
	)
}

// WriteValueWithResponseFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithResponseFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueWithResponseFunc(
			this.Ref(),
		),
	)
}

// WriteValueWithResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithResponse(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValueWithResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValueWithResponse(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasWriteValueWithoutResponse returns true if the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse" exists.
func (this BluetoothRemoteGATTCharacteristic) HasWriteValueWithoutResponse() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.Ref(),
	)
}

// WriteValueWithoutResponseFunc returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithoutResponseFunc() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicWriteValueWithoutResponseFunc(
			this.Ref(),
		),
	)
}

// WriteValueWithoutResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithoutResponse(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValueWithoutResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValueWithoutResponse(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasStartNotifications returns true if the method "BluetoothRemoteGATTCharacteristic.startNotifications" exists.
func (this BluetoothRemoteGATTCharacteristic) HasStartNotifications() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicStartNotifications(
		this.Ref(),
	)
}

// StartNotificationsFunc returns the method "BluetoothRemoteGATTCharacteristic.startNotifications".
func (this BluetoothRemoteGATTCharacteristic) StartNotificationsFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicStartNotificationsFunc(
			this.Ref(),
		),
	)
}

// StartNotifications calls the method "BluetoothRemoteGATTCharacteristic.startNotifications".
func (this BluetoothRemoteGATTCharacteristic) StartNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTCharacteristicStartNotifications(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStartNotifications calls the method "BluetoothRemoteGATTCharacteristic.startNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryStartNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicStartNotifications(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStopNotifications returns true if the method "BluetoothRemoteGATTCharacteristic.stopNotifications" exists.
func (this BluetoothRemoteGATTCharacteristic) HasStopNotifications() bool {
	return js.True == bindings.HasBluetoothRemoteGATTCharacteristicStopNotifications(
		this.Ref(),
	)
}

// StopNotificationsFunc returns the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
func (this BluetoothRemoteGATTCharacteristic) StopNotificationsFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTCharacteristicStopNotificationsFunc(
			this.Ref(),
		),
	)
}

// StopNotifications calls the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
func (this BluetoothRemoteGATTCharacteristic) StopNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTCharacteristicStopNotifications(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStopNotifications calls the method "BluetoothRemoteGATTCharacteristic.stopNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryStopNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicStopNotifications(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceDevice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTService.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceUuid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "BluetoothRemoteGATTService.isPrimary".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceIsPrimary(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetCharacteristic returns true if the method "BluetoothRemoteGATTService.getCharacteristic" exists.
func (this BluetoothRemoteGATTService) HasGetCharacteristic() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetCharacteristic(
		this.Ref(),
	)
}

// GetCharacteristicFunc returns the method "BluetoothRemoteGATTService.getCharacteristic".
func (this BluetoothRemoteGATTService) GetCharacteristicFunc() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristicFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristic calls the method "BluetoothRemoteGATTService.getCharacteristic".
func (this BluetoothRemoteGATTService) GetCharacteristic(characteristic BluetoothCharacteristicUUID) (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristic(
		this.Ref(), js.Pointer(&ret),
		characteristic.Ref(),
	)

	return
}

// TryGetCharacteristic calls the method "BluetoothRemoteGATTService.getCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristic(characteristic BluetoothCharacteristicUUID) (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristic(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		characteristic.Ref(),
	)

	return
}

// HasGetCharacteristics returns true if the method "BluetoothRemoteGATTService.getCharacteristics" exists.
func (this BluetoothRemoteGATTService) HasGetCharacteristics() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetCharacteristics(
		this.Ref(),
	)
}

// GetCharacteristicsFunc returns the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristicsFunc() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristicsFunc(
			this.Ref(),
		),
	)
}

// GetCharacteristics calls the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristics(characteristic BluetoothCharacteristicUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristics(
		this.Ref(), js.Pointer(&ret),
		characteristic.Ref(),
	)

	return
}

// TryGetCharacteristics calls the method "BluetoothRemoteGATTService.getCharacteristics"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristics(characteristic BluetoothCharacteristicUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristics(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		characteristic.Ref(),
	)

	return
}

// HasGetCharacteristics1 returns true if the method "BluetoothRemoteGATTService.getCharacteristics" exists.
func (this BluetoothRemoteGATTService) HasGetCharacteristics1() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetCharacteristics1(
		this.Ref(),
	)
}

// GetCharacteristics1Func returns the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristics1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetCharacteristics1Func(
			this.Ref(),
		),
	)
}

// GetCharacteristics1 calls the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristics1() (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristics1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCharacteristics1 calls the method "BluetoothRemoteGATTService.getCharacteristics"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristics1() (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristics1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetIncludedService returns true if the method "BluetoothRemoteGATTService.getIncludedService" exists.
func (this BluetoothRemoteGATTService) HasGetIncludedService() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetIncludedService(
		this.Ref(),
	)
}

// GetIncludedServiceFunc returns the method "BluetoothRemoteGATTService.getIncludedService".
func (this BluetoothRemoteGATTService) GetIncludedServiceFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServiceFunc(
			this.Ref(),
		),
	)
}

// GetIncludedService calls the method "BluetoothRemoteGATTService.getIncludedService".
func (this BluetoothRemoteGATTService) GetIncludedService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedService(
		this.Ref(), js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetIncludedService calls the method "BluetoothRemoteGATTService.getIncludedService"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedService(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasGetIncludedServices returns true if the method "BluetoothRemoteGATTService.getIncludedServices" exists.
func (this BluetoothRemoteGATTService) HasGetIncludedServices() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetIncludedServices(
		this.Ref(),
	)
}

// GetIncludedServicesFunc returns the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServicesFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServicesFunc(
			this.Ref(),
		),
	)
}

// GetIncludedServices calls the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedServices(
		this.Ref(), js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetIncludedServices calls the method "BluetoothRemoteGATTService.getIncludedServices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedServices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasGetIncludedServices1 returns true if the method "BluetoothRemoteGATTService.getIncludedServices" exists.
func (this BluetoothRemoteGATTService) HasGetIncludedServices1() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServiceGetIncludedServices1(
		this.Ref(),
	)
}

// GetIncludedServices1Func returns the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServices1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServiceGetIncludedServices1Func(
			this.Ref(),
		),
	)
}

// GetIncludedServices1 calls the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedServices1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetIncludedServices1 calls the method "BluetoothRemoteGATTService.getIncludedServices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedServices1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTServer) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServerDevice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Connected returns the value of property "BluetoothRemoteGATTServer.connected".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTServer) Connected() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServerConnected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasConnect returns true if the method "BluetoothRemoteGATTServer.connect" exists.
func (this BluetoothRemoteGATTServer) HasConnect() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServerConnect(
		this.Ref(),
	)
}

// ConnectFunc returns the method "BluetoothRemoteGATTServer.connect".
func (this BluetoothRemoteGATTServer) ConnectFunc() (fn js.Func[func() js.Promise[BluetoothRemoteGATTServer]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerConnectFunc(
			this.Ref(),
		),
	)
}

// Connect calls the method "BluetoothRemoteGATTServer.connect".
func (this BluetoothRemoteGATTServer) Connect() (ret js.Promise[BluetoothRemoteGATTServer]) {
	bindings.CallBluetoothRemoteGATTServerConnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryConnect calls the method "BluetoothRemoteGATTServer.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTServer) TryConnect() (ret js.Promise[BluetoothRemoteGATTServer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerConnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDisconnect returns true if the method "BluetoothRemoteGATTServer.disconnect" exists.
func (this BluetoothRemoteGATTServer) HasDisconnect() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServerDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "BluetoothRemoteGATTServer.disconnect".
func (this BluetoothRemoteGATTServer) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "BluetoothRemoteGATTServer.disconnect".
func (this BluetoothRemoteGATTServer) Disconnect() (ret js.Void) {
	bindings.CallBluetoothRemoteGATTServerDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "BluetoothRemoteGATTServer.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTServer) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPrimaryService returns true if the method "BluetoothRemoteGATTServer.getPrimaryService" exists.
func (this BluetoothRemoteGATTServer) HasGetPrimaryService() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServerGetPrimaryService(
		this.Ref(),
	)
}

// GetPrimaryServiceFunc returns the method "BluetoothRemoteGATTServer.getPrimaryService".
func (this BluetoothRemoteGATTServer) GetPrimaryServiceFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServiceFunc(
			this.Ref(),
		),
	)
}

// GetPrimaryService calls the method "BluetoothRemoteGATTServer.getPrimaryService".
func (this BluetoothRemoteGATTServer) GetPrimaryService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryService(
		this.Ref(), js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetPrimaryService calls the method "BluetoothRemoteGATTServer.getPrimaryService"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryService(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasGetPrimaryServices returns true if the method "BluetoothRemoteGATTServer.getPrimaryServices" exists.
func (this BluetoothRemoteGATTServer) HasGetPrimaryServices() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServerGetPrimaryServices(
		this.Ref(),
	)
}

// GetPrimaryServicesFunc returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServicesFunc() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServicesFunc(
			this.Ref(),
		),
	)
}

// GetPrimaryServices calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryServices(
		this.Ref(), js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetPrimaryServices calls the method "BluetoothRemoteGATTServer.getPrimaryServices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryServices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasGetPrimaryServices1 returns true if the method "BluetoothRemoteGATTServer.getPrimaryServices" exists.
func (this BluetoothRemoteGATTServer) HasGetPrimaryServices1() bool {
	return js.True == bindings.HasBluetoothRemoteGATTServerGetPrimaryServices1(
		this.Ref(),
	)
}

// GetPrimaryServices1Func returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServices1Func() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	return fn.FromRef(
		bindings.BluetoothRemoteGATTServerGetPrimaryServices1Func(
			this.Ref(),
		),
	)
}

// GetPrimaryServices1 calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryServices1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPrimaryServices1 calls the method "BluetoothRemoteGATTServer.getPrimaryServices"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryServices1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "BluetoothDevice.name".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Gatt returns the value of property "BluetoothDevice.gatt".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Gatt() (ret BluetoothRemoteGATTServer, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceGatt(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WatchingAdvertisements returns the value of property "BluetoothDevice.watchingAdvertisements".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) WatchingAdvertisements() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceWatchingAdvertisements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasForget returns true if the method "BluetoothDevice.forget" exists.
func (this BluetoothDevice) HasForget() bool {
	return js.True == bindings.HasBluetoothDeviceForget(
		this.Ref(),
	)
}

// ForgetFunc returns the method "BluetoothDevice.forget".
func (this BluetoothDevice) ForgetFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceForgetFunc(
			this.Ref(),
		),
	)
}

// Forget calls the method "BluetoothDevice.forget".
func (this BluetoothDevice) Forget() (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceForget(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryForget calls the method "BluetoothDevice.forget"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothDevice) TryForget() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceForget(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWatchAdvertisements returns true if the method "BluetoothDevice.watchAdvertisements" exists.
func (this BluetoothDevice) HasWatchAdvertisements() bool {
	return js.True == bindings.HasBluetoothDeviceWatchAdvertisements(
		this.Ref(),
	)
}

// WatchAdvertisementsFunc returns the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisementsFunc() (fn js.Func[func(options WatchAdvertisementsOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceWatchAdvertisementsFunc(
			this.Ref(),
		),
	)
}

// WatchAdvertisements calls the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisements(options WatchAdvertisementsOptions) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceWatchAdvertisements(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryWatchAdvertisements calls the method "BluetoothDevice.watchAdvertisements"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothDevice) TryWatchAdvertisements(options WatchAdvertisementsOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceWatchAdvertisements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasWatchAdvertisements1 returns true if the method "BluetoothDevice.watchAdvertisements" exists.
func (this BluetoothDevice) HasWatchAdvertisements1() bool {
	return js.True == bindings.HasBluetoothDeviceWatchAdvertisements1(
		this.Ref(),
	)
}

// WatchAdvertisements1Func returns the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisements1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.BluetoothDeviceWatchAdvertisements1Func(
			this.Ref(),
		),
	)
}

// WatchAdvertisements1 calls the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisements1() (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceWatchAdvertisements1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryWatchAdvertisements1 calls the method "BluetoothDevice.watchAdvertisements"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BluetoothDevice) TryWatchAdvertisements1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceWatchAdvertisements1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
