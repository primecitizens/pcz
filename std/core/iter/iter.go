// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Package iter.
package iter

// Core is the minimal interface of an iterator.
//
// Rationales choosing interface over function for iterator definition:
//
//   - Implicit allocations: interface values can use mark.Noescape to avoid
//     allocation but functions cannot, closures always allocate unless
//     properly inlined by the compiler, which is not something controlled
//     by application developers (at the time of go1.21).
//
// Drawbacks of interface:
//
//   - One concrete type can only implement one iterator interface, some types
//     may need multiple iterator implementations.
//     A workaround is to define `type OtherTree Tree` and use pointer type
//     cast for iterator implementation.
type Core[Elem any] interface {
	// Nth returns the n-th element, the returned bool value indicates whether
	// the element exists.
	//
	// It is required to support the 1-step [0, end) range as input sequence,
	// in which case Nth(int) MUST function exactly the same way as Next() would.
	//
	// Caller SHOULD assume the implementation only supports the iteration
	// style implied by the above requirement:
	//
	//	for i := 0; ; i++ {
	//		elem, ok := iter.Nth(i)
	//		if !ok {
	//			break
	//		}
	//
	//		_ = elem
	// 	}
	//
	// Following features are optional:
	//
	//   - Negative n (index from end)
	//   - Step more than 1 (calling with non-consecutive n)
	//   - Backward iteration (calling with n from high to low)
	//   - Arbitrary iteration (calling with unordered n)
	//
	// Rationales choosing Nth(int) instead of Next() as core iterator function:
	//
	//   - Passing an integer argument is cheap (a register in ABIInternal).
	//   - Nth(int) is Next() when called with a sequence of ascending numbers.
	//     (per implementation requirement)
	//   - Nth(int) may skip unwanted elements without extraneous call.
	//   - Nth(int) may step-backward with no external caching.
	//   - Nth(int) is more friendly to slices.
	//   - Nth(int) may be used with binary search to get the total length
	//     of the iterator (when unknown).
	//   - Stateless values can use the integer passed in to decide when to
	//     stop iteration, otherwise an additional Next() style iterator
	//     should be implemented.
	Nth(int) (Elem, bool)
}

// Func implements Core for function.
type Func[Elem any] func(int) (Elem, bool)

func (fn Func[Elem]) Nth(i int) (Elem, bool) {
	return fn(i)
}

// Finite is implemented by iterators with a finite number of elements.
type Finite interface {
	// Len returns the count of elements a iterator may produce.
	Len() int
}

// Divisible is implemented by iterators can be divided into multiple
// sub-iterators.
type Divisible[Self any] interface {
	// SliceFrom returns a sub-iterator by slicing at the `start` of current
	// itertor, it resembles the native `x[start:]` operation.
	SliceFrom(start int) Self
}

// Interface is the most complete iterator interface covers most use cases
// of an iterator.
type Interface[T, Self any] interface {
	Core[T]
	Divisible[Self]
	Finite
}

// CanSkip should be implemented by iterators support skipping elements
// during a single iteration.
//
// Hint: embed iter.MarkCanSkip.
type CanSkip interface {
	IterCanSkip()
}

// CanBackward should be implemented by iterators support backward iteration
// (from higher to lower).
//
// Hint: embed iter.MarkCanBackward.
type CanBackward interface {
	IterCanBackward()
}

// CanIndexFromEnd should be implemented by iterators support negative
// argument to Nth(int) method.
//
// Hint: embed iter.MarkCanIndexFromEnd.
type CanIndexFromEnd interface {
	IterCanIndexFromEnd()
}

// MarkCanSkip is a zero size implementation of CanSkip.
type MarkCanSkip struct{}

func (MarkCanSkip) IterCanSkip() {}

// MarkCanBackward is a zero size implementation of CanBackward.
type MarkCanBackward struct{}

func (MarkCanBackward) IterCanBackward() {}

// MarkCanIndexFromEnd is a zero size implementation of CanIndexFromEnd.
type MarkCanIndexFromEnd struct{}

func (MarkCanIndexFromEnd) IterCanIndexFromEnd() {}
