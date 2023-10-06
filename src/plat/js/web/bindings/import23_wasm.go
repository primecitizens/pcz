// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web get_IdleDeadline_DidTimeout
//go:noescape
func GetIdleDeadlineDidTimeout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IdleDeadline_TimeRemaining
//go:noescape
func HasIdleDeadlineTimeRemaining(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdleDeadline_TimeRemaining
//go:noescape
func IdleDeadlineTimeRemainingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IdleDeadline_TimeRemaining
//go:noescape
func CallIdleDeadlineTimeRemaining(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IdleDeadline_TimeRemaining
//go:noescape
func TryIdleDeadlineTimeRemaining(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_IdleRequestOptions
//go:noescape
func IdleRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdleRequestOptions
//go:noescape
func IdleRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_WriteCommandType
//go:noescape
func ConstOfWriteCommandType(str js.Ref) uint32

//go:wasmimport plat/js/web store_WriteParams
//go:noescape
func WriteParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WriteParams
//go:noescape
func WriteParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_FileSystemWritableFileStream_FileSystemWritableFileStream
//go:noescape
func NewFileSystemWritableFileStreamByFileSystemWritableFileStream(
	underlyingSink js.Ref,
	strategy unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FileSystemWritableFileStream_FileSystemWritableFileStream1
//go:noescape
func NewFileSystemWritableFileStreamByFileSystemWritableFileStream1(
	underlyingSink js.Ref) js.Ref

//go:wasmimport plat/js/web new_FileSystemWritableFileStream_FileSystemWritableFileStream2
//go:noescape
func NewFileSystemWritableFileStreamByFileSystemWritableFileStream2() js.Ref

//go:wasmimport plat/js/web has_FileSystemWritableFileStream_Write
//go:noescape
func HasFileSystemWritableFileStreamWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Write
//go:noescape
func FileSystemWritableFileStreamWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemWritableFileStream_Write
//go:noescape
func CallFileSystemWritableFileStreamWrite(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_FileSystemWritableFileStream_Write
//go:noescape
func TryFileSystemWritableFileStreamWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemWritableFileStream_Seek
//go:noescape
func HasFileSystemWritableFileStreamSeek(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Seek
//go:noescape
func FileSystemWritableFileStreamSeekFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemWritableFileStream_Seek
//go:noescape
func CallFileSystemWritableFileStreamSeek(
	this js.Ref, retPtr unsafe.Pointer,
	position float64)

//go:wasmimport plat/js/web try_FileSystemWritableFileStream_Seek
//go:noescape
func TryFileSystemWritableFileStreamSeek(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	position float64) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemWritableFileStream_Truncate
//go:noescape
func HasFileSystemWritableFileStreamTruncate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Truncate
//go:noescape
func FileSystemWritableFileStreamTruncateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemWritableFileStream_Truncate
//go:noescape
func CallFileSystemWritableFileStreamTruncate(
	this js.Ref, retPtr unsafe.Pointer,
	size float64)

//go:wasmimport plat/js/web try_FileSystemWritableFileStream_Truncate
//go:noescape
func TryFileSystemWritableFileStreamTruncate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	size float64) (ok js.Ref)

//go:wasmimport plat/js/web store_FileSystemCreateWritableOptions
//go:noescape
func FileSystemCreateWritableOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemCreateWritableOptions
//go:noescape
func FileSystemCreateWritableOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FileSystemReadWriteOptions
//go:noescape
func FileSystemReadWriteOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemReadWriteOptions
//go:noescape
func FileSystemReadWriteOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Read
//go:noescape
func HasFileSystemSyncAccessHandleRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Read
//go:noescape
func FileSystemSyncAccessHandleReadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Read
//go:noescape
func CallFileSystemSyncAccessHandleRead(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Read
//go:noescape
func TryFileSystemSyncAccessHandleRead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Read1
//go:noescape
func HasFileSystemSyncAccessHandleRead1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Read1
//go:noescape
func FileSystemSyncAccessHandleRead1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Read1
//go:noescape
func CallFileSystemSyncAccessHandleRead1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Read1
//go:noescape
func TryFileSystemSyncAccessHandleRead1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Write
//go:noescape
func HasFileSystemSyncAccessHandleWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Write
//go:noescape
func FileSystemSyncAccessHandleWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Write
//go:noescape
func CallFileSystemSyncAccessHandleWrite(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Write
//go:noescape
func TryFileSystemSyncAccessHandleWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Write1
//go:noescape
func HasFileSystemSyncAccessHandleWrite1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Write1
//go:noescape
func FileSystemSyncAccessHandleWrite1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Write1
//go:noescape
func CallFileSystemSyncAccessHandleWrite1(
	this js.Ref, retPtr unsafe.Pointer,
	buffer js.Ref)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Write1
//go:noescape
func TryFileSystemSyncAccessHandleWrite1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	buffer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Truncate
//go:noescape
func HasFileSystemSyncAccessHandleTruncate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Truncate
//go:noescape
func FileSystemSyncAccessHandleTruncateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Truncate
//go:noescape
func CallFileSystemSyncAccessHandleTruncate(
	this js.Ref, retPtr unsafe.Pointer,
	newSize float64)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Truncate
//go:noescape
func TryFileSystemSyncAccessHandleTruncate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newSize float64) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_GetSize
//go:noescape
func HasFileSystemSyncAccessHandleGetSize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_GetSize
//go:noescape
func FileSystemSyncAccessHandleGetSizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_GetSize
//go:noescape
func CallFileSystemSyncAccessHandleGetSize(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_GetSize
//go:noescape
func TryFileSystemSyncAccessHandleGetSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Flush
//go:noescape
func HasFileSystemSyncAccessHandleFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Flush
//go:noescape
func FileSystemSyncAccessHandleFlushFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Flush
//go:noescape
func CallFileSystemSyncAccessHandleFlush(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Flush
//go:noescape
func TryFileSystemSyncAccessHandleFlush(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemSyncAccessHandle_Close
//go:noescape
func HasFileSystemSyncAccessHandleClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Close
//go:noescape
func FileSystemSyncAccessHandleCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemSyncAccessHandle_Close
//go:noescape
func CallFileSystemSyncAccessHandleClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemSyncAccessHandle_Close
//go:noescape
func TryFileSystemSyncAccessHandleClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemFileHandle_GetFile
//go:noescape
func HasFileSystemFileHandleGetFile(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_GetFile
//go:noescape
func FileSystemFileHandleGetFileFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileHandle_GetFile
//go:noescape
func CallFileSystemFileHandleGetFile(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemFileHandle_GetFile
//go:noescape
func TryFileSystemFileHandleGetFile(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemFileHandle_CreateWritable
//go:noescape
func HasFileSystemFileHandleCreateWritable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateWritable
//go:noescape
func FileSystemFileHandleCreateWritableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileHandle_CreateWritable
//go:noescape
func CallFileSystemFileHandleCreateWritable(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemFileHandle_CreateWritable
//go:noescape
func TryFileSystemFileHandleCreateWritable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemFileHandle_CreateWritable1
//go:noescape
func HasFileSystemFileHandleCreateWritable1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateWritable1
//go:noescape
func FileSystemFileHandleCreateWritable1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileHandle_CreateWritable1
//go:noescape
func CallFileSystemFileHandleCreateWritable1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemFileHandle_CreateWritable1
//go:noescape
func TryFileSystemFileHandleCreateWritable1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemFileHandle_CreateSyncAccessHandle
//go:noescape
func HasFileSystemFileHandleCreateSyncAccessHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateSyncAccessHandle
//go:noescape
func FileSystemFileHandleCreateSyncAccessHandleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemFileHandle_CreateSyncAccessHandle
//go:noescape
func CallFileSystemFileHandleCreateSyncAccessHandle(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemFileHandle_CreateSyncAccessHandle
//go:noescape
func TryFileSystemFileHandleCreateSyncAccessHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_FilePickerAcceptType
//go:noescape
func FilePickerAcceptTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FilePickerAcceptType
//go:noescape
func FilePickerAcceptTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_WellKnownDirectory
//go:noescape
func ConstOfWellKnownDirectory(str js.Ref) uint32

//go:wasmimport plat/js/web store_OpenFilePickerOptions
//go:noescape
func OpenFilePickerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OpenFilePickerOptions
//go:noescape
func OpenFilePickerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SaveFilePickerOptions
//go:noescape
func SaveFilePickerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SaveFilePickerOptions
//go:noescape
func SaveFilePickerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FileSystemGetFileOptions
//go:noescape
func FileSystemGetFileOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemGetFileOptions
//go:noescape
func FileSystemGetFileOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FileSystemGetDirectoryOptions
//go:noescape
func FileSystemGetDirectoryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemGetDirectoryOptions
//go:noescape
func FileSystemGetDirectoryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FileSystemRemoveOptions
//go:noescape
func FileSystemRemoveOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemRemoveOptions
//go:noescape
func FileSystemRemoveOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_GetFileHandle
//go:noescape
func HasFileSystemDirectoryHandleGetFileHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetFileHandle
//go:noescape
func FileSystemDirectoryHandleGetFileHandleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_GetFileHandle
//go:noescape
func CallFileSystemDirectoryHandleGetFileHandle(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_GetFileHandle
//go:noescape
func TryFileSystemDirectoryHandleGetFileHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_GetFileHandle1
//go:noescape
func HasFileSystemDirectoryHandleGetFileHandle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetFileHandle1
//go:noescape
func FileSystemDirectoryHandleGetFileHandle1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_GetFileHandle1
//go:noescape
func CallFileSystemDirectoryHandleGetFileHandle1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_GetFileHandle1
//go:noescape
func TryFileSystemDirectoryHandleGetFileHandle1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_GetDirectoryHandle
//go:noescape
func HasFileSystemDirectoryHandleGetDirectoryHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetDirectoryHandle
//go:noescape
func FileSystemDirectoryHandleGetDirectoryHandleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_GetDirectoryHandle
//go:noescape
func CallFileSystemDirectoryHandleGetDirectoryHandle(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_GetDirectoryHandle
//go:noescape
func TryFileSystemDirectoryHandleGetDirectoryHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_GetDirectoryHandle1
//go:noescape
func HasFileSystemDirectoryHandleGetDirectoryHandle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetDirectoryHandle1
//go:noescape
func FileSystemDirectoryHandleGetDirectoryHandle1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_GetDirectoryHandle1
//go:noescape
func CallFileSystemDirectoryHandleGetDirectoryHandle1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_GetDirectoryHandle1
//go:noescape
func TryFileSystemDirectoryHandleGetDirectoryHandle1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_RemoveEntry
//go:noescape
func HasFileSystemDirectoryHandleRemoveEntry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_RemoveEntry
//go:noescape
func FileSystemDirectoryHandleRemoveEntryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_RemoveEntry
//go:noescape
func CallFileSystemDirectoryHandleRemoveEntry(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_RemoveEntry
//go:noescape
func TryFileSystemDirectoryHandleRemoveEntry(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_RemoveEntry1
//go:noescape
func HasFileSystemDirectoryHandleRemoveEntry1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_RemoveEntry1
//go:noescape
func FileSystemDirectoryHandleRemoveEntry1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_RemoveEntry1
//go:noescape
func CallFileSystemDirectoryHandleRemoveEntry1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_RemoveEntry1
//go:noescape
func TryFileSystemDirectoryHandleRemoveEntry1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryHandle_Resolve
//go:noescape
func HasFileSystemDirectoryHandleResolve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_Resolve
//go:noescape
func FileSystemDirectoryHandleResolveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FileSystemDirectoryHandle_Resolve
//go:noescape
func CallFileSystemDirectoryHandleResolve(
	this js.Ref, retPtr unsafe.Pointer,
	possibleDescendant js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryHandle_Resolve
//go:noescape
func TryFileSystemDirectoryHandleResolve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	possibleDescendant js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_DirectoryPickerOptions
//go:noescape
func DirectoryPickerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DirectoryPickerOptions
//go:noescape
func DirectoryPickerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_FontData_PostscriptName
//go:noescape
func GetFontDataPostscriptName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontData_FullName
//go:noescape
func GetFontDataFullName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontData_Family
//go:noescape
func GetFontDataFamily(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontData_Style
//go:noescape
func GetFontDataStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontData_Blob
//go:noescape
func HasFontDataBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontData_Blob
//go:noescape
func FontDataBlobFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontData_Blob
//go:noescape
func CallFontDataBlob(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FontData_Blob
//go:noescape
func TryFontDataBlob(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_QueryOptions
//go:noescape
func QueryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_QueryOptions
//go:noescape
func QueryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_ScreenDetails_Screens
//go:noescape
func GetScreenDetailsScreens(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScreenDetails_CurrentScreen
//go:noescape
func GetScreenDetailsCurrentScreen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_ItemType
//go:noescape
func ConstOfItemType(str js.Ref) uint32

//go:wasmimport plat/js/web store_ItemDetails
//go:noescape
func ItemDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ItemDetails
//go:noescape
func ItemDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PurchaseDetails
//go:noescape
func PurchaseDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PurchaseDetails
//go:noescape
func PurchaseDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_DigitalGoodsService_GetDetails
//go:noescape
func HasDigitalGoodsServiceGetDetails(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_GetDetails
//go:noescape
func DigitalGoodsServiceGetDetailsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DigitalGoodsService_GetDetails
//go:noescape
func CallDigitalGoodsServiceGetDetails(
	this js.Ref, retPtr unsafe.Pointer,
	itemIds js.Ref)

//go:wasmimport plat/js/web try_DigitalGoodsService_GetDetails
//go:noescape
func TryDigitalGoodsServiceGetDetails(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	itemIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DigitalGoodsService_ListPurchases
//go:noescape
func HasDigitalGoodsServiceListPurchases(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_ListPurchases
//go:noescape
func DigitalGoodsServiceListPurchasesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DigitalGoodsService_ListPurchases
//go:noescape
func CallDigitalGoodsServiceListPurchases(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DigitalGoodsService_ListPurchases
//go:noescape
func TryDigitalGoodsServiceListPurchases(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DigitalGoodsService_ListPurchaseHistory
//go:noescape
func HasDigitalGoodsServiceListPurchaseHistory(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_ListPurchaseHistory
//go:noescape
func DigitalGoodsServiceListPurchaseHistoryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DigitalGoodsService_ListPurchaseHistory
//go:noescape
func CallDigitalGoodsServiceListPurchaseHistory(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DigitalGoodsService_ListPurchaseHistory
//go:noescape
func TryDigitalGoodsServiceListPurchaseHistory(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DigitalGoodsService_Consume
//go:noescape
func HasDigitalGoodsServiceConsume(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_Consume
//go:noescape
func DigitalGoodsServiceConsumeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DigitalGoodsService_Consume
//go:noescape
func CallDigitalGoodsServiceConsume(
	this js.Ref, retPtr unsafe.Pointer,
	purchaseToken js.Ref)

//go:wasmimport plat/js/web try_DigitalGoodsService_Consume
//go:noescape
func TryDigitalGoodsServiceConsume(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	purchaseToken js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_ImageOrientation
//go:noescape
func ConstOfImageOrientation(str js.Ref) uint32

//go:wasmimport plat/js/web constof_PremultiplyAlpha
//go:noescape
func ConstOfPremultiplyAlpha(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ResizeQuality
//go:noescape
func ConstOfResizeQuality(str js.Ref) uint32

//go:wasmimport plat/js/web store_ImageBitmapOptions
//go:noescape
func ImageBitmapOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ImageBitmapOptions
//go:noescape
func ImageBitmapOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ScrollRestoration
//go:noescape
func ConstOfScrollRestoration(str js.Ref) uint32

//go:wasmimport plat/js/web get_History_Length
//go:noescape
func GetHistoryLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_History_ScrollRestoration
//go:noescape
func GetHistoryScrollRestoration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_History_ScrollRestoration
//go:noescape
func SetHistoryScrollRestoration(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_History_State
//go:noescape
func GetHistoryState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_History_Go
//go:noescape
func HasHistoryGo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Go
//go:noescape
func HistoryGoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_Go
//go:noescape
func CallHistoryGo(
	this js.Ref, retPtr unsafe.Pointer,
	delta int32)

//go:wasmimport plat/js/web try_History_Go
//go:noescape
func TryHistoryGo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	delta int32) (ok js.Ref)

//go:wasmimport plat/js/web has_History_Go1
//go:noescape
func HasHistoryGo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Go1
//go:noescape
func HistoryGo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_Go1
//go:noescape
func CallHistoryGo1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_History_Go1
//go:noescape
func TryHistoryGo1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_History_Back
//go:noescape
func HasHistoryBack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Back
//go:noescape
func HistoryBackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_Back
//go:noescape
func CallHistoryBack(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_History_Back
//go:noescape
func TryHistoryBack(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_History_Forward
//go:noescape
func HasHistoryForward(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Forward
//go:noescape
func HistoryForwardFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_Forward
//go:noescape
func CallHistoryForward(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_History_Forward
//go:noescape
func TryHistoryForward(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_History_PushState
//go:noescape
func HasHistoryPushState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_PushState
//go:noescape
func HistoryPushStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_PushState
//go:noescape
func CallHistoryPushState(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref,
	url js.Ref)

//go:wasmimport plat/js/web try_History_PushState
//go:noescape
func TryHistoryPushState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_History_PushState1
//go:noescape
func HasHistoryPushState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_PushState1
//go:noescape
func HistoryPushState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_PushState1
//go:noescape
func CallHistoryPushState1(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref)

//go:wasmimport plat/js/web try_History_PushState1
//go:noescape
func TryHistoryPushState1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_History_ReplaceState
//go:noescape
func HasHistoryReplaceState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_ReplaceState
//go:noescape
func HistoryReplaceStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_ReplaceState
//go:noescape
func CallHistoryReplaceState(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref,
	url js.Ref)

//go:wasmimport plat/js/web try_History_ReplaceState
//go:noescape
func TryHistoryReplaceState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_History_ReplaceState1
//go:noescape
func HasHistoryReplaceState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_ReplaceState1
//go:noescape
func HistoryReplaceState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_History_ReplaceState1
//go:noescape
func CallHistoryReplaceState1(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref)

//go:wasmimport plat/js/web try_History_ReplaceState1
//go:noescape
func TryHistoryReplaceState1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	unused js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationHistoryEntry_Url
//go:noescape
func GetNavigationHistoryEntryUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationHistoryEntry_Key
//go:noescape
func GetNavigationHistoryEntryKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationHistoryEntry_Id
//go:noescape
func GetNavigationHistoryEntryId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationHistoryEntry_Index
//go:noescape
func GetNavigationHistoryEntryIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NavigationHistoryEntry_SameDocument
//go:noescape
func GetNavigationHistoryEntrySameDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_NavigationHistoryEntry_GetState
//go:noescape
func HasNavigationHistoryEntryGetState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationHistoryEntry_GetState
//go:noescape
func NavigationHistoryEntryGetStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NavigationHistoryEntry_GetState
//go:noescape
func CallNavigationHistoryEntryGetState(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_NavigationHistoryEntry_GetState
//go:noescape
func TryNavigationHistoryEntryGetState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_NavigationUpdateCurrentEntryOptions
//go:noescape
func NavigationUpdateCurrentEntryOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationUpdateCurrentEntryOptions
//go:noescape
func NavigationUpdateCurrentEntryOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_NavigationResult
//go:noescape
func NavigationResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_NavigationResult
//go:noescape
func NavigationResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
