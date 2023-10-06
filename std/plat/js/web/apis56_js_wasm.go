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

type VideoEncoderEncodeOptionsForAv1 struct {
	// Quantizer is "VideoEncoderEncodeOptionsForAv1.quantizer"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quantizer MUST be set to true to make this field effective.
	Quantizer uint16

	FFI_USE_Quantizer bool // for Quantizer.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderEncodeOptionsForAv1 with all fields set.
func (p VideoEncoderEncodeOptionsForAv1) FromRef(ref js.Ref) VideoEncoderEncodeOptionsForAv1 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderEncodeOptionsForAv1 in the application heap.
func (p VideoEncoderEncodeOptionsForAv1) New() js.Ref {
	return bindings.VideoEncoderEncodeOptionsForAv1JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderEncodeOptionsForAv1) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForAv1JSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderEncodeOptionsForAv1) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForAv1JSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderEncodeOptionsForAvc struct {
	// Quantizer is "VideoEncoderEncodeOptionsForAvc.quantizer"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quantizer MUST be set to true to make this field effective.
	Quantizer uint16

	FFI_USE_Quantizer bool // for Quantizer.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderEncodeOptionsForAvc with all fields set.
func (p VideoEncoderEncodeOptionsForAvc) FromRef(ref js.Ref) VideoEncoderEncodeOptionsForAvc {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderEncodeOptionsForAvc in the application heap.
func (p VideoEncoderEncodeOptionsForAvc) New() js.Ref {
	return bindings.VideoEncoderEncodeOptionsForAvcJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderEncodeOptionsForAvc) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForAvcJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderEncodeOptionsForAvc) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForAvcJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderEncodeOptions struct {
	// KeyFrame is "VideoEncoderEncodeOptions.keyFrame"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_KeyFrame MUST be set to true to make this field effective.
	KeyFrame bool
	// Hevc is "VideoEncoderEncodeOptions.hevc"
	//
	// Optional
	//
	// NOTE: Hevc.FFI_USE MUST be set to true to get Hevc used.
	Hevc VideoEncoderEncodeOptionsForHevc
	// Vp9 is "VideoEncoderEncodeOptions.vp9"
	//
	// Optional
	//
	// NOTE: Vp9.FFI_USE MUST be set to true to get Vp9 used.
	Vp9 VideoEncoderEncodeOptionsForVp9
	// Av1 is "VideoEncoderEncodeOptions.av1"
	//
	// Optional
	//
	// NOTE: Av1.FFI_USE MUST be set to true to get Av1 used.
	Av1 VideoEncoderEncodeOptionsForAv1
	// Avc is "VideoEncoderEncodeOptions.avc"
	//
	// Optional
	//
	// NOTE: Avc.FFI_USE MUST be set to true to get Avc used.
	Avc VideoEncoderEncodeOptionsForAvc

	FFI_USE_KeyFrame bool // for KeyFrame.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderEncodeOptions with all fields set.
func (p VideoEncoderEncodeOptions) FromRef(ref js.Ref) VideoEncoderEncodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderEncodeOptions in the application heap.
func (p VideoEncoderEncodeOptions) New() js.Ref {
	return bindings.VideoEncoderEncodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderEncodeOptions) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderEncodeOptions) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderSupport struct {
	// Supported is "VideoEncoderSupport.supported"
	//
	// Optional
	//
	// NOTE: FFI_USE_Supported MUST be set to true to make this field effective.
	Supported bool
	// Config is "VideoEncoderSupport.config"
	//
	// Optional
	//
	// NOTE: Config.FFI_USE MUST be set to true to get Config used.
	Config VideoEncoderConfig

	FFI_USE_Supported bool // for Supported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderSupport with all fields set.
func (p VideoEncoderSupport) FromRef(ref js.Ref) VideoEncoderSupport {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderSupport in the application heap.
func (p VideoEncoderSupport) New() js.Ref {
	return bindings.VideoEncoderSupportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderSupport) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderSupportJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderSupport) Update(ref js.Ref) {
	bindings.VideoEncoderSupportJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewVideoEncoder(init VideoEncoderInit) (ret VideoEncoder) {
	ret.ref = bindings.NewVideoEncoderByVideoEncoder(
		js.Pointer(&init))
	return
}

type VideoEncoder struct {
	EventTarget
}

func (this VideoEncoder) Once() VideoEncoder {
	this.Ref().Once()
	return this
}

func (this VideoEncoder) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this VideoEncoder) FromRef(ref js.Ref) VideoEncoder {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this VideoEncoder) Free() {
	this.Ref().Free()
}

// State returns the value of property "VideoEncoder.state".
//
// It returns ok=false if there is no such property.
func (this VideoEncoder) State() (ret CodecState, ok bool) {
	ok = js.True == bindings.GetVideoEncoderState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EncodeQueueSize returns the value of property "VideoEncoder.encodeQueueSize".
//
// It returns ok=false if there is no such property.
func (this VideoEncoder) EncodeQueueSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoEncoderEncodeQueueSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasConfigure returns true if the method "VideoEncoder.configure" exists.
func (this VideoEncoder) HasConfigure() bool {
	return js.True == bindings.HasVideoEncoderConfigure(
		this.Ref(),
	)
}

// ConfigureFunc returns the method "VideoEncoder.configure".
func (this VideoEncoder) ConfigureFunc() (fn js.Func[func(config VideoEncoderConfig)]) {
	return fn.FromRef(
		bindings.VideoEncoderConfigureFunc(
			this.Ref(),
		),
	)
}

// Configure calls the method "VideoEncoder.configure".
func (this VideoEncoder) Configure(config VideoEncoderConfig) (ret js.Void) {
	bindings.CallVideoEncoderConfigure(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryConfigure calls the method "VideoEncoder.configure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryConfigure(config VideoEncoderConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderConfigure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasEncode returns true if the method "VideoEncoder.encode" exists.
func (this VideoEncoder) HasEncode() bool {
	return js.True == bindings.HasVideoEncoderEncode(
		this.Ref(),
	)
}

// EncodeFunc returns the method "VideoEncoder.encode".
func (this VideoEncoder) EncodeFunc() (fn js.Func[func(frame VideoFrame, options VideoEncoderEncodeOptions)]) {
	return fn.FromRef(
		bindings.VideoEncoderEncodeFunc(
			this.Ref(),
		),
	)
}

// Encode calls the method "VideoEncoder.encode".
func (this VideoEncoder) Encode(frame VideoFrame, options VideoEncoderEncodeOptions) (ret js.Void) {
	bindings.CallVideoEncoderEncode(
		this.Ref(), js.Pointer(&ret),
		frame.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryEncode calls the method "VideoEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryEncode(frame VideoFrame, options VideoEncoderEncodeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderEncode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		frame.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasEncode1 returns true if the method "VideoEncoder.encode" exists.
func (this VideoEncoder) HasEncode1() bool {
	return js.True == bindings.HasVideoEncoderEncode1(
		this.Ref(),
	)
}

// Encode1Func returns the method "VideoEncoder.encode".
func (this VideoEncoder) Encode1Func() (fn js.Func[func(frame VideoFrame)]) {
	return fn.FromRef(
		bindings.VideoEncoderEncode1Func(
			this.Ref(),
		),
	)
}

// Encode1 calls the method "VideoEncoder.encode".
func (this VideoEncoder) Encode1(frame VideoFrame) (ret js.Void) {
	bindings.CallVideoEncoderEncode1(
		this.Ref(), js.Pointer(&ret),
		frame.Ref(),
	)

	return
}

// TryEncode1 calls the method "VideoEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryEncode1(frame VideoFrame) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderEncode1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		frame.Ref(),
	)

	return
}

// HasFlush returns true if the method "VideoEncoder.flush" exists.
func (this VideoEncoder) HasFlush() bool {
	return js.True == bindings.HasVideoEncoderFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "VideoEncoder.flush".
func (this VideoEncoder) FlushFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.VideoEncoderFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "VideoEncoder.flush".
func (this VideoEncoder) Flush() (ret js.Promise[js.Void]) {
	bindings.CallVideoEncoderFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "VideoEncoder.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryFlush() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "VideoEncoder.reset" exists.
func (this VideoEncoder) HasReset() bool {
	return js.True == bindings.HasVideoEncoderReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "VideoEncoder.reset".
func (this VideoEncoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoEncoderResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "VideoEncoder.reset".
func (this VideoEncoder) Reset() (ret js.Void) {
	bindings.CallVideoEncoderReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "VideoEncoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "VideoEncoder.close" exists.
func (this VideoEncoder) HasClose() bool {
	return js.True == bindings.HasVideoEncoderClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "VideoEncoder.close".
func (this VideoEncoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoEncoderCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "VideoEncoder.close".
func (this VideoEncoder) Close() (ret js.Void) {
	bindings.CallVideoEncoderClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "VideoEncoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsConfigSupported returns true if the staticmethod "VideoEncoder.isConfigSupported" exists.
func (this VideoEncoder) HasIsConfigSupported() bool {
	return js.True == bindings.HasVideoEncoderIsConfigSupported(
		this.Ref(),
	)
}

// IsConfigSupportedFunc returns the staticmethod "VideoEncoder.isConfigSupported".
func (this VideoEncoder) IsConfigSupportedFunc() (fn js.Func[func(config VideoEncoderConfig) js.Promise[VideoEncoderSupport]]) {
	return fn.FromRef(
		bindings.VideoEncoderIsConfigSupportedFunc(
			this.Ref(),
		),
	)
}

// IsConfigSupported calls the staticmethod "VideoEncoder.isConfigSupported".
func (this VideoEncoder) IsConfigSupported(config VideoEncoderConfig) (ret js.Promise[VideoEncoderSupport]) {
	bindings.CallVideoEncoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryIsConfigSupported calls the staticmethod "VideoEncoder.isConfigSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoEncoder) TryIsConfigSupported(config VideoEncoderConfig) (ret js.Promise[VideoEncoderSupport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoEncoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

type VideoFacingModeEnum uint32

const (
	_ VideoFacingModeEnum = iota

	VideoFacingModeEnum_USER
	VideoFacingModeEnum_ENVIRONMENT
	VideoFacingModeEnum_LEFT
	VideoFacingModeEnum_RIGHT
)

func (VideoFacingModeEnum) FromRef(str js.Ref) VideoFacingModeEnum {
	return VideoFacingModeEnum(bindings.ConstOfVideoFacingModeEnum(str))
}

func (x VideoFacingModeEnum) String() (string, bool) {
	switch x {
	case VideoFacingModeEnum_USER:
		return "user", true
	case VideoFacingModeEnum_ENVIRONMENT:
		return "environment", true
	case VideoFacingModeEnum_LEFT:
		return "left", true
	case VideoFacingModeEnum_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type VideoResizeModeEnum uint32

const (
	_ VideoResizeModeEnum = iota

	VideoResizeModeEnum_NONE
	VideoResizeModeEnum_CROP_AND_SCALE
)

func (VideoResizeModeEnum) FromRef(str js.Ref) VideoResizeModeEnum {
	return VideoResizeModeEnum(bindings.ConstOfVideoResizeModeEnum(str))
}

func (x VideoResizeModeEnum) String() (string, bool) {
	switch x {
	case VideoResizeModeEnum_NONE:
		return "none", true
	case VideoResizeModeEnum_CROP_AND_SCALE:
		return "crop-and-scale", true
	default:
		return "", false
	}
}

type VideoTrackGenerator struct {
	ref js.Ref
}

func (this VideoTrackGenerator) Once() VideoTrackGenerator {
	this.Ref().Once()
	return this
}

func (this VideoTrackGenerator) Ref() js.Ref {
	return this.ref
}

func (this VideoTrackGenerator) FromRef(ref js.Ref) VideoTrackGenerator {
	this.ref = ref
	return this
}

func (this VideoTrackGenerator) Free() {
	this.Ref().Free()
}

// Writable returns the value of property "VideoTrackGenerator.writable".
//
// It returns ok=false if there is no such property.
func (this VideoTrackGenerator) Writable() (ret WritableStream, ok bool) {
	ok = js.True == bindings.GetVideoTrackGeneratorWritable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Muted returns the value of property "VideoTrackGenerator.muted".
//
// It returns ok=false if there is no such property.
func (this VideoTrackGenerator) Muted() (ret bool, ok bool) {
	ok = js.True == bindings.GetVideoTrackGeneratorMuted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMuted sets the value of property "VideoTrackGenerator.muted" to val.
//
// It returns false if the property cannot be set.
func (this VideoTrackGenerator) SetMuted(val bool) bool {
	return js.True == bindings.SetVideoTrackGeneratorMuted(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Track returns the value of property "VideoTrackGenerator.track".
//
// It returns ok=false if there is no such property.
func (this VideoTrackGenerator) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetVideoTrackGeneratorTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_CSSNumericValue_CSSKeywordValue struct {
	ref js.Ref
}

func (x OneOf_CSSNumericValue_CSSKeywordValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_CSSNumericValue_CSSKeywordValue) Free() {
	x.ref.Free()
}

func (x OneOf_CSSNumericValue_CSSKeywordValue) FromRef(ref js.Ref) OneOf_CSSNumericValue_CSSKeywordValue {
	return OneOf_CSSNumericValue_CSSKeywordValue{
		ref: ref,
	}
}

func (x OneOf_CSSNumericValue_CSSKeywordValue) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

func (x OneOf_CSSNumericValue_CSSKeywordValue) CSSKeywordValue() CSSKeywordValue {
	return CSSKeywordValue{}.FromRef(x.ref)
}

type OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue struct {
	ref js.Ref
}

func (x OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue) FromRef(ref js.Ref) OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue {
	return OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue{
		ref: ref,
	}
}

func (x OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue) ArrayOneOf_CSSNumericValue_CSSKeywordValue() js.Array[OneOf_CSSNumericValue_CSSKeywordValue] {
	return js.Array[OneOf_CSSNumericValue_CSSKeywordValue]{}.FromRef(x.ref)
}

type ViewTimelineOptions struct {
	// Subject is "ViewTimelineOptions.subject"
	//
	// Optional
	Subject Element
	// Axis is "ViewTimelineOptions.axis"
	//
	// Optional, defaults to "block".
	Axis ScrollAxis
	// Inset is "ViewTimelineOptions.inset"
	//
	// Optional, defaults to "auto".
	Inset OneOf_String_ArrayOneOf_CSSNumericValue_CSSKeywordValue

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ViewTimelineOptions with all fields set.
func (p ViewTimelineOptions) FromRef(ref js.Ref) ViewTimelineOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ViewTimelineOptions in the application heap.
func (p ViewTimelineOptions) New() js.Ref {
	return bindings.ViewTimelineOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ViewTimelineOptions) UpdateFrom(ref js.Ref) {
	bindings.ViewTimelineOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ViewTimelineOptions) Update(ref js.Ref) {
	bindings.ViewTimelineOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewViewTimeline(options ViewTimelineOptions) (ret ViewTimeline) {
	ret.ref = bindings.NewViewTimelineByViewTimeline(
		js.Pointer(&options))
	return
}

func NewViewTimelineByViewTimeline1() (ret ViewTimeline) {
	ret.ref = bindings.NewViewTimelineByViewTimeline1()
	return
}

type ViewTimeline struct {
	ScrollTimeline
}

func (this ViewTimeline) Once() ViewTimeline {
	this.Ref().Once()
	return this
}

func (this ViewTimeline) Ref() js.Ref {
	return this.ScrollTimeline.Ref()
}

func (this ViewTimeline) FromRef(ref js.Ref) ViewTimeline {
	this.ScrollTimeline = this.ScrollTimeline.FromRef(ref)
	return this
}

func (this ViewTimeline) Free() {
	this.Ref().Free()
}

// Subject returns the value of property "ViewTimeline.subject".
//
// It returns ok=false if there is no such property.
func (this ViewTimeline) Subject() (ret Element, ok bool) {
	ok = js.True == bindings.GetViewTimelineSubject(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StartOffset returns the value of property "ViewTimeline.startOffset".
//
// It returns ok=false if there is no such property.
func (this ViewTimeline) StartOffset() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetViewTimelineStartOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EndOffset returns the value of property "ViewTimeline.endOffset".
//
// It returns ok=false if there is no such property.
func (this ViewTimeline) EndOffset() (ret CSSNumericValue, ok bool) {
	ok = js.True == bindings.GetViewTimelineEndOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type VisibilityStateEntry struct {
	PerformanceEntry
}

func (this VisibilityStateEntry) Once() VisibilityStateEntry {
	this.Ref().Once()
	return this
}

func (this VisibilityStateEntry) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this VisibilityStateEntry) FromRef(ref js.Ref) VisibilityStateEntry {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this VisibilityStateEntry) Free() {
	this.Ref().Free()
}

// Name returns the value of property "VisibilityStateEntry.name".
//
// It returns ok=false if there is no such property.
func (this VisibilityStateEntry) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVisibilityStateEntryName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EntryType returns the value of property "VisibilityStateEntry.entryType".
//
// It returns ok=false if there is no such property.
func (this VisibilityStateEntry) EntryType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVisibilityStateEntryEntryType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StartTime returns the value of property "VisibilityStateEntry.startTime".
//
// It returns ok=false if there is no such property.
func (this VisibilityStateEntry) StartTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetVisibilityStateEntryStartTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "VisibilityStateEntry.duration".
//
// It returns ok=false if there is no such property.
func (this VisibilityStateEntry) Duration() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVisibilityStateEntryDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

const (
	WEBGL_blend_equation_advanced_coherent_MULTIPLY       GLenum = 0x9294
	WEBGL_blend_equation_advanced_coherent_SCREEN         GLenum = 0x9295
	WEBGL_blend_equation_advanced_coherent_OVERLAY        GLenum = 0x9296
	WEBGL_blend_equation_advanced_coherent_DARKEN         GLenum = 0x9297
	WEBGL_blend_equation_advanced_coherent_LIGHTEN        GLenum = 0x9298
	WEBGL_blend_equation_advanced_coherent_COLORDODGE     GLenum = 0x9299
	WEBGL_blend_equation_advanced_coherent_COLORBURN      GLenum = 0x929A
	WEBGL_blend_equation_advanced_coherent_HARDLIGHT      GLenum = 0x929B
	WEBGL_blend_equation_advanced_coherent_SOFTLIGHT      GLenum = 0x929C
	WEBGL_blend_equation_advanced_coherent_DIFFERENCE     GLenum = 0x929E
	WEBGL_blend_equation_advanced_coherent_EXCLUSION      GLenum = 0x92A0
	WEBGL_blend_equation_advanced_coherent_HSL_HUE        GLenum = 0x92AD
	WEBGL_blend_equation_advanced_coherent_HSL_SATURATION GLenum = 0x92AE
	WEBGL_blend_equation_advanced_coherent_HSL_COLOR      GLenum = 0x92AF
	WEBGL_blend_equation_advanced_coherent_HSL_LUMINOSITY GLenum = 0x92B0
)

type WEBGL_blend_equation_advanced_coherent struct {
	ref js.Ref
}

func (this WEBGL_blend_equation_advanced_coherent) Once() WEBGL_blend_equation_advanced_coherent {
	this.Ref().Once()
	return this
}

func (this WEBGL_blend_equation_advanced_coherent) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_blend_equation_advanced_coherent) FromRef(ref js.Ref) WEBGL_blend_equation_advanced_coherent {
	this.ref = ref
	return this
}

func (this WEBGL_blend_equation_advanced_coherent) Free() {
	this.Ref().Free()
}

const (
	WEBGL_clip_cull_distance_MAX_CLIP_DISTANCES_WEBGL                   GLenum = 0x0D32
	WEBGL_clip_cull_distance_MAX_CULL_DISTANCES_WEBGL                   GLenum = 0x82F9
	WEBGL_clip_cull_distance_MAX_COMBINED_CLIP_AND_CULL_DISTANCES_WEBGL GLenum = 0x82FA
	WEBGL_clip_cull_distance_CLIP_DISTANCE0_WEBGL                       GLenum = 0x3000
	WEBGL_clip_cull_distance_CLIP_DISTANCE1_WEBGL                       GLenum = 0x3001
	WEBGL_clip_cull_distance_CLIP_DISTANCE2_WEBGL                       GLenum = 0x3002
	WEBGL_clip_cull_distance_CLIP_DISTANCE3_WEBGL                       GLenum = 0x3003
	WEBGL_clip_cull_distance_CLIP_DISTANCE4_WEBGL                       GLenum = 0x3004
	WEBGL_clip_cull_distance_CLIP_DISTANCE5_WEBGL                       GLenum = 0x3005
	WEBGL_clip_cull_distance_CLIP_DISTANCE6_WEBGL                       GLenum = 0x3006
	WEBGL_clip_cull_distance_CLIP_DISTANCE7_WEBGL                       GLenum = 0x3007
)

type WEBGL_clip_cull_distance struct {
	ref js.Ref
}

func (this WEBGL_clip_cull_distance) Once() WEBGL_clip_cull_distance {
	this.Ref().Once()
	return this
}

func (this WEBGL_clip_cull_distance) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_clip_cull_distance) FromRef(ref js.Ref) WEBGL_clip_cull_distance {
	this.ref = ref
	return this
}

func (this WEBGL_clip_cull_distance) Free() {
	this.Ref().Free()
}

const (
	WEBGL_color_buffer_float_RGBA32F_EXT                               GLenum = 0x8814
	WEBGL_color_buffer_float_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT GLenum = 0x8211
	WEBGL_color_buffer_float_UNSIGNED_NORMALIZED_EXT                   GLenum = 0x8C17
)

type WEBGL_color_buffer_float struct {
	ref js.Ref
}

func (this WEBGL_color_buffer_float) Once() WEBGL_color_buffer_float {
	this.Ref().Once()
	return this
}

func (this WEBGL_color_buffer_float) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_color_buffer_float) FromRef(ref js.Ref) WEBGL_color_buffer_float {
	this.ref = ref
	return this
}

func (this WEBGL_color_buffer_float) Free() {
	this.Ref().Free()
}

const (
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_4x4_KHR           GLenum = 0x93B0
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_5x4_KHR           GLenum = 0x93B1
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_5x5_KHR           GLenum = 0x93B2
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_6x5_KHR           GLenum = 0x93B3
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_6x6_KHR           GLenum = 0x93B4
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_8x5_KHR           GLenum = 0x93B5
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_8x6_KHR           GLenum = 0x93B6
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_8x8_KHR           GLenum = 0x93B7
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_10x5_KHR          GLenum = 0x93B8
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_10x6_KHR          GLenum = 0x93B9
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_10x8_KHR          GLenum = 0x93BA
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_10x10_KHR         GLenum = 0x93BB
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_12x10_KHR         GLenum = 0x93BC
	WEBGL_compressed_texture_astc_COMPRESSED_RGBA_ASTC_12x12_KHR         GLenum = 0x93BD
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR   GLenum = 0x93D0
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR   GLenum = 0x93D1
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR   GLenum = 0x93D2
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR   GLenum = 0x93D3
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR   GLenum = 0x93D4
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR   GLenum = 0x93D5
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR   GLenum = 0x93D6
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR   GLenum = 0x93D7
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR  GLenum = 0x93D8
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR  GLenum = 0x93D9
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR  GLenum = 0x93DA
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR GLenum = 0x93DB
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR GLenum = 0x93DC
	WEBGL_compressed_texture_astc_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR GLenum = 0x93DD
)

type WEBGL_compressed_texture_astc struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_astc) Once() WEBGL_compressed_texture_astc {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_astc) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_astc) FromRef(ref js.Ref) WEBGL_compressed_texture_astc {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_astc) Free() {
	this.Ref().Free()
}

// HasGetSupportedProfiles returns true if the method "WEBGL_compressed_texture_astc.getSupportedProfiles" exists.
func (this WEBGL_compressed_texture_astc) HasGetSupportedProfiles() bool {
	return js.True == bindings.HasWEBGL_compressed_texture_astcGetSupportedProfiles(
		this.Ref(),
	)
}

// GetSupportedProfilesFunc returns the method "WEBGL_compressed_texture_astc.getSupportedProfiles".
func (this WEBGL_compressed_texture_astc) GetSupportedProfilesFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.WEBGL_compressed_texture_astcGetSupportedProfilesFunc(
			this.Ref(),
		),
	)
}

// GetSupportedProfiles calls the method "WEBGL_compressed_texture_astc.getSupportedProfiles".
func (this WEBGL_compressed_texture_astc) GetSupportedProfiles() (ret js.Array[js.String]) {
	bindings.CallWEBGL_compressed_texture_astcGetSupportedProfiles(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedProfiles calls the method "WEBGL_compressed_texture_astc.getSupportedProfiles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_compressed_texture_astc) TryGetSupportedProfiles() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_compressed_texture_astcGetSupportedProfiles(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

const (
	WEBGL_compressed_texture_etc_COMPRESSED_R11_EAC                        GLenum = 0x9270
	WEBGL_compressed_texture_etc_COMPRESSED_SIGNED_R11_EAC                 GLenum = 0x9271
	WEBGL_compressed_texture_etc_COMPRESSED_RG11_EAC                       GLenum = 0x9272
	WEBGL_compressed_texture_etc_COMPRESSED_SIGNED_RG11_EAC                GLenum = 0x9273
	WEBGL_compressed_texture_etc_COMPRESSED_RGB8_ETC2                      GLenum = 0x9274
	WEBGL_compressed_texture_etc_COMPRESSED_SRGB8_ETC2                     GLenum = 0x9275
	WEBGL_compressed_texture_etc_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2  GLenum = 0x9276
	WEBGL_compressed_texture_etc_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2 GLenum = 0x9277
	WEBGL_compressed_texture_etc_COMPRESSED_RGBA8_ETC2_EAC                 GLenum = 0x9278
	WEBGL_compressed_texture_etc_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC          GLenum = 0x9279
)

type WEBGL_compressed_texture_etc struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_etc) Once() WEBGL_compressed_texture_etc {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_etc) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_etc) FromRef(ref js.Ref) WEBGL_compressed_texture_etc {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_etc) Free() {
	this.Ref().Free()
}

const (
	WEBGL_compressed_texture_etc1_COMPRESSED_RGB_ETC1_WEBGL GLenum = 0x8D64
)

type WEBGL_compressed_texture_etc1 struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_etc1) Once() WEBGL_compressed_texture_etc1 {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_etc1) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_etc1) FromRef(ref js.Ref) WEBGL_compressed_texture_etc1 {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_etc1) Free() {
	this.Ref().Free()
}

const (
	WEBGL_compressed_texture_pvrtc_COMPRESSED_RGB_PVRTC_4BPPV1_IMG  GLenum = 0x8C00
	WEBGL_compressed_texture_pvrtc_COMPRESSED_RGB_PVRTC_2BPPV1_IMG  GLenum = 0x8C01
	WEBGL_compressed_texture_pvrtc_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG GLenum = 0x8C02
	WEBGL_compressed_texture_pvrtc_COMPRESSED_RGBA_PVRTC_2BPPV1_IMG GLenum = 0x8C03
)

type WEBGL_compressed_texture_pvrtc struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_pvrtc) Once() WEBGL_compressed_texture_pvrtc {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_pvrtc) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_pvrtc) FromRef(ref js.Ref) WEBGL_compressed_texture_pvrtc {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_pvrtc) Free() {
	this.Ref().Free()
}

const (
	WEBGL_compressed_texture_s3tc_COMPRESSED_RGB_S3TC_DXT1_EXT  GLenum = 0x83F0
	WEBGL_compressed_texture_s3tc_COMPRESSED_RGBA_S3TC_DXT1_EXT GLenum = 0x83F1
	WEBGL_compressed_texture_s3tc_COMPRESSED_RGBA_S3TC_DXT3_EXT GLenum = 0x83F2
	WEBGL_compressed_texture_s3tc_COMPRESSED_RGBA_S3TC_DXT5_EXT GLenum = 0x83F3
)

type WEBGL_compressed_texture_s3tc struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_s3tc) Once() WEBGL_compressed_texture_s3tc {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_s3tc) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_s3tc) FromRef(ref js.Ref) WEBGL_compressed_texture_s3tc {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_s3tc) Free() {
	this.Ref().Free()
}

const (
	WEBGL_compressed_texture_s3tc_srgb_COMPRESSED_SRGB_S3TC_DXT1_EXT       GLenum = 0x8C4C
	WEBGL_compressed_texture_s3tc_srgb_COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT GLenum = 0x8C4D
	WEBGL_compressed_texture_s3tc_srgb_COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT GLenum = 0x8C4E
	WEBGL_compressed_texture_s3tc_srgb_COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT GLenum = 0x8C4F
)

type WEBGL_compressed_texture_s3tc_srgb struct {
	ref js.Ref
}

func (this WEBGL_compressed_texture_s3tc_srgb) Once() WEBGL_compressed_texture_s3tc_srgb {
	this.Ref().Once()
	return this
}

func (this WEBGL_compressed_texture_s3tc_srgb) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_compressed_texture_s3tc_srgb) FromRef(ref js.Ref) WEBGL_compressed_texture_s3tc_srgb {
	this.ref = ref
	return this
}

func (this WEBGL_compressed_texture_s3tc_srgb) Free() {
	this.Ref().Free()
}

const (
	WEBGL_debug_renderer_info_UNMASKED_VENDOR_WEBGL   GLenum = 0x9245
	WEBGL_debug_renderer_info_UNMASKED_RENDERER_WEBGL GLenum = 0x9246
)

type WEBGL_debug_renderer_info struct {
	ref js.Ref
}

func (this WEBGL_debug_renderer_info) Once() WEBGL_debug_renderer_info {
	this.Ref().Once()
	return this
}

func (this WEBGL_debug_renderer_info) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_debug_renderer_info) FromRef(ref js.Ref) WEBGL_debug_renderer_info {
	this.ref = ref
	return this
}

func (this WEBGL_debug_renderer_info) Free() {
	this.Ref().Free()
}

type WEBGL_debug_shaders struct {
	ref js.Ref
}

func (this WEBGL_debug_shaders) Once() WEBGL_debug_shaders {
	this.Ref().Once()
	return this
}

func (this WEBGL_debug_shaders) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_debug_shaders) FromRef(ref js.Ref) WEBGL_debug_shaders {
	this.ref = ref
	return this
}

func (this WEBGL_debug_shaders) Free() {
	this.Ref().Free()
}

// HasGetTranslatedShaderSource returns true if the method "WEBGL_debug_shaders.getTranslatedShaderSource" exists.
func (this WEBGL_debug_shaders) HasGetTranslatedShaderSource() bool {
	return js.True == bindings.HasWEBGL_debug_shadersGetTranslatedShaderSource(
		this.Ref(),
	)
}

// GetTranslatedShaderSourceFunc returns the method "WEBGL_debug_shaders.getTranslatedShaderSource".
func (this WEBGL_debug_shaders) GetTranslatedShaderSourceFunc() (fn js.Func[func(shader WebGLShader) js.String]) {
	return fn.FromRef(
		bindings.WEBGL_debug_shadersGetTranslatedShaderSourceFunc(
			this.Ref(),
		),
	)
}

// GetTranslatedShaderSource calls the method "WEBGL_debug_shaders.getTranslatedShaderSource".
func (this WEBGL_debug_shaders) GetTranslatedShaderSource(shader WebGLShader) (ret js.String) {
	bindings.CallWEBGL_debug_shadersGetTranslatedShaderSource(
		this.Ref(), js.Pointer(&ret),
		shader.Ref(),
	)

	return
}

// TryGetTranslatedShaderSource calls the method "WEBGL_debug_shaders.getTranslatedShaderSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_debug_shaders) TryGetTranslatedShaderSource(shader WebGLShader) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_debug_shadersGetTranslatedShaderSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shader.Ref(),
	)

	return
}

const (
	WEBGL_depth_texture_UNSIGNED_INT_24_8_WEBGL GLenum = 0x84FA
)

type WEBGL_depth_texture struct {
	ref js.Ref
}

func (this WEBGL_depth_texture) Once() WEBGL_depth_texture {
	this.Ref().Once()
	return this
}

func (this WEBGL_depth_texture) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_depth_texture) FromRef(ref js.Ref) WEBGL_depth_texture {
	this.ref = ref
	return this
}

func (this WEBGL_depth_texture) Free() {
	this.Ref().Free()
}

const (
	WEBGL_draw_buffers_COLOR_ATTACHMENT0_WEBGL     GLenum = 0x8CE0
	WEBGL_draw_buffers_COLOR_ATTACHMENT1_WEBGL     GLenum = 0x8CE1
	WEBGL_draw_buffers_COLOR_ATTACHMENT2_WEBGL     GLenum = 0x8CE2
	WEBGL_draw_buffers_COLOR_ATTACHMENT3_WEBGL     GLenum = 0x8CE3
	WEBGL_draw_buffers_COLOR_ATTACHMENT4_WEBGL     GLenum = 0x8CE4
	WEBGL_draw_buffers_COLOR_ATTACHMENT5_WEBGL     GLenum = 0x8CE5
	WEBGL_draw_buffers_COLOR_ATTACHMENT6_WEBGL     GLenum = 0x8CE6
	WEBGL_draw_buffers_COLOR_ATTACHMENT7_WEBGL     GLenum = 0x8CE7
	WEBGL_draw_buffers_COLOR_ATTACHMENT8_WEBGL     GLenum = 0x8CE8
	WEBGL_draw_buffers_COLOR_ATTACHMENT9_WEBGL     GLenum = 0x8CE9
	WEBGL_draw_buffers_COLOR_ATTACHMENT10_WEBGL    GLenum = 0x8CEA
	WEBGL_draw_buffers_COLOR_ATTACHMENT11_WEBGL    GLenum = 0x8CEB
	WEBGL_draw_buffers_COLOR_ATTACHMENT12_WEBGL    GLenum = 0x8CEC
	WEBGL_draw_buffers_COLOR_ATTACHMENT13_WEBGL    GLenum = 0x8CED
	WEBGL_draw_buffers_COLOR_ATTACHMENT14_WEBGL    GLenum = 0x8CEE
	WEBGL_draw_buffers_COLOR_ATTACHMENT15_WEBGL    GLenum = 0x8CEF
	WEBGL_draw_buffers_DRAW_BUFFER0_WEBGL          GLenum = 0x8825
	WEBGL_draw_buffers_DRAW_BUFFER1_WEBGL          GLenum = 0x8826
	WEBGL_draw_buffers_DRAW_BUFFER2_WEBGL          GLenum = 0x8827
	WEBGL_draw_buffers_DRAW_BUFFER3_WEBGL          GLenum = 0x8828
	WEBGL_draw_buffers_DRAW_BUFFER4_WEBGL          GLenum = 0x8829
	WEBGL_draw_buffers_DRAW_BUFFER5_WEBGL          GLenum = 0x882A
	WEBGL_draw_buffers_DRAW_BUFFER6_WEBGL          GLenum = 0x882B
	WEBGL_draw_buffers_DRAW_BUFFER7_WEBGL          GLenum = 0x882C
	WEBGL_draw_buffers_DRAW_BUFFER8_WEBGL          GLenum = 0x882D
	WEBGL_draw_buffers_DRAW_BUFFER9_WEBGL          GLenum = 0x882E
	WEBGL_draw_buffers_DRAW_BUFFER10_WEBGL         GLenum = 0x882F
	WEBGL_draw_buffers_DRAW_BUFFER11_WEBGL         GLenum = 0x8830
	WEBGL_draw_buffers_DRAW_BUFFER12_WEBGL         GLenum = 0x8831
	WEBGL_draw_buffers_DRAW_BUFFER13_WEBGL         GLenum = 0x8832
	WEBGL_draw_buffers_DRAW_BUFFER14_WEBGL         GLenum = 0x8833
	WEBGL_draw_buffers_DRAW_BUFFER15_WEBGL         GLenum = 0x8834
	WEBGL_draw_buffers_MAX_COLOR_ATTACHMENTS_WEBGL GLenum = 0x8CDF
	WEBGL_draw_buffers_MAX_DRAW_BUFFERS_WEBGL      GLenum = 0x8824
)

type WEBGL_draw_buffers struct {
	ref js.Ref
}

func (this WEBGL_draw_buffers) Once() WEBGL_draw_buffers {
	this.Ref().Once()
	return this
}

func (this WEBGL_draw_buffers) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_draw_buffers) FromRef(ref js.Ref) WEBGL_draw_buffers {
	this.ref = ref
	return this
}

func (this WEBGL_draw_buffers) Free() {
	this.Ref().Free()
}

// HasDrawBuffersWEBGL returns true if the method "WEBGL_draw_buffers.drawBuffersWEBGL" exists.
func (this WEBGL_draw_buffers) HasDrawBuffersWEBGL() bool {
	return js.True == bindings.HasWEBGL_draw_buffersDrawBuffersWEBGL(
		this.Ref(),
	)
}

// DrawBuffersWEBGLFunc returns the method "WEBGL_draw_buffers.drawBuffersWEBGL".
func (this WEBGL_draw_buffers) DrawBuffersWEBGLFunc() (fn js.Func[func(buffers js.Array[GLenum])]) {
	return fn.FromRef(
		bindings.WEBGL_draw_buffersDrawBuffersWEBGLFunc(
			this.Ref(),
		),
	)
}

// DrawBuffersWEBGL calls the method "WEBGL_draw_buffers.drawBuffersWEBGL".
func (this WEBGL_draw_buffers) DrawBuffersWEBGL(buffers js.Array[GLenum]) (ret js.Void) {
	bindings.CallWEBGL_draw_buffersDrawBuffersWEBGL(
		this.Ref(), js.Pointer(&ret),
		buffers.Ref(),
	)

	return
}

// TryDrawBuffersWEBGL calls the method "WEBGL_draw_buffers.drawBuffersWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_draw_buffers) TryDrawBuffersWEBGL(buffers js.Array[GLenum]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_draw_buffersDrawBuffersWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffers.Ref(),
	)

	return
}

type WEBGL_draw_instanced_base_vertex_base_instance struct {
	ref js.Ref
}

func (this WEBGL_draw_instanced_base_vertex_base_instance) Once() WEBGL_draw_instanced_base_vertex_base_instance {
	this.Ref().Once()
	return this
}

func (this WEBGL_draw_instanced_base_vertex_base_instance) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_draw_instanced_base_vertex_base_instance) FromRef(ref js.Ref) WEBGL_draw_instanced_base_vertex_base_instance {
	this.ref = ref
	return this
}

func (this WEBGL_draw_instanced_base_vertex_base_instance) Free() {
	this.Ref().Free()
}

// HasDrawArraysInstancedBaseInstanceWEBGL returns true if the method "WEBGL_draw_instanced_base_vertex_base_instance.drawArraysInstancedBaseInstanceWEBGL" exists.
func (this WEBGL_draw_instanced_base_vertex_base_instance) HasDrawArraysInstancedBaseInstanceWEBGL() bool {
	return js.True == bindings.HasWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(),
	)
}

// DrawArraysInstancedBaseInstanceWEBGLFunc returns the method "WEBGL_draw_instanced_base_vertex_base_instance.drawArraysInstancedBaseInstanceWEBGL".
func (this WEBGL_draw_instanced_base_vertex_base_instance) DrawArraysInstancedBaseInstanceWEBGLFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei, baseInstance GLuint)]) {
	return fn.FromRef(
		bindings.WEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGLFunc(
			this.Ref(),
		),
	)
}

// DrawArraysInstancedBaseInstanceWEBGL calls the method "WEBGL_draw_instanced_base_vertex_base_instance.drawArraysInstancedBaseInstanceWEBGL".
func (this WEBGL_draw_instanced_base_vertex_base_instance) DrawArraysInstancedBaseInstanceWEBGL(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei, baseInstance GLuint) (ret js.Void) {
	bindings.CallWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
		uint32(baseInstance),
	)

	return
}

// TryDrawArraysInstancedBaseInstanceWEBGL calls the method "WEBGL_draw_instanced_base_vertex_base_instance.drawArraysInstancedBaseInstanceWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_draw_instanced_base_vertex_base_instance) TryDrawArraysInstancedBaseInstanceWEBGL(mode GLenum, first GLint, count GLsizei, instanceCount GLsizei, baseInstance GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_draw_instanced_base_vertex_base_instanceDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
		int32(instanceCount),
		uint32(baseInstance),
	)

	return
}

// HasDrawElementsInstancedBaseVertexBaseInstanceWEBGL returns true if the method "WEBGL_draw_instanced_base_vertex_base_instance.drawElementsInstancedBaseVertexBaseInstanceWEBGL" exists.
func (this WEBGL_draw_instanced_base_vertex_base_instance) HasDrawElementsInstancedBaseVertexBaseInstanceWEBGL() bool {
	return js.True == bindings.HasWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(),
	)
}

// DrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc returns the method "WEBGL_draw_instanced_base_vertex_base_instance.drawElementsInstancedBaseVertexBaseInstanceWEBGL".
func (this WEBGL_draw_instanced_base_vertex_base_instance) DrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei, baseVertex GLint, baseInstance GLuint)]) {
	return fn.FromRef(
		bindings.WEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(
			this.Ref(),
		),
	)
}

// DrawElementsInstancedBaseVertexBaseInstanceWEBGL calls the method "WEBGL_draw_instanced_base_vertex_base_instance.drawElementsInstancedBaseVertexBaseInstanceWEBGL".
func (this WEBGL_draw_instanced_base_vertex_base_instance) DrawElementsInstancedBaseVertexBaseInstanceWEBGL(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei, baseVertex GLint, baseInstance GLuint) (ret js.Void) {
	bindings.CallWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
		int32(baseVertex),
		uint32(baseInstance),
	)

	return
}

// TryDrawElementsInstancedBaseVertexBaseInstanceWEBGL calls the method "WEBGL_draw_instanced_base_vertex_base_instance.drawElementsInstancedBaseVertexBaseInstanceWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_draw_instanced_base_vertex_base_instance) TryDrawElementsInstancedBaseVertexBaseInstanceWEBGL(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, instanceCount GLsizei, baseVertex GLint, baseInstance GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_draw_instanced_base_vertex_base_instanceDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(instanceCount),
		int32(baseVertex),
		uint32(baseInstance),
	)

	return
}

type WEBGL_lose_context struct {
	ref js.Ref
}

func (this WEBGL_lose_context) Once() WEBGL_lose_context {
	this.Ref().Once()
	return this
}

func (this WEBGL_lose_context) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_lose_context) FromRef(ref js.Ref) WEBGL_lose_context {
	this.ref = ref
	return this
}

func (this WEBGL_lose_context) Free() {
	this.Ref().Free()
}

// HasLoseContext returns true if the method "WEBGL_lose_context.loseContext" exists.
func (this WEBGL_lose_context) HasLoseContext() bool {
	return js.True == bindings.HasWEBGL_lose_contextLoseContext(
		this.Ref(),
	)
}

// LoseContextFunc returns the method "WEBGL_lose_context.loseContext".
func (this WEBGL_lose_context) LoseContextFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WEBGL_lose_contextLoseContextFunc(
			this.Ref(),
		),
	)
}

// LoseContext calls the method "WEBGL_lose_context.loseContext".
func (this WEBGL_lose_context) LoseContext() (ret js.Void) {
	bindings.CallWEBGL_lose_contextLoseContext(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLoseContext calls the method "WEBGL_lose_context.loseContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_lose_context) TryLoseContext() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_lose_contextLoseContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRestoreContext returns true if the method "WEBGL_lose_context.restoreContext" exists.
func (this WEBGL_lose_context) HasRestoreContext() bool {
	return js.True == bindings.HasWEBGL_lose_contextRestoreContext(
		this.Ref(),
	)
}

// RestoreContextFunc returns the method "WEBGL_lose_context.restoreContext".
func (this WEBGL_lose_context) RestoreContextFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WEBGL_lose_contextRestoreContextFunc(
			this.Ref(),
		),
	)
}

// RestoreContext calls the method "WEBGL_lose_context.restoreContext".
func (this WEBGL_lose_context) RestoreContext() (ret js.Void) {
	bindings.CallWEBGL_lose_contextRestoreContext(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRestoreContext calls the method "WEBGL_lose_context.restoreContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_lose_context) TryRestoreContext() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_lose_contextRestoreContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_TypedArrayInt32_ArrayGLsizei struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt32_ArrayGLsizei) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt32_ArrayGLsizei) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt32_ArrayGLsizei) FromRef(ref js.Ref) OneOf_TypedArrayInt32_ArrayGLsizei {
	return OneOf_TypedArrayInt32_ArrayGLsizei{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt32_ArrayGLsizei) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt32_ArrayGLsizei) ArrayGLsizei() js.Array[GLsizei] {
	return js.Array[GLsizei]{}.FromRef(x.ref)
}

type WEBGL_multi_draw struct {
	ref js.Ref
}

func (this WEBGL_multi_draw) Once() WEBGL_multi_draw {
	this.Ref().Once()
	return this
}

func (this WEBGL_multi_draw) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_multi_draw) FromRef(ref js.Ref) WEBGL_multi_draw {
	this.ref = ref
	return this
}

func (this WEBGL_multi_draw) Free() {
	this.Ref().Free()
}

// HasMultiDrawArraysWEBGL returns true if the method "WEBGL_multi_draw.multiDrawArraysWEBGL" exists.
func (this WEBGL_multi_draw) HasMultiDrawArraysWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_drawMultiDrawArraysWEBGL(
		this.Ref(),
	)
}

// MultiDrawArraysWEBGLFunc returns the method "WEBGL_multi_draw.multiDrawArraysWEBGL".
func (this WEBGL_multi_draw) MultiDrawArraysWEBGLFunc() (fn js.Func[func(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_drawMultiDrawArraysWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawArraysWEBGL calls the method "WEBGL_multi_draw.multiDrawArraysWEBGL".
func (this WEBGL_multi_draw) MultiDrawArraysWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_drawMultiDrawArraysWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawArraysWEBGL calls the method "WEBGL_multi_draw.multiDrawArraysWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw) TryMultiDrawArraysWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_drawMultiDrawArraysWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		int32(drawcount),
	)

	return
}

// HasMultiDrawElementsWEBGL returns true if the method "WEBGL_multi_draw.multiDrawElementsWEBGL" exists.
func (this WEBGL_multi_draw) HasMultiDrawElementsWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_drawMultiDrawElementsWEBGL(
		this.Ref(),
	)
}

// MultiDrawElementsWEBGLFunc returns the method "WEBGL_multi_draw.multiDrawElementsWEBGL".
func (this WEBGL_multi_draw) MultiDrawElementsWEBGLFunc() (fn js.Func[func(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_drawMultiDrawElementsWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawElementsWEBGL calls the method "WEBGL_multi_draw.multiDrawElementsWEBGL".
func (this WEBGL_multi_draw) MultiDrawElementsWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_drawMultiDrawElementsWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawElementsWEBGL calls the method "WEBGL_multi_draw.multiDrawElementsWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw) TryMultiDrawElementsWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_drawMultiDrawElementsWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		int32(drawcount),
	)

	return
}

// HasMultiDrawArraysInstancedWEBGL returns true if the method "WEBGL_multi_draw.multiDrawArraysInstancedWEBGL" exists.
func (this WEBGL_multi_draw) HasMultiDrawArraysInstancedWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
		this.Ref(),
	)
}

// MultiDrawArraysInstancedWEBGLFunc returns the method "WEBGL_multi_draw.multiDrawArraysInstancedWEBGL".
func (this WEBGL_multi_draw) MultiDrawArraysInstancedWEBGLFunc() (fn js.Func[func(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_drawMultiDrawArraysInstancedWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawArraysInstancedWEBGL calls the method "WEBGL_multi_draw.multiDrawArraysInstancedWEBGL".
func (this WEBGL_multi_draw) MultiDrawArraysInstancedWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawArraysInstancedWEBGL calls the method "WEBGL_multi_draw.multiDrawArraysInstancedWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw) TryMultiDrawArraysInstancedWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_drawMultiDrawArraysInstancedWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		int32(drawcount),
	)

	return
}

// HasMultiDrawElementsInstancedWEBGL returns true if the method "WEBGL_multi_draw.multiDrawElementsInstancedWEBGL" exists.
func (this WEBGL_multi_draw) HasMultiDrawElementsInstancedWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
		this.Ref(),
	)
}

// MultiDrawElementsInstancedWEBGLFunc returns the method "WEBGL_multi_draw.multiDrawElementsInstancedWEBGL".
func (this WEBGL_multi_draw) MultiDrawElementsInstancedWEBGLFunc() (fn js.Func[func(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_drawMultiDrawElementsInstancedWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawElementsInstancedWEBGL calls the method "WEBGL_multi_draw.multiDrawElementsInstancedWEBGL".
func (this WEBGL_multi_draw) MultiDrawElementsInstancedWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawElementsInstancedWEBGL calls the method "WEBGL_multi_draw.multiDrawElementsInstancedWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw) TryMultiDrawElementsInstancedWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_drawMultiDrawElementsInstancedWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		int32(drawcount),
	)

	return
}

type WEBGL_multi_draw_instanced_base_vertex_base_instance struct {
	ref js.Ref
}

func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) Once() WEBGL_multi_draw_instanced_base_vertex_base_instance {
	this.Ref().Once()
	return this
}

func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) FromRef(ref js.Ref) WEBGL_multi_draw_instanced_base_vertex_base_instance {
	this.ref = ref
	return this
}

func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) Free() {
	this.Ref().Free()
}

// HasMultiDrawArraysInstancedBaseInstanceWEBGL returns true if the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawArraysInstancedBaseInstanceWEBGL" exists.
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) HasMultiDrawArraysInstancedBaseInstanceWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(),
	)
}

// MultiDrawArraysInstancedBaseInstanceWEBGLFunc returns the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawArraysInstancedBaseInstanceWEBGL".
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) MultiDrawArraysInstancedBaseInstanceWEBGLFunc() (fn js.Func[func(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawArraysInstancedBaseInstanceWEBGL calls the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawArraysInstancedBaseInstanceWEBGL".
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) MultiDrawArraysInstancedBaseInstanceWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		baseInstancesList.Ref(),
		uint32(baseInstancesOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawArraysInstancedBaseInstanceWEBGL calls the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawArraysInstancedBaseInstanceWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) TryMultiDrawArraysInstancedBaseInstanceWEBGL(mode GLenum, firstsList OneOf_TypedArrayInt32_ArrayGLint, firstsOffset GLuint, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawArraysInstancedBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		firstsList.Ref(),
		uint32(firstsOffset),
		countsList.Ref(),
		uint32(countsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		baseInstancesList.Ref(),
		uint32(baseInstancesOffset),
		int32(drawcount),
	)

	return
}

// HasMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL returns true if the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawElementsInstancedBaseVertexBaseInstanceWEBGL" exists.
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) HasMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL() bool {
	return js.True == bindings.HasWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(),
	)
}

// MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc returns the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawElementsInstancedBaseVertexBaseInstanceWEBGL".
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc() (fn js.Func[func(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseVerticesList OneOf_TypedArrayInt32_ArrayGLint, baseVerticesOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei)]) {
	return fn.FromRef(
		bindings.WEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGLFunc(
			this.Ref(),
		),
	)
}

// MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL calls the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawElementsInstancedBaseVertexBaseInstanceWEBGL".
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) MultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseVerticesList OneOf_TypedArrayInt32_ArrayGLint, baseVerticesOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei) (ret js.Void) {
	bindings.CallWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		baseVerticesList.Ref(),
		uint32(baseVerticesOffset),
		baseInstancesList.Ref(),
		uint32(baseInstancesOffset),
		int32(drawcount),
	)

	return
}

// TryMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL calls the method "WEBGL_multi_draw_instanced_base_vertex_base_instance.multiDrawElementsInstancedBaseVertexBaseInstanceWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_multi_draw_instanced_base_vertex_base_instance) TryMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(mode GLenum, countsList OneOf_TypedArrayInt32_ArrayGLsizei, countsOffset GLuint, typ GLenum, offsetsList OneOf_TypedArrayInt32_ArrayGLsizei, offsetsOffset GLuint, instanceCountsList OneOf_TypedArrayInt32_ArrayGLsizei, instanceCountsOffset GLuint, baseVerticesList OneOf_TypedArrayInt32_ArrayGLint, baseVerticesOffset GLuint, baseInstancesList OneOf_TypedArrayUint32_ArrayGLuint, baseInstancesOffset GLuint, drawcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_multi_draw_instanced_base_vertex_base_instanceMultiDrawElementsInstancedBaseVertexBaseInstanceWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		countsList.Ref(),
		uint32(countsOffset),
		uint32(typ),
		offsetsList.Ref(),
		uint32(offsetsOffset),
		instanceCountsList.Ref(),
		uint32(instanceCountsOffset),
		baseVerticesList.Ref(),
		uint32(baseVerticesOffset),
		baseInstancesList.Ref(),
		uint32(baseInstancesOffset),
		int32(drawcount),
	)

	return
}

const (
	WEBGL_provoking_vertex_FIRST_VERTEX_CONVENTION_WEBGL GLenum = 0x8E4D
	WEBGL_provoking_vertex_LAST_VERTEX_CONVENTION_WEBGL  GLenum = 0x8E4E
	WEBGL_provoking_vertex_PROVOKING_VERTEX_WEBGL        GLenum = 0x8E4F
)

type WEBGL_provoking_vertex struct {
	ref js.Ref
}

func (this WEBGL_provoking_vertex) Once() WEBGL_provoking_vertex {
	this.Ref().Once()
	return this
}

func (this WEBGL_provoking_vertex) Ref() js.Ref {
	return this.ref
}

func (this WEBGL_provoking_vertex) FromRef(ref js.Ref) WEBGL_provoking_vertex {
	this.ref = ref
	return this
}

func (this WEBGL_provoking_vertex) Free() {
	this.Ref().Free()
}

// HasProvokingVertexWEBGL returns true if the method "WEBGL_provoking_vertex.provokingVertexWEBGL" exists.
func (this WEBGL_provoking_vertex) HasProvokingVertexWEBGL() bool {
	return js.True == bindings.HasWEBGL_provoking_vertexProvokingVertexWEBGL(
		this.Ref(),
	)
}

// ProvokingVertexWEBGLFunc returns the method "WEBGL_provoking_vertex.provokingVertexWEBGL".
func (this WEBGL_provoking_vertex) ProvokingVertexWEBGLFunc() (fn js.Func[func(provokeMode GLenum)]) {
	return fn.FromRef(
		bindings.WEBGL_provoking_vertexProvokingVertexWEBGLFunc(
			this.Ref(),
		),
	)
}

// ProvokingVertexWEBGL calls the method "WEBGL_provoking_vertex.provokingVertexWEBGL".
func (this WEBGL_provoking_vertex) ProvokingVertexWEBGL(provokeMode GLenum) (ret js.Void) {
	bindings.CallWEBGL_provoking_vertexProvokingVertexWEBGL(
		this.Ref(), js.Pointer(&ret),
		uint32(provokeMode),
	)

	return
}

// TryProvokingVertexWEBGL calls the method "WEBGL_provoking_vertex.provokingVertexWEBGL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WEBGL_provoking_vertex) TryProvokingVertexWEBGL(provokeMode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWEBGL_provoking_vertexProvokingVertexWEBGL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(provokeMode),
	)

	return
}

type WebAssemblyInstantiatedSource struct {
	// Module is "WebAssemblyInstantiatedSource.module"
	//
	// Required
	Module Module
	// Instance is "WebAssemblyInstantiatedSource.instance"
	//
	// Required
	Instance Instance

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebAssemblyInstantiatedSource with all fields set.
func (p WebAssemblyInstantiatedSource) FromRef(ref js.Ref) WebAssemblyInstantiatedSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebAssemblyInstantiatedSource in the application heap.
func (p WebAssemblyInstantiatedSource) New() js.Ref {
	return bindings.WebAssemblyInstantiatedSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebAssemblyInstantiatedSource) UpdateFrom(ref js.Ref) {
	bindings.WebAssemblyInstantiatedSourceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebAssemblyInstantiatedSource) Update(ref js.Ref) {
	bindings.WebAssemblyInstantiatedSourceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WebAssembly struct{}

// HasValidate returns ture if the function "WebAssembly.validate" exists.
func (WebAssembly) HasValidate() bool {
	return js.True == bindings.HasWebAssemblyValidate()
}

// ValidateFunc returns the function "WebAssembly.validate".
func (WebAssembly) ValidateFunc() (fn js.Func[func(bytes BufferSource) bool]) {
	return fn.FromRef(
		bindings.WebAssemblyValidateFunc(),
	)
}

// Validate calls the function "WebAssembly.validate".
func (WebAssembly) Validate(bytes BufferSource) (ret bool) {
	bindings.CallWebAssemblyValidate(
		js.Pointer(&ret),
		bytes.Ref(),
	)
	return
}

// TryValidate calls the function "WebAssembly.validate"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryValidate(bytes BufferSource) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyValidate(
		js.Pointer(&ret), js.Pointer(&exception),
		bytes.Ref(),
	)
	return
}

// HasCompile returns ture if the function "WebAssembly.compile" exists.
func (WebAssembly) HasCompile() bool {
	return js.True == bindings.HasWebAssemblyCompile()
}

// CompileFunc returns the function "WebAssembly.compile".
func (WebAssembly) CompileFunc() (fn js.Func[func(bytes BufferSource) js.Promise[Module]]) {
	return fn.FromRef(
		bindings.WebAssemblyCompileFunc(),
	)
}

// Compile calls the function "WebAssembly.compile".
func (WebAssembly) Compile(bytes BufferSource) (ret js.Promise[Module]) {
	bindings.CallWebAssemblyCompile(
		js.Pointer(&ret),
		bytes.Ref(),
	)
	return
}

// TryCompile calls the function "WebAssembly.compile"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryCompile(bytes BufferSource) (ret js.Promise[Module], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyCompile(
		js.Pointer(&ret), js.Pointer(&exception),
		bytes.Ref(),
	)
	return
}

// HasInstantiate returns ture if the function "WebAssembly.instantiate" exists.
func (WebAssembly) HasInstantiate() bool {
	return js.True == bindings.HasWebAssemblyInstantiate()
}

// InstantiateFunc returns the function "WebAssembly.instantiate".
func (WebAssembly) InstantiateFunc() (fn js.Func[func(bytes BufferSource, importObject js.Object) js.Promise[WebAssemblyInstantiatedSource]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiateFunc(),
	)
}

// Instantiate calls the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate(bytes BufferSource, importObject js.Object) (ret js.Promise[WebAssemblyInstantiatedSource]) {
	bindings.CallWebAssemblyInstantiate(
		js.Pointer(&ret),
		bytes.Ref(),
		importObject.Ref(),
	)
	return
}

// TryInstantiate calls the function "WebAssembly.instantiate"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiate(bytes BufferSource, importObject js.Object) (ret js.Promise[WebAssemblyInstantiatedSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiate(
		js.Pointer(&ret), js.Pointer(&exception),
		bytes.Ref(),
		importObject.Ref(),
	)
	return
}

// HasInstantiate1 returns ture if the function "WebAssembly.instantiate" exists.
func (WebAssembly) HasInstantiate1() bool {
	return js.True == bindings.HasWebAssemblyInstantiate1()
}

// Instantiate1Func returns the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate1Func() (fn js.Func[func(bytes BufferSource) js.Promise[WebAssemblyInstantiatedSource]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiate1Func(),
	)
}

// Instantiate1 calls the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate1(bytes BufferSource) (ret js.Promise[WebAssemblyInstantiatedSource]) {
	bindings.CallWebAssemblyInstantiate1(
		js.Pointer(&ret),
		bytes.Ref(),
	)
	return
}

// TryInstantiate1 calls the function "WebAssembly.instantiate"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiate1(bytes BufferSource) (ret js.Promise[WebAssemblyInstantiatedSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiate1(
		js.Pointer(&ret), js.Pointer(&exception),
		bytes.Ref(),
	)
	return
}

// HasInstantiate2 returns ture if the function "WebAssembly.instantiate" exists.
func (WebAssembly) HasInstantiate2() bool {
	return js.True == bindings.HasWebAssemblyInstantiate2()
}

// Instantiate2Func returns the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate2Func() (fn js.Func[func(moduleObject Module, importObject js.Object) js.Promise[Instance]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiate2Func(),
	)
}

// Instantiate2 calls the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate2(moduleObject Module, importObject js.Object) (ret js.Promise[Instance]) {
	bindings.CallWebAssemblyInstantiate2(
		js.Pointer(&ret),
		moduleObject.Ref(),
		importObject.Ref(),
	)
	return
}

// TryInstantiate2 calls the function "WebAssembly.instantiate"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiate2(moduleObject Module, importObject js.Object) (ret js.Promise[Instance], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiate2(
		js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
		importObject.Ref(),
	)
	return
}

// HasInstantiate3 returns ture if the function "WebAssembly.instantiate" exists.
func (WebAssembly) HasInstantiate3() bool {
	return js.True == bindings.HasWebAssemblyInstantiate3()
}

// Instantiate3Func returns the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate3Func() (fn js.Func[func(moduleObject Module) js.Promise[Instance]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiate3Func(),
	)
}

// Instantiate3 calls the function "WebAssembly.instantiate".
func (WebAssembly) Instantiate3(moduleObject Module) (ret js.Promise[Instance]) {
	bindings.CallWebAssemblyInstantiate3(
		js.Pointer(&ret),
		moduleObject.Ref(),
	)
	return
}

// TryInstantiate3 calls the function "WebAssembly.instantiate"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiate3(moduleObject Module) (ret js.Promise[Instance], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiate3(
		js.Pointer(&ret), js.Pointer(&exception),
		moduleObject.Ref(),
	)
	return
}

// HasCompileStreaming returns ture if the function "WebAssembly.compileStreaming" exists.
func (WebAssembly) HasCompileStreaming() bool {
	return js.True == bindings.HasWebAssemblyCompileStreaming()
}

// CompileStreamingFunc returns the function "WebAssembly.compileStreaming".
func (WebAssembly) CompileStreamingFunc() (fn js.Func[func(source js.Promise[Response]) js.Promise[Module]]) {
	return fn.FromRef(
		bindings.WebAssemblyCompileStreamingFunc(),
	)
}

// CompileStreaming calls the function "WebAssembly.compileStreaming".
func (WebAssembly) CompileStreaming(source js.Promise[Response]) (ret js.Promise[Module]) {
	bindings.CallWebAssemblyCompileStreaming(
		js.Pointer(&ret),
		source.Ref(),
	)
	return
}

// TryCompileStreaming calls the function "WebAssembly.compileStreaming"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryCompileStreaming(source js.Promise[Response]) (ret js.Promise[Module], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyCompileStreaming(
		js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
	)
	return
}

// HasInstantiateStreaming returns ture if the function "WebAssembly.instantiateStreaming" exists.
func (WebAssembly) HasInstantiateStreaming() bool {
	return js.True == bindings.HasWebAssemblyInstantiateStreaming()
}

// InstantiateStreamingFunc returns the function "WebAssembly.instantiateStreaming".
func (WebAssembly) InstantiateStreamingFunc() (fn js.Func[func(source js.Promise[Response], importObject js.Object) js.Promise[WebAssemblyInstantiatedSource]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiateStreamingFunc(),
	)
}

// InstantiateStreaming calls the function "WebAssembly.instantiateStreaming".
func (WebAssembly) InstantiateStreaming(source js.Promise[Response], importObject js.Object) (ret js.Promise[WebAssemblyInstantiatedSource]) {
	bindings.CallWebAssemblyInstantiateStreaming(
		js.Pointer(&ret),
		source.Ref(),
		importObject.Ref(),
	)
	return
}

// TryInstantiateStreaming calls the function "WebAssembly.instantiateStreaming"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiateStreaming(source js.Promise[Response], importObject js.Object) (ret js.Promise[WebAssemblyInstantiatedSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiateStreaming(
		js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		importObject.Ref(),
	)
	return
}

// HasInstantiateStreaming1 returns ture if the function "WebAssembly.instantiateStreaming" exists.
func (WebAssembly) HasInstantiateStreaming1() bool {
	return js.True == bindings.HasWebAssemblyInstantiateStreaming1()
}

// InstantiateStreaming1Func returns the function "WebAssembly.instantiateStreaming".
func (WebAssembly) InstantiateStreaming1Func() (fn js.Func[func(source js.Promise[Response]) js.Promise[WebAssemblyInstantiatedSource]]) {
	return fn.FromRef(
		bindings.WebAssemblyInstantiateStreaming1Func(),
	)
}

// InstantiateStreaming1 calls the function "WebAssembly.instantiateStreaming".
func (WebAssembly) InstantiateStreaming1(source js.Promise[Response]) (ret js.Promise[WebAssemblyInstantiatedSource]) {
	bindings.CallWebAssemblyInstantiateStreaming1(
		js.Pointer(&ret),
		source.Ref(),
	)
	return
}

// TryInstantiateStreaming1 calls the function "WebAssembly.instantiateStreaming"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (WebAssembly) TryInstantiateStreaming1(source js.Promise[Response]) (ret js.Promise[WebAssemblyInstantiatedSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWebAssemblyInstantiateStreaming1(
		js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
	)
	return
}

type WebGLContextEventInit struct {
	// StatusMessage is "WebGLContextEventInit.statusMessage"
	//
	// Optional, defaults to "".
	StatusMessage js.String
	// Bubbles is "WebGLContextEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "WebGLContextEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "WebGLContextEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebGLContextEventInit with all fields set.
func (p WebGLContextEventInit) FromRef(ref js.Ref) WebGLContextEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebGLContextEventInit in the application heap.
func (p WebGLContextEventInit) New() js.Ref {
	return bindings.WebGLContextEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WebGLContextEventInit) UpdateFrom(ref js.Ref) {
	bindings.WebGLContextEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WebGLContextEventInit) Update(ref js.Ref) {
	bindings.WebGLContextEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
