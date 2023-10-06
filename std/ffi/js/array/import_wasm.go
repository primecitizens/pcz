// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package array

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

//go:wasmimport ffi/js/array new
//go:noescape
func New(
	cap uint32,
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
) bindings.Ref

//go:wasmimport ffi/js/array length
//go:noescape
func Length(arr bindings.Ref) int64

//go:wasmimport ffi/js/array slice
//go:noescape
func Slice(
	arr bindings.Ref,
	start uint32,
	end uint32,
) bindings.Ref

//go:wasmimport ffi/js/array set
//go:noescape
func Set(a, b bindings.Ref)

//go:wasmimport ffi/js/array append
//go:noescape
func Append(
	refJsArray bindings.Ref,
	take bindings.Ref, // for elemSz = 0 only
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
	offsetInJsArray uint32,
	goArray unsafe.Pointer,
	goArrayLen uint32,
) uint32

//go:wasmimport ffi/js/array copy
//go:noescape
func Copy(
	refJsArray bindings.Ref,
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
	offsetInJsArray uint32,
	outGoArray unsafe.Pointer,
	outGoArrayLen uint32,
) uint32

//go:wasmimport ffi/js/array buffer
//go:noescape
func Buffer(ref bindings.Ref, take bindings.Ref) bindings.Ref

//go:wasmimport ffi/js/array fromBuffer
//go:noescape
func FromBuffer(
	take bindings.Ref,
	refBuf bindings.Ref,
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
) bindings.Ref

//go:wasmimport ffi/js/array byteOffset
//go:noescape
func ByteOffset(ref bindings.Ref) uint32

//go:wasmimport ffi/js/array byteLength
//go:noescape
func ByteLength(ref bindings.Ref) uint32

//go:wasmimport ffi/js/array nth
//go:noescape
func Nth(
	arr bindings.Ref,
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
	nth uint32,
	ptr unsafe.Pointer,
) bindings.Ref

//go:wasmimport ffi/js/array nthNum
//go:noescape
func NthNum(
	self bindings.Ref,
	i uint32,
	outNum unsafe.Pointer,
) bindings.Ref

//go:wasmimport ffi/js/array nthBool
//go:noescape
func NthBool(
	self bindings.Ref,
	i uint32,
	outBool unsafe.Pointer,
) bindings.Ref

//go:wasmimport ffi/js/array setNth
//go:noescape
func SetNth(
	self bindings.Ref,
	elemSz uint32,
	signed bindings.Ref,
	float bindings.Ref,
	i uint32,
	take bindings.Ref,
	ptr unsafe.Pointer,
) bindings.Ref

//go:wasmimport ffi/js/array setNthNum
//go:noescape
func SetNthNum(
	self bindings.Ref,
	i uint32,
	val float64,
) bindings.Ref

//go:wasmimport ffi/js/array setNthBool
//go:noescape
func SetNthBool(
	self bindings.Ref,
	i uint32,
	val bindings.Ref,
) bindings.Ref

//go:wasmimport ffi/js/array setNthString
//go:noescape
func SetNthString(
	self bindings.Ref,
	i uint32,
	str unsafe.Pointer,
	len uint32,
) bindings.Ref
