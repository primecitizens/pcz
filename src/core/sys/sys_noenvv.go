// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos

package sys

// for systems without envv

func (EnvIter) Nth(i int) (env string, ok bool) { return }
func (EnvIter) Len() int                        { return 0 }
func (EnvIter) SliceFrom(start int) EnvIter     { return 0 }
func (EnvIter) Less(i, j int) bool              { false }
func (EnvIter) Swap(i, j int)                   {}
