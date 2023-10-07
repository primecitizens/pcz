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

//go:wasmimport plat/js/webext/inputmethodprivate constof_AutoCapitalizeType
//go:noescape
func ConstOfAutoCapitalizeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate store_FinishComposingTextArgParameters
//go:noescape
func FinishComposingTextArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_FinishComposingTextArgParameters
//go:noescape
func FinishComposingTextArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate constof_FocusReason
//go:noescape
func ConstOfFocusReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate store_GetInputMethodConfigReturnType
//go:noescape
func GetInputMethodConfigReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_GetInputMethodConfigReturnType
//go:noescape
func GetInputMethodConfigReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_GetInputMethodsReturnTypeElem
//go:noescape
func GetInputMethodsReturnTypeElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_GetInputMethodsReturnTypeElem
//go:noescape
func GetInputMethodsReturnTypeElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_GetSurroundingTextReturnType
//go:noescape
func GetSurroundingTextReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_GetSurroundingTextReturnType
//go:noescape
func GetSurroundingTextReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_GetTextFieldBoundsArgParameters
//go:noescape
func GetTextFieldBoundsArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_GetTextFieldBoundsArgParameters
//go:noescape
func GetTextFieldBoundsArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_GetTextFieldBoundsReturnType
//go:noescape
func GetTextFieldBoundsReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_GetTextFieldBoundsReturnType
//go:noescape
func GetTextFieldBoundsReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate constof_InputModeType
//go:noescape
func ConstOfInputModeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate constof_InputContextType
//go:noescape
func ConstOfInputContextType(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate store_InputContext
//go:noescape
func InputContextJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_InputContext
//go:noescape
func InputContextJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_InputMethodSettingsFieldPinyinFuzzyConfig
//go:noescape
func InputMethodSettingsFieldPinyinFuzzyConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_InputMethodSettingsFieldPinyinFuzzyConfig
//go:noescape
func InputMethodSettingsFieldPinyinFuzzyConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_InputMethodSettings
//go:noescape
func InputMethodSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_InputMethodSettings
//go:noescape
func InputMethodSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate constof_MenuItemStyle
//go:noescape
func ConstOfMenuItemStyle(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate store_MenuItem
//go:noescape
func MenuItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_MenuItem
//go:noescape
func MenuItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_OnAutocorrectArgParameters
//go:noescape
func OnAutocorrectArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_OnAutocorrectArgParameters
//go:noescape
func OnAutocorrectArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_OnCaretBoundsChangedArgCaretBounds
//go:noescape
func OnCaretBoundsChangedArgCaretBoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_OnCaretBoundsChangedArgCaretBounds
//go:noescape
func OnCaretBoundsChangedArgCaretBoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate constof_UnderlineStyle
//go:noescape
func ConstOfUnderlineStyle(str js.Ref) uint32

//go:wasmimport plat/js/webext/inputmethodprivate store_SetCompositionRangeArgParametersFieldSegmentsElem
//go:noescape
func SetCompositionRangeArgParametersFieldSegmentsElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_SetCompositionRangeArgParametersFieldSegmentsElem
//go:noescape
func SetCompositionRangeArgParametersFieldSegmentsElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate store_SetCompositionRangeArgParameters
//go:noescape
func SetCompositionRangeArgParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate load_SetCompositionRangeArgParameters
//go:noescape
func SetCompositionRangeArgParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate has_AddWordToDictionary
//go:noescape
func HasFuncAddWordToDictionary() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_AddWordToDictionary
//go:noescape
func FuncAddWordToDictionary(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_AddWordToDictionary
//go:noescape
func CallAddWordToDictionary(
	retPtr unsafe.Pointer,
	word js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_AddWordToDictionary
//go:noescape
func TryAddWordToDictionary(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	word js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_FetchAllDictionaryWords
//go:noescape
func HasFuncFetchAllDictionaryWords() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_FetchAllDictionaryWords
//go:noescape
func FuncFetchAllDictionaryWords(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_FetchAllDictionaryWords
//go:noescape
func CallFetchAllDictionaryWords(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_FetchAllDictionaryWords
//go:noescape
func TryFetchAllDictionaryWords(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_FinishComposingText
//go:noescape
func HasFuncFinishComposingText() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_FinishComposingText
//go:noescape
func FuncFinishComposingText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_FinishComposingText
//go:noescape
func CallFinishComposingText(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_FinishComposingText
//go:noescape
func TryFinishComposingText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetCurrentInputMethod
//go:noescape
func HasFuncGetCurrentInputMethod() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetCurrentInputMethod
//go:noescape
func FuncGetCurrentInputMethod(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetCurrentInputMethod
//go:noescape
func CallGetCurrentInputMethod(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetCurrentInputMethod
//go:noescape
func TryGetCurrentInputMethod(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetInputMethodConfig
//go:noescape
func HasFuncGetInputMethodConfig() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetInputMethodConfig
//go:noescape
func FuncGetInputMethodConfig(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetInputMethodConfig
//go:noescape
func CallGetInputMethodConfig(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetInputMethodConfig
//go:noescape
func TryGetInputMethodConfig(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetInputMethods
//go:noescape
func HasFuncGetInputMethods() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetInputMethods
//go:noescape
func FuncGetInputMethods(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetInputMethods
//go:noescape
func CallGetInputMethods(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetInputMethods
//go:noescape
func TryGetInputMethods(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetSettings
//go:noescape
func HasFuncGetSettings() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetSettings
//go:noescape
func FuncGetSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetSettings
//go:noescape
func CallGetSettings(
	retPtr unsafe.Pointer,
	engineID js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetSettings
//go:noescape
func TryGetSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	engineID js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetSurroundingText
//go:noescape
func HasFuncGetSurroundingText() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetSurroundingText
//go:noescape
func FuncGetSurroundingText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetSurroundingText
//go:noescape
func CallGetSurroundingText(
	retPtr unsafe.Pointer,
	beforeLength float64,
	afterLength float64)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetSurroundingText
//go:noescape
func TryGetSurroundingText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	beforeLength float64,
	afterLength float64) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_GetTextFieldBounds
//go:noescape
func HasFuncGetTextFieldBounds() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_GetTextFieldBounds
//go:noescape
func FuncGetTextFieldBounds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_GetTextFieldBounds
//go:noescape
func CallGetTextFieldBounds(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_GetTextFieldBounds
//go:noescape
func TryGetTextFieldBounds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HideInputView
//go:noescape
func HasFuncHideInputView() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HideInputView
//go:noescape
func FuncHideInputView(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HideInputView
//go:noescape
func CallHideInputView(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_HideInputView
//go:noescape
func TryHideInputView(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_NotifyInputMethodReadyForTesting
//go:noescape
func HasFuncNotifyInputMethodReadyForTesting() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_NotifyInputMethodReadyForTesting
//go:noescape
func FuncNotifyInputMethodReadyForTesting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_NotifyInputMethodReadyForTesting
//go:noescape
func CallNotifyInputMethodReadyForTesting(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_NotifyInputMethodReadyForTesting
//go:noescape
func TryNotifyInputMethodReadyForTesting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnAutocorrect
//go:noescape
func HasFuncOnAutocorrect() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnAutocorrect
//go:noescape
func FuncOnAutocorrect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnAutocorrect
//go:noescape
func CallOnAutocorrect(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnAutocorrect
//go:noescape
func TryOnAutocorrect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnCaretBoundsChanged
//go:noescape
func HasFuncOnCaretBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnCaretBoundsChanged
//go:noescape
func FuncOnCaretBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnCaretBoundsChanged
//go:noescape
func CallOnCaretBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnCaretBoundsChanged
//go:noescape
func TryOnCaretBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffCaretBoundsChanged
//go:noescape
func HasFuncOffCaretBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffCaretBoundsChanged
//go:noescape
func FuncOffCaretBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffCaretBoundsChanged
//go:noescape
func CallOffCaretBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffCaretBoundsChanged
//go:noescape
func TryOffCaretBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnCaretBoundsChanged
//go:noescape
func HasFuncHasOnCaretBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnCaretBoundsChanged
//go:noescape
func FuncHasOnCaretBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnCaretBoundsChanged
//go:noescape
func CallHasOnCaretBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnCaretBoundsChanged
//go:noescape
func TryHasOnCaretBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnDictionaryChanged
//go:noescape
func HasFuncOnDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnDictionaryChanged
//go:noescape
func FuncOnDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnDictionaryChanged
//go:noescape
func CallOnDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnDictionaryChanged
//go:noescape
func TryOnDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffDictionaryChanged
//go:noescape
func HasFuncOffDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffDictionaryChanged
//go:noescape
func FuncOffDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffDictionaryChanged
//go:noescape
func CallOffDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffDictionaryChanged
//go:noescape
func TryOffDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnDictionaryChanged
//go:noescape
func HasFuncHasOnDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnDictionaryChanged
//go:noescape
func FuncHasOnDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnDictionaryChanged
//go:noescape
func CallHasOnDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnDictionaryChanged
//go:noescape
func TryHasOnDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnDictionaryLoaded
//go:noescape
func HasFuncOnDictionaryLoaded() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnDictionaryLoaded
//go:noescape
func FuncOnDictionaryLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnDictionaryLoaded
//go:noescape
func CallOnDictionaryLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnDictionaryLoaded
//go:noescape
func TryOnDictionaryLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffDictionaryLoaded
//go:noescape
func HasFuncOffDictionaryLoaded() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffDictionaryLoaded
//go:noescape
func FuncOffDictionaryLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffDictionaryLoaded
//go:noescape
func CallOffDictionaryLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffDictionaryLoaded
//go:noescape
func TryOffDictionaryLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnDictionaryLoaded
//go:noescape
func HasFuncHasOnDictionaryLoaded() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnDictionaryLoaded
//go:noescape
func FuncHasOnDictionaryLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnDictionaryLoaded
//go:noescape
func CallHasOnDictionaryLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnDictionaryLoaded
//go:noescape
func TryHasOnDictionaryLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnFocus
//go:noescape
func HasFuncOnFocus() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnFocus
//go:noescape
func FuncOnFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnFocus
//go:noescape
func CallOnFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnFocus
//go:noescape
func TryOnFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffFocus
//go:noescape
func HasFuncOffFocus() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffFocus
//go:noescape
func FuncOffFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffFocus
//go:noescape
func CallOffFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffFocus
//go:noescape
func TryOffFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnFocus
//go:noescape
func HasFuncHasOnFocus() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnFocus
//go:noescape
func FuncHasOnFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnFocus
//go:noescape
func CallHasOnFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnFocus
//go:noescape
func TryHasOnFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnImeMenuActivationChanged
//go:noescape
func HasFuncOnImeMenuActivationChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnImeMenuActivationChanged
//go:noescape
func FuncOnImeMenuActivationChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnImeMenuActivationChanged
//go:noescape
func CallOnImeMenuActivationChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnImeMenuActivationChanged
//go:noescape
func TryOnImeMenuActivationChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffImeMenuActivationChanged
//go:noescape
func HasFuncOffImeMenuActivationChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffImeMenuActivationChanged
//go:noescape
func FuncOffImeMenuActivationChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffImeMenuActivationChanged
//go:noescape
func CallOffImeMenuActivationChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffImeMenuActivationChanged
//go:noescape
func TryOffImeMenuActivationChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnImeMenuActivationChanged
//go:noescape
func HasFuncHasOnImeMenuActivationChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnImeMenuActivationChanged
//go:noescape
func FuncHasOnImeMenuActivationChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnImeMenuActivationChanged
//go:noescape
func CallHasOnImeMenuActivationChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnImeMenuActivationChanged
//go:noescape
func TryHasOnImeMenuActivationChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnImeMenuItemsChanged
//go:noescape
func HasFuncOnImeMenuItemsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnImeMenuItemsChanged
//go:noescape
func FuncOnImeMenuItemsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnImeMenuItemsChanged
//go:noescape
func CallOnImeMenuItemsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnImeMenuItemsChanged
//go:noescape
func TryOnImeMenuItemsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffImeMenuItemsChanged
//go:noescape
func HasFuncOffImeMenuItemsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffImeMenuItemsChanged
//go:noescape
func FuncOffImeMenuItemsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffImeMenuItemsChanged
//go:noescape
func CallOffImeMenuItemsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffImeMenuItemsChanged
//go:noescape
func TryOffImeMenuItemsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnImeMenuItemsChanged
//go:noescape
func HasFuncHasOnImeMenuItemsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnImeMenuItemsChanged
//go:noescape
func FuncHasOnImeMenuItemsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnImeMenuItemsChanged
//go:noescape
func CallHasOnImeMenuItemsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnImeMenuItemsChanged
//go:noescape
func TryHasOnImeMenuItemsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnImeMenuListChanged
//go:noescape
func HasFuncOnImeMenuListChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnImeMenuListChanged
//go:noescape
func FuncOnImeMenuListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnImeMenuListChanged
//go:noescape
func CallOnImeMenuListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnImeMenuListChanged
//go:noescape
func TryOnImeMenuListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffImeMenuListChanged
//go:noescape
func HasFuncOffImeMenuListChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffImeMenuListChanged
//go:noescape
func FuncOffImeMenuListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffImeMenuListChanged
//go:noescape
func CallOffImeMenuListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffImeMenuListChanged
//go:noescape
func TryOffImeMenuListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnImeMenuListChanged
//go:noescape
func HasFuncHasOnImeMenuListChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnImeMenuListChanged
//go:noescape
func FuncHasOnImeMenuListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnImeMenuListChanged
//go:noescape
func CallHasOnImeMenuListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnImeMenuListChanged
//go:noescape
func TryHasOnImeMenuListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnInputMethodOptionsChanged
//go:noescape
func HasFuncOnInputMethodOptionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnInputMethodOptionsChanged
//go:noescape
func FuncOnInputMethodOptionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnInputMethodOptionsChanged
//go:noescape
func CallOnInputMethodOptionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnInputMethodOptionsChanged
//go:noescape
func TryOnInputMethodOptionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffInputMethodOptionsChanged
//go:noescape
func HasFuncOffInputMethodOptionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffInputMethodOptionsChanged
//go:noescape
func FuncOffInputMethodOptionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffInputMethodOptionsChanged
//go:noescape
func CallOffInputMethodOptionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffInputMethodOptionsChanged
//go:noescape
func TryOffInputMethodOptionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnInputMethodOptionsChanged
//go:noescape
func HasFuncHasOnInputMethodOptionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnInputMethodOptionsChanged
//go:noescape
func FuncHasOnInputMethodOptionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnInputMethodOptionsChanged
//go:noescape
func CallHasOnInputMethodOptionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnInputMethodOptionsChanged
//go:noescape
func TryHasOnInputMethodOptionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnScreenProjectionChanged
//go:noescape
func HasFuncOnScreenProjectionChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnScreenProjectionChanged
//go:noescape
func FuncOnScreenProjectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnScreenProjectionChanged
//go:noescape
func CallOnScreenProjectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnScreenProjectionChanged
//go:noescape
func TryOnScreenProjectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffScreenProjectionChanged
//go:noescape
func HasFuncOffScreenProjectionChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffScreenProjectionChanged
//go:noescape
func FuncOffScreenProjectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffScreenProjectionChanged
//go:noescape
func CallOffScreenProjectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffScreenProjectionChanged
//go:noescape
func TryOffScreenProjectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnScreenProjectionChanged
//go:noescape
func HasFuncHasOnScreenProjectionChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnScreenProjectionChanged
//go:noescape
func FuncHasOnScreenProjectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnScreenProjectionChanged
//go:noescape
func CallHasOnScreenProjectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnScreenProjectionChanged
//go:noescape
func TryHasOnScreenProjectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnSettingsChanged
//go:noescape
func HasFuncOnSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnSettingsChanged
//go:noescape
func FuncOnSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnSettingsChanged
//go:noescape
func CallOnSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnSettingsChanged
//go:noescape
func TryOnSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffSettingsChanged
//go:noescape
func HasFuncOffSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffSettingsChanged
//go:noescape
func FuncOffSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffSettingsChanged
//go:noescape
func CallOffSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffSettingsChanged
//go:noescape
func TryOffSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnSettingsChanged
//go:noescape
func HasFuncHasOnSettingsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnSettingsChanged
//go:noescape
func FuncHasOnSettingsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnSettingsChanged
//go:noescape
func CallHasOnSettingsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnSettingsChanged
//go:noescape
func TryHasOnSettingsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnSuggestionsChanged
//go:noescape
func HasFuncOnSuggestionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnSuggestionsChanged
//go:noescape
func FuncOnSuggestionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnSuggestionsChanged
//go:noescape
func CallOnSuggestionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnSuggestionsChanged
//go:noescape
func TryOnSuggestionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffSuggestionsChanged
//go:noescape
func HasFuncOffSuggestionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffSuggestionsChanged
//go:noescape
func FuncOffSuggestionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffSuggestionsChanged
//go:noescape
func CallOffSuggestionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffSuggestionsChanged
//go:noescape
func TryOffSuggestionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnSuggestionsChanged
//go:noescape
func HasFuncHasOnSuggestionsChanged() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnSuggestionsChanged
//go:noescape
func FuncHasOnSuggestionsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnSuggestionsChanged
//go:noescape
func CallHasOnSuggestionsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnSuggestionsChanged
//go:noescape
func TryHasOnSuggestionsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OnTouch
//go:noescape
func HasFuncOnTouch() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OnTouch
//go:noescape
func FuncOnTouch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OnTouch
//go:noescape
func CallOnTouch(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OnTouch
//go:noescape
func TryOnTouch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OffTouch
//go:noescape
func HasFuncOffTouch() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OffTouch
//go:noescape
func FuncOffTouch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OffTouch
//go:noescape
func CallOffTouch(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OffTouch
//go:noescape
func TryOffTouch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_HasOnTouch
//go:noescape
func HasFuncHasOnTouch() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_HasOnTouch
//go:noescape
func FuncHasOnTouch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_HasOnTouch
//go:noescape
func CallHasOnTouch(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_HasOnTouch
//go:noescape
func TryHasOnTouch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_OpenOptionsPage
//go:noescape
func HasFuncOpenOptionsPage() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_OpenOptionsPage
//go:noescape
func FuncOpenOptionsPage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_OpenOptionsPage
//go:noescape
func CallOpenOptionsPage(
	retPtr unsafe.Pointer,
	inputMethodId js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_OpenOptionsPage
//go:noescape
func TryOpenOptionsPage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	inputMethodId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_Reset
//go:noescape
func HasFuncReset() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_Reset
//go:noescape
func FuncReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_Reset
//go:noescape
func CallReset(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_Reset
//go:noescape
func TryReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_SetCompositionRange
//go:noescape
func HasFuncSetCompositionRange() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_SetCompositionRange
//go:noescape
func FuncSetCompositionRange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_SetCompositionRange
//go:noescape
func CallSetCompositionRange(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_SetCompositionRange
//go:noescape
func TrySetCompositionRange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_SetCurrentInputMethod
//go:noescape
func HasFuncSetCurrentInputMethod() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_SetCurrentInputMethod
//go:noescape
func FuncSetCurrentInputMethod(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_SetCurrentInputMethod
//go:noescape
func CallSetCurrentInputMethod(
	retPtr unsafe.Pointer,
	inputMethodId js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_SetCurrentInputMethod
//go:noescape
func TrySetCurrentInputMethod(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	inputMethodId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_SetSettings
//go:noescape
func HasFuncSetSettings() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_SetSettings
//go:noescape
func FuncSetSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_SetSettings
//go:noescape
func CallSetSettings(
	retPtr unsafe.Pointer,
	engineID js.Ref,
	settings unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_SetSettings
//go:noescape
func TrySetSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	engineID js.Ref,
	settings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_SetXkbLayout
//go:noescape
func HasFuncSetXkbLayout() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_SetXkbLayout
//go:noescape
func FuncSetXkbLayout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_SetXkbLayout
//go:noescape
func CallSetXkbLayout(
	retPtr unsafe.Pointer,
	xkb_name js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate try_SetXkbLayout
//go:noescape
func TrySetXkbLayout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	xkb_name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_ShowInputView
//go:noescape
func HasFuncShowInputView() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_ShowInputView
//go:noescape
func FuncShowInputView(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_ShowInputView
//go:noescape
func CallShowInputView(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_ShowInputView
//go:noescape
func TryShowInputView(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/inputmethodprivate has_SwitchToLastUsedInputMethod
//go:noescape
func HasFuncSwitchToLastUsedInputMethod() js.Ref

//go:wasmimport plat/js/webext/inputmethodprivate func_SwitchToLastUsedInputMethod
//go:noescape
func FuncSwitchToLastUsedInputMethod(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate call_SwitchToLastUsedInputMethod
//go:noescape
func CallSwitchToLastUsedInputMethod(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/inputmethodprivate try_SwitchToLastUsedInputMethod
//go:noescape
func TrySwitchToLastUsedInputMethod(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
