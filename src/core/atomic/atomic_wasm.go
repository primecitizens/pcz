// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(neelance): implement with actual atomic operations as soon as threads are available
// See https://github.com/WebAssembly/design/issues/1073

package atomic

import "unsafe"

//go:nosplit
//go:noinline
func PublicationBarrier() {}

//
// Store
//

//go:nosplit
//go:noinline
func Store8(ptr *uint8, val uint8) {
	*ptr = val
}

//go:nosplit
//go:noinline
func Store32(ptr *uint32, val uint32) {
	*ptr = val
}

//go:nosplit
//go:noinline
func Store64(ptr *uint64, val uint64) {
	*ptr = val
}

//go:nosplit
//go:noinline
func StoreUintptr(ptr *uintptr, new uintptr) {
	*ptr = new
}

// StorePointer performs *ptr = val atomically and without a write
// barrier.
//
// NO go:noescape annotation; see atomic_pointer.go.
func StorePointer(ptr unsafe.Pointer, val unsafe.Pointer)

//go:nosplit
//go:noinline
func StoreInt32(ptr *int32, new int32) {
	*ptr = new
}

//go:nosplit
//go:noinline
func StoreInt64(ptr *int64, new int64) {
	*ptr = new
}

//
// StoreRel
//

//go:nosplit
//go:noinline
func StoreRel32(ptr *uint32, val uint32) {
	*ptr = val
}

//go:nosplit
//go:noinline
func StoreRel64(ptr *uint64, val uint64) {
	*ptr = val
}

//go:nosplit
//go:noinline
func StoreRelUintptr(ptr *uintptr, val uintptr) {
	*ptr = val
}

//
// Load
//

//go:nosplit
//go:noinline
func Load8(ptr *uint8) uint8 {
	return *ptr
}

//go:nosplit
//go:noinline
func Load32(ptr *uint32) uint32 {
	return *ptr
}

//go:nosplit
//go:noinline
func Load64(ptr *uint64) uint64 {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadUint(ptr *uint) uint {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadUintptr(ptr *uintptr) uintptr {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(ptr)
}

//go:nosplit
//go:noinline
func LoadInt32(ptr *int32) int32 {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadInt64(ptr *int64) int64 {
	return *ptr
}

//
// LoadAcq
//

//go:nosplit
//go:noinline
func LoadAcq32(ptr *uint32) uint32 {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadAcq64(ptr *uint64) uint64 {
	return *ptr
}

//go:nosplit
//go:noinline
func LoadAcqUintptr(ptr *uintptr) uintptr {
	return *ptr
}

//
// bitwise
//
// NOTE: Do not add atomicxor8 (XOR is not idempotent).

//go:nosplit
//go:noinline
func And8(ptr *uint8, val uint8) {
	*ptr = *ptr & val
}

//go:nosplit
//go:noinline
func And32(ptr *uint32, val uint32) {
	*ptr = *ptr & val
}

//go:nosplit
//go:noinline
func Or8(ptr *uint8, val uint8) {
	*ptr = *ptr | val
}

//go:nosplit
//go:noinline
func Or32(ptr *uint32, val uint32) {
	*ptr = *ptr | val
}

//
// Swap
//

//go:nosplit
//go:noinline
func Swap32(ptr *uint32, new uint32) uint32 {
	old := *ptr
	*ptr = new
	return old
}

//go:nosplit
//go:noinline
func Swap64(ptr *uint64, new uint64) uint64 {
	old := *ptr
	*ptr = new
	return old
}

//go:nosplit
//go:noinline
func SwapUintptr(ptr *uintptr, new uintptr) uintptr {
	old := *ptr
	*ptr = new
	return old
}

//go:nosplit
//go:noinline
func SwapInt32(ptr *int32, new int32) int32 {
	old := *ptr
	*ptr = new
	return old
}

//go:nosplit
//go:noinline
func SwapInt64(ptr *int64, new int64) int64 {
	old := *ptr
	*ptr = new
	return old
}

//
// Add
//

//go:nosplit
//go:noinline
func Add32(ptr *uint32, delta int32) uint32 {
	new := *ptr + uint32(delta)
	*ptr = new
	return new
}

//go:nosplit
//go:noinline
func Add64(ptr *uint64, delta int64) uint64 {
	new := *ptr + uint64(delta)
	*ptr = new
	return new
}

//go:nosplit
//go:noinline
func AddUintptr(ptr *uintptr, delta uintptr) uintptr {
	new := *ptr + delta
	*ptr = new
	return new
}

//go:nosplit
//go:noinline
func AddInt32(ptr *int32, delta int32) int32 {
	new := *ptr + delta
	*ptr = new
	return new
}

//go:nosplit
//go:noinline
func AddInt64(ptr *int64, delta int64) int64 {
	new := *ptr + delta
	*ptr = new
	return new
}

//
// Compare and swap
//

//go:nosplit
//go:noinline
func Cas32(ptr *uint32, old, new uint32) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//go:nosplit
//go:noinline
func Cas64(ptr *uint64, old, new uint64) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//go:nosplit
//go:noinline
func CasUintptr(ptr *uintptr, old, new uintptr) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//go:nosplit
//go:noinline
func CasUnsafePointer(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//go:nosplit
//go:noinline
func CasInt32(ptr *int32, old, new int32) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//go:nosplit
//go:noinline
func CasInt64(ptr *int64, old, new int64) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}

//
// CasRel
//

//go:nosplit
//go:noinline
func CasRel32(ptr *uint32, old, new uint32) bool {
	if *ptr == old {
		*ptr = new
		return true
	}
	return false
}
