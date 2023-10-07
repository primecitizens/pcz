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

//go:wasmimport plat/js/webext/filemanagerprivateinternal store_EntryDescription
//go:noescape
func EntryDescriptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal load_EntryDescription
//go:noescape
func EntryDescriptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal store_IOTaskParams
//go:noescape
func IOTaskParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal load_IOTaskParams
//go:noescape
func IOTaskParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal store_ParsedTrashInfoFile
//go:noescape
func ParsedTrashInfoFileJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal load_ParsedTrashInfoFile
//go:noescape
func ParsedTrashInfoFileJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal store_SearchFilesParams
//go:noescape
func SearchFilesParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal load_SearchFilesParams
//go:noescape
func SearchFilesParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_AddFileWatch
//go:noescape
func HasFuncAddFileWatch() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_AddFileWatch
//go:noescape
func FuncAddFileWatch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_AddFileWatch
//go:noescape
func CallAddFileWatch(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_AddFileWatch
//go:noescape
func TryAddFileWatch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ComputeChecksum
//go:noescape
func HasFuncComputeChecksum() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ComputeChecksum
//go:noescape
func FuncComputeChecksum(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ComputeChecksum
//go:noescape
func CallComputeChecksum(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ComputeChecksum
//go:noescape
func TryComputeChecksum(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ExecuteCustomAction
//go:noescape
func HasFuncExecuteCustomAction() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ExecuteCustomAction
//go:noescape
func FuncExecuteCustomAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ExecuteCustomAction
//go:noescape
func CallExecuteCustomAction(
	retPtr unsafe.Pointer,
	urls js.Ref,
	actionId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ExecuteCustomAction
//go:noescape
func TryExecuteCustomAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref,
	actionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ExecuteTask
//go:noescape
func HasFuncExecuteTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ExecuteTask
//go:noescape
func FuncExecuteTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ExecuteTask
//go:noescape
func CallExecuteTask(
	retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ExecuteTask
//go:noescape
func TryExecuteTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetContentMetadata
//go:noescape
func HasFuncGetContentMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetContentMetadata
//go:noescape
func FuncGetContentMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetContentMetadata
//go:noescape
func CallGetContentMetadata(
	retPtr unsafe.Pointer,
	blobUUID js.Ref,
	mimeType js.Ref,
	includeImages js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetContentMetadata
//go:noescape
func TryGetContentMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	blobUUID js.Ref,
	mimeType js.Ref,
	includeImages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetContentMimeType
//go:noescape
func HasFuncGetContentMimeType() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetContentMimeType
//go:noescape
func FuncGetContentMimeType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetContentMimeType
//go:noescape
func CallGetContentMimeType(
	retPtr unsafe.Pointer,
	blobUUID js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetContentMimeType
//go:noescape
func TryGetContentMimeType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	blobUUID js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetCrostiniSharedPaths
//go:noescape
func HasFuncGetCrostiniSharedPaths() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetCrostiniSharedPaths
//go:noescape
func FuncGetCrostiniSharedPaths(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetCrostiniSharedPaths
//go:noescape
func CallGetCrostiniSharedPaths(
	retPtr unsafe.Pointer,
	observeFirstForSession js.Ref,
	vmName js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetCrostiniSharedPaths
//go:noescape
func TryGetCrostiniSharedPaths(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	observeFirstForSession js.Ref,
	vmName js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetCustomActions
//go:noescape
func HasFuncGetCustomActions() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetCustomActions
//go:noescape
func FuncGetCustomActions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetCustomActions
//go:noescape
func CallGetCustomActions(
	retPtr unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetCustomActions
//go:noescape
func TryGetCustomActions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetDirectorySize
//go:noescape
func HasFuncGetDirectorySize() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetDirectorySize
//go:noescape
func FuncGetDirectorySize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetDirectorySize
//go:noescape
func CallGetDirectorySize(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetDirectorySize
//go:noescape
func TryGetDirectorySize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetDisallowedTransfers
//go:noescape
func HasFuncGetDisallowedTransfers() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetDisallowedTransfers
//go:noescape
func FuncGetDisallowedTransfers(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetDisallowedTransfers
//go:noescape
func CallGetDisallowedTransfers(
	retPtr unsafe.Pointer,
	entries js.Ref,
	destinationEntry js.Ref,
	isMove js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetDisallowedTransfers
//go:noescape
func TryGetDisallowedTransfers(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	destinationEntry js.Ref,
	isMove js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetDlpMetadata
//go:noescape
func HasFuncGetDlpMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetDlpMetadata
//go:noescape
func FuncGetDlpMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetDlpMetadata
//go:noescape
func CallGetDlpMetadata(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetDlpMetadata
//go:noescape
func TryGetDlpMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetDriveQuotaMetadata
//go:noescape
func HasFuncGetDriveQuotaMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetDriveQuotaMetadata
//go:noescape
func FuncGetDriveQuotaMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetDriveQuotaMetadata
//go:noescape
func CallGetDriveQuotaMetadata(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetDriveQuotaMetadata
//go:noescape
func TryGetDriveQuotaMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetEntryProperties
//go:noescape
func HasFuncGetEntryProperties() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetEntryProperties
//go:noescape
func FuncGetEntryProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetEntryProperties
//go:noescape
func CallGetEntryProperties(
	retPtr unsafe.Pointer,
	urls js.Ref,
	names js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetEntryProperties
//go:noescape
func TryGetEntryProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref,
	names js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetFileTasks
//go:noescape
func HasFuncGetFileTasks() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetFileTasks
//go:noescape
func FuncGetFileTasks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetFileTasks
//go:noescape
func CallGetFileTasks(
	retPtr unsafe.Pointer,
	urls js.Ref,
	dlpSourceUrls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetFileTasks
//go:noescape
func TryGetFileTasks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref,
	dlpSourceUrls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetLinuxPackageInfo
//go:noescape
func HasFuncGetLinuxPackageInfo() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetLinuxPackageInfo
//go:noescape
func FuncGetLinuxPackageInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetLinuxPackageInfo
//go:noescape
func CallGetLinuxPackageInfo(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetLinuxPackageInfo
//go:noescape
func TryGetLinuxPackageInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetMimeType
//go:noescape
func HasFuncGetMimeType() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetMimeType
//go:noescape
func FuncGetMimeType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetMimeType
//go:noescape
func CallGetMimeType(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetMimeType
//go:noescape
func TryGetMimeType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetRecentFiles
//go:noescape
func HasFuncGetRecentFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetRecentFiles
//go:noescape
func FuncGetRecentFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetRecentFiles
//go:noescape
func CallGetRecentFiles(
	retPtr unsafe.Pointer,
	restriction uint32,
	file_category uint32,
	invalidate_cache js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetRecentFiles
//go:noescape
func TryGetRecentFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	restriction uint32,
	file_category uint32,
	invalidate_cache js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_GetVolumeRoot
//go:noescape
func HasFuncGetVolumeRoot() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_GetVolumeRoot
//go:noescape
func FuncGetVolumeRoot(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_GetVolumeRoot
//go:noescape
func CallGetVolumeRoot(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_GetVolumeRoot
//go:noescape
func TryGetVolumeRoot(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ImportCrostiniImage
//go:noescape
func HasFuncImportCrostiniImage() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ImportCrostiniImage
//go:noescape
func FuncImportCrostiniImage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ImportCrostiniImage
//go:noescape
func CallImportCrostiniImage(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ImportCrostiniImage
//go:noescape
func TryImportCrostiniImage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_InstallLinuxPackage
//go:noescape
func HasFuncInstallLinuxPackage() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_InstallLinuxPackage
//go:noescape
func FuncInstallLinuxPackage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_InstallLinuxPackage
//go:noescape
func CallInstallLinuxPackage(
	retPtr unsafe.Pointer,
	url js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_InstallLinuxPackage
//go:noescape
func TryInstallLinuxPackage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_InvokeSharesheet
//go:noescape
func HasFuncInvokeSharesheet() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_InvokeSharesheet
//go:noescape
func FuncInvokeSharesheet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_InvokeSharesheet
//go:noescape
func CallInvokeSharesheet(
	retPtr unsafe.Pointer,
	urls js.Ref,
	launchSource uint32,
	dlpSourceUrls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_InvokeSharesheet
//go:noescape
func TryInvokeSharesheet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref,
	launchSource uint32,
	dlpSourceUrls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ParseTrashInfoFiles
//go:noescape
func HasFuncParseTrashInfoFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ParseTrashInfoFiles
//go:noescape
func FuncParseTrashInfoFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ParseTrashInfoFiles
//go:noescape
func CallParseTrashInfoFiles(
	retPtr unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ParseTrashInfoFiles
//go:noescape
func TryParseTrashInfoFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_PinDriveFile
//go:noescape
func HasFuncPinDriveFile() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_PinDriveFile
//go:noescape
func FuncPinDriveFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_PinDriveFile
//go:noescape
func CallPinDriveFile(
	retPtr unsafe.Pointer,
	url js.Ref,
	pin js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_PinDriveFile
//go:noescape
func TryPinDriveFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	pin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_RemoveFileWatch
//go:noescape
func HasFuncRemoveFileWatch() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_RemoveFileWatch
//go:noescape
func FuncRemoveFileWatch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_RemoveFileWatch
//go:noescape
func CallRemoveFileWatch(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_RemoveFileWatch
//go:noescape
func TryRemoveFileWatch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ResolveIsolatedEntries
//go:noescape
func HasFuncResolveIsolatedEntries() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ResolveIsolatedEntries
//go:noescape
func FuncResolveIsolatedEntries(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ResolveIsolatedEntries
//go:noescape
func CallResolveIsolatedEntries(
	retPtr unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ResolveIsolatedEntries
//go:noescape
func TryResolveIsolatedEntries(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_SearchFiles
//go:noescape
func HasFuncSearchFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_SearchFiles
//go:noescape
func FuncSearchFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_SearchFiles
//go:noescape
func CallSearchFiles(
	retPtr unsafe.Pointer,
	searchParams unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_SearchFiles
//go:noescape
func TrySearchFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_SetDefaultTask
//go:noescape
func HasFuncSetDefaultTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_SetDefaultTask
//go:noescape
func FuncSetDefaultTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_SetDefaultTask
//go:noescape
func CallSetDefaultTask(
	retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	urls js.Ref,
	mimeTypes js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_SetDefaultTask
//go:noescape
func TrySetDefaultTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	urls js.Ref,
	mimeTypes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_SharePathsWithCrostini
//go:noescape
func HasFuncSharePathsWithCrostini() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_SharePathsWithCrostini
//go:noescape
func FuncSharePathsWithCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_SharePathsWithCrostini
//go:noescape
func CallSharePathsWithCrostini(
	retPtr unsafe.Pointer,
	vmName js.Ref,
	urls js.Ref,
	persist js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_SharePathsWithCrostini
//go:noescape
func TrySharePathsWithCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vmName js.Ref,
	urls js.Ref,
	persist js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_SharesheetHasTargets
//go:noescape
func HasFuncSharesheetHasTargets() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_SharesheetHasTargets
//go:noescape
func FuncSharesheetHasTargets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_SharesheetHasTargets
//go:noescape
func CallSharesheetHasTargets(
	retPtr unsafe.Pointer,
	urls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_SharesheetHasTargets
//go:noescape
func TrySharesheetHasTargets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_StartIOTask
//go:noescape
func HasFuncStartIOTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_StartIOTask
//go:noescape
func FuncStartIOTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_StartIOTask
//go:noescape
func CallStartIOTask(
	retPtr unsafe.Pointer,
	typ uint32,
	urls js.Ref,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_StartIOTask
//go:noescape
func TryStartIOTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	urls js.Ref,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ToggleAddedToHoldingSpace
//go:noescape
func HasFuncToggleAddedToHoldingSpace() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ToggleAddedToHoldingSpace
//go:noescape
func FuncToggleAddedToHoldingSpace(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ToggleAddedToHoldingSpace
//go:noescape
func CallToggleAddedToHoldingSpace(
	retPtr unsafe.Pointer,
	urls js.Ref,
	add js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ToggleAddedToHoldingSpace
//go:noescape
func TryToggleAddedToHoldingSpace(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	urls js.Ref,
	add js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_UnsharePathWithCrostini
//go:noescape
func HasFuncUnsharePathWithCrostini() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_UnsharePathWithCrostini
//go:noescape
func FuncUnsharePathWithCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_UnsharePathWithCrostini
//go:noescape
func CallUnsharePathWithCrostini(
	retPtr unsafe.Pointer,
	vmName js.Ref,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_UnsharePathWithCrostini
//go:noescape
func TryUnsharePathWithCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vmName js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal has_ValidatePathNameLength
//go:noescape
func HasFuncValidatePathNameLength() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivateinternal func_ValidatePathNameLength
//go:noescape
func FuncValidatePathNameLength(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivateinternal call_ValidatePathNameLength
//go:noescape
func CallValidatePathNameLength(
	retPtr unsafe.Pointer,
	parentUrl js.Ref,
	name js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivateinternal try_ValidatePathNameLength
//go:noescape
func TryValidatePathNameLength(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentUrl js.Ref,
	name js.Ref) (ok js.Ref)
