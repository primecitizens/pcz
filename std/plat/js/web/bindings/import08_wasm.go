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

//go:wasmimport plat/js/web constof_TextTrackMode
//go:noescape
func ConstOfTextTrackMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_TextTrackCueList_Length
//go:noescape
func GetTextTrackCueListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrackCueList_Get
//go:noescape
func HasFuncTextTrackCueListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrackCueList_Get
//go:noescape
func FuncTextTrackCueListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrackCueList_Get
//go:noescape
func CallTextTrackCueListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_TextTrackCueList_Get
//go:noescape
func TryTextTrackCueListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrackCueList_GetCueById
//go:noescape
func HasFuncTextTrackCueListGetCueById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrackCueList_GetCueById
//go:noescape
func FuncTextTrackCueListGetCueById(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrackCueList_GetCueById
//go:noescape
func CallTextTrackCueListGetCueById(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_TextTrackCueList_GetCueById
//go:noescape
func TryTextTrackCueListGetCueById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_TimeRanges_Length
//go:noescape
func GetTimeRangesLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TimeRanges_Start
//go:noescape
func HasFuncTimeRangesStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TimeRanges_Start
//go:noescape
func FuncTimeRangesStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TimeRanges_Start
//go:noescape
func CallTimeRangesStart(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_TimeRanges_Start
//go:noescape
func TryTimeRangesStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_TimeRanges_End
//go:noescape
func HasFuncTimeRangesEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TimeRanges_End
//go:noescape
func FuncTimeRangesEnd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TimeRanges_End
//go:noescape
func CallTimeRangesEnd(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_TimeRanges_End
//go:noescape
func TryTimeRangesEnd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrack_Id
//go:noescape
func GetAudioTrackId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrack_Kind
//go:noescape
func GetAudioTrackKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrack_Label
//go:noescape
func GetAudioTrackLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrack_Language
//go:noescape
func GetAudioTrackLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrack_Enabled
//go:noescape
func GetAudioTrackEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_AudioTrack_Enabled
//go:noescape
func SetAudioTrackEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_SourceBuffer
//go:noescape
func GetAudioTrackSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AudioTrackList_Length
//go:noescape
func GetAudioTrackListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioTrackList_Get
//go:noescape
func HasFuncAudioTrackListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioTrackList_Get
//go:noescape
func FuncAudioTrackListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioTrackList_Get
//go:noescape
func CallAudioTrackListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_AudioTrackList_Get
//go:noescape
func TryAudioTrackListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_AudioTrackList_GetTrackById
//go:noescape
func HasFuncAudioTrackListGetTrackById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AudioTrackList_GetTrackById
//go:noescape
func FuncAudioTrackListGetTrackById(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AudioTrackList_GetTrackById
//go:noescape
func CallAudioTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_AudioTrackList_GetTrackById
//go:noescape
func TryAudioTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrack_Id
//go:noescape
func GetVideoTrackId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrack_Kind
//go:noescape
func GetVideoTrackKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrack_Label
//go:noescape
func GetVideoTrackLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrack_Language
//go:noescape
func GetVideoTrackLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrack_Selected
//go:noescape
func GetVideoTrackSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_VideoTrack_Selected
//go:noescape
func SetVideoTrackSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_SourceBuffer
//go:noescape
func GetVideoTrackSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrackList_Length
//go:noescape
func GetVideoTrackListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VideoTrackList_SelectedIndex
//go:noescape
func GetVideoTrackListSelectedIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoTrackList_Get
//go:noescape
func HasFuncVideoTrackListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoTrackList_Get
//go:noescape
func FuncVideoTrackListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoTrackList_Get
//go:noescape
func CallVideoTrackListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_VideoTrackList_Get
//go:noescape
func TryVideoTrackListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_VideoTrackList_GetTrackById
//go:noescape
func HasFuncVideoTrackListGetTrackById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_VideoTrackList_GetTrackById
//go:noescape
func FuncVideoTrackListGetTrackById(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_VideoTrackList_GetTrackById
//go:noescape
func CallVideoTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_VideoTrackList_GetTrackById
//go:noescape
func TryVideoTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrackList_Length
//go:noescape
func GetTextTrackListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrackList_Get
//go:noescape
func HasFuncTextTrackListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrackList_Get
//go:noescape
func FuncTextTrackListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrackList_Get
//go:noescape
func CallTextTrackListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_TextTrackList_Get
//go:noescape
func TryTextTrackListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrackList_GetTrackById
//go:noescape
func HasFuncTextTrackListGetTrackById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrackList_GetTrackById
//go:noescape
func FuncTextTrackListGetTrackById(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrackList_GetTrackById
//go:noescape
func CallTextTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_TextTrackList_GetTrackById
//go:noescape
func TryTextTrackListGetTrackById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_Mode
//go:noescape
func GetSourceBufferMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SourceBuffer_Mode
//go:noescape
func SetSourceBufferMode(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_Updating
//go:noescape
func GetSourceBufferUpdating(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_Buffered
//go:noescape
func GetSourceBufferBuffered(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_TimestampOffset
//go:noescape
func GetSourceBufferTimestampOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SourceBuffer_TimestampOffset
//go:noescape
func SetSourceBufferTimestampOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_AudioTracks
//go:noescape
func GetSourceBufferAudioTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_VideoTracks
//go:noescape
func GetSourceBufferVideoTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_TextTracks
//go:noescape
func GetSourceBufferTextTracks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SourceBuffer_AppendWindowStart
//go:noescape
func GetSourceBufferAppendWindowStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SourceBuffer_AppendWindowStart
//go:noescape
func SetSourceBufferAppendWindowStart(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_AppendWindowEnd
//go:noescape
func GetSourceBufferAppendWindowEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SourceBuffer_AppendWindowEnd
//go:noescape
func SetSourceBufferAppendWindowEnd(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_SourceBuffer_AppendBuffer
//go:noescape
func HasFuncSourceBufferAppendBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_AppendBuffer
//go:noescape
func FuncSourceBufferAppendBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SourceBuffer_AppendBuffer
//go:noescape
func CallSourceBufferAppendBuffer(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_SourceBuffer_AppendBuffer
//go:noescape
func TrySourceBufferAppendBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SourceBuffer_Abort
//go:noescape
func HasFuncSourceBufferAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_Abort
//go:noescape
func FuncSourceBufferAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SourceBuffer_Abort
//go:noescape
func CallSourceBufferAbort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SourceBuffer_Abort
//go:noescape
func TrySourceBufferAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SourceBuffer_ChangeType
//go:noescape
func HasFuncSourceBufferChangeType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_ChangeType
//go:noescape
func FuncSourceBufferChangeType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SourceBuffer_ChangeType
//go:noescape
func CallSourceBufferChangeType(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_SourceBuffer_ChangeType
//go:noescape
func TrySourceBufferChangeType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SourceBuffer_Remove
//go:noescape
func HasFuncSourceBufferRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_Remove
//go:noescape
func FuncSourceBufferRemove(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SourceBuffer_Remove
//go:noescape
func CallSourceBufferRemove(
	this js.Ref, retPtr unsafe.Pointer,
	start float64,
	end float64)

//go:wasmimport plat/js/web try_SourceBuffer_Remove
//go:noescape
func TrySourceBufferRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start float64,
	end float64) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_Kind
//go:noescape
func GetTextTrackKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_Label
//go:noescape
func GetTextTrackLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_Language
//go:noescape
func GetTextTrackLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_Id
//go:noescape
func GetTextTrackId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_InBandMetadataTrackDispatchType
//go:noescape
func GetTextTrackInBandMetadataTrackDispatchType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_Mode
//go:noescape
func GetTextTrackMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_TextTrack_Mode
//go:noescape
func SetTextTrackMode(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Cues
//go:noescape
func GetTextTrackCues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_ActiveCues
//go:noescape
func GetTextTrackActiveCues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TextTrack_SourceBuffer
//go:noescape
func GetTextTrackSourceBuffer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrack_AddCue
//go:noescape
func HasFuncTextTrackAddCue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrack_AddCue
//go:noescape
func FuncTextTrackAddCue(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrack_AddCue
//go:noescape
func CallTextTrackAddCue(
	this js.Ref, retPtr unsafe.Pointer,
	cue js.Ref)

//go:wasmimport plat/js/web try_TextTrack_AddCue
//go:noescape
func TryTextTrackAddCue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cue js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_TextTrack_RemoveCue
//go:noescape
func HasFuncTextTrackRemoveCue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TextTrack_RemoveCue
//go:noescape
func FuncTextTrackRemoveCue(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_TextTrack_RemoveCue
//go:noescape
func CallTextTrackRemoveCue(
	this js.Ref, retPtr unsafe.Pointer,
	cue js.Ref)

//go:wasmimport plat/js/web try_TextTrack_RemoveCue
//go:noescape
func TryTextTrackRemoveCue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cue js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_MediaKeySessionClosedReason
//go:noescape
func ConstOfMediaKeySessionClosedReason(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MediaKeyStatus
//go:noescape
func ConstOfMediaKeyStatus(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaKeyStatusMap_Size
//go:noescape
func GetMediaKeyStatusMapSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeyStatusMap_Has
//go:noescape
func HasFuncMediaKeyStatusMapHas(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeyStatusMap_Has
//go:noescape
func FuncMediaKeyStatusMapHas(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeyStatusMap_Has
//go:noescape
func CallMediaKeyStatusMapHas(
	this js.Ref, retPtr unsafe.Pointer,
	keyId js.Ref)

//go:wasmimport plat/js/web try_MediaKeyStatusMap_Has
//go:noescape
func TryMediaKeyStatusMapHas(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeyStatusMap_Get
//go:noescape
func HasFuncMediaKeyStatusMapGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeyStatusMap_Get
//go:noescape
func FuncMediaKeyStatusMapGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeyStatusMap_Get
//go:noescape
func CallMediaKeyStatusMapGet(
	this js.Ref, retPtr unsafe.Pointer,
	keyId js.Ref)

//go:wasmimport plat/js/web try_MediaKeyStatusMap_Get
//go:noescape
func TryMediaKeyStatusMapGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaKeySession_SessionId
//go:noescape
func GetMediaKeySessionSessionId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaKeySession_Expiration
//go:noescape
func GetMediaKeySessionExpiration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaKeySession_Closed
//go:noescape
func GetMediaKeySessionClosed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaKeySession_KeyStatuses
//go:noescape
func GetMediaKeySessionKeyStatuses(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySession_GenerateRequest
//go:noescape
func HasFuncMediaKeySessionGenerateRequest(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_GenerateRequest
//go:noescape
func FuncMediaKeySessionGenerateRequest(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySession_GenerateRequest
//go:noescape
func CallMediaKeySessionGenerateRequest(
	this js.Ref, retPtr unsafe.Pointer,
	initDataType js.Ref,
	initData js.Ref)

//go:wasmimport plat/js/web try_MediaKeySession_GenerateRequest
//go:noescape
func TryMediaKeySessionGenerateRequest(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	initDataType js.Ref,
	initData js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySession_Load
//go:noescape
func HasFuncMediaKeySessionLoad(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Load
//go:noescape
func FuncMediaKeySessionLoad(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySession_Load
//go:noescape
func CallMediaKeySessionLoad(
	this js.Ref, retPtr unsafe.Pointer,
	sessionId js.Ref)

//go:wasmimport plat/js/web try_MediaKeySession_Load
//go:noescape
func TryMediaKeySessionLoad(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sessionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySession_Update
//go:noescape
func HasFuncMediaKeySessionUpdate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Update
//go:noescape
func FuncMediaKeySessionUpdate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySession_Update
//go:noescape
func CallMediaKeySessionUpdate(
	this js.Ref, retPtr unsafe.Pointer,
	response js.Ref)

//go:wasmimport plat/js/web try_MediaKeySession_Update
//go:noescape
func TryMediaKeySessionUpdate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	response js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySession_Close
//go:noescape
func HasFuncMediaKeySessionClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Close
//go:noescape
func FuncMediaKeySessionClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySession_Close
//go:noescape
func CallMediaKeySessionClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaKeySession_Close
//go:noescape
func TryMediaKeySessionClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeySession_Remove
//go:noescape
func HasFuncMediaKeySessionRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Remove
//go:noescape
func FuncMediaKeySessionRemove(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeySession_Remove
//go:noescape
func CallMediaKeySessionRemove(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaKeySession_Remove
//go:noescape
func TryMediaKeySessionRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MediaKeySessionType
//go:noescape
func ConstOfMediaKeySessionType(str js.Ref) uint32

//go:wasmimport plat/js/web has_MediaKeys_CreateSession
//go:noescape
func HasFuncMediaKeysCreateSession(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_CreateSession
//go:noescape
func FuncMediaKeysCreateSession(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeys_CreateSession
//go:noescape
func CallMediaKeysCreateSession(
	this js.Ref, retPtr unsafe.Pointer,
	sessionType uint32)

//go:wasmimport plat/js/web try_MediaKeys_CreateSession
//go:noescape
func TryMediaKeysCreateSession(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sessionType uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeys_CreateSession1
//go:noescape
func HasFuncMediaKeysCreateSession1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_CreateSession1
//go:noescape
func FuncMediaKeysCreateSession1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeys_CreateSession1
//go:noescape
func CallMediaKeysCreateSession1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaKeys_CreateSession1
//go:noescape
func TryMediaKeysCreateSession1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaKeys_SetServerCertificate
//go:noescape
func HasFuncMediaKeysSetServerCertificate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_SetServerCertificate
//go:noescape
func FuncMediaKeysSetServerCertificate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaKeys_SetServerCertificate
//go:noescape
func CallMediaKeysSetServerCertificate(
	this js.Ref, retPtr unsafe.Pointer,
	serverCertificate js.Ref)

//go:wasmimport plat/js/web try_MediaKeys_SetServerCertificate
//go:noescape
func TryMediaKeysSetServerCertificate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	serverCertificate js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_ULongRange
//go:noescape
func ULongRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ULongRange
//go:noescape
func ULongRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DoubleRange
//go:noescape
func DoubleRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DoubleRange
//go:noescape
func DoubleRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaSettingsRange
//go:noescape
func MediaSettingsRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaSettingsRange
//go:noescape
func MediaSettingsRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaTrackCapabilities
//go:noescape
func MediaTrackCapabilitiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaTrackCapabilities
//go:noescape
func MediaTrackCapabilitiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstrainULongRange
//go:noescape
func ConstrainULongRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstrainULongRange
//go:noescape
func ConstrainULongRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstrainDoubleRange
//go:noescape
func ConstrainDoubleRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstrainDoubleRange
//go:noescape
func ConstrainDoubleRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstrainDOMStringParameters
//go:noescape
func ConstrainDOMStringParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstrainDOMStringParameters
//go:noescape
func ConstrainDOMStringParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstrainBooleanParameters
//go:noescape
func ConstrainBooleanParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstrainBooleanParameters
//go:noescape
func ConstrainBooleanParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_Point2D
//go:noescape
func Point2DJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Point2D
//go:noescape
func Point2DJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ConstrainPoint2DParameters
//go:noescape
func ConstrainPoint2DParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConstrainPoint2DParameters
//go:noescape
func ConstrainPoint2DParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaTrackConstraintSet
//go:noescape
func MediaTrackConstraintSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaTrackConstraintSet
//go:noescape
func MediaTrackConstraintSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaTrackConstraints
//go:noescape
func MediaTrackConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaTrackConstraints
//go:noescape
func MediaTrackConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaTrackSettings
//go:noescape
func MediaTrackSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaTrackSettings
//go:noescape
func MediaTrackSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_CaptureAction
//go:noescape
func ConstOfCaptureAction(str js.Ref) uint32

//go:wasmimport plat/js/web store_CaptureHandle
//go:noescape
func CaptureHandleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CaptureHandle
//go:noescape
func CaptureHandleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaStreamTrackState
//go:noescape
func ConstOfMediaStreamTrackState(str js.Ref) uint32
