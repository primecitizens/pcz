// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"github.com/primecitizens/std/core/num"
)

// Ordered is a constraint that permits any ordered type except floating-point
// types: any type that supports the operators < <= >= >.
//
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface{ num.Integer | ~string }

// Less reports whether a is less than b.
func Less[T Ordered](a, b T) bool { return a < b }

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
func Compare[T Ordered](x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}

// FOrdered is Ordered plus floating-point types.
//
// Note that floating-point types may contain NaN ("not-a-number") values.
// An operator such as == or < will always report false when
// comparing a NaN value with any other value, NaN or not.
// See the [FCompare] function for a consistent way to compare NaN values.
type FOrdered interface{ Ordered | num.Float }

// FLess is Less with floating-point types.
//
// Note:
//   - A NaN is considered less than any non-NaN
//   - -0.0 is not less than (is equal to) 0.0.
func FLess[T FOrdered](x, y T) bool { return (isNaN(x) && !isNaN(y)) || x < y }

// FCompare is Compare with floating-point types allowed.
//
// Note:
//   - A NaN is considered less than any non-NaN
//   - A NaN is considered equal to a NaN
//   - and -0.0 is equal to 0.0.
func FCompare[T FOrdered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}

	return 0
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T FOrdered](x T) bool { return x != x }
