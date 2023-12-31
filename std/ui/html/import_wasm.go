// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package html

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

//go:wasmimport ui/html/builder reset
//go:noescape
func reset(elem js.Ref)

//go:wasmimport ui/html/builder free
//go:noescape
func free(elem js.Ref)

//go:wasmimport ui/html/builder integer
//go:noescape
func integer(elem js.Ref, signed js.Ref, p64 unsafe.Pointer, radix uint32)

//go:wasmimport ui/html/builder float
//go:noescape
func float(elem js.Ref, val float64)

//go:wasmimport ui/html/builder buf_str
//go:noescape
func buf_str(elem js.Ref, str unsafe.Pointer, len uint32)

//go:wasmimport ui/html/builder buf_jsstr
//go:noescape
func buf_jsstr(elem js.Ref, escape js.Ref, ptr unsafe.Pointer, len uint32)

//go:wasmimport ui/html/builder flush
//go:noescape
func flush(elem js.Ref, append js.Ref)
