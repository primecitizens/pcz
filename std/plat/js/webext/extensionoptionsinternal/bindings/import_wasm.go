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

//go:wasmimport plat/js/webext/extensionoptionsinternal store_PreferredSizeChangedOptions
//go:noescape
func PreferredSizeChangedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal load_PreferredSizeChangedOptions
//go:noescape
func PreferredSizeChangedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal store_SizeChangedOptions
//go:noescape
func SizeChangedOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal load_SizeChangedOptions
//go:noescape
func SizeChangedOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OnClose
//go:noescape
func HasFuncOnClose() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OnClose
//go:noescape
func FuncOnClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OnClose
//go:noescape
func CallOnClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OnClose
//go:noescape
func TryOnClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OffClose
//go:noescape
func HasFuncOffClose() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OffClose
//go:noescape
func FuncOffClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OffClose
//go:noescape
func CallOffClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OffClose
//go:noescape
func TryOffClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_HasOnClose
//go:noescape
func HasFuncHasOnClose() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_HasOnClose
//go:noescape
func FuncHasOnClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_HasOnClose
//go:noescape
func CallHasOnClose(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_HasOnClose
//go:noescape
func TryHasOnClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OnLoad
//go:noescape
func HasFuncOnLoad() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OnLoad
//go:noescape
func FuncOnLoad(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OnLoad
//go:noescape
func CallOnLoad(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OnLoad
//go:noescape
func TryOnLoad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OffLoad
//go:noescape
func HasFuncOffLoad() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OffLoad
//go:noescape
func FuncOffLoad(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OffLoad
//go:noescape
func CallOffLoad(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OffLoad
//go:noescape
func TryOffLoad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_HasOnLoad
//go:noescape
func HasFuncHasOnLoad() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_HasOnLoad
//go:noescape
func FuncHasOnLoad(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_HasOnLoad
//go:noescape
func CallHasOnLoad(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_HasOnLoad
//go:noescape
func TryHasOnLoad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OnPreferredSizeChanged
//go:noescape
func HasFuncOnPreferredSizeChanged() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OnPreferredSizeChanged
//go:noescape
func FuncOnPreferredSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OnPreferredSizeChanged
//go:noescape
func CallOnPreferredSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OnPreferredSizeChanged
//go:noescape
func TryOnPreferredSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_OffPreferredSizeChanged
//go:noescape
func HasFuncOffPreferredSizeChanged() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_OffPreferredSizeChanged
//go:noescape
func FuncOffPreferredSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_OffPreferredSizeChanged
//go:noescape
func CallOffPreferredSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_OffPreferredSizeChanged
//go:noescape
func TryOffPreferredSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal has_HasOnPreferredSizeChanged
//go:noescape
func HasFuncHasOnPreferredSizeChanged() js.Ref

//go:wasmimport plat/js/webext/extensionoptionsinternal func_HasOnPreferredSizeChanged
//go:noescape
func FuncHasOnPreferredSizeChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/extensionoptionsinternal call_HasOnPreferredSizeChanged
//go:noescape
func CallHasOnPreferredSizeChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/extensionoptionsinternal try_HasOnPreferredSizeChanged
//go:noescape
func TryHasOnPreferredSizeChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
