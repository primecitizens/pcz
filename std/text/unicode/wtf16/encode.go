// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wtf16

import (
	. "github.com/primecitizens/pcz/std/text/unicode/common"
	"github.com/primecitizens/pcz/std/text/unicode/utf16"
	"github.com/primecitizens/pcz/std/text/unicode/utf8"
)

// Encode returns the potentially ill-formed UTF-16 encoding of s.
func Encode(dst []uint16, s string) []uint16 {
	for i := 0; i < len(s); {
		// Cannot use 'for range s' because it expects valid
		// UTF-8 runes.
		r, size := utf8.First(s[i:])
		if r == RuneError {
			// Check if s[i:] contains a valid WTF-8 encoded surrogate.
			if sc := s[i:]; len(sc) >= 3 &&
				sc[0] == 0xED &&
				0xA0 <= sc[1] && sc[1] <= 0xBF &&
				0x80 <= sc[2] && sc[2] <= 0xBF {

				r = rune(sc[0]&mask3)<<12 + rune(sc[1]&maskx)<<6 + rune(sc[2]&maskx)
				dst = append(dst, uint16(r))
				i += 3
				continue
			}
		}
		i += size
		dst = utf16.AppendRunes(dst, r)
	}
	return dst
}
