// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"
)

//go:wasmimport core/sys sizes
//go:noescape
func Sizes(outCount unsafe.Pointer) uint32

//go:wasmimport core/sys nth
//go:noescape
func Nth(i uint32, buf unsafe.Pointer, sz uint32) int32
