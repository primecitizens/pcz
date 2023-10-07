// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cpu

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/system/cpu/bindings"
)

type CpuTime struct {
	// User is "CpuTime.user"
	//
	// Optional
	//
	// NOTE: FFI_USE_User MUST be set to true to make this field effective.
	User float64
	// Kernel is "CpuTime.kernel"
	//
	// Optional
	//
	// NOTE: FFI_USE_Kernel MUST be set to true to make this field effective.
	Kernel float64
	// Idle is "CpuTime.idle"
	//
	// Optional
	//
	// NOTE: FFI_USE_Idle MUST be set to true to make this field effective.
	Idle float64
	// Total is "CpuTime.total"
	//
	// Optional
	//
	// NOTE: FFI_USE_Total MUST be set to true to make this field effective.
	Total float64

	FFI_USE_User   bool // for User.
	FFI_USE_Kernel bool // for Kernel.
	FFI_USE_Idle   bool // for Idle.
	FFI_USE_Total  bool // for Total.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CpuTime with all fields set.
func (p CpuTime) FromRef(ref js.Ref) CpuTime {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CpuTime in the application heap.
func (p CpuTime) New() js.Ref {
	return bindings.CpuTimeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CpuTime) UpdateFrom(ref js.Ref) {
	bindings.CpuTimeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CpuTime) Update(ref js.Ref) {
	bindings.CpuTimeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CpuTime) FreeMembers(recursive bool) {
}

type ProcessorInfo struct {
	// Usage is "ProcessorInfo.usage"
	//
	// Optional
	//
	// NOTE: Usage.FFI_USE MUST be set to true to get Usage used.
	Usage CpuTime

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProcessorInfo with all fields set.
func (p ProcessorInfo) FromRef(ref js.Ref) ProcessorInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProcessorInfo in the application heap.
func (p ProcessorInfo) New() js.Ref {
	return bindings.ProcessorInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProcessorInfo) UpdateFrom(ref js.Ref) {
	bindings.ProcessorInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProcessorInfo) Update(ref js.Ref) {
	bindings.ProcessorInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProcessorInfo) FreeMembers(recursive bool) {
	if recursive {
		p.Usage.FreeMembers(true)
	}
}

type CpuInfo struct {
	// NumOfProcessors is "CpuInfo.numOfProcessors"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumOfProcessors MUST be set to true to make this field effective.
	NumOfProcessors int32
	// ArchName is "CpuInfo.archName"
	//
	// Optional
	ArchName js.String
	// ModelName is "CpuInfo.modelName"
	//
	// Optional
	ModelName js.String
	// Features is "CpuInfo.features"
	//
	// Optional
	Features js.Array[js.String]
	// Processors is "CpuInfo.processors"
	//
	// Optional
	Processors js.Array[ProcessorInfo]
	// Temperatures is "CpuInfo.temperatures"
	//
	// Optional
	Temperatures js.Array[float64]

	FFI_USE_NumOfProcessors bool // for NumOfProcessors.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CpuInfo with all fields set.
func (p CpuInfo) FromRef(ref js.Ref) CpuInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CpuInfo in the application heap.
func (p CpuInfo) New() js.Ref {
	return bindings.CpuInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CpuInfo) UpdateFrom(ref js.Ref) {
	bindings.CpuInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CpuInfo) Update(ref js.Ref) {
	bindings.CpuInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CpuInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ArchName.Ref(),
		p.ModelName.Ref(),
		p.Features.Ref(),
		p.Processors.Ref(),
		p.Temperatures.Ref(),
	)
	p.ArchName = p.ArchName.FromRef(js.Undefined)
	p.ModelName = p.ModelName.FromRef(js.Undefined)
	p.Features = p.Features.FromRef(js.Undefined)
	p.Processors = p.Processors.FromRef(js.Undefined)
	p.Temperatures = p.Temperatures.FromRef(js.Undefined)
}

type CpuInfoCallbackFunc func(this js.Ref, info *CpuInfo) js.Ref

func (fn CpuInfoCallbackFunc) Register() js.Func[func(info *CpuInfo)] {
	return js.RegisterCallback[func(info *CpuInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CpuInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CpuInfo
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

type CpuInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *CpuInfo) js.Ref
	Arg T
}

func (cb *CpuInfoCallback[T]) Register() js.Func[func(info *CpuInfo)] {
	return js.RegisterCallback[func(info *CpuInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CpuInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CpuInfo
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

// HasFuncGetInfo returns true if the function "WEBEXT.system.cpu.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.system.cpu.getInfo".
func FuncGetInfo() (fn js.Func[func() js.Promise[CpuInfo]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.system.cpu.getInfo" directly.
func GetInfo() (ret js.Promise[CpuInfo]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.system.cpu.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo() (ret js.Promise[CpuInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
