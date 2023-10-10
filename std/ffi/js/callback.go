// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"unsafe"

	stdconst "github.com/primecitizens/pcz/std/builtin/const"
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/core/math"
	"github.com/primecitizens/pcz/std/ffi"
	"github.com/primecitizens/pcz/std/ffi/js/bindings"
	"github.com/primecitizens/pcz/std/ffi/js/callback"
)

// CallbackDispatcher defines the interface of a js callback dispatcher.
type CallbackDispatcher interface {
	ffi.CallbackDispatcher[CallbackContext]
}

// UntypedCallback dispatches js callbacks without resolving argument types.
//
// By convension, args[0] refers to the js `this` context.
type UntypedCallback func(args []Ref)

// Register registers this function to js and return the reference to that js function.
func (fn UntypedCallback) Register() Func[func([]Ref)] {
	return RegisterCallback[func([]Ref)](fn, abi.FuncPCABIInternal(fn))
}

// DispatchCallback implements [CallbackDispatcher].
func (fn UntypedCallback) DispatchCallback(_ uintptr, ctx *CallbackContext) {
	fn(ctx.Args())
}

// dispFunc for type checking
type dispFunc[T any] ffi.DispatchFunc[T, CallbackContext]

// RegisterCallback creates a js function to be called from js.
//
// See [CallbackContext] for more details.
//
// NOTE: type param D MUST be a pointer type, that is, it either
//   - comes with star sign (e.g. *Foo)
//   - or being a function (e.g. PromiseCallbackHandleFunc).
func RegisterCallback[F any, D CallbackDispatcher](dispatcher D, targetPC uintptr) Func[F] {
	if unsafe.Sizeof(dispatcher) != stdconst.SizePointer {
		assert.Throw("invalid", "callback", "handler:", "not", "a", "pointer", "value")
	}

	var _ dispFunc[D] = D.DispatchCallback // type check
	return Func[F]{
		ref: callback.Func(
			unsafe.Pointer(abi.FuncPCABIInternal(D.DispatchCallback)),
			// NOTE: this is still the `dispatcher`
			// we just converted it to pointer value as assumed
			Pointer(*(**byte)(mark.NoEscapePointer(&dispatcher))),
			SizeU(targetPC),
		),
	}
}

// CallbackContext is the context of a function call from JS to Go.
//
// Limitation: At most 16 args (including `this`) are allowed.
//
// The call can be made synchronously (go -> js -> go) and
// asynchronously (js -> go), the library supports both in one form.
//
// It is user's responsibility to ensure the callback does not call await-like
// functions such as js.Promise.Await() and yield.Thread() when the js side
// expects returning result synchronously, and if not so, the js side will
// get a Promise instead of immediate value.
//
// To return value, call one of .ReturnXxx functions.
// If none of the return functions was called, it returns js value "undefined".
// If more than one return functions were called, it returns the value passed
// in the last .ReturnXxx function.
//
// After returning from the target Go function,
// All heap references available in the slice returned by .Args() are
// freed, so users MUST NOT retain any of these references, and if
// you do want to use after the callback, call Ref.Clone() to make a copy.
type CallbackContext struct {
	dispFnPC     uint32
	handler      uint32
	targetPC     uint32
	retRef       bindings.Ref
	resolveFnRef bindings.Ref
	nargs        uint32
	args         [16]bindings.Ref
}

func (ctx *CallbackContext) Args() []Ref {
	return unsafe.Slice((*Ref)(mark.NoEscape(&ctx.args[0])), ctx.nargs)
}

func (ctx *CallbackContext) Return(ref Ref) bool {
	return bindings.Ref(True) == bindings.Replace(
		ctx.retRef, bindings.Ref(ref),
	)
}

func (ctx *CallbackContext) ReturnNum(n float64) bool {
	return bindings.Ref(True) == bindings.ReplaceNum(
		ctx.retRef, n,
	)
}

func (ctx *CallbackContext) ReturnBool(b bool) bool {
	return bindings.Ref(True) == bindings.ReplaceBool(
		ctx.retRef, bindings.Ref(Bool(b)),
	)
}

func (ctx *CallbackContext) ReturnString(s string) bool {
	return bindings.Ref(True) == bindings.ReplaceString(
		ctx.retRef,
		StringData(s),
		SizeU(len(s)),
	)
}

func init() {
	const offsetsOK = true &&
		unsafe.Offsetof(CallbackContext{}.dispFnPC) == 0 &&
		unsafe.Offsetof(CallbackContext{}.handler) == 4 &&
		unsafe.Offsetof(CallbackContext{}.targetPC) == 8 &&
		unsafe.Offsetof(CallbackContext{}.retRef) == 12 &&
		unsafe.Offsetof(CallbackContext{}.resolveFnRef) == 16 &&
		unsafe.Offsetof(CallbackContext{}.nargs) == 20 &&
		unsafe.Offsetof(CallbackContext{}.args) == 24

	if !offsetsOK {
		assert.Throw("invalid", "field", "offsets")
	}
}

// handleCallback called from wasm_export_run.
//
// NOTE: When inserting code after `callDispatcher` and `DONE`
// also update the LR update code pointing to this function
// inside callDispatcher (asm implementation).
func handleCallback(ref bindings.Ref) {
	var cb CallbackContext

	// defensive, to ensure js side sets nargs.
	cb.nargs = math.MaxUint32

	callback.Context(ref, Pointer(&cb))
	if cb.nargs > 16 {
		assert.Throw("too", "many", "args", "to", "callback", "func")
	}

	callDispatcher(
		uintptr(cb.handler),
		uintptr(cb.targetPC),
		mark.NoEscape(&cb),
		uintptr(cb.dispFnPC),
	)

	// DONE:
	Func[func(bool)]{ref: cb.resolveFnRef}.CallVoid(Undefined, false)
}
