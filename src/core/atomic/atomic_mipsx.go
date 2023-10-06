// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips || mipsle

package atomic

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/cpu"
)

// TODO implement lock striping
var lock struct {
	state uint32
	pad   [cpu.CacheLinePadSize - 4]byte
}

//go:noescape
func spinLock(state *uint32)

//go:noescape
func spinUnlock(state *uint32)

//go:nosplit
func lockAndCheck(addr *uint64) {
	// ensure 8-byte alignment
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		panicUnaligned()
	}
	// force dereference before taking lock
	_ = *addr

	spinLock(&lock.state)
}

//go:nosplit
func unlock() {
	spinUnlock(&lock.state)
}

//go:nosplit
func unlockNoFence() {
	lock.state = 0
}

//
// Store
//

//go:noescape
func Store8(ptr *uint8, val uint8)

//go:noescape
func Store32(ptr *uint32, val uint32)

//go:nosplit
func Store64(addr *uint64, val uint64) {
	lockAndCheck(addr)

	*addr = val

	unlock()
	return
}

// NO go:noescape annotation; see atomic_pointer.go.
func StorePointer(ptr unsafe.Pointer, val unsafe.Pointer)

//
// StoreRel
//

//go:noescape
func StoreRel32(ptr *uint32, val uint32)

//go:noescape
func StoreRelUintptr(ptr *uintptr, val uintptr)

//
// Load
//

//go:noescape
func Load8(ptr *uint8) uint8

//go:noescape
func Load32(ptr *uint32) uint32

//go:nosplit
func Load64(addr *uint64) (val uint64) {
	lockAndCheck(addr)

	val = *addr

	unlock()
	return
}

// NO go:noescape annotation; *ptr escapes if result escapes (#31525)
func LoadPointer(ptr unsafe.Pointer) unsafe.Pointer

// LoadAcq

//go:noescape
func LoadAcq32(ptr *uint32) uint32

//go:noescape
func LoadAcqUintptr(ptr *uintptr) uintptr

//
// Swap
//

//go:noescape
func Swap32(ptr *uint32, new uint32) uint32

//go:nosplit
func Swap64(addr *uint64, new uint64) (old uint64) {
	lockAndCheck(addr)

	old = *addr
	*addr = new

	unlock()
	return
}

//go:noescape
func SwapUintptr(ptr *uintptr, new uintptr) uintptr

//
// Add
//

//go:noescape
func Add32(ptr *uint32, delta int32) uint32

//go:nosplit
func Add64(addr *uint64, delta int64) (new uint64) {
	lockAndCheck(addr)

	new = *addr + uint64(delta)
	*addr = new

	unlock()
	return
}

//go:noescape
func AddUintptr(ptr *uintptr, delta uintptr) uintptr

//
// Compare and swap
//

//go:nosplit
func Cas64(addr *uint64, old, new uint64) (swapped bool) {
	lockAndCheck(addr)

	if (*addr) == old {
		*addr = new
		unlock()
		return true
	}

	unlockNoFence()
	return false
}

//
// bitwise
//

//go:noescape
func And8(ptr *uint8, val uint8)

//go:noescape
func And32(ptr *uint32, val uint32)

//go:noescape
func Or8(ptr *uint8, val uint8)

//go:noescape
func Or32(ptr *uint32, val uint32)

//
// CasRel
//

//go:noescape
func CasRel32(addr *uint32, old, new uint32) bool
