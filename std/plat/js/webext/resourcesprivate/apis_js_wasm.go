// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package resourcesprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/resourcesprivate/bindings"
)

type Component uint32

const (
	_ Component = iota

	Component_IDENTITY
	Component_PDF
)

func (Component) FromRef(str js.Ref) Component {
	return Component(bindings.ConstOfComponent(str))
}

func (x Component) String() (string, bool) {
	switch x {
	case Component_IDENTITY:
		return "identity", true
	case Component_PDF:
		return "pdf", true
	default:
		return "", false
	}
}

type GetStringsCallbackFunc func(this js.Ref, result js.Object) js.Ref

func (fn GetStringsCallbackFunc) Register() js.Func[func(result js.Object)] {
	return js.RegisterCallback[func(result js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetStringsCallbackFunc) DispatchCallback(
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

type GetStringsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Object) js.Ref
	Arg T
}

func (cb *GetStringsCallback[T]) Register() js.Func[func(result js.Object)] {
	return js.RegisterCallback[func(result js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetStringsCallback[T]) DispatchCallback(
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

// HasFuncGetStrings returns true if the function "WEBEXT.resourcesPrivate.getStrings" exists.
func HasFuncGetStrings() bool {
	return js.True == bindings.HasFuncGetStrings()
}

// FuncGetStrings returns the function "WEBEXT.resourcesPrivate.getStrings".
func FuncGetStrings() (fn js.Func[func(component Component) js.Promise[js.Object]]) {
	bindings.FuncGetStrings(
		js.Pointer(&fn),
	)
	return
}

// GetStrings calls the function "WEBEXT.resourcesPrivate.getStrings" directly.
func GetStrings(component Component) (ret js.Promise[js.Object]) {
	bindings.CallGetStrings(
		js.Pointer(&ret),
		uint32(component),
	)

	return
}

// TryGetStrings calls the function "WEBEXT.resourcesPrivate.getStrings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetStrings(component Component) (ret js.Promise[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetStrings(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(component),
	)

	return
}
