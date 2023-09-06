// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package rand

type Core interface {
	// Rand32 returns a random uint32.
	Rand32() uint32

	// Rand64 returns a random uint64.
	Rand64() uint64
}

// Rand32n
//
// adapted from golang runtime.fastrandn
// See https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func Rand32n[C Core](core C, n uint32) uint32 {
	return uint32(uint64(core.Rand32()) * uint64(n) >> 32)
}

// Rand64n
func Rand64n[C Core](core C, n uint64) uint64 {
	return core.Rand64() % n
}
