// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utf16 implements encoding and decoding of UTF-16 sequences.
package utf16

import (
	. "github.com/primecitizens/std/text/unicode/common"
)

// EncodedSize returns the number of uint16s required to hold the UTF-16
// encoding of unicode code points in s.
func EncodedSize(s string) int {
	n := len(s)
	for _, v := range s {
		if v >= surrSelf {
			n++
		}
	}

	return n
}

// EncodeRune returns the UTF-16 surrogate pair r1, r2 for the given rune.
//
// If the rune is not a valid Unicode code point or does not need encoding,
// EncodeRune returns U+FFFD, U+FFFD.
func EncodeRune(r rune) (r1, r2 rune) {
	if r < surrSelf || r > MaxRune {
		return RuneError, RuneError
	}
	r -= surrSelf
	return surr1 + (r>>10)&0x3ff, surr2 + r&0x3ff
}

// AppendRunes appends the UTF-16 encoding of the Unicode
// code point sequence src to dst.
func AppendRunes(dst []uint16, src ...rune) []uint16 {
	for n := 0; len(src) != 0; src = src[n:] {
		if dst, n = EncodeRunes(dst, src...); n == 0 {
			// TODO(alloc): grow dst
			return dst
		}
	}

	return dst
}

func EncodeRunes(dst []uint16, s ...rune) ([]uint16, int) {
	var (
		pos = len(dst)
		i   int
		v   rune
	)
	dst = dst[:cap(dst)]

	for i = 0; i < len(s) && pos < len(dst); i++ {
		switch v = s[i]; {
		case 0 <= v && v < surr1, surr3 <= v && v < surrSelf:
			// normal rune
			dst[pos] = uint16(v)
			pos++
		case surrSelf <= v && v <= MaxRune:
			// needs surrogate sequence
			r1, r2 := EncodeRune(v)
			dst[pos] = uint16(r1)
			dst[pos+1] = uint16(r2)
			pos += 2
		default:
			dst[pos] = uint16(RuneError)
			pos++
		}
	}

	return dst[:pos], i
}
