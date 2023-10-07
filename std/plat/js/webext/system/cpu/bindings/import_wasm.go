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

//go:wasmimport plat/js/webext/system/cpu store_CpuTime
//go:noescape
func CpuTimeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/cpu load_CpuTime
//go:noescape
func CpuTimeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/cpu store_ProcessorInfo
//go:noescape
func ProcessorInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/cpu load_ProcessorInfo
//go:noescape
func ProcessorInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/cpu store_CpuInfo
//go:noescape
func CpuInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/cpu load_CpuInfo
//go:noescape
func CpuInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/cpu has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/system/cpu func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/cpu call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/system/cpu try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)
