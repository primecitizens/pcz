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

//go:wasmimport plat/js/webext/types constof_LevelOfControl
//go:noescape
func ConstOfLevelOfControl(str js.Ref) uint32

//go:wasmimport plat/js/webext/types store_GetReturnType
//go:noescape
func GetReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/types load_GetReturnType
//go:noescape
func GetReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/types store_GetArgDetails
//go:noescape
func GetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/types load_GetArgDetails
//go:noescape
func GetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/types constof_ChromeSettingScope
//go:noescape
func ConstOfChromeSettingScope(str js.Ref) uint32

//go:wasmimport plat/js/webext/types store_SetArgDetails
//go:noescape
func SetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/types load_SetArgDetails
//go:noescape
func SetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/types store_ClearArgDetails
//go:noescape
func ClearArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/types load_ClearArgDetails
//go:noescape
func ClearArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/types has_ChromeSetting_Get
//go:noescape
func HasFuncChromeSettingGet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/types func_ChromeSetting_Get
//go:noescape
func FuncChromeSettingGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/types call_ChromeSetting_Get
//go:noescape
func CallChromeSettingGet(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/types try_ChromeSetting_Get
//go:noescape
func TryChromeSettingGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/types has_ChromeSetting_Set
//go:noescape
func HasFuncChromeSettingSet(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/types func_ChromeSetting_Set
//go:noescape
func FuncChromeSettingSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/types call_ChromeSetting_Set
//go:noescape
func CallChromeSettingSet(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/types try_ChromeSetting_Set
//go:noescape
func TryChromeSettingSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/types has_ChromeSetting_Clear
//go:noescape
func HasFuncChromeSettingClear(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/types func_ChromeSetting_Clear
//go:noescape
func FuncChromeSettingClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/types call_ChromeSetting_Clear
//go:noescape
func CallChromeSettingClear(
	this js.Ref, retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/types try_ChromeSetting_Clear
//go:noescape
func TryChromeSettingClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
