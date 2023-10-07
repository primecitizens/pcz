// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package udp

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sockets/udp/bindings"
)

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

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_Persistent bool // for Persistent.
	FFI_USE_BufferSize bool // for BufferSize.
	FFI_USE_Paused     bool // for Paused.
	FFI_USE_LocalPort  bool // for LocalPort.

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
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
	// RemoteAddress is "ReceiveInfo.remoteAddress"
	//
	// Optional
	RemoteAddress js.String
	// RemotePort is "ReceiveInfo.remotePort"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemotePort MUST be set to true to make this field effective.
	RemotePort int32

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_RemotePort bool // for RemotePort.

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
		p.RemoteAddress.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.RemoteAddress = p.RemoteAddress.FromRef(js.Undefined)
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

type SetBroadcastCallbackFunc func(this js.Ref, result int32) js.Ref

func (fn SetBroadcastCallbackFunc) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetBroadcastCallbackFunc) DispatchCallback(
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

type SetBroadcastCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result int32) js.Ref
	Arg T
}

func (cb *SetBroadcastCallback[T]) Register() js.Func[func(result int32)] {
	return js.RegisterCallback[func(result int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetBroadcastCallback[T]) DispatchCallback(
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

// HasFuncBind returns true if the function "WEBEXT.sockets.udp.bind" exists.
func HasFuncBind() bool {
	return js.True == bindings.HasFuncBind()
}

// FuncBind returns the function "WEBEXT.sockets.udp.bind".
func FuncBind() (fn js.Func[func(socketId int32, address js.String, port int32, callback js.Func[func(result int32)])]) {
	bindings.FuncBind(
		js.Pointer(&fn),
	)
	return
}

// Bind calls the function "WEBEXT.sockets.udp.bind" directly.
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

// TryBind calls the function "WEBEXT.sockets.udp.bind"
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

// HasFuncClose returns true if the function "WEBEXT.sockets.udp.close" exists.
func HasFuncClose() bool {
	return js.True == bindings.HasFuncClose()
}

// FuncClose returns the function "WEBEXT.sockets.udp.close".
func FuncClose() (fn js.Func[func(socketId int32, callback js.Func[func()])]) {
	bindings.FuncClose(
		js.Pointer(&fn),
	)
	return
}

// Close calls the function "WEBEXT.sockets.udp.close" directly.
func Close(socketId int32, callback js.Func[func()]) (ret js.Void) {
	bindings.CallClose(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryClose calls the function "WEBEXT.sockets.udp.close"
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

// HasFuncCreate returns true if the function "WEBEXT.sockets.udp.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.sockets.udp.create".
func FuncCreate() (fn js.Func[func(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)])]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.sockets.udp.create" directly.
func Create(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryCreate calls the function "WEBEXT.sockets.udp.create"
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

// HasFuncGetInfo returns true if the function "WEBEXT.sockets.udp.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.sockets.udp.getInfo".
func FuncGetInfo() (fn js.Func[func(socketId int32, callback js.Func[func(socketInfo *SocketInfo)])]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.sockets.udp.getInfo" directly.
func GetInfo(socketId int32, callback js.Func[func(socketInfo *SocketInfo)]) (ret js.Void) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.sockets.udp.getInfo"
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

// HasFuncGetJoinedGroups returns true if the function "WEBEXT.sockets.udp.getJoinedGroups" exists.
func HasFuncGetJoinedGroups() bool {
	return js.True == bindings.HasFuncGetJoinedGroups()
}

// FuncGetJoinedGroups returns the function "WEBEXT.sockets.udp.getJoinedGroups".
func FuncGetJoinedGroups() (fn js.Func[func(socketId int32, callback js.Func[func(groups js.Array[js.String])])]) {
	bindings.FuncGetJoinedGroups(
		js.Pointer(&fn),
	)
	return
}

// GetJoinedGroups calls the function "WEBEXT.sockets.udp.getJoinedGroups" directly.
func GetJoinedGroups(socketId int32, callback js.Func[func(groups js.Array[js.String])]) (ret js.Void) {
	bindings.CallGetJoinedGroups(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetJoinedGroups calls the function "WEBEXT.sockets.udp.getJoinedGroups"
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

// HasFuncGetSockets returns true if the function "WEBEXT.sockets.udp.getSockets" exists.
func HasFuncGetSockets() bool {
	return js.True == bindings.HasFuncGetSockets()
}

// FuncGetSockets returns the function "WEBEXT.sockets.udp.getSockets".
func FuncGetSockets() (fn js.Func[func(callback js.Func[func(socketInfos js.Array[SocketInfo])])]) {
	bindings.FuncGetSockets(
		js.Pointer(&fn),
	)
	return
}

// GetSockets calls the function "WEBEXT.sockets.udp.getSockets" directly.
func GetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void) {
	bindings.CallGetSockets(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetSockets calls the function "WEBEXT.sockets.udp.getSockets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSockets(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncJoinGroup returns true if the function "WEBEXT.sockets.udp.joinGroup" exists.
func HasFuncJoinGroup() bool {
	return js.True == bindings.HasFuncJoinGroup()
}

// FuncJoinGroup returns the function "WEBEXT.sockets.udp.joinGroup".
func FuncJoinGroup() (fn js.Func[func(socketId int32, address js.String, callback js.Func[func(result int32)])]) {
	bindings.FuncJoinGroup(
		js.Pointer(&fn),
	)
	return
}

// JoinGroup calls the function "WEBEXT.sockets.udp.joinGroup" directly.
func JoinGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallJoinGroup(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// TryJoinGroup calls the function "WEBEXT.sockets.udp.joinGroup"
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

// HasFuncLeaveGroup returns true if the function "WEBEXT.sockets.udp.leaveGroup" exists.
func HasFuncLeaveGroup() bool {
	return js.True == bindings.HasFuncLeaveGroup()
}

// FuncLeaveGroup returns the function "WEBEXT.sockets.udp.leaveGroup".
func FuncLeaveGroup() (fn js.Func[func(socketId int32, address js.String, callback js.Func[func(result int32)])]) {
	bindings.FuncLeaveGroup(
		js.Pointer(&fn),
	)
	return
}

// LeaveGroup calls the function "WEBEXT.sockets.udp.leaveGroup" directly.
func LeaveGroup(socketId int32, address js.String, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallLeaveGroup(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		callback.Ref(),
	)

	return
}

// TryLeaveGroup calls the function "WEBEXT.sockets.udp.leaveGroup"
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

// HasFuncOnReceive returns true if the function "WEBEXT.sockets.udp.onReceive.addListener" exists.
func HasFuncOnReceive() bool {
	return js.True == bindings.HasFuncOnReceive()
}

// FuncOnReceive returns the function "WEBEXT.sockets.udp.onReceive.addListener".
func FuncOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOnReceive(
		js.Pointer(&fn),
	)
	return
}

// OnReceive calls the function "WEBEXT.sockets.udp.onReceive.addListener" directly.
func OnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceive calls the function "WEBEXT.sockets.udp.onReceive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceive returns true if the function "WEBEXT.sockets.udp.onReceive.removeListener" exists.
func HasFuncOffReceive() bool {
	return js.True == bindings.HasFuncOffReceive()
}

// FuncOffReceive returns the function "WEBEXT.sockets.udp.onReceive.removeListener".
func FuncOffReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOffReceive(
		js.Pointer(&fn),
	)
	return
}

// OffReceive calls the function "WEBEXT.sockets.udp.onReceive.removeListener" directly.
func OffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOffReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceive calls the function "WEBEXT.sockets.udp.onReceive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceive returns true if the function "WEBEXT.sockets.udp.onReceive.hasListener" exists.
func HasFuncHasOnReceive() bool {
	return js.True == bindings.HasFuncHasOnReceive()
}

// FuncHasOnReceive returns the function "WEBEXT.sockets.udp.onReceive.hasListener".
func FuncHasOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)]) bool]) {
	bindings.FuncHasOnReceive(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceive calls the function "WEBEXT.sockets.udp.onReceive.hasListener" directly.
func HasOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret bool) {
	bindings.CallHasOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceive calls the function "WEBEXT.sockets.udp.onReceive.hasListener"
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

// HasFuncOnReceiveError returns true if the function "WEBEXT.sockets.udp.onReceiveError.addListener" exists.
func HasFuncOnReceiveError() bool {
	return js.True == bindings.HasFuncOnReceiveError()
}

// FuncOnReceiveError returns the function "WEBEXT.sockets.udp.onReceiveError.addListener".
func FuncOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OnReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.addListener" directly.
func OnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceiveError returns true if the function "WEBEXT.sockets.udp.onReceiveError.removeListener" exists.
func HasFuncOffReceiveError() bool {
	return js.True == bindings.HasFuncOffReceiveError()
}

// FuncOffReceiveError returns the function "WEBEXT.sockets.udp.onReceiveError.removeListener".
func FuncOffReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOffReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OffReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.removeListener" directly.
func OffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOffReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceiveError returns true if the function "WEBEXT.sockets.udp.onReceiveError.hasListener" exists.
func HasFuncHasOnReceiveError() bool {
	return js.True == bindings.HasFuncHasOnReceiveError()
}

// FuncHasOnReceiveError returns the function "WEBEXT.sockets.udp.onReceiveError.hasListener".
func FuncHasOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)]) bool]) {
	bindings.FuncHasOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.hasListener" directly.
func HasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool) {
	bindings.CallHasOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceiveError calls the function "WEBEXT.sockets.udp.onReceiveError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSend returns true if the function "WEBEXT.sockets.udp.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.sockets.udp.send".
func FuncSend() (fn js.Func[func(socketId int32, data js.ArrayBuffer, address js.String, port int32, dnsQueryType DnsQueryType, callback js.Func[func(sendInfo *SendInfo)])]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.sockets.udp.send" directly.
func Send(socketId int32, data js.ArrayBuffer, address js.String, port int32, dnsQueryType DnsQueryType, callback js.Func[func(sendInfo *SendInfo)]) (ret js.Void) {
	bindings.CallSend(
		js.Pointer(&ret),
		int32(socketId),
		data.Ref(),
		address.Ref(),
		int32(port),
		uint32(dnsQueryType),
		callback.Ref(),
	)

	return
}

// TrySend calls the function "WEBEXT.sockets.udp.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(socketId int32, data js.ArrayBuffer, address js.String, port int32, dnsQueryType DnsQueryType, callback js.Func[func(sendInfo *SendInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		data.Ref(),
		address.Ref(),
		int32(port),
		uint32(dnsQueryType),
		callback.Ref(),
	)

	return
}

// HasFuncSetBroadcast returns true if the function "WEBEXT.sockets.udp.setBroadcast" exists.
func HasFuncSetBroadcast() bool {
	return js.True == bindings.HasFuncSetBroadcast()
}

// FuncSetBroadcast returns the function "WEBEXT.sockets.udp.setBroadcast".
func FuncSetBroadcast() (fn js.Func[func(socketId int32, enabled bool, callback js.Func[func(result int32)])]) {
	bindings.FuncSetBroadcast(
		js.Pointer(&fn),
	)
	return
}

// SetBroadcast calls the function "WEBEXT.sockets.udp.setBroadcast" directly.
func SetBroadcast(socketId int32, enabled bool, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetBroadcast(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(enabled)),
		callback.Ref(),
	)

	return
}

// TrySetBroadcast calls the function "WEBEXT.sockets.udp.setBroadcast"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetBroadcast(socketId int32, enabled bool, callback js.Func[func(result int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetBroadcast(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(enabled)),
		callback.Ref(),
	)

	return
}

// HasFuncSetMulticastLoopbackMode returns true if the function "WEBEXT.sockets.udp.setMulticastLoopbackMode" exists.
func HasFuncSetMulticastLoopbackMode() bool {
	return js.True == bindings.HasFuncSetMulticastLoopbackMode()
}

// FuncSetMulticastLoopbackMode returns the function "WEBEXT.sockets.udp.setMulticastLoopbackMode".
func FuncSetMulticastLoopbackMode() (fn js.Func[func(socketId int32, enabled bool, callback js.Func[func(result int32)])]) {
	bindings.FuncSetMulticastLoopbackMode(
		js.Pointer(&fn),
	)
	return
}

// SetMulticastLoopbackMode calls the function "WEBEXT.sockets.udp.setMulticastLoopbackMode" directly.
func SetMulticastLoopbackMode(socketId int32, enabled bool, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetMulticastLoopbackMode(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(enabled)),
		callback.Ref(),
	)

	return
}

// TrySetMulticastLoopbackMode calls the function "WEBEXT.sockets.udp.setMulticastLoopbackMode"
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

// HasFuncSetMulticastTimeToLive returns true if the function "WEBEXT.sockets.udp.setMulticastTimeToLive" exists.
func HasFuncSetMulticastTimeToLive() bool {
	return js.True == bindings.HasFuncSetMulticastTimeToLive()
}

// FuncSetMulticastTimeToLive returns the function "WEBEXT.sockets.udp.setMulticastTimeToLive".
func FuncSetMulticastTimeToLive() (fn js.Func[func(socketId int32, ttl int32, callback js.Func[func(result int32)])]) {
	bindings.FuncSetMulticastTimeToLive(
		js.Pointer(&fn),
	)
	return
}

// SetMulticastTimeToLive calls the function "WEBEXT.sockets.udp.setMulticastTimeToLive" directly.
func SetMulticastTimeToLive(socketId int32, ttl int32, callback js.Func[func(result int32)]) (ret js.Void) {
	bindings.CallSetMulticastTimeToLive(
		js.Pointer(&ret),
		int32(socketId),
		int32(ttl),
		callback.Ref(),
	)

	return
}

// TrySetMulticastTimeToLive calls the function "WEBEXT.sockets.udp.setMulticastTimeToLive"
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

// HasFuncSetPaused returns true if the function "WEBEXT.sockets.udp.setPaused" exists.
func HasFuncSetPaused() bool {
	return js.True == bindings.HasFuncSetPaused()
}

// FuncSetPaused returns the function "WEBEXT.sockets.udp.setPaused".
func FuncSetPaused() (fn js.Func[func(socketId int32, paused bool, callback js.Func[func()])]) {
	bindings.FuncSetPaused(
		js.Pointer(&fn),
	)
	return
}

// SetPaused calls the function "WEBEXT.sockets.udp.setPaused" directly.
func SetPaused(socketId int32, paused bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetPaused(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(paused)),
		callback.Ref(),
	)

	return
}

// TrySetPaused calls the function "WEBEXT.sockets.udp.setPaused"
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

// HasFuncUpdate returns true if the function "WEBEXT.sockets.udp.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.sockets.udp.update".
func FuncUpdate() (fn js.Func[func(socketId int32, properties SocketProperties, callback js.Func[func()])]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.sockets.udp.update" directly.
func Update(socketId int32, properties SocketProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryUpdate calls the function "WEBEXT.sockets.udp.update"
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
