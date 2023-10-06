// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mark

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	_ "github.com/primecitizens/pcz/std/core/mark/internal/reflect" // import reflect.noescape
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
	return (*T)(NoEscapeUnsafePointer(unsafe.Pointer(p)))
}

//go:linkname NoEscapeUnsafePointer reflect.noescape
//go:noescape
func NoEscapeUnsafePointer(p unsafe.Pointer) unsafe.Pointer

// NoEscapePointer is [NoEscape] but returns unsafe.Pointer.
//
//go:nosplit
func NoEscapePointer[T any](p *T) unsafe.Pointer {
	return NoEscapeUnsafePointer(unsafe.Pointer(p))
}

// NoEscapeSlice is [NoEscape] for slices.
//
//go:nosplit
func NoEscapeSlice[T any](p []T) []T {
	return unsafe.Slice(NoEscape(unsafe.SliceData(p)), cap(p))[:len(p)]
}

// NoEscapeSliceData is [NoEscapeSlice] but returns the pointer
// to the underlay array.
//
//go:nosplit
func NoEscapeSliceData[T any](p []T) *T {
	return NoEscape(unsafe.SliceData(p))
}

// NoEscapeSliceDataPointer is [NoEscapeSliceData] but returns an
// unsafe.Pointer.
//
//go:nosplit
func NoEscapeSliceDataPointer[T any](p []T) unsafe.Pointer {
	return NoEscapeUnsafePointer(unsafe.Pointer(unsafe.SliceData(p)))
}

// NoEscapeString is [NoEscape] for string.
//
//go:nosplit
func NoEscapeString[String ~string](s String) String {
	return String(
		unsafe.String(
			NoEscape(unsafe.StringData(string(s))),
			len(s),
		),
	)
}

// NoEscapeStringData is [NoEscapeString] but returns the pointer to
// the underlay byte array.
//
//go:nosplit
func NoEscapeStringData[String ~string](s String) *byte {
	return NoEscape(unsafe.StringData(string(s)))
}

// NoEscapeStringDataPointer is [NoEscapeStringData] but returns an
// unsafe.Pointer.
//
//go:nosplit
func NoEscapeStringDataPointer[String ~string](s String) unsafe.Pointer {
	return NoEscapeUnsafePointer(unsafe.Pointer(unsafe.StringData(string(s))))
}

// NoEscapeBytesString is [NoEscapeString] but taking bytes as the argument.
func NoEscapeBytesString[T ~byte](s []T) string {
	return unsafe.String(
		(*byte)(NoEscapeUnsafePointer(unsafe.Pointer(unsafe.SliceData(s)))),
		len(s),
	)
}
