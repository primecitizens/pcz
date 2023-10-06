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
	//
	// NOTE: FFI_USE_FullRange MUST be set to true to make this field effective.
	FullRange bool

	FFI_USE_FullRange bool // for FullRange.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoColorSpaceInit with all fields set.
func (p VideoColorSpaceInit) FromRef(ref js.Ref) VideoColorSpaceInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoColorSpaceInit in the application heap.
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
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration uint64
	// Layout is "VideoFrameBufferInit.layout"
	//
	// Optional
	Layout js.Array[PlaneLayout]
	// VisibleRect is "VideoFrameBufferInit.visibleRect"
	//
	// Optional
	//
	// NOTE: VisibleRect.FFI_USE MUST be set to true to get VisibleRect used.
	VisibleRect DOMRectInit
	// DisplayWidth is "VideoFrameBufferInit.displayWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayWidth MUST be set to true to make this field effective.
	DisplayWidth uint32
	// DisplayHeight is "VideoFrameBufferInit.displayHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayHeight MUST be set to true to make this field effective.
	DisplayHeight uint32
	// ColorSpace is "VideoFrameBufferInit.colorSpace"
	//
	// Optional
	//
	// NOTE: ColorSpace.FFI_USE MUST be set to true to get ColorSpace used.
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

// New creates a new VideoFrameBufferInit in the application heap.
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
	//
	// NOTE: Rect.FFI_USE MUST be set to true to get Rect used.
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

// New creates a new VideoFrameCopyToOptions in the application heap.
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

func NewVideoColorSpace(init VideoColorSpaceInit) (ret VideoColorSpace) {
	ret.ref = bindings.NewVideoColorSpaceByVideoColorSpace(
		js.Pointer(&init))
	return
}

func NewVideoColorSpaceByVideoColorSpace1() (ret VideoColorSpace) {
	ret.ref = bindings.NewVideoColorSpaceByVideoColorSpace1()
	return
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
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Primaries() (ret VideoColorPrimaries, ok bool) {
	ok = js.True == bindings.GetVideoColorSpacePrimaries(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transfer returns the value of property "VideoColorSpace.transfer".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Transfer() (ret VideoTransferCharacteristics, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceTransfer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "VideoColorSpace.matrix".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Matrix() (ret VideoMatrixCoefficients, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceMatrix(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FullRange returns the value of property "VideoColorSpace.fullRange".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) FullRange() (ret bool, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceFullRange(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "VideoColorSpace.toJSON" exists.
func (this VideoColorSpace) HasToJSON() bool {
	return js.True == bindings.HasVideoColorSpaceToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "VideoColorSpace.toJSON".
func (this VideoColorSpace) ToJSONFunc() (fn js.Func[func() VideoColorSpaceInit]) {
	return fn.FromRef(
		bindings.VideoColorSpaceToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "VideoColorSpace.toJSON".
func (this VideoColorSpace) ToJSON() (ret VideoColorSpaceInit) {
	bindings.CallVideoColorSpaceToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "VideoColorSpace.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoColorSpace) TryToJSON() (ret VideoColorSpaceInit, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoColorSpaceToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewVideoFrame(image CanvasImageSource, init VideoFrameInit) (ret VideoFrame) {
	ret.ref = bindings.NewVideoFrameByVideoFrame(
		image.Ref(),
		js.Pointer(&init))
	return
}

func NewVideoFrameByVideoFrame1(image CanvasImageSource) (ret VideoFrame) {
	ret.ref = bindings.NewVideoFrameByVideoFrame1(
		image.Ref())
	return
}

func NewVideoFrameByVideoFrame2(data AllowSharedBufferSource, init VideoFrameBufferInit) (ret VideoFrame) {
	ret.ref = bindings.NewVideoFrameByVideoFrame2(
		data.Ref(),
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this VideoFrame) Format() (ret VideoPixelFormat, ok bool) {
	ok = js.True == bindings.GetVideoFrameFormat(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CodedWidth returns the value of property "VideoFrame.codedWidth".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CodedHeight returns the value of property "VideoFrame.codedHeight".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CodedRect returns the value of property "VideoFrame.codedRect".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VisibleRect returns the value of property "VideoFrame.visibleRect".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) VisibleRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetVideoFrameVisibleRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DisplayWidth returns the value of property "VideoFrame.displayWidth".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) DisplayWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameDisplayWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DisplayHeight returns the value of property "VideoFrame.displayHeight".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) DisplayHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameDisplayHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "VideoFrame.duration".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) Duration() (ret uint64, ok bool) {
	ok = js.True == bindings.GetVideoFrameDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "VideoFrame.timestamp".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) Timestamp() (ret int64, ok bool) {
	ok = js.True == bindings.GetVideoFrameTimestamp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColorSpace returns the value of property "VideoFrame.colorSpace".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) ColorSpace() (ret VideoColorSpace, ok bool) {
	ok = js.True == bindings.GetVideoFrameColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasMetadata returns true if the method "VideoFrame.metadata" exists.
func (this VideoFrame) HasMetadata() bool {
	return js.True == bindings.HasVideoFrameMetadata(
		this.Ref(),
	)
}

// MetadataFunc returns the method "VideoFrame.metadata".
func (this VideoFrame) MetadataFunc() (fn js.Func[func() VideoFrameMetadata]) {
	return fn.FromRef(
		bindings.VideoFrameMetadataFunc(
			this.Ref(),
		),
	)
}

// Metadata calls the method "VideoFrame.metadata".
func (this VideoFrame) Metadata() (ret VideoFrameMetadata) {
	bindings.CallVideoFrameMetadata(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMetadata calls the method "VideoFrame.metadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryMetadata() (ret VideoFrameMetadata, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameMetadata(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAllocationSize returns true if the method "VideoFrame.allocationSize" exists.
func (this VideoFrame) HasAllocationSize() bool {
	return js.True == bindings.HasVideoFrameAllocationSize(
		this.Ref(),
	)
}

// AllocationSizeFunc returns the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSizeFunc() (fn js.Func[func(options VideoFrameCopyToOptions) uint32]) {
	return fn.FromRef(
		bindings.VideoFrameAllocationSizeFunc(
			this.Ref(),
		),
	)
}

// AllocationSize calls the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSize(options VideoFrameCopyToOptions) (ret uint32) {
	bindings.CallVideoFrameAllocationSize(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAllocationSize calls the method "VideoFrame.allocationSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryAllocationSize(options VideoFrameCopyToOptions) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameAllocationSize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasAllocationSize1 returns true if the method "VideoFrame.allocationSize" exists.
func (this VideoFrame) HasAllocationSize1() bool {
	return js.True == bindings.HasVideoFrameAllocationSize1(
		this.Ref(),
	)
}

// AllocationSize1Func returns the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSize1Func() (fn js.Func[func() uint32]) {
	return fn.FromRef(
		bindings.VideoFrameAllocationSize1Func(
			this.Ref(),
		),
	)
}

// AllocationSize1 calls the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSize1() (ret uint32) {
	bindings.CallVideoFrameAllocationSize1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAllocationSize1 calls the method "VideoFrame.allocationSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryAllocationSize1() (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameAllocationSize1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCopyTo returns true if the method "VideoFrame.copyTo" exists.
func (this VideoFrame) HasCopyTo() bool {
	return js.True == bindings.HasVideoFrameCopyTo(
		this.Ref(),
	)
}

// CopyToFunc returns the method "VideoFrame.copyTo".
func (this VideoFrame) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) js.Promise[js.Array[PlaneLayout]]]) {
	return fn.FromRef(
		bindings.VideoFrameCopyToFunc(
			this.Ref(),
		),
	)
}

// CopyTo calls the method "VideoFrame.copyTo".
func (this VideoFrame) CopyTo(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) (ret js.Promise[js.Array[PlaneLayout]]) {
	bindings.CallVideoFrameCopyTo(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCopyTo calls the method "VideoFrame.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryCopyTo(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) (ret js.Promise[js.Array[PlaneLayout]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameCopyTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCopyTo1 returns true if the method "VideoFrame.copyTo" exists.
func (this VideoFrame) HasCopyTo1() bool {
	return js.True == bindings.HasVideoFrameCopyTo1(
		this.Ref(),
	)
}

// CopyTo1Func returns the method "VideoFrame.copyTo".
func (this VideoFrame) CopyTo1Func() (fn js.Func[func(destination AllowSharedBufferSource) js.Promise[js.Array[PlaneLayout]]]) {
	return fn.FromRef(
		bindings.VideoFrameCopyTo1Func(
			this.Ref(),
		),
	)
}

// CopyTo1 calls the method "VideoFrame.copyTo".
func (this VideoFrame) CopyTo1(destination AllowSharedBufferSource) (ret js.Promise[js.Array[PlaneLayout]]) {
	bindings.CallVideoFrameCopyTo1(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryCopyTo1 calls the method "VideoFrame.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryCopyTo1(destination AllowSharedBufferSource) (ret js.Promise[js.Array[PlaneLayout]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameCopyTo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
}

// HasClone returns true if the method "VideoFrame.clone" exists.
func (this VideoFrame) HasClone() bool {
	return js.True == bindings.HasVideoFrameClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "VideoFrame.clone".
func (this VideoFrame) CloneFunc() (fn js.Func[func() VideoFrame]) {
	return fn.FromRef(
		bindings.VideoFrameCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "VideoFrame.clone".
func (this VideoFrame) Clone() (ret VideoFrame) {
	bindings.CallVideoFrameClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "VideoFrame.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryClone() (ret VideoFrame, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "VideoFrame.close" exists.
func (this VideoFrame) HasClose() bool {
	return js.True == bindings.HasVideoFrameClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "VideoFrame.close".
func (this VideoFrame) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoFrameCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "VideoFrame.close".
func (this VideoFrame) Close() (ret js.Void) {
	bindings.CallVideoFrameClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "VideoFrame.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextCanvas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferWidth returns the value of property "WebGLRenderingContext.drawingBufferWidth".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferWidth() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferHeight returns the value of property "WebGLRenderingContext.drawingBufferHeight".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferHeight() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferColorSpace returns the value of property "WebGLRenderingContext.drawingBufferColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDrawingBufferColorSpace sets the value of property "WebGLRenderingContext.drawingBufferColorSpace" to val.
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
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) UnpackColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextUnpackColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUnpackColorSpace sets the value of property "WebGLRenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGLRenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGLRenderingContextUnpackColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// HasGetContextAttributes returns true if the method "WebGLRenderingContext.getContextAttributes" exists.
func (this WebGLRenderingContext) HasGetContextAttributes() bool {
	return js.True == bindings.HasWebGLRenderingContextGetContextAttributes(
		this.Ref(),
	)
}

// GetContextAttributesFunc returns the method "WebGLRenderingContext.getContextAttributes".
func (this WebGLRenderingContext) GetContextAttributesFunc() (fn js.Func[func() WebGLContextAttributes]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// GetContextAttributes calls the method "WebGLRenderingContext.getContextAttributes".
func (this WebGLRenderingContext) GetContextAttributes() (ret WebGLContextAttributes) {
	bindings.CallWebGLRenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "WebGLRenderingContext.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetContextAttributes() (ret WebGLContextAttributes, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsContextLost returns true if the method "WebGLRenderingContext.isContextLost" exists.
func (this WebGLRenderingContext) HasIsContextLost() bool {
	return js.True == bindings.HasWebGLRenderingContextIsContextLost(
		this.Ref(),
	)
}

// IsContextLostFunc returns the method "WebGLRenderingContext.isContextLost".
func (this WebGLRenderingContext) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsContextLostFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "WebGLRenderingContext.isContextLost".
func (this WebGLRenderingContext) IsContextLost() (ret bool) {
	bindings.CallWebGLRenderingContextIsContextLost(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "WebGLRenderingContext.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsContextLost(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSupportedExtensions returns true if the method "WebGLRenderingContext.getSupportedExtensions" exists.
func (this WebGLRenderingContext) HasGetSupportedExtensions() bool {
	return js.True == bindings.HasWebGLRenderingContextGetSupportedExtensions(
		this.Ref(),
	)
}

// GetSupportedExtensionsFunc returns the method "WebGLRenderingContext.getSupportedExtensions".
func (this WebGLRenderingContext) GetSupportedExtensionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetSupportedExtensionsFunc(
			this.Ref(),
		),
	)
}

// GetSupportedExtensions calls the method "WebGLRenderingContext.getSupportedExtensions".
func (this WebGLRenderingContext) GetSupportedExtensions() (ret js.Array[js.String]) {
	bindings.CallWebGLRenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedExtensions calls the method "WebGLRenderingContext.getSupportedExtensions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetSupportedExtensions() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetExtension returns true if the method "WebGLRenderingContext.getExtension" exists.
func (this WebGLRenderingContext) HasGetExtension() bool {
	return js.True == bindings.HasWebGLRenderingContextGetExtension(
		this.Ref(),
	)
}

// GetExtensionFunc returns the method "WebGLRenderingContext.getExtension".
func (this WebGLRenderingContext) GetExtensionFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetExtensionFunc(
			this.Ref(),
		),
	)
}

// GetExtension calls the method "WebGLRenderingContext.getExtension".
func (this WebGLRenderingContext) GetExtension(name js.String) (ret js.Object) {
	bindings.CallWebGLRenderingContextGetExtension(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetExtension calls the method "WebGLRenderingContext.getExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetExtension(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetExtension(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasActiveTexture returns true if the method "WebGLRenderingContext.activeTexture" exists.
func (this WebGLRenderingContext) HasActiveTexture() bool {
	return js.True == bindings.HasWebGLRenderingContextActiveTexture(
		this.Ref(),
	)
}

// ActiveTextureFunc returns the method "WebGLRenderingContext.activeTexture".
func (this WebGLRenderingContext) ActiveTextureFunc() (fn js.Func[func(texture GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextActiveTextureFunc(
			this.Ref(),
		),
	)
}

// ActiveTexture calls the method "WebGLRenderingContext.activeTexture".
func (this WebGLRenderingContext) ActiveTexture(texture GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextActiveTexture(
		this.Ref(), js.Pointer(&ret),
		uint32(texture),
	)

	return
}

// TryActiveTexture calls the method "WebGLRenderingContext.activeTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryActiveTexture(texture GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextActiveTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(texture),
	)

	return
}

// HasAttachShader returns true if the method "WebGLRenderingContext.attachShader" exists.
func (this WebGLRenderingContext) HasAttachShader() bool {
	return js.True == bindings.HasWebGLRenderingContextAttachShader(
		this.Ref(),
	)
}

// AttachShaderFunc returns the method "WebGLRenderingContext.attachShader".
func (this WebGLRenderingContext) AttachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextAttachShaderFunc(
			this.Ref(),
		),
	)
}

// AttachShader calls the method "WebGLRenderingContext.attachShader".
func (this WebGLRenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextAttachShader(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// TryAttachShader calls the method "WebGLRenderingContext.attachShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryAttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextAttachShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasBindAttribLocation returns true if the method "WebGLRenderingContext.bindAttribLocation" exists.
func (this WebGLRenderingContext) HasBindAttribLocation() bool {
	return js.True == bindings.HasWebGLRenderingContextBindAttribLocation(
		this.Ref(),
	)
}

// BindAttribLocationFunc returns the method "WebGLRenderingContext.bindAttribLocation".
func (this WebGLRenderingContext) BindAttribLocationFunc() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindAttribLocationFunc(
			this.Ref(),
		),
	)
}

// BindAttribLocation calls the method "WebGLRenderingContext.bindAttribLocation".
func (this WebGLRenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// TryBindAttribLocation calls the method "WebGLRenderingContext.bindAttribLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// HasBindBuffer returns true if the method "WebGLRenderingContext.bindBuffer" exists.
func (this WebGLRenderingContext) HasBindBuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextBindBuffer(
		this.Ref(),
	)
}

// BindBufferFunc returns the method "WebGLRenderingContext.bindBuffer".
func (this WebGLRenderingContext) BindBufferFunc() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindBufferFunc(
			this.Ref(),
		),
	)
}

// BindBuffer calls the method "WebGLRenderingContext.bindBuffer".
func (this WebGLRenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// TryBindBuffer calls the method "WebGLRenderingContext.bindBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBindBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// HasBindFramebuffer returns true if the method "WebGLRenderingContext.bindFramebuffer" exists.
func (this WebGLRenderingContext) HasBindFramebuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextBindFramebuffer(
		this.Ref(),
	)
}

// BindFramebufferFunc returns the method "WebGLRenderingContext.bindFramebuffer".
func (this WebGLRenderingContext) BindFramebufferFunc() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindFramebufferFunc(
			this.Ref(),
		),
	)
}

// BindFramebuffer calls the method "WebGLRenderingContext.bindFramebuffer".
func (this WebGLRenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// TryBindFramebuffer calls the method "WebGLRenderingContext.bindFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// HasBindRenderbuffer returns true if the method "WebGLRenderingContext.bindRenderbuffer" exists.
func (this WebGLRenderingContext) HasBindRenderbuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextBindRenderbuffer(
		this.Ref(),
	)
}

// BindRenderbufferFunc returns the method "WebGLRenderingContext.bindRenderbuffer".
func (this WebGLRenderingContext) BindRenderbufferFunc() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindRenderbufferFunc(
			this.Ref(),
		),
	)
}

// BindRenderbuffer calls the method "WebGLRenderingContext.bindRenderbuffer".
func (this WebGLRenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// TryBindRenderbuffer calls the method "WebGLRenderingContext.bindRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// HasBindTexture returns true if the method "WebGLRenderingContext.bindTexture" exists.
func (this WebGLRenderingContext) HasBindTexture() bool {
	return js.True == bindings.HasWebGLRenderingContextBindTexture(
		this.Ref(),
	)
}

// BindTextureFunc returns the method "WebGLRenderingContext.bindTexture".
func (this WebGLRenderingContext) BindTextureFunc() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBindTextureFunc(
			this.Ref(),
		),
	)
}

// BindTexture calls the method "WebGLRenderingContext.bindTexture".
func (this WebGLRenderingContext) BindTexture(target GLenum, texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindTexture(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		texture.Ref(),
	)

	return
}

// TryBindTexture calls the method "WebGLRenderingContext.bindTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBindTexture(target GLenum, texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBindTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		texture.Ref(),
	)

	return
}

// HasBlendColor returns true if the method "WebGLRenderingContext.blendColor" exists.
func (this WebGLRenderingContext) HasBlendColor() bool {
	return js.True == bindings.HasWebGLRenderingContextBlendColor(
		this.Ref(),
	)
}

// BlendColorFunc returns the method "WebGLRenderingContext.blendColor".
func (this WebGLRenderingContext) BlendColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendColorFunc(
			this.Ref(),
		),
	)
}

// BlendColor calls the method "WebGLRenderingContext.blendColor".
func (this WebGLRenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendColor(
		this.Ref(), js.Pointer(&ret),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// TryBlendColor calls the method "WebGLRenderingContext.blendColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendColor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasBlendEquation returns true if the method "WebGLRenderingContext.blendEquation" exists.
func (this WebGLRenderingContext) HasBlendEquation() bool {
	return js.True == bindings.HasWebGLRenderingContextBlendEquation(
		this.Ref(),
	)
}

// BlendEquationFunc returns the method "WebGLRenderingContext.blendEquation".
func (this WebGLRenderingContext) BlendEquationFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendEquationFunc(
			this.Ref(),
		),
	)
}

// BlendEquation calls the method "WebGLRenderingContext.blendEquation".
func (this WebGLRenderingContext) BlendEquation(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendEquation(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryBlendEquation calls the method "WebGLRenderingContext.blendEquation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendEquation(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendEquation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasBlendEquationSeparate returns true if the method "WebGLRenderingContext.blendEquationSeparate" exists.
func (this WebGLRenderingContext) HasBlendEquationSeparate() bool {
	return js.True == bindings.HasWebGLRenderingContextBlendEquationSeparate(
		this.Ref(),
	)
}

// BlendEquationSeparateFunc returns the method "WebGLRenderingContext.blendEquationSeparate".
func (this WebGLRenderingContext) BlendEquationSeparateFunc() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendEquationSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparate calls the method "WebGLRenderingContext.blendEquationSeparate".
func (this WebGLRenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// TryBlendEquationSeparate calls the method "WebGLRenderingContext.blendEquationSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// HasBlendFunc returns true if the method "WebGLRenderingContext.blendFunc" exists.
func (this WebGLRenderingContext) HasBlendFunc() bool {
	return js.True == bindings.HasWebGLRenderingContextBlendFunc(
		this.Ref(),
	)
}

// BlendFuncFunc returns the method "WebGLRenderingContext.blendFunc".
func (this WebGLRenderingContext) BlendFuncFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendFuncFunc(
			this.Ref(),
		),
	)
}

// BlendFunc calls the method "WebGLRenderingContext.blendFunc".
func (this WebGLRenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// TryBlendFunc calls the method "WebGLRenderingContext.blendFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// HasBlendFuncSeparate returns true if the method "WebGLRenderingContext.blendFuncSeparate" exists.
func (this WebGLRenderingContext) HasBlendFuncSeparate() bool {
	return js.True == bindings.HasWebGLRenderingContextBlendFuncSeparate(
		this.Ref(),
	)
}

// BlendFuncSeparateFunc returns the method "WebGLRenderingContext.blendFuncSeparate".
func (this WebGLRenderingContext) BlendFuncSeparateFunc() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBlendFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparate calls the method "WebGLRenderingContext.blendFuncSeparate".
func (this WebGLRenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// TryBlendFuncSeparate calls the method "WebGLRenderingContext.blendFuncSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// HasCheckFramebufferStatus returns true if the method "WebGLRenderingContext.checkFramebufferStatus" exists.
func (this WebGLRenderingContext) HasCheckFramebufferStatus() bool {
	return js.True == bindings.HasWebGLRenderingContextCheckFramebufferStatus(
		this.Ref(),
	)
}

// CheckFramebufferStatusFunc returns the method "WebGLRenderingContext.checkFramebufferStatus".
func (this WebGLRenderingContext) CheckFramebufferStatusFunc() (fn js.Func[func(target GLenum) GLenum]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCheckFramebufferStatusFunc(
			this.Ref(),
		),
	)
}

// CheckFramebufferStatus calls the method "WebGLRenderingContext.checkFramebufferStatus".
func (this WebGLRenderingContext) CheckFramebufferStatus(target GLenum) (ret GLenum) {
	bindings.CallWebGLRenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryCheckFramebufferStatus calls the method "WebGLRenderingContext.checkFramebufferStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCheckFramebufferStatus(target GLenum) (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasClear returns true if the method "WebGLRenderingContext.clear" exists.
func (this WebGLRenderingContext) HasClear() bool {
	return js.True == bindings.HasWebGLRenderingContextClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "WebGLRenderingContext.clear".
func (this WebGLRenderingContext) ClearFunc() (fn js.Func[func(mask GLbitfield)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "WebGLRenderingContext.clear".
func (this WebGLRenderingContext) Clear(mask GLbitfield) (ret js.Void) {
	bindings.CallWebGLRenderingContextClear(
		this.Ref(), js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryClear calls the method "WebGLRenderingContext.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClear(mask GLbitfield) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasClearColor returns true if the method "WebGLRenderingContext.clearColor" exists.
func (this WebGLRenderingContext) HasClearColor() bool {
	return js.True == bindings.HasWebGLRenderingContextClearColor(
		this.Ref(),
	)
}

// ClearColorFunc returns the method "WebGLRenderingContext.clearColor".
func (this WebGLRenderingContext) ClearColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearColorFunc(
			this.Ref(),
		),
	)
}

// ClearColor calls the method "WebGLRenderingContext.clearColor".
func (this WebGLRenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearColor(
		this.Ref(), js.Pointer(&ret),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// TryClearColor calls the method "WebGLRenderingContext.clearColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClearColor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasClearDepth returns true if the method "WebGLRenderingContext.clearDepth" exists.
func (this WebGLRenderingContext) HasClearDepth() bool {
	return js.True == bindings.HasWebGLRenderingContextClearDepth(
		this.Ref(),
	)
}

// ClearDepthFunc returns the method "WebGLRenderingContext.clearDepth".
func (this WebGLRenderingContext) ClearDepthFunc() (fn js.Func[func(depth GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearDepthFunc(
			this.Ref(),
		),
	)
}

// ClearDepth calls the method "WebGLRenderingContext.clearDepth".
func (this WebGLRenderingContext) ClearDepth(depth GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearDepth(
		this.Ref(), js.Pointer(&ret),
		float32(depth),
	)

	return
}

// TryClearDepth calls the method "WebGLRenderingContext.clearDepth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClearDepth(depth GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClearDepth(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(depth),
	)

	return
}

// HasClearStencil returns true if the method "WebGLRenderingContext.clearStencil" exists.
func (this WebGLRenderingContext) HasClearStencil() bool {
	return js.True == bindings.HasWebGLRenderingContextClearStencil(
		this.Ref(),
	)
}

// ClearStencilFunc returns the method "WebGLRenderingContext.clearStencil".
func (this WebGLRenderingContext) ClearStencilFunc() (fn js.Func[func(s GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextClearStencilFunc(
			this.Ref(),
		),
	)
}

// ClearStencil calls the method "WebGLRenderingContext.clearStencil".
func (this WebGLRenderingContext) ClearStencil(s GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearStencil(
		this.Ref(), js.Pointer(&ret),
		int32(s),
	)

	return
}

// TryClearStencil calls the method "WebGLRenderingContext.clearStencil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClearStencil(s GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClearStencil(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(s),
	)

	return
}

// HasColorMask returns true if the method "WebGLRenderingContext.colorMask" exists.
func (this WebGLRenderingContext) HasColorMask() bool {
	return js.True == bindings.HasWebGLRenderingContextColorMask(
		this.Ref(),
	)
}

// ColorMaskFunc returns the method "WebGLRenderingContext.colorMask".
func (this WebGLRenderingContext) ColorMaskFunc() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextColorMaskFunc(
			this.Ref(),
		),
	)
}

// ColorMask calls the method "WebGLRenderingContext.colorMask".
func (this WebGLRenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextColorMask(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// TryColorMask calls the method "WebGLRenderingContext.colorMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextColorMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// HasCompileShader returns true if the method "WebGLRenderingContext.compileShader" exists.
func (this WebGLRenderingContext) HasCompileShader() bool {
	return js.True == bindings.HasWebGLRenderingContextCompileShader(
		this.Ref(),
	)
}

// CompileShaderFunc returns the method "WebGLRenderingContext.compileShader".
func (this WebGLRenderingContext) CompileShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompileShaderFunc(
			this.Ref(),
		),
	)
}

// CompileShader calls the method "WebGLRenderingContext.compileShader".
func (this WebGLRenderingContext) CompileShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompileShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryCompileShader calls the method "WebGLRenderingContext.compileShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCompileShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCompileShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasCopyTexImage2D returns true if the method "WebGLRenderingContext.copyTexImage2D" exists.
func (this WebGLRenderingContext) HasCopyTexImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextCopyTexImage2D(
		this.Ref(),
	)
}

// CopyTexImage2DFunc returns the method "WebGLRenderingContext.copyTexImage2D".
func (this WebGLRenderingContext) CopyTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCopyTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexImage2D calls the method "WebGLRenderingContext.copyTexImage2D".
func (this WebGLRenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	return
}

// TryCopyTexImage2D calls the method "WebGLRenderingContext.copyTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	return
}

// HasCopyTexSubImage2D returns true if the method "WebGLRenderingContext.copyTexSubImage2D" exists.
func (this WebGLRenderingContext) HasCopyTexSubImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextCopyTexSubImage2D(
		this.Ref(),
	)
}

// CopyTexSubImage2DFunc returns the method "WebGLRenderingContext.copyTexSubImage2D".
func (this WebGLRenderingContext) CopyTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCopyTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexSubImage2D calls the method "WebGLRenderingContext.copyTexSubImage2D".
func (this WebGLRenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryCopyTexSubImage2D calls the method "WebGLRenderingContext.copyTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasCreateBuffer returns true if the method "WebGLRenderingContext.createBuffer" exists.
func (this WebGLRenderingContext) HasCreateBuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateBuffer(
		this.Ref(),
	)
}

// CreateBufferFunc returns the method "WebGLRenderingContext.createBuffer".
func (this WebGLRenderingContext) CreateBufferFunc() (fn js.Func[func() WebGLBuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "WebGLRenderingContext.createBuffer".
func (this WebGLRenderingContext) CreateBuffer() (ret WebGLBuffer) {
	bindings.CallWebGLRenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateBuffer calls the method "WebGLRenderingContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateBuffer() (ret WebGLBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateFramebuffer returns true if the method "WebGLRenderingContext.createFramebuffer" exists.
func (this WebGLRenderingContext) HasCreateFramebuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateFramebuffer(
		this.Ref(),
	)
}

// CreateFramebufferFunc returns the method "WebGLRenderingContext.createFramebuffer".
func (this WebGLRenderingContext) CreateFramebufferFunc() (fn js.Func[func() WebGLFramebuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateFramebufferFunc(
			this.Ref(),
		),
	)
}

// CreateFramebuffer calls the method "WebGLRenderingContext.createFramebuffer".
func (this WebGLRenderingContext) CreateFramebuffer() (ret WebGLFramebuffer) {
	bindings.CallWebGLRenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateFramebuffer calls the method "WebGLRenderingContext.createFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateFramebuffer() (ret WebGLFramebuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateProgram returns true if the method "WebGLRenderingContext.createProgram" exists.
func (this WebGLRenderingContext) HasCreateProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateProgram(
		this.Ref(),
	)
}

// CreateProgramFunc returns the method "WebGLRenderingContext.createProgram".
func (this WebGLRenderingContext) CreateProgramFunc() (fn js.Func[func() WebGLProgram]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateProgramFunc(
			this.Ref(),
		),
	)
}

// CreateProgram calls the method "WebGLRenderingContext.createProgram".
func (this WebGLRenderingContext) CreateProgram() (ret WebGLProgram) {
	bindings.CallWebGLRenderingContextCreateProgram(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateProgram calls the method "WebGLRenderingContext.createProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateProgram() (ret WebGLProgram, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateRenderbuffer returns true if the method "WebGLRenderingContext.createRenderbuffer" exists.
func (this WebGLRenderingContext) HasCreateRenderbuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateRenderbuffer(
		this.Ref(),
	)
}

// CreateRenderbufferFunc returns the method "WebGLRenderingContext.createRenderbuffer".
func (this WebGLRenderingContext) CreateRenderbufferFunc() (fn js.Func[func() WebGLRenderbuffer]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateRenderbufferFunc(
			this.Ref(),
		),
	)
}

// CreateRenderbuffer calls the method "WebGLRenderingContext.createRenderbuffer".
func (this WebGLRenderingContext) CreateRenderbuffer() (ret WebGLRenderbuffer) {
	bindings.CallWebGLRenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateRenderbuffer calls the method "WebGLRenderingContext.createRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateRenderbuffer() (ret WebGLRenderbuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateShader returns true if the method "WebGLRenderingContext.createShader" exists.
func (this WebGLRenderingContext) HasCreateShader() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateShader(
		this.Ref(),
	)
}

// CreateShaderFunc returns the method "WebGLRenderingContext.createShader".
func (this WebGLRenderingContext) CreateShaderFunc() (fn js.Func[func(typ GLenum) WebGLShader]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateShaderFunc(
			this.Ref(),
		),
	)
}

// CreateShader calls the method "WebGLRenderingContext.createShader".
func (this WebGLRenderingContext) CreateShader(typ GLenum) (ret WebGLShader) {
	bindings.CallWebGLRenderingContextCreateShader(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCreateShader calls the method "WebGLRenderingContext.createShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateShader(typ GLenum) (ret WebGLShader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasCreateTexture returns true if the method "WebGLRenderingContext.createTexture" exists.
func (this WebGLRenderingContext) HasCreateTexture() bool {
	return js.True == bindings.HasWebGLRenderingContextCreateTexture(
		this.Ref(),
	)
}

// CreateTextureFunc returns the method "WebGLRenderingContext.createTexture".
func (this WebGLRenderingContext) CreateTextureFunc() (fn js.Func[func() WebGLTexture]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "WebGLRenderingContext.createTexture".
func (this WebGLRenderingContext) CreateTexture() (ret WebGLTexture) {
	bindings.CallWebGLRenderingContextCreateTexture(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTexture calls the method "WebGLRenderingContext.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateTexture() (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCullFace returns true if the method "WebGLRenderingContext.cullFace" exists.
func (this WebGLRenderingContext) HasCullFace() bool {
	return js.True == bindings.HasWebGLRenderingContextCullFace(
		this.Ref(),
	)
}

// CullFaceFunc returns the method "WebGLRenderingContext.cullFace".
func (this WebGLRenderingContext) CullFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCullFaceFunc(
			this.Ref(),
		),
	)
}

// CullFace calls the method "WebGLRenderingContext.cullFace".
func (this WebGLRenderingContext) CullFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextCullFace(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryCullFace calls the method "WebGLRenderingContext.cullFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCullFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCullFace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasDeleteBuffer returns true if the method "WebGLRenderingContext.deleteBuffer" exists.
func (this WebGLRenderingContext) HasDeleteBuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteBuffer(
		this.Ref(),
	)
}

// DeleteBufferFunc returns the method "WebGLRenderingContext.deleteBuffer".
func (this WebGLRenderingContext) DeleteBufferFunc() (fn js.Func[func(buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteBufferFunc(
			this.Ref(),
		),
	)
}

// DeleteBuffer calls the method "WebGLRenderingContext.deleteBuffer".
func (this WebGLRenderingContext) DeleteBuffer(buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryDeleteBuffer calls the method "WebGLRenderingContext.deleteBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteBuffer(buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasDeleteFramebuffer returns true if the method "WebGLRenderingContext.deleteFramebuffer" exists.
func (this WebGLRenderingContext) HasDeleteFramebuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteFramebuffer(
		this.Ref(),
	)
}

// DeleteFramebufferFunc returns the method "WebGLRenderingContext.deleteFramebuffer".
func (this WebGLRenderingContext) DeleteFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteFramebufferFunc(
			this.Ref(),
		),
	)
}

// DeleteFramebuffer calls the method "WebGLRenderingContext.deleteFramebuffer".
func (this WebGLRenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryDeleteFramebuffer calls the method "WebGLRenderingContext.deleteFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasDeleteProgram returns true if the method "WebGLRenderingContext.deleteProgram" exists.
func (this WebGLRenderingContext) HasDeleteProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteProgram(
		this.Ref(),
	)
}

// DeleteProgramFunc returns the method "WebGLRenderingContext.deleteProgram".
func (this WebGLRenderingContext) DeleteProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteProgramFunc(
			this.Ref(),
		),
	)
}

// DeleteProgram calls the method "WebGLRenderingContext.deleteProgram".
func (this WebGLRenderingContext) DeleteProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryDeleteProgram calls the method "WebGLRenderingContext.deleteProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasDeleteRenderbuffer returns true if the method "WebGLRenderingContext.deleteRenderbuffer" exists.
func (this WebGLRenderingContext) HasDeleteRenderbuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteRenderbuffer(
		this.Ref(),
	)
}

// DeleteRenderbufferFunc returns the method "WebGLRenderingContext.deleteRenderbuffer".
func (this WebGLRenderingContext) DeleteRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteRenderbufferFunc(
			this.Ref(),
		),
	)
}

// DeleteRenderbuffer calls the method "WebGLRenderingContext.deleteRenderbuffer".
func (this WebGLRenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryDeleteRenderbuffer calls the method "WebGLRenderingContext.deleteRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasDeleteShader returns true if the method "WebGLRenderingContext.deleteShader" exists.
func (this WebGLRenderingContext) HasDeleteShader() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteShader(
		this.Ref(),
	)
}

// DeleteShaderFunc returns the method "WebGLRenderingContext.deleteShader".
func (this WebGLRenderingContext) DeleteShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteShaderFunc(
			this.Ref(),
		),
	)
}

// DeleteShader calls the method "WebGLRenderingContext.deleteShader".
func (this WebGLRenderingContext) DeleteShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryDeleteShader calls the method "WebGLRenderingContext.deleteShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasDeleteTexture returns true if the method "WebGLRenderingContext.deleteTexture" exists.
func (this WebGLRenderingContext) HasDeleteTexture() bool {
	return js.True == bindings.HasWebGLRenderingContextDeleteTexture(
		this.Ref(),
	)
}

// DeleteTextureFunc returns the method "WebGLRenderingContext.deleteTexture".
func (this WebGLRenderingContext) DeleteTextureFunc() (fn js.Func[func(texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDeleteTextureFunc(
			this.Ref(),
		),
	)
}

// DeleteTexture calls the method "WebGLRenderingContext.deleteTexture".
func (this WebGLRenderingContext) DeleteTexture(texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryDeleteTexture calls the method "WebGLRenderingContext.deleteTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteTexture(texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasDepthFunc returns true if the method "WebGLRenderingContext.depthFunc" exists.
func (this WebGLRenderingContext) HasDepthFunc() bool {
	return js.True == bindings.HasWebGLRenderingContextDepthFunc(
		this.Ref(),
	)
}

// DepthFuncFunc returns the method "WebGLRenderingContext.depthFunc".
func (this WebGLRenderingContext) DepthFuncFunc() (fn js.Func[func(fn GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthFuncFunc(
			this.Ref(),
		),
	)
}

// DepthFunc calls the method "WebGLRenderingContext.depthFunc".
func (this WebGLRenderingContext) DepthFunc(fn GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(fn),
	)

	return
}

// TryDepthFunc calls the method "WebGLRenderingContext.depthFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDepthFunc(fn GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDepthFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
	)

	return
}

// HasDepthMask returns true if the method "WebGLRenderingContext.depthMask" exists.
func (this WebGLRenderingContext) HasDepthMask() bool {
	return js.True == bindings.HasWebGLRenderingContextDepthMask(
		this.Ref(),
	)
}

// DepthMaskFunc returns the method "WebGLRenderingContext.depthMask".
func (this WebGLRenderingContext) DepthMaskFunc() (fn js.Func[func(flag GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthMaskFunc(
			this.Ref(),
		),
	)
}

// DepthMask calls the method "WebGLRenderingContext.depthMask".
func (this WebGLRenderingContext) DepthMask(flag GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthMask(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(flag)),
	)

	return
}

// TryDepthMask calls the method "WebGLRenderingContext.depthMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDepthMask(flag GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDepthMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(flag)),
	)

	return
}

// HasDepthRange returns true if the method "WebGLRenderingContext.depthRange" exists.
func (this WebGLRenderingContext) HasDepthRange() bool {
	return js.True == bindings.HasWebGLRenderingContextDepthRange(
		this.Ref(),
	)
}

// DepthRangeFunc returns the method "WebGLRenderingContext.depthRange".
func (this WebGLRenderingContext) DepthRangeFunc() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDepthRangeFunc(
			this.Ref(),
		),
	)
}

// DepthRange calls the method "WebGLRenderingContext.depthRange".
func (this WebGLRenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthRange(
		this.Ref(), js.Pointer(&ret),
		float32(zNear),
		float32(zFar),
	)

	return
}

// TryDepthRange calls the method "WebGLRenderingContext.depthRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDepthRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(zNear),
		float32(zFar),
	)

	return
}

// HasDetachShader returns true if the method "WebGLRenderingContext.detachShader" exists.
func (this WebGLRenderingContext) HasDetachShader() bool {
	return js.True == bindings.HasWebGLRenderingContextDetachShader(
		this.Ref(),
	)
}

// DetachShaderFunc returns the method "WebGLRenderingContext.detachShader".
func (this WebGLRenderingContext) DetachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDetachShaderFunc(
			this.Ref(),
		),
	)
}

// DetachShader calls the method "WebGLRenderingContext.detachShader".
func (this WebGLRenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextDetachShader(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// TryDetachShader calls the method "WebGLRenderingContext.detachShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDetachShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasDisable returns true if the method "WebGLRenderingContext.disable" exists.
func (this WebGLRenderingContext) HasDisable() bool {
	return js.True == bindings.HasWebGLRenderingContextDisable(
		this.Ref(),
	)
}

// DisableFunc returns the method "WebGLRenderingContext.disable".
func (this WebGLRenderingContext) DisableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDisableFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "WebGLRenderingContext.disable".
func (this WebGLRenderingContext) Disable(cap GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextDisable(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryDisable calls the method "WebGLRenderingContext.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDisable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDisable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasDisableVertexAttribArray returns true if the method "WebGLRenderingContext.disableVertexAttribArray" exists.
func (this WebGLRenderingContext) HasDisableVertexAttribArray() bool {
	return js.True == bindings.HasWebGLRenderingContextDisableVertexAttribArray(
		this.Ref(),
	)
}

// DisableVertexAttribArrayFunc returns the method "WebGLRenderingContext.disableVertexAttribArray".
func (this WebGLRenderingContext) DisableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDisableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// DisableVertexAttribArray calls the method "WebGLRenderingContext.disableVertexAttribArray".
func (this WebGLRenderingContext) DisableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDisableVertexAttribArray calls the method "WebGLRenderingContext.disableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDisableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasDrawArrays returns true if the method "WebGLRenderingContext.drawArrays" exists.
func (this WebGLRenderingContext) HasDrawArrays() bool {
	return js.True == bindings.HasWebGLRenderingContextDrawArrays(
		this.Ref(),
	)
}

// DrawArraysFunc returns the method "WebGLRenderingContext.drawArrays".
func (this WebGLRenderingContext) DrawArraysFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDrawArraysFunc(
			this.Ref(),
		),
	)
}

// DrawArrays calls the method "WebGLRenderingContext.drawArrays".
func (this WebGLRenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextDrawArrays(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// TryDrawArrays calls the method "WebGLRenderingContext.drawArrays"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDrawArrays(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// HasDrawElements returns true if the method "WebGLRenderingContext.drawElements" exists.
func (this WebGLRenderingContext) HasDrawElements() bool {
	return js.True == bindings.HasWebGLRenderingContextDrawElements(
		this.Ref(),
	)
}

// DrawElementsFunc returns the method "WebGLRenderingContext.drawElements".
func (this WebGLRenderingContext) DrawElementsFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextDrawElementsFunc(
			this.Ref(),
		),
	)
}

// DrawElements calls the method "WebGLRenderingContext.drawElements".
func (this WebGLRenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGLRenderingContextDrawElements(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// TryDrawElements calls the method "WebGLRenderingContext.drawElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDrawElements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasEnable returns true if the method "WebGLRenderingContext.enable" exists.
func (this WebGLRenderingContext) HasEnable() bool {
	return js.True == bindings.HasWebGLRenderingContextEnable(
		this.Ref(),
	)
}

// EnableFunc returns the method "WebGLRenderingContext.enable".
func (this WebGLRenderingContext) EnableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextEnableFunc(
			this.Ref(),
		),
	)
}

// Enable calls the method "WebGLRenderingContext.enable".
func (this WebGLRenderingContext) Enable(cap GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextEnable(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryEnable calls the method "WebGLRenderingContext.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryEnable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextEnable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasEnableVertexAttribArray returns true if the method "WebGLRenderingContext.enableVertexAttribArray" exists.
func (this WebGLRenderingContext) HasEnableVertexAttribArray() bool {
	return js.True == bindings.HasWebGLRenderingContextEnableVertexAttribArray(
		this.Ref(),
	)
}

// EnableVertexAttribArrayFunc returns the method "WebGLRenderingContext.enableVertexAttribArray".
func (this WebGLRenderingContext) EnableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextEnableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// EnableVertexAttribArray calls the method "WebGLRenderingContext.enableVertexAttribArray".
func (this WebGLRenderingContext) EnableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryEnableVertexAttribArray calls the method "WebGLRenderingContext.enableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryEnableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFinish returns true if the method "WebGLRenderingContext.finish" exists.
func (this WebGLRenderingContext) HasFinish() bool {
	return js.True == bindings.HasWebGLRenderingContextFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "WebGLRenderingContext.finish".
func (this WebGLRenderingContext) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "WebGLRenderingContext.finish".
func (this WebGLRenderingContext) Finish() (ret js.Void) {
	bindings.CallWebGLRenderingContextFinish(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "WebGLRenderingContext.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFlush returns true if the method "WebGLRenderingContext.flush" exists.
func (this WebGLRenderingContext) HasFlush() bool {
	return js.True == bindings.HasWebGLRenderingContextFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "WebGLRenderingContext.flush".
func (this WebGLRenderingContext) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "WebGLRenderingContext.flush".
func (this WebGLRenderingContext) Flush() (ret js.Void) {
	bindings.CallWebGLRenderingContextFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "WebGLRenderingContext.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFramebufferRenderbuffer returns true if the method "WebGLRenderingContext.framebufferRenderbuffer" exists.
func (this WebGLRenderingContext) HasFramebufferRenderbuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextFramebufferRenderbuffer(
		this.Ref(),
	)
}

// FramebufferRenderbufferFunc returns the method "WebGLRenderingContext.framebufferRenderbuffer".
func (this WebGLRenderingContext) FramebufferRenderbufferFunc() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFramebufferRenderbufferFunc(
			this.Ref(),
		),
	)
}

// FramebufferRenderbuffer calls the method "WebGLRenderingContext.framebufferRenderbuffer".
func (this WebGLRenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// TryFramebufferRenderbuffer calls the method "WebGLRenderingContext.framebufferRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// HasFramebufferTexture2D returns true if the method "WebGLRenderingContext.framebufferTexture2D" exists.
func (this WebGLRenderingContext) HasFramebufferTexture2D() bool {
	return js.True == bindings.HasWebGLRenderingContextFramebufferTexture2D(
		this.Ref(),
	)
}

// FramebufferTexture2DFunc returns the method "WebGLRenderingContext.framebufferTexture2D".
func (this WebGLRenderingContext) FramebufferTexture2DFunc() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFramebufferTexture2DFunc(
			this.Ref(),
		),
	)
}

// FramebufferTexture2D calls the method "WebGLRenderingContext.framebufferTexture2D".
func (this WebGLRenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// TryFramebufferTexture2D calls the method "WebGLRenderingContext.framebufferTexture2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// HasFrontFace returns true if the method "WebGLRenderingContext.frontFace" exists.
func (this WebGLRenderingContext) HasFrontFace() bool {
	return js.True == bindings.HasWebGLRenderingContextFrontFace(
		this.Ref(),
	)
}

// FrontFaceFunc returns the method "WebGLRenderingContext.frontFace".
func (this WebGLRenderingContext) FrontFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextFrontFaceFunc(
			this.Ref(),
		),
	)
}

// FrontFace calls the method "WebGLRenderingContext.frontFace".
func (this WebGLRenderingContext) FrontFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextFrontFace(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryFrontFace calls the method "WebGLRenderingContext.frontFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFrontFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFrontFace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasGenerateMipmap returns true if the method "WebGLRenderingContext.generateMipmap" exists.
func (this WebGLRenderingContext) HasGenerateMipmap() bool {
	return js.True == bindings.HasWebGLRenderingContextGenerateMipmap(
		this.Ref(),
	)
}

// GenerateMipmapFunc returns the method "WebGLRenderingContext.generateMipmap".
func (this WebGLRenderingContext) GenerateMipmapFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGenerateMipmapFunc(
			this.Ref(),
		),
	)
}

// GenerateMipmap calls the method "WebGLRenderingContext.generateMipmap".
func (this WebGLRenderingContext) GenerateMipmap(target GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryGenerateMipmap calls the method "WebGLRenderingContext.generateMipmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGenerateMipmap(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasGetActiveAttrib returns true if the method "WebGLRenderingContext.getActiveAttrib" exists.
func (this WebGLRenderingContext) HasGetActiveAttrib() bool {
	return js.True == bindings.HasWebGLRenderingContextGetActiveAttrib(
		this.Ref(),
	)
}

// GetActiveAttribFunc returns the method "WebGLRenderingContext.getActiveAttrib".
func (this WebGLRenderingContext) GetActiveAttribFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetActiveAttribFunc(
			this.Ref(),
		),
	)
}

// GetActiveAttrib calls the method "WebGLRenderingContext.getActiveAttrib".
func (this WebGLRenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGLRenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
	)

	return
}

// TryGetActiveAttrib calls the method "WebGLRenderingContext.getActiveAttrib"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasGetActiveUniform returns true if the method "WebGLRenderingContext.getActiveUniform" exists.
func (this WebGLRenderingContext) HasGetActiveUniform() bool {
	return js.True == bindings.HasWebGLRenderingContextGetActiveUniform(
		this.Ref(),
	)
}

// GetActiveUniformFunc returns the method "WebGLRenderingContext.getActiveUniform".
func (this WebGLRenderingContext) GetActiveUniformFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetActiveUniformFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniform calls the method "WebGLRenderingContext.getActiveUniform".
func (this WebGLRenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGLRenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
	)

	return
}

// TryGetActiveUniform calls the method "WebGLRenderingContext.getActiveUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasGetAttachedShaders returns true if the method "WebGLRenderingContext.getAttachedShaders" exists.
func (this WebGLRenderingContext) HasGetAttachedShaders() bool {
	return js.True == bindings.HasWebGLRenderingContextGetAttachedShaders(
		this.Ref(),
	)
}

// GetAttachedShadersFunc returns the method "WebGLRenderingContext.getAttachedShaders".
func (this WebGLRenderingContext) GetAttachedShadersFunc() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetAttachedShadersFunc(
			this.Ref(),
		),
	)
}

// GetAttachedShaders calls the method "WebGLRenderingContext.getAttachedShaders".
func (this WebGLRenderingContext) GetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader]) {
	bindings.CallWebGLRenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetAttachedShaders calls the method "WebGLRenderingContext.getAttachedShaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasGetAttribLocation returns true if the method "WebGLRenderingContext.getAttribLocation" exists.
func (this WebGLRenderingContext) HasGetAttribLocation() bool {
	return js.True == bindings.HasWebGLRenderingContextGetAttribLocation(
		this.Ref(),
	)
}

// GetAttribLocationFunc returns the method "WebGLRenderingContext.getAttribLocation".
func (this WebGLRenderingContext) GetAttribLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetAttribLocationFunc(
			this.Ref(),
		),
	)
}

// GetAttribLocation calls the method "WebGLRenderingContext.getAttribLocation".
func (this WebGLRenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGLRenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		name.Ref(),
	)

	return
}

// TryGetAttribLocation calls the method "WebGLRenderingContext.getAttribLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetAttribLocation(program WebGLProgram, name js.String) (ret GLint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasGetBufferParameter returns true if the method "WebGLRenderingContext.getBufferParameter" exists.
func (this WebGLRenderingContext) HasGetBufferParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetBufferParameter(
		this.Ref(),
	)
}

// GetBufferParameterFunc returns the method "WebGLRenderingContext.getBufferParameter".
func (this WebGLRenderingContext) GetBufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetBufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetBufferParameter calls the method "WebGLRenderingContext.getBufferParameter".
func (this WebGLRenderingContext) GetBufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetBufferParameter calls the method "WebGLRenderingContext.getBufferParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetBufferParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetParameter returns true if the method "WebGLRenderingContext.getParameter" exists.
func (this WebGLRenderingContext) HasGetParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetParameter(
		this.Ref(),
	)
}

// GetParameterFunc returns the method "WebGLRenderingContext.getParameter".
func (this WebGLRenderingContext) GetParameterFunc() (fn js.Func[func(pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetParameterFunc(
			this.Ref(),
		),
	)
}

// GetParameter calls the method "WebGLRenderingContext.getParameter".
func (this WebGLRenderingContext) GetParameter(pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(pname),
	)

	return
}

// TryGetParameter calls the method "WebGLRenderingContext.getParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetParameter(pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
	)

	return
}

// HasGetError returns true if the method "WebGLRenderingContext.getError" exists.
func (this WebGLRenderingContext) HasGetError() bool {
	return js.True == bindings.HasWebGLRenderingContextGetError(
		this.Ref(),
	)
}

// GetErrorFunc returns the method "WebGLRenderingContext.getError".
func (this WebGLRenderingContext) GetErrorFunc() (fn js.Func[func() GLenum]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetErrorFunc(
			this.Ref(),
		),
	)
}

// GetError calls the method "WebGLRenderingContext.getError".
func (this WebGLRenderingContext) GetError() (ret GLenum) {
	bindings.CallWebGLRenderingContextGetError(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetError calls the method "WebGLRenderingContext.getError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetError() (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetFramebufferAttachmentParameter returns true if the method "WebGLRenderingContext.getFramebufferAttachmentParameter" exists.
func (this WebGLRenderingContext) HasGetFramebufferAttachmentParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.Ref(),
	)
}

// GetFramebufferAttachmentParameterFunc returns the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
func (this WebGLRenderingContext) GetFramebufferAttachmentParameterFunc() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetFramebufferAttachmentParameterFunc(
			this.Ref(),
		),
	)
}

// GetFramebufferAttachmentParameter calls the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
func (this WebGLRenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// TryGetFramebufferAttachmentParameter calls the method "WebGLRenderingContext.getFramebufferAttachmentParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// HasGetProgramParameter returns true if the method "WebGLRenderingContext.getProgramParameter" exists.
func (this WebGLRenderingContext) HasGetProgramParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetProgramParameter(
		this.Ref(),
	)
}

// GetProgramParameterFunc returns the method "WebGLRenderingContext.getProgramParameter".
func (this WebGLRenderingContext) GetProgramParameterFunc() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetProgramParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramParameter calls the method "WebGLRenderingContext.getProgramParameter".
func (this WebGLRenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(pname),
	)

	return
}

// TryGetProgramParameter calls the method "WebGLRenderingContext.getProgramParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(pname),
	)

	return
}

// HasGetProgramInfoLog returns true if the method "WebGLRenderingContext.getProgramInfoLog" exists.
func (this WebGLRenderingContext) HasGetProgramInfoLog() bool {
	return js.True == bindings.HasWebGLRenderingContextGetProgramInfoLog(
		this.Ref(),
	)
}

// GetProgramInfoLogFunc returns the method "WebGLRenderingContext.getProgramInfoLog".
func (this WebGLRenderingContext) GetProgramInfoLogFunc() (fn js.Func[func(program WebGLProgram) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetProgramInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetProgramInfoLog calls the method "WebGLRenderingContext.getProgramInfoLog".
func (this WebGLRenderingContext) GetProgramInfoLog(program WebGLProgram) (ret js.String) {
	bindings.CallWebGLRenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetProgramInfoLog calls the method "WebGLRenderingContext.getProgramInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetProgramInfoLog(program WebGLProgram) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasGetRenderbufferParameter returns true if the method "WebGLRenderingContext.getRenderbufferParameter" exists.
func (this WebGLRenderingContext) HasGetRenderbufferParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetRenderbufferParameter(
		this.Ref(),
	)
}

// GetRenderbufferParameterFunc returns the method "WebGLRenderingContext.getRenderbufferParameter".
func (this WebGLRenderingContext) GetRenderbufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetRenderbufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetRenderbufferParameter calls the method "WebGLRenderingContext.getRenderbufferParameter".
func (this WebGLRenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetRenderbufferParameter calls the method "WebGLRenderingContext.getRenderbufferParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetShaderParameter returns true if the method "WebGLRenderingContext.getShaderParameter" exists.
func (this WebGLRenderingContext) HasGetShaderParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetShaderParameter(
		this.Ref(),
	)
}

// GetShaderParameterFunc returns the method "WebGLRenderingContext.getShaderParameter".
func (this WebGLRenderingContext) GetShaderParameterFunc() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderParameter calls the method "WebGLRenderingContext.getShaderParameter".
func (this WebGLRenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// TryGetShaderParameter calls the method "WebGLRenderingContext.getShaderParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// HasGetShaderPrecisionFormat returns true if the method "WebGLRenderingContext.getShaderPrecisionFormat" exists.
func (this WebGLRenderingContext) HasGetShaderPrecisionFormat() bool {
	return js.True == bindings.HasWebGLRenderingContextGetShaderPrecisionFormat(
		this.Ref(),
	)
}

// GetShaderPrecisionFormatFunc returns the method "WebGLRenderingContext.getShaderPrecisionFormat".
func (this WebGLRenderingContext) GetShaderPrecisionFormatFunc() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderPrecisionFormatFunc(
			this.Ref(),
		),
	)
}

// GetShaderPrecisionFormat calls the method "WebGLRenderingContext.getShaderPrecisionFormat".
func (this WebGLRenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat) {
	bindings.CallWebGLRenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&ret),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// TryGetShaderPrecisionFormat calls the method "WebGLRenderingContext.getShaderPrecisionFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// HasGetShaderInfoLog returns true if the method "WebGLRenderingContext.getShaderInfoLog" exists.
func (this WebGLRenderingContext) HasGetShaderInfoLog() bool {
	return js.True == bindings.HasWebGLRenderingContextGetShaderInfoLog(
		this.Ref(),
	)
}

// GetShaderInfoLogFunc returns the method "WebGLRenderingContext.getShaderInfoLog".
func (this WebGLRenderingContext) GetShaderInfoLogFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetShaderInfoLog calls the method "WebGLRenderingContext.getShaderInfoLog".
func (this WebGLRenderingContext) GetShaderInfoLog(shader WebGLShader) (ret js.String) {
	bindings.CallWebGLRenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderInfoLog calls the method "WebGLRenderingContext.getShaderInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderInfoLog(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasGetShaderSource returns true if the method "WebGLRenderingContext.getShaderSource" exists.
func (this WebGLRenderingContext) HasGetShaderSource() bool {
	return js.True == bindings.HasWebGLRenderingContextGetShaderSource(
		this.Ref(),
	)
}

// GetShaderSourceFunc returns the method "WebGLRenderingContext.getShaderSource".
func (this WebGLRenderingContext) GetShaderSourceFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetShaderSourceFunc(
			this.Ref(),
		),
	)
}

// GetShaderSource calls the method "WebGLRenderingContext.getShaderSource".
func (this WebGLRenderingContext) GetShaderSource(shader WebGLShader) (ret js.String) {
	bindings.CallWebGLRenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderSource calls the method "WebGLRenderingContext.getShaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderSource(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasGetTexParameter returns true if the method "WebGLRenderingContext.getTexParameter" exists.
func (this WebGLRenderingContext) HasGetTexParameter() bool {
	return js.True == bindings.HasWebGLRenderingContextGetTexParameter(
		this.Ref(),
	)
}

// GetTexParameterFunc returns the method "WebGLRenderingContext.getTexParameter".
func (this WebGLRenderingContext) GetTexParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetTexParameterFunc(
			this.Ref(),
		),
	)
}

// GetTexParameter calls the method "WebGLRenderingContext.getTexParameter".
func (this WebGLRenderingContext) GetTexParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetTexParameter calls the method "WebGLRenderingContext.getTexParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetTexParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetUniform returns true if the method "WebGLRenderingContext.getUniform" exists.
func (this WebGLRenderingContext) HasGetUniform() bool {
	return js.True == bindings.HasWebGLRenderingContextGetUniform(
		this.Ref(),
	)
}

// GetUniformFunc returns the method "WebGLRenderingContext.getUniform".
func (this WebGLRenderingContext) GetUniformFunc() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetUniformFunc(
			this.Ref(),
		),
	)
}

// GetUniform calls the method "WebGLRenderingContext.getUniform".
func (this WebGLRenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetUniform(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		location.Ref(),
	)

	return
}

// TryGetUniform calls the method "WebGLRenderingContext.getUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetUniform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		location.Ref(),
	)

	return
}

// HasGetUniformLocation returns true if the method "WebGLRenderingContext.getUniformLocation" exists.
func (this WebGLRenderingContext) HasGetUniformLocation() bool {
	return js.True == bindings.HasWebGLRenderingContextGetUniformLocation(
		this.Ref(),
	)
}

// GetUniformLocationFunc returns the method "WebGLRenderingContext.getUniformLocation".
func (this WebGLRenderingContext) GetUniformLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetUniformLocationFunc(
			this.Ref(),
		),
	)
}

// GetUniformLocation calls the method "WebGLRenderingContext.getUniformLocation".
func (this WebGLRenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation) {
	bindings.CallWebGLRenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		name.Ref(),
	)

	return
}

// TryGetUniformLocation calls the method "WebGLRenderingContext.getUniformLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasGetVertexAttrib returns true if the method "WebGLRenderingContext.getVertexAttrib" exists.
func (this WebGLRenderingContext) HasGetVertexAttrib() bool {
	return js.True == bindings.HasWebGLRenderingContextGetVertexAttrib(
		this.Ref(),
	)
}

// GetVertexAttribFunc returns the method "WebGLRenderingContext.getVertexAttrib".
func (this WebGLRenderingContext) GetVertexAttribFunc() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetVertexAttribFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttrib calls the method "WebGLRenderingContext.getVertexAttrib".
func (this WebGLRenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(pname),
	)

	return
}

// TryGetVertexAttrib calls the method "WebGLRenderingContext.getVertexAttrib"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetVertexAttrib(index GLuint, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasGetVertexAttribOffset returns true if the method "WebGLRenderingContext.getVertexAttribOffset" exists.
func (this WebGLRenderingContext) HasGetVertexAttribOffset() bool {
	return js.True == bindings.HasWebGLRenderingContextGetVertexAttribOffset(
		this.Ref(),
	)
}

// GetVertexAttribOffsetFunc returns the method "WebGLRenderingContext.getVertexAttribOffset".
func (this WebGLRenderingContext) GetVertexAttribOffsetFunc() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextGetVertexAttribOffsetFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttribOffset calls the method "WebGLRenderingContext.getVertexAttribOffset".
func (this WebGLRenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr) {
	bindings.CallWebGLRenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(pname),
	)

	return
}

// TryGetVertexAttribOffset calls the method "WebGLRenderingContext.getVertexAttribOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasHint returns true if the method "WebGLRenderingContext.hint" exists.
func (this WebGLRenderingContext) HasHint() bool {
	return js.True == bindings.HasWebGLRenderingContextHint(
		this.Ref(),
	)
}

// HintFunc returns the method "WebGLRenderingContext.hint".
func (this WebGLRenderingContext) HintFunc() (fn js.Func[func(target GLenum, mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextHintFunc(
			this.Ref(),
		),
	)
}

// Hint calls the method "WebGLRenderingContext.hint".
func (this WebGLRenderingContext) Hint(target GLenum, mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextHint(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(mode),
	)

	return
}

// TryHint calls the method "WebGLRenderingContext.hint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryHint(target GLenum, mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextHint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(mode),
	)

	return
}

// HasIsBuffer returns true if the method "WebGLRenderingContext.isBuffer" exists.
func (this WebGLRenderingContext) HasIsBuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextIsBuffer(
		this.Ref(),
	)
}

// IsBufferFunc returns the method "WebGLRenderingContext.isBuffer".
func (this WebGLRenderingContext) IsBufferFunc() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsBufferFunc(
			this.Ref(),
		),
	)
}

// IsBuffer calls the method "WebGLRenderingContext.isBuffer".
func (this WebGLRenderingContext) IsBuffer(buffer WebGLBuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryIsBuffer calls the method "WebGLRenderingContext.isBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsBuffer(buffer WebGLBuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasIsEnabled returns true if the method "WebGLRenderingContext.isEnabled" exists.
func (this WebGLRenderingContext) HasIsEnabled() bool {
	return js.True == bindings.HasWebGLRenderingContextIsEnabled(
		this.Ref(),
	)
}

// IsEnabledFunc returns the method "WebGLRenderingContext.isEnabled".
func (this WebGLRenderingContext) IsEnabledFunc() (fn js.Func[func(cap GLenum) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsEnabledFunc(
			this.Ref(),
		),
	)
}

// IsEnabled calls the method "WebGLRenderingContext.isEnabled".
func (this WebGLRenderingContext) IsEnabled(cap GLenum) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsEnabled(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryIsEnabled calls the method "WebGLRenderingContext.isEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsEnabled(cap GLenum) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasIsFramebuffer returns true if the method "WebGLRenderingContext.isFramebuffer" exists.
func (this WebGLRenderingContext) HasIsFramebuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextIsFramebuffer(
		this.Ref(),
	)
}

// IsFramebufferFunc returns the method "WebGLRenderingContext.isFramebuffer".
func (this WebGLRenderingContext) IsFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsFramebufferFunc(
			this.Ref(),
		),
	)
}

// IsFramebuffer calls the method "WebGLRenderingContext.isFramebuffer".
func (this WebGLRenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryIsFramebuffer calls the method "WebGLRenderingContext.isFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasIsProgram returns true if the method "WebGLRenderingContext.isProgram" exists.
func (this WebGLRenderingContext) HasIsProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextIsProgram(
		this.Ref(),
	)
}

// IsProgramFunc returns the method "WebGLRenderingContext.isProgram".
func (this WebGLRenderingContext) IsProgramFunc() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsProgramFunc(
			this.Ref(),
		),
	)
}

// IsProgram calls the method "WebGLRenderingContext.isProgram".
func (this WebGLRenderingContext) IsProgram(program WebGLProgram) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryIsProgram calls the method "WebGLRenderingContext.isProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsProgram(program WebGLProgram) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasIsRenderbuffer returns true if the method "WebGLRenderingContext.isRenderbuffer" exists.
func (this WebGLRenderingContext) HasIsRenderbuffer() bool {
	return js.True == bindings.HasWebGLRenderingContextIsRenderbuffer(
		this.Ref(),
	)
}

// IsRenderbufferFunc returns the method "WebGLRenderingContext.isRenderbuffer".
func (this WebGLRenderingContext) IsRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsRenderbufferFunc(
			this.Ref(),
		),
	)
}

// IsRenderbuffer calls the method "WebGLRenderingContext.isRenderbuffer".
func (this WebGLRenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryIsRenderbuffer calls the method "WebGLRenderingContext.isRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasIsShader returns true if the method "WebGLRenderingContext.isShader" exists.
func (this WebGLRenderingContext) HasIsShader() bool {
	return js.True == bindings.HasWebGLRenderingContextIsShader(
		this.Ref(),
	)
}

// IsShaderFunc returns the method "WebGLRenderingContext.isShader".
func (this WebGLRenderingContext) IsShaderFunc() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsShaderFunc(
			this.Ref(),
		),
	)
}

// IsShader calls the method "WebGLRenderingContext.isShader".
func (this WebGLRenderingContext) IsShader(shader WebGLShader) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryIsShader calls the method "WebGLRenderingContext.isShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsShader(shader WebGLShader) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasIsTexture returns true if the method "WebGLRenderingContext.isTexture" exists.
func (this WebGLRenderingContext) HasIsTexture() bool {
	return js.True == bindings.HasWebGLRenderingContextIsTexture(
		this.Ref(),
	)
}

// IsTextureFunc returns the method "WebGLRenderingContext.isTexture".
func (this WebGLRenderingContext) IsTextureFunc() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextIsTextureFunc(
			this.Ref(),
		),
	)
}

// IsTexture calls the method "WebGLRenderingContext.isTexture".
func (this WebGLRenderingContext) IsTexture(texture WebGLTexture) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsTexture(
		this.Ref(), js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryIsTexture calls the method "WebGLRenderingContext.isTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsTexture(texture WebGLTexture) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasLineWidth returns true if the method "WebGLRenderingContext.lineWidth" exists.
func (this WebGLRenderingContext) HasLineWidth() bool {
	return js.True == bindings.HasWebGLRenderingContextLineWidth(
		this.Ref(),
	)
}

// LineWidthFunc returns the method "WebGLRenderingContext.lineWidth".
func (this WebGLRenderingContext) LineWidthFunc() (fn js.Func[func(width GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextLineWidthFunc(
			this.Ref(),
		),
	)
}

// LineWidth calls the method "WebGLRenderingContext.lineWidth".
func (this WebGLRenderingContext) LineWidth(width GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextLineWidth(
		this.Ref(), js.Pointer(&ret),
		float32(width),
	)

	return
}

// TryLineWidth calls the method "WebGLRenderingContext.lineWidth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryLineWidth(width GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextLineWidth(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(width),
	)

	return
}

// HasLinkProgram returns true if the method "WebGLRenderingContext.linkProgram" exists.
func (this WebGLRenderingContext) HasLinkProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextLinkProgram(
		this.Ref(),
	)
}

// LinkProgramFunc returns the method "WebGLRenderingContext.linkProgram".
func (this WebGLRenderingContext) LinkProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextLinkProgramFunc(
			this.Ref(),
		),
	)
}

// LinkProgram calls the method "WebGLRenderingContext.linkProgram".
func (this WebGLRenderingContext) LinkProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextLinkProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryLinkProgram calls the method "WebGLRenderingContext.linkProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryLinkProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextLinkProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasPixelStorei returns true if the method "WebGLRenderingContext.pixelStorei" exists.
func (this WebGLRenderingContext) HasPixelStorei() bool {
	return js.True == bindings.HasWebGLRenderingContextPixelStorei(
		this.Ref(),
	)
}

// PixelStoreiFunc returns the method "WebGLRenderingContext.pixelStorei".
func (this WebGLRenderingContext) PixelStoreiFunc() (fn js.Func[func(pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextPixelStoreiFunc(
			this.Ref(),
		),
	)
}

// PixelStorei calls the method "WebGLRenderingContext.pixelStorei".
func (this WebGLRenderingContext) PixelStorei(pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextPixelStorei(
		this.Ref(), js.Pointer(&ret),
		uint32(pname),
		int32(param),
	)

	return
}

// TryPixelStorei calls the method "WebGLRenderingContext.pixelStorei"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryPixelStorei(pname GLenum, param GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextPixelStorei(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
		int32(param),
	)

	return
}

// HasPolygonOffset returns true if the method "WebGLRenderingContext.polygonOffset" exists.
func (this WebGLRenderingContext) HasPolygonOffset() bool {
	return js.True == bindings.HasWebGLRenderingContextPolygonOffset(
		this.Ref(),
	)
}

// PolygonOffsetFunc returns the method "WebGLRenderingContext.polygonOffset".
func (this WebGLRenderingContext) PolygonOffsetFunc() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextPolygonOffsetFunc(
			this.Ref(),
		),
	)
}

// PolygonOffset calls the method "WebGLRenderingContext.polygonOffset".
func (this WebGLRenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&ret),
		float32(factor),
		float32(units),
	)

	return
}

// TryPolygonOffset calls the method "WebGLRenderingContext.polygonOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryPolygonOffset(factor GLfloat, units GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(factor),
		float32(units),
	)

	return
}

// HasRenderbufferStorage returns true if the method "WebGLRenderingContext.renderbufferStorage" exists.
func (this WebGLRenderingContext) HasRenderbufferStorage() bool {
	return js.True == bindings.HasWebGLRenderingContextRenderbufferStorage(
		this.Ref(),
	)
}

// RenderbufferStorageFunc returns the method "WebGLRenderingContext.renderbufferStorage".
func (this WebGLRenderingContext) RenderbufferStorageFunc() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextRenderbufferStorageFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorage calls the method "WebGLRenderingContext.renderbufferStorage".
func (this WebGLRenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// TryRenderbufferStorage calls the method "WebGLRenderingContext.renderbufferStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryRenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasSampleCoverage returns true if the method "WebGLRenderingContext.sampleCoverage" exists.
func (this WebGLRenderingContext) HasSampleCoverage() bool {
	return js.True == bindings.HasWebGLRenderingContextSampleCoverage(
		this.Ref(),
	)
}

// SampleCoverageFunc returns the method "WebGLRenderingContext.sampleCoverage".
func (this WebGLRenderingContext) SampleCoverageFunc() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextSampleCoverageFunc(
			this.Ref(),
		),
	)
}

// SampleCoverage calls the method "WebGLRenderingContext.sampleCoverage".
func (this WebGLRenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&ret),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// TrySampleCoverage calls the method "WebGLRenderingContext.sampleCoverage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TrySampleCoverage(value GLclampf, invert GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// HasScissor returns true if the method "WebGLRenderingContext.scissor" exists.
func (this WebGLRenderingContext) HasScissor() bool {
	return js.True == bindings.HasWebGLRenderingContextScissor(
		this.Ref(),
	)
}

// ScissorFunc returns the method "WebGLRenderingContext.scissor".
func (this WebGLRenderingContext) ScissorFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextScissorFunc(
			this.Ref(),
		),
	)
}

// Scissor calls the method "WebGLRenderingContext.scissor".
func (this WebGLRenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextScissor(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryScissor calls the method "WebGLRenderingContext.scissor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryScissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextScissor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasShaderSource returns true if the method "WebGLRenderingContext.shaderSource" exists.
func (this WebGLRenderingContext) HasShaderSource() bool {
	return js.True == bindings.HasWebGLRenderingContextShaderSource(
		this.Ref(),
	)
}

// ShaderSourceFunc returns the method "WebGLRenderingContext.shaderSource".
func (this WebGLRenderingContext) ShaderSourceFunc() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextShaderSourceFunc(
			this.Ref(),
		),
	)
}

// ShaderSource calls the method "WebGLRenderingContext.shaderSource".
func (this WebGLRenderingContext) ShaderSource(shader WebGLShader, source js.String) (ret js.Void) {
	bindings.CallWebGLRenderingContextShaderSource(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// TryShaderSource calls the method "WebGLRenderingContext.shaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryShaderSource(shader WebGLShader, source js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextShaderSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// HasStencilFunc returns true if the method "WebGLRenderingContext.stencilFunc" exists.
func (this WebGLRenderingContext) HasStencilFunc() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilFunc(
		this.Ref(),
	)
}

// StencilFuncFunc returns the method "WebGLRenderingContext.stencilFunc".
func (this WebGLRenderingContext) StencilFuncFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilFuncFunc(
			this.Ref(),
		),
	)
}

// StencilFunc calls the method "WebGLRenderingContext.stencilFunc".
func (this WebGLRenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// TryStencilFunc calls the method "WebGLRenderingContext.stencilFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasStencilFuncSeparate returns true if the method "WebGLRenderingContext.stencilFuncSeparate" exists.
func (this WebGLRenderingContext) HasStencilFuncSeparate() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilFuncSeparate(
		this.Ref(),
	)
}

// StencilFuncSeparateFunc returns the method "WebGLRenderingContext.stencilFuncSeparate".
func (this WebGLRenderingContext) StencilFuncSeparateFunc() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilFuncSeparate calls the method "WebGLRenderingContext.stencilFuncSeparate".
func (this WebGLRenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// TryStencilFuncSeparate calls the method "WebGLRenderingContext.stencilFuncSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasStencilMask returns true if the method "WebGLRenderingContext.stencilMask" exists.
func (this WebGLRenderingContext) HasStencilMask() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilMask(
		this.Ref(),
	)
}

// StencilMaskFunc returns the method "WebGLRenderingContext.stencilMask".
func (this WebGLRenderingContext) StencilMaskFunc() (fn js.Func[func(mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilMaskFunc(
			this.Ref(),
		),
	)
}

// StencilMask calls the method "WebGLRenderingContext.stencilMask".
func (this WebGLRenderingContext) StencilMask(mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilMask(
		this.Ref(), js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryStencilMask calls the method "WebGLRenderingContext.stencilMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilMask(mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasStencilMaskSeparate returns true if the method "WebGLRenderingContext.stencilMaskSeparate" exists.
func (this WebGLRenderingContext) HasStencilMaskSeparate() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilMaskSeparate(
		this.Ref(),
	)
}

// StencilMaskSeparateFunc returns the method "WebGLRenderingContext.stencilMaskSeparate".
func (this WebGLRenderingContext) StencilMaskSeparateFunc() (fn js.Func[func(face GLenum, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilMaskSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilMaskSeparate calls the method "WebGLRenderingContext.stencilMaskSeparate".
func (this WebGLRenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(mask),
	)

	return
}

// TryStencilMaskSeparate calls the method "WebGLRenderingContext.stencilMaskSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(mask),
	)

	return
}

// HasStencilOp returns true if the method "WebGLRenderingContext.stencilOp" exists.
func (this WebGLRenderingContext) HasStencilOp() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilOp(
		this.Ref(),
	)
}

// StencilOpFunc returns the method "WebGLRenderingContext.stencilOp".
func (this WebGLRenderingContext) StencilOpFunc() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilOpFunc(
			this.Ref(),
		),
	)
}

// StencilOp calls the method "WebGLRenderingContext.stencilOp".
func (this WebGLRenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilOp(
		this.Ref(), js.Pointer(&ret),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// TryStencilOp calls the method "WebGLRenderingContext.stencilOp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilOp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasStencilOpSeparate returns true if the method "WebGLRenderingContext.stencilOpSeparate" exists.
func (this WebGLRenderingContext) HasStencilOpSeparate() bool {
	return js.True == bindings.HasWebGLRenderingContextStencilOpSeparate(
		this.Ref(),
	)
}

// StencilOpSeparateFunc returns the method "WebGLRenderingContext.stencilOpSeparate".
func (this WebGLRenderingContext) StencilOpSeparateFunc() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextStencilOpSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilOpSeparate calls the method "WebGLRenderingContext.stencilOpSeparate".
func (this WebGLRenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// TryStencilOpSeparate calls the method "WebGLRenderingContext.stencilOpSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasTexParameterf returns true if the method "WebGLRenderingContext.texParameterf" exists.
func (this WebGLRenderingContext) HasTexParameterf() bool {
	return js.True == bindings.HasWebGLRenderingContextTexParameterf(
		this.Ref(),
	)
}

// TexParameterfFunc returns the method "WebGLRenderingContext.texParameterf".
func (this WebGLRenderingContext) TexParameterfFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexParameterfFunc(
			this.Ref(),
		),
	)
}

// TexParameterf calls the method "WebGLRenderingContext.texParameterf".
func (this WebGLRenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexParameterf(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// TryTexParameterf calls the method "WebGLRenderingContext.texParameterf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexParameterf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// HasTexParameteri returns true if the method "WebGLRenderingContext.texParameteri" exists.
func (this WebGLRenderingContext) HasTexParameteri() bool {
	return js.True == bindings.HasWebGLRenderingContextTexParameteri(
		this.Ref(),
	)
}

// TexParameteriFunc returns the method "WebGLRenderingContext.texParameteri".
func (this WebGLRenderingContext) TexParameteriFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexParameteriFunc(
			this.Ref(),
		),
	)
}

// TexParameteri calls the method "WebGLRenderingContext.texParameteri".
func (this WebGLRenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexParameteri(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// TryTexParameteri calls the method "WebGLRenderingContext.texParameteri"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexParameteri(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// HasUniform1f returns true if the method "WebGLRenderingContext.uniform1f" exists.
func (this WebGLRenderingContext) HasUniform1f() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform1f(
		this.Ref(),
	)
}

// Uniform1fFunc returns the method "WebGLRenderingContext.uniform1f".
func (this WebGLRenderingContext) Uniform1fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1fFunc(
			this.Ref(),
		),
	)
}

// Uniform1f calls the method "WebGLRenderingContext.uniform1f".
func (this WebGLRenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
	)

	return
}

// TryUniform1f calls the method "WebGLRenderingContext.uniform1f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform1f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
	)

	return
}

// HasUniform2f returns true if the method "WebGLRenderingContext.uniform2f" exists.
func (this WebGLRenderingContext) HasUniform2f() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform2f(
		this.Ref(),
	)
}

// Uniform2fFunc returns the method "WebGLRenderingContext.uniform2f".
func (this WebGLRenderingContext) Uniform2fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2fFunc(
			this.Ref(),
		),
	)
}

// Uniform2f calls the method "WebGLRenderingContext.uniform2f".
func (this WebGLRenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// TryUniform2f calls the method "WebGLRenderingContext.uniform2f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform2f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// HasUniform3f returns true if the method "WebGLRenderingContext.uniform3f" exists.
func (this WebGLRenderingContext) HasUniform3f() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform3f(
		this.Ref(),
	)
}

// Uniform3fFunc returns the method "WebGLRenderingContext.uniform3f".
func (this WebGLRenderingContext) Uniform3fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3fFunc(
			this.Ref(),
		),
	)
}

// Uniform3f calls the method "WebGLRenderingContext.uniform3f".
func (this WebGLRenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TryUniform3f calls the method "WebGLRenderingContext.uniform3f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform3f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasUniform4f returns true if the method "WebGLRenderingContext.uniform4f" exists.
func (this WebGLRenderingContext) HasUniform4f() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform4f(
		this.Ref(),
	)
}

// Uniform4fFunc returns the method "WebGLRenderingContext.uniform4f".
func (this WebGLRenderingContext) Uniform4fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4fFunc(
			this.Ref(),
		),
	)
}

// Uniform4f calls the method "WebGLRenderingContext.uniform4f".
func (this WebGLRenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// TryUniform4f calls the method "WebGLRenderingContext.uniform4f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform4f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasUniform1i returns true if the method "WebGLRenderingContext.uniform1i" exists.
func (this WebGLRenderingContext) HasUniform1i() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform1i(
		this.Ref(),
	)
}

// Uniform1iFunc returns the method "WebGLRenderingContext.uniform1i".
func (this WebGLRenderingContext) Uniform1iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1iFunc(
			this.Ref(),
		),
	)
}

// Uniform1i calls the method "WebGLRenderingContext.uniform1i".
func (this WebGLRenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
	)

	return
}

// TryUniform1i calls the method "WebGLRenderingContext.uniform1i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform1i(location WebGLUniformLocation, x GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform1i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
	)

	return
}

// HasUniform2i returns true if the method "WebGLRenderingContext.uniform2i" exists.
func (this WebGLRenderingContext) HasUniform2i() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform2i(
		this.Ref(),
	)
}

// Uniform2iFunc returns the method "WebGLRenderingContext.uniform2i".
func (this WebGLRenderingContext) Uniform2iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2iFunc(
			this.Ref(),
		),
	)
}

// Uniform2i calls the method "WebGLRenderingContext.uniform2i".
func (this WebGLRenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// TryUniform2i calls the method "WebGLRenderingContext.uniform2i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform2i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasUniform3i returns true if the method "WebGLRenderingContext.uniform3i" exists.
func (this WebGLRenderingContext) HasUniform3i() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform3i(
		this.Ref(),
	)
}

// Uniform3iFunc returns the method "WebGLRenderingContext.uniform3i".
func (this WebGLRenderingContext) Uniform3iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3iFunc(
			this.Ref(),
		),
	)
}

// Uniform3i calls the method "WebGLRenderingContext.uniform3i".
func (this WebGLRenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// TryUniform3i calls the method "WebGLRenderingContext.uniform3i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform3i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// HasUniform4i returns true if the method "WebGLRenderingContext.uniform4i" exists.
func (this WebGLRenderingContext) HasUniform4i() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform4i(
		this.Ref(),
	)
}

// Uniform4iFunc returns the method "WebGLRenderingContext.uniform4i".
func (this WebGLRenderingContext) Uniform4iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4iFunc(
			this.Ref(),
		),
	)
}

// Uniform4i calls the method "WebGLRenderingContext.uniform4i".
func (this WebGLRenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// TryUniform4i calls the method "WebGLRenderingContext.uniform4i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform4i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasUseProgram returns true if the method "WebGLRenderingContext.useProgram" exists.
func (this WebGLRenderingContext) HasUseProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextUseProgram(
		this.Ref(),
	)
}

// UseProgramFunc returns the method "WebGLRenderingContext.useProgram".
func (this WebGLRenderingContext) UseProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUseProgramFunc(
			this.Ref(),
		),
	)
}

// UseProgram calls the method "WebGLRenderingContext.useProgram".
func (this WebGLRenderingContext) UseProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextUseProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryUseProgram calls the method "WebGLRenderingContext.useProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUseProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUseProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasValidateProgram returns true if the method "WebGLRenderingContext.validateProgram" exists.
func (this WebGLRenderingContext) HasValidateProgram() bool {
	return js.True == bindings.HasWebGLRenderingContextValidateProgram(
		this.Ref(),
	)
}

// ValidateProgramFunc returns the method "WebGLRenderingContext.validateProgram".
func (this WebGLRenderingContext) ValidateProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextValidateProgramFunc(
			this.Ref(),
		),
	)
}

// ValidateProgram calls the method "WebGLRenderingContext.validateProgram".
func (this WebGLRenderingContext) ValidateProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextValidateProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryValidateProgram calls the method "WebGLRenderingContext.validateProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryValidateProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextValidateProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasVertexAttrib1f returns true if the method "WebGLRenderingContext.vertexAttrib1f" exists.
func (this WebGLRenderingContext) HasVertexAttrib1f() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib1f(
		this.Ref(),
	)
}

// VertexAttrib1fFunc returns the method "WebGLRenderingContext.vertexAttrib1f".
func (this WebGLRenderingContext) VertexAttrib1fFunc() (fn js.Func[func(index GLuint, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib1fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1f calls the method "WebGLRenderingContext.vertexAttrib1f".
func (this WebGLRenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
	)

	return
}

// TryVertexAttrib1f calls the method "WebGLRenderingContext.vertexAttrib1f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib1f(index GLuint, x GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
	)

	return
}

// HasVertexAttrib2f returns true if the method "WebGLRenderingContext.vertexAttrib2f" exists.
func (this WebGLRenderingContext) HasVertexAttrib2f() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib2f(
		this.Ref(),
	)
}

// VertexAttrib2fFunc returns the method "WebGLRenderingContext.vertexAttrib2f".
func (this WebGLRenderingContext) VertexAttrib2fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib2fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2f calls the method "WebGLRenderingContext.vertexAttrib2f".
func (this WebGLRenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// TryVertexAttrib2f calls the method "WebGLRenderingContext.vertexAttrib2f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// HasVertexAttrib3f returns true if the method "WebGLRenderingContext.vertexAttrib3f" exists.
func (this WebGLRenderingContext) HasVertexAttrib3f() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib3f(
		this.Ref(),
	)
}

// VertexAttrib3fFunc returns the method "WebGLRenderingContext.vertexAttrib3f".
func (this WebGLRenderingContext) VertexAttrib3fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib3fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3f calls the method "WebGLRenderingContext.vertexAttrib3f".
func (this WebGLRenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TryVertexAttrib3f calls the method "WebGLRenderingContext.vertexAttrib3f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasVertexAttrib4f returns true if the method "WebGLRenderingContext.vertexAttrib4f" exists.
func (this WebGLRenderingContext) HasVertexAttrib4f() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib4f(
		this.Ref(),
	)
}

// VertexAttrib4fFunc returns the method "WebGLRenderingContext.vertexAttrib4f".
func (this WebGLRenderingContext) VertexAttrib4fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib4fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4f calls the method "WebGLRenderingContext.vertexAttrib4f".
func (this WebGLRenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// TryVertexAttrib4f calls the method "WebGLRenderingContext.vertexAttrib4f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasVertexAttrib1fv returns true if the method "WebGLRenderingContext.vertexAttrib1fv" exists.
func (this WebGLRenderingContext) HasVertexAttrib1fv() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib1fv(
		this.Ref(),
	)
}

// VertexAttrib1fvFunc returns the method "WebGLRenderingContext.vertexAttrib1fv".
func (this WebGLRenderingContext) VertexAttrib1fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib1fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1fv calls the method "WebGLRenderingContext.vertexAttrib1fv".
func (this WebGLRenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib1fv calls the method "WebGLRenderingContext.vertexAttrib1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib1fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib2fv returns true if the method "WebGLRenderingContext.vertexAttrib2fv" exists.
func (this WebGLRenderingContext) HasVertexAttrib2fv() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib2fv(
		this.Ref(),
	)
}

// VertexAttrib2fvFunc returns the method "WebGLRenderingContext.vertexAttrib2fv".
func (this WebGLRenderingContext) VertexAttrib2fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib2fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2fv calls the method "WebGLRenderingContext.vertexAttrib2fv".
func (this WebGLRenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib2fv calls the method "WebGLRenderingContext.vertexAttrib2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib2fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib3fv returns true if the method "WebGLRenderingContext.vertexAttrib3fv" exists.
func (this WebGLRenderingContext) HasVertexAttrib3fv() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib3fv(
		this.Ref(),
	)
}

// VertexAttrib3fvFunc returns the method "WebGLRenderingContext.vertexAttrib3fv".
func (this WebGLRenderingContext) VertexAttrib3fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib3fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3fv calls the method "WebGLRenderingContext.vertexAttrib3fv".
func (this WebGLRenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib3fv calls the method "WebGLRenderingContext.vertexAttrib3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib3fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib4fv returns true if the method "WebGLRenderingContext.vertexAttrib4fv" exists.
func (this WebGLRenderingContext) HasVertexAttrib4fv() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttrib4fv(
		this.Ref(),
	)
}

// VertexAttrib4fvFunc returns the method "WebGLRenderingContext.vertexAttrib4fv".
func (this WebGLRenderingContext) VertexAttrib4fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttrib4fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4fv calls the method "WebGLRenderingContext.vertexAttrib4fv".
func (this WebGLRenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib4fv calls the method "WebGLRenderingContext.vertexAttrib4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttrib4fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttribPointer returns true if the method "WebGLRenderingContext.vertexAttribPointer" exists.
func (this WebGLRenderingContext) HasVertexAttribPointer() bool {
	return js.True == bindings.HasWebGLRenderingContextVertexAttribPointer(
		this.Ref(),
	)
}

// VertexAttribPointerFunc returns the method "WebGLRenderingContext.vertexAttribPointer".
func (this WebGLRenderingContext) VertexAttribPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextVertexAttribPointerFunc(
			this.Ref(),
		),
	)
}

// VertexAttribPointer calls the method "WebGLRenderingContext.vertexAttribPointer".
func (this WebGLRenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// TryVertexAttribPointer calls the method "WebGLRenderingContext.vertexAttribPointer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryVertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// HasViewport returns true if the method "WebGLRenderingContext.viewport" exists.
func (this WebGLRenderingContext) HasViewport() bool {
	return js.True == bindings.HasWebGLRenderingContextViewport(
		this.Ref(),
	)
}

// ViewportFunc returns the method "WebGLRenderingContext.viewport".
func (this WebGLRenderingContext) ViewportFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextViewportFunc(
			this.Ref(),
		),
	)
}

// Viewport calls the method "WebGLRenderingContext.viewport".
func (this WebGLRenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextViewport(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryViewport calls the method "WebGLRenderingContext.viewport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryViewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextViewport(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasMakeXRCompatible returns true if the method "WebGLRenderingContext.makeXRCompatible" exists.
func (this WebGLRenderingContext) HasMakeXRCompatible() bool {
	return js.True == bindings.HasWebGLRenderingContextMakeXRCompatible(
		this.Ref(),
	)
}

// MakeXRCompatibleFunc returns the method "WebGLRenderingContext.makeXRCompatible".
func (this WebGLRenderingContext) MakeXRCompatibleFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextMakeXRCompatibleFunc(
			this.Ref(),
		),
	)
}

// MakeXRCompatible calls the method "WebGLRenderingContext.makeXRCompatible".
func (this WebGLRenderingContext) MakeXRCompatible() (ret js.Promise[js.Void]) {
	bindings.CallWebGLRenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMakeXRCompatible calls the method "WebGLRenderingContext.makeXRCompatible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryMakeXRCompatible() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBufferData returns true if the method "WebGLRenderingContext.bufferData" exists.
func (this WebGLRenderingContext) HasBufferData() bool {
	return js.True == bindings.HasWebGLRenderingContextBufferData(
		this.Ref(),
	)
}

// BufferDataFunc returns the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferDataFunc() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferDataFunc(
			this.Ref(),
		),
	)
}

// BufferData calls the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferData(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// TryBufferData calls the method "WebGLRenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBufferData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// HasBufferData1 returns true if the method "WebGLRenderingContext.bufferData" exists.
func (this WebGLRenderingContext) HasBufferData1() bool {
	return js.True == bindings.HasWebGLRenderingContextBufferData1(
		this.Ref(),
	)
}

// BufferData1Func returns the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferData1Func() (fn js.Func[func(target GLenum, data AllowSharedBufferSource, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferData1Func(
			this.Ref(),
		),
	)
}

// BufferData1 calls the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferData1(target GLenum, data AllowSharedBufferSource, usage GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferData1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		data.Ref(),
		uint32(usage),
	)

	return
}

// TryBufferData1 calls the method "WebGLRenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBufferData1(target GLenum, data AllowSharedBufferSource, usage GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBufferData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		data.Ref(),
		uint32(usage),
	)

	return
}

// HasBufferSubData returns true if the method "WebGLRenderingContext.bufferSubData" exists.
func (this WebGLRenderingContext) HasBufferSubData() bool {
	return js.True == bindings.HasWebGLRenderingContextBufferSubData(
		this.Ref(),
	)
}

// BufferSubDataFunc returns the method "WebGLRenderingContext.bufferSubData".
func (this WebGLRenderingContext) BufferSubDataFunc() (fn js.Func[func(target GLenum, offset GLintptr, data AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// BufferSubData calls the method "WebGLRenderingContext.bufferSubData".
func (this WebGLRenderingContext) BufferSubData(target GLenum, offset GLintptr, data AllowSharedBufferSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferSubData(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(offset),
		data.Ref(),
	)

	return
}

// TryBufferSubData calls the method "WebGLRenderingContext.bufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBufferSubData(target GLenum, offset GLintptr, data AllowSharedBufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBufferSubData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(offset),
		data.Ref(),
	)

	return
}

// HasCompressedTexImage2D returns true if the method "WebGLRenderingContext.compressedTexImage2D" exists.
func (this WebGLRenderingContext) HasCompressedTexImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextCompressedTexImage2D(
		this.Ref(),
	)
}

// CompressedTexImage2DFunc returns the method "WebGLRenderingContext.compressedTexImage2D".
func (this WebGLRenderingContext) CompressedTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompressedTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D calls the method "WebGLRenderingContext.compressedTexImage2D".
func (this WebGLRenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		data.Ref(),
	)

	return
}

// TryCompressedTexImage2D calls the method "WebGLRenderingContext.compressedTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		data.Ref(),
	)

	return
}

// HasCompressedTexSubImage2D returns true if the method "WebGLRenderingContext.compressedTexSubImage2D" exists.
func (this WebGLRenderingContext) HasCompressedTexSubImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextCompressedTexSubImage2D(
		this.Ref(),
	)
}

// CompressedTexSubImage2DFunc returns the method "WebGLRenderingContext.compressedTexSubImage2D".
func (this WebGLRenderingContext) CompressedTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextCompressedTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D calls the method "WebGLRenderingContext.compressedTexSubImage2D".
func (this WebGLRenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		data.Ref(),
	)

	return
}

// TryCompressedTexSubImage2D calls the method "WebGLRenderingContext.compressedTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		data.Ref(),
	)

	return
}

// HasReadPixels returns true if the method "WebGLRenderingContext.readPixels" exists.
func (this WebGLRenderingContext) HasReadPixels() bool {
	return js.True == bindings.HasWebGLRenderingContextReadPixels(
		this.Ref(),
	)
}

// ReadPixelsFunc returns the method "WebGLRenderingContext.readPixels".
func (this WebGLRenderingContext) ReadPixelsFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextReadPixelsFunc(
			this.Ref(),
		),
	)
}

// ReadPixels calls the method "WebGLRenderingContext.readPixels".
func (this WebGLRenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextReadPixels(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	return
}

// TryReadPixels calls the method "WebGLRenderingContext.readPixels"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextReadPixels(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		pixels.Ref(),
	)

	return
}

// HasTexImage2D returns true if the method "WebGLRenderingContext.texImage2D" exists.
func (this WebGLRenderingContext) HasTexImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextTexImage2D(
		this.Ref(),
	)
}

// TexImage2DFunc returns the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexImage2DFunc(
			this.Ref(),
		),
	)
}

// TexImage2D calls the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexImage2D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage2D calls the method "WebGLRenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage2D1 returns true if the method "WebGLRenderingContext.texImage2D" exists.
func (this WebGLRenderingContext) HasTexImage2D1() bool {
	return js.True == bindings.HasWebGLRenderingContextTexImage2D1(
		this.Ref(),
	)
}

// TexImage2D1Func returns the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexImage2D1Func(
			this.Ref(),
		),
	)
}

// TexImage2D1 calls the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// TryTexImage2D1 calls the method "WebGLRenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasTexSubImage2D returns true if the method "WebGLRenderingContext.texSubImage2D" exists.
func (this WebGLRenderingContext) HasTexSubImage2D() bool {
	return js.True == bindings.HasWebGLRenderingContextTexSubImage2D(
		this.Ref(),
	)
}

// TexSubImage2DFunc returns the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage2D calls the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage2D calls the method "WebGLRenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage2D1 returns true if the method "WebGLRenderingContext.texSubImage2D" exists.
func (this WebGLRenderingContext) HasTexSubImage2D1() bool {
	return js.True == bindings.HasWebGLRenderingContextTexSubImage2D1(
		this.Ref(),
	)
}

// TexSubImage2D1Func returns the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D1 calls the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// TryTexSubImage2D1 calls the method "WebGLRenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasUniform1fv returns true if the method "WebGLRenderingContext.uniform1fv" exists.
func (this WebGLRenderingContext) HasUniform1fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform1fv(
		this.Ref(),
	)
}

// Uniform1fvFunc returns the method "WebGLRenderingContext.uniform1fv".
func (this WebGLRenderingContext) Uniform1fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1fvFunc(
			this.Ref(),
		),
	)
}

// Uniform1fv calls the method "WebGLRenderingContext.uniform1fv".
func (this WebGLRenderingContext) Uniform1fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform1fv calls the method "WebGLRenderingContext.uniform1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform1fv(location WebGLUniformLocation, v Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform1fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform2fv returns true if the method "WebGLRenderingContext.uniform2fv" exists.
func (this WebGLRenderingContext) HasUniform2fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform2fv(
		this.Ref(),
	)
}

// Uniform2fvFunc returns the method "WebGLRenderingContext.uniform2fv".
func (this WebGLRenderingContext) Uniform2fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2fvFunc(
			this.Ref(),
		),
	)
}

// Uniform2fv calls the method "WebGLRenderingContext.uniform2fv".
func (this WebGLRenderingContext) Uniform2fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform2fv calls the method "WebGLRenderingContext.uniform2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform2fv(location WebGLUniformLocation, v Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform3fv returns true if the method "WebGLRenderingContext.uniform3fv" exists.
func (this WebGLRenderingContext) HasUniform3fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform3fv(
		this.Ref(),
	)
}

// Uniform3fvFunc returns the method "WebGLRenderingContext.uniform3fv".
func (this WebGLRenderingContext) Uniform3fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3fvFunc(
			this.Ref(),
		),
	)
}

// Uniform3fv calls the method "WebGLRenderingContext.uniform3fv".
func (this WebGLRenderingContext) Uniform3fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform3fv calls the method "WebGLRenderingContext.uniform3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform3fv(location WebGLUniformLocation, v Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform4fv returns true if the method "WebGLRenderingContext.uniform4fv" exists.
func (this WebGLRenderingContext) HasUniform4fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform4fv(
		this.Ref(),
	)
}

// Uniform4fvFunc returns the method "WebGLRenderingContext.uniform4fv".
func (this WebGLRenderingContext) Uniform4fvFunc() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4fvFunc(
			this.Ref(),
		),
	)
}

// Uniform4fv calls the method "WebGLRenderingContext.uniform4fv".
func (this WebGLRenderingContext) Uniform4fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform4fv calls the method "WebGLRenderingContext.uniform4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform4fv(location WebGLUniformLocation, v Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform1iv returns true if the method "WebGLRenderingContext.uniform1iv" exists.
func (this WebGLRenderingContext) HasUniform1iv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform1iv(
		this.Ref(),
	)
}

// Uniform1ivFunc returns the method "WebGLRenderingContext.uniform1iv".
func (this WebGLRenderingContext) Uniform1ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform1ivFunc(
			this.Ref(),
		),
	)
}

// Uniform1iv calls the method "WebGLRenderingContext.uniform1iv".
func (this WebGLRenderingContext) Uniform1iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform1iv calls the method "WebGLRenderingContext.uniform1iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform1iv(location WebGLUniformLocation, v Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform1iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform2iv returns true if the method "WebGLRenderingContext.uniform2iv" exists.
func (this WebGLRenderingContext) HasUniform2iv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform2iv(
		this.Ref(),
	)
}

// Uniform2ivFunc returns the method "WebGLRenderingContext.uniform2iv".
func (this WebGLRenderingContext) Uniform2ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform2ivFunc(
			this.Ref(),
		),
	)
}

// Uniform2iv calls the method "WebGLRenderingContext.uniform2iv".
func (this WebGLRenderingContext) Uniform2iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform2iv calls the method "WebGLRenderingContext.uniform2iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform2iv(location WebGLUniformLocation, v Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform2iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform3iv returns true if the method "WebGLRenderingContext.uniform3iv" exists.
func (this WebGLRenderingContext) HasUniform3iv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform3iv(
		this.Ref(),
	)
}

// Uniform3ivFunc returns the method "WebGLRenderingContext.uniform3iv".
func (this WebGLRenderingContext) Uniform3ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform3ivFunc(
			this.Ref(),
		),
	)
}

// Uniform3iv calls the method "WebGLRenderingContext.uniform3iv".
func (this WebGLRenderingContext) Uniform3iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform3iv calls the method "WebGLRenderingContext.uniform3iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform3iv(location WebGLUniformLocation, v Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform3iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniform4iv returns true if the method "WebGLRenderingContext.uniform4iv" exists.
func (this WebGLRenderingContext) HasUniform4iv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniform4iv(
		this.Ref(),
	)
}

// Uniform4ivFunc returns the method "WebGLRenderingContext.uniform4iv".
func (this WebGLRenderingContext) Uniform4ivFunc() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniform4ivFunc(
			this.Ref(),
		),
	)
}

// Uniform4iv calls the method "WebGLRenderingContext.uniform4iv".
func (this WebGLRenderingContext) Uniform4iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		v.Ref(),
	)

	return
}

// TryUniform4iv calls the method "WebGLRenderingContext.uniform4iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniform4iv(location WebGLUniformLocation, v Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniform4iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasUniformMatrix2fv returns true if the method "WebGLRenderingContext.uniformMatrix2fv" exists.
func (this WebGLRenderingContext) HasUniformMatrix2fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniformMatrix2fv(
		this.Ref(),
	)
}

// UniformMatrix2fvFunc returns the method "WebGLRenderingContext.uniformMatrix2fv".
func (this WebGLRenderingContext) UniformMatrix2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv calls the method "WebGLRenderingContext.uniformMatrix2fv".
func (this WebGLRenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// TryUniformMatrix2fv calls the method "WebGLRenderingContext.uniformMatrix2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// HasUniformMatrix3fv returns true if the method "WebGLRenderingContext.uniformMatrix3fv" exists.
func (this WebGLRenderingContext) HasUniformMatrix3fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniformMatrix3fv(
		this.Ref(),
	)
}

// UniformMatrix3fvFunc returns the method "WebGLRenderingContext.uniformMatrix3fv".
func (this WebGLRenderingContext) UniformMatrix3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv calls the method "WebGLRenderingContext.uniformMatrix3fv".
func (this WebGLRenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// TryUniformMatrix3fv calls the method "WebGLRenderingContext.uniformMatrix3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// HasUniformMatrix4fv returns true if the method "WebGLRenderingContext.uniformMatrix4fv" exists.
func (this WebGLRenderingContext) HasUniformMatrix4fv() bool {
	return js.True == bindings.HasWebGLRenderingContextUniformMatrix4fv(
		this.Ref(),
	)
}

// UniformMatrix4fvFunc returns the method "WebGLRenderingContext.uniformMatrix4fv".
func (this WebGLRenderingContext) UniformMatrix4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	return fn.FromRef(
		bindings.WebGLRenderingContextUniformMatrix4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv calls the method "WebGLRenderingContext.uniformMatrix4fv".
func (this WebGLRenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// TryUniformMatrix4fv calls the method "WebGLRenderingContext.uniformMatrix4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextCanvas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferWidth returns the value of property "WebGL2RenderingContext.drawingBufferWidth".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferWidth() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferHeight returns the value of property "WebGL2RenderingContext.drawingBufferHeight".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferHeight() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DrawingBufferColorSpace returns the value of property "WebGL2RenderingContext.drawingBufferColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDrawingBufferColorSpace sets the value of property "WebGL2RenderingContext.drawingBufferColorSpace" to val.
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
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) UnpackColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextUnpackColorSpace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUnpackColorSpace sets the value of property "WebGL2RenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGL2RenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGL2RenderingContextUnpackColorSpace(
		this.Ref(),
		uint32(val),
	)
}

// HasBufferData returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasBufferData() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferData(
		this.Ref(),
	)
}

// BufferDataFunc returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferDataFunc() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferDataFunc(
			this.Ref(),
		),
	)
}

// BufferData calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// TryBufferData calls the method "WebGL2RenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// HasBufferData1 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasBufferData1() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferData1(
		this.Ref(),
	)
}

// BufferData1Func returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData1Func() (fn js.Func[func(target GLenum, srcData AllowSharedBufferSource, usage GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData1Func(
			this.Ref(),
		),
	)
}

// BufferData1 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData1(target GLenum, srcData AllowSharedBufferSource, usage GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
	)

	return
}

// TryBufferData1 calls the method "WebGL2RenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferData1(target GLenum, srcData AllowSharedBufferSource, usage GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
	)

	return
}

// HasBufferSubData returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasBufferSubData() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferSubData(
		this.Ref(),
	)
}

// BufferSubDataFunc returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubDataFunc() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// BufferSubData calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
	)

	return
}

// TryBufferSubData calls the method "WebGL2RenderingContext.bufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferSubData(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferSubData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
	)

	return
}

// HasBufferData2 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasBufferData2() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferData2(
		this.Ref(),
	)
}

// BufferData2Func returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData2Func() (fn js.Func[func(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData2Func(
			this.Ref(),
		),
	)
}

// BufferData2 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData2(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData2(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// TryBufferData2 calls the method "WebGL2RenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferData2(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferData2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// HasBufferData3 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasBufferData3() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferData3(
		this.Ref(),
	)
}

// BufferData3Func returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData3Func() (fn js.Func[func(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferData3Func(
			this.Ref(),
		),
	)
}

// BufferData3 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData3(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData3(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
	)

	return
}

// TryBufferData3 calls the method "WebGL2RenderingContext.bufferData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferData3(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferData3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
	)

	return
}

// HasBufferSubData1 returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasBufferSubData1() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferSubData1(
		this.Ref(),
	)
}

// BufferSubData1Func returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData1Func() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubData1Func(
			this.Ref(),
		),
	)
}

// BufferSubData1 calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData1(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// TryBufferSubData1 calls the method "WebGL2RenderingContext.bufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferSubData1(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint, length GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferSubData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// HasBufferSubData2 returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasBufferSubData2() bool {
	return js.True == bindings.HasWebGL2RenderingContextBufferSubData2(
		this.Ref(),
	)
}

// BufferSubData2Func returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData2Func() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBufferSubData2Func(
			this.Ref(),
		),
	)
}

// BufferSubData2 calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData2(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData2(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryBufferSubData2 calls the method "WebGL2RenderingContext.bufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBufferSubData2(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBufferSubData2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasTexImage2D returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasTexImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage2D(
		this.Ref(),
	)
}

// TexImage2DFunc returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2DFunc(
			this.Ref(),
		),
	)
}

// TexImage2D calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage2D calls the method "WebGL2RenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage2D1 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasTexImage2D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage2D1(
		this.Ref(),
	)
}

// TexImage2D1Func returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D1Func(
			this.Ref(),
		),
	)
}

// TexImage2D1 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// TryTexImage2D1 calls the method "WebGL2RenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasTexSubImage2D returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasTexSubImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage2D(
		this.Ref(),
	)
}

// TexSubImage2DFunc returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage2D calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage2D calls the method "WebGL2RenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage2D1 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasTexSubImage2D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage2D1(
		this.Ref(),
	)
}

// TexSubImage2D1Func returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D1 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// TryTexSubImage2D1 calls the method "WebGL2RenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasTexImage2D2 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasTexImage2D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage2D2(
		this.Ref(),
	)
}

// TexImage2D2Func returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D2Func(
			this.Ref(),
		),
	)
}

// TexImage2D2 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage2D2 calls the method "WebGL2RenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage2D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage2D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage2D3 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasTexImage2D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage2D3(
		this.Ref(),
	)
}

// TexImage2D3Func returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D3Func(
			this.Ref(),
		),
	)
}

// TexImage2D3 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage2D3 calls the method "WebGL2RenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage2D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage2D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage2D4 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasTexImage2D4() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage2D4(
		this.Ref(),
	)
}

// TexImage2D4Func returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D4Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage2D4Func(
			this.Ref(),
		),
	)
}

// TexImage2D4 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D4(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D4(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage2D4 calls the method "WebGL2RenderingContext.texImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage2D4(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage2D4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage2D2 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasTexSubImage2D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage2D2(
		this.Ref(),
	)
}

// TexSubImage2D2Func returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D2Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D2 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage2D2 calls the method "WebGL2RenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage2D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage2D3 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasTexSubImage2D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage2D3(
		this.Ref(),
	)
}

// TexSubImage2D3Func returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D3Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D3 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage2D3 calls the method "WebGL2RenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage2D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage2D4 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasTexSubImage2D4() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage2D4(
		this.Ref(),
	)
}

// TexSubImage2D4Func returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D4Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage2D4Func(
			this.Ref(),
		),
	)
}

// TexSubImage2D4 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D4(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D4(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage2D4 calls the method "WebGL2RenderingContext.texSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage2D4(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage2D4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage2D returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage2D(
		this.Ref(),
	)
}

// CompressedTexImage2DFunc returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		int32(imageSize),
		float64(offset),
	)

	return
}

// TryCompressedTexImage2D calls the method "WebGL2RenderingContext.compressedTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		int32(imageSize),
		float64(offset),
	)

	return
}

// HasCompressedTexImage2D1 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage2D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage2D1(
		this.Ref(),
	)
}

// CompressedTexImage2D1Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D1 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexImage2D1 calls the method "WebGL2RenderingContext.compressedTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage2D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage2D2 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage2D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage2D2(
		this.Ref(),
	)
}

// CompressedTexImage2D2Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D2 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D2(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryCompressedTexImage2D2 calls the method "WebGL2RenderingContext.compressedTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage2D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage2D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasCompressedTexImage2D3 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage2D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage2D3(
		this.Ref(),
	)
}

// CompressedTexImage2D3Func returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage2D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage2D3 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D3(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
	)

	return
}

// TryCompressedTexImage2D3 calls the method "WebGL2RenderingContext.compressedTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage2D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage2D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(border),
		srcData.Ref(),
	)

	return
}

// HasCompressedTexSubImage2D returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage2D(
		this.Ref(),
	)
}

// CompressedTexSubImage2DFunc returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage2D calls the method "WebGL2RenderingContext.compressedTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage2D1 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage2D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage2D1(
		this.Ref(),
	)
}

// CompressedTexSubImage2D1Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D1 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage2D1 calls the method "WebGL2RenderingContext.compressedTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage2D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage2D2 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage2D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage2D2(
		this.Ref(),
	)
}

// CompressedTexSubImage2D2Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D2 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage2D2 calls the method "WebGL2RenderingContext.compressedTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage2D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage2D3 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage2D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage2D3(
		this.Ref(),
	)
}

// CompressedTexSubImage2D3Func returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage2D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage2D3 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D3(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		srcData.Ref(),
	)

	return
}

// TryCompressedTexSubImage2D3 calls the method "WebGL2RenderingContext.compressedTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage2D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(width),
		int32(height),
		uint32(format),
		srcData.Ref(),
	)

	return
}

// HasUniform1fv returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasUniform1fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1fv(
		this.Ref(),
	)
}

// Uniform1fvFunc returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fvFunc(
			this.Ref(),
		),
	)
}

// Uniform1fv calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform1fv calls the method "WebGL2RenderingContext.uniform1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform1fv1 returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasUniform1fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1fv1(
		this.Ref(),
	)
}

// Uniform1fv1Func returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fv1Func(
			this.Ref(),
		),
	)
}

// Uniform1fv1 calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform1fv1 calls the method "WebGL2RenderingContext.uniform1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform1fv2 returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasUniform1fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1fv2(
		this.Ref(),
	)
}

// Uniform1fv2Func returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fv2Func(
			this.Ref(),
		),
	)
}

// Uniform1fv2 calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform1fv2 calls the method "WebGL2RenderingContext.uniform1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1fv2(location WebGLUniformLocation, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform2fv returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasUniform2fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2fv(
		this.Ref(),
	)
}

// Uniform2fvFunc returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fvFunc(
			this.Ref(),
		),
	)
}

// Uniform2fv calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform2fv calls the method "WebGL2RenderingContext.uniform2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform2fv1 returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasUniform2fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2fv1(
		this.Ref(),
	)
}

// Uniform2fv1Func returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fv1Func(
			this.Ref(),
		),
	)
}

// Uniform2fv1 calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform2fv1 calls the method "WebGL2RenderingContext.uniform2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform2fv2 returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasUniform2fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2fv2(
		this.Ref(),
	)
}

// Uniform2fv2Func returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fv2Func(
			this.Ref(),
		),
	)
}

// Uniform2fv2 calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform2fv2 calls the method "WebGL2RenderingContext.uniform2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2fv2(location WebGLUniformLocation, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform3fv returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasUniform3fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3fv(
		this.Ref(),
	)
}

// Uniform3fvFunc returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fvFunc(
			this.Ref(),
		),
	)
}

// Uniform3fv calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform3fv calls the method "WebGL2RenderingContext.uniform3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform3fv1 returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasUniform3fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3fv1(
		this.Ref(),
	)
}

// Uniform3fv1Func returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fv1Func(
			this.Ref(),
		),
	)
}

// Uniform3fv1 calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform3fv1 calls the method "WebGL2RenderingContext.uniform3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform3fv2 returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasUniform3fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3fv2(
		this.Ref(),
	)
}

// Uniform3fv2Func returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fv2Func(
			this.Ref(),
		),
	)
}

// Uniform3fv2 calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform3fv2 calls the method "WebGL2RenderingContext.uniform3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3fv2(location WebGLUniformLocation, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform4fv returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasUniform4fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4fv(
		this.Ref(),
	)
}

// Uniform4fvFunc returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fvFunc() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fvFunc(
			this.Ref(),
		),
	)
}

// Uniform4fv calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform4fv calls the method "WebGL2RenderingContext.uniform4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform4fv1 returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasUniform4fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4fv1(
		this.Ref(),
	)
}

// Uniform4fv1Func returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv1Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fv1Func(
			this.Ref(),
		),
	)
}

// Uniform4fv1 calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform4fv1 calls the method "WebGL2RenderingContext.uniform4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform4fv2 returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasUniform4fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4fv2(
		this.Ref(),
	)
}

// Uniform4fv2Func returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv2Func() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fv2Func(
			this.Ref(),
		),
	)
}

// Uniform4fv2 calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform4fv2 calls the method "WebGL2RenderingContext.uniform4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4fv2(location WebGLUniformLocation, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform1iv returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasUniform1iv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1iv(
		this.Ref(),
	)
}

// Uniform1ivFunc returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1ivFunc(
			this.Ref(),
		),
	)
}

// Uniform1iv calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform1iv calls the method "WebGL2RenderingContext.uniform1iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform1iv1 returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasUniform1iv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1iv1(
		this.Ref(),
	)
}

// Uniform1iv1Func returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iv1Func(
			this.Ref(),
		),
	)
}

// Uniform1iv1 calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform1iv1 calls the method "WebGL2RenderingContext.uniform1iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1iv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform1iv2 returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasUniform1iv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1iv2(
		this.Ref(),
	)
}

// Uniform1iv2Func returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iv2Func(
			this.Ref(),
		),
	)
}

// Uniform1iv2 calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform1iv2 calls the method "WebGL2RenderingContext.uniform1iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1iv2(location WebGLUniformLocation, data Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1iv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform2iv returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasUniform2iv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2iv(
		this.Ref(),
	)
}

// Uniform2ivFunc returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2ivFunc(
			this.Ref(),
		),
	)
}

// Uniform2iv calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform2iv calls the method "WebGL2RenderingContext.uniform2iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform2iv1 returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasUniform2iv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2iv1(
		this.Ref(),
	)
}

// Uniform2iv1Func returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iv1Func(
			this.Ref(),
		),
	)
}

// Uniform2iv1 calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform2iv1 calls the method "WebGL2RenderingContext.uniform2iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2iv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform2iv2 returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasUniform2iv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2iv2(
		this.Ref(),
	)
}

// Uniform2iv2Func returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iv2Func(
			this.Ref(),
		),
	)
}

// Uniform2iv2 calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform2iv2 calls the method "WebGL2RenderingContext.uniform2iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2iv2(location WebGLUniformLocation, data Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2iv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform3iv returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasUniform3iv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3iv(
		this.Ref(),
	)
}

// Uniform3ivFunc returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3ivFunc(
			this.Ref(),
		),
	)
}

// Uniform3iv calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform3iv calls the method "WebGL2RenderingContext.uniform3iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform3iv1 returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasUniform3iv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3iv1(
		this.Ref(),
	)
}

// Uniform3iv1Func returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iv1Func(
			this.Ref(),
		),
	)
}

// Uniform3iv1 calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform3iv1 calls the method "WebGL2RenderingContext.uniform3iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3iv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform3iv2 returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasUniform3iv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3iv2(
		this.Ref(),
	)
}

// Uniform3iv2Func returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iv2Func(
			this.Ref(),
		),
	)
}

// Uniform3iv2 calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform3iv2 calls the method "WebGL2RenderingContext.uniform3iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3iv2(location WebGLUniformLocation, data Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3iv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform4iv returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasUniform4iv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4iv(
		this.Ref(),
	)
}

// Uniform4ivFunc returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4ivFunc() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4ivFunc(
			this.Ref(),
		),
	)
}

// Uniform4iv calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform4iv calls the method "WebGL2RenderingContext.uniform4iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform4iv1 returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasUniform4iv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4iv1(
		this.Ref(),
	)
}

// Uniform4iv1Func returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv1Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iv1Func(
			this.Ref(),
		),
	)
}

// Uniform4iv1 calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform4iv1 calls the method "WebGL2RenderingContext.uniform4iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4iv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform4iv2 returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasUniform4iv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4iv2(
		this.Ref(),
	)
}

// Uniform4iv2Func returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv2Func() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iv2Func(
			this.Ref(),
		),
	)
}

// Uniform4iv2 calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform4iv2 calls the method "WebGL2RenderingContext.uniform4iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4iv2(location WebGLUniformLocation, data Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4iv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniformMatrix2fv returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2fv(
		this.Ref(),
	)
}

// UniformMatrix2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix2fv calls the method "WebGL2RenderingContext.uniformMatrix2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2fv1(
		this.Ref(),
	)
}

// UniformMatrix2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv1 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix2fv1 calls the method "WebGL2RenderingContext.uniformMatrix2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2fv2(
		this.Ref(),
	)
}

// UniformMatrix2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2fv2 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix2fv2 calls the method "WebGL2RenderingContext.uniformMatrix2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix3fv returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3fv(
		this.Ref(),
	)
}

// UniformMatrix3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix3fv calls the method "WebGL2RenderingContext.uniformMatrix3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3fv1(
		this.Ref(),
	)
}

// UniformMatrix3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv1 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix3fv1 calls the method "WebGL2RenderingContext.uniformMatrix3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3fv2(
		this.Ref(),
	)
}

// UniformMatrix3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3fv2 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix3fv2 calls the method "WebGL2RenderingContext.uniformMatrix3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix4fv returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4fv(
		this.Ref(),
	)
}

// UniformMatrix4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix4fv calls the method "WebGL2RenderingContext.uniformMatrix4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4fv1(
		this.Ref(),
	)
}

// UniformMatrix4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv1 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix4fv1 calls the method "WebGL2RenderingContext.uniformMatrix4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4fv2(
		this.Ref(),
	)
}

// UniformMatrix4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4fv2 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix4fv2 calls the method "WebGL2RenderingContext.uniformMatrix4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasReadPixels returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasReadPixels() bool {
	return js.True == bindings.HasWebGL2RenderingContextReadPixels(
		this.Ref(),
	)
}

// ReadPixelsFunc returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixelsFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixelsFunc(
			this.Ref(),
		),
	)
}

// ReadPixels calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
	)

	return
}

// TryReadPixels calls the method "WebGL2RenderingContext.readPixels"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextReadPixels(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
	)

	return
}

// HasReadPixels1 returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasReadPixels1() bool {
	return js.True == bindings.HasWebGL2RenderingContextReadPixels1(
		this.Ref(),
	)
}

// ReadPixels1Func returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels1Func() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixels1Func(
			this.Ref(),
		),
	)
}

// ReadPixels1 calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels1(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels1(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		float64(offset),
	)

	return
}

// TryReadPixels1 calls the method "WebGL2RenderingContext.readPixels"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryReadPixels1(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextReadPixels1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasReadPixels2 returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasReadPixels2() bool {
	return js.True == bindings.HasWebGL2RenderingContextReadPixels2(
		this.Ref(),
	)
}

// ReadPixels2Func returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels2Func() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView, dstOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadPixels2Func(
			this.Ref(),
		),
	)
}

// ReadPixels2 calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels2(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView, dstOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels2(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
		uint32(dstOffset),
	)

	return
}

// TryReadPixels2 calls the method "WebGL2RenderingContext.readPixels"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryReadPixels2(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView, dstOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextReadPixels2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		uint32(format),
		uint32(typ),
		dstData.Ref(),
		uint32(dstOffset),
	)

	return
}

// HasCopyBufferSubData returns true if the method "WebGL2RenderingContext.copyBufferSubData" exists.
func (this WebGL2RenderingContext) HasCopyBufferSubData() bool {
	return js.True == bindings.HasWebGL2RenderingContextCopyBufferSubData(
		this.Ref(),
	)
}

// CopyBufferSubDataFunc returns the method "WebGL2RenderingContext.copyBufferSubData".
func (this WebGL2RenderingContext) CopyBufferSubDataFunc() (fn js.Func[func(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// CopyBufferSubData calls the method "WebGL2RenderingContext.copyBufferSubData".
func (this WebGL2RenderingContext) CopyBufferSubData(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyBufferSubData(
		this.Ref(), js.Pointer(&ret),
		uint32(readTarget),
		uint32(writeTarget),
		float64(readOffset),
		float64(writeOffset),
		float64(size),
	)

	return
}

// TryCopyBufferSubData calls the method "WebGL2RenderingContext.copyBufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCopyBufferSubData(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCopyBufferSubData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(readTarget),
		uint32(writeTarget),
		float64(readOffset),
		float64(writeOffset),
		float64(size),
	)

	return
}

// HasGetBufferSubData returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasGetBufferSubData() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetBufferSubData(
		this.Ref(),
	)
}

// GetBufferSubDataFunc returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubDataFunc() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint, length GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubDataFunc(
			this.Ref(),
		),
	)
}

// GetBufferSubData calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
		uint32(length),
	)

	return
}

// TryGetBufferSubData calls the method "WebGL2RenderingContext.getBufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetBufferSubData(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint, length GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetBufferSubData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
		uint32(length),
	)

	return
}

// HasGetBufferSubData1 returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasGetBufferSubData1() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetBufferSubData1(
		this.Ref(),
	)
}

// GetBufferSubData1Func returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData1Func() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubData1Func(
			this.Ref(),
		),
	)
}

// GetBufferSubData1 calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData1(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData1(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
	)

	return
}

// TryGetBufferSubData1 calls the method "WebGL2RenderingContext.getBufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetBufferSubData1(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetBufferSubData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
	)

	return
}

// HasGetBufferSubData2 returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasGetBufferSubData2() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetBufferSubData2(
		this.Ref(),
	)
}

// GetBufferSubData2Func returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData2Func() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferSubData2Func(
			this.Ref(),
		),
	)
}

// GetBufferSubData2 calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData2(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData2(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
	)

	return
}

// TryGetBufferSubData2 calls the method "WebGL2RenderingContext.getBufferSubData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetBufferSubData2(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetBufferSubData2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
	)

	return
}

// HasBlitFramebuffer returns true if the method "WebGL2RenderingContext.blitFramebuffer" exists.
func (this WebGL2RenderingContext) HasBlitFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlitFramebuffer(
		this.Ref(),
	)
}

// BlitFramebufferFunc returns the method "WebGL2RenderingContext.blitFramebuffer".
func (this WebGL2RenderingContext) BlitFramebufferFunc() (fn js.Func[func(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlitFramebufferFunc(
			this.Ref(),
		),
	)
}

// BlitFramebuffer calls the method "WebGL2RenderingContext.blitFramebuffer".
func (this WebGL2RenderingContext) BlitFramebuffer(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlitFramebuffer(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryBlitFramebuffer calls the method "WebGL2RenderingContext.blitFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlitFramebuffer(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlitFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasFramebufferTextureLayer returns true if the method "WebGL2RenderingContext.framebufferTextureLayer" exists.
func (this WebGL2RenderingContext) HasFramebufferTextureLayer() bool {
	return js.True == bindings.HasWebGL2RenderingContextFramebufferTextureLayer(
		this.Ref(),
	)
}

// FramebufferTextureLayerFunc returns the method "WebGL2RenderingContext.framebufferTextureLayer".
func (this WebGL2RenderingContext) FramebufferTextureLayerFunc() (fn js.Func[func(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferTextureLayerFunc(
			this.Ref(),
		),
	)
}

// FramebufferTextureLayer calls the method "WebGL2RenderingContext.framebufferTextureLayer".
func (this WebGL2RenderingContext) FramebufferTextureLayer(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferTextureLayer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(layer),
	)

	return
}

// TryFramebufferTextureLayer calls the method "WebGL2RenderingContext.framebufferTextureLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFramebufferTextureLayer(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFramebufferTextureLayer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(layer),
	)

	return
}

// HasInvalidateFramebuffer returns true if the method "WebGL2RenderingContext.invalidateFramebuffer" exists.
func (this WebGL2RenderingContext) HasInvalidateFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextInvalidateFramebuffer(
		this.Ref(),
	)
}

// InvalidateFramebufferFunc returns the method "WebGL2RenderingContext.invalidateFramebuffer".
func (this WebGL2RenderingContext) InvalidateFramebufferFunc() (fn js.Func[func(target GLenum, attachments js.Array[GLenum])]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextInvalidateFramebufferFunc(
			this.Ref(),
		),
	)
}

// InvalidateFramebuffer calls the method "WebGL2RenderingContext.invalidateFramebuffer".
func (this WebGL2RenderingContext) InvalidateFramebuffer(target GLenum, attachments js.Array[GLenum]) (ret js.Void) {
	bindings.CallWebGL2RenderingContextInvalidateFramebuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		attachments.Ref(),
	)

	return
}

// TryInvalidateFramebuffer calls the method "WebGL2RenderingContext.invalidateFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryInvalidateFramebuffer(target GLenum, attachments js.Array[GLenum]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextInvalidateFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		attachments.Ref(),
	)

	return
}

// HasInvalidateSubFramebuffer returns true if the method "WebGL2RenderingContext.invalidateSubFramebuffer" exists.
func (this WebGL2RenderingContext) HasInvalidateSubFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextInvalidateSubFramebuffer(
		this.Ref(),
	)
}

// InvalidateSubFramebufferFunc returns the method "WebGL2RenderingContext.invalidateSubFramebuffer".
func (this WebGL2RenderingContext) InvalidateSubFramebufferFunc() (fn js.Func[func(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextInvalidateSubFramebufferFunc(
			this.Ref(),
		),
	)
}

// InvalidateSubFramebuffer calls the method "WebGL2RenderingContext.invalidateSubFramebuffer".
func (this WebGL2RenderingContext) InvalidateSubFramebuffer(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextInvalidateSubFramebuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		attachments.Ref(),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryInvalidateSubFramebuffer calls the method "WebGL2RenderingContext.invalidateSubFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryInvalidateSubFramebuffer(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextInvalidateSubFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		attachments.Ref(),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasReadBuffer returns true if the method "WebGL2RenderingContext.readBuffer" exists.
func (this WebGL2RenderingContext) HasReadBuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextReadBuffer(
		this.Ref(),
	)
}

// ReadBufferFunc returns the method "WebGL2RenderingContext.readBuffer".
func (this WebGL2RenderingContext) ReadBufferFunc() (fn js.Func[func(src GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextReadBufferFunc(
			this.Ref(),
		),
	)
}

// ReadBuffer calls the method "WebGL2RenderingContext.readBuffer".
func (this WebGL2RenderingContext) ReadBuffer(src GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(src),
	)

	return
}

// TryReadBuffer calls the method "WebGL2RenderingContext.readBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryReadBuffer(src GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextReadBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(src),
	)

	return
}

// HasGetInternalformatParameter returns true if the method "WebGL2RenderingContext.getInternalformatParameter" exists.
func (this WebGL2RenderingContext) HasGetInternalformatParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetInternalformatParameter(
		this.Ref(),
	)
}

// GetInternalformatParameterFunc returns the method "WebGL2RenderingContext.getInternalformatParameter".
func (this WebGL2RenderingContext) GetInternalformatParameterFunc() (fn js.Func[func(target GLenum, internalformat GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetInternalformatParameterFunc(
			this.Ref(),
		),
	)
}

// GetInternalformatParameter calls the method "WebGL2RenderingContext.getInternalformatParameter".
func (this WebGL2RenderingContext) GetInternalformatParameter(target GLenum, internalformat GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetInternalformatParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(internalformat),
		uint32(pname),
	)

	return
}

// TryGetInternalformatParameter calls the method "WebGL2RenderingContext.getInternalformatParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetInternalformatParameter(target GLenum, internalformat GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetInternalformatParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		uint32(pname),
	)

	return
}

// HasRenderbufferStorageMultisample returns true if the method "WebGL2RenderingContext.renderbufferStorageMultisample" exists.
func (this WebGL2RenderingContext) HasRenderbufferStorageMultisample() bool {
	return js.True == bindings.HasWebGL2RenderingContextRenderbufferStorageMultisample(
		this.Ref(),
	)
}

// RenderbufferStorageMultisampleFunc returns the method "WebGL2RenderingContext.renderbufferStorageMultisample".
func (this WebGL2RenderingContext) RenderbufferStorageMultisampleFunc() (fn js.Func[func(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextRenderbufferStorageMultisampleFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorageMultisample calls the method "WebGL2RenderingContext.renderbufferStorageMultisample".
func (this WebGL2RenderingContext) RenderbufferStorageMultisample(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextRenderbufferStorageMultisample(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(samples),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// TryRenderbufferStorageMultisample calls the method "WebGL2RenderingContext.renderbufferStorageMultisample"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryRenderbufferStorageMultisample(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextRenderbufferStorageMultisample(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(samples),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasTexStorage2D returns true if the method "WebGL2RenderingContext.texStorage2D" exists.
func (this WebGL2RenderingContext) HasTexStorage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexStorage2D(
		this.Ref(),
	)
}

// TexStorage2DFunc returns the method "WebGL2RenderingContext.texStorage2D".
func (this WebGL2RenderingContext) TexStorage2DFunc() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexStorage2DFunc(
			this.Ref(),
		),
	)
}

// TexStorage2D calls the method "WebGL2RenderingContext.texStorage2D".
func (this WebGL2RenderingContext) TexStorage2D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexStorage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// TryTexStorage2D calls the method "WebGL2RenderingContext.texStorage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexStorage2D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexStorage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasTexStorage3D returns true if the method "WebGL2RenderingContext.texStorage3D" exists.
func (this WebGL2RenderingContext) HasTexStorage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexStorage3D(
		this.Ref(),
	)
}

// TexStorage3DFunc returns the method "WebGL2RenderingContext.texStorage3D".
func (this WebGL2RenderingContext) TexStorage3DFunc() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexStorage3DFunc(
			this.Ref(),
		),
	)
}

// TexStorage3D calls the method "WebGL2RenderingContext.texStorage3D".
func (this WebGL2RenderingContext) TexStorage3D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexStorage3D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
	)

	return
}

// TryTexStorage3D calls the method "WebGL2RenderingContext.texStorage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexStorage3D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexStorage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
	)

	return
}

// HasTexImage3D returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasTexImage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage3D(
		this.Ref(),
	)
}

// TexImage3DFunc returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3DFunc(
			this.Ref(),
		),
	)
}

// TexImage3D calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage3D calls the method "WebGL2RenderingContext.texImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage3D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage3D1 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasTexImage3D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage3D1(
		this.Ref(),
	)
}

// TexImage3D1Func returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D1Func(
			this.Ref(),
		),
	)
}

// TexImage3D1 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D1(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage3D1 calls the method "WebGL2RenderingContext.texImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage3D1(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage3D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage3D2 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasTexImage3D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage3D2(
		this.Ref(),
	)
}

// TexImage3D2Func returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D2Func(
			this.Ref(),
		),
	)
}

// TexImage3D2 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage3D2 calls the method "WebGL2RenderingContext.texImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage3D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage3D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexImage3D3 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasTexImage3D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexImage3D3(
		this.Ref(),
	)
}

// TexImage3D3Func returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexImage3D3Func(
			this.Ref(),
		),
	)
}

// TexImage3D3 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexImage3D3 calls the method "WebGL2RenderingContext.texImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexImage3D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexImage3D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage3D returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasTexSubImage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage3D(
		this.Ref(),
	)
}

// TexSubImage3DFunc returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// TexSubImage3D calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage3D calls the method "WebGL2RenderingContext.texSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage3D1 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasTexSubImage3D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage3D1(
		this.Ref(),
	)
}

// TexSubImage3D1Func returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D1Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D1 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage3D1 calls the method "WebGL2RenderingContext.texSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage3D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage3D2 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasTexSubImage3D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage3D2(
		this.Ref(),
	)
}

// TexSubImage3D2Func returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D2Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D2 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage3D2 calls the method "WebGL2RenderingContext.texSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage3D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasTexSubImage3D3 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasTexSubImage3D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexSubImage3D3(
		this.Ref(),
	)
}

// TexSubImage3D3Func returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexSubImage3D3Func(
			this.Ref(),
		),
	)
}

// TexSubImage3D3 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryTexSubImage3D3 calls the method "WebGL2RenderingContext.texSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexSubImage3D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCopyTexSubImage3D returns true if the method "WebGL2RenderingContext.copyTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasCopyTexSubImage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCopyTexSubImage3D(
		this.Ref(),
	)
}

// CopyTexSubImage3DFunc returns the method "WebGL2RenderingContext.copyTexSubImage3D".
func (this WebGL2RenderingContext) CopyTexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// CopyTexSubImage3D calls the method "WebGL2RenderingContext.copyTexSubImage3D".
func (this WebGL2RenderingContext) CopyTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexSubImage3D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCopyTexSubImage3D calls the method "WebGL2RenderingContext.copyTexSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCopyTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCopyTexSubImage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage3D returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage3D(
		this.Ref(),
	)
}

// CompressedTexImage3DFunc returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexImage3D calls the method "WebGL2RenderingContext.compressedTexImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage3D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage3D1 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage3D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage3D1(
		this.Ref(),
	)
}

// CompressedTexImage3D1Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D1Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D1 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexImage3D1 calls the method "WebGL2RenderingContext.compressedTexImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage3D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage3D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage3D2 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage3D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage3D2(
		this.Ref(),
	)
}

// CompressedTexImage3D2Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D2Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D2 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexImage3D2 calls the method "WebGL2RenderingContext.compressedTexImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage3D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage3D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexImage3D3 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexImage3D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexImage3D3(
		this.Ref(),
	)
}

// CompressedTexImage3D3Func returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D3Func() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexImage3D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexImage3D3 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D3(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		srcData.Ref(),
	)

	return
}

// TryCompressedTexImage3D3 calls the method "WebGL2RenderingContext.compressedTexImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexImage3D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexImage3D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
		int32(border),
		srcData.Ref(),
	)

	return
}

// HasCompressedTexSubImage3D returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage3D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage3D(
		this.Ref(),
	)
}

// CompressedTexSubImage3DFunc returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3DFunc(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage3D calls the method "WebGL2RenderingContext.compressedTexSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage3D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage3D1 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage3D1() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage3D1(
		this.Ref(),
	)
}

// CompressedTexSubImage3D1Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D1Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D1Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D1 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage3D1 calls the method "WebGL2RenderingContext.compressedTexSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage3D1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage3D2 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage3D2() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage3D2(
		this.Ref(),
	)
}

// CompressedTexSubImage3D2Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D2Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D2Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D2 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage3D2 calls the method "WebGL2RenderingContext.compressedTexSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage3D2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasCompressedTexSubImage3D3 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasCompressedTexSubImage3D3() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompressedTexSubImage3D3(
		this.Ref(),
	)
}

// CompressedTexSubImage3D3Func returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D3Func() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompressedTexSubImage3D3Func(
			this.Ref(),
		),
	)
}

// CompressedTexSubImage3D3 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryCompressedTexSubImage3D3 calls the method "WebGL2RenderingContext.compressedTexSubImage3D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompressedTexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompressedTexSubImage3D3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasGetFragDataLocation returns true if the method "WebGL2RenderingContext.getFragDataLocation" exists.
func (this WebGL2RenderingContext) HasGetFragDataLocation() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetFragDataLocation(
		this.Ref(),
	)
}

// GetFragDataLocationFunc returns the method "WebGL2RenderingContext.getFragDataLocation".
func (this WebGL2RenderingContext) GetFragDataLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetFragDataLocationFunc(
			this.Ref(),
		),
	)
}

// GetFragDataLocation calls the method "WebGL2RenderingContext.getFragDataLocation".
func (this WebGL2RenderingContext) GetFragDataLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGL2RenderingContextGetFragDataLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		name.Ref(),
	)

	return
}

// TryGetFragDataLocation calls the method "WebGL2RenderingContext.getFragDataLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetFragDataLocation(program WebGLProgram, name js.String) (ret GLint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetFragDataLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasUniform1ui returns true if the method "WebGL2RenderingContext.uniform1ui" exists.
func (this WebGL2RenderingContext) HasUniform1ui() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1ui(
		this.Ref(),
	)
}

// Uniform1uiFunc returns the method "WebGL2RenderingContext.uniform1ui".
func (this WebGL2RenderingContext) Uniform1uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiFunc(
			this.Ref(),
		),
	)
}

// Uniform1ui calls the method "WebGL2RenderingContext.uniform1ui".
func (this WebGL2RenderingContext) Uniform1ui(location WebGLUniformLocation, v0 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1ui(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		uint32(v0),
	)

	return
}

// TryUniform1ui calls the method "WebGL2RenderingContext.uniform1ui"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1ui(location WebGLUniformLocation, v0 GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1ui(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
	)

	return
}

// HasUniform2ui returns true if the method "WebGL2RenderingContext.uniform2ui" exists.
func (this WebGL2RenderingContext) HasUniform2ui() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2ui(
		this.Ref(),
	)
}

// Uniform2uiFunc returns the method "WebGL2RenderingContext.uniform2ui".
func (this WebGL2RenderingContext) Uniform2uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiFunc(
			this.Ref(),
		),
	)
}

// Uniform2ui calls the method "WebGL2RenderingContext.uniform2ui".
func (this WebGL2RenderingContext) Uniform2ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2ui(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		uint32(v0),
		uint32(v1),
	)

	return
}

// TryUniform2ui calls the method "WebGL2RenderingContext.uniform2ui"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2ui(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
	)

	return
}

// HasUniform3ui returns true if the method "WebGL2RenderingContext.uniform3ui" exists.
func (this WebGL2RenderingContext) HasUniform3ui() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3ui(
		this.Ref(),
	)
}

// Uniform3uiFunc returns the method "WebGL2RenderingContext.uniform3ui".
func (this WebGL2RenderingContext) Uniform3uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiFunc(
			this.Ref(),
		),
	)
}

// Uniform3ui calls the method "WebGL2RenderingContext.uniform3ui".
func (this WebGL2RenderingContext) Uniform3ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3ui(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
	)

	return
}

// TryUniform3ui calls the method "WebGL2RenderingContext.uniform3ui"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3ui(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
	)

	return
}

// HasUniform4ui returns true if the method "WebGL2RenderingContext.uniform4ui" exists.
func (this WebGL2RenderingContext) HasUniform4ui() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4ui(
		this.Ref(),
	)
}

// Uniform4uiFunc returns the method "WebGL2RenderingContext.uniform4ui".
func (this WebGL2RenderingContext) Uniform4uiFunc() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiFunc(
			this.Ref(),
		),
	)
}

// Uniform4ui calls the method "WebGL2RenderingContext.uniform4ui".
func (this WebGL2RenderingContext) Uniform4ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4ui(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
		uint32(v3),
	)

	return
}

// TryUniform4ui calls the method "WebGL2RenderingContext.uniform4ui"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4ui(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
		uint32(v3),
	)

	return
}

// HasUniform1uiv returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasUniform1uiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1uiv(
		this.Ref(),
	)
}

// Uniform1uivFunc returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uivFunc(
			this.Ref(),
		),
	)
}

// Uniform1uiv calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform1uiv calls the method "WebGL2RenderingContext.uniform1uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1uiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform1uiv1 returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasUniform1uiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1uiv1(
		this.Ref(),
	)
}

// Uniform1uiv1Func returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform1uiv1 calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform1uiv1 calls the method "WebGL2RenderingContext.uniform1uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1uiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform1uiv2 returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasUniform1uiv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1uiv2(
		this.Ref(),
	)
}

// Uniform1uiv2Func returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform1uiv2 calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform1uiv2 calls the method "WebGL2RenderingContext.uniform1uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1uiv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform2uiv returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasUniform2uiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2uiv(
		this.Ref(),
	)
}

// Uniform2uivFunc returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uivFunc(
			this.Ref(),
		),
	)
}

// Uniform2uiv calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform2uiv calls the method "WebGL2RenderingContext.uniform2uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2uiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform2uiv1 returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasUniform2uiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2uiv1(
		this.Ref(),
	)
}

// Uniform2uiv1Func returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform2uiv1 calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform2uiv1 calls the method "WebGL2RenderingContext.uniform2uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2uiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform2uiv2 returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasUniform2uiv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2uiv2(
		this.Ref(),
	)
}

// Uniform2uiv2Func returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform2uiv2 calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform2uiv2 calls the method "WebGL2RenderingContext.uniform2uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2uiv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform3uiv returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasUniform3uiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3uiv(
		this.Ref(),
	)
}

// Uniform3uivFunc returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uivFunc(
			this.Ref(),
		),
	)
}

// Uniform3uiv calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform3uiv calls the method "WebGL2RenderingContext.uniform3uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3uiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform3uiv1 returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasUniform3uiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3uiv1(
		this.Ref(),
	)
}

// Uniform3uiv1Func returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform3uiv1 calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform3uiv1 calls the method "WebGL2RenderingContext.uniform3uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3uiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform3uiv2 returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasUniform3uiv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3uiv2(
		this.Ref(),
	)
}

// Uniform3uiv2Func returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform3uiv2 calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform3uiv2 calls the method "WebGL2RenderingContext.uniform3uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3uiv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniform4uiv returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasUniform4uiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4uiv(
		this.Ref(),
	)
}

// Uniform4uivFunc returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uivFunc() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uivFunc(
			this.Ref(),
		),
	)
}

// Uniform4uiv calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniform4uiv calls the method "WebGL2RenderingContext.uniform4uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4uiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniform4uiv1 returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasUniform4uiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4uiv1(
		this.Ref(),
	)
}

// Uniform4uiv1Func returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv1Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiv1Func(
			this.Ref(),
		),
	)
}

// Uniform4uiv1 calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniform4uiv1 calls the method "WebGL2RenderingContext.uniform4uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4uiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniform4uiv2 returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasUniform4uiv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4uiv2(
		this.Ref(),
	)
}

// Uniform4uiv2Func returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv2Func() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4uiv2Func(
			this.Ref(),
		),
	)
}

// Uniform4uiv2 calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		data.Ref(),
	)

	return
}

// TryUniform4uiv2 calls the method "WebGL2RenderingContext.uniform4uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4uiv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasUniformMatrix3x2fv returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x2fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x2fv(
		this.Ref(),
	)
}

// UniformMatrix3x2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix3x2fv calls the method "WebGL2RenderingContext.uniformMatrix3x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix3x2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x2fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x2fv1(
		this.Ref(),
	)
}

// UniformMatrix3x2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix3x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x2fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix3x2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x2fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x2fv2(
		this.Ref(),
	)
}

// UniformMatrix3x2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix3x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x2fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix4x2fv returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x2fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x2fv(
		this.Ref(),
	)
}

// UniformMatrix4x2fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix4x2fv calls the method "WebGL2RenderingContext.uniformMatrix4x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix4x2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x2fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x2fv1(
		this.Ref(),
	)
}

// UniformMatrix4x2fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix4x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x2fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix4x2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x2fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x2fv2(
		this.Ref(),
	)
}

// UniformMatrix4x2fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x2fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix4x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x2fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix2x3fv returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x3fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x3fv(
		this.Ref(),
	)
}

// UniformMatrix2x3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix2x3fv calls the method "WebGL2RenderingContext.uniformMatrix2x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix2x3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x3fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x3fv1(
		this.Ref(),
	)
}

// UniformMatrix2x3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix2x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x3fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix2x3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x3fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x3fv2(
		this.Ref(),
	)
}

// UniformMatrix2x3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix2x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x3fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix4x3fv returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x3fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x3fv(
		this.Ref(),
	)
}

// UniformMatrix4x3fvFunc returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix4x3fv calls the method "WebGL2RenderingContext.uniformMatrix4x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix4x3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x3fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x3fv1(
		this.Ref(),
	)
}

// UniformMatrix4x3fv1Func returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix4x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x3fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix4x3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix4x3fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix4x3fv2(
		this.Ref(),
	)
}

// UniformMatrix4x3fv2Func returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix4x3fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix4x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix4x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix4x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix4x3fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix2x4fv returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x4fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x4fv(
		this.Ref(),
	)
}

// UniformMatrix2x4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix2x4fv calls the method "WebGL2RenderingContext.uniformMatrix2x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix2x4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x4fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x4fv1(
		this.Ref(),
	)
}

// UniformMatrix2x4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix2x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x4fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix2x4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix2x4fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix2x4fv2(
		this.Ref(),
	)
}

// UniformMatrix2x4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix2x4fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix2x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix2x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix2x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix2x4fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasUniformMatrix3x4fv returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x4fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x4fv(
		this.Ref(),
	)
}

// UniformMatrix3x4fvFunc returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fvFunc() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fvFunc(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// TryUniformMatrix3x4fv calls the method "WebGL2RenderingContext.uniformMatrix3x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasUniformMatrix3x4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x4fv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x4fv1(
		this.Ref(),
	)
}

// UniformMatrix3x4fv1Func returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv1Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fv1Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv1(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryUniformMatrix3x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x4fv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasUniformMatrix3x4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasUniformMatrix3x4fv2() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformMatrix3x4fv2(
		this.Ref(),
	)
}

// UniformMatrix3x4fv2Func returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv2Func() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformMatrix3x4fv2Func(
			this.Ref(),
		),
	)
}

// UniformMatrix3x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv2(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// TryUniformMatrix3x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformMatrix3x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformMatrix3x4fv2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasVertexAttribI4i returns true if the method "WebGL2RenderingContext.vertexAttribI4i" exists.
func (this WebGL2RenderingContext) HasVertexAttribI4i() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribI4i(
		this.Ref(),
	)
}

// VertexAttribI4iFunc returns the method "WebGL2RenderingContext.vertexAttribI4i".
func (this WebGL2RenderingContext) VertexAttribI4iFunc() (fn js.Func[func(index GLuint, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4iFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4i calls the method "WebGL2RenderingContext.vertexAttribI4i".
func (this WebGL2RenderingContext) VertexAttribI4i(index GLuint, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4i(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// TryVertexAttribI4i calls the method "WebGL2RenderingContext.vertexAttribI4i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribI4i(index GLuint, x GLint, y GLint, z GLint, w GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribI4i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasVertexAttribI4iv returns true if the method "WebGL2RenderingContext.vertexAttribI4iv" exists.
func (this WebGL2RenderingContext) HasVertexAttribI4iv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribI4iv(
		this.Ref(),
	)
}

// VertexAttribI4ivFunc returns the method "WebGL2RenderingContext.vertexAttribI4iv".
func (this WebGL2RenderingContext) VertexAttribI4ivFunc() (fn js.Func[func(index GLuint, values Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4ivFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4iv calls the method "WebGL2RenderingContext.vertexAttribI4iv".
func (this WebGL2RenderingContext) VertexAttribI4iv(index GLuint, values Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4iv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttribI4iv calls the method "WebGL2RenderingContext.vertexAttribI4iv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribI4iv(index GLuint, values Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribI4iv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttribI4ui returns true if the method "WebGL2RenderingContext.vertexAttribI4ui" exists.
func (this WebGL2RenderingContext) HasVertexAttribI4ui() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribI4ui(
		this.Ref(),
	)
}

// VertexAttribI4uiFunc returns the method "WebGL2RenderingContext.vertexAttribI4ui".
func (this WebGL2RenderingContext) VertexAttribI4uiFunc() (fn js.Func[func(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4uiFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4ui calls the method "WebGL2RenderingContext.vertexAttribI4ui".
func (this WebGL2RenderingContext) VertexAttribI4ui(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4ui(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(x),
		uint32(y),
		uint32(z),
		uint32(w),
	)

	return
}

// TryVertexAttribI4ui calls the method "WebGL2RenderingContext.vertexAttribI4ui"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribI4ui(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribI4ui(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(x),
		uint32(y),
		uint32(z),
		uint32(w),
	)

	return
}

// HasVertexAttribI4uiv returns true if the method "WebGL2RenderingContext.vertexAttribI4uiv" exists.
func (this WebGL2RenderingContext) HasVertexAttribI4uiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribI4uiv(
		this.Ref(),
	)
}

// VertexAttribI4uivFunc returns the method "WebGL2RenderingContext.vertexAttribI4uiv".
func (this WebGL2RenderingContext) VertexAttribI4uivFunc() (fn js.Func[func(index GLuint, values Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribI4uivFunc(
			this.Ref(),
		),
	)
}

// VertexAttribI4uiv calls the method "WebGL2RenderingContext.vertexAttribI4uiv".
func (this WebGL2RenderingContext) VertexAttribI4uiv(index GLuint, values Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4uiv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttribI4uiv calls the method "WebGL2RenderingContext.vertexAttribI4uiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribI4uiv(index GLuint, values Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribI4uiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttribIPointer returns true if the method "WebGL2RenderingContext.vertexAttribIPointer" exists.
func (this WebGL2RenderingContext) HasVertexAttribIPointer() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribIPointer(
		this.Ref(),
	)
}

// VertexAttribIPointerFunc returns the method "WebGL2RenderingContext.vertexAttribIPointer".
func (this WebGL2RenderingContext) VertexAttribIPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribIPointerFunc(
			this.Ref(),
		),
	)
}

// VertexAttribIPointer calls the method "WebGL2RenderingContext.vertexAttribIPointer".
func (this WebGL2RenderingContext) VertexAttribIPointer(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribIPointer(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		int32(size),
		uint32(typ),
		int32(stride),
		float64(offset),
	)

	return
}

// TryVertexAttribIPointer calls the method "WebGL2RenderingContext.vertexAttribIPointer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribIPointer(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribIPointer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		int32(stride),
		float64(offset),
	)

	return
}

// HasVertexAttribDivisor returns true if the method "WebGL2RenderingContext.vertexAttribDivisor" exists.
func (this WebGL2RenderingContext) HasVertexAttribDivisor() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribDivisor(
		this.Ref(),
	)
}

// VertexAttribDivisorFunc returns the method "WebGL2RenderingContext.vertexAttribDivisor".
func (this WebGL2RenderingContext) VertexAttribDivisorFunc() (fn js.Func[func(index GLuint, divisor GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribDivisorFunc(
			this.Ref(),
		),
	)
}

// VertexAttribDivisor calls the method "WebGL2RenderingContext.vertexAttribDivisor".
func (this WebGL2RenderingContext) VertexAttribDivisor(index GLuint, divisor GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribDivisor(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(divisor),
	)

	return
}

// TryVertexAttribDivisor calls the method "WebGL2RenderingContext.vertexAttribDivisor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribDivisor(index GLuint, divisor GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribDivisor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(divisor),
	)

	return
}

// HasDrawArraysInstanced returns true if the method "WebGL2RenderingContext.drawArraysInstanced" exists.
func (this WebGL2RenderingContext) HasDrawArraysInstanced() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawArraysInstanced(
		this.Ref(),
	)
}

// DrawArraysInstancedFunc returns the method "WebGL2RenderingContext.drawArraysInstanced".
func (this WebGL2RenderingContext) DrawArraysInstancedFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawArraysInstancedFunc(
			this.Ref(),
		),
	)
}

// DrawArraysInstanced calls the method "WebGL2RenderingContext.drawArraysInstanced".
func (this WebGL2RenderingContext) DrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawArraysInstanced(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
	)

	return
}

// TryDrawArraysInstanced calls the method "WebGL2RenderingContext.drawArraysInstanced"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawArraysInstanced(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
	)

	return
}

// HasDrawElementsInstanced returns true if the method "WebGL2RenderingContext.drawElementsInstanced" exists.
func (this WebGL2RenderingContext) HasDrawElementsInstanced() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawElementsInstanced(
		this.Ref(),
	)
}

// DrawElementsInstancedFunc returns the method "WebGL2RenderingContext.drawElementsInstanced".
func (this WebGL2RenderingContext) DrawElementsInstancedFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawElementsInstancedFunc(
			this.Ref(),
		),
	)
}

// DrawElementsInstanced calls the method "WebGL2RenderingContext.drawElementsInstanced".
func (this WebGL2RenderingContext) DrawElementsInstanced(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawElementsInstanced(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
	)

	return
}

// TryDrawElementsInstanced calls the method "WebGL2RenderingContext.drawElementsInstanced"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawElementsInstanced(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawElementsInstanced(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
	)

	return
}

// HasDrawRangeElements returns true if the method "WebGL2RenderingContext.drawRangeElements" exists.
func (this WebGL2RenderingContext) HasDrawRangeElements() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawRangeElements(
		this.Ref(),
	)
}

// DrawRangeElementsFunc returns the method "WebGL2RenderingContext.drawRangeElements".
func (this WebGL2RenderingContext) DrawRangeElementsFunc() (fn js.Func[func(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawRangeElementsFunc(
			this.Ref(),
		),
	)
}

// DrawRangeElements calls the method "WebGL2RenderingContext.drawRangeElements".
func (this WebGL2RenderingContext) DrawRangeElements(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawRangeElements(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		uint32(start),
		uint32(end),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// TryDrawRangeElements calls the method "WebGL2RenderingContext.drawRangeElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawRangeElements(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawRangeElements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		uint32(start),
		uint32(end),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasDrawBuffers returns true if the method "WebGL2RenderingContext.drawBuffers" exists.
func (this WebGL2RenderingContext) HasDrawBuffers() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawBuffers(
		this.Ref(),
	)
}

// DrawBuffersFunc returns the method "WebGL2RenderingContext.drawBuffers".
func (this WebGL2RenderingContext) DrawBuffersFunc() (fn js.Func[func(buffers js.Array[GLenum])]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawBuffersFunc(
			this.Ref(),
		),
	)
}

// DrawBuffers calls the method "WebGL2RenderingContext.drawBuffers".
func (this WebGL2RenderingContext) DrawBuffers(buffers js.Array[GLenum]) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawBuffers(
		this.Ref(), js.Pointer(&ret),
		buffers.Ref(),
	)

	return
}

// TryDrawBuffers calls the method "WebGL2RenderingContext.drawBuffers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawBuffers(buffers js.Array[GLenum]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawBuffers(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffers.Ref(),
	)

	return
}

// HasClearBufferfv returns true if the method "WebGL2RenderingContext.clearBufferfv" exists.
func (this WebGL2RenderingContext) HasClearBufferfv() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferfv(
		this.Ref(),
	)
}

// ClearBufferfvFunc returns the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfvFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfvFunc(
			this.Ref(),
		),
	)
}

// ClearBufferfv calls the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfv(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfv(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryClearBufferfv calls the method "WebGL2RenderingContext.clearBufferfv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferfv(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferfv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasClearBufferfv1 returns true if the method "WebGL2RenderingContext.clearBufferfv" exists.
func (this WebGL2RenderingContext) HasClearBufferfv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferfv1(
		this.Ref(),
	)
}

// ClearBufferfv1Func returns the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferfv1 calls the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfv1(buffer GLenum, drawbuffer GLint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfv1(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// TryClearBufferfv1 calls the method "WebGL2RenderingContext.clearBufferfv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferfv1(buffer GLenum, drawbuffer GLint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferfv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasClearBufferiv returns true if the method "WebGL2RenderingContext.clearBufferiv" exists.
func (this WebGL2RenderingContext) HasClearBufferiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferiv(
		this.Ref(),
	)
}

// ClearBufferivFunc returns the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferivFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferivFunc(
			this.Ref(),
		),
	)
}

// ClearBufferiv calls the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferiv(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferiv(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryClearBufferiv calls the method "WebGL2RenderingContext.clearBufferiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferiv(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasClearBufferiv1 returns true if the method "WebGL2RenderingContext.clearBufferiv" exists.
func (this WebGL2RenderingContext) HasClearBufferiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferiv1(
		this.Ref(),
	)
}

// ClearBufferiv1Func returns the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferiv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferiv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferiv1 calls the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferiv1(buffer GLenum, drawbuffer GLint, values Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferiv1(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// TryClearBufferiv1 calls the method "WebGL2RenderingContext.clearBufferiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferiv1(buffer GLenum, drawbuffer GLint, values Int32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasClearBufferuiv returns true if the method "WebGL2RenderingContext.clearBufferuiv" exists.
func (this WebGL2RenderingContext) HasClearBufferuiv() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferuiv(
		this.Ref(),
	)
}

// ClearBufferuivFunc returns the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuivFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferuivFunc(
			this.Ref(),
		),
	)
}

// ClearBufferuiv calls the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuiv(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferuiv(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// TryClearBufferuiv calls the method "WebGL2RenderingContext.clearBufferuiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferuiv(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferuiv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasClearBufferuiv1 returns true if the method "WebGL2RenderingContext.clearBufferuiv" exists.
func (this WebGL2RenderingContext) HasClearBufferuiv1() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferuiv1(
		this.Ref(),
	)
}

// ClearBufferuiv1Func returns the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuiv1Func() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferuiv1Func(
			this.Ref(),
		),
	)
}

// ClearBufferuiv1 calls the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuiv1(buffer GLenum, drawbuffer GLint, values Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferuiv1(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// TryClearBufferuiv1 calls the method "WebGL2RenderingContext.clearBufferuiv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferuiv1(buffer GLenum, drawbuffer GLint, values Uint32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferuiv1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasClearBufferfi returns true if the method "WebGL2RenderingContext.clearBufferfi" exists.
func (this WebGL2RenderingContext) HasClearBufferfi() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearBufferfi(
		this.Ref(),
	)
}

// ClearBufferfiFunc returns the method "WebGL2RenderingContext.clearBufferfi".
func (this WebGL2RenderingContext) ClearBufferfiFunc() (fn js.Func[func(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearBufferfiFunc(
			this.Ref(),
		),
	)
}

// ClearBufferfi calls the method "WebGL2RenderingContext.clearBufferfi".
func (this WebGL2RenderingContext) ClearBufferfi(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfi(
		this.Ref(), js.Pointer(&ret),
		uint32(buffer),
		int32(drawbuffer),
		float32(depth),
		int32(stencil),
	)

	return
}

// TryClearBufferfi calls the method "WebGL2RenderingContext.clearBufferfi"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearBufferfi(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearBufferfi(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		float32(depth),
		int32(stencil),
	)

	return
}

// HasCreateQuery returns true if the method "WebGL2RenderingContext.createQuery" exists.
func (this WebGL2RenderingContext) HasCreateQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateQuery(
		this.Ref(),
	)
}

// CreateQueryFunc returns the method "WebGL2RenderingContext.createQuery".
func (this WebGL2RenderingContext) CreateQueryFunc() (fn js.Func[func() WebGLQuery]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateQueryFunc(
			this.Ref(),
		),
	)
}

// CreateQuery calls the method "WebGL2RenderingContext.createQuery".
func (this WebGL2RenderingContext) CreateQuery() (ret WebGLQuery) {
	bindings.CallWebGL2RenderingContextCreateQuery(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateQuery calls the method "WebGL2RenderingContext.createQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateQuery() (ret WebGLQuery, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteQuery returns true if the method "WebGL2RenderingContext.deleteQuery" exists.
func (this WebGL2RenderingContext) HasDeleteQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteQuery(
		this.Ref(),
	)
}

// DeleteQueryFunc returns the method "WebGL2RenderingContext.deleteQuery".
func (this WebGL2RenderingContext) DeleteQueryFunc() (fn js.Func[func(query WebGLQuery)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteQueryFunc(
			this.Ref(),
		),
	)
}

// DeleteQuery calls the method "WebGL2RenderingContext.deleteQuery".
func (this WebGL2RenderingContext) DeleteQuery(query WebGLQuery) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteQuery(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryDeleteQuery calls the method "WebGL2RenderingContext.deleteQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteQuery(query WebGLQuery) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasIsQuery returns true if the method "WebGL2RenderingContext.isQuery" exists.
func (this WebGL2RenderingContext) HasIsQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsQuery(
		this.Ref(),
	)
}

// IsQueryFunc returns the method "WebGL2RenderingContext.isQuery".
func (this WebGL2RenderingContext) IsQueryFunc() (fn js.Func[func(query WebGLQuery) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsQueryFunc(
			this.Ref(),
		),
	)
}

// IsQuery calls the method "WebGL2RenderingContext.isQuery".
func (this WebGL2RenderingContext) IsQuery(query WebGLQuery) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsQuery(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryIsQuery calls the method "WebGL2RenderingContext.isQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsQuery(query WebGLQuery) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasBeginQuery returns true if the method "WebGL2RenderingContext.beginQuery" exists.
func (this WebGL2RenderingContext) HasBeginQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextBeginQuery(
		this.Ref(),
	)
}

// BeginQueryFunc returns the method "WebGL2RenderingContext.beginQuery".
func (this WebGL2RenderingContext) BeginQueryFunc() (fn js.Func[func(target GLenum, query WebGLQuery)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBeginQueryFunc(
			this.Ref(),
		),
	)
}

// BeginQuery calls the method "WebGL2RenderingContext.beginQuery".
func (this WebGL2RenderingContext) BeginQuery(target GLenum, query WebGLQuery) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBeginQuery(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		query.Ref(),
	)

	return
}

// TryBeginQuery calls the method "WebGL2RenderingContext.beginQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBeginQuery(target GLenum, query WebGLQuery) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBeginQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		query.Ref(),
	)

	return
}

// HasEndQuery returns true if the method "WebGL2RenderingContext.endQuery" exists.
func (this WebGL2RenderingContext) HasEndQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextEndQuery(
		this.Ref(),
	)
}

// EndQueryFunc returns the method "WebGL2RenderingContext.endQuery".
func (this WebGL2RenderingContext) EndQueryFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEndQueryFunc(
			this.Ref(),
		),
	)
}

// EndQuery calls the method "WebGL2RenderingContext.endQuery".
func (this WebGL2RenderingContext) EndQuery(target GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEndQuery(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryEndQuery calls the method "WebGL2RenderingContext.endQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEndQuery(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEndQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasGetQuery returns true if the method "WebGL2RenderingContext.getQuery" exists.
func (this WebGL2RenderingContext) HasGetQuery() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetQuery(
		this.Ref(),
	)
}

// GetQueryFunc returns the method "WebGL2RenderingContext.getQuery".
func (this WebGL2RenderingContext) GetQueryFunc() (fn js.Func[func(target GLenum, pname GLenum) WebGLQuery]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetQueryFunc(
			this.Ref(),
		),
	)
}

// GetQuery calls the method "WebGL2RenderingContext.getQuery".
func (this WebGL2RenderingContext) GetQuery(target GLenum, pname GLenum) (ret WebGLQuery) {
	bindings.CallWebGL2RenderingContextGetQuery(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetQuery calls the method "WebGL2RenderingContext.getQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetQuery(target GLenum, pname GLenum) (ret WebGLQuery, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetQueryParameter returns true if the method "WebGL2RenderingContext.getQueryParameter" exists.
func (this WebGL2RenderingContext) HasGetQueryParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetQueryParameter(
		this.Ref(),
	)
}

// GetQueryParameterFunc returns the method "WebGL2RenderingContext.getQueryParameter".
func (this WebGL2RenderingContext) GetQueryParameterFunc() (fn js.Func[func(query WebGLQuery, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetQueryParameterFunc(
			this.Ref(),
		),
	)
}

// GetQueryParameter calls the method "WebGL2RenderingContext.getQueryParameter".
func (this WebGL2RenderingContext) GetQueryParameter(query WebGLQuery, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetQueryParameter(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(pname),
	)

	return
}

// TryGetQueryParameter calls the method "WebGL2RenderingContext.getQueryParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetQueryParameter(query WebGLQuery, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetQueryParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(pname),
	)

	return
}

// HasCreateSampler returns true if the method "WebGL2RenderingContext.createSampler" exists.
func (this WebGL2RenderingContext) HasCreateSampler() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateSampler(
		this.Ref(),
	)
}

// CreateSamplerFunc returns the method "WebGL2RenderingContext.createSampler".
func (this WebGL2RenderingContext) CreateSamplerFunc() (fn js.Func[func() WebGLSampler]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateSamplerFunc(
			this.Ref(),
		),
	)
}

// CreateSampler calls the method "WebGL2RenderingContext.createSampler".
func (this WebGL2RenderingContext) CreateSampler() (ret WebGLSampler) {
	bindings.CallWebGL2RenderingContextCreateSampler(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSampler calls the method "WebGL2RenderingContext.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateSampler() (ret WebGLSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateSampler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteSampler returns true if the method "WebGL2RenderingContext.deleteSampler" exists.
func (this WebGL2RenderingContext) HasDeleteSampler() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteSampler(
		this.Ref(),
	)
}

// DeleteSamplerFunc returns the method "WebGL2RenderingContext.deleteSampler".
func (this WebGL2RenderingContext) DeleteSamplerFunc() (fn js.Func[func(sampler WebGLSampler)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteSamplerFunc(
			this.Ref(),
		),
	)
}

// DeleteSampler calls the method "WebGL2RenderingContext.deleteSampler".
func (this WebGL2RenderingContext) DeleteSampler(sampler WebGLSampler) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteSampler(
		this.Ref(), js.Pointer(&ret),
		sampler.Ref(),
	)

	return
}

// TryDeleteSampler calls the method "WebGL2RenderingContext.deleteSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteSampler(sampler WebGLSampler) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteSampler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
	)

	return
}

// HasIsSampler returns true if the method "WebGL2RenderingContext.isSampler" exists.
func (this WebGL2RenderingContext) HasIsSampler() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsSampler(
		this.Ref(),
	)
}

// IsSamplerFunc returns the method "WebGL2RenderingContext.isSampler".
func (this WebGL2RenderingContext) IsSamplerFunc() (fn js.Func[func(sampler WebGLSampler) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsSamplerFunc(
			this.Ref(),
		),
	)
}

// IsSampler calls the method "WebGL2RenderingContext.isSampler".
func (this WebGL2RenderingContext) IsSampler(sampler WebGLSampler) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsSampler(
		this.Ref(), js.Pointer(&ret),
		sampler.Ref(),
	)

	return
}

// TryIsSampler calls the method "WebGL2RenderingContext.isSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsSampler(sampler WebGLSampler) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsSampler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
	)

	return
}

// HasBindSampler returns true if the method "WebGL2RenderingContext.bindSampler" exists.
func (this WebGL2RenderingContext) HasBindSampler() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindSampler(
		this.Ref(),
	)
}

// BindSamplerFunc returns the method "WebGL2RenderingContext.bindSampler".
func (this WebGL2RenderingContext) BindSamplerFunc() (fn js.Func[func(unit GLuint, sampler WebGLSampler)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindSamplerFunc(
			this.Ref(),
		),
	)
}

// BindSampler calls the method "WebGL2RenderingContext.bindSampler".
func (this WebGL2RenderingContext) BindSampler(unit GLuint, sampler WebGLSampler) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindSampler(
		this.Ref(), js.Pointer(&ret),
		uint32(unit),
		sampler.Ref(),
	)

	return
}

// TryBindSampler calls the method "WebGL2RenderingContext.bindSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindSampler(unit GLuint, sampler WebGLSampler) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindSampler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(unit),
		sampler.Ref(),
	)

	return
}

// HasSamplerParameteri returns true if the method "WebGL2RenderingContext.samplerParameteri" exists.
func (this WebGL2RenderingContext) HasSamplerParameteri() bool {
	return js.True == bindings.HasWebGL2RenderingContextSamplerParameteri(
		this.Ref(),
	)
}

// SamplerParameteriFunc returns the method "WebGL2RenderingContext.samplerParameteri".
func (this WebGL2RenderingContext) SamplerParameteriFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSamplerParameteriFunc(
			this.Ref(),
		),
	)
}

// SamplerParameteri calls the method "WebGL2RenderingContext.samplerParameteri".
func (this WebGL2RenderingContext) SamplerParameteri(sampler WebGLSampler, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSamplerParameteri(
		this.Ref(), js.Pointer(&ret),
		sampler.Ref(),
		uint32(pname),
		int32(param),
	)

	return
}

// TrySamplerParameteri calls the method "WebGL2RenderingContext.samplerParameteri"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TrySamplerParameteri(sampler WebGLSampler, pname GLenum, param GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextSamplerParameteri(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
		int32(param),
	)

	return
}

// HasSamplerParameterf returns true if the method "WebGL2RenderingContext.samplerParameterf" exists.
func (this WebGL2RenderingContext) HasSamplerParameterf() bool {
	return js.True == bindings.HasWebGL2RenderingContextSamplerParameterf(
		this.Ref(),
	)
}

// SamplerParameterfFunc returns the method "WebGL2RenderingContext.samplerParameterf".
func (this WebGL2RenderingContext) SamplerParameterfFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSamplerParameterfFunc(
			this.Ref(),
		),
	)
}

// SamplerParameterf calls the method "WebGL2RenderingContext.samplerParameterf".
func (this WebGL2RenderingContext) SamplerParameterf(sampler WebGLSampler, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSamplerParameterf(
		this.Ref(), js.Pointer(&ret),
		sampler.Ref(),
		uint32(pname),
		float32(param),
	)

	return
}

// TrySamplerParameterf calls the method "WebGL2RenderingContext.samplerParameterf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TrySamplerParameterf(sampler WebGLSampler, pname GLenum, param GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextSamplerParameterf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
		float32(param),
	)

	return
}

// HasGetSamplerParameter returns true if the method "WebGL2RenderingContext.getSamplerParameter" exists.
func (this WebGL2RenderingContext) HasGetSamplerParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetSamplerParameter(
		this.Ref(),
	)
}

// GetSamplerParameterFunc returns the method "WebGL2RenderingContext.getSamplerParameter".
func (this WebGL2RenderingContext) GetSamplerParameterFunc() (fn js.Func[func(sampler WebGLSampler, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSamplerParameterFunc(
			this.Ref(),
		),
	)
}

// GetSamplerParameter calls the method "WebGL2RenderingContext.getSamplerParameter".
func (this WebGL2RenderingContext) GetSamplerParameter(sampler WebGLSampler, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetSamplerParameter(
		this.Ref(), js.Pointer(&ret),
		sampler.Ref(),
		uint32(pname),
	)

	return
}

// TryGetSamplerParameter calls the method "WebGL2RenderingContext.getSamplerParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetSamplerParameter(sampler WebGLSampler, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetSamplerParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
	)

	return
}

// HasFenceSync returns true if the method "WebGL2RenderingContext.fenceSync" exists.
func (this WebGL2RenderingContext) HasFenceSync() bool {
	return js.True == bindings.HasWebGL2RenderingContextFenceSync(
		this.Ref(),
	)
}

// FenceSyncFunc returns the method "WebGL2RenderingContext.fenceSync".
func (this WebGL2RenderingContext) FenceSyncFunc() (fn js.Func[func(condition GLenum, flags GLbitfield) WebGLSync]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFenceSyncFunc(
			this.Ref(),
		),
	)
}

// FenceSync calls the method "WebGL2RenderingContext.fenceSync".
func (this WebGL2RenderingContext) FenceSync(condition GLenum, flags GLbitfield) (ret WebGLSync) {
	bindings.CallWebGL2RenderingContextFenceSync(
		this.Ref(), js.Pointer(&ret),
		uint32(condition),
		uint32(flags),
	)

	return
}

// TryFenceSync calls the method "WebGL2RenderingContext.fenceSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFenceSync(condition GLenum, flags GLbitfield) (ret WebGLSync, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFenceSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(condition),
		uint32(flags),
	)

	return
}

// HasIsSync returns true if the method "WebGL2RenderingContext.isSync" exists.
func (this WebGL2RenderingContext) HasIsSync() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsSync(
		this.Ref(),
	)
}

// IsSyncFunc returns the method "WebGL2RenderingContext.isSync".
func (this WebGL2RenderingContext) IsSyncFunc() (fn js.Func[func(sync WebGLSync) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsSyncFunc(
			this.Ref(),
		),
	)
}

// IsSync calls the method "WebGL2RenderingContext.isSync".
func (this WebGL2RenderingContext) IsSync(sync WebGLSync) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsSync(
		this.Ref(), js.Pointer(&ret),
		sync.Ref(),
	)

	return
}

// TryIsSync calls the method "WebGL2RenderingContext.isSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsSync(sync WebGLSync) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
	)

	return
}

// HasDeleteSync returns true if the method "WebGL2RenderingContext.deleteSync" exists.
func (this WebGL2RenderingContext) HasDeleteSync() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteSync(
		this.Ref(),
	)
}

// DeleteSyncFunc returns the method "WebGL2RenderingContext.deleteSync".
func (this WebGL2RenderingContext) DeleteSyncFunc() (fn js.Func[func(sync WebGLSync)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteSyncFunc(
			this.Ref(),
		),
	)
}

// DeleteSync calls the method "WebGL2RenderingContext.deleteSync".
func (this WebGL2RenderingContext) DeleteSync(sync WebGLSync) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteSync(
		this.Ref(), js.Pointer(&ret),
		sync.Ref(),
	)

	return
}

// TryDeleteSync calls the method "WebGL2RenderingContext.deleteSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteSync(sync WebGLSync) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
	)

	return
}

// HasClientWaitSync returns true if the method "WebGL2RenderingContext.clientWaitSync" exists.
func (this WebGL2RenderingContext) HasClientWaitSync() bool {
	return js.True == bindings.HasWebGL2RenderingContextClientWaitSync(
		this.Ref(),
	)
}

// ClientWaitSyncFunc returns the method "WebGL2RenderingContext.clientWaitSync".
func (this WebGL2RenderingContext) ClientWaitSyncFunc() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLuint64) GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClientWaitSyncFunc(
			this.Ref(),
		),
	)
}

// ClientWaitSync calls the method "WebGL2RenderingContext.clientWaitSync".
func (this WebGL2RenderingContext) ClientWaitSync(sync WebGLSync, flags GLbitfield, timeout GLuint64) (ret GLenum) {
	bindings.CallWebGL2RenderingContextClientWaitSync(
		this.Ref(), js.Pointer(&ret),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// TryClientWaitSync calls the method "WebGL2RenderingContext.clientWaitSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClientWaitSync(sync WebGLSync, flags GLbitfield, timeout GLuint64) (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClientWaitSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// HasWaitSync returns true if the method "WebGL2RenderingContext.waitSync" exists.
func (this WebGL2RenderingContext) HasWaitSync() bool {
	return js.True == bindings.HasWebGL2RenderingContextWaitSync(
		this.Ref(),
	)
}

// WaitSyncFunc returns the method "WebGL2RenderingContext.waitSync".
func (this WebGL2RenderingContext) WaitSyncFunc() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLint64)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextWaitSyncFunc(
			this.Ref(),
		),
	)
}

// WaitSync calls the method "WebGL2RenderingContext.waitSync".
func (this WebGL2RenderingContext) WaitSync(sync WebGLSync, flags GLbitfield, timeout GLint64) (ret js.Void) {
	bindings.CallWebGL2RenderingContextWaitSync(
		this.Ref(), js.Pointer(&ret),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// TryWaitSync calls the method "WebGL2RenderingContext.waitSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryWaitSync(sync WebGLSync, flags GLbitfield, timeout GLint64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextWaitSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// HasGetSyncParameter returns true if the method "WebGL2RenderingContext.getSyncParameter" exists.
func (this WebGL2RenderingContext) HasGetSyncParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetSyncParameter(
		this.Ref(),
	)
}

// GetSyncParameterFunc returns the method "WebGL2RenderingContext.getSyncParameter".
func (this WebGL2RenderingContext) GetSyncParameterFunc() (fn js.Func[func(sync WebGLSync, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSyncParameterFunc(
			this.Ref(),
		),
	)
}

// GetSyncParameter calls the method "WebGL2RenderingContext.getSyncParameter".
func (this WebGL2RenderingContext) GetSyncParameter(sync WebGLSync, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetSyncParameter(
		this.Ref(), js.Pointer(&ret),
		sync.Ref(),
		uint32(pname),
	)

	return
}

// TryGetSyncParameter calls the method "WebGL2RenderingContext.getSyncParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetSyncParameter(sync WebGLSync, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetSyncParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(pname),
	)

	return
}

// HasCreateTransformFeedback returns true if the method "WebGL2RenderingContext.createTransformFeedback" exists.
func (this WebGL2RenderingContext) HasCreateTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateTransformFeedback(
		this.Ref(),
	)
}

// CreateTransformFeedbackFunc returns the method "WebGL2RenderingContext.createTransformFeedback".
func (this WebGL2RenderingContext) CreateTransformFeedbackFunc() (fn js.Func[func() WebGLTransformFeedback]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// CreateTransformFeedback calls the method "WebGL2RenderingContext.createTransformFeedback".
func (this WebGL2RenderingContext) CreateTransformFeedback() (ret WebGLTransformFeedback) {
	bindings.CallWebGL2RenderingContextCreateTransformFeedback(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTransformFeedback calls the method "WebGL2RenderingContext.createTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateTransformFeedback() (ret WebGLTransformFeedback, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteTransformFeedback returns true if the method "WebGL2RenderingContext.deleteTransformFeedback" exists.
func (this WebGL2RenderingContext) HasDeleteTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteTransformFeedback(
		this.Ref(),
	)
}

// DeleteTransformFeedbackFunc returns the method "WebGL2RenderingContext.deleteTransformFeedback".
func (this WebGL2RenderingContext) DeleteTransformFeedbackFunc() (fn js.Func[func(tf WebGLTransformFeedback)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// DeleteTransformFeedback calls the method "WebGL2RenderingContext.deleteTransformFeedback".
func (this WebGL2RenderingContext) DeleteTransformFeedback(tf WebGLTransformFeedback) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteTransformFeedback(
		this.Ref(), js.Pointer(&ret),
		tf.Ref(),
	)

	return
}

// TryDeleteTransformFeedback calls the method "WebGL2RenderingContext.deleteTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteTransformFeedback(tf WebGLTransformFeedback) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tf.Ref(),
	)

	return
}

// HasIsTransformFeedback returns true if the method "WebGL2RenderingContext.isTransformFeedback" exists.
func (this WebGL2RenderingContext) HasIsTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsTransformFeedback(
		this.Ref(),
	)
}

// IsTransformFeedbackFunc returns the method "WebGL2RenderingContext.isTransformFeedback".
func (this WebGL2RenderingContext) IsTransformFeedbackFunc() (fn js.Func[func(tf WebGLTransformFeedback) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// IsTransformFeedback calls the method "WebGL2RenderingContext.isTransformFeedback".
func (this WebGL2RenderingContext) IsTransformFeedback(tf WebGLTransformFeedback) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsTransformFeedback(
		this.Ref(), js.Pointer(&ret),
		tf.Ref(),
	)

	return
}

// TryIsTransformFeedback calls the method "WebGL2RenderingContext.isTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsTransformFeedback(tf WebGLTransformFeedback) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tf.Ref(),
	)

	return
}

// HasBindTransformFeedback returns true if the method "WebGL2RenderingContext.bindTransformFeedback" exists.
func (this WebGL2RenderingContext) HasBindTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindTransformFeedback(
		this.Ref(),
	)
}

// BindTransformFeedbackFunc returns the method "WebGL2RenderingContext.bindTransformFeedback".
func (this WebGL2RenderingContext) BindTransformFeedbackFunc() (fn js.Func[func(target GLenum, tf WebGLTransformFeedback)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// BindTransformFeedback calls the method "WebGL2RenderingContext.bindTransformFeedback".
func (this WebGL2RenderingContext) BindTransformFeedback(target GLenum, tf WebGLTransformFeedback) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindTransformFeedback(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		tf.Ref(),
	)

	return
}

// TryBindTransformFeedback calls the method "WebGL2RenderingContext.bindTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindTransformFeedback(target GLenum, tf WebGLTransformFeedback) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		tf.Ref(),
	)

	return
}

// HasBeginTransformFeedback returns true if the method "WebGL2RenderingContext.beginTransformFeedback" exists.
func (this WebGL2RenderingContext) HasBeginTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextBeginTransformFeedback(
		this.Ref(),
	)
}

// BeginTransformFeedbackFunc returns the method "WebGL2RenderingContext.beginTransformFeedback".
func (this WebGL2RenderingContext) BeginTransformFeedbackFunc() (fn js.Func[func(primitiveMode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBeginTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// BeginTransformFeedback calls the method "WebGL2RenderingContext.beginTransformFeedback".
func (this WebGL2RenderingContext) BeginTransformFeedback(primitiveMode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBeginTransformFeedback(
		this.Ref(), js.Pointer(&ret),
		uint32(primitiveMode),
	)

	return
}

// TryBeginTransformFeedback calls the method "WebGL2RenderingContext.beginTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBeginTransformFeedback(primitiveMode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBeginTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(primitiveMode),
	)

	return
}

// HasEndTransformFeedback returns true if the method "WebGL2RenderingContext.endTransformFeedback" exists.
func (this WebGL2RenderingContext) HasEndTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextEndTransformFeedback(
		this.Ref(),
	)
}

// EndTransformFeedbackFunc returns the method "WebGL2RenderingContext.endTransformFeedback".
func (this WebGL2RenderingContext) EndTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEndTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// EndTransformFeedback calls the method "WebGL2RenderingContext.endTransformFeedback".
func (this WebGL2RenderingContext) EndTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextEndTransformFeedback(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEndTransformFeedback calls the method "WebGL2RenderingContext.endTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEndTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEndTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTransformFeedbackVaryings returns true if the method "WebGL2RenderingContext.transformFeedbackVaryings" exists.
func (this WebGL2RenderingContext) HasTransformFeedbackVaryings() bool {
	return js.True == bindings.HasWebGL2RenderingContextTransformFeedbackVaryings(
		this.Ref(),
	)
}

// TransformFeedbackVaryingsFunc returns the method "WebGL2RenderingContext.transformFeedbackVaryings".
func (this WebGL2RenderingContext) TransformFeedbackVaryingsFunc() (fn js.Func[func(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTransformFeedbackVaryingsFunc(
			this.Ref(),
		),
	)
}

// TransformFeedbackVaryings calls the method "WebGL2RenderingContext.transformFeedbackVaryings".
func (this WebGL2RenderingContext) TransformFeedbackVaryings(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTransformFeedbackVaryings(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		varyings.Ref(),
		uint32(bufferMode),
	)

	return
}

// TryTransformFeedbackVaryings calls the method "WebGL2RenderingContext.transformFeedbackVaryings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTransformFeedbackVaryings(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTransformFeedbackVaryings(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		varyings.Ref(),
		uint32(bufferMode),
	)

	return
}

// HasGetTransformFeedbackVarying returns true if the method "WebGL2RenderingContext.getTransformFeedbackVarying" exists.
func (this WebGL2RenderingContext) HasGetTransformFeedbackVarying() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetTransformFeedbackVarying(
		this.Ref(),
	)
}

// GetTransformFeedbackVaryingFunc returns the method "WebGL2RenderingContext.getTransformFeedbackVarying".
func (this WebGL2RenderingContext) GetTransformFeedbackVaryingFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetTransformFeedbackVaryingFunc(
			this.Ref(),
		),
	)
}

// GetTransformFeedbackVarying calls the method "WebGL2RenderingContext.getTransformFeedbackVarying".
func (this WebGL2RenderingContext) GetTransformFeedbackVarying(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetTransformFeedbackVarying(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
	)

	return
}

// TryGetTransformFeedbackVarying calls the method "WebGL2RenderingContext.getTransformFeedbackVarying"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetTransformFeedbackVarying(program WebGLProgram, index GLuint) (ret WebGLActiveInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetTransformFeedbackVarying(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasPauseTransformFeedback returns true if the method "WebGL2RenderingContext.pauseTransformFeedback" exists.
func (this WebGL2RenderingContext) HasPauseTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextPauseTransformFeedback(
		this.Ref(),
	)
}

// PauseTransformFeedbackFunc returns the method "WebGL2RenderingContext.pauseTransformFeedback".
func (this WebGL2RenderingContext) PauseTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPauseTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// PauseTransformFeedback calls the method "WebGL2RenderingContext.pauseTransformFeedback".
func (this WebGL2RenderingContext) PauseTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextPauseTransformFeedback(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPauseTransformFeedback calls the method "WebGL2RenderingContext.pauseTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryPauseTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextPauseTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResumeTransformFeedback returns true if the method "WebGL2RenderingContext.resumeTransformFeedback" exists.
func (this WebGL2RenderingContext) HasResumeTransformFeedback() bool {
	return js.True == bindings.HasWebGL2RenderingContextResumeTransformFeedback(
		this.Ref(),
	)
}

// ResumeTransformFeedbackFunc returns the method "WebGL2RenderingContext.resumeTransformFeedback".
func (this WebGL2RenderingContext) ResumeTransformFeedbackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextResumeTransformFeedbackFunc(
			this.Ref(),
		),
	)
}

// ResumeTransformFeedback calls the method "WebGL2RenderingContext.resumeTransformFeedback".
func (this WebGL2RenderingContext) ResumeTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextResumeTransformFeedback(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResumeTransformFeedback calls the method "WebGL2RenderingContext.resumeTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryResumeTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextResumeTransformFeedback(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBindBufferBase returns true if the method "WebGL2RenderingContext.bindBufferBase" exists.
func (this WebGL2RenderingContext) HasBindBufferBase() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindBufferBase(
		this.Ref(),
	)
}

// BindBufferBaseFunc returns the method "WebGL2RenderingContext.bindBufferBase".
func (this WebGL2RenderingContext) BindBufferBaseFunc() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferBaseFunc(
			this.Ref(),
		),
	)
}

// BindBufferBase calls the method "WebGL2RenderingContext.bindBufferBase".
func (this WebGL2RenderingContext) BindBufferBase(target GLenum, index GLuint, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBufferBase(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(index),
		buffer.Ref(),
	)

	return
}

// TryBindBufferBase calls the method "WebGL2RenderingContext.bindBufferBase"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindBufferBase(target GLenum, index GLuint, buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindBufferBase(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
		buffer.Ref(),
	)

	return
}

// HasBindBufferRange returns true if the method "WebGL2RenderingContext.bindBufferRange" exists.
func (this WebGL2RenderingContext) HasBindBufferRange() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindBufferRange(
		this.Ref(),
	)
}

// BindBufferRangeFunc returns the method "WebGL2RenderingContext.bindBufferRange".
func (this WebGL2RenderingContext) BindBufferRangeFunc() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferRangeFunc(
			this.Ref(),
		),
	)
}

// BindBufferRange calls the method "WebGL2RenderingContext.bindBufferRange".
func (this WebGL2RenderingContext) BindBufferRange(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBufferRange(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(index),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// TryBindBufferRange calls the method "WebGL2RenderingContext.bindBufferRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindBufferRange(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindBufferRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasGetIndexedParameter returns true if the method "WebGL2RenderingContext.getIndexedParameter" exists.
func (this WebGL2RenderingContext) HasGetIndexedParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetIndexedParameter(
		this.Ref(),
	)
}

// GetIndexedParameterFunc returns the method "WebGL2RenderingContext.getIndexedParameter".
func (this WebGL2RenderingContext) GetIndexedParameterFunc() (fn js.Func[func(target GLenum, index GLuint) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetIndexedParameterFunc(
			this.Ref(),
		),
	)
}

// GetIndexedParameter calls the method "WebGL2RenderingContext.getIndexedParameter".
func (this WebGL2RenderingContext) GetIndexedParameter(target GLenum, index GLuint) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetIndexedParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(index),
	)

	return
}

// TryGetIndexedParameter calls the method "WebGL2RenderingContext.getIndexedParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetIndexedParameter(target GLenum, index GLuint) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetIndexedParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
	)

	return
}

// HasGetUniformIndices returns true if the method "WebGL2RenderingContext.getUniformIndices" exists.
func (this WebGL2RenderingContext) HasGetUniformIndices() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetUniformIndices(
		this.Ref(),
	)
}

// GetUniformIndicesFunc returns the method "WebGL2RenderingContext.getUniformIndices".
func (this WebGL2RenderingContext) GetUniformIndicesFunc() (fn js.Func[func(program WebGLProgram, uniformNames js.Array[js.String]) js.Array[GLuint]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformIndicesFunc(
			this.Ref(),
		),
	)
}

// GetUniformIndices calls the method "WebGL2RenderingContext.getUniformIndices".
func (this WebGL2RenderingContext) GetUniformIndices(program WebGLProgram, uniformNames js.Array[js.String]) (ret js.Array[GLuint]) {
	bindings.CallWebGL2RenderingContextGetUniformIndices(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uniformNames.Ref(),
	)

	return
}

// TryGetUniformIndices calls the method "WebGL2RenderingContext.getUniformIndices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetUniformIndices(program WebGLProgram, uniformNames js.Array[js.String]) (ret js.Array[GLuint], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetUniformIndices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformNames.Ref(),
	)

	return
}

// HasGetActiveUniforms returns true if the method "WebGL2RenderingContext.getActiveUniforms" exists.
func (this WebGL2RenderingContext) HasGetActiveUniforms() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetActiveUniforms(
		this.Ref(),
	)
}

// GetActiveUniformsFunc returns the method "WebGL2RenderingContext.getActiveUniforms".
func (this WebGL2RenderingContext) GetActiveUniformsFunc() (fn js.Func[func(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformsFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniforms calls the method "WebGL2RenderingContext.getActiveUniforms".
func (this WebGL2RenderingContext) GetActiveUniforms(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetActiveUniforms(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uniformIndices.Ref(),
		uint32(pname),
	)

	return
}

// TryGetActiveUniforms calls the method "WebGL2RenderingContext.getActiveUniforms"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetActiveUniforms(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetActiveUniforms(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformIndices.Ref(),
		uint32(pname),
	)

	return
}

// HasGetUniformBlockIndex returns true if the method "WebGL2RenderingContext.getUniformBlockIndex" exists.
func (this WebGL2RenderingContext) HasGetUniformBlockIndex() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetUniformBlockIndex(
		this.Ref(),
	)
}

// GetUniformBlockIndexFunc returns the method "WebGL2RenderingContext.getUniformBlockIndex".
func (this WebGL2RenderingContext) GetUniformBlockIndexFunc() (fn js.Func[func(program WebGLProgram, uniformBlockName js.String) GLuint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformBlockIndexFunc(
			this.Ref(),
		),
	)
}

// GetUniformBlockIndex calls the method "WebGL2RenderingContext.getUniformBlockIndex".
func (this WebGL2RenderingContext) GetUniformBlockIndex(program WebGLProgram, uniformBlockName js.String) (ret GLuint) {
	bindings.CallWebGL2RenderingContextGetUniformBlockIndex(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uniformBlockName.Ref(),
	)

	return
}

// TryGetUniformBlockIndex calls the method "WebGL2RenderingContext.getUniformBlockIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetUniformBlockIndex(program WebGLProgram, uniformBlockName js.String) (ret GLuint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetUniformBlockIndex(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformBlockName.Ref(),
	)

	return
}

// HasGetActiveUniformBlockParameter returns true if the method "WebGL2RenderingContext.getActiveUniformBlockParameter" exists.
func (this WebGL2RenderingContext) HasGetActiveUniformBlockParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.Ref(),
	)
}

// GetActiveUniformBlockParameterFunc returns the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
func (this WebGL2RenderingContext) GetActiveUniformBlockParameterFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformBlockParameterFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniformBlockParameter calls the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
func (this WebGL2RenderingContext) GetActiveUniformBlockParameter(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(pname),
	)

	return
}

// TryGetActiveUniformBlockParameter calls the method "WebGL2RenderingContext.getActiveUniformBlockParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetActiveUniformBlockParameter(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(pname),
	)

	return
}

// HasGetActiveUniformBlockName returns true if the method "WebGL2RenderingContext.getActiveUniformBlockName" exists.
func (this WebGL2RenderingContext) HasGetActiveUniformBlockName() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetActiveUniformBlockName(
		this.Ref(),
	)
}

// GetActiveUniformBlockNameFunc returns the method "WebGL2RenderingContext.getActiveUniformBlockName".
func (this WebGL2RenderingContext) GetActiveUniformBlockNameFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformBlockNameFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniformBlockName calls the method "WebGL2RenderingContext.getActiveUniformBlockName".
func (this WebGL2RenderingContext) GetActiveUniformBlockName(program WebGLProgram, uniformBlockIndex GLuint) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetActiveUniformBlockName(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(uniformBlockIndex),
	)

	return
}

// TryGetActiveUniformBlockName calls the method "WebGL2RenderingContext.getActiveUniformBlockName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetActiveUniformBlockName(program WebGLProgram, uniformBlockIndex GLuint) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetActiveUniformBlockName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
	)

	return
}

// HasUniformBlockBinding returns true if the method "WebGL2RenderingContext.uniformBlockBinding" exists.
func (this WebGL2RenderingContext) HasUniformBlockBinding() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniformBlockBinding(
		this.Ref(),
	)
}

// UniformBlockBindingFunc returns the method "WebGL2RenderingContext.uniformBlockBinding".
func (this WebGL2RenderingContext) UniformBlockBindingFunc() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniformBlockBindingFunc(
			this.Ref(),
		),
	)
}

// UniformBlockBinding calls the method "WebGL2RenderingContext.uniformBlockBinding".
func (this WebGL2RenderingContext) UniformBlockBinding(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformBlockBinding(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(uniformBlockBinding),
	)

	return
}

// TryUniformBlockBinding calls the method "WebGL2RenderingContext.uniformBlockBinding"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniformBlockBinding(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniformBlockBinding(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(uniformBlockBinding),
	)

	return
}

// HasCreateVertexArray returns true if the method "WebGL2RenderingContext.createVertexArray" exists.
func (this WebGL2RenderingContext) HasCreateVertexArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateVertexArray(
		this.Ref(),
	)
}

// CreateVertexArrayFunc returns the method "WebGL2RenderingContext.createVertexArray".
func (this WebGL2RenderingContext) CreateVertexArrayFunc() (fn js.Func[func() WebGLVertexArrayObject]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateVertexArrayFunc(
			this.Ref(),
		),
	)
}

// CreateVertexArray calls the method "WebGL2RenderingContext.createVertexArray".
func (this WebGL2RenderingContext) CreateVertexArray() (ret WebGLVertexArrayObject) {
	bindings.CallWebGL2RenderingContextCreateVertexArray(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateVertexArray calls the method "WebGL2RenderingContext.createVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateVertexArray() (ret WebGLVertexArrayObject, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateVertexArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteVertexArray returns true if the method "WebGL2RenderingContext.deleteVertexArray" exists.
func (this WebGL2RenderingContext) HasDeleteVertexArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteVertexArray(
		this.Ref(),
	)
}

// DeleteVertexArrayFunc returns the method "WebGL2RenderingContext.deleteVertexArray".
func (this WebGL2RenderingContext) DeleteVertexArrayFunc() (fn js.Func[func(vertexArray WebGLVertexArrayObject)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteVertexArrayFunc(
			this.Ref(),
		),
	)
}

// DeleteVertexArray calls the method "WebGL2RenderingContext.deleteVertexArray".
func (this WebGL2RenderingContext) DeleteVertexArray(vertexArray WebGLVertexArrayObject) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteVertexArray(
		this.Ref(), js.Pointer(&ret),
		vertexArray.Ref(),
	)

	return
}

// TryDeleteVertexArray calls the method "WebGL2RenderingContext.deleteVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteVertexArray(vertexArray WebGLVertexArrayObject) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteVertexArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		vertexArray.Ref(),
	)

	return
}

// HasIsVertexArray returns true if the method "WebGL2RenderingContext.isVertexArray" exists.
func (this WebGL2RenderingContext) HasIsVertexArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsVertexArray(
		this.Ref(),
	)
}

// IsVertexArrayFunc returns the method "WebGL2RenderingContext.isVertexArray".
func (this WebGL2RenderingContext) IsVertexArrayFunc() (fn js.Func[func(vertexArray WebGLVertexArrayObject) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsVertexArrayFunc(
			this.Ref(),
		),
	)
}

// IsVertexArray calls the method "WebGL2RenderingContext.isVertexArray".
func (this WebGL2RenderingContext) IsVertexArray(vertexArray WebGLVertexArrayObject) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsVertexArray(
		this.Ref(), js.Pointer(&ret),
		vertexArray.Ref(),
	)

	return
}

// TryIsVertexArray calls the method "WebGL2RenderingContext.isVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsVertexArray(vertexArray WebGLVertexArrayObject) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsVertexArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		vertexArray.Ref(),
	)

	return
}

// HasBindVertexArray returns true if the method "WebGL2RenderingContext.bindVertexArray" exists.
func (this WebGL2RenderingContext) HasBindVertexArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindVertexArray(
		this.Ref(),
	)
}

// BindVertexArrayFunc returns the method "WebGL2RenderingContext.bindVertexArray".
func (this WebGL2RenderingContext) BindVertexArrayFunc() (fn js.Func[func(array WebGLVertexArrayObject)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindVertexArrayFunc(
			this.Ref(),
		),
	)
}

// BindVertexArray calls the method "WebGL2RenderingContext.bindVertexArray".
func (this WebGL2RenderingContext) BindVertexArray(array WebGLVertexArrayObject) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindVertexArray(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryBindVertexArray calls the method "WebGL2RenderingContext.bindVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindVertexArray(array WebGLVertexArrayObject) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindVertexArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasGetContextAttributes returns true if the method "WebGL2RenderingContext.getContextAttributes" exists.
func (this WebGL2RenderingContext) HasGetContextAttributes() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetContextAttributes(
		this.Ref(),
	)
}

// GetContextAttributesFunc returns the method "WebGL2RenderingContext.getContextAttributes".
func (this WebGL2RenderingContext) GetContextAttributesFunc() (fn js.Func[func() WebGLContextAttributes]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetContextAttributesFunc(
			this.Ref(),
		),
	)
}

// GetContextAttributes calls the method "WebGL2RenderingContext.getContextAttributes".
func (this WebGL2RenderingContext) GetContextAttributes() (ret WebGLContextAttributes) {
	bindings.CallWebGL2RenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "WebGL2RenderingContext.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetContextAttributes() (ret WebGLContextAttributes, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetContextAttributes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsContextLost returns true if the method "WebGL2RenderingContext.isContextLost" exists.
func (this WebGL2RenderingContext) HasIsContextLost() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsContextLost(
		this.Ref(),
	)
}

// IsContextLostFunc returns the method "WebGL2RenderingContext.isContextLost".
func (this WebGL2RenderingContext) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsContextLostFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "WebGL2RenderingContext.isContextLost".
func (this WebGL2RenderingContext) IsContextLost() (ret bool) {
	bindings.CallWebGL2RenderingContextIsContextLost(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "WebGL2RenderingContext.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsContextLost(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSupportedExtensions returns true if the method "WebGL2RenderingContext.getSupportedExtensions" exists.
func (this WebGL2RenderingContext) HasGetSupportedExtensions() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetSupportedExtensions(
		this.Ref(),
	)
}

// GetSupportedExtensionsFunc returns the method "WebGL2RenderingContext.getSupportedExtensions".
func (this WebGL2RenderingContext) GetSupportedExtensionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetSupportedExtensionsFunc(
			this.Ref(),
		),
	)
}

// GetSupportedExtensions calls the method "WebGL2RenderingContext.getSupportedExtensions".
func (this WebGL2RenderingContext) GetSupportedExtensions() (ret js.Array[js.String]) {
	bindings.CallWebGL2RenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedExtensions calls the method "WebGL2RenderingContext.getSupportedExtensions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetSupportedExtensions() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetSupportedExtensions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetExtension returns true if the method "WebGL2RenderingContext.getExtension" exists.
func (this WebGL2RenderingContext) HasGetExtension() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetExtension(
		this.Ref(),
	)
}

// GetExtensionFunc returns the method "WebGL2RenderingContext.getExtension".
func (this WebGL2RenderingContext) GetExtensionFunc() (fn js.Func[func(name js.String) js.Object]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetExtensionFunc(
			this.Ref(),
		),
	)
}

// GetExtension calls the method "WebGL2RenderingContext.getExtension".
func (this WebGL2RenderingContext) GetExtension(name js.String) (ret js.Object) {
	bindings.CallWebGL2RenderingContextGetExtension(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetExtension calls the method "WebGL2RenderingContext.getExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetExtension(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetExtension(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasActiveTexture returns true if the method "WebGL2RenderingContext.activeTexture" exists.
func (this WebGL2RenderingContext) HasActiveTexture() bool {
	return js.True == bindings.HasWebGL2RenderingContextActiveTexture(
		this.Ref(),
	)
}

// ActiveTextureFunc returns the method "WebGL2RenderingContext.activeTexture".
func (this WebGL2RenderingContext) ActiveTextureFunc() (fn js.Func[func(texture GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextActiveTextureFunc(
			this.Ref(),
		),
	)
}

// ActiveTexture calls the method "WebGL2RenderingContext.activeTexture".
func (this WebGL2RenderingContext) ActiveTexture(texture GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextActiveTexture(
		this.Ref(), js.Pointer(&ret),
		uint32(texture),
	)

	return
}

// TryActiveTexture calls the method "WebGL2RenderingContext.activeTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryActiveTexture(texture GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextActiveTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(texture),
	)

	return
}

// HasAttachShader returns true if the method "WebGL2RenderingContext.attachShader" exists.
func (this WebGL2RenderingContext) HasAttachShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextAttachShader(
		this.Ref(),
	)
}

// AttachShaderFunc returns the method "WebGL2RenderingContext.attachShader".
func (this WebGL2RenderingContext) AttachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextAttachShaderFunc(
			this.Ref(),
		),
	)
}

// AttachShader calls the method "WebGL2RenderingContext.attachShader".
func (this WebGL2RenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextAttachShader(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// TryAttachShader calls the method "WebGL2RenderingContext.attachShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryAttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextAttachShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasBindAttribLocation returns true if the method "WebGL2RenderingContext.bindAttribLocation" exists.
func (this WebGL2RenderingContext) HasBindAttribLocation() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindAttribLocation(
		this.Ref(),
	)
}

// BindAttribLocationFunc returns the method "WebGL2RenderingContext.bindAttribLocation".
func (this WebGL2RenderingContext) BindAttribLocationFunc() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindAttribLocationFunc(
			this.Ref(),
		),
	)
}

// BindAttribLocation calls the method "WebGL2RenderingContext.bindAttribLocation".
func (this WebGL2RenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// TryBindAttribLocation calls the method "WebGL2RenderingContext.bindAttribLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindAttribLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// HasBindBuffer returns true if the method "WebGL2RenderingContext.bindBuffer" exists.
func (this WebGL2RenderingContext) HasBindBuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindBuffer(
		this.Ref(),
	)
}

// BindBufferFunc returns the method "WebGL2RenderingContext.bindBuffer".
func (this WebGL2RenderingContext) BindBufferFunc() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindBufferFunc(
			this.Ref(),
		),
	)
}

// BindBuffer calls the method "WebGL2RenderingContext.bindBuffer".
func (this WebGL2RenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// TryBindBuffer calls the method "WebGL2RenderingContext.bindBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// HasBindFramebuffer returns true if the method "WebGL2RenderingContext.bindFramebuffer" exists.
func (this WebGL2RenderingContext) HasBindFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindFramebuffer(
		this.Ref(),
	)
}

// BindFramebufferFunc returns the method "WebGL2RenderingContext.bindFramebuffer".
func (this WebGL2RenderingContext) BindFramebufferFunc() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindFramebufferFunc(
			this.Ref(),
		),
	)
}

// BindFramebuffer calls the method "WebGL2RenderingContext.bindFramebuffer".
func (this WebGL2RenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// TryBindFramebuffer calls the method "WebGL2RenderingContext.bindFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// HasBindRenderbuffer returns true if the method "WebGL2RenderingContext.bindRenderbuffer" exists.
func (this WebGL2RenderingContext) HasBindRenderbuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindRenderbuffer(
		this.Ref(),
	)
}

// BindRenderbufferFunc returns the method "WebGL2RenderingContext.bindRenderbuffer".
func (this WebGL2RenderingContext) BindRenderbufferFunc() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindRenderbufferFunc(
			this.Ref(),
		),
	)
}

// BindRenderbuffer calls the method "WebGL2RenderingContext.bindRenderbuffer".
func (this WebGL2RenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// TryBindRenderbuffer calls the method "WebGL2RenderingContext.bindRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// HasBindTexture returns true if the method "WebGL2RenderingContext.bindTexture" exists.
func (this WebGL2RenderingContext) HasBindTexture() bool {
	return js.True == bindings.HasWebGL2RenderingContextBindTexture(
		this.Ref(),
	)
}

// BindTextureFunc returns the method "WebGL2RenderingContext.bindTexture".
func (this WebGL2RenderingContext) BindTextureFunc() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBindTextureFunc(
			this.Ref(),
		),
	)
}

// BindTexture calls the method "WebGL2RenderingContext.bindTexture".
func (this WebGL2RenderingContext) BindTexture(target GLenum, texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindTexture(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		texture.Ref(),
	)

	return
}

// TryBindTexture calls the method "WebGL2RenderingContext.bindTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindTexture(target GLenum, texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		texture.Ref(),
	)

	return
}

// HasBlendColor returns true if the method "WebGL2RenderingContext.blendColor" exists.
func (this WebGL2RenderingContext) HasBlendColor() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlendColor(
		this.Ref(),
	)
}

// BlendColorFunc returns the method "WebGL2RenderingContext.blendColor".
func (this WebGL2RenderingContext) BlendColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendColorFunc(
			this.Ref(),
		),
	)
}

// BlendColor calls the method "WebGL2RenderingContext.blendColor".
func (this WebGL2RenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendColor(
		this.Ref(), js.Pointer(&ret),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// TryBlendColor calls the method "WebGL2RenderingContext.blendColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendColor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasBlendEquation returns true if the method "WebGL2RenderingContext.blendEquation" exists.
func (this WebGL2RenderingContext) HasBlendEquation() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlendEquation(
		this.Ref(),
	)
}

// BlendEquationFunc returns the method "WebGL2RenderingContext.blendEquation".
func (this WebGL2RenderingContext) BlendEquationFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendEquationFunc(
			this.Ref(),
		),
	)
}

// BlendEquation calls the method "WebGL2RenderingContext.blendEquation".
func (this WebGL2RenderingContext) BlendEquation(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendEquation(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryBlendEquation calls the method "WebGL2RenderingContext.blendEquation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendEquation(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendEquation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasBlendEquationSeparate returns true if the method "WebGL2RenderingContext.blendEquationSeparate" exists.
func (this WebGL2RenderingContext) HasBlendEquationSeparate() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlendEquationSeparate(
		this.Ref(),
	)
}

// BlendEquationSeparateFunc returns the method "WebGL2RenderingContext.blendEquationSeparate".
func (this WebGL2RenderingContext) BlendEquationSeparateFunc() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendEquationSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparate calls the method "WebGL2RenderingContext.blendEquationSeparate".
func (this WebGL2RenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// TryBlendEquationSeparate calls the method "WebGL2RenderingContext.blendEquationSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendEquationSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// HasBlendFunc returns true if the method "WebGL2RenderingContext.blendFunc" exists.
func (this WebGL2RenderingContext) HasBlendFunc() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlendFunc(
		this.Ref(),
	)
}

// BlendFuncFunc returns the method "WebGL2RenderingContext.blendFunc".
func (this WebGL2RenderingContext) BlendFuncFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendFuncFunc(
			this.Ref(),
		),
	)
}

// BlendFunc calls the method "WebGL2RenderingContext.blendFunc".
func (this WebGL2RenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// TryBlendFunc calls the method "WebGL2RenderingContext.blendFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// HasBlendFuncSeparate returns true if the method "WebGL2RenderingContext.blendFuncSeparate" exists.
func (this WebGL2RenderingContext) HasBlendFuncSeparate() bool {
	return js.True == bindings.HasWebGL2RenderingContextBlendFuncSeparate(
		this.Ref(),
	)
}

// BlendFuncSeparateFunc returns the method "WebGL2RenderingContext.blendFuncSeparate".
func (this WebGL2RenderingContext) BlendFuncSeparateFunc() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextBlendFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparate calls the method "WebGL2RenderingContext.blendFuncSeparate".
func (this WebGL2RenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// TryBlendFuncSeparate calls the method "WebGL2RenderingContext.blendFuncSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendFuncSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// HasCheckFramebufferStatus returns true if the method "WebGL2RenderingContext.checkFramebufferStatus" exists.
func (this WebGL2RenderingContext) HasCheckFramebufferStatus() bool {
	return js.True == bindings.HasWebGL2RenderingContextCheckFramebufferStatus(
		this.Ref(),
	)
}

// CheckFramebufferStatusFunc returns the method "WebGL2RenderingContext.checkFramebufferStatus".
func (this WebGL2RenderingContext) CheckFramebufferStatusFunc() (fn js.Func[func(target GLenum) GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCheckFramebufferStatusFunc(
			this.Ref(),
		),
	)
}

// CheckFramebufferStatus calls the method "WebGL2RenderingContext.checkFramebufferStatus".
func (this WebGL2RenderingContext) CheckFramebufferStatus(target GLenum) (ret GLenum) {
	bindings.CallWebGL2RenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryCheckFramebufferStatus calls the method "WebGL2RenderingContext.checkFramebufferStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCheckFramebufferStatus(target GLenum) (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCheckFramebufferStatus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasClear returns true if the method "WebGL2RenderingContext.clear" exists.
func (this WebGL2RenderingContext) HasClear() bool {
	return js.True == bindings.HasWebGL2RenderingContextClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "WebGL2RenderingContext.clear".
func (this WebGL2RenderingContext) ClearFunc() (fn js.Func[func(mask GLbitfield)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "WebGL2RenderingContext.clear".
func (this WebGL2RenderingContext) Clear(mask GLbitfield) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClear(
		this.Ref(), js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryClear calls the method "WebGL2RenderingContext.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClear(mask GLbitfield) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasClearColor returns true if the method "WebGL2RenderingContext.clearColor" exists.
func (this WebGL2RenderingContext) HasClearColor() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearColor(
		this.Ref(),
	)
}

// ClearColorFunc returns the method "WebGL2RenderingContext.clearColor".
func (this WebGL2RenderingContext) ClearColorFunc() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearColorFunc(
			this.Ref(),
		),
	)
}

// ClearColor calls the method "WebGL2RenderingContext.clearColor".
func (this WebGL2RenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearColor(
		this.Ref(), js.Pointer(&ret),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// TryClearColor calls the method "WebGL2RenderingContext.clearColor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearColor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasClearDepth returns true if the method "WebGL2RenderingContext.clearDepth" exists.
func (this WebGL2RenderingContext) HasClearDepth() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearDepth(
		this.Ref(),
	)
}

// ClearDepthFunc returns the method "WebGL2RenderingContext.clearDepth".
func (this WebGL2RenderingContext) ClearDepthFunc() (fn js.Func[func(depth GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearDepthFunc(
			this.Ref(),
		),
	)
}

// ClearDepth calls the method "WebGL2RenderingContext.clearDepth".
func (this WebGL2RenderingContext) ClearDepth(depth GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearDepth(
		this.Ref(), js.Pointer(&ret),
		float32(depth),
	)

	return
}

// TryClearDepth calls the method "WebGL2RenderingContext.clearDepth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearDepth(depth GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearDepth(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(depth),
	)

	return
}

// HasClearStencil returns true if the method "WebGL2RenderingContext.clearStencil" exists.
func (this WebGL2RenderingContext) HasClearStencil() bool {
	return js.True == bindings.HasWebGL2RenderingContextClearStencil(
		this.Ref(),
	)
}

// ClearStencilFunc returns the method "WebGL2RenderingContext.clearStencil".
func (this WebGL2RenderingContext) ClearStencilFunc() (fn js.Func[func(s GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextClearStencilFunc(
			this.Ref(),
		),
	)
}

// ClearStencil calls the method "WebGL2RenderingContext.clearStencil".
func (this WebGL2RenderingContext) ClearStencil(s GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearStencil(
		this.Ref(), js.Pointer(&ret),
		int32(s),
	)

	return
}

// TryClearStencil calls the method "WebGL2RenderingContext.clearStencil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearStencil(s GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearStencil(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(s),
	)

	return
}

// HasColorMask returns true if the method "WebGL2RenderingContext.colorMask" exists.
func (this WebGL2RenderingContext) HasColorMask() bool {
	return js.True == bindings.HasWebGL2RenderingContextColorMask(
		this.Ref(),
	)
}

// ColorMaskFunc returns the method "WebGL2RenderingContext.colorMask".
func (this WebGL2RenderingContext) ColorMaskFunc() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextColorMaskFunc(
			this.Ref(),
		),
	)
}

// ColorMask calls the method "WebGL2RenderingContext.colorMask".
func (this WebGL2RenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextColorMask(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// TryColorMask calls the method "WebGL2RenderingContext.colorMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextColorMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// HasCompileShader returns true if the method "WebGL2RenderingContext.compileShader" exists.
func (this WebGL2RenderingContext) HasCompileShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextCompileShader(
		this.Ref(),
	)
}

// CompileShaderFunc returns the method "WebGL2RenderingContext.compileShader".
func (this WebGL2RenderingContext) CompileShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCompileShaderFunc(
			this.Ref(),
		),
	)
}

// CompileShader calls the method "WebGL2RenderingContext.compileShader".
func (this WebGL2RenderingContext) CompileShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompileShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryCompileShader calls the method "WebGL2RenderingContext.compileShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompileShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompileShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasCopyTexImage2D returns true if the method "WebGL2RenderingContext.copyTexImage2D" exists.
func (this WebGL2RenderingContext) HasCopyTexImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCopyTexImage2D(
		this.Ref(),
	)
}

// CopyTexImage2DFunc returns the method "WebGL2RenderingContext.copyTexImage2D".
func (this WebGL2RenderingContext) CopyTexImage2DFunc() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexImage2D calls the method "WebGL2RenderingContext.copyTexImage2D".
func (this WebGL2RenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	return
}

// TryCopyTexImage2D calls the method "WebGL2RenderingContext.copyTexImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCopyTexImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		uint32(internalformat),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
		int32(border),
	)

	return
}

// HasCopyTexSubImage2D returns true if the method "WebGL2RenderingContext.copyTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasCopyTexSubImage2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextCopyTexSubImage2D(
		this.Ref(),
	)
}

// CopyTexSubImage2DFunc returns the method "WebGL2RenderingContext.copyTexSubImage2D".
func (this WebGL2RenderingContext) CopyTexSubImage2DFunc() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCopyTexSubImage2DFunc(
			this.Ref(),
		),
	)
}

// CopyTexSubImage2D calls the method "WebGL2RenderingContext.copyTexSubImage2D".
func (this WebGL2RenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryCopyTexSubImage2D calls the method "WebGL2RenderingContext.copyTexSubImage2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCopyTexSubImage2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(xoffset),
		int32(yoffset),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasCreateBuffer returns true if the method "WebGL2RenderingContext.createBuffer" exists.
func (this WebGL2RenderingContext) HasCreateBuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateBuffer(
		this.Ref(),
	)
}

// CreateBufferFunc returns the method "WebGL2RenderingContext.createBuffer".
func (this WebGL2RenderingContext) CreateBufferFunc() (fn js.Func[func() WebGLBuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "WebGL2RenderingContext.createBuffer".
func (this WebGL2RenderingContext) CreateBuffer() (ret WebGLBuffer) {
	bindings.CallWebGL2RenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateBuffer calls the method "WebGL2RenderingContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateBuffer() (ret WebGLBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateFramebuffer returns true if the method "WebGL2RenderingContext.createFramebuffer" exists.
func (this WebGL2RenderingContext) HasCreateFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateFramebuffer(
		this.Ref(),
	)
}

// CreateFramebufferFunc returns the method "WebGL2RenderingContext.createFramebuffer".
func (this WebGL2RenderingContext) CreateFramebufferFunc() (fn js.Func[func() WebGLFramebuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateFramebufferFunc(
			this.Ref(),
		),
	)
}

// CreateFramebuffer calls the method "WebGL2RenderingContext.createFramebuffer".
func (this WebGL2RenderingContext) CreateFramebuffer() (ret WebGLFramebuffer) {
	bindings.CallWebGL2RenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateFramebuffer calls the method "WebGL2RenderingContext.createFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateFramebuffer() (ret WebGLFramebuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateProgram returns true if the method "WebGL2RenderingContext.createProgram" exists.
func (this WebGL2RenderingContext) HasCreateProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateProgram(
		this.Ref(),
	)
}

// CreateProgramFunc returns the method "WebGL2RenderingContext.createProgram".
func (this WebGL2RenderingContext) CreateProgramFunc() (fn js.Func[func() WebGLProgram]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateProgramFunc(
			this.Ref(),
		),
	)
}

// CreateProgram calls the method "WebGL2RenderingContext.createProgram".
func (this WebGL2RenderingContext) CreateProgram() (ret WebGLProgram) {
	bindings.CallWebGL2RenderingContextCreateProgram(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateProgram calls the method "WebGL2RenderingContext.createProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateProgram() (ret WebGLProgram, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateRenderbuffer returns true if the method "WebGL2RenderingContext.createRenderbuffer" exists.
func (this WebGL2RenderingContext) HasCreateRenderbuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateRenderbuffer(
		this.Ref(),
	)
}

// CreateRenderbufferFunc returns the method "WebGL2RenderingContext.createRenderbuffer".
func (this WebGL2RenderingContext) CreateRenderbufferFunc() (fn js.Func[func() WebGLRenderbuffer]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateRenderbufferFunc(
			this.Ref(),
		),
	)
}

// CreateRenderbuffer calls the method "WebGL2RenderingContext.createRenderbuffer".
func (this WebGL2RenderingContext) CreateRenderbuffer() (ret WebGLRenderbuffer) {
	bindings.CallWebGL2RenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateRenderbuffer calls the method "WebGL2RenderingContext.createRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateRenderbuffer() (ret WebGLRenderbuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateShader returns true if the method "WebGL2RenderingContext.createShader" exists.
func (this WebGL2RenderingContext) HasCreateShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateShader(
		this.Ref(),
	)
}

// CreateShaderFunc returns the method "WebGL2RenderingContext.createShader".
func (this WebGL2RenderingContext) CreateShaderFunc() (fn js.Func[func(typ GLenum) WebGLShader]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateShaderFunc(
			this.Ref(),
		),
	)
}

// CreateShader calls the method "WebGL2RenderingContext.createShader".
func (this WebGL2RenderingContext) CreateShader(typ GLenum) (ret WebGLShader) {
	bindings.CallWebGL2RenderingContextCreateShader(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCreateShader calls the method "WebGL2RenderingContext.createShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateShader(typ GLenum) (ret WebGLShader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasCreateTexture returns true if the method "WebGL2RenderingContext.createTexture" exists.
func (this WebGL2RenderingContext) HasCreateTexture() bool {
	return js.True == bindings.HasWebGL2RenderingContextCreateTexture(
		this.Ref(),
	)
}

// CreateTextureFunc returns the method "WebGL2RenderingContext.createTexture".
func (this WebGL2RenderingContext) CreateTextureFunc() (fn js.Func[func() WebGLTexture]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "WebGL2RenderingContext.createTexture".
func (this WebGL2RenderingContext) CreateTexture() (ret WebGLTexture) {
	bindings.CallWebGL2RenderingContextCreateTexture(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTexture calls the method "WebGL2RenderingContext.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateTexture() (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCullFace returns true if the method "WebGL2RenderingContext.cullFace" exists.
func (this WebGL2RenderingContext) HasCullFace() bool {
	return js.True == bindings.HasWebGL2RenderingContextCullFace(
		this.Ref(),
	)
}

// CullFaceFunc returns the method "WebGL2RenderingContext.cullFace".
func (this WebGL2RenderingContext) CullFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextCullFaceFunc(
			this.Ref(),
		),
	)
}

// CullFace calls the method "WebGL2RenderingContext.cullFace".
func (this WebGL2RenderingContext) CullFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCullFace(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryCullFace calls the method "WebGL2RenderingContext.cullFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCullFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCullFace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasDeleteBuffer returns true if the method "WebGL2RenderingContext.deleteBuffer" exists.
func (this WebGL2RenderingContext) HasDeleteBuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteBuffer(
		this.Ref(),
	)
}

// DeleteBufferFunc returns the method "WebGL2RenderingContext.deleteBuffer".
func (this WebGL2RenderingContext) DeleteBufferFunc() (fn js.Func[func(buffer WebGLBuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteBufferFunc(
			this.Ref(),
		),
	)
}

// DeleteBuffer calls the method "WebGL2RenderingContext.deleteBuffer".
func (this WebGL2RenderingContext) DeleteBuffer(buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryDeleteBuffer calls the method "WebGL2RenderingContext.deleteBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteBuffer(buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasDeleteFramebuffer returns true if the method "WebGL2RenderingContext.deleteFramebuffer" exists.
func (this WebGL2RenderingContext) HasDeleteFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteFramebuffer(
		this.Ref(),
	)
}

// DeleteFramebufferFunc returns the method "WebGL2RenderingContext.deleteFramebuffer".
func (this WebGL2RenderingContext) DeleteFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteFramebufferFunc(
			this.Ref(),
		),
	)
}

// DeleteFramebuffer calls the method "WebGL2RenderingContext.deleteFramebuffer".
func (this WebGL2RenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryDeleteFramebuffer calls the method "WebGL2RenderingContext.deleteFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasDeleteProgram returns true if the method "WebGL2RenderingContext.deleteProgram" exists.
func (this WebGL2RenderingContext) HasDeleteProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteProgram(
		this.Ref(),
	)
}

// DeleteProgramFunc returns the method "WebGL2RenderingContext.deleteProgram".
func (this WebGL2RenderingContext) DeleteProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteProgramFunc(
			this.Ref(),
		),
	)
}

// DeleteProgram calls the method "WebGL2RenderingContext.deleteProgram".
func (this WebGL2RenderingContext) DeleteProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryDeleteProgram calls the method "WebGL2RenderingContext.deleteProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasDeleteRenderbuffer returns true if the method "WebGL2RenderingContext.deleteRenderbuffer" exists.
func (this WebGL2RenderingContext) HasDeleteRenderbuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteRenderbuffer(
		this.Ref(),
	)
}

// DeleteRenderbufferFunc returns the method "WebGL2RenderingContext.deleteRenderbuffer".
func (this WebGL2RenderingContext) DeleteRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteRenderbufferFunc(
			this.Ref(),
		),
	)
}

// DeleteRenderbuffer calls the method "WebGL2RenderingContext.deleteRenderbuffer".
func (this WebGL2RenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryDeleteRenderbuffer calls the method "WebGL2RenderingContext.deleteRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasDeleteShader returns true if the method "WebGL2RenderingContext.deleteShader" exists.
func (this WebGL2RenderingContext) HasDeleteShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteShader(
		this.Ref(),
	)
}

// DeleteShaderFunc returns the method "WebGL2RenderingContext.deleteShader".
func (this WebGL2RenderingContext) DeleteShaderFunc() (fn js.Func[func(shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteShaderFunc(
			this.Ref(),
		),
	)
}

// DeleteShader calls the method "WebGL2RenderingContext.deleteShader".
func (this WebGL2RenderingContext) DeleteShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryDeleteShader calls the method "WebGL2RenderingContext.deleteShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasDeleteTexture returns true if the method "WebGL2RenderingContext.deleteTexture" exists.
func (this WebGL2RenderingContext) HasDeleteTexture() bool {
	return js.True == bindings.HasWebGL2RenderingContextDeleteTexture(
		this.Ref(),
	)
}

// DeleteTextureFunc returns the method "WebGL2RenderingContext.deleteTexture".
func (this WebGL2RenderingContext) DeleteTextureFunc() (fn js.Func[func(texture WebGLTexture)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDeleteTextureFunc(
			this.Ref(),
		),
	)
}

// DeleteTexture calls the method "WebGL2RenderingContext.deleteTexture".
func (this WebGL2RenderingContext) DeleteTexture(texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryDeleteTexture calls the method "WebGL2RenderingContext.deleteTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteTexture(texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasDepthFunc returns true if the method "WebGL2RenderingContext.depthFunc" exists.
func (this WebGL2RenderingContext) HasDepthFunc() bool {
	return js.True == bindings.HasWebGL2RenderingContextDepthFunc(
		this.Ref(),
	)
}

// DepthFuncFunc returns the method "WebGL2RenderingContext.depthFunc".
func (this WebGL2RenderingContext) DepthFuncFunc() (fn js.Func[func(fn GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthFuncFunc(
			this.Ref(),
		),
	)
}

// DepthFunc calls the method "WebGL2RenderingContext.depthFunc".
func (this WebGL2RenderingContext) DepthFunc(fn GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(fn),
	)

	return
}

// TryDepthFunc calls the method "WebGL2RenderingContext.depthFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDepthFunc(fn GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDepthFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
	)

	return
}

// HasDepthMask returns true if the method "WebGL2RenderingContext.depthMask" exists.
func (this WebGL2RenderingContext) HasDepthMask() bool {
	return js.True == bindings.HasWebGL2RenderingContextDepthMask(
		this.Ref(),
	)
}

// DepthMaskFunc returns the method "WebGL2RenderingContext.depthMask".
func (this WebGL2RenderingContext) DepthMaskFunc() (fn js.Func[func(flag GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthMaskFunc(
			this.Ref(),
		),
	)
}

// DepthMask calls the method "WebGL2RenderingContext.depthMask".
func (this WebGL2RenderingContext) DepthMask(flag GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthMask(
		this.Ref(), js.Pointer(&ret),
		js.Bool(bool(flag)),
	)

	return
}

// TryDepthMask calls the method "WebGL2RenderingContext.depthMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDepthMask(flag GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDepthMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(flag)),
	)

	return
}

// HasDepthRange returns true if the method "WebGL2RenderingContext.depthRange" exists.
func (this WebGL2RenderingContext) HasDepthRange() bool {
	return js.True == bindings.HasWebGL2RenderingContextDepthRange(
		this.Ref(),
	)
}

// DepthRangeFunc returns the method "WebGL2RenderingContext.depthRange".
func (this WebGL2RenderingContext) DepthRangeFunc() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDepthRangeFunc(
			this.Ref(),
		),
	)
}

// DepthRange calls the method "WebGL2RenderingContext.depthRange".
func (this WebGL2RenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthRange(
		this.Ref(), js.Pointer(&ret),
		float32(zNear),
		float32(zFar),
	)

	return
}

// TryDepthRange calls the method "WebGL2RenderingContext.depthRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDepthRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(zNear),
		float32(zFar),
	)

	return
}

// HasDetachShader returns true if the method "WebGL2RenderingContext.detachShader" exists.
func (this WebGL2RenderingContext) HasDetachShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextDetachShader(
		this.Ref(),
	)
}

// DetachShaderFunc returns the method "WebGL2RenderingContext.detachShader".
func (this WebGL2RenderingContext) DetachShaderFunc() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDetachShaderFunc(
			this.Ref(),
		),
	)
}

// DetachShader calls the method "WebGL2RenderingContext.detachShader".
func (this WebGL2RenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDetachShader(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// TryDetachShader calls the method "WebGL2RenderingContext.detachShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDetachShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasDisable returns true if the method "WebGL2RenderingContext.disable" exists.
func (this WebGL2RenderingContext) HasDisable() bool {
	return js.True == bindings.HasWebGL2RenderingContextDisable(
		this.Ref(),
	)
}

// DisableFunc returns the method "WebGL2RenderingContext.disable".
func (this WebGL2RenderingContext) DisableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDisableFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "WebGL2RenderingContext.disable".
func (this WebGL2RenderingContext) Disable(cap GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDisable(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryDisable calls the method "WebGL2RenderingContext.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDisable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDisable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasDisableVertexAttribArray returns true if the method "WebGL2RenderingContext.disableVertexAttribArray" exists.
func (this WebGL2RenderingContext) HasDisableVertexAttribArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextDisableVertexAttribArray(
		this.Ref(),
	)
}

// DisableVertexAttribArrayFunc returns the method "WebGL2RenderingContext.disableVertexAttribArray".
func (this WebGL2RenderingContext) DisableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDisableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// DisableVertexAttribArray calls the method "WebGL2RenderingContext.disableVertexAttribArray".
func (this WebGL2RenderingContext) DisableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDisableVertexAttribArray calls the method "WebGL2RenderingContext.disableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDisableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDisableVertexAttribArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasDrawArrays returns true if the method "WebGL2RenderingContext.drawArrays" exists.
func (this WebGL2RenderingContext) HasDrawArrays() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawArrays(
		this.Ref(),
	)
}

// DrawArraysFunc returns the method "WebGL2RenderingContext.drawArrays".
func (this WebGL2RenderingContext) DrawArraysFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawArraysFunc(
			this.Ref(),
		),
	)
}

// DrawArrays calls the method "WebGL2RenderingContext.drawArrays".
func (this WebGL2RenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawArrays(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// TryDrawArrays calls the method "WebGL2RenderingContext.drawArrays"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawArrays(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// HasDrawElements returns true if the method "WebGL2RenderingContext.drawElements" exists.
func (this WebGL2RenderingContext) HasDrawElements() bool {
	return js.True == bindings.HasWebGL2RenderingContextDrawElements(
		this.Ref(),
	)
}

// DrawElementsFunc returns the method "WebGL2RenderingContext.drawElements".
func (this WebGL2RenderingContext) DrawElementsFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextDrawElementsFunc(
			this.Ref(),
		),
	)
}

// DrawElements calls the method "WebGL2RenderingContext.drawElements".
func (this WebGL2RenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawElements(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// TryDrawElements calls the method "WebGL2RenderingContext.drawElements"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawElements(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasEnable returns true if the method "WebGL2RenderingContext.enable" exists.
func (this WebGL2RenderingContext) HasEnable() bool {
	return js.True == bindings.HasWebGL2RenderingContextEnable(
		this.Ref(),
	)
}

// EnableFunc returns the method "WebGL2RenderingContext.enable".
func (this WebGL2RenderingContext) EnableFunc() (fn js.Func[func(cap GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEnableFunc(
			this.Ref(),
		),
	)
}

// Enable calls the method "WebGL2RenderingContext.enable".
func (this WebGL2RenderingContext) Enable(cap GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEnable(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryEnable calls the method "WebGL2RenderingContext.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEnable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEnable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasEnableVertexAttribArray returns true if the method "WebGL2RenderingContext.enableVertexAttribArray" exists.
func (this WebGL2RenderingContext) HasEnableVertexAttribArray() bool {
	return js.True == bindings.HasWebGL2RenderingContextEnableVertexAttribArray(
		this.Ref(),
	)
}

// EnableVertexAttribArrayFunc returns the method "WebGL2RenderingContext.enableVertexAttribArray".
func (this WebGL2RenderingContext) EnableVertexAttribArrayFunc() (fn js.Func[func(index GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextEnableVertexAttribArrayFunc(
			this.Ref(),
		),
	)
}

// EnableVertexAttribArray calls the method "WebGL2RenderingContext.enableVertexAttribArray".
func (this WebGL2RenderingContext) EnableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryEnableVertexAttribArray calls the method "WebGL2RenderingContext.enableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEnableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEnableVertexAttribArray(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFinish returns true if the method "WebGL2RenderingContext.finish" exists.
func (this WebGL2RenderingContext) HasFinish() bool {
	return js.True == bindings.HasWebGL2RenderingContextFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "WebGL2RenderingContext.finish".
func (this WebGL2RenderingContext) FinishFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "WebGL2RenderingContext.finish".
func (this WebGL2RenderingContext) Finish() (ret js.Void) {
	bindings.CallWebGL2RenderingContextFinish(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "WebGL2RenderingContext.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFlush returns true if the method "WebGL2RenderingContext.flush" exists.
func (this WebGL2RenderingContext) HasFlush() bool {
	return js.True == bindings.HasWebGL2RenderingContextFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "WebGL2RenderingContext.flush".
func (this WebGL2RenderingContext) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "WebGL2RenderingContext.flush".
func (this WebGL2RenderingContext) Flush() (ret js.Void) {
	bindings.CallWebGL2RenderingContextFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "WebGL2RenderingContext.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFramebufferRenderbuffer returns true if the method "WebGL2RenderingContext.framebufferRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFramebufferRenderbuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextFramebufferRenderbuffer(
		this.Ref(),
	)
}

// FramebufferRenderbufferFunc returns the method "WebGL2RenderingContext.framebufferRenderbuffer".
func (this WebGL2RenderingContext) FramebufferRenderbufferFunc() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferRenderbufferFunc(
			this.Ref(),
		),
	)
}

// FramebufferRenderbuffer calls the method "WebGL2RenderingContext.framebufferRenderbuffer".
func (this WebGL2RenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// TryFramebufferRenderbuffer calls the method "WebGL2RenderingContext.framebufferRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFramebufferRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// HasFramebufferTexture2D returns true if the method "WebGL2RenderingContext.framebufferTexture2D" exists.
func (this WebGL2RenderingContext) HasFramebufferTexture2D() bool {
	return js.True == bindings.HasWebGL2RenderingContextFramebufferTexture2D(
		this.Ref(),
	)
}

// FramebufferTexture2DFunc returns the method "WebGL2RenderingContext.framebufferTexture2D".
func (this WebGL2RenderingContext) FramebufferTexture2DFunc() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFramebufferTexture2DFunc(
			this.Ref(),
		),
	)
}

// FramebufferTexture2D calls the method "WebGL2RenderingContext.framebufferTexture2D".
func (this WebGL2RenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// TryFramebufferTexture2D calls the method "WebGL2RenderingContext.framebufferTexture2D"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFramebufferTexture2D(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// HasFrontFace returns true if the method "WebGL2RenderingContext.frontFace" exists.
func (this WebGL2RenderingContext) HasFrontFace() bool {
	return js.True == bindings.HasWebGL2RenderingContextFrontFace(
		this.Ref(),
	)
}

// FrontFaceFunc returns the method "WebGL2RenderingContext.frontFace".
func (this WebGL2RenderingContext) FrontFaceFunc() (fn js.Func[func(mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextFrontFaceFunc(
			this.Ref(),
		),
	)
}

// FrontFace calls the method "WebGL2RenderingContext.frontFace".
func (this WebGL2RenderingContext) FrontFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFrontFace(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryFrontFace calls the method "WebGL2RenderingContext.frontFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFrontFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFrontFace(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasGenerateMipmap returns true if the method "WebGL2RenderingContext.generateMipmap" exists.
func (this WebGL2RenderingContext) HasGenerateMipmap() bool {
	return js.True == bindings.HasWebGL2RenderingContextGenerateMipmap(
		this.Ref(),
	)
}

// GenerateMipmapFunc returns the method "WebGL2RenderingContext.generateMipmap".
func (this WebGL2RenderingContext) GenerateMipmapFunc() (fn js.Func[func(target GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGenerateMipmapFunc(
			this.Ref(),
		),
	)
}

// GenerateMipmap calls the method "WebGL2RenderingContext.generateMipmap".
func (this WebGL2RenderingContext) GenerateMipmap(target GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryGenerateMipmap calls the method "WebGL2RenderingContext.generateMipmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGenerateMipmap(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGenerateMipmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasGetActiveAttrib returns true if the method "WebGL2RenderingContext.getActiveAttrib" exists.
func (this WebGL2RenderingContext) HasGetActiveAttrib() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetActiveAttrib(
		this.Ref(),
	)
}

// GetActiveAttribFunc returns the method "WebGL2RenderingContext.getActiveAttrib".
func (this WebGL2RenderingContext) GetActiveAttribFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveAttribFunc(
			this.Ref(),
		),
	)
}

// GetActiveAttrib calls the method "WebGL2RenderingContext.getActiveAttrib".
func (this WebGL2RenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
	)

	return
}

// TryGetActiveAttrib calls the method "WebGL2RenderingContext.getActiveAttrib"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetActiveAttrib(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasGetActiveUniform returns true if the method "WebGL2RenderingContext.getActiveUniform" exists.
func (this WebGL2RenderingContext) HasGetActiveUniform() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetActiveUniform(
		this.Ref(),
	)
}

// GetActiveUniformFunc returns the method "WebGL2RenderingContext.getActiveUniform".
func (this WebGL2RenderingContext) GetActiveUniformFunc() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetActiveUniformFunc(
			this.Ref(),
		),
	)
}

// GetActiveUniform calls the method "WebGL2RenderingContext.getActiveUniform".
func (this WebGL2RenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(index),
	)

	return
}

// TryGetActiveUniform calls the method "WebGL2RenderingContext.getActiveUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetActiveUniform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasGetAttachedShaders returns true if the method "WebGL2RenderingContext.getAttachedShaders" exists.
func (this WebGL2RenderingContext) HasGetAttachedShaders() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetAttachedShaders(
		this.Ref(),
	)
}

// GetAttachedShadersFunc returns the method "WebGL2RenderingContext.getAttachedShaders".
func (this WebGL2RenderingContext) GetAttachedShadersFunc() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetAttachedShadersFunc(
			this.Ref(),
		),
	)
}

// GetAttachedShaders calls the method "WebGL2RenderingContext.getAttachedShaders".
func (this WebGL2RenderingContext) GetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader]) {
	bindings.CallWebGL2RenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetAttachedShaders calls the method "WebGL2RenderingContext.getAttachedShaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetAttachedShaders(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasGetAttribLocation returns true if the method "WebGL2RenderingContext.getAttribLocation" exists.
func (this WebGL2RenderingContext) HasGetAttribLocation() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetAttribLocation(
		this.Ref(),
	)
}

// GetAttribLocationFunc returns the method "WebGL2RenderingContext.getAttribLocation".
func (this WebGL2RenderingContext) GetAttribLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetAttribLocationFunc(
			this.Ref(),
		),
	)
}

// GetAttribLocation calls the method "WebGL2RenderingContext.getAttribLocation".
func (this WebGL2RenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGL2RenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		name.Ref(),
	)

	return
}

// TryGetAttribLocation calls the method "WebGL2RenderingContext.getAttribLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetAttribLocation(program WebGLProgram, name js.String) (ret GLint, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetAttribLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasGetBufferParameter returns true if the method "WebGL2RenderingContext.getBufferParameter" exists.
func (this WebGL2RenderingContext) HasGetBufferParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetBufferParameter(
		this.Ref(),
	)
}

// GetBufferParameterFunc returns the method "WebGL2RenderingContext.getBufferParameter".
func (this WebGL2RenderingContext) GetBufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetBufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetBufferParameter calls the method "WebGL2RenderingContext.getBufferParameter".
func (this WebGL2RenderingContext) GetBufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetBufferParameter calls the method "WebGL2RenderingContext.getBufferParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetBufferParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetBufferParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetParameter returns true if the method "WebGL2RenderingContext.getParameter" exists.
func (this WebGL2RenderingContext) HasGetParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetParameter(
		this.Ref(),
	)
}

// GetParameterFunc returns the method "WebGL2RenderingContext.getParameter".
func (this WebGL2RenderingContext) GetParameterFunc() (fn js.Func[func(pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetParameterFunc(
			this.Ref(),
		),
	)
}

// GetParameter calls the method "WebGL2RenderingContext.getParameter".
func (this WebGL2RenderingContext) GetParameter(pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(pname),
	)

	return
}

// TryGetParameter calls the method "WebGL2RenderingContext.getParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetParameter(pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
	)

	return
}

// HasGetError returns true if the method "WebGL2RenderingContext.getError" exists.
func (this WebGL2RenderingContext) HasGetError() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetError(
		this.Ref(),
	)
}

// GetErrorFunc returns the method "WebGL2RenderingContext.getError".
func (this WebGL2RenderingContext) GetErrorFunc() (fn js.Func[func() GLenum]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetErrorFunc(
			this.Ref(),
		),
	)
}

// GetError calls the method "WebGL2RenderingContext.getError".
func (this WebGL2RenderingContext) GetError() (ret GLenum) {
	bindings.CallWebGL2RenderingContextGetError(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetError calls the method "WebGL2RenderingContext.getError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetError() (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetError(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetFramebufferAttachmentParameter returns true if the method "WebGL2RenderingContext.getFramebufferAttachmentParameter" exists.
func (this WebGL2RenderingContext) HasGetFramebufferAttachmentParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.Ref(),
	)
}

// GetFramebufferAttachmentParameterFunc returns the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
func (this WebGL2RenderingContext) GetFramebufferAttachmentParameterFunc() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetFramebufferAttachmentParameterFunc(
			this.Ref(),
		),
	)
}

// GetFramebufferAttachmentParameter calls the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
func (this WebGL2RenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// TryGetFramebufferAttachmentParameter calls the method "WebGL2RenderingContext.getFramebufferAttachmentParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// HasGetProgramParameter returns true if the method "WebGL2RenderingContext.getProgramParameter" exists.
func (this WebGL2RenderingContext) HasGetProgramParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetProgramParameter(
		this.Ref(),
	)
}

// GetProgramParameterFunc returns the method "WebGL2RenderingContext.getProgramParameter".
func (this WebGL2RenderingContext) GetProgramParameterFunc() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetProgramParameterFunc(
			this.Ref(),
		),
	)
}

// GetProgramParameter calls the method "WebGL2RenderingContext.getProgramParameter".
func (this WebGL2RenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		uint32(pname),
	)

	return
}

// TryGetProgramParameter calls the method "WebGL2RenderingContext.getProgramParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetProgramParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(pname),
	)

	return
}

// HasGetProgramInfoLog returns true if the method "WebGL2RenderingContext.getProgramInfoLog" exists.
func (this WebGL2RenderingContext) HasGetProgramInfoLog() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetProgramInfoLog(
		this.Ref(),
	)
}

// GetProgramInfoLogFunc returns the method "WebGL2RenderingContext.getProgramInfoLog".
func (this WebGL2RenderingContext) GetProgramInfoLogFunc() (fn js.Func[func(program WebGLProgram) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetProgramInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetProgramInfoLog calls the method "WebGL2RenderingContext.getProgramInfoLog".
func (this WebGL2RenderingContext) GetProgramInfoLog(program WebGLProgram) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetProgramInfoLog calls the method "WebGL2RenderingContext.getProgramInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetProgramInfoLog(program WebGLProgram) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetProgramInfoLog(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasGetRenderbufferParameter returns true if the method "WebGL2RenderingContext.getRenderbufferParameter" exists.
func (this WebGL2RenderingContext) HasGetRenderbufferParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetRenderbufferParameter(
		this.Ref(),
	)
}

// GetRenderbufferParameterFunc returns the method "WebGL2RenderingContext.getRenderbufferParameter".
func (this WebGL2RenderingContext) GetRenderbufferParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetRenderbufferParameterFunc(
			this.Ref(),
		),
	)
}

// GetRenderbufferParameter calls the method "WebGL2RenderingContext.getRenderbufferParameter".
func (this WebGL2RenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetRenderbufferParameter calls the method "WebGL2RenderingContext.getRenderbufferParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetRenderbufferParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetShaderParameter returns true if the method "WebGL2RenderingContext.getShaderParameter" exists.
func (this WebGL2RenderingContext) HasGetShaderParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetShaderParameter(
		this.Ref(),
	)
}

// GetShaderParameterFunc returns the method "WebGL2RenderingContext.getShaderParameter".
func (this WebGL2RenderingContext) GetShaderParameterFunc() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderParameterFunc(
			this.Ref(),
		),
	)
}

// GetShaderParameter calls the method "WebGL2RenderingContext.getShaderParameter".
func (this WebGL2RenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// TryGetShaderParameter calls the method "WebGL2RenderingContext.getShaderParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// HasGetShaderPrecisionFormat returns true if the method "WebGL2RenderingContext.getShaderPrecisionFormat" exists.
func (this WebGL2RenderingContext) HasGetShaderPrecisionFormat() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetShaderPrecisionFormat(
		this.Ref(),
	)
}

// GetShaderPrecisionFormatFunc returns the method "WebGL2RenderingContext.getShaderPrecisionFormat".
func (this WebGL2RenderingContext) GetShaderPrecisionFormatFunc() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderPrecisionFormatFunc(
			this.Ref(),
		),
	)
}

// GetShaderPrecisionFormat calls the method "WebGL2RenderingContext.getShaderPrecisionFormat".
func (this WebGL2RenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat) {
	bindings.CallWebGL2RenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&ret),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// TryGetShaderPrecisionFormat calls the method "WebGL2RenderingContext.getShaderPrecisionFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderPrecisionFormat(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// HasGetShaderInfoLog returns true if the method "WebGL2RenderingContext.getShaderInfoLog" exists.
func (this WebGL2RenderingContext) HasGetShaderInfoLog() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetShaderInfoLog(
		this.Ref(),
	)
}

// GetShaderInfoLogFunc returns the method "WebGL2RenderingContext.getShaderInfoLog".
func (this WebGL2RenderingContext) GetShaderInfoLogFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderInfoLogFunc(
			this.Ref(),
		),
	)
}

// GetShaderInfoLog calls the method "WebGL2RenderingContext.getShaderInfoLog".
func (this WebGL2RenderingContext) GetShaderInfoLog(shader WebGLShader) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderInfoLog calls the method "WebGL2RenderingContext.getShaderInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderInfoLog(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderInfoLog(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasGetShaderSource returns true if the method "WebGL2RenderingContext.getShaderSource" exists.
func (this WebGL2RenderingContext) HasGetShaderSource() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetShaderSource(
		this.Ref(),
	)
}

// GetShaderSourceFunc returns the method "WebGL2RenderingContext.getShaderSource".
func (this WebGL2RenderingContext) GetShaderSourceFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetShaderSourceFunc(
			this.Ref(),
		),
	)
}

// GetShaderSource calls the method "WebGL2RenderingContext.getShaderSource".
func (this WebGL2RenderingContext) GetShaderSource(shader WebGLShader) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderSource calls the method "WebGL2RenderingContext.getShaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderSource(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasGetTexParameter returns true if the method "WebGL2RenderingContext.getTexParameter" exists.
func (this WebGL2RenderingContext) HasGetTexParameter() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetTexParameter(
		this.Ref(),
	)
}

// GetTexParameterFunc returns the method "WebGL2RenderingContext.getTexParameter".
func (this WebGL2RenderingContext) GetTexParameterFunc() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetTexParameterFunc(
			this.Ref(),
		),
	)
}

// GetTexParameter calls the method "WebGL2RenderingContext.getTexParameter".
func (this WebGL2RenderingContext) GetTexParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetTexParameter calls the method "WebGL2RenderingContext.getTexParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetTexParameter(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetTexParameter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasGetUniform returns true if the method "WebGL2RenderingContext.getUniform" exists.
func (this WebGL2RenderingContext) HasGetUniform() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetUniform(
		this.Ref(),
	)
}

// GetUniformFunc returns the method "WebGL2RenderingContext.getUniform".
func (this WebGL2RenderingContext) GetUniformFunc() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformFunc(
			this.Ref(),
		),
	)
}

// GetUniform calls the method "WebGL2RenderingContext.getUniform".
func (this WebGL2RenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetUniform(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		location.Ref(),
	)

	return
}

// TryGetUniform calls the method "WebGL2RenderingContext.getUniform"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetUniform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		location.Ref(),
	)

	return
}

// HasGetUniformLocation returns true if the method "WebGL2RenderingContext.getUniformLocation" exists.
func (this WebGL2RenderingContext) HasGetUniformLocation() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetUniformLocation(
		this.Ref(),
	)
}

// GetUniformLocationFunc returns the method "WebGL2RenderingContext.getUniformLocation".
func (this WebGL2RenderingContext) GetUniformLocationFunc() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetUniformLocationFunc(
			this.Ref(),
		),
	)
}

// GetUniformLocation calls the method "WebGL2RenderingContext.getUniformLocation".
func (this WebGL2RenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation) {
	bindings.CallWebGL2RenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
		name.Ref(),
	)

	return
}

// TryGetUniformLocation calls the method "WebGL2RenderingContext.getUniformLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetUniformLocation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasGetVertexAttrib returns true if the method "WebGL2RenderingContext.getVertexAttrib" exists.
func (this WebGL2RenderingContext) HasGetVertexAttrib() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetVertexAttrib(
		this.Ref(),
	)
}

// GetVertexAttribFunc returns the method "WebGL2RenderingContext.getVertexAttrib".
func (this WebGL2RenderingContext) GetVertexAttribFunc() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetVertexAttribFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttrib calls the method "WebGL2RenderingContext.getVertexAttrib".
func (this WebGL2RenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(pname),
	)

	return
}

// TryGetVertexAttrib calls the method "WebGL2RenderingContext.getVertexAttrib"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetVertexAttrib(index GLuint, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetVertexAttrib(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasGetVertexAttribOffset returns true if the method "WebGL2RenderingContext.getVertexAttribOffset" exists.
func (this WebGL2RenderingContext) HasGetVertexAttribOffset() bool {
	return js.True == bindings.HasWebGL2RenderingContextGetVertexAttribOffset(
		this.Ref(),
	)
}

// GetVertexAttribOffsetFunc returns the method "WebGL2RenderingContext.getVertexAttribOffset".
func (this WebGL2RenderingContext) GetVertexAttribOffsetFunc() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextGetVertexAttribOffsetFunc(
			this.Ref(),
		),
	)
}

// GetVertexAttribOffset calls the method "WebGL2RenderingContext.getVertexAttribOffset".
func (this WebGL2RenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr) {
	bindings.CallWebGL2RenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		uint32(pname),
	)

	return
}

// TryGetVertexAttribOffset calls the method "WebGL2RenderingContext.getVertexAttribOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetVertexAttribOffset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasHint returns true if the method "WebGL2RenderingContext.hint" exists.
func (this WebGL2RenderingContext) HasHint() bool {
	return js.True == bindings.HasWebGL2RenderingContextHint(
		this.Ref(),
	)
}

// HintFunc returns the method "WebGL2RenderingContext.hint".
func (this WebGL2RenderingContext) HintFunc() (fn js.Func[func(target GLenum, mode GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextHintFunc(
			this.Ref(),
		),
	)
}

// Hint calls the method "WebGL2RenderingContext.hint".
func (this WebGL2RenderingContext) Hint(target GLenum, mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextHint(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(mode),
	)

	return
}

// TryHint calls the method "WebGL2RenderingContext.hint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryHint(target GLenum, mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextHint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(mode),
	)

	return
}

// HasIsBuffer returns true if the method "WebGL2RenderingContext.isBuffer" exists.
func (this WebGL2RenderingContext) HasIsBuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsBuffer(
		this.Ref(),
	)
}

// IsBufferFunc returns the method "WebGL2RenderingContext.isBuffer".
func (this WebGL2RenderingContext) IsBufferFunc() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsBufferFunc(
			this.Ref(),
		),
	)
}

// IsBuffer calls the method "WebGL2RenderingContext.isBuffer".
func (this WebGL2RenderingContext) IsBuffer(buffer WebGLBuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryIsBuffer calls the method "WebGL2RenderingContext.isBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsBuffer(buffer WebGLBuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasIsEnabled returns true if the method "WebGL2RenderingContext.isEnabled" exists.
func (this WebGL2RenderingContext) HasIsEnabled() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsEnabled(
		this.Ref(),
	)
}

// IsEnabledFunc returns the method "WebGL2RenderingContext.isEnabled".
func (this WebGL2RenderingContext) IsEnabledFunc() (fn js.Func[func(cap GLenum) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsEnabledFunc(
			this.Ref(),
		),
	)
}

// IsEnabled calls the method "WebGL2RenderingContext.isEnabled".
func (this WebGL2RenderingContext) IsEnabled(cap GLenum) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsEnabled(
		this.Ref(), js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryIsEnabled calls the method "WebGL2RenderingContext.isEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsEnabled(cap GLenum) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasIsFramebuffer returns true if the method "WebGL2RenderingContext.isFramebuffer" exists.
func (this WebGL2RenderingContext) HasIsFramebuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsFramebuffer(
		this.Ref(),
	)
}

// IsFramebufferFunc returns the method "WebGL2RenderingContext.isFramebuffer".
func (this WebGL2RenderingContext) IsFramebufferFunc() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsFramebufferFunc(
			this.Ref(),
		),
	)
}

// IsFramebuffer calls the method "WebGL2RenderingContext.isFramebuffer".
func (this WebGL2RenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryIsFramebuffer calls the method "WebGL2RenderingContext.isFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsFramebuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasIsProgram returns true if the method "WebGL2RenderingContext.isProgram" exists.
func (this WebGL2RenderingContext) HasIsProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsProgram(
		this.Ref(),
	)
}

// IsProgramFunc returns the method "WebGL2RenderingContext.isProgram".
func (this WebGL2RenderingContext) IsProgramFunc() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsProgramFunc(
			this.Ref(),
		),
	)
}

// IsProgram calls the method "WebGL2RenderingContext.isProgram".
func (this WebGL2RenderingContext) IsProgram(program WebGLProgram) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryIsProgram calls the method "WebGL2RenderingContext.isProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsProgram(program WebGLProgram) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasIsRenderbuffer returns true if the method "WebGL2RenderingContext.isRenderbuffer" exists.
func (this WebGL2RenderingContext) HasIsRenderbuffer() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsRenderbuffer(
		this.Ref(),
	)
}

// IsRenderbufferFunc returns the method "WebGL2RenderingContext.isRenderbuffer".
func (this WebGL2RenderingContext) IsRenderbufferFunc() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsRenderbufferFunc(
			this.Ref(),
		),
	)
}

// IsRenderbuffer calls the method "WebGL2RenderingContext.isRenderbuffer".
func (this WebGL2RenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryIsRenderbuffer calls the method "WebGL2RenderingContext.isRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsRenderbuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasIsShader returns true if the method "WebGL2RenderingContext.isShader" exists.
func (this WebGL2RenderingContext) HasIsShader() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsShader(
		this.Ref(),
	)
}

// IsShaderFunc returns the method "WebGL2RenderingContext.isShader".
func (this WebGL2RenderingContext) IsShaderFunc() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsShaderFunc(
			this.Ref(),
		),
	)
}

// IsShader calls the method "WebGL2RenderingContext.isShader".
func (this WebGL2RenderingContext) IsShader(shader WebGLShader) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsShader(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryIsShader calls the method "WebGL2RenderingContext.isShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsShader(shader WebGLShader) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsShader(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasIsTexture returns true if the method "WebGL2RenderingContext.isTexture" exists.
func (this WebGL2RenderingContext) HasIsTexture() bool {
	return js.True == bindings.HasWebGL2RenderingContextIsTexture(
		this.Ref(),
	)
}

// IsTextureFunc returns the method "WebGL2RenderingContext.isTexture".
func (this WebGL2RenderingContext) IsTextureFunc() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextIsTextureFunc(
			this.Ref(),
		),
	)
}

// IsTexture calls the method "WebGL2RenderingContext.isTexture".
func (this WebGL2RenderingContext) IsTexture(texture WebGLTexture) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsTexture(
		this.Ref(), js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryIsTexture calls the method "WebGL2RenderingContext.isTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsTexture(texture WebGLTexture) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasLineWidth returns true if the method "WebGL2RenderingContext.lineWidth" exists.
func (this WebGL2RenderingContext) HasLineWidth() bool {
	return js.True == bindings.HasWebGL2RenderingContextLineWidth(
		this.Ref(),
	)
}

// LineWidthFunc returns the method "WebGL2RenderingContext.lineWidth".
func (this WebGL2RenderingContext) LineWidthFunc() (fn js.Func[func(width GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextLineWidthFunc(
			this.Ref(),
		),
	)
}

// LineWidth calls the method "WebGL2RenderingContext.lineWidth".
func (this WebGL2RenderingContext) LineWidth(width GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextLineWidth(
		this.Ref(), js.Pointer(&ret),
		float32(width),
	)

	return
}

// TryLineWidth calls the method "WebGL2RenderingContext.lineWidth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryLineWidth(width GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextLineWidth(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(width),
	)

	return
}

// HasLinkProgram returns true if the method "WebGL2RenderingContext.linkProgram" exists.
func (this WebGL2RenderingContext) HasLinkProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextLinkProgram(
		this.Ref(),
	)
}

// LinkProgramFunc returns the method "WebGL2RenderingContext.linkProgram".
func (this WebGL2RenderingContext) LinkProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextLinkProgramFunc(
			this.Ref(),
		),
	)
}

// LinkProgram calls the method "WebGL2RenderingContext.linkProgram".
func (this WebGL2RenderingContext) LinkProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextLinkProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryLinkProgram calls the method "WebGL2RenderingContext.linkProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryLinkProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextLinkProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasPixelStorei returns true if the method "WebGL2RenderingContext.pixelStorei" exists.
func (this WebGL2RenderingContext) HasPixelStorei() bool {
	return js.True == bindings.HasWebGL2RenderingContextPixelStorei(
		this.Ref(),
	)
}

// PixelStoreiFunc returns the method "WebGL2RenderingContext.pixelStorei".
func (this WebGL2RenderingContext) PixelStoreiFunc() (fn js.Func[func(pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPixelStoreiFunc(
			this.Ref(),
		),
	)
}

// PixelStorei calls the method "WebGL2RenderingContext.pixelStorei".
func (this WebGL2RenderingContext) PixelStorei(pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextPixelStorei(
		this.Ref(), js.Pointer(&ret),
		uint32(pname),
		int32(param),
	)

	return
}

// TryPixelStorei calls the method "WebGL2RenderingContext.pixelStorei"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryPixelStorei(pname GLenum, param GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextPixelStorei(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
		int32(param),
	)

	return
}

// HasPolygonOffset returns true if the method "WebGL2RenderingContext.polygonOffset" exists.
func (this WebGL2RenderingContext) HasPolygonOffset() bool {
	return js.True == bindings.HasWebGL2RenderingContextPolygonOffset(
		this.Ref(),
	)
}

// PolygonOffsetFunc returns the method "WebGL2RenderingContext.polygonOffset".
func (this WebGL2RenderingContext) PolygonOffsetFunc() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextPolygonOffsetFunc(
			this.Ref(),
		),
	)
}

// PolygonOffset calls the method "WebGL2RenderingContext.polygonOffset".
func (this WebGL2RenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&ret),
		float32(factor),
		float32(units),
	)

	return
}

// TryPolygonOffset calls the method "WebGL2RenderingContext.polygonOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryPolygonOffset(factor GLfloat, units GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextPolygonOffset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(factor),
		float32(units),
	)

	return
}

// HasRenderbufferStorage returns true if the method "WebGL2RenderingContext.renderbufferStorage" exists.
func (this WebGL2RenderingContext) HasRenderbufferStorage() bool {
	return js.True == bindings.HasWebGL2RenderingContextRenderbufferStorage(
		this.Ref(),
	)
}

// RenderbufferStorageFunc returns the method "WebGL2RenderingContext.renderbufferStorage".
func (this WebGL2RenderingContext) RenderbufferStorageFunc() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextRenderbufferStorageFunc(
			this.Ref(),
		),
	)
}

// RenderbufferStorage calls the method "WebGL2RenderingContext.renderbufferStorage".
func (this WebGL2RenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// TryRenderbufferStorage calls the method "WebGL2RenderingContext.renderbufferStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryRenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextRenderbufferStorage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasSampleCoverage returns true if the method "WebGL2RenderingContext.sampleCoverage" exists.
func (this WebGL2RenderingContext) HasSampleCoverage() bool {
	return js.True == bindings.HasWebGL2RenderingContextSampleCoverage(
		this.Ref(),
	)
}

// SampleCoverageFunc returns the method "WebGL2RenderingContext.sampleCoverage".
func (this WebGL2RenderingContext) SampleCoverageFunc() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextSampleCoverageFunc(
			this.Ref(),
		),
	)
}

// SampleCoverage calls the method "WebGL2RenderingContext.sampleCoverage".
func (this WebGL2RenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&ret),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// TrySampleCoverage calls the method "WebGL2RenderingContext.sampleCoverage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TrySampleCoverage(value GLclampf, invert GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextSampleCoverage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// HasScissor returns true if the method "WebGL2RenderingContext.scissor" exists.
func (this WebGL2RenderingContext) HasScissor() bool {
	return js.True == bindings.HasWebGL2RenderingContextScissor(
		this.Ref(),
	)
}

// ScissorFunc returns the method "WebGL2RenderingContext.scissor".
func (this WebGL2RenderingContext) ScissorFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextScissorFunc(
			this.Ref(),
		),
	)
}

// Scissor calls the method "WebGL2RenderingContext.scissor".
func (this WebGL2RenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextScissor(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryScissor calls the method "WebGL2RenderingContext.scissor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryScissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextScissor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasShaderSource returns true if the method "WebGL2RenderingContext.shaderSource" exists.
func (this WebGL2RenderingContext) HasShaderSource() bool {
	return js.True == bindings.HasWebGL2RenderingContextShaderSource(
		this.Ref(),
	)
}

// ShaderSourceFunc returns the method "WebGL2RenderingContext.shaderSource".
func (this WebGL2RenderingContext) ShaderSourceFunc() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextShaderSourceFunc(
			this.Ref(),
		),
	)
}

// ShaderSource calls the method "WebGL2RenderingContext.shaderSource".
func (this WebGL2RenderingContext) ShaderSource(shader WebGLShader, source js.String) (ret js.Void) {
	bindings.CallWebGL2RenderingContextShaderSource(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// TryShaderSource calls the method "WebGL2RenderingContext.shaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryShaderSource(shader WebGLShader, source js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextShaderSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// HasStencilFunc returns true if the method "WebGL2RenderingContext.stencilFunc" exists.
func (this WebGL2RenderingContext) HasStencilFunc() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilFunc(
		this.Ref(),
	)
}

// StencilFuncFunc returns the method "WebGL2RenderingContext.stencilFunc".
func (this WebGL2RenderingContext) StencilFuncFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilFuncFunc(
			this.Ref(),
		),
	)
}

// StencilFunc calls the method "WebGL2RenderingContext.stencilFunc".
func (this WebGL2RenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilFunc(
		this.Ref(), js.Pointer(&ret),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// TryStencilFunc calls the method "WebGL2RenderingContext.stencilFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilFunc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasStencilFuncSeparate returns true if the method "WebGL2RenderingContext.stencilFuncSeparate" exists.
func (this WebGL2RenderingContext) HasStencilFuncSeparate() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilFuncSeparate(
		this.Ref(),
	)
}

// StencilFuncSeparateFunc returns the method "WebGL2RenderingContext.stencilFuncSeparate".
func (this WebGL2RenderingContext) StencilFuncSeparateFunc() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilFuncSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilFuncSeparate calls the method "WebGL2RenderingContext.stencilFuncSeparate".
func (this WebGL2RenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// TryStencilFuncSeparate calls the method "WebGL2RenderingContext.stencilFuncSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilFuncSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasStencilMask returns true if the method "WebGL2RenderingContext.stencilMask" exists.
func (this WebGL2RenderingContext) HasStencilMask() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilMask(
		this.Ref(),
	)
}

// StencilMaskFunc returns the method "WebGL2RenderingContext.stencilMask".
func (this WebGL2RenderingContext) StencilMaskFunc() (fn js.Func[func(mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilMaskFunc(
			this.Ref(),
		),
	)
}

// StencilMask calls the method "WebGL2RenderingContext.stencilMask".
func (this WebGL2RenderingContext) StencilMask(mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilMask(
		this.Ref(), js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryStencilMask calls the method "WebGL2RenderingContext.stencilMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilMask(mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilMask(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasStencilMaskSeparate returns true if the method "WebGL2RenderingContext.stencilMaskSeparate" exists.
func (this WebGL2RenderingContext) HasStencilMaskSeparate() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilMaskSeparate(
		this.Ref(),
	)
}

// StencilMaskSeparateFunc returns the method "WebGL2RenderingContext.stencilMaskSeparate".
func (this WebGL2RenderingContext) StencilMaskSeparateFunc() (fn js.Func[func(face GLenum, mask GLuint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilMaskSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilMaskSeparate calls the method "WebGL2RenderingContext.stencilMaskSeparate".
func (this WebGL2RenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(mask),
	)

	return
}

// TryStencilMaskSeparate calls the method "WebGL2RenderingContext.stencilMaskSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilMaskSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(mask),
	)

	return
}

// HasStencilOp returns true if the method "WebGL2RenderingContext.stencilOp" exists.
func (this WebGL2RenderingContext) HasStencilOp() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilOp(
		this.Ref(),
	)
}

// StencilOpFunc returns the method "WebGL2RenderingContext.stencilOp".
func (this WebGL2RenderingContext) StencilOpFunc() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilOpFunc(
			this.Ref(),
		),
	)
}

// StencilOp calls the method "WebGL2RenderingContext.stencilOp".
func (this WebGL2RenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilOp(
		this.Ref(), js.Pointer(&ret),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// TryStencilOp calls the method "WebGL2RenderingContext.stencilOp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilOp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasStencilOpSeparate returns true if the method "WebGL2RenderingContext.stencilOpSeparate" exists.
func (this WebGL2RenderingContext) HasStencilOpSeparate() bool {
	return js.True == bindings.HasWebGL2RenderingContextStencilOpSeparate(
		this.Ref(),
	)
}

// StencilOpSeparateFunc returns the method "WebGL2RenderingContext.stencilOpSeparate".
func (this WebGL2RenderingContext) StencilOpSeparateFunc() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextStencilOpSeparateFunc(
			this.Ref(),
		),
	)
}

// StencilOpSeparate calls the method "WebGL2RenderingContext.stencilOpSeparate".
func (this WebGL2RenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&ret),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// TryStencilOpSeparate calls the method "WebGL2RenderingContext.stencilOpSeparate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilOpSeparate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasTexParameterf returns true if the method "WebGL2RenderingContext.texParameterf" exists.
func (this WebGL2RenderingContext) HasTexParameterf() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexParameterf(
		this.Ref(),
	)
}

// TexParameterfFunc returns the method "WebGL2RenderingContext.texParameterf".
func (this WebGL2RenderingContext) TexParameterfFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexParameterfFunc(
			this.Ref(),
		),
	)
}

// TexParameterf calls the method "WebGL2RenderingContext.texParameterf".
func (this WebGL2RenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexParameterf(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// TryTexParameterf calls the method "WebGL2RenderingContext.texParameterf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexParameterf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// HasTexParameteri returns true if the method "WebGL2RenderingContext.texParameteri" exists.
func (this WebGL2RenderingContext) HasTexParameteri() bool {
	return js.True == bindings.HasWebGL2RenderingContextTexParameteri(
		this.Ref(),
	)
}

// TexParameteriFunc returns the method "WebGL2RenderingContext.texParameteri".
func (this WebGL2RenderingContext) TexParameteriFunc() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextTexParameteriFunc(
			this.Ref(),
		),
	)
}

// TexParameteri calls the method "WebGL2RenderingContext.texParameteri".
func (this WebGL2RenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexParameteri(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// TryTexParameteri calls the method "WebGL2RenderingContext.texParameteri"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryTexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextTexParameteri(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// HasUniform1f returns true if the method "WebGL2RenderingContext.uniform1f" exists.
func (this WebGL2RenderingContext) HasUniform1f() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1f(
		this.Ref(),
	)
}

// Uniform1fFunc returns the method "WebGL2RenderingContext.uniform1f".
func (this WebGL2RenderingContext) Uniform1fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1fFunc(
			this.Ref(),
		),
	)
}

// Uniform1f calls the method "WebGL2RenderingContext.uniform1f".
func (this WebGL2RenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
	)

	return
}

// TryUniform1f calls the method "WebGL2RenderingContext.uniform1f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
	)

	return
}

// HasUniform2f returns true if the method "WebGL2RenderingContext.uniform2f" exists.
func (this WebGL2RenderingContext) HasUniform2f() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2f(
		this.Ref(),
	)
}

// Uniform2fFunc returns the method "WebGL2RenderingContext.uniform2f".
func (this WebGL2RenderingContext) Uniform2fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2fFunc(
			this.Ref(),
		),
	)
}

// Uniform2f calls the method "WebGL2RenderingContext.uniform2f".
func (this WebGL2RenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// TryUniform2f calls the method "WebGL2RenderingContext.uniform2f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// HasUniform3f returns true if the method "WebGL2RenderingContext.uniform3f" exists.
func (this WebGL2RenderingContext) HasUniform3f() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3f(
		this.Ref(),
	)
}

// Uniform3fFunc returns the method "WebGL2RenderingContext.uniform3f".
func (this WebGL2RenderingContext) Uniform3fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3fFunc(
			this.Ref(),
		),
	)
}

// Uniform3f calls the method "WebGL2RenderingContext.uniform3f".
func (this WebGL2RenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TryUniform3f calls the method "WebGL2RenderingContext.uniform3f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasUniform4f returns true if the method "WebGL2RenderingContext.uniform4f" exists.
func (this WebGL2RenderingContext) HasUniform4f() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4f(
		this.Ref(),
	)
}

// Uniform4fFunc returns the method "WebGL2RenderingContext.uniform4f".
func (this WebGL2RenderingContext) Uniform4fFunc() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4fFunc(
			this.Ref(),
		),
	)
}

// Uniform4f calls the method "WebGL2RenderingContext.uniform4f".
func (this WebGL2RenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4f(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// TryUniform4f calls the method "WebGL2RenderingContext.uniform4f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasUniform1i returns true if the method "WebGL2RenderingContext.uniform1i" exists.
func (this WebGL2RenderingContext) HasUniform1i() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform1i(
		this.Ref(),
	)
}

// Uniform1iFunc returns the method "WebGL2RenderingContext.uniform1i".
func (this WebGL2RenderingContext) Uniform1iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform1iFunc(
			this.Ref(),
		),
	)
}

// Uniform1i calls the method "WebGL2RenderingContext.uniform1i".
func (this WebGL2RenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
	)

	return
}

// TryUniform1i calls the method "WebGL2RenderingContext.uniform1i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform1i(location WebGLUniformLocation, x GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform1i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
	)

	return
}

// HasUniform2i returns true if the method "WebGL2RenderingContext.uniform2i" exists.
func (this WebGL2RenderingContext) HasUniform2i() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform2i(
		this.Ref(),
	)
}

// Uniform2iFunc returns the method "WebGL2RenderingContext.uniform2i".
func (this WebGL2RenderingContext) Uniform2iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform2iFunc(
			this.Ref(),
		),
	)
}

// Uniform2i calls the method "WebGL2RenderingContext.uniform2i".
func (this WebGL2RenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// TryUniform2i calls the method "WebGL2RenderingContext.uniform2i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform2i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasUniform3i returns true if the method "WebGL2RenderingContext.uniform3i" exists.
func (this WebGL2RenderingContext) HasUniform3i() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform3i(
		this.Ref(),
	)
}

// Uniform3iFunc returns the method "WebGL2RenderingContext.uniform3i".
func (this WebGL2RenderingContext) Uniform3iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform3iFunc(
			this.Ref(),
		),
	)
}

// Uniform3i calls the method "WebGL2RenderingContext.uniform3i".
func (this WebGL2RenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// TryUniform3i calls the method "WebGL2RenderingContext.uniform3i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform3i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// HasUniform4i returns true if the method "WebGL2RenderingContext.uniform4i" exists.
func (this WebGL2RenderingContext) HasUniform4i() bool {
	return js.True == bindings.HasWebGL2RenderingContextUniform4i(
		this.Ref(),
	)
}

// Uniform4iFunc returns the method "WebGL2RenderingContext.uniform4i".
func (this WebGL2RenderingContext) Uniform4iFunc() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUniform4iFunc(
			this.Ref(),
		),
	)
}

// Uniform4i calls the method "WebGL2RenderingContext.uniform4i".
func (this WebGL2RenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4i(
		this.Ref(), js.Pointer(&ret),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// TryUniform4i calls the method "WebGL2RenderingContext.uniform4i"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUniform4i(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasUseProgram returns true if the method "WebGL2RenderingContext.useProgram" exists.
func (this WebGL2RenderingContext) HasUseProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextUseProgram(
		this.Ref(),
	)
}

// UseProgramFunc returns the method "WebGL2RenderingContext.useProgram".
func (this WebGL2RenderingContext) UseProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextUseProgramFunc(
			this.Ref(),
		),
	)
}

// UseProgram calls the method "WebGL2RenderingContext.useProgram".
func (this WebGL2RenderingContext) UseProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUseProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryUseProgram calls the method "WebGL2RenderingContext.useProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUseProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUseProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasValidateProgram returns true if the method "WebGL2RenderingContext.validateProgram" exists.
func (this WebGL2RenderingContext) HasValidateProgram() bool {
	return js.True == bindings.HasWebGL2RenderingContextValidateProgram(
		this.Ref(),
	)
}

// ValidateProgramFunc returns the method "WebGL2RenderingContext.validateProgram".
func (this WebGL2RenderingContext) ValidateProgramFunc() (fn js.Func[func(program WebGLProgram)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextValidateProgramFunc(
			this.Ref(),
		),
	)
}

// ValidateProgram calls the method "WebGL2RenderingContext.validateProgram".
func (this WebGL2RenderingContext) ValidateProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextValidateProgram(
		this.Ref(), js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryValidateProgram calls the method "WebGL2RenderingContext.validateProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryValidateProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextValidateProgram(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasVertexAttrib1f returns true if the method "WebGL2RenderingContext.vertexAttrib1f" exists.
func (this WebGL2RenderingContext) HasVertexAttrib1f() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib1f(
		this.Ref(),
	)
}

// VertexAttrib1fFunc returns the method "WebGL2RenderingContext.vertexAttrib1f".
func (this WebGL2RenderingContext) VertexAttrib1fFunc() (fn js.Func[func(index GLuint, x GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib1fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1f calls the method "WebGL2RenderingContext.vertexAttrib1f".
func (this WebGL2RenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
	)

	return
}

// TryVertexAttrib1f calls the method "WebGL2RenderingContext.vertexAttrib1f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib1f(index GLuint, x GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib1f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
	)

	return
}

// HasVertexAttrib2f returns true if the method "WebGL2RenderingContext.vertexAttrib2f" exists.
func (this WebGL2RenderingContext) HasVertexAttrib2f() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib2f(
		this.Ref(),
	)
}

// VertexAttrib2fFunc returns the method "WebGL2RenderingContext.vertexAttrib2f".
func (this WebGL2RenderingContext) VertexAttrib2fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib2fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2f calls the method "WebGL2RenderingContext.vertexAttrib2f".
func (this WebGL2RenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// TryVertexAttrib2f calls the method "WebGL2RenderingContext.vertexAttrib2f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib2f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// HasVertexAttrib3f returns true if the method "WebGL2RenderingContext.vertexAttrib3f" exists.
func (this WebGL2RenderingContext) HasVertexAttrib3f() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib3f(
		this.Ref(),
	)
}

// VertexAttrib3fFunc returns the method "WebGL2RenderingContext.vertexAttrib3f".
func (this WebGL2RenderingContext) VertexAttrib3fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib3fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3f calls the method "WebGL2RenderingContext.vertexAttrib3f".
func (this WebGL2RenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TryVertexAttrib3f calls the method "WebGL2RenderingContext.vertexAttrib3f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib3f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasVertexAttrib4f returns true if the method "WebGL2RenderingContext.vertexAttrib4f" exists.
func (this WebGL2RenderingContext) HasVertexAttrib4f() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib4f(
		this.Ref(),
	)
}

// VertexAttrib4fFunc returns the method "WebGL2RenderingContext.vertexAttrib4f".
func (this WebGL2RenderingContext) VertexAttrib4fFunc() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib4fFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4f calls the method "WebGL2RenderingContext.vertexAttrib4f".
func (this WebGL2RenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// TryVertexAttrib4f calls the method "WebGL2RenderingContext.vertexAttrib4f"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib4f(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasVertexAttrib1fv returns true if the method "WebGL2RenderingContext.vertexAttrib1fv" exists.
func (this WebGL2RenderingContext) HasVertexAttrib1fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib1fv(
		this.Ref(),
	)
}

// VertexAttrib1fvFunc returns the method "WebGL2RenderingContext.vertexAttrib1fv".
func (this WebGL2RenderingContext) VertexAttrib1fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib1fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib1fv calls the method "WebGL2RenderingContext.vertexAttrib1fv".
func (this WebGL2RenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib1fv calls the method "WebGL2RenderingContext.vertexAttrib1fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib1fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib1fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib2fv returns true if the method "WebGL2RenderingContext.vertexAttrib2fv" exists.
func (this WebGL2RenderingContext) HasVertexAttrib2fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib2fv(
		this.Ref(),
	)
}

// VertexAttrib2fvFunc returns the method "WebGL2RenderingContext.vertexAttrib2fv".
func (this WebGL2RenderingContext) VertexAttrib2fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib2fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib2fv calls the method "WebGL2RenderingContext.vertexAttrib2fv".
func (this WebGL2RenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib2fv calls the method "WebGL2RenderingContext.vertexAttrib2fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib2fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib2fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib3fv returns true if the method "WebGL2RenderingContext.vertexAttrib3fv" exists.
func (this WebGL2RenderingContext) HasVertexAttrib3fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib3fv(
		this.Ref(),
	)
}

// VertexAttrib3fvFunc returns the method "WebGL2RenderingContext.vertexAttrib3fv".
func (this WebGL2RenderingContext) VertexAttrib3fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib3fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib3fv calls the method "WebGL2RenderingContext.vertexAttrib3fv".
func (this WebGL2RenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib3fv calls the method "WebGL2RenderingContext.vertexAttrib3fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib3fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib3fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttrib4fv returns true if the method "WebGL2RenderingContext.vertexAttrib4fv" exists.
func (this WebGL2RenderingContext) HasVertexAttrib4fv() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttrib4fv(
		this.Ref(),
	)
}

// VertexAttrib4fvFunc returns the method "WebGL2RenderingContext.vertexAttrib4fv".
func (this WebGL2RenderingContext) VertexAttrib4fvFunc() (fn js.Func[func(index GLuint, values Float32List)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttrib4fvFunc(
			this.Ref(),
		),
	)
}

// VertexAttrib4fv calls the method "WebGL2RenderingContext.vertexAttrib4fv".
func (this WebGL2RenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		values.Ref(),
	)

	return
}

// TryVertexAttrib4fv calls the method "WebGL2RenderingContext.vertexAttrib4fv"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttrib4fv(index GLuint, values Float32List) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttrib4fv(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasVertexAttribPointer returns true if the method "WebGL2RenderingContext.vertexAttribPointer" exists.
func (this WebGL2RenderingContext) HasVertexAttribPointer() bool {
	return js.True == bindings.HasWebGL2RenderingContextVertexAttribPointer(
		this.Ref(),
	)
}

// VertexAttribPointerFunc returns the method "WebGL2RenderingContext.vertexAttribPointer".
func (this WebGL2RenderingContext) VertexAttribPointerFunc() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextVertexAttribPointerFunc(
			this.Ref(),
		),
	)
}

// VertexAttribPointer calls the method "WebGL2RenderingContext.vertexAttribPointer".
func (this WebGL2RenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// TryVertexAttribPointer calls the method "WebGL2RenderingContext.vertexAttribPointer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryVertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextVertexAttribPointer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// HasViewport returns true if the method "WebGL2RenderingContext.viewport" exists.
func (this WebGL2RenderingContext) HasViewport() bool {
	return js.True == bindings.HasWebGL2RenderingContextViewport(
		this.Ref(),
	)
}

// ViewportFunc returns the method "WebGL2RenderingContext.viewport".
func (this WebGL2RenderingContext) ViewportFunc() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextViewportFunc(
			this.Ref(),
		),
	)
}

// Viewport calls the method "WebGL2RenderingContext.viewport".
func (this WebGL2RenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextViewport(
		this.Ref(), js.Pointer(&ret),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// TryViewport calls the method "WebGL2RenderingContext.viewport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryViewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextViewport(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasMakeXRCompatible returns true if the method "WebGL2RenderingContext.makeXRCompatible" exists.
func (this WebGL2RenderingContext) HasMakeXRCompatible() bool {
	return js.True == bindings.HasWebGL2RenderingContextMakeXRCompatible(
		this.Ref(),
	)
}

// MakeXRCompatibleFunc returns the method "WebGL2RenderingContext.makeXRCompatible".
func (this WebGL2RenderingContext) MakeXRCompatibleFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.WebGL2RenderingContextMakeXRCompatibleFunc(
			this.Ref(),
		),
	)
}

// MakeXRCompatible calls the method "WebGL2RenderingContext.makeXRCompatible".
func (this WebGL2RenderingContext) MakeXRCompatible() (ret js.Promise[js.Void]) {
	bindings.CallWebGL2RenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMakeXRCompatible calls the method "WebGL2RenderingContext.makeXRCompatible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryMakeXRCompatible() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextMakeXRCompatible(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this GPUBuffer) Size() (ret GPUSize64Out, ok bool) {
	ok = js.True == bindings.GetGPUBufferSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Usage returns the value of property "GPUBuffer.usage".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) Usage() (ret GPUFlagsConstant, ok bool) {
	ok = js.True == bindings.GetGPUBufferUsage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MapState returns the value of property "GPUBuffer.mapState".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) MapState() (ret GPUBufferMapState, ok bool) {
	ok = js.True == bindings.GetGPUBufferMapState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUBuffer.label".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUBufferLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUBuffer.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBuffer) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBufferLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasMapAsync returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasMapAsync() bool {
	return js.True == bindings.HasGPUBufferMapAsync(
		this.Ref(),
	)
}

// MapAsyncFunc returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsyncFunc() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsyncFunc(
			this.Ref(),
		),
	)
}

// MapAsync calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		float64(offset),
		float64(size),
	)

	return
}

// TryMapAsync calls the method "GPUBuffer.mapAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryMapAsync(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferMapAsync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		float64(offset),
		float64(size),
	)

	return
}

// HasMapAsync1 returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasMapAsync1() bool {
	return js.True == bindings.HasGPUBufferMapAsync1(
		this.Ref(),
	)
}

// MapAsync1Func returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync1Func() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsync1Func(
			this.Ref(),
		),
	)
}

// MapAsync1 calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync1(mode GPUMapModeFlags, offset GPUSize64) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync1(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		float64(offset),
	)

	return
}

// TryMapAsync1 calls the method "GPUBuffer.mapAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryMapAsync1(mode GPUMapModeFlags, offset GPUSize64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferMapAsync1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		float64(offset),
	)

	return
}

// HasMapAsync2 returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasMapAsync2() bool {
	return js.True == bindings.HasGPUBufferMapAsync2(
		this.Ref(),
	)
}

// MapAsync2Func returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync2Func() (fn js.Func[func(mode GPUMapModeFlags) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUBufferMapAsync2Func(
			this.Ref(),
		),
	)
}

// MapAsync2 calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync2(mode GPUMapModeFlags) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync2(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryMapAsync2 calls the method "GPUBuffer.mapAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryMapAsync2(mode GPUMapModeFlags) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferMapAsync2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasGetMappedRange returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasGetMappedRange() bool {
	return js.True == bindings.HasGPUBufferGetMappedRange(
		this.Ref(),
	)
}

// GetMappedRangeFunc returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRangeFunc() (fn js.Func[func(offset GPUSize64, size GPUSize64) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRangeFunc(
			this.Ref(),
		),
	)
}

// GetMappedRange calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange(offset GPUSize64, size GPUSize64) (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange(
		this.Ref(), js.Pointer(&ret),
		float64(offset),
		float64(size),
	)

	return
}

// TryGetMappedRange calls the method "GPUBuffer.getMappedRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryGetMappedRange(offset GPUSize64, size GPUSize64) (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferGetMappedRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
		float64(size),
	)

	return
}

// HasGetMappedRange1 returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasGetMappedRange1() bool {
	return js.True == bindings.HasGPUBufferGetMappedRange1(
		this.Ref(),
	)
}

// GetMappedRange1Func returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange1Func() (fn js.Func[func(offset GPUSize64) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRange1Func(
			this.Ref(),
		),
	)
}

// GetMappedRange1 calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange1(offset GPUSize64) (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange1(
		this.Ref(), js.Pointer(&ret),
		float64(offset),
	)

	return
}

// TryGetMappedRange1 calls the method "GPUBuffer.getMappedRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryGetMappedRange1(offset GPUSize64) (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferGetMappedRange1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
	)

	return
}

// HasGetMappedRange2 returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasGetMappedRange2() bool {
	return js.True == bindings.HasGPUBufferGetMappedRange2(
		this.Ref(),
	)
}

// GetMappedRange2Func returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange2Func() (fn js.Func[func() js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.GPUBufferGetMappedRange2Func(
			this.Ref(),
		),
	)
}

// GetMappedRange2 calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange2() (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetMappedRange2 calls the method "GPUBuffer.getMappedRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryGetMappedRange2() (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferGetMappedRange2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnmap returns true if the method "GPUBuffer.unmap" exists.
func (this GPUBuffer) HasUnmap() bool {
	return js.True == bindings.HasGPUBufferUnmap(
		this.Ref(),
	)
}

// UnmapFunc returns the method "GPUBuffer.unmap".
func (this GPUBuffer) UnmapFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUBufferUnmapFunc(
			this.Ref(),
		),
	)
}

// Unmap calls the method "GPUBuffer.unmap".
func (this GPUBuffer) Unmap() (ret js.Void) {
	bindings.CallGPUBufferUnmap(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnmap calls the method "GPUBuffer.unmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryUnmap() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferUnmap(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDestroy returns true if the method "GPUBuffer.destroy" exists.
func (this GPUBuffer) HasDestroy() bool {
	return js.True == bindings.HasGPUBufferDestroy(
		this.Ref(),
	)
}

// DestroyFunc returns the method "GPUBuffer.destroy".
func (this GPUBuffer) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUBufferDestroyFunc(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUBuffer.destroy".
func (this GPUBuffer) Destroy() (ret js.Void) {
	bindings.CallGPUBufferDestroy(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUBuffer.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferDestroy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_MappedAtCreation MUST be set to true to make this field effective.
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

// New creates a new GPUBufferDescriptor in the application heap.
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
// It returns ok=false if there is no such property.
func (this GPUTextureView) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUTextureViewLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUTextureView.label" to val.
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
	//
	// NOTE: FFI_USE_BaseMipLevel MUST be set to true to make this field effective.
	BaseMipLevel GPUIntegerCoordinate
	// MipLevelCount is "GPUTextureViewDescriptor.mipLevelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_MipLevelCount MUST be set to true to make this field effective.
	MipLevelCount GPUIntegerCoordinate
	// BaseArrayLayer is "GPUTextureViewDescriptor.baseArrayLayer"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_BaseArrayLayer MUST be set to true to make this field effective.
	BaseArrayLayer GPUIntegerCoordinate
	// ArrayLayerCount is "GPUTextureViewDescriptor.arrayLayerCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ArrayLayerCount MUST be set to true to make this field effective.
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

// New creates a new GPUTextureViewDescriptor in the application heap.
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
// It returns ok=false if there is no such property.
func (this GPUTexture) Width() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "GPUTexture.height".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Height() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DepthOrArrayLayers returns the value of property "GPUTexture.depthOrArrayLayers".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) DepthOrArrayLayers() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureDepthOrArrayLayers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MipLevelCount returns the value of property "GPUTexture.mipLevelCount".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) MipLevelCount() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureMipLevelCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SampleCount returns the value of property "GPUTexture.sampleCount".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) SampleCount() (ret GPUSize32Out, ok bool) {
	ok = js.True == bindings.GetGPUTextureSampleCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dimension returns the value of property "GPUTexture.dimension".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Dimension() (ret GPUTextureDimension, ok bool) {
	ok = js.True == bindings.GetGPUTextureDimension(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Format returns the value of property "GPUTexture.format".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Format() (ret GPUTextureFormat, ok bool) {
	ok = js.True == bindings.GetGPUTextureFormat(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Usage returns the value of property "GPUTexture.usage".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Usage() (ret GPUFlagsConstant, ok bool) {
	ok = js.True == bindings.GetGPUTextureUsage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUTexture.label".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUTextureLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUTexture.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUTexture) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUTextureLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasCreateView returns true if the method "GPUTexture.createView" exists.
func (this GPUTexture) HasCreateView() bool {
	return js.True == bindings.HasGPUTextureCreateView(
		this.Ref(),
	)
}

// CreateViewFunc returns the method "GPUTexture.createView".
func (this GPUTexture) CreateViewFunc() (fn js.Func[func(descriptor GPUTextureViewDescriptor) GPUTextureView]) {
	return fn.FromRef(
		bindings.GPUTextureCreateViewFunc(
			this.Ref(),
		),
	)
}

// CreateView calls the method "GPUTexture.createView".
func (this GPUTexture) CreateView(descriptor GPUTextureViewDescriptor) (ret GPUTextureView) {
	bindings.CallGPUTextureCreateView(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateView calls the method "GPUTexture.createView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryCreateView(descriptor GPUTextureViewDescriptor) (ret GPUTextureView, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureCreateView(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateView1 returns true if the method "GPUTexture.createView" exists.
func (this GPUTexture) HasCreateView1() bool {
	return js.True == bindings.HasGPUTextureCreateView1(
		this.Ref(),
	)
}

// CreateView1Func returns the method "GPUTexture.createView".
func (this GPUTexture) CreateView1Func() (fn js.Func[func() GPUTextureView]) {
	return fn.FromRef(
		bindings.GPUTextureCreateView1Func(
			this.Ref(),
		),
	)
}

// CreateView1 calls the method "GPUTexture.createView".
func (this GPUTexture) CreateView1() (ret GPUTextureView) {
	bindings.CallGPUTextureCreateView1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateView1 calls the method "GPUTexture.createView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryCreateView1() (ret GPUTextureView, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureCreateView1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDestroy returns true if the method "GPUTexture.destroy" exists.
func (this GPUTexture) HasDestroy() bool {
	return js.True == bindings.HasGPUTextureDestroy(
		this.Ref(),
	)
}

// DestroyFunc returns the method "GPUTexture.destroy".
func (this GPUTexture) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUTextureDestroyFunc(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUTexture.destroy".
func (this GPUTexture) Destroy() (ret js.Void) {
	bindings.CallGPUTextureDestroy(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUTexture.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureDestroy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
