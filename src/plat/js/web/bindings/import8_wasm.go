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

//go:wasmimport plat/js/web constof_TextTrackMode
//go:noescape

func ConstOfTextTrackMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_TextTrackCueList_Length
//go:noescape

func GetTextTrackCueListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_TextTrackCueList_Get
//go:noescape

func CallTextTrackCueListGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_TextTrackCueList_Get
//go:noescape

func TextTrackCueListGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextTrackCueList_GetCueById
//go:noescape

func CallTextTrackCueListGetCueById(
	this js.Ref,
	ptr unsafe.Pointer,

	id js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextTrackCueList_GetCueById
//go:noescape

func TextTrackCueListGetCueByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TimeRanges_Length
//go:noescape

func GetTimeRangesLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_TimeRanges_Start
//go:noescape

func CallTimeRangesStart(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) float64

//go:wasmimport plat/js/web func_TimeRanges_Start
//go:noescape

func TimeRangesStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TimeRanges_End
//go:noescape

func CallTimeRangesEnd(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) float64

//go:wasmimport plat/js/web func_TimeRanges_End
//go:noescape

func TimeRangesEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_Id
//go:noescape

func GetAudioTrackId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_Kind
//go:noescape

func GetAudioTrackKind(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_Label
//go:noescape

func GetAudioTrackLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_Language
//go:noescape

func GetAudioTrackLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_Enabled
//go:noescape

func GetAudioTrackEnabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_AudioTrack_Enabled
//go:noescape

func SetAudioTrackEnabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrack_SourceBuffer
//go:noescape

func GetAudioTrackSourceBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AudioTrackList_Length
//go:noescape

func GetAudioTrackListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_AudioTrackList_Get
//go:noescape

func CallAudioTrackListGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_AudioTrackList_Get
//go:noescape

func AudioTrackListGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AudioTrackList_GetTrackById
//go:noescape

func CallAudioTrackListGetTrackById(
	this js.Ref,
	ptr unsafe.Pointer,

	id js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AudioTrackList_GetTrackById
//go:noescape

func AudioTrackListGetTrackByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_Id
//go:noescape

func GetVideoTrackId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_Kind
//go:noescape

func GetVideoTrackKind(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_Label
//go:noescape

func GetVideoTrackLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_Language
//go:noescape

func GetVideoTrackLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_Selected
//go:noescape

func GetVideoTrackSelected(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_VideoTrack_Selected
//go:noescape

func SetVideoTrackSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrack_SourceBuffer
//go:noescape

func GetVideoTrackSourceBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_VideoTrackList_Length
//go:noescape

func GetVideoTrackListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_VideoTrackList_SelectedIndex
//go:noescape

func GetVideoTrackListSelectedIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web call_VideoTrackList_Get
//go:noescape

func CallVideoTrackListGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_VideoTrackList_Get
//go:noescape

func VideoTrackListGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_VideoTrackList_GetTrackById
//go:noescape

func CallVideoTrackListGetTrackById(
	this js.Ref,
	ptr unsafe.Pointer,

	id js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_VideoTrackList_GetTrackById
//go:noescape

func VideoTrackListGetTrackByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TextTrackList_Length
//go:noescape

func GetTextTrackListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_TextTrackList_Get
//go:noescape

func CallTextTrackListGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_TextTrackList_Get
//go:noescape

func TextTrackListGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextTrackList_GetTrackById
//go:noescape

func CallTextTrackListGetTrackById(
	this js.Ref,
	ptr unsafe.Pointer,

	id js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextTrackList_GetTrackById
//go:noescape

func TextTrackListGetTrackByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_Mode
//go:noescape

func GetSourceBufferMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_SourceBuffer_Mode
//go:noescape

func SetSourceBufferMode(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_Updating
//go:noescape

func GetSourceBufferUpdating(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_Buffered
//go:noescape

func GetSourceBufferBuffered(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_TimestampOffset
//go:noescape

func GetSourceBufferTimestampOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_SourceBuffer_TimestampOffset
//go:noescape

func SetSourceBufferTimestampOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_AudioTracks
//go:noescape

func GetSourceBufferAudioTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_VideoTracks
//go:noescape

func GetSourceBufferVideoTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_TextTracks
//go:noescape

func GetSourceBufferTextTracks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_AppendWindowStart
//go:noescape

func GetSourceBufferAppendWindowStart(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_SourceBuffer_AppendWindowStart
//go:noescape

func SetSourceBufferAppendWindowStart(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_SourceBuffer_AppendWindowEnd
//go:noescape

func GetSourceBufferAppendWindowEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_SourceBuffer_AppendWindowEnd
//go:noescape

func SetSourceBufferAppendWindowEnd(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_SourceBuffer_AppendBuffer
//go:noescape

func CallSourceBufferAppendBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_AppendBuffer
//go:noescape

func SourceBufferAppendBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SourceBuffer_Abort
//go:noescape

func CallSourceBufferAbort(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_Abort
//go:noescape

func SourceBufferAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SourceBuffer_ChangeType
//go:noescape

func CallSourceBufferChangeType(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_ChangeType
//go:noescape

func SourceBufferChangeTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SourceBuffer_Remove
//go:noescape

func CallSourceBufferRemove(
	this js.Ref,
	ptr unsafe.Pointer,

	start float64,
	end float64,
) js.Ref

//go:wasmimport plat/js/web func_SourceBuffer_Remove
//go:noescape

func SourceBufferRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Kind
//go:noescape

func GetTextTrackKind(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TextTrack_Label
//go:noescape

func GetTextTrackLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Language
//go:noescape

func GetTextTrackLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Id
//go:noescape

func GetTextTrackId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_InBandMetadataTrackDispatchType
//go:noescape

func GetTextTrackInBandMetadataTrackDispatchType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Mode
//go:noescape

func GetTextTrackMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_TextTrack_Mode
//go:noescape

func SetTextTrackMode(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_Cues
//go:noescape

func GetTextTrackCues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_ActiveCues
//go:noescape

func GetTextTrackActiveCues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TextTrack_SourceBuffer
//go:noescape

func GetTextTrackSourceBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_TextTrack_AddCue
//go:noescape

func CallTextTrackAddCue(
	this js.Ref,
	ptr unsafe.Pointer,

	cue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextTrack_AddCue
//go:noescape

func TextTrackAddCueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TextTrack_RemoveCue
//go:noescape

func CallTextTrackRemoveCue(
	this js.Ref,
	ptr unsafe.Pointer,

	cue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_TextTrack_RemoveCue
//go:noescape

func TextTrackRemoveCueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaKeySessionClosedReason
//go:noescape

func ConstOfMediaKeySessionClosedReason(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MediaKeyStatus
//go:noescape

func ConstOfMediaKeyStatus(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaKeyStatusMap_Size
//go:noescape

func GetMediaKeyStatusMapSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_MediaKeyStatusMap_Has
//go:noescape

func CallMediaKeyStatusMapHas(
	this js.Ref,
	ptr unsafe.Pointer,

	keyId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeyStatusMap_Has
//go:noescape

func MediaKeyStatusMapHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeyStatusMap_Get
//go:noescape

func CallMediaKeyStatusMapGet(
	this js.Ref,
	ptr unsafe.Pointer,

	keyId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeyStatusMap_Get
//go:noescape

func MediaKeyStatusMapGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaKeySession_SessionId
//go:noescape

func GetMediaKeySessionSessionId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaKeySession_Expiration
//go:noescape

func GetMediaKeySessionExpiration(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_MediaKeySession_Closed
//go:noescape

func GetMediaKeySessionClosed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_MediaKeySession_KeyStatuses
//go:noescape

func GetMediaKeySessionKeyStatuses(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_MediaKeySession_GenerateRequest
//go:noescape

func CallMediaKeySessionGenerateRequest(
	this js.Ref,
	ptr unsafe.Pointer,

	initDataType js.Ref,
	initData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_GenerateRequest
//go:noescape

func MediaKeySessionGenerateRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeySession_Load
//go:noescape

func CallMediaKeySessionLoad(
	this js.Ref,
	ptr unsafe.Pointer,

	sessionId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Load
//go:noescape

func MediaKeySessionLoadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeySession_Update
//go:noescape

func CallMediaKeySessionUpdate(
	this js.Ref,
	ptr unsafe.Pointer,

	response js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Update
//go:noescape

func MediaKeySessionUpdateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeySession_Close
//go:noescape

func CallMediaKeySessionClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Close
//go:noescape

func MediaKeySessionCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeySession_Remove
//go:noescape

func CallMediaKeySessionRemove(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaKeySession_Remove
//go:noescape

func MediaKeySessionRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaKeySessionType
//go:noescape

func ConstOfMediaKeySessionType(str js.Ref) uint32

//go:wasmimport plat/js/web call_MediaKeys_CreateSession
//go:noescape

func CallMediaKeysCreateSession(
	this js.Ref,
	ptr unsafe.Pointer,

	sessionType uint32,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_CreateSession
//go:noescape

func MediaKeysCreateSessionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeys_CreateSession1
//go:noescape

func CallMediaKeysCreateSession1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_CreateSession1
//go:noescape

func MediaKeysCreateSession1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaKeys_SetServerCertificate
//go:noescape

func CallMediaKeysSetServerCertificate(
	this js.Ref,
	ptr unsafe.Pointer,

	serverCertificate js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaKeys_SetServerCertificate
//go:noescape

func MediaKeysSetServerCertificateFunc(this js.Ref) js.Ref

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
