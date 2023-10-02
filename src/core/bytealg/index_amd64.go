// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && amd64

package bytealg

import (
	"github.com/primecitizens/std/core/cpu"
)

var (
	hasSSE42  bool
	hasAVX2   bool
	hasPOPCNT bool

	// indexArgBMaxLen is the maximum length of the string to be searched for (argument b) in Index.
	// If indexArgBMaxLen is not 0, make sure indexArgBMaxLen >= 4.
	indexArgBMaxLen int
)

func init() {
	if cpu.X86.HasAll(cpu.X86Feature_avx2) {
		indexArgBMaxLen = 63
		hasAVX2 = true
	} else {
		indexArgBMaxLen = 31
	}

	hasSSE42 = cpu.X86.HasAll(cpu.X86Feature_sse42)
	hasPOPCNT = cpu.X86.HasAll(cpu.X86Feature_popcnt)
}

const MaxBruteForce = 64

// cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
}

// A backup implementation to use by assembly.
func countGeneric(b []byte, c byte) int {
	n := 0
	for _, x := range b {
		if x == c {
			n++
		}
	}
	return n
}

func countGenericString(s string, c byte) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			n++
		}
	}
	return n
}
