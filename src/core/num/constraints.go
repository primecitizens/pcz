// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package num

import "unsafe"

// Uint is the type union of unsigned integers.
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Int is the type union of signed integers.
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Integer is the type union of both signed and unsigned integers.
type Integer interface {
	Int | Uint
}

// Float is the type union of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Real is the type union of all real number types.
type Real interface {
	Integer | Float
}

// Complex is the type union of all complex types.
type Complex interface {
	~complex64 | ~complex128
}

// Number is the type union of all numeric types.
type Number interface {
	Real | Complex
}

// Pointer represents all numeric types that can be a pointer value
type Pointer interface {
	~uintptr | ~unsafe.Pointer
}
