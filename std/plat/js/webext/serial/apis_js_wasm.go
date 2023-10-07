// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package serial

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/serial/bindings"
)

type ClearBreakCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn ClearBreakCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ClearBreakCallbackFunc) DispatchCallback(
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

type ClearBreakCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *ClearBreakCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ClearBreakCallback[T]) DispatchCallback(
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

type ConnectCallbackFunc func(this js.Ref, connectionInfo *ConnectionInfo) js.Ref

func (fn ConnectCallbackFunc) Register() js.Func[func(connectionInfo *ConnectionInfo)] {
	return js.RegisterCallback[func(connectionInfo *ConnectionInfo)](
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
	var arg0 ConnectionInfo
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

type ConnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, connectionInfo *ConnectionInfo) js.Ref
	Arg T
}

func (cb *ConnectCallback[T]) Register() js.Func[func(connectionInfo *ConnectionInfo)] {
	return js.RegisterCallback[func(connectionInfo *ConnectionInfo)](
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
	var arg0 ConnectionInfo
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

type DataBits uint32

const (
	_ DataBits = iota

	DataBits_SEVEN
	DataBits_EIGHT
)

func (DataBits) FromRef(str js.Ref) DataBits {
	return DataBits(bindings.ConstOfDataBits(str))
}

func (x DataBits) String() (string, bool) {
	switch x {
	case DataBits_SEVEN:
		return "seven", true
	case DataBits_EIGHT:
		return "eight", true
	default:
		return "", false
	}
}

type ParityBit uint32

const (
	_ ParityBit = iota

	ParityBit_NO
	ParityBit_ODD
	ParityBit_EVEN
)

func (ParityBit) FromRef(str js.Ref) ParityBit {
	return ParityBit(bindings.ConstOfParityBit(str))
}

func (x ParityBit) String() (string, bool) {
	switch x {
	case ParityBit_NO:
		return "no", true
	case ParityBit_ODD:
		return "odd", true
	case ParityBit_EVEN:
		return "even", true
	default:
		return "", false
	}
}

type StopBits uint32

const (
	_ StopBits = iota

	StopBits_ONE
	StopBits_TWO
)

func (StopBits) FromRef(str js.Ref) StopBits {
	return StopBits(bindings.ConstOfStopBits(str))
}

func (x StopBits) String() (string, bool) {
	switch x {
	case StopBits_ONE:
		return "one", true
	case StopBits_TWO:
		return "two", true
	default:
		return "", false
	}
}

type ConnectionInfo struct {
	// ConnectionId is "ConnectionInfo.connectionId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConnectionId MUST be set to true to make this field effective.
	ConnectionId int32
	// Paused is "ConnectionInfo.paused"
	//
	// Optional
	//
	// NOTE: FFI_USE_Paused MUST be set to true to make this field effective.
	Paused bool
	// Persistent is "ConnectionInfo.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool
	// Name is "ConnectionInfo.name"
	//
	// Optional
	Name js.String
	// BufferSize is "ConnectionInfo.bufferSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_BufferSize MUST be set to true to make this field effective.
	BufferSize int32
	// ReceiveTimeout is "ConnectionInfo.receiveTimeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReceiveTimeout MUST be set to true to make this field effective.
	ReceiveTimeout int32
	// SendTimeout is "ConnectionInfo.sendTimeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendTimeout MUST be set to true to make this field effective.
	SendTimeout int32
	// Bitrate is "ConnectionInfo.bitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bitrate MUST be set to true to make this field effective.
	Bitrate int32
	// DataBits is "ConnectionInfo.dataBits"
	//
	// Optional
	DataBits DataBits
	// ParityBit is "ConnectionInfo.parityBit"
	//
	// Optional
	ParityBit ParityBit
	// StopBits is "ConnectionInfo.stopBits"
	//
	// Optional
	StopBits StopBits
	// CtsFlowControl is "ConnectionInfo.ctsFlowControl"
	//
	// Optional
	//
	// NOTE: FFI_USE_CtsFlowControl MUST be set to true to make this field effective.
	CtsFlowControl bool

	FFI_USE_ConnectionId   bool // for ConnectionId.
	FFI_USE_Paused         bool // for Paused.
	FFI_USE_Persistent     bool // for Persistent.
	FFI_USE_BufferSize     bool // for BufferSize.
	FFI_USE_ReceiveTimeout bool // for ReceiveTimeout.
	FFI_USE_SendTimeout    bool // for SendTimeout.
	FFI_USE_Bitrate        bool // for Bitrate.
	FFI_USE_CtsFlowControl bool // for CtsFlowControl.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectionInfo with all fields set.
func (p ConnectionInfo) FromRef(ref js.Ref) ConnectionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectionInfo in the application heap.
func (p ConnectionInfo) New() js.Ref {
	return bindings.ConnectionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectionInfo) UpdateFrom(ref js.Ref) {
	bindings.ConnectionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectionInfo) Update(ref js.Ref) {
	bindings.ConnectionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type ConnectionOptions struct {
	// Persistent is "ConnectionOptions.persistent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Persistent MUST be set to true to make this field effective.
	Persistent bool
	// Name is "ConnectionOptions.name"
	//
	// Optional
	Name js.String
	// BufferSize is "ConnectionOptions.bufferSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_BufferSize MUST be set to true to make this field effective.
	BufferSize int32
	// Bitrate is "ConnectionOptions.bitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bitrate MUST be set to true to make this field effective.
	Bitrate int32
	// DataBits is "ConnectionOptions.dataBits"
	//
	// Optional
	DataBits DataBits
	// ParityBit is "ConnectionOptions.parityBit"
	//
	// Optional
	ParityBit ParityBit
	// StopBits is "ConnectionOptions.stopBits"
	//
	// Optional
	StopBits StopBits
	// CtsFlowControl is "ConnectionOptions.ctsFlowControl"
	//
	// Optional
	//
	// NOTE: FFI_USE_CtsFlowControl MUST be set to true to make this field effective.
	CtsFlowControl bool
	// ReceiveTimeout is "ConnectionOptions.receiveTimeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_ReceiveTimeout MUST be set to true to make this field effective.
	ReceiveTimeout int32
	// SendTimeout is "ConnectionOptions.sendTimeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_SendTimeout MUST be set to true to make this field effective.
	SendTimeout int32

	FFI_USE_Persistent     bool // for Persistent.
	FFI_USE_BufferSize     bool // for BufferSize.
	FFI_USE_Bitrate        bool // for Bitrate.
	FFI_USE_CtsFlowControl bool // for CtsFlowControl.
	FFI_USE_ReceiveTimeout bool // for ReceiveTimeout.
	FFI_USE_SendTimeout    bool // for SendTimeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ConnectionOptions with all fields set.
func (p ConnectionOptions) FromRef(ref js.Ref) ConnectionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ConnectionOptions in the application heap.
func (p ConnectionOptions) New() js.Ref {
	return bindings.ConnectionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ConnectionOptions) UpdateFrom(ref js.Ref) {
	bindings.ConnectionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ConnectionOptions) Update(ref js.Ref) {
	bindings.ConnectionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ConnectionOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type DeviceControlSignals struct {
	// Dcd is "DeviceControlSignals.dcd"
	//
	// Optional
	//
	// NOTE: FFI_USE_Dcd MUST be set to true to make this field effective.
	Dcd bool
	// Cts is "DeviceControlSignals.cts"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cts MUST be set to true to make this field effective.
	Cts bool
	// Ri is "DeviceControlSignals.ri"
	//
	// Optional
	//
	// NOTE: FFI_USE_Ri MUST be set to true to make this field effective.
	Ri bool
	// Dsr is "DeviceControlSignals.dsr"
	//
	// Optional
	//
	// NOTE: FFI_USE_Dsr MUST be set to true to make this field effective.
	Dsr bool

	FFI_USE_Dcd bool // for Dcd.
	FFI_USE_Cts bool // for Cts.
	FFI_USE_Ri  bool // for Ri.
	FFI_USE_Dsr bool // for Dsr.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceControlSignals with all fields set.
func (p DeviceControlSignals) FromRef(ref js.Ref) DeviceControlSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceControlSignals in the application heap.
func (p DeviceControlSignals) New() js.Ref {
	return bindings.DeviceControlSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceControlSignals) UpdateFrom(ref js.Ref) {
	bindings.DeviceControlSignalsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceControlSignals) Update(ref js.Ref) {
	bindings.DeviceControlSignalsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceControlSignals) FreeMembers(recursive bool) {
}

type DeviceInfo struct {
	// Path is "DeviceInfo.path"
	//
	// Optional
	Path js.String
	// VendorId is "DeviceInfo.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "DeviceInfo.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// DisplayName is "DeviceInfo.displayName"
	//
	// Optional
	DisplayName js.String

	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceInfo with all fields set.
func (p DeviceInfo) FromRef(ref js.Ref) DeviceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceInfo in the application heap.
func (p DeviceInfo) New() js.Ref {
	return bindings.DeviceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceInfo) UpdateFrom(ref js.Ref) {
	bindings.DeviceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceInfo) Update(ref js.Ref) {
	bindings.DeviceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Path.Ref(),
		p.DisplayName.Ref(),
	)
	p.Path = p.Path.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type DisconnectCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn DisconnectCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisconnectCallbackFunc) DispatchCallback(
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

type DisconnectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *DisconnectCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisconnectCallback[T]) DispatchCallback(
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

type FlushCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn FlushCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FlushCallbackFunc) DispatchCallback(
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

type FlushCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *FlushCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FlushCallback[T]) DispatchCallback(
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

type GetConnectionsCallbackFunc func(this js.Ref, connectionInfos js.Array[ConnectionInfo]) js.Ref

func (fn GetConnectionsCallbackFunc) Register() js.Func[func(connectionInfos js.Array[ConnectionInfo])] {
	return js.RegisterCallback[func(connectionInfos js.Array[ConnectionInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetConnectionsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ConnectionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetConnectionsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, connectionInfos js.Array[ConnectionInfo]) js.Ref
	Arg T
}

func (cb *GetConnectionsCallback[T]) Register() js.Func[func(connectionInfos js.Array[ConnectionInfo])] {
	return js.RegisterCallback[func(connectionInfos js.Array[ConnectionInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetConnectionsCallback[T]) DispatchCallback(
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

		js.Array[ConnectionInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetControlSignalsCallbackFunc func(this js.Ref, signals *DeviceControlSignals) js.Ref

func (fn GetControlSignalsCallbackFunc) Register() js.Func[func(signals *DeviceControlSignals)] {
	return js.RegisterCallback[func(signals *DeviceControlSignals)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetControlSignalsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceControlSignals
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

type GetControlSignalsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, signals *DeviceControlSignals) js.Ref
	Arg T
}

func (cb *GetControlSignalsCallback[T]) Register() js.Func[func(signals *DeviceControlSignals)] {
	return js.RegisterCallback[func(signals *DeviceControlSignals)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetControlSignalsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DeviceControlSignals
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

type GetDevicesCallbackFunc func(this js.Ref, ports js.Array[DeviceInfo]) js.Ref

func (fn GetDevicesCallbackFunc) Register() js.Func[func(ports js.Array[DeviceInfo])] {
	return js.RegisterCallback[func(ports js.Array[DeviceInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, ports js.Array[DeviceInfo]) js.Ref
	Arg T
}

func (cb *GetDevicesCallback[T]) Register() js.Func[func(ports js.Array[DeviceInfo])] {
	return js.RegisterCallback[func(ports js.Array[DeviceInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDevicesCallback[T]) DispatchCallback(
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

		js.Array[DeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetInfoCallbackFunc func(this js.Ref, connectionInfo *ConnectionInfo) js.Ref

func (fn GetInfoCallbackFunc) Register() js.Func[func(connectionInfo *ConnectionInfo)] {
	return js.RegisterCallback[func(connectionInfo *ConnectionInfo)](
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
	var arg0 ConnectionInfo
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
	Fn  func(arg T, this js.Ref, connectionInfo *ConnectionInfo) js.Ref
	Arg T
}

func (cb *GetInfoCallback[T]) Register() js.Func[func(connectionInfo *ConnectionInfo)] {
	return js.RegisterCallback[func(connectionInfo *ConnectionInfo)](
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
	var arg0 ConnectionInfo
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

type HostControlSignals struct {
	// Dtr is "HostControlSignals.dtr"
	//
	// Optional
	//
	// NOTE: FFI_USE_Dtr MUST be set to true to make this field effective.
	Dtr bool
	// Rts is "HostControlSignals.rts"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rts MUST be set to true to make this field effective.
	Rts bool

	FFI_USE_Dtr bool // for Dtr.
	FFI_USE_Rts bool // for Rts.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HostControlSignals with all fields set.
func (p HostControlSignals) FromRef(ref js.Ref) HostControlSignals {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HostControlSignals in the application heap.
func (p HostControlSignals) New() js.Ref {
	return bindings.HostControlSignalsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HostControlSignals) UpdateFrom(ref js.Ref) {
	bindings.HostControlSignalsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HostControlSignals) Update(ref js.Ref) {
	bindings.HostControlSignalsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HostControlSignals) FreeMembers(recursive bool) {
}

type ReceiveError uint32

const (
	_ ReceiveError = iota

	ReceiveError_DISCONNECTED
	ReceiveError_TIMEOUT
	ReceiveError_DEVICE_LOST
	ReceiveError_BREAK
	ReceiveError_FRAME_ERROR
	ReceiveError_OVERRUN
	ReceiveError_BUFFER_OVERFLOW
	ReceiveError_PARITY_ERROR
	ReceiveError_SYSTEM_ERROR
)

func (ReceiveError) FromRef(str js.Ref) ReceiveError {
	return ReceiveError(bindings.ConstOfReceiveError(str))
}

func (x ReceiveError) String() (string, bool) {
	switch x {
	case ReceiveError_DISCONNECTED:
		return "disconnected", true
	case ReceiveError_TIMEOUT:
		return "timeout", true
	case ReceiveError_DEVICE_LOST:
		return "device_lost", true
	case ReceiveError_BREAK:
		return "break", true
	case ReceiveError_FRAME_ERROR:
		return "frame_error", true
	case ReceiveError_OVERRUN:
		return "overrun", true
	case ReceiveError_BUFFER_OVERFLOW:
		return "buffer_overflow", true
	case ReceiveError_PARITY_ERROR:
		return "parity_error", true
	case ReceiveError_SYSTEM_ERROR:
		return "system_error", true
	default:
		return "", false
	}
}

type ReceiveErrorInfo struct {
	// ConnectionId is "ReceiveErrorInfo.connectionId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConnectionId MUST be set to true to make this field effective.
	ConnectionId int32
	// Error is "ReceiveErrorInfo.error"
	//
	// Optional
	Error ReceiveError

	FFI_USE_ConnectionId bool // for ConnectionId.

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
	// ConnectionId is "ReceiveInfo.connectionId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConnectionId MUST be set to true to make this field effective.
	ConnectionId int32
	// Data is "ReceiveInfo.data"
	//
	// Optional
	Data js.ArrayBuffer

	FFI_USE_ConnectionId bool // for ConnectionId.

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

type SendError uint32

const (
	_ SendError = iota

	SendError_DISCONNECTED
	SendError_PENDING
	SendError_TIMEOUT
	SendError_SYSTEM_ERROR
)

func (SendError) FromRef(str js.Ref) SendError {
	return SendError(bindings.ConstOfSendError(str))
}

func (x SendError) String() (string, bool) {
	switch x {
	case SendError_DISCONNECTED:
		return "disconnected", true
	case SendError_PENDING:
		return "pending", true
	case SendError_TIMEOUT:
		return "timeout", true
	case SendError_SYSTEM_ERROR:
		return "system_error", true
	default:
		return "", false
	}
}

type SendInfo struct {
	// BytesSent is "SendInfo.bytesSent"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesSent MUST be set to true to make this field effective.
	BytesSent int32
	// Error is "SendInfo.error"
	//
	// Optional
	Error SendError

	FFI_USE_BytesSent bool // for BytesSent.

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

type SetBreakCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn SetBreakCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetBreakCallbackFunc) DispatchCallback(
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

type SetBreakCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *SetBreakCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetBreakCallback[T]) DispatchCallback(
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

type SetControlSignalsCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn SetControlSignalsCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetControlSignalsCallbackFunc) DispatchCallback(
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

type SetControlSignalsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *SetControlSignalsCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetControlSignalsCallback[T]) DispatchCallback(
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

type UpdateCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn UpdateCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateCallbackFunc) DispatchCallback(
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

type UpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *UpdateCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateCallback[T]) DispatchCallback(
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

// HasFuncClearBreak returns true if the function "WEBEXT.serial.clearBreak" exists.
func HasFuncClearBreak() bool {
	return js.True == bindings.HasFuncClearBreak()
}

// FuncClearBreak returns the function "WEBEXT.serial.clearBreak".
func FuncClearBreak() (fn js.Func[func(connectionId int32) js.Promise[js.Boolean]]) {
	bindings.FuncClearBreak(
		js.Pointer(&fn),
	)
	return
}

// ClearBreak calls the function "WEBEXT.serial.clearBreak" directly.
func ClearBreak(connectionId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallClearBreak(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryClearBreak calls the function "WEBEXT.serial.clearBreak"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearBreak(connectionId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearBreak(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncConnect returns true if the function "WEBEXT.serial.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.serial.connect".
func FuncConnect() (fn js.Func[func(path js.String, options ConnectionOptions) js.Promise[ConnectionInfo]]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.serial.connect" directly.
func Connect(path js.String, options ConnectionOptions) (ret js.Promise[ConnectionInfo]) {
	bindings.CallConnect(
		js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConnect calls the function "WEBEXT.serial.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(path js.String, options ConnectionOptions) (ret js.Promise[ConnectionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.serial.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.serial.disconnect".
func FuncDisconnect() (fn js.Func[func(connectionId int32) js.Promise[js.Boolean]]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.serial.disconnect" directly.
func Disconnect(connectionId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.serial.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(connectionId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncFlush returns true if the function "WEBEXT.serial.flush" exists.
func HasFuncFlush() bool {
	return js.True == bindings.HasFuncFlush()
}

// FuncFlush returns the function "WEBEXT.serial.flush".
func FuncFlush() (fn js.Func[func(connectionId int32) js.Promise[js.Boolean]]) {
	bindings.FuncFlush(
		js.Pointer(&fn),
	)
	return
}

// Flush calls the function "WEBEXT.serial.flush" directly.
func Flush(connectionId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallFlush(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryFlush calls the function "WEBEXT.serial.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFlush(connectionId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFlush(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncGetConnections returns true if the function "WEBEXT.serial.getConnections" exists.
func HasFuncGetConnections() bool {
	return js.True == bindings.HasFuncGetConnections()
}

// FuncGetConnections returns the function "WEBEXT.serial.getConnections".
func FuncGetConnections() (fn js.Func[func() js.Promise[js.Array[ConnectionInfo]]]) {
	bindings.FuncGetConnections(
		js.Pointer(&fn),
	)
	return
}

// GetConnections calls the function "WEBEXT.serial.getConnections" directly.
func GetConnections() (ret js.Promise[js.Array[ConnectionInfo]]) {
	bindings.CallGetConnections(
		js.Pointer(&ret),
	)

	return
}

// TryGetConnections calls the function "WEBEXT.serial.getConnections"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetConnections() (ret js.Promise[js.Array[ConnectionInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetConnections(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetControlSignals returns true if the function "WEBEXT.serial.getControlSignals" exists.
func HasFuncGetControlSignals() bool {
	return js.True == bindings.HasFuncGetControlSignals()
}

// FuncGetControlSignals returns the function "WEBEXT.serial.getControlSignals".
func FuncGetControlSignals() (fn js.Func[func(connectionId int32) js.Promise[DeviceControlSignals]]) {
	bindings.FuncGetControlSignals(
		js.Pointer(&fn),
	)
	return
}

// GetControlSignals calls the function "WEBEXT.serial.getControlSignals" directly.
func GetControlSignals(connectionId int32) (ret js.Promise[DeviceControlSignals]) {
	bindings.CallGetControlSignals(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryGetControlSignals calls the function "WEBEXT.serial.getControlSignals"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetControlSignals(connectionId int32) (ret js.Promise[DeviceControlSignals], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetControlSignals(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncGetDevices returns true if the function "WEBEXT.serial.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.serial.getDevices".
func FuncGetDevices() (fn js.Func[func() js.Promise[js.Array[DeviceInfo]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.serial.getDevices" directly.
func GetDevices() (ret js.Promise[js.Array[DeviceInfo]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.serial.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices() (ret js.Promise[js.Array[DeviceInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.serial.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.serial.getInfo".
func FuncGetInfo() (fn js.Func[func(connectionId int32) js.Promise[ConnectionInfo]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.serial.getInfo" directly.
func GetInfo(connectionId int32) (ret js.Promise[ConnectionInfo]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.serial.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo(connectionId int32) (ret js.Promise[ConnectionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
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

// HasFuncOnReceive returns true if the function "WEBEXT.serial.onReceive.addListener" exists.
func HasFuncOnReceive() bool {
	return js.True == bindings.HasFuncOnReceive()
}

// FuncOnReceive returns the function "WEBEXT.serial.onReceive.addListener".
func FuncOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOnReceive(
		js.Pointer(&fn),
	)
	return
}

// OnReceive calls the function "WEBEXT.serial.onReceive.addListener" directly.
func OnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceive calls the function "WEBEXT.serial.onReceive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceive returns true if the function "WEBEXT.serial.onReceive.removeListener" exists.
func HasFuncOffReceive() bool {
	return js.True == bindings.HasFuncOffReceive()
}

// FuncOffReceive returns the function "WEBEXT.serial.onReceive.removeListener".
func FuncOffReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)])]) {
	bindings.FuncOffReceive(
		js.Pointer(&fn),
	)
	return
}

// OffReceive calls the function "WEBEXT.serial.onReceive.removeListener" directly.
func OffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void) {
	bindings.CallOffReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceive calls the function "WEBEXT.serial.onReceive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceive(callback js.Func[func(info *ReceiveInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceive returns true if the function "WEBEXT.serial.onReceive.hasListener" exists.
func HasFuncHasOnReceive() bool {
	return js.True == bindings.HasFuncHasOnReceive()
}

// FuncHasOnReceive returns the function "WEBEXT.serial.onReceive.hasListener".
func FuncHasOnReceive() (fn js.Func[func(callback js.Func[func(info *ReceiveInfo)]) bool]) {
	bindings.FuncHasOnReceive(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceive calls the function "WEBEXT.serial.onReceive.hasListener" directly.
func HasOnReceive(callback js.Func[func(info *ReceiveInfo)]) (ret bool) {
	bindings.CallHasOnReceive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceive calls the function "WEBEXT.serial.onReceive.hasListener"
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

// HasFuncOnReceiveError returns true if the function "WEBEXT.serial.onReceiveError.addListener" exists.
func HasFuncOnReceiveError() bool {
	return js.True == bindings.HasFuncOnReceiveError()
}

// FuncOnReceiveError returns the function "WEBEXT.serial.onReceiveError.addListener".
func FuncOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OnReceiveError calls the function "WEBEXT.serial.onReceiveError.addListener" directly.
func OnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnReceiveError calls the function "WEBEXT.serial.onReceiveError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffReceiveError returns true if the function "WEBEXT.serial.onReceiveError.removeListener" exists.
func HasFuncOffReceiveError() bool {
	return js.True == bindings.HasFuncOffReceiveError()
}

// FuncOffReceiveError returns the function "WEBEXT.serial.onReceiveError.removeListener".
func FuncOffReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)])]) {
	bindings.FuncOffReceiveError(
		js.Pointer(&fn),
	)
	return
}

// OffReceiveError calls the function "WEBEXT.serial.onReceiveError.removeListener" directly.
func OffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void) {
	bindings.CallOffReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffReceiveError calls the function "WEBEXT.serial.onReceiveError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnReceiveError returns true if the function "WEBEXT.serial.onReceiveError.hasListener" exists.
func HasFuncHasOnReceiveError() bool {
	return js.True == bindings.HasFuncHasOnReceiveError()
}

// FuncHasOnReceiveError returns the function "WEBEXT.serial.onReceiveError.hasListener".
func FuncHasOnReceiveError() (fn js.Func[func(callback js.Func[func(info *ReceiveErrorInfo)]) bool]) {
	bindings.FuncHasOnReceiveError(
		js.Pointer(&fn),
	)
	return
}

// HasOnReceiveError calls the function "WEBEXT.serial.onReceiveError.hasListener" directly.
func HasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool) {
	bindings.CallHasOnReceiveError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnReceiveError calls the function "WEBEXT.serial.onReceiveError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnReceiveError(callback js.Func[func(info *ReceiveErrorInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnReceiveError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSend returns true if the function "WEBEXT.serial.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.serial.send".
func FuncSend() (fn js.Func[func(connectionId int32, data js.ArrayBuffer) js.Promise[SendInfo]]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.serial.send" directly.
func Send(connectionId int32, data js.ArrayBuffer) (ret js.Promise[SendInfo]) {
	bindings.CallSend(
		js.Pointer(&ret),
		int32(connectionId),
		data.Ref(),
	)

	return
}

// TrySend calls the function "WEBEXT.serial.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(connectionId int32, data js.ArrayBuffer) (ret js.Promise[SendInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		data.Ref(),
	)

	return
}

// HasFuncSetBreak returns true if the function "WEBEXT.serial.setBreak" exists.
func HasFuncSetBreak() bool {
	return js.True == bindings.HasFuncSetBreak()
}

// FuncSetBreak returns the function "WEBEXT.serial.setBreak".
func FuncSetBreak() (fn js.Func[func(connectionId int32) js.Promise[js.Boolean]]) {
	bindings.FuncSetBreak(
		js.Pointer(&fn),
	)
	return
}

// SetBreak calls the function "WEBEXT.serial.setBreak" directly.
func SetBreak(connectionId int32) (ret js.Promise[js.Boolean]) {
	bindings.CallSetBreak(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TrySetBreak calls the function "WEBEXT.serial.setBreak"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetBreak(connectionId int32) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetBreak(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncSetControlSignals returns true if the function "WEBEXT.serial.setControlSignals" exists.
func HasFuncSetControlSignals() bool {
	return js.True == bindings.HasFuncSetControlSignals()
}

// FuncSetControlSignals returns the function "WEBEXT.serial.setControlSignals".
func FuncSetControlSignals() (fn js.Func[func(connectionId int32, signals HostControlSignals) js.Promise[js.Boolean]]) {
	bindings.FuncSetControlSignals(
		js.Pointer(&fn),
	)
	return
}

// SetControlSignals calls the function "WEBEXT.serial.setControlSignals" directly.
func SetControlSignals(connectionId int32, signals HostControlSignals) (ret js.Promise[js.Boolean]) {
	bindings.CallSetControlSignals(
		js.Pointer(&ret),
		int32(connectionId),
		js.Pointer(&signals),
	)

	return
}

// TrySetControlSignals calls the function "WEBEXT.serial.setControlSignals"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetControlSignals(connectionId int32, signals HostControlSignals) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetControlSignals(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		js.Pointer(&signals),
	)

	return
}

// HasFuncSetPaused returns true if the function "WEBEXT.serial.setPaused" exists.
func HasFuncSetPaused() bool {
	return js.True == bindings.HasFuncSetPaused()
}

// FuncSetPaused returns the function "WEBEXT.serial.setPaused".
func FuncSetPaused() (fn js.Func[func(connectionId int32, paused bool) js.Promise[js.Void]]) {
	bindings.FuncSetPaused(
		js.Pointer(&fn),
	)
	return
}

// SetPaused calls the function "WEBEXT.serial.setPaused" directly.
func SetPaused(connectionId int32, paused bool) (ret js.Promise[js.Void]) {
	bindings.CallSetPaused(
		js.Pointer(&ret),
		int32(connectionId),
		js.Bool(bool(paused)),
	)

	return
}

// TrySetPaused calls the function "WEBEXT.serial.setPaused"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPaused(connectionId int32, paused bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPaused(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		js.Bool(bool(paused)),
	)

	return
}

// HasFuncUpdate returns true if the function "WEBEXT.serial.update" exists.
func HasFuncUpdate() bool {
	return js.True == bindings.HasFuncUpdate()
}

// FuncUpdate returns the function "WEBEXT.serial.update".
func FuncUpdate() (fn js.Func[func(connectionId int32, options ConnectionOptions) js.Promise[js.Boolean]]) {
	bindings.FuncUpdate(
		js.Pointer(&fn),
	)
	return
}

// Update calls the function "WEBEXT.serial.update" directly.
func Update(connectionId int32, options ConnectionOptions) (ret js.Promise[js.Boolean]) {
	bindings.CallUpdate(
		js.Pointer(&ret),
		int32(connectionId),
		js.Pointer(&options),
	)

	return
}

// TryUpdate calls the function "WEBEXT.serial.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdate(connectionId int32, options ConnectionOptions) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdate(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		js.Pointer(&options),
	)

	return
}
