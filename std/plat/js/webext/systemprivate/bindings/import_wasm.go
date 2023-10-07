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

//go:wasmimport plat/js/webext/systemprivate constof_GetIncognitoModeAvailabilityValue
//go:noescape
func ConstOfGetIncognitoModeAvailabilityValue(str js.Ref) uint32

//go:wasmimport plat/js/webext/systemprivate constof_UpdateStatusState
//go:noescape
func ConstOfUpdateStatusState(str js.Ref) uint32

//go:wasmimport plat/js/webext/systemprivate store_UpdateStatus
//go:noescape
func UpdateStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/systemprivate load_UpdateStatus
//go:noescape
func UpdateStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/systemprivate has_GetApiKey
//go:noescape
func HasFuncGetApiKey() js.Ref

//go:wasmimport plat/js/webext/systemprivate func_GetApiKey
//go:noescape
func FuncGetApiKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate call_GetApiKey
//go:noescape
func CallGetApiKey(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate try_GetApiKey
//go:noescape
func TryGetApiKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/systemprivate has_GetIncognitoModeAvailability
//go:noescape
func HasFuncGetIncognitoModeAvailability() js.Ref

//go:wasmimport plat/js/webext/systemprivate func_GetIncognitoModeAvailability
//go:noescape
func FuncGetIncognitoModeAvailability(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate call_GetIncognitoModeAvailability
//go:noescape
func CallGetIncognitoModeAvailability(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate try_GetIncognitoModeAvailability
//go:noescape
func TryGetIncognitoModeAvailability(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/systemprivate has_GetUpdateStatus
//go:noescape
func HasFuncGetUpdateStatus() js.Ref

//go:wasmimport plat/js/webext/systemprivate func_GetUpdateStatus
//go:noescape
func FuncGetUpdateStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate call_GetUpdateStatus
//go:noescape
func CallGetUpdateStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/systemprivate try_GetUpdateStatus
//go:noescape
func TryGetUpdateStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
