// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !wasm

package atomic

import "unsafe"

// PublicationBarrier performs a store/store barrier (a "publication"
// or "export" barrier). Some form of synchronization is required
// between initializing an object and making that object accessible to
// another processor. Without synchronization, the initialization
// writes and the "publication" write may be reordered, allowing the
// other processor to follow the pointer and observe an uninitialized
// object. In general, higher-level synchronization should be used,
// such as locking or an atomic pointer write. publicationBarrier is
// for when those aren't an option, such as in the implementation of
// the memory manager.
//
// There's no corresponding barrier for the read side because the read
// side naturally has a data dependency order. All architectures that
// Go supports or seems likely to ever support automatically enforce
// data dependency ordering.
func PublicationBarrier()

//
// Store
//

//go:noescape
func StoreInt32(ptr *int32, new int32)

//go:noescape
func StoreInt64(ptr *int64, new int64)

//go:noescape
func StoreUintptr(ptr *uintptr, new uintptr)

// TODO(matloob): Should these functions have the go:noescape annotation?

//
// Load
//

//go:noescape
func LoadUintptr(ptr *uintptr) uintptr

//go:noescape
func LoadUint(ptr *uint) uint

//go:noescape
func LoadInt32(ptr *int32) int32

//go:noescape
func LoadInt64(ptr *int64) int64

//
// Swap
//

//go:noescape
func SwapInt32(ptr *int32, new int32) int32

//go:noescape
func SwapInt64(ptr *int64, new int64) int64

//
// Add
//

//go:noescape
func AddInt32(ptr *int32, delta int32) int32

//go:noescape
func AddInt64(ptr *int64, delta int64) int64

//
// Compare and swap
//

//go:noescape
func Cas32(ptr *uint32, old, new uint32) bool

//go:noescape
func CasUintptr(ptr *uintptr, old, new uintptr) bool

// NO go:noescape annotation; see atomic_pointer.go.
func CasUnsafePointer(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool

//go:noescape
func CasInt32(ptr *int32, old, new int32) bool

//go:noescape
func CasInt64(ptr *int64, old, new int64) bool
