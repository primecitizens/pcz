// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf16

// The conditions replacementChar==unicode.ReplacementChar and
// maxRune==unicode.MaxRune are verified in the tests.
// Defining them locally avoids this package depending on package unicode.

const (
	ReplacementChar = '\uFFFD'     // Unicode replacement character
	MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
)

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair.
	// the value is those 20 bits plus 0x10000.
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000

	surrSelf = 0x10000
)

// IsSurrogate reports whether the specified Unicode code point
// can appear in a surrogate pair.
func IsSurrogate(r rune) bool {
	return surr1 <= r && r < surr3
}
