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

//go:wasmimport plat/js/webext/input/ime constof_AssistiveWindowButton
//go:noescape
func ConstOfAssistiveWindowButton(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime constof_AssistiveWindowType
//go:noescape
func ConstOfAssistiveWindowType(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_AssistiveWindowProperties
//go:noescape
func AssistiveWindowPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_AssistiveWindowProperties
//go:noescape
func AssistiveWindowPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_AutoCapitalizeType
//go:noescape
func ConstOfAutoCapitalizeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_ClearCompositionArgParameters
//go:noescape
func ClearCompositionArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_ClearCompositionArgParameters
//go:noescape
func ClearCompositionArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_CommitTextArgParameters
//go:noescape
func CommitTextArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_CommitTextArgParameters
//go:noescape
func CommitTextArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_DeleteSurroundingTextArgParameters
//go:noescape
func DeleteSurroundingTextArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_DeleteSurroundingTextArgParameters
//go:noescape
func DeleteSurroundingTextArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_InputContextType
//go:noescape
func ConstOfInputContextType(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_InputContext
//go:noescape
func InputContextJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_InputContext
//go:noescape
func InputContextJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_KeyboardEventType
//go:noescape
func ConstOfKeyboardEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_KeyboardEvent
//go:noescape
func KeyboardEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_KeyboardEvent
//go:noescape
func KeyboardEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_MenuItemStyle
//go:noescape
func ConstOfMenuItemStyle(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_MenuItem
//go:noescape
func MenuItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_MenuItem
//go:noescape
func MenuItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_MenuParameters
//go:noescape
func MenuParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_MenuParameters
//go:noescape
func MenuParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_MouseButton
//go:noescape
func ConstOfMouseButton(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_OnAssistiveWindowButtonClickedArgDetails
//go:noescape
func OnAssistiveWindowButtonClickedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_OnAssistiveWindowButtonClickedArgDetails
//go:noescape
func OnAssistiveWindowButtonClickedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_OnSurroundingTextChangedArgSurroundingInfo
//go:noescape
func OnSurroundingTextChangedArgSurroundingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_OnSurroundingTextChangedArgSurroundingInfo
//go:noescape
func OnSurroundingTextChangedArgSurroundingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_ScreenType
//go:noescape
func ConstOfScreenType(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_SendKeyEventsArgParameters
//go:noescape
func SendKeyEventsArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SendKeyEventsArgParameters
//go:noescape
func SendKeyEventsArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetAssistiveWindowButtonHighlightedArgParameters
//go:noescape
func SetAssistiveWindowButtonHighlightedArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetAssistiveWindowButtonHighlightedArgParameters
//go:noescape
func SetAssistiveWindowButtonHighlightedArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetAssistiveWindowPropertiesArgParameters
//go:noescape
func SetAssistiveWindowPropertiesArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetAssistiveWindowPropertiesArgParameters
//go:noescape
func SetAssistiveWindowPropertiesArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_WindowPosition
//go:noescape
func ConstOfWindowPosition(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_SetCandidateWindowPropertiesArgParametersFieldProperties
//go:noescape
func SetCandidateWindowPropertiesArgParametersFieldPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCandidateWindowPropertiesArgParametersFieldProperties
//go:noescape
func SetCandidateWindowPropertiesArgParametersFieldPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCandidateWindowPropertiesArgParameters
//go:noescape
func SetCandidateWindowPropertiesArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCandidateWindowPropertiesArgParameters
//go:noescape
func SetCandidateWindowPropertiesArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCandidatesArgParametersFieldCandidatesElemFieldUsage
//go:noescape
func SetCandidatesArgParametersFieldCandidatesElemFieldUsageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCandidatesArgParametersFieldCandidatesElemFieldUsage
//go:noescape
func SetCandidatesArgParametersFieldCandidatesElemFieldUsageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCandidatesArgParametersFieldCandidatesElem
//go:noescape
func SetCandidatesArgParametersFieldCandidatesElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCandidatesArgParametersFieldCandidatesElem
//go:noescape
func SetCandidatesArgParametersFieldCandidatesElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCandidatesArgParameters
//go:noescape
func SetCandidatesArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCandidatesArgParameters
//go:noescape
func SetCandidatesArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime constof_UnderlineStyle
//go:noescape
func ConstOfUnderlineStyle(str js.Ref) uint32

//go:wasmimport plat/js/webext/input/ime store_SetCompositionArgParametersFieldSegmentsElem
//go:noescape
func SetCompositionArgParametersFieldSegmentsElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCompositionArgParametersFieldSegmentsElem
//go:noescape
func SetCompositionArgParametersFieldSegmentsElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCompositionArgParameters
//go:noescape
func SetCompositionArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCompositionArgParameters
//go:noescape
func SetCompositionArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime store_SetCursorPositionArgParameters
//go:noescape
func SetCursorPositionArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/input/ime load_SetCursorPositionArgParameters
//go:noescape
func SetCursorPositionArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/input/ime has_ClearComposition
//go:noescape
func HasFuncClearComposition() js.Ref

//go:wasmimport plat/js/webext/input/ime func_ClearComposition
//go:noescape
func FuncClearComposition(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_ClearComposition
//go:noescape
func CallClearComposition(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_ClearComposition
//go:noescape
func TryClearComposition(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_CommitText
//go:noescape
func HasFuncCommitText() js.Ref

//go:wasmimport plat/js/webext/input/ime func_CommitText
//go:noescape
func FuncCommitText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_CommitText
//go:noescape
func CallCommitText(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_CommitText
//go:noescape
func TryCommitText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_DeleteSurroundingText
//go:noescape
func HasFuncDeleteSurroundingText() js.Ref

//go:wasmimport plat/js/webext/input/ime func_DeleteSurroundingText
//go:noescape
func FuncDeleteSurroundingText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_DeleteSurroundingText
//go:noescape
func CallDeleteSurroundingText(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_DeleteSurroundingText
//go:noescape
func TryDeleteSurroundingText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HideInputView
//go:noescape
func HasFuncHideInputView() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HideInputView
//go:noescape
func FuncHideInputView(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HideInputView
//go:noescape
func CallHideInputView(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_HideInputView
//go:noescape
func TryHideInputView(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_KeyEventHandled
//go:noescape
func HasFuncKeyEventHandled() js.Ref

//go:wasmimport plat/js/webext/input/ime func_KeyEventHandled
//go:noescape
func FuncKeyEventHandled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_KeyEventHandled
//go:noescape
func CallKeyEventHandled(
	retPtr unsafe.Pointer,
	requestId js.Ref,
	response js.Ref)

//go:wasmimport plat/js/webext/input/ime try_KeyEventHandled
//go:noescape
func TryKeyEventHandled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId js.Ref,
	response js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnActivate
//go:noescape
func HasFuncOnActivate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnActivate
//go:noescape
func FuncOnActivate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnActivate
//go:noescape
func CallOnActivate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnActivate
//go:noescape
func TryOnActivate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffActivate
//go:noescape
func HasFuncOffActivate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffActivate
//go:noescape
func FuncOffActivate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffActivate
//go:noescape
func CallOffActivate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffActivate
//go:noescape
func TryOffActivate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnActivate
//go:noescape
func HasFuncHasOnActivate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnActivate
//go:noescape
func FuncHasOnActivate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnActivate
//go:noescape
func CallHasOnActivate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnActivate
//go:noescape
func TryHasOnActivate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnAssistiveWindowButtonClicked
//go:noescape
func HasFuncOnAssistiveWindowButtonClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnAssistiveWindowButtonClicked
//go:noescape
func FuncOnAssistiveWindowButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnAssistiveWindowButtonClicked
//go:noescape
func CallOnAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnAssistiveWindowButtonClicked
//go:noescape
func TryOnAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffAssistiveWindowButtonClicked
//go:noescape
func HasFuncOffAssistiveWindowButtonClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffAssistiveWindowButtonClicked
//go:noescape
func FuncOffAssistiveWindowButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffAssistiveWindowButtonClicked
//go:noescape
func CallOffAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffAssistiveWindowButtonClicked
//go:noescape
func TryOffAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnAssistiveWindowButtonClicked
//go:noescape
func HasFuncHasOnAssistiveWindowButtonClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnAssistiveWindowButtonClicked
//go:noescape
func FuncHasOnAssistiveWindowButtonClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnAssistiveWindowButtonClicked
//go:noescape
func CallHasOnAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnAssistiveWindowButtonClicked
//go:noescape
func TryHasOnAssistiveWindowButtonClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnBlur
//go:noescape
func HasFuncOnBlur() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnBlur
//go:noescape
func FuncOnBlur(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnBlur
//go:noescape
func CallOnBlur(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnBlur
//go:noescape
func TryOnBlur(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffBlur
//go:noescape
func HasFuncOffBlur() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffBlur
//go:noescape
func FuncOffBlur(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffBlur
//go:noescape
func CallOffBlur(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffBlur
//go:noescape
func TryOffBlur(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnBlur
//go:noescape
func HasFuncHasOnBlur() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnBlur
//go:noescape
func FuncHasOnBlur(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnBlur
//go:noescape
func CallHasOnBlur(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnBlur
//go:noescape
func TryHasOnBlur(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnCandidateClicked
//go:noescape
func HasFuncOnCandidateClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnCandidateClicked
//go:noescape
func FuncOnCandidateClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnCandidateClicked
//go:noescape
func CallOnCandidateClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnCandidateClicked
//go:noescape
func TryOnCandidateClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffCandidateClicked
//go:noescape
func HasFuncOffCandidateClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffCandidateClicked
//go:noescape
func FuncOffCandidateClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffCandidateClicked
//go:noescape
func CallOffCandidateClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffCandidateClicked
//go:noescape
func TryOffCandidateClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnCandidateClicked
//go:noescape
func HasFuncHasOnCandidateClicked() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnCandidateClicked
//go:noescape
func FuncHasOnCandidateClicked(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnCandidateClicked
//go:noescape
func CallHasOnCandidateClicked(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnCandidateClicked
//go:noescape
func TryHasOnCandidateClicked(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnDeactivated
//go:noescape
func HasFuncOnDeactivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnDeactivated
//go:noescape
func FuncOnDeactivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnDeactivated
//go:noescape
func CallOnDeactivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnDeactivated
//go:noescape
func TryOnDeactivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffDeactivated
//go:noescape
func HasFuncOffDeactivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffDeactivated
//go:noescape
func FuncOffDeactivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffDeactivated
//go:noescape
func CallOffDeactivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffDeactivated
//go:noescape
func TryOffDeactivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnDeactivated
//go:noescape
func HasFuncHasOnDeactivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnDeactivated
//go:noescape
func FuncHasOnDeactivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnDeactivated
//go:noescape
func CallHasOnDeactivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnDeactivated
//go:noescape
func TryHasOnDeactivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnFocus
//go:noescape
func HasFuncOnFocus() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnFocus
//go:noescape
func FuncOnFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnFocus
//go:noescape
func CallOnFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnFocus
//go:noescape
func TryOnFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffFocus
//go:noescape
func HasFuncOffFocus() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffFocus
//go:noescape
func FuncOffFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffFocus
//go:noescape
func CallOffFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffFocus
//go:noescape
func TryOffFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnFocus
//go:noescape
func HasFuncHasOnFocus() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnFocus
//go:noescape
func FuncHasOnFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnFocus
//go:noescape
func CallHasOnFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnFocus
//go:noescape
func TryHasOnFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnInputContextUpdate
//go:noescape
func HasFuncOnInputContextUpdate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnInputContextUpdate
//go:noescape
func FuncOnInputContextUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnInputContextUpdate
//go:noescape
func CallOnInputContextUpdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnInputContextUpdate
//go:noescape
func TryOnInputContextUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffInputContextUpdate
//go:noescape
func HasFuncOffInputContextUpdate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffInputContextUpdate
//go:noescape
func FuncOffInputContextUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffInputContextUpdate
//go:noescape
func CallOffInputContextUpdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffInputContextUpdate
//go:noescape
func TryOffInputContextUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnInputContextUpdate
//go:noescape
func HasFuncHasOnInputContextUpdate() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnInputContextUpdate
//go:noescape
func FuncHasOnInputContextUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnInputContextUpdate
//go:noescape
func CallHasOnInputContextUpdate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnInputContextUpdate
//go:noescape
func TryHasOnInputContextUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnKeyEvent
//go:noescape
func HasFuncOnKeyEvent() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnKeyEvent
//go:noescape
func FuncOnKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnKeyEvent
//go:noescape
func CallOnKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnKeyEvent
//go:noescape
func TryOnKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffKeyEvent
//go:noescape
func HasFuncOffKeyEvent() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffKeyEvent
//go:noescape
func FuncOffKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffKeyEvent
//go:noescape
func CallOffKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffKeyEvent
//go:noescape
func TryOffKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnKeyEvent
//go:noescape
func HasFuncHasOnKeyEvent() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnKeyEvent
//go:noescape
func FuncHasOnKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnKeyEvent
//go:noescape
func CallHasOnKeyEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnKeyEvent
//go:noescape
func TryHasOnKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnMenuItemActivated
//go:noescape
func HasFuncOnMenuItemActivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnMenuItemActivated
//go:noescape
func FuncOnMenuItemActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnMenuItemActivated
//go:noescape
func CallOnMenuItemActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnMenuItemActivated
//go:noescape
func TryOnMenuItemActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffMenuItemActivated
//go:noescape
func HasFuncOffMenuItemActivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffMenuItemActivated
//go:noescape
func FuncOffMenuItemActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffMenuItemActivated
//go:noescape
func CallOffMenuItemActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffMenuItemActivated
//go:noescape
func TryOffMenuItemActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnMenuItemActivated
//go:noescape
func HasFuncHasOnMenuItemActivated() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnMenuItemActivated
//go:noescape
func FuncHasOnMenuItemActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnMenuItemActivated
//go:noescape
func CallHasOnMenuItemActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnMenuItemActivated
//go:noescape
func TryHasOnMenuItemActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnReset
//go:noescape
func HasFuncOnReset() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnReset
//go:noescape
func FuncOnReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnReset
//go:noescape
func CallOnReset(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnReset
//go:noescape
func TryOnReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffReset
//go:noescape
func HasFuncOffReset() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffReset
//go:noescape
func FuncOffReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffReset
//go:noescape
func CallOffReset(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffReset
//go:noescape
func TryOffReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnReset
//go:noescape
func HasFuncHasOnReset() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnReset
//go:noescape
func FuncHasOnReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnReset
//go:noescape
func CallHasOnReset(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnReset
//go:noescape
func TryHasOnReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OnSurroundingTextChanged
//go:noescape
func HasFuncOnSurroundingTextChanged() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OnSurroundingTextChanged
//go:noescape
func FuncOnSurroundingTextChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OnSurroundingTextChanged
//go:noescape
func CallOnSurroundingTextChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OnSurroundingTextChanged
//go:noescape
func TryOnSurroundingTextChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_OffSurroundingTextChanged
//go:noescape
func HasFuncOffSurroundingTextChanged() js.Ref

//go:wasmimport plat/js/webext/input/ime func_OffSurroundingTextChanged
//go:noescape
func FuncOffSurroundingTextChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_OffSurroundingTextChanged
//go:noescape
func CallOffSurroundingTextChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_OffSurroundingTextChanged
//go:noescape
func TryOffSurroundingTextChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_HasOnSurroundingTextChanged
//go:noescape
func HasFuncHasOnSurroundingTextChanged() js.Ref

//go:wasmimport plat/js/webext/input/ime func_HasOnSurroundingTextChanged
//go:noescape
func FuncHasOnSurroundingTextChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_HasOnSurroundingTextChanged
//go:noescape
func CallHasOnSurroundingTextChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/input/ime try_HasOnSurroundingTextChanged
//go:noescape
func TryHasOnSurroundingTextChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SendKeyEvents
//go:noescape
func HasFuncSendKeyEvents() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SendKeyEvents
//go:noescape
func FuncSendKeyEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SendKeyEvents
//go:noescape
func CallSendKeyEvents(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SendKeyEvents
//go:noescape
func TrySendKeyEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetAssistiveWindowButtonHighlighted
//go:noescape
func HasFuncSetAssistiveWindowButtonHighlighted() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetAssistiveWindowButtonHighlighted
//go:noescape
func FuncSetAssistiveWindowButtonHighlighted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetAssistiveWindowButtonHighlighted
//go:noescape
func CallSetAssistiveWindowButtonHighlighted(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetAssistiveWindowButtonHighlighted
//go:noescape
func TrySetAssistiveWindowButtonHighlighted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetAssistiveWindowProperties
//go:noescape
func HasFuncSetAssistiveWindowProperties() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetAssistiveWindowProperties
//go:noescape
func FuncSetAssistiveWindowProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetAssistiveWindowProperties
//go:noescape
func CallSetAssistiveWindowProperties(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetAssistiveWindowProperties
//go:noescape
func TrySetAssistiveWindowProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetCandidateWindowProperties
//go:noescape
func HasFuncSetCandidateWindowProperties() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetCandidateWindowProperties
//go:noescape
func FuncSetCandidateWindowProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetCandidateWindowProperties
//go:noescape
func CallSetCandidateWindowProperties(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetCandidateWindowProperties
//go:noescape
func TrySetCandidateWindowProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetCandidates
//go:noescape
func HasFuncSetCandidates() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetCandidates
//go:noescape
func FuncSetCandidates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetCandidates
//go:noescape
func CallSetCandidates(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetCandidates
//go:noescape
func TrySetCandidates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetComposition
//go:noescape
func HasFuncSetComposition() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetComposition
//go:noescape
func FuncSetComposition(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetComposition
//go:noescape
func CallSetComposition(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetComposition
//go:noescape
func TrySetComposition(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetCursorPosition
//go:noescape
func HasFuncSetCursorPosition() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetCursorPosition
//go:noescape
func FuncSetCursorPosition(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetCursorPosition
//go:noescape
func CallSetCursorPosition(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetCursorPosition
//go:noescape
func TrySetCursorPosition(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_SetMenuItems
//go:noescape
func HasFuncSetMenuItems() js.Ref

//go:wasmimport plat/js/webext/input/ime func_SetMenuItems
//go:noescape
func FuncSetMenuItems(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_SetMenuItems
//go:noescape
func CallSetMenuItems(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_SetMenuItems
//go:noescape
func TrySetMenuItems(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/input/ime has_UpdateMenuItems
//go:noescape
func HasFuncUpdateMenuItems() js.Ref

//go:wasmimport plat/js/webext/input/ime func_UpdateMenuItems
//go:noescape
func FuncUpdateMenuItems(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime call_UpdateMenuItems
//go:noescape
func CallUpdateMenuItems(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/input/ime try_UpdateMenuItems
//go:noescape
func TryUpdateMenuItems(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)
