// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdint

import (
	"github.com/primecitizens/pcz/std/core/bits"
)

type Int128 Uint128

// Uint128 represents a Uint128 using two uint64s.
//
// When the methods below mention a bit number, bit 0 is the most
// significant bit (in hi) and bit 127 is the lowest (lo&1).
//
// NOTE: Uint128 is kept as struct instead of array because array of size > 1
// doesn't go through SSA, see https://github.com/golang/go/issues/24416
type Uint128 struct {
	hi uint64
	lo uint64
}

func (u Uint128) Equals(o Uint128) bool { return u.hi == o.hi && u.lo == o.lo }

// IsZero reports whether u == 0.
//
// It's faster than u == (uint128{}) because the compiler (as of Go
// 1.15/1.16b1) doesn't do this trick and instead inserts a branch in
// its eq alg's generated code.
func (u Uint128) IsZero() bool { return u.hi|u.lo == 0 }

// AND returns the bitwise AND of u AND m (u&m).
func (u Uint128) AND(m Uint128) Uint128 { return Uint128{hi: u.hi & m.hi, lo: u.lo & m.lo} }

// XOR returns the bitwise XOR of u and m (u^m).
func (u Uint128) XOR(m Uint128) Uint128 { return Uint128{hi: u.hi ^ m.hi, lo: u.lo ^ m.lo} }

// OR returns the bitwise OR of u and m (u|m).
func (u Uint128) OR(m Uint128) Uint128 { return Uint128{hi: u.hi | m.hi, lo: u.lo | m.lo} }

// NOT returns the bitwise NOT of u.
func (u Uint128) NOT() Uint128 { return Uint128{hi: ^u.hi, lo: ^u.lo} }

// SubOne returns u - 1.
func (u Uint128) SubOne() Uint128 {
	lo, borrow := bits.Sub64(u.lo, 1, 0)
	return Uint128{hi: u.hi - borrow, lo: lo}
}

// AddOne returns u + 1.
func (u Uint128) AddOne() Uint128 {
	lo, carry := bits.Add64(u.lo, 1, 0)
	return Uint128{hi: u.hi + carry, lo: lo}
}

func (u Uint128) Add(n Uint128) Uint128 {
	lo, carry := bits.Add64(u.lo, n.lo, 0)
	hi, carry := bits.Add64(u.hi, u.lo, carry)
	if carry != 0 {
		// overflow?
	}
	return Uint128{hi: hi, lo: lo}
}

func (u Uint128) Sub64(n uint64) Uint128 {
	lo, carry := bits.Sub64(u.lo, n, 0)
	return Uint128{hi: u.hi + carry, lo: lo}
}

func (u Uint128) Add64(n uint64) Uint128 {
	lo, carry := bits.Add64(u.lo, n, 0)
	return Uint128{hi: u.hi + carry, lo: lo}
}

// bitsSetFrom returns a copy of u with the given bit
// and all subsequent ones set.
func (u Uint128) bitsSetFrom(bit uint8) Uint128 {
	return u.OR(mask6(int(bit)).NOT())
}

// bitsClearedFrom returns a copy of u with the given bit
// and all subsequent ones cleared.
func (u Uint128) bitsClearedFrom(bit uint8) Uint128 {
	return u.AND(mask6(int(bit)))
}

// mask6 returns a uint128 bitmask with the topmost n bits of a
// 128-bit number.
func mask6(n int) Uint128 {
	return Uint128{hi: ^(^uint64(0) >> n), lo: ^uint64(0) << (128 - n)}
}
