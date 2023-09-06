// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (js || wasip1 || darwin || linux || freebsd || netbsd || openbsd || solaris || dragonfly || windows)

package sys

import (
	"github.com/primecitizens/std/core/iter"
)

func (it EnvIter) Nth(i int) (env string, ok bool) {
	if it < 0 {
		return
	}

	i, ok = iter.Index(i, len(envs)-int(it))
	if ok {
		return nthenv(i), true
	}

	return
}

func (i EnvIter) Len() int { return len(envs) - int(i) }

func (i EnvIter) SliceFrom(start int) EnvIter {
	start, ok := iter.Bound(start, len(envs)-int(i))
	if ok {
		return i + EnvIter(start)
	}

	return EnvIter(len(envs))
}

func (EnvIter) Less(i, j int) bool { return nthenv(i) < nthenv(j) }
func (EnvIter) Swap(i, j int)      { envs[i], envs[j] = envs[j], envs[i] }
