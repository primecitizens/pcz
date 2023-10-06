// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows UTF-16 strings can contain unpaired surrogates, which can't be
// decoded into a valid UTF-8 string. This file defines a set of functions
// that can be used to encode and decode potentially ill-formed UTF-16 strings
// by using the [the WTF-8 encoding](https://simonsapin.github.io/wtf-8/).
//
// WTF-8 is a strict superset of UTF-8, i.e. any string that is
// well-formed in UTF-8 is also well-formed in WTF-8 and the content
// is unchanged. Also, the conversion never fails and is lossless.
//
// The benefit of using WTF-8 instead of UTF-8 when decoding a UTF-16 string
// is that the conversion is lossless even for ill-formed UTF-16 strings.
// This property allows to read an ill-formed UTF-16 string, convert it
// to a Go string, and convert it back to the same original UTF-16 string.
//
// See go.dev/issues/59971 for more info.
package wtf16

import (
	"github.com/primecitizens/pcz/std/text/unicode/common"
	"github.com/primecitizens/pcz/std/text/unicode/utf16"
)

func AsString(s []uint16) String {
	return String(utf16.AsString(s))
}

type String string

func (s String) Slice() []uint16 {
	return utf16.String(s).Slice()
}

const (
	surr1 = common.SurrogateMin
	surr2 = 0xdc00
	surr3 = common.SurrogateMax + 1

	tx    = 0b10000000
	t3    = 0b11100000
	maskx = 0b00111111
	mask3 = 0b00001111

	rune1Max = 1<<7 - 1
	rune2Max = 1<<11 - 1
)
