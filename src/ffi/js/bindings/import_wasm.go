// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

//
// Array
//

//go:wasmimport ffi/js array
func Array(cap uint32, elemSz uint32, signed, float Ref) Ref

//go:wasmimport ffi/js slice
func Slice(arr Ref, start, end uint32) Ref

//go:wasmimport ffi/js append
func Append(arr, free Ref, elemSz uint32, signed, float Ref, off uint32, ptr Uintptr32, n uint32) uint32

//go:wasmimport ffi/js copy
func Copy(arr Ref, elemSz uint32, signed, float Ref, off uint32, ptr Uintptr32, n uint32) uint32

//go:wasmimport ffi/js nth
func Nth(arr Ref, elemSz, signed, float uint32, i uint32, outRef Uintptr32) Ref

//go:wasmimport ffi/js nthNum
func NthNum(self Ref, i uint32, outNum Uintptr32) Ref

//go:wasmimport ffi/js nthBool
func NthBool(self Ref, i uint32, outBool Uintptr32) Ref

//go:wasmimport ffi/js setNth
func SetNth(self Ref, i int32, free, val Ref) Ref

//go:wasmimport ffi/js setNthNum
func SetNthNum(self Ref, i int32, val Number) Ref

//go:wasmimport ffi/js setNthBool
func SetNthBool(self Ref, i int32, val uint32) Ref

//go:wasmimport ffi/js setNthString
func SetNthString(self Ref, i int32, ptr Uintptr32, sz uint32) Ref

//
// Func
//

//go:wasmimport ffi/js func
func Func(dispFnPC, handler, targetPC uint32) Ref

//go:wasmimport ffi/js cbctx
func CallbackContext(ctx Ref, ptr Uintptr32)

//go:wasmimport ffi/js try
func Try(fn, this, free, catch, finally Ref, ptr Uintptr32, n uint32, pCaught Uintptr32) Ref

//go:wasmimport ffi/js call
func Call(fn, this, free Ref, ptr Uintptr32, n uint32) Ref

//go:wasmimport ffi/js callVoid
func CallVoid(fn, this, free Ref, ptr Uintptr32, n uint32)

//go:wasmimport ffi/js callNum
func CallNum(fn, this, free Ref, ptr Uintptr32, n uint32) Number

//go:wasmimport ffi/js callBool
func CallBool(fn, this, free Ref, ptr Uintptr32, n uint32) Ref

//
// Heap
//

//go:wasmimport ffi/js instanceof
func Instanceof(a, b Ref) Ref

//go:wasmimport ffi/js fill
func Fill(old, new Ref)

//go:wasmimport ffi/js fillNum
func FillNum(old Ref, n Number)

//go:wasmimport ffi/js fillBool
func FillBool(old Ref, b uint32)

//go:wasmimport ffi/js fillString
func FillString(old Ref, ptr Uintptr32, sz uint32)

//go:wasmimport ffi/js free
func Free(Ref)

//go:wasmimport ffi/js batchFree
func BatchFree(ptr Uintptr32, n uint32)

//
// Number
//

//go:wasmimport ffi/js number
func NewNumber(n float64) Ref

//
// Object
//

//go:wasmimport ffi/js new
func New(constructor Ref, ptr Uintptr32, n uint32) Ref

//go:wasmimport ffi/js getPrototype
func GetPrototype(o Ref) Ref

//go:wasmimport ffi/js setPrototype
func SetPrototype(self, free, proto Ref) Ref

//go:wasmimport ffi/js getPropDesc
func GetPropDesc(self Ref, ptr Uintptr32, sz uint32) uint32

//go:wasmimport ffi/js defineProp
func DefineProp(self Ref, ptr Uintptr32, sz uint32, flags uint32, getter, setter Ref) Ref

//go:wasmimport ffi/js deleteProp
func DeleteProp(self Ref, ptr Uintptr32, sz uint32) Ref

//go:wasmimport ffi/js prop
func Prop(this Ref, ptr Uintptr32, sz uint32) Ref

//go:wasmimport ffi/js numProp
func NumProp(this Ref, ptr Uintptr32, sz uint32) Number

//go:wasmimport ffi/js boolProp
func BoolProp(this Ref, ptr Uintptr32, sz uint32) Ref

//go:wasmimport ffi/js propByString
func PropByString(this, name Ref) Ref

//go:wasmimport ffi/js numPropByString
func NumPropByString(this, name Ref) Number

//go:wasmimport ffi/js boolPropByString
func BoolPropByString(this, name Ref) Ref

//go:wasmimport ffi/js setProp
func SetProp(this Ref, ptr Uintptr32, sz uint32, free Ref, val Ref) Ref

//go:wasmimport ffi/js setNumProp
func SetNumProp(this Ref, ptr Uintptr32, sz uint32, val Number) Ref

//go:wasmimport ffi/js setBoolProp
func SetBoolProp(this Ref, ptr Uintptr32, sz uint32, val Ref) Ref

//go:wasmimport ffi/js setStringProp
func SetStringProp(this Ref, nameptr Uintptr32, namesz uint32, valptr Uintptr32, valsz uint32) Ref

//go:wasmimport ffi/js setPropByString
func SetPropByString(this, prop, free, val Ref) Ref

//go:wasmimport ffi/js setNumPropByString
func SetNumPropByString(this, prop Ref, val Number) Ref

//go:wasmimport ffi/js setBoolPropByString
func SetBoolPropByString(this, prop, val Ref) Ref

//go:wasmimport ffi/js setStringPropByString
func SetStringPropByString(this, prop, free, val Ref) Ref

//
// Promise
//

//go:wasmimport ffi/js promise
func Promise(cb Ref) Ref

//go:wasmimport ffi/js then
func Then(onfulfilled, onrejected Ref)

//go:wasmimport ffi/js catch
func Catch(onrejected Ref)

//go:wasmimport ffi/js finally
func Finally(onfinally Ref)

//
// String
//

//go:wasmimport ffi/js stringFromUTF8
func StringFromUTF8(ptr Uintptr32, sz uint32) Ref

//go:wasmimport ffi/js stringEqualsUTF8
func StringEqualsUTF8(self Ref, ptr Uintptr32, sz uint32) Ref

//go:wasmimport ffi/js stringIntoUTF8
func StringIntoUTF8(s Ref, ptr Uintptr32, len uint32) uint32

//go:wasmimport ffi/js stringSizeUTF8
func StringSizeUTF8(s Ref) Number
