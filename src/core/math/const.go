// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/primecitizens/std/core/arch"
)

const (
	MaxUintptr = ^uintptr(0)
	MaxUint    = ^uint(0)
	MaxUint8   = ^uint8(0)
	MaxUint16  = ^uint16(0)
	MaxUint32  = ^uint32(0)
	MaxUint64  = ^uint64(0)
)

const (
	MaxInt   = 1<<(arch.PtrSize-1) - 1
	MaxInt8  = 1<<7 - 1
	MaxInt16 = 1<<15 - 1
	MaxInt32 = 1<<31 - 1
	MaxInt64 = 1<<63 - 1

	MinInt   = -1 << (arch.PtrSize - 1)
	MinInt8  = -1 << 7
	MinInt16 = -1 << 15
	MinInt32 = -1 << 31
	MinInt64 = -1 << 63
)
