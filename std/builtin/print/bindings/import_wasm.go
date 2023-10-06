// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"
)

//go:wasmimport stdprint print
//go:noescape
func Print(str unsafe.Pointer, sz uint32)
