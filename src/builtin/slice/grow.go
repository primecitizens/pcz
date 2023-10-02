// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdslice

import (
	"unsafe"

	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/assert"
)

func Append[E any, Alloc alloc.M](alloc Alloc, dst []E, elems ...E) []E {
	assert.TODO()
	return dst
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[E any, Alloc alloc.M](alloc Alloc, s []E, n int) []E {
	if n < 0 {
		assert.Throw("cannot", "be", "negative")
	}
	if n -= cap(s) - len(s); n > 0 {
		s = Append(alloc, s[:cap(s)], make([]E, n)...)[:len(s)]
	}
	return s
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[E any, Alloc alloc.M](alloc Alloc, s []E, i int, v ...E) []E {
	m := len(v)
	if m == 0 {
		return s
	}
	n := len(s)
	if i == n {
		return Append(alloc, s, v...)
	}
	if n+m > cap(s) {
		// Use append rather than make so that we bump the size of
		// the slice up to the next storage class.
		// This is what Grow does but we don't call Grow because
		// that might copy the values twice.
		s2 := Append(alloc, s[:i], make([]E, n+m-i)...)
		copy(s2[i:], v)
		copy(s2[i+m:], s[i:])
		return s2
	}
	s = s[:n+m]

	// before:
	// s: aaaaaaaabbbbccccccccdddd
	//            ^   ^       ^   ^
	//            i  i+m      n  n+m
	// after:
	// s: aaaaaaaavvvvbbbbcccccccc
	//            ^   ^       ^   ^
	//            i  i+m      n  n+m
	//
	// a are the values that don't move in s.
	// v are the values copied in from v.
	// b and c are the values from s that are shifted up in index.
	// d are the values that get overwritten, never to be seen again.

	if !overlaps(v, s[i+m:]) {
		// Easy case - v does not overlap either the c or d regions.
		// (It might be in some of a or b, or elsewhere entirely.)
		// The data we copy up doesn't write to v at all, so just do it.

		copy(s[i+m:], s[i:])

		// Now we have
		// s: aaaaaaaabbbbbbbbcccccccc
		//            ^   ^       ^   ^
		//            i  i+m      n  n+m
		// Note the b values are duplicated.

		copy(s[i:], v)

		// Now we have
		// s: aaaaaaaavvvvbbbbcccccccc
		//            ^   ^       ^   ^
		//            i  i+m      n  n+m
		// That's the result we want.
		return s
	}

	// The hard case - v overlaps c or d. We can't just shift up
	// the data because we'd move or clobber the values we're trying
	// to insert.
	// So instead, write v on top of d, then rotate.
	copy(s[n:], v)

	// Now we have
	// s: aaaaaaaabbbbccccccccvvvv
	//            ^   ^       ^   ^
	//            i  i+m      n  n+m

	rotateRight(s[i:], m)

	// Now we have
	// s: aaaaaaaavvvvbbbbcccccccc
	//            ^   ^       ^   ^
	//            i  i+m      n  n+m
	// That's the result we want.
	return s
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice. Replace panics if s[i:j] is not a valid slice of s.
func Replace[E any, Alloc alloc.M](alloc Alloc, s []E, i, j int, v ...E) []E {
	_ = s[i:j] // verify that i:j is a valid subslice

	if i == j {
		return Insert(alloc, s, i, v...)
	}
	if j == len(s) {
		return Append(alloc, s[:i], v...)
	}

	tot := len(s[:i]) + len(v) + len(s[j:])
	if tot > cap(s) {
		// Too big to fit, allocate and copy over.
		s2 := Append(alloc, s[:i], make([]E, tot-i)...) // See Insert
		copy(s2[i:], v)
		copy(s2[i+len(v):], s[j:])
		return s2
	}

	r := s[:tot]

	if i+len(v) <= j {
		// Easy, as v fits in the deleted portion.
		copy(r[i:], v)
		if i+len(v) != j {
			copy(r[i+len(v):], s[j:])
		}
		return r
	}

	// We are expanding (v is bigger than j-i).
	// The situation is something like this:
	// (example has i=4,j=8,len(s)=16,len(v)=6)
	// s: aaaaxxxxbbbbbbbbyy
	//        ^   ^       ^ ^
	//        i   j  len(s) tot
	// a: prefix of s
	// x: deleted range
	// b: more of s
	// y: area to expand into

	if !overlaps(r[i+len(v):], v) {
		// Easy, as v is not clobbered by the first copy.
		copy(r[i+len(v):], s[j:])
		copy(r[i:], v)
		return r
	}

	// This is a situation where we don't have a single place to which
	// we can copy v. Parts of it need to go to two different places.
	// We want to copy the prefix of v into y and the suffix into x, then
	// rotate |y| spots to the right.
	//
	//        v[2:]      v[:2]
	//         |           |
	// s: aaaavvvvbbbbbbbbvv
	//        ^   ^       ^ ^
	//        i   j  len(s) tot
	//
	// If either of those two destinations don't alias v, then we're good.
	y := len(v) - (j - i) // length of y portion

	if !overlaps(r[i:j], v) {
		copy(r[i:j], v[y:])
		copy(r[len(s):], v[:y])
		rotateRight(r[i:], y)
		return r
	}
	if !overlaps(r[len(s):], v) {
		copy(r[len(s):], v[:y])
		copy(r[i:j], v[y:])
		rotateRight(r[i:], y)
		return r
	}

	// Now we know that v overlaps both x and y.
	// That means that the entirety of b is *inside* v.
	// So we don't need to preserve b at all; instead we
	// can copy v first, then copy the b part of v out of
	// v to the right destination.
	k := startIdx(v, s[j:])
	copy(r[i:], v)
	copy(r[i+len(v):], r[i+k:])
	return r
}

// Rotation algorithm explanation:
//
// rotate left by 2
// start with
//   0123456789
// split up like this
//   01 234567 89
// swap first 2 and last 2
//   89 234567 01
// join first parts
//   89234567 01
// recursively rotate first left part by 2
//   23456789 01
// join at the end
//   2345678901
//
// rotate left by 8
// start with
//   0123456789
// split up like this
//   01 234567 89
// swap first 2 and last 2
//   89 234567 01
// join last parts
//   89 23456701
// recursively rotate second part left by 6
//   89 01234567
// join at the end
//   8901234567

// TODO: There are other rotate algorithms.
// This algorithm has the desirable property that it moves each element exactly twice.
// The triple-reverse algorithm is simpler and more cache friendly, but takes more writes.
// The follow-cycles algorithm can be 1-write but it is not very cache friendly.

// rotateLeft rotates b left by n spaces.
// s_final[i] = s_orig[i+r], wrapping around.
func rotateLeft[E any](s []E, r int) {
	for r != 0 && r != len(s) {
		if r*2 <= len(s) {
			swap(s[:r], s[len(s)-r:])
			s = s[:len(s)-r]
		} else {
			swap(s[:len(s)-r], s[r:])
			s, r = s[len(s)-r:], r*2-len(s)
		}
	}
}

func rotateRight[E any](s []E, r int) {
	rotateLeft(s, len(s)-r)
}

// swap swaps the contents of x and y. x and y must be equal length and disjoint.
func swap[E any](x, y []E) {
	for i := 0; i < len(x); i++ {
		x[i], y[i] = y[i], x[i]
	}
}

// overlaps reports whether the memory ranges a[0:len(a)] and b[0:len(b)] overlap.
func overlaps[E any](a, b []E) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	elemSize := unsafe.Sizeof(a[0])
	if elemSize == 0 {
		return false
	}
	// TODO: use a runtime/unsafe facility once one becomes available. See issue 12445.
	// Also see crypto/internal/alias/alias.go:AnyOverlap
	return uintptr(unsafe.Pointer(&a[0])) <= uintptr(unsafe.Pointer(&b[len(b)-1]))+(elemSize-1) &&
		uintptr(unsafe.Pointer(&b[0])) <= uintptr(unsafe.Pointer(&a[len(a)-1]))+(elemSize-1)
}

// startIdx returns the index in haystack where the needle starts.
// prerequisite: the needle must be aliased entirely inside the haystack.
func startIdx[E any](haystack, needle []E) int {
	p := &needle[0]
	for i := range haystack {
		if p == &haystack[i] {
			return i
		}
	}
	// TODO: what if the overlap is by a non-integral number of Es?
	assert.Throw("needle", "not", "found")
	return -1
}
