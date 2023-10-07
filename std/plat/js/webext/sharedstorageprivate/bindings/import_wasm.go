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

//go:wasmimport plat/js/webext/sharedstorageprivate has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/sharedstorageprivate func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sharedstorageprivate call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/sharedstorageprivate try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sharedstorageprivate has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/sharedstorageprivate func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sharedstorageprivate call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	keys js.Ref)

//go:wasmimport plat/js/webext/sharedstorageprivate try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keys js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/sharedstorageprivate has_Set
//go:noescape
func HasFuncSet() js.Ref

//go:wasmimport plat/js/webext/sharedstorageprivate func_Set
//go:noescape
func FuncSet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sharedstorageprivate call_Set
//go:noescape
func CallSet(
	retPtr unsafe.Pointer,
	items js.Ref)

//go:wasmimport plat/js/webext/sharedstorageprivate try_Set
//go:noescape
func TrySet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	items js.Ref) (ok js.Ref)
