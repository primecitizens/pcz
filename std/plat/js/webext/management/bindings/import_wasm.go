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

//go:wasmimport plat/js/webext/management constof_ExtensionDisabledReason
//go:noescape
func ConstOfExtensionDisabledReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/management constof_LaunchType
//go:noescape
func ConstOfLaunchType(str js.Ref) uint32

//go:wasmimport plat/js/webext/management store_IconInfo
//go:noescape
func IconInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/management load_IconInfo
//go:noescape
func IconInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/management constof_ExtensionInstallType
//go:noescape
func ConstOfExtensionInstallType(str js.Ref) uint32

//go:wasmimport plat/js/webext/management constof_ExtensionType
//go:noescape
func ConstOfExtensionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/management store_ExtensionInfo
//go:noescape
func ExtensionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/management load_ExtensionInfo
//go:noescape
func ExtensionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/management store_UninstallOptions
//go:noescape
func UninstallOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/management load_UninstallOptions
//go:noescape
func UninstallOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/management has_CreateAppShortcut
//go:noescape
func HasFuncCreateAppShortcut() js.Ref

//go:wasmimport plat/js/webext/management func_CreateAppShortcut
//go:noescape
func FuncCreateAppShortcut(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_CreateAppShortcut
//go:noescape
func CallCreateAppShortcut(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/management try_CreateAppShortcut
//go:noescape
func TryCreateAppShortcut(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_GenerateAppForLink
//go:noescape
func HasFuncGenerateAppForLink() js.Ref

//go:wasmimport plat/js/webext/management func_GenerateAppForLink
//go:noescape
func FuncGenerateAppForLink(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_GenerateAppForLink
//go:noescape
func CallGenerateAppForLink(
	retPtr unsafe.Pointer,
	url js.Ref,
	title js.Ref)

//go:wasmimport plat/js/webext/management try_GenerateAppForLink
//go:noescape
func TryGenerateAppForLink(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	title js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/management func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/management try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/management func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/management try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_GetPermissionWarningsById
//go:noescape
func HasFuncGetPermissionWarningsById() js.Ref

//go:wasmimport plat/js/webext/management func_GetPermissionWarningsById
//go:noescape
func FuncGetPermissionWarningsById(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_GetPermissionWarningsById
//go:noescape
func CallGetPermissionWarningsById(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/management try_GetPermissionWarningsById
//go:noescape
func TryGetPermissionWarningsById(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_GetPermissionWarningsByManifest
//go:noescape
func HasFuncGetPermissionWarningsByManifest() js.Ref

//go:wasmimport plat/js/webext/management func_GetPermissionWarningsByManifest
//go:noescape
func FuncGetPermissionWarningsByManifest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_GetPermissionWarningsByManifest
//go:noescape
func CallGetPermissionWarningsByManifest(
	retPtr unsafe.Pointer,
	manifestStr js.Ref)

//go:wasmimport plat/js/webext/management try_GetPermissionWarningsByManifest
//go:noescape
func TryGetPermissionWarningsByManifest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	manifestStr js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_GetSelf
//go:noescape
func HasFuncGetSelf() js.Ref

//go:wasmimport plat/js/webext/management func_GetSelf
//go:noescape
func FuncGetSelf(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_GetSelf
//go:noescape
func CallGetSelf(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/management try_GetSelf
//go:noescape
func TryGetSelf(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_InstallReplacementWebApp
//go:noescape
func HasFuncInstallReplacementWebApp() js.Ref

//go:wasmimport plat/js/webext/management func_InstallReplacementWebApp
//go:noescape
func FuncInstallReplacementWebApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_InstallReplacementWebApp
//go:noescape
func CallInstallReplacementWebApp(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/management try_InstallReplacementWebApp
//go:noescape
func TryInstallReplacementWebApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_LaunchApp
//go:noescape
func HasFuncLaunchApp() js.Ref

//go:wasmimport plat/js/webext/management func_LaunchApp
//go:noescape
func FuncLaunchApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_LaunchApp
//go:noescape
func CallLaunchApp(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/management try_LaunchApp
//go:noescape
func TryLaunchApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OnDisabled
//go:noescape
func HasFuncOnDisabled() js.Ref

//go:wasmimport plat/js/webext/management func_OnDisabled
//go:noescape
func FuncOnDisabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OnDisabled
//go:noescape
func CallOnDisabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OnDisabled
//go:noescape
func TryOnDisabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OffDisabled
//go:noescape
func HasFuncOffDisabled() js.Ref

//go:wasmimport plat/js/webext/management func_OffDisabled
//go:noescape
func FuncOffDisabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OffDisabled
//go:noescape
func CallOffDisabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OffDisabled
//go:noescape
func TryOffDisabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_HasOnDisabled
//go:noescape
func HasFuncHasOnDisabled() js.Ref

//go:wasmimport plat/js/webext/management func_HasOnDisabled
//go:noescape
func FuncHasOnDisabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_HasOnDisabled
//go:noescape
func CallHasOnDisabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_HasOnDisabled
//go:noescape
func TryHasOnDisabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OnEnabled
//go:noescape
func HasFuncOnEnabled() js.Ref

//go:wasmimport plat/js/webext/management func_OnEnabled
//go:noescape
func FuncOnEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OnEnabled
//go:noescape
func CallOnEnabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OnEnabled
//go:noescape
func TryOnEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OffEnabled
//go:noescape
func HasFuncOffEnabled() js.Ref

//go:wasmimport plat/js/webext/management func_OffEnabled
//go:noescape
func FuncOffEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OffEnabled
//go:noescape
func CallOffEnabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OffEnabled
//go:noescape
func TryOffEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_HasOnEnabled
//go:noescape
func HasFuncHasOnEnabled() js.Ref

//go:wasmimport plat/js/webext/management func_HasOnEnabled
//go:noescape
func FuncHasOnEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_HasOnEnabled
//go:noescape
func CallHasOnEnabled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_HasOnEnabled
//go:noescape
func TryHasOnEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OnInstalled
//go:noescape
func HasFuncOnInstalled() js.Ref

//go:wasmimport plat/js/webext/management func_OnInstalled
//go:noescape
func FuncOnInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OnInstalled
//go:noescape
func CallOnInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OnInstalled
//go:noescape
func TryOnInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OffInstalled
//go:noescape
func HasFuncOffInstalled() js.Ref

//go:wasmimport plat/js/webext/management func_OffInstalled
//go:noescape
func FuncOffInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OffInstalled
//go:noescape
func CallOffInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OffInstalled
//go:noescape
func TryOffInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_HasOnInstalled
//go:noescape
func HasFuncHasOnInstalled() js.Ref

//go:wasmimport plat/js/webext/management func_HasOnInstalled
//go:noescape
func FuncHasOnInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_HasOnInstalled
//go:noescape
func CallHasOnInstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_HasOnInstalled
//go:noescape
func TryHasOnInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OnUninstalled
//go:noescape
func HasFuncOnUninstalled() js.Ref

//go:wasmimport plat/js/webext/management func_OnUninstalled
//go:noescape
func FuncOnUninstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OnUninstalled
//go:noescape
func CallOnUninstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OnUninstalled
//go:noescape
func TryOnUninstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_OffUninstalled
//go:noescape
func HasFuncOffUninstalled() js.Ref

//go:wasmimport plat/js/webext/management func_OffUninstalled
//go:noescape
func FuncOffUninstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_OffUninstalled
//go:noescape
func CallOffUninstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_OffUninstalled
//go:noescape
func TryOffUninstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_HasOnUninstalled
//go:noescape
func HasFuncHasOnUninstalled() js.Ref

//go:wasmimport plat/js/webext/management func_HasOnUninstalled
//go:noescape
func FuncHasOnUninstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_HasOnUninstalled
//go:noescape
func CallHasOnUninstalled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/management try_HasOnUninstalled
//go:noescape
func TryHasOnUninstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_SetEnabled
//go:noescape
func HasFuncSetEnabled() js.Ref

//go:wasmimport plat/js/webext/management func_SetEnabled
//go:noescape
func FuncSetEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_SetEnabled
//go:noescape
func CallSetEnabled(
	retPtr unsafe.Pointer,
	id js.Ref,
	enabled js.Ref)

//go:wasmimport plat/js/webext/management try_SetEnabled
//go:noescape
func TrySetEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_SetLaunchType
//go:noescape
func HasFuncSetLaunchType() js.Ref

//go:wasmimport plat/js/webext/management func_SetLaunchType
//go:noescape
func FuncSetLaunchType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_SetLaunchType
//go:noescape
func CallSetLaunchType(
	retPtr unsafe.Pointer,
	id js.Ref,
	launchType uint32)

//go:wasmimport plat/js/webext/management try_SetLaunchType
//go:noescape
func TrySetLaunchType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	launchType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_Uninstall
//go:noescape
func HasFuncUninstall() js.Ref

//go:wasmimport plat/js/webext/management func_Uninstall
//go:noescape
func FuncUninstall(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_Uninstall
//go:noescape
func CallUninstall(
	retPtr unsafe.Pointer,
	id js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/management try_Uninstall
//go:noescape
func TryUninstall(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/management has_UninstallSelf
//go:noescape
func HasFuncUninstallSelf() js.Ref

//go:wasmimport plat/js/webext/management func_UninstallSelf
//go:noescape
func FuncUninstallSelf(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/management call_UninstallSelf
//go:noescape
func CallUninstallSelf(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/management try_UninstallSelf
//go:noescape
func TryUninstallSelf(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
