// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sharedstorageprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sharedstorageprivate/bindings"
)

type GetCallbackFunc func(this js.Ref, items js.Object) js.Ref

func (fn GetCallbackFunc) Register() js.Func[func(items js.Object)] {
	return js.RegisterCallback[func(items js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCallback[T any] struct {
	Fn  func(arg T, this js.Ref, items js.Object) js.Ref
	Arg T
}

func (cb *GetCallback[T]) Register() js.Func[func(items js.Object)] {
	return js.RegisterCallback[func(items js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemoveCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemoveCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetCallbackFunc func(this js.Ref) js.Ref

func (fn SetCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGet returns true if the function "WEBEXT.sharedStoragePrivate.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.sharedStoragePrivate.get".
func FuncGet() (fn js.Func[func() js.Promise[js.Object]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.sharedStoragePrivate.get" directly.
func Get() (ret js.Promise[js.Object]) {
	bindings.CallGet(
		js.Pointer(&ret),
	)

	return
}

// TryGet calls the function "WEBEXT.sharedStoragePrivate.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet() (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemove returns true if the function "WEBEXT.sharedStoragePrivate.remove" exists.
func HasFuncRemove() bool {
	return js.True == bindings.HasFuncRemove()
}

// FuncRemove returns the function "WEBEXT.sharedStoragePrivate.remove".
func FuncRemove() (fn js.Func[func(keys js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncRemove(
		js.Pointer(&fn),
	)
	return
}

// Remove calls the function "WEBEXT.sharedStoragePrivate.remove" directly.
func Remove(keys js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallRemove(
		js.Pointer(&ret),
		keys.Ref(),
	)

	return
}

// TryRemove calls the function "WEBEXT.sharedStoragePrivate.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemove(keys js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemove(
		js.Pointer(&ret), js.Pointer(&exception),
		keys.Ref(),
	)

	return
}

// HasFuncSet returns true if the function "WEBEXT.sharedStoragePrivate.set" exists.
func HasFuncSet() bool {
	return js.True == bindings.HasFuncSet()
}

// FuncSet returns the function "WEBEXT.sharedStoragePrivate.set".
func FuncSet() (fn js.Func[func(items js.Object) js.Promise[js.Void]]) {
	bindings.FuncSet(
		js.Pointer(&fn),
	)
	return
}

// Set calls the function "WEBEXT.sharedStoragePrivate.set" directly.
func Set(items js.Object) (ret js.Promise[js.Void]) {
	bindings.CallSet(
		js.Pointer(&ret),
		items.Ref(),
	)

	return
}

// TrySet calls the function "WEBEXT.sharedStoragePrivate.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySet(items js.Object) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySet(
		js.Pointer(&ret), js.Pointer(&exception),
		items.Ref(),
	)

	return
}
