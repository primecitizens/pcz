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

//go:wasmimport plat/js/webext/i18n store_DetectLanguageReturnTypeFieldLanguagesElem
//go:noescape
func DetectLanguageReturnTypeFieldLanguagesElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/i18n load_DetectLanguageReturnTypeFieldLanguagesElem
//go:noescape
func DetectLanguageReturnTypeFieldLanguagesElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/i18n store_DetectLanguageReturnType
//go:noescape
func DetectLanguageReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/i18n load_DetectLanguageReturnType
//go:noescape
func DetectLanguageReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/i18n store_GetMessageArgOptions
//go:noescape
func GetMessageArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/i18n load_GetMessageArgOptions
//go:noescape
func GetMessageArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/i18n has_DetectLanguage
//go:noescape
func HasFuncDetectLanguage() js.Ref

//go:wasmimport plat/js/webext/i18n func_DetectLanguage
//go:noescape
func FuncDetectLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n call_DetectLanguage
//go:noescape
func CallDetectLanguage(
	retPtr unsafe.Pointer,
	text js.Ref)

//go:wasmimport plat/js/webext/i18n try_DetectLanguage
//go:noescape
func TryDetectLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/i18n has_GetAcceptLanguages
//go:noescape
func HasFuncGetAcceptLanguages() js.Ref

//go:wasmimport plat/js/webext/i18n func_GetAcceptLanguages
//go:noescape
func FuncGetAcceptLanguages(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n call_GetAcceptLanguages
//go:noescape
func CallGetAcceptLanguages(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n try_GetAcceptLanguages
//go:noescape
func TryGetAcceptLanguages(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/i18n has_GetMessage
//go:noescape
func HasFuncGetMessage() js.Ref

//go:wasmimport plat/js/webext/i18n func_GetMessage
//go:noescape
func FuncGetMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n call_GetMessage
//go:noescape
func CallGetMessage(
	retPtr unsafe.Pointer,
	messageName js.Ref,
	substitutions js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n try_GetMessage
//go:noescape
func TryGetMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	messageName js.Ref,
	substitutions js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/i18n has_GetUILanguage
//go:noescape
func HasFuncGetUILanguage() js.Ref

//go:wasmimport plat/js/webext/i18n func_GetUILanguage
//go:noescape
func FuncGetUILanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n call_GetUILanguage
//go:noescape
func CallGetUILanguage(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/i18n try_GetUILanguage
//go:noescape
func TryGetUILanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
