// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

const smallIntCacheSize = 128

const (
	Undefined   Ref = iota // undefined
	Null                   // null
	Zero                   // 0
	False                  // false
	True                   // true
	NaN                    // NaN
	Inf                    // Infinity
	NegativeInf            // -Infinity
	Global                 // globalThis

	firstSmallIntCache                                          // = 1
	lastSmallIntCache  = firstSmallIntCache + smallIntCacheSize // = 128
)

type Ref bindings.Ref

func (r Ref) Undefined() bool { return r == Undefined }

// Falsy returns true when r is the reference to one of following values:
//   - Undefined
//   - Null
//   - 0
//   - False
//
// Otherwise, it returns false.
func (r Ref) Falsy() bool { return r < True }

// Truthy returns true when r.Falsy() returns false.
func (r Ref) Truthy() bool { return r > False }

func (r Ref) Once() Ref {
	bindings.Once(bindings.Ref(r))
	return r
}

// Clone returns a new heap reference to the same object referenced by r.
func (r Ref) Clone() Ref {
	return Ref(bindings.Clone(bindings.Ref(r)))
}

// Free the referenced js value.
func (r Ref) Free() {
	bindings.Free(bindings.Ref(r))
}

func (r Ref) Any() Any {
	return Any{
		ref: bindings.Ref(r),
	}
}

func Free[T ~uint32](refs ...T) {
	if len(refs) == 0 {
		return
	}

	bindings.BatchFree(
		SliceData(refs),
		SizeU(len(refs)),
	)
}

type Any struct {
	ref bindings.Ref
}

func (x Any) FromRef(ref Ref) Any {
	return Any{ref: bindings.Ref(ref)}
}

func (x Any) Ref() Ref {
	return Ref(x.ref)
}

func (x Any) Once() Any {
	bindings.Once(x.ref)
	return x
}

func (x Any) Free() {
	bindings.Free(x.ref)
}

func (x Any) AsNumber() Number[float64] {
	return Number[float64]{ref: x.ref}
}

func (x Any) AsFunc() Func[Any] {
	return Func[Any]{ref: x.ref}
}

func (x Any) AsArray() Array[Any] {
	return Array[Any]{ref: x.ref}
}

func (x Any) AsString() String {
	return String{ref: x.ref}
}

func (x Any) AsObject() Object {
	return Object{ref: x.ref}
}

func (x Any) InstanceOf(o Ref) bool {
	return bindings.Ref(True) == bindings.Instanceof(
		x.ref,
		bindings.Ref(o),
	)
}

type Void struct{}

func (Void) FromRef(Ref) Void { return Void{} }
func (Void) Ref() Ref         { return Undefined }
func (Void) Once() Void       { return Void{} }
func (Void) Free()            {}

type Boolean struct {
	ref Ref
}

func (b Boolean) FromRef(ref Ref) Boolean {
	if ref != True && ref != False {
		assert.Throw("not", "a", "boolean", "ref")
	}

	return Boolean{
		ref: ref,
	}
}

func (b Boolean) Ref() Ref {
	return b.ref
}

func (b Boolean) Once() Boolean {
	return b
}

func (b Boolean) Free() {}

func (b Boolean) Get() bool {
	return b.ref == True
}
