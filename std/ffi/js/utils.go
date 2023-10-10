// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/core/num"
)

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
