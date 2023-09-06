// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/ffi/js/bindings"
)

type Func struct {
	ref bindings.Ref
}

func (fn Func) Ref() Ref {
	return Ref(fn.ref)
}

// Try runs the fn in a trycatch block, return heap reference of the result
// of either fn or catch.
//
// When free is true, all args after it (catch, finally, args...) are freed
// after the call.
//
// The bool return value indicates whether the try was successful (catch was
// not called)
func (fn Func) Try(this Ref, free bool, catch, finally Func, args ...Ref) (Ref, bool) {
	var caught bool
	ret := bindings.Try(
		fn.ref,
		uint32(this),
		uint32(Bool(free)),
		catch.ref,
		finally.ref,
		SliceData(args),
		SizeU(len(args)),
		Pointer(&caught),
	)

	return Ref(ret), !caught
}

// Call calls js function fn and returns the heap reference of the js return
// value.
//
// When free is ture, all args after it are freed after the call.
func (fn Func) Call(this Ref, free bool, args ...Ref) Ref {
	return Ref(bindings.Call(
		fn.ref,
		uint32(this),
		uint32(Bool(free)),
		SliceData(args),
		SizeU(len(args)),
	))
}

// CallVoid calls js function fn and discard any return value.
//
// When free is ture, all args after it are freed after the call.
func (fn Func) CallVoid(this Ref, free bool, args ...Ref) {
	bindings.CallVoid(
		fn.ref,
		uint32(this),
		uint32(Bool(free)),
		SliceData(args),
		SizeU(len(args)),
	)
}

// CallNum calls js function fn and treate the return value as a js number.
//
// When free is ture, all args after it are freed after the call.
func (fn Func) CallNum(this Ref, free bool, args ...Ref) Number {
	return bindings.CallNum(
		fn.ref,
		uint32(this),
		uint32(Bool(free)),
		SliceData(args),
		SizeU(len(args)),
	)
}

// CallBool calls js function fn and treate the return value as a js boolean.
//
// When free is ture, all args after it are freed after the call.
func (fn Func) CallBool(this Ref, free bool, args ...Ref) bool {
	return bindings.CallBool(
		fn.ref,
		uint32(this),
		uint32(Bool(free)),
		SliceData(args),
		SizeU(len(args)),
	) == bindings.Ref(True)
}

func (fn Func) Free() {
	bindings.Free(fn.ref)
}
