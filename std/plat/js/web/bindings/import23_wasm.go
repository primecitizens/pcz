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

//go:wasmimport plat/js/web get_IdleDeadline_DidTimeout
//go:noescape
func GetIdleDeadlineDidTimeout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IdleDeadline_TimeRemaining
//go:noescape
func HasFuncIdleDeadlineTimeRemaining(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IdleDeadline_TimeRemaining
//go:noescape
func FuncIdleDeadlineTimeRemaining(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemWritableFileStreamWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Write
//go:noescape
func FuncFileSystemWritableFileStreamWrite(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemWritableFileStreamSeek(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Seek
//go:noescape
func FuncFileSystemWritableFileStreamSeek(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemWritableFileStreamTruncate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemWritableFileStream_Truncate
//go:noescape
func FuncFileSystemWritableFileStreamTruncate(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Read
//go:noescape
func FuncFileSystemSyncAccessHandleRead(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleRead1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Read1
//go:noescape
func FuncFileSystemSyncAccessHandleRead1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Write
//go:noescape
func FuncFileSystemSyncAccessHandleWrite(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleWrite1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Write1
//go:noescape
func FuncFileSystemSyncAccessHandleWrite1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleTruncate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Truncate
//go:noescape
func FuncFileSystemSyncAccessHandleTruncate(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleGetSize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_GetSize
//go:noescape
func FuncFileSystemSyncAccessHandleGetSize(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleFlush(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Flush
//go:noescape
func FuncFileSystemSyncAccessHandleFlush(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemSyncAccessHandleClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemSyncAccessHandle_Close
//go:noescape
func FuncFileSystemSyncAccessHandleClose(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemFileHandleGetFile(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_GetFile
//go:noescape
func FuncFileSystemFileHandleGetFile(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemFileHandleCreateWritable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateWritable
//go:noescape
func FuncFileSystemFileHandleCreateWritable(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemFileHandleCreateWritable1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateWritable1
//go:noescape
func FuncFileSystemFileHandleCreateWritable1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemFileHandleCreateSyncAccessHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemFileHandle_CreateSyncAccessHandle
//go:noescape
func FuncFileSystemFileHandleCreateSyncAccessHandle(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleGetFileHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetFileHandle
//go:noescape
func FuncFileSystemDirectoryHandleGetFileHandle(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleGetFileHandle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetFileHandle1
//go:noescape
func FuncFileSystemDirectoryHandleGetFileHandle1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleGetDirectoryHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetDirectoryHandle
//go:noescape
func FuncFileSystemDirectoryHandleGetDirectoryHandle(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleGetDirectoryHandle1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_GetDirectoryHandle1
//go:noescape
func FuncFileSystemDirectoryHandleGetDirectoryHandle1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleRemoveEntry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_RemoveEntry
//go:noescape
func FuncFileSystemDirectoryHandleRemoveEntry(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleRemoveEntry1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_RemoveEntry1
//go:noescape
func FuncFileSystemDirectoryHandleRemoveEntry1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFileSystemDirectoryHandleResolve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryHandle_Resolve
//go:noescape
func FuncFileSystemDirectoryHandleResolve(this js.Ref, fn unsafe.Pointer)

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
func HasFuncFontDataBlob(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontData_Blob
//go:noescape
func FuncFontDataBlob(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDigitalGoodsServiceGetDetails(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_GetDetails
//go:noescape
func FuncDigitalGoodsServiceGetDetails(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDigitalGoodsServiceListPurchases(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_ListPurchases
//go:noescape
func FuncDigitalGoodsServiceListPurchases(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDigitalGoodsServiceListPurchaseHistory(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_ListPurchaseHistory
//go:noescape
func FuncDigitalGoodsServiceListPurchaseHistory(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDigitalGoodsServiceConsume(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DigitalGoodsService_Consume
//go:noescape
func FuncDigitalGoodsServiceConsume(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryGo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Go
//go:noescape
func FuncHistoryGo(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryGo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Go1
//go:noescape
func FuncHistoryGo1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryBack(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Back
//go:noescape
func FuncHistoryBack(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryForward(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_Forward
//go:noescape
func FuncHistoryForward(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryPushState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_PushState
//go:noescape
func FuncHistoryPushState(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryPushState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_PushState1
//go:noescape
func FuncHistoryPushState1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryReplaceState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_ReplaceState
//go:noescape
func FuncHistoryReplaceState(this js.Ref, fn unsafe.Pointer)

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
func HasFuncHistoryReplaceState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_History_ReplaceState1
//go:noescape
func FuncHistoryReplaceState1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncNavigationHistoryEntryGetState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_NavigationHistoryEntry_GetState
//go:noescape
func FuncNavigationHistoryEntryGetState(this js.Ref, fn unsafe.Pointer)

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
