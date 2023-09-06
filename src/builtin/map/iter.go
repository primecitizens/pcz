// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdmap

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
)

// A hash iteration structure.
// If you modify HashIter, also change cmd/compile/internal/reflectdata/reflect.go
// and reflect/value.go to match the layout of this structure.
type HashIter struct {
	Key         unsafe.Pointer // Must be in first position.  Write nil to indicate iteration end (see cmd/compile/internal/walk/range.go).
	Elem        unsafe.Pointer // Must be in second position (see cmd/compile/internal/walk/range.go).
	T           *abi.MapType
	H           *hmap
	Buckets     unsafe.Pointer // bucket ptr at hash_iter initialization time
	Bptr        *bmap          // current bucket
	Overflow    *[]*bmap       // keeps overflow buckets of hmap.buckets alive
	OldOverflow *[]*bmap       // keeps overflow buckets of hmap.oldbuckets alive
	StartBucket uintptr        // bucket iteration started at
	Offset      uint8          // intra-bucket offset to start from during iteration (should be big enough to hold bucketCnt-1)
	Wrapped     bool           // already wrapped around from end of bucket array to beginning
	B           uint8
	I           uint8
	Bucket      uintptr
	CheckBucket uintptr
}

func MapIterInit(t *abi.MapType, h *hmap, it *hiter) {}

func MapIterNext(it *hiter) {}
