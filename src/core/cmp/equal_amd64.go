// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cmp

import (
	"github.com/primecitizens/std/core/cpu"
)

var hasAVX2 bool

func init() {
	hasAVX2 = cpu.X86.HasAll(cpu.X86Feature_avx2)
}
