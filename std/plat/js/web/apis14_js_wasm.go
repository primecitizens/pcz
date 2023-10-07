// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *VideoColorSpaceInit) UpdateFrom(ref js.Ref) {
	bindings.VideoColorSpaceInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoColorSpaceInit) Update(ref js.Ref) {
	bindings.VideoColorSpaceInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoColorSpaceInit) FreeMembers(recursive bool) {
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
func (p *VideoFrameBufferInit) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameBufferInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoFrameBufferInit) Update(ref js.Ref) {
	bindings.VideoFrameBufferInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoFrameBufferInit) FreeMembers(recursive bool) {
	js.Free(
		p.Layout.Ref(),
		p.Transfer.Ref(),
	)
	p.Layout = p.Layout.FromRef(js.Undefined)
	p.Transfer = p.Transfer.FromRef(js.Undefined)
	if recursive {
		p.VisibleRect.FreeMembers(true)
		p.ColorSpace.FreeMembers(true)
	}
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
func (p *VideoFrameCopyToOptions) UpdateFrom(ref js.Ref) {
	bindings.VideoFrameCopyToOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoFrameCopyToOptions) Update(ref js.Ref) {
	bindings.VideoFrameCopyToOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoFrameCopyToOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Layout.Ref(),
	)
	p.Layout = p.Layout.FromRef(js.Undefined)
	if recursive {
		p.Rect.FreeMembers(true)
	}
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
	this.ref.Once()
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
	this.ref.Free()
}

// Primaries returns the value of property "VideoColorSpace.primaries".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Primaries() (ret VideoColorPrimaries, ok bool) {
	ok = js.True == bindings.GetVideoColorSpacePrimaries(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transfer returns the value of property "VideoColorSpace.transfer".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Transfer() (ret VideoTransferCharacteristics, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceTransfer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "VideoColorSpace.matrix".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) Matrix() (ret VideoMatrixCoefficients, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FullRange returns the value of property "VideoColorSpace.fullRange".
//
// It returns ok=false if there is no such property.
func (this VideoColorSpace) FullRange() (ret bool, ok bool) {
	ok = js.True == bindings.GetVideoColorSpaceFullRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "VideoColorSpace.toJSON" exists.
func (this VideoColorSpace) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncVideoColorSpaceToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "VideoColorSpace.toJSON".
func (this VideoColorSpace) FuncToJSON() (fn js.Func[func() VideoColorSpaceInit]) {
	bindings.FuncVideoColorSpaceToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "VideoColorSpace.toJSON".
func (this VideoColorSpace) ToJSON() (ret VideoColorSpaceInit) {
	bindings.CallVideoColorSpaceToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "VideoColorSpace.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoColorSpace) TryToJSON() (ret VideoColorSpaceInit, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoColorSpaceToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Format returns the value of property "VideoFrame.format".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) Format() (ret VideoPixelFormat, ok bool) {
	ok = js.True == bindings.GetVideoFrameFormat(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CodedWidth returns the value of property "VideoFrame.codedWidth".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CodedHeight returns the value of property "VideoFrame.codedHeight".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CodedRect returns the value of property "VideoFrame.codedRect".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) CodedRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetVideoFrameCodedRect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VisibleRect returns the value of property "VideoFrame.visibleRect".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) VisibleRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetVideoFrameVisibleRect(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DisplayWidth returns the value of property "VideoFrame.displayWidth".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) DisplayWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameDisplayWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DisplayHeight returns the value of property "VideoFrame.displayHeight".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) DisplayHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoFrameDisplayHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "VideoFrame.duration".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) Duration() (ret uint64, ok bool) {
	ok = js.True == bindings.GetVideoFrameDuration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "VideoFrame.timestamp".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) Timestamp() (ret int64, ok bool) {
	ok = js.True == bindings.GetVideoFrameTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColorSpace returns the value of property "VideoFrame.colorSpace".
//
// It returns ok=false if there is no such property.
func (this VideoFrame) ColorSpace() (ret VideoColorSpace, ok bool) {
	ok = js.True == bindings.GetVideoFrameColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncMetadata returns true if the method "VideoFrame.metadata" exists.
func (this VideoFrame) HasFuncMetadata() bool {
	return js.True == bindings.HasFuncVideoFrameMetadata(
		this.ref,
	)
}

// FuncMetadata returns the method "VideoFrame.metadata".
func (this VideoFrame) FuncMetadata() (fn js.Func[func() VideoFrameMetadata]) {
	bindings.FuncVideoFrameMetadata(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Metadata calls the method "VideoFrame.metadata".
func (this VideoFrame) Metadata() (ret VideoFrameMetadata) {
	bindings.CallVideoFrameMetadata(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMetadata calls the method "VideoFrame.metadata"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryMetadata() (ret VideoFrameMetadata, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameMetadata(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAllocationSize returns true if the method "VideoFrame.allocationSize" exists.
func (this VideoFrame) HasFuncAllocationSize() bool {
	return js.True == bindings.HasFuncVideoFrameAllocationSize(
		this.ref,
	)
}

// FuncAllocationSize returns the method "VideoFrame.allocationSize".
func (this VideoFrame) FuncAllocationSize() (fn js.Func[func(options VideoFrameCopyToOptions) uint32]) {
	bindings.FuncVideoFrameAllocationSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AllocationSize calls the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSize(options VideoFrameCopyToOptions) (ret uint32) {
	bindings.CallVideoFrameAllocationSize(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAllocationSize calls the method "VideoFrame.allocationSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryAllocationSize(options VideoFrameCopyToOptions) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameAllocationSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncAllocationSize1 returns true if the method "VideoFrame.allocationSize" exists.
func (this VideoFrame) HasFuncAllocationSize1() bool {
	return js.True == bindings.HasFuncVideoFrameAllocationSize1(
		this.ref,
	)
}

// FuncAllocationSize1 returns the method "VideoFrame.allocationSize".
func (this VideoFrame) FuncAllocationSize1() (fn js.Func[func() uint32]) {
	bindings.FuncVideoFrameAllocationSize1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AllocationSize1 calls the method "VideoFrame.allocationSize".
func (this VideoFrame) AllocationSize1() (ret uint32) {
	bindings.CallVideoFrameAllocationSize1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllocationSize1 calls the method "VideoFrame.allocationSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryAllocationSize1() (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameAllocationSize1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCopyTo returns true if the method "VideoFrame.copyTo" exists.
func (this VideoFrame) HasFuncCopyTo() bool {
	return js.True == bindings.HasFuncVideoFrameCopyTo(
		this.ref,
	)
}

// FuncCopyTo returns the method "VideoFrame.copyTo".
func (this VideoFrame) FuncCopyTo() (fn js.Func[func(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) js.Promise[js.Array[PlaneLayout]]]) {
	bindings.FuncVideoFrameCopyTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTo calls the method "VideoFrame.copyTo".
func (this VideoFrame) CopyTo(destination AllowSharedBufferSource, options VideoFrameCopyToOptions) (ret js.Promise[js.Array[PlaneLayout]]) {
	bindings.CallVideoFrameCopyTo(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCopyTo1 returns true if the method "VideoFrame.copyTo" exists.
func (this VideoFrame) HasFuncCopyTo1() bool {
	return js.True == bindings.HasFuncVideoFrameCopyTo1(
		this.ref,
	)
}

// FuncCopyTo1 returns the method "VideoFrame.copyTo".
func (this VideoFrame) FuncCopyTo1() (fn js.Func[func(destination AllowSharedBufferSource) js.Promise[js.Array[PlaneLayout]]]) {
	bindings.FuncVideoFrameCopyTo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTo1 calls the method "VideoFrame.copyTo".
func (this VideoFrame) CopyTo1(destination AllowSharedBufferSource) (ret js.Promise[js.Array[PlaneLayout]]) {
	bindings.CallVideoFrameCopyTo1(
		this.ref, js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryCopyTo1 calls the method "VideoFrame.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryCopyTo1(destination AllowSharedBufferSource) (ret js.Promise[js.Array[PlaneLayout]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameCopyTo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
}

// HasFuncClone returns true if the method "VideoFrame.clone" exists.
func (this VideoFrame) HasFuncClone() bool {
	return js.True == bindings.HasFuncVideoFrameClone(
		this.ref,
	)
}

// FuncClone returns the method "VideoFrame.clone".
func (this VideoFrame) FuncClone() (fn js.Func[func() VideoFrame]) {
	bindings.FuncVideoFrameClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "VideoFrame.clone".
func (this VideoFrame) Clone() (ret VideoFrame) {
	bindings.CallVideoFrameClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "VideoFrame.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryClone() (ret VideoFrame, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "VideoFrame.close" exists.
func (this VideoFrame) HasFuncClose() bool {
	return js.True == bindings.HasFuncVideoFrameClose(
		this.ref,
	)
}

// FuncClose returns the method "VideoFrame.close".
func (this VideoFrame) FuncClose() (fn js.Func[func()]) {
	bindings.FuncVideoFrameClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "VideoFrame.close".
func (this VideoFrame) Close() (ret js.Void) {
	bindings.CallVideoFrameClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "VideoFrame.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoFrame) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoFrameClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Canvas returns the value of property "WebGLRenderingContext.canvas".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextCanvas(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferWidth returns the value of property "WebGLRenderingContext.drawingBufferWidth".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferWidth() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferHeight returns the value of property "WebGLRenderingContext.drawingBufferHeight".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferHeight() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferColorSpace returns the value of property "WebGLRenderingContext.drawingBufferColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) DrawingBufferColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextDrawingBufferColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDrawingBufferColorSpace sets the value of property "WebGLRenderingContext.drawingBufferColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGLRenderingContext) SetDrawingBufferColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGLRenderingContextDrawingBufferColorSpace(
		this.ref,
		uint32(val),
	)
}

// UnpackColorSpace returns the value of property "WebGLRenderingContext.unpackColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGLRenderingContext) UnpackColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGLRenderingContextUnpackColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUnpackColorSpace sets the value of property "WebGLRenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGLRenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGLRenderingContextUnpackColorSpace(
		this.ref,
		uint32(val),
	)
}

// HasFuncGetContextAttributes returns true if the method "WebGLRenderingContext.getContextAttributes" exists.
func (this WebGLRenderingContext) HasFuncGetContextAttributes() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetContextAttributes(
		this.ref,
	)
}

// FuncGetContextAttributes returns the method "WebGLRenderingContext.getContextAttributes".
func (this WebGLRenderingContext) FuncGetContextAttributes() (fn js.Func[func() WebGLContextAttributes]) {
	bindings.FuncWebGLRenderingContextGetContextAttributes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContextAttributes calls the method "WebGLRenderingContext.getContextAttributes".
func (this WebGLRenderingContext) GetContextAttributes() (ret WebGLContextAttributes) {
	bindings.CallWebGLRenderingContextGetContextAttributes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "WebGLRenderingContext.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetContextAttributes() (ret WebGLContextAttributes, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetContextAttributes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsContextLost returns true if the method "WebGLRenderingContext.isContextLost" exists.
func (this WebGLRenderingContext) HasFuncIsContextLost() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsContextLost(
		this.ref,
	)
}

// FuncIsContextLost returns the method "WebGLRenderingContext.isContextLost".
func (this WebGLRenderingContext) FuncIsContextLost() (fn js.Func[func() bool]) {
	bindings.FuncWebGLRenderingContextIsContextLost(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsContextLost calls the method "WebGLRenderingContext.isContextLost".
func (this WebGLRenderingContext) IsContextLost() (ret bool) {
	bindings.CallWebGLRenderingContextIsContextLost(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "WebGLRenderingContext.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsContextLost(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSupportedExtensions returns true if the method "WebGLRenderingContext.getSupportedExtensions" exists.
func (this WebGLRenderingContext) HasFuncGetSupportedExtensions() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetSupportedExtensions(
		this.ref,
	)
}

// FuncGetSupportedExtensions returns the method "WebGLRenderingContext.getSupportedExtensions".
func (this WebGLRenderingContext) FuncGetSupportedExtensions() (fn js.Func[func() js.Array[js.String]]) {
	bindings.FuncWebGLRenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSupportedExtensions calls the method "WebGLRenderingContext.getSupportedExtensions".
func (this WebGLRenderingContext) GetSupportedExtensions() (ret js.Array[js.String]) {
	bindings.CallWebGLRenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSupportedExtensions calls the method "WebGLRenderingContext.getSupportedExtensions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetSupportedExtensions() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetExtension returns true if the method "WebGLRenderingContext.getExtension" exists.
func (this WebGLRenderingContext) HasFuncGetExtension() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetExtension(
		this.ref,
	)
}

// FuncGetExtension returns the method "WebGLRenderingContext.getExtension".
func (this WebGLRenderingContext) FuncGetExtension() (fn js.Func[func(name js.String) js.Object]) {
	bindings.FuncWebGLRenderingContextGetExtension(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetExtension calls the method "WebGLRenderingContext.getExtension".
func (this WebGLRenderingContext) GetExtension(name js.String) (ret js.Object) {
	bindings.CallWebGLRenderingContextGetExtension(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetExtension calls the method "WebGLRenderingContext.getExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetExtension(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetExtension(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncActiveTexture returns true if the method "WebGLRenderingContext.activeTexture" exists.
func (this WebGLRenderingContext) HasFuncActiveTexture() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextActiveTexture(
		this.ref,
	)
}

// FuncActiveTexture returns the method "WebGLRenderingContext.activeTexture".
func (this WebGLRenderingContext) FuncActiveTexture() (fn js.Func[func(texture GLenum)]) {
	bindings.FuncWebGLRenderingContextActiveTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ActiveTexture calls the method "WebGLRenderingContext.activeTexture".
func (this WebGLRenderingContext) ActiveTexture(texture GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextActiveTexture(
		this.ref, js.Pointer(&ret),
		uint32(texture),
	)

	return
}

// TryActiveTexture calls the method "WebGLRenderingContext.activeTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryActiveTexture(texture GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextActiveTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(texture),
	)

	return
}

// HasFuncAttachShader returns true if the method "WebGLRenderingContext.attachShader" exists.
func (this WebGLRenderingContext) HasFuncAttachShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextAttachShader(
		this.ref,
	)
}

// FuncAttachShader returns the method "WebGLRenderingContext.attachShader".
func (this WebGLRenderingContext) FuncAttachShader() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	bindings.FuncWebGLRenderingContextAttachShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AttachShader calls the method "WebGLRenderingContext.attachShader".
func (this WebGLRenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextAttachShader(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasFuncBindAttribLocation returns true if the method "WebGLRenderingContext.bindAttribLocation" exists.
func (this WebGLRenderingContext) HasFuncBindAttribLocation() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBindAttribLocation(
		this.ref,
	)
}

// FuncBindAttribLocation returns the method "WebGLRenderingContext.bindAttribLocation".
func (this WebGLRenderingContext) FuncBindAttribLocation() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	bindings.FuncWebGLRenderingContextBindAttribLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindAttribLocation calls the method "WebGLRenderingContext.bindAttribLocation".
func (this WebGLRenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindAttribLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// HasFuncBindBuffer returns true if the method "WebGLRenderingContext.bindBuffer" exists.
func (this WebGLRenderingContext) HasFuncBindBuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBindBuffer(
		this.ref,
	)
}

// FuncBindBuffer returns the method "WebGLRenderingContext.bindBuffer".
func (this WebGLRenderingContext) FuncBindBuffer() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	bindings.FuncWebGLRenderingContextBindBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindBuffer calls the method "WebGLRenderingContext.bindBuffer".
func (this WebGLRenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// HasFuncBindFramebuffer returns true if the method "WebGLRenderingContext.bindFramebuffer" exists.
func (this WebGLRenderingContext) HasFuncBindFramebuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBindFramebuffer(
		this.ref,
	)
}

// FuncBindFramebuffer returns the method "WebGLRenderingContext.bindFramebuffer".
func (this WebGLRenderingContext) FuncBindFramebuffer() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	bindings.FuncWebGLRenderingContextBindFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindFramebuffer calls the method "WebGLRenderingContext.bindFramebuffer".
func (this WebGLRenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindFramebuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// HasFuncBindRenderbuffer returns true if the method "WebGLRenderingContext.bindRenderbuffer" exists.
func (this WebGLRenderingContext) HasFuncBindRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBindRenderbuffer(
		this.ref,
	)
}

// FuncBindRenderbuffer returns the method "WebGLRenderingContext.bindRenderbuffer".
func (this WebGLRenderingContext) FuncBindRenderbuffer() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGLRenderingContextBindRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindRenderbuffer calls the method "WebGLRenderingContext.bindRenderbuffer".
func (this WebGLRenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindRenderbuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncBindTexture returns true if the method "WebGLRenderingContext.bindTexture" exists.
func (this WebGLRenderingContext) HasFuncBindTexture() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBindTexture(
		this.ref,
	)
}

// FuncBindTexture returns the method "WebGLRenderingContext.bindTexture".
func (this WebGLRenderingContext) FuncBindTexture() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	bindings.FuncWebGLRenderingContextBindTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindTexture calls the method "WebGLRenderingContext.bindTexture".
func (this WebGLRenderingContext) BindTexture(target GLenum, texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGLRenderingContextBindTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		texture.Ref(),
	)

	return
}

// HasFuncBlendColor returns true if the method "WebGLRenderingContext.blendColor" exists.
func (this WebGLRenderingContext) HasFuncBlendColor() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBlendColor(
		this.ref,
	)
}

// FuncBlendColor returns the method "WebGLRenderingContext.blendColor".
func (this WebGLRenderingContext) FuncBlendColor() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	bindings.FuncWebGLRenderingContextBlendColor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendColor calls the method "WebGLRenderingContext.blendColor".
func (this WebGLRenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendColor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasFuncBlendEquation returns true if the method "WebGLRenderingContext.blendEquation" exists.
func (this WebGLRenderingContext) HasFuncBlendEquation() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBlendEquation(
		this.ref,
	)
}

// FuncBlendEquation returns the method "WebGLRenderingContext.blendEquation".
func (this WebGLRenderingContext) FuncBlendEquation() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGLRenderingContextBlendEquation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendEquation calls the method "WebGLRenderingContext.blendEquation".
func (this WebGLRenderingContext) BlendEquation(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendEquation(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryBlendEquation calls the method "WebGLRenderingContext.blendEquation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryBlendEquation(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextBlendEquation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncBlendEquationSeparate returns true if the method "WebGLRenderingContext.blendEquationSeparate" exists.
func (this WebGLRenderingContext) HasFuncBlendEquationSeparate() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBlendEquationSeparate(
		this.ref,
	)
}

// FuncBlendEquationSeparate returns the method "WebGLRenderingContext.blendEquationSeparate".
func (this WebGLRenderingContext) FuncBlendEquationSeparate() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	bindings.FuncWebGLRenderingContextBlendEquationSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendEquationSeparate calls the method "WebGLRenderingContext.blendEquationSeparate".
func (this WebGLRenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendEquationSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// HasFuncBlendFunc returns true if the method "WebGLRenderingContext.blendFunc" exists.
func (this WebGLRenderingContext) HasFuncBlendFunc() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBlendFunc(
		this.ref,
	)
}

// FuncBlendFunc returns the method "WebGLRenderingContext.blendFunc".
func (this WebGLRenderingContext) FuncBlendFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	bindings.FuncWebGLRenderingContextBlendFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendFunc calls the method "WebGLRenderingContext.blendFunc".
func (this WebGLRenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendFunc(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// HasFuncBlendFuncSeparate returns true if the method "WebGLRenderingContext.blendFuncSeparate" exists.
func (this WebGLRenderingContext) HasFuncBlendFuncSeparate() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBlendFuncSeparate(
		this.ref,
	)
}

// FuncBlendFuncSeparate returns the method "WebGLRenderingContext.blendFuncSeparate".
func (this WebGLRenderingContext) FuncBlendFuncSeparate() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	bindings.FuncWebGLRenderingContextBlendFuncSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendFuncSeparate calls the method "WebGLRenderingContext.blendFuncSeparate".
func (this WebGLRenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBlendFuncSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// HasFuncCheckFramebufferStatus returns true if the method "WebGLRenderingContext.checkFramebufferStatus" exists.
func (this WebGLRenderingContext) HasFuncCheckFramebufferStatus() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCheckFramebufferStatus(
		this.ref,
	)
}

// FuncCheckFramebufferStatus returns the method "WebGLRenderingContext.checkFramebufferStatus".
func (this WebGLRenderingContext) FuncCheckFramebufferStatus() (fn js.Func[func(target GLenum) GLenum]) {
	bindings.FuncWebGLRenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckFramebufferStatus calls the method "WebGLRenderingContext.checkFramebufferStatus".
func (this WebGLRenderingContext) CheckFramebufferStatus(target GLenum) (ret GLenum) {
	bindings.CallWebGLRenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryCheckFramebufferStatus calls the method "WebGLRenderingContext.checkFramebufferStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCheckFramebufferStatus(target GLenum) (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncClear returns true if the method "WebGLRenderingContext.clear" exists.
func (this WebGLRenderingContext) HasFuncClear() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextClear(
		this.ref,
	)
}

// FuncClear returns the method "WebGLRenderingContext.clear".
func (this WebGLRenderingContext) FuncClear() (fn js.Func[func(mask GLbitfield)]) {
	bindings.FuncWebGLRenderingContextClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "WebGLRenderingContext.clear".
func (this WebGLRenderingContext) Clear(mask GLbitfield) (ret js.Void) {
	bindings.CallWebGLRenderingContextClear(
		this.ref, js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryClear calls the method "WebGLRenderingContext.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClear(mask GLbitfield) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasFuncClearColor returns true if the method "WebGLRenderingContext.clearColor" exists.
func (this WebGLRenderingContext) HasFuncClearColor() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextClearColor(
		this.ref,
	)
}

// FuncClearColor returns the method "WebGLRenderingContext.clearColor".
func (this WebGLRenderingContext) FuncClearColor() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	bindings.FuncWebGLRenderingContextClearColor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearColor calls the method "WebGLRenderingContext.clearColor".
func (this WebGLRenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearColor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasFuncClearDepth returns true if the method "WebGLRenderingContext.clearDepth" exists.
func (this WebGLRenderingContext) HasFuncClearDepth() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextClearDepth(
		this.ref,
	)
}

// FuncClearDepth returns the method "WebGLRenderingContext.clearDepth".
func (this WebGLRenderingContext) FuncClearDepth() (fn js.Func[func(depth GLclampf)]) {
	bindings.FuncWebGLRenderingContextClearDepth(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearDepth calls the method "WebGLRenderingContext.clearDepth".
func (this WebGLRenderingContext) ClearDepth(depth GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearDepth(
		this.ref, js.Pointer(&ret),
		float32(depth),
	)

	return
}

// TryClearDepth calls the method "WebGLRenderingContext.clearDepth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClearDepth(depth GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClearDepth(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(depth),
	)

	return
}

// HasFuncClearStencil returns true if the method "WebGLRenderingContext.clearStencil" exists.
func (this WebGLRenderingContext) HasFuncClearStencil() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextClearStencil(
		this.ref,
	)
}

// FuncClearStencil returns the method "WebGLRenderingContext.clearStencil".
func (this WebGLRenderingContext) FuncClearStencil() (fn js.Func[func(s GLint)]) {
	bindings.FuncWebGLRenderingContextClearStencil(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearStencil calls the method "WebGLRenderingContext.clearStencil".
func (this WebGLRenderingContext) ClearStencil(s GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextClearStencil(
		this.ref, js.Pointer(&ret),
		int32(s),
	)

	return
}

// TryClearStencil calls the method "WebGLRenderingContext.clearStencil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryClearStencil(s GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextClearStencil(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(s),
	)

	return
}

// HasFuncColorMask returns true if the method "WebGLRenderingContext.colorMask" exists.
func (this WebGLRenderingContext) HasFuncColorMask() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextColorMask(
		this.ref,
	)
}

// FuncColorMask returns the method "WebGLRenderingContext.colorMask".
func (this WebGLRenderingContext) FuncColorMask() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	bindings.FuncWebGLRenderingContextColorMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ColorMask calls the method "WebGLRenderingContext.colorMask".
func (this WebGLRenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextColorMask(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// HasFuncCompileShader returns true if the method "WebGLRenderingContext.compileShader" exists.
func (this WebGLRenderingContext) HasFuncCompileShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCompileShader(
		this.ref,
	)
}

// FuncCompileShader returns the method "WebGLRenderingContext.compileShader".
func (this WebGLRenderingContext) FuncCompileShader() (fn js.Func[func(shader WebGLShader)]) {
	bindings.FuncWebGLRenderingContextCompileShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompileShader calls the method "WebGLRenderingContext.compileShader".
func (this WebGLRenderingContext) CompileShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompileShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryCompileShader calls the method "WebGLRenderingContext.compileShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCompileShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCompileShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncCopyTexImage2D returns true if the method "WebGLRenderingContext.copyTexImage2D" exists.
func (this WebGLRenderingContext) HasFuncCopyTexImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCopyTexImage2D(
		this.ref,
	)
}

// FuncCopyTexImage2D returns the method "WebGLRenderingContext.copyTexImage2D".
func (this WebGLRenderingContext) FuncCopyTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	bindings.FuncWebGLRenderingContextCopyTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTexImage2D calls the method "WebGLRenderingContext.copyTexImage2D".
func (this WebGLRenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextCopyTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCopyTexSubImage2D returns true if the method "WebGLRenderingContext.copyTexSubImage2D" exists.
func (this WebGLRenderingContext) HasFuncCopyTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCopyTexSubImage2D(
		this.ref,
	)
}

// FuncCopyTexSubImage2D returns the method "WebGLRenderingContext.copyTexSubImage2D".
func (this WebGLRenderingContext) FuncCopyTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGLRenderingContextCopyTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTexSubImage2D calls the method "WebGLRenderingContext.copyTexSubImage2D".
func (this WebGLRenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextCopyTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCreateBuffer returns true if the method "WebGLRenderingContext.createBuffer" exists.
func (this WebGLRenderingContext) HasFuncCreateBuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateBuffer(
		this.ref,
	)
}

// FuncCreateBuffer returns the method "WebGLRenderingContext.createBuffer".
func (this WebGLRenderingContext) FuncCreateBuffer() (fn js.Func[func() WebGLBuffer]) {
	bindings.FuncWebGLRenderingContextCreateBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBuffer calls the method "WebGLRenderingContext.createBuffer".
func (this WebGLRenderingContext) CreateBuffer() (ret WebGLBuffer) {
	bindings.CallWebGLRenderingContextCreateBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateBuffer calls the method "WebGLRenderingContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateBuffer() (ret WebGLBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateFramebuffer returns true if the method "WebGLRenderingContext.createFramebuffer" exists.
func (this WebGLRenderingContext) HasFuncCreateFramebuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateFramebuffer(
		this.ref,
	)
}

// FuncCreateFramebuffer returns the method "WebGLRenderingContext.createFramebuffer".
func (this WebGLRenderingContext) FuncCreateFramebuffer() (fn js.Func[func() WebGLFramebuffer]) {
	bindings.FuncWebGLRenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateFramebuffer calls the method "WebGLRenderingContext.createFramebuffer".
func (this WebGLRenderingContext) CreateFramebuffer() (ret WebGLFramebuffer) {
	bindings.CallWebGLRenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateFramebuffer calls the method "WebGLRenderingContext.createFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateFramebuffer() (ret WebGLFramebuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateProgram returns true if the method "WebGLRenderingContext.createProgram" exists.
func (this WebGLRenderingContext) HasFuncCreateProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateProgram(
		this.ref,
	)
}

// FuncCreateProgram returns the method "WebGLRenderingContext.createProgram".
func (this WebGLRenderingContext) FuncCreateProgram() (fn js.Func[func() WebGLProgram]) {
	bindings.FuncWebGLRenderingContextCreateProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateProgram calls the method "WebGLRenderingContext.createProgram".
func (this WebGLRenderingContext) CreateProgram() (ret WebGLProgram) {
	bindings.CallWebGLRenderingContextCreateProgram(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateProgram calls the method "WebGLRenderingContext.createProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateProgram() (ret WebGLProgram, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateRenderbuffer returns true if the method "WebGLRenderingContext.createRenderbuffer" exists.
func (this WebGLRenderingContext) HasFuncCreateRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateRenderbuffer(
		this.ref,
	)
}

// FuncCreateRenderbuffer returns the method "WebGLRenderingContext.createRenderbuffer".
func (this WebGLRenderingContext) FuncCreateRenderbuffer() (fn js.Func[func() WebGLRenderbuffer]) {
	bindings.FuncWebGLRenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRenderbuffer calls the method "WebGLRenderingContext.createRenderbuffer".
func (this WebGLRenderingContext) CreateRenderbuffer() (ret WebGLRenderbuffer) {
	bindings.CallWebGLRenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateRenderbuffer calls the method "WebGLRenderingContext.createRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateRenderbuffer() (ret WebGLRenderbuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateShader returns true if the method "WebGLRenderingContext.createShader" exists.
func (this WebGLRenderingContext) HasFuncCreateShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateShader(
		this.ref,
	)
}

// FuncCreateShader returns the method "WebGLRenderingContext.createShader".
func (this WebGLRenderingContext) FuncCreateShader() (fn js.Func[func(typ GLenum) WebGLShader]) {
	bindings.FuncWebGLRenderingContextCreateShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateShader calls the method "WebGLRenderingContext.createShader".
func (this WebGLRenderingContext) CreateShader(typ GLenum) (ret WebGLShader) {
	bindings.CallWebGLRenderingContextCreateShader(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCreateShader calls the method "WebGLRenderingContext.createShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateShader(typ GLenum) (ret WebGLShader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncCreateTexture returns true if the method "WebGLRenderingContext.createTexture" exists.
func (this WebGLRenderingContext) HasFuncCreateTexture() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCreateTexture(
		this.ref,
	)
}

// FuncCreateTexture returns the method "WebGLRenderingContext.createTexture".
func (this WebGLRenderingContext) FuncCreateTexture() (fn js.Func[func() WebGLTexture]) {
	bindings.FuncWebGLRenderingContextCreateTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTexture calls the method "WebGLRenderingContext.createTexture".
func (this WebGLRenderingContext) CreateTexture() (ret WebGLTexture) {
	bindings.CallWebGLRenderingContextCreateTexture(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTexture calls the method "WebGLRenderingContext.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCreateTexture() (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCreateTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCullFace returns true if the method "WebGLRenderingContext.cullFace" exists.
func (this WebGLRenderingContext) HasFuncCullFace() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCullFace(
		this.ref,
	)
}

// FuncCullFace returns the method "WebGLRenderingContext.cullFace".
func (this WebGLRenderingContext) FuncCullFace() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGLRenderingContextCullFace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CullFace calls the method "WebGLRenderingContext.cullFace".
func (this WebGLRenderingContext) CullFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextCullFace(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryCullFace calls the method "WebGLRenderingContext.cullFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryCullFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextCullFace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncDeleteBuffer returns true if the method "WebGLRenderingContext.deleteBuffer" exists.
func (this WebGLRenderingContext) HasFuncDeleteBuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteBuffer(
		this.ref,
	)
}

// FuncDeleteBuffer returns the method "WebGLRenderingContext.deleteBuffer".
func (this WebGLRenderingContext) FuncDeleteBuffer() (fn js.Func[func(buffer WebGLBuffer)]) {
	bindings.FuncWebGLRenderingContextDeleteBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteBuffer calls the method "WebGLRenderingContext.deleteBuffer".
func (this WebGLRenderingContext) DeleteBuffer(buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteBuffer(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryDeleteBuffer calls the method "WebGLRenderingContext.deleteBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteBuffer(buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncDeleteFramebuffer returns true if the method "WebGLRenderingContext.deleteFramebuffer" exists.
func (this WebGLRenderingContext) HasFuncDeleteFramebuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteFramebuffer(
		this.ref,
	)
}

// FuncDeleteFramebuffer returns the method "WebGLRenderingContext.deleteFramebuffer".
func (this WebGLRenderingContext) FuncDeleteFramebuffer() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	bindings.FuncWebGLRenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteFramebuffer calls the method "WebGLRenderingContext.deleteFramebuffer".
func (this WebGLRenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryDeleteFramebuffer calls the method "WebGLRenderingContext.deleteFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasFuncDeleteProgram returns true if the method "WebGLRenderingContext.deleteProgram" exists.
func (this WebGLRenderingContext) HasFuncDeleteProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteProgram(
		this.ref,
	)
}

// FuncDeleteProgram returns the method "WebGLRenderingContext.deleteProgram".
func (this WebGLRenderingContext) FuncDeleteProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGLRenderingContextDeleteProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteProgram calls the method "WebGLRenderingContext.deleteProgram".
func (this WebGLRenderingContext) DeleteProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryDeleteProgram calls the method "WebGLRenderingContext.deleteProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncDeleteRenderbuffer returns true if the method "WebGLRenderingContext.deleteRenderbuffer" exists.
func (this WebGLRenderingContext) HasFuncDeleteRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteRenderbuffer(
		this.ref,
	)
}

// FuncDeleteRenderbuffer returns the method "WebGLRenderingContext.deleteRenderbuffer".
func (this WebGLRenderingContext) FuncDeleteRenderbuffer() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGLRenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRenderbuffer calls the method "WebGLRenderingContext.deleteRenderbuffer".
func (this WebGLRenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryDeleteRenderbuffer calls the method "WebGLRenderingContext.deleteRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncDeleteShader returns true if the method "WebGLRenderingContext.deleteShader" exists.
func (this WebGLRenderingContext) HasFuncDeleteShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteShader(
		this.ref,
	)
}

// FuncDeleteShader returns the method "WebGLRenderingContext.deleteShader".
func (this WebGLRenderingContext) FuncDeleteShader() (fn js.Func[func(shader WebGLShader)]) {
	bindings.FuncWebGLRenderingContextDeleteShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteShader calls the method "WebGLRenderingContext.deleteShader".
func (this WebGLRenderingContext) DeleteShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryDeleteShader calls the method "WebGLRenderingContext.deleteShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncDeleteTexture returns true if the method "WebGLRenderingContext.deleteTexture" exists.
func (this WebGLRenderingContext) HasFuncDeleteTexture() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDeleteTexture(
		this.ref,
	)
}

// FuncDeleteTexture returns the method "WebGLRenderingContext.deleteTexture".
func (this WebGLRenderingContext) FuncDeleteTexture() (fn js.Func[func(texture WebGLTexture)]) {
	bindings.FuncWebGLRenderingContextDeleteTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteTexture calls the method "WebGLRenderingContext.deleteTexture".
func (this WebGLRenderingContext) DeleteTexture(texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGLRenderingContextDeleteTexture(
		this.ref, js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryDeleteTexture calls the method "WebGLRenderingContext.deleteTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDeleteTexture(texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDeleteTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasFuncDepthFunc returns true if the method "WebGLRenderingContext.depthFunc" exists.
func (this WebGLRenderingContext) HasFuncDepthFunc() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDepthFunc(
		this.ref,
	)
}

// FuncDepthFunc returns the method "WebGLRenderingContext.depthFunc".
func (this WebGLRenderingContext) FuncDepthFunc() (fn js.Func[func(fn GLenum)]) {
	bindings.FuncWebGLRenderingContextDepthFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthFunc calls the method "WebGLRenderingContext.depthFunc".
func (this WebGLRenderingContext) DepthFunc(fn GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthFunc(
		this.ref, js.Pointer(&ret),
		uint32(fn),
	)

	return
}

// TryDepthFunc calls the method "WebGLRenderingContext.depthFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDepthFunc(fn GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDepthFunc(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
	)

	return
}

// HasFuncDepthMask returns true if the method "WebGLRenderingContext.depthMask" exists.
func (this WebGLRenderingContext) HasFuncDepthMask() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDepthMask(
		this.ref,
	)
}

// FuncDepthMask returns the method "WebGLRenderingContext.depthMask".
func (this WebGLRenderingContext) FuncDepthMask() (fn js.Func[func(flag GLboolean)]) {
	bindings.FuncWebGLRenderingContextDepthMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthMask calls the method "WebGLRenderingContext.depthMask".
func (this WebGLRenderingContext) DepthMask(flag GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthMask(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(flag)),
	)

	return
}

// TryDepthMask calls the method "WebGLRenderingContext.depthMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDepthMask(flag GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDepthMask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(flag)),
	)

	return
}

// HasFuncDepthRange returns true if the method "WebGLRenderingContext.depthRange" exists.
func (this WebGLRenderingContext) HasFuncDepthRange() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDepthRange(
		this.ref,
	)
}

// FuncDepthRange returns the method "WebGLRenderingContext.depthRange".
func (this WebGLRenderingContext) FuncDepthRange() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	bindings.FuncWebGLRenderingContextDepthRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthRange calls the method "WebGLRenderingContext.depthRange".
func (this WebGLRenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void) {
	bindings.CallWebGLRenderingContextDepthRange(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(zNear),
		float32(zFar),
	)

	return
}

// HasFuncDetachShader returns true if the method "WebGLRenderingContext.detachShader" exists.
func (this WebGLRenderingContext) HasFuncDetachShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDetachShader(
		this.ref,
	)
}

// FuncDetachShader returns the method "WebGLRenderingContext.detachShader".
func (this WebGLRenderingContext) FuncDetachShader() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	bindings.FuncWebGLRenderingContextDetachShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DetachShader calls the method "WebGLRenderingContext.detachShader".
func (this WebGLRenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGLRenderingContextDetachShader(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasFuncDisable returns true if the method "WebGLRenderingContext.disable" exists.
func (this WebGLRenderingContext) HasFuncDisable() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDisable(
		this.ref,
	)
}

// FuncDisable returns the method "WebGLRenderingContext.disable".
func (this WebGLRenderingContext) FuncDisable() (fn js.Func[func(cap GLenum)]) {
	bindings.FuncWebGLRenderingContextDisable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disable calls the method "WebGLRenderingContext.disable".
func (this WebGLRenderingContext) Disable(cap GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextDisable(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryDisable calls the method "WebGLRenderingContext.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDisable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDisable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncDisableVertexAttribArray returns true if the method "WebGLRenderingContext.disableVertexAttribArray" exists.
func (this WebGLRenderingContext) HasFuncDisableVertexAttribArray() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDisableVertexAttribArray(
		this.ref,
	)
}

// FuncDisableVertexAttribArray returns the method "WebGLRenderingContext.disableVertexAttribArray".
func (this WebGLRenderingContext) FuncDisableVertexAttribArray() (fn js.Func[func(index GLuint)]) {
	bindings.FuncWebGLRenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DisableVertexAttribArray calls the method "WebGLRenderingContext.disableVertexAttribArray".
func (this WebGLRenderingContext) DisableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDisableVertexAttribArray calls the method "WebGLRenderingContext.disableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryDisableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncDrawArrays returns true if the method "WebGLRenderingContext.drawArrays" exists.
func (this WebGLRenderingContext) HasFuncDrawArrays() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDrawArrays(
		this.ref,
	)
}

// FuncDrawArrays returns the method "WebGLRenderingContext.drawArrays".
func (this WebGLRenderingContext) FuncDrawArrays() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	bindings.FuncWebGLRenderingContextDrawArrays(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawArrays calls the method "WebGLRenderingContext.drawArrays".
func (this WebGLRenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextDrawArrays(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// HasFuncDrawElements returns true if the method "WebGLRenderingContext.drawElements" exists.
func (this WebGLRenderingContext) HasFuncDrawElements() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextDrawElements(
		this.ref,
	)
}

// FuncDrawElements returns the method "WebGLRenderingContext.drawElements".
func (this WebGLRenderingContext) FuncDrawElements() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	bindings.FuncWebGLRenderingContextDrawElements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawElements calls the method "WebGLRenderingContext.drawElements".
func (this WebGLRenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGLRenderingContextDrawElements(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasFuncEnable returns true if the method "WebGLRenderingContext.enable" exists.
func (this WebGLRenderingContext) HasFuncEnable() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextEnable(
		this.ref,
	)
}

// FuncEnable returns the method "WebGLRenderingContext.enable".
func (this WebGLRenderingContext) FuncEnable() (fn js.Func[func(cap GLenum)]) {
	bindings.FuncWebGLRenderingContextEnable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Enable calls the method "WebGLRenderingContext.enable".
func (this WebGLRenderingContext) Enable(cap GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextEnable(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryEnable calls the method "WebGLRenderingContext.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryEnable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextEnable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncEnableVertexAttribArray returns true if the method "WebGLRenderingContext.enableVertexAttribArray" exists.
func (this WebGLRenderingContext) HasFuncEnableVertexAttribArray() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextEnableVertexAttribArray(
		this.ref,
	)
}

// FuncEnableVertexAttribArray returns the method "WebGLRenderingContext.enableVertexAttribArray".
func (this WebGLRenderingContext) FuncEnableVertexAttribArray() (fn js.Func[func(index GLuint)]) {
	bindings.FuncWebGLRenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EnableVertexAttribArray calls the method "WebGLRenderingContext.enableVertexAttribArray".
func (this WebGLRenderingContext) EnableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryEnableVertexAttribArray calls the method "WebGLRenderingContext.enableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryEnableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncFinish returns true if the method "WebGLRenderingContext.finish" exists.
func (this WebGLRenderingContext) HasFuncFinish() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextFinish(
		this.ref,
	)
}

// FuncFinish returns the method "WebGLRenderingContext.finish".
func (this WebGLRenderingContext) FuncFinish() (fn js.Func[func()]) {
	bindings.FuncWebGLRenderingContextFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "WebGLRenderingContext.finish".
func (this WebGLRenderingContext) Finish() (ret js.Void) {
	bindings.CallWebGLRenderingContextFinish(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "WebGLRenderingContext.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFlush returns true if the method "WebGLRenderingContext.flush" exists.
func (this WebGLRenderingContext) HasFuncFlush() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextFlush(
		this.ref,
	)
}

// FuncFlush returns the method "WebGLRenderingContext.flush".
func (this WebGLRenderingContext) FuncFlush() (fn js.Func[func()]) {
	bindings.FuncWebGLRenderingContextFlush(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Flush calls the method "WebGLRenderingContext.flush".
func (this WebGLRenderingContext) Flush() (ret js.Void) {
	bindings.CallWebGLRenderingContextFlush(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "WebGLRenderingContext.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFlush(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFramebufferRenderbuffer returns true if the method "WebGLRenderingContext.framebufferRenderbuffer" exists.
func (this WebGLRenderingContext) HasFuncFramebufferRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextFramebufferRenderbuffer(
		this.ref,
	)
}

// FuncFramebufferRenderbuffer returns the method "WebGLRenderingContext.framebufferRenderbuffer".
func (this WebGLRenderingContext) FuncFramebufferRenderbuffer() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGLRenderingContextFramebufferRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FramebufferRenderbuffer calls the method "WebGLRenderingContext.framebufferRenderbuffer".
func (this WebGLRenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGLRenderingContextFramebufferRenderbuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncFramebufferTexture2D returns true if the method "WebGLRenderingContext.framebufferTexture2D" exists.
func (this WebGLRenderingContext) HasFuncFramebufferTexture2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextFramebufferTexture2D(
		this.ref,
	)
}

// FuncFramebufferTexture2D returns the method "WebGLRenderingContext.framebufferTexture2D".
func (this WebGLRenderingContext) FuncFramebufferTexture2D() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	bindings.FuncWebGLRenderingContextFramebufferTexture2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FramebufferTexture2D calls the method "WebGLRenderingContext.framebufferTexture2D".
func (this WebGLRenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextFramebufferTexture2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// HasFuncFrontFace returns true if the method "WebGLRenderingContext.frontFace" exists.
func (this WebGLRenderingContext) HasFuncFrontFace() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextFrontFace(
		this.ref,
	)
}

// FuncFrontFace returns the method "WebGLRenderingContext.frontFace".
func (this WebGLRenderingContext) FuncFrontFace() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGLRenderingContextFrontFace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FrontFace calls the method "WebGLRenderingContext.frontFace".
func (this WebGLRenderingContext) FrontFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextFrontFace(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryFrontFace calls the method "WebGLRenderingContext.frontFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryFrontFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextFrontFace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncGenerateMipmap returns true if the method "WebGLRenderingContext.generateMipmap" exists.
func (this WebGLRenderingContext) HasFuncGenerateMipmap() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGenerateMipmap(
		this.ref,
	)
}

// FuncGenerateMipmap returns the method "WebGLRenderingContext.generateMipmap".
func (this WebGLRenderingContext) FuncGenerateMipmap() (fn js.Func[func(target GLenum)]) {
	bindings.FuncWebGLRenderingContextGenerateMipmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GenerateMipmap calls the method "WebGLRenderingContext.generateMipmap".
func (this WebGLRenderingContext) GenerateMipmap(target GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextGenerateMipmap(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryGenerateMipmap calls the method "WebGLRenderingContext.generateMipmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGenerateMipmap(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGenerateMipmap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncGetActiveAttrib returns true if the method "WebGLRenderingContext.getActiveAttrib" exists.
func (this WebGLRenderingContext) HasFuncGetActiveAttrib() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetActiveAttrib(
		this.ref,
	)
}

// FuncGetActiveAttrib returns the method "WebGLRenderingContext.getActiveAttrib".
func (this WebGLRenderingContext) FuncGetActiveAttrib() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	bindings.FuncWebGLRenderingContextGetActiveAttrib(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveAttrib calls the method "WebGLRenderingContext.getActiveAttrib".
func (this WebGLRenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGLRenderingContextGetActiveAttrib(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasFuncGetActiveUniform returns true if the method "WebGLRenderingContext.getActiveUniform" exists.
func (this WebGLRenderingContext) HasFuncGetActiveUniform() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetActiveUniform(
		this.ref,
	)
}

// FuncGetActiveUniform returns the method "WebGLRenderingContext.getActiveUniform".
func (this WebGLRenderingContext) FuncGetActiveUniform() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	bindings.FuncWebGLRenderingContextGetActiveUniform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveUniform calls the method "WebGLRenderingContext.getActiveUniform".
func (this WebGLRenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGLRenderingContextGetActiveUniform(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasFuncGetAttachedShaders returns true if the method "WebGLRenderingContext.getAttachedShaders" exists.
func (this WebGLRenderingContext) HasFuncGetAttachedShaders() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetAttachedShaders(
		this.ref,
	)
}

// FuncGetAttachedShaders returns the method "WebGLRenderingContext.getAttachedShaders".
func (this WebGLRenderingContext) FuncGetAttachedShaders() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	bindings.FuncWebGLRenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttachedShaders calls the method "WebGLRenderingContext.getAttachedShaders".
func (this WebGLRenderingContext) GetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader]) {
	bindings.CallWebGLRenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetAttachedShaders calls the method "WebGLRenderingContext.getAttachedShaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncGetAttribLocation returns true if the method "WebGLRenderingContext.getAttribLocation" exists.
func (this WebGLRenderingContext) HasFuncGetAttribLocation() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetAttribLocation(
		this.ref,
	)
}

// FuncGetAttribLocation returns the method "WebGLRenderingContext.getAttribLocation".
func (this WebGLRenderingContext) FuncGetAttribLocation() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	bindings.FuncWebGLRenderingContextGetAttribLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttribLocation calls the method "WebGLRenderingContext.getAttribLocation".
func (this WebGLRenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGLRenderingContextGetAttribLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncGetBufferParameter returns true if the method "WebGLRenderingContext.getBufferParameter" exists.
func (this WebGLRenderingContext) HasFuncGetBufferParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetBufferParameter(
		this.ref,
	)
}

// FuncGetBufferParameter returns the method "WebGLRenderingContext.getBufferParameter".
func (this WebGLRenderingContext) FuncGetBufferParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetBufferParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBufferParameter calls the method "WebGLRenderingContext.getBufferParameter".
func (this WebGLRenderingContext) GetBufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetBufferParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetParameter returns true if the method "WebGLRenderingContext.getParameter" exists.
func (this WebGLRenderingContext) HasFuncGetParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetParameter(
		this.ref,
	)
}

// FuncGetParameter returns the method "WebGLRenderingContext.getParameter".
func (this WebGLRenderingContext) FuncGetParameter() (fn js.Func[func(pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParameter calls the method "WebGLRenderingContext.getParameter".
func (this WebGLRenderingContext) GetParameter(pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetParameter(
		this.ref, js.Pointer(&ret),
		uint32(pname),
	)

	return
}

// TryGetParameter calls the method "WebGLRenderingContext.getParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetParameter(pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetParameter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
	)

	return
}

// HasFuncGetError returns true if the method "WebGLRenderingContext.getError" exists.
func (this WebGLRenderingContext) HasFuncGetError() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetError(
		this.ref,
	)
}

// FuncGetError returns the method "WebGLRenderingContext.getError".
func (this WebGLRenderingContext) FuncGetError() (fn js.Func[func() GLenum]) {
	bindings.FuncWebGLRenderingContextGetError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetError calls the method "WebGLRenderingContext.getError".
func (this WebGLRenderingContext) GetError() (ret GLenum) {
	bindings.CallWebGLRenderingContextGetError(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetError calls the method "WebGLRenderingContext.getError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetError() (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetFramebufferAttachmentParameter returns true if the method "WebGLRenderingContext.getFramebufferAttachmentParameter" exists.
func (this WebGLRenderingContext) HasFuncGetFramebufferAttachmentParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.ref,
	)
}

// FuncGetFramebufferAttachmentParameter returns the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
func (this WebGLRenderingContext) FuncGetFramebufferAttachmentParameter() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFramebufferAttachmentParameter calls the method "WebGLRenderingContext.getFramebufferAttachmentParameter".
func (this WebGLRenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetFramebufferAttachmentParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// HasFuncGetProgramParameter returns true if the method "WebGLRenderingContext.getProgramParameter" exists.
func (this WebGLRenderingContext) HasFuncGetProgramParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetProgramParameter(
		this.ref,
	)
}

// FuncGetProgramParameter returns the method "WebGLRenderingContext.getProgramParameter".
func (this WebGLRenderingContext) FuncGetProgramParameter() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetProgramParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetProgramParameter calls the method "WebGLRenderingContext.getProgramParameter".
func (this WebGLRenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetProgramParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncGetProgramInfoLog returns true if the method "WebGLRenderingContext.getProgramInfoLog" exists.
func (this WebGLRenderingContext) HasFuncGetProgramInfoLog() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetProgramInfoLog(
		this.ref,
	)
}

// FuncGetProgramInfoLog returns the method "WebGLRenderingContext.getProgramInfoLog".
func (this WebGLRenderingContext) FuncGetProgramInfoLog() (fn js.Func[func(program WebGLProgram) js.String]) {
	bindings.FuncWebGLRenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetProgramInfoLog calls the method "WebGLRenderingContext.getProgramInfoLog".
func (this WebGLRenderingContext) GetProgramInfoLog(program WebGLProgram) (ret js.String) {
	bindings.CallWebGLRenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetProgramInfoLog calls the method "WebGLRenderingContext.getProgramInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetProgramInfoLog(program WebGLProgram) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncGetRenderbufferParameter returns true if the method "WebGLRenderingContext.getRenderbufferParameter" exists.
func (this WebGLRenderingContext) HasFuncGetRenderbufferParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetRenderbufferParameter(
		this.ref,
	)
}

// FuncGetRenderbufferParameter returns the method "WebGLRenderingContext.getRenderbufferParameter".
func (this WebGLRenderingContext) FuncGetRenderbufferParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetRenderbufferParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRenderbufferParameter calls the method "WebGLRenderingContext.getRenderbufferParameter".
func (this WebGLRenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetRenderbufferParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetShaderParameter returns true if the method "WebGLRenderingContext.getShaderParameter" exists.
func (this WebGLRenderingContext) HasFuncGetShaderParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetShaderParameter(
		this.ref,
	)
}

// FuncGetShaderParameter returns the method "WebGLRenderingContext.getShaderParameter".
func (this WebGLRenderingContext) FuncGetShaderParameter() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetShaderParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderParameter calls the method "WebGLRenderingContext.getShaderParameter".
func (this WebGLRenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetShaderParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncGetShaderPrecisionFormat returns true if the method "WebGLRenderingContext.getShaderPrecisionFormat" exists.
func (this WebGLRenderingContext) HasFuncGetShaderPrecisionFormat() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetShaderPrecisionFormat(
		this.ref,
	)
}

// FuncGetShaderPrecisionFormat returns the method "WebGLRenderingContext.getShaderPrecisionFormat".
func (this WebGLRenderingContext) FuncGetShaderPrecisionFormat() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	bindings.FuncWebGLRenderingContextGetShaderPrecisionFormat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderPrecisionFormat calls the method "WebGLRenderingContext.getShaderPrecisionFormat".
func (this WebGLRenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat) {
	bindings.CallWebGLRenderingContextGetShaderPrecisionFormat(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// HasFuncGetShaderInfoLog returns true if the method "WebGLRenderingContext.getShaderInfoLog" exists.
func (this WebGLRenderingContext) HasFuncGetShaderInfoLog() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetShaderInfoLog(
		this.ref,
	)
}

// FuncGetShaderInfoLog returns the method "WebGLRenderingContext.getShaderInfoLog".
func (this WebGLRenderingContext) FuncGetShaderInfoLog() (fn js.Func[func(shader WebGLShader) js.String]) {
	bindings.FuncWebGLRenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderInfoLog calls the method "WebGLRenderingContext.getShaderInfoLog".
func (this WebGLRenderingContext) GetShaderInfoLog(shader WebGLShader) (ret js.String) {
	bindings.CallWebGLRenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderInfoLog calls the method "WebGLRenderingContext.getShaderInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderInfoLog(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncGetShaderSource returns true if the method "WebGLRenderingContext.getShaderSource" exists.
func (this WebGLRenderingContext) HasFuncGetShaderSource() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetShaderSource(
		this.ref,
	)
}

// FuncGetShaderSource returns the method "WebGLRenderingContext.getShaderSource".
func (this WebGLRenderingContext) FuncGetShaderSource() (fn js.Func[func(shader WebGLShader) js.String]) {
	bindings.FuncWebGLRenderingContextGetShaderSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderSource calls the method "WebGLRenderingContext.getShaderSource".
func (this WebGLRenderingContext) GetShaderSource(shader WebGLShader) (ret js.String) {
	bindings.CallWebGLRenderingContextGetShaderSource(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderSource calls the method "WebGLRenderingContext.getShaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryGetShaderSource(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextGetShaderSource(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncGetTexParameter returns true if the method "WebGLRenderingContext.getTexParameter" exists.
func (this WebGLRenderingContext) HasFuncGetTexParameter() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetTexParameter(
		this.ref,
	)
}

// FuncGetTexParameter returns the method "WebGLRenderingContext.getTexParameter".
func (this WebGLRenderingContext) FuncGetTexParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetTexParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTexParameter calls the method "WebGLRenderingContext.getTexParameter".
func (this WebGLRenderingContext) GetTexParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetTexParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetUniform returns true if the method "WebGLRenderingContext.getUniform" exists.
func (this WebGLRenderingContext) HasFuncGetUniform() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetUniform(
		this.ref,
	)
}

// FuncGetUniform returns the method "WebGLRenderingContext.getUniform".
func (this WebGLRenderingContext) FuncGetUniform() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	bindings.FuncWebGLRenderingContextGetUniform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniform calls the method "WebGLRenderingContext.getUniform".
func (this WebGLRenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetUniform(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		location.Ref(),
	)

	return
}

// HasFuncGetUniformLocation returns true if the method "WebGLRenderingContext.getUniformLocation" exists.
func (this WebGLRenderingContext) HasFuncGetUniformLocation() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetUniformLocation(
		this.ref,
	)
}

// FuncGetUniformLocation returns the method "WebGLRenderingContext.getUniformLocation".
func (this WebGLRenderingContext) FuncGetUniformLocation() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	bindings.FuncWebGLRenderingContextGetUniformLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniformLocation calls the method "WebGLRenderingContext.getUniformLocation".
func (this WebGLRenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation) {
	bindings.CallWebGLRenderingContextGetUniformLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncGetVertexAttrib returns true if the method "WebGLRenderingContext.getVertexAttrib" exists.
func (this WebGLRenderingContext) HasFuncGetVertexAttrib() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetVertexAttrib(
		this.ref,
	)
}

// FuncGetVertexAttrib returns the method "WebGLRenderingContext.getVertexAttrib".
func (this WebGLRenderingContext) FuncGetVertexAttrib() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	bindings.FuncWebGLRenderingContextGetVertexAttrib(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetVertexAttrib calls the method "WebGLRenderingContext.getVertexAttrib".
func (this WebGLRenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGLRenderingContextGetVertexAttrib(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasFuncGetVertexAttribOffset returns true if the method "WebGLRenderingContext.getVertexAttribOffset" exists.
func (this WebGLRenderingContext) HasFuncGetVertexAttribOffset() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextGetVertexAttribOffset(
		this.ref,
	)
}

// FuncGetVertexAttribOffset returns the method "WebGLRenderingContext.getVertexAttribOffset".
func (this WebGLRenderingContext) FuncGetVertexAttribOffset() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	bindings.FuncWebGLRenderingContextGetVertexAttribOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetVertexAttribOffset calls the method "WebGLRenderingContext.getVertexAttribOffset".
func (this WebGLRenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr) {
	bindings.CallWebGLRenderingContextGetVertexAttribOffset(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasFuncHint returns true if the method "WebGLRenderingContext.hint" exists.
func (this WebGLRenderingContext) HasFuncHint() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextHint(
		this.ref,
	)
}

// FuncHint returns the method "WebGLRenderingContext.hint".
func (this WebGLRenderingContext) FuncHint() (fn js.Func[func(target GLenum, mode GLenum)]) {
	bindings.FuncWebGLRenderingContextHint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Hint calls the method "WebGLRenderingContext.hint".
func (this WebGLRenderingContext) Hint(target GLenum, mode GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextHint(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(mode),
	)

	return
}

// HasFuncIsBuffer returns true if the method "WebGLRenderingContext.isBuffer" exists.
func (this WebGLRenderingContext) HasFuncIsBuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsBuffer(
		this.ref,
	)
}

// FuncIsBuffer returns the method "WebGLRenderingContext.isBuffer".
func (this WebGLRenderingContext) FuncIsBuffer() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsBuffer calls the method "WebGLRenderingContext.isBuffer".
func (this WebGLRenderingContext) IsBuffer(buffer WebGLBuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsBuffer(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryIsBuffer calls the method "WebGLRenderingContext.isBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsBuffer(buffer WebGLBuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncIsEnabled returns true if the method "WebGLRenderingContext.isEnabled" exists.
func (this WebGLRenderingContext) HasFuncIsEnabled() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsEnabled(
		this.ref,
	)
}

// FuncIsEnabled returns the method "WebGLRenderingContext.isEnabled".
func (this WebGLRenderingContext) FuncIsEnabled() (fn js.Func[func(cap GLenum) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsEnabled calls the method "WebGLRenderingContext.isEnabled".
func (this WebGLRenderingContext) IsEnabled(cap GLenum) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsEnabled(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryIsEnabled calls the method "WebGLRenderingContext.isEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsEnabled(cap GLenum) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncIsFramebuffer returns true if the method "WebGLRenderingContext.isFramebuffer" exists.
func (this WebGLRenderingContext) HasFuncIsFramebuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsFramebuffer(
		this.ref,
	)
}

// FuncIsFramebuffer returns the method "WebGLRenderingContext.isFramebuffer".
func (this WebGLRenderingContext) FuncIsFramebuffer() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsFramebuffer calls the method "WebGLRenderingContext.isFramebuffer".
func (this WebGLRenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsFramebuffer(
		this.ref, js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryIsFramebuffer calls the method "WebGLRenderingContext.isFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasFuncIsProgram returns true if the method "WebGLRenderingContext.isProgram" exists.
func (this WebGLRenderingContext) HasFuncIsProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsProgram(
		this.ref,
	)
}

// FuncIsProgram returns the method "WebGLRenderingContext.isProgram".
func (this WebGLRenderingContext) FuncIsProgram() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsProgram calls the method "WebGLRenderingContext.isProgram".
func (this WebGLRenderingContext) IsProgram(program WebGLProgram) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryIsProgram calls the method "WebGLRenderingContext.isProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsProgram(program WebGLProgram) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncIsRenderbuffer returns true if the method "WebGLRenderingContext.isRenderbuffer" exists.
func (this WebGLRenderingContext) HasFuncIsRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsRenderbuffer(
		this.ref,
	)
}

// FuncIsRenderbuffer returns the method "WebGLRenderingContext.isRenderbuffer".
func (this WebGLRenderingContext) FuncIsRenderbuffer() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsRenderbuffer calls the method "WebGLRenderingContext.isRenderbuffer".
func (this WebGLRenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryIsRenderbuffer calls the method "WebGLRenderingContext.isRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncIsShader returns true if the method "WebGLRenderingContext.isShader" exists.
func (this WebGLRenderingContext) HasFuncIsShader() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsShader(
		this.ref,
	)
}

// FuncIsShader returns the method "WebGLRenderingContext.isShader".
func (this WebGLRenderingContext) FuncIsShader() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsShader calls the method "WebGLRenderingContext.isShader".
func (this WebGLRenderingContext) IsShader(shader WebGLShader) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryIsShader calls the method "WebGLRenderingContext.isShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsShader(shader WebGLShader) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncIsTexture returns true if the method "WebGLRenderingContext.isTexture" exists.
func (this WebGLRenderingContext) HasFuncIsTexture() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextIsTexture(
		this.ref,
	)
}

// FuncIsTexture returns the method "WebGLRenderingContext.isTexture".
func (this WebGLRenderingContext) FuncIsTexture() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	bindings.FuncWebGLRenderingContextIsTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTexture calls the method "WebGLRenderingContext.isTexture".
func (this WebGLRenderingContext) IsTexture(texture WebGLTexture) (ret GLboolean) {
	bindings.CallWebGLRenderingContextIsTexture(
		this.ref, js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryIsTexture calls the method "WebGLRenderingContext.isTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryIsTexture(texture WebGLTexture) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextIsTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasFuncLineWidth returns true if the method "WebGLRenderingContext.lineWidth" exists.
func (this WebGLRenderingContext) HasFuncLineWidth() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextLineWidth(
		this.ref,
	)
}

// FuncLineWidth returns the method "WebGLRenderingContext.lineWidth".
func (this WebGLRenderingContext) FuncLineWidth() (fn js.Func[func(width GLfloat)]) {
	bindings.FuncWebGLRenderingContextLineWidth(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LineWidth calls the method "WebGLRenderingContext.lineWidth".
func (this WebGLRenderingContext) LineWidth(width GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextLineWidth(
		this.ref, js.Pointer(&ret),
		float32(width),
	)

	return
}

// TryLineWidth calls the method "WebGLRenderingContext.lineWidth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryLineWidth(width GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextLineWidth(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(width),
	)

	return
}

// HasFuncLinkProgram returns true if the method "WebGLRenderingContext.linkProgram" exists.
func (this WebGLRenderingContext) HasFuncLinkProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextLinkProgram(
		this.ref,
	)
}

// FuncLinkProgram returns the method "WebGLRenderingContext.linkProgram".
func (this WebGLRenderingContext) FuncLinkProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGLRenderingContextLinkProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LinkProgram calls the method "WebGLRenderingContext.linkProgram".
func (this WebGLRenderingContext) LinkProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextLinkProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryLinkProgram calls the method "WebGLRenderingContext.linkProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryLinkProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextLinkProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncPixelStorei returns true if the method "WebGLRenderingContext.pixelStorei" exists.
func (this WebGLRenderingContext) HasFuncPixelStorei() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextPixelStorei(
		this.ref,
	)
}

// FuncPixelStorei returns the method "WebGLRenderingContext.pixelStorei".
func (this WebGLRenderingContext) FuncPixelStorei() (fn js.Func[func(pname GLenum, param GLint)]) {
	bindings.FuncWebGLRenderingContextPixelStorei(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PixelStorei calls the method "WebGLRenderingContext.pixelStorei".
func (this WebGLRenderingContext) PixelStorei(pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextPixelStorei(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
		int32(param),
	)

	return
}

// HasFuncPolygonOffset returns true if the method "WebGLRenderingContext.polygonOffset" exists.
func (this WebGLRenderingContext) HasFuncPolygonOffset() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextPolygonOffset(
		this.ref,
	)
}

// FuncPolygonOffset returns the method "WebGLRenderingContext.polygonOffset".
func (this WebGLRenderingContext) FuncPolygonOffset() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	bindings.FuncWebGLRenderingContextPolygonOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PolygonOffset calls the method "WebGLRenderingContext.polygonOffset".
func (this WebGLRenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextPolygonOffset(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(factor),
		float32(units),
	)

	return
}

// HasFuncRenderbufferStorage returns true if the method "WebGLRenderingContext.renderbufferStorage" exists.
func (this WebGLRenderingContext) HasFuncRenderbufferStorage() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextRenderbufferStorage(
		this.ref,
	)
}

// FuncRenderbufferStorage returns the method "WebGLRenderingContext.renderbufferStorage".
func (this WebGLRenderingContext) FuncRenderbufferStorage() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGLRenderingContextRenderbufferStorage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RenderbufferStorage calls the method "WebGLRenderingContext.renderbufferStorage".
func (this WebGLRenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextRenderbufferStorage(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncSampleCoverage returns true if the method "WebGLRenderingContext.sampleCoverage" exists.
func (this WebGLRenderingContext) HasFuncSampleCoverage() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextSampleCoverage(
		this.ref,
	)
}

// FuncSampleCoverage returns the method "WebGLRenderingContext.sampleCoverage".
func (this WebGLRenderingContext) FuncSampleCoverage() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	bindings.FuncWebGLRenderingContextSampleCoverage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SampleCoverage calls the method "WebGLRenderingContext.sampleCoverage".
func (this WebGLRenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (ret js.Void) {
	bindings.CallWebGLRenderingContextSampleCoverage(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// HasFuncScissor returns true if the method "WebGLRenderingContext.scissor" exists.
func (this WebGLRenderingContext) HasFuncScissor() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextScissor(
		this.ref,
	)
}

// FuncScissor returns the method "WebGLRenderingContext.scissor".
func (this WebGLRenderingContext) FuncScissor() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGLRenderingContextScissor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scissor calls the method "WebGLRenderingContext.scissor".
func (this WebGLRenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextScissor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncShaderSource returns true if the method "WebGLRenderingContext.shaderSource" exists.
func (this WebGLRenderingContext) HasFuncShaderSource() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextShaderSource(
		this.ref,
	)
}

// FuncShaderSource returns the method "WebGLRenderingContext.shaderSource".
func (this WebGLRenderingContext) FuncShaderSource() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	bindings.FuncWebGLRenderingContextShaderSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShaderSource calls the method "WebGLRenderingContext.shaderSource".
func (this WebGLRenderingContext) ShaderSource(shader WebGLShader, source js.String) (ret js.Void) {
	bindings.CallWebGLRenderingContextShaderSource(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// HasFuncStencilFunc returns true if the method "WebGLRenderingContext.stencilFunc" exists.
func (this WebGLRenderingContext) HasFuncStencilFunc() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilFunc(
		this.ref,
	)
}

// FuncStencilFunc returns the method "WebGLRenderingContext.stencilFunc".
func (this WebGLRenderingContext) FuncStencilFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	bindings.FuncWebGLRenderingContextStencilFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilFunc calls the method "WebGLRenderingContext.stencilFunc".
func (this WebGLRenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilFunc(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasFuncStencilFuncSeparate returns true if the method "WebGLRenderingContext.stencilFuncSeparate" exists.
func (this WebGLRenderingContext) HasFuncStencilFuncSeparate() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilFuncSeparate(
		this.ref,
	)
}

// FuncStencilFuncSeparate returns the method "WebGLRenderingContext.stencilFuncSeparate".
func (this WebGLRenderingContext) FuncStencilFuncSeparate() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	bindings.FuncWebGLRenderingContextStencilFuncSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilFuncSeparate calls the method "WebGLRenderingContext.stencilFuncSeparate".
func (this WebGLRenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilFuncSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasFuncStencilMask returns true if the method "WebGLRenderingContext.stencilMask" exists.
func (this WebGLRenderingContext) HasFuncStencilMask() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilMask(
		this.ref,
	)
}

// FuncStencilMask returns the method "WebGLRenderingContext.stencilMask".
func (this WebGLRenderingContext) FuncStencilMask() (fn js.Func[func(mask GLuint)]) {
	bindings.FuncWebGLRenderingContextStencilMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilMask calls the method "WebGLRenderingContext.stencilMask".
func (this WebGLRenderingContext) StencilMask(mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilMask(
		this.ref, js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryStencilMask calls the method "WebGLRenderingContext.stencilMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryStencilMask(mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextStencilMask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasFuncStencilMaskSeparate returns true if the method "WebGLRenderingContext.stencilMaskSeparate" exists.
func (this WebGLRenderingContext) HasFuncStencilMaskSeparate() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilMaskSeparate(
		this.ref,
	)
}

// FuncStencilMaskSeparate returns the method "WebGLRenderingContext.stencilMaskSeparate".
func (this WebGLRenderingContext) FuncStencilMaskSeparate() (fn js.Func[func(face GLenum, mask GLuint)]) {
	bindings.FuncWebGLRenderingContextStencilMaskSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilMaskSeparate calls the method "WebGLRenderingContext.stencilMaskSeparate".
func (this WebGLRenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilMaskSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(mask),
	)

	return
}

// HasFuncStencilOp returns true if the method "WebGLRenderingContext.stencilOp" exists.
func (this WebGLRenderingContext) HasFuncStencilOp() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilOp(
		this.ref,
	)
}

// FuncStencilOp returns the method "WebGLRenderingContext.stencilOp".
func (this WebGLRenderingContext) FuncStencilOp() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	bindings.FuncWebGLRenderingContextStencilOp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilOp calls the method "WebGLRenderingContext.stencilOp".
func (this WebGLRenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilOp(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasFuncStencilOpSeparate returns true if the method "WebGLRenderingContext.stencilOpSeparate" exists.
func (this WebGLRenderingContext) HasFuncStencilOpSeparate() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextStencilOpSeparate(
		this.ref,
	)
}

// FuncStencilOpSeparate returns the method "WebGLRenderingContext.stencilOpSeparate".
func (this WebGLRenderingContext) FuncStencilOpSeparate() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	bindings.FuncWebGLRenderingContextStencilOpSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilOpSeparate calls the method "WebGLRenderingContext.stencilOpSeparate".
func (this WebGLRenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextStencilOpSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasFuncTexParameterf returns true if the method "WebGLRenderingContext.texParameterf" exists.
func (this WebGLRenderingContext) HasFuncTexParameterf() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexParameterf(
		this.ref,
	)
}

// FuncTexParameterf returns the method "WebGLRenderingContext.texParameterf".
func (this WebGLRenderingContext) FuncTexParameterf() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	bindings.FuncWebGLRenderingContextTexParameterf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexParameterf calls the method "WebGLRenderingContext.texParameterf".
func (this WebGLRenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexParameterf(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// HasFuncTexParameteri returns true if the method "WebGLRenderingContext.texParameteri" exists.
func (this WebGLRenderingContext) HasFuncTexParameteri() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexParameteri(
		this.ref,
	)
}

// FuncTexParameteri returns the method "WebGLRenderingContext.texParameteri".
func (this WebGLRenderingContext) FuncTexParameteri() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	bindings.FuncWebGLRenderingContextTexParameteri(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexParameteri calls the method "WebGLRenderingContext.texParameteri".
func (this WebGLRenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexParameteri(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// HasFuncUniform1f returns true if the method "WebGLRenderingContext.uniform1f" exists.
func (this WebGLRenderingContext) HasFuncUniform1f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform1f(
		this.ref,
	)
}

// FuncUniform1f returns the method "WebGLRenderingContext.uniform1f".
func (this WebGLRenderingContext) FuncUniform1f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	bindings.FuncWebGLRenderingContextUniform1f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1f calls the method "WebGLRenderingContext.uniform1f".
func (this WebGLRenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
	)

	return
}

// HasFuncUniform2f returns true if the method "WebGLRenderingContext.uniform2f" exists.
func (this WebGLRenderingContext) HasFuncUniform2f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform2f(
		this.ref,
	)
}

// FuncUniform2f returns the method "WebGLRenderingContext.uniform2f".
func (this WebGLRenderingContext) FuncUniform2f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	bindings.FuncWebGLRenderingContextUniform2f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2f calls the method "WebGLRenderingContext.uniform2f".
func (this WebGLRenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// HasFuncUniform3f returns true if the method "WebGLRenderingContext.uniform3f" exists.
func (this WebGLRenderingContext) HasFuncUniform3f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform3f(
		this.ref,
	)
}

// FuncUniform3f returns the method "WebGLRenderingContext.uniform3f".
func (this WebGLRenderingContext) FuncUniform3f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	bindings.FuncWebGLRenderingContextUniform3f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3f calls the method "WebGLRenderingContext.uniform3f".
func (this WebGLRenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncUniform4f returns true if the method "WebGLRenderingContext.uniform4f" exists.
func (this WebGLRenderingContext) HasFuncUniform4f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform4f(
		this.ref,
	)
}

// FuncUniform4f returns the method "WebGLRenderingContext.uniform4f".
func (this WebGLRenderingContext) FuncUniform4f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	bindings.FuncWebGLRenderingContextUniform4f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4f calls the method "WebGLRenderingContext.uniform4f".
func (this WebGLRenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasFuncUniform1i returns true if the method "WebGLRenderingContext.uniform1i" exists.
func (this WebGLRenderingContext) HasFuncUniform1i() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform1i(
		this.ref,
	)
}

// FuncUniform1i returns the method "WebGLRenderingContext.uniform1i".
func (this WebGLRenderingContext) FuncUniform1i() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	bindings.FuncWebGLRenderingContextUniform1i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1i calls the method "WebGLRenderingContext.uniform1i".
func (this WebGLRenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
	)

	return
}

// HasFuncUniform2i returns true if the method "WebGLRenderingContext.uniform2i" exists.
func (this WebGLRenderingContext) HasFuncUniform2i() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform2i(
		this.ref,
	)
}

// FuncUniform2i returns the method "WebGLRenderingContext.uniform2i".
func (this WebGLRenderingContext) FuncUniform2i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	bindings.FuncWebGLRenderingContextUniform2i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2i calls the method "WebGLRenderingContext.uniform2i".
func (this WebGLRenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncUniform3i returns true if the method "WebGLRenderingContext.uniform3i" exists.
func (this WebGLRenderingContext) HasFuncUniform3i() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform3i(
		this.ref,
	)
}

// FuncUniform3i returns the method "WebGLRenderingContext.uniform3i".
func (this WebGLRenderingContext) FuncUniform3i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	bindings.FuncWebGLRenderingContextUniform3i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3i calls the method "WebGLRenderingContext.uniform3i".
func (this WebGLRenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// HasFuncUniform4i returns true if the method "WebGLRenderingContext.uniform4i" exists.
func (this WebGLRenderingContext) HasFuncUniform4i() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform4i(
		this.ref,
	)
}

// FuncUniform4i returns the method "WebGLRenderingContext.uniform4i".
func (this WebGLRenderingContext) FuncUniform4i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	bindings.FuncWebGLRenderingContextUniform4i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4i calls the method "WebGLRenderingContext.uniform4i".
func (this WebGLRenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasFuncUseProgram returns true if the method "WebGLRenderingContext.useProgram" exists.
func (this WebGLRenderingContext) HasFuncUseProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUseProgram(
		this.ref,
	)
}

// FuncUseProgram returns the method "WebGLRenderingContext.useProgram".
func (this WebGLRenderingContext) FuncUseProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGLRenderingContextUseProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UseProgram calls the method "WebGLRenderingContext.useProgram".
func (this WebGLRenderingContext) UseProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextUseProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryUseProgram calls the method "WebGLRenderingContext.useProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryUseProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextUseProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncValidateProgram returns true if the method "WebGLRenderingContext.validateProgram" exists.
func (this WebGLRenderingContext) HasFuncValidateProgram() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextValidateProgram(
		this.ref,
	)
}

// FuncValidateProgram returns the method "WebGLRenderingContext.validateProgram".
func (this WebGLRenderingContext) FuncValidateProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGLRenderingContextValidateProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ValidateProgram calls the method "WebGLRenderingContext.validateProgram".
func (this WebGLRenderingContext) ValidateProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGLRenderingContextValidateProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryValidateProgram calls the method "WebGLRenderingContext.validateProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryValidateProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextValidateProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncVertexAttrib1f returns true if the method "WebGLRenderingContext.vertexAttrib1f" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib1f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib1f(
		this.ref,
	)
}

// FuncVertexAttrib1f returns the method "WebGLRenderingContext.vertexAttrib1f".
func (this WebGLRenderingContext) FuncVertexAttrib1f() (fn js.Func[func(index GLuint, x GLfloat)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib1f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib1f calls the method "WebGLRenderingContext.vertexAttrib1f".
func (this WebGLRenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib1f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
	)

	return
}

// HasFuncVertexAttrib2f returns true if the method "WebGLRenderingContext.vertexAttrib2f" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib2f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib2f(
		this.ref,
	)
}

// FuncVertexAttrib2f returns the method "WebGLRenderingContext.vertexAttrib2f".
func (this WebGLRenderingContext) FuncVertexAttrib2f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib2f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib2f calls the method "WebGLRenderingContext.vertexAttrib2f".
func (this WebGLRenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib2f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// HasFuncVertexAttrib3f returns true if the method "WebGLRenderingContext.vertexAttrib3f" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib3f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib3f(
		this.ref,
	)
}

// FuncVertexAttrib3f returns the method "WebGLRenderingContext.vertexAttrib3f".
func (this WebGLRenderingContext) FuncVertexAttrib3f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib3f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib3f calls the method "WebGLRenderingContext.vertexAttrib3f".
func (this WebGLRenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib3f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncVertexAttrib4f returns true if the method "WebGLRenderingContext.vertexAttrib4f" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib4f() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib4f(
		this.ref,
	)
}

// FuncVertexAttrib4f returns the method "WebGLRenderingContext.vertexAttrib4f".
func (this WebGLRenderingContext) FuncVertexAttrib4f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib4f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib4f calls the method "WebGLRenderingContext.vertexAttrib4f".
func (this WebGLRenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib4f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasFuncVertexAttrib1fv returns true if the method "WebGLRenderingContext.vertexAttrib1fv" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib1fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib1fv(
		this.ref,
	)
}

// FuncVertexAttrib1fv returns the method "WebGLRenderingContext.vertexAttrib1fv".
func (this WebGLRenderingContext) FuncVertexAttrib1fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib1fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib1fv calls the method "WebGLRenderingContext.vertexAttrib1fv".
func (this WebGLRenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib1fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib2fv returns true if the method "WebGLRenderingContext.vertexAttrib2fv" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib2fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib2fv(
		this.ref,
	)
}

// FuncVertexAttrib2fv returns the method "WebGLRenderingContext.vertexAttrib2fv".
func (this WebGLRenderingContext) FuncVertexAttrib2fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib2fv calls the method "WebGLRenderingContext.vertexAttrib2fv".
func (this WebGLRenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib3fv returns true if the method "WebGLRenderingContext.vertexAttrib3fv" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib3fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib3fv(
		this.ref,
	)
}

// FuncVertexAttrib3fv returns the method "WebGLRenderingContext.vertexAttrib3fv".
func (this WebGLRenderingContext) FuncVertexAttrib3fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib3fv calls the method "WebGLRenderingContext.vertexAttrib3fv".
func (this WebGLRenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib4fv returns true if the method "WebGLRenderingContext.vertexAttrib4fv" exists.
func (this WebGLRenderingContext) HasFuncVertexAttrib4fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttrib4fv(
		this.ref,
	)
}

// FuncVertexAttrib4fv returns the method "WebGLRenderingContext.vertexAttrib4fv".
func (this WebGLRenderingContext) FuncVertexAttrib4fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGLRenderingContextVertexAttrib4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib4fv calls the method "WebGLRenderingContext.vertexAttrib4fv".
func (this WebGLRenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttrib4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttribPointer returns true if the method "WebGLRenderingContext.vertexAttribPointer" exists.
func (this WebGLRenderingContext) HasFuncVertexAttribPointer() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextVertexAttribPointer(
		this.ref,
	)
}

// FuncVertexAttribPointer returns the method "WebGLRenderingContext.vertexAttribPointer".
func (this WebGLRenderingContext) FuncVertexAttribPointer() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	bindings.FuncWebGLRenderingContextVertexAttribPointer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribPointer calls the method "WebGLRenderingContext.vertexAttribPointer".
func (this WebGLRenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGLRenderingContextVertexAttribPointer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// HasFuncViewport returns true if the method "WebGLRenderingContext.viewport" exists.
func (this WebGLRenderingContext) HasFuncViewport() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextViewport(
		this.ref,
	)
}

// FuncViewport returns the method "WebGLRenderingContext.viewport".
func (this WebGLRenderingContext) FuncViewport() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGLRenderingContextViewport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Viewport calls the method "WebGLRenderingContext.viewport".
func (this WebGLRenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGLRenderingContextViewport(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncMakeXRCompatible returns true if the method "WebGLRenderingContext.makeXRCompatible" exists.
func (this WebGLRenderingContext) HasFuncMakeXRCompatible() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextMakeXRCompatible(
		this.ref,
	)
}

// FuncMakeXRCompatible returns the method "WebGLRenderingContext.makeXRCompatible".
func (this WebGLRenderingContext) FuncMakeXRCompatible() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWebGLRenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MakeXRCompatible calls the method "WebGLRenderingContext.makeXRCompatible".
func (this WebGLRenderingContext) MakeXRCompatible() (ret js.Promise[js.Void]) {
	bindings.CallWebGLRenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMakeXRCompatible calls the method "WebGLRenderingContext.makeXRCompatible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGLRenderingContext) TryMakeXRCompatible() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGLRenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBufferData returns true if the method "WebGLRenderingContext.bufferData" exists.
func (this WebGLRenderingContext) HasFuncBufferData() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBufferData(
		this.ref,
	)
}

// FuncBufferData returns the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) FuncBufferData() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	bindings.FuncWebGLRenderingContextBufferData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData calls the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// HasFuncBufferData1 returns true if the method "WebGLRenderingContext.bufferData" exists.
func (this WebGLRenderingContext) HasFuncBufferData1() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBufferData1(
		this.ref,
	)
}

// FuncBufferData1 returns the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) FuncBufferData1() (fn js.Func[func(target GLenum, data AllowSharedBufferSource, usage GLenum)]) {
	bindings.FuncWebGLRenderingContextBufferData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData1 calls the method "WebGLRenderingContext.bufferData".
func (this WebGLRenderingContext) BufferData1(target GLenum, data AllowSharedBufferSource, usage GLenum) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferData1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		data.Ref(),
		uint32(usage),
	)

	return
}

// HasFuncBufferSubData returns true if the method "WebGLRenderingContext.bufferSubData" exists.
func (this WebGLRenderingContext) HasFuncBufferSubData() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextBufferSubData(
		this.ref,
	)
}

// FuncBufferSubData returns the method "WebGLRenderingContext.bufferSubData".
func (this WebGLRenderingContext) FuncBufferSubData() (fn js.Func[func(target GLenum, offset GLintptr, data AllowSharedBufferSource)]) {
	bindings.FuncWebGLRenderingContextBufferSubData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferSubData calls the method "WebGLRenderingContext.bufferSubData".
func (this WebGLRenderingContext) BufferSubData(target GLenum, offset GLintptr, data AllowSharedBufferSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextBufferSubData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(offset),
		data.Ref(),
	)

	return
}

// HasFuncCompressedTexImage2D returns true if the method "WebGLRenderingContext.compressedTexImage2D" exists.
func (this WebGLRenderingContext) HasFuncCompressedTexImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCompressedTexImage2D(
		this.ref,
	)
}

// FuncCompressedTexImage2D returns the method "WebGLRenderingContext.compressedTexImage2D".
func (this WebGLRenderingContext) FuncCompressedTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data js.ArrayBufferView)]) {
	bindings.FuncWebGLRenderingContextCompressedTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage2D calls the method "WebGLRenderingContext.compressedTexImage2D".
func (this WebGLRenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, data js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompressedTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage2D returns true if the method "WebGLRenderingContext.compressedTexSubImage2D" exists.
func (this WebGLRenderingContext) HasFuncCompressedTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextCompressedTexSubImage2D(
		this.ref,
	)
}

// FuncCompressedTexSubImage2D returns the method "WebGLRenderingContext.compressedTexSubImage2D".
func (this WebGLRenderingContext) FuncCompressedTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data js.ArrayBufferView)]) {
	bindings.FuncWebGLRenderingContextCompressedTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage2D calls the method "WebGLRenderingContext.compressedTexSubImage2D".
func (this WebGLRenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, data js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextCompressedTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncReadPixels returns true if the method "WebGLRenderingContext.readPixels" exists.
func (this WebGLRenderingContext) HasFuncReadPixels() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextReadPixels(
		this.ref,
	)
}

// FuncReadPixels returns the method "WebGLRenderingContext.readPixels".
func (this WebGLRenderingContext) FuncReadPixels() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	bindings.FuncWebGLRenderingContextReadPixels(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadPixels calls the method "WebGLRenderingContext.readPixels".
func (this WebGLRenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextReadPixels(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D returns true if the method "WebGLRenderingContext.texImage2D" exists.
func (this WebGLRenderingContext) HasFuncTexImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexImage2D(
		this.ref,
	)
}

// FuncTexImage2D returns the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) FuncTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	bindings.FuncWebGLRenderingContextTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D calls the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D1 returns true if the method "WebGLRenderingContext.texImage2D" exists.
func (this WebGLRenderingContext) HasFuncTexImage2D1() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexImage2D1(
		this.ref,
	)
}

// FuncTexImage2D1 returns the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) FuncTexImage2D1() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGLRenderingContextTexImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D1 calls the method "WebGLRenderingContext.texImage2D".
func (this WebGLRenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasFuncTexSubImage2D returns true if the method "WebGLRenderingContext.texSubImage2D" exists.
func (this WebGLRenderingContext) HasFuncTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexSubImage2D(
		this.ref,
	)
}

// FuncTexSubImage2D returns the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) FuncTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	bindings.FuncWebGLRenderingContextTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D calls the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage2D1 returns true if the method "WebGLRenderingContext.texSubImage2D" exists.
func (this WebGLRenderingContext) HasFuncTexSubImage2D1() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextTexSubImage2D1(
		this.ref,
	)
}

// FuncTexSubImage2D1 returns the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) FuncTexSubImage2D1() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGLRenderingContextTexSubImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D1 calls the method "WebGLRenderingContext.texSubImage2D".
func (this WebGLRenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGLRenderingContextTexSubImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncUniform1fv returns true if the method "WebGLRenderingContext.uniform1fv" exists.
func (this WebGLRenderingContext) HasFuncUniform1fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform1fv(
		this.ref,
	)
}

// FuncUniform1fv returns the method "WebGLRenderingContext.uniform1fv".
func (this WebGLRenderingContext) FuncUniform1fv() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	bindings.FuncWebGLRenderingContextUniform1fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1fv calls the method "WebGLRenderingContext.uniform1fv".
func (this WebGLRenderingContext) Uniform1fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform2fv returns true if the method "WebGLRenderingContext.uniform2fv" exists.
func (this WebGLRenderingContext) HasFuncUniform2fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform2fv(
		this.ref,
	)
}

// FuncUniform2fv returns the method "WebGLRenderingContext.uniform2fv".
func (this WebGLRenderingContext) FuncUniform2fv() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	bindings.FuncWebGLRenderingContextUniform2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2fv calls the method "WebGLRenderingContext.uniform2fv".
func (this WebGLRenderingContext) Uniform2fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform3fv returns true if the method "WebGLRenderingContext.uniform3fv" exists.
func (this WebGLRenderingContext) HasFuncUniform3fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform3fv(
		this.ref,
	)
}

// FuncUniform3fv returns the method "WebGLRenderingContext.uniform3fv".
func (this WebGLRenderingContext) FuncUniform3fv() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	bindings.FuncWebGLRenderingContextUniform3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3fv calls the method "WebGLRenderingContext.uniform3fv".
func (this WebGLRenderingContext) Uniform3fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform4fv returns true if the method "WebGLRenderingContext.uniform4fv" exists.
func (this WebGLRenderingContext) HasFuncUniform4fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform4fv(
		this.ref,
	)
}

// FuncUniform4fv returns the method "WebGLRenderingContext.uniform4fv".
func (this WebGLRenderingContext) FuncUniform4fv() (fn js.Func[func(location WebGLUniformLocation, v Float32List)]) {
	bindings.FuncWebGLRenderingContextUniform4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4fv calls the method "WebGLRenderingContext.uniform4fv".
func (this WebGLRenderingContext) Uniform4fv(location WebGLUniformLocation, v Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform1iv returns true if the method "WebGLRenderingContext.uniform1iv" exists.
func (this WebGLRenderingContext) HasFuncUniform1iv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform1iv(
		this.ref,
	)
}

// FuncUniform1iv returns the method "WebGLRenderingContext.uniform1iv".
func (this WebGLRenderingContext) FuncUniform1iv() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	bindings.FuncWebGLRenderingContextUniform1iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1iv calls the method "WebGLRenderingContext.uniform1iv".
func (this WebGLRenderingContext) Uniform1iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform1iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform2iv returns true if the method "WebGLRenderingContext.uniform2iv" exists.
func (this WebGLRenderingContext) HasFuncUniform2iv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform2iv(
		this.ref,
	)
}

// FuncUniform2iv returns the method "WebGLRenderingContext.uniform2iv".
func (this WebGLRenderingContext) FuncUniform2iv() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	bindings.FuncWebGLRenderingContextUniform2iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2iv calls the method "WebGLRenderingContext.uniform2iv".
func (this WebGLRenderingContext) Uniform2iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform2iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform3iv returns true if the method "WebGLRenderingContext.uniform3iv" exists.
func (this WebGLRenderingContext) HasFuncUniform3iv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform3iv(
		this.ref,
	)
}

// FuncUniform3iv returns the method "WebGLRenderingContext.uniform3iv".
func (this WebGLRenderingContext) FuncUniform3iv() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	bindings.FuncWebGLRenderingContextUniform3iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3iv calls the method "WebGLRenderingContext.uniform3iv".
func (this WebGLRenderingContext) Uniform3iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform3iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniform4iv returns true if the method "WebGLRenderingContext.uniform4iv" exists.
func (this WebGLRenderingContext) HasFuncUniform4iv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniform4iv(
		this.ref,
	)
}

// FuncUniform4iv returns the method "WebGLRenderingContext.uniform4iv".
func (this WebGLRenderingContext) FuncUniform4iv() (fn js.Func[func(location WebGLUniformLocation, v Int32List)]) {
	bindings.FuncWebGLRenderingContextUniform4iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4iv calls the method "WebGLRenderingContext.uniform4iv".
func (this WebGLRenderingContext) Uniform4iv(location WebGLUniformLocation, v Int32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniform4iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		v.Ref(),
	)

	return
}

// HasFuncUniformMatrix2fv returns true if the method "WebGLRenderingContext.uniformMatrix2fv" exists.
func (this WebGLRenderingContext) HasFuncUniformMatrix2fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniformMatrix2fv(
		this.ref,
	)
}

// FuncUniformMatrix2fv returns the method "WebGLRenderingContext.uniformMatrix2fv".
func (this WebGLRenderingContext) FuncUniformMatrix2fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	bindings.FuncWebGLRenderingContextUniformMatrix2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2fv calls the method "WebGLRenderingContext.uniformMatrix2fv".
func (this WebGLRenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// HasFuncUniformMatrix3fv returns true if the method "WebGLRenderingContext.uniformMatrix3fv" exists.
func (this WebGLRenderingContext) HasFuncUniformMatrix3fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniformMatrix3fv(
		this.ref,
	)
}

// FuncUniformMatrix3fv returns the method "WebGLRenderingContext.uniformMatrix3fv".
func (this WebGLRenderingContext) FuncUniformMatrix3fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	bindings.FuncWebGLRenderingContextUniformMatrix3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3fv calls the method "WebGLRenderingContext.uniformMatrix3fv".
func (this WebGLRenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		value.Ref(),
	)

	return
}

// HasFuncUniformMatrix4fv returns true if the method "WebGLRenderingContext.uniformMatrix4fv" exists.
func (this WebGLRenderingContext) HasFuncUniformMatrix4fv() bool {
	return js.True == bindings.HasFuncWebGLRenderingContextUniformMatrix4fv(
		this.ref,
	)
}

// FuncUniformMatrix4fv returns the method "WebGLRenderingContext.uniformMatrix4fv".
func (this WebGLRenderingContext) FuncUniformMatrix4fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, value Float32List)]) {
	bindings.FuncWebGLRenderingContextUniformMatrix4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4fv calls the method "WebGLRenderingContext.uniformMatrix4fv".
func (this WebGLRenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, value Float32List) (ret js.Void) {
	bindings.CallWebGLRenderingContextUniformMatrix4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

type WebGLSampler struct {
	WebGLObject
}

func (this WebGLSampler) Once() WebGLSampler {
	this.ref.Once()
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
	this.ref.Free()
}

type WebGLSync struct {
	WebGLObject
}

func (this WebGLSync) Once() WebGLSync {
	this.ref.Once()
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
	this.ref.Free()
}

type GLuint64 uint64

type WebGLTransformFeedback struct {
	WebGLObject
}

func (this WebGLTransformFeedback) Once() WebGLTransformFeedback {
	this.ref.Once()
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
	this.ref.Free()
}

type WebGLVertexArrayObject struct {
	WebGLObject
}

func (this WebGLVertexArrayObject) Once() WebGLVertexArrayObject {
	this.ref.Once()
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
	this.ref.Free()
}

type WebGL2RenderingContext struct {
	ref js.Ref
}

func (this WebGL2RenderingContext) Once() WebGL2RenderingContext {
	this.ref.Once()
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
	this.ref.Free()
}

// Canvas returns the value of property "WebGL2RenderingContext.canvas".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextCanvas(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferWidth returns the value of property "WebGL2RenderingContext.drawingBufferWidth".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferWidth() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferHeight returns the value of property "WebGL2RenderingContext.drawingBufferHeight".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferHeight() (ret GLsizei, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DrawingBufferColorSpace returns the value of property "WebGL2RenderingContext.drawingBufferColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) DrawingBufferColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextDrawingBufferColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDrawingBufferColorSpace sets the value of property "WebGL2RenderingContext.drawingBufferColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGL2RenderingContext) SetDrawingBufferColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGL2RenderingContextDrawingBufferColorSpace(
		this.ref,
		uint32(val),
	)
}

// UnpackColorSpace returns the value of property "WebGL2RenderingContext.unpackColorSpace".
//
// It returns ok=false if there is no such property.
func (this WebGL2RenderingContext) UnpackColorSpace() (ret PredefinedColorSpace, ok bool) {
	ok = js.True == bindings.GetWebGL2RenderingContextUnpackColorSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUnpackColorSpace sets the value of property "WebGL2RenderingContext.unpackColorSpace" to val.
//
// It returns false if the property cannot be set.
func (this WebGL2RenderingContext) SetUnpackColorSpace(val PredefinedColorSpace) bool {
	return js.True == bindings.SetWebGL2RenderingContextUnpackColorSpace(
		this.ref,
		uint32(val),
	)
}

// HasFuncBufferData returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasFuncBufferData() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferData(
		this.ref,
	)
}

// FuncBufferData returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) FuncBufferData() (fn js.Func[func(target GLenum, size GLsizeiptr, usage GLenum)]) {
	bindings.FuncWebGL2RenderingContextBufferData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData(target GLenum, size GLsizeiptr, usage GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(size),
		uint32(usage),
	)

	return
}

// HasFuncBufferData1 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasFuncBufferData1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferData1(
		this.ref,
	)
}

// FuncBufferData1 returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) FuncBufferData1() (fn js.Func[func(target GLenum, srcData AllowSharedBufferSource, usage GLenum)]) {
	bindings.FuncWebGL2RenderingContextBufferData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData1 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData1(target GLenum, srcData AllowSharedBufferSource, usage GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
	)

	return
}

// HasFuncBufferSubData returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncBufferSubData() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferSubData(
		this.ref,
	)
}

// FuncBufferSubData returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) FuncBufferSubData() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource)]) {
	bindings.FuncWebGL2RenderingContextBufferSubData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferSubData calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData(target GLenum, dstByteOffset GLintptr, srcData AllowSharedBufferSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
	)

	return
}

// HasFuncBufferData2 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasFuncBufferData2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferData2(
		this.ref,
	)
}

// FuncBufferData2 returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) FuncBufferData2() (fn js.Func[func(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint)]) {
	bindings.FuncWebGL2RenderingContextBufferData2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData2 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData2(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// HasFuncBufferData3 returns true if the method "WebGL2RenderingContext.bufferData" exists.
func (this WebGL2RenderingContext) HasFuncBufferData3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferData3(
		this.ref,
	)
}

// FuncBufferData3 returns the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) FuncBufferData3() (fn js.Func[func(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextBufferData3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferData3 calls the method "WebGL2RenderingContext.bufferData".
func (this WebGL2RenderingContext) BufferData3(target GLenum, srcData js.ArrayBufferView, usage GLenum, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferData3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		srcData.Ref(),
		uint32(usage),
		uint32(srcOffset),
	)

	return
}

// HasFuncBufferSubData1 returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncBufferSubData1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferSubData1(
		this.ref,
	)
}

// FuncBufferSubData1 returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) FuncBufferSubData1() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint, length GLuint)]) {
	bindings.FuncWebGL2RenderingContextBufferSubData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferSubData1 calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData1(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
		uint32(length),
	)

	return
}

// HasFuncBufferSubData2 returns true if the method "WebGL2RenderingContext.bufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncBufferSubData2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBufferSubData2(
		this.ref,
	)
}

// FuncBufferSubData2 returns the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) FuncBufferSubData2() (fn js.Func[func(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextBufferSubData2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BufferSubData2 calls the method "WebGL2RenderingContext.bufferSubData".
func (this WebGL2RenderingContext) BufferSubData2(target GLenum, dstByteOffset GLintptr, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBufferSubData2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(dstByteOffset),
		srcData.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncTexImage2D returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage2D(
		this.ref,
	)
}

// FuncTexImage2D returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) FuncTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D1 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage2D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage2D1(
		this.ref,
	)
}

// FuncTexImage2D1 returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) FuncTexImage2D1() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D1 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D1(target GLenum, level GLint, internalformat GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(level),
		int32(internalformat),
		uint32(format),
		uint32(typ),
		source.Ref(),
	)

	return
}

// HasFuncTexSubImage2D returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage2D(
		this.ref,
	)
}

// FuncTexSubImage2D returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) FuncTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pixels js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage2D1 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage2D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage2D1(
		this.ref,
	)
}

// FuncTexSubImage2D1 returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) FuncTexSubImage2D1() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D1 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D2 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage2D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage2D2(
		this.ref,
	)
}

// FuncTexImage2D2 returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) FuncTexImage2D2() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextTexImage2D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D2 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D3 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage2D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage2D3(
		this.ref,
	)
}

// FuncTexImage2D3 returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) FuncTexImage2D3() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexImage2D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D3 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage2D4 returns true if the method "WebGL2RenderingContext.texImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage2D4() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage2D4(
		this.ref,
	)
}

// FuncTexImage2D4 returns the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) FuncTexImage2D4() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextTexImage2D4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage2D4 calls the method "WebGL2RenderingContext.texImage2D".
func (this WebGL2RenderingContext) TexImage2D4(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage2D4(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage2D2 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage2D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage2D2(
		this.ref,
	)
}

// FuncTexSubImage2D2 returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) FuncTexSubImage2D2() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage2D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D2 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage2D3 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage2D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage2D3(
		this.ref,
	)
}

// FuncTexSubImage2D3 returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) FuncTexSubImage2D3() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage2D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D3 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage2D4 returns true if the method "WebGL2RenderingContext.texSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage2D4() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage2D4(
		this.ref,
	)
}

// FuncTexSubImage2D4 returns the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) FuncTexSubImage2D4() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage2D4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage2D4 calls the method "WebGL2RenderingContext.texSubImage2D".
func (this WebGL2RenderingContext) TexSubImage2D4(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage2D4(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage2D returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage2D(
		this.ref,
	)
}

// FuncCompressedTexImage2D returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage2D calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage2D1 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage2D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage2D1(
		this.ref,
	)
}

// FuncCompressedTexImage2D1 returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexImage2D1() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage2D1 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage2D2 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage2D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage2D2(
		this.ref,
	)
}

// FuncCompressedTexImage2D2 returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexImage2D2() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage2D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage2D2 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage2D3 returns true if the method "WebGL2RenderingContext.compressedTexImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage2D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage2D3(
		this.ref,
	)
}

// FuncCompressedTexImage2D3 returns the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexImage2D3() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage2D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage2D3 calls the method "WebGL2RenderingContext.compressedTexImage2D".
func (this WebGL2RenderingContext) CompressedTexImage2D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage2D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage2D returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage2D(
		this.ref,
	)
}

// FuncCompressedTexSubImage2D returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage2D calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage2D1 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage2D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage2D1(
		this.ref,
	)
}

// FuncCompressedTexSubImage2D1 returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage2D1() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage2D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage2D1 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage2D2 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage2D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage2D2(
		this.ref,
	)
}

// FuncCompressedTexSubImage2D2 returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage2D2() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage2D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage2D2 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage2D3 returns true if the method "WebGL2RenderingContext.compressedTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage2D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage2D3(
		this.ref,
	)
}

// FuncCompressedTexSubImage2D3 returns the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage2D3() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage2D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage2D3 calls the method "WebGL2RenderingContext.compressedTexSubImage2D".
func (this WebGL2RenderingContext) CompressedTexSubImage2D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage2D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncUniform1fv returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1fv(
		this.ref,
	)
}

// FuncUniform1fv returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) FuncUniform1fv() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1fv calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform1fv1 returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1fv1(
		this.ref,
	)
}

// FuncUniform1fv1 returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) FuncUniform1fv1() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1fv1 calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform1fv2 returns true if the method "WebGL2RenderingContext.uniform1fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1fv2(
		this.ref,
	)
}

// FuncUniform1fv2 returns the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) FuncUniform1fv2() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniform1fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1fv2 calls the method "WebGL2RenderingContext.uniform1fv".
func (this WebGL2RenderingContext) Uniform1fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform2fv returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2fv(
		this.ref,
	)
}

// FuncUniform2fv returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) FuncUniform2fv() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2fv calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform2fv1 returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2fv1(
		this.ref,
	)
}

// FuncUniform2fv1 returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) FuncUniform2fv1() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2fv1 calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform2fv2 returns true if the method "WebGL2RenderingContext.uniform2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2fv2(
		this.ref,
	)
}

// FuncUniform2fv2 returns the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) FuncUniform2fv2() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniform2fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2fv2 calls the method "WebGL2RenderingContext.uniform2fv".
func (this WebGL2RenderingContext) Uniform2fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform3fv returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3fv(
		this.ref,
	)
}

// FuncUniform3fv returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) FuncUniform3fv() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3fv calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform3fv1 returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3fv1(
		this.ref,
	)
}

// FuncUniform3fv1 returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) FuncUniform3fv1() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3fv1 calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform3fv2 returns true if the method "WebGL2RenderingContext.uniform3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3fv2(
		this.ref,
	)
}

// FuncUniform3fv2 returns the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) FuncUniform3fv2() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniform3fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3fv2 calls the method "WebGL2RenderingContext.uniform3fv".
func (this WebGL2RenderingContext) Uniform3fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform4fv returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4fv(
		this.ref,
	)
}

// FuncUniform4fv returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) FuncUniform4fv() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4fv calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv(location WebGLUniformLocation, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform4fv1 returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4fv1(
		this.ref,
	)
}

// FuncUniform4fv1 returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) FuncUniform4fv1() (fn js.Func[func(location WebGLUniformLocation, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4fv1 calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv1(location WebGLUniformLocation, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform4fv2 returns true if the method "WebGL2RenderingContext.uniform4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4fv2(
		this.ref,
	)
}

// FuncUniform4fv2 returns the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) FuncUniform4fv2() (fn js.Func[func(location WebGLUniformLocation, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniform4fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4fv2 calls the method "WebGL2RenderingContext.uniform4fv".
func (this WebGL2RenderingContext) Uniform4fv2(location WebGLUniformLocation, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform1iv returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1iv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1iv(
		this.ref,
	)
}

// FuncUniform1iv returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) FuncUniform1iv() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1iv calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform1iv1 returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1iv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1iv1(
		this.ref,
	)
}

// FuncUniform1iv1 returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) FuncUniform1iv1() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1iv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1iv1 calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform1iv2 returns true if the method "WebGL2RenderingContext.uniform1iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1iv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1iv2(
		this.ref,
	)
}

// FuncUniform1iv2 returns the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) FuncUniform1iv2() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	bindings.FuncWebGL2RenderingContextUniform1iv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1iv2 calls the method "WebGL2RenderingContext.uniform1iv".
func (this WebGL2RenderingContext) Uniform1iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1iv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform2iv returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2iv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2iv(
		this.ref,
	)
}

// FuncUniform2iv returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) FuncUniform2iv() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2iv calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform2iv1 returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2iv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2iv1(
		this.ref,
	)
}

// FuncUniform2iv1 returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) FuncUniform2iv1() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2iv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2iv1 calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform2iv2 returns true if the method "WebGL2RenderingContext.uniform2iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2iv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2iv2(
		this.ref,
	)
}

// FuncUniform2iv2 returns the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) FuncUniform2iv2() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	bindings.FuncWebGL2RenderingContextUniform2iv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2iv2 calls the method "WebGL2RenderingContext.uniform2iv".
func (this WebGL2RenderingContext) Uniform2iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2iv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform3iv returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3iv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3iv(
		this.ref,
	)
}

// FuncUniform3iv returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) FuncUniform3iv() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3iv calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform3iv1 returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3iv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3iv1(
		this.ref,
	)
}

// FuncUniform3iv1 returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) FuncUniform3iv1() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3iv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3iv1 calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform3iv2 returns true if the method "WebGL2RenderingContext.uniform3iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3iv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3iv2(
		this.ref,
	)
}

// FuncUniform3iv2 returns the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) FuncUniform3iv2() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	bindings.FuncWebGL2RenderingContextUniform3iv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3iv2 calls the method "WebGL2RenderingContext.uniform3iv".
func (this WebGL2RenderingContext) Uniform3iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3iv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform4iv returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4iv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4iv(
		this.ref,
	)
}

// FuncUniform4iv returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) FuncUniform4iv() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4iv calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv(location WebGLUniformLocation, data Int32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform4iv1 returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4iv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4iv1(
		this.ref,
	)
}

// FuncUniform4iv1 returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) FuncUniform4iv1() (fn js.Func[func(location WebGLUniformLocation, data Int32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4iv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4iv1 calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv1(location WebGLUniformLocation, data Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform4iv2 returns true if the method "WebGL2RenderingContext.uniform4iv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4iv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4iv2(
		this.ref,
	)
}

// FuncUniform4iv2 returns the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) FuncUniform4iv2() (fn js.Func[func(location WebGLUniformLocation, data Int32List)]) {
	bindings.FuncWebGL2RenderingContextUniform4iv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4iv2 calls the method "WebGL2RenderingContext.uniform4iv".
func (this WebGL2RenderingContext) Uniform4iv2(location WebGLUniformLocation, data Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4iv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix2fv returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2fv(
		this.ref,
	)
}

// FuncUniformMatrix2fv returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2fv calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2fv1(
		this.ref,
	)
}

// FuncUniformMatrix2fv1 returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2fv1 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2fv2(
		this.ref,
	)
}

// FuncUniformMatrix2fv2 returns the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2fv2 calls the method "WebGL2RenderingContext.uniformMatrix2fv".
func (this WebGL2RenderingContext) UniformMatrix2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix3fv returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3fv(
		this.ref,
	)
}

// FuncUniformMatrix3fv returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3fv calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3fv1(
		this.ref,
	)
}

// FuncUniformMatrix3fv1 returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3fv1 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3fv2(
		this.ref,
	)
}

// FuncUniformMatrix3fv2 returns the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3fv2 calls the method "WebGL2RenderingContext.uniformMatrix3fv".
func (this WebGL2RenderingContext) UniformMatrix3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix4fv returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4fv(
		this.ref,
	)
}

// FuncUniformMatrix4fv returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4fv calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4fv1(
		this.ref,
	)
}

// FuncUniformMatrix4fv1 returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4fv1 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4fv2(
		this.ref,
	)
}

// FuncUniformMatrix4fv2 returns the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4fv2 calls the method "WebGL2RenderingContext.uniformMatrix4fv".
func (this WebGL2RenderingContext) UniformMatrix4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncReadPixels returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasFuncReadPixels() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextReadPixels(
		this.ref,
	)
}

// FuncReadPixels returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) FuncReadPixels() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextReadPixels(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadPixels calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncReadPixels1 returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasFuncReadPixels1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextReadPixels1(
		this.ref,
	)
}

// FuncReadPixels1 returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) FuncReadPixels1() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextReadPixels1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadPixels1 calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels1(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncReadPixels2 returns true if the method "WebGL2RenderingContext.readPixels" exists.
func (this WebGL2RenderingContext) HasFuncReadPixels2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextReadPixels2(
		this.ref,
	)
}

// FuncReadPixels2 returns the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) FuncReadPixels2() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView, dstOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextReadPixels2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadPixels2 calls the method "WebGL2RenderingContext.readPixels".
func (this WebGL2RenderingContext) ReadPixels2(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, typ GLenum, dstData js.ArrayBufferView, dstOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadPixels2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCopyBufferSubData returns true if the method "WebGL2RenderingContext.copyBufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncCopyBufferSubData() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCopyBufferSubData(
		this.ref,
	)
}

// FuncCopyBufferSubData returns the method "WebGL2RenderingContext.copyBufferSubData".
func (this WebGL2RenderingContext) FuncCopyBufferSubData() (fn js.Func[func(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr)]) {
	bindings.FuncWebGL2RenderingContextCopyBufferSubData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyBufferSubData calls the method "WebGL2RenderingContext.copyBufferSubData".
func (this WebGL2RenderingContext) CopyBufferSubData(readTarget GLenum, writeTarget GLenum, readOffset GLintptr, writeOffset GLintptr, size GLsizeiptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyBufferSubData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(readTarget),
		uint32(writeTarget),
		float64(readOffset),
		float64(writeOffset),
		float64(size),
	)

	return
}

// HasFuncGetBufferSubData returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncGetBufferSubData() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetBufferSubData(
		this.ref,
	)
}

// FuncGetBufferSubData returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) FuncGetBufferSubData() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint, length GLuint)]) {
	bindings.FuncWebGL2RenderingContextGetBufferSubData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBufferSubData calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint, length GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
		uint32(length),
	)

	return
}

// HasFuncGetBufferSubData1 returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncGetBufferSubData1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetBufferSubData1(
		this.ref,
	)
}

// FuncGetBufferSubData1 returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) FuncGetBufferSubData1() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextGetBufferSubData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBufferSubData1 calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData1(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView, dstOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
		uint32(dstOffset),
	)

	return
}

// HasFuncGetBufferSubData2 returns true if the method "WebGL2RenderingContext.getBufferSubData" exists.
func (this WebGL2RenderingContext) HasFuncGetBufferSubData2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetBufferSubData2(
		this.ref,
	)
}

// FuncGetBufferSubData2 returns the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) FuncGetBufferSubData2() (fn js.Func[func(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextGetBufferSubData2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBufferSubData2 calls the method "WebGL2RenderingContext.getBufferSubData".
func (this WebGL2RenderingContext) GetBufferSubData2(target GLenum, srcByteOffset GLintptr, dstBuffer js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGetBufferSubData2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		float64(srcByteOffset),
		dstBuffer.Ref(),
	)

	return
}

// HasFuncBlitFramebuffer returns true if the method "WebGL2RenderingContext.blitFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncBlitFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlitFramebuffer(
		this.ref,
	)
}

// FuncBlitFramebuffer returns the method "WebGL2RenderingContext.blitFramebuffer".
func (this WebGL2RenderingContext) FuncBlitFramebuffer() (fn js.Func[func(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum)]) {
	bindings.FuncWebGL2RenderingContextBlitFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlitFramebuffer calls the method "WebGL2RenderingContext.blitFramebuffer".
func (this WebGL2RenderingContext) BlitFramebuffer(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlitFramebuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncFramebufferTextureLayer returns true if the method "WebGL2RenderingContext.framebufferTextureLayer" exists.
func (this WebGL2RenderingContext) HasFuncFramebufferTextureLayer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFramebufferTextureLayer(
		this.ref,
	)
}

// FuncFramebufferTextureLayer returns the method "WebGL2RenderingContext.framebufferTextureLayer".
func (this WebGL2RenderingContext) FuncFramebufferTextureLayer() (fn js.Func[func(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint)]) {
	bindings.FuncWebGL2RenderingContextFramebufferTextureLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FramebufferTextureLayer calls the method "WebGL2RenderingContext.framebufferTextureLayer".
func (this WebGL2RenderingContext) FramebufferTextureLayer(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, layer GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferTextureLayer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(layer),
	)

	return
}

// HasFuncInvalidateFramebuffer returns true if the method "WebGL2RenderingContext.invalidateFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncInvalidateFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextInvalidateFramebuffer(
		this.ref,
	)
}

// FuncInvalidateFramebuffer returns the method "WebGL2RenderingContext.invalidateFramebuffer".
func (this WebGL2RenderingContext) FuncInvalidateFramebuffer() (fn js.Func[func(target GLenum, attachments js.Array[GLenum])]) {
	bindings.FuncWebGL2RenderingContextInvalidateFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InvalidateFramebuffer calls the method "WebGL2RenderingContext.invalidateFramebuffer".
func (this WebGL2RenderingContext) InvalidateFramebuffer(target GLenum, attachments js.Array[GLenum]) (ret js.Void) {
	bindings.CallWebGL2RenderingContextInvalidateFramebuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		attachments.Ref(),
	)

	return
}

// HasFuncInvalidateSubFramebuffer returns true if the method "WebGL2RenderingContext.invalidateSubFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncInvalidateSubFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextInvalidateSubFramebuffer(
		this.ref,
	)
}

// FuncInvalidateSubFramebuffer returns the method "WebGL2RenderingContext.invalidateSubFramebuffer".
func (this WebGL2RenderingContext) FuncInvalidateSubFramebuffer() (fn js.Func[func(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextInvalidateSubFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InvalidateSubFramebuffer calls the method "WebGL2RenderingContext.invalidateSubFramebuffer".
func (this WebGL2RenderingContext) InvalidateSubFramebuffer(target GLenum, attachments js.Array[GLenum], x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextInvalidateSubFramebuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		attachments.Ref(),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncReadBuffer returns true if the method "WebGL2RenderingContext.readBuffer" exists.
func (this WebGL2RenderingContext) HasFuncReadBuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextReadBuffer(
		this.ref,
	)
}

// FuncReadBuffer returns the method "WebGL2RenderingContext.readBuffer".
func (this WebGL2RenderingContext) FuncReadBuffer() (fn js.Func[func(src GLenum)]) {
	bindings.FuncWebGL2RenderingContextReadBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReadBuffer calls the method "WebGL2RenderingContext.readBuffer".
func (this WebGL2RenderingContext) ReadBuffer(src GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextReadBuffer(
		this.ref, js.Pointer(&ret),
		uint32(src),
	)

	return
}

// TryReadBuffer calls the method "WebGL2RenderingContext.readBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryReadBuffer(src GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextReadBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(src),
	)

	return
}

// HasFuncGetInternalformatParameter returns true if the method "WebGL2RenderingContext.getInternalformatParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetInternalformatParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetInternalformatParameter(
		this.ref,
	)
}

// FuncGetInternalformatParameter returns the method "WebGL2RenderingContext.getInternalformatParameter".
func (this WebGL2RenderingContext) FuncGetInternalformatParameter() (fn js.Func[func(target GLenum, internalformat GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetInternalformatParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetInternalformatParameter calls the method "WebGL2RenderingContext.getInternalformatParameter".
func (this WebGL2RenderingContext) GetInternalformatParameter(target GLenum, internalformat GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetInternalformatParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		uint32(pname),
	)

	return
}

// HasFuncRenderbufferStorageMultisample returns true if the method "WebGL2RenderingContext.renderbufferStorageMultisample" exists.
func (this WebGL2RenderingContext) HasFuncRenderbufferStorageMultisample() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextRenderbufferStorageMultisample(
		this.ref,
	)
}

// FuncRenderbufferStorageMultisample returns the method "WebGL2RenderingContext.renderbufferStorageMultisample".
func (this WebGL2RenderingContext) FuncRenderbufferStorageMultisample() (fn js.Func[func(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextRenderbufferStorageMultisample(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RenderbufferStorageMultisample calls the method "WebGL2RenderingContext.renderbufferStorageMultisample".
func (this WebGL2RenderingContext) RenderbufferStorageMultisample(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextRenderbufferStorageMultisample(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(samples),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncTexStorage2D returns true if the method "WebGL2RenderingContext.texStorage2D" exists.
func (this WebGL2RenderingContext) HasFuncTexStorage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexStorage2D(
		this.ref,
	)
}

// FuncTexStorage2D returns the method "WebGL2RenderingContext.texStorage2D".
func (this WebGL2RenderingContext) FuncTexStorage2D() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextTexStorage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexStorage2D calls the method "WebGL2RenderingContext.texStorage2D".
func (this WebGL2RenderingContext) TexStorage2D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexStorage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncTexStorage3D returns true if the method "WebGL2RenderingContext.texStorage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexStorage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexStorage3D(
		this.ref,
	)
}

// FuncTexStorage3D returns the method "WebGL2RenderingContext.texStorage3D".
func (this WebGL2RenderingContext) FuncTexStorage3D() (fn js.Func[func(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei)]) {
	bindings.FuncWebGL2RenderingContextTexStorage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexStorage3D calls the method "WebGL2RenderingContext.texStorage3D".
func (this WebGL2RenderingContext) TexStorage3D(target GLenum, levels GLsizei, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexStorage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		int32(levels),
		uint32(internalformat),
		int32(width),
		int32(height),
		int32(depth),
	)

	return
}

// HasFuncTexImage3D returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage3D(
		this.ref,
	)
}

// FuncTexImage3D returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) FuncTexImage3D() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextTexImage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage3D calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage3D1 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage3D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage3D1(
		this.ref,
	)
}

// FuncTexImage3D1 returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) FuncTexImage3D1() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexImage3D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage3D1 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D1(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage3D2 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage3D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage3D2(
		this.ref,
	)
}

// FuncTexImage3D2 returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) FuncTexImage3D2() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextTexImage3D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage3D2 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D2(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexImage3D3 returns true if the method "WebGL2RenderingContext.texImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexImage3D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexImage3D3(
		this.ref,
	)
}

// FuncTexImage3D3 returns the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) FuncTexImage3D3() (fn js.Func[func(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextTexImage3D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexImage3D3 calls the method "WebGL2RenderingContext.texImage3D".
func (this WebGL2RenderingContext) TexImage3D3(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, depth GLsizei, border GLint, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexImage3D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage3D returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage3D(
		this.ref,
	)
}

// FuncTexSubImage3D returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) FuncTexSubImage3D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage3D calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, pboOffset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage3D1 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage3D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage3D1(
		this.ref,
	)
}

// FuncTexSubImage3D1 returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) FuncTexSubImage3D1() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage3D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage3D1 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, source TexImageSource) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage3D2 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage3D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage3D2(
		this.ref,
	)
}

// FuncTexSubImage3D2 returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) FuncTexSubImage3D2() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage3D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage3D2 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncTexSubImage3D3 returns true if the method "WebGL2RenderingContext.texSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncTexSubImage3D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexSubImage3D3(
		this.ref,
	)
}

// FuncTexSubImage3D3 returns the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) FuncTexSubImage3D3() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextTexSubImage3D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexSubImage3D3 calls the method "WebGL2RenderingContext.texSubImage3D".
func (this WebGL2RenderingContext) TexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, typ GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexSubImage3D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCopyTexSubImage3D returns true if the method "WebGL2RenderingContext.copyTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCopyTexSubImage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCopyTexSubImage3D(
		this.ref,
	)
}

// FuncCopyTexSubImage3D returns the method "WebGL2RenderingContext.copyTexSubImage3D".
func (this WebGL2RenderingContext) FuncCopyTexSubImage3D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextCopyTexSubImage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTexSubImage3D calls the method "WebGL2RenderingContext.copyTexSubImage3D".
func (this WebGL2RenderingContext) CopyTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexSubImage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage3D returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage3D(
		this.ref,
	)
}

// FuncCompressedTexImage3D returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexImage3D() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage3D calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage3D1 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage3D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage3D1(
		this.ref,
	)
}

// FuncCompressedTexImage3D1 returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexImage3D1() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage3D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage3D1 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D1(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage3D2 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage3D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage3D2(
		this.ref,
	)
}

// FuncCompressedTexImage3D2 returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexImage3D2() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage3D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage3D2 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D2(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexImage3D3 returns true if the method "WebGL2RenderingContext.compressedTexImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexImage3D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexImage3D3(
		this.ref,
	)
}

// FuncCompressedTexImage3D3 returns the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexImage3D3() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexImage3D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexImage3D3 calls the method "WebGL2RenderingContext.compressedTexImage3D".
func (this WebGL2RenderingContext) CompressedTexImage3D3(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexImage3D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage3D returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage3D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage3D(
		this.ref,
	)
}

// FuncCompressedTexSubImage3D returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage3D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage3D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage3D calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage3D1 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage3D1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage3D1(
		this.ref,
	)
}

// FuncCompressedTexSubImage3D1 returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage3D1() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage3D1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage3D1 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D1(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint, srcLengthOverride GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage3D2 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage3D2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage3D2(
		this.ref,
	)
}

// FuncCompressedTexSubImage3D2 returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage3D2() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage3D2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage3D2 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D2(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCompressedTexSubImage3D3 returns true if the method "WebGL2RenderingContext.compressedTexSubImage3D" exists.
func (this WebGL2RenderingContext) HasFuncCompressedTexSubImage3D3() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompressedTexSubImage3D3(
		this.ref,
	)
}

// FuncCompressedTexSubImage3D3 returns the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) FuncCompressedTexSubImage3D3() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView)]) {
	bindings.FuncWebGL2RenderingContextCompressedTexSubImage3D3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompressedTexSubImage3D3 calls the method "WebGL2RenderingContext.compressedTexSubImage3D".
func (this WebGL2RenderingContext) CompressedTexSubImage3D3(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, srcData js.ArrayBufferView) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompressedTexSubImage3D3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncGetFragDataLocation returns true if the method "WebGL2RenderingContext.getFragDataLocation" exists.
func (this WebGL2RenderingContext) HasFuncGetFragDataLocation() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetFragDataLocation(
		this.ref,
	)
}

// FuncGetFragDataLocation returns the method "WebGL2RenderingContext.getFragDataLocation".
func (this WebGL2RenderingContext) FuncGetFragDataLocation() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	bindings.FuncWebGL2RenderingContextGetFragDataLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFragDataLocation calls the method "WebGL2RenderingContext.getFragDataLocation".
func (this WebGL2RenderingContext) GetFragDataLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGL2RenderingContextGetFragDataLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncUniform1ui returns true if the method "WebGL2RenderingContext.uniform1ui" exists.
func (this WebGL2RenderingContext) HasFuncUniform1ui() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1ui(
		this.ref,
	)
}

// FuncUniform1ui returns the method "WebGL2RenderingContext.uniform1ui".
func (this WebGL2RenderingContext) FuncUniform1ui() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1ui(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1ui calls the method "WebGL2RenderingContext.uniform1ui".
func (this WebGL2RenderingContext) Uniform1ui(location WebGLUniformLocation, v0 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1ui(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
	)

	return
}

// HasFuncUniform2ui returns true if the method "WebGL2RenderingContext.uniform2ui" exists.
func (this WebGL2RenderingContext) HasFuncUniform2ui() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2ui(
		this.ref,
	)
}

// FuncUniform2ui returns the method "WebGL2RenderingContext.uniform2ui".
func (this WebGL2RenderingContext) FuncUniform2ui() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2ui(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2ui calls the method "WebGL2RenderingContext.uniform2ui".
func (this WebGL2RenderingContext) Uniform2ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2ui(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
	)

	return
}

// HasFuncUniform3ui returns true if the method "WebGL2RenderingContext.uniform3ui" exists.
func (this WebGL2RenderingContext) HasFuncUniform3ui() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3ui(
		this.ref,
	)
}

// FuncUniform3ui returns the method "WebGL2RenderingContext.uniform3ui".
func (this WebGL2RenderingContext) FuncUniform3ui() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3ui(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3ui calls the method "WebGL2RenderingContext.uniform3ui".
func (this WebGL2RenderingContext) Uniform3ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3ui(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
	)

	return
}

// HasFuncUniform4ui returns true if the method "WebGL2RenderingContext.uniform4ui" exists.
func (this WebGL2RenderingContext) HasFuncUniform4ui() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4ui(
		this.ref,
	)
}

// FuncUniform4ui returns the method "WebGL2RenderingContext.uniform4ui".
func (this WebGL2RenderingContext) FuncUniform4ui() (fn js.Func[func(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4ui(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4ui calls the method "WebGL2RenderingContext.uniform4ui".
func (this WebGL2RenderingContext) Uniform4ui(location WebGLUniformLocation, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4ui(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		uint32(v0),
		uint32(v1),
		uint32(v2),
		uint32(v3),
	)

	return
}

// HasFuncUniform1uiv returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1uiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1uiv(
		this.ref,
	)
}

// FuncUniform1uiv returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) FuncUniform1uiv() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1uiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1uiv calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform1uiv1 returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1uiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1uiv1(
		this.ref,
	)
}

// FuncUniform1uiv1 returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) FuncUniform1uiv1() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform1uiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1uiv1 calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform1uiv2 returns true if the method "WebGL2RenderingContext.uniform1uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform1uiv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1uiv2(
		this.ref,
	)
}

// FuncUniform1uiv2 returns the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) FuncUniform1uiv2() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	bindings.FuncWebGL2RenderingContextUniform1uiv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1uiv2 calls the method "WebGL2RenderingContext.uniform1uiv".
func (this WebGL2RenderingContext) Uniform1uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1uiv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform2uiv returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2uiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2uiv(
		this.ref,
	)
}

// FuncUniform2uiv returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) FuncUniform2uiv() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2uiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2uiv calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform2uiv1 returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2uiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2uiv1(
		this.ref,
	)
}

// FuncUniform2uiv1 returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) FuncUniform2uiv1() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform2uiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2uiv1 calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform2uiv2 returns true if the method "WebGL2RenderingContext.uniform2uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform2uiv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2uiv2(
		this.ref,
	)
}

// FuncUniform2uiv2 returns the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) FuncUniform2uiv2() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	bindings.FuncWebGL2RenderingContextUniform2uiv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2uiv2 calls the method "WebGL2RenderingContext.uniform2uiv".
func (this WebGL2RenderingContext) Uniform2uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2uiv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform3uiv returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3uiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3uiv(
		this.ref,
	)
}

// FuncUniform3uiv returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) FuncUniform3uiv() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3uiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3uiv calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform3uiv1 returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3uiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3uiv1(
		this.ref,
	)
}

// FuncUniform3uiv1 returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) FuncUniform3uiv1() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform3uiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3uiv1 calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform3uiv2 returns true if the method "WebGL2RenderingContext.uniform3uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform3uiv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3uiv2(
		this.ref,
	)
}

// FuncUniform3uiv2 returns the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) FuncUniform3uiv2() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	bindings.FuncWebGL2RenderingContextUniform3uiv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3uiv2 calls the method "WebGL2RenderingContext.uniform3uiv".
func (this WebGL2RenderingContext) Uniform3uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3uiv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniform4uiv returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4uiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4uiv(
		this.ref,
	)
}

// FuncUniform4uiv returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) FuncUniform4uiv() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4uiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4uiv calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv(location WebGLUniformLocation, data Uint32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniform4uiv1 returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4uiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4uiv1(
		this.ref,
	)
}

// FuncUniform4uiv1 returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) FuncUniform4uiv1() (fn js.Func[func(location WebGLUniformLocation, data Uint32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniform4uiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4uiv1 calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv1(location WebGLUniformLocation, data Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniform4uiv2 returns true if the method "WebGL2RenderingContext.uniform4uiv" exists.
func (this WebGL2RenderingContext) HasFuncUniform4uiv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4uiv2(
		this.ref,
	)
}

// FuncUniform4uiv2 returns the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) FuncUniform4uiv2() (fn js.Func[func(location WebGLUniformLocation, data Uint32List)]) {
	bindings.FuncWebGL2RenderingContextUniform4uiv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4uiv2 calls the method "WebGL2RenderingContext.uniform4uiv".
func (this WebGL2RenderingContext) Uniform4uiv2(location WebGLUniformLocation, data Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4uiv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix3x2fv returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x2fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x2fv(
		this.ref,
	)
}

// FuncUniformMatrix3x2fv returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x2fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x2fv calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix3x2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x2fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x2fv1(
		this.ref,
	)
}

// FuncUniformMatrix3x2fv1 returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x2fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x2fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix3x2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x2fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x2fv2(
		this.ref,
	)
}

// FuncUniformMatrix3x2fv2 returns the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x2fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x2fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x2fv".
func (this WebGL2RenderingContext) UniformMatrix3x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x2fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix4x2fv returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x2fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x2fv(
		this.ref,
	)
}

// FuncUniformMatrix4x2fv returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x2fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x2fv calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix4x2fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x2fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x2fv1(
		this.ref,
	)
}

// FuncUniformMatrix4x2fv1 returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x2fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x2fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x2fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix4x2fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4x2fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x2fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x2fv2(
		this.ref,
	)
}

// FuncUniformMatrix4x2fv2 returns the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x2fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x2fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x2fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x2fv".
func (this WebGL2RenderingContext) UniformMatrix4x2fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x2fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix2x3fv returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x3fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x3fv(
		this.ref,
	)
}

// FuncUniformMatrix2x3fv returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x3fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x3fv calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix2x3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x3fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x3fv1(
		this.ref,
	)
}

// FuncUniformMatrix2x3fv1 returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x3fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x3fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix2x3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x3fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x3fv2(
		this.ref,
	)
}

// FuncUniformMatrix2x3fv2 returns the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x3fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x3fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x3fv".
func (this WebGL2RenderingContext) UniformMatrix2x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x3fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix4x3fv returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x3fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x3fv(
		this.ref,
	)
}

// FuncUniformMatrix4x3fv returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x3fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x3fv calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix4x3fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x3fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x3fv1(
		this.ref,
	)
}

// FuncUniformMatrix4x3fv1 returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x3fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x3fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x3fv1 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix4x3fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix4x3fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix4x3fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix4x3fv2(
		this.ref,
	)
}

// FuncUniformMatrix4x3fv2 returns the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) FuncUniformMatrix4x3fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix4x3fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix4x3fv2 calls the method "WebGL2RenderingContext.uniformMatrix4x3fv".
func (this WebGL2RenderingContext) UniformMatrix4x3fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix4x3fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix2x4fv returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x4fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x4fv(
		this.ref,
	)
}

// FuncUniformMatrix2x4fv returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x4fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x4fv calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix2x4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x4fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x4fv1(
		this.ref,
	)
}

// FuncUniformMatrix2x4fv1 returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x4fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x4fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix2x4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix2x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix2x4fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix2x4fv2(
		this.ref,
	)
}

// FuncUniformMatrix2x4fv2 returns the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix2x4fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix2x4fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix2x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix2x4fv".
func (this WebGL2RenderingContext) UniformMatrix2x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix2x4fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncUniformMatrix3x4fv returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x4fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x4fv(
		this.ref,
	)
}

// FuncUniformMatrix3x4fv returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x4fv() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x4fv calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint, srcLength GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
		uint32(srcLength),
	)

	return
}

// HasFuncUniformMatrix3x4fv1 returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x4fv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x4fv1(
		this.ref,
	)
}

// FuncUniformMatrix3x4fv1 returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x4fv1() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x4fv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x4fv1 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv1(location WebGLUniformLocation, transpose GLboolean, data Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncUniformMatrix3x4fv2 returns true if the method "WebGL2RenderingContext.uniformMatrix3x4fv" exists.
func (this WebGL2RenderingContext) HasFuncUniformMatrix3x4fv2() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformMatrix3x4fv2(
		this.ref,
	)
}

// FuncUniformMatrix3x4fv2 returns the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) FuncUniformMatrix3x4fv2() (fn js.Func[func(location WebGLUniformLocation, transpose GLboolean, data Float32List)]) {
	bindings.FuncWebGL2RenderingContextUniformMatrix3x4fv2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformMatrix3x4fv2 calls the method "WebGL2RenderingContext.uniformMatrix3x4fv".
func (this WebGL2RenderingContext) UniformMatrix3x4fv2(location WebGLUniformLocation, transpose GLboolean, data Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformMatrix3x4fv2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		js.Bool(bool(transpose)),
		data.Ref(),
	)

	return
}

// HasFuncVertexAttribI4i returns true if the method "WebGL2RenderingContext.vertexAttribI4i" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribI4i() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribI4i(
		this.ref,
	)
}

// FuncVertexAttribI4i returns the method "WebGL2RenderingContext.vertexAttribI4i".
func (this WebGL2RenderingContext) FuncVertexAttribI4i() (fn js.Func[func(index GLuint, x GLint, y GLint, z GLint, w GLint)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribI4i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribI4i calls the method "WebGL2RenderingContext.vertexAttribI4i".
func (this WebGL2RenderingContext) VertexAttribI4i(index GLuint, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasFuncVertexAttribI4iv returns true if the method "WebGL2RenderingContext.vertexAttribI4iv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribI4iv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribI4iv(
		this.ref,
	)
}

// FuncVertexAttribI4iv returns the method "WebGL2RenderingContext.vertexAttribI4iv".
func (this WebGL2RenderingContext) FuncVertexAttribI4iv() (fn js.Func[func(index GLuint, values Int32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribI4iv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribI4iv calls the method "WebGL2RenderingContext.vertexAttribI4iv".
func (this WebGL2RenderingContext) VertexAttribI4iv(index GLuint, values Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4iv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttribI4ui returns true if the method "WebGL2RenderingContext.vertexAttribI4ui" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribI4ui() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribI4ui(
		this.ref,
	)
}

// FuncVertexAttribI4ui returns the method "WebGL2RenderingContext.vertexAttribI4ui".
func (this WebGL2RenderingContext) FuncVertexAttribI4ui() (fn js.Func[func(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribI4ui(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribI4ui calls the method "WebGL2RenderingContext.vertexAttribI4ui".
func (this WebGL2RenderingContext) VertexAttribI4ui(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4ui(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(x),
		uint32(y),
		uint32(z),
		uint32(w),
	)

	return
}

// HasFuncVertexAttribI4uiv returns true if the method "WebGL2RenderingContext.vertexAttribI4uiv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribI4uiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribI4uiv(
		this.ref,
	)
}

// FuncVertexAttribI4uiv returns the method "WebGL2RenderingContext.vertexAttribI4uiv".
func (this WebGL2RenderingContext) FuncVertexAttribI4uiv() (fn js.Func[func(index GLuint, values Uint32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribI4uiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribI4uiv calls the method "WebGL2RenderingContext.vertexAttribI4uiv".
func (this WebGL2RenderingContext) VertexAttribI4uiv(index GLuint, values Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribI4uiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttribIPointer returns true if the method "WebGL2RenderingContext.vertexAttribIPointer" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribIPointer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribIPointer(
		this.ref,
	)
}

// FuncVertexAttribIPointer returns the method "WebGL2RenderingContext.vertexAttribIPointer".
func (this WebGL2RenderingContext) FuncVertexAttribIPointer() (fn js.Func[func(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribIPointer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribIPointer calls the method "WebGL2RenderingContext.vertexAttribIPointer".
func (this WebGL2RenderingContext) VertexAttribIPointer(index GLuint, size GLint, typ GLenum, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribIPointer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		int32(stride),
		float64(offset),
	)

	return
}

// HasFuncVertexAttribDivisor returns true if the method "WebGL2RenderingContext.vertexAttribDivisor" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribDivisor() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribDivisor(
		this.ref,
	)
}

// FuncVertexAttribDivisor returns the method "WebGL2RenderingContext.vertexAttribDivisor".
func (this WebGL2RenderingContext) FuncVertexAttribDivisor() (fn js.Func[func(index GLuint, divisor GLuint)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribDivisor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribDivisor calls the method "WebGL2RenderingContext.vertexAttribDivisor".
func (this WebGL2RenderingContext) VertexAttribDivisor(index GLuint, divisor GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribDivisor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(divisor),
	)

	return
}

// HasFuncDrawArraysInstanced returns true if the method "WebGL2RenderingContext.drawArraysInstanced" exists.
func (this WebGL2RenderingContext) HasFuncDrawArraysInstanced() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawArraysInstanced(
		this.ref,
	)
}

// FuncDrawArraysInstanced returns the method "WebGL2RenderingContext.drawArraysInstanced".
func (this WebGL2RenderingContext) FuncDrawArraysInstanced() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei)]) {
	bindings.FuncWebGL2RenderingContextDrawArraysInstanced(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawArraysInstanced calls the method "WebGL2RenderingContext.drawArraysInstanced".
func (this WebGL2RenderingContext) DrawArraysInstanced(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawArraysInstanced(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
	)

	return
}

// HasFuncDrawElementsInstanced returns true if the method "WebGL2RenderingContext.drawElementsInstanced" exists.
func (this WebGL2RenderingContext) HasFuncDrawElementsInstanced() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawElementsInstanced(
		this.ref,
	)
}

// FuncDrawElementsInstanced returns the method "WebGL2RenderingContext.drawElementsInstanced".
func (this WebGL2RenderingContext) FuncDrawElementsInstanced() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei)]) {
	bindings.FuncWebGL2RenderingContextDrawElementsInstanced(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawElementsInstanced calls the method "WebGL2RenderingContext.drawElementsInstanced".
func (this WebGL2RenderingContext) DrawElementsInstanced(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawElementsInstanced(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
	)

	return
}

// HasFuncDrawRangeElements returns true if the method "WebGL2RenderingContext.drawRangeElements" exists.
func (this WebGL2RenderingContext) HasFuncDrawRangeElements() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawRangeElements(
		this.ref,
	)
}

// FuncDrawRangeElements returns the method "WebGL2RenderingContext.drawRangeElements".
func (this WebGL2RenderingContext) FuncDrawRangeElements() (fn js.Func[func(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextDrawRangeElements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawRangeElements calls the method "WebGL2RenderingContext.drawRangeElements".
func (this WebGL2RenderingContext) DrawRangeElements(mode GLenum, start GLuint, end GLuint, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawRangeElements(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		uint32(start),
		uint32(end),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasFuncDrawBuffers returns true if the method "WebGL2RenderingContext.drawBuffers" exists.
func (this WebGL2RenderingContext) HasFuncDrawBuffers() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawBuffers(
		this.ref,
	)
}

// FuncDrawBuffers returns the method "WebGL2RenderingContext.drawBuffers".
func (this WebGL2RenderingContext) FuncDrawBuffers() (fn js.Func[func(buffers js.Array[GLenum])]) {
	bindings.FuncWebGL2RenderingContextDrawBuffers(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawBuffers calls the method "WebGL2RenderingContext.drawBuffers".
func (this WebGL2RenderingContext) DrawBuffers(buffers js.Array[GLenum]) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawBuffers(
		this.ref, js.Pointer(&ret),
		buffers.Ref(),
	)

	return
}

// TryDrawBuffers calls the method "WebGL2RenderingContext.drawBuffers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDrawBuffers(buffers js.Array[GLenum]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDrawBuffers(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffers.Ref(),
	)

	return
}

// HasFuncClearBufferfv returns true if the method "WebGL2RenderingContext.clearBufferfv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferfv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferfv(
		this.ref,
	)
}

// FuncClearBufferfv returns the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) FuncClearBufferfv() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextClearBufferfv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferfv calls the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfv(buffer GLenum, drawbuffer GLint, values Float32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncClearBufferfv1 returns true if the method "WebGL2RenderingContext.clearBufferfv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferfv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferfv1(
		this.ref,
	)
}

// FuncClearBufferfv1 returns the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) FuncClearBufferfv1() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Float32List)]) {
	bindings.FuncWebGL2RenderingContextClearBufferfv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferfv1 calls the method "WebGL2RenderingContext.clearBufferfv".
func (this WebGL2RenderingContext) ClearBufferfv1(buffer GLenum, drawbuffer GLint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasFuncClearBufferiv returns true if the method "WebGL2RenderingContext.clearBufferiv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferiv(
		this.ref,
	)
}

// FuncClearBufferiv returns the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) FuncClearBufferiv() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextClearBufferiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferiv calls the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferiv(buffer GLenum, drawbuffer GLint, values Int32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncClearBufferiv1 returns true if the method "WebGL2RenderingContext.clearBufferiv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferiv1(
		this.ref,
	)
}

// FuncClearBufferiv1 returns the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) FuncClearBufferiv1() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Int32List)]) {
	bindings.FuncWebGL2RenderingContextClearBufferiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferiv1 calls the method "WebGL2RenderingContext.clearBufferiv".
func (this WebGL2RenderingContext) ClearBufferiv1(buffer GLenum, drawbuffer GLint, values Int32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasFuncClearBufferuiv returns true if the method "WebGL2RenderingContext.clearBufferuiv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferuiv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferuiv(
		this.ref,
	)
}

// FuncClearBufferuiv returns the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) FuncClearBufferuiv() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint)]) {
	bindings.FuncWebGL2RenderingContextClearBufferuiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferuiv calls the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuiv(buffer GLenum, drawbuffer GLint, values Uint32List, srcOffset GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferuiv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
		uint32(srcOffset),
	)

	return
}

// HasFuncClearBufferuiv1 returns true if the method "WebGL2RenderingContext.clearBufferuiv" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferuiv1() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferuiv1(
		this.ref,
	)
}

// FuncClearBufferuiv1 returns the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) FuncClearBufferuiv1() (fn js.Func[func(buffer GLenum, drawbuffer GLint, values Uint32List)]) {
	bindings.FuncWebGL2RenderingContextClearBufferuiv1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferuiv1 calls the method "WebGL2RenderingContext.clearBufferuiv".
func (this WebGL2RenderingContext) ClearBufferuiv1(buffer GLenum, drawbuffer GLint, values Uint32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferuiv1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		values.Ref(),
	)

	return
}

// HasFuncClearBufferfi returns true if the method "WebGL2RenderingContext.clearBufferfi" exists.
func (this WebGL2RenderingContext) HasFuncClearBufferfi() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearBufferfi(
		this.ref,
	)
}

// FuncClearBufferfi returns the method "WebGL2RenderingContext.clearBufferfi".
func (this WebGL2RenderingContext) FuncClearBufferfi() (fn js.Func[func(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint)]) {
	bindings.FuncWebGL2RenderingContextClearBufferfi(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBufferfi calls the method "WebGL2RenderingContext.clearBufferfi".
func (this WebGL2RenderingContext) ClearBufferfi(buffer GLenum, drawbuffer GLint, depth GLfloat, stencil GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearBufferfi(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(buffer),
		int32(drawbuffer),
		float32(depth),
		int32(stencil),
	)

	return
}

// HasFuncCreateQuery returns true if the method "WebGL2RenderingContext.createQuery" exists.
func (this WebGL2RenderingContext) HasFuncCreateQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateQuery(
		this.ref,
	)
}

// FuncCreateQuery returns the method "WebGL2RenderingContext.createQuery".
func (this WebGL2RenderingContext) FuncCreateQuery() (fn js.Func[func() WebGLQuery]) {
	bindings.FuncWebGL2RenderingContextCreateQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuery calls the method "WebGL2RenderingContext.createQuery".
func (this WebGL2RenderingContext) CreateQuery() (ret WebGLQuery) {
	bindings.CallWebGL2RenderingContextCreateQuery(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateQuery calls the method "WebGL2RenderingContext.createQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateQuery() (ret WebGLQuery, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteQuery returns true if the method "WebGL2RenderingContext.deleteQuery" exists.
func (this WebGL2RenderingContext) HasFuncDeleteQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteQuery(
		this.ref,
	)
}

// FuncDeleteQuery returns the method "WebGL2RenderingContext.deleteQuery".
func (this WebGL2RenderingContext) FuncDeleteQuery() (fn js.Func[func(query WebGLQuery)]) {
	bindings.FuncWebGL2RenderingContextDeleteQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteQuery calls the method "WebGL2RenderingContext.deleteQuery".
func (this WebGL2RenderingContext) DeleteQuery(query WebGLQuery) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteQuery(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryDeleteQuery calls the method "WebGL2RenderingContext.deleteQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteQuery(query WebGLQuery) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncIsQuery returns true if the method "WebGL2RenderingContext.isQuery" exists.
func (this WebGL2RenderingContext) HasFuncIsQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsQuery(
		this.ref,
	)
}

// FuncIsQuery returns the method "WebGL2RenderingContext.isQuery".
func (this WebGL2RenderingContext) FuncIsQuery() (fn js.Func[func(query WebGLQuery) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsQuery calls the method "WebGL2RenderingContext.isQuery".
func (this WebGL2RenderingContext) IsQuery(query WebGLQuery) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsQuery(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryIsQuery calls the method "WebGL2RenderingContext.isQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsQuery(query WebGLQuery) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncBeginQuery returns true if the method "WebGL2RenderingContext.beginQuery" exists.
func (this WebGL2RenderingContext) HasFuncBeginQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBeginQuery(
		this.ref,
	)
}

// FuncBeginQuery returns the method "WebGL2RenderingContext.beginQuery".
func (this WebGL2RenderingContext) FuncBeginQuery() (fn js.Func[func(target GLenum, query WebGLQuery)]) {
	bindings.FuncWebGL2RenderingContextBeginQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginQuery calls the method "WebGL2RenderingContext.beginQuery".
func (this WebGL2RenderingContext) BeginQuery(target GLenum, query WebGLQuery) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBeginQuery(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		query.Ref(),
	)

	return
}

// HasFuncEndQuery returns true if the method "WebGL2RenderingContext.endQuery" exists.
func (this WebGL2RenderingContext) HasFuncEndQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextEndQuery(
		this.ref,
	)
}

// FuncEndQuery returns the method "WebGL2RenderingContext.endQuery".
func (this WebGL2RenderingContext) FuncEndQuery() (fn js.Func[func(target GLenum)]) {
	bindings.FuncWebGL2RenderingContextEndQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndQuery calls the method "WebGL2RenderingContext.endQuery".
func (this WebGL2RenderingContext) EndQuery(target GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEndQuery(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryEndQuery calls the method "WebGL2RenderingContext.endQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEndQuery(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEndQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncGetQuery returns true if the method "WebGL2RenderingContext.getQuery" exists.
func (this WebGL2RenderingContext) HasFuncGetQuery() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetQuery(
		this.ref,
	)
}

// FuncGetQuery returns the method "WebGL2RenderingContext.getQuery".
func (this WebGL2RenderingContext) FuncGetQuery() (fn js.Func[func(target GLenum, pname GLenum) WebGLQuery]) {
	bindings.FuncWebGL2RenderingContextGetQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetQuery calls the method "WebGL2RenderingContext.getQuery".
func (this WebGL2RenderingContext) GetQuery(target GLenum, pname GLenum) (ret WebGLQuery) {
	bindings.CallWebGL2RenderingContextGetQuery(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetQueryParameter returns true if the method "WebGL2RenderingContext.getQueryParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetQueryParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetQueryParameter(
		this.ref,
	)
}

// FuncGetQueryParameter returns the method "WebGL2RenderingContext.getQueryParameter".
func (this WebGL2RenderingContext) FuncGetQueryParameter() (fn js.Func[func(query WebGLQuery, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetQueryParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetQueryParameter calls the method "WebGL2RenderingContext.getQueryParameter".
func (this WebGL2RenderingContext) GetQueryParameter(query WebGLQuery, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetQueryParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncCreateSampler returns true if the method "WebGL2RenderingContext.createSampler" exists.
func (this WebGL2RenderingContext) HasFuncCreateSampler() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateSampler(
		this.ref,
	)
}

// FuncCreateSampler returns the method "WebGL2RenderingContext.createSampler".
func (this WebGL2RenderingContext) FuncCreateSampler() (fn js.Func[func() WebGLSampler]) {
	bindings.FuncWebGL2RenderingContextCreateSampler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSampler calls the method "WebGL2RenderingContext.createSampler".
func (this WebGL2RenderingContext) CreateSampler() (ret WebGLSampler) {
	bindings.CallWebGL2RenderingContextCreateSampler(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSampler calls the method "WebGL2RenderingContext.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateSampler() (ret WebGLSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateSampler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteSampler returns true if the method "WebGL2RenderingContext.deleteSampler" exists.
func (this WebGL2RenderingContext) HasFuncDeleteSampler() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteSampler(
		this.ref,
	)
}

// FuncDeleteSampler returns the method "WebGL2RenderingContext.deleteSampler".
func (this WebGL2RenderingContext) FuncDeleteSampler() (fn js.Func[func(sampler WebGLSampler)]) {
	bindings.FuncWebGL2RenderingContextDeleteSampler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteSampler calls the method "WebGL2RenderingContext.deleteSampler".
func (this WebGL2RenderingContext) DeleteSampler(sampler WebGLSampler) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteSampler(
		this.ref, js.Pointer(&ret),
		sampler.Ref(),
	)

	return
}

// TryDeleteSampler calls the method "WebGL2RenderingContext.deleteSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteSampler(sampler WebGLSampler) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteSampler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
	)

	return
}

// HasFuncIsSampler returns true if the method "WebGL2RenderingContext.isSampler" exists.
func (this WebGL2RenderingContext) HasFuncIsSampler() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsSampler(
		this.ref,
	)
}

// FuncIsSampler returns the method "WebGL2RenderingContext.isSampler".
func (this WebGL2RenderingContext) FuncIsSampler() (fn js.Func[func(sampler WebGLSampler) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsSampler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsSampler calls the method "WebGL2RenderingContext.isSampler".
func (this WebGL2RenderingContext) IsSampler(sampler WebGLSampler) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsSampler(
		this.ref, js.Pointer(&ret),
		sampler.Ref(),
	)

	return
}

// TryIsSampler calls the method "WebGL2RenderingContext.isSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsSampler(sampler WebGLSampler) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsSampler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
	)

	return
}

// HasFuncBindSampler returns true if the method "WebGL2RenderingContext.bindSampler" exists.
func (this WebGL2RenderingContext) HasFuncBindSampler() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindSampler(
		this.ref,
	)
}

// FuncBindSampler returns the method "WebGL2RenderingContext.bindSampler".
func (this WebGL2RenderingContext) FuncBindSampler() (fn js.Func[func(unit GLuint, sampler WebGLSampler)]) {
	bindings.FuncWebGL2RenderingContextBindSampler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindSampler calls the method "WebGL2RenderingContext.bindSampler".
func (this WebGL2RenderingContext) BindSampler(unit GLuint, sampler WebGLSampler) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindSampler(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(unit),
		sampler.Ref(),
	)

	return
}

// HasFuncSamplerParameteri returns true if the method "WebGL2RenderingContext.samplerParameteri" exists.
func (this WebGL2RenderingContext) HasFuncSamplerParameteri() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextSamplerParameteri(
		this.ref,
	)
}

// FuncSamplerParameteri returns the method "WebGL2RenderingContext.samplerParameteri".
func (this WebGL2RenderingContext) FuncSamplerParameteri() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLint)]) {
	bindings.FuncWebGL2RenderingContextSamplerParameteri(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SamplerParameteri calls the method "WebGL2RenderingContext.samplerParameteri".
func (this WebGL2RenderingContext) SamplerParameteri(sampler WebGLSampler, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSamplerParameteri(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
		int32(param),
	)

	return
}

// HasFuncSamplerParameterf returns true if the method "WebGL2RenderingContext.samplerParameterf" exists.
func (this WebGL2RenderingContext) HasFuncSamplerParameterf() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextSamplerParameterf(
		this.ref,
	)
}

// FuncSamplerParameterf returns the method "WebGL2RenderingContext.samplerParameterf".
func (this WebGL2RenderingContext) FuncSamplerParameterf() (fn js.Func[func(sampler WebGLSampler, pname GLenum, param GLfloat)]) {
	bindings.FuncWebGL2RenderingContextSamplerParameterf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SamplerParameterf calls the method "WebGL2RenderingContext.samplerParameterf".
func (this WebGL2RenderingContext) SamplerParameterf(sampler WebGLSampler, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSamplerParameterf(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
		float32(param),
	)

	return
}

// HasFuncGetSamplerParameter returns true if the method "WebGL2RenderingContext.getSamplerParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetSamplerParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetSamplerParameter(
		this.ref,
	)
}

// FuncGetSamplerParameter returns the method "WebGL2RenderingContext.getSamplerParameter".
func (this WebGL2RenderingContext) FuncGetSamplerParameter() (fn js.Func[func(sampler WebGLSampler, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetSamplerParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSamplerParameter calls the method "WebGL2RenderingContext.getSamplerParameter".
func (this WebGL2RenderingContext) GetSamplerParameter(sampler WebGLSampler, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetSamplerParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sampler.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncFenceSync returns true if the method "WebGL2RenderingContext.fenceSync" exists.
func (this WebGL2RenderingContext) HasFuncFenceSync() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFenceSync(
		this.ref,
	)
}

// FuncFenceSync returns the method "WebGL2RenderingContext.fenceSync".
func (this WebGL2RenderingContext) FuncFenceSync() (fn js.Func[func(condition GLenum, flags GLbitfield) WebGLSync]) {
	bindings.FuncWebGL2RenderingContextFenceSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FenceSync calls the method "WebGL2RenderingContext.fenceSync".
func (this WebGL2RenderingContext) FenceSync(condition GLenum, flags GLbitfield) (ret WebGLSync) {
	bindings.CallWebGL2RenderingContextFenceSync(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(condition),
		uint32(flags),
	)

	return
}

// HasFuncIsSync returns true if the method "WebGL2RenderingContext.isSync" exists.
func (this WebGL2RenderingContext) HasFuncIsSync() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsSync(
		this.ref,
	)
}

// FuncIsSync returns the method "WebGL2RenderingContext.isSync".
func (this WebGL2RenderingContext) FuncIsSync() (fn js.Func[func(sync WebGLSync) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsSync calls the method "WebGL2RenderingContext.isSync".
func (this WebGL2RenderingContext) IsSync(sync WebGLSync) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsSync(
		this.ref, js.Pointer(&ret),
		sync.Ref(),
	)

	return
}

// TryIsSync calls the method "WebGL2RenderingContext.isSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsSync(sync WebGLSync) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
	)

	return
}

// HasFuncDeleteSync returns true if the method "WebGL2RenderingContext.deleteSync" exists.
func (this WebGL2RenderingContext) HasFuncDeleteSync() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteSync(
		this.ref,
	)
}

// FuncDeleteSync returns the method "WebGL2RenderingContext.deleteSync".
func (this WebGL2RenderingContext) FuncDeleteSync() (fn js.Func[func(sync WebGLSync)]) {
	bindings.FuncWebGL2RenderingContextDeleteSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteSync calls the method "WebGL2RenderingContext.deleteSync".
func (this WebGL2RenderingContext) DeleteSync(sync WebGLSync) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteSync(
		this.ref, js.Pointer(&ret),
		sync.Ref(),
	)

	return
}

// TryDeleteSync calls the method "WebGL2RenderingContext.deleteSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteSync(sync WebGLSync) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
	)

	return
}

// HasFuncClientWaitSync returns true if the method "WebGL2RenderingContext.clientWaitSync" exists.
func (this WebGL2RenderingContext) HasFuncClientWaitSync() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClientWaitSync(
		this.ref,
	)
}

// FuncClientWaitSync returns the method "WebGL2RenderingContext.clientWaitSync".
func (this WebGL2RenderingContext) FuncClientWaitSync() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLuint64) GLenum]) {
	bindings.FuncWebGL2RenderingContextClientWaitSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClientWaitSync calls the method "WebGL2RenderingContext.clientWaitSync".
func (this WebGL2RenderingContext) ClientWaitSync(sync WebGLSync, flags GLbitfield, timeout GLuint64) (ret GLenum) {
	bindings.CallWebGL2RenderingContextClientWaitSync(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// HasFuncWaitSync returns true if the method "WebGL2RenderingContext.waitSync" exists.
func (this WebGL2RenderingContext) HasFuncWaitSync() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextWaitSync(
		this.ref,
	)
}

// FuncWaitSync returns the method "WebGL2RenderingContext.waitSync".
func (this WebGL2RenderingContext) FuncWaitSync() (fn js.Func[func(sync WebGLSync, flags GLbitfield, timeout GLint64)]) {
	bindings.FuncWebGL2RenderingContextWaitSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WaitSync calls the method "WebGL2RenderingContext.waitSync".
func (this WebGL2RenderingContext) WaitSync(sync WebGLSync, flags GLbitfield, timeout GLint64) (ret js.Void) {
	bindings.CallWebGL2RenderingContextWaitSync(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(flags),
		float64(timeout),
	)

	return
}

// HasFuncGetSyncParameter returns true if the method "WebGL2RenderingContext.getSyncParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetSyncParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetSyncParameter(
		this.ref,
	)
}

// FuncGetSyncParameter returns the method "WebGL2RenderingContext.getSyncParameter".
func (this WebGL2RenderingContext) FuncGetSyncParameter() (fn js.Func[func(sync WebGLSync, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetSyncParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSyncParameter calls the method "WebGL2RenderingContext.getSyncParameter".
func (this WebGL2RenderingContext) GetSyncParameter(sync WebGLSync, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetSyncParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sync.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncCreateTransformFeedback returns true if the method "WebGL2RenderingContext.createTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncCreateTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateTransformFeedback(
		this.ref,
	)
}

// FuncCreateTransformFeedback returns the method "WebGL2RenderingContext.createTransformFeedback".
func (this WebGL2RenderingContext) FuncCreateTransformFeedback() (fn js.Func[func() WebGLTransformFeedback]) {
	bindings.FuncWebGL2RenderingContextCreateTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTransformFeedback calls the method "WebGL2RenderingContext.createTransformFeedback".
func (this WebGL2RenderingContext) CreateTransformFeedback() (ret WebGLTransformFeedback) {
	bindings.CallWebGL2RenderingContextCreateTransformFeedback(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTransformFeedback calls the method "WebGL2RenderingContext.createTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateTransformFeedback() (ret WebGLTransformFeedback, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteTransformFeedback returns true if the method "WebGL2RenderingContext.deleteTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncDeleteTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteTransformFeedback(
		this.ref,
	)
}

// FuncDeleteTransformFeedback returns the method "WebGL2RenderingContext.deleteTransformFeedback".
func (this WebGL2RenderingContext) FuncDeleteTransformFeedback() (fn js.Func[func(tf WebGLTransformFeedback)]) {
	bindings.FuncWebGL2RenderingContextDeleteTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteTransformFeedback calls the method "WebGL2RenderingContext.deleteTransformFeedback".
func (this WebGL2RenderingContext) DeleteTransformFeedback(tf WebGLTransformFeedback) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteTransformFeedback(
		this.ref, js.Pointer(&ret),
		tf.Ref(),
	)

	return
}

// TryDeleteTransformFeedback calls the method "WebGL2RenderingContext.deleteTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteTransformFeedback(tf WebGLTransformFeedback) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		tf.Ref(),
	)

	return
}

// HasFuncIsTransformFeedback returns true if the method "WebGL2RenderingContext.isTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncIsTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsTransformFeedback(
		this.ref,
	)
}

// FuncIsTransformFeedback returns the method "WebGL2RenderingContext.isTransformFeedback".
func (this WebGL2RenderingContext) FuncIsTransformFeedback() (fn js.Func[func(tf WebGLTransformFeedback) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTransformFeedback calls the method "WebGL2RenderingContext.isTransformFeedback".
func (this WebGL2RenderingContext) IsTransformFeedback(tf WebGLTransformFeedback) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsTransformFeedback(
		this.ref, js.Pointer(&ret),
		tf.Ref(),
	)

	return
}

// TryIsTransformFeedback calls the method "WebGL2RenderingContext.isTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsTransformFeedback(tf WebGLTransformFeedback) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		tf.Ref(),
	)

	return
}

// HasFuncBindTransformFeedback returns true if the method "WebGL2RenderingContext.bindTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncBindTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindTransformFeedback(
		this.ref,
	)
}

// FuncBindTransformFeedback returns the method "WebGL2RenderingContext.bindTransformFeedback".
func (this WebGL2RenderingContext) FuncBindTransformFeedback() (fn js.Func[func(target GLenum, tf WebGLTransformFeedback)]) {
	bindings.FuncWebGL2RenderingContextBindTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindTransformFeedback calls the method "WebGL2RenderingContext.bindTransformFeedback".
func (this WebGL2RenderingContext) BindTransformFeedback(target GLenum, tf WebGLTransformFeedback) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindTransformFeedback(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		tf.Ref(),
	)

	return
}

// HasFuncBeginTransformFeedback returns true if the method "WebGL2RenderingContext.beginTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncBeginTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBeginTransformFeedback(
		this.ref,
	)
}

// FuncBeginTransformFeedback returns the method "WebGL2RenderingContext.beginTransformFeedback".
func (this WebGL2RenderingContext) FuncBeginTransformFeedback() (fn js.Func[func(primitiveMode GLenum)]) {
	bindings.FuncWebGL2RenderingContextBeginTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginTransformFeedback calls the method "WebGL2RenderingContext.beginTransformFeedback".
func (this WebGL2RenderingContext) BeginTransformFeedback(primitiveMode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBeginTransformFeedback(
		this.ref, js.Pointer(&ret),
		uint32(primitiveMode),
	)

	return
}

// TryBeginTransformFeedback calls the method "WebGL2RenderingContext.beginTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBeginTransformFeedback(primitiveMode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBeginTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(primitiveMode),
	)

	return
}

// HasFuncEndTransformFeedback returns true if the method "WebGL2RenderingContext.endTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncEndTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextEndTransformFeedback(
		this.ref,
	)
}

// FuncEndTransformFeedback returns the method "WebGL2RenderingContext.endTransformFeedback".
func (this WebGL2RenderingContext) FuncEndTransformFeedback() (fn js.Func[func()]) {
	bindings.FuncWebGL2RenderingContextEndTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndTransformFeedback calls the method "WebGL2RenderingContext.endTransformFeedback".
func (this WebGL2RenderingContext) EndTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextEndTransformFeedback(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEndTransformFeedback calls the method "WebGL2RenderingContext.endTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEndTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEndTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTransformFeedbackVaryings returns true if the method "WebGL2RenderingContext.transformFeedbackVaryings" exists.
func (this WebGL2RenderingContext) HasFuncTransformFeedbackVaryings() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTransformFeedbackVaryings(
		this.ref,
	)
}

// FuncTransformFeedbackVaryings returns the method "WebGL2RenderingContext.transformFeedbackVaryings".
func (this WebGL2RenderingContext) FuncTransformFeedbackVaryings() (fn js.Func[func(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum)]) {
	bindings.FuncWebGL2RenderingContextTransformFeedbackVaryings(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransformFeedbackVaryings calls the method "WebGL2RenderingContext.transformFeedbackVaryings".
func (this WebGL2RenderingContext) TransformFeedbackVaryings(program WebGLProgram, varyings js.Array[js.String], bufferMode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTransformFeedbackVaryings(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		varyings.Ref(),
		uint32(bufferMode),
	)

	return
}

// HasFuncGetTransformFeedbackVarying returns true if the method "WebGL2RenderingContext.getTransformFeedbackVarying" exists.
func (this WebGL2RenderingContext) HasFuncGetTransformFeedbackVarying() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetTransformFeedbackVarying(
		this.ref,
	)
}

// FuncGetTransformFeedbackVarying returns the method "WebGL2RenderingContext.getTransformFeedbackVarying".
func (this WebGL2RenderingContext) FuncGetTransformFeedbackVarying() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	bindings.FuncWebGL2RenderingContextGetTransformFeedbackVarying(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTransformFeedbackVarying calls the method "WebGL2RenderingContext.getTransformFeedbackVarying".
func (this WebGL2RenderingContext) GetTransformFeedbackVarying(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetTransformFeedbackVarying(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasFuncPauseTransformFeedback returns true if the method "WebGL2RenderingContext.pauseTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncPauseTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextPauseTransformFeedback(
		this.ref,
	)
}

// FuncPauseTransformFeedback returns the method "WebGL2RenderingContext.pauseTransformFeedback".
func (this WebGL2RenderingContext) FuncPauseTransformFeedback() (fn js.Func[func()]) {
	bindings.FuncWebGL2RenderingContextPauseTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PauseTransformFeedback calls the method "WebGL2RenderingContext.pauseTransformFeedback".
func (this WebGL2RenderingContext) PauseTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextPauseTransformFeedback(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPauseTransformFeedback calls the method "WebGL2RenderingContext.pauseTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryPauseTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextPauseTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncResumeTransformFeedback returns true if the method "WebGL2RenderingContext.resumeTransformFeedback" exists.
func (this WebGL2RenderingContext) HasFuncResumeTransformFeedback() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextResumeTransformFeedback(
		this.ref,
	)
}

// FuncResumeTransformFeedback returns the method "WebGL2RenderingContext.resumeTransformFeedback".
func (this WebGL2RenderingContext) FuncResumeTransformFeedback() (fn js.Func[func()]) {
	bindings.FuncWebGL2RenderingContextResumeTransformFeedback(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResumeTransformFeedback calls the method "WebGL2RenderingContext.resumeTransformFeedback".
func (this WebGL2RenderingContext) ResumeTransformFeedback() (ret js.Void) {
	bindings.CallWebGL2RenderingContextResumeTransformFeedback(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryResumeTransformFeedback calls the method "WebGL2RenderingContext.resumeTransformFeedback"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryResumeTransformFeedback() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextResumeTransformFeedback(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBindBufferBase returns true if the method "WebGL2RenderingContext.bindBufferBase" exists.
func (this WebGL2RenderingContext) HasFuncBindBufferBase() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindBufferBase(
		this.ref,
	)
}

// FuncBindBufferBase returns the method "WebGL2RenderingContext.bindBufferBase".
func (this WebGL2RenderingContext) FuncBindBufferBase() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer)]) {
	bindings.FuncWebGL2RenderingContextBindBufferBase(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindBufferBase calls the method "WebGL2RenderingContext.bindBufferBase".
func (this WebGL2RenderingContext) BindBufferBase(target GLenum, index GLuint, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBufferBase(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
		buffer.Ref(),
	)

	return
}

// HasFuncBindBufferRange returns true if the method "WebGL2RenderingContext.bindBufferRange" exists.
func (this WebGL2RenderingContext) HasFuncBindBufferRange() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindBufferRange(
		this.ref,
	)
}

// FuncBindBufferRange returns the method "WebGL2RenderingContext.bindBufferRange".
func (this WebGL2RenderingContext) FuncBindBufferRange() (fn js.Func[func(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr)]) {
	bindings.FuncWebGL2RenderingContextBindBufferRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindBufferRange calls the method "WebGL2RenderingContext.bindBufferRange".
func (this WebGL2RenderingContext) BindBufferRange(target GLenum, index GLuint, buffer WebGLBuffer, offset GLintptr, size GLsizeiptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBufferRange(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncGetIndexedParameter returns true if the method "WebGL2RenderingContext.getIndexedParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetIndexedParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetIndexedParameter(
		this.ref,
	)
}

// FuncGetIndexedParameter returns the method "WebGL2RenderingContext.getIndexedParameter".
func (this WebGL2RenderingContext) FuncGetIndexedParameter() (fn js.Func[func(target GLenum, index GLuint) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetIndexedParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetIndexedParameter calls the method "WebGL2RenderingContext.getIndexedParameter".
func (this WebGL2RenderingContext) GetIndexedParameter(target GLenum, index GLuint) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetIndexedParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
	)

	return
}

// HasFuncGetUniformIndices returns true if the method "WebGL2RenderingContext.getUniformIndices" exists.
func (this WebGL2RenderingContext) HasFuncGetUniformIndices() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetUniformIndices(
		this.ref,
	)
}

// FuncGetUniformIndices returns the method "WebGL2RenderingContext.getUniformIndices".
func (this WebGL2RenderingContext) FuncGetUniformIndices() (fn js.Func[func(program WebGLProgram, uniformNames js.Array[js.String]) js.Array[GLuint]]) {
	bindings.FuncWebGL2RenderingContextGetUniformIndices(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniformIndices calls the method "WebGL2RenderingContext.getUniformIndices".
func (this WebGL2RenderingContext) GetUniformIndices(program WebGLProgram, uniformNames js.Array[js.String]) (ret js.Array[GLuint]) {
	bindings.CallWebGL2RenderingContextGetUniformIndices(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformNames.Ref(),
	)

	return
}

// HasFuncGetActiveUniforms returns true if the method "WebGL2RenderingContext.getActiveUniforms" exists.
func (this WebGL2RenderingContext) HasFuncGetActiveUniforms() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetActiveUniforms(
		this.ref,
	)
}

// FuncGetActiveUniforms returns the method "WebGL2RenderingContext.getActiveUniforms".
func (this WebGL2RenderingContext) FuncGetActiveUniforms() (fn js.Func[func(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetActiveUniforms(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveUniforms calls the method "WebGL2RenderingContext.getActiveUniforms".
func (this WebGL2RenderingContext) GetActiveUniforms(program WebGLProgram, uniformIndices js.Array[GLuint], pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetActiveUniforms(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformIndices.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncGetUniformBlockIndex returns true if the method "WebGL2RenderingContext.getUniformBlockIndex" exists.
func (this WebGL2RenderingContext) HasFuncGetUniformBlockIndex() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetUniformBlockIndex(
		this.ref,
	)
}

// FuncGetUniformBlockIndex returns the method "WebGL2RenderingContext.getUniformBlockIndex".
func (this WebGL2RenderingContext) FuncGetUniformBlockIndex() (fn js.Func[func(program WebGLProgram, uniformBlockName js.String) GLuint]) {
	bindings.FuncWebGL2RenderingContextGetUniformBlockIndex(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniformBlockIndex calls the method "WebGL2RenderingContext.getUniformBlockIndex".
func (this WebGL2RenderingContext) GetUniformBlockIndex(program WebGLProgram, uniformBlockName js.String) (ret GLuint) {
	bindings.CallWebGL2RenderingContextGetUniformBlockIndex(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uniformBlockName.Ref(),
	)

	return
}

// HasFuncGetActiveUniformBlockParameter returns true if the method "WebGL2RenderingContext.getActiveUniformBlockParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetActiveUniformBlockParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.ref,
	)
}

// FuncGetActiveUniformBlockParameter returns the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
func (this WebGL2RenderingContext) FuncGetActiveUniformBlockParameter() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveUniformBlockParameter calls the method "WebGL2RenderingContext.getActiveUniformBlockParameter".
func (this WebGL2RenderingContext) GetActiveUniformBlockParameter(program WebGLProgram, uniformBlockIndex GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetActiveUniformBlockParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(pname),
	)

	return
}

// HasFuncGetActiveUniformBlockName returns true if the method "WebGL2RenderingContext.getActiveUniformBlockName" exists.
func (this WebGL2RenderingContext) HasFuncGetActiveUniformBlockName() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetActiveUniformBlockName(
		this.ref,
	)
}

// FuncGetActiveUniformBlockName returns the method "WebGL2RenderingContext.getActiveUniformBlockName".
func (this WebGL2RenderingContext) FuncGetActiveUniformBlockName() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint) js.String]) {
	bindings.FuncWebGL2RenderingContextGetActiveUniformBlockName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveUniformBlockName calls the method "WebGL2RenderingContext.getActiveUniformBlockName".
func (this WebGL2RenderingContext) GetActiveUniformBlockName(program WebGLProgram, uniformBlockIndex GLuint) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetActiveUniformBlockName(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
	)

	return
}

// HasFuncUniformBlockBinding returns true if the method "WebGL2RenderingContext.uniformBlockBinding" exists.
func (this WebGL2RenderingContext) HasFuncUniformBlockBinding() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniformBlockBinding(
		this.ref,
	)
}

// FuncUniformBlockBinding returns the method "WebGL2RenderingContext.uniformBlockBinding".
func (this WebGL2RenderingContext) FuncUniformBlockBinding() (fn js.Func[func(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint)]) {
	bindings.FuncWebGL2RenderingContextUniformBlockBinding(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UniformBlockBinding calls the method "WebGL2RenderingContext.uniformBlockBinding".
func (this WebGL2RenderingContext) UniformBlockBinding(program WebGLProgram, uniformBlockIndex GLuint, uniformBlockBinding GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniformBlockBinding(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(uniformBlockIndex),
		uint32(uniformBlockBinding),
	)

	return
}

// HasFuncCreateVertexArray returns true if the method "WebGL2RenderingContext.createVertexArray" exists.
func (this WebGL2RenderingContext) HasFuncCreateVertexArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateVertexArray(
		this.ref,
	)
}

// FuncCreateVertexArray returns the method "WebGL2RenderingContext.createVertexArray".
func (this WebGL2RenderingContext) FuncCreateVertexArray() (fn js.Func[func() WebGLVertexArrayObject]) {
	bindings.FuncWebGL2RenderingContextCreateVertexArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateVertexArray calls the method "WebGL2RenderingContext.createVertexArray".
func (this WebGL2RenderingContext) CreateVertexArray() (ret WebGLVertexArrayObject) {
	bindings.CallWebGL2RenderingContextCreateVertexArray(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateVertexArray calls the method "WebGL2RenderingContext.createVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateVertexArray() (ret WebGLVertexArrayObject, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateVertexArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteVertexArray returns true if the method "WebGL2RenderingContext.deleteVertexArray" exists.
func (this WebGL2RenderingContext) HasFuncDeleteVertexArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteVertexArray(
		this.ref,
	)
}

// FuncDeleteVertexArray returns the method "WebGL2RenderingContext.deleteVertexArray".
func (this WebGL2RenderingContext) FuncDeleteVertexArray() (fn js.Func[func(vertexArray WebGLVertexArrayObject)]) {
	bindings.FuncWebGL2RenderingContextDeleteVertexArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteVertexArray calls the method "WebGL2RenderingContext.deleteVertexArray".
func (this WebGL2RenderingContext) DeleteVertexArray(vertexArray WebGLVertexArrayObject) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteVertexArray(
		this.ref, js.Pointer(&ret),
		vertexArray.Ref(),
	)

	return
}

// TryDeleteVertexArray calls the method "WebGL2RenderingContext.deleteVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteVertexArray(vertexArray WebGLVertexArrayObject) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteVertexArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		vertexArray.Ref(),
	)

	return
}

// HasFuncIsVertexArray returns true if the method "WebGL2RenderingContext.isVertexArray" exists.
func (this WebGL2RenderingContext) HasFuncIsVertexArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsVertexArray(
		this.ref,
	)
}

// FuncIsVertexArray returns the method "WebGL2RenderingContext.isVertexArray".
func (this WebGL2RenderingContext) FuncIsVertexArray() (fn js.Func[func(vertexArray WebGLVertexArrayObject) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsVertexArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsVertexArray calls the method "WebGL2RenderingContext.isVertexArray".
func (this WebGL2RenderingContext) IsVertexArray(vertexArray WebGLVertexArrayObject) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsVertexArray(
		this.ref, js.Pointer(&ret),
		vertexArray.Ref(),
	)

	return
}

// TryIsVertexArray calls the method "WebGL2RenderingContext.isVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsVertexArray(vertexArray WebGLVertexArrayObject) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsVertexArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		vertexArray.Ref(),
	)

	return
}

// HasFuncBindVertexArray returns true if the method "WebGL2RenderingContext.bindVertexArray" exists.
func (this WebGL2RenderingContext) HasFuncBindVertexArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindVertexArray(
		this.ref,
	)
}

// FuncBindVertexArray returns the method "WebGL2RenderingContext.bindVertexArray".
func (this WebGL2RenderingContext) FuncBindVertexArray() (fn js.Func[func(array WebGLVertexArrayObject)]) {
	bindings.FuncWebGL2RenderingContextBindVertexArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindVertexArray calls the method "WebGL2RenderingContext.bindVertexArray".
func (this WebGL2RenderingContext) BindVertexArray(array WebGLVertexArrayObject) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindVertexArray(
		this.ref, js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryBindVertexArray calls the method "WebGL2RenderingContext.bindVertexArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBindVertexArray(array WebGLVertexArrayObject) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBindVertexArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasFuncGetContextAttributes returns true if the method "WebGL2RenderingContext.getContextAttributes" exists.
func (this WebGL2RenderingContext) HasFuncGetContextAttributes() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetContextAttributes(
		this.ref,
	)
}

// FuncGetContextAttributes returns the method "WebGL2RenderingContext.getContextAttributes".
func (this WebGL2RenderingContext) FuncGetContextAttributes() (fn js.Func[func() WebGLContextAttributes]) {
	bindings.FuncWebGL2RenderingContextGetContextAttributes(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContextAttributes calls the method "WebGL2RenderingContext.getContextAttributes".
func (this WebGL2RenderingContext) GetContextAttributes() (ret WebGLContextAttributes) {
	bindings.CallWebGL2RenderingContextGetContextAttributes(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetContextAttributes calls the method "WebGL2RenderingContext.getContextAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetContextAttributes() (ret WebGLContextAttributes, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetContextAttributes(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsContextLost returns true if the method "WebGL2RenderingContext.isContextLost" exists.
func (this WebGL2RenderingContext) HasFuncIsContextLost() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsContextLost(
		this.ref,
	)
}

// FuncIsContextLost returns the method "WebGL2RenderingContext.isContextLost".
func (this WebGL2RenderingContext) FuncIsContextLost() (fn js.Func[func() bool]) {
	bindings.FuncWebGL2RenderingContextIsContextLost(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsContextLost calls the method "WebGL2RenderingContext.isContextLost".
func (this WebGL2RenderingContext) IsContextLost() (ret bool) {
	bindings.CallWebGL2RenderingContextIsContextLost(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "WebGL2RenderingContext.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsContextLost(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSupportedExtensions returns true if the method "WebGL2RenderingContext.getSupportedExtensions" exists.
func (this WebGL2RenderingContext) HasFuncGetSupportedExtensions() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetSupportedExtensions(
		this.ref,
	)
}

// FuncGetSupportedExtensions returns the method "WebGL2RenderingContext.getSupportedExtensions".
func (this WebGL2RenderingContext) FuncGetSupportedExtensions() (fn js.Func[func() js.Array[js.String]]) {
	bindings.FuncWebGL2RenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSupportedExtensions calls the method "WebGL2RenderingContext.getSupportedExtensions".
func (this WebGL2RenderingContext) GetSupportedExtensions() (ret js.Array[js.String]) {
	bindings.CallWebGL2RenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSupportedExtensions calls the method "WebGL2RenderingContext.getSupportedExtensions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetSupportedExtensions() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetSupportedExtensions(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetExtension returns true if the method "WebGL2RenderingContext.getExtension" exists.
func (this WebGL2RenderingContext) HasFuncGetExtension() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetExtension(
		this.ref,
	)
}

// FuncGetExtension returns the method "WebGL2RenderingContext.getExtension".
func (this WebGL2RenderingContext) FuncGetExtension() (fn js.Func[func(name js.String) js.Object]) {
	bindings.FuncWebGL2RenderingContextGetExtension(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetExtension calls the method "WebGL2RenderingContext.getExtension".
func (this WebGL2RenderingContext) GetExtension(name js.String) (ret js.Object) {
	bindings.CallWebGL2RenderingContextGetExtension(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetExtension calls the method "WebGL2RenderingContext.getExtension"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetExtension(name js.String) (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetExtension(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncActiveTexture returns true if the method "WebGL2RenderingContext.activeTexture" exists.
func (this WebGL2RenderingContext) HasFuncActiveTexture() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextActiveTexture(
		this.ref,
	)
}

// FuncActiveTexture returns the method "WebGL2RenderingContext.activeTexture".
func (this WebGL2RenderingContext) FuncActiveTexture() (fn js.Func[func(texture GLenum)]) {
	bindings.FuncWebGL2RenderingContextActiveTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ActiveTexture calls the method "WebGL2RenderingContext.activeTexture".
func (this WebGL2RenderingContext) ActiveTexture(texture GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextActiveTexture(
		this.ref, js.Pointer(&ret),
		uint32(texture),
	)

	return
}

// TryActiveTexture calls the method "WebGL2RenderingContext.activeTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryActiveTexture(texture GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextActiveTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(texture),
	)

	return
}

// HasFuncAttachShader returns true if the method "WebGL2RenderingContext.attachShader" exists.
func (this WebGL2RenderingContext) HasFuncAttachShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextAttachShader(
		this.ref,
	)
}

// FuncAttachShader returns the method "WebGL2RenderingContext.attachShader".
func (this WebGL2RenderingContext) FuncAttachShader() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	bindings.FuncWebGL2RenderingContextAttachShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AttachShader calls the method "WebGL2RenderingContext.attachShader".
func (this WebGL2RenderingContext) AttachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextAttachShader(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasFuncBindAttribLocation returns true if the method "WebGL2RenderingContext.bindAttribLocation" exists.
func (this WebGL2RenderingContext) HasFuncBindAttribLocation() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindAttribLocation(
		this.ref,
	)
}

// FuncBindAttribLocation returns the method "WebGL2RenderingContext.bindAttribLocation".
func (this WebGL2RenderingContext) FuncBindAttribLocation() (fn js.Func[func(program WebGLProgram, index GLuint, name js.String)]) {
	bindings.FuncWebGL2RenderingContextBindAttribLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindAttribLocation calls the method "WebGL2RenderingContext.bindAttribLocation".
func (this WebGL2RenderingContext) BindAttribLocation(program WebGLProgram, index GLuint, name js.String) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindAttribLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
		name.Ref(),
	)

	return
}

// HasFuncBindBuffer returns true if the method "WebGL2RenderingContext.bindBuffer" exists.
func (this WebGL2RenderingContext) HasFuncBindBuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindBuffer(
		this.ref,
	)
}

// FuncBindBuffer returns the method "WebGL2RenderingContext.bindBuffer".
func (this WebGL2RenderingContext) FuncBindBuffer() (fn js.Func[func(target GLenum, buffer WebGLBuffer)]) {
	bindings.FuncWebGL2RenderingContextBindBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindBuffer calls the method "WebGL2RenderingContext.bindBuffer".
func (this WebGL2RenderingContext) BindBuffer(target GLenum, buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		buffer.Ref(),
	)

	return
}

// HasFuncBindFramebuffer returns true if the method "WebGL2RenderingContext.bindFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncBindFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindFramebuffer(
		this.ref,
	)
}

// FuncBindFramebuffer returns the method "WebGL2RenderingContext.bindFramebuffer".
func (this WebGL2RenderingContext) FuncBindFramebuffer() (fn js.Func[func(target GLenum, framebuffer WebGLFramebuffer)]) {
	bindings.FuncWebGL2RenderingContextBindFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindFramebuffer calls the method "WebGL2RenderingContext.bindFramebuffer".
func (this WebGL2RenderingContext) BindFramebuffer(target GLenum, framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindFramebuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		framebuffer.Ref(),
	)

	return
}

// HasFuncBindRenderbuffer returns true if the method "WebGL2RenderingContext.bindRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFuncBindRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindRenderbuffer(
		this.ref,
	)
}

// FuncBindRenderbuffer returns the method "WebGL2RenderingContext.bindRenderbuffer".
func (this WebGL2RenderingContext) FuncBindRenderbuffer() (fn js.Func[func(target GLenum, renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGL2RenderingContextBindRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindRenderbuffer calls the method "WebGL2RenderingContext.bindRenderbuffer".
func (this WebGL2RenderingContext) BindRenderbuffer(target GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindRenderbuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncBindTexture returns true if the method "WebGL2RenderingContext.bindTexture" exists.
func (this WebGL2RenderingContext) HasFuncBindTexture() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBindTexture(
		this.ref,
	)
}

// FuncBindTexture returns the method "WebGL2RenderingContext.bindTexture".
func (this WebGL2RenderingContext) FuncBindTexture() (fn js.Func[func(target GLenum, texture WebGLTexture)]) {
	bindings.FuncWebGL2RenderingContextBindTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BindTexture calls the method "WebGL2RenderingContext.bindTexture".
func (this WebGL2RenderingContext) BindTexture(target GLenum, texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBindTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		texture.Ref(),
	)

	return
}

// HasFuncBlendColor returns true if the method "WebGL2RenderingContext.blendColor" exists.
func (this WebGL2RenderingContext) HasFuncBlendColor() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlendColor(
		this.ref,
	)
}

// FuncBlendColor returns the method "WebGL2RenderingContext.blendColor".
func (this WebGL2RenderingContext) FuncBlendColor() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	bindings.FuncWebGL2RenderingContextBlendColor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendColor calls the method "WebGL2RenderingContext.blendColor".
func (this WebGL2RenderingContext) BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendColor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasFuncBlendEquation returns true if the method "WebGL2RenderingContext.blendEquation" exists.
func (this WebGL2RenderingContext) HasFuncBlendEquation() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlendEquation(
		this.ref,
	)
}

// FuncBlendEquation returns the method "WebGL2RenderingContext.blendEquation".
func (this WebGL2RenderingContext) FuncBlendEquation() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGL2RenderingContextBlendEquation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendEquation calls the method "WebGL2RenderingContext.blendEquation".
func (this WebGL2RenderingContext) BlendEquation(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendEquation(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryBlendEquation calls the method "WebGL2RenderingContext.blendEquation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryBlendEquation(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextBlendEquation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncBlendEquationSeparate returns true if the method "WebGL2RenderingContext.blendEquationSeparate" exists.
func (this WebGL2RenderingContext) HasFuncBlendEquationSeparate() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlendEquationSeparate(
		this.ref,
	)
}

// FuncBlendEquationSeparate returns the method "WebGL2RenderingContext.blendEquationSeparate".
func (this WebGL2RenderingContext) FuncBlendEquationSeparate() (fn js.Func[func(modeRGB GLenum, modeAlpha GLenum)]) {
	bindings.FuncWebGL2RenderingContextBlendEquationSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendEquationSeparate calls the method "WebGL2RenderingContext.blendEquationSeparate".
func (this WebGL2RenderingContext) BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendEquationSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// HasFuncBlendFunc returns true if the method "WebGL2RenderingContext.blendFunc" exists.
func (this WebGL2RenderingContext) HasFuncBlendFunc() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlendFunc(
		this.ref,
	)
}

// FuncBlendFunc returns the method "WebGL2RenderingContext.blendFunc".
func (this WebGL2RenderingContext) FuncBlendFunc() (fn js.Func[func(sfactor GLenum, dfactor GLenum)]) {
	bindings.FuncWebGL2RenderingContextBlendFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendFunc calls the method "WebGL2RenderingContext.blendFunc".
func (this WebGL2RenderingContext) BlendFunc(sfactor GLenum, dfactor GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendFunc(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(sfactor),
		uint32(dfactor),
	)

	return
}

// HasFuncBlendFuncSeparate returns true if the method "WebGL2RenderingContext.blendFuncSeparate" exists.
func (this WebGL2RenderingContext) HasFuncBlendFuncSeparate() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextBlendFuncSeparate(
		this.ref,
	)
}

// FuncBlendFuncSeparate returns the method "WebGL2RenderingContext.blendFuncSeparate".
func (this WebGL2RenderingContext) FuncBlendFuncSeparate() (fn js.Func[func(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	bindings.FuncWebGL2RenderingContextBlendFuncSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BlendFuncSeparate calls the method "WebGL2RenderingContext.blendFuncSeparate".
func (this WebGL2RenderingContext) BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextBlendFuncSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// HasFuncCheckFramebufferStatus returns true if the method "WebGL2RenderingContext.checkFramebufferStatus" exists.
func (this WebGL2RenderingContext) HasFuncCheckFramebufferStatus() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCheckFramebufferStatus(
		this.ref,
	)
}

// FuncCheckFramebufferStatus returns the method "WebGL2RenderingContext.checkFramebufferStatus".
func (this WebGL2RenderingContext) FuncCheckFramebufferStatus() (fn js.Func[func(target GLenum) GLenum]) {
	bindings.FuncWebGL2RenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckFramebufferStatus calls the method "WebGL2RenderingContext.checkFramebufferStatus".
func (this WebGL2RenderingContext) CheckFramebufferStatus(target GLenum) (ret GLenum) {
	bindings.CallWebGL2RenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryCheckFramebufferStatus calls the method "WebGL2RenderingContext.checkFramebufferStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCheckFramebufferStatus(target GLenum) (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCheckFramebufferStatus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncClear returns true if the method "WebGL2RenderingContext.clear" exists.
func (this WebGL2RenderingContext) HasFuncClear() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClear(
		this.ref,
	)
}

// FuncClear returns the method "WebGL2RenderingContext.clear".
func (this WebGL2RenderingContext) FuncClear() (fn js.Func[func(mask GLbitfield)]) {
	bindings.FuncWebGL2RenderingContextClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "WebGL2RenderingContext.clear".
func (this WebGL2RenderingContext) Clear(mask GLbitfield) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClear(
		this.ref, js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryClear calls the method "WebGL2RenderingContext.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClear(mask GLbitfield) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasFuncClearColor returns true if the method "WebGL2RenderingContext.clearColor" exists.
func (this WebGL2RenderingContext) HasFuncClearColor() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearColor(
		this.ref,
	)
}

// FuncClearColor returns the method "WebGL2RenderingContext.clearColor".
func (this WebGL2RenderingContext) FuncClearColor() (fn js.Func[func(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf)]) {
	bindings.FuncWebGL2RenderingContextClearColor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearColor calls the method "WebGL2RenderingContext.clearColor".
func (this WebGL2RenderingContext) ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearColor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(red),
		float32(green),
		float32(blue),
		float32(alpha),
	)

	return
}

// HasFuncClearDepth returns true if the method "WebGL2RenderingContext.clearDepth" exists.
func (this WebGL2RenderingContext) HasFuncClearDepth() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearDepth(
		this.ref,
	)
}

// FuncClearDepth returns the method "WebGL2RenderingContext.clearDepth".
func (this WebGL2RenderingContext) FuncClearDepth() (fn js.Func[func(depth GLclampf)]) {
	bindings.FuncWebGL2RenderingContextClearDepth(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearDepth calls the method "WebGL2RenderingContext.clearDepth".
func (this WebGL2RenderingContext) ClearDepth(depth GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearDepth(
		this.ref, js.Pointer(&ret),
		float32(depth),
	)

	return
}

// TryClearDepth calls the method "WebGL2RenderingContext.clearDepth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearDepth(depth GLclampf) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearDepth(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(depth),
	)

	return
}

// HasFuncClearStencil returns true if the method "WebGL2RenderingContext.clearStencil" exists.
func (this WebGL2RenderingContext) HasFuncClearStencil() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextClearStencil(
		this.ref,
	)
}

// FuncClearStencil returns the method "WebGL2RenderingContext.clearStencil".
func (this WebGL2RenderingContext) FuncClearStencil() (fn js.Func[func(s GLint)]) {
	bindings.FuncWebGL2RenderingContextClearStencil(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearStencil calls the method "WebGL2RenderingContext.clearStencil".
func (this WebGL2RenderingContext) ClearStencil(s GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextClearStencil(
		this.ref, js.Pointer(&ret),
		int32(s),
	)

	return
}

// TryClearStencil calls the method "WebGL2RenderingContext.clearStencil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryClearStencil(s GLint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextClearStencil(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(s),
	)

	return
}

// HasFuncColorMask returns true if the method "WebGL2RenderingContext.colorMask" exists.
func (this WebGL2RenderingContext) HasFuncColorMask() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextColorMask(
		this.ref,
	)
}

// FuncColorMask returns the method "WebGL2RenderingContext.colorMask".
func (this WebGL2RenderingContext) FuncColorMask() (fn js.Func[func(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean)]) {
	bindings.FuncWebGL2RenderingContextColorMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ColorMask calls the method "WebGL2RenderingContext.colorMask".
func (this WebGL2RenderingContext) ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextColorMask(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(red)),
		js.Bool(bool(green)),
		js.Bool(bool(blue)),
		js.Bool(bool(alpha)),
	)

	return
}

// HasFuncCompileShader returns true if the method "WebGL2RenderingContext.compileShader" exists.
func (this WebGL2RenderingContext) HasFuncCompileShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCompileShader(
		this.ref,
	)
}

// FuncCompileShader returns the method "WebGL2RenderingContext.compileShader".
func (this WebGL2RenderingContext) FuncCompileShader() (fn js.Func[func(shader WebGLShader)]) {
	bindings.FuncWebGL2RenderingContextCompileShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CompileShader calls the method "WebGL2RenderingContext.compileShader".
func (this WebGL2RenderingContext) CompileShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCompileShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryCompileShader calls the method "WebGL2RenderingContext.compileShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCompileShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCompileShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncCopyTexImage2D returns true if the method "WebGL2RenderingContext.copyTexImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCopyTexImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCopyTexImage2D(
		this.ref,
	)
}

// FuncCopyTexImage2D returns the method "WebGL2RenderingContext.copyTexImage2D".
func (this WebGL2RenderingContext) FuncCopyTexImage2D() (fn js.Func[func(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint)]) {
	bindings.FuncWebGL2RenderingContextCopyTexImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTexImage2D calls the method "WebGL2RenderingContext.copyTexImage2D".
func (this WebGL2RenderingContext) CopyTexImage2D(target GLenum, level GLint, internalformat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCopyTexSubImage2D returns true if the method "WebGL2RenderingContext.copyTexSubImage2D" exists.
func (this WebGL2RenderingContext) HasFuncCopyTexSubImage2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCopyTexSubImage2D(
		this.ref,
	)
}

// FuncCopyTexSubImage2D returns the method "WebGL2RenderingContext.copyTexSubImage2D".
func (this WebGL2RenderingContext) FuncCopyTexSubImage2D() (fn js.Func[func(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextCopyTexSubImage2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTexSubImage2D calls the method "WebGL2RenderingContext.copyTexSubImage2D".
func (this WebGL2RenderingContext) CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCopyTexSubImage2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncCreateBuffer returns true if the method "WebGL2RenderingContext.createBuffer" exists.
func (this WebGL2RenderingContext) HasFuncCreateBuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateBuffer(
		this.ref,
	)
}

// FuncCreateBuffer returns the method "WebGL2RenderingContext.createBuffer".
func (this WebGL2RenderingContext) FuncCreateBuffer() (fn js.Func[func() WebGLBuffer]) {
	bindings.FuncWebGL2RenderingContextCreateBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBuffer calls the method "WebGL2RenderingContext.createBuffer".
func (this WebGL2RenderingContext) CreateBuffer() (ret WebGLBuffer) {
	bindings.CallWebGL2RenderingContextCreateBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateBuffer calls the method "WebGL2RenderingContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateBuffer() (ret WebGLBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateFramebuffer returns true if the method "WebGL2RenderingContext.createFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncCreateFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateFramebuffer(
		this.ref,
	)
}

// FuncCreateFramebuffer returns the method "WebGL2RenderingContext.createFramebuffer".
func (this WebGL2RenderingContext) FuncCreateFramebuffer() (fn js.Func[func() WebGLFramebuffer]) {
	bindings.FuncWebGL2RenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateFramebuffer calls the method "WebGL2RenderingContext.createFramebuffer".
func (this WebGL2RenderingContext) CreateFramebuffer() (ret WebGLFramebuffer) {
	bindings.CallWebGL2RenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateFramebuffer calls the method "WebGL2RenderingContext.createFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateFramebuffer() (ret WebGLFramebuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateProgram returns true if the method "WebGL2RenderingContext.createProgram" exists.
func (this WebGL2RenderingContext) HasFuncCreateProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateProgram(
		this.ref,
	)
}

// FuncCreateProgram returns the method "WebGL2RenderingContext.createProgram".
func (this WebGL2RenderingContext) FuncCreateProgram() (fn js.Func[func() WebGLProgram]) {
	bindings.FuncWebGL2RenderingContextCreateProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateProgram calls the method "WebGL2RenderingContext.createProgram".
func (this WebGL2RenderingContext) CreateProgram() (ret WebGLProgram) {
	bindings.CallWebGL2RenderingContextCreateProgram(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateProgram calls the method "WebGL2RenderingContext.createProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateProgram() (ret WebGLProgram, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateRenderbuffer returns true if the method "WebGL2RenderingContext.createRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFuncCreateRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateRenderbuffer(
		this.ref,
	)
}

// FuncCreateRenderbuffer returns the method "WebGL2RenderingContext.createRenderbuffer".
func (this WebGL2RenderingContext) FuncCreateRenderbuffer() (fn js.Func[func() WebGLRenderbuffer]) {
	bindings.FuncWebGL2RenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRenderbuffer calls the method "WebGL2RenderingContext.createRenderbuffer".
func (this WebGL2RenderingContext) CreateRenderbuffer() (ret WebGLRenderbuffer) {
	bindings.CallWebGL2RenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateRenderbuffer calls the method "WebGL2RenderingContext.createRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateRenderbuffer() (ret WebGLRenderbuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateShader returns true if the method "WebGL2RenderingContext.createShader" exists.
func (this WebGL2RenderingContext) HasFuncCreateShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateShader(
		this.ref,
	)
}

// FuncCreateShader returns the method "WebGL2RenderingContext.createShader".
func (this WebGL2RenderingContext) FuncCreateShader() (fn js.Func[func(typ GLenum) WebGLShader]) {
	bindings.FuncWebGL2RenderingContextCreateShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateShader calls the method "WebGL2RenderingContext.createShader".
func (this WebGL2RenderingContext) CreateShader(typ GLenum) (ret WebGLShader) {
	bindings.CallWebGL2RenderingContextCreateShader(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCreateShader calls the method "WebGL2RenderingContext.createShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateShader(typ GLenum) (ret WebGLShader, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncCreateTexture returns true if the method "WebGL2RenderingContext.createTexture" exists.
func (this WebGL2RenderingContext) HasFuncCreateTexture() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCreateTexture(
		this.ref,
	)
}

// FuncCreateTexture returns the method "WebGL2RenderingContext.createTexture".
func (this WebGL2RenderingContext) FuncCreateTexture() (fn js.Func[func() WebGLTexture]) {
	bindings.FuncWebGL2RenderingContextCreateTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTexture calls the method "WebGL2RenderingContext.createTexture".
func (this WebGL2RenderingContext) CreateTexture() (ret WebGLTexture) {
	bindings.CallWebGL2RenderingContextCreateTexture(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTexture calls the method "WebGL2RenderingContext.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCreateTexture() (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCreateTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCullFace returns true if the method "WebGL2RenderingContext.cullFace" exists.
func (this WebGL2RenderingContext) HasFuncCullFace() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextCullFace(
		this.ref,
	)
}

// FuncCullFace returns the method "WebGL2RenderingContext.cullFace".
func (this WebGL2RenderingContext) FuncCullFace() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGL2RenderingContextCullFace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CullFace calls the method "WebGL2RenderingContext.cullFace".
func (this WebGL2RenderingContext) CullFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextCullFace(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryCullFace calls the method "WebGL2RenderingContext.cullFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryCullFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextCullFace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncDeleteBuffer returns true if the method "WebGL2RenderingContext.deleteBuffer" exists.
func (this WebGL2RenderingContext) HasFuncDeleteBuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteBuffer(
		this.ref,
	)
}

// FuncDeleteBuffer returns the method "WebGL2RenderingContext.deleteBuffer".
func (this WebGL2RenderingContext) FuncDeleteBuffer() (fn js.Func[func(buffer WebGLBuffer)]) {
	bindings.FuncWebGL2RenderingContextDeleteBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteBuffer calls the method "WebGL2RenderingContext.deleteBuffer".
func (this WebGL2RenderingContext) DeleteBuffer(buffer WebGLBuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteBuffer(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryDeleteBuffer calls the method "WebGL2RenderingContext.deleteBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteBuffer(buffer WebGLBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncDeleteFramebuffer returns true if the method "WebGL2RenderingContext.deleteFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncDeleteFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteFramebuffer(
		this.ref,
	)
}

// FuncDeleteFramebuffer returns the method "WebGL2RenderingContext.deleteFramebuffer".
func (this WebGL2RenderingContext) FuncDeleteFramebuffer() (fn js.Func[func(framebuffer WebGLFramebuffer)]) {
	bindings.FuncWebGL2RenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteFramebuffer calls the method "WebGL2RenderingContext.deleteFramebuffer".
func (this WebGL2RenderingContext) DeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryDeleteFramebuffer calls the method "WebGL2RenderingContext.deleteFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteFramebuffer(framebuffer WebGLFramebuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasFuncDeleteProgram returns true if the method "WebGL2RenderingContext.deleteProgram" exists.
func (this WebGL2RenderingContext) HasFuncDeleteProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteProgram(
		this.ref,
	)
}

// FuncDeleteProgram returns the method "WebGL2RenderingContext.deleteProgram".
func (this WebGL2RenderingContext) FuncDeleteProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGL2RenderingContextDeleteProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteProgram calls the method "WebGL2RenderingContext.deleteProgram".
func (this WebGL2RenderingContext) DeleteProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryDeleteProgram calls the method "WebGL2RenderingContext.deleteProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncDeleteRenderbuffer returns true if the method "WebGL2RenderingContext.deleteRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFuncDeleteRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteRenderbuffer(
		this.ref,
	)
}

// FuncDeleteRenderbuffer returns the method "WebGL2RenderingContext.deleteRenderbuffer".
func (this WebGL2RenderingContext) FuncDeleteRenderbuffer() (fn js.Func[func(renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGL2RenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRenderbuffer calls the method "WebGL2RenderingContext.deleteRenderbuffer".
func (this WebGL2RenderingContext) DeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryDeleteRenderbuffer calls the method "WebGL2RenderingContext.deleteRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteRenderbuffer(renderbuffer WebGLRenderbuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncDeleteShader returns true if the method "WebGL2RenderingContext.deleteShader" exists.
func (this WebGL2RenderingContext) HasFuncDeleteShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteShader(
		this.ref,
	)
}

// FuncDeleteShader returns the method "WebGL2RenderingContext.deleteShader".
func (this WebGL2RenderingContext) FuncDeleteShader() (fn js.Func[func(shader WebGLShader)]) {
	bindings.FuncWebGL2RenderingContextDeleteShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteShader calls the method "WebGL2RenderingContext.deleteShader".
func (this WebGL2RenderingContext) DeleteShader(shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryDeleteShader calls the method "WebGL2RenderingContext.deleteShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteShader(shader WebGLShader) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncDeleteTexture returns true if the method "WebGL2RenderingContext.deleteTexture" exists.
func (this WebGL2RenderingContext) HasFuncDeleteTexture() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDeleteTexture(
		this.ref,
	)
}

// FuncDeleteTexture returns the method "WebGL2RenderingContext.deleteTexture".
func (this WebGL2RenderingContext) FuncDeleteTexture() (fn js.Func[func(texture WebGLTexture)]) {
	bindings.FuncWebGL2RenderingContextDeleteTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteTexture calls the method "WebGL2RenderingContext.deleteTexture".
func (this WebGL2RenderingContext) DeleteTexture(texture WebGLTexture) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDeleteTexture(
		this.ref, js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryDeleteTexture calls the method "WebGL2RenderingContext.deleteTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDeleteTexture(texture WebGLTexture) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDeleteTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasFuncDepthFunc returns true if the method "WebGL2RenderingContext.depthFunc" exists.
func (this WebGL2RenderingContext) HasFuncDepthFunc() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDepthFunc(
		this.ref,
	)
}

// FuncDepthFunc returns the method "WebGL2RenderingContext.depthFunc".
func (this WebGL2RenderingContext) FuncDepthFunc() (fn js.Func[func(fn GLenum)]) {
	bindings.FuncWebGL2RenderingContextDepthFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthFunc calls the method "WebGL2RenderingContext.depthFunc".
func (this WebGL2RenderingContext) DepthFunc(fn GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthFunc(
		this.ref, js.Pointer(&ret),
		uint32(fn),
	)

	return
}

// TryDepthFunc calls the method "WebGL2RenderingContext.depthFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDepthFunc(fn GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDepthFunc(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
	)

	return
}

// HasFuncDepthMask returns true if the method "WebGL2RenderingContext.depthMask" exists.
func (this WebGL2RenderingContext) HasFuncDepthMask() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDepthMask(
		this.ref,
	)
}

// FuncDepthMask returns the method "WebGL2RenderingContext.depthMask".
func (this WebGL2RenderingContext) FuncDepthMask() (fn js.Func[func(flag GLboolean)]) {
	bindings.FuncWebGL2RenderingContextDepthMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthMask calls the method "WebGL2RenderingContext.depthMask".
func (this WebGL2RenderingContext) DepthMask(flag GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthMask(
		this.ref, js.Pointer(&ret),
		js.Bool(bool(flag)),
	)

	return
}

// TryDepthMask calls the method "WebGL2RenderingContext.depthMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDepthMask(flag GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDepthMask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(flag)),
	)

	return
}

// HasFuncDepthRange returns true if the method "WebGL2RenderingContext.depthRange" exists.
func (this WebGL2RenderingContext) HasFuncDepthRange() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDepthRange(
		this.ref,
	)
}

// FuncDepthRange returns the method "WebGL2RenderingContext.depthRange".
func (this WebGL2RenderingContext) FuncDepthRange() (fn js.Func[func(zNear GLclampf, zFar GLclampf)]) {
	bindings.FuncWebGL2RenderingContextDepthRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DepthRange calls the method "WebGL2RenderingContext.depthRange".
func (this WebGL2RenderingContext) DepthRange(zNear GLclampf, zFar GLclampf) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDepthRange(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(zNear),
		float32(zFar),
	)

	return
}

// HasFuncDetachShader returns true if the method "WebGL2RenderingContext.detachShader" exists.
func (this WebGL2RenderingContext) HasFuncDetachShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDetachShader(
		this.ref,
	)
}

// FuncDetachShader returns the method "WebGL2RenderingContext.detachShader".
func (this WebGL2RenderingContext) FuncDetachShader() (fn js.Func[func(program WebGLProgram, shader WebGLShader)]) {
	bindings.FuncWebGL2RenderingContextDetachShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DetachShader calls the method "WebGL2RenderingContext.detachShader".
func (this WebGL2RenderingContext) DetachShader(program WebGLProgram, shader WebGLShader) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDetachShader(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		shader.Ref(),
	)

	return
}

// HasFuncDisable returns true if the method "WebGL2RenderingContext.disable" exists.
func (this WebGL2RenderingContext) HasFuncDisable() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDisable(
		this.ref,
	)
}

// FuncDisable returns the method "WebGL2RenderingContext.disable".
func (this WebGL2RenderingContext) FuncDisable() (fn js.Func[func(cap GLenum)]) {
	bindings.FuncWebGL2RenderingContextDisable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disable calls the method "WebGL2RenderingContext.disable".
func (this WebGL2RenderingContext) Disable(cap GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDisable(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryDisable calls the method "WebGL2RenderingContext.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDisable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDisable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncDisableVertexAttribArray returns true if the method "WebGL2RenderingContext.disableVertexAttribArray" exists.
func (this WebGL2RenderingContext) HasFuncDisableVertexAttribArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDisableVertexAttribArray(
		this.ref,
	)
}

// FuncDisableVertexAttribArray returns the method "WebGL2RenderingContext.disableVertexAttribArray".
func (this WebGL2RenderingContext) FuncDisableVertexAttribArray() (fn js.Func[func(index GLuint)]) {
	bindings.FuncWebGL2RenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DisableVertexAttribArray calls the method "WebGL2RenderingContext.disableVertexAttribArray".
func (this WebGL2RenderingContext) DisableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryDisableVertexAttribArray calls the method "WebGL2RenderingContext.disableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryDisableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextDisableVertexAttribArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncDrawArrays returns true if the method "WebGL2RenderingContext.drawArrays" exists.
func (this WebGL2RenderingContext) HasFuncDrawArrays() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawArrays(
		this.ref,
	)
}

// FuncDrawArrays returns the method "WebGL2RenderingContext.drawArrays".
func (this WebGL2RenderingContext) FuncDrawArrays() (fn js.Func[func(mode GLenum, first GLint, count GLsizei)]) {
	bindings.FuncWebGL2RenderingContextDrawArrays(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawArrays calls the method "WebGL2RenderingContext.drawArrays".
func (this WebGL2RenderingContext) DrawArrays(mode GLenum, first GLint, count GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawArrays(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
	)

	return
}

// HasFuncDrawElements returns true if the method "WebGL2RenderingContext.drawElements" exists.
func (this WebGL2RenderingContext) HasFuncDrawElements() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextDrawElements(
		this.ref,
	)
}

// FuncDrawElements returns the method "WebGL2RenderingContext.drawElements".
func (this WebGL2RenderingContext) FuncDrawElements() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextDrawElements(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawElements calls the method "WebGL2RenderingContext.drawElements".
func (this WebGL2RenderingContext) DrawElements(mode GLenum, count GLsizei, typ GLenum, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextDrawElements(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
	)

	return
}

// HasFuncEnable returns true if the method "WebGL2RenderingContext.enable" exists.
func (this WebGL2RenderingContext) HasFuncEnable() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextEnable(
		this.ref,
	)
}

// FuncEnable returns the method "WebGL2RenderingContext.enable".
func (this WebGL2RenderingContext) FuncEnable() (fn js.Func[func(cap GLenum)]) {
	bindings.FuncWebGL2RenderingContextEnable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Enable calls the method "WebGL2RenderingContext.enable".
func (this WebGL2RenderingContext) Enable(cap GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEnable(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryEnable calls the method "WebGL2RenderingContext.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEnable(cap GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEnable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncEnableVertexAttribArray returns true if the method "WebGL2RenderingContext.enableVertexAttribArray" exists.
func (this WebGL2RenderingContext) HasFuncEnableVertexAttribArray() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextEnableVertexAttribArray(
		this.ref,
	)
}

// FuncEnableVertexAttribArray returns the method "WebGL2RenderingContext.enableVertexAttribArray".
func (this WebGL2RenderingContext) FuncEnableVertexAttribArray() (fn js.Func[func(index GLuint)]) {
	bindings.FuncWebGL2RenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EnableVertexAttribArray calls the method "WebGL2RenderingContext.enableVertexAttribArray".
func (this WebGL2RenderingContext) EnableVertexAttribArray(index GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryEnableVertexAttribArray calls the method "WebGL2RenderingContext.enableVertexAttribArray"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryEnableVertexAttribArray(index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextEnableVertexAttribArray(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncFinish returns true if the method "WebGL2RenderingContext.finish" exists.
func (this WebGL2RenderingContext) HasFuncFinish() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFinish(
		this.ref,
	)
}

// FuncFinish returns the method "WebGL2RenderingContext.finish".
func (this WebGL2RenderingContext) FuncFinish() (fn js.Func[func()]) {
	bindings.FuncWebGL2RenderingContextFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "WebGL2RenderingContext.finish".
func (this WebGL2RenderingContext) Finish() (ret js.Void) {
	bindings.CallWebGL2RenderingContextFinish(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish calls the method "WebGL2RenderingContext.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFinish() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFlush returns true if the method "WebGL2RenderingContext.flush" exists.
func (this WebGL2RenderingContext) HasFuncFlush() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFlush(
		this.ref,
	)
}

// FuncFlush returns the method "WebGL2RenderingContext.flush".
func (this WebGL2RenderingContext) FuncFlush() (fn js.Func[func()]) {
	bindings.FuncWebGL2RenderingContextFlush(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Flush calls the method "WebGL2RenderingContext.flush".
func (this WebGL2RenderingContext) Flush() (ret js.Void) {
	bindings.CallWebGL2RenderingContextFlush(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "WebGL2RenderingContext.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFlush(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFramebufferRenderbuffer returns true if the method "WebGL2RenderingContext.framebufferRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFuncFramebufferRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFramebufferRenderbuffer(
		this.ref,
	)
}

// FuncFramebufferRenderbuffer returns the method "WebGL2RenderingContext.framebufferRenderbuffer".
func (this WebGL2RenderingContext) FuncFramebufferRenderbuffer() (fn js.Func[func(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer)]) {
	bindings.FuncWebGL2RenderingContextFramebufferRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FramebufferRenderbuffer calls the method "WebGL2RenderingContext.framebufferRenderbuffer".
func (this WebGL2RenderingContext) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum, renderbuffer WebGLRenderbuffer) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferRenderbuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(renderbuffertarget),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncFramebufferTexture2D returns true if the method "WebGL2RenderingContext.framebufferTexture2D" exists.
func (this WebGL2RenderingContext) HasFuncFramebufferTexture2D() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFramebufferTexture2D(
		this.ref,
	)
}

// FuncFramebufferTexture2D returns the method "WebGL2RenderingContext.framebufferTexture2D".
func (this WebGL2RenderingContext) FuncFramebufferTexture2D() (fn js.Func[func(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint)]) {
	bindings.FuncWebGL2RenderingContextFramebufferTexture2D(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FramebufferTexture2D calls the method "WebGL2RenderingContext.framebufferTexture2D".
func (this WebGL2RenderingContext) FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture WebGLTexture, level GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFramebufferTexture2D(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(textarget),
		texture.Ref(),
		int32(level),
	)

	return
}

// HasFuncFrontFace returns true if the method "WebGL2RenderingContext.frontFace" exists.
func (this WebGL2RenderingContext) HasFuncFrontFace() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextFrontFace(
		this.ref,
	)
}

// FuncFrontFace returns the method "WebGL2RenderingContext.frontFace".
func (this WebGL2RenderingContext) FuncFrontFace() (fn js.Func[func(mode GLenum)]) {
	bindings.FuncWebGL2RenderingContextFrontFace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FrontFace calls the method "WebGL2RenderingContext.frontFace".
func (this WebGL2RenderingContext) FrontFace(mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextFrontFace(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryFrontFace calls the method "WebGL2RenderingContext.frontFace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryFrontFace(mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextFrontFace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncGenerateMipmap returns true if the method "WebGL2RenderingContext.generateMipmap" exists.
func (this WebGL2RenderingContext) HasFuncGenerateMipmap() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGenerateMipmap(
		this.ref,
	)
}

// FuncGenerateMipmap returns the method "WebGL2RenderingContext.generateMipmap".
func (this WebGL2RenderingContext) FuncGenerateMipmap() (fn js.Func[func(target GLenum)]) {
	bindings.FuncWebGL2RenderingContextGenerateMipmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GenerateMipmap calls the method "WebGL2RenderingContext.generateMipmap".
func (this WebGL2RenderingContext) GenerateMipmap(target GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextGenerateMipmap(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryGenerateMipmap calls the method "WebGL2RenderingContext.generateMipmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGenerateMipmap(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGenerateMipmap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncGetActiveAttrib returns true if the method "WebGL2RenderingContext.getActiveAttrib" exists.
func (this WebGL2RenderingContext) HasFuncGetActiveAttrib() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetActiveAttrib(
		this.ref,
	)
}

// FuncGetActiveAttrib returns the method "WebGL2RenderingContext.getActiveAttrib".
func (this WebGL2RenderingContext) FuncGetActiveAttrib() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	bindings.FuncWebGL2RenderingContextGetActiveAttrib(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveAttrib calls the method "WebGL2RenderingContext.getActiveAttrib".
func (this WebGL2RenderingContext) GetActiveAttrib(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetActiveAttrib(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasFuncGetActiveUniform returns true if the method "WebGL2RenderingContext.getActiveUniform" exists.
func (this WebGL2RenderingContext) HasFuncGetActiveUniform() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetActiveUniform(
		this.ref,
	)
}

// FuncGetActiveUniform returns the method "WebGL2RenderingContext.getActiveUniform".
func (this WebGL2RenderingContext) FuncGetActiveUniform() (fn js.Func[func(program WebGLProgram, index GLuint) WebGLActiveInfo]) {
	bindings.FuncWebGL2RenderingContextGetActiveUniform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetActiveUniform calls the method "WebGL2RenderingContext.getActiveUniform".
func (this WebGL2RenderingContext) GetActiveUniform(program WebGLProgram, index GLuint) (ret WebGLActiveInfo) {
	bindings.CallWebGL2RenderingContextGetActiveUniform(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(index),
	)

	return
}

// HasFuncGetAttachedShaders returns true if the method "WebGL2RenderingContext.getAttachedShaders" exists.
func (this WebGL2RenderingContext) HasFuncGetAttachedShaders() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetAttachedShaders(
		this.ref,
	)
}

// FuncGetAttachedShaders returns the method "WebGL2RenderingContext.getAttachedShaders".
func (this WebGL2RenderingContext) FuncGetAttachedShaders() (fn js.Func[func(program WebGLProgram) js.Array[WebGLShader]]) {
	bindings.FuncWebGL2RenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttachedShaders calls the method "WebGL2RenderingContext.getAttachedShaders".
func (this WebGL2RenderingContext) GetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader]) {
	bindings.CallWebGL2RenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetAttachedShaders calls the method "WebGL2RenderingContext.getAttachedShaders"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetAttachedShaders(program WebGLProgram) (ret js.Array[WebGLShader], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetAttachedShaders(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncGetAttribLocation returns true if the method "WebGL2RenderingContext.getAttribLocation" exists.
func (this WebGL2RenderingContext) HasFuncGetAttribLocation() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetAttribLocation(
		this.ref,
	)
}

// FuncGetAttribLocation returns the method "WebGL2RenderingContext.getAttribLocation".
func (this WebGL2RenderingContext) FuncGetAttribLocation() (fn js.Func[func(program WebGLProgram, name js.String) GLint]) {
	bindings.FuncWebGL2RenderingContextGetAttribLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAttribLocation calls the method "WebGL2RenderingContext.getAttribLocation".
func (this WebGL2RenderingContext) GetAttribLocation(program WebGLProgram, name js.String) (ret GLint) {
	bindings.CallWebGL2RenderingContextGetAttribLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncGetBufferParameter returns true if the method "WebGL2RenderingContext.getBufferParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetBufferParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetBufferParameter(
		this.ref,
	)
}

// FuncGetBufferParameter returns the method "WebGL2RenderingContext.getBufferParameter".
func (this WebGL2RenderingContext) FuncGetBufferParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetBufferParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBufferParameter calls the method "WebGL2RenderingContext.getBufferParameter".
func (this WebGL2RenderingContext) GetBufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetBufferParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetParameter returns true if the method "WebGL2RenderingContext.getParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetParameter(
		this.ref,
	)
}

// FuncGetParameter returns the method "WebGL2RenderingContext.getParameter".
func (this WebGL2RenderingContext) FuncGetParameter() (fn js.Func[func(pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParameter calls the method "WebGL2RenderingContext.getParameter".
func (this WebGL2RenderingContext) GetParameter(pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetParameter(
		this.ref, js.Pointer(&ret),
		uint32(pname),
	)

	return
}

// TryGetParameter calls the method "WebGL2RenderingContext.getParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetParameter(pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetParameter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
	)

	return
}

// HasFuncGetError returns true if the method "WebGL2RenderingContext.getError" exists.
func (this WebGL2RenderingContext) HasFuncGetError() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetError(
		this.ref,
	)
}

// FuncGetError returns the method "WebGL2RenderingContext.getError".
func (this WebGL2RenderingContext) FuncGetError() (fn js.Func[func() GLenum]) {
	bindings.FuncWebGL2RenderingContextGetError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetError calls the method "WebGL2RenderingContext.getError".
func (this WebGL2RenderingContext) GetError() (ret GLenum) {
	bindings.CallWebGL2RenderingContextGetError(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetError calls the method "WebGL2RenderingContext.getError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetError() (ret GLenum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetFramebufferAttachmentParameter returns true if the method "WebGL2RenderingContext.getFramebufferAttachmentParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetFramebufferAttachmentParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.ref,
	)
}

// FuncGetFramebufferAttachmentParameter returns the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
func (this WebGL2RenderingContext) FuncGetFramebufferAttachmentParameter() (fn js.Func[func(target GLenum, attachment GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFramebufferAttachmentParameter calls the method "WebGL2RenderingContext.getFramebufferAttachmentParameter".
func (this WebGL2RenderingContext) GetFramebufferAttachmentParameter(target GLenum, attachment GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetFramebufferAttachmentParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		uint32(pname),
	)

	return
}

// HasFuncGetProgramParameter returns true if the method "WebGL2RenderingContext.getProgramParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetProgramParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetProgramParameter(
		this.ref,
	)
}

// FuncGetProgramParameter returns the method "WebGL2RenderingContext.getProgramParameter".
func (this WebGL2RenderingContext) FuncGetProgramParameter() (fn js.Func[func(program WebGLProgram, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetProgramParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetProgramParameter calls the method "WebGL2RenderingContext.getProgramParameter".
func (this WebGL2RenderingContext) GetProgramParameter(program WebGLProgram, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetProgramParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncGetProgramInfoLog returns true if the method "WebGL2RenderingContext.getProgramInfoLog" exists.
func (this WebGL2RenderingContext) HasFuncGetProgramInfoLog() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetProgramInfoLog(
		this.ref,
	)
}

// FuncGetProgramInfoLog returns the method "WebGL2RenderingContext.getProgramInfoLog".
func (this WebGL2RenderingContext) FuncGetProgramInfoLog() (fn js.Func[func(program WebGLProgram) js.String]) {
	bindings.FuncWebGL2RenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetProgramInfoLog calls the method "WebGL2RenderingContext.getProgramInfoLog".
func (this WebGL2RenderingContext) GetProgramInfoLog(program WebGLProgram) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryGetProgramInfoLog calls the method "WebGL2RenderingContext.getProgramInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetProgramInfoLog(program WebGLProgram) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetProgramInfoLog(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncGetRenderbufferParameter returns true if the method "WebGL2RenderingContext.getRenderbufferParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetRenderbufferParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetRenderbufferParameter(
		this.ref,
	)
}

// FuncGetRenderbufferParameter returns the method "WebGL2RenderingContext.getRenderbufferParameter".
func (this WebGL2RenderingContext) FuncGetRenderbufferParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetRenderbufferParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRenderbufferParameter calls the method "WebGL2RenderingContext.getRenderbufferParameter".
func (this WebGL2RenderingContext) GetRenderbufferParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetRenderbufferParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetShaderParameter returns true if the method "WebGL2RenderingContext.getShaderParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetShaderParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetShaderParameter(
		this.ref,
	)
}

// FuncGetShaderParameter returns the method "WebGL2RenderingContext.getShaderParameter".
func (this WebGL2RenderingContext) FuncGetShaderParameter() (fn js.Func[func(shader WebGLShader, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetShaderParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderParameter calls the method "WebGL2RenderingContext.getShaderParameter".
func (this WebGL2RenderingContext) GetShaderParameter(shader WebGLShader, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetShaderParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		uint32(pname),
	)

	return
}

// HasFuncGetShaderPrecisionFormat returns true if the method "WebGL2RenderingContext.getShaderPrecisionFormat" exists.
func (this WebGL2RenderingContext) HasFuncGetShaderPrecisionFormat() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetShaderPrecisionFormat(
		this.ref,
	)
}

// FuncGetShaderPrecisionFormat returns the method "WebGL2RenderingContext.getShaderPrecisionFormat".
func (this WebGL2RenderingContext) FuncGetShaderPrecisionFormat() (fn js.Func[func(shadertype GLenum, precisiontype GLenum) WebGLShaderPrecisionFormat]) {
	bindings.FuncWebGL2RenderingContextGetShaderPrecisionFormat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderPrecisionFormat calls the method "WebGL2RenderingContext.getShaderPrecisionFormat".
func (this WebGL2RenderingContext) GetShaderPrecisionFormat(shadertype GLenum, precisiontype GLenum) (ret WebGLShaderPrecisionFormat) {
	bindings.CallWebGL2RenderingContextGetShaderPrecisionFormat(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(shadertype),
		uint32(precisiontype),
	)

	return
}

// HasFuncGetShaderInfoLog returns true if the method "WebGL2RenderingContext.getShaderInfoLog" exists.
func (this WebGL2RenderingContext) HasFuncGetShaderInfoLog() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetShaderInfoLog(
		this.ref,
	)
}

// FuncGetShaderInfoLog returns the method "WebGL2RenderingContext.getShaderInfoLog".
func (this WebGL2RenderingContext) FuncGetShaderInfoLog() (fn js.Func[func(shader WebGLShader) js.String]) {
	bindings.FuncWebGL2RenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderInfoLog calls the method "WebGL2RenderingContext.getShaderInfoLog".
func (this WebGL2RenderingContext) GetShaderInfoLog(shader WebGLShader) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderInfoLog calls the method "WebGL2RenderingContext.getShaderInfoLog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderInfoLog(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderInfoLog(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncGetShaderSource returns true if the method "WebGL2RenderingContext.getShaderSource" exists.
func (this WebGL2RenderingContext) HasFuncGetShaderSource() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetShaderSource(
		this.ref,
	)
}

// FuncGetShaderSource returns the method "WebGL2RenderingContext.getShaderSource".
func (this WebGL2RenderingContext) FuncGetShaderSource() (fn js.Func[func(shader WebGLShader) js.String]) {
	bindings.FuncWebGL2RenderingContextGetShaderSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetShaderSource calls the method "WebGL2RenderingContext.getShaderSource".
func (this WebGL2RenderingContext) GetShaderSource(shader WebGLShader) (ret js.String) {
	bindings.CallWebGL2RenderingContextGetShaderSource(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetShaderSource calls the method "WebGL2RenderingContext.getShaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryGetShaderSource(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextGetShaderSource(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncGetTexParameter returns true if the method "WebGL2RenderingContext.getTexParameter" exists.
func (this WebGL2RenderingContext) HasFuncGetTexParameter() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetTexParameter(
		this.ref,
	)
}

// FuncGetTexParameter returns the method "WebGL2RenderingContext.getTexParameter".
func (this WebGL2RenderingContext) FuncGetTexParameter() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetTexParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTexParameter calls the method "WebGL2RenderingContext.getTexParameter".
func (this WebGL2RenderingContext) GetTexParameter(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetTexParameter(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetUniform returns true if the method "WebGL2RenderingContext.getUniform" exists.
func (this WebGL2RenderingContext) HasFuncGetUniform() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetUniform(
		this.ref,
	)
}

// FuncGetUniform returns the method "WebGL2RenderingContext.getUniform".
func (this WebGL2RenderingContext) FuncGetUniform() (fn js.Func[func(program WebGLProgram, location WebGLUniformLocation) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetUniform(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniform calls the method "WebGL2RenderingContext.getUniform".
func (this WebGL2RenderingContext) GetUniform(program WebGLProgram, location WebGLUniformLocation) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetUniform(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		location.Ref(),
	)

	return
}

// HasFuncGetUniformLocation returns true if the method "WebGL2RenderingContext.getUniformLocation" exists.
func (this WebGL2RenderingContext) HasFuncGetUniformLocation() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetUniformLocation(
		this.ref,
	)
}

// FuncGetUniformLocation returns the method "WebGL2RenderingContext.getUniformLocation".
func (this WebGL2RenderingContext) FuncGetUniformLocation() (fn js.Func[func(program WebGLProgram, name js.String) WebGLUniformLocation]) {
	bindings.FuncWebGL2RenderingContextGetUniformLocation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUniformLocation calls the method "WebGL2RenderingContext.getUniformLocation".
func (this WebGL2RenderingContext) GetUniformLocation(program WebGLProgram, name js.String) (ret WebGLUniformLocation) {
	bindings.CallWebGL2RenderingContextGetUniformLocation(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
		name.Ref(),
	)

	return
}

// HasFuncGetVertexAttrib returns true if the method "WebGL2RenderingContext.getVertexAttrib" exists.
func (this WebGL2RenderingContext) HasFuncGetVertexAttrib() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetVertexAttrib(
		this.ref,
	)
}

// FuncGetVertexAttrib returns the method "WebGL2RenderingContext.getVertexAttrib".
func (this WebGL2RenderingContext) FuncGetVertexAttrib() (fn js.Func[func(index GLuint, pname GLenum) js.Any]) {
	bindings.FuncWebGL2RenderingContextGetVertexAttrib(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetVertexAttrib calls the method "WebGL2RenderingContext.getVertexAttrib".
func (this WebGL2RenderingContext) GetVertexAttrib(index GLuint, pname GLenum) (ret js.Any) {
	bindings.CallWebGL2RenderingContextGetVertexAttrib(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasFuncGetVertexAttribOffset returns true if the method "WebGL2RenderingContext.getVertexAttribOffset" exists.
func (this WebGL2RenderingContext) HasFuncGetVertexAttribOffset() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextGetVertexAttribOffset(
		this.ref,
	)
}

// FuncGetVertexAttribOffset returns the method "WebGL2RenderingContext.getVertexAttribOffset".
func (this WebGL2RenderingContext) FuncGetVertexAttribOffset() (fn js.Func[func(index GLuint, pname GLenum) GLintptr]) {
	bindings.FuncWebGL2RenderingContextGetVertexAttribOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetVertexAttribOffset calls the method "WebGL2RenderingContext.getVertexAttribOffset".
func (this WebGL2RenderingContext) GetVertexAttribOffset(index GLuint, pname GLenum) (ret GLintptr) {
	bindings.CallWebGL2RenderingContextGetVertexAttribOffset(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(pname),
	)

	return
}

// HasFuncHint returns true if the method "WebGL2RenderingContext.hint" exists.
func (this WebGL2RenderingContext) HasFuncHint() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextHint(
		this.ref,
	)
}

// FuncHint returns the method "WebGL2RenderingContext.hint".
func (this WebGL2RenderingContext) FuncHint() (fn js.Func[func(target GLenum, mode GLenum)]) {
	bindings.FuncWebGL2RenderingContextHint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Hint calls the method "WebGL2RenderingContext.hint".
func (this WebGL2RenderingContext) Hint(target GLenum, mode GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextHint(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(mode),
	)

	return
}

// HasFuncIsBuffer returns true if the method "WebGL2RenderingContext.isBuffer" exists.
func (this WebGL2RenderingContext) HasFuncIsBuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsBuffer(
		this.ref,
	)
}

// FuncIsBuffer returns the method "WebGL2RenderingContext.isBuffer".
func (this WebGL2RenderingContext) FuncIsBuffer() (fn js.Func[func(buffer WebGLBuffer) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsBuffer calls the method "WebGL2RenderingContext.isBuffer".
func (this WebGL2RenderingContext) IsBuffer(buffer WebGLBuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsBuffer(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryIsBuffer calls the method "WebGL2RenderingContext.isBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsBuffer(buffer WebGLBuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncIsEnabled returns true if the method "WebGL2RenderingContext.isEnabled" exists.
func (this WebGL2RenderingContext) HasFuncIsEnabled() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsEnabled(
		this.ref,
	)
}

// FuncIsEnabled returns the method "WebGL2RenderingContext.isEnabled".
func (this WebGL2RenderingContext) FuncIsEnabled() (fn js.Func[func(cap GLenum) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsEnabled calls the method "WebGL2RenderingContext.isEnabled".
func (this WebGL2RenderingContext) IsEnabled(cap GLenum) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsEnabled(
		this.ref, js.Pointer(&ret),
		uint32(cap),
	)

	return
}

// TryIsEnabled calls the method "WebGL2RenderingContext.isEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsEnabled(cap GLenum) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(cap),
	)

	return
}

// HasFuncIsFramebuffer returns true if the method "WebGL2RenderingContext.isFramebuffer" exists.
func (this WebGL2RenderingContext) HasFuncIsFramebuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsFramebuffer(
		this.ref,
	)
}

// FuncIsFramebuffer returns the method "WebGL2RenderingContext.isFramebuffer".
func (this WebGL2RenderingContext) FuncIsFramebuffer() (fn js.Func[func(framebuffer WebGLFramebuffer) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsFramebuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsFramebuffer calls the method "WebGL2RenderingContext.isFramebuffer".
func (this WebGL2RenderingContext) IsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsFramebuffer(
		this.ref, js.Pointer(&ret),
		framebuffer.Ref(),
	)

	return
}

// TryIsFramebuffer calls the method "WebGL2RenderingContext.isFramebuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsFramebuffer(framebuffer WebGLFramebuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsFramebuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		framebuffer.Ref(),
	)

	return
}

// HasFuncIsProgram returns true if the method "WebGL2RenderingContext.isProgram" exists.
func (this WebGL2RenderingContext) HasFuncIsProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsProgram(
		this.ref,
	)
}

// FuncIsProgram returns the method "WebGL2RenderingContext.isProgram".
func (this WebGL2RenderingContext) FuncIsProgram() (fn js.Func[func(program WebGLProgram) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsProgram calls the method "WebGL2RenderingContext.isProgram".
func (this WebGL2RenderingContext) IsProgram(program WebGLProgram) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryIsProgram calls the method "WebGL2RenderingContext.isProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsProgram(program WebGLProgram) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncIsRenderbuffer returns true if the method "WebGL2RenderingContext.isRenderbuffer" exists.
func (this WebGL2RenderingContext) HasFuncIsRenderbuffer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsRenderbuffer(
		this.ref,
	)
}

// FuncIsRenderbuffer returns the method "WebGL2RenderingContext.isRenderbuffer".
func (this WebGL2RenderingContext) FuncIsRenderbuffer() (fn js.Func[func(renderbuffer WebGLRenderbuffer) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsRenderbuffer calls the method "WebGL2RenderingContext.isRenderbuffer".
func (this WebGL2RenderingContext) IsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&ret),
		renderbuffer.Ref(),
	)

	return
}

// TryIsRenderbuffer calls the method "WebGL2RenderingContext.isRenderbuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsRenderbuffer(renderbuffer WebGLRenderbuffer) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsRenderbuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		renderbuffer.Ref(),
	)

	return
}

// HasFuncIsShader returns true if the method "WebGL2RenderingContext.isShader" exists.
func (this WebGL2RenderingContext) HasFuncIsShader() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsShader(
		this.ref,
	)
}

// FuncIsShader returns the method "WebGL2RenderingContext.isShader".
func (this WebGL2RenderingContext) FuncIsShader() (fn js.Func[func(shader WebGLShader) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsShader(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsShader calls the method "WebGL2RenderingContext.isShader".
func (this WebGL2RenderingContext) IsShader(shader WebGLShader) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsShader(
		this.ref, js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryIsShader calls the method "WebGL2RenderingContext.isShader"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsShader(shader WebGLShader) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsShader(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

// HasFuncIsTexture returns true if the method "WebGL2RenderingContext.isTexture" exists.
func (this WebGL2RenderingContext) HasFuncIsTexture() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextIsTexture(
		this.ref,
	)
}

// FuncIsTexture returns the method "WebGL2RenderingContext.isTexture".
func (this WebGL2RenderingContext) FuncIsTexture() (fn js.Func[func(texture WebGLTexture) GLboolean]) {
	bindings.FuncWebGL2RenderingContextIsTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTexture calls the method "WebGL2RenderingContext.isTexture".
func (this WebGL2RenderingContext) IsTexture(texture WebGLTexture) (ret GLboolean) {
	bindings.CallWebGL2RenderingContextIsTexture(
		this.ref, js.Pointer(&ret),
		texture.Ref(),
	)

	return
}

// TryIsTexture calls the method "WebGL2RenderingContext.isTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryIsTexture(texture WebGLTexture) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextIsTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		texture.Ref(),
	)

	return
}

// HasFuncLineWidth returns true if the method "WebGL2RenderingContext.lineWidth" exists.
func (this WebGL2RenderingContext) HasFuncLineWidth() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextLineWidth(
		this.ref,
	)
}

// FuncLineWidth returns the method "WebGL2RenderingContext.lineWidth".
func (this WebGL2RenderingContext) FuncLineWidth() (fn js.Func[func(width GLfloat)]) {
	bindings.FuncWebGL2RenderingContextLineWidth(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LineWidth calls the method "WebGL2RenderingContext.lineWidth".
func (this WebGL2RenderingContext) LineWidth(width GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextLineWidth(
		this.ref, js.Pointer(&ret),
		float32(width),
	)

	return
}

// TryLineWidth calls the method "WebGL2RenderingContext.lineWidth"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryLineWidth(width GLfloat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextLineWidth(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(width),
	)

	return
}

// HasFuncLinkProgram returns true if the method "WebGL2RenderingContext.linkProgram" exists.
func (this WebGL2RenderingContext) HasFuncLinkProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextLinkProgram(
		this.ref,
	)
}

// FuncLinkProgram returns the method "WebGL2RenderingContext.linkProgram".
func (this WebGL2RenderingContext) FuncLinkProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGL2RenderingContextLinkProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LinkProgram calls the method "WebGL2RenderingContext.linkProgram".
func (this WebGL2RenderingContext) LinkProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextLinkProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryLinkProgram calls the method "WebGL2RenderingContext.linkProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryLinkProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextLinkProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncPixelStorei returns true if the method "WebGL2RenderingContext.pixelStorei" exists.
func (this WebGL2RenderingContext) HasFuncPixelStorei() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextPixelStorei(
		this.ref,
	)
}

// FuncPixelStorei returns the method "WebGL2RenderingContext.pixelStorei".
func (this WebGL2RenderingContext) FuncPixelStorei() (fn js.Func[func(pname GLenum, param GLint)]) {
	bindings.FuncWebGL2RenderingContextPixelStorei(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PixelStorei calls the method "WebGL2RenderingContext.pixelStorei".
func (this WebGL2RenderingContext) PixelStorei(pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextPixelStorei(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(pname),
		int32(param),
	)

	return
}

// HasFuncPolygonOffset returns true if the method "WebGL2RenderingContext.polygonOffset" exists.
func (this WebGL2RenderingContext) HasFuncPolygonOffset() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextPolygonOffset(
		this.ref,
	)
}

// FuncPolygonOffset returns the method "WebGL2RenderingContext.polygonOffset".
func (this WebGL2RenderingContext) FuncPolygonOffset() (fn js.Func[func(factor GLfloat, units GLfloat)]) {
	bindings.FuncWebGL2RenderingContextPolygonOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PolygonOffset calls the method "WebGL2RenderingContext.polygonOffset".
func (this WebGL2RenderingContext) PolygonOffset(factor GLfloat, units GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextPolygonOffset(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(factor),
		float32(units),
	)

	return
}

// HasFuncRenderbufferStorage returns true if the method "WebGL2RenderingContext.renderbufferStorage" exists.
func (this WebGL2RenderingContext) HasFuncRenderbufferStorage() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextRenderbufferStorage(
		this.ref,
	)
}

// FuncRenderbufferStorage returns the method "WebGL2RenderingContext.renderbufferStorage".
func (this WebGL2RenderingContext) FuncRenderbufferStorage() (fn js.Func[func(target GLenum, internalformat GLenum, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextRenderbufferStorage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RenderbufferStorage calls the method "WebGL2RenderingContext.renderbufferStorage".
func (this WebGL2RenderingContext) RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextRenderbufferStorage(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(internalformat),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncSampleCoverage returns true if the method "WebGL2RenderingContext.sampleCoverage" exists.
func (this WebGL2RenderingContext) HasFuncSampleCoverage() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextSampleCoverage(
		this.ref,
	)
}

// FuncSampleCoverage returns the method "WebGL2RenderingContext.sampleCoverage".
func (this WebGL2RenderingContext) FuncSampleCoverage() (fn js.Func[func(value GLclampf, invert GLboolean)]) {
	bindings.FuncWebGL2RenderingContextSampleCoverage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SampleCoverage calls the method "WebGL2RenderingContext.sampleCoverage".
func (this WebGL2RenderingContext) SampleCoverage(value GLclampf, invert GLboolean) (ret js.Void) {
	bindings.CallWebGL2RenderingContextSampleCoverage(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(value),
		js.Bool(bool(invert)),
	)

	return
}

// HasFuncScissor returns true if the method "WebGL2RenderingContext.scissor" exists.
func (this WebGL2RenderingContext) HasFuncScissor() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextScissor(
		this.ref,
	)
}

// FuncScissor returns the method "WebGL2RenderingContext.scissor".
func (this WebGL2RenderingContext) FuncScissor() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextScissor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scissor calls the method "WebGL2RenderingContext.scissor".
func (this WebGL2RenderingContext) Scissor(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextScissor(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncShaderSource returns true if the method "WebGL2RenderingContext.shaderSource" exists.
func (this WebGL2RenderingContext) HasFuncShaderSource() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextShaderSource(
		this.ref,
	)
}

// FuncShaderSource returns the method "WebGL2RenderingContext.shaderSource".
func (this WebGL2RenderingContext) FuncShaderSource() (fn js.Func[func(shader WebGLShader, source js.String)]) {
	bindings.FuncWebGL2RenderingContextShaderSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShaderSource calls the method "WebGL2RenderingContext.shaderSource".
func (this WebGL2RenderingContext) ShaderSource(shader WebGLShader, source js.String) (ret js.Void) {
	bindings.CallWebGL2RenderingContextShaderSource(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
		source.Ref(),
	)

	return
}

// HasFuncStencilFunc returns true if the method "WebGL2RenderingContext.stencilFunc" exists.
func (this WebGL2RenderingContext) HasFuncStencilFunc() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilFunc(
		this.ref,
	)
}

// FuncStencilFunc returns the method "WebGL2RenderingContext.stencilFunc".
func (this WebGL2RenderingContext) FuncStencilFunc() (fn js.Func[func(fn GLenum, ref GLint, mask GLuint)]) {
	bindings.FuncWebGL2RenderingContextStencilFunc(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilFunc calls the method "WebGL2RenderingContext.stencilFunc".
func (this WebGL2RenderingContext) StencilFunc(fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilFunc(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasFuncStencilFuncSeparate returns true if the method "WebGL2RenderingContext.stencilFuncSeparate" exists.
func (this WebGL2RenderingContext) HasFuncStencilFuncSeparate() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilFuncSeparate(
		this.ref,
	)
}

// FuncStencilFuncSeparate returns the method "WebGL2RenderingContext.stencilFuncSeparate".
func (this WebGL2RenderingContext) FuncStencilFuncSeparate() (fn js.Func[func(face GLenum, fn GLenum, ref GLint, mask GLuint)]) {
	bindings.FuncWebGL2RenderingContextStencilFuncSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilFuncSeparate calls the method "WebGL2RenderingContext.stencilFuncSeparate".
func (this WebGL2RenderingContext) StencilFuncSeparate(face GLenum, fn GLenum, ref GLint, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilFuncSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fn),
		int32(ref),
		uint32(mask),
	)

	return
}

// HasFuncStencilMask returns true if the method "WebGL2RenderingContext.stencilMask" exists.
func (this WebGL2RenderingContext) HasFuncStencilMask() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilMask(
		this.ref,
	)
}

// FuncStencilMask returns the method "WebGL2RenderingContext.stencilMask".
func (this WebGL2RenderingContext) FuncStencilMask() (fn js.Func[func(mask GLuint)]) {
	bindings.FuncWebGL2RenderingContextStencilMask(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilMask calls the method "WebGL2RenderingContext.stencilMask".
func (this WebGL2RenderingContext) StencilMask(mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilMask(
		this.ref, js.Pointer(&ret),
		uint32(mask),
	)

	return
}

// TryStencilMask calls the method "WebGL2RenderingContext.stencilMask"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryStencilMask(mask GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextStencilMask(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mask),
	)

	return
}

// HasFuncStencilMaskSeparate returns true if the method "WebGL2RenderingContext.stencilMaskSeparate" exists.
func (this WebGL2RenderingContext) HasFuncStencilMaskSeparate() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilMaskSeparate(
		this.ref,
	)
}

// FuncStencilMaskSeparate returns the method "WebGL2RenderingContext.stencilMaskSeparate".
func (this WebGL2RenderingContext) FuncStencilMaskSeparate() (fn js.Func[func(face GLenum, mask GLuint)]) {
	bindings.FuncWebGL2RenderingContextStencilMaskSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilMaskSeparate calls the method "WebGL2RenderingContext.stencilMaskSeparate".
func (this WebGL2RenderingContext) StencilMaskSeparate(face GLenum, mask GLuint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilMaskSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(mask),
	)

	return
}

// HasFuncStencilOp returns true if the method "WebGL2RenderingContext.stencilOp" exists.
func (this WebGL2RenderingContext) HasFuncStencilOp() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilOp(
		this.ref,
	)
}

// FuncStencilOp returns the method "WebGL2RenderingContext.stencilOp".
func (this WebGL2RenderingContext) FuncStencilOp() (fn js.Func[func(fail GLenum, zfail GLenum, zpass GLenum)]) {
	bindings.FuncWebGL2RenderingContextStencilOp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilOp calls the method "WebGL2RenderingContext.stencilOp".
func (this WebGL2RenderingContext) StencilOp(fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilOp(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasFuncStencilOpSeparate returns true if the method "WebGL2RenderingContext.stencilOpSeparate" exists.
func (this WebGL2RenderingContext) HasFuncStencilOpSeparate() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextStencilOpSeparate(
		this.ref,
	)
}

// FuncStencilOpSeparate returns the method "WebGL2RenderingContext.stencilOpSeparate".
func (this WebGL2RenderingContext) FuncStencilOpSeparate() (fn js.Func[func(face GLenum, fail GLenum, zfail GLenum, zpass GLenum)]) {
	bindings.FuncWebGL2RenderingContextStencilOpSeparate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StencilOpSeparate calls the method "WebGL2RenderingContext.stencilOpSeparate".
func (this WebGL2RenderingContext) StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) (ret js.Void) {
	bindings.CallWebGL2RenderingContextStencilOpSeparate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(face),
		uint32(fail),
		uint32(zfail),
		uint32(zpass),
	)

	return
}

// HasFuncTexParameterf returns true if the method "WebGL2RenderingContext.texParameterf" exists.
func (this WebGL2RenderingContext) HasFuncTexParameterf() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexParameterf(
		this.ref,
	)
}

// FuncTexParameterf returns the method "WebGL2RenderingContext.texParameterf".
func (this WebGL2RenderingContext) FuncTexParameterf() (fn js.Func[func(target GLenum, pname GLenum, param GLfloat)]) {
	bindings.FuncWebGL2RenderingContextTexParameterf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexParameterf calls the method "WebGL2RenderingContext.texParameterf".
func (this WebGL2RenderingContext) TexParameterf(target GLenum, pname GLenum, param GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexParameterf(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		float32(param),
	)

	return
}

// HasFuncTexParameteri returns true if the method "WebGL2RenderingContext.texParameteri" exists.
func (this WebGL2RenderingContext) HasFuncTexParameteri() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextTexParameteri(
		this.ref,
	)
}

// FuncTexParameteri returns the method "WebGL2RenderingContext.texParameteri".
func (this WebGL2RenderingContext) FuncTexParameteri() (fn js.Func[func(target GLenum, pname GLenum, param GLint)]) {
	bindings.FuncWebGL2RenderingContextTexParameteri(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TexParameteri calls the method "WebGL2RenderingContext.texParameteri".
func (this WebGL2RenderingContext) TexParameteri(target GLenum, pname GLenum, param GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextTexParameteri(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
		int32(param),
	)

	return
}

// HasFuncUniform1f returns true if the method "WebGL2RenderingContext.uniform1f" exists.
func (this WebGL2RenderingContext) HasFuncUniform1f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1f(
		this.ref,
	)
}

// FuncUniform1f returns the method "WebGL2RenderingContext.uniform1f".
func (this WebGL2RenderingContext) FuncUniform1f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat)]) {
	bindings.FuncWebGL2RenderingContextUniform1f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1f calls the method "WebGL2RenderingContext.uniform1f".
func (this WebGL2RenderingContext) Uniform1f(location WebGLUniformLocation, x GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
	)

	return
}

// HasFuncUniform2f returns true if the method "WebGL2RenderingContext.uniform2f" exists.
func (this WebGL2RenderingContext) HasFuncUniform2f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2f(
		this.ref,
	)
}

// FuncUniform2f returns the method "WebGL2RenderingContext.uniform2f".
func (this WebGL2RenderingContext) FuncUniform2f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat)]) {
	bindings.FuncWebGL2RenderingContextUniform2f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2f calls the method "WebGL2RenderingContext.uniform2f".
func (this WebGL2RenderingContext) Uniform2f(location WebGLUniformLocation, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
	)

	return
}

// HasFuncUniform3f returns true if the method "WebGL2RenderingContext.uniform3f" exists.
func (this WebGL2RenderingContext) HasFuncUniform3f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3f(
		this.ref,
	)
}

// FuncUniform3f returns the method "WebGL2RenderingContext.uniform3f".
func (this WebGL2RenderingContext) FuncUniform3f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat)]) {
	bindings.FuncWebGL2RenderingContextUniform3f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3f calls the method "WebGL2RenderingContext.uniform3f".
func (this WebGL2RenderingContext) Uniform3f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncUniform4f returns true if the method "WebGL2RenderingContext.uniform4f" exists.
func (this WebGL2RenderingContext) HasFuncUniform4f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4f(
		this.ref,
	)
}

// FuncUniform4f returns the method "WebGL2RenderingContext.uniform4f".
func (this WebGL2RenderingContext) FuncUniform4f() (fn js.Func[func(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	bindings.FuncWebGL2RenderingContextUniform4f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4f calls the method "WebGL2RenderingContext.uniform4f".
func (this WebGL2RenderingContext) Uniform4f(location WebGLUniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasFuncUniform1i returns true if the method "WebGL2RenderingContext.uniform1i" exists.
func (this WebGL2RenderingContext) HasFuncUniform1i() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform1i(
		this.ref,
	)
}

// FuncUniform1i returns the method "WebGL2RenderingContext.uniform1i".
func (this WebGL2RenderingContext) FuncUniform1i() (fn js.Func[func(location WebGLUniformLocation, x GLint)]) {
	bindings.FuncWebGL2RenderingContextUniform1i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform1i calls the method "WebGL2RenderingContext.uniform1i".
func (this WebGL2RenderingContext) Uniform1i(location WebGLUniformLocation, x GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform1i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
	)

	return
}

// HasFuncUniform2i returns true if the method "WebGL2RenderingContext.uniform2i" exists.
func (this WebGL2RenderingContext) HasFuncUniform2i() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform2i(
		this.ref,
	)
}

// FuncUniform2i returns the method "WebGL2RenderingContext.uniform2i".
func (this WebGL2RenderingContext) FuncUniform2i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint)]) {
	bindings.FuncWebGL2RenderingContextUniform2i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform2i calls the method "WebGL2RenderingContext.uniform2i".
func (this WebGL2RenderingContext) Uniform2i(location WebGLUniformLocation, x GLint, y GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform2i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncUniform3i returns true if the method "WebGL2RenderingContext.uniform3i" exists.
func (this WebGL2RenderingContext) HasFuncUniform3i() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform3i(
		this.ref,
	)
}

// FuncUniform3i returns the method "WebGL2RenderingContext.uniform3i".
func (this WebGL2RenderingContext) FuncUniform3i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint)]) {
	bindings.FuncWebGL2RenderingContextUniform3i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform3i calls the method "WebGL2RenderingContext.uniform3i".
func (this WebGL2RenderingContext) Uniform3i(location WebGLUniformLocation, x GLint, y GLint, z GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform3i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
	)

	return
}

// HasFuncUniform4i returns true if the method "WebGL2RenderingContext.uniform4i" exists.
func (this WebGL2RenderingContext) HasFuncUniform4i() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUniform4i(
		this.ref,
	)
}

// FuncUniform4i returns the method "WebGL2RenderingContext.uniform4i".
func (this WebGL2RenderingContext) FuncUniform4i() (fn js.Func[func(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint)]) {
	bindings.FuncWebGL2RenderingContextUniform4i(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Uniform4i calls the method "WebGL2RenderingContext.uniform4i".
func (this WebGL2RenderingContext) Uniform4i(location WebGLUniformLocation, x GLint, y GLint, z GLint, w GLint) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUniform4i(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		location.Ref(),
		int32(x),
		int32(y),
		int32(z),
		int32(w),
	)

	return
}

// HasFuncUseProgram returns true if the method "WebGL2RenderingContext.useProgram" exists.
func (this WebGL2RenderingContext) HasFuncUseProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextUseProgram(
		this.ref,
	)
}

// FuncUseProgram returns the method "WebGL2RenderingContext.useProgram".
func (this WebGL2RenderingContext) FuncUseProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGL2RenderingContextUseProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UseProgram calls the method "WebGL2RenderingContext.useProgram".
func (this WebGL2RenderingContext) UseProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextUseProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryUseProgram calls the method "WebGL2RenderingContext.useProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryUseProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextUseProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncValidateProgram returns true if the method "WebGL2RenderingContext.validateProgram" exists.
func (this WebGL2RenderingContext) HasFuncValidateProgram() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextValidateProgram(
		this.ref,
	)
}

// FuncValidateProgram returns the method "WebGL2RenderingContext.validateProgram".
func (this WebGL2RenderingContext) FuncValidateProgram() (fn js.Func[func(program WebGLProgram)]) {
	bindings.FuncWebGL2RenderingContextValidateProgram(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ValidateProgram calls the method "WebGL2RenderingContext.validateProgram".
func (this WebGL2RenderingContext) ValidateProgram(program WebGLProgram) (ret js.Void) {
	bindings.CallWebGL2RenderingContextValidateProgram(
		this.ref, js.Pointer(&ret),
		program.Ref(),
	)

	return
}

// TryValidateProgram calls the method "WebGL2RenderingContext.validateProgram"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryValidateProgram(program WebGLProgram) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextValidateProgram(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		program.Ref(),
	)

	return
}

// HasFuncVertexAttrib1f returns true if the method "WebGL2RenderingContext.vertexAttrib1f" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib1f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib1f(
		this.ref,
	)
}

// FuncVertexAttrib1f returns the method "WebGL2RenderingContext.vertexAttrib1f".
func (this WebGL2RenderingContext) FuncVertexAttrib1f() (fn js.Func[func(index GLuint, x GLfloat)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib1f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib1f calls the method "WebGL2RenderingContext.vertexAttrib1f".
func (this WebGL2RenderingContext) VertexAttrib1f(index GLuint, x GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib1f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
	)

	return
}

// HasFuncVertexAttrib2f returns true if the method "WebGL2RenderingContext.vertexAttrib2f" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib2f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib2f(
		this.ref,
	)
}

// FuncVertexAttrib2f returns the method "WebGL2RenderingContext.vertexAttrib2f".
func (this WebGL2RenderingContext) FuncVertexAttrib2f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib2f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib2f calls the method "WebGL2RenderingContext.vertexAttrib2f".
func (this WebGL2RenderingContext) VertexAttrib2f(index GLuint, x GLfloat, y GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib2f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
	)

	return
}

// HasFuncVertexAttrib3f returns true if the method "WebGL2RenderingContext.vertexAttrib3f" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib3f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib3f(
		this.ref,
	)
}

// FuncVertexAttrib3f returns the method "WebGL2RenderingContext.vertexAttrib3f".
func (this WebGL2RenderingContext) FuncVertexAttrib3f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib3f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib3f calls the method "WebGL2RenderingContext.vertexAttrib3f".
func (this WebGL2RenderingContext) VertexAttrib3f(index GLuint, x GLfloat, y GLfloat, z GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib3f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncVertexAttrib4f returns true if the method "WebGL2RenderingContext.vertexAttrib4f" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib4f() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib4f(
		this.ref,
	)
}

// FuncVertexAttrib4f returns the method "WebGL2RenderingContext.vertexAttrib4f".
func (this WebGL2RenderingContext) FuncVertexAttrib4f() (fn js.Func[func(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib4f(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib4f calls the method "WebGL2RenderingContext.vertexAttrib4f".
func (this WebGL2RenderingContext) VertexAttrib4f(index GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib4f(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		float32(x),
		float32(y),
		float32(z),
		float32(w),
	)

	return
}

// HasFuncVertexAttrib1fv returns true if the method "WebGL2RenderingContext.vertexAttrib1fv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib1fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib1fv(
		this.ref,
	)
}

// FuncVertexAttrib1fv returns the method "WebGL2RenderingContext.vertexAttrib1fv".
func (this WebGL2RenderingContext) FuncVertexAttrib1fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib1fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib1fv calls the method "WebGL2RenderingContext.vertexAttrib1fv".
func (this WebGL2RenderingContext) VertexAttrib1fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib1fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib2fv returns true if the method "WebGL2RenderingContext.vertexAttrib2fv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib2fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib2fv(
		this.ref,
	)
}

// FuncVertexAttrib2fv returns the method "WebGL2RenderingContext.vertexAttrib2fv".
func (this WebGL2RenderingContext) FuncVertexAttrib2fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib2fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib2fv calls the method "WebGL2RenderingContext.vertexAttrib2fv".
func (this WebGL2RenderingContext) VertexAttrib2fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib2fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib3fv returns true if the method "WebGL2RenderingContext.vertexAttrib3fv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib3fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib3fv(
		this.ref,
	)
}

// FuncVertexAttrib3fv returns the method "WebGL2RenderingContext.vertexAttrib3fv".
func (this WebGL2RenderingContext) FuncVertexAttrib3fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib3fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib3fv calls the method "WebGL2RenderingContext.vertexAttrib3fv".
func (this WebGL2RenderingContext) VertexAttrib3fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib3fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttrib4fv returns true if the method "WebGL2RenderingContext.vertexAttrib4fv" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttrib4fv() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttrib4fv(
		this.ref,
	)
}

// FuncVertexAttrib4fv returns the method "WebGL2RenderingContext.vertexAttrib4fv".
func (this WebGL2RenderingContext) FuncVertexAttrib4fv() (fn js.Func[func(index GLuint, values Float32List)]) {
	bindings.FuncWebGL2RenderingContextVertexAttrib4fv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttrib4fv calls the method "WebGL2RenderingContext.vertexAttrib4fv".
func (this WebGL2RenderingContext) VertexAttrib4fv(index GLuint, values Float32List) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttrib4fv(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		values.Ref(),
	)

	return
}

// HasFuncVertexAttribPointer returns true if the method "WebGL2RenderingContext.vertexAttribPointer" exists.
func (this WebGL2RenderingContext) HasFuncVertexAttribPointer() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextVertexAttribPointer(
		this.ref,
	)
}

// FuncVertexAttribPointer returns the method "WebGL2RenderingContext.vertexAttribPointer".
func (this WebGL2RenderingContext) FuncVertexAttribPointer() (fn js.Func[func(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr)]) {
	bindings.FuncWebGL2RenderingContextVertexAttribPointer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribPointer calls the method "WebGL2RenderingContext.vertexAttribPointer".
func (this WebGL2RenderingContext) VertexAttribPointer(index GLuint, size GLint, typ GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) (ret js.Void) {
	bindings.CallWebGL2RenderingContextVertexAttribPointer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		int32(size),
		uint32(typ),
		js.Bool(bool(normalized)),
		int32(stride),
		float64(offset),
	)

	return
}

// HasFuncViewport returns true if the method "WebGL2RenderingContext.viewport" exists.
func (this WebGL2RenderingContext) HasFuncViewport() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextViewport(
		this.ref,
	)
}

// FuncViewport returns the method "WebGL2RenderingContext.viewport".
func (this WebGL2RenderingContext) FuncViewport() (fn js.Func[func(x GLint, y GLint, width GLsizei, height GLsizei)]) {
	bindings.FuncWebGL2RenderingContextViewport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Viewport calls the method "WebGL2RenderingContext.viewport".
func (this WebGL2RenderingContext) Viewport(x GLint, y GLint, width GLsizei, height GLsizei) (ret js.Void) {
	bindings.CallWebGL2RenderingContextViewport(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		int32(width),
		int32(height),
	)

	return
}

// HasFuncMakeXRCompatible returns true if the method "WebGL2RenderingContext.makeXRCompatible" exists.
func (this WebGL2RenderingContext) HasFuncMakeXRCompatible() bool {
	return js.True == bindings.HasFuncWebGL2RenderingContextMakeXRCompatible(
		this.ref,
	)
}

// FuncMakeXRCompatible returns the method "WebGL2RenderingContext.makeXRCompatible".
func (this WebGL2RenderingContext) FuncMakeXRCompatible() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncWebGL2RenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MakeXRCompatible calls the method "WebGL2RenderingContext.makeXRCompatible".
func (this WebGL2RenderingContext) MakeXRCompatible() (ret js.Promise[js.Void]) {
	bindings.CallWebGL2RenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMakeXRCompatible calls the method "WebGL2RenderingContext.makeXRCompatible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WebGL2RenderingContext) TryMakeXRCompatible() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebGL2RenderingContextMakeXRCompatible(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Size returns the value of property "GPUBuffer.size".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) Size() (ret GPUSize64Out, ok bool) {
	ok = js.True == bindings.GetGPUBufferSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Usage returns the value of property "GPUBuffer.usage".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) Usage() (ret GPUFlagsConstant, ok bool) {
	ok = js.True == bindings.GetGPUBufferUsage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MapState returns the value of property "GPUBuffer.mapState".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) MapState() (ret GPUBufferMapState, ok bool) {
	ok = js.True == bindings.GetGPUBufferMapState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUBuffer.label".
//
// It returns ok=false if there is no such property.
func (this GPUBuffer) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUBufferLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUBuffer.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBuffer) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBufferLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncMapAsync returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasFuncMapAsync() bool {
	return js.True == bindings.HasFuncGPUBufferMapAsync(
		this.ref,
	)
}

// FuncMapAsync returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) FuncMapAsync() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) js.Promise[js.Void]]) {
	bindings.FuncGPUBufferMapAsync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MapAsync calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync(mode GPUMapModeFlags, offset GPUSize64, size GPUSize64) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncMapAsync1 returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasFuncMapAsync1() bool {
	return js.True == bindings.HasFuncGPUBufferMapAsync1(
		this.ref,
	)
}

// FuncMapAsync1 returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) FuncMapAsync1() (fn js.Func[func(mode GPUMapModeFlags, offset GPUSize64) js.Promise[js.Void]]) {
	bindings.FuncGPUBufferMapAsync1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MapAsync1 calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync1(mode GPUMapModeFlags, offset GPUSize64) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		float64(offset),
	)

	return
}

// HasFuncMapAsync2 returns true if the method "GPUBuffer.mapAsync" exists.
func (this GPUBuffer) HasFuncMapAsync2() bool {
	return js.True == bindings.HasFuncGPUBufferMapAsync2(
		this.ref,
	)
}

// FuncMapAsync2 returns the method "GPUBuffer.mapAsync".
func (this GPUBuffer) FuncMapAsync2() (fn js.Func[func(mode GPUMapModeFlags) js.Promise[js.Void]]) {
	bindings.FuncGPUBufferMapAsync2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MapAsync2 calls the method "GPUBuffer.mapAsync".
func (this GPUBuffer) MapAsync2(mode GPUMapModeFlags) (ret js.Promise[js.Void]) {
	bindings.CallGPUBufferMapAsync2(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryMapAsync2 calls the method "GPUBuffer.mapAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryMapAsync2(mode GPUMapModeFlags) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferMapAsync2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncGetMappedRange returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasFuncGetMappedRange() bool {
	return js.True == bindings.HasFuncGPUBufferGetMappedRange(
		this.ref,
	)
}

// FuncGetMappedRange returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) FuncGetMappedRange() (fn js.Func[func(offset GPUSize64, size GPUSize64) js.ArrayBuffer]) {
	bindings.FuncGPUBufferGetMappedRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetMappedRange calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange(offset GPUSize64, size GPUSize64) (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncGetMappedRange1 returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasFuncGetMappedRange1() bool {
	return js.True == bindings.HasFuncGPUBufferGetMappedRange1(
		this.ref,
	)
}

// FuncGetMappedRange1 returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) FuncGetMappedRange1() (fn js.Func[func(offset GPUSize64) js.ArrayBuffer]) {
	bindings.FuncGPUBufferGetMappedRange1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetMappedRange1 calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange1(offset GPUSize64) (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange1(
		this.ref, js.Pointer(&ret),
		float64(offset),
	)

	return
}

// TryGetMappedRange1 calls the method "GPUBuffer.getMappedRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryGetMappedRange1(offset GPUSize64) (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferGetMappedRange1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(offset),
	)

	return
}

// HasFuncGetMappedRange2 returns true if the method "GPUBuffer.getMappedRange" exists.
func (this GPUBuffer) HasFuncGetMappedRange2() bool {
	return js.True == bindings.HasFuncGPUBufferGetMappedRange2(
		this.ref,
	)
}

// FuncGetMappedRange2 returns the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) FuncGetMappedRange2() (fn js.Func[func() js.ArrayBuffer]) {
	bindings.FuncGPUBufferGetMappedRange2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetMappedRange2 calls the method "GPUBuffer.getMappedRange".
func (this GPUBuffer) GetMappedRange2() (ret js.ArrayBuffer) {
	bindings.CallGPUBufferGetMappedRange2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetMappedRange2 calls the method "GPUBuffer.getMappedRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryGetMappedRange2() (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferGetMappedRange2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUnmap returns true if the method "GPUBuffer.unmap" exists.
func (this GPUBuffer) HasFuncUnmap() bool {
	return js.True == bindings.HasFuncGPUBufferUnmap(
		this.ref,
	)
}

// FuncUnmap returns the method "GPUBuffer.unmap".
func (this GPUBuffer) FuncUnmap() (fn js.Func[func()]) {
	bindings.FuncGPUBufferUnmap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Unmap calls the method "GPUBuffer.unmap".
func (this GPUBuffer) Unmap() (ret js.Void) {
	bindings.CallGPUBufferUnmap(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUnmap calls the method "GPUBuffer.unmap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryUnmap() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferUnmap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDestroy returns true if the method "GPUBuffer.destroy" exists.
func (this GPUBuffer) HasFuncDestroy() bool {
	return js.True == bindings.HasFuncGPUBufferDestroy(
		this.ref,
	)
}

// FuncDestroy returns the method "GPUBuffer.destroy".
func (this GPUBuffer) FuncDestroy() (fn js.Func[func()]) {
	bindings.FuncGPUBufferDestroy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Destroy calls the method "GPUBuffer.destroy".
func (this GPUBuffer) Destroy() (ret js.Void) {
	bindings.CallGPUBufferDestroy(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUBuffer.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUBuffer) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUBufferDestroy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPUBufferDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBufferDescriptor) Update(ref js.Ref) {
	bindings.GPUBufferDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBufferDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUTextureView struct {
	ref js.Ref
}

func (this GPUTextureView) Once() GPUTextureView {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUTextureView.label".
//
// It returns ok=false if there is no such property.
func (this GPUTextureView) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUTextureViewLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUTextureView.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUTextureView) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUTextureViewLabel(
		this.ref,
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
func (p *GPUTextureViewDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureViewDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUTextureViewDescriptor) Update(ref js.Ref) {
	bindings.GPUTextureViewDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUTextureViewDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "GPUTexture.width".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Width() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "GPUTexture.height".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Height() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthOrArrayLayers returns the value of property "GPUTexture.depthOrArrayLayers".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) DepthOrArrayLayers() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureDepthOrArrayLayers(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MipLevelCount returns the value of property "GPUTexture.mipLevelCount".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) MipLevelCount() (ret GPUIntegerCoordinateOut, ok bool) {
	ok = js.True == bindings.GetGPUTextureMipLevelCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SampleCount returns the value of property "GPUTexture.sampleCount".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) SampleCount() (ret GPUSize32Out, ok bool) {
	ok = js.True == bindings.GetGPUTextureSampleCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dimension returns the value of property "GPUTexture.dimension".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Dimension() (ret GPUTextureDimension, ok bool) {
	ok = js.True == bindings.GetGPUTextureDimension(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Format returns the value of property "GPUTexture.format".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Format() (ret GPUTextureFormat, ok bool) {
	ok = js.True == bindings.GetGPUTextureFormat(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Usage returns the value of property "GPUTexture.usage".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Usage() (ret GPUFlagsConstant, ok bool) {
	ok = js.True == bindings.GetGPUTextureUsage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUTexture.label".
//
// It returns ok=false if there is no such property.
func (this GPUTexture) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUTextureLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUTexture.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUTexture) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUTextureLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCreateView returns true if the method "GPUTexture.createView" exists.
func (this GPUTexture) HasFuncCreateView() bool {
	return js.True == bindings.HasFuncGPUTextureCreateView(
		this.ref,
	)
}

// FuncCreateView returns the method "GPUTexture.createView".
func (this GPUTexture) FuncCreateView() (fn js.Func[func(descriptor GPUTextureViewDescriptor) GPUTextureView]) {
	bindings.FuncGPUTextureCreateView(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateView calls the method "GPUTexture.createView".
func (this GPUTexture) CreateView(descriptor GPUTextureViewDescriptor) (ret GPUTextureView) {
	bindings.CallGPUTextureCreateView(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateView calls the method "GPUTexture.createView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryCreateView(descriptor GPUTextureViewDescriptor) (ret GPUTextureView, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureCreateView(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateView1 returns true if the method "GPUTexture.createView" exists.
func (this GPUTexture) HasFuncCreateView1() bool {
	return js.True == bindings.HasFuncGPUTextureCreateView1(
		this.ref,
	)
}

// FuncCreateView1 returns the method "GPUTexture.createView".
func (this GPUTexture) FuncCreateView1() (fn js.Func[func() GPUTextureView]) {
	bindings.FuncGPUTextureCreateView1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateView1 calls the method "GPUTexture.createView".
func (this GPUTexture) CreateView1() (ret GPUTextureView) {
	bindings.CallGPUTextureCreateView1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateView1 calls the method "GPUTexture.createView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryCreateView1() (ret GPUTextureView, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureCreateView1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDestroy returns true if the method "GPUTexture.destroy" exists.
func (this GPUTexture) HasFuncDestroy() bool {
	return js.True == bindings.HasFuncGPUTextureDestroy(
		this.ref,
	)
}

// FuncDestroy returns the method "GPUTexture.destroy".
func (this GPUTexture) FuncDestroy() (fn js.Func[func()]) {
	bindings.FuncGPUTextureDestroy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Destroy calls the method "GPUTexture.destroy".
func (this GPUTexture) Destroy() (ret js.Void) {
	bindings.CallGPUTextureDestroy(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUTexture.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUTexture) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUTextureDestroy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
