// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sort

var (
	_ Interface = ReverseSorter[SliceSorter[byte]]{}
	_ Interface = SliceSorter[struct{}]{}
)
