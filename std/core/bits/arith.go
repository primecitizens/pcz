// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits

import (
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
)

// Add returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Add(x, y, carry uint) (sum, carryOut uint) {
	if arch.UintBits == 32 {
		s32, c32 := Add32(uint32(x), uint32(y), uint32(carry))
		return uint(s32), uint(c32)
	}
	s64, c64 := Add64(uint64(x), uint64(y), uint64(carry))
	return uint(s64), uint(c64)
}

// Add32 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Add32(x, y, carry uint32) (sum, carryOut uint32) {
	sum64 := uint64(x) + uint64(y) + uint64(carry)
	sum = uint32(sum64)
	carryOut = uint32(sum64 >> 32)
	return
}

// Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Add64(x, y, carry uint64) (sum, carryOut uint64) {
	sum = x + y + carry
	// The sum will overflow if both top bits are set (x & y) or if one of them
	// is (x | y), and a carry from the lower place happened. If such a carry
	// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
	carryOut = ((x & y) | ((x | y) &^ sum)) >> 63
	return
}

// Sub returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Sub(x, y, borrow uint) (diff, borrowOut uint) {
	if arch.UintBits == 32 {
		d32, b32 := Sub32(uint32(x), uint32(y), uint32(borrow))
		return uint(d32), uint(b32)
	}
	d64, b64 := Sub64(uint64(x), uint64(y), uint64(borrow))
	return uint(d64), uint(b64)
}

// Sub32 returns the difference of x, y and borrow, diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32) {
	diff = x - y - borrow
	// The difference will underflow if the top bit of x is not set and the top
	// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
	// from the lower place happens. If that borrow happens, the result will be
	// 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).
	borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 31
	return
}

// Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64) {
	diff = x - y - borrow
	// See Sub32 for the bit logic.
	borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 63
	return
}

// --- Full-width multiply ---

// Mul returns the full-width product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
//
// This function's execution time does not depend on the inputs.
func Mul(x, y uint) (hi, lo uint) {
	if arch.UintBits == 32 {
		h, l := Mul32(uint32(x), uint32(y))
		return uint(h), uint(l)
	}
	h, l := Mul64(uint64(x), uint64(y))
	return uint(h), uint(l)
}

// Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
//
// This function's execution time does not depend on the inputs.
func Mul32(x, y uint32) (hi, lo uint32) {
	tmp := uint64(x) * uint64(y)
	hi, lo = uint32(tmp>>32), uint32(tmp)
	return
}

// Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
//
// This function's execution time does not depend on the inputs.
func Mul64(x, y uint64) (hi, lo uint64) {
	const mask32 = 1<<32 - 1
	x0 := x & mask32
	x1 := x >> 32
	y0 := y & mask32
	y1 := y >> 32
	w0 := x0 * y0
	t := x1*y0 + w0>>32
	w1 := t & mask32
	w2 := t >> 32
	w1 += x0 * y1
	hi = x1*y1 + w2 + w1>>32
	lo = x * y
	return
}

// --- Full-width divide ---

// Div returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div(hi, lo, y uint) (quo, rem uint) {
	if arch.UintBits == 32 {
		q, r := Div32(uint32(hi), uint32(lo), uint32(y))
		return uint(q), uint(r)
	}
	q, r := Div64(uint64(hi), uint64(lo), uint64(y))
	return uint(q), uint(r)
}

// Div32 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div32(hi, lo, y uint32) (quo, rem uint32) {
	if y != 0 && y <= hi {
		assert.Panic(ErrOverflow{})
	}
	z := uint64(hi)<<32 | uint64(lo)
	quo, rem = uint32(z/uint64(y)), uint32(z%uint64(y))
	return
}

// Div64 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div64(hi, lo, y uint64) (quo, rem uint64) {
	if y == 0 {
		assert.Panic(ErrDivideByZero{})
	}
	if y <= hi {
		assert.Panic(ErrOverflow{})
	}

	// If high part is zero, we can directly return the results.
	if hi == 0 {
		return lo / y, lo % y
	}

	s := uint(LeadingZeros64(y))
	y <<= s

	const (
		two32  = 1 << 32
		mask32 = two32 - 1
	)
	yn1 := y >> 32
	yn0 := y & mask32
	un32 := hi<<s | lo>>(64-s)
	un10 := lo << s
	un1 := un10 >> 32
	un0 := un10 & mask32
	q1 := un32 / yn1
	rhat := un32 - q1*yn1

	for q1 >= two32 || q1*yn0 > two32*rhat+un1 {
		q1--
		rhat += yn1
		if rhat >= two32 {
			break
		}
	}

	un21 := un32*two32 + un1 - q1*y
	q0 := un21 / yn1
	rhat = un21 - q0*yn1

	for q0 >= two32 || q0*yn0 > two32*rhat+un0 {
		q0--
		rhat += yn1
		if rhat >= two32 {
			break
		}
	}

	return q1*two32 + q0, (un21*two32 + un0 - q0*y) >> s
}

// Rem returns the remainder of (hi, lo) divided by y. Rem panics for
// y == 0 (division by zero) but, unlike Div, it doesn't panic on a
// quotient overflow.
func Rem(hi, lo, y uint) uint {
	if arch.UintBits == 32 {
		return uint(Rem32(uint32(hi), uint32(lo), uint32(y)))
	}
	return uint(Rem64(uint64(hi), uint64(lo), uint64(y)))
}

// Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics
// for y == 0 (division by zero) but, unlike Div32, it doesn't panic
// on a quotient overflow.
func Rem32(hi, lo, y uint32) uint32 {
	return uint32((uint64(hi)<<32 | uint64(lo)) % uint64(y))
}

// Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics
// for y == 0 (division by zero) but, unlike Div64, it doesn't panic
// on a quotient overflow.
func Rem64(hi, lo, y uint64) uint64 {
	// We scale down hi so that hi < y, then use Div64 to compute the
	// rem with the guarantee that it won't panic on quotient overflow.
	// Given that
	//   hi ≡ hi%y    (mod y)
	// we have
	//   hi<<64 + lo ≡ (hi%y)<<64 + lo    (mod y)
	_, rem := Div64(hi%y, lo, y)
	return rem
}
