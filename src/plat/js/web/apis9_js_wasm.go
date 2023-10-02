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

type MediaStreamTrack struct {
	EventTarget
}

func (this MediaStreamTrack) Once() MediaStreamTrack {
	this.Ref().Once()
	return this
}

func (this MediaStreamTrack) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaStreamTrack) FromRef(ref js.Ref) MediaStreamTrack {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaStreamTrack) Free() {
	this.Ref().Free()
}

// Kind returns the value of property "MediaStreamTrack.kind".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Kind() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "MediaStreamTrack.id".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label returns the value of property "MediaStreamTrack.label".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Enabled returns the value of property "MediaStreamTrack.enabled".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Enabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Enabled sets the value of property "MediaStreamTrack.enabled" to val.
//
// It returns false if the property cannot be set.
func (this MediaStreamTrack) SetEnabled(val bool) bool {
	return js.True == bindings.SetMediaStreamTrackEnabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Muted returns the value of property "MediaStreamTrack.muted".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Muted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackMuted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReadyState returns the value of property "MediaStreamTrack.readyState".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) ReadyState() (MediaStreamTrackState, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrackState(_ret), _ok
}

// ContentHint returns the value of property "MediaStreamTrack.contentHint".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) ContentHint() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackContentHint(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContentHint sets the value of property "MediaStreamTrack.contentHint" to val.
//
// It returns false if the property cannot be set.
func (this MediaStreamTrack) SetContentHint(val js.String) bool {
	return js.True == bindings.SetMediaStreamTrackContentHint(
		this.Ref(),
		val.Ref(),
	)
}

// Isolated returns the value of property "MediaStreamTrack.isolated".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrack) Isolated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackIsolated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Clone calls the method "MediaStreamTrack.clone".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) Clone() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "MediaStreamTrack.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) CloneFunc() (fn js.Func[func() MediaStreamTrack]) {
	return fn.FromRef(
		bindings.MediaStreamTrackCloneFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "MediaStreamTrack.stop".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "MediaStreamTrack.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaStreamTrackStopFunc(
			this.Ref(),
		),
	)
}

// GetCapabilities calls the method "MediaStreamTrack.getCapabilities".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) GetCapabilities() (MediaTrackCapabilities, bool) {
	var _ret MediaTrackCapabilities
	_ok := js.True == bindings.CallMediaStreamTrackGetCapabilities(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetCapabilitiesFunc returns the method "MediaStreamTrack.getCapabilities".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) GetCapabilitiesFunc() (fn js.Func[func() MediaTrackCapabilities]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetConstraints calls the method "MediaStreamTrack.getConstraints".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) GetConstraints() (MediaTrackConstraints, bool) {
	var _ret MediaTrackConstraints
	_ok := js.True == bindings.CallMediaStreamTrackGetConstraints(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetConstraintsFunc returns the method "MediaStreamTrack.getConstraints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) GetConstraintsFunc() (fn js.Func[func() MediaTrackConstraints]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetConstraintsFunc(
			this.Ref(),
		),
	)
}

// GetSettings calls the method "MediaStreamTrack.getSettings".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) GetSettings() (MediaTrackSettings, bool) {
	var _ret MediaTrackSettings
	_ok := js.True == bindings.CallMediaStreamTrackGetSettings(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetSettingsFunc returns the method "MediaStreamTrack.getSettings".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) GetSettingsFunc() (fn js.Func[func() MediaTrackSettings]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetSettingsFunc(
			this.Ref(),
		),
	)
}

// ApplyConstraints calls the method "MediaStreamTrack.applyConstraints".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) ApplyConstraints(constraints MediaTrackConstraints) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackApplyConstraints(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&constraints),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ApplyConstraintsFunc returns the method "MediaStreamTrack.applyConstraints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) ApplyConstraintsFunc() (fn js.Func[func(constraints MediaTrackConstraints) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackApplyConstraintsFunc(
			this.Ref(),
		),
	)
}

// ApplyConstraints1 calls the method "MediaStreamTrack.applyConstraints".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) ApplyConstraints1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackApplyConstraints1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ApplyConstraints1Func returns the method "MediaStreamTrack.applyConstraints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) ApplyConstraints1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackApplyConstraints1Func(
			this.Ref(),
		),
	)
}

// GetSupportedCaptureActions calls the method "MediaStreamTrack.getSupportedCaptureActions".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) GetSupportedCaptureActions() (js.Array[js.String], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackGetSupportedCaptureActions(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.String]{}.FromRef(_ret), _ok
}

// GetSupportedCaptureActionsFunc returns the method "MediaStreamTrack.getSupportedCaptureActions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) GetSupportedCaptureActionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetSupportedCaptureActionsFunc(
			this.Ref(),
		),
	)
}

// SendCaptureAction calls the method "MediaStreamTrack.sendCaptureAction".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) SendCaptureAction(action CaptureAction) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamTrackSendCaptureAction(
		this.Ref(), js.Pointer(&_ok),
		uint32(action),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SendCaptureActionFunc returns the method "MediaStreamTrack.sendCaptureAction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) SendCaptureActionFunc() (fn js.Func[func(action CaptureAction) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackSendCaptureActionFunc(
			this.Ref(),
		),
	)
}

// GetCaptureHandle calls the method "MediaStreamTrack.getCaptureHandle".
//
// The returned bool will be false if there is no such method.
func (this MediaStreamTrack) GetCaptureHandle() (CaptureHandle, bool) {
	var _ret CaptureHandle
	_ok := js.True == bindings.CallMediaStreamTrackGetCaptureHandle(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetCaptureHandleFunc returns the method "MediaStreamTrack.getCaptureHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStreamTrack) GetCaptureHandleFunc() (fn js.Func[func() CaptureHandle]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetCaptureHandleFunc(
			this.Ref(),
		),
	)
}

func NewMediaStream(stream MediaStream) MediaStream {
	return MediaStream{}.FromRef(
		bindings.NewMediaStreamByMediaStream(
			stream.Ref()),
	)
}

func NewMediaStreamByMediaStream1(tracks js.Array[MediaStreamTrack]) MediaStream {
	return MediaStream{}.FromRef(
		bindings.NewMediaStreamByMediaStream1(
			tracks.Ref()),
	)
}

type MediaStream struct {
	EventTarget
}

func (this MediaStream) Once() MediaStream {
	this.Ref().Once()
	return this
}

func (this MediaStream) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaStream) FromRef(ref js.Ref) MediaStream {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaStream) Free() {
	this.Ref().Free()
}

// Id returns the value of property "MediaStream.id".
//
// The returned bool will be false if there is no such property.
func (this MediaStream) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Active returns the value of property "MediaStream.active".
//
// The returned bool will be false if there is no such property.
func (this MediaStream) Active() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamActive(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetAudioTracks calls the method "MediaStream.getAudioTracks".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) GetAudioTracks() (js.Array[MediaStreamTrack], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamGetAudioTracks(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[MediaStreamTrack]{}.FromRef(_ret), _ok
}

// GetAudioTracksFunc returns the method "MediaStream.getAudioTracks".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) GetAudioTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetAudioTracksFunc(
			this.Ref(),
		),
	)
}

// GetVideoTracks calls the method "MediaStream.getVideoTracks".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) GetVideoTracks() (js.Array[MediaStreamTrack], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamGetVideoTracks(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[MediaStreamTrack]{}.FromRef(_ret), _ok
}

// GetVideoTracksFunc returns the method "MediaStream.getVideoTracks".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) GetVideoTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetVideoTracksFunc(
			this.Ref(),
		),
	)
}

// GetTracks calls the method "MediaStream.getTracks".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) GetTracks() (js.Array[MediaStreamTrack], bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamGetTracks(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[MediaStreamTrack]{}.FromRef(_ret), _ok
}

// GetTracksFunc returns the method "MediaStream.getTracks".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) GetTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetTracksFunc(
			this.Ref(),
		),
	)
}

// GetTrackById calls the method "MediaStream.getTrackById".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) GetTrackById(trackId js.String) (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamGetTrackById(
		this.Ref(), js.Pointer(&_ok),
		trackId.Ref(),
	)

	return MediaStreamTrack{}.FromRef(_ret), _ok
}

// GetTrackByIdFunc returns the method "MediaStream.getTrackById".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) GetTrackByIdFunc() (fn js.Func[func(trackId js.String) MediaStreamTrack]) {
	return fn.FromRef(
		bindings.MediaStreamGetTrackByIdFunc(
			this.Ref(),
		),
	)
}

// AddTrack calls the method "MediaStream.addTrack".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) AddTrack(track MediaStreamTrack) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamAddTrack(
		this.Ref(), js.Pointer(&_ok),
		track.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddTrackFunc returns the method "MediaStream.addTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) AddTrackFunc() (fn js.Func[func(track MediaStreamTrack)]) {
	return fn.FromRef(
		bindings.MediaStreamAddTrackFunc(
			this.Ref(),
		),
	)
}

// RemoveTrack calls the method "MediaStream.removeTrack".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) RemoveTrack(track MediaStreamTrack) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamRemoveTrack(
		this.Ref(), js.Pointer(&_ok),
		track.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveTrackFunc returns the method "MediaStream.removeTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) RemoveTrackFunc() (fn js.Func[func(track MediaStreamTrack)]) {
	return fn.FromRef(
		bindings.MediaStreamRemoveTrackFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "MediaStream.clone".
//
// The returned bool will be false if there is no such method.
func (this MediaStream) Clone() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.CallMediaStreamClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaStream{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "MediaStream.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaStream) CloneFunc() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.MediaStreamCloneFunc(
			this.Ref(),
		),
	)
}

const (
	MediaError_MEDIA_ERR_ABORTED           uint16 = 1
	MediaError_MEDIA_ERR_NETWORK           uint16 = 2
	MediaError_MEDIA_ERR_DECODE            uint16 = 3
	MediaError_MEDIA_ERR_SRC_NOT_SUPPORTED uint16 = 4
)

type MediaError struct {
	ref js.Ref
}

func (this MediaError) Once() MediaError {
	this.Ref().Once()
	return this
}

func (this MediaError) Ref() js.Ref {
	return this.ref
}

func (this MediaError) FromRef(ref js.Ref) MediaError {
	this.ref = ref
	return this
}

func (this MediaError) Free() {
	this.Ref().Free()
}

// Code returns the value of property "MediaError.code".
//
// The returned bool will be false if there is no such property.
func (this MediaError) Code() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetMediaErrorCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Message returns the value of property "MediaError.message".
//
// The returned bool will be false if there is no such property.
func (this MediaError) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaErrorMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type EndOfStreamError uint32

const (
	_ EndOfStreamError = iota

	EndOfStreamError_NETWORK
	EndOfStreamError_DECODE
)

func (EndOfStreamError) FromRef(str js.Ref) EndOfStreamError {
	return EndOfStreamError(bindings.ConstOfEndOfStreamError(str))
}

func (x EndOfStreamError) String() (string, bool) {
	switch x {
	case EndOfStreamError_NETWORK:
		return "network", true
	case EndOfStreamError_DECODE:
		return "decode", true
	default:
		return "", false
	}
}

type MediaSourceHandle struct {
	ref js.Ref
}

func (this MediaSourceHandle) Once() MediaSourceHandle {
	this.Ref().Once()
	return this
}

func (this MediaSourceHandle) Ref() js.Ref {
	return this.ref
}

func (this MediaSourceHandle) FromRef(ref js.Ref) MediaSourceHandle {
	this.ref = ref
	return this
}

func (this MediaSourceHandle) Free() {
	this.Ref().Free()
}

type SourceBufferList struct {
	EventTarget
}

func (this SourceBufferList) Once() SourceBufferList {
	this.Ref().Once()
	return this
}

func (this SourceBufferList) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SourceBufferList) FromRef(ref js.Ref) SourceBufferList {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SourceBufferList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SourceBufferList.length".
//
// The returned bool will be false if there is no such property.
func (this SourceBufferList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSourceBufferListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "SourceBufferList.".
//
// The returned bool will be false if there is no such method.
func (this SourceBufferList) Get(index uint32) (SourceBuffer, bool) {
	var _ok bool
	_ret := bindings.CallSourceBufferListGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SourceBuffer{}.FromRef(_ret), _ok
}

// GetFunc returns the method "SourceBufferList.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SourceBufferList) GetFunc() (fn js.Func[func(index uint32) SourceBuffer]) {
	return fn.FromRef(
		bindings.SourceBufferListGetFunc(
			this.Ref(),
		),
	)
}

type ReadyState uint32

const (
	_ ReadyState = iota

	ReadyState_CLOSED
	ReadyState_OPEN
	ReadyState_ENDED
)

func (ReadyState) FromRef(str js.Ref) ReadyState {
	return ReadyState(bindings.ConstOfReadyState(str))
}

func (x ReadyState) String() (string, bool) {
	switch x {
	case ReadyState_CLOSED:
		return "closed", true
	case ReadyState_OPEN:
		return "open", true
	case ReadyState_ENDED:
		return "ended", true
	default:
		return "", false
	}
}

type MediaSource struct {
	EventTarget
}

func (this MediaSource) Once() MediaSource {
	this.Ref().Once()
	return this
}

func (this MediaSource) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaSource) FromRef(ref js.Ref) MediaSource {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaSource) Free() {
	this.Ref().Free()
}

// Handle returns the value of property "MediaSource.handle".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) Handle() (MediaSourceHandle, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceHandle(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaSourceHandle{}.FromRef(_ret), _ok
}

// SourceBuffers returns the value of property "MediaSource.sourceBuffers".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) SourceBuffers() (SourceBufferList, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceSourceBuffers(
		this.Ref(), js.Pointer(&_ok),
	)
	return SourceBufferList{}.FromRef(_ret), _ok
}

// ActiveSourceBuffers returns the value of property "MediaSource.activeSourceBuffers".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) ActiveSourceBuffers() (SourceBufferList, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceActiveSourceBuffers(
		this.Ref(), js.Pointer(&_ok),
	)
	return SourceBufferList{}.FromRef(_ret), _ok
}

// ReadyState returns the value of property "MediaSource.readyState".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) ReadyState() (ReadyState, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadyState(_ret), _ok
}

// Duration returns the value of property "MediaSource.duration".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) Duration() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Duration sets the value of property "MediaSource.duration" to val.
//
// It returns false if the property cannot be set.
func (this MediaSource) SetDuration(val float64) bool {
	return js.True == bindings.SetMediaSourceDuration(
		this.Ref(),
		float64(val),
	)
}

// CanConstructInDedicatedWorker returns the value of property "MediaSource.canConstructInDedicatedWorker".
//
// The returned bool will be false if there is no such property.
func (this MediaSource) CanConstructInDedicatedWorker() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaSourceCanConstructInDedicatedWorker(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AddSourceBuffer calls the method "MediaSource.addSourceBuffer".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) AddSourceBuffer(typ js.String) (SourceBuffer, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceAddSourceBuffer(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return SourceBuffer{}.FromRef(_ret), _ok
}

// AddSourceBufferFunc returns the method "MediaSource.addSourceBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) AddSourceBufferFunc() (fn js.Func[func(typ js.String) SourceBuffer]) {
	return fn.FromRef(
		bindings.MediaSourceAddSourceBufferFunc(
			this.Ref(),
		),
	)
}

// RemoveSourceBuffer calls the method "MediaSource.removeSourceBuffer".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) RemoveSourceBuffer(sourceBuffer SourceBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceRemoveSourceBuffer(
		this.Ref(), js.Pointer(&_ok),
		sourceBuffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveSourceBufferFunc returns the method "MediaSource.removeSourceBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) RemoveSourceBufferFunc() (fn js.Func[func(sourceBuffer SourceBuffer)]) {
	return fn.FromRef(
		bindings.MediaSourceRemoveSourceBufferFunc(
			this.Ref(),
		),
	)
}

// EndOfStream calls the method "MediaSource.endOfStream".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) EndOfStream(err EndOfStreamError) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceEndOfStream(
		this.Ref(), js.Pointer(&_ok),
		uint32(err),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndOfStreamFunc returns the method "MediaSource.endOfStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) EndOfStreamFunc() (fn js.Func[func(err EndOfStreamError)]) {
	return fn.FromRef(
		bindings.MediaSourceEndOfStreamFunc(
			this.Ref(),
		),
	)
}

// EndOfStream1 calls the method "MediaSource.endOfStream".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) EndOfStream1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceEndOfStream1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndOfStream1Func returns the method "MediaSource.endOfStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) EndOfStream1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSourceEndOfStream1Func(
			this.Ref(),
		),
	)
}

// SetLiveSeekableRange calls the method "MediaSource.setLiveSeekableRange".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) SetLiveSeekableRange(start float64, end float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceSetLiveSeekableRange(
		this.Ref(), js.Pointer(&_ok),
		float64(start),
		float64(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetLiveSeekableRangeFunc returns the method "MediaSource.setLiveSeekableRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) SetLiveSeekableRangeFunc() (fn js.Func[func(start float64, end float64)]) {
	return fn.FromRef(
		bindings.MediaSourceSetLiveSeekableRangeFunc(
			this.Ref(),
		),
	)
}

// ClearLiveSeekableRange calls the method "MediaSource.clearLiveSeekableRange".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) ClearLiveSeekableRange() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceClearLiveSeekableRange(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearLiveSeekableRangeFunc returns the method "MediaSource.clearLiveSeekableRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) ClearLiveSeekableRangeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSourceClearLiveSeekableRangeFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "MediaSource.isTypeSupported".
//
// The returned bool will be false if there is no such method.
func (this MediaSource) IsTypeSupported(typ js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallMediaSourceIsTypeSupported(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return _ret == js.True, _ok
}

// IsTypeSupportedFunc returns the staticmethod "MediaSource.isTypeSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaSource) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.MediaSourceIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

type OneOf_MediaStream_MediaSource_Blob struct {
	ref js.Ref
}

func (x OneOf_MediaStream_MediaSource_Blob) Ref() js.Ref {
	return x.ref
}

func (x OneOf_MediaStream_MediaSource_Blob) Free() {
	x.ref.Free()
}

func (x OneOf_MediaStream_MediaSource_Blob) FromRef(ref js.Ref) OneOf_MediaStream_MediaSource_Blob {
	return OneOf_MediaStream_MediaSource_Blob{
		ref: ref,
	}
}

func (x OneOf_MediaStream_MediaSource_Blob) MediaStream() MediaStream {
	return MediaStream{}.FromRef(x.ref)
}

func (x OneOf_MediaStream_MediaSource_Blob) MediaSource() MediaSource {
	return MediaSource{}.FromRef(x.ref)
}

func (x OneOf_MediaStream_MediaSource_Blob) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

type MediaProvider = OneOf_MediaStream_MediaSource_Blob

type RemotePlaybackAvailabilityCallbackFunc func(this js.Ref, available bool) js.Ref

func (fn RemotePlaybackAvailabilityCallbackFunc) Register() js.Func[func(available bool)] {
	return js.RegisterCallback[func(available bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemotePlaybackAvailabilityCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemotePlaybackAvailabilityCallback[T any] struct {
	Fn  func(arg T, this js.Ref, available bool) js.Ref
	Arg T
}

func (cb *RemotePlaybackAvailabilityCallback[T]) Register() js.Func[func(available bool)] {
	return js.RegisterCallback[func(available bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemotePlaybackAvailabilityCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemotePlaybackState uint32

const (
	_ RemotePlaybackState = iota

	RemotePlaybackState_CONNECTING
	RemotePlaybackState_CONNECTED
	RemotePlaybackState_DISCONNECTED
)

func (RemotePlaybackState) FromRef(str js.Ref) RemotePlaybackState {
	return RemotePlaybackState(bindings.ConstOfRemotePlaybackState(str))
}

func (x RemotePlaybackState) String() (string, bool) {
	switch x {
	case RemotePlaybackState_CONNECTING:
		return "connecting", true
	case RemotePlaybackState_CONNECTED:
		return "connected", true
	case RemotePlaybackState_DISCONNECTED:
		return "disconnected", true
	default:
		return "", false
	}
}

type RemotePlayback struct {
	EventTarget
}

func (this RemotePlayback) Once() RemotePlayback {
	this.Ref().Once()
	return this
}

func (this RemotePlayback) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this RemotePlayback) FromRef(ref js.Ref) RemotePlayback {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this RemotePlayback) Free() {
	this.Ref().Free()
}

// State returns the value of property "RemotePlayback.state".
//
// The returned bool will be false if there is no such property.
func (this RemotePlayback) State() (RemotePlaybackState, bool) {
	var _ok bool
	_ret := bindings.GetRemotePlaybackState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RemotePlaybackState(_ret), _ok
}

// WatchAvailability calls the method "RemotePlayback.watchAvailability".
//
// The returned bool will be false if there is no such method.
func (this RemotePlayback) WatchAvailability(callback js.Func[func(available bool)]) (js.Promise[js.Number[int32]], bool) {
	var _ok bool
	_ret := bindings.CallRemotePlaybackWatchAvailability(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return js.Promise[js.Number[int32]]{}.FromRef(_ret), _ok
}

// WatchAvailabilityFunc returns the method "RemotePlayback.watchAvailability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RemotePlayback) WatchAvailabilityFunc() (fn js.Func[func(callback js.Func[func(available bool)]) js.Promise[js.Number[int32]]]) {
	return fn.FromRef(
		bindings.RemotePlaybackWatchAvailabilityFunc(
			this.Ref(),
		),
	)
}

// CancelWatchAvailability calls the method "RemotePlayback.cancelWatchAvailability".
//
// The returned bool will be false if there is no such method.
func (this RemotePlayback) CancelWatchAvailability(id int32) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRemotePlaybackCancelWatchAvailability(
		this.Ref(), js.Pointer(&_ok),
		int32(id),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CancelWatchAvailabilityFunc returns the method "RemotePlayback.cancelWatchAvailability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RemotePlayback) CancelWatchAvailabilityFunc() (fn js.Func[func(id int32) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackCancelWatchAvailabilityFunc(
			this.Ref(),
		),
	)
}

// CancelWatchAvailability1 calls the method "RemotePlayback.cancelWatchAvailability".
//
// The returned bool will be false if there is no such method.
func (this RemotePlayback) CancelWatchAvailability1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRemotePlaybackCancelWatchAvailability1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CancelWatchAvailability1Func returns the method "RemotePlayback.cancelWatchAvailability".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RemotePlayback) CancelWatchAvailability1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackCancelWatchAvailability1Func(
			this.Ref(),
		),
	)
}

// Prompt calls the method "RemotePlayback.prompt".
//
// The returned bool will be false if there is no such method.
func (this RemotePlayback) Prompt() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallRemotePlaybackPrompt(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PromptFunc returns the method "RemotePlayback.prompt".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RemotePlayback) PromptFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackPromptFunc(
			this.Ref(),
		),
	)
}

type HTMLMediaElement struct {
	HTMLElement
}

func (this HTMLMediaElement) Once() HTMLMediaElement {
	this.Ref().Once()
	return this
}

func (this HTMLMediaElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMediaElement) FromRef(ref js.Ref) HTMLMediaElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMediaElement) Free() {
	this.Ref().Free()
}

// Error returns the value of property "HTMLMediaElement.error".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Error() (MediaError, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementError(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaError{}.FromRef(_ret), _ok
}

// Src returns the value of property "HTMLMediaElement.src".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLMediaElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLMediaElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// SrcObject returns the value of property "HTMLMediaElement.srcObject".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) SrcObject() (MediaProvider, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementSrcObject(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaProvider{}.FromRef(_ret), _ok
}

// SrcObject sets the value of property "HTMLMediaElement.srcObject" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetSrcObject(val MediaProvider) bool {
	return js.True == bindings.SetHTMLMediaElementSrcObject(
		this.Ref(),
		val.Ref(),
	)
}

// CurrentSrc returns the value of property "HTMLMediaElement.currentSrc".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) CurrentSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementCurrentSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin returns the value of property "HTMLMediaElement.crossOrigin".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "HTMLMediaElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLMediaElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// NetworkState returns the value of property "HTMLMediaElement.networkState".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) NetworkState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementNetworkState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Preload returns the value of property "HTMLMediaElement.preload".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Preload() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementPreload(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Preload sets the value of property "HTMLMediaElement.preload" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetPreload(val js.String) bool {
	return js.True == bindings.SetHTMLMediaElementPreload(
		this.Ref(),
		val.Ref(),
	)
}

// Buffered returns the value of property "HTMLMediaElement.buffered".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Buffered() (TimeRanges, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementBuffered(
		this.Ref(), js.Pointer(&_ok),
	)
	return TimeRanges{}.FromRef(_ret), _ok
}

// ReadyState returns the value of property "HTMLMediaElement.readyState".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Seeking returns the value of property "HTMLMediaElement.seeking".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Seeking() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementSeeking(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CurrentTime returns the value of property "HTMLMediaElement.currentTime".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) CurrentTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// CurrentTime sets the value of property "HTMLMediaElement.currentTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetCurrentTime(val float64) bool {
	return js.True == bindings.SetHTMLMediaElementCurrentTime(
		this.Ref(),
		float64(val),
	)
}

// Duration returns the value of property "HTMLMediaElement.duration".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Duration() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Paused returns the value of property "HTMLMediaElement.paused".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Paused() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementPaused(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DefaultPlaybackRate returns the value of property "HTMLMediaElement.defaultPlaybackRate".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) DefaultPlaybackRate() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementDefaultPlaybackRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DefaultPlaybackRate sets the value of property "HTMLMediaElement.defaultPlaybackRate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetDefaultPlaybackRate(val float64) bool {
	return js.True == bindings.SetHTMLMediaElementDefaultPlaybackRate(
		this.Ref(),
		float64(val),
	)
}

// PlaybackRate returns the value of property "HTMLMediaElement.playbackRate".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) PlaybackRate() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementPlaybackRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PlaybackRate sets the value of property "HTMLMediaElement.playbackRate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetPlaybackRate(val float64) bool {
	return js.True == bindings.SetHTMLMediaElementPlaybackRate(
		this.Ref(),
		float64(val),
	)
}

// PreservesPitch returns the value of property "HTMLMediaElement.preservesPitch".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) PreservesPitch() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementPreservesPitch(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// PreservesPitch sets the value of property "HTMLMediaElement.preservesPitch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetPreservesPitch(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementPreservesPitch(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Played returns the value of property "HTMLMediaElement.played".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Played() (TimeRanges, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementPlayed(
		this.Ref(), js.Pointer(&_ok),
	)
	return TimeRanges{}.FromRef(_ret), _ok
}

// Seekable returns the value of property "HTMLMediaElement.seekable".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Seekable() (TimeRanges, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementSeekable(
		this.Ref(), js.Pointer(&_ok),
	)
	return TimeRanges{}.FromRef(_ret), _ok
}

// Ended returns the value of property "HTMLMediaElement.ended".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Ended() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementEnded(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Autoplay returns the value of property "HTMLMediaElement.autoplay".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Autoplay() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementAutoplay(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Autoplay sets the value of property "HTMLMediaElement.autoplay" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetAutoplay(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementAutoplay(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Loop returns the value of property "HTMLMediaElement.loop".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Loop() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementLoop(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Loop sets the value of property "HTMLMediaElement.loop" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetLoop(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementLoop(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Controls returns the value of property "HTMLMediaElement.controls".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Controls() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementControls(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Controls sets the value of property "HTMLMediaElement.controls" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetControls(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementControls(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Volume returns the value of property "HTMLMediaElement.volume".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Volume() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementVolume(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Volume sets the value of property "HTMLMediaElement.volume" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetVolume(val float64) bool {
	return js.True == bindings.SetHTMLMediaElementVolume(
		this.Ref(),
		float64(val),
	)
}

// Muted returns the value of property "HTMLMediaElement.muted".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Muted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementMuted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Muted sets the value of property "HTMLMediaElement.muted" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetMuted(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementMuted(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// DefaultMuted returns the value of property "HTMLMediaElement.defaultMuted".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) DefaultMuted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementDefaultMuted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DefaultMuted sets the value of property "HTMLMediaElement.defaultMuted" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetDefaultMuted(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementDefaultMuted(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// AudioTracks returns the value of property "HTMLMediaElement.audioTracks".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) AudioTracks() (AudioTrackList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementAudioTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioTrackList{}.FromRef(_ret), _ok
}

// VideoTracks returns the value of property "HTMLMediaElement.videoTracks".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) VideoTracks() (VideoTrackList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementVideoTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return VideoTrackList{}.FromRef(_ret), _ok
}

// TextTracks returns the value of property "HTMLMediaElement.textTracks".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) TextTracks() (TextTrackList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementTextTracks(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrackList{}.FromRef(_ret), _ok
}

// MediaKeys returns the value of property "HTMLMediaElement.mediaKeys".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) MediaKeys() (MediaKeys, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementMediaKeys(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaKeys{}.FromRef(_ret), _ok
}

// Remote returns the value of property "HTMLMediaElement.remote".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) Remote() (RemotePlayback, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementRemote(
		this.Ref(), js.Pointer(&_ok),
	)
	return RemotePlayback{}.FromRef(_ret), _ok
}

// DisableRemotePlayback returns the value of property "HTMLMediaElement.disableRemotePlayback".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) DisableRemotePlayback() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementDisableRemotePlayback(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DisableRemotePlayback sets the value of property "HTMLMediaElement.disableRemotePlayback" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMediaElement) SetDisableRemotePlayback(val bool) bool {
	return js.True == bindings.SetHTMLMediaElementDisableRemotePlayback(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// SinkId returns the value of property "HTMLMediaElement.sinkId".
//
// The returned bool will be false if there is no such property.
func (this HTMLMediaElement) SinkId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMediaElementSinkId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Load calls the method "HTMLMediaElement.load".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) Load() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementLoad(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LoadFunc returns the method "HTMLMediaElement.load".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) LoadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMediaElementLoadFunc(
			this.Ref(),
		),
	)
}

// CanPlayType calls the method "HTMLMediaElement.canPlayType".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) CanPlayType(typ js.String) (CanPlayTypeResult, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementCanPlayType(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return CanPlayTypeResult(_ret), _ok
}

// CanPlayTypeFunc returns the method "HTMLMediaElement.canPlayType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) CanPlayTypeFunc() (fn js.Func[func(typ js.String) CanPlayTypeResult]) {
	return fn.FromRef(
		bindings.HTMLMediaElementCanPlayTypeFunc(
			this.Ref(),
		),
	)
}

// FastSeek calls the method "HTMLMediaElement.fastSeek".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) FastSeek(time float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementFastSeek(
		this.Ref(), js.Pointer(&_ok),
		float64(time),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FastSeekFunc returns the method "HTMLMediaElement.fastSeek".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) FastSeekFunc() (fn js.Func[func(time float64)]) {
	return fn.FromRef(
		bindings.HTMLMediaElementFastSeekFunc(
			this.Ref(),
		),
	)
}

// GetStartDate calls the method "HTMLMediaElement.getStartDate".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) GetStartDate() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementGetStartDate(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// GetStartDateFunc returns the method "HTMLMediaElement.getStartDate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) GetStartDateFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.HTMLMediaElementGetStartDateFunc(
			this.Ref(),
		),
	)
}

// Play calls the method "HTMLMediaElement.play".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) Play() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementPlay(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// PlayFunc returns the method "HTMLMediaElement.play".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) PlayFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementPlayFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "HTMLMediaElement.pause".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) Pause() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementPause(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseFunc returns the method "HTMLMediaElement.pause".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMediaElementPauseFunc(
			this.Ref(),
		),
	)
}

// AddTextTrack calls the method "HTMLMediaElement.addTextTrack".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) AddTextTrack(kind TextTrackKind, label js.String, language js.String) (TextTrack, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementAddTextTrack(
		this.Ref(), js.Pointer(&_ok),
		uint32(kind),
		label.Ref(),
		language.Ref(),
	)

	return TextTrack{}.FromRef(_ret), _ok
}

// AddTextTrackFunc returns the method "HTMLMediaElement.addTextTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) AddTextTrackFunc() (fn js.Func[func(kind TextTrackKind, label js.String, language js.String) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrackFunc(
			this.Ref(),
		),
	)
}

// AddTextTrack1 calls the method "HTMLMediaElement.addTextTrack".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) AddTextTrack1(kind TextTrackKind, label js.String) (TextTrack, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementAddTextTrack1(
		this.Ref(), js.Pointer(&_ok),
		uint32(kind),
		label.Ref(),
	)

	return TextTrack{}.FromRef(_ret), _ok
}

// AddTextTrack1Func returns the method "HTMLMediaElement.addTextTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) AddTextTrack1Func() (fn js.Func[func(kind TextTrackKind, label js.String) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrack1Func(
			this.Ref(),
		),
	)
}

// AddTextTrack2 calls the method "HTMLMediaElement.addTextTrack".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) AddTextTrack2(kind TextTrackKind) (TextTrack, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementAddTextTrack2(
		this.Ref(), js.Pointer(&_ok),
		uint32(kind),
	)

	return TextTrack{}.FromRef(_ret), _ok
}

// AddTextTrack2Func returns the method "HTMLMediaElement.addTextTrack".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) AddTextTrack2Func() (fn js.Func[func(kind TextTrackKind) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrack2Func(
			this.Ref(),
		),
	)
}

// SetMediaKeys calls the method "HTMLMediaElement.setMediaKeys".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) SetMediaKeys(mediaKeys MediaKeys) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementSetMediaKeys(
		this.Ref(), js.Pointer(&_ok),
		mediaKeys.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetMediaKeysFunc returns the method "HTMLMediaElement.setMediaKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) SetMediaKeysFunc() (fn js.Func[func(mediaKeys MediaKeys) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementSetMediaKeysFunc(
			this.Ref(),
		),
	)
}

// CaptureStream calls the method "HTMLMediaElement.captureStream".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) CaptureStream() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementCaptureStream(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaStream{}.FromRef(_ret), _ok
}

// CaptureStreamFunc returns the method "HTMLMediaElement.captureStream".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) CaptureStreamFunc() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.HTMLMediaElementCaptureStreamFunc(
			this.Ref(),
		),
	)
}

// SetSinkId calls the method "HTMLMediaElement.setSinkId".
//
// The returned bool will be false if there is no such method.
func (this HTMLMediaElement) SetSinkId(sinkId js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLMediaElementSetSinkId(
		this.Ref(), js.Pointer(&_ok),
		sinkId.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetSinkIdFunc returns the method "HTMLMediaElement.setSinkId".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMediaElement) SetSinkIdFunc() (fn js.Func[func(sinkId js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementSetSinkIdFunc(
			this.Ref(),
		),
	)
}

type MediaElementAudioSourceOptions struct {
	// MediaElement is "MediaElementAudioSourceOptions.mediaElement"
	//
	// Required
	MediaElement HTMLMediaElement

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaElementAudioSourceOptions with all fields set.
func (p MediaElementAudioSourceOptions) FromRef(ref js.Ref) MediaElementAudioSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaElementAudioSourceOptions in the application heap.
func (p MediaElementAudioSourceOptions) New() js.Ref {
	return bindings.MediaElementAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaElementAudioSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.MediaElementAudioSourceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaElementAudioSourceOptions) Update(ref js.Ref) {
	bindings.MediaElementAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaElementAudioSourceNode(context AudioContext, options MediaElementAudioSourceOptions) MediaElementAudioSourceNode {
	return MediaElementAudioSourceNode{}.FromRef(
		bindings.NewMediaElementAudioSourceNodeByMediaElementAudioSourceNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

type MediaElementAudioSourceNode struct {
	AudioNode
}

func (this MediaElementAudioSourceNode) Once() MediaElementAudioSourceNode {
	this.Ref().Once()
	return this
}

func (this MediaElementAudioSourceNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this MediaElementAudioSourceNode) FromRef(ref js.Ref) MediaElementAudioSourceNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this MediaElementAudioSourceNode) Free() {
	this.Ref().Free()
}

// MediaElement returns the value of property "MediaElementAudioSourceNode.mediaElement".
//
// The returned bool will be false if there is no such property.
func (this MediaElementAudioSourceNode) MediaElement() (HTMLMediaElement, bool) {
	var _ok bool
	_ret := bindings.GetMediaElementAudioSourceNodeMediaElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLMediaElement{}.FromRef(_ret), _ok
}

type MediaStreamAudioSourceOptions struct {
	// MediaStream is "MediaStreamAudioSourceOptions.mediaStream"
	//
	// Required
	MediaStream MediaStream

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamAudioSourceOptions with all fields set.
func (p MediaStreamAudioSourceOptions) FromRef(ref js.Ref) MediaStreamAudioSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamAudioSourceOptions in the application heap.
func (p MediaStreamAudioSourceOptions) New() js.Ref {
	return bindings.MediaStreamAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaStreamAudioSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamAudioSourceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaStreamAudioSourceOptions) Update(ref js.Ref) {
	bindings.MediaStreamAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaStreamAudioSourceNode(context AudioContext, options MediaStreamAudioSourceOptions) MediaStreamAudioSourceNode {
	return MediaStreamAudioSourceNode{}.FromRef(
		bindings.NewMediaStreamAudioSourceNodeByMediaStreamAudioSourceNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

type MediaStreamAudioSourceNode struct {
	AudioNode
}

func (this MediaStreamAudioSourceNode) Once() MediaStreamAudioSourceNode {
	this.Ref().Once()
	return this
}

func (this MediaStreamAudioSourceNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this MediaStreamAudioSourceNode) FromRef(ref js.Ref) MediaStreamAudioSourceNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this MediaStreamAudioSourceNode) Free() {
	this.Ref().Free()
}

// MediaStream returns the value of property "MediaStreamAudioSourceNode.mediaStream".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamAudioSourceNode) MediaStream() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamAudioSourceNodeMediaStream(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStream{}.FromRef(_ret), _ok
}

type MediaStreamTrackAudioSourceOptions struct {
	// MediaStreamTrack is "MediaStreamTrackAudioSourceOptions.mediaStreamTrack"
	//
	// Required
	MediaStreamTrack MediaStreamTrack

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamTrackAudioSourceOptions with all fields set.
func (p MediaStreamTrackAudioSourceOptions) FromRef(ref js.Ref) MediaStreamTrackAudioSourceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamTrackAudioSourceOptions in the application heap.
func (p MediaStreamTrackAudioSourceOptions) New() js.Ref {
	return bindings.MediaStreamTrackAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaStreamTrackAudioSourceOptions) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamTrackAudioSourceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaStreamTrackAudioSourceOptions) Update(ref js.Ref) {
	bindings.MediaStreamTrackAudioSourceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaStreamTrackAudioSourceNode(context AudioContext, options MediaStreamTrackAudioSourceOptions) MediaStreamTrackAudioSourceNode {
	return MediaStreamTrackAudioSourceNode{}.FromRef(
		bindings.NewMediaStreamTrackAudioSourceNodeByMediaStreamTrackAudioSourceNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

type MediaStreamTrackAudioSourceNode struct {
	AudioNode
}

func (this MediaStreamTrackAudioSourceNode) Once() MediaStreamTrackAudioSourceNode {
	this.Ref().Once()
	return this
}

func (this MediaStreamTrackAudioSourceNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this MediaStreamTrackAudioSourceNode) FromRef(ref js.Ref) MediaStreamTrackAudioSourceNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this MediaStreamTrackAudioSourceNode) Free() {
	this.Ref().Free()
}

type AudioNodeOptions struct {
	// ChannelCount is "AudioNodeOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount uint32
	// ChannelCountMode is "AudioNodeOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "AudioNodeOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioNodeOptions with all fields set.
func (p AudioNodeOptions) FromRef(ref js.Ref) AudioNodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioNodeOptions in the application heap.
func (p AudioNodeOptions) New() js.Ref {
	return bindings.AudioNodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioNodeOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioNodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioNodeOptions) Update(ref js.Ref) {
	bindings.AudioNodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaStreamAudioDestinationNode(context AudioContext, options AudioNodeOptions) MediaStreamAudioDestinationNode {
	return MediaStreamAudioDestinationNode{}.FromRef(
		bindings.NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode1(context AudioContext) MediaStreamAudioDestinationNode {
	return MediaStreamAudioDestinationNode{}.FromRef(
		bindings.NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode1(
			context.Ref()),
	)
}

type MediaStreamAudioDestinationNode struct {
	AudioNode
}

func (this MediaStreamAudioDestinationNode) Once() MediaStreamAudioDestinationNode {
	this.Ref().Once()
	return this
}

func (this MediaStreamAudioDestinationNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this MediaStreamAudioDestinationNode) FromRef(ref js.Ref) MediaStreamAudioDestinationNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this MediaStreamAudioDestinationNode) Free() {
	this.Ref().Free()
}

// Stream returns the value of property "MediaStreamAudioDestinationNode.stream".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamAudioDestinationNode) Stream() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamAudioDestinationNodeStream(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStream{}.FromRef(_ret), _ok
}

type AudioSinkInfo struct {
	ref js.Ref
}

func (this AudioSinkInfo) Once() AudioSinkInfo {
	this.Ref().Once()
	return this
}

func (this AudioSinkInfo) Ref() js.Ref {
	return this.ref
}

func (this AudioSinkInfo) FromRef(ref js.Ref) AudioSinkInfo {
	this.ref = ref
	return this
}

func (this AudioSinkInfo) Free() {
	this.Ref().Free()
}

// Type returns the value of property "AudioSinkInfo.type".
//
// The returned bool will be false if there is no such property.
func (this AudioSinkInfo) Type() (AudioSinkType, bool) {
	var _ok bool
	_ret := bindings.GetAudioSinkInfoType(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioSinkType(_ret), _ok
}

type OneOf_String_AudioSinkInfo struct {
	ref js.Ref
}

func (x OneOf_String_AudioSinkInfo) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_AudioSinkInfo) Free() {
	x.ref.Free()
}

func (x OneOf_String_AudioSinkInfo) FromRef(ref js.Ref) OneOf_String_AudioSinkInfo {
	return OneOf_String_AudioSinkInfo{
		ref: ref,
	}
}

func (x OneOf_String_AudioSinkInfo) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_AudioSinkInfo) AudioSinkInfo() AudioSinkInfo {
	return AudioSinkInfo{}.FromRef(x.ref)
}

type AudioRenderCapacityOptions struct {
	// UpdateInterval is "AudioRenderCapacityOptions.updateInterval"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_UpdateInterval MUST be set to true to make this field effective.
	UpdateInterval float64

	FFI_USE_UpdateInterval bool // for UpdateInterval.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioRenderCapacityOptions with all fields set.
func (p AudioRenderCapacityOptions) FromRef(ref js.Ref) AudioRenderCapacityOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioRenderCapacityOptions in the application heap.
func (p AudioRenderCapacityOptions) New() js.Ref {
	return bindings.AudioRenderCapacityOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioRenderCapacityOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioRenderCapacityOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioRenderCapacityOptions) Update(ref js.Ref) {
	bindings.AudioRenderCapacityOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioRenderCapacity struct {
	EventTarget
}

func (this AudioRenderCapacity) Once() AudioRenderCapacity {
	this.Ref().Once()
	return this
}

func (this AudioRenderCapacity) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AudioRenderCapacity) FromRef(ref js.Ref) AudioRenderCapacity {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AudioRenderCapacity) Free() {
	this.Ref().Free()
}

// Start calls the method "AudioRenderCapacity.start".
//
// The returned bool will be false if there is no such method.
func (this AudioRenderCapacity) Start(options AudioRenderCapacityOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioRenderCapacityStart(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "AudioRenderCapacity.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioRenderCapacity) StartFunc() (fn js.Func[func(options AudioRenderCapacityOptions)]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStartFunc(
			this.Ref(),
		),
	)
}

// Start1 calls the method "AudioRenderCapacity.start".
//
// The returned bool will be false if there is no such method.
func (this AudioRenderCapacity) Start1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioRenderCapacityStart1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start1Func returns the method "AudioRenderCapacity.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioRenderCapacity) Start1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStart1Func(
			this.Ref(),
		),
	)
}

// Stop calls the method "AudioRenderCapacity.stop".
//
// The returned bool will be false if there is no such method.
func (this AudioRenderCapacity) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioRenderCapacityStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "AudioRenderCapacity.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioRenderCapacity) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStopFunc(
			this.Ref(),
		),
	)
}

func NewAudioContext(contextOptions AudioContextOptions) AudioContext {
	return AudioContext{}.FromRef(
		bindings.NewAudioContextByAudioContext(
			js.Pointer(&contextOptions)),
	)
}

func NewAudioContextByAudioContext1() AudioContext {
	return AudioContext{}.FromRef(
		bindings.NewAudioContextByAudioContext1(),
	)
}

type AudioContext struct {
	BaseAudioContext
}

func (this AudioContext) Once() AudioContext {
	this.Ref().Once()
	return this
}

func (this AudioContext) Ref() js.Ref {
	return this.BaseAudioContext.Ref()
}

func (this AudioContext) FromRef(ref js.Ref) AudioContext {
	this.BaseAudioContext = this.BaseAudioContext.FromRef(ref)
	return this
}

func (this AudioContext) Free() {
	this.Ref().Free()
}

// BaseLatency returns the value of property "AudioContext.baseLatency".
//
// The returned bool will be false if there is no such property.
func (this AudioContext) BaseLatency() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAudioContextBaseLatency(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OutputLatency returns the value of property "AudioContext.outputLatency".
//
// The returned bool will be false if there is no such property.
func (this AudioContext) OutputLatency() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAudioContextOutputLatency(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// SinkId returns the value of property "AudioContext.sinkId".
//
// The returned bool will be false if there is no such property.
func (this AudioContext) SinkId() (OneOf_String_AudioSinkInfo, bool) {
	var _ok bool
	_ret := bindings.GetAudioContextSinkId(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_AudioSinkInfo{}.FromRef(_ret), _ok
}

// RenderCapacity returns the value of property "AudioContext.renderCapacity".
//
// The returned bool will be false if there is no such property.
func (this AudioContext) RenderCapacity() (AudioRenderCapacity, bool) {
	var _ok bool
	_ret := bindings.GetAudioContextRenderCapacity(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioRenderCapacity{}.FromRef(_ret), _ok
}

// GetOutputTimestamp calls the method "AudioContext.getOutputTimestamp".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) GetOutputTimestamp() (AudioTimestamp, bool) {
	var _ret AudioTimestamp
	_ok := js.True == bindings.CallAudioContextGetOutputTimestamp(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetOutputTimestampFunc returns the method "AudioContext.getOutputTimestamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) GetOutputTimestampFunc() (fn js.Func[func() AudioTimestamp]) {
	return fn.FromRef(
		bindings.AudioContextGetOutputTimestampFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "AudioContext.resume".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) Resume() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallAudioContextResume(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ResumeFunc returns the method "AudioContext.resume".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) ResumeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextResumeFunc(
			this.Ref(),
		),
	)
}

// Suspend calls the method "AudioContext.suspend".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) Suspend() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallAudioContextSuspend(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SuspendFunc returns the method "AudioContext.suspend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) SuspendFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextSuspendFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioContext.close".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) Close() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallAudioContextClose(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CloseFunc returns the method "AudioContext.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextCloseFunc(
			this.Ref(),
		),
	)
}

// SetSinkId calls the method "AudioContext.setSinkId".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) SetSinkId(sinkId OneOf_String_AudioSinkOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallAudioContextSetSinkId(
		this.Ref(), js.Pointer(&_ok),
		sinkId.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetSinkIdFunc returns the method "AudioContext.setSinkId".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) SetSinkIdFunc() (fn js.Func[func(sinkId OneOf_String_AudioSinkOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextSetSinkIdFunc(
			this.Ref(),
		),
	)
}

// CreateMediaElementSource calls the method "AudioContext.createMediaElementSource".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) CreateMediaElementSource(mediaElement HTMLMediaElement) (MediaElementAudioSourceNode, bool) {
	var _ok bool
	_ret := bindings.CallAudioContextCreateMediaElementSource(
		this.Ref(), js.Pointer(&_ok),
		mediaElement.Ref(),
	)

	return MediaElementAudioSourceNode{}.FromRef(_ret), _ok
}

// CreateMediaElementSourceFunc returns the method "AudioContext.createMediaElementSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) CreateMediaElementSourceFunc() (fn js.Func[func(mediaElement HTMLMediaElement) MediaElementAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaElementSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamSource calls the method "AudioContext.createMediaStreamSource".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) CreateMediaStreamSource(mediaStream MediaStream) (MediaStreamAudioSourceNode, bool) {
	var _ok bool
	_ret := bindings.CallAudioContextCreateMediaStreamSource(
		this.Ref(), js.Pointer(&_ok),
		mediaStream.Ref(),
	)

	return MediaStreamAudioSourceNode{}.FromRef(_ret), _ok
}

// CreateMediaStreamSourceFunc returns the method "AudioContext.createMediaStreamSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) CreateMediaStreamSourceFunc() (fn js.Func[func(mediaStream MediaStream) MediaStreamAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamTrackSource calls the method "AudioContext.createMediaStreamTrackSource".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) CreateMediaStreamTrackSource(mediaStreamTrack MediaStreamTrack) (MediaStreamTrackAudioSourceNode, bool) {
	var _ok bool
	_ret := bindings.CallAudioContextCreateMediaStreamTrackSource(
		this.Ref(), js.Pointer(&_ok),
		mediaStreamTrack.Ref(),
	)

	return MediaStreamTrackAudioSourceNode{}.FromRef(_ret), _ok
}

// CreateMediaStreamTrackSourceFunc returns the method "AudioContext.createMediaStreamTrackSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) CreateMediaStreamTrackSourceFunc() (fn js.Func[func(mediaStreamTrack MediaStreamTrack) MediaStreamTrackAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamTrackSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamDestination calls the method "AudioContext.createMediaStreamDestination".
//
// The returned bool will be false if there is no such method.
func (this AudioContext) CreateMediaStreamDestination() (MediaStreamAudioDestinationNode, bool) {
	var _ok bool
	_ret := bindings.CallAudioContextCreateMediaStreamDestination(
		this.Ref(), js.Pointer(&_ok),
	)

	return MediaStreamAudioDestinationNode{}.FromRef(_ret), _ok
}

// CreateMediaStreamDestinationFunc returns the method "AudioContext.createMediaStreamDestination".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioContext) CreateMediaStreamDestinationFunc() (fn js.Func[func() MediaStreamAudioDestinationNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamDestinationFunc(
			this.Ref(),
		),
	)
}

type AudioSampleFormat uint32

const (
	_ AudioSampleFormat = iota

	AudioSampleFormat_U8
	AudioSampleFormat_S16
	AudioSampleFormat_S32
	AudioSampleFormat_F32
	AudioSampleFormat_U8_PLANAR
	AudioSampleFormat_S16_PLANAR
	AudioSampleFormat_S32_PLANAR
	AudioSampleFormat_F32_PLANAR
)

func (AudioSampleFormat) FromRef(str js.Ref) AudioSampleFormat {
	return AudioSampleFormat(bindings.ConstOfAudioSampleFormat(str))
}

func (x AudioSampleFormat) String() (string, bool) {
	switch x {
	case AudioSampleFormat_U8:
		return "u8", true
	case AudioSampleFormat_S16:
		return "s16", true
	case AudioSampleFormat_S32:
		return "s32", true
	case AudioSampleFormat_F32:
		return "f32", true
	case AudioSampleFormat_U8_PLANAR:
		return "u8-planar", true
	case AudioSampleFormat_S16_PLANAR:
		return "s16-planar", true
	case AudioSampleFormat_S32_PLANAR:
		return "s32-planar", true
	case AudioSampleFormat_F32_PLANAR:
		return "f32-planar", true
	default:
		return "", false
	}
}

type AudioDataInit struct {
	// Format is "AudioDataInit.format"
	//
	// Required
	Format AudioSampleFormat
	// SampleRate is "AudioDataInit.sampleRate"
	//
	// Required
	SampleRate float32
	// NumberOfFrames is "AudioDataInit.numberOfFrames"
	//
	// Required
	NumberOfFrames uint32
	// NumberOfChannels is "AudioDataInit.numberOfChannels"
	//
	// Required
	NumberOfChannels uint32
	// Timestamp is "AudioDataInit.timestamp"
	//
	// Required
	Timestamp int64
	// Data is "AudioDataInit.data"
	//
	// Required
	Data BufferSource
	// Transfer is "AudioDataInit.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.ArrayBuffer]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDataInit with all fields set.
func (p AudioDataInit) FromRef(ref js.Ref) AudioDataInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDataInit in the application heap.
func (p AudioDataInit) New() js.Ref {
	return bindings.AudioDataInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioDataInit) UpdateFrom(ref js.Ref) {
	bindings.AudioDataInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioDataInit) Update(ref js.Ref) {
	bindings.AudioDataInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioDataCopyToOptions struct {
	// PlaneIndex is "AudioDataCopyToOptions.planeIndex"
	//
	// Required
	PlaneIndex uint32
	// FrameOffset is "AudioDataCopyToOptions.frameOffset"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_FrameOffset MUST be set to true to make this field effective.
	FrameOffset uint32
	// FrameCount is "AudioDataCopyToOptions.frameCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameCount MUST be set to true to make this field effective.
	FrameCount uint32
	// Format is "AudioDataCopyToOptions.format"
	//
	// Optional
	Format AudioSampleFormat

	FFI_USE_FrameOffset bool // for FrameOffset.
	FFI_USE_FrameCount  bool // for FrameCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDataCopyToOptions with all fields set.
func (p AudioDataCopyToOptions) FromRef(ref js.Ref) AudioDataCopyToOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDataCopyToOptions in the application heap.
func (p AudioDataCopyToOptions) New() js.Ref {
	return bindings.AudioDataCopyToOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioDataCopyToOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioDataCopyToOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioDataCopyToOptions) Update(ref js.Ref) {
	bindings.AudioDataCopyToOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAudioData(init AudioDataInit) AudioData {
	return AudioData{}.FromRef(
		bindings.NewAudioDataByAudioData(
			js.Pointer(&init)),
	)
}

type AudioData struct {
	ref js.Ref
}

func (this AudioData) Once() AudioData {
	this.Ref().Once()
	return this
}

func (this AudioData) Ref() js.Ref {
	return this.ref
}

func (this AudioData) FromRef(ref js.Ref) AudioData {
	this.ref = ref
	return this
}

func (this AudioData) Free() {
	this.Ref().Free()
}

// Format returns the value of property "AudioData.format".
//
// The returned bool will be false if there is no such property.
func (this AudioData) Format() (AudioSampleFormat, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataFormat(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioSampleFormat(_ret), _ok
}

// SampleRate returns the value of property "AudioData.sampleRate".
//
// The returned bool will be false if there is no such property.
func (this AudioData) SampleRate() (float32, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataSampleRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// NumberOfFrames returns the value of property "AudioData.numberOfFrames".
//
// The returned bool will be false if there is no such property.
func (this AudioData) NumberOfFrames() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataNumberOfFrames(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NumberOfChannels returns the value of property "AudioData.numberOfChannels".
//
// The returned bool will be false if there is no such property.
func (this AudioData) NumberOfChannels() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataNumberOfChannels(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Duration returns the value of property "AudioData.duration".
//
// The returned bool will be false if there is no such property.
func (this AudioData) Duration() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Timestamp returns the value of property "AudioData.timestamp".
//
// The returned bool will be false if there is no such property.
func (this AudioData) Timestamp() (int64, bool) {
	var _ok bool
	_ret := bindings.GetAudioDataTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// AllocationSize calls the method "AudioData.allocationSize".
//
// The returned bool will be false if there is no such method.
func (this AudioData) AllocationSize(options AudioDataCopyToOptions) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallAudioDataAllocationSize(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return uint32(_ret), _ok
}

// AllocationSizeFunc returns the method "AudioData.allocationSize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioData) AllocationSizeFunc() (fn js.Func[func(options AudioDataCopyToOptions) uint32]) {
	return fn.FromRef(
		bindings.AudioDataAllocationSizeFunc(
			this.Ref(),
		),
	)
}

// CopyTo calls the method "AudioData.copyTo".
//
// The returned bool will be false if there is no such method.
func (this AudioData) CopyTo(destination AllowSharedBufferSource, options AudioDataCopyToOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDataCopyTo(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyToFunc returns the method "AudioData.copyTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioData) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource, options AudioDataCopyToOptions)]) {
	return fn.FromRef(
		bindings.AudioDataCopyToFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "AudioData.clone".
//
// The returned bool will be false if there is no such method.
func (this AudioData) Clone() (AudioData, bool) {
	var _ok bool
	_ret := bindings.CallAudioDataClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return AudioData{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "AudioData.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioData) CloneFunc() (fn js.Func[func() AudioData]) {
	return fn.FromRef(
		bindings.AudioDataCloneFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioData.close".
//
// The returned bool will be false if there is no such method.
func (this AudioData) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDataClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "AudioData.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioData) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDataCloseFunc(
			this.Ref(),
		),
	)
}

type AudioDataOutputCallbackFunc func(this js.Ref, output AudioData) js.Ref

func (fn AudioDataOutputCallbackFunc) Register() js.Func[func(output AudioData)] {
	return js.RegisterCallback[func(output AudioData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AudioDataOutputCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AudioData{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioDataOutputCallback[T any] struct {
	Fn  func(arg T, this js.Ref, output AudioData) js.Ref
	Arg T
}

func (cb *AudioDataOutputCallback[T]) Register() js.Func[func(output AudioData)] {
	return js.RegisterCallback[func(output AudioData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AudioDataOutputCallback[T]) DispatchCallback(
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

		AudioData{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WebCodecsErrorCallbackFunc func(this js.Ref, err DOMException) js.Ref

func (fn WebCodecsErrorCallbackFunc) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WebCodecsErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WebCodecsErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err DOMException) js.Ref
	Arg T
}

func (cb *WebCodecsErrorCallback[T]) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WebCodecsErrorCallback[T]) DispatchCallback(
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

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioDecoderInit struct {
	// Output is "AudioDecoderInit.output"
	//
	// Required
	Output js.Func[func(output AudioData)]
	// Error is "AudioDecoderInit.error"
	//
	// Required
	Error js.Func[func(err DOMException)]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDecoderInit with all fields set.
func (p AudioDecoderInit) FromRef(ref js.Ref) AudioDecoderInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDecoderInit in the application heap.
func (p AudioDecoderInit) New() js.Ref {
	return bindings.AudioDecoderInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioDecoderInit) UpdateFrom(ref js.Ref) {
	bindings.AudioDecoderInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioDecoderInit) Update(ref js.Ref) {
	bindings.AudioDecoderInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioDecoderConfig struct {
	// Codec is "AudioDecoderConfig.codec"
	//
	// Required
	Codec js.String
	// SampleRate is "AudioDecoderConfig.sampleRate"
	//
	// Required
	SampleRate uint32
	// NumberOfChannels is "AudioDecoderConfig.numberOfChannels"
	//
	// Required
	NumberOfChannels uint32
	// Description is "AudioDecoderConfig.description"
	//
	// Optional
	Description BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDecoderConfig with all fields set.
func (p AudioDecoderConfig) FromRef(ref js.Ref) AudioDecoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDecoderConfig in the application heap.
func (p AudioDecoderConfig) New() js.Ref {
	return bindings.AudioDecoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioDecoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AudioDecoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioDecoderConfig) Update(ref js.Ref) {
	bindings.AudioDecoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type EncodedAudioChunkType uint32

const (
	_ EncodedAudioChunkType = iota

	EncodedAudioChunkType_KEY
	EncodedAudioChunkType_DELTA
)

func (EncodedAudioChunkType) FromRef(str js.Ref) EncodedAudioChunkType {
	return EncodedAudioChunkType(bindings.ConstOfEncodedAudioChunkType(str))
}

func (x EncodedAudioChunkType) String() (string, bool) {
	switch x {
	case EncodedAudioChunkType_KEY:
		return "key", true
	case EncodedAudioChunkType_DELTA:
		return "delta", true
	default:
		return "", false
	}
}

type EncodedAudioChunkInit struct {
	// Type is "EncodedAudioChunkInit.type"
	//
	// Required
	Type EncodedAudioChunkType
	// Timestamp is "EncodedAudioChunkInit.timestamp"
	//
	// Required
	Timestamp int64
	// Duration is "EncodedAudioChunkInit.duration"
	//
	// Optional
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration uint64
	// Data is "EncodedAudioChunkInit.data"
	//
	// Required
	Data BufferSource

	FFI_USE_Duration bool // for Duration.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EncodedAudioChunkInit with all fields set.
func (p EncodedAudioChunkInit) FromRef(ref js.Ref) EncodedAudioChunkInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EncodedAudioChunkInit in the application heap.
func (p EncodedAudioChunkInit) New() js.Ref {
	return bindings.EncodedAudioChunkInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EncodedAudioChunkInit) UpdateFrom(ref js.Ref) {
	bindings.EncodedAudioChunkInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EncodedAudioChunkInit) Update(ref js.Ref) {
	bindings.EncodedAudioChunkInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewEncodedAudioChunk(init EncodedAudioChunkInit) EncodedAudioChunk {
	return EncodedAudioChunk{}.FromRef(
		bindings.NewEncodedAudioChunkByEncodedAudioChunk(
			js.Pointer(&init)),
	)
}

type EncodedAudioChunk struct {
	ref js.Ref
}

func (this EncodedAudioChunk) Once() EncodedAudioChunk {
	this.Ref().Once()
	return this
}

func (this EncodedAudioChunk) Ref() js.Ref {
	return this.ref
}

func (this EncodedAudioChunk) FromRef(ref js.Ref) EncodedAudioChunk {
	this.ref = ref
	return this
}

func (this EncodedAudioChunk) Free() {
	this.Ref().Free()
}

// Type returns the value of property "EncodedAudioChunk.type".
//
// The returned bool will be false if there is no such property.
func (this EncodedAudioChunk) Type() (EncodedAudioChunkType, bool) {
	var _ok bool
	_ret := bindings.GetEncodedAudioChunkType(
		this.Ref(), js.Pointer(&_ok),
	)
	return EncodedAudioChunkType(_ret), _ok
}

// Timestamp returns the value of property "EncodedAudioChunk.timestamp".
//
// The returned bool will be false if there is no such property.
func (this EncodedAudioChunk) Timestamp() (int64, bool) {
	var _ok bool
	_ret := bindings.GetEncodedAudioChunkTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// Duration returns the value of property "EncodedAudioChunk.duration".
//
// The returned bool will be false if there is no such property.
func (this EncodedAudioChunk) Duration() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetEncodedAudioChunkDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ByteLength returns the value of property "EncodedAudioChunk.byteLength".
//
// The returned bool will be false if there is no such property.
func (this EncodedAudioChunk) ByteLength() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetEncodedAudioChunkByteLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CopyTo calls the method "EncodedAudioChunk.copyTo".
//
// The returned bool will be false if there is no such method.
func (this EncodedAudioChunk) CopyTo(destination AllowSharedBufferSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEncodedAudioChunkCopyTo(
		this.Ref(), js.Pointer(&_ok),
		destination.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyToFunc returns the method "EncodedAudioChunk.copyTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EncodedAudioChunk) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.EncodedAudioChunkCopyToFunc(
			this.Ref(),
		),
	)
}

type AudioDecoderSupport struct {
	// Supported is "AudioDecoderSupport.supported"
	//
	// Optional
	//
	// NOTE: FFI_USE_Supported MUST be set to true to make this field effective.
	Supported bool
	// Config is "AudioDecoderSupport.config"
	//
	// Optional
	Config AudioDecoderConfig

	FFI_USE_Supported bool // for Supported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioDecoderSupport with all fields set.
func (p AudioDecoderSupport) FromRef(ref js.Ref) AudioDecoderSupport {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioDecoderSupport in the application heap.
func (p AudioDecoderSupport) New() js.Ref {
	return bindings.AudioDecoderSupportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioDecoderSupport) UpdateFrom(ref js.Ref) {
	bindings.AudioDecoderSupportJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioDecoderSupport) Update(ref js.Ref) {
	bindings.AudioDecoderSupportJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CodecState uint32

const (
	_ CodecState = iota

	CodecState_UNCONFIGURED
	CodecState_CONFIGURED
	CodecState_CLOSED
)

func (CodecState) FromRef(str js.Ref) CodecState {
	return CodecState(bindings.ConstOfCodecState(str))
}

func (x CodecState) String() (string, bool) {
	switch x {
	case CodecState_UNCONFIGURED:
		return "unconfigured", true
	case CodecState_CONFIGURED:
		return "configured", true
	case CodecState_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

func NewAudioDecoder(init AudioDecoderInit) AudioDecoder {
	return AudioDecoder{}.FromRef(
		bindings.NewAudioDecoderByAudioDecoder(
			js.Pointer(&init)),
	)
}

type AudioDecoder struct {
	EventTarget
}

func (this AudioDecoder) Once() AudioDecoder {
	this.Ref().Once()
	return this
}

func (this AudioDecoder) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AudioDecoder) FromRef(ref js.Ref) AudioDecoder {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AudioDecoder) Free() {
	this.Ref().Free()
}

// State returns the value of property "AudioDecoder.state".
//
// The returned bool will be false if there is no such property.
func (this AudioDecoder) State() (CodecState, bool) {
	var _ok bool
	_ret := bindings.GetAudioDecoderState(
		this.Ref(), js.Pointer(&_ok),
	)
	return CodecState(_ret), _ok
}

// DecodeQueueSize returns the value of property "AudioDecoder.decodeQueueSize".
//
// The returned bool will be false if there is no such property.
func (this AudioDecoder) DecodeQueueSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioDecoderDecodeQueueSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Configure calls the method "AudioDecoder.configure".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) Configure(config AudioDecoderConfig) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderConfigure(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ConfigureFunc returns the method "AudioDecoder.configure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) ConfigureFunc() (fn js.Func[func(config AudioDecoderConfig)]) {
	return fn.FromRef(
		bindings.AudioDecoderConfigureFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "AudioDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) Decode(chunk EncodedAudioChunk) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderDecode(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DecodeFunc returns the method "AudioDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) DecodeFunc() (fn js.Func[func(chunk EncodedAudioChunk)]) {
	return fn.FromRef(
		bindings.AudioDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "AudioDecoder.flush".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) Flush() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderFlush(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// FlushFunc returns the method "AudioDecoder.flush".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) FlushFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioDecoderFlushFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "AudioDecoder.reset".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "AudioDecoder.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDecoderResetFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioDecoder.close".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "AudioDecoder.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDecoderCloseFunc(
			this.Ref(),
		),
	)
}

// IsConfigSupported calls the staticmethod "AudioDecoder.isConfigSupported".
//
// The returned bool will be false if there is no such method.
func (this AudioDecoder) IsConfigSupported(config AudioDecoderConfig) (js.Promise[AudioDecoderSupport], bool) {
	var _ok bool
	_ret := bindings.CallAudioDecoderIsConfigSupported(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	return js.Promise[AudioDecoderSupport]{}.FromRef(_ret), _ok
}

// IsConfigSupportedFunc returns the staticmethod "AudioDecoder.isConfigSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioDecoder) IsConfigSupportedFunc() (fn js.Func[func(config AudioDecoderConfig) js.Promise[AudioDecoderSupport]]) {
	return fn.FromRef(
		bindings.AudioDecoderIsConfigSupportedFunc(
			this.Ref(),
		),
	)
}

type EncodedAudioChunkOutputCallbackFunc func(this js.Ref, output EncodedAudioChunk, metadata EncodedAudioChunkMetadata) js.Ref

func (fn EncodedAudioChunkOutputCallbackFunc) Register() js.Func[func(output EncodedAudioChunk, metadata EncodedAudioChunkMetadata)] {
	return js.RegisterCallback[func(output EncodedAudioChunk, metadata EncodedAudioChunkMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EncodedAudioChunkOutputCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		EncodedAudioChunk{}.FromRef(args[0+1]),
		EncodedAudioChunkMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EncodedAudioChunkOutputCallback[T any] struct {
	Fn  func(arg T, this js.Ref, output EncodedAudioChunk, metadata EncodedAudioChunkMetadata) js.Ref
	Arg T
}

func (cb *EncodedAudioChunkOutputCallback[T]) Register() js.Func[func(output EncodedAudioChunk, metadata EncodedAudioChunkMetadata)] {
	return js.RegisterCallback[func(output EncodedAudioChunk, metadata EncodedAudioChunkMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EncodedAudioChunkOutputCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		EncodedAudioChunk{}.FromRef(args[0+1]),
		EncodedAudioChunkMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EncodedAudioChunkMetadata struct {
	// DecoderConfig is "EncodedAudioChunkMetadata.decoderConfig"
	//
	// Optional
	DecoderConfig AudioDecoderConfig

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EncodedAudioChunkMetadata with all fields set.
func (p EncodedAudioChunkMetadata) FromRef(ref js.Ref) EncodedAudioChunkMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EncodedAudioChunkMetadata in the application heap.
func (p EncodedAudioChunkMetadata) New() js.Ref {
	return bindings.EncodedAudioChunkMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EncodedAudioChunkMetadata) UpdateFrom(ref js.Ref) {
	bindings.EncodedAudioChunkMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EncodedAudioChunkMetadata) Update(ref js.Ref) {
	bindings.EncodedAudioChunkMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
