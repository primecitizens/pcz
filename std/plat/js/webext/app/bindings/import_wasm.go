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

//go:wasmimport plat/js/webext/app store_DOMWindow
//go:noescape
func DOMWindowJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app load_DOMWindow
//go:noescape
func DOMWindowJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app store_Details
//go:noescape
func DetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app load_Details
//go:noescape
func DetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app constof_InstallStateType
//go:noescape
func ConstOfInstallStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/app constof_RunningStateType
//go:noescape
func ConstOfRunningStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/app has_GetDetails
//go:noescape
func HasFuncGetDetails() js.Ref

//go:wasmimport plat/js/webext/app func_GetDetails
//go:noescape
func FuncGetDetails(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app call_GetDetails
//go:noescape
func CallGetDetails(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app try_GetDetails
//go:noescape
func TryGetDetails(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app has_GetIsInstalled
//go:noescape
func HasFuncGetIsInstalled() js.Ref

//go:wasmimport plat/js/webext/app func_GetIsInstalled
//go:noescape
func FuncGetIsInstalled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app call_GetIsInstalled
//go:noescape
func CallGetIsInstalled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app try_GetIsInstalled
//go:noescape
func TryGetIsInstalled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app has_InstallState
//go:noescape
func HasFuncInstallState() js.Ref

//go:wasmimport plat/js/webext/app func_InstallState
//go:noescape
func FuncInstallState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app call_InstallState
//go:noescape
func CallInstallState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app try_InstallState
//go:noescape
func TryInstallState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app has_RunningState
//go:noescape
func HasFuncRunningState() js.Ref

//go:wasmimport plat/js/webext/app func_RunningState
//go:noescape
func FuncRunningState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app call_RunningState
//go:noescape
func CallRunningState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app try_RunningState
//go:noescape
func TryRunningState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
