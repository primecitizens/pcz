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

//go:wasmimport plat/js/web get_HTMLAllCollection_Length
//go:noescape

func GetHTMLAllCollectionLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_HTMLAllCollection_Get
//go:noescape

func CallHTMLAllCollectionGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Get
//go:noescape

func HTMLAllCollectionGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLAllCollection_NamedItem
//go:noescape

func CallHTMLAllCollectionNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_NamedItem
//go:noescape

func HTMLAllCollectionNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLAllCollection_Item
//go:noescape

func CallHTMLAllCollectionItem(
	this js.Ref,
	ptr unsafe.Pointer,

	nameOrIndex js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Item
//go:noescape

func HTMLAllCollectionItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLAllCollection_Item1
//go:noescape

func CallHTMLAllCollectionItem1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Item1
//go:noescape

func HTMLAllCollectionItem1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_DOMStringList_Length
//go:noescape

func GetDOMStringListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_DOMStringList_Item
//go:noescape

func CallDOMStringListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_DOMStringList_Item
//go:noescape

func DOMStringListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringList_Contains
//go:noescape

func CallDOMStringListContains(
	this js.Ref,
	ptr unsafe.Pointer,

	string js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMStringList_Contains
//go:noescape

func DOMStringListContainsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_Location_Href
//go:noescape

func GetLocationHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Href
//go:noescape

func SetLocationHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Origin
//go:noescape

func GetLocationOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Location_Protocol
//go:noescape

func GetLocationProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Protocol
//go:noescape

func SetLocationProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Host
//go:noescape

func GetLocationHost(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Host
//go:noescape

func SetLocationHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Hostname
//go:noescape

func GetLocationHostname(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Hostname
//go:noescape

func SetLocationHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Port
//go:noescape

func GetLocationPort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Port
//go:noescape

func SetLocationPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Pathname
//go:noescape

func GetLocationPathname(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Pathname
//go:noescape

func SetLocationPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Search
//go:noescape

func GetLocationSearch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Search
//go:noescape

func SetLocationSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Hash
//go:noescape

func GetLocationHash(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Location_Hash
//go:noescape

func SetLocationHash(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_AncestorOrigins
//go:noescape

func GetLocationAncestorOrigins(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Location_Assign
//go:noescape

func CallLocationAssign(
	this js.Ref,
	ptr unsafe.Pointer,

	url js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Location_Assign
//go:noescape

func LocationAssignFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Location_Replace
//go:noescape

func CallLocationReplace(
	this js.Ref,
	ptr unsafe.Pointer,

	url js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Location_Replace
//go:noescape

func LocationReplaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Location_Reload
//go:noescape

func CallLocationReload(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Location_Reload
//go:noescape

func LocationReloadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_DocumentReadyState
//go:noescape

func ConstOfDocumentReadyState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_EndingType
//go:noescape

func ConstOfEndingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_BlobPropertyBag
//go:noescape

func BlobPropertyBagJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BlobPropertyBag
//go:noescape

func BlobPropertyBagJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_QueuingStrategy
//go:noescape

func QueuingStrategyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_QueuingStrategy
//go:noescape

func QueuingStrategyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReadableStreamReadResult
//go:noescape

func ReadableStreamReadResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReadableStreamReadResult
//go:noescape

func ReadableStreamReadResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ReadableStreamDefaultReader_ReadableStreamDefaultReader
//go:noescape

func NewReadableStreamDefaultReaderByReadableStreamDefaultReader(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web get_ReadableStreamDefaultReader_Closed
//go:noescape

func GetReadableStreamDefaultReaderClosed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Read
//go:noescape

func CallReadableStreamDefaultReaderRead(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Read
//go:noescape

func ReadableStreamDefaultReaderReadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_ReleaseLock
//go:noescape

func CallReadableStreamDefaultReaderReleaseLock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_ReleaseLock
//go:noescape

func ReadableStreamDefaultReaderReleaseLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Cancel
//go:noescape

func CallReadableStreamDefaultReaderCancel(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Cancel
//go:noescape

func ReadableStreamDefaultReaderCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Cancel1
//go:noescape

func CallReadableStreamDefaultReaderCancel1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Cancel1
//go:noescape

func ReadableStreamDefaultReaderCancel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_ReadableStreamBYOBReader_ReadableStreamBYOBReader
//go:noescape

func NewReadableStreamBYOBReaderByReadableStreamBYOBReader(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web get_ReadableStreamBYOBReader_Closed
//go:noescape

func GetReadableStreamBYOBReaderClosed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Read
//go:noescape

func CallReadableStreamBYOBReaderRead(
	this js.Ref,
	ptr unsafe.Pointer,

	view js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Read
//go:noescape

func ReadableStreamBYOBReaderReadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_ReleaseLock
//go:noescape

func CallReadableStreamBYOBReaderReleaseLock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_ReleaseLock
//go:noescape

func ReadableStreamBYOBReaderReleaseLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Cancel
//go:noescape

func CallReadableStreamBYOBReaderCancel(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Cancel
//go:noescape

func ReadableStreamBYOBReaderCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Cancel1
//go:noescape

func CallReadableStreamBYOBReaderCancel1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Cancel1
//go:noescape

func ReadableStreamBYOBReaderCancel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ReadableStreamReaderMode
//go:noescape

func ConstOfReadableStreamReaderMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_ReadableStreamGetReaderOptions
//go:noescape

func ReadableStreamGetReaderOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReadableStreamGetReaderOptions
//go:noescape

func ReadableStreamGetReaderOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_WritableStreamDefaultWriter_WritableStreamDefaultWriter
//go:noescape

func NewWritableStreamDefaultWriterByWritableStreamDefaultWriter(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web get_WritableStreamDefaultWriter_Closed
//go:noescape

func GetWritableStreamDefaultWriterClosed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_WritableStreamDefaultWriter_DesiredSize
//go:noescape

func GetWritableStreamDefaultWriterDesiredSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_WritableStreamDefaultWriter_Ready
//go:noescape

func GetWritableStreamDefaultWriterReady(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Abort
//go:noescape

func CallWritableStreamDefaultWriterAbort(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Abort
//go:noescape

func WritableStreamDefaultWriterAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Abort1
//go:noescape

func CallWritableStreamDefaultWriterAbort1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Abort1
//go:noescape

func WritableStreamDefaultWriterAbort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Close
//go:noescape

func CallWritableStreamDefaultWriterClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Close
//go:noescape

func WritableStreamDefaultWriterCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_ReleaseLock
//go:noescape

func CallWritableStreamDefaultWriterReleaseLock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_ReleaseLock
//go:noescape

func WritableStreamDefaultWriterReleaseLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Write
//go:noescape

func CallWritableStreamDefaultWriterWrite(
	this js.Ref,
	ptr unsafe.Pointer,

	chunk js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Write
//go:noescape

func WritableStreamDefaultWriterWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Write1
//go:noescape

func CallWritableStreamDefaultWriterWrite1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Write1
//go:noescape

func WritableStreamDefaultWriterWrite1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_WritableStream_WritableStream
//go:noescape

func NewWritableStreamByWritableStream(
	underlyingSink js.Ref,
	strategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_WritableStream_WritableStream1
//go:noescape

func NewWritableStreamByWritableStream1(
	underlyingSink js.Ref) js.Ref

//go:wasmimport plat/js/web new_WritableStream_WritableStream2
//go:noescape

func NewWritableStreamByWritableStream2() js.Ref

//go:wasmimport plat/js/web get_WritableStream_Locked
//go:noescape

func GetWritableStreamLocked(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_WritableStream_Abort
//go:noescape

func CallWritableStreamAbort(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Abort
//go:noescape

func WritableStreamAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStream_Abort1
//go:noescape

func CallWritableStreamAbort1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Abort1
//go:noescape

func WritableStreamAbort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStream_Close
//go:noescape

func CallWritableStreamClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Close
//go:noescape

func WritableStreamCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WritableStream_GetWriter
//go:noescape

func CallWritableStreamGetWriter(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_WritableStream_GetWriter
//go:noescape

func WritableStreamGetWriterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ReadableWritablePair
//go:noescape

func ReadableWritablePairJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ReadableWritablePair
//go:noescape

func ReadableWritablePairJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_StreamPipeOptions
//go:noescape

func StreamPipeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StreamPipeOptions
//go:noescape

func StreamPipeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ReadableStream_ReadableStream
//go:noescape

func NewReadableStreamByReadableStream(
	underlyingSource js.Ref,
	strategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ReadableStream_ReadableStream1
//go:noescape

func NewReadableStreamByReadableStream1(
	underlyingSource js.Ref) js.Ref

//go:wasmimport plat/js/web new_ReadableStream_ReadableStream2
//go:noescape

func NewReadableStreamByReadableStream2() js.Ref

//go:wasmimport plat/js/web get_ReadableStream_Locked
//go:noescape

func GetReadableStreamLocked(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_From
//go:noescape

func CallReadableStreamFrom(
	this js.Ref,
	ptr unsafe.Pointer,

	asyncIterable js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_From
//go:noescape

func ReadableStreamFromFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_Cancel
//go:noescape

func CallReadableStreamCancel(
	this js.Ref,
	ptr unsafe.Pointer,

	reason js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Cancel
//go:noescape

func ReadableStreamCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_Cancel1
//go:noescape

func CallReadableStreamCancel1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Cancel1
//go:noescape

func ReadableStreamCancel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_GetReader
//go:noescape

func CallReadableStreamGetReader(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_GetReader
//go:noescape

func ReadableStreamGetReaderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_GetReader1
//go:noescape

func CallReadableStreamGetReader1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_GetReader1
//go:noescape

func ReadableStreamGetReader1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_PipeThrough
//go:noescape

func CallReadableStreamPipeThrough(
	this js.Ref,
	ptr unsafe.Pointer,

	transform unsafe.Pointer,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeThrough
//go:noescape

func ReadableStreamPipeThroughFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_PipeThrough1
//go:noescape

func CallReadableStreamPipeThrough1(
	this js.Ref,
	ptr unsafe.Pointer,

	transform unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeThrough1
//go:noescape

func ReadableStreamPipeThrough1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_PipeTo
//go:noescape

func CallReadableStreamPipeTo(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeTo
//go:noescape

func ReadableStreamPipeToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_PipeTo1
//go:noescape

func CallReadableStreamPipeTo1(
	this js.Ref,
	ptr unsafe.Pointer,

	destination js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeTo1
//go:noescape

func ReadableStreamPipeTo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ReadableStream_Tee
//go:noescape

func CallReadableStreamTee(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Tee
//go:noescape

func ReadableStreamTeeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_Blob_Blob
//go:noescape

func NewBlobByBlob(
	blobParts js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Blob_Blob1
//go:noescape

func NewBlobByBlob1(
	blobParts js.Ref) js.Ref

//go:wasmimport plat/js/web new_Blob_Blob2
//go:noescape

func NewBlobByBlob2() js.Ref

//go:wasmimport plat/js/web get_Blob_Size
//go:noescape

func GetBlobSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web get_Blob_Type
//go:noescape

func GetBlobType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Blob_Slice
//go:noescape

func CallBlobSlice(
	this js.Ref,
	ptr unsafe.Pointer,

	start float64,
	end float64,
	contentType js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice
//go:noescape

func BlobSliceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_Slice1
//go:noescape

func CallBlobSlice1(
	this js.Ref,
	ptr unsafe.Pointer,

	start float64,
	end float64,
) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice1
//go:noescape

func BlobSlice1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_Slice2
//go:noescape

func CallBlobSlice2(
	this js.Ref,
	ptr unsafe.Pointer,

	start float64,
) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice2
//go:noescape

func BlobSlice2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_Slice3
//go:noescape

func CallBlobSlice3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice3
//go:noescape

func BlobSlice3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_Stream
//go:noescape

func CallBlobStream(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Blob_Stream
//go:noescape

func BlobStreamFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_Text
//go:noescape

func CallBlobText(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Blob_Text
//go:noescape

func BlobTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Blob_ArrayBuffer
//go:noescape

func CallBlobArrayBuffer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Blob_ArrayBuffer
//go:noescape

func BlobArrayBufferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_FilePropertyBag
//go:noescape

func FilePropertyBagJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FilePropertyBag
//go:noescape

func FilePropertyBagJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_File_File
//go:noescape

func NewFileByFile(
	fileBits js.Ref,
	fileName js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_File_File1
//go:noescape

func NewFileByFile1(
	fileBits js.Ref,
	fileName js.Ref) js.Ref

//go:wasmimport plat/js/web get_File_Name
//go:noescape

func GetFileName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_File_LastModified
//go:noescape

func GetFileLastModified(
	this js.Ref,
	ptr unsafe.Pointer,
) int64

//go:wasmimport plat/js/web get_File_WebkitRelativePath
//go:noescape

func GetFileWebkitRelativePath(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RadioNodeList_Value
//go:noescape

func GetRadioNodeListValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RadioNodeList_Value
//go:noescape

func SetRadioNodeListValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_HTMLFormControlsCollection_NamedItem
//go:noescape

func CallHTMLFormControlsCollectionNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLFormControlsCollection_NamedItem
//go:noescape

func HTMLFormControlsCollectionNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_HTMLFormElement_HTMLFormElement
//go:noescape

func NewHTMLFormElementByHTMLFormElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_AcceptCharset
//go:noescape

func GetHTMLFormElementAcceptCharset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_AcceptCharset
//go:noescape

func SetHTMLFormElementAcceptCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Action
//go:noescape

func GetHTMLFormElementAction(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Action
//go:noescape

func SetHTMLFormElementAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Autocomplete
//go:noescape

func GetHTMLFormElementAutocomplete(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Autocomplete
//go:noescape

func SetHTMLFormElementAutocomplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Enctype
//go:noescape

func GetHTMLFormElementEnctype(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Enctype
//go:noescape

func SetHTMLFormElementEnctype(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Encoding
//go:noescape

func GetHTMLFormElementEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Encoding
//go:noescape

func SetHTMLFormElementEncoding(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Method
//go:noescape

func GetHTMLFormElementMethod(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Method
//go:noescape

func SetHTMLFormElementMethod(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Name
//go:noescape

func GetHTMLFormElementName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Name
//go:noescape

func SetHTMLFormElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_NoValidate
//go:noescape

func GetHTMLFormElementNoValidate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_NoValidate
//go:noescape

func SetHTMLFormElementNoValidate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Target
//go:noescape

func GetHTMLFormElementTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Target
//go:noescape

func SetHTMLFormElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Rel
//go:noescape

func GetHTMLFormElementRel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLFormElement_Rel
//go:noescape

func SetHTMLFormElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_RelList
//go:noescape

func GetHTMLFormElementRelList(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Elements
//go:noescape

func GetHTMLFormElementElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Length
//go:noescape

func GetHTMLFormElementLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_HTMLFormElement_Get
//go:noescape

func CallHTMLFormElementGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Get
//go:noescape

func HTMLFormElementGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_Get1
//go:noescape

func CallHTMLFormElementGet1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Get1
//go:noescape

func HTMLFormElementGet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_Submit
//go:noescape

func CallHTMLFormElementSubmit(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Submit
//go:noescape

func HTMLFormElementSubmitFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_RequestSubmit
//go:noescape

func CallHTMLFormElementRequestSubmit(
	this js.Ref,
	ptr unsafe.Pointer,

	submitter js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_RequestSubmit
//go:noescape

func HTMLFormElementRequestSubmitFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_RequestSubmit1
//go:noescape

func CallHTMLFormElementRequestSubmit1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_RequestSubmit1
//go:noescape

func HTMLFormElementRequestSubmit1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_Reset
//go:noescape

func CallHTMLFormElementReset(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Reset
//go:noescape

func HTMLFormElementResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_CheckValidity
//go:noescape

func CallHTMLFormElementCheckValidity(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_CheckValidity
//go:noescape

func HTMLFormElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFormElement_ReportValidity
//go:noescape

func CallHTMLFormElementReportValidity(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_ReportValidity
//go:noescape

func HTMLFormElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_FormData_FormData
//go:noescape

func NewFormDataByFormData(
	form js.Ref,
	submitter js.Ref) js.Ref

//go:wasmimport plat/js/web new_FormData_FormData1
//go:noescape

func NewFormDataByFormData1(
	form js.Ref) js.Ref

//go:wasmimport plat/js/web new_FormData_FormData2
//go:noescape

func NewFormDataByFormData2() js.Ref

//go:wasmimport plat/js/web call_FormData_Append
//go:noescape

func CallFormDataAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Append
//go:noescape

func FormDataAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Append1
//go:noescape

func CallFormDataAppend1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	blobValue js.Ref,
	filename js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Append1
//go:noescape

func FormDataAppend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Append2
//go:noescape

func CallFormDataAppend2(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	blobValue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Append2
//go:noescape

func FormDataAppend2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Delete
//go:noescape

func CallFormDataDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Delete
//go:noescape

func FormDataDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Get
//go:noescape

func CallFormDataGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Get
//go:noescape

func FormDataGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_GetAll
//go:noescape

func CallFormDataGetAll(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_GetAll
//go:noescape

func FormDataGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Has
//go:noescape

func CallFormDataHas(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Has
//go:noescape

func FormDataHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Set
//go:noescape

func CallFormDataSet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Set
//go:noescape

func FormDataSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Set1
//go:noescape

func CallFormDataSet1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	blobValue js.Ref,
	filename js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Set1
//go:noescape

func FormDataSet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FormData_Set2
//go:noescape

func CallFormDataSet2(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	blobValue js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FormData_Set2
//go:noescape

func FormDataSet2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ValidityStateFlags
//go:noescape

func ValidityStateFlagsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ValidityStateFlags
//go:noescape

func ValidityStateFlagsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ValidityState_ValueMissing
//go:noescape

func GetValidityStateValueMissing(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_TypeMismatch
//go:noescape

func GetValidityStateTypeMismatch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_PatternMismatch
//go:noescape

func GetValidityStatePatternMismatch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_TooLong
//go:noescape

func GetValidityStateTooLong(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_TooShort
//go:noescape

func GetValidityStateTooShort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_RangeUnderflow
//go:noescape

func GetValidityStateRangeUnderflow(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_RangeOverflow
//go:noescape

func GetValidityStateRangeOverflow(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_StepMismatch
//go:noescape

func GetValidityStateStepMismatch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_BadInput
//go:noescape

func GetValidityStateBadInput(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_CustomError
//go:noescape

func GetValidityStateCustomError(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ValidityState_Valid
//go:noescape

func GetValidityStateValid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_CustomStateSet_Add
//go:noescape

func CallCustomStateSetAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CustomStateSet_Add
//go:noescape

func CustomStateSetAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_ShadowRoot
//go:noescape

func GetElementInternalsShadowRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_Form
//go:noescape

func GetElementInternalsForm(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_WillValidate
//go:noescape

func GetElementInternalsWillValidate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_Validity
//go:noescape

func GetElementInternalsValidity(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_ValidationMessage
//go:noescape

func GetElementInternalsValidationMessage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_Labels
//go:noescape

func GetElementInternalsLabels(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_States
//go:noescape

func GetElementInternalsStates(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_Role
//go:noescape

func GetElementInternalsRole(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_Role
//go:noescape

func SetElementInternalsRole(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaActiveDescendantElement
//go:noescape

func GetElementInternalsAriaActiveDescendantElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaActiveDescendantElement
//go:noescape

func SetElementInternalsAriaActiveDescendantElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaAtomic
//go:noescape

func GetElementInternalsAriaAtomic(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaAtomic
//go:noescape

func SetElementInternalsAriaAtomic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaAutoComplete
//go:noescape

func GetElementInternalsAriaAutoComplete(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaAutoComplete
//go:noescape

func SetElementInternalsAriaAutoComplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaBusy
//go:noescape

func GetElementInternalsAriaBusy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaBusy
//go:noescape

func SetElementInternalsAriaBusy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaChecked
//go:noescape

func GetElementInternalsAriaChecked(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaChecked
//go:noescape

func SetElementInternalsAriaChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColCount
//go:noescape

func GetElementInternalsAriaColCount(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaColCount
//go:noescape

func SetElementInternalsAriaColCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColIndex
//go:noescape

func GetElementInternalsAriaColIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaColIndex
//go:noescape

func SetElementInternalsAriaColIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColIndexText
//go:noescape

func GetElementInternalsAriaColIndexText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaColIndexText
//go:noescape

func SetElementInternalsAriaColIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColSpan
//go:noescape

func GetElementInternalsAriaColSpan(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaColSpan
//go:noescape

func SetElementInternalsAriaColSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaControlsElements
//go:noescape

func GetElementInternalsAriaControlsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaControlsElements
//go:noescape

func SetElementInternalsAriaControlsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaCurrent
//go:noescape

func GetElementInternalsAriaCurrent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaCurrent
//go:noescape

func SetElementInternalsAriaCurrent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDescribedByElements
//go:noescape

func GetElementInternalsAriaDescribedByElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaDescribedByElements
//go:noescape

func SetElementInternalsAriaDescribedByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDescription
//go:noescape

func GetElementInternalsAriaDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaDescription
//go:noescape

func SetElementInternalsAriaDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDetailsElements
//go:noescape

func GetElementInternalsAriaDetailsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaDetailsElements
//go:noescape

func SetElementInternalsAriaDetailsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDisabled
//go:noescape

func GetElementInternalsAriaDisabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaDisabled
//go:noescape

func SetElementInternalsAriaDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaErrorMessageElements
//go:noescape

func GetElementInternalsAriaErrorMessageElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaErrorMessageElements
//go:noescape

func SetElementInternalsAriaErrorMessageElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaExpanded
//go:noescape

func GetElementInternalsAriaExpanded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaExpanded
//go:noescape

func SetElementInternalsAriaExpanded(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaFlowToElements
//go:noescape

func GetElementInternalsAriaFlowToElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaFlowToElements
//go:noescape

func SetElementInternalsAriaFlowToElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaHasPopup
//go:noescape

func GetElementInternalsAriaHasPopup(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaHasPopup
//go:noescape

func SetElementInternalsAriaHasPopup(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaHidden
//go:noescape

func GetElementInternalsAriaHidden(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaHidden
//go:noescape

func SetElementInternalsAriaHidden(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaInvalid
//go:noescape

func GetElementInternalsAriaInvalid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaInvalid
//go:noescape

func SetElementInternalsAriaInvalid(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaKeyShortcuts
//go:noescape

func GetElementInternalsAriaKeyShortcuts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaKeyShortcuts
//go:noescape

func SetElementInternalsAriaKeyShortcuts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLabel
//go:noescape

func GetElementInternalsAriaLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaLabel
//go:noescape

func SetElementInternalsAriaLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLabelledByElements
//go:noescape

func GetElementInternalsAriaLabelledByElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaLabelledByElements
//go:noescape

func SetElementInternalsAriaLabelledByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLevel
//go:noescape

func GetElementInternalsAriaLevel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaLevel
//go:noescape

func SetElementInternalsAriaLevel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLive
//go:noescape

func GetElementInternalsAriaLive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaLive
//go:noescape

func SetElementInternalsAriaLive(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaModal
//go:noescape

func GetElementInternalsAriaModal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaModal
//go:noescape

func SetElementInternalsAriaModal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaMultiLine
//go:noescape

func GetElementInternalsAriaMultiLine(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaMultiLine
//go:noescape

func SetElementInternalsAriaMultiLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaMultiSelectable
//go:noescape

func GetElementInternalsAriaMultiSelectable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaMultiSelectable
//go:noescape

func SetElementInternalsAriaMultiSelectable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaOrientation
//go:noescape

func GetElementInternalsAriaOrientation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaOrientation
//go:noescape

func SetElementInternalsAriaOrientation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaOwnsElements
//go:noescape

func GetElementInternalsAriaOwnsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaOwnsElements
//go:noescape

func SetElementInternalsAriaOwnsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPlaceholder
//go:noescape

func GetElementInternalsAriaPlaceholder(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaPlaceholder
//go:noescape

func SetElementInternalsAriaPlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPosInSet
//go:noescape

func GetElementInternalsAriaPosInSet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaPosInSet
//go:noescape

func SetElementInternalsAriaPosInSet(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPressed
//go:noescape

func GetElementInternalsAriaPressed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaPressed
//go:noescape

func SetElementInternalsAriaPressed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaReadOnly
//go:noescape

func GetElementInternalsAriaReadOnly(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaReadOnly
//go:noescape

func SetElementInternalsAriaReadOnly(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRequired
//go:noescape

func GetElementInternalsAriaRequired(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRequired
//go:noescape

func SetElementInternalsAriaRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRoleDescription
//go:noescape

func GetElementInternalsAriaRoleDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRoleDescription
//go:noescape

func SetElementInternalsAriaRoleDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowCount
//go:noescape

func GetElementInternalsAriaRowCount(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRowCount
//go:noescape

func SetElementInternalsAriaRowCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowIndex
//go:noescape

func GetElementInternalsAriaRowIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRowIndex
//go:noescape

func SetElementInternalsAriaRowIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowIndexText
//go:noescape

func GetElementInternalsAriaRowIndexText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRowIndexText
//go:noescape

func SetElementInternalsAriaRowIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowSpan
//go:noescape

func GetElementInternalsAriaRowSpan(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaRowSpan
//go:noescape

func SetElementInternalsAriaRowSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSelected
//go:noescape

func GetElementInternalsAriaSelected(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaSelected
//go:noescape

func SetElementInternalsAriaSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSetSize
//go:noescape

func GetElementInternalsAriaSetSize(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaSetSize
//go:noescape

func SetElementInternalsAriaSetSize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSort
//go:noescape

func GetElementInternalsAriaSort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaSort
//go:noescape

func SetElementInternalsAriaSort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueMax
//go:noescape

func GetElementInternalsAriaValueMax(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaValueMax
//go:noescape

func SetElementInternalsAriaValueMax(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueMin
//go:noescape

func GetElementInternalsAriaValueMin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaValueMin
//go:noescape

func SetElementInternalsAriaValueMin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueNow
//go:noescape

func GetElementInternalsAriaValueNow(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaValueNow
//go:noescape

func SetElementInternalsAriaValueNow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueText
//go:noescape

func GetElementInternalsAriaValueText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ElementInternals_AriaValueText
//go:noescape

func SetElementInternalsAriaValueText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetFormValue
//go:noescape

func CallElementInternalsSetFormValue(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
	state js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetFormValue
//go:noescape

func ElementInternalsSetFormValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetFormValue1
//go:noescape

func CallElementInternalsSetFormValue1(
	this js.Ref,
	ptr unsafe.Pointer,

	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetFormValue1
//go:noescape

func ElementInternalsSetFormValue1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetValidity
//go:noescape

func CallElementInternalsSetValidity(
	this js.Ref,
	ptr unsafe.Pointer,

	flags unsafe.Pointer,
	message js.Ref,
	anchor js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity
//go:noescape

func ElementInternalsSetValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetValidity1
//go:noescape

func CallElementInternalsSetValidity1(
	this js.Ref,
	ptr unsafe.Pointer,

	flags unsafe.Pointer,
	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity1
//go:noescape

func ElementInternalsSetValidity1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetValidity2
//go:noescape

func CallElementInternalsSetValidity2(
	this js.Ref,
	ptr unsafe.Pointer,

	flags unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity2
//go:noescape

func ElementInternalsSetValidity2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_SetValidity3
//go:noescape

func CallElementInternalsSetValidity3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity3
//go:noescape

func ElementInternalsSetValidity3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_CheckValidity
//go:noescape

func CallElementInternalsCheckValidity(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_CheckValidity
//go:noescape

func ElementInternalsCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ElementInternals_ReportValidity
//go:noescape

func CallElementInternalsReportValidity(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_ReportValidity
//go:noescape

func ElementInternalsReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_EditContextInit
//go:noescape

func EditContextInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_EditContextInit
//go:noescape

func EditContextInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_EditContext_EditContext
//go:noescape

func NewEditContextByEditContext(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_EditContext_EditContext1
//go:noescape

func NewEditContextByEditContext1() js.Ref

//go:wasmimport plat/js/web get_EditContext_Text
//go:noescape

func GetEditContextText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EditContext_SelectionStart
//go:noescape

func GetEditContextSelectionStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_EditContext_SelectionEnd
//go:noescape

func GetEditContextSelectionEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_EditContext_CompositionRangeStart
//go:noescape

func GetEditContextCompositionRangeStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_EditContext_CompositionRangeEnd
//go:noescape

func GetEditContextCompositionRangeEnd(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_EditContext_IsComposing
//go:noescape

func GetEditContextIsComposing(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EditContext_ControlBounds
//go:noescape

func GetEditContextControlBounds(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EditContext_SelectionBounds
//go:noescape

func GetEditContextSelectionBounds(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_EditContext_CharacterBoundsRangeStart
//go:noescape

func GetEditContextCharacterBoundsRangeStart(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_EditContext_UpdateText
//go:noescape

func CallEditContextUpdateText(
	this js.Ref,
	ptr unsafe.Pointer,

	rangeStart uint32,
	rangeEnd uint32,
	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateText
//go:noescape

func EditContextUpdateTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_UpdateSelection
//go:noescape

func CallEditContextUpdateSelection(
	this js.Ref,
	ptr unsafe.Pointer,

	start uint32,
	end uint32,
) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateSelection
//go:noescape

func EditContextUpdateSelectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_UpdateControlBounds
//go:noescape

func CallEditContextUpdateControlBounds(
	this js.Ref,
	ptr unsafe.Pointer,

	controlBounds js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateControlBounds
//go:noescape

func EditContextUpdateControlBoundsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_UpdateSelectionBounds
//go:noescape

func CallEditContextUpdateSelectionBounds(
	this js.Ref,
	ptr unsafe.Pointer,

	selectionBounds js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateSelectionBounds
//go:noescape

func EditContextUpdateSelectionBoundsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_UpdateCharacterBounds
//go:noescape

func CallEditContextUpdateCharacterBounds(
	this js.Ref,
	ptr unsafe.Pointer,

	rangeStart uint32,
	characterBounds js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateCharacterBounds
//go:noescape

func EditContextUpdateCharacterBoundsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_AttachedElements
//go:noescape

func CallEditContextAttachedElements(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_EditContext_AttachedElements
//go:noescape

func EditContextAttachedElementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_EditContext_CharacterBounds
//go:noescape

func CallEditContextCharacterBounds(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_EditContext_CharacterBounds
//go:noescape

func EditContextCharacterBoundsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_HTMLElement_HTMLElement
//go:noescape

func NewHTMLElementByHTMLElement() js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Title
//go:noescape

func GetHTMLElementTitle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Title
//go:noescape

func SetHTMLElementTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Lang
//go:noescape

func GetHTMLElementLang(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Lang
//go:noescape

func SetHTMLElementLang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Translate
//go:noescape

func GetHTMLElementTranslate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Translate
//go:noescape

func SetHTMLElementTranslate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Dir
//go:noescape

func GetHTMLElementDir(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Dir
//go:noescape

func SetHTMLElementDir(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Hidden
//go:noescape

func GetHTMLElementHidden(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Hidden
//go:noescape

func SetHTMLElementHidden(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Inert
//go:noescape

func GetHTMLElementInert(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Inert
//go:noescape

func SetHTMLElementInert(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_AccessKey
//go:noescape

func GetHTMLElementAccessKey(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_AccessKey
//go:noescape

func SetHTMLElementAccessKey(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_AccessKeyLabel
//go:noescape

func GetHTMLElementAccessKeyLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Draggable
//go:noescape

func GetHTMLElementDraggable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Draggable
//go:noescape

func SetHTMLElementDraggable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Spellcheck
//go:noescape

func GetHTMLElementSpellcheck(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Spellcheck
//go:noescape

func SetHTMLElementSpellcheck(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Autocapitalize
//go:noescape

func GetHTMLElementAutocapitalize(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Autocapitalize
//go:noescape

func SetHTMLElementAutocapitalize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_InnerText
//go:noescape

func GetHTMLElementInnerText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_InnerText
//go:noescape

func SetHTMLElementInnerText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_OuterText
//go:noescape

func GetHTMLElementOuterText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_OuterText
//go:noescape

func SetHTMLElementOuterText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Popover
//go:noescape

func GetHTMLElementPopover(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Popover
//go:noescape

func SetHTMLElementPopover(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_OffsetParent
//go:noescape

func GetHTMLElementOffsetParent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_OffsetTop
//go:noescape

func GetHTMLElementOffsetTop(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_HTMLElement_OffsetLeft
//go:noescape

func GetHTMLElementOffsetLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_HTMLElement_OffsetWidth
//go:noescape

func GetHTMLElementOffsetWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_HTMLElement_OffsetHeight
//go:noescape

func GetHTMLElementOffsetHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_HTMLElement_EditContext
//go:noescape

func GetHTMLElementEditContext(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_EditContext
//go:noescape

func SetHTMLElementEditContext(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Style
//go:noescape

func GetHTMLElementStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_AttributeStyleMap
//go:noescape

func GetHTMLElementAttributeStyleMap(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Dataset
//go:noescape

func GetHTMLElementDataset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Nonce
//go:noescape

func GetHTMLElementNonce(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Nonce
//go:noescape

func SetHTMLElementNonce(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Autofocus
//go:noescape

func GetHTMLElementAutofocus(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_Autofocus
//go:noescape

func SetHTMLElementAutofocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_TabIndex
//go:noescape

func GetHTMLElementTabIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web set_HTMLElement_TabIndex
//go:noescape

func SetHTMLElementTabIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_ContentEditable
//go:noescape

func GetHTMLElementContentEditable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_ContentEditable
//go:noescape

func SetHTMLElementContentEditable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_EnterKeyHint
//go:noescape

func GetHTMLElementEnterKeyHint(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_EnterKeyHint
//go:noescape

func SetHTMLElementEnterKeyHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_IsContentEditable
//go:noescape

func GetHTMLElementIsContentEditable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_InputMode
//go:noescape

func GetHTMLElementInputMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_InputMode
//go:noescape

func SetHTMLElementInputMode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_VirtualKeyboardPolicy
//go:noescape

func GetHTMLElementVirtualKeyboardPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLElement_VirtualKeyboardPolicy
//go:noescape

func SetHTMLElementVirtualKeyboardPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_Click
//go:noescape

func CallHTMLElementClick(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Click
//go:noescape

func HTMLElementClickFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_AttachInternals
//go:noescape

func CallHTMLElementAttachInternals(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_AttachInternals
//go:noescape

func HTMLElementAttachInternalsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_ShowPopover
//go:noescape

func CallHTMLElementShowPopover(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_ShowPopover
//go:noescape

func HTMLElementShowPopoverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_HidePopover
//go:noescape

func CallHTMLElementHidePopover(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_HidePopover
//go:noescape

func HTMLElementHidePopoverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_TogglePopover
//go:noescape

func CallHTMLElementTogglePopover(
	this js.Ref,
	ptr unsafe.Pointer,

	force js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_TogglePopover
//go:noescape

func HTMLElementTogglePopoverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_TogglePopover1
//go:noescape

func CallHTMLElementTogglePopover1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_TogglePopover1
//go:noescape

func HTMLElementTogglePopover1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_Focus
//go:noescape

func CallHTMLElementFocus(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Focus
//go:noescape

func HTMLElementFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_Focus1
//go:noescape

func CallHTMLElementFocus1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Focus1
//go:noescape

func HTMLElementFocus1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLElement_Blur
//go:noescape

func CallHTMLElementBlur(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Blur
//go:noescape

func HTMLElementBlurFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_HTMLHeadElement_HTMLHeadElement
//go:noescape

func NewHTMLHeadElementByHTMLHeadElement() js.Ref

//go:wasmimport plat/js/web new_HTMLScriptElement_HTMLScriptElement
//go:noescape

func NewHTMLScriptElementByHTMLScriptElement() js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Src
//go:noescape

func GetHTMLScriptElementSrc(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Src
//go:noescape

func SetHTMLScriptElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Type
//go:noescape

func GetHTMLScriptElementType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Type
//go:noescape

func SetHTMLScriptElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_NoModule
//go:noescape

func GetHTMLScriptElementNoModule(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_NoModule
//go:noescape

func SetHTMLScriptElementNoModule(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Async
//go:noescape

func GetHTMLScriptElementAsync(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Async
//go:noescape

func SetHTMLScriptElementAsync(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Defer
//go:noescape

func GetHTMLScriptElementDefer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Defer
//go:noescape

func SetHTMLScriptElementDefer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_CrossOrigin
//go:noescape

func GetHTMLScriptElementCrossOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_CrossOrigin
//go:noescape

func SetHTMLScriptElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Text
//go:noescape

func GetHTMLScriptElementText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Text
//go:noescape

func SetHTMLScriptElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Integrity
//go:noescape

func GetHTMLScriptElementIntegrity(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Integrity
//go:noescape

func SetHTMLScriptElementIntegrity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_ReferrerPolicy
//go:noescape

func GetHTMLScriptElementReferrerPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_ReferrerPolicy
//go:noescape

func SetHTMLScriptElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Blocking
//go:noescape

func GetHTMLScriptElementBlocking(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_FetchPriority
//go:noescape

func GetHTMLScriptElementFetchPriority(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_FetchPriority
//go:noescape

func SetHTMLScriptElementFetchPriority(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Charset
//go:noescape

func GetHTMLScriptElementCharset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Charset
//go:noescape

func SetHTMLScriptElementCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Event
//go:noescape

func GetHTMLScriptElementEvent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_Event
//go:noescape

func SetHTMLScriptElementEvent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_HtmlFor
//go:noescape

func GetHTMLScriptElementHtmlFor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_HtmlFor
//go:noescape

func SetHTMLScriptElementHtmlFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_AttributionSrc
//go:noescape

func GetHTMLScriptElementAttributionSrc(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLScriptElement_AttributionSrc
//go:noescape

func SetHTMLScriptElementAttributionSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_HTMLScriptElement_Supports
//go:noescape

func CallHTMLScriptElementSupports(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLScriptElement_Supports
//go:noescape

func HTMLScriptElementSupportsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGScriptElement_Type
//go:noescape

func GetSVGScriptElementType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGScriptElement_Type
//go:noescape

func SetSVGScriptElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGScriptElement_CrossOrigin
//go:noescape

func GetSVGScriptElementCrossOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGScriptElement_CrossOrigin
//go:noescape

func SetSVGScriptElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGScriptElement_Href
//go:noescape

func GetSVGScriptElementHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref
