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

//go:wasmimport plat/js/webext/brailledisplayprivate store_DisplayState
//go:noescape
func DisplayStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate load_DisplayState
//go:noescape
func DisplayStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate constof_KeyCommand
//go:noescape
func ConstOfKeyCommand(str js.Ref) uint32

//go:wasmimport plat/js/webext/brailledisplayprivate store_KeyEvent
//go:noescape
func KeyEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate load_KeyEvent
//go:noescape
func KeyEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate has_GetDisplayState
//go:noescape
func HasFuncGetDisplayState() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_GetDisplayState
//go:noescape
func FuncGetDisplayState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_GetDisplayState
//go:noescape
func CallGetDisplayState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate try_GetDisplayState
//go:noescape
func TryGetDisplayState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_OnDisplayStateChanged
//go:noescape
func HasFuncOnDisplayStateChanged() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_OnDisplayStateChanged
//go:noescape
func FuncOnDisplayStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_OnDisplayStateChanged
//go:noescape
func CallOnDisplayStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_OnDisplayStateChanged
//go:noescape
func TryOnDisplayStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_OffDisplayStateChanged
//go:noescape
func HasFuncOffDisplayStateChanged() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_OffDisplayStateChanged
//go:noescape
func FuncOffDisplayStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_OffDisplayStateChanged
//go:noescape
func CallOffDisplayStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_OffDisplayStateChanged
//go:noescape
func TryOffDisplayStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_HasOnDisplayStateChanged
//go:noescape
func HasFuncHasOnDisplayStateChanged() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_HasOnDisplayStateChanged
//go:noescape
func FuncHasOnDisplayStateChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_HasOnDisplayStateChanged
//go:noescape
func CallHasOnDisplayStateChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_HasOnDisplayStateChanged
//go:noescape
func TryHasOnDisplayStateChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_OnKeyEvent
//go:noescape
func HasFuncOnKeyEvent() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_OnKeyEvent
//go:noescape
func FuncOnKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_OnKeyEvent
//go:noescape
func CallOnKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_OnKeyEvent
//go:noescape
func TryOnKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_OffKeyEvent
//go:noescape
func HasFuncOffKeyEvent() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_OffKeyEvent
//go:noescape
func FuncOffKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_OffKeyEvent
//go:noescape
func CallOffKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_OffKeyEvent
//go:noescape
func TryOffKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_HasOnKeyEvent
//go:noescape
func HasFuncHasOnKeyEvent() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_HasOnKeyEvent
//go:noescape
func FuncHasOnKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_HasOnKeyEvent
//go:noescape
func CallHasOnKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_HasOnKeyEvent
//go:noescape
func TryHasOnKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_UpdateBluetoothBrailleDisplayAddress
//go:noescape
func HasFuncUpdateBluetoothBrailleDisplayAddress() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_UpdateBluetoothBrailleDisplayAddress
//go:noescape
func FuncUpdateBluetoothBrailleDisplayAddress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_UpdateBluetoothBrailleDisplayAddress
//go:noescape
func CallUpdateBluetoothBrailleDisplayAddress(
	retPtr unsafe.Pointer,
	address js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate try_UpdateBluetoothBrailleDisplayAddress
//go:noescape
func TryUpdateBluetoothBrailleDisplayAddress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	address js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/brailledisplayprivate has_WriteDots
//go:noescape
func HasFuncWriteDots() js.Ref

//go:wasmimport plat/js/webext/brailledisplayprivate func_WriteDots
//go:noescape
func FuncWriteDots(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/brailledisplayprivate call_WriteDots
//go:noescape
func CallWriteDots(
	retPtr unsafe.Pointer,
	cells js.Ref,
	columns int32,
	rows int32)

//go:wasmimport plat/js/webext/brailledisplayprivate try_WriteDots
//go:noescape
func TryWriteDots(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cells js.Ref,
	columns int32,
	rows int32) (ok js.Ref)
