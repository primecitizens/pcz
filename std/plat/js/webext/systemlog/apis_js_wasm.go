// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package systemlog

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/systemlog/bindings"
)

type MessageOptions struct {
	// Message is "MessageOptions.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MessageOptions with all fields set.
func (p MessageOptions) FromRef(ref js.Ref) MessageOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MessageOptions in the application heap.
func (p MessageOptions) New() js.Ref {
	return bindings.MessageOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MessageOptions) UpdateFrom(ref js.Ref) {
	bindings.MessageOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MessageOptions) Update(ref js.Ref) {
	bindings.MessageOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MessageOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncAdd returns true if the function "WEBEXT.systemLog.add" exists.
func HasFuncAdd() bool {
	return js.True == bindings.HasFuncAdd()
}

// FuncAdd returns the function "WEBEXT.systemLog.add".
func FuncAdd() (fn js.Func[func(options MessageOptions) js.Promise[js.Void]]) {
	bindings.FuncAdd(
		js.Pointer(&fn),
	)
	return
}

// Add calls the function "WEBEXT.systemLog.add" directly.
func Add(options MessageOptions) (ret js.Promise[js.Void]) {
	bindings.CallAdd(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAdd calls the function "WEBEXT.systemLog.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAdd(options MessageOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAdd(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}
