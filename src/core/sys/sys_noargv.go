// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos

package sys

// for systems without argv

func (ArgIter) Nth(i int) (arg string, ok bool) { return }
func (ArgIter) Len() int                        { return 0 }
func (ArgIter) SliceFrom(start int) ArgIter     { return 0 }
