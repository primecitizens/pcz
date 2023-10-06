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
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackKind(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "MediaStreamTrack.id".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "MediaStreamTrack.label".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Enabled returns the value of property "MediaStreamTrack.enabled".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Enabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEnabled sets the value of property "MediaStreamTrack.enabled" to val.
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
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Muted() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackMuted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "MediaStreamTrack.readyState".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) ReadyState() (ret MediaStreamTrackState, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentHint returns the value of property "MediaStreamTrack.contentHint".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) ContentHint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackContentHint(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetContentHint sets the value of property "MediaStreamTrack.contentHint" to val.
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
// It returns ok=false if there is no such property.
func (this MediaStreamTrack) Isolated() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackIsolated(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClone returns true if the method "MediaStreamTrack.clone" exists.
func (this MediaStreamTrack) HasClone() bool {
	return js.True == bindings.HasMediaStreamTrackClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "MediaStreamTrack.clone".
func (this MediaStreamTrack) CloneFunc() (fn js.Func[func() MediaStreamTrack]) {
	return fn.FromRef(
		bindings.MediaStreamTrackCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "MediaStreamTrack.clone".
func (this MediaStreamTrack) Clone() (ret MediaStreamTrack) {
	bindings.CallMediaStreamTrackClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "MediaStreamTrack.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryClone() (ret MediaStreamTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "MediaStreamTrack.stop" exists.
func (this MediaStreamTrack) HasStop() bool {
	return js.True == bindings.HasMediaStreamTrackStop(
		this.Ref(),
	)
}

// StopFunc returns the method "MediaStreamTrack.stop".
func (this MediaStreamTrack) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaStreamTrackStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "MediaStreamTrack.stop".
func (this MediaStreamTrack) Stop() (ret js.Void) {
	bindings.CallMediaStreamTrackStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "MediaStreamTrack.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetCapabilities returns true if the method "MediaStreamTrack.getCapabilities" exists.
func (this MediaStreamTrack) HasGetCapabilities() bool {
	return js.True == bindings.HasMediaStreamTrackGetCapabilities(
		this.Ref(),
	)
}

// GetCapabilitiesFunc returns the method "MediaStreamTrack.getCapabilities".
func (this MediaStreamTrack) GetCapabilitiesFunc() (fn js.Func[func() MediaTrackCapabilities]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetCapabilitiesFunc(
			this.Ref(),
		),
	)
}

// GetCapabilities calls the method "MediaStreamTrack.getCapabilities".
func (this MediaStreamTrack) GetCapabilities() (ret MediaTrackCapabilities) {
	bindings.CallMediaStreamTrackGetCapabilities(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCapabilities calls the method "MediaStreamTrack.getCapabilities"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryGetCapabilities() (ret MediaTrackCapabilities, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackGetCapabilities(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetConstraints returns true if the method "MediaStreamTrack.getConstraints" exists.
func (this MediaStreamTrack) HasGetConstraints() bool {
	return js.True == bindings.HasMediaStreamTrackGetConstraints(
		this.Ref(),
	)
}

// GetConstraintsFunc returns the method "MediaStreamTrack.getConstraints".
func (this MediaStreamTrack) GetConstraintsFunc() (fn js.Func[func() MediaTrackConstraints]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetConstraintsFunc(
			this.Ref(),
		),
	)
}

// GetConstraints calls the method "MediaStreamTrack.getConstraints".
func (this MediaStreamTrack) GetConstraints() (ret MediaTrackConstraints) {
	bindings.CallMediaStreamTrackGetConstraints(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetConstraints calls the method "MediaStreamTrack.getConstraints"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryGetConstraints() (ret MediaTrackConstraints, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackGetConstraints(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSettings returns true if the method "MediaStreamTrack.getSettings" exists.
func (this MediaStreamTrack) HasGetSettings() bool {
	return js.True == bindings.HasMediaStreamTrackGetSettings(
		this.Ref(),
	)
}

// GetSettingsFunc returns the method "MediaStreamTrack.getSettings".
func (this MediaStreamTrack) GetSettingsFunc() (fn js.Func[func() MediaTrackSettings]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetSettingsFunc(
			this.Ref(),
		),
	)
}

// GetSettings calls the method "MediaStreamTrack.getSettings".
func (this MediaStreamTrack) GetSettings() (ret MediaTrackSettings) {
	bindings.CallMediaStreamTrackGetSettings(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSettings calls the method "MediaStreamTrack.getSettings"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryGetSettings() (ret MediaTrackSettings, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackGetSettings(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasApplyConstraints returns true if the method "MediaStreamTrack.applyConstraints" exists.
func (this MediaStreamTrack) HasApplyConstraints() bool {
	return js.True == bindings.HasMediaStreamTrackApplyConstraints(
		this.Ref(),
	)
}

// ApplyConstraintsFunc returns the method "MediaStreamTrack.applyConstraints".
func (this MediaStreamTrack) ApplyConstraintsFunc() (fn js.Func[func(constraints MediaTrackConstraints) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackApplyConstraintsFunc(
			this.Ref(),
		),
	)
}

// ApplyConstraints calls the method "MediaStreamTrack.applyConstraints".
func (this MediaStreamTrack) ApplyConstraints(constraints MediaTrackConstraints) (ret js.Promise[js.Void]) {
	bindings.CallMediaStreamTrackApplyConstraints(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&constraints),
	)

	return
}

// TryApplyConstraints calls the method "MediaStreamTrack.applyConstraints"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryApplyConstraints(constraints MediaTrackConstraints) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackApplyConstraints(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
	)

	return
}

// HasApplyConstraints1 returns true if the method "MediaStreamTrack.applyConstraints" exists.
func (this MediaStreamTrack) HasApplyConstraints1() bool {
	return js.True == bindings.HasMediaStreamTrackApplyConstraints1(
		this.Ref(),
	)
}

// ApplyConstraints1Func returns the method "MediaStreamTrack.applyConstraints".
func (this MediaStreamTrack) ApplyConstraints1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackApplyConstraints1Func(
			this.Ref(),
		),
	)
}

// ApplyConstraints1 calls the method "MediaStreamTrack.applyConstraints".
func (this MediaStreamTrack) ApplyConstraints1() (ret js.Promise[js.Void]) {
	bindings.CallMediaStreamTrackApplyConstraints1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryApplyConstraints1 calls the method "MediaStreamTrack.applyConstraints"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryApplyConstraints1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackApplyConstraints1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSupportedCaptureActions returns true if the method "MediaStreamTrack.getSupportedCaptureActions" exists.
func (this MediaStreamTrack) HasGetSupportedCaptureActions() bool {
	return js.True == bindings.HasMediaStreamTrackGetSupportedCaptureActions(
		this.Ref(),
	)
}

// GetSupportedCaptureActionsFunc returns the method "MediaStreamTrack.getSupportedCaptureActions".
func (this MediaStreamTrack) GetSupportedCaptureActionsFunc() (fn js.Func[func() js.Array[js.String]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetSupportedCaptureActionsFunc(
			this.Ref(),
		),
	)
}

// GetSupportedCaptureActions calls the method "MediaStreamTrack.getSupportedCaptureActions".
func (this MediaStreamTrack) GetSupportedCaptureActions() (ret js.Array[js.String]) {
	bindings.CallMediaStreamTrackGetSupportedCaptureActions(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedCaptureActions calls the method "MediaStreamTrack.getSupportedCaptureActions"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryGetSupportedCaptureActions() (ret js.Array[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackGetSupportedCaptureActions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSendCaptureAction returns true if the method "MediaStreamTrack.sendCaptureAction" exists.
func (this MediaStreamTrack) HasSendCaptureAction() bool {
	return js.True == bindings.HasMediaStreamTrackSendCaptureAction(
		this.Ref(),
	)
}

// SendCaptureActionFunc returns the method "MediaStreamTrack.sendCaptureAction".
func (this MediaStreamTrack) SendCaptureActionFunc() (fn js.Func[func(action CaptureAction) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.MediaStreamTrackSendCaptureActionFunc(
			this.Ref(),
		),
	)
}

// SendCaptureAction calls the method "MediaStreamTrack.sendCaptureAction".
func (this MediaStreamTrack) SendCaptureAction(action CaptureAction) (ret js.Promise[js.Void]) {
	bindings.CallMediaStreamTrackSendCaptureAction(
		this.Ref(), js.Pointer(&ret),
		uint32(action),
	)

	return
}

// TrySendCaptureAction calls the method "MediaStreamTrack.sendCaptureAction"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TrySendCaptureAction(action CaptureAction) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackSendCaptureAction(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(action),
	)

	return
}

// HasGetCaptureHandle returns true if the method "MediaStreamTrack.getCaptureHandle" exists.
func (this MediaStreamTrack) HasGetCaptureHandle() bool {
	return js.True == bindings.HasMediaStreamTrackGetCaptureHandle(
		this.Ref(),
	)
}

// GetCaptureHandleFunc returns the method "MediaStreamTrack.getCaptureHandle".
func (this MediaStreamTrack) GetCaptureHandleFunc() (fn js.Func[func() CaptureHandle]) {
	return fn.FromRef(
		bindings.MediaStreamTrackGetCaptureHandleFunc(
			this.Ref(),
		),
	)
}

// GetCaptureHandle calls the method "MediaStreamTrack.getCaptureHandle".
func (this MediaStreamTrack) GetCaptureHandle() (ret CaptureHandle) {
	bindings.CallMediaStreamTrackGetCaptureHandle(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCaptureHandle calls the method "MediaStreamTrack.getCaptureHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStreamTrack) TryGetCaptureHandle() (ret CaptureHandle, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamTrackGetCaptureHandle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewMediaStream(stream MediaStream) (ret MediaStream) {
	ret.ref = bindings.NewMediaStreamByMediaStream(
		stream.Ref())
	return
}

func NewMediaStreamByMediaStream1(tracks js.Array[MediaStreamTrack]) (ret MediaStream) {
	ret.ref = bindings.NewMediaStreamByMediaStream1(
		tracks.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MediaStream) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaStreamId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Active returns the value of property "MediaStream.active".
//
// It returns ok=false if there is no such property.
func (this MediaStream) Active() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaStreamActive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetAudioTracks returns true if the method "MediaStream.getAudioTracks" exists.
func (this MediaStream) HasGetAudioTracks() bool {
	return js.True == bindings.HasMediaStreamGetAudioTracks(
		this.Ref(),
	)
}

// GetAudioTracksFunc returns the method "MediaStream.getAudioTracks".
func (this MediaStream) GetAudioTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetAudioTracksFunc(
			this.Ref(),
		),
	)
}

// GetAudioTracks calls the method "MediaStream.getAudioTracks".
func (this MediaStream) GetAudioTracks() (ret js.Array[MediaStreamTrack]) {
	bindings.CallMediaStreamGetAudioTracks(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAudioTracks calls the method "MediaStream.getAudioTracks"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryGetAudioTracks() (ret js.Array[MediaStreamTrack], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamGetAudioTracks(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetVideoTracks returns true if the method "MediaStream.getVideoTracks" exists.
func (this MediaStream) HasGetVideoTracks() bool {
	return js.True == bindings.HasMediaStreamGetVideoTracks(
		this.Ref(),
	)
}

// GetVideoTracksFunc returns the method "MediaStream.getVideoTracks".
func (this MediaStream) GetVideoTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetVideoTracksFunc(
			this.Ref(),
		),
	)
}

// GetVideoTracks calls the method "MediaStream.getVideoTracks".
func (this MediaStream) GetVideoTracks() (ret js.Array[MediaStreamTrack]) {
	bindings.CallMediaStreamGetVideoTracks(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetVideoTracks calls the method "MediaStream.getVideoTracks"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryGetVideoTracks() (ret js.Array[MediaStreamTrack], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamGetVideoTracks(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetTracks returns true if the method "MediaStream.getTracks" exists.
func (this MediaStream) HasGetTracks() bool {
	return js.True == bindings.HasMediaStreamGetTracks(
		this.Ref(),
	)
}

// GetTracksFunc returns the method "MediaStream.getTracks".
func (this MediaStream) GetTracksFunc() (fn js.Func[func() js.Array[MediaStreamTrack]]) {
	return fn.FromRef(
		bindings.MediaStreamGetTracksFunc(
			this.Ref(),
		),
	)
}

// GetTracks calls the method "MediaStream.getTracks".
func (this MediaStream) GetTracks() (ret js.Array[MediaStreamTrack]) {
	bindings.CallMediaStreamGetTracks(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTracks calls the method "MediaStream.getTracks"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryGetTracks() (ret js.Array[MediaStreamTrack], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamGetTracks(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetTrackById returns true if the method "MediaStream.getTrackById" exists.
func (this MediaStream) HasGetTrackById() bool {
	return js.True == bindings.HasMediaStreamGetTrackById(
		this.Ref(),
	)
}

// GetTrackByIdFunc returns the method "MediaStream.getTrackById".
func (this MediaStream) GetTrackByIdFunc() (fn js.Func[func(trackId js.String) MediaStreamTrack]) {
	return fn.FromRef(
		bindings.MediaStreamGetTrackByIdFunc(
			this.Ref(),
		),
	)
}

// GetTrackById calls the method "MediaStream.getTrackById".
func (this MediaStream) GetTrackById(trackId js.String) (ret MediaStreamTrack) {
	bindings.CallMediaStreamGetTrackById(
		this.Ref(), js.Pointer(&ret),
		trackId.Ref(),
	)

	return
}

// TryGetTrackById calls the method "MediaStream.getTrackById"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryGetTrackById(trackId js.String) (ret MediaStreamTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamGetTrackById(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		trackId.Ref(),
	)

	return
}

// HasAddTrack returns true if the method "MediaStream.addTrack" exists.
func (this MediaStream) HasAddTrack() bool {
	return js.True == bindings.HasMediaStreamAddTrack(
		this.Ref(),
	)
}

// AddTrackFunc returns the method "MediaStream.addTrack".
func (this MediaStream) AddTrackFunc() (fn js.Func[func(track MediaStreamTrack)]) {
	return fn.FromRef(
		bindings.MediaStreamAddTrackFunc(
			this.Ref(),
		),
	)
}

// AddTrack calls the method "MediaStream.addTrack".
func (this MediaStream) AddTrack(track MediaStreamTrack) (ret js.Void) {
	bindings.CallMediaStreamAddTrack(
		this.Ref(), js.Pointer(&ret),
		track.Ref(),
	)

	return
}

// TryAddTrack calls the method "MediaStream.addTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryAddTrack(track MediaStreamTrack) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamAddTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		track.Ref(),
	)

	return
}

// HasRemoveTrack returns true if the method "MediaStream.removeTrack" exists.
func (this MediaStream) HasRemoveTrack() bool {
	return js.True == bindings.HasMediaStreamRemoveTrack(
		this.Ref(),
	)
}

// RemoveTrackFunc returns the method "MediaStream.removeTrack".
func (this MediaStream) RemoveTrackFunc() (fn js.Func[func(track MediaStreamTrack)]) {
	return fn.FromRef(
		bindings.MediaStreamRemoveTrackFunc(
			this.Ref(),
		),
	)
}

// RemoveTrack calls the method "MediaStream.removeTrack".
func (this MediaStream) RemoveTrack(track MediaStreamTrack) (ret js.Void) {
	bindings.CallMediaStreamRemoveTrack(
		this.Ref(), js.Pointer(&ret),
		track.Ref(),
	)

	return
}

// TryRemoveTrack calls the method "MediaStream.removeTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryRemoveTrack(track MediaStreamTrack) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamRemoveTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		track.Ref(),
	)

	return
}

// HasClone returns true if the method "MediaStream.clone" exists.
func (this MediaStream) HasClone() bool {
	return js.True == bindings.HasMediaStreamClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "MediaStream.clone".
func (this MediaStream) CloneFunc() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.MediaStreamCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "MediaStream.clone".
func (this MediaStream) Clone() (ret MediaStream) {
	bindings.CallMediaStreamClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "MediaStream.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaStream) TryClone() (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaStreamClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this MediaError) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetMediaErrorCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "MediaError.message".
//
// It returns ok=false if there is no such property.
func (this MediaError) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaErrorMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this SourceBufferList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSourceBufferListLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "SourceBufferList." exists.
func (this SourceBufferList) HasGet() bool {
	return js.True == bindings.HasSourceBufferListGet(
		this.Ref(),
	)
}

// GetFunc returns the method "SourceBufferList.".
func (this SourceBufferList) GetFunc() (fn js.Func[func(index uint32) SourceBuffer]) {
	return fn.FromRef(
		bindings.SourceBufferListGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "SourceBufferList.".
func (this SourceBufferList) Get(index uint32) (ret SourceBuffer) {
	bindings.CallSourceBufferListGet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "SourceBufferList."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this SourceBufferList) TryGet(index uint32) (ret SourceBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySourceBufferListGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
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
// It returns ok=false if there is no such property.
func (this MediaSource) Handle() (ret MediaSourceHandle, ok bool) {
	ok = js.True == bindings.GetMediaSourceHandle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SourceBuffers returns the value of property "MediaSource.sourceBuffers".
//
// It returns ok=false if there is no such property.
func (this MediaSource) SourceBuffers() (ret SourceBufferList, ok bool) {
	ok = js.True == bindings.GetMediaSourceSourceBuffers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActiveSourceBuffers returns the value of property "MediaSource.activeSourceBuffers".
//
// It returns ok=false if there is no such property.
func (this MediaSource) ActiveSourceBuffers() (ret SourceBufferList, ok bool) {
	ok = js.True == bindings.GetMediaSourceActiveSourceBuffers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "MediaSource.readyState".
//
// It returns ok=false if there is no such property.
func (this MediaSource) ReadyState() (ret ReadyState, ok bool) {
	ok = js.True == bindings.GetMediaSourceReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "MediaSource.duration".
//
// It returns ok=false if there is no such property.
func (this MediaSource) Duration() (ret float64, ok bool) {
	ok = js.True == bindings.GetMediaSourceDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDuration sets the value of property "MediaSource.duration" to val.
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
// It returns ok=false if there is no such property.
func (this MediaSource) CanConstructInDedicatedWorker() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaSourceCanConstructInDedicatedWorker(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAddSourceBuffer returns true if the method "MediaSource.addSourceBuffer" exists.
func (this MediaSource) HasAddSourceBuffer() bool {
	return js.True == bindings.HasMediaSourceAddSourceBuffer(
		this.Ref(),
	)
}

// AddSourceBufferFunc returns the method "MediaSource.addSourceBuffer".
func (this MediaSource) AddSourceBufferFunc() (fn js.Func[func(typ js.String) SourceBuffer]) {
	return fn.FromRef(
		bindings.MediaSourceAddSourceBufferFunc(
			this.Ref(),
		),
	)
}

// AddSourceBuffer calls the method "MediaSource.addSourceBuffer".
func (this MediaSource) AddSourceBuffer(typ js.String) (ret SourceBuffer) {
	bindings.CallMediaSourceAddSourceBuffer(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryAddSourceBuffer calls the method "MediaSource.addSourceBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryAddSourceBuffer(typ js.String) (ret SourceBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceAddSourceBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasRemoveSourceBuffer returns true if the method "MediaSource.removeSourceBuffer" exists.
func (this MediaSource) HasRemoveSourceBuffer() bool {
	return js.True == bindings.HasMediaSourceRemoveSourceBuffer(
		this.Ref(),
	)
}

// RemoveSourceBufferFunc returns the method "MediaSource.removeSourceBuffer".
func (this MediaSource) RemoveSourceBufferFunc() (fn js.Func[func(sourceBuffer SourceBuffer)]) {
	return fn.FromRef(
		bindings.MediaSourceRemoveSourceBufferFunc(
			this.Ref(),
		),
	)
}

// RemoveSourceBuffer calls the method "MediaSource.removeSourceBuffer".
func (this MediaSource) RemoveSourceBuffer(sourceBuffer SourceBuffer) (ret js.Void) {
	bindings.CallMediaSourceRemoveSourceBuffer(
		this.Ref(), js.Pointer(&ret),
		sourceBuffer.Ref(),
	)

	return
}

// TryRemoveSourceBuffer calls the method "MediaSource.removeSourceBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryRemoveSourceBuffer(sourceBuffer SourceBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceRemoveSourceBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sourceBuffer.Ref(),
	)

	return
}

// HasEndOfStream returns true if the method "MediaSource.endOfStream" exists.
func (this MediaSource) HasEndOfStream() bool {
	return js.True == bindings.HasMediaSourceEndOfStream(
		this.Ref(),
	)
}

// EndOfStreamFunc returns the method "MediaSource.endOfStream".
func (this MediaSource) EndOfStreamFunc() (fn js.Func[func(err EndOfStreamError)]) {
	return fn.FromRef(
		bindings.MediaSourceEndOfStreamFunc(
			this.Ref(),
		),
	)
}

// EndOfStream calls the method "MediaSource.endOfStream".
func (this MediaSource) EndOfStream(err EndOfStreamError) (ret js.Void) {
	bindings.CallMediaSourceEndOfStream(
		this.Ref(), js.Pointer(&ret),
		uint32(err),
	)

	return
}

// TryEndOfStream calls the method "MediaSource.endOfStream"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryEndOfStream(err EndOfStreamError) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceEndOfStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(err),
	)

	return
}

// HasEndOfStream1 returns true if the method "MediaSource.endOfStream" exists.
func (this MediaSource) HasEndOfStream1() bool {
	return js.True == bindings.HasMediaSourceEndOfStream1(
		this.Ref(),
	)
}

// EndOfStream1Func returns the method "MediaSource.endOfStream".
func (this MediaSource) EndOfStream1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSourceEndOfStream1Func(
			this.Ref(),
		),
	)
}

// EndOfStream1 calls the method "MediaSource.endOfStream".
func (this MediaSource) EndOfStream1() (ret js.Void) {
	bindings.CallMediaSourceEndOfStream1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEndOfStream1 calls the method "MediaSource.endOfStream"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryEndOfStream1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceEndOfStream1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetLiveSeekableRange returns true if the method "MediaSource.setLiveSeekableRange" exists.
func (this MediaSource) HasSetLiveSeekableRange() bool {
	return js.True == bindings.HasMediaSourceSetLiveSeekableRange(
		this.Ref(),
	)
}

// SetLiveSeekableRangeFunc returns the method "MediaSource.setLiveSeekableRange".
func (this MediaSource) SetLiveSeekableRangeFunc() (fn js.Func[func(start float64, end float64)]) {
	return fn.FromRef(
		bindings.MediaSourceSetLiveSeekableRangeFunc(
			this.Ref(),
		),
	)
}

// SetLiveSeekableRange calls the method "MediaSource.setLiveSeekableRange".
func (this MediaSource) SetLiveSeekableRange(start float64, end float64) (ret js.Void) {
	bindings.CallMediaSourceSetLiveSeekableRange(
		this.Ref(), js.Pointer(&ret),
		float64(start),
		float64(end),
	)

	return
}

// TrySetLiveSeekableRange calls the method "MediaSource.setLiveSeekableRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TrySetLiveSeekableRange(start float64, end float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceSetLiveSeekableRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(start),
		float64(end),
	)

	return
}

// HasClearLiveSeekableRange returns true if the method "MediaSource.clearLiveSeekableRange" exists.
func (this MediaSource) HasClearLiveSeekableRange() bool {
	return js.True == bindings.HasMediaSourceClearLiveSeekableRange(
		this.Ref(),
	)
}

// ClearLiveSeekableRangeFunc returns the method "MediaSource.clearLiveSeekableRange".
func (this MediaSource) ClearLiveSeekableRangeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaSourceClearLiveSeekableRangeFunc(
			this.Ref(),
		),
	)
}

// ClearLiveSeekableRange calls the method "MediaSource.clearLiveSeekableRange".
func (this MediaSource) ClearLiveSeekableRange() (ret js.Void) {
	bindings.CallMediaSourceClearLiveSeekableRange(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearLiveSeekableRange calls the method "MediaSource.clearLiveSeekableRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryClearLiveSeekableRange() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceClearLiveSeekableRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsTypeSupported returns true if the staticmethod "MediaSource.isTypeSupported" exists.
func (this MediaSource) HasIsTypeSupported() bool {
	return js.True == bindings.HasMediaSourceIsTypeSupported(
		this.Ref(),
	)
}

// IsTypeSupportedFunc returns the staticmethod "MediaSource.isTypeSupported".
func (this MediaSource) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.MediaSourceIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "MediaSource.isTypeSupported".
func (this MediaSource) IsTypeSupported(typ js.String) (ret bool) {
	bindings.CallMediaSourceIsTypeSupported(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryIsTypeSupported calls the staticmethod "MediaSource.isTypeSupported"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaSource) TryIsTypeSupported(typ js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaSourceIsTypeSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this RemotePlayback) State() (ret RemotePlaybackState, ok bool) {
	ok = js.True == bindings.GetRemotePlaybackState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasWatchAvailability returns true if the method "RemotePlayback.watchAvailability" exists.
func (this RemotePlayback) HasWatchAvailability() bool {
	return js.True == bindings.HasRemotePlaybackWatchAvailability(
		this.Ref(),
	)
}

// WatchAvailabilityFunc returns the method "RemotePlayback.watchAvailability".
func (this RemotePlayback) WatchAvailabilityFunc() (fn js.Func[func(callback js.Func[func(available bool)]) js.Promise[js.Number[int32]]]) {
	return fn.FromRef(
		bindings.RemotePlaybackWatchAvailabilityFunc(
			this.Ref(),
		),
	)
}

// WatchAvailability calls the method "RemotePlayback.watchAvailability".
func (this RemotePlayback) WatchAvailability(callback js.Func[func(available bool)]) (ret js.Promise[js.Number[int32]]) {
	bindings.CallRemotePlaybackWatchAvailability(
		this.Ref(), js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryWatchAvailability calls the method "RemotePlayback.watchAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RemotePlayback) TryWatchAvailability(callback js.Func[func(available bool)]) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemotePlaybackWatchAvailability(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasCancelWatchAvailability returns true if the method "RemotePlayback.cancelWatchAvailability" exists.
func (this RemotePlayback) HasCancelWatchAvailability() bool {
	return js.True == bindings.HasRemotePlaybackCancelWatchAvailability(
		this.Ref(),
	)
}

// CancelWatchAvailabilityFunc returns the method "RemotePlayback.cancelWatchAvailability".
func (this RemotePlayback) CancelWatchAvailabilityFunc() (fn js.Func[func(id int32) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackCancelWatchAvailabilityFunc(
			this.Ref(),
		),
	)
}

// CancelWatchAvailability calls the method "RemotePlayback.cancelWatchAvailability".
func (this RemotePlayback) CancelWatchAvailability(id int32) (ret js.Promise[js.Void]) {
	bindings.CallRemotePlaybackCancelWatchAvailability(
		this.Ref(), js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryCancelWatchAvailability calls the method "RemotePlayback.cancelWatchAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RemotePlayback) TryCancelWatchAvailability(id int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemotePlaybackCancelWatchAvailability(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasCancelWatchAvailability1 returns true if the method "RemotePlayback.cancelWatchAvailability" exists.
func (this RemotePlayback) HasCancelWatchAvailability1() bool {
	return js.True == bindings.HasRemotePlaybackCancelWatchAvailability1(
		this.Ref(),
	)
}

// CancelWatchAvailability1Func returns the method "RemotePlayback.cancelWatchAvailability".
func (this RemotePlayback) CancelWatchAvailability1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackCancelWatchAvailability1Func(
			this.Ref(),
		),
	)
}

// CancelWatchAvailability1 calls the method "RemotePlayback.cancelWatchAvailability".
func (this RemotePlayback) CancelWatchAvailability1() (ret js.Promise[js.Void]) {
	bindings.CallRemotePlaybackCancelWatchAvailability1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCancelWatchAvailability1 calls the method "RemotePlayback.cancelWatchAvailability"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RemotePlayback) TryCancelWatchAvailability1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemotePlaybackCancelWatchAvailability1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPrompt returns true if the method "RemotePlayback.prompt" exists.
func (this RemotePlayback) HasPrompt() bool {
	return js.True == bindings.HasRemotePlaybackPrompt(
		this.Ref(),
	)
}

// PromptFunc returns the method "RemotePlayback.prompt".
func (this RemotePlayback) PromptFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.RemotePlaybackPromptFunc(
			this.Ref(),
		),
	)
}

// Prompt calls the method "RemotePlayback.prompt".
func (this RemotePlayback) Prompt() (ret js.Promise[js.Void]) {
	bindings.CallRemotePlaybackPrompt(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPrompt calls the method "RemotePlayback.prompt"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this RemotePlayback) TryPrompt() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemotePlaybackPrompt(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Error() (ret MediaError, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementError(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Src returns the value of property "HTMLMediaElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLMediaElement.src" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) SrcObject() (ret MediaProvider, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementSrcObject(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrcObject sets the value of property "HTMLMediaElement.srcObject" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) CurrentSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementCurrentSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CrossOrigin returns the value of property "HTMLMediaElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLMediaElement.crossOrigin" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) NetworkState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementNetworkState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Preload returns the value of property "HTMLMediaElement.preload".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Preload() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementPreload(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPreload sets the value of property "HTMLMediaElement.preload" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Buffered() (ret TimeRanges, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementBuffered(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "HTMLMediaElement.readyState".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Seeking returns the value of property "HTMLMediaElement.seeking".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Seeking() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementSeeking(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentTime returns the value of property "HTMLMediaElement.currentTime".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) CurrentTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCurrentTime sets the value of property "HTMLMediaElement.currentTime" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Duration() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Paused returns the value of property "HTMLMediaElement.paused".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Paused() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementPaused(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultPlaybackRate returns the value of property "HTMLMediaElement.defaultPlaybackRate".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) DefaultPlaybackRate() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementDefaultPlaybackRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultPlaybackRate sets the value of property "HTMLMediaElement.defaultPlaybackRate" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) PlaybackRate() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementPlaybackRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaybackRate sets the value of property "HTMLMediaElement.playbackRate" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) PreservesPitch() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementPreservesPitch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPreservesPitch sets the value of property "HTMLMediaElement.preservesPitch" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Played() (ret TimeRanges, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementPlayed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Seekable returns the value of property "HTMLMediaElement.seekable".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Seekable() (ret TimeRanges, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementSeekable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ended returns the value of property "HTMLMediaElement.ended".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Ended() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementEnded(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Autoplay returns the value of property "HTMLMediaElement.autoplay".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Autoplay() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementAutoplay(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutoplay sets the value of property "HTMLMediaElement.autoplay" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Loop() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementLoop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLoop sets the value of property "HTMLMediaElement.loop" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Controls() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementControls(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetControls sets the value of property "HTMLMediaElement.controls" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Volume() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementVolume(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVolume sets the value of property "HTMLMediaElement.volume" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Muted() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementMuted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMuted sets the value of property "HTMLMediaElement.muted" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) DefaultMuted() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementDefaultMuted(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultMuted sets the value of property "HTMLMediaElement.defaultMuted" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) AudioTracks() (ret AudioTrackList, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementAudioTracks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VideoTracks returns the value of property "HTMLMediaElement.videoTracks".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) VideoTracks() (ret VideoTrackList, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementVideoTracks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TextTracks returns the value of property "HTMLMediaElement.textTracks".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) TextTracks() (ret TextTrackList, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementTextTracks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaKeys returns the value of property "HTMLMediaElement.mediaKeys".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) MediaKeys() (ret MediaKeys, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementMediaKeys(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Remote returns the value of property "HTMLMediaElement.remote".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) Remote() (ret RemotePlayback, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementRemote(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DisableRemotePlayback returns the value of property "HTMLMediaElement.disableRemotePlayback".
//
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) DisableRemotePlayback() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementDisableRemotePlayback(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisableRemotePlayback sets the value of property "HTMLMediaElement.disableRemotePlayback" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLMediaElement) SinkId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMediaElementSinkId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasLoad returns true if the method "HTMLMediaElement.load" exists.
func (this HTMLMediaElement) HasLoad() bool {
	return js.True == bindings.HasHTMLMediaElementLoad(
		this.Ref(),
	)
}

// LoadFunc returns the method "HTMLMediaElement.load".
func (this HTMLMediaElement) LoadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMediaElementLoadFunc(
			this.Ref(),
		),
	)
}

// Load calls the method "HTMLMediaElement.load".
func (this HTMLMediaElement) Load() (ret js.Void) {
	bindings.CallHTMLMediaElementLoad(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLoad calls the method "HTMLMediaElement.load"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryLoad() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementLoad(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCanPlayType returns true if the method "HTMLMediaElement.canPlayType" exists.
func (this HTMLMediaElement) HasCanPlayType() bool {
	return js.True == bindings.HasHTMLMediaElementCanPlayType(
		this.Ref(),
	)
}

// CanPlayTypeFunc returns the method "HTMLMediaElement.canPlayType".
func (this HTMLMediaElement) CanPlayTypeFunc() (fn js.Func[func(typ js.String) CanPlayTypeResult]) {
	return fn.FromRef(
		bindings.HTMLMediaElementCanPlayTypeFunc(
			this.Ref(),
		),
	)
}

// CanPlayType calls the method "HTMLMediaElement.canPlayType".
func (this HTMLMediaElement) CanPlayType(typ js.String) (ret CanPlayTypeResult) {
	bindings.CallHTMLMediaElementCanPlayType(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryCanPlayType calls the method "HTMLMediaElement.canPlayType"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryCanPlayType(typ js.String) (ret CanPlayTypeResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementCanPlayType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasFastSeek returns true if the method "HTMLMediaElement.fastSeek" exists.
func (this HTMLMediaElement) HasFastSeek() bool {
	return js.True == bindings.HasHTMLMediaElementFastSeek(
		this.Ref(),
	)
}

// FastSeekFunc returns the method "HTMLMediaElement.fastSeek".
func (this HTMLMediaElement) FastSeekFunc() (fn js.Func[func(time float64)]) {
	return fn.FromRef(
		bindings.HTMLMediaElementFastSeekFunc(
			this.Ref(),
		),
	)
}

// FastSeek calls the method "HTMLMediaElement.fastSeek".
func (this HTMLMediaElement) FastSeek(time float64) (ret js.Void) {
	bindings.CallHTMLMediaElementFastSeek(
		this.Ref(), js.Pointer(&ret),
		float64(time),
	)

	return
}

// TryFastSeek calls the method "HTMLMediaElement.fastSeek"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryFastSeek(time float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementFastSeek(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(time),
	)

	return
}

// HasGetStartDate returns true if the method "HTMLMediaElement.getStartDate" exists.
func (this HTMLMediaElement) HasGetStartDate() bool {
	return js.True == bindings.HasHTMLMediaElementGetStartDate(
		this.Ref(),
	)
}

// GetStartDateFunc returns the method "HTMLMediaElement.getStartDate".
func (this HTMLMediaElement) GetStartDateFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.HTMLMediaElementGetStartDateFunc(
			this.Ref(),
		),
	)
}

// GetStartDate calls the method "HTMLMediaElement.getStartDate".
func (this HTMLMediaElement) GetStartDate() (ret js.Object) {
	bindings.CallHTMLMediaElementGetStartDate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetStartDate calls the method "HTMLMediaElement.getStartDate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryGetStartDate() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementGetStartDate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPlay returns true if the method "HTMLMediaElement.play" exists.
func (this HTMLMediaElement) HasPlay() bool {
	return js.True == bindings.HasHTMLMediaElementPlay(
		this.Ref(),
	)
}

// PlayFunc returns the method "HTMLMediaElement.play".
func (this HTMLMediaElement) PlayFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementPlayFunc(
			this.Ref(),
		),
	)
}

// Play calls the method "HTMLMediaElement.play".
func (this HTMLMediaElement) Play() (ret js.Promise[js.Void]) {
	bindings.CallHTMLMediaElementPlay(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPlay calls the method "HTMLMediaElement.play"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryPlay() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementPlay(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPause returns true if the method "HTMLMediaElement.pause" exists.
func (this HTMLMediaElement) HasPause() bool {
	return js.True == bindings.HasHTMLMediaElementPause(
		this.Ref(),
	)
}

// PauseFunc returns the method "HTMLMediaElement.pause".
func (this HTMLMediaElement) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMediaElementPauseFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "HTMLMediaElement.pause".
func (this HTMLMediaElement) Pause() (ret js.Void) {
	bindings.CallHTMLMediaElementPause(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "HTMLMediaElement.pause"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementPause(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAddTextTrack returns true if the method "HTMLMediaElement.addTextTrack" exists.
func (this HTMLMediaElement) HasAddTextTrack() bool {
	return js.True == bindings.HasHTMLMediaElementAddTextTrack(
		this.Ref(),
	)
}

// AddTextTrackFunc returns the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrackFunc() (fn js.Func[func(kind TextTrackKind, label js.String, language js.String) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrackFunc(
			this.Ref(),
		),
	)
}

// AddTextTrack calls the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrack(kind TextTrackKind, label js.String, language js.String) (ret TextTrack) {
	bindings.CallHTMLMediaElementAddTextTrack(
		this.Ref(), js.Pointer(&ret),
		uint32(kind),
		label.Ref(),
		language.Ref(),
	)

	return
}

// TryAddTextTrack calls the method "HTMLMediaElement.addTextTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryAddTextTrack(kind TextTrackKind, label js.String, language js.String) (ret TextTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementAddTextTrack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(kind),
		label.Ref(),
		language.Ref(),
	)

	return
}

// HasAddTextTrack1 returns true if the method "HTMLMediaElement.addTextTrack" exists.
func (this HTMLMediaElement) HasAddTextTrack1() bool {
	return js.True == bindings.HasHTMLMediaElementAddTextTrack1(
		this.Ref(),
	)
}

// AddTextTrack1Func returns the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrack1Func() (fn js.Func[func(kind TextTrackKind, label js.String) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrack1Func(
			this.Ref(),
		),
	)
}

// AddTextTrack1 calls the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrack1(kind TextTrackKind, label js.String) (ret TextTrack) {
	bindings.CallHTMLMediaElementAddTextTrack1(
		this.Ref(), js.Pointer(&ret),
		uint32(kind),
		label.Ref(),
	)

	return
}

// TryAddTextTrack1 calls the method "HTMLMediaElement.addTextTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryAddTextTrack1(kind TextTrackKind, label js.String) (ret TextTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementAddTextTrack1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(kind),
		label.Ref(),
	)

	return
}

// HasAddTextTrack2 returns true if the method "HTMLMediaElement.addTextTrack" exists.
func (this HTMLMediaElement) HasAddTextTrack2() bool {
	return js.True == bindings.HasHTMLMediaElementAddTextTrack2(
		this.Ref(),
	)
}

// AddTextTrack2Func returns the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrack2Func() (fn js.Func[func(kind TextTrackKind) TextTrack]) {
	return fn.FromRef(
		bindings.HTMLMediaElementAddTextTrack2Func(
			this.Ref(),
		),
	)
}

// AddTextTrack2 calls the method "HTMLMediaElement.addTextTrack".
func (this HTMLMediaElement) AddTextTrack2(kind TextTrackKind) (ret TextTrack) {
	bindings.CallHTMLMediaElementAddTextTrack2(
		this.Ref(), js.Pointer(&ret),
		uint32(kind),
	)

	return
}

// TryAddTextTrack2 calls the method "HTMLMediaElement.addTextTrack"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryAddTextTrack2(kind TextTrackKind) (ret TextTrack, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementAddTextTrack2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(kind),
	)

	return
}

// HasSetMediaKeys returns true if the method "HTMLMediaElement.setMediaKeys" exists.
func (this HTMLMediaElement) HasSetMediaKeys() bool {
	return js.True == bindings.HasHTMLMediaElementSetMediaKeys(
		this.Ref(),
	)
}

// SetMediaKeysFunc returns the method "HTMLMediaElement.setMediaKeys".
func (this HTMLMediaElement) SetMediaKeysFunc() (fn js.Func[func(mediaKeys MediaKeys) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementSetMediaKeysFunc(
			this.Ref(),
		),
	)
}

// SetMediaKeys calls the method "HTMLMediaElement.setMediaKeys".
func (this HTMLMediaElement) SetMediaKeys(mediaKeys MediaKeys) (ret js.Promise[js.Void]) {
	bindings.CallHTMLMediaElementSetMediaKeys(
		this.Ref(), js.Pointer(&ret),
		mediaKeys.Ref(),
	)

	return
}

// TrySetMediaKeys calls the method "HTMLMediaElement.setMediaKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TrySetMediaKeys(mediaKeys MediaKeys) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementSetMediaKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mediaKeys.Ref(),
	)

	return
}

// HasCaptureStream returns true if the method "HTMLMediaElement.captureStream" exists.
func (this HTMLMediaElement) HasCaptureStream() bool {
	return js.True == bindings.HasHTMLMediaElementCaptureStream(
		this.Ref(),
	)
}

// CaptureStreamFunc returns the method "HTMLMediaElement.captureStream".
func (this HTMLMediaElement) CaptureStreamFunc() (fn js.Func[func() MediaStream]) {
	return fn.FromRef(
		bindings.HTMLMediaElementCaptureStreamFunc(
			this.Ref(),
		),
	)
}

// CaptureStream calls the method "HTMLMediaElement.captureStream".
func (this HTMLMediaElement) CaptureStream() (ret MediaStream) {
	bindings.CallHTMLMediaElementCaptureStream(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCaptureStream calls the method "HTMLMediaElement.captureStream"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TryCaptureStream() (ret MediaStream, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementCaptureStream(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetSinkId returns true if the method "HTMLMediaElement.setSinkId" exists.
func (this HTMLMediaElement) HasSetSinkId() bool {
	return js.True == bindings.HasHTMLMediaElementSetSinkId(
		this.Ref(),
	)
}

// SetSinkIdFunc returns the method "HTMLMediaElement.setSinkId".
func (this HTMLMediaElement) SetSinkIdFunc() (fn js.Func[func(sinkId js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLMediaElementSetSinkIdFunc(
			this.Ref(),
		),
	)
}

// SetSinkId calls the method "HTMLMediaElement.setSinkId".
func (this HTMLMediaElement) SetSinkId(sinkId js.String) (ret js.Promise[js.Void]) {
	bindings.CallHTMLMediaElementSetSinkId(
		this.Ref(), js.Pointer(&ret),
		sinkId.Ref(),
	)

	return
}

// TrySetSinkId calls the method "HTMLMediaElement.setSinkId"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLMediaElement) TrySetSinkId(sinkId js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMediaElementSetSinkId(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sinkId.Ref(),
	)

	return
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

func NewMediaElementAudioSourceNode(context AudioContext, options MediaElementAudioSourceOptions) (ret MediaElementAudioSourceNode) {
	ret.ref = bindings.NewMediaElementAudioSourceNodeByMediaElementAudioSourceNode(
		context.Ref(),
		js.Pointer(&options))
	return
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
// It returns ok=false if there is no such property.
func (this MediaElementAudioSourceNode) MediaElement() (ret HTMLMediaElement, ok bool) {
	ok = js.True == bindings.GetMediaElementAudioSourceNodeMediaElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewMediaStreamAudioSourceNode(context AudioContext, options MediaStreamAudioSourceOptions) (ret MediaStreamAudioSourceNode) {
	ret.ref = bindings.NewMediaStreamAudioSourceNodeByMediaStreamAudioSourceNode(
		context.Ref(),
		js.Pointer(&options))
	return
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
// It returns ok=false if there is no such property.
func (this MediaStreamAudioSourceNode) MediaStream() (ret MediaStream, ok bool) {
	ok = js.True == bindings.GetMediaStreamAudioSourceNodeMediaStream(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewMediaStreamTrackAudioSourceNode(context AudioContext, options MediaStreamTrackAudioSourceOptions) (ret MediaStreamTrackAudioSourceNode) {
	ret.ref = bindings.NewMediaStreamTrackAudioSourceNodeByMediaStreamTrackAudioSourceNode(
		context.Ref(),
		js.Pointer(&options))
	return
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

func NewMediaStreamAudioDestinationNode(context AudioContext, options AudioNodeOptions) (ret MediaStreamAudioDestinationNode) {
	ret.ref = bindings.NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode1(context AudioContext) (ret MediaStreamAudioDestinationNode) {
	ret.ref = bindings.NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode1(
		context.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MediaStreamAudioDestinationNode) Stream() (ret MediaStream, ok bool) {
	ok = js.True == bindings.GetMediaStreamAudioDestinationNodeStream(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this AudioSinkInfo) Type() (ret AudioSinkType, ok bool) {
	ok = js.True == bindings.GetAudioSinkInfoType(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasStart returns true if the method "AudioRenderCapacity.start" exists.
func (this AudioRenderCapacity) HasStart() bool {
	return js.True == bindings.HasAudioRenderCapacityStart(
		this.Ref(),
	)
}

// StartFunc returns the method "AudioRenderCapacity.start".
func (this AudioRenderCapacity) StartFunc() (fn js.Func[func(options AudioRenderCapacityOptions)]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "AudioRenderCapacity.start".
func (this AudioRenderCapacity) Start(options AudioRenderCapacityOptions) (ret js.Void) {
	bindings.CallAudioRenderCapacityStart(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryStart calls the method "AudioRenderCapacity.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioRenderCapacity) TryStart(options AudioRenderCapacityOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioRenderCapacityStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasStart1 returns true if the method "AudioRenderCapacity.start" exists.
func (this AudioRenderCapacity) HasStart1() bool {
	return js.True == bindings.HasAudioRenderCapacityStart1(
		this.Ref(),
	)
}

// Start1Func returns the method "AudioRenderCapacity.start".
func (this AudioRenderCapacity) Start1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStart1Func(
			this.Ref(),
		),
	)
}

// Start1 calls the method "AudioRenderCapacity.start".
func (this AudioRenderCapacity) Start1() (ret js.Void) {
	bindings.CallAudioRenderCapacityStart1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "AudioRenderCapacity.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioRenderCapacity) TryStart1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioRenderCapacityStart1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "AudioRenderCapacity.stop" exists.
func (this AudioRenderCapacity) HasStop() bool {
	return js.True == bindings.HasAudioRenderCapacityStop(
		this.Ref(),
	)
}

// StopFunc returns the method "AudioRenderCapacity.stop".
func (this AudioRenderCapacity) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioRenderCapacityStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "AudioRenderCapacity.stop".
func (this AudioRenderCapacity) Stop() (ret js.Void) {
	bindings.CallAudioRenderCapacityStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "AudioRenderCapacity.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioRenderCapacity) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioRenderCapacityStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewAudioContext(contextOptions AudioContextOptions) (ret AudioContext) {
	ret.ref = bindings.NewAudioContextByAudioContext(
		js.Pointer(&contextOptions))
	return
}

func NewAudioContextByAudioContext1() (ret AudioContext) {
	ret.ref = bindings.NewAudioContextByAudioContext1()
	return
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
// It returns ok=false if there is no such property.
func (this AudioContext) BaseLatency() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioContextBaseLatency(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OutputLatency returns the value of property "AudioContext.outputLatency".
//
// It returns ok=false if there is no such property.
func (this AudioContext) OutputLatency() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioContextOutputLatency(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SinkId returns the value of property "AudioContext.sinkId".
//
// It returns ok=false if there is no such property.
func (this AudioContext) SinkId() (ret OneOf_String_AudioSinkInfo, ok bool) {
	ok = js.True == bindings.GetAudioContextSinkId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RenderCapacity returns the value of property "AudioContext.renderCapacity".
//
// It returns ok=false if there is no such property.
func (this AudioContext) RenderCapacity() (ret AudioRenderCapacity, ok bool) {
	ok = js.True == bindings.GetAudioContextRenderCapacity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetOutputTimestamp returns true if the method "AudioContext.getOutputTimestamp" exists.
func (this AudioContext) HasGetOutputTimestamp() bool {
	return js.True == bindings.HasAudioContextGetOutputTimestamp(
		this.Ref(),
	)
}

// GetOutputTimestampFunc returns the method "AudioContext.getOutputTimestamp".
func (this AudioContext) GetOutputTimestampFunc() (fn js.Func[func() AudioTimestamp]) {
	return fn.FromRef(
		bindings.AudioContextGetOutputTimestampFunc(
			this.Ref(),
		),
	)
}

// GetOutputTimestamp calls the method "AudioContext.getOutputTimestamp".
func (this AudioContext) GetOutputTimestamp() (ret AudioTimestamp) {
	bindings.CallAudioContextGetOutputTimestamp(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetOutputTimestamp calls the method "AudioContext.getOutputTimestamp"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryGetOutputTimestamp() (ret AudioTimestamp, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextGetOutputTimestamp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResume returns true if the method "AudioContext.resume" exists.
func (this AudioContext) HasResume() bool {
	return js.True == bindings.HasAudioContextResume(
		this.Ref(),
	)
}

// ResumeFunc returns the method "AudioContext.resume".
func (this AudioContext) ResumeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextResumeFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "AudioContext.resume".
func (this AudioContext) Resume() (ret js.Promise[js.Void]) {
	bindings.CallAudioContextResume(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResume calls the method "AudioContext.resume"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryResume() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextResume(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSuspend returns true if the method "AudioContext.suspend" exists.
func (this AudioContext) HasSuspend() bool {
	return js.True == bindings.HasAudioContextSuspend(
		this.Ref(),
	)
}

// SuspendFunc returns the method "AudioContext.suspend".
func (this AudioContext) SuspendFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextSuspendFunc(
			this.Ref(),
		),
	)
}

// Suspend calls the method "AudioContext.suspend".
func (this AudioContext) Suspend() (ret js.Promise[js.Void]) {
	bindings.CallAudioContextSuspend(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySuspend calls the method "AudioContext.suspend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TrySuspend() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextSuspend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "AudioContext.close" exists.
func (this AudioContext) HasClose() bool {
	return js.True == bindings.HasAudioContextClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "AudioContext.close".
func (this AudioContext) CloseFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioContext.close".
func (this AudioContext) Close() (ret js.Promise[js.Void]) {
	bindings.CallAudioContextClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AudioContext.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryClose() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetSinkId returns true if the method "AudioContext.setSinkId" exists.
func (this AudioContext) HasSetSinkId() bool {
	return js.True == bindings.HasAudioContextSetSinkId(
		this.Ref(),
	)
}

// SetSinkIdFunc returns the method "AudioContext.setSinkId".
func (this AudioContext) SetSinkIdFunc() (fn js.Func[func(sinkId OneOf_String_AudioSinkOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioContextSetSinkIdFunc(
			this.Ref(),
		),
	)
}

// SetSinkId calls the method "AudioContext.setSinkId".
func (this AudioContext) SetSinkId(sinkId OneOf_String_AudioSinkOptions) (ret js.Promise[js.Void]) {
	bindings.CallAudioContextSetSinkId(
		this.Ref(), js.Pointer(&ret),
		sinkId.Ref(),
	)

	return
}

// TrySetSinkId calls the method "AudioContext.setSinkId"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TrySetSinkId(sinkId OneOf_String_AudioSinkOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextSetSinkId(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		sinkId.Ref(),
	)

	return
}

// HasCreateMediaElementSource returns true if the method "AudioContext.createMediaElementSource" exists.
func (this AudioContext) HasCreateMediaElementSource() bool {
	return js.True == bindings.HasAudioContextCreateMediaElementSource(
		this.Ref(),
	)
}

// CreateMediaElementSourceFunc returns the method "AudioContext.createMediaElementSource".
func (this AudioContext) CreateMediaElementSourceFunc() (fn js.Func[func(mediaElement HTMLMediaElement) MediaElementAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaElementSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaElementSource calls the method "AudioContext.createMediaElementSource".
func (this AudioContext) CreateMediaElementSource(mediaElement HTMLMediaElement) (ret MediaElementAudioSourceNode) {
	bindings.CallAudioContextCreateMediaElementSource(
		this.Ref(), js.Pointer(&ret),
		mediaElement.Ref(),
	)

	return
}

// TryCreateMediaElementSource calls the method "AudioContext.createMediaElementSource"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryCreateMediaElementSource(mediaElement HTMLMediaElement) (ret MediaElementAudioSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextCreateMediaElementSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mediaElement.Ref(),
	)

	return
}

// HasCreateMediaStreamSource returns true if the method "AudioContext.createMediaStreamSource" exists.
func (this AudioContext) HasCreateMediaStreamSource() bool {
	return js.True == bindings.HasAudioContextCreateMediaStreamSource(
		this.Ref(),
	)
}

// CreateMediaStreamSourceFunc returns the method "AudioContext.createMediaStreamSource".
func (this AudioContext) CreateMediaStreamSourceFunc() (fn js.Func[func(mediaStream MediaStream) MediaStreamAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamSource calls the method "AudioContext.createMediaStreamSource".
func (this AudioContext) CreateMediaStreamSource(mediaStream MediaStream) (ret MediaStreamAudioSourceNode) {
	bindings.CallAudioContextCreateMediaStreamSource(
		this.Ref(), js.Pointer(&ret),
		mediaStream.Ref(),
	)

	return
}

// TryCreateMediaStreamSource calls the method "AudioContext.createMediaStreamSource"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryCreateMediaStreamSource(mediaStream MediaStream) (ret MediaStreamAudioSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextCreateMediaStreamSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mediaStream.Ref(),
	)

	return
}

// HasCreateMediaStreamTrackSource returns true if the method "AudioContext.createMediaStreamTrackSource" exists.
func (this AudioContext) HasCreateMediaStreamTrackSource() bool {
	return js.True == bindings.HasAudioContextCreateMediaStreamTrackSource(
		this.Ref(),
	)
}

// CreateMediaStreamTrackSourceFunc returns the method "AudioContext.createMediaStreamTrackSource".
func (this AudioContext) CreateMediaStreamTrackSourceFunc() (fn js.Func[func(mediaStreamTrack MediaStreamTrack) MediaStreamTrackAudioSourceNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamTrackSourceFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamTrackSource calls the method "AudioContext.createMediaStreamTrackSource".
func (this AudioContext) CreateMediaStreamTrackSource(mediaStreamTrack MediaStreamTrack) (ret MediaStreamTrackAudioSourceNode) {
	bindings.CallAudioContextCreateMediaStreamTrackSource(
		this.Ref(), js.Pointer(&ret),
		mediaStreamTrack.Ref(),
	)

	return
}

// TryCreateMediaStreamTrackSource calls the method "AudioContext.createMediaStreamTrackSource"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryCreateMediaStreamTrackSource(mediaStreamTrack MediaStreamTrack) (ret MediaStreamTrackAudioSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextCreateMediaStreamTrackSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mediaStreamTrack.Ref(),
	)

	return
}

// HasCreateMediaStreamDestination returns true if the method "AudioContext.createMediaStreamDestination" exists.
func (this AudioContext) HasCreateMediaStreamDestination() bool {
	return js.True == bindings.HasAudioContextCreateMediaStreamDestination(
		this.Ref(),
	)
}

// CreateMediaStreamDestinationFunc returns the method "AudioContext.createMediaStreamDestination".
func (this AudioContext) CreateMediaStreamDestinationFunc() (fn js.Func[func() MediaStreamAudioDestinationNode]) {
	return fn.FromRef(
		bindings.AudioContextCreateMediaStreamDestinationFunc(
			this.Ref(),
		),
	)
}

// CreateMediaStreamDestination calls the method "AudioContext.createMediaStreamDestination".
func (this AudioContext) CreateMediaStreamDestination() (ret MediaStreamAudioDestinationNode) {
	bindings.CallAudioContextCreateMediaStreamDestination(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateMediaStreamDestination calls the method "AudioContext.createMediaStreamDestination"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioContext) TryCreateMediaStreamDestination() (ret MediaStreamAudioDestinationNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioContextCreateMediaStreamDestination(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewAudioData(init AudioDataInit) (ret AudioData) {
	ret.ref = bindings.NewAudioDataByAudioData(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this AudioData) Format() (ret AudioSampleFormat, ok bool) {
	ok = js.True == bindings.GetAudioDataFormat(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SampleRate returns the value of property "AudioData.sampleRate".
//
// It returns ok=false if there is no such property.
func (this AudioData) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioDataSampleRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfFrames returns the value of property "AudioData.numberOfFrames".
//
// It returns ok=false if there is no such property.
func (this AudioData) NumberOfFrames() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioDataNumberOfFrames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfChannels returns the value of property "AudioData.numberOfChannels".
//
// It returns ok=false if there is no such property.
func (this AudioData) NumberOfChannels() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioDataNumberOfChannels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "AudioData.duration".
//
// It returns ok=false if there is no such property.
func (this AudioData) Duration() (ret uint64, ok bool) {
	ok = js.True == bindings.GetAudioDataDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "AudioData.timestamp".
//
// It returns ok=false if there is no such property.
func (this AudioData) Timestamp() (ret int64, ok bool) {
	ok = js.True == bindings.GetAudioDataTimestamp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAllocationSize returns true if the method "AudioData.allocationSize" exists.
func (this AudioData) HasAllocationSize() bool {
	return js.True == bindings.HasAudioDataAllocationSize(
		this.Ref(),
	)
}

// AllocationSizeFunc returns the method "AudioData.allocationSize".
func (this AudioData) AllocationSizeFunc() (fn js.Func[func(options AudioDataCopyToOptions) uint32]) {
	return fn.FromRef(
		bindings.AudioDataAllocationSizeFunc(
			this.Ref(),
		),
	)
}

// AllocationSize calls the method "AudioData.allocationSize".
func (this AudioData) AllocationSize(options AudioDataCopyToOptions) (ret uint32) {
	bindings.CallAudioDataAllocationSize(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAllocationSize calls the method "AudioData.allocationSize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioData) TryAllocationSize(options AudioDataCopyToOptions) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDataAllocationSize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCopyTo returns true if the method "AudioData.copyTo" exists.
func (this AudioData) HasCopyTo() bool {
	return js.True == bindings.HasAudioDataCopyTo(
		this.Ref(),
	)
}

// CopyToFunc returns the method "AudioData.copyTo".
func (this AudioData) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource, options AudioDataCopyToOptions)]) {
	return fn.FromRef(
		bindings.AudioDataCopyToFunc(
			this.Ref(),
		),
	)
}

// CopyTo calls the method "AudioData.copyTo".
func (this AudioData) CopyTo(destination AllowSharedBufferSource, options AudioDataCopyToOptions) (ret js.Void) {
	bindings.CallAudioDataCopyTo(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCopyTo calls the method "AudioData.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioData) TryCopyTo(destination AllowSharedBufferSource, options AudioDataCopyToOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDataCopyTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasClone returns true if the method "AudioData.clone" exists.
func (this AudioData) HasClone() bool {
	return js.True == bindings.HasAudioDataClone(
		this.Ref(),
	)
}

// CloneFunc returns the method "AudioData.clone".
func (this AudioData) CloneFunc() (fn js.Func[func() AudioData]) {
	return fn.FromRef(
		bindings.AudioDataCloneFunc(
			this.Ref(),
		),
	)
}

// Clone calls the method "AudioData.clone".
func (this AudioData) Clone() (ret AudioData) {
	bindings.CallAudioDataClone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "AudioData.clone"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioData) TryClone() (ret AudioData, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDataClone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "AudioData.close" exists.
func (this AudioData) HasClose() bool {
	return js.True == bindings.HasAudioDataClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "AudioData.close".
func (this AudioData) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDataCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioData.close".
func (this AudioData) Close() (ret js.Void) {
	bindings.CallAudioDataClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AudioData.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioData) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDataClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewEncodedAudioChunk(init EncodedAudioChunkInit) (ret EncodedAudioChunk) {
	ret.ref = bindings.NewEncodedAudioChunkByEncodedAudioChunk(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this EncodedAudioChunk) Type() (ret EncodedAudioChunkType, ok bool) {
	ok = js.True == bindings.GetEncodedAudioChunkType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "EncodedAudioChunk.timestamp".
//
// It returns ok=false if there is no such property.
func (this EncodedAudioChunk) Timestamp() (ret int64, ok bool) {
	ok = js.True == bindings.GetEncodedAudioChunkTimestamp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "EncodedAudioChunk.duration".
//
// It returns ok=false if there is no such property.
func (this EncodedAudioChunk) Duration() (ret uint64, ok bool) {
	ok = js.True == bindings.GetEncodedAudioChunkDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ByteLength returns the value of property "EncodedAudioChunk.byteLength".
//
// It returns ok=false if there is no such property.
func (this EncodedAudioChunk) ByteLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEncodedAudioChunkByteLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCopyTo returns true if the method "EncodedAudioChunk.copyTo" exists.
func (this EncodedAudioChunk) HasCopyTo() bool {
	return js.True == bindings.HasEncodedAudioChunkCopyTo(
		this.Ref(),
	)
}

// CopyToFunc returns the method "EncodedAudioChunk.copyTo".
func (this EncodedAudioChunk) CopyToFunc() (fn js.Func[func(destination AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.EncodedAudioChunkCopyToFunc(
			this.Ref(),
		),
	)
}

// CopyTo calls the method "EncodedAudioChunk.copyTo".
func (this EncodedAudioChunk) CopyTo(destination AllowSharedBufferSource) (ret js.Void) {
	bindings.CallEncodedAudioChunkCopyTo(
		this.Ref(), js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryCopyTo calls the method "EncodedAudioChunk.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this EncodedAudioChunk) TryCopyTo(destination AllowSharedBufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEncodedAudioChunkCopyTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
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
	//
	// NOTE: Config.FFI_USE MUST be set to true to get Config used.
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

func NewAudioDecoder(init AudioDecoderInit) (ret AudioDecoder) {
	ret.ref = bindings.NewAudioDecoderByAudioDecoder(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this AudioDecoder) State() (ret CodecState, ok bool) {
	ok = js.True == bindings.GetAudioDecoderState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DecodeQueueSize returns the value of property "AudioDecoder.decodeQueueSize".
//
// It returns ok=false if there is no such property.
func (this AudioDecoder) DecodeQueueSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioDecoderDecodeQueueSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasConfigure returns true if the method "AudioDecoder.configure" exists.
func (this AudioDecoder) HasConfigure() bool {
	return js.True == bindings.HasAudioDecoderConfigure(
		this.Ref(),
	)
}

// ConfigureFunc returns the method "AudioDecoder.configure".
func (this AudioDecoder) ConfigureFunc() (fn js.Func[func(config AudioDecoderConfig)]) {
	return fn.FromRef(
		bindings.AudioDecoderConfigureFunc(
			this.Ref(),
		),
	)
}

// Configure calls the method "AudioDecoder.configure".
func (this AudioDecoder) Configure(config AudioDecoderConfig) (ret js.Void) {
	bindings.CallAudioDecoderConfigure(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryConfigure calls the method "AudioDecoder.configure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryConfigure(config AudioDecoderConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderConfigure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasDecode returns true if the method "AudioDecoder.decode" exists.
func (this AudioDecoder) HasDecode() bool {
	return js.True == bindings.HasAudioDecoderDecode(
		this.Ref(),
	)
}

// DecodeFunc returns the method "AudioDecoder.decode".
func (this AudioDecoder) DecodeFunc() (fn js.Func[func(chunk EncodedAudioChunk)]) {
	return fn.FromRef(
		bindings.AudioDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "AudioDecoder.decode".
func (this AudioDecoder) Decode(chunk EncodedAudioChunk) (ret js.Void) {
	bindings.CallAudioDecoderDecode(
		this.Ref(), js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryDecode calls the method "AudioDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryDecode(chunk EncodedAudioChunk) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderDecode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasFlush returns true if the method "AudioDecoder.flush" exists.
func (this AudioDecoder) HasFlush() bool {
	return js.True == bindings.HasAudioDecoderFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "AudioDecoder.flush".
func (this AudioDecoder) FlushFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioDecoderFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "AudioDecoder.flush".
func (this AudioDecoder) Flush() (ret js.Promise[js.Void]) {
	bindings.CallAudioDecoderFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "AudioDecoder.flush"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryFlush() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "AudioDecoder.reset" exists.
func (this AudioDecoder) HasReset() bool {
	return js.True == bindings.HasAudioDecoderReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "AudioDecoder.reset".
func (this AudioDecoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDecoderResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "AudioDecoder.reset".
func (this AudioDecoder) Reset() (ret js.Void) {
	bindings.CallAudioDecoderReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "AudioDecoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "AudioDecoder.close" exists.
func (this AudioDecoder) HasClose() bool {
	return js.True == bindings.HasAudioDecoderClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "AudioDecoder.close".
func (this AudioDecoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioDecoderCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioDecoder.close".
func (this AudioDecoder) Close() (ret js.Void) {
	bindings.CallAudioDecoderClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AudioDecoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsConfigSupported returns true if the staticmethod "AudioDecoder.isConfigSupported" exists.
func (this AudioDecoder) HasIsConfigSupported() bool {
	return js.True == bindings.HasAudioDecoderIsConfigSupported(
		this.Ref(),
	)
}

// IsConfigSupportedFunc returns the staticmethod "AudioDecoder.isConfigSupported".
func (this AudioDecoder) IsConfigSupportedFunc() (fn js.Func[func(config AudioDecoderConfig) js.Promise[AudioDecoderSupport]]) {
	return fn.FromRef(
		bindings.AudioDecoderIsConfigSupportedFunc(
			this.Ref(),
		),
	)
}

// IsConfigSupported calls the staticmethod "AudioDecoder.isConfigSupported".
func (this AudioDecoder) IsConfigSupported(config AudioDecoderConfig) (ret js.Promise[AudioDecoderSupport]) {
	bindings.CallAudioDecoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryIsConfigSupported calls the staticmethod "AudioDecoder.isConfigSupported"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioDecoder) TryIsConfigSupported(config AudioDecoderConfig) (ret js.Promise[AudioDecoderSupport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioDecoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
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
	//
	// NOTE: DecoderConfig.FFI_USE MUST be set to true to get DecoderConfig used.
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
