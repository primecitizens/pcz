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

//go:wasmimport plat/js/webext/systemindicator store_SetIconDetails
//go:noescape
func SetIconDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/systemindicator load_SetIconDetails
//go:noescape
func SetIconDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/systemindicator has_Disable
//go:noescape
func HasFuncDisable() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_Disable
//go:noescape
func FuncDisable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_Disable
//go:noescape
func CallDisable(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator try_Disable
//go:noescape
func TryDisable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/systemindicator has_Enable
//go:noescape
func HasFuncEnable() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_Enable
//go:noescape
func FuncEnable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_Enable
//go:noescape
func CallEnable(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator try_Enable
//go:noescape
func TryEnable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/systemindicator has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/systemindicator try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/systemindicator has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/systemindicator try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/systemindicator has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/systemindicator try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/systemindicator has_SetIcon
//go:noescape
func HasFuncSetIcon() js.Ref

//go:wasmimport plat/js/webext/systemindicator func_SetIcon
//go:noescape
func FuncSetIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator call_SetIcon
//go:noescape
func CallSetIcon(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/systemindicator try_SetIcon
//go:noescape
func TrySetIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
