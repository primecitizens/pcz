// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/math"
	"github.com/primecitizens/pcz/std/core/num"
	"github.com/primecitizens/pcz/std/ffi/js/array"
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

func NewArray[T any](sz int) Array[T] {
	if sz < 0 {
		assert.Throw("invalid", "negative", "capacity")
	}

	if uint64(sz) > uint64(math.MaxUint32) {
		assert.Throw("array", "size", "too", "large")
	}

	return Array[T]{
		ref: array.New(SizeU(sz), 0, 0, 0),
	}
}

func NewArrayOf[T any](take bool, elems ...Ref) Array[T] {
	return NewArray[T](len(elems)).MustFill(0, take, elems...)
}

type Array[T any] struct {
	ref bindings.Ref
}

func (a Array[T]) FromRef(ref Ref) Array[T] {
	return Array[T]{
		ref: bindings.Ref(ref),
	}
}

func (a Array[T]) Ref() Ref {
	return Ref(a.ref)
}

func (a Array[T]) Once() Array[T] {
	bindings.Once(a.ref)
	return a
}

func (a Array[T]) Length() int {
	return int(array.Length(a.ref))
}

func (a Array[T]) Slice(start, end int) Array[T] {
	return Array[T]{
		ref: array.Slice(a.ref, SizeU(start), SizeU(end)),
	}
}

func (a Array[T]) SliceFrom(start int) Array[T] {
	return Array[T]{
		ref: array.Slice(a.ref, SizeU(start), 0),
	}
}

func (a Array[T]) Fill(start int, take bool, elems ...Ref) int {
	return int(
		array.Append(
			a.ref,
			bindings.Ref(Bool(take)),
			0,                   // elemSz (for typed array only)
			bindings.Ref(False), // signed (for typed array only)
			bindings.Ref(False), // float (for typed array only)
			SizeU(start),        // offset
			SliceData(elems),
			SizeU(len(elems)),
		),
	)
}

func (a Array[T]) MustFill(start int, take bool, elems ...Ref) Array[T] {
	if len(elems) != a.Fill(start, take, elems...) {
		assert.Throw("not", "all", "elements", "inserted")
	}

	return a
}

func (a Array[T]) Copy(start int, outBuf []Ref) int {
	return int(
		array.Copy(
			a.ref,
			0,                   // elemSz (for typed array only)
			bindings.Ref(False), // signed (for typed array only)
			bindings.Ref(False), // float (for typed array only)
			SizeU(start),
			SliceData(outBuf),
			SizeU(len(outBuf)),
		),
	)
}

func (a Array[T]) Nth(i int) (Ref, bool) {
	var ret Ref
	if bindings.Ref(True) == array.Nth(
		a.ref,
		0,                   // elemSz (for typed array only)
		bindings.Ref(False), // signed (for typed array only)
		bindings.Ref(False), // float (for typed array only)
		SizeU(i),
		Pointer(&ret),
	) {
		return ret, true
	}

	return Undefined, false
}

func (a Array[T]) NthNum(i int) (float64, bool) {
	var ret float64
	if bindings.Ref(True) == array.NthNum(
		a.ref,
		SizeU(i),
		Pointer(&ret),
	) {
		return ret, true
	}

	return 0, false
}

func (a Array[T]) NthBool(i int) (bool, bool) {
	var ret bool
	if bindings.Ref(True) == array.NthBool(
		a.ref,
		SizeU(i),
		Pointer(&ret),
	) {
		return ret, true
	}

	return false, false
}

func (a Array[T]) SetNth(i int, take bool, val Ref) bool {
	return bindings.Ref(True) == array.SetNth(
		a.ref,
		0,                   // elemSz (for typed array only)
		bindings.Ref(False), // signed (for typed array only)
		bindings.Ref(False), // float (for typed array only)
		SizeU(i),
		bindings.Ref(Bool(take)),
		Pointer(&val),
	)
}

func (a Array[T]) SetNthNum(i int, val float64) bool {
	return bindings.Ref(True) == array.SetNthNum(
		a.ref,
		SizeU(i),
		val,
	)
}

func (a Array[T]) SetNthBool(i int, val bool) bool {
	return bindings.Ref(True) == array.SetNthBool(
		a.ref,
		SizeU(i),
		bindings.Ref(Bool(val)),
	)
}

func (a Array[T]) SetNthString(i int, s string) bool {
	return bindings.Ref(True) == array.SetNthString(
		a.ref,
		SizeU(i),
		StringData(s),
		SizeU(len(s)),
	)
}

func (a Array[T]) Free() {
	bindings.Free(a.ref)
}

// TODO
type FrozenArray[T any] struct {
	ref bindings.Ref
}

func (a FrozenArray[T]) Ref() Ref {
	return Ref(a.ref)
}

func (a FrozenArray[T]) FromRef(ref Ref) FrozenArray[T] {
	return FrozenArray[T]{
		ref: bindings.Ref(ref),
	}
}

// TODO
type ObservableArray[T any] struct {
	ref bindings.Ref
}

func (a ObservableArray[T]) Ref() Ref {
	return Ref(a.ref)
}

func (a ObservableArray[T]) FromRef(ref Ref) ObservableArray[T] {
	return ObservableArray[T]{
		ref: bindings.Ref(ref),
	}
}

// TODO
type DataView struct {
	ref bindings.Ref
}

func (v DataView) Ref() Ref {
	return Ref(v.ref)
}

func (v DataView) FromRef(ref Ref) DataView {
	return DataView{
		ref: bindings.Ref(ref),
	}
}

// TODO
type ArrayBuffer struct {
	ref bindings.Ref
}

func (buf ArrayBuffer) FromRef(ref Ref) ArrayBuffer {
	return ArrayBuffer{
		ref: bindings.Ref(ref),
	}
}

func (buf ArrayBuffer) Ref() Ref {
	return Ref(buf.ref)
}

func (buf ArrayBuffer) Once() ArrayBuffer {
	bindings.Once(buf.ref)
	return buf
}

func (buf ArrayBuffer) Free() {
	bindings.Free(buf.ref)
	return
}

// TODO
type SharedArrayBuffer struct {
	ref bindings.Ref
}

func (buf SharedArrayBuffer) FromRef(ref Ref) SharedArrayBuffer {
	return SharedArrayBuffer{
		ref: bindings.Ref(ref),
	}
}

func (buf SharedArrayBuffer) Ref() Ref {
	return Ref(buf.ref)
}

func (buf SharedArrayBuffer) Once() SharedArrayBuffer {
	bindings.Once(buf.ref)
	return buf
}

func (buf SharedArrayBuffer) Free() {
	bindings.Free(buf.ref)
	return
}

// TODO
type ArrayBufferView struct {
	ref bindings.Ref
}

func (v ArrayBufferView) FromRef(ref Ref) ArrayBufferView {
	return ArrayBufferView{
		ref: bindings.Ref(ref),
	}
}

func (v ArrayBufferView) Ref() Ref {
	return Ref(v.ref)
}

func (v ArrayBufferView) Once() ArrayBufferView {
	bindings.Once(v.ref)
	return v
}

func (v ArrayBufferView) Free() {
	bindings.Free(v.ref)
	return
}

type (
	Int8Array         = TypedArray[int8]
	Int16Array        = TypedArray[int16]
	Int32Array        = TypedArray[int32]
	Uint8Array        = TypedArray[uint8]
	Uint8ClampedArray = TypedArray[uint8]
	Uint16Array       = TypedArray[uint16]
	Uint32Array       = TypedArray[uint32]
	Float32Array      = TypedArray[float32]
	Float64Array      = TypedArray[float64]
	BigInt64Array     = TypedArray[int64]
	BigUint64Array    = TypedArray[uint64]
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

	elemSz, signed, float := num.CheckType[T]()
	return TypedArray[T]{
		ref: array.New(
			SizeU(sz),
			SizeU(elemSz),
			bindings.Ref(Bool(signed)),
			bindings.Ref(Bool(float)),
		),
	}
}

func NewTypedArrayOf[T elem](elems ...T) TypedArray[T] {
	return NewTypedArray[T](len(elems)).MustFill(0, elems...)
}

type TypedArray[T elem] struct {
	ref bindings.Ref
}

func (a TypedArray[T]) FromRef(ref Ref) TypedArray[T] {
	return TypedArray[T]{
		ref: bindings.Ref(ref),
	}
}

func (a TypedArray[T]) Ref() Ref {
	return Ref(a.ref)
}

func (a TypedArray[T]) Once() TypedArray[T] {
	a.Ref().Once()
	return a
}

func (a TypedArray[T]) Free() {
	bindings.Free(a.ref)
}

func (a TypedArray[T]) FromArrayBuffer(take bool, buf ArrayBuffer) TypedArray[T] {
	elemSz, signed, float := num.CheckType[T]()
	return TypedArray[T]{
		ref: array.FromBuffer(
			bindings.Ref(Bool(take)),
			buf.ref,
			SizeU(elemSz),
			bindings.Ref(Bool(signed)),
			bindings.Ref(Bool(float)),
		),
	}
}

func (a TypedArray[T]) FromSharedArrayBuffer(take bool, buf SharedArrayBuffer) TypedArray[T] {
	elemSz, signed, float := num.CheckType[T]()
	return TypedArray[T]{
		ref: array.FromBuffer(
			bindings.Ref(Bool(take)),
			buf.ref,
			SizeU(elemSz),
			bindings.Ref(Bool(signed)),
			bindings.Ref(Bool(float)),
		),
	}
}

func (a TypedArray[T]) Buffer(take bool) Ref {
	return Ref(array.Buffer(bindings.Ref(Bool(take)), a.ref))
}

func (a TypedArray[T]) ByteOffset() int {
	return int(array.ByteOffset(a.ref))
}

func (a TypedArray[T]) ByteLength() int {
	return int(array.ByteLength(a.ref))
}

func (a TypedArray[T]) Length() int {
	return int(array.Length(a.ref))
}

func (a TypedArray[T]) Slice(start, end int) TypedArray[T] {
	return TypedArray[T]{
		ref: array.Slice(
			a.ref, SizeU(start), SizeU(end),
		),
	}
}

func (a TypedArray[T]) SliceFrom(start int) TypedArray[T] {
	return TypedArray[T]{
		ref: array.Slice(
			a.ref, SizeU(start), 0,
		),
	}
}

func (a TypedArray[T]) Nth(i int) (T, bool) {
	var out T
	elemSz, signed, float := num.CheckType[T]()
	if bindings.Ref(True) == array.Nth(
		a.ref,
		SizeU(elemSz),
		bindings.Ref(Bool(signed)),
		bindings.Ref(Bool(float)),
		SizeU(i),
		Pointer(&out),
	) {
		return 0, false
	}

	return out, true
}

func (a TypedArray[T]) SetNth(i int, val T) bool {
	elemSz, signed, float := num.CheckType[T]()
	return bindings.Ref(True) == array.SetNth(
		a.ref,
		SizeU(elemSz),
		bindings.Ref(Bool(signed)),
		bindings.Ref(Bool(float)),
		SizeU(i),
		bindings.Ref(False),
		Pointer(&val),
	)
}

func (a TypedArray[T]) Fill(start int, elems ...T) int {
	elemSz, signed, float := num.CheckType[T]()
	return int(
		array.Append(
			a.ref,
			bindings.Ref(False), // take (nop for typed array)
			SizeU(elemSz),
			bindings.Ref(Bool(signed)),
			bindings.Ref(Bool(float)),
			SizeU(start),
			SliceData(elems),
			SizeU(len(elems)),
		),
	)
}

func (a TypedArray[T]) MustFill(start int, elems ...T) TypedArray[T] {
	if len(elems) != a.Fill(start, elems...) {
		assert.Throw("not", "all", "elements", "inserted")
	}

	return a
}

func (a TypedArray[T]) Set(other TypedArray[T]) TypedArray[T] {
	array.Set(a.ref, other.ref)
	return a
}

func (a TypedArray[T]) Copy(start int, outBuf []T) int {
	elemSz, signed, float := num.CheckType[T]()
	return int(
		array.Copy(
			a.ref,
			SizeU(elemSz),
			bindings.Ref(Bool(signed)),
			bindings.Ref(Bool(float)),
			SizeU(start),
			SliceData(outBuf),
			SizeU(len(outBuf)),
		),
	)
}
