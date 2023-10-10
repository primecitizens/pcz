// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"
)

//
// Func
//

//go:wasmimport ffi/js try
//go:noescape
func Try(
	fn Ref,
	this Ref,
	take bool,
	catch Ref,
	finally Ref,
	pArgs unsafe.Pointer,
	count uint32,
	pCaught unsafe.Pointer,
) Ref

//go:wasmimport ffi/js call
//go:noescape
func Call(
	fn Ref,
	this Ref,
	take Ref,
	pArgs unsafe.Pointer,
	count uint32,
) Ref

//go:wasmimport ffi/js callVoid
//go:noescape
func CallVoid(fn, this, take Ref, pArgs unsafe.Pointer, count uint32)

//go:wasmimport ffi/js callNum
//go:noescape
func CallNum(fn, this, take Ref, pArgs unsafe.Pointer, count uint32) float64

//go:wasmimport ffi/js callBool
//go:noescape
func CallBool(fn, this, take Ref, pArgs unsafe.Pointer, count uint32) bool

//
// Heap
//

//go:wasmimport ffi/js once
//go:noescape
func Once(ref Ref)

//go:wasmimport ffi/js clone
//go:noescape
func Clone(ref Ref) Ref

//go:wasmimport ffi/js instanceof
//go:noescape
func Instanceof(a, b Ref) Ref

//go:wasmimport ffi/js replace
//go:noescape
func Replace(old, new Ref) Ref

//go:wasmimport ffi/js replaceNum
//go:noescape
func ReplaceNum(old Ref, n float64) Ref

//go:wasmimport ffi/js replaceBigInt
//go:noescape
func ReplaceBigInt(old Ref, signed Ref, new unsafe.Pointer) Ref

//go:wasmimport ffi/js replaceBool
//go:noescape
func ReplaceBool(old Ref, boolean Ref) Ref

//go:wasmimport ffi/js replaceString
//go:noescape
func ReplaceString(old Ref, str unsafe.Pointer, len uint32) Ref

//go:wasmimport ffi/js free
//go:noescape
func Free(Ref)

//go:wasmimport ffi/js batchFree
//go:noescape
func BatchFree(pRefs unsafe.Pointer, count uint32)

//
// Number
//

//go:wasmimport ffi/js number
//go:noescape
func Number(n float64) Ref

//go:wasmimport ffi/js getnum
//go:noescape
func GetNum(ref Ref, pF64 unsafe.Pointer) Ref

//go:wasmimport ffi/js bigint
//go:noescape
func BigInt(signed Ref, ptr unsafe.Pointer) Ref

//go:wasmimport ffi/js getBigInt
//go:noescape
func GetBigInt(ref Ref, signed Ref, ptr unsafe.Pointer) Ref

//
// Object
//

//go:wasmimport ffi/js new
//go:noescape
func New(constructor Ref, pArgs unsafe.Pointer, count uint32) Ref

//go:wasmimport ffi/js getPrototype
//go:noescape
func GetPrototype(obj Ref) Ref

//go:wasmimport ffi/js setPrototype
//go:noescape
func SetPrototype(self Ref, take bool, proto Ref) Ref

//go:wasmimport ffi/js getPropDesc
//go:noescape
func GetPropDesc(self Ref, pName unsafe.Pointer, lenName uint32) uint32

//go:wasmimport ffi/js defineProp
//go:noescape
func DefineProp(self Ref, pName unsafe.Pointer, lenName uint32, flags uint32, getter, setter Ref) bool

//go:wasmimport ffi/js deleteProp
//go:noescape
func DeleteProp(self Ref, pName unsafe.Pointer, sz uint32) bool

//go:wasmimport ffi/js prop
//go:noescape
func Prop(this Ref, pName unsafe.Pointer, lenName uint32) Ref

//go:wasmimport ffi/js numProp
//go:noescape
func NumProp(this Ref, pName unsafe.Pointer, lenName uint32) float64

//go:wasmimport ffi/js boolProp
//go:noescape
func BoolProp(this Ref, pName unsafe.Pointer, lenName uint32) bool

//go:wasmimport ffi/js propByString
//go:noescape
func PropByString(this, name Ref) Ref

//go:wasmimport ffi/js numPropByString
//go:noescape
func NumPropByString(this, name Ref) float64

//go:wasmimport ffi/js boolPropByString
//go:noescape
func BoolPropByString(this, name Ref) bool

//go:wasmimport ffi/js setProp
//go:noescape
func SetProp(
	this Ref,
	pName unsafe.Pointer,
	lenName uint32,
	take bool,
	val Ref,
) bool

//go:wasmimport ffi/js setNumProp
//go:noescape
func SetNumProp(
	this Ref,
	pName unsafe.Pointer,
	lenName uint32,
	val float64,
) bool

//go:wasmimport ffi/js setBoolProp
//go:noescape
func SetBoolProp(
	this Ref,
	pName unsafe.Pointer,
	lenName uint32,
	val bool,
) bool

//go:wasmimport ffi/js setStringProp
//go:noescape
func SetStringProp(
	this Ref,
	nameptr unsafe.Pointer,
	namesz uint32,
	valptr unsafe.Pointer,
	valsz uint32,
) bool

//go:wasmimport ffi/js setPropByString
//go:noescape
func SetPropByString(this, prop, take, val Ref) bool

//go:wasmimport ffi/js setNumPropByString
//go:noescape
func SetNumPropByString(this, prop Ref, val float64) bool

//go:wasmimport ffi/js setBoolPropByString
//go:noescape
func SetBoolPropByString(this, prop, val Ref) bool

//go:wasmimport ffi/js setStringPropByString
//go:noescape
func SetStringPropByString(this, prop, take, val Ref) bool

//
// String
//

//go:wasmimport ffi/js decodeUTF8
//go:noescape
func DecodeUTF8(ptr unsafe.Pointer, len uint32) Ref

//go:wasmimport ffi/js equalsUTF8
//go:noescape
func EqualsUTF8(self Ref, str unsafe.Pointer, len uint32) bool

//go:wasmimport ffi/js prependUTF8
//go:noescape
func PrependUTF8(self Ref, pStrSlice unsafe.Pointer, count uint32)

//go:wasmimport ffi/js prependString
//go:noescape
func PrependString(self Ref, take bool, pRefSlice unsafe.Pointer, count uint32)

//go:wasmimport ffi/js appendUTF8
//go:noescape
func AppendUTF8(self Ref, pStrSlice unsafe.Pointer, count uint32)

//go:wasmimport ffi/js appendString
//go:noescape
func AppendString(self Ref, take bool, pRefSlice unsafe.Pointer, count uint32)

//go:wasmimport ffi/js encodeUTF8
//go:noescape
func EncodeUTF8(s Ref, str unsafe.Pointer, len uint32) uint32

//go:wasmimport ffi/js sizeUTF8
//go:noescape
func SizeUTF8(s Ref) uint32
