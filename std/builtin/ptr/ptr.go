// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdptr

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/num"
	"github.com/primecitizens/pcz/std/core/thread"
)

// OnStack reports whether the pointer points to somewhere of
// current goroutine's stack.
func OnStack(ptr uintptr) bool {
	return thread.G().Stack.PointerOnStack(ptr)
}

// Cast casts numeric pointer value to *R
//
//go:nosplit
func Cast[T any, Ptr num.Pointer](p Ptr) *T {
	return (*T)(unsafe.Pointer(p))
}

// Assign value of type T to the address: *(base+offset) = v
//
//go:nosplit
func Assign[T any, Ptr num.Pointer](base Ptr, offset uintptr, v T) {
	*(*T)(unsafe.Pointer(uintptr(base) + offset)) = v
}

// Add adds offset off to p.
//
//go:nosplit
func Add[T any, O num.Integer](p *T, off O) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(p), off))
}

func CheckAlignment(p unsafe.Pointer, elem *abi.Type, n uintptr) {
	// nil pointer is always suitably aligned (#47430).
	if p == nil {
		return
	}

	// Check that (*[n]elem)(p) is appropriately aligned.
	// Note that we allow unaligned pointers if the types they point to contain
	// no pointers themselves. See issue 37298.
	// TODO(mdempsky): What about fieldAlign?
	if elem.PtrBytes != 0 && uintptr(p)&(uintptr(elem.Align_)-1) != 0 {
		assert.Throw("checkptr: misaligned pointer conversion")
	}

	// Check that (*[n]elem)(p) doesn't straddle multiple heap objects.
	// TODO(mdempsky): Fix #46938 so we don't need to worry about overflow here.
	if CheckptrStraddles(p, n*elem.Size_) {
		assert.Throw("checkptr: converted pointer straddles multiple allocations")
	}
}

// CheckptrStraddles reports whether the first size-bytes of memory
// addressed by ptr is known to straddle more than one Go allocation.
func CheckptrStraddles(ptr unsafe.Pointer, size uintptr) bool {
	if size <= 1 {
		return false
	}

	// Check that add(ptr, size-1) won't overflow. This avoids the risk
	// of producing an illegal pointer value (assuming ptr is legal).
	if uintptr(ptr) >= -(size - 1) {
		return true
	}
	end := unsafe.Add(ptr, size-1)

	// TODO(mdempsky): Detect when [ptr, end] contains Go allocations,
	// but neither ptr nor end point into one themselves.

	return CheckptrBase(ptr) != CheckptrBase(end)
}

func CheckArithmetic(p unsafe.Pointer, originals []unsafe.Pointer) {
	// MinLegalPointer is the smallest possible legal pointer.
	// This is the smallest possible architectural page size,
	// since we assume that the first page is never mapped.
	//
	// This should agree with minZeroPage in the compiler.
	const MinLegalPointer uintptr = 4096

	if 0 < uintptr(p) && uintptr(p) < MinLegalPointer {
		assert.Throw("checkptr: pointer arithmetic computed bad pointer value")
	}

	// Check that if the computed pointer p points into a heap
	// object, then one of the original pointers must have pointed
	// into the same object.
	base := CheckptrBase(p)
	if base == 0 {
		return
	}

	for _, original := range originals {
		if base == CheckptrBase(original) {
			return
		}
	}

	assert.Throw("checkptr: pointer arithmetic result points to invalid allocation")
}

// CheckptrBase returns the base address for the allocation containing
// the address p.
//
// Importantly, if p1 and p2 point into the same variable, then
// CheckptrBase(p1) == CheckptrBase(p2). However, the converse/inverse
// is not necessarily true as allocations can have trailing padding,
// and multiple variables may be packed into a single allocation.
func CheckptrBase(p unsafe.Pointer) uintptr {
	// 	// stack
	// 	if gp := getg(); gp.stack.lo <= uintptr(p) && uintptr(p) < gp.stack.hi {
	// 		// TODO(mdempsky): Walk the stack to identify the
	// 		// specific stack frame or even stack object that p
	// 		// points into.
	// 		//
	// 		// In the mean time, use "1" as a pseudo-address to
	// 		// represent the stack. This is an invalid address on
	// 		// all platforms, so it's guaranteed to be distinct
	// 		// from any of the addresses we might return below.
	// 		return 1
	// 	}
	//
	// 	// heap (must check after stack because of #35068)
	// 	if base, _, _ := findObject(uintptr(p), 0, 0); base != 0 {
	// 		return base
	// 	}
	//
	// 	// data or bss
	// 	for _, datap := range activeModules() {
	// 		if datap.data <= uintptr(p) && uintptr(p) < datap.edata {
	// 			return datap.data
	// 		}
	// 		if datap.bss <= uintptr(p) && uintptr(p) < datap.ebss {
	// 			return datap.bss
	// 		}
	// 	}

	return 0
}
