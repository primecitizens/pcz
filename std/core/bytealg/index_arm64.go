// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && arm64

package bytealg

// Optimize cases where the length of the substring is less than 32 bytes
const (
	// Empirical data shows that using Index can get better
	// performance when len(s) <= 16.
	MaxBruteForce = 16

	// indexArgBMaxLen is the maximum length of the string to be searched for (argument b) in Index.
	// If indexArgBMaxLen is not 0, make sure indexArgBMaxLen >= 4.
	indexArgBMaxLen = 32
)

// cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func cutover(n int) int {
	// 1 error per 16 characters, plus a few slop to start.
	return 4 + n>>4
}
