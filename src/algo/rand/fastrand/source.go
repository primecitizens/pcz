// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Package fastrand is adapted from the standard Go runtime fastrand
// implementation.
//
// wyrand (https://github.com/wangyi-fudan/wyhash) on 64bit platforms and
// xorshift64+ (https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf)
// on 32bit platforms.
package fastrand

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/math"
)

type Source struct {
	fastrand uint64
}

//go:nosplit
func (s *Source) Rand32() uint32 {
	// Only the platform that math.Mul64 can be lowered
	// by the compiler should be in this list.
	if arch.IsAmd64|arch.IsArm64|arch.IsPpc64|
		arch.IsPpc64le|arch.IsMips64|arch.IsMips64le|
		arch.IsS390x|arch.IsRiscv64 == 1 {
		s.fastrand += 0xa0761d6478bd642f
		hi, lo := math.Mul64(s.fastrand, s.fastrand^0xe7037ed1a0b428db)
		return uint32(hi ^ lo)
	}

	// Implement xorshift64+: 2 32-bit xorshift sequences added together.
	// Shift triplet [17,7,16] was calculated as indicated in Marsaglia's
	// Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
	// This generator passes the SmallCrush suite, part of TestU01 framework:
	// http://simul.iro.umontreal.ca/testu01/tu01.html
	t := (*[2]uint32)(unsafe.Pointer(&s.fastrand))
	s1, s0 := t[0], t[1]
	s1 ^= s1 << 17
	s1 = s1 ^ s0 ^ s1>>7 ^ s0>>16
	t[0], t[1] = s0, s1
	return s0 + s1
}

func (s *Source) Rand64() uint64 {
	// Implement wyrand: https://github.com/wangyi-fudan/wyhash
	// Only the platform that math.Mul64 can be lowered
	// by the compiler should be in this list.
	if arch.IsAmd64|arch.IsArm64|arch.IsPpc64|
		arch.IsPpc64le|arch.IsMips64|arch.IsMips64le|
		arch.IsS390x|arch.IsRiscv64 == 1 {
		s.fastrand += 0xa0761d6478bd642f
		hi, lo := math.Mul64(s.fastrand, s.fastrand^0xe7037ed1a0b428db)
		return hi ^ lo
	}

	// Implement xorshift64+: 2 32-bit xorshift sequences added together.
	// Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
	// This generator passes the SmallCrush suite, part of TestU01 framework:
	// http://simul.iro.umontreal.ca/testu01/tu01.html
	t := (*[2]uint32)(unsafe.Pointer(&s.fastrand))
	s1, s0 := t[0], t[1]
	s1 ^= s1 << 17
	s1 = s1 ^ s0 ^ s1>>7 ^ s0>>16
	r := uint64(s0 + s1)

	s0, s1 = s1, s0
	s1 ^= s1 << 17
	s1 = s1 ^ s0 ^ s1>>7 ^ s0>>16
	r += uint64(s0+s1) << 32

	t[0], t[1] = s0, s1
	return r
}
