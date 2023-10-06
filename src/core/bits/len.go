// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits

import (
	"github.com/primecitizens/pcz/std/core/arch"
)

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len(x uint) int {
	if arch.UintBits == 32 {
		return Len32(uint32(x))
	}
	return Len64(uint64(x))
}

const Len8tab = "" +
	"\x00\x01\x02\x02\x03\x03\x03\x03\x04\x04\x04\x04\x04\x04\x04\x04" +
	"\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08"

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len8(x uint8) int {
	return int(Len8tab[x])
}

// Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len16(x uint16) (n int) {
	if x >= 1<<8 {
		x >>= 8
		n = 8
	}
	return n + int(Len8tab[x])
}

// Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len32(x uint32) (n int) {
	if x >= 1<<16 {
		x >>= 16
		n = 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(Len8tab[x])
}

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(Len8tab[x])
}
