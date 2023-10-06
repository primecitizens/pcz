// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package cmp

import (
	"bytes"
	"unsafe"
)

func Bytes(a, b []byte) int {
	return bytes.Compare(a, b)
}

func String(a, b string) int {
	return bytes.Compare(
		unsafe.Slice(unsafe.StringData(a), len(a)),
		unsafe.Slice(unsafe.StringData(b), len(b)),
	)
}
