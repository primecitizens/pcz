// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package network

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/system/network/bindings"
)

type GetNetworkInterfacesCallbackFunc func(this js.Ref, networkInterfaces js.Array[NetworkInterface]) js.Ref

func (fn GetNetworkInterfacesCallbackFunc) Register() js.Func[func(networkInterfaces js.Array[NetworkInterface])] {
	return js.RegisterCallback[func(networkInterfaces js.Array[NetworkInterface])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetNetworkInterfacesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[NetworkInterface]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetNetworkInterfacesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, networkInterfaces js.Array[NetworkInterface]) js.Ref
	Arg T
}

func (cb *GetNetworkInterfacesCallback[T]) Register() js.Func[func(networkInterfaces js.Array[NetworkInterface])] {
	return js.RegisterCallback[func(networkInterfaces js.Array[NetworkInterface])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetNetworkInterfacesCallback[T]) DispatchCallback(
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

		js.Array[NetworkInterface]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NetworkInterface struct {
	// Name is "NetworkInterface.name"
	//
	// Optional
	Name js.String
	// Address is "NetworkInterface.address"
	//
	// Optional
	Address js.String
	// PrefixLength is "NetworkInterface.prefixLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_PrefixLength MUST be set to true to make this field effective.
	PrefixLength int32

	FFI_USE_PrefixLength bool // for PrefixLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkInterface with all fields set.
func (p NetworkInterface) FromRef(ref js.Ref) NetworkInterface {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkInterface in the application heap.
func (p NetworkInterface) New() js.Ref {
	return bindings.NetworkInterfaceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkInterface) UpdateFrom(ref js.Ref) {
	bindings.NetworkInterfaceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkInterface) Update(ref js.Ref) {
	bindings.NetworkInterfaceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkInterface) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Address.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Address = p.Address.FromRef(js.Undefined)
}

// HasFuncGetNetworkInterfaces returns true if the function "WEBEXT.system.network.getNetworkInterfaces" exists.
func HasFuncGetNetworkInterfaces() bool {
	return js.True == bindings.HasFuncGetNetworkInterfaces()
}

// FuncGetNetworkInterfaces returns the function "WEBEXT.system.network.getNetworkInterfaces".
func FuncGetNetworkInterfaces() (fn js.Func[func() js.Promise[js.Array[NetworkInterface]]]) {
	bindings.FuncGetNetworkInterfaces(
		js.Pointer(&fn),
	)
	return
}

// GetNetworkInterfaces calls the function "WEBEXT.system.network.getNetworkInterfaces" directly.
func GetNetworkInterfaces() (ret js.Promise[js.Array[NetworkInterface]]) {
	bindings.CallGetNetworkInterfaces(
		js.Pointer(&ret),
	)

	return
}

// TryGetNetworkInterfaces calls the function "WEBEXT.system.network.getNetworkInterfaces"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNetworkInterfaces() (ret js.Promise[js.Array[NetworkInterface]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNetworkInterfaces(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
