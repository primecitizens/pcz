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

//go:wasmimport plat/js/webext/offscreen constof_Reason
//go:noescape
func ConstOfReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/offscreen store_CreateParameters
//go:noescape
func CreateParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/offscreen load_CreateParameters
//go:noescape
func CreateParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/offscreen has_CloseDocument
//go:noescape
func HasFuncCloseDocument() js.Ref

//go:wasmimport plat/js/webext/offscreen func_CloseDocument
//go:noescape
func FuncCloseDocument(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen call_CloseDocument
//go:noescape
func CallCloseDocument(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen try_CloseDocument
//go:noescape
func TryCloseDocument(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/offscreen has_CreateDocument
//go:noescape
func HasFuncCreateDocument() js.Ref

//go:wasmimport plat/js/webext/offscreen func_CreateDocument
//go:noescape
func FuncCreateDocument(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen call_CreateDocument
//go:noescape
func CallCreateDocument(
	retPtr unsafe.Pointer,
	parameters unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen try_CreateDocument
//go:noescape
func TryCreateDocument(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/offscreen has_HasDocument
//go:noescape
func HasFuncHasDocument() js.Ref

//go:wasmimport plat/js/webext/offscreen func_HasDocument
//go:noescape
func FuncHasDocument(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen call_HasDocument
//go:noescape
func CallHasDocument(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/offscreen try_HasDocument
//go:noescape
func TryHasDocument(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
