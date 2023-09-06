// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdslice

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
)

func Make(typ *abi.Type, len int, cap int) unsafe.Pointer {
	assert.TODO()
	return nil
}

func MakeCopy(typ *abi.Type, tolen int, fromlen int, from unsafe.Pointer) unsafe.Pointer {
	assert.TODO()
	return nil
}

// Grow allocates new backing store for a slice.
//
// arguments:
//
//	oldPtr = pointer to the slice's backing array
//	newLen = new length (= oldLen + num)
//	oldCap = original slice's capacity.
//	   num = number of elements being added
//	    et = element type
//
// return values:
//
//	newPtr = pointer to the new backing store
//	newLen = same value as the argument
//	newCap = capacity of the new backing store
//
// Requires that uint(newLen) > uint(oldCap).
// Assumes the original slice length is newLen - num
//
// A new backing store is allocated with space for at least newLen elements.
// Existing entries [0, oldLen) are copied over to the new backing store.
// Added entries [oldLen, newLen) are not initialized by growslice
// (although for pointer-containing element types, they are zeroed). They
// must be initialized by the caller.
// Trailing entries [newLen, newCap) are zeroed.
//
// growslice's odd calling convention makes the generated code that calls
// this function simpler. In particular, it accepts and returns the
// new length so that the old length is not live (does not need to be
// spilled/restored) and the new length is returned (also does not need
// to be spilled/restored).
func Grow(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *abi.Type) Header {
	assert.TODO()
	return Header{}
}

func Copy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, wid uintptr) int {
	assert.TODO()
	return 0
}
