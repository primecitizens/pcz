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

//go:wasmimport plat/js/webext/filemanagerprivate store_IconSet
//go:noescape
func IconSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_IconSet
//go:noescape
func IconSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_AndroidApp
//go:noescape
func AndroidAppJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_AndroidApp
//go:noescape
func AndroidAppJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_AttachedImages
//go:noescape
func AttachedImagesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_AttachedImages
//go:noescape
func AttachedImagesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_BulkPinStage
//go:noescape
func ConstOfBulkPinStage(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_BulkPinProgress
//go:noescape
func BulkPinProgressJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_BulkPinProgress
//go:noescape
func BulkPinProgressJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_ChangeType
//go:noescape
func ConstOfChangeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_ConflictPauseParams
//go:noescape
func ConflictPauseParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ConflictPauseParams
//go:noescape
func ConflictPauseParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_ConflictResumeParams
//go:noescape
func ConflictResumeParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ConflictResumeParams
//go:noescape
func ConflictResumeParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_CrostiniEventType
//go:noescape
func ConstOfCrostiniEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_CrostiniEvent
//go:noescape
func CrostiniEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_CrostiniEvent
//go:noescape
func CrostiniEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DeviceConnectionState
//go:noescape
func ConstOfDeviceConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_DeviceEventType
//go:noescape
func ConstOfDeviceEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DeviceEvent
//go:noescape
func DeviceEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DeviceEvent
//go:noescape
func DeviceEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DeviceType
//go:noescape
func ConstOfDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_VolumeType
//go:noescape
func ConstOfVolumeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DialogCallerInformation
//go:noescape
func DialogCallerInformationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DialogCallerInformation
//go:noescape
func DialogCallerInformationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DlpLevel
//go:noescape
func ConstOfDlpLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DlpMetadata
//go:noescape
func DlpMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DlpMetadata
//go:noescape
func DlpMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_DlpRestrictionDetails
//go:noescape
func DlpRestrictionDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DlpRestrictionDetails
//go:noescape
func DlpRestrictionDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveConfirmDialogType
//go:noescape
func ConstOfDriveConfirmDialogType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DriveConfirmDialogEvent
//go:noescape
func DriveConfirmDialogEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DriveConfirmDialogEvent
//go:noescape
func DriveConfirmDialogEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveConnectionStateType
//go:noescape
func ConstOfDriveConnectionStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveOfflineReason
//go:noescape
func ConstOfDriveOfflineReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DriveConnectionState
//go:noescape
func DriveConnectionStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DriveConnectionState
//go:noescape
func DriveConnectionStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveDialogResult
//go:noescape
func ConstOfDriveDialogResult(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DriveMetadataSearchResult
//go:noescape
func DriveMetadataSearchResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DriveMetadataSearchResult
//go:noescape
func DriveMetadataSearchResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_UserType
//go:noescape
func ConstOfUserType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DriveQuotaMetadata
//go:noescape
func DriveQuotaMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DriveQuotaMetadata
//go:noescape
func DriveQuotaMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveShareType
//go:noescape
func ConstOfDriveShareType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_DriveSyncErrorType
//go:noescape
func ConstOfDriveSyncErrorType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_DriveSyncErrorEvent
//go:noescape
func DriveSyncErrorEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_DriveSyncErrorEvent
//go:noescape
func DriveSyncErrorEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_RecentDateBucket
//go:noescape
func ConstOfRecentDateBucket(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_SyncStatus
//go:noescape
func ConstOfSyncStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_EntryProperties
//go:noescape
func EntryPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_EntryProperties
//go:noescape
func EntryPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_EntryPropertyName
//go:noescape
func ConstOfEntryPropertyName(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_TaskResult
//go:noescape
func ConstOfTaskResult(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_FileCategory
//go:noescape
func ConstOfFileCategory(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_FileChange
//go:noescape
func FileChangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileChange
//go:noescape
func FileChangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_FileSystemProviderAction
//go:noescape
func FileSystemProviderActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileSystemProviderAction
//go:noescape
func FileSystemProviderActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_FileTaskDescriptor
//go:noescape
func FileTaskDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileTaskDescriptor
//go:noescape
func FileTaskDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_FileTask
//go:noescape
func FileTaskJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileTask
//go:noescape
func FileTaskJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_TransferState
//go:noescape
func ConstOfTransferState(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_FileTransferStatus
//go:noescape
func FileTransferStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileTransferStatus
//go:noescape
func FileTransferStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_FileWatchEventType
//go:noescape
func ConstOfFileWatchEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_FileWatchEvent
//go:noescape
func FileWatchEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_FileWatchEvent
//go:noescape
func FileWatchEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_FormatFileSystemType
//go:noescape
func ConstOfFormatFileSystemType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_StreamInfo
//go:noescape
func StreamInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_StreamInfo
//go:noescape
func StreamInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_MediaMetadata
//go:noescape
func MediaMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_MediaMetadata
//go:noescape
func MediaMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_PolicyDefaultHandlerStatus
//go:noescape
func ConstOfPolicyDefaultHandlerStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_ResultingTasks
//go:noescape
func ResultingTasksJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ResultingTasks
//go:noescape
func ResultingTasksJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_LinuxPackageInfo
//go:noescape
func LinuxPackageInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_LinuxPackageInfo
//go:noescape
func LinuxPackageInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_Preferences
//go:noescape
func PreferencesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_Preferences
//go:noescape
func PreferencesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_ProfileInfo
//go:noescape
func ProfileInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ProfileInfo
//go:noescape
func ProfileInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_ProviderSource
//go:noescape
func ConstOfProviderSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_Provider
//go:noescape
func ProviderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_Provider
//go:noescape
func ProviderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_MountPointSizeStats
//go:noescape
func MountPointSizeStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_MountPointSizeStats
//go:noescape
func MountPointSizeStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_Source
//go:noescape
func ConstOfSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_MountError
//go:noescape
func ConstOfMountError(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_MountContext
//go:noescape
func ConstOfMountContext(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_VmType
//go:noescape
func ConstOfVmType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_VolumeMetadata
//go:noescape
func VolumeMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_VolumeMetadata
//go:noescape
func VolumeMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_GetVolumeRootOptions
//go:noescape
func GetVolumeRootOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_GetVolumeRootOptions
//go:noescape
func GetVolumeRootOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_HoldingSpaceState
//go:noescape
func HoldingSpaceStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_HoldingSpaceState
//go:noescape
func HoldingSpaceStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_IOTaskParams
//go:noescape
func IOTaskParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_IOTaskParams
//go:noescape
func IOTaskParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_IOTaskState
//go:noescape
func ConstOfIOTaskState(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_IOTaskType
//go:noescape
func ConstOfIOTaskType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_InspectionType
//go:noescape
func ConstOfInspectionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_InstallLinuxPackageResponse
//go:noescape
func ConstOfInstallLinuxPackageResponse(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_MountableGuest
//go:noescape
func MountableGuestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_MountableGuest
//go:noescape
func MountableGuestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_MountCompletedEventType
//go:noescape
func ConstOfMountCompletedEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_MountCompletedEvent
//go:noescape
func MountCompletedEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_MountCompletedEvent
//go:noescape
func MountCompletedEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_OpenWindowParams
//go:noescape
func OpenWindowParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_OpenWindowParams
//go:noescape
func OpenWindowParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_ParsedTrashInfoFile
//go:noescape
func ParsedTrashInfoFileJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ParsedTrashInfoFile
//go:noescape
func ParsedTrashInfoFileJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_PolicyErrorType
//go:noescape
func ConstOfPolicyErrorType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_PolicyPauseParams
//go:noescape
func PolicyPauseParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_PolicyPauseParams
//go:noescape
func PolicyPauseParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_PauseParams
//go:noescape
func PauseParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_PauseParams
//go:noescape
func PauseParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_PolicyDialogType
//go:noescape
func ConstOfPolicyDialogType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_PolicyError
//go:noescape
func PolicyErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_PolicyError
//go:noescape
func PolicyErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_PolicyResumeParams
//go:noescape
func PolicyResumeParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_PolicyResumeParams
//go:noescape
func PolicyResumeParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_PreferencesChange
//go:noescape
func PreferencesChangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_PreferencesChange
//go:noescape
func PreferencesChangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_ProgressStatus
//go:noescape
func ProgressStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ProgressStatus
//go:noescape
func ProgressStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_ResumeParams
//go:noescape
func ResumeParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_ResumeParams
//go:noescape
func ResumeParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_SearchType
//go:noescape
func ConstOfSearchType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_SearchMetadataParams
//go:noescape
func SearchMetadataParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_SearchMetadataParams
//go:noescape
func SearchMetadataParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate store_SearchParams
//go:noescape
func SearchParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_SearchParams
//go:noescape
func SearchParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_SharesheetLaunchSource
//go:noescape
func ConstOfSharesheetLaunchSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate constof_SourceRestriction
//go:noescape
func ConstOfSourceRestriction(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate store_SyncState
//go:noescape
func SyncStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate load_SyncState
//go:noescape
func SyncStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate constof_ZoomOperationType
//go:noescape
func ConstOfZoomOperationType(str js.Ref) uint32

//go:wasmimport plat/js/webext/filemanagerprivate has_AddFileWatch
//go:noescape
func HasFuncAddFileWatch() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_AddFileWatch
//go:noescape
func FuncAddFileWatch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_AddFileWatch
//go:noescape
func CallAddFileWatch(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_AddFileWatch
//go:noescape
func TryAddFileWatch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_AddMount
//go:noescape
func HasFuncAddMount() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_AddMount
//go:noescape
func FuncAddMount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_AddMount
//go:noescape
func CallAddMount(
	retPtr unsafe.Pointer,
	fileUrl js.Ref,
	password js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_AddMount
//go:noescape
func TryAddMount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileUrl js.Ref,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_AddProvidedFileSystem
//go:noescape
func HasFuncAddProvidedFileSystem() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_AddProvidedFileSystem
//go:noescape
func FuncAddProvidedFileSystem(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_AddProvidedFileSystem
//go:noescape
func CallAddProvidedFileSystem(
	retPtr unsafe.Pointer,
	providerId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_AddProvidedFileSystem
//go:noescape
func TryAddProvidedFileSystem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	providerId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_CalculateBulkPinRequiredSpace
//go:noescape
func HasFuncCalculateBulkPinRequiredSpace() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_CalculateBulkPinRequiredSpace
//go:noescape
func FuncCalculateBulkPinRequiredSpace(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_CalculateBulkPinRequiredSpace
//go:noescape
func CallCalculateBulkPinRequiredSpace(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_CalculateBulkPinRequiredSpace
//go:noescape
func TryCalculateBulkPinRequiredSpace(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_CancelDialog
//go:noescape
func HasFuncCancelDialog() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_CancelDialog
//go:noescape
func FuncCancelDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_CancelDialog
//go:noescape
func CallCancelDialog(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_CancelDialog
//go:noescape
func TryCancelDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_CancelIOTask
//go:noescape
func HasFuncCancelIOTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_CancelIOTask
//go:noescape
func FuncCancelIOTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_CancelIOTask
//go:noescape
func CallCancelIOTask(
	retPtr unsafe.Pointer,
	taskId int32)

//go:wasmimport plat/js/webext/filemanagerprivate try_CancelIOTask
//go:noescape
func TryCancelIOTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	taskId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_CancelMounting
//go:noescape
func HasFuncCancelMounting() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_CancelMounting
//go:noescape
func FuncCancelMounting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_CancelMounting
//go:noescape
func CallCancelMounting(
	retPtr unsafe.Pointer,
	fileUrl js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_CancelMounting
//go:noescape
func TryCancelMounting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ComputeChecksum
//go:noescape
func HasFuncComputeChecksum() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ComputeChecksum
//go:noescape
func FuncComputeChecksum(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ComputeChecksum
//go:noescape
func CallComputeChecksum(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ComputeChecksum
//go:noescape
func TryComputeChecksum(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ConfigureVolume
//go:noescape
func HasFuncConfigureVolume() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ConfigureVolume
//go:noescape
func FuncConfigureVolume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ConfigureVolume
//go:noescape
func CallConfigureVolume(
	retPtr unsafe.Pointer,
	volumeId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ConfigureVolume
//go:noescape
func TryConfigureVolume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_DismissIOTask
//go:noescape
func HasFuncDismissIOTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_DismissIOTask
//go:noescape
func FuncDismissIOTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_DismissIOTask
//go:noescape
func CallDismissIOTask(
	retPtr unsafe.Pointer,
	taskId int32)

//go:wasmimport plat/js/webext/filemanagerprivate try_DismissIOTask
//go:noescape
func TryDismissIOTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	taskId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_EnableExternalFileScheme
//go:noescape
func HasFuncEnableExternalFileScheme() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_EnableExternalFileScheme
//go:noescape
func FuncEnableExternalFileScheme(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_EnableExternalFileScheme
//go:noescape
func CallEnableExternalFileScheme(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_EnableExternalFileScheme
//go:noescape
func TryEnableExternalFileScheme(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ExecuteCustomAction
//go:noescape
func HasFuncExecuteCustomAction() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ExecuteCustomAction
//go:noescape
func FuncExecuteCustomAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ExecuteCustomAction
//go:noescape
func CallExecuteCustomAction(
	retPtr unsafe.Pointer,
	entries js.Ref,
	actionId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ExecuteCustomAction
//go:noescape
func TryExecuteCustomAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	actionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ExecuteTask
//go:noescape
func HasFuncExecuteTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ExecuteTask
//go:noescape
func FuncExecuteTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ExecuteTask
//go:noescape
func CallExecuteTask(
	retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ExecuteTask
//go:noescape
func TryExecuteTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_FormatVolume
//go:noescape
func HasFuncFormatVolume() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_FormatVolume
//go:noescape
func FuncFormatVolume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_FormatVolume
//go:noescape
func CallFormatVolume(
	retPtr unsafe.Pointer,
	volumeId js.Ref,
	filesystem uint32,
	volumeLabel js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_FormatVolume
//go:noescape
func TryFormatVolume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref,
	filesystem uint32,
	volumeLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetAndroidPickerApps
//go:noescape
func HasFuncGetAndroidPickerApps() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetAndroidPickerApps
//go:noescape
func FuncGetAndroidPickerApps(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetAndroidPickerApps
//go:noescape
func CallGetAndroidPickerApps(
	retPtr unsafe.Pointer,
	extensions js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetAndroidPickerApps
//go:noescape
func TryGetAndroidPickerApps(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensions js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetBulkPinProgress
//go:noescape
func HasFuncGetBulkPinProgress() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetBulkPinProgress
//go:noescape
func FuncGetBulkPinProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetBulkPinProgress
//go:noescape
func CallGetBulkPinProgress(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetBulkPinProgress
//go:noescape
func TryGetBulkPinProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetContentMetadata
//go:noescape
func HasFuncGetContentMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetContentMetadata
//go:noescape
func FuncGetContentMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetContentMetadata
//go:noescape
func CallGetContentMetadata(
	retPtr unsafe.Pointer,
	fileEntry js.Ref,
	mimeType js.Ref,
	includeImages js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetContentMetadata
//go:noescape
func TryGetContentMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileEntry js.Ref,
	mimeType js.Ref,
	includeImages js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetContentMimeType
//go:noescape
func HasFuncGetContentMimeType() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetContentMimeType
//go:noescape
func FuncGetContentMimeType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetContentMimeType
//go:noescape
func CallGetContentMimeType(
	retPtr unsafe.Pointer,
	fileEntry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetContentMimeType
//go:noescape
func TryGetContentMimeType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileEntry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetCrostiniSharedPaths
//go:noescape
func HasFuncGetCrostiniSharedPaths() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetCrostiniSharedPaths
//go:noescape
func FuncGetCrostiniSharedPaths(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetCrostiniSharedPaths
//go:noescape
func CallGetCrostiniSharedPaths(
	retPtr unsafe.Pointer,
	observeFirstForSession js.Ref,
	vmName js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetCrostiniSharedPaths
//go:noescape
func TryGetCrostiniSharedPaths(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	observeFirstForSession js.Ref,
	vmName js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetCustomActions
//go:noescape
func HasFuncGetCustomActions() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetCustomActions
//go:noescape
func FuncGetCustomActions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetCustomActions
//go:noescape
func CallGetCustomActions(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetCustomActions
//go:noescape
func TryGetCustomActions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDeviceConnectionState
//go:noescape
func HasFuncGetDeviceConnectionState() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDeviceConnectionState
//go:noescape
func FuncGetDeviceConnectionState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDeviceConnectionState
//go:noescape
func CallGetDeviceConnectionState(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDeviceConnectionState
//go:noescape
func TryGetDeviceConnectionState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDialogCaller
//go:noescape
func HasFuncGetDialogCaller() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDialogCaller
//go:noescape
func FuncGetDialogCaller(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDialogCaller
//go:noescape
func CallGetDialogCaller(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDialogCaller
//go:noescape
func TryGetDialogCaller(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDirectorySize
//go:noescape
func HasFuncGetDirectorySize() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDirectorySize
//go:noescape
func FuncGetDirectorySize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDirectorySize
//go:noescape
func CallGetDirectorySize(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDirectorySize
//go:noescape
func TryGetDirectorySize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDisallowedTransfers
//go:noescape
func HasFuncGetDisallowedTransfers() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDisallowedTransfers
//go:noescape
func FuncGetDisallowedTransfers(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDisallowedTransfers
//go:noescape
func CallGetDisallowedTransfers(
	retPtr unsafe.Pointer,
	entries js.Ref,
	destinationEntry js.Ref,
	isMove js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDisallowedTransfers
//go:noescape
func TryGetDisallowedTransfers(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	destinationEntry js.Ref,
	isMove js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDlpBlockedComponents
//go:noescape
func HasFuncGetDlpBlockedComponents() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDlpBlockedComponents
//go:noescape
func FuncGetDlpBlockedComponents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDlpBlockedComponents
//go:noescape
func CallGetDlpBlockedComponents(
	retPtr unsafe.Pointer,
	sourceUrl js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDlpBlockedComponents
//go:noescape
func TryGetDlpBlockedComponents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sourceUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDlpMetadata
//go:noescape
func HasFuncGetDlpMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDlpMetadata
//go:noescape
func FuncGetDlpMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDlpMetadata
//go:noescape
func CallGetDlpMetadata(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDlpMetadata
//go:noescape
func TryGetDlpMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDlpRestrictionDetails
//go:noescape
func HasFuncGetDlpRestrictionDetails() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDlpRestrictionDetails
//go:noescape
func FuncGetDlpRestrictionDetails(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDlpRestrictionDetails
//go:noescape
func CallGetDlpRestrictionDetails(
	retPtr unsafe.Pointer,
	sourceUrl js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDlpRestrictionDetails
//go:noescape
func TryGetDlpRestrictionDetails(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sourceUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDriveConnectionState
//go:noescape
func HasFuncGetDriveConnectionState() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDriveConnectionState
//go:noescape
func FuncGetDriveConnectionState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDriveConnectionState
//go:noescape
func CallGetDriveConnectionState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDriveConnectionState
//go:noescape
func TryGetDriveConnectionState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetDriveQuotaMetadata
//go:noescape
func HasFuncGetDriveQuotaMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetDriveQuotaMetadata
//go:noescape
func FuncGetDriveQuotaMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetDriveQuotaMetadata
//go:noescape
func CallGetDriveQuotaMetadata(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetDriveQuotaMetadata
//go:noescape
func TryGetDriveQuotaMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetEntryProperties
//go:noescape
func HasFuncGetEntryProperties() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetEntryProperties
//go:noescape
func FuncGetEntryProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetEntryProperties
//go:noescape
func CallGetEntryProperties(
	retPtr unsafe.Pointer,
	entries js.Ref,
	names js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetEntryProperties
//go:noescape
func TryGetEntryProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	names js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetFileTasks
//go:noescape
func HasFuncGetFileTasks() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetFileTasks
//go:noescape
func FuncGetFileTasks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetFileTasks
//go:noescape
func CallGetFileTasks(
	retPtr unsafe.Pointer,
	entries js.Ref,
	dlpSourceUrls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetFileTasks
//go:noescape
func TryGetFileTasks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	dlpSourceUrls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetHoldingSpaceState
//go:noescape
func HasFuncGetHoldingSpaceState() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetHoldingSpaceState
//go:noescape
func FuncGetHoldingSpaceState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetHoldingSpaceState
//go:noescape
func CallGetHoldingSpaceState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetHoldingSpaceState
//go:noescape
func TryGetHoldingSpaceState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetLinuxPackageInfo
//go:noescape
func HasFuncGetLinuxPackageInfo() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetLinuxPackageInfo
//go:noescape
func FuncGetLinuxPackageInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetLinuxPackageInfo
//go:noescape
func CallGetLinuxPackageInfo(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetLinuxPackageInfo
//go:noescape
func TryGetLinuxPackageInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetMimeType
//go:noescape
func HasFuncGetMimeType() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetMimeType
//go:noescape
func FuncGetMimeType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetMimeType
//go:noescape
func CallGetMimeType(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetMimeType
//go:noescape
func TryGetMimeType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetPreferences
//go:noescape
func HasFuncGetPreferences() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetPreferences
//go:noescape
func FuncGetPreferences(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetPreferences
//go:noescape
func CallGetPreferences(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetPreferences
//go:noescape
func TryGetPreferences(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetProfiles
//go:noescape
func HasFuncGetProfiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetProfiles
//go:noescape
func FuncGetProfiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetProfiles
//go:noescape
func CallGetProfiles(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetProfiles
//go:noescape
func TryGetProfiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetProviders
//go:noescape
func HasFuncGetProviders() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetProviders
//go:noescape
func FuncGetProviders(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetProviders
//go:noescape
func CallGetProviders(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetProviders
//go:noescape
func TryGetProviders(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetRecentFiles
//go:noescape
func HasFuncGetRecentFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetRecentFiles
//go:noescape
func FuncGetRecentFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetRecentFiles
//go:noescape
func CallGetRecentFiles(
	retPtr unsafe.Pointer,
	restriction uint32,
	fileCategory uint32,
	invalidateCache js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetRecentFiles
//go:noescape
func TryGetRecentFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	restriction uint32,
	fileCategory uint32,
	invalidateCache js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetSizeStats
//go:noescape
func HasFuncGetSizeStats() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetSizeStats
//go:noescape
func FuncGetSizeStats(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetSizeStats
//go:noescape
func CallGetSizeStats(
	retPtr unsafe.Pointer,
	volumeId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetSizeStats
//go:noescape
func TryGetSizeStats(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetStrings
//go:noescape
func HasFuncGetStrings() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetStrings
//go:noescape
func FuncGetStrings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetStrings
//go:noescape
func CallGetStrings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetStrings
//go:noescape
func TryGetStrings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetVolumeMetadataList
//go:noescape
func HasFuncGetVolumeMetadataList() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetVolumeMetadataList
//go:noescape
func FuncGetVolumeMetadataList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetVolumeMetadataList
//go:noescape
func CallGetVolumeMetadataList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetVolumeMetadataList
//go:noescape
func TryGetVolumeMetadataList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GetVolumeRoot
//go:noescape
func HasFuncGetVolumeRoot() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GetVolumeRoot
//go:noescape
func FuncGetVolumeRoot(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GetVolumeRoot
//go:noescape
func CallGetVolumeRoot(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_GetVolumeRoot
//go:noescape
func TryGetVolumeRoot(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_GrantAccess
//go:noescape
func HasFuncGrantAccess() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_GrantAccess
//go:noescape
func FuncGrantAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_GrantAccess
//go:noescape
func CallGrantAccess(
	retPtr unsafe.Pointer,
	entryUrls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_GrantAccess
//go:noescape
func TryGrantAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entryUrls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ImportCrostiniImage
//go:noescape
func HasFuncImportCrostiniImage() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ImportCrostiniImage
//go:noescape
func FuncImportCrostiniImage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ImportCrostiniImage
//go:noescape
func CallImportCrostiniImage(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ImportCrostiniImage
//go:noescape
func TryImportCrostiniImage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_InstallLinuxPackage
//go:noescape
func HasFuncInstallLinuxPackage() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_InstallLinuxPackage
//go:noescape
func FuncInstallLinuxPackage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_InstallLinuxPackage
//go:noescape
func CallInstallLinuxPackage(
	retPtr unsafe.Pointer,
	entry js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_InstallLinuxPackage
//go:noescape
func TryInstallLinuxPackage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_InvokeSharesheet
//go:noescape
func HasFuncInvokeSharesheet() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_InvokeSharesheet
//go:noescape
func FuncInvokeSharesheet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_InvokeSharesheet
//go:noescape
func CallInvokeSharesheet(
	retPtr unsafe.Pointer,
	entries js.Ref,
	launchSource uint32,
	dlpSourceUrls js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_InvokeSharesheet
//go:noescape
func TryInvokeSharesheet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	launchSource uint32,
	dlpSourceUrls js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_IsTabletModeEnabled
//go:noescape
func HasFuncIsTabletModeEnabled() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_IsTabletModeEnabled
//go:noescape
func FuncIsTabletModeEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_IsTabletModeEnabled
//go:noescape
func CallIsTabletModeEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_IsTabletModeEnabled
//go:noescape
func TryIsTabletModeEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ListMountableGuests
//go:noescape
func HasFuncListMountableGuests() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ListMountableGuests
//go:noescape
func FuncListMountableGuests(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ListMountableGuests
//go:noescape
func CallListMountableGuests(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_ListMountableGuests
//go:noescape
func TryListMountableGuests(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_MountCrostini
//go:noescape
func HasFuncMountCrostini() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_MountCrostini
//go:noescape
func FuncMountCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_MountCrostini
//go:noescape
func CallMountCrostini(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_MountCrostini
//go:noescape
func TryMountCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_MountGuest
//go:noescape
func HasFuncMountGuest() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_MountGuest
//go:noescape
func FuncMountGuest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_MountGuest
//go:noescape
func CallMountGuest(
	retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/webext/filemanagerprivate try_MountGuest
//go:noescape
func TryMountGuest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_NotifyDriveDialogResult
//go:noescape
func HasFuncNotifyDriveDialogResult() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_NotifyDriveDialogResult
//go:noescape
func FuncNotifyDriveDialogResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_NotifyDriveDialogResult
//go:noescape
func CallNotifyDriveDialogResult(
	retPtr unsafe.Pointer,
	result uint32)

//go:wasmimport plat/js/webext/filemanagerprivate try_NotifyDriveDialogResult
//go:noescape
func TryNotifyDriveDialogResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	result uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnAppsUpdated
//go:noescape
func HasFuncOnAppsUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnAppsUpdated
//go:noescape
func FuncOnAppsUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnAppsUpdated
//go:noescape
func CallOnAppsUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnAppsUpdated
//go:noescape
func TryOnAppsUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffAppsUpdated
//go:noescape
func HasFuncOffAppsUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffAppsUpdated
//go:noescape
func FuncOffAppsUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffAppsUpdated
//go:noescape
func CallOffAppsUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffAppsUpdated
//go:noescape
func TryOffAppsUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnAppsUpdated
//go:noescape
func HasFuncHasOnAppsUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnAppsUpdated
//go:noescape
func FuncHasOnAppsUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnAppsUpdated
//go:noescape
func CallHasOnAppsUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnAppsUpdated
//go:noescape
func TryHasOnAppsUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnBulkPinProgress
//go:noescape
func HasFuncOnBulkPinProgress() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnBulkPinProgress
//go:noescape
func FuncOnBulkPinProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnBulkPinProgress
//go:noescape
func CallOnBulkPinProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnBulkPinProgress
//go:noescape
func TryOnBulkPinProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffBulkPinProgress
//go:noescape
func HasFuncOffBulkPinProgress() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffBulkPinProgress
//go:noescape
func FuncOffBulkPinProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffBulkPinProgress
//go:noescape
func CallOffBulkPinProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffBulkPinProgress
//go:noescape
func TryOffBulkPinProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnBulkPinProgress
//go:noescape
func HasFuncHasOnBulkPinProgress() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnBulkPinProgress
//go:noescape
func FuncHasOnBulkPinProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnBulkPinProgress
//go:noescape
func CallHasOnBulkPinProgress(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnBulkPinProgress
//go:noescape
func TryHasOnBulkPinProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnCrostiniChanged
//go:noescape
func HasFuncOnCrostiniChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnCrostiniChanged
//go:noescape
func FuncOnCrostiniChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnCrostiniChanged
//go:noescape
func CallOnCrostiniChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnCrostiniChanged
//go:noescape
func TryOnCrostiniChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffCrostiniChanged
//go:noescape
func HasFuncOffCrostiniChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffCrostiniChanged
//go:noescape
func FuncOffCrostiniChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffCrostiniChanged
//go:noescape
func CallOffCrostiniChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffCrostiniChanged
//go:noescape
func TryOffCrostiniChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnCrostiniChanged
//go:noescape
func HasFuncHasOnCrostiniChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnCrostiniChanged
//go:noescape
func FuncHasOnCrostiniChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnCrostiniChanged
//go:noescape
func CallHasOnCrostiniChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnCrostiniChanged
//go:noescape
func TryHasOnCrostiniChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDeviceChanged
//go:noescape
func HasFuncOnDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDeviceChanged
//go:noescape
func FuncOnDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDeviceChanged
//go:noescape
func CallOnDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDeviceChanged
//go:noescape
func TryOnDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDeviceChanged
//go:noescape
func HasFuncOffDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDeviceChanged
//go:noescape
func FuncOffDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDeviceChanged
//go:noescape
func CallOffDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDeviceChanged
//go:noescape
func TryOffDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDeviceChanged
//go:noescape
func HasFuncHasOnDeviceChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDeviceChanged
//go:noescape
func FuncHasOnDeviceChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDeviceChanged
//go:noescape
func CallHasOnDeviceChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDeviceChanged
//go:noescape
func TryHasOnDeviceChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDeviceConnectionStatusChanged
//go:noescape
func HasFuncOnDeviceConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDeviceConnectionStatusChanged
//go:noescape
func FuncOnDeviceConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDeviceConnectionStatusChanged
//go:noescape
func CallOnDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDeviceConnectionStatusChanged
//go:noescape
func TryOnDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDeviceConnectionStatusChanged
//go:noescape
func HasFuncOffDeviceConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDeviceConnectionStatusChanged
//go:noescape
func FuncOffDeviceConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDeviceConnectionStatusChanged
//go:noescape
func CallOffDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDeviceConnectionStatusChanged
//go:noescape
func TryOffDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDeviceConnectionStatusChanged
//go:noescape
func HasFuncHasOnDeviceConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDeviceConnectionStatusChanged
//go:noescape
func FuncHasOnDeviceConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDeviceConnectionStatusChanged
//go:noescape
func CallHasOnDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDeviceConnectionStatusChanged
//go:noescape
func TryHasOnDeviceConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDirectoryChanged
//go:noescape
func HasFuncOnDirectoryChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDirectoryChanged
//go:noescape
func FuncOnDirectoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDirectoryChanged
//go:noescape
func CallOnDirectoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDirectoryChanged
//go:noescape
func TryOnDirectoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDirectoryChanged
//go:noescape
func HasFuncOffDirectoryChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDirectoryChanged
//go:noescape
func FuncOffDirectoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDirectoryChanged
//go:noescape
func CallOffDirectoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDirectoryChanged
//go:noescape
func TryOffDirectoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDirectoryChanged
//go:noescape
func HasFuncHasOnDirectoryChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDirectoryChanged
//go:noescape
func FuncHasOnDirectoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDirectoryChanged
//go:noescape
func CallHasOnDirectoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDirectoryChanged
//go:noescape
func TryHasOnDirectoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDriveConfirmDialog
//go:noescape
func HasFuncOnDriveConfirmDialog() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDriveConfirmDialog
//go:noescape
func FuncOnDriveConfirmDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDriveConfirmDialog
//go:noescape
func CallOnDriveConfirmDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDriveConfirmDialog
//go:noescape
func TryOnDriveConfirmDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDriveConfirmDialog
//go:noescape
func HasFuncOffDriveConfirmDialog() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDriveConfirmDialog
//go:noescape
func FuncOffDriveConfirmDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDriveConfirmDialog
//go:noescape
func CallOffDriveConfirmDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDriveConfirmDialog
//go:noescape
func TryOffDriveConfirmDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDriveConfirmDialog
//go:noescape
func HasFuncHasOnDriveConfirmDialog() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDriveConfirmDialog
//go:noescape
func FuncHasOnDriveConfirmDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDriveConfirmDialog
//go:noescape
func CallHasOnDriveConfirmDialog(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDriveConfirmDialog
//go:noescape
func TryHasOnDriveConfirmDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDriveConnectionStatusChanged
//go:noescape
func HasFuncOnDriveConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDriveConnectionStatusChanged
//go:noescape
func FuncOnDriveConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDriveConnectionStatusChanged
//go:noescape
func CallOnDriveConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDriveConnectionStatusChanged
//go:noescape
func TryOnDriveConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDriveConnectionStatusChanged
//go:noescape
func HasFuncOffDriveConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDriveConnectionStatusChanged
//go:noescape
func FuncOffDriveConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDriveConnectionStatusChanged
//go:noescape
func CallOffDriveConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDriveConnectionStatusChanged
//go:noescape
func TryOffDriveConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDriveConnectionStatusChanged
//go:noescape
func HasFuncHasOnDriveConnectionStatusChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDriveConnectionStatusChanged
//go:noescape
func FuncHasOnDriveConnectionStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDriveConnectionStatusChanged
//go:noescape
func CallHasOnDriveConnectionStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDriveConnectionStatusChanged
//go:noescape
func TryHasOnDriveConnectionStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnDriveSyncError
//go:noescape
func HasFuncOnDriveSyncError() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnDriveSyncError
//go:noescape
func FuncOnDriveSyncError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnDriveSyncError
//go:noescape
func CallOnDriveSyncError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnDriveSyncError
//go:noescape
func TryOnDriveSyncError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffDriveSyncError
//go:noescape
func HasFuncOffDriveSyncError() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffDriveSyncError
//go:noescape
func FuncOffDriveSyncError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffDriveSyncError
//go:noescape
func CallOffDriveSyncError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffDriveSyncError
//go:noescape
func TryOffDriveSyncError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnDriveSyncError
//go:noescape
func HasFuncHasOnDriveSyncError() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnDriveSyncError
//go:noescape
func FuncHasOnDriveSyncError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnDriveSyncError
//go:noescape
func CallHasOnDriveSyncError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnDriveSyncError
//go:noescape
func TryHasOnDriveSyncError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnFileTransfersUpdated
//go:noescape
func HasFuncOnFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnFileTransfersUpdated
//go:noescape
func FuncOnFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnFileTransfersUpdated
//go:noescape
func CallOnFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnFileTransfersUpdated
//go:noescape
func TryOnFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffFileTransfersUpdated
//go:noescape
func HasFuncOffFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffFileTransfersUpdated
//go:noescape
func FuncOffFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffFileTransfersUpdated
//go:noescape
func CallOffFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffFileTransfersUpdated
//go:noescape
func TryOffFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnFileTransfersUpdated
//go:noescape
func HasFuncHasOnFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnFileTransfersUpdated
//go:noescape
func FuncHasOnFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnFileTransfersUpdated
//go:noescape
func CallHasOnFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnFileTransfersUpdated
//go:noescape
func TryHasOnFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnIOTaskProgressStatus
//go:noescape
func HasFuncOnIOTaskProgressStatus() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnIOTaskProgressStatus
//go:noescape
func FuncOnIOTaskProgressStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnIOTaskProgressStatus
//go:noescape
func CallOnIOTaskProgressStatus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnIOTaskProgressStatus
//go:noescape
func TryOnIOTaskProgressStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffIOTaskProgressStatus
//go:noescape
func HasFuncOffIOTaskProgressStatus() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffIOTaskProgressStatus
//go:noescape
func FuncOffIOTaskProgressStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffIOTaskProgressStatus
//go:noescape
func CallOffIOTaskProgressStatus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffIOTaskProgressStatus
//go:noescape
func TryOffIOTaskProgressStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnIOTaskProgressStatus
//go:noescape
func HasFuncHasOnIOTaskProgressStatus() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnIOTaskProgressStatus
//go:noescape
func FuncHasOnIOTaskProgressStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnIOTaskProgressStatus
//go:noescape
func CallHasOnIOTaskProgressStatus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnIOTaskProgressStatus
//go:noescape
func TryHasOnIOTaskProgressStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnIndividualFileTransfersUpdated
//go:noescape
func HasFuncOnIndividualFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnIndividualFileTransfersUpdated
//go:noescape
func FuncOnIndividualFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnIndividualFileTransfersUpdated
//go:noescape
func CallOnIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnIndividualFileTransfersUpdated
//go:noescape
func TryOnIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffIndividualFileTransfersUpdated
//go:noescape
func HasFuncOffIndividualFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffIndividualFileTransfersUpdated
//go:noescape
func FuncOffIndividualFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffIndividualFileTransfersUpdated
//go:noescape
func CallOffIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffIndividualFileTransfersUpdated
//go:noescape
func TryOffIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnIndividualFileTransfersUpdated
//go:noescape
func HasFuncHasOnIndividualFileTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnIndividualFileTransfersUpdated
//go:noescape
func FuncHasOnIndividualFileTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnIndividualFileTransfersUpdated
//go:noescape
func CallHasOnIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnIndividualFileTransfersUpdated
//go:noescape
func TryHasOnIndividualFileTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnMountCompleted
//go:noescape
func HasFuncOnMountCompleted() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnMountCompleted
//go:noescape
func FuncOnMountCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnMountCompleted
//go:noescape
func CallOnMountCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnMountCompleted
//go:noescape
func TryOnMountCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffMountCompleted
//go:noescape
func HasFuncOffMountCompleted() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffMountCompleted
//go:noescape
func FuncOffMountCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffMountCompleted
//go:noescape
func CallOffMountCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffMountCompleted
//go:noescape
func TryOffMountCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnMountCompleted
//go:noescape
func HasFuncHasOnMountCompleted() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnMountCompleted
//go:noescape
func FuncHasOnMountCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnMountCompleted
//go:noescape
func CallHasOnMountCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnMountCompleted
//go:noescape
func TryHasOnMountCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnMountableGuestsChanged
//go:noescape
func HasFuncOnMountableGuestsChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnMountableGuestsChanged
//go:noescape
func FuncOnMountableGuestsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnMountableGuestsChanged
//go:noescape
func CallOnMountableGuestsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnMountableGuestsChanged
//go:noescape
func TryOnMountableGuestsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffMountableGuestsChanged
//go:noescape
func HasFuncOffMountableGuestsChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffMountableGuestsChanged
//go:noescape
func FuncOffMountableGuestsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffMountableGuestsChanged
//go:noescape
func CallOffMountableGuestsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffMountableGuestsChanged
//go:noescape
func TryOffMountableGuestsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnMountableGuestsChanged
//go:noescape
func HasFuncHasOnMountableGuestsChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnMountableGuestsChanged
//go:noescape
func FuncHasOnMountableGuestsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnMountableGuestsChanged
//go:noescape
func CallHasOnMountableGuestsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnMountableGuestsChanged
//go:noescape
func TryHasOnMountableGuestsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnPinTransfersUpdated
//go:noescape
func HasFuncOnPinTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnPinTransfersUpdated
//go:noescape
func FuncOnPinTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnPinTransfersUpdated
//go:noescape
func CallOnPinTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnPinTransfersUpdated
//go:noescape
func TryOnPinTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffPinTransfersUpdated
//go:noescape
func HasFuncOffPinTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffPinTransfersUpdated
//go:noescape
func FuncOffPinTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffPinTransfersUpdated
//go:noescape
func CallOffPinTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffPinTransfersUpdated
//go:noescape
func TryOffPinTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnPinTransfersUpdated
//go:noescape
func HasFuncHasOnPinTransfersUpdated() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnPinTransfersUpdated
//go:noescape
func FuncHasOnPinTransfersUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnPinTransfersUpdated
//go:noescape
func CallHasOnPinTransfersUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnPinTransfersUpdated
//go:noescape
func TryHasOnPinTransfersUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnPreferencesChanged
//go:noescape
func HasFuncOnPreferencesChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnPreferencesChanged
//go:noescape
func FuncOnPreferencesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnPreferencesChanged
//go:noescape
func CallOnPreferencesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnPreferencesChanged
//go:noescape
func TryOnPreferencesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffPreferencesChanged
//go:noescape
func HasFuncOffPreferencesChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffPreferencesChanged
//go:noescape
func FuncOffPreferencesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffPreferencesChanged
//go:noescape
func CallOffPreferencesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffPreferencesChanged
//go:noescape
func TryOffPreferencesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnPreferencesChanged
//go:noescape
func HasFuncHasOnPreferencesChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnPreferencesChanged
//go:noescape
func FuncHasOnPreferencesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnPreferencesChanged
//go:noescape
func CallHasOnPreferencesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnPreferencesChanged
//go:noescape
func TryHasOnPreferencesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OnTabletModeChanged
//go:noescape
func HasFuncOnTabletModeChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OnTabletModeChanged
//go:noescape
func FuncOnTabletModeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OnTabletModeChanged
//go:noescape
func CallOnTabletModeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OnTabletModeChanged
//go:noescape
func TryOnTabletModeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OffTabletModeChanged
//go:noescape
func HasFuncOffTabletModeChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OffTabletModeChanged
//go:noescape
func FuncOffTabletModeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OffTabletModeChanged
//go:noescape
func CallOffTabletModeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OffTabletModeChanged
//go:noescape
func TryOffTabletModeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_HasOnTabletModeChanged
//go:noescape
func HasFuncHasOnTabletModeChanged() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_HasOnTabletModeChanged
//go:noescape
func FuncHasOnTabletModeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_HasOnTabletModeChanged
//go:noescape
func CallHasOnTabletModeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_HasOnTabletModeChanged
//go:noescape
func TryHasOnTabletModeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OpenInspector
//go:noescape
func HasFuncOpenInspector() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OpenInspector
//go:noescape
func FuncOpenInspector(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OpenInspector
//go:noescape
func CallOpenInspector(
	retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/webext/filemanagerprivate try_OpenInspector
//go:noescape
func TryOpenInspector(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OpenManageSyncSettings
//go:noescape
func HasFuncOpenManageSyncSettings() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OpenManageSyncSettings
//go:noescape
func FuncOpenManageSyncSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OpenManageSyncSettings
//go:noescape
func CallOpenManageSyncSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_OpenManageSyncSettings
//go:noescape
func TryOpenManageSyncSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OpenSettingsSubpage
//go:noescape
func HasFuncOpenSettingsSubpage() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OpenSettingsSubpage
//go:noescape
func FuncOpenSettingsSubpage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OpenSettingsSubpage
//go:noescape
func CallOpenSettingsSubpage(
	retPtr unsafe.Pointer,
	sub_page js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OpenSettingsSubpage
//go:noescape
func TryOpenSettingsSubpage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sub_page js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OpenURL
//go:noescape
func HasFuncOpenURL() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OpenURL
//go:noescape
func FuncOpenURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OpenURL
//go:noescape
func CallOpenURL(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_OpenURL
//go:noescape
func TryOpenURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_OpenWindow
//go:noescape
func HasFuncOpenWindow() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_OpenWindow
//go:noescape
func FuncOpenWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_OpenWindow
//go:noescape
func CallOpenWindow(
	retPtr unsafe.Pointer,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_OpenWindow
//go:noescape
func TryOpenWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ParseTrashInfoFiles
//go:noescape
func HasFuncParseTrashInfoFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ParseTrashInfoFiles
//go:noescape
func FuncParseTrashInfoFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ParseTrashInfoFiles
//go:noescape
func CallParseTrashInfoFiles(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ParseTrashInfoFiles
//go:noescape
func TryParseTrashInfoFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_PinDriveFile
//go:noescape
func HasFuncPinDriveFile() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_PinDriveFile
//go:noescape
func FuncPinDriveFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_PinDriveFile
//go:noescape
func CallPinDriveFile(
	retPtr unsafe.Pointer,
	entry js.Ref,
	pin js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_PinDriveFile
//go:noescape
func TryPinDriveFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref,
	pin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_PollDriveHostedFilePinStates
//go:noescape
func HasFuncPollDriveHostedFilePinStates() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_PollDriveHostedFilePinStates
//go:noescape
func FuncPollDriveHostedFilePinStates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_PollDriveHostedFilePinStates
//go:noescape
func CallPollDriveHostedFilePinStates(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_PollDriveHostedFilePinStates
//go:noescape
func TryPollDriveHostedFilePinStates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ProgressPausedTasks
//go:noescape
func HasFuncProgressPausedTasks() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ProgressPausedTasks
//go:noescape
func FuncProgressPausedTasks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ProgressPausedTasks
//go:noescape
func CallProgressPausedTasks(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_ProgressPausedTasks
//go:noescape
func TryProgressPausedTasks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_RemoveFileWatch
//go:noescape
func HasFuncRemoveFileWatch() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_RemoveFileWatch
//go:noescape
func FuncRemoveFileWatch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_RemoveFileWatch
//go:noescape
func CallRemoveFileWatch(
	retPtr unsafe.Pointer,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_RemoveFileWatch
//go:noescape
func TryRemoveFileWatch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_RemoveMount
//go:noescape
func HasFuncRemoveMount() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_RemoveMount
//go:noescape
func FuncRemoveMount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_RemoveMount
//go:noescape
func CallRemoveMount(
	retPtr unsafe.Pointer,
	volumeId js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_RemoveMount
//go:noescape
func TryRemoveMount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_RenameVolume
//go:noescape
func HasFuncRenameVolume() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_RenameVolume
//go:noescape
func FuncRenameVolume(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_RenameVolume
//go:noescape
func CallRenameVolume(
	retPtr unsafe.Pointer,
	volumeId js.Ref,
	newName js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_RenameVolume
//go:noescape
func TryRenameVolume(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref,
	newName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ResolveIsolatedEntries
//go:noescape
func HasFuncResolveIsolatedEntries() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ResolveIsolatedEntries
//go:noescape
func FuncResolveIsolatedEntries(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ResolveIsolatedEntries
//go:noescape
func CallResolveIsolatedEntries(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ResolveIsolatedEntries
//go:noescape
func TryResolveIsolatedEntries(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ResumeIOTask
//go:noescape
func HasFuncResumeIOTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ResumeIOTask
//go:noescape
func FuncResumeIOTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ResumeIOTask
//go:noescape
func CallResumeIOTask(
	retPtr unsafe.Pointer,
	taskId int32,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_ResumeIOTask
//go:noescape
func TryResumeIOTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	taskId int32,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SearchDrive
//go:noescape
func HasFuncSearchDrive() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SearchDrive
//go:noescape
func FuncSearchDrive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SearchDrive
//go:noescape
func CallSearchDrive(
	retPtr unsafe.Pointer,
	searchParams unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SearchDrive
//go:noescape
func TrySearchDrive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchParams unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SearchDriveMetadata
//go:noescape
func HasFuncSearchDriveMetadata() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SearchDriveMetadata
//go:noescape
func FuncSearchDriveMetadata(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SearchDriveMetadata
//go:noescape
func CallSearchDriveMetadata(
	retPtr unsafe.Pointer,
	searchParams unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_SearchDriveMetadata
//go:noescape
func TrySearchDriveMetadata(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SearchFiles
//go:noescape
func HasFuncSearchFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SearchFiles
//go:noescape
func FuncSearchFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SearchFiles
//go:noescape
func CallSearchFiles(
	retPtr unsafe.Pointer,
	searchParams unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_SearchFiles
//go:noescape
func TrySearchFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SearchFilesByHashes
//go:noescape
func HasFuncSearchFilesByHashes() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SearchFilesByHashes
//go:noescape
func FuncSearchFilesByHashes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SearchFilesByHashes
//go:noescape
func CallSearchFilesByHashes(
	retPtr unsafe.Pointer,
	volumeId js.Ref,
	hashList js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SearchFilesByHashes
//go:noescape
func TrySearchFilesByHashes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	volumeId js.Ref,
	hashList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SelectAndroidPickerApp
//go:noescape
func HasFuncSelectAndroidPickerApp() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SelectAndroidPickerApp
//go:noescape
func FuncSelectAndroidPickerApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SelectAndroidPickerApp
//go:noescape
func CallSelectAndroidPickerApp(
	retPtr unsafe.Pointer,
	androidApp unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_SelectAndroidPickerApp
//go:noescape
func TrySelectAndroidPickerApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	androidApp unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SelectFile
//go:noescape
func HasFuncSelectFile() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SelectFile
//go:noescape
func FuncSelectFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SelectFile
//go:noescape
func CallSelectFile(
	retPtr unsafe.Pointer,
	selectedPath js.Ref,
	index int32,
	forOpening js.Ref,
	shouldReturnLocalPath js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SelectFile
//go:noescape
func TrySelectFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectedPath js.Ref,
	index int32,
	forOpening js.Ref,
	shouldReturnLocalPath js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SelectFiles
//go:noescape
func HasFuncSelectFiles() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SelectFiles
//go:noescape
func FuncSelectFiles(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SelectFiles
//go:noescape
func CallSelectFiles(
	retPtr unsafe.Pointer,
	selectedPaths js.Ref,
	shouldReturnLocalPath js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SelectFiles
//go:noescape
func TrySelectFiles(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectedPaths js.Ref,
	shouldReturnLocalPath js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SendFeedback
//go:noescape
func HasFuncSendFeedback() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SendFeedback
//go:noescape
func FuncSendFeedback(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SendFeedback
//go:noescape
func CallSendFeedback(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_SendFeedback
//go:noescape
func TrySendFeedback(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SetDefaultTask
//go:noescape
func HasFuncSetDefaultTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SetDefaultTask
//go:noescape
func FuncSetDefaultTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SetDefaultTask
//go:noescape
func CallSetDefaultTask(
	retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	entries js.Ref,
	mimeTypes js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SetDefaultTask
//go:noescape
func TrySetDefaultTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	entries js.Ref,
	mimeTypes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SetPreferences
//go:noescape
func HasFuncSetPreferences() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SetPreferences
//go:noescape
func FuncSetPreferences(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SetPreferences
//go:noescape
func CallSetPreferences(
	retPtr unsafe.Pointer,
	changeInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_SetPreferences
//go:noescape
func TrySetPreferences(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	changeInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SharePathsWithCrostini
//go:noescape
func HasFuncSharePathsWithCrostini() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SharePathsWithCrostini
//go:noescape
func FuncSharePathsWithCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SharePathsWithCrostini
//go:noescape
func CallSharePathsWithCrostini(
	retPtr unsafe.Pointer,
	vmName js.Ref,
	entries js.Ref,
	persist js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SharePathsWithCrostini
//go:noescape
func TrySharePathsWithCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vmName js.Ref,
	entries js.Ref,
	persist js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SharesheetHasTargets
//go:noescape
func HasFuncSharesheetHasTargets() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SharesheetHasTargets
//go:noescape
func FuncSharesheetHasTargets(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SharesheetHasTargets
//go:noescape
func CallSharesheetHasTargets(
	retPtr unsafe.Pointer,
	entries js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SharesheetHasTargets
//go:noescape
func TrySharesheetHasTargets(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ShowPolicyDialog
//go:noescape
func HasFuncShowPolicyDialog() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ShowPolicyDialog
//go:noescape
func FuncShowPolicyDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ShowPolicyDialog
//go:noescape
func CallShowPolicyDialog(
	retPtr unsafe.Pointer,
	taskId int32,
	typ uint32)

//go:wasmimport plat/js/webext/filemanagerprivate try_ShowPolicyDialog
//go:noescape
func TryShowPolicyDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	taskId int32,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_SinglePartitionFormat
//go:noescape
func HasFuncSinglePartitionFormat() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_SinglePartitionFormat
//go:noescape
func FuncSinglePartitionFormat(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_SinglePartitionFormat
//go:noescape
func CallSinglePartitionFormat(
	retPtr unsafe.Pointer,
	deviceStoragePath js.Ref,
	filesystem uint32,
	volumeLabel js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_SinglePartitionFormat
//go:noescape
func TrySinglePartitionFormat(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceStoragePath js.Ref,
	filesystem uint32,
	volumeLabel js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_StartIOTask
//go:noescape
func HasFuncStartIOTask() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_StartIOTask
//go:noescape
func FuncStartIOTask(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_StartIOTask
//go:noescape
func CallStartIOTask(
	retPtr unsafe.Pointer,
	typ uint32,
	entries js.Ref,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate try_StartIOTask
//go:noescape
func TryStartIOTask(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	entries js.Ref,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ToggleAddedToHoldingSpace
//go:noescape
func HasFuncToggleAddedToHoldingSpace() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ToggleAddedToHoldingSpace
//go:noescape
func FuncToggleAddedToHoldingSpace(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ToggleAddedToHoldingSpace
//go:noescape
func CallToggleAddedToHoldingSpace(
	retPtr unsafe.Pointer,
	entries js.Ref,
	added js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ToggleAddedToHoldingSpace
//go:noescape
func TryToggleAddedToHoldingSpace(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	entries js.Ref,
	added js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_UnsharePathWithCrostini
//go:noescape
func HasFuncUnsharePathWithCrostini() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_UnsharePathWithCrostini
//go:noescape
func FuncUnsharePathWithCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_UnsharePathWithCrostini
//go:noescape
func CallUnsharePathWithCrostini(
	retPtr unsafe.Pointer,
	vmName js.Ref,
	entry js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_UnsharePathWithCrostini
//go:noescape
func TryUnsharePathWithCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vmName js.Ref,
	entry js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_ValidatePathNameLength
//go:noescape
func HasFuncValidatePathNameLength() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_ValidatePathNameLength
//go:noescape
func FuncValidatePathNameLength(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_ValidatePathNameLength
//go:noescape
func CallValidatePathNameLength(
	retPtr unsafe.Pointer,
	parentEntry js.Ref,
	name js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate try_ValidatePathNameLength
//go:noescape
func TryValidatePathNameLength(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentEntry js.Ref,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filemanagerprivate has_Zoom
//go:noescape
func HasFuncZoom() js.Ref

//go:wasmimport plat/js/webext/filemanagerprivate func_Zoom
//go:noescape
func FuncZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filemanagerprivate call_Zoom
//go:noescape
func CallZoom(
	retPtr unsafe.Pointer,
	operation uint32)

//go:wasmimport plat/js/webext/filemanagerprivate try_Zoom
//go:noescape
func TryZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	operation uint32) (ok js.Ref)
