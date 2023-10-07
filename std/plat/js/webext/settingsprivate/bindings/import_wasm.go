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

//go:wasmimport plat/js/webext/settingsprivate constof_ControlledBy
//go:noescape
func ConstOfControlledBy(str js.Ref) uint32

//go:wasmimport plat/js/webext/settingsprivate constof_Enforcement
//go:noescape
func ConstOfEnforcement(str js.Ref) uint32

//go:wasmimport plat/js/webext/settingsprivate constof_PrefType
//go:noescape
func ConstOfPrefType(str js.Ref) uint32

//go:wasmimport plat/js/webext/settingsprivate store_PrefObject
//go:noescape
func PrefObjectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/settingsprivate load_PrefObject
//go:noescape
func PrefObjectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/settingsprivate has_GetAllPrefs
//go:noescape
func HasFuncGetAllPrefs() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_GetAllPrefs
//go:noescape
func FuncGetAllPrefs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_GetAllPrefs
//go:noescape
func CallGetAllPrefs(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate try_GetAllPrefs
//go:noescape
func TryGetAllPrefs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_GetDefaultZoom
//go:noescape
func HasFuncGetDefaultZoom() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_GetDefaultZoom
//go:noescape
func FuncGetDefaultZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_GetDefaultZoom
//go:noescape
func CallGetDefaultZoom(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate try_GetDefaultZoom
//go:noescape
func TryGetDefaultZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_GetPref
//go:noescape
func HasFuncGetPref() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_GetPref
//go:noescape
func FuncGetPref(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_GetPref
//go:noescape
func CallGetPref(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/settingsprivate try_GetPref
//go:noescape
func TryGetPref(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_OnPrefsChanged
//go:noescape
func HasFuncOnPrefsChanged() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_OnPrefsChanged
//go:noescape
func FuncOnPrefsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_OnPrefsChanged
//go:noescape
func CallOnPrefsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/settingsprivate try_OnPrefsChanged
//go:noescape
func TryOnPrefsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_OffPrefsChanged
//go:noescape
func HasFuncOffPrefsChanged() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_OffPrefsChanged
//go:noescape
func FuncOffPrefsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_OffPrefsChanged
//go:noescape
func CallOffPrefsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/settingsprivate try_OffPrefsChanged
//go:noescape
func TryOffPrefsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_HasOnPrefsChanged
//go:noescape
func HasFuncHasOnPrefsChanged() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_HasOnPrefsChanged
//go:noescape
func FuncHasOnPrefsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_HasOnPrefsChanged
//go:noescape
func CallHasOnPrefsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/settingsprivate try_HasOnPrefsChanged
//go:noescape
func TryHasOnPrefsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_SetDefaultZoom
//go:noescape
func HasFuncSetDefaultZoom() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_SetDefaultZoom
//go:noescape
func FuncSetDefaultZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_SetDefaultZoom
//go:noescape
func CallSetDefaultZoom(
	retPtr unsafe.Pointer,
	zoom float64)

//go:wasmimport plat/js/webext/settingsprivate try_SetDefaultZoom
//go:noescape
func TrySetDefaultZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	zoom float64) (ok js.Ref)

//go:wasmimport plat/js/webext/settingsprivate has_SetPref
//go:noescape
func HasFuncSetPref() js.Ref

//go:wasmimport plat/js/webext/settingsprivate func_SetPref
//go:noescape
func FuncSetPref(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/settingsprivate call_SetPref
//go:noescape
func CallSetPref(
	retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref,
	pageId js.Ref)

//go:wasmimport plat/js/webext/settingsprivate try_SetPref
//go:noescape
func TrySetPref(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref,
	pageId js.Ref) (ok js.Ref)
