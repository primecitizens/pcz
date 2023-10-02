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

type VideoTransferCharacteristics uint32

const (
	_ VideoTransferCharacteristics = iota

	VideoTransferCharacteristics_BT709
	VideoTransferCharacteristics_SMPTE_170M
	VideoTransferCharacteristics_IEC61966_2_1
	VideoTransferCharacteristics_LINEAR
	VideoTransferCharacteristics_PQ
	VideoTransferCharacteristics_HLG
)

func (VideoTransferCharacteristics) FromRef(str js.Ref) VideoTransferCharacteristics {
	return VideoTransferCharacteristics(bindings.ConstOfVideoTransferCharacteristics(str))
}

func (x VideoTransferCharacteristics) String() (string, bool) {
	switch x {
	case VideoTransferCharacteristics_BT709:
		return "bt709", true
	case VideoTransferCharacteristics_SMPTE_170M:
		return "smpte170m", true
	case VideoTransferCharacteristics_IEC61966_2_1:
		return "iec61966-2-1", true
	case VideoTransferCharacteristics_LINEAR:
		return "linear", true
	case VideoTransferCharacteristics_PQ:
		return "pq", true
	case VideoTransferCharacteristics_HLG:
		return "hlg", true
	default:
		return "", false
	}
}

type VideoMatrixCoefficients uint32

const (
	_ VideoMatrixCoefficients = iota

	VideoMatrixCoefficients_RGB
	VideoMatrixCoefficients_BT709
	VideoMatrixCoefficients_BT_470BG
	VideoMatrixCoefficients_SMPTE_170M
	VideoMatrixCoefficients_BT2020_NCL
)

func (VideoMatrixCoefficients) FromRef(str js.Ref) VideoMatrixCoefficients {
	return VideoMatrixCoefficients(bindings.ConstOfVideoMatrixCoefficients(str))
}

func (x VideoMatrixCoefficients) String() (string, bool) {
	switch x {
	case VideoMatrixCoefficients_RGB:
		return "rgb", true
	case VideoMatrixCoefficients_BT709:
		return "bt709", true
	case VideoMatrixCoefficients_BT_470BG:
		return "bt470bg", true
	case VideoMatrixCoefficients_SMPTE_170M:
		return "smpte170m", true
	case VideoMatrixCoefficients_BT2020_NCL:
		return "bt2020-ncl", true
	default:
		return "", false
	}
}

type VideoColorSpaceInit struct {
	// Primaries is "VideoColorSpaceInit.primaries"
	//
	// Optional, defaults to null.
	Primaries VideoColorPrimaries
	// Transfer is "VideoColorSpaceInit.transfer"
	//
	// Optional, defaults to null.
	Transfer VideoTransferCharacteristics
	// Matrix is "VideoColorSpaceInit.matrix"
	//
	// Optional, defaults to null.
	Matrix VideoMatrixCoefficients
	// FullRange is "VideoColorSpaceInit.fullRange"
	//
	// Optional, defaults to null.
	FullRange bool

	FFI_USE_FullRange bool // for FullRange.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoColorSpaceInit with all fields set.
func (p VideoColorSpaceInit) FromRef(ref js.Ref) VideoColorSpaceInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 VideoColorSpaceInit VideoColorSpaceInit [// VideoColorSpaceInit] [0x14000312c80 0x14000312d20 0x14000312dc0 0x14000312e60 0x14000312f00] 0x14000baa630 {0 0}} in the application heap.
func (p VideoColorSpaceInit) New() js.Ref {
	return bindings.VideoColorSpaceInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoColorSpaceInit) UpdateFrom(ref js.Ref) {
	bindings.VideoColorSpaceInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoColorSpaceInit) Update(ref js.Ref) {
	bindings.VideoColorSpaceInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoFrameBufferInit struct {
	// Format is "VideoFrameBufferInit.format"
	//
	// Required
	Format VideoPixelFormat
	// CodedWidth is "VideoFrameBufferInit.codedWidth"
	//
	// Required
	CodedWidth uint32
	// CodedHeight is "VideoFrameBufferInit.codedHeight"
	//
	// Required
	CodedHeight uint32
	// Timestamp is "VideoFrameBufferInit.timestamp"
	//
	// Required
	Timestamp int64
	// Duration is "VideoFrameBufferInit.duration"
	//
	// Optional
	Duration uint64
	// Layout is "VideoFrameBufferInit.layout"
	//
	// Optional
	Layout js.Array[PlaneLayout]
	// VisibleRect is "VideoFrameBufferInit.visibleRect"
	//
	// Optional
	VisibleRect DOMRectInit
	// DisplayWidth is "VideoFrameBufferInit.displayWidth"
	//
	// Optional
	DisplayWidth uint32
	// DisplayHeight is "VideoFrameBufferInit.displayHeight"
	//
	// Optional
	DisplayHeight uint32
	// ColorSpace is "VideoFrameBufferInit.colorSpace"
	//
	// Optional
	ColorSpace VideoColorSpaceInit
	// Transfer is "VideoFrameBufferInit.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.ArrayBuffer]

	FFI_USE_Duration      bool // for Duration.
	FFI_USE_DisplayWidth  bool // for DisplayWidth.
	FFI_USE_DisplayHeight bool // for DisplayHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoFrameBufferInit with all fields set.
func (p VideoFrameBufferInit) FromRef(ref js.Ref) VideoFrameBufferInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 VideoFrameBufferInit VideoFrameBufferInit [// VideoFrameBufferInit] [0x140003123c0 0x14000312460 0x14000312500 0x140003125a0 0x14000312640 0x140003128c0 0x14000312960 0x14000312a00 0x14000312b40 0x14000312fa0 0x14000313040 0x140003126e0 0x14000312aa0 0x14000312be0] 0x14000baa108 {0 0}} in the application heap.
func (p VideoFrameBufferInit) New() js.Ref {
	return bindings.VideoFrameBufferInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoFrameBufferInit) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameBufferInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoFrameBufferInit) Update(ref js.Ref) {
	bindings.VideoFrameBufferInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoFrameCopyToOptions struct {
	// Rect is "VideoFrameCopyToOptions.rect"
	//
	// Optional
	Rect DOMRectInit
	// Layout is "VideoFrameCopyToOptions.layout"
	//
	// Optional
	Layout js.Array[PlaneLayout]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoFrameCopyToOptions with all fields set.
func (p VideoFrameCopyToOptions) FromRef(ref js.Ref) VideoFrameCopyToOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 VideoFrameCopyToOptions VideoFrameCopyToOptions [// VideoFrameCopyToOptions] [0x140003130e0 0x14000313180] 0x14000baa738 {0 0}} in the application heap.
func (p VideoFrameCopyToOptions) New() js.Ref {
	return bindings.VideoFrameCopyToOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoFrameCopyToOptions) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameCopyToOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoFrameCopyToOptions) Update(ref js.Ref) {
	bindings.VideoFrameCopyToOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewVideoColorSpace(init VideoColorSpaceInit) VideoColorSpace {
	return VideoColorSpace{}.FromRef(
		bindings.NewVideoColorSpaceByVideoColorSpace(
			js.Pointer(&init)),
	)
}

func NewVideoColorSpaceByVideoColorSpace1() VideoColorSpace {
	return VideoColorSpace{}.FromRef(
		bindings.NewVideoColorSpaceByVideoColorSpace1(),
	)
}

type VideoColorSpace struct {
	ref js.Ref
}

func (this VideoColorSpace) Once() VideoColorSpace {
	this.Ref().Once()
	return this
}

func (this VideoColorSpace) Ref() js.Ref {
	return this.ref
}

func (this VideoColorSpace) FromRef(ref js.Ref) VideoColorSpace {
	this.ref = ref
	return this
}

func (this VideoColorSpace) Free() {
	this.Ref().Free()
}

// Primaries returns the value of property "VideoColorSpace.primaries".
//
// The returned bool will be false if there is no such property.
func (this VideoColorSpace) Primaries() (VideoColorPrimaries, bool) {
	var _ok bool
	_ret := bindings.GetVideoColorSpacePrimaries(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoColorPrimaries(_ret), _ok
}

// Transfer returns the value of property "VideoColorSpace.transfer".
//
// The returned bool will be false if there is no such property.
func (this VideoColorSpace) Transfer() (VideoTransferCharacteristics, bool) {
	var _ok bool
	_ret := bindings.GetVideoColorSpaceTransfer(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoTransferCharacteristics(_ret), _ok
}

// Matrix returns the value of property "VideoColorSpace.matrix".
//
// The returned bool will be false if there is no such property.
func (this VideoColorSpace) Matrix() (VideoMatrixCoefficients, bool) {
	var _ok bool
	_ret := bindings.GetVideoColorSpaceMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoMatrixCoefficients(_ret), _ok
}

// FullRange returns the value of property "VideoColorSpace.fullRange".
//
// The returned bool will be false if there is no such property.
func (this VideoColorSpace) FullRange() (bool, bool) {
	var _ok bool
	_ret := bindings.GetVideoColorSpaceFullRange(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ToJSON calls the method "VideoColorSpace.toJSON".
//
// The returned bool will be false if there is no such method.
func (this VideoColorSpace) ToJSON() (VideoColorSpaceInit, bool) {
	var _ret VideoColorSpaceInit
	_ok := js.True == bindings.CallVideoColorSpaceToJSON(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// ToJSONFunc returns the method "VideoColorSpace.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoColorSpace) ToJSONFunc() (fn js.Func[func() VideoColorSpaceInit]) {
	return fn.FromRef(
		bindings.VideoColorSpaceToJSONFunc(
			this.Ref(),
		),
	)
}

func NewVideoFrame(image CanvasImageSource, init VideoFrameInit) VideoFrame {
	return VideoFrame{}.FromRef(
		bindings.NewVideoFrameByVideoFrame(
			image.Ref(),
			js.Pointer(&init)),
	)
}

func NewVideoFrameByVideoFrame1(image CanvasImageSource) VideoFrame {
	return VideoFrame{}.FromRef(
		bindings.NewVideoFrameByVideoFrame1(
			image.Ref()),
	)
}

func NewVideoFrameByVideoFrame2(data AllowSharedBufferSource, init VideoFrameBufferInit) VideoFrame {
	return VideoFrame{}.FromRef(
		bindings.NewVideoFrameByVideoFrame2(
			data.Ref(),
			js.Pointer(&init)),
	)
}

type VideoFrame struct {
	ref js.Ref
}

func (this VideoFrame) Once() VideoFrame {
	this.Ref().Once()
	return this
}

func (this VideoFrame) Ref() js.Ref {
	return this.ref
}

func (this VideoFrame) FromRef(ref js.Ref) VideoFrame {
	this.ref = ref
	return this
}

func (this VideoFrame) Free() {
	this.Ref().Free()
}

// Format returns the value of property "VideoFrame.format".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) Format() (VideoPixelFormat, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameFormat(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoPixelFormat(_ret), _ok
}

// CodedWidth returns the value of property "VideoFrame.codedWidth".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) CodedWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameCodedWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CodedHeight returns the value of property "VideoFrame.codedHeight".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) CodedHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameCodedHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CodedRect returns the value of property "VideoFrame.codedRect".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) CodedRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameCodedRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// VisibleRect returns the value of property "VideoFrame.visibleRect".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) VisibleRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameVisibleRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// DisplayWidth returns the value of property "VideoFrame.displayWidth".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) DisplayWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameDisplayWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// DisplayHeight returns the value of property "VideoFrame.displayHeight".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) DisplayHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameDisplayHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Duration returns the value of property "VideoFrame.duration".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) Duration() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Timestamp returns the value of property "VideoFrame.timestamp".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) Timestamp() (int64, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// ColorSpace returns the value of property "VideoFrame.colorSpace".
//
// The returned bool will be false if there is no such property.
func (this VideoFrame) ColorSpace() (VideoColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetVideoFrameColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoColorSpace{}.FromRef(_ret), _ok
}

// Metadata calls the method "VideoFrame.metadata".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) Metadata() (VideoFrameMetadata, bool) {
	var _ret VideoFrameMetadata
	_ok := js.True == bindings.CallVideoFrameMetadata(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// MetadataFunc returns the method "VideoFrame.metadata".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) MetadataFunc() (fn js.Func[func() VideoFrameMetadata]) {
	return fn.FromRef(
		bindings.VideoFrameMetadataFunc(
			this.Ref(),
		),
	)
}

// AllocationSize calls the method "VideoFrame.allocationSize".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) AllocationSize(options VideoFrameCopyToOptions) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameAllocationSize(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return uint32(_ret), _ok
}

// AllocationSizeFunc returns the method "VideoFrame.allocationSize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) AllocationSizeFunc() (fn js.Func[func(options VideoFrameCopyToOptions) uint32]) {
	return fn.FromRef(
		bindings.VideoFrameAllocationSizeFunc(
			this.Ref(),
		),
	)
}

// AllocationSize1 calls the method "VideoFrame.allocationSize".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) AllocationSize1() (uint32, bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameAllocationSize1(
		this.Ref(), js.Pointer(&_ok),
	)

	return uint32(_ret), _ok
}

// AllocationSize1Func returns the method "VideoFrame.allocationSize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) AllocationSize1Func() (fn js.Func[func() uint32]) {
	return fn.FromRef(
		bindings.VideoFrameAllocationSize1Func(
			this.Ref(),
		),
	)
}

// CopyTo calls the method "VideoFrame.copyTo".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) CopyTo(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) (js.Promise[js.Array[PlaneLayout]], bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameCopyTo(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[PlaneLayout]]{}.FromRef(_ret), _ok
}

// CopyToFunc returns the method "VideoFrame.copyTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) js.Promise[js.Array[PlaneLayout]]]) {
	return fn.FromRef(
		bindings.VideoFrameCopyToFunc(
			this.Ref(),
		),
	)
}

// CopyTo1 calls the method "VideoFrame.copyTo".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) CopyTo1(destination AllowSharedBufferSource) (js.Promise[js.Array[PlaneLayout]], bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameCopyTo1(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
	)

	return js.Promise[js.Array[PlaneLayout]]{}.FromRef(_ret), _ok
}

// CopyTo1Func returns the method "VideoFrame.copyTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) CopyTo1Func() (fn js.Func[func(destination AllowSharedBufferSource) js.Promise[js.Array[PlaneLayout]]]) {
	return fn.FromRef(
		bindings.VideoFrameCopyTo1Func(
			this.Ref(),
		),
	)
}

// Clone calls the method "VideoFrame.clone".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) Clone() (VideoFrame, bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return VideoFrame{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "VideoFrame.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) CloneFunc() (fn js.Func[func() VideoFrame]) {
	return fn.FromRef(
		bindings.VideoFrameCloneFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "VideoFrame.close".
//
// The returned bool will be false if there is no such method.
func (this VideoFrame) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVideoFrameClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "VideoFrame.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoFrame) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoFrameCloseFunc(
			this.Ref(),
		),
	)
}

type OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame struct {
	ref js.Ref
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) Free() {
	x.ref.Free()
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) FromRef(ref js.Ref) OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame {
	return OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame{
		ref: ref,
	}
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) ImageBitmap() ImageBitmap {
	return ImageBitmap{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) ImageData() ImageData {
	return ImageData{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) HTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) HTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) HTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) OffscreenCanvas() OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame) VideoFrame() VideoFrame {
	return VideoFrame{}.FromRef(x.ref)
}

type TexImageSource = OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLCanvasElement_HTMLVideoElement_OffscreenCanvas_VideoFrame

type OneOf_TypedArrayInt32_ArrayGLint struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt32_ArrayGLint) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt32_ArrayGLint) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt32_ArrayGLint) FromRef(ref js.Ref) OneOf_TypedArrayInt32_ArrayGLint {
	return OneOf_TypedArrayInt32_ArrayGLint{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt32_ArrayGLint) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt32_ArrayGLint) ArrayGLint() js.Array[GLint] {
	return js.Array[GLint]{}.FromRef(x.ref)
}

type Int32List = OneOf_TypedArrayInt32_ArrayGLint

type WebGLRenderingContext struct {
	ref js.Ref
}

func (this WebGLRenderingContext) Once() WebGLRenderingContext {
	this.Ref().Once()
	return this
}

func (this WebGLRenderingContext) Ref() js.Ref {
	return this.ref
}

func (this WebGLRenderingContext) FromRef(ref js.Ref) WebGLRenderingContext {
	this.ref = ref
	return this
}

func (this WebGLRenderingContext) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "WebGLRenderingContext.canvas".
//
// The returned bool will be false if there is no such property.
func (this WebGLRenderingContext) Canvas() (OneOf_HTMLCanvasElement_OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.GetWebGLRenderingContextCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_HTMLCanvasElement_OffscreenCanvas{}.FromRef(_ret), _ok
}

// DrawingBufferWidth returns the value of property "WebGLRenderingContext.drawingBufferWidth".
//
// The returned bool will be false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferWidth() (GLsizei, bool) {
	var _ok bool
	_ret := bindings.GetWebGLRenderingContextDrawingBufferWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLsizei(_ret), _ok
}

// DrawingBufferHeight returns the value of property "WebGLRenderingContext.drawingBufferHeight".
//
// The returned bool will be false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferHeight() (GLsizei, bool) {
	var _ok bool
	_ret := bindings.GetWebGLRenderingContextDrawingBufferHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLsizei(_ret), _ok
}

// DrawingBufferColorSpace returns the value of property "WebGLRenderingContext.drawingBufferColorSpace".
//
// The returned bool will be false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferColorSpace() (PredefinedColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetWebGLRenderingContextDrawingBufferColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return PredefinedColorSpace(_ret), _ok
}

// DrawingBufferColorSpace sets the value of property "WebGLRenderingContext.drawingBufferColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGLRenderingContext) SetDrawingBufferColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGLRenderingContextDrawingBufferColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// UnpackColorSpace returns the value of property "WebGLRenderingContext.unpackColorSpace".
//
// The returned bool will be false if there is no such property.
func (this WebGLRenderingContext) UnpackColorSpace() (PredefinedColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetWebGLRenderingContextUnpackColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return PredefinedColorSpace(_ret), _ok
}

// UnpackColorSpace sets the value of property "WebGLRenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGLRenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGLRenderingContextUnpackColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// GetContextAttributes calls the method "WebGLRenderingContext.getContextAttributes".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetContextAttributes() (WebGLContextAttributes, bool) {
	var _ret WebGLContextAttributes
	_ok := js.True == bindings.CallWebGLRenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetContextAttributesFunc returns the method "WebGLRenderingContext.getContextAttributes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetContextAttributesFunc() (fn js.Func[func() WebGLContextAttributes]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "WebGLRenderingContext.isContextLost".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsContextLost() (bool, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsContextLost(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsContextLostFunc returns the method "WebGLRenderingContext.isContextLost".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsContextLostFunc(
			this.Ref(),
		),
	)
}

// GetSupportedExtensions calls the method "WebGLRenderingContext.getSupportedExtensions".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetSupportedExtensions() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// GetSupportedExtensionsFunc returns the method "WebGLRenderingContext.getSupportedExtensions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetSupportedExtensionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetSupportedExtensionsFunc(
			this.Ref(),
		),
	)
}

// GetExtension calls the method "WebGLRenderingContext.getExtension".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetExtension(name js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetExtension(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// GetExtensionFunc returns the method "WebGLRenderingContext.getExtension".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetExtensionFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetExtensionFunc(
			this.Ref(),
		),
	)
}

// ActiveTexture calls the method "WebGLRenderingContext.activeTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ActiveTexture(texture GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextActiveTexture(
		this.Ref(), js.Pointer(&_ok),
		uint32(texture),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ActiveTextureFunc returns the method "WebGLRenderingContext.activeTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ActiveTextureFunc() (fn js.Func[func(texture GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextActiveTextureFunc(
			this.Ref(),
		),
	)
}

// AttachShader calls the method "WebGLRenderingContext.attachShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextAttachShader(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AttachShaderFunc returns the method "WebGLRenderingContext.attachShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) AttachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextAttachShaderFunc(
			this.Ref(),
		),
	)
}

// BindAttribLocation calls the method "WebGLRenderingContext.bindAttribLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindAttribLocationFunc returns the method "WebGLRenderingContext.bindAttribLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BindAttribLocationFunc() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindAttribLocationFunc(
			this.Ref(),
		),
	)
}

// BindBuffer calls the method "WebGLRenderingContext.bindBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBindBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindBufferFunc returns the method "WebGLRenderingContext.bindBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BindBufferFunc() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindBufferFunc(
			this.Ref(),
		),
	)
}

// BindFramebuffer calls the method "WebGLRenderingContext.bindFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		framebuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindFramebufferFunc returns the method "WebGLRenderingContext.bindFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BindFramebufferFunc() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindFramebufferFunc(
			this.Ref(),
		),
	)
}

// BindRenderbuffer calls the method "WebGLRenderingContext.bindRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindRenderbufferFunc returns the method "WebGLRenderingContext.bindRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BindRenderbufferFunc() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindRenderbufferFunc(
			this.Ref(),
		),
	)
}

// BindTexture calls the method "WebGLRenderingContext.bindTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BindTexture(target GLenum, texture WebGLTexture) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBindTexture(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		texture.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindTextureFunc returns the method "WebGLRenderingContext.bindTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BindTextureFunc() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindTextureFunc(
			this.Ref(),
		),
	)
}

// BlendColor calls the method "WebGLRenderingContext.blendColor".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBlendColor(
		this.Ref(), js.Pointer(&_ok),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendColorFunc returns the method "WebGLRenderingContext.blendColor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BlendColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendColorFunc(
			this.Ref(),
		),
	)
}

// BlendEquation calls the method "WebGLRenderingContext.blendEquation".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BlendEquation(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBlendEquation(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationFunc returns the method "WebGLRenderingContext.blendEquation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BlendEquationFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendEquationFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparate calls the method "WebGLRenderingContext.blendEquationSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationSeparateFunc returns the method "WebGLRenderingContext.blendEquationSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BlendEquationSeparateFunc() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendEquationSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendFunc calls the method "WebGLRenderingContext.blendFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBlendFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(sfactor),
		uint32(dfactor),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFuncFunc returns the method "WebGLRenderingContext.blendFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BlendFuncFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendFuncFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparate calls the method "WebGLRenderingContext.blendFuncSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFuncSeparateFunc returns the method "WebGLRenderingContext.blendFuncSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BlendFuncSeparateFunc() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// CheckFramebufferStatus calls the method "WebGLRenderingContext.checkFramebufferStatus".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CheckFramebufferStatus(target GLenum) (GLenum, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
	)

	return GLenum(_ret), _ok
}

// CheckFramebufferStatusFunc returns the method "WebGLRenderingContext.checkFramebufferStatus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CheckFramebufferStatusFunc() (fn js.Func[func(target GLenum) GLenum]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCheckFramebufferStatusFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "WebGLRenderingContext.clear".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Clear(mask GLbitfield) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextClear(
		this.Ref(), js.Pointer(&_ok),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "WebGLRenderingContext.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ClearFunc() (fn js.Func[func(mask GLbitfield)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearFunc(
			this.Ref(),
		),
	)
}

// ClearColor calls the method "WebGLRenderingContext.clearColor".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextClearColor(
		this.Ref(), js.Pointer(&_ok),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearColorFunc returns the method "WebGLRenderingContext.clearColor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ClearColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearColorFunc(
			this.Ref(),
		),
	)
}

// ClearDepth calls the method "WebGLRenderingContext.clearDepth".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ClearDepth(depth GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextClearDepth(
		this.Ref(), js.Pointer(&_ok),
		float32(depth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearDepthFunc returns the method "WebGLRenderingContext.clearDepth".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ClearDepthFunc() (fn js.Func[func(depth GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearDepthFunc(
			this.Ref(),
		),
	)
}

// ClearStencil calls the method "WebGLRenderingContext.clearStencil".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ClearStencil(s GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextClearStencil(
		this.Ref(), js.Pointer(&_ok),
		int32(s),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearStencilFunc returns the method "WebGLRenderingContext.clearStencil".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ClearStencilFunc() (fn js.Func[func(s GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearStencilFunc(
			this.Ref(),
		),
	)
}

// ColorMask calls the method "WebGLRenderingContext.colorMask".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextColorMask(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ColorMaskFunc returns the method "WebGLRenderingContext.colorMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ColorMaskFunc() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextColorMaskFunc(
			this.Ref(),
		),
	)
}

// CompileShader calls the method "WebGLRenderingContext.compileShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CompileShader(shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCompileShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompileShaderFunc returns the method "WebGLRenderingContext.compileShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CompileShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompileShaderFunc(
			this.Ref(),
		),
	)
}

// CopyTexImage2D calls the method "WebGLRenderingContext.copyTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTexImage2DFunc returns the method "WebGLRenderingContext.copyTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CopyTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCopyTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexSubImage2D calls the method "WebGLRenderingContext.copyTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTexSubImage2DFunc returns the method "WebGLRenderingContext.copyTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CopyTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCopyTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "WebGLRenderingContext.createBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateBuffer() (WebGLBuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLBuffer{}.FromRef(_ret), _ok
}

// CreateBufferFunc returns the method "WebGLRenderingContext.createBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateBufferFunc() (fn js.Func[func() WebGLBuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateFramebuffer calls the method "WebGLRenderingContext.createFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateFramebuffer() (WebGLFramebuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLFramebuffer{}.FromRef(_ret), _ok
}

// CreateFramebufferFunc returns the method "WebGLRenderingContext.createFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateFramebufferFunc() (fn js.Func[func() WebGLFramebuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateFramebufferFunc(
			this.Ref(),
		),
	)
}

// CreateProgram calls the method "WebGLRenderingContext.createProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateProgram() (WebGLProgram, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateProgram(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLProgram{}.FromRef(_ret), _ok
}

// CreateProgramFunc returns the method "WebGLRenderingContext.createProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateProgramFunc() (fn js.Func[func() WebGLProgram]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateProgramFunc(
			this.Ref(),
		),
	)
}

// CreateRenderbuffer calls the method "WebGLRenderingContext.createRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateRenderbuffer() (WebGLRenderbuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLRenderbuffer{}.FromRef(_ret), _ok
}

// CreateRenderbufferFunc returns the method "WebGLRenderingContext.createRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateRenderbufferFunc() (fn js.Func[func() WebGLRenderbuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateRenderbufferFunc(
			this.Ref(),
		),
	)
}

// CreateShader calls the method "WebGLRenderingContext.createShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateShader(typ GLenum) (WebGLShader, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateShader(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return WebGLShader{}.FromRef(_ret), _ok
}

// CreateShaderFunc returns the method "WebGLRenderingContext.createShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateShaderFunc() (fn js.Func[func(typ GLenum) WebGLShader]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateShaderFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "WebGLRenderingContext.createTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CreateTexture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCreateTexture(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLTexture{}.FromRef(_ret), _ok
}

// CreateTextureFunc returns the method "WebGLRenderingContext.createTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CreateTextureFunc() (fn js.Func[func() WebGLTexture]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CullFace calls the method "WebGLRenderingContext.cullFace".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CullFace(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCullFace(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CullFaceFunc returns the method "WebGLRenderingContext.cullFace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CullFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCullFaceFunc(
			this.Ref(),
		),
	)
}

// DeleteBuffer calls the method "WebGLRenderingContext.deleteBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteBuffer(buffer WebGLBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteBufferFunc returns the method "WebGLRenderingContext.deleteBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteBufferFunc() (fn js.Func[func(buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteBufferFunc(
			this.Ref(),
		),
	)
}

// DeleteFramebuffer calls the method "WebGLRenderingContext.deleteFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		framebuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFramebufferFunc returns the method "WebGLRenderingContext.deleteFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteFramebufferFunc(
			this.Ref(),
		),
	)
}

// DeleteProgram calls the method "WebGLRenderingContext.deleteProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteProgramFunc returns the method "WebGLRenderingContext.deleteProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteProgramFunc(
			this.Ref(),
		),
	)
}

// DeleteRenderbuffer calls the method "WebGLRenderingContext.deleteRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRenderbufferFunc returns the method "WebGLRenderingContext.deleteRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteRenderbufferFunc(
			this.Ref(),
		),
	)
}

// DeleteShader calls the method "WebGLRenderingContext.deleteShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteShader(shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteShaderFunc returns the method "WebGLRenderingContext.deleteShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteShaderFunc(
			this.Ref(),
		),
	)
}

// DeleteTexture calls the method "WebGLRenderingContext.deleteTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DeleteTexture(texture WebGLTexture) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&_ok),
		texture.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteTextureFunc returns the method "WebGLRenderingContext.deleteTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DeleteTextureFunc() (fn js.Func[func(texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteTextureFunc(
			this.Ref(),
		),
	)
}

// DepthFunc calls the method "WebGLRenderingContext.depthFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DepthFunc(fn GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDepthFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(fn),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthFuncFunc returns the method "WebGLRenderingContext.depthFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DepthFuncFunc() (fn js.Func[func(fn GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthFuncFunc(
			this.Ref(),
		),
	)
}

// DepthMask calls the method "WebGLRenderingContext.depthMask".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DepthMask(flag GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDepthMask(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(flag)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthMaskFunc returns the method "WebGLRenderingContext.depthMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DepthMaskFunc() (fn js.Func[func(flag GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthMaskFunc(
			this.Ref(),
		),
	)
}

// DepthRange calls the method "WebGLRenderingContext.depthRange".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDepthRange(
		this.Ref(), js.Pointer(&_ok),
		float32(zNear),
		float32(zFar),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthRangeFunc returns the method "WebGLRenderingContext.depthRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DepthRangeFunc() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthRangeFunc(
			this.Ref(),
		),
	)
}

// DetachShader calls the method "WebGLRenderingContext.detachShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDetachShader(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DetachShaderFunc returns the method "WebGLRenderingContext.detachShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DetachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDetachShaderFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "WebGLRenderingContext.disable".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Disable(cap GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDisable(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisableFunc returns the method "WebGLRenderingContext.disable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DisableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDisableFunc(
			this.Ref(),
		),
	)
}

// DisableVertexAttribArray calls the method "WebGLRenderingContext.disableVertexAttribArray".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DisableVertexAttribArray(index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisableVertexAttribArrayFunc returns the method "WebGLRenderingContext.disableVertexAttribArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DisableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDisableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// DrawArrays calls the method "WebGLRenderingContext.drawArrays".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDrawArrays(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(first),
		int32(count),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawArraysFunc returns the method "WebGLRenderingContext.drawArrays".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DrawArraysFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDrawArraysFunc(
			this.Ref(),
		),
	)
}

// DrawElements calls the method "WebGLRenderingContext.drawElements".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextDrawElements(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawElementsFunc returns the method "WebGLRenderingContext.drawElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) DrawElementsFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDrawElementsFunc(
			this.Ref(),
		),
	)
}

// Enable calls the method "WebGLRenderingContext.enable".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Enable(cap GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextEnable(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnableFunc returns the method "WebGLRenderingContext.enable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) EnableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextEnableFunc(
			this.Ref(),
		),
	)
}

// EnableVertexAttribArray calls the method "WebGLRenderingContext.enableVertexAttribArray".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) EnableVertexAttribArray(index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnableVertexAttribArrayFunc returns the method "WebGLRenderingContext.enableVertexAttribArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) EnableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextEnableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "WebGLRenderingContext.finish".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Finish() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextFinish(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FinishFunc returns the method "WebGLRenderingContext.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFinishFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "WebGLRenderingContext.flush".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Flush() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextFlush(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FlushFunc returns the method "WebGLRenderingContext.flush".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFlushFunc(
			this.Ref(),
		),
	)
}

// FramebufferRenderbuffer calls the method "WebGLRenderingContext.framebufferRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferRenderbufferFunc returns the method "WebGLRenderingContext.framebufferRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) FramebufferRenderbufferFunc() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFramebufferRenderbufferFunc(
			this.Ref(),
		),
	)
}

// FramebufferTexture2D calls the method "WebGLRenderingContext.framebufferTexture2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferTexture2DFunc returns the method "WebGLRenderingContext.framebufferTexture2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) FramebufferTexture2DFunc() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFramebufferTexture2DFunc(
			this.Ref(),
		),
	)
}

// FrontFace calls the method "WebGLRenderingContext.frontFace".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) FrontFace(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextFrontFace(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FrontFaceFunc returns the method "WebGLRenderingContext.frontFace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) FrontFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFrontFaceFunc(
			this.Ref(),
		),
	)
}

// GenerateMipmap calls the method "WebGLRenderingContext.generateMipmap".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GenerateMipmap(target GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GenerateMipmapFunc returns the method "WebGLRenderingContext.generateMipmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GenerateMipmapFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGenerateMipmapFunc(
			this.Ref(),
		),
	)
}

// GetActiveAttrib calls the method "WebGLRenderingContext.getActiveAttrib".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (WebGLActiveInfo, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
	)

	return WebGLActiveInfo{}.FromRef(_ret), _ok
}

// GetActiveAttribFunc returns the method "WebGLRenderingContext.getActiveAttrib".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetActiveAttribFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetActiveAttribFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniform calls the method "WebGLRenderingContext.getActiveUniform".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (WebGLActiveInfo, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
	)

	return WebGLActiveInfo{}.FromRef(_ret), _ok
}

// GetActiveUniformFunc returns the method "WebGLRenderingContext.getActiveUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetActiveUniformFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetActiveUniformFunc(
			this.Ref(),
		),
	)
}

// GetAttachedShaders calls the method "WebGLRenderingContext.getAttachedShaders".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetAttachedShaders(program WebGLProgram) (js.Array[WebGLShader], bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return js.Array[WebGLShader]{}.FromRef(_ret), _ok
}

// GetAttachedShadersFunc returns the method "WebGLRenderingContext.getAttachedShaders".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetAttachedShadersFunc() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetAttachedShadersFunc(
			this.Ref(),
		),
	)
}

// GetAttribLocation calls the method "WebGLRenderingContext.getAttribLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (GLint, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		name.Ref(),
	)

	return GLint(_ret), _ok
}

// GetAttribLocationFunc returns the method "WebGLRenderingContext.getAttribLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetAttribLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetAttribLocationFunc(
			this.Ref(),
		),
	)
}

// GetBufferParameter calls the method "WebGLRenderingContext.getBufferParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetBufferParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetBufferParameterFunc returns the method "WebGLRenderingContext.getBufferParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetBufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetBufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetParameter calls the method "WebGLRenderingContext.getParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetParameter(pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetParameterFunc returns the method "WebGLRenderingContext.getParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetParameterFunc() (fn js.Func[func(pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetParameterFunc(
			this.Ref(),
		),
	)
}

// GetError calls the method "WebGLRenderingContext.getError".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetError() (GLenum, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetError(
		this.Ref(), js.Pointer(&_ok),
	)

	return GLenum(_ret), _ok
}

// GetErrorFunc returns the method "WebGLRenderingContext.getError".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetErrorFunc() (fn js.Func[func() GLenum]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetErrorFunc(
			this.Ref(),
		),
	)
}

// GetFramebufferAttachmentParameter calls the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetFramebufferAttachmentParameterFunc returns the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetFramebufferAttachmentParameterFunc() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetFramebufferAttachmentParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramParameter calls the method "WebGLRenderingContext.getProgramParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetProgramParameterFunc returns the method "WebGLRenderingContext.getProgramParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetProgramParameterFunc() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetProgramParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramInfoLog calls the method "WebGLRenderingContext.getProgramInfoLog".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetProgramInfoLog(program WebGLProgram) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetProgramInfoLogFunc returns the method "WebGLRenderingContext.getProgramInfoLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetProgramInfoLogFunc() (fn js.Func[func(program WebGLProgram) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetProgramInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetRenderbufferParameter calls the method "WebGLRenderingContext.getRenderbufferParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetRenderbufferParameterFunc returns the method "WebGLRenderingContext.getRenderbufferParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetRenderbufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetRenderbufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderParameter calls the method "WebGLRenderingContext.getShaderParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetShaderParameterFunc returns the method "WebGLRenderingContext.getShaderParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetShaderParameterFunc() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderPrecisionFormat calls the method "WebGLRenderingContext.getShaderPrecisionFormat".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (WebGLShaderPrecisionFormat, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&_ok),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return WebGLShaderPrecisionFormat{}.FromRef(_ret), _ok
}

// GetShaderPrecisionFormatFunc returns the method "WebGLRenderingContext.getShaderPrecisionFormat".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetShaderPrecisionFormatFunc() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderPrecisionFormatFunc(
			this.Ref(),
		),
	)
}

// GetShaderInfoLog calls the method "WebGLRenderingContext.getShaderInfoLog".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetShaderInfoLog(shader WebGLShader) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetShaderInfoLogFunc returns the method "WebGLRenderingContext.getShaderInfoLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetShaderInfoLogFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetShaderSource calls the method "WebGLRenderingContext.getShaderSource".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetShaderSource(shader WebGLShader) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetShaderSourceFunc returns the method "WebGLRenderingContext.getShaderSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetShaderSourceFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderSourceFunc(
			this.Ref(),
		),
	)
}

// GetTexParameter calls the method "WebGLRenderingContext.getTexParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetTexParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetTexParameterFunc returns the method "WebGLRenderingContext.getTexParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetTexParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetTexParameterFunc(
			this.Ref(),
		),
	)
}

// GetUniform calls the method "WebGLRenderingContext.getUniform".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetUniform(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		location.Ref(),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetUniformFunc returns the method "WebGLRenderingContext.getUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetUniformFunc() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetUniformFunc(
			this.Ref(),
		),
	)
}

// GetUniformLocation calls the method "WebGLRenderingContext.getUniformLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (WebGLUniformLocation, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		name.Ref(),
	)

	return WebGLUniformLocation{}.FromRef(_ret), _ok
}

// GetUniformLocationFunc returns the method "WebGLRenderingContext.getUniformLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetUniformLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetUniformLocationFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttrib calls the method "WebGLRenderingContext.getVertexAttrib".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetVertexAttribFunc returns the method "WebGLRenderingContext.getVertexAttrib".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetVertexAttribFunc() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetVertexAttribFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttribOffset calls the method "WebGLRenderingContext.getVertexAttribOffset".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (GLintptr, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(pname),
	)

	return GLintptr(_ret), _ok
}

// GetVertexAttribOffsetFunc returns the method "WebGLRenderingContext.getVertexAttribOffset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) GetVertexAttribOffsetFunc() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetVertexAttribOffsetFunc(
			this.Ref(),
		),
	)
}

// Hint calls the method "WebGLRenderingContext.hint".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Hint(target GLenum, mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextHint(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// HintFunc returns the method "WebGLRenderingContext.hint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) HintFunc() (fn js.Func[func(target GLenum, mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextHintFunc(
			this.Ref(),
		),
	)
}

// IsBuffer calls the method "WebGLRenderingContext.isBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsBuffer(buffer WebGLBuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsBufferFunc returns the method "WebGLRenderingContext.isBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsBufferFunc() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsBufferFunc(
			this.Ref(),
		),
	)
}

// IsEnabled calls the method "WebGLRenderingContext.isEnabled".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsEnabled(cap GLenum) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsEnabled(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	return _ret == js.True, _ok
}

// IsEnabledFunc returns the method "WebGLRenderingContext.isEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsEnabledFunc() (fn js.Func[func(cap GLenum) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsEnabledFunc(
			this.Ref(),
		),
	)
}

// IsFramebuffer calls the method "WebGLRenderingContext.isFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		framebuffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsFramebufferFunc returns the method "WebGLRenderingContext.isFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsFramebufferFunc(
			this.Ref(),
		),
	)
}

// IsProgram calls the method "WebGLRenderingContext.isProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsProgram(program WebGLProgram) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return _ret == js.True, _ok
}

// IsProgramFunc returns the method "WebGLRenderingContext.isProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsProgramFunc() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsProgramFunc(
			this.Ref(),
		),
	)
}

// IsRenderbuffer calls the method "WebGLRenderingContext.isRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		renderbuffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsRenderbufferFunc returns the method "WebGLRenderingContext.isRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsRenderbufferFunc(
			this.Ref(),
		),
	)
}

// IsShader calls the method "WebGLRenderingContext.isShader".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsShader(shader WebGLShader) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return _ret == js.True, _ok
}

// IsShaderFunc returns the method "WebGLRenderingContext.isShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsShaderFunc() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsShaderFunc(
			this.Ref(),
		),
	)
}

// IsTexture calls the method "WebGLRenderingContext.isTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) IsTexture(texture WebGLTexture) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextIsTexture(
		this.Ref(), js.Pointer(&_ok),
		texture.Ref(),
	)

	return _ret == js.True, _ok
}

// IsTextureFunc returns the method "WebGLRenderingContext.isTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) IsTextureFunc() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsTextureFunc(
			this.Ref(),
		),
	)
}

// LineWidth calls the method "WebGLRenderingContext.lineWidth".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) LineWidth(width GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextLineWidth(
		this.Ref(), js.Pointer(&_ok),
		float32(width),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineWidthFunc returns the method "WebGLRenderingContext.lineWidth".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) LineWidthFunc() (fn js.Func[func(width GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextLineWidthFunc(
			this.Ref(),
		),
	)
}

// LinkProgram calls the method "WebGLRenderingContext.linkProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) LinkProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextLinkProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LinkProgramFunc returns the method "WebGLRenderingContext.linkProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) LinkProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextLinkProgramFunc(
			this.Ref(),
		),
	)
}

// PixelStorei calls the method "WebGLRenderingContext.pixelStorei".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) PixelStorei(pname GLenum, param GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextPixelStorei(
		this.Ref(), js.Pointer(&_ok),
		uint32(pname),
		int32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PixelStoreiFunc returns the method "WebGLRenderingContext.pixelStorei".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) PixelStoreiFunc() (fn js.Func[func(pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextPixelStoreiFunc(
			this.Ref(),
		),
	)
}

// PolygonOffset calls the method "WebGLRenderingContext.polygonOffset".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&_ok),
		float32(factor),
		float32(units),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PolygonOffsetFunc returns the method "WebGLRenderingContext.polygonOffset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) PolygonOffsetFunc() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextPolygonOffsetFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorage calls the method "WebGLRenderingContext.renderbufferStorage".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RenderbufferStorageFunc returns the method "WebGLRenderingContext.renderbufferStorage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) RenderbufferStorageFunc() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextRenderbufferStorageFunc(
			this.Ref(),
		),
	)
}

// SampleCoverage calls the method "WebGLRenderingContext.sampleCoverage".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&_ok),
		float32(value),
		js.Bool(bool(invert)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SampleCoverageFunc returns the method "WebGLRenderingContext.sampleCoverage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) SampleCoverageFunc() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextSampleCoverageFunc(
			this.Ref(),
		),
	)
}

// Scissor calls the method "WebGLRenderingContext.scissor".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextScissor(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScissorFunc returns the method "WebGLRenderingContext.scissor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ScissorFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextScissorFunc(
			this.Ref(),
		),
	)
}

// ShaderSource calls the method "WebGLRenderingContext.shaderSource".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ShaderSource(shader WebGLShader, source js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextShaderSource(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShaderSourceFunc returns the method "WebGLRenderingContext.shaderSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ShaderSourceFunc() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextShaderSourceFunc(
			this.Ref(),
		),
	)
}

// StencilFunc calls the method "WebGLRenderingContext.stencilFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilFuncFunc returns the method "WebGLRenderingContext.stencilFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilFuncFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilFuncFunc(
			this.Ref(),
		),
	)
}

// StencilFuncSeparate calls the method "WebGLRenderingContext.stencilFuncSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilFuncSeparateFunc returns the method "WebGLRenderingContext.stencilFuncSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilFuncSeparateFunc() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilMask calls the method "WebGLRenderingContext.stencilMask".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilMask(mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilMask(
		this.Ref(), js.Pointer(&_ok),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilMaskFunc returns the method "WebGLRenderingContext.stencilMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilMaskFunc() (fn js.Func[func(mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilMaskFunc(
			this.Ref(),
		),
	)
}

// StencilMaskSeparate calls the method "WebGLRenderingContext.stencilMaskSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilMaskSeparateFunc returns the method "WebGLRenderingContext.stencilMaskSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilMaskSeparateFunc() (fn js.Func[func(face GLenum, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilMaskSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilOp calls the method "WebGLRenderingContext.stencilOp".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilOp(
		this.Ref(), js.Pointer(&_ok),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilOpFunc returns the method "WebGLRenderingContext.stencilOp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilOpFunc() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilOpFunc(
			this.Ref(),
		),
	)
}

// StencilOpSeparate calls the method "WebGLRenderingContext.stencilOpSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilOpSeparateFunc returns the method "WebGLRenderingContext.stencilOpSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) StencilOpSeparateFunc() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilOpSeparateFunc(
			this.Ref(),
		),
	)
}

// TexParameterf calls the method "WebGLRenderingContext.texParameterf".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexParameterf(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexParameterfFunc returns the method "WebGLRenderingContext.texParameterf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexParameterfFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexParameterfFunc(
			this.Ref(),
		),
	)
}

// TexParameteri calls the method "WebGLRenderingContext.texParameteri".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexParameteri(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexParameteriFunc returns the method "WebGLRenderingContext.texParameteri".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexParameteriFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexParameteriFunc(
			this.Ref(),
		),
	)
}

// Uniform1f calls the method "WebGLRenderingContext.uniform1f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform1f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fFunc returns the method "WebGLRenderingContext.uniform1f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform1fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1fFunc(
			this.Ref(),
		),
	)
}

// Uniform2f calls the method "WebGLRenderingContext.uniform2f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform2f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fFunc returns the method "WebGLRenderingContext.uniform2f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform2fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2fFunc(
			this.Ref(),
		),
	)
}

// Uniform3f calls the method "WebGLRenderingContext.uniform3f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform3f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fFunc returns the method "WebGLRenderingContext.uniform3f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform3fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3fFunc(
			this.Ref(),
		),
	)
}

// Uniform4f calls the method "WebGLRenderingContext.uniform4f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform4f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fFunc returns the method "WebGLRenderingContext.uniform4f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform4fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4fFunc(
			this.Ref(),
		),
	)
}

// Uniform1i calls the method "WebGLRenderingContext.uniform1i".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform1i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1iFunc returns the method "WebGLRenderingContext.uniform1i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform1iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1iFunc(
			this.Ref(),
		),
	)
}

// Uniform2i calls the method "WebGLRenderingContext.uniform2i".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform2i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2iFunc returns the method "WebGLRenderingContext.uniform2i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform2iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2iFunc(
			this.Ref(),
		),
	)
}

// Uniform3i calls the method "WebGLRenderingContext.uniform3i".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform3i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3iFunc returns the method "WebGLRenderingContext.uniform3i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform3iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3iFunc(
			this.Ref(),
		),
	)
}

// Uniform4i calls the method "WebGLRenderingContext.uniform4i".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform4i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4iFunc returns the method "WebGLRenderingContext.uniform4i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform4iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4iFunc(
			this.Ref(),
		),
	)
}

// UseProgram calls the method "WebGLRenderingContext.useProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) UseProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUseProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UseProgramFunc returns the method "WebGLRenderingContext.useProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) UseProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUseProgramFunc(
			this.Ref(),
		),
	)
}

// ValidateProgram calls the method "WebGLRenderingContext.validateProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ValidateProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextValidateProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ValidateProgramFunc returns the method "WebGLRenderingContext.validateProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ValidateProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextValidateProgramFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1f calls the method "WebGLRenderingContext.vertexAttrib1f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib1fFunc returns the method "WebGLRenderingContext.vertexAttrib1f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib1fFunc() (fn js.Func[func(index GLuint, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib1fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2f calls the method "WebGLRenderingContext.vertexAttrib2f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib2fFunc returns the method "WebGLRenderingContext.vertexAttrib2f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib2fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib2fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3f calls the method "WebGLRenderingContext.vertexAttrib3f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib3fFunc returns the method "WebGLRenderingContext.vertexAttrib3f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib3fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib3fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4f calls the method "WebGLRenderingContext.vertexAttrib4f".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib4fFunc returns the method "WebGLRenderingContext.vertexAttrib4f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib4fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib4fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1fv calls the method "WebGLRenderingContext.vertexAttrib1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib1fvFunc returns the method "WebGLRenderingContext.vertexAttrib1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib1fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib1fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2fv calls the method "WebGLRenderingContext.vertexAttrib2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib2fvFunc returns the method "WebGLRenderingContext.vertexAttrib2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib2fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib2fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3fv calls the method "WebGLRenderingContext.vertexAttrib3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib3fvFunc returns the method "WebGLRenderingContext.vertexAttrib3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib3fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib3fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4fv calls the method "WebGLRenderingContext.vertexAttrib4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib4fvFunc returns the method "WebGLRenderingContext.vertexAttrib4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttrib4fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib4fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttribPointer calls the method "WebGLRenderingContext.vertexAttribPointer".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribPointerFunc returns the method "WebGLRenderingContext.vertexAttribPointer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) VertexAttribPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttribPointerFunc(
			this.Ref(),
		),
	)
}

// Viewport calls the method "WebGLRenderingContext.viewport".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextViewport(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ViewportFunc returns the method "WebGLRenderingContext.viewport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ViewportFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextViewportFunc(
			this.Ref(),
		),
	)
}

// MakeXRCompatible calls the method "WebGLRenderingContext.makeXRCompatible".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) MakeXRCompatible() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MakeXRCompatibleFunc returns the method "WebGLRenderingContext.makeXRCompatible".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) MakeXRCompatibleFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextMakeXRCompatibleFunc(
			this.Ref(),
		),
	)
}

// BufferData calls the method "WebGLRenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBufferData(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferDataFunc returns the method "WebGLRenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BufferDataFunc() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferDataFunc(
			this.Ref(),
		),
	)
}

// BufferData1 calls the method "WebGLRenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BufferData1(target GLenum, data AllowSharedBufferSource, usage GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBufferData1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		data.Ref(),
		uint32(usage),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferData1Func returns the method "WebGLRenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BufferData1Func() (fn js.Func[func(target GLenum, data AllowSharedBufferSource, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferData1Func(
			this.Ref(),
		),
	)
}

// BufferSubData calls the method "WebGLRenderingContext.bufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) BufferSubData(target GLenum, offset GLintptr, data AllowSharedBufferSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextBufferSubData(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(offset),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferSubDataFunc returns the method "WebGLRenderingContext.bufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) BufferSubDataFunc() (fn js.Func[func(target GLenum, offset GLintptr, data AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D calls the method "WebGLRenderingContext.compressedTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage2DFunc returns the method "WebGLRenderingContext.compressedTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CompressedTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompressedTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D calls the method "WebGLRenderingContext.compressedTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage2DFunc returns the method "WebGLRenderingContext.compressedTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) CompressedTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompressedTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// ReadPixels calls the method "WebGLRenderingContext.readPixels".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextReadPixels(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadPixelsFunc returns the method "WebGLRenderingContext.readPixels".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) ReadPixelsFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextReadPixelsFunc(
			this.Ref(),
		),
	)
}

// TexImage2D calls the method "WebGLRenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2DFunc returns the method "WebGLRenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexImage2DFunc(
			this.Ref(),
		),
	)
}

// TexImage2D1 calls the method "WebGLRenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2D1Func returns the method "WebGLRenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexImage2D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D calls the method "WebGLRenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2DFunc returns the method "WebGLRenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage2D1 calls the method "WebGLRenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2D1Func returns the method "WebGLRenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) TexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// Uniform1fv calls the method "WebGLRenderingContext.uniform1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform1fv(location WebGLUniformLocation, v Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform1fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fvFunc returns the method "WebGLRenderingContext.uniform1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform1fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1fvFunc(
			this.Ref(),
		),
	)
}

// Uniform2fv calls the method "WebGLRenderingContext.uniform2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform2fv(location WebGLUniformLocation, v Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fvFunc returns the method "WebGLRenderingContext.uniform2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform2fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2fvFunc(
			this.Ref(),
		),
	)
}

// Uniform3fv calls the method "WebGLRenderingContext.uniform3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform3fv(location WebGLUniformLocation, v Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fvFunc returns the method "WebGLRenderingContext.uniform3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform3fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3fvFunc(
			this.Ref(),
		),
	)
}

// Uniform4fv calls the method "WebGLRenderingContext.uniform4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform4fv(location WebGLUniformLocation, v Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fvFunc returns the method "WebGLRenderingContext.uniform4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform4fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4fvFunc(
			this.Ref(),
		),
	)
}

// Uniform1iv calls the method "WebGLRenderingContext.uniform1iv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform1iv(location WebGLUniformLocation, v Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform1iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1ivFunc returns the method "WebGLRenderingContext.uniform1iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform1ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1ivFunc(
			this.Ref(),
		),
	)
}

// Uniform2iv calls the method "WebGLRenderingContext.uniform2iv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform2iv(location WebGLUniformLocation, v Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform2iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2ivFunc returns the method "WebGLRenderingContext.uniform2iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform2ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2ivFunc(
			this.Ref(),
		),
	)
}

// Uniform3iv calls the method "WebGLRenderingContext.uniform3iv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform3iv(location WebGLUniformLocation, v Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform3iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3ivFunc returns the method "WebGLRenderingContext.uniform3iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform3ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3ivFunc(
			this.Ref(),
		),
	)
}

// Uniform4iv calls the method "WebGLRenderingContext.uniform4iv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) Uniform4iv(location WebGLUniformLocation, v Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniform4iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		v.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4ivFunc returns the method "WebGLRenderingContext.uniform4iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) Uniform4ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4ivFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv calls the method "WebGLRenderingContext.uniformMatrix2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2fvFunc returns the method "WebGLRenderingContext.uniformMatrix2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) UniformMatrix2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv calls the method "WebGLRenderingContext.uniformMatrix3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3fvFunc returns the method "WebGLRenderingContext.uniformMatrix3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) UniformMatrix3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv calls the method "WebGLRenderingContext.uniformMatrix4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGLRenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGLRenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4fvFunc returns the method "WebGLRenderingContext.uniformMatrix4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGLRenderingContext) UniformMatrix4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix4fvFunc(
			this.Ref(),
		),
	)
}

type GLint64 int64

const (
	WebGL2RenderingContext_READ_BUFFER                                   GLenum = 0x0C02
	WebGL2RenderingContext_UNPACK_ROW_LENGTH                             GLenum = 0x0CF2
	WebGL2RenderingContext_UNPACK_SKIP_ROWS                              GLenum = 0x0CF3
	WebGL2RenderingContext_UNPACK_SKIP_PIXELS                            GLenum = 0x0CF4
	WebGL2RenderingContext_PACK_ROW_LENGTH                               GLenum = 0x0D02
	WebGL2RenderingContext_PACK_SKIP_ROWS                                GLenum = 0x0D03
	WebGL2RenderingContext_PACK_SKIP_PIXELS                              GLenum = 0x0D04
	WebGL2RenderingContext_COLOR                                         GLenum = 0x1800
	WebGL2RenderingContext_DEPTH                                         GLenum = 0x1801
	WebGL2RenderingContext_STENCIL                                       GLenum = 0x1802
	WebGL2RenderingContext_RED                                           GLenum = 0x1903
	WebGL2RenderingContext_RGB8                                          GLenum = 0x8051
	WebGL2RenderingContext_RGBA8                                         GLenum = 0x8058
	WebGL2RenderingContext_RGB10_A2                                      GLenum = 0x8059
	WebGL2RenderingContext_TEXTURE_BINDING_3D                            GLenum = 0x806A
	WebGL2RenderingContext_UNPACK_SKIP_IMAGES                            GLenum = 0x806D
	WebGL2RenderingContext_UNPACK_IMAGE_HEIGHT                           GLenum = 0x806E
	WebGL2RenderingContext_TEXTURE_3D                                    GLenum = 0x806F
	WebGL2RenderingContext_TEXTURE_WRAP_R                                GLenum = 0x8072
	WebGL2RenderingContext_MAX_3D_TEXTURE_SIZE                           GLenum = 0x8073
	WebGL2RenderingContext_UNSIGNED_INT_2_10_10_10_REV                   GLenum = 0x8368
	WebGL2RenderingContext_MAX_ELEMENTS_VERTICES                         GLenum = 0x80E8
	WebGL2RenderingContext_MAX_ELEMENTS_INDICES                          GLenum = 0x80E9
	WebGL2RenderingContext_TEXTURE_MIN_LOD                               GLenum = 0x813A
	WebGL2RenderingContext_TEXTURE_MAX_LOD                               GLenum = 0x813B
	WebGL2RenderingContext_TEXTURE_BASE_LEVEL                            GLenum = 0x813C
	WebGL2RenderingContext_TEXTURE_MAX_LEVEL                             GLenum = 0x813D
	WebGL2RenderingContext_MIN                                           GLenum = 0x8007
	WebGL2RenderingContext_MAX                                           GLenum = 0x8008
	WebGL2RenderingContext_DEPTH_COMPONENT24                             GLenum = 0x81A6
	WebGL2RenderingContext_MAX_TEXTURE_LOD_BIAS                          GLenum = 0x84FD
	WebGL2RenderingContext_TEXTURE_COMPARE_MODE                          GLenum = 0x884C
	WebGL2RenderingContext_TEXTURE_COMPARE_FUNC                          GLenum = 0x884D
	WebGL2RenderingContext_CURRENT_QUERY                                 GLenum = 0x8865
	WebGL2RenderingContext_QUERY_RESULT                                  GLenum = 0x8866
	WebGL2RenderingContext_QUERY_RESULT_AVAILABLE                        GLenum = 0x8867
	WebGL2RenderingContext_STREAM_READ                                   GLenum = 0x88E1
	WebGL2RenderingContext_STREAM_COPY                                   GLenum = 0x88E2
	WebGL2RenderingContext_STATIC_READ                                   GLenum = 0x88E5
	WebGL2RenderingContext_STATIC_COPY                                   GLenum = 0x88E6
	WebGL2RenderingContext_DYNAMIC_READ                                  GLenum = 0x88E9
	WebGL2RenderingContext_DYNAMIC_COPY                                  GLenum = 0x88EA
	WebGL2RenderingContext_MAX_DRAW_BUFFERS                              GLenum = 0x8824
	WebGL2RenderingContext_DRAW_BUFFER0                                  GLenum = 0x8825
	WebGL2RenderingContext_DRAW_BUFFER1                                  GLenum = 0x8826
	WebGL2RenderingContext_DRAW_BUFFER2                                  GLenum = 0x8827
	WebGL2RenderingContext_DRAW_BUFFER3                                  GLenum = 0x8828
	WebGL2RenderingContext_DRAW_BUFFER4                                  GLenum = 0x8829
	WebGL2RenderingContext_DRAW_BUFFER5                                  GLenum = 0x882A
	WebGL2RenderingContext_DRAW_BUFFER6                                  GLenum = 0x882B
	WebGL2RenderingContext_DRAW_BUFFER7                                  GLenum = 0x882C
	WebGL2RenderingContext_DRAW_BUFFER8                                  GLenum = 0x882D
	WebGL2RenderingContext_DRAW_BUFFER9                                  GLenum = 0x882E
	WebGL2RenderingContext_DRAW_BUFFER10                                 GLenum = 0x882F
	WebGL2RenderingContext_DRAW_BUFFER11                                 GLenum = 0x8830
	WebGL2RenderingContext_DRAW_BUFFER12                                 GLenum = 0x8831
	WebGL2RenderingContext_DRAW_BUFFER13                                 GLenum = 0x8832
	WebGL2RenderingContext_DRAW_BUFFER14                                 GLenum = 0x8833
	WebGL2RenderingContext_DRAW_BUFFER15                                 GLenum = 0x8834
	WebGL2RenderingContext_MAX_FRAGMENT_UNIFORM_COMPONENTS               GLenum = 0x8B49
	WebGL2RenderingContext_MAX_VERTEX_UNIFORM_COMPONENTS                 GLenum = 0x8B4A
	WebGL2RenderingContext_SAMPLER_3D                                    GLenum = 0x8B5F
	WebGL2RenderingContext_SAMPLER_2D_SHADOW                             GLenum = 0x8B62
	WebGL2RenderingContext_FRAGMENT_SHADER_DERIVATIVE_HINT               GLenum = 0x8B8B
	WebGL2RenderingContext_PIXEL_PACK_BUFFER                             GLenum = 0x88EB
	WebGL2RenderingContext_PIXEL_UNPACK_BUFFER                           GLenum = 0x88EC
	WebGL2RenderingContext_PIXEL_PACK_BUFFER_BINDING                     GLenum = 0x88ED
	WebGL2RenderingContext_PIXEL_UNPACK_BUFFER_BINDING                   GLenum = 0x88EF
	WebGL2RenderingContext_FLOAT_MAT2x3                                  GLenum = 0x8B65
	WebGL2RenderingContext_FLOAT_MAT2x4                                  GLenum = 0x8B66
	WebGL2RenderingContext_FLOAT_MAT3x2                                  GLenum = 0x8B67
	WebGL2RenderingContext_FLOAT_MAT3x4                                  GLenum = 0x8B68
	WebGL2RenderingContext_FLOAT_MAT4x2                                  GLenum = 0x8B69
	WebGL2RenderingContext_FLOAT_MAT4x3                                  GLenum = 0x8B6A
	WebGL2RenderingContext_SRGB                                          GLenum = 0x8C40
	WebGL2RenderingContext_SRGB8                                         GLenum = 0x8C41
	WebGL2RenderingContext_SRGB8_ALPHA8                                  GLenum = 0x8C43
	WebGL2RenderingContext_COMPARE_REF_TO_TEXTURE                        GLenum = 0x884E
	WebGL2RenderingContext_RGBA32F                                       GLenum = 0x8814
	WebGL2RenderingContext_RGB32F                                        GLenum = 0x8815
	WebGL2RenderingContext_RGBA16F                                       GLenum = 0x881A
	WebGL2RenderingContext_RGB16F                                        GLenum = 0x881B
	WebGL2RenderingContext_VERTEX_ATTRIB_ARRAY_INTEGER                   GLenum = 0x88FD
	WebGL2RenderingContext_MAX_ARRAY_TEXTURE_LAYERS                      GLenum = 0x88FF
	WebGL2RenderingContext_MIN_PROGRAM_TEXEL_OFFSET                      GLenum = 0x8904
	WebGL2RenderingContext_MAX_PROGRAM_TEXEL_OFFSET                      GLenum = 0x8905
	WebGL2RenderingContext_MAX_VARYING_COMPONENTS                        GLenum = 0x8B4B
	WebGL2RenderingContext_TEXTURE_2D_ARRAY                              GLenum = 0x8C1A
	WebGL2RenderingContext_TEXTURE_BINDING_2D_ARRAY                      GLenum = 0x8C1D
	WebGL2RenderingContext_R11F_G11F_B10F                                GLenum = 0x8C3A
	WebGL2RenderingContext_UNSIGNED_INT_10F_11F_11F_REV                  GLenum = 0x8C3B
	WebGL2RenderingContext_RGB9_E5                                       GLenum = 0x8C3D
	WebGL2RenderingContext_UNSIGNED_INT_5_9_9_9_REV                      GLenum = 0x8C3E
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BUFFER_MODE                GLenum = 0x8C7F
	WebGL2RenderingContext_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS    GLenum = 0x8C80
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_VARYINGS                   GLenum = 0x8C83
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BUFFER_START               GLenum = 0x8C84
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BUFFER_SIZE                GLenum = 0x8C85
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN         GLenum = 0x8C88
	WebGL2RenderingContext_RASTERIZER_DISCARD                            GLenum = 0x8C89
	WebGL2RenderingContext_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS GLenum = 0x8C8A
	WebGL2RenderingContext_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS       GLenum = 0x8C8B
	WebGL2RenderingContext_INTERLEAVED_ATTRIBS                           GLenum = 0x8C8C
	WebGL2RenderingContext_SEPARATE_ATTRIBS                              GLenum = 0x8C8D
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BUFFER                     GLenum = 0x8C8E
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BUFFER_BINDING             GLenum = 0x8C8F
	WebGL2RenderingContext_RGBA32UI                                      GLenum = 0x8D70
	WebGL2RenderingContext_RGB32UI                                       GLenum = 0x8D71
	WebGL2RenderingContext_RGBA16UI                                      GLenum = 0x8D76
	WebGL2RenderingContext_RGB16UI                                       GLenum = 0x8D77
	WebGL2RenderingContext_RGBA8UI                                       GLenum = 0x8D7C
	WebGL2RenderingContext_RGB8UI                                        GLenum = 0x8D7D
	WebGL2RenderingContext_RGBA32I                                       GLenum = 0x8D82
	WebGL2RenderingContext_RGB32I                                        GLenum = 0x8D83
	WebGL2RenderingContext_RGBA16I                                       GLenum = 0x8D88
	WebGL2RenderingContext_RGB16I                                        GLenum = 0x8D89
	WebGL2RenderingContext_RGBA8I                                        GLenum = 0x8D8E
	WebGL2RenderingContext_RGB8I                                         GLenum = 0x8D8F
	WebGL2RenderingContext_RED_INTEGER                                   GLenum = 0x8D94
	WebGL2RenderingContext_RGB_INTEGER                                   GLenum = 0x8D98
	WebGL2RenderingContext_RGBA_INTEGER                                  GLenum = 0x8D99
	WebGL2RenderingContext_SAMPLER_2D_ARRAY                              GLenum = 0x8DC1
	WebGL2RenderingContext_SAMPLER_2D_ARRAY_SHADOW                       GLenum = 0x8DC4
	WebGL2RenderingContext_SAMPLER_CUBE_SHADOW                           GLenum = 0x8DC5
	WebGL2RenderingContext_UNSIGNED_INT_VEC2                             GLenum = 0x8DC6
	WebGL2RenderingContext_UNSIGNED_INT_VEC3                             GLenum = 0x8DC7
	WebGL2RenderingContext_UNSIGNED_INT_VEC4                             GLenum = 0x8DC8
	WebGL2RenderingContext_INT_SAMPLER_2D                                GLenum = 0x8DCA
	WebGL2RenderingContext_INT_SAMPLER_3D                                GLenum = 0x8DCB
	WebGL2RenderingContext_INT_SAMPLER_CUBE                              GLenum = 0x8DCC
	WebGL2RenderingContext_INT_SAMPLER_2D_ARRAY                          GLenum = 0x8DCF
	WebGL2RenderingContext_UNSIGNED_INT_SAMPLER_2D                       GLenum = 0x8DD2
	WebGL2RenderingContext_UNSIGNED_INT_SAMPLER_3D                       GLenum = 0x8DD3
	WebGL2RenderingContext_UNSIGNED_INT_SAMPLER_CUBE                     GLenum = 0x8DD4
	WebGL2RenderingContext_UNSIGNED_INT_SAMPLER_2D_ARRAY                 GLenum = 0x8DD7
	WebGL2RenderingContext_DEPTH_COMPONENT32F                            GLenum = 0x8CAC
	WebGL2RenderingContext_DEPTH32F_STENCIL8                             GLenum = 0x8CAD
	WebGL2RenderingContext_FLOAT_32_UNSIGNED_INT_24_8_REV                GLenum = 0x8DAD
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING         GLenum = 0x8210
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE         GLenum = 0x8211
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_RED_SIZE               GLenum = 0x8212
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE             GLenum = 0x8213
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE              GLenum = 0x8214
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE             GLenum = 0x8215
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE             GLenum = 0x8216
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE           GLenum = 0x8217
	WebGL2RenderingContext_FRAMEBUFFER_DEFAULT                           GLenum = 0x8218
	WebGL2RenderingContext_UNSIGNED_INT_24_8                             GLenum = 0x84FA
	WebGL2RenderingContext_DEPTH24_STENCIL8                              GLenum = 0x88F0
	WebGL2RenderingContext_UNSIGNED_NORMALIZED                           GLenum = 0x8C17
	WebGL2RenderingContext_DRAW_FRAMEBUFFER_BINDING                      GLenum = 0x8CA6
	WebGL2RenderingContext_READ_FRAMEBUFFER                              GLenum = 0x8CA8
	WebGL2RenderingContext_DRAW_FRAMEBUFFER                              GLenum = 0x8CA9
	WebGL2RenderingContext_READ_FRAMEBUFFER_BINDING                      GLenum = 0x8CAA
	WebGL2RenderingContext_RENDERBUFFER_SAMPLES                          GLenum = 0x8CAB
	WebGL2RenderingContext_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER          GLenum = 0x8CD4
	WebGL2RenderingContext_MAX_COLOR_ATTACHMENTS                         GLenum = 0x8CDF
	WebGL2RenderingContext_COLOR_ATTACHMENT1                             GLenum = 0x8CE1
	WebGL2RenderingContext_COLOR_ATTACHMENT2                             GLenum = 0x8CE2
	WebGL2RenderingContext_COLOR_ATTACHMENT3                             GLenum = 0x8CE3
	WebGL2RenderingContext_COLOR_ATTACHMENT4                             GLenum = 0x8CE4
	WebGL2RenderingContext_COLOR_ATTACHMENT5                             GLenum = 0x8CE5
	WebGL2RenderingContext_COLOR_ATTACHMENT6                             GLenum = 0x8CE6
	WebGL2RenderingContext_COLOR_ATTACHMENT7                             GLenum = 0x8CE7
	WebGL2RenderingContext_COLOR_ATTACHMENT8                             GLenum = 0x8CE8
	WebGL2RenderingContext_COLOR_ATTACHMENT9                             GLenum = 0x8CE9
	WebGL2RenderingContext_COLOR_ATTACHMENT10                            GLenum = 0x8CEA
	WebGL2RenderingContext_COLOR_ATTACHMENT11                            GLenum = 0x8CEB
	WebGL2RenderingContext_COLOR_ATTACHMENT12                            GLenum = 0x8CEC
	WebGL2RenderingContext_COLOR_ATTACHMENT13                            GLenum = 0x8CED
	WebGL2RenderingContext_COLOR_ATTACHMENT14                            GLenum = 0x8CEE
	WebGL2RenderingContext_COLOR_ATTACHMENT15                            GLenum = 0x8CEF
	WebGL2RenderingContext_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE            GLenum = 0x8D56
	WebGL2RenderingContext_MAX_SAMPLES                                   GLenum = 0x8D57
	WebGL2RenderingContext_HALF_FLOAT                                    GLenum = 0x140B
	WebGL2RenderingContext_RG                                            GLenum = 0x8227
	WebGL2RenderingContext_RG_INTEGER                                    GLenum = 0x8228
	WebGL2RenderingContext_R8                                            GLenum = 0x8229
	WebGL2RenderingContext_RG8                                           GLenum = 0x822B
	WebGL2RenderingContext_R16F                                          GLenum = 0x822D
	WebGL2RenderingContext_R32F                                          GLenum = 0x822E
	WebGL2RenderingContext_RG16F                                         GLenum = 0x822F
	WebGL2RenderingContext_RG32F                                         GLenum = 0x8230
	WebGL2RenderingContext_R8I                                           GLenum = 0x8231
	WebGL2RenderingContext_R8UI                                          GLenum = 0x8232
	WebGL2RenderingContext_R16I                                          GLenum = 0x8233
	WebGL2RenderingContext_R16UI                                         GLenum = 0x8234
	WebGL2RenderingContext_R32I                                          GLenum = 0x8235
	WebGL2RenderingContext_R32UI                                         GLenum = 0x8236
	WebGL2RenderingContext_RG8I                                          GLenum = 0x8237
	WebGL2RenderingContext_RG8UI                                         GLenum = 0x8238
	WebGL2RenderingContext_RG16I                                         GLenum = 0x8239
	WebGL2RenderingContext_RG16UI                                        GLenum = 0x823A
	WebGL2RenderingContext_RG32I                                         GLenum = 0x823B
	WebGL2RenderingContext_RG32UI                                        GLenum = 0x823C
	WebGL2RenderingContext_VERTEX_ARRAY_BINDING                          GLenum = 0x85B5
	WebGL2RenderingContext_R8_SNORM                                      GLenum = 0x8F94
	WebGL2RenderingContext_RG8_SNORM                                     GLenum = 0x8F95
	WebGL2RenderingContext_RGB8_SNORM                                    GLenum = 0x8F96
	WebGL2RenderingContext_RGBA8_SNORM                                   GLenum = 0x8F97
	WebGL2RenderingContext_SIGNED_NORMALIZED                             GLenum = 0x8F9C
	WebGL2RenderingContext_COPY_READ_BUFFER                              GLenum = 0x8F36
	WebGL2RenderingContext_COPY_WRITE_BUFFER                             GLenum = 0x8F37
	WebGL2RenderingContext_COPY_READ_BUFFER_BINDING                      GLenum = 0x8F36
	WebGL2RenderingContext_COPY_WRITE_BUFFER_BINDING                     GLenum = 0x8F37
	WebGL2RenderingContext_UNIFORM_BUFFER                                GLenum = 0x8A11
	WebGL2RenderingContext_UNIFORM_BUFFER_BINDING                        GLenum = 0x8A28
	WebGL2RenderingContext_UNIFORM_BUFFER_START                          GLenum = 0x8A29
	WebGL2RenderingContext_UNIFORM_BUFFER_SIZE                           GLenum = 0x8A2A
	WebGL2RenderingContext_MAX_VERTEX_UNIFORM_BLOCKS                     GLenum = 0x8A2B
	WebGL2RenderingContext_MAX_FRAGMENT_UNIFORM_BLOCKS                   GLenum = 0x8A2D
	WebGL2RenderingContext_MAX_COMBINED_UNIFORM_BLOCKS                   GLenum = 0x8A2E
	WebGL2RenderingContext_MAX_UNIFORM_BUFFER_BINDINGS                   GLenum = 0x8A2F
	WebGL2RenderingContext_MAX_UNIFORM_BLOCK_SIZE                        GLenum = 0x8A30
	WebGL2RenderingContext_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS        GLenum = 0x8A31
	WebGL2RenderingContext_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS      GLenum = 0x8A33
	WebGL2RenderingContext_UNIFORM_BUFFER_OFFSET_ALIGNMENT               GLenum = 0x8A34
	WebGL2RenderingContext_ACTIVE_UNIFORM_BLOCKS                         GLenum = 0x8A36
	WebGL2RenderingContext_UNIFORM_TYPE                                  GLenum = 0x8A37
	WebGL2RenderingContext_UNIFORM_SIZE                                  GLenum = 0x8A38
	WebGL2RenderingContext_UNIFORM_BLOCK_INDEX                           GLenum = 0x8A3A
	WebGL2RenderingContext_UNIFORM_OFFSET                                GLenum = 0x8A3B
	WebGL2RenderingContext_UNIFORM_ARRAY_STRIDE                          GLenum = 0x8A3C
	WebGL2RenderingContext_UNIFORM_MATRIX_STRIDE                         GLenum = 0x8A3D
	WebGL2RenderingContext_UNIFORM_IS_ROW_MAJOR                          GLenum = 0x8A3E
	WebGL2RenderingContext_UNIFORM_BLOCK_BINDING                         GLenum = 0x8A3F
	WebGL2RenderingContext_UNIFORM_BLOCK_DATA_SIZE                       GLenum = 0x8A40
	WebGL2RenderingContext_UNIFORM_BLOCK_ACTIVE_UNIFORMS                 GLenum = 0x8A42
	WebGL2RenderingContext_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES          GLenum = 0x8A43
	WebGL2RenderingContext_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER     GLenum = 0x8A44
	WebGL2RenderingContext_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER   GLenum = 0x8A46
	WebGL2RenderingContext_INVALID_INDEX                                 GLenum = 0xFFFFFFFF
	WebGL2RenderingContext_MAX_VERTEX_OUTPUT_COMPONENTS                  GLenum = 0x9122
	WebGL2RenderingContext_MAX_FRAGMENT_INPUT_COMPONENTS                 GLenum = 0x9125
	WebGL2RenderingContext_MAX_SERVER_WAIT_TIMEOUT                       GLenum = 0x9111
	WebGL2RenderingContext_OBJECT_TYPE                                   GLenum = 0x9112
	WebGL2RenderingContext_SYNC_CONDITION                                GLenum = 0x9113
	WebGL2RenderingContext_SYNC_STATUS                                   GLenum = 0x9114
	WebGL2RenderingContext_SYNC_FLAGS                                    GLenum = 0x9115
	WebGL2RenderingContext_SYNC_FENCE                                    GLenum = 0x9116
	WebGL2RenderingContext_SYNC_GPU_COMMANDS_COMPLETE                    GLenum = 0x9117
	WebGL2RenderingContext_UNSIGNALED                                    GLenum = 0x9118
	WebGL2RenderingContext_SIGNALED                                      GLenum = 0x9119
	WebGL2RenderingContext_ALREADY_SIGNALED                              GLenum = 0x911A
	WebGL2RenderingContext_TIMEOUT_EXPIRED                               GLenum = 0x911B
	WebGL2RenderingContext_CONDITION_SATISFIED                           GLenum = 0x911C
	WebGL2RenderingContext_WAIT_FAILED                                   GLenum = 0x911D
	WebGL2RenderingContext_SYNC_FLUSH_COMMANDS_BIT                       GLenum = 0x00000001
	WebGL2RenderingContext_VERTEX_ATTRIB_ARRAY_DIVISOR                   GLenum = 0x88FE
	WebGL2RenderingContext_ANY_SAMPLES_PASSED                            GLenum = 0x8C2F
	WebGL2RenderingContext_ANY_SAMPLES_PASSED_CONSERVATIVE               GLenum = 0x8D6A
	WebGL2RenderingContext_SAMPLER_BINDING                               GLenum = 0x8919
	WebGL2RenderingContext_RGB10_A2UI                                    GLenum = 0x906F
	WebGL2RenderingContext_INT_2_10_10_10_REV                            GLenum = 0x8D9F
	WebGL2RenderingContext_TRANSFORM_FEEDBACK                            GLenum = 0x8E22
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_PAUSED                     GLenum = 0x8E23
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_ACTIVE                     GLenum = 0x8E24
	WebGL2RenderingContext_TRANSFORM_FEEDBACK_BINDING                    GLenum = 0x8E25
	WebGL2RenderingContext_TEXTURE_IMMUTABLE_FORMAT                      GLenum = 0x912F
	WebGL2RenderingContext_MAX_ELEMENT_INDEX                             GLenum = 0x8D6B
	WebGL2RenderingContext_TEXTURE_IMMUTABLE_LEVELS                      GLenum = 0x82DF
)

const (
	WebGL2RenderingContext_TIMEOUT_IGNORED GLint64 = -1
)

const (
	WebGL2RenderingContext_MAX_CLIENT_WAIT_TIMEOUT_WEBGL GLenum = 0x9247
)

type OneOf_TypedArrayUint32_ArrayGLuint struct {
	ref js.Ref
}

func (x OneOf_TypedArrayUint32_ArrayGLuint) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayUint32_ArrayGLuint) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayUint32_ArrayGLuint) FromRef(ref js.Ref) OneOf_TypedArrayUint32_ArrayGLuint {
	return OneOf_TypedArrayUint32_ArrayGLuint{
		ref: ref,
	}
}

func (x OneOf_TypedArrayUint32_ArrayGLuint) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayUint32_ArrayGLuint) ArrayGLuint() js.Array[GLuint] {
	return js.Array[GLuint]{}.FromRef(x.ref)
}

type Uint32List = OneOf_TypedArrayUint32_ArrayGLuint

type WebGLQuery struct {
	WebGLObject
}

func (this WebGLQuery) Once() WebGLQuery {
	this.Ref().Once()
	return this
}

func (this WebGLQuery) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLQuery) FromRef(ref js.Ref) WebGLQuery {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLQuery) Free() {
	this.Ref().Free()
}

type WebGLSampler struct {
	WebGLObject
}

func (this WebGLSampler) Once() WebGLSampler {
	this.Ref().Once()
	return this
}

func (this WebGLSampler) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLSampler) FromRef(ref js.Ref) WebGLSampler {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLSampler) Free() {
	this.Ref().Free()
}

type WebGLSync struct {
	WebGLObject
}

func (this WebGLSync) Once() WebGLSync {
	this.Ref().Once()
	return this
}

func (this WebGLSync) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLSync) FromRef(ref js.Ref) WebGLSync {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLSync) Free() {
	this.Ref().Free()
}

type GLuint64 uint64

type WebGLTransformFeedback struct {
	WebGLObject
}

func (this WebGLTransformFeedback) Once() WebGLTransformFeedback {
	this.Ref().Once()
	return this
}

func (this WebGLTransformFeedback) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLTransformFeedback) FromRef(ref js.Ref) WebGLTransformFeedback {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLTransformFeedback) Free() {
	this.Ref().Free()
}

type WebGLVertexArrayObject struct {
	WebGLObject
}

func (this WebGLVertexArrayObject) Once() WebGLVertexArrayObject {
	this.Ref().Once()
	return this
}

func (this WebGLVertexArrayObject) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLVertexArrayObject) FromRef(ref js.Ref) WebGLVertexArrayObject {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLVertexArrayObject) Free() {
	this.Ref().Free()
}

type WebGL2RenderingContext struct {
	ref js.Ref
}

func (this WebGL2RenderingContext) Once() WebGL2RenderingContext {
	this.Ref().Once()
	return this
}

func (this WebGL2RenderingContext) Ref() js.Ref {
	return this.ref
}

func (this WebGL2RenderingContext) FromRef(ref js.Ref) WebGL2RenderingContext {
	this.ref = ref
	return this
}

func (this WebGL2RenderingContext) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "WebGL2RenderingContext.canvas".
//
// The returned bool will be false if there is no such property.
func (this WebGL2RenderingContext) Canvas() (OneOf_HTMLCanvasElement_OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.GetWebGL2RenderingContextCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_HTMLCanvasElement_OffscreenCanvas{}.FromRef(_ret), _ok
}

// DrawingBufferWidth returns the value of property "WebGL2RenderingContext.drawingBufferWidth".
//
// The returned bool will be false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferWidth() (GLsizei, bool) {
	var _ok bool
	_ret := bindings.GetWebGL2RenderingContextDrawingBufferWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLsizei(_ret), _ok
}

// DrawingBufferHeight returns the value of property "WebGL2RenderingContext.drawingBufferHeight".
//
// The returned bool will be false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferHeight() (GLsizei, bool) {
	var _ok bool
	_ret := bindings.GetWebGL2RenderingContextDrawingBufferHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return GLsizei(_ret), _ok
}

// DrawingBufferColorSpace returns the value of property "WebGL2RenderingContext.drawingBufferColorSpace".
//
// The returned bool will be false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferColorSpace() (PredefinedColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetWebGL2RenderingContextDrawingBufferColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return PredefinedColorSpace(_ret), _ok
}

// DrawingBufferColorSpace sets the value of property "WebGL2RenderingContext.drawingBufferColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGL2RenderingContext) SetDrawingBufferColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGL2RenderingContextDrawingBufferColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// UnpackColorSpace returns the value of property "WebGL2RenderingContext.unpackColorSpace".
//
// The returned bool will be false if there is no such property.
func (this WebGL2RenderingContext) UnpackColorSpace() (PredefinedColorSpace, bool) {
	var _ok bool
	_ret := bindings.GetWebGL2RenderingContextUnpackColorSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return PredefinedColorSpace(_ret), _ok
}

// UnpackColorSpace sets the value of property "WebGL2RenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGL2RenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGL2RenderingContextUnpackColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// BufferData calls the method "WebGL2RenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferData(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferDataFunc returns the method "WebGL2RenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferDataFunc() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferDataFunc(
			this.Ref(),
		),
	)
}

// BufferData1 calls the method "WebGL2RenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferData1(target GLenum, srcData AllowSharedBufferSource, usage GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferData1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferData1Func returns the method "WebGL2RenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferData1Func() (fn js.Func[func(target GLenum, srcData AllowSharedBufferSource, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData1Func(
			this.Ref(),
		),
	)
}

// BufferSubData calls the method "WebGL2RenderingContext.bufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferSubData(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferSubData(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferSubDataFunc returns the method "WebGL2RenderingContext.bufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferSubDataFunc() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// BufferData2 calls the method "WebGL2RenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferData2(target GLenum, srcData ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferData2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
		uint32(length),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferData2Func returns the method "WebGL2RenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferData2Func() (fn js.Func[func(target GLenum, srcData ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData2Func(
			this.Ref(),
		),
	)
}

// BufferData3 calls the method "WebGL2RenderingContext.bufferData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferData3(target GLenum, srcData ArrayBufferView, usage GLenum, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferData3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferData3Func returns the method "WebGL2RenderingContext.bufferData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferData3Func() (fn js.Func[func(target GLenum, srcData ArrayBufferView, usage GLenum, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData3Func(
			this.Ref(),
		),
	)
}

// BufferSubData1 calls the method "WebGL2RenderingContext.bufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferSubData1(target GLenum, dstByteOffset GLintptr, srcData ArrayBufferView, srcOffset GLuint, length GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferSubData1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(length),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferSubData1Func returns the method "WebGL2RenderingContext.bufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferSubData1Func() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData ArrayBufferView, srcOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubData1Func(
			this.Ref(),
		),
	)
}

// BufferSubData2 calls the method "WebGL2RenderingContext.bufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BufferSubData2(target GLenum, dstByteOffset GLintptr, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBufferSubData2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BufferSubData2Func returns the method "WebGL2RenderingContext.bufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BufferSubData2Func() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubData2Func(
			this.Ref(),
		),
	)
}

// TexImage2D calls the method "WebGL2RenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2DFunc returns the method "WebGL2RenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2DFunc(
			this.Ref(),
		),
	)
}

// TexImage2D1 calls the method "WebGL2RenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2D1Func returns the method "WebGL2RenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D calls the method "WebGL2RenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2DFunc returns the method "WebGL2RenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage2D1 calls the method "WebGL2RenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2D1Func returns the method "WebGL2RenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// TexImage2D2 calls the method "WebGL2RenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage2D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage2D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		uint32(format),
		uint32(typ),
		float64(pboOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2D2Func returns the method "WebGL2RenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage2D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D2Func(
			this.Ref(),
		),
	)
}

// TexImage2D3 calls the method "WebGL2RenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage2D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage2D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2D3Func returns the method "WebGL2RenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage2D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D3Func(
			this.Ref(),
		),
	)
}

// TexImage2D4 calls the method "WebGL2RenderingContext.texImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage2D4(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage2D4(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage2D4Func returns the method "WebGL2RenderingContext.texImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage2D4Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D4Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D2 calls the method "WebGL2RenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage2D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		float64(pboOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2D2Func returns the method "WebGL2RenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D2Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D3 calls the method "WebGL2RenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage2D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2D3Func returns the method "WebGL2RenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D3Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D4 calls the method "WebGL2RenderingContext.texSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D4(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage2D4(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage2D4Func returns the method "WebGL2RenderingContext.texSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage2D4Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D4Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D calls the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		int32(imageSize),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage2DFunc returns the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D1 calls the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(srcLengthOverride),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage2D1Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D2 calls the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage2D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage2D2Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D3 calls the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage2D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage2D3Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage2D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		int32(imageSize),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage2DFunc returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D1 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage2D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(srcLengthOverride),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage2D1Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D2 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage2D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage2D2Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D3 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage2D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage2D3Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage2D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D3Func(
			this.Ref(),
		),
	)
}

// Uniform1fv calls the method "WebGL2RenderingContext.uniform1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fvFunc returns the method "WebGL2RenderingContext.uniform1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fvFunc(
			this.Ref(),
		),
	)
}

// Uniform1fv1 calls the method "WebGL2RenderingContext.uniform1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fv1Func returns the method "WebGL2RenderingContext.uniform1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fv1Func(
			this.Ref(),
		),
	)
}

// Uniform1fv2 calls the method "WebGL2RenderingContext.uniform1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1fv2(location WebGLUniformLocation, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fv2Func returns the method "WebGL2RenderingContext.uniform1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fv2Func(
			this.Ref(),
		),
	)
}

// Uniform2fv calls the method "WebGL2RenderingContext.uniform2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fvFunc returns the method "WebGL2RenderingContext.uniform2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fvFunc(
			this.Ref(),
		),
	)
}

// Uniform2fv1 calls the method "WebGL2RenderingContext.uniform2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fv1Func returns the method "WebGL2RenderingContext.uniform2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fv1Func(
			this.Ref(),
		),
	)
}

// Uniform2fv2 calls the method "WebGL2RenderingContext.uniform2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2fv2(location WebGLUniformLocation, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fv2Func returns the method "WebGL2RenderingContext.uniform2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fv2Func(
			this.Ref(),
		),
	)
}

// Uniform3fv calls the method "WebGL2RenderingContext.uniform3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fvFunc returns the method "WebGL2RenderingContext.uniform3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fvFunc(
			this.Ref(),
		),
	)
}

// Uniform3fv1 calls the method "WebGL2RenderingContext.uniform3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fv1Func returns the method "WebGL2RenderingContext.uniform3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fv1Func(
			this.Ref(),
		),
	)
}

// Uniform3fv2 calls the method "WebGL2RenderingContext.uniform3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3fv2(location WebGLUniformLocation, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fv2Func returns the method "WebGL2RenderingContext.uniform3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fv2Func(
			this.Ref(),
		),
	)
}

// Uniform4fv calls the method "WebGL2RenderingContext.uniform4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fvFunc returns the method "WebGL2RenderingContext.uniform4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fvFunc(
			this.Ref(),
		),
	)
}

// Uniform4fv1 calls the method "WebGL2RenderingContext.uniform4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fv1Func returns the method "WebGL2RenderingContext.uniform4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fv1Func(
			this.Ref(),
		),
	)
}

// Uniform4fv2 calls the method "WebGL2RenderingContext.uniform4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4fv2(location WebGLUniformLocation, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fv2Func returns the method "WebGL2RenderingContext.uniform4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fv2Func(
			this.Ref(),
		),
	)
}

// Uniform1iv calls the method "WebGL2RenderingContext.uniform1iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1ivFunc returns the method "WebGL2RenderingContext.uniform1iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1ivFunc(
			this.Ref(),
		),
	)
}

// Uniform1iv1 calls the method "WebGL2RenderingContext.uniform1iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1iv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1iv1Func returns the method "WebGL2RenderingContext.uniform1iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iv1Func(
			this.Ref(),
		),
	)
}

// Uniform1iv2 calls the method "WebGL2RenderingContext.uniform1iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1iv2(location WebGLUniformLocation, data Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1iv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1iv2Func returns the method "WebGL2RenderingContext.uniform1iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iv2Func(
			this.Ref(),
		),
	)
}

// Uniform2iv calls the method "WebGL2RenderingContext.uniform2iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2ivFunc returns the method "WebGL2RenderingContext.uniform2iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2ivFunc(
			this.Ref(),
		),
	)
}

// Uniform2iv1 calls the method "WebGL2RenderingContext.uniform2iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2iv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2iv1Func returns the method "WebGL2RenderingContext.uniform2iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iv1Func(
			this.Ref(),
		),
	)
}

// Uniform2iv2 calls the method "WebGL2RenderingContext.uniform2iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2iv2(location WebGLUniformLocation, data Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2iv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2iv2Func returns the method "WebGL2RenderingContext.uniform2iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iv2Func(
			this.Ref(),
		),
	)
}

// Uniform3iv calls the method "WebGL2RenderingContext.uniform3iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3ivFunc returns the method "WebGL2RenderingContext.uniform3iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3ivFunc(
			this.Ref(),
		),
	)
}

// Uniform3iv1 calls the method "WebGL2RenderingContext.uniform3iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3iv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3iv1Func returns the method "WebGL2RenderingContext.uniform3iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iv1Func(
			this.Ref(),
		),
	)
}

// Uniform3iv2 calls the method "WebGL2RenderingContext.uniform3iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3iv2(location WebGLUniformLocation, data Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3iv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3iv2Func returns the method "WebGL2RenderingContext.uniform3iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iv2Func(
			this.Ref(),
		),
	)
}

// Uniform4iv calls the method "WebGL2RenderingContext.uniform4iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4iv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4ivFunc returns the method "WebGL2RenderingContext.uniform4iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4ivFunc(
			this.Ref(),
		),
	)
}

// Uniform4iv1 calls the method "WebGL2RenderingContext.uniform4iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4iv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4iv1Func returns the method "WebGL2RenderingContext.uniform4iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iv1Func(
			this.Ref(),
		),
	)
}

// Uniform4iv2 calls the method "WebGL2RenderingContext.uniform4iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4iv2(location WebGLUniformLocation, data Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4iv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4iv2Func returns the method "WebGL2RenderingContext.uniform4iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv calls the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv1 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv2 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv calls the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv1 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv2 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv calls the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv1 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv2 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fv2Func(
			this.Ref(),
		),
	)
}

// ReadPixels calls the method "WebGL2RenderingContext.readPixels".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextReadPixels(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadPixelsFunc returns the method "WebGL2RenderingContext.readPixels".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ReadPixelsFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixelsFunc(
			this.Ref(),
		),
	)
}

// ReadPixels1 calls the method "WebGL2RenderingContext.readPixels".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ReadPixels1(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextReadPixels1(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadPixels1Func returns the method "WebGL2RenderingContext.readPixels".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ReadPixels1Func() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixels1Func(
			this.Ref(),
		),
	)
}

// ReadPixels2 calls the method "WebGL2RenderingContext.readPixels".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ReadPixels2(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData ArrayBufferView, dstOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextReadPixels2(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
		uint32(dstOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadPixels2Func returns the method "WebGL2RenderingContext.readPixels".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ReadPixels2Func() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData ArrayBufferView, dstOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixels2Func(
			this.Ref(),
		),
	)
}

// CopyBufferSubData calls the method "WebGL2RenderingContext.copyBufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CopyBufferSubData(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCopyBufferSubData(
		this.Ref(), js.Pointer(&_ok),
		uint32(readTarget),
		uint32(writeTarget),
		float64(readOffset),
		float64(writeOffset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyBufferSubDataFunc returns the method "WebGL2RenderingContext.copyBufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CopyBufferSubDataFunc() (fn js.Func[func(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// GetBufferSubData calls the method "WebGL2RenderingContext.getBufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubData(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView, dstOffset GLuint, length GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetBufferSubData(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
		uint32(length),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetBufferSubDataFunc returns the method "WebGL2RenderingContext.getBufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubDataFunc() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView, dstOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// GetBufferSubData1 calls the method "WebGL2RenderingContext.getBufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubData1(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView, dstOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetBufferSubData1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetBufferSubData1Func returns the method "WebGL2RenderingContext.getBufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubData1Func() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView, dstOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubData1Func(
			this.Ref(),
		),
	)
}

// GetBufferSubData2 calls the method "WebGL2RenderingContext.getBufferSubData".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubData2(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetBufferSubData2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetBufferSubData2Func returns the method "WebGL2RenderingContext.getBufferSubData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetBufferSubData2Func() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubData2Func(
			this.Ref(),
		),
	)
}

// BlitFramebuffer calls the method "WebGL2RenderingContext.blitFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlitFramebuffer(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlitFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		int32(srcX0),
		int32(srcY0),
		int32(srcX1),
		int32(srcY1),
		int32(dstX0),
		int32(dstY0),
		int32(dstX1),
		int32(dstY1),
		uint32(mask),
		uint32(filter),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlitFramebufferFunc returns the method "WebGL2RenderingContext.blitFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlitFramebufferFunc() (fn js.Func[func(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlitFramebufferFunc(
			this.Ref(),
		),
	)
}

// FramebufferTextureLayer calls the method "WebGL2RenderingContext.framebufferTextureLayer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) FramebufferTextureLayer(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFramebufferTextureLayer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(layer),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferTextureLayerFunc returns the method "WebGL2RenderingContext.framebufferTextureLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FramebufferTextureLayerFunc() (fn js.Func[func(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferTextureLayerFunc(
			this.Ref(),
		),
	)
}

// InvalidateFramebuffer calls the method "WebGL2RenderingContext.invalidateFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) InvalidateFramebuffer(target GLenum, attachments js.Array[GLenum]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextInvalidateFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		attachments.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InvalidateFramebufferFunc returns the method "WebGL2RenderingContext.invalidateFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) InvalidateFramebufferFunc() (fn js.Func[func(target GLenum, attachments js.Array[GLenum])]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextInvalidateFramebufferFunc(
			this.Ref(),
		),
	)
}

// InvalidateSubFramebuffer calls the method "WebGL2RenderingContext.invalidateSubFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) InvalidateSubFramebuffer(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextInvalidateSubFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		attachments.Ref(),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InvalidateSubFramebufferFunc returns the method "WebGL2RenderingContext.invalidateSubFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) InvalidateSubFramebufferFunc() (fn js.Func[func(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextInvalidateSubFramebufferFunc(
			this.Ref(),
		),
	)
}

// ReadBuffer calls the method "WebGL2RenderingContext.readBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ReadBuffer(src GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextReadBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(src),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadBufferFunc returns the method "WebGL2RenderingContext.readBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ReadBufferFunc() (fn js.Func[func(src GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadBufferFunc(
			this.Ref(),
		),
	)
}

// GetInternalformatParameter calls the method "WebGL2RenderingContext.getInternalformatParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetInternalformatParameter(target GLenum, internalformat GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetInternalformatParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(internalformat),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetInternalformatParameterFunc returns the method "WebGL2RenderingContext.getInternalformatParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetInternalformatParameterFunc() (fn js.Func[func(target GLenum, internalformat GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetInternalformatParameterFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorageMultisample calls the method "WebGL2RenderingContext.renderbufferStorageMultisample".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) RenderbufferStorageMultisample(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextRenderbufferStorageMultisample(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(samples),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RenderbufferStorageMultisampleFunc returns the method "WebGL2RenderingContext.renderbufferStorageMultisample".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) RenderbufferStorageMultisampleFunc() (fn js.Func[func(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextRenderbufferStorageMultisampleFunc(
			this.Ref(),
		),
	)
}

// TexStorage2D calls the method "WebGL2RenderingContext.texStorage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexStorage2D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexStorage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexStorage2DFunc returns the method "WebGL2RenderingContext.texStorage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexStorage2DFunc() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexStorage2DFunc(
			this.Ref(),
		),
	)
}

// TexStorage3D calls the method "WebGL2RenderingContext.texStorage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexStorage3D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexStorage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexStorage3DFunc returns the method "WebGL2RenderingContext.texStorage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexStorage3DFunc() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexStorage3DFunc(
			this.Ref(),
		),
	)
}

// TexImage3D calls the method "WebGL2RenderingContext.texImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage3D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		uint32(format),
		uint32(typ),
		float64(pboOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage3DFunc returns the method "WebGL2RenderingContext.texImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage3DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3DFunc(
			this.Ref(),
		),
	)
}

// TexImage3D1 calls the method "WebGL2RenderingContext.texImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage3D1(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage3D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage3D1Func returns the method "WebGL2RenderingContext.texImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage3D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D1Func(
			this.Ref(),
		),
	)
}

// TexImage3D2 calls the method "WebGL2RenderingContext.texImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage3D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage3D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage3D2Func returns the method "WebGL2RenderingContext.texImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage3D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D2Func(
			this.Ref(),
		),
	)
}

// TexImage3D3 calls the method "WebGL2RenderingContext.texImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexImage3D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexImage3D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexImage3D3Func returns the method "WebGL2RenderingContext.texImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexImage3D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D3Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D calls the method "WebGL2RenderingContext.texSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		uint32(typ),
		float64(pboOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage3DFunc returns the method "WebGL2RenderingContext.texSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage3D1 calls the method "WebGL2RenderingContext.texSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage3D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage3D1Func returns the method "WebGL2RenderingContext.texSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D2 calls the method "WebGL2RenderingContext.texSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage3D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage3D2Func returns the method "WebGL2RenderingContext.texSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D2Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D3 calls the method "WebGL2RenderingContext.texSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexSubImage3D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		uint32(typ),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexSubImage3D3Func returns the method "WebGL2RenderingContext.texSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexSubImage3D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D3Func(
			this.Ref(),
		),
	)
}

// CopyTexSubImage3D calls the method "WebGL2RenderingContext.copyTexSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CopyTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCopyTexSubImage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTexSubImage3DFunc returns the method "WebGL2RenderingContext.copyTexSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CopyTexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D calls the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		int32(imageSize),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage3DFunc returns the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D1 calls the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage3D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(srcLengthOverride),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage3D1Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D2 calls the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage3D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage3D2Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D3 calls the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexImage3D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexImage3D3Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexImage3D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage3D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		int32(imageSize),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage3DFunc returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D1 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage3D1(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(srcLengthOverride),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage3D1Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D2 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage3D2(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		srcData.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage3D2Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D3 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompressedTexSubImage3D3(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(zoffset),
		int32(width),
		int32(height),
		int32(depth),
		uint32(format),
		srcData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompressedTexSubImage3D3Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompressedTexSubImage3D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D3Func(
			this.Ref(),
		),
	)
}

// GetFragDataLocation calls the method "WebGL2RenderingContext.getFragDataLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetFragDataLocation(program WebGLProgram, name js.String) (GLint, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetFragDataLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		name.Ref(),
	)

	return GLint(_ret), _ok
}

// GetFragDataLocationFunc returns the method "WebGL2RenderingContext.getFragDataLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetFragDataLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetFragDataLocationFunc(
			this.Ref(),
		),
	)
}

// Uniform1ui calls the method "WebGL2RenderingContext.uniform1ui".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1ui(location WebGLUniformLocation, v0 GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1ui(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		uint32(v0),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1uiFunc returns the method "WebGL2RenderingContext.uniform1ui".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiFunc(
			this.Ref(),
		),
	)
}

// Uniform2ui calls the method "WebGL2RenderingContext.uniform2ui".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2ui(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		uint32(v0),
		uint32(v1),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2uiFunc returns the method "WebGL2RenderingContext.uniform2ui".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiFunc(
			this.Ref(),
		),
	)
}

// Uniform3ui calls the method "WebGL2RenderingContext.uniform3ui".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3ui(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3uiFunc returns the method "WebGL2RenderingContext.uniform3ui".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiFunc(
			this.Ref(),
		),
	)
}

// Uniform4ui calls the method "WebGL2RenderingContext.uniform4ui".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4ui(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
		uint32(v3),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4uiFunc returns the method "WebGL2RenderingContext.uniform4ui".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiFunc(
			this.Ref(),
		),
	)
}

// Uniform1uiv calls the method "WebGL2RenderingContext.uniform1uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1uiv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1uivFunc returns the method "WebGL2RenderingContext.uniform1uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uivFunc(
			this.Ref(),
		),
	)
}

// Uniform1uiv1 calls the method "WebGL2RenderingContext.uniform1uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1uiv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1uiv1Func returns the method "WebGL2RenderingContext.uniform1uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform1uiv2 calls the method "WebGL2RenderingContext.uniform1uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiv2(location WebGLUniformLocation, data Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1uiv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1uiv2Func returns the method "WebGL2RenderingContext.uniform1uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform2uiv calls the method "WebGL2RenderingContext.uniform2uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2uiv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2uivFunc returns the method "WebGL2RenderingContext.uniform2uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uivFunc(
			this.Ref(),
		),
	)
}

// Uniform2uiv1 calls the method "WebGL2RenderingContext.uniform2uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2uiv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2uiv1Func returns the method "WebGL2RenderingContext.uniform2uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform2uiv2 calls the method "WebGL2RenderingContext.uniform2uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiv2(location WebGLUniformLocation, data Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2uiv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2uiv2Func returns the method "WebGL2RenderingContext.uniform2uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform3uiv calls the method "WebGL2RenderingContext.uniform3uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3uiv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3uivFunc returns the method "WebGL2RenderingContext.uniform3uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uivFunc(
			this.Ref(),
		),
	)
}

// Uniform3uiv1 calls the method "WebGL2RenderingContext.uniform3uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3uiv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3uiv1Func returns the method "WebGL2RenderingContext.uniform3uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform3uiv2 calls the method "WebGL2RenderingContext.uniform3uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiv2(location WebGLUniformLocation, data Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3uiv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3uiv2Func returns the method "WebGL2RenderingContext.uniform3uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform4uiv calls the method "WebGL2RenderingContext.uniform4uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4uiv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4uivFunc returns the method "WebGL2RenderingContext.uniform4uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uivFunc(
			this.Ref(),
		),
	)
}

// Uniform4uiv1 calls the method "WebGL2RenderingContext.uniform4uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4uiv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4uiv1Func returns the method "WebGL2RenderingContext.uniform4uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform4uiv2 calls the method "WebGL2RenderingContext.uniform4uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiv2(location WebGLUniformLocation, data Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4uiv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4uiv2Func returns the method "WebGL2RenderingContext.uniform4uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x2fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x2fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x2fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x2fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x2fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x3fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x3fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x3fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x3fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix4x3fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix4x3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix4x3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x4fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix2x4fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix2x4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix2x4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x4fv(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x4fv1(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformMatrix3x4fv2(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformMatrix3x4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformMatrix3x4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fv2Func(
			this.Ref(),
		),
	)
}

// VertexAttribI4i calls the method "WebGL2RenderingContext.vertexAttribI4i".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4i(index GLuint, x GLint, y GLint, z GLint, w GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribI4i(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribI4iFunc returns the method "WebGL2RenderingContext.vertexAttribI4i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4iFunc() (fn js.Func[func(index GLuint, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4iFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4iv calls the method "WebGL2RenderingContext.vertexAttribI4iv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4iv(index GLuint, values Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribI4iv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribI4ivFunc returns the method "WebGL2RenderingContext.vertexAttribI4iv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4ivFunc() (fn js.Func[func(index GLuint, values Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4ivFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4ui calls the method "WebGL2RenderingContext.vertexAttribI4ui".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4ui(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribI4ui(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(x),
		uint32(y),
		uint32(z),
		uint32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribI4uiFunc returns the method "WebGL2RenderingContext.vertexAttribI4ui".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4uiFunc() (fn js.Func[func(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4uiFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4uiv calls the method "WebGL2RenderingContext.vertexAttribI4uiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4uiv(index GLuint, values Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribI4uiv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribI4uivFunc returns the method "WebGL2RenderingContext.vertexAttribI4uiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribI4uivFunc() (fn js.Func[func(index GLuint, values Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4uivFunc(
			this.Ref(),
		),
	)
}

// VertexAttribIPointer calls the method "WebGL2RenderingContext.vertexAttribIPointer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribIPointer(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribIPointer(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		int32(size),
		uint32(typ),
		int32(stride),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribIPointerFunc returns the method "WebGL2RenderingContext.vertexAttribIPointer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribIPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribIPointerFunc(
			this.Ref(),
		),
	)
}

// VertexAttribDivisor calls the method "WebGL2RenderingContext.vertexAttribDivisor".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribDivisor(index GLuint, divisor GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribDivisor(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(divisor),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribDivisorFunc returns the method "WebGL2RenderingContext.vertexAttribDivisor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribDivisorFunc() (fn js.Func[func(index GLuint, divisor GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribDivisorFunc(
			this.Ref(),
		),
	)
}

// DrawArraysInstanced calls the method "WebGL2RenderingContext.drawArraysInstanced".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawArraysInstanced(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawArraysInstancedFunc returns the method "WebGL2RenderingContext.drawArraysInstanced".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawArraysInstancedFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawArraysInstancedFunc(
			this.Ref(),
		),
	)
}

// DrawElementsInstanced calls the method "WebGL2RenderingContext.drawElementsInstanced".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawElementsInstanced(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawElementsInstanced(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawElementsInstancedFunc returns the method "WebGL2RenderingContext.drawElementsInstanced".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawElementsInstancedFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawElementsInstancedFunc(
			this.Ref(),
		),
	)
}

// DrawRangeElements calls the method "WebGL2RenderingContext.drawRangeElements".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawRangeElements(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawRangeElements(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		uint32(start),
		uint32(end),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawRangeElementsFunc returns the method "WebGL2RenderingContext.drawRangeElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawRangeElementsFunc() (fn js.Func[func(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawRangeElementsFunc(
			this.Ref(),
		),
	)
}

// DrawBuffers calls the method "WebGL2RenderingContext.drawBuffers".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawBuffers(buffers js.Array[GLenum]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawBuffers(
		this.Ref(), js.Pointer(&_ok),
		buffers.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawBuffersFunc returns the method "WebGL2RenderingContext.drawBuffers".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawBuffersFunc() (fn js.Func[func(buffers js.Array[GLenum])]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawBuffersFunc(
			this.Ref(),
		),
	)
}

// ClearBufferfv calls the method "WebGL2RenderingContext.clearBufferfv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfv(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferfv(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferfvFunc returns the method "WebGL2RenderingContext.clearBufferfv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfvFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfvFunc(
			this.Ref(),
		),
	)
}

// ClearBufferfv1 calls the method "WebGL2RenderingContext.clearBufferfv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfv1(buffer GLenum, drawbuffer GLint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferfv1(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferfv1Func returns the method "WebGL2RenderingContext.clearBufferfv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferiv calls the method "WebGL2RenderingContext.clearBufferiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferiv(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferiv(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferivFunc returns the method "WebGL2RenderingContext.clearBufferiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferivFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferivFunc(
			this.Ref(),
		),
	)
}

// ClearBufferiv1 calls the method "WebGL2RenderingContext.clearBufferiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferiv1(buffer GLenum, drawbuffer GLint, values Int32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferiv1(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferiv1Func returns the method "WebGL2RenderingContext.clearBufferiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferiv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferiv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferuiv calls the method "WebGL2RenderingContext.clearBufferuiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferuiv(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferuiv(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferuivFunc returns the method "WebGL2RenderingContext.clearBufferuiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferuivFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferuivFunc(
			this.Ref(),
		),
	)
}

// ClearBufferuiv1 calls the method "WebGL2RenderingContext.clearBufferuiv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferuiv1(buffer GLenum, drawbuffer GLint, values Uint32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferuiv1(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferuiv1Func returns the method "WebGL2RenderingContext.clearBufferuiv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferuiv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferuiv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferfi calls the method "WebGL2RenderingContext.clearBufferfi".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfi(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearBufferfi(
		this.Ref(), js.Pointer(&_ok),
		uint32(buffer),
		int32(drawbuffer),
		float32(depth),
		int32(stencil),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferfiFunc returns the method "WebGL2RenderingContext.clearBufferfi".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearBufferfiFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfiFunc(
			this.Ref(),
		),
	)
}

// CreateQuery calls the method "WebGL2RenderingContext.createQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateQuery() (WebGLQuery, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateQuery(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLQuery{}.FromRef(_ret), _ok
}

// CreateQueryFunc returns the method "WebGL2RenderingContext.createQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateQueryFunc() (fn js.Func[func() WebGLQuery]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateQueryFunc(
			this.Ref(),
		),
	)
}

// DeleteQuery calls the method "WebGL2RenderingContext.deleteQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteQuery(query WebGLQuery) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteQuery(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteQueryFunc returns the method "WebGL2RenderingContext.deleteQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteQueryFunc() (fn js.Func[func(query WebGLQuery)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteQueryFunc(
			this.Ref(),
		),
	)
}

// IsQuery calls the method "WebGL2RenderingContext.isQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsQuery(query WebGLQuery) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsQuery(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return _ret == js.True, _ok
}

// IsQueryFunc returns the method "WebGL2RenderingContext.isQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsQueryFunc() (fn js.Func[func(query WebGLQuery) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsQueryFunc(
			this.Ref(),
		),
	)
}

// BeginQuery calls the method "WebGL2RenderingContext.beginQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BeginQuery(target GLenum, query WebGLQuery) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBeginQuery(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		query.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginQueryFunc returns the method "WebGL2RenderingContext.beginQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BeginQueryFunc() (fn js.Func[func(target GLenum, query WebGLQuery)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBeginQueryFunc(
			this.Ref(),
		),
	)
}

// EndQuery calls the method "WebGL2RenderingContext.endQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) EndQuery(target GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextEndQuery(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndQueryFunc returns the method "WebGL2RenderingContext.endQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) EndQueryFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEndQueryFunc(
			this.Ref(),
		),
	)
}

// GetQuery calls the method "WebGL2RenderingContext.getQuery".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetQuery(target GLenum, pname GLenum) (WebGLQuery, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetQuery(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return WebGLQuery{}.FromRef(_ret), _ok
}

// GetQueryFunc returns the method "WebGL2RenderingContext.getQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetQueryFunc() (fn js.Func[func(target GLenum, pname GLenum) WebGLQuery]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetQueryFunc(
			this.Ref(),
		),
	)
}

// GetQueryParameter calls the method "WebGL2RenderingContext.getQueryParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetQueryParameter(query WebGLQuery, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetQueryParameter(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetQueryParameterFunc returns the method "WebGL2RenderingContext.getQueryParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetQueryParameterFunc() (fn js.Func[func(query WebGLQuery, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetQueryParameterFunc(
			this.Ref(),
		),
	)
}

// CreateSampler calls the method "WebGL2RenderingContext.createSampler".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateSampler() (WebGLSampler, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateSampler(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLSampler{}.FromRef(_ret), _ok
}

// CreateSamplerFunc returns the method "WebGL2RenderingContext.createSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateSamplerFunc() (fn js.Func[func() WebGLSampler]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateSamplerFunc(
			this.Ref(),
		),
	)
}

// DeleteSampler calls the method "WebGL2RenderingContext.deleteSampler".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteSampler(sampler WebGLSampler) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteSampler(
		this.Ref(), js.Pointer(&_ok),
		sampler.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteSamplerFunc returns the method "WebGL2RenderingContext.deleteSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteSamplerFunc() (fn js.Func[func(sampler WebGLSampler)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteSamplerFunc(
			this.Ref(),
		),
	)
}

// IsSampler calls the method "WebGL2RenderingContext.isSampler".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsSampler(sampler WebGLSampler) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsSampler(
		this.Ref(), js.Pointer(&_ok),
		sampler.Ref(),
	)

	return _ret == js.True, _ok
}

// IsSamplerFunc returns the method "WebGL2RenderingContext.isSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsSamplerFunc() (fn js.Func[func(sampler WebGLSampler) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsSamplerFunc(
			this.Ref(),
		),
	)
}

// BindSampler calls the method "WebGL2RenderingContext.bindSampler".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindSampler(unit GLuint, sampler WebGLSampler) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindSampler(
		this.Ref(), js.Pointer(&_ok),
		uint32(unit),
		sampler.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindSamplerFunc returns the method "WebGL2RenderingContext.bindSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindSamplerFunc() (fn js.Func[func(unit GLuint, sampler WebGLSampler)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindSamplerFunc(
			this.Ref(),
		),
	)
}

// SamplerParameteri calls the method "WebGL2RenderingContext.samplerParameteri".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) SamplerParameteri(sampler WebGLSampler, pname GLenum, param GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextSamplerParameteri(
		this.Ref(), js.Pointer(&_ok),
		sampler.Ref(),
		uint32(pname),
		int32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SamplerParameteriFunc returns the method "WebGL2RenderingContext.samplerParameteri".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) SamplerParameteriFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSamplerParameteriFunc(
			this.Ref(),
		),
	)
}

// SamplerParameterf calls the method "WebGL2RenderingContext.samplerParameterf".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) SamplerParameterf(sampler WebGLSampler, pname GLenum, param GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextSamplerParameterf(
		this.Ref(), js.Pointer(&_ok),
		sampler.Ref(),
		uint32(pname),
		float32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SamplerParameterfFunc returns the method "WebGL2RenderingContext.samplerParameterf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) SamplerParameterfFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSamplerParameterfFunc(
			this.Ref(),
		),
	)
}

// GetSamplerParameter calls the method "WebGL2RenderingContext.getSamplerParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetSamplerParameter(sampler WebGLSampler, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetSamplerParameter(
		this.Ref(), js.Pointer(&_ok),
		sampler.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetSamplerParameterFunc returns the method "WebGL2RenderingContext.getSamplerParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetSamplerParameterFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSamplerParameterFunc(
			this.Ref(),
		),
	)
}

// FenceSync calls the method "WebGL2RenderingContext.fenceSync".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) FenceSync(condition GLenum, flags GLbitfield) (WebGLSync, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFenceSync(
		this.Ref(), js.Pointer(&_ok),
		uint32(condition),
		uint32(flags),
	)

	return WebGLSync{}.FromRef(_ret), _ok
}

// FenceSyncFunc returns the method "WebGL2RenderingContext.fenceSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FenceSyncFunc() (fn js.Func[func(condition GLenum, flags GLbitfield) WebGLSync]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFenceSyncFunc(
			this.Ref(),
		),
	)
}

// IsSync calls the method "WebGL2RenderingContext.isSync".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsSync(sync WebGLSync) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsSync(
		this.Ref(), js.Pointer(&_ok),
		sync.Ref(),
	)

	return _ret == js.True, _ok
}

// IsSyncFunc returns the method "WebGL2RenderingContext.isSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsSyncFunc() (fn js.Func[func(sync WebGLSync) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsSyncFunc(
			this.Ref(),
		),
	)
}

// DeleteSync calls the method "WebGL2RenderingContext.deleteSync".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteSync(sync WebGLSync) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteSync(
		this.Ref(), js.Pointer(&_ok),
		sync.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteSyncFunc returns the method "WebGL2RenderingContext.deleteSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteSyncFunc() (fn js.Func[func(sync WebGLSync)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteSyncFunc(
			this.Ref(),
		),
	)
}

// ClientWaitSync calls the method "WebGL2RenderingContext.clientWaitSync".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClientWaitSync(sync WebGLSync, flags GLbitfield, timeout GLuint64) (GLenum, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClientWaitSync(
		this.Ref(), js.Pointer(&_ok),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return GLenum(_ret), _ok
}

// ClientWaitSyncFunc returns the method "WebGL2RenderingContext.clientWaitSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClientWaitSyncFunc() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLuint64) GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClientWaitSyncFunc(
			this.Ref(),
		),
	)
}

// WaitSync calls the method "WebGL2RenderingContext.waitSync".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) WaitSync(sync WebGLSync, flags GLbitfield, timeout GLint64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextWaitSync(
		this.Ref(), js.Pointer(&_ok),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WaitSyncFunc returns the method "WebGL2RenderingContext.waitSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) WaitSyncFunc() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLint64)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextWaitSyncFunc(
			this.Ref(),
		),
	)
}

// GetSyncParameter calls the method "WebGL2RenderingContext.getSyncParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetSyncParameter(sync WebGLSync, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetSyncParameter(
		this.Ref(), js.Pointer(&_ok),
		sync.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetSyncParameterFunc returns the method "WebGL2RenderingContext.getSyncParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetSyncParameterFunc() (fn js.Func[func(sync WebGLSync, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSyncParameterFunc(
			this.Ref(),
		),
	)
}

// CreateTransformFeedback calls the method "WebGL2RenderingContext.createTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateTransformFeedback() (WebGLTransformFeedback, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLTransformFeedback{}.FromRef(_ret), _ok
}

// CreateTransformFeedbackFunc returns the method "WebGL2RenderingContext.createTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateTransformFeedbackFunc() (fn js.Func[func() WebGLTransformFeedback]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// DeleteTransformFeedback calls the method "WebGL2RenderingContext.deleteTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteTransformFeedback(tf WebGLTransformFeedback) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
		tf.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteTransformFeedbackFunc returns the method "WebGL2RenderingContext.deleteTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteTransformFeedbackFunc() (fn js.Func[func(tf WebGLTransformFeedback)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// IsTransformFeedback calls the method "WebGL2RenderingContext.isTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsTransformFeedback(tf WebGLTransformFeedback) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
		tf.Ref(),
	)

	return _ret == js.True, _ok
}

// IsTransformFeedbackFunc returns the method "WebGL2RenderingContext.isTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsTransformFeedbackFunc() (fn js.Func[func(tf WebGLTransformFeedback) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// BindTransformFeedback calls the method "WebGL2RenderingContext.bindTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindTransformFeedback(target GLenum, tf WebGLTransformFeedback) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		tf.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindTransformFeedbackFunc returns the method "WebGL2RenderingContext.bindTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindTransformFeedbackFunc() (fn js.Func[func(target GLenum, tf WebGLTransformFeedback)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// BeginTransformFeedback calls the method "WebGL2RenderingContext.beginTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BeginTransformFeedback(primitiveMode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBeginTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
		uint32(primitiveMode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginTransformFeedbackFunc returns the method "WebGL2RenderingContext.beginTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BeginTransformFeedbackFunc() (fn js.Func[func(primitiveMode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBeginTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// EndTransformFeedback calls the method "WebGL2RenderingContext.endTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) EndTransformFeedback() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextEndTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndTransformFeedbackFunc returns the method "WebGL2RenderingContext.endTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) EndTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEndTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// TransformFeedbackVaryings calls the method "WebGL2RenderingContext.transformFeedbackVaryings".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TransformFeedbackVaryings(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTransformFeedbackVaryings(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		varyings.Ref(),
		uint32(bufferMode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TransformFeedbackVaryingsFunc returns the method "WebGL2RenderingContext.transformFeedbackVaryings".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TransformFeedbackVaryingsFunc() (fn js.Func[func(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTransformFeedbackVaryingsFunc(
			this.Ref(),
		),
	)
}

// GetTransformFeedbackVarying calls the method "WebGL2RenderingContext.getTransformFeedbackVarying".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetTransformFeedbackVarying(program WebGLProgram, index GLuint) (WebGLActiveInfo, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetTransformFeedbackVarying(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
	)

	return WebGLActiveInfo{}.FromRef(_ret), _ok
}

// GetTransformFeedbackVaryingFunc returns the method "WebGL2RenderingContext.getTransformFeedbackVarying".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetTransformFeedbackVaryingFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetTransformFeedbackVaryingFunc(
			this.Ref(),
		),
	)
}

// PauseTransformFeedback calls the method "WebGL2RenderingContext.pauseTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) PauseTransformFeedback() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextPauseTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseTransformFeedbackFunc returns the method "WebGL2RenderingContext.pauseTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) PauseTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPauseTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// ResumeTransformFeedback calls the method "WebGL2RenderingContext.resumeTransformFeedback".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ResumeTransformFeedback() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextResumeTransformFeedback(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResumeTransformFeedbackFunc returns the method "WebGL2RenderingContext.resumeTransformFeedback".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ResumeTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextResumeTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// BindBufferBase calls the method "WebGL2RenderingContext.bindBufferBase".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindBufferBase(target GLenum, index GLuint, buffer WebGLBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindBufferBase(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(index),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindBufferBaseFunc returns the method "WebGL2RenderingContext.bindBufferBase".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindBufferBaseFunc() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferBaseFunc(
			this.Ref(),
		),
	)
}

// BindBufferRange calls the method "WebGL2RenderingContext.bindBufferRange".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindBufferRange(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindBufferRange(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(index),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindBufferRangeFunc returns the method "WebGL2RenderingContext.bindBufferRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindBufferRangeFunc() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferRangeFunc(
			this.Ref(),
		),
	)
}

// GetIndexedParameter calls the method "WebGL2RenderingContext.getIndexedParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetIndexedParameter(target GLenum, index GLuint) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetIndexedParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(index),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetIndexedParameterFunc returns the method "WebGL2RenderingContext.getIndexedParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetIndexedParameterFunc() (fn js.Func[func(target GLenum, index GLuint) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetIndexedParameterFunc(
			this.Ref(),
		),
	)
}

// GetUniformIndices calls the method "WebGL2RenderingContext.getUniformIndices".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetUniformIndices(program WebGLProgram, uniformNames js.Array[js.String]) (js.Array[GLuint], bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetUniformIndices(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uniformNames.Ref(),
	)

	return js.Array[GLuint]{}.FromRef(_ret), _ok
}

// GetUniformIndicesFunc returns the method "WebGL2RenderingContext.getUniformIndices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetUniformIndicesFunc() (fn js.Func[func(program WebGLProgram, uniformNames js.Array[js.String]) js.Array[GLuint]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformIndicesFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniforms calls the method "WebGL2RenderingContext.getActiveUniforms".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniforms(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetActiveUniforms(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uniformIndices.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetActiveUniformsFunc returns the method "WebGL2RenderingContext.getActiveUniforms".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformsFunc() (fn js.Func[func(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformsFunc(
			this.Ref(),
		),
	)
}

// GetUniformBlockIndex calls the method "WebGL2RenderingContext.getUniformBlockIndex".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetUniformBlockIndex(program WebGLProgram, uniformBlockName js.String) (GLuint, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetUniformBlockIndex(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uniformBlockName.Ref(),
	)

	return GLuint(_ret), _ok
}

// GetUniformBlockIndexFunc returns the method "WebGL2RenderingContext.getUniformBlockIndex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetUniformBlockIndexFunc() (fn js.Func[func(program WebGLProgram, uniformBlockName js.String) GLuint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformBlockIndexFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniformBlockParameter calls the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformBlockParameter(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetActiveUniformBlockParameterFunc returns the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformBlockParameterFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformBlockParameterFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniformBlockName calls the method "WebGL2RenderingContext.getActiveUniformBlockName".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformBlockName(program WebGLProgram, uniformBlockIndex GLuint) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetActiveUniformBlockName(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(uniformBlockIndex),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetActiveUniformBlockNameFunc returns the method "WebGL2RenderingContext.getActiveUniformBlockName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformBlockNameFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformBlockNameFunc(
			this.Ref(),
		),
	)
}

// UniformBlockBinding calls the method "WebGL2RenderingContext.uniformBlockBinding".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UniformBlockBinding(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniformBlockBinding(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(uniformBlockBinding),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UniformBlockBindingFunc returns the method "WebGL2RenderingContext.uniformBlockBinding".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UniformBlockBindingFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformBlockBindingFunc(
			this.Ref(),
		),
	)
}

// CreateVertexArray calls the method "WebGL2RenderingContext.createVertexArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateVertexArray() (WebGLVertexArrayObject, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateVertexArray(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLVertexArrayObject{}.FromRef(_ret), _ok
}

// CreateVertexArrayFunc returns the method "WebGL2RenderingContext.createVertexArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateVertexArrayFunc() (fn js.Func[func() WebGLVertexArrayObject]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateVertexArrayFunc(
			this.Ref(),
		),
	)
}

// DeleteVertexArray calls the method "WebGL2RenderingContext.deleteVertexArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteVertexArray(vertexArray WebGLVertexArrayObject) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteVertexArray(
		this.Ref(), js.Pointer(&_ok),
		vertexArray.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteVertexArrayFunc returns the method "WebGL2RenderingContext.deleteVertexArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteVertexArrayFunc() (fn js.Func[func(vertexArray WebGLVertexArrayObject)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteVertexArrayFunc(
			this.Ref(),
		),
	)
}

// IsVertexArray calls the method "WebGL2RenderingContext.isVertexArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsVertexArray(vertexArray WebGLVertexArrayObject) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsVertexArray(
		this.Ref(), js.Pointer(&_ok),
		vertexArray.Ref(),
	)

	return _ret == js.True, _ok
}

// IsVertexArrayFunc returns the method "WebGL2RenderingContext.isVertexArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsVertexArrayFunc() (fn js.Func[func(vertexArray WebGLVertexArrayObject) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsVertexArrayFunc(
			this.Ref(),
		),
	)
}

// BindVertexArray calls the method "WebGL2RenderingContext.bindVertexArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindVertexArray(array WebGLVertexArrayObject) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindVertexArray(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindVertexArrayFunc returns the method "WebGL2RenderingContext.bindVertexArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindVertexArrayFunc() (fn js.Func[func(array WebGLVertexArrayObject)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindVertexArrayFunc(
			this.Ref(),
		),
	)
}

// GetContextAttributes calls the method "WebGL2RenderingContext.getContextAttributes".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetContextAttributes() (WebGLContextAttributes, bool) {
	var _ret WebGLContextAttributes
	_ok := js.True == bindings.CallWebGL2RenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetContextAttributesFunc returns the method "WebGL2RenderingContext.getContextAttributes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetContextAttributesFunc() (fn js.Func[func() WebGLContextAttributes]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "WebGL2RenderingContext.isContextLost".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsContextLost() (bool, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsContextLost(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsContextLostFunc returns the method "WebGL2RenderingContext.isContextLost".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsContextLostFunc(
			this.Ref(),
		),
	)
}

// GetSupportedExtensions calls the method "WebGL2RenderingContext.getSupportedExtensions".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetSupportedExtensions() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// GetSupportedExtensionsFunc returns the method "WebGL2RenderingContext.getSupportedExtensions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetSupportedExtensionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSupportedExtensionsFunc(
			this.Ref(),
		),
	)
}

// GetExtension calls the method "WebGL2RenderingContext.getExtension".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetExtension(name js.String) (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetExtension(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// GetExtensionFunc returns the method "WebGL2RenderingContext.getExtension".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetExtensionFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetExtensionFunc(
			this.Ref(),
		),
	)
}

// ActiveTexture calls the method "WebGL2RenderingContext.activeTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ActiveTexture(texture GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextActiveTexture(
		this.Ref(), js.Pointer(&_ok),
		uint32(texture),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ActiveTextureFunc returns the method "WebGL2RenderingContext.activeTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ActiveTextureFunc() (fn js.Func[func(texture GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextActiveTextureFunc(
			this.Ref(),
		),
	)
}

// AttachShader calls the method "WebGL2RenderingContext.attachShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextAttachShader(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AttachShaderFunc returns the method "WebGL2RenderingContext.attachShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) AttachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextAttachShaderFunc(
			this.Ref(),
		),
	)
}

// BindAttribLocation calls the method "WebGL2RenderingContext.bindAttribLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindAttribLocationFunc returns the method "WebGL2RenderingContext.bindAttribLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindAttribLocationFunc() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindAttribLocationFunc(
			this.Ref(),
		),
	)
}

// BindBuffer calls the method "WebGL2RenderingContext.bindBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindBufferFunc returns the method "WebGL2RenderingContext.bindBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindBufferFunc() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferFunc(
			this.Ref(),
		),
	)
}

// BindFramebuffer calls the method "WebGL2RenderingContext.bindFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		framebuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindFramebufferFunc returns the method "WebGL2RenderingContext.bindFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindFramebufferFunc() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindFramebufferFunc(
			this.Ref(),
		),
	)
}

// BindRenderbuffer calls the method "WebGL2RenderingContext.bindRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindRenderbufferFunc returns the method "WebGL2RenderingContext.bindRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindRenderbufferFunc() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindRenderbufferFunc(
			this.Ref(),
		),
	)
}

// BindTexture calls the method "WebGL2RenderingContext.bindTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BindTexture(target GLenum, texture WebGLTexture) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBindTexture(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		texture.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindTextureFunc returns the method "WebGL2RenderingContext.bindTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BindTextureFunc() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindTextureFunc(
			this.Ref(),
		),
	)
}

// BlendColor calls the method "WebGL2RenderingContext.blendColor".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlendColor(
		this.Ref(), js.Pointer(&_ok),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendColorFunc returns the method "WebGL2RenderingContext.blendColor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlendColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendColorFunc(
			this.Ref(),
		),
	)
}

// BlendEquation calls the method "WebGL2RenderingContext.blendEquation".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlendEquation(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlendEquation(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationFunc returns the method "WebGL2RenderingContext.blendEquation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlendEquationFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendEquationFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparate calls the method "WebGL2RenderingContext.blendEquationSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationSeparateFunc returns the method "WebGL2RenderingContext.blendEquationSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlendEquationSeparateFunc() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendEquationSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendFunc calls the method "WebGL2RenderingContext.blendFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlendFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(sfactor),
		uint32(dfactor),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFuncFunc returns the method "WebGL2RenderingContext.blendFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlendFuncFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendFuncFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparate calls the method "WebGL2RenderingContext.blendFuncSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFuncSeparateFunc returns the method "WebGL2RenderingContext.blendFuncSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) BlendFuncSeparateFunc() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// CheckFramebufferStatus calls the method "WebGL2RenderingContext.checkFramebufferStatus".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CheckFramebufferStatus(target GLenum) (GLenum, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
	)

	return GLenum(_ret), _ok
}

// CheckFramebufferStatusFunc returns the method "WebGL2RenderingContext.checkFramebufferStatus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CheckFramebufferStatusFunc() (fn js.Func[func(target GLenum) GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCheckFramebufferStatusFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "WebGL2RenderingContext.clear".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Clear(mask GLbitfield) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClear(
		this.Ref(), js.Pointer(&_ok),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the method "WebGL2RenderingContext.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearFunc() (fn js.Func[func(mask GLbitfield)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearFunc(
			this.Ref(),
		),
	)
}

// ClearColor calls the method "WebGL2RenderingContext.clearColor".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearColor(
		this.Ref(), js.Pointer(&_ok),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearColorFunc returns the method "WebGL2RenderingContext.clearColor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearColorFunc(
			this.Ref(),
		),
	)
}

// ClearDepth calls the method "WebGL2RenderingContext.clearDepth".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearDepth(depth GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearDepth(
		this.Ref(), js.Pointer(&_ok),
		float32(depth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearDepthFunc returns the method "WebGL2RenderingContext.clearDepth".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearDepthFunc() (fn js.Func[func(depth GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearDepthFunc(
			this.Ref(),
		),
	)
}

// ClearStencil calls the method "WebGL2RenderingContext.clearStencil".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ClearStencil(s GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextClearStencil(
		this.Ref(), js.Pointer(&_ok),
		int32(s),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearStencilFunc returns the method "WebGL2RenderingContext.clearStencil".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ClearStencilFunc() (fn js.Func[func(s GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearStencilFunc(
			this.Ref(),
		),
	)
}

// ColorMask calls the method "WebGL2RenderingContext.colorMask".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextColorMask(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ColorMaskFunc returns the method "WebGL2RenderingContext.colorMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ColorMaskFunc() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextColorMaskFunc(
			this.Ref(),
		),
	)
}

// CompileShader calls the method "WebGL2RenderingContext.compileShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CompileShader(shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCompileShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CompileShaderFunc returns the method "WebGL2RenderingContext.compileShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CompileShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompileShaderFunc(
			this.Ref(),
		),
	)
}

// CopyTexImage2D calls the method "WebGL2RenderingContext.copyTexImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTexImage2DFunc returns the method "WebGL2RenderingContext.copyTexImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CopyTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexSubImage2D calls the method "WebGL2RenderingContext.copyTexSubImage2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTexSubImage2DFunc returns the method "WebGL2RenderingContext.copyTexSubImage2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CopyTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "WebGL2RenderingContext.createBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateBuffer() (WebGLBuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLBuffer{}.FromRef(_ret), _ok
}

// CreateBufferFunc returns the method "WebGL2RenderingContext.createBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateBufferFunc() (fn js.Func[func() WebGLBuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateFramebuffer calls the method "WebGL2RenderingContext.createFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateFramebuffer() (WebGLFramebuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLFramebuffer{}.FromRef(_ret), _ok
}

// CreateFramebufferFunc returns the method "WebGL2RenderingContext.createFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateFramebufferFunc() (fn js.Func[func() WebGLFramebuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateFramebufferFunc(
			this.Ref(),
		),
	)
}

// CreateProgram calls the method "WebGL2RenderingContext.createProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateProgram() (WebGLProgram, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateProgram(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLProgram{}.FromRef(_ret), _ok
}

// CreateProgramFunc returns the method "WebGL2RenderingContext.createProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateProgramFunc() (fn js.Func[func() WebGLProgram]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateProgramFunc(
			this.Ref(),
		),
	)
}

// CreateRenderbuffer calls the method "WebGL2RenderingContext.createRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateRenderbuffer() (WebGLRenderbuffer, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLRenderbuffer{}.FromRef(_ret), _ok
}

// CreateRenderbufferFunc returns the method "WebGL2RenderingContext.createRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateRenderbufferFunc() (fn js.Func[func() WebGLRenderbuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateRenderbufferFunc(
			this.Ref(),
		),
	)
}

// CreateShader calls the method "WebGL2RenderingContext.createShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateShader(typ GLenum) (WebGLShader, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateShader(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return WebGLShader{}.FromRef(_ret), _ok
}

// CreateShaderFunc returns the method "WebGL2RenderingContext.createShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateShaderFunc() (fn js.Func[func(typ GLenum) WebGLShader]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateShaderFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "WebGL2RenderingContext.createTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CreateTexture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCreateTexture(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLTexture{}.FromRef(_ret), _ok
}

// CreateTextureFunc returns the method "WebGL2RenderingContext.createTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CreateTextureFunc() (fn js.Func[func() WebGLTexture]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CullFace calls the method "WebGL2RenderingContext.cullFace".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) CullFace(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextCullFace(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CullFaceFunc returns the method "WebGL2RenderingContext.cullFace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) CullFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCullFaceFunc(
			this.Ref(),
		),
	)
}

// DeleteBuffer calls the method "WebGL2RenderingContext.deleteBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteBuffer(buffer WebGLBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteBufferFunc returns the method "WebGL2RenderingContext.deleteBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteBufferFunc() (fn js.Func[func(buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteBufferFunc(
			this.Ref(),
		),
	)
}

// DeleteFramebuffer calls the method "WebGL2RenderingContext.deleteFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		framebuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteFramebufferFunc returns the method "WebGL2RenderingContext.deleteFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteFramebufferFunc(
			this.Ref(),
		),
	)
}

// DeleteProgram calls the method "WebGL2RenderingContext.deleteProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteProgramFunc returns the method "WebGL2RenderingContext.deleteProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteProgramFunc(
			this.Ref(),
		),
	)
}

// DeleteRenderbuffer calls the method "WebGL2RenderingContext.deleteRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRenderbufferFunc returns the method "WebGL2RenderingContext.deleteRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteRenderbufferFunc(
			this.Ref(),
		),
	)
}

// DeleteShader calls the method "WebGL2RenderingContext.deleteShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteShader(shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteShaderFunc returns the method "WebGL2RenderingContext.deleteShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteShaderFunc(
			this.Ref(),
		),
	)
}

// DeleteTexture calls the method "WebGL2RenderingContext.deleteTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DeleteTexture(texture WebGLTexture) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&_ok),
		texture.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteTextureFunc returns the method "WebGL2RenderingContext.deleteTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DeleteTextureFunc() (fn js.Func[func(texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteTextureFunc(
			this.Ref(),
		),
	)
}

// DepthFunc calls the method "WebGL2RenderingContext.depthFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DepthFunc(fn GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDepthFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(fn),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthFuncFunc returns the method "WebGL2RenderingContext.depthFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DepthFuncFunc() (fn js.Func[func(fn GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthFuncFunc(
			this.Ref(),
		),
	)
}

// DepthMask calls the method "WebGL2RenderingContext.depthMask".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DepthMask(flag GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDepthMask(
		this.Ref(), js.Pointer(&_ok),
		js.Bool(bool(flag)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthMaskFunc returns the method "WebGL2RenderingContext.depthMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DepthMaskFunc() (fn js.Func[func(flag GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthMaskFunc(
			this.Ref(),
		),
	)
}

// DepthRange calls the method "WebGL2RenderingContext.depthRange".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDepthRange(
		this.Ref(), js.Pointer(&_ok),
		float32(zNear),
		float32(zFar),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DepthRangeFunc returns the method "WebGL2RenderingContext.depthRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DepthRangeFunc() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthRangeFunc(
			this.Ref(),
		),
	)
}

// DetachShader calls the method "WebGL2RenderingContext.detachShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDetachShader(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		shader.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DetachShaderFunc returns the method "WebGL2RenderingContext.detachShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DetachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDetachShaderFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "WebGL2RenderingContext.disable".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Disable(cap GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDisable(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisableFunc returns the method "WebGL2RenderingContext.disable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DisableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDisableFunc(
			this.Ref(),
		),
	)
}

// DisableVertexAttribArray calls the method "WebGL2RenderingContext.disableVertexAttribArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DisableVertexAttribArray(index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisableVertexAttribArrayFunc returns the method "WebGL2RenderingContext.disableVertexAttribArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DisableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDisableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// DrawArrays calls the method "WebGL2RenderingContext.drawArrays".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawArrays(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(first),
		int32(count),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawArraysFunc returns the method "WebGL2RenderingContext.drawArrays".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawArraysFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawArraysFunc(
			this.Ref(),
		),
	)
}

// DrawElements calls the method "WebGL2RenderingContext.drawElements".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextDrawElements(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawElementsFunc returns the method "WebGL2RenderingContext.drawElements".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) DrawElementsFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawElementsFunc(
			this.Ref(),
		),
	)
}

// Enable calls the method "WebGL2RenderingContext.enable".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Enable(cap GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextEnable(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnableFunc returns the method "WebGL2RenderingContext.enable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) EnableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEnableFunc(
			this.Ref(),
		),
	)
}

// EnableVertexAttribArray calls the method "WebGL2RenderingContext.enableVertexAttribArray".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) EnableVertexAttribArray(index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnableVertexAttribArrayFunc returns the method "WebGL2RenderingContext.enableVertexAttribArray".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) EnableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEnableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "WebGL2RenderingContext.finish".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Finish() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFinish(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FinishFunc returns the method "WebGL2RenderingContext.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFinishFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "WebGL2RenderingContext.flush".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Flush() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFlush(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FlushFunc returns the method "WebGL2RenderingContext.flush".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFlushFunc(
			this.Ref(),
		),
	)
}

// FramebufferRenderbuffer calls the method "WebGL2RenderingContext.framebufferRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferRenderbufferFunc returns the method "WebGL2RenderingContext.framebufferRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FramebufferRenderbufferFunc() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferRenderbufferFunc(
			this.Ref(),
		),
	)
}

// FramebufferTexture2D calls the method "WebGL2RenderingContext.framebufferTexture2D".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferTexture2DFunc returns the method "WebGL2RenderingContext.framebufferTexture2D".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FramebufferTexture2DFunc() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferTexture2DFunc(
			this.Ref(),
		),
	)
}

// FrontFace calls the method "WebGL2RenderingContext.frontFace".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) FrontFace(mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextFrontFace(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FrontFaceFunc returns the method "WebGL2RenderingContext.frontFace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) FrontFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFrontFaceFunc(
			this.Ref(),
		),
	)
}

// GenerateMipmap calls the method "WebGL2RenderingContext.generateMipmap".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GenerateMipmap(target GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GenerateMipmapFunc returns the method "WebGL2RenderingContext.generateMipmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GenerateMipmapFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGenerateMipmapFunc(
			this.Ref(),
		),
	)
}

// GetActiveAttrib calls the method "WebGL2RenderingContext.getActiveAttrib".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (WebGLActiveInfo, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
	)

	return WebGLActiveInfo{}.FromRef(_ret), _ok
}

// GetActiveAttribFunc returns the method "WebGL2RenderingContext.getActiveAttrib".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetActiveAttribFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveAttribFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniform calls the method "WebGL2RenderingContext.getActiveUniform".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (WebGLActiveInfo, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(index),
	)

	return WebGLActiveInfo{}.FromRef(_ret), _ok
}

// GetActiveUniformFunc returns the method "WebGL2RenderingContext.getActiveUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetActiveUniformFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformFunc(
			this.Ref(),
		),
	)
}

// GetAttachedShaders calls the method "WebGL2RenderingContext.getAttachedShaders".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetAttachedShaders(program WebGLProgram) (js.Array[WebGLShader], bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return js.Array[WebGLShader]{}.FromRef(_ret), _ok
}

// GetAttachedShadersFunc returns the method "WebGL2RenderingContext.getAttachedShaders".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetAttachedShadersFunc() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetAttachedShadersFunc(
			this.Ref(),
		),
	)
}

// GetAttribLocation calls the method "WebGL2RenderingContext.getAttribLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (GLint, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		name.Ref(),
	)

	return GLint(_ret), _ok
}

// GetAttribLocationFunc returns the method "WebGL2RenderingContext.getAttribLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetAttribLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetAttribLocationFunc(
			this.Ref(),
		),
	)
}

// GetBufferParameter calls the method "WebGL2RenderingContext.getBufferParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetBufferParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetBufferParameterFunc returns the method "WebGL2RenderingContext.getBufferParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetBufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetParameter calls the method "WebGL2RenderingContext.getParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetParameter(pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetParameterFunc returns the method "WebGL2RenderingContext.getParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetParameterFunc() (fn js.Func[func(pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetParameterFunc(
			this.Ref(),
		),
	)
}

// GetError calls the method "WebGL2RenderingContext.getError".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetError() (GLenum, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetError(
		this.Ref(), js.Pointer(&_ok),
	)

	return GLenum(_ret), _ok
}

// GetErrorFunc returns the method "WebGL2RenderingContext.getError".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetErrorFunc() (fn js.Func[func() GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetErrorFunc(
			this.Ref(),
		),
	)
}

// GetFramebufferAttachmentParameter calls the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetFramebufferAttachmentParameterFunc returns the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetFramebufferAttachmentParameterFunc() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetFramebufferAttachmentParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramParameter calls the method "WebGL2RenderingContext.getProgramParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetProgramParameterFunc returns the method "WebGL2RenderingContext.getProgramParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetProgramParameterFunc() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetProgramParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramInfoLog calls the method "WebGL2RenderingContext.getProgramInfoLog".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetProgramInfoLog(program WebGLProgram) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetProgramInfoLogFunc returns the method "WebGL2RenderingContext.getProgramInfoLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetProgramInfoLogFunc() (fn js.Func[func(program WebGLProgram) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetProgramInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetRenderbufferParameter calls the method "WebGL2RenderingContext.getRenderbufferParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetRenderbufferParameterFunc returns the method "WebGL2RenderingContext.getRenderbufferParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetRenderbufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetRenderbufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderParameter calls the method "WebGL2RenderingContext.getShaderParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetShaderParameterFunc returns the method "WebGL2RenderingContext.getShaderParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetShaderParameterFunc() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderPrecisionFormat calls the method "WebGL2RenderingContext.getShaderPrecisionFormat".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (WebGLShaderPrecisionFormat, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&_ok),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return WebGLShaderPrecisionFormat{}.FromRef(_ret), _ok
}

// GetShaderPrecisionFormatFunc returns the method "WebGL2RenderingContext.getShaderPrecisionFormat".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetShaderPrecisionFormatFunc() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderPrecisionFormatFunc(
			this.Ref(),
		),
	)
}

// GetShaderInfoLog calls the method "WebGL2RenderingContext.getShaderInfoLog".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetShaderInfoLog(shader WebGLShader) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetShaderInfoLogFunc returns the method "WebGL2RenderingContext.getShaderInfoLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetShaderInfoLogFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetShaderSource calls the method "WebGL2RenderingContext.getShaderSource".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetShaderSource(shader WebGLShader) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetShaderSourceFunc returns the method "WebGL2RenderingContext.getShaderSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetShaderSourceFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderSourceFunc(
			this.Ref(),
		),
	)
}

// GetTexParameter calls the method "WebGL2RenderingContext.getTexParameter".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetTexParameter(target GLenum, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetTexParameterFunc returns the method "WebGL2RenderingContext.getTexParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetTexParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetTexParameterFunc(
			this.Ref(),
		),
	)
}

// GetUniform calls the method "WebGL2RenderingContext.getUniform".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetUniform(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		location.Ref(),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetUniformFunc returns the method "WebGL2RenderingContext.getUniform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetUniformFunc() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformFunc(
			this.Ref(),
		),
	)
}

// GetUniformLocation calls the method "WebGL2RenderingContext.getUniformLocation".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (WebGLUniformLocation, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
		name.Ref(),
	)

	return WebGLUniformLocation{}.FromRef(_ret), _ok
}

// GetUniformLocationFunc returns the method "WebGL2RenderingContext.getUniformLocation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetUniformLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformLocationFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttrib calls the method "WebGL2RenderingContext.getVertexAttrib".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(pname),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetVertexAttribFunc returns the method "WebGL2RenderingContext.getVertexAttrib".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetVertexAttribFunc() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetVertexAttribFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttribOffset calls the method "WebGL2RenderingContext.getVertexAttribOffset".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (GLintptr, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(pname),
	)

	return GLintptr(_ret), _ok
}

// GetVertexAttribOffsetFunc returns the method "WebGL2RenderingContext.getVertexAttribOffset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) GetVertexAttribOffsetFunc() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetVertexAttribOffsetFunc(
			this.Ref(),
		),
	)
}

// Hint calls the method "WebGL2RenderingContext.hint".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Hint(target GLenum, mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextHint(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// HintFunc returns the method "WebGL2RenderingContext.hint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) HintFunc() (fn js.Func[func(target GLenum, mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextHintFunc(
			this.Ref(),
		),
	)
}

// IsBuffer calls the method "WebGL2RenderingContext.isBuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsBuffer(buffer WebGLBuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsBufferFunc returns the method "WebGL2RenderingContext.isBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsBufferFunc() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsBufferFunc(
			this.Ref(),
		),
	)
}

// IsEnabled calls the method "WebGL2RenderingContext.isEnabled".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsEnabled(cap GLenum) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsEnabled(
		this.Ref(), js.Pointer(&_ok),
		uint32(cap),
	)

	return _ret == js.True, _ok
}

// IsEnabledFunc returns the method "WebGL2RenderingContext.isEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsEnabledFunc() (fn js.Func[func(cap GLenum) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsEnabledFunc(
			this.Ref(),
		),
	)
}

// IsFramebuffer calls the method "WebGL2RenderingContext.isFramebuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&_ok),
		framebuffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsFramebufferFunc returns the method "WebGL2RenderingContext.isFramebuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsFramebufferFunc(
			this.Ref(),
		),
	)
}

// IsProgram calls the method "WebGL2RenderingContext.isProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsProgram(program WebGLProgram) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	return _ret == js.True, _ok
}

// IsProgramFunc returns the method "WebGL2RenderingContext.isProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsProgramFunc() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsProgramFunc(
			this.Ref(),
		),
	)
}

// IsRenderbuffer calls the method "WebGL2RenderingContext.isRenderbuffer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&_ok),
		renderbuffer.Ref(),
	)

	return _ret == js.True, _ok
}

// IsRenderbufferFunc returns the method "WebGL2RenderingContext.isRenderbuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsRenderbufferFunc(
			this.Ref(),
		),
	)
}

// IsShader calls the method "WebGL2RenderingContext.isShader".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsShader(shader WebGLShader) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsShader(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
	)

	return _ret == js.True, _ok
}

// IsShaderFunc returns the method "WebGL2RenderingContext.isShader".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsShaderFunc() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsShaderFunc(
			this.Ref(),
		),
	)
}

// IsTexture calls the method "WebGL2RenderingContext.isTexture".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) IsTexture(texture WebGLTexture) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextIsTexture(
		this.Ref(), js.Pointer(&_ok),
		texture.Ref(),
	)

	return _ret == js.True, _ok
}

// IsTextureFunc returns the method "WebGL2RenderingContext.isTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) IsTextureFunc() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsTextureFunc(
			this.Ref(),
		),
	)
}

// LineWidth calls the method "WebGL2RenderingContext.lineWidth".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) LineWidth(width GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextLineWidth(
		this.Ref(), js.Pointer(&_ok),
		float32(width),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineWidthFunc returns the method "WebGL2RenderingContext.lineWidth".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) LineWidthFunc() (fn js.Func[func(width GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextLineWidthFunc(
			this.Ref(),
		),
	)
}

// LinkProgram calls the method "WebGL2RenderingContext.linkProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) LinkProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextLinkProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LinkProgramFunc returns the method "WebGL2RenderingContext.linkProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) LinkProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextLinkProgramFunc(
			this.Ref(),
		),
	)
}

// PixelStorei calls the method "WebGL2RenderingContext.pixelStorei".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) PixelStorei(pname GLenum, param GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextPixelStorei(
		this.Ref(), js.Pointer(&_ok),
		uint32(pname),
		int32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PixelStoreiFunc returns the method "WebGL2RenderingContext.pixelStorei".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) PixelStoreiFunc() (fn js.Func[func(pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPixelStoreiFunc(
			this.Ref(),
		),
	)
}

// PolygonOffset calls the method "WebGL2RenderingContext.polygonOffset".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&_ok),
		float32(factor),
		float32(units),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PolygonOffsetFunc returns the method "WebGL2RenderingContext.polygonOffset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) PolygonOffsetFunc() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPolygonOffsetFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorage calls the method "WebGL2RenderingContext.renderbufferStorage".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RenderbufferStorageFunc returns the method "WebGL2RenderingContext.renderbufferStorage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) RenderbufferStorageFunc() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextRenderbufferStorageFunc(
			this.Ref(),
		),
	)
}

// SampleCoverage calls the method "WebGL2RenderingContext.sampleCoverage".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&_ok),
		float32(value),
		js.Bool(bool(invert)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SampleCoverageFunc returns the method "WebGL2RenderingContext.sampleCoverage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) SampleCoverageFunc() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSampleCoverageFunc(
			this.Ref(),
		),
	)
}

// Scissor calls the method "WebGL2RenderingContext.scissor".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextScissor(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScissorFunc returns the method "WebGL2RenderingContext.scissor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ScissorFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextScissorFunc(
			this.Ref(),
		),
	)
}

// ShaderSource calls the method "WebGL2RenderingContext.shaderSource".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ShaderSource(shader WebGLShader, source js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextShaderSource(
		this.Ref(), js.Pointer(&_ok),
		shader.Ref(),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShaderSourceFunc returns the method "WebGL2RenderingContext.shaderSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ShaderSourceFunc() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextShaderSourceFunc(
			this.Ref(),
		),
	)
}

// StencilFunc calls the method "WebGL2RenderingContext.stencilFunc".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilFunc(
		this.Ref(), js.Pointer(&_ok),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilFuncFunc returns the method "WebGL2RenderingContext.stencilFunc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilFuncFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilFuncFunc(
			this.Ref(),
		),
	)
}

// StencilFuncSeparate calls the method "WebGL2RenderingContext.stencilFuncSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilFuncSeparateFunc returns the method "WebGL2RenderingContext.stencilFuncSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilFuncSeparateFunc() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilMask calls the method "WebGL2RenderingContext.stencilMask".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilMask(mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilMask(
		this.Ref(), js.Pointer(&_ok),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilMaskFunc returns the method "WebGL2RenderingContext.stencilMask".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilMaskFunc() (fn js.Func[func(mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilMaskFunc(
			this.Ref(),
		),
	)
}

// StencilMaskSeparate calls the method "WebGL2RenderingContext.stencilMaskSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(mask),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilMaskSeparateFunc returns the method "WebGL2RenderingContext.stencilMaskSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilMaskSeparateFunc() (fn js.Func[func(face GLenum, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilMaskSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilOp calls the method "WebGL2RenderingContext.stencilOp".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilOp(
		this.Ref(), js.Pointer(&_ok),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilOpFunc returns the method "WebGL2RenderingContext.stencilOp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilOpFunc() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilOpFunc(
			this.Ref(),
		),
	)
}

// StencilOpSeparate calls the method "WebGL2RenderingContext.stencilOpSeparate".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&_ok),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StencilOpSeparateFunc returns the method "WebGL2RenderingContext.stencilOpSeparate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) StencilOpSeparateFunc() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilOpSeparateFunc(
			this.Ref(),
		),
	)
}

// TexParameterf calls the method "WebGL2RenderingContext.texParameterf".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexParameterf(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexParameterfFunc returns the method "WebGL2RenderingContext.texParameterf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexParameterfFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexParameterfFunc(
			this.Ref(),
		),
	)
}

// TexParameteri calls the method "WebGL2RenderingContext.texParameteri".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextTexParameteri(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TexParameteriFunc returns the method "WebGL2RenderingContext.texParameteri".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) TexParameteriFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexParameteriFunc(
			this.Ref(),
		),
	)
}

// Uniform1f calls the method "WebGL2RenderingContext.uniform1f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1fFunc returns the method "WebGL2RenderingContext.uniform1f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fFunc(
			this.Ref(),
		),
	)
}

// Uniform2f calls the method "WebGL2RenderingContext.uniform2f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2fFunc returns the method "WebGL2RenderingContext.uniform2f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fFunc(
			this.Ref(),
		),
	)
}

// Uniform3f calls the method "WebGL2RenderingContext.uniform3f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3fFunc returns the method "WebGL2RenderingContext.uniform3f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fFunc(
			this.Ref(),
		),
	)
}

// Uniform4f calls the method "WebGL2RenderingContext.uniform4f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4f(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4fFunc returns the method "WebGL2RenderingContext.uniform4f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fFunc(
			this.Ref(),
		),
	)
}

// Uniform1i calls the method "WebGL2RenderingContext.uniform1i".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform1i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform1iFunc returns the method "WebGL2RenderingContext.uniform1i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform1iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iFunc(
			this.Ref(),
		),
	)
}

// Uniform2i calls the method "WebGL2RenderingContext.uniform2i".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform2i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform2iFunc returns the method "WebGL2RenderingContext.uniform2i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform2iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iFunc(
			this.Ref(),
		),
	)
}

// Uniform3i calls the method "WebGL2RenderingContext.uniform3i".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform3i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform3iFunc returns the method "WebGL2RenderingContext.uniform3i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform3iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iFunc(
			this.Ref(),
		),
	)
}

// Uniform4i calls the method "WebGL2RenderingContext.uniform4i".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUniform4i(
		this.Ref(), js.Pointer(&_ok),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Uniform4iFunc returns the method "WebGL2RenderingContext.uniform4i".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) Uniform4iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iFunc(
			this.Ref(),
		),
	)
}

// UseProgram calls the method "WebGL2RenderingContext.useProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) UseProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextUseProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UseProgramFunc returns the method "WebGL2RenderingContext.useProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) UseProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUseProgramFunc(
			this.Ref(),
		),
	)
}

// ValidateProgram calls the method "WebGL2RenderingContext.validateProgram".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) ValidateProgram(program WebGLProgram) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextValidateProgram(
		this.Ref(), js.Pointer(&_ok),
		program.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ValidateProgramFunc returns the method "WebGL2RenderingContext.validateProgram".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ValidateProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextValidateProgramFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1f calls the method "WebGL2RenderingContext.vertexAttrib1f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib1fFunc returns the method "WebGL2RenderingContext.vertexAttrib1f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib1fFunc() (fn js.Func[func(index GLuint, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib1fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2f calls the method "WebGL2RenderingContext.vertexAttrib2f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib2fFunc returns the method "WebGL2RenderingContext.vertexAttrib2f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib2fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib2fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3f calls the method "WebGL2RenderingContext.vertexAttrib3f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib3fFunc returns the method "WebGL2RenderingContext.vertexAttrib3f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib3fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib3fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4f calls the method "WebGL2RenderingContext.vertexAttrib4f".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib4fFunc returns the method "WebGL2RenderingContext.vertexAttrib4f".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib4fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib4fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1fv calls the method "WebGL2RenderingContext.vertexAttrib1fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib1fvFunc returns the method "WebGL2RenderingContext.vertexAttrib1fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib1fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib1fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2fv calls the method "WebGL2RenderingContext.vertexAttrib2fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib2fvFunc returns the method "WebGL2RenderingContext.vertexAttrib2fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib2fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib2fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3fv calls the method "WebGL2RenderingContext.vertexAttrib3fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib3fvFunc returns the method "WebGL2RenderingContext.vertexAttrib3fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib3fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib3fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4fv calls the method "WebGL2RenderingContext.vertexAttrib4fv".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		values.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttrib4fvFunc returns the method "WebGL2RenderingContext.vertexAttrib4fv".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttrib4fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib4fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttribPointer calls the method "WebGL2RenderingContext.vertexAttribPointer".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribPointerFunc returns the method "WebGL2RenderingContext.vertexAttribPointer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) VertexAttribPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribPointerFunc(
			this.Ref(),
		),
	)
}

// Viewport calls the method "WebGL2RenderingContext.viewport".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextViewport(
		this.Ref(), js.Pointer(&_ok),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ViewportFunc returns the method "WebGL2RenderingContext.viewport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) ViewportFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextViewportFunc(
			this.Ref(),
		),
	)
}

// MakeXRCompatible calls the method "WebGL2RenderingContext.makeXRCompatible".
//
// The returned bool will be false if there is no such method.
func (this WebGL2RenderingContext) MakeXRCompatible() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallWebGL2RenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MakeXRCompatibleFunc returns the method "WebGL2RenderingContext.makeXRCompatible".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WebGL2RenderingContext) MakeXRCompatibleFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextMakeXRCompatibleFunc(
			this.Ref(),
		),
	)
}

type GPUMapModeFlags uint32

type GPUSize64 uint64

type GPUSize64Out uint64

type GPUFlagsConstant uint32

type GPUBufferMapState uint32

const (
	_ GPUBufferMapState = iota

	GPUBufferMapState_UNMAPPED
	GPUBufferMapState_PENDING
	GPUBufferMapState_MAPPED
)

func (GPUBufferMapState) FromRef(str js.Ref) GPUBufferMapState {
	return GPUBufferMapState(bindings.ConstOfGPUBufferMapState(str))
}

func (x GPUBufferMapState) String() (string, bool) {
	switch x {
	case GPUBufferMapState_UNMAPPED:
		return "unmapped", true
	case GPUBufferMapState_PENDING:
		return "pending", true
	case GPUBufferMapState_MAPPED:
		return "mapped", true
	default:
		return "", false
	}
}

type GPUBuffer struct {
	ref js.Ref
}

func (this GPUBuffer) Once() GPUBuffer {
	this.Ref().Once()
	return this
}

func (this GPUBuffer) Ref() js.Ref {
	return this.ref
}

func (this GPUBuffer) FromRef(ref js.Ref) GPUBuffer {
	this.ref = ref
	return this
}

func (this GPUBuffer) Free() {
	this.Ref().Free()
}

// Size returns the value of property "GPUBuffer.size".
//
// The returned bool will be false if there is no such property.
func (this GPUBuffer) Size() (GPUSize64Out, bool) {
	var _ok bool
	_ret := bindings.GetGPUBufferSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSize64Out(_ret), _ok
}

// Usage returns the value of property "GPUBuffer.usage".
//
// The returned bool will be false if there is no such property.
func (this GPUBuffer) Usage() (GPUFlagsConstant, bool) {
	var _ok bool
	_ret := bindings.GetGPUBufferUsage(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUFlagsConstant(_ret), _ok
}

// MapState returns the value of property "GPUBuffer.mapState".
//
// The returned bool will be false if there is no such property.
func (this GPUBuffer) MapState() (GPUBufferMapState, bool) {
	var _ok bool
	_ret := bindings.GetGPUBufferMapState(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUBufferMapState(_ret), _ok
}

// Label returns the value of property "GPUBuffer.label".
//
// The returned bool will be false if there is no such property.
func (this GPUBuffer) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUBufferLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUBuffer.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBuffer) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBufferLabel(
		this.Ref(),
		val.Ref(),
	)
}

// MapAsync calls the method "GPUBuffer.mapAsync".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) MapAsync(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferMapAsync(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		float64(offset),
		float64(size),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MapAsyncFunc returns the method "GPUBuffer.mapAsync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) MapAsyncFunc() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsyncFunc(
			this.Ref(),
		),
	)
}

// MapAsync1 calls the method "GPUBuffer.mapAsync".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) MapAsync1(mode GPUMapModeFlags, offset GPUSize64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferMapAsync1(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		float64(offset),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MapAsync1Func returns the method "GPUBuffer.mapAsync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) MapAsync1Func() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsync1Func(
			this.Ref(),
		),
	)
}

// MapAsync2 calls the method "GPUBuffer.mapAsync".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) MapAsync2(mode GPUMapModeFlags) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferMapAsync2(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MapAsync2Func returns the method "GPUBuffer.mapAsync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) MapAsync2Func() (fn js.Func[func(mode GPUMapModeFlags) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsync2Func(
			this.Ref(),
		),
	)
}

// GetMappedRange calls the method "GPUBuffer.getMappedRange".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) GetMappedRange(offset GPUSize64, size GPUSize64) (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferGetMappedRange(
		this.Ref(), js.Pointer(&_ok),
		float64(offset),
		float64(size),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// GetMappedRangeFunc returns the method "GPUBuffer.getMappedRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) GetMappedRangeFunc() (fn js.Func[func(offset GPUSize64, size GPUSize64) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRangeFunc(
			this.Ref(),
		),
	)
}

// GetMappedRange1 calls the method "GPUBuffer.getMappedRange".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) GetMappedRange1(offset GPUSize64) (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferGetMappedRange1(
		this.Ref(), js.Pointer(&_ok),
		float64(offset),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// GetMappedRange1Func returns the method "GPUBuffer.getMappedRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) GetMappedRange1Func() (fn js.Func[func(offset GPUSize64) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRange1Func(
			this.Ref(),
		),
	)
}

// GetMappedRange2 calls the method "GPUBuffer.getMappedRange".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) GetMappedRange2() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferGetMappedRange2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// GetMappedRange2Func returns the method "GPUBuffer.getMappedRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) GetMappedRange2Func() (fn js.Func[func() js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRange2Func(
			this.Ref(),
		),
	)
}

// Unmap calls the method "GPUBuffer.unmap".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) Unmap() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferUnmap(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnmapFunc returns the method "GPUBuffer.unmap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) UnmapFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUBufferUnmapFunc(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUBuffer.destroy".
//
// The returned bool will be false if there is no such method.
func (this GPUBuffer) Destroy() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUBufferDestroy(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DestroyFunc returns the method "GPUBuffer.destroy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUBuffer) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUBufferDestroyFunc(
			this.Ref(),
		),
	)
}

type GPUBufferUsageFlags uint32

type GPUBufferDescriptor struct {
	// Size is "GPUBufferDescriptor.size"
	//
	// Required
	Size GPUSize64
	// Usage is "GPUBufferDescriptor.usage"
	//
	// Required
	Usage GPUBufferUsageFlags
	// MappedAtCreation is "GPUBufferDescriptor.mappedAtCreation"
	//
	// Optional, defaults to false.
	MappedAtCreation bool
	// Label is "GPUBufferDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_MappedAtCreation bool // for MappedAtCreation.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBufferDescriptor with all fields set.
func (p GPUBufferDescriptor) FromRef(ref js.Ref) GPUBufferDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBufferDescriptor GPUBufferDescriptor [// GPUBufferDescriptor] [0x14000313540 0x140003135e0 0x14000313680 0x140003137c0 0x14000313720] 0x1400037c618 {0 0}} in the application heap.
func (p GPUBufferDescriptor) New() js.Ref {
	return bindings.GPUBufferDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBufferDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBufferDescriptor) Update(ref js.Ref) {
	bindings.GPUBufferDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUTextureView struct {
	ref js.Ref
}

func (this GPUTextureView) Once() GPUTextureView {
	this.Ref().Once()
	return this
}

func (this GPUTextureView) Ref() js.Ref {
	return this.ref
}

func (this GPUTextureView) FromRef(ref js.Ref) GPUTextureView {
	this.ref = ref
	return this
}

func (this GPUTextureView) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUTextureView.label".
//
// The returned bool will be false if there is no such property.
func (this GPUTextureView) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureViewLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUTextureView.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUTextureView) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUTextureViewLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUTextureFormat uint32

const (
	_ GPUTextureFormat = iota

	GPUTextureFormat_R_8UNORM
	GPUTextureFormat_R_8SNORM
	GPUTextureFormat_R_8UINT
	GPUTextureFormat_R_8SINT
	GPUTextureFormat_R_16UINT
	GPUTextureFormat_R_16SINT
	GPUTextureFormat_R_16FLOAT
	GPUTextureFormat_RG_8UNORM
	GPUTextureFormat_RG_8SNORM
	GPUTextureFormat_RG_8UINT
	GPUTextureFormat_RG_8SINT
	GPUTextureFormat_R_32UINT
	GPUTextureFormat_R_32SINT
	GPUTextureFormat_R_32FLOAT
	GPUTextureFormat_RG_16UINT
	GPUTextureFormat_RG_16SINT
	GPUTextureFormat_RG_16FLOAT
	GPUTextureFormat_RGBA_8UNORM
	GPUTextureFormat_RGBA_8UNORM_SRGB
	GPUTextureFormat_RGBA_8SNORM
	GPUTextureFormat_RGBA_8UINT
	GPUTextureFormat_RGBA_8SINT
	GPUTextureFormat_BGRA_8UNORM
	GPUTextureFormat_BGRA_8UNORM_SRGB
	GPUTextureFormat_RGB_9E5UFLOAT
	GPUTextureFormat_RGB_10A2UINT
	GPUTextureFormat_RGB_10A2UNORM
	GPUTextureFormat_RG_11B10UFLOAT
	GPUTextureFormat_RG_32UINT
	GPUTextureFormat_RG_32SINT
	GPUTextureFormat_RG_32FLOAT
	GPUTextureFormat_RGBA_16UINT
	GPUTextureFormat_RGBA_16SINT
	GPUTextureFormat_RGBA_16FLOAT
	GPUTextureFormat_RGBA_32UINT
	GPUTextureFormat_RGBA_32SINT
	GPUTextureFormat_RGBA_32FLOAT
	GPUTextureFormat_STENCIL8
	GPUTextureFormat_DEPTH_16UNORM
	GPUTextureFormat_DEPTH_24PLUS
	GPUTextureFormat_DEPTH_24PLUS_STENCIL8
	GPUTextureFormat_DEPTH_32FLOAT
	GPUTextureFormat_DEPTH_32FLOAT_STENCIL8
	GPUTextureFormat_BC1_RGBA_UNORM
	GPUTextureFormat_BC1_RGBA_UNORM_SRGB
	GPUTextureFormat_BC2_RGBA_UNORM
	GPUTextureFormat_BC2_RGBA_UNORM_SRGB
	GPUTextureFormat_BC3_RGBA_UNORM
	GPUTextureFormat_BC3_RGBA_UNORM_SRGB
	GPUTextureFormat_BC4_R_UNORM
	GPUTextureFormat_BC4_R_SNORM
	GPUTextureFormat_BC5_RG_UNORM
	GPUTextureFormat_BC5_RG_SNORM
	GPUTextureFormat_BC_6H_RGB_UFLOAT
	GPUTextureFormat_BC_6H_RGB_FLOAT
	GPUTextureFormat_BC7_RGBA_UNORM
	GPUTextureFormat_BC7_RGBA_UNORM_SRGB
	GPUTextureFormat_ETC2_RGB_8UNORM
	GPUTextureFormat_ETC2_RGB_8UNORM_SRGB
	GPUTextureFormat_ETC2_RGB_8A1UNORM
	GPUTextureFormat_ETC2_RGB_8A1UNORM_SRGB
	GPUTextureFormat_ETC2_RGBA_8UNORM
	GPUTextureFormat_ETC2_RGBA_8UNORM_SRGB
	GPUTextureFormat_EAC_R_11UNORM
	GPUTextureFormat_EAC_R_11SNORM
	GPUTextureFormat_EAC_RG_11UNORM
	GPUTextureFormat_EAC_RG_11SNORM
	GPUTextureFormat_ASTC_4X4_UNORM
	GPUTextureFormat_ASTC_4X4_UNORM_SRGB
	GPUTextureFormat_ASTC_5X4_UNORM
	GPUTextureFormat_ASTC_5X4_UNORM_SRGB
	GPUTextureFormat_ASTC_5X5_UNORM
	GPUTextureFormat_ASTC_5X5_UNORM_SRGB
	GPUTextureFormat_ASTC_6X5_UNORM
	GPUTextureFormat_ASTC_6X5_UNORM_SRGB
	GPUTextureFormat_ASTC_6X6_UNORM
	GPUTextureFormat_ASTC_6X6_UNORM_SRGB
	GPUTextureFormat_ASTC_8X5_UNORM
	GPUTextureFormat_ASTC_8X5_UNORM_SRGB
	GPUTextureFormat_ASTC_8X6_UNORM
	GPUTextureFormat_ASTC_8X6_UNORM_SRGB
	GPUTextureFormat_ASTC_8X8_UNORM
	GPUTextureFormat_ASTC_8X8_UNORM_SRGB
	GPUTextureFormat_ASTC_10X5_UNORM
	GPUTextureFormat_ASTC_10X5_UNORM_SRGB
	GPUTextureFormat_ASTC_10X6_UNORM
	GPUTextureFormat_ASTC_10X6_UNORM_SRGB
	GPUTextureFormat_ASTC_10X8_UNORM
	GPUTextureFormat_ASTC_10X8_UNORM_SRGB
	GPUTextureFormat_ASTC_10X10_UNORM
	GPUTextureFormat_ASTC_10X10_UNORM_SRGB
	GPUTextureFormat_ASTC_12X10_UNORM
	GPUTextureFormat_ASTC_12X10_UNORM_SRGB
	GPUTextureFormat_ASTC_12X12_UNORM
	GPUTextureFormat_ASTC_12X12_UNORM_SRGB
)

func (GPUTextureFormat) FromRef(str js.Ref) GPUTextureFormat {
	return GPUTextureFormat(bindings.ConstOfGPUTextureFormat(str))
}

func (x GPUTextureFormat) String() (string, bool) {
	switch x {
	case GPUTextureFormat_R_8UNORM:
		return "r8unorm", true
	case GPUTextureFormat_R_8SNORM:
		return "r8snorm", true
	case GPUTextureFormat_R_8UINT:
		return "r8uint", true
	case GPUTextureFormat_R_8SINT:
		return "r8sint", true
	case GPUTextureFormat_R_16UINT:
		return "r16uint", true
	case GPUTextureFormat_R_16SINT:
		return "r16sint", true
	case GPUTextureFormat_R_16FLOAT:
		return "r16float", true
	case GPUTextureFormat_RG_8UNORM:
		return "rg8unorm", true
	case GPUTextureFormat_RG_8SNORM:
		return "rg8snorm", true
	case GPUTextureFormat_RG_8UINT:
		return "rg8uint", true
	case GPUTextureFormat_RG_8SINT:
		return "rg8sint", true
	case GPUTextureFormat_R_32UINT:
		return "r32uint", true
	case GPUTextureFormat_R_32SINT:
		return "r32sint", true
	case GPUTextureFormat_R_32FLOAT:
		return "r32float", true
	case GPUTextureFormat_RG_16UINT:
		return "rg16uint", true
	case GPUTextureFormat_RG_16SINT:
		return "rg16sint", true
	case GPUTextureFormat_RG_16FLOAT:
		return "rg16float", true
	case GPUTextureFormat_RGBA_8UNORM:
		return "rgba8unorm", true
	case GPUTextureFormat_RGBA_8UNORM_SRGB:
		return "rgba8unorm-srgb", true
	case GPUTextureFormat_RGBA_8SNORM:
		return "rgba8snorm", true
	case GPUTextureFormat_RGBA_8UINT:
		return "rgba8uint", true
	case GPUTextureFormat_RGBA_8SINT:
		return "rgba8sint", true
	case GPUTextureFormat_BGRA_8UNORM:
		return "bgra8unorm", true
	case GPUTextureFormat_BGRA_8UNORM_SRGB:
		return "bgra8unorm-srgb", true
	case GPUTextureFormat_RGB_9E5UFLOAT:
		return "rgb9e5ufloat", true
	case GPUTextureFormat_RGB_10A2UINT:
		return "rgb10a2uint", true
	case GPUTextureFormat_RGB_10A2UNORM:
		return "rgb10a2unorm", true
	case GPUTextureFormat_RG_11B10UFLOAT:
		return "rg11b10ufloat", true
	case GPUTextureFormat_RG_32UINT:
		return "rg32uint", true
	case GPUTextureFormat_RG_32SINT:
		return "rg32sint", true
	case GPUTextureFormat_RG_32FLOAT:
		return "rg32float", true
	case GPUTextureFormat_RGBA_16UINT:
		return "rgba16uint", true
	case GPUTextureFormat_RGBA_16SINT:
		return "rgba16sint", true
	case GPUTextureFormat_RGBA_16FLOAT:
		return "rgba16float", true
	case GPUTextureFormat_RGBA_32UINT:
		return "rgba32uint", true
	case GPUTextureFormat_RGBA_32SINT:
		return "rgba32sint", true
	case GPUTextureFormat_RGBA_32FLOAT:
		return "rgba32float", true
	case GPUTextureFormat_STENCIL8:
		return "stencil8", true
	case GPUTextureFormat_DEPTH_16UNORM:
		return "depth16unorm", true
	case GPUTextureFormat_DEPTH_24PLUS:
		return "depth24plus", true
	case GPUTextureFormat_DEPTH_24PLUS_STENCIL8:
		return "depth24plus-stencil8", true
	case GPUTextureFormat_DEPTH_32FLOAT:
		return "depth32float", true
	case GPUTextureFormat_DEPTH_32FLOAT_STENCIL8:
		return "depth32float-stencil8", true
	case GPUTextureFormat_BC1_RGBA_UNORM:
		return "bc1-rgba-unorm", true
	case GPUTextureFormat_BC1_RGBA_UNORM_SRGB:
		return "bc1-rgba-unorm-srgb", true
	case GPUTextureFormat_BC2_RGBA_UNORM:
		return "bc2-rgba-unorm", true
	case GPUTextureFormat_BC2_RGBA_UNORM_SRGB:
		return "bc2-rgba-unorm-srgb", true
	case GPUTextureFormat_BC3_RGBA_UNORM:
		return "bc3-rgba-unorm", true
	case GPUTextureFormat_BC3_RGBA_UNORM_SRGB:
		return "bc3-rgba-unorm-srgb", true
	case GPUTextureFormat_BC4_R_UNORM:
		return "bc4-r-unorm", true
	case GPUTextureFormat_BC4_R_SNORM:
		return "bc4-r-snorm", true
	case GPUTextureFormat_BC5_RG_UNORM:
		return "bc5-rg-unorm", true
	case GPUTextureFormat_BC5_RG_SNORM:
		return "bc5-rg-snorm", true
	case GPUTextureFormat_BC_6H_RGB_UFLOAT:
		return "bc6h-rgb-ufloat", true
	case GPUTextureFormat_BC_6H_RGB_FLOAT:
		return "bc6h-rgb-float", true
	case GPUTextureFormat_BC7_RGBA_UNORM:
		return "bc7-rgba-unorm", true
	case GPUTextureFormat_BC7_RGBA_UNORM_SRGB:
		return "bc7-rgba-unorm-srgb", true
	case GPUTextureFormat_ETC2_RGB_8UNORM:
		return "etc2-rgb8unorm", true
	case GPUTextureFormat_ETC2_RGB_8UNORM_SRGB:
		return "etc2-rgb8unorm-srgb", true
	case GPUTextureFormat_ETC2_RGB_8A1UNORM:
		return "etc2-rgb8a1unorm", true
	case GPUTextureFormat_ETC2_RGB_8A1UNORM_SRGB:
		return "etc2-rgb8a1unorm-srgb", true
	case GPUTextureFormat_ETC2_RGBA_8UNORM:
		return "etc2-rgba8unorm", true
	case GPUTextureFormat_ETC2_RGBA_8UNORM_SRGB:
		return "etc2-rgba8unorm-srgb", true
	case GPUTextureFormat_EAC_R_11UNORM:
		return "eac-r11unorm", true
	case GPUTextureFormat_EAC_R_11SNORM:
		return "eac-r11snorm", true
	case GPUTextureFormat_EAC_RG_11UNORM:
		return "eac-rg11unorm", true
	case GPUTextureFormat_EAC_RG_11SNORM:
		return "eac-rg11snorm", true
	case GPUTextureFormat_ASTC_4X4_UNORM:
		return "astc-4x4-unorm", true
	case GPUTextureFormat_ASTC_4X4_UNORM_SRGB:
		return "astc-4x4-unorm-srgb", true
	case GPUTextureFormat_ASTC_5X4_UNORM:
		return "astc-5x4-unorm", true
	case GPUTextureFormat_ASTC_5X4_UNORM_SRGB:
		return "astc-5x4-unorm-srgb", true
	case GPUTextureFormat_ASTC_5X5_UNORM:
		return "astc-5x5-unorm", true
	case GPUTextureFormat_ASTC_5X5_UNORM_SRGB:
		return "astc-5x5-unorm-srgb", true
	case GPUTextureFormat_ASTC_6X5_UNORM:
		return "astc-6x5-unorm", true
	case GPUTextureFormat_ASTC_6X5_UNORM_SRGB:
		return "astc-6x5-unorm-srgb", true
	case GPUTextureFormat_ASTC_6X6_UNORM:
		return "astc-6x6-unorm", true
	case GPUTextureFormat_ASTC_6X6_UNORM_SRGB:
		return "astc-6x6-unorm-srgb", true
	case GPUTextureFormat_ASTC_8X5_UNORM:
		return "astc-8x5-unorm", true
	case GPUTextureFormat_ASTC_8X5_UNORM_SRGB:
		return "astc-8x5-unorm-srgb", true
	case GPUTextureFormat_ASTC_8X6_UNORM:
		return "astc-8x6-unorm", true
	case GPUTextureFormat_ASTC_8X6_UNORM_SRGB:
		return "astc-8x6-unorm-srgb", true
	case GPUTextureFormat_ASTC_8X8_UNORM:
		return "astc-8x8-unorm", true
	case GPUTextureFormat_ASTC_8X8_UNORM_SRGB:
		return "astc-8x8-unorm-srgb", true
	case GPUTextureFormat_ASTC_10X5_UNORM:
		return "astc-10x5-unorm", true
	case GPUTextureFormat_ASTC_10X5_UNORM_SRGB:
		return "astc-10x5-unorm-srgb", true
	case GPUTextureFormat_ASTC_10X6_UNORM:
		return "astc-10x6-unorm", true
	case GPUTextureFormat_ASTC_10X6_UNORM_SRGB:
		return "astc-10x6-unorm-srgb", true
	case GPUTextureFormat_ASTC_10X8_UNORM:
		return "astc-10x8-unorm", true
	case GPUTextureFormat_ASTC_10X8_UNORM_SRGB:
		return "astc-10x8-unorm-srgb", true
	case GPUTextureFormat_ASTC_10X10_UNORM:
		return "astc-10x10-unorm", true
	case GPUTextureFormat_ASTC_10X10_UNORM_SRGB:
		return "astc-10x10-unorm-srgb", true
	case GPUTextureFormat_ASTC_12X10_UNORM:
		return "astc-12x10-unorm", true
	case GPUTextureFormat_ASTC_12X10_UNORM_SRGB:
		return "astc-12x10-unorm-srgb", true
	case GPUTextureFormat_ASTC_12X12_UNORM:
		return "astc-12x12-unorm", true
	case GPUTextureFormat_ASTC_12X12_UNORM_SRGB:
		return "astc-12x12-unorm-srgb", true
	default:
		return "", false
	}
}

type GPUTextureViewDimension uint32

const (
	_ GPUTextureViewDimension = iota

	GPUTextureViewDimension_1D
	GPUTextureViewDimension_2D
	GPUTextureViewDimension_2D_ARRAY
	GPUTextureViewDimension_CUBE
	GPUTextureViewDimension_CUBE_ARRAY
	GPUTextureViewDimension_3D
)

func (GPUTextureViewDimension) FromRef(str js.Ref) GPUTextureViewDimension {
	return GPUTextureViewDimension(bindings.ConstOfGPUTextureViewDimension(str))
}

func (x GPUTextureViewDimension) String() (string, bool) {
	switch x {
	case GPUTextureViewDimension_1D:
		return "1d", true
	case GPUTextureViewDimension_2D:
		return "2d", true
	case GPUTextureViewDimension_2D_ARRAY:
		return "2d-array", true
	case GPUTextureViewDimension_CUBE:
		return "cube", true
	case GPUTextureViewDimension_CUBE_ARRAY:
		return "cube-array", true
	case GPUTextureViewDimension_3D:
		return "3d", true
	default:
		return "", false
	}
}

type GPUTextureAspect uint32

const (
	_ GPUTextureAspect = iota

	GPUTextureAspect_ALL
	GPUTextureAspect_STENCIL_ONLY
	GPUTextureAspect_DEPTH_ONLY
)

func (GPUTextureAspect) FromRef(str js.Ref) GPUTextureAspect {
	return GPUTextureAspect(bindings.ConstOfGPUTextureAspect(str))
}

func (x GPUTextureAspect) String() (string, bool) {
	switch x {
	case GPUTextureAspect_ALL:
		return "all", true
	case GPUTextureAspect_STENCIL_ONLY:
		return "stencil-only", true
	case GPUTextureAspect_DEPTH_ONLY:
		return "depth-only", true
	default:
		return "", false
	}
}

type GPUIntegerCoordinate uint32

type GPUTextureViewDescriptor struct {
	// Format is "GPUTextureViewDescriptor.format"
	//
	// Optional
	Format GPUTextureFormat
	// Dimension is "GPUTextureViewDescriptor.dimension"
	//
	// Optional
	Dimension GPUTextureViewDimension
	// Aspect is "GPUTextureViewDescriptor.aspect"
	//
	// Optional, defaults to "all".
	Aspect GPUTextureAspect
	// BaseMipLevel is "GPUTextureViewDescriptor.baseMipLevel"
	//
	// Optional, defaults to 0.
	BaseMipLevel GPUIntegerCoordinate
	// MipLevelCount is "GPUTextureViewDescriptor.mipLevelCount"
	//
	// Optional
	MipLevelCount GPUIntegerCoordinate
	// BaseArrayLayer is "GPUTextureViewDescriptor.baseArrayLayer"
	//
	// Optional, defaults to 0.
	BaseArrayLayer GPUIntegerCoordinate
	// ArrayLayerCount is "GPUTextureViewDescriptor.arrayLayerCount"
	//
	// Optional
	ArrayLayerCount GPUIntegerCoordinate
	// Label is "GPUTextureViewDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_BaseMipLevel    bool // for BaseMipLevel.
	FFI_USE_MipLevelCount   bool // for MipLevelCount.
	FFI_USE_BaseArrayLayer  bool // for BaseArrayLayer.
	FFI_USE_ArrayLayerCount bool // for ArrayLayerCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUTextureViewDescriptor with all fields set.
func (p GPUTextureViewDescriptor) FromRef(ref js.Ref) GPUTextureViewDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUTextureViewDescriptor GPUTextureViewDescriptor [// GPUTextureViewDescriptor] [0x14000313860 0x14000313900 0x140003139a0 0x14000313a40 0x14000313b80 0x14000313cc0 0x14000313e00 0x14000313f40 0x14000313ae0 0x14000313c20 0x14000313d60 0x14000313ea0] 0x1400037c678 {0 0}} in the application heap.
func (p GPUTextureViewDescriptor) New() js.Ref {
	return bindings.GPUTextureViewDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUTextureViewDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureViewDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUTextureViewDescriptor) Update(ref js.Ref) {
	bindings.GPUTextureViewDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUIntegerCoordinateOut uint32

type GPUSize32Out uint32

type GPUTextureDimension uint32

const (
	_ GPUTextureDimension = iota

	GPUTextureDimension_1D
	GPUTextureDimension_2D
	GPUTextureDimension_3D
)

func (GPUTextureDimension) FromRef(str js.Ref) GPUTextureDimension {
	return GPUTextureDimension(bindings.ConstOfGPUTextureDimension(str))
}

func (x GPUTextureDimension) String() (string, bool) {
	switch x {
	case GPUTextureDimension_1D:
		return "1d", true
	case GPUTextureDimension_2D:
		return "2d", true
	case GPUTextureDimension_3D:
		return "3d", true
	default:
		return "", false
	}
}

type GPUTexture struct {
	ref js.Ref
}

func (this GPUTexture) Once() GPUTexture {
	this.Ref().Once()
	return this
}

func (this GPUTexture) Ref() js.Ref {
	return this.ref
}

func (this GPUTexture) FromRef(ref js.Ref) GPUTexture {
	this.ref = ref
	return this
}

func (this GPUTexture) Free() {
	this.Ref().Free()
}

// Width returns the value of property "GPUTexture.width".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Width() (GPUIntegerCoordinateOut, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUIntegerCoordinateOut(_ret), _ok
}

// Height returns the value of property "GPUTexture.height".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Height() (GPUIntegerCoordinateOut, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUIntegerCoordinateOut(_ret), _ok
}

// DepthOrArrayLayers returns the value of property "GPUTexture.depthOrArrayLayers".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) DepthOrArrayLayers() (GPUIntegerCoordinateOut, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureDepthOrArrayLayers(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUIntegerCoordinateOut(_ret), _ok
}

// MipLevelCount returns the value of property "GPUTexture.mipLevelCount".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) MipLevelCount() (GPUIntegerCoordinateOut, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureMipLevelCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUIntegerCoordinateOut(_ret), _ok
}

// SampleCount returns the value of property "GPUTexture.sampleCount".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) SampleCount() (GPUSize32Out, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureSampleCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSize32Out(_ret), _ok
}

// Dimension returns the value of property "GPUTexture.dimension".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Dimension() (GPUTextureDimension, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureDimension(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUTextureDimension(_ret), _ok
}

// Format returns the value of property "GPUTexture.format".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Format() (GPUTextureFormat, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureFormat(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUTextureFormat(_ret), _ok
}

// Usage returns the value of property "GPUTexture.usage".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Usage() (GPUFlagsConstant, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureUsage(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUFlagsConstant(_ret), _ok
}

// Label returns the value of property "GPUTexture.label".
//
// The returned bool will be false if there is no such property.
func (this GPUTexture) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUTextureLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUTexture.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUTexture) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUTextureLabel(
		this.Ref(),
		val.Ref(),
	)
}

// CreateView calls the method "GPUTexture.createView".
//
// The returned bool will be false if there is no such method.
func (this GPUTexture) CreateView(descriptor GPUTextureViewDescriptor) (GPUTextureView, bool) {
	var _ok bool
	_ret := bindings.CallGPUTextureCreateView(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUTextureView{}.FromRef(_ret), _ok
}

// CreateViewFunc returns the method "GPUTexture.createView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUTexture) CreateViewFunc() (fn js.Func[func(descriptor GPUTextureViewDescriptor) GPUTextureView]) {
	return fn.FromRef(
		bindings.GPUTextureCreateViewFunc(
			this.Ref(),
		),
	)
}

// CreateView1 calls the method "GPUTexture.createView".
//
// The returned bool will be false if there is no such method.
func (this GPUTexture) CreateView1() (GPUTextureView, bool) {
	var _ok bool
	_ret := bindings.CallGPUTextureCreateView1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUTextureView{}.FromRef(_ret), _ok
}

// CreateView1Func returns the method "GPUTexture.createView".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUTexture) CreateView1Func() (fn js.Func[func() GPUTextureView]) {
	return fn.FromRef(
		bindings.GPUTextureCreateView1Func(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUTexture.destroy".
//
// The returned bool will be false if there is no such method.
func (this GPUTexture) Destroy() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUTextureDestroy(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DestroyFunc returns the method "GPUTexture.destroy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUTexture) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUTextureDestroyFunc(
			this.Ref(),
		),
	)
}
