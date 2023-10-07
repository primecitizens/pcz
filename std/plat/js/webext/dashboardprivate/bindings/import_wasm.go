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

//go:wasmimport plat/js/webext/dashboardprivate constof_Result
//go:noescape
func ConstOfResult(str js.Ref) uint32

//go:wasmimport plat/js/webext/dashboardprivate store_ShowPermissionPromptForDelegatedInstallArgDetails
//go:noescape
func ShowPermissionPromptForDelegatedInstallArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/dashboardprivate load_ShowPermissionPromptForDelegatedInstallArgDetails
//go:noescape
func ShowPermissionPromptForDelegatedInstallArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/dashboardprivate has_ShowPermissionPromptForDelegatedInstall
//go:noescape
func HasFuncShowPermissionPromptForDelegatedInstall() js.Ref

//go:wasmimport plat/js/webext/dashboardprivate func_ShowPermissionPromptForDelegatedInstall
//go:noescape
func FuncShowPermissionPromptForDelegatedInstall(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/dashboardprivate call_ShowPermissionPromptForDelegatedInstall
//go:noescape
func CallShowPermissionPromptForDelegatedInstall(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/dashboardprivate try_ShowPermissionPromptForDelegatedInstall
//go:noescape
func TryShowPermissionPromptForDelegatedInstall(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
