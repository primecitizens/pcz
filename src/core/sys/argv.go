// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos

package sys

import (
	"github.com/primecitizens/std/core/iter"
)

func (it ArgIter) Nth(i int) (arg string, ok bool) {
	if it < 0 {
		return
	}

	i, ok = iter.Index(i, len(args)-int(it))
	if ok {
		return ntharg(i), true
	}

	return
}

func (it ArgIter) Len() int { return len(args) - int(it) }

func (it ArgIter) SliceFrom(start int) ArgIter {
	start, ok := iter.Bound(start, len(args)-int(it))
	if ok {
		return it + ArgIter(start)
	}

	return ArgIter(len(args))
}
