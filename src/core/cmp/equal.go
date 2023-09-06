// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"unsafe"
)

//go:noescape
func MemEqual(p, q unsafe.Pointer, sz uintptr) bool

// BytesEqual reports whether a and b are the same length and contain the
// same bytes.
// A nil argument is equivalent to an empty slice.
func BytesEqual(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	// There is a test for this in package bytes.
	return string(a) == string(b)
}
