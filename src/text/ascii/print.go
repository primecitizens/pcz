// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ascii

const (
	Max = '\u007F'
)

// EqualFold reports whether s and t are equal, ASCII-case-insensitively.
func EqualFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if Lower(s[i]) != Lower(t[i]) {
			return false
		}
	}
	return true
}

// Lower returns the ASCII lowercase version of b.
func Lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// Upper returns the ASCII uppercase version of b.
func Upper(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return b - ('a' - 'A')
	}
	return b
}

// CharIsPrint returns whether b is ASCII and printable according to
// https://tools.ietf.org/html/rfc20#section-4.2.
func CharIsPrint(b byte) bool {
	if b < ' ' || b > '~' {
		return false
	}

	return true
}

func IsPrint(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < ' ' || s[i] > '~' {
			return false
		}
	}
	return true
}

// Valid returns true when s consists of ASCII characters only.
func Valid(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > Max {
			return false
		}
	}
	return true
}
