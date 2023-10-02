// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wtf16

import (
	. "github.com/primecitizens/std/text/unicode/common"
	"github.com/primecitizens/std/text/unicode/utf16"
	"github.com/primecitizens/std/text/unicode/utf8"
)

func WTF8DecodedSize(s ...uint16) (n int, canDecodeInPlace bool) {
	// use utf16.UTF8DecodedSize is fine because WTF8Decode handles
	// invalid surrogates as 3-byte utf8 append, which is the same
	// as utf8.RuneErrorLen
	return utf16.UTF8DecodedSize(s...)
}

func WTF8DecodeAll(dst []byte, src ...uint16) []byte {
	for n := 0; len(src) != 0; src = src[n:] {
		if dst, n = WTF8Decode(dst, src...); n == 0 {
			// TODO(alloc): grow dst
			return dst
		}
	}

	return dst
}

// WTF8Decode returns the WTF-8 encoding of the potentially ill-formed
// UTF-16 src.
func WTF8Decode(dst []byte, src ...uint16) ([]byte, int) {
	var (
		r     rune
		sz, n int
	)

Loop:
	for ; n < len(src); n++ {
		switch r = rune(src[n]); {
		case r < surr1, surr3 <= r:
			// normal rune
			dst, sz = utf8.EncodeRune(dst, r)
		case surr1 <= r && r < surr2 &&
			n+1 < len(src) &&
			surr2 <= src[n+1] && src[n+1] < surr3:
			// valid surrogate sequence
			n++
			dst, sz = utf8.EncodeRune(dst, utf16.DecodeRune(r, rune(src[n])))
		default:
			// WTF-8 fallback.
			// This only handles the 3-byte case of utf8.AppendRune,
			// as surrogates always fall in that case.

			if cap(dst)-len(dst) < 3 {
				break Loop
			}

			if r > MaxRune {
				r = RuneError
			}

			dst = append(dst, t3|byte(r>>12), tx|byte(r>>6)&maskx, tx|byte(r)&maskx)
		}

		if sz == 0 {
			break
		}
	}

	return dst, n
}
