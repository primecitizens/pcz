// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/core/mark"
)

//go:wasmimport core/assert append
//go:noescape
func append(str unsafe.Pointer, len uint32)

//go:wasmimport core/assert throw
//go:noescape
func Throw()

func Append(s ...string) {
	for _, a := range s {
		append(mark.NoEscapeStringDataPointer(a), uint32(len(a)))
	}
}
