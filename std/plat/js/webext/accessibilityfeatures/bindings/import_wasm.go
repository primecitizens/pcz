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

//go:wasmimport plat/js/webext/accessibilityfeatures get_AnimationPolicy
//go:noescape
func GetAnimationPolicy(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_AnimationPolicy
//go:noescape
func SetAnimationPolicy(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_Autoclick
//go:noescape
func GetAutoclick(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_Autoclick
//go:noescape
func SetAutoclick(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_CaretHighlight
//go:noescape
func GetCaretHighlight(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_CaretHighlight
//go:noescape
func SetCaretHighlight(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_CursorColor
//go:noescape
func GetCursorColor(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_CursorColor
//go:noescape
func SetCursorColor(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_CursorHighlight
//go:noescape
func GetCursorHighlight(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_CursorHighlight
//go:noescape
func SetCursorHighlight(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_Dictation
//go:noescape
func GetDictation(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_Dictation
//go:noescape
func SetDictation(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_DockedMagnifier
//go:noescape
func GetDockedMagnifier(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_DockedMagnifier
//go:noescape
func SetDockedMagnifier(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_FocusHighlight
//go:noescape
func GetFocusHighlight(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_FocusHighlight
//go:noescape
func SetFocusHighlight(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_HighContrast
//go:noescape
func GetHighContrast(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_HighContrast
//go:noescape
func SetHighContrast(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_LargeCursor
//go:noescape
func GetLargeCursor(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_LargeCursor
//go:noescape
func SetLargeCursor(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_ScreenMagnifier
//go:noescape
func GetScreenMagnifier(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_ScreenMagnifier
//go:noescape
func SetScreenMagnifier(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_SelectToSpeak
//go:noescape
func GetSelectToSpeak(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_SelectToSpeak
//go:noescape
func SetSelectToSpeak(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_SpokenFeedback
//go:noescape
func GetSpokenFeedback(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_SpokenFeedback
//go:noescape
func SetSpokenFeedback(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_StickyKeys
//go:noescape
func GetStickyKeys(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_StickyKeys
//go:noescape
func SetStickyKeys(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_SwitchAccess
//go:noescape
func GetSwitchAccess(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_SwitchAccess
//go:noescape
func SetSwitchAccess(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures get_VirtualKeyboard
//go:noescape
func GetVirtualKeyboard(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityfeatures set_VirtualKeyboard
//go:noescape
func SetVirtualKeyboard(
	val js.Ref) js.Ref
