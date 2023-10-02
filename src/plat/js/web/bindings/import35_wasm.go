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

//go:wasmimport plat/js/web store_EncodedVideoChunkMetadata
//go:noescape

func EncodedVideoChunkMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EncodedVideoChunkMetadata
//go:noescape

func EncodedVideoChunkMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ErrorEventInit
//go:noescape

func ErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ErrorEventInit
//go:noescape

func ErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ErrorEvent_ErrorEvent
//go:noescape

func NewErrorEventByErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ErrorEvent_ErrorEvent1
//go:noescape

func NewErrorEventByErrorEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ErrorEvent_Message
//go:noescape

func GetErrorEventMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ErrorEvent_Filename
//go:noescape

func GetErrorEventFilename(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ErrorEvent_Lineno
//go:noescape

func GetErrorEventLineno(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ErrorEvent_Colno
//go:noescape

func GetErrorEventColno(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ErrorEvent_Error
//go:noescape

func GetErrorEventError(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_EventModifierInit
//go:noescape

func EventModifierInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EventModifierInit
//go:noescape

func EventModifierInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_EventSourceInit
//go:noescape

func EventSourceInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EventSourceInit
//go:noescape

func EventSourceInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_EventSource_EventSource
//go:noescape

func NewEventSourceByEventSource(
	url js.Ref,
	eventSourceInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_EventSource_EventSource1
//go:noescape

func NewEventSourceByEventSource1(
	url js.Ref) js.Ref

//go:wasmimport plat/js/web get_EventSource_Url
//go:noescape

func GetEventSourceUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EventSource_WithCredentials
//go:noescape

func GetEventSourceWithCredentials(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EventSource_ReadyState
//go:noescape

func GetEventSourceReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_EventSource_Close
//go:noescape

func CallEventSourceClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_EventSource_Close
//go:noescape

func EventSourceCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ExtendableCookieChangeEventInit
//go:noescape

func ExtendableCookieChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ExtendableCookieChangeEventInit
//go:noescape

func ExtendableCookieChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ExtendableCookieChangeEvent_ExtendableCookieChangeEvent
//go:noescape

func NewExtendableCookieChangeEventByExtendableCookieChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ExtendableCookieChangeEvent_ExtendableCookieChangeEvent1
//go:noescape

func NewExtendableCookieChangeEventByExtendableCookieChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ExtendableCookieChangeEvent_Changed
//go:noescape

func GetExtendableCookieChangeEventChanged(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ExtendableCookieChangeEvent_Deleted
//go:noescape

func GetExtendableCookieChangeEventDeleted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_ExtendableEventInit
//go:noescape

func ExtendableEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ExtendableEventInit
//go:noescape

func ExtendableEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ExtendableEvent_ExtendableEvent
//go:noescape

func NewExtendableEventByExtendableEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ExtendableEvent_ExtendableEvent1
//go:noescape

func NewExtendableEventByExtendableEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web call_ExtendableEvent_WaitUntil
//go:noescape

func CallExtendableEventWaitUntil(
	this js.Ref,
	ptr unsafe.Pointer,

	f js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ExtendableEvent_WaitUntil
//go:noescape

func ExtendableEventWaitUntilFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ExtendableMessageEventInit
//go:noescape

func ExtendableMessageEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ExtendableMessageEventInit
//go:noescape

func ExtendableMessageEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ExtendableMessageEvent_ExtendableMessageEvent
//go:noescape

func NewExtendableMessageEventByExtendableMessageEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ExtendableMessageEvent_ExtendableMessageEvent1
//go:noescape

func NewExtendableMessageEventByExtendableMessageEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ExtendableMessageEvent_Data
//go:noescape

func GetExtendableMessageEventData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ExtendableMessageEvent_Origin
//go:noescape

func GetExtendableMessageEventOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ExtendableMessageEvent_LastEventId
//go:noescape

func GetExtendableMessageEventLastEventId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ExtendableMessageEvent_Source
//go:noescape

func GetExtendableMessageEventSource(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ExtendableMessageEvent_Ports
//go:noescape

func GetExtendableMessageEventPorts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_EyeDropper_Open
//go:noescape

func CallEyeDropperOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_EyeDropper_Open
//go:noescape

func EyeDropperOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EyeDropper_Open1
//go:noescape

func CallEyeDropperOpen1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_EyeDropper_Open1
//go:noescape

func EyeDropperOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_FaceDetectorOptions
//go:noescape

func FaceDetectorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FaceDetectorOptions
//go:noescape

func FaceDetectorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FaceDetector_FaceDetector
//go:noescape

func NewFaceDetectorByFaceDetector(
	faceDetectorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FaceDetector_FaceDetector1
//go:noescape

func NewFaceDetectorByFaceDetector1() js.Ref

//go:wasmimport plat/js/web call_FaceDetector_Detect
//go:noescape

func CallFaceDetectorDetect(
	this js.Ref,
	ptr unsafe.Pointer,

	image js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FaceDetector_Detect
//go:noescape

func FaceDetectorDetectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_FederatedCredential_FederatedCredential
//go:noescape

func NewFederatedCredentialByFederatedCredential(
	data unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_FederatedCredential_Provider
//go:noescape

func GetFederatedCredentialProvider(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FederatedCredential_Protocol
//go:noescape

func GetFederatedCredentialProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FederatedCredential_Name
//go:noescape

func GetFederatedCredentialName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FederatedCredential_IconURL
//go:noescape

func GetFederatedCredentialIconURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_FetchEventInit
//go:noescape

func FetchEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FetchEventInit
//go:noescape

func FetchEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FetchEvent_FetchEvent
//go:noescape

func NewFetchEventByFetchEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_Request
//go:noescape

func GetFetchEventRequest(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_PreloadResponse
//go:noescape

func GetFetchEventPreloadResponse(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_ClientId
//go:noescape

func GetFetchEventClientId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_ResultingClientId
//go:noescape

func GetFetchEventResultingClientId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_ReplacesClientId
//go:noescape

func GetFetchEventReplacesClientId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FetchEvent_Handled
//go:noescape

func GetFetchEventHandled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_FetchEvent_RespondWith
//go:noescape

func CallFetchEventRespondWith(
	this js.Ref,
	ptr unsafe.Pointer,

	r js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FetchEvent_RespondWith
//go:noescape

func FetchEventRespondWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_FilePickerOptions
//go:noescape

func FilePickerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FilePickerOptions
//go:noescape

func FilePickerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_FileReader_ReadyState
//go:noescape

func GetFileReaderReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_FileReader_Result
//go:noescape

func GetFileReaderResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FileReader_Error
//go:noescape

func GetFileReaderError(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_FileReader_ReadAsArrayBuffer
//go:noescape

func CallFileReaderReadAsArrayBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReader_ReadAsArrayBuffer
//go:noescape

func FileReaderReadAsArrayBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReader_ReadAsBinaryString
//go:noescape

func CallFileReaderReadAsBinaryString(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReader_ReadAsBinaryString
//go:noescape

func FileReaderReadAsBinaryStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReader_ReadAsText
//go:noescape

func CallFileReaderReadAsText(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
	encoding js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReader_ReadAsText
//go:noescape

func FileReaderReadAsTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReader_ReadAsText1
//go:noescape

func CallFileReaderReadAsText1(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReader_ReadAsText1
//go:noescape

func FileReaderReadAsText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReader_ReadAsDataURL
//go:noescape

func CallFileReaderReadAsDataURL(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReader_ReadAsDataURL
//go:noescape

func FileReaderReadAsDataURLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReader_Abort
//go:noescape

func CallFileReaderAbort(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_FileReader_Abort
//go:noescape

func FileReaderAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReaderSync_ReadAsArrayBuffer
//go:noescape

func CallFileReaderSyncReadAsArrayBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReaderSync_ReadAsArrayBuffer
//go:noescape

func FileReaderSyncReadAsArrayBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReaderSync_ReadAsBinaryString
//go:noescape

func CallFileReaderSyncReadAsBinaryString(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReaderSync_ReadAsBinaryString
//go:noescape

func FileReaderSyncReadAsBinaryStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReaderSync_ReadAsText
//go:noescape

func CallFileReaderSyncReadAsText(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
	encoding js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReaderSync_ReadAsText
//go:noescape

func FileReaderSyncReadAsTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReaderSync_ReadAsText1
//go:noescape

func CallFileReaderSyncReadAsText1(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReaderSync_ReadAsText1
//go:noescape

func FileReaderSyncReadAsText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileReaderSync_ReadAsDataURL
//go:noescape

func CallFileReaderSyncReadAsDataURL(
	this js.Ref,
	ptr unsafe.Pointer,

	blob js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileReaderSync_ReadAsDataURL
//go:noescape

func FileReaderSyncReadAsDataURLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileEntry_File
//go:noescape

func CallFileSystemFileEntryFile(
	this js.Ref,
	ptr unsafe.Pointer,

	successCallback js.Ref,
	errorCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileEntry_File
//go:noescape

func FileSystemFileEntryFileFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileEntry_File1
//go:noescape

func CallFileSystemFileEntryFile1(
	this js.Ref,
	ptr unsafe.Pointer,

	successCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileEntry_File1
//go:noescape

func FileSystemFileEntryFile1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_FileSystemPermissionDescriptor
//go:noescape

func FileSystemPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemPermissionDescriptor
//go:noescape

func FileSystemPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_FillLightMode
//go:noescape

func ConstOfFillLightMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_FocusEventInit
//go:noescape

func FocusEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FocusEventInit
//go:noescape

func FocusEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FocusEvent_FocusEvent
//go:noescape

func NewFocusEventByFocusEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FocusEvent_FocusEvent1
//go:noescape

func NewFocusEventByFocusEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_FocusEvent_RelatedTarget
//go:noescape

func GetFocusEventRelatedTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_FontFaceSetLoadEventInit
//go:noescape

func FontFaceSetLoadEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FontFaceSetLoadEventInit
//go:noescape

func FontFaceSetLoadEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FontFaceSetLoadEvent_FontFaceSetLoadEvent
//go:noescape

func NewFontFaceSetLoadEventByFontFaceSetLoadEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FontFaceSetLoadEvent_FontFaceSetLoadEvent1
//go:noescape

func NewFontFaceSetLoadEventByFontFaceSetLoadEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_FontFaceSetLoadEvent_Fontfaces
//go:noescape

func GetFontFaceSetLoadEventFontfaces(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFaceVariationAxis_Name
//go:noescape

func GetFontFaceVariationAxisName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFaceVariationAxis_AxisTag
//go:noescape

func GetFontFaceVariationAxisAxisTag(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFaceVariationAxis_MinimumValue
//go:noescape

func GetFontFaceVariationAxisMinimumValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontFaceVariationAxis_MaximumValue
//go:noescape

func GetFontFaceVariationAxisMaximumValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontFaceVariationAxis_DefaultValue
//go:noescape

func GetFontFaceVariationAxisDefaultValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web store_FormDataEventInit
//go:noescape

func FormDataEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FormDataEventInit
//go:noescape

func FormDataEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FormDataEvent_FormDataEvent
//go:noescape

func NewFormDataEventByFormDataEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_FormDataEvent_FormData
//go:noescape

func GetFormDataEventFormData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_FragmentResultOptions
//go:noescape

func FragmentResultOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FragmentResultOptions
//go:noescape

func FragmentResultOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FragmentResult_FragmentResult
//go:noescape

func NewFragmentResultByFragmentResult(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FragmentResult_FragmentResult1
//go:noescape

func NewFragmentResultByFragmentResult1() js.Ref

//go:wasmimport plat/js/web get_FragmentResult_InlineSize
//go:noescape

func GetFragmentResultInlineSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FragmentResult_BlockSize
//go:noescape

func GetFragmentResultBlockSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64
