// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	"unsafe"

	"github.com/primecitizens/std/core/mark"
)

// Int32 is an atomically accessed int32 value.
//
// An Int32 must not be copied.
type Int32[T ~int32] struct {
	noCopy mark.NoCopy
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (i *Int32[T]) Load() T {
	return T(LoadInt32((*int32)(unsafe.Pointer(&i.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (i *Int32[T]) Store(value T) {
	StoreInt32((*int32)(unsafe.Pointer(&i.Value)), int32(value))
}

// Cas atomically compares i's value with old,
// and if they're equal, swaps i's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (i *Int32[T]) Cas(old, new T) bool {
	return CasInt32((*int32)(unsafe.Pointer(&i.Value)), int32(old), int32(new))
}

// Swap replaces i's value with new, returning
// i's value before the replacement.
//
//go:nosplit
func (i *Int32[T]) Swap(new T) T {
	return T(SwapInt32((*int32)(unsafe.Pointer(&i.Value)), int32(new)))
}

// Add adds delta to i atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (i *Int32[T]) Add(delta T) T {
	return T(AddInt32((*int32)(unsafe.Pointer(&i.Value)), int32(delta)))
}

// Int64 is an atomically accessed int64 value.
//
// 8-byte aligned on all platforms, unlike a regular int64.
//
// An Int64 must not be copied.
type Int64[T ~int64] struct {
	noCopy mark.NoCopy
	_      mark.Align64
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (i *Int64[T]) Load() T {
	return T(LoadInt64((*int64)(unsafe.Pointer(&i.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (i *Int64[T]) Store(value T) {
	StoreInt64((*int64)(unsafe.Pointer(&i.Value)), int64(value))
}

// Cas atomically compares i's value with old,
// and if they're equal, swaps i's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (i *Int64[T]) Cas(old, new T) bool {
	return CasInt64((*int64)(unsafe.Pointer(&i.Value)), int64(old), int64(new))
}

// Swap replaces i's value with new, returning
// i's value before the replacement.
//
//go:nosplit
func (i *Int64[T]) Swap(new T) T {
	return T(SwapInt64((*int64)(unsafe.Pointer(&i.Value)), int64(new)))
}

// Add adds delta to i atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (i *Int64[T]) Add(delta T) T {
	return T(AddInt64((*int64)(unsafe.Pointer(&i.Value)), int64(delta)))
}

// Uint8 is an atomically accessed uint8 value.
//
// A Uint8 must not be copied.
type Uint8[T ~uint8] struct {
	noCopy mark.NoCopy
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint8[T]) Load() T {
	return T(Load8((*uint8)(unsafe.Pointer(&u.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint8[T]) Store(value T) {
	Store8((*uint8)(unsafe.Pointer(&u.Value)), uint8(value))
}

// And takes value and performs a bit-wise
// "and" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint8[T]) And(value T) {
	And8((*uint8)(unsafe.Pointer(&u.Value)), uint8(value))
}

// Or takes value and performs a bit-wise
// "or" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint8[T]) Or(value T) {
	Or8((*uint8)(unsafe.Pointer(&u.Value)), uint8(value))
}

// Bool is an atomically accessed bool value.
//
// A Bool must not be copied.
type Bool[T ~bool] struct {
	// Inherits noCopy from Uint8.
	Uint8[uint8]
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (b *Bool[T]) Load() T {
	return b.Uint8.Load() != 0
}

// Store updates the value atomically.
//
//go:nosplit
func (b *Bool[T]) Store(value T) {
	if value {
		b.Uint8.Store(1)
	} else {
		b.Uint8.Store(0)
	}
}

// Uint32 is an atomically accessed uint32 value.
//
// A Uint32 must not be copied.
type Uint32[T ~uint32] struct {
	noCopy mark.NoCopy
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint32[T]) Load() T {
	return T(Load32((*uint32)(unsafe.Pointer(&u.Value))))
}

// LoadAcquire is a partially unsynchronized version
// of Load that relaxes ordering constraints. Other threads
// may observe operations that precede this operation to
// occur after it, but no operation that occurs after it
// on this thread can be observed to occur before it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32[T]) LoadAcquire() T {
	return T(LoadAcq32((*uint32)(unsafe.Pointer(&u.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint32[T]) Store(value T) {
	Store32((*uint32)(unsafe.Pointer(&u.Value)), uint32(value))
}

// StoreRelease is a partially unsynchronized version
// of Store that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32[T]) StoreRelease(value T) {
	StoreRel32((*uint32)(unsafe.Pointer(&u.Value)), uint32(value))
}

// Cas atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uint32[T]) Cas(old, new T) bool {
	return CasRel32((*uint32)(unsafe.Pointer(&u.Value)), uint32(old), uint32(new))
}

// CasRelease is a partially unsynchronized version
// of Cas that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
// It reports whether the swap ran.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uint32[T]) CasRelease(old, new T) bool {
	return CasRel32((*uint32)(unsafe.Pointer(&u.Value)), uint32(old), uint32(new))
}

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uint32[T]) Swap(value T) T {
	return T(Swap32((*uint32)(unsafe.Pointer(&u.Value)), uint32(value)))
}

// And takes value and performs a bit-wise
// "and" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint32[T]) And(value T) {
	And32((*uint32)(unsafe.Pointer(&u.Value)), uint32(value))
}

// Or takes value and performs a bit-wise
// "or" operation with the value of u, storing
// the result into u.
//
// The full process is performed atomically.
//
//go:nosplit
func (u *Uint32[T]) Or(value T) {
	Or32((*uint32)(unsafe.Pointer(&u.Value)), uint32(value))
}

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uint32[T]) Add(delta int32) T {
	return T(Add32((*uint32)(unsafe.Pointer(&u.Value)), delta))
}

// Uint64 is an atomically accessed uint64 value.
//
// 8-byte aligned on all platforms, unlike a regular uint64.
//
// A Uint64 must not be copied.
type Uint64[T ~uint64] struct {
	noCopy mark.NoCopy
	_      mark.Align64
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uint64[T]) Load() T {
	return T(Load64((*uint64)(unsafe.Pointer(&u.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (u *Uint64[T]) Store(value T) {
	Store64((*uint64)(unsafe.Pointer(&u.Value)), uint64(value))
}

// Cas atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uint64[T]) Cas(old, new T) bool {
	return Cas64((*uint64)(unsafe.Pointer(&u.Value)), uint64(old), uint64(new))
}

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uint64[T]) Swap(value T) T {
	return T(Swap64((*uint64)(unsafe.Pointer(&u.Value)), uint64(value)))
}

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uint64[T]) Add(delta int64) T {
	return T(Add64((*uint64)(unsafe.Pointer(&u.Value)), delta))
}

// Uintptr is an atomically accessed uintptr value.
//
// A Uintptr must not be copied.
type Uintptr[T ~uintptr] struct {
	noCopy mark.NoCopy
	Value  T
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *Uintptr[T]) Load() T {
	return T(LoadUintptr((*uintptr)(unsafe.Pointer(&u.Value))))
}

// LoadAcquire is a partially unsynchronized version
// of Load that relaxes ordering constraints. Other threads
// may observe operations that precede this operation to
// occur after it, but no operation that occurs after it
// on this thread can be observed to occur before it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uintptr[T]) LoadAcquire() T {
	return T(LoadAcqUintptr((*uintptr)(unsafe.Pointer(&u.Value))))
}

// Store updates the value atomically.
//
//go:nosplit
func (u *Uintptr[T]) Store(value T) {
	StoreUintptr((*uintptr)(unsafe.Pointer(&u.Value)), uintptr(value))
}

// StoreRelease is a partially unsynchronized version
// of Store that relaxes ordering constraints. Other threads
// may observe operations that occur after this operation to
// precede it, but no operation that precedes it
// on this thread can be observed to occur after it.
//
// WARNING: Use sparingly and with great care.
//
//go:nosplit
func (u *Uintptr[T]) StoreRelease(value T) {
	StoreRelUintptr((*uintptr)(unsafe.Pointer(&u.Value)), uintptr(value))
}

// Cas atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
//
//go:nosplit
func (u *Uintptr[T]) Cas(old, new T) bool {
	return CasUintptr((*uintptr)(unsafe.Pointer(&u.Value)), uintptr(old), uintptr(new))
}

// Swap replaces u's value with new, returning
// u's value before the replacement.
//
//go:nosplit
func (u *Uintptr[T]) Swap(value T) T {
	return T(SwapUintptr((*uintptr)(unsafe.Pointer(&u.Value)), uintptr(value)))
}

// Add adds delta to u atomically, returning
// the new updated value.
//
// This operation wraps around in the usual
// two's-complement way.
//
//go:nosplit
func (u *Uintptr[T]) Add(delta T) T {
	return T(AddUintptr((*uintptr)(unsafe.Pointer(&u.Value)), uintptr(delta)))
}

// Float64 is an atomically accessed float64 value.
//
// 8-byte aligned on all platforms, unlike a regular float64.
//
// A Float64 must not be copied.
type Float64[T ~float64] struct {
	// Inherits noCopy and align64 from Uint64.
	Uint64[uint64]
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (f *Float64[T]) Load() T {
	r := f.Uint64.Load()
	return T(*(*float64)(unsafe.Pointer(&r)))
}

// Store updates the value atomically.
//
//go:nosplit
func (f *Float64[T]) Store(value T) {
	f.Uint64.Store(*(*uint64)(unsafe.Pointer(&value)))
}

// UnsafePointer is an atomically accessed unsafe.Pointer value.
//
// Note that because of the atomicity guarantees, stores to values
// of this type never trigger a write barrier, and the relevant
// methods are suffixed with "NoWB" to indicate that explicitly.
// As a result, this type should be used carefully, and sparingly,
// mostly with values that do not live in the Go heap anyway.
//
// An UnsafePointer must not be copied.
type UnsafePointer struct {
	noCopy mark.NoCopy
	Value  unsafe.Pointer
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (u *UnsafePointer) Load() unsafe.Pointer {
	return LoadPointer(unsafe.Pointer(&u.Value))
}

// StoreNoWB updates the value atomically.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Store instead.
//
//go:nosplit
func (u *UnsafePointer) StoreNoWB(value unsafe.Pointer) {
	StorePointer(unsafe.Pointer(&u.Value), value)
}

// Store updates the value atomically.
func (u *UnsafePointer) Store(value unsafe.Pointer) {
	storePointer(&u.Value, value)
}

// provided by runtime
//
//go:linkname storePointer
func storePointer(ptr *unsafe.Pointer, new unsafe.Pointer)

// CasNoWB atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Cas instead.
//
//go:nosplit
func (u *UnsafePointer) CasNoWB(old, new unsafe.Pointer) bool {
	return CasUnsafePointer(&u.Value, old, new)
}

// Cas atomically compares u's value with old,
// and if they're equal, swaps u's value with new.
// It reports whether the swap ran.
func (u *UnsafePointer) Cas(old, new unsafe.Pointer) bool {
	return casPointer(&u.Value, old, new)
}

func casPointer(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool

// Pointer is an atomic pointer of type *T.
type Pointer[T any] struct {
	UnsafePointer
}

// Load accesses and returns the value atomically.
//
//go:nosplit
func (p *Pointer[T]) Load() *T {
	return (*T)(p.UnsafePointer.Load())
}

// Store updates the value atomically.
//
//go:nosplit
func (p *Pointer[T]) Store(value *T) {
	p.UnsafePointer.Store(unsafe.Pointer(value))
}

// StoreNoWB updates the value atomically.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Store instead.
//
//go:nosplit
func (p *Pointer[T]) StoreNoWB(value *T) {
	p.UnsafePointer.StoreNoWB(unsafe.Pointer(value))
}

// Cas atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
func (p *Pointer[T]) Cas(old, new *T) bool {
	return p.UnsafePointer.Cas(unsafe.Pointer(old), unsafe.Pointer(new))
}

// CasNoWB atomically (with respect to other methods)
// compares u's value with old, and if they're equal,
// swaps u's value with new.
// It reports whether the swap ran.
//
// WARNING: As the name implies this operation does *not*
// perform a write barrier on value, and so this operation may
// hide pointers from the GC. Use with care and sparingly.
// It is safe to use with values not found in the Go heap.
// Prefer Cas instead.
//
//go:nosplit
func (p *Pointer[T]) CasNoWB(old, new *T) bool {
	return p.UnsafePointer.CasNoWB(unsafe.Pointer(old), unsafe.Pointer(new))
}
