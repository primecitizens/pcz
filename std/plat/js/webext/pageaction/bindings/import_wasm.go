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

//go:wasmimport plat/js/webext/pageaction store_ImageDataType
//go:noescape
func ImageDataTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pageaction load_ImageDataType
//go:noescape
func ImageDataTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pageaction store_SetIconArgDetails
//go:noescape
func SetIconArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pageaction load_SetIconArgDetails
//go:noescape
func SetIconArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pageaction store_SetPopupArgDetails
//go:noescape
func SetPopupArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pageaction load_SetPopupArgDetails
//go:noescape
func SetPopupArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pageaction store_SetTitleArgDetails
//go:noescape
func SetTitleArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pageaction load_SetTitleArgDetails
//go:noescape
func SetTitleArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pageaction store_TabDetails
//go:noescape
func TabDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pageaction load_TabDetails
//go:noescape
func TabDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pageaction has_GetPopup
//go:noescape
func HasFuncGetPopup() js.Ref

//go:wasmimport plat/js/webext/pageaction func_GetPopup
//go:noescape
func FuncGetPopup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_GetPopup
//go:noescape
func CallGetPopup(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction try_GetPopup
//go:noescape
func TryGetPopup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_GetTitle
//go:noescape
func HasFuncGetTitle() js.Ref

//go:wasmimport plat/js/webext/pageaction func_GetTitle
//go:noescape
func FuncGetTitle(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_GetTitle
//go:noescape
func CallGetTitle(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction try_GetTitle
//go:noescape
func TryGetTitle(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_Hide
//go:noescape
func HasFuncHide() js.Ref

//go:wasmimport plat/js/webext/pageaction func_Hide
//go:noescape
func FuncHide(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_Hide
//go:noescape
func CallHide(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/pageaction try_Hide
//go:noescape
func TryHide(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/pageaction func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pageaction try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/pageaction func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pageaction try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/pageaction func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pageaction try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_SetIcon
//go:noescape
func HasFuncSetIcon() js.Ref

//go:wasmimport plat/js/webext/pageaction func_SetIcon
//go:noescape
func FuncSetIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_SetIcon
//go:noescape
func CallSetIcon(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction try_SetIcon
//go:noescape
func TrySetIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_SetPopup
//go:noescape
func HasFuncSetPopup() js.Ref

//go:wasmimport plat/js/webext/pageaction func_SetPopup
//go:noescape
func FuncSetPopup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_SetPopup
//go:noescape
func CallSetPopup(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction try_SetPopup
//go:noescape
func TrySetPopup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_SetTitle
//go:noescape
func HasFuncSetTitle() js.Ref

//go:wasmimport plat/js/webext/pageaction func_SetTitle
//go:noescape
func FuncSetTitle(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_SetTitle
//go:noescape
func CallSetTitle(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction try_SetTitle
//go:noescape
func TrySetTitle(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pageaction has_Show
//go:noescape
func HasFuncShow() js.Ref

//go:wasmimport plat/js/webext/pageaction func_Show
//go:noescape
func FuncShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pageaction call_Show
//go:noescape
func CallShow(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/pageaction try_Show
//go:noescape
func TryShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)
