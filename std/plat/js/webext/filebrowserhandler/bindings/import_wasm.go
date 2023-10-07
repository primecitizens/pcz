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

//go:wasmimport plat/js/webext/filebrowserhandler store_FileHandlerExecuteEventDetails
//go:noescape
func FileHandlerExecuteEventDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler load_FileHandlerExecuteEventDetails
//go:noescape
func FileHandlerExecuteEventDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/filebrowserhandler has_OnExecute
//go:noescape
func HasFuncOnExecute() js.Ref

//go:wasmimport plat/js/webext/filebrowserhandler func_OnExecute
//go:noescape
func FuncOnExecute(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filebrowserhandler call_OnExecute
//go:noescape
func CallOnExecute(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler try_OnExecute
//go:noescape
func TryOnExecute(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler has_OffExecute
//go:noescape
func HasFuncOffExecute() js.Ref

//go:wasmimport plat/js/webext/filebrowserhandler func_OffExecute
//go:noescape
func FuncOffExecute(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filebrowserhandler call_OffExecute
//go:noescape
func CallOffExecute(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler try_OffExecute
//go:noescape
func TryOffExecute(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler has_HasOnExecute
//go:noescape
func HasFuncHasOnExecute() js.Ref

//go:wasmimport plat/js/webext/filebrowserhandler func_HasOnExecute
//go:noescape
func FuncHasOnExecute(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filebrowserhandler call_HasOnExecute
//go:noescape
func CallHasOnExecute(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/filebrowserhandler try_HasOnExecute
//go:noescape
func TryHasOnExecute(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
