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

//go:wasmimport plat/js/webext/omnibox constof_DescriptionStyleType
//go:noescape
func ConstOfDescriptionStyleType(str js.Ref) uint32

//go:wasmimport plat/js/webext/omnibox store_MatchClassification
//go:noescape
func MatchClassificationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/omnibox load_MatchClassification
//go:noescape
func MatchClassificationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/omnibox store_DefaultSuggestResult
//go:noescape
func DefaultSuggestResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/omnibox load_DefaultSuggestResult
//go:noescape
func DefaultSuggestResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/omnibox store_SuggestResult
//go:noescape
func SuggestResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/omnibox load_SuggestResult
//go:noescape
func SuggestResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/omnibox constof_OnInputEnteredDisposition
//go:noescape
func ConstOfOnInputEnteredDisposition(str js.Ref) uint32

//go:wasmimport plat/js/webext/omnibox has_OnDeleteSuggestion
//go:noescape
func HasFuncOnDeleteSuggestion() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OnDeleteSuggestion
//go:noescape
func FuncOnDeleteSuggestion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OnDeleteSuggestion
//go:noescape
func CallOnDeleteSuggestion(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OnDeleteSuggestion
//go:noescape
func TryOnDeleteSuggestion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OffDeleteSuggestion
//go:noescape
func HasFuncOffDeleteSuggestion() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OffDeleteSuggestion
//go:noescape
func FuncOffDeleteSuggestion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OffDeleteSuggestion
//go:noescape
func CallOffDeleteSuggestion(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OffDeleteSuggestion
//go:noescape
func TryOffDeleteSuggestion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_HasOnDeleteSuggestion
//go:noescape
func HasFuncHasOnDeleteSuggestion() js.Ref

//go:wasmimport plat/js/webext/omnibox func_HasOnDeleteSuggestion
//go:noescape
func FuncHasOnDeleteSuggestion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_HasOnDeleteSuggestion
//go:noescape
func CallHasOnDeleteSuggestion(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_HasOnDeleteSuggestion
//go:noescape
func TryHasOnDeleteSuggestion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OnInputCancelled
//go:noescape
func HasFuncOnInputCancelled() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OnInputCancelled
//go:noescape
func FuncOnInputCancelled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OnInputCancelled
//go:noescape
func CallOnInputCancelled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OnInputCancelled
//go:noescape
func TryOnInputCancelled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OffInputCancelled
//go:noescape
func HasFuncOffInputCancelled() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OffInputCancelled
//go:noescape
func FuncOffInputCancelled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OffInputCancelled
//go:noescape
func CallOffInputCancelled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OffInputCancelled
//go:noescape
func TryOffInputCancelled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_HasOnInputCancelled
//go:noescape
func HasFuncHasOnInputCancelled() js.Ref

//go:wasmimport plat/js/webext/omnibox func_HasOnInputCancelled
//go:noescape
func FuncHasOnInputCancelled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_HasOnInputCancelled
//go:noescape
func CallHasOnInputCancelled(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_HasOnInputCancelled
//go:noescape
func TryHasOnInputCancelled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OnInputChanged
//go:noescape
func HasFuncOnInputChanged() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OnInputChanged
//go:noescape
func FuncOnInputChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OnInputChanged
//go:noescape
func CallOnInputChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OnInputChanged
//go:noescape
func TryOnInputChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OffInputChanged
//go:noescape
func HasFuncOffInputChanged() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OffInputChanged
//go:noescape
func FuncOffInputChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OffInputChanged
//go:noescape
func CallOffInputChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OffInputChanged
//go:noescape
func TryOffInputChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_HasOnInputChanged
//go:noescape
func HasFuncHasOnInputChanged() js.Ref

//go:wasmimport plat/js/webext/omnibox func_HasOnInputChanged
//go:noescape
func FuncHasOnInputChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_HasOnInputChanged
//go:noescape
func CallHasOnInputChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_HasOnInputChanged
//go:noescape
func TryHasOnInputChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OnInputEntered
//go:noescape
func HasFuncOnInputEntered() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OnInputEntered
//go:noescape
func FuncOnInputEntered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OnInputEntered
//go:noescape
func CallOnInputEntered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OnInputEntered
//go:noescape
func TryOnInputEntered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OffInputEntered
//go:noescape
func HasFuncOffInputEntered() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OffInputEntered
//go:noescape
func FuncOffInputEntered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OffInputEntered
//go:noescape
func CallOffInputEntered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OffInputEntered
//go:noescape
func TryOffInputEntered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_HasOnInputEntered
//go:noescape
func HasFuncHasOnInputEntered() js.Ref

//go:wasmimport plat/js/webext/omnibox func_HasOnInputEntered
//go:noescape
func FuncHasOnInputEntered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_HasOnInputEntered
//go:noescape
func CallHasOnInputEntered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_HasOnInputEntered
//go:noescape
func TryHasOnInputEntered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OnInputStarted
//go:noescape
func HasFuncOnInputStarted() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OnInputStarted
//go:noescape
func FuncOnInputStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OnInputStarted
//go:noescape
func CallOnInputStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OnInputStarted
//go:noescape
func TryOnInputStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_OffInputStarted
//go:noescape
func HasFuncOffInputStarted() js.Ref

//go:wasmimport plat/js/webext/omnibox func_OffInputStarted
//go:noescape
func FuncOffInputStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_OffInputStarted
//go:noescape
func CallOffInputStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_OffInputStarted
//go:noescape
func TryOffInputStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_HasOnInputStarted
//go:noescape
func HasFuncHasOnInputStarted() js.Ref

//go:wasmimport plat/js/webext/omnibox func_HasOnInputStarted
//go:noescape
func FuncHasOnInputStarted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_HasOnInputStarted
//go:noescape
func CallHasOnInputStarted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/omnibox try_HasOnInputStarted
//go:noescape
func TryHasOnInputStarted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_SendSuggestions
//go:noescape
func HasFuncSendSuggestions() js.Ref

//go:wasmimport plat/js/webext/omnibox func_SendSuggestions
//go:noescape
func FuncSendSuggestions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_SendSuggestions
//go:noescape
func CallSendSuggestions(
	retPtr unsafe.Pointer,
	requestId float64,
	suggestResults js.Ref)

//go:wasmimport plat/js/webext/omnibox try_SendSuggestions
//go:noescape
func TrySendSuggestions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId float64,
	suggestResults js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/omnibox has_SetDefaultSuggestion
//go:noescape
func HasFuncSetDefaultSuggestion() js.Ref

//go:wasmimport plat/js/webext/omnibox func_SetDefaultSuggestion
//go:noescape
func FuncSetDefaultSuggestion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox call_SetDefaultSuggestion
//go:noescape
func CallSetDefaultSuggestion(
	retPtr unsafe.Pointer,
	suggestion unsafe.Pointer)

//go:wasmimport plat/js/webext/omnibox try_SetDefaultSuggestion
//go:noescape
func TrySetDefaultSuggestion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	suggestion unsafe.Pointer) (ok js.Ref)
