// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package hid

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/hid/bindings"
)

type ConnectCallbackFunc func(this js.Ref, connection *HidConnectInfo) js.Ref

func (fn ConnectCallbackFunc) Register() js.Func[func(connection *HidConnectInfo)] {
	return js.RegisterCallback[func(connection *HidConnectInfo)](
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
	var arg0 HidConnectInfo
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
	Fn  func(arg T, this js.Ref, connection *HidConnectInfo) js.Ref
	Arg T
}

func (cb *ConnectCallback[T]) Register() js.Func[func(connection *HidConnectInfo)] {
	return js.RegisterCallback[func(connection *HidConnectInfo)](
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
	var arg0 HidConnectInfo
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

type HidConnectInfo struct {
	// ConnectionId is "HidConnectInfo.connectionId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConnectionId MUST be set to true to make this field effective.
	ConnectionId int32

	FFI_USE_ConnectionId bool // for ConnectionId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HidConnectInfo with all fields set.
func (p HidConnectInfo) FromRef(ref js.Ref) HidConnectInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HidConnectInfo in the application heap.
func (p HidConnectInfo) New() js.Ref {
	return bindings.HidConnectInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HidConnectInfo) UpdateFrom(ref js.Ref) {
	bindings.HidConnectInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HidConnectInfo) Update(ref js.Ref) {
	bindings.HidConnectInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HidConnectInfo) FreeMembers(recursive bool) {
}

type DeviceFilter struct {
	// VendorId is "DeviceFilter.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "DeviceFilter.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// UsagePage is "DeviceFilter.usagePage"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsagePage MUST be set to true to make this field effective.
	UsagePage int32
	// Usage is "DeviceFilter.usage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
	Usage int32

	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.
	FFI_USE_UsagePage bool // for UsagePage.
	FFI_USE_Usage     bool // for Usage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceFilter with all fields set.
func (p DeviceFilter) FromRef(ref js.Ref) DeviceFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceFilter in the application heap.
func (p DeviceFilter) New() js.Ref {
	return bindings.DeviceFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceFilter) UpdateFrom(ref js.Ref) {
	bindings.DeviceFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceFilter) Update(ref js.Ref) {
	bindings.DeviceFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceFilter) FreeMembers(recursive bool) {
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

type GetDevicesCallbackFunc func(this js.Ref, devices js.Array[HidDeviceInfo]) js.Ref

func (fn GetDevicesCallbackFunc) Register() js.Func[func(devices js.Array[HidDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[HidDeviceInfo])](
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

		js.Array[HidDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, devices js.Array[HidDeviceInfo]) js.Ref
	Arg T
}

func (cb *GetDevicesCallback[T]) Register() js.Func[func(devices js.Array[HidDeviceInfo])] {
	return js.RegisterCallback[func(devices js.Array[HidDeviceInfo])](
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

		js.Array[HidDeviceInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type HidCollectionInfo struct {
	// UsagePage is "HidCollectionInfo.usagePage"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsagePage MUST be set to true to make this field effective.
	UsagePage int32
	// Usage is "HidCollectionInfo.usage"
	//
	// Optional
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
	Usage int32
	// ReportIds is "HidCollectionInfo.reportIds"
	//
	// Optional
	ReportIds js.Array[int32]

	FFI_USE_UsagePage bool // for UsagePage.
	FFI_USE_Usage     bool // for Usage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HidCollectionInfo with all fields set.
func (p HidCollectionInfo) FromRef(ref js.Ref) HidCollectionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HidCollectionInfo in the application heap.
func (p HidCollectionInfo) New() js.Ref {
	return bindings.HidCollectionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HidCollectionInfo) UpdateFrom(ref js.Ref) {
	bindings.HidCollectionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HidCollectionInfo) Update(ref js.Ref) {
	bindings.HidCollectionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HidCollectionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ReportIds.Ref(),
	)
	p.ReportIds = p.ReportIds.FromRef(js.Undefined)
}

type HidDeviceInfo struct {
	// DeviceId is "HidDeviceInfo.deviceId"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceId MUST be set to true to make this field effective.
	DeviceId int32
	// VendorId is "HidDeviceInfo.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "HidDeviceInfo.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// ProductName is "HidDeviceInfo.productName"
	//
	// Optional
	ProductName js.String
	// SerialNumber is "HidDeviceInfo.serialNumber"
	//
	// Optional
	SerialNumber js.String
	// Collections is "HidDeviceInfo.collections"
	//
	// Optional
	Collections js.Array[HidCollectionInfo]
	// MaxInputReportSize is "HidDeviceInfo.maxInputReportSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxInputReportSize MUST be set to true to make this field effective.
	MaxInputReportSize int32
	// MaxOutputReportSize is "HidDeviceInfo.maxOutputReportSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxOutputReportSize MUST be set to true to make this field effective.
	MaxOutputReportSize int32
	// MaxFeatureReportSize is "HidDeviceInfo.maxFeatureReportSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxFeatureReportSize MUST be set to true to make this field effective.
	MaxFeatureReportSize int32
	// ReportDescriptor is "HidDeviceInfo.reportDescriptor"
	//
	// Optional
	ReportDescriptor js.ArrayBuffer

	FFI_USE_DeviceId             bool // for DeviceId.
	FFI_USE_VendorId             bool // for VendorId.
	FFI_USE_ProductId            bool // for ProductId.
	FFI_USE_MaxInputReportSize   bool // for MaxInputReportSize.
	FFI_USE_MaxOutputReportSize  bool // for MaxOutputReportSize.
	FFI_USE_MaxFeatureReportSize bool // for MaxFeatureReportSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HidDeviceInfo with all fields set.
func (p HidDeviceInfo) FromRef(ref js.Ref) HidDeviceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HidDeviceInfo in the application heap.
func (p HidDeviceInfo) New() js.Ref {
	return bindings.HidDeviceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HidDeviceInfo) UpdateFrom(ref js.Ref) {
	bindings.HidDeviceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HidDeviceInfo) Update(ref js.Ref) {
	bindings.HidDeviceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HidDeviceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ProductName.Ref(),
		p.SerialNumber.Ref(),
		p.Collections.Ref(),
		p.ReportDescriptor.Ref(),
	)
	p.ProductName = p.ProductName.FromRef(js.Undefined)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
	p.Collections = p.Collections.FromRef(js.Undefined)
	p.ReportDescriptor = p.ReportDescriptor.FromRef(js.Undefined)
}

type GetDevicesOptions struct {
	// VendorId is "GetDevicesOptions.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId int32
	// ProductId is "GetDevicesOptions.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int32
	// Filters is "GetDevicesOptions.filters"
	//
	// Optional
	Filters js.Array[DeviceFilter]

	FFI_USE_VendorId  bool // for VendorId.
	FFI_USE_ProductId bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDevicesOptions with all fields set.
func (p GetDevicesOptions) FromRef(ref js.Ref) GetDevicesOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDevicesOptions in the application heap.
func (p GetDevicesOptions) New() js.Ref {
	return bindings.GetDevicesOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDevicesOptions) UpdateFrom(ref js.Ref) {
	bindings.GetDevicesOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDevicesOptions) Update(ref js.Ref) {
	bindings.GetDevicesOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDevicesOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
}

type ReceiveCallbackFunc func(this js.Ref, reportId int32, data js.ArrayBuffer) js.Ref

func (fn ReceiveCallbackFunc) Register() js.Func[func(reportId int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(reportId int32, data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReceiveCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReceiveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reportId int32, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *ReceiveCallback[T]) Register() js.Func[func(reportId int32, data js.ArrayBuffer)] {
	return js.RegisterCallback[func(reportId int32, data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReceiveCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
		js.ArrayBuffer{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReceiveFeatureReportCallbackFunc func(this js.Ref, data js.ArrayBuffer) js.Ref

func (fn ReceiveFeatureReportCallbackFunc) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ReceiveFeatureReportCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ReceiveFeatureReportCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *ReceiveFeatureReportCallback[T]) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ReceiveFeatureReportCallback[T]) DispatchCallback(
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

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SendCallbackFunc func(this js.Ref) js.Ref

func (fn SendCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SendCallbackFunc) DispatchCallback(
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

type SendCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SendCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SendCallback[T]) DispatchCallback(
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

// HasFuncConnect returns true if the function "WEBEXT.hid.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.hid.connect".
func FuncConnect() (fn js.Func[func(deviceId int32) js.Promise[HidConnectInfo]]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.hid.connect" directly.
func Connect(deviceId int32) (ret js.Promise[HidConnectInfo]) {
	bindings.CallConnect(
		js.Pointer(&ret),
		int32(deviceId),
	)

	return
}

// TryConnect calls the function "WEBEXT.hid.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(deviceId int32) (ret js.Promise[HidConnectInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(deviceId),
	)

	return
}

// HasFuncDisconnect returns true if the function "WEBEXT.hid.disconnect" exists.
func HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncDisconnect()
}

// FuncDisconnect returns the function "WEBEXT.hid.disconnect".
func FuncDisconnect() (fn js.Func[func(connectionId int32) js.Promise[js.Void]]) {
	bindings.FuncDisconnect(
		js.Pointer(&fn),
	)
	return
}

// Disconnect calls the function "WEBEXT.hid.disconnect" directly.
func Disconnect(connectionId int32) (ret js.Promise[js.Void]) {
	bindings.CallDisconnect(
		js.Pointer(&ret),
		int32(connectionId),
	)

	return
}

// TryDisconnect calls the function "WEBEXT.hid.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDisconnect(connectionId int32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDisconnect(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
	)

	return
}

// HasFuncGetDevices returns true if the function "WEBEXT.hid.getDevices" exists.
func HasFuncGetDevices() bool {
	return js.True == bindings.HasFuncGetDevices()
}

// FuncGetDevices returns the function "WEBEXT.hid.getDevices".
func FuncGetDevices() (fn js.Func[func(options GetDevicesOptions) js.Promise[js.Array[HidDeviceInfo]]]) {
	bindings.FuncGetDevices(
		js.Pointer(&fn),
	)
	return
}

// GetDevices calls the function "WEBEXT.hid.getDevices" directly.
func GetDevices(options GetDevicesOptions) (ret js.Promise[js.Array[HidDeviceInfo]]) {
	bindings.CallGetDevices(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetDevices calls the function "WEBEXT.hid.getDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDevices(options GetDevicesOptions) (ret js.Promise[js.Array[HidDeviceInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDevices(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type OnDeviceAddedEventCallbackFunc func(this js.Ref, device *HidDeviceInfo) js.Ref

func (fn OnDeviceAddedEventCallbackFunc) Register() js.Func[func(device *HidDeviceInfo)] {
	return js.RegisterCallback[func(device *HidDeviceInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceAddedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HidDeviceInfo
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

type OnDeviceAddedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, device *HidDeviceInfo) js.Ref
	Arg T
}

func (cb *OnDeviceAddedEventCallback[T]) Register() js.Func[func(device *HidDeviceInfo)] {
	return js.RegisterCallback[func(device *HidDeviceInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceAddedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HidDeviceInfo
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

// HasFuncOnDeviceAdded returns true if the function "WEBEXT.hid.onDeviceAdded.addListener" exists.
func HasFuncOnDeviceAdded() bool {
	return js.True == bindings.HasFuncOnDeviceAdded()
}

// FuncOnDeviceAdded returns the function "WEBEXT.hid.onDeviceAdded.addListener".
func FuncOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *HidDeviceInfo)])]) {
	bindings.FuncOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.addListener" directly.
func OnDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret js.Void) {
	bindings.CallOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceAdded returns true if the function "WEBEXT.hid.onDeviceAdded.removeListener" exists.
func HasFuncOffDeviceAdded() bool {
	return js.True == bindings.HasFuncOffDeviceAdded()
}

// FuncOffDeviceAdded returns the function "WEBEXT.hid.onDeviceAdded.removeListener".
func FuncOffDeviceAdded() (fn js.Func[func(callback js.Func[func(device *HidDeviceInfo)])]) {
	bindings.FuncOffDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.removeListener" directly.
func OffDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret js.Void) {
	bindings.CallOffDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceAdded returns true if the function "WEBEXT.hid.onDeviceAdded.hasListener" exists.
func HasFuncHasOnDeviceAdded() bool {
	return js.True == bindings.HasFuncHasOnDeviceAdded()
}

// FuncHasOnDeviceAdded returns the function "WEBEXT.hid.onDeviceAdded.hasListener".
func FuncHasOnDeviceAdded() (fn js.Func[func(callback js.Func[func(device *HidDeviceInfo)]) bool]) {
	bindings.FuncHasOnDeviceAdded(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.hasListener" directly.
func HasOnDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret bool) {
	bindings.CallHasOnDeviceAdded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceAdded calls the function "WEBEXT.hid.onDeviceAdded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceAdded(callback js.Func[func(device *HidDeviceInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceAdded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDeviceRemovedEventCallbackFunc func(this js.Ref, deviceId int32) js.Ref

func (fn OnDeviceRemovedEventCallbackFunc) Register() js.Func[func(deviceId int32)] {
	return js.RegisterCallback[func(deviceId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDeviceRemovedEventCallbackFunc) DispatchCallback(
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

type OnDeviceRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deviceId int32) js.Ref
	Arg T
}

func (cb *OnDeviceRemovedEventCallback[T]) Register() js.Func[func(deviceId int32)] {
	return js.RegisterCallback[func(deviceId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDeviceRemovedEventCallback[T]) DispatchCallback(
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

// HasFuncOnDeviceRemoved returns true if the function "WEBEXT.hid.onDeviceRemoved.addListener" exists.
func HasFuncOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncOnDeviceRemoved()
}

// FuncOnDeviceRemoved returns the function "WEBEXT.hid.onDeviceRemoved.addListener".
func FuncOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(deviceId int32)])]) {
	bindings.FuncOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.addListener" directly.
func OnDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret js.Void) {
	bindings.CallOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDeviceRemoved returns true if the function "WEBEXT.hid.onDeviceRemoved.removeListener" exists.
func HasFuncOffDeviceRemoved() bool {
	return js.True == bindings.HasFuncOffDeviceRemoved()
}

// FuncOffDeviceRemoved returns the function "WEBEXT.hid.onDeviceRemoved.removeListener".
func FuncOffDeviceRemoved() (fn js.Func[func(callback js.Func[func(deviceId int32)])]) {
	bindings.FuncOffDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.removeListener" directly.
func OffDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret js.Void) {
	bindings.CallOffDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDeviceRemoved returns true if the function "WEBEXT.hid.onDeviceRemoved.hasListener" exists.
func HasFuncHasOnDeviceRemoved() bool {
	return js.True == bindings.HasFuncHasOnDeviceRemoved()
}

// FuncHasOnDeviceRemoved returns the function "WEBEXT.hid.onDeviceRemoved.hasListener".
func FuncHasOnDeviceRemoved() (fn js.Func[func(callback js.Func[func(deviceId int32)]) bool]) {
	bindings.FuncHasOnDeviceRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.hasListener" directly.
func HasOnDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret bool) {
	bindings.CallHasOnDeviceRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDeviceRemoved calls the function "WEBEXT.hid.onDeviceRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDeviceRemoved(callback js.Func[func(deviceId int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDeviceRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncReceive returns true if the function "WEBEXT.hid.receive" exists.
func HasFuncReceive() bool {
	return js.True == bindings.HasFuncReceive()
}

// FuncReceive returns the function "WEBEXT.hid.receive".
func FuncReceive() (fn js.Func[func(connectionId int32, callback js.Func[func(reportId int32, data js.ArrayBuffer)])]) {
	bindings.FuncReceive(
		js.Pointer(&fn),
	)
	return
}

// Receive calls the function "WEBEXT.hid.receive" directly.
func Receive(connectionId int32, callback js.Func[func(reportId int32, data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallReceive(
		js.Pointer(&ret),
		int32(connectionId),
		callback.Ref(),
	)

	return
}

// TryReceive calls the function "WEBEXT.hid.receive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReceive(connectionId int32, callback js.Func[func(reportId int32, data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReceive(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		callback.Ref(),
	)

	return
}

// HasFuncReceiveFeatureReport returns true if the function "WEBEXT.hid.receiveFeatureReport" exists.
func HasFuncReceiveFeatureReport() bool {
	return js.True == bindings.HasFuncReceiveFeatureReport()
}

// FuncReceiveFeatureReport returns the function "WEBEXT.hid.receiveFeatureReport".
func FuncReceiveFeatureReport() (fn js.Func[func(connectionId int32, reportId int32) js.Promise[js.ArrayBuffer]]) {
	bindings.FuncReceiveFeatureReport(
		js.Pointer(&fn),
	)
	return
}

// ReceiveFeatureReport calls the function "WEBEXT.hid.receiveFeatureReport" directly.
func ReceiveFeatureReport(connectionId int32, reportId int32) (ret js.Promise[js.ArrayBuffer]) {
	bindings.CallReceiveFeatureReport(
		js.Pointer(&ret),
		int32(connectionId),
		int32(reportId),
	)

	return
}

// TryReceiveFeatureReport calls the function "WEBEXT.hid.receiveFeatureReport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReceiveFeatureReport(connectionId int32, reportId int32) (ret js.Promise[js.ArrayBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReceiveFeatureReport(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		int32(reportId),
	)

	return
}

// HasFuncSend returns true if the function "WEBEXT.hid.send" exists.
func HasFuncSend() bool {
	return js.True == bindings.HasFuncSend()
}

// FuncSend returns the function "WEBEXT.hid.send".
func FuncSend() (fn js.Func[func(connectionId int32, reportId int32, data js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncSend(
		js.Pointer(&fn),
	)
	return
}

// Send calls the function "WEBEXT.hid.send" directly.
func Send(connectionId int32, reportId int32, data js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallSend(
		js.Pointer(&ret),
		int32(connectionId),
		int32(reportId),
		data.Ref(),
	)

	return
}

// TrySend calls the function "WEBEXT.hid.send"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySend(connectionId int32, reportId int32, data js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySend(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		int32(reportId),
		data.Ref(),
	)

	return
}

// HasFuncSendFeatureReport returns true if the function "WEBEXT.hid.sendFeatureReport" exists.
func HasFuncSendFeatureReport() bool {
	return js.True == bindings.HasFuncSendFeatureReport()
}

// FuncSendFeatureReport returns the function "WEBEXT.hid.sendFeatureReport".
func FuncSendFeatureReport() (fn js.Func[func(connectionId int32, reportId int32, data js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncSendFeatureReport(
		js.Pointer(&fn),
	)
	return
}

// SendFeatureReport calls the function "WEBEXT.hid.sendFeatureReport" directly.
func SendFeatureReport(connectionId int32, reportId int32, data js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallSendFeatureReport(
		js.Pointer(&ret),
		int32(connectionId),
		int32(reportId),
		data.Ref(),
	)

	return
}

// TrySendFeatureReport calls the function "WEBEXT.hid.sendFeatureReport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendFeatureReport(connectionId int32, reportId int32, data js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendFeatureReport(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(connectionId),
		int32(reportId),
		data.Ref(),
	)

	return
}
