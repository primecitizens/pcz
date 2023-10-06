// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf16

import (
	"github.com/primecitizens/pcz/std/core/assert"
	. "github.com/primecitizens/pcz/std/text/unicode/common"
	"github.com/primecitizens/pcz/std/text/unicode/utf8"
)

// UTF8DecodedSize returns UTF-8 bytes required to store the decoded
// UTF-16 encoding of unicode code points.
func UTF8DecodedSize(utf16 ...uint16) (n int, canDecodeInPlace bool) {
	var (
		i  int
		ar rune
	)

	// when canDecodeInPlace is true: i*2 >= n
	for canDecodeInPlace = true; i < len(utf16); {
		switch ar = rune(utf16[i]); {
		case ar < surr1, surr3 <= ar:
			// normal rune
			i++
			n += utf8.RuneLen(ar)
			if canDecodeInPlace && i*2 /* space avail */ < n /* space required */ {
				canDecodeInPlace = false
			}
		case surr1 <= ar && ar < surr2 &&
			i+1 < len(utf16) &&
			surr2 <= utf16[i+1] && utf16[i+1] < surr3:
			// valid surrogate sequence
			// in this case, we have 4-bytes available for utf-8, so it
			// can always decode to utf8 in-place, no need to update
			// canDecodeInPlace
			i++
			n += utf8.RuneLen(DecodeRune(ar, rune(utf16[i])))
			i++
		default:
			// invalid surrogate sequence
			i++
			n += utf8.RuneErrorLen
			if canDecodeInPlace && i*2 /* space avail */ < n /* space required */ {
				canDecodeInPlace = false
			}
		}
	}

	return
}

// DecodeRune returns the UTF-16 decoding of a surrogate pair.
// If the pair is not a valid UTF-16 surrogate pair, DecodeRune returns
// the Unicode replacement code point U+FFFD.
func DecodeRune(r1, r2 rune) rune {
	if surr1 <= r1 && r1 < surr2 && surr2 <= r2 && r2 < surr3 {
		return (r1-surr1)<<10 | (r2 - surr2) + surrSelf
	}
	return RuneError
}

// UTF8Decode transcodes UTF-16 encoding of unicode code points
// to UTF-8 bytes, returns updated dst and count of consumed
// uint16s from src.
func UTF8Decode(dst []byte, utf16 ...uint16) ([]byte, int) {
	var (
		sz, n int
		r     rune
	)

	for ; n < len(utf16); n++ {
		switch r = rune(utf16[n]); {
		case r < surr1, surr3 <= r:
			// normal rune
			dst, sz = utf8.EncodeRune(dst, r)
		case surr1 <= r && r < surr2 &&
			n+1 < len(utf16) &&
			surr2 <= utf16[n+1] && utf16[n+1] < surr3:
			// valid surrogate sequence
			n++
			dst, sz = utf8.EncodeRune(dst, DecodeRune(r, rune(utf16[n])))
		default:
			// invalid surrogate sequence
			dst, sz = utf8.EncodeRune(dst, RuneError)
		}

		if sz == 0 {
			break
		}
	}

	return dst, n
}

func UTF8Append(dst []byte, src ...uint16) []byte {
	for n := 0; len(src) != 0; src = src[n:] {
		if dst, n = UTF8Decode(dst, src...); n == 0 {
			assert.TODO("grow")
		}
	}

	return dst
}

// RunesDecode decodes UTF-16 code points in src into the dst, returns
// updated dst and count of consumed uint16s from src.
func RunesDecode(d []rune, src ...uint16) (dst []rune, n int) {
	var (
		r   rune
		off = len(d)
	)
	dst = d[:cap(d)]

	for n = 0; n < len(src) && off < len(dst); n++ {
		switch r = rune(src[n]); {
		case r < surr1, surr3 <= r:
			// normal rune
		case surr1 <= r && r < surr2 &&
			n+1 < len(src) &&
			surr2 <= src[n+1] && src[n+1] < surr3:
			// valid surrogate sequence
			n++
			r = DecodeRune(r, rune(src[n]))
		default:
			// invalid surrogate sequence
			r = RuneError
		}

		dst[off] = r
		off++
	}

	return dst[:off], n
}

// RunesAppend decodes all UTF-16 code points in src into the dst, returns
// the decoded runes.
func RunesAppend(dst []rune, utf16 ...uint16) []rune {
	for n := 0; len(utf16) != 0; utf16 = utf16[n:] {
		if dst, n = RunesDecode(dst, utf16...); n == 0 {
			assert.TODO("grow")
		}
	}

	return dst
}
