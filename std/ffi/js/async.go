// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/yield"
	"github.com/primecitizens/pcz/std/ffi/js/async"
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
)

// TODO: support batch awaiting (Promise.{allSettled, all, race})

// ExecutorFunc defines the go function signature of a promise executor.
type ExecutorFunc[T any] func(arg T, this Any) (value, reason Ref, fulfilled bool)

type jsExecutorFunc = func(resolve, reject Func[Ref])

// Executor is a helper for constructing new Promises in js.
//
// TODO: support handling multiple promises simultaneously.
type Executor[T any] struct {
	// executor is the executor function to be called from js.
	executor ExecutorFunc[T]
	arg      T

	fn      Func[jsExecutorFunc]
	promise bindings.Ref
}

func (p *Executor[T]) Free() {
	Free(p.fn.ref, p.promise)
}

// Reset prepares a new promise for use in js.
func (p *Executor[T]) Reset(arg T, fn ExecutorFunc[T]) Ref {
	p.arg = arg

	if p.executor == nil {
		bindings.Free(p.promise)
		p.executor = fn
	} else if abi.FuncPCABIInternal(p.executor) != abi.FuncPCABIInternal(fn) {
		p.Free()
		p.fn.ref = bindings.Ref(Undefined)
		p.executor = fn
	}

	p.promise = async.Promise(p.register().ref)
	return Ref(p.promise)
}

func (p *Executor[T]) register() Func[jsExecutorFunc] {
	if p == nil || p.executor == nil {
		assert.Throw("nil", "context", "or", "executor")
	}

	if p.fn.ref != bindings.Ref(Undefined) {
		return p.fn
	}

	p.fn = RegisterCallback[jsExecutorFunc](p, abi.FuncPCABIInternal(p.executor))
	return p.fn
}

// DispatchCallback implements CallbackDispatcher.
func (p *Executor[T]) DispatchCallback(targetPC uintptr, ctx *CallbackContext) {
	args := ctx.Args()
	if len(args) != 3 || targetPC != abi.FuncPCABIInternal(p.executor) {
		assert.Throw("invalid", "promise", "callback", "call")
	}

	value, reason, fulfilled := p.executor(
		p.arg,
		Any{
			ref: bindings.Ref(args[0]),
		},
	)

	if fulfilled {
		Func[func(Any)]{ref: bindings.Ref(args[1])}.CallVoid(args[0], true, value)
	} else {
		Func[func(Any)]{ref: bindings.Ref(args[2])}.CallVoid(args[0], true, reason)
	}
}

func NewResolvedPromise[
	T interface{ FromRef(Ref) T },
](data Ref) Promise[T] {
	return Promise[T]{
		ref: async.Resolved(bindings.Ref(data)),
	}
}

func NewRejectedPromise[
	T interface{ FromRef(Ref) T },
](reason Ref) Promise[T] {
	return Promise[T]{
		ref: async.Rejected(bindings.Ref(reason)),
	}
}

// TryAwait is a helper function to consume Promises returned
// from TryXxx functions.
//
// It calls p.Await(true) when ok is true.
func TryAwait[
	T interface{ FromRef(Ref) T },
](
	p Promise[T], exception Any, ok bool,
) (
	value T, err Any, fulfilled bool,
) {
	if !ok {
		err = exception
		return
	}

	return p.Await(true)
}

type Promise[T interface{ FromRef(Ref) T }] struct {
	ref bindings.Ref
}

func (p Promise[T]) FromRef(ref Ref) Promise[T] {
	return Promise[T]{ref: bindings.Ref(ref)}
}

func (p Promise[T]) Ref() Ref {
	return Ref(p.ref)
}

func (p Promise[T]) Once() Promise[T] {
	bindings.Once(p.ref)
	return p
}

func (p Promise[T]) Free() {
	bindings.Free(p.ref)
}

// Then calls Promise.then(onfulfilled, onrejected) and
// if onfinally is not undefined/null value, call Promise.finally(onfinally).
//
// NOTE: It is generally discouraged to use this function for callbacks to Go code.
func (p Promise[T]) Then(onfulfilled, onrejected Func[func(Ref)], onfinally Func[func()]) Promise[T] {
	async.Then(onfulfilled.ref, onrejected.ref, onfinally.ref)
	return p
}

// Await works like `await`
//
// When take is true, it frees the Promise p from the application heap.
//
//go:nosplit
func (p Promise[T]) Await(take bool) (value T, err Any, fulfilled bool) {
	var (
		ref Ref
		ok  bool
	)

	if bindings.Ref(True) == async.ShouldWait(
		bindings.Ref(Bool(take)), p.ref, Pointer(&ok), Pointer(&ref),
	) {
		yield.Thread()
		// notified from js as promise finished
		ok = bindings.Ref(True) == async.TakeWaited(bindings.Ref(ref), Pointer(&ref))
	}

	if ok {
		return value.FromRef(ref), err, true
	}

	return value, err.FromRef(ref), false
}

func (p Promise[T]) MustAwait(take bool) T {
	val, err, ok := p.Await(take)
	if !ok {
		// TODO: log error
		err.Free()
		assert.Throw("promise", "not", "fulfilled")
	}

	return val
}
