// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package bytealg

import (
	"bytes"
	"strings"
)

func CountSlice(b []byte, c byte) (n int) {
	for _, x := range b {
		if x == c {
			n++
		}
	}

	return
}

func IndexSlice(s, sep []byte) int {
	return bytes.Index(s, sep)
}

func IndexSliceByte(b []byte, c byte) int {
	return bytes.IndexByte(b, c)
}

func Count(s string, c byte) (n int) {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			n++
		}
	}

	return
}

func Index(s, substr string) int {
	return strings.Index(s, substr)
}

func IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}
