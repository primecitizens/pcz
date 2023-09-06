// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js/bindings"
)

// PromiseFunc is a simple wrapper for your function.
type PromiseFunc func(resolve, reject Func)

// HandleJSCallback implements CallbackHandler.
func (fn PromiseFunc) DispatchCallback(targetPC uintptr, ctx *CallbackContext) {
	args := ctx.Args()
	if len(args) != 2 || targetPC != abi.FuncPCABIInternal(fn) {
		assert.Throw("invalid", "promise", "callback", "call")
	}

	fn(args[0].AsFunc(), args[1].AsFunc())
}

// Register registers self to the js for callback operation.
func (fn PromiseFunc) Register() Func {
	return RegisterCallback(fn, abi.FuncPCABIInternal(fn))
}

func NewPromise[T any](cb Func) Promise[T] {
	return Promise[T]{
		bindings.Promise(cb.ref),
	}
}

type Promise[T any] struct {
	ref bindings.Ref
}

func (p Promise[T]) Ref() Ref {
	return Ref(p.ref)
}

func (p Promise[T]) Free() {
	bindings.Free(p.ref)
}

func (p Promise[T]) FromRef(ref Ref) Promise[T] {
	return Promise[T]{ref: bindings.Ref(ref)}
}

func (p Promise[T]) Then(onfulfilled, onrejected Func) Promise[T] {
	bindings.Then(onfulfilled.ref, onrejected.ref)
	return p
}

func (p Promise[T]) Catch(onrejected Func) Promise[T] {
	bindings.Catch(onrejected.ref)
	return p
}

func (p Promise[T]) Finally(onfinally Func) Promise[T] {
	bindings.Finally(onfinally.ref)
	return p
}
