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

//go:wasmimport plat/js/web get_HTMLAllCollection_Length
//go:noescape
func GetHTMLAllCollectionLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLAllCollection_Get
//go:noescape
func HasFuncHTMLAllCollectionGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Get
//go:noescape
func FuncHTMLAllCollectionGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLAllCollection_Get
//go:noescape
func CallHTMLAllCollectionGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_HTMLAllCollection_Get
//go:noescape
func TryHTMLAllCollectionGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLAllCollection_NamedItem
//go:noescape
func HasFuncHTMLAllCollectionNamedItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_NamedItem
//go:noescape
func FuncHTMLAllCollectionNamedItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLAllCollection_NamedItem
//go:noescape
func CallHTMLAllCollectionNamedItem(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_HTMLAllCollection_NamedItem
//go:noescape
func TryHTMLAllCollectionNamedItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLAllCollection_Item
//go:noescape
func HasFuncHTMLAllCollectionItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Item
//go:noescape
func FuncHTMLAllCollectionItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLAllCollection_Item
//go:noescape
func CallHTMLAllCollectionItem(
	this js.Ref, retPtr unsafe.Pointer,
	nameOrIndex js.Ref)

//go:wasmimport plat/js/web try_HTMLAllCollection_Item
//go:noescape
func TryHTMLAllCollectionItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nameOrIndex js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLAllCollection_Item1
//go:noescape
func HasFuncHTMLAllCollectionItem1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLAllCollection_Item1
//go:noescape
func FuncHTMLAllCollectionItem1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLAllCollection_Item1
//go:noescape
func CallHTMLAllCollectionItem1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLAllCollection_Item1
//go:noescape
func TryHTMLAllCollectionItem1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMStringList_Length
//go:noescape
func GetDOMStringListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMStringList_Item
//go:noescape
func HasFuncDOMStringListItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringList_Item
//go:noescape
func FuncDOMStringListItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMStringList_Item
//go:noescape
func CallDOMStringListItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_DOMStringList_Item
//go:noescape
func TryDOMStringListItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMStringList_Contains
//go:noescape
func HasFuncDOMStringListContains(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringList_Contains
//go:noescape
func FuncDOMStringListContains(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DOMStringList_Contains
//go:noescape
func CallDOMStringListContains(
	this js.Ref, retPtr unsafe.Pointer,
	string js.Ref)

//go:wasmimport plat/js/web try_DOMStringList_Contains
//go:noescape
func TryDOMStringListContains(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	string js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_Location_Href
//go:noescape
func GetLocationHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Href
//go:noescape
func SetLocationHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Origin
//go:noescape
func GetLocationOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Location_Protocol
//go:noescape
func GetLocationProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Protocol
//go:noescape
func SetLocationProtocol(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Host
//go:noescape
func GetLocationHost(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Host
//go:noescape
func SetLocationHost(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Hostname
//go:noescape
func GetLocationHostname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Hostname
//go:noescape
func SetLocationHostname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Port
//go:noescape
func GetLocationPort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Port
//go:noescape
func SetLocationPort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Pathname
//go:noescape
func GetLocationPathname(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Pathname
//go:noescape
func SetLocationPathname(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Search
//go:noescape
func GetLocationSearch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Search
//go:noescape
func SetLocationSearch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_Hash
//go:noescape
func GetLocationHash(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Location_Hash
//go:noescape
func SetLocationHash(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Location_AncestorOrigins
//go:noescape
func GetLocationAncestorOrigins(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Location_Assign
//go:noescape
func HasFuncLocationAssign(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Location_Assign
//go:noescape
func FuncLocationAssign(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Location_Assign
//go:noescape
func CallLocationAssign(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Location_Assign
//go:noescape
func TryLocationAssign(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Location_Replace
//go:noescape
func HasFuncLocationReplace(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Location_Replace
//go:noescape
func FuncLocationReplace(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Location_Replace
//go:noescape
func CallLocationReplace(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Location_Replace
//go:noescape
func TryLocationReplace(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Location_Reload
//go:noescape
func HasFuncLocationReload(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Location_Reload
//go:noescape
func FuncLocationReload(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Location_Reload
//go:noescape
func CallLocationReload(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Location_Reload
//go:noescape
func TryLocationReload(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultReader_Read
//go:noescape
func HasFuncReadableStreamDefaultReaderRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Read
//go:noescape
func FuncReadableStreamDefaultReaderRead(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Read
//go:noescape
func CallReadableStreamDefaultReaderRead(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultReader_Read
//go:noescape
func TryReadableStreamDefaultReaderRead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultReader_ReleaseLock
//go:noescape
func HasFuncReadableStreamDefaultReaderReleaseLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_ReleaseLock
//go:noescape
func FuncReadableStreamDefaultReaderReleaseLock(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_ReleaseLock
//go:noescape
func CallReadableStreamDefaultReaderReleaseLock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultReader_ReleaseLock
//go:noescape
func TryReadableStreamDefaultReaderReleaseLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultReader_Cancel
//go:noescape
func HasFuncReadableStreamDefaultReaderCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Cancel
//go:noescape
func FuncReadableStreamDefaultReaderCancel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Cancel
//go:noescape
func CallReadableStreamDefaultReaderCancel(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamDefaultReader_Cancel
//go:noescape
func TryReadableStreamDefaultReaderCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamDefaultReader_Cancel1
//go:noescape
func HasFuncReadableStreamDefaultReaderCancel1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamDefaultReader_Cancel1
//go:noescape
func FuncReadableStreamDefaultReaderCancel1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamDefaultReader_Cancel1
//go:noescape
func CallReadableStreamDefaultReaderCancel1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamDefaultReader_Cancel1
//go:noescape
func TryReadableStreamDefaultReaderCancel1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_ReadableStreamBYOBReader_ReadableStreamBYOBReader
//go:noescape
func NewReadableStreamBYOBReaderByReadableStreamBYOBReader(
	stream js.Ref) js.Ref

//go:wasmimport plat/js/web get_ReadableStreamBYOBReader_Closed
//go:noescape
func GetReadableStreamBYOBReaderClosed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBReader_Read
//go:noescape
func HasFuncReadableStreamBYOBReaderRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Read
//go:noescape
func FuncReadableStreamBYOBReaderRead(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Read
//go:noescape
func CallReadableStreamBYOBReaderRead(
	this js.Ref, retPtr unsafe.Pointer,
	view js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamBYOBReader_Read
//go:noescape
func TryReadableStreamBYOBReaderRead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBReader_ReleaseLock
//go:noescape
func HasFuncReadableStreamBYOBReaderReleaseLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_ReleaseLock
//go:noescape
func FuncReadableStreamBYOBReaderReleaseLock(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_ReleaseLock
//go:noescape
func CallReadableStreamBYOBReaderReleaseLock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamBYOBReader_ReleaseLock
//go:noescape
func TryReadableStreamBYOBReaderReleaseLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBReader_Cancel
//go:noescape
func HasFuncReadableStreamBYOBReaderCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Cancel
//go:noescape
func FuncReadableStreamBYOBReaderCancel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Cancel
//go:noescape
func CallReadableStreamBYOBReaderCancel(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_ReadableStreamBYOBReader_Cancel
//go:noescape
func TryReadableStreamBYOBReaderCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStreamBYOBReader_Cancel1
//go:noescape
func HasFuncReadableStreamBYOBReaderCancel1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStreamBYOBReader_Cancel1
//go:noescape
func FuncReadableStreamBYOBReaderCancel1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStreamBYOBReader_Cancel1
//go:noescape
func CallReadableStreamBYOBReaderCancel1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStreamBYOBReader_Cancel1
//go:noescape
func TryReadableStreamBYOBReaderCancel1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WritableStreamDefaultWriter_DesiredSize
//go:noescape
func GetWritableStreamDefaultWriterDesiredSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WritableStreamDefaultWriter_Ready
//go:noescape
func GetWritableStreamDefaultWriterReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_Abort
//go:noescape
func HasFuncWritableStreamDefaultWriterAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Abort
//go:noescape
func FuncWritableStreamDefaultWriterAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Abort
//go:noescape
func CallWritableStreamDefaultWriterAbort(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_Abort
//go:noescape
func TryWritableStreamDefaultWriterAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_Abort1
//go:noescape
func HasFuncWritableStreamDefaultWriterAbort1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Abort1
//go:noescape
func FuncWritableStreamDefaultWriterAbort1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Abort1
//go:noescape
func CallWritableStreamDefaultWriterAbort1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_Abort1
//go:noescape
func TryWritableStreamDefaultWriterAbort1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_Close
//go:noescape
func HasFuncWritableStreamDefaultWriterClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Close
//go:noescape
func FuncWritableStreamDefaultWriterClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Close
//go:noescape
func CallWritableStreamDefaultWriterClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_Close
//go:noescape
func TryWritableStreamDefaultWriterClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_ReleaseLock
//go:noescape
func HasFuncWritableStreamDefaultWriterReleaseLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_ReleaseLock
//go:noescape
func FuncWritableStreamDefaultWriterReleaseLock(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_ReleaseLock
//go:noescape
func CallWritableStreamDefaultWriterReleaseLock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_ReleaseLock
//go:noescape
func TryWritableStreamDefaultWriterReleaseLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_Write
//go:noescape
func HasFuncWritableStreamDefaultWriterWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Write
//go:noescape
func FuncWritableStreamDefaultWriterWrite(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Write
//go:noescape
func CallWritableStreamDefaultWriterWrite(
	this js.Ref, retPtr unsafe.Pointer,
	chunk js.Ref)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_Write
//go:noescape
func TryWritableStreamDefaultWriterWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	chunk js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStreamDefaultWriter_Write1
//go:noescape
func HasFuncWritableStreamDefaultWriterWrite1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStreamDefaultWriter_Write1
//go:noescape
func FuncWritableStreamDefaultWriterWrite1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStreamDefaultWriter_Write1
//go:noescape
func CallWritableStreamDefaultWriterWrite1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStreamDefaultWriter_Write1
//go:noescape
func TryWritableStreamDefaultWriterWrite1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStream_Abort
//go:noescape
func HasFuncWritableStreamAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Abort
//go:noescape
func FuncWritableStreamAbort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStream_Abort
//go:noescape
func CallWritableStreamAbort(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_WritableStream_Abort
//go:noescape
func TryWritableStreamAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStream_Abort1
//go:noescape
func HasFuncWritableStreamAbort1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Abort1
//go:noescape
func FuncWritableStreamAbort1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStream_Abort1
//go:noescape
func CallWritableStreamAbort1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStream_Abort1
//go:noescape
func TryWritableStreamAbort1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStream_Close
//go:noescape
func HasFuncWritableStreamClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStream_Close
//go:noescape
func FuncWritableStreamClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStream_Close
//go:noescape
func CallWritableStreamClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStream_Close
//go:noescape
func TryWritableStreamClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WritableStream_GetWriter
//go:noescape
func HasFuncWritableStreamGetWriter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WritableStream_GetWriter
//go:noescape
func FuncWritableStreamGetWriter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WritableStream_GetWriter
//go:noescape
func CallWritableStreamGetWriter(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WritableStream_GetWriter
//go:noescape
func TryWritableStreamGetWriter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_From
//go:noescape
func HasFuncReadableStreamFrom(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_From
//go:noescape
func FuncReadableStreamFrom(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_From
//go:noescape
func CallReadableStreamFrom(
	this js.Ref, retPtr unsafe.Pointer,
	asyncIterable js.Ref)

//go:wasmimport plat/js/web try_ReadableStream_From
//go:noescape
func TryReadableStreamFrom(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	asyncIterable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_Cancel
//go:noescape
func HasFuncReadableStreamCancel(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Cancel
//go:noescape
func FuncReadableStreamCancel(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_Cancel
//go:noescape
func CallReadableStreamCancel(
	this js.Ref, retPtr unsafe.Pointer,
	reason js.Ref)

//go:wasmimport plat/js/web try_ReadableStream_Cancel
//go:noescape
func TryReadableStreamCancel(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reason js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_Cancel1
//go:noescape
func HasFuncReadableStreamCancel1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Cancel1
//go:noescape
func FuncReadableStreamCancel1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_Cancel1
//go:noescape
func CallReadableStreamCancel1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_Cancel1
//go:noescape
func TryReadableStreamCancel1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_GetReader
//go:noescape
func HasFuncReadableStreamGetReader(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_GetReader
//go:noescape
func FuncReadableStreamGetReader(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_GetReader
//go:noescape
func CallReadableStreamGetReader(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_GetReader
//go:noescape
func TryReadableStreamGetReader(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_GetReader1
//go:noescape
func HasFuncReadableStreamGetReader1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_GetReader1
//go:noescape
func FuncReadableStreamGetReader1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_GetReader1
//go:noescape
func CallReadableStreamGetReader1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_GetReader1
//go:noescape
func TryReadableStreamGetReader1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_PipeThrough
//go:noescape
func HasFuncReadableStreamPipeThrough(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeThrough
//go:noescape
func FuncReadableStreamPipeThrough(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_PipeThrough
//go:noescape
func CallReadableStreamPipeThrough(
	this js.Ref, retPtr unsafe.Pointer,
	transform unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_PipeThrough
//go:noescape
func TryReadableStreamPipeThrough(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transform unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_PipeThrough1
//go:noescape
func HasFuncReadableStreamPipeThrough1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeThrough1
//go:noescape
func FuncReadableStreamPipeThrough1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_PipeThrough1
//go:noescape
func CallReadableStreamPipeThrough1(
	this js.Ref, retPtr unsafe.Pointer,
	transform unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_PipeThrough1
//go:noescape
func TryReadableStreamPipeThrough1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transform unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_PipeTo
//go:noescape
func HasFuncReadableStreamPipeTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeTo
//go:noescape
func FuncReadableStreamPipeTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_PipeTo
//go:noescape
func CallReadableStreamPipeTo(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_PipeTo
//go:noescape
func TryReadableStreamPipeTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_PipeTo1
//go:noescape
func HasFuncReadableStreamPipeTo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_PipeTo1
//go:noescape
func FuncReadableStreamPipeTo1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_PipeTo1
//go:noescape
func CallReadableStreamPipeTo1(
	this js.Ref, retPtr unsafe.Pointer,
	destination js.Ref)

//go:wasmimport plat/js/web try_ReadableStream_PipeTo1
//go:noescape
func TryReadableStreamPipeTo1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	destination js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ReadableStream_Tee
//go:noescape
func HasFuncReadableStreamTee(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ReadableStream_Tee
//go:noescape
func FuncReadableStreamTee(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ReadableStream_Tee
//go:noescape
func CallReadableStreamTee(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ReadableStream_Tee
//go:noescape
func TryReadableStreamTee(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Blob_Type
//go:noescape
func GetBlobType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Slice
//go:noescape
func HasFuncBlobSlice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice
//go:noescape
func FuncBlobSlice(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Slice
//go:noescape
func CallBlobSlice(
	this js.Ref, retPtr unsafe.Pointer,
	start float64,
	end float64,
	contentType js.Ref)

//go:wasmimport plat/js/web try_Blob_Slice
//go:noescape
func TryBlobSlice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start float64,
	end float64,
	contentType js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Slice1
//go:noescape
func HasFuncBlobSlice1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice1
//go:noescape
func FuncBlobSlice1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Slice1
//go:noescape
func CallBlobSlice1(
	this js.Ref, retPtr unsafe.Pointer,
	start float64,
	end float64)

//go:wasmimport plat/js/web try_Blob_Slice1
//go:noescape
func TryBlobSlice1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start float64,
	end float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Slice2
//go:noescape
func HasFuncBlobSlice2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice2
//go:noescape
func FuncBlobSlice2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Slice2
//go:noescape
func CallBlobSlice2(
	this js.Ref, retPtr unsafe.Pointer,
	start float64)

//go:wasmimport plat/js/web try_Blob_Slice2
//go:noescape
func TryBlobSlice2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Slice3
//go:noescape
func HasFuncBlobSlice3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Slice3
//go:noescape
func FuncBlobSlice3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Slice3
//go:noescape
func CallBlobSlice3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Blob_Slice3
//go:noescape
func TryBlobSlice3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Stream
//go:noescape
func HasFuncBlobStream(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Stream
//go:noescape
func FuncBlobStream(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Stream
//go:noescape
func CallBlobStream(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Blob_Stream
//go:noescape
func TryBlobStream(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_Text
//go:noescape
func HasFuncBlobText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_Text
//go:noescape
func FuncBlobText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_Text
//go:noescape
func CallBlobText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Blob_Text
//go:noescape
func TryBlobText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Blob_ArrayBuffer
//go:noescape
func HasFuncBlobArrayBuffer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Blob_ArrayBuffer
//go:noescape
func FuncBlobArrayBuffer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Blob_ArrayBuffer
//go:noescape
func CallBlobArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Blob_ArrayBuffer
//go:noescape
func TryBlobArrayBuffer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_File_LastModified
//go:noescape
func GetFileLastModified(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_File_WebkitRelativePath
//go:noescape
func GetFileWebkitRelativePath(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RadioNodeList_Value
//go:noescape
func GetRadioNodeListValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RadioNodeList_Value
//go:noescape
func SetRadioNodeListValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLFormControlsCollection_NamedItem
//go:noescape
func HasFuncHTMLFormControlsCollectionNamedItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormControlsCollection_NamedItem
//go:noescape
func FuncHTMLFormControlsCollectionNamedItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormControlsCollection_NamedItem
//go:noescape
func CallHTMLFormControlsCollectionNamedItem(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_HTMLFormControlsCollection_NamedItem
//go:noescape
func TryHTMLFormControlsCollectionNamedItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFormElement_AcceptCharset
//go:noescape
func GetHTMLFormElementAcceptCharset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_AcceptCharset
//go:noescape
func SetHTMLFormElementAcceptCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Action
//go:noescape
func GetHTMLFormElementAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Action
//go:noescape
func SetHTMLFormElementAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Autocomplete
//go:noescape
func GetHTMLFormElementAutocomplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Autocomplete
//go:noescape
func SetHTMLFormElementAutocomplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Enctype
//go:noescape
func GetHTMLFormElementEnctype(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Enctype
//go:noescape
func SetHTMLFormElementEnctype(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Encoding
//go:noescape
func GetHTMLFormElementEncoding(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Encoding
//go:noescape
func SetHTMLFormElementEncoding(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Method
//go:noescape
func GetHTMLFormElementMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Method
//go:noescape
func SetHTMLFormElementMethod(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Name
//go:noescape
func GetHTMLFormElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Name
//go:noescape
func SetHTMLFormElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_NoValidate
//go:noescape
func GetHTMLFormElementNoValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_NoValidate
//go:noescape
func SetHTMLFormElementNoValidate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Target
//go:noescape
func GetHTMLFormElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Target
//go:noescape
func SetHTMLFormElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_Rel
//go:noescape
func GetHTMLFormElementRel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFormElement_Rel
//go:noescape
func SetHTMLFormElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFormElement_RelList
//go:noescape
func GetHTMLFormElementRelList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFormElement_Elements
//go:noescape
func GetHTMLFormElementElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFormElement_Length
//go:noescape
func GetHTMLFormElementLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_Get
//go:noescape
func HasFuncHTMLFormElementGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Get
//go:noescape
func FuncHTMLFormElementGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_Get
//go:noescape
func CallHTMLFormElementGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_HTMLFormElement_Get
//go:noescape
func TryHTMLFormElementGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_Get1
//go:noescape
func HasFuncHTMLFormElementGet1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Get1
//go:noescape
func FuncHTMLFormElementGet1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_Get1
//go:noescape
func CallHTMLFormElementGet1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_HTMLFormElement_Get1
//go:noescape
func TryHTMLFormElementGet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_Submit
//go:noescape
func HasFuncHTMLFormElementSubmit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Submit
//go:noescape
func FuncHTMLFormElementSubmit(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_Submit
//go:noescape
func CallHTMLFormElementSubmit(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFormElement_Submit
//go:noescape
func TryHTMLFormElementSubmit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_RequestSubmit
//go:noescape
func HasFuncHTMLFormElementRequestSubmit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_RequestSubmit
//go:noescape
func FuncHTMLFormElementRequestSubmit(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_RequestSubmit
//go:noescape
func CallHTMLFormElementRequestSubmit(
	this js.Ref, retPtr unsafe.Pointer,
	submitter js.Ref)

//go:wasmimport plat/js/web try_HTMLFormElement_RequestSubmit
//go:noescape
func TryHTMLFormElementRequestSubmit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	submitter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_RequestSubmit1
//go:noescape
func HasFuncHTMLFormElementRequestSubmit1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_RequestSubmit1
//go:noescape
func FuncHTMLFormElementRequestSubmit1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_RequestSubmit1
//go:noescape
func CallHTMLFormElementRequestSubmit1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFormElement_RequestSubmit1
//go:noescape
func TryHTMLFormElementRequestSubmit1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_Reset
//go:noescape
func HasFuncHTMLFormElementReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_Reset
//go:noescape
func FuncHTMLFormElementReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_Reset
//go:noescape
func CallHTMLFormElementReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFormElement_Reset
//go:noescape
func TryHTMLFormElementReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_CheckValidity
//go:noescape
func HasFuncHTMLFormElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_CheckValidity
//go:noescape
func FuncHTMLFormElementCheckValidity(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_CheckValidity
//go:noescape
func CallHTMLFormElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFormElement_CheckValidity
//go:noescape
func TryHTMLFormElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFormElement_ReportValidity
//go:noescape
func HasFuncHTMLFormElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFormElement_ReportValidity
//go:noescape
func FuncHTMLFormElementReportValidity(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLFormElement_ReportValidity
//go:noescape
func CallHTMLFormElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFormElement_ReportValidity
//go:noescape
func TryHTMLFormElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_FormData_Append
//go:noescape
func HasFuncFormDataAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Append
//go:noescape
func FuncFormDataAppend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Append
//go:noescape
func CallFormDataAppend(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_FormData_Append
//go:noescape
func TryFormDataAppend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Append1
//go:noescape
func HasFuncFormDataAppend1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Append1
//go:noescape
func FuncFormDataAppend1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Append1
//go:noescape
func CallFormDataAppend1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref,
	filename js.Ref)

//go:wasmimport plat/js/web try_FormData_Append1
//go:noescape
func TryFormDataAppend1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref,
	filename js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Append2
//go:noescape
func HasFuncFormDataAppend2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Append2
//go:noescape
func FuncFormDataAppend2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Append2
//go:noescape
func CallFormDataAppend2(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref)

//go:wasmimport plat/js/web try_FormData_Append2
//go:noescape
func TryFormDataAppend2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Delete
//go:noescape
func HasFuncFormDataDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Delete
//go:noescape
func FuncFormDataDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Delete
//go:noescape
func CallFormDataDelete(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FormData_Delete
//go:noescape
func TryFormDataDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Get
//go:noescape
func HasFuncFormDataGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Get
//go:noescape
func FuncFormDataGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Get
//go:noescape
func CallFormDataGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FormData_Get
//go:noescape
func TryFormDataGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_GetAll
//go:noescape
func HasFuncFormDataGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_GetAll
//go:noescape
func FuncFormDataGetAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_GetAll
//go:noescape
func CallFormDataGetAll(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FormData_GetAll
//go:noescape
func TryFormDataGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Has
//go:noescape
func HasFuncFormDataHas(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Has
//go:noescape
func FuncFormDataHas(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Has
//go:noescape
func CallFormDataHas(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FormData_Has
//go:noescape
func TryFormDataHas(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Set
//go:noescape
func HasFuncFormDataSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Set
//go:noescape
func FuncFormDataSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Set
//go:noescape
func CallFormDataSet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_FormData_Set
//go:noescape
func TryFormDataSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Set1
//go:noescape
func HasFuncFormDataSet1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Set1
//go:noescape
func FuncFormDataSet1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Set1
//go:noescape
func CallFormDataSet1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref,
	filename js.Ref)

//go:wasmimport plat/js/web try_FormData_Set1
//go:noescape
func TryFormDataSet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref,
	filename js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FormData_Set2
//go:noescape
func HasFuncFormDataSet2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FormData_Set2
//go:noescape
func FuncFormDataSet2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FormData_Set2
//go:noescape
func CallFormDataSet2(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref)

//go:wasmimport plat/js/web try_FormData_Set2
//go:noescape
func TryFormDataSet2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	blobValue js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_TypeMismatch
//go:noescape
func GetValidityStateTypeMismatch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_PatternMismatch
//go:noescape
func GetValidityStatePatternMismatch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_TooLong
//go:noescape
func GetValidityStateTooLong(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_TooShort
//go:noescape
func GetValidityStateTooShort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_RangeUnderflow
//go:noescape
func GetValidityStateRangeUnderflow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_RangeOverflow
//go:noescape
func GetValidityStateRangeOverflow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_StepMismatch
//go:noescape
func GetValidityStateStepMismatch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_BadInput
//go:noescape
func GetValidityStateBadInput(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_CustomError
//go:noescape
func GetValidityStateCustomError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ValidityState_Valid
//go:noescape
func GetValidityStateValid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CustomStateSet_Add
//go:noescape
func HasFuncCustomStateSetAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CustomStateSet_Add
//go:noescape
func FuncCustomStateSetAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CustomStateSet_Add
//go:noescape
func CallCustomStateSetAdd(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_CustomStateSet_Add
//go:noescape
func TryCustomStateSetAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_ShadowRoot
//go:noescape
func GetElementInternalsShadowRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_Form
//go:noescape
func GetElementInternalsForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_WillValidate
//go:noescape
func GetElementInternalsWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_Validity
//go:noescape
func GetElementInternalsValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_ValidationMessage
//go:noescape
func GetElementInternalsValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_Labels
//go:noescape
func GetElementInternalsLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_States
//go:noescape
func GetElementInternalsStates(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ElementInternals_Role
//go:noescape
func GetElementInternalsRole(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_Role
//go:noescape
func SetElementInternalsRole(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaActiveDescendantElement
//go:noescape
func GetElementInternalsAriaActiveDescendantElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaActiveDescendantElement
//go:noescape
func SetElementInternalsAriaActiveDescendantElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaAtomic
//go:noescape
func GetElementInternalsAriaAtomic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaAtomic
//go:noescape
func SetElementInternalsAriaAtomic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaAutoComplete
//go:noescape
func GetElementInternalsAriaAutoComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaAutoComplete
//go:noescape
func SetElementInternalsAriaAutoComplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaBusy
//go:noescape
func GetElementInternalsAriaBusy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaBusy
//go:noescape
func SetElementInternalsAriaBusy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaChecked
//go:noescape
func GetElementInternalsAriaChecked(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaChecked
//go:noescape
func SetElementInternalsAriaChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColCount
//go:noescape
func GetElementInternalsAriaColCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaColCount
//go:noescape
func SetElementInternalsAriaColCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColIndex
//go:noescape
func GetElementInternalsAriaColIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaColIndex
//go:noescape
func SetElementInternalsAriaColIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColIndexText
//go:noescape
func GetElementInternalsAriaColIndexText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaColIndexText
//go:noescape
func SetElementInternalsAriaColIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaColSpan
//go:noescape
func GetElementInternalsAriaColSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaColSpan
//go:noescape
func SetElementInternalsAriaColSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaControlsElements
//go:noescape
func GetElementInternalsAriaControlsElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaControlsElements
//go:noescape
func SetElementInternalsAriaControlsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaCurrent
//go:noescape
func GetElementInternalsAriaCurrent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaCurrent
//go:noescape
func SetElementInternalsAriaCurrent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDescribedByElements
//go:noescape
func GetElementInternalsAriaDescribedByElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaDescribedByElements
//go:noescape
func SetElementInternalsAriaDescribedByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDescription
//go:noescape
func GetElementInternalsAriaDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaDescription
//go:noescape
func SetElementInternalsAriaDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDetailsElements
//go:noescape
func GetElementInternalsAriaDetailsElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaDetailsElements
//go:noescape
func SetElementInternalsAriaDetailsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaDisabled
//go:noescape
func GetElementInternalsAriaDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaDisabled
//go:noescape
func SetElementInternalsAriaDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaErrorMessageElements
//go:noescape
func GetElementInternalsAriaErrorMessageElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaErrorMessageElements
//go:noescape
func SetElementInternalsAriaErrorMessageElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaExpanded
//go:noescape
func GetElementInternalsAriaExpanded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaExpanded
//go:noescape
func SetElementInternalsAriaExpanded(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaFlowToElements
//go:noescape
func GetElementInternalsAriaFlowToElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaFlowToElements
//go:noescape
func SetElementInternalsAriaFlowToElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaHasPopup
//go:noescape
func GetElementInternalsAriaHasPopup(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaHasPopup
//go:noescape
func SetElementInternalsAriaHasPopup(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaHidden
//go:noescape
func GetElementInternalsAriaHidden(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaHidden
//go:noescape
func SetElementInternalsAriaHidden(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaInvalid
//go:noescape
func GetElementInternalsAriaInvalid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaInvalid
//go:noescape
func SetElementInternalsAriaInvalid(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaKeyShortcuts
//go:noescape
func GetElementInternalsAriaKeyShortcuts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaKeyShortcuts
//go:noescape
func SetElementInternalsAriaKeyShortcuts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLabel
//go:noescape
func GetElementInternalsAriaLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaLabel
//go:noescape
func SetElementInternalsAriaLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLabelledByElements
//go:noescape
func GetElementInternalsAriaLabelledByElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaLabelledByElements
//go:noescape
func SetElementInternalsAriaLabelledByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLevel
//go:noescape
func GetElementInternalsAriaLevel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaLevel
//go:noescape
func SetElementInternalsAriaLevel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaLive
//go:noescape
func GetElementInternalsAriaLive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaLive
//go:noescape
func SetElementInternalsAriaLive(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaModal
//go:noescape
func GetElementInternalsAriaModal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaModal
//go:noescape
func SetElementInternalsAriaModal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaMultiLine
//go:noescape
func GetElementInternalsAriaMultiLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaMultiLine
//go:noescape
func SetElementInternalsAriaMultiLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaMultiSelectable
//go:noescape
func GetElementInternalsAriaMultiSelectable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaMultiSelectable
//go:noescape
func SetElementInternalsAriaMultiSelectable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaOrientation
//go:noescape
func GetElementInternalsAriaOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaOrientation
//go:noescape
func SetElementInternalsAriaOrientation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaOwnsElements
//go:noescape
func GetElementInternalsAriaOwnsElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaOwnsElements
//go:noescape
func SetElementInternalsAriaOwnsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPlaceholder
//go:noescape
func GetElementInternalsAriaPlaceholder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaPlaceholder
//go:noescape
func SetElementInternalsAriaPlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPosInSet
//go:noescape
func GetElementInternalsAriaPosInSet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaPosInSet
//go:noescape
func SetElementInternalsAriaPosInSet(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaPressed
//go:noescape
func GetElementInternalsAriaPressed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaPressed
//go:noescape
func SetElementInternalsAriaPressed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaReadOnly
//go:noescape
func GetElementInternalsAriaReadOnly(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaReadOnly
//go:noescape
func SetElementInternalsAriaReadOnly(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRequired
//go:noescape
func GetElementInternalsAriaRequired(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRequired
//go:noescape
func SetElementInternalsAriaRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRoleDescription
//go:noescape
func GetElementInternalsAriaRoleDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRoleDescription
//go:noescape
func SetElementInternalsAriaRoleDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowCount
//go:noescape
func GetElementInternalsAriaRowCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRowCount
//go:noescape
func SetElementInternalsAriaRowCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowIndex
//go:noescape
func GetElementInternalsAriaRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRowIndex
//go:noescape
func SetElementInternalsAriaRowIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowIndexText
//go:noescape
func GetElementInternalsAriaRowIndexText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRowIndexText
//go:noescape
func SetElementInternalsAriaRowIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaRowSpan
//go:noescape
func GetElementInternalsAriaRowSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaRowSpan
//go:noescape
func SetElementInternalsAriaRowSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSelected
//go:noescape
func GetElementInternalsAriaSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaSelected
//go:noescape
func SetElementInternalsAriaSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSetSize
//go:noescape
func GetElementInternalsAriaSetSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaSetSize
//go:noescape
func SetElementInternalsAriaSetSize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaSort
//go:noescape
func GetElementInternalsAriaSort(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaSort
//go:noescape
func SetElementInternalsAriaSort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueMax
//go:noescape
func GetElementInternalsAriaValueMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaValueMax
//go:noescape
func SetElementInternalsAriaValueMax(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueMin
//go:noescape
func GetElementInternalsAriaValueMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaValueMin
//go:noescape
func SetElementInternalsAriaValueMin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueNow
//go:noescape
func GetElementInternalsAriaValueNow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaValueNow
//go:noescape
func SetElementInternalsAriaValueNow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ElementInternals_AriaValueText
//go:noescape
func GetElementInternalsAriaValueText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_ElementInternals_AriaValueText
//go:noescape
func SetElementInternalsAriaValueText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_ElementInternals_SetFormValue
//go:noescape
func HasFuncElementInternalsSetFormValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetFormValue
//go:noescape
func FuncElementInternalsSetFormValue(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetFormValue
//go:noescape
func CallElementInternalsSetFormValue(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref,
	state js.Ref)

//go:wasmimport plat/js/web try_ElementInternals_SetFormValue
//go:noescape
func TryElementInternalsSetFormValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref,
	state js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_SetFormValue1
//go:noescape
func HasFuncElementInternalsSetFormValue1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetFormValue1
//go:noescape
func FuncElementInternalsSetFormValue1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetFormValue1
//go:noescape
func CallElementInternalsSetFormValue1(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_ElementInternals_SetFormValue1
//go:noescape
func TryElementInternalsSetFormValue1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_SetValidity
//go:noescape
func HasFuncElementInternalsSetValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity
//go:noescape
func FuncElementInternalsSetValidity(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetValidity
//go:noescape
func CallElementInternalsSetValidity(
	this js.Ref, retPtr unsafe.Pointer,
	flags unsafe.Pointer,
	message js.Ref,
	anchor js.Ref)

//go:wasmimport plat/js/web try_ElementInternals_SetValidity
//go:noescape
func TryElementInternalsSetValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	flags unsafe.Pointer,
	message js.Ref,
	anchor js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_SetValidity1
//go:noescape
func HasFuncElementInternalsSetValidity1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity1
//go:noescape
func FuncElementInternalsSetValidity1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetValidity1
//go:noescape
func CallElementInternalsSetValidity1(
	this js.Ref, retPtr unsafe.Pointer,
	flags unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_ElementInternals_SetValidity1
//go:noescape
func TryElementInternalsSetValidity1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	flags unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_SetValidity2
//go:noescape
func HasFuncElementInternalsSetValidity2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity2
//go:noescape
func FuncElementInternalsSetValidity2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetValidity2
//go:noescape
func CallElementInternalsSetValidity2(
	this js.Ref, retPtr unsafe.Pointer,
	flags unsafe.Pointer)

//go:wasmimport plat/js/web try_ElementInternals_SetValidity2
//go:noescape
func TryElementInternalsSetValidity2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	flags unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_SetValidity3
//go:noescape
func HasFuncElementInternalsSetValidity3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_SetValidity3
//go:noescape
func FuncElementInternalsSetValidity3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_SetValidity3
//go:noescape
func CallElementInternalsSetValidity3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ElementInternals_SetValidity3
//go:noescape
func TryElementInternalsSetValidity3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_CheckValidity
//go:noescape
func HasFuncElementInternalsCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_CheckValidity
//go:noescape
func FuncElementInternalsCheckValidity(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_CheckValidity
//go:noescape
func CallElementInternalsCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ElementInternals_CheckValidity
//go:noescape
func TryElementInternalsCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ElementInternals_ReportValidity
//go:noescape
func HasFuncElementInternalsReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ElementInternals_ReportValidity
//go:noescape
func FuncElementInternalsReportValidity(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ElementInternals_ReportValidity
//go:noescape
func CallElementInternalsReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ElementInternals_ReportValidity
//go:noescape
func TryElementInternalsReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_SelectionStart
//go:noescape
func GetEditContextSelectionStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_SelectionEnd
//go:noescape
func GetEditContextSelectionEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_CompositionRangeStart
//go:noescape
func GetEditContextCompositionRangeStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_CompositionRangeEnd
//go:noescape
func GetEditContextCompositionRangeEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_IsComposing
//go:noescape
func GetEditContextIsComposing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_ControlBounds
//go:noescape
func GetEditContextControlBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_SelectionBounds
//go:noescape
func GetEditContextSelectionBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_EditContext_CharacterBoundsRangeStart
//go:noescape
func GetEditContextCharacterBoundsRangeStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_UpdateText
//go:noescape
func HasFuncEditContextUpdateText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateText
//go:noescape
func FuncEditContextUpdateText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_UpdateText
//go:noescape
func CallEditContextUpdateText(
	this js.Ref, retPtr unsafe.Pointer,
	rangeStart uint32,
	rangeEnd uint32,
	text js.Ref)

//go:wasmimport plat/js/web try_EditContext_UpdateText
//go:noescape
func TryEditContextUpdateText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rangeStart uint32,
	rangeEnd uint32,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_UpdateSelection
//go:noescape
func HasFuncEditContextUpdateSelection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateSelection
//go:noescape
func FuncEditContextUpdateSelection(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_UpdateSelection
//go:noescape
func CallEditContextUpdateSelection(
	this js.Ref, retPtr unsafe.Pointer,
	start uint32,
	end uint32)

//go:wasmimport plat/js/web try_EditContext_UpdateSelection
//go:noescape
func TryEditContextUpdateSelection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start uint32,
	end uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_UpdateControlBounds
//go:noescape
func HasFuncEditContextUpdateControlBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateControlBounds
//go:noescape
func FuncEditContextUpdateControlBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_UpdateControlBounds
//go:noescape
func CallEditContextUpdateControlBounds(
	this js.Ref, retPtr unsafe.Pointer,
	controlBounds js.Ref)

//go:wasmimport plat/js/web try_EditContext_UpdateControlBounds
//go:noescape
func TryEditContextUpdateControlBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	controlBounds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_UpdateSelectionBounds
//go:noescape
func HasFuncEditContextUpdateSelectionBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateSelectionBounds
//go:noescape
func FuncEditContextUpdateSelectionBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_UpdateSelectionBounds
//go:noescape
func CallEditContextUpdateSelectionBounds(
	this js.Ref, retPtr unsafe.Pointer,
	selectionBounds js.Ref)

//go:wasmimport plat/js/web try_EditContext_UpdateSelectionBounds
//go:noescape
func TryEditContextUpdateSelectionBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectionBounds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_UpdateCharacterBounds
//go:noescape
func HasFuncEditContextUpdateCharacterBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_UpdateCharacterBounds
//go:noescape
func FuncEditContextUpdateCharacterBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_UpdateCharacterBounds
//go:noescape
func CallEditContextUpdateCharacterBounds(
	this js.Ref, retPtr unsafe.Pointer,
	rangeStart uint32,
	characterBounds js.Ref)

//go:wasmimport plat/js/web try_EditContext_UpdateCharacterBounds
//go:noescape
func TryEditContextUpdateCharacterBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rangeStart uint32,
	characterBounds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_AttachedElements
//go:noescape
func HasFuncEditContextAttachedElements(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_AttachedElements
//go:noescape
func FuncEditContextAttachedElements(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_AttachedElements
//go:noescape
func CallEditContextAttachedElements(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_EditContext_AttachedElements
//go:noescape
func TryEditContextAttachedElements(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_EditContext_CharacterBounds
//go:noescape
func HasFuncEditContextCharacterBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_EditContext_CharacterBounds
//go:noescape
func FuncEditContextCharacterBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_EditContext_CharacterBounds
//go:noescape
func CallEditContextCharacterBounds(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_EditContext_CharacterBounds
//go:noescape
func TryEditContextCharacterBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_Title
//go:noescape
func GetHTMLElementTitle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Title
//go:noescape
func SetHTMLElementTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Lang
//go:noescape
func GetHTMLElementLang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Lang
//go:noescape
func SetHTMLElementLang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Translate
//go:noescape
func GetHTMLElementTranslate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Translate
//go:noescape
func SetHTMLElementTranslate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Dir
//go:noescape
func GetHTMLElementDir(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Dir
//go:noescape
func SetHTMLElementDir(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Hidden
//go:noescape
func GetHTMLElementHidden(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Hidden
//go:noescape
func SetHTMLElementHidden(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Inert
//go:noescape
func GetHTMLElementInert(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Inert
//go:noescape
func SetHTMLElementInert(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_AccessKey
//go:noescape
func GetHTMLElementAccessKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_AccessKey
//go:noescape
func SetHTMLElementAccessKey(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_AccessKeyLabel
//go:noescape
func GetHTMLElementAccessKeyLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_Draggable
//go:noescape
func GetHTMLElementDraggable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Draggable
//go:noescape
func SetHTMLElementDraggable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Spellcheck
//go:noescape
func GetHTMLElementSpellcheck(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Spellcheck
//go:noescape
func SetHTMLElementSpellcheck(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Autocapitalize
//go:noescape
func GetHTMLElementAutocapitalize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Autocapitalize
//go:noescape
func SetHTMLElementAutocapitalize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_InnerText
//go:noescape
func GetHTMLElementInnerText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_InnerText
//go:noescape
func SetHTMLElementInnerText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_OuterText
//go:noescape
func GetHTMLElementOuterText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_OuterText
//go:noescape
func SetHTMLElementOuterText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Popover
//go:noescape
func GetHTMLElementPopover(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Popover
//go:noescape
func SetHTMLElementPopover(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_OffsetParent
//go:noescape
func GetHTMLElementOffsetParent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_OffsetTop
//go:noescape
func GetHTMLElementOffsetTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_OffsetLeft
//go:noescape
func GetHTMLElementOffsetLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_OffsetWidth
//go:noescape
func GetHTMLElementOffsetWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_OffsetHeight
//go:noescape
func GetHTMLElementOffsetHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_EditContext
//go:noescape
func GetHTMLElementEditContext(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_EditContext
//go:noescape
func SetHTMLElementEditContext(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Style
//go:noescape
func GetHTMLElementStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_AttributeStyleMap
//go:noescape
func GetHTMLElementAttributeStyleMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_Dataset
//go:noescape
func GetHTMLElementDataset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_Nonce
//go:noescape
func GetHTMLElementNonce(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Nonce
//go:noescape
func SetHTMLElementNonce(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_Autofocus
//go:noescape
func GetHTMLElementAutofocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_Autofocus
//go:noescape
func SetHTMLElementAutofocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_TabIndex
//go:noescape
func GetHTMLElementTabIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_TabIndex
//go:noescape
func SetHTMLElementTabIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_ContentEditable
//go:noescape
func GetHTMLElementContentEditable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_ContentEditable
//go:noescape
func SetHTMLElementContentEditable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_EnterKeyHint
//go:noescape
func GetHTMLElementEnterKeyHint(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_EnterKeyHint
//go:noescape
func SetHTMLElementEnterKeyHint(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_IsContentEditable
//go:noescape
func GetHTMLElementIsContentEditable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLElement_InputMode
//go:noescape
func GetHTMLElementInputMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_InputMode
//go:noescape
func SetHTMLElementInputMode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLElement_VirtualKeyboardPolicy
//go:noescape
func GetHTMLElementVirtualKeyboardPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLElement_VirtualKeyboardPolicy
//go:noescape
func SetHTMLElementVirtualKeyboardPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLElement_Click
//go:noescape
func HasFuncHTMLElementClick(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Click
//go:noescape
func FuncHTMLElementClick(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_Click
//go:noescape
func CallHTMLElementClick(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_Click
//go:noescape
func TryHTMLElementClick(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_AttachInternals
//go:noescape
func HasFuncHTMLElementAttachInternals(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_AttachInternals
//go:noescape
func FuncHTMLElementAttachInternals(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_AttachInternals
//go:noescape
func CallHTMLElementAttachInternals(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_AttachInternals
//go:noescape
func TryHTMLElementAttachInternals(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_ShowPopover
//go:noescape
func HasFuncHTMLElementShowPopover(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_ShowPopover
//go:noescape
func FuncHTMLElementShowPopover(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_ShowPopover
//go:noescape
func CallHTMLElementShowPopover(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_ShowPopover
//go:noescape
func TryHTMLElementShowPopover(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_HidePopover
//go:noescape
func HasFuncHTMLElementHidePopover(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_HidePopover
//go:noescape
func FuncHTMLElementHidePopover(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_HidePopover
//go:noescape
func CallHTMLElementHidePopover(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_HidePopover
//go:noescape
func TryHTMLElementHidePopover(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_TogglePopover
//go:noescape
func HasFuncHTMLElementTogglePopover(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_TogglePopover
//go:noescape
func FuncHTMLElementTogglePopover(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_TogglePopover
//go:noescape
func CallHTMLElementTogglePopover(
	this js.Ref, retPtr unsafe.Pointer,
	force js.Ref)

//go:wasmimport plat/js/web try_HTMLElement_TogglePopover
//go:noescape
func TryHTMLElementTogglePopover(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	force js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_TogglePopover1
//go:noescape
func HasFuncHTMLElementTogglePopover1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_TogglePopover1
//go:noescape
func FuncHTMLElementTogglePopover1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_TogglePopover1
//go:noescape
func CallHTMLElementTogglePopover1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_TogglePopover1
//go:noescape
func TryHTMLElementTogglePopover1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_Focus
//go:noescape
func HasFuncHTMLElementFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Focus
//go:noescape
func FuncHTMLElementFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_Focus
//go:noescape
func CallHTMLElementFocus(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_Focus
//go:noescape
func TryHTMLElementFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_Focus1
//go:noescape
func HasFuncHTMLElementFocus1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Focus1
//go:noescape
func FuncHTMLElementFocus1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_Focus1
//go:noescape
func CallHTMLElementFocus1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_Focus1
//go:noescape
func TryHTMLElementFocus1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLElement_Blur
//go:noescape
func HasFuncHTMLElementBlur(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLElement_Blur
//go:noescape
func FuncHTMLElementBlur(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLElement_Blur
//go:noescape
func CallHTMLElementBlur(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLElement_Blur
//go:noescape
func TryHTMLElementBlur(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLScriptElement_Src
//go:noescape
func GetHTMLScriptElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Src
//go:noescape
func SetHTMLScriptElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Type
//go:noescape
func GetHTMLScriptElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Type
//go:noescape
func SetHTMLScriptElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_NoModule
//go:noescape
func GetHTMLScriptElementNoModule(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_NoModule
//go:noescape
func SetHTMLScriptElementNoModule(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Async
//go:noescape
func GetHTMLScriptElementAsync(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Async
//go:noescape
func SetHTMLScriptElementAsync(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Defer
//go:noescape
func GetHTMLScriptElementDefer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Defer
//go:noescape
func SetHTMLScriptElementDefer(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_CrossOrigin
//go:noescape
func GetHTMLScriptElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_CrossOrigin
//go:noescape
func SetHTMLScriptElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Text
//go:noescape
func GetHTMLScriptElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Text
//go:noescape
func SetHTMLScriptElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Integrity
//go:noescape
func GetHTMLScriptElementIntegrity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Integrity
//go:noescape
func SetHTMLScriptElementIntegrity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_ReferrerPolicy
//go:noescape
func GetHTMLScriptElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_ReferrerPolicy
//go:noescape
func SetHTMLScriptElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Blocking
//go:noescape
func GetHTMLScriptElementBlocking(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLScriptElement_FetchPriority
//go:noescape
func GetHTMLScriptElementFetchPriority(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_FetchPriority
//go:noescape
func SetHTMLScriptElementFetchPriority(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Charset
//go:noescape
func GetHTMLScriptElementCharset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Charset
//go:noescape
func SetHTMLScriptElementCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_Event
//go:noescape
func GetHTMLScriptElementEvent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_Event
//go:noescape
func SetHTMLScriptElementEvent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_HtmlFor
//go:noescape
func GetHTMLScriptElementHtmlFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_HtmlFor
//go:noescape
func SetHTMLScriptElementHtmlFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLScriptElement_AttributionSrc
//go:noescape
func GetHTMLScriptElementAttributionSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLScriptElement_AttributionSrc
//go:noescape
func SetHTMLScriptElementAttributionSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLScriptElement_Supports
//go:noescape
func HasFuncHTMLScriptElementSupports(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLScriptElement_Supports
//go:noescape
func FuncHTMLScriptElementSupports(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_HTMLScriptElement_Supports
//go:noescape
func CallHTMLScriptElementSupports(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_HTMLScriptElement_Supports
//go:noescape
func TryHTMLScriptElementSupports(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGScriptElement_Type
//go:noescape
func GetSVGScriptElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGScriptElement_Type
//go:noescape
func SetSVGScriptElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGScriptElement_CrossOrigin
//go:noescape
func GetSVGScriptElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGScriptElement_CrossOrigin
//go:noescape
func SetSVGScriptElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGScriptElement_Href
//go:noescape
func GetSVGScriptElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)
