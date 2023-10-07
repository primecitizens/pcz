// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tcp

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sockets/tcp/bindings"
)

type CloseCallbackFunc func(this js.Ref) js.Ref

func (fn CloseCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CloseCallbackFunc) DispatchCallback(
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

type CloseCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CloseCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CloseCallback[T]) DispatchCallback(
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

type DisconnectCallbackFunc func(this js.Ref) js.Ref

func (fn DisconnectCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisconnectCallbackFunc) DispatchCallback(
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

type DisconnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *DisconnectCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisconnectCallback[T]) DispatchCallback(
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

type DnsQueryType uint32

const (
	_ DnsQueryType = iota

	DnsQueryType_ANY
	DnsQueryType_IPV4
	DnsQueryType_IPV6
)

func (DnsQueryType) FromRef(str js.Ref) DnsQueryType {
	return DnsQueryType(bindings.ConstOfDnsQueryType(str))
}

func (x DnsQueryType) String() (string, bool) {
	switch x {
	case DnsQueryType_ANY:
		return "any", true
	case DnsQueryType_IPV4:
		return "ipv4", true
	case DnsQueryType_IPV6:
		return "ipv6", true
	default:
		return "", false
	}
}

type GetInfoCallbackFunc func(this js.Ref, socketInfo *SocketInfo) js.Ref

func (fn GetInfoCallbackFunc) Register() js.Func[func(socketInfo *SocketInfo)] {
	return js.RegisterCallback[func(socketInfo *SocketInfo)](
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
	Fn  func(arg T, this js.Ref, socketInfo *SocketInfo) js.Ref
	Arg T
}

func (cb *GetInfoCallback[T]) Register() js.Func[func(socketInfo *SocketInfo)] {
	return js.RegisterCallback[func(socketInfo *SocketInfo)](
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

type SocketInfo struct {
	// SocketId is "SocketInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// Persistent is "SocketInfo.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool
	// Name is "SocketInfo.name"
	//
	// Optional
	Name js.String
	// BufferSize is "SocketInfo.bufferSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_BufferSize MUST be set to true to make this field effective.
	BufferSize int32
	// Paused is "SocketInfo.paused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Paused MUST be set to true to make this field effective.
	Paused bool
	// Connected is "SocketInfo.connected"
	//
	// Optional
	//
	// NOTE: FFI_USE_Connected MUST be set to true to make this field effective.
	Connected bool
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

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_Persistent bool // for Persistent.
	FFI_USE_BufferSize bool // for BufferSize.
	FFI_USE_Paused     bool // for Paused.
	FFI_USE_Connected  bool // for Connected.
	FFI_USE_LocalPort  bool // for LocalPort.
	FFI_USE_PeerPort   bool // for PeerPort.

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
		p.Name.Ref(),
		p.LocalAddress.Ref(),
		p.PeerAddress.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.LocalAddress = p.LocalAddress.FromRef(js.Undefined)
	p.PeerAddress = p.PeerAddress.FromRef(js.Undefined)
}

type GetSocketsCallbackFunc func(this js.Ref, socketInfos js.Array[SocketInfo]) js.Ref

func (fn GetSocketsCallbackFunc) Register() js.Func[func(socketInfos js.Array[SocketInfo])] {
	return js.RegisterCallback[func(socketInfos js.Array[SocketInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSocketsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SocketInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSocketsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, socketInfos js.Array[SocketInfo]) js.Ref
	Arg T
}

func (cb *GetSocketsCallback[T]) Register() js.Func[func(socketInfos js.Array[SocketInfo])] {
	return js.RegisterCallback[func(socketInfos js.Array[SocketInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSocketsCallback[T]) DispatchCallback(
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

		js.Array[SocketInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReceiveErrorInfo struct {
	// SocketId is "ReceiveErrorInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// ResultCode is "ReceiveErrorInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_ResultCode bool // for ResultCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReceiveErrorInfo with all fields set.
func (p ReceiveErrorInfo) FromRef(ref js.Ref) ReceiveErrorInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReceiveErrorInfo in the application heap.
func (p ReceiveErrorInfo) New() js.Ref {
	return bindings.ReceiveErrorInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReceiveErrorInfo) UpdateFrom(ref js.Ref) {
	bindings.ReceiveErrorInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReceiveErrorInfo) Update(ref js.Ref) {
	bindings.ReceiveErrorInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReceiveErrorInfo) FreeMembers(recursive bool) {
}

type ReceiveInfo struct {
	// SocketId is "ReceiveInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// Data is "ReceiveInfo.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_SocketId bool // for SocketId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReceiveInfo with all fields set.
func (p ReceiveInfo) FromRef(ref js.Ref) ReceiveInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReceiveInfo in the application heap.
func (p ReceiveInfo) New() js.Ref {
	return bindings.ReceiveInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReceiveInfo) UpdateFrom(ref js.Ref) {
	bindings.ReceiveInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReceiveInfo) Update(ref js.Ref) {
	bindings.ReceiveInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReceiveInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
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

type SendCallbackFunc func(this js.Ref, sendInfo *SendInfo) js.Ref

func (fn SendCallbackFunc) Register() js.Func[func(sendInfo *SendInfo)] {
	return js.RegisterCallback[func(sendInfo *SendInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendInfo
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

type SendCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sendInfo *SendInfo) js.Ref
	Arg T
}

func (cb *SendCallback[T]) Register() js.Func[func(sendInfo *SendInfo)] {
	return js.RegisterCallback[func(sendInfo *SendInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SendInfo
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

type SendInfo struct {
	// ResultCode is "SendInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32
	// BytesSent is "SendInfo.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent int32

	FFI_USE_ResultCode bool // for ResultCode.
	FFI_USE_BytesSent  bool // for BytesSent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SendInfo with all fields set.
func (p SendInfo) FromRef(ref js.Ref) SendInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SendInfo in the application heap.
func (p SendInfo) New() js.Ref {
	return bindings.SendInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SendInfo) UpdateFrom(ref js.Ref) {
	bindings.SendInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SendInfo) Update(ref js.Ref) {
	bindings.SendInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SendInfo) FreeMembers(recursive bool) {
}

type SetKeepAliveCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SetKeepAliveCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetKeepAliveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SetKeepAliveCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetNoDelayCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SetNoDelayCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetNoDelayCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SetNoDelayCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
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

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetPausedCallbackFunc func(this js.Ref) js.Ref

func (fn SetPausedCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetPausedCallbackFunc) DispatchCallback(
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

type SetPausedCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetPausedCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetPausedCallback[T]) DispatchCallback(
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

type SocketProperties struct {
	// Persistent is "SocketProperties.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool
	// Name is "SocketProperties.name"
	//
	// Optional
	Name js.String
	// BufferSize is "SocketProperties.bufferSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_BufferSize MUST be set to true to make this field effective.
	BufferSize int32

	FFI_USE_Persistent bool // for Persistent.
	FFI_USE_BufferSize bool // for BufferSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SocketProperties with all fields set.
func (p SocketProperties) FromRef(ref js.Ref) SocketProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SocketProperties in the application heap.
func (p SocketProperties) New() js.Ref {
	return bindings.SocketPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SocketProperties) UpdateFrom(ref js.Ref) {
	bindings.SocketPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SocketProperties) Update(ref js.Ref) {
	bindings.SocketPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SocketProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type UpdateCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateCallbackFunc) DispatchCallback(
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

type UpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateCallback[T]) DispatchCallback(
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

// HasFuncClose returns true if the function "WEBEXT.sockets.tcp.close" exists.
func HasFuncClose() bool {
	return js.True == bindings.HasFuncClose()
}

// FuncClose returns the function "WEBEXT.sockets.tcp.close".
func FuncClose() (fn js.Func[func(socketId int32, callback js.Func[func()])]) {
	bindings.FuncClose(
		js.Pointer(&fn),
	)
	return
}

// Close calls the function "WEBEXT.sockets.tcp.close" directly.
func Close(socketId int32, callback js.Func[func()]) (ret js.Void) {
	bindings.CallClose(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryClose calls the function "WEBEXT.sockets.tcp.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClose(socketId int32, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClose(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncConnect returns true if the function "WEBEXT.sockets.tcp.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.sockets.tcp.connect".
func FuncConnect() (fn js.Func[func(socketId int32, peerAddress js.String, peerPort int32, dnsQueryType DnsQueryType, callback js.Func[func(result int32)])]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.sockets.tcp.connect" directly.
func Connect(socketId int32, peerAddress js.String, peerPort int32, dnsQueryType DnsQueryType, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallConnect(
		js.Pointer(&ret),
		int32(socketId),
		peerAddress.Ref(),
		int32(peerPort),
		uint32(dnsQueryType),
		callback.Ref(),
	)

	return
}

// TryConnect calls the function "WEBEXT.sockets.tcp.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(socketId int32, peerAddress js.String, peerPort int32, dnsQueryType DnsQueryType, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		peerAddress.Ref(),
		int32(peerPort),
		uint32(dnsQueryType),
		callback.Ref(),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.sockets.tcp.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.sockets.tcp.create".
func FuncCreate() (fn js.Func[func(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)])]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.sockets.tcp.create" directly.
func Create(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryCreate calls the function "WEBEXT.sockets.tcp.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.sockets.tcp.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.sockets.tcp.disconnect".
func FuncDisconnect() (fn js.Func[func(socketId int32, callback js.Func[func()])]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.sockets.tcp.disconnect" directly.
func Disconnect(socketId int32, callback js.Func[func()]) (ret js.Void) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.sockets.tcp.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(socketId int32, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.sockets.tcp.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.sockets.tcp.getInfo".
func FuncGetInfo() (fn js.Func[func(socketId int32, callback js.Func[func(socketInfo *SocketInfo)])]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.sockets.tcp.getInfo" directly.
func GetInfo(socketId int32, callback js.Func[func(socketInfo *SocketInfo)]) (ret js.Void) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.sockets.tcp.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo(socketId int32, callback js.Func[func(socketInfo *SocketInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// HasFuncGetSockets returns true if the function "WEBEXT.sockets.tcp.getSockets" exists.
func HasFuncGetSockets() bool {
	return js.True == bindings.HasFuncGetSockets()
}

// FuncGetSockets returns the function "WEBEXT.sockets.tcp.getSockets".
func FuncGetSockets() (fn js.Func[func(callback js.Func[func(socketInfos js.Array[SocketInfo])])]) {
	bindings.FuncGetSockets(
		js.Pointer(&fn),
	)
	return
}

// GetSockets calls the function "WEBEXT.sockets.tcp.getSockets" directly.
func GetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void) {
	bindings.CallGetSockets(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetSockets calls the function "WEBEXT.sockets.tcp.getSockets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSockets(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReceiveEventCallbackFunc func(this js.Ref, info *ReceiveInfo) js.Ref

func (fn OnReceiveEventCallbackFunc) Register() js.Func[func(info *ReceiveInfo)] {
	return js.RegisterCallback[func(info *ReceiveInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReceiveEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReceiveInfo
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

type OnReceiveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ReceiveInfo) js.Ref
	Arg T
}

func (cb *OnReceiveEventCallback[T]) Register() js.Func[func(info *ReceiveInfo)] {
	return js.RegisterCallback[func(info *ReceiveInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReceiveEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReceiveInfo
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

// HasFuncOnReceive returns true if the function "WEBEXT.sockets.tcp.onReceive.addListener" exists.
func HasFuncOnReceive() bool {
	return js.True == bindings.HasFuncOnReceive()
}

// FuncOnReceive returns the function "WEBEXT.sockets.tcp.onReceive.addListener".
func FuncOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOnReceive(
		js.Pointer(&fn),
	)
	return
}

// OnReceive calls the function "WEBEXT.sockets.tcp.onReceive.addListener" directly.
func OnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceive calls the function "WEBEXT.sockets.tcp.onReceive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceive returns true if the function "WEBEXT.sockets.tcp.onReceive.removeListener" exists.
func HasFuncOffReceive() bool {
	return js.True == bindings.HasFuncOffReceive()
}

// FuncOffReceive returns the function "WEBEXT.sockets.tcp.onReceive.removeListener".
func FuncOffReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOffReceive(
		js.Pointer(&fn),
	)
	return
}

// OffReceive calls the function "WEBEXT.sockets.tcp.onReceive.removeListener" directly.
func OffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOffReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceive calls the function "WEBEXT.sockets.tcp.onReceive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceive returns true if the function "WEBEXT.sockets.tcp.onReceive.hasListener" exists.
func HasFuncHasOnReceive() bool {
	return js.True == bindings.HasFuncHasOnReceive()
}

// FuncHasOnReceive returns the function "WEBEXT.sockets.tcp.onReceive.hasListener".
func FuncHasOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)]) bool]) {
	bindings.FuncHasOnReceive(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceive calls the function "WEBEXT.sockets.tcp.onReceive.hasListener" directly.
func HasOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret bool) {
	bindings.CallHasOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceive calls the function "WEBEXT.sockets.tcp.onReceive.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnReceiveErrorEventCallbackFunc func(this js.Ref, info *ReceiveErrorInfo) js.Ref

func (fn OnReceiveErrorEventCallbackFunc) Register() js.Func[func(info *ReceiveErrorInfo)] {
	return js.RegisterCallback[func(info *ReceiveErrorInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnReceiveErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReceiveErrorInfo
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

type OnReceiveErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *ReceiveErrorInfo) js.Ref
	Arg T
}

func (cb *OnReceiveErrorEventCallback[T]) Register() js.Func[func(info *ReceiveErrorInfo)] {
	return js.RegisterCallback[func(info *ReceiveErrorInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnReceiveErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ReceiveErrorInfo
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

// HasFuncOnReceiveError returns true if the function "WEBEXT.sockets.tcp.onReceiveError.addListener" exists.
func HasFuncOnReceiveError() bool {
	return js.True == bindings.HasFuncOnReceiveError()
}

// FuncOnReceiveError returns the function "WEBEXT.sockets.tcp.onReceiveError.addListener".
func FuncOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OnReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.addListener" directly.
func OnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceiveError returns true if the function "WEBEXT.sockets.tcp.onReceiveError.removeListener" exists.
func HasFuncOffReceiveError() bool {
	return js.True == bindings.HasFuncOffReceiveError()
}

// FuncOffReceiveError returns the function "WEBEXT.sockets.tcp.onReceiveError.removeListener".
func FuncOffReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOffReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OffReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.removeListener" directly.
func OffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOffReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceiveError returns true if the function "WEBEXT.sockets.tcp.onReceiveError.hasListener" exists.
func HasFuncHasOnReceiveError() bool {
	return js.True == bindings.HasFuncHasOnReceiveError()
}

// FuncHasOnReceiveError returns the function "WEBEXT.sockets.tcp.onReceiveError.hasListener".
func FuncHasOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)]) bool]) {
	bindings.FuncHasOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.hasListener" directly.
func HasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool) {
	bindings.CallHasOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceiveError calls the function "WEBEXT.sockets.tcp.onReceiveError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSecure returns true if the function "WEBEXT.sockets.tcp.secure" exists.
func HasFuncSecure() bool {
	return js.True == bindings.HasFuncSecure()
}

// FuncSecure returns the function "WEBEXT.sockets.tcp.secure".
func FuncSecure() (fn js.Func[func(socketId int32, options SecureOptions, callback js.Func[func(result int32)])]) {
	bindings.FuncSecure(
		js.Pointer(&fn),
	)
	return
}

// Secure calls the function "WEBEXT.sockets.tcp.secure" directly.
func Secure(socketId int32, options SecureOptions, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSecure(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TrySecure calls the function "WEBEXT.sockets.tcp.secure"
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

// HasFuncSend returns true if the function "WEBEXT.sockets.tcp.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.sockets.tcp.send".
func FuncSend() (fn js.Func[func(socketId int32, data js.ArrayBuffer, callback js.Func[func(sendInfo *SendInfo)])]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.sockets.tcp.send" directly.
func Send(socketId int32, data js.ArrayBuffer, callback js.Func[func(sendInfo *SendInfo)]) (ret js.Void) {
	bindings.CallSend(
		js.Pointer(&ret),
		int32(socketId),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// TrySend calls the function "WEBEXT.sockets.tcp.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(socketId int32, data js.ArrayBuffer, callback js.Func[func(sendInfo *SendInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncSetKeepAlive returns true if the function "WEBEXT.sockets.tcp.setKeepAlive" exists.
func HasFuncSetKeepAlive() bool {
	return js.True == bindings.HasFuncSetKeepAlive()
}

// FuncSetKeepAlive returns the function "WEBEXT.sockets.tcp.setKeepAlive".
func FuncSetKeepAlive() (fn js.Func[func(socketId int32, enable bool, delay int32, callback js.Func[func(result int32)])]) {
	bindings.FuncSetKeepAlive(
		js.Pointer(&fn),
	)
	return
}

// SetKeepAlive calls the function "WEBEXT.sockets.tcp.setKeepAlive" directly.
func SetKeepAlive(socketId int32, enable bool, delay int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetKeepAlive(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(enable)),
		int32(delay),
		callback.Ref(),
	)

	return
}

// TrySetKeepAlive calls the function "WEBEXT.sockets.tcp.setKeepAlive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetKeepAlive(socketId int32, enable bool, delay int32, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetKeepAlive(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(enable)),
		int32(delay),
		callback.Ref(),
	)

	return
}

// HasFuncSetNoDelay returns true if the function "WEBEXT.sockets.tcp.setNoDelay" exists.
func HasFuncSetNoDelay() bool {
	return js.True == bindings.HasFuncSetNoDelay()
}

// FuncSetNoDelay returns the function "WEBEXT.sockets.tcp.setNoDelay".
func FuncSetNoDelay() (fn js.Func[func(socketId int32, noDelay bool, callback js.Func[func(result int32)])]) {
	bindings.FuncSetNoDelay(
		js.Pointer(&fn),
	)
	return
}

// SetNoDelay calls the function "WEBEXT.sockets.tcp.setNoDelay" directly.
func SetNoDelay(socketId int32, noDelay bool, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetNoDelay(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(noDelay)),
		callback.Ref(),
	)

	return
}

// TrySetNoDelay calls the function "WEBEXT.sockets.tcp.setNoDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetNoDelay(socketId int32, noDelay bool, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetNoDelay(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(noDelay)),
		callback.Ref(),
	)

	return
}

// HasFuncSetPaused returns true if the function "WEBEXT.sockets.tcp.setPaused" exists.
func HasFuncSetPaused() bool {
	return js.True == bindings.HasFuncSetPaused()
}

// FuncSetPaused returns the function "WEBEXT.sockets.tcp.setPaused".
func FuncSetPaused() (fn js.Func[func(socketId int32, paused bool, callback js.Func[func()])]) {
	bindings.FuncSetPaused(
		js.Pointer(&fn),
	)
	return
}

// SetPaused calls the function "WEBEXT.sockets.tcp.setPaused" directly.
func SetPaused(socketId int32, paused bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetPaused(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(paused)),
		callback.Ref(),
	)

	return
}

// TrySetPaused calls the function "WEBEXT.sockets.tcp.setPaused"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPaused(socketId int32, paused bool, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPaused(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(paused)),
		callback.Ref(),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.sockets.tcp.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.sockets.tcp.update".
func FuncUpdate() (fn js.Func[func(socketId int32, properties SocketProperties, callback js.Func[func()])]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.sockets.tcp.update" directly.
func Update(socketId int32, properties SocketProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryUpdate calls the function "WEBEXT.sockets.tcp.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(socketId int32, properties SocketProperties, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}
