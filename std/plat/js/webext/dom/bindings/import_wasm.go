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

//go:wasmimport plat/js/webext/dom has_OpenOrClosedShadowRoot
//go:noescape
func HasFuncOpenOrClosedShadowRoot() js.Ref

//go:wasmimport plat/js/webext/dom func_OpenOrClosedShadowRoot
//go:noescape
func FuncOpenOrClosedShadowRoot(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/dom call_OpenOrClosedShadowRoot
//go:noescape
func CallOpenOrClosedShadowRoot(
	retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/webext/dom try_OpenOrClosedShadowRoot
//go:noescape
func TryOpenOrClosedShadowRoot(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)
