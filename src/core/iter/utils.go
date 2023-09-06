// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

import (
	"github.com/primecitizens/std/core/assert"
)

// Contains returns true if there is x in iter.
func Contains[T comparable, Iter Core[T]](iter Iter, x T) bool {
	for i := 0; ; i++ {
		item, ok := iter.Nth(i)
		if !ok {
			return false
		}

		if x == item {
			return true
		}
	}
}

// Index returns the index in range [0, n) from i.
//
// assuming n > 0:
//
//   - when i >= 0, validate it
//   - when i < 0, index from the n
//
// NOTE: the returned index is only valid when valid is true.
func Index(i, n int) (index int, valid bool) {
	if i >= 0 {
		return i, i < n
	}

	if n > 0 {
		i = n + i
		return i, i >= 0 && i < n
	}

	return -1, false
}

// Bound returns value used for slice bounding.
//
// assuming n >= 0:
//
//   - i >= 0 && i < n: Bound(i, n) = Index(i, n)
//   - i == n: returns (i, true)
//   - i < 0: Bound(i, n) = Index(i, n) + 1
//
// NOTE: the returned bound is only valid when valid is true.
func Bound(i, n int) (bound int, valid bool) {
	if i >= 0 {
		return i, i <= n
	}

	if n < 0 {
		return -1, false
	}

	i = n + i + 1
	return i, i >= 0 && i <= n
}

// Range returns the [start:end] in range [0, n]
func Range(start, end, n int) (int, int, bool) {
	if start >= 0 && end >= 0 {
		return start, end, start <= end && end <= n
	}

	if start < 0 {
		start = n + start + 1
		if start < 0 {
			return -1, -1, false
		}
	}

	if end < 0 {
		end = n + end + 1
		if end < 0 {
			return -1, -1, false
		}
	}

	return start, end, start <= end && end <= n
}

// Each calls fn on each element from the iter.
func Each[Iter Core[Elem], Elem any](
	iter Iter, fn func(i int, v Elem) bool,
) int {
	return EachEx(iter, fn, func(i int, v Elem, fn func(i int, v Elem) bool) bool {
		return fn(i, v)
	})
}

// EachEx is like Each but calls fn with an extra arg.
func EachEx[Iter Core[Elem], Elem, Arg any](
	iter Iter, arg Arg, fn func(i int, v Elem, arg Arg) bool,
) int {
	var (
		val Elem
		ok  bool
	)

	for i := 0; ; i++ {
		if val, ok = iter.Nth(i); ok {
			if fn(i, val, arg) {
				continue
			}
		}

		return i
	}
}

// Step calls fn on every next step element from the iter.
func Step[Elem any](
	iter Core[Elem], from, step int, fn func(i int, v Elem) bool,
) int {
	var (
		val Elem
		ok  bool
	)

	if step == 0 {
		assert.Throw("invald", "zero", "step")
	}

	if step < 0 {
		if _, ok = iter.(CanBackward); ok {
			assert.Throw("no", "backward", "iteration", "support")
		}
	}

	if _, ok = iter.(CanSkip); ok && (step > 1 || step < -1) {
		if step < 0 {
			for i := from; ; i-- {
				if val, ok = iter.Nth(i); ok {
					if i == from {
						if fn(i, val) {
							from += step
							continue
						}
					} else {
						continue
					}
				}

				return i
			}
		} else {
			for i := from; ; i++ {
				if val, ok = iter.Nth(i); ok {
					if i == from {
						if fn(i, val) {
							from += step
							continue
						}
					} else {
						continue
					}
				}

				return from
			}
		}

		assert.Unreachable()
	}

	for ; ; from += step {
		if val, ok = iter.Nth(from); ok {
			if fn(from, val) {
				continue
			}
		}

		return from
	}
}

// StepEx is like Step but calls fn with an extra arg.
//
// It panics if from or step is invalid for the iter.
func StepEx[Elem, Arg any](
	iter Core[Elem], from, step int, arg Arg, fn func(i int, v Elem, arg Arg) bool,
) {
	var (
		val Elem
		ok  bool
	)

	if step == 0 {
		assert.Throw("invald", "zero", "step")
	}

	if step < 0 {
		if _, ok = iter.(CanBackward); ok {
			assert.Throw("no", "backward", "iteration", "support")
		}
	}

	if _, ok = iter.(CanSkip); ok {
		if step < 0 {
			for i := from; ; i-- {
				if val, ok = iter.Nth(i); ok {
					if i == from {
						if fn(i, val, arg) {
							from += step
							continue
						}
					} else {
						continue
					}
				}

				return
			}

			assert.Unreachable()
		}

		for i := from; ; i++ {
			if val, ok = iter.Nth(i); ok {
				if i == from {
					if fn(i, val, arg) {
						from += step
						continue
					}
				} else {
					continue
				}
			}

			return
		}

		assert.Unreachable()
	}

	for ; ; from += step {
		if val, ok = iter.Nth(from); ok {
			if fn(from, val, arg) {
				continue
			}
		}

		return
	}
}
