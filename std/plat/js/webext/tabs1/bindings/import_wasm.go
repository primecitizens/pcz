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

//go:wasmimport plat/js/webext/tabs1 has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/tabs1 func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs1 call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	tabId float64,
	connectInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs1 try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	connectInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs1 has_Highlight
//go:noescape
func HasFuncHighlight() js.Ref

//go:wasmimport plat/js/webext/tabs1 func_Highlight
//go:noescape
func FuncHighlight(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs1 call_Highlight
//go:noescape
func CallHighlight(
	retPtr unsafe.Pointer,
	highlightInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs1 try_Highlight
//go:noescape
func TryHighlight(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	highlightInfo unsafe.Pointer) (ok js.Ref)
