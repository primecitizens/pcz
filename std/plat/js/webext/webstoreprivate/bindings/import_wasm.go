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

//go:wasmimport plat/js/webext/webstoreprivate constof_ExtensionInstallStatus
//go:noescape
func ConstOfExtensionInstallStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/webstoreprivate store_GetBrowserLoginReturnType
//go:noescape
func GetBrowserLoginReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate load_GetBrowserLoginReturnType
//go:noescape
func GetBrowserLoginReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webstoreprivate constof_Result
//go:noescape
func ConstOfResult(str js.Ref) uint32

//go:wasmimport plat/js/webext/webstoreprivate constof_WebGlStatus
//go:noescape
func ConstOfWebGlStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/webstoreprivate has_BeginInstallWithManifest3
//go:noescape
func HasFuncBeginInstallWithManifest3() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_BeginInstallWithManifest3
//go:noescape
func FuncBeginInstallWithManifest3(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_BeginInstallWithManifest3
//go:noescape
func CallBeginInstallWithManifest3(
	retPtr unsafe.Pointer,
	details js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_BeginInstallWithManifest3
//go:noescape
func TryBeginInstallWithManifest3(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_CompleteInstall
//go:noescape
func HasFuncCompleteInstall() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_CompleteInstall
//go:noescape
func FuncCompleteInstall(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_CompleteInstall
//go:noescape
func CallCompleteInstall(
	retPtr unsafe.Pointer,
	expected_id js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_CompleteInstall
//go:noescape
func TryCompleteInstall(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expected_id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_EnableAppLauncher
//go:noescape
func HasFuncEnableAppLauncher() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_EnableAppLauncher
//go:noescape
func FuncEnableAppLauncher(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_EnableAppLauncher
//go:noescape
func CallEnableAppLauncher(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_EnableAppLauncher
//go:noescape
func TryEnableAppLauncher(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetBrowserLogin
//go:noescape
func HasFuncGetBrowserLogin() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetBrowserLogin
//go:noescape
func FuncGetBrowserLogin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetBrowserLogin
//go:noescape
func CallGetBrowserLogin(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_GetBrowserLogin
//go:noescape
func TryGetBrowserLogin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetExtensionStatus
//go:noescape
func HasFuncGetExtensionStatus() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetExtensionStatus
//go:noescape
func FuncGetExtensionStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetExtensionStatus
//go:noescape
func CallGetExtensionStatus(
	retPtr unsafe.Pointer,
	id js.Ref,
	manifest js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_GetExtensionStatus
//go:noescape
func TryGetExtensionStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	manifest js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetIsLauncherEnabled
//go:noescape
func HasFuncGetIsLauncherEnabled() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetIsLauncherEnabled
//go:noescape
func FuncGetIsLauncherEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetIsLauncherEnabled
//go:noescape
func CallGetIsLauncherEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_GetIsLauncherEnabled
//go:noescape
func TryGetIsLauncherEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetReferrerChain
//go:noescape
func HasFuncGetReferrerChain() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetReferrerChain
//go:noescape
func FuncGetReferrerChain(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetReferrerChain
//go:noescape
func CallGetReferrerChain(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_GetReferrerChain
//go:noescape
func TryGetReferrerChain(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetStoreLogin
//go:noescape
func HasFuncGetStoreLogin() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetStoreLogin
//go:noescape
func FuncGetStoreLogin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetStoreLogin
//go:noescape
func CallGetStoreLogin(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_GetStoreLogin
//go:noescape
func TryGetStoreLogin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_GetWebGLStatus
//go:noescape
func HasFuncGetWebGLStatus() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_GetWebGLStatus
//go:noescape
func FuncGetWebGLStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_GetWebGLStatus
//go:noescape
func CallGetWebGLStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_GetWebGLStatus
//go:noescape
func TryGetWebGLStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_Install
//go:noescape
func HasFuncInstall() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_Install
//go:noescape
func FuncInstall(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_Install
//go:noescape
func CallInstall(
	retPtr unsafe.Pointer,
	expected_id js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_Install
//go:noescape
func TryInstall(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expected_id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_IsInIncognitoMode
//go:noescape
func HasFuncIsInIncognitoMode() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_IsInIncognitoMode
//go:noescape
func FuncIsInIncognitoMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_IsInIncognitoMode
//go:noescape
func CallIsInIncognitoMode(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate try_IsInIncognitoMode
//go:noescape
func TryIsInIncognitoMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_IsPendingCustodianApproval
//go:noescape
func HasFuncIsPendingCustodianApproval() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_IsPendingCustodianApproval
//go:noescape
func FuncIsPendingCustodianApproval(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_IsPendingCustodianApproval
//go:noescape
func CallIsPendingCustodianApproval(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_IsPendingCustodianApproval
//go:noescape
func TryIsPendingCustodianApproval(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate has_SetStoreLogin
//go:noescape
func HasFuncSetStoreLogin() js.Ref

//go:wasmimport plat/js/webext/webstoreprivate func_SetStoreLogin
//go:noescape
func FuncSetStoreLogin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webstoreprivate call_SetStoreLogin
//go:noescape
func CallSetStoreLogin(
	retPtr unsafe.Pointer,
	login js.Ref)

//go:wasmimport plat/js/webext/webstoreprivate try_SetStoreLogin
//go:noescape
func TrySetStoreLogin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	login js.Ref) (ok js.Ref)
