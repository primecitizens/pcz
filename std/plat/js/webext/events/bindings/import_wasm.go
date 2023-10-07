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

//go:wasmimport plat/js/webext/events store_Rule
//go:noescape
func RuleJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/events load_Rule
//go:noescape
func RuleJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/events has_Event_AddListener
//go:noescape
func HasFuncEventAddListener(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_AddListener
//go:noescape
func FuncEventAddListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_AddListener
//go:noescape
func CallEventAddListener(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_AddListener
//go:noescape
func TryEventAddListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_RemoveListener
//go:noescape
func HasFuncEventRemoveListener(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_RemoveListener
//go:noescape
func FuncEventRemoveListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_RemoveListener
//go:noescape
func CallEventRemoveListener(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_RemoveListener
//go:noescape
func TryEventRemoveListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_HasListener
//go:noescape
func HasFuncEventHasListener(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_HasListener
//go:noescape
func FuncEventHasListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_HasListener
//go:noescape
func CallEventHasListener(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_HasListener
//go:noescape
func TryEventHasListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_HasListeners
//go:noescape
func HasFuncEventHasListeners(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_HasListeners
//go:noescape
func FuncEventHasListeners(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_HasListeners
//go:noescape
func CallEventHasListeners(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/events try_Event_HasListeners
//go:noescape
func TryEventHasListeners(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_AddRules
//go:noescape
func HasFuncEventAddRules(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_AddRules
//go:noescape
func FuncEventAddRules(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_AddRules
//go:noescape
func CallEventAddRules(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	rules js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_AddRules
//go:noescape
func TryEventAddRules(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	rules js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_AddRules1
//go:noescape
func HasFuncEventAddRules1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_AddRules1
//go:noescape
func FuncEventAddRules1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_AddRules1
//go:noescape
func CallEventAddRules1(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	rules js.Ref)

//go:wasmimport plat/js/webext/events try_Event_AddRules1
//go:noescape
func TryEventAddRules1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	rules js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_GetRules
//go:noescape
func HasFuncEventGetRules(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_GetRules
//go:noescape
func FuncEventGetRules(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_GetRules
//go:noescape
func CallEventGetRules(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_GetRules
//go:noescape
func TryEventGetRules(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_GetRules1
//go:noescape
func HasFuncEventGetRules1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_GetRules1
//go:noescape
func FuncEventGetRules1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_GetRules1
//go:noescape
func CallEventGetRules1(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64)

//go:wasmimport plat/js/webext/events try_Event_GetRules1
//go:noescape
func TryEventGetRules1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_RemoveRules
//go:noescape
func HasFuncEventRemoveRules(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_RemoveRules
//go:noescape
func FuncEventRemoveRules(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_RemoveRules
//go:noescape
func CallEventRemoveRules(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/events try_Event_RemoveRules
//go:noescape
func TryEventRemoveRules(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_RemoveRules1
//go:noescape
func HasFuncEventRemoveRules1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_RemoveRules1
//go:noescape
func FuncEventRemoveRules1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_RemoveRules1
//go:noescape
func CallEventRemoveRules1(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref)

//go:wasmimport plat/js/webext/events try_Event_RemoveRules1
//go:noescape
func TryEventRemoveRules1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64,
	ruleIdentifiers js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/events has_Event_RemoveRules2
//go:noescape
func HasFuncEventRemoveRules2(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/events func_Event_RemoveRules2
//go:noescape
func FuncEventRemoveRules2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/events call_Event_RemoveRules2
//go:noescape
func CallEventRemoveRules2(
	this js.Ref, retPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64)

//go:wasmimport plat/js/webext/events try_Event_RemoveRules2
//go:noescape
func TryEventRemoveRules2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventName js.Ref,
	webViewInstanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/events store_UrlFilter
//go:noescape
func UrlFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/events load_UrlFilter
//go:noescape
func UrlFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
