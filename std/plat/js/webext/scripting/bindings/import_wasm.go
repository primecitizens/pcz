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

//go:wasmimport plat/js/webext/scripting store_InjectionTarget
//go:noescape
func InjectionTargetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_InjectionTarget
//go:noescape
func InjectionTargetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting constof_StyleOrigin
//go:noescape
func ConstOfStyleOrigin(str js.Ref) uint32

//go:wasmimport plat/js/webext/scripting store_CSSInjection
//go:noescape
func CSSInjectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_CSSInjection
//go:noescape
func CSSInjectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting store_ContentScriptFilter
//go:noescape
func ContentScriptFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_ContentScriptFilter
//go:noescape
func ContentScriptFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting constof_ExecutionWorld
//go:noescape
func ConstOfExecutionWorld(str js.Ref) uint32

//go:wasmimport plat/js/webext/scripting store_RegisteredContentScript
//go:noescape
func RegisteredContentScriptJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_RegisteredContentScript
//go:noescape
func RegisteredContentScriptJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting store_InjectionResult
//go:noescape
func InjectionResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_InjectionResult
//go:noescape
func InjectionResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting has_Properties_GlobalParams
//go:noescape
func HasFuncPropertiesGlobalParams(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting func_Properties_GlobalParams
//go:noescape
func FuncPropertiesGlobalParams(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_Properties_GlobalParams
//go:noescape
func CallPropertiesGlobalParams(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_Properties_GlobalParams
//go:noescape
func TryPropertiesGlobalParams(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting store_ScriptInjection
//go:noescape
func ScriptInjectionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/scripting load_ScriptInjection
//go:noescape
func ScriptInjectionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/scripting has_ExecuteScript
//go:noescape
func HasFuncExecuteScript() js.Ref

//go:wasmimport plat/js/webext/scripting func_ExecuteScript
//go:noescape
func FuncExecuteScript(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_ExecuteScript
//go:noescape
func CallExecuteScript(
	retPtr unsafe.Pointer,
	injection unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_ExecuteScript
//go:noescape
func TryExecuteScript(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	injection unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_GetRegisteredContentScripts
//go:noescape
func HasFuncGetRegisteredContentScripts() js.Ref

//go:wasmimport plat/js/webext/scripting func_GetRegisteredContentScripts
//go:noescape
func FuncGetRegisteredContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_GetRegisteredContentScripts
//go:noescape
func CallGetRegisteredContentScripts(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_GetRegisteredContentScripts
//go:noescape
func TryGetRegisteredContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_InsertCSS
//go:noescape
func HasFuncInsertCSS() js.Ref

//go:wasmimport plat/js/webext/scripting func_InsertCSS
//go:noescape
func FuncInsertCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_InsertCSS
//go:noescape
func CallInsertCSS(
	retPtr unsafe.Pointer,
	injection unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_InsertCSS
//go:noescape
func TryInsertCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	injection unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_RegisterContentScripts
//go:noescape
func HasFuncRegisterContentScripts() js.Ref

//go:wasmimport plat/js/webext/scripting func_RegisterContentScripts
//go:noescape
func FuncRegisterContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_RegisterContentScripts
//go:noescape
func CallRegisterContentScripts(
	retPtr unsafe.Pointer,
	scripts js.Ref)

//go:wasmimport plat/js/webext/scripting try_RegisterContentScripts
//go:noescape
func TryRegisterContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scripts js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_RemoveCSS
//go:noescape
func HasFuncRemoveCSS() js.Ref

//go:wasmimport plat/js/webext/scripting func_RemoveCSS
//go:noescape
func FuncRemoveCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_RemoveCSS
//go:noescape
func CallRemoveCSS(
	retPtr unsafe.Pointer,
	injection unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_RemoveCSS
//go:noescape
func TryRemoveCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	injection unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_UnregisterContentScripts
//go:noescape
func HasFuncUnregisterContentScripts() js.Ref

//go:wasmimport plat/js/webext/scripting func_UnregisterContentScripts
//go:noescape
func FuncUnregisterContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_UnregisterContentScripts
//go:noescape
func CallUnregisterContentScripts(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting try_UnregisterContentScripts
//go:noescape
func TryUnregisterContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/scripting has_UpdateContentScripts
//go:noescape
func HasFuncUpdateContentScripts() js.Ref

//go:wasmimport plat/js/webext/scripting func_UpdateContentScripts
//go:noescape
func FuncUpdateContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/scripting call_UpdateContentScripts
//go:noescape
func CallUpdateContentScripts(
	retPtr unsafe.Pointer,
	scripts js.Ref)

//go:wasmimport plat/js/webext/scripting try_UpdateContentScripts
//go:noescape
func TryUpdateContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scripts js.Ref) (ok js.Ref)
