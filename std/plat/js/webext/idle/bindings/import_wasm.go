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

//go:wasmimport plat/js/webext/idle constof_IdleState
//go:noescape
func ConstOfIdleState(str js.Ref) uint32

//go:wasmimport plat/js/webext/idle has_GetAutoLockDelay
//go:noescape
func HasFuncGetAutoLockDelay() js.Ref

//go:wasmimport plat/js/webext/idle func_GetAutoLockDelay
//go:noescape
func FuncGetAutoLockDelay(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_GetAutoLockDelay
//go:noescape
func CallGetAutoLockDelay(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/idle try_GetAutoLockDelay
//go:noescape
func TryGetAutoLockDelay(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/idle has_OnStateChanged
//go:noescape
func HasFuncOnStateChanged() js.Ref

//go:wasmimport plat/js/webext/idle func_OnStateChanged
//go:noescape
func FuncOnStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_OnStateChanged
//go:noescape
func CallOnStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/idle try_OnStateChanged
//go:noescape
func TryOnStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/idle has_OffStateChanged
//go:noescape
func HasFuncOffStateChanged() js.Ref

//go:wasmimport plat/js/webext/idle func_OffStateChanged
//go:noescape
func FuncOffStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_OffStateChanged
//go:noescape
func CallOffStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/idle try_OffStateChanged
//go:noescape
func TryOffStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/idle has_HasOnStateChanged
//go:noescape
func HasFuncHasOnStateChanged() js.Ref

//go:wasmimport plat/js/webext/idle func_HasOnStateChanged
//go:noescape
func FuncHasOnStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_HasOnStateChanged
//go:noescape
func CallHasOnStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/idle try_HasOnStateChanged
//go:noescape
func TryHasOnStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/idle has_QueryState
//go:noescape
func HasFuncQueryState() js.Ref

//go:wasmimport plat/js/webext/idle func_QueryState
//go:noescape
func FuncQueryState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_QueryState
//go:noescape
func CallQueryState(
	retPtr unsafe.Pointer,
	detectionIntervalInSeconds float64)

//go:wasmimport plat/js/webext/idle try_QueryState
//go:noescape
func TryQueryState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	detectionIntervalInSeconds float64) (ok js.Ref)

//go:wasmimport plat/js/webext/idle has_SetDetectionInterval
//go:noescape
func HasFuncSetDetectionInterval() js.Ref

//go:wasmimport plat/js/webext/idle func_SetDetectionInterval
//go:noescape
func FuncSetDetectionInterval(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idle call_SetDetectionInterval
//go:noescape
func CallSetDetectionInterval(
	retPtr unsafe.Pointer,
	intervalInSeconds float64)

//go:wasmimport plat/js/webext/idle try_SetDetectionInterval
//go:noescape
func TrySetDetectionInterval(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	intervalInSeconds float64) (ok js.Ref)
