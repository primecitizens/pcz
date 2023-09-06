// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build 386

package emu64

// Floating point control word values.
// Bits 0-5 are bits to disable floating-point exceptions.
// Bits 8-9 are the precision control:
//
//	0 = single precision a.k.a. float32
//	2 = double precision a.k.a. float64
//
// Bits 10-11 are the rounding mode:
//
//	0 = round to nearest (even on a tie)
//	3 = round toward zero
const (
	controlWord64      uint16 = 0x3f + 2<<8 + 0<<10
	controlWord64trunc uint16 = 0x3f + 2<<8 + 3<<10
)

func Float64ToUint32(float64) uint32
func Uint32ToFloat64(uint32) float64

//go:noescape
func _mul64by32(lo64 *uint64, a uint64, b uint32) (hi32 uint32)

//go:noescape
func _div64by32(a uint64, b uint32, r *uint32) (q uint32)

//go:nosplit
func dodiv(n, d uint64) (q, r uint64) {
	if d > n {
		return 0, n
	}

	if uint32(d>>32) != 0 {
		t := uint32(n>>32) / uint32(d>>32)
		var lo64 uint64
		hi32 := _mul64by32(&lo64, d, t)
		if hi32 != 0 || lo64 > n {
			return slowdodiv(n, d)
		}
		return uint64(t), n - lo64
	}

	// d is 32 bit
	var qhi uint32
	if uint32(n>>32) >= uint32(d) {
		if uint32(d) == 0 {
			stdpanic.PanicDivide()
		}
		qhi = uint32(n>>32) / uint32(d)
		n -= uint64(uint32(d)*qhi) << 32
	} else {
		qhi = 0
	}

	var rlo uint32
	qlo := _div64by32(n, uint32(d), &rlo)
	return uint64(qhi)<<32 + uint64(qlo), uint64(rlo)
}
