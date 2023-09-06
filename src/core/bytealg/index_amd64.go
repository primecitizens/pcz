// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytealg

import (
	"github.com/primecitizens/std/core/cpu"
)

const MaxBruteForce = 64

var (
	hasSSE42  bool
	hasAVX2   bool
	hasPOPCNT bool
)

func init() {
	if cpu.X86.HasAll(cpu.X86Feature_avx2) {
		MaxLen = 63
		hasAVX2 = true
	} else {
		MaxLen = 31
	}

	hasSSE42 = cpu.X86.HasAll(cpu.X86Feature_sse42)
	hasPOPCNT = cpu.X86.HasAll(cpu.X86Feature_popcnt)
}

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func Cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
}
