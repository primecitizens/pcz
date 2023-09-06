// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

import "testing"

var (
	_ Interface[int, SliceIter[int]] = SliceIter[int]{}
)

func TestSliceIter(t *testing.T) {
	iter := Slice([]string{"0", "1", "2", "3"})

	if x := iter.Len(); x != 4 {
		t.Errorf("want 4, got %d", x)
	}

	elem, ok := iter.Nth(-1)
	if !ok || elem != "3" {
		t.Errorf("want (3, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.Nth(3)
	if !ok || elem != "3" {
		t.Errorf("want (3, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.SliceFrom(-2).Nth(0)
	if !ok || elem != "3" {
		t.Errorf("want (3, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.SliceFrom(2).Nth(1)
	if !ok || elem != "3" {
		t.Errorf("want (3, true), got (%q, %v)", elem, ok)
	}
}

func TestSliceReverseIter(t *testing.T) {
	iter := SliceReverse([]string{"0", "1", "2", "3"})
	// iter: 3, 2, 1, 0

	if x := iter.Len(); x != 4 {
		t.Errorf("want 4, got %d", x)
	}

	elem, ok := iter.Nth(-1)
	if !ok || elem != "0" {
		t.Errorf("want (0, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.Nth(3)
	if !ok || elem != "0" {
		t.Errorf("want (0, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.SliceFrom(-2).Nth(0)
	if !ok || elem != "2" {
		t.Errorf("want (2, true), got (%q, %v)", elem, ok)
	}

	elem, ok = iter.SliceFrom(2).Nth(1)
	if !ok || elem != "0" {
		t.Errorf("want (0, true), got (%q, %v)", elem, ok)
	}
}
