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

//go:wasmimport plat/js/webext/commandlineprivate has_HasSwitch
//go:noescape
func HasFuncHasSwitch() js.Ref

//go:wasmimport plat/js/webext/commandlineprivate func_HasSwitch
//go:noescape
func FuncHasSwitch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/commandlineprivate call_HasSwitch
//go:noescape
func CallHasSwitch(
	retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/webext/commandlineprivate try_HasSwitch
//go:noescape
func TryHasSwitch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)
