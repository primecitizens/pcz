// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

// Slice returns an iterator for the given slice.
func Slice[Elem any](x []Elem) SliceIter[Elem] {
	return SliceIter[Elem](x)
}

type SliceIter[Elem any] []Elem

// Nth implements Core[T]
func (s SliceIter[E]) Nth(i int) (elem E, ok bool) {
	i, ok = Index(i, len(s))
	if ok {
		return s[i], true
	}

	return
}

// Len implements Finite.
func (s SliceIter[E]) Len() int { return len(s) }

// SliceFrom implements Divisible[SliceIter[E]]
func (s SliceIter[E]) SliceFrom(start int) SliceIter[E] {
	start, ok := Bound(start, len(s))
	if ok {
		return s[start:]
	}

	return nil
}

type SliceIterEx[Elem any, T []Elem] int

func (x SliceIterEx[Elem, T]) NthEx(s T, i int) (e Elem, ok bool) {
	if int(x) >= len(s) || x < 0 {
		return
	}

	i, ok = Index(i, len(s))
	if ok {
		return s[i], true
	}

	return
}

func (x SliceIterEx[Elem, T]) LenEx(s T) int {
	if int(x) >= len(s) || x < 0 {
		return 0
	}

	return len(s) - int(x)
}

func (x SliceIterEx[Elem, T]) SliceFromEx(s T, start int) SliceIterEx[Elem, T] {
	start, ok := Bound(start, x.LenEx(s))
	if ok {
		return x + SliceIterEx[Elem, T](start)
	}

	return SliceIterEx[Elem, T](len(s))
}

func SliceReverse[Elem any](x []Elem) SliceReverseIter[Elem] {
	return SliceReverseIter[Elem](x)
}

type SliceReverseIter[Elem any] []Elem

// Nth implements Core[T]
func (s SliceReverseIter[E]) Nth(i int) (elem E, ok bool) {
	if i < 0 {
		i = -i - 1
		ok = /* implicitly true: i >= 0 && */ i < len(s)
	} else {
		i = len(s) - i - 1
		ok = i >= 0 && i < len(s)
	}

	if ok {
		return s[i], true
	}

	return
}

// Len implements Finite.
func (s SliceReverseIter[E]) Len() int { return len(s) }

// SliceFrom implements Divisible[SliceReverseIter[E]]
func (s SliceReverseIter[E]) SliceFrom(start int) SliceReverseIter[E] {
	if start < 0 {
		start = len(s) + start + 1
	} else {
		start = len(s) - start
	}

	if start >= 0 && start <= len(s) {
		return s[:start]
	}

	return nil
}

type SliceReverseIterEx[Elem any, T []Elem] int

// Nth implements Core[T]
func (x SliceReverseIterEx[Elem, T]) NthEx(s T, i int) (e Elem, ok bool) {
	if i < 0 {
		i = -i - 1
		ok = /* implicitly true: i >= 0 && */ i < len(s)
	} else {
		i = len(s) - i - 1
		ok = i >= 0 && i < len(s)
	}

	if ok {
		return s[i], true
	}

	return
}

// Len implements Finite.
func (x SliceReverseIterEx[Elem, T]) LenEx(s T) int { return len(s) }

// // SliceFrom implements Divisible[SliceReverseIterEx[Elem, T]]
// func (x SliceReverseIterEx[Elem, T]) SliceFromEx(s T, start int) SliceReverseIterEx[Elem, T] {
// 	if start < 0 {
// 		start = len(s) + start + 1
// 	} else {
// 		start = len(s) - start
// 	}
//
// 	if start >= 0 && start <= len(s) {
// 		return s[:start]
// 	}
//
// 	return nil
// }
