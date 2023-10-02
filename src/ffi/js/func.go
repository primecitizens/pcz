// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js/bindings"
)

func ThrowInvalidCallbackInvocation() {
	assert.Throw("invalid", "callback", "invocation")
}

func ThrowCallbackValueNotReturned() {
	assert.Throw(
		"value", "not", "returned:", "internal", "callback", "bindings", "error",
	)
}

type Func[T any] struct {
	ref bindings.Ref
}

func (fn Func[T]) FromRef(ref Ref) Func[T] {
	return Func[T]{
		ref: bindings.Ref(ref),
	}
}

func (fn Func[T]) Ref() Ref {
	return Ref(fn.ref)
}

func (fn Func[T]) Once() Func[T] {
	Ref(fn.ref).Once()
	return fn
}

func (fn Func[T]) Free() {
	bindings.Free(fn.ref)
}

// Try runs the fn in a trycatch block, return heap reference of the result
// of either fn or catch.
//
// When free is true, all args after it (catch, finally, args...) are freed
// after the call.
//
// The bool return value indicates whether the try was successful (catch was
// not called)
func (fn Func[T]) Try(
	this Ref,
	take bool,
	catch Func[func(err Any, args []Ref) Ref],
	finally Func[func(args []Ref)],
	args ...Ref,
) (Ref, bool) {
	var caught bool
	ret := bindings.Try(
		fn.ref,
		bindings.Ref(this),
		bindings.Ref(Bool(take)),
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
func (fn Func[T]) Call(this Ref, take bool, args ...Ref) Ref {
	return Ref(bindings.Call(
		fn.ref,
		bindings.Ref(this),
		bindings.Ref(Bool(take)),
		SliceData(args),
		SizeU(len(args)),
	))
}

// CallVoid calls js function fn and discard any return value.
//
// When free is ture, all args after it are freed after the call.
func (fn Func[T]) CallVoid(this Ref, take bool, args ...Ref) {
	bindings.CallVoid(
		fn.ref,
		bindings.Ref(this),
		bindings.Ref(Bool(take)),
		SliceData(args),
		SizeU(len(args)),
	)
}

// CallNum calls js function fn and treate the return value as a js number.
//
// When free is ture, all args after it are freed after the call.
func (fn Func[T]) CallNum(this Ref, take bool, args ...Ref) float64 {
	return bindings.CallNum(
		fn.ref,
		bindings.Ref(this),
		bindings.Ref(Bool(take)),
		SliceData(args),
		SizeU(len(args)),
	)
}

// CallBool calls js function fn and treate the return value as a js boolean.
//
// When free is ture, all args after it are freed after the call.
func (fn Func[T]) CallBool(this Ref, take bool, args ...Ref) bool {
	return bindings.Ref(True) == bindings.CallBool(
		fn.ref,
		bindings.Ref(this),
		bindings.Ref(Bool(take)),
		SliceData(args),
		SizeU(len(args)),
	)
}
