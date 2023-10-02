// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf8

import (
	"github.com/primecitizens/std/core/assert"
	. "github.com/primecitizens/std/text/unicode/common"
)

// RuneLen returns the number of bytes required to encode the rune.
//
// It returns 0 if the rune is not a valid value to encode in UTF-8.
func RuneLen(r rune) int {
	if r >= 0 && r <= Rune1Max {
		return 1
	}

	return RuneLenNonASCII(r)
}

func RuneLenNonASCII(r rune) int {
	switch {
	case r < 0:
	case r <= Rune2Max:
		return 2
	case SurrogateMin <= r && r <= SurrogateMax:
	case r <= Rune3Max:
		return 3
	case r <= MaxRune:
		return 4
	}
	return 0
}

// EncodeRune writes into dst the UTF-8 encoding of the rune.
//
// If the rune is out of range, it writes the encoding of RuneError.
//
// It returns the number of bytes written, and when dst is too small to
// write the complete byte sequence of the rune, it returns 0.
func EncodeRune(dst []byte, r rune) ([]byte, int) {
	if len(dst) == cap(dst) {
		return dst, 0
	} else if uint32(r) <= Rune1Max {
		return append(dst, byte(r)), 1
	}

	return EncodeRuneNonASCII(dst, r)
}

func EncodeRuneNonASCII(dst []byte, r rune) ([]byte, int) {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch u := uint32(r); {
	case u <= Rune2Max:
		return append(dst, t2|byte(r>>6), tx|byte(r)&maskx), 2
	case u > MaxRune,
		SurrogateMin <= u && u <= SurrogateMax:
		r = RuneError
		fallthrough
	case u <= Rune3Max:
		return append(dst, t3|byte(r>>12), tx|byte(r>>6)&maskx, tx|byte(r)&maskx), 3
	default:
		return append(dst, t4|byte(r>>18), tx|byte(r>>12)&maskx, tx|byte(r>>6)&maskx, tx|byte(r)&maskx), 4
	}
}

// AppendRunes appends the UTF-8 encoding of r to the end of p and
// returns the extended buffer. If the rune is out of range,
// it appends the encoding of RuneError.
func AppendRunes(dst []byte, src ...rune) []byte {
	var n int
	for _, r := range src {
		if dst, n = EncodeRune(dst, r); n == 0 {
			assert.TODO("grow")
		}
	}

	return dst
}
