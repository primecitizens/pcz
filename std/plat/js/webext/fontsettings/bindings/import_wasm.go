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

//go:wasmimport plat/js/webext/fontsettings store_ClearDefaultFixedFontSizeArgDetails
//go:noescape
func ClearDefaultFixedFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_ClearDefaultFixedFontSizeArgDetails
//go:noescape
func ClearDefaultFixedFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_ClearDefaultFontSizeArgDetails
//go:noescape
func ClearDefaultFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_ClearDefaultFontSizeArgDetails
//go:noescape
func ClearDefaultFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings constof_GenericFamily
//go:noescape
func ConstOfGenericFamily(str js.Ref) uint32

//go:wasmimport plat/js/webext/fontsettings constof_ScriptCode
//go:noescape
func ConstOfScriptCode(str js.Ref) uint32

//go:wasmimport plat/js/webext/fontsettings store_ClearFontArgDetails
//go:noescape
func ClearFontArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_ClearFontArgDetails
//go:noescape
func ClearFontArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_ClearMinimumFontSizeArgDetails
//go:noescape
func ClearMinimumFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_ClearMinimumFontSizeArgDetails
//go:noescape
func ClearMinimumFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_FontName
//go:noescape
func FontNameJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_FontName
//go:noescape
func FontNameJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetDefaultFixedFontSizeArgDetails
//go:noescape
func GetDefaultFixedFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetDefaultFixedFontSizeArgDetails
//go:noescape
func GetDefaultFixedFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings constof_LevelOfControl
//go:noescape
func ConstOfLevelOfControl(str js.Ref) uint32

//go:wasmimport plat/js/webext/fontsettings store_GetDefaultFixedFontSizeReturnType
//go:noescape
func GetDefaultFixedFontSizeReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetDefaultFixedFontSizeReturnType
//go:noescape
func GetDefaultFixedFontSizeReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetDefaultFontSizeArgDetails
//go:noescape
func GetDefaultFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetDefaultFontSizeArgDetails
//go:noescape
func GetDefaultFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetDefaultFontSizeReturnType
//go:noescape
func GetDefaultFontSizeReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetDefaultFontSizeReturnType
//go:noescape
func GetDefaultFontSizeReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetFontArgDetails
//go:noescape
func GetFontArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetFontArgDetails
//go:noescape
func GetFontArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetFontReturnType
//go:noescape
func GetFontReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetFontReturnType
//go:noescape
func GetFontReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetMinimumFontSizeArgDetails
//go:noescape
func GetMinimumFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetMinimumFontSizeArgDetails
//go:noescape
func GetMinimumFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_GetMinimumFontSizeReturnType
//go:noescape
func GetMinimumFontSizeReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_GetMinimumFontSizeReturnType
//go:noescape
func GetMinimumFontSizeReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_OnDefaultFixedFontSizeChangedArgDetails
//go:noescape
func OnDefaultFixedFontSizeChangedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_OnDefaultFixedFontSizeChangedArgDetails
//go:noescape
func OnDefaultFixedFontSizeChangedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_OnDefaultFontSizeChangedArgDetails
//go:noescape
func OnDefaultFontSizeChangedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_OnDefaultFontSizeChangedArgDetails
//go:noescape
func OnDefaultFontSizeChangedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_OnFontChangedArgDetails
//go:noescape
func OnFontChangedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_OnFontChangedArgDetails
//go:noescape
func OnFontChangedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_OnMinimumFontSizeChangedArgDetails
//go:noescape
func OnMinimumFontSizeChangedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_OnMinimumFontSizeChangedArgDetails
//go:noescape
func OnMinimumFontSizeChangedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_SetDefaultFixedFontSizeArgDetails
//go:noescape
func SetDefaultFixedFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_SetDefaultFixedFontSizeArgDetails
//go:noescape
func SetDefaultFixedFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_SetDefaultFontSizeArgDetails
//go:noescape
func SetDefaultFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_SetDefaultFontSizeArgDetails
//go:noescape
func SetDefaultFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_SetFontArgDetails
//go:noescape
func SetFontArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_SetFontArgDetails
//go:noescape
func SetFontArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings store_SetMinimumFontSizeArgDetails
//go:noescape
func SetMinimumFontSizeArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/fontsettings load_SetMinimumFontSizeArgDetails
//go:noescape
func SetMinimumFontSizeArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/fontsettings has_ClearDefaultFixedFontSize
//go:noescape
func HasFuncClearDefaultFixedFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_ClearDefaultFixedFontSize
//go:noescape
func FuncClearDefaultFixedFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_ClearDefaultFixedFontSize
//go:noescape
func CallClearDefaultFixedFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_ClearDefaultFixedFontSize
//go:noescape
func TryClearDefaultFixedFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_ClearDefaultFontSize
//go:noescape
func HasFuncClearDefaultFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_ClearDefaultFontSize
//go:noescape
func FuncClearDefaultFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_ClearDefaultFontSize
//go:noescape
func CallClearDefaultFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_ClearDefaultFontSize
//go:noescape
func TryClearDefaultFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_ClearFont
//go:noescape
func HasFuncClearFont() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_ClearFont
//go:noescape
func FuncClearFont(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_ClearFont
//go:noescape
func CallClearFont(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_ClearFont
//go:noescape
func TryClearFont(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_ClearMinimumFontSize
//go:noescape
func HasFuncClearMinimumFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_ClearMinimumFontSize
//go:noescape
func FuncClearMinimumFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_ClearMinimumFontSize
//go:noescape
func CallClearMinimumFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_ClearMinimumFontSize
//go:noescape
func TryClearMinimumFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_GetDefaultFixedFontSize
//go:noescape
func HasFuncGetDefaultFixedFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_GetDefaultFixedFontSize
//go:noescape
func FuncGetDefaultFixedFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_GetDefaultFixedFontSize
//go:noescape
func CallGetDefaultFixedFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_GetDefaultFixedFontSize
//go:noescape
func TryGetDefaultFixedFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_GetDefaultFontSize
//go:noescape
func HasFuncGetDefaultFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_GetDefaultFontSize
//go:noescape
func FuncGetDefaultFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_GetDefaultFontSize
//go:noescape
func CallGetDefaultFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_GetDefaultFontSize
//go:noescape
func TryGetDefaultFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_GetFont
//go:noescape
func HasFuncGetFont() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_GetFont
//go:noescape
func FuncGetFont(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_GetFont
//go:noescape
func CallGetFont(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_GetFont
//go:noescape
func TryGetFont(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_GetFontList
//go:noescape
func HasFuncGetFontList() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_GetFontList
//go:noescape
func FuncGetFontList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_GetFontList
//go:noescape
func CallGetFontList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_GetFontList
//go:noescape
func TryGetFontList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_GetMinimumFontSize
//go:noescape
func HasFuncGetMinimumFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_GetMinimumFontSize
//go:noescape
func FuncGetMinimumFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_GetMinimumFontSize
//go:noescape
func CallGetMinimumFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_GetMinimumFontSize
//go:noescape
func TryGetMinimumFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OnDefaultFixedFontSizeChanged
//go:noescape
func HasFuncOnDefaultFixedFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OnDefaultFixedFontSizeChanged
//go:noescape
func FuncOnDefaultFixedFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OnDefaultFixedFontSizeChanged
//go:noescape
func CallOnDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OnDefaultFixedFontSizeChanged
//go:noescape
func TryOnDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OffDefaultFixedFontSizeChanged
//go:noescape
func HasFuncOffDefaultFixedFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OffDefaultFixedFontSizeChanged
//go:noescape
func FuncOffDefaultFixedFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OffDefaultFixedFontSizeChanged
//go:noescape
func CallOffDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OffDefaultFixedFontSizeChanged
//go:noescape
func TryOffDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_HasOnDefaultFixedFontSizeChanged
//go:noescape
func HasFuncHasOnDefaultFixedFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_HasOnDefaultFixedFontSizeChanged
//go:noescape
func FuncHasOnDefaultFixedFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_HasOnDefaultFixedFontSizeChanged
//go:noescape
func CallHasOnDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_HasOnDefaultFixedFontSizeChanged
//go:noescape
func TryHasOnDefaultFixedFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OnDefaultFontSizeChanged
//go:noescape
func HasFuncOnDefaultFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OnDefaultFontSizeChanged
//go:noescape
func FuncOnDefaultFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OnDefaultFontSizeChanged
//go:noescape
func CallOnDefaultFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OnDefaultFontSizeChanged
//go:noescape
func TryOnDefaultFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OffDefaultFontSizeChanged
//go:noescape
func HasFuncOffDefaultFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OffDefaultFontSizeChanged
//go:noescape
func FuncOffDefaultFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OffDefaultFontSizeChanged
//go:noescape
func CallOffDefaultFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OffDefaultFontSizeChanged
//go:noescape
func TryOffDefaultFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_HasOnDefaultFontSizeChanged
//go:noescape
func HasFuncHasOnDefaultFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_HasOnDefaultFontSizeChanged
//go:noescape
func FuncHasOnDefaultFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_HasOnDefaultFontSizeChanged
//go:noescape
func CallHasOnDefaultFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_HasOnDefaultFontSizeChanged
//go:noescape
func TryHasOnDefaultFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OnFontChanged
//go:noescape
func HasFuncOnFontChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OnFontChanged
//go:noescape
func FuncOnFontChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OnFontChanged
//go:noescape
func CallOnFontChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OnFontChanged
//go:noescape
func TryOnFontChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OffFontChanged
//go:noescape
func HasFuncOffFontChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OffFontChanged
//go:noescape
func FuncOffFontChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OffFontChanged
//go:noescape
func CallOffFontChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OffFontChanged
//go:noescape
func TryOffFontChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_HasOnFontChanged
//go:noescape
func HasFuncHasOnFontChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_HasOnFontChanged
//go:noescape
func FuncHasOnFontChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_HasOnFontChanged
//go:noescape
func CallHasOnFontChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_HasOnFontChanged
//go:noescape
func TryHasOnFontChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OnMinimumFontSizeChanged
//go:noescape
func HasFuncOnMinimumFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OnMinimumFontSizeChanged
//go:noescape
func FuncOnMinimumFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OnMinimumFontSizeChanged
//go:noescape
func CallOnMinimumFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OnMinimumFontSizeChanged
//go:noescape
func TryOnMinimumFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_OffMinimumFontSizeChanged
//go:noescape
func HasFuncOffMinimumFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_OffMinimumFontSizeChanged
//go:noescape
func FuncOffMinimumFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_OffMinimumFontSizeChanged
//go:noescape
func CallOffMinimumFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_OffMinimumFontSizeChanged
//go:noescape
func TryOffMinimumFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_HasOnMinimumFontSizeChanged
//go:noescape
func HasFuncHasOnMinimumFontSizeChanged() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_HasOnMinimumFontSizeChanged
//go:noescape
func FuncHasOnMinimumFontSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_HasOnMinimumFontSizeChanged
//go:noescape
func CallHasOnMinimumFontSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/fontsettings try_HasOnMinimumFontSizeChanged
//go:noescape
func TryHasOnMinimumFontSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_SetDefaultFixedFontSize
//go:noescape
func HasFuncSetDefaultFixedFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_SetDefaultFixedFontSize
//go:noescape
func FuncSetDefaultFixedFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_SetDefaultFixedFontSize
//go:noescape
func CallSetDefaultFixedFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_SetDefaultFixedFontSize
//go:noescape
func TrySetDefaultFixedFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_SetDefaultFontSize
//go:noescape
func HasFuncSetDefaultFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_SetDefaultFontSize
//go:noescape
func FuncSetDefaultFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_SetDefaultFontSize
//go:noescape
func CallSetDefaultFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_SetDefaultFontSize
//go:noescape
func TrySetDefaultFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_SetFont
//go:noescape
func HasFuncSetFont() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_SetFont
//go:noescape
func FuncSetFont(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_SetFont
//go:noescape
func CallSetFont(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_SetFont
//go:noescape
func TrySetFont(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/fontsettings has_SetMinimumFontSize
//go:noescape
func HasFuncSetMinimumFontSize() js.Ref

//go:wasmimport plat/js/webext/fontsettings func_SetMinimumFontSize
//go:noescape
func FuncSetMinimumFontSize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings call_SetMinimumFontSize
//go:noescape
func CallSetMinimumFontSize(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/fontsettings try_SetMinimumFontSize
//go:noescape
func TrySetMinimumFontSize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
