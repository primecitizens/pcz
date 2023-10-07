// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

const (
	_ TextTrackMode = iota

	TextTrackMode_DISABLED
	TextTrackMode_HIDDEN
	TextTrackMode_SHOWING
)

func (TextTrackMode) FromRef(str js.Ref) TextTrackMode {
	return TextTrackMode(bindings.ConstOfTextTrackMode(str))
}

func (x TextTrackMode) String() (string, bool) {
	switch x {
	case TextTrackMode_DISABLED:
		return "disabled", true
	case TextTrackMode_HIDDEN:
		return "hidden", true
	case TextTrackMode_SHOWING:
		return "showing", true
	default:
		return "", false
	}
}

type TextTrackCueList struct {
	ref js.Ref
}

func (this TextTrackCueList) Once() TextTrackCueList {
	this.ref.Once()
	return this
}

func (this TextTrackCueList) Ref() js.Ref {
	return this.ref
}

func (this TextTrackCueList) FromRef(ref js.Ref) TextTrackCueList {
	this.ref = ref
	return this
}

func (this TextTrackCueList) Free() {
	this.ref.Free()
}

// Length returns the value of property "TextTrackCueList.length".
//
// It returns ok=false if there is no such property.
func (this TextTrackCueList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextTrackCueListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "TextTrackCueList." exists.
func (this TextTrackCueList) HasFuncGet() bool {
	return js.True == bindings.HasFuncTextTrackCueListGet(
		this.ref,
	)
}

// FuncGet returns the method "TextTrackCueList.".
func (this TextTrackCueList) FuncGet() (fn js.Func[func(index uint32) TextTrackCue]) {
	bindings.FuncTextTrackCueListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "TextTrackCueList.".
func (this TextTrackCueList) Get(index uint32) (ret TextTrackCue) {
	bindings.CallTextTrackCueListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "TextTrackCueList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrackCueList) TryGet(index uint32) (ret TextTrackCue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackCueListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetCueById returns true if the method "TextTrackCueList.getCueById" exists.
func (this TextTrackCueList) HasFuncGetCueById() bool {
	return js.True == bindings.HasFuncTextTrackCueListGetCueById(
		this.ref,
	)
}

// FuncGetCueById returns the method "TextTrackCueList.getCueById".
func (this TextTrackCueList) FuncGetCueById() (fn js.Func[func(id js.String) TextTrackCue]) {
	bindings.FuncTextTrackCueListGetCueById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCueById calls the method "TextTrackCueList.getCueById".
func (this TextTrackCueList) GetCueById(id js.String) (ret TextTrackCue) {
	bindings.CallTextTrackCueListGetCueById(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetCueById calls the method "TextTrackCueList.getCueById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrackCueList) TryGetCueById(id js.String) (ret TextTrackCue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackCueListGetCueById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type TimeRanges struct {
	ref js.Ref
}

func (this TimeRanges) Once() TimeRanges {
	this.ref.Once()
	return this
}

func (this TimeRanges) Ref() js.Ref {
	return this.ref
}

func (this TimeRanges) FromRef(ref js.Ref) TimeRanges {
	this.ref = ref
	return this
}

func (this TimeRanges) Free() {
	this.ref.Free()
}

// Length returns the value of property "TimeRanges.length".
//
// It returns ok=false if there is no such property.
func (this TimeRanges) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTimeRangesLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncStart returns true if the method "TimeRanges.start" exists.
func (this TimeRanges) HasFuncStart() bool {
	return js.True == bindings.HasFuncTimeRangesStart(
		this.ref,
	)
}

// FuncStart returns the method "TimeRanges.start".
func (this TimeRanges) FuncStart() (fn js.Func[func(index uint32) float64]) {
	bindings.FuncTimeRangesStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "TimeRanges.start".
func (this TimeRanges) Start(index uint32) (ret float64) {
	bindings.CallTimeRangesStart(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryStart calls the method "TimeRanges.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TimeRanges) TryStart(index uint32) (ret float64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTimeRangesStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncEnd returns true if the method "TimeRanges.end" exists.
func (this TimeRanges) HasFuncEnd() bool {
	return js.True == bindings.HasFuncTimeRangesEnd(
		this.ref,
	)
}

// FuncEnd returns the method "TimeRanges.end".
func (this TimeRanges) FuncEnd() (fn js.Func[func(index uint32) float64]) {
	bindings.FuncTimeRangesEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// End calls the method "TimeRanges.end".
func (this TimeRanges) End(index uint32) (ret float64) {
	bindings.CallTimeRangesEnd(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryEnd calls the method "TimeRanges.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TimeRanges) TryEnd(index uint32) (ret float64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTimeRangesEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type AudioTrack struct {
	ref js.Ref
}

func (this AudioTrack) Once() AudioTrack {
	this.ref.Once()
	return this
}

func (this AudioTrack) Ref() js.Ref {
	return this.ref
}

func (this AudioTrack) FromRef(ref js.Ref) AudioTrack {
	this.ref = ref
	return this
}

func (this AudioTrack) Free() {
	this.ref.Free()
}

// Id returns the value of property "AudioTrack.id".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAudioTrackId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Kind returns the value of property "AudioTrack.kind".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAudioTrackKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "AudioTrack.label".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAudioTrackLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "AudioTrack.language".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAudioTrackLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Enabled returns the value of property "AudioTrack.enabled".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) Enabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetAudioTrackEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEnabled sets the value of property "AudioTrack.enabled" to val.
//
// It returns false if the property cannot be set.
func (this AudioTrack) SetEnabled(val bool) bool {
	return js.True == bindings.SetAudioTrackEnabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// SourceBuffer returns the value of property "AudioTrack.sourceBuffer".
//
// It returns ok=false if there is no such property.
func (this AudioTrack) SourceBuffer() (ret SourceBuffer, ok bool) {
	ok = js.True == bindings.GetAudioTrackSourceBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioTrackList struct {
	EventTarget
}

func (this AudioTrackList) Once() AudioTrackList {
	this.ref.Once()
	return this
}

func (this AudioTrackList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AudioTrackList) FromRef(ref js.Ref) AudioTrackList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AudioTrackList) Free() {
	this.ref.Free()
}

// Length returns the value of property "AudioTrackList.length".
//
// It returns ok=false if there is no such property.
func (this AudioTrackList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioTrackListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "AudioTrackList." exists.
func (this AudioTrackList) HasFuncGet() bool {
	return js.True == bindings.HasFuncAudioTrackListGet(
		this.ref,
	)
}

// FuncGet returns the method "AudioTrackList.".
func (this AudioTrackList) FuncGet() (fn js.Func[func(index uint32) AudioTrack]) {
	bindings.FuncAudioTrackListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "AudioTrackList.".
func (this AudioTrackList) Get(index uint32) (ret AudioTrack) {
	bindings.CallAudioTrackListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "AudioTrackList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioTrackList) TryGet(index uint32) (ret AudioTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioTrackListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetTrackById returns true if the method "AudioTrackList.getTrackById" exists.
func (this AudioTrackList) HasFuncGetTrackById() bool {
	return js.True == bindings.HasFuncAudioTrackListGetTrackById(
		this.ref,
	)
}

// FuncGetTrackById returns the method "AudioTrackList.getTrackById".
func (this AudioTrackList) FuncGetTrackById() (fn js.Func[func(id js.String) AudioTrack]) {
	bindings.FuncAudioTrackListGetTrackById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTrackById calls the method "AudioTrackList.getTrackById".
func (this AudioTrackList) GetTrackById(id js.String) (ret AudioTrack) {
	bindings.CallAudioTrackListGetTrackById(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetTrackById calls the method "AudioTrackList.getTrackById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioTrackList) TryGetTrackById(id js.String) (ret AudioTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioTrackListGetTrackById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type VideoTrack struct {
	ref js.Ref
}

func (this VideoTrack) Once() VideoTrack {
	this.ref.Once()
	return this
}

func (this VideoTrack) Ref() js.Ref {
	return this.ref
}

func (this VideoTrack) FromRef(ref js.Ref) VideoTrack {
	this.ref = ref
	return this
}

func (this VideoTrack) Free() {
	this.ref.Free()
}

// Id returns the value of property "VideoTrack.id".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVideoTrackId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Kind returns the value of property "VideoTrack.kind".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVideoTrackKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "VideoTrack.label".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVideoTrackLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "VideoTrack.language".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVideoTrackLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Selected returns the value of property "VideoTrack.selected".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetVideoTrackSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "VideoTrack.selected" to val.
//
// It returns false if the property cannot be set.
func (this VideoTrack) SetSelected(val bool) bool {
	return js.True == bindings.SetVideoTrackSelected(
		this.ref,
		js.Bool(bool(val)),
	)
}

// SourceBuffer returns the value of property "VideoTrack.sourceBuffer".
//
// It returns ok=false if there is no such property.
func (this VideoTrack) SourceBuffer() (ret SourceBuffer, ok bool) {
	ok = js.True == bindings.GetVideoTrackSourceBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

type VideoTrackList struct {
	EventTarget
}

func (this VideoTrackList) Once() VideoTrackList {
	this.ref.Once()
	return this
}

func (this VideoTrackList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this VideoTrackList) FromRef(ref js.Ref) VideoTrackList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this VideoTrackList) Free() {
	this.ref.Free()
}

// Length returns the value of property "VideoTrackList.length".
//
// It returns ok=false if there is no such property.
func (this VideoTrackList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoTrackListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectedIndex returns the value of property "VideoTrackList.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this VideoTrackList) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetVideoTrackListSelectedIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "VideoTrackList." exists.
func (this VideoTrackList) HasFuncGet() bool {
	return js.True == bindings.HasFuncVideoTrackListGet(
		this.ref,
	)
}

// FuncGet returns the method "VideoTrackList.".
func (this VideoTrackList) FuncGet() (fn js.Func[func(index uint32) VideoTrack]) {
	bindings.FuncVideoTrackListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "VideoTrackList.".
func (this VideoTrackList) Get(index uint32) (ret VideoTrack) {
	bindings.CallVideoTrackListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "VideoTrackList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoTrackList) TryGet(index uint32) (ret VideoTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoTrackListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetTrackById returns true if the method "VideoTrackList.getTrackById" exists.
func (this VideoTrackList) HasFuncGetTrackById() bool {
	return js.True == bindings.HasFuncVideoTrackListGetTrackById(
		this.ref,
	)
}

// FuncGetTrackById returns the method "VideoTrackList.getTrackById".
func (this VideoTrackList) FuncGetTrackById() (fn js.Func[func(id js.String) VideoTrack]) {
	bindings.FuncVideoTrackListGetTrackById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTrackById calls the method "VideoTrackList.getTrackById".
func (this VideoTrackList) GetTrackById(id js.String) (ret VideoTrack) {
	bindings.CallVideoTrackListGetTrackById(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetTrackById calls the method "VideoTrackList.getTrackById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoTrackList) TryGetTrackById(id js.String) (ret VideoTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoTrackListGetTrackById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type TextTrackList struct {
	EventTarget
}

func (this TextTrackList) Once() TextTrackList {
	this.ref.Once()
	return this
}

func (this TextTrackList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this TextTrackList) FromRef(ref js.Ref) TextTrackList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this TextTrackList) Free() {
	this.ref.Free()
}

// Length returns the value of property "TextTrackList.length".
//
// It returns ok=false if there is no such property.
func (this TextTrackList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetTextTrackListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "TextTrackList." exists.
func (this TextTrackList) HasFuncGet() bool {
	return js.True == bindings.HasFuncTextTrackListGet(
		this.ref,
	)
}

// FuncGet returns the method "TextTrackList.".
func (this TextTrackList) FuncGet() (fn js.Func[func(index uint32) TextTrack]) {
	bindings.FuncTextTrackListGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "TextTrackList.".
func (this TextTrackList) Get(index uint32) (ret TextTrack) {
	bindings.CallTextTrackListGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "TextTrackList."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrackList) TryGet(index uint32) (ret TextTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackListGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncGetTrackById returns true if the method "TextTrackList.getTrackById" exists.
func (this TextTrackList) HasFuncGetTrackById() bool {
	return js.True == bindings.HasFuncTextTrackListGetTrackById(
		this.ref,
	)
}

// FuncGetTrackById returns the method "TextTrackList.getTrackById".
func (this TextTrackList) FuncGetTrackById() (fn js.Func[func(id js.String) TextTrack]) {
	bindings.FuncTextTrackListGetTrackById(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetTrackById calls the method "TextTrackList.getTrackById".
func (this TextTrackList) GetTrackById(id js.String) (ret TextTrack) {
	bindings.CallTextTrackListGetTrackById(
		this.ref, js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetTrackById calls the method "TextTrackList.getTrackById"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrackList) TryGetTrackById(id js.String) (ret TextTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackListGetTrackById(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

type SourceBuffer struct {
	EventTarget
}

func (this SourceBuffer) Once() SourceBuffer {
	this.ref.Once()
	return this
}

func (this SourceBuffer) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SourceBuffer) FromRef(ref js.Ref) SourceBuffer {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SourceBuffer) Free() {
	this.ref.Free()
}

// Mode returns the value of property "SourceBuffer.mode".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) Mode() (ret AppendMode, ok bool) {
	ok = js.True == bindings.GetSourceBufferMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMode sets the value of property "SourceBuffer.mode" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetMode(val AppendMode) bool {
	return js.True == bindings.SetSourceBufferMode(
		this.ref,
		uint32(val),
	)
}

// Updating returns the value of property "SourceBuffer.updating".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) Updating() (ret bool, ok bool) {
	ok = js.True == bindings.GetSourceBufferUpdating(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Buffered returns the value of property "SourceBuffer.buffered".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) Buffered() (ret TimeRanges, ok bool) {
	ok = js.True == bindings.GetSourceBufferBuffered(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TimestampOffset returns the value of property "SourceBuffer.timestampOffset".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) TimestampOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetSourceBufferTimestampOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTimestampOffset sets the value of property "SourceBuffer.timestampOffset" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetTimestampOffset(val float64) bool {
	return js.True == bindings.SetSourceBufferTimestampOffset(
		this.ref,
		float64(val),
	)
}

// AudioTracks returns the value of property "SourceBuffer.audioTracks".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) AudioTracks() (ret AudioTrackList, ok bool) {
	ok = js.True == bindings.GetSourceBufferAudioTracks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VideoTracks returns the value of property "SourceBuffer.videoTracks".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) VideoTracks() (ret VideoTrackList, ok bool) {
	ok = js.True == bindings.GetSourceBufferVideoTracks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TextTracks returns the value of property "SourceBuffer.textTracks".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) TextTracks() (ret TextTrackList, ok bool) {
	ok = js.True == bindings.GetSourceBufferTextTracks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppendWindowStart returns the value of property "SourceBuffer.appendWindowStart".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) AppendWindowStart() (ret float64, ok bool) {
	ok = js.True == bindings.GetSourceBufferAppendWindowStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAppendWindowStart sets the value of property "SourceBuffer.appendWindowStart" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetAppendWindowStart(val float64) bool {
	return js.True == bindings.SetSourceBufferAppendWindowStart(
		this.ref,
		float64(val),
	)
}

// AppendWindowEnd returns the value of property "SourceBuffer.appendWindowEnd".
//
// It returns ok=false if there is no such property.
func (this SourceBuffer) AppendWindowEnd() (ret float64, ok bool) {
	ok = js.True == bindings.GetSourceBufferAppendWindowEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAppendWindowEnd sets the value of property "SourceBuffer.appendWindowEnd" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetAppendWindowEnd(val float64) bool {
	return js.True == bindings.SetSourceBufferAppendWindowEnd(
		this.ref,
		float64(val),
	)
}

// HasFuncAppendBuffer returns true if the method "SourceBuffer.appendBuffer" exists.
func (this SourceBuffer) HasFuncAppendBuffer() bool {
	return js.True == bindings.HasFuncSourceBufferAppendBuffer(
		this.ref,
	)
}

// FuncAppendBuffer returns the method "SourceBuffer.appendBuffer".
func (this SourceBuffer) FuncAppendBuffer() (fn js.Func[func(data BufferSource)]) {
	bindings.FuncSourceBufferAppendBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AppendBuffer calls the method "SourceBuffer.appendBuffer".
func (this SourceBuffer) AppendBuffer(data BufferSource) (ret js.Void) {
	bindings.CallSourceBufferAppendBuffer(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryAppendBuffer calls the method "SourceBuffer.appendBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SourceBuffer) TryAppendBuffer(data BufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySourceBufferAppendBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncAbort returns true if the method "SourceBuffer.abort" exists.
func (this SourceBuffer) HasFuncAbort() bool {
	return js.True == bindings.HasFuncSourceBufferAbort(
		this.ref,
	)
}

// FuncAbort returns the method "SourceBuffer.abort".
func (this SourceBuffer) FuncAbort() (fn js.Func[func()]) {
	bindings.FuncSourceBufferAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "SourceBuffer.abort".
func (this SourceBuffer) Abort() (ret js.Void) {
	bindings.CallSourceBufferAbort(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "SourceBuffer.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SourceBuffer) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySourceBufferAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncChangeType returns true if the method "SourceBuffer.changeType" exists.
func (this SourceBuffer) HasFuncChangeType() bool {
	return js.True == bindings.HasFuncSourceBufferChangeType(
		this.ref,
	)
}

// FuncChangeType returns the method "SourceBuffer.changeType".
func (this SourceBuffer) FuncChangeType() (fn js.Func[func(typ js.String)]) {
	bindings.FuncSourceBufferChangeType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ChangeType calls the method "SourceBuffer.changeType".
func (this SourceBuffer) ChangeType(typ js.String) (ret js.Void) {
	bindings.CallSourceBufferChangeType(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryChangeType calls the method "SourceBuffer.changeType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SourceBuffer) TryChangeType(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySourceBufferChangeType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFuncRemove returns true if the method "SourceBuffer.remove" exists.
func (this SourceBuffer) HasFuncRemove() bool {
	return js.True == bindings.HasFuncSourceBufferRemove(
		this.ref,
	)
}

// FuncRemove returns the method "SourceBuffer.remove".
func (this SourceBuffer) FuncRemove() (fn js.Func[func(start float64, end float64)]) {
	bindings.FuncSourceBufferRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "SourceBuffer.remove".
func (this SourceBuffer) Remove(start float64, end float64) (ret js.Void) {
	bindings.CallSourceBufferRemove(
		this.ref, js.Pointer(&ret),
		float64(start),
		float64(end),
	)

	return
}

// TryRemove calls the method "SourceBuffer.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SourceBuffer) TryRemove(start float64, end float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySourceBufferRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
	)

	return
}

type TextTrack struct {
	EventTarget
}

func (this TextTrack) Once() TextTrack {
	this.ref.Once()
	return this
}

func (this TextTrack) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this TextTrack) FromRef(ref js.Ref) TextTrack {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this TextTrack) Free() {
	this.ref.Free()
}

// Kind returns the value of property "TextTrack.kind".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Kind() (ret TextTrackKind, ok bool) {
	ok = js.True == bindings.GetTextTrackKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "TextTrack.label".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "TextTrack.language".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "TextTrack.id".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InBandMetadataTrackDispatchType returns the value of property "TextTrack.inBandMetadataTrackDispatchType".
//
// It returns ok=false if there is no such property.
func (this TextTrack) InBandMetadataTrackDispatchType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackInBandMetadataTrackDispatchType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "TextTrack.mode".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Mode() (ret TextTrackMode, ok bool) {
	ok = js.True == bindings.GetTextTrackMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMode sets the value of property "TextTrack.mode" to val.
//
// It returns false if the property cannot be set.
func (this TextTrack) SetMode(val TextTrackMode) bool {
	return js.True == bindings.SetTextTrackMode(
		this.ref,
		uint32(val),
	)
}

// Cues returns the value of property "TextTrack.cues".
//
// It returns ok=false if there is no such property.
func (this TextTrack) Cues() (ret TextTrackCueList, ok bool) {
	ok = js.True == bindings.GetTextTrackCues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ActiveCues returns the value of property "TextTrack.activeCues".
//
// It returns ok=false if there is no such property.
func (this TextTrack) ActiveCues() (ret TextTrackCueList, ok bool) {
	ok = js.True == bindings.GetTextTrackActiveCues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SourceBuffer returns the value of property "TextTrack.sourceBuffer".
//
// It returns ok=false if there is no such property.
func (this TextTrack) SourceBuffer() (ret SourceBuffer, ok bool) {
	ok = js.True == bindings.GetTextTrackSourceBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAddCue returns true if the method "TextTrack.addCue" exists.
func (this TextTrack) HasFuncAddCue() bool {
	return js.True == bindings.HasFuncTextTrackAddCue(
		this.ref,
	)
}

// FuncAddCue returns the method "TextTrack.addCue".
func (this TextTrack) FuncAddCue() (fn js.Func[func(cue TextTrackCue)]) {
	bindings.FuncTextTrackAddCue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddCue calls the method "TextTrack.addCue".
func (this TextTrack) AddCue(cue TextTrackCue) (ret js.Void) {
	bindings.CallTextTrackAddCue(
		this.ref, js.Pointer(&ret),
		cue.Ref(),
	)

	return
}

// TryAddCue calls the method "TextTrack.addCue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrack) TryAddCue(cue TextTrackCue) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackAddCue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		cue.Ref(),
	)

	return
}

// HasFuncRemoveCue returns true if the method "TextTrack.removeCue" exists.
func (this TextTrack) HasFuncRemoveCue() bool {
	return js.True == bindings.HasFuncTextTrackRemoveCue(
		this.ref,
	)
}

// FuncRemoveCue returns the method "TextTrack.removeCue".
func (this TextTrack) FuncRemoveCue() (fn js.Func[func(cue TextTrackCue)]) {
	bindings.FuncTextTrackRemoveCue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveCue calls the method "TextTrack.removeCue".
func (this TextTrack) RemoveCue(cue TextTrackCue) (ret js.Void) {
	bindings.CallTextTrackRemoveCue(
		this.ref, js.Pointer(&ret),
		cue.Ref(),
	)

	return
}

// TryRemoveCue calls the method "TextTrack.removeCue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TextTrack) TryRemoveCue(cue TextTrackCue) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTextTrackRemoveCue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		cue.Ref(),
	)

	return
}

type MediaKeySessionClosedReason uint32

const (
	_ MediaKeySessionClosedReason = iota

	MediaKeySessionClosedReason_INTERNAL_ERROR
	MediaKeySessionClosedReason_CLOSED_BY_APPLICATION
	MediaKeySessionClosedReason_RELEASE_ACKNOWLEDGED
	MediaKeySessionClosedReason_HARDWARE_CONTEXT_RESET
	MediaKeySessionClosedReason_RESOURCE_EVICTED
)

func (MediaKeySessionClosedReason) FromRef(str js.Ref) MediaKeySessionClosedReason {
	return MediaKeySessionClosedReason(bindings.ConstOfMediaKeySessionClosedReason(str))
}

func (x MediaKeySessionClosedReason) String() (string, bool) {
	switch x {
	case MediaKeySessionClosedReason_INTERNAL_ERROR:
		return "internal-error", true
	case MediaKeySessionClosedReason_CLOSED_BY_APPLICATION:
		return "closed-by-application", true
	case MediaKeySessionClosedReason_RELEASE_ACKNOWLEDGED:
		return "release-acknowledged", true
	case MediaKeySessionClosedReason_HARDWARE_CONTEXT_RESET:
		return "hardware-context-reset", true
	case MediaKeySessionClosedReason_RESOURCE_EVICTED:
		return "resource-evicted", true
	default:
		return "", false
	}
}

type MediaKeyStatus uint32

const (
	_ MediaKeyStatus = iota

	MediaKeyStatus_USABLE
	MediaKeyStatus_EXPIRED
	MediaKeyStatus_RELEASED
	MediaKeyStatus_OUTPUT_RESTRICTED
	MediaKeyStatus_OUTPUT_DOWNSCALED
	MediaKeyStatus_USABLE_IN_FUTURE
	MediaKeyStatus_STATUS_PENDING
	MediaKeyStatus_INTERNAL_ERROR
)

func (MediaKeyStatus) FromRef(str js.Ref) MediaKeyStatus {
	return MediaKeyStatus(bindings.ConstOfMediaKeyStatus(str))
}

func (x MediaKeyStatus) String() (string, bool) {
	switch x {
	case MediaKeyStatus_USABLE:
		return "usable", true
	case MediaKeyStatus_EXPIRED:
		return "expired", true
	case MediaKeyStatus_RELEASED:
		return "released", true
	case MediaKeyStatus_OUTPUT_RESTRICTED:
		return "output-restricted", true
	case MediaKeyStatus_OUTPUT_DOWNSCALED:
		return "output-downscaled", true
	case MediaKeyStatus_USABLE_IN_FUTURE:
		return "usable-in-future", true
	case MediaKeyStatus_STATUS_PENDING:
		return "status-pending", true
	case MediaKeyStatus_INTERNAL_ERROR:
		return "internal-error", true
	default:
		return "", false
	}
}

type OneOf_MediaKeyStatus_undefined struct {
	ref js.Ref
}

func (x OneOf_MediaKeyStatus_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_MediaKeyStatus_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_MediaKeyStatus_undefined) FromRef(ref js.Ref) OneOf_MediaKeyStatus_undefined {
	return OneOf_MediaKeyStatus_undefined{
		ref: ref,
	}
}

func (x OneOf_MediaKeyStatus_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_MediaKeyStatus_undefined) MediaKeyStatus() MediaKeyStatus {
	return MediaKeyStatus(0).FromRef(x.ref)
}

type MediaKeyStatusMap struct {
	ref js.Ref
}

func (this MediaKeyStatusMap) Once() MediaKeyStatusMap {
	this.ref.Once()
	return this
}

func (this MediaKeyStatusMap) Ref() js.Ref {
	return this.ref
}

func (this MediaKeyStatusMap) FromRef(ref js.Ref) MediaKeyStatusMap {
	this.ref = ref
	return this
}

func (this MediaKeyStatusMap) Free() {
	this.ref.Free()
}

// Size returns the value of property "MediaKeyStatusMap.size".
//
// It returns ok=false if there is no such property.
func (this MediaKeyStatusMap) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaKeyStatusMapSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncHas returns true if the method "MediaKeyStatusMap.has" exists.
func (this MediaKeyStatusMap) HasFuncHas() bool {
	return js.True == bindings.HasFuncMediaKeyStatusMapHas(
		this.ref,
	)
}

// FuncHas returns the method "MediaKeyStatusMap.has".
func (this MediaKeyStatusMap) FuncHas() (fn js.Func[func(keyId BufferSource) bool]) {
	bindings.FuncMediaKeyStatusMapHas(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Has calls the method "MediaKeyStatusMap.has".
func (this MediaKeyStatusMap) Has(keyId BufferSource) (ret bool) {
	bindings.CallMediaKeyStatusMapHas(
		this.ref, js.Pointer(&ret),
		keyId.Ref(),
	)

	return
}

// TryHas calls the method "MediaKeyStatusMap.has"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeyStatusMap) TryHas(keyId BufferSource) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeyStatusMapHas(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keyId.Ref(),
	)

	return
}

// HasFuncGet returns true if the method "MediaKeyStatusMap.get" exists.
func (this MediaKeyStatusMap) HasFuncGet() bool {
	return js.True == bindings.HasFuncMediaKeyStatusMapGet(
		this.ref,
	)
}

// FuncGet returns the method "MediaKeyStatusMap.get".
func (this MediaKeyStatusMap) FuncGet() (fn js.Func[func(keyId BufferSource) OneOf_MediaKeyStatus_undefined]) {
	bindings.FuncMediaKeyStatusMapGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "MediaKeyStatusMap.get".
func (this MediaKeyStatusMap) Get(keyId BufferSource) (ret OneOf_MediaKeyStatus_undefined) {
	bindings.CallMediaKeyStatusMapGet(
		this.ref, js.Pointer(&ret),
		keyId.Ref(),
	)

	return
}

// TryGet calls the method "MediaKeyStatusMap.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeyStatusMap) TryGet(keyId BufferSource) (ret OneOf_MediaKeyStatus_undefined, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeyStatusMapGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keyId.Ref(),
	)

	return
}

type MediaKeySession struct {
	EventTarget
}

func (this MediaKeySession) Once() MediaKeySession {
	this.ref.Once()
	return this
}

func (this MediaKeySession) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaKeySession) FromRef(ref js.Ref) MediaKeySession {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaKeySession) Free() {
	this.ref.Free()
}

// SessionId returns the value of property "MediaKeySession.sessionId".
//
// It returns ok=false if there is no such property.
func (this MediaKeySession) SessionId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaKeySessionSessionId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Expiration returns the value of property "MediaKeySession.expiration".
//
// It returns ok=false if there is no such property.
func (this MediaKeySession) Expiration() (ret float64, ok bool) {
	ok = js.True == bindings.GetMediaKeySessionExpiration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Closed returns the value of property "MediaKeySession.closed".
//
// It returns ok=false if there is no such property.
func (this MediaKeySession) Closed() (ret js.Promise[MediaKeySessionClosedReason], ok bool) {
	ok = js.True == bindings.GetMediaKeySessionClosed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KeyStatuses returns the value of property "MediaKeySession.keyStatuses".
//
// It returns ok=false if there is no such property.
func (this MediaKeySession) KeyStatuses() (ret MediaKeyStatusMap, ok bool) {
	ok = js.True == bindings.GetMediaKeySessionKeyStatuses(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGenerateRequest returns true if the method "MediaKeySession.generateRequest" exists.
func (this MediaKeySession) HasFuncGenerateRequest() bool {
	return js.True == bindings.HasFuncMediaKeySessionGenerateRequest(
		this.ref,
	)
}

// FuncGenerateRequest returns the method "MediaKeySession.generateRequest".
func (this MediaKeySession) FuncGenerateRequest() (fn js.Func[func(initDataType js.String, initData BufferSource) js.Promise[js.Void]]) {
	bindings.FuncMediaKeySessionGenerateRequest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GenerateRequest calls the method "MediaKeySession.generateRequest".
func (this MediaKeySession) GenerateRequest(initDataType js.String, initData BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallMediaKeySessionGenerateRequest(
		this.ref, js.Pointer(&ret),
		initDataType.Ref(),
		initData.Ref(),
	)

	return
}

// TryGenerateRequest calls the method "MediaKeySession.generateRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySession) TryGenerateRequest(initDataType js.String, initData BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySessionGenerateRequest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		initDataType.Ref(),
		initData.Ref(),
	)

	return
}

// HasFuncLoad returns true if the method "MediaKeySession.load" exists.
func (this MediaKeySession) HasFuncLoad() bool {
	return js.True == bindings.HasFuncMediaKeySessionLoad(
		this.ref,
	)
}

// FuncLoad returns the method "MediaKeySession.load".
func (this MediaKeySession) FuncLoad() (fn js.Func[func(sessionId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncMediaKeySessionLoad(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Load calls the method "MediaKeySession.load".
func (this MediaKeySession) Load(sessionId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallMediaKeySessionLoad(
		this.ref, js.Pointer(&ret),
		sessionId.Ref(),
	)

	return
}

// TryLoad calls the method "MediaKeySession.load"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySession) TryLoad(sessionId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySessionLoad(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		sessionId.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the method "MediaKeySession.update" exists.
func (this MediaKeySession) HasFuncUpdate() bool {
	return js.True == bindings.HasFuncMediaKeySessionUpdate(
		this.ref,
	)
}

// FuncUpdate returns the method "MediaKeySession.update".
func (this MediaKeySession) FuncUpdate() (fn js.Func[func(response BufferSource) js.Promise[js.Void]]) {
	bindings.FuncMediaKeySessionUpdate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Update calls the method "MediaKeySession.update".
func (this MediaKeySession) Update(response BufferSource) (ret js.Promise[js.Void]) {
	bindings.CallMediaKeySessionUpdate(
		this.ref, js.Pointer(&ret),
		response.Ref(),
	)

	return
}

// TryUpdate calls the method "MediaKeySession.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySession) TryUpdate(response BufferSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySessionUpdate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		response.Ref(),
	)

	return
}

// HasFuncClose returns true if the method "MediaKeySession.close" exists.
func (this MediaKeySession) HasFuncClose() bool {
	return js.True == bindings.HasFuncMediaKeySessionClose(
		this.ref,
	)
}

// FuncClose returns the method "MediaKeySession.close".
func (this MediaKeySession) FuncClose() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncMediaKeySessionClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "MediaKeySession.close".
func (this MediaKeySession) Close() (ret js.Promise[js.Void]) {
	bindings.CallMediaKeySessionClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "MediaKeySession.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySession) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySessionClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemove returns true if the method "MediaKeySession.remove" exists.
func (this MediaKeySession) HasFuncRemove() bool {
	return js.True == bindings.HasFuncMediaKeySessionRemove(
		this.ref,
	)
}

// FuncRemove returns the method "MediaKeySession.remove".
func (this MediaKeySession) FuncRemove() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncMediaKeySessionRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "MediaKeySession.remove".
func (this MediaKeySession) Remove() (ret js.Promise[js.Void]) {
	bindings.CallMediaKeySessionRemove(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "MediaKeySession.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySession) TryRemove() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySessionRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MediaKeySessionType uint32

const (
	_ MediaKeySessionType = iota

	MediaKeySessionType_TEMPORARY
	MediaKeySessionType_PERSISTENT_LICENSE
)

func (MediaKeySessionType) FromRef(str js.Ref) MediaKeySessionType {
	return MediaKeySessionType(bindings.ConstOfMediaKeySessionType(str))
}

func (x MediaKeySessionType) String() (string, bool) {
	switch x {
	case MediaKeySessionType_TEMPORARY:
		return "temporary", true
	case MediaKeySessionType_PERSISTENT_LICENSE:
		return "persistent-license", true
	default:
		return "", false
	}
}

type MediaKeys struct {
	ref js.Ref
}

func (this MediaKeys) Once() MediaKeys {
	this.ref.Once()
	return this
}

func (this MediaKeys) Ref() js.Ref {
	return this.ref
}

func (this MediaKeys) FromRef(ref js.Ref) MediaKeys {
	this.ref = ref
	return this
}

func (this MediaKeys) Free() {
	this.ref.Free()
}

// HasFuncCreateSession returns true if the method "MediaKeys.createSession" exists.
func (this MediaKeys) HasFuncCreateSession() bool {
	return js.True == bindings.HasFuncMediaKeysCreateSession(
		this.ref,
	)
}

// FuncCreateSession returns the method "MediaKeys.createSession".
func (this MediaKeys) FuncCreateSession() (fn js.Func[func(sessionType MediaKeySessionType) MediaKeySession]) {
	bindings.FuncMediaKeysCreateSession(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSession calls the method "MediaKeys.createSession".
func (this MediaKeys) CreateSession(sessionType MediaKeySessionType) (ret MediaKeySession) {
	bindings.CallMediaKeysCreateSession(
		this.ref, js.Pointer(&ret),
		uint32(sessionType),
	)

	return
}

// TryCreateSession calls the method "MediaKeys.createSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeys) TryCreateSession(sessionType MediaKeySessionType) (ret MediaKeySession, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeysCreateSession(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(sessionType),
	)

	return
}

// HasFuncCreateSession1 returns true if the method "MediaKeys.createSession" exists.
func (this MediaKeys) HasFuncCreateSession1() bool {
	return js.True == bindings.HasFuncMediaKeysCreateSession1(
		this.ref,
	)
}

// FuncCreateSession1 returns the method "MediaKeys.createSession".
func (this MediaKeys) FuncCreateSession1() (fn js.Func[func() MediaKeySession]) {
	bindings.FuncMediaKeysCreateSession1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSession1 calls the method "MediaKeys.createSession".
func (this MediaKeys) CreateSession1() (ret MediaKeySession) {
	bindings.CallMediaKeysCreateSession1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSession1 calls the method "MediaKeys.createSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeys) TryCreateSession1() (ret MediaKeySession, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeysCreateSession1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetServerCertificate returns true if the method "MediaKeys.setServerCertificate" exists.
func (this MediaKeys) HasFuncSetServerCertificate() bool {
	return js.True == bindings.HasFuncMediaKeysSetServerCertificate(
		this.ref,
	)
}

// FuncSetServerCertificate returns the method "MediaKeys.setServerCertificate".
func (this MediaKeys) FuncSetServerCertificate() (fn js.Func[func(serverCertificate BufferSource) js.Promise[js.Boolean]]) {
	bindings.FuncMediaKeysSetServerCertificate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetServerCertificate calls the method "MediaKeys.setServerCertificate".
func (this MediaKeys) SetServerCertificate(serverCertificate BufferSource) (ret js.Promise[js.Boolean]) {
	bindings.CallMediaKeysSetServerCertificate(
		this.ref, js.Pointer(&ret),
		serverCertificate.Ref(),
	)

	return
}

// TrySetServerCertificate calls the method "MediaKeys.setServerCertificate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeys) TrySetServerCertificate(serverCertificate BufferSource) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeysSetServerCertificate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		serverCertificate.Ref(),
	)

	return
}

type ULongRange struct {
	// Max is "ULongRange.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max uint32
	// Min is "ULongRange.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min uint32

	FFI_USE_Max bool // for Max.
	FFI_USE_Min bool // for Min.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ULongRange with all fields set.
func (p ULongRange) FromRef(ref js.Ref) ULongRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ULongRange in the application heap.
func (p ULongRange) New() js.Ref {
	return bindings.ULongRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ULongRange) UpdateFrom(ref js.Ref) {
	bindings.ULongRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ULongRange) Update(ref js.Ref) {
	bindings.ULongRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ULongRange) FreeMembers(recursive bool) {
}

type DoubleRange struct {
	// Max is "DoubleRange.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max float64
	// Min is "DoubleRange.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min float64

	FFI_USE_Max bool // for Max.
	FFI_USE_Min bool // for Min.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DoubleRange with all fields set.
func (p DoubleRange) FromRef(ref js.Ref) DoubleRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DoubleRange in the application heap.
func (p DoubleRange) New() js.Ref {
	return bindings.DoubleRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DoubleRange) UpdateFrom(ref js.Ref) {
	bindings.DoubleRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DoubleRange) Update(ref js.Ref) {
	bindings.DoubleRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DoubleRange) FreeMembers(recursive bool) {
}

type MediaSettingsRange struct {
	// Max is "MediaSettingsRange.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max float64
	// Min is "MediaSettingsRange.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min float64
	// Step is "MediaSettingsRange.step"
	//
	// Optional
	//
	// NOTE: FFI_USE_Step MUST be set to true to make this field effective.
	Step float64

	FFI_USE_Max  bool // for Max.
	FFI_USE_Min  bool // for Min.
	FFI_USE_Step bool // for Step.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaSettingsRange with all fields set.
func (p MediaSettingsRange) FromRef(ref js.Ref) MediaSettingsRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaSettingsRange in the application heap.
func (p MediaSettingsRange) New() js.Ref {
	return bindings.MediaSettingsRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaSettingsRange) UpdateFrom(ref js.Ref) {
	bindings.MediaSettingsRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaSettingsRange) Update(ref js.Ref) {
	bindings.MediaSettingsRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaSettingsRange) FreeMembers(recursive bool) {
}

type MediaTrackCapabilities struct {
	// Width is "MediaTrackCapabilities.width"
	//
	// Optional
	//
	// NOTE: Width.FFI_USE MUST be set to true to get Width used.
	Width ULongRange
	// Height is "MediaTrackCapabilities.height"
	//
	// Optional
	//
	// NOTE: Height.FFI_USE MUST be set to true to get Height used.
	Height ULongRange
	// AspectRatio is "MediaTrackCapabilities.aspectRatio"
	//
	// Optional
	//
	// NOTE: AspectRatio.FFI_USE MUST be set to true to get AspectRatio used.
	AspectRatio DoubleRange
	// FrameRate is "MediaTrackCapabilities.frameRate"
	//
	// Optional
	//
	// NOTE: FrameRate.FFI_USE MUST be set to true to get FrameRate used.
	FrameRate DoubleRange
	// FacingMode is "MediaTrackCapabilities.facingMode"
	//
	// Optional
	FacingMode js.Array[js.String]
	// ResizeMode is "MediaTrackCapabilities.resizeMode"
	//
	// Optional
	ResizeMode js.Array[js.String]
	// SampleRate is "MediaTrackCapabilities.sampleRate"
	//
	// Optional
	//
	// NOTE: SampleRate.FFI_USE MUST be set to true to get SampleRate used.
	SampleRate ULongRange
	// SampleSize is "MediaTrackCapabilities.sampleSize"
	//
	// Optional
	//
	// NOTE: SampleSize.FFI_USE MUST be set to true to get SampleSize used.
	SampleSize ULongRange
	// EchoCancellation is "MediaTrackCapabilities.echoCancellation"
	//
	// Optional
	EchoCancellation js.Array[bool]
	// AutoGainControl is "MediaTrackCapabilities.autoGainControl"
	//
	// Optional
	AutoGainControl js.Array[bool]
	// NoiseSuppression is "MediaTrackCapabilities.noiseSuppression"
	//
	// Optional
	NoiseSuppression js.Array[bool]
	// Latency is "MediaTrackCapabilities.latency"
	//
	// Optional
	//
	// NOTE: Latency.FFI_USE MUST be set to true to get Latency used.
	Latency DoubleRange
	// ChannelCount is "MediaTrackCapabilities.channelCount"
	//
	// Optional
	//
	// NOTE: ChannelCount.FFI_USE MUST be set to true to get ChannelCount used.
	ChannelCount ULongRange
	// DeviceId is "MediaTrackCapabilities.deviceId"
	//
	// Optional
	DeviceId js.String
	// GroupId is "MediaTrackCapabilities.groupId"
	//
	// Optional
	GroupId js.String
	// WhiteBalanceMode is "MediaTrackCapabilities.whiteBalanceMode"
	//
	// Optional
	WhiteBalanceMode js.Array[js.String]
	// ExposureMode is "MediaTrackCapabilities.exposureMode"
	//
	// Optional
	ExposureMode js.Array[js.String]
	// FocusMode is "MediaTrackCapabilities.focusMode"
	//
	// Optional
	FocusMode js.Array[js.String]
	// ExposureCompensation is "MediaTrackCapabilities.exposureCompensation"
	//
	// Optional
	//
	// NOTE: ExposureCompensation.FFI_USE MUST be set to true to get ExposureCompensation used.
	ExposureCompensation MediaSettingsRange
	// ExposureTime is "MediaTrackCapabilities.exposureTime"
	//
	// Optional
	//
	// NOTE: ExposureTime.FFI_USE MUST be set to true to get ExposureTime used.
	ExposureTime MediaSettingsRange
	// ColorTemperature is "MediaTrackCapabilities.colorTemperature"
	//
	// Optional
	//
	// NOTE: ColorTemperature.FFI_USE MUST be set to true to get ColorTemperature used.
	ColorTemperature MediaSettingsRange
	// Iso is "MediaTrackCapabilities.iso"
	//
	// Optional
	//
	// NOTE: Iso.FFI_USE MUST be set to true to get Iso used.
	Iso MediaSettingsRange
	// Brightness is "MediaTrackCapabilities.brightness"
	//
	// Optional
	//
	// NOTE: Brightness.FFI_USE MUST be set to true to get Brightness used.
	Brightness MediaSettingsRange
	// Contrast is "MediaTrackCapabilities.contrast"
	//
	// Optional
	//
	// NOTE: Contrast.FFI_USE MUST be set to true to get Contrast used.
	Contrast MediaSettingsRange
	// Saturation is "MediaTrackCapabilities.saturation"
	//
	// Optional
	//
	// NOTE: Saturation.FFI_USE MUST be set to true to get Saturation used.
	Saturation MediaSettingsRange
	// Sharpness is "MediaTrackCapabilities.sharpness"
	//
	// Optional
	//
	// NOTE: Sharpness.FFI_USE MUST be set to true to get Sharpness used.
	Sharpness MediaSettingsRange
	// FocusDistance is "MediaTrackCapabilities.focusDistance"
	//
	// Optional
	//
	// NOTE: FocusDistance.FFI_USE MUST be set to true to get FocusDistance used.
	FocusDistance MediaSettingsRange
	// Pan is "MediaTrackCapabilities.pan"
	//
	// Optional
	//
	// NOTE: Pan.FFI_USE MUST be set to true to get Pan used.
	Pan MediaSettingsRange
	// Tilt is "MediaTrackCapabilities.tilt"
	//
	// Optional
	//
	// NOTE: Tilt.FFI_USE MUST be set to true to get Tilt used.
	Tilt MediaSettingsRange
	// Zoom is "MediaTrackCapabilities.zoom"
	//
	// Optional
	//
	// NOTE: Zoom.FFI_USE MUST be set to true to get Zoom used.
	Zoom MediaSettingsRange
	// Torch is "MediaTrackCapabilities.torch"
	//
	// Optional
	//
	// NOTE: FFI_USE_Torch MUST be set to true to make this field effective.
	Torch bool
	// DisplaySurface is "MediaTrackCapabilities.displaySurface"
	//
	// Optional
	DisplaySurface js.String
	// LogicalSurface is "MediaTrackCapabilities.logicalSurface"
	//
	// Optional
	//
	// NOTE: FFI_USE_LogicalSurface MUST be set to true to make this field effective.
	LogicalSurface bool
	// Cursor is "MediaTrackCapabilities.cursor"
	//
	// Optional
	Cursor js.Array[js.String]

	FFI_USE_Torch          bool // for Torch.
	FFI_USE_LogicalSurface bool // for LogicalSurface.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaTrackCapabilities with all fields set.
func (p MediaTrackCapabilities) FromRef(ref js.Ref) MediaTrackCapabilities {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaTrackCapabilities in the application heap.
func (p MediaTrackCapabilities) New() js.Ref {
	return bindings.MediaTrackCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaTrackCapabilities) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackCapabilitiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaTrackCapabilities) Update(ref js.Ref) {
	bindings.MediaTrackCapabilitiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaTrackCapabilities) FreeMembers(recursive bool) {
	js.Free(
		p.FacingMode.Ref(),
		p.ResizeMode.Ref(),
		p.EchoCancellation.Ref(),
		p.AutoGainControl.Ref(),
		p.NoiseSuppression.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
		p.WhiteBalanceMode.Ref(),
		p.ExposureMode.Ref(),
		p.FocusMode.Ref(),
		p.DisplaySurface.Ref(),
		p.Cursor.Ref(),
	)
	p.FacingMode = p.FacingMode.FromRef(js.Undefined)
	p.ResizeMode = p.ResizeMode.FromRef(js.Undefined)
	p.EchoCancellation = p.EchoCancellation.FromRef(js.Undefined)
	p.AutoGainControl = p.AutoGainControl.FromRef(js.Undefined)
	p.NoiseSuppression = p.NoiseSuppression.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
	p.WhiteBalanceMode = p.WhiteBalanceMode.FromRef(js.Undefined)
	p.ExposureMode = p.ExposureMode.FromRef(js.Undefined)
	p.FocusMode = p.FocusMode.FromRef(js.Undefined)
	p.DisplaySurface = p.DisplaySurface.FromRef(js.Undefined)
	p.Cursor = p.Cursor.FromRef(js.Undefined)
	if recursive {
		p.Width.FreeMembers(true)
		p.Height.FreeMembers(true)
		p.AspectRatio.FreeMembers(true)
		p.FrameRate.FreeMembers(true)
		p.SampleRate.FreeMembers(true)
		p.SampleSize.FreeMembers(true)
		p.Latency.FreeMembers(true)
		p.ChannelCount.FreeMembers(true)
		p.ExposureCompensation.FreeMembers(true)
		p.ExposureTime.FreeMembers(true)
		p.ColorTemperature.FreeMembers(true)
		p.Iso.FreeMembers(true)
		p.Brightness.FreeMembers(true)
		p.Contrast.FreeMembers(true)
		p.Saturation.FreeMembers(true)
		p.Sharpness.FreeMembers(true)
		p.FocusDistance.FreeMembers(true)
		p.Pan.FreeMembers(true)
		p.Tilt.FreeMembers(true)
		p.Zoom.FreeMembers(true)
	}
}

type ConstrainULongRange struct {
	// Exact is "ConstrainULongRange.exact"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exact MUST be set to true to make this field effective.
	Exact uint32
	// Ideal is "ConstrainULongRange.ideal"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ideal MUST be set to true to make this field effective.
	Ideal uint32
	// Max is "ConstrainULongRange.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max uint32
	// Min is "ConstrainULongRange.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min uint32

	FFI_USE_Exact bool // for Exact.
	FFI_USE_Ideal bool // for Ideal.
	FFI_USE_Max   bool // for Max.
	FFI_USE_Min   bool // for Min.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstrainULongRange with all fields set.
func (p ConstrainULongRange) FromRef(ref js.Ref) ConstrainULongRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstrainULongRange in the application heap.
func (p ConstrainULongRange) New() js.Ref {
	return bindings.ConstrainULongRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstrainULongRange) UpdateFrom(ref js.Ref) {
	bindings.ConstrainULongRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstrainULongRange) Update(ref js.Ref) {
	bindings.ConstrainULongRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstrainULongRange) FreeMembers(recursive bool) {
}

type OneOf_Uint32_ConstrainULongRange struct {
	ref js.Ref
}

func (x OneOf_Uint32_ConstrainULongRange) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Uint32_ConstrainULongRange) Free() {
	x.ref.Free()
}

func (x OneOf_Uint32_ConstrainULongRange) FromRef(ref js.Ref) OneOf_Uint32_ConstrainULongRange {
	return OneOf_Uint32_ConstrainULongRange{
		ref: ref,
	}
}

func (x OneOf_Uint32_ConstrainULongRange) Uint32() uint32 {
	return js.Number[uint32]{}.FromRef(x.ref).Get()
}

func (x OneOf_Uint32_ConstrainULongRange) ConstrainULongRange() ConstrainULongRange {
	var ret ConstrainULongRange
	ret.UpdateFrom(x.ref)
	return ret
}

type ConstrainULong = OneOf_Uint32_ConstrainULongRange

type ConstrainDoubleRange struct {
	// Exact is "ConstrainDoubleRange.exact"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exact MUST be set to true to make this field effective.
	Exact float64
	// Ideal is "ConstrainDoubleRange.ideal"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ideal MUST be set to true to make this field effective.
	Ideal float64
	// Max is "ConstrainDoubleRange.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max float64
	// Min is "ConstrainDoubleRange.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min float64

	FFI_USE_Exact bool // for Exact.
	FFI_USE_Ideal bool // for Ideal.
	FFI_USE_Max   bool // for Max.
	FFI_USE_Min   bool // for Min.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstrainDoubleRange with all fields set.
func (p ConstrainDoubleRange) FromRef(ref js.Ref) ConstrainDoubleRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstrainDoubleRange in the application heap.
func (p ConstrainDoubleRange) New() js.Ref {
	return bindings.ConstrainDoubleRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstrainDoubleRange) UpdateFrom(ref js.Ref) {
	bindings.ConstrainDoubleRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstrainDoubleRange) Update(ref js.Ref) {
	bindings.ConstrainDoubleRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstrainDoubleRange) FreeMembers(recursive bool) {
}

type OneOf_Float64_ConstrainDoubleRange struct {
	ref js.Ref
}

func (x OneOf_Float64_ConstrainDoubleRange) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_ConstrainDoubleRange) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_ConstrainDoubleRange) FromRef(ref js.Ref) OneOf_Float64_ConstrainDoubleRange {
	return OneOf_Float64_ConstrainDoubleRange{
		ref: ref,
	}
}

func (x OneOf_Float64_ConstrainDoubleRange) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_ConstrainDoubleRange) ConstrainDoubleRange() ConstrainDoubleRange {
	var ret ConstrainDoubleRange
	ret.UpdateFrom(x.ref)
	return ret
}

type ConstrainDouble = OneOf_Float64_ConstrainDoubleRange

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type ConstrainDOMStringParameters struct {
	// Exact is "ConstrainDOMStringParameters.exact"
	//
	// Optional
	Exact OneOf_String_ArrayString
	// Ideal is "ConstrainDOMStringParameters.ideal"
	//
	// Optional
	Ideal OneOf_String_ArrayString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstrainDOMStringParameters with all fields set.
func (p ConstrainDOMStringParameters) FromRef(ref js.Ref) ConstrainDOMStringParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstrainDOMStringParameters in the application heap.
func (p ConstrainDOMStringParameters) New() js.Ref {
	return bindings.ConstrainDOMStringParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstrainDOMStringParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainDOMStringParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstrainDOMStringParameters) Update(ref js.Ref) {
	bindings.ConstrainDOMStringParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstrainDOMStringParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Exact.Ref(),
		p.Ideal.Ref(),
	)
	p.Exact = p.Exact.FromRef(js.Undefined)
	p.Ideal = p.Ideal.FromRef(js.Undefined)
}

type OneOf_String_ArrayString_ConstrainDOMStringParameters struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) FromRef(ref js.Ref) OneOf_String_ArrayString_ConstrainDOMStringParameters {
	return OneOf_String_ArrayString_ConstrainDOMStringParameters{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString_ConstrainDOMStringParameters) ConstrainDOMStringParameters() ConstrainDOMStringParameters {
	var ret ConstrainDOMStringParameters
	ret.UpdateFrom(x.ref)
	return ret
}

type ConstrainDOMString = OneOf_String_ArrayString_ConstrainDOMStringParameters

type ConstrainBooleanParameters struct {
	// Exact is "ConstrainBooleanParameters.exact"
	//
	// Optional
	//
	// NOTE: FFI_USE_Exact MUST be set to true to make this field effective.
	Exact bool
	// Ideal is "ConstrainBooleanParameters.ideal"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ideal MUST be set to true to make this field effective.
	Ideal bool

	FFI_USE_Exact bool // for Exact.
	FFI_USE_Ideal bool // for Ideal.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstrainBooleanParameters with all fields set.
func (p ConstrainBooleanParameters) FromRef(ref js.Ref) ConstrainBooleanParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstrainBooleanParameters in the application heap.
func (p ConstrainBooleanParameters) New() js.Ref {
	return bindings.ConstrainBooleanParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstrainBooleanParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainBooleanParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstrainBooleanParameters) Update(ref js.Ref) {
	bindings.ConstrainBooleanParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstrainBooleanParameters) FreeMembers(recursive bool) {
}

type OneOf_Bool_ConstrainBooleanParameters struct {
	ref js.Ref
}

func (x OneOf_Bool_ConstrainBooleanParameters) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_ConstrainBooleanParameters) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_ConstrainBooleanParameters) FromRef(ref js.Ref) OneOf_Bool_ConstrainBooleanParameters {
	return OneOf_Bool_ConstrainBooleanParameters{
		ref: ref,
	}
}

func (x OneOf_Bool_ConstrainBooleanParameters) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_ConstrainBooleanParameters) ConstrainBooleanParameters() ConstrainBooleanParameters {
	var ret ConstrainBooleanParameters
	ret.UpdateFrom(x.ref)
	return ret
}

type ConstrainBoolean = OneOf_Bool_ConstrainBooleanParameters

type Point2D struct {
	// X is "Point2D.x"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "Point2D.y"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Point2D with all fields set.
func (p Point2D) FromRef(ref js.Ref) Point2D {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Point2D in the application heap.
func (p Point2D) New() js.Ref {
	return bindings.Point2DJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Point2D) UpdateFrom(ref js.Ref) {
	bindings.Point2DJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Point2D) Update(ref js.Ref) {
	bindings.Point2DJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Point2D) FreeMembers(recursive bool) {
}

type ConstrainPoint2DParameters struct {
	// Exact is "ConstrainPoint2DParameters.exact"
	//
	// Optional
	Exact js.Array[Point2D]
	// Ideal is "ConstrainPoint2DParameters.ideal"
	//
	// Optional
	Ideal js.Array[Point2D]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConstrainPoint2DParameters with all fields set.
func (p ConstrainPoint2DParameters) FromRef(ref js.Ref) ConstrainPoint2DParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConstrainPoint2DParameters in the application heap.
func (p ConstrainPoint2DParameters) New() js.Ref {
	return bindings.ConstrainPoint2DParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConstrainPoint2DParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainPoint2DParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConstrainPoint2DParameters) Update(ref js.Ref) {
	bindings.ConstrainPoint2DParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConstrainPoint2DParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Exact.Ref(),
		p.Ideal.Ref(),
	)
	p.Exact = p.Exact.FromRef(js.Undefined)
	p.Ideal = p.Ideal.FromRef(js.Undefined)
}

type OneOf_ArrayPoint2D_ConstrainPoint2DParameters struct {
	ref js.Ref
}

func (x OneOf_ArrayPoint2D_ConstrainPoint2DParameters) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayPoint2D_ConstrainPoint2DParameters) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayPoint2D_ConstrainPoint2DParameters) FromRef(ref js.Ref) OneOf_ArrayPoint2D_ConstrainPoint2DParameters {
	return OneOf_ArrayPoint2D_ConstrainPoint2DParameters{
		ref: ref,
	}
}

func (x OneOf_ArrayPoint2D_ConstrainPoint2DParameters) ArrayPoint2D() js.Array[Point2D] {
	return js.Array[Point2D]{}.FromRef(x.ref)
}

func (x OneOf_ArrayPoint2D_ConstrainPoint2DParameters) ConstrainPoint2DParameters() ConstrainPoint2DParameters {
	var ret ConstrainPoint2DParameters
	ret.UpdateFrom(x.ref)
	return ret
}

type ConstrainPoint2D = OneOf_ArrayPoint2D_ConstrainPoint2DParameters

type OneOf_Bool_Float64_ConstrainDoubleRange struct {
	ref js.Ref
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) FromRef(ref js.Ref) OneOf_Bool_Float64_ConstrainDoubleRange {
	return OneOf_Bool_Float64_ConstrainDoubleRange{
		ref: ref,
	}
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Bool_Float64_ConstrainDoubleRange) ConstrainDoubleRange() ConstrainDoubleRange {
	var ret ConstrainDoubleRange
	ret.UpdateFrom(x.ref)
	return ret
}

type MediaTrackConstraintSet struct {
	// Width is "MediaTrackConstraintSet.width"
	//
	// Optional
	Width ConstrainULong
	// Height is "MediaTrackConstraintSet.height"
	//
	// Optional
	Height ConstrainULong
	// AspectRatio is "MediaTrackConstraintSet.aspectRatio"
	//
	// Optional
	AspectRatio ConstrainDouble
	// FrameRate is "MediaTrackConstraintSet.frameRate"
	//
	// Optional
	FrameRate ConstrainDouble
	// FacingMode is "MediaTrackConstraintSet.facingMode"
	//
	// Optional
	FacingMode ConstrainDOMString
	// ResizeMode is "MediaTrackConstraintSet.resizeMode"
	//
	// Optional
	ResizeMode ConstrainDOMString
	// SampleRate is "MediaTrackConstraintSet.sampleRate"
	//
	// Optional
	SampleRate ConstrainULong
	// SampleSize is "MediaTrackConstraintSet.sampleSize"
	//
	// Optional
	SampleSize ConstrainULong
	// EchoCancellation is "MediaTrackConstraintSet.echoCancellation"
	//
	// Optional
	EchoCancellation ConstrainBoolean
	// AutoGainControl is "MediaTrackConstraintSet.autoGainControl"
	//
	// Optional
	AutoGainControl ConstrainBoolean
	// NoiseSuppression is "MediaTrackConstraintSet.noiseSuppression"
	//
	// Optional
	NoiseSuppression ConstrainBoolean
	// Latency is "MediaTrackConstraintSet.latency"
	//
	// Optional
	Latency ConstrainDouble
	// ChannelCount is "MediaTrackConstraintSet.channelCount"
	//
	// Optional
	ChannelCount ConstrainULong
	// DeviceId is "MediaTrackConstraintSet.deviceId"
	//
	// Optional
	DeviceId ConstrainDOMString
	// GroupId is "MediaTrackConstraintSet.groupId"
	//
	// Optional
	GroupId ConstrainDOMString
	// WhiteBalanceMode is "MediaTrackConstraintSet.whiteBalanceMode"
	//
	// Optional
	WhiteBalanceMode ConstrainDOMString
	// ExposureMode is "MediaTrackConstraintSet.exposureMode"
	//
	// Optional
	ExposureMode ConstrainDOMString
	// FocusMode is "MediaTrackConstraintSet.focusMode"
	//
	// Optional
	FocusMode ConstrainDOMString
	// PointsOfInterest is "MediaTrackConstraintSet.pointsOfInterest"
	//
	// Optional
	PointsOfInterest ConstrainPoint2D
	// ExposureCompensation is "MediaTrackConstraintSet.exposureCompensation"
	//
	// Optional
	ExposureCompensation ConstrainDouble
	// ExposureTime is "MediaTrackConstraintSet.exposureTime"
	//
	// Optional
	ExposureTime ConstrainDouble
	// ColorTemperature is "MediaTrackConstraintSet.colorTemperature"
	//
	// Optional
	ColorTemperature ConstrainDouble
	// Iso is "MediaTrackConstraintSet.iso"
	//
	// Optional
	Iso ConstrainDouble
	// Brightness is "MediaTrackConstraintSet.brightness"
	//
	// Optional
	Brightness ConstrainDouble
	// Contrast is "MediaTrackConstraintSet.contrast"
	//
	// Optional
	Contrast ConstrainDouble
	// Saturation is "MediaTrackConstraintSet.saturation"
	//
	// Optional
	Saturation ConstrainDouble
	// Sharpness is "MediaTrackConstraintSet.sharpness"
	//
	// Optional
	Sharpness ConstrainDouble
	// FocusDistance is "MediaTrackConstraintSet.focusDistance"
	//
	// Optional
	FocusDistance ConstrainDouble
	// Pan is "MediaTrackConstraintSet.pan"
	//
	// Optional
	Pan OneOf_Bool_Float64_ConstrainDoubleRange
	// Tilt is "MediaTrackConstraintSet.tilt"
	//
	// Optional
	Tilt OneOf_Bool_Float64_ConstrainDoubleRange
	// Zoom is "MediaTrackConstraintSet.zoom"
	//
	// Optional
	Zoom OneOf_Bool_Float64_ConstrainDoubleRange
	// Torch is "MediaTrackConstraintSet.torch"
	//
	// Optional
	Torch ConstrainBoolean
	// DisplaySurface is "MediaTrackConstraintSet.displaySurface"
	//
	// Optional
	DisplaySurface ConstrainDOMString
	// LogicalSurface is "MediaTrackConstraintSet.logicalSurface"
	//
	// Optional
	LogicalSurface ConstrainBoolean
	// Cursor is "MediaTrackConstraintSet.cursor"
	//
	// Optional
	Cursor ConstrainDOMString
	// RestrictOwnAudio is "MediaTrackConstraintSet.restrictOwnAudio"
	//
	// Optional
	RestrictOwnAudio ConstrainBoolean
	// SuppressLocalAudioPlayback is "MediaTrackConstraintSet.suppressLocalAudioPlayback"
	//
	// Optional
	SuppressLocalAudioPlayback ConstrainBoolean

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaTrackConstraintSet with all fields set.
func (p MediaTrackConstraintSet) FromRef(ref js.Ref) MediaTrackConstraintSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaTrackConstraintSet in the application heap.
func (p MediaTrackConstraintSet) New() js.Ref {
	return bindings.MediaTrackConstraintSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaTrackConstraintSet) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackConstraintSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaTrackConstraintSet) Update(ref js.Ref) {
	bindings.MediaTrackConstraintSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaTrackConstraintSet) FreeMembers(recursive bool) {
	js.Free(
		p.Width.Ref(),
		p.Height.Ref(),
		p.AspectRatio.Ref(),
		p.FrameRate.Ref(),
		p.FacingMode.Ref(),
		p.ResizeMode.Ref(),
		p.SampleRate.Ref(),
		p.SampleSize.Ref(),
		p.EchoCancellation.Ref(),
		p.AutoGainControl.Ref(),
		p.NoiseSuppression.Ref(),
		p.Latency.Ref(),
		p.ChannelCount.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
		p.WhiteBalanceMode.Ref(),
		p.ExposureMode.Ref(),
		p.FocusMode.Ref(),
		p.PointsOfInterest.Ref(),
		p.ExposureCompensation.Ref(),
		p.ExposureTime.Ref(),
		p.ColorTemperature.Ref(),
		p.Iso.Ref(),
		p.Brightness.Ref(),
		p.Contrast.Ref(),
		p.Saturation.Ref(),
		p.Sharpness.Ref(),
		p.FocusDistance.Ref(),
		p.Pan.Ref(),
		p.Tilt.Ref(),
		p.Zoom.Ref(),
		p.Torch.Ref(),
		p.DisplaySurface.Ref(),
		p.LogicalSurface.Ref(),
		p.Cursor.Ref(),
		p.RestrictOwnAudio.Ref(),
		p.SuppressLocalAudioPlayback.Ref(),
	)
	p.Width = p.Width.FromRef(js.Undefined)
	p.Height = p.Height.FromRef(js.Undefined)
	p.AspectRatio = p.AspectRatio.FromRef(js.Undefined)
	p.FrameRate = p.FrameRate.FromRef(js.Undefined)
	p.FacingMode = p.FacingMode.FromRef(js.Undefined)
	p.ResizeMode = p.ResizeMode.FromRef(js.Undefined)
	p.SampleRate = p.SampleRate.FromRef(js.Undefined)
	p.SampleSize = p.SampleSize.FromRef(js.Undefined)
	p.EchoCancellation = p.EchoCancellation.FromRef(js.Undefined)
	p.AutoGainControl = p.AutoGainControl.FromRef(js.Undefined)
	p.NoiseSuppression = p.NoiseSuppression.FromRef(js.Undefined)
	p.Latency = p.Latency.FromRef(js.Undefined)
	p.ChannelCount = p.ChannelCount.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
	p.WhiteBalanceMode = p.WhiteBalanceMode.FromRef(js.Undefined)
	p.ExposureMode = p.ExposureMode.FromRef(js.Undefined)
	p.FocusMode = p.FocusMode.FromRef(js.Undefined)
	p.PointsOfInterest = p.PointsOfInterest.FromRef(js.Undefined)
	p.ExposureCompensation = p.ExposureCompensation.FromRef(js.Undefined)
	p.ExposureTime = p.ExposureTime.FromRef(js.Undefined)
	p.ColorTemperature = p.ColorTemperature.FromRef(js.Undefined)
	p.Iso = p.Iso.FromRef(js.Undefined)
	p.Brightness = p.Brightness.FromRef(js.Undefined)
	p.Contrast = p.Contrast.FromRef(js.Undefined)
	p.Saturation = p.Saturation.FromRef(js.Undefined)
	p.Sharpness = p.Sharpness.FromRef(js.Undefined)
	p.FocusDistance = p.FocusDistance.FromRef(js.Undefined)
	p.Pan = p.Pan.FromRef(js.Undefined)
	p.Tilt = p.Tilt.FromRef(js.Undefined)
	p.Zoom = p.Zoom.FromRef(js.Undefined)
	p.Torch = p.Torch.FromRef(js.Undefined)
	p.DisplaySurface = p.DisplaySurface.FromRef(js.Undefined)
	p.LogicalSurface = p.LogicalSurface.FromRef(js.Undefined)
	p.Cursor = p.Cursor.FromRef(js.Undefined)
	p.RestrictOwnAudio = p.RestrictOwnAudio.FromRef(js.Undefined)
	p.SuppressLocalAudioPlayback = p.SuppressLocalAudioPlayback.FromRef(js.Undefined)
}

type MediaTrackConstraints struct {
	// Advanced is "MediaTrackConstraints.advanced"
	//
	// Optional
	Advanced js.Array[MediaTrackConstraintSet]
	// WhiteBalanceMode is "MediaTrackConstraints.whiteBalanceMode"
	//
	// Optional
	WhiteBalanceMode ConstrainDOMString
	// ExposureMode is "MediaTrackConstraints.exposureMode"
	//
	// Optional
	ExposureMode ConstrainDOMString
	// FocusMode is "MediaTrackConstraints.focusMode"
	//
	// Optional
	FocusMode ConstrainDOMString
	// PointsOfInterest is "MediaTrackConstraints.pointsOfInterest"
	//
	// Optional
	PointsOfInterest ConstrainPoint2D
	// ExposureCompensation is "MediaTrackConstraints.exposureCompensation"
	//
	// Optional
	ExposureCompensation ConstrainDouble
	// ExposureTime is "MediaTrackConstraints.exposureTime"
	//
	// Optional
	ExposureTime ConstrainDouble
	// ColorTemperature is "MediaTrackConstraints.colorTemperature"
	//
	// Optional
	ColorTemperature ConstrainDouble
	// Iso is "MediaTrackConstraints.iso"
	//
	// Optional
	Iso ConstrainDouble
	// Brightness is "MediaTrackConstraints.brightness"
	//
	// Optional
	Brightness ConstrainDouble
	// Contrast is "MediaTrackConstraints.contrast"
	//
	// Optional
	Contrast ConstrainDouble
	// Saturation is "MediaTrackConstraints.saturation"
	//
	// Optional
	Saturation ConstrainDouble
	// Sharpness is "MediaTrackConstraints.sharpness"
	//
	// Optional
	Sharpness ConstrainDouble
	// FocusDistance is "MediaTrackConstraints.focusDistance"
	//
	// Optional
	FocusDistance ConstrainDouble
	// Pan is "MediaTrackConstraints.pan"
	//
	// Optional
	Pan OneOf_Bool_Float64_ConstrainDoubleRange
	// Tilt is "MediaTrackConstraints.tilt"
	//
	// Optional
	Tilt OneOf_Bool_Float64_ConstrainDoubleRange
	// Zoom is "MediaTrackConstraints.zoom"
	//
	// Optional
	Zoom OneOf_Bool_Float64_ConstrainDoubleRange
	// Torch is "MediaTrackConstraints.torch"
	//
	// Optional
	Torch ConstrainBoolean

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaTrackConstraints with all fields set.
func (p MediaTrackConstraints) FromRef(ref js.Ref) MediaTrackConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaTrackConstraints in the application heap.
func (p MediaTrackConstraints) New() js.Ref {
	return bindings.MediaTrackConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaTrackConstraints) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackConstraintsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaTrackConstraints) Update(ref js.Ref) {
	bindings.MediaTrackConstraintsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaTrackConstraints) FreeMembers(recursive bool) {
	js.Free(
		p.Advanced.Ref(),
		p.WhiteBalanceMode.Ref(),
		p.ExposureMode.Ref(),
		p.FocusMode.Ref(),
		p.PointsOfInterest.Ref(),
		p.ExposureCompensation.Ref(),
		p.ExposureTime.Ref(),
		p.ColorTemperature.Ref(),
		p.Iso.Ref(),
		p.Brightness.Ref(),
		p.Contrast.Ref(),
		p.Saturation.Ref(),
		p.Sharpness.Ref(),
		p.FocusDistance.Ref(),
		p.Pan.Ref(),
		p.Tilt.Ref(),
		p.Zoom.Ref(),
		p.Torch.Ref(),
	)
	p.Advanced = p.Advanced.FromRef(js.Undefined)
	p.WhiteBalanceMode = p.WhiteBalanceMode.FromRef(js.Undefined)
	p.ExposureMode = p.ExposureMode.FromRef(js.Undefined)
	p.FocusMode = p.FocusMode.FromRef(js.Undefined)
	p.PointsOfInterest = p.PointsOfInterest.FromRef(js.Undefined)
	p.ExposureCompensation = p.ExposureCompensation.FromRef(js.Undefined)
	p.ExposureTime = p.ExposureTime.FromRef(js.Undefined)
	p.ColorTemperature = p.ColorTemperature.FromRef(js.Undefined)
	p.Iso = p.Iso.FromRef(js.Undefined)
	p.Brightness = p.Brightness.FromRef(js.Undefined)
	p.Contrast = p.Contrast.FromRef(js.Undefined)
	p.Saturation = p.Saturation.FromRef(js.Undefined)
	p.Sharpness = p.Sharpness.FromRef(js.Undefined)
	p.FocusDistance = p.FocusDistance.FromRef(js.Undefined)
	p.Pan = p.Pan.FromRef(js.Undefined)
	p.Tilt = p.Tilt.FromRef(js.Undefined)
	p.Zoom = p.Zoom.FromRef(js.Undefined)
	p.Torch = p.Torch.FromRef(js.Undefined)
}

type MediaTrackSettings struct {
	// Width is "MediaTrackSettings.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width uint32
	// Height is "MediaTrackSettings.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height uint32
	// AspectRatio is "MediaTrackSettings.aspectRatio"
	//
	// Optional
	//
	// NOTE: FFI_USE_AspectRatio MUST be set to true to make this field effective.
	AspectRatio float64
	// FrameRate is "MediaTrackSettings.frameRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameRate MUST be set to true to make this field effective.
	FrameRate float64
	// FacingMode is "MediaTrackSettings.facingMode"
	//
	// Optional
	FacingMode js.String
	// ResizeMode is "MediaTrackSettings.resizeMode"
	//
	// Optional
	ResizeMode js.String
	// SampleRate is "MediaTrackSettings.sampleRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate uint32
	// SampleSize is "MediaTrackSettings.sampleSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_SampleSize MUST be set to true to make this field effective.
	SampleSize uint32
	// EchoCancellation is "MediaTrackSettings.echoCancellation"
	//
	// Optional
	//
	// NOTE: FFI_USE_EchoCancellation MUST be set to true to make this field effective.
	EchoCancellation bool
	// AutoGainControl is "MediaTrackSettings.autoGainControl"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoGainControl MUST be set to true to make this field effective.
	AutoGainControl bool
	// NoiseSuppression is "MediaTrackSettings.noiseSuppression"
	//
	// Optional
	//
	// NOTE: FFI_USE_NoiseSuppression MUST be set to true to make this field effective.
	NoiseSuppression bool
	// Latency is "MediaTrackSettings.latency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Latency MUST be set to true to make this field effective.
	Latency float64
	// ChannelCount is "MediaTrackSettings.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount uint32
	// DeviceId is "MediaTrackSettings.deviceId"
	//
	// Optional
	DeviceId js.String
	// GroupId is "MediaTrackSettings.groupId"
	//
	// Optional
	GroupId js.String
	// WhiteBalanceMode is "MediaTrackSettings.whiteBalanceMode"
	//
	// Optional
	WhiteBalanceMode js.String
	// ExposureMode is "MediaTrackSettings.exposureMode"
	//
	// Optional
	ExposureMode js.String
	// FocusMode is "MediaTrackSettings.focusMode"
	//
	// Optional
	FocusMode js.String
	// PointsOfInterest is "MediaTrackSettings.pointsOfInterest"
	//
	// Optional
	PointsOfInterest js.Array[Point2D]
	// ExposureCompensation is "MediaTrackSettings.exposureCompensation"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExposureCompensation MUST be set to true to make this field effective.
	ExposureCompensation float64
	// ExposureTime is "MediaTrackSettings.exposureTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ExposureTime MUST be set to true to make this field effective.
	ExposureTime float64
	// ColorTemperature is "MediaTrackSettings.colorTemperature"
	//
	// Optional
	//
	// NOTE: FFI_USE_ColorTemperature MUST be set to true to make this field effective.
	ColorTemperature float64
	// Iso is "MediaTrackSettings.iso"
	//
	// Optional
	//
	// NOTE: FFI_USE_Iso MUST be set to true to make this field effective.
	Iso float64
	// Brightness is "MediaTrackSettings.brightness"
	//
	// Optional
	//
	// NOTE: FFI_USE_Brightness MUST be set to true to make this field effective.
	Brightness float64
	// Contrast is "MediaTrackSettings.contrast"
	//
	// Optional
	//
	// NOTE: FFI_USE_Contrast MUST be set to true to make this field effective.
	Contrast float64
	// Saturation is "MediaTrackSettings.saturation"
	//
	// Optional
	//
	// NOTE: FFI_USE_Saturation MUST be set to true to make this field effective.
	Saturation float64
	// Sharpness is "MediaTrackSettings.sharpness"
	//
	// Optional
	//
	// NOTE: FFI_USE_Sharpness MUST be set to true to make this field effective.
	Sharpness float64
	// FocusDistance is "MediaTrackSettings.focusDistance"
	//
	// Optional
	//
	// NOTE: FFI_USE_FocusDistance MUST be set to true to make this field effective.
	FocusDistance float64
	// Pan is "MediaTrackSettings.pan"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pan MUST be set to true to make this field effective.
	Pan float64
	// Tilt is "MediaTrackSettings.tilt"
	//
	// Optional
	//
	// NOTE: FFI_USE_Tilt MUST be set to true to make this field effective.
	Tilt float64
	// Zoom is "MediaTrackSettings.zoom"
	//
	// Optional
	//
	// NOTE: FFI_USE_Zoom MUST be set to true to make this field effective.
	Zoom float64
	// Torch is "MediaTrackSettings.torch"
	//
	// Optional
	//
	// NOTE: FFI_USE_Torch MUST be set to true to make this field effective.
	Torch bool
	// DisplaySurface is "MediaTrackSettings.displaySurface"
	//
	// Optional
	DisplaySurface js.String
	// LogicalSurface is "MediaTrackSettings.logicalSurface"
	//
	// Optional
	//
	// NOTE: FFI_USE_LogicalSurface MUST be set to true to make this field effective.
	LogicalSurface bool
	// Cursor is "MediaTrackSettings.cursor"
	//
	// Optional
	Cursor js.String
	// RestrictOwnAudio is "MediaTrackSettings.restrictOwnAudio"
	//
	// Optional
	//
	// NOTE: FFI_USE_RestrictOwnAudio MUST be set to true to make this field effective.
	RestrictOwnAudio bool
	// SuppressLocalAudioPlayback is "MediaTrackSettings.suppressLocalAudioPlayback"
	//
	// Optional
	//
	// NOTE: FFI_USE_SuppressLocalAudioPlayback MUST be set to true to make this field effective.
	SuppressLocalAudioPlayback bool

	FFI_USE_Width                      bool // for Width.
	FFI_USE_Height                     bool // for Height.
	FFI_USE_AspectRatio                bool // for AspectRatio.
	FFI_USE_FrameRate                  bool // for FrameRate.
	FFI_USE_SampleRate                 bool // for SampleRate.
	FFI_USE_SampleSize                 bool // for SampleSize.
	FFI_USE_EchoCancellation           bool // for EchoCancellation.
	FFI_USE_AutoGainControl            bool // for AutoGainControl.
	FFI_USE_NoiseSuppression           bool // for NoiseSuppression.
	FFI_USE_Latency                    bool // for Latency.
	FFI_USE_ChannelCount               bool // for ChannelCount.
	FFI_USE_ExposureCompensation       bool // for ExposureCompensation.
	FFI_USE_ExposureTime               bool // for ExposureTime.
	FFI_USE_ColorTemperature           bool // for ColorTemperature.
	FFI_USE_Iso                        bool // for Iso.
	FFI_USE_Brightness                 bool // for Brightness.
	FFI_USE_Contrast                   bool // for Contrast.
	FFI_USE_Saturation                 bool // for Saturation.
	FFI_USE_Sharpness                  bool // for Sharpness.
	FFI_USE_FocusDistance              bool // for FocusDistance.
	FFI_USE_Pan                        bool // for Pan.
	FFI_USE_Tilt                       bool // for Tilt.
	FFI_USE_Zoom                       bool // for Zoom.
	FFI_USE_Torch                      bool // for Torch.
	FFI_USE_LogicalSurface             bool // for LogicalSurface.
	FFI_USE_RestrictOwnAudio           bool // for RestrictOwnAudio.
	FFI_USE_SuppressLocalAudioPlayback bool // for SuppressLocalAudioPlayback.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaTrackSettings with all fields set.
func (p MediaTrackSettings) FromRef(ref js.Ref) MediaTrackSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaTrackSettings in the application heap.
func (p MediaTrackSettings) New() js.Ref {
	return bindings.MediaTrackSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaTrackSettings) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackSettingsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaTrackSettings) Update(ref js.Ref) {
	bindings.MediaTrackSettingsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaTrackSettings) FreeMembers(recursive bool) {
	js.Free(
		p.FacingMode.Ref(),
		p.ResizeMode.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
		p.WhiteBalanceMode.Ref(),
		p.ExposureMode.Ref(),
		p.FocusMode.Ref(),
		p.PointsOfInterest.Ref(),
		p.DisplaySurface.Ref(),
		p.Cursor.Ref(),
	)
	p.FacingMode = p.FacingMode.FromRef(js.Undefined)
	p.ResizeMode = p.ResizeMode.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
	p.WhiteBalanceMode = p.WhiteBalanceMode.FromRef(js.Undefined)
	p.ExposureMode = p.ExposureMode.FromRef(js.Undefined)
	p.FocusMode = p.FocusMode.FromRef(js.Undefined)
	p.PointsOfInterest = p.PointsOfInterest.FromRef(js.Undefined)
	p.DisplaySurface = p.DisplaySurface.FromRef(js.Undefined)
	p.Cursor = p.Cursor.FromRef(js.Undefined)
}

type CaptureAction uint32

const (
	_ CaptureAction = iota

	CaptureAction_NEXT
	CaptureAction_PREVIOUS
	CaptureAction_FIRST
	CaptureAction_LAST
)

func (CaptureAction) FromRef(str js.Ref) CaptureAction {
	return CaptureAction(bindings.ConstOfCaptureAction(str))
}

func (x CaptureAction) String() (string, bool) {
	switch x {
	case CaptureAction_NEXT:
		return "next", true
	case CaptureAction_PREVIOUS:
		return "previous", true
	case CaptureAction_FIRST:
		return "first", true
	case CaptureAction_LAST:
		return "last", true
	default:
		return "", false
	}
}

type CaptureHandle struct {
	// Origin is "CaptureHandle.origin"
	//
	// Optional
	Origin js.String
	// Handle is "CaptureHandle.handle"
	//
	// Optional
	Handle js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CaptureHandle with all fields set.
func (p CaptureHandle) FromRef(ref js.Ref) CaptureHandle {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CaptureHandle in the application heap.
func (p CaptureHandle) New() js.Ref {
	return bindings.CaptureHandleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CaptureHandle) UpdateFrom(ref js.Ref) {
	bindings.CaptureHandleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CaptureHandle) Update(ref js.Ref) {
	bindings.CaptureHandleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CaptureHandle) FreeMembers(recursive bool) {
	js.Free(
		p.Origin.Ref(),
		p.Handle.Ref(),
	)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.Handle = p.Handle.FromRef(js.Undefined)
}

type MediaStreamTrackState uint32

const (
	_ MediaStreamTrackState = iota

	MediaStreamTrackState_LIVE
	MediaStreamTrackState_ENDED
)

func (MediaStreamTrackState) FromRef(str js.Ref) MediaStreamTrackState {
	return MediaStreamTrackState(bindings.ConstOfMediaStreamTrackState(str))
}

func (x MediaStreamTrackState) String() (string, bool) {
	switch x {
	case MediaStreamTrackState_LIVE:
		return "live", true
	case MediaStreamTrackState_ENDED:
		return "ended", true
	default:
		return "", false
	}
}
