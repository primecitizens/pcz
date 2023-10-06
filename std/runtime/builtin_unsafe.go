// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	stdptr "github.com/primecitizens/pcz/std/builtin/ptr"
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/math"
)

// functions required by compiler for unsafe.Slice and unsafe.String
//
// see $GOROOT/src/runtime/unsafe.go

//
// unsafe.Slice
//

// This is called only from compiler-generated code, so we can get the
// source of the panic.
func panicunsafeslicelen()    { panicunsafeslicelen1(getcallerpc()) }
func panicunsafeslicenilptr() { panicunsafeslicenilptr1(getcallerpc()) }

func panicunsafeslicelen1(pc uintptr) {
	// panicCheck1(pc, "unsafe.Slice: len out of range")
	assert.Panic("unsafe.Slice: len out of range")
}

func panicunsafeslicenilptr1(pc uintptr) {
	// panicCheck1(pc, "unsafe.Slice: ptr is nil and len is not zero")
	assert.Panic("unsafe.Slice:", "ptr", "is", "nil", "and", "len", "is", "not", "zero")
}

// Keep this code in sync with cmd/compile/internal/walk/builtin.go:walkUnsafeSlice
func unsafeslice(et *abi.Type, ptr unsafe.Pointer, len int) {
	if len < 0 {
		panicunsafeslicelen1(getcallerpc())
	}

	if et.Size_ == 0 {
		if ptr == nil && len > 0 {
			panicunsafeslicenilptr1(getcallerpc())
		}
	}

	mem, overflow := math.MulUintptr(et.Size_, uintptr(len))
	if overflow || mem > -uintptr(ptr) {
		if ptr == nil {
			panicunsafeslicenilptr1(getcallerpc())
		}
		panicunsafeslicelen1(getcallerpc())
	}
}

// Keep this code in sync with cmd/compile/internal/walk/builtin.go:walkUnsafeSlice
func unsafeslice64(et *abi.Type, ptr unsafe.Pointer, len64 int64) {
	len := int(len64)
	if int64(len) != len64 {
		panicunsafeslicelen1(getcallerpc())
	}
	unsafeslice(et, ptr, len)
}

func unsafeslicecheckptr(et *abi.Type, ptr unsafe.Pointer, len64 int64) {
	unsafeslice64(et, ptr, len64)

	// Check that underlying array doesn't straddle multiple heap objects.
	// unsafeslice64 has already checked for overflow.
	if stdptr.CheckptrStraddles(ptr, uintptr(len64)*et.Size_) {
		assert.Throw("checkptr:", "unsafe.Slice", "result", "straddles", "multiple", "allocations")
	}
}

//
// unsafe.String
//

func panicunsafestringlen() {
	assert.Panic("unsafe.String:", "len", "out", "of", "range")
}

func panicunsafestringnilptr() {
	assert.Panic("unsafe.String:", "ptr", "is", "nil", "and", "len", "is", "not", "zero")
}

func unsafestring(ptr unsafe.Pointer, len int) {
	if len < 0 {
		panicunsafestringlen()
	}

	if uintptr(len) > -uintptr(ptr) {
		if ptr == nil {
			panicunsafestringnilptr()
		}
		panicunsafestringlen()
	}
}

// Keep this code in sync with cmd/compile/internal/walk/builtin.go:walkUnsafeString
func unsafestring64(ptr unsafe.Pointer, len64 int64) {
	len := int(len64)
	if int64(len) != len64 {
		panicunsafestringlen()
	}
	unsafestring(ptr, len)
}

func unsafestringcheckptr(ptr unsafe.Pointer, len64 int64) {
	unsafestring64(ptr, len64)

	// Check that underlying array doesn't straddle multiple heap objects.
	// unsafestring64 has already checked for overflow.
	if stdptr.CheckptrStraddles(ptr, uintptr(len64)) {
		assert.Throw("checkptr:", "unsafe.String", "result", "straddles", "multiple", "allocations")
	}
}
