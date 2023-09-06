// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

import (
	"github.com/primecitizens/std/core/bits"
	"github.com/primecitizens/std/core/cmp"
	"github.com/primecitizens/std/core/iter"
)

// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Interface interface {
	// A sortable collection must have finite elements.
	iter.Finite

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Ascend sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Ascend[T Interface](data T) {
	n := data.Len()
	if n <= 1 {
		return
	}
	limit := bits.Len(uint(n))
	pdqsort(data, 0, n, limit)
}

func StableAscend[T Interface](data T) {
	stable(data, data.Len())
}

// IsAscend reports whether data is sorted.
func IsAscend[T Interface](data T) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func SliceAscend[T any](data []T, cmp func(data []T, i, j int) bool) {
	Ascend(SliceSorter[T]{Data: data, LessFunc: cmp})
}

func SliceStableAscend[T any](data []T, cmp func(data []T, i, j int) bool) {
	stable(SliceSorter[T]{Data: data, LessFunc: cmp}, len(data))
}

func SliceIsAscend[T any](data []T, cmp func(data []T, i, j int) bool) bool {
	return IsAscend(SliceSorter[T]{Data: data, LessFunc: cmp})
}

func BuiltinAscend[T cmp.FOrdered](data []T) {
	SliceAscend(data, BuiltinLessFunc[T])
}

func BuiltinIsAscend[T cmp.FOrdered](data []T) bool {
	return SliceIsAscend(data, BuiltinLessFunc[T])
}
