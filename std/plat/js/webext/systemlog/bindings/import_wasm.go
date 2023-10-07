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

//go:wasmimport plat/js/webext/systemlog store_MessageOptions
//go:noescape
func MessageOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/systemlog load_MessageOptions
//go:noescape
func MessageOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/systemlog has_Add
//go:noescape
func HasFuncAdd() js.Ref

//go:wasmimport plat/js/webext/systemlog func_Add
//go:noescape
func FuncAdd(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/systemlog call_Add
//go:noescape
func CallAdd(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/systemlog try_Add
//go:noescape
func TryAdd(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
