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

//go:wasmimport plat/js/webext/cecprivate constof_DisplayCecPowerState
//go:noescape
func ConstOfDisplayCecPowerState(str js.Ref) uint32

//go:wasmimport plat/js/webext/cecprivate has_QueryDisplayCecPowerState
//go:noescape
func HasFuncQueryDisplayCecPowerState() js.Ref

//go:wasmimport plat/js/webext/cecprivate func_QueryDisplayCecPowerState
//go:noescape
func FuncQueryDisplayCecPowerState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate call_QueryDisplayCecPowerState
//go:noescape
func CallQueryDisplayCecPowerState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate try_QueryDisplayCecPowerState
//go:noescape
func TryQueryDisplayCecPowerState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cecprivate has_SendStandBy
//go:noescape
func HasFuncSendStandBy() js.Ref

//go:wasmimport plat/js/webext/cecprivate func_SendStandBy
//go:noescape
func FuncSendStandBy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate call_SendStandBy
//go:noescape
func CallSendStandBy(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate try_SendStandBy
//go:noescape
func TrySendStandBy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/cecprivate has_SendWakeUp
//go:noescape
func HasFuncSendWakeUp() js.Ref

//go:wasmimport plat/js/webext/cecprivate func_SendWakeUp
//go:noescape
func FuncSendWakeUp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate call_SendWakeUp
//go:noescape
func CallSendWakeUp(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/cecprivate try_SendWakeUp
//go:noescape
func TrySendWakeUp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
