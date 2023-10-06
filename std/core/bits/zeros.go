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

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int { return arch.UintBits - Len(x) }

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int { return 8 - Len8(x) }

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int { return 16 - Len16(x) }

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int { return 32 - Len32(x) }

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int { return 64 - Len64(x) }

// See http://supertech.csail.mit.edu/papers/debruijn.pdf
const (
	DeBruijn32      = 0x077CB531
	DeBruijn32Table = "" +
		"\x00\x01\x28\x02\x29\x14\x24\x03\x30\x22\x20\x15\x25\x17\x04\x08" +
		"\x31\x27\x13\x23\x21\x19\x16\x07\x26\x12\x18\x06\x11\x05\x10\x09"
)

const (
	DeBruijn64      = 0x03f79d71b4ca8b09
	DeBruijn64Table = "" +
		"\x00\x01\x56\x02\x57\x49\x28\x03\x61\x58\x42\x50\x38\x29\x17\x04" +
		"\x62\x47\x59\x36\x45\x43\x51\x22\x53\x39\x33\x30\x24\x18\x12\x05" +
		"\x63\x55\x48\x27\x60\x41\x37\x16\x46\x35\x44\x21\x52\x32\x23\x11" +
		"\x54\x26\x40\x15\x34\x20\x31\x10\x25\x14\x19\x09\x13\x08\x07\x06"
)

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
func TrailingZeros(x uint) int {
	if arch.UintBits == 32 {
		return TrailingZeros32(uint32(x))
	}
	return TrailingZeros64(uint64(x))
}

const Ntz8Table = "" +
	"\x08\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x06\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x07\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x06\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x05\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00" +
	"\x04\x00\x01\x00\x02\x00\x01\x00\x03\x00\x01\x00\x02\x00\x01\x00"

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func TrailingZeros8(x uint8) int {
	return int(Ntz8Table[x])
}

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func TrailingZeros16(x uint16) int {
	if x == 0 {
		return 16
	}
	// see comment in TrailingZeros64
	return int(DeBruijn32Table[uint32(x&-x)*DeBruijn32>>(32-5)])
}

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int {
	if x == 0 {
		return 32
	}
	// see comment in TrailingZeros64
	return int(DeBruijn32Table[(x&-x)*DeBruijn32>>(32-5)])
}

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}
	// If popcount is fast, replace code below with return popcount(^x & (x - 1)).
	//
	// x & -x leaves only the right-most bit set in the word. Let k be the
	// index of that bit. Since only a single bit is set, the value is two
	// to the power of k. Multiplying by a power of two is equivalent to
	// left shifting, in this case by k bits. The de Bruijn (64 bit) constant
	// is such that all six bit, consecutive substrings are distinct.
	// Therefore, if we have a left shifted version of this constant we can
	// find by how many bits it was shifted by looking at which six bit
	// substring ended up at the top of the word.
	// (Knuth, volume 4, section 7.3.1)
	return int(DeBruijn64Table[(x&-x)*DeBruijn64>>(64-6)])
}
