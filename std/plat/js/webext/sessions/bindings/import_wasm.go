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

//go:wasmimport plat/js/webext/sessions store_Session
//go:noescape
func SessionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sessions load_Session
//go:noescape
func SessionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sessions store_Device
//go:noescape
func DeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sessions load_Device
//go:noescape
func DeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sessions store_Filter
//go:noescape
func FilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sessions load_Filter
//go:noescape
func FilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sessions get_MAX_SESSION_RESULTS
//go:noescape
func GetMAX_SESSION_RESULTS(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/sessions set_MAX_SESSION_RESULTS
//go:noescape
func SetMAX_SESSION_RESULTS(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/sessions has_GetDevices
//go:noescape
func HasFuncGetDevices() js.Ref

//go:wasmimport plat/js/webext/sessions func_GetDevices
//go:noescape
func FuncGetDevices(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_GetDevices
//go:noescape
func CallGetDevices(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions try_GetDevices
//go:noescape
func TryGetDevices(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sessions has_GetRecentlyClosed
//go:noescape
func HasFuncGetRecentlyClosed() js.Ref

//go:wasmimport plat/js/webext/sessions func_GetRecentlyClosed
//go:noescape
func FuncGetRecentlyClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_GetRecentlyClosed
//go:noescape
func CallGetRecentlyClosed(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions try_GetRecentlyClosed
//go:noescape
func TryGetRecentlyClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sessions has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/sessions func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sessions try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sessions has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/sessions func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sessions try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sessions has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/sessions func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/sessions try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sessions has_Restore
//go:noescape
func HasFuncRestore() js.Ref

//go:wasmimport plat/js/webext/sessions func_Restore
//go:noescape
func FuncRestore(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sessions call_Restore
//go:noescape
func CallRestore(
	retPtr unsafe.Pointer,
	sessionId js.Ref)

//go:wasmimport plat/js/webext/sessions try_Restore
//go:noescape
func TryRestore(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sessionId js.Ref) (ok js.Ref)
