// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdstring

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/bytealg"
	"github.com/primecitizens/pcz/std/core/os"
)

// FindNull returns the index of first \x00 byte, for UTF-8 encoding
//
//go:nosplit
func FindNull(s *byte) int {
	if s == nil {
		return 0
	}

	// Avoid IndexByteString on Plan 9 because it uses SSE instructions
	// on x86 machines, and those are classified as floating point instructions,
	// which are illegal in a note handler.
	if os.IsPlan9 != 0 {
		p := (*[os.MaxAlloc/2 - 1]uint8)(unsafe.Pointer(s))
		l := 0
		for p[l] != 0 {
			l++
		}
		return l
	}

	// pageSize is the unit we scan at a time looking for NULL.
	// It must be the minimum page size for any architecture Go
	// runs on. It's okay (just a minor performance loss) if the
	// actual system page size is larger than this value.
	const pageSize = 4096

	offset := 0
	ptr := unsafe.Pointer(s)
	// IndexByteString uses wide reads, so we need to be careful
	// with page boundaries. Call IndexByteString on
	// [ptr, endOfPage) interval.
	safeLen := int(pageSize - uintptr(ptr)%pageSize)

	for {
		t := unsafe.String((*byte)(ptr), safeLen)
		// Check one page at a time.
		if i := bytealg.IndexByte(t, 0); i != -1 {
			return offset + i
		}
		// Move to next page
		ptr = unsafe.Pointer(uintptr(ptr) + uintptr(safeLen))
		offset += safeLen
		safeLen = pageSize
	}
}

// FindNull2 returns the first \x00 wide-char, intended for UTF-16 encoding
//
//go:nosplit
func FindNull2(s *uint16) int {
	if s == nil {
		return 0
	}
	p := (*[os.MaxAlloc/2/2 - 1]uint16)(unsafe.Pointer(s))
	l := 0
	for p[l] != 0 {
		l++
	}
	return l
}
