// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !wasm

package bindings

//
// Array
//

func Array(cap uint32, elemSz uint32, signed, float Ref) Ref { return 0 }
func Slice(arr Ref, start, end uint32) Ref                   { return 0 }
func Append(arr, free Ref, elemSz uint32, signed, float Ref, off uint32, ptr Uintptr32, n uint32) uint32 {
	return 0
}
func Copy(arr Ref, elemSz uint32, signed, float Ref, off uint32, ptr Uintptr32, n uint32) uint32 {
	return 0
}
func Nth(arr Ref, elemSz, signed, float uint32, i uint32, outRef Uintptr32) Ref { return 0 }
func NthNum(self Ref, i uint32, outNum Uintptr32) Ref                           { return 0 }
func NthBool(self Ref, i uint32, outBool Uintptr32) Ref                         { return 0 }
func SetNth(self Ref, i int32, free, val Ref) Ref                               { return 0 }
func SetNthNum(self Ref, i int32, val Number) Ref                               { return 0 }
func SetNthBool(self Ref, i int32, val uint32) Ref                              { return 0 }
func SetNthString(self Ref, i int32, ptr Uintptr32, sz uint32) Ref              { return 0 }

//
// Func
//

func Func(dispFnPC, handler, targetPC uint32) Ref { return 0 }
func CallbackContext(ctx Ref, ptr Uintptr32)      {}

func Try(fn, this, free, catch, finally Ref, ptr Uintptr32, n uint32, pCaught Uintptr32) Ref {
	return 0
}
func Call(fn, this, free Ref, ptr Uintptr32, n uint32) Ref { return 0 }
func CallVoid(fn, this, free Ref, ptr Uintptr32, n uint32) {}

func CallNum(fn, this, free Ref, ptr Uintptr32, n uint32) Number { return 0 }
func CallBool(fn, this, free Ref, ptr Uintptr32, n uint32) Ref   { return 0 }

//
// Heap
//

func Instanceof(a, b Ref) Ref                      { return 0 }
func Fill(old, new Ref)                            {}
func FillNum(old Ref, n Number)                    {}
func FillBool(old Ref, b uint32)                   {}
func FillString(old Ref, ptr Uintptr32, sz uint32) {}
func Free(Ref)                                     {}
func BatchFree(ptr Uintptr32, n uint32)            {}

// Number

func NewNumber(n float64) Ref { return 0 }

// Object

func New(constructor Ref, ptr Uintptr32, n uint32) Ref                                    { return 0 }
func GetPrototype(o Ref) Ref                                                              { return 0 }
func SetPrototype(self, free, proto Ref) Ref                                              { return 0 }
func GetPropDesc(self Ref, ptr Uintptr32, sz uint32) uint32                               { return 0 }
func DefineProp(self Ref, ptr Uintptr32, sz uint32, flags uint32, getter, setter Ref) Ref { return 0 }
func DeleteProp(self Ref, ptr Uintptr32, sz uint32) Ref                                   { return 0 }
func Prop(this Ref, ptr Uintptr32, sz uint32) Ref                                         { return 0 }
func NumProp(this Ref, ptr Uintptr32, sz uint32) Number                                   { return 0 }
func BoolProp(this Ref, ptr Uintptr32, sz uint32) Ref                                     { return 0 }
func PropByString(this, name Ref) Ref                                                     { return 0 }
func NumPropByString(this, name Ref) Number                                               { return 0 }
func BoolPropByString(this, name Ref) Ref                                                 { return 0 }
func SetProp(this Ref, ptr Uintptr32, sz uint32, free Ref, val Ref) Ref                   { return 0 }
func SetNumProp(this Ref, ptr Uintptr32, sz uint32, val Number) Ref                       { return 0 }
func SetBoolProp(this Ref, ptr Uintptr32, sz uint32, val Ref) Ref                         { return 0 }
func SetStringProp(this Ref, nameptr Uintptr32, namesz uint32, valptr Uintptr32, valsz uint32) Ref {
	return 0
}
func SetPropByString(this, prop, free, val Ref) Ref       { return 0 }
func SetNumPropByString(this, prop Ref, val Number) Ref   { return 0 }
func SetBoolPropByString(this, prop, val Ref) Ref         { return 0 }
func SetStringPropByString(this, prop, free, val Ref) Ref { return 0 }

// Promise

func Promise(cb Ref) Ref               { return 0 }
func Then(onfulfilled, onrejected Ref) {}
func Catch(onrejected Ref)             {}
func Finally(onfinally Ref)            {}

// String

func StringFromUTF8(ptr Uintptr32, sz uint32) Ref             { return 0 }
func StringEqualsUTF8(self Ref, ptr Uintptr32, sz uint32) Ref { return 0 }
func StringIntoUTF8(s Ref, ptr Uintptr32, len uint32) uint32  { return 0 }
func StringSizeUTF8(s Ref) Number                             { return 0 }
