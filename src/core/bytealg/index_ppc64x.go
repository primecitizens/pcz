// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && (ppc64 || ppc64le)

package bytealg

import (
	"github.com/primecitizens/std/core/cpu"
)

const (
	MaxBruteForce = 16

	// indexArgBMaxLen is the maximum length of the string to be searched for (argument b) in Index.
	// If indexArgBMaxLen is not 0, make sure indexArgBMaxLen >= 4.
	indexArgBMaxLen = 32
)

var (
	isPOWER9 bool
)

func init() {
	isPOWER9 = cpu.PPC64.HasAll(PPC64Feature_is_power9)
}

// cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
}
