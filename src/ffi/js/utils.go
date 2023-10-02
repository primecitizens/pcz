// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/core/num"
)

func isSigned[T num.Integer]() bool {
	x := T(1)
	x = -x
	return T(x) < 0
}

func checkType[T elem]() (elemSz uint32, signed, float bool) {
	var x T = 1
	elemSz = uint32(unsafe.Sizeof(x))
	if x/3 > 0 {
		return elemSz, true, true
	}

	x = 1
	x = -x
	if T(x) > 0 {
		return elemSz, false, false
	}

	return elemSz, true, false
}

func Must[T any](x T, ok bool) T {
	if ok {
		return x
	}

	assert.Throw("value", "not", "present")
	return x
}

func Bool(b bool) Ref {
	if b {
		return True
	}

	return False
}

func Pointer[T any](p *T) unsafe.Pointer {
	return mark.NoEscapeUnsafePointer(unsafe.Pointer(p))
}

func StringData(s string) unsafe.Pointer {
	return mark.NoEscapeStringDataPointer(s)
}

func SliceData[T any](s []T) unsafe.Pointer {
	return mark.NoEscapeSliceDataPointer(s)
}

func SizeU[T num.Integer | num.Pointer](v T) (x uint32) {
	if x = uint32(uintptr(v)); uintptr(x) != uintptr(v) {
		assert.Throw("overflow")
	}

	return x
}

func SizeI[T num.Integer | num.Pointer](v T) (x int32) {
	if x = int32(uintptr(v)); uintptr(x) != uintptr(v) {
		assert.Throw("overflow")
	}

	return x
}
