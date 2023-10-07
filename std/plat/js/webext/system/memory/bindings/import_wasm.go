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

//go:wasmimport plat/js/webext/system/memory store_MemoryInfo
//go:noescape
func MemoryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/memory load_MemoryInfo
//go:noescape
func MemoryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/memory has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/system/memory func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/memory call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/system/memory try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
