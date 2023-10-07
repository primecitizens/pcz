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

//go:wasmimport plat/js/webext/filesystemprovider store_AbortRequestedOptions
//go:noescape
func AbortRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_AbortRequestedOptions
//go:noescape
func AbortRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_Action
//go:noescape
func ActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_Action
//go:noescape
func ActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_AddWatcherRequestedOptions
//go:noescape
func AddWatcherRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_AddWatcherRequestedOptions
//go:noescape
func AddWatcherRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider constof_ChangeType
//go:noescape
func ConstOfChangeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filesystemprovider store_Change
//go:noescape
func ChangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_Change
//go:noescape
func ChangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_CloseFileRequestedOptions
//go:noescape
func CloseFileRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_CloseFileRequestedOptions
//go:noescape
func CloseFileRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_CloudIdentifier
//go:noescape
func CloudIdentifierJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_CloudIdentifier
//go:noescape
func CloudIdentifierJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider constof_CommonActionId
//go:noescape
func ConstOfCommonActionId(str js.Ref) uint32

//go:wasmimport plat/js/webext/filesystemprovider store_ConfigureRequestedOptions
//go:noescape
func ConfigureRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_ConfigureRequestedOptions
//go:noescape
func ConfigureRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_CopyEntryRequestedOptions
//go:noescape
func CopyEntryRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_CopyEntryRequestedOptions
//go:noescape
func CopyEntryRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_CreateDirectoryRequestedOptions
//go:noescape
func CreateDirectoryRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_CreateDirectoryRequestedOptions
//go:noescape
func CreateDirectoryRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_CreateFileRequestedOptions
//go:noescape
func CreateFileRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_CreateFileRequestedOptions
//go:noescape
func CreateFileRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_DeleteEntryRequestedOptions
//go:noescape
func DeleteEntryRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_DeleteEntryRequestedOptions
//go:noescape
func DeleteEntryRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_EntryMetadata
//go:noescape
func EntryMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_EntryMetadata
//go:noescape
func EntryMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_ExecuteActionRequestedOptions
//go:noescape
func ExecuteActionRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_ExecuteActionRequestedOptions
//go:noescape
func ExecuteActionRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider constof_OpenFileMode
//go:noescape
func ConstOfOpenFileMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/filesystemprovider store_OpenedFile
//go:noescape
func OpenedFileJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_OpenedFile
//go:noescape
func OpenedFileJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_Watcher
//go:noescape
func WatcherJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_Watcher
//go:noescape
func WatcherJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_FileSystemInfo
//go:noescape
func FileSystemInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_FileSystemInfo
//go:noescape
func FileSystemInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_GetActionsRequestedOptions
//go:noescape
func GetActionsRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_GetActionsRequestedOptions
//go:noescape
func GetActionsRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_GetMetadataRequestedOptions
//go:noescape
func GetMetadataRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_GetMetadataRequestedOptions
//go:noescape
func GetMetadataRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_MountOptions
//go:noescape
func MountOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_MountOptions
//go:noescape
func MountOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_MoveEntryRequestedOptions
//go:noescape
func MoveEntryRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_MoveEntryRequestedOptions
//go:noescape
func MoveEntryRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_NotifyOptions
//go:noescape
func NotifyOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_NotifyOptions
//go:noescape
func NotifyOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_OpenFileRequestedOptions
//go:noescape
func OpenFileRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_OpenFileRequestedOptions
//go:noescape
func OpenFileRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider constof_ProviderError
//go:noescape
func ConstOfProviderError(str js.Ref) uint32

//go:wasmimport plat/js/webext/filesystemprovider store_ReadDirectoryRequestedOptions
//go:noescape
func ReadDirectoryRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_ReadDirectoryRequestedOptions
//go:noescape
func ReadDirectoryRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_ReadFileRequestedOptions
//go:noescape
func ReadFileRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_ReadFileRequestedOptions
//go:noescape
func ReadFileRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_RemoveWatcherRequestedOptions
//go:noescape
func RemoveWatcherRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_RemoveWatcherRequestedOptions
//go:noescape
func RemoveWatcherRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_TruncateRequestedOptions
//go:noescape
func TruncateRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_TruncateRequestedOptions
//go:noescape
func TruncateRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_UnmountOptions
//go:noescape
func UnmountOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_UnmountOptions
//go:noescape
func UnmountOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_UnmountRequestedOptions
//go:noescape
func UnmountRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_UnmountRequestedOptions
//go:noescape
func UnmountRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider store_WriteFileRequestedOptions
//go:noescape
func WriteFileRequestedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider load_WriteFileRequestedOptions
//go:noescape
func WriteFileRequestedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filesystemprovider has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_Mount
//go:noescape
func HasFuncMount() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_Mount
//go:noescape
func FuncMount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_Mount
//go:noescape
func CallMount(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider try_Mount
//go:noescape
func TryMount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_Notify
//go:noescape
func HasFuncNotify() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_Notify
//go:noescape
func FuncNotify(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_Notify
//go:noescape
func CallNotify(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider try_Notify
//go:noescape
func TryNotify(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnAbortRequested
//go:noescape
func HasFuncOnAbortRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnAbortRequested
//go:noescape
func FuncOnAbortRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnAbortRequested
//go:noescape
func CallOnAbortRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnAbortRequested
//go:noescape
func TryOnAbortRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffAbortRequested
//go:noescape
func HasFuncOffAbortRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffAbortRequested
//go:noescape
func FuncOffAbortRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffAbortRequested
//go:noescape
func CallOffAbortRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffAbortRequested
//go:noescape
func TryOffAbortRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnAbortRequested
//go:noescape
func HasFuncHasOnAbortRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnAbortRequested
//go:noescape
func FuncHasOnAbortRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnAbortRequested
//go:noescape
func CallHasOnAbortRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnAbortRequested
//go:noescape
func TryHasOnAbortRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnAddWatcherRequested
//go:noescape
func HasFuncOnAddWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnAddWatcherRequested
//go:noescape
func FuncOnAddWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnAddWatcherRequested
//go:noescape
func CallOnAddWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnAddWatcherRequested
//go:noescape
func TryOnAddWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffAddWatcherRequested
//go:noescape
func HasFuncOffAddWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffAddWatcherRequested
//go:noescape
func FuncOffAddWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffAddWatcherRequested
//go:noescape
func CallOffAddWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffAddWatcherRequested
//go:noescape
func TryOffAddWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnAddWatcherRequested
//go:noescape
func HasFuncHasOnAddWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnAddWatcherRequested
//go:noescape
func FuncHasOnAddWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnAddWatcherRequested
//go:noescape
func CallHasOnAddWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnAddWatcherRequested
//go:noescape
func TryHasOnAddWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnCloseFileRequested
//go:noescape
func HasFuncOnCloseFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnCloseFileRequested
//go:noescape
func FuncOnCloseFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnCloseFileRequested
//go:noescape
func CallOnCloseFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnCloseFileRequested
//go:noescape
func TryOnCloseFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffCloseFileRequested
//go:noescape
func HasFuncOffCloseFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffCloseFileRequested
//go:noescape
func FuncOffCloseFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffCloseFileRequested
//go:noescape
func CallOffCloseFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffCloseFileRequested
//go:noescape
func TryOffCloseFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnCloseFileRequested
//go:noescape
func HasFuncHasOnCloseFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnCloseFileRequested
//go:noescape
func FuncHasOnCloseFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnCloseFileRequested
//go:noescape
func CallHasOnCloseFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnCloseFileRequested
//go:noescape
func TryHasOnCloseFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnConfigureRequested
//go:noescape
func HasFuncOnConfigureRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnConfigureRequested
//go:noescape
func FuncOnConfigureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnConfigureRequested
//go:noescape
func CallOnConfigureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnConfigureRequested
//go:noescape
func TryOnConfigureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffConfigureRequested
//go:noescape
func HasFuncOffConfigureRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffConfigureRequested
//go:noescape
func FuncOffConfigureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffConfigureRequested
//go:noescape
func CallOffConfigureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffConfigureRequested
//go:noescape
func TryOffConfigureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnConfigureRequested
//go:noescape
func HasFuncHasOnConfigureRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnConfigureRequested
//go:noescape
func FuncHasOnConfigureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnConfigureRequested
//go:noescape
func CallHasOnConfigureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnConfigureRequested
//go:noescape
func TryHasOnConfigureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnCopyEntryRequested
//go:noescape
func HasFuncOnCopyEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnCopyEntryRequested
//go:noescape
func FuncOnCopyEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnCopyEntryRequested
//go:noescape
func CallOnCopyEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnCopyEntryRequested
//go:noescape
func TryOnCopyEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffCopyEntryRequested
//go:noescape
func HasFuncOffCopyEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffCopyEntryRequested
//go:noescape
func FuncOffCopyEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffCopyEntryRequested
//go:noescape
func CallOffCopyEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffCopyEntryRequested
//go:noescape
func TryOffCopyEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnCopyEntryRequested
//go:noescape
func HasFuncHasOnCopyEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnCopyEntryRequested
//go:noescape
func FuncHasOnCopyEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnCopyEntryRequested
//go:noescape
func CallHasOnCopyEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnCopyEntryRequested
//go:noescape
func TryHasOnCopyEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnCreateDirectoryRequested
//go:noescape
func HasFuncOnCreateDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnCreateDirectoryRequested
//go:noescape
func FuncOnCreateDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnCreateDirectoryRequested
//go:noescape
func CallOnCreateDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnCreateDirectoryRequested
//go:noescape
func TryOnCreateDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffCreateDirectoryRequested
//go:noescape
func HasFuncOffCreateDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffCreateDirectoryRequested
//go:noescape
func FuncOffCreateDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffCreateDirectoryRequested
//go:noescape
func CallOffCreateDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffCreateDirectoryRequested
//go:noescape
func TryOffCreateDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnCreateDirectoryRequested
//go:noescape
func HasFuncHasOnCreateDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnCreateDirectoryRequested
//go:noescape
func FuncHasOnCreateDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnCreateDirectoryRequested
//go:noescape
func CallHasOnCreateDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnCreateDirectoryRequested
//go:noescape
func TryHasOnCreateDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnCreateFileRequested
//go:noescape
func HasFuncOnCreateFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnCreateFileRequested
//go:noescape
func FuncOnCreateFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnCreateFileRequested
//go:noescape
func CallOnCreateFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnCreateFileRequested
//go:noescape
func TryOnCreateFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffCreateFileRequested
//go:noescape
func HasFuncOffCreateFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffCreateFileRequested
//go:noescape
func FuncOffCreateFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffCreateFileRequested
//go:noescape
func CallOffCreateFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffCreateFileRequested
//go:noescape
func TryOffCreateFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnCreateFileRequested
//go:noescape
func HasFuncHasOnCreateFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnCreateFileRequested
//go:noescape
func FuncHasOnCreateFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnCreateFileRequested
//go:noescape
func CallHasOnCreateFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnCreateFileRequested
//go:noescape
func TryHasOnCreateFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnDeleteEntryRequested
//go:noescape
func HasFuncOnDeleteEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnDeleteEntryRequested
//go:noescape
func FuncOnDeleteEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnDeleteEntryRequested
//go:noescape
func CallOnDeleteEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnDeleteEntryRequested
//go:noescape
func TryOnDeleteEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffDeleteEntryRequested
//go:noescape
func HasFuncOffDeleteEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffDeleteEntryRequested
//go:noescape
func FuncOffDeleteEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffDeleteEntryRequested
//go:noescape
func CallOffDeleteEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffDeleteEntryRequested
//go:noescape
func TryOffDeleteEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnDeleteEntryRequested
//go:noescape
func HasFuncHasOnDeleteEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnDeleteEntryRequested
//go:noescape
func FuncHasOnDeleteEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnDeleteEntryRequested
//go:noescape
func CallHasOnDeleteEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnDeleteEntryRequested
//go:noescape
func TryHasOnDeleteEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnExecuteActionRequested
//go:noescape
func HasFuncOnExecuteActionRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnExecuteActionRequested
//go:noescape
func FuncOnExecuteActionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnExecuteActionRequested
//go:noescape
func CallOnExecuteActionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnExecuteActionRequested
//go:noescape
func TryOnExecuteActionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffExecuteActionRequested
//go:noescape
func HasFuncOffExecuteActionRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffExecuteActionRequested
//go:noescape
func FuncOffExecuteActionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffExecuteActionRequested
//go:noescape
func CallOffExecuteActionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffExecuteActionRequested
//go:noescape
func TryOffExecuteActionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnExecuteActionRequested
//go:noescape
func HasFuncHasOnExecuteActionRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnExecuteActionRequested
//go:noescape
func FuncHasOnExecuteActionRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnExecuteActionRequested
//go:noescape
func CallHasOnExecuteActionRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnExecuteActionRequested
//go:noescape
func TryHasOnExecuteActionRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnGetActionsRequested
//go:noescape
func HasFuncOnGetActionsRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnGetActionsRequested
//go:noescape
func FuncOnGetActionsRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnGetActionsRequested
//go:noescape
func CallOnGetActionsRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnGetActionsRequested
//go:noescape
func TryOnGetActionsRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffGetActionsRequested
//go:noescape
func HasFuncOffGetActionsRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffGetActionsRequested
//go:noescape
func FuncOffGetActionsRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffGetActionsRequested
//go:noescape
func CallOffGetActionsRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffGetActionsRequested
//go:noescape
func TryOffGetActionsRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnGetActionsRequested
//go:noescape
func HasFuncHasOnGetActionsRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnGetActionsRequested
//go:noescape
func FuncHasOnGetActionsRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnGetActionsRequested
//go:noescape
func CallHasOnGetActionsRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnGetActionsRequested
//go:noescape
func TryHasOnGetActionsRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnGetMetadataRequested
//go:noescape
func HasFuncOnGetMetadataRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnGetMetadataRequested
//go:noescape
func FuncOnGetMetadataRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnGetMetadataRequested
//go:noescape
func CallOnGetMetadataRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnGetMetadataRequested
//go:noescape
func TryOnGetMetadataRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffGetMetadataRequested
//go:noescape
func HasFuncOffGetMetadataRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffGetMetadataRequested
//go:noescape
func FuncOffGetMetadataRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffGetMetadataRequested
//go:noescape
func CallOffGetMetadataRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffGetMetadataRequested
//go:noescape
func TryOffGetMetadataRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnGetMetadataRequested
//go:noescape
func HasFuncHasOnGetMetadataRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnGetMetadataRequested
//go:noescape
func FuncHasOnGetMetadataRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnGetMetadataRequested
//go:noescape
func CallHasOnGetMetadataRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnGetMetadataRequested
//go:noescape
func TryHasOnGetMetadataRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnMountRequested
//go:noescape
func HasFuncOnMountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnMountRequested
//go:noescape
func FuncOnMountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnMountRequested
//go:noescape
func CallOnMountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnMountRequested
//go:noescape
func TryOnMountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffMountRequested
//go:noescape
func HasFuncOffMountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffMountRequested
//go:noescape
func FuncOffMountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffMountRequested
//go:noescape
func CallOffMountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffMountRequested
//go:noescape
func TryOffMountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnMountRequested
//go:noescape
func HasFuncHasOnMountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnMountRequested
//go:noescape
func FuncHasOnMountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnMountRequested
//go:noescape
func CallHasOnMountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnMountRequested
//go:noescape
func TryHasOnMountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnMoveEntryRequested
//go:noescape
func HasFuncOnMoveEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnMoveEntryRequested
//go:noescape
func FuncOnMoveEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnMoveEntryRequested
//go:noescape
func CallOnMoveEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnMoveEntryRequested
//go:noescape
func TryOnMoveEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffMoveEntryRequested
//go:noescape
func HasFuncOffMoveEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffMoveEntryRequested
//go:noescape
func FuncOffMoveEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffMoveEntryRequested
//go:noescape
func CallOffMoveEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffMoveEntryRequested
//go:noescape
func TryOffMoveEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnMoveEntryRequested
//go:noescape
func HasFuncHasOnMoveEntryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnMoveEntryRequested
//go:noescape
func FuncHasOnMoveEntryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnMoveEntryRequested
//go:noescape
func CallHasOnMoveEntryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnMoveEntryRequested
//go:noescape
func TryHasOnMoveEntryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnOpenFileRequested
//go:noescape
func HasFuncOnOpenFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnOpenFileRequested
//go:noescape
func FuncOnOpenFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnOpenFileRequested
//go:noescape
func CallOnOpenFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnOpenFileRequested
//go:noescape
func TryOnOpenFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffOpenFileRequested
//go:noescape
func HasFuncOffOpenFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffOpenFileRequested
//go:noescape
func FuncOffOpenFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffOpenFileRequested
//go:noescape
func CallOffOpenFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffOpenFileRequested
//go:noescape
func TryOffOpenFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnOpenFileRequested
//go:noescape
func HasFuncHasOnOpenFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnOpenFileRequested
//go:noescape
func FuncHasOnOpenFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnOpenFileRequested
//go:noescape
func CallHasOnOpenFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnOpenFileRequested
//go:noescape
func TryHasOnOpenFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnReadDirectoryRequested
//go:noescape
func HasFuncOnReadDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnReadDirectoryRequested
//go:noescape
func FuncOnReadDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnReadDirectoryRequested
//go:noescape
func CallOnReadDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnReadDirectoryRequested
//go:noescape
func TryOnReadDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffReadDirectoryRequested
//go:noescape
func HasFuncOffReadDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffReadDirectoryRequested
//go:noescape
func FuncOffReadDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffReadDirectoryRequested
//go:noescape
func CallOffReadDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffReadDirectoryRequested
//go:noescape
func TryOffReadDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnReadDirectoryRequested
//go:noescape
func HasFuncHasOnReadDirectoryRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnReadDirectoryRequested
//go:noescape
func FuncHasOnReadDirectoryRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnReadDirectoryRequested
//go:noescape
func CallHasOnReadDirectoryRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnReadDirectoryRequested
//go:noescape
func TryHasOnReadDirectoryRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnReadFileRequested
//go:noescape
func HasFuncOnReadFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnReadFileRequested
//go:noescape
func FuncOnReadFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnReadFileRequested
//go:noescape
func CallOnReadFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnReadFileRequested
//go:noescape
func TryOnReadFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffReadFileRequested
//go:noescape
func HasFuncOffReadFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffReadFileRequested
//go:noescape
func FuncOffReadFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffReadFileRequested
//go:noescape
func CallOffReadFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffReadFileRequested
//go:noescape
func TryOffReadFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnReadFileRequested
//go:noescape
func HasFuncHasOnReadFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnReadFileRequested
//go:noescape
func FuncHasOnReadFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnReadFileRequested
//go:noescape
func CallHasOnReadFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnReadFileRequested
//go:noescape
func TryHasOnReadFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnRemoveWatcherRequested
//go:noescape
func HasFuncOnRemoveWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnRemoveWatcherRequested
//go:noescape
func FuncOnRemoveWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnRemoveWatcherRequested
//go:noescape
func CallOnRemoveWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnRemoveWatcherRequested
//go:noescape
func TryOnRemoveWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffRemoveWatcherRequested
//go:noescape
func HasFuncOffRemoveWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffRemoveWatcherRequested
//go:noescape
func FuncOffRemoveWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffRemoveWatcherRequested
//go:noescape
func CallOffRemoveWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffRemoveWatcherRequested
//go:noescape
func TryOffRemoveWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnRemoveWatcherRequested
//go:noescape
func HasFuncHasOnRemoveWatcherRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnRemoveWatcherRequested
//go:noescape
func FuncHasOnRemoveWatcherRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnRemoveWatcherRequested
//go:noescape
func CallHasOnRemoveWatcherRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnRemoveWatcherRequested
//go:noescape
func TryHasOnRemoveWatcherRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnTruncateRequested
//go:noescape
func HasFuncOnTruncateRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnTruncateRequested
//go:noescape
func FuncOnTruncateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnTruncateRequested
//go:noescape
func CallOnTruncateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnTruncateRequested
//go:noescape
func TryOnTruncateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffTruncateRequested
//go:noescape
func HasFuncOffTruncateRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffTruncateRequested
//go:noescape
func FuncOffTruncateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffTruncateRequested
//go:noescape
func CallOffTruncateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffTruncateRequested
//go:noescape
func TryOffTruncateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnTruncateRequested
//go:noescape
func HasFuncHasOnTruncateRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnTruncateRequested
//go:noescape
func FuncHasOnTruncateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnTruncateRequested
//go:noescape
func CallHasOnTruncateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnTruncateRequested
//go:noescape
func TryHasOnTruncateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnUnmountRequested
//go:noescape
func HasFuncOnUnmountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnUnmountRequested
//go:noescape
func FuncOnUnmountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnUnmountRequested
//go:noescape
func CallOnUnmountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnUnmountRequested
//go:noescape
func TryOnUnmountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffUnmountRequested
//go:noescape
func HasFuncOffUnmountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffUnmountRequested
//go:noescape
func FuncOffUnmountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffUnmountRequested
//go:noescape
func CallOffUnmountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffUnmountRequested
//go:noescape
func TryOffUnmountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnUnmountRequested
//go:noescape
func HasFuncHasOnUnmountRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnUnmountRequested
//go:noescape
func FuncHasOnUnmountRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnUnmountRequested
//go:noescape
func CallHasOnUnmountRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnUnmountRequested
//go:noescape
func TryHasOnUnmountRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OnWriteFileRequested
//go:noescape
func HasFuncOnWriteFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OnWriteFileRequested
//go:noescape
func FuncOnWriteFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OnWriteFileRequested
//go:noescape
func CallOnWriteFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OnWriteFileRequested
//go:noescape
func TryOnWriteFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_OffWriteFileRequested
//go:noescape
func HasFuncOffWriteFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_OffWriteFileRequested
//go:noescape
func FuncOffWriteFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_OffWriteFileRequested
//go:noescape
func CallOffWriteFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_OffWriteFileRequested
//go:noescape
func TryOffWriteFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_HasOnWriteFileRequested
//go:noescape
func HasFuncHasOnWriteFileRequested() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_HasOnWriteFileRequested
//go:noescape
func FuncHasOnWriteFileRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_HasOnWriteFileRequested
//go:noescape
func CallHasOnWriteFileRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider try_HasOnWriteFileRequested
//go:noescape
func TryHasOnWriteFileRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemprovider has_Unmount
//go:noescape
func HasFuncUnmount() js.Ref

//go:wasmimport plat/js/webext/filesystemprovider func_Unmount
//go:noescape
func FuncUnmount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider call_Unmount
//go:noescape
func CallUnmount(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemprovider try_Unmount
//go:noescape
func TryUnmount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
