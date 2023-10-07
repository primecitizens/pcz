// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mdns

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/mdns/bindings"
)

type MDnsService struct {
	// ServiceName is "MDnsService.serviceName"
	//
	// Optional
	ServiceName js.String
	// ServiceHostPort is "MDnsService.serviceHostPort"
	//
	// Optional
	ServiceHostPort js.String
	// IpAddress is "MDnsService.ipAddress"
	//
	// Optional
	IpAddress js.String
	// ServiceData is "MDnsService.serviceData"
	//
	// Optional
	ServiceData js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MDnsService with all fields set.
func (p MDnsService) FromRef(ref js.Ref) MDnsService {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MDnsService in the application heap.
func (p MDnsService) New() js.Ref {
	return bindings.MDnsServiceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MDnsService) UpdateFrom(ref js.Ref) {
	bindings.MDnsServiceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MDnsService) Update(ref js.Ref) {
	bindings.MDnsServiceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MDnsService) FreeMembers(recursive bool) {
	js.Free(
		p.ServiceName.Ref(),
		p.ServiceHostPort.Ref(),
		p.IpAddress.Ref(),
		p.ServiceData.Ref(),
	)
	p.ServiceName = p.ServiceName.FromRef(js.Undefined)
	p.ServiceHostPort = p.ServiceHostPort.FromRef(js.Undefined)
	p.IpAddress = p.IpAddress.FromRef(js.Undefined)
	p.ServiceData = p.ServiceData.FromRef(js.Undefined)
}

type Properties struct {
	ref js.Ref
}

func (this Properties) Once() Properties {
	this.ref.Once()
	return this
}

func (this Properties) Ref() js.Ref {
	return this.ref
}

func (this Properties) FromRef(ref js.Ref) Properties {
	this.ref = ref
	return this
}

func (this Properties) Free() {
	this.ref.Free()
}

// HasFuncMAX_SERVICE_INSTANCES_PER_EVENT returns true if the static method "Properties.MAX_SERVICE_INSTANCES_PER_EVENT" exists.
func (this Properties) HasFuncMAX_SERVICE_INSTANCES_PER_EVENT() bool {
	return js.True == bindings.HasFuncPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
		this.ref,
	)
}

// FuncMAX_SERVICE_INSTANCES_PER_EVENT returns the static method "Properties.MAX_SERVICE_INSTANCES_PER_EVENT".
func (this Properties) FuncMAX_SERVICE_INSTANCES_PER_EVENT() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MAX_SERVICE_INSTANCES_PER_EVENT calls the static method "Properties.MAX_SERVICE_INSTANCES_PER_EVENT".
func (this Properties) MAX_SERVICE_INSTANCES_PER_EVENT() (ret int32) {
	bindings.CallPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMAX_SERVICE_INSTANCES_PER_EVENT calls the static method "Properties.MAX_SERVICE_INSTANCES_PER_EVENT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryMAX_SERVICE_INSTANCES_PER_EVENT() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesMAX_SERVICE_INSTANCES_PER_EVENT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasFuncForceDiscovery returns true if the function "WEBEXT.mdns.forceDiscovery" exists.
func HasFuncForceDiscovery() bool {
	return js.True == bindings.HasFuncForceDiscovery()
}

// FuncForceDiscovery returns the function "WEBEXT.mdns.forceDiscovery".
func FuncForceDiscovery() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncForceDiscovery(
		js.Pointer(&fn),
	)
	return
}

// ForceDiscovery calls the function "WEBEXT.mdns.forceDiscovery" directly.
func ForceDiscovery() (ret js.Promise[js.Void]) {
	bindings.CallForceDiscovery(
		js.Pointer(&ret),
	)

	return
}

// TryForceDiscovery calls the function "WEBEXT.mdns.forceDiscovery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForceDiscovery() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryForceDiscovery(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnServiceListEventCallbackFunc func(this js.Ref, services js.Array[MDnsService]) js.Ref

func (fn OnServiceListEventCallbackFunc) Register() js.Func[func(services js.Array[MDnsService])] {
	return js.RegisterCallback[func(services js.Array[MDnsService])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnServiceListEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[MDnsService]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnServiceListEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, services js.Array[MDnsService]) js.Ref
	Arg T
}

func (cb *OnServiceListEventCallback[T]) Register() js.Func[func(services js.Array[MDnsService])] {
	return js.RegisterCallback[func(services js.Array[MDnsService])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnServiceListEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Array[MDnsService]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnServiceList returns true if the function "WEBEXT.mdns.onServiceList.addListener" exists.
func HasFuncOnServiceList() bool {
	return js.True == bindings.HasFuncOnServiceList()
}

// FuncOnServiceList returns the function "WEBEXT.mdns.onServiceList.addListener".
func FuncOnServiceList() (fn js.Func[func(callback js.Func[func(services js.Array[MDnsService])])]) {
	bindings.FuncOnServiceList(
		js.Pointer(&fn),
	)
	return
}

// OnServiceList calls the function "WEBEXT.mdns.onServiceList.addListener" directly.
func OnServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret js.Void) {
	bindings.CallOnServiceList(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnServiceList calls the function "WEBEXT.mdns.onServiceList.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnServiceList(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffServiceList returns true if the function "WEBEXT.mdns.onServiceList.removeListener" exists.
func HasFuncOffServiceList() bool {
	return js.True == bindings.HasFuncOffServiceList()
}

// FuncOffServiceList returns the function "WEBEXT.mdns.onServiceList.removeListener".
func FuncOffServiceList() (fn js.Func[func(callback js.Func[func(services js.Array[MDnsService])])]) {
	bindings.FuncOffServiceList(
		js.Pointer(&fn),
	)
	return
}

// OffServiceList calls the function "WEBEXT.mdns.onServiceList.removeListener" directly.
func OffServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret js.Void) {
	bindings.CallOffServiceList(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffServiceList calls the function "WEBEXT.mdns.onServiceList.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffServiceList(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnServiceList returns true if the function "WEBEXT.mdns.onServiceList.hasListener" exists.
func HasFuncHasOnServiceList() bool {
	return js.True == bindings.HasFuncHasOnServiceList()
}

// FuncHasOnServiceList returns the function "WEBEXT.mdns.onServiceList.hasListener".
func FuncHasOnServiceList() (fn js.Func[func(callback js.Func[func(services js.Array[MDnsService])]) bool]) {
	bindings.FuncHasOnServiceList(
		js.Pointer(&fn),
	)
	return
}

// HasOnServiceList calls the function "WEBEXT.mdns.onServiceList.hasListener" directly.
func HasOnServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret bool) {
	bindings.CallHasOnServiceList(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnServiceList calls the function "WEBEXT.mdns.onServiceList.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnServiceList(callback js.Func[func(services js.Array[MDnsService])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnServiceList(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
