// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mem

import (
	"unsafe"

	"github.com/primecitizens/std/core/arch"
)

// Note: These routines perform the read with a native endianness.
func ReadUnaligned32(p unsafe.Pointer) uint32 {
	q := (*[4]byte)(p)
	if arch.BigEndian {
		return uint32(q[3]) | uint32(q[2])<<8 | uint32(q[1])<<16 | uint32(q[0])<<24
	}
	return uint32(q[0]) | uint32(q[1])<<8 | uint32(q[2])<<16 | uint32(q[3])<<24
}

func ReadUnaligned64(p unsafe.Pointer) uint64 {
	q := (*[8]byte)(p)
	if arch.BigEndian {
		return uint64(q[7]) | uint64(q[6])<<8 | uint64(q[5])<<16 | uint64(q[4])<<24 |
			uint64(q[3])<<32 | uint64(q[2])<<40 | uint64(q[1])<<48 | uint64(q[0])<<56
	}
	return uint64(q[0]) | uint64(q[1])<<8 | uint64(q[2])<<16 | uint64(q[3])<<24 | uint64(q[4])<<32 | uint64(q[5])<<40 | uint64(q[6])<<48 | uint64(q[7])<<56
}
