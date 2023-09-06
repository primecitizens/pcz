// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64

package atomic

import (
	"unsafe"

	"github.com/primecitizens/std/core/cpu"
)

var (
	useARM64Atomics bool
)

func init() {
	useARM64Atomics = cpu.ARM64.HasAll(cpu.ARM64Feature_atomics)
}

//
// Store
//

//go:noescape
func Store8(ptr *uint8, val uint8)

//go:noescape
func Store32(ptr *uint32, val uint32)

//go:noescape
func Store64(ptr *uint64, val uint64)

// NO go:noescape annotation; see atomic_pointer.go.
func StorePointer(ptr unsafe.Pointer, val unsafe.Pointer)

//
// StoreRel
//

//go:noescape
func StoreRel32(ptr *uint32, val uint32)

//go:noescape
func StoreRel64(ptr *uint64, val uint64)

//go:noescape
func StoreRelUintptr(ptr *uintptr, val uintptr)

//
// Load
//

//go:noescape
func Load8(ptr *uint8) uint8

//go:noescape
func Load32(ptr *uint32) uint32

//go:noescape
func Load64(ptr *uint64) uint64

// NO go:noescape annotation; *ptr escapes if result escapes (#31525)
func LoadPointer(ptr unsafe.Pointer) unsafe.Pointer

//
// LoadAcq
//

//go:noescape
func LoadAcq32(addr *uint32) uint32

//go:noescape
func LoadAcq64(ptr *uint64) uint64

//go:noescape
func LoadAcqUintptr(ptr *uintptr) uintptr

//
// Swap
//

//go:noescape
func Swap32(ptr *uint32, new uint32) uint32

//go:noescape
func Swap64(ptr *uint64, new uint64) uint64

//go:noescape
func SwapUintptr(ptr *uintptr, new uintptr) uintptr

//
// Add
//

//go:noescape
func Add32(ptr *uint32, delta int32) uint32

//go:noescape
func Add64(ptr *uint64, delta int64) uint64

//go:noescape
func AddUintptr(ptr *uintptr, delta uintptr) uintptr

//
// bitwise
//

//go:noescape
func Or8(ptr *uint8, val uint8)

//go:noescape
func And8(ptr *uint8, val uint8)

//go:noescape
func Or32(ptr *uint32, val uint32)

//go:noescape
func And32(ptr *uint32, val uint32)

//
// Compare and swap
//

//go:noescape
func Cas64(ptr *uint64, old, new uint64) bool

//
// CasRel
//

//go:noescape
func CasRel32(ptr *uint32, old, new uint32) bool
