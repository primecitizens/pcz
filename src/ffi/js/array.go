// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/iter"
	"github.com/primecitizens/std/core/math"
	"github.com/primecitizens/std/ffi/js/bindings"
)

func NewArray(sz int) Array {
	if sz < 0 {
		assert.Throw("invalid", "negative", "capacity")
	}

	if uint64(sz) > uint64(math.MaxUint32) {
		assert.Throw("array", "size", "too", "large")
	}

	return Array{
		ref:    bindings.Array(SizeU(sz), 0, 0, 0),
		offset: 0,
		len:    0,
		cap:    sz,
	}
}

type Array struct {
	ref    bindings.Ref
	offset uint32
	len    int
	cap    int
}

func (a Array) Ref() Ref {
	return Ref(a.ref)
}

func (a Array) Len() int {
	return a.len
}

func (a Array) Cap() int {
	return a.cap
}

func (a Array) Slice(start, end int) Array {
	return a
}

func (a Array) SliceFrom(start int) Array {
	return a
}

func (a Array) Append(free bool, elems ...Ref) Array {
	bindings.Append(
		a.ref,
		bindings.Ref(Bool(free)),
		0, // elemSz (for typed array)
		0, // signed (for typed array)
		0, // float (for typed array)
		a.offset+SizeU(a.len),
		SliceData(elems),
		SizeU(len(elems)),
	)
	return a
}

func (a Array) Into(buf []Ref, offset int) int {
	return int(bindings.Copy(
		a.ref,
		0, // elemSz (for typed array)
		0, // signed (for typed array)
		0, // float (for typed array)
		SizeU(offset)+a.offset,
		SliceData(buf),
		SizeU(len(buf)),
	))
}

func (a Array) Nth(i int) (Ref, bool) {
	i, ok := iter.Index(i, a.cap)
	if !ok {
		return 0, false
	}

	var ret Ref
	if bindings.Ref(True) == bindings.Nth(
		a.ref,
		0, // elemSz (for typed array)
		0, // signed (for typed array)
		0, // float (for typed array)
		SizeU(i)+a.offset,
		Pointer(&ret),
	) {
		return ret, true
	}

	return ret, false
}

func (a Array) NthNum(i int) (Number, bool) {
	i, ok := iter.Index(i, a.cap)
	if !ok {
		return 0, false
	}

	var ret Number
	if bindings.Ref(True) == bindings.NthNum(
		a.ref,
		SizeU(i)+a.offset,
		Pointer(&ret),
	) {
		return ret, true
	}

	return ret, false
}

func (a Array) NthBool(i int) (bool, bool) {
	var ret bool
	if bindings.Ref(True) == bindings.NthBool(
		a.ref,
		SizeU(i)+a.offset,
		Pointer(&ret),
	) {
		return ret, true
	}

	return ret, false
}

func (a Array) SetNth(i int, free bool, val Ref) bool {
	return bindings.Ref(True) == bindings.SetNth(
		a.ref,
		int32(i),
		bindings.Ref(Bool(free)),
		bindings.Ref(val),
	)
}

func (a Array) SetNthNum(i int, val Number) bool {
	return bindings.Ref(True) == bindings.SetNthNum(
		a.ref,
		int32(i),
		float64(val),
	)
}

func (a Array) SetNthBool(i int, val bool) bool {
	return bindings.Ref(True) == bindings.SetNthBool(
		a.ref,
		int32(i),
		bindings.Ref(Bool(val)),
	)
}

func (a Array) SetNthString(i int, s string) bool {
	return bindings.Ref(True) == bindings.SetNthString(
		a.ref,
		int32(i),
		StringData(s),
		SizeU(len(s)),
	)
}

func (a Array) Free() {
	bindings.Free(a.ref)
}

type (
	ArrayBuffer       Ref
	SharedArrayBuffer Ref
	DataView          Ref
)

type Uint8ClampedArray[T ~uint8] struct {
	TypedArray[T]
}

type (
	Int8Array      = TypedArray[int8]
	Int16Array     = TypedArray[int16]
	Int32Array     = TypedArray[int32]
	Uint8Array     = TypedArray[uint8]
	Uint16Array    = TypedArray[uint16]
	Uint32Array    = TypedArray[uint32]
	Float32Array   = TypedArray[float32]
	Float64Array   = TypedArray[float64]
	BigInt64Array  = TypedArray[int64]
	BigUint64Array = TypedArray[uint64]
)

type elem interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func NewTypedArray[T elem](sz int) TypedArray[T] {
	if sz < 0 {
		assert.Throw("invalid", "negative", "capacity")
	}

	if uint64(sz) > uint64(math.MaxUint32) {
		assert.Throw("array", "size", "too", "large")
	}

	elemSz, signed, float := checkType[T]()
	return TypedArray[T]{
		ref: bindings.Array(SizeU(sz), elemSz, signed, float),
		len: 0,
		cap: sz,
	}
}

type TypedArray[T elem] struct {
	ref    bindings.Ref
	offset uint32
	len    int
	cap    int
}

func (a TypedArray[T]) Ref() Ref {
	return Ref(a.ref)
}

func (a TypedArray[T]) Free() {
	bindings.Free(a.ref)
}

func (a TypedArray[T]) Len() int {
	return a.len
}

func (a TypedArray[T]) Cap() int {
	return a.cap
}

func (a TypedArray[T]) SliceFrom(i int) (TypedArray[T], bool) {
	i, ok := iter.Bound(i, a.cap)
	if !ok {
		return TypedArray[T]{}, false
	}

	return TypedArray[T]{
		ref:    a.ref,
		offset: a.offset + SizeU(i),
		len:    a.len - i,
		cap:    a.cap - i,
	}, false
}

func (a TypedArray[T]) Nth(i int) (T, bool) {
	var out T
	elemSz, signed, float := checkType[T]()
	ok := bindings.Ref(True) == bindings.Nth(
		a.ref,
		elemSz,
		signed,
		float,
		SizeU(i)+a.offset,
		Pointer(&out),
	)
	return out, ok
}

func (a *TypedArray[T]) Append(elems ...T) *TypedArray[T] {
	elemSz, signed, float := checkType[T]()
	a.len += int(bindings.Append(
		a.ref,
		bindings.Ref(True), // free (nop for typed array)
		elemSz,
		signed,
		float,
		SizeU(a.len)+a.offset,
		SliceData(elems),
		SizeU(len(elems)),
	))
	return a
}

func (a TypedArray[T]) Into(buf []T, offset int) int {
	elemSz, signed, float := checkType[T]()
	return int(bindings.Copy(
		a.ref,
		elemSz,
		signed,
		float,
		SizeU(offset)+a.offset,
		SliceData(buf),
		SizeU(len(buf)),
	))
}

func checkType[T elem]() (elemSz uint32, signed, float bindings.Ref) {
	var x T = 1
	elemSz = uint32(unsafe.Sizeof(x))
	x /= 2
	if x > 0 {
		return elemSz, bindings.Ref(True), bindings.Ref(True)
	}

	x = 1
	x = -x
	if T(x) > 0 {
		return elemSz, bindings.Ref(False), bindings.Ref(False)
	}

	return elemSz, bindings.Ref(True), bindings.Ref(False)
}
