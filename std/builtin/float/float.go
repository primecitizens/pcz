// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdfloat

import "unsafe"

var inf = Float64FromBits(0x7FF0000000000000)

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN(f float64) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	return f != f
}

// IsFinite reports whether f is neither NaN nor an infinity.
func IsFinite(f float64) bool {
	return !IsNaN(f - f)
}

// IsInf reports whether f is an infinity.
func IsInf(f float64) bool {
	return !IsNaN(f) && !IsFinite(f)
}

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float64) float64 {
	const sign = 1 << 63
	return Float64FromBits(Float64Bits(x) &^ sign)
}

// CopySign returns a value with the magnitude
// of x and the sign of y.
func CopySign(x, y float64) float64 {
	const sign = 1 << 63
	return Float64FromBits(Float64Bits(x)&^sign | Float64Bits(y)&sign)
}

// Float64Bits returns the IEEE 754 binary representation of f.
func Float64Bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

// Float64FromBits returns the floating point number corresponding
// the IEEE 754 binary representation b.
func Float64FromBits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}
