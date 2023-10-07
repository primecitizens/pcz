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

//go:wasmimport plat/js/webext/downloads store_BooleanDelta
//go:noescape
func BooleanDeltaJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_BooleanDelta
//go:noescape
func BooleanDeltaJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads constof_DangerType
//go:noescape
func ConstOfDangerType(str js.Ref) uint32

//go:wasmimport plat/js/webext/downloads store_DoubleDelta
//go:noescape
func DoubleDeltaJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_DoubleDelta
//go:noescape
func DoubleDeltaJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_StringDelta
//go:noescape
func StringDeltaJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_StringDelta
//go:noescape
func StringDeltaJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_DownloadDelta
//go:noescape
func DownloadDeltaJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_DownloadDelta
//go:noescape
func DownloadDeltaJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads constof_State
//go:noescape
func ConstOfState(str js.Ref) uint32

//go:wasmimport plat/js/webext/downloads constof_InterruptReason
//go:noescape
func ConstOfInterruptReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/downloads store_DownloadItem
//go:noescape
func DownloadItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_DownloadItem
//go:noescape
func DownloadItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads constof_FilenameConflictAction
//go:noescape
func ConstOfFilenameConflictAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/downloads constof_HttpMethod
//go:noescape
func ConstOfHttpMethod(str js.Ref) uint32

//go:wasmimport plat/js/webext/downloads store_HeaderNameValuePair
//go:noescape
func HeaderNameValuePairJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_HeaderNameValuePair
//go:noescape
func HeaderNameValuePairJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_DownloadOptions
//go:noescape
func DownloadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_DownloadOptions
//go:noescape
func DownloadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_DownloadQuery
//go:noescape
func DownloadQueryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_DownloadQuery
//go:noescape
func DownloadQueryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_FilenameSuggestion
//go:noescape
func FilenameSuggestionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_FilenameSuggestion
//go:noescape
func FilenameSuggestionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_GetFileIconOptions
//go:noescape
func GetFileIconOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_GetFileIconOptions
//go:noescape
func GetFileIconOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads store_UiOptions
//go:noescape
func UiOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/downloads load_UiOptions
//go:noescape
func UiOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/downloads has_AcceptDanger
//go:noescape
func HasFuncAcceptDanger() js.Ref

//go:wasmimport plat/js/webext/downloads func_AcceptDanger
//go:noescape
func FuncAcceptDanger(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_AcceptDanger
//go:noescape
func CallAcceptDanger(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_AcceptDanger
//go:noescape
func TryAcceptDanger(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Cancel
//go:noescape
func HasFuncCancel() js.Ref

//go:wasmimport plat/js/webext/downloads func_Cancel
//go:noescape
func FuncCancel(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Cancel
//go:noescape
func CallCancel(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_Cancel
//go:noescape
func TryCancel(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Download
//go:noescape
func HasFuncDownload() js.Ref

//go:wasmimport plat/js/webext/downloads func_Download
//go:noescape
func FuncDownload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Download
//go:noescape
func CallDownload(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_Download
//go:noescape
func TryDownload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Erase
//go:noescape
func HasFuncErase() js.Ref

//go:wasmimport plat/js/webext/downloads func_Erase
//go:noescape
func FuncErase(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Erase
//go:noescape
func CallErase(
	retPtr unsafe.Pointer,
	query unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_Erase
//go:noescape
func TryErase(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_GetFileIcon
//go:noescape
func HasFuncGetFileIcon() js.Ref

//go:wasmimport plat/js/webext/downloads func_GetFileIcon
//go:noescape
func FuncGetFileIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_GetFileIcon
//go:noescape
func CallGetFileIcon(
	retPtr unsafe.Pointer,
	downloadId int32,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_GetFileIcon
//go:noescape
func TryGetFileIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/downloads func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/downloads func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/downloads func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/downloads func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/downloads func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/downloads func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OnDeterminingFilename
//go:noescape
func HasFuncOnDeterminingFilename() js.Ref

//go:wasmimport plat/js/webext/downloads func_OnDeterminingFilename
//go:noescape
func FuncOnDeterminingFilename(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OnDeterminingFilename
//go:noescape
func CallOnDeterminingFilename(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OnDeterminingFilename
//go:noescape
func TryOnDeterminingFilename(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OffDeterminingFilename
//go:noescape
func HasFuncOffDeterminingFilename() js.Ref

//go:wasmimport plat/js/webext/downloads func_OffDeterminingFilename
//go:noescape
func FuncOffDeterminingFilename(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OffDeterminingFilename
//go:noescape
func CallOffDeterminingFilename(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OffDeterminingFilename
//go:noescape
func TryOffDeterminingFilename(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_HasOnDeterminingFilename
//go:noescape
func HasFuncHasOnDeterminingFilename() js.Ref

//go:wasmimport plat/js/webext/downloads func_HasOnDeterminingFilename
//go:noescape
func FuncHasOnDeterminingFilename(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_HasOnDeterminingFilename
//go:noescape
func CallHasOnDeterminingFilename(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_HasOnDeterminingFilename
//go:noescape
func TryHasOnDeterminingFilename(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OnErased
//go:noescape
func HasFuncOnErased() js.Ref

//go:wasmimport plat/js/webext/downloads func_OnErased
//go:noescape
func FuncOnErased(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OnErased
//go:noescape
func CallOnErased(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OnErased
//go:noescape
func TryOnErased(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_OffErased
//go:noescape
func HasFuncOffErased() js.Ref

//go:wasmimport plat/js/webext/downloads func_OffErased
//go:noescape
func FuncOffErased(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_OffErased
//go:noescape
func CallOffErased(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_OffErased
//go:noescape
func TryOffErased(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_HasOnErased
//go:noescape
func HasFuncHasOnErased() js.Ref

//go:wasmimport plat/js/webext/downloads func_HasOnErased
//go:noescape
func FuncHasOnErased(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_HasOnErased
//go:noescape
func CallHasOnErased(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/downloads try_HasOnErased
//go:noescape
func TryHasOnErased(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Open
//go:noescape
func HasFuncOpen() js.Ref

//go:wasmimport plat/js/webext/downloads func_Open
//go:noescape
func FuncOpen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Open
//go:noescape
func CallOpen(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_Open
//go:noescape
func TryOpen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Pause
//go:noescape
func HasFuncPause() js.Ref

//go:wasmimport plat/js/webext/downloads func_Pause
//go:noescape
func FuncPause(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Pause
//go:noescape
func CallPause(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_Pause
//go:noescape
func TryPause(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_RemoveFile
//go:noescape
func HasFuncRemoveFile() js.Ref

//go:wasmimport plat/js/webext/downloads func_RemoveFile
//go:noescape
func FuncRemoveFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_RemoveFile
//go:noescape
func CallRemoveFile(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_RemoveFile
//go:noescape
func TryRemoveFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Resume
//go:noescape
func HasFuncResume() js.Ref

//go:wasmimport plat/js/webext/downloads func_Resume
//go:noescape
func FuncResume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Resume
//go:noescape
func CallResume(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_Resume
//go:noescape
func TryResume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Search
//go:noescape
func HasFuncSearch() js.Ref

//go:wasmimport plat/js/webext/downloads func_Search
//go:noescape
func FuncSearch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Search
//go:noescape
func CallSearch(
	retPtr unsafe.Pointer,
	query unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_Search
//go:noescape
func TrySearch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_SetShelfEnabled
//go:noescape
func HasFuncSetShelfEnabled() js.Ref

//go:wasmimport plat/js/webext/downloads func_SetShelfEnabled
//go:noescape
func FuncSetShelfEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_SetShelfEnabled
//go:noescape
func CallSetShelfEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/downloads try_SetShelfEnabled
//go:noescape
func TrySetShelfEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_SetUiOptions
//go:noescape
func HasFuncSetUiOptions() js.Ref

//go:wasmimport plat/js/webext/downloads func_SetUiOptions
//go:noescape
func FuncSetUiOptions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_SetUiOptions
//go:noescape
func CallSetUiOptions(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_SetUiOptions
//go:noescape
func TrySetUiOptions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_Show
//go:noescape
func HasFuncShow() js.Ref

//go:wasmimport plat/js/webext/downloads func_Show
//go:noescape
func FuncShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_Show
//go:noescape
func CallShow(
	retPtr unsafe.Pointer,
	downloadId int32)

//go:wasmimport plat/js/webext/downloads try_Show
//go:noescape
func TryShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	downloadId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/downloads has_ShowDefaultFolder
//go:noescape
func HasFuncShowDefaultFolder() js.Ref

//go:wasmimport plat/js/webext/downloads func_ShowDefaultFolder
//go:noescape
func FuncShowDefaultFolder(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads call_ShowDefaultFolder
//go:noescape
func CallShowDefaultFolder(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/downloads try_ShowDefaultFolder
//go:noescape
func TryShowDefaultFolder(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
