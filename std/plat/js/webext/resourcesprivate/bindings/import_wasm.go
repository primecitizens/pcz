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

//go:wasmimport plat/js/webext/resourcesprivate constof_Component
//go:noescape
func ConstOfComponent(str js.Ref) uint32

//go:wasmimport plat/js/webext/resourcesprivate has_GetStrings
//go:noescape
func HasFuncGetStrings() js.Ref

//go:wasmimport plat/js/webext/resourcesprivate func_GetStrings
//go:noescape
func FuncGetStrings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/resourcesprivate call_GetStrings
//go:noescape
func CallGetStrings(
	retPtr unsafe.Pointer,
	component uint32)

//go:wasmimport plat/js/webext/resourcesprivate try_GetStrings
//go:noescape
func TryGetStrings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	component uint32) (ok js.Ref)
