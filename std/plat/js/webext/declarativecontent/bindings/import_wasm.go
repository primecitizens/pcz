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

//go:wasmimport plat/js/webext/declarativecontent constof_PageStateMatcherInstanceType
//go:noescape
func ConstOfPageStateMatcherInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativecontent store_PageStateMatcher
//go:noescape
func PageStateMatcherJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativecontent load_PageStateMatcher
//go:noescape
func PageStateMatcherJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativecontent constof_RequestContentScriptInstanceType
//go:noescape
func ConstOfRequestContentScriptInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativecontent store_RequestContentScript
//go:noescape
func RequestContentScriptJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativecontent load_RequestContentScript
//go:noescape
func RequestContentScriptJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativecontent constof_SetIconInstanceType
//go:noescape
func ConstOfSetIconInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativecontent store_SetIcon
//go:noescape
func SetIconJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativecontent load_SetIcon
//go:noescape
func SetIconJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativecontent constof_ShowActionInstanceType
//go:noescape
func ConstOfShowActionInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativecontent store_ShowAction
//go:noescape
func ShowActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativecontent load_ShowAction
//go:noescape
func ShowActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativecontent constof_ShowPageActionInstanceType
//go:noescape
func ConstOfShowPageActionInstanceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/declarativecontent store_ShowPageAction
//go:noescape
func ShowPageActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/declarativecontent load_ShowPageAction
//go:noescape
func ShowPageActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/declarativecontent has_OnPageChanged
//go:noescape
func HasFuncOnPageChanged() js.Ref

//go:wasmimport plat/js/webext/declarativecontent func_OnPageChanged
//go:noescape
func FuncOnPageChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativecontent call_OnPageChanged
//go:noescape
func CallOnPageChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativecontent try_OnPageChanged
//go:noescape
func TryOnPageChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativecontent has_OffPageChanged
//go:noescape
func HasFuncOffPageChanged() js.Ref

//go:wasmimport plat/js/webext/declarativecontent func_OffPageChanged
//go:noescape
func FuncOffPageChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativecontent call_OffPageChanged
//go:noescape
func CallOffPageChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativecontent try_OffPageChanged
//go:noescape
func TryOffPageChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/declarativecontent has_HasOnPageChanged
//go:noescape
func HasFuncHasOnPageChanged() js.Ref

//go:wasmimport plat/js/webext/declarativecontent func_HasOnPageChanged
//go:noescape
func FuncHasOnPageChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/declarativecontent call_HasOnPageChanged
//go:noescape
func CallHasOnPageChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/declarativecontent try_HasOnPageChanged
//go:noescape
func TryHasOnPageChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
