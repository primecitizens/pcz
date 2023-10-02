// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	stdslice "github.com/primecitizens/std/builtin/slice"
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/asan"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/bits"
	"github.com/primecitizens/std/core/math"
	"github.com/primecitizens/std/core/mem"
	"github.com/primecitizens/std/core/msan"
	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/core/os"
	"github.com/primecitizens/std/core/race"
)

// see $GOROOT/src/runtime/slice.go

func makeslice(et *abi.Type, len, cap int) unsafe.Pointer {
	return alloc.MakeTyped(getg().G().DefaultAlloc(), et, len, cap)
}

func makeslice64(et *abi.Type, len64, cap64 int64) unsafe.Pointer {
	len := int(len64)
	if int64(len) != len64 {
		panicmakeslicelen()
	}

	cap := int(cap64)
	if int64(cap) != cap64 {
		panicmakeslicecap()
	}

	return makeslice(et, len, cap)
}

func makeslicecopy(et *abi.Type, tolen int, fromlen int, from unsafe.Pointer) unsafe.Pointer {
	var tomem, copymem uintptr
	if uintptr(tolen) > uintptr(fromlen) {
		var overflow bool
		tomem, overflow = math.MulUintptr(et.Size_, uintptr(tolen))
		if overflow || tomem > os.MaxAlloc || tolen < 0 {
			panicmakeslicelen()
		}
		copymem = et.Size_ * uintptr(fromlen)
	} else {
		// fromlen is a known good length providing and equal or greater than tolen,
		// thereby making tolen a good slice length too as from and to slices have the
		// same element width.
		tomem = et.Size_ * uintptr(tolen)
		copymem = tomem
	}

	var to unsafe.Pointer
	if et.PtrBytes == 0 {
		to = explicit_mallocgc(nil, tomem, false)
		if copymem < tomem {
			mem.Clear(unsafe.Add(to, copymem), tomem-copymem)
		}
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		to = explicit_mallocgc(et, uintptr(tolen), true)
		if copymem > 0 && writeBarrier.enabled {
			// Only shade the pointers in old.array since we know the destination slice to
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(to), uintptr(from), copymem)
		}
	}

	if race.Enabled {
		callerpc := getcallerpc()
		pc := abi.FuncPCABIInternal(makeslicecopy)
		race.ReadRangePC(from, copymem, callerpc, pc)
	}
	if msan.Enabled {
		msan.Read(from, copymem)
	}
	if asan.Enabled {
		asan.Read(from, copymem)
	}

	mem.Move(to, from, copymem)
	return to
}

// growslice allocates new backing store for a slice.
//
// arguments:
//
//	oldPtr = pointer to the slice's backing array
//	newLen = new length (= oldLen + num)
//	oldCap = original slice's capacity.
//	   num = number of elements being added
//	    et = element type
//
// return values:
//
//	newPtr = pointer to the new backing store
//	newLen = same value as the argument
//	newCap = capacity of the new backing store
//
// Requires that uint(newLen) > uint(oldCap).
// Assumes the original slice length is newLen - num
//
// A new backing store is allocated with space for at least newLen elements.
// Existing entries [0, oldLen) are copied over to the new backing store.
// Added entries [oldLen, newLen) are not initialized by growslice
// (although for pointer-containing element types, they are zeroed). They
// must be initialized by the caller.
// Trailing entries [newLen, newCap) are zeroed.
//
// growslice's odd calling convention makes the generated code that calls
// this function simpler. In particular, it accepts and returns the
// new length so that the old length is not live (does not need to be
// spilled/restored) and the new length is returned (also does not need
// to be spilled/restored).
func growslice(oldPtr unsafe.Pointer, newLen, oldCap, n int, et *abi.Type) stdslice.Header {
	oldLen := newLen - n
	if race.Enabled {
		callerpc := getcallerpc()
		race.ReadRangePC(oldPtr, uintptr(oldLen*int(et.Size_)), callerpc, abi.FuncPCABIInternal(growslice))
	}
	if msan.Enabled {
		msan.Read(oldPtr, uintptr(oldLen*int(et.Size_)))
	}
	if asan.Enabled {
		asan.Read(oldPtr, uintptr(oldLen*int(et.Size_)))
	}

	if newLen < 0 {
		assert.Throw("growslice:", "len", "out", "of", "range")
	}

	if et.Size_ == 0 {
		// append should not create a slice with nil pointer but non-zero len.
		// We assume that append doesn't need to preserve oldPtr in this case.
		return stdslice.Header{Array: alloc.ZeroSized(), Len: newLen, Cap: newLen}
	}

	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		newcap = newLen
	} else {
		const threshold = 256
		if oldCap < threshold {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < newLen {
				// Transition from growing 2x for small slices
				// to growing 1.25x for large slices. This formula
				// gives a smooth-ish transition between the two.
				newcap += (newcap + 3*threshold) / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = newLen
			}
		}
	}

	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.Size.
	// For 1 we don't need any division/multiplication.
	// For arch.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	case et.Size_ == 1:
		lenmem = uintptr(oldLen)
		newlenmem = uintptr(newLen)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > os.MaxAlloc
		newcap = int(capmem)
	case et.Size_ == arch.PtrSize:
		lenmem = uintptr(oldLen) * arch.PtrSize
		newlenmem = uintptr(newLen) * arch.PtrSize
		capmem = roundupsize(uintptr(newcap) * arch.PtrSize)
		overflow = uintptr(newcap) > os.MaxAlloc/arch.PtrSize
		newcap = int(capmem / arch.PtrSize)
	case num.IsPowerOfTwo(et.Size_):
		var shift uintptr
		if arch.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(bits.TrailingZeros64(uint64(et.Size_))) & 63
		} else {
			shift = uintptr(bits.TrailingZeros32(uint32(et.Size_))) & 31
		}
		lenmem = uintptr(oldLen) << shift
		newlenmem = uintptr(newLen) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (os.MaxAlloc >> shift)
		newcap = int(capmem >> shift)
		capmem = uintptr(newcap) << shift
	default:
		lenmem = uintptr(oldLen) * et.Size_
		newlenmem = uintptr(newLen) * et.Size_
		capmem, overflow = math.MulUintptr(et.Size_, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.Size_)
		capmem = uintptr(newcap) * et.Size_
	}

	// The check of overflow in addition to capmem > os.MaxAlloc is needed
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if overflow || capmem > os.MaxAlloc {
		assert.Throw("growslice:", "len", "out", "of", "range")
	}

	var p unsafe.Pointer
	if et.PtrBytes == 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from oldLen to newLen.
		// Only clear the part that will not be overwritten.
		// The reflect_growslice() that calls growslice will manually clear
		// the region not cleared here.
		mem.Clear(unsafe.Add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// Only shade the pointers in oldPtr since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(oldPtr), lenmem-et.Size_+et.PtrBytes)
		}
	}
	mem.Move(p, oldPtr, lenmem)

	return stdslice.Header{Array: p, Len: newLen, Cap: newcap}
}

// slicecopy is used to copy from a string or slice of pointerless elements into a slice.
func slicecopy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, width uintptr) int {
	if fromLen == 0 || toLen == 0 {
		return 0
	}

	n := fromLen
	if toLen < n {
		n = toLen
	}

	if width == 0 {
		return n
	}

	size := uintptr(n) * width
	if race.Enabled {
		callerpc := getcallerpc()
		pc := abi.FuncPCABIInternal(slicecopy)
		race.ReadRangePC(fromPtr, size, callerpc, pc)
		race.WriteRangePC(toPtr, size, callerpc, pc)
	}
	if msan.Enabled {
		msan.Read(fromPtr, size)
		msan.Write(toPtr, size)
	}
	if asan.Enabled {
		asan.Read(fromPtr, size)
		asan.Write(toPtr, size)
	}

	if size == 1 { // common case worth about 2x to do here
		// TODO: is this still worth it with new memmove impl?
		*(*byte)(toPtr) = *(*byte)(fromPtr) // known to be a byte pointer
	} else {
		mem.Move(toPtr, fromPtr, size)
	}
	return n
}

func panicmakeslicelen() { assert.Throw("makeslice:", "len", "out", "of", "range") }
func panicmakeslicecap() { assert.Throw("makeslice:", "cap", "out", "of", "range") }
