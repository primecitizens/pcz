// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web get_MediaStreamTrack_Kind
//go:noescape

func GetMediaStreamTrackKind(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Id
//go:noescape

func GetMediaStreamTrackId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Label
//go:noescape

func GetMediaStreamTrackLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Enabled
//go:noescape

func GetMediaStreamTrackEnabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaStreamTrack_Enabled
//go:noescape

func SetMediaStreamTrackEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Muted
//go:noescape

func GetMediaStreamTrackMuted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_ReadyState
//go:noescape

func GetMediaStreamTrackReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaStreamTrack_ContentHint
//go:noescape

func GetMediaStreamTrackContentHint(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaStreamTrack_ContentHint
//go:noescape

func SetMediaStreamTrackContentHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Isolated
//go:noescape

func GetMediaStreamTrackIsolated(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_Clone
//go:noescape

func CallMediaStreamTrackClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_Clone
//go:noescape

func MediaStreamTrackCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_Stop
//go:noescape

func CallMediaStreamTrackStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_Stop
//go:noescape

func MediaStreamTrackStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_GetCapabilities
//go:noescape

func CallMediaStreamTrackGetCapabilities(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetCapabilities
//go:noescape

func MediaStreamTrackGetCapabilitiesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_GetConstraints
//go:noescape

func CallMediaStreamTrackGetConstraints(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetConstraints
//go:noescape

func MediaStreamTrackGetConstraintsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_GetSettings
//go:noescape

func CallMediaStreamTrackGetSettings(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetSettings
//go:noescape

func MediaStreamTrackGetSettingsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_ApplyConstraints
//go:noescape

func CallMediaStreamTrackApplyConstraints(
	this js.Ref,
	ptr unsafe.Pointer,

	constraints unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_ApplyConstraints
//go:noescape

func MediaStreamTrackApplyConstraintsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_ApplyConstraints1
//go:noescape

func CallMediaStreamTrackApplyConstraints1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_ApplyConstraints1
//go:noescape

func MediaStreamTrackApplyConstraints1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape

func CallMediaStreamTrackGetSupportedCaptureActions(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape

func MediaStreamTrackGetSupportedCaptureActionsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_SendCaptureAction
//go:noescape

func CallMediaStreamTrackSendCaptureAction(
	this js.Ref,
	ptr unsafe.Pointer,

	action uint32,
) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_SendCaptureAction
//go:noescape

func MediaStreamTrackSendCaptureActionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStreamTrack_GetCaptureHandle
//go:noescape

func CallMediaStreamTrackGetCaptureHandle(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetCaptureHandle
//go:noescape

func MediaStreamTrackGetCaptureHandleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStream_MediaStream
//go:noescape

func NewMediaStreamByMediaStream(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStream_MediaStream1
//go:noescape

func NewMediaStreamByMediaStream1(
	tracks js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaStream_Id
//go:noescape

func GetMediaStreamId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaStream_Active
//go:noescape

func GetMediaStreamActive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MediaStream_GetAudioTracks
//go:noescape

func CallMediaStreamGetAudioTracks(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetAudioTracks
//go:noescape

func MediaStreamGetAudioTracksFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_GetVideoTracks
//go:noescape

func CallMediaStreamGetVideoTracks(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetVideoTracks
//go:noescape

func MediaStreamGetVideoTracksFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_GetTracks
//go:noescape

func CallMediaStreamGetTracks(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetTracks
//go:noescape

func MediaStreamGetTracksFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_GetTrackById
//go:noescape

func CallMediaStreamGetTrackById(
	this js.Ref,
	ptr unsafe.Pointer,

	trackId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetTrackById
//go:noescape

func MediaStreamGetTrackByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_AddTrack
//go:noescape

func CallMediaStreamAddTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	track js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaStream_AddTrack
//go:noescape

func MediaStreamAddTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_RemoveTrack
//go:noescape

func CallMediaStreamRemoveTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	track js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaStream_RemoveTrack
//go:noescape

func MediaStreamRemoveTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaStream_Clone
//go:noescape

func CallMediaStreamClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaStream_Clone
//go:noescape

func MediaStreamCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaError_Code
//go:noescape

func GetMediaErrorCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaError_Message
//go:noescape

func GetMediaErrorMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_EndOfStreamError
//go:noescape

func ConstOfEndOfStreamError(str js.Ref) uint32

//go:wasmimport plat/js/web get_SourceBufferList_Length
//go:noescape

func GetSourceBufferListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SourceBufferList_Get
//go:noescape

func CallSourceBufferListGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SourceBufferList_Get
//go:noescape

func SourceBufferListGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ReadyState
//go:noescape

func ConstOfReadyState(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaSource_Handle
//go:noescape

func GetMediaSourceHandle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaSource_SourceBuffers
//go:noescape

func GetMediaSourceSourceBuffers(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaSource_ActiveSourceBuffers
//go:noescape

func GetMediaSourceActiveSourceBuffers(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaSource_ReadyState
//go:noescape

func GetMediaSourceReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_MediaSource_Duration
//go:noescape

func GetMediaSourceDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_MediaSource_Duration
//go:noescape

func SetMediaSourceDuration(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_MediaSource_CanConstructInDedicatedWorker
//go:noescape

func GetMediaSourceCanConstructInDedicatedWorker(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MediaSource_AddSourceBuffer
//go:noescape

func CallMediaSourceAddSourceBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSource_AddSourceBuffer
//go:noescape

func MediaSourceAddSourceBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_RemoveSourceBuffer
//go:noescape

func CallMediaSourceRemoveSourceBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	sourceBuffer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSource_RemoveSourceBuffer
//go:noescape

func MediaSourceRemoveSourceBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_EndOfStream
//go:noescape

func CallMediaSourceEndOfStream(
	this js.Ref,
	ptr unsafe.Pointer,

	err uint32,
) js.Ref

//go:wasmimport plat/js/web func_MediaSource_EndOfStream
//go:noescape

func MediaSourceEndOfStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_EndOfStream1
//go:noescape

func CallMediaSourceEndOfStream1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaSource_EndOfStream1
//go:noescape

func MediaSourceEndOfStream1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_SetLiveSeekableRange
//go:noescape

func CallMediaSourceSetLiveSeekableRange(
	this js.Ref,
	ptr unsafe.Pointer,

	start float64,
	end float64,
) js.Ref

//go:wasmimport plat/js/web func_MediaSource_SetLiveSeekableRange
//go:noescape

func MediaSourceSetLiveSeekableRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_ClearLiveSeekableRange
//go:noescape

func CallMediaSourceClearLiveSeekableRange(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaSource_ClearLiveSeekableRange
//go:noescape

func MediaSourceClearLiveSeekableRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSource_IsTypeSupported
//go:noescape

func CallMediaSourceIsTypeSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSource_IsTypeSupported
//go:noescape

func MediaSourceIsTypeSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RemotePlaybackState
//go:noescape

func ConstOfRemotePlaybackState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RemotePlayback_State
//go:noescape

func GetRemotePlaybackState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_RemotePlayback_WatchAvailability
//go:noescape

func CallRemotePlaybackWatchAvailability(
	this js.Ref,
	ptr unsafe.Pointer,

	callback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_WatchAvailability
//go:noescape

func RemotePlaybackWatchAvailabilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RemotePlayback_CancelWatchAvailability
//go:noescape

func CallRemotePlaybackCancelWatchAvailability(
	this js.Ref,
	ptr unsafe.Pointer,

	id int32,
) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_CancelWatchAvailability
//go:noescape

func RemotePlaybackCancelWatchAvailabilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RemotePlayback_CancelWatchAvailability1
//go:noescape

func CallRemotePlaybackCancelWatchAvailability1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_CancelWatchAvailability1
//go:noescape

func RemotePlaybackCancelWatchAvailability1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RemotePlayback_Prompt
//go:noescape

func CallRemotePlaybackPrompt(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_Prompt
//go:noescape

func RemotePlaybackPromptFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Error
//go:noescape

func GetHTMLMediaElementError(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Src
//go:noescape

func GetHTMLMediaElementSrc(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Src
//go:noescape

func SetHTMLMediaElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_SrcObject
//go:noescape

func GetHTMLMediaElementSrcObject(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_SrcObject
//go:noescape

func SetHTMLMediaElementSrcObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_CurrentSrc
//go:noescape

func GetHTMLMediaElementCurrentSrc(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_CrossOrigin
//go:noescape

func GetHTMLMediaElementCrossOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_CrossOrigin
//go:noescape

func SetHTMLMediaElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_NetworkState
//go:noescape

func GetHTMLMediaElementNetworkState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_HTMLMediaElement_Preload
//go:noescape

func GetHTMLMediaElementPreload(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Preload
//go:noescape

func SetHTMLMediaElementPreload(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Buffered
//go:noescape

func GetHTMLMediaElementBuffered(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_ReadyState
//go:noescape

func GetHTMLMediaElementReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_HTMLMediaElement_Seeking
//go:noescape

func GetHTMLMediaElementSeeking(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_CurrentTime
//go:noescape

func GetHTMLMediaElementCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_HTMLMediaElement_CurrentTime
//go:noescape

func SetHTMLMediaElementCurrentTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Duration
//go:noescape

func GetHTMLMediaElementDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_HTMLMediaElement_Paused
//go:noescape

func GetHTMLMediaElementPaused(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_DefaultPlaybackRate
//go:noescape

func GetHTMLMediaElementDefaultPlaybackRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_HTMLMediaElement_DefaultPlaybackRate
//go:noescape

func SetHTMLMediaElementDefaultPlaybackRate(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_PlaybackRate
//go:noescape

func GetHTMLMediaElementPlaybackRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_HTMLMediaElement_PlaybackRate
//go:noescape

func SetHTMLMediaElementPlaybackRate(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_PreservesPitch
//go:noescape

func GetHTMLMediaElementPreservesPitch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_PreservesPitch
//go:noescape

func SetHTMLMediaElementPreservesPitch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Played
//go:noescape

func GetHTMLMediaElementPlayed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Seekable
//go:noescape

func GetHTMLMediaElementSeekable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Ended
//go:noescape

func GetHTMLMediaElementEnded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Autoplay
//go:noescape

func GetHTMLMediaElementAutoplay(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Autoplay
//go:noescape

func SetHTMLMediaElementAutoplay(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Loop
//go:noescape

func GetHTMLMediaElementLoop(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Loop
//go:noescape

func SetHTMLMediaElementLoop(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Controls
//go:noescape

func GetHTMLMediaElementControls(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Controls
//go:noescape

func SetHTMLMediaElementControls(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Volume
//go:noescape

func GetHTMLMediaElementVolume(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_HTMLMediaElement_Volume
//go:noescape

func SetHTMLMediaElementVolume(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Muted
//go:noescape

func GetHTMLMediaElementMuted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_Muted
//go:noescape

func SetHTMLMediaElementMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_DefaultMuted
//go:noescape

func GetHTMLMediaElementDefaultMuted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_DefaultMuted
//go:noescape

func SetHTMLMediaElementDefaultMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_AudioTracks
//go:noescape

func GetHTMLMediaElementAudioTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_VideoTracks
//go:noescape

func GetHTMLMediaElementVideoTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_TextTracks
//go:noescape

func GetHTMLMediaElementTextTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_MediaKeys
//go:noescape

func GetHTMLMediaElementMediaKeys(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Remote
//go:noescape

func GetHTMLMediaElementRemote(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_DisableRemotePlayback
//go:noescape

func GetHTMLMediaElementDisableRemotePlayback(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLMediaElement_DisableRemotePlayback
//go:noescape

func SetHTMLMediaElementDisableRemotePlayback(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_SinkId
//go:noescape

func GetHTMLMediaElementSinkId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_Load
//go:noescape

func CallHTMLMediaElementLoad(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Load
//go:noescape

func HTMLMediaElementLoadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_CanPlayType
//go:noescape

func CallHTMLMediaElementCanPlayType(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) uint32

//go:wasmimport plat/js/web func_HTMLMediaElement_CanPlayType
//go:noescape

func HTMLMediaElementCanPlayTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_FastSeek
//go:noescape

func CallHTMLMediaElementFastSeek(
	this js.Ref,
	ptr unsafe.Pointer,

	time float64,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_FastSeek
//go:noescape

func HTMLMediaElementFastSeekFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_GetStartDate
//go:noescape

func CallHTMLMediaElementGetStartDate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_GetStartDate
//go:noescape

func HTMLMediaElementGetStartDateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_Play
//go:noescape

func CallHTMLMediaElementPlay(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Play
//go:noescape

func HTMLMediaElementPlayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_Pause
//go:noescape

func CallHTMLMediaElementPause(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Pause
//go:noescape

func HTMLMediaElementPauseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack
//go:noescape

func CallHTMLMediaElementAddTextTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	kind uint32,
	label js.Ref,
	language js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack
//go:noescape

func HTMLMediaElementAddTextTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack1
//go:noescape

func CallHTMLMediaElementAddTextTrack1(
	this js.Ref,
	ptr unsafe.Pointer,

	kind uint32,
	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack1
//go:noescape

func HTMLMediaElementAddTextTrack1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack2
//go:noescape

func CallHTMLMediaElementAddTextTrack2(
	this js.Ref,
	ptr unsafe.Pointer,

	kind uint32,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack2
//go:noescape

func HTMLMediaElementAddTextTrack2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_SetMediaKeys
//go:noescape

func CallHTMLMediaElementSetMediaKeys(
	this js.Ref,
	ptr unsafe.Pointer,

	mediaKeys js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_SetMediaKeys
//go:noescape

func HTMLMediaElementSetMediaKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_CaptureStream
//go:noescape

func CallHTMLMediaElementCaptureStream(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_CaptureStream
//go:noescape

func HTMLMediaElementCaptureStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMediaElement_SetSinkId
//go:noescape

func CallHTMLMediaElementSetSinkId(
	this js.Ref,
	ptr unsafe.Pointer,

	sinkId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_SetSinkId
//go:noescape

func HTMLMediaElementSetSinkIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaElementAudioSourceOptions
//go:noescape

func MediaElementAudioSourceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaElementAudioSourceOptions
//go:noescape

func MediaElementAudioSourceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaElementAudioSourceNode_MediaElementAudioSourceNode
//go:noescape

func NewMediaElementAudioSourceNodeByMediaElementAudioSourceNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_MediaElementAudioSourceNode_MediaElement
//go:noescape

func GetMediaElementAudioSourceNodeMediaElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_MediaStreamAudioSourceOptions
//go:noescape

func MediaStreamAudioSourceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaStreamAudioSourceOptions
//go:noescape

func MediaStreamAudioSourceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStreamAudioSourceNode_MediaStreamAudioSourceNode
//go:noescape

func NewMediaStreamAudioSourceNodeByMediaStreamAudioSourceNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_MediaStreamAudioSourceNode_MediaStream
//go:noescape

func GetMediaStreamAudioSourceNodeMediaStream(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_MediaStreamTrackAudioSourceOptions
//go:noescape

func MediaStreamTrackAudioSourceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaStreamTrackAudioSourceOptions
//go:noescape

func MediaStreamTrackAudioSourceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStreamTrackAudioSourceNode_MediaStreamTrackAudioSourceNode
//go:noescape

func NewMediaStreamTrackAudioSourceNodeByMediaStreamTrackAudioSourceNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web store_AudioNodeOptions
//go:noescape

func AudioNodeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioNodeOptions
//go:noescape

func AudioNodeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaStreamAudioDestinationNode_MediaStreamAudioDestinationNode
//go:noescape

func NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode(
	context js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MediaStreamAudioDestinationNode_MediaStreamAudioDestinationNode1
//go:noescape

func NewMediaStreamAudioDestinationNodeByMediaStreamAudioDestinationNode1(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaStreamAudioDestinationNode_Stream
//go:noescape

func GetMediaStreamAudioDestinationNodeStream(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioSinkInfo_Type
//go:noescape

func GetAudioSinkInfoType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web store_AudioRenderCapacityOptions
//go:noescape

func AudioRenderCapacityOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioRenderCapacityOptions
//go:noescape

func AudioRenderCapacityOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioRenderCapacity_Start
//go:noescape

func CallAudioRenderCapacityStart(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Start
//go:noescape

func AudioRenderCapacityStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioRenderCapacity_Start1
//go:noescape

func CallAudioRenderCapacityStart1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Start1
//go:noescape

func AudioRenderCapacityStart1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioRenderCapacity_Stop
//go:noescape

func CallAudioRenderCapacityStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Stop
//go:noescape

func AudioRenderCapacityStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioContext_AudioContext
//go:noescape

func NewAudioContextByAudioContext(
	contextOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AudioContext_AudioContext1
//go:noescape

func NewAudioContextByAudioContext1() js.Ref

//go:wasmimport plat/js/web get_AudioContext_BaseLatency
//go:noescape

func GetAudioContextBaseLatency(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_AudioContext_OutputLatency
//go:noescape

func GetAudioContextOutputLatency(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_AudioContext_SinkId
//go:noescape

func GetAudioContextSinkId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioContext_RenderCapacity
//go:noescape

func GetAudioContextRenderCapacity(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AudioContext_GetOutputTimestamp
//go:noescape

func CallAudioContextGetOutputTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioContext_GetOutputTimestamp
//go:noescape

func AudioContextGetOutputTimestampFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_Resume
//go:noescape

func CallAudioContextResume(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Resume
//go:noescape

func AudioContextResumeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_Suspend
//go:noescape

func CallAudioContextSuspend(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Suspend
//go:noescape

func AudioContextSuspendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_Close
//go:noescape

func CallAudioContextClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Close
//go:noescape

func AudioContextCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_SetSinkId
//go:noescape

func CallAudioContextSetSinkId(
	this js.Ref,
	ptr unsafe.Pointer,

	sinkId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioContext_SetSinkId
//go:noescape

func AudioContextSetSinkIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_CreateMediaElementSource
//go:noescape

func CallAudioContextCreateMediaElementSource(
	this js.Ref,
	ptr unsafe.Pointer,

	mediaElement js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaElementSource
//go:noescape

func AudioContextCreateMediaElementSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamSource
//go:noescape

func CallAudioContextCreateMediaStreamSource(
	this js.Ref,
	ptr unsafe.Pointer,

	mediaStream js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamSource
//go:noescape

func AudioContextCreateMediaStreamSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamTrackSource
//go:noescape

func CallAudioContextCreateMediaStreamTrackSource(
	this js.Ref,
	ptr unsafe.Pointer,

	mediaStreamTrack js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamTrackSource
//go:noescape

func AudioContextCreateMediaStreamTrackSourceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamDestination
//go:noescape

func CallAudioContextCreateMediaStreamDestination(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamDestination
//go:noescape

func AudioContextCreateMediaStreamDestinationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AudioSampleFormat
//go:noescape

func ConstOfAudioSampleFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_AudioDataInit
//go:noescape

func AudioDataInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioDataInit
//go:noescape

func AudioDataInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioDataCopyToOptions
//go:noescape

func AudioDataCopyToOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioDataCopyToOptions
//go:noescape

func AudioDataCopyToOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AudioData_AudioData
//go:noescape

func NewAudioDataByAudioData(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_AudioData_Format
//go:noescape

func GetAudioDataFormat(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AudioData_SampleRate
//go:noescape

func GetAudioDataSampleRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_AudioData_NumberOfFrames
//go:noescape

func GetAudioDataNumberOfFrames(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AudioData_NumberOfChannels
//go:noescape

func GetAudioDataNumberOfChannels(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AudioData_Duration
//go:noescape

func GetAudioDataDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_AudioData_Timestamp
//go:noescape

func GetAudioDataTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web call_AudioData_AllocationSize
//go:noescape

func CallAudioDataAllocationSize(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web func_AudioData_AllocationSize
//go:noescape

func AudioDataAllocationSizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioData_CopyTo
//go:noescape

func CallAudioDataCopyTo(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_AudioData_CopyTo
//go:noescape

func AudioDataCopyToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioData_Clone
//go:noescape

func CallAudioDataClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioData_Clone
//go:noescape

func AudioDataCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioData_Close
//go:noescape

func CallAudioDataClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioData_Close
//go:noescape

func AudioDataCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioDecoderInit
//go:noescape

func AudioDecoderInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioDecoderInit
//go:noescape

func AudioDecoderInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioDecoderConfig
//go:noescape

func AudioDecoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioDecoderConfig
//go:noescape

func AudioDecoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_EncodedAudioChunkType
//go:noescape

func ConstOfEncodedAudioChunkType(str js.Ref) uint32

//go:wasmimport plat/js/web store_EncodedAudioChunkInit
//go:noescape

func EncodedAudioChunkInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EncodedAudioChunkInit
//go:noescape

func EncodedAudioChunkInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_EncodedAudioChunk_EncodedAudioChunk
//go:noescape

func NewEncodedAudioChunkByEncodedAudioChunk(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_EncodedAudioChunk_Type
//go:noescape

func GetEncodedAudioChunkType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_EncodedAudioChunk_Timestamp
//go:noescape

func GetEncodedAudioChunkTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web get_EncodedAudioChunk_Duration
//go:noescape

func GetEncodedAudioChunkDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_EncodedAudioChunk_ByteLength
//go:noescape

func GetEncodedAudioChunkByteLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_EncodedAudioChunk_CopyTo
//go:noescape

func CallEncodedAudioChunkCopyTo(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EncodedAudioChunk_CopyTo
//go:noescape

func EncodedAudioChunkCopyToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AudioDecoderSupport
//go:noescape

func AudioDecoderSupportJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioDecoderSupport
//go:noescape

func AudioDecoderSupportJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_CodecState
//go:noescape

func ConstOfCodecState(str js.Ref) uint32

//go:wasmimport plat/js/web new_AudioDecoder_AudioDecoder
//go:noescape

func NewAudioDecoderByAudioDecoder(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_AudioDecoder_State
//go:noescape

func GetAudioDecoderState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AudioDecoder_DecodeQueueSize
//go:noescape

func GetAudioDecoderDecodeQueueSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_AudioDecoder_Configure
//go:noescape

func CallAudioDecoderConfigure(
	this js.Ref,
	ptr unsafe.Pointer,

	config unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Configure
//go:noescape

func AudioDecoderConfigureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioDecoder_Decode
//go:noescape

func CallAudioDecoderDecode(
	this js.Ref,
	ptr unsafe.Pointer,

	chunk js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Decode
//go:noescape

func AudioDecoderDecodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioDecoder_Flush
//go:noescape

func CallAudioDecoderFlush(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Flush
//go:noescape

func AudioDecoderFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioDecoder_Reset
//go:noescape

func CallAudioDecoderReset(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Reset
//go:noescape

func AudioDecoderResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioDecoder_Close
//go:noescape

func CallAudioDecoderClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Close
//go:noescape

func AudioDecoderCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioDecoder_IsConfigSupported
//go:noescape

func CallAudioDecoderIsConfigSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	config unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_IsConfigSupported
//go:noescape

func AudioDecoderIsConfigSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_EncodedAudioChunkMetadata
//go:noescape

func EncodedAudioChunkMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EncodedAudioChunkMetadata
//go:noescape

func EncodedAudioChunkMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
