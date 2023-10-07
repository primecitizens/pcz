// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tcpserver

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sockets/tcpserver/bindings"
)

type AcceptErrorInfo struct {
	// SocketId is "AcceptErrorInfo.socketId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SocketId MUST be set to true to make this field effective.
	SocketId int32
	// ResultCode is "AcceptErrorInfo.resultCode"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResultCode MUST be set to true to make this field effective.
	ResultCode int32

	FFI_USE_SocketId   bool // for SocketId.
	FFI_USE_ResultCode bool // for ResultCode.

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

	FFI_USE_Persistent bool // for Persistent.

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

// HasFuncClose returns true if the function "WEBEXT.sockets.tcpServer.close" exists.
func HasFuncClose() bool {
	return js.True == bindings.HasFuncClose()
}

// FuncClose returns the function "WEBEXT.sockets.tcpServer.close".
func FuncClose() (fn js.Func[func(socketId int32, callback js.Func[func()])]) {
	bindings.FuncClose(
		js.Pointer(&fn),
	)
	return
}

// Close calls the function "WEBEXT.sockets.tcpServer.close" directly.
func Close(socketId int32, callback js.Func[func()]) (ret js.Void) {
	bindings.CallClose(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryClose calls the function "WEBEXT.sockets.tcpServer.close"
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

// HasFuncCreate returns true if the function "WEBEXT.sockets.tcpServer.create" exists.
func HasFuncCreate() bool {
	return js.True == bindings.HasFuncCreate()
}

// FuncCreate returns the function "WEBEXT.sockets.tcpServer.create".
func FuncCreate() (fn js.Func[func(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)])]) {
	bindings.FuncCreate(
		js.Pointer(&fn),
	)
	return
}

// Create calls the function "WEBEXT.sockets.tcpServer.create" directly.
func Create(properties SocketProperties, callback js.Func[func(createInfo *CreateInfo)]) (ret js.Void) {
	bindings.CallCreate(
		js.Pointer(&ret),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryCreate calls the function "WEBEXT.sockets.tcpServer.create"
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

// HasFuncDisconnect returns true if the function "WEBEXT.sockets.tcpServer.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.sockets.tcpServer.disconnect".
func FuncDisconnect() (fn js.Func[func(socketId int32, callback js.Func[func()])]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.sockets.tcpServer.disconnect" directly.
func Disconnect(socketId int32, callback js.Func[func()]) (ret js.Void) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.sockets.tcpServer.disconnect"
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

// HasFuncGetInfo returns true if the function "WEBEXT.sockets.tcpServer.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.sockets.tcpServer.getInfo".
func FuncGetInfo() (fn js.Func[func(socketId int32, callback js.Func[func(socketInfo *SocketInfo)])]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.sockets.tcpServer.getInfo" directly.
func GetInfo(socketId int32, callback js.Func[func(socketInfo *SocketInfo)]) (ret js.Void) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(socketId),
		callback.Ref(),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.sockets.tcpServer.getInfo"
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

// HasFuncGetSockets returns true if the function "WEBEXT.sockets.tcpServer.getSockets" exists.
func HasFuncGetSockets() bool {
	return js.True == bindings.HasFuncGetSockets()
}

// FuncGetSockets returns the function "WEBEXT.sockets.tcpServer.getSockets".
func FuncGetSockets() (fn js.Func[func(callback js.Func[func(socketInfos js.Array[SocketInfo])])]) {
	bindings.FuncGetSockets(
		js.Pointer(&fn),
	)
	return
}

// GetSockets calls the function "WEBEXT.sockets.tcpServer.getSockets" directly.
func GetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void) {
	bindings.CallGetSockets(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetSockets calls the function "WEBEXT.sockets.tcpServer.getSockets"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSockets(callback js.Func[func(socketInfos js.Array[SocketInfo])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSockets(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncListen returns true if the function "WEBEXT.sockets.tcpServer.listen" exists.
func HasFuncListen() bool {
	return js.True == bindings.HasFuncListen()
}

// FuncListen returns the function "WEBEXT.sockets.tcpServer.listen".
func FuncListen() (fn js.Func[func(socketId int32, address js.String, port int32, backlog int32, callback js.Func[func(result int32)])]) {
	bindings.FuncListen(
		js.Pointer(&fn),
	)
	return
}

// Listen calls the function "WEBEXT.sockets.tcpServer.listen" directly.
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

// TryListen calls the function "WEBEXT.sockets.tcpServer.listen"
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

// HasFuncOnAccept returns true if the function "WEBEXT.sockets.tcpServer.onAccept.addListener" exists.
func HasFuncOnAccept() bool {
	return js.True == bindings.HasFuncOnAccept()
}

// FuncOnAccept returns the function "WEBEXT.sockets.tcpServer.onAccept.addListener".
func FuncOnAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)])]) {
	bindings.FuncOnAccept(
		js.Pointer(&fn),
	)
	return
}

// OnAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.addListener" directly.
func OnAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void) {
	bindings.CallOnAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccept returns true if the function "WEBEXT.sockets.tcpServer.onAccept.removeListener" exists.
func HasFuncOffAccept() bool {
	return js.True == bindings.HasFuncOffAccept()
}

// FuncOffAccept returns the function "WEBEXT.sockets.tcpServer.onAccept.removeListener".
func FuncOffAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)])]) {
	bindings.FuncOffAccept(
		js.Pointer(&fn),
	)
	return
}

// OffAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.removeListener" directly.
func OffAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void) {
	bindings.CallOffAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccept(callback js.Func[func(info *AcceptInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccept(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccept returns true if the function "WEBEXT.sockets.tcpServer.onAccept.hasListener" exists.
func HasFuncHasOnAccept() bool {
	return js.True == bindings.HasFuncHasOnAccept()
}

// FuncHasOnAccept returns the function "WEBEXT.sockets.tcpServer.onAccept.hasListener".
func FuncHasOnAccept() (fn js.Func[func(callback js.Func[func(info *AcceptInfo)]) bool]) {
	bindings.FuncHasOnAccept(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.hasListener" directly.
func HasOnAccept(callback js.Func[func(info *AcceptInfo)]) (ret bool) {
	bindings.CallHasOnAccept(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccept calls the function "WEBEXT.sockets.tcpServer.onAccept.hasListener"
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

// HasFuncOnAcceptError returns true if the function "WEBEXT.sockets.tcpServer.onAcceptError.addListener" exists.
func HasFuncOnAcceptError() bool {
	return js.True == bindings.HasFuncOnAcceptError()
}

// FuncOnAcceptError returns the function "WEBEXT.sockets.tcpServer.onAcceptError.addListener".
func FuncOnAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)])]) {
	bindings.FuncOnAcceptError(
		js.Pointer(&fn),
	)
	return
}

// OnAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.addListener" directly.
func OnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void) {
	bindings.CallOnAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAcceptError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAcceptError returns true if the function "WEBEXT.sockets.tcpServer.onAcceptError.removeListener" exists.
func HasFuncOffAcceptError() bool {
	return js.True == bindings.HasFuncOffAcceptError()
}

// FuncOffAcceptError returns the function "WEBEXT.sockets.tcpServer.onAcceptError.removeListener".
func FuncOffAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)])]) {
	bindings.FuncOffAcceptError(
		js.Pointer(&fn),
	)
	return
}

// OffAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.removeListener" directly.
func OffAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void) {
	bindings.CallOffAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAcceptError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAcceptError returns true if the function "WEBEXT.sockets.tcpServer.onAcceptError.hasListener" exists.
func HasFuncHasOnAcceptError() bool {
	return js.True == bindings.HasFuncHasOnAcceptError()
}

// FuncHasOnAcceptError returns the function "WEBEXT.sockets.tcpServer.onAcceptError.hasListener".
func FuncHasOnAcceptError() (fn js.Func[func(callback js.Func[func(info *AcceptErrorInfo)]) bool]) {
	bindings.FuncHasOnAcceptError(
		js.Pointer(&fn),
	)
	return
}

// HasOnAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.hasListener" directly.
func HasOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret bool) {
	bindings.CallHasOnAcceptError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAcceptError calls the function "WEBEXT.sockets.tcpServer.onAcceptError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAcceptError(callback js.Func[func(info *AcceptErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAcceptError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetPaused returns true if the function "WEBEXT.sockets.tcpServer.setPaused" exists.
func HasFuncSetPaused() bool {
	return js.True == bindings.HasFuncSetPaused()
}

// FuncSetPaused returns the function "WEBEXT.sockets.tcpServer.setPaused".
func FuncSetPaused() (fn js.Func[func(socketId int32, paused bool, callback js.Func[func()])]) {
	bindings.FuncSetPaused(
		js.Pointer(&fn),
	)
	return
}

// SetPaused calls the function "WEBEXT.sockets.tcpServer.setPaused" directly.
func SetPaused(socketId int32, paused bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetPaused(
		js.Pointer(&ret),
		int32(socketId),
		js.Bool(bool(paused)),
		callback.Ref(),
	)

	return
}

// TrySetPaused calls the function "WEBEXT.sockets.tcpServer.setPaused"
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

// HasFuncUpdate returns true if the function "WEBEXT.sockets.tcpServer.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.sockets.tcpServer.update".
func FuncUpdate() (fn js.Func[func(socketId int32, properties SocketProperties, callback js.Func[func()])]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.sockets.tcpServer.update" directly.
func Update(socketId int32, properties SocketProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		int32(socketId),
		js.Pointer(&properties),
		callback.Ref(),
	)

	return
}

// TryUpdate calls the function "WEBEXT.sockets.tcpServer.update"
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
