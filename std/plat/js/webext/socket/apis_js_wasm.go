// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package socket

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/socket/bindings"
)

type AcceptCallbackFunc func(this js.Ref, acceptInfo *AcceptInfo) js.Ref

func (fn AcceptCallbackFunc) Register() js.Func[func(acceptInfo *AcceptInfo)] {
	return js.RegisterCallback[func(acceptInfo *AcceptInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AcceptCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AcceptInfo
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

type AcceptCallback[T any] struct {
	Fn  func(arg T, this js.Ref, acceptInfo *AcceptInfo) js.Ref
	Arg T
}

func (cb *AcceptCallback[T]) Register() js.Func[func(acceptInfo *AcceptInfo)] {
	return js.RegisterCallback[func(acceptInfo *AcceptInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AcceptCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AcceptInfo
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

type AcceptInfo struct {
	// ResultCode is "AcceptInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// SocketId is "AcceptInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32

	FFI_USE_ResultCode bool // for ResultCode.
	FFI_USE_SocketId   bool // for SocketId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AcceptInfo with all fields set.
func (p AcceptInfo) FromRef(ref js.Ref) AcceptInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AcceptInfo in the application heap.
func (p AcceptInfo) New() js.Ref {
	return bindings.AcceptInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AcceptInfo) UpdateFrom(ref js.Ref) {
	bindings.AcceptInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AcceptInfo) Update(ref js.Ref) {
	bindings.AcceptInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AcceptInfo) FreeMembers(recursive bool) {
}

type BindCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn BindCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BindCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BindCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *BindCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BindCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConnectCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn ConnectCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ConnectCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *ConnectCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ConnectCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateCallbackFunc func(this js.Ref, createInfo *CreateInfo) js.Ref

func (fn CreateCallbackFunc) Register() js.Func[func(createInfo *CreateInfo)] {
	return js.RegisterCallback[func(createInfo *CreateInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateInfo
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

type CreateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, createInfo *CreateInfo) js.Ref
	Arg T
}

func (cb *CreateCallback[T]) Register() js.Func[func(createInfo *CreateInfo)] {
	return js.RegisterCallback[func(createInfo *CreateInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateInfo
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

type CreateInfo struct {
	// SocketId is "CreateInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32

	FFI_USE_SocketId bool // for SocketId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateInfo with all fields set.
func (p CreateInfo) FromRef(ref js.Ref) CreateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateInfo in the application heap.
func (p CreateInfo) New() js.Ref {
	return bindings.CreateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateInfo) UpdateFrom(ref js.Ref) {
	bindings.CreateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateInfo) Update(ref js.Ref) {
	bindings.CreateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateInfo) FreeMembers(recursive bool) {
}

type CreateOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateOptions with all fields set.
func (p CreateOptions) FromRef(ref js.Ref) CreateOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateOptions in the application heap.
func (p CreateOptions) New() js.Ref {
	return bindings.CreateOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateOptions) UpdateFrom(ref js.Ref) {
	bindings.CreateOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateOptions) Update(ref js.Ref) {
	bindings.CreateOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateOptions) FreeMembers(recursive bool) {
}

type GetInfoCallbackFunc func(this js.Ref, result *SocketInfo) js.Ref

func (fn GetInfoCallbackFunc) Register() js.Func[func(result *SocketInfo)] {
	return js.RegisterCallback[func(result *SocketInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SocketInfo
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

type GetInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *SocketInfo) js.Ref
	Arg T
}

func (cb *GetInfoCallback[T]) Register() js.Func[func(result *SocketInfo)] {
	return js.RegisterCallback[func(result *SocketInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SocketInfo
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

type SocketType uint32

const (
	_ SocketType = iota

	SocketType_TCP
	SocketType_UDP
)

func (SocketType) FromRef(str js.Ref) SocketType {
	return SocketType(bindings.ConstOfSocketType(str))
}

func (x SocketType) String() (string, bool) {
	switch x {
	case SocketType_TCP:
		return "tcp", true
	case SocketType_UDP:
		return "udp", true
	default:
		return "", false
	}
}

type SocketInfo struct {
	// SocketType is "SocketInfo.socketType"
	//
	// Optional
	SocketType SocketType
	// Connected is "SocketInfo.connected"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connected MUST be set to true to make this field effective.
	Connected bool
	// PeerAddress is "SocketInfo.peerAddress"
	//
	// Optional
	PeerAddress js.String
	// PeerPort is "SocketInfo.peerPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_PeerPort MUST be set to true to make this field effective.
	PeerPort int32
	// LocalAddress is "SocketInfo.localAddress"
	//
	// Optional
	LocalAddress js.String
	// LocalPort is "SocketInfo.localPort"
	//
	// Optional
	//
	// NOTE: FFI_USE_LocalPort MUST be set to true to make this field effective.
	LocalPort int32

	FFI_USE_Connected bool // for Connected.
	FFI_USE_PeerPort  bool // for PeerPort.
	FFI_USE_LocalPort bool // for LocalPort.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SocketInfo with all fields set.
func (p SocketInfo) FromRef(ref js.Ref) SocketInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SocketInfo in the application heap.
func (p SocketInfo) New() js.Ref {
	return bindings.SocketInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SocketInfo) UpdateFrom(ref js.Ref) {
	bindings.SocketInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SocketInfo) Update(ref js.Ref) {
	bindings.SocketInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SocketInfo) FreeMembers(recursive bool) {
	js.Free(
		p.PeerAddress.Ref(),
		p.LocalAddress.Ref(),
	)
	p.PeerAddress = p.PeerAddress.FromRef(js.Undefined)
	p.LocalAddress = p.LocalAddress.FromRef(js.Undefined)
}

type GetJoinedGroupsCallbackFunc func(this js.Ref, groups js.Array[js.String]) js.Ref

func (fn GetJoinedGroupsCallbackFunc) Register() js.Func[func(groups js.Array[js.String])] {
	return js.RegisterCallback[func(groups js.Array[js.String])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetJoinedGroupsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetJoinedGroupsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, groups js.Array[js.String]) js.Ref
	Arg T
}

func (cb *GetJoinedGroupsCallback[T]) Register() js.Func[func(groups js.Array[js.String])] {
	return js.RegisterCallback[func(groups js.Array[js.String])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetJoinedGroupsCallback[T]) DispatchCallback(
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

		js.Array[js.String]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetNetworkCallbackFunc func(this js.Ref, result js.Array[NetworkInterface]) js.Ref

func (fn GetNetworkCallbackFunc) Register() js.Func[func(result js.Array[NetworkInterface])] {
	return js.RegisterCallback[func(result js.Array[NetworkInterface])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetNetworkCallbackFunc) DispatchCallback(
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

type GetNetworkCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[NetworkInterface]) js.Ref
	Arg T
}

func (cb *GetNetworkCallback[T]) Register() js.Func[func(result js.Array[NetworkInterface])] {
	return js.RegisterCallback[func(result js.Array[NetworkInterface])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetNetworkCallback[T]) DispatchCallback(
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

type JoinGroupCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn JoinGroupCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn JoinGroupCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type JoinGroupCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *JoinGroupCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *JoinGroupCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LeaveGroupCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn LeaveGroupCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LeaveGroupCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LeaveGroupCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *LeaveGroupCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LeaveGroupCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListenCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn ListenCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListenCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ListenCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *ListenCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListenCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReadCallbackFunc func(this js.Ref, readInfo *ReadInfo) js.Ref

func (fn ReadCallbackFunc) Register() js.Func[func(readInfo *ReadInfo)] {
	return js.RegisterCallback[func(readInfo *ReadInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReadCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadInfo
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

type ReadCallback[T any] struct {
	Fn  func(arg T, this js.Ref, readInfo *ReadInfo) js.Ref
	Arg T
}

func (cb *ReadCallback[T]) Register() js.Func[func(readInfo *ReadInfo)] {
	return js.RegisterCallback[func(readInfo *ReadInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReadCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReadInfo
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

type ReadInfo struct {
	// ResultCode is "ReadInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// Data is "ReadInfo.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_ResultCode bool // for ResultCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadInfo with all fields set.
func (p ReadInfo) FromRef(ref js.Ref) ReadInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadInfo in the application heap.
func (p ReadInfo) New() js.Ref {
	return bindings.ReadInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadInfo) UpdateFrom(ref js.Ref) {
	bindings.ReadInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadInfo) Update(ref js.Ref) {
	bindings.ReadInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type RecvFromCallbackFunc func(this js.Ref, recvFromInfo *RecvFromInfo) js.Ref

func (fn RecvFromCallbackFunc) Register() js.Func[func(recvFromInfo *RecvFromInfo)] {
	return js.RegisterCallback[func(recvFromInfo *RecvFromInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RecvFromCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RecvFromInfo
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

type RecvFromCallback[T any] struct {
	Fn  func(arg T, this js.Ref, recvFromInfo *RecvFromInfo) js.Ref
	Arg T
}

func (cb *RecvFromCallback[T]) Register() js.Func[func(recvFromInfo *RecvFromInfo)] {
	return js.RegisterCallback[func(recvFromInfo *RecvFromInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RecvFromCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 RecvFromInfo
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

type RecvFromInfo struct {
	// ResultCode is "RecvFromInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// Data is "RecvFromInfo.data"
	//
	// Optional
	Data js.ArrayBuffer
	// Address is "RecvFromInfo.address"
	//
	// Optional
	Address js.String
	// Port is "RecvFromInfo.port"
	//
	// Optional
	//
	// NOTE: FFI_USE_Port MUST be set to true to make this field effective.
	Port int32

	FFI_USE_ResultCode bool // for ResultCode.
	FFI_USE_Port       bool // for Port.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RecvFromInfo with all fields set.
func (p RecvFromInfo) FromRef(ref js.Ref) RecvFromInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RecvFromInfo in the application heap.
func (p RecvFromInfo) New() js.Ref {
	return bindings.RecvFromInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RecvFromInfo) UpdateFrom(ref js.Ref) {
	bindings.RecvFromInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RecvFromInfo) Update(ref js.Ref) {
	bindings.RecvFromInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RecvFromInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Address.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Address = p.Address.FromRef(js.Undefined)
}

type SecureCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SecureCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SecureCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SecureCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SecureCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SecureCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TLSVersionConstraints struct {
	// Min is "TLSVersionConstraints.min"
	//
	// Optional
	Min js.String
	// Max is "TLSVersionConstraints.max"
	//
	// Optional
	Max js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TLSVersionConstraints with all fields set.
func (p TLSVersionConstraints) FromRef(ref js.Ref) TLSVersionConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TLSVersionConstraints in the application heap.
func (p TLSVersionConstraints) New() js.Ref {
	return bindings.TLSVersionConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TLSVersionConstraints) UpdateFrom(ref js.Ref) {
	bindings.TLSVersionConstraintsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TLSVersionConstraints) Update(ref js.Ref) {
	bindings.TLSVersionConstraintsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TLSVersionConstraints) FreeMembers(recursive bool) {
	js.Free(
		p.Min.Ref(),
		p.Max.Ref(),
	)
	p.Min = p.Min.FromRef(js.Undefined)
	p.Max = p.Max.FromRef(js.Undefined)
}

type SecureOptions struct {
	// TlsVersion is "SecureOptions.tlsVersion"
	//
	// Optional
	//
	// NOTE: TlsVersion.FFI_USE MUST be set to true to get TlsVersion used.
	TlsVersion TLSVersionConstraints

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SecureOptions with all fields set.
func (p SecureOptions) FromRef(ref js.Ref) SecureOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SecureOptions in the application heap.
func (p SecureOptions) New() js.Ref {
	return bindings.SecureOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SecureOptions) UpdateFrom(ref js.Ref) {
	bindings.SecureOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SecureOptions) Update(ref js.Ref) {
	bindings.SecureOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SecureOptions) FreeMembers(recursive bool) {
	if recursive {
		p.TlsVersion.FreeMembers(true)
	}
}

type SendToCallbackFunc func(this js.Ref, writeInfo *WriteInfo) js.Ref

func (fn SendToCallbackFunc) Register() js.Func[func(writeInfo *WriteInfo)] {
	return js.RegisterCallback[func(writeInfo *WriteInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendToCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteInfo
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

type SendToCallback[T any] struct {
	Fn  func(arg T, this js.Ref, writeInfo *WriteInfo) js.Ref
	Arg T
}

func (cb *SendToCallback[T]) Register() js.Func[func(writeInfo *WriteInfo)] {
	return js.RegisterCallback[func(writeInfo *WriteInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendToCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteInfo
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

type WriteInfo struct {
	// BytesWritten is "WriteInfo.bytesWritten"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesWritten MUST be set to true to make this field effective.
	BytesWritten int32

	FFI_USE_BytesWritten bool // for BytesWritten.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WriteInfo with all fields set.
func (p WriteInfo) FromRef(ref js.Ref) WriteInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WriteInfo in the application heap.
func (p WriteInfo) New() js.Ref {
	return bindings.WriteInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WriteInfo) UpdateFrom(ref js.Ref) {
	bindings.WriteInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WriteInfo) Update(ref js.Ref) {
	bindings.WriteInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WriteInfo) FreeMembers(recursive bool) {
}

type SetKeepAliveCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn SetKeepAliveCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetKeepAliveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetKeepAliveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *SetKeepAliveCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetKeepAliveCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMulticastLoopbackModeCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SetMulticastLoopbackModeCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetMulticastLoopbackModeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMulticastLoopbackModeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SetMulticastLoopbackModeCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetMulticastLoopbackModeCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMulticastTimeToLiveCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SetMulticastTimeToLiveCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetMulticastTimeToLiveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMulticastTimeToLiveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SetMulticastTimeToLiveCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetMulticastTimeToLiveCallback[T]) DispatchCallback(
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetNoDelayCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn SetNoDelayCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetNoDelayCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetNoDelayCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *SetNoDelayCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetNoDelayCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WriteCallbackFunc func(this js.Ref, writeInfo *WriteInfo) js.Ref

func (fn WriteCallbackFunc) Register() js.Func[func(writeInfo *WriteInfo)] {
	return js.RegisterCallback[func(writeInfo *WriteInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WriteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteInfo
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

type WriteCallback[T any] struct {
	Fn  func(arg T, this js.Ref, writeInfo *WriteInfo) js.Ref
	Arg T
}

func (cb *WriteCallback[T]) Register() js.Func[func(writeInfo *WriteInfo)] {
	return js.RegisterCallback[func(writeInfo *WriteInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WriteCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WriteInfo
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

// HasFuncAccept returns true if the function "WEBEXT.socket.accept" exists.
func HasFuncAccept() bool {
	return js.True == bindings.HasFuncAccept()
}

// FuncAccept returns the function "WEBEXT.socket.accept".
func FuncAccept() (fn js.Func[func(socketId int32, callback js.Func[func(acceptInfo *AcceptInfo)])]) {
	bindings.FuncAccept(
		js.Pointer(&fn),
	)
	return
}

// Accept calls the function "WEBEXT.socket.accept" directly.
func Accept(socketId int32, callback js.Func[func(acceptInfo *AcceptInfo)]) (ret js.Void) {
	bindings.CallAccept(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryAccept calls the function "WEBEXT.socket.accept"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAccept(socketId int32, callback js.Func[func(acceptInfo *AcceptInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncBind returns true if the function "WEBEXT.socket.bind" exists.
func HasFuncBind() bool {
	return js.True == bindings.HasFuncBind()
}

// FuncBind returns the function "WEBEXT.socket.bind".
func FuncBind() (fn js.Func[func(socketId int32, address js.String, port int32, callback js.Func[func(result int32)])]) {
	bindings.FuncBind(
		js.Pointer(&fn),
	)
	return
}

// Bind calls the function "WEBEXT.socket.bind" directly.
func Bind(socketId int32, address js.String, port int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallBind(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// TryBind calls the function "WEBEXT.socket.bind"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryBind(socketId int32, address js.String, port int32, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBind(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		address.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// HasFuncConnect returns true if the function "WEBEXT.socket.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.socket.connect".
func FuncConnect() (fn js.Func[func(socketId int32, hostname js.String, port int32, callback js.Func[func(result int32)])]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.socket.connect" directly.
func Connect(socketId int32, hostname js.String, port int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallConnect(
		js.Pointer(&ret),
		int32(socketId),
		hostname.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// TryConnect calls the function "WEBEXT.socket.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(socketId int32, hostname js.String, port int32, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		hostname.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.socket.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.socket.create".
func FuncCreate() (fn js.Func[func(typ SocketType, options CreateOptions, callback js.Func[func(createInfo *CreateInfo)])]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.socket.create" directly.
func Create(typ SocketType, options CreateOptions, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void) {
	bindings.CallCreate(
		js.Pointer(&ret),
		uint32(typ),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryCreate calls the function "WEBEXT.socket.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(typ SocketType, options CreateOptions, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncDestroy returns true if the function "WEBEXT.socket.destroy" exists.
func HasFuncDestroy() bool {
	return js.True == bindings.HasFuncDestroy()
}

// FuncDestroy returns the function "WEBEXT.socket.destroy".
func FuncDestroy() (fn js.Func[func(socketId int32)]) {
	bindings.FuncDestroy(
		js.Pointer(&fn),
	)
	return
}

// Destroy calls the function "WEBEXT.socket.destroy" directly.
func Destroy(socketId int32) (ret js.Void) {
	bindings.CallDestroy(
		js.Pointer(&ret),
		int32(socketId),
	)

	return
}

// TryDestroy calls the function "WEBEXT.socket.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDestroy(socketId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDestroy(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.socket.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.socket.disconnect".
func FuncDisconnect() (fn js.Func[func(socketId int32)]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.socket.disconnect" directly.
func Disconnect(socketId int32) (ret js.Void) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(socketId),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.socket.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(socketId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.socket.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.socket.getInfo".
func FuncGetInfo() (fn js.Func[func(socketId int32, callback js.Func[func(result *SocketInfo)])]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.socket.getInfo" directly.
func GetInfo(socketId int32, callback js.Func[func(result *SocketInfo)]) (ret js.Void) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.socket.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo(socketId int32, callback js.Func[func(result *SocketInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncGetJoinedGroups returns true if the function "WEBEXT.socket.getJoinedGroups" exists.
func HasFuncGetJoinedGroups() bool {
	return js.True == bindings.HasFuncGetJoinedGroups()
}

// FuncGetJoinedGroups returns the function "WEBEXT.socket.getJoinedGroups".
func FuncGetJoinedGroups() (fn js.Func[func(socketId int32, callback js.Func[func(groups js.Array[js.String])])]) {
	bindings.FuncGetJoinedGroups(
		js.Pointer(&fn),
	)
	return
}

// GetJoinedGroups calls the function "WEBEXT.socket.getJoinedGroups" directly.
func GetJoinedGroups(socketId int32, callback js.Func[func(groups js.Array[js.String])]) (ret js.Void) {
	bindings.CallGetJoinedGroups(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetJoinedGroups calls the function "WEBEXT.socket.getJoinedGroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetJoinedGroups(socketId int32, callback js.Func[func(groups js.Array[js.String])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetJoinedGroups(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncGetNetworkList returns true if the function "WEBEXT.socket.getNetworkList" exists.
func HasFuncGetNetworkList() bool {
	return js.True == bindings.HasFuncGetNetworkList()
}

// FuncGetNetworkList returns the function "WEBEXT.socket.getNetworkList".
func FuncGetNetworkList() (fn js.Func[func(callback js.Func[func(result js.Array[NetworkInterface])])]) {
	bindings.FuncGetNetworkList(
		js.Pointer(&fn),
	)
	return
}

// GetNetworkList calls the function "WEBEXT.socket.getNetworkList" directly.
func GetNetworkList(callback js.Func[func(result js.Array[NetworkInterface])]) (ret js.Void) {
	bindings.CallGetNetworkList(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetNetworkList calls the function "WEBEXT.socket.getNetworkList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNetworkList(callback js.Func[func(result js.Array[NetworkInterface])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNetworkList(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncJoinGroup returns true if the function "WEBEXT.socket.joinGroup" exists.
func HasFuncJoinGroup() bool {
	return js.True == bindings.HasFuncJoinGroup()
}

// FuncJoinGroup returns the function "WEBEXT.socket.joinGroup".
func FuncJoinGroup() (fn js.Func[func(socketId int32, address js.String, callback js.Func[func(result int32)])]) {
	bindings.FuncJoinGroup(
		js.Pointer(&fn),
	)
	return
}

// JoinGroup calls the function "WEBEXT.socket.joinGroup" directly.
func JoinGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallJoinGroup(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// TryJoinGroup calls the function "WEBEXT.socket.joinGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryJoinGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryJoinGroup(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncLeaveGroup returns true if the function "WEBEXT.socket.leaveGroup" exists.
func HasFuncLeaveGroup() bool {
	return js.True == bindings.HasFuncLeaveGroup()
}

// FuncLeaveGroup returns the function "WEBEXT.socket.leaveGroup".
func FuncLeaveGroup() (fn js.Func[func(socketId int32, address js.String, callback js.Func[func(result int32)])]) {
	bindings.FuncLeaveGroup(
		js.Pointer(&fn),
	)
	return
}

// LeaveGroup calls the function "WEBEXT.socket.leaveGroup" directly.
func LeaveGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallLeaveGroup(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// TryLeaveGroup calls the function "WEBEXT.socket.leaveGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLeaveGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLeaveGroup(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncListen returns true if the function "WEBEXT.socket.listen" exists.
func HasFuncListen() bool {
	return js.True == bindings.HasFuncListen()
}

// FuncListen returns the function "WEBEXT.socket.listen".
func FuncListen() (fn js.Func[func(socketId int32, address js.String, port int32, backlog int32, callback js.Func[func(result int32)])]) {
	bindings.FuncListen(
		js.Pointer(&fn),
	)
	return
}

// Listen calls the function "WEBEXT.socket.listen" directly.
func Listen(socketId int32, address js.String, port int32, backlog int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallListen(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		int32(port),
		int32(backlog),
		callback.Ref(),
	)

	return
}

// TryListen calls the function "WEBEXT.socket.listen"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListen(socketId int32, address js.String, port int32, backlog int32, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryListen(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		address.Ref(),
		int32(port),
		int32(backlog),
		callback.Ref(),
	)

	return
}

// HasFuncRead returns true if the function "WEBEXT.socket.read" exists.
func HasFuncRead() bool {
	return js.True == bindings.HasFuncRead()
}

// FuncRead returns the function "WEBEXT.socket.read".
func FuncRead() (fn js.Func[func(socketId int32, bufferSize int32, callback js.Func[func(readInfo *ReadInfo)])]) {
	bindings.FuncRead(
		js.Pointer(&fn),
	)
	return
}

// Read calls the function "WEBEXT.socket.read" directly.
func Read(socketId int32, bufferSize int32, callback js.Func[func(readInfo *ReadInfo)]) (ret js.Void) {
	bindings.CallRead(
		js.Pointer(&ret),
		int32(socketId),
		int32(bufferSize),
		callback.Ref(),
	)

	return
}

// TryRead calls the function "WEBEXT.socket.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRead(socketId int32, bufferSize int32, callback js.Func[func(readInfo *ReadInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRead(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		int32(bufferSize),
		callback.Ref(),
	)

	return
}

// HasFuncRecvFrom returns true if the function "WEBEXT.socket.recvFrom" exists.
func HasFuncRecvFrom() bool {
	return js.True == bindings.HasFuncRecvFrom()
}

// FuncRecvFrom returns the function "WEBEXT.socket.recvFrom".
func FuncRecvFrom() (fn js.Func[func(socketId int32, bufferSize int32, callback js.Func[func(recvFromInfo *RecvFromInfo)])]) {
	bindings.FuncRecvFrom(
		js.Pointer(&fn),
	)
	return
}

// RecvFrom calls the function "WEBEXT.socket.recvFrom" directly.
func RecvFrom(socketId int32, bufferSize int32, callback js.Func[func(recvFromInfo *RecvFromInfo)]) (ret js.Void) {
	bindings.CallRecvFrom(
		js.Pointer(&ret),
		int32(socketId),
		int32(bufferSize),
		callback.Ref(),
	)

	return
}

// TryRecvFrom calls the function "WEBEXT.socket.recvFrom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecvFrom(socketId int32, bufferSize int32, callback js.Func[func(recvFromInfo *RecvFromInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecvFrom(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		int32(bufferSize),
		callback.Ref(),
	)

	return
}

// HasFuncSecure returns true if the function "WEBEXT.socket.secure" exists.
func HasFuncSecure() bool {
	return js.True == bindings.HasFuncSecure()
}

// FuncSecure returns the function "WEBEXT.socket.secure".
func FuncSecure() (fn js.Func[func(socketId int32, options SecureOptions, callback js.Func[func(result int32)])]) {
	bindings.FuncSecure(
		js.Pointer(&fn),
	)
	return
}

// Secure calls the function "WEBEXT.socket.secure" directly.
func Secure(socketId int32, options SecureOptions, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSecure(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TrySecure calls the function "WEBEXT.socket.secure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySecure(socketId int32, options SecureOptions, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySecure(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncSendTo returns true if the function "WEBEXT.socket.sendTo" exists.
func HasFuncSendTo() bool {
	return js.True == bindings.HasFuncSendTo()
}

// FuncSendTo returns the function "WEBEXT.socket.sendTo".
func FuncSendTo() (fn js.Func[func(socketId int32, data js.ArrayBuffer, address js.String, port int32, callback js.Func[func(writeInfo *WriteInfo)])]) {
	bindings.FuncSendTo(
		js.Pointer(&fn),
	)
	return
}

// SendTo calls the function "WEBEXT.socket.sendTo" directly.
func SendTo(socketId int32, data js.ArrayBuffer, address js.String, port int32, callback js.Func[func(writeInfo *WriteInfo)]) (ret js.Void) {
	bindings.CallSendTo(
		js.Pointer(&ret),
		int32(socketId),
		data.Ref(),
		address.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// TrySendTo calls the function "WEBEXT.socket.sendTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendTo(socketId int32, data js.ArrayBuffer, address js.String, port int32, callback js.Func[func(writeInfo *WriteInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendTo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		data.Ref(),
		address.Ref(),
		int32(port),
		callback.Ref(),
	)

	return
}

// HasFuncSetKeepAlive returns true if the function "WEBEXT.socket.setKeepAlive" exists.
func HasFuncSetKeepAlive() bool {
	return js.True == bindings.HasFuncSetKeepAlive()
}

// FuncSetKeepAlive returns the function "WEBEXT.socket.setKeepAlive".
func FuncSetKeepAlive() (fn js.Func[func(socketId int32, enable bool, delay int32, callback js.Func[func(result bool)])]) {
	bindings.FuncSetKeepAlive(
		js.Pointer(&fn),
	)
	return
}

// SetKeepAlive calls the function "WEBEXT.socket.setKeepAlive" directly.
func SetKeepAlive(socketId int32, enable bool, delay int32, callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallSetKeepAlive(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(enable)),
		int32(delay),
		callback.Ref(),
	)

	return
}

// TrySetKeepAlive calls the function "WEBEXT.socket.setKeepAlive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetKeepAlive(socketId int32, enable bool, delay int32, callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetKeepAlive(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(enable)),
		int32(delay),
		callback.Ref(),
	)

	return
}

// HasFuncSetMulticastLoopbackMode returns true if the function "WEBEXT.socket.setMulticastLoopbackMode" exists.
func HasFuncSetMulticastLoopbackMode() bool {
	return js.True == bindings.HasFuncSetMulticastLoopbackMode()
}

// FuncSetMulticastLoopbackMode returns the function "WEBEXT.socket.setMulticastLoopbackMode".
func FuncSetMulticastLoopbackMode() (fn js.Func[func(socketId int32, enabled bool, callback js.Func[func(result int32)])]) {
	bindings.FuncSetMulticastLoopbackMode(
		js.Pointer(&fn),
	)
	return
}

// SetMulticastLoopbackMode calls the function "WEBEXT.socket.setMulticastLoopbackMode" directly.
func SetMulticastLoopbackMode(socketId int32, enabled bool, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetMulticastLoopbackMode(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(enabled)),
		callback.Ref(),
	)

	return
}

// TrySetMulticastLoopbackMode calls the function "WEBEXT.socket.setMulticastLoopbackMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMulticastLoopbackMode(socketId int32, enabled bool, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMulticastLoopbackMode(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(enabled)),
		callback.Ref(),
	)

	return
}

// HasFuncSetMulticastTimeToLive returns true if the function "WEBEXT.socket.setMulticastTimeToLive" exists.
func HasFuncSetMulticastTimeToLive() bool {
	return js.True == bindings.HasFuncSetMulticastTimeToLive()
}

// FuncSetMulticastTimeToLive returns the function "WEBEXT.socket.setMulticastTimeToLive".
func FuncSetMulticastTimeToLive() (fn js.Func[func(socketId int32, ttl int32, callback js.Func[func(result int32)])]) {
	bindings.FuncSetMulticastTimeToLive(
		js.Pointer(&fn),
	)
	return
}

// SetMulticastTimeToLive calls the function "WEBEXT.socket.setMulticastTimeToLive" directly.
func SetMulticastTimeToLive(socketId int32, ttl int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetMulticastTimeToLive(
		js.Pointer(&ret),
		int32(socketId),
		int32(ttl),
		callback.Ref(),
	)

	return
}

// TrySetMulticastTimeToLive calls the function "WEBEXT.socket.setMulticastTimeToLive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMulticastTimeToLive(socketId int32, ttl int32, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMulticastTimeToLive(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		int32(ttl),
		callback.Ref(),
	)

	return
}

// HasFuncSetNoDelay returns true if the function "WEBEXT.socket.setNoDelay" exists.
func HasFuncSetNoDelay() bool {
	return js.True == bindings.HasFuncSetNoDelay()
}

// FuncSetNoDelay returns the function "WEBEXT.socket.setNoDelay".
func FuncSetNoDelay() (fn js.Func[func(socketId int32, noDelay bool, callback js.Func[func(result bool)])]) {
	bindings.FuncSetNoDelay(
		js.Pointer(&fn),
	)
	return
}

// SetNoDelay calls the function "WEBEXT.socket.setNoDelay" directly.
func SetNoDelay(socketId int32, noDelay bool, callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallSetNoDelay(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(noDelay)),
		callback.Ref(),
	)

	return
}

// TrySetNoDelay calls the function "WEBEXT.socket.setNoDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetNoDelay(socketId int32, noDelay bool, callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetNoDelay(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(noDelay)),
		callback.Ref(),
	)

	return
}

// HasFuncWrite returns true if the function "WEBEXT.socket.write" exists.
func HasFuncWrite() bool {
	return js.True == bindings.HasFuncWrite()
}

// FuncWrite returns the function "WEBEXT.socket.write".
func FuncWrite() (fn js.Func[func(socketId int32, data js.ArrayBuffer, callback js.Func[func(writeInfo *WriteInfo)])]) {
	bindings.FuncWrite(
		js.Pointer(&fn),
	)
	return
}

// Write calls the function "WEBEXT.socket.write" directly.
func Write(socketId int32, data js.ArrayBuffer, callback js.Func[func(writeInfo *WriteInfo)]) (ret js.Void) {
	bindings.CallWrite(
		js.Pointer(&ret),
		int32(socketId),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// TryWrite calls the function "WEBEXT.socket.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryWrite(socketId int32, data js.ArrayBuffer, callback js.Func[func(writeInfo *WriteInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWrite(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		data.Ref(),
		callback.Ref(),
	)

	return
}
