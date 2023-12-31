// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/web get_MediaStreamTrack_Kind
//go:noescape
func GetMediaStreamTrackKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStreamTrack_Id
//go:noescape
func GetMediaStreamTrackId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStreamTrack_Label
//go:noescape
func GetMediaStreamTrackLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStreamTrack_Enabled
//go:noescape
func GetMediaStreamTrackEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaStreamTrack_Enabled
//go:noescape
func SetMediaStreamTrackEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Muted
//go:noescape
func GetMediaStreamTrackMuted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStreamTrack_ReadyState
//go:noescape
func GetMediaStreamTrackReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStreamTrack_ContentHint
//go:noescape
func GetMediaStreamTrackContentHint(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaStreamTrack_ContentHint
//go:noescape
func SetMediaStreamTrackContentHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaStreamTrack_Isolated
//go:noescape
func GetMediaStreamTrackIsolated(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_Clone
//go:noescape
func HasFuncMediaStreamTrackClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_Clone
//go:noescape
func FuncMediaStreamTrackClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_Clone
//go:noescape
func CallMediaStreamTrackClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_Clone
//go:noescape
func TryMediaStreamTrackClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_Stop
//go:noescape
func HasFuncMediaStreamTrackStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_Stop
//go:noescape
func FuncMediaStreamTrackStop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_Stop
//go:noescape
func CallMediaStreamTrackStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_Stop
//go:noescape
func TryMediaStreamTrackStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_GetCapabilities
//go:noescape
func HasFuncMediaStreamTrackGetCapabilities(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetCapabilities
//go:noescape
func FuncMediaStreamTrackGetCapabilities(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_GetCapabilities
//go:noescape
func CallMediaStreamTrackGetCapabilities(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_GetCapabilities
//go:noescape
func TryMediaStreamTrackGetCapabilities(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_GetConstraints
//go:noescape
func HasFuncMediaStreamTrackGetConstraints(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetConstraints
//go:noescape
func FuncMediaStreamTrackGetConstraints(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_GetConstraints
//go:noescape
func CallMediaStreamTrackGetConstraints(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_GetConstraints
//go:noescape
func TryMediaStreamTrackGetConstraints(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_GetSettings
//go:noescape
func HasFuncMediaStreamTrackGetSettings(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetSettings
//go:noescape
func FuncMediaStreamTrackGetSettings(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_GetSettings
//go:noescape
func CallMediaStreamTrackGetSettings(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_GetSettings
//go:noescape
func TryMediaStreamTrackGetSettings(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_ApplyConstraints
//go:noescape
func HasFuncMediaStreamTrackApplyConstraints(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_ApplyConstraints
//go:noescape
func FuncMediaStreamTrackApplyConstraints(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_ApplyConstraints
//go:noescape
func CallMediaStreamTrackApplyConstraints(
	this js.Ref, retPtr unsafe.Pointer,
	constraints unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_ApplyConstraints
//go:noescape
func TryMediaStreamTrackApplyConstraints(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constraints unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_ApplyConstraints1
//go:noescape
func HasFuncMediaStreamTrackApplyConstraints1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_ApplyConstraints1
//go:noescape
func FuncMediaStreamTrackApplyConstraints1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_ApplyConstraints1
//go:noescape
func CallMediaStreamTrackApplyConstraints1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_ApplyConstraints1
//go:noescape
func TryMediaStreamTrackApplyConstraints1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape
func HasFuncMediaStreamTrackGetSupportedCaptureActions(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape
func FuncMediaStreamTrackGetSupportedCaptureActions(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape
func CallMediaStreamTrackGetSupportedCaptureActions(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_GetSupportedCaptureActions
//go:noescape
func TryMediaStreamTrackGetSupportedCaptureActions(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_SendCaptureAction
//go:noescape
func HasFuncMediaStreamTrackSendCaptureAction(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_SendCaptureAction
//go:noescape
func FuncMediaStreamTrackSendCaptureAction(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_SendCaptureAction
//go:noescape
func CallMediaStreamTrackSendCaptureAction(
	this js.Ref, retPtr unsafe.Pointer,
	action uint32)

//go:wasmimport plat/js/web try_MediaStreamTrack_SendCaptureAction
//go:noescape
func TryMediaStreamTrackSendCaptureAction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	action uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStreamTrack_GetCaptureHandle
//go:noescape
func HasFuncMediaStreamTrackGetCaptureHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStreamTrack_GetCaptureHandle
//go:noescape
func FuncMediaStreamTrackGetCaptureHandle(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStreamTrack_GetCaptureHandle
//go:noescape
func CallMediaStreamTrackGetCaptureHandle(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStreamTrack_GetCaptureHandle
//go:noescape
func TryMediaStreamTrackGetCaptureHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaStream_Active
//go:noescape
func GetMediaStreamActive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_GetAudioTracks
//go:noescape
func HasFuncMediaStreamGetAudioTracks(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetAudioTracks
//go:noescape
func FuncMediaStreamGetAudioTracks(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_GetAudioTracks
//go:noescape
func CallMediaStreamGetAudioTracks(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStream_GetAudioTracks
//go:noescape
func TryMediaStreamGetAudioTracks(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_GetVideoTracks
//go:noescape
func HasFuncMediaStreamGetVideoTracks(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetVideoTracks
//go:noescape
func FuncMediaStreamGetVideoTracks(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_GetVideoTracks
//go:noescape
func CallMediaStreamGetVideoTracks(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStream_GetVideoTracks
//go:noescape
func TryMediaStreamGetVideoTracks(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_GetTracks
//go:noescape
func HasFuncMediaStreamGetTracks(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetTracks
//go:noescape
func FuncMediaStreamGetTracks(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_GetTracks
//go:noescape
func CallMediaStreamGetTracks(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStream_GetTracks
//go:noescape
func TryMediaStreamGetTracks(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_GetTrackById
//go:noescape
func HasFuncMediaStreamGetTrackById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_GetTrackById
//go:noescape
func FuncMediaStreamGetTrackById(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_GetTrackById
//go:noescape
func CallMediaStreamGetTrackById(
	this js.Ref, retPtr unsafe.Pointer,
	trackId js.Ref)

//go:wasmimport plat/js/web try_MediaStream_GetTrackById
//go:noescape
func TryMediaStreamGetTrackById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	trackId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_AddTrack
//go:noescape
func HasFuncMediaStreamAddTrack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_AddTrack
//go:noescape
func FuncMediaStreamAddTrack(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_AddTrack
//go:noescape
func CallMediaStreamAddTrack(
	this js.Ref, retPtr unsafe.Pointer,
	track js.Ref)

//go:wasmimport plat/js/web try_MediaStream_AddTrack
//go:noescape
func TryMediaStreamAddTrack(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	track js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_RemoveTrack
//go:noescape
func HasFuncMediaStreamRemoveTrack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_RemoveTrack
//go:noescape
func FuncMediaStreamRemoveTrack(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_RemoveTrack
//go:noescape
func CallMediaStreamRemoveTrack(
	this js.Ref, retPtr unsafe.Pointer,
	track js.Ref)

//go:wasmimport plat/js/web try_MediaStream_RemoveTrack
//go:noescape
func TryMediaStreamRemoveTrack(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	track js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaStream_Clone
//go:noescape
func HasFuncMediaStreamClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaStream_Clone
//go:noescape
func FuncMediaStreamClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaStream_Clone
//go:noescape
func CallMediaStreamClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaStream_Clone
//go:noescape
func TryMediaStreamClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaError_Code
//go:noescape
func GetMediaErrorCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaError_Message
//go:noescape
func GetMediaErrorMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_EndOfStreamError
//go:noescape
func ConstOfEndOfStreamError(str js.Ref) uint32

//go:wasmimport plat/js/web get_SourceBufferList_Length
//go:noescape
func GetSourceBufferListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SourceBufferList_Get
//go:noescape
func HasFuncSourceBufferListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SourceBufferList_Get
//go:noescape
func FuncSourceBufferListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SourceBufferList_Get
//go:noescape
func CallSourceBufferListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SourceBufferList_Get
//go:noescape
func TrySourceBufferListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web constof_ReadyState
//go:noescape
func ConstOfReadyState(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaSource_Handle
//go:noescape
func GetMediaSourceHandle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaSource_SourceBuffers
//go:noescape
func GetMediaSourceSourceBuffers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaSource_ActiveSourceBuffers
//go:noescape
func GetMediaSourceActiveSourceBuffers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaSource_ReadyState
//go:noescape
func GetMediaSourceReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaSource_Duration
//go:noescape
func GetMediaSourceDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaSource_Duration
//go:noescape
func SetMediaSourceDuration(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_MediaSource_CanConstructInDedicatedWorker
//go:noescape
func GetMediaSourceCanConstructInDedicatedWorker(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_AddSourceBuffer
//go:noescape
func HasFuncMediaSourceAddSourceBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_AddSourceBuffer
//go:noescape
func FuncMediaSourceAddSourceBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_AddSourceBuffer
//go:noescape
func CallMediaSourceAddSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_MediaSource_AddSourceBuffer
//go:noescape
func TryMediaSourceAddSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_RemoveSourceBuffer
//go:noescape
func HasFuncMediaSourceRemoveSourceBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_RemoveSourceBuffer
//go:noescape
func FuncMediaSourceRemoveSourceBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_RemoveSourceBuffer
//go:noescape
func CallMediaSourceRemoveSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	sourceBuffer js.Ref)

//go:wasmimport plat/js/web try_MediaSource_RemoveSourceBuffer
//go:noescape
func TryMediaSourceRemoveSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sourceBuffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_EndOfStream
//go:noescape
func HasFuncMediaSourceEndOfStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_EndOfStream
//go:noescape
func FuncMediaSourceEndOfStream(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_EndOfStream
//go:noescape
func CallMediaSourceEndOfStream(
	this js.Ref, retPtr unsafe.Pointer,
	err uint32)

//go:wasmimport plat/js/web try_MediaSource_EndOfStream
//go:noescape
func TryMediaSourceEndOfStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_EndOfStream1
//go:noescape
func HasFuncMediaSourceEndOfStream1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_EndOfStream1
//go:noescape
func FuncMediaSourceEndOfStream1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_EndOfStream1
//go:noescape
func CallMediaSourceEndOfStream1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaSource_EndOfStream1
//go:noescape
func TryMediaSourceEndOfStream1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_SetLiveSeekableRange
//go:noescape
func HasFuncMediaSourceSetLiveSeekableRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_SetLiveSeekableRange
//go:noescape
func FuncMediaSourceSetLiveSeekableRange(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_SetLiveSeekableRange
//go:noescape
func CallMediaSourceSetLiveSeekableRange(
	this js.Ref, retPtr unsafe.Pointer,
	start float64,
	end float64)

//go:wasmimport plat/js/web try_MediaSource_SetLiveSeekableRange
//go:noescape
func TryMediaSourceSetLiveSeekableRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start float64,
	end float64) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_ClearLiveSeekableRange
//go:noescape
func HasFuncMediaSourceClearLiveSeekableRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_ClearLiveSeekableRange
//go:noescape
func FuncMediaSourceClearLiveSeekableRange(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_ClearLiveSeekableRange
//go:noescape
func CallMediaSourceClearLiveSeekableRange(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaSource_ClearLiveSeekableRange
//go:noescape
func TryMediaSourceClearLiveSeekableRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSource_IsTypeSupported
//go:noescape
func HasFuncMediaSourceIsTypeSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSource_IsTypeSupported
//go:noescape
func FuncMediaSourceIsTypeSupported(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaSource_IsTypeSupported
//go:noescape
func CallMediaSourceIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_MediaSource_IsTypeSupported
//go:noescape
func TryMediaSourceIsTypeSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_RemotePlaybackState
//go:noescape
func ConstOfRemotePlaybackState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RemotePlayback_State
//go:noescape
func GetRemotePlaybackState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RemotePlayback_WatchAvailability
//go:noescape
func HasFuncRemotePlaybackWatchAvailability(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_WatchAvailability
//go:noescape
func FuncRemotePlaybackWatchAvailability(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RemotePlayback_WatchAvailability
//go:noescape
func CallRemotePlaybackWatchAvailability(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_RemotePlayback_WatchAvailability
//go:noescape
func TryRemotePlaybackWatchAvailability(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_RemotePlayback_CancelWatchAvailability
//go:noescape
func HasFuncRemotePlaybackCancelWatchAvailability(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_CancelWatchAvailability
//go:noescape
func FuncRemotePlaybackCancelWatchAvailability(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RemotePlayback_CancelWatchAvailability
//go:noescape
func CallRemotePlaybackCancelWatchAvailability(
	this js.Ref, retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/web try_RemotePlayback_CancelWatchAvailability
//go:noescape
func TryRemotePlaybackCancelWatchAvailability(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/web has_RemotePlayback_CancelWatchAvailability1
//go:noescape
func HasFuncRemotePlaybackCancelWatchAvailability1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_CancelWatchAvailability1
//go:noescape
func FuncRemotePlaybackCancelWatchAvailability1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RemotePlayback_CancelWatchAvailability1
//go:noescape
func CallRemotePlaybackCancelWatchAvailability1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RemotePlayback_CancelWatchAvailability1
//go:noescape
func TryRemotePlaybackCancelWatchAvailability1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RemotePlayback_Prompt
//go:noescape
func HasFuncRemotePlaybackPrompt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RemotePlayback_Prompt
//go:noescape
func FuncRemotePlaybackPrompt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RemotePlayback_Prompt
//go:noescape
func CallRemotePlaybackPrompt(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RemotePlayback_Prompt
//go:noescape
func TryRemotePlaybackPrompt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Error
//go:noescape
func GetHTMLMediaElementError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Src
//go:noescape
func GetHTMLMediaElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Src
//go:noescape
func SetHTMLMediaElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_SrcObject
//go:noescape
func GetHTMLMediaElementSrcObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_SrcObject
//go:noescape
func SetHTMLMediaElementSrcObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_CurrentSrc
//go:noescape
func GetHTMLMediaElementCurrentSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_CrossOrigin
//go:noescape
func GetHTMLMediaElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_CrossOrigin
//go:noescape
func SetHTMLMediaElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_NetworkState
//go:noescape
func GetHTMLMediaElementNetworkState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Preload
//go:noescape
func GetHTMLMediaElementPreload(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Preload
//go:noescape
func SetHTMLMediaElementPreload(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Buffered
//go:noescape
func GetHTMLMediaElementBuffered(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_ReadyState
//go:noescape
func GetHTMLMediaElementReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Seeking
//go:noescape
func GetHTMLMediaElementSeeking(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_CurrentTime
//go:noescape
func GetHTMLMediaElementCurrentTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_CurrentTime
//go:noescape
func SetHTMLMediaElementCurrentTime(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Duration
//go:noescape
func GetHTMLMediaElementDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Paused
//go:noescape
func GetHTMLMediaElementPaused(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_DefaultPlaybackRate
//go:noescape
func GetHTMLMediaElementDefaultPlaybackRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_DefaultPlaybackRate
//go:noescape
func SetHTMLMediaElementDefaultPlaybackRate(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_PlaybackRate
//go:noescape
func GetHTMLMediaElementPlaybackRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_PlaybackRate
//go:noescape
func SetHTMLMediaElementPlaybackRate(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_PreservesPitch
//go:noescape
func GetHTMLMediaElementPreservesPitch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_PreservesPitch
//go:noescape
func SetHTMLMediaElementPreservesPitch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Played
//go:noescape
func GetHTMLMediaElementPlayed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Seekable
//go:noescape
func GetHTMLMediaElementSeekable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Ended
//go:noescape
func GetHTMLMediaElementEnded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Autoplay
//go:noescape
func GetHTMLMediaElementAutoplay(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Autoplay
//go:noescape
func SetHTMLMediaElementAutoplay(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Loop
//go:noescape
func GetHTMLMediaElementLoop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Loop
//go:noescape
func SetHTMLMediaElementLoop(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Controls
//go:noescape
func GetHTMLMediaElementControls(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Controls
//go:noescape
func SetHTMLMediaElementControls(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Volume
//go:noescape
func GetHTMLMediaElementVolume(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Volume
//go:noescape
func SetHTMLMediaElementVolume(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_Muted
//go:noescape
func GetHTMLMediaElementMuted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_Muted
//go:noescape
func SetHTMLMediaElementMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_DefaultMuted
//go:noescape
func GetHTMLMediaElementDefaultMuted(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_DefaultMuted
//go:noescape
func SetHTMLMediaElementDefaultMuted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_AudioTracks
//go:noescape
func GetHTMLMediaElementAudioTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_VideoTracks
//go:noescape
func GetHTMLMediaElementVideoTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_TextTracks
//go:noescape
func GetHTMLMediaElementTextTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_MediaKeys
//go:noescape
func GetHTMLMediaElementMediaKeys(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_Remote
//go:noescape
func GetHTMLMediaElementRemote(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLMediaElement_DisableRemotePlayback
//go:noescape
func GetHTMLMediaElementDisableRemotePlayback(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMediaElement_DisableRemotePlayback
//go:noescape
func SetHTMLMediaElementDisableRemotePlayback(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMediaElement_SinkId
//go:noescape
func GetHTMLMediaElementSinkId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_Load
//go:noescape
func HasFuncHTMLMediaElementLoad(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Load
//go:noescape
func FuncHTMLMediaElementLoad(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_Load
//go:noescape
func CallHTMLMediaElementLoad(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMediaElement_Load
//go:noescape
func TryHTMLMediaElementLoad(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_CanPlayType
//go:noescape
func HasFuncHTMLMediaElementCanPlayType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_CanPlayType
//go:noescape
func FuncHTMLMediaElementCanPlayType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_CanPlayType
//go:noescape
func CallHTMLMediaElementCanPlayType(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_HTMLMediaElement_CanPlayType
//go:noescape
func TryHTMLMediaElementCanPlayType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_FastSeek
//go:noescape
func HasFuncHTMLMediaElementFastSeek(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_FastSeek
//go:noescape
func FuncHTMLMediaElementFastSeek(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_FastSeek
//go:noescape
func CallHTMLMediaElementFastSeek(
	this js.Ref, retPtr unsafe.Pointer,
	time float64)

//go:wasmimport plat/js/web try_HTMLMediaElement_FastSeek
//go:noescape
func TryHTMLMediaElementFastSeek(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	time float64) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_GetStartDate
//go:noescape
func HasFuncHTMLMediaElementGetStartDate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_GetStartDate
//go:noescape
func FuncHTMLMediaElementGetStartDate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_GetStartDate
//go:noescape
func CallHTMLMediaElementGetStartDate(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMediaElement_GetStartDate
//go:noescape
func TryHTMLMediaElementGetStartDate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_Play
//go:noescape
func HasFuncHTMLMediaElementPlay(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Play
//go:noescape
func FuncHTMLMediaElementPlay(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_Play
//go:noescape
func CallHTMLMediaElementPlay(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMediaElement_Play
//go:noescape
func TryHTMLMediaElementPlay(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_Pause
//go:noescape
func HasFuncHTMLMediaElementPause(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_Pause
//go:noescape
func FuncHTMLMediaElementPause(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_Pause
//go:noescape
func CallHTMLMediaElementPause(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMediaElement_Pause
//go:noescape
func TryHTMLMediaElementPause(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_AddTextTrack
//go:noescape
func HasFuncHTMLMediaElementAddTextTrack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack
//go:noescape
func FuncHTMLMediaElementAddTextTrack(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack
//go:noescape
func CallHTMLMediaElementAddTextTrack(
	this js.Ref, retPtr unsafe.Pointer,
	kind uint32,
	label js.Ref,
	language js.Ref)

//go:wasmimport plat/js/web try_HTMLMediaElement_AddTextTrack
//go:noescape
func TryHTMLMediaElementAddTextTrack(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	kind uint32,
	label js.Ref,
	language js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_AddTextTrack1
//go:noescape
func HasFuncHTMLMediaElementAddTextTrack1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack1
//go:noescape
func FuncHTMLMediaElementAddTextTrack1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack1
//go:noescape
func CallHTMLMediaElementAddTextTrack1(
	this js.Ref, retPtr unsafe.Pointer,
	kind uint32,
	label js.Ref)

//go:wasmimport plat/js/web try_HTMLMediaElement_AddTextTrack1
//go:noescape
func TryHTMLMediaElementAddTextTrack1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	kind uint32,
	label js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_AddTextTrack2
//go:noescape
func HasFuncHTMLMediaElementAddTextTrack2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_AddTextTrack2
//go:noescape
func FuncHTMLMediaElementAddTextTrack2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_AddTextTrack2
//go:noescape
func CallHTMLMediaElementAddTextTrack2(
	this js.Ref, retPtr unsafe.Pointer,
	kind uint32)

//go:wasmimport plat/js/web try_HTMLMediaElement_AddTextTrack2
//go:noescape
func TryHTMLMediaElementAddTextTrack2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	kind uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_SetMediaKeys
//go:noescape
func HasFuncHTMLMediaElementSetMediaKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_SetMediaKeys
//go:noescape
func FuncHTMLMediaElementSetMediaKeys(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_SetMediaKeys
//go:noescape
func CallHTMLMediaElementSetMediaKeys(
	this js.Ref, retPtr unsafe.Pointer,
	mediaKeys js.Ref)

//go:wasmimport plat/js/web try_HTMLMediaElement_SetMediaKeys
//go:noescape
func TryHTMLMediaElementSetMediaKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mediaKeys js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_CaptureStream
//go:noescape
func HasFuncHTMLMediaElementCaptureStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_CaptureStream
//go:noescape
func FuncHTMLMediaElementCaptureStream(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_CaptureStream
//go:noescape
func CallHTMLMediaElementCaptureStream(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMediaElement_CaptureStream
//go:noescape
func TryHTMLMediaElementCaptureStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMediaElement_SetSinkId
//go:noescape
func HasFuncHTMLMediaElementSetSinkId(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMediaElement_SetSinkId
//go:noescape
func FuncHTMLMediaElementSetSinkId(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLMediaElement_SetSinkId
//go:noescape
func CallHTMLMediaElementSetSinkId(
	this js.Ref, retPtr unsafe.Pointer,
	sinkId js.Ref)

//go:wasmimport plat/js/web try_HTMLMediaElement_SetSinkId
//go:noescape
func TryHTMLMediaElementSetSinkId(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sinkId js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioSinkInfo_Type
//go:noescape
func GetAudioSinkInfoType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_AudioRenderCapacityOptions
//go:noescape
func AudioRenderCapacityOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AudioRenderCapacityOptions
//go:noescape
func AudioRenderCapacityOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_AudioRenderCapacity_Start
//go:noescape
func HasFuncAudioRenderCapacityStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Start
//go:noescape
func FuncAudioRenderCapacityStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioRenderCapacity_Start
//go:noescape
func CallAudioRenderCapacityStart(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioRenderCapacity_Start
//go:noescape
func TryAudioRenderCapacityStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioRenderCapacity_Start1
//go:noescape
func HasFuncAudioRenderCapacityStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Start1
//go:noescape
func FuncAudioRenderCapacityStart1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioRenderCapacity_Start1
//go:noescape
func CallAudioRenderCapacityStart1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioRenderCapacity_Start1
//go:noescape
func TryAudioRenderCapacityStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioRenderCapacity_Stop
//go:noescape
func HasFuncAudioRenderCapacityStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioRenderCapacity_Stop
//go:noescape
func FuncAudioRenderCapacityStop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioRenderCapacity_Stop
//go:noescape
func CallAudioRenderCapacityStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioRenderCapacity_Stop
//go:noescape
func TryAudioRenderCapacityStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioContext_OutputLatency
//go:noescape
func GetAudioContextOutputLatency(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioContext_SinkId
//go:noescape
func GetAudioContextSinkId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioContext_RenderCapacity
//go:noescape
func GetAudioContextRenderCapacity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_GetOutputTimestamp
//go:noescape
func HasFuncAudioContextGetOutputTimestamp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_GetOutputTimestamp
//go:noescape
func FuncAudioContextGetOutputTimestamp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_GetOutputTimestamp
//go:noescape
func CallAudioContextGetOutputTimestamp(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioContext_GetOutputTimestamp
//go:noescape
func TryAudioContextGetOutputTimestamp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_Resume
//go:noescape
func HasFuncAudioContextResume(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Resume
//go:noescape
func FuncAudioContextResume(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_Resume
//go:noescape
func CallAudioContextResume(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioContext_Resume
//go:noescape
func TryAudioContextResume(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_Suspend
//go:noescape
func HasFuncAudioContextSuspend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Suspend
//go:noescape
func FuncAudioContextSuspend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_Suspend
//go:noescape
func CallAudioContextSuspend(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioContext_Suspend
//go:noescape
func TryAudioContextSuspend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_Close
//go:noescape
func HasFuncAudioContextClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_Close
//go:noescape
func FuncAudioContextClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_Close
//go:noescape
func CallAudioContextClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioContext_Close
//go:noescape
func TryAudioContextClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_SetSinkId
//go:noescape
func HasFuncAudioContextSetSinkId(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_SetSinkId
//go:noescape
func FuncAudioContextSetSinkId(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_SetSinkId
//go:noescape
func CallAudioContextSetSinkId(
	this js.Ref, retPtr unsafe.Pointer,
	sinkId js.Ref)

//go:wasmimport plat/js/web try_AudioContext_SetSinkId
//go:noescape
func TryAudioContextSetSinkId(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sinkId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_CreateMediaElementSource
//go:noescape
func HasFuncAudioContextCreateMediaElementSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaElementSource
//go:noescape
func FuncAudioContextCreateMediaElementSource(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_CreateMediaElementSource
//go:noescape
func CallAudioContextCreateMediaElementSource(
	this js.Ref, retPtr unsafe.Pointer,
	mediaElement js.Ref)

//go:wasmimport plat/js/web try_AudioContext_CreateMediaElementSource
//go:noescape
func TryAudioContextCreateMediaElementSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mediaElement js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_CreateMediaStreamSource
//go:noescape
func HasFuncAudioContextCreateMediaStreamSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamSource
//go:noescape
func FuncAudioContextCreateMediaStreamSource(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamSource
//go:noescape
func CallAudioContextCreateMediaStreamSource(
	this js.Ref, retPtr unsafe.Pointer,
	mediaStream js.Ref)

//go:wasmimport plat/js/web try_AudioContext_CreateMediaStreamSource
//go:noescape
func TryAudioContextCreateMediaStreamSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mediaStream js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_CreateMediaStreamTrackSource
//go:noescape
func HasFuncAudioContextCreateMediaStreamTrackSource(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamTrackSource
//go:noescape
func FuncAudioContextCreateMediaStreamTrackSource(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamTrackSource
//go:noescape
func CallAudioContextCreateMediaStreamTrackSource(
	this js.Ref, retPtr unsafe.Pointer,
	mediaStreamTrack js.Ref)

//go:wasmimport plat/js/web try_AudioContext_CreateMediaStreamTrackSource
//go:noescape
func TryAudioContextCreateMediaStreamTrackSource(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mediaStreamTrack js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioContext_CreateMediaStreamDestination
//go:noescape
func HasFuncAudioContextCreateMediaStreamDestination(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioContext_CreateMediaStreamDestination
//go:noescape
func FuncAudioContextCreateMediaStreamDestination(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioContext_CreateMediaStreamDestination
//go:noescape
func CallAudioContextCreateMediaStreamDestination(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioContext_CreateMediaStreamDestination
//go:noescape
func TryAudioContextCreateMediaStreamDestination(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioData_SampleRate
//go:noescape
func GetAudioDataSampleRate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioData_NumberOfFrames
//go:noescape
func GetAudioDataNumberOfFrames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioData_NumberOfChannels
//go:noescape
func GetAudioDataNumberOfChannels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioData_Duration
//go:noescape
func GetAudioDataDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioData_Timestamp
//go:noescape
func GetAudioDataTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioData_AllocationSize
//go:noescape
func HasFuncAudioDataAllocationSize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioData_AllocationSize
//go:noescape
func FuncAudioDataAllocationSize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioData_AllocationSize
//go:noescape
func CallAudioDataAllocationSize(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioData_AllocationSize
//go:noescape
func TryAudioDataAllocationSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioData_CopyTo
//go:noescape
func HasFuncAudioDataCopyTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioData_CopyTo
//go:noescape
func FuncAudioDataCopyTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioData_CopyTo
//go:noescape
func CallAudioDataCopyTo(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioData_CopyTo
//go:noescape
func TryAudioDataCopyTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioData_Clone
//go:noescape
func HasFuncAudioDataClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioData_Clone
//go:noescape
func FuncAudioDataClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioData_Clone
//go:noescape
func CallAudioDataClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioData_Clone
//go:noescape
func TryAudioDataClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioData_Close
//go:noescape
func HasFuncAudioDataClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioData_Close
//go:noescape
func FuncAudioDataClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioData_Close
//go:noescape
func CallAudioDataClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioData_Close
//go:noescape
func TryAudioDataClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedAudioChunk_Timestamp
//go:noescape
func GetEncodedAudioChunkTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedAudioChunk_Duration
//go:noescape
func GetEncodedAudioChunkDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EncodedAudioChunk_ByteLength
//go:noescape
func GetEncodedAudioChunkByteLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EncodedAudioChunk_CopyTo
//go:noescape
func HasFuncEncodedAudioChunkCopyTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EncodedAudioChunk_CopyTo
//go:noescape
func FuncEncodedAudioChunkCopyTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EncodedAudioChunk_CopyTo
//go:noescape
func CallEncodedAudioChunkCopyTo(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref)

//go:wasmimport plat/js/web try_EncodedAudioChunk_CopyTo
//go:noescape
func TryEncodedAudioChunkCopyTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioDecoder_DecodeQueueSize
//go:noescape
func GetAudioDecoderDecodeQueueSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_Configure
//go:noescape
func HasFuncAudioDecoderConfigure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Configure
//go:noescape
func FuncAudioDecoderConfigure(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_Configure
//go:noescape
func CallAudioDecoderConfigure(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioDecoder_Configure
//go:noescape
func TryAudioDecoderConfigure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_Decode
//go:noescape
func HasFuncAudioDecoderDecode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Decode
//go:noescape
func FuncAudioDecoderDecode(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_Decode
//go:noescape
func CallAudioDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer,
	chunk js.Ref)

//go:wasmimport plat/js/web try_AudioDecoder_Decode
//go:noescape
func TryAudioDecoderDecode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	chunk js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_Flush
//go:noescape
func HasFuncAudioDecoderFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Flush
//go:noescape
func FuncAudioDecoderFlush(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_Flush
//go:noescape
func CallAudioDecoderFlush(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioDecoder_Flush
//go:noescape
func TryAudioDecoderFlush(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_Reset
//go:noescape
func HasFuncAudioDecoderReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Reset
//go:noescape
func FuncAudioDecoderReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_Reset
//go:noescape
func CallAudioDecoderReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioDecoder_Reset
//go:noescape
func TryAudioDecoderReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_Close
//go:noescape
func HasFuncAudioDecoderClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_Close
//go:noescape
func FuncAudioDecoderClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_Close
//go:noescape
func CallAudioDecoderClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioDecoder_Close
//go:noescape
func TryAudioDecoderClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioDecoder_IsConfigSupported
//go:noescape
func HasFuncAudioDecoderIsConfigSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioDecoder_IsConfigSupported
//go:noescape
func FuncAudioDecoderIsConfigSupported(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioDecoder_IsConfigSupported
//go:noescape
func CallAudioDecoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_AudioDecoder_IsConfigSupported
//go:noescape
func TryAudioDecoderIsConfigSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_EncodedAudioChunkMetadata
//go:noescape
func EncodedAudioChunkMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EncodedAudioChunkMetadata
//go:noescape
func EncodedAudioChunkMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
