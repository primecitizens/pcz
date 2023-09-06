// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/thread"
)

//go:nosplit
func Fastrand() uint32 {
	return thread.GetG().G().Fastrand32()
}

//go:nosplit
func FastrandN(n uint32) uint32 {
	// This is similar to fastrand() % n, but faster.
	// See https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32(uint64(Fastrand()) * uint64(n) >> 32)
}

//go:nosplit
func Fastrand64() uint64 {
	return thread.GetG().G().Fastrand64()
}

//go:nosplit
func FastrandU() uint {
	if arch.UintBits == 32 {
		return uint(Fastrand())
	}
	return uint(Fastrand64())
}
