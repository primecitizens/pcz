// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package num

import (
	"unsafe"
)

func IsSignedType[T Real]() bool {
	x := T(1)
	x = -x
	return x < 0
}

func IsSigned[T Real](T) bool {
	return IsSignedType[T]()
}

func IsFloatType[T Real]() bool {
	x := T(1)
	x /= 3
	return x > 0
}

func IsFloat[T Real](T) bool {
	return IsFloatType[T]()
}

func CheckType[T Real]() (szBytes uintptr, float, signed bool) {
	var x T = 1
	szBytes = unsafe.Sizeof(x)
	if x/3 > 0 {
		return szBytes, true, true
	}

	x = 1
	x = -x
	if T(x) > 0 {
		return szBytes, false, false
	}

	return szBytes, true, false
}

func Check[T Real](T) (szBytes uintptr, float, signed bool) {
	return CheckType[T]()
}
