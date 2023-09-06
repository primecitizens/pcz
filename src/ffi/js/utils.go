// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/std/core/debug"
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/ffi/js/bindings"
)

func Bool(b bool) Ref {
	if b {
		return True
	}

	return False
}

func Pointer[T any](p *T) bindings.Uintptr32 {
	return bindings.Uintptr32(SizeU(
		uintptr(unsafe.Pointer(mark.NoEscape(p))),
	))
}

func StringData(s string) bindings.Uintptr32 {
	return bindings.Uintptr32(SizeU(
		uintptr(
			unsafe.Pointer(
				mark.NoEscape(unsafe.StringData(s)),
			),
		),
	))
}

func SliceData[T any](s []T) bindings.Uintptr32 {
	return bindings.Uintptr32(
		SizeU(
			uintptr(
				unsafe.Pointer(
					mark.NoEscape(unsafe.SliceData(s)),
				),
			),
		),
	)
}

func SizeU[T num.Integer | num.Pointer](v T) (x uint32) {
	if x = uint32(uintptr(v)); uintptr(x) != uintptr(v) {
		debug.Abort()
	}

	return x
}

func SizeI[T num.Integer | num.Pointer](v T) (x int32) {
	if x = int32(uintptr(v)); uintptr(x) != uintptr(v) {
		debug.Abort()
	}

	return x
}
