// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdstring

import (
	"unsafe"
)

type Header struct {
	Str unsafe.Pointer
	Len int
}

// FromBytes converts a byte slice to string without allocation
func FromBytes[T ~byte](b []T) string {
	return unsafe.String((*byte)(unsafe.Pointer(unsafe.SliceData(b))), len(b))
}

// FromByteArray converts a sequance of NULL terminated bytes into string
// assuming UTF-8 encoding.
//
//go:nosplit
func FromByteArray(arr *byte) string {
	return unsafe.String(arr, FindNULL(arr))
}

//go:nosplit
func ToBytes[T ~byte](s string) []T {
	return unsafe.Slice((*T)(unsafe.Pointer(unsafe.StringData(s))), len(s))
}
