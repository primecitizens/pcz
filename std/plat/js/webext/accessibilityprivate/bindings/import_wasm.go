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

//go:wasmimport plat/js/webext/accessibilityprivate constof_AcceleratorAction
//go:noescape
func ConstOfAcceleratorAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_AccessibilityFeature
//go:noescape
func ConstOfAccessibilityFeature(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_AlertInfo
//go:noescape
func AlertInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_AlertInfo
//go:noescape
func AlertInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_AssistiveTechnologyType
//go:noescape
func ConstOfAssistiveTechnologyType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_DictationBubbleHintType
//go:noescape
func ConstOfDictationBubbleHintType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_DictationBubbleIconType
//go:noescape
func ConstOfDictationBubbleIconType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_DictationBubbleProperties
//go:noescape
func DictationBubblePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_DictationBubbleProperties
//go:noescape
func DictationBubblePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_DlcType
//go:noescape
func ConstOfDlcType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_ScreenRect
//go:noescape
func ScreenRectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_ScreenRect
//go:noescape
func ScreenRectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_FocusRingStackingOrder
//go:noescape
func ConstOfFocusRingStackingOrder(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_FocusType
//go:noescape
func ConstOfFocusType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_FocusRingInfo
//go:noescape
func FocusRingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_FocusRingInfo
//go:noescape
func FocusRingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_Gesture
//go:noescape
func ConstOfGesture(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate get_IS_DEFAULT_EVENT_SOURCE_TOUCH
//go:noescape
func GetIS_DEFAULT_EVENT_SOURCE_TOUCH(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate set_IS_DEFAULT_EVENT_SOURCE_TOUCH
//go:noescape
func SetIS_DEFAULT_EVENT_SOURCE_TOUCH(
	val float64) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_MagnifierCommand
//go:noescape
func ConstOfMagnifierCommand(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_PointScanPoint
//go:noescape
func PointScanPointJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_PointScanPoint
//go:noescape
func PointScanPointJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_PointScanState
//go:noescape
func ConstOfPointScanState(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_PumpkinData
//go:noescape
func PumpkinDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_PumpkinData
//go:noescape
func PumpkinDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate store_ScreenPoint
//go:noescape
func ScreenPointJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_ScreenPoint
//go:noescape
func ScreenPointJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_SelectToSpeakPanelAction
//go:noescape
func ConstOfSelectToSpeakPanelAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SelectToSpeakState
//go:noescape
func ConstOfSelectToSpeakState(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SetNativeChromeVoxResponse
//go:noescape
func ConstOfSetNativeChromeVoxResponse(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SwitchAccessBubble
//go:noescape
func ConstOfSwitchAccessBubble(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SwitchAccessCommand
//go:noescape
func ConstOfSwitchAccessCommand(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SwitchAccessMenuAction
//go:noescape
func ConstOfSwitchAccessMenuAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_SyntheticKeyboardModifiers
//go:noescape
func SyntheticKeyboardModifiersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_SyntheticKeyboardModifiers
//go:noescape
func SyntheticKeyboardModifiersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_SyntheticKeyboardEventType
//go:noescape
func ConstOfSyntheticKeyboardEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_SyntheticKeyboardEvent
//go:noescape
func SyntheticKeyboardEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_SyntheticKeyboardEvent
//go:noescape
func SyntheticKeyboardEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_SyntheticMouseEventButton
//go:noescape
func ConstOfSyntheticMouseEventButton(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate constof_SyntheticMouseEventType
//go:noescape
func ConstOfSyntheticMouseEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate store_SyntheticMouseEvent
//go:noescape
func SyntheticMouseEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate load_SyntheticMouseEvent
//go:noescape
func SyntheticMouseEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate constof_ToastType
//go:noescape
func ConstOfToastType(str js.Ref) uint32

//go:wasmimport plat/js/webext/accessibilityprivate has_DarkenScreen
//go:noescape
func HasFuncDarkenScreen() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_DarkenScreen
//go:noescape
func FuncDarkenScreen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_DarkenScreen
//go:noescape
func CallDarkenScreen(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_DarkenScreen
//go:noescape
func TryDarkenScreen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_EnableMouseEvents
//go:noescape
func HasFuncEnableMouseEvents() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_EnableMouseEvents
//go:noescape
func FuncEnableMouseEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_EnableMouseEvents
//go:noescape
func CallEnableMouseEvents(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_EnableMouseEvents
//go:noescape
func TryEnableMouseEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_ForwardKeyEventsToSwitchAccess
//go:noescape
func HasFuncForwardKeyEventsToSwitchAccess() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_ForwardKeyEventsToSwitchAccess
//go:noescape
func FuncForwardKeyEventsToSwitchAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_ForwardKeyEventsToSwitchAccess
//go:noescape
func CallForwardKeyEventsToSwitchAccess(
	retPtr unsafe.Pointer,
	shouldForward js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_ForwardKeyEventsToSwitchAccess
//go:noescape
func TryForwardKeyEventsToSwitchAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shouldForward js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_GetBatteryDescription
//go:noescape
func HasFuncGetBatteryDescription() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_GetBatteryDescription
//go:noescape
func FuncGetBatteryDescription(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_GetBatteryDescription
//go:noescape
func CallGetBatteryDescription(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_GetBatteryDescription
//go:noescape
func TryGetBatteryDescription(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_GetDisplayNameForLocale
//go:noescape
func HasFuncGetDisplayNameForLocale() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_GetDisplayNameForLocale
//go:noescape
func FuncGetDisplayNameForLocale(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_GetDisplayNameForLocale
//go:noescape
func CallGetDisplayNameForLocale(
	retPtr unsafe.Pointer,
	localeCodeToTranslate js.Ref,
	displayLocaleCode js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_GetDisplayNameForLocale
//go:noescape
func TryGetDisplayNameForLocale(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	localeCodeToTranslate js.Ref,
	displayLocaleCode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_GetDlcContents
//go:noescape
func HasFuncGetDlcContents() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_GetDlcContents
//go:noescape
func FuncGetDlcContents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_GetDlcContents
//go:noescape
func CallGetDlcContents(
	retPtr unsafe.Pointer,
	dlc uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_GetDlcContents
//go:noescape
func TryGetDlcContents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dlc uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_GetLocalizedDomKeyStringForKeyCode
//go:noescape
func HasFuncGetLocalizedDomKeyStringForKeyCode() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_GetLocalizedDomKeyStringForKeyCode
//go:noescape
func FuncGetLocalizedDomKeyStringForKeyCode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_GetLocalizedDomKeyStringForKeyCode
//go:noescape
func CallGetLocalizedDomKeyStringForKeyCode(
	retPtr unsafe.Pointer,
	keyCode float64)

//go:wasmimport plat/js/webext/accessibilityprivate try_GetLocalizedDomKeyStringForKeyCode
//go:noescape
func TryGetLocalizedDomKeyStringForKeyCode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyCode float64) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HandleScrollableBoundsForPointFound
//go:noescape
func HasFuncHandleScrollableBoundsForPointFound() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HandleScrollableBoundsForPointFound
//go:noescape
func FuncHandleScrollableBoundsForPointFound(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HandleScrollableBoundsForPointFound
//go:noescape
func CallHandleScrollableBoundsForPointFound(
	retPtr unsafe.Pointer,
	rect unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_HandleScrollableBoundsForPointFound
//go:noescape
func TryHandleScrollableBoundsForPointFound(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_InstallPumpkinForDictation
//go:noescape
func HasFuncInstallPumpkinForDictation() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_InstallPumpkinForDictation
//go:noescape
func FuncInstallPumpkinForDictation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_InstallPumpkinForDictation
//go:noescape
func CallInstallPumpkinForDictation(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_InstallPumpkinForDictation
//go:noescape
func TryInstallPumpkinForDictation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_IsFeatureEnabled
//go:noescape
func HasFuncIsFeatureEnabled() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_IsFeatureEnabled
//go:noescape
func FuncIsFeatureEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_IsFeatureEnabled
//go:noescape
func CallIsFeatureEnabled(
	retPtr unsafe.Pointer,
	feature uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_IsFeatureEnabled
//go:noescape
func TryIsFeatureEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_IsLacrosPrimary
//go:noescape
func HasFuncIsLacrosPrimary() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_IsLacrosPrimary
//go:noescape
func FuncIsLacrosPrimary(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_IsLacrosPrimary
//go:noescape
func CallIsLacrosPrimary(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_IsLacrosPrimary
//go:noescape
func TryIsLacrosPrimary(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_MagnifierCenterOnPoint
//go:noescape
func HasFuncMagnifierCenterOnPoint() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_MagnifierCenterOnPoint
//go:noescape
func FuncMagnifierCenterOnPoint(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_MagnifierCenterOnPoint
//go:noescape
func CallMagnifierCenterOnPoint(
	retPtr unsafe.Pointer,
	point unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_MagnifierCenterOnPoint
//go:noescape
func TryMagnifierCenterOnPoint(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_MoveMagnifierToRect
//go:noescape
func HasFuncMoveMagnifierToRect() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_MoveMagnifierToRect
//go:noescape
func FuncMoveMagnifierToRect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_MoveMagnifierToRect
//go:noescape
func CallMoveMagnifierToRect(
	retPtr unsafe.Pointer,
	rect unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_MoveMagnifierToRect
//go:noescape
func TryMoveMagnifierToRect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnAccessibilityGesture
//go:noescape
func HasFuncOnAccessibilityGesture() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnAccessibilityGesture
//go:noescape
func FuncOnAccessibilityGesture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnAccessibilityGesture
//go:noescape
func CallOnAccessibilityGesture(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnAccessibilityGesture
//go:noescape
func TryOnAccessibilityGesture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffAccessibilityGesture
//go:noescape
func HasFuncOffAccessibilityGesture() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffAccessibilityGesture
//go:noescape
func FuncOffAccessibilityGesture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffAccessibilityGesture
//go:noescape
func CallOffAccessibilityGesture(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffAccessibilityGesture
//go:noescape
func TryOffAccessibilityGesture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnAccessibilityGesture
//go:noescape
func HasFuncHasOnAccessibilityGesture() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnAccessibilityGesture
//go:noescape
func FuncHasOnAccessibilityGesture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnAccessibilityGesture
//go:noescape
func CallHasOnAccessibilityGesture(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnAccessibilityGesture
//go:noescape
func TryHasOnAccessibilityGesture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnAnnounceForAccessibility
//go:noescape
func HasFuncOnAnnounceForAccessibility() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnAnnounceForAccessibility
//go:noescape
func FuncOnAnnounceForAccessibility(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnAnnounceForAccessibility
//go:noescape
func CallOnAnnounceForAccessibility(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnAnnounceForAccessibility
//go:noescape
func TryOnAnnounceForAccessibility(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffAnnounceForAccessibility
//go:noescape
func HasFuncOffAnnounceForAccessibility() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffAnnounceForAccessibility
//go:noescape
func FuncOffAnnounceForAccessibility(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffAnnounceForAccessibility
//go:noescape
func CallOffAnnounceForAccessibility(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffAnnounceForAccessibility
//go:noescape
func TryOffAnnounceForAccessibility(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnAnnounceForAccessibility
//go:noescape
func HasFuncHasOnAnnounceForAccessibility() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnAnnounceForAccessibility
//go:noescape
func FuncHasOnAnnounceForAccessibility(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnAnnounceForAccessibility
//go:noescape
func CallHasOnAnnounceForAccessibility(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnAnnounceForAccessibility
//go:noescape
func TryHasOnAnnounceForAccessibility(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnCustomSpokenFeedbackToggled
//go:noescape
func HasFuncOnCustomSpokenFeedbackToggled() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnCustomSpokenFeedbackToggled
//go:noescape
func FuncOnCustomSpokenFeedbackToggled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnCustomSpokenFeedbackToggled
//go:noescape
func CallOnCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnCustomSpokenFeedbackToggled
//go:noescape
func TryOnCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffCustomSpokenFeedbackToggled
//go:noescape
func HasFuncOffCustomSpokenFeedbackToggled() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffCustomSpokenFeedbackToggled
//go:noescape
func FuncOffCustomSpokenFeedbackToggled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffCustomSpokenFeedbackToggled
//go:noescape
func CallOffCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffCustomSpokenFeedbackToggled
//go:noescape
func TryOffCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnCustomSpokenFeedbackToggled
//go:noescape
func HasFuncHasOnCustomSpokenFeedbackToggled() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnCustomSpokenFeedbackToggled
//go:noescape
func FuncHasOnCustomSpokenFeedbackToggled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnCustomSpokenFeedbackToggled
//go:noescape
func CallHasOnCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnCustomSpokenFeedbackToggled
//go:noescape
func TryHasOnCustomSpokenFeedbackToggled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnIntroduceChromeVox
//go:noescape
func HasFuncOnIntroduceChromeVox() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnIntroduceChromeVox
//go:noescape
func FuncOnIntroduceChromeVox(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnIntroduceChromeVox
//go:noescape
func CallOnIntroduceChromeVox(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnIntroduceChromeVox
//go:noescape
func TryOnIntroduceChromeVox(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffIntroduceChromeVox
//go:noescape
func HasFuncOffIntroduceChromeVox() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffIntroduceChromeVox
//go:noescape
func FuncOffIntroduceChromeVox(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffIntroduceChromeVox
//go:noescape
func CallOffIntroduceChromeVox(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffIntroduceChromeVox
//go:noescape
func TryOffIntroduceChromeVox(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnIntroduceChromeVox
//go:noescape
func HasFuncHasOnIntroduceChromeVox() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnIntroduceChromeVox
//go:noescape
func FuncHasOnIntroduceChromeVox(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnIntroduceChromeVox
//go:noescape
func CallHasOnIntroduceChromeVox(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnIntroduceChromeVox
//go:noescape
func TryHasOnIntroduceChromeVox(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnMagnifierBoundsChanged
//go:noescape
func HasFuncOnMagnifierBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnMagnifierBoundsChanged
//go:noescape
func FuncOnMagnifierBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnMagnifierBoundsChanged
//go:noescape
func CallOnMagnifierBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnMagnifierBoundsChanged
//go:noescape
func TryOnMagnifierBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffMagnifierBoundsChanged
//go:noescape
func HasFuncOffMagnifierBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffMagnifierBoundsChanged
//go:noescape
func FuncOffMagnifierBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffMagnifierBoundsChanged
//go:noescape
func CallOffMagnifierBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffMagnifierBoundsChanged
//go:noescape
func TryOffMagnifierBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnMagnifierBoundsChanged
//go:noescape
func HasFuncHasOnMagnifierBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnMagnifierBoundsChanged
//go:noescape
func FuncHasOnMagnifierBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnMagnifierBoundsChanged
//go:noescape
func CallHasOnMagnifierBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnMagnifierBoundsChanged
//go:noescape
func TryHasOnMagnifierBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnMagnifierCommand
//go:noescape
func HasFuncOnMagnifierCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnMagnifierCommand
//go:noescape
func FuncOnMagnifierCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnMagnifierCommand
//go:noescape
func CallOnMagnifierCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnMagnifierCommand
//go:noescape
func TryOnMagnifierCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffMagnifierCommand
//go:noescape
func HasFuncOffMagnifierCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffMagnifierCommand
//go:noescape
func FuncOffMagnifierCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffMagnifierCommand
//go:noescape
func CallOffMagnifierCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffMagnifierCommand
//go:noescape
func TryOffMagnifierCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnMagnifierCommand
//go:noescape
func HasFuncHasOnMagnifierCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnMagnifierCommand
//go:noescape
func FuncHasOnMagnifierCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnMagnifierCommand
//go:noescape
func CallHasOnMagnifierCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnMagnifierCommand
//go:noescape
func TryHasOnMagnifierCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnPointScanSet
//go:noescape
func HasFuncOnPointScanSet() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnPointScanSet
//go:noescape
func FuncOnPointScanSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnPointScanSet
//go:noescape
func CallOnPointScanSet(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnPointScanSet
//go:noescape
func TryOnPointScanSet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffPointScanSet
//go:noescape
func HasFuncOffPointScanSet() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffPointScanSet
//go:noescape
func FuncOffPointScanSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffPointScanSet
//go:noescape
func CallOffPointScanSet(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffPointScanSet
//go:noescape
func TryOffPointScanSet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnPointScanSet
//go:noescape
func HasFuncHasOnPointScanSet() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnPointScanSet
//go:noescape
func FuncHasOnPointScanSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnPointScanSet
//go:noescape
func CallHasOnPointScanSet(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnPointScanSet
//go:noescape
func TryHasOnPointScanSet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnScrollableBoundsForPointRequested
//go:noescape
func HasFuncOnScrollableBoundsForPointRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnScrollableBoundsForPointRequested
//go:noescape
func FuncOnScrollableBoundsForPointRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnScrollableBoundsForPointRequested
//go:noescape
func CallOnScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnScrollableBoundsForPointRequested
//go:noescape
func TryOnScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffScrollableBoundsForPointRequested
//go:noescape
func HasFuncOffScrollableBoundsForPointRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffScrollableBoundsForPointRequested
//go:noescape
func FuncOffScrollableBoundsForPointRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffScrollableBoundsForPointRequested
//go:noescape
func CallOffScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffScrollableBoundsForPointRequested
//go:noescape
func TryOffScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnScrollableBoundsForPointRequested
//go:noescape
func HasFuncHasOnScrollableBoundsForPointRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnScrollableBoundsForPointRequested
//go:noescape
func FuncHasOnScrollableBoundsForPointRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnScrollableBoundsForPointRequested
//go:noescape
func CallHasOnScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnScrollableBoundsForPointRequested
//go:noescape
func TryHasOnScrollableBoundsForPointRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSelectToSpeakContextMenuClicked
//go:noescape
func HasFuncOnSelectToSpeakContextMenuClicked() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSelectToSpeakContextMenuClicked
//go:noescape
func FuncOnSelectToSpeakContextMenuClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSelectToSpeakContextMenuClicked
//go:noescape
func CallOnSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSelectToSpeakContextMenuClicked
//go:noescape
func TryOnSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSelectToSpeakContextMenuClicked
//go:noescape
func HasFuncOffSelectToSpeakContextMenuClicked() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSelectToSpeakContextMenuClicked
//go:noescape
func FuncOffSelectToSpeakContextMenuClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSelectToSpeakContextMenuClicked
//go:noescape
func CallOffSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSelectToSpeakContextMenuClicked
//go:noescape
func TryOffSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSelectToSpeakContextMenuClicked
//go:noescape
func HasFuncHasOnSelectToSpeakContextMenuClicked() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSelectToSpeakContextMenuClicked
//go:noescape
func FuncHasOnSelectToSpeakContextMenuClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSelectToSpeakContextMenuClicked
//go:noescape
func CallHasOnSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSelectToSpeakContextMenuClicked
//go:noescape
func TryHasOnSelectToSpeakContextMenuClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSelectToSpeakKeysPressedChanged
//go:noescape
func HasFuncOnSelectToSpeakKeysPressedChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSelectToSpeakKeysPressedChanged
//go:noescape
func FuncOnSelectToSpeakKeysPressedChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSelectToSpeakKeysPressedChanged
//go:noescape
func CallOnSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSelectToSpeakKeysPressedChanged
//go:noescape
func TryOnSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSelectToSpeakKeysPressedChanged
//go:noescape
func HasFuncOffSelectToSpeakKeysPressedChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSelectToSpeakKeysPressedChanged
//go:noescape
func FuncOffSelectToSpeakKeysPressedChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSelectToSpeakKeysPressedChanged
//go:noescape
func CallOffSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSelectToSpeakKeysPressedChanged
//go:noescape
func TryOffSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSelectToSpeakKeysPressedChanged
//go:noescape
func HasFuncHasOnSelectToSpeakKeysPressedChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSelectToSpeakKeysPressedChanged
//go:noescape
func FuncHasOnSelectToSpeakKeysPressedChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSelectToSpeakKeysPressedChanged
//go:noescape
func CallHasOnSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSelectToSpeakKeysPressedChanged
//go:noescape
func TryHasOnSelectToSpeakKeysPressedChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSelectToSpeakMouseChanged
//go:noescape
func HasFuncOnSelectToSpeakMouseChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSelectToSpeakMouseChanged
//go:noescape
func FuncOnSelectToSpeakMouseChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSelectToSpeakMouseChanged
//go:noescape
func CallOnSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSelectToSpeakMouseChanged
//go:noescape
func TryOnSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSelectToSpeakMouseChanged
//go:noescape
func HasFuncOffSelectToSpeakMouseChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSelectToSpeakMouseChanged
//go:noescape
func FuncOffSelectToSpeakMouseChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSelectToSpeakMouseChanged
//go:noescape
func CallOffSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSelectToSpeakMouseChanged
//go:noescape
func TryOffSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSelectToSpeakMouseChanged
//go:noescape
func HasFuncHasOnSelectToSpeakMouseChanged() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSelectToSpeakMouseChanged
//go:noescape
func FuncHasOnSelectToSpeakMouseChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSelectToSpeakMouseChanged
//go:noescape
func CallHasOnSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSelectToSpeakMouseChanged
//go:noescape
func TryHasOnSelectToSpeakMouseChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSelectToSpeakPanelAction
//go:noescape
func HasFuncOnSelectToSpeakPanelAction() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSelectToSpeakPanelAction
//go:noescape
func FuncOnSelectToSpeakPanelAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSelectToSpeakPanelAction
//go:noescape
func CallOnSelectToSpeakPanelAction(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSelectToSpeakPanelAction
//go:noescape
func TryOnSelectToSpeakPanelAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSelectToSpeakPanelAction
//go:noescape
func HasFuncOffSelectToSpeakPanelAction() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSelectToSpeakPanelAction
//go:noescape
func FuncOffSelectToSpeakPanelAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSelectToSpeakPanelAction
//go:noescape
func CallOffSelectToSpeakPanelAction(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSelectToSpeakPanelAction
//go:noescape
func TryOffSelectToSpeakPanelAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSelectToSpeakPanelAction
//go:noescape
func HasFuncHasOnSelectToSpeakPanelAction() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSelectToSpeakPanelAction
//go:noescape
func FuncHasOnSelectToSpeakPanelAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSelectToSpeakPanelAction
//go:noescape
func CallHasOnSelectToSpeakPanelAction(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSelectToSpeakPanelAction
//go:noescape
func TryHasOnSelectToSpeakPanelAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSelectToSpeakStateChangeRequested
//go:noescape
func HasFuncOnSelectToSpeakStateChangeRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSelectToSpeakStateChangeRequested
//go:noescape
func FuncOnSelectToSpeakStateChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSelectToSpeakStateChangeRequested
//go:noescape
func CallOnSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSelectToSpeakStateChangeRequested
//go:noescape
func TryOnSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSelectToSpeakStateChangeRequested
//go:noescape
func HasFuncOffSelectToSpeakStateChangeRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSelectToSpeakStateChangeRequested
//go:noescape
func FuncOffSelectToSpeakStateChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSelectToSpeakStateChangeRequested
//go:noescape
func CallOffSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSelectToSpeakStateChangeRequested
//go:noescape
func TryOffSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSelectToSpeakStateChangeRequested
//go:noescape
func HasFuncHasOnSelectToSpeakStateChangeRequested() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSelectToSpeakStateChangeRequested
//go:noescape
func FuncHasOnSelectToSpeakStateChangeRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSelectToSpeakStateChangeRequested
//go:noescape
func CallHasOnSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSelectToSpeakStateChangeRequested
//go:noescape
func TryHasOnSelectToSpeakStateChangeRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnShowChromeVoxTutorial
//go:noescape
func HasFuncOnShowChromeVoxTutorial() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnShowChromeVoxTutorial
//go:noescape
func FuncOnShowChromeVoxTutorial(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnShowChromeVoxTutorial
//go:noescape
func CallOnShowChromeVoxTutorial(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnShowChromeVoxTutorial
//go:noescape
func TryOnShowChromeVoxTutorial(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffShowChromeVoxTutorial
//go:noescape
func HasFuncOffShowChromeVoxTutorial() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffShowChromeVoxTutorial
//go:noescape
func FuncOffShowChromeVoxTutorial(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffShowChromeVoxTutorial
//go:noescape
func CallOffShowChromeVoxTutorial(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffShowChromeVoxTutorial
//go:noescape
func TryOffShowChromeVoxTutorial(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnShowChromeVoxTutorial
//go:noescape
func HasFuncHasOnShowChromeVoxTutorial() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnShowChromeVoxTutorial
//go:noescape
func FuncHasOnShowChromeVoxTutorial(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnShowChromeVoxTutorial
//go:noescape
func CallHasOnShowChromeVoxTutorial(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnShowChromeVoxTutorial
//go:noescape
func TryHasOnShowChromeVoxTutorial(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnSwitchAccessCommand
//go:noescape
func HasFuncOnSwitchAccessCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnSwitchAccessCommand
//go:noescape
func FuncOnSwitchAccessCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnSwitchAccessCommand
//go:noescape
func CallOnSwitchAccessCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnSwitchAccessCommand
//go:noescape
func TryOnSwitchAccessCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffSwitchAccessCommand
//go:noescape
func HasFuncOffSwitchAccessCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffSwitchAccessCommand
//go:noescape
func FuncOffSwitchAccessCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffSwitchAccessCommand
//go:noescape
func CallOffSwitchAccessCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffSwitchAccessCommand
//go:noescape
func TryOffSwitchAccessCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnSwitchAccessCommand
//go:noescape
func HasFuncHasOnSwitchAccessCommand() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnSwitchAccessCommand
//go:noescape
func FuncHasOnSwitchAccessCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnSwitchAccessCommand
//go:noescape
func CallHasOnSwitchAccessCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnSwitchAccessCommand
//go:noescape
func TryHasOnSwitchAccessCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnToggleDictation
//go:noescape
func HasFuncOnToggleDictation() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnToggleDictation
//go:noescape
func FuncOnToggleDictation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnToggleDictation
//go:noescape
func CallOnToggleDictation(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnToggleDictation
//go:noescape
func TryOnToggleDictation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffToggleDictation
//go:noescape
func HasFuncOffToggleDictation() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffToggleDictation
//go:noescape
func FuncOffToggleDictation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffToggleDictation
//go:noescape
func CallOffToggleDictation(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffToggleDictation
//go:noescape
func TryOffToggleDictation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnToggleDictation
//go:noescape
func HasFuncHasOnToggleDictation() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnToggleDictation
//go:noescape
func FuncHasOnToggleDictation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnToggleDictation
//go:noescape
func CallHasOnToggleDictation(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnToggleDictation
//go:noescape
func TryHasOnToggleDictation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnTwoFingerTouchStart
//go:noescape
func HasFuncOnTwoFingerTouchStart() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnTwoFingerTouchStart
//go:noescape
func FuncOnTwoFingerTouchStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnTwoFingerTouchStart
//go:noescape
func CallOnTwoFingerTouchStart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnTwoFingerTouchStart
//go:noescape
func TryOnTwoFingerTouchStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffTwoFingerTouchStart
//go:noescape
func HasFuncOffTwoFingerTouchStart() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffTwoFingerTouchStart
//go:noescape
func FuncOffTwoFingerTouchStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffTwoFingerTouchStart
//go:noescape
func CallOffTwoFingerTouchStart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffTwoFingerTouchStart
//go:noescape
func TryOffTwoFingerTouchStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnTwoFingerTouchStart
//go:noescape
func HasFuncHasOnTwoFingerTouchStart() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnTwoFingerTouchStart
//go:noescape
func FuncHasOnTwoFingerTouchStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnTwoFingerTouchStart
//go:noescape
func CallHasOnTwoFingerTouchStart(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnTwoFingerTouchStart
//go:noescape
func TryHasOnTwoFingerTouchStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OnTwoFingerTouchStop
//go:noescape
func HasFuncOnTwoFingerTouchStop() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OnTwoFingerTouchStop
//go:noescape
func FuncOnTwoFingerTouchStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OnTwoFingerTouchStop
//go:noescape
func CallOnTwoFingerTouchStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OnTwoFingerTouchStop
//go:noescape
func TryOnTwoFingerTouchStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OffTwoFingerTouchStop
//go:noescape
func HasFuncOffTwoFingerTouchStop() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OffTwoFingerTouchStop
//go:noescape
func FuncOffTwoFingerTouchStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OffTwoFingerTouchStop
//go:noescape
func CallOffTwoFingerTouchStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OffTwoFingerTouchStop
//go:noescape
func TryOffTwoFingerTouchStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_HasOnTwoFingerTouchStop
//go:noescape
func HasFuncHasOnTwoFingerTouchStop() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_HasOnTwoFingerTouchStop
//go:noescape
func FuncHasOnTwoFingerTouchStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_HasOnTwoFingerTouchStop
//go:noescape
func CallHasOnTwoFingerTouchStop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_HasOnTwoFingerTouchStop
//go:noescape
func TryHasOnTwoFingerTouchStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_OpenSettingsSubpage
//go:noescape
func HasFuncOpenSettingsSubpage() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_OpenSettingsSubpage
//go:noescape
func FuncOpenSettingsSubpage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_OpenSettingsSubpage
//go:noescape
func CallOpenSettingsSubpage(
	retPtr unsafe.Pointer,
	subpage js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_OpenSettingsSubpage
//go:noescape
func TryOpenSettingsSubpage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	subpage js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_PerformAcceleratorAction
//go:noescape
func HasFuncPerformAcceleratorAction() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_PerformAcceleratorAction
//go:noescape
func FuncPerformAcceleratorAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_PerformAcceleratorAction
//go:noescape
func CallPerformAcceleratorAction(
	retPtr unsafe.Pointer,
	acceleratorAction uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_PerformAcceleratorAction
//go:noescape
func TryPerformAcceleratorAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	acceleratorAction uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SendSyntheticKeyEvent
//go:noescape
func HasFuncSendSyntheticKeyEvent() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SendSyntheticKeyEvent
//go:noescape
func FuncSendSyntheticKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SendSyntheticKeyEvent
//go:noescape
func CallSendSyntheticKeyEvent(
	retPtr unsafe.Pointer,
	keyEvent unsafe.Pointer,
	useRewriters js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SendSyntheticKeyEvent
//go:noescape
func TrySendSyntheticKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyEvent unsafe.Pointer,
	useRewriters js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SendSyntheticMouseEvent
//go:noescape
func HasFuncSendSyntheticMouseEvent() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SendSyntheticMouseEvent
//go:noescape
func FuncSendSyntheticMouseEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SendSyntheticMouseEvent
//go:noescape
func CallSendSyntheticMouseEvent(
	retPtr unsafe.Pointer,
	mouseEvent unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_SendSyntheticMouseEvent
//go:noescape
func TrySendSyntheticMouseEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	mouseEvent unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetFocusRings
//go:noescape
func HasFuncSetFocusRings() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetFocusRings
//go:noescape
func FuncSetFocusRings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetFocusRings
//go:noescape
func CallSetFocusRings(
	retPtr unsafe.Pointer,
	focusRings js.Ref,
	atType uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetFocusRings
//go:noescape
func TrySetFocusRings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	focusRings js.Ref,
	atType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetHighlights
//go:noescape
func HasFuncSetHighlights() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetHighlights
//go:noescape
func FuncSetHighlights(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetHighlights
//go:noescape
func CallSetHighlights(
	retPtr unsafe.Pointer,
	rects js.Ref,
	color js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetHighlights
//go:noescape
func TrySetHighlights(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rects js.Ref,
	color js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetKeyboardListener
//go:noescape
func HasFuncSetKeyboardListener() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetKeyboardListener
//go:noescape
func FuncSetKeyboardListener(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetKeyboardListener
//go:noescape
func CallSetKeyboardListener(
	retPtr unsafe.Pointer,
	enabled js.Ref,
	capture js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetKeyboardListener
//go:noescape
func TrySetKeyboardListener(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref,
	capture js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetNativeAccessibilityEnabled
//go:noescape
func HasFuncSetNativeAccessibilityEnabled() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetNativeAccessibilityEnabled
//go:noescape
func FuncSetNativeAccessibilityEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetNativeAccessibilityEnabled
//go:noescape
func CallSetNativeAccessibilityEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetNativeAccessibilityEnabled
//go:noescape
func TrySetNativeAccessibilityEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetNativeChromeVoxArcSupportForCurrentApp
//go:noescape
func HasFuncSetNativeChromeVoxArcSupportForCurrentApp() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetNativeChromeVoxArcSupportForCurrentApp
//go:noescape
func FuncSetNativeChromeVoxArcSupportForCurrentApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetNativeChromeVoxArcSupportForCurrentApp
//go:noescape
func CallSetNativeChromeVoxArcSupportForCurrentApp(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetNativeChromeVoxArcSupportForCurrentApp
//go:noescape
func TrySetNativeChromeVoxArcSupportForCurrentApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetPointScanState
//go:noescape
func HasFuncSetPointScanState() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetPointScanState
//go:noescape
func FuncSetPointScanState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetPointScanState
//go:noescape
func CallSetPointScanState(
	retPtr unsafe.Pointer,
	state uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetPointScanState
//go:noescape
func TrySetPointScanState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetSelectToSpeakState
//go:noescape
func HasFuncSetSelectToSpeakState() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetSelectToSpeakState
//go:noescape
func FuncSetSelectToSpeakState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetSelectToSpeakState
//go:noescape
func CallSetSelectToSpeakState(
	retPtr unsafe.Pointer,
	state uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetSelectToSpeakState
//go:noescape
func TrySetSelectToSpeakState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SetVirtualKeyboardVisible
//go:noescape
func HasFuncSetVirtualKeyboardVisible() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SetVirtualKeyboardVisible
//go:noescape
func FuncSetVirtualKeyboardVisible(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SetVirtualKeyboardVisible
//go:noescape
func CallSetVirtualKeyboardVisible(
	retPtr unsafe.Pointer,
	isVisible js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_SetVirtualKeyboardVisible
//go:noescape
func TrySetVirtualKeyboardVisible(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	isVisible js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_ShowConfirmationDialog
//go:noescape
func HasFuncShowConfirmationDialog() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_ShowConfirmationDialog
//go:noescape
func FuncShowConfirmationDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_ShowConfirmationDialog
//go:noescape
func CallShowConfirmationDialog(
	retPtr unsafe.Pointer,
	title js.Ref,
	description js.Ref,
	cancelName js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_ShowConfirmationDialog
//go:noescape
func TryShowConfirmationDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	title js.Ref,
	description js.Ref,
	cancelName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_ShowToast
//go:noescape
func HasFuncShowToast() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_ShowToast
//go:noescape
func FuncShowToast(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_ShowToast
//go:noescape
func CallShowToast(
	retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/webext/accessibilityprivate try_ShowToast
//go:noescape
func TryShowToast(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_SilenceSpokenFeedback
//go:noescape
func HasFuncSilenceSpokenFeedback() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_SilenceSpokenFeedback
//go:noescape
func FuncSilenceSpokenFeedback(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_SilenceSpokenFeedback
//go:noescape
func CallSilenceSpokenFeedback(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_SilenceSpokenFeedback
//go:noescape
func TrySilenceSpokenFeedback(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_ToggleDictation
//go:noescape
func HasFuncToggleDictation() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_ToggleDictation
//go:noescape
func FuncToggleDictation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_ToggleDictation
//go:noescape
func CallToggleDictation(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_ToggleDictation
//go:noescape
func TryToggleDictation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_UpdateDictationBubble
//go:noescape
func HasFuncUpdateDictationBubble() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_UpdateDictationBubble
//go:noescape
func FuncUpdateDictationBubble(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_UpdateDictationBubble
//go:noescape
func CallUpdateDictationBubble(
	retPtr unsafe.Pointer,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate try_UpdateDictationBubble
//go:noescape
func TryUpdateDictationBubble(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_UpdateSelectToSpeakPanel
//go:noescape
func HasFuncUpdateSelectToSpeakPanel() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_UpdateSelectToSpeakPanel
//go:noescape
func FuncUpdateSelectToSpeakPanel(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_UpdateSelectToSpeakPanel
//go:noescape
func CallUpdateSelectToSpeakPanel(
	retPtr unsafe.Pointer,
	show js.Ref,
	anchor unsafe.Pointer,
	isPaused js.Ref,
	speed float64)

//go:wasmimport plat/js/webext/accessibilityprivate try_UpdateSelectToSpeakPanel
//go:noescape
func TryUpdateSelectToSpeakPanel(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	show js.Ref,
	anchor unsafe.Pointer,
	isPaused js.Ref,
	speed float64) (ok js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate has_UpdateSwitchAccessBubble
//go:noescape
func HasFuncUpdateSwitchAccessBubble() js.Ref

//go:wasmimport plat/js/webext/accessibilityprivate func_UpdateSwitchAccessBubble
//go:noescape
func FuncUpdateSwitchAccessBubble(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/accessibilityprivate call_UpdateSwitchAccessBubble
//go:noescape
func CallUpdateSwitchAccessBubble(
	retPtr unsafe.Pointer,
	bubble uint32,
	show js.Ref,
	anchor unsafe.Pointer,
	actions js.Ref)

//go:wasmimport plat/js/webext/accessibilityprivate try_UpdateSwitchAccessBubble
//go:noescape
func TryUpdateSwitchAccessBubble(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bubble uint32,
	show js.Ref,
	anchor unsafe.Pointer,
	actions js.Ref) (ok js.Ref)
