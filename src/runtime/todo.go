// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/assert"
)

// TODO

// TODO(alloc): standardize allocation size like go sizeclass?
func roundupsize(sz uintptr) uintptr {
	return sz
}

// bulkBarrierPreWriteSrcOnly is like bulkBarrierPreWrite but
// does not execute write barriers for [dst, dst+size).
//
// In addition to the requirements of bulkBarrierPreWrite
// callers need to ensure [dst, dst+size) is zeroed.
//
// This is used for special cases where e.g. dst was just
// created and zeroed with malloc.
//
// See ${GOROOT}/src/runtime/mbitmap.go#func:bulkBarrierPreWriteSrcOnly
//
//go:nosplit
func bulkBarrierPreWriteSrcOnly(dst, src, size uintptr) {
	if (dst|src|size)&(arch.PtrSize-1) != 0 {
		assert.Throw("bulkBarrierPreWrite:", "unaligned", "arguments")
	}
	if !writeBarrier.needed {
		return
	}

	// TODO
}

// bulkBarrierPreWrite executes a write barrier
// for every pointer slot in the memory range [src, src+size),
// using pointer/scalar information from [dst, dst+size).
// This executes the write barriers necessary before a memmove.
// src, dst, and size must be pointer-aligned.
// The range [dst, dst+size) must lie within a single object.
// It does not perform the actual writes.
//
// As a special case, src == 0 indicates that this is being used for a
// memclr. bulkBarrierPreWrite will pass 0 for the src of each write
// barrier.
//
// Callers should call bulkBarrierPreWrite immediately before
// calling memmove(dst, src, size). This function is marked nosplit
// to avoid being preempted; the GC must not stop the goroutine
// between the memmove and the execution of the barriers.
// The caller is also responsible for cgo pointer checks if this
// may be writing Go pointers into non-Go memory.
//
// The pointer bitmap is not maintained for allocations containing
// no pointers at all; any caller of bulkBarrierPreWrite must first
// make sure the underlying allocation contains pointers, usually
// by checking typ.PtrBytes.
//
// Callers must perform cgo checks if goexperiment.CgoCheck2.
//
// See ${GOROOT}/src/runtime/mbitmap.go#func:bulkBarrierPreWrite
//
//go:nosplit
func bulkBarrierPreWrite(dst, src, size uintptr) {
	if (dst|src|size)&(arch.PtrSize-1) != 0 {
		assert.Throw("bulkBarrierPreWrite:", "unaligned", "arguments")
	}

	if !writeBarrier.needed {
		return
	}

	// TODO
}
