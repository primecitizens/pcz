// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package dns

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/dns/bindings"
)

type ResolveCallbackFunc func(this js.Ref, resolveInfo *ResolveCallbackResolveInfo) js.Ref

func (fn ResolveCallbackFunc) Register() js.Func[func(resolveInfo *ResolveCallbackResolveInfo)] {
	return js.RegisterCallback[func(resolveInfo *ResolveCallbackResolveInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResolveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ResolveCallbackResolveInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResolveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, resolveInfo *ResolveCallbackResolveInfo) js.Ref
	Arg T
}

func (cb *ResolveCallback[T]) Register() js.Func[func(resolveInfo *ResolveCallbackResolveInfo)] {
	return js.RegisterCallback[func(resolveInfo *ResolveCallbackResolveInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResolveCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ResolveCallbackResolveInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResolveCallbackResolveInfo struct {
	// ResultCode is "ResolveCallbackResolveInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// Address is "ResolveCallbackResolveInfo.address"
	//
	// Optional
	Address js.String

	FFI_USE_ResultCode bool // for ResultCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResolveCallbackResolveInfo with all fields set.
func (p ResolveCallbackResolveInfo) FromRef(ref js.Ref) ResolveCallbackResolveInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResolveCallbackResolveInfo in the application heap.
func (p ResolveCallbackResolveInfo) New() js.Ref {
	return bindings.ResolveCallbackResolveInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResolveCallbackResolveInfo) UpdateFrom(ref js.Ref) {
	bindings.ResolveCallbackResolveInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResolveCallbackResolveInfo) Update(ref js.Ref) {
	bindings.ResolveCallbackResolveInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResolveCallbackResolveInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Address.Ref(),
	)
	p.Address = p.Address.FromRef(js.Undefined)
}

// HasFuncResolve returns true if the function "WEBEXT.dns.resolve" exists.
func HasFuncResolve() bool {
	return js.True == bindings.HasFuncResolve()
}

// FuncResolve returns the function "WEBEXT.dns.resolve".
func FuncResolve() (fn js.Func[func(hostname js.String) js.Promise[ResolveCallbackResolveInfo]]) {
	bindings.FuncResolve(
		js.Pointer(&fn),
	)
	return
}

// Resolve calls the function "WEBEXT.dns.resolve" directly.
func Resolve(hostname js.String) (ret js.Promise[ResolveCallbackResolveInfo]) {
	bindings.CallResolve(
		js.Pointer(&ret),
		hostname.Ref(),
	)

	return
}

// TryResolve calls the function "WEBEXT.dns.resolve"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResolve(hostname js.String) (ret js.Promise[ResolveCallbackResolveInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResolve(
		js.Pointer(&ret), js.Pointer(&exception),
		hostname.Ref(),
	)

	return
}
