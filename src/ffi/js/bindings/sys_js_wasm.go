// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js && wasm

package bindings

import (
	"unsafe"
	_ "unsafe" // for go:linkname
)

func Print(s string) {
	print(
		Uintptr32(uintptr(
			unsafe.Pointer(
				unsafe.StringData(s),
			),
		)),
		uint32(len(s)),
	)
}

//go:wasmimport ffi/js print
func print(ptr Uintptr32, sz uint32)

//go:wasmimport ffi/js exit
func Exit(code int32)

//go:wasmimport ffi/js resetMemoryDataView
func ResetMemoryDataView()

//go:wasmimport ffi/js argsEnvsSize
func ArgsEnvsSize(outCount Uintptr32) uint32

//go:wasmimport ffi/js nthsysInto
func NthSysInto(i uint32, buf Uintptr32, sz uint32) int32
