// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package stdptr

import (
	"unsafe"

	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/os"
)

const (
	// addrBits is the number of bits needed to represent a virtual address.
	//
	// See heapAddrBits for a table of address space sizes on
	// various architectures. 48 bits is enough for all
	// architectures except s390x.
	//
	// On AMD64, virtual addresses are 48-bit (or 57-bit) numbers sign extended to 64.
	// We shift the address left 16 to eliminate the sign extended part and make
	// room in the bottom for the count.
	//
	// On s390x, virtual addresses are 64-bit. There's not much we
	// can do about this, so we just hope that the kernel doesn't
	// get to really high addresses and panic if it does.
	addrBits = 48

	// In addition to the 16 bits taken from the top, we can take 3 from the
	// bottom, because node must be pointer-aligned, giving a total of 19 bits
	// of count.
	tagBits = 64 - addrBits + 3

	// On AIX, 64-bit addresses are split into 36-bit segment number and 28-bit
	// offset in segment.  Segment numbers in the range 0x0A0000000-0x0AFFFFFFF(LSA)
	// are available for mmap.
	// We assume all tagged addresses are from memory allocated with mmap.
	// We use one bit to distinguish between the two ranges.
	aixAddrBits = 57
	aixTagBits  = 64 - aixAddrBits + 3

	// riscv64 SV57 mode gives 56 bits of userspace VA.
	// tagged pointer code supports it,
	// but broader support for SV57 mode is incomplete,
	// and there may be other issues (see #54104).
	riscv64AddrBits = 56
	riscv64TagBits  = 64 - riscv64AddrBits + 3
)

// The number of bits stored in the numeric tag of a TaggedPointer
const TaggedPointerBits = (os.IsAix * aixTagBits) + (arch.IsRiscv64 * riscv64TagBits) + ((1 - os.IsAix) * (1 - arch.IsRiscv64) * tagBits)

// TaggedPointerPack created a TaggedPointer from a pointer and a tag.
// Tag bits that don't fit in the result are discarded.
func TaggedPointerPack(ptr unsafe.Pointer, tag uintptr) TaggedPointer {
	if os.IsAix == 1 {
		if arch.IsPpc64 == 0 {
			panic("check this code for aix on non-ppc64")
		}
		return TaggedPointer(uint64(uintptr(ptr))<<(64-aixAddrBits) | uint64(tag&(1<<aixTagBits-1)))
	}
	if arch.IsRiscv64 == 1 {
		return TaggedPointer(uint64(uintptr(ptr))<<(64-riscv64AddrBits) | uint64(tag&(1<<riscv64TagBits-1)))
	}
	return TaggedPointer(uint64(uintptr(ptr))<<(64-addrBits) | uint64(tag&(1<<tagBits-1)))
}

// Pointer returns the pointer from a TaggedPointer.
func (tp TaggedPointer) pointer() unsafe.Pointer {
	if arch.IsAmd64 == 1 {
		// amd64 systems can place the stack above the VA hole, so we need to sign extend
		// val before unpacking.
		return unsafe.Pointer(uintptr(int64(tp) >> tagBits << 3))
	}
	if os.IsAix == 1 {
		return unsafe.Pointer(uintptr((tp >> aixTagBits << 3) | 0xa<<56))
	}
	if arch.IsRiscv64 == 1 {
		return unsafe.Pointer(uintptr(tp >> riscv64TagBits << 3))
	}
	return unsafe.Pointer(uintptr(tp >> tagBits << 3))
}

// Tag returns the tag from a TaggedPointer.
func (tp TaggedPointer) tag() uintptr {
	return uintptr(tp & (1<<TaggedPointerBits - 1))
}
