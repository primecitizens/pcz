// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm && js

package sbrkalloc

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/alloc/sbrkalloc/bindings"
)

func init() {
	// web browsers provides a default 16MB memry, not matching
	// the actual size of the wasm blob.
	for i, iter := 0, abi.ModuleIter(0); ; i++ {
		md, ok := iter.Nth(i)
		if !ok {
			break
		}

		if md.HasMain == 1 {
			_start = md.End
			break
		}
	}

	if _start == 0 {
		_start = uintptr(Sbrk(0))
		base = _start
		end = _start
	} else {
		base = _start
		end = uintptr(Sbrk(0))
	}

	lastFree = nil
}

func resetMemoryDataView() {
	bindings.ResetMemoryDataView()
}
