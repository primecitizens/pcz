// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build arm || mips || mipsle

package emu64

func Float64ToUint32(n float64) uint32 { return uint32(n) }
func Uint32ToFloat64(n uint32) float64 { return float64(n) }

//go:nosplit
func dodiv(n, d uint64) (q, r uint64) {
	// arm doesn't have a division instruction, so
	// slowdodiv is the best that we can do.

	// No _div64by32 on mips and using only _mul64by32 doesn't bring much benefit
	return slowdodiv(n, d)
}
