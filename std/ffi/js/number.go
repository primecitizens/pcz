// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/math"
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

type bigintTypes interface {
	~uint | ~uint64 | ~uintptr | ~int64
}

func NewBigInt[T bigintTypes](value T) BigInt[T] {
	if unsafe.Sizeof(value) != 8 {
		assert.Throw("unexpected", "size")
	}

	return BigInt[T]{
		ref: bindings.BigInt(
			bindings.Ref(Bool(isSigned[T]())),
			Pointer(&value),
		),
	}
}

type BigInt[T bigintTypes] struct {
	ref bindings.Ref
}

func (b BigInt[T]) FromRef(ref Ref) BigInt[T] {
	return BigInt[T]{
		ref: bindings.Ref(ref),
	}
}

func (b BigInt[T]) Ref() Ref {
	return Ref(b.ref)
}

func (b BigInt[T]) Once() BigInt[T] {
	bindings.Once(b.ref)
	return b
}

func (b BigInt[T]) Free() {
	bindings.Free(b.ref)
}

func (b BigInt[T]) Get() T {
	var value T
	if unsafe.Sizeof(value) != 8 {
		assert.Throw("unexpected", "size")
	}

	if bindings.Ref(True) != bindings.GetBigInt(
		b.ref,
		bindings.Ref(Bool(isSigned[T]())),
		Pointer(&value),
	) {
		assert.Throw("not", "a", "bigint")
	}

	return value
}

func (b BigInt[T]) Set(value T) bool {
	if unsafe.Sizeof(value) != 8 {
		assert.Throw("unexpected", "size")
	}

	return bindings.Ref(True) == bindings.ReplaceBigInt(
		b.ref,
		bindings.Ref(Bool(isSigned[T]())),
		Pointer(&value),
	)
}

type numTypes interface {
	~uint8 | ~uint16 | ~uint32 | ~int8 | ~int16 | ~int32 | ~float32 | ~float64
}

func NewNumber[T numTypes](x T) Number[T] {
	maxCache := smallIntCacheSize
	if x != x {
		return Number[T]{
			ref: bindings.Ref(NaN),
		}
	}

	_, _, float := checkType[T]()
	if float {
		switch math.Float64bits(float64(x)) {
		case math.Inf:
			return Number[T]{
				ref: bindings.Ref(Inf),
			}
		case math.NegInf:
			return Number[T]{
				ref: bindings.Ref(NegativeInf),
			}
		}
	}

	if x < T(0) || // negative number
		x > T(maxCache) || // uncached
		x/2 != 0 { // floating point
		return Number[T]{
			ref: bindings.Ref(bindings.Number(float64(x))),
		}
	}

	return Number[T]{
		ref: bindings.Ref(Ref(x) + firstSmallIntCache),
	}
}

// Number represents a js number living in the heap.
//
// NOTE: This type is defined for code not working directly with imported js functions
// (using go:wasmimport) and should be avoided whenever possible.
type Number[T numTypes] struct {
	ref bindings.Ref
}

func (b Number[T]) FromRef(ref Ref) Number[T] {
	return Number[T]{
		ref: bindings.Ref(ref),
	}
}

func (b Number[T]) Ref() Ref {
	return Ref(b.ref)
}

func (b Number[T]) Once() Number[T] {
	bindings.Once(b.ref)
	return b
}

func (b Number[T]) Free() {
	bindings.Free(b.ref)
}

func (b Number[T]) Get() T {
	elemSz, signed, float := checkType[T]()

	if b.ref < bindings.Ref(firstSmallIntCache) {
		var value float64
		switch b.ref {
		case bindings.Ref(NaN):
			value = math.Float64frombits(math.NaN)
		case bindings.Ref(Inf):
			value = math.Float64frombits(math.Inf)
		case bindings.Ref(NegativeInf):
			value = math.Float64frombits(math.NegInf)
		default:
			assert.Throw("not", "a", "number")
		}

		if !float {
			assert.Throw("not", "float")
		}

		switch elemSz {
		case 4, 8:
		default:
			assert.Unreachable()
		}

		return T(value)
	}

	if b.ref <= bindings.Ref(lastSmallIntCache) {
		x := b.ref - bindings.Ref(firstSmallIntCache)
		return T(x)
	}

	var v float64
	if bindings.Ref(True) != bindings.GetNum(b.ref, Pointer(&v)) {
		assert.Throw("not", "a", "number")
	}

	if float {
		// is floating-point type
		switch elemSz {
		case 4:
			if v > math.MaxFloat32 || v < -math.MaxFloat32 {
				assert.Throw("out", "of", "range")
			}
		case 8:
		default:
			assert.Unreachable()
		}

		return T(v)
	}

	// is integer type
	if signed {
		switch elemSz {
		case 1:
			if v > math.MaxInt8 || v < math.MinInt8 {
				assert.Throw("out", "of", "range")
			}
		case 2:
			if v > math.MaxInt16 || v < math.MinInt16 {
				assert.Throw("out", "of", "range")
			}
		case 4:
			if v > math.MaxInt32 || v < math.MinInt32 {
				assert.Throw("out", "of", "range")
			}
		default:
			assert.Unreachable()
		}
	}

	if v < 0 {
		assert.Throw("underflow")
	}

	switch elemSz {
	case 1:
		if v > math.MaxUint8 {
			assert.Throw("overflow")
		}
	case 2:
		if v > math.MaxUint16 {
			assert.Throw("overflow")
		}
	case 4:
		if v > math.MaxUint32 {
			assert.Throw("overflow")
		}
	default:
		assert.Unreachable()
	}

	return T(v)
}

func (b *Number[T]) Set(value T) bool {
	if b.ref <= bindings.Ref(lastSmallIntCache) {
		b.ref = NewNumber(value).ref
		return true
	}

	return bindings.Ref(True) == bindings.ReplaceNum(
		b.ref, float64(value),
	)
}

func (b Number[T]) IsNaN() bool {
	return b.ref == bindings.Ref(NaN)
}

func (b Number[T]) IsInf(sign int) bool {
	return (sign >= 0 && b.ref == bindings.Ref(Inf)) ||
		(sign <= 0 && b.ref == bindings.Ref(NegativeInf))
}
