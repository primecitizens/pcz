// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && 386

package mem

import "github.com/primecitizens/pcz/std/core/cpu"

var (
	hasERMS bool // required by memclr & memmove
)

func init() {
	hasERMS = cpu.X86.HasAll(cpu.X86Feature_erms)
}
