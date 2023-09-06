// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

package sysalloc

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/arch"
)

// Sbrk
//
//go:nosplit
func Sbrk(size uintptr) unsafe.Pointer {
	i := int32(divRoundUp(size, arch.DefaultPhysPageSize))
	i = growMemory(i)
	if i < 0 {
		return nil
	}
	resetMemoryDataView()
	return unsafe.Pointer(uintptr(i) * arch.DefaultPhysPageSize)
}

// growMemory is a simple wrapper for wasm instruction memory.grow.
//
// When return value >= 0, it is the previous size (in pages).
func growMemory(pages int32) int32
