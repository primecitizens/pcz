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

//go:wasmimport plat/js/webext/loginscreenui store_ShowOptions
//go:noescape
func ShowOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/loginscreenui load_ShowOptions
//go:noescape
func ShowOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/loginscreenui has_Close
//go:noescape
func HasFuncClose() js.Ref

//go:wasmimport plat/js/webext/loginscreenui func_Close
//go:noescape
func FuncClose(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenui call_Close
//go:noescape
func CallClose(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenui try_Close
//go:noescape
func TryClose(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/loginscreenui has_Show
//go:noescape
func HasFuncShow() js.Ref

//go:wasmimport plat/js/webext/loginscreenui func_Show
//go:noescape
func FuncShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenui call_Show
//go:noescape
func CallShow(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/loginscreenui try_Show
//go:noescape
func TryShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
