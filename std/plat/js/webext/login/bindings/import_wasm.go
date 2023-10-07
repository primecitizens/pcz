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

//go:wasmimport plat/js/webext/login store_SamlUserSessionProperties
//go:noescape
func SamlUserSessionPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/login load_SamlUserSessionProperties
//go:noescape
func SamlUserSessionPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/login has_EndSharedSession
//go:noescape
func HasFuncEndSharedSession() js.Ref

//go:wasmimport plat/js/webext/login func_EndSharedSession
//go:noescape
func FuncEndSharedSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_EndSharedSession
//go:noescape
func CallEndSharedSession(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_EndSharedSession
//go:noescape
func TryEndSharedSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_EnterSharedSession
//go:noescape
func HasFuncEnterSharedSession() js.Ref

//go:wasmimport plat/js/webext/login func_EnterSharedSession
//go:noescape
func FuncEnterSharedSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_EnterSharedSession
//go:noescape
func CallEnterSharedSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_EnterSharedSession
//go:noescape
func TryEnterSharedSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_ExitCurrentSession
//go:noescape
func HasFuncExitCurrentSession() js.Ref

//go:wasmimport plat/js/webext/login func_ExitCurrentSession
//go:noescape
func FuncExitCurrentSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_ExitCurrentSession
//go:noescape
func CallExitCurrentSession(
	retPtr unsafe.Pointer,
	dataForNextLoginAttempt js.Ref)

//go:wasmimport plat/js/webext/login try_ExitCurrentSession
//go:noescape
func TryExitCurrentSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataForNextLoginAttempt js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_FetchDataForNextLoginAttempt
//go:noescape
func HasFuncFetchDataForNextLoginAttempt() js.Ref

//go:wasmimport plat/js/webext/login func_FetchDataForNextLoginAttempt
//go:noescape
func FuncFetchDataForNextLoginAttempt(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_FetchDataForNextLoginAttempt
//go:noescape
func CallFetchDataForNextLoginAttempt(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_FetchDataForNextLoginAttempt
//go:noescape
func TryFetchDataForNextLoginAttempt(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_LaunchManagedGuestSession
//go:noescape
func HasFuncLaunchManagedGuestSession() js.Ref

//go:wasmimport plat/js/webext/login func_LaunchManagedGuestSession
//go:noescape
func FuncLaunchManagedGuestSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_LaunchManagedGuestSession
//go:noescape
func CallLaunchManagedGuestSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_LaunchManagedGuestSession
//go:noescape
func TryLaunchManagedGuestSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_LaunchSamlUserSession
//go:noescape
func HasFuncLaunchSamlUserSession() js.Ref

//go:wasmimport plat/js/webext/login func_LaunchSamlUserSession
//go:noescape
func FuncLaunchSamlUserSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_LaunchSamlUserSession
//go:noescape
func CallLaunchSamlUserSession(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_LaunchSamlUserSession
//go:noescape
func TryLaunchSamlUserSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_LaunchSharedManagedGuestSession
//go:noescape
func HasFuncLaunchSharedManagedGuestSession() js.Ref

//go:wasmimport plat/js/webext/login func_LaunchSharedManagedGuestSession
//go:noescape
func FuncLaunchSharedManagedGuestSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_LaunchSharedManagedGuestSession
//go:noescape
func CallLaunchSharedManagedGuestSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_LaunchSharedManagedGuestSession
//go:noescape
func TryLaunchSharedManagedGuestSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_LockCurrentSession
//go:noescape
func HasFuncLockCurrentSession() js.Ref

//go:wasmimport plat/js/webext/login func_LockCurrentSession
//go:noescape
func FuncLockCurrentSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_LockCurrentSession
//go:noescape
func CallLockCurrentSession(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_LockCurrentSession
//go:noescape
func TryLockCurrentSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_LockManagedGuestSession
//go:noescape
func HasFuncLockManagedGuestSession() js.Ref

//go:wasmimport plat/js/webext/login func_LockManagedGuestSession
//go:noescape
func FuncLockManagedGuestSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_LockManagedGuestSession
//go:noescape
func CallLockManagedGuestSession(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_LockManagedGuestSession
//go:noescape
func TryLockManagedGuestSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_NotifyExternalLogoutDone
//go:noescape
func HasFuncNotifyExternalLogoutDone() js.Ref

//go:wasmimport plat/js/webext/login func_NotifyExternalLogoutDone
//go:noescape
func FuncNotifyExternalLogoutDone(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_NotifyExternalLogoutDone
//go:noescape
func CallNotifyExternalLogoutDone(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_NotifyExternalLogoutDone
//go:noescape
func TryNotifyExternalLogoutDone(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_OnExternalLogoutDone
//go:noescape
func HasFuncOnExternalLogoutDone() js.Ref

//go:wasmimport plat/js/webext/login func_OnExternalLogoutDone
//go:noescape
func FuncOnExternalLogoutDone(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_OnExternalLogoutDone
//go:noescape
func CallOnExternalLogoutDone(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_OnExternalLogoutDone
//go:noescape
func TryOnExternalLogoutDone(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_OffExternalLogoutDone
//go:noescape
func HasFuncOffExternalLogoutDone() js.Ref

//go:wasmimport plat/js/webext/login func_OffExternalLogoutDone
//go:noescape
func FuncOffExternalLogoutDone(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_OffExternalLogoutDone
//go:noescape
func CallOffExternalLogoutDone(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_OffExternalLogoutDone
//go:noescape
func TryOffExternalLogoutDone(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_HasOnExternalLogoutDone
//go:noescape
func HasFuncHasOnExternalLogoutDone() js.Ref

//go:wasmimport plat/js/webext/login func_HasOnExternalLogoutDone
//go:noescape
func FuncHasOnExternalLogoutDone(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_HasOnExternalLogoutDone
//go:noescape
func CallHasOnExternalLogoutDone(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_HasOnExternalLogoutDone
//go:noescape
func TryHasOnExternalLogoutDone(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_OnRequestExternalLogout
//go:noescape
func HasFuncOnRequestExternalLogout() js.Ref

//go:wasmimport plat/js/webext/login func_OnRequestExternalLogout
//go:noescape
func FuncOnRequestExternalLogout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_OnRequestExternalLogout
//go:noescape
func CallOnRequestExternalLogout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_OnRequestExternalLogout
//go:noescape
func TryOnRequestExternalLogout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_OffRequestExternalLogout
//go:noescape
func HasFuncOffRequestExternalLogout() js.Ref

//go:wasmimport plat/js/webext/login func_OffRequestExternalLogout
//go:noescape
func FuncOffRequestExternalLogout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_OffRequestExternalLogout
//go:noescape
func CallOffRequestExternalLogout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_OffRequestExternalLogout
//go:noescape
func TryOffRequestExternalLogout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_HasOnRequestExternalLogout
//go:noescape
func HasFuncHasOnRequestExternalLogout() js.Ref

//go:wasmimport plat/js/webext/login func_HasOnRequestExternalLogout
//go:noescape
func FuncHasOnRequestExternalLogout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_HasOnRequestExternalLogout
//go:noescape
func CallHasOnRequestExternalLogout(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/login try_HasOnRequestExternalLogout
//go:noescape
func TryHasOnRequestExternalLogout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_RequestExternalLogout
//go:noescape
func HasFuncRequestExternalLogout() js.Ref

//go:wasmimport plat/js/webext/login func_RequestExternalLogout
//go:noescape
func FuncRequestExternalLogout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_RequestExternalLogout
//go:noescape
func CallRequestExternalLogout(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/login try_RequestExternalLogout
//go:noescape
func TryRequestExternalLogout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_SetDataForNextLoginAttempt
//go:noescape
func HasFuncSetDataForNextLoginAttempt() js.Ref

//go:wasmimport plat/js/webext/login func_SetDataForNextLoginAttempt
//go:noescape
func FuncSetDataForNextLoginAttempt(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_SetDataForNextLoginAttempt
//go:noescape
func CallSetDataForNextLoginAttempt(
	retPtr unsafe.Pointer,
	dataForNextLoginAttempt js.Ref)

//go:wasmimport plat/js/webext/login try_SetDataForNextLoginAttempt
//go:noescape
func TrySetDataForNextLoginAttempt(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataForNextLoginAttempt js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_UnlockCurrentSession
//go:noescape
func HasFuncUnlockCurrentSession() js.Ref

//go:wasmimport plat/js/webext/login func_UnlockCurrentSession
//go:noescape
func FuncUnlockCurrentSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_UnlockCurrentSession
//go:noescape
func CallUnlockCurrentSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_UnlockCurrentSession
//go:noescape
func TryUnlockCurrentSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_UnlockManagedGuestSession
//go:noescape
func HasFuncUnlockManagedGuestSession() js.Ref

//go:wasmimport plat/js/webext/login func_UnlockManagedGuestSession
//go:noescape
func FuncUnlockManagedGuestSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_UnlockManagedGuestSession
//go:noescape
func CallUnlockManagedGuestSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_UnlockManagedGuestSession
//go:noescape
func TryUnlockManagedGuestSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/login has_UnlockSharedSession
//go:noescape
func HasFuncUnlockSharedSession() js.Ref

//go:wasmimport plat/js/webext/login func_UnlockSharedSession
//go:noescape
func FuncUnlockSharedSession(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/login call_UnlockSharedSession
//go:noescape
func CallUnlockSharedSession(
	retPtr unsafe.Pointer,
	password js.Ref)

//go:wasmimport plat/js/webext/login try_UnlockSharedSession
//go:noescape
func TryUnlockSharedSession(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	password js.Ref) (ok js.Ref)
