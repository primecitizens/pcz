// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package idltest

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/idltest/bindings"
)

type ArrayBufferCallbackFunc func(this js.Ref, buffer js.ArrayBuffer) js.Ref

func (fn ArrayBufferCallbackFunc) Register() js.Func[func(buffer js.ArrayBuffer)] {
	return js.RegisterCallback[func(buffer js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ArrayBufferCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ArrayBufferCallback[T any] struct {
	Fn  func(arg T, this js.Ref, buffer js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *ArrayBufferCallback[T]) Register() js.Func[func(buffer js.ArrayBuffer)] {
	return js.RegisterCallback[func(buffer js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ArrayBufferCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LongArrayCallbackFunc func(this js.Ref, array js.Array[int32]) js.Ref

func (fn LongArrayCallbackFunc) Register() js.Func[func(array js.Array[int32])] {
	return js.RegisterCallback[func(array js.Array[int32])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LongArrayCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LongArrayCallback[T any] struct {
	Fn  func(arg T, this js.Ref, array js.Array[int32]) js.Ref
	Arg T
}

func (cb *LongArrayCallback[T]) Register() js.Func[func(array js.Array[int32])] {
	return js.RegisterCallback[func(array js.Array[int32])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LongArrayCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[int32]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetArrayBuffer returns true if the function "WEBEXT.idltest.getArrayBuffer" exists.
func HasFuncGetArrayBuffer() bool {
	return js.True == bindings.HasFuncGetArrayBuffer()
}

// FuncGetArrayBuffer returns the function "WEBEXT.idltest.getArrayBuffer".
func FuncGetArrayBuffer() (fn js.Func[func() js.Promise[js.ArrayBuffer]]) {
	bindings.FuncGetArrayBuffer(
		js.Pointer(&fn),
	)
	return
}

// GetArrayBuffer calls the function "WEBEXT.idltest.getArrayBuffer" directly.
func GetArrayBuffer() (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallGetArrayBuffer(
		js.Pointer(&ret),
	)

	return
}

// TryGetArrayBuffer calls the function "WEBEXT.idltest.getArrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArrayBuffer() (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArrayBuffer(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncNocompileFunc returns true if the function "WEBEXT.idltest.nocompileFunc" exists.
func HasFuncNocompileFunc() bool {
	return js.True == bindings.HasFuncNocompileFunc()
}

// FuncNocompileFunc returns the function "WEBEXT.idltest.nocompileFunc".
func FuncNocompileFunc() (fn js.Func[func(switch_ int32)]) {
	bindings.FuncNocompileFunc(
		js.Pointer(&fn),
	)
	return
}

// NocompileFunc calls the function "WEBEXT.idltest.nocompileFunc" directly.
func NocompileFunc(switch_ int32) (ret js.Void) {
	bindings.CallNocompileFunc(
		js.Pointer(&ret),
		int32(switch_),
	)

	return
}

// TryNocompileFunc calls the function "WEBEXT.idltest.nocompileFunc"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNocompileFunc(switch_ int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNocompileFunc(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(switch_),
	)

	return
}

// HasFuncSendArrayBuffer returns true if the function "WEBEXT.idltest.sendArrayBuffer" exists.
func HasFuncSendArrayBuffer() bool {
	return js.True == bindings.HasFuncSendArrayBuffer()
}

// FuncSendArrayBuffer returns the function "WEBEXT.idltest.sendArrayBuffer".
func FuncSendArrayBuffer() (fn js.Func[func(input js.ArrayBuffer) js.Promise[js.Array[int32]]]) {
	bindings.FuncSendArrayBuffer(
		js.Pointer(&fn),
	)
	return
}

// SendArrayBuffer calls the function "WEBEXT.idltest.sendArrayBuffer" directly.
func SendArrayBuffer(input js.ArrayBuffer) (ret js.Promise[js.Array[int32]]) {
	bindings.CallSendArrayBuffer(
		js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySendArrayBuffer calls the function "WEBEXT.idltest.sendArrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendArrayBuffer(input js.ArrayBuffer) (ret js.Promise[js.Array[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendArrayBuffer(
		js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSendArrayBufferView returns true if the function "WEBEXT.idltest.sendArrayBufferView" exists.
func HasFuncSendArrayBufferView() bool {
	return js.True == bindings.HasFuncSendArrayBufferView()
}

// FuncSendArrayBufferView returns the function "WEBEXT.idltest.sendArrayBufferView".
func FuncSendArrayBufferView() (fn js.Func[func(input js.ArrayBufferView) js.Promise[js.Array[int32]]]) {
	bindings.FuncSendArrayBufferView(
		js.Pointer(&fn),
	)
	return
}

// SendArrayBufferView calls the function "WEBEXT.idltest.sendArrayBufferView" directly.
func SendArrayBufferView(input js.ArrayBufferView) (ret js.Promise[js.Array[int32]]) {
	bindings.CallSendArrayBufferView(
		js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySendArrayBufferView calls the function "WEBEXT.idltest.sendArrayBufferView"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendArrayBufferView(input js.ArrayBufferView) (ret js.Promise[js.Array[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendArrayBufferView(
		js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}
