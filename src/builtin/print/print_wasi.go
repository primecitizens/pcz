// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasip1

package stdprint

import (
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/ffi/wasm/wasi"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	var (
		iovs = wasi.IOBuffer{
			Ptr: wasi.Uintptr(mark.NoEscapeSliceDataPointer(b)),
			Len: wasi.Size(len(b)),
		}
		n wasi.Size
	)

	wasi.Write(
		wasi.Stderr,
		mark.NoEscapePointer(&iovs),
		1,
		mark.NoEscapePointer(&n),
	)
}
