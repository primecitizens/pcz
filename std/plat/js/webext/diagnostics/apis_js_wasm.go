// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package diagnostics

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/diagnostics/bindings"
)

type SendPacketCallbackFunc func(this js.Ref, result *SendPacketResult) js.Ref

func (fn SendPacketCallbackFunc) Register() js.Func[func(result *SendPacketResult)] {
	return js.RegisterCallback[func(result *SendPacketResult)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendPacketCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendPacketResult
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

type SendPacketCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *SendPacketResult) js.Ref
	Arg T
}

func (cb *SendPacketCallback[T]) Register() js.Func[func(result *SendPacketResult)] {
	return js.RegisterCallback[func(result *SendPacketResult)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendPacketCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendPacketResult
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

type SendPacketResult struct {
	// Ip is "SendPacketResult.ip"
	//
	// Optional
	Ip js.String
	// Latency is "SendPacketResult.latency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Latency MUST be set to true to make this field effective.
	Latency float64

	FFI_USE_Latency bool // for Latency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendPacketResult with all fields set.
func (p SendPacketResult) FromRef(ref js.Ref) SendPacketResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendPacketResult in the application heap.
func (p SendPacketResult) New() js.Ref {
	return bindings.SendPacketResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendPacketResult) UpdateFrom(ref js.Ref) {
	bindings.SendPacketResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendPacketResult) Update(ref js.Ref) {
	bindings.SendPacketResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendPacketResult) FreeMembers(recursive bool) {
	js.Free(
		p.Ip.Ref(),
	)
	p.Ip = p.Ip.FromRef(js.Undefined)
}

type SendPacketOptions struct {
	// Ip is "SendPacketOptions.ip"
	//
	// Optional
	Ip js.String
	// Ttl is "SendPacketOptions.ttl"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ttl MUST be set to true to make this field effective.
	Ttl int32
	// Timeout is "SendPacketOptions.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout int32
	// Size is "SendPacketOptions.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size int32

	FFI_USE_Ttl     bool // for Ttl.
	FFI_USE_Timeout bool // for Timeout.
	FFI_USE_Size    bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendPacketOptions with all fields set.
func (p SendPacketOptions) FromRef(ref js.Ref) SendPacketOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendPacketOptions in the application heap.
func (p SendPacketOptions) New() js.Ref {
	return bindings.SendPacketOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendPacketOptions) UpdateFrom(ref js.Ref) {
	bindings.SendPacketOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendPacketOptions) Update(ref js.Ref) {
	bindings.SendPacketOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendPacketOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Ip.Ref(),
	)
	p.Ip = p.Ip.FromRef(js.Undefined)
}

// HasFuncSendPacket returns true if the function "WEBEXT.diagnostics.sendPacket" exists.
func HasFuncSendPacket() bool {
	return js.True == bindings.HasFuncSendPacket()
}

// FuncSendPacket returns the function "WEBEXT.diagnostics.sendPacket".
func FuncSendPacket() (fn js.Func[func(options SendPacketOptions) js.Promise[SendPacketResult]]) {
	bindings.FuncSendPacket(
		js.Pointer(&fn),
	)
	return
}

// SendPacket calls the function "WEBEXT.diagnostics.sendPacket" directly.
func SendPacket(options SendPacketOptions) (ret js.Promise[SendPacketResult]) {
	bindings.CallSendPacket(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySendPacket calls the function "WEBEXT.diagnostics.sendPacket"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendPacket(options SendPacketOptions) (ret js.Promise[SendPacketResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendPacket(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}
