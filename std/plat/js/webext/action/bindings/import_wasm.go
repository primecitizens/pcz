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

//go:wasmimport plat/js/webext/action store_OpenPopupOptions
//go:noescape
func OpenPopupOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_OpenPopupOptions
//go:noescape
func OpenPopupOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetBadgeBackgroundColorArgDetails
//go:noescape
func SetBadgeBackgroundColorArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetBadgeBackgroundColorArgDetails
//go:noescape
func SetBadgeBackgroundColorArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetBadgeTextArgDetails
//go:noescape
func SetBadgeTextArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetBadgeTextArgDetails
//go:noescape
func SetBadgeTextArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetBadgeTextColorArgDetails
//go:noescape
func SetBadgeTextColorArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetBadgeTextColorArgDetails
//go:noescape
func SetBadgeTextColorArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetIconArgDetails
//go:noescape
func SetIconArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetIconArgDetails
//go:noescape
func SetIconArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetPopupArgDetails
//go:noescape
func SetPopupArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetPopupArgDetails
//go:noescape
func SetPopupArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_SetTitleArgDetails
//go:noescape
func SetTitleArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_SetTitleArgDetails
//go:noescape
func SetTitleArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_TabDetails
//go:noescape
func TabDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_TabDetails
//go:noescape
func TabDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action store_UserSettings
//go:noescape
func UserSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/action load_UserSettings
//go:noescape
func UserSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/action has_Disable
//go:noescape
func HasFuncDisable() js.Ref

//go:wasmimport plat/js/webext/action func_Disable
//go:noescape
func FuncDisable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_Disable
//go:noescape
func CallDisable(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/action try_Disable
//go:noescape
func TryDisable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_Enable
//go:noescape
func HasFuncEnable() js.Ref

//go:wasmimport plat/js/webext/action func_Enable
//go:noescape
func FuncEnable(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_Enable
//go:noescape
func CallEnable(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/action try_Enable
//go:noescape
func TryEnable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetBadgeBackgroundColor
//go:noescape
func HasFuncGetBadgeBackgroundColor() js.Ref

//go:wasmimport plat/js/webext/action func_GetBadgeBackgroundColor
//go:noescape
func FuncGetBadgeBackgroundColor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetBadgeBackgroundColor
//go:noescape
func CallGetBadgeBackgroundColor(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetBadgeBackgroundColor
//go:noescape
func TryGetBadgeBackgroundColor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetBadgeText
//go:noescape
func HasFuncGetBadgeText() js.Ref

//go:wasmimport plat/js/webext/action func_GetBadgeText
//go:noescape
func FuncGetBadgeText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetBadgeText
//go:noescape
func CallGetBadgeText(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetBadgeText
//go:noescape
func TryGetBadgeText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetBadgeTextColor
//go:noescape
func HasFuncGetBadgeTextColor() js.Ref

//go:wasmimport plat/js/webext/action func_GetBadgeTextColor
//go:noescape
func FuncGetBadgeTextColor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetBadgeTextColor
//go:noescape
func CallGetBadgeTextColor(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetBadgeTextColor
//go:noescape
func TryGetBadgeTextColor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetPopup
//go:noescape
func HasFuncGetPopup() js.Ref

//go:wasmimport plat/js/webext/action func_GetPopup
//go:noescape
func FuncGetPopup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetPopup
//go:noescape
func CallGetPopup(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetPopup
//go:noescape
func TryGetPopup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetTitle
//go:noescape
func HasFuncGetTitle() js.Ref

//go:wasmimport plat/js/webext/action func_GetTitle
//go:noescape
func FuncGetTitle(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetTitle
//go:noescape
func CallGetTitle(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetTitle
//go:noescape
func TryGetTitle(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_GetUserSettings
//go:noescape
func HasFuncGetUserSettings() js.Ref

//go:wasmimport plat/js/webext/action func_GetUserSettings
//go:noescape
func FuncGetUserSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_GetUserSettings
//go:noescape
func CallGetUserSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_GetUserSettings
//go:noescape
func TryGetUserSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_IsEnabled
//go:noescape
func HasFuncIsEnabled() js.Ref

//go:wasmimport plat/js/webext/action func_IsEnabled
//go:noescape
func FuncIsEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_IsEnabled
//go:noescape
func CallIsEnabled(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/action try_IsEnabled
//go:noescape
func TryIsEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_OnClicked
//go:noescape
func HasFuncOnClicked() js.Ref

//go:wasmimport plat/js/webext/action func_OnClicked
//go:noescape
func FuncOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_OnClicked
//go:noescape
func CallOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/action try_OnClicked
//go:noescape
func TryOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_OffClicked
//go:noescape
func HasFuncOffClicked() js.Ref

//go:wasmimport plat/js/webext/action func_OffClicked
//go:noescape
func FuncOffClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_OffClicked
//go:noescape
func CallOffClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/action try_OffClicked
//go:noescape
func TryOffClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_HasOnClicked
//go:noescape
func HasFuncHasOnClicked() js.Ref

//go:wasmimport plat/js/webext/action func_HasOnClicked
//go:noescape
func FuncHasOnClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_HasOnClicked
//go:noescape
func CallHasOnClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/action try_HasOnClicked
//go:noescape
func TryHasOnClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_OpenPopup
//go:noescape
func HasFuncOpenPopup() js.Ref

//go:wasmimport plat/js/webext/action func_OpenPopup
//go:noescape
func FuncOpenPopup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_OpenPopup
//go:noescape
func CallOpenPopup(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_OpenPopup
//go:noescape
func TryOpenPopup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetBadgeBackgroundColor
//go:noescape
func HasFuncSetBadgeBackgroundColor() js.Ref

//go:wasmimport plat/js/webext/action func_SetBadgeBackgroundColor
//go:noescape
func FuncSetBadgeBackgroundColor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetBadgeBackgroundColor
//go:noescape
func CallSetBadgeBackgroundColor(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetBadgeBackgroundColor
//go:noescape
func TrySetBadgeBackgroundColor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetBadgeText
//go:noescape
func HasFuncSetBadgeText() js.Ref

//go:wasmimport plat/js/webext/action func_SetBadgeText
//go:noescape
func FuncSetBadgeText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetBadgeText
//go:noescape
func CallSetBadgeText(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetBadgeText
//go:noescape
func TrySetBadgeText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetBadgeTextColor
//go:noescape
func HasFuncSetBadgeTextColor() js.Ref

//go:wasmimport plat/js/webext/action func_SetBadgeTextColor
//go:noescape
func FuncSetBadgeTextColor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetBadgeTextColor
//go:noescape
func CallSetBadgeTextColor(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetBadgeTextColor
//go:noescape
func TrySetBadgeTextColor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetIcon
//go:noescape
func HasFuncSetIcon() js.Ref

//go:wasmimport plat/js/webext/action func_SetIcon
//go:noescape
func FuncSetIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetIcon
//go:noescape
func CallSetIcon(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetIcon
//go:noescape
func TrySetIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetPopup
//go:noescape
func HasFuncSetPopup() js.Ref

//go:wasmimport plat/js/webext/action func_SetPopup
//go:noescape
func FuncSetPopup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetPopup
//go:noescape
func CallSetPopup(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetPopup
//go:noescape
func TrySetPopup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/action has_SetTitle
//go:noescape
func HasFuncSetTitle() js.Ref

//go:wasmimport plat/js/webext/action func_SetTitle
//go:noescape
func FuncSetTitle(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/action call_SetTitle
//go:noescape
func CallSetTitle(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/action try_SetTitle
//go:noescape
func TrySetTitle(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
