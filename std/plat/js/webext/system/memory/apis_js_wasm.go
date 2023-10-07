// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package memory

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/system/memory/bindings"
)

type MemoryInfo struct {
	// Capacity is "MemoryInfo.capacity"
	//
	// Optional
	//
	// NOTE: FFI_USE_Capacity MUST be set to true to make this field effective.
	Capacity float64
	// AvailableCapacity is "MemoryInfo.availableCapacity"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableCapacity MUST be set to true to make this field effective.
	AvailableCapacity float64

	FFI_USE_Capacity          bool // for Capacity.
	FFI_USE_AvailableCapacity bool // for AvailableCapacity.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryInfo with all fields set.
func (p MemoryInfo) FromRef(ref js.Ref) MemoryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryInfo in the application heap.
func (p MemoryInfo) New() js.Ref {
	return bindings.MemoryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MemoryInfo) UpdateFrom(ref js.Ref) {
	bindings.MemoryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MemoryInfo) Update(ref js.Ref) {
	bindings.MemoryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MemoryInfo) FreeMembers(recursive bool) {
}

type MemoryInfoCallbackFunc func(this js.Ref, info *MemoryInfo) js.Ref

func (fn MemoryInfoCallbackFunc) Register() js.Func[func(info *MemoryInfo)] {
	return js.RegisterCallback[func(info *MemoryInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MemoryInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryInfo
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

type MemoryInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *MemoryInfo) js.Ref
	Arg T
}

func (cb *MemoryInfoCallback[T]) Register() js.Func[func(info *MemoryInfo)] {
	return js.RegisterCallback[func(info *MemoryInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MemoryInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryInfo
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

// HasFuncGetInfo returns true if the function "WEBEXT.system.memory.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.system.memory.getInfo".
func FuncGetInfo() (fn js.Func[func() js.Promise[MemoryInfo]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.system.memory.getInfo" directly.
func GetInfo() (ret js.Promise[MemoryInfo]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.system.memory.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo() (ret js.Promise[MemoryInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
