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

//go:wasmimport plat/js/webext/languagesettingsprivate store_InputMethod
//go:noescape
func InputMethodJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate load_InputMethod
//go:noescape
func InputMethodJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate store_InputMethodLists
//go:noescape
func InputMethodListsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate load_InputMethodLists
//go:noescape
func InputMethodListsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate store_Language
//go:noescape
func LanguageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate load_Language
//go:noescape
func LanguageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate store_SpellcheckDictionaryStatus
//go:noescape
func SpellcheckDictionaryStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate load_SpellcheckDictionaryStatus
//go:noescape
func SpellcheckDictionaryStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate constof_MoveType
//go:noescape
func ConstOfMoveType(str js.Ref) uint32

//go:wasmimport plat/js/webext/languagesettingsprivate has_AddInputMethod
//go:noescape
func HasFuncAddInputMethod() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_AddInputMethod
//go:noescape
func FuncAddInputMethod(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_AddInputMethod
//go:noescape
func CallAddInputMethod(
	retPtr unsafe.Pointer,
	inputMethodId js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_AddInputMethod
//go:noescape
func TryAddInputMethod(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	inputMethodId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_AddSpellcheckWord
//go:noescape
func HasFuncAddSpellcheckWord() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_AddSpellcheckWord
//go:noescape
func FuncAddSpellcheckWord(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_AddSpellcheckWord
//go:noescape
func CallAddSpellcheckWord(
	retPtr unsafe.Pointer,
	word js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_AddSpellcheckWord
//go:noescape
func TryAddSpellcheckWord(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	word js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_DisableLanguage
//go:noescape
func HasFuncDisableLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_DisableLanguage
//go:noescape
func FuncDisableLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_DisableLanguage
//go:noescape
func CallDisableLanguage(
	retPtr unsafe.Pointer,
	languageCode js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_DisableLanguage
//go:noescape
func TryDisableLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_EnableLanguage
//go:noescape
func HasFuncEnableLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_EnableLanguage
//go:noescape
func FuncEnableLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_EnableLanguage
//go:noescape
func CallEnableLanguage(
	retPtr unsafe.Pointer,
	languageCode js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_EnableLanguage
//go:noescape
func TryEnableLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetAlwaysTranslateLanguages
//go:noescape
func HasFuncGetAlwaysTranslateLanguages() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetAlwaysTranslateLanguages
//go:noescape
func FuncGetAlwaysTranslateLanguages(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetAlwaysTranslateLanguages
//go:noescape
func CallGetAlwaysTranslateLanguages(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetAlwaysTranslateLanguages
//go:noescape
func TryGetAlwaysTranslateLanguages(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetInputMethodLists
//go:noescape
func HasFuncGetInputMethodLists() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetInputMethodLists
//go:noescape
func FuncGetInputMethodLists(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetInputMethodLists
//go:noescape
func CallGetInputMethodLists(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetInputMethodLists
//go:noescape
func TryGetInputMethodLists(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetLanguageList
//go:noescape
func HasFuncGetLanguageList() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetLanguageList
//go:noescape
func FuncGetLanguageList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetLanguageList
//go:noescape
func CallGetLanguageList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetLanguageList
//go:noescape
func TryGetLanguageList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetNeverTranslateLanguages
//go:noescape
func HasFuncGetNeverTranslateLanguages() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetNeverTranslateLanguages
//go:noescape
func FuncGetNeverTranslateLanguages(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetNeverTranslateLanguages
//go:noescape
func CallGetNeverTranslateLanguages(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetNeverTranslateLanguages
//go:noescape
func TryGetNeverTranslateLanguages(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetSpellcheckDictionaryStatuses
//go:noescape
func HasFuncGetSpellcheckDictionaryStatuses() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetSpellcheckDictionaryStatuses
//go:noescape
func FuncGetSpellcheckDictionaryStatuses(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetSpellcheckDictionaryStatuses
//go:noescape
func CallGetSpellcheckDictionaryStatuses(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetSpellcheckDictionaryStatuses
//go:noescape
func TryGetSpellcheckDictionaryStatuses(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetSpellcheckWords
//go:noescape
func HasFuncGetSpellcheckWords() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetSpellcheckWords
//go:noescape
func FuncGetSpellcheckWords(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetSpellcheckWords
//go:noescape
func CallGetSpellcheckWords(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetSpellcheckWords
//go:noescape
func TryGetSpellcheckWords(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_GetTranslateTargetLanguage
//go:noescape
func HasFuncGetTranslateTargetLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_GetTranslateTargetLanguage
//go:noescape
func FuncGetTranslateTargetLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_GetTranslateTargetLanguage
//go:noescape
func CallGetTranslateTargetLanguage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate try_GetTranslateTargetLanguage
//go:noescape
func TryGetTranslateTargetLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_MoveLanguage
//go:noescape
func HasFuncMoveLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_MoveLanguage
//go:noescape
func FuncMoveLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_MoveLanguage
//go:noescape
func CallMoveLanguage(
	retPtr unsafe.Pointer,
	languageCode js.Ref,
	moveType uint32)

//go:wasmimport plat/js/webext/languagesettingsprivate try_MoveLanguage
//go:noescape
func TryMoveLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref,
	moveType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OnCustomDictionaryChanged
//go:noescape
func HasFuncOnCustomDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OnCustomDictionaryChanged
//go:noescape
func FuncOnCustomDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OnCustomDictionaryChanged
//go:noescape
func CallOnCustomDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OnCustomDictionaryChanged
//go:noescape
func TryOnCustomDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OffCustomDictionaryChanged
//go:noescape
func HasFuncOffCustomDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OffCustomDictionaryChanged
//go:noescape
func FuncOffCustomDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OffCustomDictionaryChanged
//go:noescape
func CallOffCustomDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OffCustomDictionaryChanged
//go:noescape
func TryOffCustomDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_HasOnCustomDictionaryChanged
//go:noescape
func HasFuncHasOnCustomDictionaryChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_HasOnCustomDictionaryChanged
//go:noescape
func FuncHasOnCustomDictionaryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_HasOnCustomDictionaryChanged
//go:noescape
func CallHasOnCustomDictionaryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_HasOnCustomDictionaryChanged
//go:noescape
func TryHasOnCustomDictionaryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OnInputMethodAdded
//go:noescape
func HasFuncOnInputMethodAdded() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OnInputMethodAdded
//go:noescape
func FuncOnInputMethodAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OnInputMethodAdded
//go:noescape
func CallOnInputMethodAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OnInputMethodAdded
//go:noescape
func TryOnInputMethodAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OffInputMethodAdded
//go:noescape
func HasFuncOffInputMethodAdded() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OffInputMethodAdded
//go:noescape
func FuncOffInputMethodAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OffInputMethodAdded
//go:noescape
func CallOffInputMethodAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OffInputMethodAdded
//go:noescape
func TryOffInputMethodAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_HasOnInputMethodAdded
//go:noescape
func HasFuncHasOnInputMethodAdded() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_HasOnInputMethodAdded
//go:noescape
func FuncHasOnInputMethodAdded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_HasOnInputMethodAdded
//go:noescape
func CallHasOnInputMethodAdded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_HasOnInputMethodAdded
//go:noescape
func TryHasOnInputMethodAdded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OnInputMethodRemoved
//go:noescape
func HasFuncOnInputMethodRemoved() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OnInputMethodRemoved
//go:noescape
func FuncOnInputMethodRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OnInputMethodRemoved
//go:noescape
func CallOnInputMethodRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OnInputMethodRemoved
//go:noescape
func TryOnInputMethodRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OffInputMethodRemoved
//go:noescape
func HasFuncOffInputMethodRemoved() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OffInputMethodRemoved
//go:noescape
func FuncOffInputMethodRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OffInputMethodRemoved
//go:noescape
func CallOffInputMethodRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OffInputMethodRemoved
//go:noescape
func TryOffInputMethodRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_HasOnInputMethodRemoved
//go:noescape
func HasFuncHasOnInputMethodRemoved() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_HasOnInputMethodRemoved
//go:noescape
func FuncHasOnInputMethodRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_HasOnInputMethodRemoved
//go:noescape
func CallHasOnInputMethodRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_HasOnInputMethodRemoved
//go:noescape
func TryHasOnInputMethodRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OnSpellcheckDictionariesChanged
//go:noescape
func HasFuncOnSpellcheckDictionariesChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OnSpellcheckDictionariesChanged
//go:noescape
func FuncOnSpellcheckDictionariesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OnSpellcheckDictionariesChanged
//go:noescape
func CallOnSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OnSpellcheckDictionariesChanged
//go:noescape
func TryOnSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_OffSpellcheckDictionariesChanged
//go:noescape
func HasFuncOffSpellcheckDictionariesChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_OffSpellcheckDictionariesChanged
//go:noescape
func FuncOffSpellcheckDictionariesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_OffSpellcheckDictionariesChanged
//go:noescape
func CallOffSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_OffSpellcheckDictionariesChanged
//go:noescape
func TryOffSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_HasOnSpellcheckDictionariesChanged
//go:noescape
func HasFuncHasOnSpellcheckDictionariesChanged() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_HasOnSpellcheckDictionariesChanged
//go:noescape
func FuncHasOnSpellcheckDictionariesChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_HasOnSpellcheckDictionariesChanged
//go:noescape
func CallHasOnSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_HasOnSpellcheckDictionariesChanged
//go:noescape
func TryHasOnSpellcheckDictionariesChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_RemoveInputMethod
//go:noescape
func HasFuncRemoveInputMethod() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_RemoveInputMethod
//go:noescape
func FuncRemoveInputMethod(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_RemoveInputMethod
//go:noescape
func CallRemoveInputMethod(
	retPtr unsafe.Pointer,
	inputMethodId js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_RemoveInputMethod
//go:noescape
func TryRemoveInputMethod(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	inputMethodId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_RemoveSpellcheckWord
//go:noescape
func HasFuncRemoveSpellcheckWord() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_RemoveSpellcheckWord
//go:noescape
func FuncRemoveSpellcheckWord(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_RemoveSpellcheckWord
//go:noescape
func CallRemoveSpellcheckWord(
	retPtr unsafe.Pointer,
	word js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_RemoveSpellcheckWord
//go:noescape
func TryRemoveSpellcheckWord(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	word js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_RetryDownloadDictionary
//go:noescape
func HasFuncRetryDownloadDictionary() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_RetryDownloadDictionary
//go:noescape
func FuncRetryDownloadDictionary(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_RetryDownloadDictionary
//go:noescape
func CallRetryDownloadDictionary(
	retPtr unsafe.Pointer,
	languageCode js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_RetryDownloadDictionary
//go:noescape
func TryRetryDownloadDictionary(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_SetEnableTranslationForLanguage
//go:noescape
func HasFuncSetEnableTranslationForLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_SetEnableTranslationForLanguage
//go:noescape
func FuncSetEnableTranslationForLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_SetEnableTranslationForLanguage
//go:noescape
func CallSetEnableTranslationForLanguage(
	retPtr unsafe.Pointer,
	languageCode js.Ref,
	enable js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_SetEnableTranslationForLanguage
//go:noescape
func TrySetEnableTranslationForLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref,
	enable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_SetLanguageAlwaysTranslateState
//go:noescape
func HasFuncSetLanguageAlwaysTranslateState() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_SetLanguageAlwaysTranslateState
//go:noescape
func FuncSetLanguageAlwaysTranslateState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_SetLanguageAlwaysTranslateState
//go:noescape
func CallSetLanguageAlwaysTranslateState(
	retPtr unsafe.Pointer,
	languageCode js.Ref,
	alwaysTranslate js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_SetLanguageAlwaysTranslateState
//go:noescape
func TrySetLanguageAlwaysTranslateState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref,
	alwaysTranslate js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate has_SetTranslateTargetLanguage
//go:noescape
func HasFuncSetTranslateTargetLanguage() js.Ref

//go:wasmimport plat/js/webext/languagesettingsprivate func_SetTranslateTargetLanguage
//go:noescape
func FuncSetTranslateTargetLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/languagesettingsprivate call_SetTranslateTargetLanguage
//go:noescape
func CallSetTranslateTargetLanguage(
	retPtr unsafe.Pointer,
	languageCode js.Ref)

//go:wasmimport plat/js/webext/languagesettingsprivate try_SetTranslateTargetLanguage
//go:noescape
func TrySetTranslateTargetLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	languageCode js.Ref) (ok js.Ref)
