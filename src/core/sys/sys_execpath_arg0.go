// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (js || wasip1 || linux || freebsd || netbsd || openbsd || solaris || dragonfly || windows)

package sys

func ExecutablePath() string {
	if len(args) > 0 {
		return ntharg(0)
	}

	return ""
}
