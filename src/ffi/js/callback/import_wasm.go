// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package callback

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

//go:wasmimport ffi/js/callback func
//go:noescape
func Func(
	dispFnPC,
	handler unsafe.Pointer,
	targetPC uint32,
) bindings.Ref

//go:wasmimport ffi/js/callback ctx
//go:noescape
func Context(
	ctx bindings.Ref,
	ptr unsafe.Pointer,
)
