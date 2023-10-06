// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdptr

// TaggedPointer is a pointer with a numeric tag.
// The size of the numeric tag is GOARCH-dependent,
// currently at least 10 bits.
// This should only be used with pointers allocated outside the Go heap.
type TaggedPointer uint64

// MinTagBits is the minimum number of tag bits that we expect.
const MinTagBits = 10
