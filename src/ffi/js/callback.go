// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/core/math"
	"github.com/primecitizens/std/ffi"
	"github.com/primecitizens/std/ffi/js/bindings"
)

type CallbackDispatcher ffi.CallbackDispatcher[CallbackContext]

type DummyDispatcher func()

func (fn DummyDispatcher) DispatchCallback(_ uintptr, ctx *CallbackContext) {
	fn()
}

func (fn DummyDispatcher) Register() Func {
	return RegisterCallback(fn, abi.FuncPCABIInternal(fn))
}

// dispFunc for type checking
type dispFunc[T any] ffi.DispatchFunc[T, CallbackContext]

// RegisterCallback
//
// NOTE: type param H MUST be a concrete pointer type, that is, it either
//   - comes with star sign (e.g. *Foo)
//   - or being a function (e.g. [PromiseCallbackHandleFunc]).
func RegisterCallback[D CallbackDispatcher](dispatcher D, targetPC uintptr) Func {
	if unsafe.Sizeof(unsafe.Pointer(nil)) != unsafe.Sizeof(uintptr(0)) {
		assert.Throw("unexpected", "pointer", "size", "not", "match")
	}

	if unsafe.Sizeof(dispatcher) != unsafe.Sizeof((*byte)(nil)) {
		assert.Throw("invalid", "callback", "handler:", "not", "a", "pointer", "value")
	}

	var _ dispFunc[D] = D.DispatchCallback // type check
	return Func{
		ref: bindings.Func(
			SizeU(abi.FuncPCABIInternal(D.DispatchCallback)),
			// this is still the `handler` as passed in
			// we just converted it to pointer value as assumed
			SizeU(uintptr(
				unsafe.Pointer(
					*(**byte)(
						unsafe.Pointer(mark.NoEscape(&dispatcher)),
					),
				),
			)),
			SizeU(targetPC),
		),
	}
}

type CallbackContext struct {
	dispFnPC uint32
	handler  uint32
	targetPC uint32
	retRef   bindings.Ref
	nargs    uint32
	args     [16]bindings.Ref
}

func (ctx *CallbackContext) Args() []Ref {
	return unsafe.Slice((*Ref)(mark.NoEscape(&ctx.args[0])), ctx.nargs)
}

func (ctx *CallbackContext) Return(ref Ref) {
	bindings.Fill(ctx.retRef, bindings.Ref(ref))
}

func (ctx *CallbackContext) ReturnNum(n Number) {
	bindings.FillNum(ctx.retRef, float64(n))
}

func (ctx *CallbackContext) ReturnBool(b bool) {
	if b {
		bindings.FillBool(ctx.retRef, 1)
	} else {
		bindings.FillBool(ctx.retRef, 0)
	}
}

func (ctx *CallbackContext) ReturnString(s string) {
	bindings.FillString(
		ctx.retRef,
		StringData(s),
		SizeU(len(s)),
	)
}

// handleCallback called from wasm_export_run.
func handleCallback(ref bindings.Ref, unused uint32) {
	const validOffsets = true &&
		unsafe.Offsetof(CallbackContext{}.dispFnPC) == 0 &&
		unsafe.Offsetof(CallbackContext{}.handler) == 4 &&
		unsafe.Offsetof(CallbackContext{}.targetPC) == 8 &&
		unsafe.Offsetof(CallbackContext{}.retRef) == 12 &&
		unsafe.Offsetof(CallbackContext{}.nargs) == 16 &&
		unsafe.Offsetof(CallbackContext{}.args) == 20
	if !validOffsets {
		assert.Throw("unexpected", "invalid", "callback", "context", "field", "offsets")
	}

	var cb CallbackContext

	// defensive, to ensure js side sets nargs.
	cb.nargs = math.MaxUint32

	bindings.CallbackContext(ref, Pointer(&cb))
	if cb.nargs > 16 {
		assert.Throw("too", "many", "args", "to", "callback", "func")
	}

	pc := uintptr(cb.dispFnPC)
	callHandler(
		uintptr(cb.handler),
		uintptr(cb.targetPC),
		mark.NoEscape(&cb),
		(*(*dispFunc[uintptr])(unsafe.Pointer(mark.NoEscape(&pc)))),
	)
}
