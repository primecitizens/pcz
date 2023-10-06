// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || arm || mips || mipsle

package stdptr

import "unsafe"

// The number of bits stored in the numeric tag of a TaggedPointer
const TaggedPointerBits = 32

// On 32-bit systems, TaggedPointer has a 32-bit pointer and 32-bit count.

// TaggedPointerPack created a TaggedPointer from a pointer and a tag.
// Tag bits that don't fit in the result are discarded.
func TaggedPointerPack(ptr unsafe.Pointer, tag uintptr) TaggedPointer {
	return TaggedPointer(uintptr(ptr))<<32 | TaggedPointer(tag)
}

// Pointer returns the pointer from a TaggedPointer.
func (tp TaggedPointer) pointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(tp >> 32))
}

// Tag returns the tag from a TaggedPointer.
func (tp TaggedPointer) tag() uintptr {
	return uintptr(tp)
}
