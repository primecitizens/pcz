// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package networkingattributes

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/networkingattributes/bindings"
)

type GetNetworkDetailsCallbackFunc func(this js.Ref, networkAddresses *NetworkDetails) js.Ref

func (fn GetNetworkDetailsCallbackFunc) Register() js.Func[func(networkAddresses *NetworkDetails)] {
	return js.RegisterCallback[func(networkAddresses *NetworkDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetNetworkDetailsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkDetails
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

type GetNetworkDetailsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, networkAddresses *NetworkDetails) js.Ref
	Arg T
}

func (cb *GetNetworkDetailsCallback[T]) Register() js.Func[func(networkAddresses *NetworkDetails)] {
	return js.RegisterCallback[func(networkAddresses *NetworkDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetNetworkDetailsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NetworkDetails
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

type NetworkDetails struct {
	// MacAddress is "NetworkDetails.macAddress"
	//
	// Optional
	MacAddress js.String
	// Ipv4 is "NetworkDetails.ipv4"
	//
	// Optional
	Ipv4 js.String
	// Ipv6 is "NetworkDetails.ipv6"
	//
	// Optional
	Ipv6 js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkDetails with all fields set.
func (p NetworkDetails) FromRef(ref js.Ref) NetworkDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkDetails in the application heap.
func (p NetworkDetails) New() js.Ref {
	return bindings.NetworkDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkDetails) UpdateFrom(ref js.Ref) {
	bindings.NetworkDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkDetails) Update(ref js.Ref) {
	bindings.NetworkDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkDetails) FreeMembers(recursive bool) {
	js.Free(
		p.MacAddress.Ref(),
		p.Ipv4.Ref(),
		p.Ipv6.Ref(),
	)
	p.MacAddress = p.MacAddress.FromRef(js.Undefined)
	p.Ipv4 = p.Ipv4.FromRef(js.Undefined)
	p.Ipv6 = p.Ipv6.FromRef(js.Undefined)
}

// HasFuncGetNetworkDetails returns true if the function "WEBEXT.enterprise.networkingAttributes.getNetworkDetails" exists.
func HasFuncGetNetworkDetails() bool {
	return js.True == bindings.HasFuncGetNetworkDetails()
}

// FuncGetNetworkDetails returns the function "WEBEXT.enterprise.networkingAttributes.getNetworkDetails".
func FuncGetNetworkDetails() (fn js.Func[func() js.Promise[NetworkDetails]]) {
	bindings.FuncGetNetworkDetails(
		js.Pointer(&fn),
	)
	return
}

// GetNetworkDetails calls the function "WEBEXT.enterprise.networkingAttributes.getNetworkDetails" directly.
func GetNetworkDetails() (ret js.Promise[NetworkDetails]) {
	bindings.CallGetNetworkDetails(
		js.Pointer(&ret),
	)

	return
}

// TryGetNetworkDetails calls the function "WEBEXT.enterprise.networkingAttributes.getNetworkDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNetworkDetails() (ret js.Promise[NetworkDetails], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNetworkDetails(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
