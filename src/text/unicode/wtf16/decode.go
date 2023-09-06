// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wtf16

import (
	"unicode/utf16"
	"unicode/utf8"
)

// EecodeWTF16 returns the WTF-8 encoding of the potentially ill-formed
// UTF-16 s.
func EecodeWTF16(s []uint16, buf []byte) []byte {
	for i := 0; i < len(s); i++ {
		var ar rune
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			ar = rune(r)
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			ar = utf16.DecodeRune(rune(r), rune(s[i+1]))
			i++
		default:
			// WTF-8 fallback.
			// This only handles the 3-byte case of utf8.AppendRune,
			// as surrogates always fall in that case.
			ar = rune(r)
			if ar > utf8.MaxRune {
				ar = utf8.RuneError
			}
			buf = append(buf, t3|byte(ar>>12), tx|byte(ar>>6)&maskx, tx|byte(ar)&maskx)
			continue
		}
		buf = utf8.AppendRune(buf, ar)
	}
	return buf
}
