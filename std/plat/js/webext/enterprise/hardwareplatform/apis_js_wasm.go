// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package hardwareplatform

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/hardwareplatform/bindings"
)

type HardwarePlatformInfo struct {
	// Model is "HardwarePlatformInfo.model"
	//
	// Optional
	Model js.String
	// Manufacturer is "HardwarePlatformInfo.manufacturer"
	//
	// Optional
	Manufacturer js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HardwarePlatformInfo with all fields set.
func (p HardwarePlatformInfo) FromRef(ref js.Ref) HardwarePlatformInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HardwarePlatformInfo in the application heap.
func (p HardwarePlatformInfo) New() js.Ref {
	return bindings.HardwarePlatformInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HardwarePlatformInfo) UpdateFrom(ref js.Ref) {
	bindings.HardwarePlatformInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HardwarePlatformInfo) Update(ref js.Ref) {
	bindings.HardwarePlatformInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HardwarePlatformInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Model.Ref(),
		p.Manufacturer.Ref(),
	)
	p.Model = p.Model.FromRef(js.Undefined)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
}

type HardwarePlatformInfoCallbackFunc func(this js.Ref, info *HardwarePlatformInfo) js.Ref

func (fn HardwarePlatformInfoCallbackFunc) Register() js.Func[func(info *HardwarePlatformInfo)] {
	return js.RegisterCallback[func(info *HardwarePlatformInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn HardwarePlatformInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HardwarePlatformInfo
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

type HardwarePlatformInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *HardwarePlatformInfo) js.Ref
	Arg T
}

func (cb *HardwarePlatformInfoCallback[T]) Register() js.Func[func(info *HardwarePlatformInfo)] {
	return js.RegisterCallback[func(info *HardwarePlatformInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *HardwarePlatformInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HardwarePlatformInfo
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

// HasFuncGetHardwarePlatformInfo returns true if the function "WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo" exists.
func HasFuncGetHardwarePlatformInfo() bool {
	return js.True == bindings.HasFuncGetHardwarePlatformInfo()
}

// FuncGetHardwarePlatformInfo returns the function "WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo".
func FuncGetHardwarePlatformInfo() (fn js.Func[func() js.Promise[HardwarePlatformInfo]]) {
	bindings.FuncGetHardwarePlatformInfo(
		js.Pointer(&fn),
	)
	return
}

// GetHardwarePlatformInfo calls the function "WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo" directly.
func GetHardwarePlatformInfo() (ret js.Promise[HardwarePlatformInfo]) {
	bindings.CallGetHardwarePlatformInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetHardwarePlatformInfo calls the function "WEBEXT.enterprise.hardwarePlatform.getHardwarePlatformInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetHardwarePlatformInfo() (ret js.Promise[HardwarePlatformInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetHardwarePlatformInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
