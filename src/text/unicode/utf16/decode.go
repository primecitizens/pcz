// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf16

// DecodeRune returns the UTF-16 decoding of a surrogate pair.
// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
// the Unicode replacement code point U+FFFD.
func DecodeRune(r1, r2 rune) rune {
	if surr1 <= r1 && r1 < surr2 && surr2 <= r2 && r2 < surr3 {
		return (r1-surr1)<<10 | (r2 - surr2) + surrSelf
	}
	return ReplacementChar
}

// Decode returns the Unicode code point sequence represented
// by the UTF-16 encoding s.
func Decode(s []uint16) []rune {
	a := make([]rune, len(s))
	n := 0
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			a[n] = rune(r)
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			a[n] = DecodeRune(rune(r), rune(s[i+1]))
			i++
		default:
			// invalid surrogate sequence
			a[n] = ReplacementChar
		}
		n++
	}
	return a[:n]
}
