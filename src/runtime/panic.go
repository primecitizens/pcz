// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/debug"
)

// see $GOROOT/src/runtime/panic.go

func throwinit() {
	// TODO
}

func panicdivide() { assert.Throw("integer", "divide", "by", "zero") }
func panicshift()  { assert.Throw("negative", "shift", "amount") }

// panicwrap generates a panic for a call to a wrapped value method
// with a nil pointer receiver.
//
// It is called from the generated wrapper code.
func panicwrap() {
	pc := getcallerpc()
	println("value", "method", abi.FindFunc(pc).Name(), "called", "using", "nil", "pointer")
	debug.Abort()
}

// failures in the comparisons for s[x], 0 <= x < y (y == len(s))
func goPanicIndex(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsIndex})
}

func goPanicIndexU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsIndex})
}

// failures in the comparisons for s[:x], 0 <= x <= y (y == len(s) or cap(s))
func goPanicSliceAlen(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSliceAlen})
}

func goPanicSliceAlenU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSliceAlen})
}

func goPanicSliceAcap(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSliceAcap})
}

func goPanicSliceAcapU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSliceAcap})
}

// failures in the comparisons for s[x:y], 0 <= x <= y
func goPanicSliceB(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSliceB})
}

func goPanicSliceBU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSliceB})
}

// failures in the comparisons for s[::x], 0 <= x <= y (y == len(s) or cap(s))
func goPanicSlice3Alen(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSlice3Alen})
}

func goPanicSlice3AlenU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSlice3Alen})
}

func goPanicSlice3Acap(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSlice3Acap})
}

func goPanicSlice3AcapU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSlice3Acap})
}

func goPanicSlice3B(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSlice3B})
}

func goPanicSlice3BU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSlice3B})
}

func goPanicSlice3C(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsSlice3C})
}

func goPanicSlice3CU(x uint, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: false, y: y, code: boundsSlice3C})
}

// failures in the conversion ([x]T)(s) or (*[x]T)(s), 0 <= x <= y, y == len(s)
func goPanicSliceConvert(x int, y int) {
	panicBoundsError(boundsError{x: int64(x), signed: true, y: y, code: boundsConvert})
}

func panicBoundsError(e boundsError) {
	if e.code == boundsConvert {
		println(
			"runtime", "error:", "cannot", "convert", "slice", "with", "length", e.y,
			"to", "array", "or", "pointer", "to", "array", "with", "length", e.x,
		)
		debug.Abort()
		return
	}

	if e.signed && e.x < 0 {
		before, after := negBoundsErrorCode(e.code).fmt()
		println("runtime", "error:", before, e.x, after)
		debug.Abort()
		return
	}

	before, between, after := e.code.fmt()
	println("runtime", "error:", before, e.x, between, e.y, after)
	debug.Abort()
	return
}

// A boundsError represents an indexing or slicing operation gone wrong.
type boundsError struct {
	x int64
	y int
	// Values in an index or slice expression can be signed or unsigned.
	// That means we'd need 65 bits to encode all possible indexes, from -2^63 to 2^64-1.
	// Instead, we keep track of whether x should be interpreted as signed or unsigned.
	// y is known to be nonnegative and to fit in an int.
	signed bool
	code   boundsErrorCode
}

const (
	boundsIndex = iota // s[x], 0 <= x < len(s) failed

	boundsSliceAlen // s[?:x], 0 <= x <= len(s) failed
	boundsSliceAcap // s[?:x], 0 <= x <= cap(s) failed
	boundsSliceB    // s[x:y], 0 <= x <= y failed (but boundsSliceA didn't happen)

	boundsSlice3Alen // s[?:?:x], 0 <= x <= len(s) failed
	boundsSlice3Acap // s[?:?:x], 0 <= x <= cap(s) failed
	boundsSlice3B    // s[?:x:y], 0 <= x <= y failed (but boundsSlice3A didn't happen)
	boundsSlice3C    // s[x:y:?], 0 <= x <= y failed (but boundsSlice3A/B didn't happen)

	boundsConvert // (*[x]T)(s), 0 <= x <= len(s) failed
	// Note: in the above, len(s) and cap(s) are stored in y
)

type boundsErrorCode uint8

func (c boundsErrorCode) fmt() (before, between, after string) {
	switch c {
	case boundsIndex:
		return "index out of range [", "] with length", ""
	case boundsSliceAlen:
		return "slice bounds out of range [:", "] with length", ""
	case boundsSliceAcap:
		return "slice bounds out of range [:", "] with capacity", ""
	case boundsSliceB:
		return "slice bounds out of range [", ":", "]"
	case boundsSlice3Alen:
		return "slice bounds out of range [::", "] with length", ""
	case boundsSlice3Acap:
		return "slice bounds out of range [::", "] with capacity", ""
	case boundsSlice3B:
		return "slice bounds out of range [:", ":", "]"
	case boundsSlice3C:
		return "slice bounds out of range [", ":", ":]"
	default:
		assert.Unreachable()
		return
	}
}

type negBoundsErrorCode boundsErrorCode

func (c negBoundsErrorCode) fmt() (before, after string) {
	switch c {
	case boundsIndex:
		return "index out of range [", "]"
	case boundsSliceAlen:
		return "slice bounds out of range [:", "]"
	case boundsSliceAcap:
		return "slice bounds out of range [:", "]"
	case boundsSliceB:
		return "slice bounds out of range [", ":]"
	case boundsSlice3Alen:
		return "slice bounds out of range [::", "]"
	case boundsSlice3Acap:
		return "slice bounds out of range [::", "]"
	case boundsSlice3B:
		return "slice bounds out of range [:", ":]"
	case boundsSlice3C:
		return "slice bounds out of range [", "::]"
	default:
		assert.Unreachable()
		return
	}
}

// Implemented in assembly, as they take arguments in registers.
// Declared here to mark them as ABIInternal.

func panicIndex(x int, y int)
func panicIndexU(x uint, y int)
func panicSliceAlen(x int, y int)
func panicSliceAlenU(x uint, y int)
func panicSliceAcap(x int, y int)
func panicSliceAcapU(x uint, y int)
func panicSliceB(x int, y int)
func panicSliceBU(x uint, y int)
func panicSlice3Alen(x int, y int)
func panicSlice3AlenU(x uint, y int)
func panicSlice3Acap(x int, y int)
func panicSlice3AcapU(x uint, y int)
func panicSlice3B(x int, y int)
func panicSlice3BU(x uint, y int)
func panicSlice3C(x int, y int)
func panicSlice3CU(x uint, y int)
func panicSliceConvert(x int, y int)
