// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && (linux || freebsd || netbsd || solaris || dragonfly)

package sys

import (
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/iter"
)

func (it AuxvIter) Nth(i int) (aux Auxv, ok bool) {
	if it < 0 {
		return
	}

	i, ok = iter.Index(i, len(auxvs)-int(it))
	if ok {
		return nthauxv(i), true
	}

	return
}

func (it AuxvIter) Len() int { return len(auxvs) - int(it) }

func (it AuxvIter) SliceFrom(start int) AuxvIter {
	start, ok := iter.Bound(start, len(auxvs)-int(it))
	if ok {
		return it + AuxvIter(start)
	}

	return AuxvIter(len(auxvs))
}
