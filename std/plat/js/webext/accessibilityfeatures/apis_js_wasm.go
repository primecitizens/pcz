// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package accessibilityfeatures

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/accessibilityfeatures/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/types"
)

// AnimationPolicy returns the value of property "WEBEXT.accessibilityFeatures.animationPolicy".
//
// The returned bool will be false if there is no such property.
func AnimationPolicy() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetAnimationPolicy(
		js.Pointer(&ret),
	)

	return
}

// SetAnimationPolicy sets the value of property "WEBEXT.accessibilityFeatures.animationPolicy" to val.
//
// It returns false if the property cannot be set.
func SetAnimationPolicy(val types.ChromeSetting) bool {
	return js.True == bindings.SetAnimationPolicy(
		val.Ref())
}

// Autoclick returns the value of property "WEBEXT.accessibilityFeatures.autoclick".
//
// The returned bool will be false if there is no such property.
func Autoclick() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetAutoclick(
		js.Pointer(&ret),
	)

	return
}

// SetAutoclick sets the value of property "WEBEXT.accessibilityFeatures.autoclick" to val.
//
// It returns false if the property cannot be set.
func SetAutoclick(val types.ChromeSetting) bool {
	return js.True == bindings.SetAutoclick(
		val.Ref())
}

// CaretHighlight returns the value of property "WEBEXT.accessibilityFeatures.caretHighlight".
//
// The returned bool will be false if there is no such property.
func CaretHighlight() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetCaretHighlight(
		js.Pointer(&ret),
	)

	return
}

// SetCaretHighlight sets the value of property "WEBEXT.accessibilityFeatures.caretHighlight" to val.
//
// It returns false if the property cannot be set.
func SetCaretHighlight(val types.ChromeSetting) bool {
	return js.True == bindings.SetCaretHighlight(
		val.Ref())
}

// CursorColor returns the value of property "WEBEXT.accessibilityFeatures.cursorColor".
//
// The returned bool will be false if there is no such property.
func CursorColor() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetCursorColor(
		js.Pointer(&ret),
	)

	return
}

// SetCursorColor sets the value of property "WEBEXT.accessibilityFeatures.cursorColor" to val.
//
// It returns false if the property cannot be set.
func SetCursorColor(val types.ChromeSetting) bool {
	return js.True == bindings.SetCursorColor(
		val.Ref())
}

// CursorHighlight returns the value of property "WEBEXT.accessibilityFeatures.cursorHighlight".
//
// The returned bool will be false if there is no such property.
func CursorHighlight() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetCursorHighlight(
		js.Pointer(&ret),
	)

	return
}

// SetCursorHighlight sets the value of property "WEBEXT.accessibilityFeatures.cursorHighlight" to val.
//
// It returns false if the property cannot be set.
func SetCursorHighlight(val types.ChromeSetting) bool {
	return js.True == bindings.SetCursorHighlight(
		val.Ref())
}

// Dictation returns the value of property "WEBEXT.accessibilityFeatures.dictation".
//
// The returned bool will be false if there is no such property.
func Dictation() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetDictation(
		js.Pointer(&ret),
	)

	return
}

// SetDictation sets the value of property "WEBEXT.accessibilityFeatures.dictation" to val.
//
// It returns false if the property cannot be set.
func SetDictation(val types.ChromeSetting) bool {
	return js.True == bindings.SetDictation(
		val.Ref())
}

// DockedMagnifier returns the value of property "WEBEXT.accessibilityFeatures.dockedMagnifier".
//
// The returned bool will be false if there is no such property.
func DockedMagnifier() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetDockedMagnifier(
		js.Pointer(&ret),
	)

	return
}

// SetDockedMagnifier sets the value of property "WEBEXT.accessibilityFeatures.dockedMagnifier" to val.
//
// It returns false if the property cannot be set.
func SetDockedMagnifier(val types.ChromeSetting) bool {
	return js.True == bindings.SetDockedMagnifier(
		val.Ref())
}

// FocusHighlight returns the value of property "WEBEXT.accessibilityFeatures.focusHighlight".
//
// The returned bool will be false if there is no such property.
func FocusHighlight() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetFocusHighlight(
		js.Pointer(&ret),
	)

	return
}

// SetFocusHighlight sets the value of property "WEBEXT.accessibilityFeatures.focusHighlight" to val.
//
// It returns false if the property cannot be set.
func SetFocusHighlight(val types.ChromeSetting) bool {
	return js.True == bindings.SetFocusHighlight(
		val.Ref())
}

// HighContrast returns the value of property "WEBEXT.accessibilityFeatures.highContrast".
//
// The returned bool will be false if there is no such property.
func HighContrast() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetHighContrast(
		js.Pointer(&ret),
	)

	return
}

// SetHighContrast sets the value of property "WEBEXT.accessibilityFeatures.highContrast" to val.
//
// It returns false if the property cannot be set.
func SetHighContrast(val types.ChromeSetting) bool {
	return js.True == bindings.SetHighContrast(
		val.Ref())
}

// LargeCursor returns the value of property "WEBEXT.accessibilityFeatures.largeCursor".
//
// The returned bool will be false if there is no such property.
func LargeCursor() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetLargeCursor(
		js.Pointer(&ret),
	)

	return
}

// SetLargeCursor sets the value of property "WEBEXT.accessibilityFeatures.largeCursor" to val.
//
// It returns false if the property cannot be set.
func SetLargeCursor(val types.ChromeSetting) bool {
	return js.True == bindings.SetLargeCursor(
		val.Ref())
}

// ScreenMagnifier returns the value of property "WEBEXT.accessibilityFeatures.screenMagnifier".
//
// The returned bool will be false if there is no such property.
func ScreenMagnifier() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetScreenMagnifier(
		js.Pointer(&ret),
	)

	return
}

// SetScreenMagnifier sets the value of property "WEBEXT.accessibilityFeatures.screenMagnifier" to val.
//
// It returns false if the property cannot be set.
func SetScreenMagnifier(val types.ChromeSetting) bool {
	return js.True == bindings.SetScreenMagnifier(
		val.Ref())
}

// SelectToSpeak returns the value of property "WEBEXT.accessibilityFeatures.selectToSpeak".
//
// The returned bool will be false if there is no such property.
func SelectToSpeak() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetSelectToSpeak(
		js.Pointer(&ret),
	)

	return
}

// SetSelectToSpeak sets the value of property "WEBEXT.accessibilityFeatures.selectToSpeak" to val.
//
// It returns false if the property cannot be set.
func SetSelectToSpeak(val types.ChromeSetting) bool {
	return js.True == bindings.SetSelectToSpeak(
		val.Ref())
}

// SpokenFeedback returns the value of property "WEBEXT.accessibilityFeatures.spokenFeedback".
//
// The returned bool will be false if there is no such property.
func SpokenFeedback() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetSpokenFeedback(
		js.Pointer(&ret),
	)

	return
}

// SetSpokenFeedback sets the value of property "WEBEXT.accessibilityFeatures.spokenFeedback" to val.
//
// It returns false if the property cannot be set.
func SetSpokenFeedback(val types.ChromeSetting) bool {
	return js.True == bindings.SetSpokenFeedback(
		val.Ref())
}

// StickyKeys returns the value of property "WEBEXT.accessibilityFeatures.stickyKeys".
//
// The returned bool will be false if there is no such property.
func StickyKeys() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetStickyKeys(
		js.Pointer(&ret),
	)

	return
}

// SetStickyKeys sets the value of property "WEBEXT.accessibilityFeatures.stickyKeys" to val.
//
// It returns false if the property cannot be set.
func SetStickyKeys(val types.ChromeSetting) bool {
	return js.True == bindings.SetStickyKeys(
		val.Ref())
}

// SwitchAccess returns the value of property "WEBEXT.accessibilityFeatures.switchAccess".
//
// The returned bool will be false if there is no such property.
func SwitchAccess() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetSwitchAccess(
		js.Pointer(&ret),
	)

	return
}

// SetSwitchAccess sets the value of property "WEBEXT.accessibilityFeatures.switchAccess" to val.
//
// It returns false if the property cannot be set.
func SetSwitchAccess(val types.ChromeSetting) bool {
	return js.True == bindings.SetSwitchAccess(
		val.Ref())
}

// VirtualKeyboard returns the value of property "WEBEXT.accessibilityFeatures.virtualKeyboard".
//
// The returned bool will be false if there is no such property.
func VirtualKeyboard() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetVirtualKeyboard(
		js.Pointer(&ret),
	)

	return
}

// SetVirtualKeyboard sets the value of property "WEBEXT.accessibilityFeatures.virtualKeyboard" to val.
//
// It returns false if the property cannot be set.
func SetVirtualKeyboard(val types.ChromeSetting) bool {
	return js.True == bindings.SetVirtualKeyboard(
		val.Ref())
}
