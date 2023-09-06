// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos || js || wasip1 || darwin || openbsd || windows

package sys

type Auxv uintptr

// for systems without auxv

func (it AuxvIter) Nth(i int) (aux Auxv, ok bool) { return }
func (it AuxvIter) Len() int                      { return 0 }
func (it AuxvIter) SliceFrom(start int) AuxvIter  { return 0 }
