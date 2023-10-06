// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && s390x

package bytealg

import (
	"github.com/primecitizens/pcz/std/core/cpu"
)

const (
	MaxBruteForce = 64
)

var (
	hasVX bool
	// indexArgBMaxLen is the maximum length of the string to be searched for (argument b) in Index.
	// If indexArgBMaxLen is not 0, make sure indexArgBMaxLen >= 4.
	indexArgBMaxLen int
)

func init() {
	// Note: we're kind of lucky that this flag is available at this point.
	// The runtime sets HasVX when processing auxv records, and that happens
	// to happen *before* running the init functions of packages that
	// the runtime depends on.
	// TODO: it would really be nicer for core/cpu to figure out this
	// flag by itself. Then we wouldn't need to depend on quirks of
	// early startup initialization order.

	hasVX = cpu.S390X.HasAll(S390xFeature_vx)
	if hasVX {
		indexArgBMaxLen = 64
	}
}

// cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
}
