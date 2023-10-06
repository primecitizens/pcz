// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm

package atomic

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/cpu"
)

type spinlock struct {
	v uint32
}

//go:nosplit
func (l *spinlock) lock() {
	for {
		if Cas32(&l.v, 0, 1) {
			return
		}
	}
}

//go:nosplit
func (l *spinlock) unlock() {
	Store32(&l.v, 0)
}

var locktab [57]struct {
	l   spinlock
	pad [cpu.CacheLinePadSize - unsafe.Sizeof(spinlock{})]byte
}

func addrLock(addr *uint64) *spinlock {
	return &locktab[(uintptr(unsafe.Pointer(addr))>>3)%uintptr(len(locktab))].l
}

//go:nosplit
func armcas(ptr *uint32, old, new uint32) bool

//
// Store
//

//go:noescape
func Store8(addr *uint8, v uint8)

//go:noescape
func Store32(addr *uint32, v uint32)

//go:noescape
func Store64(addr *uint64, v uint64)

// Not noescape -- it installs a pointer to addr.
func StorePointer(addr unsafe.Pointer, v unsafe.Pointer)

//go:nosplit
func goStore64(addr *uint64, v uint64) {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		*(*int)(nil) = 0 // crash on unaligned uint64
	}
	_ = *addr // if nil, fault before taking the lock
	addrLock(addr).lock()
	*addr = v
	addrLock(addr).unlock()
}

//
// StoreRel
//

//go:noescape
func StoreRel32(addr *uint32, v uint32)

//go:noescape
func StoreRelUintptr(addr *uintptr, v uintptr)

//
// Load
//

//go:noescape
func Load8(addr *uint8) uint8

//go:noescape
func Load32(addr *uint32) uint32

//go:noescape
func Load64(addr *uint64) uint64

// NO go:noescape annotation; *addr escapes if result escapes (#31525)
func LoadPointer(addr unsafe.Pointer) unsafe.Pointer

//go:nosplit
func goLoad64(addr *uint64) uint64 {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		*(*int)(nil) = 0 // crash on unaligned uint64
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr
	addrLock(addr).unlock()
	return r
}

//
// LoadAcq
//

//go:noescape
func LoadAcq32(addr *uint32) uint32

//go:noescape
func LoadAcqUintptr(ptr *uintptr) uintptr

//
// bitwise
//

//go:nosplit
func Or8(addr *uint8, v uint8) {
	// Align down to 4 bytes and use 32-bit CAS.
	uaddr := uintptr(unsafe.Pointer(addr))
	addr32 := (*uint32)(unsafe.Pointer(uaddr &^ 3))
	word := uint32(v) << ((uaddr & 3) * 8) // little endian
	for {
		old := *addr32
		if Cas32(addr32, old, old|word) {
			return
		}
	}
}

//go:nosplit
func Or32(addr *uint32, v uint32) {
	for {
		old := *addr
		if Cas32(addr, old, old|v) {
			return
		}
	}
}

//go:nosplit
func And8(addr *uint8, v uint8) {
	// Align down to 4 bytes and use 32-bit CAS.
	uaddr := uintptr(unsafe.Pointer(addr))
	addr32 := (*uint32)(unsafe.Pointer(uaddr &^ 3))
	word := uint32(v) << ((uaddr & 3) * 8)    // little endian
	mask := uint32(0xFF) << ((uaddr & 3) * 8) // little endian
	word |= ^mask
	for {
		old := *addr32
		if Cas32(addr32, old, old&word) {
			return
		}
	}
}

//go:nosplit
func And32(addr *uint32, v uint32) {
	for {
		old := *addr
		if Cas32(addr, old, old&v) {
			return
		}
	}
}

//
// Swap
//

//go:nosplit
func Swap32(addr *uint32, v uint32) uint32 {
	for {
		old := *addr
		if Cas32(addr, old, v) {
			return old
		}
	}
}

//go:noescape
func Swap64(addr *uint64, v uint64) uint64

//go:nosplit
func SwapUintptr(addr *uintptr, v uintptr) uintptr {
	return uintptr(Swap32((*uint32)(unsafe.Pointer(addr)), uint32(v)))
}

//go:nosplit
func goSwap64(addr *uint64, v uint64) uint64 {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		*(*int)(nil) = 0 // crash on unaligned uint64
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr
	*addr = v
	addrLock(addr).unlock()
	return r
}

//
// Add
//

// Atomic add and return new value.
//
//go:nosplit
func Add32(val *uint32, delta int32) uint32 {
	for {
		oval := *val
		nval := oval + uint32(delta)
		if Cas32(val, oval, nval) {
			return nval
		}
	}
}

//go:noescape
func Add64(addr *uint64, delta int64) uint64

//go:noescape
func AddUintptr(ptr *uintptr, delta uintptr) uintptr

//go:nosplit
func goAdd64(addr *uint64, delta int64) uint64 {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		*(*int)(nil) = 0 // crash on unaligned uint64
	}
	_ = *addr // if nil, fault before taking the lock
	var r uint64
	addrLock(addr).lock()
	r = *addr + uint64(delta)
	*addr = r
	addrLock(addr).unlock()
	return r
}

//
// Compare and swap
//

//go:noescape
func Cas64(addr *uint64, old, new uint64) bool

//go:nosplit
func goCas64(addr *uint64, old, new uint64) bool {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		*(*int)(nil) = 0 // crash on unaligned uint64
	}
	_ = *addr // if nil, fault before taking the lock
	var ok bool
	addrLock(addr).lock()
	if *addr == old {
		*addr = new
		ok = true
	}
	addrLock(addr).unlock()
	return ok
}

//
// CasRel
//

//go:noescape
func CasRel32(addr *uint32, old, new uint32) bool
