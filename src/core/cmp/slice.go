// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"github.com/primecitizens/std/core/mark"
)

type Context[T any] struct {
	A, B T
}

// Compare compares the elements of a and b, using cmp on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
//
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if a == b, -1 if a < b, and +1 if a > b.
func Slice[E any](a, b []E, cmp func(ctx *Context[[]E], i int) int) int {
	ctx := Context[[]E]{
		A: a,
		B: b,
	}

	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			return +1
		}
		if c := cmp(mark.NoEscape(&ctx), i); c != 0 {
			return c
		}
	}
	if len(a) < len(b) {
		return -1
	}
	return 0
}

// SliceEx is like [Slice], but takes an extra arg.
func SliceEx[E, T any](a, b []E, arg T, cmp func(arg T, ctx *Context[[]E], i int) int) int {
	ctx := Context[[]E]{
		A: a,
		B: b,
	}

	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			return +1
		}
		if c := cmp(arg, mark.NoEscape(&ctx), i); c != 0 {
			return c
		}
	}
	if len(a) < len(b) {
		return -1
	}

	return 0
}
