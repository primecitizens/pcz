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

//go:wasmimport plat/js/webext/topsites store_MostVisitedURL
//go:noescape
func MostVisitedURLJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/topsites load_MostVisitedURL
//go:noescape
func MostVisitedURLJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/topsites has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/topsites func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/topsites call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/topsites try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
