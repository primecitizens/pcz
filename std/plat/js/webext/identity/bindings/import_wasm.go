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

//go:wasmimport plat/js/webext/identity store_AccountInfo
//go:noescape
func AccountInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_AccountInfo
//go:noescape
func AccountInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity constof_AccountStatus
//go:noescape
func ConstOfAccountStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/identity store_GetAuthTokenResult
//go:noescape
func GetAuthTokenResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_GetAuthTokenResult
//go:noescape
func GetAuthTokenResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity store_ProfileUserInfo
//go:noescape
func ProfileUserInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_ProfileUserInfo
//go:noescape
func ProfileUserInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity store_InvalidTokenDetails
//go:noescape
func InvalidTokenDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_InvalidTokenDetails
//go:noescape
func InvalidTokenDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity store_ProfileDetails
//go:noescape
func ProfileDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_ProfileDetails
//go:noescape
func ProfileDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity store_TokenDetails
//go:noescape
func TokenDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_TokenDetails
//go:noescape
func TokenDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity store_WebAuthFlowDetails
//go:noescape
func WebAuthFlowDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/identity load_WebAuthFlowDetails
//go:noescape
func WebAuthFlowDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/identity has_ClearAllCachedAuthTokens
//go:noescape
func HasFuncClearAllCachedAuthTokens() js.Ref

//go:wasmimport plat/js/webext/identity func_ClearAllCachedAuthTokens
//go:noescape
func FuncClearAllCachedAuthTokens(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_ClearAllCachedAuthTokens
//go:noescape
func CallClearAllCachedAuthTokens(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_ClearAllCachedAuthTokens
//go:noescape
func TryClearAllCachedAuthTokens(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_GetAccounts
//go:noescape
func HasFuncGetAccounts() js.Ref

//go:wasmimport plat/js/webext/identity func_GetAccounts
//go:noescape
func FuncGetAccounts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_GetAccounts
//go:noescape
func CallGetAccounts(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_GetAccounts
//go:noescape
func TryGetAccounts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_GetAuthToken
//go:noescape
func HasFuncGetAuthToken() js.Ref

//go:wasmimport plat/js/webext/identity func_GetAuthToken
//go:noescape
func FuncGetAuthToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_GetAuthToken
//go:noescape
func CallGetAuthToken(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_GetAuthToken
//go:noescape
func TryGetAuthToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_GetProfileUserInfo
//go:noescape
func HasFuncGetProfileUserInfo() js.Ref

//go:wasmimport plat/js/webext/identity func_GetProfileUserInfo
//go:noescape
func FuncGetProfileUserInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_GetProfileUserInfo
//go:noescape
func CallGetProfileUserInfo(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_GetProfileUserInfo
//go:noescape
func TryGetProfileUserInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_GetRedirectURL
//go:noescape
func HasFuncGetRedirectURL() js.Ref

//go:wasmimport plat/js/webext/identity func_GetRedirectURL
//go:noescape
func FuncGetRedirectURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_GetRedirectURL
//go:noescape
func CallGetRedirectURL(
	retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/webext/identity try_GetRedirectURL
//go:noescape
func TryGetRedirectURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_LaunchWebAuthFlow
//go:noescape
func HasFuncLaunchWebAuthFlow() js.Ref

//go:wasmimport plat/js/webext/identity func_LaunchWebAuthFlow
//go:noescape
func FuncLaunchWebAuthFlow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_LaunchWebAuthFlow
//go:noescape
func CallLaunchWebAuthFlow(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_LaunchWebAuthFlow
//go:noescape
func TryLaunchWebAuthFlow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_OnSignInChanged
//go:noescape
func HasFuncOnSignInChanged() js.Ref

//go:wasmimport plat/js/webext/identity func_OnSignInChanged
//go:noescape
func FuncOnSignInChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_OnSignInChanged
//go:noescape
func CallOnSignInChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/identity try_OnSignInChanged
//go:noescape
func TryOnSignInChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_OffSignInChanged
//go:noescape
func HasFuncOffSignInChanged() js.Ref

//go:wasmimport plat/js/webext/identity func_OffSignInChanged
//go:noescape
func FuncOffSignInChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_OffSignInChanged
//go:noescape
func CallOffSignInChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/identity try_OffSignInChanged
//go:noescape
func TryOffSignInChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_HasOnSignInChanged
//go:noescape
func HasFuncHasOnSignInChanged() js.Ref

//go:wasmimport plat/js/webext/identity func_HasOnSignInChanged
//go:noescape
func FuncHasOnSignInChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_HasOnSignInChanged
//go:noescape
func CallHasOnSignInChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/identity try_HasOnSignInChanged
//go:noescape
func TryHasOnSignInChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/identity has_RemoveCachedAuthToken
//go:noescape
func HasFuncRemoveCachedAuthToken() js.Ref

//go:wasmimport plat/js/webext/identity func_RemoveCachedAuthToken
//go:noescape
func FuncRemoveCachedAuthToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/identity call_RemoveCachedAuthToken
//go:noescape
func CallRemoveCachedAuthToken(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/identity try_RemoveCachedAuthToken
//go:noescape
func TryRemoveCachedAuthToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
