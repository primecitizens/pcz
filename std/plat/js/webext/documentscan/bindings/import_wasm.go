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

//go:wasmimport plat/js/webext/documentscan store_ScanResults
//go:noescape
func ScanResultsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/documentscan load_ScanResults
//go:noescape
func ScanResultsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/documentscan store_ScanOptions
//go:noescape
func ScanOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/documentscan load_ScanOptions
//go:noescape
func ScanOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/documentscan has_Scan
//go:noescape
func HasFuncScan() js.Ref

//go:wasmimport plat/js/webext/documentscan func_Scan
//go:noescape
func FuncScan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/documentscan call_Scan
//go:noescape
func CallScan(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/documentscan try_Scan
//go:noescape
func TryScan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)
