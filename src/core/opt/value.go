// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package opt

import "github.com/primecitizens/std/core/assert"

func Maybe[T any](value T, hasValue bool) Value[T] {
	if hasValue {
		return Some(value)
	}

	return Value[T]{}
}

func Some[T any](v T) Value[T] {
	return Value[T]{
		V: v,

		some: true,
	}
}

// Value.
//
// Zero value means None.
type Value[T any] struct {
	V T

	some bool
}

func (o *Value[T]) Set(value T) {
	o.V = value
	o.some = true
}

func (o Value[T]) Get() T {
	if o.some {
		return o.V
	}

	assert.Throw("no", "value")
	return o.V
}

func (o *Value[T]) Erase() (hadValue bool) {
	if o == nil {
		return false
	}

	hadValue = o.some

	var zero T
	o.some = false
	o.V = zero
	return
}

func (o *Value[T]) IsNone() bool { return o == nil || !o.some }
func (o *Value[T]) IsSome() bool { return o != nil && o.some }
