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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "TextTrackCueList.length".
//
// The returned bool will be false if there is no such property.
func (this TextTrackCueList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCueListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "TextTrackCueList.".
//
// The returned bool will be false if there is no such method.
func (this TextTrackCueList) Get(index uint32) (TextTrackCue, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackCueListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return TextTrackCue{}.FromRef(_ret), _ok
}

// GetFunc returns the method "TextTrackCueList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrackCueList) GetFunc() (fn js.Func[func(index uint32) TextTrackCue]) {
	return fn.FromRef(
		bindings.TextTrackCueListGetFunc(
			this.Ref(),
		),
	)
}

// GetCueById calls the method "TextTrackCueList.getCueById".
//
// The returned bool will be false if there is no such method.
func (this TextTrackCueList) GetCueById(id js.String) (TextTrackCue, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackCueListGetCueById(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return TextTrackCue{}.FromRef(_ret), _ok
}

// GetCueByIdFunc returns the method "TextTrackCueList.getCueById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrackCueList) GetCueByIdFunc() (fn js.Func[func(id js.String) TextTrackCue]) {
	return fn.FromRef(
		bindings.TextTrackCueListGetCueByIdFunc(
			this.Ref(),
		),
	)
}

type TimeRanges struct {
	ref js.Ref
}

func (this TimeRanges) Once() TimeRanges {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "TimeRanges.length".
//
// The returned bool will be false if there is no such property.
func (this TimeRanges) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTimeRangesLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Start calls the method "TimeRanges.start".
//
// The returned bool will be false if there is no such method.
func (this TimeRanges) Start(index uint32) (float64, bool) {
	var _ok bool
	_ret := bindings.CallTimeRangesStart(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return float64(_ret), _ok
}

// StartFunc returns the method "TimeRanges.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TimeRanges) StartFunc() (fn js.Func[func(index uint32) float64]) {
	return fn.FromRef(
		bindings.TimeRangesStartFunc(
			this.Ref(),
		),
	)
}

// End calls the method "TimeRanges.end".
//
// The returned bool will be false if there is no such method.
func (this TimeRanges) End(index uint32) (float64, bool) {
	var _ok bool
	_ret := bindings.CallTimeRangesEnd(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return float64(_ret), _ok
}

// EndFunc returns the method "TimeRanges.end".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TimeRanges) EndFunc() (fn js.Func[func(index uint32) float64]) {
	return fn.FromRef(
		bindings.TimeRangesEndFunc(
			this.Ref(),
		),
	)
}

type AudioTrack struct {
	ref js.Ref
}

func (this AudioTrack) Once() AudioTrack {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Id returns the value of property "AudioTrack.id".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Kind returns the value of property "AudioTrack.kind".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) Kind() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label returns the value of property "AudioTrack.label".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "AudioTrack.language".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Enabled returns the value of property "AudioTrack.enabled".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) Enabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Enabled sets the value of property "AudioTrack.enabled" to val.
//
// It returns false if the property cannot be set.
func (this AudioTrack) SetEnabled(val bool) bool {
	return js.True == bindings.SetAudioTrackEnabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// SourceBuffer returns the value of property "AudioTrack.sourceBuffer".
//
// The returned bool will be false if there is no such property.
func (this AudioTrack) SourceBuffer() (SourceBuffer, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackSourceBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return SourceBuffer{}.FromRef(_ret), _ok
}

type AudioTrackList struct {
	EventTarget
}

func (this AudioTrackList) Once() AudioTrackList {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "AudioTrackList.length".
//
// The returned bool will be false if there is no such property.
func (this AudioTrackList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioTrackListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "AudioTrackList.".
//
// The returned bool will be false if there is no such method.
func (this AudioTrackList) Get(index uint32) (AudioTrack, bool) {
	var _ok bool
	_ret := bindings.CallAudioTrackListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return AudioTrack{}.FromRef(_ret), _ok
}

// GetFunc returns the method "AudioTrackList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioTrackList) GetFunc() (fn js.Func[func(index uint32) AudioTrack]) {
	return fn.FromRef(
		bindings.AudioTrackListGetFunc(
			this.Ref(),
		),
	)
}

// GetTrackById calls the method "AudioTrackList.getTrackById".
//
// The returned bool will be false if there is no such method.
func (this AudioTrackList) GetTrackById(id js.String) (AudioTrack, bool) {
	var _ok bool
	_ret := bindings.CallAudioTrackListGetTrackById(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return AudioTrack{}.FromRef(_ret), _ok
}

// GetTrackByIdFunc returns the method "AudioTrackList.getTrackById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioTrackList) GetTrackByIdFunc() (fn js.Func[func(id js.String) AudioTrack]) {
	return fn.FromRef(
		bindings.AudioTrackListGetTrackByIdFunc(
			this.Ref(),
		),
	)
}

type VideoTrack struct {
	ref js.Ref
}

func (this VideoTrack) Once() VideoTrack {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Id returns the value of property "VideoTrack.id".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Kind returns the value of property "VideoTrack.kind".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) Kind() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label returns the value of property "VideoTrack.label".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "VideoTrack.language".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Selected returns the value of property "VideoTrack.selected".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) Selected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Selected sets the value of property "VideoTrack.selected" to val.
//
// It returns false if the property cannot be set.
func (this VideoTrack) SetSelected(val bool) bool {
	return js.True == bindings.SetVideoTrackSelected(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// SourceBuffer returns the value of property "VideoTrack.sourceBuffer".
//
// The returned bool will be false if there is no such property.
func (this VideoTrack) SourceBuffer() (SourceBuffer, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackSourceBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return SourceBuffer{}.FromRef(_ret), _ok
}

type VideoTrackList struct {
	EventTarget
}

func (this VideoTrackList) Once() VideoTrackList {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "VideoTrackList.length".
//
// The returned bool will be false if there is no such property.
func (this VideoTrackList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectedIndex returns the value of property "VideoTrackList.selectedIndex".
//
// The returned bool will be false if there is no such property.
func (this VideoTrackList) SelectedIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetVideoTrackListSelectedIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Get calls the method "VideoTrackList.".
//
// The returned bool will be false if there is no such method.
func (this VideoTrackList) Get(index uint32) (VideoTrack, bool) {
	var _ok bool
	_ret := bindings.CallVideoTrackListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return VideoTrack{}.FromRef(_ret), _ok
}

// GetFunc returns the method "VideoTrackList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoTrackList) GetFunc() (fn js.Func[func(index uint32) VideoTrack]) {
	return fn.FromRef(
		bindings.VideoTrackListGetFunc(
			this.Ref(),
		),
	)
}

// GetTrackById calls the method "VideoTrackList.getTrackById".
//
// The returned bool will be false if there is no such method.
func (this VideoTrackList) GetTrackById(id js.String) (VideoTrack, bool) {
	var _ok bool
	_ret := bindings.CallVideoTrackListGetTrackById(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return VideoTrack{}.FromRef(_ret), _ok
}

// GetTrackByIdFunc returns the method "VideoTrackList.getTrackById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoTrackList) GetTrackByIdFunc() (fn js.Func[func(id js.String) VideoTrack]) {
	return fn.FromRef(
		bindings.VideoTrackListGetTrackByIdFunc(
			this.Ref(),
		),
	)
}

type TextTrackList struct {
	EventTarget
}

func (this TextTrackList) Once() TextTrackList {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "TextTrackList.length".
//
// The returned bool will be false if there is no such property.
func (this TextTrackList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "TextTrackList.".
//
// The returned bool will be false if there is no such method.
func (this TextTrackList) Get(index uint32) (TextTrack, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return TextTrack{}.FromRef(_ret), _ok
}

// GetFunc returns the method "TextTrackList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrackList) GetFunc() (fn js.Func[func(index uint32) TextTrack]) {
	return fn.FromRef(
		bindings.TextTrackListGetFunc(
			this.Ref(),
		),
	)
}

// GetTrackById calls the method "TextTrackList.getTrackById".
//
// The returned bool will be false if there is no such method.
func (this TextTrackList) GetTrackById(id js.String) (TextTrack, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackListGetTrackById(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return TextTrack{}.FromRef(_ret), _ok
}

// GetTrackByIdFunc returns the method "TextTrackList.getTrackById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrackList) GetTrackByIdFunc() (fn js.Func[func(id js.String) TextTrack]) {
	return fn.FromRef(
		bindings.TextTrackListGetTrackByIdFunc(
			this.Ref(),
		),
	)
}

type SourceBuffer struct {
	EventTarget
}

func (this SourceBuffer) Once() SourceBuffer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Mode returns the value of property "SourceBuffer.mode".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) Mode() (AppendMode, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return AppendMode(_ret), _ok
}

// Mode sets the value of property "SourceBuffer.mode" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetMode(val AppendMode) bool {
	return js.True == bindings.SetSourceBufferMode(
		this.Ref(),
		uint32(val),
	)
}

// Updating returns the value of property "SourceBuffer.updating".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) Updating() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferUpdating(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Buffered returns the value of property "SourceBuffer.buffered".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) Buffered() (TimeRanges, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferBuffered(
		this.Ref(), js.Pointer(&_ok),
	)
	return TimeRanges{}.FromRef(_ret), _ok
}

// TimestampOffset returns the value of property "SourceBuffer.timestampOffset".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) TimestampOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferTimestampOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// TimestampOffset sets the value of property "SourceBuffer.timestampOffset" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetTimestampOffset(val float64) bool {
	return js.True == bindings.SetSourceBufferTimestampOffset(
		this.Ref(),
		float64(val),
	)
}

// AudioTracks returns the value of property "SourceBuffer.audioTracks".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) AudioTracks() (AudioTrackList, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferAudioTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioTrackList{}.FromRef(_ret), _ok
}

// VideoTracks returns the value of property "SourceBuffer.videoTracks".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) VideoTracks() (VideoTrackList, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferVideoTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoTrackList{}.FromRef(_ret), _ok
}

// TextTracks returns the value of property "SourceBuffer.textTracks".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) TextTracks() (TextTrackList, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferTextTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackList{}.FromRef(_ret), _ok
}

// AppendWindowStart returns the value of property "SourceBuffer.appendWindowStart".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) AppendWindowStart() (float64, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferAppendWindowStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AppendWindowStart sets the value of property "SourceBuffer.appendWindowStart" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetAppendWindowStart(val float64) bool {
	return js.True == bindings.SetSourceBufferAppendWindowStart(
		this.Ref(),
		float64(val),
	)
}

// AppendWindowEnd returns the value of property "SourceBuffer.appendWindowEnd".
//
// The returned bool will be false if there is no such property.
func (this SourceBuffer) AppendWindowEnd() (float64, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferAppendWindowEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AppendWindowEnd sets the value of property "SourceBuffer.appendWindowEnd" to val.
//
// It returns false if the property cannot be set.
func (this SourceBuffer) SetAppendWindowEnd(val float64) bool {
	return js.True == bindings.SetSourceBufferAppendWindowEnd(
		this.Ref(),
		float64(val),
	)
}

// AppendBuffer calls the method "SourceBuffer.appendBuffer".
//
// The returned bool will be false if there is no such method.
func (this SourceBuffer) AppendBuffer(data BufferSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSourceBufferAppendBuffer(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AppendBufferFunc returns the method "SourceBuffer.appendBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SourceBuffer) AppendBufferFunc() (fn js.Func[func(data BufferSource)]) {
	return fn.FromRef(
		bindings.SourceBufferAppendBufferFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "SourceBuffer.abort".
//
// The returned bool will be false if there is no such method.
func (this SourceBuffer) Abort() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSourceBufferAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "SourceBuffer.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SourceBuffer) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SourceBufferAbortFunc(
			this.Ref(),
		),
	)
}

// ChangeType calls the method "SourceBuffer.changeType".
//
// The returned bool will be false if there is no such method.
func (this SourceBuffer) ChangeType(typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSourceBufferChangeType(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ChangeTypeFunc returns the method "SourceBuffer.changeType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SourceBuffer) ChangeTypeFunc() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.SourceBufferChangeTypeFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "SourceBuffer.remove".
//
// The returned bool will be false if there is no such method.
func (this SourceBuffer) Remove(start float64, end float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSourceBufferRemove(
		this.Ref(), js.Pointer(&_ok),
		float64(start),
		float64(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "SourceBuffer.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SourceBuffer) RemoveFunc() (fn js.Func[func(start float64, end float64)]) {
	return fn.FromRef(
		bindings.SourceBufferRemoveFunc(
			this.Ref(),
		),
	)
}

type TextTrack struct {
	EventTarget
}

func (this TextTrack) Once() TextTrack {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Kind returns the value of property "TextTrack.kind".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Kind() (TextTrackKind, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackKind(_ret), _ok
}

// Label returns the value of property "TextTrack.label".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "TextTrack.language".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "TextTrack.id".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InBandMetadataTrackDispatchType returns the value of property "TextTrack.inBandMetadataTrackDispatchType".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) InBandMetadataTrackDispatchType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackInBandMetadataTrackDispatchType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Mode returns the value of property "TextTrack.mode".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Mode() (TextTrackMode, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackMode(_ret), _ok
}

// Mode sets the value of property "TextTrack.mode" to val.
//
// It returns false if the property cannot be set.
func (this TextTrack) SetMode(val TextTrackMode) bool {
	return js.True == bindings.SetTextTrackMode(
		this.Ref(),
		uint32(val),
	)
}

// Cues returns the value of property "TextTrack.cues".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) Cues() (TextTrackCueList, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCues(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackCueList{}.FromRef(_ret), _ok
}

// ActiveCues returns the value of property "TextTrack.activeCues".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) ActiveCues() (TextTrackCueList, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackActiveCues(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackCueList{}.FromRef(_ret), _ok
}

// SourceBuffer returns the value of property "TextTrack.sourceBuffer".
//
// The returned bool will be false if there is no such property.
func (this TextTrack) SourceBuffer() (SourceBuffer, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackSourceBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return SourceBuffer{}.FromRef(_ret), _ok
}

// AddCue calls the method "TextTrack.addCue".
//
// The returned bool will be false if there is no such method.
func (this TextTrack) AddCue(cue TextTrackCue) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackAddCue(
		this.Ref(), js.Pointer(&_ok),
		cue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddCueFunc returns the method "TextTrack.addCue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrack) AddCueFunc() (fn js.Func[func(cue TextTrackCue)]) {
	return fn.FromRef(
		bindings.TextTrackAddCueFunc(
			this.Ref(),
		),
	)
}

// RemoveCue calls the method "TextTrack.removeCue".
//
// The returned bool will be false if there is no such method.
func (this TextTrack) RemoveCue(cue TextTrackCue) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallTextTrackRemoveCue(
		this.Ref(), js.Pointer(&_ok),
		cue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveCueFunc returns the method "TextTrack.removeCue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TextTrack) RemoveCueFunc() (fn js.Func[func(cue TextTrackCue)]) {
	return fn.FromRef(
		bindings.TextTrackRemoveCueFunc(
			this.Ref(),
		),
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Size returns the value of property "MediaKeyStatusMap.size".
//
// The returned bool will be false if there is no such property.
func (this MediaKeyStatusMap) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeyStatusMapSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Has calls the method "MediaKeyStatusMap.has".
//
// The returned bool will be false if there is no such method.
func (this MediaKeyStatusMap) Has(keyId BufferSource) (bool, bool) {
	var _ok bool
	_ret := bindings.CallMediaKeyStatusMapHas(
		this.Ref(), js.Pointer(&_ok),
		keyId.Ref(),
	)

	return _ret == js.True, _ok
}

// HasFunc returns the method "MediaKeyStatusMap.has".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeyStatusMap) HasFunc() (fn js.Func[func(keyId BufferSource) bool]) {
	return fn.FromRef(
		bindings.MediaKeyStatusMapHasFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "MediaKeyStatusMap.get".
//
// The returned bool will be false if there is no such method.
func (this MediaKeyStatusMap) Get(keyId BufferSource) (OneOf_MediaKeyStatus_undefined, bool) {
	var _ok bool
	_ret := bindings.CallMediaKeyStatusMapGet(
		this.Ref(), js.Pointer(&_ok),
		keyId.Ref(),
	)

	return OneOf_MediaKeyStatus_undefined{}.FromRef(_ret), _ok
}

// GetFunc returns the method "MediaKeyStatusMap.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeyStatusMap) GetFunc() (fn js.Func[func(keyId BufferSource) OneOf_MediaKeyStatus_undefined]) {
	return fn.FromRef(
		bindings.MediaKeyStatusMapGetFunc(
			this.Ref(),
		),
	)
}

type MediaKeySession struct {
	EventTarget
}

func (this MediaKeySession) Once() MediaKeySession {
	this.Ref().Once()
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
	this.Ref().Free()
}

// SessionId returns the value of property "MediaKeySession.sessionId".
//
// The returned bool will be false if there is no such property.
func (this MediaKeySession) SessionId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeySessionSessionId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Expiration returns the value of property "MediaKeySession.expiration".
//
// The returned bool will be false if there is no such property.
func (this MediaKeySession) Expiration() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeySessionExpiration(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Closed returns the value of property "MediaKeySession.closed".
//
// The returned bool will be false if there is no such property.
func (this MediaKeySession) Closed() (js.Promise[MediaKeySessionClosedReason], bool) {
	var _ok bool
	_ret := bindings.GetMediaKeySessionClosed(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[MediaKeySessionClosedReason]{}.FromRef(_ret), _ok
}

// KeyStatuses returns the value of property "MediaKeySession.keyStatuses".
//
// The returned bool will be false if there is no such property.
func (this MediaKeySession) KeyStatuses() (MediaKeyStatusMap, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeySessionKeyStatuses(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaKeyStatusMap{}.FromRef(_ret), _ok
}

// GenerateRequest calls the method "MediaKeySession.generateRequest".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySession) GenerateRequest(initDataType js.String, initData BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySessionGenerateRequest(
		this.Ref(), js.Pointer(&_ok),
		initDataType.Ref(),
		initData.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// GenerateRequestFunc returns the method "MediaKeySession.generateRequest".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySession) GenerateRequestFunc() (fn js.Func[func(initDataType js.String, initData BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaKeySessionGenerateRequestFunc(
			this.Ref(),
		),
	)
}

// Load calls the method "MediaKeySession.load".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySession) Load(sessionId js.String) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySessionLoad(
		this.Ref(), js.Pointer(&_ok),
		sessionId.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// LoadFunc returns the method "MediaKeySession.load".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySession) LoadFunc() (fn js.Func[func(sessionId js.String) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.MediaKeySessionLoadFunc(
			this.Ref(),
		),
	)
}

// Update calls the method "MediaKeySession.update".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySession) Update(response BufferSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySessionUpdate(
		this.Ref(), js.Pointer(&_ok),
		response.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UpdateFunc returns the method "MediaKeySession.update".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySession) UpdateFunc() (fn js.Func[func(response BufferSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaKeySessionUpdateFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "MediaKeySession.close".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySession) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySessionClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "MediaKeySession.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySession) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaKeySessionCloseFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "MediaKeySession.remove".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySession) Remove() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySessionRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RemoveFunc returns the method "MediaKeySession.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySession) RemoveFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaKeySessionRemoveFunc(
			this.Ref(),
		),
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// CreateSession calls the method "MediaKeys.createSession".
//
// The returned bool will be false if there is no such method.
func (this MediaKeys) CreateSession(sessionType MediaKeySessionType) (MediaKeySession, bool) {
	var _ok bool
	_ret := bindings.CallMediaKeysCreateSession(
		this.Ref(), js.Pointer(&_ok),
		uint32(sessionType),
	)

	return MediaKeySession{}.FromRef(_ret), _ok
}

// CreateSessionFunc returns the method "MediaKeys.createSession".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeys) CreateSessionFunc() (fn js.Func[func(sessionType MediaKeySessionType) MediaKeySession]) {
	return fn.FromRef(
		bindings.MediaKeysCreateSessionFunc(
			this.Ref(),
		),
	)
}

// CreateSession1 calls the method "MediaKeys.createSession".
//
// The returned bool will be false if there is no such method.
func (this MediaKeys) CreateSession1() (MediaKeySession, bool) {
	var _ok bool
	_ret := bindings.CallMediaKeysCreateSession1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaKeySession{}.FromRef(_ret), _ok
}

// CreateSession1Func returns the method "MediaKeys.createSession".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeys) CreateSession1Func() (fn js.Func[func() MediaKeySession]) {
	return fn.FromRef(
		bindings.MediaKeysCreateSession1Func(
			this.Ref(),
		),
	)
}

// SetServerCertificate calls the method "MediaKeys.setServerCertificate".
//
// The returned bool will be false if there is no such method.
func (this MediaKeys) SetServerCertificate(serverCertificate BufferSource) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeysSetServerCertificate(
		this.Ref(), js.Pointer(&_ok),
		serverCertificate.Ref(),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// SetServerCertificateFunc returns the method "MediaKeys.setServerCertificate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeys) SetServerCertificateFunc() (fn js.Func[func(serverCertificate BufferSource) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.MediaKeysSetServerCertificateFunc(
			this.Ref(),
		),
	)
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
func (p ULongRange) UpdateFrom(ref js.Ref) {
	bindings.ULongRangeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ULongRange) Update(ref js.Ref) {
	bindings.ULongRangeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p DoubleRange) UpdateFrom(ref js.Ref) {
	bindings.DoubleRangeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DoubleRange) Update(ref js.Ref) {
	bindings.DoubleRangeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaSettingsRange) UpdateFrom(ref js.Ref) {
	bindings.MediaSettingsRangeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaSettingsRange) Update(ref js.Ref) {
	bindings.MediaSettingsRangeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaTrackCapabilities struct {
	// Width is "MediaTrackCapabilities.width"
	//
	// Optional
	Width ULongRange
	// Height is "MediaTrackCapabilities.height"
	//
	// Optional
	Height ULongRange
	// AspectRatio is "MediaTrackCapabilities.aspectRatio"
	//
	// Optional
	AspectRatio DoubleRange
	// FrameRate is "MediaTrackCapabilities.frameRate"
	//
	// Optional
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
	SampleRate ULongRange
	// SampleSize is "MediaTrackCapabilities.sampleSize"
	//
	// Optional
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
	Latency DoubleRange
	// ChannelCount is "MediaTrackCapabilities.channelCount"
	//
	// Optional
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
	ExposureCompensation MediaSettingsRange
	// ExposureTime is "MediaTrackCapabilities.exposureTime"
	//
	// Optional
	ExposureTime MediaSettingsRange
	// ColorTemperature is "MediaTrackCapabilities.colorTemperature"
	//
	// Optional
	ColorTemperature MediaSettingsRange
	// Iso is "MediaTrackCapabilities.iso"
	//
	// Optional
	Iso MediaSettingsRange
	// Brightness is "MediaTrackCapabilities.brightness"
	//
	// Optional
	Brightness MediaSettingsRange
	// Contrast is "MediaTrackCapabilities.contrast"
	//
	// Optional
	Contrast MediaSettingsRange
	// Saturation is "MediaTrackCapabilities.saturation"
	//
	// Optional
	Saturation MediaSettingsRange
	// Sharpness is "MediaTrackCapabilities.sharpness"
	//
	// Optional
	Sharpness MediaSettingsRange
	// FocusDistance is "MediaTrackCapabilities.focusDistance"
	//
	// Optional
	FocusDistance MediaSettingsRange
	// Pan is "MediaTrackCapabilities.pan"
	//
	// Optional
	Pan MediaSettingsRange
	// Tilt is "MediaTrackCapabilities.tilt"
	//
	// Optional
	Tilt MediaSettingsRange
	// Zoom is "MediaTrackCapabilities.zoom"
	//
	// Optional
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
func (p MediaTrackCapabilities) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackCapabilitiesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaTrackCapabilities) Update(ref js.Ref) {
	bindings.MediaTrackCapabilitiesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p ConstrainULongRange) UpdateFrom(ref js.Ref) {
	bindings.ConstrainULongRangeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstrainULongRange) Update(ref js.Ref) {
	bindings.ConstrainULongRangeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p ConstrainDoubleRange) UpdateFrom(ref js.Ref) {
	bindings.ConstrainDoubleRangeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstrainDoubleRange) Update(ref js.Ref) {
	bindings.ConstrainDoubleRangeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p ConstrainDOMStringParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainDOMStringParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstrainDOMStringParameters) Update(ref js.Ref) {
	bindings.ConstrainDOMStringParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p ConstrainBooleanParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainBooleanParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstrainBooleanParameters) Update(ref js.Ref) {
	bindings.ConstrainBooleanParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p Point2D) UpdateFrom(ref js.Ref) {
	bindings.Point2DJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p Point2D) Update(ref js.Ref) {
	bindings.Point2DJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p ConstrainPoint2DParameters) UpdateFrom(ref js.Ref) {
	bindings.ConstrainPoint2DParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ConstrainPoint2DParameters) Update(ref js.Ref) {
	bindings.ConstrainPoint2DParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaTrackConstraintSet) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackConstraintSetJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaTrackConstraintSet) Update(ref js.Ref) {
	bindings.MediaTrackConstraintSetJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaTrackConstraints) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackConstraintsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaTrackConstraints) Update(ref js.Ref) {
	bindings.MediaTrackConstraintsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaTrackSettings) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaTrackSettings) Update(ref js.Ref) {
	bindings.MediaTrackSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p CaptureHandle) UpdateFrom(ref js.Ref) {
	bindings.CaptureHandleJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CaptureHandle) Update(ref js.Ref) {
	bindings.CaptureHandleJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
