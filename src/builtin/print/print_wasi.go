// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasip1

package stdprint

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/wasm/wasi"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	wasi.Write(wasi.Stderr, wasi.IOBuffer{
		Ptr: uint32(uintptr(unsafe.Pointer(&b[0]))),
		Len: wasi.Size(len(b)),
	})
}
