// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdslice

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
//
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// When Compact discards m elements in total, it might not modify the elements
// s[len(s)-m:len(s)]. If those elements contain pointers you might consider
// zeroing those elements so that objects they reference can be garbage collected.
func Compact[E any](s []E, eq func(s []E, prev, cur int) bool) []E {
	if len(s) < 2 {
		return s
	}
	i := 1
	for k := 1; k < len(s); k++ {
		if eq(s, k-1, k) {
			continue
		}

		if i != k {
			s[i] = s[k]
		}
		i++
	}
	return s[:i]
}

// CompactEx is like [Compact], but takes an extra arg.
func CompactEx[E, T any](s []E, arg T, eq func(arg T, s []E, prev, cur int) bool) []E {
	if len(s) < 2 {
		return s
	}
	i := 1
	for k := 1; k < len(s); k++ {
		if eq(arg, s, k-1, k) {
			continue
		}

		if i != k {
			s[i] = s[k]
		}
		i++
	}
	return s[:i]
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if s[i:j] is not a valid slice of s.
// Delete is O(len(s)-j), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete might not modify the elements s[len(s)-(j-i):len(s)]. If those
// elements contain pointers you might consider zeroing those elements so that
// objects they reference can be garbage collected.
func Delete[E any](s []E, i, j int) []E {
	_ = s[i:j] // bounds check
	return append(s[:i], s[j:]...)
}

func Cut[E any](s []E) (used, remaining []E) {
	return s[:len(s):len(s)], s[len(s):len(s):cap(s)]
}
