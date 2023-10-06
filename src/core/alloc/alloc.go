// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alloc

import (
	"unsafe"

	stdtype "github.com/primecitizens/pcz/std/builtin/type"
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/core/math"
	"github.com/primecitizens/pcz/std/core/mem"
	"github.com/primecitizens/pcz/std/core/os"
)

// Hint for freeing the allocated memory
//
// TODO: TBD
type Hint uint32

// P for persistant allocations.
type P interface {
	Palloc(size uintptr) unsafe.Pointer
}

// M for general purpose allocations.
type M interface {
	// Malloc allocates memory for n elements of typ.
	//
	// NOTE: for bytes allocations, typ may be nil.
	Malloc(typ *abi.Type, n uintptr, zeroize bool) unsafe.Pointer

	// Free memory allocated for n elements.
	//
	// NOTE: for freeing types without heap pointer, typ may be nil
	// and in that case, n is the bytes allocated.
	Free(typ *abi.Type, n uintptr, ptr unsafe.Pointer) Hint
}

func SizeOutOfRange(typSz uintptr, n int) (outOfRange bool) {
	if n < 0 {
		return true
	}

	mem, overflow := math.MulUintptr(typSz, uintptr(n))
	if overflow || mem > os.MaxAlloc {
		return true
	}

	return false
}

func elemTypeOf(ptr any) *abi.Type {
	return stdtype.TypeOf(ptr).PointerTypeUnsafe().Elem
}

func New[T any](alloc M) *T {
	return (*T)(alloc.Malloc(elemTypeOf((*T)(nil)), 1, true))
}

func Make[E any](alloc M, len, cap int) []E {
	return unsafe.Slice(
		(*E)(MakeTyped(alloc, elemTypeOf((*E)(nil)), len, cap)),
		cap,
	)[:len]
}

func MakeTyped(alloc M, et *abi.Type, len, cap int) unsafe.Pointer {
	if len > cap {
		assert.Throw("makeslice:", "len", ">", "cap")
	}

	if et != nil && SizeOutOfRange(et.Size_, cap) {
		assert.Throw("makeslice:", "size", "out", "of", "range")
	}

	return alloc.Malloc(et, uintptr(cap), true)
}

func Clone[T any](alloc M, x T) T {
	typ := elemTypeOf((*T)(nil))
	switch typ.Kind() {
	case abi.KindString:
		str := *(*string)(mark.NoEscapePointer(&x))
		data := unsafe.Slice(
			(*byte)(alloc.Malloc(nil, uintptr(len(str)), false)),
			len(str),
		)

		str = unsafe.String(unsafe.SliceData(data), copy(data, str))
		return *(*T)(mark.NoEscapePointer(&str))
	case abi.KindSlice:
		slice := *(*[]byte)(mark.NoEscapePointer(&x))
		typ = typ.SliceType().Elem
		data := MakeTyped(alloc, typ, len(slice), len(slice))
		mem.Move(data, mark.NoEscapeSliceDataPointer(slice), uintptr(len(slice))*typ.Size_)
		return *(*T)(mark.NoEscapePointer(&slice))
	case abi.KindPointer:
		typ = typ.PointerType().Elem
		ptr := alloc.Malloc(typ, 1, false)
		mem.TypedMove(typ, ptr, mark.NoEscapePointer(&x))
		return *(*T)(mark.NoEscapePointer(&ptr))
	default:
		assert.TODO()
	}

	return x
}

func Free[T any](alloc M, x T) Hint {
	typ := elemTypeOf((*T)(nil))
	switch typ.Kind() {
	case abi.KindPointer:
		return alloc.Free(typ.PointerType().Elem, 1, *(*unsafe.Pointer)(unsafe.Pointer(&x)))
	case abi.KindString:
		str := *(*string)(mark.NoEscapePointer(&x))
		return alloc.Free(nil, uintptr(len(str)), unsafe.Pointer(unsafe.StringData(str)))
	case abi.KindSlice:
		slice := *(*[]byte)(mark.NoEscapePointer(&x))
		return alloc.Free(typ.SliceType().Elem, uintptr(len(slice)), unsafe.Pointer(unsafe.SliceData(slice)))
	default:
		assert.TODO()
	}

	return 0
}

// DenyNoneZero implements M that only allows allocation of zero sized types.
type DenyNoneZero struct{}

func (DenyNoneZero) Malloc(typ *abi.Type, n uintptr, zeroize bool) unsafe.Pointer {
	if n == 0 {
		return ZeroSized()
	}

	assert.Throw("no", "non-zero", "allocation", "allowed")
	return nil
}

func (DenyNoneZero) Palloc(size uintptr) unsafe.Pointer {
	if size == 0 {
		return ZeroSized()
	}

	assert.Throw("no", "non-zero", "allocation", "allowed")
	return nil
}

func (DenyNoneZero) Free(typ *abi.Type, n uintptr, ptr unsafe.Pointer) Hint { return 0 }

// S for stack allocations.
type S interface {
	NewStack(size uintptr) unsafe.Pointer
}

// H for heap allocations.
type H interface {
	M
	MallocHeap(typ *abi.Type, n uintptr, zerioze bool) unsafe.Pointer
}
