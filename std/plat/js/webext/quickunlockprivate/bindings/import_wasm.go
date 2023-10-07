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

//go:wasmimport plat/js/webext/quickunlockprivate constof_CredentialProblem
//go:noescape
func ConstOfCredentialProblem(str js.Ref) uint32

//go:wasmimport plat/js/webext/quickunlockprivate store_CredentialCheck
//go:noescape
func CredentialCheckJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate load_CredentialCheck
//go:noescape
func CredentialCheckJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate store_CredentialRequirements
//go:noescape
func CredentialRequirementsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate load_CredentialRequirements
//go:noescape
func CredentialRequirementsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate constof_QuickUnlockMode
//go:noescape
func ConstOfQuickUnlockMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/quickunlockprivate store_TokenInfo
//go:noescape
func TokenInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate load_TokenInfo
//go:noescape
func TokenInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate has_CanAuthenticatePin
//go:noescape
func HasFuncCanAuthenticatePin() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_CanAuthenticatePin
//go:noescape
func FuncCanAuthenticatePin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_CanAuthenticatePin
//go:noescape
func CallCanAuthenticatePin(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate try_CanAuthenticatePin
//go:noescape
func TryCanAuthenticatePin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_CheckCredential
//go:noescape
func HasFuncCheckCredential() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_CheckCredential
//go:noescape
func FuncCheckCredential(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_CheckCredential
//go:noescape
func CallCheckCredential(
	retPtr unsafe.Pointer,
	mode uint32,
	credential js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_CheckCredential
//go:noescape
func TryCheckCredential(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32,
	credential js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_GetActiveModes
//go:noescape
func HasFuncGetActiveModes() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_GetActiveModes
//go:noescape
func FuncGetActiveModes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_GetActiveModes
//go:noescape
func CallGetActiveModes(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate try_GetActiveModes
//go:noescape
func TryGetActiveModes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_GetAuthToken
//go:noescape
func HasFuncGetAuthToken() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_GetAuthToken
//go:noescape
func FuncGetAuthToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_GetAuthToken
//go:noescape
func CallGetAuthToken(
	retPtr unsafe.Pointer,
	accountPassword js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_GetAuthToken
//go:noescape
func TryGetAuthToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	accountPassword js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_GetAvailableModes
//go:noescape
func HasFuncGetAvailableModes() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_GetAvailableModes
//go:noescape
func FuncGetAvailableModes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_GetAvailableModes
//go:noescape
func CallGetAvailableModes(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate try_GetAvailableModes
//go:noescape
func TryGetAvailableModes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_GetCredentialRequirements
//go:noescape
func HasFuncGetCredentialRequirements() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_GetCredentialRequirements
//go:noescape
func FuncGetCredentialRequirements(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_GetCredentialRequirements
//go:noescape
func CallGetCredentialRequirements(
	retPtr unsafe.Pointer,
	mode uint32)

//go:wasmimport plat/js/webext/quickunlockprivate try_GetCredentialRequirements
//go:noescape
func TryGetCredentialRequirements(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mode uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_OnActiveModesChanged
//go:noescape
func HasFuncOnActiveModesChanged() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_OnActiveModesChanged
//go:noescape
func FuncOnActiveModesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_OnActiveModesChanged
//go:noescape
func CallOnActiveModesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_OnActiveModesChanged
//go:noescape
func TryOnActiveModesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_OffActiveModesChanged
//go:noescape
func HasFuncOffActiveModesChanged() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_OffActiveModesChanged
//go:noescape
func FuncOffActiveModesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_OffActiveModesChanged
//go:noescape
func CallOffActiveModesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_OffActiveModesChanged
//go:noescape
func TryOffActiveModesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_HasOnActiveModesChanged
//go:noescape
func HasFuncHasOnActiveModesChanged() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_HasOnActiveModesChanged
//go:noescape
func FuncHasOnActiveModesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_HasOnActiveModesChanged
//go:noescape
func CallHasOnActiveModesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_HasOnActiveModesChanged
//go:noescape
func TryHasOnActiveModesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_SetLockScreenEnabled
//go:noescape
func HasFuncSetLockScreenEnabled() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_SetLockScreenEnabled
//go:noescape
func FuncSetLockScreenEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_SetLockScreenEnabled
//go:noescape
func CallSetLockScreenEnabled(
	retPtr unsafe.Pointer,
	token js.Ref,
	enabled js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_SetLockScreenEnabled
//go:noescape
func TrySetLockScreenEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	token js.Ref,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_SetModes
//go:noescape
func HasFuncSetModes() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_SetModes
//go:noescape
func FuncSetModes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_SetModes
//go:noescape
func CallSetModes(
	retPtr unsafe.Pointer,
	token js.Ref,
	modes js.Ref,
	credentials js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_SetModes
//go:noescape
func TrySetModes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	token js.Ref,
	modes js.Ref,
	credentials js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate has_SetPinAutosubmitEnabled
//go:noescape
func HasFuncSetPinAutosubmitEnabled() js.Ref

//go:wasmimport plat/js/webext/quickunlockprivate func_SetPinAutosubmitEnabled
//go:noescape
func FuncSetPinAutosubmitEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/quickunlockprivate call_SetPinAutosubmitEnabled
//go:noescape
func CallSetPinAutosubmitEnabled(
	retPtr unsafe.Pointer,
	token js.Ref,
	pin js.Ref,
	enabled js.Ref)

//go:wasmimport plat/js/webext/quickunlockprivate try_SetPinAutosubmitEnabled
//go:noescape
func TrySetPinAutosubmitEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	token js.Ref,
	pin js.Ref,
	enabled js.Ref) (ok js.Ref)
