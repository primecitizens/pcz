// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf8

import (
	. "github.com/primecitizens/std/text/unicode/common"
)

// functions used by the runtime

// IterRune returns the rune at the start of p[pos:] and the index
// after the rune in p.
//
// If the string appears to be incomplete or decoding problems are
// encountered (RuneError, pos+1) is returned to ensure
// progress when IterRune is used to iterate over a string.
func IterRune(p string, pos int) (r rune, next int) {
	n := len(p) - pos
	if n < 1 {
		return RuneError, pos + 1
	}

	p = p[pos:]

	p0 := p[0]
	x := first[p0]
	if x >= as {
		// The following code simulates an additional check for x == xx and
		// handling the ASCII and invalid cases accordingly. This mask-and-or
		// approach prevents an additional branch.
		mask := rune(x) << 31 >> 31 // Create 0x0000 or 0xFFFF.
		return rune(p0)&^mask | RuneError&mask, pos + 1
	}

	sz := int(x & 7)
	if n < sz {
		return RuneError, pos + 1
	}

	b1 := p[1]
	accept := acceptRanges[x>>4]
	if b1 < accept.lo || accept.hi < b1 {
		return RuneError, pos + 1
	}

	if sz <= 2 { // <= instead of == to help the compiler eliminate some bounds checks
		return rune(p0&mask2)<<6 | rune(b1&maskx), pos + 2
	}

	b2 := p[2]
	if b2 < locb || hicb < b2 {
		return RuneError, pos + 1
	}

	if sz <= 3 {
		return rune(p0&mask3)<<12 | rune(b1&maskx)<<6 | rune(b2&maskx), pos + 3
	}

	b3 := p[3]
	if b3 < locb || hicb < b3 {
		return RuneError, pos + 1
	}

	return rune(p0&mask4)<<18 | rune(b1&maskx)<<12 | rune(b2&maskx)<<6 | rune(b3&maskx), pos + 4
}
