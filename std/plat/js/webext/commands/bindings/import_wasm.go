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

//go:wasmimport plat/js/webext/commands store_Command
//go:noescape
func CommandJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/commands load_Command
//go:noescape
func CommandJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/commands has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/commands func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/commands call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/commands try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/commands has_OnCommand
//go:noescape
func HasFuncOnCommand() js.Ref

//go:wasmimport plat/js/webext/commands func_OnCommand
//go:noescape
func FuncOnCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/commands call_OnCommand
//go:noescape
func CallOnCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/commands try_OnCommand
//go:noescape
func TryOnCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/commands has_OffCommand
//go:noescape
func HasFuncOffCommand() js.Ref

//go:wasmimport plat/js/webext/commands func_OffCommand
//go:noescape
func FuncOffCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/commands call_OffCommand
//go:noescape
func CallOffCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/commands try_OffCommand
//go:noescape
func TryOffCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/commands has_HasOnCommand
//go:noescape
func HasFuncHasOnCommand() js.Ref

//go:wasmimport plat/js/webext/commands func_HasOnCommand
//go:noescape
func FuncHasOnCommand(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/commands call_HasOnCommand
//go:noescape
func CallHasOnCommand(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/commands try_HasOnCommand
//go:noescape
func TryHasOnCommand(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
