// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hashing algorithm inspired by
// wyhash: https://github.com/wangyi-fudan/wyhash/blob/ceb019b530e2c1c14d70b79bfa2bc49de7d95bc1/Modern%20Non-Cryptographic%20Hash%20Function%20and%20Pseudorandom%20Number%20Generator.pdf

//go:build 386 || arm || mips || mipsle

package hash

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/mem"
)

func MemHash(p unsafe.Pointer, seed, s uintptr) uintptr {
	a, b := mix32(uint32(seed), uint32(s^hashkey[0]))
	if s == 0 {
		return uintptr(a ^ b)
	}
	for ; s > 8; s -= 8 {
		a ^= mem.ReadUnaligned32(p)
		b ^= mem.ReadUnaligned32(unsafe.Add(p, 4))
		a, b = mix32(a, b)
		p = unsafe.Add(p, 8)
	}
	if s >= 4 {
		a ^= mem.ReadUnaligned32(p)
		b ^= mem.ReadUnaligned32(unsafe.Add(p, s-4))
	} else {
		t := uint32(*(*byte)(p))
		t |= uint32(*(*byte)(unsafe.Add(p, s>>1))) << 8
		t |= uint32(*(*byte)(unsafe.Add(p, s-1))) << 16
		b ^= t
	}
	a, b = mix32(a, b)
	a, b = mix32(a, b)
	return uintptr(a ^ b)
}

func MemHash32(p unsafe.Pointer, seed uintptr) uintptr {
	a, b := mix32(uint32(seed), uint32(4^hashkey[0]))
	t := mem.ReadUnaligned32(p)
	a ^= t
	b ^= t
	a, b = mix32(a, b)
	a, b = mix32(a, b)
	return uintptr(a ^ b)
}

func MemHash64(p unsafe.Pointer, seed uintptr) uintptr {
	a, b := mix32(uint32(seed), uint32(8^hashkey[0]))
	a ^= mem.ReadUnaligned32(p)
	b ^= mem.ReadUnaligned32(unsafe.Add(p, 4))
	a, b = mix32(a, b)
	a, b = mix32(a, b)
	return uintptr(a ^ b)
}

func mix32(a, b uint32) (uint32, uint32) {
	c := uint64(a^uint32(hashkey[1])) * uint64(b^uint32(hashkey[2]))
	return uint32(c), uint32(c >> 32)
}
