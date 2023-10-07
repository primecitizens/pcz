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

//go:wasmimport plat/js/webext/userscripts store_ScriptSource
//go:noescape
func ScriptSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/userscripts load_ScriptSource
//go:noescape
func ScriptSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/userscripts store_RegisteredUserScript
//go:noescape
func RegisteredUserScriptJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/userscripts load_RegisteredUserScript
//go:noescape
func RegisteredUserScriptJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/userscripts store_UserScriptFilter
//go:noescape
func UserScriptFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/userscripts load_UserScriptFilter
//go:noescape
func UserScriptFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/userscripts has_GetScripts
//go:noescape
func HasFuncGetScripts() js.Ref

//go:wasmimport plat/js/webext/userscripts func_GetScripts
//go:noescape
func FuncGetScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/userscripts call_GetScripts
//go:noescape
func CallGetScripts(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/userscripts try_GetScripts
//go:noescape
func TryGetScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/userscripts has_Register
//go:noescape
func HasFuncRegister() js.Ref

//go:wasmimport plat/js/webext/userscripts func_Register
//go:noescape
func FuncRegister(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/userscripts call_Register
//go:noescape
func CallRegister(
	retPtr unsafe.Pointer,
	scripts js.Ref)

//go:wasmimport plat/js/webext/userscripts try_Register
//go:noescape
func TryRegister(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scripts js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/userscripts has_Unregister
//go:noescape
func HasFuncUnregister() js.Ref

//go:wasmimport plat/js/webext/userscripts func_Unregister
//go:noescape
func FuncUnregister(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/userscripts call_Unregister
//go:noescape
func CallUnregister(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/userscripts try_Unregister
//go:noescape
func TryUnregister(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)
