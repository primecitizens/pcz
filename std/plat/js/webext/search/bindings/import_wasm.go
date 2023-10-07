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

//go:wasmimport plat/js/webext/search constof_Disposition
//go:noescape
func ConstOfDisposition(str js.Ref) uint32

//go:wasmimport plat/js/webext/search store_QueryInfo
//go:noescape
func QueryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/search load_QueryInfo
//go:noescape
func QueryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/search has_Query
//go:noescape
func HasFuncQuery() js.Ref

//go:wasmimport plat/js/webext/search func_Query
//go:noescape
func FuncQuery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/search call_Query
//go:noescape
func CallQuery(
	retPtr unsafe.Pointer,
	queryInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/search try_Query
//go:noescape
func TryQuery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryInfo unsafe.Pointer) (ok js.Ref)
