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

//go:wasmimport plat/js/webext/mojoprivate has_RequireAsync
//go:noescape
func HasFuncRequireAsync() js.Ref

//go:wasmimport plat/js/webext/mojoprivate func_RequireAsync
//go:noescape
func FuncRequireAsync(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mojoprivate call_RequireAsync
//go:noescape
func CallRequireAsync(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/mojoprivate try_RequireAsync
//go:noescape
func TryRequireAsync(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)
