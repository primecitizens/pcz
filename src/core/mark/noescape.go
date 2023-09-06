// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mark

import (
	"unsafe"
)

// NoEscape hides a pointer from escape analysis. USE CAREFULLY!
//
// It is the identity function but escape analysis doesn't think the output
// depends on the input.
//
// It is inlined and currently compiles down to zero instructions.
//
//go:nosplit
func NoEscape[T any](p *T) *T {
	x := uintptr(unsafe.Pointer(p)) // this line is required
	return (*T)(unsafe.Pointer(x ^ 0))
}

// NoEscapeUnsafePointer is [NoEscape] but taking and returning unsafe.Pointer.
//
//go:nosplit
func NoEscapeUnsafePointer(ptr unsafe.Pointer) unsafe.Pointer {
	x := uintptr(ptr) // this line is required
	return unsafe.Pointer(x ^ 0)
}

// NoEscapeSlice is [NoEscape] for slices.
//
//go:nosplit
func NoEscapeSlice[T any](p []T) (ret []T) {
	ret = unsafe.Slice(
		NoEscape(unsafe.SliceData(p)),
		cap(p),
	)
	return ret[:len(p)]
}

// NoescapeSlice is [NoEscape] for strings.
//
//go:nosplit
func NoEscapeString[T ~string](s T) T {
	return T(unsafe.String(
		NoEscape(unsafe.StringData(string(s))),
		len(s),
	))
}

// NoEscapeBytesString is [NoEscapeSlice] for bytes but returning string
// instead.
func NoEscapeBytesString[T ~byte](s []T) string {
	return unsafe.String(
		(*byte)(unsafe.Pointer(NoEscape(unsafe.SliceData(s)))),
		len(s),
	)
}
