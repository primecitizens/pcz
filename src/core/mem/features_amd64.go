// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mem

import "github.com/primecitizens/std/core/cpu"

var (
	hasERMS       bool // required by memclr & memmove
	useAVXmemmove bool // required by memmove
)

func init() {
	hasERMS = cpu.X86.HasAll(cpu.X86Feature_erms)

	// TODO: need to check cpu family: see golang std src/runtime/cpuflags_amd64.go
	useAVXmemmove = cpu.X86.HasAll(cpu.X86Feature_avx)
}
