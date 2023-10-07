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

//go:wasmimport plat/js/web get_WindowClient_VisibilityState
//go:noescape
func GetWindowClientVisibilityState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WindowClient_Focused
//go:noescape
func GetWindowClientFocused(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WindowClient_AncestorOrigins
//go:noescape
func GetWindowClientAncestorOrigins(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowClient_Focus
//go:noescape
func HasFuncWindowClientFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowClient_Focus
//go:noescape
func FuncWindowClientFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowClient_Focus
//go:noescape
func CallWindowClientFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WindowClient_Focus
//go:noescape
func TryWindowClientFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowClient_Navigate
//go:noescape
func HasFuncWindowClientNavigate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowClient_Navigate
//go:noescape
func FuncWindowClientNavigate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_WindowClient_Navigate
//go:noescape
func CallWindowClientNavigate(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_WindowClient_Navigate
//go:noescape
func TryWindowClientNavigate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Clients_Get
//go:noescape
func HasFuncClientsGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clients_Get
//go:noescape
func FuncClientsGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clients_Get
//go:noescape
func CallClientsGet(
	this js.Ref, retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/web try_Clients_Get
//go:noescape
func TryClientsGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Clients_MatchAll
//go:noescape
func HasFuncClientsMatchAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clients_MatchAll
//go:noescape
func FuncClientsMatchAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clients_MatchAll
//go:noescape
func CallClientsMatchAll(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Clients_MatchAll
//go:noescape
func TryClientsMatchAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Clients_MatchAll1
//go:noescape
func HasFuncClientsMatchAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clients_MatchAll1
//go:noescape
func FuncClientsMatchAll1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clients_MatchAll1
//go:noescape
func CallClientsMatchAll1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Clients_MatchAll1
//go:noescape
func TryClientsMatchAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Clients_OpenWindow
//go:noescape
func HasFuncClientsOpenWindow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clients_OpenWindow
//go:noescape
func FuncClientsOpenWindow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clients_OpenWindow
//go:noescape
func CallClientsOpenWindow(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Clients_OpenWindow
//go:noescape
func TryClientsOpenWindow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Clients_Claim
//go:noescape
func HasFuncClientsClaim(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clients_Claim
//go:noescape
func FuncClientsClaim(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clients_Claim
//go:noescape
func CallClientsClaim(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Clients_Claim
//go:noescape
func TryClientsClaim(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PresentationStyle
//go:noescape
func ConstOfPresentationStyle(str js.Ref) uint32

//go:wasmimport plat/js/web store_ClipboardItemOptions
//go:noescape
func ClipboardItemOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ClipboardItemOptions
//go:noescape
func ClipboardItemOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ClipboardItem_ClipboardItem
//go:noescape
func NewClipboardItemByClipboardItem(
	items js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ClipboardItem_ClipboardItem1
//go:noescape
func NewClipboardItemByClipboardItem1(
	items js.Ref) js.Ref

//go:wasmimport plat/js/web get_ClipboardItem_PresentationStyle
//go:noescape
func GetClipboardItemPresentationStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ClipboardItem_Types
//go:noescape
func GetClipboardItemTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ClipboardItem_GetType
//go:noescape
func HasFuncClipboardItemGetType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ClipboardItem_GetType
//go:noescape
func FuncClipboardItemGetType(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_ClipboardItem_GetType
//go:noescape
func CallClipboardItemGetType(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_ClipboardItem_GetType
//go:noescape
func TryClipboardItemGetType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Clipboard_Read
//go:noescape
func HasFuncClipboardRead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clipboard_Read
//go:noescape
func FuncClipboardRead(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clipboard_Read
//go:noescape
func CallClipboardRead(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Clipboard_Read
//go:noescape
func TryClipboardRead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Clipboard_ReadText
//go:noescape
func HasFuncClipboardReadText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clipboard_ReadText
//go:noescape
func FuncClipboardReadText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clipboard_ReadText
//go:noescape
func CallClipboardReadText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Clipboard_ReadText
//go:noescape
func TryClipboardReadText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Clipboard_Write
//go:noescape
func HasFuncClipboardWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clipboard_Write
//go:noescape
func FuncClipboardWrite(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clipboard_Write
//go:noescape
func CallClipboardWrite(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Clipboard_Write
//go:noescape
func TryClipboardWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Clipboard_WriteText
//go:noescape
func HasFuncClipboardWriteText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Clipboard_WriteText
//go:noescape
func FuncClipboardWriteText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Clipboard_WriteText
//go:noescape
func CallClipboardWriteText(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Clipboard_WriteText
//go:noescape
func TryClipboardWriteText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_PermissionState
//go:noescape
func ConstOfPermissionState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_FileSystemPermissionMode
//go:noescape
func ConstOfFileSystemPermissionMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_FileSystemHandlePermissionDescriptor
//go:noescape
func FileSystemHandlePermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemHandlePermissionDescriptor
//go:noescape
func FileSystemHandlePermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_FileSystemHandleKind
//go:noescape
func ConstOfFileSystemHandleKind(str js.Ref) uint32

//go:wasmimport plat/js/web get_FileSystemHandle_Kind
//go:noescape
func GetFileSystemHandleKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemHandle_Name
//go:noescape
func GetFileSystemHandleName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemHandle_IsSameEntry
//go:noescape
func HasFuncFileSystemHandleIsSameEntry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemHandle_IsSameEntry
//go:noescape
func FuncFileSystemHandleIsSameEntry(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemHandle_IsSameEntry
//go:noescape
func CallFileSystemHandleIsSameEntry(
	this js.Ref, retPtr unsafe.Pointer,
	other js.Ref)

//go:wasmimport plat/js/web try_FileSystemHandle_IsSameEntry
//go:noescape
func TryFileSystemHandleIsSameEntry(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemHandle_QueryPermission
//go:noescape
func HasFuncFileSystemHandleQueryPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemHandle_QueryPermission
//go:noescape
func FuncFileSystemHandleQueryPermission(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemHandle_QueryPermission
//go:noescape
func CallFileSystemHandleQueryPermission(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemHandle_QueryPermission
//go:noescape
func TryFileSystemHandleQueryPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemHandle_QueryPermission1
//go:noescape
func HasFuncFileSystemHandleQueryPermission1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemHandle_QueryPermission1
//go:noescape
func FuncFileSystemHandleQueryPermission1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemHandle_QueryPermission1
//go:noescape
func CallFileSystemHandleQueryPermission1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemHandle_QueryPermission1
//go:noescape
func TryFileSystemHandleQueryPermission1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemHandle_RequestPermission
//go:noescape
func HasFuncFileSystemHandleRequestPermission(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemHandle_RequestPermission
//go:noescape
func FuncFileSystemHandleRequestPermission(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemHandle_RequestPermission
//go:noescape
func CallFileSystemHandleRequestPermission(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemHandle_RequestPermission
//go:noescape
func TryFileSystemHandleRequestPermission(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemHandle_RequestPermission1
//go:noescape
func HasFuncFileSystemHandleRequestPermission1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemHandle_RequestPermission1
//go:noescape
func FuncFileSystemHandleRequestPermission1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemHandle_RequestPermission1
//go:noescape
func CallFileSystemHandleRequestPermission1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemHandle_RequestPermission1
//go:noescape
func TryFileSystemHandleRequestPermission1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryReader_ReadEntries
//go:noescape
func HasFuncFileSystemDirectoryReaderReadEntries(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryReader_ReadEntries
//go:noescape
func FuncFileSystemDirectoryReaderReadEntries(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryReader_ReadEntries
//go:noescape
func CallFileSystemDirectoryReaderReadEntries(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryReader_ReadEntries
//go:noescape
func TryFileSystemDirectoryReaderReadEntries(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryReader_ReadEntries1
//go:noescape
func HasFuncFileSystemDirectoryReaderReadEntries1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryReader_ReadEntries1
//go:noescape
func FuncFileSystemDirectoryReaderReadEntries1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryReader_ReadEntries1
//go:noescape
func CallFileSystemDirectoryReaderReadEntries1(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryReader_ReadEntries1
//go:noescape
func TryFileSystemDirectoryReaderReadEntries1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_FileSystemFlags
//go:noescape
func FileSystemFlagsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FileSystemFlags
//go:noescape
func FileSystemFlagsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_CreateReader
//go:noescape
func HasFuncFileSystemDirectoryEntryCreateReader(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_CreateReader
//go:noescape
func FuncFileSystemDirectoryEntryCreateReader(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_CreateReader
//go:noescape
func CallFileSystemDirectoryEntryCreateReader(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_CreateReader
//go:noescape
func TryFileSystemDirectoryEntryCreateReader(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetFile
//go:noescape
func HasFuncFileSystemDirectoryEntryGetFile(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetFile
//go:noescape
func FuncFileSystemDirectoryEntryGetFile(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetFile
//go:noescape
func CallFileSystemDirectoryEntryGetFile(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetFile
//go:noescape
func TryFileSystemDirectoryEntryGetFile(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetFile1
//go:noescape
func HasFuncFileSystemDirectoryEntryGetFile1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetFile1
//go:noescape
func FuncFileSystemDirectoryEntryGetFile1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetFile1
//go:noescape
func CallFileSystemDirectoryEntryGetFile1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetFile1
//go:noescape
func TryFileSystemDirectoryEntryGetFile1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetFile2
//go:noescape
func HasFuncFileSystemDirectoryEntryGetFile2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetFile2
//go:noescape
func FuncFileSystemDirectoryEntryGetFile2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetFile2
//go:noescape
func CallFileSystemDirectoryEntryGetFile2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetFile2
//go:noescape
func TryFileSystemDirectoryEntryGetFile2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetFile3
//go:noescape
func HasFuncFileSystemDirectoryEntryGetFile3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetFile3
//go:noescape
func FuncFileSystemDirectoryEntryGetFile3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetFile3
//go:noescape
func CallFileSystemDirectoryEntryGetFile3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetFile3
//go:noescape
func TryFileSystemDirectoryEntryGetFile3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetFile4
//go:noescape
func HasFuncFileSystemDirectoryEntryGetFile4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetFile4
//go:noescape
func FuncFileSystemDirectoryEntryGetFile4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetFile4
//go:noescape
func CallFileSystemDirectoryEntryGetFile4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetFile4
//go:noescape
func TryFileSystemDirectoryEntryGetFile4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetDirectory
//go:noescape
func HasFuncFileSystemDirectoryEntryGetDirectory(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetDirectory
//go:noescape
func FuncFileSystemDirectoryEntryGetDirectory(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetDirectory
//go:noescape
func CallFileSystemDirectoryEntryGetDirectory(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetDirectory
//go:noescape
func TryFileSystemDirectoryEntryGetDirectory(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetDirectory1
//go:noescape
func HasFuncFileSystemDirectoryEntryGetDirectory1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetDirectory1
//go:noescape
func FuncFileSystemDirectoryEntryGetDirectory1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetDirectory1
//go:noescape
func CallFileSystemDirectoryEntryGetDirectory1(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetDirectory1
//go:noescape
func TryFileSystemDirectoryEntryGetDirectory1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetDirectory2
//go:noescape
func HasFuncFileSystemDirectoryEntryGetDirectory2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetDirectory2
//go:noescape
func FuncFileSystemDirectoryEntryGetDirectory2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetDirectory2
//go:noescape
func CallFileSystemDirectoryEntryGetDirectory2(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetDirectory2
//go:noescape
func TryFileSystemDirectoryEntryGetDirectory2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetDirectory3
//go:noescape
func HasFuncFileSystemDirectoryEntryGetDirectory3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetDirectory3
//go:noescape
func FuncFileSystemDirectoryEntryGetDirectory3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetDirectory3
//go:noescape
func CallFileSystemDirectoryEntryGetDirectory3(
	this js.Ref, retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetDirectory3
//go:noescape
func TryFileSystemDirectoryEntryGetDirectory3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemDirectoryEntry_GetDirectory4
//go:noescape
func HasFuncFileSystemDirectoryEntryGetDirectory4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemDirectoryEntry_GetDirectory4
//go:noescape
func FuncFileSystemDirectoryEntryGetDirectory4(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemDirectoryEntry_GetDirectory4
//go:noescape
func CallFileSystemDirectoryEntryGetDirectory4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemDirectoryEntry_GetDirectory4
//go:noescape
func TryFileSystemDirectoryEntryGetDirectory4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystem_Name
//go:noescape
func GetFileSystemName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystem_Root
//go:noescape
func GetFileSystemRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemEntry_IsFile
//go:noescape
func GetFileSystemEntryIsFile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemEntry_IsDirectory
//go:noescape
func GetFileSystemEntryIsDirectory(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemEntry_Name
//go:noescape
func GetFileSystemEntryName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemEntry_FullPath
//go:noescape
func GetFileSystemEntryFullPath(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileSystemEntry_Filesystem
//go:noescape
func GetFileSystemEntryFilesystem(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemEntry_GetParent
//go:noescape
func HasFuncFileSystemEntryGetParent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemEntry_GetParent
//go:noescape
func FuncFileSystemEntryGetParent(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemEntry_GetParent
//go:noescape
func CallFileSystemEntryGetParent(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemEntry_GetParent
//go:noescape
func TryFileSystemEntryGetParent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemEntry_GetParent1
//go:noescape
func HasFuncFileSystemEntryGetParent1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemEntry_GetParent1
//go:noescape
func FuncFileSystemEntryGetParent1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemEntry_GetParent1
//go:noescape
func CallFileSystemEntryGetParent1(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_FileSystemEntry_GetParent1
//go:noescape
func TryFileSystemEntryGetParent1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FileSystemEntry_GetParent2
//go:noescape
func HasFuncFileSystemEntryGetParent2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileSystemEntry_GetParent2
//go:noescape
func FuncFileSystemEntryGetParent2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileSystemEntry_GetParent2
//go:noescape
func CallFileSystemEntryGetParent2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FileSystemEntry_GetParent2
//go:noescape
func TryFileSystemEntryGetParent2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransferItem_Kind
//go:noescape
func GetDataTransferItemKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransferItem_Type
//go:noescape
func GetDataTransferItemType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItem_GetAsString
//go:noescape
func HasFuncDataTransferItemGetAsString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItem_GetAsString
//go:noescape
func FuncDataTransferItemGetAsString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItem_GetAsString
//go:noescape
func CallDataTransferItemGetAsString(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_DataTransferItem_GetAsString
//go:noescape
func TryDataTransferItemGetAsString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItem_GetAsFile
//go:noescape
func HasFuncDataTransferItemGetAsFile(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItem_GetAsFile
//go:noescape
func FuncDataTransferItemGetAsFile(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItem_GetAsFile
//go:noescape
func CallDataTransferItemGetAsFile(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DataTransferItem_GetAsFile
//go:noescape
func TryDataTransferItemGetAsFile(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItem_GetAsFileSystemHandle
//go:noescape
func HasFuncDataTransferItemGetAsFileSystemHandle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItem_GetAsFileSystemHandle
//go:noescape
func FuncDataTransferItemGetAsFileSystemHandle(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItem_GetAsFileSystemHandle
//go:noescape
func CallDataTransferItemGetAsFileSystemHandle(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DataTransferItem_GetAsFileSystemHandle
//go:noescape
func TryDataTransferItemGetAsFileSystemHandle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItem_WebkitGetAsEntry
//go:noescape
func HasFuncDataTransferItemWebkitGetAsEntry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItem_WebkitGetAsEntry
//go:noescape
func FuncDataTransferItemWebkitGetAsEntry(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItem_WebkitGetAsEntry
//go:noescape
func CallDataTransferItemWebkitGetAsEntry(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DataTransferItem_WebkitGetAsEntry
//go:noescape
func TryDataTransferItemWebkitGetAsEntry(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransferItemList_Length
//go:noescape
func GetDataTransferItemListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItemList_Get
//go:noescape
func HasFuncDataTransferItemListGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItemList_Get
//go:noescape
func FuncDataTransferItemListGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItemList_Get
//go:noescape
func CallDataTransferItemListGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_DataTransferItemList_Get
//go:noescape
func TryDataTransferItemListGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItemList_Add
//go:noescape
func HasFuncDataTransferItemListAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItemList_Add
//go:noescape
func FuncDataTransferItemListAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItemList_Add
//go:noescape
func CallDataTransferItemListAdd(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref,
	typ js.Ref)

//go:wasmimport plat/js/web try_DataTransferItemList_Add
//go:noescape
func TryDataTransferItemListAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItemList_Add1
//go:noescape
func HasFuncDataTransferItemListAdd1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItemList_Add1
//go:noescape
func FuncDataTransferItemListAdd1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItemList_Add1
//go:noescape
func CallDataTransferItemListAdd1(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_DataTransferItemList_Add1
//go:noescape
func TryDataTransferItemListAdd1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItemList_Remove
//go:noescape
func HasFuncDataTransferItemListRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItemList_Remove
//go:noescape
func FuncDataTransferItemListRemove(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItemList_Remove
//go:noescape
func CallDataTransferItemListRemove(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_DataTransferItemList_Remove
//go:noescape
func TryDataTransferItemListRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransferItemList_Clear
//go:noescape
func HasFuncDataTransferItemListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransferItemList_Clear
//go:noescape
func FuncDataTransferItemListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransferItemList_Clear
//go:noescape
func CallDataTransferItemListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DataTransferItemList_Clear
//go:noescape
func TryDataTransferItemListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FileList_Length
//go:noescape
func GetFileListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FileList_Item
//go:noescape
func HasFuncFileListItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FileList_Item
//go:noescape
func FuncFileListItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_FileList_Item
//go:noescape
func CallFileListItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_FileList_Item
//go:noescape
func TryFileListItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransfer_DropEffect
//go:noescape
func GetDataTransferDropEffect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DataTransfer_DropEffect
//go:noescape
func SetDataTransferDropEffect(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_DataTransfer_EffectAllowed
//go:noescape
func GetDataTransferEffectAllowed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DataTransfer_EffectAllowed
//go:noescape
func SetDataTransferEffectAllowed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_DataTransfer_Items
//go:noescape
func GetDataTransferItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransfer_Types
//go:noescape
func GetDataTransferTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DataTransfer_Files
//go:noescape
func GetDataTransferFiles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransfer_SetDragImage
//go:noescape
func HasFuncDataTransferSetDragImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransfer_SetDragImage
//go:noescape
func FuncDataTransferSetDragImage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransfer_SetDragImage
//go:noescape
func CallDataTransferSetDragImage(
	this js.Ref, retPtr unsafe.Pointer,
	image js.Ref,
	x int32,
	y int32)

//go:wasmimport plat/js/web try_DataTransfer_SetDragImage
//go:noescape
func TryDataTransferSetDragImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	image js.Ref,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransfer_GetData
//go:noescape
func HasFuncDataTransferGetData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransfer_GetData
//go:noescape
func FuncDataTransferGetData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransfer_GetData
//go:noescape
func CallDataTransferGetData(
	this js.Ref, retPtr unsafe.Pointer,
	format js.Ref)

//go:wasmimport plat/js/web try_DataTransfer_GetData
//go:noescape
func TryDataTransferGetData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransfer_SetData
//go:noescape
func HasFuncDataTransferSetData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransfer_SetData
//go:noescape
func FuncDataTransferSetData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransfer_SetData
//go:noescape
func CallDataTransferSetData(
	this js.Ref, retPtr unsafe.Pointer,
	format js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_DataTransfer_SetData
//go:noescape
func TryDataTransferSetData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransfer_ClearData
//go:noescape
func HasFuncDataTransferClearData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransfer_ClearData
//go:noescape
func FuncDataTransferClearData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransfer_ClearData
//go:noescape
func CallDataTransferClearData(
	this js.Ref, retPtr unsafe.Pointer,
	format js.Ref)

//go:wasmimport plat/js/web try_DataTransfer_ClearData
//go:noescape
func TryDataTransferClearData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	format js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DataTransfer_ClearData1
//go:noescape
func HasFuncDataTransferClearData1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DataTransfer_ClearData1
//go:noescape
func FuncDataTransferClearData1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_DataTransfer_ClearData1
//go:noescape
func CallDataTransferClearData1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DataTransfer_ClearData1
//go:noescape
func TryDataTransferClearData1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ClipboardEventInit
//go:noescape
func ClipboardEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ClipboardEventInit
//go:noescape
func ClipboardEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ClipboardEvent_ClipboardEvent
//go:noescape
func NewClipboardEventByClipboardEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_ClipboardEvent_ClipboardEvent1
//go:noescape
func NewClipboardEventByClipboardEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_ClipboardEvent_ClipboardData
//go:noescape
func GetClipboardEventClipboardData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_ClipboardPermissionDescriptor
//go:noescape
func ClipboardPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ClipboardPermissionDescriptor
//go:noescape
func ClipboardPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CloseEventInit
//go:noescape
func CloseEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CloseEventInit
//go:noescape
func CloseEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CloseEvent_CloseEvent
//go:noescape
func NewCloseEventByCloseEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CloseEvent_CloseEvent1
//go:noescape
func NewCloseEventByCloseEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_CloseEvent_WasClean
//go:noescape
func GetCloseEventWasClean(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CloseEvent_Code
//go:noescape
func GetCloseEventCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CloseEvent_Reason
//go:noescape
func GetCloseEventReason(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_CollectedClientAdditionalPaymentData
//go:noescape
func CollectedClientAdditionalPaymentDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CollectedClientAdditionalPaymentData
//go:noescape
func CollectedClientAdditionalPaymentDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CollectedClientData
//go:noescape
func CollectedClientDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CollectedClientData
//go:noescape
func CollectedClientDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CollectedClientPaymentData
//go:noescape
func CollectedClientPaymentDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CollectedClientPaymentData
//go:noescape
func CollectedClientPaymentDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ColorGamut
//go:noescape
func ConstOfColorGamut(str js.Ref) uint32

//go:wasmimport plat/js/web store_ColorSelectionOptions
//go:noescape
func ColorSelectionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ColorSelectionOptions
//go:noescape
func ColorSelectionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ColorSelectionResult
//go:noescape
func ColorSelectionResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ColorSelectionResult
//go:noescape
func ColorSelectionResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ColorSpaceConversion
//go:noescape
func ConstOfColorSpaceConversion(str js.Ref) uint32

//go:wasmimport plat/js/web constof_CompositeOperation
//go:noescape
func ConstOfCompositeOperation(str js.Ref) uint32

//go:wasmimport plat/js/web store_WindowPostMessageOptions
//go:noescape
func WindowPostMessageOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WindowPostMessageOptions
//go:noescape
func WindowPostMessageOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaQueryList_Media
//go:noescape
func GetMediaQueryListMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaQueryList_Matches
//go:noescape
func GetMediaQueryListMatches(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaQueryList_AddListener
//go:noescape
func HasFuncMediaQueryListAddListener(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaQueryList_AddListener
//go:noescape
func FuncMediaQueryListAddListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaQueryList_AddListener
//go:noescape
func CallMediaQueryListAddListener(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_MediaQueryList_AddListener
//go:noescape
func TryMediaQueryListAddListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaQueryList_RemoveListener
//go:noescape
func HasFuncMediaQueryListRemoveListener(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaQueryList_RemoveListener
//go:noescape
func FuncMediaQueryListRemoveListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MediaQueryList_RemoveListener
//go:noescape
func CallMediaQueryListRemoveListener(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/web try_MediaQueryList_RemoveListener
//go:noescape
func TryMediaQueryListRemoveListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
