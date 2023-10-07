// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package storage

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/system/storage/bindings"
)

type EjectDeviceCallbackFunc func(this js.Ref, result EjectDeviceResultCode) js.Ref

func (fn EjectDeviceCallbackFunc) Register() js.Func[func(result EjectDeviceResultCode)] {
	return js.RegisterCallback[func(result EjectDeviceResultCode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EjectDeviceCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		EjectDeviceResultCode(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EjectDeviceCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result EjectDeviceResultCode) js.Ref
	Arg T
}

func (cb *EjectDeviceCallback[T]) Register() js.Func[func(result EjectDeviceResultCode)] {
	return js.RegisterCallback[func(result EjectDeviceResultCode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EjectDeviceCallback[T]) DispatchCallback(
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

		EjectDeviceResultCode(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EjectDeviceResultCode uint32

const (
	_ EjectDeviceResultCode = iota

	EjectDeviceResultCode_SUCCESS
	EjectDeviceResultCode_IN_USE
	EjectDeviceResultCode_NO_SUCH_DEVICE
	EjectDeviceResultCode_FAILURE
)

func (EjectDeviceResultCode) FromRef(str js.Ref) EjectDeviceResultCode {
	return EjectDeviceResultCode(bindings.ConstOfEjectDeviceResultCode(str))
}

func (x EjectDeviceResultCode) String() (string, bool) {
	switch x {
	case EjectDeviceResultCode_SUCCESS:
		return "success", true
	case EjectDeviceResultCode_IN_USE:
		return "in_use", true
	case EjectDeviceResultCode_NO_SUCH_DEVICE:
		return "no_such_device", true
	case EjectDeviceResultCode_FAILURE:
		return "failure", true
	default:
		return "", false
	}
}

type GetAvailableCapacityCallbackFunc func(this js.Ref, info *StorageAvailableCapacityInfo) js.Ref

func (fn GetAvailableCapacityCallbackFunc) Register() js.Func[func(info *StorageAvailableCapacityInfo)] {
	return js.RegisterCallback[func(info *StorageAvailableCapacityInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAvailableCapacityCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageAvailableCapacityInfo
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

type GetAvailableCapacityCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *StorageAvailableCapacityInfo) js.Ref
	Arg T
}

func (cb *GetAvailableCapacityCallback[T]) Register() js.Func[func(info *StorageAvailableCapacityInfo)] {
	return js.RegisterCallback[func(info *StorageAvailableCapacityInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAvailableCapacityCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageAvailableCapacityInfo
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

type StorageAvailableCapacityInfo struct {
	// Id is "StorageAvailableCapacityInfo.id"
	//
	// Optional
	Id js.String
	// AvailableCapacity is "StorageAvailableCapacityInfo.availableCapacity"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableCapacity MUST be set to true to make this field effective.
	AvailableCapacity float64

	FFI_USE_AvailableCapacity bool // for AvailableCapacity.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageAvailableCapacityInfo with all fields set.
func (p StorageAvailableCapacityInfo) FromRef(ref js.Ref) StorageAvailableCapacityInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageAvailableCapacityInfo in the application heap.
func (p StorageAvailableCapacityInfo) New() js.Ref {
	return bindings.StorageAvailableCapacityInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageAvailableCapacityInfo) UpdateFrom(ref js.Ref) {
	bindings.StorageAvailableCapacityInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageAvailableCapacityInfo) Update(ref js.Ref) {
	bindings.StorageAvailableCapacityInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageAvailableCapacityInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type StorageInfoCallbackFunc func(this js.Ref, info js.Array[StorageUnitInfo]) js.Ref

func (fn StorageInfoCallbackFunc) Register() js.Func[func(info js.Array[StorageUnitInfo])] {
	return js.RegisterCallback[func(info js.Array[StorageUnitInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StorageInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[StorageUnitInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StorageInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info js.Array[StorageUnitInfo]) js.Ref
	Arg T
}

func (cb *StorageInfoCallback[T]) Register() js.Func[func(info js.Array[StorageUnitInfo])] {
	return js.RegisterCallback[func(info js.Array[StorageUnitInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StorageInfoCallback[T]) DispatchCallback(
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

		js.Array[StorageUnitInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StorageUnitType uint32

const (
	_ StorageUnitType = iota

	StorageUnitType_FIXED
	StorageUnitType_REMOVABLE
	StorageUnitType_UNKNOWN
)

func (StorageUnitType) FromRef(str js.Ref) StorageUnitType {
	return StorageUnitType(bindings.ConstOfStorageUnitType(str))
}

func (x StorageUnitType) String() (string, bool) {
	switch x {
	case StorageUnitType_FIXED:
		return "fixed", true
	case StorageUnitType_REMOVABLE:
		return "removable", true
	case StorageUnitType_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type StorageUnitInfo struct {
	// Id is "StorageUnitInfo.id"
	//
	// Optional
	Id js.String
	// Name is "StorageUnitInfo.name"
	//
	// Optional
	Name js.String
	// Type is "StorageUnitInfo.type"
	//
	// Optional
	Type StorageUnitType
	// Capacity is "StorageUnitInfo.capacity"
	//
	// Optional
	//
	// NOTE: FFI_USE_Capacity MUST be set to true to make this field effective.
	Capacity float64

	FFI_USE_Capacity bool // for Capacity.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageUnitInfo with all fields set.
func (p StorageUnitInfo) FromRef(ref js.Ref) StorageUnitInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageUnitInfo in the application heap.
func (p StorageUnitInfo) New() js.Ref {
	return bindings.StorageUnitInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageUnitInfo) UpdateFrom(ref js.Ref) {
	bindings.StorageUnitInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageUnitInfo) Update(ref js.Ref) {
	bindings.StorageUnitInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageUnitInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

// HasFuncEjectDevice returns true if the function "WEBEXT.system.storage.ejectDevice" exists.
func HasFuncEjectDevice() bool {
	return js.True == bindings.HasFuncEjectDevice()
}

// FuncEjectDevice returns the function "WEBEXT.system.storage.ejectDevice".
func FuncEjectDevice() (fn js.Func[func(id js.String) js.Promise[EjectDeviceResultCode]]) {
	bindings.FuncEjectDevice(
		js.Pointer(&fn),
	)
	return
}

// EjectDevice calls the function "WEBEXT.system.storage.ejectDevice" directly.
func EjectDevice(id js.String) (ret js.Promise[EjectDeviceResultCode]) {
	bindings.CallEjectDevice(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryEjectDevice calls the function "WEBEXT.system.storage.ejectDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEjectDevice(id js.String) (ret js.Promise[EjectDeviceResultCode], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEjectDevice(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetAvailableCapacity returns true if the function "WEBEXT.system.storage.getAvailableCapacity" exists.
func HasFuncGetAvailableCapacity() bool {
	return js.True == bindings.HasFuncGetAvailableCapacity()
}

// FuncGetAvailableCapacity returns the function "WEBEXT.system.storage.getAvailableCapacity".
func FuncGetAvailableCapacity() (fn js.Func[func(id js.String) js.Promise[StorageAvailableCapacityInfo]]) {
	bindings.FuncGetAvailableCapacity(
		js.Pointer(&fn),
	)
	return
}

// GetAvailableCapacity calls the function "WEBEXT.system.storage.getAvailableCapacity" directly.
func GetAvailableCapacity(id js.String) (ret js.Promise[StorageAvailableCapacityInfo]) {
	bindings.CallGetAvailableCapacity(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryGetAvailableCapacity calls the function "WEBEXT.system.storage.getAvailableCapacity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAvailableCapacity(id js.String) (ret js.Promise[StorageAvailableCapacityInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAvailableCapacity(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.system.storage.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.system.storage.getInfo".
func FuncGetInfo() (fn js.Func[func() js.Promise[js.Array[StorageUnitInfo]]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.system.storage.getInfo" directly.
func GetInfo() (ret js.Promise[js.Array[StorageUnitInfo]]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.system.storage.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo() (ret js.Promise[js.Array[StorageUnitInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnAttachedEventCallbackFunc func(this js.Ref, info *StorageUnitInfo) js.Ref

func (fn OnAttachedEventCallbackFunc) Register() js.Func[func(info *StorageUnitInfo)] {
	return js.RegisterCallback[func(info *StorageUnitInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAttachedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageUnitInfo
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

type OnAttachedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *StorageUnitInfo) js.Ref
	Arg T
}

func (cb *OnAttachedEventCallback[T]) Register() js.Func[func(info *StorageUnitInfo)] {
	return js.RegisterCallback[func(info *StorageUnitInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAttachedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StorageUnitInfo
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

// HasFuncOnAttached returns true if the function "WEBEXT.system.storage.onAttached.addListener" exists.
func HasFuncOnAttached() bool {
	return js.True == bindings.HasFuncOnAttached()
}

// FuncOnAttached returns the function "WEBEXT.system.storage.onAttached.addListener".
func FuncOnAttached() (fn js.Func[func(callback js.Func[func(info *StorageUnitInfo)])]) {
	bindings.FuncOnAttached(
		js.Pointer(&fn),
	)
	return
}

// OnAttached calls the function "WEBEXT.system.storage.onAttached.addListener" directly.
func OnAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret js.Void) {
	bindings.CallOnAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAttached calls the function "WEBEXT.system.storage.onAttached.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAttached returns true if the function "WEBEXT.system.storage.onAttached.removeListener" exists.
func HasFuncOffAttached() bool {
	return js.True == bindings.HasFuncOffAttached()
}

// FuncOffAttached returns the function "WEBEXT.system.storage.onAttached.removeListener".
func FuncOffAttached() (fn js.Func[func(callback js.Func[func(info *StorageUnitInfo)])]) {
	bindings.FuncOffAttached(
		js.Pointer(&fn),
	)
	return
}

// OffAttached calls the function "WEBEXT.system.storage.onAttached.removeListener" directly.
func OffAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret js.Void) {
	bindings.CallOffAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAttached calls the function "WEBEXT.system.storage.onAttached.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAttached returns true if the function "WEBEXT.system.storage.onAttached.hasListener" exists.
func HasFuncHasOnAttached() bool {
	return js.True == bindings.HasFuncHasOnAttached()
}

// FuncHasOnAttached returns the function "WEBEXT.system.storage.onAttached.hasListener".
func FuncHasOnAttached() (fn js.Func[func(callback js.Func[func(info *StorageUnitInfo)]) bool]) {
	bindings.FuncHasOnAttached(
		js.Pointer(&fn),
	)
	return
}

// HasOnAttached calls the function "WEBEXT.system.storage.onAttached.hasListener" directly.
func HasOnAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret bool) {
	bindings.CallHasOnAttached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAttached calls the function "WEBEXT.system.storage.onAttached.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAttached(callback js.Func[func(info *StorageUnitInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAttached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDetachedEventCallbackFunc func(this js.Ref, id js.String) js.Ref

func (fn OnDetachedEventCallbackFunc) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDetachedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDetachedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String) js.Ref
	Arg T
}

func (cb *OnDetachedEventCallback[T]) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDetachedEventCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDetached returns true if the function "WEBEXT.system.storage.onDetached.addListener" exists.
func HasFuncOnDetached() bool {
	return js.True == bindings.HasFuncOnDetached()
}

// FuncOnDetached returns the function "WEBEXT.system.storage.onDetached.addListener".
func FuncOnDetached() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOnDetached(
		js.Pointer(&fn),
	)
	return
}

// OnDetached calls the function "WEBEXT.system.storage.onDetached.addListener" directly.
func OnDetached(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOnDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDetached calls the function "WEBEXT.system.storage.onDetached.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDetached(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDetached returns true if the function "WEBEXT.system.storage.onDetached.removeListener" exists.
func HasFuncOffDetached() bool {
	return js.True == bindings.HasFuncOffDetached()
}

// FuncOffDetached returns the function "WEBEXT.system.storage.onDetached.removeListener".
func FuncOffDetached() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOffDetached(
		js.Pointer(&fn),
	)
	return
}

// OffDetached calls the function "WEBEXT.system.storage.onDetached.removeListener" directly.
func OffDetached(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOffDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDetached calls the function "WEBEXT.system.storage.onDetached.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDetached(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDetached returns true if the function "WEBEXT.system.storage.onDetached.hasListener" exists.
func HasFuncHasOnDetached() bool {
	return js.True == bindings.HasFuncHasOnDetached()
}

// FuncHasOnDetached returns the function "WEBEXT.system.storage.onDetached.hasListener".
func FuncHasOnDetached() (fn js.Func[func(callback js.Func[func(id js.String)]) bool]) {
	bindings.FuncHasOnDetached(
		js.Pointer(&fn),
	)
	return
}

// HasOnDetached calls the function "WEBEXT.system.storage.onDetached.hasListener" directly.
func HasOnDetached(callback js.Func[func(id js.String)]) (ret bool) {
	bindings.CallHasOnDetached(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDetached calls the function "WEBEXT.system.storage.onDetached.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDetached(callback js.Func[func(id js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDetached(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}
