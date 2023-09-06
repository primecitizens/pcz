// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sort

import (
	"github.com/primecitizens/std/core/cmp"
	"github.com/primecitizens/std/core/num"
)

func BuiltinLessFunc[T cmp.FOrdered](s []T, i, j int) bool {
	return cmp.FLess(s[i], s[j])
}

func BuiltinCmpFunc[T num.Integer](s []T, i, j int) int {
	return cmp.FCompare(s[i], s[j])
}

type ReverseSorter[T Interface] struct {
	Inner T
}

// Len implements Interface
func (r ReverseSorter[T]) Len() int { return r.Inner.Len() }

// Swap implements Interface
func (r ReverseSorter[T]) Swap(i int, j int) { r.Inner.Swap(i, j) }

// Less returns the opposite of the embedded implementation's Less method.
func (r ReverseSorter[T]) Less(i, j int) bool { return !r.Inner.Less(j, i) }

type SliceSorter[Elem any] struct {
	Data     []Elem
	LessFunc func(data []Elem, i, j int) bool
}

// Len implements Interface
func (s SliceSorter[E]) Len() int { return len(s.Data) }

// Less implements Interface
func (s SliceSorter[E]) Less(i, j int) bool { return s.LessFunc(s.Data, i, j) }

// Swap implements Interface
func (s SliceSorter[E]) Swap(i, j int) { s.Data[i], s.Data[j] = s.Data[j], s.Data[i] }
