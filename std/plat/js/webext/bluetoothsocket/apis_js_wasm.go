// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bluetoothsocket

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bluetoothsocket/bindings"
)

type AcceptError uint32

const (
	_ AcceptError = iota

	AcceptError_SYSTEM_ERROR
	AcceptError_NOT_LISTENING
)

func (AcceptError) FromRef(str js.Ref) AcceptError {
	return AcceptError(bindings.ConstOfAcceptError(str))
}

func (x AcceptError) String() (string, bool) {
	switch x {
	case AcceptError_SYSTEM_ERROR:
		return "system_error", true
	case AcceptError_NOT_LISTENING:
		return "not_listening", true
	default:
		return "", false
	}
}

type AcceptErrorInfo struct {
	// SocketId is "AcceptErrorInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// ErrorMessage is "AcceptErrorInfo.errorMessage"
	//
	// Optional
	ErrorMessage js.String
	// Error is "AcceptErrorInfo.error"
	//
	// Optional
	Error AcceptError

	FFI_USE_SocketId bool // for SocketId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AcceptErrorInfo with all fields set.
func (p AcceptErrorInfo) FromRef(ref js.Ref) AcceptErrorInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AcceptErrorInfo in the application heap.
func (p AcceptErrorInfo) New() js.Ref {
	return bindings.AcceptErrorInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AcceptErrorInfo) UpdateFrom(ref js.Ref) {
	bindings.AcceptErrorInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AcceptErrorInfo) Update(ref js.Ref) {
	bindings.AcceptErrorInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AcceptErrorInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ErrorMessage.Ref(),
	)
	p.ErrorMessage = p.ErrorMessage.FromRef(js.Undefined)
}

type AcceptInfo struct {
	// SocketId is "AcceptInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// ClientSocketId is "AcceptInfo.clientSocketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientSocketId MUST be set to true to make this field effective.
	ClientSocketId int32

	FFI_USE_SocketId       bool // for SocketId.
	FFI_USE_ClientSocketId bool // for ClientSocketId.

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

type ConnectCallbackFunc func(this js.Ref) js.Ref

func (fn ConnectCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ConnectCallbackFunc) DispatchCallback(
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

type ConnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ConnectCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ConnectCallback[T]) DispatchCallback(
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
	// Address is "SocketInfo.address"
	//
	// Optional
	Address js.String
	// Uuid is "SocketInfo.uuid"
	//
	// Optional
	Uuid js.String

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_Persistent bool // for Persistent.
	FFI_USE_BufferSize bool // for BufferSize.
	FFI_USE_Paused     bool // for Paused.
	FFI_USE_Connected  bool // for Connected.

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
		p.Address.Ref(),
		p.Uuid.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Address = p.Address.FromRef(js.Undefined)
	p.Uuid = p.Uuid.FromRef(js.Undefined)
}

type GetSocketsCallbackFunc func(this js.Ref, sockets js.Array[SocketInfo]) js.Ref

func (fn GetSocketsCallbackFunc) Register() js.Func[func(sockets js.Array[SocketInfo])] {
	return js.RegisterCallback[func(sockets js.Array[SocketInfo])](
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
	Fn  func(arg T, this js.Ref, sockets js.Array[SocketInfo]) js.Ref
	Arg T
}

func (cb *GetSocketsCallback[T]) Register() js.Func[func(sockets js.Array[SocketInfo])] {
	return js.RegisterCallback[func(sockets js.Array[SocketInfo])](
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

type ListenCallbackFunc func(this js.Ref) js.Ref

func (fn ListenCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ListenCallbackFunc) DispatchCallback(
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

type ListenCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ListenCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ListenCallback[T]) DispatchCallback(
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

type ListenOptions struct {
	// Channel is "ListenOptions.channel"
	//
	// Optional
	//
	// NOTE: FFI_USE_Channel MUST be set to true to make this field effective.
	Channel int32
	// Psm is "ListenOptions.psm"
	//
	// Optional
	//
	// NOTE: FFI_USE_Psm MUST be set to true to make this field effective.
	Psm int32
	// Backlog is "ListenOptions.backlog"
	//
	// Optional
	//
	// NOTE: FFI_USE_Backlog MUST be set to true to make this field effective.
	Backlog int32

	FFI_USE_Channel bool // for Channel.
	FFI_USE_Psm     bool // for Psm.
	FFI_USE_Backlog bool // for Backlog.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ListenOptions with all fields set.
func (p ListenOptions) FromRef(ref js.Ref) ListenOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ListenOptions in the application heap.
func (p ListenOptions) New() js.Ref {
	return bindings.ListenOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ListenOptions) UpdateFrom(ref js.Ref) {
	bindings.ListenOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ListenOptions) Update(ref js.Ref) {
	bindings.ListenOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ListenOptions) FreeMembers(recursive bool) {
}

type ReceiveError uint32

const (
	_ ReceiveError = iota

	ReceiveError_DISCONNECTED
	ReceiveError_SYSTEM_ERROR
	ReceiveError_NOT_CONNECTED
)

func (ReceiveError) FromRef(str js.Ref) ReceiveError {
	return ReceiveError(bindings.ConstOfReceiveError(str))
}

func (x ReceiveError) String() (string, bool) {
	switch x {
	case ReceiveError_DISCONNECTED:
		return "disconnected", true
	case ReceiveError_SYSTEM_ERROR:
		return "system_error", true
	case ReceiveError_NOT_CONNECTED:
		return "not_connected", true
	default:
		return "", false
	}
}

type ReceiveErrorInfo struct {
	// SocketId is "ReceiveErrorInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// ErrorMessage is "ReceiveErrorInfo.errorMessage"
	//
	// Optional
	ErrorMessage js.String
	// Error is "ReceiveErrorInfo.error"
	//
	// Optional
	Error ReceiveError

	FFI_USE_SocketId bool // for SocketId.

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
	js.Free(
		p.ErrorMessage.Ref(),
	)
	p.ErrorMessage = p.ErrorMessage.FromRef(js.Undefined)
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

type SendCallbackFunc func(this js.Ref, bytesSent int32) js.Ref

func (fn SendCallbackFunc) Register() js.Func[func(bytesSent int32)] {
	return js.RegisterCallback[func(bytesSent int32)](
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

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SendCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bytesSent int32) js.Ref
	Arg T
}

func (cb *SendCallback[T]) Register() js.Func[func(bytesSent int32)] {
	return js.RegisterCallback[func(bytesSent int32)](
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

// HasFuncClose returns true if the function "WEBEXT.bluetoothSocket.close" exists.
func HasFuncClose() bool {
	return js.True == bindings.HasFuncClose()
}

// FuncClose returns the function "WEBEXT.bluetoothSocket.close".
func FuncClose() (fn js.Func[func(socketId int32) js.Promise[js.Void]]) {
	bindings.FuncClose(
		js.Pointer(&fn),
	)
	return
}

// Close calls the function "WEBEXT.bluetoothSocket.close" directly.
func Close(socketId int32) (ret js.Promise[js.Void]) {
	bindings.CallClose(
		js.Pointer(&ret),
		int32(socketId),
	)

	return
}

// TryClose calls the function "WEBEXT.bluetoothSocket.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClose(socketId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClose(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
	)

	return
}

// HasFuncConnect returns true if the function "WEBEXT.bluetoothSocket.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.bluetoothSocket.connect".
func FuncConnect() (fn js.Func[func(socketId int32, address js.String, uuid js.String) js.Promise[js.Void]]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.bluetoothSocket.connect" directly.
func Connect(socketId int32, address js.String, uuid js.String) (ret js.Promise[js.Void]) {
	bindings.CallConnect(
		js.Pointer(&ret),
		int32(socketId),
		address.Ref(),
		uuid.Ref(),
	)

	return
}

// TryConnect calls the function "WEBEXT.bluetoothSocket.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(socketId int32, address js.String, uuid js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		address.Ref(),
		uuid.Ref(),
	)

	return
}

// HasFuncCreate returns true if the function "WEBEXT.bluetoothSocket.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.bluetoothSocket.create".
func FuncCreate() (fn js.Func[func(properties SocketProperties) js.Promise[CreateInfo]]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.bluetoothSocket.create" directly.
func Create(properties SocketProperties) (ret js.Promise[CreateInfo]) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&properties),
	)

	return
}

// TryCreate calls the function "WEBEXT.bluetoothSocket.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreate(properties SocketProperties) (ret js.Promise[CreateInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreate(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&properties),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.bluetoothSocket.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.bluetoothSocket.disconnect".
func FuncDisconnect() (fn js.Func[func(socketId int32) js.Promise[js.Void]]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.bluetoothSocket.disconnect" directly.
func Disconnect(socketId int32) (ret js.Promise[js.Void]) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(socketId),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.bluetoothSocket.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(socketId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.bluetoothSocket.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.bluetoothSocket.getInfo".
func FuncGetInfo() (fn js.Func[func(socketId int32) js.Promise[SocketInfo]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.bluetoothSocket.getInfo" directly.
func GetInfo(socketId int32) (ret js.Promise[SocketInfo]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(socketId),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.bluetoothSocket.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo(socketId int32) (ret js.Promise[SocketInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
	)

	return
}

// HasFuncGetSockets returns true if the function "WEBEXT.bluetoothSocket.getSockets" exists.
func HasFuncGetSockets() bool {
	return js.True == bindings.HasFuncGetSockets()
}

// FuncGetSockets returns the function "WEBEXT.bluetoothSocket.getSockets".
func FuncGetSockets() (fn js.Func[func() js.Promise[js.Array[SocketInfo]]]) {
	bindings.FuncGetSockets(
		js.Pointer(&fn),
	)
	return
}

// GetSockets calls the function "WEBEXT.bluetoothSocket.getSockets" directly.
func GetSockets() (ret js.Promise[js.Array[SocketInfo]]) {
	bindings.CallGetSockets(
		js.Pointer(&ret),
	)

	return
}

// TryGetSockets calls the function "WEBEXT.bluetoothSocket.getSockets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSockets() (ret js.Promise[js.Array[SocketInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSockets(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncListenUsingL2cap returns true if the function "WEBEXT.bluetoothSocket.listenUsingL2cap" exists.
func HasFuncListenUsingL2cap() bool {
	return js.True == bindings.HasFuncListenUsingL2cap()
}

// FuncListenUsingL2cap returns the function "WEBEXT.bluetoothSocket.listenUsingL2cap".
func FuncListenUsingL2cap() (fn js.Func[func(socketId int32, uuid js.String, options ListenOptions) js.Promise[js.Void]]) {
	bindings.FuncListenUsingL2cap(
		js.Pointer(&fn),
	)
	return
}

// ListenUsingL2cap calls the function "WEBEXT.bluetoothSocket.listenUsingL2cap" directly.
func ListenUsingL2cap(socketId int32, uuid js.String, options ListenOptions) (ret js.Promise[js.Void]) {
	bindings.CallListenUsingL2cap(
		js.Pointer(&ret),
		int32(socketId),
		uuid.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryListenUsingL2cap calls the function "WEBEXT.bluetoothSocket.listenUsingL2cap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListenUsingL2cap(socketId int32, uuid js.String, options ListenOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryListenUsingL2cap(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		uuid.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncListenUsingRfcomm returns true if the function "WEBEXT.bluetoothSocket.listenUsingRfcomm" exists.
func HasFuncListenUsingRfcomm() bool {
	return js.True == bindings.HasFuncListenUsingRfcomm()
}

// FuncListenUsingRfcomm returns the function "WEBEXT.bluetoothSocket.listenUsingRfcomm".
func FuncListenUsingRfcomm() (fn js.Func[func(socketId int32, uuid js.String, options ListenOptions) js.Promise[js.Void]]) {
	bindings.FuncListenUsingRfcomm(
		js.Pointer(&fn),
	)
	return
}

// ListenUsingRfcomm calls the function "WEBEXT.bluetoothSocket.listenUsingRfcomm" directly.
func ListenUsingRfcomm(socketId int32, uuid js.String, options ListenOptions) (ret js.Promise[js.Void]) {
	bindings.CallListenUsingRfcomm(
		js.Pointer(&ret),
		int32(socketId),
		uuid.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryListenUsingRfcomm calls the function "WEBEXT.bluetoothSocket.listenUsingRfcomm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryListenUsingRfcomm(socketId int32, uuid js.String, options ListenOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryListenUsingRfcomm(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		uuid.Ref(),
		js.Pointer(&options),
	)

	return
}

type OnAcceptEventCallbackFunc func(this js.Ref, info *AcceptInfo) js.Ref

func (fn OnAcceptEventCallbackFunc) Register() js.Func[func(info *AcceptInfo)] {
	return js.RegisterCallback[func(info *AcceptInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAcceptEventCallbackFunc) DispatchCallback(
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

type OnAcceptEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *AcceptInfo) js.Ref
	Arg T
}

func (cb *OnAcceptEventCallback[T]) Register() js.Func[func(info *AcceptInfo)] {
	return js.RegisterCallback[func(info *AcceptInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAcceptEventCallback[T]) DispatchCallback(
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

// HasFuncOnAccept returns true if the function "WEBEXT.bluetoothSocket.onAccept.addListener" exists.
func HasFuncOnAccept() bool {
	return js.True == bindings.HasFuncOnAccept()
}

// FuncOnAccept returns the function "WEBEXT.bluetoothSocket.onAccept.addListener".
func FuncOnAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)])]) {
	bindings.FuncOnAccept(
		js.Pointer(&fn),
	)
	return
}

// OnAccept calls the function "WEBEXT.bluetoothSocket.onAccept.addListener" directly.
func OnAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void) {
	bindings.CallOnAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccept calls the function "WEBEXT.bluetoothSocket.onAccept.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccept returns true if the function "WEBEXT.bluetoothSocket.onAccept.removeListener" exists.
func HasFuncOffAccept() bool {
	return js.True == bindings.HasFuncOffAccept()
}

// FuncOffAccept returns the function "WEBEXT.bluetoothSocket.onAccept.removeListener".
func FuncOffAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)])]) {
	bindings.FuncOffAccept(
		js.Pointer(&fn),
	)
	return
}

// OffAccept calls the function "WEBEXT.bluetoothSocket.onAccept.removeListener" directly.
func OffAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void) {
	bindings.CallOffAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccept calls the function "WEBEXT.bluetoothSocket.onAccept.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccept returns true if the function "WEBEXT.bluetoothSocket.onAccept.hasListener" exists.
func HasFuncHasOnAccept() bool {
	return js.True == bindings.HasFuncHasOnAccept()
}

// FuncHasOnAccept returns the function "WEBEXT.bluetoothSocket.onAccept.hasListener".
func FuncHasOnAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)]) bool]) {
	bindings.FuncHasOnAccept(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccept calls the function "WEBEXT.bluetoothSocket.onAccept.hasListener" directly.
func HasOnAccept(callback js.Func[func(info *AcceptInfo)]) (ret bool) {
	bindings.CallHasOnAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccept calls the function "WEBEXT.bluetoothSocket.onAccept.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccept(callback js.Func[func(info *AcceptInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnAcceptErrorEventCallbackFunc func(this js.Ref, info *AcceptErrorInfo) js.Ref

func (fn OnAcceptErrorEventCallbackFunc) Register() js.Func[func(info *AcceptErrorInfo)] {
	return js.RegisterCallback[func(info *AcceptErrorInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAcceptErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AcceptErrorInfo
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

type OnAcceptErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *AcceptErrorInfo) js.Ref
	Arg T
}

func (cb *OnAcceptErrorEventCallback[T]) Register() js.Func[func(info *AcceptErrorInfo)] {
	return js.RegisterCallback[func(info *AcceptErrorInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAcceptErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AcceptErrorInfo
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

// HasFuncOnAcceptError returns true if the function "WEBEXT.bluetoothSocket.onAcceptError.addListener" exists.
func HasFuncOnAcceptError() bool {
	return js.True == bindings.HasFuncOnAcceptError()
}

// FuncOnAcceptError returns the function "WEBEXT.bluetoothSocket.onAcceptError.addListener".
func FuncOnAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)])]) {
	bindings.FuncOnAcceptError(
		js.Pointer(&fn),
	)
	return
}

// OnAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.addListener" directly.
func OnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void) {
	bindings.CallOnAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAcceptError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAcceptError returns true if the function "WEBEXT.bluetoothSocket.onAcceptError.removeListener" exists.
func HasFuncOffAcceptError() bool {
	return js.True == bindings.HasFuncOffAcceptError()
}

// FuncOffAcceptError returns the function "WEBEXT.bluetoothSocket.onAcceptError.removeListener".
func FuncOffAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)])]) {
	bindings.FuncOffAcceptError(
		js.Pointer(&fn),
	)
	return
}

// OffAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.removeListener" directly.
func OffAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void) {
	bindings.CallOffAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAcceptError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAcceptError returns true if the function "WEBEXT.bluetoothSocket.onAcceptError.hasListener" exists.
func HasFuncHasOnAcceptError() bool {
	return js.True == bindings.HasFuncHasOnAcceptError()
}

// FuncHasOnAcceptError returns the function "WEBEXT.bluetoothSocket.onAcceptError.hasListener".
func FuncHasOnAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)]) bool]) {
	bindings.FuncHasOnAcceptError(
		js.Pointer(&fn),
	)
	return
}

// HasOnAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.hasListener" directly.
func HasOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret bool) {
	bindings.CallHasOnAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAcceptError calls the function "WEBEXT.bluetoothSocket.onAcceptError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAcceptError(
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

// HasFuncOnReceive returns true if the function "WEBEXT.bluetoothSocket.onReceive.addListener" exists.
func HasFuncOnReceive() bool {
	return js.True == bindings.HasFuncOnReceive()
}

// FuncOnReceive returns the function "WEBEXT.bluetoothSocket.onReceive.addListener".
func FuncOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOnReceive(
		js.Pointer(&fn),
	)
	return
}

// OnReceive calls the function "WEBEXT.bluetoothSocket.onReceive.addListener" directly.
func OnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceive calls the function "WEBEXT.bluetoothSocket.onReceive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceive returns true if the function "WEBEXT.bluetoothSocket.onReceive.removeListener" exists.
func HasFuncOffReceive() bool {
	return js.True == bindings.HasFuncOffReceive()
}

// FuncOffReceive returns the function "WEBEXT.bluetoothSocket.onReceive.removeListener".
func FuncOffReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOffReceive(
		js.Pointer(&fn),
	)
	return
}

// OffReceive calls the function "WEBEXT.bluetoothSocket.onReceive.removeListener" directly.
func OffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOffReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceive calls the function "WEBEXT.bluetoothSocket.onReceive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceive returns true if the function "WEBEXT.bluetoothSocket.onReceive.hasListener" exists.
func HasFuncHasOnReceive() bool {
	return js.True == bindings.HasFuncHasOnReceive()
}

// FuncHasOnReceive returns the function "WEBEXT.bluetoothSocket.onReceive.hasListener".
func FuncHasOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)]) bool]) {
	bindings.FuncHasOnReceive(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceive calls the function "WEBEXT.bluetoothSocket.onReceive.hasListener" directly.
func HasOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret bool) {
	bindings.CallHasOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceive calls the function "WEBEXT.bluetoothSocket.onReceive.hasListener"
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

// HasFuncOnReceiveError returns true if the function "WEBEXT.bluetoothSocket.onReceiveError.addListener" exists.
func HasFuncOnReceiveError() bool {
	return js.True == bindings.HasFuncOnReceiveError()
}

// FuncOnReceiveError returns the function "WEBEXT.bluetoothSocket.onReceiveError.addListener".
func FuncOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OnReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.addListener" directly.
func OnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceiveError returns true if the function "WEBEXT.bluetoothSocket.onReceiveError.removeListener" exists.
func HasFuncOffReceiveError() bool {
	return js.True == bindings.HasFuncOffReceiveError()
}

// FuncOffReceiveError returns the function "WEBEXT.bluetoothSocket.onReceiveError.removeListener".
func FuncOffReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOffReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OffReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.removeListener" directly.
func OffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOffReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceiveError returns true if the function "WEBEXT.bluetoothSocket.onReceiveError.hasListener" exists.
func HasFuncHasOnReceiveError() bool {
	return js.True == bindings.HasFuncHasOnReceiveError()
}

// FuncHasOnReceiveError returns the function "WEBEXT.bluetoothSocket.onReceiveError.hasListener".
func FuncHasOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)]) bool]) {
	bindings.FuncHasOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.hasListener" directly.
func HasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool) {
	bindings.CallHasOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceiveError calls the function "WEBEXT.bluetoothSocket.onReceiveError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSend returns true if the function "WEBEXT.bluetoothSocket.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.bluetoothSocket.send".
func FuncSend() (fn js.Func[func(socketId int32, data js.ArrayBuffer) js.Promise[js.Number[int32]]]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.bluetoothSocket.send" directly.
func Send(socketId int32, data js.ArrayBuffer) (ret js.Promise[js.Number[int32]]) {
	bindings.CallSend(
		js.Pointer(&ret),
		int32(socketId),
		data.Ref(),
	)

	return
}

// TrySend calls the function "WEBEXT.bluetoothSocket.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(socketId int32, data js.ArrayBuffer) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		data.Ref(),
	)

	return
}

// HasFuncSetPaused returns true if the function "WEBEXT.bluetoothSocket.setPaused" exists.
func HasFuncSetPaused() bool {
	return js.True == bindings.HasFuncSetPaused()
}

// FuncSetPaused returns the function "WEBEXT.bluetoothSocket.setPaused".
func FuncSetPaused() (fn js.Func[func(socketId int32, paused bool) js.Promise[js.Void]]) {
	bindings.FuncSetPaused(
		js.Pointer(&fn),
	)
	return
}

// SetPaused calls the function "WEBEXT.bluetoothSocket.setPaused" directly.
func SetPaused(socketId int32, paused bool) (ret js.Promise[js.Void]) {
	bindings.CallSetPaused(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(paused)),
	)

	return
}

// TrySetPaused calls the function "WEBEXT.bluetoothSocket.setPaused"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPaused(socketId int32, paused bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPaused(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Bool(bool(paused)),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.bluetoothSocket.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.bluetoothSocket.update".
func FuncUpdate() (fn js.Func[func(socketId int32, properties SocketProperties) js.Promise[js.Void]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.bluetoothSocket.update" directly.
func Update(socketId int32, properties SocketProperties) (ret js.Promise[js.Void]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&properties),
	)

	return
}

// TryUpdate calls the function "WEBEXT.bluetoothSocket.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(socketId int32, properties SocketProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(socketId),
		js.Pointer(&properties),
	)

	return
}
