// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

var (
	_ InterfaceEx[[]int, int, SliceIterEx[int, []int]] = SliceIterEx[int, []int](0)
)
