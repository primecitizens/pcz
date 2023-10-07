// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *ImageEncodeOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageEncodeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageEncodeOptions) Update(ref js.Ref) {
	bindings.ImageEncodeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageEncodeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Type.Ref(),
	)
	p.Type = p.Type.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "OffscreenCanvas.width".
//
// It returns ok=false if there is no such property.
func (this OffscreenCanvas) Width() (ret uint64, ok bool) {
	ok = js.True == bindings.GetOffscreenCanvasWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "OffscreenCanvas.width" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvas) SetWidth(val uint64) bool {
	return js.True == bindings.SetOffscreenCanvasWidth(
		this.ref,
		float64(val),
	)
}

// Height returns the value of property "OffscreenCanvas.height".
//
// It returns ok=false if there is no such property.
func (this OffscreenCanvas) Height() (ret uint64, ok bool) {
	ok = js.True == bindings.GetOffscreenCanvasHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "OffscreenCanvas.height" to val.
//
// It returns false if the property cannot be set.
func (this OffscreenCanvas) SetHeight(val uint64) bool {
	return js.True == bindings.SetOffscreenCanvasHeight(
		this.ref,
		float64(val),
	)
}

// HasFuncGetContext returns true if the method "OffscreenCanvas.getContext" exists.
func (this OffscreenCanvas) HasFuncGetContext() bool {
	return js.True == bindings.HasFuncOffscreenCanvasGetContext(
		this.ref,
	)
}

// FuncGetContext returns the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) FuncGetContext() (fn js.Func[func(contextId OffscreenRenderingContextId, options js.Any) OffscreenRenderingContext]) {
	bindings.FuncOffscreenCanvasGetContext(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContext calls the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContext(contextId OffscreenRenderingContextId, options js.Any) (ret OffscreenRenderingContext) {
	bindings.CallOffscreenCanvasGetContext(
		this.ref, js.Pointer(&ret),
		uint32(contextId),
		options.Ref(),
	)

	return
}

// TryGetContext calls the method "OffscreenCanvas.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OffscreenCanvas) TryGetContext(contextId OffscreenRenderingContextId, options js.Any) (ret OffscreenRenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasGetContext(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(contextId),
		options.Ref(),
	)

	return
}

// HasFuncGetContext1 returns true if the method "OffscreenCanvas.getContext" exists.
func (this OffscreenCanvas) HasFuncGetContext1() bool {
	return js.True == bindings.HasFuncOffscreenCanvasGetContext1(
		this.ref,
	)
}

// FuncGetContext1 returns the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) FuncGetContext1() (fn js.Func[func(contextId OffscreenRenderingContextId) OffscreenRenderingContext]) {
	bindings.FuncOffscreenCanvasGetContext1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContext1 calls the method "OffscreenCanvas.getContext".
func (this OffscreenCanvas) GetContext1(contextId OffscreenRenderingContextId) (ret OffscreenRenderingContext) {
	bindings.CallOffscreenCanvasGetContext1(
		this.ref, js.Pointer(&ret),
		uint32(contextId),
	)

	return
}

// TryGetContext1 calls the method "OffscreenCanvas.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OffscreenCanvas) TryGetContext1(contextId OffscreenRenderingContextId) (ret OffscreenRenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasGetContext1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(contextId),
	)

	return
}

// HasFuncTransferToImageBitmap returns true if the method "OffscreenCanvas.transferToImageBitmap" exists.
func (this OffscreenCanvas) HasFuncTransferToImageBitmap() bool {
	return js.True == bindings.HasFuncOffscreenCanvasTransferToImageBitmap(
		this.ref,
	)
}

// FuncTransferToImageBitmap returns the method "OffscreenCanvas.transferToImageBitmap".
func (this OffscreenCanvas) FuncTransferToImageBitmap() (fn js.Func[func() ImageBitmap]) {
	bindings.FuncOffscreenCanvasTransferToImageBitmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransferToImageBitmap calls the method "OffscreenCanvas.transferToImageBitmap".
func (this OffscreenCanvas) TransferToImageBitmap() (ret ImageBitmap) {
	bindings.CallOffscreenCanvasTransferToImageBitmap(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTransferToImageBitmap calls the method "OffscreenCanvas.transferToImageBitmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OffscreenCanvas) TryTransferToImageBitmap() (ret ImageBitmap, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasTransferToImageBitmap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvertToBlob returns true if the method "OffscreenCanvas.convertToBlob" exists.
func (this OffscreenCanvas) HasFuncConvertToBlob() bool {
	return js.True == bindings.HasFuncOffscreenCanvasConvertToBlob(
		this.ref,
	)
}

// FuncConvertToBlob returns the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) FuncConvertToBlob() (fn js.Func[func(options ImageEncodeOptions) js.Promise[Blob]]) {
	bindings.FuncOffscreenCanvasConvertToBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertToBlob calls the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlob(options ImageEncodeOptions) (ret js.Promise[Blob]) {
	bindings.CallOffscreenCanvasConvertToBlob(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryConvertToBlob calls the method "OffscreenCanvas.convertToBlob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OffscreenCanvas) TryConvertToBlob(options ImageEncodeOptions) (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasConvertToBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvertToBlob1 returns true if the method "OffscreenCanvas.convertToBlob" exists.
func (this OffscreenCanvas) HasFuncConvertToBlob1() bool {
	return js.True == bindings.HasFuncOffscreenCanvasConvertToBlob1(
		this.ref,
	)
}

// FuncConvertToBlob1 returns the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) FuncConvertToBlob1() (fn js.Func[func() js.Promise[Blob]]) {
	bindings.FuncOffscreenCanvasConvertToBlob1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvertToBlob1 calls the method "OffscreenCanvas.convertToBlob".
func (this OffscreenCanvas) ConvertToBlob1() (ret js.Promise[Blob]) {
	bindings.CallOffscreenCanvasConvertToBlob1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryConvertToBlob1 calls the method "OffscreenCanvas.convertToBlob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this OffscreenCanvas) TryConvertToBlob1() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffscreenCanvasConvertToBlob1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Canvas returns the value of property "CanvasRenderingContext2D.canvas".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Canvas() (ret HTMLCanvasElement, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DCanvas(
		this.ref, js.Pointer(&ret),
	)
	return
}

// GlobalAlpha returns the value of property "CanvasRenderingContext2D.globalAlpha".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) GlobalAlpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DGlobalAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetGlobalAlpha sets the value of property "CanvasRenderingContext2D.globalAlpha" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetGlobalAlpha(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DGlobalAlpha(
		this.ref,
		float64(val),
	)
}

// GlobalCompositeOperation returns the value of property "CanvasRenderingContext2D.globalCompositeOperation".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) GlobalCompositeOperation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DGlobalCompositeOperation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetGlobalCompositeOperation sets the value of property "CanvasRenderingContext2D.globalCompositeOperation" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetGlobalCompositeOperation(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DGlobalCompositeOperation(
		this.ref,
		val.Ref(),
	)
}

// ImageSmoothingEnabled returns the value of property "CanvasRenderingContext2D.imageSmoothingEnabled".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DImageSmoothingEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingEnabled sets the value of property "CanvasRenderingContext2D.imageSmoothingEnabled" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetImageSmoothingEnabled(val bool) bool {
	return js.True == bindings.SetCanvasRenderingContext2DImageSmoothingEnabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ImageSmoothingQuality returns the value of property "CanvasRenderingContext2D.imageSmoothingQuality".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ImageSmoothingQuality() (ret ImageSmoothingQuality, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DImageSmoothingQuality(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingQuality sets the value of property "CanvasRenderingContext2D.imageSmoothingQuality" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) bool {
	return js.True == bindings.SetCanvasRenderingContext2DImageSmoothingQuality(
		this.ref,
		uint32(val),
	)
}

// StrokeStyle returns the value of property "CanvasRenderingContext2D.strokeStyle".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) StrokeStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DStrokeStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStrokeStyle sets the value of property "CanvasRenderingContext2D.strokeStyle" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetStrokeStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetCanvasRenderingContext2DStrokeStyle(
		this.ref,
		val.Ref(),
	)
}

// FillStyle returns the value of property "CanvasRenderingContext2D.fillStyle".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FillStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFillStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFillStyle sets the value of property "CanvasRenderingContext2D.fillStyle" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFillStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFillStyle(
		this.ref,
		val.Ref(),
	)
}

// ShadowOffsetX returns the value of property "CanvasRenderingContext2D.shadowOffsetX".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetX() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowOffsetX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetX sets the value of property "CanvasRenderingContext2D.shadowOffsetX" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowOffsetX(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowOffsetX(
		this.ref,
		float64(val),
	)
}

// ShadowOffsetY returns the value of property "CanvasRenderingContext2D.shadowOffsetY".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowOffsetY() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowOffsetY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetY sets the value of property "CanvasRenderingContext2D.shadowOffsetY" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowOffsetY(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowOffsetY(
		this.ref,
		float64(val),
	)
}

// ShadowBlur returns the value of property "CanvasRenderingContext2D.shadowBlur".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowBlur() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowBlur(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShadowBlur sets the value of property "CanvasRenderingContext2D.shadowBlur" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowBlur(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowBlur(
		this.ref,
		float64(val),
	)
}

// ShadowColor returns the value of property "CanvasRenderingContext2D.shadowColor".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) ShadowColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DShadowColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShadowColor sets the value of property "CanvasRenderingContext2D.shadowColor" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetShadowColor(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DShadowColor(
		this.ref,
		val.Ref(),
	)
}

// Filter returns the value of property "CanvasRenderingContext2D.filter".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Filter() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFilter(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFilter sets the value of property "CanvasRenderingContext2D.filter" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFilter(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFilter(
		this.ref,
		val.Ref(),
	)
}

// LineWidth returns the value of property "CanvasRenderingContext2D.lineWidth".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineWidth() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineWidth sets the value of property "CanvasRenderingContext2D.lineWidth" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineWidth(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineWidth(
		this.ref,
		float64(val),
	)
}

// LineCap returns the value of property "CanvasRenderingContext2D.lineCap".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineCap() (ret CanvasLineCap, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineCap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineCap sets the value of property "CanvasRenderingContext2D.lineCap" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineCap(val CanvasLineCap) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineCap(
		this.ref,
		uint32(val),
	)
}

// LineJoin returns the value of property "CanvasRenderingContext2D.lineJoin".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineJoin() (ret CanvasLineJoin, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineJoin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineJoin sets the value of property "CanvasRenderingContext2D.lineJoin" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineJoin(val CanvasLineJoin) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineJoin(
		this.ref,
		uint32(val),
	)
}

// MiterLimit returns the value of property "CanvasRenderingContext2D.miterLimit".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) MiterLimit() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DMiterLimit(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMiterLimit sets the value of property "CanvasRenderingContext2D.miterLimit" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetMiterLimit(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DMiterLimit(
		this.ref,
		float64(val),
	)
}

// LineDashOffset returns the value of property "CanvasRenderingContext2D.lineDashOffset".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LineDashOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLineDashOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineDashOffset sets the value of property "CanvasRenderingContext2D.lineDashOffset" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLineDashOffset(val float64) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLineDashOffset(
		this.ref,
		float64(val),
	)
}

// Font returns the value of property "CanvasRenderingContext2D.font".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Font() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFont(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFont sets the value of property "CanvasRenderingContext2D.font" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFont(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFont(
		this.ref,
		val.Ref(),
	)
}

// TextAlign returns the value of property "CanvasRenderingContext2D.textAlign".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextAlign() (ret CanvasTextAlign, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextAlign sets the value of property "CanvasRenderingContext2D.textAlign" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextAlign(val CanvasTextAlign) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextAlign(
		this.ref,
		uint32(val),
	)
}

// TextBaseline returns the value of property "CanvasRenderingContext2D.textBaseline".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextBaseline() (ret CanvasTextBaseline, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextBaseline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextBaseline sets the value of property "CanvasRenderingContext2D.textBaseline" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextBaseline(val CanvasTextBaseline) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextBaseline(
		this.ref,
		uint32(val),
	)
}

// Direction returns the value of property "CanvasRenderingContext2D.direction".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) Direction() (ret CanvasDirection, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDirection sets the value of property "CanvasRenderingContext2D.direction" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetDirection(val CanvasDirection) bool {
	return js.True == bindings.SetCanvasRenderingContext2DDirection(
		this.ref,
		uint32(val),
	)
}

// LetterSpacing returns the value of property "CanvasRenderingContext2D.letterSpacing".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) LetterSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DLetterSpacing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLetterSpacing sets the value of property "CanvasRenderingContext2D.letterSpacing" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetLetterSpacing(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DLetterSpacing(
		this.ref,
		val.Ref(),
	)
}

// FontKerning returns the value of property "CanvasRenderingContext2D.fontKerning".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontKerning() (ret CanvasFontKerning, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontKerning(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontKerning sets the value of property "CanvasRenderingContext2D.fontKerning" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontKerning(val CanvasFontKerning) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontKerning(
		this.ref,
		uint32(val),
	)
}

// FontStretch returns the value of property "CanvasRenderingContext2D.fontStretch".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontStretch() (ret CanvasFontStretch, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontStretch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontStretch sets the value of property "CanvasRenderingContext2D.fontStretch" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontStretch(val CanvasFontStretch) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontStretch(
		this.ref,
		uint32(val),
	)
}

// FontVariantCaps returns the value of property "CanvasRenderingContext2D.fontVariantCaps".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) FontVariantCaps() (ret CanvasFontVariantCaps, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DFontVariantCaps(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontVariantCaps sets the value of property "CanvasRenderingContext2D.fontVariantCaps" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetFontVariantCaps(val CanvasFontVariantCaps) bool {
	return js.True == bindings.SetCanvasRenderingContext2DFontVariantCaps(
		this.ref,
		uint32(val),
	)
}

// TextRendering returns the value of property "CanvasRenderingContext2D.textRendering".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) TextRendering() (ret CanvasTextRendering, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DTextRendering(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextRendering sets the value of property "CanvasRenderingContext2D.textRendering" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetTextRendering(val CanvasTextRendering) bool {
	return js.True == bindings.SetCanvasRenderingContext2DTextRendering(
		this.ref,
		uint32(val),
	)
}

// WordSpacing returns the value of property "CanvasRenderingContext2D.wordSpacing".
//
// It returns ok=false if there is no such property.
func (this CanvasRenderingContext2D) WordSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCanvasRenderingContext2DWordSpacing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWordSpacing sets the value of property "CanvasRenderingContext2D.wordSpacing" to val.
//
// It returns false if the property cannot be set.
func (this CanvasRenderingContext2D) SetWordSpacing(val js.String) bool {
	return js.True == bindings.SetCanvasRenderingContext2DWordSpacing(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetContextAttributes returns true if the method "CanvasRenderingContext2D.getContextAttributes" exists.
func (this CanvasRenderingContext2D) HasFuncGetContextAttributes() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DGetContextAttributes(
		this.ref,
	)
}

// FuncGetContextAttributes returns the method "CanvasRenderingContext2D.getContextAttributes".
func (this CanvasRenderingContext2D) FuncGetContextAttributes() (fn js.Func[func() CanvasRenderingContext2DSettings]) {
	bindings.FuncCanvasRenderingContext2DGetContextAttributes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContextAttributes calls the method "CanvasRenderingContext2D.getContextAttributes".
func (this CanvasRenderingContext2D) GetContextAttributes() (ret CanvasRenderingContext2DSettings) {
	bindings.CallCanvasRenderingContext2DGetContextAttributes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "CanvasRenderingContext2D.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryGetContextAttributes() (ret CanvasRenderingContext2DSettings, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetContextAttributes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSave returns true if the method "CanvasRenderingContext2D.save" exists.
func (this CanvasRenderingContext2D) HasFuncSave() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DSave(
		this.ref,
	)
}

// FuncSave returns the method "CanvasRenderingContext2D.save".
func (this CanvasRenderingContext2D) FuncSave() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DSave(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Save calls the method "CanvasRenderingContext2D.save".
func (this CanvasRenderingContext2D) Save() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSave(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySave calls the method "CanvasRenderingContext2D.save"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TrySave() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSave(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestore returns true if the method "CanvasRenderingContext2D.restore" exists.
func (this CanvasRenderingContext2D) HasFuncRestore() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DRestore(
		this.ref,
	)
}

// FuncRestore returns the method "CanvasRenderingContext2D.restore".
func (this CanvasRenderingContext2D) FuncRestore() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DRestore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Restore calls the method "CanvasRenderingContext2D.restore".
func (this CanvasRenderingContext2D) Restore() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRestore(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRestore calls the method "CanvasRenderingContext2D.restore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryRestore() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRestore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "CanvasRenderingContext2D.reset" exists.
func (this CanvasRenderingContext2D) HasFuncReset() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DReset(
		this.ref,
	)
}

// FuncReset returns the method "CanvasRenderingContext2D.reset".
func (this CanvasRenderingContext2D) FuncReset() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "CanvasRenderingContext2D.reset".
func (this CanvasRenderingContext2D) Reset() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "CanvasRenderingContext2D.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsContextLost returns true if the method "CanvasRenderingContext2D.isContextLost" exists.
func (this CanvasRenderingContext2D) HasFuncIsContextLost() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsContextLost(
		this.ref,
	)
}

// FuncIsContextLost returns the method "CanvasRenderingContext2D.isContextLost".
func (this CanvasRenderingContext2D) FuncIsContextLost() (fn js.Func[func() bool]) {
	bindings.FuncCanvasRenderingContext2DIsContextLost(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsContextLost calls the method "CanvasRenderingContext2D.isContextLost".
func (this CanvasRenderingContext2D) IsContextLost() (ret bool) {
	bindings.CallCanvasRenderingContext2DIsContextLost(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "CanvasRenderingContext2D.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsContextLost(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScale returns true if the method "CanvasRenderingContext2D.scale" exists.
func (this CanvasRenderingContext2D) HasFuncScale() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DScale(
		this.ref,
	)
}

// FuncScale returns the method "CanvasRenderingContext2D.scale".
func (this CanvasRenderingContext2D) FuncScale() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DScale(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scale calls the method "CanvasRenderingContext2D.scale".
func (this CanvasRenderingContext2D) Scale(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScale(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScale calls the method "CanvasRenderingContext2D.scale"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryScale(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScale(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncRotate returns true if the method "CanvasRenderingContext2D.rotate" exists.
func (this CanvasRenderingContext2D) HasFuncRotate() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DRotate(
		this.ref,
	)
}

// FuncRotate returns the method "CanvasRenderingContext2D.rotate".
func (this CanvasRenderingContext2D) FuncRotate() (fn js.Func[func(angle float64)]) {
	bindings.FuncCanvasRenderingContext2DRotate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rotate calls the method "CanvasRenderingContext2D.rotate".
func (this CanvasRenderingContext2D) Rotate(angle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRotate(
		this.ref, js.Pointer(&ret),
		float64(angle),
	)

	return
}

// TryRotate calls the method "CanvasRenderingContext2D.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryRotate(angle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRotate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(angle),
	)

	return
}

// HasFuncTranslate returns true if the method "CanvasRenderingContext2D.translate" exists.
func (this CanvasRenderingContext2D) HasFuncTranslate() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DTranslate(
		this.ref,
	)
}

// FuncTranslate returns the method "CanvasRenderingContext2D.translate".
func (this CanvasRenderingContext2D) FuncTranslate() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DTranslate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Translate calls the method "CanvasRenderingContext2D.translate".
func (this CanvasRenderingContext2D) Translate(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DTranslate(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryTranslate calls the method "CanvasRenderingContext2D.translate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryTranslate(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DTranslate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncTransform returns true if the method "CanvasRenderingContext2D.transform" exists.
func (this CanvasRenderingContext2D) HasFuncTransform() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DTransform(
		this.ref,
	)
}

// FuncTransform returns the method "CanvasRenderingContext2D.transform".
func (this CanvasRenderingContext2D) FuncTransform() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	bindings.FuncCanvasRenderingContext2DTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transform calls the method "CanvasRenderingContext2D.transform".
func (this CanvasRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DTransform(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasFuncGetTransform returns true if the method "CanvasRenderingContext2D.getTransform" exists.
func (this CanvasRenderingContext2D) HasFuncGetTransform() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DGetTransform(
		this.ref,
	)
}

// FuncGetTransform returns the method "CanvasRenderingContext2D.getTransform".
func (this CanvasRenderingContext2D) FuncGetTransform() (fn js.Func[func() DOMMatrix]) {
	bindings.FuncCanvasRenderingContext2DGetTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTransform calls the method "CanvasRenderingContext2D.getTransform".
func (this CanvasRenderingContext2D) GetTransform() (ret DOMMatrix) {
	bindings.CallCanvasRenderingContext2DGetTransform(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetTransform calls the method "CanvasRenderingContext2D.getTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryGetTransform() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetTransform returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasFuncSetTransform() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DSetTransform(
		this.ref,
	)
}

// FuncSetTransform returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) FuncSetTransform() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	bindings.FuncCanvasRenderingContext2DSetTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTransform calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasFuncSetTransform1 returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasFuncSetTransform1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DSetTransform1(
		this.ref,
	)
}

// FuncSetTransform1 returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) FuncSetTransform1() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	bindings.FuncCanvasRenderingContext2DSetTransform1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTransform1 calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TrySetTransform1 calls the method "CanvasRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform1(transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasFuncSetTransform2 returns true if the method "CanvasRenderingContext2D.setTransform" exists.
func (this CanvasRenderingContext2D) HasFuncSetTransform2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DSetTransform2(
		this.ref,
	)
}

// FuncSetTransform2 returns the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) FuncSetTransform2() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DSetTransform2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetTransform2 calls the method "CanvasRenderingContext2D.setTransform".
func (this CanvasRenderingContext2D) SetTransform2() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetTransform2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetTransform2 calls the method "CanvasRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TrySetTransform2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetTransform2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncResetTransform returns true if the method "CanvasRenderingContext2D.resetTransform" exists.
func (this CanvasRenderingContext2D) HasFuncResetTransform() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DResetTransform(
		this.ref,
	)
}

// FuncResetTransform returns the method "CanvasRenderingContext2D.resetTransform".
func (this CanvasRenderingContext2D) FuncResetTransform() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DResetTransform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResetTransform calls the method "CanvasRenderingContext2D.resetTransform".
func (this CanvasRenderingContext2D) ResetTransform() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DResetTransform(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryResetTransform calls the method "CanvasRenderingContext2D.resetTransform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryResetTransform() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DResetTransform(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateLinearGradient returns true if the method "CanvasRenderingContext2D.createLinearGradient" exists.
func (this CanvasRenderingContext2D) HasFuncCreateLinearGradient() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateLinearGradient(
		this.ref,
	)
}

// FuncCreateLinearGradient returns the method "CanvasRenderingContext2D.createLinearGradient".
func (this CanvasRenderingContext2D) FuncCreateLinearGradient() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	bindings.FuncCanvasRenderingContext2DCreateLinearGradient(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateLinearGradient calls the method "CanvasRenderingContext2D.createLinearGradient".
func (this CanvasRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateLinearGradient(
		this.ref, js.Pointer(&ret),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// TryCreateLinearGradient calls the method "CanvasRenderingContext2D.createLinearGradient"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateLinearGradient(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// HasFuncCreateRadialGradient returns true if the method "CanvasRenderingContext2D.createRadialGradient" exists.
func (this CanvasRenderingContext2D) HasFuncCreateRadialGradient() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateRadialGradient(
		this.ref,
	)
}

// FuncCreateRadialGradient returns the method "CanvasRenderingContext2D.createRadialGradient".
func (this CanvasRenderingContext2D) FuncCreateRadialGradient() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	bindings.FuncCanvasRenderingContext2DCreateRadialGradient(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRadialGradient calls the method "CanvasRenderingContext2D.createRadialGradient".
func (this CanvasRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateRadialGradient(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateRadialGradient(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return
}

// HasFuncCreateConicGradient returns true if the method "CanvasRenderingContext2D.createConicGradient" exists.
func (this CanvasRenderingContext2D) HasFuncCreateConicGradient() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateConicGradient(
		this.ref,
	)
}

// FuncCreateConicGradient returns the method "CanvasRenderingContext2D.createConicGradient".
func (this CanvasRenderingContext2D) FuncCreateConicGradient() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	bindings.FuncCanvasRenderingContext2DCreateConicGradient(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateConicGradient calls the method "CanvasRenderingContext2D.createConicGradient".
func (this CanvasRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient) {
	bindings.CallCanvasRenderingContext2DCreateConicGradient(
		this.ref, js.Pointer(&ret),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// TryCreateConicGradient calls the method "CanvasRenderingContext2D.createConicGradient"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateConicGradient(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncCreatePattern returns true if the method "CanvasRenderingContext2D.createPattern" exists.
func (this CanvasRenderingContext2D) HasFuncCreatePattern() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreatePattern(
		this.ref,
	)
}

// FuncCreatePattern returns the method "CanvasRenderingContext2D.createPattern".
func (this CanvasRenderingContext2D) FuncCreatePattern() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	bindings.FuncCanvasRenderingContext2DCreatePattern(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePattern calls the method "CanvasRenderingContext2D.createPattern".
func (this CanvasRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern) {
	bindings.CallCanvasRenderingContext2DCreatePattern(
		this.ref, js.Pointer(&ret),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// TryCreatePattern calls the method "CanvasRenderingContext2D.createPattern"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreatePattern(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// HasFuncClearRect returns true if the method "CanvasRenderingContext2D.clearRect" exists.
func (this CanvasRenderingContext2D) HasFuncClearRect() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClearRect(
		this.ref,
	)
}

// FuncClearRect returns the method "CanvasRenderingContext2D.clearRect".
func (this CanvasRenderingContext2D) FuncClearRect() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncCanvasRenderingContext2DClearRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearRect calls the method "CanvasRenderingContext2D.clearRect".
func (this CanvasRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClearRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryClearRect calls the method "CanvasRenderingContext2D.clearRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClearRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClearRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncFillRect returns true if the method "CanvasRenderingContext2D.fillRect" exists.
func (this CanvasRenderingContext2D) HasFuncFillRect() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFillRect(
		this.ref,
	)
}

// FuncFillRect returns the method "CanvasRenderingContext2D.fillRect".
func (this CanvasRenderingContext2D) FuncFillRect() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncCanvasRenderingContext2DFillRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FillRect calls the method "CanvasRenderingContext2D.fillRect".
func (this CanvasRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryFillRect calls the method "CanvasRenderingContext2D.fillRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFillRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncStrokeRect returns true if the method "CanvasRenderingContext2D.strokeRect" exists.
func (this CanvasRenderingContext2D) HasFuncStrokeRect() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DStrokeRect(
		this.ref,
	)
}

// FuncStrokeRect returns the method "CanvasRenderingContext2D.strokeRect".
func (this CanvasRenderingContext2D) FuncStrokeRect() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncCanvasRenderingContext2DStrokeRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StrokeRect calls the method "CanvasRenderingContext2D.strokeRect".
func (this CanvasRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryStrokeRect calls the method "CanvasRenderingContext2D.strokeRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncBeginPath returns true if the method "CanvasRenderingContext2D.beginPath" exists.
func (this CanvasRenderingContext2D) HasFuncBeginPath() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DBeginPath(
		this.ref,
	)
}

// FuncBeginPath returns the method "CanvasRenderingContext2D.beginPath".
func (this CanvasRenderingContext2D) FuncBeginPath() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DBeginPath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginPath calls the method "CanvasRenderingContext2D.beginPath".
func (this CanvasRenderingContext2D) BeginPath() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DBeginPath(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBeginPath calls the method "CanvasRenderingContext2D.beginPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryBeginPath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DBeginPath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFill returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFuncFill() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFill(
		this.ref,
	)
}

// FuncFill returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) FuncFill() (fn js.Func[func(fillRule CanvasFillRule)]) {
	bindings.FuncCanvasRenderingContext2DFill(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fill calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill(
		this.ref, js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryFill calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFill(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasFuncFill1 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFuncFill1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFill1(
		this.ref,
	)
}

// FuncFill1 returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) FuncFill1() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DFill1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fill1 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill1() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFill1 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFill1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFill2 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFuncFill2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFill2(
		this.ref,
	)
}

// FuncFill2 returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) FuncFill2() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	bindings.FuncCanvasRenderingContext2DFill2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fill2 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill2(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryFill2 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFill2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasFuncFill3 returns true if the method "CanvasRenderingContext2D.fill" exists.
func (this CanvasRenderingContext2D) HasFuncFill3() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFill3(
		this.ref,
	)
}

// FuncFill3 returns the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) FuncFill3() (fn js.Func[func(path Path2D)]) {
	bindings.FuncCanvasRenderingContext2DFill3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Fill3 calls the method "CanvasRenderingContext2D.fill".
func (this CanvasRenderingContext2D) Fill3(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFill3(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryFill3 calls the method "CanvasRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFill3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFill3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncStroke returns true if the method "CanvasRenderingContext2D.stroke" exists.
func (this CanvasRenderingContext2D) HasFuncStroke() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DStroke(
		this.ref,
	)
}

// FuncStroke returns the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) FuncStroke() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DStroke(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stroke calls the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) Stroke() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStroke(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStroke calls the method "CanvasRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryStroke() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStroke(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStroke1 returns true if the method "CanvasRenderingContext2D.stroke" exists.
func (this CanvasRenderingContext2D) HasFuncStroke1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DStroke1(
		this.ref,
	)
}

// FuncStroke1 returns the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) FuncStroke1() (fn js.Func[func(path Path2D)]) {
	bindings.FuncCanvasRenderingContext2DStroke1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stroke1 calls the method "CanvasRenderingContext2D.stroke".
func (this CanvasRenderingContext2D) Stroke1(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStroke1(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryStroke1 calls the method "CanvasRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryStroke1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStroke1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncClip returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasFuncClip() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClip(
		this.ref,
	)
}

// FuncClip returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) FuncClip() (fn js.Func[func(fillRule CanvasFillRule)]) {
	bindings.FuncCanvasRenderingContext2DClip(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clip calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip(
		this.ref, js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryClip calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClip(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasFuncClip1 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasFuncClip1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClip1(
		this.ref,
	)
}

// FuncClip1 returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) FuncClip1() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DClip1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clip1 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip1() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClip1 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClip1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClip2 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasFuncClip2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClip2(
		this.ref,
	)
}

// FuncClip2 returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) FuncClip2() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	bindings.FuncCanvasRenderingContext2DClip2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clip2 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip2(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryClip2 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClip2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasFuncClip3 returns true if the method "CanvasRenderingContext2D.clip" exists.
func (this CanvasRenderingContext2D) HasFuncClip3() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClip3(
		this.ref,
	)
}

// FuncClip3 returns the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) FuncClip3() (fn js.Func[func(path Path2D)]) {
	bindings.FuncCanvasRenderingContext2DClip3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clip3 calls the method "CanvasRenderingContext2D.clip".
func (this CanvasRenderingContext2D) Clip3(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClip3(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryClip3 calls the method "CanvasRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClip3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClip3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncIsPointInPath returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInPath() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInPath(
		this.ref,
	)
}

// FuncIsPointInPath returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) FuncIsPointInPath() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInPath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInPath calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasFuncIsPointInPath1 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInPath1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInPath1(
		this.ref,
	)
}

// FuncIsPointInPath1 returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) FuncIsPointInPath1() (fn js.Func[func(x float64, y float64) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInPath1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInPath1 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath1(x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath1 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath1(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncIsPointInPath2 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInPath2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInPath2(
		this.ref,
	)
}

// FuncIsPointInPath2 returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) FuncIsPointInPath2() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInPath2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInPath2 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath2(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath2 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasFuncIsPointInPath3 returns true if the method "CanvasRenderingContext2D.isPointInPath" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInPath3() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInPath3(
		this.ref,
	)
}

// FuncIsPointInPath3 returns the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) FuncIsPointInPath3() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInPath3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInPath3 calls the method "CanvasRenderingContext2D.isPointInPath".
func (this CanvasRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInPath3(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath3 calls the method "CanvasRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInPath3(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInPath3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncIsPointInStroke returns true if the method "CanvasRenderingContext2D.isPointInStroke" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInStroke() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInStroke(
		this.ref,
	)
}

// FuncIsPointInStroke returns the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) FuncIsPointInStroke() (fn js.Func[func(x float64, y float64) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInStroke(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInStroke calls the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStroke(x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInStroke(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke calls the method "CanvasRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInStroke(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInStroke(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncIsPointInStroke1 returns true if the method "CanvasRenderingContext2D.isPointInStroke" exists.
func (this CanvasRenderingContext2D) HasFuncIsPointInStroke1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DIsPointInStroke1(
		this.ref,
	)
}

// FuncIsPointInStroke1 returns the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) FuncIsPointInStroke1() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	bindings.FuncCanvasRenderingContext2DIsPointInStroke1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPointInStroke1 calls the method "CanvasRenderingContext2D.isPointInStroke".
func (this CanvasRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallCanvasRenderingContext2DIsPointInStroke1(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke1 calls the method "CanvasRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryIsPointInStroke1(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DIsPointInStroke1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncDrawFocusIfNeeded returns true if the method "CanvasRenderingContext2D.drawFocusIfNeeded" exists.
func (this CanvasRenderingContext2D) HasFuncDrawFocusIfNeeded() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DDrawFocusIfNeeded(
		this.ref,
	)
}

// FuncDrawFocusIfNeeded returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) FuncDrawFocusIfNeeded() (fn js.Func[func(element Element)]) {
	bindings.FuncCanvasRenderingContext2DDrawFocusIfNeeded(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawFocusIfNeeded calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeeded(element Element) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryDrawFocusIfNeeded calls the method "CanvasRenderingContext2D.drawFocusIfNeeded"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawFocusIfNeeded(element Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawFocusIfNeeded(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasFuncDrawFocusIfNeeded1 returns true if the method "CanvasRenderingContext2D.drawFocusIfNeeded" exists.
func (this CanvasRenderingContext2D) HasFuncDrawFocusIfNeeded1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.ref,
	)
}

// FuncDrawFocusIfNeeded1 returns the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) FuncDrawFocusIfNeeded1() (fn js.Func[func(path Path2D, element Element)]) {
	bindings.FuncCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawFocusIfNeeded1 calls the method "CanvasRenderingContext2D.drawFocusIfNeeded".
func (this CanvasRenderingContext2D) DrawFocusIfNeeded1(path Path2D, element Element) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.ref, js.Pointer(&ret),
		path.Ref(),
		element.Ref(),
	)

	return
}

// TryDrawFocusIfNeeded1 calls the method "CanvasRenderingContext2D.drawFocusIfNeeded"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawFocusIfNeeded1(path Path2D, element Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawFocusIfNeeded1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		element.Ref(),
	)

	return
}

// HasFuncScrollPathIntoView returns true if the method "CanvasRenderingContext2D.scrollPathIntoView" exists.
func (this CanvasRenderingContext2D) HasFuncScrollPathIntoView() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DScrollPathIntoView(
		this.ref,
	)
}

// FuncScrollPathIntoView returns the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) FuncScrollPathIntoView() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DScrollPathIntoView(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollPathIntoView calls the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoView() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScrollPathIntoView(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollPathIntoView calls the method "CanvasRenderingContext2D.scrollPathIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryScrollPathIntoView() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScrollPathIntoView(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollPathIntoView1 returns true if the method "CanvasRenderingContext2D.scrollPathIntoView" exists.
func (this CanvasRenderingContext2D) HasFuncScrollPathIntoView1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DScrollPathIntoView1(
		this.ref,
	)
}

// FuncScrollPathIntoView1 returns the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) FuncScrollPathIntoView1() (fn js.Func[func(path Path2D)]) {
	bindings.FuncCanvasRenderingContext2DScrollPathIntoView1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollPathIntoView1 calls the method "CanvasRenderingContext2D.scrollPathIntoView".
func (this CanvasRenderingContext2D) ScrollPathIntoView1(path Path2D) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DScrollPathIntoView1(
		this.ref, js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryScrollPathIntoView1 calls the method "CanvasRenderingContext2D.scrollPathIntoView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryScrollPathIntoView1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DScrollPathIntoView1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncFillText returns true if the method "CanvasRenderingContext2D.fillText" exists.
func (this CanvasRenderingContext2D) HasFuncFillText() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFillText(
		this.ref,
	)
}

// FuncFillText returns the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FuncFillText() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	bindings.FuncCanvasRenderingContext2DFillText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FillText calls the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillText(
		this.ref, js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// TryFillText calls the method "CanvasRenderingContext2D.fillText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFillText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// HasFuncFillText1 returns true if the method "CanvasRenderingContext2D.fillText" exists.
func (this CanvasRenderingContext2D) HasFuncFillText1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DFillText1(
		this.ref,
	)
}

// FuncFillText1 returns the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FuncFillText1() (fn js.Func[func(text js.String, x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DFillText1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FillText1 calls the method "CanvasRenderingContext2D.fillText".
func (this CanvasRenderingContext2D) FillText1(text js.String, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DFillText1(
		this.ref, js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryFillText1 calls the method "CanvasRenderingContext2D.fillText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryFillText1(text js.String, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DFillText1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncStrokeText returns true if the method "CanvasRenderingContext2D.strokeText" exists.
func (this CanvasRenderingContext2D) HasFuncStrokeText() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DStrokeText(
		this.ref,
	)
}

// FuncStrokeText returns the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) FuncStrokeText() (fn js.Func[func(text js.String, x float64, y float64, maxWidth float64)]) {
	bindings.FuncCanvasRenderingContext2DStrokeText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StrokeText calls the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeText(
		this.ref, js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// TryStrokeText calls the method "CanvasRenderingContext2D.strokeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeText(text js.String, x float64, y float64, maxWidth float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
		float64(maxWidth),
	)

	return
}

// HasFuncStrokeText1 returns true if the method "CanvasRenderingContext2D.strokeText" exists.
func (this CanvasRenderingContext2D) HasFuncStrokeText1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DStrokeText1(
		this.ref,
	)
}

// FuncStrokeText1 returns the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) FuncStrokeText1() (fn js.Func[func(text js.String, x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DStrokeText1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StrokeText1 calls the method "CanvasRenderingContext2D.strokeText".
func (this CanvasRenderingContext2D) StrokeText1(text js.String, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DStrokeText1(
		this.ref, js.Pointer(&ret),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryStrokeText1 calls the method "CanvasRenderingContext2D.strokeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryStrokeText1(text js.String, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DStrokeText1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncMeasureText returns true if the method "CanvasRenderingContext2D.measureText" exists.
func (this CanvasRenderingContext2D) HasFuncMeasureText() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DMeasureText(
		this.ref,
	)
}

// FuncMeasureText returns the method "CanvasRenderingContext2D.measureText".
func (this CanvasRenderingContext2D) FuncMeasureText() (fn js.Func[func(text js.String) TextMetrics]) {
	bindings.FuncCanvasRenderingContext2DMeasureText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MeasureText calls the method "CanvasRenderingContext2D.measureText".
func (this CanvasRenderingContext2D) MeasureText(text js.String) (ret TextMetrics) {
	bindings.CallCanvasRenderingContext2DMeasureText(
		this.ref, js.Pointer(&ret),
		text.Ref(),
	)

	return
}

// TryMeasureText calls the method "CanvasRenderingContext2D.measureText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryMeasureText(text js.String) (ret TextMetrics, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DMeasureText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		text.Ref(),
	)

	return
}

// HasFuncDrawImage returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasFuncDrawImage() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DDrawImage(
		this.ref,
	)
}

// FuncDrawImage returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) FuncDrawImage() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	bindings.FuncCanvasRenderingContext2DDrawImage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawImage calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage(
		this.ref, js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// TryDrawImage calls the method "CanvasRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// HasFuncDrawImage1 returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasFuncDrawImage1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DDrawImage1(
		this.ref,
	)
}

// FuncDrawImage1 returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) FuncDrawImage1() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	bindings.FuncCanvasRenderingContext2DDrawImage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawImage1 calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage1(
		this.ref, js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// TryDrawImage1 calls the method "CanvasRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// HasFuncDrawImage2 returns true if the method "CanvasRenderingContext2D.drawImage" exists.
func (this CanvasRenderingContext2D) HasFuncDrawImage2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DDrawImage2(
		this.ref,
	)
}

// FuncDrawImage2 returns the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) FuncDrawImage2() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	bindings.FuncCanvasRenderingContext2DDrawImage2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawImage2 calls the method "CanvasRenderingContext2D.drawImage".
func (this CanvasRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DDrawImage2(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryDrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DDrawImage2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCreateImageData returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasFuncCreateImageData() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateImageData(
		this.ref,
	)
}

// FuncCreateImageData returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) FuncCreateImageData() (fn js.Func[func(sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	bindings.FuncCanvasRenderingContext2DCreateImageData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageData calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData(sw int32, sh int32, settings ImageDataSettings) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData(
		this.ref, js.Pointer(&ret),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// TryCreateImageData calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData(sw int32, sh int32, settings ImageDataSettings) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// HasFuncCreateImageData1 returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasFuncCreateImageData1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateImageData1(
		this.ref,
	)
}

// FuncCreateImageData1 returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) FuncCreateImageData1() (fn js.Func[func(sw int32, sh int32) ImageData]) {
	bindings.FuncCanvasRenderingContext2DCreateImageData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageData1 calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData1(sw int32, sh int32) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData1(
		this.ref, js.Pointer(&ret),
		int32(sw),
		int32(sh),
	)

	return
}

// TryCreateImageData1 calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData1(sw int32, sh int32) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(sw),
		int32(sh),
	)

	return
}

// HasFuncCreateImageData2 returns true if the method "CanvasRenderingContext2D.createImageData" exists.
func (this CanvasRenderingContext2D) HasFuncCreateImageData2() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DCreateImageData2(
		this.ref,
	)
}

// FuncCreateImageData2 returns the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) FuncCreateImageData2() (fn js.Func[func(imagedata ImageData) ImageData]) {
	bindings.FuncCanvasRenderingContext2DCreateImageData2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateImageData2 calls the method "CanvasRenderingContext2D.createImageData".
func (this CanvasRenderingContext2D) CreateImageData2(imagedata ImageData) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DCreateImageData2(
		this.ref, js.Pointer(&ret),
		imagedata.Ref(),
	)

	return
}

// TryCreateImageData2 calls the method "CanvasRenderingContext2D.createImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryCreateImageData2(imagedata ImageData) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DCreateImageData2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		imagedata.Ref(),
	)

	return
}

// HasFuncGetImageData returns true if the method "CanvasRenderingContext2D.getImageData" exists.
func (this CanvasRenderingContext2D) HasFuncGetImageData() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DGetImageData(
		this.ref,
	)
}

// FuncGetImageData returns the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) FuncGetImageData() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) ImageData]) {
	bindings.FuncCanvasRenderingContext2DGetImageData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetImageData calls the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DGetImageData(
		this.ref, js.Pointer(&ret),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// TryGetImageData calls the method "CanvasRenderingContext2D.getImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryGetImageData(sx int32, sy int32, sw int32, sh int32, settings ImageDataSettings) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetImageData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
		js.Pointer(&settings),
	)

	return
}

// HasFuncGetImageData1 returns true if the method "CanvasRenderingContext2D.getImageData" exists.
func (this CanvasRenderingContext2D) HasFuncGetImageData1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DGetImageData1(
		this.ref,
	)
}

// FuncGetImageData1 returns the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) FuncGetImageData1() (fn js.Func[func(sx int32, sy int32, sw int32, sh int32) ImageData]) {
	bindings.FuncCanvasRenderingContext2DGetImageData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetImageData1 calls the method "CanvasRenderingContext2D.getImageData".
func (this CanvasRenderingContext2D) GetImageData1(sx int32, sy int32, sw int32, sh int32) (ret ImageData) {
	bindings.CallCanvasRenderingContext2DGetImageData1(
		this.ref, js.Pointer(&ret),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// TryGetImageData1 calls the method "CanvasRenderingContext2D.getImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryGetImageData1(sx int32, sy int32, sw int32, sh int32) (ret ImageData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetImageData1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(sx),
		int32(sy),
		int32(sw),
		int32(sh),
	)

	return
}

// HasFuncPutImageData returns true if the method "CanvasRenderingContext2D.putImageData" exists.
func (this CanvasRenderingContext2D) HasFuncPutImageData() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DPutImageData(
		this.ref,
	)
}

// FuncPutImageData returns the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) FuncPutImageData() (fn js.Func[func(imagedata ImageData, dx int32, dy int32)]) {
	bindings.FuncCanvasRenderingContext2DPutImageData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PutImageData calls the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageData(imagedata ImageData, dx int32, dy int32) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DPutImageData(
		this.ref, js.Pointer(&ret),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	return
}

// TryPutImageData calls the method "CanvasRenderingContext2D.putImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryPutImageData(imagedata ImageData, dx int32, dy int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DPutImageData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		imagedata.Ref(),
		int32(dx),
		int32(dy),
	)

	return
}

// HasFuncPutImageData1 returns true if the method "CanvasRenderingContext2D.putImageData" exists.
func (this CanvasRenderingContext2D) HasFuncPutImageData1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DPutImageData1(
		this.ref,
	)
}

// FuncPutImageData1 returns the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) FuncPutImageData1() (fn js.Func[func(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32)]) {
	bindings.FuncCanvasRenderingContext2DPutImageData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PutImageData1 calls the method "CanvasRenderingContext2D.putImageData".
func (this CanvasRenderingContext2D) PutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DPutImageData1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryPutImageData1(imagedata ImageData, dx int32, dy int32, dirtyX int32, dirtyY int32, dirtyWidth int32, dirtyHeight int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DPutImageData1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncSetLineDash returns true if the method "CanvasRenderingContext2D.setLineDash" exists.
func (this CanvasRenderingContext2D) HasFuncSetLineDash() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DSetLineDash(
		this.ref,
	)
}

// FuncSetLineDash returns the method "CanvasRenderingContext2D.setLineDash".
func (this CanvasRenderingContext2D) FuncSetLineDash() (fn js.Func[func(segments js.Array[float64])]) {
	bindings.FuncCanvasRenderingContext2DSetLineDash(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetLineDash calls the method "CanvasRenderingContext2D.setLineDash".
func (this CanvasRenderingContext2D) SetLineDash(segments js.Array[float64]) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DSetLineDash(
		this.ref, js.Pointer(&ret),
		segments.Ref(),
	)

	return
}

// TrySetLineDash calls the method "CanvasRenderingContext2D.setLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TrySetLineDash(segments js.Array[float64]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DSetLineDash(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		segments.Ref(),
	)

	return
}

// HasFuncGetLineDash returns true if the method "CanvasRenderingContext2D.getLineDash" exists.
func (this CanvasRenderingContext2D) HasFuncGetLineDash() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DGetLineDash(
		this.ref,
	)
}

// FuncGetLineDash returns the method "CanvasRenderingContext2D.getLineDash".
func (this CanvasRenderingContext2D) FuncGetLineDash() (fn js.Func[func() js.Array[float64]]) {
	bindings.FuncCanvasRenderingContext2DGetLineDash(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetLineDash calls the method "CanvasRenderingContext2D.getLineDash".
func (this CanvasRenderingContext2D) GetLineDash() (ret js.Array[float64]) {
	bindings.CallCanvasRenderingContext2DGetLineDash(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetLineDash calls the method "CanvasRenderingContext2D.getLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryGetLineDash() (ret js.Array[float64], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DGetLineDash(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClosePath returns true if the method "CanvasRenderingContext2D.closePath" exists.
func (this CanvasRenderingContext2D) HasFuncClosePath() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DClosePath(
		this.ref,
	)
}

// FuncClosePath returns the method "CanvasRenderingContext2D.closePath".
func (this CanvasRenderingContext2D) FuncClosePath() (fn js.Func[func()]) {
	bindings.FuncCanvasRenderingContext2DClosePath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClosePath calls the method "CanvasRenderingContext2D.closePath".
func (this CanvasRenderingContext2D) ClosePath() (ret js.Void) {
	bindings.CallCanvasRenderingContext2DClosePath(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClosePath calls the method "CanvasRenderingContext2D.closePath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryClosePath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DClosePath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveTo returns true if the method "CanvasRenderingContext2D.moveTo" exists.
func (this CanvasRenderingContext2D) HasFuncMoveTo() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DMoveTo(
		this.ref,
	)
}

// FuncMoveTo returns the method "CanvasRenderingContext2D.moveTo".
func (this CanvasRenderingContext2D) FuncMoveTo() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DMoveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveTo calls the method "CanvasRenderingContext2D.moveTo".
func (this CanvasRenderingContext2D) MoveTo(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DMoveTo(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryMoveTo calls the method "CanvasRenderingContext2D.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryMoveTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DMoveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncLineTo returns true if the method "CanvasRenderingContext2D.lineTo" exists.
func (this CanvasRenderingContext2D) HasFuncLineTo() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DLineTo(
		this.ref,
	)
}

// FuncLineTo returns the method "CanvasRenderingContext2D.lineTo".
func (this CanvasRenderingContext2D) FuncLineTo() (fn js.Func[func(x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DLineTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LineTo calls the method "CanvasRenderingContext2D.lineTo".
func (this CanvasRenderingContext2D) LineTo(x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DLineTo(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryLineTo calls the method "CanvasRenderingContext2D.lineTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryLineTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DLineTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncQuadraticCurveTo returns true if the method "CanvasRenderingContext2D.quadraticCurveTo" exists.
func (this CanvasRenderingContext2D) HasFuncQuadraticCurveTo() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DQuadraticCurveTo(
		this.ref,
	)
}

// FuncQuadraticCurveTo returns the method "CanvasRenderingContext2D.quadraticCurveTo".
func (this CanvasRenderingContext2D) FuncQuadraticCurveTo() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DQuadraticCurveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QuadraticCurveTo calls the method "CanvasRenderingContext2D.quadraticCurveTo".
func (this CanvasRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DQuadraticCurveTo(
		this.ref, js.Pointer(&ret),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// TryQuadraticCurveTo calls the method "CanvasRenderingContext2D.quadraticCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryQuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DQuadraticCurveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncBezierCurveTo returns true if the method "CanvasRenderingContext2D.bezierCurveTo" exists.
func (this CanvasRenderingContext2D) HasFuncBezierCurveTo() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DBezierCurveTo(
		this.ref,
	)
}

// FuncBezierCurveTo returns the method "CanvasRenderingContext2D.bezierCurveTo".
func (this CanvasRenderingContext2D) FuncBezierCurveTo() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	bindings.FuncCanvasRenderingContext2DBezierCurveTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BezierCurveTo calls the method "CanvasRenderingContext2D.bezierCurveTo".
func (this CanvasRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DBezierCurveTo(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryBezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DBezierCurveTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncArcTo returns true if the method "CanvasRenderingContext2D.arcTo" exists.
func (this CanvasRenderingContext2D) HasFuncArcTo() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DArcTo(
		this.ref,
	)
}

// FuncArcTo returns the method "CanvasRenderingContext2D.arcTo".
func (this CanvasRenderingContext2D) FuncArcTo() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	bindings.FuncCanvasRenderingContext2DArcTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArcTo calls the method "CanvasRenderingContext2D.arcTo".
func (this CanvasRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArcTo(
		this.ref, js.Pointer(&ret),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// TryArcTo calls the method "CanvasRenderingContext2D.arcTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArcTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// HasFuncRect returns true if the method "CanvasRenderingContext2D.rect" exists.
func (this CanvasRenderingContext2D) HasFuncRect() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DRect(
		this.ref,
	)
}

// FuncRect returns the method "CanvasRenderingContext2D.rect".
func (this CanvasRenderingContext2D) FuncRect() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncCanvasRenderingContext2DRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Rect calls the method "CanvasRenderingContext2D.rect".
func (this CanvasRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRect calls the method "CanvasRenderingContext2D.rect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncRoundRect returns true if the method "CanvasRenderingContext2D.roundRect" exists.
func (this CanvasRenderingContext2D) HasFuncRoundRect() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DRoundRect(
		this.ref,
	)
}

// FuncRoundRect returns the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) FuncRoundRect() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	bindings.FuncCanvasRenderingContext2DRoundRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RoundRect calls the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRoundRect(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// TryRoundRect calls the method "CanvasRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryRoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRoundRect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// HasFuncRoundRect1 returns true if the method "CanvasRenderingContext2D.roundRect" exists.
func (this CanvasRenderingContext2D) HasFuncRoundRect1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DRoundRect1(
		this.ref,
	)
}

// FuncRoundRect1 returns the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) FuncRoundRect1() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	bindings.FuncCanvasRenderingContext2DRoundRect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RoundRect1 calls the method "CanvasRenderingContext2D.roundRect".
func (this CanvasRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DRoundRect1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRoundRect1 calls the method "CanvasRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryRoundRect1(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DRoundRect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFuncArc returns true if the method "CanvasRenderingContext2D.arc" exists.
func (this CanvasRenderingContext2D) HasFuncArc() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DArc(
		this.ref,
	)
}

// FuncArc returns the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) FuncArc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	bindings.FuncCanvasRenderingContext2DArc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Arc calls the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArc(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryArc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArc(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasFuncArc1 returns true if the method "CanvasRenderingContext2D.arc" exists.
func (this CanvasRenderingContext2D) HasFuncArc1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DArc1(
		this.ref,
	)
}

// FuncArc1 returns the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) FuncArc1() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	bindings.FuncCanvasRenderingContext2DArc1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Arc1 calls the method "CanvasRenderingContext2D.arc".
func (this CanvasRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DArc1(
		this.ref, js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryArc1 calls the method "CanvasRenderingContext2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryArc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DArc1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasFuncEllipse returns true if the method "CanvasRenderingContext2D.ellipse" exists.
func (this CanvasRenderingContext2D) HasFuncEllipse() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DEllipse(
		this.ref,
	)
}

// FuncEllipse returns the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) FuncEllipse() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	bindings.FuncCanvasRenderingContext2DEllipse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ellipse calls the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DEllipse(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryEllipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DEllipse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncEllipse1 returns true if the method "CanvasRenderingContext2D.ellipse" exists.
func (this CanvasRenderingContext2D) HasFuncEllipse1() bool {
	return js.True == bindings.HasFuncCanvasRenderingContext2DEllipse1(
		this.ref,
	)
}

// FuncEllipse1 returns the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) FuncEllipse1() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	bindings.FuncCanvasRenderingContext2DEllipse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ellipse1 calls the method "CanvasRenderingContext2D.ellipse".
func (this CanvasRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallCanvasRenderingContext2DEllipse1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CanvasRenderingContext2D) TryEllipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanvasRenderingContext2DEllipse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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

type HTMLCanvasElement struct {
	HTMLElement
}

func (this HTMLCanvasElement) Once() HTMLCanvasElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "HTMLCanvasElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLCanvasElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCanvasElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLCanvasElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLCanvasElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLCanvasElementWidth(
		this.ref,
		uint32(val),
	)
}

// Height returns the value of property "HTMLCanvasElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLCanvasElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLCanvasElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLCanvasElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLCanvasElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLCanvasElementHeight(
		this.ref,
		uint32(val),
	)
}

// HasFuncGetContext returns true if the method "HTMLCanvasElement.getContext" exists.
func (this HTMLCanvasElement) HasFuncGetContext() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementGetContext(
		this.ref,
	)
}

// FuncGetContext returns the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) FuncGetContext() (fn js.Func[func(contextId js.String, options js.Any) RenderingContext]) {
	bindings.FuncHTMLCanvasElementGetContext(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContext calls the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContext(contextId js.String, options js.Any) (ret RenderingContext) {
	bindings.CallHTMLCanvasElementGetContext(
		this.ref, js.Pointer(&ret),
		contextId.Ref(),
		options.Ref(),
	)

	return
}

// TryGetContext calls the method "HTMLCanvasElement.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryGetContext(contextId js.String, options js.Any) (ret RenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementGetContext(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		contextId.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncGetContext1 returns true if the method "HTMLCanvasElement.getContext" exists.
func (this HTMLCanvasElement) HasFuncGetContext1() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementGetContext1(
		this.ref,
	)
}

// FuncGetContext1 returns the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) FuncGetContext1() (fn js.Func[func(contextId js.String) RenderingContext]) {
	bindings.FuncHTMLCanvasElementGetContext1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContext1 calls the method "HTMLCanvasElement.getContext".
func (this HTMLCanvasElement) GetContext1(contextId js.String) (ret RenderingContext) {
	bindings.CallHTMLCanvasElementGetContext1(
		this.ref, js.Pointer(&ret),
		contextId.Ref(),
	)

	return
}

// TryGetContext1 calls the method "HTMLCanvasElement.getContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryGetContext1(contextId js.String) (ret RenderingContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementGetContext1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		contextId.Ref(),
	)

	return
}

// HasFuncToDataURL returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasFuncToDataURL() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToDataURL(
		this.ref,
	)
}

// FuncToDataURL returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) FuncToDataURL() (fn js.Func[func(typ js.String, quality js.Any) js.String]) {
	bindings.FuncHTMLCanvasElementToDataURL(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToDataURL calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL(typ js.String, quality js.Any) (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// TryToDataURL calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL(typ js.String, quality js.Any) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// HasFuncToDataURL1 returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasFuncToDataURL1() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToDataURL1(
		this.ref,
	)
}

// FuncToDataURL1 returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) FuncToDataURL1() (fn js.Func[func(typ js.String) js.String]) {
	bindings.FuncHTMLCanvasElementToDataURL1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToDataURL1 calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL1(typ js.String) (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryToDataURL1 calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL1(typ js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncToDataURL2 returns true if the method "HTMLCanvasElement.toDataURL" exists.
func (this HTMLCanvasElement) HasFuncToDataURL2() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToDataURL2(
		this.ref,
	)
}

// FuncToDataURL2 returns the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) FuncToDataURL2() (fn js.Func[func() js.String]) {
	bindings.FuncHTMLCanvasElementToDataURL2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToDataURL2 calls the method "HTMLCanvasElement.toDataURL".
func (this HTMLCanvasElement) ToDataURL2() (ret js.String) {
	bindings.CallHTMLCanvasElementToDataURL2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToDataURL2 calls the method "HTMLCanvasElement.toDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToDataURL2() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToDataURL2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToBlob returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasFuncToBlob() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToBlob(
		this.ref,
	)
}

// FuncToBlob returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) FuncToBlob() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String, quality js.Any)]) {
	bindings.FuncHTMLCanvasElementToBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToBlob calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob(callback js.Func[func(blob Blob)], typ js.String, quality js.Any) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// TryToBlob calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToBlob(callback js.Func[func(blob Blob)], typ js.String, quality js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		typ.Ref(),
		quality.Ref(),
	)

	return
}

// HasFuncToBlob1 returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasFuncToBlob1() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToBlob1(
		this.ref,
	)
}

// FuncToBlob1 returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) FuncToBlob1() (fn js.Func[func(callback js.Func[func(blob Blob)], typ js.String)]) {
	bindings.FuncHTMLCanvasElementToBlob1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToBlob1 calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob1(callback js.Func[func(blob Blob)], typ js.String) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob1(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
		typ.Ref(),
	)

	return
}

// TryToBlob1 calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToBlob1(callback js.Func[func(blob Blob)], typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		typ.Ref(),
	)

	return
}

// HasFuncToBlob2 returns true if the method "HTMLCanvasElement.toBlob" exists.
func (this HTMLCanvasElement) HasFuncToBlob2() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementToBlob2(
		this.ref,
	)
}

// FuncToBlob2 returns the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) FuncToBlob2() (fn js.Func[func(callback js.Func[func(blob Blob)])]) {
	bindings.FuncHTMLCanvasElementToBlob2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToBlob2 calls the method "HTMLCanvasElement.toBlob".
func (this HTMLCanvasElement) ToBlob2(callback js.Func[func(blob Blob)]) (ret js.Void) {
	bindings.CallHTMLCanvasElementToBlob2(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryToBlob2 calls the method "HTMLCanvasElement.toBlob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryToBlob2(callback js.Func[func(blob Blob)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementToBlob2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncTransferControlToOffscreen returns true if the method "HTMLCanvasElement.transferControlToOffscreen" exists.
func (this HTMLCanvasElement) HasFuncTransferControlToOffscreen() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementTransferControlToOffscreen(
		this.ref,
	)
}

// FuncTransferControlToOffscreen returns the method "HTMLCanvasElement.transferControlToOffscreen".
func (this HTMLCanvasElement) FuncTransferControlToOffscreen() (fn js.Func[func() OffscreenCanvas]) {
	bindings.FuncHTMLCanvasElementTransferControlToOffscreen(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransferControlToOffscreen calls the method "HTMLCanvasElement.transferControlToOffscreen".
func (this HTMLCanvasElement) TransferControlToOffscreen() (ret OffscreenCanvas) {
	bindings.CallHTMLCanvasElementTransferControlToOffscreen(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTransferControlToOffscreen calls the method "HTMLCanvasElement.transferControlToOffscreen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryTransferControlToOffscreen() (ret OffscreenCanvas, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementTransferControlToOffscreen(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCaptureStream returns true if the method "HTMLCanvasElement.captureStream" exists.
func (this HTMLCanvasElement) HasFuncCaptureStream() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementCaptureStream(
		this.ref,
	)
}

// FuncCaptureStream returns the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) FuncCaptureStream() (fn js.Func[func(frameRequestRate float64) MediaStream]) {
	bindings.FuncHTMLCanvasElementCaptureStream(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CaptureStream calls the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStream(frameRequestRate float64) (ret MediaStream) {
	bindings.CallHTMLCanvasElementCaptureStream(
		this.ref, js.Pointer(&ret),
		float64(frameRequestRate),
	)

	return
}

// TryCaptureStream calls the method "HTMLCanvasElement.captureStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryCaptureStream(frameRequestRate float64) (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementCaptureStream(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(frameRequestRate),
	)

	return
}

// HasFuncCaptureStream1 returns true if the method "HTMLCanvasElement.captureStream" exists.
func (this HTMLCanvasElement) HasFuncCaptureStream1() bool {
	return js.True == bindings.HasFuncHTMLCanvasElementCaptureStream1(
		this.ref,
	)
}

// FuncCaptureStream1 returns the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) FuncCaptureStream1() (fn js.Func[func() MediaStream]) {
	bindings.FuncHTMLCanvasElementCaptureStream1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CaptureStream1 calls the method "HTMLCanvasElement.captureStream".
func (this HTMLCanvasElement) CaptureStream1() (ret MediaStream) {
	bindings.CallHTMLCanvasElementCaptureStream1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCaptureStream1 calls the method "HTMLCanvasElement.captureStream"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLCanvasElement) TryCaptureStream1() (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLCanvasElementCaptureStream1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGetSupportedFormats returns true if the static method "BarcodeDetector.getSupportedFormats" exists.
func (this BarcodeDetector) HasFuncGetSupportedFormats() bool {
	return js.True == bindings.HasFuncBarcodeDetectorGetSupportedFormats(
		this.ref,
	)
}

// FuncGetSupportedFormats returns the static method "BarcodeDetector.getSupportedFormats".
func (this BarcodeDetector) FuncGetSupportedFormats() (fn js.Func[func() js.Promise[js.Array[BarcodeFormat]]]) {
	bindings.FuncBarcodeDetectorGetSupportedFormats(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSupportedFormats calls the static method "BarcodeDetector.getSupportedFormats".
func (this BarcodeDetector) GetSupportedFormats() (ret js.Promise[js.Array[BarcodeFormat]]) {
	bindings.CallBarcodeDetectorGetSupportedFormats(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSupportedFormats calls the static method "BarcodeDetector.getSupportedFormats"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BarcodeDetector) TryGetSupportedFormats() (ret js.Promise[js.Array[BarcodeFormat]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBarcodeDetectorGetSupportedFormats(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDetect returns true if the method "BarcodeDetector.detect" exists.
func (this BarcodeDetector) HasFuncDetect() bool {
	return js.True == bindings.HasFuncBarcodeDetectorDetect(
		this.ref,
	)
}

// FuncDetect returns the method "BarcodeDetector.detect".
func (this BarcodeDetector) FuncDetect() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedBarcode]]]) {
	bindings.FuncBarcodeDetectorDetect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Detect calls the method "BarcodeDetector.detect".
func (this BarcodeDetector) Detect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedBarcode]]) {
	bindings.CallBarcodeDetectorDetect(
		this.ref, js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryDetect calls the method "BarcodeDetector.detect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BarcodeDetector) TryDetect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedBarcode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBarcodeDetectorDetect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BaseComputedKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BaseComputedKeyframeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BaseComputedKeyframe) Update(ref js.Ref) {
	bindings.BaseComputedKeyframeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BaseComputedKeyframe) FreeMembers(recursive bool) {
	js.Free(
		p.Easing.Ref(),
	)
	p.Easing = p.Easing.FromRef(js.Undefined)
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
func (p *BaseKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BaseKeyframeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BaseKeyframe) Update(ref js.Ref) {
	bindings.BaseKeyframeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BaseKeyframe) FreeMembers(recursive bool) {
	js.Free(
		p.Easing.Ref(),
	)
	p.Easing = p.Easing.FromRef(js.Undefined)
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
func (p *BasePropertyIndexedKeyframe) UpdateFrom(ref js.Ref) {
	bindings.BasePropertyIndexedKeyframeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BasePropertyIndexedKeyframe) Update(ref js.Ref) {
	bindings.BasePropertyIndexedKeyframeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BasePropertyIndexedKeyframe) FreeMembers(recursive bool) {
	js.Free(
		p.Offset.Ref(),
		p.Easing.Ref(),
		p.Composite.Ref(),
	)
	p.Offset = p.Offset.FromRef(js.Undefined)
	p.Easing = p.Easing.FromRef(js.Undefined)
	p.Composite = p.Composite.FromRef(js.Undefined)
}

type BatteryManager struct {
	EventTarget
}

func (this BatteryManager) Once() BatteryManager {
	this.ref.Once()
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
	this.ref.Free()
}

// Charging returns the value of property "BatteryManager.charging".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) Charging() (ret bool, ok bool) {
	ok = js.True == bindings.GetBatteryManagerCharging(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChargingTime returns the value of property "BatteryManager.chargingTime".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) ChargingTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerChargingTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DischargingTime returns the value of property "BatteryManager.dischargingTime".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) DischargingTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerDischargingTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Level returns the value of property "BatteryManager.level".
//
// It returns ok=false if there is no such property.
func (this BatteryManager) Level() (ret float64, ok bool) {
	ok = js.True == bindings.GetBatteryManagerLevel(
		this.ref, js.Pointer(&ret),
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
func (p *PromptResponseObject) UpdateFrom(ref js.Ref) {
	bindings.PromptResponseObjectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PromptResponseObject) Update(ref js.Ref) {
	bindings.PromptResponseObjectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PromptResponseObject) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncPrompt returns true if the method "BeforeInstallPromptEvent.prompt" exists.
func (this BeforeInstallPromptEvent) HasFuncPrompt() bool {
	return js.True == bindings.HasFuncBeforeInstallPromptEventPrompt(
		this.ref,
	)
}

// FuncPrompt returns the method "BeforeInstallPromptEvent.prompt".
func (this BeforeInstallPromptEvent) FuncPrompt() (fn js.Func[func() js.Promise[PromptResponseObject]]) {
	bindings.FuncBeforeInstallPromptEventPrompt(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prompt calls the method "BeforeInstallPromptEvent.prompt".
func (this BeforeInstallPromptEvent) Prompt() (ret js.Promise[PromptResponseObject]) {
	bindings.CallBeforeInstallPromptEventPrompt(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPrompt calls the method "BeforeInstallPromptEvent.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BeforeInstallPromptEvent) TryPrompt() (ret js.Promise[PromptResponseObject], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBeforeInstallPromptEventPrompt(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// ReturnValue returns the value of property "BeforeUnloadEvent.returnValue".
//
// It returns ok=false if there is no such property.
func (this BeforeUnloadEvent) ReturnValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBeforeUnloadEventReturnValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReturnValue sets the value of property "BeforeUnloadEvent.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this BeforeUnloadEvent) SetReturnValue(val js.String) bool {
	return js.True == bindings.SetBeforeUnloadEventReturnValue(
		this.ref,
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
func (p *PreviousWin) UpdateFrom(ref js.Ref) {
	bindings.PreviousWinJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PreviousWin) Update(ref js.Ref) {
	bindings.PreviousWinJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PreviousWin) FreeMembers(recursive bool) {
	js.Free(
		p.AdJSON.Ref(),
	)
	p.AdJSON = p.AdJSON.FromRef(js.Undefined)
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
func (p *BiddingBrowserSignals) UpdateFrom(ref js.Ref) {
	bindings.BiddingBrowserSignalsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BiddingBrowserSignals) Update(ref js.Ref) {
	bindings.BiddingBrowserSignalsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BiddingBrowserSignals) FreeMembers(recursive bool) {
	js.Free(
		p.TopWindowHostname.Ref(),
		p.Seller.Ref(),
		p.TopLevelSeller.Ref(),
		p.PrevWinsMs.Ref(),
		p.WasmHelper.Ref(),
	)
	p.TopWindowHostname = p.TopWindowHostname.FromRef(js.Undefined)
	p.Seller = p.Seller.FromRef(js.Undefined)
	p.TopLevelSeller = p.TopLevelSeller.FromRef(js.Undefined)
	p.PrevWinsMs = p.PrevWinsMs.FromRef(js.Undefined)
	p.WasmHelper = p.WasmHelper.FromRef(js.Undefined)
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
func (p *BlobEventInit) UpdateFrom(ref js.Ref) {
	bindings.BlobEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BlobEventInit) Update(ref js.Ref) {
	bindings.BlobEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BlobEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "BlobEvent.data".
//
// It returns ok=false if there is no such property.
func (this BlobEvent) Data() (ret Blob, ok bool) {
	ok = js.True == bindings.GetBlobEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timecode returns the value of property "BlobEvent.timecode".
//
// It returns ok=false if there is no such property.
func (this BlobEvent) Timecode() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetBlobEventTimecode(
		this.ref, js.Pointer(&ret),
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
func (p *WatchAdvertisementsOptions) UpdateFrom(ref js.Ref) {
	bindings.WatchAdvertisementsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WatchAdvertisementsOptions) Update(ref js.Ref) {
	bindings.WatchAdvertisementsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WatchAdvertisementsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
}

type BluetoothRemoteGATTDescriptor struct {
	ref js.Ref
}

func (this BluetoothRemoteGATTDescriptor) Once() BluetoothRemoteGATTDescriptor {
	this.ref.Once()
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
	this.ref.Free()
}

// Characteristic returns the value of property "BluetoothRemoteGATTDescriptor.characteristic".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Characteristic() (ret BluetoothRemoteGATTCharacteristic, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorCharacteristic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTDescriptor.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorUuid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "BluetoothRemoteGATTDescriptor.value".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTDescriptor) Value() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTDescriptorValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncReadValue returns true if the method "BluetoothRemoteGATTDescriptor.readValue" exists.
func (this BluetoothRemoteGATTDescriptor) HasFuncReadValue() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTDescriptorReadValue(
		this.ref,
	)
}

// FuncReadValue returns the method "BluetoothRemoteGATTDescriptor.readValue".
func (this BluetoothRemoteGATTDescriptor) FuncReadValue() (fn js.Func[func() js.Promise[js.DataView]]) {
	bindings.FuncBluetoothRemoteGATTDescriptorReadValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadValue calls the method "BluetoothRemoteGATTDescriptor.readValue".
func (this BluetoothRemoteGATTDescriptor) ReadValue() (ret js.Promise[js.DataView]) {
	bindings.CallBluetoothRemoteGATTDescriptorReadValue(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReadValue calls the method "BluetoothRemoteGATTDescriptor.readValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTDescriptor) TryReadValue() (ret js.Promise[js.DataView], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTDescriptorReadValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWriteValue returns true if the method "BluetoothRemoteGATTDescriptor.writeValue" exists.
func (this BluetoothRemoteGATTDescriptor) HasFuncWriteValue() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTDescriptorWriteValue(
		this.ref,
	)
}

// FuncWriteValue returns the method "BluetoothRemoteGATTDescriptor.writeValue".
func (this BluetoothRemoteGATTDescriptor) FuncWriteValue() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	bindings.FuncBluetoothRemoteGATTDescriptorWriteValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteValue calls the method "BluetoothRemoteGATTDescriptor.writeValue".
func (this BluetoothRemoteGATTDescriptor) WriteValue(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTDescriptorWriteValue(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValue calls the method "BluetoothRemoteGATTDescriptor.writeValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTDescriptor) TryWriteValue(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTDescriptorWriteValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Broadcast returns the value of property "BluetoothCharacteristicProperties.broadcast".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Broadcast() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesBroadcast(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Read returns the value of property "BluetoothCharacteristicProperties.read".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Read() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesRead(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WriteWithoutResponse returns the value of property "BluetoothCharacteristicProperties.writeWithoutResponse".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) WriteWithoutResponse() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWriteWithoutResponse(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Write returns the value of property "BluetoothCharacteristicProperties.write".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Write() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWrite(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Notify returns the value of property "BluetoothCharacteristicProperties.notify".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Notify() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesNotify(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Indicate returns the value of property "BluetoothCharacteristicProperties.indicate".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) Indicate() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesIndicate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AuthenticatedSignedWrites returns the value of property "BluetoothCharacteristicProperties.authenticatedSignedWrites".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) AuthenticatedSignedWrites() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesAuthenticatedSignedWrites(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReliableWrite returns the value of property "BluetoothCharacteristicProperties.reliableWrite".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) ReliableWrite() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesReliableWrite(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WritableAuxiliaries returns the value of property "BluetoothCharacteristicProperties.writableAuxiliaries".
//
// It returns ok=false if there is no such property.
func (this BluetoothCharacteristicProperties) WritableAuxiliaries() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothCharacteristicPropertiesWritableAuxiliaries(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BluetoothRemoteGATTCharacteristic struct {
	EventTarget
}

func (this BluetoothRemoteGATTCharacteristic) Once() BluetoothRemoteGATTCharacteristic {
	this.ref.Once()
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
	this.ref.Free()
}

// Service returns the value of property "BluetoothRemoteGATTCharacteristic.service".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Service() (ret BluetoothRemoteGATTService, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicService(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTCharacteristic.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicUuid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Properties returns the value of property "BluetoothRemoteGATTCharacteristic.properties".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Properties() (ret BluetoothCharacteristicProperties, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicProperties(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "BluetoothRemoteGATTCharacteristic.value".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTCharacteristic) Value() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTCharacteristicValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetDescriptor returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptor" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncGetDescriptor() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.ref,
	)
}

// FuncGetDescriptor returns the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
func (this BluetoothRemoteGATTCharacteristic) FuncGetDescriptor() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[BluetoothRemoteGATTDescriptor]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDescriptor calls the method "BluetoothRemoteGATTCharacteristic.getDescriptor".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptor(descriptor BluetoothDescriptorUUID) (ret js.Promise[BluetoothRemoteGATTDescriptor]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.ref, js.Pointer(&ret),
		descriptor.Ref(),
	)

	return
}

// TryGetDescriptor calls the method "BluetoothRemoteGATTCharacteristic.getDescriptor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptor(descriptor BluetoothDescriptorUUID) (ret js.Promise[BluetoothRemoteGATTDescriptor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		descriptor.Ref(),
	)

	return
}

// HasFuncGetDescriptors returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptors" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncGetDescriptors() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.ref,
	)
}

// FuncGetDescriptors returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) FuncGetDescriptors() (fn js.Func[func(descriptor BluetoothDescriptorUUID) js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDescriptors calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors(descriptor BluetoothDescriptorUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.ref, js.Pointer(&ret),
		descriptor.Ref(),
	)

	return
}

// TryGetDescriptors calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptors(descriptor BluetoothDescriptorUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptors(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		descriptor.Ref(),
	)

	return
}

// HasFuncGetDescriptors1 returns true if the method "BluetoothRemoteGATTCharacteristic.getDescriptors" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncGetDescriptors1() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.ref,
	)
}

// FuncGetDescriptors1 returns the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) FuncGetDescriptors1() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDescriptors1 calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors".
func (this BluetoothRemoteGATTCharacteristic) GetDescriptors1() (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]]) {
	bindings.CallBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetDescriptors1 calls the method "BluetoothRemoteGATTCharacteristic.getDescriptors"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryGetDescriptors1() (ret js.Promise[js.Array[BluetoothRemoteGATTDescriptor]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicGetDescriptors1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReadValue returns true if the method "BluetoothRemoteGATTCharacteristic.readValue" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncReadValue() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicReadValue(
		this.ref,
	)
}

// FuncReadValue returns the method "BluetoothRemoteGATTCharacteristic.readValue".
func (this BluetoothRemoteGATTCharacteristic) FuncReadValue() (fn js.Func[func() js.Promise[js.DataView]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicReadValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadValue calls the method "BluetoothRemoteGATTCharacteristic.readValue".
func (this BluetoothRemoteGATTCharacteristic) ReadValue() (ret js.Promise[js.DataView]) {
	bindings.CallBluetoothRemoteGATTCharacteristicReadValue(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReadValue calls the method "BluetoothRemoteGATTCharacteristic.readValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryReadValue() (ret js.Promise[js.DataView], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicReadValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWriteValue returns true if the method "BluetoothRemoteGATTCharacteristic.writeValue" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncWriteValue() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicWriteValue(
		this.ref,
	)
}

// FuncWriteValue returns the method "BluetoothRemoteGATTCharacteristic.writeValue".
func (this BluetoothRemoteGATTCharacteristic) FuncWriteValue() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicWriteValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteValue calls the method "BluetoothRemoteGATTCharacteristic.writeValue".
func (this BluetoothRemoteGATTCharacteristic) WriteValue(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValue(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValue calls the method "BluetoothRemoteGATTCharacteristic.writeValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValue(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncWriteValueWithResponse returns true if the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncWriteValueWithResponse() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.ref,
	)
}

// FuncWriteValueWithResponse returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
func (this BluetoothRemoteGATTCharacteristic) FuncWriteValueWithResponse() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteValueWithResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithResponse(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValueWithResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValueWithResponse(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValueWithResponse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncWriteValueWithoutResponse returns true if the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncWriteValueWithoutResponse() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.ref,
	)
}

// FuncWriteValueWithoutResponse returns the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
func (this BluetoothRemoteGATTCharacteristic) FuncWriteValueWithoutResponse() (fn js.Func[func(value BufferSource) js.Promise[js.Void]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteValueWithoutResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse".
func (this BluetoothRemoteGATTCharacteristic) WriteValueWithoutResponse(value BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryWriteValueWithoutResponse calls the method "BluetoothRemoteGATTCharacteristic.writeValueWithoutResponse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryWriteValueWithoutResponse(value BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicWriteValueWithoutResponse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncStartNotifications returns true if the method "BluetoothRemoteGATTCharacteristic.startNotifications" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncStartNotifications() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicStartNotifications(
		this.ref,
	)
}

// FuncStartNotifications returns the method "BluetoothRemoteGATTCharacteristic.startNotifications".
func (this BluetoothRemoteGATTCharacteristic) FuncStartNotifications() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicStartNotifications(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StartNotifications calls the method "BluetoothRemoteGATTCharacteristic.startNotifications".
func (this BluetoothRemoteGATTCharacteristic) StartNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTCharacteristicStartNotifications(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStartNotifications calls the method "BluetoothRemoteGATTCharacteristic.startNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryStartNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicStartNotifications(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopNotifications returns true if the method "BluetoothRemoteGATTCharacteristic.stopNotifications" exists.
func (this BluetoothRemoteGATTCharacteristic) HasFuncStopNotifications() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTCharacteristicStopNotifications(
		this.ref,
	)
}

// FuncStopNotifications returns the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
func (this BluetoothRemoteGATTCharacteristic) FuncStopNotifications() (fn js.Func[func() js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	bindings.FuncBluetoothRemoteGATTCharacteristicStopNotifications(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StopNotifications calls the method "BluetoothRemoteGATTCharacteristic.stopNotifications".
func (this BluetoothRemoteGATTCharacteristic) StopNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTCharacteristicStopNotifications(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStopNotifications calls the method "BluetoothRemoteGATTCharacteristic.stopNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTCharacteristic) TryStopNotifications() (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTCharacteristicStopNotifications(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BluetoothCharacteristicUUID = OneOf_String_Uint32

type BluetoothServiceUUID = OneOf_String_Uint32

type BluetoothRemoteGATTService struct {
	EventTarget
}

func (this BluetoothRemoteGATTService) Once() BluetoothRemoteGATTService {
	this.ref.Once()
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
	this.ref.Free()
}

// Device returns the value of property "BluetoothRemoteGATTService.device".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Uuid returns the value of property "BluetoothRemoteGATTService.uuid".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) Uuid() (ret UUID, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceUuid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsPrimary returns the value of property "BluetoothRemoteGATTService.isPrimary".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTService) IsPrimary() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServiceIsPrimary(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetCharacteristic returns true if the method "BluetoothRemoteGATTService.getCharacteristic" exists.
func (this BluetoothRemoteGATTService) HasFuncGetCharacteristic() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetCharacteristic(
		this.ref,
	)
}

// FuncGetCharacteristic returns the method "BluetoothRemoteGATTService.getCharacteristic".
func (this BluetoothRemoteGATTService) FuncGetCharacteristic() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[BluetoothRemoteGATTCharacteristic]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetCharacteristic(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCharacteristic calls the method "BluetoothRemoteGATTService.getCharacteristic".
func (this BluetoothRemoteGATTService) GetCharacteristic(characteristic BluetoothCharacteristicUUID) (ret js.Promise[BluetoothRemoteGATTCharacteristic]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristic(
		this.ref, js.Pointer(&ret),
		characteristic.Ref(),
	)

	return
}

// TryGetCharacteristic calls the method "BluetoothRemoteGATTService.getCharacteristic"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristic(characteristic BluetoothCharacteristicUUID) (ret js.Promise[BluetoothRemoteGATTCharacteristic], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristic(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		characteristic.Ref(),
	)

	return
}

// HasFuncGetCharacteristics returns true if the method "BluetoothRemoteGATTService.getCharacteristics" exists.
func (this BluetoothRemoteGATTService) HasFuncGetCharacteristics() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetCharacteristics(
		this.ref,
	)
}

// FuncGetCharacteristics returns the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) FuncGetCharacteristics() (fn js.Func[func(characteristic BluetoothCharacteristicUUID) js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetCharacteristics(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCharacteristics calls the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristics(characteristic BluetoothCharacteristicUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristics(
		this.ref, js.Pointer(&ret),
		characteristic.Ref(),
	)

	return
}

// TryGetCharacteristics calls the method "BluetoothRemoteGATTService.getCharacteristics"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristics(characteristic BluetoothCharacteristicUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristics(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		characteristic.Ref(),
	)

	return
}

// HasFuncGetCharacteristics1 returns true if the method "BluetoothRemoteGATTService.getCharacteristics" exists.
func (this BluetoothRemoteGATTService) HasFuncGetCharacteristics1() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetCharacteristics1(
		this.ref,
	)
}

// FuncGetCharacteristics1 returns the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) FuncGetCharacteristics1() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetCharacteristics1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCharacteristics1 calls the method "BluetoothRemoteGATTService.getCharacteristics".
func (this BluetoothRemoteGATTService) GetCharacteristics1() (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]]) {
	bindings.CallBluetoothRemoteGATTServiceGetCharacteristics1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCharacteristics1 calls the method "BluetoothRemoteGATTService.getCharacteristics"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetCharacteristics1() (ret js.Promise[js.Array[BluetoothRemoteGATTCharacteristic]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetCharacteristics1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetIncludedService returns true if the method "BluetoothRemoteGATTService.getIncludedService" exists.
func (this BluetoothRemoteGATTService) HasFuncGetIncludedService() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetIncludedService(
		this.ref,
	)
}

// FuncGetIncludedService returns the method "BluetoothRemoteGATTService.getIncludedService".
func (this BluetoothRemoteGATTService) FuncGetIncludedService() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetIncludedService(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIncludedService calls the method "BluetoothRemoteGATTService.getIncludedService".
func (this BluetoothRemoteGATTService) GetIncludedService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedService(
		this.ref, js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetIncludedService calls the method "BluetoothRemoteGATTService.getIncludedService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedService(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasFuncGetIncludedServices returns true if the method "BluetoothRemoteGATTService.getIncludedServices" exists.
func (this BluetoothRemoteGATTService) HasFuncGetIncludedServices() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetIncludedServices(
		this.ref,
	)
}

// FuncGetIncludedServices returns the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) FuncGetIncludedServices() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetIncludedServices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIncludedServices calls the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedServices(
		this.ref, js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetIncludedServices calls the method "BluetoothRemoteGATTService.getIncludedServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedServices(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasFuncGetIncludedServices1 returns true if the method "BluetoothRemoteGATTService.getIncludedServices" exists.
func (this BluetoothRemoteGATTService) HasFuncGetIncludedServices1() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServiceGetIncludedServices1(
		this.ref,
	)
}

// FuncGetIncludedServices1 returns the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) FuncGetIncludedServices1() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	bindings.FuncBluetoothRemoteGATTServiceGetIncludedServices1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIncludedServices1 calls the method "BluetoothRemoteGATTService.getIncludedServices".
func (this BluetoothRemoteGATTService) GetIncludedServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServiceGetIncludedServices1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetIncludedServices1 calls the method "BluetoothRemoteGATTService.getIncludedServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTService) TryGetIncludedServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServiceGetIncludedServices1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BluetoothRemoteGATTServer struct {
	ref js.Ref
}

func (this BluetoothRemoteGATTServer) Once() BluetoothRemoteGATTServer {
	this.ref.Once()
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
	this.ref.Free()
}

// Device returns the value of property "BluetoothRemoteGATTServer.device".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTServer) Device() (ret BluetoothDevice, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServerDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Connected returns the value of property "BluetoothRemoteGATTServer.connected".
//
// It returns ok=false if there is no such property.
func (this BluetoothRemoteGATTServer) Connected() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothRemoteGATTServerConnected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncConnect returns true if the method "BluetoothRemoteGATTServer.connect" exists.
func (this BluetoothRemoteGATTServer) HasFuncConnect() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServerConnect(
		this.ref,
	)
}

// FuncConnect returns the method "BluetoothRemoteGATTServer.connect".
func (this BluetoothRemoteGATTServer) FuncConnect() (fn js.Func[func() js.Promise[BluetoothRemoteGATTServer]]) {
	bindings.FuncBluetoothRemoteGATTServerConnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect calls the method "BluetoothRemoteGATTServer.connect".
func (this BluetoothRemoteGATTServer) Connect() (ret js.Promise[BluetoothRemoteGATTServer]) {
	bindings.CallBluetoothRemoteGATTServerConnect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryConnect calls the method "BluetoothRemoteGATTServer.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTServer) TryConnect() (ret js.Promise[BluetoothRemoteGATTServer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerConnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDisconnect returns true if the method "BluetoothRemoteGATTServer.disconnect" exists.
func (this BluetoothRemoteGATTServer) HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServerDisconnect(
		this.ref,
	)
}

// FuncDisconnect returns the method "BluetoothRemoteGATTServer.disconnect".
func (this BluetoothRemoteGATTServer) FuncDisconnect() (fn js.Func[func()]) {
	bindings.FuncBluetoothRemoteGATTServerDisconnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect calls the method "BluetoothRemoteGATTServer.disconnect".
func (this BluetoothRemoteGATTServer) Disconnect() (ret js.Void) {
	bindings.CallBluetoothRemoteGATTServerDisconnect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "BluetoothRemoteGATTServer.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTServer) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerDisconnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPrimaryService returns true if the method "BluetoothRemoteGATTServer.getPrimaryService" exists.
func (this BluetoothRemoteGATTServer) HasFuncGetPrimaryService() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServerGetPrimaryService(
		this.ref,
	)
}

// FuncGetPrimaryService returns the method "BluetoothRemoteGATTServer.getPrimaryService".
func (this BluetoothRemoteGATTServer) FuncGetPrimaryService() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[BluetoothRemoteGATTService]]) {
	bindings.FuncBluetoothRemoteGATTServerGetPrimaryService(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPrimaryService calls the method "BluetoothRemoteGATTServer.getPrimaryService".
func (this BluetoothRemoteGATTServer) GetPrimaryService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryService(
		this.ref, js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetPrimaryService calls the method "BluetoothRemoteGATTServer.getPrimaryService"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryService(service BluetoothServiceUUID) (ret js.Promise[BluetoothRemoteGATTService], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryService(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasFuncGetPrimaryServices returns true if the method "BluetoothRemoteGATTServer.getPrimaryServices" exists.
func (this BluetoothRemoteGATTServer) HasFuncGetPrimaryServices() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServerGetPrimaryServices(
		this.ref,
	)
}

// FuncGetPrimaryServices returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) FuncGetPrimaryServices() (fn js.Func[func(service BluetoothServiceUUID) js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	bindings.FuncBluetoothRemoteGATTServerGetPrimaryServices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPrimaryServices calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryServices(
		this.ref, js.Pointer(&ret),
		service.Ref(),
	)

	return
}

// TryGetPrimaryServices calls the method "BluetoothRemoteGATTServer.getPrimaryServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryServices(service BluetoothServiceUUID) (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryServices(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		service.Ref(),
	)

	return
}

// HasFuncGetPrimaryServices1 returns true if the method "BluetoothRemoteGATTServer.getPrimaryServices" exists.
func (this BluetoothRemoteGATTServer) HasFuncGetPrimaryServices1() bool {
	return js.True == bindings.HasFuncBluetoothRemoteGATTServerGetPrimaryServices1(
		this.ref,
	)
}

// FuncGetPrimaryServices1 returns the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) FuncGetPrimaryServices1() (fn js.Func[func() js.Promise[js.Array[BluetoothRemoteGATTService]]]) {
	bindings.FuncBluetoothRemoteGATTServerGetPrimaryServices1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPrimaryServices1 calls the method "BluetoothRemoteGATTServer.getPrimaryServices".
func (this BluetoothRemoteGATTServer) GetPrimaryServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]]) {
	bindings.CallBluetoothRemoteGATTServerGetPrimaryServices1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetPrimaryServices1 calls the method "BluetoothRemoteGATTServer.getPrimaryServices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothRemoteGATTServer) TryGetPrimaryServices1() (ret js.Promise[js.Array[BluetoothRemoteGATTService]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothRemoteGATTServerGetPrimaryServices1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type BluetoothDevice struct {
	EventTarget
}

func (this BluetoothDevice) Once() BluetoothDevice {
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "BluetoothDevice.id".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "BluetoothDevice.name".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gatt returns the value of property "BluetoothDevice.gatt".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) Gatt() (ret BluetoothRemoteGATTServer, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceGatt(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WatchingAdvertisements returns the value of property "BluetoothDevice.watchingAdvertisements".
//
// It returns ok=false if there is no such property.
func (this BluetoothDevice) WatchingAdvertisements() (ret bool, ok bool) {
	ok = js.True == bindings.GetBluetoothDeviceWatchingAdvertisements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncForget returns true if the method "BluetoothDevice.forget" exists.
func (this BluetoothDevice) HasFuncForget() bool {
	return js.True == bindings.HasFuncBluetoothDeviceForget(
		this.ref,
	)
}

// FuncForget returns the method "BluetoothDevice.forget".
func (this BluetoothDevice) FuncForget() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncBluetoothDeviceForget(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forget calls the method "BluetoothDevice.forget".
func (this BluetoothDevice) Forget() (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceForget(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForget calls the method "BluetoothDevice.forget"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothDevice) TryForget() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceForget(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWatchAdvertisements returns true if the method "BluetoothDevice.watchAdvertisements" exists.
func (this BluetoothDevice) HasFuncWatchAdvertisements() bool {
	return js.True == bindings.HasFuncBluetoothDeviceWatchAdvertisements(
		this.ref,
	)
}

// FuncWatchAdvertisements returns the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) FuncWatchAdvertisements() (fn js.Func[func(options WatchAdvertisementsOptions) js.Promise[js.Void]]) {
	bindings.FuncBluetoothDeviceWatchAdvertisements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WatchAdvertisements calls the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisements(options WatchAdvertisementsOptions) (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceWatchAdvertisements(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryWatchAdvertisements calls the method "BluetoothDevice.watchAdvertisements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothDevice) TryWatchAdvertisements(options WatchAdvertisementsOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceWatchAdvertisements(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncWatchAdvertisements1 returns true if the method "BluetoothDevice.watchAdvertisements" exists.
func (this BluetoothDevice) HasFuncWatchAdvertisements1() bool {
	return js.True == bindings.HasFuncBluetoothDeviceWatchAdvertisements1(
		this.ref,
	)
}

// FuncWatchAdvertisements1 returns the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) FuncWatchAdvertisements1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncBluetoothDeviceWatchAdvertisements1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WatchAdvertisements1 calls the method "BluetoothDevice.watchAdvertisements".
func (this BluetoothDevice) WatchAdvertisements1() (ret js.Promise[js.Void]) {
	bindings.CallBluetoothDeviceWatchAdvertisements1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryWatchAdvertisements1 calls the method "BluetoothDevice.watchAdvertisements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BluetoothDevice) TryWatchAdvertisements1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBluetoothDeviceWatchAdvertisements1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *BluetoothManufacturerDataFilterInit) UpdateFrom(ref js.Ref) {
	bindings.BluetoothManufacturerDataFilterInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BluetoothManufacturerDataFilterInit) Update(ref js.Ref) {
	bindings.BluetoothManufacturerDataFilterInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BluetoothManufacturerDataFilterInit) FreeMembers(recursive bool) {
	js.Free(
		p.DataPrefix.Ref(),
		p.Mask.Ref(),
	)
	p.DataPrefix = p.DataPrefix.FromRef(js.Undefined)
	p.Mask = p.Mask.FromRef(js.Undefined)
}
