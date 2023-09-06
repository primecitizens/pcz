// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdint

import (
	"testing"
)

func TestUint128AddSub(t *testing.T) {
	const add1 = 1
	const sub1 = -1
	tests := []struct {
		in   Uint128
		op   int // +1 or -1 to add vs subtract
		want Uint128
	}{
		{Uint128{0, 0}, add1, Uint128{0, 1}},
		{Uint128{0, 1}, add1, Uint128{0, 2}},
		{Uint128{1, 0}, add1, Uint128{1, 1}},
		{Uint128{0, ^uint64(0)}, add1, Uint128{1, 0}},
		{Uint128{^uint64(0), ^uint64(0)}, add1, Uint128{0, 0}},

		{Uint128{0, 0}, sub1, Uint128{^uint64(0), ^uint64(0)}},
		{Uint128{0, 1}, sub1, Uint128{0, 0}},
		{Uint128{0, 2}, sub1, Uint128{0, 1}},
		{Uint128{1, 0}, sub1, Uint128{0, ^uint64(0)}},
		{Uint128{1, 1}, sub1, Uint128{1, 0}},
	}
	for _, tt := range tests {
		var got Uint128
		switch tt.op {
		case add1:
			got = tt.in.AddOne()
		case sub1:
			got = tt.in.SubOne()
		default:
			panic("bogus op")
		}
		if got != tt.want {
			t.Errorf("%v add %d = %v; want %v", tt.in, tt.op, got, tt.want)
		}
	}
}

func TestBitsSetFrom(t *testing.T) {
	tests := []struct {
		bit  uint8
		want Uint128
	}{
		{0, Uint128{^uint64(0), ^uint64(0)}},
		{1, Uint128{^uint64(0) >> 1, ^uint64(0)}},
		{63, Uint128{1, ^uint64(0)}},
		{64, Uint128{0, ^uint64(0)}},
		{65, Uint128{0, ^uint64(0) >> 1}},
		{127, Uint128{0, 1}},
		{128, Uint128{0, 0}},
	}
	for _, tt := range tests {
		var zero Uint128
		got := zero.bitsSetFrom(tt.bit)
		if got != tt.want {
			t.Errorf("0.bitsSetFrom(%d) = %064b want %064b", tt.bit, got, tt.want)
		}
	}
}

func TestBitsClearedFrom(t *testing.T) {
	tests := []struct {
		bit  uint8
		want Uint128
	}{
		{0, Uint128{0, 0}},
		{1, Uint128{1 << 63, 0}},
		{63, Uint128{^uint64(0) &^ 1, 0}},
		{64, Uint128{^uint64(0), 0}},
		{65, Uint128{^uint64(0), 1 << 63}},
		{127, Uint128{^uint64(0), ^uint64(0) &^ 1}},
		{128, Uint128{^uint64(0), ^uint64(0)}},
	}
	for _, tt := range tests {
		ones := Uint128{^uint64(0), ^uint64(0)}
		got := ones.bitsClearedFrom(tt.bit)
		if got != tt.want {
			t.Errorf("ones.bitsClearedFrom(%d) = %064b want %064b", tt.bit, got, tt.want)
		}
	}
}
