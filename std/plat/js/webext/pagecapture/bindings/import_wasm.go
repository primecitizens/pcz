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

//go:wasmimport plat/js/webext/pagecapture store_SaveAsMHTMLArgDetails
//go:noescape
func SaveAsMHTMLArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/pagecapture load_SaveAsMHTMLArgDetails
//go:noescape
func SaveAsMHTMLArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/pagecapture has_SaveAsMHTML
//go:noescape
func HasFuncSaveAsMHTML() js.Ref

//go:wasmimport plat/js/webext/pagecapture func_SaveAsMHTML
//go:noescape
func FuncSaveAsMHTML(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pagecapture call_SaveAsMHTML
//go:noescape
func CallSaveAsMHTML(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/pagecapture try_SaveAsMHTML
//go:noescape
func TrySaveAsMHTML(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)
