// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mem

import (
	"unsafe" // also for go:linkname

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/num"
)

func memclr(ptr, n uintptr)

// MemclrNoHeapPointers clears n bytes starting at ptr.
//
// Usually you should use typedmemclr. MemclrNoHeapPointers should be
// used only when the caller knows that *ptr contains no heap pointers
// because either:
//
// *ptr is initialized memory and its type is pointer-free, or
//
// *ptr is uninitialized memory (e.g., memory that's being reused
// for a new allocation) and hence contains only "junk".
//
// MemclrNoHeapPointers ensures that if ptr is pointer-aligned, and n
// is a multiple of the pointer size, then any pointer-aligned,
// pointer-sized portion is cleared atomically. Despite the function
// name, this is necessary because this function is the underlying
// implementation of typedmemclr and memclrHasPointers. See the doc of
// memmove for more details.
//
// The (CPU-specific) implementations of this function are in memclr_*.s.
func MemclrNoHeapPointers[T num.Pointer](ptr T, n uintptr) {
	memclr(uintptr(ptr), n)
}

func MemclrBytes(b []byte) {
	memclr(uintptr(unsafe.Pointer(unsafe.SliceData(b))), uintptr(len(b)))
}

func Memclr[T num.Pointer](ptr T, n uintptr) {
	memclr(uintptr(ptr), n)
}

func MemclrHasPointers[T num.Pointer](ptr T, n uintptr) {
	// TODO: free linked heap memory, how?
	memclr(uintptr(ptr), n)
}

func TypedMemClr(typ *abi.Type, dst unsafe.Pointer) {
	if typ.PtrBytes != 0 {
		memclr(uintptr(dst), typ.Size_)
	} else {
		memclr(uintptr(dst), typ.Size_)
	}
}
