// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf16

import (
	"unsafe"

	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/bits"
	. "github.com/primecitizens/std/text/unicode/common"
)

func AsString(s []uint16) String {
	if arch.BigEndian {
		for i, x := range s {
			s[i] = bits.ReverseBytes16(x)
		}
	}

	return String(
		unsafe.String(
			(*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*2,
		),
	)
}

// A String represents a UTF-16 encoded string.
//
// NOTE: DO NOT use as go string.
type String string

func (s String) Slice() []uint16 {
	ret := unsafe.Slice(
		(*uint16)(unsafe.Pointer(unsafe.StringData(string(s)))), len(s)/2,
	)

	if arch.BigEndian {
		for i, x := range ret {
			ret[i] = bits.ReverseBytes16(x)
		}
	}

	return ret
}

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair.
	// the value is those 20 bits plus 0x10000.
	surr1 = SurrogateMin
	surr2 = 0xdc00
	surr3 = SurrogateMax + 1

	surrSelf = 0x10000
)

// IsSurrogate reports whether the specified Unicode code point
// can appear in a surrogate pair.
func IsSurrogate(r rune) bool {
	return surr1 <= r && r < surr3
}
