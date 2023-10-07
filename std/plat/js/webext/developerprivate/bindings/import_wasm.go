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

//go:wasmimport plat/js/webext/developerprivate store_AccessModifier
//go:noescape
func AccessModifierJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_AccessModifier
//go:noescape
func AccessModifierJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_CommandScope
//go:noescape
func ConstOfCommandScope(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_Command
//go:noescape
func CommandJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_Command
//go:noescape
func CommandJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ControlledInfo
//go:noescape
func ControlledInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ControlledInfo
//go:noescape
func ControlledInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_ErrorType
//go:noescape
func ConstOfErrorType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_DeleteExtensionErrorsProperties
//go:noescape
func DeleteExtensionErrorsPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_DeleteExtensionErrorsProperties
//go:noescape
func DeleteExtensionErrorsPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_DependentExtension
//go:noescape
func DependentExtensionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_DependentExtension
//go:noescape
func DependentExtensionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_DisableReasons
//go:noescape
func DisableReasonsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_DisableReasons
//go:noescape
func DisableReasonsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ErrorFileSource
//go:noescape
func ErrorFileSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ErrorFileSource
//go:noescape
func ErrorFileSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_ErrorLevel
//go:noescape
func ConstOfErrorLevel(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate constof_EventType
//go:noescape
func ConstOfEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_SafetyCheckStrings
//go:noescape
func SafetyCheckStringsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_SafetyCheckStrings
//go:noescape
func SafetyCheckStringsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_HomePage
//go:noescape
func HomePageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_HomePage
//go:noescape
func HomePageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_Location
//go:noescape
func ConstOfLocation(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_ManifestError
//go:noescape
func ManifestErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ManifestError
//go:noescape
func ManifestErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_OptionsPage
//go:noescape
func OptionsPageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_OptionsPage
//go:noescape
func OptionsPageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_Permission
//go:noescape
func PermissionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_Permission
//go:noescape
func PermissionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_HostAccess
//go:noescape
func ConstOfHostAccess(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_SiteControl
//go:noescape
func SiteControlJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_SiteControl
//go:noescape
func SiteControlJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_RuntimeHostPermissions
//go:noescape
func RuntimeHostPermissionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_RuntimeHostPermissions
//go:noescape
func RuntimeHostPermissionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_Permissions
//go:noescape
func PermissionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_Permissions
//go:noescape
func PermissionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_StackFrame
//go:noescape
func StackFrameJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_StackFrame
//go:noescape
func StackFrameJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_RuntimeError
//go:noescape
func RuntimeErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_RuntimeError
//go:noescape
func RuntimeErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_ExtensionState
//go:noescape
func ConstOfExtensionState(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate constof_ExtensionType
//go:noescape
func ConstOfExtensionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate constof_ViewType
//go:noescape
func ConstOfViewType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_ExtensionView
//go:noescape
func ExtensionViewJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ExtensionView
//go:noescape
func ExtensionViewJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ExtensionInfo
//go:noescape
func ExtensionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ExtensionInfo
//go:noescape
func ExtensionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_EventData
//go:noescape
func EventDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_EventData
//go:noescape
func EventDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ExtensionCommandUpdate
//go:noescape
func ExtensionCommandUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ExtensionCommandUpdate
//go:noescape
func ExtensionCommandUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ExtensionConfigurationUpdate
//go:noescape
func ExtensionConfigurationUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ExtensionConfigurationUpdate
//go:noescape
func ExtensionConfigurationUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ExtensionSiteAccessUpdate
//go:noescape
func ExtensionSiteAccessUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ExtensionSiteAccessUpdate
//go:noescape
func ExtensionSiteAccessUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_FileType
//go:noescape
func ConstOfFileType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_GetExtensionsInfoOptions
//go:noescape
func GetExtensionsInfoOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_GetExtensionsInfoOptions
//go:noescape
func GetExtensionsInfoOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_MatchingExtensionInfo
//go:noescape
func MatchingExtensionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_MatchingExtensionInfo
//go:noescape
func MatchingExtensionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ProjectInfo
//go:noescape
func ProjectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ProjectInfo
//go:noescape
func ProjectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_InspectOptions
//go:noescape
func InspectOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_InspectOptions
//go:noescape
func InspectOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_InstallWarning
//go:noescape
func InstallWarningJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_InstallWarning
//go:noescape
func InstallWarningJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_ItemType
//go:noescape
func ConstOfItemType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_ItemInspectView
//go:noescape
func ItemInspectViewJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ItemInspectView
//go:noescape
func ItemInspectViewJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ItemInfo
//go:noescape
func ItemInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ItemInfo
//go:noescape
func ItemInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_LoadError
//go:noescape
func LoadErrorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_LoadError
//go:noescape
func LoadErrorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_LoadUnpackedOptions
//go:noescape
func LoadUnpackedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_LoadUnpackedOptions
//go:noescape
func LoadUnpackedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_OpenDevToolsProperties
//go:noescape
func OpenDevToolsPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_OpenDevToolsProperties
//go:noescape
func OpenDevToolsPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_PackStatus
//go:noescape
func ConstOfPackStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_PackDirectoryResponse
//go:noescape
func PackDirectoryResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_PackDirectoryResponse
//go:noescape
func PackDirectoryResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ProfileConfigurationUpdate
//go:noescape
func ProfileConfigurationUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ProfileConfigurationUpdate
//go:noescape
func ProfileConfigurationUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ProfileInfo
//go:noescape
func ProfileInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ProfileInfo
//go:noescape
func ProfileInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_ReloadOptions
//go:noescape
func ReloadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_ReloadOptions
//go:noescape
func ReloadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_RequestFileSourceResponse
//go:noescape
func RequestFileSourceResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_RequestFileSourceResponse
//go:noescape
func RequestFileSourceResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_RequestFileSourceProperties
//go:noescape
func RequestFileSourcePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_RequestFileSourceProperties
//go:noescape
func RequestFileSourcePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate constof_SelectType
//go:noescape
func ConstOfSelectType(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate constof_SiteSet
//go:noescape
func ConstOfSiteSet(str js.Ref) uint32

//go:wasmimport plat/js/webext/developerprivate store_SiteInfo
//go:noescape
func SiteInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_SiteInfo
//go:noescape
func SiteInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_SiteGroup
//go:noescape
func SiteGroupJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_SiteGroup
//go:noescape
func SiteGroupJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_UserSiteSettings
//go:noescape
func UserSiteSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_UserSiteSettings
//go:noescape
func UserSiteSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate store_UserSiteSettingsOptions
//go:noescape
func UserSiteSettingsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/developerprivate load_UserSiteSettingsOptions
//go:noescape
func UserSiteSettingsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/developerprivate has_AddHostPermission
//go:noescape
func HasFuncAddHostPermission() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_AddHostPermission
//go:noescape
func FuncAddHostPermission(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_AddHostPermission
//go:noescape
func CallAddHostPermission(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	host js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_AddHostPermission
//go:noescape
func TryAddHostPermission(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	host js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_AddUserSpecifiedSites
//go:noescape
func HasFuncAddUserSpecifiedSites() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_AddUserSpecifiedSites
//go:noescape
func FuncAddUserSpecifiedSites(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_AddUserSpecifiedSites
//go:noescape
func CallAddUserSpecifiedSites(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_AddUserSpecifiedSites
//go:noescape
func TryAddUserSpecifiedSites(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_AutoUpdate
//go:noescape
func HasFuncAutoUpdate() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_AutoUpdate
//go:noescape
func FuncAutoUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_AutoUpdate
//go:noescape
func CallAutoUpdate(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_AutoUpdate
//go:noescape
func TryAutoUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_ChoosePath
//go:noescape
func HasFuncChoosePath() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_ChoosePath
//go:noescape
func FuncChoosePath(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_ChoosePath
//go:noescape
func CallChoosePath(
	retPtr unsafe.Pointer,
	selectType uint32,
	fileType uint32)

//go:wasmimport plat/js/webext/developerprivate try_ChoosePath
//go:noescape
func TryChoosePath(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectType uint32,
	fileType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_DeleteExtensionErrors
//go:noescape
func HasFuncDeleteExtensionErrors() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_DeleteExtensionErrors
//go:noescape
func FuncDeleteExtensionErrors(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_DeleteExtensionErrors
//go:noescape
func CallDeleteExtensionErrors(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_DeleteExtensionErrors
//go:noescape
func TryDeleteExtensionErrors(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetExtensionInfo
//go:noescape
func HasFuncGetExtensionInfo() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetExtensionInfo
//go:noescape
func FuncGetExtensionInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetExtensionInfo
//go:noescape
func CallGetExtensionInfo(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_GetExtensionInfo
//go:noescape
func TryGetExtensionInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetExtensionSize
//go:noescape
func HasFuncGetExtensionSize() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetExtensionSize
//go:noescape
func FuncGetExtensionSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetExtensionSize
//go:noescape
func CallGetExtensionSize(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_GetExtensionSize
//go:noescape
func TryGetExtensionSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetExtensionsInfo
//go:noescape
func HasFuncGetExtensionsInfo() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetExtensionsInfo
//go:noescape
func FuncGetExtensionsInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetExtensionsInfo
//go:noescape
func CallGetExtensionsInfo(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_GetExtensionsInfo
//go:noescape
func TryGetExtensionsInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetMatchingExtensionsForSite
//go:noescape
func HasFuncGetMatchingExtensionsForSite() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetMatchingExtensionsForSite
//go:noescape
func FuncGetMatchingExtensionsForSite(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetMatchingExtensionsForSite
//go:noescape
func CallGetMatchingExtensionsForSite(
	retPtr unsafe.Pointer,
	site js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_GetMatchingExtensionsForSite
//go:noescape
func TryGetMatchingExtensionsForSite(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	site js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetProfileConfiguration
//go:noescape
func HasFuncGetProfileConfiguration() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetProfileConfiguration
//go:noescape
func FuncGetProfileConfiguration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetProfileConfiguration
//go:noescape
func CallGetProfileConfiguration(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_GetProfileConfiguration
//go:noescape
func TryGetProfileConfiguration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetUserAndExtensionSitesByEtld
//go:noescape
func HasFuncGetUserAndExtensionSitesByEtld() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetUserAndExtensionSitesByEtld
//go:noescape
func FuncGetUserAndExtensionSitesByEtld(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetUserAndExtensionSitesByEtld
//go:noescape
func CallGetUserAndExtensionSitesByEtld(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_GetUserAndExtensionSitesByEtld
//go:noescape
func TryGetUserAndExtensionSitesByEtld(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_GetUserSiteSettings
//go:noescape
func HasFuncGetUserSiteSettings() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_GetUserSiteSettings
//go:noescape
func FuncGetUserSiteSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_GetUserSiteSettings
//go:noescape
func CallGetUserSiteSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_GetUserSiteSettings
//go:noescape
func TryGetUserSiteSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_Inspect
//go:noescape
func HasFuncInspect() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_Inspect
//go:noescape
func FuncInspect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_Inspect
//go:noescape
func CallInspect(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_Inspect
//go:noescape
func TryInspect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_InstallDroppedFile
//go:noescape
func HasFuncInstallDroppedFile() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_InstallDroppedFile
//go:noescape
func FuncInstallDroppedFile(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_InstallDroppedFile
//go:noescape
func CallInstallDroppedFile(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_InstallDroppedFile
//go:noescape
func TryInstallDroppedFile(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_IsProfileManaged
//go:noescape
func HasFuncIsProfileManaged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_IsProfileManaged
//go:noescape
func FuncIsProfileManaged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_IsProfileManaged
//go:noescape
func CallIsProfileManaged(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_IsProfileManaged
//go:noescape
func TryIsProfileManaged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_LoadDirectory
//go:noescape
func HasFuncLoadDirectory() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_LoadDirectory
//go:noescape
func FuncLoadDirectory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_LoadDirectory
//go:noescape
func CallLoadDirectory(
	retPtr unsafe.Pointer,
	directory js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_LoadDirectory
//go:noescape
func TryLoadDirectory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	directory js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_LoadUnpacked
//go:noescape
func HasFuncLoadUnpacked() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_LoadUnpacked
//go:noescape
func FuncLoadUnpacked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_LoadUnpacked
//go:noescape
func CallLoadUnpacked(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_LoadUnpacked
//go:noescape
func TryLoadUnpacked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_NotifyDragInstallInProgress
//go:noescape
func HasFuncNotifyDragInstallInProgress() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_NotifyDragInstallInProgress
//go:noescape
func FuncNotifyDragInstallInProgress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_NotifyDragInstallInProgress
//go:noescape
func CallNotifyDragInstallInProgress(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_NotifyDragInstallInProgress
//go:noescape
func TryNotifyDragInstallInProgress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OnItemStateChanged
//go:noescape
func HasFuncOnItemStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OnItemStateChanged
//go:noescape
func FuncOnItemStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OnItemStateChanged
//go:noescape
func CallOnItemStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OnItemStateChanged
//go:noescape
func TryOnItemStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OffItemStateChanged
//go:noescape
func HasFuncOffItemStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OffItemStateChanged
//go:noescape
func FuncOffItemStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OffItemStateChanged
//go:noescape
func CallOffItemStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OffItemStateChanged
//go:noescape
func TryOffItemStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_HasOnItemStateChanged
//go:noescape
func HasFuncHasOnItemStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_HasOnItemStateChanged
//go:noescape
func FuncHasOnItemStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_HasOnItemStateChanged
//go:noescape
func CallHasOnItemStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_HasOnItemStateChanged
//go:noescape
func TryHasOnItemStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OnProfileStateChanged
//go:noescape
func HasFuncOnProfileStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OnProfileStateChanged
//go:noescape
func FuncOnProfileStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OnProfileStateChanged
//go:noescape
func CallOnProfileStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OnProfileStateChanged
//go:noescape
func TryOnProfileStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OffProfileStateChanged
//go:noescape
func HasFuncOffProfileStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OffProfileStateChanged
//go:noescape
func FuncOffProfileStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OffProfileStateChanged
//go:noescape
func CallOffProfileStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OffProfileStateChanged
//go:noescape
func TryOffProfileStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_HasOnProfileStateChanged
//go:noescape
func HasFuncHasOnProfileStateChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_HasOnProfileStateChanged
//go:noescape
func FuncHasOnProfileStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_HasOnProfileStateChanged
//go:noescape
func CallHasOnProfileStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_HasOnProfileStateChanged
//go:noescape
func TryHasOnProfileStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OnUserSiteSettingsChanged
//go:noescape
func HasFuncOnUserSiteSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OnUserSiteSettingsChanged
//go:noescape
func FuncOnUserSiteSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OnUserSiteSettingsChanged
//go:noescape
func CallOnUserSiteSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OnUserSiteSettingsChanged
//go:noescape
func TryOnUserSiteSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OffUserSiteSettingsChanged
//go:noescape
func HasFuncOffUserSiteSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OffUserSiteSettingsChanged
//go:noescape
func FuncOffUserSiteSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OffUserSiteSettingsChanged
//go:noescape
func CallOffUserSiteSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_OffUserSiteSettingsChanged
//go:noescape
func TryOffUserSiteSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_HasOnUserSiteSettingsChanged
//go:noescape
func HasFuncHasOnUserSiteSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_HasOnUserSiteSettingsChanged
//go:noescape
func FuncHasOnUserSiteSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_HasOnUserSiteSettingsChanged
//go:noescape
func CallHasOnUserSiteSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_HasOnUserSiteSettingsChanged
//go:noescape
func TryHasOnUserSiteSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_OpenDevTools
//go:noescape
func HasFuncOpenDevTools() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_OpenDevTools
//go:noescape
func FuncOpenDevTools(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_OpenDevTools
//go:noescape
func CallOpenDevTools(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_OpenDevTools
//go:noescape
func TryOpenDevTools(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_PackDirectory
//go:noescape
func HasFuncPackDirectory() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_PackDirectory
//go:noescape
func FuncPackDirectory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_PackDirectory
//go:noescape
func CallPackDirectory(
	retPtr unsafe.Pointer,
	path js.Ref,
	privateKeyPath js.Ref,
	flags int32)

//go:wasmimport plat/js/webext/developerprivate try_PackDirectory
//go:noescape
func TryPackDirectory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref,
	privateKeyPath js.Ref,
	flags int32) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_Reload
//go:noescape
func HasFuncReload() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_Reload
//go:noescape
func FuncReload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_Reload
//go:noescape
func CallReload(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_Reload
//go:noescape
func TryReload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_RemoveHostPermission
//go:noescape
func HasFuncRemoveHostPermission() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_RemoveHostPermission
//go:noescape
func FuncRemoveHostPermission(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_RemoveHostPermission
//go:noescape
func CallRemoveHostPermission(
	retPtr unsafe.Pointer,
	extensionId js.Ref,
	host js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_RemoveHostPermission
//go:noescape
func TryRemoveHostPermission(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref,
	host js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_RemoveMultipleExtensions
//go:noescape
func HasFuncRemoveMultipleExtensions() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_RemoveMultipleExtensions
//go:noescape
func FuncRemoveMultipleExtensions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_RemoveMultipleExtensions
//go:noescape
func CallRemoveMultipleExtensions(
	retPtr unsafe.Pointer,
	extensionIds js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_RemoveMultipleExtensions
//go:noescape
func TryRemoveMultipleExtensions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_RemoveUserSpecifiedSites
//go:noescape
func HasFuncRemoveUserSpecifiedSites() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_RemoveUserSpecifiedSites
//go:noescape
func FuncRemoveUserSpecifiedSites(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_RemoveUserSpecifiedSites
//go:noescape
func CallRemoveUserSpecifiedSites(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_RemoveUserSpecifiedSites
//go:noescape
func TryRemoveUserSpecifiedSites(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_RepairExtension
//go:noescape
func HasFuncRepairExtension() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_RepairExtension
//go:noescape
func FuncRepairExtension(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_RepairExtension
//go:noescape
func CallRepairExtension(
	retPtr unsafe.Pointer,
	extensionId js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_RepairExtension
//go:noescape
func TryRepairExtension(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_RequestFileSource
//go:noescape
func HasFuncRequestFileSource() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_RequestFileSource
//go:noescape
func FuncRequestFileSource(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_RequestFileSource
//go:noescape
func CallRequestFileSource(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_RequestFileSource
//go:noescape
func TryRequestFileSource(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_SetShortcutHandlingSuspended
//go:noescape
func HasFuncSetShortcutHandlingSuspended() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_SetShortcutHandlingSuspended
//go:noescape
func FuncSetShortcutHandlingSuspended(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_SetShortcutHandlingSuspended
//go:noescape
func CallSetShortcutHandlingSuspended(
	retPtr unsafe.Pointer,
	isSuspended js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_SetShortcutHandlingSuspended
//go:noescape
func TrySetShortcutHandlingSuspended(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	isSuspended js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_ShowOptions
//go:noescape
func HasFuncShowOptions() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_ShowOptions
//go:noescape
func FuncShowOptions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_ShowOptions
//go:noescape
func CallShowOptions(
	retPtr unsafe.Pointer,
	extensionId js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_ShowOptions
//go:noescape
func TryShowOptions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_ShowPath
//go:noescape
func HasFuncShowPath() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_ShowPath
//go:noescape
func FuncShowPath(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_ShowPath
//go:noescape
func CallShowPath(
	retPtr unsafe.Pointer,
	extensionId js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_ShowPath
//go:noescape
func TryShowPath(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_UpdateExtensionCommand
//go:noescape
func HasFuncUpdateExtensionCommand() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_UpdateExtensionCommand
//go:noescape
func FuncUpdateExtensionCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_UpdateExtensionCommand
//go:noescape
func CallUpdateExtensionCommand(
	retPtr unsafe.Pointer,
	update unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_UpdateExtensionCommand
//go:noescape
func TryUpdateExtensionCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	update unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_UpdateExtensionConfiguration
//go:noescape
func HasFuncUpdateExtensionConfiguration() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_UpdateExtensionConfiguration
//go:noescape
func FuncUpdateExtensionConfiguration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_UpdateExtensionConfiguration
//go:noescape
func CallUpdateExtensionConfiguration(
	retPtr unsafe.Pointer,
	update unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_UpdateExtensionConfiguration
//go:noescape
func TryUpdateExtensionConfiguration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	update unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_UpdateProfileConfiguration
//go:noescape
func HasFuncUpdateProfileConfiguration() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_UpdateProfileConfiguration
//go:noescape
func FuncUpdateProfileConfiguration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_UpdateProfileConfiguration
//go:noescape
func CallUpdateProfileConfiguration(
	retPtr unsafe.Pointer,
	update unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate try_UpdateProfileConfiguration
//go:noescape
func TryUpdateProfileConfiguration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	update unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/developerprivate has_UpdateSiteAccess
//go:noescape
func HasFuncUpdateSiteAccess() js.Ref

//go:wasmimport plat/js/webext/developerprivate func_UpdateSiteAccess
//go:noescape
func FuncUpdateSiteAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/developerprivate call_UpdateSiteAccess
//go:noescape
func CallUpdateSiteAccess(
	retPtr unsafe.Pointer,
	site js.Ref,
	updates js.Ref)

//go:wasmimport plat/js/webext/developerprivate try_UpdateSiteAccess
//go:noescape
func TryUpdateSiteAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	site js.Ref,
	updates js.Ref) (ok js.Ref)
